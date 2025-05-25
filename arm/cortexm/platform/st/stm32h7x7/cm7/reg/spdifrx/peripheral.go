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
	Cr    registerCrType
	Imr   registerImrType
	Sr    registerSrType
	Ifcr  registerIfcrType
	Dr_00 registerDr_00Type
	Dr_01 registerDr_01Type
	Dr_10 registerDr_10Type
	Csr   registerCsrType
	Dir   registerDirType
	_     [984]byte
	Verr  registerVerrType
	Idr   registerIdrType
	Sidr  registerSidrType
}

// registerCrType Control register
type registerCrType uint32

const (
	RegisterCrFieldSpdifrxenShift = 0
	RegisterCrFieldSpdifrxenMask  = 0x3
)

// GetSpdifrxen Peripheral Block Enable
func (r *registerCrType) GetSpdifrxen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSpdifrxenMask) >> RegisterCrFieldSpdifrxenShift)
}

// SetSpdifrxen Peripheral Block Enable
func (r *registerCrType) SetSpdifrxen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSpdifrxenMask)|(uint32(value)<<RegisterCrFieldSpdifrxenShift))
}

const (
	RegisterCrFieldRxdmaenShift = 2
	RegisterCrFieldRxdmaenMask  = 0x4
)

// GetRxdmaen Receiver DMA ENable for data flow
func (r *registerCrType) GetRxdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRxdmaenMask) != 0
}

// SetRxdmaen Receiver DMA ENable for data flow
func (r *registerCrType) SetRxdmaen(value bool) {
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
func (r *registerCrType) GetRxsteo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRxsteoMask) != 0
}

// SetRxsteo STerEO Mode
func (r *registerCrType) SetRxsteo(value bool) {
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
func (r *registerCrType) GetDrfmt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDrfmtMask) >> RegisterCrFieldDrfmtShift)
}

// SetDrfmt RX Data format
func (r *registerCrType) SetDrfmt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDrfmtMask)|(uint32(value)<<RegisterCrFieldDrfmtShift))
}

const (
	RegisterCrFieldPmskShift = 6
	RegisterCrFieldPmskMask  = 0x40
)

// GetPmsk Mask Parity error bit
func (r *registerCrType) GetPmsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPmskMask) != 0
}

// SetPmsk Mask Parity error bit
func (r *registerCrType) SetPmsk(value bool) {
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
func (r *registerCrType) GetVmsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldVmskMask) != 0
}

// SetVmsk Mask of Validity bit
func (r *registerCrType) SetVmsk(value bool) {
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
func (r *registerCrType) GetCumsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCumskMask) != 0
}

// SetCumsk Mask of channel status and user bits
func (r *registerCrType) SetCumsk(value bool) {
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
func (r *registerCrType) GetPtmsk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPtmskMask) != 0
}

// SetPtmsk Mask of Preamble Type bits
func (r *registerCrType) SetPtmsk(value bool) {
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
func (r *registerCrType) GetCbdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCbdmaenMask) != 0
}

// SetCbdmaen Control Buffer DMA ENable for control flow
func (r *registerCrType) SetCbdmaen(value bool) {
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
func (r *registerCrType) GetChsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldChselMask) != 0
}

// SetChsel Channel Selection
func (r *registerCrType) SetChsel(value bool) {
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
func (r *registerCrType) GetNbtr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldNbtrMask) >> RegisterCrFieldNbtrShift)
}

// SetNbtr Maximum allowed re-tries during synchronization phase
func (r *registerCrType) SetNbtr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldNbtrMask)|(uint32(value)<<RegisterCrFieldNbtrShift))
}

const (
	RegisterCrFieldWfaShift = 14
	RegisterCrFieldWfaMask  = 0x4000
)

// GetWfa Wait For Activity
func (r *registerCrType) GetWfa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldWfaMask) != 0
}

// SetWfa Wait For Activity
func (r *registerCrType) SetWfa(value bool) {
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
func (r *registerCrType) GetInsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldInselMask) >> RegisterCrFieldInselShift)
}

// SetInsel input selection
func (r *registerCrType) SetInsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldInselMask)|(uint32(value)<<RegisterCrFieldInselShift))
}

const (
	RegisterCrFieldCksenShift = 20
	RegisterCrFieldCksenMask  = 0x100000
)

// GetCksen Symbol Clock Enable
func (r *registerCrType) GetCksen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCksenMask) != 0
}

// SetCksen Symbol Clock Enable
func (r *registerCrType) SetCksen(value bool) {
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
func (r *registerCrType) GetCksbkpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCksbkpenMask) != 0
}

// SetCksbkpen Backup Symbol Clock Enable
func (r *registerCrType) SetCksbkpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCksbkpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCksbkpenMask)
	}
}

// registerImrType Interrupt mask register
type registerImrType uint32

const (
	RegisterImrFieldRxneieShift = 0
	RegisterImrFieldRxneieMask  = 0x1
)

// GetRxneie RXNE interrupt enable
func (r *registerImrType) GetRxneie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldRxneieMask) != 0
}

// SetRxneie RXNE interrupt enable
func (r *registerImrType) SetRxneie(value bool) {
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
func (r *registerImrType) GetCsrneie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldCsrneieMask) != 0
}

// SetCsrneie Control Buffer Ready Interrupt Enable
func (r *registerImrType) SetCsrneie(value bool) {
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
func (r *registerImrType) GetPerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldPerrieMask) != 0
}

// SetPerrie Parity error interrupt enable
func (r *registerImrType) SetPerrie(value bool) {
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
func (r *registerImrType) GetOvrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldOvrieMask) != 0
}

// SetOvrie Overrun error Interrupt Enable
func (r *registerImrType) SetOvrie(value bool) {
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
func (r *registerImrType) GetSblkie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldSblkieMask) != 0
}

// SetSblkie Synchronization Block Detected Interrupt Enable
func (r *registerImrType) SetSblkie(value bool) {
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
func (r *registerImrType) GetSyncdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldSyncdieMask) != 0
}

// SetSyncdie Synchronization Done
func (r *registerImrType) SetSyncdie(value bool) {
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
func (r *registerImrType) GetIfeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterImrFieldIfeieMask) != 0
}

// SetIfeie Serial Interface Error Interrupt Enable
func (r *registerImrType) SetIfeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterImrFieldIfeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterImrFieldIfeieMask)
	}
}

// registerSrType Status register
type registerSrType uint32

const (
	RegisterSrFieldRxneShift = 0
	RegisterSrFieldRxneMask  = 0x1
)

// GetRxne Read data register not empty
func (r *registerSrType) GetRxne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRxneMask) != 0
}

// SetRxne Read data register not empty
func (r *registerSrType) SetRxne(value bool) {
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
func (r *registerSrType) GetCsrne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCsrneMask) != 0
}

// SetCsrne Control Buffer register is not empty
func (r *registerSrType) SetCsrne(value bool) {
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
func (r *registerSrType) GetPerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldPerrMask) != 0
}

// SetPerr Parity error
func (r *registerSrType) SetPerr(value bool) {
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
func (r *registerSrType) GetOvr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldOvrMask) != 0
}

// SetOvr Overrun error
func (r *registerSrType) SetOvr(value bool) {
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
func (r *registerSrType) GetSbd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSbdMask) != 0
}

// SetSbd Synchronization Block Detected
func (r *registerSrType) SetSbd(value bool) {
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
func (r *registerSrType) GetSyncd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSyncdMask) != 0
}

// SetSyncd Synchronization Done
func (r *registerSrType) SetSyncd(value bool) {
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
func (r *registerSrType) GetFerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldFerrMask) != 0
}

// SetFerr Framing error
func (r *registerSrType) SetFerr(value bool) {
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
func (r *registerSrType) GetSerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSerrMask) != 0
}

// SetSerr Synchronization error
func (r *registerSrType) SetSerr(value bool) {
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
func (r *registerSrType) GetTerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTerrMask) != 0
}

// SetTerr Time-out error
func (r *registerSrType) SetTerr(value bool) {
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
func (r *registerSrType) GetWidth5() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldWidth5Mask) >> RegisterSrFieldWidth5Shift)
}

// SetWidth5 Duration of 5 symbols counted with SPDIF_CLK
func (r *registerSrType) SetWidth5(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldWidth5Mask)|(uint32(value)<<RegisterSrFieldWidth5Shift))
}

// registerIfcrType Interrupt Flag Clear register
type registerIfcrType uint32

const (
	RegisterIfcrFieldPerrcfShift = 2
	RegisterIfcrFieldPerrcfMask  = 0x4
)

// GetPerrcf Clears the Parity error flag
func (r *registerIfcrType) GetPerrcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldPerrcfMask) != 0
}

// SetPerrcf Clears the Parity error flag
func (r *registerIfcrType) SetPerrcf(value bool) {
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
func (r *registerIfcrType) GetOvrcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldOvrcfMask) != 0
}

// SetOvrcf Clears the Overrun error flag
func (r *registerIfcrType) SetOvrcf(value bool) {
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
func (r *registerIfcrType) GetSbdcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldSbdcfMask) != 0
}

// SetSbdcf Clears the Synchronization Block Detected flag
func (r *registerIfcrType) SetSbdcf(value bool) {
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
func (r *registerIfcrType) GetSyncdcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldSyncdcfMask) != 0
}

// SetSyncdcf Clears the Synchronization Done flag
func (r *registerIfcrType) SetSyncdcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldSyncdcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldSyncdcfMask)
	}
}

// registerDr_00Type Data input register
type registerDr_00Type uint32

const (
	RegisterDr_00FieldDrShift = 0
	RegisterDr_00FieldDrMask  = 0xffffff
)

// GetDr Parity Error bit
func (r *registerDr_00Type) GetDr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDr_00FieldDrMask) >> RegisterDr_00FieldDrShift)
}

// SetDr Parity Error bit
func (r *registerDr_00Type) SetDr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr_00FieldDrMask)|(uint32(value)<<RegisterDr_00FieldDrShift))
}

const (
	RegisterDr_00FieldPeShift = 24
	RegisterDr_00FieldPeMask  = 0x1000000
)

// GetPe Parity Error bit
func (r *registerDr_00Type) GetPe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr_00FieldPeMask) != 0
}

// SetPe Parity Error bit
func (r *registerDr_00Type) SetPe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr_00FieldPeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr_00FieldPeMask)
	}
}

const (
	RegisterDr_00FieldVShift = 25
	RegisterDr_00FieldVMask  = 0x2000000
)

// GetV Validity bit
func (r *registerDr_00Type) GetV() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr_00FieldVMask) != 0
}

// SetV Validity bit
func (r *registerDr_00Type) SetV(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr_00FieldVMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr_00FieldVMask)
	}
}

const (
	RegisterDr_00FieldUShift = 26
	RegisterDr_00FieldUMask  = 0x4000000
)

// GetU User bit
func (r *registerDr_00Type) GetU() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr_00FieldUMask) != 0
}

// SetU User bit
func (r *registerDr_00Type) SetU(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr_00FieldUMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr_00FieldUMask)
	}
}

const (
	RegisterDr_00FieldCShift = 27
	RegisterDr_00FieldCMask  = 0x8000000
)

// GetC Channel Status bit
func (r *registerDr_00Type) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr_00FieldCMask) != 0
}

// SetC Channel Status bit
func (r *registerDr_00Type) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr_00FieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr_00FieldCMask)
	}
}

const (
	RegisterDr_00FieldPtShift = 28
	RegisterDr_00FieldPtMask  = 0x30000000
)

// GetPt Preamble Type
func (r *registerDr_00Type) GetPt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDr_00FieldPtMask) >> RegisterDr_00FieldPtShift)
}

// SetPt Preamble Type
func (r *registerDr_00Type) SetPt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr_00FieldPtMask)|(uint32(value)<<RegisterDr_00FieldPtShift))
}

// registerDr_01Type Data input register
type registerDr_01Type uint32

const (
	RegisterDr_01FieldPeShift = 0
	RegisterDr_01FieldPeMask  = 0x1
)

// GetPe Parity Error bit
func (r *registerDr_01Type) GetPe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr_01FieldPeMask) != 0
}

// SetPe Parity Error bit
func (r *registerDr_01Type) SetPe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr_01FieldPeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr_01FieldPeMask)
	}
}

const (
	RegisterDr_01FieldVShift = 1
	RegisterDr_01FieldVMask  = 0x2
)

// GetV Validity bit
func (r *registerDr_01Type) GetV() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr_01FieldVMask) != 0
}

// SetV Validity bit
func (r *registerDr_01Type) SetV(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr_01FieldVMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr_01FieldVMask)
	}
}

const (
	RegisterDr_01FieldUShift = 2
	RegisterDr_01FieldUMask  = 0x4
)

// GetU User bit
func (r *registerDr_01Type) GetU() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr_01FieldUMask) != 0
}

// SetU User bit
func (r *registerDr_01Type) SetU(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr_01FieldUMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr_01FieldUMask)
	}
}

const (
	RegisterDr_01FieldCShift = 3
	RegisterDr_01FieldCMask  = 0x8
)

// GetC Channel Status bit
func (r *registerDr_01Type) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDr_01FieldCMask) != 0
}

// SetC Channel Status bit
func (r *registerDr_01Type) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDr_01FieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDr_01FieldCMask)
	}
}

const (
	RegisterDr_01FieldPtShift = 4
	RegisterDr_01FieldPtMask  = 0x30
)

// GetPt Preamble Type
func (r *registerDr_01Type) GetPt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDr_01FieldPtMask) >> RegisterDr_01FieldPtShift)
}

// SetPt Preamble Type
func (r *registerDr_01Type) SetPt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr_01FieldPtMask)|(uint32(value)<<RegisterDr_01FieldPtShift))
}

const (
	RegisterDr_01FieldDrShift = 8
	RegisterDr_01FieldDrMask  = 0xffffff00
)

// GetDr Data value
func (r *registerDr_01Type) GetDr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDr_01FieldDrMask) >> RegisterDr_01FieldDrShift)
}

// SetDr Data value
func (r *registerDr_01Type) SetDr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr_01FieldDrMask)|(uint32(value)<<RegisterDr_01FieldDrShift))
}

// registerDr_10Type Data input register
type registerDr_10Type uint32

const (
	RegisterDr_10FieldDrnl1Shift = 0
	RegisterDr_10FieldDrnl1Mask  = 0xffff
)

// GetDrnl1 Data value
func (r *registerDr_10Type) GetDrnl1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDr_10FieldDrnl1Mask) >> RegisterDr_10FieldDrnl1Shift)
}

// SetDrnl1 Data value
func (r *registerDr_10Type) SetDrnl1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr_10FieldDrnl1Mask)|(uint32(value)<<RegisterDr_10FieldDrnl1Shift))
}

const (
	RegisterDr_10FieldDrnl2Shift = 16
	RegisterDr_10FieldDrnl2Mask  = 0xffff0000
)

// GetDrnl2 Data value
func (r *registerDr_10Type) GetDrnl2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDr_10FieldDrnl2Mask) >> RegisterDr_10FieldDrnl2Shift)
}

// SetDrnl2 Data value
func (r *registerDr_10Type) SetDrnl2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDr_10FieldDrnl2Mask)|(uint32(value)<<RegisterDr_10FieldDrnl2Shift))
}

// registerCsrType Channel Status register
type registerCsrType uint32

const (
	RegisterCsrFieldUsrShift = 0
	RegisterCsrFieldUsrMask  = 0xffff
)

// GetUsr User data information
func (r *registerCsrType) GetUsr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldUsrMask) >> RegisterCsrFieldUsrShift)
}

// SetUsr User data information
func (r *registerCsrType) SetUsr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldUsrMask)|(uint32(value)<<RegisterCsrFieldUsrShift))
}

const (
	RegisterCsrFieldCsShift = 16
	RegisterCsrFieldCsMask  = 0xff0000
)

// GetCs Channel A status information
func (r *registerCsrType) GetCs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldCsMask) >> RegisterCsrFieldCsShift)
}

// SetCs Channel A status information
func (r *registerCsrType) SetCs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldCsMask)|(uint32(value)<<RegisterCsrFieldCsShift))
}

const (
	RegisterCsrFieldSobShift = 24
	RegisterCsrFieldSobMask  = 0x1000000
)

// GetSob Start Of Block
func (r *registerCsrType) GetSob() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldSobMask) != 0
}

// SetSob Start Of Block
func (r *registerCsrType) SetSob(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldSobMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldSobMask)
	}
}

// registerDirType Debug Information register
type registerDirType uint32

const (
	RegisterDirFieldThiShift = 0
	RegisterDirFieldThiMask  = 0x1fff
)

// GetThi Threshold HIGH
func (r *registerDirType) GetThi() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDirFieldThiMask) >> RegisterDirFieldThiShift)
}

// SetThi Threshold HIGH
func (r *registerDirType) SetThi(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDirFieldThiMask)|(uint32(value)<<RegisterDirFieldThiShift))
}

const (
	RegisterDirFieldTloShift = 16
	RegisterDirFieldTloMask  = 0x1fff0000
)

// GetTlo Threshold LOW
func (r *registerDirType) GetTlo() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDirFieldTloMask) >> RegisterDirFieldTloShift)
}

// SetTlo Threshold LOW
func (r *registerDirType) SetTlo(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDirFieldTloMask)|(uint32(value)<<RegisterDirFieldTloShift))
}

// registerVerrType SPDIFRX version register
type registerVerrType uint32

const (
	RegisterVerrFieldMinrevShift = 0
	RegisterVerrFieldMinrevMask  = 0xf
)

// GetMinrev Minor revision
func (r *registerVerrType) GetMinrev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterVerrFieldMinrevMask) >> RegisterVerrFieldMinrevShift)
}

// SetMinrev Minor revision
func (r *registerVerrType) SetMinrev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterVerrFieldMinrevMask)|(uint32(value)<<RegisterVerrFieldMinrevShift))
}

const (
	RegisterVerrFieldMajrevShift = 4
	RegisterVerrFieldMajrevMask  = 0xf0
)

// GetMajrev Major revision
func (r *registerVerrType) GetMajrev() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterVerrFieldMajrevMask) >> RegisterVerrFieldMajrevShift)
}

// SetMajrev Major revision
func (r *registerVerrType) SetMajrev(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterVerrFieldMajrevMask)|(uint32(value)<<RegisterVerrFieldMajrevShift))
}

// registerIdrType SPDIFRX identification register
type registerIdrType uint32

const (
	RegisterIdrFieldIdShift = 0
	RegisterIdrFieldIdMask  = 0xffffffff
)

// GetId SPDIFRX identifier
func (r *registerIdrType) GetId() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIdrFieldIdMask) >> RegisterIdrFieldIdShift)
}

// SetId SPDIFRX identifier
func (r *registerIdrType) SetId(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdrFieldIdMask)|(uint32(value)<<RegisterIdrFieldIdShift))
}

// registerSidrType SPDIFRX size identification register
type registerSidrType uint32

const (
	RegisterSidrFieldSidShift = 0
	RegisterSidrFieldSidMask  = 0xffffffff
)

// GetSid Size identification
func (r *registerSidrType) GetSid() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterSidrFieldSidMask) >> RegisterSidrFieldSidShift)
}

// SetSid Size identification
func (r *registerSidrType) SetSid(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSidrFieldSidMask)|(uint32(value)<<RegisterSidrFieldSidShift))
}
