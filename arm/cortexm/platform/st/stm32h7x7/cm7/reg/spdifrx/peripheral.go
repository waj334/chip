//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package spdifrx

import (
	"unsafe"
	"volatile"
)

var (
	Spdifrx = (*_spdifrx)(unsafe.Pointer(uintptr(0x40004000)))
)

type _spdifrx struct {
	Cr   RegisterCrType
	Imr  RegisterImrType
	Sr   RegisterSrType
	Ifcr RegisterIfcrType
	Dr00 RegisterDr00Type
	Dr01 RegisterDr01Type
	Dr10 RegisterDr10Type
	Csr  RegisterCsrType
	Dir  RegisterDirType
	_    [984]byte
	Verr RegisterVerrType
	Idr  RegisterIdrType
	Sidr RegisterSidrType
}

// RegisterCrType Control register
type RegisterCrType uint32

func (r *RegisterCrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrFieldSpdifrxenShift = 0
	RegisterCrFieldSpdifrxenMask  = 0x3
)

// GetSpdifrxen Peripheral Block Enable
func (r *RegisterCrType) GetSpdifrxen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSpdifrxenMask) >> RegisterCrFieldSpdifrxenShift)
}

// SetSpdifrxen Peripheral Block Enable
func (r *RegisterCrType) SetSpdifrxen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSpdifrxenMask)|(uint32(value)<<RegisterCrFieldSpdifrxenShift))
}

const (
	RegisterCrFieldRxdmaenShift = 2
	RegisterCrFieldRxdmaenMask  = 0x4
)

// GetRxdmaen Receiver DMA ENable for data flow
func (r *RegisterCrType) GetRxdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRxdmaenMask) != 0
}

// SetRxdmaen Receiver DMA ENable for data flow
func (r *RegisterCrType) SetRxdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRxdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRxdmaenMask)
	}
}

const (
	RegisterCrFieldRxsteoShift = 3
	RegisterCrFieldRxsteoMask  = 0x8
)

// GetRxsteo STerEO Mode
func (r *RegisterCrType) GetRxsteo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRxsteoMask) != 0
}

// SetRxsteo STerEO Mode
func (r *RegisterCrType) SetRxsteo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRxsteoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRxsteoMask)
	}
}

const (
	RegisterCrFieldDrfmtShift = 4
	RegisterCrFieldDrfmtMask  = 0x30
)

// GetDrfmt RX Data format
func (r *RegisterCrType) GetDrfmt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDrfmtMask) >> RegisterCrFieldDrfmtShift)
}

// SetDrfmt RX Data format
func (r *RegisterCrType) SetDrfmt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDrfmtMask)|(uint32(value)<<RegisterCrFieldDrfmtShift))
}

const (
	RegisterCrFieldPmskShift = 6
	RegisterCrFieldPmskMask  = 0x40
)

// GetPmsk Mask Parity error bit
func (r *RegisterCrType) GetPmsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPmskMask) != 0
}

// SetPmsk Mask Parity error bit
func (r *RegisterCrType) SetPmsk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPmskMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPmskMask)
	}
}

const (
	RegisterCrFieldVmskShift = 7
	RegisterCrFieldVmskMask  = 0x80
)

// GetVmsk Mask of Validity bit
func (r *RegisterCrType) GetVmsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldVmskMask) != 0
}

// SetVmsk Mask of Validity bit
func (r *RegisterCrType) SetVmsk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldVmskMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldVmskMask)
	}
}

const (
	RegisterCrFieldCumskShift = 8
	RegisterCrFieldCumskMask  = 0x100
)

// GetCumsk Mask of channel status and user bits
func (r *RegisterCrType) GetCumsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCumskMask) != 0
}

// SetCumsk Mask of channel status and user bits
func (r *RegisterCrType) SetCumsk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCumskMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCumskMask)
	}
}

const (
	RegisterCrFieldPtmskShift = 9
	RegisterCrFieldPtmskMask  = 0x200
)

// GetPtmsk Mask of Preamble Type bits
func (r *RegisterCrType) GetPtmsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPtmskMask) != 0
}

// SetPtmsk Mask of Preamble Type bits
func (r *RegisterCrType) SetPtmsk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPtmskMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPtmskMask)
	}
}

const (
	RegisterCrFieldCbdmaenShift = 10
	RegisterCrFieldCbdmaenMask  = 0x400
)

// GetCbdmaen Control Buffer DMA ENable for control flow
func (r *RegisterCrType) GetCbdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCbdmaenMask) != 0
}

// SetCbdmaen Control Buffer DMA ENable for control flow
func (r *RegisterCrType) SetCbdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCbdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCbdmaenMask)
	}
}

const (
	RegisterCrFieldChselShift = 11
	RegisterCrFieldChselMask  = 0x800
)

// GetChsel Channel Selection
func (r *RegisterCrType) GetChsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldChselMask) != 0
}

// SetChsel Channel Selection
func (r *RegisterCrType) SetChsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldChselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldChselMask)
	}
}

const (
	RegisterCrFieldNbtrShift = 12
	RegisterCrFieldNbtrMask  = 0x3000
)

// GetNbtr Maximum allowed re-tries during synchronization phase
func (r *RegisterCrType) GetNbtr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldNbtrMask) >> RegisterCrFieldNbtrShift)
}

// SetNbtr Maximum allowed re-tries during synchronization phase
func (r *RegisterCrType) SetNbtr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldNbtrMask)|(uint32(value)<<RegisterCrFieldNbtrShift))
}

const (
	RegisterCrFieldWfaShift = 14
	RegisterCrFieldWfaMask  = 0x4000
)

// GetWfa Wait For Activity
func (r *RegisterCrType) GetWfa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldWfaMask) != 0
}

// SetWfa Wait For Activity
func (r *RegisterCrType) SetWfa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldWfaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldWfaMask)
	}
}

const (
	RegisterCrFieldInselShift = 16
	RegisterCrFieldInselMask  = 0x70000
)

// GetInsel input selection
func (r *RegisterCrType) GetInsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldInselMask) >> RegisterCrFieldInselShift)
}

// SetInsel input selection
func (r *RegisterCrType) SetInsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldInselMask)|(uint32(value)<<RegisterCrFieldInselShift))
}

const (
	RegisterCrFieldCksenShift = 20
	RegisterCrFieldCksenMask  = 0x100000
)

// GetCksen Symbol Clock Enable
func (r *RegisterCrType) GetCksen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCksenMask) != 0
}

// SetCksen Symbol Clock Enable
func (r *RegisterCrType) SetCksen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCksenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCksenMask)
	}
}

const (
	RegisterCrFieldCksbkpenShift = 21
	RegisterCrFieldCksbkpenMask  = 0x200000
)

// GetCksbkpen Backup Symbol Clock Enable
func (r *RegisterCrType) GetCksbkpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCksbkpenMask) != 0
}

// SetCksbkpen Backup Symbol Clock Enable
func (r *RegisterCrType) SetCksbkpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCksbkpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCksbkpenMask)
	}
}

// RegisterImrType Interrupt mask register
type RegisterImrType uint32

func (r *RegisterImrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterImrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterImrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterImrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterImrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterImrFieldRxneieShift = 0
	RegisterImrFieldRxneieMask  = 0x1
)

// GetRxneie RXNE interrupt enable
func (r *RegisterImrType) GetRxneie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldRxneieMask) != 0
}

// SetRxneie RXNE interrupt enable
func (r *RegisterImrType) SetRxneie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImrFieldRxneieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImrFieldRxneieMask)
	}
}

const (
	RegisterImrFieldCsrneieShift = 1
	RegisterImrFieldCsrneieMask  = 0x2
)

// GetCsrneie Control Buffer Ready Interrupt Enable
func (r *RegisterImrType) GetCsrneie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldCsrneieMask) != 0
}

// SetCsrneie Control Buffer Ready Interrupt Enable
func (r *RegisterImrType) SetCsrneie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImrFieldCsrneieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImrFieldCsrneieMask)
	}
}

const (
	RegisterImrFieldPerrieShift = 2
	RegisterImrFieldPerrieMask  = 0x4
)

// GetPerrie Parity error interrupt enable
func (r *RegisterImrType) GetPerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldPerrieMask) != 0
}

// SetPerrie Parity error interrupt enable
func (r *RegisterImrType) SetPerrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImrFieldPerrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImrFieldPerrieMask)
	}
}

const (
	RegisterImrFieldOvrieShift = 3
	RegisterImrFieldOvrieMask  = 0x8
)

// GetOvrie Overrun error Interrupt Enable
func (r *RegisterImrType) GetOvrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldOvrieMask) != 0
}

// SetOvrie Overrun error Interrupt Enable
func (r *RegisterImrType) SetOvrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImrFieldOvrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImrFieldOvrieMask)
	}
}

const (
	RegisterImrFieldSblkieShift = 4
	RegisterImrFieldSblkieMask  = 0x10
)

// GetSblkie Synchronization Block Detected Interrupt Enable
func (r *RegisterImrType) GetSblkie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldSblkieMask) != 0
}

// SetSblkie Synchronization Block Detected Interrupt Enable
func (r *RegisterImrType) SetSblkie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImrFieldSblkieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImrFieldSblkieMask)
	}
}

const (
	RegisterImrFieldSyncdieShift = 5
	RegisterImrFieldSyncdieMask  = 0x20
)

// GetSyncdie Synchronization Done
func (r *RegisterImrType) GetSyncdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldSyncdieMask) != 0
}

// SetSyncdie Synchronization Done
func (r *RegisterImrType) SetSyncdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImrFieldSyncdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImrFieldSyncdieMask)
	}
}

const (
	RegisterImrFieldIfeieShift = 6
	RegisterImrFieldIfeieMask  = 0x40
)

// GetIfeie Serial Interface Error Interrupt Enable
func (r *RegisterImrType) GetIfeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldIfeieMask) != 0
}

// SetIfeie Serial Interface Error Interrupt Enable
func (r *RegisterImrType) SetIfeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImrFieldIfeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImrFieldIfeieMask)
	}
}

// RegisterSrType Status register
type RegisterSrType uint32

func (r *RegisterSrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSrFieldRxneShift = 0
	RegisterSrFieldRxneMask  = 0x1
)

// GetRxne Read data register not empty
func (r *RegisterSrType) GetRxne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRxneMask) != 0
}

// SetRxne Read data register not empty
func (r *RegisterSrType) SetRxne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldRxneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldRxneMask)
	}
}

const (
	RegisterSrFieldCsrneShift = 1
	RegisterSrFieldCsrneMask  = 0x2
)

// GetCsrne Control Buffer register is not empty
func (r *RegisterSrType) GetCsrne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCsrneMask) != 0
}

// SetCsrne Control Buffer register is not empty
func (r *RegisterSrType) SetCsrne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCsrneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCsrneMask)
	}
}

const (
	RegisterSrFieldPerrShift = 2
	RegisterSrFieldPerrMask  = 0x4
)

// GetPerr Parity error
func (r *RegisterSrType) GetPerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldPerrMask) != 0
}

// SetPerr Parity error
func (r *RegisterSrType) SetPerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldPerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldPerrMask)
	}
}

const (
	RegisterSrFieldOvrShift = 3
	RegisterSrFieldOvrMask  = 0x8
)

// GetOvr Overrun error
func (r *RegisterSrType) GetOvr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldOvrMask) != 0
}

// SetOvr Overrun error
func (r *RegisterSrType) SetOvr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldOvrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldOvrMask)
	}
}

const (
	RegisterSrFieldSbdShift = 4
	RegisterSrFieldSbdMask  = 0x10
)

// GetSbd Synchronization Block Detected
func (r *RegisterSrType) GetSbd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSbdMask) != 0
}

// SetSbd Synchronization Block Detected
func (r *RegisterSrType) SetSbd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldSbdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldSbdMask)
	}
}

const (
	RegisterSrFieldSyncdShift = 5
	RegisterSrFieldSyncdMask  = 0x20
)

// GetSyncd Synchronization Done
func (r *RegisterSrType) GetSyncd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSyncdMask) != 0
}

// SetSyncd Synchronization Done
func (r *RegisterSrType) SetSyncd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldSyncdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldSyncdMask)
	}
}

const (
	RegisterSrFieldFerrShift = 6
	RegisterSrFieldFerrMask  = 0x40
)

// GetFerr Framing error
func (r *RegisterSrType) GetFerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldFerrMask) != 0
}

// SetFerr Framing error
func (r *RegisterSrType) SetFerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldFerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldFerrMask)
	}
}

const (
	RegisterSrFieldSerrShift = 7
	RegisterSrFieldSerrMask  = 0x80
)

// GetSerr Synchronization error
func (r *RegisterSrType) GetSerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSerrMask) != 0
}

// SetSerr Synchronization error
func (r *RegisterSrType) SetSerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldSerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldSerrMask)
	}
}

const (
	RegisterSrFieldTerrShift = 8
	RegisterSrFieldTerrMask  = 0x100
)

// GetTerr Time-out error
func (r *RegisterSrType) GetTerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTerrMask) != 0
}

// SetTerr Time-out error
func (r *RegisterSrType) SetTerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldTerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldTerrMask)
	}
}

const (
	RegisterSrFieldWidth5Shift = 16
	RegisterSrFieldWidth5Mask  = 0x7fff0000
)

// GetWidth5 Duration of 5 symbols counted with SPDIF_CLK
func (r *RegisterSrType) GetWidth5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldWidth5Mask) >> RegisterSrFieldWidth5Shift)
}

// SetWidth5 Duration of 5 symbols counted with SPDIF_CLK
func (r *RegisterSrType) SetWidth5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldWidth5Mask)|(uint32(value)<<RegisterSrFieldWidth5Shift))
}

// RegisterIfcrType Interrupt Flag Clear register
type RegisterIfcrType uint32

func (r *RegisterIfcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIfcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIfcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIfcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIfcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIfcrFieldPerrcfShift = 2
	RegisterIfcrFieldPerrcfMask  = 0x4
)

// GetPerrcf Clears the Parity error flag
func (r *RegisterIfcrType) GetPerrcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldPerrcfMask) != 0
}

// SetPerrcf Clears the Parity error flag
func (r *RegisterIfcrType) SetPerrcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldPerrcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldPerrcfMask)
	}
}

const (
	RegisterIfcrFieldOvrcfShift = 3
	RegisterIfcrFieldOvrcfMask  = 0x8
)

// GetOvrcf Clears the Overrun error flag
func (r *RegisterIfcrType) GetOvrcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldOvrcfMask) != 0
}

// SetOvrcf Clears the Overrun error flag
func (r *RegisterIfcrType) SetOvrcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldOvrcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldOvrcfMask)
	}
}

const (
	RegisterIfcrFieldSbdcfShift = 4
	RegisterIfcrFieldSbdcfMask  = 0x10
)

// GetSbdcf Clears the Synchronization Block Detected flag
func (r *RegisterIfcrType) GetSbdcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldSbdcfMask) != 0
}

// SetSbdcf Clears the Synchronization Block Detected flag
func (r *RegisterIfcrType) SetSbdcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldSbdcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldSbdcfMask)
	}
}

const (
	RegisterIfcrFieldSyncdcfShift = 5
	RegisterIfcrFieldSyncdcfMask  = 0x20
)

// GetSyncdcf Clears the Synchronization Done flag
func (r *RegisterIfcrType) GetSyncdcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldSyncdcfMask) != 0
}

// SetSyncdcf Clears the Synchronization Done flag
func (r *RegisterIfcrType) SetSyncdcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldSyncdcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldSyncdcfMask)
	}
}

// RegisterDr00Type Data input register
type RegisterDr00Type uint32

func (r *RegisterDr00Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDr00Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDr00Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDr00Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDr00Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDr00FieldDrShift = 0
	RegisterDr00FieldDrMask  = 0xffffff
)

// GetDr Parity Error bit
func (r *RegisterDr00Type) GetDr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDr00FieldDrMask) >> RegisterDr00FieldDrShift)
}

// SetDr Parity Error bit
func (r *RegisterDr00Type) SetDr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr00FieldDrMask)|(uint32(value)<<RegisterDr00FieldDrShift))
}

const (
	RegisterDr00FieldPeShift = 24
	RegisterDr00FieldPeMask  = 0x1000000
)

// GetPe Parity Error bit
func (r *RegisterDr00Type) GetPe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr00FieldPeMask) != 0
}

// SetPe Parity Error bit
func (r *RegisterDr00Type) SetPe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr00FieldPeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr00FieldPeMask)
	}
}

const (
	RegisterDr00FieldVShift = 25
	RegisterDr00FieldVMask  = 0x2000000
)

// GetV Validity bit
func (r *RegisterDr00Type) GetV() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr00FieldVMask) != 0
}

// SetV Validity bit
func (r *RegisterDr00Type) SetV(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr00FieldVMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr00FieldVMask)
	}
}

const (
	RegisterDr00FieldUShift = 26
	RegisterDr00FieldUMask  = 0x4000000
)

// GetU User bit
func (r *RegisterDr00Type) GetU() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr00FieldUMask) != 0
}

// SetU User bit
func (r *RegisterDr00Type) SetU(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr00FieldUMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr00FieldUMask)
	}
}

const (
	RegisterDr00FieldCShift = 27
	RegisterDr00FieldCMask  = 0x8000000
)

// GetC Channel Status bit
func (r *RegisterDr00Type) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr00FieldCMask) != 0
}

// SetC Channel Status bit
func (r *RegisterDr00Type) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr00FieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr00FieldCMask)
	}
}

const (
	RegisterDr00FieldPtShift = 28
	RegisterDr00FieldPtMask  = 0x30000000
)

// GetPt Preamble Type
func (r *RegisterDr00Type) GetPt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDr00FieldPtMask) >> RegisterDr00FieldPtShift)
}

// SetPt Preamble Type
func (r *RegisterDr00Type) SetPt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr00FieldPtMask)|(uint32(value)<<RegisterDr00FieldPtShift))
}

// RegisterDr01Type Data input register
type RegisterDr01Type uint32

func (r *RegisterDr01Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDr01Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDr01Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDr01Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDr01Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDr01FieldPeShift = 0
	RegisterDr01FieldPeMask  = 0x1
)

// GetPe Parity Error bit
func (r *RegisterDr01Type) GetPe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr01FieldPeMask) != 0
}

// SetPe Parity Error bit
func (r *RegisterDr01Type) SetPe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr01FieldPeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr01FieldPeMask)
	}
}

const (
	RegisterDr01FieldVShift = 1
	RegisterDr01FieldVMask  = 0x2
)

// GetV Validity bit
func (r *RegisterDr01Type) GetV() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr01FieldVMask) != 0
}

// SetV Validity bit
func (r *RegisterDr01Type) SetV(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr01FieldVMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr01FieldVMask)
	}
}

const (
	RegisterDr01FieldUShift = 2
	RegisterDr01FieldUMask  = 0x4
)

// GetU User bit
func (r *RegisterDr01Type) GetU() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr01FieldUMask) != 0
}

// SetU User bit
func (r *RegisterDr01Type) SetU(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr01FieldUMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr01FieldUMask)
	}
}

const (
	RegisterDr01FieldCShift = 3
	RegisterDr01FieldCMask  = 0x8
)

// GetC Channel Status bit
func (r *RegisterDr01Type) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr01FieldCMask) != 0
}

// SetC Channel Status bit
func (r *RegisterDr01Type) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr01FieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr01FieldCMask)
	}
}

const (
	RegisterDr01FieldPtShift = 4
	RegisterDr01FieldPtMask  = 0x30
)

// GetPt Preamble Type
func (r *RegisterDr01Type) GetPt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDr01FieldPtMask) >> RegisterDr01FieldPtShift)
}

// SetPt Preamble Type
func (r *RegisterDr01Type) SetPt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr01FieldPtMask)|(uint32(value)<<RegisterDr01FieldPtShift))
}

const (
	RegisterDr01FieldDrShift = 8
	RegisterDr01FieldDrMask  = 0xffffff00
)

// GetDr Data value
func (r *RegisterDr01Type) GetDr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDr01FieldDrMask) >> RegisterDr01FieldDrShift)
}

// SetDr Data value
func (r *RegisterDr01Type) SetDr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr01FieldDrMask)|(uint32(value)<<RegisterDr01FieldDrShift))
}

// RegisterDr10Type Data input register
type RegisterDr10Type uint32

func (r *RegisterDr10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDr10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDr10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDr10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDr10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDr10FieldDrnl1Shift = 0
	RegisterDr10FieldDrnl1Mask  = 0xffff
)

// GetDrnl1 Data value
func (r *RegisterDr10Type) GetDrnl1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDr10FieldDrnl1Mask) >> RegisterDr10FieldDrnl1Shift)
}

// SetDrnl1 Data value
func (r *RegisterDr10Type) SetDrnl1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr10FieldDrnl1Mask)|(uint32(value)<<RegisterDr10FieldDrnl1Shift))
}

const (
	RegisterDr10FieldDrnl2Shift = 16
	RegisterDr10FieldDrnl2Mask  = 0xffff0000
)

// GetDrnl2 Data value
func (r *RegisterDr10Type) GetDrnl2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDr10FieldDrnl2Mask) >> RegisterDr10FieldDrnl2Shift)
}

// SetDrnl2 Data value
func (r *RegisterDr10Type) SetDrnl2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr10FieldDrnl2Mask)|(uint32(value)<<RegisterDr10FieldDrnl2Shift))
}

// RegisterCsrType Channel Status register
type RegisterCsrType uint32

func (r *RegisterCsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsrFieldUsrShift = 0
	RegisterCsrFieldUsrMask  = 0xffff
)

// GetUsr User data information
func (r *RegisterCsrType) GetUsr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldUsrMask) >> RegisterCsrFieldUsrShift)
}

// SetUsr User data information
func (r *RegisterCsrType) SetUsr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldUsrMask)|(uint32(value)<<RegisterCsrFieldUsrShift))
}

const (
	RegisterCsrFieldCsShift = 16
	RegisterCsrFieldCsMask  = 0xff0000
)

// GetCs Channel A status information
func (r *RegisterCsrType) GetCs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldCsMask) >> RegisterCsrFieldCsShift)
}

// SetCs Channel A status information
func (r *RegisterCsrType) SetCs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldCsMask)|(uint32(value)<<RegisterCsrFieldCsShift))
}

const (
	RegisterCsrFieldSobShift = 24
	RegisterCsrFieldSobMask  = 0x1000000
)

// GetSob Start Of Block
func (r *RegisterCsrType) GetSob() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldSobMask) != 0
}

// SetSob Start Of Block
func (r *RegisterCsrType) SetSob(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldSobMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldSobMask)
	}
}

// RegisterDirType Debug Information register
type RegisterDirType uint32

func (r *RegisterDirType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDirType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDirType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDirType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDirType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDirFieldThiShift = 0
	RegisterDirFieldThiMask  = 0x1fff
)

// GetThi Threshold HIGH
func (r *RegisterDirType) GetThi() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDirFieldThiMask) >> RegisterDirFieldThiShift)
}

// SetThi Threshold HIGH
func (r *RegisterDirType) SetThi(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDirFieldThiMask)|(uint32(value)<<RegisterDirFieldThiShift))
}

const (
	RegisterDirFieldTloShift = 16
	RegisterDirFieldTloMask  = 0x1fff0000
)

// GetTlo Threshold LOW
func (r *RegisterDirType) GetTlo() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDirFieldTloMask) >> RegisterDirFieldTloShift)
}

// SetTlo Threshold LOW
func (r *RegisterDirType) SetTlo(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDirFieldTloMask)|(uint32(value)<<RegisterDirFieldTloShift))
}

// RegisterVerrType SPDIFRX version register
type RegisterVerrType uint32

func (r *RegisterVerrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterVerrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterVerrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterVerrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterVerrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterVerrFieldMinrevShift = 0
	RegisterVerrFieldMinrevMask  = 0xf
)

// GetMinrev Minor revision
func (r *RegisterVerrType) GetMinrev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterVerrFieldMinrevMask) >> RegisterVerrFieldMinrevShift)
}

// SetMinrev Minor revision
func (r *RegisterVerrType) SetMinrev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterVerrFieldMinrevMask)|(uint32(value)<<RegisterVerrFieldMinrevShift))
}

const (
	RegisterVerrFieldMajrevShift = 4
	RegisterVerrFieldMajrevMask  = 0xf0
)

// GetMajrev Major revision
func (r *RegisterVerrType) GetMajrev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterVerrFieldMajrevMask) >> RegisterVerrFieldMajrevShift)
}

// SetMajrev Major revision
func (r *RegisterVerrType) SetMajrev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterVerrFieldMajrevMask)|(uint32(value)<<RegisterVerrFieldMajrevShift))
}

// RegisterIdrType SPDIFRX identification register
type RegisterIdrType uint32

func (r *RegisterIdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIdrFieldIdShift = 0
	RegisterIdrFieldIdMask  = 0xffffffff
)

// GetId SPDIFRX identifier
func (r *RegisterIdrType) GetId() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldIdMask) >> RegisterIdrFieldIdShift)
}

// SetId SPDIFRX identifier
func (r *RegisterIdrType) SetId(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldIdMask)|(uint32(value)<<RegisterIdrFieldIdShift))
}

// RegisterSidrType SPDIFRX size identification register
type RegisterSidrType uint32

func (r *RegisterSidrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSidrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSidrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSidrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSidrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSidrFieldSidShift = 0
	RegisterSidrFieldSidMask  = 0xffffffff
)

// GetSid Size identification
func (r *RegisterSidrType) GetSid() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterSidrFieldSidMask) >> RegisterSidrFieldSidShift)
}

// SetSid Size identification
func (r *RegisterSidrType) SetSid(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSidrFieldSidMask)|(uint32(value)<<RegisterSidrFieldSidShift))
}
