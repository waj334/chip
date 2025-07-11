//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package quadspi

import (
	"unsafe"
	"volatile"
)

var (
	Quadspi = (*_quadspi)(unsafe.Pointer(uintptr(0x52005000)))
)

type _quadspi struct {
	Cr    RegisterCrType
	Dcr   RegisterDcrType
	Sr    RegisterSrType
	Fcr   RegisterFcrType
	Dlr   RegisterDlrType
	Ccr   RegisterCcrType
	Ar    RegisterArType
	Abr   RegisterAbrType
	Dr    RegisterDrType
	Psmkr RegisterPsmkrType
	Psmar RegisterPsmarType
	Pir   RegisterPirType
	Lptr  RegisterLptrType
}

// RegisterCrType QUADSPI control register
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
	RegisterCrFieldEnShift = 0
	RegisterCrFieldEnMask  = 0x1
)

// GetEn Enable Enable the QUADSPI.
func (r *RegisterCrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEnMask) != 0
}

// SetEn Enable Enable the QUADSPI.
func (r *RegisterCrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEnMask)
	}
}

const (
	RegisterCrFieldAbortShift = 1
	RegisterCrFieldAbortMask  = 0x2
)

// GetAbort Abort request This bit aborts the on-going command sequence. It is automatically reset once the abort is complete. This bit stops the current transfer. In polling mode or memory-mapped mode, this bit also reset the APM bit or the DM bit.
func (r *RegisterCrType) GetAbort() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAbortMask) != 0
}

// SetAbort Abort request This bit aborts the on-going command sequence. It is automatically reset once the abort is complete. This bit stops the current transfer. In polling mode or memory-mapped mode, this bit also reset the APM bit or the DM bit.
func (r *RegisterCrType) SetAbort(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAbortMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAbortMask)
	}
}

const (
	RegisterCrFieldDmaenShift = 2
	RegisterCrFieldDmaenMask  = 0x4
)

// GetDmaen DMA enable In indirect mode, DMA can be used to input or output data via the QUADSPI_DR register. DMA transfers are initiated when the FIFO threshold flag, FTF, is set.
func (r *RegisterCrType) GetDmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDmaenMask) != 0
}

// SetDmaen DMA enable In indirect mode, DMA can be used to input or output data via the QUADSPI_DR register. DMA transfers are initiated when the FIFO threshold flag, FTF, is set.
func (r *RegisterCrType) SetDmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDmaenMask)
	}
}

const (
	RegisterCrFieldTcenShift = 3
	RegisterCrFieldTcenMask  = 0x8
)

// GetTcen Timeout counter enable This bit is valid only when memory-mapped mode (FMODE = 11) is selected. Activating this bit causes the chip select (nCS) to be released (and thus reduces consumption) if there has not been an access after a certain amount of time, where this time is defined by TIMEOUT[15:0] (QUADSPI_LPTR). Enable the timeout counter. By default, the QUADSPI never stops its prefetch operation, keeping the previous read operation active with nCS maintained low, even if no access to the Flash memory occurs for a long time. Since Flash memories tend to consume more when nCS is held low, the application might want to activate the timeout counter (TCEN = 1, QUADSPI_CR[3]) so that nCS is released after a period of TIMEOUT[15:0] (QUADSPI_LPTR) cycles have elapsed without an access since when the FIFO becomes full with prefetch data. This bit can be modified only when BUSY = 0.
func (r *RegisterCrType) GetTcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTcenMask) != 0
}

// SetTcen Timeout counter enable This bit is valid only when memory-mapped mode (FMODE = 11) is selected. Activating this bit causes the chip select (nCS) to be released (and thus reduces consumption) if there has not been an access after a certain amount of time, where this time is defined by TIMEOUT[15:0] (QUADSPI_LPTR). Enable the timeout counter. By default, the QUADSPI never stops its prefetch operation, keeping the previous read operation active with nCS maintained low, even if no access to the Flash memory occurs for a long time. Since Flash memories tend to consume more when nCS is held low, the application might want to activate the timeout counter (TCEN = 1, QUADSPI_CR[3]) so that nCS is released after a period of TIMEOUT[15:0] (QUADSPI_LPTR) cycles have elapsed without an access since when the FIFO becomes full with prefetch data. This bit can be modified only when BUSY = 0.
func (r *RegisterCrType) SetTcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTcenMask)
	}
}

const (
	RegisterCrFieldSshiftShift = 4
	RegisterCrFieldSshiftMask  = 0x10
)

// GetSshift Sample shift By default, the QUADSPI samples data 1/2 of a CLK cycle after the data is driven by the Flash memory. This bit allows the data is to be sampled later in order to account for external signal delays. Firmware must assure that SSHIFT = 0 when in DDR mode (when DDRM = 1). This field can be modified only when BUSY = 0.
func (r *RegisterCrType) GetSshift() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSshiftMask) != 0
}

// SetSshift Sample shift By default, the QUADSPI samples data 1/2 of a CLK cycle after the data is driven by the Flash memory. This bit allows the data is to be sampled later in order to account for external signal delays. Firmware must assure that SSHIFT = 0 when in DDR mode (when DDRM = 1). This field can be modified only when BUSY = 0.
func (r *RegisterCrType) SetSshift(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSshiftMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSshiftMask)
	}
}

const (
	RegisterCrFieldDfmShift = 6
	RegisterCrFieldDfmMask  = 0x40
)

// GetDfm Dual-flash mode This bit activates dual-flash mode, where two external Flash memories are used simultaneously to double throughput and capacity. This bit can be modified only when BUSY = 0.
func (r *RegisterCrType) GetDfm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDfmMask) != 0
}

// SetDfm Dual-flash mode This bit activates dual-flash mode, where two external Flash memories are used simultaneously to double throughput and capacity. This bit can be modified only when BUSY = 0.
func (r *RegisterCrType) SetDfm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDfmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDfmMask)
	}
}

const (
	RegisterCrFieldFselShift = 7
	RegisterCrFieldFselMask  = 0x80
)

// GetFsel Flash memory selection This bit selects the Flash memory to be addressed in single flash mode (when DFM = 0). This bit can be modified only when BUSY = 0. This bit is ignored when DFM = 1.
func (r *RegisterCrType) GetFsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldFselMask) != 0
}

// SetFsel Flash memory selection This bit selects the Flash memory to be addressed in single flash mode (when DFM = 0). This bit can be modified only when BUSY = 0. This bit is ignored when DFM = 1.
func (r *RegisterCrType) SetFsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldFselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldFselMask)
	}
}

const (
	RegisterCrFieldFthresShift = 8
	RegisterCrFieldFthresMask  = 0x1f00
)

// GetFthres FIFO threshold level Defines, in indirect mode, the threshold number of bytes in the FIFO that will cause the FIFO threshold flag (FTF, QUADSPI_SR[2]) to be set. In indirect write mode (FMODE = 00): ... In indirect read mode (FMODE = 01): ... If DMAEN = 1, then the DMA controller for the corresponding channel must be disabled before changing the FTHRES value.
func (r *RegisterCrType) GetFthres() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldFthresMask) >> RegisterCrFieldFthresShift)
}

// SetFthres FIFO threshold level Defines, in indirect mode, the threshold number of bytes in the FIFO that will cause the FIFO threshold flag (FTF, QUADSPI_SR[2]) to be set. In indirect write mode (FMODE = 00): ... In indirect read mode (FMODE = 01): ... If DMAEN = 1, then the DMA controller for the corresponding channel must be disabled before changing the FTHRES value.
func (r *RegisterCrType) SetFthres(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldFthresMask)|(uint32(value)<<RegisterCrFieldFthresShift))
}

const (
	RegisterCrFieldTeieShift = 16
	RegisterCrFieldTeieMask  = 0x10000
)

// GetTeie Transfer error interrupt enable This bit enables the transfer error interrupt.
func (r *RegisterCrType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTeieMask) != 0
}

// SetTeie Transfer error interrupt enable This bit enables the transfer error interrupt.
func (r *RegisterCrType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTeieMask)
	}
}

const (
	RegisterCrFieldTcieShift = 17
	RegisterCrFieldTcieMask  = 0x20000
)

// GetTcie Transfer complete interrupt enable This bit enables the transfer complete interrupt.
func (r *RegisterCrType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTcieMask) != 0
}

// SetTcie Transfer complete interrupt enable This bit enables the transfer complete interrupt.
func (r *RegisterCrType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTcieMask)
	}
}

const (
	RegisterCrFieldFtieShift = 18
	RegisterCrFieldFtieMask  = 0x40000
)

// GetFtie FIFO threshold interrupt enable This bit enables the FIFO threshold interrupt.
func (r *RegisterCrType) GetFtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldFtieMask) != 0
}

// SetFtie FIFO threshold interrupt enable This bit enables the FIFO threshold interrupt.
func (r *RegisterCrType) SetFtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldFtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldFtieMask)
	}
}

const (
	RegisterCrFieldSmieShift = 19
	RegisterCrFieldSmieMask  = 0x80000
)

// GetSmie Status match interrupt enable This bit enables the status match interrupt.
func (r *RegisterCrType) GetSmie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSmieMask) != 0
}

// SetSmie Status match interrupt enable This bit enables the status match interrupt.
func (r *RegisterCrType) SetSmie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSmieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSmieMask)
	}
}

const (
	RegisterCrFieldToieShift = 20
	RegisterCrFieldToieMask  = 0x100000
)

// GetToie TimeOut interrupt enable This bit enables the TimeOut interrupt.
func (r *RegisterCrType) GetToie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldToieMask) != 0
}

// SetToie TimeOut interrupt enable This bit enables the TimeOut interrupt.
func (r *RegisterCrType) SetToie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldToieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldToieMask)
	}
}

const (
	RegisterCrFieldApmsShift = 22
	RegisterCrFieldApmsMask  = 0x400000
)

// GetApms Automatic poll mode stop This bit determines if automatic polling is stopped after a match. This bit can be modified only when BUSY = 0.
func (r *RegisterCrType) GetApms() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldApmsMask) != 0
}

// SetApms Automatic poll mode stop This bit determines if automatic polling is stopped after a match. This bit can be modified only when BUSY = 0.
func (r *RegisterCrType) SetApms(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldApmsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldApmsMask)
	}
}

const (
	RegisterCrFieldPmmShift = 23
	RegisterCrFieldPmmMask  = 0x800000
)

// GetPmm Polling match mode This bit indicates which method should be used for determining a match during automatic polling mode. This bit can be modified only when BUSY = 0.
func (r *RegisterCrType) GetPmm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPmmMask) != 0
}

// SetPmm Polling match mode This bit indicates which method should be used for determining a match during automatic polling mode. This bit can be modified only when BUSY = 0.
func (r *RegisterCrType) SetPmm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPmmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPmmMask)
	}
}

const (
	RegisterCrFieldPrescalerShift = 24
	RegisterCrFieldPrescalerMask  = 0xff000000
)

// GetPrescaler clock prescaler
func (r *RegisterCrType) GetPrescaler() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPrescalerMask) >> RegisterCrFieldPrescalerShift)
}

// SetPrescaler clock prescaler
func (r *RegisterCrType) SetPrescaler(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPrescalerMask)|(uint32(value)<<RegisterCrFieldPrescalerShift))
}

// RegisterDcrType QUADSPI device configuration register
type RegisterDcrType uint32

func (r *RegisterDcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDcrFieldCkmodeShift = 0
	RegisterDcrFieldCkmodeMask  = 0x1
)

// GetCkmode indicates the level that clk takes between command
func (r *RegisterDcrType) GetCkmode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDcrFieldCkmodeMask) != 0
}

// SetCkmode indicates the level that clk takes between command
func (r *RegisterDcrType) SetCkmode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDcrFieldCkmodeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDcrFieldCkmodeMask)
	}
}

const (
	RegisterDcrFieldCshtShift = 8
	RegisterDcrFieldCshtMask  = 0x700
)

// GetCsht Chip select high time CSHT+1 defines the minimum number of CLK cycles which the chip select (nCS) must remain high between commands issued to the Flash memory. ... This field can be modified only when BUSY = 0.
func (r *RegisterDcrType) GetCsht() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcrFieldCshtMask) >> RegisterDcrFieldCshtShift)
}

// SetCsht Chip select high time CSHT+1 defines the minimum number of CLK cycles which the chip select (nCS) must remain high between commands issued to the Flash memory. ... This field can be modified only when BUSY = 0.
func (r *RegisterDcrType) SetCsht(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcrFieldCshtMask)|(uint32(value)<<RegisterDcrFieldCshtShift))
}

const (
	RegisterDcrFieldFsizeShift = 16
	RegisterDcrFieldFsizeMask  = 0x1f0000
)

// GetFsize Flash memory size This field defines the size of external memory using the following formula: Number of bytes in Flash memory = 2[FSIZE+1] FSIZE+1 is effectively the number of address bits required to address the Flash memory. The Flash memory capacity can be up to 4GB (addressed using 32 bits) in indirect mode, but the addressable space in memory-mapped mode is limited to 256MB. If DFM = 1, FSIZE indicates the total capacity of the two Flash memories together. This field can be modified only when BUSY = 0.
func (r *RegisterDcrType) GetFsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDcrFieldFsizeMask) >> RegisterDcrFieldFsizeShift)
}

// SetFsize Flash memory size This field defines the size of external memory using the following formula: Number of bytes in Flash memory = 2[FSIZE+1] FSIZE+1 is effectively the number of address bits required to address the Flash memory. The Flash memory capacity can be up to 4GB (addressed using 32 bits) in indirect mode, but the addressable space in memory-mapped mode is limited to 256MB. If DFM = 1, FSIZE indicates the total capacity of the two Flash memories together. This field can be modified only when BUSY = 0.
func (r *RegisterDcrType) SetFsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcrFieldFsizeMask)|(uint32(value)<<RegisterDcrFieldFsizeShift))
}

// RegisterSrType QUADSPI status register
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
	RegisterSrFieldTefShift = 0
	RegisterSrFieldTefMask  = 0x1
)

// GetTef Transfer error flag This bit is set in indirect mode when an invalid address is being accessed in indirect mode. It is cleared by writing 1 to CTEF.
func (r *RegisterSrType) GetTef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTefMask) != 0
}

// SetTef Transfer error flag This bit is set in indirect mode when an invalid address is being accessed in indirect mode. It is cleared by writing 1 to CTEF.
func (r *RegisterSrType) SetTef(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldTefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldTefMask)
	}
}

const (
	RegisterSrFieldTcfShift = 1
	RegisterSrFieldTcfMask  = 0x2
)

// GetTcf Transfer complete flag This bit is set in indirect mode when the programmed number of data has been transferred or in any mode when the transfer has been aborted.It is cleared by writing 1 to CTCF.
func (r *RegisterSrType) GetTcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTcfMask) != 0
}

// SetTcf Transfer complete flag This bit is set in indirect mode when the programmed number of data has been transferred or in any mode when the transfer has been aborted.It is cleared by writing 1 to CTCF.
func (r *RegisterSrType) SetTcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldTcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldTcfMask)
	}
}

const (
	RegisterSrFieldFtfShift = 2
	RegisterSrFieldFtfMask  = 0x4
)

// GetFtf FIFO threshold flag In indirect mode, this bit is set when the FIFO threshold has been reached, or if there is any data left in the FIFO after reads from the Flash memory are complete. It is cleared automatically as soon as threshold condition is no longer true. In automatic polling mode this bit is set every time the status register is read, and the bit is cleared when the data register is read.
func (r *RegisterSrType) GetFtf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldFtfMask) != 0
}

// SetFtf FIFO threshold flag In indirect mode, this bit is set when the FIFO threshold has been reached, or if there is any data left in the FIFO after reads from the Flash memory are complete. It is cleared automatically as soon as threshold condition is no longer true. In automatic polling mode this bit is set every time the status register is read, and the bit is cleared when the data register is read.
func (r *RegisterSrType) SetFtf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldFtfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldFtfMask)
	}
}

const (
	RegisterSrFieldSmfShift = 3
	RegisterSrFieldSmfMask  = 0x8
)

// GetSmf Status match flag This bit is set in automatic polling mode when the unmasked received data matches the corresponding bits in the match register (QUADSPI_PSMAR). It is cleared by writing 1 to CSMF.
func (r *RegisterSrType) GetSmf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSmfMask) != 0
}

// SetSmf Status match flag This bit is set in automatic polling mode when the unmasked received data matches the corresponding bits in the match register (QUADSPI_PSMAR). It is cleared by writing 1 to CSMF.
func (r *RegisterSrType) SetSmf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldSmfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldSmfMask)
	}
}

const (
	RegisterSrFieldTofShift = 4
	RegisterSrFieldTofMask  = 0x10
)

// GetTof Timeout flag This bit is set when timeout occurs. It is cleared by writing 1 to CTOF.
func (r *RegisterSrType) GetTof() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTofMask) != 0
}

// SetTof Timeout flag This bit is set when timeout occurs. It is cleared by writing 1 to CTOF.
func (r *RegisterSrType) SetTof(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldTofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldTofMask)
	}
}

const (
	RegisterSrFieldBusyShift = 5
	RegisterSrFieldBusyMask  = 0x20
)

// GetBusy Busy This bit is set when an operation is on going. This bit clears automatically when the operation with the Flash memory is finished and the FIFO is empty.
func (r *RegisterSrType) GetBusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldBusyMask) != 0
}

// SetBusy Busy This bit is set when an operation is on going. This bit clears automatically when the operation with the Flash memory is finished and the FIFO is empty.
func (r *RegisterSrType) SetBusy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldBusyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldBusyMask)
	}
}

const (
	RegisterSrFieldFlevelShift = 8
	RegisterSrFieldFlevelMask  = 0x3f00
)

// GetFlevel FIFO level This field gives the number of valid bytes which are being held in the FIFO. FLEVEL = 0 when the FIFO is empty, and 16 when it is full. In memory-mapped mode and in automatic status polling mode, FLEVEL is zero.
func (r *RegisterSrType) GetFlevel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldFlevelMask) >> RegisterSrFieldFlevelShift)
}

// SetFlevel FIFO level This field gives the number of valid bytes which are being held in the FIFO. FLEVEL = 0 when the FIFO is empty, and 16 when it is full. In memory-mapped mode and in automatic status polling mode, FLEVEL is zero.
func (r *RegisterSrType) SetFlevel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldFlevelMask)|(uint32(value)<<RegisterSrFieldFlevelShift))
}

// RegisterFcrType QUADSPI flag clear register
type RegisterFcrType uint32

func (r *RegisterFcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterFcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterFcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterFcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterFcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterFcrFieldCtefShift = 0
	RegisterFcrFieldCtefMask  = 0x1
)

// GetCtef Clear transfer error flag Writing 1 clears the TEF flag in the QUADSPI_SR register
func (r *RegisterFcrType) GetCtef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFcrFieldCtefMask) != 0
}

// SetCtef Clear transfer error flag Writing 1 clears the TEF flag in the QUADSPI_SR register
func (r *RegisterFcrType) SetCtef(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFcrFieldCtefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFcrFieldCtefMask)
	}
}

const (
	RegisterFcrFieldCtcfShift = 1
	RegisterFcrFieldCtcfMask  = 0x2
)

// GetCtcf Clear transfer complete flag Writing 1 clears the TCF flag in the QUADSPI_SR register
func (r *RegisterFcrType) GetCtcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFcrFieldCtcfMask) != 0
}

// SetCtcf Clear transfer complete flag Writing 1 clears the TCF flag in the QUADSPI_SR register
func (r *RegisterFcrType) SetCtcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFcrFieldCtcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFcrFieldCtcfMask)
	}
}

const (
	RegisterFcrFieldCsmfShift = 3
	RegisterFcrFieldCsmfMask  = 0x8
)

// GetCsmf Clear status match flag Writing 1 clears the SMF flag in the QUADSPI_SR register
func (r *RegisterFcrType) GetCsmf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFcrFieldCsmfMask) != 0
}

// SetCsmf Clear status match flag Writing 1 clears the SMF flag in the QUADSPI_SR register
func (r *RegisterFcrType) SetCsmf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFcrFieldCsmfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFcrFieldCsmfMask)
	}
}

const (
	RegisterFcrFieldCtofShift = 4
	RegisterFcrFieldCtofMask  = 0x10
)

// GetCtof Clear timeout flag Writing 1 clears the TOF flag in the QUADSPI_SR register
func (r *RegisterFcrType) GetCtof() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterFcrFieldCtofMask) != 0
}

// SetCtof Clear timeout flag Writing 1 clears the TOF flag in the QUADSPI_SR register
func (r *RegisterFcrType) SetCtof(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterFcrFieldCtofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterFcrFieldCtofMask)
	}
}

// RegisterDlrType QUADSPI data length register
type RegisterDlrType uint32

func (r *RegisterDlrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDlrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDlrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDlrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDlrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDlrFieldDlShift = 0
	RegisterDlrFieldDlMask  = 0xffffffff
)

// GetDl Data length Number of data to be retrieved (value+1) in indirect and status-polling modes. A value no greater than 3 (indicating 4 bytes) should be used for status-polling mode. All 1s in indirect mode means undefined length, where QUADSPI will continue until the end of memory, as defined by FSIZE. 0x0000_0000: 1 byte is to be transferred 0x0000_0001: 2 bytes are to be transferred 0x0000_0002: 3 bytes are to be transferred 0x0000_0003: 4 bytes are to be transferred ... 0xFFFF_FFFD: 4,294,967,294 (4G-2) bytes are to be transferred 0xFFFF_FFFE: 4,294,967,295 (4G-1) bytes are to be transferred 0xFFFF_FFFF: undefined length -- all bytes until the end of Flash memory (as defined by FSIZE) are to be transferred. Continue reading indefinitely if FSIZE = 0x1F. DL[0] is stuck at 1 in dual-flash mode (DFM = 1) even when 0 is written to this bit, thus assuring that each access transfers an even number of bytes. This field has no effect when in memory-mapped mode (FMODE = 10). This field can be written only when BUSY = 0.
func (r *RegisterDlrType) GetDl() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDlrFieldDlMask) >> RegisterDlrFieldDlShift)
}

// SetDl Data length Number of data to be retrieved (value+1) in indirect and status-polling modes. A value no greater than 3 (indicating 4 bytes) should be used for status-polling mode. All 1s in indirect mode means undefined length, where QUADSPI will continue until the end of memory, as defined by FSIZE. 0x0000_0000: 1 byte is to be transferred 0x0000_0001: 2 bytes are to be transferred 0x0000_0002: 3 bytes are to be transferred 0x0000_0003: 4 bytes are to be transferred ... 0xFFFF_FFFD: 4,294,967,294 (4G-2) bytes are to be transferred 0xFFFF_FFFE: 4,294,967,295 (4G-1) bytes are to be transferred 0xFFFF_FFFF: undefined length -- all bytes until the end of Flash memory (as defined by FSIZE) are to be transferred. Continue reading indefinitely if FSIZE = 0x1F. DL[0] is stuck at 1 in dual-flash mode (DFM = 1) even when 0 is written to this bit, thus assuring that each access transfers an even number of bytes. This field has no effect when in memory-mapped mode (FMODE = 10). This field can be written only when BUSY = 0.
func (r *RegisterDlrType) SetDl(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDlrFieldDlMask)|(uint32(value)<<RegisterDlrFieldDlShift))
}

// RegisterCcrType QUADSPI communication configuration register
type RegisterCcrType uint32

func (r *RegisterCcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcrFieldInstructionShift = 0
	RegisterCcrFieldInstructionMask  = 0xff
)

// GetInstruction Instruction Instruction to be send to the external SPI device. This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetInstruction() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldInstructionMask) >> RegisterCcrFieldInstructionShift)
}

// SetInstruction Instruction Instruction to be send to the external SPI device. This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetInstruction(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldInstructionMask)|(uint32(value)<<RegisterCcrFieldInstructionShift))
}

const (
	RegisterCcrFieldImodeShift = 8
	RegisterCcrFieldImodeMask  = 0x300
)

// GetImode Instruction mode This field defines the instruction phase mode of operation: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetImode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldImodeMask) >> RegisterCcrFieldImodeShift)
}

// SetImode Instruction mode This field defines the instruction phase mode of operation: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetImode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldImodeMask)|(uint32(value)<<RegisterCcrFieldImodeShift))
}

const (
	RegisterCcrFieldAdmodeShift = 10
	RegisterCcrFieldAdmodeMask  = 0xc00
)

// GetAdmode Address mode This field defines the address phase mode of operation: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetAdmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldAdmodeMask) >> RegisterCcrFieldAdmodeShift)
}

// SetAdmode Address mode This field defines the address phase mode of operation: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetAdmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldAdmodeMask)|(uint32(value)<<RegisterCcrFieldAdmodeShift))
}

const (
	RegisterCcrFieldAdsizeShift = 12
	RegisterCcrFieldAdsizeMask  = 0x3000
)

// GetAdsize Address size This bit defines address size: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetAdsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldAdsizeMask) >> RegisterCcrFieldAdsizeShift)
}

// SetAdsize Address size This bit defines address size: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetAdsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldAdsizeMask)|(uint32(value)<<RegisterCcrFieldAdsizeShift))
}

const (
	RegisterCcrFieldAbmodeShift = 14
	RegisterCcrFieldAbmodeMask  = 0xc000
)

// GetAbmode Alternate bytes mode This field defines the alternate-bytes phase mode of operation: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetAbmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldAbmodeMask) >> RegisterCcrFieldAbmodeShift)
}

// SetAbmode Alternate bytes mode This field defines the alternate-bytes phase mode of operation: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetAbmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldAbmodeMask)|(uint32(value)<<RegisterCcrFieldAbmodeShift))
}

const (
	RegisterCcrFieldAbsizeShift = 16
	RegisterCcrFieldAbsizeMask  = 0x30000
)

// GetAbsize Alternate bytes size This bit defines alternate bytes size: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetAbsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldAbsizeMask) >> RegisterCcrFieldAbsizeShift)
}

// SetAbsize Alternate bytes size This bit defines alternate bytes size: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetAbsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldAbsizeMask)|(uint32(value)<<RegisterCcrFieldAbsizeShift))
}

const (
	RegisterCcrFieldDcycShift = 18
	RegisterCcrFieldDcycMask  = 0x7c0000
)

// GetDcyc Number of dummy cycles This field defines the duration of the dummy phase. In both SDR and DDR modes, it specifies a number of CLK cycles (0-31). This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetDcyc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDcycMask) >> RegisterCcrFieldDcycShift)
}

// SetDcyc Number of dummy cycles This field defines the duration of the dummy phase. In both SDR and DDR modes, it specifies a number of CLK cycles (0-31). This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetDcyc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDcycMask)|(uint32(value)<<RegisterCcrFieldDcycShift))
}

const (
	RegisterCcrFieldDmodeShift = 24
	RegisterCcrFieldDmodeMask  = 0x3000000
)

// GetDmode Data mode This field defines the data phases mode of operation: This field also determines the dummy phase mode of operation. This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetDmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDmodeMask) >> RegisterCcrFieldDmodeShift)
}

// SetDmode Data mode This field defines the data phases mode of operation: This field also determines the dummy phase mode of operation. This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetDmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDmodeMask)|(uint32(value)<<RegisterCcrFieldDmodeShift))
}

const (
	RegisterCcrFieldFmodeShift = 26
	RegisterCcrFieldFmodeMask  = 0xc000000
)

// GetFmode Functional mode This field defines the QUADSPI functional mode of operation. If DMAEN = 1 already, then the DMA controller for the corresponding channel must be disabled before changing the FMODE value. This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetFmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldFmodeMask) >> RegisterCcrFieldFmodeShift)
}

// SetFmode Functional mode This field defines the QUADSPI functional mode of operation. If DMAEN = 1 already, then the DMA controller for the corresponding channel must be disabled before changing the FMODE value. This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetFmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldFmodeMask)|(uint32(value)<<RegisterCcrFieldFmodeShift))
}

const (
	RegisterCcrFieldSiooShift = 28
	RegisterCcrFieldSiooMask  = 0x10000000
)

// GetSioo Send instruction only once mode See Section15.3.11: Sending the instruction only once on page13. This bit has no effect when IMODE = 00. This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetSioo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldSiooMask) != 0
}

// SetSioo Send instruction only once mode See Section15.3.11: Sending the instruction only once on page13. This bit has no effect when IMODE = 00. This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetSioo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldSiooMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldSiooMask)
	}
}

const (
	RegisterCcrFieldDhhcShift = 30
	RegisterCcrFieldDhhcMask  = 0x40000000
)

// GetDhhc DDR hold Delay the data output by 1/4 of the QUADSPI output clock cycle in DDR mode: This feature is only active in DDR mode. This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetDhhc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDhhcMask) != 0
}

// SetDhhc DDR hold Delay the data output by 1/4 of the QUADSPI output clock cycle in DDR mode: This feature is only active in DDR mode. This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetDhhc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldDhhcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDhhcMask)
	}
}

const (
	RegisterCcrFieldDdrmShift = 31
	RegisterCcrFieldDdrmMask  = 0x80000000
)

// GetDdrm Double data rate mode This bit sets the DDR mode for the address, alternate byte and data phase: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) GetDdrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDdrmMask) != 0
}

// SetDdrm Double data rate mode This bit sets the DDR mode for the address, alternate byte and data phase: This field can be written only when BUSY = 0.
func (r *RegisterCcrType) SetDdrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldDdrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDdrmMask)
	}
}

// RegisterArType QUADSPI address register
type RegisterArType uint32

func (r *RegisterArType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterArType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterArType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterArType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterArType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterArFieldAddressShift = 0
	RegisterArFieldAddressMask  = 0xffffffff
)

// GetAddress [31 0]: Address Address to be send to the external Flash memory Writes to this field are ignored when BUSY = 0 or when FMODE = 11 (memory-mapped mode). In dual flash mode, ADDRESS[0] is automatically stuck to 0 as the address should always be even
func (r *RegisterArType) GetAddress() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterArFieldAddressMask) >> RegisterArFieldAddressShift)
}

// SetAddress [31 0]: Address Address to be send to the external Flash memory Writes to this field are ignored when BUSY = 0 or when FMODE = 11 (memory-mapped mode). In dual flash mode, ADDRESS[0] is automatically stuck to 0 as the address should always be even
func (r *RegisterArType) SetAddress(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArFieldAddressMask)|(uint32(value)<<RegisterArFieldAddressShift))
}

// RegisterAbrType QUADSPI alternate bytes registers
type RegisterAbrType uint32

func (r *RegisterAbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAbrFieldAlternateShift = 0
	RegisterAbrFieldAlternateMask  = 0xffffffff
)

// GetAlternate Alternate Bytes Optional data to be send to the external SPI device right after the address. This field can be written only when BUSY = 0.
func (r *RegisterAbrType) GetAlternate() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterAbrFieldAlternateMask) >> RegisterAbrFieldAlternateShift)
}

// SetAlternate Alternate Bytes Optional data to be send to the external SPI device right after the address. This field can be written only when BUSY = 0.
func (r *RegisterAbrType) SetAlternate(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAbrFieldAlternateMask)|(uint32(value)<<RegisterAbrFieldAlternateShift))
}

// RegisterDrType QUADSPI data register
type RegisterDrType uint32

func (r *RegisterDrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDrFieldDataShift = 0
	RegisterDrFieldDataMask  = 0xffffffff
)

// GetData Data Data to be sent/received to/from the external SPI device. In indirect write mode, data written to this register is stored on the FIFO before it is sent to the Flash memory during the data phase. If the FIFO is too full, a write operation is stalled until the FIFO has enough space to accept the amount of data being written. In indirect read mode, reading this register gives (via the FIFO) the data which was received from the Flash memory. If the FIFO does not have as many bytes as requested by the read operation and if BUSY=1, the read operation is stalled until enough data is present or until the transfer is complete, whichever happens first. In automatic polling mode, this register contains the last data read from the Flash memory (without masking). Word, halfword, and byte accesses to this register are supported. In indirect write mode, a byte write adds 1 byte to the FIFO, a halfword write 2, and a word write 4. Similarly, in indirect read mode, a byte read removes 1 byte from the FIFO, a halfword read 2, and a word read 4. Accesses in indirect mode must be aligned to the bottom of this register: a byte read must read DATA[7:0] and a halfword read must read DATA[15:0].
func (r *RegisterDrType) GetData() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDrFieldDataMask) >> RegisterDrFieldDataShift)
}

// SetData Data Data to be sent/received to/from the external SPI device. In indirect write mode, data written to this register is stored on the FIFO before it is sent to the Flash memory during the data phase. If the FIFO is too full, a write operation is stalled until the FIFO has enough space to accept the amount of data being written. In indirect read mode, reading this register gives (via the FIFO) the data which was received from the Flash memory. If the FIFO does not have as many bytes as requested by the read operation and if BUSY=1, the read operation is stalled until enough data is present or until the transfer is complete, whichever happens first. In automatic polling mode, this register contains the last data read from the Flash memory (without masking). Word, halfword, and byte accesses to this register are supported. In indirect write mode, a byte write adds 1 byte to the FIFO, a halfword write 2, and a word write 4. Similarly, in indirect read mode, a byte read removes 1 byte from the FIFO, a halfword read 2, and a word read 4. Accesses in indirect mode must be aligned to the bottom of this register: a byte read must read DATA[7:0] and a halfword read must read DATA[15:0].
func (r *RegisterDrType) SetData(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDrFieldDataMask)|(uint32(value)<<RegisterDrFieldDataShift))
}

// RegisterPsmkrType QUADSPI polling status mask register
type RegisterPsmkrType uint32

func (r *RegisterPsmkrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPsmkrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPsmkrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPsmkrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPsmkrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPsmkrFieldMaskShift = 0
	RegisterPsmkrFieldMaskMask  = 0xffffffff
)

// GetMask Status mask Mask to be applied to the status bytes received in polling mode. For bit n: This field can be written only when BUSY = 0.
func (r *RegisterPsmkrType) GetMask() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterPsmkrFieldMaskMask) >> RegisterPsmkrFieldMaskShift)
}

// SetMask Status mask Mask to be applied to the status bytes received in polling mode. For bit n: This field can be written only when BUSY = 0.
func (r *RegisterPsmkrType) SetMask(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPsmkrFieldMaskMask)|(uint32(value)<<RegisterPsmkrFieldMaskShift))
}

// RegisterPsmarType QUADSPI polling status match register
type RegisterPsmarType uint32

func (r *RegisterPsmarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPsmarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPsmarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPsmarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPsmarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPsmarFieldMatchShift = 0
	RegisterPsmarFieldMatchMask  = 0xffffffff
)

// GetMatch Status match Value to be compared with the masked status register to get a match. This field can be written only when BUSY = 0.
func (r *RegisterPsmarType) GetMatch() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterPsmarFieldMatchMask) >> RegisterPsmarFieldMatchShift)
}

// SetMatch Status match Value to be compared with the masked status register to get a match. This field can be written only when BUSY = 0.
func (r *RegisterPsmarType) SetMatch(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPsmarFieldMatchMask)|(uint32(value)<<RegisterPsmarFieldMatchShift))
}

// RegisterPirType QUADSPI polling interval register
type RegisterPirType uint32

func (r *RegisterPirType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPirType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPirType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPirType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPirType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPirFieldIntervalShift = 0
	RegisterPirFieldIntervalMask  = 0xffff
)

// GetInterval Polling interval Number of CLK cycles between to read during automatic polling phases. This field can be written only when BUSY = 0.
func (r *RegisterPirType) GetInterval() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPirFieldIntervalMask) >> RegisterPirFieldIntervalShift)
}

// SetInterval Polling interval Number of CLK cycles between to read during automatic polling phases. This field can be written only when BUSY = 0.
func (r *RegisterPirType) SetInterval(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPirFieldIntervalMask)|(uint32(value)<<RegisterPirFieldIntervalShift))
}

// RegisterLptrType QUADSPI low-power timeout register
type RegisterLptrType uint32

func (r *RegisterLptrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterLptrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterLptrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterLptrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterLptrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterLptrFieldTimeoutShift = 0
	RegisterLptrFieldTimeoutMask  = 0xffff
)

// GetTimeout Timeout period After each access in memory-mapped mode, the QUADSPI prefetches the subsequent bytes and holds these bytes in the FIFO. This field indicates how many CLK cycles the QUADSPI waits after the FIFO becomes full until it raises nCS, putting the Flash memory in a lower-consumption state. This field can be written only when BUSY = 0.
func (r *RegisterLptrType) GetTimeout() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterLptrFieldTimeoutMask) >> RegisterLptrFieldTimeoutShift)
}

// SetTimeout Timeout period After each access in memory-mapped mode, the QUADSPI prefetches the subsequent bytes and holds these bytes in the FIFO. This field indicates how many CLK cycles the QUADSPI waits after the FIFO becomes full until it raises nCS, putting the Flash memory in a lower-consumption state. This field can be written only when BUSY = 0.
func (r *RegisterLptrType) SetTimeout(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterLptrFieldTimeoutMask)|(uint32(value)<<RegisterLptrFieldTimeoutShift))
}
