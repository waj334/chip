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
	Itguctrl registerItguctrlType
	Itgucfg  registerItgucfgType
	_        [8]byte
	Itgulut  [16]registerItgulutType
	_        [176]byte
	Dtguctrl registerDtguctrlType
	Dtgucfg  registerDtgucfgType
	_        [8]byte
	Dtgulut  [16]registerDtgulutType
}

// registerItguctrlType Main TCM Gate Unit (TGU) control registers for the ITCM
type registerItguctrlType uint32

const (
	RegisterItguctrlFieldDbfenShift = 0
	RegisterItguctrlFieldDbfenMask  = 0x1
)

// GetDbfen Enable data side BusFault for TGU fault
func (r *registerItguctrlType) GetDbfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterItguctrlFieldDbfenMask) != 0
}

// SetDbfen Enable data side BusFault for TGU fault
func (r *registerItguctrlType) SetDbfen(value bool) {
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
func (r *registerItguctrlType) GetDeren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterItguctrlFieldDerenMask) != 0
}

// SetDeren Enable Slave AHB (S-AHB) error response for TGU fault
func (r *registerItguctrlType) SetDeren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterItguctrlFieldDerenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterItguctrlFieldDerenMask)
	}
}

// registerItgucfgType Provides access to configuration values for the ITGU
type registerItgucfgType uint32

const (
	RegisterItgucfgFieldBlkszShift = 0
	RegisterItgucfgFieldBlkszMask  = 0xf
)

// GetBlksz TGU block size in bytes
func (r *registerItgucfgType) GetBlksz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterItgucfgFieldBlkszMask) >> RegisterItgucfgFieldBlkszShift)
}

// SetBlksz TGU block size in bytes
func (r *registerItgucfgType) SetBlksz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterItgucfgFieldBlkszMask)|(uint32(value)<<RegisterItgucfgFieldBlkszShift))
}

const (
	RegisterItgucfgFieldNumblksShift = 8
	RegisterItgucfgFieldNumblksMask  = 0xf00
)

// GetNumblks This value is used to calculate the number of TCM blocks
func (r *registerItgucfgType) GetNumblks() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterItgucfgFieldNumblksMask) >> RegisterItgucfgFieldNumblksShift)
}

// SetNumblks This value is used to calculate the number of TCM blocks
func (r *registerItgucfgType) SetNumblks(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterItgucfgFieldNumblksMask)|(uint32(value)<<RegisterItgucfgFieldNumblksShift))
}

const (
	RegisterItgucfgFieldPresentShift = 31
	RegisterItgucfgFieldPresentMask  = 0x80000000
)

// GetPresent Determines if the TGU is present
func (r *registerItgucfgType) GetPresent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterItgucfgFieldPresentMask) != 0
}

// SetPresent Determines if the TGU is present
func (r *registerItgucfgType) SetPresent(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterItgucfgFieldPresentMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterItgucfgFieldPresentMask)
	}
}

// registerItgulutType Identifying the ITGU blocks as being Secure or Non-secure
type registerItgulutType uint32

// registerDtguctrlType Main TCM Gate Unit (TGU) control registers for the DTCM
type registerDtguctrlType uint32

const (
	RegisterDtguctrlFieldDbfenShift = 0
	RegisterDtguctrlFieldDbfenMask  = 0x1
)

// GetDbfen Enable data side BusFault for TGU fault
func (r *registerDtguctrlType) GetDbfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtguctrlFieldDbfenMask) != 0
}

// SetDbfen Enable data side BusFault for TGU fault
func (r *registerDtguctrlType) SetDbfen(value bool) {
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
func (r *registerDtguctrlType) GetDeren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtguctrlFieldDerenMask) != 0
}

// SetDeren Enable Slave AHB (S-AHB) error response for TGU fault
func (r *registerDtguctrlType) SetDeren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtguctrlFieldDerenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtguctrlFieldDerenMask)
	}
}

// registerDtgucfgType Provides access to configuration values for the DTGU
type registerDtgucfgType uint32

const (
	RegisterDtgucfgFieldBlkszShift = 0
	RegisterDtgucfgFieldBlkszMask  = 0xf
)

// GetBlksz TGU block size in bytes
func (r *registerDtgucfgType) GetBlksz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtgucfgFieldBlkszMask) >> RegisterDtgucfgFieldBlkszShift)
}

// SetBlksz TGU block size in bytes
func (r *registerDtgucfgType) SetBlksz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtgucfgFieldBlkszMask)|(uint32(value)<<RegisterDtgucfgFieldBlkszShift))
}

const (
	RegisterDtgucfgFieldNumblksShift = 8
	RegisterDtgucfgFieldNumblksMask  = 0xf00
)

// GetNumblks This value is used to calculate the number of TCM blocks
func (r *registerDtgucfgType) GetNumblks() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDtgucfgFieldNumblksMask) >> RegisterDtgucfgFieldNumblksShift)
}

// SetNumblks This value is used to calculate the number of TCM blocks
func (r *registerDtgucfgType) SetNumblks(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtgucfgFieldNumblksMask)|(uint32(value)<<RegisterDtgucfgFieldNumblksShift))
}

const (
	RegisterDtgucfgFieldPresentShift = 31
	RegisterDtgucfgFieldPresentMask  = 0x80000000
)

// GetPresent Determines if the TGU is present
func (r *registerDtgucfgType) GetPresent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDtgucfgFieldPresentMask) != 0
}

// SetPresent Determines if the TGU is present
func (r *registerDtgucfgType) SetPresent(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDtgucfgFieldPresentMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDtgucfgFieldPresentMask)
	}
}

// registerDtgulutType Identifying the DTGU blocks as being Secure or Non-secure
type registerDtgulutType uint32
