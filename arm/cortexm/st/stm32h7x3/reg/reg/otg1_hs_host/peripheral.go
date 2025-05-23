package otg1_hs_host

import (
	"unsafe"
	"volatile"
)

var (
	Otg1_hs_host = (*_otg1_hs_host)(unsafe.Pointer(uintptr(0x40040400)))
	Otg2_hs_host = (*_otg1_hs_host)(unsafe.Pointer(uintptr(0x40080400)))
)

type _otg1_hs_host struct {
	Otg_hs_hcfg       registerOtg_hs_hcfgType
	Otg_hs_hfir       registerOtg_hs_hfirType
	Otg_hs_hfnum      registerOtg_hs_hfnumType
	_                 [4]byte
	Otg_hs_hptxsts    registerOtg_hs_hptxstsType
	Otg_hs_haint      registerOtg_hs_haintType
	Otg_hs_haintmsk   registerOtg_hs_haintmskType
	_                 [36]byte
	Otg_hs_hprt       registerOtg_hs_hprtType
	_                 [188]byte
	Otg_hs_hcchar0    registerOtg_hs_hcchar0Type
	_                 [28]byte
	Otg_hs_hcchar1    registerOtg_hs_hcchar1Type
	_                 [28]byte
	Otg_hs_hcchar2    registerOtg_hs_hcchar2Type
	_                 [28]byte
	Otg_hs_hcchar3    registerOtg_hs_hcchar3Type
	_                 [28]byte
	Otg_hs_hcchar4    registerOtg_hs_hcchar4Type
	_                 [28]byte
	Otg_hs_hcchar5    registerOtg_hs_hcchar5Type
	_                 [28]byte
	Otg_hs_hcchar6    registerOtg_hs_hcchar6Type
	_                 [28]byte
	Otg_hs_hcchar7    registerOtg_hs_hcchar7Type
	_                 [28]byte
	Otg_hs_hcchar8    registerOtg_hs_hcchar8Type
	_                 [28]byte
	Otg_hs_hcchar9    registerOtg_hs_hcchar9Type
	_                 [28]byte
	Otg_hs_hcchar10   registerOtg_hs_hcchar10Type
	_                 [28]byte
	Otg_hs_hcchar11   registerOtg_hs_hcchar11Type
	Otg_hs_hcsplt0    registerOtg_hs_hcsplt0Type
	_                 [28]byte
	Otg_hs_hcsplt1    registerOtg_hs_hcsplt1Type
	_                 [28]byte
	Otg_hs_hcsplt2    registerOtg_hs_hcsplt2Type
	_                 [28]byte
	Otg_hs_hcsplt3    registerOtg_hs_hcsplt3Type
	_                 [28]byte
	Otg_hs_hcsplt4    registerOtg_hs_hcsplt4Type
	_                 [28]byte
	Otg_hs_hcsplt5    registerOtg_hs_hcsplt5Type
	_                 [28]byte
	Otg_hs_hcsplt6    registerOtg_hs_hcsplt6Type
	_                 [28]byte
	Otg_hs_hcsplt7    registerOtg_hs_hcsplt7Type
	_                 [28]byte
	Otg_hs_hcsplt8    registerOtg_hs_hcsplt8Type
	_                 [28]byte
	Otg_hs_hcsplt9    registerOtg_hs_hcsplt9Type
	_                 [28]byte
	Otg_hs_hcsplt10   registerOtg_hs_hcsplt10Type
	_                 [28]byte
	Otg_hs_hcsplt11   registerOtg_hs_hcsplt11Type
	Otg_hs_hcint0     registerOtg_hs_hcint0Type
	_                 [28]byte
	Otg_hs_hcint1     registerOtg_hs_hcint1Type
	_                 [28]byte
	Otg_hs_hcint2     registerOtg_hs_hcint2Type
	_                 [28]byte
	Otg_hs_hcint3     registerOtg_hs_hcint3Type
	_                 [28]byte
	Otg_hs_hcint4     registerOtg_hs_hcint4Type
	_                 [28]byte
	Otg_hs_hcint5     registerOtg_hs_hcint5Type
	_                 [28]byte
	Otg_hs_hcint6     registerOtg_hs_hcint6Type
	_                 [28]byte
	Otg_hs_hcint7     registerOtg_hs_hcint7Type
	_                 [28]byte
	Otg_hs_hcint8     registerOtg_hs_hcint8Type
	_                 [28]byte
	Otg_hs_hcint9     registerOtg_hs_hcint9Type
	_                 [28]byte
	Otg_hs_hcint10    registerOtg_hs_hcint10Type
	_                 [28]byte
	Otg_hs_hcint11    registerOtg_hs_hcint11Type
	Otg_hs_hcintmsk0  registerOtg_hs_hcintmsk0Type
	_                 [28]byte
	Otg_hs_hcintmsk1  registerOtg_hs_hcintmsk1Type
	_                 [28]byte
	Otg_hs_hcintmsk2  registerOtg_hs_hcintmsk2Type
	_                 [28]byte
	Otg_hs_hcintmsk3  registerOtg_hs_hcintmsk3Type
	_                 [28]byte
	Otg_hs_hcintmsk4  registerOtg_hs_hcintmsk4Type
	_                 [28]byte
	Otg_hs_hcintmsk5  registerOtg_hs_hcintmsk5Type
	_                 [28]byte
	Otg_hs_hcintmsk6  registerOtg_hs_hcintmsk6Type
	_                 [28]byte
	Otg_hs_hcintmsk7  registerOtg_hs_hcintmsk7Type
	_                 [28]byte
	Otg_hs_hcintmsk8  registerOtg_hs_hcintmsk8Type
	_                 [28]byte
	Otg_hs_hcintmsk9  registerOtg_hs_hcintmsk9Type
	_                 [28]byte
	Otg_hs_hcintmsk10 registerOtg_hs_hcintmsk10Type
	_                 [28]byte
	Otg_hs_hcintmsk11 registerOtg_hs_hcintmsk11Type
	Otg_hs_hctsiz0    registerOtg_hs_hctsiz0Type
	_                 [28]byte
	Otg_hs_hctsiz1    registerOtg_hs_hctsiz1Type
	_                 [28]byte
	Otg_hs_hctsiz2    registerOtg_hs_hctsiz2Type
	_                 [28]byte
	Otg_hs_hctsiz3    registerOtg_hs_hctsiz3Type
	_                 [28]byte
	Otg_hs_hctsiz4    registerOtg_hs_hctsiz4Type
	_                 [28]byte
	Otg_hs_hctsiz5    registerOtg_hs_hctsiz5Type
	_                 [28]byte
	Otg_hs_hctsiz6    registerOtg_hs_hctsiz6Type
	_                 [28]byte
	Otg_hs_hctsiz7    registerOtg_hs_hctsiz7Type
	_                 [28]byte
	Otg_hs_hctsiz8    registerOtg_hs_hctsiz8Type
	_                 [28]byte
	Otg_hs_hctsiz9    registerOtg_hs_hctsiz9Type
	_                 [28]byte
	Otg_hs_hctsiz10   registerOtg_hs_hctsiz10Type
	_                 [28]byte
	Otg_hs_hctsiz11   registerOtg_hs_hctsiz11Type
	Otg_hs_hcdma0     registerOtg_hs_hcdma0Type
	_                 [28]byte
	Otg_hs_hcdma1     registerOtg_hs_hcdma1Type
	_                 [28]byte
	Otg_hs_hcdma2     registerOtg_hs_hcdma2Type
	_                 [28]byte
	Otg_hs_hcdma3     registerOtg_hs_hcdma3Type
	_                 [28]byte
	Otg_hs_hcdma4     registerOtg_hs_hcdma4Type
	_                 [28]byte
	Otg_hs_hcdma5     registerOtg_hs_hcdma5Type
	_                 [28]byte
	Otg_hs_hcdma6     registerOtg_hs_hcdma6Type
	_                 [28]byte
	Otg_hs_hcdma7     registerOtg_hs_hcdma7Type
	_                 [28]byte
	Otg_hs_hcdma8     registerOtg_hs_hcdma8Type
	_                 [28]byte
	Otg_hs_hcdma9     registerOtg_hs_hcdma9Type
	_                 [28]byte
	Otg_hs_hcdma10    registerOtg_hs_hcdma10Type
	_                 [28]byte
	Otg_hs_hcdma11    registerOtg_hs_hcdma11Type
	Otg_hs_hcchar12   registerOtg_hs_hcchar12Type
	Otg_hs_hcsplt12   registerOtg_hs_hcsplt12Type
	Otg_hs_hcint12    registerOtg_hs_hcint12Type
	Otg_hs_hcintmsk12 registerOtg_hs_hcintmsk12Type
	Otg_hs_hctsiz12   registerOtg_hs_hctsiz12Type
	Otg_hs_hcdma12    registerOtg_hs_hcdma12Type
	Otg_hs_hcchar13   registerOtg_hs_hcchar13Type
	Otg_hs_hcsplt13   registerOtg_hs_hcsplt13Type
	Otg_hs_hcint13    registerOtg_hs_hcint13Type
	Otg_hs_hcintmsk13 registerOtg_hs_hcintmsk13Type
	Otg_hs_hctsiz13   registerOtg_hs_hctsiz13Type
	Otg_hs_hcdma13    registerOtg_hs_hcdma13Type
	Otg_hs_hcchar14   registerOtg_hs_hcchar14Type
	Otg_hs_hcsplt14   registerOtg_hs_hcsplt14Type
	Otg_hs_hcint14    registerOtg_hs_hcint14Type
	Otg_hs_hcintmsk14 registerOtg_hs_hcintmsk14Type
	Otg_hs_hctsiz14   registerOtg_hs_hctsiz14Type
	Otg_hs_hcdma14    registerOtg_hs_hcdma14Type
	Otg_hs_hcchar15   registerOtg_hs_hcchar15Type
	Otg_hs_hcsplt15   registerOtg_hs_hcsplt15Type
	Otg_hs_hcint15    registerOtg_hs_hcint15Type
	Otg_hs_hcintmsk15 registerOtg_hs_hcintmsk15Type
	Otg_hs_hctsiz15   registerOtg_hs_hctsiz15Type
	Otg_hs_hcdma15    registerOtg_hs_hcdma15Type
}

// registerOtg_hs_hcfgType OTG_HS host configuration register
type registerOtg_hs_hcfgType uint32

const (
	RegisterOtg_hs_hcfgFieldFslspcsShift = 0
	RegisterOtg_hs_hcfgFieldFslspcsMask  = 0x3
)

// GetFslspcs FS/LS PHY clock select
func (r *registerOtg_hs_hcfgType) GetFslspcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcfgFieldFslspcsMask) >> RegisterOtg_hs_hcfgFieldFslspcsShift)
}

// SetFslspcs FS/LS PHY clock select
func (r *registerOtg_hs_hcfgType) SetFslspcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcfgFieldFslspcsMask)|(uint32(value)<<RegisterOtg_hs_hcfgFieldFslspcsShift))
}

const (
	RegisterOtg_hs_hcfgFieldFslssShift = 2
	RegisterOtg_hs_hcfgFieldFslssMask  = 0x4
)

// GetFslss FS- and LS-only support
func (r *registerOtg_hs_hcfgType) GetFslss() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcfgFieldFslssMask) != 0
}

// registerOtg_hs_hfirType OTG_HS Host frame interval register
type registerOtg_hs_hfirType uint32

const (
	RegisterOtg_hs_hfirFieldFrivlShift = 0
	RegisterOtg_hs_hfirFieldFrivlMask  = 0xffff
)

// GetFrivl Frame interval
func (r *registerOtg_hs_hfirType) GetFrivl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hfirFieldFrivlMask) >> RegisterOtg_hs_hfirFieldFrivlShift)
}

// SetFrivl Frame interval
func (r *registerOtg_hs_hfirType) SetFrivl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hfirFieldFrivlMask)|(uint32(value)<<RegisterOtg_hs_hfirFieldFrivlShift))
}

// registerOtg_hs_hfnumType OTG_HS host frame number/frame time remaining register
type registerOtg_hs_hfnumType uint32

const (
	RegisterOtg_hs_hfnumFieldFrnumShift = 0
	RegisterOtg_hs_hfnumFieldFrnumMask  = 0xffff
)

// GetFrnum Frame number
func (r *registerOtg_hs_hfnumType) GetFrnum() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hfnumFieldFrnumMask) >> RegisterOtg_hs_hfnumFieldFrnumShift)
}

// SetFrnum Frame number
func (r *registerOtg_hs_hfnumType) SetFrnum(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hfnumFieldFrnumMask)|(uint32(value)<<RegisterOtg_hs_hfnumFieldFrnumShift))
}

const (
	RegisterOtg_hs_hfnumFieldFtremShift = 16
	RegisterOtg_hs_hfnumFieldFtremMask  = 0xffff0000
)

// GetFtrem Frame time remaining
func (r *registerOtg_hs_hfnumType) GetFtrem() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hfnumFieldFtremMask) >> RegisterOtg_hs_hfnumFieldFtremShift)
}

// SetFtrem Frame time remaining
func (r *registerOtg_hs_hfnumType) SetFtrem(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hfnumFieldFtremMask)|(uint32(value)<<RegisterOtg_hs_hfnumFieldFtremShift))
}

// registerOtg_hs_hptxstsType OTG_HS_Host periodic transmit FIFO/queue status register
type registerOtg_hs_hptxstsType uint32

const (
	RegisterOtg_hs_hptxstsFieldPtxfsavlShift = 0
	RegisterOtg_hs_hptxstsFieldPtxfsavlMask  = 0xffff
)

// GetPtxfsavl Periodic transmit data FIFO space available
func (r *registerOtg_hs_hptxstsType) GetPtxfsavl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hptxstsFieldPtxfsavlMask) >> RegisterOtg_hs_hptxstsFieldPtxfsavlShift)
}

// SetPtxfsavl Periodic transmit data FIFO space available
func (r *registerOtg_hs_hptxstsType) SetPtxfsavl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hptxstsFieldPtxfsavlMask)|(uint32(value)<<RegisterOtg_hs_hptxstsFieldPtxfsavlShift))
}

const (
	RegisterOtg_hs_hptxstsFieldPtxqsavShift = 16
	RegisterOtg_hs_hptxstsFieldPtxqsavMask  = 0xff0000
)

// GetPtxqsav Periodic transmit request queue space available
func (r *registerOtg_hs_hptxstsType) GetPtxqsav() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hptxstsFieldPtxqsavMask) >> RegisterOtg_hs_hptxstsFieldPtxqsavShift)
}

const (
	RegisterOtg_hs_hptxstsFieldPtxqtopShift = 24
	RegisterOtg_hs_hptxstsFieldPtxqtopMask  = 0xff000000
)

// GetPtxqtop Top of the periodic transmit request queue
func (r *registerOtg_hs_hptxstsType) GetPtxqtop() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hptxstsFieldPtxqtopMask) >> RegisterOtg_hs_hptxstsFieldPtxqtopShift)
}

// registerOtg_hs_haintType OTG_HS Host all channels interrupt register
type registerOtg_hs_haintType uint32

const (
	RegisterOtg_hs_haintFieldHaintShift = 0
	RegisterOtg_hs_haintFieldHaintMask  = 0xffff
)

// GetHaint Channel interrupts
func (r *registerOtg_hs_haintType) GetHaint() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_haintFieldHaintMask) >> RegisterOtg_hs_haintFieldHaintShift)
}

// SetHaint Channel interrupts
func (r *registerOtg_hs_haintType) SetHaint(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_haintFieldHaintMask)|(uint32(value)<<RegisterOtg_hs_haintFieldHaintShift))
}

// registerOtg_hs_haintmskType OTG_HS host all channels interrupt mask register
type registerOtg_hs_haintmskType uint32

const (
	RegisterOtg_hs_haintmskFieldHaintmShift = 0
	RegisterOtg_hs_haintmskFieldHaintmMask  = 0xffff
)

// GetHaintm Channel interrupt mask
func (r *registerOtg_hs_haintmskType) GetHaintm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_haintmskFieldHaintmMask) >> RegisterOtg_hs_haintmskFieldHaintmShift)
}

// SetHaintm Channel interrupt mask
func (r *registerOtg_hs_haintmskType) SetHaintm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_haintmskFieldHaintmMask)|(uint32(value)<<RegisterOtg_hs_haintmskFieldHaintmShift))
}

// registerOtg_hs_hprtType OTG_HS host port control and status register
type registerOtg_hs_hprtType uint32

const (
	RegisterOtg_hs_hprtFieldPcstsShift = 0
	RegisterOtg_hs_hprtFieldPcstsMask  = 0x1
)

// GetPcsts Port connect status
func (r *registerOtg_hs_hprtType) GetPcsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPcstsMask) != 0
}

const (
	RegisterOtg_hs_hprtFieldPcdetShift = 1
	RegisterOtg_hs_hprtFieldPcdetMask  = 0x2
)

// GetPcdet Port connect detected
func (r *registerOtg_hs_hprtType) GetPcdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPcdetMask) != 0
}

// SetPcdet Port connect detected
func (r *registerOtg_hs_hprtType) SetPcdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hprtFieldPcdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hprtFieldPcdetMask)
	}
}

const (
	RegisterOtg_hs_hprtFieldPenaShift = 2
	RegisterOtg_hs_hprtFieldPenaMask  = 0x4
)

// GetPena Port enable
func (r *registerOtg_hs_hprtType) GetPena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPenaMask) != 0
}

// SetPena Port enable
func (r *registerOtg_hs_hprtType) SetPena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hprtFieldPenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hprtFieldPenaMask)
	}
}

const (
	RegisterOtg_hs_hprtFieldPenchngShift = 3
	RegisterOtg_hs_hprtFieldPenchngMask  = 0x8
)

// GetPenchng Port enable/disable change
func (r *registerOtg_hs_hprtType) GetPenchng() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPenchngMask) != 0
}

// SetPenchng Port enable/disable change
func (r *registerOtg_hs_hprtType) SetPenchng(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hprtFieldPenchngMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hprtFieldPenchngMask)
	}
}

const (
	RegisterOtg_hs_hprtFieldPocaShift = 4
	RegisterOtg_hs_hprtFieldPocaMask  = 0x10
)

// GetPoca Port overcurrent active
func (r *registerOtg_hs_hprtType) GetPoca() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPocaMask) != 0
}

const (
	RegisterOtg_hs_hprtFieldPocchngShift = 5
	RegisterOtg_hs_hprtFieldPocchngMask  = 0x20
)

// GetPocchng Port overcurrent change
func (r *registerOtg_hs_hprtType) GetPocchng() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPocchngMask) != 0
}

// SetPocchng Port overcurrent change
func (r *registerOtg_hs_hprtType) SetPocchng(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hprtFieldPocchngMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hprtFieldPocchngMask)
	}
}

const (
	RegisterOtg_hs_hprtFieldPresShift = 6
	RegisterOtg_hs_hprtFieldPresMask  = 0x40
)

// GetPres Port resume
func (r *registerOtg_hs_hprtType) GetPres() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPresMask) != 0
}

// SetPres Port resume
func (r *registerOtg_hs_hprtType) SetPres(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hprtFieldPresMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hprtFieldPresMask)
	}
}

const (
	RegisterOtg_hs_hprtFieldPsuspShift = 7
	RegisterOtg_hs_hprtFieldPsuspMask  = 0x80
)

// GetPsusp Port suspend
func (r *registerOtg_hs_hprtType) GetPsusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPsuspMask) != 0
}

// SetPsusp Port suspend
func (r *registerOtg_hs_hprtType) SetPsusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hprtFieldPsuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hprtFieldPsuspMask)
	}
}

const (
	RegisterOtg_hs_hprtFieldPrstShift = 8
	RegisterOtg_hs_hprtFieldPrstMask  = 0x100
)

// GetPrst Port reset
func (r *registerOtg_hs_hprtType) GetPrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPrstMask) != 0
}

// SetPrst Port reset
func (r *registerOtg_hs_hprtType) SetPrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hprtFieldPrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hprtFieldPrstMask)
	}
}

const (
	RegisterOtg_hs_hprtFieldPlstsShift = 10
	RegisterOtg_hs_hprtFieldPlstsMask  = 0xc00
)

// GetPlsts Port line status
func (r *registerOtg_hs_hprtType) GetPlsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPlstsMask) >> RegisterOtg_hs_hprtFieldPlstsShift)
}

const (
	RegisterOtg_hs_hprtFieldPpwrShift = 12
	RegisterOtg_hs_hprtFieldPpwrMask  = 0x1000
)

// GetPpwr Port power
func (r *registerOtg_hs_hprtType) GetPpwr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPpwrMask) != 0
}

// SetPpwr Port power
func (r *registerOtg_hs_hprtType) SetPpwr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hprtFieldPpwrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hprtFieldPpwrMask)
	}
}

const (
	RegisterOtg_hs_hprtFieldPtctlShift = 13
	RegisterOtg_hs_hprtFieldPtctlMask  = 0x1e000
)

// GetPtctl Port test control
func (r *registerOtg_hs_hprtType) GetPtctl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPtctlMask) >> RegisterOtg_hs_hprtFieldPtctlShift)
}

// SetPtctl Port test control
func (r *registerOtg_hs_hprtType) SetPtctl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hprtFieldPtctlMask)|(uint32(value)<<RegisterOtg_hs_hprtFieldPtctlShift))
}

const (
	RegisterOtg_hs_hprtFieldPspdShift = 17
	RegisterOtg_hs_hprtFieldPspdMask  = 0x60000
)

// GetPspd Port speed
func (r *registerOtg_hs_hprtType) GetPspd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hprtFieldPspdMask) >> RegisterOtg_hs_hprtFieldPspdShift)
}

// registerOtg_hs_hcchar0Type OTG_HS host channel-0 characteristics register
type registerOtg_hs_hcchar0Type uint32

const (
	RegisterOtg_hs_hcchar0FieldMpsizShift = 0
	RegisterOtg_hs_hcchar0FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar0Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar0FieldMpsizMask) >> RegisterOtg_hs_hcchar0FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar0Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar0FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar0FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar0FieldEpnumShift = 11
	RegisterOtg_hs_hcchar0FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar0Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar0FieldEpnumMask) >> RegisterOtg_hs_hcchar0FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar0Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar0FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar0FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar0FieldEpdirShift = 15
	RegisterOtg_hs_hcchar0FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar0Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar0FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar0Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar0FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar0FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar0FieldLsdevShift = 17
	RegisterOtg_hs_hcchar0FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar0Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar0FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar0Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar0FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar0FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar0FieldEptypShift = 18
	RegisterOtg_hs_hcchar0FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar0Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar0FieldEptypMask) >> RegisterOtg_hs_hcchar0FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar0Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar0FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar0FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar0FieldMcShift = 20
	RegisterOtg_hs_hcchar0FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar0Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar0FieldMcMask) >> RegisterOtg_hs_hcchar0FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar0Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar0FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar0FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar0FieldDadShift = 22
	RegisterOtg_hs_hcchar0FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar0Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar0FieldDadMask) >> RegisterOtg_hs_hcchar0FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar0Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar0FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar0FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar0FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar0FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar0Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar0FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar0Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar0FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar0FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar0FieldChdisShift = 30
	RegisterOtg_hs_hcchar0FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar0Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar0FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar0Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar0FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar0FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar0FieldChenaShift = 31
	RegisterOtg_hs_hcchar0FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar0Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar0FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar0Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar0FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar0FieldChenaMask)
	}
}

// registerOtg_hs_hcchar1Type OTG_HS host channel-1 characteristics register
type registerOtg_hs_hcchar1Type uint32

const (
	RegisterOtg_hs_hcchar1FieldMpsizShift = 0
	RegisterOtg_hs_hcchar1FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar1Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar1FieldMpsizMask) >> RegisterOtg_hs_hcchar1FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar1Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar1FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar1FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar1FieldEpnumShift = 11
	RegisterOtg_hs_hcchar1FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar1Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar1FieldEpnumMask) >> RegisterOtg_hs_hcchar1FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar1Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar1FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar1FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar1FieldEpdirShift = 15
	RegisterOtg_hs_hcchar1FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar1Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar1FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar1Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar1FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar1FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar1FieldLsdevShift = 17
	RegisterOtg_hs_hcchar1FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar1Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar1FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar1Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar1FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar1FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar1FieldEptypShift = 18
	RegisterOtg_hs_hcchar1FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar1Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar1FieldEptypMask) >> RegisterOtg_hs_hcchar1FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar1Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar1FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar1FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar1FieldMcShift = 20
	RegisterOtg_hs_hcchar1FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar1Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar1FieldMcMask) >> RegisterOtg_hs_hcchar1FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar1Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar1FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar1FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar1FieldDadShift = 22
	RegisterOtg_hs_hcchar1FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar1Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar1FieldDadMask) >> RegisterOtg_hs_hcchar1FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar1Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar1FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar1FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar1FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar1FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar1Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar1FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar1Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar1FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar1FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar1FieldChdisShift = 30
	RegisterOtg_hs_hcchar1FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar1Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar1FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar1Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar1FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar1FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar1FieldChenaShift = 31
	RegisterOtg_hs_hcchar1FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar1Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar1FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar1Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar1FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar1FieldChenaMask)
	}
}

// registerOtg_hs_hcchar2Type OTG_HS host channel-2 characteristics register
type registerOtg_hs_hcchar2Type uint32

const (
	RegisterOtg_hs_hcchar2FieldMpsizShift = 0
	RegisterOtg_hs_hcchar2FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar2Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar2FieldMpsizMask) >> RegisterOtg_hs_hcchar2FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar2Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar2FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar2FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar2FieldEpnumShift = 11
	RegisterOtg_hs_hcchar2FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar2Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar2FieldEpnumMask) >> RegisterOtg_hs_hcchar2FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar2Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar2FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar2FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar2FieldEpdirShift = 15
	RegisterOtg_hs_hcchar2FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar2Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar2FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar2Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar2FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar2FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar2FieldLsdevShift = 17
	RegisterOtg_hs_hcchar2FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar2Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar2FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar2Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar2FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar2FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar2FieldEptypShift = 18
	RegisterOtg_hs_hcchar2FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar2Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar2FieldEptypMask) >> RegisterOtg_hs_hcchar2FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar2Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar2FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar2FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar2FieldMcShift = 20
	RegisterOtg_hs_hcchar2FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar2Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar2FieldMcMask) >> RegisterOtg_hs_hcchar2FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar2Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar2FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar2FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar2FieldDadShift = 22
	RegisterOtg_hs_hcchar2FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar2Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar2FieldDadMask) >> RegisterOtg_hs_hcchar2FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar2Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar2FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar2FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar2FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar2FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar2Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar2FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar2Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar2FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar2FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar2FieldChdisShift = 30
	RegisterOtg_hs_hcchar2FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar2Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar2FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar2Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar2FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar2FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar2FieldChenaShift = 31
	RegisterOtg_hs_hcchar2FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar2Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar2FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar2Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar2FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar2FieldChenaMask)
	}
}

// registerOtg_hs_hcchar3Type OTG_HS host channel-3 characteristics register
type registerOtg_hs_hcchar3Type uint32

const (
	RegisterOtg_hs_hcchar3FieldMpsizShift = 0
	RegisterOtg_hs_hcchar3FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar3Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar3FieldMpsizMask) >> RegisterOtg_hs_hcchar3FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar3Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar3FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar3FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar3FieldEpnumShift = 11
	RegisterOtg_hs_hcchar3FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar3Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar3FieldEpnumMask) >> RegisterOtg_hs_hcchar3FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar3Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar3FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar3FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar3FieldEpdirShift = 15
	RegisterOtg_hs_hcchar3FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar3Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar3FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar3Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar3FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar3FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar3FieldLsdevShift = 17
	RegisterOtg_hs_hcchar3FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar3Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar3FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar3Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar3FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar3FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar3FieldEptypShift = 18
	RegisterOtg_hs_hcchar3FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar3Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar3FieldEptypMask) >> RegisterOtg_hs_hcchar3FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar3Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar3FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar3FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar3FieldMcShift = 20
	RegisterOtg_hs_hcchar3FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar3Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar3FieldMcMask) >> RegisterOtg_hs_hcchar3FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar3Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar3FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar3FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar3FieldDadShift = 22
	RegisterOtg_hs_hcchar3FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar3Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar3FieldDadMask) >> RegisterOtg_hs_hcchar3FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar3Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar3FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar3FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar3FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar3FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar3Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar3FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar3Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar3FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar3FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar3FieldChdisShift = 30
	RegisterOtg_hs_hcchar3FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar3Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar3FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar3Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar3FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar3FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar3FieldChenaShift = 31
	RegisterOtg_hs_hcchar3FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar3Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar3FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar3Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar3FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar3FieldChenaMask)
	}
}

// registerOtg_hs_hcchar4Type OTG_HS host channel-4 characteristics register
type registerOtg_hs_hcchar4Type uint32

const (
	RegisterOtg_hs_hcchar4FieldMpsizShift = 0
	RegisterOtg_hs_hcchar4FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar4Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar4FieldMpsizMask) >> RegisterOtg_hs_hcchar4FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar4Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar4FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar4FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar4FieldEpnumShift = 11
	RegisterOtg_hs_hcchar4FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar4Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar4FieldEpnumMask) >> RegisterOtg_hs_hcchar4FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar4Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar4FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar4FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar4FieldEpdirShift = 15
	RegisterOtg_hs_hcchar4FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar4Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar4FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar4Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar4FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar4FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar4FieldLsdevShift = 17
	RegisterOtg_hs_hcchar4FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar4Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar4FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar4Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar4FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar4FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar4FieldEptypShift = 18
	RegisterOtg_hs_hcchar4FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar4Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar4FieldEptypMask) >> RegisterOtg_hs_hcchar4FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar4Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar4FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar4FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar4FieldMcShift = 20
	RegisterOtg_hs_hcchar4FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar4Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar4FieldMcMask) >> RegisterOtg_hs_hcchar4FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar4Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar4FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar4FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar4FieldDadShift = 22
	RegisterOtg_hs_hcchar4FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar4Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar4FieldDadMask) >> RegisterOtg_hs_hcchar4FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar4Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar4FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar4FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar4FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar4FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar4Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar4FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar4Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar4FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar4FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar4FieldChdisShift = 30
	RegisterOtg_hs_hcchar4FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar4Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar4FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar4Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar4FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar4FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar4FieldChenaShift = 31
	RegisterOtg_hs_hcchar4FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar4Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar4FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar4Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar4FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar4FieldChenaMask)
	}
}

// registerOtg_hs_hcchar5Type OTG_HS host channel-5 characteristics register
type registerOtg_hs_hcchar5Type uint32

const (
	RegisterOtg_hs_hcchar5FieldMpsizShift = 0
	RegisterOtg_hs_hcchar5FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar5Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar5FieldMpsizMask) >> RegisterOtg_hs_hcchar5FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar5Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar5FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar5FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar5FieldEpnumShift = 11
	RegisterOtg_hs_hcchar5FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar5Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar5FieldEpnumMask) >> RegisterOtg_hs_hcchar5FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar5Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar5FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar5FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar5FieldEpdirShift = 15
	RegisterOtg_hs_hcchar5FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar5Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar5FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar5Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar5FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar5FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar5FieldLsdevShift = 17
	RegisterOtg_hs_hcchar5FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar5Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar5FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar5Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar5FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar5FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar5FieldEptypShift = 18
	RegisterOtg_hs_hcchar5FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar5Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar5FieldEptypMask) >> RegisterOtg_hs_hcchar5FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar5Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar5FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar5FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar5FieldMcShift = 20
	RegisterOtg_hs_hcchar5FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar5Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar5FieldMcMask) >> RegisterOtg_hs_hcchar5FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar5Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar5FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar5FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar5FieldDadShift = 22
	RegisterOtg_hs_hcchar5FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar5Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar5FieldDadMask) >> RegisterOtg_hs_hcchar5FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar5Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar5FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar5FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar5FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar5FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar5Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar5FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar5Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar5FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar5FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar5FieldChdisShift = 30
	RegisterOtg_hs_hcchar5FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar5Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar5FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar5Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar5FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar5FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar5FieldChenaShift = 31
	RegisterOtg_hs_hcchar5FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar5Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar5FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar5Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar5FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar5FieldChenaMask)
	}
}

// registerOtg_hs_hcchar6Type OTG_HS host channel-6 characteristics register
type registerOtg_hs_hcchar6Type uint32

const (
	RegisterOtg_hs_hcchar6FieldMpsizShift = 0
	RegisterOtg_hs_hcchar6FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar6Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar6FieldMpsizMask) >> RegisterOtg_hs_hcchar6FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar6Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar6FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar6FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar6FieldEpnumShift = 11
	RegisterOtg_hs_hcchar6FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar6Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar6FieldEpnumMask) >> RegisterOtg_hs_hcchar6FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar6Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar6FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar6FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar6FieldEpdirShift = 15
	RegisterOtg_hs_hcchar6FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar6Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar6FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar6Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar6FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar6FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar6FieldLsdevShift = 17
	RegisterOtg_hs_hcchar6FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar6Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar6FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar6Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar6FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar6FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar6FieldEptypShift = 18
	RegisterOtg_hs_hcchar6FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar6Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar6FieldEptypMask) >> RegisterOtg_hs_hcchar6FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar6Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar6FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar6FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar6FieldMcShift = 20
	RegisterOtg_hs_hcchar6FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar6Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar6FieldMcMask) >> RegisterOtg_hs_hcchar6FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar6Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar6FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar6FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar6FieldDadShift = 22
	RegisterOtg_hs_hcchar6FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar6Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar6FieldDadMask) >> RegisterOtg_hs_hcchar6FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar6Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar6FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar6FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar6FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar6FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar6Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar6FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar6Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar6FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar6FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar6FieldChdisShift = 30
	RegisterOtg_hs_hcchar6FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar6Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar6FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar6Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar6FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar6FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar6FieldChenaShift = 31
	RegisterOtg_hs_hcchar6FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar6Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar6FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar6Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar6FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar6FieldChenaMask)
	}
}

// registerOtg_hs_hcchar7Type OTG_HS host channel-7 characteristics register
type registerOtg_hs_hcchar7Type uint32

const (
	RegisterOtg_hs_hcchar7FieldMpsizShift = 0
	RegisterOtg_hs_hcchar7FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar7Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar7FieldMpsizMask) >> RegisterOtg_hs_hcchar7FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar7Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar7FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar7FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar7FieldEpnumShift = 11
	RegisterOtg_hs_hcchar7FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar7Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar7FieldEpnumMask) >> RegisterOtg_hs_hcchar7FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar7Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar7FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar7FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar7FieldEpdirShift = 15
	RegisterOtg_hs_hcchar7FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar7Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar7FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar7Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar7FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar7FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar7FieldLsdevShift = 17
	RegisterOtg_hs_hcchar7FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar7Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar7FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar7Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar7FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar7FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar7FieldEptypShift = 18
	RegisterOtg_hs_hcchar7FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar7Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar7FieldEptypMask) >> RegisterOtg_hs_hcchar7FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar7Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar7FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar7FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar7FieldMcShift = 20
	RegisterOtg_hs_hcchar7FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar7Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar7FieldMcMask) >> RegisterOtg_hs_hcchar7FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar7Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar7FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar7FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar7FieldDadShift = 22
	RegisterOtg_hs_hcchar7FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar7Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar7FieldDadMask) >> RegisterOtg_hs_hcchar7FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar7Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar7FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar7FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar7FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar7FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar7Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar7FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar7Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar7FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar7FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar7FieldChdisShift = 30
	RegisterOtg_hs_hcchar7FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar7Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar7FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar7Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar7FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar7FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar7FieldChenaShift = 31
	RegisterOtg_hs_hcchar7FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar7Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar7FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar7Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar7FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar7FieldChenaMask)
	}
}

// registerOtg_hs_hcchar8Type OTG_HS host channel-8 characteristics register
type registerOtg_hs_hcchar8Type uint32

const (
	RegisterOtg_hs_hcchar8FieldMpsizShift = 0
	RegisterOtg_hs_hcchar8FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar8Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar8FieldMpsizMask) >> RegisterOtg_hs_hcchar8FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar8Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar8FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar8FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar8FieldEpnumShift = 11
	RegisterOtg_hs_hcchar8FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar8Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar8FieldEpnumMask) >> RegisterOtg_hs_hcchar8FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar8Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar8FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar8FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar8FieldEpdirShift = 15
	RegisterOtg_hs_hcchar8FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar8Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar8FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar8Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar8FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar8FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar8FieldLsdevShift = 17
	RegisterOtg_hs_hcchar8FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar8Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar8FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar8Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar8FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar8FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar8FieldEptypShift = 18
	RegisterOtg_hs_hcchar8FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar8Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar8FieldEptypMask) >> RegisterOtg_hs_hcchar8FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar8Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar8FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar8FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar8FieldMcShift = 20
	RegisterOtg_hs_hcchar8FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar8Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar8FieldMcMask) >> RegisterOtg_hs_hcchar8FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar8Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar8FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar8FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar8FieldDadShift = 22
	RegisterOtg_hs_hcchar8FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar8Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar8FieldDadMask) >> RegisterOtg_hs_hcchar8FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar8Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar8FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar8FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar8FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar8FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar8Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar8FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar8Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar8FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar8FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar8FieldChdisShift = 30
	RegisterOtg_hs_hcchar8FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar8Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar8FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar8Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar8FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar8FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar8FieldChenaShift = 31
	RegisterOtg_hs_hcchar8FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar8Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar8FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar8Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar8FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar8FieldChenaMask)
	}
}

// registerOtg_hs_hcchar9Type OTG_HS host channel-9 characteristics register
type registerOtg_hs_hcchar9Type uint32

const (
	RegisterOtg_hs_hcchar9FieldMpsizShift = 0
	RegisterOtg_hs_hcchar9FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar9Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar9FieldMpsizMask) >> RegisterOtg_hs_hcchar9FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar9Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar9FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar9FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar9FieldEpnumShift = 11
	RegisterOtg_hs_hcchar9FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar9Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar9FieldEpnumMask) >> RegisterOtg_hs_hcchar9FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar9Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar9FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar9FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar9FieldEpdirShift = 15
	RegisterOtg_hs_hcchar9FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar9Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar9FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar9Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar9FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar9FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar9FieldLsdevShift = 17
	RegisterOtg_hs_hcchar9FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar9Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar9FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar9Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar9FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar9FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar9FieldEptypShift = 18
	RegisterOtg_hs_hcchar9FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar9Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar9FieldEptypMask) >> RegisterOtg_hs_hcchar9FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar9Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar9FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar9FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar9FieldMcShift = 20
	RegisterOtg_hs_hcchar9FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar9Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar9FieldMcMask) >> RegisterOtg_hs_hcchar9FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar9Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar9FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar9FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar9FieldDadShift = 22
	RegisterOtg_hs_hcchar9FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar9Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar9FieldDadMask) >> RegisterOtg_hs_hcchar9FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar9Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar9FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar9FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar9FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar9FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar9Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar9FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar9Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar9FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar9FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar9FieldChdisShift = 30
	RegisterOtg_hs_hcchar9FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar9Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar9FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar9Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar9FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar9FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar9FieldChenaShift = 31
	RegisterOtg_hs_hcchar9FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar9Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar9FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar9Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar9FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar9FieldChenaMask)
	}
}

// registerOtg_hs_hcchar10Type OTG_HS host channel-10 characteristics register
type registerOtg_hs_hcchar10Type uint32

const (
	RegisterOtg_hs_hcchar10FieldMpsizShift = 0
	RegisterOtg_hs_hcchar10FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar10Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar10FieldMpsizMask) >> RegisterOtg_hs_hcchar10FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar10Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar10FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar10FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar10FieldEpnumShift = 11
	RegisterOtg_hs_hcchar10FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar10Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar10FieldEpnumMask) >> RegisterOtg_hs_hcchar10FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar10Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar10FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar10FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar10FieldEpdirShift = 15
	RegisterOtg_hs_hcchar10FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar10Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar10FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar10Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar10FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar10FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar10FieldLsdevShift = 17
	RegisterOtg_hs_hcchar10FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar10Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar10FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar10Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar10FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar10FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar10FieldEptypShift = 18
	RegisterOtg_hs_hcchar10FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar10Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar10FieldEptypMask) >> RegisterOtg_hs_hcchar10FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar10Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar10FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar10FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar10FieldMcShift = 20
	RegisterOtg_hs_hcchar10FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar10Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar10FieldMcMask) >> RegisterOtg_hs_hcchar10FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar10Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar10FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar10FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar10FieldDadShift = 22
	RegisterOtg_hs_hcchar10FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar10Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar10FieldDadMask) >> RegisterOtg_hs_hcchar10FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar10Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar10FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar10FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar10FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar10FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar10Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar10FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar10Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar10FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar10FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar10FieldChdisShift = 30
	RegisterOtg_hs_hcchar10FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar10Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar10FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar10Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar10FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar10FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar10FieldChenaShift = 31
	RegisterOtg_hs_hcchar10FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar10Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar10FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar10Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar10FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar10FieldChenaMask)
	}
}

// registerOtg_hs_hcchar11Type OTG_HS host channel-11 characteristics register
type registerOtg_hs_hcchar11Type uint32

const (
	RegisterOtg_hs_hcchar11FieldMpsizShift = 0
	RegisterOtg_hs_hcchar11FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar11Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar11FieldMpsizMask) >> RegisterOtg_hs_hcchar11FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar11Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar11FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar11FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar11FieldEpnumShift = 11
	RegisterOtg_hs_hcchar11FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar11Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar11FieldEpnumMask) >> RegisterOtg_hs_hcchar11FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar11Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar11FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar11FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar11FieldEpdirShift = 15
	RegisterOtg_hs_hcchar11FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar11Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar11FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar11Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar11FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar11FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar11FieldLsdevShift = 17
	RegisterOtg_hs_hcchar11FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar11Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar11FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar11Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar11FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar11FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar11FieldEptypShift = 18
	RegisterOtg_hs_hcchar11FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar11Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar11FieldEptypMask) >> RegisterOtg_hs_hcchar11FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar11Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar11FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar11FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar11FieldMcShift = 20
	RegisterOtg_hs_hcchar11FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar11Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar11FieldMcMask) >> RegisterOtg_hs_hcchar11FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar11Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar11FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar11FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar11FieldDadShift = 22
	RegisterOtg_hs_hcchar11FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar11Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar11FieldDadMask) >> RegisterOtg_hs_hcchar11FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar11Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar11FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar11FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar11FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar11FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar11Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar11FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar11Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar11FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar11FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar11FieldChdisShift = 30
	RegisterOtg_hs_hcchar11FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar11Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar11FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar11Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar11FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar11FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar11FieldChenaShift = 31
	RegisterOtg_hs_hcchar11FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar11Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar11FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar11Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar11FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar11FieldChenaMask)
	}
}

// registerOtg_hs_hcsplt0Type OTG_HS host channel-0 split control register
type registerOtg_hs_hcsplt0Type uint32

const (
	RegisterOtg_hs_hcsplt0FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt0FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt0Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt0FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt0FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt0Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt0FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt0FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt0FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt0FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt0Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt0FieldHubaddrMask) >> RegisterOtg_hs_hcsplt0FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt0Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt0FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt0FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt0FieldXactposShift = 14
	RegisterOtg_hs_hcsplt0FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt0Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt0FieldXactposMask) >> RegisterOtg_hs_hcsplt0FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt0Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt0FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt0FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt0FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt0FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt0Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt0FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt0Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt0FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt0FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt0FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt0FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt0Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt0FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt0Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt0FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt0FieldSplitenMask)
	}
}

// registerOtg_hs_hcsplt1Type OTG_HS host channel-1 split control register
type registerOtg_hs_hcsplt1Type uint32

const (
	RegisterOtg_hs_hcsplt1FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt1FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt1Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt1FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt1FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt1Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt1FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt1FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt1FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt1FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt1Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt1FieldHubaddrMask) >> RegisterOtg_hs_hcsplt1FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt1Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt1FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt1FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt1FieldXactposShift = 14
	RegisterOtg_hs_hcsplt1FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt1Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt1FieldXactposMask) >> RegisterOtg_hs_hcsplt1FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt1Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt1FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt1FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt1FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt1FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt1Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt1FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt1Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt1FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt1FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt1FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt1FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt1Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt1FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt1Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt1FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt1FieldSplitenMask)
	}
}

// registerOtg_hs_hcsplt2Type OTG_HS host channel-2 split control register
type registerOtg_hs_hcsplt2Type uint32

const (
	RegisterOtg_hs_hcsplt2FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt2FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt2Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt2FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt2FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt2Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt2FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt2FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt2FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt2FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt2Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt2FieldHubaddrMask) >> RegisterOtg_hs_hcsplt2FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt2Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt2FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt2FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt2FieldXactposShift = 14
	RegisterOtg_hs_hcsplt2FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt2Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt2FieldXactposMask) >> RegisterOtg_hs_hcsplt2FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt2Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt2FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt2FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt2FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt2FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt2Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt2FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt2Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt2FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt2FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt2FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt2FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt2Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt2FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt2Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt2FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt2FieldSplitenMask)
	}
}

// registerOtg_hs_hcsplt3Type OTG_HS host channel-3 split control register
type registerOtg_hs_hcsplt3Type uint32

const (
	RegisterOtg_hs_hcsplt3FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt3FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt3Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt3FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt3FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt3Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt3FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt3FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt3FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt3FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt3Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt3FieldHubaddrMask) >> RegisterOtg_hs_hcsplt3FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt3Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt3FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt3FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt3FieldXactposShift = 14
	RegisterOtg_hs_hcsplt3FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt3Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt3FieldXactposMask) >> RegisterOtg_hs_hcsplt3FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt3Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt3FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt3FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt3FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt3FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt3Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt3FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt3Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt3FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt3FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt3FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt3FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt3Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt3FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt3Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt3FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt3FieldSplitenMask)
	}
}

// registerOtg_hs_hcsplt4Type OTG_HS host channel-4 split control register
type registerOtg_hs_hcsplt4Type uint32

const (
	RegisterOtg_hs_hcsplt4FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt4FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt4Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt4FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt4FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt4Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt4FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt4FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt4FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt4FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt4Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt4FieldHubaddrMask) >> RegisterOtg_hs_hcsplt4FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt4Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt4FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt4FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt4FieldXactposShift = 14
	RegisterOtg_hs_hcsplt4FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt4Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt4FieldXactposMask) >> RegisterOtg_hs_hcsplt4FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt4Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt4FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt4FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt4FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt4FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt4Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt4FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt4Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt4FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt4FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt4FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt4FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt4Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt4FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt4Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt4FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt4FieldSplitenMask)
	}
}

// registerOtg_hs_hcsplt5Type OTG_HS host channel-5 split control register
type registerOtg_hs_hcsplt5Type uint32

const (
	RegisterOtg_hs_hcsplt5FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt5FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt5Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt5FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt5FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt5Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt5FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt5FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt5FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt5FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt5Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt5FieldHubaddrMask) >> RegisterOtg_hs_hcsplt5FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt5Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt5FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt5FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt5FieldXactposShift = 14
	RegisterOtg_hs_hcsplt5FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt5Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt5FieldXactposMask) >> RegisterOtg_hs_hcsplt5FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt5Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt5FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt5FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt5FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt5FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt5Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt5FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt5Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt5FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt5FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt5FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt5FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt5Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt5FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt5Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt5FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt5FieldSplitenMask)
	}
}

// registerOtg_hs_hcsplt6Type OTG_HS host channel-6 split control register
type registerOtg_hs_hcsplt6Type uint32

const (
	RegisterOtg_hs_hcsplt6FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt6FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt6Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt6FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt6FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt6Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt6FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt6FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt6FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt6FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt6Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt6FieldHubaddrMask) >> RegisterOtg_hs_hcsplt6FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt6Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt6FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt6FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt6FieldXactposShift = 14
	RegisterOtg_hs_hcsplt6FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt6Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt6FieldXactposMask) >> RegisterOtg_hs_hcsplt6FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt6Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt6FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt6FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt6FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt6FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt6Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt6FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt6Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt6FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt6FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt6FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt6FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt6Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt6FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt6Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt6FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt6FieldSplitenMask)
	}
}

// registerOtg_hs_hcsplt7Type OTG_HS host channel-7 split control register
type registerOtg_hs_hcsplt7Type uint32

const (
	RegisterOtg_hs_hcsplt7FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt7FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt7Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt7FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt7FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt7Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt7FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt7FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt7FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt7FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt7Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt7FieldHubaddrMask) >> RegisterOtg_hs_hcsplt7FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt7Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt7FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt7FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt7FieldXactposShift = 14
	RegisterOtg_hs_hcsplt7FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt7Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt7FieldXactposMask) >> RegisterOtg_hs_hcsplt7FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt7Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt7FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt7FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt7FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt7FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt7Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt7FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt7Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt7FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt7FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt7FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt7FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt7Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt7FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt7Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt7FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt7FieldSplitenMask)
	}
}

// registerOtg_hs_hcsplt8Type OTG_HS host channel-8 split control register
type registerOtg_hs_hcsplt8Type uint32

const (
	RegisterOtg_hs_hcsplt8FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt8FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt8Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt8FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt8FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt8Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt8FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt8FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt8FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt8FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt8Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt8FieldHubaddrMask) >> RegisterOtg_hs_hcsplt8FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt8Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt8FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt8FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt8FieldXactposShift = 14
	RegisterOtg_hs_hcsplt8FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt8Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt8FieldXactposMask) >> RegisterOtg_hs_hcsplt8FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt8Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt8FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt8FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt8FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt8FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt8Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt8FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt8Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt8FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt8FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt8FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt8FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt8Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt8FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt8Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt8FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt8FieldSplitenMask)
	}
}

// registerOtg_hs_hcsplt9Type OTG_HS host channel-9 split control register
type registerOtg_hs_hcsplt9Type uint32

const (
	RegisterOtg_hs_hcsplt9FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt9FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt9Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt9FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt9FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt9Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt9FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt9FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt9FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt9FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt9Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt9FieldHubaddrMask) >> RegisterOtg_hs_hcsplt9FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt9Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt9FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt9FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt9FieldXactposShift = 14
	RegisterOtg_hs_hcsplt9FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt9Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt9FieldXactposMask) >> RegisterOtg_hs_hcsplt9FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt9Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt9FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt9FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt9FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt9FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt9Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt9FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt9Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt9FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt9FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt9FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt9FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt9Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt9FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt9Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt9FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt9FieldSplitenMask)
	}
}

// registerOtg_hs_hcsplt10Type OTG_HS host channel-10 split control register
type registerOtg_hs_hcsplt10Type uint32

const (
	RegisterOtg_hs_hcsplt10FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt10FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt10Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt10FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt10FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt10Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt10FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt10FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt10FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt10FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt10Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt10FieldHubaddrMask) >> RegisterOtg_hs_hcsplt10FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt10Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt10FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt10FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt10FieldXactposShift = 14
	RegisterOtg_hs_hcsplt10FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt10Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt10FieldXactposMask) >> RegisterOtg_hs_hcsplt10FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt10Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt10FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt10FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt10FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt10FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt10Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt10FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt10Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt10FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt10FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt10FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt10FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt10Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt10FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt10Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt10FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt10FieldSplitenMask)
	}
}

// registerOtg_hs_hcsplt11Type OTG_HS host channel-11 split control register
type registerOtg_hs_hcsplt11Type uint32

const (
	RegisterOtg_hs_hcsplt11FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt11FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt11Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt11FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt11FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt11Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt11FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt11FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt11FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt11FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt11Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt11FieldHubaddrMask) >> RegisterOtg_hs_hcsplt11FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt11Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt11FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt11FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt11FieldXactposShift = 14
	RegisterOtg_hs_hcsplt11FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt11Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt11FieldXactposMask) >> RegisterOtg_hs_hcsplt11FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt11Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt11FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt11FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt11FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt11FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt11Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt11FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt11Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt11FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt11FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt11FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt11FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt11Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt11FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt11Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt11FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt11FieldSplitenMask)
	}
}

// registerOtg_hs_hcint0Type OTG_HS host channel-11 interrupt register
type registerOtg_hs_hcint0Type uint32

const (
	RegisterOtg_hs_hcint0FieldXfrcShift = 0
	RegisterOtg_hs_hcint0FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint0Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint0FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint0Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint0FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint0FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint0FieldChhShift = 1
	RegisterOtg_hs_hcint0FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint0Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint0FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint0Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint0FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint0FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint0FieldAhberrShift = 2
	RegisterOtg_hs_hcint0FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint0Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint0FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint0Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint0FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint0FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint0FieldStallShift = 3
	RegisterOtg_hs_hcint0FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint0Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint0FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint0Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint0FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint0FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint0FieldNakShift = 4
	RegisterOtg_hs_hcint0FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint0Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint0FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint0Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint0FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint0FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint0FieldAckShift = 5
	RegisterOtg_hs_hcint0FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint0Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint0FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint0Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint0FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint0FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint0FieldNyetShift = 6
	RegisterOtg_hs_hcint0FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint0Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint0FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint0Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint0FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint0FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint0FieldTxerrShift = 7
	RegisterOtg_hs_hcint0FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint0Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint0FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint0Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint0FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint0FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint0FieldBberrShift = 8
	RegisterOtg_hs_hcint0FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint0Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint0FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint0Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint0FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint0FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint0FieldFrmorShift = 9
	RegisterOtg_hs_hcint0FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint0Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint0FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint0Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint0FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint0FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint0FieldDterrShift = 10
	RegisterOtg_hs_hcint0FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint0Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint0FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint0Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint0FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint0FieldDterrMask)
	}
}

// registerOtg_hs_hcint1Type OTG_HS host channel-1 interrupt register
type registerOtg_hs_hcint1Type uint32

const (
	RegisterOtg_hs_hcint1FieldXfrcShift = 0
	RegisterOtg_hs_hcint1FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint1Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint1FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint1Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint1FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint1FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint1FieldChhShift = 1
	RegisterOtg_hs_hcint1FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint1Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint1FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint1Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint1FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint1FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint1FieldAhberrShift = 2
	RegisterOtg_hs_hcint1FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint1Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint1FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint1Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint1FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint1FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint1FieldStallShift = 3
	RegisterOtg_hs_hcint1FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint1Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint1FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint1Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint1FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint1FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint1FieldNakShift = 4
	RegisterOtg_hs_hcint1FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint1Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint1FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint1Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint1FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint1FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint1FieldAckShift = 5
	RegisterOtg_hs_hcint1FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint1Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint1FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint1Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint1FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint1FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint1FieldNyetShift = 6
	RegisterOtg_hs_hcint1FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint1Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint1FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint1Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint1FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint1FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint1FieldTxerrShift = 7
	RegisterOtg_hs_hcint1FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint1Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint1FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint1Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint1FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint1FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint1FieldBberrShift = 8
	RegisterOtg_hs_hcint1FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint1Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint1FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint1Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint1FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint1FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint1FieldFrmorShift = 9
	RegisterOtg_hs_hcint1FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint1Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint1FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint1Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint1FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint1FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint1FieldDterrShift = 10
	RegisterOtg_hs_hcint1FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint1Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint1FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint1Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint1FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint1FieldDterrMask)
	}
}

// registerOtg_hs_hcint2Type OTG_HS host channel-2 interrupt register
type registerOtg_hs_hcint2Type uint32

const (
	RegisterOtg_hs_hcint2FieldXfrcShift = 0
	RegisterOtg_hs_hcint2FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint2Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint2FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint2Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint2FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint2FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint2FieldChhShift = 1
	RegisterOtg_hs_hcint2FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint2Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint2FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint2Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint2FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint2FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint2FieldAhberrShift = 2
	RegisterOtg_hs_hcint2FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint2Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint2FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint2Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint2FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint2FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint2FieldStallShift = 3
	RegisterOtg_hs_hcint2FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint2Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint2FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint2Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint2FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint2FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint2FieldNakShift = 4
	RegisterOtg_hs_hcint2FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint2Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint2FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint2Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint2FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint2FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint2FieldAckShift = 5
	RegisterOtg_hs_hcint2FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint2Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint2FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint2Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint2FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint2FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint2FieldNyetShift = 6
	RegisterOtg_hs_hcint2FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint2Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint2FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint2Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint2FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint2FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint2FieldTxerrShift = 7
	RegisterOtg_hs_hcint2FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint2Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint2FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint2Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint2FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint2FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint2FieldBberrShift = 8
	RegisterOtg_hs_hcint2FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint2Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint2FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint2Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint2FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint2FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint2FieldFrmorShift = 9
	RegisterOtg_hs_hcint2FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint2Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint2FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint2Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint2FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint2FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint2FieldDterrShift = 10
	RegisterOtg_hs_hcint2FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint2Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint2FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint2Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint2FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint2FieldDterrMask)
	}
}

// registerOtg_hs_hcint3Type OTG_HS host channel-3 interrupt register
type registerOtg_hs_hcint3Type uint32

const (
	RegisterOtg_hs_hcint3FieldXfrcShift = 0
	RegisterOtg_hs_hcint3FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint3Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint3FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint3Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint3FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint3FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint3FieldChhShift = 1
	RegisterOtg_hs_hcint3FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint3Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint3FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint3Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint3FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint3FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint3FieldAhberrShift = 2
	RegisterOtg_hs_hcint3FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint3Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint3FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint3Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint3FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint3FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint3FieldStallShift = 3
	RegisterOtg_hs_hcint3FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint3Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint3FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint3Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint3FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint3FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint3FieldNakShift = 4
	RegisterOtg_hs_hcint3FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint3Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint3FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint3Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint3FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint3FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint3FieldAckShift = 5
	RegisterOtg_hs_hcint3FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint3Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint3FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint3Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint3FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint3FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint3FieldNyetShift = 6
	RegisterOtg_hs_hcint3FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint3Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint3FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint3Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint3FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint3FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint3FieldTxerrShift = 7
	RegisterOtg_hs_hcint3FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint3Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint3FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint3Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint3FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint3FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint3FieldBberrShift = 8
	RegisterOtg_hs_hcint3FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint3Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint3FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint3Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint3FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint3FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint3FieldFrmorShift = 9
	RegisterOtg_hs_hcint3FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint3Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint3FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint3Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint3FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint3FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint3FieldDterrShift = 10
	RegisterOtg_hs_hcint3FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint3Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint3FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint3Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint3FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint3FieldDterrMask)
	}
}

// registerOtg_hs_hcint4Type OTG_HS host channel-4 interrupt register
type registerOtg_hs_hcint4Type uint32

const (
	RegisterOtg_hs_hcint4FieldXfrcShift = 0
	RegisterOtg_hs_hcint4FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint4Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint4FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint4Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint4FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint4FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint4FieldChhShift = 1
	RegisterOtg_hs_hcint4FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint4Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint4FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint4Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint4FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint4FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint4FieldAhberrShift = 2
	RegisterOtg_hs_hcint4FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint4Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint4FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint4Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint4FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint4FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint4FieldStallShift = 3
	RegisterOtg_hs_hcint4FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint4Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint4FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint4Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint4FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint4FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint4FieldNakShift = 4
	RegisterOtg_hs_hcint4FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint4Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint4FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint4Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint4FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint4FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint4FieldAckShift = 5
	RegisterOtg_hs_hcint4FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint4Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint4FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint4Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint4FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint4FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint4FieldNyetShift = 6
	RegisterOtg_hs_hcint4FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint4Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint4FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint4Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint4FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint4FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint4FieldTxerrShift = 7
	RegisterOtg_hs_hcint4FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint4Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint4FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint4Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint4FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint4FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint4FieldBberrShift = 8
	RegisterOtg_hs_hcint4FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint4Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint4FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint4Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint4FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint4FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint4FieldFrmorShift = 9
	RegisterOtg_hs_hcint4FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint4Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint4FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint4Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint4FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint4FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint4FieldDterrShift = 10
	RegisterOtg_hs_hcint4FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint4Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint4FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint4Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint4FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint4FieldDterrMask)
	}
}

// registerOtg_hs_hcint5Type OTG_HS host channel-5 interrupt register
type registerOtg_hs_hcint5Type uint32

const (
	RegisterOtg_hs_hcint5FieldXfrcShift = 0
	RegisterOtg_hs_hcint5FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint5Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint5FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint5Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint5FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint5FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint5FieldChhShift = 1
	RegisterOtg_hs_hcint5FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint5Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint5FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint5Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint5FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint5FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint5FieldAhberrShift = 2
	RegisterOtg_hs_hcint5FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint5Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint5FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint5Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint5FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint5FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint5FieldStallShift = 3
	RegisterOtg_hs_hcint5FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint5Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint5FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint5Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint5FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint5FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint5FieldNakShift = 4
	RegisterOtg_hs_hcint5FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint5Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint5FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint5Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint5FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint5FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint5FieldAckShift = 5
	RegisterOtg_hs_hcint5FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint5Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint5FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint5Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint5FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint5FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint5FieldNyetShift = 6
	RegisterOtg_hs_hcint5FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint5Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint5FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint5Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint5FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint5FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint5FieldTxerrShift = 7
	RegisterOtg_hs_hcint5FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint5Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint5FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint5Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint5FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint5FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint5FieldBberrShift = 8
	RegisterOtg_hs_hcint5FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint5Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint5FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint5Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint5FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint5FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint5FieldFrmorShift = 9
	RegisterOtg_hs_hcint5FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint5Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint5FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint5Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint5FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint5FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint5FieldDterrShift = 10
	RegisterOtg_hs_hcint5FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint5Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint5FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint5Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint5FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint5FieldDterrMask)
	}
}

// registerOtg_hs_hcint6Type OTG_HS host channel-6 interrupt register
type registerOtg_hs_hcint6Type uint32

const (
	RegisterOtg_hs_hcint6FieldXfrcShift = 0
	RegisterOtg_hs_hcint6FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint6Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint6FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint6Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint6FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint6FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint6FieldChhShift = 1
	RegisterOtg_hs_hcint6FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint6Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint6FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint6Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint6FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint6FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint6FieldAhberrShift = 2
	RegisterOtg_hs_hcint6FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint6Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint6FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint6Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint6FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint6FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint6FieldStallShift = 3
	RegisterOtg_hs_hcint6FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint6Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint6FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint6Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint6FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint6FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint6FieldNakShift = 4
	RegisterOtg_hs_hcint6FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint6Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint6FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint6Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint6FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint6FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint6FieldAckShift = 5
	RegisterOtg_hs_hcint6FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint6Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint6FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint6Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint6FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint6FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint6FieldNyetShift = 6
	RegisterOtg_hs_hcint6FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint6Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint6FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint6Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint6FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint6FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint6FieldTxerrShift = 7
	RegisterOtg_hs_hcint6FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint6Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint6FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint6Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint6FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint6FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint6FieldBberrShift = 8
	RegisterOtg_hs_hcint6FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint6Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint6FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint6Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint6FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint6FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint6FieldFrmorShift = 9
	RegisterOtg_hs_hcint6FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint6Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint6FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint6Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint6FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint6FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint6FieldDterrShift = 10
	RegisterOtg_hs_hcint6FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint6Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint6FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint6Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint6FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint6FieldDterrMask)
	}
}

// registerOtg_hs_hcint7Type OTG_HS host channel-7 interrupt register
type registerOtg_hs_hcint7Type uint32

const (
	RegisterOtg_hs_hcint7FieldXfrcShift = 0
	RegisterOtg_hs_hcint7FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint7Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint7FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint7Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint7FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint7FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint7FieldChhShift = 1
	RegisterOtg_hs_hcint7FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint7Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint7FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint7Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint7FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint7FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint7FieldAhberrShift = 2
	RegisterOtg_hs_hcint7FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint7Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint7FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint7Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint7FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint7FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint7FieldStallShift = 3
	RegisterOtg_hs_hcint7FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint7Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint7FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint7Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint7FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint7FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint7FieldNakShift = 4
	RegisterOtg_hs_hcint7FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint7Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint7FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint7Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint7FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint7FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint7FieldAckShift = 5
	RegisterOtg_hs_hcint7FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint7Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint7FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint7Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint7FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint7FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint7FieldNyetShift = 6
	RegisterOtg_hs_hcint7FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint7Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint7FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint7Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint7FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint7FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint7FieldTxerrShift = 7
	RegisterOtg_hs_hcint7FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint7Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint7FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint7Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint7FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint7FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint7FieldBberrShift = 8
	RegisterOtg_hs_hcint7FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint7Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint7FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint7Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint7FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint7FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint7FieldFrmorShift = 9
	RegisterOtg_hs_hcint7FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint7Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint7FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint7Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint7FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint7FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint7FieldDterrShift = 10
	RegisterOtg_hs_hcint7FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint7Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint7FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint7Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint7FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint7FieldDterrMask)
	}
}

// registerOtg_hs_hcint8Type OTG_HS host channel-8 interrupt register
type registerOtg_hs_hcint8Type uint32

const (
	RegisterOtg_hs_hcint8FieldXfrcShift = 0
	RegisterOtg_hs_hcint8FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint8Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint8FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint8Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint8FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint8FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint8FieldChhShift = 1
	RegisterOtg_hs_hcint8FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint8Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint8FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint8Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint8FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint8FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint8FieldAhberrShift = 2
	RegisterOtg_hs_hcint8FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint8Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint8FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint8Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint8FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint8FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint8FieldStallShift = 3
	RegisterOtg_hs_hcint8FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint8Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint8FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint8Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint8FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint8FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint8FieldNakShift = 4
	RegisterOtg_hs_hcint8FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint8Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint8FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint8Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint8FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint8FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint8FieldAckShift = 5
	RegisterOtg_hs_hcint8FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint8Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint8FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint8Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint8FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint8FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint8FieldNyetShift = 6
	RegisterOtg_hs_hcint8FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint8Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint8FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint8Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint8FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint8FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint8FieldTxerrShift = 7
	RegisterOtg_hs_hcint8FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint8Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint8FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint8Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint8FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint8FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint8FieldBberrShift = 8
	RegisterOtg_hs_hcint8FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint8Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint8FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint8Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint8FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint8FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint8FieldFrmorShift = 9
	RegisterOtg_hs_hcint8FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint8Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint8FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint8Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint8FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint8FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint8FieldDterrShift = 10
	RegisterOtg_hs_hcint8FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint8Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint8FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint8Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint8FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint8FieldDterrMask)
	}
}

// registerOtg_hs_hcint9Type OTG_HS host channel-9 interrupt register
type registerOtg_hs_hcint9Type uint32

const (
	RegisterOtg_hs_hcint9FieldXfrcShift = 0
	RegisterOtg_hs_hcint9FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint9Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint9FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint9Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint9FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint9FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint9FieldChhShift = 1
	RegisterOtg_hs_hcint9FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint9Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint9FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint9Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint9FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint9FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint9FieldAhberrShift = 2
	RegisterOtg_hs_hcint9FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint9Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint9FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint9Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint9FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint9FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint9FieldStallShift = 3
	RegisterOtg_hs_hcint9FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint9Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint9FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint9Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint9FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint9FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint9FieldNakShift = 4
	RegisterOtg_hs_hcint9FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint9Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint9FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint9Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint9FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint9FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint9FieldAckShift = 5
	RegisterOtg_hs_hcint9FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint9Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint9FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint9Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint9FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint9FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint9FieldNyetShift = 6
	RegisterOtg_hs_hcint9FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint9Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint9FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint9Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint9FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint9FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint9FieldTxerrShift = 7
	RegisterOtg_hs_hcint9FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint9Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint9FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint9Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint9FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint9FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint9FieldBberrShift = 8
	RegisterOtg_hs_hcint9FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint9Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint9FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint9Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint9FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint9FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint9FieldFrmorShift = 9
	RegisterOtg_hs_hcint9FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint9Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint9FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint9Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint9FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint9FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint9FieldDterrShift = 10
	RegisterOtg_hs_hcint9FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint9Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint9FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint9Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint9FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint9FieldDterrMask)
	}
}

// registerOtg_hs_hcint10Type OTG_HS host channel-10 interrupt register
type registerOtg_hs_hcint10Type uint32

const (
	RegisterOtg_hs_hcint10FieldXfrcShift = 0
	RegisterOtg_hs_hcint10FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint10Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint10FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint10Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint10FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint10FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint10FieldChhShift = 1
	RegisterOtg_hs_hcint10FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint10Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint10FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint10Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint10FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint10FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint10FieldAhberrShift = 2
	RegisterOtg_hs_hcint10FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint10Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint10FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint10Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint10FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint10FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint10FieldStallShift = 3
	RegisterOtg_hs_hcint10FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint10Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint10FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint10Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint10FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint10FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint10FieldNakShift = 4
	RegisterOtg_hs_hcint10FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint10Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint10FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint10Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint10FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint10FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint10FieldAckShift = 5
	RegisterOtg_hs_hcint10FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint10Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint10FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint10Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint10FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint10FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint10FieldNyetShift = 6
	RegisterOtg_hs_hcint10FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint10Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint10FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint10Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint10FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint10FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint10FieldTxerrShift = 7
	RegisterOtg_hs_hcint10FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint10Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint10FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint10Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint10FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint10FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint10FieldBberrShift = 8
	RegisterOtg_hs_hcint10FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint10Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint10FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint10Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint10FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint10FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint10FieldFrmorShift = 9
	RegisterOtg_hs_hcint10FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint10Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint10FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint10Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint10FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint10FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint10FieldDterrShift = 10
	RegisterOtg_hs_hcint10FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint10Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint10FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint10Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint10FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint10FieldDterrMask)
	}
}

// registerOtg_hs_hcint11Type OTG_HS host channel-11 interrupt register
type registerOtg_hs_hcint11Type uint32

const (
	RegisterOtg_hs_hcint11FieldXfrcShift = 0
	RegisterOtg_hs_hcint11FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint11Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint11FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint11Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint11FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint11FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint11FieldChhShift = 1
	RegisterOtg_hs_hcint11FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint11Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint11FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint11Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint11FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint11FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint11FieldAhberrShift = 2
	RegisterOtg_hs_hcint11FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint11Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint11FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint11Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint11FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint11FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint11FieldStallShift = 3
	RegisterOtg_hs_hcint11FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint11Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint11FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint11Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint11FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint11FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint11FieldNakShift = 4
	RegisterOtg_hs_hcint11FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint11Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint11FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint11Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint11FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint11FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint11FieldAckShift = 5
	RegisterOtg_hs_hcint11FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint11Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint11FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint11Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint11FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint11FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint11FieldNyetShift = 6
	RegisterOtg_hs_hcint11FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint11Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint11FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint11Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint11FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint11FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint11FieldTxerrShift = 7
	RegisterOtg_hs_hcint11FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint11Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint11FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint11Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint11FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint11FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint11FieldBberrShift = 8
	RegisterOtg_hs_hcint11FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint11Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint11FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint11Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint11FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint11FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint11FieldFrmorShift = 9
	RegisterOtg_hs_hcint11FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint11Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint11FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint11Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint11FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint11FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint11FieldDterrShift = 10
	RegisterOtg_hs_hcint11FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint11Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint11FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint11Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint11FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint11FieldDterrMask)
	}
}

// registerOtg_hs_hcintmsk0Type OTG_HS host channel-11 interrupt mask register
type registerOtg_hs_hcintmsk0Type uint32

const (
	RegisterOtg_hs_hcintmsk0FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk0FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk0Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk0FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk0Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk0FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk0FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk0FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk0FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk0Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk0FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk0Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk0FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk0FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk0FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk0FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk0Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk0FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk0Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk0FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk0FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk0FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk0FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk0Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk0FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk0Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk0FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk0FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk0FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk0FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk0Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk0FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk0Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk0FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk0FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk0FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk0FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk0Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk0FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk0Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk0FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk0FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk0FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk0FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk0Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk0FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk0Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk0FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk0FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk0FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk0FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk0Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk0FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk0Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk0FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk0FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk0FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk0FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk0Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk0FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk0Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk0FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk0FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk0FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk0FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk0Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk0FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk0Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk0FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk0FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk0FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk0FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk0Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk0FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk0Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk0FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk0FieldDterrmMask)
	}
}

// registerOtg_hs_hcintmsk1Type OTG_HS host channel-1 interrupt mask register
type registerOtg_hs_hcintmsk1Type uint32

const (
	RegisterOtg_hs_hcintmsk1FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk1FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk1Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk1FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk1Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk1FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk1FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk1FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk1FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk1Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk1FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk1Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk1FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk1FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk1FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk1FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk1Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk1FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk1Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk1FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk1FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk1FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk1FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk1Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk1FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk1Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk1FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk1FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk1FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk1FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk1Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk1FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk1Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk1FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk1FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk1FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk1FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk1Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk1FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk1Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk1FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk1FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk1FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk1FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk1Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk1FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk1Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk1FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk1FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk1FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk1FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk1Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk1FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk1Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk1FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk1FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk1FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk1FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk1Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk1FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk1Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk1FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk1FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk1FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk1FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk1Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk1FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk1Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk1FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk1FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk1FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk1FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk1Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk1FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk1Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk1FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk1FieldDterrmMask)
	}
}

// registerOtg_hs_hcintmsk2Type OTG_HS host channel-2 interrupt mask register
type registerOtg_hs_hcintmsk2Type uint32

const (
	RegisterOtg_hs_hcintmsk2FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk2FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk2Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk2FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk2Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk2FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk2FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk2FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk2FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk2Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk2FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk2Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk2FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk2FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk2FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk2FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk2Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk2FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk2Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk2FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk2FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk2FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk2FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk2Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk2FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk2Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk2FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk2FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk2FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk2FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk2Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk2FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk2Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk2FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk2FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk2FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk2FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk2Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk2FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk2Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk2FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk2FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk2FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk2FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk2Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk2FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk2Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk2FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk2FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk2FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk2FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk2Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk2FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk2Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk2FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk2FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk2FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk2FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk2Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk2FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk2Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk2FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk2FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk2FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk2FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk2Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk2FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk2Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk2FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk2FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk2FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk2FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk2Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk2FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk2Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk2FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk2FieldDterrmMask)
	}
}

// registerOtg_hs_hcintmsk3Type OTG_HS host channel-3 interrupt mask register
type registerOtg_hs_hcintmsk3Type uint32

const (
	RegisterOtg_hs_hcintmsk3FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk3FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk3Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk3FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk3Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk3FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk3FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk3FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk3FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk3Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk3FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk3Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk3FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk3FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk3FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk3FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk3Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk3FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk3Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk3FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk3FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk3FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk3FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk3Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk3FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk3Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk3FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk3FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk3FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk3FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk3Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk3FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk3Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk3FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk3FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk3FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk3FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk3Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk3FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk3Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk3FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk3FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk3FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk3FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk3Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk3FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk3Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk3FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk3FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk3FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk3FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk3Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk3FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk3Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk3FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk3FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk3FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk3FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk3Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk3FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk3Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk3FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk3FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk3FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk3FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk3Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk3FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk3Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk3FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk3FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk3FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk3FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk3Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk3FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk3Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk3FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk3FieldDterrmMask)
	}
}

// registerOtg_hs_hcintmsk4Type OTG_HS host channel-4 interrupt mask register
type registerOtg_hs_hcintmsk4Type uint32

const (
	RegisterOtg_hs_hcintmsk4FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk4FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk4Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk4FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk4Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk4FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk4FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk4FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk4FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk4Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk4FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk4Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk4FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk4FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk4FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk4FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk4Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk4FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk4Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk4FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk4FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk4FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk4FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk4Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk4FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk4Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk4FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk4FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk4FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk4FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk4Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk4FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk4Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk4FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk4FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk4FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk4FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk4Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk4FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk4Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk4FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk4FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk4FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk4FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk4Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk4FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk4Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk4FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk4FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk4FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk4FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk4Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk4FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk4Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk4FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk4FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk4FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk4FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk4Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk4FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk4Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk4FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk4FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk4FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk4FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk4Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk4FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk4Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk4FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk4FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk4FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk4FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk4Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk4FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk4Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk4FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk4FieldDterrmMask)
	}
}

// registerOtg_hs_hcintmsk5Type OTG_HS host channel-5 interrupt mask register
type registerOtg_hs_hcintmsk5Type uint32

const (
	RegisterOtg_hs_hcintmsk5FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk5FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk5Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk5FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk5Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk5FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk5FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk5FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk5FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk5Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk5FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk5Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk5FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk5FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk5FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk5FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk5Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk5FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk5Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk5FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk5FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk5FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk5FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk5Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk5FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk5Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk5FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk5FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk5FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk5FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk5Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk5FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk5Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk5FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk5FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk5FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk5FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk5Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk5FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk5Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk5FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk5FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk5FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk5FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk5Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk5FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk5Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk5FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk5FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk5FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk5FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk5Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk5FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk5Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk5FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk5FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk5FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk5FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk5Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk5FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk5Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk5FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk5FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk5FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk5FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk5Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk5FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk5Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk5FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk5FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk5FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk5FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk5Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk5FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk5Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk5FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk5FieldDterrmMask)
	}
}

// registerOtg_hs_hcintmsk6Type OTG_HS host channel-6 interrupt mask register
type registerOtg_hs_hcintmsk6Type uint32

const (
	RegisterOtg_hs_hcintmsk6FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk6FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk6Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk6FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk6Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk6FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk6FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk6FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk6FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk6Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk6FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk6Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk6FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk6FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk6FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk6FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk6Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk6FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk6Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk6FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk6FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk6FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk6FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk6Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk6FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk6Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk6FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk6FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk6FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk6FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk6Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk6FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk6Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk6FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk6FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk6FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk6FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk6Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk6FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk6Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk6FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk6FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk6FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk6FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk6Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk6FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk6Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk6FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk6FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk6FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk6FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk6Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk6FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk6Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk6FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk6FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk6FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk6FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk6Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk6FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk6Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk6FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk6FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk6FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk6FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk6Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk6FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk6Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk6FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk6FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk6FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk6FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk6Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk6FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk6Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk6FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk6FieldDterrmMask)
	}
}

// registerOtg_hs_hcintmsk7Type OTG_HS host channel-7 interrupt mask register
type registerOtg_hs_hcintmsk7Type uint32

const (
	RegisterOtg_hs_hcintmsk7FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk7FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk7Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk7FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk7Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk7FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk7FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk7FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk7FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk7Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk7FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk7Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk7FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk7FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk7FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk7FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk7Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk7FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk7Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk7FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk7FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk7FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk7FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk7Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk7FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk7Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk7FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk7FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk7FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk7FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk7Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk7FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk7Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk7FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk7FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk7FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk7FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk7Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk7FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk7Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk7FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk7FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk7FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk7FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk7Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk7FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk7Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk7FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk7FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk7FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk7FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk7Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk7FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk7Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk7FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk7FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk7FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk7FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk7Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk7FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk7Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk7FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk7FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk7FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk7FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk7Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk7FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk7Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk7FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk7FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk7FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk7FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk7Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk7FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk7Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk7FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk7FieldDterrmMask)
	}
}

// registerOtg_hs_hcintmsk8Type OTG_HS host channel-8 interrupt mask register
type registerOtg_hs_hcintmsk8Type uint32

const (
	RegisterOtg_hs_hcintmsk8FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk8FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk8Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk8FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk8Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk8FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk8FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk8FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk8FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk8Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk8FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk8Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk8FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk8FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk8FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk8FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk8Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk8FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk8Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk8FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk8FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk8FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk8FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk8Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk8FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk8Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk8FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk8FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk8FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk8FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk8Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk8FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk8Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk8FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk8FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk8FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk8FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk8Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk8FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk8Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk8FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk8FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk8FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk8FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk8Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk8FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk8Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk8FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk8FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk8FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk8FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk8Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk8FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk8Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk8FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk8FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk8FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk8FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk8Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk8FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk8Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk8FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk8FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk8FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk8FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk8Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk8FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk8Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk8FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk8FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk8FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk8FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk8Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk8FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk8Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk8FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk8FieldDterrmMask)
	}
}

// registerOtg_hs_hcintmsk9Type OTG_HS host channel-9 interrupt mask register
type registerOtg_hs_hcintmsk9Type uint32

const (
	RegisterOtg_hs_hcintmsk9FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk9FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk9Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk9FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk9Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk9FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk9FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk9FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk9FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk9Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk9FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk9Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk9FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk9FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk9FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk9FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk9Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk9FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk9Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk9FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk9FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk9FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk9FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk9Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk9FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk9Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk9FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk9FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk9FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk9FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk9Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk9FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk9Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk9FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk9FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk9FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk9FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk9Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk9FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk9Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk9FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk9FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk9FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk9FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk9Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk9FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk9Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk9FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk9FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk9FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk9FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk9Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk9FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk9Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk9FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk9FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk9FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk9FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk9Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk9FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk9Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk9FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk9FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk9FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk9FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk9Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk9FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk9Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk9FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk9FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk9FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk9FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk9Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk9FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk9Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk9FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk9FieldDterrmMask)
	}
}

// registerOtg_hs_hcintmsk10Type OTG_HS host channel-10 interrupt mask register
type registerOtg_hs_hcintmsk10Type uint32

const (
	RegisterOtg_hs_hcintmsk10FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk10FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk10Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk10FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk10Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk10FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk10FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk10FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk10FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk10Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk10FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk10Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk10FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk10FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk10FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk10FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk10Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk10FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk10Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk10FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk10FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk10FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk10FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk10Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk10FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk10Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk10FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk10FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk10FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk10FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk10Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk10FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk10Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk10FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk10FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk10FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk10FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk10Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk10FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk10Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk10FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk10FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk10FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk10FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk10Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk10FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk10Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk10FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk10FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk10FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk10FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk10Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk10FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk10Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk10FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk10FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk10FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk10FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk10Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk10FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk10Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk10FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk10FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk10FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk10FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk10Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk10FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk10Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk10FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk10FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk10FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk10FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk10Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk10FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk10Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk10FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk10FieldDterrmMask)
	}
}

// registerOtg_hs_hcintmsk11Type OTG_HS host channel-11 interrupt mask register
type registerOtg_hs_hcintmsk11Type uint32

const (
	RegisterOtg_hs_hcintmsk11FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk11FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk11Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk11FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk11Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk11FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk11FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk11FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk11FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk11Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk11FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk11Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk11FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk11FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk11FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk11FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk11Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk11FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk11Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk11FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk11FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk11FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk11FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk11Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk11FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk11Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk11FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk11FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk11FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk11FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk11Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk11FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk11Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk11FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk11FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk11FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk11FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk11Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk11FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk11Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk11FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk11FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk11FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk11FieldNyetMask  = 0x40
)

// GetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk11Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk11FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *registerOtg_hs_hcintmsk11Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk11FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk11FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk11FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk11FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk11Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk11FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *registerOtg_hs_hcintmsk11Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk11FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk11FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk11FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk11FieldBberrmMask  = 0x100
)

// GetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk11Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk11FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *registerOtg_hs_hcintmsk11Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk11FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk11FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk11FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk11FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk11Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk11FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk11Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk11FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk11FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk11FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk11FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk11Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk11FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk11Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk11FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk11FieldDterrmMask)
	}
}

// registerOtg_hs_hctsiz0Type OTG_HS host channel-11 transfer size register
type registerOtg_hs_hctsiz0Type uint32

const (
	RegisterOtg_hs_hctsiz0FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz0FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz0Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz0FieldXfrsizMask) >> RegisterOtg_hs_hctsiz0FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz0Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz0FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz0FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz0FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz0FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz0Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz0FieldPktcntMask) >> RegisterOtg_hs_hctsiz0FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz0Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz0FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz0FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz0FieldDpidShift = 29
	RegisterOtg_hs_hctsiz0FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz0Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz0FieldDpidMask) >> RegisterOtg_hs_hctsiz0FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz0Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz0FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz0FieldDpidShift))
}

// registerOtg_hs_hctsiz1Type OTG_HS host channel-1 transfer size register
type registerOtg_hs_hctsiz1Type uint32

const (
	RegisterOtg_hs_hctsiz1FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz1FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz1Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz1FieldXfrsizMask) >> RegisterOtg_hs_hctsiz1FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz1Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz1FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz1FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz1FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz1FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz1Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz1FieldPktcntMask) >> RegisterOtg_hs_hctsiz1FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz1Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz1FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz1FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz1FieldDpidShift = 29
	RegisterOtg_hs_hctsiz1FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz1Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz1FieldDpidMask) >> RegisterOtg_hs_hctsiz1FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz1Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz1FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz1FieldDpidShift))
}

// registerOtg_hs_hctsiz2Type OTG_HS host channel-2 transfer size register
type registerOtg_hs_hctsiz2Type uint32

const (
	RegisterOtg_hs_hctsiz2FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz2FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz2Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz2FieldXfrsizMask) >> RegisterOtg_hs_hctsiz2FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz2Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz2FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz2FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz2FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz2FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz2Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz2FieldPktcntMask) >> RegisterOtg_hs_hctsiz2FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz2Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz2FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz2FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz2FieldDpidShift = 29
	RegisterOtg_hs_hctsiz2FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz2Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz2FieldDpidMask) >> RegisterOtg_hs_hctsiz2FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz2Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz2FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz2FieldDpidShift))
}

// registerOtg_hs_hctsiz3Type OTG_HS host channel-3 transfer size register
type registerOtg_hs_hctsiz3Type uint32

const (
	RegisterOtg_hs_hctsiz3FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz3FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz3Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz3FieldXfrsizMask) >> RegisterOtg_hs_hctsiz3FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz3Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz3FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz3FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz3FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz3FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz3Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz3FieldPktcntMask) >> RegisterOtg_hs_hctsiz3FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz3Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz3FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz3FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz3FieldDpidShift = 29
	RegisterOtg_hs_hctsiz3FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz3Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz3FieldDpidMask) >> RegisterOtg_hs_hctsiz3FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz3Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz3FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz3FieldDpidShift))
}

// registerOtg_hs_hctsiz4Type OTG_HS host channel-4 transfer size register
type registerOtg_hs_hctsiz4Type uint32

const (
	RegisterOtg_hs_hctsiz4FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz4FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz4Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz4FieldXfrsizMask) >> RegisterOtg_hs_hctsiz4FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz4Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz4FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz4FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz4FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz4FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz4Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz4FieldPktcntMask) >> RegisterOtg_hs_hctsiz4FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz4Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz4FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz4FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz4FieldDpidShift = 29
	RegisterOtg_hs_hctsiz4FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz4Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz4FieldDpidMask) >> RegisterOtg_hs_hctsiz4FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz4Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz4FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz4FieldDpidShift))
}

// registerOtg_hs_hctsiz5Type OTG_HS host channel-5 transfer size register
type registerOtg_hs_hctsiz5Type uint32

const (
	RegisterOtg_hs_hctsiz5FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz5FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz5Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz5FieldXfrsizMask) >> RegisterOtg_hs_hctsiz5FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz5Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz5FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz5FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz5FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz5FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz5Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz5FieldPktcntMask) >> RegisterOtg_hs_hctsiz5FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz5Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz5FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz5FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz5FieldDpidShift = 29
	RegisterOtg_hs_hctsiz5FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz5Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz5FieldDpidMask) >> RegisterOtg_hs_hctsiz5FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz5Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz5FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz5FieldDpidShift))
}

// registerOtg_hs_hctsiz6Type OTG_HS host channel-6 transfer size register
type registerOtg_hs_hctsiz6Type uint32

const (
	RegisterOtg_hs_hctsiz6FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz6FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz6Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz6FieldXfrsizMask) >> RegisterOtg_hs_hctsiz6FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz6Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz6FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz6FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz6FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz6FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz6Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz6FieldPktcntMask) >> RegisterOtg_hs_hctsiz6FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz6Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz6FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz6FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz6FieldDpidShift = 29
	RegisterOtg_hs_hctsiz6FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz6Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz6FieldDpidMask) >> RegisterOtg_hs_hctsiz6FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz6Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz6FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz6FieldDpidShift))
}

// registerOtg_hs_hctsiz7Type OTG_HS host channel-7 transfer size register
type registerOtg_hs_hctsiz7Type uint32

const (
	RegisterOtg_hs_hctsiz7FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz7FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz7Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz7FieldXfrsizMask) >> RegisterOtg_hs_hctsiz7FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz7Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz7FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz7FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz7FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz7FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz7Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz7FieldPktcntMask) >> RegisterOtg_hs_hctsiz7FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz7Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz7FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz7FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz7FieldDpidShift = 29
	RegisterOtg_hs_hctsiz7FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz7Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz7FieldDpidMask) >> RegisterOtg_hs_hctsiz7FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz7Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz7FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz7FieldDpidShift))
}

// registerOtg_hs_hctsiz8Type OTG_HS host channel-8 transfer size register
type registerOtg_hs_hctsiz8Type uint32

const (
	RegisterOtg_hs_hctsiz8FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz8FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz8Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz8FieldXfrsizMask) >> RegisterOtg_hs_hctsiz8FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz8Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz8FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz8FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz8FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz8FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz8Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz8FieldPktcntMask) >> RegisterOtg_hs_hctsiz8FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz8Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz8FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz8FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz8FieldDpidShift = 29
	RegisterOtg_hs_hctsiz8FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz8Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz8FieldDpidMask) >> RegisterOtg_hs_hctsiz8FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz8Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz8FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz8FieldDpidShift))
}

// registerOtg_hs_hctsiz9Type OTG_HS host channel-9 transfer size register
type registerOtg_hs_hctsiz9Type uint32

const (
	RegisterOtg_hs_hctsiz9FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz9FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz9Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz9FieldXfrsizMask) >> RegisterOtg_hs_hctsiz9FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz9Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz9FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz9FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz9FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz9FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz9Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz9FieldPktcntMask) >> RegisterOtg_hs_hctsiz9FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz9Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz9FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz9FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz9FieldDpidShift = 29
	RegisterOtg_hs_hctsiz9FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz9Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz9FieldDpidMask) >> RegisterOtg_hs_hctsiz9FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz9Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz9FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz9FieldDpidShift))
}

// registerOtg_hs_hctsiz10Type OTG_HS host channel-10 transfer size register
type registerOtg_hs_hctsiz10Type uint32

const (
	RegisterOtg_hs_hctsiz10FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz10FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz10Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz10FieldXfrsizMask) >> RegisterOtg_hs_hctsiz10FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz10Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz10FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz10FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz10FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz10FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz10Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz10FieldPktcntMask) >> RegisterOtg_hs_hctsiz10FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz10Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz10FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz10FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz10FieldDpidShift = 29
	RegisterOtg_hs_hctsiz10FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz10Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz10FieldDpidMask) >> RegisterOtg_hs_hctsiz10FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz10Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz10FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz10FieldDpidShift))
}

// registerOtg_hs_hctsiz11Type OTG_HS host channel-11 transfer size register
type registerOtg_hs_hctsiz11Type uint32

const (
	RegisterOtg_hs_hctsiz11FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz11FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz11Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz11FieldXfrsizMask) >> RegisterOtg_hs_hctsiz11FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz11Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz11FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz11FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz11FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz11FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz11Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz11FieldPktcntMask) >> RegisterOtg_hs_hctsiz11FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz11Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz11FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz11FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz11FieldDpidShift = 29
	RegisterOtg_hs_hctsiz11FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz11Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz11FieldDpidMask) >> RegisterOtg_hs_hctsiz11FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz11Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz11FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz11FieldDpidShift))
}

// registerOtg_hs_hcdma0Type OTG_HS host channel-0 DMA address register
type registerOtg_hs_hcdma0Type uint32

const (
	RegisterOtg_hs_hcdma0FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma0FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma0Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma0FieldDmaaddrMask) >> RegisterOtg_hs_hcdma0FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma0Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma0FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma0FieldDmaaddrShift))
}

// registerOtg_hs_hcdma1Type OTG_HS host channel-1 DMA address register
type registerOtg_hs_hcdma1Type uint32

const (
	RegisterOtg_hs_hcdma1FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma1FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma1Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma1FieldDmaaddrMask) >> RegisterOtg_hs_hcdma1FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma1Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma1FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma1FieldDmaaddrShift))
}

// registerOtg_hs_hcdma2Type OTG_HS host channel-2 DMA address register
type registerOtg_hs_hcdma2Type uint32

const (
	RegisterOtg_hs_hcdma2FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma2FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma2Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma2FieldDmaaddrMask) >> RegisterOtg_hs_hcdma2FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma2Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma2FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma2FieldDmaaddrShift))
}

// registerOtg_hs_hcdma3Type OTG_HS host channel-3 DMA address register
type registerOtg_hs_hcdma3Type uint32

const (
	RegisterOtg_hs_hcdma3FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma3FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma3Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma3FieldDmaaddrMask) >> RegisterOtg_hs_hcdma3FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma3Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma3FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma3FieldDmaaddrShift))
}

// registerOtg_hs_hcdma4Type OTG_HS host channel-4 DMA address register
type registerOtg_hs_hcdma4Type uint32

const (
	RegisterOtg_hs_hcdma4FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma4FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma4Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma4FieldDmaaddrMask) >> RegisterOtg_hs_hcdma4FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma4Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma4FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma4FieldDmaaddrShift))
}

// registerOtg_hs_hcdma5Type OTG_HS host channel-5 DMA address register
type registerOtg_hs_hcdma5Type uint32

const (
	RegisterOtg_hs_hcdma5FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma5FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma5Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma5FieldDmaaddrMask) >> RegisterOtg_hs_hcdma5FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma5Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma5FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma5FieldDmaaddrShift))
}

// registerOtg_hs_hcdma6Type OTG_HS host channel-6 DMA address register
type registerOtg_hs_hcdma6Type uint32

const (
	RegisterOtg_hs_hcdma6FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma6FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma6Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma6FieldDmaaddrMask) >> RegisterOtg_hs_hcdma6FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma6Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma6FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma6FieldDmaaddrShift))
}

// registerOtg_hs_hcdma7Type OTG_HS host channel-7 DMA address register
type registerOtg_hs_hcdma7Type uint32

const (
	RegisterOtg_hs_hcdma7FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma7FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma7Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma7FieldDmaaddrMask) >> RegisterOtg_hs_hcdma7FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma7Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma7FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma7FieldDmaaddrShift))
}

// registerOtg_hs_hcdma8Type OTG_HS host channel-8 DMA address register
type registerOtg_hs_hcdma8Type uint32

const (
	RegisterOtg_hs_hcdma8FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma8FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma8Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma8FieldDmaaddrMask) >> RegisterOtg_hs_hcdma8FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma8Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma8FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma8FieldDmaaddrShift))
}

// registerOtg_hs_hcdma9Type OTG_HS host channel-9 DMA address register
type registerOtg_hs_hcdma9Type uint32

const (
	RegisterOtg_hs_hcdma9FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma9FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma9Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma9FieldDmaaddrMask) >> RegisterOtg_hs_hcdma9FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma9Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma9FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma9FieldDmaaddrShift))
}

// registerOtg_hs_hcdma10Type OTG_HS host channel-10 DMA address register
type registerOtg_hs_hcdma10Type uint32

const (
	RegisterOtg_hs_hcdma10FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma10FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma10Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma10FieldDmaaddrMask) >> RegisterOtg_hs_hcdma10FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma10Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma10FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma10FieldDmaaddrShift))
}

// registerOtg_hs_hcdma11Type OTG_HS host channel-11 DMA address register
type registerOtg_hs_hcdma11Type uint32

const (
	RegisterOtg_hs_hcdma11FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma11FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma11Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma11FieldDmaaddrMask) >> RegisterOtg_hs_hcdma11FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma11Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma11FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma11FieldDmaaddrShift))
}

// registerOtg_hs_hcchar12Type OTG_HS host channel-12 characteristics register
type registerOtg_hs_hcchar12Type uint32

const (
	RegisterOtg_hs_hcchar12FieldMpsizShift = 0
	RegisterOtg_hs_hcchar12FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar12Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar12FieldMpsizMask) >> RegisterOtg_hs_hcchar12FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar12Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar12FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar12FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar12FieldEpnumShift = 11
	RegisterOtg_hs_hcchar12FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar12Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar12FieldEpnumMask) >> RegisterOtg_hs_hcchar12FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar12Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar12FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar12FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar12FieldEpdirShift = 15
	RegisterOtg_hs_hcchar12FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar12Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar12FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar12Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar12FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar12FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar12FieldLsdevShift = 17
	RegisterOtg_hs_hcchar12FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar12Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar12FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar12Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar12FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar12FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar12FieldEptypShift = 18
	RegisterOtg_hs_hcchar12FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar12Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar12FieldEptypMask) >> RegisterOtg_hs_hcchar12FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar12Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar12FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar12FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar12FieldMcShift = 20
	RegisterOtg_hs_hcchar12FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar12Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar12FieldMcMask) >> RegisterOtg_hs_hcchar12FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar12Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar12FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar12FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar12FieldDadShift = 22
	RegisterOtg_hs_hcchar12FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar12Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar12FieldDadMask) >> RegisterOtg_hs_hcchar12FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar12Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar12FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar12FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar12FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar12FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar12Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar12FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar12Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar12FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar12FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar12FieldChdisShift = 30
	RegisterOtg_hs_hcchar12FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar12Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar12FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar12Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar12FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar12FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar12FieldChenaShift = 31
	RegisterOtg_hs_hcchar12FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar12Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar12FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar12Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar12FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar12FieldChenaMask)
	}
}

// registerOtg_hs_hcsplt12Type OTG_HS host channel-12 split control register
type registerOtg_hs_hcsplt12Type uint32

const (
	RegisterOtg_hs_hcsplt12FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt12FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt12Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt12FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt12FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt12Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt12FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt12FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt12FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt12FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt12Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt12FieldHubaddrMask) >> RegisterOtg_hs_hcsplt12FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt12Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt12FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt12FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt12FieldXactposShift = 14
	RegisterOtg_hs_hcsplt12FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt12Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt12FieldXactposMask) >> RegisterOtg_hs_hcsplt12FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt12Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt12FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt12FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt12FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt12FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt12Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt12FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt12Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt12FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt12FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt12FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt12FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt12Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt12FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt12Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt12FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt12FieldSplitenMask)
	}
}

// registerOtg_hs_hcint12Type OTG_HS host channel-12 interrupt register
type registerOtg_hs_hcint12Type uint32

const (
	RegisterOtg_hs_hcint12FieldXfrcShift = 0
	RegisterOtg_hs_hcint12FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint12Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint12FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint12Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint12FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint12FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint12FieldChhShift = 1
	RegisterOtg_hs_hcint12FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint12Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint12FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint12Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint12FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint12FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint12FieldAhberrShift = 2
	RegisterOtg_hs_hcint12FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint12Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint12FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint12Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint12FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint12FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint12FieldStallShift = 3
	RegisterOtg_hs_hcint12FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint12Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint12FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint12Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint12FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint12FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint12FieldNakShift = 4
	RegisterOtg_hs_hcint12FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint12Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint12FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint12Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint12FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint12FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint12FieldAckShift = 5
	RegisterOtg_hs_hcint12FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint12Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint12FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint12Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint12FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint12FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint12FieldNyetShift = 6
	RegisterOtg_hs_hcint12FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint12Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint12FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint12Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint12FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint12FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint12FieldTxerrShift = 7
	RegisterOtg_hs_hcint12FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint12Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint12FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint12Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint12FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint12FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint12FieldBberrShift = 8
	RegisterOtg_hs_hcint12FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint12Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint12FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint12Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint12FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint12FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint12FieldFrmorShift = 9
	RegisterOtg_hs_hcint12FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint12Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint12FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint12Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint12FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint12FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint12FieldDterrShift = 10
	RegisterOtg_hs_hcint12FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint12Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint12FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint12Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint12FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint12FieldDterrMask)
	}
}

// registerOtg_hs_hcintmsk12Type OTG_HS host channel-12 interrupt mask register
type registerOtg_hs_hcintmsk12Type uint32

const (
	RegisterOtg_hs_hcintmsk12FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk12FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk12Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk12FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk12Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk12FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk12FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk12FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk12FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk12Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk12FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk12Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk12FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk12FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk12FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk12FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk12Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk12FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk12Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk12FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk12FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk12FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk12FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk12Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk12FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk12Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk12FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk12FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk12FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk12FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk12Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk12FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk12Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk12FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk12FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk12FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk12FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk12Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk12FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk12Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk12FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk12FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk12FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk12FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcintmsk12Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk12FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcintmsk12Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk12FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk12FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk12FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk12FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error
func (r *registerOtg_hs_hcintmsk12Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk12FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *registerOtg_hs_hcintmsk12Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk12FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk12FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk12FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk12FieldBberrmMask  = 0x100
)

// GetBberrm Babble error
func (r *registerOtg_hs_hcintmsk12Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk12FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *registerOtg_hs_hcintmsk12Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk12FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk12FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk12FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk12FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk12Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk12FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk12Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk12FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk12FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk12FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk12FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk12Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk12FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk12Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk12FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk12FieldDterrmMask)
	}
}

// registerOtg_hs_hctsiz12Type OTG_HS host channel-12 transfer size register
type registerOtg_hs_hctsiz12Type uint32

const (
	RegisterOtg_hs_hctsiz12FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz12FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz12Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz12FieldXfrsizMask) >> RegisterOtg_hs_hctsiz12FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz12Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz12FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz12FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz12FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz12FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz12Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz12FieldPktcntMask) >> RegisterOtg_hs_hctsiz12FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz12Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz12FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz12FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz12FieldDpidShift = 29
	RegisterOtg_hs_hctsiz12FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz12Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz12FieldDpidMask) >> RegisterOtg_hs_hctsiz12FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz12Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz12FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz12FieldDpidShift))
}

// registerOtg_hs_hcdma12Type OTG_HS host channel-12 DMA address register
type registerOtg_hs_hcdma12Type uint32

const (
	RegisterOtg_hs_hcdma12FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma12FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma12Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma12FieldDmaaddrMask) >> RegisterOtg_hs_hcdma12FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma12Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma12FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma12FieldDmaaddrShift))
}

// registerOtg_hs_hcchar13Type OTG_HS host channel-13 characteristics register
type registerOtg_hs_hcchar13Type uint32

const (
	RegisterOtg_hs_hcchar13FieldMpsizShift = 0
	RegisterOtg_hs_hcchar13FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar13Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar13FieldMpsizMask) >> RegisterOtg_hs_hcchar13FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar13Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar13FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar13FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar13FieldEpnumShift = 11
	RegisterOtg_hs_hcchar13FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar13Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar13FieldEpnumMask) >> RegisterOtg_hs_hcchar13FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar13Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar13FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar13FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar13FieldEpdirShift = 15
	RegisterOtg_hs_hcchar13FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar13Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar13FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar13Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar13FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar13FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar13FieldLsdevShift = 17
	RegisterOtg_hs_hcchar13FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar13Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar13FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar13Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar13FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar13FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar13FieldEptypShift = 18
	RegisterOtg_hs_hcchar13FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar13Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar13FieldEptypMask) >> RegisterOtg_hs_hcchar13FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar13Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar13FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar13FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar13FieldMcShift = 20
	RegisterOtg_hs_hcchar13FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar13Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar13FieldMcMask) >> RegisterOtg_hs_hcchar13FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar13Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar13FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar13FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar13FieldDadShift = 22
	RegisterOtg_hs_hcchar13FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar13Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar13FieldDadMask) >> RegisterOtg_hs_hcchar13FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar13Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar13FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar13FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar13FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar13FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar13Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar13FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar13Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar13FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar13FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar13FieldChdisShift = 30
	RegisterOtg_hs_hcchar13FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar13Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar13FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar13Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar13FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar13FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar13FieldChenaShift = 31
	RegisterOtg_hs_hcchar13FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar13Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar13FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar13Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar13FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar13FieldChenaMask)
	}
}

// registerOtg_hs_hcsplt13Type OTG_HS host channel-13 split control register
type registerOtg_hs_hcsplt13Type uint32

const (
	RegisterOtg_hs_hcsplt13FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt13FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt13Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt13FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt13FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt13Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt13FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt13FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt13FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt13FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt13Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt13FieldHubaddrMask) >> RegisterOtg_hs_hcsplt13FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt13Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt13FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt13FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt13FieldXactposShift = 14
	RegisterOtg_hs_hcsplt13FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt13Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt13FieldXactposMask) >> RegisterOtg_hs_hcsplt13FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt13Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt13FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt13FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt13FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt13FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt13Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt13FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt13Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt13FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt13FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt13FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt13FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt13Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt13FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt13Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt13FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt13FieldSplitenMask)
	}
}

// registerOtg_hs_hcint13Type OTG_HS host channel-13 interrupt register
type registerOtg_hs_hcint13Type uint32

const (
	RegisterOtg_hs_hcint13FieldXfrcShift = 0
	RegisterOtg_hs_hcint13FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint13Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint13FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint13Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint13FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint13FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint13FieldChhShift = 1
	RegisterOtg_hs_hcint13FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint13Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint13FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint13Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint13FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint13FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint13FieldAhberrShift = 2
	RegisterOtg_hs_hcint13FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint13Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint13FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint13Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint13FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint13FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint13FieldStallShift = 3
	RegisterOtg_hs_hcint13FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint13Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint13FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint13Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint13FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint13FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint13FieldNakShift = 4
	RegisterOtg_hs_hcint13FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint13Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint13FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint13Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint13FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint13FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint13FieldAckShift = 5
	RegisterOtg_hs_hcint13FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint13Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint13FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint13Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint13FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint13FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint13FieldNyetShift = 6
	RegisterOtg_hs_hcint13FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint13Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint13FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint13Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint13FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint13FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint13FieldTxerrShift = 7
	RegisterOtg_hs_hcint13FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint13Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint13FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint13Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint13FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint13FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint13FieldBberrShift = 8
	RegisterOtg_hs_hcint13FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint13Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint13FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint13Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint13FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint13FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint13FieldFrmorShift = 9
	RegisterOtg_hs_hcint13FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint13Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint13FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint13Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint13FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint13FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint13FieldDterrShift = 10
	RegisterOtg_hs_hcint13FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint13Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint13FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint13Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint13FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint13FieldDterrMask)
	}
}

// registerOtg_hs_hcintmsk13Type OTG_HS host channel-13 interrupt mask register
type registerOtg_hs_hcintmsk13Type uint32

const (
	RegisterOtg_hs_hcintmsk13FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk13FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk13Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk13FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk13Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk13FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk13FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk13FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk13FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk13Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk13FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk13Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk13FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk13FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk13FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk13FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk13Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk13FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk13Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk13FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk13FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk13FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk13FieldStallmMask  = 0x8
)

// GetStallm STALLM response received interrupt mask
func (r *registerOtg_hs_hcintmsk13Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk13FieldStallmMask) != 0
}

// SetStallm STALLM response received interrupt mask
func (r *registerOtg_hs_hcintmsk13Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk13FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk13FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk13FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk13FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk13Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk13FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk13Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk13FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk13FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk13FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk13FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk13Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk13FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk13Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk13FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk13FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk13FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk13FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcintmsk13Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk13FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcintmsk13Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk13FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk13FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk13FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk13FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error
func (r *registerOtg_hs_hcintmsk13Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk13FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *registerOtg_hs_hcintmsk13Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk13FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk13FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk13FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk13FieldBberrmMask  = 0x100
)

// GetBberrm Babble error
func (r *registerOtg_hs_hcintmsk13Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk13FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *registerOtg_hs_hcintmsk13Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk13FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk13FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk13FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk13FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk13Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk13FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk13Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk13FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk13FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk13FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk13FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk13Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk13FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk13Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk13FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk13FieldDterrmMask)
	}
}

// registerOtg_hs_hctsiz13Type OTG_HS host channel-13 transfer size register
type registerOtg_hs_hctsiz13Type uint32

const (
	RegisterOtg_hs_hctsiz13FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz13FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz13Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz13FieldXfrsizMask) >> RegisterOtg_hs_hctsiz13FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz13Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz13FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz13FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz13FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz13FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz13Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz13FieldPktcntMask) >> RegisterOtg_hs_hctsiz13FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz13Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz13FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz13FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz13FieldDpidShift = 29
	RegisterOtg_hs_hctsiz13FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz13Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz13FieldDpidMask) >> RegisterOtg_hs_hctsiz13FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz13Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz13FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz13FieldDpidShift))
}

// registerOtg_hs_hcdma13Type OTG_HS host channel-13 DMA address register
type registerOtg_hs_hcdma13Type uint32

const (
	RegisterOtg_hs_hcdma13FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma13FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma13Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma13FieldDmaaddrMask) >> RegisterOtg_hs_hcdma13FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma13Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma13FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma13FieldDmaaddrShift))
}

// registerOtg_hs_hcchar14Type OTG_HS host channel-14 characteristics register
type registerOtg_hs_hcchar14Type uint32

const (
	RegisterOtg_hs_hcchar14FieldMpsizShift = 0
	RegisterOtg_hs_hcchar14FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar14Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar14FieldMpsizMask) >> RegisterOtg_hs_hcchar14FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar14Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar14FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar14FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar14FieldEpnumShift = 11
	RegisterOtg_hs_hcchar14FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar14Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar14FieldEpnumMask) >> RegisterOtg_hs_hcchar14FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar14Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar14FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar14FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar14FieldEpdirShift = 15
	RegisterOtg_hs_hcchar14FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar14Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar14FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar14Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar14FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar14FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar14FieldLsdevShift = 17
	RegisterOtg_hs_hcchar14FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar14Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar14FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar14Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar14FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar14FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar14FieldEptypShift = 18
	RegisterOtg_hs_hcchar14FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar14Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar14FieldEptypMask) >> RegisterOtg_hs_hcchar14FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar14Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar14FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar14FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar14FieldMcShift = 20
	RegisterOtg_hs_hcchar14FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar14Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar14FieldMcMask) >> RegisterOtg_hs_hcchar14FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar14Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar14FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar14FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar14FieldDadShift = 22
	RegisterOtg_hs_hcchar14FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar14Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar14FieldDadMask) >> RegisterOtg_hs_hcchar14FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar14Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar14FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar14FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar14FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar14FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar14Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar14FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar14Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar14FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar14FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar14FieldChdisShift = 30
	RegisterOtg_hs_hcchar14FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar14Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar14FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar14Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar14FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar14FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar14FieldChenaShift = 31
	RegisterOtg_hs_hcchar14FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar14Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar14FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar14Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar14FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar14FieldChenaMask)
	}
}

// registerOtg_hs_hcsplt14Type OTG_HS host channel-14 split control register
type registerOtg_hs_hcsplt14Type uint32

const (
	RegisterOtg_hs_hcsplt14FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt14FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt14Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt14FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt14FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt14Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt14FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt14FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt14FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt14FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt14Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt14FieldHubaddrMask) >> RegisterOtg_hs_hcsplt14FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt14Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt14FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt14FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt14FieldXactposShift = 14
	RegisterOtg_hs_hcsplt14FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt14Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt14FieldXactposMask) >> RegisterOtg_hs_hcsplt14FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt14Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt14FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt14FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt14FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt14FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt14Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt14FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt14Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt14FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt14FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt14FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt14FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt14Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt14FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt14Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt14FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt14FieldSplitenMask)
	}
}

// registerOtg_hs_hcint14Type OTG_HS host channel-14 interrupt register
type registerOtg_hs_hcint14Type uint32

const (
	RegisterOtg_hs_hcint14FieldXfrcShift = 0
	RegisterOtg_hs_hcint14FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint14Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint14FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint14Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint14FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint14FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint14FieldChhShift = 1
	RegisterOtg_hs_hcint14FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint14Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint14FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint14Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint14FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint14FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint14FieldAhberrShift = 2
	RegisterOtg_hs_hcint14FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint14Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint14FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint14Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint14FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint14FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint14FieldStallShift = 3
	RegisterOtg_hs_hcint14FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint14Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint14FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint14Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint14FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint14FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint14FieldNakShift = 4
	RegisterOtg_hs_hcint14FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint14Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint14FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint14Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint14FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint14FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint14FieldAckShift = 5
	RegisterOtg_hs_hcint14FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint14Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint14FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint14Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint14FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint14FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint14FieldNyetShift = 6
	RegisterOtg_hs_hcint14FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint14Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint14FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint14Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint14FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint14FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint14FieldTxerrShift = 7
	RegisterOtg_hs_hcint14FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint14Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint14FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint14Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint14FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint14FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint14FieldBberrShift = 8
	RegisterOtg_hs_hcint14FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint14Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint14FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint14Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint14FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint14FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint14FieldFrmorShift = 9
	RegisterOtg_hs_hcint14FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint14Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint14FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint14Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint14FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint14FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint14FieldDterrShift = 10
	RegisterOtg_hs_hcint14FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint14Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint14FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint14Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint14FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint14FieldDterrMask)
	}
}

// registerOtg_hs_hcintmsk14Type OTG_HS host channel-14 interrupt mask register
type registerOtg_hs_hcintmsk14Type uint32

const (
	RegisterOtg_hs_hcintmsk14FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk14FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk14Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk14FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk14Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk14FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk14FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk14FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk14FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk14Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk14FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk14Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk14FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk14FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk14FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk14FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk14Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk14FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk14Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk14FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk14FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk14FieldStallmShift = 3
	RegisterOtg_hs_hcintmsk14FieldStallmMask  = 0x8
)

// GetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk14Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk14FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk14Type) SetStallm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk14FieldStallmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk14FieldStallmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk14FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk14FieldNakmMask  = 0x10
)

// GetNakm NAKM response received interrupt mask
func (r *registerOtg_hs_hcintmsk14Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk14FieldNakmMask) != 0
}

// SetNakm NAKM response received interrupt mask
func (r *registerOtg_hs_hcintmsk14Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk14FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk14FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk14FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk14FieldAckmMask  = 0x20
)

// GetAckm ACKM response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk14Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk14FieldAckmMask) != 0
}

// SetAckm ACKM response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk14Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk14FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk14FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk14FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk14FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcintmsk14Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk14FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcintmsk14Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk14FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk14FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk14FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk14FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error
func (r *registerOtg_hs_hcintmsk14Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk14FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *registerOtg_hs_hcintmsk14Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk14FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk14FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk14FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk14FieldBberrmMask  = 0x100
)

// GetBberrm Babble error
func (r *registerOtg_hs_hcintmsk14Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk14FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *registerOtg_hs_hcintmsk14Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk14FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk14FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk14FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk14FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk14Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk14FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk14Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk14FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk14FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk14FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk14FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk14Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk14FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk14Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk14FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk14FieldDterrmMask)
	}
}

// registerOtg_hs_hctsiz14Type OTG_HS host channel-14 transfer size register
type registerOtg_hs_hctsiz14Type uint32

const (
	RegisterOtg_hs_hctsiz14FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz14FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz14Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz14FieldXfrsizMask) >> RegisterOtg_hs_hctsiz14FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz14Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz14FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz14FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz14FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz14FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz14Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz14FieldPktcntMask) >> RegisterOtg_hs_hctsiz14FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz14Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz14FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz14FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz14FieldDpidShift = 29
	RegisterOtg_hs_hctsiz14FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz14Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz14FieldDpidMask) >> RegisterOtg_hs_hctsiz14FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz14Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz14FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz14FieldDpidShift))
}

// registerOtg_hs_hcdma14Type OTG_HS host channel-14 DMA address register
type registerOtg_hs_hcdma14Type uint32

const (
	RegisterOtg_hs_hcdma14FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma14FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma14Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma14FieldDmaaddrMask) >> RegisterOtg_hs_hcdma14FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma14Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma14FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma14FieldDmaaddrShift))
}

// registerOtg_hs_hcchar15Type OTG_HS host channel-15 characteristics register
type registerOtg_hs_hcchar15Type uint32

const (
	RegisterOtg_hs_hcchar15FieldMpsizShift = 0
	RegisterOtg_hs_hcchar15FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar15Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar15FieldMpsizMask) >> RegisterOtg_hs_hcchar15FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *registerOtg_hs_hcchar15Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar15FieldMpsizMask)|(uint32(value)<<RegisterOtg_hs_hcchar15FieldMpsizShift))
}

const (
	RegisterOtg_hs_hcchar15FieldEpnumShift = 11
	RegisterOtg_hs_hcchar15FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *registerOtg_hs_hcchar15Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar15FieldEpnumMask) >> RegisterOtg_hs_hcchar15FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *registerOtg_hs_hcchar15Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar15FieldEpnumMask)|(uint32(value)<<RegisterOtg_hs_hcchar15FieldEpnumShift))
}

const (
	RegisterOtg_hs_hcchar15FieldEpdirShift = 15
	RegisterOtg_hs_hcchar15FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar15Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar15FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *registerOtg_hs_hcchar15Type) SetEpdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar15FieldEpdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar15FieldEpdirMask)
	}
}

const (
	RegisterOtg_hs_hcchar15FieldLsdevShift = 17
	RegisterOtg_hs_hcchar15FieldLsdevMask  = 0x20000
)

// GetLsdev Low-speed device
func (r *registerOtg_hs_hcchar15Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar15FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *registerOtg_hs_hcchar15Type) SetLsdev(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar15FieldLsdevMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar15FieldLsdevMask)
	}
}

const (
	RegisterOtg_hs_hcchar15FieldEptypShift = 18
	RegisterOtg_hs_hcchar15FieldEptypMask  = 0xc0000
)

// GetEptyp Endpoint type
func (r *registerOtg_hs_hcchar15Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar15FieldEptypMask) >> RegisterOtg_hs_hcchar15FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *registerOtg_hs_hcchar15Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar15FieldEptypMask)|(uint32(value)<<RegisterOtg_hs_hcchar15FieldEptypShift))
}

const (
	RegisterOtg_hs_hcchar15FieldMcShift = 20
	RegisterOtg_hs_hcchar15FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar15Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar15FieldMcMask) >> RegisterOtg_hs_hcchar15FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *registerOtg_hs_hcchar15Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar15FieldMcMask)|(uint32(value)<<RegisterOtg_hs_hcchar15FieldMcShift))
}

const (
	RegisterOtg_hs_hcchar15FieldDadShift = 22
	RegisterOtg_hs_hcchar15FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *registerOtg_hs_hcchar15Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar15FieldDadMask) >> RegisterOtg_hs_hcchar15FieldDadShift)
}

// SetDad Device address
func (r *registerOtg_hs_hcchar15Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar15FieldDadMask)|(uint32(value)<<RegisterOtg_hs_hcchar15FieldDadShift))
}

const (
	RegisterOtg_hs_hcchar15FieldOddfrmShift = 29
	RegisterOtg_hs_hcchar15FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *registerOtg_hs_hcchar15Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar15FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *registerOtg_hs_hcchar15Type) SetOddfrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar15FieldOddfrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar15FieldOddfrmMask)
	}
}

const (
	RegisterOtg_hs_hcchar15FieldChdisShift = 30
	RegisterOtg_hs_hcchar15FieldChdisMask  = 0x40000000
)

// GetChdis Channel disable
func (r *registerOtg_hs_hcchar15Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar15FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *registerOtg_hs_hcchar15Type) SetChdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar15FieldChdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar15FieldChdisMask)
	}
}

const (
	RegisterOtg_hs_hcchar15FieldChenaShift = 31
	RegisterOtg_hs_hcchar15FieldChenaMask  = 0x80000000
)

// GetChena Channel enable
func (r *registerOtg_hs_hcchar15Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcchar15FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *registerOtg_hs_hcchar15Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcchar15FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcchar15FieldChenaMask)
	}
}

// registerOtg_hs_hcsplt15Type OTG_HS host channel-15 split control register
type registerOtg_hs_hcsplt15Type uint32

const (
	RegisterOtg_hs_hcsplt15FieldPrtaddrShift = 0
	RegisterOtg_hs_hcsplt15FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *registerOtg_hs_hcsplt15Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt15FieldPrtaddrMask) >> RegisterOtg_hs_hcsplt15FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *registerOtg_hs_hcsplt15Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt15FieldPrtaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt15FieldPrtaddrShift))
}

const (
	RegisterOtg_hs_hcsplt15FieldHubaddrShift = 7
	RegisterOtg_hs_hcsplt15FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *registerOtg_hs_hcsplt15Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt15FieldHubaddrMask) >> RegisterOtg_hs_hcsplt15FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *registerOtg_hs_hcsplt15Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt15FieldHubaddrMask)|(uint32(value)<<RegisterOtg_hs_hcsplt15FieldHubaddrShift))
}

const (
	RegisterOtg_hs_hcsplt15FieldXactposShift = 14
	RegisterOtg_hs_hcsplt15FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt15Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt15FieldXactposMask) >> RegisterOtg_hs_hcsplt15FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *registerOtg_hs_hcsplt15Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt15FieldXactposMask)|(uint32(value)<<RegisterOtg_hs_hcsplt15FieldXactposShift))
}

const (
	RegisterOtg_hs_hcsplt15FieldComplspltShift = 16
	RegisterOtg_hs_hcsplt15FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt15Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt15FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *registerOtg_hs_hcsplt15Type) SetComplsplt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt15FieldComplspltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt15FieldComplspltMask)
	}
}

const (
	RegisterOtg_hs_hcsplt15FieldSplitenShift = 31
	RegisterOtg_hs_hcsplt15FieldSplitenMask  = 0x80000000
)

// GetSpliten Split enable
func (r *registerOtg_hs_hcsplt15Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcsplt15FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *registerOtg_hs_hcsplt15Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcsplt15FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcsplt15FieldSplitenMask)
	}
}

// registerOtg_hs_hcint15Type OTG_HS host channel-15 interrupt register
type registerOtg_hs_hcint15Type uint32

const (
	RegisterOtg_hs_hcint15FieldXfrcShift = 0
	RegisterOtg_hs_hcint15FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *registerOtg_hs_hcint15Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint15FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *registerOtg_hs_hcint15Type) SetXfrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint15FieldXfrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint15FieldXfrcMask)
	}
}

const (
	RegisterOtg_hs_hcint15FieldChhShift = 1
	RegisterOtg_hs_hcint15FieldChhMask  = 0x2
)

// GetChh Channel halted
func (r *registerOtg_hs_hcint15Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint15FieldChhMask) != 0
}

// SetChh Channel halted
func (r *registerOtg_hs_hcint15Type) SetChh(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint15FieldChhMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint15FieldChhMask)
	}
}

const (
	RegisterOtg_hs_hcint15FieldAhberrShift = 2
	RegisterOtg_hs_hcint15FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcint15Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint15FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcint15Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint15FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint15FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcint15FieldStallShift = 3
	RegisterOtg_hs_hcint15FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt
func (r *registerOtg_hs_hcint15Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint15FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *registerOtg_hs_hcint15Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint15FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint15FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcint15FieldNakShift = 4
	RegisterOtg_hs_hcint15FieldNakMask  = 0x10
)

// GetNak NAK response received interrupt
func (r *registerOtg_hs_hcint15Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint15FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *registerOtg_hs_hcint15Type) SetNak(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint15FieldNakMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint15FieldNakMask)
	}
}

const (
	RegisterOtg_hs_hcint15FieldAckShift = 5
	RegisterOtg_hs_hcint15FieldAckMask  = 0x20
)

// GetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint15Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint15FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *registerOtg_hs_hcint15Type) SetAck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint15FieldAckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint15FieldAckMask)
	}
}

const (
	RegisterOtg_hs_hcint15FieldNyetShift = 6
	RegisterOtg_hs_hcint15FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcint15Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint15FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcint15Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint15FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint15FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcint15FieldTxerrShift = 7
	RegisterOtg_hs_hcint15FieldTxerrMask  = 0x80
)

// GetTxerr Transaction error
func (r *registerOtg_hs_hcint15Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint15FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *registerOtg_hs_hcint15Type) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint15FieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint15FieldTxerrMask)
	}
}

const (
	RegisterOtg_hs_hcint15FieldBberrShift = 8
	RegisterOtg_hs_hcint15FieldBberrMask  = 0x100
)

// GetBberr Babble error
func (r *registerOtg_hs_hcint15Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint15FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *registerOtg_hs_hcint15Type) SetBberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint15FieldBberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint15FieldBberrMask)
	}
}

const (
	RegisterOtg_hs_hcint15FieldFrmorShift = 9
	RegisterOtg_hs_hcint15FieldFrmorMask  = 0x200
)

// GetFrmor Frame overrun
func (r *registerOtg_hs_hcint15Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint15FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *registerOtg_hs_hcint15Type) SetFrmor(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint15FieldFrmorMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint15FieldFrmorMask)
	}
}

const (
	RegisterOtg_hs_hcint15FieldDterrShift = 10
	RegisterOtg_hs_hcint15FieldDterrMask  = 0x400
)

// GetDterr Data toggle error
func (r *registerOtg_hs_hcint15Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcint15FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *registerOtg_hs_hcint15Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcint15FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcint15FieldDterrMask)
	}
}

// registerOtg_hs_hcintmsk15Type OTG_HS host channel-15 interrupt mask register
type registerOtg_hs_hcintmsk15Type uint32

const (
	RegisterOtg_hs_hcintmsk15FieldXfrcmShift = 0
	RegisterOtg_hs_hcintmsk15FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk15Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk15FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *registerOtg_hs_hcintmsk15Type) SetXfrcm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk15FieldXfrcmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk15FieldXfrcmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk15FieldChhmShift = 1
	RegisterOtg_hs_hcintmsk15FieldChhmMask  = 0x2
)

// GetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk15Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk15FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *registerOtg_hs_hcintmsk15Type) SetChhm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk15FieldChhmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk15FieldChhmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk15FieldAhberrShift = 2
	RegisterOtg_hs_hcintmsk15FieldAhberrMask  = 0x4
)

// GetAhberr AHB error
func (r *registerOtg_hs_hcintmsk15Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk15FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *registerOtg_hs_hcintmsk15Type) SetAhberr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk15FieldAhberrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk15FieldAhberrMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk15FieldStallShift = 3
	RegisterOtg_hs_hcintmsk15FieldStallMask  = 0x8
)

// GetStall STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk15Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk15FieldStallMask) != 0
}

// SetStall STALL response received interrupt mask
func (r *registerOtg_hs_hcintmsk15Type) SetStall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk15FieldStallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk15FieldStallMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk15FieldNakmShift = 4
	RegisterOtg_hs_hcintmsk15FieldNakmMask  = 0x10
)

// GetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk15Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk15FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *registerOtg_hs_hcintmsk15Type) SetNakm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk15FieldNakmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk15FieldNakmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk15FieldAckmShift = 5
	RegisterOtg_hs_hcintmsk15FieldAckmMask  = 0x20
)

// GetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk15Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk15FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *registerOtg_hs_hcintmsk15Type) SetAckm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk15FieldAckmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk15FieldAckmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk15FieldNyetShift = 6
	RegisterOtg_hs_hcintmsk15FieldNyetMask  = 0x40
)

// GetNyet Response received interrupt
func (r *registerOtg_hs_hcintmsk15Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk15FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *registerOtg_hs_hcintmsk15Type) SetNyet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk15FieldNyetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk15FieldNyetMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk15FieldTxerrmShift = 7
	RegisterOtg_hs_hcintmsk15FieldTxerrmMask  = 0x80
)

// GetTxerrm Transaction error
func (r *registerOtg_hs_hcintmsk15Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk15FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *registerOtg_hs_hcintmsk15Type) SetTxerrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk15FieldTxerrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk15FieldTxerrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk15FieldBberrmShift = 8
	RegisterOtg_hs_hcintmsk15FieldBberrmMask  = 0x100
)

// GetBberrm Babble error
func (r *registerOtg_hs_hcintmsk15Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk15FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *registerOtg_hs_hcintmsk15Type) SetBberrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk15FieldBberrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk15FieldBberrmMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk15FieldFrmormShift = 9
	RegisterOtg_hs_hcintmsk15FieldFrmormMask  = 0x200
)

// GetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk15Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk15FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *registerOtg_hs_hcintmsk15Type) SetFrmorm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk15FieldFrmormMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk15FieldFrmormMask)
	}
}

const (
	RegisterOtg_hs_hcintmsk15FieldDterrmShift = 10
	RegisterOtg_hs_hcintmsk15FieldDterrmMask  = 0x400
)

// GetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk15Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcintmsk15FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *registerOtg_hs_hcintmsk15Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtg_hs_hcintmsk15FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcintmsk15FieldDterrmMask)
	}
}

// registerOtg_hs_hctsiz15Type OTG_HS host channel-15 transfer size register
type registerOtg_hs_hctsiz15Type uint32

const (
	RegisterOtg_hs_hctsiz15FieldXfrsizShift = 0
	RegisterOtg_hs_hctsiz15FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz15Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz15FieldXfrsizMask) >> RegisterOtg_hs_hctsiz15FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *registerOtg_hs_hctsiz15Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz15FieldXfrsizMask)|(uint32(value)<<RegisterOtg_hs_hctsiz15FieldXfrsizShift))
}

const (
	RegisterOtg_hs_hctsiz15FieldPktcntShift = 19
	RegisterOtg_hs_hctsiz15FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *registerOtg_hs_hctsiz15Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz15FieldPktcntMask) >> RegisterOtg_hs_hctsiz15FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *registerOtg_hs_hctsiz15Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz15FieldPktcntMask)|(uint32(value)<<RegisterOtg_hs_hctsiz15FieldPktcntShift))
}

const (
	RegisterOtg_hs_hctsiz15FieldDpidShift = 29
	RegisterOtg_hs_hctsiz15FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *registerOtg_hs_hctsiz15Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hctsiz15FieldDpidMask) >> RegisterOtg_hs_hctsiz15FieldDpidShift)
}

// SetDpid Data PID
func (r *registerOtg_hs_hctsiz15Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hctsiz15FieldDpidMask)|(uint32(value)<<RegisterOtg_hs_hctsiz15FieldDpidShift))
}

// registerOtg_hs_hcdma15Type OTG_HS host channel-15 DMA address register
type registerOtg_hs_hcdma15Type uint32

const (
	RegisterOtg_hs_hcdma15FieldDmaaddrShift = 0
	RegisterOtg_hs_hcdma15FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *registerOtg_hs_hcdma15Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtg_hs_hcdma15FieldDmaaddrMask) >> RegisterOtg_hs_hcdma15FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *registerOtg_hs_hcdma15Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtg_hs_hcdma15FieldDmaaddrMask)|(uint32(value)<<RegisterOtg_hs_hcdma15FieldDmaaddrShift))
}
