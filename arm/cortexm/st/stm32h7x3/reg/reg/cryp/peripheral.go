package cryp

import (
	"unsafe"
	"volatile"
)

var (
	Cryp = (*_cryp)(unsafe.Pointer(uintptr(0x48021000)))
)

type _cryp struct {
	Cr         registerCrType
	Sr         registerSrType
	Din        registerDinType
	Dout       registerDoutType
	Dmacr      registerDmacrType
	Imscr      registerImscrType
	Risr       registerRisrType
	Misr       registerMisrType
	K0lr       registerK0lrType
	K0rr       registerK0rrType
	K1lr       registerK1lrType
	K1rr       registerK1rrType
	K2lr       registerK2lrType
	K2rr       registerK2rrType
	K3lr       registerK3lrType
	K3rr       registerK3rrType
	Iv0lr      registerIv0lrType
	Iv0rr      registerIv0rrType
	Iv1lr      registerIv1lrType
	Iv1rr      registerIv1rrType
	Csgcmccm0r registerCsgcmccm0rType
	Csgcmccm1r registerCsgcmccm1rType
	Csgcmccm2r registerCsgcmccm2rType
	Csgcmccm3r registerCsgcmccm3rType
	Csgcmccm4r registerCsgcmccm4rType
	Csgcmccm5r registerCsgcmccm5rType
	Csgcmccm6r registerCsgcmccm6rType
	Csgcmccm7r registerCsgcmccm7rType
	Csgcm0r    registerCsgcm0rType
	Csgcm1r    registerCsgcm1rType
	Csgcm2r    registerCsgcm2rType
	Csgcm3r    registerCsgcm3rType
	Csgcm4r    registerCsgcm4rType
	Csgcm5r    registerCsgcm5rType
	Csgcm6r    registerCsgcm6rType
	Csgcm7r    registerCsgcm7rType
}

// registerCrType control register
type registerCrType uint32

const (
	RegisterCrFieldAlgodirShift = 2
	RegisterCrFieldAlgodirMask  = 0x4
)

// GetAlgodir Algorithm direction
func (r *registerCrType) GetAlgodir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAlgodirMask) != 0
}

// SetAlgodir Algorithm direction
func (r *registerCrType) SetAlgodir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAlgodirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAlgodirMask)
	}
}

const (
	RegisterCrFieldAlgomode0Shift = 3
	RegisterCrFieldAlgomode0Mask  = 0x38
)

// GetAlgomode0 Algorithm mode
func (r *registerCrType) GetAlgomode0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAlgomode0Mask) >> RegisterCrFieldAlgomode0Shift)
}

// SetAlgomode0 Algorithm mode
func (r *registerCrType) SetAlgomode0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAlgomode0Mask)|(uint32(value)<<RegisterCrFieldAlgomode0Shift))
}

const (
	RegisterCrFieldDatatypeShift = 6
	RegisterCrFieldDatatypeMask  = 0xc0
)

// GetDatatype Data type selection
func (r *registerCrType) GetDatatype() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDatatypeMask) >> RegisterCrFieldDatatypeShift)
}

// SetDatatype Data type selection
func (r *registerCrType) SetDatatype(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDatatypeMask)|(uint32(value)<<RegisterCrFieldDatatypeShift))
}

const (
	RegisterCrFieldKeysizeShift = 8
	RegisterCrFieldKeysizeMask  = 0x300
)

// GetKeysize Key size selection (AES mode only)
func (r *registerCrType) GetKeysize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldKeysizeMask) >> RegisterCrFieldKeysizeShift)
}

// SetKeysize Key size selection (AES mode only)
func (r *registerCrType) SetKeysize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldKeysizeMask)|(uint32(value)<<RegisterCrFieldKeysizeShift))
}

const (
	RegisterCrFieldFflushShift = 14
	RegisterCrFieldFflushMask  = 0x4000
)

// SetFflush FIFO flush
func (r *registerCrType) SetFflush(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldFflushMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldFflushMask)
	}
}

const (
	RegisterCrFieldCrypenShift = 15
	RegisterCrFieldCrypenMask  = 0x8000
)

// GetCrypen Cryptographic processor enable
func (r *registerCrType) GetCrypen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCrypenMask) != 0
}

// SetCrypen Cryptographic processor enable
func (r *registerCrType) SetCrypen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCrypenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCrypenMask)
	}
}

const (
	RegisterCrFieldGcm_ccmphShift = 16
	RegisterCrFieldGcm_ccmphMask  = 0x30000
)

// GetGcm_ccmph GCM_CCMPH
func (r *registerCrType) GetGcm_ccmph() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldGcm_ccmphMask) >> RegisterCrFieldGcm_ccmphShift)
}

// SetGcm_ccmph GCM_CCMPH
func (r *registerCrType) SetGcm_ccmph(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldGcm_ccmphMask)|(uint32(value)<<RegisterCrFieldGcm_ccmphShift))
}

const (
	RegisterCrFieldAlgomode3Shift = 19
	RegisterCrFieldAlgomode3Mask  = 0x80000
)

// GetAlgomode3 ALGOMODE
func (r *registerCrType) GetAlgomode3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAlgomode3Mask) != 0
}

// SetAlgomode3 ALGOMODE
func (r *registerCrType) SetAlgomode3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAlgomode3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAlgomode3Mask)
	}
}

// registerSrType status register
type registerSrType uint32

const (
	RegisterSrFieldBusyShift = 4
	RegisterSrFieldBusyMask  = 0x10
)

// GetBusy Busy bit
func (r *registerSrType) GetBusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldBusyMask) != 0
}

// SetBusy Busy bit
func (r *registerSrType) SetBusy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldBusyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldBusyMask)
	}
}

const (
	RegisterSrFieldOffuShift = 3
	RegisterSrFieldOffuMask  = 0x8
)

// GetOffu Output FIFO full
func (r *registerSrType) GetOffu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldOffuMask) != 0
}

// SetOffu Output FIFO full
func (r *registerSrType) SetOffu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldOffuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldOffuMask)
	}
}

const (
	RegisterSrFieldOfneShift = 2
	RegisterSrFieldOfneMask  = 0x4
)

// GetOfne Output FIFO not empty
func (r *registerSrType) GetOfne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldOfneMask) != 0
}

// SetOfne Output FIFO not empty
func (r *registerSrType) SetOfne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldOfneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldOfneMask)
	}
}

const (
	RegisterSrFieldIfnfShift = 1
	RegisterSrFieldIfnfMask  = 0x2
)

// GetIfnf Input FIFO not full
func (r *registerSrType) GetIfnf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIfnfMask) != 0
}

// SetIfnf Input FIFO not full
func (r *registerSrType) SetIfnf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldIfnfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldIfnfMask)
	}
}

const (
	RegisterSrFieldIfemShift = 0
	RegisterSrFieldIfemMask  = 0x1
)

// GetIfem Input FIFO empty
func (r *registerSrType) GetIfem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIfemMask) != 0
}

// SetIfem Input FIFO empty
func (r *registerSrType) SetIfem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldIfemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldIfemMask)
	}
}

// registerDinType data input register
type registerDinType uint32

const (
	RegisterDinFieldDatainShift = 0
	RegisterDinFieldDatainMask  = 0xffffffff
)

// GetDatain Data input
func (r *registerDinType) GetDatain() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDinFieldDatainMask) >> RegisterDinFieldDatainShift)
}

// SetDatain Data input
func (r *registerDinType) SetDatain(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDinFieldDatainMask)|(uint32(value)<<RegisterDinFieldDatainShift))
}

// registerDoutType data output register
type registerDoutType uint32

const (
	RegisterDoutFieldDataoutShift = 0
	RegisterDoutFieldDataoutMask  = 0xffffffff
)

// GetDataout Data output
func (r *registerDoutType) GetDataout() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDoutFieldDataoutMask) >> RegisterDoutFieldDataoutShift)
}

// SetDataout Data output
func (r *registerDoutType) SetDataout(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDoutFieldDataoutMask)|(uint32(value)<<RegisterDoutFieldDataoutShift))
}

// registerDmacrType DMA control register
type registerDmacrType uint32

const (
	RegisterDmacrFieldDoenShift = 1
	RegisterDmacrFieldDoenMask  = 0x2
)

// GetDoen DMA output enable
func (r *registerDmacrType) GetDoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacrFieldDoenMask) != 0
}

// SetDoen DMA output enable
func (r *registerDmacrType) SetDoen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacrFieldDoenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacrFieldDoenMask)
	}
}

const (
	RegisterDmacrFieldDienShift = 0
	RegisterDmacrFieldDienMask  = 0x1
)

// GetDien DMA input enable
func (r *registerDmacrType) GetDien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacrFieldDienMask) != 0
}

// SetDien DMA input enable
func (r *registerDmacrType) SetDien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacrFieldDienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacrFieldDienMask)
	}
}

// registerImscrType interrupt mask set/clear register
type registerImscrType uint32

const (
	RegisterImscrFieldOutimShift = 1
	RegisterImscrFieldOutimMask  = 0x2
)

// GetOutim Output FIFO service interrupt mask
func (r *registerImscrType) GetOutim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImscrFieldOutimMask) != 0
}

// SetOutim Output FIFO service interrupt mask
func (r *registerImscrType) SetOutim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImscrFieldOutimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImscrFieldOutimMask)
	}
}

const (
	RegisterImscrFieldInimShift = 0
	RegisterImscrFieldInimMask  = 0x1
)

// GetInim Input FIFO service interrupt mask
func (r *registerImscrType) GetInim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImscrFieldInimMask) != 0
}

// SetInim Input FIFO service interrupt mask
func (r *registerImscrType) SetInim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImscrFieldInimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImscrFieldInimMask)
	}
}

// registerRisrType raw interrupt status register
type registerRisrType uint32

const (
	RegisterRisrFieldOutrisShift = 1
	RegisterRisrFieldOutrisMask  = 0x2
)

// GetOutris Output FIFO service raw interrupt status
func (r *registerRisrType) GetOutris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisrFieldOutrisMask) != 0
}

// SetOutris Output FIFO service raw interrupt status
func (r *registerRisrType) SetOutris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisrFieldOutrisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisrFieldOutrisMask)
	}
}

const (
	RegisterRisrFieldInrisShift = 0
	RegisterRisrFieldInrisMask  = 0x1
)

// GetInris Input FIFO service raw interrupt status
func (r *registerRisrType) GetInris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRisrFieldInrisMask) != 0
}

// SetInris Input FIFO service raw interrupt status
func (r *registerRisrType) SetInris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRisrFieldInrisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRisrFieldInrisMask)
	}
}

// registerMisrType masked interrupt status register
type registerMisrType uint32

const (
	RegisterMisrFieldOutmisShift = 1
	RegisterMisrFieldOutmisMask  = 0x2
)

// GetOutmis Output FIFO service masked interrupt status
func (r *registerMisrType) GetOutmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldOutmisMask) != 0
}

// SetOutmis Output FIFO service masked interrupt status
func (r *registerMisrType) SetOutmis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisrFieldOutmisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisrFieldOutmisMask)
	}
}

const (
	RegisterMisrFieldInmisShift = 0
	RegisterMisrFieldInmisMask  = 0x1
)

// GetInmis Input FIFO service masked interrupt status
func (r *registerMisrType) GetInmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldInmisMask) != 0
}

// SetInmis Input FIFO service masked interrupt status
func (r *registerMisrType) SetInmis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisrFieldInmisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisrFieldInmisMask)
	}
}

// registerK0lrType key registers
type registerK0lrType uint32

const (
	RegisterK0lrFieldK224Shift = 0
	RegisterK0lrFieldK224Mask  = 0x1
)

// GetK224 K224
func (r *registerK0lrType) GetK224() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK224Mask) != 0
}

// SetK224 K224
func (r *registerK0lrType) SetK224(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK224Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK224Mask)
	}
}

const (
	RegisterK0lrFieldK225Shift = 1
	RegisterK0lrFieldK225Mask  = 0x2
)

// GetK225 K225
func (r *registerK0lrType) GetK225() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK225Mask) != 0
}

// SetK225 K225
func (r *registerK0lrType) SetK225(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK225Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK225Mask)
	}
}

const (
	RegisterK0lrFieldK226Shift = 2
	RegisterK0lrFieldK226Mask  = 0x4
)

// GetK226 K226
func (r *registerK0lrType) GetK226() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK226Mask) != 0
}

// SetK226 K226
func (r *registerK0lrType) SetK226(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK226Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK226Mask)
	}
}

const (
	RegisterK0lrFieldK227Shift = 3
	RegisterK0lrFieldK227Mask  = 0x8
)

// GetK227 K227
func (r *registerK0lrType) GetK227() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK227Mask) != 0
}

// SetK227 K227
func (r *registerK0lrType) SetK227(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK227Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK227Mask)
	}
}

const (
	RegisterK0lrFieldK228Shift = 4
	RegisterK0lrFieldK228Mask  = 0x10
)

// GetK228 K228
func (r *registerK0lrType) GetK228() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK228Mask) != 0
}

// SetK228 K228
func (r *registerK0lrType) SetK228(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK228Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK228Mask)
	}
}

const (
	RegisterK0lrFieldK229Shift = 5
	RegisterK0lrFieldK229Mask  = 0x20
)

// GetK229 K229
func (r *registerK0lrType) GetK229() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK229Mask) != 0
}

// SetK229 K229
func (r *registerK0lrType) SetK229(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK229Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK229Mask)
	}
}

const (
	RegisterK0lrFieldK230Shift = 6
	RegisterK0lrFieldK230Mask  = 0x40
)

// GetK230 K230
func (r *registerK0lrType) GetK230() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK230Mask) != 0
}

// SetK230 K230
func (r *registerK0lrType) SetK230(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK230Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK230Mask)
	}
}

const (
	RegisterK0lrFieldK231Shift = 7
	RegisterK0lrFieldK231Mask  = 0x80
)

// GetK231 K231
func (r *registerK0lrType) GetK231() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK231Mask) != 0
}

// SetK231 K231
func (r *registerK0lrType) SetK231(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK231Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK231Mask)
	}
}

const (
	RegisterK0lrFieldK232Shift = 8
	RegisterK0lrFieldK232Mask  = 0x100
)

// GetK232 K232
func (r *registerK0lrType) GetK232() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK232Mask) != 0
}

// SetK232 K232
func (r *registerK0lrType) SetK232(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK232Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK232Mask)
	}
}

const (
	RegisterK0lrFieldK233Shift = 9
	RegisterK0lrFieldK233Mask  = 0x200
)

// GetK233 K233
func (r *registerK0lrType) GetK233() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK233Mask) != 0
}

// SetK233 K233
func (r *registerK0lrType) SetK233(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK233Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK233Mask)
	}
}

const (
	RegisterK0lrFieldK234Shift = 10
	RegisterK0lrFieldK234Mask  = 0x400
)

// GetK234 K234
func (r *registerK0lrType) GetK234() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK234Mask) != 0
}

// SetK234 K234
func (r *registerK0lrType) SetK234(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK234Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK234Mask)
	}
}

const (
	RegisterK0lrFieldK235Shift = 11
	RegisterK0lrFieldK235Mask  = 0x800
)

// GetK235 K235
func (r *registerK0lrType) GetK235() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK235Mask) != 0
}

// SetK235 K235
func (r *registerK0lrType) SetK235(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK235Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK235Mask)
	}
}

const (
	RegisterK0lrFieldK236Shift = 12
	RegisterK0lrFieldK236Mask  = 0x1000
)

// GetK236 K236
func (r *registerK0lrType) GetK236() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK236Mask) != 0
}

// SetK236 K236
func (r *registerK0lrType) SetK236(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK236Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK236Mask)
	}
}

const (
	RegisterK0lrFieldK237Shift = 13
	RegisterK0lrFieldK237Mask  = 0x2000
)

// GetK237 K237
func (r *registerK0lrType) GetK237() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK237Mask) != 0
}

// SetK237 K237
func (r *registerK0lrType) SetK237(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK237Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK237Mask)
	}
}

const (
	RegisterK0lrFieldK238Shift = 14
	RegisterK0lrFieldK238Mask  = 0x4000
)

// GetK238 K238
func (r *registerK0lrType) GetK238() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK238Mask) != 0
}

// SetK238 K238
func (r *registerK0lrType) SetK238(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK238Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK238Mask)
	}
}

const (
	RegisterK0lrFieldK239Shift = 15
	RegisterK0lrFieldK239Mask  = 0x8000
)

// GetK239 K239
func (r *registerK0lrType) GetK239() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK239Mask) != 0
}

// SetK239 K239
func (r *registerK0lrType) SetK239(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK239Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK239Mask)
	}
}

const (
	RegisterK0lrFieldK240Shift = 16
	RegisterK0lrFieldK240Mask  = 0x10000
)

// GetK240 K240
func (r *registerK0lrType) GetK240() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK240Mask) != 0
}

// SetK240 K240
func (r *registerK0lrType) SetK240(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK240Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK240Mask)
	}
}

const (
	RegisterK0lrFieldK241Shift = 17
	RegisterK0lrFieldK241Mask  = 0x20000
)

// GetK241 K241
func (r *registerK0lrType) GetK241() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK241Mask) != 0
}

// SetK241 K241
func (r *registerK0lrType) SetK241(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK241Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK241Mask)
	}
}

const (
	RegisterK0lrFieldK242Shift = 18
	RegisterK0lrFieldK242Mask  = 0x40000
)

// GetK242 K242
func (r *registerK0lrType) GetK242() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK242Mask) != 0
}

// SetK242 K242
func (r *registerK0lrType) SetK242(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK242Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK242Mask)
	}
}

const (
	RegisterK0lrFieldK243Shift = 19
	RegisterK0lrFieldK243Mask  = 0x80000
)

// GetK243 K243
func (r *registerK0lrType) GetK243() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK243Mask) != 0
}

// SetK243 K243
func (r *registerK0lrType) SetK243(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK243Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK243Mask)
	}
}

const (
	RegisterK0lrFieldK244Shift = 20
	RegisterK0lrFieldK244Mask  = 0x100000
)

// GetK244 K244
func (r *registerK0lrType) GetK244() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK244Mask) != 0
}

// SetK244 K244
func (r *registerK0lrType) SetK244(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK244Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK244Mask)
	}
}

const (
	RegisterK0lrFieldK245Shift = 21
	RegisterK0lrFieldK245Mask  = 0x200000
)

// GetK245 K245
func (r *registerK0lrType) GetK245() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK245Mask) != 0
}

// SetK245 K245
func (r *registerK0lrType) SetK245(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK245Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK245Mask)
	}
}

const (
	RegisterK0lrFieldK246Shift = 22
	RegisterK0lrFieldK246Mask  = 0x400000
)

// GetK246 K246
func (r *registerK0lrType) GetK246() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK246Mask) != 0
}

// SetK246 K246
func (r *registerK0lrType) SetK246(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK246Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK246Mask)
	}
}

const (
	RegisterK0lrFieldK247Shift = 23
	RegisterK0lrFieldK247Mask  = 0x800000
)

// GetK247 K247
func (r *registerK0lrType) GetK247() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK247Mask) != 0
}

// SetK247 K247
func (r *registerK0lrType) SetK247(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK247Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK247Mask)
	}
}

const (
	RegisterK0lrFieldK248Shift = 24
	RegisterK0lrFieldK248Mask  = 0x1000000
)

// GetK248 K248
func (r *registerK0lrType) GetK248() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK248Mask) != 0
}

// SetK248 K248
func (r *registerK0lrType) SetK248(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK248Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK248Mask)
	}
}

const (
	RegisterK0lrFieldK249Shift = 25
	RegisterK0lrFieldK249Mask  = 0x2000000
)

// GetK249 K249
func (r *registerK0lrType) GetK249() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK249Mask) != 0
}

// SetK249 K249
func (r *registerK0lrType) SetK249(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK249Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK249Mask)
	}
}

const (
	RegisterK0lrFieldK250Shift = 26
	RegisterK0lrFieldK250Mask  = 0x4000000
)

// GetK250 K250
func (r *registerK0lrType) GetK250() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK250Mask) != 0
}

// SetK250 K250
func (r *registerK0lrType) SetK250(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK250Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK250Mask)
	}
}

const (
	RegisterK0lrFieldK251Shift = 27
	RegisterK0lrFieldK251Mask  = 0x8000000
)

// GetK251 K251
func (r *registerK0lrType) GetK251() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK251Mask) != 0
}

// SetK251 K251
func (r *registerK0lrType) SetK251(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK251Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK251Mask)
	}
}

const (
	RegisterK0lrFieldK252Shift = 28
	RegisterK0lrFieldK252Mask  = 0x10000000
)

// GetK252 K252
func (r *registerK0lrType) GetK252() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK252Mask) != 0
}

// SetK252 K252
func (r *registerK0lrType) SetK252(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK252Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK252Mask)
	}
}

const (
	RegisterK0lrFieldK253Shift = 29
	RegisterK0lrFieldK253Mask  = 0x20000000
)

// GetK253 K253
func (r *registerK0lrType) GetK253() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK253Mask) != 0
}

// SetK253 K253
func (r *registerK0lrType) SetK253(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK253Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK253Mask)
	}
}

const (
	RegisterK0lrFieldK254Shift = 30
	RegisterK0lrFieldK254Mask  = 0x40000000
)

// GetK254 K254
func (r *registerK0lrType) GetK254() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK254Mask) != 0
}

// SetK254 K254
func (r *registerK0lrType) SetK254(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK254Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK254Mask)
	}
}

const (
	RegisterK0lrFieldK255Shift = 31
	RegisterK0lrFieldK255Mask  = 0x80000000
)

// GetK255 K255
func (r *registerK0lrType) GetK255() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldK255Mask) != 0
}

// SetK255 K255
func (r *registerK0lrType) SetK255(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldK255Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldK255Mask)
	}
}

// registerK0rrType key registers
type registerK0rrType uint32

const (
	RegisterK0rrFieldK192Shift = 0
	RegisterK0rrFieldK192Mask  = 0x1
)

// GetK192 K192
func (r *registerK0rrType) GetK192() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK192Mask) != 0
}

// SetK192 K192
func (r *registerK0rrType) SetK192(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK192Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK192Mask)
	}
}

const (
	RegisterK0rrFieldK193Shift = 1
	RegisterK0rrFieldK193Mask  = 0x2
)

// GetK193 K193
func (r *registerK0rrType) GetK193() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK193Mask) != 0
}

// SetK193 K193
func (r *registerK0rrType) SetK193(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK193Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK193Mask)
	}
}

const (
	RegisterK0rrFieldK194Shift = 2
	RegisterK0rrFieldK194Mask  = 0x4
)

// GetK194 K194
func (r *registerK0rrType) GetK194() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK194Mask) != 0
}

// SetK194 K194
func (r *registerK0rrType) SetK194(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK194Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK194Mask)
	}
}

const (
	RegisterK0rrFieldK195Shift = 3
	RegisterK0rrFieldK195Mask  = 0x8
)

// GetK195 K195
func (r *registerK0rrType) GetK195() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK195Mask) != 0
}

// SetK195 K195
func (r *registerK0rrType) SetK195(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK195Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK195Mask)
	}
}

const (
	RegisterK0rrFieldK196Shift = 4
	RegisterK0rrFieldK196Mask  = 0x10
)

// GetK196 K196
func (r *registerK0rrType) GetK196() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK196Mask) != 0
}

// SetK196 K196
func (r *registerK0rrType) SetK196(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK196Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK196Mask)
	}
}

const (
	RegisterK0rrFieldK197Shift = 5
	RegisterK0rrFieldK197Mask  = 0x20
)

// GetK197 K197
func (r *registerK0rrType) GetK197() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK197Mask) != 0
}

// SetK197 K197
func (r *registerK0rrType) SetK197(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK197Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK197Mask)
	}
}

const (
	RegisterK0rrFieldK198Shift = 6
	RegisterK0rrFieldK198Mask  = 0x40
)

// GetK198 K198
func (r *registerK0rrType) GetK198() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK198Mask) != 0
}

// SetK198 K198
func (r *registerK0rrType) SetK198(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK198Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK198Mask)
	}
}

const (
	RegisterK0rrFieldK199Shift = 7
	RegisterK0rrFieldK199Mask  = 0x80
)

// GetK199 K199
func (r *registerK0rrType) GetK199() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK199Mask) != 0
}

// SetK199 K199
func (r *registerK0rrType) SetK199(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK199Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK199Mask)
	}
}

const (
	RegisterK0rrFieldK200Shift = 8
	RegisterK0rrFieldK200Mask  = 0x100
)

// GetK200 K200
func (r *registerK0rrType) GetK200() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK200Mask) != 0
}

// SetK200 K200
func (r *registerK0rrType) SetK200(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK200Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK200Mask)
	}
}

const (
	RegisterK0rrFieldK201Shift = 9
	RegisterK0rrFieldK201Mask  = 0x200
)

// GetK201 K201
func (r *registerK0rrType) GetK201() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK201Mask) != 0
}

// SetK201 K201
func (r *registerK0rrType) SetK201(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK201Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK201Mask)
	}
}

const (
	RegisterK0rrFieldK202Shift = 10
	RegisterK0rrFieldK202Mask  = 0x400
)

// GetK202 K202
func (r *registerK0rrType) GetK202() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK202Mask) != 0
}

// SetK202 K202
func (r *registerK0rrType) SetK202(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK202Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK202Mask)
	}
}

const (
	RegisterK0rrFieldK203Shift = 11
	RegisterK0rrFieldK203Mask  = 0x800
)

// GetK203 K203
func (r *registerK0rrType) GetK203() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK203Mask) != 0
}

// SetK203 K203
func (r *registerK0rrType) SetK203(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK203Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK203Mask)
	}
}

const (
	RegisterK0rrFieldK204Shift = 12
	RegisterK0rrFieldK204Mask  = 0x1000
)

// GetK204 K204
func (r *registerK0rrType) GetK204() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK204Mask) != 0
}

// SetK204 K204
func (r *registerK0rrType) SetK204(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK204Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK204Mask)
	}
}

const (
	RegisterK0rrFieldK205Shift = 13
	RegisterK0rrFieldK205Mask  = 0x2000
)

// GetK205 K205
func (r *registerK0rrType) GetK205() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK205Mask) != 0
}

// SetK205 K205
func (r *registerK0rrType) SetK205(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK205Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK205Mask)
	}
}

const (
	RegisterK0rrFieldK206Shift = 14
	RegisterK0rrFieldK206Mask  = 0x4000
)

// GetK206 K206
func (r *registerK0rrType) GetK206() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK206Mask) != 0
}

// SetK206 K206
func (r *registerK0rrType) SetK206(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK206Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK206Mask)
	}
}

const (
	RegisterK0rrFieldK207Shift = 15
	RegisterK0rrFieldK207Mask  = 0x8000
)

// GetK207 K207
func (r *registerK0rrType) GetK207() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK207Mask) != 0
}

// SetK207 K207
func (r *registerK0rrType) SetK207(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK207Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK207Mask)
	}
}

const (
	RegisterK0rrFieldK208Shift = 16
	RegisterK0rrFieldK208Mask  = 0x10000
)

// GetK208 K208
func (r *registerK0rrType) GetK208() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK208Mask) != 0
}

// SetK208 K208
func (r *registerK0rrType) SetK208(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK208Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK208Mask)
	}
}

const (
	RegisterK0rrFieldK209Shift = 17
	RegisterK0rrFieldK209Mask  = 0x20000
)

// GetK209 K209
func (r *registerK0rrType) GetK209() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK209Mask) != 0
}

// SetK209 K209
func (r *registerK0rrType) SetK209(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK209Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK209Mask)
	}
}

const (
	RegisterK0rrFieldK210Shift = 18
	RegisterK0rrFieldK210Mask  = 0x40000
)

// GetK210 K210
func (r *registerK0rrType) GetK210() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK210Mask) != 0
}

// SetK210 K210
func (r *registerK0rrType) SetK210(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK210Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK210Mask)
	}
}

const (
	RegisterK0rrFieldK211Shift = 19
	RegisterK0rrFieldK211Mask  = 0x80000
)

// GetK211 K211
func (r *registerK0rrType) GetK211() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK211Mask) != 0
}

// SetK211 K211
func (r *registerK0rrType) SetK211(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK211Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK211Mask)
	}
}

const (
	RegisterK0rrFieldK212Shift = 20
	RegisterK0rrFieldK212Mask  = 0x100000
)

// GetK212 K212
func (r *registerK0rrType) GetK212() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK212Mask) != 0
}

// SetK212 K212
func (r *registerK0rrType) SetK212(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK212Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK212Mask)
	}
}

const (
	RegisterK0rrFieldK213Shift = 21
	RegisterK0rrFieldK213Mask  = 0x200000
)

// GetK213 K213
func (r *registerK0rrType) GetK213() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK213Mask) != 0
}

// SetK213 K213
func (r *registerK0rrType) SetK213(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK213Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK213Mask)
	}
}

const (
	RegisterK0rrFieldK214Shift = 22
	RegisterK0rrFieldK214Mask  = 0x400000
)

// GetK214 K214
func (r *registerK0rrType) GetK214() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK214Mask) != 0
}

// SetK214 K214
func (r *registerK0rrType) SetK214(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK214Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK214Mask)
	}
}

const (
	RegisterK0rrFieldK215Shift = 23
	RegisterK0rrFieldK215Mask  = 0x800000
)

// GetK215 K215
func (r *registerK0rrType) GetK215() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK215Mask) != 0
}

// SetK215 K215
func (r *registerK0rrType) SetK215(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK215Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK215Mask)
	}
}

const (
	RegisterK0rrFieldK216Shift = 24
	RegisterK0rrFieldK216Mask  = 0x1000000
)

// GetK216 K216
func (r *registerK0rrType) GetK216() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK216Mask) != 0
}

// SetK216 K216
func (r *registerK0rrType) SetK216(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK216Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK216Mask)
	}
}

const (
	RegisterK0rrFieldK217Shift = 25
	RegisterK0rrFieldK217Mask  = 0x2000000
)

// GetK217 K217
func (r *registerK0rrType) GetK217() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK217Mask) != 0
}

// SetK217 K217
func (r *registerK0rrType) SetK217(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK217Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK217Mask)
	}
}

const (
	RegisterK0rrFieldK218Shift = 26
	RegisterK0rrFieldK218Mask  = 0x4000000
)

// GetK218 K218
func (r *registerK0rrType) GetK218() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK218Mask) != 0
}

// SetK218 K218
func (r *registerK0rrType) SetK218(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK218Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK218Mask)
	}
}

const (
	RegisterK0rrFieldK219Shift = 27
	RegisterK0rrFieldK219Mask  = 0x8000000
)

// GetK219 K219
func (r *registerK0rrType) GetK219() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK219Mask) != 0
}

// SetK219 K219
func (r *registerK0rrType) SetK219(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK219Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK219Mask)
	}
}

const (
	RegisterK0rrFieldK220Shift = 28
	RegisterK0rrFieldK220Mask  = 0x10000000
)

// GetK220 K220
func (r *registerK0rrType) GetK220() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK220Mask) != 0
}

// SetK220 K220
func (r *registerK0rrType) SetK220(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK220Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK220Mask)
	}
}

const (
	RegisterK0rrFieldK221Shift = 29
	RegisterK0rrFieldK221Mask  = 0x20000000
)

// GetK221 K221
func (r *registerK0rrType) GetK221() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK221Mask) != 0
}

// SetK221 K221
func (r *registerK0rrType) SetK221(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK221Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK221Mask)
	}
}

const (
	RegisterK0rrFieldK222Shift = 30
	RegisterK0rrFieldK222Mask  = 0x40000000
)

// GetK222 K222
func (r *registerK0rrType) GetK222() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK222Mask) != 0
}

// SetK222 K222
func (r *registerK0rrType) SetK222(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK222Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK222Mask)
	}
}

const (
	RegisterK0rrFieldK223Shift = 31
	RegisterK0rrFieldK223Mask  = 0x80000000
)

// GetK223 K223
func (r *registerK0rrType) GetK223() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldK223Mask) != 0
}

// SetK223 K223
func (r *registerK0rrType) SetK223(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldK223Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldK223Mask)
	}
}

// registerK1lrType key registers
type registerK1lrType uint32

const (
	RegisterK1lrFieldK160Shift = 0
	RegisterK1lrFieldK160Mask  = 0x1
)

// GetK160 K160
func (r *registerK1lrType) GetK160() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK160Mask) != 0
}

// SetK160 K160
func (r *registerK1lrType) SetK160(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK160Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK160Mask)
	}
}

const (
	RegisterK1lrFieldK161Shift = 1
	RegisterK1lrFieldK161Mask  = 0x2
)

// GetK161 K161
func (r *registerK1lrType) GetK161() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK161Mask) != 0
}

// SetK161 K161
func (r *registerK1lrType) SetK161(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK161Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK161Mask)
	}
}

const (
	RegisterK1lrFieldK162Shift = 2
	RegisterK1lrFieldK162Mask  = 0x4
)

// GetK162 K162
func (r *registerK1lrType) GetK162() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK162Mask) != 0
}

// SetK162 K162
func (r *registerK1lrType) SetK162(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK162Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK162Mask)
	}
}

const (
	RegisterK1lrFieldK163Shift = 3
	RegisterK1lrFieldK163Mask  = 0x8
)

// GetK163 K163
func (r *registerK1lrType) GetK163() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK163Mask) != 0
}

// SetK163 K163
func (r *registerK1lrType) SetK163(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK163Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK163Mask)
	}
}

const (
	RegisterK1lrFieldK164Shift = 4
	RegisterK1lrFieldK164Mask  = 0x10
)

// GetK164 K164
func (r *registerK1lrType) GetK164() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK164Mask) != 0
}

// SetK164 K164
func (r *registerK1lrType) SetK164(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK164Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK164Mask)
	}
}

const (
	RegisterK1lrFieldK165Shift = 5
	RegisterK1lrFieldK165Mask  = 0x20
)

// GetK165 K165
func (r *registerK1lrType) GetK165() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK165Mask) != 0
}

// SetK165 K165
func (r *registerK1lrType) SetK165(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK165Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK165Mask)
	}
}

const (
	RegisterK1lrFieldK166Shift = 6
	RegisterK1lrFieldK166Mask  = 0x40
)

// GetK166 K166
func (r *registerK1lrType) GetK166() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK166Mask) != 0
}

// SetK166 K166
func (r *registerK1lrType) SetK166(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK166Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK166Mask)
	}
}

const (
	RegisterK1lrFieldK167Shift = 7
	RegisterK1lrFieldK167Mask  = 0x80
)

// GetK167 K167
func (r *registerK1lrType) GetK167() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK167Mask) != 0
}

// SetK167 K167
func (r *registerK1lrType) SetK167(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK167Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK167Mask)
	}
}

const (
	RegisterK1lrFieldK168Shift = 8
	RegisterK1lrFieldK168Mask  = 0x100
)

// GetK168 K168
func (r *registerK1lrType) GetK168() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK168Mask) != 0
}

// SetK168 K168
func (r *registerK1lrType) SetK168(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK168Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK168Mask)
	}
}

const (
	RegisterK1lrFieldK169Shift = 9
	RegisterK1lrFieldK169Mask  = 0x200
)

// GetK169 K169
func (r *registerK1lrType) GetK169() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK169Mask) != 0
}

// SetK169 K169
func (r *registerK1lrType) SetK169(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK169Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK169Mask)
	}
}

const (
	RegisterK1lrFieldK170Shift = 10
	RegisterK1lrFieldK170Mask  = 0x400
)

// GetK170 K170
func (r *registerK1lrType) GetK170() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK170Mask) != 0
}

// SetK170 K170
func (r *registerK1lrType) SetK170(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK170Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK170Mask)
	}
}

const (
	RegisterK1lrFieldK171Shift = 11
	RegisterK1lrFieldK171Mask  = 0x800
)

// GetK171 K171
func (r *registerK1lrType) GetK171() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK171Mask) != 0
}

// SetK171 K171
func (r *registerK1lrType) SetK171(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK171Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK171Mask)
	}
}

const (
	RegisterK1lrFieldK172Shift = 12
	RegisterK1lrFieldK172Mask  = 0x1000
)

// GetK172 K172
func (r *registerK1lrType) GetK172() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK172Mask) != 0
}

// SetK172 K172
func (r *registerK1lrType) SetK172(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK172Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK172Mask)
	}
}

const (
	RegisterK1lrFieldK173Shift = 13
	RegisterK1lrFieldK173Mask  = 0x2000
)

// GetK173 K173
func (r *registerK1lrType) GetK173() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK173Mask) != 0
}

// SetK173 K173
func (r *registerK1lrType) SetK173(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK173Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK173Mask)
	}
}

const (
	RegisterK1lrFieldK174Shift = 14
	RegisterK1lrFieldK174Mask  = 0x4000
)

// GetK174 K174
func (r *registerK1lrType) GetK174() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK174Mask) != 0
}

// SetK174 K174
func (r *registerK1lrType) SetK174(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK174Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK174Mask)
	}
}

const (
	RegisterK1lrFieldK175Shift = 15
	RegisterK1lrFieldK175Mask  = 0x8000
)

// GetK175 K175
func (r *registerK1lrType) GetK175() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK175Mask) != 0
}

// SetK175 K175
func (r *registerK1lrType) SetK175(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK175Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK175Mask)
	}
}

const (
	RegisterK1lrFieldK176Shift = 16
	RegisterK1lrFieldK176Mask  = 0x10000
)

// GetK176 K176
func (r *registerK1lrType) GetK176() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK176Mask) != 0
}

// SetK176 K176
func (r *registerK1lrType) SetK176(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK176Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK176Mask)
	}
}

const (
	RegisterK1lrFieldK177Shift = 17
	RegisterK1lrFieldK177Mask  = 0x20000
)

// GetK177 K177
func (r *registerK1lrType) GetK177() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK177Mask) != 0
}

// SetK177 K177
func (r *registerK1lrType) SetK177(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK177Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK177Mask)
	}
}

const (
	RegisterK1lrFieldK178Shift = 18
	RegisterK1lrFieldK178Mask  = 0x40000
)

// GetK178 K178
func (r *registerK1lrType) GetK178() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK178Mask) != 0
}

// SetK178 K178
func (r *registerK1lrType) SetK178(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK178Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK178Mask)
	}
}

const (
	RegisterK1lrFieldK179Shift = 19
	RegisterK1lrFieldK179Mask  = 0x80000
)

// GetK179 K179
func (r *registerK1lrType) GetK179() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK179Mask) != 0
}

// SetK179 K179
func (r *registerK1lrType) SetK179(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK179Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK179Mask)
	}
}

const (
	RegisterK1lrFieldK180Shift = 20
	RegisterK1lrFieldK180Mask  = 0x100000
)

// GetK180 K180
func (r *registerK1lrType) GetK180() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK180Mask) != 0
}

// SetK180 K180
func (r *registerK1lrType) SetK180(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK180Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK180Mask)
	}
}

const (
	RegisterK1lrFieldK181Shift = 21
	RegisterK1lrFieldK181Mask  = 0x200000
)

// GetK181 K181
func (r *registerK1lrType) GetK181() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK181Mask) != 0
}

// SetK181 K181
func (r *registerK1lrType) SetK181(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK181Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK181Mask)
	}
}

const (
	RegisterK1lrFieldK182Shift = 22
	RegisterK1lrFieldK182Mask  = 0x400000
)

// GetK182 K182
func (r *registerK1lrType) GetK182() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK182Mask) != 0
}

// SetK182 K182
func (r *registerK1lrType) SetK182(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK182Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK182Mask)
	}
}

const (
	RegisterK1lrFieldK183Shift = 23
	RegisterK1lrFieldK183Mask  = 0x800000
)

// GetK183 K183
func (r *registerK1lrType) GetK183() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK183Mask) != 0
}

// SetK183 K183
func (r *registerK1lrType) SetK183(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK183Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK183Mask)
	}
}

const (
	RegisterK1lrFieldK184Shift = 24
	RegisterK1lrFieldK184Mask  = 0x1000000
)

// GetK184 K184
func (r *registerK1lrType) GetK184() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK184Mask) != 0
}

// SetK184 K184
func (r *registerK1lrType) SetK184(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK184Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK184Mask)
	}
}

const (
	RegisterK1lrFieldK185Shift = 25
	RegisterK1lrFieldK185Mask  = 0x2000000
)

// GetK185 K185
func (r *registerK1lrType) GetK185() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK185Mask) != 0
}

// SetK185 K185
func (r *registerK1lrType) SetK185(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK185Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK185Mask)
	}
}

const (
	RegisterK1lrFieldK186Shift = 26
	RegisterK1lrFieldK186Mask  = 0x4000000
)

// GetK186 K186
func (r *registerK1lrType) GetK186() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK186Mask) != 0
}

// SetK186 K186
func (r *registerK1lrType) SetK186(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK186Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK186Mask)
	}
}

const (
	RegisterK1lrFieldK187Shift = 27
	RegisterK1lrFieldK187Mask  = 0x8000000
)

// GetK187 K187
func (r *registerK1lrType) GetK187() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK187Mask) != 0
}

// SetK187 K187
func (r *registerK1lrType) SetK187(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK187Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK187Mask)
	}
}

const (
	RegisterK1lrFieldK188Shift = 28
	RegisterK1lrFieldK188Mask  = 0x10000000
)

// GetK188 K188
func (r *registerK1lrType) GetK188() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK188Mask) != 0
}

// SetK188 K188
func (r *registerK1lrType) SetK188(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK188Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK188Mask)
	}
}

const (
	RegisterK1lrFieldK189Shift = 29
	RegisterK1lrFieldK189Mask  = 0x20000000
)

// GetK189 K189
func (r *registerK1lrType) GetK189() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK189Mask) != 0
}

// SetK189 K189
func (r *registerK1lrType) SetK189(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK189Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK189Mask)
	}
}

const (
	RegisterK1lrFieldK190Shift = 30
	RegisterK1lrFieldK190Mask  = 0x40000000
)

// GetK190 K190
func (r *registerK1lrType) GetK190() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK190Mask) != 0
}

// SetK190 K190
func (r *registerK1lrType) SetK190(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK190Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK190Mask)
	}
}

const (
	RegisterK1lrFieldK191Shift = 31
	RegisterK1lrFieldK191Mask  = 0x80000000
)

// GetK191 K191
func (r *registerK1lrType) GetK191() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldK191Mask) != 0
}

// SetK191 K191
func (r *registerK1lrType) SetK191(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldK191Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldK191Mask)
	}
}

// registerK1rrType key registers
type registerK1rrType uint32

const (
	RegisterK1rrFieldK128Shift = 0
	RegisterK1rrFieldK128Mask  = 0x1
)

// GetK128 K128
func (r *registerK1rrType) GetK128() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK128Mask) != 0
}

// SetK128 K128
func (r *registerK1rrType) SetK128(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK128Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK128Mask)
	}
}

const (
	RegisterK1rrFieldK129Shift = 1
	RegisterK1rrFieldK129Mask  = 0x2
)

// GetK129 K129
func (r *registerK1rrType) GetK129() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK129Mask) != 0
}

// SetK129 K129
func (r *registerK1rrType) SetK129(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK129Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK129Mask)
	}
}

const (
	RegisterK1rrFieldK130Shift = 2
	RegisterK1rrFieldK130Mask  = 0x4
)

// GetK130 K130
func (r *registerK1rrType) GetK130() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK130Mask) != 0
}

// SetK130 K130
func (r *registerK1rrType) SetK130(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK130Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK130Mask)
	}
}

const (
	RegisterK1rrFieldK131Shift = 3
	RegisterK1rrFieldK131Mask  = 0x8
)

// GetK131 K131
func (r *registerK1rrType) GetK131() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK131Mask) != 0
}

// SetK131 K131
func (r *registerK1rrType) SetK131(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK131Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK131Mask)
	}
}

const (
	RegisterK1rrFieldK132Shift = 4
	RegisterK1rrFieldK132Mask  = 0x10
)

// GetK132 K132
func (r *registerK1rrType) GetK132() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK132Mask) != 0
}

// SetK132 K132
func (r *registerK1rrType) SetK132(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK132Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK132Mask)
	}
}

const (
	RegisterK1rrFieldK133Shift = 5
	RegisterK1rrFieldK133Mask  = 0x20
)

// GetK133 K133
func (r *registerK1rrType) GetK133() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK133Mask) != 0
}

// SetK133 K133
func (r *registerK1rrType) SetK133(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK133Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK133Mask)
	}
}

const (
	RegisterK1rrFieldK134Shift = 6
	RegisterK1rrFieldK134Mask  = 0x40
)

// GetK134 K134
func (r *registerK1rrType) GetK134() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK134Mask) != 0
}

// SetK134 K134
func (r *registerK1rrType) SetK134(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK134Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK134Mask)
	}
}

const (
	RegisterK1rrFieldK135Shift = 7
	RegisterK1rrFieldK135Mask  = 0x80
)

// GetK135 K135
func (r *registerK1rrType) GetK135() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK135Mask) != 0
}

// SetK135 K135
func (r *registerK1rrType) SetK135(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK135Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK135Mask)
	}
}

const (
	RegisterK1rrFieldK136Shift = 8
	RegisterK1rrFieldK136Mask  = 0x100
)

// GetK136 K136
func (r *registerK1rrType) GetK136() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK136Mask) != 0
}

// SetK136 K136
func (r *registerK1rrType) SetK136(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK136Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK136Mask)
	}
}

const (
	RegisterK1rrFieldK137Shift = 9
	RegisterK1rrFieldK137Mask  = 0x200
)

// GetK137 K137
func (r *registerK1rrType) GetK137() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK137Mask) != 0
}

// SetK137 K137
func (r *registerK1rrType) SetK137(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK137Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK137Mask)
	}
}

const (
	RegisterK1rrFieldK138Shift = 10
	RegisterK1rrFieldK138Mask  = 0x400
)

// GetK138 K138
func (r *registerK1rrType) GetK138() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK138Mask) != 0
}

// SetK138 K138
func (r *registerK1rrType) SetK138(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK138Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK138Mask)
	}
}

const (
	RegisterK1rrFieldK139Shift = 11
	RegisterK1rrFieldK139Mask  = 0x800
)

// GetK139 K139
func (r *registerK1rrType) GetK139() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK139Mask) != 0
}

// SetK139 K139
func (r *registerK1rrType) SetK139(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK139Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK139Mask)
	}
}

const (
	RegisterK1rrFieldK140Shift = 12
	RegisterK1rrFieldK140Mask  = 0x1000
)

// GetK140 K140
func (r *registerK1rrType) GetK140() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK140Mask) != 0
}

// SetK140 K140
func (r *registerK1rrType) SetK140(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK140Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK140Mask)
	}
}

const (
	RegisterK1rrFieldK141Shift = 13
	RegisterK1rrFieldK141Mask  = 0x2000
)

// GetK141 K141
func (r *registerK1rrType) GetK141() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK141Mask) != 0
}

// SetK141 K141
func (r *registerK1rrType) SetK141(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK141Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK141Mask)
	}
}

const (
	RegisterK1rrFieldK142Shift = 14
	RegisterK1rrFieldK142Mask  = 0x4000
)

// GetK142 K142
func (r *registerK1rrType) GetK142() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK142Mask) != 0
}

// SetK142 K142
func (r *registerK1rrType) SetK142(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK142Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK142Mask)
	}
}

const (
	RegisterK1rrFieldK143Shift = 15
	RegisterK1rrFieldK143Mask  = 0x8000
)

// GetK143 K143
func (r *registerK1rrType) GetK143() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK143Mask) != 0
}

// SetK143 K143
func (r *registerK1rrType) SetK143(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK143Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK143Mask)
	}
}

const (
	RegisterK1rrFieldK144Shift = 16
	RegisterK1rrFieldK144Mask  = 0x10000
)

// GetK144 K144
func (r *registerK1rrType) GetK144() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK144Mask) != 0
}

// SetK144 K144
func (r *registerK1rrType) SetK144(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK144Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK144Mask)
	}
}

const (
	RegisterK1rrFieldK145Shift = 17
	RegisterK1rrFieldK145Mask  = 0x20000
)

// GetK145 K145
func (r *registerK1rrType) GetK145() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK145Mask) != 0
}

// SetK145 K145
func (r *registerK1rrType) SetK145(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK145Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK145Mask)
	}
}

const (
	RegisterK1rrFieldK146Shift = 18
	RegisterK1rrFieldK146Mask  = 0x40000
)

// GetK146 K146
func (r *registerK1rrType) GetK146() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK146Mask) != 0
}

// SetK146 K146
func (r *registerK1rrType) SetK146(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK146Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK146Mask)
	}
}

const (
	RegisterK1rrFieldK147Shift = 19
	RegisterK1rrFieldK147Mask  = 0x80000
)

// GetK147 K147
func (r *registerK1rrType) GetK147() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK147Mask) != 0
}

// SetK147 K147
func (r *registerK1rrType) SetK147(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK147Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK147Mask)
	}
}

const (
	RegisterK1rrFieldK148Shift = 20
	RegisterK1rrFieldK148Mask  = 0x100000
)

// GetK148 K148
func (r *registerK1rrType) GetK148() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK148Mask) != 0
}

// SetK148 K148
func (r *registerK1rrType) SetK148(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK148Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK148Mask)
	}
}

const (
	RegisterK1rrFieldK149Shift = 21
	RegisterK1rrFieldK149Mask  = 0x200000
)

// GetK149 K149
func (r *registerK1rrType) GetK149() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK149Mask) != 0
}

// SetK149 K149
func (r *registerK1rrType) SetK149(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK149Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK149Mask)
	}
}

const (
	RegisterK1rrFieldK150Shift = 22
	RegisterK1rrFieldK150Mask  = 0x400000
)

// GetK150 K150
func (r *registerK1rrType) GetK150() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK150Mask) != 0
}

// SetK150 K150
func (r *registerK1rrType) SetK150(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK150Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK150Mask)
	}
}

const (
	RegisterK1rrFieldK151Shift = 23
	RegisterK1rrFieldK151Mask  = 0x800000
)

// GetK151 K151
func (r *registerK1rrType) GetK151() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK151Mask) != 0
}

// SetK151 K151
func (r *registerK1rrType) SetK151(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK151Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK151Mask)
	}
}

const (
	RegisterK1rrFieldK152Shift = 24
	RegisterK1rrFieldK152Mask  = 0x1000000
)

// GetK152 K152
func (r *registerK1rrType) GetK152() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK152Mask) != 0
}

// SetK152 K152
func (r *registerK1rrType) SetK152(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK152Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK152Mask)
	}
}

const (
	RegisterK1rrFieldK153Shift = 25
	RegisterK1rrFieldK153Mask  = 0x2000000
)

// GetK153 K153
func (r *registerK1rrType) GetK153() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK153Mask) != 0
}

// SetK153 K153
func (r *registerK1rrType) SetK153(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK153Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK153Mask)
	}
}

const (
	RegisterK1rrFieldK154Shift = 26
	RegisterK1rrFieldK154Mask  = 0x4000000
)

// GetK154 K154
func (r *registerK1rrType) GetK154() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK154Mask) != 0
}

// SetK154 K154
func (r *registerK1rrType) SetK154(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK154Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK154Mask)
	}
}

const (
	RegisterK1rrFieldK155Shift = 27
	RegisterK1rrFieldK155Mask  = 0x8000000
)

// GetK155 K155
func (r *registerK1rrType) GetK155() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK155Mask) != 0
}

// SetK155 K155
func (r *registerK1rrType) SetK155(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK155Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK155Mask)
	}
}

const (
	RegisterK1rrFieldK156Shift = 28
	RegisterK1rrFieldK156Mask  = 0x10000000
)

// GetK156 K156
func (r *registerK1rrType) GetK156() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK156Mask) != 0
}

// SetK156 K156
func (r *registerK1rrType) SetK156(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK156Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK156Mask)
	}
}

const (
	RegisterK1rrFieldK157Shift = 29
	RegisterK1rrFieldK157Mask  = 0x20000000
)

// GetK157 K157
func (r *registerK1rrType) GetK157() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK157Mask) != 0
}

// SetK157 K157
func (r *registerK1rrType) SetK157(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK157Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK157Mask)
	}
}

const (
	RegisterK1rrFieldK158Shift = 30
	RegisterK1rrFieldK158Mask  = 0x40000000
)

// GetK158 K158
func (r *registerK1rrType) GetK158() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK158Mask) != 0
}

// SetK158 K158
func (r *registerK1rrType) SetK158(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK158Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK158Mask)
	}
}

const (
	RegisterK1rrFieldK159Shift = 31
	RegisterK1rrFieldK159Mask  = 0x80000000
)

// GetK159 K159
func (r *registerK1rrType) GetK159() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldK159Mask) != 0
}

// SetK159 K159
func (r *registerK1rrType) SetK159(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldK159Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldK159Mask)
	}
}

// registerK2lrType key registers
type registerK2lrType uint32

const (
	RegisterK2lrFieldK96Shift = 0
	RegisterK2lrFieldK96Mask  = 0x1
)

// GetK96 K96
func (r *registerK2lrType) GetK96() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK96Mask) != 0
}

// SetK96 K96
func (r *registerK2lrType) SetK96(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK96Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK96Mask)
	}
}

const (
	RegisterK2lrFieldK97Shift = 1
	RegisterK2lrFieldK97Mask  = 0x2
)

// GetK97 K97
func (r *registerK2lrType) GetK97() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK97Mask) != 0
}

// SetK97 K97
func (r *registerK2lrType) SetK97(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK97Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK97Mask)
	}
}

const (
	RegisterK2lrFieldK98Shift = 2
	RegisterK2lrFieldK98Mask  = 0x4
)

// GetK98 K98
func (r *registerK2lrType) GetK98() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK98Mask) != 0
}

// SetK98 K98
func (r *registerK2lrType) SetK98(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK98Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK98Mask)
	}
}

const (
	RegisterK2lrFieldK99Shift = 3
	RegisterK2lrFieldK99Mask  = 0x8
)

// GetK99 K99
func (r *registerK2lrType) GetK99() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK99Mask) != 0
}

// SetK99 K99
func (r *registerK2lrType) SetK99(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK99Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK99Mask)
	}
}

const (
	RegisterK2lrFieldK100Shift = 4
	RegisterK2lrFieldK100Mask  = 0x10
)

// GetK100 K100
func (r *registerK2lrType) GetK100() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK100Mask) != 0
}

// SetK100 K100
func (r *registerK2lrType) SetK100(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK100Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK100Mask)
	}
}

const (
	RegisterK2lrFieldK101Shift = 5
	RegisterK2lrFieldK101Mask  = 0x20
)

// GetK101 K101
func (r *registerK2lrType) GetK101() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK101Mask) != 0
}

// SetK101 K101
func (r *registerK2lrType) SetK101(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK101Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK101Mask)
	}
}

const (
	RegisterK2lrFieldK102Shift = 6
	RegisterK2lrFieldK102Mask  = 0x40
)

// GetK102 K102
func (r *registerK2lrType) GetK102() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK102Mask) != 0
}

// SetK102 K102
func (r *registerK2lrType) SetK102(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK102Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK102Mask)
	}
}

const (
	RegisterK2lrFieldK103Shift = 7
	RegisterK2lrFieldK103Mask  = 0x80
)

// GetK103 K103
func (r *registerK2lrType) GetK103() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK103Mask) != 0
}

// SetK103 K103
func (r *registerK2lrType) SetK103(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK103Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK103Mask)
	}
}

const (
	RegisterK2lrFieldK104Shift = 8
	RegisterK2lrFieldK104Mask  = 0x100
)

// GetK104 K104
func (r *registerK2lrType) GetK104() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK104Mask) != 0
}

// SetK104 K104
func (r *registerK2lrType) SetK104(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK104Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK104Mask)
	}
}

const (
	RegisterK2lrFieldK105Shift = 9
	RegisterK2lrFieldK105Mask  = 0x200
)

// GetK105 K105
func (r *registerK2lrType) GetK105() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK105Mask) != 0
}

// SetK105 K105
func (r *registerK2lrType) SetK105(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK105Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK105Mask)
	}
}

const (
	RegisterK2lrFieldK106Shift = 10
	RegisterK2lrFieldK106Mask  = 0x400
)

// GetK106 K106
func (r *registerK2lrType) GetK106() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK106Mask) != 0
}

// SetK106 K106
func (r *registerK2lrType) SetK106(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK106Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK106Mask)
	}
}

const (
	RegisterK2lrFieldK107Shift = 11
	RegisterK2lrFieldK107Mask  = 0x800
)

// GetK107 K107
func (r *registerK2lrType) GetK107() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK107Mask) != 0
}

// SetK107 K107
func (r *registerK2lrType) SetK107(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK107Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK107Mask)
	}
}

const (
	RegisterK2lrFieldK108Shift = 12
	RegisterK2lrFieldK108Mask  = 0x1000
)

// GetK108 K108
func (r *registerK2lrType) GetK108() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK108Mask) != 0
}

// SetK108 K108
func (r *registerK2lrType) SetK108(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK108Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK108Mask)
	}
}

const (
	RegisterK2lrFieldK109Shift = 13
	RegisterK2lrFieldK109Mask  = 0x2000
)

// GetK109 K109
func (r *registerK2lrType) GetK109() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK109Mask) != 0
}

// SetK109 K109
func (r *registerK2lrType) SetK109(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK109Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK109Mask)
	}
}

const (
	RegisterK2lrFieldK110Shift = 14
	RegisterK2lrFieldK110Mask  = 0x4000
)

// GetK110 K110
func (r *registerK2lrType) GetK110() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK110Mask) != 0
}

// SetK110 K110
func (r *registerK2lrType) SetK110(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK110Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK110Mask)
	}
}

const (
	RegisterK2lrFieldK111Shift = 15
	RegisterK2lrFieldK111Mask  = 0x8000
)

// GetK111 K111
func (r *registerK2lrType) GetK111() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK111Mask) != 0
}

// SetK111 K111
func (r *registerK2lrType) SetK111(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK111Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK111Mask)
	}
}

const (
	RegisterK2lrFieldK112Shift = 16
	RegisterK2lrFieldK112Mask  = 0x10000
)

// GetK112 K112
func (r *registerK2lrType) GetK112() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK112Mask) != 0
}

// SetK112 K112
func (r *registerK2lrType) SetK112(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK112Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK112Mask)
	}
}

const (
	RegisterK2lrFieldK113Shift = 17
	RegisterK2lrFieldK113Mask  = 0x20000
)

// GetK113 K113
func (r *registerK2lrType) GetK113() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK113Mask) != 0
}

// SetK113 K113
func (r *registerK2lrType) SetK113(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK113Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK113Mask)
	}
}

const (
	RegisterK2lrFieldK114Shift = 18
	RegisterK2lrFieldK114Mask  = 0x40000
)

// GetK114 K114
func (r *registerK2lrType) GetK114() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK114Mask) != 0
}

// SetK114 K114
func (r *registerK2lrType) SetK114(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK114Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK114Mask)
	}
}

const (
	RegisterK2lrFieldK115Shift = 19
	RegisterK2lrFieldK115Mask  = 0x80000
)

// GetK115 K115
func (r *registerK2lrType) GetK115() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK115Mask) != 0
}

// SetK115 K115
func (r *registerK2lrType) SetK115(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK115Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK115Mask)
	}
}

const (
	RegisterK2lrFieldK116Shift = 20
	RegisterK2lrFieldK116Mask  = 0x100000
)

// GetK116 K116
func (r *registerK2lrType) GetK116() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK116Mask) != 0
}

// SetK116 K116
func (r *registerK2lrType) SetK116(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK116Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK116Mask)
	}
}

const (
	RegisterK2lrFieldK117Shift = 21
	RegisterK2lrFieldK117Mask  = 0x200000
)

// GetK117 K117
func (r *registerK2lrType) GetK117() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK117Mask) != 0
}

// SetK117 K117
func (r *registerK2lrType) SetK117(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK117Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK117Mask)
	}
}

const (
	RegisterK2lrFieldK118Shift = 22
	RegisterK2lrFieldK118Mask  = 0x400000
)

// GetK118 K118
func (r *registerK2lrType) GetK118() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK118Mask) != 0
}

// SetK118 K118
func (r *registerK2lrType) SetK118(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK118Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK118Mask)
	}
}

const (
	RegisterK2lrFieldK119Shift = 23
	RegisterK2lrFieldK119Mask  = 0x800000
)

// GetK119 K119
func (r *registerK2lrType) GetK119() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK119Mask) != 0
}

// SetK119 K119
func (r *registerK2lrType) SetK119(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK119Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK119Mask)
	}
}

const (
	RegisterK2lrFieldK120Shift = 24
	RegisterK2lrFieldK120Mask  = 0x1000000
)

// GetK120 K120
func (r *registerK2lrType) GetK120() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK120Mask) != 0
}

// SetK120 K120
func (r *registerK2lrType) SetK120(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK120Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK120Mask)
	}
}

const (
	RegisterK2lrFieldK121Shift = 25
	RegisterK2lrFieldK121Mask  = 0x2000000
)

// GetK121 K121
func (r *registerK2lrType) GetK121() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK121Mask) != 0
}

// SetK121 K121
func (r *registerK2lrType) SetK121(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK121Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK121Mask)
	}
}

const (
	RegisterK2lrFieldK122Shift = 26
	RegisterK2lrFieldK122Mask  = 0x4000000
)

// GetK122 K122
func (r *registerK2lrType) GetK122() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK122Mask) != 0
}

// SetK122 K122
func (r *registerK2lrType) SetK122(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK122Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK122Mask)
	}
}

const (
	RegisterK2lrFieldK123Shift = 27
	RegisterK2lrFieldK123Mask  = 0x8000000
)

// GetK123 K123
func (r *registerK2lrType) GetK123() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK123Mask) != 0
}

// SetK123 K123
func (r *registerK2lrType) SetK123(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK123Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK123Mask)
	}
}

const (
	RegisterK2lrFieldK124Shift = 28
	RegisterK2lrFieldK124Mask  = 0x10000000
)

// GetK124 K124
func (r *registerK2lrType) GetK124() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK124Mask) != 0
}

// SetK124 K124
func (r *registerK2lrType) SetK124(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK124Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK124Mask)
	}
}

const (
	RegisterK2lrFieldK125Shift = 29
	RegisterK2lrFieldK125Mask  = 0x20000000
)

// GetK125 K125
func (r *registerK2lrType) GetK125() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK125Mask) != 0
}

// SetK125 K125
func (r *registerK2lrType) SetK125(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK125Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK125Mask)
	}
}

const (
	RegisterK2lrFieldK126Shift = 30
	RegisterK2lrFieldK126Mask  = 0x40000000
)

// GetK126 K126
func (r *registerK2lrType) GetK126() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK126Mask) != 0
}

// SetK126 K126
func (r *registerK2lrType) SetK126(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK126Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK126Mask)
	}
}

const (
	RegisterK2lrFieldK127Shift = 31
	RegisterK2lrFieldK127Mask  = 0x80000000
)

// GetK127 K127
func (r *registerK2lrType) GetK127() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldK127Mask) != 0
}

// SetK127 K127
func (r *registerK2lrType) SetK127(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldK127Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldK127Mask)
	}
}

// registerK2rrType key registers
type registerK2rrType uint32

const (
	RegisterK2rrFieldK64Shift = 0
	RegisterK2rrFieldK64Mask  = 0x1
)

// GetK64 K64
func (r *registerK2rrType) GetK64() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK64Mask) != 0
}

// SetK64 K64
func (r *registerK2rrType) SetK64(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK64Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK64Mask)
	}
}

const (
	RegisterK2rrFieldK65Shift = 1
	RegisterK2rrFieldK65Mask  = 0x2
)

// GetK65 K65
func (r *registerK2rrType) GetK65() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK65Mask) != 0
}

// SetK65 K65
func (r *registerK2rrType) SetK65(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK65Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK65Mask)
	}
}

const (
	RegisterK2rrFieldK66Shift = 2
	RegisterK2rrFieldK66Mask  = 0x4
)

// GetK66 K66
func (r *registerK2rrType) GetK66() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK66Mask) != 0
}

// SetK66 K66
func (r *registerK2rrType) SetK66(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK66Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK66Mask)
	}
}

const (
	RegisterK2rrFieldK67Shift = 3
	RegisterK2rrFieldK67Mask  = 0x8
)

// GetK67 K67
func (r *registerK2rrType) GetK67() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK67Mask) != 0
}

// SetK67 K67
func (r *registerK2rrType) SetK67(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK67Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK67Mask)
	}
}

const (
	RegisterK2rrFieldK68Shift = 4
	RegisterK2rrFieldK68Mask  = 0x10
)

// GetK68 K68
func (r *registerK2rrType) GetK68() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK68Mask) != 0
}

// SetK68 K68
func (r *registerK2rrType) SetK68(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK68Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK68Mask)
	}
}

const (
	RegisterK2rrFieldK69Shift = 5
	RegisterK2rrFieldK69Mask  = 0x20
)

// GetK69 K69
func (r *registerK2rrType) GetK69() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK69Mask) != 0
}

// SetK69 K69
func (r *registerK2rrType) SetK69(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK69Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK69Mask)
	}
}

const (
	RegisterK2rrFieldK70Shift = 6
	RegisterK2rrFieldK70Mask  = 0x40
)

// GetK70 K70
func (r *registerK2rrType) GetK70() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK70Mask) != 0
}

// SetK70 K70
func (r *registerK2rrType) SetK70(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK70Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK70Mask)
	}
}

const (
	RegisterK2rrFieldK71Shift = 7
	RegisterK2rrFieldK71Mask  = 0x80
)

// GetK71 K71
func (r *registerK2rrType) GetK71() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK71Mask) != 0
}

// SetK71 K71
func (r *registerK2rrType) SetK71(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK71Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK71Mask)
	}
}

const (
	RegisterK2rrFieldK72Shift = 8
	RegisterK2rrFieldK72Mask  = 0x100
)

// GetK72 K72
func (r *registerK2rrType) GetK72() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK72Mask) != 0
}

// SetK72 K72
func (r *registerK2rrType) SetK72(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK72Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK72Mask)
	}
}

const (
	RegisterK2rrFieldK73Shift = 9
	RegisterK2rrFieldK73Mask  = 0x200
)

// GetK73 K73
func (r *registerK2rrType) GetK73() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK73Mask) != 0
}

// SetK73 K73
func (r *registerK2rrType) SetK73(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK73Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK73Mask)
	}
}

const (
	RegisterK2rrFieldK74Shift = 10
	RegisterK2rrFieldK74Mask  = 0x400
)

// GetK74 K74
func (r *registerK2rrType) GetK74() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK74Mask) != 0
}

// SetK74 K74
func (r *registerK2rrType) SetK74(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK74Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK74Mask)
	}
}

const (
	RegisterK2rrFieldK75Shift = 11
	RegisterK2rrFieldK75Mask  = 0x800
)

// GetK75 K75
func (r *registerK2rrType) GetK75() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK75Mask) != 0
}

// SetK75 K75
func (r *registerK2rrType) SetK75(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK75Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK75Mask)
	}
}

const (
	RegisterK2rrFieldK76Shift = 12
	RegisterK2rrFieldK76Mask  = 0x1000
)

// GetK76 K76
func (r *registerK2rrType) GetK76() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK76Mask) != 0
}

// SetK76 K76
func (r *registerK2rrType) SetK76(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK76Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK76Mask)
	}
}

const (
	RegisterK2rrFieldK77Shift = 13
	RegisterK2rrFieldK77Mask  = 0x2000
)

// GetK77 K77
func (r *registerK2rrType) GetK77() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK77Mask) != 0
}

// SetK77 K77
func (r *registerK2rrType) SetK77(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK77Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK77Mask)
	}
}

const (
	RegisterK2rrFieldK78Shift = 14
	RegisterK2rrFieldK78Mask  = 0x4000
)

// GetK78 K78
func (r *registerK2rrType) GetK78() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK78Mask) != 0
}

// SetK78 K78
func (r *registerK2rrType) SetK78(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK78Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK78Mask)
	}
}

const (
	RegisterK2rrFieldK79Shift = 15
	RegisterK2rrFieldK79Mask  = 0x8000
)

// GetK79 K79
func (r *registerK2rrType) GetK79() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK79Mask) != 0
}

// SetK79 K79
func (r *registerK2rrType) SetK79(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK79Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK79Mask)
	}
}

const (
	RegisterK2rrFieldK80Shift = 16
	RegisterK2rrFieldK80Mask  = 0x10000
)

// GetK80 K80
func (r *registerK2rrType) GetK80() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK80Mask) != 0
}

// SetK80 K80
func (r *registerK2rrType) SetK80(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK80Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK80Mask)
	}
}

const (
	RegisterK2rrFieldK81Shift = 17
	RegisterK2rrFieldK81Mask  = 0x20000
)

// GetK81 K81
func (r *registerK2rrType) GetK81() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK81Mask) != 0
}

// SetK81 K81
func (r *registerK2rrType) SetK81(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK81Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK81Mask)
	}
}

const (
	RegisterK2rrFieldK82Shift = 18
	RegisterK2rrFieldK82Mask  = 0x40000
)

// GetK82 K82
func (r *registerK2rrType) GetK82() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK82Mask) != 0
}

// SetK82 K82
func (r *registerK2rrType) SetK82(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK82Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK82Mask)
	}
}

const (
	RegisterK2rrFieldK83Shift = 19
	RegisterK2rrFieldK83Mask  = 0x80000
)

// GetK83 K83
func (r *registerK2rrType) GetK83() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK83Mask) != 0
}

// SetK83 K83
func (r *registerK2rrType) SetK83(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK83Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK83Mask)
	}
}

const (
	RegisterK2rrFieldK84Shift = 20
	RegisterK2rrFieldK84Mask  = 0x100000
)

// GetK84 K84
func (r *registerK2rrType) GetK84() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK84Mask) != 0
}

// SetK84 K84
func (r *registerK2rrType) SetK84(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK84Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK84Mask)
	}
}

const (
	RegisterK2rrFieldK85Shift = 21
	RegisterK2rrFieldK85Mask  = 0x200000
)

// GetK85 K85
func (r *registerK2rrType) GetK85() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK85Mask) != 0
}

// SetK85 K85
func (r *registerK2rrType) SetK85(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK85Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK85Mask)
	}
}

const (
	RegisterK2rrFieldK86Shift = 22
	RegisterK2rrFieldK86Mask  = 0x400000
)

// GetK86 K86
func (r *registerK2rrType) GetK86() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK86Mask) != 0
}

// SetK86 K86
func (r *registerK2rrType) SetK86(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK86Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK86Mask)
	}
}

const (
	RegisterK2rrFieldK87Shift = 23
	RegisterK2rrFieldK87Mask  = 0x800000
)

// GetK87 K87
func (r *registerK2rrType) GetK87() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK87Mask) != 0
}

// SetK87 K87
func (r *registerK2rrType) SetK87(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK87Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK87Mask)
	}
}

const (
	RegisterK2rrFieldK88Shift = 24
	RegisterK2rrFieldK88Mask  = 0x1000000
)

// GetK88 K88
func (r *registerK2rrType) GetK88() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK88Mask) != 0
}

// SetK88 K88
func (r *registerK2rrType) SetK88(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK88Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK88Mask)
	}
}

const (
	RegisterK2rrFieldK89Shift = 25
	RegisterK2rrFieldK89Mask  = 0x2000000
)

// GetK89 K89
func (r *registerK2rrType) GetK89() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK89Mask) != 0
}

// SetK89 K89
func (r *registerK2rrType) SetK89(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK89Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK89Mask)
	}
}

const (
	RegisterK2rrFieldK90Shift = 26
	RegisterK2rrFieldK90Mask  = 0x4000000
)

// GetK90 K90
func (r *registerK2rrType) GetK90() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK90Mask) != 0
}

// SetK90 K90
func (r *registerK2rrType) SetK90(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK90Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK90Mask)
	}
}

const (
	RegisterK2rrFieldK91Shift = 27
	RegisterK2rrFieldK91Mask  = 0x8000000
)

// GetK91 K91
func (r *registerK2rrType) GetK91() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK91Mask) != 0
}

// SetK91 K91
func (r *registerK2rrType) SetK91(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK91Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK91Mask)
	}
}

const (
	RegisterK2rrFieldK92Shift = 28
	RegisterK2rrFieldK92Mask  = 0x10000000
)

// GetK92 K92
func (r *registerK2rrType) GetK92() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK92Mask) != 0
}

// SetK92 K92
func (r *registerK2rrType) SetK92(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK92Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK92Mask)
	}
}

const (
	RegisterK2rrFieldK93Shift = 29
	RegisterK2rrFieldK93Mask  = 0x20000000
)

// GetK93 K93
func (r *registerK2rrType) GetK93() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK93Mask) != 0
}

// SetK93 K93
func (r *registerK2rrType) SetK93(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK93Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK93Mask)
	}
}

const (
	RegisterK2rrFieldK94Shift = 30
	RegisterK2rrFieldK94Mask  = 0x40000000
)

// GetK94 K94
func (r *registerK2rrType) GetK94() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK94Mask) != 0
}

// SetK94 K94
func (r *registerK2rrType) SetK94(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK94Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK94Mask)
	}
}

const (
	RegisterK2rrFieldK95Shift = 31
	RegisterK2rrFieldK95Mask  = 0x80000000
)

// GetK95 K95
func (r *registerK2rrType) GetK95() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldK95Mask) != 0
}

// SetK95 K95
func (r *registerK2rrType) SetK95(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldK95Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldK95Mask)
	}
}

// registerK3lrType key registers
type registerK3lrType uint32

const (
	RegisterK3lrFieldK32Shift = 0
	RegisterK3lrFieldK32Mask  = 0x1
)

// GetK32 K32
func (r *registerK3lrType) GetK32() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK32Mask) != 0
}

// SetK32 K32
func (r *registerK3lrType) SetK32(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK32Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK32Mask)
	}
}

const (
	RegisterK3lrFieldK33Shift = 1
	RegisterK3lrFieldK33Mask  = 0x2
)

// GetK33 K33
func (r *registerK3lrType) GetK33() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK33Mask) != 0
}

// SetK33 K33
func (r *registerK3lrType) SetK33(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK33Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK33Mask)
	}
}

const (
	RegisterK3lrFieldK34Shift = 2
	RegisterK3lrFieldK34Mask  = 0x4
)

// GetK34 K34
func (r *registerK3lrType) GetK34() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK34Mask) != 0
}

// SetK34 K34
func (r *registerK3lrType) SetK34(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK34Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK34Mask)
	}
}

const (
	RegisterK3lrFieldK35Shift = 3
	RegisterK3lrFieldK35Mask  = 0x8
)

// GetK35 K35
func (r *registerK3lrType) GetK35() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK35Mask) != 0
}

// SetK35 K35
func (r *registerK3lrType) SetK35(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK35Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK35Mask)
	}
}

const (
	RegisterK3lrFieldK36Shift = 4
	RegisterK3lrFieldK36Mask  = 0x10
)

// GetK36 K36
func (r *registerK3lrType) GetK36() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK36Mask) != 0
}

// SetK36 K36
func (r *registerK3lrType) SetK36(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK36Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK36Mask)
	}
}

const (
	RegisterK3lrFieldK37Shift = 5
	RegisterK3lrFieldK37Mask  = 0x20
)

// GetK37 K37
func (r *registerK3lrType) GetK37() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK37Mask) != 0
}

// SetK37 K37
func (r *registerK3lrType) SetK37(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK37Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK37Mask)
	}
}

const (
	RegisterK3lrFieldK38Shift = 6
	RegisterK3lrFieldK38Mask  = 0x40
)

// GetK38 K38
func (r *registerK3lrType) GetK38() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK38Mask) != 0
}

// SetK38 K38
func (r *registerK3lrType) SetK38(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK38Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK38Mask)
	}
}

const (
	RegisterK3lrFieldK39Shift = 7
	RegisterK3lrFieldK39Mask  = 0x80
)

// GetK39 K39
func (r *registerK3lrType) GetK39() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK39Mask) != 0
}

// SetK39 K39
func (r *registerK3lrType) SetK39(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK39Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK39Mask)
	}
}

const (
	RegisterK3lrFieldK40Shift = 8
	RegisterK3lrFieldK40Mask  = 0x100
)

// GetK40 K40
func (r *registerK3lrType) GetK40() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK40Mask) != 0
}

// SetK40 K40
func (r *registerK3lrType) SetK40(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK40Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK40Mask)
	}
}

const (
	RegisterK3lrFieldK41Shift = 9
	RegisterK3lrFieldK41Mask  = 0x200
)

// GetK41 K41
func (r *registerK3lrType) GetK41() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK41Mask) != 0
}

// SetK41 K41
func (r *registerK3lrType) SetK41(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK41Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK41Mask)
	}
}

const (
	RegisterK3lrFieldK42Shift = 10
	RegisterK3lrFieldK42Mask  = 0x400
)

// GetK42 K42
func (r *registerK3lrType) GetK42() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK42Mask) != 0
}

// SetK42 K42
func (r *registerK3lrType) SetK42(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK42Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK42Mask)
	}
}

const (
	RegisterK3lrFieldK43Shift = 11
	RegisterK3lrFieldK43Mask  = 0x800
)

// GetK43 K43
func (r *registerK3lrType) GetK43() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK43Mask) != 0
}

// SetK43 K43
func (r *registerK3lrType) SetK43(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK43Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK43Mask)
	}
}

const (
	RegisterK3lrFieldK44Shift = 12
	RegisterK3lrFieldK44Mask  = 0x1000
)

// GetK44 K44
func (r *registerK3lrType) GetK44() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK44Mask) != 0
}

// SetK44 K44
func (r *registerK3lrType) SetK44(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK44Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK44Mask)
	}
}

const (
	RegisterK3lrFieldK45Shift = 13
	RegisterK3lrFieldK45Mask  = 0x2000
)

// GetK45 K45
func (r *registerK3lrType) GetK45() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK45Mask) != 0
}

// SetK45 K45
func (r *registerK3lrType) SetK45(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK45Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK45Mask)
	}
}

const (
	RegisterK3lrFieldK46Shift = 14
	RegisterK3lrFieldK46Mask  = 0x4000
)

// GetK46 K46
func (r *registerK3lrType) GetK46() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK46Mask) != 0
}

// SetK46 K46
func (r *registerK3lrType) SetK46(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK46Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK46Mask)
	}
}

const (
	RegisterK3lrFieldK47Shift = 15
	RegisterK3lrFieldK47Mask  = 0x8000
)

// GetK47 K47
func (r *registerK3lrType) GetK47() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK47Mask) != 0
}

// SetK47 K47
func (r *registerK3lrType) SetK47(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK47Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK47Mask)
	}
}

const (
	RegisterK3lrFieldK48Shift = 16
	RegisterK3lrFieldK48Mask  = 0x10000
)

// GetK48 K48
func (r *registerK3lrType) GetK48() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK48Mask) != 0
}

// SetK48 K48
func (r *registerK3lrType) SetK48(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK48Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK48Mask)
	}
}

const (
	RegisterK3lrFieldK49Shift = 17
	RegisterK3lrFieldK49Mask  = 0x20000
)

// GetK49 K49
func (r *registerK3lrType) GetK49() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK49Mask) != 0
}

// SetK49 K49
func (r *registerK3lrType) SetK49(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK49Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK49Mask)
	}
}

const (
	RegisterK3lrFieldK50Shift = 18
	RegisterK3lrFieldK50Mask  = 0x40000
)

// GetK50 K50
func (r *registerK3lrType) GetK50() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK50Mask) != 0
}

// SetK50 K50
func (r *registerK3lrType) SetK50(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK50Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK50Mask)
	}
}

const (
	RegisterK3lrFieldK51Shift = 19
	RegisterK3lrFieldK51Mask  = 0x80000
)

// GetK51 K51
func (r *registerK3lrType) GetK51() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK51Mask) != 0
}

// SetK51 K51
func (r *registerK3lrType) SetK51(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK51Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK51Mask)
	}
}

const (
	RegisterK3lrFieldK52Shift = 20
	RegisterK3lrFieldK52Mask  = 0x100000
)

// GetK52 K52
func (r *registerK3lrType) GetK52() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK52Mask) != 0
}

// SetK52 K52
func (r *registerK3lrType) SetK52(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK52Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK52Mask)
	}
}

const (
	RegisterK3lrFieldK53Shift = 21
	RegisterK3lrFieldK53Mask  = 0x200000
)

// GetK53 K53
func (r *registerK3lrType) GetK53() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK53Mask) != 0
}

// SetK53 K53
func (r *registerK3lrType) SetK53(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK53Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK53Mask)
	}
}

const (
	RegisterK3lrFieldK54Shift = 22
	RegisterK3lrFieldK54Mask  = 0x400000
)

// GetK54 K54
func (r *registerK3lrType) GetK54() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK54Mask) != 0
}

// SetK54 K54
func (r *registerK3lrType) SetK54(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK54Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK54Mask)
	}
}

const (
	RegisterK3lrFieldK55Shift = 23
	RegisterK3lrFieldK55Mask  = 0x800000
)

// GetK55 K55
func (r *registerK3lrType) GetK55() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK55Mask) != 0
}

// SetK55 K55
func (r *registerK3lrType) SetK55(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK55Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK55Mask)
	}
}

const (
	RegisterK3lrFieldK56Shift = 24
	RegisterK3lrFieldK56Mask  = 0x1000000
)

// GetK56 K56
func (r *registerK3lrType) GetK56() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK56Mask) != 0
}

// SetK56 K56
func (r *registerK3lrType) SetK56(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK56Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK56Mask)
	}
}

const (
	RegisterK3lrFieldK57Shift = 25
	RegisterK3lrFieldK57Mask  = 0x2000000
)

// GetK57 K57
func (r *registerK3lrType) GetK57() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK57Mask) != 0
}

// SetK57 K57
func (r *registerK3lrType) SetK57(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK57Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK57Mask)
	}
}

const (
	RegisterK3lrFieldK58Shift = 26
	RegisterK3lrFieldK58Mask  = 0x4000000
)

// GetK58 K58
func (r *registerK3lrType) GetK58() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK58Mask) != 0
}

// SetK58 K58
func (r *registerK3lrType) SetK58(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK58Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK58Mask)
	}
}

const (
	RegisterK3lrFieldK59Shift = 27
	RegisterK3lrFieldK59Mask  = 0x8000000
)

// GetK59 K59
func (r *registerK3lrType) GetK59() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK59Mask) != 0
}

// SetK59 K59
func (r *registerK3lrType) SetK59(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK59Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK59Mask)
	}
}

const (
	RegisterK3lrFieldK60Shift = 28
	RegisterK3lrFieldK60Mask  = 0x10000000
)

// GetK60 K60
func (r *registerK3lrType) GetK60() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK60Mask) != 0
}

// SetK60 K60
func (r *registerK3lrType) SetK60(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK60Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK60Mask)
	}
}

const (
	RegisterK3lrFieldK61Shift = 29
	RegisterK3lrFieldK61Mask  = 0x20000000
)

// GetK61 K61
func (r *registerK3lrType) GetK61() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK61Mask) != 0
}

// SetK61 K61
func (r *registerK3lrType) SetK61(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK61Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK61Mask)
	}
}

const (
	RegisterK3lrFieldK62Shift = 30
	RegisterK3lrFieldK62Mask  = 0x40000000
)

// GetK62 K62
func (r *registerK3lrType) GetK62() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK62Mask) != 0
}

// SetK62 K62
func (r *registerK3lrType) SetK62(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK62Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK62Mask)
	}
}

const (
	RegisterK3lrFieldK63Shift = 31
	RegisterK3lrFieldK63Mask  = 0x80000000
)

// GetK63 K63
func (r *registerK3lrType) GetK63() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldK63Mask) != 0
}

// SetK63 K63
func (r *registerK3lrType) SetK63(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldK63Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldK63Mask)
	}
}

// registerK3rrType key registers
type registerK3rrType uint32

const (
	RegisterK3rrFieldK0Shift = 0
	RegisterK3rrFieldK0Mask  = 0x1
)

// GetK0 K0
func (r *registerK3rrType) GetK0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK0Mask) != 0
}

// SetK0 K0
func (r *registerK3rrType) SetK0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK0Mask)
	}
}

const (
	RegisterK3rrFieldK1Shift = 1
	RegisterK3rrFieldK1Mask  = 0x2
)

// GetK1 K1
func (r *registerK3rrType) GetK1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK1Mask) != 0
}

// SetK1 K1
func (r *registerK3rrType) SetK1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK1Mask)
	}
}

const (
	RegisterK3rrFieldK2Shift = 2
	RegisterK3rrFieldK2Mask  = 0x4
)

// GetK2 K2
func (r *registerK3rrType) GetK2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK2Mask) != 0
}

// SetK2 K2
func (r *registerK3rrType) SetK2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK2Mask)
	}
}

const (
	RegisterK3rrFieldK3Shift = 3
	RegisterK3rrFieldK3Mask  = 0x8
)

// GetK3 K3
func (r *registerK3rrType) GetK3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK3Mask) != 0
}

// SetK3 K3
func (r *registerK3rrType) SetK3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK3Mask)
	}
}

const (
	RegisterK3rrFieldK4Shift = 4
	RegisterK3rrFieldK4Mask  = 0x10
)

// GetK4 K4
func (r *registerK3rrType) GetK4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK4Mask) != 0
}

// SetK4 K4
func (r *registerK3rrType) SetK4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK4Mask)
	}
}

const (
	RegisterK3rrFieldK5Shift = 5
	RegisterK3rrFieldK5Mask  = 0x20
)

// GetK5 K5
func (r *registerK3rrType) GetK5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK5Mask) != 0
}

// SetK5 K5
func (r *registerK3rrType) SetK5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK5Mask)
	}
}

const (
	RegisterK3rrFieldK6Shift = 6
	RegisterK3rrFieldK6Mask  = 0x40
)

// GetK6 K6
func (r *registerK3rrType) GetK6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK6Mask) != 0
}

// SetK6 K6
func (r *registerK3rrType) SetK6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK6Mask)
	}
}

const (
	RegisterK3rrFieldK7Shift = 7
	RegisterK3rrFieldK7Mask  = 0x80
)

// GetK7 K7
func (r *registerK3rrType) GetK7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK7Mask) != 0
}

// SetK7 K7
func (r *registerK3rrType) SetK7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK7Mask)
	}
}

const (
	RegisterK3rrFieldK8Shift = 8
	RegisterK3rrFieldK8Mask  = 0x100
)

// GetK8 K8
func (r *registerK3rrType) GetK8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK8Mask) != 0
}

// SetK8 K8
func (r *registerK3rrType) SetK8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK8Mask)
	}
}

const (
	RegisterK3rrFieldK9Shift = 9
	RegisterK3rrFieldK9Mask  = 0x200
)

// GetK9 K9
func (r *registerK3rrType) GetK9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK9Mask) != 0
}

// SetK9 K9
func (r *registerK3rrType) SetK9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK9Mask)
	}
}

const (
	RegisterK3rrFieldK10Shift = 10
	RegisterK3rrFieldK10Mask  = 0x400
)

// GetK10 K10
func (r *registerK3rrType) GetK10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK10Mask) != 0
}

// SetK10 K10
func (r *registerK3rrType) SetK10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK10Mask)
	}
}

const (
	RegisterK3rrFieldK11Shift = 11
	RegisterK3rrFieldK11Mask  = 0x800
)

// GetK11 K11
func (r *registerK3rrType) GetK11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK11Mask) != 0
}

// SetK11 K11
func (r *registerK3rrType) SetK11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK11Mask)
	}
}

const (
	RegisterK3rrFieldK12Shift = 12
	RegisterK3rrFieldK12Mask  = 0x1000
)

// GetK12 K12
func (r *registerK3rrType) GetK12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK12Mask) != 0
}

// SetK12 K12
func (r *registerK3rrType) SetK12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK12Mask)
	}
}

const (
	RegisterK3rrFieldK13Shift = 13
	RegisterK3rrFieldK13Mask  = 0x2000
)

// GetK13 K13
func (r *registerK3rrType) GetK13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK13Mask) != 0
}

// SetK13 K13
func (r *registerK3rrType) SetK13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK13Mask)
	}
}

const (
	RegisterK3rrFieldK14Shift = 14
	RegisterK3rrFieldK14Mask  = 0x4000
)

// GetK14 K14
func (r *registerK3rrType) GetK14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK14Mask) != 0
}

// SetK14 K14
func (r *registerK3rrType) SetK14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK14Mask)
	}
}

const (
	RegisterK3rrFieldK15Shift = 15
	RegisterK3rrFieldK15Mask  = 0x8000
)

// GetK15 K15
func (r *registerK3rrType) GetK15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK15Mask) != 0
}

// SetK15 K15
func (r *registerK3rrType) SetK15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK15Mask)
	}
}

const (
	RegisterK3rrFieldK16Shift = 16
	RegisterK3rrFieldK16Mask  = 0x10000
)

// GetK16 K16
func (r *registerK3rrType) GetK16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK16Mask) != 0
}

// SetK16 K16
func (r *registerK3rrType) SetK16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK16Mask)
	}
}

const (
	RegisterK3rrFieldK17Shift = 17
	RegisterK3rrFieldK17Mask  = 0x20000
)

// GetK17 K17
func (r *registerK3rrType) GetK17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK17Mask) != 0
}

// SetK17 K17
func (r *registerK3rrType) SetK17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK17Mask)
	}
}

const (
	RegisterK3rrFieldK18Shift = 18
	RegisterK3rrFieldK18Mask  = 0x40000
)

// GetK18 K18
func (r *registerK3rrType) GetK18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK18Mask) != 0
}

// SetK18 K18
func (r *registerK3rrType) SetK18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK18Mask)
	}
}

const (
	RegisterK3rrFieldK19Shift = 19
	RegisterK3rrFieldK19Mask  = 0x80000
)

// GetK19 K19
func (r *registerK3rrType) GetK19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK19Mask) != 0
}

// SetK19 K19
func (r *registerK3rrType) SetK19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK19Mask)
	}
}

const (
	RegisterK3rrFieldK20Shift = 20
	RegisterK3rrFieldK20Mask  = 0x100000
)

// GetK20 K20
func (r *registerK3rrType) GetK20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK20Mask) != 0
}

// SetK20 K20
func (r *registerK3rrType) SetK20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK20Mask)
	}
}

const (
	RegisterK3rrFieldK21Shift = 21
	RegisterK3rrFieldK21Mask  = 0x200000
)

// GetK21 K21
func (r *registerK3rrType) GetK21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK21Mask) != 0
}

// SetK21 K21
func (r *registerK3rrType) SetK21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK21Mask)
	}
}

const (
	RegisterK3rrFieldK22Shift = 22
	RegisterK3rrFieldK22Mask  = 0x400000
)

// GetK22 K22
func (r *registerK3rrType) GetK22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK22Mask) != 0
}

// SetK22 K22
func (r *registerK3rrType) SetK22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK22Mask)
	}
}

const (
	RegisterK3rrFieldK23Shift = 23
	RegisterK3rrFieldK23Mask  = 0x800000
)

// GetK23 K23
func (r *registerK3rrType) GetK23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK23Mask) != 0
}

// SetK23 K23
func (r *registerK3rrType) SetK23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK23Mask)
	}
}

const (
	RegisterK3rrFieldK24Shift = 24
	RegisterK3rrFieldK24Mask  = 0x1000000
)

// GetK24 K24
func (r *registerK3rrType) GetK24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK24Mask) != 0
}

// SetK24 K24
func (r *registerK3rrType) SetK24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK24Mask)
	}
}

const (
	RegisterK3rrFieldK25Shift = 25
	RegisterK3rrFieldK25Mask  = 0x2000000
)

// GetK25 K25
func (r *registerK3rrType) GetK25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK25Mask) != 0
}

// SetK25 K25
func (r *registerK3rrType) SetK25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK25Mask)
	}
}

const (
	RegisterK3rrFieldK26Shift = 26
	RegisterK3rrFieldK26Mask  = 0x4000000
)

// GetK26 K26
func (r *registerK3rrType) GetK26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK26Mask) != 0
}

// SetK26 K26
func (r *registerK3rrType) SetK26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK26Mask)
	}
}

const (
	RegisterK3rrFieldK27Shift = 27
	RegisterK3rrFieldK27Mask  = 0x8000000
)

// GetK27 K27
func (r *registerK3rrType) GetK27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK27Mask) != 0
}

// SetK27 K27
func (r *registerK3rrType) SetK27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK27Mask)
	}
}

const (
	RegisterK3rrFieldK28Shift = 28
	RegisterK3rrFieldK28Mask  = 0x10000000
)

// GetK28 K28
func (r *registerK3rrType) GetK28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK28Mask) != 0
}

// SetK28 K28
func (r *registerK3rrType) SetK28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK28Mask)
	}
}

const (
	RegisterK3rrFieldK29Shift = 29
	RegisterK3rrFieldK29Mask  = 0x20000000
)

// GetK29 K29
func (r *registerK3rrType) GetK29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK29Mask) != 0
}

// SetK29 K29
func (r *registerK3rrType) SetK29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK29Mask)
	}
}

const (
	RegisterK3rrFieldK30Shift = 30
	RegisterK3rrFieldK30Mask  = 0x40000000
)

// GetK30 K30
func (r *registerK3rrType) GetK30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK30Mask) != 0
}

// SetK30 K30
func (r *registerK3rrType) SetK30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK30Mask)
	}
}

const (
	RegisterK3rrFieldK31Shift = 31
	RegisterK3rrFieldK31Mask  = 0x80000000
)

// GetK31 K31
func (r *registerK3rrType) GetK31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldK31Mask) != 0
}

// SetK31 K31
func (r *registerK3rrType) SetK31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldK31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldK31Mask)
	}
}

// registerIv0lrType initialization vector registers
type registerIv0lrType uint32

const (
	RegisterIv0lrFieldIv31Shift = 0
	RegisterIv0lrFieldIv31Mask  = 0x1
)

// GetIv31 IV31
func (r *registerIv0lrType) GetIv31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv31Mask) != 0
}

// SetIv31 IV31
func (r *registerIv0lrType) SetIv31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv31Mask)
	}
}

const (
	RegisterIv0lrFieldIv30Shift = 1
	RegisterIv0lrFieldIv30Mask  = 0x2
)

// GetIv30 IV30
func (r *registerIv0lrType) GetIv30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv30Mask) != 0
}

// SetIv30 IV30
func (r *registerIv0lrType) SetIv30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv30Mask)
	}
}

const (
	RegisterIv0lrFieldIv29Shift = 2
	RegisterIv0lrFieldIv29Mask  = 0x4
)

// GetIv29 IV29
func (r *registerIv0lrType) GetIv29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv29Mask) != 0
}

// SetIv29 IV29
func (r *registerIv0lrType) SetIv29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv29Mask)
	}
}

const (
	RegisterIv0lrFieldIv28Shift = 3
	RegisterIv0lrFieldIv28Mask  = 0x8
)

// GetIv28 IV28
func (r *registerIv0lrType) GetIv28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv28Mask) != 0
}

// SetIv28 IV28
func (r *registerIv0lrType) SetIv28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv28Mask)
	}
}

const (
	RegisterIv0lrFieldIv27Shift = 4
	RegisterIv0lrFieldIv27Mask  = 0x10
)

// GetIv27 IV27
func (r *registerIv0lrType) GetIv27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv27Mask) != 0
}

// SetIv27 IV27
func (r *registerIv0lrType) SetIv27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv27Mask)
	}
}

const (
	RegisterIv0lrFieldIv26Shift = 5
	RegisterIv0lrFieldIv26Mask  = 0x20
)

// GetIv26 IV26
func (r *registerIv0lrType) GetIv26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv26Mask) != 0
}

// SetIv26 IV26
func (r *registerIv0lrType) SetIv26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv26Mask)
	}
}

const (
	RegisterIv0lrFieldIv25Shift = 6
	RegisterIv0lrFieldIv25Mask  = 0x40
)

// GetIv25 IV25
func (r *registerIv0lrType) GetIv25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv25Mask) != 0
}

// SetIv25 IV25
func (r *registerIv0lrType) SetIv25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv25Mask)
	}
}

const (
	RegisterIv0lrFieldIv24Shift = 7
	RegisterIv0lrFieldIv24Mask  = 0x80
)

// GetIv24 IV24
func (r *registerIv0lrType) GetIv24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv24Mask) != 0
}

// SetIv24 IV24
func (r *registerIv0lrType) SetIv24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv24Mask)
	}
}

const (
	RegisterIv0lrFieldIv23Shift = 8
	RegisterIv0lrFieldIv23Mask  = 0x100
)

// GetIv23 IV23
func (r *registerIv0lrType) GetIv23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv23Mask) != 0
}

// SetIv23 IV23
func (r *registerIv0lrType) SetIv23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv23Mask)
	}
}

const (
	RegisterIv0lrFieldIv22Shift = 9
	RegisterIv0lrFieldIv22Mask  = 0x200
)

// GetIv22 IV22
func (r *registerIv0lrType) GetIv22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv22Mask) != 0
}

// SetIv22 IV22
func (r *registerIv0lrType) SetIv22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv22Mask)
	}
}

const (
	RegisterIv0lrFieldIv21Shift = 10
	RegisterIv0lrFieldIv21Mask  = 0x400
)

// GetIv21 IV21
func (r *registerIv0lrType) GetIv21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv21Mask) != 0
}

// SetIv21 IV21
func (r *registerIv0lrType) SetIv21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv21Mask)
	}
}

const (
	RegisterIv0lrFieldIv20Shift = 11
	RegisterIv0lrFieldIv20Mask  = 0x800
)

// GetIv20 IV20
func (r *registerIv0lrType) GetIv20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv20Mask) != 0
}

// SetIv20 IV20
func (r *registerIv0lrType) SetIv20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv20Mask)
	}
}

const (
	RegisterIv0lrFieldIv19Shift = 12
	RegisterIv0lrFieldIv19Mask  = 0x1000
)

// GetIv19 IV19
func (r *registerIv0lrType) GetIv19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv19Mask) != 0
}

// SetIv19 IV19
func (r *registerIv0lrType) SetIv19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv19Mask)
	}
}

const (
	RegisterIv0lrFieldIv18Shift = 13
	RegisterIv0lrFieldIv18Mask  = 0x2000
)

// GetIv18 IV18
func (r *registerIv0lrType) GetIv18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv18Mask) != 0
}

// SetIv18 IV18
func (r *registerIv0lrType) SetIv18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv18Mask)
	}
}

const (
	RegisterIv0lrFieldIv17Shift = 14
	RegisterIv0lrFieldIv17Mask  = 0x4000
)

// GetIv17 IV17
func (r *registerIv0lrType) GetIv17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv17Mask) != 0
}

// SetIv17 IV17
func (r *registerIv0lrType) SetIv17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv17Mask)
	}
}

const (
	RegisterIv0lrFieldIv16Shift = 15
	RegisterIv0lrFieldIv16Mask  = 0x8000
)

// GetIv16 IV16
func (r *registerIv0lrType) GetIv16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv16Mask) != 0
}

// SetIv16 IV16
func (r *registerIv0lrType) SetIv16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv16Mask)
	}
}

const (
	RegisterIv0lrFieldIv15Shift = 16
	RegisterIv0lrFieldIv15Mask  = 0x10000
)

// GetIv15 IV15
func (r *registerIv0lrType) GetIv15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv15Mask) != 0
}

// SetIv15 IV15
func (r *registerIv0lrType) SetIv15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv15Mask)
	}
}

const (
	RegisterIv0lrFieldIv14Shift = 17
	RegisterIv0lrFieldIv14Mask  = 0x20000
)

// GetIv14 IV14
func (r *registerIv0lrType) GetIv14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv14Mask) != 0
}

// SetIv14 IV14
func (r *registerIv0lrType) SetIv14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv14Mask)
	}
}

const (
	RegisterIv0lrFieldIv13Shift = 18
	RegisterIv0lrFieldIv13Mask  = 0x40000
)

// GetIv13 IV13
func (r *registerIv0lrType) GetIv13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv13Mask) != 0
}

// SetIv13 IV13
func (r *registerIv0lrType) SetIv13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv13Mask)
	}
}

const (
	RegisterIv0lrFieldIv12Shift = 19
	RegisterIv0lrFieldIv12Mask  = 0x80000
)

// GetIv12 IV12
func (r *registerIv0lrType) GetIv12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv12Mask) != 0
}

// SetIv12 IV12
func (r *registerIv0lrType) SetIv12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv12Mask)
	}
}

const (
	RegisterIv0lrFieldIv11Shift = 20
	RegisterIv0lrFieldIv11Mask  = 0x100000
)

// GetIv11 IV11
func (r *registerIv0lrType) GetIv11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv11Mask) != 0
}

// SetIv11 IV11
func (r *registerIv0lrType) SetIv11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv11Mask)
	}
}

const (
	RegisterIv0lrFieldIv10Shift = 21
	RegisterIv0lrFieldIv10Mask  = 0x200000
)

// GetIv10 IV10
func (r *registerIv0lrType) GetIv10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv10Mask) != 0
}

// SetIv10 IV10
func (r *registerIv0lrType) SetIv10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv10Mask)
	}
}

const (
	RegisterIv0lrFieldIv9Shift = 22
	RegisterIv0lrFieldIv9Mask  = 0x400000
)

// GetIv9 IV9
func (r *registerIv0lrType) GetIv9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv9Mask) != 0
}

// SetIv9 IV9
func (r *registerIv0lrType) SetIv9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv9Mask)
	}
}

const (
	RegisterIv0lrFieldIv8Shift = 23
	RegisterIv0lrFieldIv8Mask  = 0x800000
)

// GetIv8 IV8
func (r *registerIv0lrType) GetIv8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv8Mask) != 0
}

// SetIv8 IV8
func (r *registerIv0lrType) SetIv8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv8Mask)
	}
}

const (
	RegisterIv0lrFieldIv7Shift = 24
	RegisterIv0lrFieldIv7Mask  = 0x1000000
)

// GetIv7 IV7
func (r *registerIv0lrType) GetIv7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv7Mask) != 0
}

// SetIv7 IV7
func (r *registerIv0lrType) SetIv7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv7Mask)
	}
}

const (
	RegisterIv0lrFieldIv6Shift = 25
	RegisterIv0lrFieldIv6Mask  = 0x2000000
)

// GetIv6 IV6
func (r *registerIv0lrType) GetIv6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv6Mask) != 0
}

// SetIv6 IV6
func (r *registerIv0lrType) SetIv6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv6Mask)
	}
}

const (
	RegisterIv0lrFieldIv5Shift = 26
	RegisterIv0lrFieldIv5Mask  = 0x4000000
)

// GetIv5 IV5
func (r *registerIv0lrType) GetIv5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv5Mask) != 0
}

// SetIv5 IV5
func (r *registerIv0lrType) SetIv5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv5Mask)
	}
}

const (
	RegisterIv0lrFieldIv4Shift = 27
	RegisterIv0lrFieldIv4Mask  = 0x8000000
)

// GetIv4 IV4
func (r *registerIv0lrType) GetIv4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv4Mask) != 0
}

// SetIv4 IV4
func (r *registerIv0lrType) SetIv4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv4Mask)
	}
}

const (
	RegisterIv0lrFieldIv3Shift = 28
	RegisterIv0lrFieldIv3Mask  = 0x10000000
)

// GetIv3 IV3
func (r *registerIv0lrType) GetIv3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv3Mask) != 0
}

// SetIv3 IV3
func (r *registerIv0lrType) SetIv3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv3Mask)
	}
}

const (
	RegisterIv0lrFieldIv2Shift = 29
	RegisterIv0lrFieldIv2Mask  = 0x20000000
)

// GetIv2 IV2
func (r *registerIv0lrType) GetIv2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv2Mask) != 0
}

// SetIv2 IV2
func (r *registerIv0lrType) SetIv2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv2Mask)
	}
}

const (
	RegisterIv0lrFieldIv1Shift = 30
	RegisterIv0lrFieldIv1Mask  = 0x40000000
)

// GetIv1 IV1
func (r *registerIv0lrType) GetIv1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv1Mask) != 0
}

// SetIv1 IV1
func (r *registerIv0lrType) SetIv1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv1Mask)
	}
}

const (
	RegisterIv0lrFieldIv0Shift = 31
	RegisterIv0lrFieldIv0Mask  = 0x80000000
)

// GetIv0 IV0
func (r *registerIv0lrType) GetIv0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0lrFieldIv0Mask) != 0
}

// SetIv0 IV0
func (r *registerIv0lrType) SetIv0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0lrFieldIv0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0lrFieldIv0Mask)
	}
}

// registerIv0rrType initialization vector registers
type registerIv0rrType uint32

const (
	RegisterIv0rrFieldIv63Shift = 0
	RegisterIv0rrFieldIv63Mask  = 0x1
)

// GetIv63 IV63
func (r *registerIv0rrType) GetIv63() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv63Mask) != 0
}

// SetIv63 IV63
func (r *registerIv0rrType) SetIv63(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv63Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv63Mask)
	}
}

const (
	RegisterIv0rrFieldIv62Shift = 1
	RegisterIv0rrFieldIv62Mask  = 0x2
)

// GetIv62 IV62
func (r *registerIv0rrType) GetIv62() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv62Mask) != 0
}

// SetIv62 IV62
func (r *registerIv0rrType) SetIv62(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv62Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv62Mask)
	}
}

const (
	RegisterIv0rrFieldIv61Shift = 2
	RegisterIv0rrFieldIv61Mask  = 0x4
)

// GetIv61 IV61
func (r *registerIv0rrType) GetIv61() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv61Mask) != 0
}

// SetIv61 IV61
func (r *registerIv0rrType) SetIv61(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv61Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv61Mask)
	}
}

const (
	RegisterIv0rrFieldIv60Shift = 3
	RegisterIv0rrFieldIv60Mask  = 0x8
)

// GetIv60 IV60
func (r *registerIv0rrType) GetIv60() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv60Mask) != 0
}

// SetIv60 IV60
func (r *registerIv0rrType) SetIv60(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv60Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv60Mask)
	}
}

const (
	RegisterIv0rrFieldIv59Shift = 4
	RegisterIv0rrFieldIv59Mask  = 0x10
)

// GetIv59 IV59
func (r *registerIv0rrType) GetIv59() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv59Mask) != 0
}

// SetIv59 IV59
func (r *registerIv0rrType) SetIv59(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv59Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv59Mask)
	}
}

const (
	RegisterIv0rrFieldIv58Shift = 5
	RegisterIv0rrFieldIv58Mask  = 0x20
)

// GetIv58 IV58
func (r *registerIv0rrType) GetIv58() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv58Mask) != 0
}

// SetIv58 IV58
func (r *registerIv0rrType) SetIv58(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv58Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv58Mask)
	}
}

const (
	RegisterIv0rrFieldIv57Shift = 6
	RegisterIv0rrFieldIv57Mask  = 0x40
)

// GetIv57 IV57
func (r *registerIv0rrType) GetIv57() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv57Mask) != 0
}

// SetIv57 IV57
func (r *registerIv0rrType) SetIv57(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv57Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv57Mask)
	}
}

const (
	RegisterIv0rrFieldIv56Shift = 7
	RegisterIv0rrFieldIv56Mask  = 0x80
)

// GetIv56 IV56
func (r *registerIv0rrType) GetIv56() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv56Mask) != 0
}

// SetIv56 IV56
func (r *registerIv0rrType) SetIv56(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv56Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv56Mask)
	}
}

const (
	RegisterIv0rrFieldIv55Shift = 8
	RegisterIv0rrFieldIv55Mask  = 0x100
)

// GetIv55 IV55
func (r *registerIv0rrType) GetIv55() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv55Mask) != 0
}

// SetIv55 IV55
func (r *registerIv0rrType) SetIv55(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv55Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv55Mask)
	}
}

const (
	RegisterIv0rrFieldIv54Shift = 9
	RegisterIv0rrFieldIv54Mask  = 0x200
)

// GetIv54 IV54
func (r *registerIv0rrType) GetIv54() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv54Mask) != 0
}

// SetIv54 IV54
func (r *registerIv0rrType) SetIv54(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv54Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv54Mask)
	}
}

const (
	RegisterIv0rrFieldIv53Shift = 10
	RegisterIv0rrFieldIv53Mask  = 0x400
)

// GetIv53 IV53
func (r *registerIv0rrType) GetIv53() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv53Mask) != 0
}

// SetIv53 IV53
func (r *registerIv0rrType) SetIv53(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv53Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv53Mask)
	}
}

const (
	RegisterIv0rrFieldIv52Shift = 11
	RegisterIv0rrFieldIv52Mask  = 0x800
)

// GetIv52 IV52
func (r *registerIv0rrType) GetIv52() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv52Mask) != 0
}

// SetIv52 IV52
func (r *registerIv0rrType) SetIv52(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv52Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv52Mask)
	}
}

const (
	RegisterIv0rrFieldIv51Shift = 12
	RegisterIv0rrFieldIv51Mask  = 0x1000
)

// GetIv51 IV51
func (r *registerIv0rrType) GetIv51() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv51Mask) != 0
}

// SetIv51 IV51
func (r *registerIv0rrType) SetIv51(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv51Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv51Mask)
	}
}

const (
	RegisterIv0rrFieldIv50Shift = 13
	RegisterIv0rrFieldIv50Mask  = 0x2000
)

// GetIv50 IV50
func (r *registerIv0rrType) GetIv50() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv50Mask) != 0
}

// SetIv50 IV50
func (r *registerIv0rrType) SetIv50(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv50Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv50Mask)
	}
}

const (
	RegisterIv0rrFieldIv49Shift = 14
	RegisterIv0rrFieldIv49Mask  = 0x4000
)

// GetIv49 IV49
func (r *registerIv0rrType) GetIv49() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv49Mask) != 0
}

// SetIv49 IV49
func (r *registerIv0rrType) SetIv49(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv49Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv49Mask)
	}
}

const (
	RegisterIv0rrFieldIv48Shift = 15
	RegisterIv0rrFieldIv48Mask  = 0x8000
)

// GetIv48 IV48
func (r *registerIv0rrType) GetIv48() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv48Mask) != 0
}

// SetIv48 IV48
func (r *registerIv0rrType) SetIv48(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv48Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv48Mask)
	}
}

const (
	RegisterIv0rrFieldIv47Shift = 16
	RegisterIv0rrFieldIv47Mask  = 0x10000
)

// GetIv47 IV47
func (r *registerIv0rrType) GetIv47() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv47Mask) != 0
}

// SetIv47 IV47
func (r *registerIv0rrType) SetIv47(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv47Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv47Mask)
	}
}

const (
	RegisterIv0rrFieldIv46Shift = 17
	RegisterIv0rrFieldIv46Mask  = 0x20000
)

// GetIv46 IV46
func (r *registerIv0rrType) GetIv46() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv46Mask) != 0
}

// SetIv46 IV46
func (r *registerIv0rrType) SetIv46(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv46Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv46Mask)
	}
}

const (
	RegisterIv0rrFieldIv45Shift = 18
	RegisterIv0rrFieldIv45Mask  = 0x40000
)

// GetIv45 IV45
func (r *registerIv0rrType) GetIv45() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv45Mask) != 0
}

// SetIv45 IV45
func (r *registerIv0rrType) SetIv45(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv45Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv45Mask)
	}
}

const (
	RegisterIv0rrFieldIv44Shift = 19
	RegisterIv0rrFieldIv44Mask  = 0x80000
)

// GetIv44 IV44
func (r *registerIv0rrType) GetIv44() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv44Mask) != 0
}

// SetIv44 IV44
func (r *registerIv0rrType) SetIv44(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv44Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv44Mask)
	}
}

const (
	RegisterIv0rrFieldIv43Shift = 20
	RegisterIv0rrFieldIv43Mask  = 0x100000
)

// GetIv43 IV43
func (r *registerIv0rrType) GetIv43() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv43Mask) != 0
}

// SetIv43 IV43
func (r *registerIv0rrType) SetIv43(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv43Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv43Mask)
	}
}

const (
	RegisterIv0rrFieldIv42Shift = 21
	RegisterIv0rrFieldIv42Mask  = 0x200000
)

// GetIv42 IV42
func (r *registerIv0rrType) GetIv42() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv42Mask) != 0
}

// SetIv42 IV42
func (r *registerIv0rrType) SetIv42(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv42Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv42Mask)
	}
}

const (
	RegisterIv0rrFieldIv41Shift = 22
	RegisterIv0rrFieldIv41Mask  = 0x400000
)

// GetIv41 IV41
func (r *registerIv0rrType) GetIv41() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv41Mask) != 0
}

// SetIv41 IV41
func (r *registerIv0rrType) SetIv41(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv41Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv41Mask)
	}
}

const (
	RegisterIv0rrFieldIv40Shift = 23
	RegisterIv0rrFieldIv40Mask  = 0x800000
)

// GetIv40 IV40
func (r *registerIv0rrType) GetIv40() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv40Mask) != 0
}

// SetIv40 IV40
func (r *registerIv0rrType) SetIv40(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv40Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv40Mask)
	}
}

const (
	RegisterIv0rrFieldIv39Shift = 24
	RegisterIv0rrFieldIv39Mask  = 0x1000000
)

// GetIv39 IV39
func (r *registerIv0rrType) GetIv39() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv39Mask) != 0
}

// SetIv39 IV39
func (r *registerIv0rrType) SetIv39(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv39Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv39Mask)
	}
}

const (
	RegisterIv0rrFieldIv38Shift = 25
	RegisterIv0rrFieldIv38Mask  = 0x2000000
)

// GetIv38 IV38
func (r *registerIv0rrType) GetIv38() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv38Mask) != 0
}

// SetIv38 IV38
func (r *registerIv0rrType) SetIv38(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv38Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv38Mask)
	}
}

const (
	RegisterIv0rrFieldIv37Shift = 26
	RegisterIv0rrFieldIv37Mask  = 0x4000000
)

// GetIv37 IV37
func (r *registerIv0rrType) GetIv37() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv37Mask) != 0
}

// SetIv37 IV37
func (r *registerIv0rrType) SetIv37(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv37Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv37Mask)
	}
}

const (
	RegisterIv0rrFieldIv36Shift = 27
	RegisterIv0rrFieldIv36Mask  = 0x8000000
)

// GetIv36 IV36
func (r *registerIv0rrType) GetIv36() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv36Mask) != 0
}

// SetIv36 IV36
func (r *registerIv0rrType) SetIv36(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv36Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv36Mask)
	}
}

const (
	RegisterIv0rrFieldIv35Shift = 28
	RegisterIv0rrFieldIv35Mask  = 0x10000000
)

// GetIv35 IV35
func (r *registerIv0rrType) GetIv35() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv35Mask) != 0
}

// SetIv35 IV35
func (r *registerIv0rrType) SetIv35(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv35Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv35Mask)
	}
}

const (
	RegisterIv0rrFieldIv34Shift = 29
	RegisterIv0rrFieldIv34Mask  = 0x20000000
)

// GetIv34 IV34
func (r *registerIv0rrType) GetIv34() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv34Mask) != 0
}

// SetIv34 IV34
func (r *registerIv0rrType) SetIv34(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv34Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv34Mask)
	}
}

const (
	RegisterIv0rrFieldIv33Shift = 30
	RegisterIv0rrFieldIv33Mask  = 0x40000000
)

// GetIv33 IV33
func (r *registerIv0rrType) GetIv33() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv33Mask) != 0
}

// SetIv33 IV33
func (r *registerIv0rrType) SetIv33(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv33Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv33Mask)
	}
}

const (
	RegisterIv0rrFieldIv32Shift = 31
	RegisterIv0rrFieldIv32Mask  = 0x80000000
)

// GetIv32 IV32
func (r *registerIv0rrType) GetIv32() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv0rrFieldIv32Mask) != 0
}

// SetIv32 IV32
func (r *registerIv0rrType) SetIv32(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv0rrFieldIv32Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv0rrFieldIv32Mask)
	}
}

// registerIv1lrType initialization vector registers
type registerIv1lrType uint32

const (
	RegisterIv1lrFieldIv95Shift = 0
	RegisterIv1lrFieldIv95Mask  = 0x1
)

// GetIv95 IV95
func (r *registerIv1lrType) GetIv95() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv95Mask) != 0
}

// SetIv95 IV95
func (r *registerIv1lrType) SetIv95(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv95Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv95Mask)
	}
}

const (
	RegisterIv1lrFieldIv94Shift = 1
	RegisterIv1lrFieldIv94Mask  = 0x2
)

// GetIv94 IV94
func (r *registerIv1lrType) GetIv94() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv94Mask) != 0
}

// SetIv94 IV94
func (r *registerIv1lrType) SetIv94(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv94Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv94Mask)
	}
}

const (
	RegisterIv1lrFieldIv93Shift = 2
	RegisterIv1lrFieldIv93Mask  = 0x4
)

// GetIv93 IV93
func (r *registerIv1lrType) GetIv93() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv93Mask) != 0
}

// SetIv93 IV93
func (r *registerIv1lrType) SetIv93(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv93Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv93Mask)
	}
}

const (
	RegisterIv1lrFieldIv92Shift = 3
	RegisterIv1lrFieldIv92Mask  = 0x8
)

// GetIv92 IV92
func (r *registerIv1lrType) GetIv92() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv92Mask) != 0
}

// SetIv92 IV92
func (r *registerIv1lrType) SetIv92(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv92Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv92Mask)
	}
}

const (
	RegisterIv1lrFieldIv91Shift = 4
	RegisterIv1lrFieldIv91Mask  = 0x10
)

// GetIv91 IV91
func (r *registerIv1lrType) GetIv91() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv91Mask) != 0
}

// SetIv91 IV91
func (r *registerIv1lrType) SetIv91(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv91Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv91Mask)
	}
}

const (
	RegisterIv1lrFieldIv90Shift = 5
	RegisterIv1lrFieldIv90Mask  = 0x20
)

// GetIv90 IV90
func (r *registerIv1lrType) GetIv90() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv90Mask) != 0
}

// SetIv90 IV90
func (r *registerIv1lrType) SetIv90(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv90Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv90Mask)
	}
}

const (
	RegisterIv1lrFieldIv89Shift = 6
	RegisterIv1lrFieldIv89Mask  = 0x40
)

// GetIv89 IV89
func (r *registerIv1lrType) GetIv89() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv89Mask) != 0
}

// SetIv89 IV89
func (r *registerIv1lrType) SetIv89(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv89Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv89Mask)
	}
}

const (
	RegisterIv1lrFieldIv88Shift = 7
	RegisterIv1lrFieldIv88Mask  = 0x80
)

// GetIv88 IV88
func (r *registerIv1lrType) GetIv88() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv88Mask) != 0
}

// SetIv88 IV88
func (r *registerIv1lrType) SetIv88(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv88Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv88Mask)
	}
}

const (
	RegisterIv1lrFieldIv87Shift = 8
	RegisterIv1lrFieldIv87Mask  = 0x100
)

// GetIv87 IV87
func (r *registerIv1lrType) GetIv87() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv87Mask) != 0
}

// SetIv87 IV87
func (r *registerIv1lrType) SetIv87(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv87Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv87Mask)
	}
}

const (
	RegisterIv1lrFieldIv86Shift = 9
	RegisterIv1lrFieldIv86Mask  = 0x200
)

// GetIv86 IV86
func (r *registerIv1lrType) GetIv86() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv86Mask) != 0
}

// SetIv86 IV86
func (r *registerIv1lrType) SetIv86(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv86Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv86Mask)
	}
}

const (
	RegisterIv1lrFieldIv85Shift = 10
	RegisterIv1lrFieldIv85Mask  = 0x400
)

// GetIv85 IV85
func (r *registerIv1lrType) GetIv85() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv85Mask) != 0
}

// SetIv85 IV85
func (r *registerIv1lrType) SetIv85(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv85Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv85Mask)
	}
}

const (
	RegisterIv1lrFieldIv84Shift = 11
	RegisterIv1lrFieldIv84Mask  = 0x800
)

// GetIv84 IV84
func (r *registerIv1lrType) GetIv84() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv84Mask) != 0
}

// SetIv84 IV84
func (r *registerIv1lrType) SetIv84(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv84Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv84Mask)
	}
}

const (
	RegisterIv1lrFieldIv83Shift = 12
	RegisterIv1lrFieldIv83Mask  = 0x1000
)

// GetIv83 IV83
func (r *registerIv1lrType) GetIv83() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv83Mask) != 0
}

// SetIv83 IV83
func (r *registerIv1lrType) SetIv83(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv83Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv83Mask)
	}
}

const (
	RegisterIv1lrFieldIv82Shift = 13
	RegisterIv1lrFieldIv82Mask  = 0x2000
)

// GetIv82 IV82
func (r *registerIv1lrType) GetIv82() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv82Mask) != 0
}

// SetIv82 IV82
func (r *registerIv1lrType) SetIv82(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv82Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv82Mask)
	}
}

const (
	RegisterIv1lrFieldIv81Shift = 14
	RegisterIv1lrFieldIv81Mask  = 0x4000
)

// GetIv81 IV81
func (r *registerIv1lrType) GetIv81() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv81Mask) != 0
}

// SetIv81 IV81
func (r *registerIv1lrType) SetIv81(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv81Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv81Mask)
	}
}

const (
	RegisterIv1lrFieldIv80Shift = 15
	RegisterIv1lrFieldIv80Mask  = 0x8000
)

// GetIv80 IV80
func (r *registerIv1lrType) GetIv80() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv80Mask) != 0
}

// SetIv80 IV80
func (r *registerIv1lrType) SetIv80(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv80Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv80Mask)
	}
}

const (
	RegisterIv1lrFieldIv79Shift = 16
	RegisterIv1lrFieldIv79Mask  = 0x10000
)

// GetIv79 IV79
func (r *registerIv1lrType) GetIv79() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv79Mask) != 0
}

// SetIv79 IV79
func (r *registerIv1lrType) SetIv79(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv79Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv79Mask)
	}
}

const (
	RegisterIv1lrFieldIv78Shift = 17
	RegisterIv1lrFieldIv78Mask  = 0x20000
)

// GetIv78 IV78
func (r *registerIv1lrType) GetIv78() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv78Mask) != 0
}

// SetIv78 IV78
func (r *registerIv1lrType) SetIv78(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv78Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv78Mask)
	}
}

const (
	RegisterIv1lrFieldIv77Shift = 18
	RegisterIv1lrFieldIv77Mask  = 0x40000
)

// GetIv77 IV77
func (r *registerIv1lrType) GetIv77() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv77Mask) != 0
}

// SetIv77 IV77
func (r *registerIv1lrType) SetIv77(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv77Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv77Mask)
	}
}

const (
	RegisterIv1lrFieldIv76Shift = 19
	RegisterIv1lrFieldIv76Mask  = 0x80000
)

// GetIv76 IV76
func (r *registerIv1lrType) GetIv76() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv76Mask) != 0
}

// SetIv76 IV76
func (r *registerIv1lrType) SetIv76(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv76Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv76Mask)
	}
}

const (
	RegisterIv1lrFieldIv75Shift = 20
	RegisterIv1lrFieldIv75Mask  = 0x100000
)

// GetIv75 IV75
func (r *registerIv1lrType) GetIv75() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv75Mask) != 0
}

// SetIv75 IV75
func (r *registerIv1lrType) SetIv75(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv75Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv75Mask)
	}
}

const (
	RegisterIv1lrFieldIv74Shift = 21
	RegisterIv1lrFieldIv74Mask  = 0x200000
)

// GetIv74 IV74
func (r *registerIv1lrType) GetIv74() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv74Mask) != 0
}

// SetIv74 IV74
func (r *registerIv1lrType) SetIv74(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv74Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv74Mask)
	}
}

const (
	RegisterIv1lrFieldIv73Shift = 22
	RegisterIv1lrFieldIv73Mask  = 0x400000
)

// GetIv73 IV73
func (r *registerIv1lrType) GetIv73() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv73Mask) != 0
}

// SetIv73 IV73
func (r *registerIv1lrType) SetIv73(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv73Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv73Mask)
	}
}

const (
	RegisterIv1lrFieldIv72Shift = 23
	RegisterIv1lrFieldIv72Mask  = 0x800000
)

// GetIv72 IV72
func (r *registerIv1lrType) GetIv72() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv72Mask) != 0
}

// SetIv72 IV72
func (r *registerIv1lrType) SetIv72(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv72Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv72Mask)
	}
}

const (
	RegisterIv1lrFieldIv71Shift = 24
	RegisterIv1lrFieldIv71Mask  = 0x1000000
)

// GetIv71 IV71
func (r *registerIv1lrType) GetIv71() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv71Mask) != 0
}

// SetIv71 IV71
func (r *registerIv1lrType) SetIv71(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv71Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv71Mask)
	}
}

const (
	RegisterIv1lrFieldIv70Shift = 25
	RegisterIv1lrFieldIv70Mask  = 0x2000000
)

// GetIv70 IV70
func (r *registerIv1lrType) GetIv70() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv70Mask) != 0
}

// SetIv70 IV70
func (r *registerIv1lrType) SetIv70(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv70Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv70Mask)
	}
}

const (
	RegisterIv1lrFieldIv69Shift = 26
	RegisterIv1lrFieldIv69Mask  = 0x4000000
)

// GetIv69 IV69
func (r *registerIv1lrType) GetIv69() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv69Mask) != 0
}

// SetIv69 IV69
func (r *registerIv1lrType) SetIv69(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv69Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv69Mask)
	}
}

const (
	RegisterIv1lrFieldIv68Shift = 27
	RegisterIv1lrFieldIv68Mask  = 0x8000000
)

// GetIv68 IV68
func (r *registerIv1lrType) GetIv68() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv68Mask) != 0
}

// SetIv68 IV68
func (r *registerIv1lrType) SetIv68(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv68Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv68Mask)
	}
}

const (
	RegisterIv1lrFieldIv67Shift = 28
	RegisterIv1lrFieldIv67Mask  = 0x10000000
)

// GetIv67 IV67
func (r *registerIv1lrType) GetIv67() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv67Mask) != 0
}

// SetIv67 IV67
func (r *registerIv1lrType) SetIv67(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv67Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv67Mask)
	}
}

const (
	RegisterIv1lrFieldIv66Shift = 29
	RegisterIv1lrFieldIv66Mask  = 0x20000000
)

// GetIv66 IV66
func (r *registerIv1lrType) GetIv66() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv66Mask) != 0
}

// SetIv66 IV66
func (r *registerIv1lrType) SetIv66(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv66Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv66Mask)
	}
}

const (
	RegisterIv1lrFieldIv65Shift = 30
	RegisterIv1lrFieldIv65Mask  = 0x40000000
)

// GetIv65 IV65
func (r *registerIv1lrType) GetIv65() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv65Mask) != 0
}

// SetIv65 IV65
func (r *registerIv1lrType) SetIv65(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv65Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv65Mask)
	}
}

const (
	RegisterIv1lrFieldIv64Shift = 31
	RegisterIv1lrFieldIv64Mask  = 0x80000000
)

// GetIv64 IV64
func (r *registerIv1lrType) GetIv64() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1lrFieldIv64Mask) != 0
}

// SetIv64 IV64
func (r *registerIv1lrType) SetIv64(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1lrFieldIv64Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1lrFieldIv64Mask)
	}
}

// registerIv1rrType initialization vector registers
type registerIv1rrType uint32

const (
	RegisterIv1rrFieldIv127Shift = 0
	RegisterIv1rrFieldIv127Mask  = 0x1
)

// GetIv127 IV127
func (r *registerIv1rrType) GetIv127() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv127Mask) != 0
}

// SetIv127 IV127
func (r *registerIv1rrType) SetIv127(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv127Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv127Mask)
	}
}

const (
	RegisterIv1rrFieldIv126Shift = 1
	RegisterIv1rrFieldIv126Mask  = 0x2
)

// GetIv126 IV126
func (r *registerIv1rrType) GetIv126() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv126Mask) != 0
}

// SetIv126 IV126
func (r *registerIv1rrType) SetIv126(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv126Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv126Mask)
	}
}

const (
	RegisterIv1rrFieldIv125Shift = 2
	RegisterIv1rrFieldIv125Mask  = 0x4
)

// GetIv125 IV125
func (r *registerIv1rrType) GetIv125() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv125Mask) != 0
}

// SetIv125 IV125
func (r *registerIv1rrType) SetIv125(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv125Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv125Mask)
	}
}

const (
	RegisterIv1rrFieldIv124Shift = 3
	RegisterIv1rrFieldIv124Mask  = 0x8
)

// GetIv124 IV124
func (r *registerIv1rrType) GetIv124() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv124Mask) != 0
}

// SetIv124 IV124
func (r *registerIv1rrType) SetIv124(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv124Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv124Mask)
	}
}

const (
	RegisterIv1rrFieldIv123Shift = 4
	RegisterIv1rrFieldIv123Mask  = 0x10
)

// GetIv123 IV123
func (r *registerIv1rrType) GetIv123() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv123Mask) != 0
}

// SetIv123 IV123
func (r *registerIv1rrType) SetIv123(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv123Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv123Mask)
	}
}

const (
	RegisterIv1rrFieldIv122Shift = 5
	RegisterIv1rrFieldIv122Mask  = 0x20
)

// GetIv122 IV122
func (r *registerIv1rrType) GetIv122() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv122Mask) != 0
}

// SetIv122 IV122
func (r *registerIv1rrType) SetIv122(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv122Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv122Mask)
	}
}

const (
	RegisterIv1rrFieldIv121Shift = 6
	RegisterIv1rrFieldIv121Mask  = 0x40
)

// GetIv121 IV121
func (r *registerIv1rrType) GetIv121() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv121Mask) != 0
}

// SetIv121 IV121
func (r *registerIv1rrType) SetIv121(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv121Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv121Mask)
	}
}

const (
	RegisterIv1rrFieldIv120Shift = 7
	RegisterIv1rrFieldIv120Mask  = 0x80
)

// GetIv120 IV120
func (r *registerIv1rrType) GetIv120() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv120Mask) != 0
}

// SetIv120 IV120
func (r *registerIv1rrType) SetIv120(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv120Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv120Mask)
	}
}

const (
	RegisterIv1rrFieldIv119Shift = 8
	RegisterIv1rrFieldIv119Mask  = 0x100
)

// GetIv119 IV119
func (r *registerIv1rrType) GetIv119() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv119Mask) != 0
}

// SetIv119 IV119
func (r *registerIv1rrType) SetIv119(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv119Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv119Mask)
	}
}

const (
	RegisterIv1rrFieldIv118Shift = 9
	RegisterIv1rrFieldIv118Mask  = 0x200
)

// GetIv118 IV118
func (r *registerIv1rrType) GetIv118() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv118Mask) != 0
}

// SetIv118 IV118
func (r *registerIv1rrType) SetIv118(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv118Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv118Mask)
	}
}

const (
	RegisterIv1rrFieldIv117Shift = 10
	RegisterIv1rrFieldIv117Mask  = 0x400
)

// GetIv117 IV117
func (r *registerIv1rrType) GetIv117() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv117Mask) != 0
}

// SetIv117 IV117
func (r *registerIv1rrType) SetIv117(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv117Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv117Mask)
	}
}

const (
	RegisterIv1rrFieldIv116Shift = 11
	RegisterIv1rrFieldIv116Mask  = 0x800
)

// GetIv116 IV116
func (r *registerIv1rrType) GetIv116() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv116Mask) != 0
}

// SetIv116 IV116
func (r *registerIv1rrType) SetIv116(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv116Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv116Mask)
	}
}

const (
	RegisterIv1rrFieldIv115Shift = 12
	RegisterIv1rrFieldIv115Mask  = 0x1000
)

// GetIv115 IV115
func (r *registerIv1rrType) GetIv115() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv115Mask) != 0
}

// SetIv115 IV115
func (r *registerIv1rrType) SetIv115(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv115Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv115Mask)
	}
}

const (
	RegisterIv1rrFieldIv114Shift = 13
	RegisterIv1rrFieldIv114Mask  = 0x2000
)

// GetIv114 IV114
func (r *registerIv1rrType) GetIv114() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv114Mask) != 0
}

// SetIv114 IV114
func (r *registerIv1rrType) SetIv114(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv114Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv114Mask)
	}
}

const (
	RegisterIv1rrFieldIv113Shift = 14
	RegisterIv1rrFieldIv113Mask  = 0x4000
)

// GetIv113 IV113
func (r *registerIv1rrType) GetIv113() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv113Mask) != 0
}

// SetIv113 IV113
func (r *registerIv1rrType) SetIv113(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv113Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv113Mask)
	}
}

const (
	RegisterIv1rrFieldIv112Shift = 15
	RegisterIv1rrFieldIv112Mask  = 0x8000
)

// GetIv112 IV112
func (r *registerIv1rrType) GetIv112() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv112Mask) != 0
}

// SetIv112 IV112
func (r *registerIv1rrType) SetIv112(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv112Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv112Mask)
	}
}

const (
	RegisterIv1rrFieldIv111Shift = 16
	RegisterIv1rrFieldIv111Mask  = 0x10000
)

// GetIv111 IV111
func (r *registerIv1rrType) GetIv111() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv111Mask) != 0
}

// SetIv111 IV111
func (r *registerIv1rrType) SetIv111(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv111Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv111Mask)
	}
}

const (
	RegisterIv1rrFieldIv110Shift = 17
	RegisterIv1rrFieldIv110Mask  = 0x20000
)

// GetIv110 IV110
func (r *registerIv1rrType) GetIv110() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv110Mask) != 0
}

// SetIv110 IV110
func (r *registerIv1rrType) SetIv110(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv110Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv110Mask)
	}
}

const (
	RegisterIv1rrFieldIv109Shift = 18
	RegisterIv1rrFieldIv109Mask  = 0x40000
)

// GetIv109 IV109
func (r *registerIv1rrType) GetIv109() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv109Mask) != 0
}

// SetIv109 IV109
func (r *registerIv1rrType) SetIv109(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv109Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv109Mask)
	}
}

const (
	RegisterIv1rrFieldIv108Shift = 19
	RegisterIv1rrFieldIv108Mask  = 0x80000
)

// GetIv108 IV108
func (r *registerIv1rrType) GetIv108() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv108Mask) != 0
}

// SetIv108 IV108
func (r *registerIv1rrType) SetIv108(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv108Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv108Mask)
	}
}

const (
	RegisterIv1rrFieldIv107Shift = 20
	RegisterIv1rrFieldIv107Mask  = 0x100000
)

// GetIv107 IV107
func (r *registerIv1rrType) GetIv107() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv107Mask) != 0
}

// SetIv107 IV107
func (r *registerIv1rrType) SetIv107(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv107Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv107Mask)
	}
}

const (
	RegisterIv1rrFieldIv106Shift = 21
	RegisterIv1rrFieldIv106Mask  = 0x200000
)

// GetIv106 IV106
func (r *registerIv1rrType) GetIv106() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv106Mask) != 0
}

// SetIv106 IV106
func (r *registerIv1rrType) SetIv106(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv106Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv106Mask)
	}
}

const (
	RegisterIv1rrFieldIv105Shift = 22
	RegisterIv1rrFieldIv105Mask  = 0x400000
)

// GetIv105 IV105
func (r *registerIv1rrType) GetIv105() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv105Mask) != 0
}

// SetIv105 IV105
func (r *registerIv1rrType) SetIv105(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv105Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv105Mask)
	}
}

const (
	RegisterIv1rrFieldIv104Shift = 23
	RegisterIv1rrFieldIv104Mask  = 0x800000
)

// GetIv104 IV104
func (r *registerIv1rrType) GetIv104() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv104Mask) != 0
}

// SetIv104 IV104
func (r *registerIv1rrType) SetIv104(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv104Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv104Mask)
	}
}

const (
	RegisterIv1rrFieldIv103Shift = 24
	RegisterIv1rrFieldIv103Mask  = 0x1000000
)

// GetIv103 IV103
func (r *registerIv1rrType) GetIv103() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv103Mask) != 0
}

// SetIv103 IV103
func (r *registerIv1rrType) SetIv103(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv103Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv103Mask)
	}
}

const (
	RegisterIv1rrFieldIv102Shift = 25
	RegisterIv1rrFieldIv102Mask  = 0x2000000
)

// GetIv102 IV102
func (r *registerIv1rrType) GetIv102() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv102Mask) != 0
}

// SetIv102 IV102
func (r *registerIv1rrType) SetIv102(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv102Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv102Mask)
	}
}

const (
	RegisterIv1rrFieldIv101Shift = 26
	RegisterIv1rrFieldIv101Mask  = 0x4000000
)

// GetIv101 IV101
func (r *registerIv1rrType) GetIv101() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv101Mask) != 0
}

// SetIv101 IV101
func (r *registerIv1rrType) SetIv101(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv101Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv101Mask)
	}
}

const (
	RegisterIv1rrFieldIv100Shift = 27
	RegisterIv1rrFieldIv100Mask  = 0x8000000
)

// GetIv100 IV100
func (r *registerIv1rrType) GetIv100() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv100Mask) != 0
}

// SetIv100 IV100
func (r *registerIv1rrType) SetIv100(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv100Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv100Mask)
	}
}

const (
	RegisterIv1rrFieldIv99Shift = 28
	RegisterIv1rrFieldIv99Mask  = 0x10000000
)

// GetIv99 IV99
func (r *registerIv1rrType) GetIv99() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv99Mask) != 0
}

// SetIv99 IV99
func (r *registerIv1rrType) SetIv99(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv99Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv99Mask)
	}
}

const (
	RegisterIv1rrFieldIv98Shift = 29
	RegisterIv1rrFieldIv98Mask  = 0x20000000
)

// GetIv98 IV98
func (r *registerIv1rrType) GetIv98() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv98Mask) != 0
}

// SetIv98 IV98
func (r *registerIv1rrType) SetIv98(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv98Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv98Mask)
	}
}

const (
	RegisterIv1rrFieldIv97Shift = 30
	RegisterIv1rrFieldIv97Mask  = 0x40000000
)

// GetIv97 IV97
func (r *registerIv1rrType) GetIv97() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv97Mask) != 0
}

// SetIv97 IV97
func (r *registerIv1rrType) SetIv97(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv97Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv97Mask)
	}
}

const (
	RegisterIv1rrFieldIv96Shift = 31
	RegisterIv1rrFieldIv96Mask  = 0x80000000
)

// GetIv96 IV96
func (r *registerIv1rrType) GetIv96() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIv1rrFieldIv96Mask) != 0
}

// SetIv96 IV96
func (r *registerIv1rrType) SetIv96(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIv1rrFieldIv96Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIv1rrFieldIv96Mask)
	}
}

// registerCsgcmccm0rType context swap register
type registerCsgcmccm0rType uint32

const (
	RegisterCsgcmccm0rFieldCsgcmccm0Shift = 0
	RegisterCsgcmccm0rFieldCsgcmccm0Mask  = 0xffffffff
)

// GetCsgcmccm0 CSGCMCCM0
func (r *registerCsgcmccm0rType) GetCsgcmccm0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm0rFieldCsgcmccm0Mask) >> RegisterCsgcmccm0rFieldCsgcmccm0Shift)
}

// SetCsgcmccm0 CSGCMCCM0
func (r *registerCsgcmccm0rType) SetCsgcmccm0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm0rFieldCsgcmccm0Mask)|(uint32(value)<<RegisterCsgcmccm0rFieldCsgcmccm0Shift))
}

// registerCsgcmccm1rType context swap register
type registerCsgcmccm1rType uint32

const (
	RegisterCsgcmccm1rFieldCsgcmccm1Shift = 0
	RegisterCsgcmccm1rFieldCsgcmccm1Mask  = 0xffffffff
)

// GetCsgcmccm1 CSGCMCCM1
func (r *registerCsgcmccm1rType) GetCsgcmccm1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm1rFieldCsgcmccm1Mask) >> RegisterCsgcmccm1rFieldCsgcmccm1Shift)
}

// SetCsgcmccm1 CSGCMCCM1
func (r *registerCsgcmccm1rType) SetCsgcmccm1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm1rFieldCsgcmccm1Mask)|(uint32(value)<<RegisterCsgcmccm1rFieldCsgcmccm1Shift))
}

// registerCsgcmccm2rType context swap register
type registerCsgcmccm2rType uint32

const (
	RegisterCsgcmccm2rFieldCsgcmccm2Shift = 0
	RegisterCsgcmccm2rFieldCsgcmccm2Mask  = 0xffffffff
)

// GetCsgcmccm2 CSGCMCCM2
func (r *registerCsgcmccm2rType) GetCsgcmccm2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm2rFieldCsgcmccm2Mask) >> RegisterCsgcmccm2rFieldCsgcmccm2Shift)
}

// SetCsgcmccm2 CSGCMCCM2
func (r *registerCsgcmccm2rType) SetCsgcmccm2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm2rFieldCsgcmccm2Mask)|(uint32(value)<<RegisterCsgcmccm2rFieldCsgcmccm2Shift))
}

// registerCsgcmccm3rType context swap register
type registerCsgcmccm3rType uint32

const (
	RegisterCsgcmccm3rFieldCsgcmccm3Shift = 0
	RegisterCsgcmccm3rFieldCsgcmccm3Mask  = 0xffffffff
)

// GetCsgcmccm3 CSGCMCCM3
func (r *registerCsgcmccm3rType) GetCsgcmccm3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm3rFieldCsgcmccm3Mask) >> RegisterCsgcmccm3rFieldCsgcmccm3Shift)
}

// SetCsgcmccm3 CSGCMCCM3
func (r *registerCsgcmccm3rType) SetCsgcmccm3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm3rFieldCsgcmccm3Mask)|(uint32(value)<<RegisterCsgcmccm3rFieldCsgcmccm3Shift))
}

// registerCsgcmccm4rType context swap register
type registerCsgcmccm4rType uint32

const (
	RegisterCsgcmccm4rFieldCsgcmccm4Shift = 0
	RegisterCsgcmccm4rFieldCsgcmccm4Mask  = 0xffffffff
)

// GetCsgcmccm4 CSGCMCCM4
func (r *registerCsgcmccm4rType) GetCsgcmccm4() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm4rFieldCsgcmccm4Mask) >> RegisterCsgcmccm4rFieldCsgcmccm4Shift)
}

// SetCsgcmccm4 CSGCMCCM4
func (r *registerCsgcmccm4rType) SetCsgcmccm4(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm4rFieldCsgcmccm4Mask)|(uint32(value)<<RegisterCsgcmccm4rFieldCsgcmccm4Shift))
}

// registerCsgcmccm5rType context swap register
type registerCsgcmccm5rType uint32

const (
	RegisterCsgcmccm5rFieldCsgcmccm5Shift = 0
	RegisterCsgcmccm5rFieldCsgcmccm5Mask  = 0xffffffff
)

// GetCsgcmccm5 CSGCMCCM5
func (r *registerCsgcmccm5rType) GetCsgcmccm5() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm5rFieldCsgcmccm5Mask) >> RegisterCsgcmccm5rFieldCsgcmccm5Shift)
}

// SetCsgcmccm5 CSGCMCCM5
func (r *registerCsgcmccm5rType) SetCsgcmccm5(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm5rFieldCsgcmccm5Mask)|(uint32(value)<<RegisterCsgcmccm5rFieldCsgcmccm5Shift))
}

// registerCsgcmccm6rType context swap register
type registerCsgcmccm6rType uint32

const (
	RegisterCsgcmccm6rFieldCsgcmccm6Shift = 0
	RegisterCsgcmccm6rFieldCsgcmccm6Mask  = 0xffffffff
)

// GetCsgcmccm6 CSGCMCCM6
func (r *registerCsgcmccm6rType) GetCsgcmccm6() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm6rFieldCsgcmccm6Mask) >> RegisterCsgcmccm6rFieldCsgcmccm6Shift)
}

// SetCsgcmccm6 CSGCMCCM6
func (r *registerCsgcmccm6rType) SetCsgcmccm6(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm6rFieldCsgcmccm6Mask)|(uint32(value)<<RegisterCsgcmccm6rFieldCsgcmccm6Shift))
}

// registerCsgcmccm7rType context swap register
type registerCsgcmccm7rType uint32

const (
	RegisterCsgcmccm7rFieldCsgcmccm7Shift = 0
	RegisterCsgcmccm7rFieldCsgcmccm7Mask  = 0xffffffff
)

// GetCsgcmccm7 CSGCMCCM7
func (r *registerCsgcmccm7rType) GetCsgcmccm7() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm7rFieldCsgcmccm7Mask) >> RegisterCsgcmccm7rFieldCsgcmccm7Shift)
}

// SetCsgcmccm7 CSGCMCCM7
func (r *registerCsgcmccm7rType) SetCsgcmccm7(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm7rFieldCsgcmccm7Mask)|(uint32(value)<<RegisterCsgcmccm7rFieldCsgcmccm7Shift))
}

// registerCsgcm0rType context swap register
type registerCsgcm0rType uint32

const (
	RegisterCsgcm0rFieldCsgcm0Shift = 0
	RegisterCsgcm0rFieldCsgcm0Mask  = 0xffffffff
)

// GetCsgcm0 CSGCM0
func (r *registerCsgcm0rType) GetCsgcm0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm0rFieldCsgcm0Mask) >> RegisterCsgcm0rFieldCsgcm0Shift)
}

// SetCsgcm0 CSGCM0
func (r *registerCsgcm0rType) SetCsgcm0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm0rFieldCsgcm0Mask)|(uint32(value)<<RegisterCsgcm0rFieldCsgcm0Shift))
}

// registerCsgcm1rType context swap register
type registerCsgcm1rType uint32

const (
	RegisterCsgcm1rFieldCsgcm1Shift = 0
	RegisterCsgcm1rFieldCsgcm1Mask  = 0xffffffff
)

// GetCsgcm1 CSGCM1
func (r *registerCsgcm1rType) GetCsgcm1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm1rFieldCsgcm1Mask) >> RegisterCsgcm1rFieldCsgcm1Shift)
}

// SetCsgcm1 CSGCM1
func (r *registerCsgcm1rType) SetCsgcm1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm1rFieldCsgcm1Mask)|(uint32(value)<<RegisterCsgcm1rFieldCsgcm1Shift))
}

// registerCsgcm2rType context swap register
type registerCsgcm2rType uint32

const (
	RegisterCsgcm2rFieldCsgcm2Shift = 0
	RegisterCsgcm2rFieldCsgcm2Mask  = 0xffffffff
)

// GetCsgcm2 CSGCM2
func (r *registerCsgcm2rType) GetCsgcm2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm2rFieldCsgcm2Mask) >> RegisterCsgcm2rFieldCsgcm2Shift)
}

// SetCsgcm2 CSGCM2
func (r *registerCsgcm2rType) SetCsgcm2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm2rFieldCsgcm2Mask)|(uint32(value)<<RegisterCsgcm2rFieldCsgcm2Shift))
}

// registerCsgcm3rType context swap register
type registerCsgcm3rType uint32

const (
	RegisterCsgcm3rFieldCsgcm3Shift = 0
	RegisterCsgcm3rFieldCsgcm3Mask  = 0xffffffff
)

// GetCsgcm3 CSGCM3
func (r *registerCsgcm3rType) GetCsgcm3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm3rFieldCsgcm3Mask) >> RegisterCsgcm3rFieldCsgcm3Shift)
}

// SetCsgcm3 CSGCM3
func (r *registerCsgcm3rType) SetCsgcm3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm3rFieldCsgcm3Mask)|(uint32(value)<<RegisterCsgcm3rFieldCsgcm3Shift))
}

// registerCsgcm4rType context swap register
type registerCsgcm4rType uint32

const (
	RegisterCsgcm4rFieldCsgcm4Shift = 0
	RegisterCsgcm4rFieldCsgcm4Mask  = 0xffffffff
)

// GetCsgcm4 CSGCM4
func (r *registerCsgcm4rType) GetCsgcm4() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm4rFieldCsgcm4Mask) >> RegisterCsgcm4rFieldCsgcm4Shift)
}

// SetCsgcm4 CSGCM4
func (r *registerCsgcm4rType) SetCsgcm4(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm4rFieldCsgcm4Mask)|(uint32(value)<<RegisterCsgcm4rFieldCsgcm4Shift))
}

// registerCsgcm5rType context swap register
type registerCsgcm5rType uint32

const (
	RegisterCsgcm5rFieldCsgcm5Shift = 0
	RegisterCsgcm5rFieldCsgcm5Mask  = 0xffffffff
)

// GetCsgcm5 CSGCM5
func (r *registerCsgcm5rType) GetCsgcm5() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm5rFieldCsgcm5Mask) >> RegisterCsgcm5rFieldCsgcm5Shift)
}

// SetCsgcm5 CSGCM5
func (r *registerCsgcm5rType) SetCsgcm5(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm5rFieldCsgcm5Mask)|(uint32(value)<<RegisterCsgcm5rFieldCsgcm5Shift))
}

// registerCsgcm6rType context swap register
type registerCsgcm6rType uint32

const (
	RegisterCsgcm6rFieldCsgcm6Shift = 0
	RegisterCsgcm6rFieldCsgcm6Mask  = 0xffffffff
)

// GetCsgcm6 CSGCM6
func (r *registerCsgcm6rType) GetCsgcm6() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm6rFieldCsgcm6Mask) >> RegisterCsgcm6rFieldCsgcm6Shift)
}

// SetCsgcm6 CSGCM6
func (r *registerCsgcm6rType) SetCsgcm6(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm6rFieldCsgcm6Mask)|(uint32(value)<<RegisterCsgcm6rFieldCsgcm6Shift))
}

// registerCsgcm7rType context swap register
type registerCsgcm7rType uint32

const (
	RegisterCsgcm7rFieldCsgcm7Shift = 0
	RegisterCsgcm7rFieldCsgcm7Mask  = 0xffffffff
)

// GetCsgcm7 CSGCM7
func (r *registerCsgcm7rType) GetCsgcm7() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm7rFieldCsgcm7Mask) >> RegisterCsgcm7rFieldCsgcm7Shift)
}

// SetCsgcm7 CSGCM7
func (r *registerCsgcm7rType) SetCsgcm7(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm7rFieldCsgcm7Mask)|(uint32(value)<<RegisterCsgcm7rFieldCsgcm7Shift))
}
