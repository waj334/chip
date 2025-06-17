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
	Otghshcfg       RegisterOtghshcfgType
	Otghshfir       RegisterOtghshfirType
	Otghshfnum      RegisterOtghshfnumType
	_               [4]byte
	Otghshptxsts    RegisterOtghshptxstsType
	Otghshaint      RegisterOtghshaintType
	Otghshaintmsk   RegisterOtghshaintmskType
	_               [36]byte
	Otghshprt       RegisterOtghshprtType
	_               [188]byte
	Otghshcchar0    RegisterOtghshcchar0Type
	Otghshcsplt0    RegisterOtghshcsplt0Type
	Otghshcint0     RegisterOtghshcint0Type
	Otghshcintmsk0  RegisterOtghshcintmsk0Type
	Otghshctsiz0    RegisterOtghshctsiz0Type
	Otghshcdma0     RegisterOtghshcdma0Type
	_               [8]byte
	Otghshcchar1    RegisterOtghshcchar1Type
	Otghshcsplt1    RegisterOtghshcsplt1Type
	Otghshcint1     RegisterOtghshcint1Type
	Otghshcintmsk1  RegisterOtghshcintmsk1Type
	Otghshctsiz1    RegisterOtghshctsiz1Type
	Otghshcdma1     RegisterOtghshcdma1Type
	_               [8]byte
	Otghshcchar2    RegisterOtghshcchar2Type
	Otghshcsplt2    RegisterOtghshcsplt2Type
	Otghshcint2     RegisterOtghshcint2Type
	Otghshcintmsk2  RegisterOtghshcintmsk2Type
	Otghshctsiz2    RegisterOtghshctsiz2Type
	Otghshcdma2     RegisterOtghshcdma2Type
	_               [8]byte
	Otghshcchar3    RegisterOtghshcchar3Type
	Otghshcsplt3    RegisterOtghshcsplt3Type
	Otghshcint3     RegisterOtghshcint3Type
	Otghshcintmsk3  RegisterOtghshcintmsk3Type
	Otghshctsiz3    RegisterOtghshctsiz3Type
	Otghshcdma3     RegisterOtghshcdma3Type
	_               [8]byte
	Otghshcchar4    RegisterOtghshcchar4Type
	Otghshcsplt4    RegisterOtghshcsplt4Type
	Otghshcint4     RegisterOtghshcint4Type
	Otghshcintmsk4  RegisterOtghshcintmsk4Type
	Otghshctsiz4    RegisterOtghshctsiz4Type
	Otghshcdma4     RegisterOtghshcdma4Type
	_               [8]byte
	Otghshcchar5    RegisterOtghshcchar5Type
	Otghshcsplt5    RegisterOtghshcsplt5Type
	Otghshcint5     RegisterOtghshcint5Type
	Otghshcintmsk5  RegisterOtghshcintmsk5Type
	Otghshctsiz5    RegisterOtghshctsiz5Type
	Otghshcdma5     RegisterOtghshcdma5Type
	_               [8]byte
	Otghshcchar6    RegisterOtghshcchar6Type
	Otghshcsplt6    RegisterOtghshcsplt6Type
	Otghshcint6     RegisterOtghshcint6Type
	Otghshcintmsk6  RegisterOtghshcintmsk6Type
	Otghshctsiz6    RegisterOtghshctsiz6Type
	Otghshcdma6     RegisterOtghshcdma6Type
	_               [8]byte
	Otghshcchar7    RegisterOtghshcchar7Type
	Otghshcsplt7    RegisterOtghshcsplt7Type
	Otghshcint7     RegisterOtghshcint7Type
	Otghshcintmsk7  RegisterOtghshcintmsk7Type
	Otghshctsiz7    RegisterOtghshctsiz7Type
	Otghshcdma7     RegisterOtghshcdma7Type
	_               [8]byte
	Otghshcchar8    RegisterOtghshcchar8Type
	Otghshcsplt8    RegisterOtghshcsplt8Type
	Otghshcint8     RegisterOtghshcint8Type
	Otghshcintmsk8  RegisterOtghshcintmsk8Type
	Otghshctsiz8    RegisterOtghshctsiz8Type
	Otghshcdma8     RegisterOtghshcdma8Type
	_               [8]byte
	Otghshcchar9    RegisterOtghshcchar9Type
	Otghshcsplt9    RegisterOtghshcsplt9Type
	Otghshcint9     RegisterOtghshcint9Type
	Otghshcintmsk9  RegisterOtghshcintmsk9Type
	Otghshctsiz9    RegisterOtghshctsiz9Type
	Otghshcdma9     RegisterOtghshcdma9Type
	_               [8]byte
	Otghshcchar10   RegisterOtghshcchar10Type
	Otghshcsplt10   RegisterOtghshcsplt10Type
	Otghshcint10    RegisterOtghshcint10Type
	Otghshcintmsk10 RegisterOtghshcintmsk10Type
	Otghshctsiz10   RegisterOtghshctsiz10Type
	Otghshcdma10    RegisterOtghshcdma10Type
	_               [8]byte
	Otghshcchar11   RegisterOtghshcchar11Type
	Otghshcsplt11   RegisterOtghshcsplt11Type
	Otghshcint11    RegisterOtghshcint11Type
	Otghshcintmsk11 RegisterOtghshcintmsk11Type
	Otghshctsiz11   RegisterOtghshctsiz11Type
	Otghshcdma11    RegisterOtghshcdma11Type
	Otghshcchar12   RegisterOtghshcchar12Type
	Otghshcsplt12   RegisterOtghshcsplt12Type
	Otghshcint12    RegisterOtghshcint12Type
	Otghshcintmsk12 RegisterOtghshcintmsk12Type
	Otghshctsiz12   RegisterOtghshctsiz12Type
	Otghshcdma12    RegisterOtghshcdma12Type
	Otghshcchar13   RegisterOtghshcchar13Type
	Otghshcsplt13   RegisterOtghshcsplt13Type
	Otghshcint13    RegisterOtghshcint13Type
	Otghshcintmsk13 RegisterOtghshcintmsk13Type
	Otghshctsiz13   RegisterOtghshctsiz13Type
	Otghshcdma13    RegisterOtghshcdma13Type
	Otghshcchar14   RegisterOtghshcchar14Type
	Otghshcsplt14   RegisterOtghshcsplt14Type
	Otghshcint14    RegisterOtghshcint14Type
	Otghshcintmsk14 RegisterOtghshcintmsk14Type
	Otghshctsiz14   RegisterOtghshctsiz14Type
	Otghshcdma14    RegisterOtghshcdma14Type
	Otghshcchar15   RegisterOtghshcchar15Type
	Otghshcsplt15   RegisterOtghshcsplt15Type
	Otghshcint15    RegisterOtghshcint15Type
	Otghshcintmsk15 RegisterOtghshcintmsk15Type
	Otghshctsiz15   RegisterOtghshctsiz15Type
	Otghshcdma15    RegisterOtghshcdma15Type
}

// RegisterOtghshcfgType OTG_HS host configuration register
type RegisterOtghshcfgType uint32

func (r *RegisterOtghshcfgType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcfgType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcfgType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcfgType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcfgType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcfgFieldFslspcsShift = 0
	RegisterOtghshcfgFieldFslspcsMask  = 0x3
)

// GetFslspcs FS/LS PHY clock select
func (r *RegisterOtghshcfgType) GetFslspcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcfgFieldFslspcsMask) >> RegisterOtghshcfgFieldFslspcsShift)
}

// SetFslspcs FS/LS PHY clock select
func (r *RegisterOtghshcfgType) SetFslspcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcfgFieldFslspcsMask)|(uint32(value)<<RegisterOtghshcfgFieldFslspcsShift))
}

const (
	RegisterOtghshcfgFieldFslssShift = 2
	RegisterOtghshcfgFieldFslssMask  = 0x4
)

// GetFslss FS- and LS-only support
func (r *RegisterOtghshcfgType) GetFslss() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcfgFieldFslssMask) != 0
}

// RegisterOtghshfirType OTG_HS Host frame interval register
type RegisterOtghshfirType uint32

func (r *RegisterOtghshfirType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshfirType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshfirType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshfirType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshfirType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshfirFieldFrivlShift = 0
	RegisterOtghshfirFieldFrivlMask  = 0xffff
)

// GetFrivl Frame interval
func (r *RegisterOtghshfirType) GetFrivl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshfirFieldFrivlMask) >> RegisterOtghshfirFieldFrivlShift)
}

// SetFrivl Frame interval
func (r *RegisterOtghshfirType) SetFrivl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshfirFieldFrivlMask)|(uint32(value)<<RegisterOtghshfirFieldFrivlShift))
}

// RegisterOtghshfnumType OTG_HS host frame number/frame time remaining register
type RegisterOtghshfnumType uint32

func (r *RegisterOtghshfnumType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshfnumType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshfnumType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshfnumType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshfnumType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshfnumFieldFrnumShift = 0
	RegisterOtghshfnumFieldFrnumMask  = 0xffff
)

// GetFrnum Frame number
func (r *RegisterOtghshfnumType) GetFrnum() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshfnumFieldFrnumMask) >> RegisterOtghshfnumFieldFrnumShift)
}

// SetFrnum Frame number
func (r *RegisterOtghshfnumType) SetFrnum(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshfnumFieldFrnumMask)|(uint32(value)<<RegisterOtghshfnumFieldFrnumShift))
}

const (
	RegisterOtghshfnumFieldFtremShift = 16
	RegisterOtghshfnumFieldFtremMask  = 0xffff0000
)

// GetFtrem Frame time remaining
func (r *RegisterOtghshfnumType) GetFtrem() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshfnumFieldFtremMask) >> RegisterOtghshfnumFieldFtremShift)
}

// SetFtrem Frame time remaining
func (r *RegisterOtghshfnumType) SetFtrem(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshfnumFieldFtremMask)|(uint32(value)<<RegisterOtghshfnumFieldFtremShift))
}

// RegisterOtghshptxstsType OTG_HS_Host periodic transmit FIFO/queue status register
type RegisterOtghshptxstsType uint32

func (r *RegisterOtghshptxstsType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshptxstsType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshptxstsType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshptxstsType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshptxstsType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshptxstsFieldPtxfsavlShift = 0
	RegisterOtghshptxstsFieldPtxfsavlMask  = 0xffff
)

// GetPtxfsavl Periodic transmit data FIFO space available
func (r *RegisterOtghshptxstsType) GetPtxfsavl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshptxstsFieldPtxfsavlMask) >> RegisterOtghshptxstsFieldPtxfsavlShift)
}

// SetPtxfsavl Periodic transmit data FIFO space available
func (r *RegisterOtghshptxstsType) SetPtxfsavl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshptxstsFieldPtxfsavlMask)|(uint32(value)<<RegisterOtghshptxstsFieldPtxfsavlShift))
}

const (
	RegisterOtghshptxstsFieldPtxqsavShift = 16
	RegisterOtghshptxstsFieldPtxqsavMask  = 0xff0000
)

// GetPtxqsav Periodic transmit request queue space available
func (r *RegisterOtghshptxstsType) GetPtxqsav() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshptxstsFieldPtxqsavMask) >> RegisterOtghshptxstsFieldPtxqsavShift)
}

const (
	RegisterOtghshptxstsFieldPtxqtopShift = 24
	RegisterOtghshptxstsFieldPtxqtopMask  = 0xff000000
)

// GetPtxqtop Top of the periodic transmit request queue
func (r *RegisterOtghshptxstsType) GetPtxqtop() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshptxstsFieldPtxqtopMask) >> RegisterOtghshptxstsFieldPtxqtopShift)
}

// RegisterOtghshaintType OTG_HS Host all channels interrupt register
type RegisterOtghshaintType uint32

func (r *RegisterOtghshaintType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshaintType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshaintType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshaintType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshaintType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshaintFieldHaintShift = 0
	RegisterOtghshaintFieldHaintMask  = 0xffff
)

// GetHaint Channel interrupts
func (r *RegisterOtghshaintType) GetHaint() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshaintFieldHaintMask) >> RegisterOtghshaintFieldHaintShift)
}

// SetHaint Channel interrupts
func (r *RegisterOtghshaintType) SetHaint(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshaintFieldHaintMask)|(uint32(value)<<RegisterOtghshaintFieldHaintShift))
}

// RegisterOtghshaintmskType OTG_HS host all channels interrupt mask register
type RegisterOtghshaintmskType uint32

func (r *RegisterOtghshaintmskType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshaintmskType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshaintmskType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshaintmskType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshaintmskType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshaintmskFieldHaintmShift = 0
	RegisterOtghshaintmskFieldHaintmMask  = 0xffff
)

// GetHaintm Channel interrupt mask
func (r *RegisterOtghshaintmskType) GetHaintm() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshaintmskFieldHaintmMask) >> RegisterOtghshaintmskFieldHaintmShift)
}

// SetHaintm Channel interrupt mask
func (r *RegisterOtghshaintmskType) SetHaintm(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshaintmskFieldHaintmMask)|(uint32(value)<<RegisterOtghshaintmskFieldHaintmShift))
}

// RegisterOtghshprtType OTG_HS host port control and status register
type RegisterOtghshprtType uint32

func (r *RegisterOtghshprtType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshprtType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshprtType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshprtType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshprtType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshprtFieldPcstsShift = 0
	RegisterOtghshprtFieldPcstsMask  = 0x1
)

// GetPcsts Port connect status
func (r *RegisterOtghshprtType) GetPcsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPcstsMask) != 0
}

const (
	RegisterOtghshprtFieldPcdetShift = 1
	RegisterOtghshprtFieldPcdetMask  = 0x2
)

// GetPcdet Port connect detected
func (r *RegisterOtghshprtType) GetPcdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPcdetMask) != 0
}

// SetPcdet Port connect detected
func (r *RegisterOtghshprtType) SetPcdet(value bool) {
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
func (r *RegisterOtghshprtType) GetPena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPenaMask) != 0
}

// SetPena Port enable
func (r *RegisterOtghshprtType) SetPena(value bool) {
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
func (r *RegisterOtghshprtType) GetPenchng() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPenchngMask) != 0
}

// SetPenchng Port enable/disable change
func (r *RegisterOtghshprtType) SetPenchng(value bool) {
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
func (r *RegisterOtghshprtType) GetPoca() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPocaMask) != 0
}

const (
	RegisterOtghshprtFieldPocchngShift = 5
	RegisterOtghshprtFieldPocchngMask  = 0x20
)

// GetPocchng Port overcurrent change
func (r *RegisterOtghshprtType) GetPocchng() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPocchngMask) != 0
}

// SetPocchng Port overcurrent change
func (r *RegisterOtghshprtType) SetPocchng(value bool) {
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
func (r *RegisterOtghshprtType) GetPres() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPresMask) != 0
}

// SetPres Port resume
func (r *RegisterOtghshprtType) SetPres(value bool) {
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
func (r *RegisterOtghshprtType) GetPsusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPsuspMask) != 0
}

// SetPsusp Port suspend
func (r *RegisterOtghshprtType) SetPsusp(value bool) {
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
func (r *RegisterOtghshprtType) GetPrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPrstMask) != 0
}

// SetPrst Port reset
func (r *RegisterOtghshprtType) SetPrst(value bool) {
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
func (r *RegisterOtghshprtType) GetPlsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPlstsMask) >> RegisterOtghshprtFieldPlstsShift)
}

const (
	RegisterOtghshprtFieldPpwrShift = 12
	RegisterOtghshprtFieldPpwrMask  = 0x1000
)

// GetPpwr Port power
func (r *RegisterOtghshprtType) GetPpwr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPpwrMask) != 0
}

// SetPpwr Port power
func (r *RegisterOtghshprtType) SetPpwr(value bool) {
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
func (r *RegisterOtghshprtType) GetPtctl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPtctlMask) >> RegisterOtghshprtFieldPtctlShift)
}

// SetPtctl Port test control
func (r *RegisterOtghshprtType) SetPtctl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshprtFieldPtctlMask)|(uint32(value)<<RegisterOtghshprtFieldPtctlShift))
}

const (
	RegisterOtghshprtFieldPspdShift = 17
	RegisterOtghshprtFieldPspdMask  = 0x60000
)

// GetPspd Port speed
func (r *RegisterOtghshprtType) GetPspd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshprtFieldPspdMask) >> RegisterOtghshprtFieldPspdShift)
}

// RegisterOtghshcchar0Type OTG_HS host channel-0 characteristics register
type RegisterOtghshcchar0Type uint32

func (r *RegisterOtghshcchar0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar0FieldMpsizShift = 0
	RegisterOtghshcchar0FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar0Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldMpsizMask) >> RegisterOtghshcchar0FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar0Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar0FieldMpsizShift))
}

const (
	RegisterOtghshcchar0FieldEpnumShift = 11
	RegisterOtghshcchar0FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar0Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldEpnumMask) >> RegisterOtghshcchar0FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar0Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar0FieldEpnumShift))
}

const (
	RegisterOtghshcchar0FieldEpdirShift = 15
	RegisterOtghshcchar0FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar0Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar0Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar0Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar0Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar0Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldEptypMask) >> RegisterOtghshcchar0FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar0Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar0FieldEptypShift))
}

const (
	RegisterOtghshcchar0FieldMcShift = 20
	RegisterOtghshcchar0FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar0Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldMcMask) >> RegisterOtghshcchar0FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar0Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldMcMask)|(uint32(value)<<RegisterOtghshcchar0FieldMcShift))
}

const (
	RegisterOtghshcchar0FieldDadShift = 22
	RegisterOtghshcchar0FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar0Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldDadMask) >> RegisterOtghshcchar0FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar0Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldDadMask)|(uint32(value)<<RegisterOtghshcchar0FieldDadShift))
}

const (
	RegisterOtghshcchar0FieldOddfrmShift = 29
	RegisterOtghshcchar0FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar0Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar0Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar0Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar0Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar0Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar0FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar0Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar0FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar0FieldChenaMask)
	}
}

// RegisterOtghshcsplt0Type OTG_HS host channel-0 split control register
type RegisterOtghshcsplt0Type uint32

func (r *RegisterOtghshcsplt0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt0FieldPrtaddrShift = 0
	RegisterOtghshcsplt0FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt0Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt0FieldPrtaddrMask) >> RegisterOtghshcsplt0FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt0Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt0FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt0FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt0FieldHubaddrShift = 7
	RegisterOtghshcsplt0FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt0Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt0FieldHubaddrMask) >> RegisterOtghshcsplt0FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt0Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt0FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt0FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt0FieldXactposShift = 14
	RegisterOtghshcsplt0FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt0Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt0FieldXactposMask) >> RegisterOtghshcsplt0FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt0Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt0FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt0FieldXactposShift))
}

const (
	RegisterOtghshcsplt0FieldComplspltShift = 16
	RegisterOtghshcsplt0FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt0Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt0FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt0Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt0Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt0FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt0Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt0FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt0FieldSplitenMask)
	}
}

// RegisterOtghshcint0Type OTG_HS host channel-11 interrupt register
type RegisterOtghshcint0Type uint32

func (r *RegisterOtghshcint0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint0FieldXfrcShift = 0
	RegisterOtghshcint0FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint0Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint0Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint0Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint0Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint0Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint0Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint0Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint0Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint0Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint0Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint0Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint0Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint0Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint0Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint0Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint0Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint0Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint0Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint0Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint0Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint0Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint0FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint0Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint0FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint0FieldDterrMask)
	}
}

// RegisterOtghshcintmsk0Type OTG_HS host channel-11 interrupt mask register
type RegisterOtghshcintmsk0Type uint32

func (r *RegisterOtghshcintmsk0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk0FieldXfrcmShift = 0
	RegisterOtghshcintmsk0FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk0Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk0Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk0Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk0Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk0Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk0Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk0Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk0Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk0Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk0Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk0Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk0Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk0Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk0Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk0Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk0Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk0Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk0Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk0Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk0Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk0Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk0FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk0Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk0FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk0FieldDterrmMask)
	}
}

// RegisterOtghshctsiz0Type OTG_HS host channel-11 transfer size register
type RegisterOtghshctsiz0Type uint32

func (r *RegisterOtghshctsiz0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz0FieldXfrsizShift = 0
	RegisterOtghshctsiz0FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz0Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz0FieldXfrsizMask) >> RegisterOtghshctsiz0FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz0Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz0FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz0FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz0FieldPktcntShift = 19
	RegisterOtghshctsiz0FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz0Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz0FieldPktcntMask) >> RegisterOtghshctsiz0FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz0Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz0FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz0FieldPktcntShift))
}

const (
	RegisterOtghshctsiz0FieldDpidShift = 29
	RegisterOtghshctsiz0FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz0Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz0FieldDpidMask) >> RegisterOtghshctsiz0FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz0Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz0FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz0FieldDpidShift))
}

// RegisterOtghshcdma0Type OTG_HS host channel-0 DMA address register
type RegisterOtghshcdma0Type uint32

func (r *RegisterOtghshcdma0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma0FieldDmaaddrShift = 0
	RegisterOtghshcdma0FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma0Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma0FieldDmaaddrMask) >> RegisterOtghshcdma0FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma0Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma0FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma0FieldDmaaddrShift))
}

// RegisterOtghshcchar1Type OTG_HS host channel-1 characteristics register
type RegisterOtghshcchar1Type uint32

func (r *RegisterOtghshcchar1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar1FieldMpsizShift = 0
	RegisterOtghshcchar1FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar1Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldMpsizMask) >> RegisterOtghshcchar1FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar1Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar1FieldMpsizShift))
}

const (
	RegisterOtghshcchar1FieldEpnumShift = 11
	RegisterOtghshcchar1FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar1Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldEpnumMask) >> RegisterOtghshcchar1FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar1Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar1FieldEpnumShift))
}

const (
	RegisterOtghshcchar1FieldEpdirShift = 15
	RegisterOtghshcchar1FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar1Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar1Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar1Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar1Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar1Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldEptypMask) >> RegisterOtghshcchar1FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar1Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar1FieldEptypShift))
}

const (
	RegisterOtghshcchar1FieldMcShift = 20
	RegisterOtghshcchar1FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar1Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldMcMask) >> RegisterOtghshcchar1FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar1Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldMcMask)|(uint32(value)<<RegisterOtghshcchar1FieldMcShift))
}

const (
	RegisterOtghshcchar1FieldDadShift = 22
	RegisterOtghshcchar1FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar1Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldDadMask) >> RegisterOtghshcchar1FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar1Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldDadMask)|(uint32(value)<<RegisterOtghshcchar1FieldDadShift))
}

const (
	RegisterOtghshcchar1FieldOddfrmShift = 29
	RegisterOtghshcchar1FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar1Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar1Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar1Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar1Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar1Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar1FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar1Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar1FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar1FieldChenaMask)
	}
}

// RegisterOtghshcsplt1Type OTG_HS host channel-1 split control register
type RegisterOtghshcsplt1Type uint32

func (r *RegisterOtghshcsplt1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt1FieldPrtaddrShift = 0
	RegisterOtghshcsplt1FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt1Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt1FieldPrtaddrMask) >> RegisterOtghshcsplt1FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt1Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt1FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt1FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt1FieldHubaddrShift = 7
	RegisterOtghshcsplt1FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt1Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt1FieldHubaddrMask) >> RegisterOtghshcsplt1FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt1Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt1FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt1FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt1FieldXactposShift = 14
	RegisterOtghshcsplt1FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt1Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt1FieldXactposMask) >> RegisterOtghshcsplt1FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt1Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt1FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt1FieldXactposShift))
}

const (
	RegisterOtghshcsplt1FieldComplspltShift = 16
	RegisterOtghshcsplt1FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt1Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt1FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt1Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt1Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt1FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt1Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt1FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt1FieldSplitenMask)
	}
}

// RegisterOtghshcint1Type OTG_HS host channel-1 interrupt register
type RegisterOtghshcint1Type uint32

func (r *RegisterOtghshcint1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint1FieldXfrcShift = 0
	RegisterOtghshcint1FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint1Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint1Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint1Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint1Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint1Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint1Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint1Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint1Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint1Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint1Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint1Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint1Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint1Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint1Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint1Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint1Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint1Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint1Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint1Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint1Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint1Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint1FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint1Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint1FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint1FieldDterrMask)
	}
}

// RegisterOtghshcintmsk1Type OTG_HS host channel-1 interrupt mask register
type RegisterOtghshcintmsk1Type uint32

func (r *RegisterOtghshcintmsk1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk1FieldXfrcmShift = 0
	RegisterOtghshcintmsk1FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk1Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk1Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk1Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk1Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk1Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk1Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk1Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk1Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk1Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk1Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk1Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk1Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk1Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk1Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk1Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk1Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk1Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk1Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk1Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk1Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk1Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk1FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk1Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk1FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk1FieldDterrmMask)
	}
}

// RegisterOtghshctsiz1Type OTG_HS host channel-1 transfer size register
type RegisterOtghshctsiz1Type uint32

func (r *RegisterOtghshctsiz1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz1FieldXfrsizShift = 0
	RegisterOtghshctsiz1FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz1Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz1FieldXfrsizMask) >> RegisterOtghshctsiz1FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz1Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz1FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz1FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz1FieldPktcntShift = 19
	RegisterOtghshctsiz1FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz1Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz1FieldPktcntMask) >> RegisterOtghshctsiz1FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz1Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz1FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz1FieldPktcntShift))
}

const (
	RegisterOtghshctsiz1FieldDpidShift = 29
	RegisterOtghshctsiz1FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz1Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz1FieldDpidMask) >> RegisterOtghshctsiz1FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz1Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz1FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz1FieldDpidShift))
}

// RegisterOtghshcdma1Type OTG_HS host channel-1 DMA address register
type RegisterOtghshcdma1Type uint32

func (r *RegisterOtghshcdma1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma1FieldDmaaddrShift = 0
	RegisterOtghshcdma1FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma1Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma1FieldDmaaddrMask) >> RegisterOtghshcdma1FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma1Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma1FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma1FieldDmaaddrShift))
}

// RegisterOtghshcchar2Type OTG_HS host channel-2 characteristics register
type RegisterOtghshcchar2Type uint32

func (r *RegisterOtghshcchar2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar2FieldMpsizShift = 0
	RegisterOtghshcchar2FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar2Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldMpsizMask) >> RegisterOtghshcchar2FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar2Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar2FieldMpsizShift))
}

const (
	RegisterOtghshcchar2FieldEpnumShift = 11
	RegisterOtghshcchar2FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar2Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldEpnumMask) >> RegisterOtghshcchar2FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar2Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar2FieldEpnumShift))
}

const (
	RegisterOtghshcchar2FieldEpdirShift = 15
	RegisterOtghshcchar2FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar2Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar2Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar2Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar2Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar2Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldEptypMask) >> RegisterOtghshcchar2FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar2Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar2FieldEptypShift))
}

const (
	RegisterOtghshcchar2FieldMcShift = 20
	RegisterOtghshcchar2FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar2Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldMcMask) >> RegisterOtghshcchar2FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar2Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldMcMask)|(uint32(value)<<RegisterOtghshcchar2FieldMcShift))
}

const (
	RegisterOtghshcchar2FieldDadShift = 22
	RegisterOtghshcchar2FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar2Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldDadMask) >> RegisterOtghshcchar2FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar2Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldDadMask)|(uint32(value)<<RegisterOtghshcchar2FieldDadShift))
}

const (
	RegisterOtghshcchar2FieldOddfrmShift = 29
	RegisterOtghshcchar2FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar2Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar2Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar2Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar2Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar2Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar2FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar2Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar2FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar2FieldChenaMask)
	}
}

// RegisterOtghshcsplt2Type OTG_HS host channel-2 split control register
type RegisterOtghshcsplt2Type uint32

func (r *RegisterOtghshcsplt2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt2FieldPrtaddrShift = 0
	RegisterOtghshcsplt2FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt2Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt2FieldPrtaddrMask) >> RegisterOtghshcsplt2FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt2Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt2FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt2FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt2FieldHubaddrShift = 7
	RegisterOtghshcsplt2FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt2Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt2FieldHubaddrMask) >> RegisterOtghshcsplt2FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt2Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt2FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt2FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt2FieldXactposShift = 14
	RegisterOtghshcsplt2FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt2Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt2FieldXactposMask) >> RegisterOtghshcsplt2FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt2Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt2FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt2FieldXactposShift))
}

const (
	RegisterOtghshcsplt2FieldComplspltShift = 16
	RegisterOtghshcsplt2FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt2Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt2FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt2Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt2Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt2FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt2Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt2FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt2FieldSplitenMask)
	}
}

// RegisterOtghshcint2Type OTG_HS host channel-2 interrupt register
type RegisterOtghshcint2Type uint32

func (r *RegisterOtghshcint2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint2FieldXfrcShift = 0
	RegisterOtghshcint2FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint2Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint2Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint2Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint2Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint2Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint2Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint2Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint2Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint2Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint2Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint2Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint2Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint2Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint2Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint2Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint2Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint2Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint2Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint2Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint2Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint2Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint2FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint2Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint2FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint2FieldDterrMask)
	}
}

// RegisterOtghshcintmsk2Type OTG_HS host channel-2 interrupt mask register
type RegisterOtghshcintmsk2Type uint32

func (r *RegisterOtghshcintmsk2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk2FieldXfrcmShift = 0
	RegisterOtghshcintmsk2FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk2Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk2Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk2Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk2Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk2Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk2Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk2Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk2Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk2Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk2Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk2Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk2Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk2Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk2Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk2Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk2Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk2Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk2Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk2Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk2Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk2Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk2FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk2Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk2FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk2FieldDterrmMask)
	}
}

// RegisterOtghshctsiz2Type OTG_HS host channel-2 transfer size register
type RegisterOtghshctsiz2Type uint32

func (r *RegisterOtghshctsiz2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz2FieldXfrsizShift = 0
	RegisterOtghshctsiz2FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz2Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz2FieldXfrsizMask) >> RegisterOtghshctsiz2FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz2Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz2FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz2FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz2FieldPktcntShift = 19
	RegisterOtghshctsiz2FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz2Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz2FieldPktcntMask) >> RegisterOtghshctsiz2FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz2Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz2FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz2FieldPktcntShift))
}

const (
	RegisterOtghshctsiz2FieldDpidShift = 29
	RegisterOtghshctsiz2FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz2Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz2FieldDpidMask) >> RegisterOtghshctsiz2FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz2Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz2FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz2FieldDpidShift))
}

// RegisterOtghshcdma2Type OTG_HS host channel-2 DMA address register
type RegisterOtghshcdma2Type uint32

func (r *RegisterOtghshcdma2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma2FieldDmaaddrShift = 0
	RegisterOtghshcdma2FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma2Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma2FieldDmaaddrMask) >> RegisterOtghshcdma2FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma2Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma2FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma2FieldDmaaddrShift))
}

// RegisterOtghshcchar3Type OTG_HS host channel-3 characteristics register
type RegisterOtghshcchar3Type uint32

func (r *RegisterOtghshcchar3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar3FieldMpsizShift = 0
	RegisterOtghshcchar3FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar3Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldMpsizMask) >> RegisterOtghshcchar3FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar3Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar3FieldMpsizShift))
}

const (
	RegisterOtghshcchar3FieldEpnumShift = 11
	RegisterOtghshcchar3FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar3Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldEpnumMask) >> RegisterOtghshcchar3FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar3Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar3FieldEpnumShift))
}

const (
	RegisterOtghshcchar3FieldEpdirShift = 15
	RegisterOtghshcchar3FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar3Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar3Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar3Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar3Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar3Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldEptypMask) >> RegisterOtghshcchar3FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar3Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar3FieldEptypShift))
}

const (
	RegisterOtghshcchar3FieldMcShift = 20
	RegisterOtghshcchar3FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar3Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldMcMask) >> RegisterOtghshcchar3FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar3Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldMcMask)|(uint32(value)<<RegisterOtghshcchar3FieldMcShift))
}

const (
	RegisterOtghshcchar3FieldDadShift = 22
	RegisterOtghshcchar3FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar3Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldDadMask) >> RegisterOtghshcchar3FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar3Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldDadMask)|(uint32(value)<<RegisterOtghshcchar3FieldDadShift))
}

const (
	RegisterOtghshcchar3FieldOddfrmShift = 29
	RegisterOtghshcchar3FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar3Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar3Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar3Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar3Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar3Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar3FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar3Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar3FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar3FieldChenaMask)
	}
}

// RegisterOtghshcsplt3Type OTG_HS host channel-3 split control register
type RegisterOtghshcsplt3Type uint32

func (r *RegisterOtghshcsplt3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt3FieldPrtaddrShift = 0
	RegisterOtghshcsplt3FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt3Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt3FieldPrtaddrMask) >> RegisterOtghshcsplt3FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt3Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt3FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt3FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt3FieldHubaddrShift = 7
	RegisterOtghshcsplt3FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt3Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt3FieldHubaddrMask) >> RegisterOtghshcsplt3FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt3Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt3FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt3FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt3FieldXactposShift = 14
	RegisterOtghshcsplt3FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt3Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt3FieldXactposMask) >> RegisterOtghshcsplt3FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt3Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt3FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt3FieldXactposShift))
}

const (
	RegisterOtghshcsplt3FieldComplspltShift = 16
	RegisterOtghshcsplt3FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt3Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt3FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt3Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt3Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt3FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt3Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt3FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt3FieldSplitenMask)
	}
}

// RegisterOtghshcint3Type OTG_HS host channel-3 interrupt register
type RegisterOtghshcint3Type uint32

func (r *RegisterOtghshcint3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint3FieldXfrcShift = 0
	RegisterOtghshcint3FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint3Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint3Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint3Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint3Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint3Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint3Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint3Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint3Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint3Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint3Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint3Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint3Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint3Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint3Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint3Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint3Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint3Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint3Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint3Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint3Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint3Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint3FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint3Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint3FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint3FieldDterrMask)
	}
}

// RegisterOtghshcintmsk3Type OTG_HS host channel-3 interrupt mask register
type RegisterOtghshcintmsk3Type uint32

func (r *RegisterOtghshcintmsk3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk3FieldXfrcmShift = 0
	RegisterOtghshcintmsk3FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk3Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk3Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk3Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk3Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk3Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk3Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk3Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk3Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk3Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk3Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk3Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk3Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk3Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk3Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk3Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk3Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk3Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk3Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk3Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk3Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk3Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk3FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk3Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk3FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk3FieldDterrmMask)
	}
}

// RegisterOtghshctsiz3Type OTG_HS host channel-3 transfer size register
type RegisterOtghshctsiz3Type uint32

func (r *RegisterOtghshctsiz3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz3FieldXfrsizShift = 0
	RegisterOtghshctsiz3FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz3Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz3FieldXfrsizMask) >> RegisterOtghshctsiz3FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz3Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz3FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz3FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz3FieldPktcntShift = 19
	RegisterOtghshctsiz3FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz3Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz3FieldPktcntMask) >> RegisterOtghshctsiz3FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz3Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz3FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz3FieldPktcntShift))
}

const (
	RegisterOtghshctsiz3FieldDpidShift = 29
	RegisterOtghshctsiz3FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz3Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz3FieldDpidMask) >> RegisterOtghshctsiz3FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz3Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz3FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz3FieldDpidShift))
}

// RegisterOtghshcdma3Type OTG_HS host channel-3 DMA address register
type RegisterOtghshcdma3Type uint32

func (r *RegisterOtghshcdma3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma3FieldDmaaddrShift = 0
	RegisterOtghshcdma3FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma3Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma3FieldDmaaddrMask) >> RegisterOtghshcdma3FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma3Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma3FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma3FieldDmaaddrShift))
}

// RegisterOtghshcchar4Type OTG_HS host channel-4 characteristics register
type RegisterOtghshcchar4Type uint32

func (r *RegisterOtghshcchar4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar4FieldMpsizShift = 0
	RegisterOtghshcchar4FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar4Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldMpsizMask) >> RegisterOtghshcchar4FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar4Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar4FieldMpsizShift))
}

const (
	RegisterOtghshcchar4FieldEpnumShift = 11
	RegisterOtghshcchar4FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar4Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldEpnumMask) >> RegisterOtghshcchar4FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar4Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar4FieldEpnumShift))
}

const (
	RegisterOtghshcchar4FieldEpdirShift = 15
	RegisterOtghshcchar4FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar4Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar4Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar4Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar4Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar4Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldEptypMask) >> RegisterOtghshcchar4FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar4Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar4FieldEptypShift))
}

const (
	RegisterOtghshcchar4FieldMcShift = 20
	RegisterOtghshcchar4FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar4Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldMcMask) >> RegisterOtghshcchar4FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar4Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldMcMask)|(uint32(value)<<RegisterOtghshcchar4FieldMcShift))
}

const (
	RegisterOtghshcchar4FieldDadShift = 22
	RegisterOtghshcchar4FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar4Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldDadMask) >> RegisterOtghshcchar4FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar4Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldDadMask)|(uint32(value)<<RegisterOtghshcchar4FieldDadShift))
}

const (
	RegisterOtghshcchar4FieldOddfrmShift = 29
	RegisterOtghshcchar4FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar4Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar4Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar4Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar4Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar4Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar4FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar4Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar4FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar4FieldChenaMask)
	}
}

// RegisterOtghshcsplt4Type OTG_HS host channel-4 split control register
type RegisterOtghshcsplt4Type uint32

func (r *RegisterOtghshcsplt4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt4FieldPrtaddrShift = 0
	RegisterOtghshcsplt4FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt4Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt4FieldPrtaddrMask) >> RegisterOtghshcsplt4FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt4Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt4FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt4FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt4FieldHubaddrShift = 7
	RegisterOtghshcsplt4FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt4Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt4FieldHubaddrMask) >> RegisterOtghshcsplt4FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt4Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt4FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt4FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt4FieldXactposShift = 14
	RegisterOtghshcsplt4FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt4Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt4FieldXactposMask) >> RegisterOtghshcsplt4FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt4Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt4FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt4FieldXactposShift))
}

const (
	RegisterOtghshcsplt4FieldComplspltShift = 16
	RegisterOtghshcsplt4FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt4Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt4FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt4Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt4Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt4FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt4Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt4FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt4FieldSplitenMask)
	}
}

// RegisterOtghshcint4Type OTG_HS host channel-4 interrupt register
type RegisterOtghshcint4Type uint32

func (r *RegisterOtghshcint4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint4FieldXfrcShift = 0
	RegisterOtghshcint4FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint4Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint4Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint4Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint4Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint4Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint4Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint4Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint4Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint4Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint4Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint4Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint4Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint4Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint4Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint4Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint4Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint4Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint4Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint4Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint4Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint4Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint4FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint4Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint4FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint4FieldDterrMask)
	}
}

// RegisterOtghshcintmsk4Type OTG_HS host channel-4 interrupt mask register
type RegisterOtghshcintmsk4Type uint32

func (r *RegisterOtghshcintmsk4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk4FieldXfrcmShift = 0
	RegisterOtghshcintmsk4FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk4Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk4Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk4Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk4Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk4Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk4Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk4Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk4Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk4Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk4Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk4Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk4Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk4Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk4Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk4Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk4Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk4Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk4Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk4Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk4Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk4Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk4FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk4Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk4FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk4FieldDterrmMask)
	}
}

// RegisterOtghshctsiz4Type OTG_HS host channel-4 transfer size register
type RegisterOtghshctsiz4Type uint32

func (r *RegisterOtghshctsiz4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz4FieldXfrsizShift = 0
	RegisterOtghshctsiz4FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz4Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz4FieldXfrsizMask) >> RegisterOtghshctsiz4FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz4Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz4FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz4FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz4FieldPktcntShift = 19
	RegisterOtghshctsiz4FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz4Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz4FieldPktcntMask) >> RegisterOtghshctsiz4FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz4Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz4FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz4FieldPktcntShift))
}

const (
	RegisterOtghshctsiz4FieldDpidShift = 29
	RegisterOtghshctsiz4FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz4Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz4FieldDpidMask) >> RegisterOtghshctsiz4FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz4Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz4FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz4FieldDpidShift))
}

// RegisterOtghshcdma4Type OTG_HS host channel-4 DMA address register
type RegisterOtghshcdma4Type uint32

func (r *RegisterOtghshcdma4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma4FieldDmaaddrShift = 0
	RegisterOtghshcdma4FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma4Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma4FieldDmaaddrMask) >> RegisterOtghshcdma4FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma4Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma4FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma4FieldDmaaddrShift))
}

// RegisterOtghshcchar5Type OTG_HS host channel-5 characteristics register
type RegisterOtghshcchar5Type uint32

func (r *RegisterOtghshcchar5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar5FieldMpsizShift = 0
	RegisterOtghshcchar5FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar5Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldMpsizMask) >> RegisterOtghshcchar5FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar5Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar5FieldMpsizShift))
}

const (
	RegisterOtghshcchar5FieldEpnumShift = 11
	RegisterOtghshcchar5FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar5Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldEpnumMask) >> RegisterOtghshcchar5FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar5Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar5FieldEpnumShift))
}

const (
	RegisterOtghshcchar5FieldEpdirShift = 15
	RegisterOtghshcchar5FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar5Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar5Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar5Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar5Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar5Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldEptypMask) >> RegisterOtghshcchar5FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar5Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar5FieldEptypShift))
}

const (
	RegisterOtghshcchar5FieldMcShift = 20
	RegisterOtghshcchar5FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar5Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldMcMask) >> RegisterOtghshcchar5FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar5Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldMcMask)|(uint32(value)<<RegisterOtghshcchar5FieldMcShift))
}

const (
	RegisterOtghshcchar5FieldDadShift = 22
	RegisterOtghshcchar5FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar5Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldDadMask) >> RegisterOtghshcchar5FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar5Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldDadMask)|(uint32(value)<<RegisterOtghshcchar5FieldDadShift))
}

const (
	RegisterOtghshcchar5FieldOddfrmShift = 29
	RegisterOtghshcchar5FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar5Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar5Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar5Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar5Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar5Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar5FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar5Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar5FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar5FieldChenaMask)
	}
}

// RegisterOtghshcsplt5Type OTG_HS host channel-5 split control register
type RegisterOtghshcsplt5Type uint32

func (r *RegisterOtghshcsplt5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt5FieldPrtaddrShift = 0
	RegisterOtghshcsplt5FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt5Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt5FieldPrtaddrMask) >> RegisterOtghshcsplt5FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt5Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt5FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt5FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt5FieldHubaddrShift = 7
	RegisterOtghshcsplt5FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt5Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt5FieldHubaddrMask) >> RegisterOtghshcsplt5FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt5Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt5FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt5FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt5FieldXactposShift = 14
	RegisterOtghshcsplt5FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt5Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt5FieldXactposMask) >> RegisterOtghshcsplt5FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt5Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt5FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt5FieldXactposShift))
}

const (
	RegisterOtghshcsplt5FieldComplspltShift = 16
	RegisterOtghshcsplt5FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt5Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt5FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt5Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt5Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt5FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt5Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt5FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt5FieldSplitenMask)
	}
}

// RegisterOtghshcint5Type OTG_HS host channel-5 interrupt register
type RegisterOtghshcint5Type uint32

func (r *RegisterOtghshcint5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint5FieldXfrcShift = 0
	RegisterOtghshcint5FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint5Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint5Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint5Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint5Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint5Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint5Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint5Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint5Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint5Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint5Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint5Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint5Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint5Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint5Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint5Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint5Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint5Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint5Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint5Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint5Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint5Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint5FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint5Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint5FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint5FieldDterrMask)
	}
}

// RegisterOtghshcintmsk5Type OTG_HS host channel-5 interrupt mask register
type RegisterOtghshcintmsk5Type uint32

func (r *RegisterOtghshcintmsk5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk5FieldXfrcmShift = 0
	RegisterOtghshcintmsk5FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk5Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk5Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk5Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk5Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk5Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk5Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk5Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk5Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk5Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk5Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk5Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk5Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk5Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk5Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk5Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk5Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk5Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk5Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk5Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk5Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk5Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk5FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk5Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk5FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk5FieldDterrmMask)
	}
}

// RegisterOtghshctsiz5Type OTG_HS host channel-5 transfer size register
type RegisterOtghshctsiz5Type uint32

func (r *RegisterOtghshctsiz5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz5FieldXfrsizShift = 0
	RegisterOtghshctsiz5FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz5Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz5FieldXfrsizMask) >> RegisterOtghshctsiz5FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz5Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz5FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz5FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz5FieldPktcntShift = 19
	RegisterOtghshctsiz5FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz5Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz5FieldPktcntMask) >> RegisterOtghshctsiz5FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz5Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz5FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz5FieldPktcntShift))
}

const (
	RegisterOtghshctsiz5FieldDpidShift = 29
	RegisterOtghshctsiz5FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz5Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz5FieldDpidMask) >> RegisterOtghshctsiz5FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz5Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz5FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz5FieldDpidShift))
}

// RegisterOtghshcdma5Type OTG_HS host channel-5 DMA address register
type RegisterOtghshcdma5Type uint32

func (r *RegisterOtghshcdma5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma5FieldDmaaddrShift = 0
	RegisterOtghshcdma5FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma5Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma5FieldDmaaddrMask) >> RegisterOtghshcdma5FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma5Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma5FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma5FieldDmaaddrShift))
}

// RegisterOtghshcchar6Type OTG_HS host channel-6 characteristics register
type RegisterOtghshcchar6Type uint32

func (r *RegisterOtghshcchar6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar6FieldMpsizShift = 0
	RegisterOtghshcchar6FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar6Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldMpsizMask) >> RegisterOtghshcchar6FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar6Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar6FieldMpsizShift))
}

const (
	RegisterOtghshcchar6FieldEpnumShift = 11
	RegisterOtghshcchar6FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar6Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldEpnumMask) >> RegisterOtghshcchar6FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar6Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar6FieldEpnumShift))
}

const (
	RegisterOtghshcchar6FieldEpdirShift = 15
	RegisterOtghshcchar6FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar6Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar6Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar6Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar6Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar6Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldEptypMask) >> RegisterOtghshcchar6FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar6Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar6FieldEptypShift))
}

const (
	RegisterOtghshcchar6FieldMcShift = 20
	RegisterOtghshcchar6FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar6Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldMcMask) >> RegisterOtghshcchar6FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar6Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldMcMask)|(uint32(value)<<RegisterOtghshcchar6FieldMcShift))
}

const (
	RegisterOtghshcchar6FieldDadShift = 22
	RegisterOtghshcchar6FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar6Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldDadMask) >> RegisterOtghshcchar6FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar6Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldDadMask)|(uint32(value)<<RegisterOtghshcchar6FieldDadShift))
}

const (
	RegisterOtghshcchar6FieldOddfrmShift = 29
	RegisterOtghshcchar6FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar6Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar6Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar6Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar6Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar6Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar6FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar6Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar6FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar6FieldChenaMask)
	}
}

// RegisterOtghshcsplt6Type OTG_HS host channel-6 split control register
type RegisterOtghshcsplt6Type uint32

func (r *RegisterOtghshcsplt6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt6FieldPrtaddrShift = 0
	RegisterOtghshcsplt6FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt6Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt6FieldPrtaddrMask) >> RegisterOtghshcsplt6FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt6Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt6FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt6FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt6FieldHubaddrShift = 7
	RegisterOtghshcsplt6FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt6Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt6FieldHubaddrMask) >> RegisterOtghshcsplt6FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt6Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt6FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt6FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt6FieldXactposShift = 14
	RegisterOtghshcsplt6FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt6Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt6FieldXactposMask) >> RegisterOtghshcsplt6FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt6Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt6FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt6FieldXactposShift))
}

const (
	RegisterOtghshcsplt6FieldComplspltShift = 16
	RegisterOtghshcsplt6FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt6Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt6FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt6Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt6Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt6FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt6Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt6FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt6FieldSplitenMask)
	}
}

// RegisterOtghshcint6Type OTG_HS host channel-6 interrupt register
type RegisterOtghshcint6Type uint32

func (r *RegisterOtghshcint6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint6FieldXfrcShift = 0
	RegisterOtghshcint6FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint6Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint6Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint6Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint6Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint6Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint6Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint6Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint6Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint6Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint6Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint6Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint6Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint6Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint6Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint6Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint6Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint6Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint6Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint6Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint6Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint6Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint6FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint6Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint6FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint6FieldDterrMask)
	}
}

// RegisterOtghshcintmsk6Type OTG_HS host channel-6 interrupt mask register
type RegisterOtghshcintmsk6Type uint32

func (r *RegisterOtghshcintmsk6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk6FieldXfrcmShift = 0
	RegisterOtghshcintmsk6FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk6Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk6Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk6Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk6Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk6Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk6Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk6Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk6Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk6Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk6Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk6Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk6Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk6Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk6Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk6Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk6Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk6Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk6Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk6Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk6Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk6Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk6FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk6Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk6FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk6FieldDterrmMask)
	}
}

// RegisterOtghshctsiz6Type OTG_HS host channel-6 transfer size register
type RegisterOtghshctsiz6Type uint32

func (r *RegisterOtghshctsiz6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz6FieldXfrsizShift = 0
	RegisterOtghshctsiz6FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz6Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz6FieldXfrsizMask) >> RegisterOtghshctsiz6FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz6Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz6FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz6FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz6FieldPktcntShift = 19
	RegisterOtghshctsiz6FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz6Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz6FieldPktcntMask) >> RegisterOtghshctsiz6FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz6Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz6FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz6FieldPktcntShift))
}

const (
	RegisterOtghshctsiz6FieldDpidShift = 29
	RegisterOtghshctsiz6FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz6Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz6FieldDpidMask) >> RegisterOtghshctsiz6FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz6Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz6FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz6FieldDpidShift))
}

// RegisterOtghshcdma6Type OTG_HS host channel-6 DMA address register
type RegisterOtghshcdma6Type uint32

func (r *RegisterOtghshcdma6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma6FieldDmaaddrShift = 0
	RegisterOtghshcdma6FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma6Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma6FieldDmaaddrMask) >> RegisterOtghshcdma6FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma6Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma6FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma6FieldDmaaddrShift))
}

// RegisterOtghshcchar7Type OTG_HS host channel-7 characteristics register
type RegisterOtghshcchar7Type uint32

func (r *RegisterOtghshcchar7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar7FieldMpsizShift = 0
	RegisterOtghshcchar7FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar7Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldMpsizMask) >> RegisterOtghshcchar7FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar7Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar7FieldMpsizShift))
}

const (
	RegisterOtghshcchar7FieldEpnumShift = 11
	RegisterOtghshcchar7FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar7Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldEpnumMask) >> RegisterOtghshcchar7FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar7Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar7FieldEpnumShift))
}

const (
	RegisterOtghshcchar7FieldEpdirShift = 15
	RegisterOtghshcchar7FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar7Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar7Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar7Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar7Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar7Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldEptypMask) >> RegisterOtghshcchar7FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar7Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar7FieldEptypShift))
}

const (
	RegisterOtghshcchar7FieldMcShift = 20
	RegisterOtghshcchar7FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar7Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldMcMask) >> RegisterOtghshcchar7FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar7Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldMcMask)|(uint32(value)<<RegisterOtghshcchar7FieldMcShift))
}

const (
	RegisterOtghshcchar7FieldDadShift = 22
	RegisterOtghshcchar7FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar7Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldDadMask) >> RegisterOtghshcchar7FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar7Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldDadMask)|(uint32(value)<<RegisterOtghshcchar7FieldDadShift))
}

const (
	RegisterOtghshcchar7FieldOddfrmShift = 29
	RegisterOtghshcchar7FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar7Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar7Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar7Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar7Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar7Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar7FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar7Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar7FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar7FieldChenaMask)
	}
}

// RegisterOtghshcsplt7Type OTG_HS host channel-7 split control register
type RegisterOtghshcsplt7Type uint32

func (r *RegisterOtghshcsplt7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt7FieldPrtaddrShift = 0
	RegisterOtghshcsplt7FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt7Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt7FieldPrtaddrMask) >> RegisterOtghshcsplt7FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt7Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt7FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt7FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt7FieldHubaddrShift = 7
	RegisterOtghshcsplt7FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt7Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt7FieldHubaddrMask) >> RegisterOtghshcsplt7FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt7Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt7FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt7FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt7FieldXactposShift = 14
	RegisterOtghshcsplt7FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt7Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt7FieldXactposMask) >> RegisterOtghshcsplt7FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt7Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt7FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt7FieldXactposShift))
}

const (
	RegisterOtghshcsplt7FieldComplspltShift = 16
	RegisterOtghshcsplt7FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt7Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt7FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt7Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt7Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt7FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt7Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt7FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt7FieldSplitenMask)
	}
}

// RegisterOtghshcint7Type OTG_HS host channel-7 interrupt register
type RegisterOtghshcint7Type uint32

func (r *RegisterOtghshcint7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint7FieldXfrcShift = 0
	RegisterOtghshcint7FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint7Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint7Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint7Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint7Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint7Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint7Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint7Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint7Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint7Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint7Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint7Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint7Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint7Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint7Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint7Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint7Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint7Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint7Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint7Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint7Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint7Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint7FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint7Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint7FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint7FieldDterrMask)
	}
}

// RegisterOtghshcintmsk7Type OTG_HS host channel-7 interrupt mask register
type RegisterOtghshcintmsk7Type uint32

func (r *RegisterOtghshcintmsk7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk7FieldXfrcmShift = 0
	RegisterOtghshcintmsk7FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk7Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk7Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk7Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk7Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk7Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk7Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk7Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk7Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk7Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk7Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk7Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk7Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk7Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk7Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk7Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk7Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk7Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk7Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk7Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk7Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk7Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk7FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk7Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk7FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk7FieldDterrmMask)
	}
}

// RegisterOtghshctsiz7Type OTG_HS host channel-7 transfer size register
type RegisterOtghshctsiz7Type uint32

func (r *RegisterOtghshctsiz7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz7FieldXfrsizShift = 0
	RegisterOtghshctsiz7FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz7Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz7FieldXfrsizMask) >> RegisterOtghshctsiz7FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz7Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz7FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz7FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz7FieldPktcntShift = 19
	RegisterOtghshctsiz7FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz7Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz7FieldPktcntMask) >> RegisterOtghshctsiz7FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz7Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz7FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz7FieldPktcntShift))
}

const (
	RegisterOtghshctsiz7FieldDpidShift = 29
	RegisterOtghshctsiz7FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz7Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz7FieldDpidMask) >> RegisterOtghshctsiz7FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz7Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz7FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz7FieldDpidShift))
}

// RegisterOtghshcdma7Type OTG_HS host channel-7 DMA address register
type RegisterOtghshcdma7Type uint32

func (r *RegisterOtghshcdma7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma7FieldDmaaddrShift = 0
	RegisterOtghshcdma7FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma7Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma7FieldDmaaddrMask) >> RegisterOtghshcdma7FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma7Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma7FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma7FieldDmaaddrShift))
}

// RegisterOtghshcchar8Type OTG_HS host channel-8 characteristics register
type RegisterOtghshcchar8Type uint32

func (r *RegisterOtghshcchar8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar8FieldMpsizShift = 0
	RegisterOtghshcchar8FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar8Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldMpsizMask) >> RegisterOtghshcchar8FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar8Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar8FieldMpsizShift))
}

const (
	RegisterOtghshcchar8FieldEpnumShift = 11
	RegisterOtghshcchar8FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar8Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldEpnumMask) >> RegisterOtghshcchar8FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar8Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar8FieldEpnumShift))
}

const (
	RegisterOtghshcchar8FieldEpdirShift = 15
	RegisterOtghshcchar8FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar8Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar8Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar8Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar8Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar8Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldEptypMask) >> RegisterOtghshcchar8FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar8Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar8FieldEptypShift))
}

const (
	RegisterOtghshcchar8FieldMcShift = 20
	RegisterOtghshcchar8FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar8Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldMcMask) >> RegisterOtghshcchar8FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar8Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldMcMask)|(uint32(value)<<RegisterOtghshcchar8FieldMcShift))
}

const (
	RegisterOtghshcchar8FieldDadShift = 22
	RegisterOtghshcchar8FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar8Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldDadMask) >> RegisterOtghshcchar8FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar8Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldDadMask)|(uint32(value)<<RegisterOtghshcchar8FieldDadShift))
}

const (
	RegisterOtghshcchar8FieldOddfrmShift = 29
	RegisterOtghshcchar8FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar8Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar8Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar8Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar8Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar8Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar8FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar8Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar8FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar8FieldChenaMask)
	}
}

// RegisterOtghshcsplt8Type OTG_HS host channel-8 split control register
type RegisterOtghshcsplt8Type uint32

func (r *RegisterOtghshcsplt8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt8FieldPrtaddrShift = 0
	RegisterOtghshcsplt8FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt8Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt8FieldPrtaddrMask) >> RegisterOtghshcsplt8FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt8Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt8FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt8FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt8FieldHubaddrShift = 7
	RegisterOtghshcsplt8FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt8Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt8FieldHubaddrMask) >> RegisterOtghshcsplt8FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt8Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt8FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt8FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt8FieldXactposShift = 14
	RegisterOtghshcsplt8FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt8Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt8FieldXactposMask) >> RegisterOtghshcsplt8FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt8Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt8FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt8FieldXactposShift))
}

const (
	RegisterOtghshcsplt8FieldComplspltShift = 16
	RegisterOtghshcsplt8FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt8Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt8FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt8Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt8Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt8FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt8Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt8FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt8FieldSplitenMask)
	}
}

// RegisterOtghshcint8Type OTG_HS host channel-8 interrupt register
type RegisterOtghshcint8Type uint32

func (r *RegisterOtghshcint8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint8FieldXfrcShift = 0
	RegisterOtghshcint8FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint8Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint8Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint8Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint8Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint8Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint8Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint8Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint8Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint8Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint8Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint8Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint8Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint8Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint8Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint8Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint8Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint8Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint8Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint8Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint8Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint8Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint8FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint8Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint8FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint8FieldDterrMask)
	}
}

// RegisterOtghshcintmsk8Type OTG_HS host channel-8 interrupt mask register
type RegisterOtghshcintmsk8Type uint32

func (r *RegisterOtghshcintmsk8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk8FieldXfrcmShift = 0
	RegisterOtghshcintmsk8FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk8Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk8Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk8Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk8Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk8Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk8Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk8Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk8Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk8Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk8Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk8Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk8Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk8Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk8Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk8Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk8Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk8Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk8Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk8Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk8Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk8Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk8FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk8Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk8FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk8FieldDterrmMask)
	}
}

// RegisterOtghshctsiz8Type OTG_HS host channel-8 transfer size register
type RegisterOtghshctsiz8Type uint32

func (r *RegisterOtghshctsiz8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz8FieldXfrsizShift = 0
	RegisterOtghshctsiz8FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz8Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz8FieldXfrsizMask) >> RegisterOtghshctsiz8FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz8Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz8FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz8FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz8FieldPktcntShift = 19
	RegisterOtghshctsiz8FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz8Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz8FieldPktcntMask) >> RegisterOtghshctsiz8FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz8Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz8FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz8FieldPktcntShift))
}

const (
	RegisterOtghshctsiz8FieldDpidShift = 29
	RegisterOtghshctsiz8FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz8Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz8FieldDpidMask) >> RegisterOtghshctsiz8FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz8Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz8FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz8FieldDpidShift))
}

// RegisterOtghshcdma8Type OTG_HS host channel-8 DMA address register
type RegisterOtghshcdma8Type uint32

func (r *RegisterOtghshcdma8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma8FieldDmaaddrShift = 0
	RegisterOtghshcdma8FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma8Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma8FieldDmaaddrMask) >> RegisterOtghshcdma8FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma8Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma8FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma8FieldDmaaddrShift))
}

// RegisterOtghshcchar9Type OTG_HS host channel-9 characteristics register
type RegisterOtghshcchar9Type uint32

func (r *RegisterOtghshcchar9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar9FieldMpsizShift = 0
	RegisterOtghshcchar9FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar9Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldMpsizMask) >> RegisterOtghshcchar9FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar9Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar9FieldMpsizShift))
}

const (
	RegisterOtghshcchar9FieldEpnumShift = 11
	RegisterOtghshcchar9FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar9Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldEpnumMask) >> RegisterOtghshcchar9FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar9Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar9FieldEpnumShift))
}

const (
	RegisterOtghshcchar9FieldEpdirShift = 15
	RegisterOtghshcchar9FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar9Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar9Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar9Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar9Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar9Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldEptypMask) >> RegisterOtghshcchar9FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar9Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar9FieldEptypShift))
}

const (
	RegisterOtghshcchar9FieldMcShift = 20
	RegisterOtghshcchar9FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar9Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldMcMask) >> RegisterOtghshcchar9FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar9Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldMcMask)|(uint32(value)<<RegisterOtghshcchar9FieldMcShift))
}

const (
	RegisterOtghshcchar9FieldDadShift = 22
	RegisterOtghshcchar9FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar9Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldDadMask) >> RegisterOtghshcchar9FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar9Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldDadMask)|(uint32(value)<<RegisterOtghshcchar9FieldDadShift))
}

const (
	RegisterOtghshcchar9FieldOddfrmShift = 29
	RegisterOtghshcchar9FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar9Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar9Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar9Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar9Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar9Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar9FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar9Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar9FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar9FieldChenaMask)
	}
}

// RegisterOtghshcsplt9Type OTG_HS host channel-9 split control register
type RegisterOtghshcsplt9Type uint32

func (r *RegisterOtghshcsplt9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt9FieldPrtaddrShift = 0
	RegisterOtghshcsplt9FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt9Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt9FieldPrtaddrMask) >> RegisterOtghshcsplt9FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt9Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt9FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt9FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt9FieldHubaddrShift = 7
	RegisterOtghshcsplt9FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt9Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt9FieldHubaddrMask) >> RegisterOtghshcsplt9FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt9Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt9FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt9FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt9FieldXactposShift = 14
	RegisterOtghshcsplt9FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt9Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt9FieldXactposMask) >> RegisterOtghshcsplt9FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt9Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt9FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt9FieldXactposShift))
}

const (
	RegisterOtghshcsplt9FieldComplspltShift = 16
	RegisterOtghshcsplt9FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt9Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt9FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt9Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt9Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt9FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt9Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt9FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt9FieldSplitenMask)
	}
}

// RegisterOtghshcint9Type OTG_HS host channel-9 interrupt register
type RegisterOtghshcint9Type uint32

func (r *RegisterOtghshcint9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint9FieldXfrcShift = 0
	RegisterOtghshcint9FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint9Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint9Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint9Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint9Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint9Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint9Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint9Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint9Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint9Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint9Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint9Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint9Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint9Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint9Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint9Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint9Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint9Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint9Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint9Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint9Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint9Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint9FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint9Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint9FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint9FieldDterrMask)
	}
}

// RegisterOtghshcintmsk9Type OTG_HS host channel-9 interrupt mask register
type RegisterOtghshcintmsk9Type uint32

func (r *RegisterOtghshcintmsk9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk9FieldXfrcmShift = 0
	RegisterOtghshcintmsk9FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk9Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk9Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk9Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk9Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk9Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk9Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk9Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk9Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk9Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk9Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk9Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk9Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk9Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk9Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk9Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk9Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk9Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk9Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk9Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk9Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk9Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk9FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk9Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk9FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk9FieldDterrmMask)
	}
}

// RegisterOtghshctsiz9Type OTG_HS host channel-9 transfer size register
type RegisterOtghshctsiz9Type uint32

func (r *RegisterOtghshctsiz9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz9FieldXfrsizShift = 0
	RegisterOtghshctsiz9FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz9Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz9FieldXfrsizMask) >> RegisterOtghshctsiz9FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz9Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz9FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz9FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz9FieldPktcntShift = 19
	RegisterOtghshctsiz9FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz9Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz9FieldPktcntMask) >> RegisterOtghshctsiz9FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz9Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz9FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz9FieldPktcntShift))
}

const (
	RegisterOtghshctsiz9FieldDpidShift = 29
	RegisterOtghshctsiz9FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz9Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz9FieldDpidMask) >> RegisterOtghshctsiz9FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz9Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz9FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz9FieldDpidShift))
}

// RegisterOtghshcdma9Type OTG_HS host channel-9 DMA address register
type RegisterOtghshcdma9Type uint32

func (r *RegisterOtghshcdma9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma9FieldDmaaddrShift = 0
	RegisterOtghshcdma9FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma9Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma9FieldDmaaddrMask) >> RegisterOtghshcdma9FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma9Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma9FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma9FieldDmaaddrShift))
}

// RegisterOtghshcchar10Type OTG_HS host channel-10 characteristics register
type RegisterOtghshcchar10Type uint32

func (r *RegisterOtghshcchar10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar10FieldMpsizShift = 0
	RegisterOtghshcchar10FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar10Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldMpsizMask) >> RegisterOtghshcchar10FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar10Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar10FieldMpsizShift))
}

const (
	RegisterOtghshcchar10FieldEpnumShift = 11
	RegisterOtghshcchar10FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar10Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldEpnumMask) >> RegisterOtghshcchar10FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar10Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar10FieldEpnumShift))
}

const (
	RegisterOtghshcchar10FieldEpdirShift = 15
	RegisterOtghshcchar10FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar10Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar10Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar10Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar10Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar10Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldEptypMask) >> RegisterOtghshcchar10FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar10Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar10FieldEptypShift))
}

const (
	RegisterOtghshcchar10FieldMcShift = 20
	RegisterOtghshcchar10FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar10Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldMcMask) >> RegisterOtghshcchar10FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar10Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldMcMask)|(uint32(value)<<RegisterOtghshcchar10FieldMcShift))
}

const (
	RegisterOtghshcchar10FieldDadShift = 22
	RegisterOtghshcchar10FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar10Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldDadMask) >> RegisterOtghshcchar10FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar10Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldDadMask)|(uint32(value)<<RegisterOtghshcchar10FieldDadShift))
}

const (
	RegisterOtghshcchar10FieldOddfrmShift = 29
	RegisterOtghshcchar10FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar10Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar10Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar10Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar10Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar10Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar10FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar10Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar10FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar10FieldChenaMask)
	}
}

// RegisterOtghshcsplt10Type OTG_HS host channel-10 split control register
type RegisterOtghshcsplt10Type uint32

func (r *RegisterOtghshcsplt10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt10FieldPrtaddrShift = 0
	RegisterOtghshcsplt10FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt10Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt10FieldPrtaddrMask) >> RegisterOtghshcsplt10FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt10Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt10FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt10FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt10FieldHubaddrShift = 7
	RegisterOtghshcsplt10FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt10Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt10FieldHubaddrMask) >> RegisterOtghshcsplt10FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt10Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt10FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt10FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt10FieldXactposShift = 14
	RegisterOtghshcsplt10FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt10Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt10FieldXactposMask) >> RegisterOtghshcsplt10FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt10Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt10FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt10FieldXactposShift))
}

const (
	RegisterOtghshcsplt10FieldComplspltShift = 16
	RegisterOtghshcsplt10FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt10Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt10FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt10Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt10Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt10FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt10Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt10FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt10FieldSplitenMask)
	}
}

// RegisterOtghshcint10Type OTG_HS host channel-10 interrupt register
type RegisterOtghshcint10Type uint32

func (r *RegisterOtghshcint10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint10FieldXfrcShift = 0
	RegisterOtghshcint10FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint10Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint10Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint10Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint10Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint10Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint10Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint10Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint10Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint10Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint10Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint10Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint10Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint10Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint10Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint10Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint10Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint10Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint10Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint10Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint10Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint10Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint10FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint10Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint10FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint10FieldDterrMask)
	}
}

// RegisterOtghshcintmsk10Type OTG_HS host channel-10 interrupt mask register
type RegisterOtghshcintmsk10Type uint32

func (r *RegisterOtghshcintmsk10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk10FieldXfrcmShift = 0
	RegisterOtghshcintmsk10FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk10Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk10Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk10Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk10Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk10Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk10Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk10Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk10Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk10Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk10Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk10Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk10Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk10Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk10Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk10Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk10Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk10Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk10Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk10Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk10Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk10Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk10FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk10Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk10FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk10FieldDterrmMask)
	}
}

// RegisterOtghshctsiz10Type OTG_HS host channel-10 transfer size register
type RegisterOtghshctsiz10Type uint32

func (r *RegisterOtghshctsiz10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz10FieldXfrsizShift = 0
	RegisterOtghshctsiz10FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz10Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz10FieldXfrsizMask) >> RegisterOtghshctsiz10FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz10Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz10FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz10FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz10FieldPktcntShift = 19
	RegisterOtghshctsiz10FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz10Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz10FieldPktcntMask) >> RegisterOtghshctsiz10FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz10Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz10FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz10FieldPktcntShift))
}

const (
	RegisterOtghshctsiz10FieldDpidShift = 29
	RegisterOtghshctsiz10FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz10Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz10FieldDpidMask) >> RegisterOtghshctsiz10FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz10Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz10FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz10FieldDpidShift))
}

// RegisterOtghshcdma10Type OTG_HS host channel-10 DMA address register
type RegisterOtghshcdma10Type uint32

func (r *RegisterOtghshcdma10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma10FieldDmaaddrShift = 0
	RegisterOtghshcdma10FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma10Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma10FieldDmaaddrMask) >> RegisterOtghshcdma10FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma10Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma10FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma10FieldDmaaddrShift))
}

// RegisterOtghshcchar11Type OTG_HS host channel-11 characteristics register
type RegisterOtghshcchar11Type uint32

func (r *RegisterOtghshcchar11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar11FieldMpsizShift = 0
	RegisterOtghshcchar11FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar11Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldMpsizMask) >> RegisterOtghshcchar11FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar11Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar11FieldMpsizShift))
}

const (
	RegisterOtghshcchar11FieldEpnumShift = 11
	RegisterOtghshcchar11FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar11Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldEpnumMask) >> RegisterOtghshcchar11FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar11Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar11FieldEpnumShift))
}

const (
	RegisterOtghshcchar11FieldEpdirShift = 15
	RegisterOtghshcchar11FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar11Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar11Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar11Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar11Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar11Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldEptypMask) >> RegisterOtghshcchar11FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar11Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar11FieldEptypShift))
}

const (
	RegisterOtghshcchar11FieldMcShift = 20
	RegisterOtghshcchar11FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar11Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldMcMask) >> RegisterOtghshcchar11FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar11Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldMcMask)|(uint32(value)<<RegisterOtghshcchar11FieldMcShift))
}

const (
	RegisterOtghshcchar11FieldDadShift = 22
	RegisterOtghshcchar11FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar11Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldDadMask) >> RegisterOtghshcchar11FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar11Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldDadMask)|(uint32(value)<<RegisterOtghshcchar11FieldDadShift))
}

const (
	RegisterOtghshcchar11FieldOddfrmShift = 29
	RegisterOtghshcchar11FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar11Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar11Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar11Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar11Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar11Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar11FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar11Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar11FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar11FieldChenaMask)
	}
}

// RegisterOtghshcsplt11Type OTG_HS host channel-11 split control register
type RegisterOtghshcsplt11Type uint32

func (r *RegisterOtghshcsplt11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt11FieldPrtaddrShift = 0
	RegisterOtghshcsplt11FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt11Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt11FieldPrtaddrMask) >> RegisterOtghshcsplt11FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt11Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt11FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt11FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt11FieldHubaddrShift = 7
	RegisterOtghshcsplt11FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt11Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt11FieldHubaddrMask) >> RegisterOtghshcsplt11FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt11Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt11FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt11FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt11FieldXactposShift = 14
	RegisterOtghshcsplt11FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt11Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt11FieldXactposMask) >> RegisterOtghshcsplt11FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt11Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt11FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt11FieldXactposShift))
}

const (
	RegisterOtghshcsplt11FieldComplspltShift = 16
	RegisterOtghshcsplt11FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt11Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt11FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt11Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt11Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt11FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt11Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt11FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt11FieldSplitenMask)
	}
}

// RegisterOtghshcint11Type OTG_HS host channel-11 interrupt register
type RegisterOtghshcint11Type uint32

func (r *RegisterOtghshcint11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint11FieldXfrcShift = 0
	RegisterOtghshcint11FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint11Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint11Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint11Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint11Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint11Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint11Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint11Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint11Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint11Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint11Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint11Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint11Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint11Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint11Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint11Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint11Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint11Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint11Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint11Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint11Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint11Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint11FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint11Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint11FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint11FieldDterrMask)
	}
}

// RegisterOtghshcintmsk11Type OTG_HS host channel-11 interrupt mask register
type RegisterOtghshcintmsk11Type uint32

func (r *RegisterOtghshcintmsk11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk11FieldXfrcmShift = 0
	RegisterOtghshcintmsk11FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk11Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk11Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk11Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk11Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk11Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk11Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk11Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk11Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk11Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk11Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk11Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk11Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk11Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldNyetMask) != 0
}

// SetNyet response received interrupt mask
func (r *RegisterOtghshcintmsk11Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk11Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error mask
func (r *RegisterOtghshcintmsk11Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk11Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldBberrmMask) != 0
}

// SetBberrm Babble error mask
func (r *RegisterOtghshcintmsk11Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk11Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk11Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk11Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk11FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk11Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk11FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk11FieldDterrmMask)
	}
}

// RegisterOtghshctsiz11Type OTG_HS host channel-11 transfer size register
type RegisterOtghshctsiz11Type uint32

func (r *RegisterOtghshctsiz11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz11FieldXfrsizShift = 0
	RegisterOtghshctsiz11FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz11Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz11FieldXfrsizMask) >> RegisterOtghshctsiz11FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz11Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz11FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz11FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz11FieldPktcntShift = 19
	RegisterOtghshctsiz11FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz11Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz11FieldPktcntMask) >> RegisterOtghshctsiz11FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz11Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz11FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz11FieldPktcntShift))
}

const (
	RegisterOtghshctsiz11FieldDpidShift = 29
	RegisterOtghshctsiz11FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz11Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz11FieldDpidMask) >> RegisterOtghshctsiz11FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz11Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz11FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz11FieldDpidShift))
}

// RegisterOtghshcdma11Type OTG_HS host channel-11 DMA address register
type RegisterOtghshcdma11Type uint32

func (r *RegisterOtghshcdma11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma11FieldDmaaddrShift = 0
	RegisterOtghshcdma11FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma11Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma11FieldDmaaddrMask) >> RegisterOtghshcdma11FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma11Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma11FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma11FieldDmaaddrShift))
}

// RegisterOtghshcchar12Type OTG_HS host channel-12 characteristics register
type RegisterOtghshcchar12Type uint32

func (r *RegisterOtghshcchar12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar12FieldMpsizShift = 0
	RegisterOtghshcchar12FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar12Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldMpsizMask) >> RegisterOtghshcchar12FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar12Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar12FieldMpsizShift))
}

const (
	RegisterOtghshcchar12FieldEpnumShift = 11
	RegisterOtghshcchar12FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar12Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldEpnumMask) >> RegisterOtghshcchar12FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar12Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar12FieldEpnumShift))
}

const (
	RegisterOtghshcchar12FieldEpdirShift = 15
	RegisterOtghshcchar12FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar12Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar12Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar12Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar12Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar12Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldEptypMask) >> RegisterOtghshcchar12FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar12Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar12FieldEptypShift))
}

const (
	RegisterOtghshcchar12FieldMcShift = 20
	RegisterOtghshcchar12FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar12Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldMcMask) >> RegisterOtghshcchar12FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar12Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldMcMask)|(uint32(value)<<RegisterOtghshcchar12FieldMcShift))
}

const (
	RegisterOtghshcchar12FieldDadShift = 22
	RegisterOtghshcchar12FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar12Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldDadMask) >> RegisterOtghshcchar12FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar12Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldDadMask)|(uint32(value)<<RegisterOtghshcchar12FieldDadShift))
}

const (
	RegisterOtghshcchar12FieldOddfrmShift = 29
	RegisterOtghshcchar12FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar12Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar12Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar12Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar12Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar12Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar12FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar12Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar12FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar12FieldChenaMask)
	}
}

// RegisterOtghshcsplt12Type OTG_HS host channel-12 split control register
type RegisterOtghshcsplt12Type uint32

func (r *RegisterOtghshcsplt12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt12FieldPrtaddrShift = 0
	RegisterOtghshcsplt12FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt12Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt12FieldPrtaddrMask) >> RegisterOtghshcsplt12FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt12Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt12FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt12FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt12FieldHubaddrShift = 7
	RegisterOtghshcsplt12FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt12Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt12FieldHubaddrMask) >> RegisterOtghshcsplt12FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt12Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt12FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt12FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt12FieldXactposShift = 14
	RegisterOtghshcsplt12FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt12Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt12FieldXactposMask) >> RegisterOtghshcsplt12FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt12Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt12FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt12FieldXactposShift))
}

const (
	RegisterOtghshcsplt12FieldComplspltShift = 16
	RegisterOtghshcsplt12FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt12Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt12FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt12Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt12Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt12FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt12Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt12FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt12FieldSplitenMask)
	}
}

// RegisterOtghshcint12Type OTG_HS host channel-12 interrupt register
type RegisterOtghshcint12Type uint32

func (r *RegisterOtghshcint12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint12FieldXfrcShift = 0
	RegisterOtghshcint12FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint12Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint12Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint12Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint12Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint12Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint12Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint12Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint12Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint12Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint12Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint12Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint12Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint12Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint12Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint12Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint12Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint12Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint12Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint12Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint12Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint12Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint12FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint12Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint12FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint12FieldDterrMask)
	}
}

// RegisterOtghshcintmsk12Type OTG_HS host channel-12 interrupt mask register
type RegisterOtghshcintmsk12Type uint32

func (r *RegisterOtghshcintmsk12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk12FieldXfrcmShift = 0
	RegisterOtghshcintmsk12FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk12Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk12Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk12Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk12Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk12Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk12Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk12Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk12Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk12Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk12Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk12Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk12Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk12Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcintmsk12Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk12Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *RegisterOtghshcintmsk12Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk12Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *RegisterOtghshcintmsk12Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk12Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk12Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk12Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk12FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk12Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk12FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk12FieldDterrmMask)
	}
}

// RegisterOtghshctsiz12Type OTG_HS host channel-12 transfer size register
type RegisterOtghshctsiz12Type uint32

func (r *RegisterOtghshctsiz12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz12FieldXfrsizShift = 0
	RegisterOtghshctsiz12FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz12Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz12FieldXfrsizMask) >> RegisterOtghshctsiz12FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz12Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz12FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz12FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz12FieldPktcntShift = 19
	RegisterOtghshctsiz12FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz12Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz12FieldPktcntMask) >> RegisterOtghshctsiz12FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz12Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz12FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz12FieldPktcntShift))
}

const (
	RegisterOtghshctsiz12FieldDpidShift = 29
	RegisterOtghshctsiz12FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz12Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz12FieldDpidMask) >> RegisterOtghshctsiz12FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz12Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz12FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz12FieldDpidShift))
}

// RegisterOtghshcdma12Type OTG_HS host channel-12 DMA address register
type RegisterOtghshcdma12Type uint32

func (r *RegisterOtghshcdma12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma12FieldDmaaddrShift = 0
	RegisterOtghshcdma12FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma12Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma12FieldDmaaddrMask) >> RegisterOtghshcdma12FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma12Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma12FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma12FieldDmaaddrShift))
}

// RegisterOtghshcchar13Type OTG_HS host channel-13 characteristics register
type RegisterOtghshcchar13Type uint32

func (r *RegisterOtghshcchar13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar13FieldMpsizShift = 0
	RegisterOtghshcchar13FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar13Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldMpsizMask) >> RegisterOtghshcchar13FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar13Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar13FieldMpsizShift))
}

const (
	RegisterOtghshcchar13FieldEpnumShift = 11
	RegisterOtghshcchar13FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar13Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldEpnumMask) >> RegisterOtghshcchar13FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar13Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar13FieldEpnumShift))
}

const (
	RegisterOtghshcchar13FieldEpdirShift = 15
	RegisterOtghshcchar13FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar13Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar13Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar13Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar13Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar13Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldEptypMask) >> RegisterOtghshcchar13FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar13Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar13FieldEptypShift))
}

const (
	RegisterOtghshcchar13FieldMcShift = 20
	RegisterOtghshcchar13FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar13Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldMcMask) >> RegisterOtghshcchar13FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar13Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldMcMask)|(uint32(value)<<RegisterOtghshcchar13FieldMcShift))
}

const (
	RegisterOtghshcchar13FieldDadShift = 22
	RegisterOtghshcchar13FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar13Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldDadMask) >> RegisterOtghshcchar13FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar13Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldDadMask)|(uint32(value)<<RegisterOtghshcchar13FieldDadShift))
}

const (
	RegisterOtghshcchar13FieldOddfrmShift = 29
	RegisterOtghshcchar13FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar13Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar13Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar13Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar13Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar13Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar13FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar13Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar13FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar13FieldChenaMask)
	}
}

// RegisterOtghshcsplt13Type OTG_HS host channel-13 split control register
type RegisterOtghshcsplt13Type uint32

func (r *RegisterOtghshcsplt13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt13FieldPrtaddrShift = 0
	RegisterOtghshcsplt13FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt13Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt13FieldPrtaddrMask) >> RegisterOtghshcsplt13FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt13Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt13FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt13FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt13FieldHubaddrShift = 7
	RegisterOtghshcsplt13FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt13Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt13FieldHubaddrMask) >> RegisterOtghshcsplt13FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt13Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt13FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt13FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt13FieldXactposShift = 14
	RegisterOtghshcsplt13FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt13Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt13FieldXactposMask) >> RegisterOtghshcsplt13FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt13Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt13FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt13FieldXactposShift))
}

const (
	RegisterOtghshcsplt13FieldComplspltShift = 16
	RegisterOtghshcsplt13FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt13Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt13FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt13Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt13Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt13FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt13Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt13FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt13FieldSplitenMask)
	}
}

// RegisterOtghshcint13Type OTG_HS host channel-13 interrupt register
type RegisterOtghshcint13Type uint32

func (r *RegisterOtghshcint13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint13FieldXfrcShift = 0
	RegisterOtghshcint13FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint13Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint13Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint13Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint13Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint13Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint13Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint13Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint13Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint13Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint13Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint13Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint13Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint13Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint13Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint13Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint13Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint13Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint13Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint13Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint13Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint13Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint13FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint13Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint13FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint13FieldDterrMask)
	}
}

// RegisterOtghshcintmsk13Type OTG_HS host channel-13 interrupt mask register
type RegisterOtghshcintmsk13Type uint32

func (r *RegisterOtghshcintmsk13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk13FieldXfrcmShift = 0
	RegisterOtghshcintmsk13FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk13Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk13Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk13Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk13Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk13Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk13Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk13Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldStallmMask) != 0
}

// SetStallm STALLM response received interrupt mask
func (r *RegisterOtghshcintmsk13Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk13Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk13Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk13Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk13Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk13Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcintmsk13Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk13Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *RegisterOtghshcintmsk13Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk13Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *RegisterOtghshcintmsk13Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk13Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk13Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk13Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk13FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk13Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk13FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk13FieldDterrmMask)
	}
}

// RegisterOtghshctsiz13Type OTG_HS host channel-13 transfer size register
type RegisterOtghshctsiz13Type uint32

func (r *RegisterOtghshctsiz13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz13FieldXfrsizShift = 0
	RegisterOtghshctsiz13FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz13Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz13FieldXfrsizMask) >> RegisterOtghshctsiz13FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz13Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz13FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz13FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz13FieldPktcntShift = 19
	RegisterOtghshctsiz13FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz13Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz13FieldPktcntMask) >> RegisterOtghshctsiz13FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz13Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz13FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz13FieldPktcntShift))
}

const (
	RegisterOtghshctsiz13FieldDpidShift = 29
	RegisterOtghshctsiz13FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz13Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz13FieldDpidMask) >> RegisterOtghshctsiz13FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz13Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz13FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz13FieldDpidShift))
}

// RegisterOtghshcdma13Type OTG_HS host channel-13 DMA address register
type RegisterOtghshcdma13Type uint32

func (r *RegisterOtghshcdma13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma13FieldDmaaddrShift = 0
	RegisterOtghshcdma13FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma13Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma13FieldDmaaddrMask) >> RegisterOtghshcdma13FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma13Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma13FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma13FieldDmaaddrShift))
}

// RegisterOtghshcchar14Type OTG_HS host channel-14 characteristics register
type RegisterOtghshcchar14Type uint32

func (r *RegisterOtghshcchar14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar14FieldMpsizShift = 0
	RegisterOtghshcchar14FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar14Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldMpsizMask) >> RegisterOtghshcchar14FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar14Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar14FieldMpsizShift))
}

const (
	RegisterOtghshcchar14FieldEpnumShift = 11
	RegisterOtghshcchar14FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar14Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldEpnumMask) >> RegisterOtghshcchar14FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar14Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar14FieldEpnumShift))
}

const (
	RegisterOtghshcchar14FieldEpdirShift = 15
	RegisterOtghshcchar14FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar14Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar14Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar14Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar14Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar14Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldEptypMask) >> RegisterOtghshcchar14FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar14Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar14FieldEptypShift))
}

const (
	RegisterOtghshcchar14FieldMcShift = 20
	RegisterOtghshcchar14FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar14Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldMcMask) >> RegisterOtghshcchar14FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar14Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldMcMask)|(uint32(value)<<RegisterOtghshcchar14FieldMcShift))
}

const (
	RegisterOtghshcchar14FieldDadShift = 22
	RegisterOtghshcchar14FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar14Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldDadMask) >> RegisterOtghshcchar14FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar14Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldDadMask)|(uint32(value)<<RegisterOtghshcchar14FieldDadShift))
}

const (
	RegisterOtghshcchar14FieldOddfrmShift = 29
	RegisterOtghshcchar14FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar14Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar14Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar14Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar14Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar14Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar14FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar14Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar14FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar14FieldChenaMask)
	}
}

// RegisterOtghshcsplt14Type OTG_HS host channel-14 split control register
type RegisterOtghshcsplt14Type uint32

func (r *RegisterOtghshcsplt14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt14FieldPrtaddrShift = 0
	RegisterOtghshcsplt14FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt14Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt14FieldPrtaddrMask) >> RegisterOtghshcsplt14FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt14Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt14FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt14FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt14FieldHubaddrShift = 7
	RegisterOtghshcsplt14FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt14Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt14FieldHubaddrMask) >> RegisterOtghshcsplt14FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt14Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt14FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt14FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt14FieldXactposShift = 14
	RegisterOtghshcsplt14FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt14Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt14FieldXactposMask) >> RegisterOtghshcsplt14FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt14Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt14FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt14FieldXactposShift))
}

const (
	RegisterOtghshcsplt14FieldComplspltShift = 16
	RegisterOtghshcsplt14FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt14Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt14FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt14Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt14Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt14FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt14Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt14FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt14FieldSplitenMask)
	}
}

// RegisterOtghshcint14Type OTG_HS host channel-14 interrupt register
type RegisterOtghshcint14Type uint32

func (r *RegisterOtghshcint14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint14FieldXfrcShift = 0
	RegisterOtghshcint14FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint14Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint14Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint14Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint14Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint14Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint14Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint14Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint14Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint14Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint14Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint14Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint14Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint14Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint14Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint14Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint14Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint14Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint14Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint14Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint14Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint14Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint14FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint14Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint14FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint14FieldDterrMask)
	}
}

// RegisterOtghshcintmsk14Type OTG_HS host channel-14 interrupt mask register
type RegisterOtghshcintmsk14Type uint32

func (r *RegisterOtghshcintmsk14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk14FieldXfrcmShift = 0
	RegisterOtghshcintmsk14FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk14Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk14Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk14Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk14Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk14Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk14Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk14Type) GetStallm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldStallmMask) != 0
}

// SetStallm STALL response received interrupt mask
func (r *RegisterOtghshcintmsk14Type) SetStallm(value bool) {
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
func (r *RegisterOtghshcintmsk14Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldNakmMask) != 0
}

// SetNakm NAKM response received interrupt mask
func (r *RegisterOtghshcintmsk14Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk14Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldAckmMask) != 0
}

// SetAckm ACKM response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk14Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk14Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcintmsk14Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk14Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *RegisterOtghshcintmsk14Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk14Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *RegisterOtghshcintmsk14Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk14Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk14Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk14Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk14FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk14Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk14FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk14FieldDterrmMask)
	}
}

// RegisterOtghshctsiz14Type OTG_HS host channel-14 transfer size register
type RegisterOtghshctsiz14Type uint32

func (r *RegisterOtghshctsiz14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz14FieldXfrsizShift = 0
	RegisterOtghshctsiz14FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz14Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz14FieldXfrsizMask) >> RegisterOtghshctsiz14FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz14Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz14FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz14FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz14FieldPktcntShift = 19
	RegisterOtghshctsiz14FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz14Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz14FieldPktcntMask) >> RegisterOtghshctsiz14FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz14Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz14FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz14FieldPktcntShift))
}

const (
	RegisterOtghshctsiz14FieldDpidShift = 29
	RegisterOtghshctsiz14FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz14Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz14FieldDpidMask) >> RegisterOtghshctsiz14FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz14Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz14FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz14FieldDpidShift))
}

// RegisterOtghshcdma14Type OTG_HS host channel-14 DMA address register
type RegisterOtghshcdma14Type uint32

func (r *RegisterOtghshcdma14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma14FieldDmaaddrShift = 0
	RegisterOtghshcdma14FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma14Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma14FieldDmaaddrMask) >> RegisterOtghshcdma14FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma14Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma14FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma14FieldDmaaddrShift))
}

// RegisterOtghshcchar15Type OTG_HS host channel-15 characteristics register
type RegisterOtghshcchar15Type uint32

func (r *RegisterOtghshcchar15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcchar15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcchar15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcchar15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcchar15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcchar15FieldMpsizShift = 0
	RegisterOtghshcchar15FieldMpsizMask  = 0x7ff
)

// GetMpsiz Maximum packet size
func (r *RegisterOtghshcchar15Type) GetMpsiz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldMpsizMask) >> RegisterOtghshcchar15FieldMpsizShift)
}

// SetMpsiz Maximum packet size
func (r *RegisterOtghshcchar15Type) SetMpsiz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldMpsizMask)|(uint32(value)<<RegisterOtghshcchar15FieldMpsizShift))
}

const (
	RegisterOtghshcchar15FieldEpnumShift = 11
	RegisterOtghshcchar15FieldEpnumMask  = 0x7800
)

// GetEpnum Endpoint number
func (r *RegisterOtghshcchar15Type) GetEpnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldEpnumMask) >> RegisterOtghshcchar15FieldEpnumShift)
}

// SetEpnum Endpoint number
func (r *RegisterOtghshcchar15Type) SetEpnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldEpnumMask)|(uint32(value)<<RegisterOtghshcchar15FieldEpnumShift))
}

const (
	RegisterOtghshcchar15FieldEpdirShift = 15
	RegisterOtghshcchar15FieldEpdirMask  = 0x8000
)

// GetEpdir Endpoint direction
func (r *RegisterOtghshcchar15Type) GetEpdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldEpdirMask) != 0
}

// SetEpdir Endpoint direction
func (r *RegisterOtghshcchar15Type) SetEpdir(value bool) {
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
func (r *RegisterOtghshcchar15Type) GetLsdev() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldLsdevMask) != 0
}

// SetLsdev Low-speed device
func (r *RegisterOtghshcchar15Type) SetLsdev(value bool) {
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
func (r *RegisterOtghshcchar15Type) GetEptyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldEptypMask) >> RegisterOtghshcchar15FieldEptypShift)
}

// SetEptyp Endpoint type
func (r *RegisterOtghshcchar15Type) SetEptyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldEptypMask)|(uint32(value)<<RegisterOtghshcchar15FieldEptypShift))
}

const (
	RegisterOtghshcchar15FieldMcShift = 20
	RegisterOtghshcchar15FieldMcMask  = 0x300000
)

// GetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar15Type) GetMc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldMcMask) >> RegisterOtghshcchar15FieldMcShift)
}

// SetMc Multi Count (MC) / Error Count (EC)
func (r *RegisterOtghshcchar15Type) SetMc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldMcMask)|(uint32(value)<<RegisterOtghshcchar15FieldMcShift))
}

const (
	RegisterOtghshcchar15FieldDadShift = 22
	RegisterOtghshcchar15FieldDadMask  = 0x1fc00000
)

// GetDad Device address
func (r *RegisterOtghshcchar15Type) GetDad() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldDadMask) >> RegisterOtghshcchar15FieldDadShift)
}

// SetDad Device address
func (r *RegisterOtghshcchar15Type) SetDad(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldDadMask)|(uint32(value)<<RegisterOtghshcchar15FieldDadShift))
}

const (
	RegisterOtghshcchar15FieldOddfrmShift = 29
	RegisterOtghshcchar15FieldOddfrmMask  = 0x20000000
)

// GetOddfrm Odd frame
func (r *RegisterOtghshcchar15Type) GetOddfrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldOddfrmMask) != 0
}

// SetOddfrm Odd frame
func (r *RegisterOtghshcchar15Type) SetOddfrm(value bool) {
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
func (r *RegisterOtghshcchar15Type) GetChdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldChdisMask) != 0
}

// SetChdis Channel disable
func (r *RegisterOtghshcchar15Type) SetChdis(value bool) {
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
func (r *RegisterOtghshcchar15Type) GetChena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcchar15FieldChenaMask) != 0
}

// SetChena Channel enable
func (r *RegisterOtghshcchar15Type) SetChena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcchar15FieldChenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcchar15FieldChenaMask)
	}
}

// RegisterOtghshcsplt15Type OTG_HS host channel-15 split control register
type RegisterOtghshcsplt15Type uint32

func (r *RegisterOtghshcsplt15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcsplt15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcsplt15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcsplt15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcsplt15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcsplt15FieldPrtaddrShift = 0
	RegisterOtghshcsplt15FieldPrtaddrMask  = 0x7f
)

// GetPrtaddr Port address
func (r *RegisterOtghshcsplt15Type) GetPrtaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt15FieldPrtaddrMask) >> RegisterOtghshcsplt15FieldPrtaddrShift)
}

// SetPrtaddr Port address
func (r *RegisterOtghshcsplt15Type) SetPrtaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt15FieldPrtaddrMask)|(uint32(value)<<RegisterOtghshcsplt15FieldPrtaddrShift))
}

const (
	RegisterOtghshcsplt15FieldHubaddrShift = 7
	RegisterOtghshcsplt15FieldHubaddrMask  = 0x3f80
)

// GetHubaddr Hub address
func (r *RegisterOtghshcsplt15Type) GetHubaddr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt15FieldHubaddrMask) >> RegisterOtghshcsplt15FieldHubaddrShift)
}

// SetHubaddr Hub address
func (r *RegisterOtghshcsplt15Type) SetHubaddr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt15FieldHubaddrMask)|(uint32(value)<<RegisterOtghshcsplt15FieldHubaddrShift))
}

const (
	RegisterOtghshcsplt15FieldXactposShift = 14
	RegisterOtghshcsplt15FieldXactposMask  = 0xc000
)

// GetXactpos XACTPOS
func (r *RegisterOtghshcsplt15Type) GetXactpos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt15FieldXactposMask) >> RegisterOtghshcsplt15FieldXactposShift)
}

// SetXactpos XACTPOS
func (r *RegisterOtghshcsplt15Type) SetXactpos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt15FieldXactposMask)|(uint32(value)<<RegisterOtghshcsplt15FieldXactposShift))
}

const (
	RegisterOtghshcsplt15FieldComplspltShift = 16
	RegisterOtghshcsplt15FieldComplspltMask  = 0x10000
)

// GetComplsplt Do complete split
func (r *RegisterOtghshcsplt15Type) GetComplsplt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt15FieldComplspltMask) != 0
}

// SetComplsplt Do complete split
func (r *RegisterOtghshcsplt15Type) SetComplsplt(value bool) {
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
func (r *RegisterOtghshcsplt15Type) GetSpliten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcsplt15FieldSplitenMask) != 0
}

// SetSpliten Split enable
func (r *RegisterOtghshcsplt15Type) SetSpliten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcsplt15FieldSplitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcsplt15FieldSplitenMask)
	}
}

// RegisterOtghshcint15Type OTG_HS host channel-15 interrupt register
type RegisterOtghshcint15Type uint32

func (r *RegisterOtghshcint15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcint15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcint15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcint15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcint15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcint15FieldXfrcShift = 0
	RegisterOtghshcint15FieldXfrcMask  = 0x1
)

// GetXfrc Transfer completed
func (r *RegisterOtghshcint15Type) GetXfrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldXfrcMask) != 0
}

// SetXfrc Transfer completed
func (r *RegisterOtghshcint15Type) SetXfrc(value bool) {
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
func (r *RegisterOtghshcint15Type) GetChh() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldChhMask) != 0
}

// SetChh Channel halted
func (r *RegisterOtghshcint15Type) SetChh(value bool) {
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
func (r *RegisterOtghshcint15Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcint15Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcint15Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldStallMask) != 0
}

// SetStall STALL response received interrupt
func (r *RegisterOtghshcint15Type) SetStall(value bool) {
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
func (r *RegisterOtghshcint15Type) GetNak() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldNakMask) != 0
}

// SetNak NAK response received interrupt
func (r *RegisterOtghshcint15Type) SetNak(value bool) {
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
func (r *RegisterOtghshcint15Type) GetAck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldAckMask) != 0
}

// SetAck ACK response received/transmitted interrupt
func (r *RegisterOtghshcint15Type) SetAck(value bool) {
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
func (r *RegisterOtghshcint15Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcint15Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcint15Type) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldTxerrMask) != 0
}

// SetTxerr Transaction error
func (r *RegisterOtghshcint15Type) SetTxerr(value bool) {
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
func (r *RegisterOtghshcint15Type) GetBberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldBberrMask) != 0
}

// SetBberr Babble error
func (r *RegisterOtghshcint15Type) SetBberr(value bool) {
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
func (r *RegisterOtghshcint15Type) GetFrmor() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldFrmorMask) != 0
}

// SetFrmor Frame overrun
func (r *RegisterOtghshcint15Type) SetFrmor(value bool) {
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
func (r *RegisterOtghshcint15Type) GetDterr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcint15FieldDterrMask) != 0
}

// SetDterr Data toggle error
func (r *RegisterOtghshcint15Type) SetDterr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcint15FieldDterrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcint15FieldDterrMask)
	}
}

// RegisterOtghshcintmsk15Type OTG_HS host channel-15 interrupt mask register
type RegisterOtghshcintmsk15Type uint32

func (r *RegisterOtghshcintmsk15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcintmsk15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcintmsk15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcintmsk15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcintmsk15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcintmsk15FieldXfrcmShift = 0
	RegisterOtghshcintmsk15FieldXfrcmMask  = 0x1
)

// GetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk15Type) GetXfrcm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldXfrcmMask) != 0
}

// SetXfrcm Transfer completed mask
func (r *RegisterOtghshcintmsk15Type) SetXfrcm(value bool) {
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
func (r *RegisterOtghshcintmsk15Type) GetChhm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldChhmMask) != 0
}

// SetChhm Channel halted mask
func (r *RegisterOtghshcintmsk15Type) SetChhm(value bool) {
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
func (r *RegisterOtghshcintmsk15Type) GetAhberr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldAhberrMask) != 0
}

// SetAhberr AHB error
func (r *RegisterOtghshcintmsk15Type) SetAhberr(value bool) {
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
func (r *RegisterOtghshcintmsk15Type) GetStall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldStallMask) != 0
}

// SetStall STALL response received interrupt mask
func (r *RegisterOtghshcintmsk15Type) SetStall(value bool) {
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
func (r *RegisterOtghshcintmsk15Type) GetNakm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldNakmMask) != 0
}

// SetNakm NAK response received interrupt mask
func (r *RegisterOtghshcintmsk15Type) SetNakm(value bool) {
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
func (r *RegisterOtghshcintmsk15Type) GetAckm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldAckmMask) != 0
}

// SetAckm ACK response received/transmitted interrupt mask
func (r *RegisterOtghshcintmsk15Type) SetAckm(value bool) {
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
func (r *RegisterOtghshcintmsk15Type) GetNyet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldNyetMask) != 0
}

// SetNyet Response received interrupt
func (r *RegisterOtghshcintmsk15Type) SetNyet(value bool) {
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
func (r *RegisterOtghshcintmsk15Type) GetTxerrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldTxerrmMask) != 0
}

// SetTxerrm Transaction error
func (r *RegisterOtghshcintmsk15Type) SetTxerrm(value bool) {
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
func (r *RegisterOtghshcintmsk15Type) GetBberrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldBberrmMask) != 0
}

// SetBberrm Babble error
func (r *RegisterOtghshcintmsk15Type) SetBberrm(value bool) {
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
func (r *RegisterOtghshcintmsk15Type) GetFrmorm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldFrmormMask) != 0
}

// SetFrmorm Frame overrun mask
func (r *RegisterOtghshcintmsk15Type) SetFrmorm(value bool) {
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
func (r *RegisterOtghshcintmsk15Type) GetDterrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcintmsk15FieldDterrmMask) != 0
}

// SetDterrm Data toggle error mask
func (r *RegisterOtghshcintmsk15Type) SetDterrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOtghshcintmsk15FieldDterrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcintmsk15FieldDterrmMask)
	}
}

// RegisterOtghshctsiz15Type OTG_HS host channel-15 transfer size register
type RegisterOtghshctsiz15Type uint32

func (r *RegisterOtghshctsiz15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshctsiz15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshctsiz15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshctsiz15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshctsiz15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshctsiz15FieldXfrsizShift = 0
	RegisterOtghshctsiz15FieldXfrsizMask  = 0x7ffff
)

// GetXfrsiz Transfer size
func (r *RegisterOtghshctsiz15Type) GetXfrsiz() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz15FieldXfrsizMask) >> RegisterOtghshctsiz15FieldXfrsizShift)
}

// SetXfrsiz Transfer size
func (r *RegisterOtghshctsiz15Type) SetXfrsiz(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz15FieldXfrsizMask)|(uint32(value)<<RegisterOtghshctsiz15FieldXfrsizShift))
}

const (
	RegisterOtghshctsiz15FieldPktcntShift = 19
	RegisterOtghshctsiz15FieldPktcntMask  = 0x1ff80000
)

// GetPktcnt Packet count
func (r *RegisterOtghshctsiz15Type) GetPktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz15FieldPktcntMask) >> RegisterOtghshctsiz15FieldPktcntShift)
}

// SetPktcnt Packet count
func (r *RegisterOtghshctsiz15Type) SetPktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz15FieldPktcntMask)|(uint32(value)<<RegisterOtghshctsiz15FieldPktcntShift))
}

const (
	RegisterOtghshctsiz15FieldDpidShift = 29
	RegisterOtghshctsiz15FieldDpidMask  = 0x60000000
)

// GetDpid Data PID
func (r *RegisterOtghshctsiz15Type) GetDpid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshctsiz15FieldDpidMask) >> RegisterOtghshctsiz15FieldDpidShift)
}

// SetDpid Data PID
func (r *RegisterOtghshctsiz15Type) SetDpid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshctsiz15FieldDpidMask)|(uint32(value)<<RegisterOtghshctsiz15FieldDpidShift))
}

// RegisterOtghshcdma15Type OTG_HS host channel-15 DMA address register
type RegisterOtghshcdma15Type uint32

func (r *RegisterOtghshcdma15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterOtghshcdma15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterOtghshcdma15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterOtghshcdma15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterOtghshcdma15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterOtghshcdma15FieldDmaaddrShift = 0
	RegisterOtghshcdma15FieldDmaaddrMask  = 0xffffffff
)

// GetDmaaddr DMA address
func (r *RegisterOtghshcdma15Type) GetDmaaddr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOtghshcdma15FieldDmaaddrMask) >> RegisterOtghshcdma15FieldDmaaddrShift)
}

// SetDmaaddr DMA address
func (r *RegisterOtghshcdma15Type) SetDmaaddr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOtghshcdma15FieldDmaaddrMask)|(uint32(value)<<RegisterOtghshcdma15FieldDmaaddrShift))
}
