//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

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
	RegisterCrFieldGcmccmphShift = 16
	RegisterCrFieldGcmccmphMask  = 0x30000
)

// GetGcmccmph GCM_CCMPH
func (r *registerCrType) GetGcmccmph() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldGcmccmphMask) >> RegisterCrFieldGcmccmphShift)
}

// SetGcmccmph GCM_CCMPH
func (r *registerCrType) SetGcmccmph(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldGcmccmphMask)|(uint32(value)<<RegisterCrFieldGcmccmphShift))
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

// registerImscrType interrupt mask set/clear register
type registerImscrType uint32

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

// registerRisrType raw interrupt status register
type registerRisrType uint32

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

// registerMisrType masked interrupt status register
type registerMisrType uint32

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

// registerK0lrType key registers
type registerK0lrType uint32

const (
	RegisterK0lrFieldB224Shift = 0
	RegisterK0lrFieldB224Mask  = 0x1
)

// GetB224 b224
func (r *registerK0lrType) GetB224() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB224Mask) != 0
}

// SetB224 b224
func (r *registerK0lrType) SetB224(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB224Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB224Mask)
	}
}

const (
	RegisterK0lrFieldB225Shift = 1
	RegisterK0lrFieldB225Mask  = 0x2
)

// GetB225 b225
func (r *registerK0lrType) GetB225() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB225Mask) != 0
}

// SetB225 b225
func (r *registerK0lrType) SetB225(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB225Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB225Mask)
	}
}

const (
	RegisterK0lrFieldB226Shift = 2
	RegisterK0lrFieldB226Mask  = 0x4
)

// GetB226 b226
func (r *registerK0lrType) GetB226() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB226Mask) != 0
}

// SetB226 b226
func (r *registerK0lrType) SetB226(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB226Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB226Mask)
	}
}

const (
	RegisterK0lrFieldB227Shift = 3
	RegisterK0lrFieldB227Mask  = 0x8
)

// GetB227 b227
func (r *registerK0lrType) GetB227() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB227Mask) != 0
}

// SetB227 b227
func (r *registerK0lrType) SetB227(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB227Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB227Mask)
	}
}

const (
	RegisterK0lrFieldB228Shift = 4
	RegisterK0lrFieldB228Mask  = 0x10
)

// GetB228 b228
func (r *registerK0lrType) GetB228() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB228Mask) != 0
}

// SetB228 b228
func (r *registerK0lrType) SetB228(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB228Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB228Mask)
	}
}

const (
	RegisterK0lrFieldB229Shift = 5
	RegisterK0lrFieldB229Mask  = 0x20
)

// GetB229 b229
func (r *registerK0lrType) GetB229() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB229Mask) != 0
}

// SetB229 b229
func (r *registerK0lrType) SetB229(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB229Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB229Mask)
	}
}

const (
	RegisterK0lrFieldB230Shift = 6
	RegisterK0lrFieldB230Mask  = 0x40
)

// GetB230 b230
func (r *registerK0lrType) GetB230() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB230Mask) != 0
}

// SetB230 b230
func (r *registerK0lrType) SetB230(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB230Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB230Mask)
	}
}

const (
	RegisterK0lrFieldB231Shift = 7
	RegisterK0lrFieldB231Mask  = 0x80
)

// GetB231 b231
func (r *registerK0lrType) GetB231() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB231Mask) != 0
}

// SetB231 b231
func (r *registerK0lrType) SetB231(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB231Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB231Mask)
	}
}

const (
	RegisterK0lrFieldB232Shift = 8
	RegisterK0lrFieldB232Mask  = 0x100
)

// GetB232 b232
func (r *registerK0lrType) GetB232() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB232Mask) != 0
}

// SetB232 b232
func (r *registerK0lrType) SetB232(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB232Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB232Mask)
	}
}

const (
	RegisterK0lrFieldB233Shift = 9
	RegisterK0lrFieldB233Mask  = 0x200
)

// GetB233 b233
func (r *registerK0lrType) GetB233() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB233Mask) != 0
}

// SetB233 b233
func (r *registerK0lrType) SetB233(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB233Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB233Mask)
	}
}

const (
	RegisterK0lrFieldB234Shift = 10
	RegisterK0lrFieldB234Mask  = 0x400
)

// GetB234 b234
func (r *registerK0lrType) GetB234() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB234Mask) != 0
}

// SetB234 b234
func (r *registerK0lrType) SetB234(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB234Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB234Mask)
	}
}

const (
	RegisterK0lrFieldB235Shift = 11
	RegisterK0lrFieldB235Mask  = 0x800
)

// GetB235 b235
func (r *registerK0lrType) GetB235() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB235Mask) != 0
}

// SetB235 b235
func (r *registerK0lrType) SetB235(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB235Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB235Mask)
	}
}

const (
	RegisterK0lrFieldB236Shift = 12
	RegisterK0lrFieldB236Mask  = 0x1000
)

// GetB236 b236
func (r *registerK0lrType) GetB236() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB236Mask) != 0
}

// SetB236 b236
func (r *registerK0lrType) SetB236(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB236Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB236Mask)
	}
}

const (
	RegisterK0lrFieldB237Shift = 13
	RegisterK0lrFieldB237Mask  = 0x2000
)

// GetB237 b237
func (r *registerK0lrType) GetB237() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB237Mask) != 0
}

// SetB237 b237
func (r *registerK0lrType) SetB237(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB237Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB237Mask)
	}
}

const (
	RegisterK0lrFieldB238Shift = 14
	RegisterK0lrFieldB238Mask  = 0x4000
)

// GetB238 b238
func (r *registerK0lrType) GetB238() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB238Mask) != 0
}

// SetB238 b238
func (r *registerK0lrType) SetB238(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB238Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB238Mask)
	}
}

const (
	RegisterK0lrFieldB239Shift = 15
	RegisterK0lrFieldB239Mask  = 0x8000
)

// GetB239 b239
func (r *registerK0lrType) GetB239() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB239Mask) != 0
}

// SetB239 b239
func (r *registerK0lrType) SetB239(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB239Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB239Mask)
	}
}

const (
	RegisterK0lrFieldB240Shift = 16
	RegisterK0lrFieldB240Mask  = 0x10000
)

// GetB240 b240
func (r *registerK0lrType) GetB240() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB240Mask) != 0
}

// SetB240 b240
func (r *registerK0lrType) SetB240(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB240Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB240Mask)
	}
}

const (
	RegisterK0lrFieldB241Shift = 17
	RegisterK0lrFieldB241Mask  = 0x20000
)

// GetB241 b241
func (r *registerK0lrType) GetB241() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB241Mask) != 0
}

// SetB241 b241
func (r *registerK0lrType) SetB241(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB241Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB241Mask)
	}
}

const (
	RegisterK0lrFieldB242Shift = 18
	RegisterK0lrFieldB242Mask  = 0x40000
)

// GetB242 b242
func (r *registerK0lrType) GetB242() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB242Mask) != 0
}

// SetB242 b242
func (r *registerK0lrType) SetB242(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB242Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB242Mask)
	}
}

const (
	RegisterK0lrFieldB243Shift = 19
	RegisterK0lrFieldB243Mask  = 0x80000
)

// GetB243 b243
func (r *registerK0lrType) GetB243() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB243Mask) != 0
}

// SetB243 b243
func (r *registerK0lrType) SetB243(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB243Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB243Mask)
	}
}

const (
	RegisterK0lrFieldB244Shift = 20
	RegisterK0lrFieldB244Mask  = 0x100000
)

// GetB244 b244
func (r *registerK0lrType) GetB244() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB244Mask) != 0
}

// SetB244 b244
func (r *registerK0lrType) SetB244(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB244Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB244Mask)
	}
}

const (
	RegisterK0lrFieldB245Shift = 21
	RegisterK0lrFieldB245Mask  = 0x200000
)

// GetB245 b245
func (r *registerK0lrType) GetB245() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB245Mask) != 0
}

// SetB245 b245
func (r *registerK0lrType) SetB245(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB245Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB245Mask)
	}
}

const (
	RegisterK0lrFieldB246Shift = 22
	RegisterK0lrFieldB246Mask  = 0x400000
)

// GetB246 b246
func (r *registerK0lrType) GetB246() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB246Mask) != 0
}

// SetB246 b246
func (r *registerK0lrType) SetB246(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB246Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB246Mask)
	}
}

const (
	RegisterK0lrFieldB247Shift = 23
	RegisterK0lrFieldB247Mask  = 0x800000
)

// GetB247 b247
func (r *registerK0lrType) GetB247() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB247Mask) != 0
}

// SetB247 b247
func (r *registerK0lrType) SetB247(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB247Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB247Mask)
	}
}

const (
	RegisterK0lrFieldB248Shift = 24
	RegisterK0lrFieldB248Mask  = 0x1000000
)

// GetB248 b248
func (r *registerK0lrType) GetB248() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB248Mask) != 0
}

// SetB248 b248
func (r *registerK0lrType) SetB248(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB248Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB248Mask)
	}
}

const (
	RegisterK0lrFieldB249Shift = 25
	RegisterK0lrFieldB249Mask  = 0x2000000
)

// GetB249 b249
func (r *registerK0lrType) GetB249() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB249Mask) != 0
}

// SetB249 b249
func (r *registerK0lrType) SetB249(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB249Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB249Mask)
	}
}

const (
	RegisterK0lrFieldB250Shift = 26
	RegisterK0lrFieldB250Mask  = 0x4000000
)

// GetB250 b250
func (r *registerK0lrType) GetB250() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB250Mask) != 0
}

// SetB250 b250
func (r *registerK0lrType) SetB250(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB250Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB250Mask)
	}
}

const (
	RegisterK0lrFieldB251Shift = 27
	RegisterK0lrFieldB251Mask  = 0x8000000
)

// GetB251 b251
func (r *registerK0lrType) GetB251() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB251Mask) != 0
}

// SetB251 b251
func (r *registerK0lrType) SetB251(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB251Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB251Mask)
	}
}

const (
	RegisterK0lrFieldB252Shift = 28
	RegisterK0lrFieldB252Mask  = 0x10000000
)

// GetB252 b252
func (r *registerK0lrType) GetB252() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB252Mask) != 0
}

// SetB252 b252
func (r *registerK0lrType) SetB252(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB252Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB252Mask)
	}
}

const (
	RegisterK0lrFieldB253Shift = 29
	RegisterK0lrFieldB253Mask  = 0x20000000
)

// GetB253 b253
func (r *registerK0lrType) GetB253() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB253Mask) != 0
}

// SetB253 b253
func (r *registerK0lrType) SetB253(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB253Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB253Mask)
	}
}

const (
	RegisterK0lrFieldB254Shift = 30
	RegisterK0lrFieldB254Mask  = 0x40000000
)

// GetB254 b254
func (r *registerK0lrType) GetB254() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB254Mask) != 0
}

// SetB254 b254
func (r *registerK0lrType) SetB254(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB254Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB254Mask)
	}
}

const (
	RegisterK0lrFieldB255Shift = 31
	RegisterK0lrFieldB255Mask  = 0x80000000
)

// GetB255 b255
func (r *registerK0lrType) GetB255() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0lrFieldB255Mask) != 0
}

// SetB255 b255
func (r *registerK0lrType) SetB255(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0lrFieldB255Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0lrFieldB255Mask)
	}
}

// registerK0rrType key registers
type registerK0rrType uint32

const (
	RegisterK0rrFieldB192Shift = 0
	RegisterK0rrFieldB192Mask  = 0x1
)

// GetB192 b192
func (r *registerK0rrType) GetB192() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB192Mask) != 0
}

// SetB192 b192
func (r *registerK0rrType) SetB192(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB192Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB192Mask)
	}
}

const (
	RegisterK0rrFieldB193Shift = 1
	RegisterK0rrFieldB193Mask  = 0x2
)

// GetB193 b193
func (r *registerK0rrType) GetB193() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB193Mask) != 0
}

// SetB193 b193
func (r *registerK0rrType) SetB193(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB193Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB193Mask)
	}
}

const (
	RegisterK0rrFieldB194Shift = 2
	RegisterK0rrFieldB194Mask  = 0x4
)

// GetB194 b194
func (r *registerK0rrType) GetB194() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB194Mask) != 0
}

// SetB194 b194
func (r *registerK0rrType) SetB194(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB194Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB194Mask)
	}
}

const (
	RegisterK0rrFieldB195Shift = 3
	RegisterK0rrFieldB195Mask  = 0x8
)

// GetB195 b195
func (r *registerK0rrType) GetB195() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB195Mask) != 0
}

// SetB195 b195
func (r *registerK0rrType) SetB195(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB195Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB195Mask)
	}
}

const (
	RegisterK0rrFieldB196Shift = 4
	RegisterK0rrFieldB196Mask  = 0x10
)

// GetB196 b196
func (r *registerK0rrType) GetB196() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB196Mask) != 0
}

// SetB196 b196
func (r *registerK0rrType) SetB196(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB196Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB196Mask)
	}
}

const (
	RegisterK0rrFieldB197Shift = 5
	RegisterK0rrFieldB197Mask  = 0x20
)

// GetB197 b197
func (r *registerK0rrType) GetB197() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB197Mask) != 0
}

// SetB197 b197
func (r *registerK0rrType) SetB197(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB197Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB197Mask)
	}
}

const (
	RegisterK0rrFieldB198Shift = 6
	RegisterK0rrFieldB198Mask  = 0x40
)

// GetB198 b198
func (r *registerK0rrType) GetB198() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB198Mask) != 0
}

// SetB198 b198
func (r *registerK0rrType) SetB198(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB198Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB198Mask)
	}
}

const (
	RegisterK0rrFieldB199Shift = 7
	RegisterK0rrFieldB199Mask  = 0x80
)

// GetB199 b199
func (r *registerK0rrType) GetB199() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB199Mask) != 0
}

// SetB199 b199
func (r *registerK0rrType) SetB199(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB199Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB199Mask)
	}
}

const (
	RegisterK0rrFieldB200Shift = 8
	RegisterK0rrFieldB200Mask  = 0x100
)

// GetB200 b200
func (r *registerK0rrType) GetB200() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB200Mask) != 0
}

// SetB200 b200
func (r *registerK0rrType) SetB200(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB200Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB200Mask)
	}
}

const (
	RegisterK0rrFieldB201Shift = 9
	RegisterK0rrFieldB201Mask  = 0x200
)

// GetB201 b201
func (r *registerK0rrType) GetB201() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB201Mask) != 0
}

// SetB201 b201
func (r *registerK0rrType) SetB201(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB201Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB201Mask)
	}
}

const (
	RegisterK0rrFieldB202Shift = 10
	RegisterK0rrFieldB202Mask  = 0x400
)

// GetB202 b202
func (r *registerK0rrType) GetB202() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB202Mask) != 0
}

// SetB202 b202
func (r *registerK0rrType) SetB202(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB202Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB202Mask)
	}
}

const (
	RegisterK0rrFieldB203Shift = 11
	RegisterK0rrFieldB203Mask  = 0x800
)

// GetB203 b203
func (r *registerK0rrType) GetB203() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB203Mask) != 0
}

// SetB203 b203
func (r *registerK0rrType) SetB203(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB203Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB203Mask)
	}
}

const (
	RegisterK0rrFieldB204Shift = 12
	RegisterK0rrFieldB204Mask  = 0x1000
)

// GetB204 b204
func (r *registerK0rrType) GetB204() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB204Mask) != 0
}

// SetB204 b204
func (r *registerK0rrType) SetB204(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB204Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB204Mask)
	}
}

const (
	RegisterK0rrFieldB205Shift = 13
	RegisterK0rrFieldB205Mask  = 0x2000
)

// GetB205 b205
func (r *registerK0rrType) GetB205() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB205Mask) != 0
}

// SetB205 b205
func (r *registerK0rrType) SetB205(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB205Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB205Mask)
	}
}

const (
	RegisterK0rrFieldB206Shift = 14
	RegisterK0rrFieldB206Mask  = 0x4000
)

// GetB206 b206
func (r *registerK0rrType) GetB206() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB206Mask) != 0
}

// SetB206 b206
func (r *registerK0rrType) SetB206(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB206Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB206Mask)
	}
}

const (
	RegisterK0rrFieldB207Shift = 15
	RegisterK0rrFieldB207Mask  = 0x8000
)

// GetB207 b207
func (r *registerK0rrType) GetB207() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB207Mask) != 0
}

// SetB207 b207
func (r *registerK0rrType) SetB207(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB207Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB207Mask)
	}
}

const (
	RegisterK0rrFieldB208Shift = 16
	RegisterK0rrFieldB208Mask  = 0x10000
)

// GetB208 b208
func (r *registerK0rrType) GetB208() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB208Mask) != 0
}

// SetB208 b208
func (r *registerK0rrType) SetB208(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB208Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB208Mask)
	}
}

const (
	RegisterK0rrFieldB209Shift = 17
	RegisterK0rrFieldB209Mask  = 0x20000
)

// GetB209 b209
func (r *registerK0rrType) GetB209() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB209Mask) != 0
}

// SetB209 b209
func (r *registerK0rrType) SetB209(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB209Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB209Mask)
	}
}

const (
	RegisterK0rrFieldB210Shift = 18
	RegisterK0rrFieldB210Mask  = 0x40000
)

// GetB210 b210
func (r *registerK0rrType) GetB210() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB210Mask) != 0
}

// SetB210 b210
func (r *registerK0rrType) SetB210(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB210Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB210Mask)
	}
}

const (
	RegisterK0rrFieldB211Shift = 19
	RegisterK0rrFieldB211Mask  = 0x80000
)

// GetB211 b211
func (r *registerK0rrType) GetB211() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB211Mask) != 0
}

// SetB211 b211
func (r *registerK0rrType) SetB211(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB211Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB211Mask)
	}
}

const (
	RegisterK0rrFieldB212Shift = 20
	RegisterK0rrFieldB212Mask  = 0x100000
)

// GetB212 b212
func (r *registerK0rrType) GetB212() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB212Mask) != 0
}

// SetB212 b212
func (r *registerK0rrType) SetB212(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB212Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB212Mask)
	}
}

const (
	RegisterK0rrFieldB213Shift = 21
	RegisterK0rrFieldB213Mask  = 0x200000
)

// GetB213 b213
func (r *registerK0rrType) GetB213() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB213Mask) != 0
}

// SetB213 b213
func (r *registerK0rrType) SetB213(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB213Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB213Mask)
	}
}

const (
	RegisterK0rrFieldB214Shift = 22
	RegisterK0rrFieldB214Mask  = 0x400000
)

// GetB214 b214
func (r *registerK0rrType) GetB214() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB214Mask) != 0
}

// SetB214 b214
func (r *registerK0rrType) SetB214(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB214Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB214Mask)
	}
}

const (
	RegisterK0rrFieldB215Shift = 23
	RegisterK0rrFieldB215Mask  = 0x800000
)

// GetB215 b215
func (r *registerK0rrType) GetB215() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB215Mask) != 0
}

// SetB215 b215
func (r *registerK0rrType) SetB215(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB215Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB215Mask)
	}
}

const (
	RegisterK0rrFieldB216Shift = 24
	RegisterK0rrFieldB216Mask  = 0x1000000
)

// GetB216 b216
func (r *registerK0rrType) GetB216() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB216Mask) != 0
}

// SetB216 b216
func (r *registerK0rrType) SetB216(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB216Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB216Mask)
	}
}

const (
	RegisterK0rrFieldB217Shift = 25
	RegisterK0rrFieldB217Mask  = 0x2000000
)

// GetB217 b217
func (r *registerK0rrType) GetB217() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB217Mask) != 0
}

// SetB217 b217
func (r *registerK0rrType) SetB217(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB217Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB217Mask)
	}
}

const (
	RegisterK0rrFieldB218Shift = 26
	RegisterK0rrFieldB218Mask  = 0x4000000
)

// GetB218 b218
func (r *registerK0rrType) GetB218() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB218Mask) != 0
}

// SetB218 b218
func (r *registerK0rrType) SetB218(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB218Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB218Mask)
	}
}

const (
	RegisterK0rrFieldB219Shift = 27
	RegisterK0rrFieldB219Mask  = 0x8000000
)

// GetB219 b219
func (r *registerK0rrType) GetB219() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB219Mask) != 0
}

// SetB219 b219
func (r *registerK0rrType) SetB219(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB219Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB219Mask)
	}
}

const (
	RegisterK0rrFieldB220Shift = 28
	RegisterK0rrFieldB220Mask  = 0x10000000
)

// GetB220 b220
func (r *registerK0rrType) GetB220() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB220Mask) != 0
}

// SetB220 b220
func (r *registerK0rrType) SetB220(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB220Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB220Mask)
	}
}

const (
	RegisterK0rrFieldB221Shift = 29
	RegisterK0rrFieldB221Mask  = 0x20000000
)

// GetB221 b221
func (r *registerK0rrType) GetB221() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB221Mask) != 0
}

// SetB221 b221
func (r *registerK0rrType) SetB221(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB221Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB221Mask)
	}
}

const (
	RegisterK0rrFieldB222Shift = 30
	RegisterK0rrFieldB222Mask  = 0x40000000
)

// GetB222 b222
func (r *registerK0rrType) GetB222() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB222Mask) != 0
}

// SetB222 b222
func (r *registerK0rrType) SetB222(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB222Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB222Mask)
	}
}

const (
	RegisterK0rrFieldB223Shift = 31
	RegisterK0rrFieldB223Mask  = 0x80000000
)

// GetB223 b223
func (r *registerK0rrType) GetB223() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK0rrFieldB223Mask) != 0
}

// SetB223 b223
func (r *registerK0rrType) SetB223(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK0rrFieldB223Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK0rrFieldB223Mask)
	}
}

// registerK1lrType key registers
type registerK1lrType uint32

const (
	RegisterK1lrFieldB160Shift = 0
	RegisterK1lrFieldB160Mask  = 0x1
)

// GetB160 b160
func (r *registerK1lrType) GetB160() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB160Mask) != 0
}

// SetB160 b160
func (r *registerK1lrType) SetB160(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB160Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB160Mask)
	}
}

const (
	RegisterK1lrFieldB161Shift = 1
	RegisterK1lrFieldB161Mask  = 0x2
)

// GetB161 b161
func (r *registerK1lrType) GetB161() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB161Mask) != 0
}

// SetB161 b161
func (r *registerK1lrType) SetB161(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB161Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB161Mask)
	}
}

const (
	RegisterK1lrFieldB162Shift = 2
	RegisterK1lrFieldB162Mask  = 0x4
)

// GetB162 b162
func (r *registerK1lrType) GetB162() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB162Mask) != 0
}

// SetB162 b162
func (r *registerK1lrType) SetB162(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB162Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB162Mask)
	}
}

const (
	RegisterK1lrFieldB163Shift = 3
	RegisterK1lrFieldB163Mask  = 0x8
)

// GetB163 b163
func (r *registerK1lrType) GetB163() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB163Mask) != 0
}

// SetB163 b163
func (r *registerK1lrType) SetB163(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB163Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB163Mask)
	}
}

const (
	RegisterK1lrFieldB164Shift = 4
	RegisterK1lrFieldB164Mask  = 0x10
)

// GetB164 b164
func (r *registerK1lrType) GetB164() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB164Mask) != 0
}

// SetB164 b164
func (r *registerK1lrType) SetB164(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB164Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB164Mask)
	}
}

const (
	RegisterK1lrFieldB165Shift = 5
	RegisterK1lrFieldB165Mask  = 0x20
)

// GetB165 b165
func (r *registerK1lrType) GetB165() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB165Mask) != 0
}

// SetB165 b165
func (r *registerK1lrType) SetB165(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB165Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB165Mask)
	}
}

const (
	RegisterK1lrFieldB166Shift = 6
	RegisterK1lrFieldB166Mask  = 0x40
)

// GetB166 b166
func (r *registerK1lrType) GetB166() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB166Mask) != 0
}

// SetB166 b166
func (r *registerK1lrType) SetB166(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB166Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB166Mask)
	}
}

const (
	RegisterK1lrFieldB167Shift = 7
	RegisterK1lrFieldB167Mask  = 0x80
)

// GetB167 b167
func (r *registerK1lrType) GetB167() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB167Mask) != 0
}

// SetB167 b167
func (r *registerK1lrType) SetB167(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB167Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB167Mask)
	}
}

const (
	RegisterK1lrFieldB168Shift = 8
	RegisterK1lrFieldB168Mask  = 0x100
)

// GetB168 b168
func (r *registerK1lrType) GetB168() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB168Mask) != 0
}

// SetB168 b168
func (r *registerK1lrType) SetB168(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB168Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB168Mask)
	}
}

const (
	RegisterK1lrFieldB169Shift = 9
	RegisterK1lrFieldB169Mask  = 0x200
)

// GetB169 b169
func (r *registerK1lrType) GetB169() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB169Mask) != 0
}

// SetB169 b169
func (r *registerK1lrType) SetB169(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB169Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB169Mask)
	}
}

const (
	RegisterK1lrFieldB170Shift = 10
	RegisterK1lrFieldB170Mask  = 0x400
)

// GetB170 b170
func (r *registerK1lrType) GetB170() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB170Mask) != 0
}

// SetB170 b170
func (r *registerK1lrType) SetB170(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB170Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB170Mask)
	}
}

const (
	RegisterK1lrFieldB171Shift = 11
	RegisterK1lrFieldB171Mask  = 0x800
)

// GetB171 b171
func (r *registerK1lrType) GetB171() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB171Mask) != 0
}

// SetB171 b171
func (r *registerK1lrType) SetB171(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB171Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB171Mask)
	}
}

const (
	RegisterK1lrFieldB172Shift = 12
	RegisterK1lrFieldB172Mask  = 0x1000
)

// GetB172 b172
func (r *registerK1lrType) GetB172() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB172Mask) != 0
}

// SetB172 b172
func (r *registerK1lrType) SetB172(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB172Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB172Mask)
	}
}

const (
	RegisterK1lrFieldB173Shift = 13
	RegisterK1lrFieldB173Mask  = 0x2000
)

// GetB173 b173
func (r *registerK1lrType) GetB173() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB173Mask) != 0
}

// SetB173 b173
func (r *registerK1lrType) SetB173(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB173Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB173Mask)
	}
}

const (
	RegisterK1lrFieldB174Shift = 14
	RegisterK1lrFieldB174Mask  = 0x4000
)

// GetB174 b174
func (r *registerK1lrType) GetB174() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB174Mask) != 0
}

// SetB174 b174
func (r *registerK1lrType) SetB174(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB174Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB174Mask)
	}
}

const (
	RegisterK1lrFieldB175Shift = 15
	RegisterK1lrFieldB175Mask  = 0x8000
)

// GetB175 b175
func (r *registerK1lrType) GetB175() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB175Mask) != 0
}

// SetB175 b175
func (r *registerK1lrType) SetB175(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB175Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB175Mask)
	}
}

const (
	RegisterK1lrFieldB176Shift = 16
	RegisterK1lrFieldB176Mask  = 0x10000
)

// GetB176 b176
func (r *registerK1lrType) GetB176() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB176Mask) != 0
}

// SetB176 b176
func (r *registerK1lrType) SetB176(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB176Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB176Mask)
	}
}

const (
	RegisterK1lrFieldB177Shift = 17
	RegisterK1lrFieldB177Mask  = 0x20000
)

// GetB177 b177
func (r *registerK1lrType) GetB177() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB177Mask) != 0
}

// SetB177 b177
func (r *registerK1lrType) SetB177(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB177Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB177Mask)
	}
}

const (
	RegisterK1lrFieldB178Shift = 18
	RegisterK1lrFieldB178Mask  = 0x40000
)

// GetB178 b178
func (r *registerK1lrType) GetB178() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB178Mask) != 0
}

// SetB178 b178
func (r *registerK1lrType) SetB178(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB178Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB178Mask)
	}
}

const (
	RegisterK1lrFieldB179Shift = 19
	RegisterK1lrFieldB179Mask  = 0x80000
)

// GetB179 b179
func (r *registerK1lrType) GetB179() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB179Mask) != 0
}

// SetB179 b179
func (r *registerK1lrType) SetB179(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB179Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB179Mask)
	}
}

const (
	RegisterK1lrFieldB180Shift = 20
	RegisterK1lrFieldB180Mask  = 0x100000
)

// GetB180 b180
func (r *registerK1lrType) GetB180() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB180Mask) != 0
}

// SetB180 b180
func (r *registerK1lrType) SetB180(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB180Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB180Mask)
	}
}

const (
	RegisterK1lrFieldB181Shift = 21
	RegisterK1lrFieldB181Mask  = 0x200000
)

// GetB181 b181
func (r *registerK1lrType) GetB181() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB181Mask) != 0
}

// SetB181 b181
func (r *registerK1lrType) SetB181(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB181Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB181Mask)
	}
}

const (
	RegisterK1lrFieldB182Shift = 22
	RegisterK1lrFieldB182Mask  = 0x400000
)

// GetB182 b182
func (r *registerK1lrType) GetB182() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB182Mask) != 0
}

// SetB182 b182
func (r *registerK1lrType) SetB182(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB182Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB182Mask)
	}
}

const (
	RegisterK1lrFieldB183Shift = 23
	RegisterK1lrFieldB183Mask  = 0x800000
)

// GetB183 b183
func (r *registerK1lrType) GetB183() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB183Mask) != 0
}

// SetB183 b183
func (r *registerK1lrType) SetB183(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB183Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB183Mask)
	}
}

const (
	RegisterK1lrFieldB184Shift = 24
	RegisterK1lrFieldB184Mask  = 0x1000000
)

// GetB184 b184
func (r *registerK1lrType) GetB184() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB184Mask) != 0
}

// SetB184 b184
func (r *registerK1lrType) SetB184(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB184Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB184Mask)
	}
}

const (
	RegisterK1lrFieldB185Shift = 25
	RegisterK1lrFieldB185Mask  = 0x2000000
)

// GetB185 b185
func (r *registerK1lrType) GetB185() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB185Mask) != 0
}

// SetB185 b185
func (r *registerK1lrType) SetB185(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB185Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB185Mask)
	}
}

const (
	RegisterK1lrFieldB186Shift = 26
	RegisterK1lrFieldB186Mask  = 0x4000000
)

// GetB186 b186
func (r *registerK1lrType) GetB186() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB186Mask) != 0
}

// SetB186 b186
func (r *registerK1lrType) SetB186(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB186Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB186Mask)
	}
}

const (
	RegisterK1lrFieldB187Shift = 27
	RegisterK1lrFieldB187Mask  = 0x8000000
)

// GetB187 b187
func (r *registerK1lrType) GetB187() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB187Mask) != 0
}

// SetB187 b187
func (r *registerK1lrType) SetB187(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB187Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB187Mask)
	}
}

const (
	RegisterK1lrFieldB188Shift = 28
	RegisterK1lrFieldB188Mask  = 0x10000000
)

// GetB188 b188
func (r *registerK1lrType) GetB188() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB188Mask) != 0
}

// SetB188 b188
func (r *registerK1lrType) SetB188(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB188Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB188Mask)
	}
}

const (
	RegisterK1lrFieldB189Shift = 29
	RegisterK1lrFieldB189Mask  = 0x20000000
)

// GetB189 b189
func (r *registerK1lrType) GetB189() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB189Mask) != 0
}

// SetB189 b189
func (r *registerK1lrType) SetB189(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB189Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB189Mask)
	}
}

const (
	RegisterK1lrFieldB190Shift = 30
	RegisterK1lrFieldB190Mask  = 0x40000000
)

// GetB190 b190
func (r *registerK1lrType) GetB190() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB190Mask) != 0
}

// SetB190 b190
func (r *registerK1lrType) SetB190(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB190Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB190Mask)
	}
}

const (
	RegisterK1lrFieldB191Shift = 31
	RegisterK1lrFieldB191Mask  = 0x80000000
)

// GetB191 b191
func (r *registerK1lrType) GetB191() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1lrFieldB191Mask) != 0
}

// SetB191 b191
func (r *registerK1lrType) SetB191(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1lrFieldB191Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1lrFieldB191Mask)
	}
}

// registerK1rrType key registers
type registerK1rrType uint32

const (
	RegisterK1rrFieldB128Shift = 0
	RegisterK1rrFieldB128Mask  = 0x1
)

// GetB128 b128
func (r *registerK1rrType) GetB128() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB128Mask) != 0
}

// SetB128 b128
func (r *registerK1rrType) SetB128(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB128Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB128Mask)
	}
}

const (
	RegisterK1rrFieldB129Shift = 1
	RegisterK1rrFieldB129Mask  = 0x2
)

// GetB129 b129
func (r *registerK1rrType) GetB129() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB129Mask) != 0
}

// SetB129 b129
func (r *registerK1rrType) SetB129(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB129Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB129Mask)
	}
}

const (
	RegisterK1rrFieldB130Shift = 2
	RegisterK1rrFieldB130Mask  = 0x4
)

// GetB130 b130
func (r *registerK1rrType) GetB130() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB130Mask) != 0
}

// SetB130 b130
func (r *registerK1rrType) SetB130(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB130Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB130Mask)
	}
}

const (
	RegisterK1rrFieldB131Shift = 3
	RegisterK1rrFieldB131Mask  = 0x8
)

// GetB131 b131
func (r *registerK1rrType) GetB131() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB131Mask) != 0
}

// SetB131 b131
func (r *registerK1rrType) SetB131(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB131Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB131Mask)
	}
}

const (
	RegisterK1rrFieldB132Shift = 4
	RegisterK1rrFieldB132Mask  = 0x10
)

// GetB132 b132
func (r *registerK1rrType) GetB132() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB132Mask) != 0
}

// SetB132 b132
func (r *registerK1rrType) SetB132(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB132Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB132Mask)
	}
}

const (
	RegisterK1rrFieldB133Shift = 5
	RegisterK1rrFieldB133Mask  = 0x20
)

// GetB133 b133
func (r *registerK1rrType) GetB133() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB133Mask) != 0
}

// SetB133 b133
func (r *registerK1rrType) SetB133(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB133Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB133Mask)
	}
}

const (
	RegisterK1rrFieldB134Shift = 6
	RegisterK1rrFieldB134Mask  = 0x40
)

// GetB134 b134
func (r *registerK1rrType) GetB134() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB134Mask) != 0
}

// SetB134 b134
func (r *registerK1rrType) SetB134(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB134Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB134Mask)
	}
}

const (
	RegisterK1rrFieldB135Shift = 7
	RegisterK1rrFieldB135Mask  = 0x80
)

// GetB135 b135
func (r *registerK1rrType) GetB135() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB135Mask) != 0
}

// SetB135 b135
func (r *registerK1rrType) SetB135(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB135Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB135Mask)
	}
}

const (
	RegisterK1rrFieldB136Shift = 8
	RegisterK1rrFieldB136Mask  = 0x100
)

// GetB136 b136
func (r *registerK1rrType) GetB136() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB136Mask) != 0
}

// SetB136 b136
func (r *registerK1rrType) SetB136(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB136Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB136Mask)
	}
}

const (
	RegisterK1rrFieldB137Shift = 9
	RegisterK1rrFieldB137Mask  = 0x200
)

// GetB137 b137
func (r *registerK1rrType) GetB137() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB137Mask) != 0
}

// SetB137 b137
func (r *registerK1rrType) SetB137(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB137Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB137Mask)
	}
}

const (
	RegisterK1rrFieldB138Shift = 10
	RegisterK1rrFieldB138Mask  = 0x400
)

// GetB138 b138
func (r *registerK1rrType) GetB138() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB138Mask) != 0
}

// SetB138 b138
func (r *registerK1rrType) SetB138(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB138Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB138Mask)
	}
}

const (
	RegisterK1rrFieldB139Shift = 11
	RegisterK1rrFieldB139Mask  = 0x800
)

// GetB139 b139
func (r *registerK1rrType) GetB139() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB139Mask) != 0
}

// SetB139 b139
func (r *registerK1rrType) SetB139(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB139Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB139Mask)
	}
}

const (
	RegisterK1rrFieldB140Shift = 12
	RegisterK1rrFieldB140Mask  = 0x1000
)

// GetB140 b140
func (r *registerK1rrType) GetB140() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB140Mask) != 0
}

// SetB140 b140
func (r *registerK1rrType) SetB140(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB140Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB140Mask)
	}
}

const (
	RegisterK1rrFieldB141Shift = 13
	RegisterK1rrFieldB141Mask  = 0x2000
)

// GetB141 b141
func (r *registerK1rrType) GetB141() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB141Mask) != 0
}

// SetB141 b141
func (r *registerK1rrType) SetB141(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB141Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB141Mask)
	}
}

const (
	RegisterK1rrFieldB142Shift = 14
	RegisterK1rrFieldB142Mask  = 0x4000
)

// GetB142 b142
func (r *registerK1rrType) GetB142() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB142Mask) != 0
}

// SetB142 b142
func (r *registerK1rrType) SetB142(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB142Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB142Mask)
	}
}

const (
	RegisterK1rrFieldB143Shift = 15
	RegisterK1rrFieldB143Mask  = 0x8000
)

// GetB143 b143
func (r *registerK1rrType) GetB143() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB143Mask) != 0
}

// SetB143 b143
func (r *registerK1rrType) SetB143(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB143Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB143Mask)
	}
}

const (
	RegisterK1rrFieldB144Shift = 16
	RegisterK1rrFieldB144Mask  = 0x10000
)

// GetB144 b144
func (r *registerK1rrType) GetB144() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB144Mask) != 0
}

// SetB144 b144
func (r *registerK1rrType) SetB144(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB144Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB144Mask)
	}
}

const (
	RegisterK1rrFieldB145Shift = 17
	RegisterK1rrFieldB145Mask  = 0x20000
)

// GetB145 b145
func (r *registerK1rrType) GetB145() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB145Mask) != 0
}

// SetB145 b145
func (r *registerK1rrType) SetB145(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB145Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB145Mask)
	}
}

const (
	RegisterK1rrFieldB146Shift = 18
	RegisterK1rrFieldB146Mask  = 0x40000
)

// GetB146 b146
func (r *registerK1rrType) GetB146() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB146Mask) != 0
}

// SetB146 b146
func (r *registerK1rrType) SetB146(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB146Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB146Mask)
	}
}

const (
	RegisterK1rrFieldB147Shift = 19
	RegisterK1rrFieldB147Mask  = 0x80000
)

// GetB147 b147
func (r *registerK1rrType) GetB147() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB147Mask) != 0
}

// SetB147 b147
func (r *registerK1rrType) SetB147(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB147Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB147Mask)
	}
}

const (
	RegisterK1rrFieldB148Shift = 20
	RegisterK1rrFieldB148Mask  = 0x100000
)

// GetB148 b148
func (r *registerK1rrType) GetB148() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB148Mask) != 0
}

// SetB148 b148
func (r *registerK1rrType) SetB148(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB148Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB148Mask)
	}
}

const (
	RegisterK1rrFieldB149Shift = 21
	RegisterK1rrFieldB149Mask  = 0x200000
)

// GetB149 b149
func (r *registerK1rrType) GetB149() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB149Mask) != 0
}

// SetB149 b149
func (r *registerK1rrType) SetB149(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB149Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB149Mask)
	}
}

const (
	RegisterK1rrFieldB150Shift = 22
	RegisterK1rrFieldB150Mask  = 0x400000
)

// GetB150 b150
func (r *registerK1rrType) GetB150() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB150Mask) != 0
}

// SetB150 b150
func (r *registerK1rrType) SetB150(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB150Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB150Mask)
	}
}

const (
	RegisterK1rrFieldB151Shift = 23
	RegisterK1rrFieldB151Mask  = 0x800000
)

// GetB151 b151
func (r *registerK1rrType) GetB151() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB151Mask) != 0
}

// SetB151 b151
func (r *registerK1rrType) SetB151(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB151Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB151Mask)
	}
}

const (
	RegisterK1rrFieldB152Shift = 24
	RegisterK1rrFieldB152Mask  = 0x1000000
)

// GetB152 b152
func (r *registerK1rrType) GetB152() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB152Mask) != 0
}

// SetB152 b152
func (r *registerK1rrType) SetB152(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB152Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB152Mask)
	}
}

const (
	RegisterK1rrFieldB153Shift = 25
	RegisterK1rrFieldB153Mask  = 0x2000000
)

// GetB153 b153
func (r *registerK1rrType) GetB153() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB153Mask) != 0
}

// SetB153 b153
func (r *registerK1rrType) SetB153(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB153Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB153Mask)
	}
}

const (
	RegisterK1rrFieldB154Shift = 26
	RegisterK1rrFieldB154Mask  = 0x4000000
)

// GetB154 b154
func (r *registerK1rrType) GetB154() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB154Mask) != 0
}

// SetB154 b154
func (r *registerK1rrType) SetB154(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB154Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB154Mask)
	}
}

const (
	RegisterK1rrFieldB155Shift = 27
	RegisterK1rrFieldB155Mask  = 0x8000000
)

// GetB155 b155
func (r *registerK1rrType) GetB155() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB155Mask) != 0
}

// SetB155 b155
func (r *registerK1rrType) SetB155(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB155Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB155Mask)
	}
}

const (
	RegisterK1rrFieldB156Shift = 28
	RegisterK1rrFieldB156Mask  = 0x10000000
)

// GetB156 b156
func (r *registerK1rrType) GetB156() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB156Mask) != 0
}

// SetB156 b156
func (r *registerK1rrType) SetB156(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB156Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB156Mask)
	}
}

const (
	RegisterK1rrFieldB157Shift = 29
	RegisterK1rrFieldB157Mask  = 0x20000000
)

// GetB157 b157
func (r *registerK1rrType) GetB157() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB157Mask) != 0
}

// SetB157 b157
func (r *registerK1rrType) SetB157(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB157Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB157Mask)
	}
}

const (
	RegisterK1rrFieldB158Shift = 30
	RegisterK1rrFieldB158Mask  = 0x40000000
)

// GetB158 b158
func (r *registerK1rrType) GetB158() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB158Mask) != 0
}

// SetB158 b158
func (r *registerK1rrType) SetB158(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB158Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB158Mask)
	}
}

const (
	RegisterK1rrFieldB159Shift = 31
	RegisterK1rrFieldB159Mask  = 0x80000000
)

// GetB159 b159
func (r *registerK1rrType) GetB159() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK1rrFieldB159Mask) != 0
}

// SetB159 b159
func (r *registerK1rrType) SetB159(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK1rrFieldB159Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK1rrFieldB159Mask)
	}
}

// registerK2lrType key registers
type registerK2lrType uint32

const (
	RegisterK2lrFieldB96Shift = 0
	RegisterK2lrFieldB96Mask  = 0x1
)

// GetB96 b96
func (r *registerK2lrType) GetB96() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB96Mask) != 0
}

// SetB96 b96
func (r *registerK2lrType) SetB96(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB96Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB96Mask)
	}
}

const (
	RegisterK2lrFieldB97Shift = 1
	RegisterK2lrFieldB97Mask  = 0x2
)

// GetB97 b97
func (r *registerK2lrType) GetB97() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB97Mask) != 0
}

// SetB97 b97
func (r *registerK2lrType) SetB97(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB97Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB97Mask)
	}
}

const (
	RegisterK2lrFieldB98Shift = 2
	RegisterK2lrFieldB98Mask  = 0x4
)

// GetB98 b98
func (r *registerK2lrType) GetB98() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB98Mask) != 0
}

// SetB98 b98
func (r *registerK2lrType) SetB98(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB98Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB98Mask)
	}
}

const (
	RegisterK2lrFieldB99Shift = 3
	RegisterK2lrFieldB99Mask  = 0x8
)

// GetB99 b99
func (r *registerK2lrType) GetB99() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB99Mask) != 0
}

// SetB99 b99
func (r *registerK2lrType) SetB99(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB99Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB99Mask)
	}
}

const (
	RegisterK2lrFieldB100Shift = 4
	RegisterK2lrFieldB100Mask  = 0x10
)

// GetB100 b100
func (r *registerK2lrType) GetB100() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB100Mask) != 0
}

// SetB100 b100
func (r *registerK2lrType) SetB100(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB100Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB100Mask)
	}
}

const (
	RegisterK2lrFieldB101Shift = 5
	RegisterK2lrFieldB101Mask  = 0x20
)

// GetB101 b101
func (r *registerK2lrType) GetB101() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB101Mask) != 0
}

// SetB101 b101
func (r *registerK2lrType) SetB101(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB101Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB101Mask)
	}
}

const (
	RegisterK2lrFieldB102Shift = 6
	RegisterK2lrFieldB102Mask  = 0x40
)

// GetB102 b102
func (r *registerK2lrType) GetB102() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB102Mask) != 0
}

// SetB102 b102
func (r *registerK2lrType) SetB102(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB102Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB102Mask)
	}
}

const (
	RegisterK2lrFieldB103Shift = 7
	RegisterK2lrFieldB103Mask  = 0x80
)

// GetB103 b103
func (r *registerK2lrType) GetB103() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB103Mask) != 0
}

// SetB103 b103
func (r *registerK2lrType) SetB103(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB103Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB103Mask)
	}
}

const (
	RegisterK2lrFieldB104Shift = 8
	RegisterK2lrFieldB104Mask  = 0x100
)

// GetB104 b104
func (r *registerK2lrType) GetB104() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB104Mask) != 0
}

// SetB104 b104
func (r *registerK2lrType) SetB104(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB104Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB104Mask)
	}
}

const (
	RegisterK2lrFieldB105Shift = 9
	RegisterK2lrFieldB105Mask  = 0x200
)

// GetB105 b105
func (r *registerK2lrType) GetB105() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB105Mask) != 0
}

// SetB105 b105
func (r *registerK2lrType) SetB105(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB105Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB105Mask)
	}
}

const (
	RegisterK2lrFieldB106Shift = 10
	RegisterK2lrFieldB106Mask  = 0x400
)

// GetB106 b106
func (r *registerK2lrType) GetB106() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB106Mask) != 0
}

// SetB106 b106
func (r *registerK2lrType) SetB106(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB106Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB106Mask)
	}
}

const (
	RegisterK2lrFieldB107Shift = 11
	RegisterK2lrFieldB107Mask  = 0x800
)

// GetB107 b107
func (r *registerK2lrType) GetB107() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB107Mask) != 0
}

// SetB107 b107
func (r *registerK2lrType) SetB107(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB107Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB107Mask)
	}
}

const (
	RegisterK2lrFieldB108Shift = 12
	RegisterK2lrFieldB108Mask  = 0x1000
)

// GetB108 b108
func (r *registerK2lrType) GetB108() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB108Mask) != 0
}

// SetB108 b108
func (r *registerK2lrType) SetB108(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB108Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB108Mask)
	}
}

const (
	RegisterK2lrFieldB109Shift = 13
	RegisterK2lrFieldB109Mask  = 0x2000
)

// GetB109 b109
func (r *registerK2lrType) GetB109() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB109Mask) != 0
}

// SetB109 b109
func (r *registerK2lrType) SetB109(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB109Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB109Mask)
	}
}

const (
	RegisterK2lrFieldB110Shift = 14
	RegisterK2lrFieldB110Mask  = 0x4000
)

// GetB110 b110
func (r *registerK2lrType) GetB110() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB110Mask) != 0
}

// SetB110 b110
func (r *registerK2lrType) SetB110(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB110Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB110Mask)
	}
}

const (
	RegisterK2lrFieldB111Shift = 15
	RegisterK2lrFieldB111Mask  = 0x8000
)

// GetB111 b111
func (r *registerK2lrType) GetB111() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB111Mask) != 0
}

// SetB111 b111
func (r *registerK2lrType) SetB111(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB111Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB111Mask)
	}
}

const (
	RegisterK2lrFieldB112Shift = 16
	RegisterK2lrFieldB112Mask  = 0x10000
)

// GetB112 b112
func (r *registerK2lrType) GetB112() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB112Mask) != 0
}

// SetB112 b112
func (r *registerK2lrType) SetB112(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB112Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB112Mask)
	}
}

const (
	RegisterK2lrFieldB113Shift = 17
	RegisterK2lrFieldB113Mask  = 0x20000
)

// GetB113 b113
func (r *registerK2lrType) GetB113() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB113Mask) != 0
}

// SetB113 b113
func (r *registerK2lrType) SetB113(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB113Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB113Mask)
	}
}

const (
	RegisterK2lrFieldB114Shift = 18
	RegisterK2lrFieldB114Mask  = 0x40000
)

// GetB114 b114
func (r *registerK2lrType) GetB114() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB114Mask) != 0
}

// SetB114 b114
func (r *registerK2lrType) SetB114(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB114Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB114Mask)
	}
}

const (
	RegisterK2lrFieldB115Shift = 19
	RegisterK2lrFieldB115Mask  = 0x80000
)

// GetB115 b115
func (r *registerK2lrType) GetB115() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB115Mask) != 0
}

// SetB115 b115
func (r *registerK2lrType) SetB115(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB115Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB115Mask)
	}
}

const (
	RegisterK2lrFieldB116Shift = 20
	RegisterK2lrFieldB116Mask  = 0x100000
)

// GetB116 b116
func (r *registerK2lrType) GetB116() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB116Mask) != 0
}

// SetB116 b116
func (r *registerK2lrType) SetB116(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB116Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB116Mask)
	}
}

const (
	RegisterK2lrFieldB117Shift = 21
	RegisterK2lrFieldB117Mask  = 0x200000
)

// GetB117 b117
func (r *registerK2lrType) GetB117() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB117Mask) != 0
}

// SetB117 b117
func (r *registerK2lrType) SetB117(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB117Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB117Mask)
	}
}

const (
	RegisterK2lrFieldB118Shift = 22
	RegisterK2lrFieldB118Mask  = 0x400000
)

// GetB118 b118
func (r *registerK2lrType) GetB118() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB118Mask) != 0
}

// SetB118 b118
func (r *registerK2lrType) SetB118(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB118Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB118Mask)
	}
}

const (
	RegisterK2lrFieldB119Shift = 23
	RegisterK2lrFieldB119Mask  = 0x800000
)

// GetB119 b119
func (r *registerK2lrType) GetB119() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB119Mask) != 0
}

// SetB119 b119
func (r *registerK2lrType) SetB119(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB119Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB119Mask)
	}
}

const (
	RegisterK2lrFieldB120Shift = 24
	RegisterK2lrFieldB120Mask  = 0x1000000
)

// GetB120 b120
func (r *registerK2lrType) GetB120() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB120Mask) != 0
}

// SetB120 b120
func (r *registerK2lrType) SetB120(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB120Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB120Mask)
	}
}

const (
	RegisterK2lrFieldB121Shift = 25
	RegisterK2lrFieldB121Mask  = 0x2000000
)

// GetB121 b121
func (r *registerK2lrType) GetB121() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB121Mask) != 0
}

// SetB121 b121
func (r *registerK2lrType) SetB121(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB121Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB121Mask)
	}
}

const (
	RegisterK2lrFieldB122Shift = 26
	RegisterK2lrFieldB122Mask  = 0x4000000
)

// GetB122 b122
func (r *registerK2lrType) GetB122() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB122Mask) != 0
}

// SetB122 b122
func (r *registerK2lrType) SetB122(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB122Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB122Mask)
	}
}

const (
	RegisterK2lrFieldB123Shift = 27
	RegisterK2lrFieldB123Mask  = 0x8000000
)

// GetB123 b123
func (r *registerK2lrType) GetB123() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB123Mask) != 0
}

// SetB123 b123
func (r *registerK2lrType) SetB123(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB123Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB123Mask)
	}
}

const (
	RegisterK2lrFieldB124Shift = 28
	RegisterK2lrFieldB124Mask  = 0x10000000
)

// GetB124 b124
func (r *registerK2lrType) GetB124() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB124Mask) != 0
}

// SetB124 b124
func (r *registerK2lrType) SetB124(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB124Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB124Mask)
	}
}

const (
	RegisterK2lrFieldB125Shift = 29
	RegisterK2lrFieldB125Mask  = 0x20000000
)

// GetB125 b125
func (r *registerK2lrType) GetB125() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB125Mask) != 0
}

// SetB125 b125
func (r *registerK2lrType) SetB125(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB125Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB125Mask)
	}
}

const (
	RegisterK2lrFieldB126Shift = 30
	RegisterK2lrFieldB126Mask  = 0x40000000
)

// GetB126 b126
func (r *registerK2lrType) GetB126() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB126Mask) != 0
}

// SetB126 b126
func (r *registerK2lrType) SetB126(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB126Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB126Mask)
	}
}

const (
	RegisterK2lrFieldB127Shift = 31
	RegisterK2lrFieldB127Mask  = 0x80000000
)

// GetB127 b127
func (r *registerK2lrType) GetB127() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2lrFieldB127Mask) != 0
}

// SetB127 b127
func (r *registerK2lrType) SetB127(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2lrFieldB127Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2lrFieldB127Mask)
	}
}

// registerK2rrType key registers
type registerK2rrType uint32

const (
	RegisterK2rrFieldB64Shift = 0
	RegisterK2rrFieldB64Mask  = 0x1
)

// GetB64 b64
func (r *registerK2rrType) GetB64() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB64Mask) != 0
}

// SetB64 b64
func (r *registerK2rrType) SetB64(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB64Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB64Mask)
	}
}

const (
	RegisterK2rrFieldB65Shift = 1
	RegisterK2rrFieldB65Mask  = 0x2
)

// GetB65 b65
func (r *registerK2rrType) GetB65() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB65Mask) != 0
}

// SetB65 b65
func (r *registerK2rrType) SetB65(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB65Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB65Mask)
	}
}

const (
	RegisterK2rrFieldB66Shift = 2
	RegisterK2rrFieldB66Mask  = 0x4
)

// GetB66 b66
func (r *registerK2rrType) GetB66() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB66Mask) != 0
}

// SetB66 b66
func (r *registerK2rrType) SetB66(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB66Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB66Mask)
	}
}

const (
	RegisterK2rrFieldB67Shift = 3
	RegisterK2rrFieldB67Mask  = 0x8
)

// GetB67 b67
func (r *registerK2rrType) GetB67() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB67Mask) != 0
}

// SetB67 b67
func (r *registerK2rrType) SetB67(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB67Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB67Mask)
	}
}

const (
	RegisterK2rrFieldB68Shift = 4
	RegisterK2rrFieldB68Mask  = 0x10
)

// GetB68 b68
func (r *registerK2rrType) GetB68() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB68Mask) != 0
}

// SetB68 b68
func (r *registerK2rrType) SetB68(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB68Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB68Mask)
	}
}

const (
	RegisterK2rrFieldB69Shift = 5
	RegisterK2rrFieldB69Mask  = 0x20
)

// GetB69 b69
func (r *registerK2rrType) GetB69() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB69Mask) != 0
}

// SetB69 b69
func (r *registerK2rrType) SetB69(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB69Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB69Mask)
	}
}

const (
	RegisterK2rrFieldB70Shift = 6
	RegisterK2rrFieldB70Mask  = 0x40
)

// GetB70 b70
func (r *registerK2rrType) GetB70() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB70Mask) != 0
}

// SetB70 b70
func (r *registerK2rrType) SetB70(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB70Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB70Mask)
	}
}

const (
	RegisterK2rrFieldB71Shift = 7
	RegisterK2rrFieldB71Mask  = 0x80
)

// GetB71 b71
func (r *registerK2rrType) GetB71() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB71Mask) != 0
}

// SetB71 b71
func (r *registerK2rrType) SetB71(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB71Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB71Mask)
	}
}

const (
	RegisterK2rrFieldB72Shift = 8
	RegisterK2rrFieldB72Mask  = 0x100
)

// GetB72 b72
func (r *registerK2rrType) GetB72() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB72Mask) != 0
}

// SetB72 b72
func (r *registerK2rrType) SetB72(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB72Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB72Mask)
	}
}

const (
	RegisterK2rrFieldB73Shift = 9
	RegisterK2rrFieldB73Mask  = 0x200
)

// GetB73 b73
func (r *registerK2rrType) GetB73() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB73Mask) != 0
}

// SetB73 b73
func (r *registerK2rrType) SetB73(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB73Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB73Mask)
	}
}

const (
	RegisterK2rrFieldB74Shift = 10
	RegisterK2rrFieldB74Mask  = 0x400
)

// GetB74 b74
func (r *registerK2rrType) GetB74() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB74Mask) != 0
}

// SetB74 b74
func (r *registerK2rrType) SetB74(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB74Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB74Mask)
	}
}

const (
	RegisterK2rrFieldB75Shift = 11
	RegisterK2rrFieldB75Mask  = 0x800
)

// GetB75 b75
func (r *registerK2rrType) GetB75() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB75Mask) != 0
}

// SetB75 b75
func (r *registerK2rrType) SetB75(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB75Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB75Mask)
	}
}

const (
	RegisterK2rrFieldB76Shift = 12
	RegisterK2rrFieldB76Mask  = 0x1000
)

// GetB76 b76
func (r *registerK2rrType) GetB76() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB76Mask) != 0
}

// SetB76 b76
func (r *registerK2rrType) SetB76(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB76Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB76Mask)
	}
}

const (
	RegisterK2rrFieldB77Shift = 13
	RegisterK2rrFieldB77Mask  = 0x2000
)

// GetB77 b77
func (r *registerK2rrType) GetB77() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB77Mask) != 0
}

// SetB77 b77
func (r *registerK2rrType) SetB77(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB77Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB77Mask)
	}
}

const (
	RegisterK2rrFieldB78Shift = 14
	RegisterK2rrFieldB78Mask  = 0x4000
)

// GetB78 b78
func (r *registerK2rrType) GetB78() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB78Mask) != 0
}

// SetB78 b78
func (r *registerK2rrType) SetB78(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB78Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB78Mask)
	}
}

const (
	RegisterK2rrFieldB79Shift = 15
	RegisterK2rrFieldB79Mask  = 0x8000
)

// GetB79 b79
func (r *registerK2rrType) GetB79() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB79Mask) != 0
}

// SetB79 b79
func (r *registerK2rrType) SetB79(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB79Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB79Mask)
	}
}

const (
	RegisterK2rrFieldB80Shift = 16
	RegisterK2rrFieldB80Mask  = 0x10000
)

// GetB80 b80
func (r *registerK2rrType) GetB80() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB80Mask) != 0
}

// SetB80 b80
func (r *registerK2rrType) SetB80(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB80Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB80Mask)
	}
}

const (
	RegisterK2rrFieldB81Shift = 17
	RegisterK2rrFieldB81Mask  = 0x20000
)

// GetB81 b81
func (r *registerK2rrType) GetB81() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB81Mask) != 0
}

// SetB81 b81
func (r *registerK2rrType) SetB81(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB81Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB81Mask)
	}
}

const (
	RegisterK2rrFieldB82Shift = 18
	RegisterK2rrFieldB82Mask  = 0x40000
)

// GetB82 b82
func (r *registerK2rrType) GetB82() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB82Mask) != 0
}

// SetB82 b82
func (r *registerK2rrType) SetB82(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB82Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB82Mask)
	}
}

const (
	RegisterK2rrFieldB83Shift = 19
	RegisterK2rrFieldB83Mask  = 0x80000
)

// GetB83 b83
func (r *registerK2rrType) GetB83() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB83Mask) != 0
}

// SetB83 b83
func (r *registerK2rrType) SetB83(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB83Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB83Mask)
	}
}

const (
	RegisterK2rrFieldB84Shift = 20
	RegisterK2rrFieldB84Mask  = 0x100000
)

// GetB84 b84
func (r *registerK2rrType) GetB84() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB84Mask) != 0
}

// SetB84 b84
func (r *registerK2rrType) SetB84(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB84Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB84Mask)
	}
}

const (
	RegisterK2rrFieldB85Shift = 21
	RegisterK2rrFieldB85Mask  = 0x200000
)

// GetB85 b85
func (r *registerK2rrType) GetB85() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB85Mask) != 0
}

// SetB85 b85
func (r *registerK2rrType) SetB85(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB85Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB85Mask)
	}
}

const (
	RegisterK2rrFieldB86Shift = 22
	RegisterK2rrFieldB86Mask  = 0x400000
)

// GetB86 b86
func (r *registerK2rrType) GetB86() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB86Mask) != 0
}

// SetB86 b86
func (r *registerK2rrType) SetB86(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB86Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB86Mask)
	}
}

const (
	RegisterK2rrFieldB87Shift = 23
	RegisterK2rrFieldB87Mask  = 0x800000
)

// GetB87 b87
func (r *registerK2rrType) GetB87() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB87Mask) != 0
}

// SetB87 b87
func (r *registerK2rrType) SetB87(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB87Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB87Mask)
	}
}

const (
	RegisterK2rrFieldB88Shift = 24
	RegisterK2rrFieldB88Mask  = 0x1000000
)

// GetB88 b88
func (r *registerK2rrType) GetB88() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB88Mask) != 0
}

// SetB88 b88
func (r *registerK2rrType) SetB88(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB88Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB88Mask)
	}
}

const (
	RegisterK2rrFieldB89Shift = 25
	RegisterK2rrFieldB89Mask  = 0x2000000
)

// GetB89 b89
func (r *registerK2rrType) GetB89() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB89Mask) != 0
}

// SetB89 b89
func (r *registerK2rrType) SetB89(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB89Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB89Mask)
	}
}

const (
	RegisterK2rrFieldB90Shift = 26
	RegisterK2rrFieldB90Mask  = 0x4000000
)

// GetB90 b90
func (r *registerK2rrType) GetB90() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB90Mask) != 0
}

// SetB90 b90
func (r *registerK2rrType) SetB90(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB90Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB90Mask)
	}
}

const (
	RegisterK2rrFieldB91Shift = 27
	RegisterK2rrFieldB91Mask  = 0x8000000
)

// GetB91 b91
func (r *registerK2rrType) GetB91() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB91Mask) != 0
}

// SetB91 b91
func (r *registerK2rrType) SetB91(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB91Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB91Mask)
	}
}

const (
	RegisterK2rrFieldB92Shift = 28
	RegisterK2rrFieldB92Mask  = 0x10000000
)

// GetB92 b92
func (r *registerK2rrType) GetB92() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB92Mask) != 0
}

// SetB92 b92
func (r *registerK2rrType) SetB92(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB92Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB92Mask)
	}
}

const (
	RegisterK2rrFieldB93Shift = 29
	RegisterK2rrFieldB93Mask  = 0x20000000
)

// GetB93 b93
func (r *registerK2rrType) GetB93() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB93Mask) != 0
}

// SetB93 b93
func (r *registerK2rrType) SetB93(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB93Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB93Mask)
	}
}

const (
	RegisterK2rrFieldB94Shift = 30
	RegisterK2rrFieldB94Mask  = 0x40000000
)

// GetB94 b94
func (r *registerK2rrType) GetB94() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB94Mask) != 0
}

// SetB94 b94
func (r *registerK2rrType) SetB94(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB94Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB94Mask)
	}
}

const (
	RegisterK2rrFieldB95Shift = 31
	RegisterK2rrFieldB95Mask  = 0x80000000
)

// GetB95 b95
func (r *registerK2rrType) GetB95() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK2rrFieldB95Mask) != 0
}

// SetB95 b95
func (r *registerK2rrType) SetB95(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK2rrFieldB95Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK2rrFieldB95Mask)
	}
}

// registerK3lrType key registers
type registerK3lrType uint32

const (
	RegisterK3lrFieldB32Shift = 0
	RegisterK3lrFieldB32Mask  = 0x1
)

// GetB32 b32
func (r *registerK3lrType) GetB32() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB32Mask) != 0
}

// SetB32 b32
func (r *registerK3lrType) SetB32(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB32Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB32Mask)
	}
}

const (
	RegisterK3lrFieldB33Shift = 1
	RegisterK3lrFieldB33Mask  = 0x2
)

// GetB33 b33
func (r *registerK3lrType) GetB33() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB33Mask) != 0
}

// SetB33 b33
func (r *registerK3lrType) SetB33(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB33Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB33Mask)
	}
}

const (
	RegisterK3lrFieldB34Shift = 2
	RegisterK3lrFieldB34Mask  = 0x4
)

// GetB34 b34
func (r *registerK3lrType) GetB34() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB34Mask) != 0
}

// SetB34 b34
func (r *registerK3lrType) SetB34(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB34Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB34Mask)
	}
}

const (
	RegisterK3lrFieldB35Shift = 3
	RegisterK3lrFieldB35Mask  = 0x8
)

// GetB35 b35
func (r *registerK3lrType) GetB35() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB35Mask) != 0
}

// SetB35 b35
func (r *registerK3lrType) SetB35(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB35Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB35Mask)
	}
}

const (
	RegisterK3lrFieldB36Shift = 4
	RegisterK3lrFieldB36Mask  = 0x10
)

// GetB36 b36
func (r *registerK3lrType) GetB36() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB36Mask) != 0
}

// SetB36 b36
func (r *registerK3lrType) SetB36(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB36Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB36Mask)
	}
}

const (
	RegisterK3lrFieldB37Shift = 5
	RegisterK3lrFieldB37Mask  = 0x20
)

// GetB37 b37
func (r *registerK3lrType) GetB37() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB37Mask) != 0
}

// SetB37 b37
func (r *registerK3lrType) SetB37(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB37Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB37Mask)
	}
}

const (
	RegisterK3lrFieldB38Shift = 6
	RegisterK3lrFieldB38Mask  = 0x40
)

// GetB38 b38
func (r *registerK3lrType) GetB38() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB38Mask) != 0
}

// SetB38 b38
func (r *registerK3lrType) SetB38(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB38Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB38Mask)
	}
}

const (
	RegisterK3lrFieldB39Shift = 7
	RegisterK3lrFieldB39Mask  = 0x80
)

// GetB39 b39
func (r *registerK3lrType) GetB39() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB39Mask) != 0
}

// SetB39 b39
func (r *registerK3lrType) SetB39(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB39Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB39Mask)
	}
}

const (
	RegisterK3lrFieldB40Shift = 8
	RegisterK3lrFieldB40Mask  = 0x100
)

// GetB40 b40
func (r *registerK3lrType) GetB40() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB40Mask) != 0
}

// SetB40 b40
func (r *registerK3lrType) SetB40(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB40Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB40Mask)
	}
}

const (
	RegisterK3lrFieldB41Shift = 9
	RegisterK3lrFieldB41Mask  = 0x200
)

// GetB41 b41
func (r *registerK3lrType) GetB41() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB41Mask) != 0
}

// SetB41 b41
func (r *registerK3lrType) SetB41(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB41Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB41Mask)
	}
}

const (
	RegisterK3lrFieldB42Shift = 10
	RegisterK3lrFieldB42Mask  = 0x400
)

// GetB42 b42
func (r *registerK3lrType) GetB42() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB42Mask) != 0
}

// SetB42 b42
func (r *registerK3lrType) SetB42(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB42Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB42Mask)
	}
}

const (
	RegisterK3lrFieldB43Shift = 11
	RegisterK3lrFieldB43Mask  = 0x800
)

// GetB43 b43
func (r *registerK3lrType) GetB43() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB43Mask) != 0
}

// SetB43 b43
func (r *registerK3lrType) SetB43(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB43Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB43Mask)
	}
}

const (
	RegisterK3lrFieldB44Shift = 12
	RegisterK3lrFieldB44Mask  = 0x1000
)

// GetB44 b44
func (r *registerK3lrType) GetB44() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB44Mask) != 0
}

// SetB44 b44
func (r *registerK3lrType) SetB44(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB44Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB44Mask)
	}
}

const (
	RegisterK3lrFieldB45Shift = 13
	RegisterK3lrFieldB45Mask  = 0x2000
)

// GetB45 b45
func (r *registerK3lrType) GetB45() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB45Mask) != 0
}

// SetB45 b45
func (r *registerK3lrType) SetB45(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB45Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB45Mask)
	}
}

const (
	RegisterK3lrFieldB46Shift = 14
	RegisterK3lrFieldB46Mask  = 0x4000
)

// GetB46 b46
func (r *registerK3lrType) GetB46() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB46Mask) != 0
}

// SetB46 b46
func (r *registerK3lrType) SetB46(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB46Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB46Mask)
	}
}

const (
	RegisterK3lrFieldB47Shift = 15
	RegisterK3lrFieldB47Mask  = 0x8000
)

// GetB47 b47
func (r *registerK3lrType) GetB47() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB47Mask) != 0
}

// SetB47 b47
func (r *registerK3lrType) SetB47(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB47Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB47Mask)
	}
}

const (
	RegisterK3lrFieldB48Shift = 16
	RegisterK3lrFieldB48Mask  = 0x10000
)

// GetB48 b48
func (r *registerK3lrType) GetB48() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB48Mask) != 0
}

// SetB48 b48
func (r *registerK3lrType) SetB48(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB48Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB48Mask)
	}
}

const (
	RegisterK3lrFieldB49Shift = 17
	RegisterK3lrFieldB49Mask  = 0x20000
)

// GetB49 b49
func (r *registerK3lrType) GetB49() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB49Mask) != 0
}

// SetB49 b49
func (r *registerK3lrType) SetB49(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB49Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB49Mask)
	}
}

const (
	RegisterK3lrFieldB50Shift = 18
	RegisterK3lrFieldB50Mask  = 0x40000
)

// GetB50 b50
func (r *registerK3lrType) GetB50() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB50Mask) != 0
}

// SetB50 b50
func (r *registerK3lrType) SetB50(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB50Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB50Mask)
	}
}

const (
	RegisterK3lrFieldB51Shift = 19
	RegisterK3lrFieldB51Mask  = 0x80000
)

// GetB51 b51
func (r *registerK3lrType) GetB51() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB51Mask) != 0
}

// SetB51 b51
func (r *registerK3lrType) SetB51(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB51Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB51Mask)
	}
}

const (
	RegisterK3lrFieldB52Shift = 20
	RegisterK3lrFieldB52Mask  = 0x100000
)

// GetB52 b52
func (r *registerK3lrType) GetB52() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB52Mask) != 0
}

// SetB52 b52
func (r *registerK3lrType) SetB52(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB52Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB52Mask)
	}
}

const (
	RegisterK3lrFieldB53Shift = 21
	RegisterK3lrFieldB53Mask  = 0x200000
)

// GetB53 b53
func (r *registerK3lrType) GetB53() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB53Mask) != 0
}

// SetB53 b53
func (r *registerK3lrType) SetB53(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB53Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB53Mask)
	}
}

const (
	RegisterK3lrFieldB54Shift = 22
	RegisterK3lrFieldB54Mask  = 0x400000
)

// GetB54 b54
func (r *registerK3lrType) GetB54() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB54Mask) != 0
}

// SetB54 b54
func (r *registerK3lrType) SetB54(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB54Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB54Mask)
	}
}

const (
	RegisterK3lrFieldB55Shift = 23
	RegisterK3lrFieldB55Mask  = 0x800000
)

// GetB55 b55
func (r *registerK3lrType) GetB55() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB55Mask) != 0
}

// SetB55 b55
func (r *registerK3lrType) SetB55(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB55Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB55Mask)
	}
}

const (
	RegisterK3lrFieldB56Shift = 24
	RegisterK3lrFieldB56Mask  = 0x1000000
)

// GetB56 b56
func (r *registerK3lrType) GetB56() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB56Mask) != 0
}

// SetB56 b56
func (r *registerK3lrType) SetB56(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB56Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB56Mask)
	}
}

const (
	RegisterK3lrFieldB57Shift = 25
	RegisterK3lrFieldB57Mask  = 0x2000000
)

// GetB57 b57
func (r *registerK3lrType) GetB57() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB57Mask) != 0
}

// SetB57 b57
func (r *registerK3lrType) SetB57(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB57Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB57Mask)
	}
}

const (
	RegisterK3lrFieldB58Shift = 26
	RegisterK3lrFieldB58Mask  = 0x4000000
)

// GetB58 b58
func (r *registerK3lrType) GetB58() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB58Mask) != 0
}

// SetB58 b58
func (r *registerK3lrType) SetB58(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB58Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB58Mask)
	}
}

const (
	RegisterK3lrFieldB59Shift = 27
	RegisterK3lrFieldB59Mask  = 0x8000000
)

// GetB59 b59
func (r *registerK3lrType) GetB59() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB59Mask) != 0
}

// SetB59 b59
func (r *registerK3lrType) SetB59(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB59Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB59Mask)
	}
}

const (
	RegisterK3lrFieldB60Shift = 28
	RegisterK3lrFieldB60Mask  = 0x10000000
)

// GetB60 b60
func (r *registerK3lrType) GetB60() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB60Mask) != 0
}

// SetB60 b60
func (r *registerK3lrType) SetB60(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB60Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB60Mask)
	}
}

const (
	RegisterK3lrFieldB61Shift = 29
	RegisterK3lrFieldB61Mask  = 0x20000000
)

// GetB61 b61
func (r *registerK3lrType) GetB61() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB61Mask) != 0
}

// SetB61 b61
func (r *registerK3lrType) SetB61(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB61Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB61Mask)
	}
}

const (
	RegisterK3lrFieldB62Shift = 30
	RegisterK3lrFieldB62Mask  = 0x40000000
)

// GetB62 b62
func (r *registerK3lrType) GetB62() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB62Mask) != 0
}

// SetB62 b62
func (r *registerK3lrType) SetB62(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB62Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB62Mask)
	}
}

const (
	RegisterK3lrFieldB63Shift = 31
	RegisterK3lrFieldB63Mask  = 0x80000000
)

// GetB63 b63
func (r *registerK3lrType) GetB63() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3lrFieldB63Mask) != 0
}

// SetB63 b63
func (r *registerK3lrType) SetB63(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3lrFieldB63Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3lrFieldB63Mask)
	}
}

// registerK3rrType key registers
type registerK3rrType uint32

const (
	RegisterK3rrFieldB0Shift = 0
	RegisterK3rrFieldB0Mask  = 0x1
)

// GetB0 b0
func (r *registerK3rrType) GetB0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB0Mask) != 0
}

// SetB0 b0
func (r *registerK3rrType) SetB0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB0Mask)
	}
}

const (
	RegisterK3rrFieldB1Shift = 1
	RegisterK3rrFieldB1Mask  = 0x2
)

// GetB1 b1
func (r *registerK3rrType) GetB1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB1Mask) != 0
}

// SetB1 b1
func (r *registerK3rrType) SetB1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB1Mask)
	}
}

const (
	RegisterK3rrFieldB2Shift = 2
	RegisterK3rrFieldB2Mask  = 0x4
)

// GetB2 b2
func (r *registerK3rrType) GetB2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB2Mask) != 0
}

// SetB2 b2
func (r *registerK3rrType) SetB2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB2Mask)
	}
}

const (
	RegisterK3rrFieldB3Shift = 3
	RegisterK3rrFieldB3Mask  = 0x8
)

// GetB3 b3
func (r *registerK3rrType) GetB3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB3Mask) != 0
}

// SetB3 b3
func (r *registerK3rrType) SetB3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB3Mask)
	}
}

const (
	RegisterK3rrFieldB4Shift = 4
	RegisterK3rrFieldB4Mask  = 0x10
)

// GetB4 b4
func (r *registerK3rrType) GetB4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB4Mask) != 0
}

// SetB4 b4
func (r *registerK3rrType) SetB4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB4Mask)
	}
}

const (
	RegisterK3rrFieldB5Shift = 5
	RegisterK3rrFieldB5Mask  = 0x20
)

// GetB5 b5
func (r *registerK3rrType) GetB5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB5Mask) != 0
}

// SetB5 b5
func (r *registerK3rrType) SetB5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB5Mask)
	}
}

const (
	RegisterK3rrFieldB6Shift = 6
	RegisterK3rrFieldB6Mask  = 0x40
)

// GetB6 b6
func (r *registerK3rrType) GetB6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB6Mask) != 0
}

// SetB6 b6
func (r *registerK3rrType) SetB6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB6Mask)
	}
}

const (
	RegisterK3rrFieldB7Shift = 7
	RegisterK3rrFieldB7Mask  = 0x80
)

// GetB7 b7
func (r *registerK3rrType) GetB7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB7Mask) != 0
}

// SetB7 b7
func (r *registerK3rrType) SetB7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB7Mask)
	}
}

const (
	RegisterK3rrFieldB8Shift = 8
	RegisterK3rrFieldB8Mask  = 0x100
)

// GetB8 b8
func (r *registerK3rrType) GetB8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB8Mask) != 0
}

// SetB8 b8
func (r *registerK3rrType) SetB8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB8Mask)
	}
}

const (
	RegisterK3rrFieldB9Shift = 9
	RegisterK3rrFieldB9Mask  = 0x200
)

// GetB9 b9
func (r *registerK3rrType) GetB9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB9Mask) != 0
}

// SetB9 b9
func (r *registerK3rrType) SetB9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB9Mask)
	}
}

const (
	RegisterK3rrFieldB10Shift = 10
	RegisterK3rrFieldB10Mask  = 0x400
)

// GetB10 b10
func (r *registerK3rrType) GetB10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB10Mask) != 0
}

// SetB10 b10
func (r *registerK3rrType) SetB10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB10Mask)
	}
}

const (
	RegisterK3rrFieldB11Shift = 11
	RegisterK3rrFieldB11Mask  = 0x800
)

// GetB11 b11
func (r *registerK3rrType) GetB11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB11Mask) != 0
}

// SetB11 b11
func (r *registerK3rrType) SetB11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB11Mask)
	}
}

const (
	RegisterK3rrFieldB12Shift = 12
	RegisterK3rrFieldB12Mask  = 0x1000
)

// GetB12 b12
func (r *registerK3rrType) GetB12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB12Mask) != 0
}

// SetB12 b12
func (r *registerK3rrType) SetB12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB12Mask)
	}
}

const (
	RegisterK3rrFieldB13Shift = 13
	RegisterK3rrFieldB13Mask  = 0x2000
)

// GetB13 b13
func (r *registerK3rrType) GetB13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB13Mask) != 0
}

// SetB13 b13
func (r *registerK3rrType) SetB13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB13Mask)
	}
}

const (
	RegisterK3rrFieldB14Shift = 14
	RegisterK3rrFieldB14Mask  = 0x4000
)

// GetB14 b14
func (r *registerK3rrType) GetB14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB14Mask) != 0
}

// SetB14 b14
func (r *registerK3rrType) SetB14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB14Mask)
	}
}

const (
	RegisterK3rrFieldB15Shift = 15
	RegisterK3rrFieldB15Mask  = 0x8000
)

// GetB15 b15
func (r *registerK3rrType) GetB15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB15Mask) != 0
}

// SetB15 b15
func (r *registerK3rrType) SetB15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB15Mask)
	}
}

const (
	RegisterK3rrFieldB16Shift = 16
	RegisterK3rrFieldB16Mask  = 0x10000
)

// GetB16 b16
func (r *registerK3rrType) GetB16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB16Mask) != 0
}

// SetB16 b16
func (r *registerK3rrType) SetB16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB16Mask)
	}
}

const (
	RegisterK3rrFieldB17Shift = 17
	RegisterK3rrFieldB17Mask  = 0x20000
)

// GetB17 b17
func (r *registerK3rrType) GetB17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB17Mask) != 0
}

// SetB17 b17
func (r *registerK3rrType) SetB17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB17Mask)
	}
}

const (
	RegisterK3rrFieldB18Shift = 18
	RegisterK3rrFieldB18Mask  = 0x40000
)

// GetB18 b18
func (r *registerK3rrType) GetB18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB18Mask) != 0
}

// SetB18 b18
func (r *registerK3rrType) SetB18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB18Mask)
	}
}

const (
	RegisterK3rrFieldB19Shift = 19
	RegisterK3rrFieldB19Mask  = 0x80000
)

// GetB19 b19
func (r *registerK3rrType) GetB19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB19Mask) != 0
}

// SetB19 b19
func (r *registerK3rrType) SetB19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB19Mask)
	}
}

const (
	RegisterK3rrFieldB20Shift = 20
	RegisterK3rrFieldB20Mask  = 0x100000
)

// GetB20 b20
func (r *registerK3rrType) GetB20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB20Mask) != 0
}

// SetB20 b20
func (r *registerK3rrType) SetB20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB20Mask)
	}
}

const (
	RegisterK3rrFieldB21Shift = 21
	RegisterK3rrFieldB21Mask  = 0x200000
)

// GetB21 b21
func (r *registerK3rrType) GetB21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB21Mask) != 0
}

// SetB21 b21
func (r *registerK3rrType) SetB21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB21Mask)
	}
}

const (
	RegisterK3rrFieldB22Shift = 22
	RegisterK3rrFieldB22Mask  = 0x400000
)

// GetB22 b22
func (r *registerK3rrType) GetB22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB22Mask) != 0
}

// SetB22 b22
func (r *registerK3rrType) SetB22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB22Mask)
	}
}

const (
	RegisterK3rrFieldB23Shift = 23
	RegisterK3rrFieldB23Mask  = 0x800000
)

// GetB23 b23
func (r *registerK3rrType) GetB23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB23Mask) != 0
}

// SetB23 b23
func (r *registerK3rrType) SetB23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB23Mask)
	}
}

const (
	RegisterK3rrFieldB24Shift = 24
	RegisterK3rrFieldB24Mask  = 0x1000000
)

// GetB24 b24
func (r *registerK3rrType) GetB24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB24Mask) != 0
}

// SetB24 b24
func (r *registerK3rrType) SetB24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB24Mask)
	}
}

const (
	RegisterK3rrFieldB25Shift = 25
	RegisterK3rrFieldB25Mask  = 0x2000000
)

// GetB25 b25
func (r *registerK3rrType) GetB25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB25Mask) != 0
}

// SetB25 b25
func (r *registerK3rrType) SetB25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB25Mask)
	}
}

const (
	RegisterK3rrFieldB26Shift = 26
	RegisterK3rrFieldB26Mask  = 0x4000000
)

// GetB26 b26
func (r *registerK3rrType) GetB26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB26Mask) != 0
}

// SetB26 b26
func (r *registerK3rrType) SetB26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB26Mask)
	}
}

const (
	RegisterK3rrFieldB27Shift = 27
	RegisterK3rrFieldB27Mask  = 0x8000000
)

// GetB27 b27
func (r *registerK3rrType) GetB27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB27Mask) != 0
}

// SetB27 b27
func (r *registerK3rrType) SetB27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB27Mask)
	}
}

const (
	RegisterK3rrFieldB28Shift = 28
	RegisterK3rrFieldB28Mask  = 0x10000000
)

// GetB28 b28
func (r *registerK3rrType) GetB28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB28Mask) != 0
}

// SetB28 b28
func (r *registerK3rrType) SetB28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB28Mask)
	}
}

const (
	RegisterK3rrFieldB29Shift = 29
	RegisterK3rrFieldB29Mask  = 0x20000000
)

// GetB29 b29
func (r *registerK3rrType) GetB29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB29Mask) != 0
}

// SetB29 b29
func (r *registerK3rrType) SetB29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB29Mask)
	}
}

const (
	RegisterK3rrFieldB30Shift = 30
	RegisterK3rrFieldB30Mask  = 0x40000000
)

// GetB30 b30
func (r *registerK3rrType) GetB30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB30Mask) != 0
}

// SetB30 b30
func (r *registerK3rrType) SetB30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB30Mask)
	}
}

const (
	RegisterK3rrFieldB31Shift = 31
	RegisterK3rrFieldB31Mask  = 0x80000000
)

// GetB31 b31
func (r *registerK3rrType) GetB31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterK3rrFieldB31Mask) != 0
}

// SetB31 b31
func (r *registerK3rrType) SetB31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterK3rrFieldB31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterK3rrFieldB31Mask)
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
	RegisterCsgcmccm0rFieldCsgcmccm0rShift = 0
	RegisterCsgcmccm0rFieldCsgcmccm0rMask  = 0xffffffff
)

// GetCsgcmccm0r CSGCMCCM0R
func (r *registerCsgcmccm0rType) GetCsgcmccm0r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm0rFieldCsgcmccm0rMask) >> RegisterCsgcmccm0rFieldCsgcmccm0rShift)
}

// SetCsgcmccm0r CSGCMCCM0R
func (r *registerCsgcmccm0rType) SetCsgcmccm0r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm0rFieldCsgcmccm0rMask)|(uint32(value)<<RegisterCsgcmccm0rFieldCsgcmccm0rShift))
}

// registerCsgcmccm1rType context swap register
type registerCsgcmccm1rType uint32

const (
	RegisterCsgcmccm1rFieldCsgcmccm1rShift = 0
	RegisterCsgcmccm1rFieldCsgcmccm1rMask  = 0xffffffff
)

// GetCsgcmccm1r CSGCMCCM1R
func (r *registerCsgcmccm1rType) GetCsgcmccm1r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm1rFieldCsgcmccm1rMask) >> RegisterCsgcmccm1rFieldCsgcmccm1rShift)
}

// SetCsgcmccm1r CSGCMCCM1R
func (r *registerCsgcmccm1rType) SetCsgcmccm1r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm1rFieldCsgcmccm1rMask)|(uint32(value)<<RegisterCsgcmccm1rFieldCsgcmccm1rShift))
}

// registerCsgcmccm2rType context swap register
type registerCsgcmccm2rType uint32

const (
	RegisterCsgcmccm2rFieldCsgcmccm2rShift = 0
	RegisterCsgcmccm2rFieldCsgcmccm2rMask  = 0xffffffff
)

// GetCsgcmccm2r CSGCMCCM2R
func (r *registerCsgcmccm2rType) GetCsgcmccm2r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm2rFieldCsgcmccm2rMask) >> RegisterCsgcmccm2rFieldCsgcmccm2rShift)
}

// SetCsgcmccm2r CSGCMCCM2R
func (r *registerCsgcmccm2rType) SetCsgcmccm2r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm2rFieldCsgcmccm2rMask)|(uint32(value)<<RegisterCsgcmccm2rFieldCsgcmccm2rShift))
}

// registerCsgcmccm3rType context swap register
type registerCsgcmccm3rType uint32

const (
	RegisterCsgcmccm3rFieldCsgcmccm3rShift = 0
	RegisterCsgcmccm3rFieldCsgcmccm3rMask  = 0xffffffff
)

// GetCsgcmccm3r CSGCMCCM3R
func (r *registerCsgcmccm3rType) GetCsgcmccm3r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm3rFieldCsgcmccm3rMask) >> RegisterCsgcmccm3rFieldCsgcmccm3rShift)
}

// SetCsgcmccm3r CSGCMCCM3R
func (r *registerCsgcmccm3rType) SetCsgcmccm3r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm3rFieldCsgcmccm3rMask)|(uint32(value)<<RegisterCsgcmccm3rFieldCsgcmccm3rShift))
}

// registerCsgcmccm4rType context swap register
type registerCsgcmccm4rType uint32

const (
	RegisterCsgcmccm4rFieldCsgcmccm4rShift = 0
	RegisterCsgcmccm4rFieldCsgcmccm4rMask  = 0xffffffff
)

// GetCsgcmccm4r CSGCMCCM4R
func (r *registerCsgcmccm4rType) GetCsgcmccm4r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm4rFieldCsgcmccm4rMask) >> RegisterCsgcmccm4rFieldCsgcmccm4rShift)
}

// SetCsgcmccm4r CSGCMCCM4R
func (r *registerCsgcmccm4rType) SetCsgcmccm4r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm4rFieldCsgcmccm4rMask)|(uint32(value)<<RegisterCsgcmccm4rFieldCsgcmccm4rShift))
}

// registerCsgcmccm5rType context swap register
type registerCsgcmccm5rType uint32

const (
	RegisterCsgcmccm5rFieldCsgcmccm5rShift = 0
	RegisterCsgcmccm5rFieldCsgcmccm5rMask  = 0xffffffff
)

// GetCsgcmccm5r CSGCMCCM5R
func (r *registerCsgcmccm5rType) GetCsgcmccm5r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm5rFieldCsgcmccm5rMask) >> RegisterCsgcmccm5rFieldCsgcmccm5rShift)
}

// SetCsgcmccm5r CSGCMCCM5R
func (r *registerCsgcmccm5rType) SetCsgcmccm5r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm5rFieldCsgcmccm5rMask)|(uint32(value)<<RegisterCsgcmccm5rFieldCsgcmccm5rShift))
}

// registerCsgcmccm6rType context swap register
type registerCsgcmccm6rType uint32

const (
	RegisterCsgcmccm6rFieldCsgcmccm6rShift = 0
	RegisterCsgcmccm6rFieldCsgcmccm6rMask  = 0xffffffff
)

// GetCsgcmccm6r CSGCMCCM6R
func (r *registerCsgcmccm6rType) GetCsgcmccm6r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm6rFieldCsgcmccm6rMask) >> RegisterCsgcmccm6rFieldCsgcmccm6rShift)
}

// SetCsgcmccm6r CSGCMCCM6R
func (r *registerCsgcmccm6rType) SetCsgcmccm6r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm6rFieldCsgcmccm6rMask)|(uint32(value)<<RegisterCsgcmccm6rFieldCsgcmccm6rShift))
}

// registerCsgcmccm7rType context swap register
type registerCsgcmccm7rType uint32

const (
	RegisterCsgcmccm7rFieldCsgcmccm7rShift = 0
	RegisterCsgcmccm7rFieldCsgcmccm7rMask  = 0xffffffff
)

// GetCsgcmccm7r CSGCMCCM7R
func (r *registerCsgcmccm7rType) GetCsgcmccm7r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcmccm7rFieldCsgcmccm7rMask) >> RegisterCsgcmccm7rFieldCsgcmccm7rShift)
}

// SetCsgcmccm7r CSGCMCCM7R
func (r *registerCsgcmccm7rType) SetCsgcmccm7r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcmccm7rFieldCsgcmccm7rMask)|(uint32(value)<<RegisterCsgcmccm7rFieldCsgcmccm7rShift))
}

// registerCsgcm0rType context swap register
type registerCsgcm0rType uint32

const (
	RegisterCsgcm0rFieldCsgcm0rShift = 0
	RegisterCsgcm0rFieldCsgcm0rMask  = 0xffffffff
)

// GetCsgcm0r CSGCM0R
func (r *registerCsgcm0rType) GetCsgcm0r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm0rFieldCsgcm0rMask) >> RegisterCsgcm0rFieldCsgcm0rShift)
}

// SetCsgcm0r CSGCM0R
func (r *registerCsgcm0rType) SetCsgcm0r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm0rFieldCsgcm0rMask)|(uint32(value)<<RegisterCsgcm0rFieldCsgcm0rShift))
}

// registerCsgcm1rType context swap register
type registerCsgcm1rType uint32

const (
	RegisterCsgcm1rFieldCsgcm1rShift = 0
	RegisterCsgcm1rFieldCsgcm1rMask  = 0xffffffff
)

// GetCsgcm1r CSGCM1R
func (r *registerCsgcm1rType) GetCsgcm1r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm1rFieldCsgcm1rMask) >> RegisterCsgcm1rFieldCsgcm1rShift)
}

// SetCsgcm1r CSGCM1R
func (r *registerCsgcm1rType) SetCsgcm1r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm1rFieldCsgcm1rMask)|(uint32(value)<<RegisterCsgcm1rFieldCsgcm1rShift))
}

// registerCsgcm2rType context swap register
type registerCsgcm2rType uint32

const (
	RegisterCsgcm2rFieldCsgcm2rShift = 0
	RegisterCsgcm2rFieldCsgcm2rMask  = 0xffffffff
)

// GetCsgcm2r CSGCM2R
func (r *registerCsgcm2rType) GetCsgcm2r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm2rFieldCsgcm2rMask) >> RegisterCsgcm2rFieldCsgcm2rShift)
}

// SetCsgcm2r CSGCM2R
func (r *registerCsgcm2rType) SetCsgcm2r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm2rFieldCsgcm2rMask)|(uint32(value)<<RegisterCsgcm2rFieldCsgcm2rShift))
}

// registerCsgcm3rType context swap register
type registerCsgcm3rType uint32

const (
	RegisterCsgcm3rFieldCsgcm3rShift = 0
	RegisterCsgcm3rFieldCsgcm3rMask  = 0xffffffff
)

// GetCsgcm3r CSGCM3R
func (r *registerCsgcm3rType) GetCsgcm3r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm3rFieldCsgcm3rMask) >> RegisterCsgcm3rFieldCsgcm3rShift)
}

// SetCsgcm3r CSGCM3R
func (r *registerCsgcm3rType) SetCsgcm3r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm3rFieldCsgcm3rMask)|(uint32(value)<<RegisterCsgcm3rFieldCsgcm3rShift))
}

// registerCsgcm4rType context swap register
type registerCsgcm4rType uint32

const (
	RegisterCsgcm4rFieldCsgcm4rShift = 0
	RegisterCsgcm4rFieldCsgcm4rMask  = 0xffffffff
)

// GetCsgcm4r CSGCM4R
func (r *registerCsgcm4rType) GetCsgcm4r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm4rFieldCsgcm4rMask) >> RegisterCsgcm4rFieldCsgcm4rShift)
}

// SetCsgcm4r CSGCM4R
func (r *registerCsgcm4rType) SetCsgcm4r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm4rFieldCsgcm4rMask)|(uint32(value)<<RegisterCsgcm4rFieldCsgcm4rShift))
}

// registerCsgcm5rType context swap register
type registerCsgcm5rType uint32

const (
	RegisterCsgcm5rFieldCsgcm5rShift = 0
	RegisterCsgcm5rFieldCsgcm5rMask  = 0xffffffff
)

// GetCsgcm5r CSGCM5R
func (r *registerCsgcm5rType) GetCsgcm5r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm5rFieldCsgcm5rMask) >> RegisterCsgcm5rFieldCsgcm5rShift)
}

// SetCsgcm5r CSGCM5R
func (r *registerCsgcm5rType) SetCsgcm5r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm5rFieldCsgcm5rMask)|(uint32(value)<<RegisterCsgcm5rFieldCsgcm5rShift))
}

// registerCsgcm6rType context swap register
type registerCsgcm6rType uint32

const (
	RegisterCsgcm6rFieldCsgcm6rShift = 0
	RegisterCsgcm6rFieldCsgcm6rMask  = 0xffffffff
)

// GetCsgcm6r CSGCM6R
func (r *registerCsgcm6rType) GetCsgcm6r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm6rFieldCsgcm6rMask) >> RegisterCsgcm6rFieldCsgcm6rShift)
}

// SetCsgcm6r CSGCM6R
func (r *registerCsgcm6rType) SetCsgcm6r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm6rFieldCsgcm6rMask)|(uint32(value)<<RegisterCsgcm6rFieldCsgcm6rShift))
}

// registerCsgcm7rType context swap register
type registerCsgcm7rType uint32

const (
	RegisterCsgcm7rFieldCsgcm7rShift = 0
	RegisterCsgcm7rFieldCsgcm7rMask  = 0xffffffff
)

// GetCsgcm7r CSGCM7R
func (r *registerCsgcm7rType) GetCsgcm7r() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCsgcm7rFieldCsgcm7rMask) >> RegisterCsgcm7rFieldCsgcm7rShift)
}

// SetCsgcm7r CSGCM7R
func (r *registerCsgcm7rType) SetCsgcm7r(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsgcm7rFieldCsgcm7rMask)|(uint32(value)<<RegisterCsgcm7rFieldCsgcm7rShift))
}
