//go:build arm && cortexm

package tcm

import (
	"unsafe"
	"volatile"
)

var (
	Tcm = (*_tcm)(unsafe.Pointer(uintptr(0xe001e500)))
)

type _tcm struct {
	Itguctrl RegisterItguctrlType
	Itgucfg  RegisterItgucfgType
	_        [8]byte
	Itgulut  [16]RegisterItgulutType
	_        [176]byte
	Dtguctrl RegisterDtguctrlType
	Dtgucfg  RegisterDtgucfgType
	_        [8]byte
	Dtgulut  [16]RegisterDtgulutType
}

// RegisterItguctrlType Main TCM Gate Unit (TGU) control registers for the ITCM
type RegisterItguctrlType uint32

func (r *RegisterItguctrlType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterItguctrlType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterItguctrlType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterItguctrlType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterItguctrlType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterItguctrlFieldDbfenShift = 0
	RegisterItguctrlFieldDbfenMask  = 0x1
)

// GetDbfen Enable data side BusFault for TGU fault
func (r *RegisterItguctrlType) GetDbfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterItguctrlFieldDbfenMask) != 0
}

// SetDbfen Enable data side BusFault for TGU fault
func (r *RegisterItguctrlType) SetDbfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterItguctrlFieldDbfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterItguctrlFieldDbfenMask)
	}
}

const (
	RegisterItguctrlFieldDerenShift = 1
	RegisterItguctrlFieldDerenMask  = 0x2
)

// GetDeren Enable Slave AHB (S-AHB) error response for TGU fault
func (r *RegisterItguctrlType) GetDeren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterItguctrlFieldDerenMask) != 0
}

// SetDeren Enable Slave AHB (S-AHB) error response for TGU fault
func (r *RegisterItguctrlType) SetDeren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterItguctrlFieldDerenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterItguctrlFieldDerenMask)
	}
}

// RegisterItgucfgType Provides access to configuration values for the ITGU
type RegisterItgucfgType uint32

func (r *RegisterItgucfgType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterItgucfgType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterItgucfgType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterItgucfgType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterItgucfgType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterItgucfgFieldBlkszShift = 0
	RegisterItgucfgFieldBlkszMask  = 0xf
)

// GetBlksz TGU block size in bytes
func (r *RegisterItgucfgType) GetBlksz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterItgucfgFieldBlkszMask) >> RegisterItgucfgFieldBlkszShift)
}

// SetBlksz TGU block size in bytes
func (r *RegisterItgucfgType) SetBlksz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterItgucfgFieldBlkszMask)|(uint32(value)<<RegisterItgucfgFieldBlkszShift))
}

const (
	RegisterItgucfgFieldNumblksShift = 8
	RegisterItgucfgFieldNumblksMask  = 0xf00
)

// GetNumblks This value is used to calculate the number of TCM blocks
func (r *RegisterItgucfgType) GetNumblks() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterItgucfgFieldNumblksMask) >> RegisterItgucfgFieldNumblksShift)
}

// SetNumblks This value is used to calculate the number of TCM blocks
func (r *RegisterItgucfgType) SetNumblks(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterItgucfgFieldNumblksMask)|(uint32(value)<<RegisterItgucfgFieldNumblksShift))
}

const (
	RegisterItgucfgFieldPresentShift = 31
	RegisterItgucfgFieldPresentMask  = 0x80000000
)

// GetPresent Determines if the TGU is present
func (r *RegisterItgucfgType) GetPresent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterItgucfgFieldPresentMask) != 0
}

// SetPresent Determines if the TGU is present
func (r *RegisterItgucfgType) SetPresent(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterItgucfgFieldPresentMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterItgucfgFieldPresentMask)
	}
}

// RegisterItgulutType Identifying the ITGU blocks as being Secure or Non-secure
type RegisterItgulutType uint32

func (r *RegisterItgulutType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterItgulutType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterItgulutType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterItgulutType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterItgulutType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

// RegisterDtguctrlType Main TCM Gate Unit (TGU) control registers for the DTCM
type RegisterDtguctrlType uint32

func (r *RegisterDtguctrlType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDtguctrlType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDtguctrlType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDtguctrlType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDtguctrlType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDtguctrlFieldDbfenShift = 0
	RegisterDtguctrlFieldDbfenMask  = 0x1
)

// GetDbfen Enable data side BusFault for TGU fault
func (r *RegisterDtguctrlType) GetDbfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtguctrlFieldDbfenMask) != 0
}

// SetDbfen Enable data side BusFault for TGU fault
func (r *RegisterDtguctrlType) SetDbfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtguctrlFieldDbfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtguctrlFieldDbfenMask)
	}
}

const (
	RegisterDtguctrlFieldDerenShift = 1
	RegisterDtguctrlFieldDerenMask  = 0x2
)

// GetDeren Enable Slave AHB (S-AHB) error response for TGU fault
func (r *RegisterDtguctrlType) GetDeren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtguctrlFieldDerenMask) != 0
}

// SetDeren Enable Slave AHB (S-AHB) error response for TGU fault
func (r *RegisterDtguctrlType) SetDeren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtguctrlFieldDerenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtguctrlFieldDerenMask)
	}
}

// RegisterDtgucfgType Provides access to configuration values for the DTGU
type RegisterDtgucfgType uint32

func (r *RegisterDtgucfgType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDtgucfgType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDtgucfgType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDtgucfgType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDtgucfgType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDtgucfgFieldBlkszShift = 0
	RegisterDtgucfgFieldBlkszMask  = 0xf
)

// GetBlksz TGU block size in bytes
func (r *RegisterDtgucfgType) GetBlksz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtgucfgFieldBlkszMask) >> RegisterDtgucfgFieldBlkszShift)
}

// SetBlksz TGU block size in bytes
func (r *RegisterDtgucfgType) SetBlksz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtgucfgFieldBlkszMask)|(uint32(value)<<RegisterDtgucfgFieldBlkszShift))
}

const (
	RegisterDtgucfgFieldNumblksShift = 8
	RegisterDtgucfgFieldNumblksMask  = 0xf00
)

// GetNumblks This value is used to calculate the number of TCM blocks
func (r *RegisterDtgucfgType) GetNumblks() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtgucfgFieldNumblksMask) >> RegisterDtgucfgFieldNumblksShift)
}

// SetNumblks This value is used to calculate the number of TCM blocks
func (r *RegisterDtgucfgType) SetNumblks(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtgucfgFieldNumblksMask)|(uint32(value)<<RegisterDtgucfgFieldNumblksShift))
}

const (
	RegisterDtgucfgFieldPresentShift = 31
	RegisterDtgucfgFieldPresentMask  = 0x80000000
)

// GetPresent Determines if the TGU is present
func (r *RegisterDtgucfgType) GetPresent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtgucfgFieldPresentMask) != 0
}

// SetPresent Determines if the TGU is present
func (r *RegisterDtgucfgType) SetPresent(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtgucfgFieldPresentMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtgucfgFieldPresentMask)
	}
}

// RegisterDtgulutType Identifying the DTGU blocks as being Secure or Non-secure
type RegisterDtgulutType uint32

func (r *RegisterDtgulutType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDtgulutType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDtgulutType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDtgulutType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDtgulutType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}
