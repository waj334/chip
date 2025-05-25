//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package fmc

import (
	"unsafe"
	"volatile"
)

var (
	Fmc = (*_fmc)(unsafe.Pointer(uintptr(0x52004000)))
)

type _fmc struct {
	Bcr1  registerBcr1Type
	Btr1  registerBtr1Type
	Bcr2  registerBcr2Type
	Btr2  registerBtr2Type
	Bcr3  registerBcr3Type
	Btr3  registerBtr3Type
	Bcr4  registerBcr4Type
	Btr4  registerBtr4Type
	_     [96]byte
	Pcr   registerPcrType
	Sr    registerSrType
	Pmem  registerPmemType
	Patt  registerPattType
	_     [4]byte
	Eccr  registerEccrType
	_     [108]byte
	Bwtr1 registerBwtr1Type
	_     [4]byte
	Bwtr2 registerBwtr2Type
	_     [4]byte
	Bwtr3 registerBwtr3Type
	_     [4]byte
	Bwtr4 registerBwtr4Type
	_     [32]byte
	Sdcr1 registerSdcr1Type
	Sdcr2 registerSdcr2Type
	Sdtr1 registerSdtr1Type
	Sdtr2 registerSdtr2Type
	Sdcmr registerSdcmrType
	Sdrtr registerSdrtrType
	Sdsr  registerSdsrType
}

// registerBcr1Type This register contains the control information of each memory bank, used for SRAMs, PSRAM and NOR Flash memories.
type registerBcr1Type uint32

const (
	RegisterBcr1FieldMbkenShift = 0
	RegisterBcr1FieldMbkenMask  = 0x1
)

// GetMbken Memory bank enable bit This bit enables the memory bank. After reset Bank1 is enabled, all others are disabled. Accessing a disabled bank causes an ERROR on AXI bus.
func (r *registerBcr1Type) GetMbken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldMbkenMask) != 0
}

// SetMbken Memory bank enable bit This bit enables the memory bank. After reset Bank1 is enabled, all others are disabled. Accessing a disabled bank causes an ERROR on AXI bus.
func (r *registerBcr1Type) SetMbken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldMbkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldMbkenMask)
	}
}

const (
	RegisterBcr1FieldMuxenShift = 1
	RegisterBcr1FieldMuxenMask  = 0x2
)

// GetMuxen Address/data multiplexing enable bit When this bit is set, the address and data values are multiplexed on the data bus, valid only with NOR and PSRAM memories:
func (r *registerBcr1Type) GetMuxen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldMuxenMask) != 0
}

// SetMuxen Address/data multiplexing enable bit When this bit is set, the address and data values are multiplexed on the data bus, valid only with NOR and PSRAM memories:
func (r *registerBcr1Type) SetMuxen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldMuxenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldMuxenMask)
	}
}

const (
	RegisterBcr1FieldMtypShift = 2
	RegisterBcr1FieldMtypMask  = 0xc
)

// GetMtyp Memory type These bits define the type of external memory attached to the corresponding memory bank:
func (r *registerBcr1Type) GetMtyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldMtypMask) >> RegisterBcr1FieldMtypShift)
}

// SetMtyp Memory type These bits define the type of external memory attached to the corresponding memory bank:
func (r *registerBcr1Type) SetMtyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldMtypMask)|(uint32(value)<<RegisterBcr1FieldMtypShift))
}

const (
	RegisterBcr1FieldMwidShift = 4
	RegisterBcr1FieldMwidMask  = 0x30
)

// GetMwid Memory data bus width Defines the external memory device width, valid for all type of memories.
func (r *registerBcr1Type) GetMwid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldMwidMask) >> RegisterBcr1FieldMwidShift)
}

// SetMwid Memory data bus width Defines the external memory device width, valid for all type of memories.
func (r *registerBcr1Type) SetMwid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldMwidMask)|(uint32(value)<<RegisterBcr1FieldMwidShift))
}

const (
	RegisterBcr1FieldFaccenShift = 6
	RegisterBcr1FieldFaccenMask  = 0x40
)

// GetFaccen Flash access enable This bit enables NOR Flash memory access operations.
func (r *registerBcr1Type) GetFaccen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldFaccenMask) != 0
}

// SetFaccen Flash access enable This bit enables NOR Flash memory access operations.
func (r *registerBcr1Type) SetFaccen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldFaccenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldFaccenMask)
	}
}

const (
	RegisterBcr1FieldBurstenShift = 8
	RegisterBcr1FieldBurstenMask  = 0x100
)

// GetBursten Burst enable bit This bit enables/disables synchronous accesses during read operations. It is valid only for synchronous memories operating in Burst mode:
func (r *registerBcr1Type) GetBursten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldBurstenMask) != 0
}

// SetBursten Burst enable bit This bit enables/disables synchronous accesses during read operations. It is valid only for synchronous memories operating in Burst mode:
func (r *registerBcr1Type) SetBursten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldBurstenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldBurstenMask)
	}
}

const (
	RegisterBcr1FieldWaitpolShift = 9
	RegisterBcr1FieldWaitpolMask  = 0x200
)

// GetWaitpol Wait signal polarity bit This bit defines the polarity of the wait signal from memory used for either in synchronous or asynchronous mode:
func (r *registerBcr1Type) GetWaitpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldWaitpolMask) != 0
}

// SetWaitpol Wait signal polarity bit This bit defines the polarity of the wait signal from memory used for either in synchronous or asynchronous mode:
func (r *registerBcr1Type) SetWaitpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldWaitpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldWaitpolMask)
	}
}

const (
	RegisterBcr1FieldWaitcfgShift = 11
	RegisterBcr1FieldWaitcfgMask  = 0x800
)

// GetWaitcfg Wait timing configuration The NWAIT signal indicates whether the data from the memory are valid or if a wait state must be inserted when accessing the memory in synchronous mode. This configuration bit determines if NWAIT is asserted by the memory one clock cycle before the wait state or during the wait state:
func (r *registerBcr1Type) GetWaitcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldWaitcfgMask) != 0
}

// SetWaitcfg Wait timing configuration The NWAIT signal indicates whether the data from the memory are valid or if a wait state must be inserted when accessing the memory in synchronous mode. This configuration bit determines if NWAIT is asserted by the memory one clock cycle before the wait state or during the wait state:
func (r *registerBcr1Type) SetWaitcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldWaitcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldWaitcfgMask)
	}
}

const (
	RegisterBcr1FieldWrenShift = 12
	RegisterBcr1FieldWrenMask  = 0x1000
)

// GetWren Write enable bit This bit indicates whether write operations are enabled/disabled in the bank by the FMC:
func (r *registerBcr1Type) GetWren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldWrenMask) != 0
}

// SetWren Write enable bit This bit indicates whether write operations are enabled/disabled in the bank by the FMC:
func (r *registerBcr1Type) SetWren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldWrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldWrenMask)
	}
}

const (
	RegisterBcr1FieldWaitenShift = 13
	RegisterBcr1FieldWaitenMask  = 0x2000
)

// GetWaiten Wait enable bit This bit enables/disables wait-state insertion via the NWAIT signal when accessing the memory in synchronous mode.
func (r *registerBcr1Type) GetWaiten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldWaitenMask) != 0
}

// SetWaiten Wait enable bit This bit enables/disables wait-state insertion via the NWAIT signal when accessing the memory in synchronous mode.
func (r *registerBcr1Type) SetWaiten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldWaitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldWaitenMask)
	}
}

const (
	RegisterBcr1FieldExtmodShift = 14
	RegisterBcr1FieldExtmodMask  = 0x4000
)

// GetExtmod Extended mode enable. This bit enables the FMC to program the write timings for asynchronous accesses inside the FMC_BWTR register, thus resulting in different timings for read and write operations. Note: When the extended mode is disabled, the FMC can operate in Mode1 or Mode2 as follows: ** Mode 1 is the default mode when the SRAM/PSRAM memory type is selected (MTYP =0x0 or 0x01) ** Mode 2 is the default mode when the NOR memory type is selected (MTYP = 0x10).
func (r *registerBcr1Type) GetExtmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldExtmodMask) != 0
}

// SetExtmod Extended mode enable. This bit enables the FMC to program the write timings for asynchronous accesses inside the FMC_BWTR register, thus resulting in different timings for read and write operations. Note: When the extended mode is disabled, the FMC can operate in Mode1 or Mode2 as follows: ** Mode 1 is the default mode when the SRAM/PSRAM memory type is selected (MTYP =0x0 or 0x01) ** Mode 2 is the default mode when the NOR memory type is selected (MTYP = 0x10).
func (r *registerBcr1Type) SetExtmod(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldExtmodMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldExtmodMask)
	}
}

const (
	RegisterBcr1FieldAsyncwaitShift = 15
	RegisterBcr1FieldAsyncwaitMask  = 0x8000
)

// GetAsyncwait Wait signal during asynchronous transfers This bit enables/disables the FMC to use the wait signal even during an asynchronous protocol.
func (r *registerBcr1Type) GetAsyncwait() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldAsyncwaitMask) != 0
}

// SetAsyncwait Wait signal during asynchronous transfers This bit enables/disables the FMC to use the wait signal even during an asynchronous protocol.
func (r *registerBcr1Type) SetAsyncwait(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldAsyncwaitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldAsyncwaitMask)
	}
}

const (
	RegisterBcr1FieldCpsizeShift = 16
	RegisterBcr1FieldCpsizeMask  = 0x70000
)

// GetCpsize CRAM Page Size These are used for Cellular RAM 1.5 which does not allow burst access to cross the address boundaries between pages. When these bits are configured, the FMC controller splits automatically the burst access when the memory page size is reached (refer to memory datasheet for page size). Other configuration: reserved.
func (r *registerBcr1Type) GetCpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldCpsizeMask) >> RegisterBcr1FieldCpsizeShift)
}

// SetCpsize CRAM Page Size These are used for Cellular RAM 1.5 which does not allow burst access to cross the address boundaries between pages. When these bits are configured, the FMC controller splits automatically the burst access when the memory page size is reached (refer to memory datasheet for page size). Other configuration: reserved.
func (r *registerBcr1Type) SetCpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldCpsizeMask)|(uint32(value)<<RegisterBcr1FieldCpsizeShift))
}

const (
	RegisterBcr1FieldCburstrwShift = 19
	RegisterBcr1FieldCburstrwMask  = 0x80000
)

// GetCburstrw Write burst enable For PSRAM (CRAM) operating in Burst mode, the bit enables synchronous accesses during write operations. The enable bit for synchronous read accesses is the BURSTEN bit in the FMC_BCRx register.
func (r *registerBcr1Type) GetCburstrw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldCburstrwMask) != 0
}

// SetCburstrw Write burst enable For PSRAM (CRAM) operating in Burst mode, the bit enables synchronous accesses during write operations. The enable bit for synchronous read accesses is the BURSTEN bit in the FMC_BCRx register.
func (r *registerBcr1Type) SetCburstrw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldCburstrwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldCburstrwMask)
	}
}

const (
	RegisterBcr1FieldCclkenShift = 20
	RegisterBcr1FieldCclkenMask  = 0x100000
)

// GetCclken Continuous Clock Enable This bit enables the FMC_CLK clock output to external memory devices. Note: The CCLKEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register. Bank 1 must be configured in synchronous mode to generate the FMC_CLK continuous clock. If CCLKEN bit is set, the FMC_CLK clock ratio is specified by CLKDIV value in the FMC_BTR1 register. CLKDIV in FMC_BWTR1 is dont care. If the synchronous mode is used and CCLKEN bit is set, the synchronous memories connected to other banks than Bank 1 are clocked by the same clock (the CLKDIV value in the FMC_BTR2..4 and FMC_BWTR2..4 registers for other banks has no effect.)
func (r *registerBcr1Type) GetCclken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldCclkenMask) != 0
}

// SetCclken Continuous Clock Enable This bit enables the FMC_CLK clock output to external memory devices. Note: The CCLKEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register. Bank 1 must be configured in synchronous mode to generate the FMC_CLK continuous clock. If CCLKEN bit is set, the FMC_CLK clock ratio is specified by CLKDIV value in the FMC_BTR1 register. CLKDIV in FMC_BWTR1 is dont care. If the synchronous mode is used and CCLKEN bit is set, the synchronous memories connected to other banks than Bank 1 are clocked by the same clock (the CLKDIV value in the FMC_BTR2..4 and FMC_BWTR2..4 registers for other banks has no effect.)
func (r *registerBcr1Type) SetCclken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldCclkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldCclkenMask)
	}
}

const (
	RegisterBcr1FieldWfdisShift = 21
	RegisterBcr1FieldWfdisMask  = 0x200000
)

// GetWfdis Write FIFO Disable This bit disables the Write FIFO used by the FMC controller. Note: The WFDIS bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr1Type) GetWfdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldWfdisMask) != 0
}

// SetWfdis Write FIFO Disable This bit disables the Write FIFO used by the FMC controller. Note: The WFDIS bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr1Type) SetWfdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldWfdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldWfdisMask)
	}
}

const (
	RegisterBcr1FieldBmapShift = 24
	RegisterBcr1FieldBmapMask  = 0x3000000
)

// GetBmap FMC bank mapping These bits allows different to remap SDRAM bank2 or swap the FMC NOR/PSRAM and SDRAM banks.Refer to Table 10 for Note: The BMAP bits of the FMC_BCR2..4 registers are dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr1Type) GetBmap() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldBmapMask) >> RegisterBcr1FieldBmapShift)
}

// SetBmap FMC bank mapping These bits allows different to remap SDRAM bank2 or swap the FMC NOR/PSRAM and SDRAM banks.Refer to Table 10 for Note: The BMAP bits of the FMC_BCR2..4 registers are dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr1Type) SetBmap(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldBmapMask)|(uint32(value)<<RegisterBcr1FieldBmapShift))
}

const (
	RegisterBcr1FieldFmcenShift = 31
	RegisterBcr1FieldFmcenMask  = 0x80000000
)

// GetFmcen FMC controller Enable This bit enables/disables the FMC controller. Note: The FMCEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr1Type) GetFmcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr1FieldFmcenMask) != 0
}

// SetFmcen FMC controller Enable This bit enables/disables the FMC controller. Note: The FMCEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr1Type) SetFmcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr1FieldFmcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr1FieldFmcenMask)
	}
}

// registerBtr1Type This register contains the control information of each memory bank, used for SRAMs, PSRAM and NOR Flash memories.If the EXTMOD bit is set in the FMC_BCRx register, then this register is partitioned for write and read access, that is, 2 registers are available: one to configure read accesses (this register) and one to configure write accesses (FMC_BWTRx registers).
type registerBtr1Type uint32

const (
	RegisterBtr1FieldAddsetShift = 0
	RegisterBtr1FieldAddsetMask  = 0xf
)

// GetAddset Address setup phase duration These bits are written by software to define the duration of the address setup phase (refer to Figure81 to Figure93), used in SRAMs, ROMs and asynchronous NOR Flash: For each access mode address setup phase duration, please refer to the respective figure (refer to Figure81 to Figure93). Note: In synchronous accesses, this value is dont care. In Muxed mode or Mode D, the minimum value for ADDSET is 1.
func (r *registerBtr1Type) GetAddset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr1FieldAddsetMask) >> RegisterBtr1FieldAddsetShift)
}

// SetAddset Address setup phase duration These bits are written by software to define the duration of the address setup phase (refer to Figure81 to Figure93), used in SRAMs, ROMs and asynchronous NOR Flash: For each access mode address setup phase duration, please refer to the respective figure (refer to Figure81 to Figure93). Note: In synchronous accesses, this value is dont care. In Muxed mode or Mode D, the minimum value for ADDSET is 1.
func (r *registerBtr1Type) SetAddset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr1FieldAddsetMask)|(uint32(value)<<RegisterBtr1FieldAddsetShift))
}

const (
	RegisterBtr1FieldAddhldShift = 4
	RegisterBtr1FieldAddhldMask  = 0xf0
)

// GetAddhld Address-hold phase duration These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in mode D or multiplexed accesses: For each access mode address-hold phase duration, please refer to the respective figure (Figure81 to Figure93). Note: In synchronous accesses, this value is not used, the address hold phase is always 1 memory clock period duration.
func (r *registerBtr1Type) GetAddhld() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr1FieldAddhldMask) >> RegisterBtr1FieldAddhldShift)
}

// SetAddhld Address-hold phase duration These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in mode D or multiplexed accesses: For each access mode address-hold phase duration, please refer to the respective figure (Figure81 to Figure93). Note: In synchronous accesses, this value is not used, the address hold phase is always 1 memory clock period duration.
func (r *registerBtr1Type) SetAddhld(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr1FieldAddhldMask)|(uint32(value)<<RegisterBtr1FieldAddhldShift))
}

const (
	RegisterBtr1FieldDatastShift = 8
	RegisterBtr1FieldDatastMask  = 0xff00
)

// GetDatast Data-phase duration These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous accesses: For each memory type and access mode data-phase duration, please refer to the respective figure (Figure81 to Figure93). Example: Mode1, write access, DATAST=1: Data-phase duration= DATAST+1 = 2 KCK_FMC clock cycles. Note: In synchronous accesses, this value is dont care.
func (r *registerBtr1Type) GetDatast() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr1FieldDatastMask) >> RegisterBtr1FieldDatastShift)
}

// SetDatast Data-phase duration These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous accesses: For each memory type and access mode data-phase duration, please refer to the respective figure (Figure81 to Figure93). Example: Mode1, write access, DATAST=1: Data-phase duration= DATAST+1 = 2 KCK_FMC clock cycles. Note: In synchronous accesses, this value is dont care.
func (r *registerBtr1Type) SetDatast(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr1FieldDatastMask)|(uint32(value)<<RegisterBtr1FieldDatastShift))
}

const (
	RegisterBtr1FieldBusturnShift = 16
	RegisterBtr1FieldBusturnMask  = 0xf0000
)

// GetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write-to-read or read-to write transaction. The programmed bus turnaround delay is inserted between an asynchronous read (in muxed or mode D) or write transaction and any other asynchronous /synchronous read/write from/to a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except in muxed mode and mode D. There is a bus turnaround delay of 1 FMC clock cycle between: Two consecutive asynchronous read transfers to the same static memory bank except for modes muxed and D. An asynchronous read to an asynchronous or synchronous write to any static bank or dynamic bank except in modes muxed and D mode. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank. A synchronous write (burst or single) access and an asynchronous write or read transfer to or from static memory bank (the bank can be the same or a different one in case of a read operation. Two consecutive synchronous read operations (in Burst or Single mode) followed by any synchronous/asynchronous read or write from/to another static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write access (in Burst or Single mode) and a synchronous read from the same or a different bank. The bus turnaround delay allows to match the minimum time between consecutive transactions (tEHEL from NEx high to NEx low) and the maximum time required by the memory to free the data bus after a read access (tEHQZ): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin and (BUSTRUN + 2)KCK_FMC period &#8805; tEHQZmax if EXTMOD = 0 (BUSTRUN + 2)KCK_FMC period &#8805; max (tEHELmin, tEHQZmax) if EXTMOD = 126. ...
func (r *registerBtr1Type) GetBusturn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr1FieldBusturnMask) >> RegisterBtr1FieldBusturnShift)
}

// SetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write-to-read or read-to write transaction. The programmed bus turnaround delay is inserted between an asynchronous read (in muxed or mode D) or write transaction and any other asynchronous /synchronous read/write from/to a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except in muxed mode and mode D. There is a bus turnaround delay of 1 FMC clock cycle between: Two consecutive asynchronous read transfers to the same static memory bank except for modes muxed and D. An asynchronous read to an asynchronous or synchronous write to any static bank or dynamic bank except in modes muxed and D mode. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank. A synchronous write (burst or single) access and an asynchronous write or read transfer to or from static memory bank (the bank can be the same or a different one in case of a read operation. Two consecutive synchronous read operations (in Burst or Single mode) followed by any synchronous/asynchronous read or write from/to another static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write access (in Burst or Single mode) and a synchronous read from the same or a different bank. The bus turnaround delay allows to match the minimum time between consecutive transactions (tEHEL from NEx high to NEx low) and the maximum time required by the memory to free the data bus after a read access (tEHQZ): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin and (BUSTRUN + 2)KCK_FMC period &#8805; tEHQZmax if EXTMOD = 0 (BUSTRUN + 2)KCK_FMC period &#8805; max (tEHELmin, tEHQZmax) if EXTMOD = 126. ...
func (r *registerBtr1Type) SetBusturn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr1FieldBusturnMask)|(uint32(value)<<RegisterBtr1FieldBusturnShift))
}

const (
	RegisterBtr1FieldClkdivShift = 20
	RegisterBtr1FieldClkdivMask  = 0xf00000
)

// GetClkdiv Clock divide ratio (for FMC_CLK signal) These bits define the period of FMC_CLK clock output signal, expressed in number of KCK_FMC cycles: In asynchronous NOR Flash, SRAM or PSRAM accesses, this value is dont care. Note: Refer to Section20.6.5: Synchronous transactions for FMC_CLK divider ratio formula)
func (r *registerBtr1Type) GetClkdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr1FieldClkdivMask) >> RegisterBtr1FieldClkdivShift)
}

// SetClkdiv Clock divide ratio (for FMC_CLK signal) These bits define the period of FMC_CLK clock output signal, expressed in number of KCK_FMC cycles: In asynchronous NOR Flash, SRAM or PSRAM accesses, this value is dont care. Note: Refer to Section20.6.5: Synchronous transactions for FMC_CLK divider ratio formula)
func (r *registerBtr1Type) SetClkdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr1FieldClkdivMask)|(uint32(value)<<RegisterBtr1FieldClkdivShift))
}

const (
	RegisterBtr1FieldDatlatShift = 24
	RegisterBtr1FieldDatlatMask  = 0xf000000
)

// GetDatlat Data latency for synchronous memory For synchronous access with read write burst mode enabled these bits define the number of memory clock cycles
func (r *registerBtr1Type) GetDatlat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr1FieldDatlatMask) >> RegisterBtr1FieldDatlatShift)
}

// SetDatlat Data latency for synchronous memory For synchronous access with read write burst mode enabled these bits define the number of memory clock cycles
func (r *registerBtr1Type) SetDatlat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr1FieldDatlatMask)|(uint32(value)<<RegisterBtr1FieldDatlatShift))
}

const (
	RegisterBtr1FieldAccmodShift = 28
	RegisterBtr1FieldAccmodMask  = 0x30000000
)

// GetAccmod Access mode These bits specify the asynchronous access modes as shown in the timing diagrams. They are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBtr1Type) GetAccmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr1FieldAccmodMask) >> RegisterBtr1FieldAccmodShift)
}

// SetAccmod Access mode These bits specify the asynchronous access modes as shown in the timing diagrams. They are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBtr1Type) SetAccmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr1FieldAccmodMask)|(uint32(value)<<RegisterBtr1FieldAccmodShift))
}

// registerBcr2Type This register contains the control information of each memory bank, used for SRAMs, PSRAM and NOR Flash memories.
type registerBcr2Type uint32

const (
	RegisterBcr2FieldMbkenShift = 0
	RegisterBcr2FieldMbkenMask  = 0x1
)

// GetMbken Memory bank enable bit This bit enables the memory bank. After reset Bank1 is enabled, all others are disabled. Accessing a disabled bank causes an ERROR on AXI bus.
func (r *registerBcr2Type) GetMbken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldMbkenMask) != 0
}

// SetMbken Memory bank enable bit This bit enables the memory bank. After reset Bank1 is enabled, all others are disabled. Accessing a disabled bank causes an ERROR on AXI bus.
func (r *registerBcr2Type) SetMbken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldMbkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldMbkenMask)
	}
}

const (
	RegisterBcr2FieldMuxenShift = 1
	RegisterBcr2FieldMuxenMask  = 0x2
)

// GetMuxen Address/data multiplexing enable bit When this bit is set, the address and data values are multiplexed on the data bus, valid only with NOR and PSRAM memories:
func (r *registerBcr2Type) GetMuxen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldMuxenMask) != 0
}

// SetMuxen Address/data multiplexing enable bit When this bit is set, the address and data values are multiplexed on the data bus, valid only with NOR and PSRAM memories:
func (r *registerBcr2Type) SetMuxen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldMuxenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldMuxenMask)
	}
}

const (
	RegisterBcr2FieldMtypShift = 2
	RegisterBcr2FieldMtypMask  = 0xc
)

// GetMtyp Memory type These bits define the type of external memory attached to the corresponding memory bank:
func (r *registerBcr2Type) GetMtyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldMtypMask) >> RegisterBcr2FieldMtypShift)
}

// SetMtyp Memory type These bits define the type of external memory attached to the corresponding memory bank:
func (r *registerBcr2Type) SetMtyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldMtypMask)|(uint32(value)<<RegisterBcr2FieldMtypShift))
}

const (
	RegisterBcr2FieldMwidShift = 4
	RegisterBcr2FieldMwidMask  = 0x30
)

// GetMwid Memory data bus width Defines the external memory device width, valid for all type of memories.
func (r *registerBcr2Type) GetMwid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldMwidMask) >> RegisterBcr2FieldMwidShift)
}

// SetMwid Memory data bus width Defines the external memory device width, valid for all type of memories.
func (r *registerBcr2Type) SetMwid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldMwidMask)|(uint32(value)<<RegisterBcr2FieldMwidShift))
}

const (
	RegisterBcr2FieldFaccenShift = 6
	RegisterBcr2FieldFaccenMask  = 0x40
)

// GetFaccen Flash access enable This bit enables NOR Flash memory access operations.
func (r *registerBcr2Type) GetFaccen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldFaccenMask) != 0
}

// SetFaccen Flash access enable This bit enables NOR Flash memory access operations.
func (r *registerBcr2Type) SetFaccen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldFaccenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldFaccenMask)
	}
}

const (
	RegisterBcr2FieldBurstenShift = 8
	RegisterBcr2FieldBurstenMask  = 0x100
)

// GetBursten Burst enable bit This bit enables/disables synchronous accesses during read operations. It is valid only for synchronous memories operating in Burst mode:
func (r *registerBcr2Type) GetBursten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldBurstenMask) != 0
}

// SetBursten Burst enable bit This bit enables/disables synchronous accesses during read operations. It is valid only for synchronous memories operating in Burst mode:
func (r *registerBcr2Type) SetBursten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldBurstenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldBurstenMask)
	}
}

const (
	RegisterBcr2FieldWaitpolShift = 9
	RegisterBcr2FieldWaitpolMask  = 0x200
)

// GetWaitpol Wait signal polarity bit This bit defines the polarity of the wait signal from memory used for either in synchronous or asynchronous mode:
func (r *registerBcr2Type) GetWaitpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldWaitpolMask) != 0
}

// SetWaitpol Wait signal polarity bit This bit defines the polarity of the wait signal from memory used for either in synchronous or asynchronous mode:
func (r *registerBcr2Type) SetWaitpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldWaitpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldWaitpolMask)
	}
}

const (
	RegisterBcr2FieldWaitcfgShift = 11
	RegisterBcr2FieldWaitcfgMask  = 0x800
)

// GetWaitcfg Wait timing configuration The NWAIT signal indicates whether the data from the memory are valid or if a wait state must be inserted when accessing the memory in synchronous mode. This configuration bit determines if NWAIT is asserted by the memory one clock cycle before the wait state or during the wait state:
func (r *registerBcr2Type) GetWaitcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldWaitcfgMask) != 0
}

// SetWaitcfg Wait timing configuration The NWAIT signal indicates whether the data from the memory are valid or if a wait state must be inserted when accessing the memory in synchronous mode. This configuration bit determines if NWAIT is asserted by the memory one clock cycle before the wait state or during the wait state:
func (r *registerBcr2Type) SetWaitcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldWaitcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldWaitcfgMask)
	}
}

const (
	RegisterBcr2FieldWrenShift = 12
	RegisterBcr2FieldWrenMask  = 0x1000
)

// GetWren Write enable bit This bit indicates whether write operations are enabled/disabled in the bank by the FMC:
func (r *registerBcr2Type) GetWren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldWrenMask) != 0
}

// SetWren Write enable bit This bit indicates whether write operations are enabled/disabled in the bank by the FMC:
func (r *registerBcr2Type) SetWren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldWrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldWrenMask)
	}
}

const (
	RegisterBcr2FieldWaitenShift = 13
	RegisterBcr2FieldWaitenMask  = 0x2000
)

// GetWaiten Wait enable bit This bit enables/disables wait-state insertion via the NWAIT signal when accessing the memory in synchronous mode.
func (r *registerBcr2Type) GetWaiten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldWaitenMask) != 0
}

// SetWaiten Wait enable bit This bit enables/disables wait-state insertion via the NWAIT signal when accessing the memory in synchronous mode.
func (r *registerBcr2Type) SetWaiten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldWaitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldWaitenMask)
	}
}

const (
	RegisterBcr2FieldExtmodShift = 14
	RegisterBcr2FieldExtmodMask  = 0x4000
)

// GetExtmod Extended mode enable. This bit enables the FMC to program the write timings for asynchronous accesses inside the FMC_BWTR register, thus resulting in different timings for read and write operations. Note: When the extended mode is disabled, the FMC can operate in Mode1 or Mode2 as follows: ** Mode 1 is the default mode when the SRAM/PSRAM memory type is selected (MTYP =0x0 or 0x01) ** Mode 2 is the default mode when the NOR memory type is selected (MTYP = 0x10).
func (r *registerBcr2Type) GetExtmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldExtmodMask) != 0
}

// SetExtmod Extended mode enable. This bit enables the FMC to program the write timings for asynchronous accesses inside the FMC_BWTR register, thus resulting in different timings for read and write operations. Note: When the extended mode is disabled, the FMC can operate in Mode1 or Mode2 as follows: ** Mode 1 is the default mode when the SRAM/PSRAM memory type is selected (MTYP =0x0 or 0x01) ** Mode 2 is the default mode when the NOR memory type is selected (MTYP = 0x10).
func (r *registerBcr2Type) SetExtmod(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldExtmodMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldExtmodMask)
	}
}

const (
	RegisterBcr2FieldAsyncwaitShift = 15
	RegisterBcr2FieldAsyncwaitMask  = 0x8000
)

// GetAsyncwait Wait signal during asynchronous transfers This bit enables/disables the FMC to use the wait signal even during an asynchronous protocol.
func (r *registerBcr2Type) GetAsyncwait() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldAsyncwaitMask) != 0
}

// SetAsyncwait Wait signal during asynchronous transfers This bit enables/disables the FMC to use the wait signal even during an asynchronous protocol.
func (r *registerBcr2Type) SetAsyncwait(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldAsyncwaitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldAsyncwaitMask)
	}
}

const (
	RegisterBcr2FieldCpsizeShift = 16
	RegisterBcr2FieldCpsizeMask  = 0x70000
)

// GetCpsize CRAM Page Size These are used for Cellular RAM 1.5 which does not allow burst access to cross the address boundaries between pages. When these bits are configured, the FMC controller splits automatically the burst access when the memory page size is reached (refer to memory datasheet for page size). Other configuration: reserved.
func (r *registerBcr2Type) GetCpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldCpsizeMask) >> RegisterBcr2FieldCpsizeShift)
}

// SetCpsize CRAM Page Size These are used for Cellular RAM 1.5 which does not allow burst access to cross the address boundaries between pages. When these bits are configured, the FMC controller splits automatically the burst access when the memory page size is reached (refer to memory datasheet for page size). Other configuration: reserved.
func (r *registerBcr2Type) SetCpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldCpsizeMask)|(uint32(value)<<RegisterBcr2FieldCpsizeShift))
}

const (
	RegisterBcr2FieldCburstrwShift = 19
	RegisterBcr2FieldCburstrwMask  = 0x80000
)

// GetCburstrw Write burst enable For PSRAM (CRAM) operating in Burst mode, the bit enables synchronous accesses during write operations. The enable bit for synchronous read accesses is the BURSTEN bit in the FMC_BCRx register.
func (r *registerBcr2Type) GetCburstrw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldCburstrwMask) != 0
}

// SetCburstrw Write burst enable For PSRAM (CRAM) operating in Burst mode, the bit enables synchronous accesses during write operations. The enable bit for synchronous read accesses is the BURSTEN bit in the FMC_BCRx register.
func (r *registerBcr2Type) SetCburstrw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldCburstrwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldCburstrwMask)
	}
}

const (
	RegisterBcr2FieldCclkenShift = 20
	RegisterBcr2FieldCclkenMask  = 0x100000
)

// GetCclken Continuous Clock Enable This bit enables the FMC_CLK clock output to external memory devices. Note: The CCLKEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register. Bank 1 must be configured in synchronous mode to generate the FMC_CLK continuous clock. If CCLKEN bit is set, the FMC_CLK clock ratio is specified by CLKDIV value in the FMC_BTR1 register. CLKDIV in FMC_BWTR1 is dont care. If the synchronous mode is used and CCLKEN bit is set, the synchronous memories connected to other banks than Bank 1 are clocked by the same clock (the CLKDIV value in the FMC_BTR2..4 and FMC_BWTR2..4 registers for other banks has no effect.)
func (r *registerBcr2Type) GetCclken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldCclkenMask) != 0
}

// SetCclken Continuous Clock Enable This bit enables the FMC_CLK clock output to external memory devices. Note: The CCLKEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register. Bank 1 must be configured in synchronous mode to generate the FMC_CLK continuous clock. If CCLKEN bit is set, the FMC_CLK clock ratio is specified by CLKDIV value in the FMC_BTR1 register. CLKDIV in FMC_BWTR1 is dont care. If the synchronous mode is used and CCLKEN bit is set, the synchronous memories connected to other banks than Bank 1 are clocked by the same clock (the CLKDIV value in the FMC_BTR2..4 and FMC_BWTR2..4 registers for other banks has no effect.)
func (r *registerBcr2Type) SetCclken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldCclkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldCclkenMask)
	}
}

const (
	RegisterBcr2FieldWfdisShift = 21
	RegisterBcr2FieldWfdisMask  = 0x200000
)

// GetWfdis Write FIFO Disable This bit disables the Write FIFO used by the FMC controller. Note: The WFDIS bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr2Type) GetWfdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldWfdisMask) != 0
}

// SetWfdis Write FIFO Disable This bit disables the Write FIFO used by the FMC controller. Note: The WFDIS bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr2Type) SetWfdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldWfdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldWfdisMask)
	}
}

const (
	RegisterBcr2FieldBmapShift = 24
	RegisterBcr2FieldBmapMask  = 0x3000000
)

// GetBmap FMC bank mapping These bits allows different to remap SDRAM bank2 or swap the FMC NOR/PSRAM and SDRAM banks.Refer to Table 10 for Note: The BMAP bits of the FMC_BCR2..4 registers are dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr2Type) GetBmap() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldBmapMask) >> RegisterBcr2FieldBmapShift)
}

// SetBmap FMC bank mapping These bits allows different to remap SDRAM bank2 or swap the FMC NOR/PSRAM and SDRAM banks.Refer to Table 10 for Note: The BMAP bits of the FMC_BCR2..4 registers are dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr2Type) SetBmap(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldBmapMask)|(uint32(value)<<RegisterBcr2FieldBmapShift))
}

const (
	RegisterBcr2FieldFmcenShift = 31
	RegisterBcr2FieldFmcenMask  = 0x80000000
)

// GetFmcen FMC controller Enable This bit enables/disables the FMC controller. Note: The FMCEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr2Type) GetFmcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr2FieldFmcenMask) != 0
}

// SetFmcen FMC controller Enable This bit enables/disables the FMC controller. Note: The FMCEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr2Type) SetFmcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr2FieldFmcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr2FieldFmcenMask)
	}
}

// registerBtr2Type This register contains the control information of each memory bank, used for SRAMs, PSRAM and NOR Flash memories.If the EXTMOD bit is set in the FMC_BCRx register, then this register is partitioned for write and read access, that is, 2 registers are available: one to configure read accesses (this register) and one to configure write accesses (FMC_BWTRx registers).
type registerBtr2Type uint32

const (
	RegisterBtr2FieldAddsetShift = 0
	RegisterBtr2FieldAddsetMask  = 0xf
)

// GetAddset Address setup phase duration These bits are written by software to define the duration of the address setup phase (refer to Figure81 to Figure93), used in SRAMs, ROMs and asynchronous NOR Flash: For each access mode address setup phase duration, please refer to the respective figure (refer to Figure81 to Figure93). Note: In synchronous accesses, this value is dont care. In Muxed mode or Mode D, the minimum value for ADDSET is 1.
func (r *registerBtr2Type) GetAddset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr2FieldAddsetMask) >> RegisterBtr2FieldAddsetShift)
}

// SetAddset Address setup phase duration These bits are written by software to define the duration of the address setup phase (refer to Figure81 to Figure93), used in SRAMs, ROMs and asynchronous NOR Flash: For each access mode address setup phase duration, please refer to the respective figure (refer to Figure81 to Figure93). Note: In synchronous accesses, this value is dont care. In Muxed mode or Mode D, the minimum value for ADDSET is 1.
func (r *registerBtr2Type) SetAddset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr2FieldAddsetMask)|(uint32(value)<<RegisterBtr2FieldAddsetShift))
}

const (
	RegisterBtr2FieldAddhldShift = 4
	RegisterBtr2FieldAddhldMask  = 0xf0
)

// GetAddhld Address-hold phase duration These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in mode D or multiplexed accesses: For each access mode address-hold phase duration, please refer to the respective figure (Figure81 to Figure93). Note: In synchronous accesses, this value is not used, the address hold phase is always 1 memory clock period duration.
func (r *registerBtr2Type) GetAddhld() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr2FieldAddhldMask) >> RegisterBtr2FieldAddhldShift)
}

// SetAddhld Address-hold phase duration These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in mode D or multiplexed accesses: For each access mode address-hold phase duration, please refer to the respective figure (Figure81 to Figure93). Note: In synchronous accesses, this value is not used, the address hold phase is always 1 memory clock period duration.
func (r *registerBtr2Type) SetAddhld(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr2FieldAddhldMask)|(uint32(value)<<RegisterBtr2FieldAddhldShift))
}

const (
	RegisterBtr2FieldDatastShift = 8
	RegisterBtr2FieldDatastMask  = 0xff00
)

// GetDatast Data-phase duration These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous accesses: For each memory type and access mode data-phase duration, please refer to the respective figure (Figure81 to Figure93). Example: Mode1, write access, DATAST=1: Data-phase duration= DATAST+1 = 2 KCK_FMC clock cycles. Note: In synchronous accesses, this value is dont care.
func (r *registerBtr2Type) GetDatast() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr2FieldDatastMask) >> RegisterBtr2FieldDatastShift)
}

// SetDatast Data-phase duration These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous accesses: For each memory type and access mode data-phase duration, please refer to the respective figure (Figure81 to Figure93). Example: Mode1, write access, DATAST=1: Data-phase duration= DATAST+1 = 2 KCK_FMC clock cycles. Note: In synchronous accesses, this value is dont care.
func (r *registerBtr2Type) SetDatast(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr2FieldDatastMask)|(uint32(value)<<RegisterBtr2FieldDatastShift))
}

const (
	RegisterBtr2FieldBusturnShift = 16
	RegisterBtr2FieldBusturnMask  = 0xf0000
)

// GetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write-to-read or read-to write transaction. The programmed bus turnaround delay is inserted between an asynchronous read (in muxed or mode D) or write transaction and any other asynchronous /synchronous read/write from/to a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except in muxed mode and mode D. There is a bus turnaround delay of 1 FMC clock cycle between: Two consecutive asynchronous read transfers to the same static memory bank except for modes muxed and D. An asynchronous read to an asynchronous or synchronous write to any static bank or dynamic bank except in modes muxed and D mode. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank. A synchronous write (burst or single) access and an asynchronous write or read transfer to or from static memory bank (the bank can be the same or a different one in case of a read operation. Two consecutive synchronous read operations (in Burst or Single mode) followed by any synchronous/asynchronous read or write from/to another static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write access (in Burst or Single mode) and a synchronous read from the same or a different bank. The bus turnaround delay allows to match the minimum time between consecutive transactions (tEHEL from NEx high to NEx low) and the maximum time required by the memory to free the data bus after a read access (tEHQZ): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin and (BUSTRUN + 2)KCK_FMC period &#8805; tEHQZmax if EXTMOD = 0 (BUSTRUN + 2)KCK_FMC period &#8805; max (tEHELmin, tEHQZmax) if EXTMOD = 1. ...
func (r *registerBtr2Type) GetBusturn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr2FieldBusturnMask) >> RegisterBtr2FieldBusturnShift)
}

// SetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write-to-read or read-to write transaction. The programmed bus turnaround delay is inserted between an asynchronous read (in muxed or mode D) or write transaction and any other asynchronous /synchronous read/write from/to a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except in muxed mode and mode D. There is a bus turnaround delay of 1 FMC clock cycle between: Two consecutive asynchronous read transfers to the same static memory bank except for modes muxed and D. An asynchronous read to an asynchronous or synchronous write to any static bank or dynamic bank except in modes muxed and D mode. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank. A synchronous write (burst or single) access and an asynchronous write or read transfer to or from static memory bank (the bank can be the same or a different one in case of a read operation. Two consecutive synchronous read operations (in Burst or Single mode) followed by any synchronous/asynchronous read or write from/to another static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write access (in Burst or Single mode) and a synchronous read from the same or a different bank. The bus turnaround delay allows to match the minimum time between consecutive transactions (tEHEL from NEx high to NEx low) and the maximum time required by the memory to free the data bus after a read access (tEHQZ): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin and (BUSTRUN + 2)KCK_FMC period &#8805; tEHQZmax if EXTMOD = 0 (BUSTRUN + 2)KCK_FMC period &#8805; max (tEHELmin, tEHQZmax) if EXTMOD = 1. ...
func (r *registerBtr2Type) SetBusturn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr2FieldBusturnMask)|(uint32(value)<<RegisterBtr2FieldBusturnShift))
}

const (
	RegisterBtr2FieldClkdivShift = 20
	RegisterBtr2FieldClkdivMask  = 0xf00000
)

// GetClkdiv Clock divide ratio (for FMC_CLK signal) These bits define the period of FMC_CLK clock output signal, expressed in number of KCK_FMC cycles: In asynchronous NOR Flash, SRAM or PSRAM accesses, this value is dont care. Note: Refer to Section20.6.5: Synchronous transactions for FMC_CLK divider ratio formula)
func (r *registerBtr2Type) GetClkdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr2FieldClkdivMask) >> RegisterBtr2FieldClkdivShift)
}

// SetClkdiv Clock divide ratio (for FMC_CLK signal) These bits define the period of FMC_CLK clock output signal, expressed in number of KCK_FMC cycles: In asynchronous NOR Flash, SRAM or PSRAM accesses, this value is dont care. Note: Refer to Section20.6.5: Synchronous transactions for FMC_CLK divider ratio formula)
func (r *registerBtr2Type) SetClkdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr2FieldClkdivMask)|(uint32(value)<<RegisterBtr2FieldClkdivShift))
}

const (
	RegisterBtr2FieldDatlatShift = 24
	RegisterBtr2FieldDatlatMask  = 0xf000000
)

// GetDatlat Data latency for synchronous memory For synchronous access with read write burst mode enabled these bits define the number of memory clock cycles
func (r *registerBtr2Type) GetDatlat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr2FieldDatlatMask) >> RegisterBtr2FieldDatlatShift)
}

// SetDatlat Data latency for synchronous memory For synchronous access with read write burst mode enabled these bits define the number of memory clock cycles
func (r *registerBtr2Type) SetDatlat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr2FieldDatlatMask)|(uint32(value)<<RegisterBtr2FieldDatlatShift))
}

const (
	RegisterBtr2FieldAccmodShift = 28
	RegisterBtr2FieldAccmodMask  = 0x30000000
)

// GetAccmod Access mode These bits specify the asynchronous access modes as shown in the timing diagrams. They are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBtr2Type) GetAccmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr2FieldAccmodMask) >> RegisterBtr2FieldAccmodShift)
}

// SetAccmod Access mode These bits specify the asynchronous access modes as shown in the timing diagrams. They are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBtr2Type) SetAccmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr2FieldAccmodMask)|(uint32(value)<<RegisterBtr2FieldAccmodShift))
}

// registerBcr3Type This register contains the control information of each memory bank, used for SRAMs, PSRAM and NOR Flash memories.
type registerBcr3Type uint32

const (
	RegisterBcr3FieldMbkenShift = 0
	RegisterBcr3FieldMbkenMask  = 0x1
)

// GetMbken Memory bank enable bit This bit enables the memory bank. After reset Bank1 is enabled, all others are disabled. Accessing a disabled bank causes an ERROR on AXI bus.
func (r *registerBcr3Type) GetMbken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldMbkenMask) != 0
}

// SetMbken Memory bank enable bit This bit enables the memory bank. After reset Bank1 is enabled, all others are disabled. Accessing a disabled bank causes an ERROR on AXI bus.
func (r *registerBcr3Type) SetMbken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldMbkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldMbkenMask)
	}
}

const (
	RegisterBcr3FieldMuxenShift = 1
	RegisterBcr3FieldMuxenMask  = 0x2
)

// GetMuxen Address/data multiplexing enable bit When this bit is set, the address and data values are multiplexed on the data bus, valid only with NOR and PSRAM memories:
func (r *registerBcr3Type) GetMuxen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldMuxenMask) != 0
}

// SetMuxen Address/data multiplexing enable bit When this bit is set, the address and data values are multiplexed on the data bus, valid only with NOR and PSRAM memories:
func (r *registerBcr3Type) SetMuxen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldMuxenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldMuxenMask)
	}
}

const (
	RegisterBcr3FieldMtypShift = 2
	RegisterBcr3FieldMtypMask  = 0xc
)

// GetMtyp Memory type These bits define the type of external memory attached to the corresponding memory bank:
func (r *registerBcr3Type) GetMtyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldMtypMask) >> RegisterBcr3FieldMtypShift)
}

// SetMtyp Memory type These bits define the type of external memory attached to the corresponding memory bank:
func (r *registerBcr3Type) SetMtyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldMtypMask)|(uint32(value)<<RegisterBcr3FieldMtypShift))
}

const (
	RegisterBcr3FieldMwidShift = 4
	RegisterBcr3FieldMwidMask  = 0x30
)

// GetMwid Memory data bus width Defines the external memory device width, valid for all type of memories.
func (r *registerBcr3Type) GetMwid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldMwidMask) >> RegisterBcr3FieldMwidShift)
}

// SetMwid Memory data bus width Defines the external memory device width, valid for all type of memories.
func (r *registerBcr3Type) SetMwid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldMwidMask)|(uint32(value)<<RegisterBcr3FieldMwidShift))
}

const (
	RegisterBcr3FieldFaccenShift = 6
	RegisterBcr3FieldFaccenMask  = 0x40
)

// GetFaccen Flash access enable This bit enables NOR Flash memory access operations.
func (r *registerBcr3Type) GetFaccen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldFaccenMask) != 0
}

// SetFaccen Flash access enable This bit enables NOR Flash memory access operations.
func (r *registerBcr3Type) SetFaccen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldFaccenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldFaccenMask)
	}
}

const (
	RegisterBcr3FieldBurstenShift = 8
	RegisterBcr3FieldBurstenMask  = 0x100
)

// GetBursten Burst enable bit This bit enables/disables synchronous accesses during read operations. It is valid only for synchronous memories operating in Burst mode:
func (r *registerBcr3Type) GetBursten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldBurstenMask) != 0
}

// SetBursten Burst enable bit This bit enables/disables synchronous accesses during read operations. It is valid only for synchronous memories operating in Burst mode:
func (r *registerBcr3Type) SetBursten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldBurstenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldBurstenMask)
	}
}

const (
	RegisterBcr3FieldWaitpolShift = 9
	RegisterBcr3FieldWaitpolMask  = 0x200
)

// GetWaitpol Wait signal polarity bit This bit defines the polarity of the wait signal from memory used for either in synchronous or asynchronous mode:
func (r *registerBcr3Type) GetWaitpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldWaitpolMask) != 0
}

// SetWaitpol Wait signal polarity bit This bit defines the polarity of the wait signal from memory used for either in synchronous or asynchronous mode:
func (r *registerBcr3Type) SetWaitpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldWaitpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldWaitpolMask)
	}
}

const (
	RegisterBcr3FieldWaitcfgShift = 11
	RegisterBcr3FieldWaitcfgMask  = 0x800
)

// GetWaitcfg Wait timing configuration The NWAIT signal indicates whether the data from the memory are valid or if a wait state must be inserted when accessing the memory in synchronous mode. This configuration bit determines if NWAIT is asserted by the memory one clock cycle before the wait state or during the wait state:
func (r *registerBcr3Type) GetWaitcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldWaitcfgMask) != 0
}

// SetWaitcfg Wait timing configuration The NWAIT signal indicates whether the data from the memory are valid or if a wait state must be inserted when accessing the memory in synchronous mode. This configuration bit determines if NWAIT is asserted by the memory one clock cycle before the wait state or during the wait state:
func (r *registerBcr3Type) SetWaitcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldWaitcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldWaitcfgMask)
	}
}

const (
	RegisterBcr3FieldWrenShift = 12
	RegisterBcr3FieldWrenMask  = 0x1000
)

// GetWren Write enable bit This bit indicates whether write operations are enabled/disabled in the bank by the FMC:
func (r *registerBcr3Type) GetWren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldWrenMask) != 0
}

// SetWren Write enable bit This bit indicates whether write operations are enabled/disabled in the bank by the FMC:
func (r *registerBcr3Type) SetWren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldWrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldWrenMask)
	}
}

const (
	RegisterBcr3FieldWaitenShift = 13
	RegisterBcr3FieldWaitenMask  = 0x2000
)

// GetWaiten Wait enable bit This bit enables/disables wait-state insertion via the NWAIT signal when accessing the memory in synchronous mode.
func (r *registerBcr3Type) GetWaiten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldWaitenMask) != 0
}

// SetWaiten Wait enable bit This bit enables/disables wait-state insertion via the NWAIT signal when accessing the memory in synchronous mode.
func (r *registerBcr3Type) SetWaiten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldWaitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldWaitenMask)
	}
}

const (
	RegisterBcr3FieldExtmodShift = 14
	RegisterBcr3FieldExtmodMask  = 0x4000
)

// GetExtmod Extended mode enable. This bit enables the FMC to program the write timings for asynchronous accesses inside the FMC_BWTR register, thus resulting in different timings for read and write operations. Note: When the extended mode is disabled, the FMC can operate in Mode1 or Mode2 as follows: ** Mode 1 is the default mode when the SRAM/PSRAM memory type is selected (MTYP =0x0 or 0x01) ** Mode 2 is the default mode when the NOR memory type is selected (MTYP = 0x10).
func (r *registerBcr3Type) GetExtmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldExtmodMask) != 0
}

// SetExtmod Extended mode enable. This bit enables the FMC to program the write timings for asynchronous accesses inside the FMC_BWTR register, thus resulting in different timings for read and write operations. Note: When the extended mode is disabled, the FMC can operate in Mode1 or Mode2 as follows: ** Mode 1 is the default mode when the SRAM/PSRAM memory type is selected (MTYP =0x0 or 0x01) ** Mode 2 is the default mode when the NOR memory type is selected (MTYP = 0x10).
func (r *registerBcr3Type) SetExtmod(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldExtmodMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldExtmodMask)
	}
}

const (
	RegisterBcr3FieldAsyncwaitShift = 15
	RegisterBcr3FieldAsyncwaitMask  = 0x8000
)

// GetAsyncwait Wait signal during asynchronous transfers This bit enables/disables the FMC to use the wait signal even during an asynchronous protocol.
func (r *registerBcr3Type) GetAsyncwait() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldAsyncwaitMask) != 0
}

// SetAsyncwait Wait signal during asynchronous transfers This bit enables/disables the FMC to use the wait signal even during an asynchronous protocol.
func (r *registerBcr3Type) SetAsyncwait(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldAsyncwaitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldAsyncwaitMask)
	}
}

const (
	RegisterBcr3FieldCpsizeShift = 16
	RegisterBcr3FieldCpsizeMask  = 0x70000
)

// GetCpsize CRAM Page Size These are used for Cellular RAM 1.5 which does not allow burst access to cross the address boundaries between pages. When these bits are configured, the FMC controller splits automatically the burst access when the memory page size is reached (refer to memory datasheet for page size). Other configuration: reserved.
func (r *registerBcr3Type) GetCpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldCpsizeMask) >> RegisterBcr3FieldCpsizeShift)
}

// SetCpsize CRAM Page Size These are used for Cellular RAM 1.5 which does not allow burst access to cross the address boundaries between pages. When these bits are configured, the FMC controller splits automatically the burst access when the memory page size is reached (refer to memory datasheet for page size). Other configuration: reserved.
func (r *registerBcr3Type) SetCpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldCpsizeMask)|(uint32(value)<<RegisterBcr3FieldCpsizeShift))
}

const (
	RegisterBcr3FieldCburstrwShift = 19
	RegisterBcr3FieldCburstrwMask  = 0x80000
)

// GetCburstrw Write burst enable For PSRAM (CRAM) operating in Burst mode, the bit enables synchronous accesses during write operations. The enable bit for synchronous read accesses is the BURSTEN bit in the FMC_BCRx register.
func (r *registerBcr3Type) GetCburstrw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldCburstrwMask) != 0
}

// SetCburstrw Write burst enable For PSRAM (CRAM) operating in Burst mode, the bit enables synchronous accesses during write operations. The enable bit for synchronous read accesses is the BURSTEN bit in the FMC_BCRx register.
func (r *registerBcr3Type) SetCburstrw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldCburstrwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldCburstrwMask)
	}
}

const (
	RegisterBcr3FieldCclkenShift = 20
	RegisterBcr3FieldCclkenMask  = 0x100000
)

// GetCclken Continuous Clock Enable This bit enables the FMC_CLK clock output to external memory devices. Note: The CCLKEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register. Bank 1 must be configured in synchronous mode to generate the FMC_CLK continuous clock. If CCLKEN bit is set, the FMC_CLK clock ratio is specified by CLKDIV value in the FMC_BTR1 register. CLKDIV in FMC_BWTR1 is dont care. If the synchronous mode is used and CCLKEN bit is set, the synchronous memories connected to other banks than Bank 1 are clocked by the same clock (the CLKDIV value in the FMC_BTR2..4 and FMC_BWTR2..4 registers for other banks has no effect.)
func (r *registerBcr3Type) GetCclken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldCclkenMask) != 0
}

// SetCclken Continuous Clock Enable This bit enables the FMC_CLK clock output to external memory devices. Note: The CCLKEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register. Bank 1 must be configured in synchronous mode to generate the FMC_CLK continuous clock. If CCLKEN bit is set, the FMC_CLK clock ratio is specified by CLKDIV value in the FMC_BTR1 register. CLKDIV in FMC_BWTR1 is dont care. If the synchronous mode is used and CCLKEN bit is set, the synchronous memories connected to other banks than Bank 1 are clocked by the same clock (the CLKDIV value in the FMC_BTR2..4 and FMC_BWTR2..4 registers for other banks has no effect.)
func (r *registerBcr3Type) SetCclken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldCclkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldCclkenMask)
	}
}

const (
	RegisterBcr3FieldWfdisShift = 21
	RegisterBcr3FieldWfdisMask  = 0x200000
)

// GetWfdis Write FIFO Disable This bit disables the Write FIFO used by the FMC controller. Note: The WFDIS bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr3Type) GetWfdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldWfdisMask) != 0
}

// SetWfdis Write FIFO Disable This bit disables the Write FIFO used by the FMC controller. Note: The WFDIS bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr3Type) SetWfdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldWfdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldWfdisMask)
	}
}

const (
	RegisterBcr3FieldBmapShift = 24
	RegisterBcr3FieldBmapMask  = 0x3000000
)

// GetBmap FMC bank mapping These bits allows different to remap SDRAM bank2 or swap the FMC NOR/PSRAM and SDRAM banks.Refer to Table 10 for Note: The BMAP bits of the FMC_BCR2..4 registers are dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr3Type) GetBmap() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldBmapMask) >> RegisterBcr3FieldBmapShift)
}

// SetBmap FMC bank mapping These bits allows different to remap SDRAM bank2 or swap the FMC NOR/PSRAM and SDRAM banks.Refer to Table 10 for Note: The BMAP bits of the FMC_BCR2..4 registers are dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr3Type) SetBmap(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldBmapMask)|(uint32(value)<<RegisterBcr3FieldBmapShift))
}

const (
	RegisterBcr3FieldFmcenShift = 31
	RegisterBcr3FieldFmcenMask  = 0x80000000
)

// GetFmcen FMC controller Enable This bit enables/disables the FMC controller. Note: The FMCEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr3Type) GetFmcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr3FieldFmcenMask) != 0
}

// SetFmcen FMC controller Enable This bit enables/disables the FMC controller. Note: The FMCEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr3Type) SetFmcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr3FieldFmcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr3FieldFmcenMask)
	}
}

// registerBtr3Type This register contains the control information of each memory bank, used for SRAMs, PSRAM and NOR Flash memories.If the EXTMOD bit is set in the FMC_BCRx register, then this register is partitioned for write and read access, that is, 2 registers are available: one to configure read accesses (this register) and one to configure write accesses (FMC_BWTRx registers).
type registerBtr3Type uint32

const (
	RegisterBtr3FieldAddsetShift = 0
	RegisterBtr3FieldAddsetMask  = 0xf
)

// GetAddset Address setup phase duration These bits are written by software to define the duration of the address setup phase (refer to Figure81 to Figure93), used in SRAMs, ROMs and asynchronous NOR Flash: For each access mode address setup phase duration, please refer to the respective figure (refer to Figure81 to Figure93). Note: In synchronous accesses, this value is dont care. In Muxed mode or Mode D, the minimum value for ADDSET is 1.
func (r *registerBtr3Type) GetAddset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr3FieldAddsetMask) >> RegisterBtr3FieldAddsetShift)
}

// SetAddset Address setup phase duration These bits are written by software to define the duration of the address setup phase (refer to Figure81 to Figure93), used in SRAMs, ROMs and asynchronous NOR Flash: For each access mode address setup phase duration, please refer to the respective figure (refer to Figure81 to Figure93). Note: In synchronous accesses, this value is dont care. In Muxed mode or Mode D, the minimum value for ADDSET is 1.
func (r *registerBtr3Type) SetAddset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr3FieldAddsetMask)|(uint32(value)<<RegisterBtr3FieldAddsetShift))
}

const (
	RegisterBtr3FieldAddhldShift = 4
	RegisterBtr3FieldAddhldMask  = 0xf0
)

// GetAddhld Address-hold phase duration These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in mode D or multiplexed accesses: For each access mode address-hold phase duration, please refer to the respective figure (Figure81 to Figure93). Note: In synchronous accesses, this value is not used, the address hold phase is always 1 memory clock period duration.
func (r *registerBtr3Type) GetAddhld() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr3FieldAddhldMask) >> RegisterBtr3FieldAddhldShift)
}

// SetAddhld Address-hold phase duration These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in mode D or multiplexed accesses: For each access mode address-hold phase duration, please refer to the respective figure (Figure81 to Figure93). Note: In synchronous accesses, this value is not used, the address hold phase is always 1 memory clock period duration.
func (r *registerBtr3Type) SetAddhld(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr3FieldAddhldMask)|(uint32(value)<<RegisterBtr3FieldAddhldShift))
}

const (
	RegisterBtr3FieldDatastShift = 8
	RegisterBtr3FieldDatastMask  = 0xff00
)

// GetDatast Data-phase duration These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous accesses: For each memory type and access mode data-phase duration, please refer to the respective figure (Figure81 to Figure93). Example: Mode1, write access, DATAST=1: Data-phase duration= DATAST+1 = 2 KCK_FMC clock cycles. Note: In synchronous accesses, this value is dont care.
func (r *registerBtr3Type) GetDatast() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr3FieldDatastMask) >> RegisterBtr3FieldDatastShift)
}

// SetDatast Data-phase duration These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous accesses: For each memory type and access mode data-phase duration, please refer to the respective figure (Figure81 to Figure93). Example: Mode1, write access, DATAST=1: Data-phase duration= DATAST+1 = 2 KCK_FMC clock cycles. Note: In synchronous accesses, this value is dont care.
func (r *registerBtr3Type) SetDatast(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr3FieldDatastMask)|(uint32(value)<<RegisterBtr3FieldDatastShift))
}

const (
	RegisterBtr3FieldBusturnShift = 16
	RegisterBtr3FieldBusturnMask  = 0xf0000
)

// GetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write-to-read or read-to write transaction. The programmed bus turnaround delay is inserted between an asynchronous read (in muxed or mode D) or write transaction and any other asynchronous /synchronous read/write from/to a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except in muxed mode and mode D. There is a bus turnaround delay of 1 FMC clock cycle between: Two consecutive asynchronous read transfers to the same static memory bank except for modes muxed and D. An asynchronous read to an asynchronous or synchronous write to any static bank or dynamic bank except in modes muxed and D mode. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank. A synchronous write (burst or single) access and an asynchronous write or read transfer to or from static memory bank (the bank can be the same or a different one in case of a read operation. Two consecutive synchronous read operations (in Burst or Single mode) followed by any synchronous/asynchronous read or write from/to another static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write access (in Burst or Single mode) and a synchronous read from the same or a different bank. The bus turnaround delay allows to match the minimum time between consecutive transactions (tEHEL from NEx high to NEx low) and the maximum time required by the memory to free the data bus after a read access (tEHQZ): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin and (BUSTRUN + 2)KCK_FMC period &#8805; tEHQZmax if EXTMOD = 0 (BUSTRUN + 2)KCK_FMC period &#8805; max (tEHELmin, tEHQZmax) if EXTMOD =1. ...
func (r *registerBtr3Type) GetBusturn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr3FieldBusturnMask) >> RegisterBtr3FieldBusturnShift)
}

// SetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write-to-read or read-to write transaction. The programmed bus turnaround delay is inserted between an asynchronous read (in muxed or mode D) or write transaction and any other asynchronous /synchronous read/write from/to a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except in muxed mode and mode D. There is a bus turnaround delay of 1 FMC clock cycle between: Two consecutive asynchronous read transfers to the same static memory bank except for modes muxed and D. An asynchronous read to an asynchronous or synchronous write to any static bank or dynamic bank except in modes muxed and D mode. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank. A synchronous write (burst or single) access and an asynchronous write or read transfer to or from static memory bank (the bank can be the same or a different one in case of a read operation. Two consecutive synchronous read operations (in Burst or Single mode) followed by any synchronous/asynchronous read or write from/to another static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write access (in Burst or Single mode) and a synchronous read from the same or a different bank. The bus turnaround delay allows to match the minimum time between consecutive transactions (tEHEL from NEx high to NEx low) and the maximum time required by the memory to free the data bus after a read access (tEHQZ): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin and (BUSTRUN + 2)KCK_FMC period &#8805; tEHQZmax if EXTMOD = 0 (BUSTRUN + 2)KCK_FMC period &#8805; max (tEHELmin, tEHQZmax) if EXTMOD =1. ...
func (r *registerBtr3Type) SetBusturn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr3FieldBusturnMask)|(uint32(value)<<RegisterBtr3FieldBusturnShift))
}

const (
	RegisterBtr3FieldClkdivShift = 20
	RegisterBtr3FieldClkdivMask  = 0xf00000
)

// GetClkdiv Clock divide ratio (for FMC_CLK signal) These bits define the period of FMC_CLK clock output signal, expressed in number of KCK_FMC cycles: In asynchronous NOR Flash, SRAM or PSRAM accesses, this value is dont care. Note: Refer to Section20.6.5: Synchronous transactions for FMC_CLK divider ratio formula)
func (r *registerBtr3Type) GetClkdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr3FieldClkdivMask) >> RegisterBtr3FieldClkdivShift)
}

// SetClkdiv Clock divide ratio (for FMC_CLK signal) These bits define the period of FMC_CLK clock output signal, expressed in number of KCK_FMC cycles: In asynchronous NOR Flash, SRAM or PSRAM accesses, this value is dont care. Note: Refer to Section20.6.5: Synchronous transactions for FMC_CLK divider ratio formula)
func (r *registerBtr3Type) SetClkdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr3FieldClkdivMask)|(uint32(value)<<RegisterBtr3FieldClkdivShift))
}

const (
	RegisterBtr3FieldDatlatShift = 24
	RegisterBtr3FieldDatlatMask  = 0xf000000
)

// GetDatlat Data latency for synchronous memory For synchronous access with read write burst mode enabled these bits define the number of memory clock cycles
func (r *registerBtr3Type) GetDatlat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr3FieldDatlatMask) >> RegisterBtr3FieldDatlatShift)
}

// SetDatlat Data latency for synchronous memory For synchronous access with read write burst mode enabled these bits define the number of memory clock cycles
func (r *registerBtr3Type) SetDatlat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr3FieldDatlatMask)|(uint32(value)<<RegisterBtr3FieldDatlatShift))
}

const (
	RegisterBtr3FieldAccmodShift = 28
	RegisterBtr3FieldAccmodMask  = 0x30000000
)

// GetAccmod Access mode These bits specify the asynchronous access modes as shown in the timing diagrams. They are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBtr3Type) GetAccmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr3FieldAccmodMask) >> RegisterBtr3FieldAccmodShift)
}

// SetAccmod Access mode These bits specify the asynchronous access modes as shown in the timing diagrams. They are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBtr3Type) SetAccmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr3FieldAccmodMask)|(uint32(value)<<RegisterBtr3FieldAccmodShift))
}

// registerBcr4Type This register contains the control information of each memory bank, used for SRAMs, PSRAM and NOR Flash memories.
type registerBcr4Type uint32

const (
	RegisterBcr4FieldMbkenShift = 0
	RegisterBcr4FieldMbkenMask  = 0x1
)

// GetMbken Memory bank enable bit This bit enables the memory bank. After reset Bank1 is enabled, all others are disabled. Accessing a disabled bank causes an ERROR on AXI bus.
func (r *registerBcr4Type) GetMbken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldMbkenMask) != 0
}

// SetMbken Memory bank enable bit This bit enables the memory bank. After reset Bank1 is enabled, all others are disabled. Accessing a disabled bank causes an ERROR on AXI bus.
func (r *registerBcr4Type) SetMbken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldMbkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldMbkenMask)
	}
}

const (
	RegisterBcr4FieldMuxenShift = 1
	RegisterBcr4FieldMuxenMask  = 0x2
)

// GetMuxen Address/data multiplexing enable bit When this bit is set, the address and data values are multiplexed on the data bus, valid only with NOR and PSRAM memories:
func (r *registerBcr4Type) GetMuxen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldMuxenMask) != 0
}

// SetMuxen Address/data multiplexing enable bit When this bit is set, the address and data values are multiplexed on the data bus, valid only with NOR and PSRAM memories:
func (r *registerBcr4Type) SetMuxen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldMuxenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldMuxenMask)
	}
}

const (
	RegisterBcr4FieldMtypShift = 2
	RegisterBcr4FieldMtypMask  = 0xc
)

// GetMtyp Memory type These bits define the type of external memory attached to the corresponding memory bank:
func (r *registerBcr4Type) GetMtyp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldMtypMask) >> RegisterBcr4FieldMtypShift)
}

// SetMtyp Memory type These bits define the type of external memory attached to the corresponding memory bank:
func (r *registerBcr4Type) SetMtyp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldMtypMask)|(uint32(value)<<RegisterBcr4FieldMtypShift))
}

const (
	RegisterBcr4FieldMwidShift = 4
	RegisterBcr4FieldMwidMask  = 0x30
)

// GetMwid Memory data bus width Defines the external memory device width, valid for all type of memories.
func (r *registerBcr4Type) GetMwid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldMwidMask) >> RegisterBcr4FieldMwidShift)
}

// SetMwid Memory data bus width Defines the external memory device width, valid for all type of memories.
func (r *registerBcr4Type) SetMwid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldMwidMask)|(uint32(value)<<RegisterBcr4FieldMwidShift))
}

const (
	RegisterBcr4FieldFaccenShift = 6
	RegisterBcr4FieldFaccenMask  = 0x40
)

// GetFaccen Flash access enable This bit enables NOR Flash memory access operations.
func (r *registerBcr4Type) GetFaccen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldFaccenMask) != 0
}

// SetFaccen Flash access enable This bit enables NOR Flash memory access operations.
func (r *registerBcr4Type) SetFaccen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldFaccenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldFaccenMask)
	}
}

const (
	RegisterBcr4FieldBurstenShift = 8
	RegisterBcr4FieldBurstenMask  = 0x100
)

// GetBursten Burst enable bit This bit enables/disables synchronous accesses during read operations. It is valid only for synchronous memories operating in Burst mode:
func (r *registerBcr4Type) GetBursten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldBurstenMask) != 0
}

// SetBursten Burst enable bit This bit enables/disables synchronous accesses during read operations. It is valid only for synchronous memories operating in Burst mode:
func (r *registerBcr4Type) SetBursten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldBurstenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldBurstenMask)
	}
}

const (
	RegisterBcr4FieldWaitpolShift = 9
	RegisterBcr4FieldWaitpolMask  = 0x200
)

// GetWaitpol Wait signal polarity bit This bit defines the polarity of the wait signal from memory used for either in synchronous or asynchronous mode:
func (r *registerBcr4Type) GetWaitpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldWaitpolMask) != 0
}

// SetWaitpol Wait signal polarity bit This bit defines the polarity of the wait signal from memory used for either in synchronous or asynchronous mode:
func (r *registerBcr4Type) SetWaitpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldWaitpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldWaitpolMask)
	}
}

const (
	RegisterBcr4FieldWaitcfgShift = 11
	RegisterBcr4FieldWaitcfgMask  = 0x800
)

// GetWaitcfg Wait timing configuration The NWAIT signal indicates whether the data from the memory are valid or if a wait state must be inserted when accessing the memory in synchronous mode. This configuration bit determines if NWAIT is asserted by the memory one clock cycle before the wait state or during the wait state:
func (r *registerBcr4Type) GetWaitcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldWaitcfgMask) != 0
}

// SetWaitcfg Wait timing configuration The NWAIT signal indicates whether the data from the memory are valid or if a wait state must be inserted when accessing the memory in synchronous mode. This configuration bit determines if NWAIT is asserted by the memory one clock cycle before the wait state or during the wait state:
func (r *registerBcr4Type) SetWaitcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldWaitcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldWaitcfgMask)
	}
}

const (
	RegisterBcr4FieldWrenShift = 12
	RegisterBcr4FieldWrenMask  = 0x1000
)

// GetWren Write enable bit This bit indicates whether write operations are enabled/disabled in the bank by the FMC:
func (r *registerBcr4Type) GetWren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldWrenMask) != 0
}

// SetWren Write enable bit This bit indicates whether write operations are enabled/disabled in the bank by the FMC:
func (r *registerBcr4Type) SetWren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldWrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldWrenMask)
	}
}

const (
	RegisterBcr4FieldWaitenShift = 13
	RegisterBcr4FieldWaitenMask  = 0x2000
)

// GetWaiten Wait enable bit This bit enables/disables wait-state insertion via the NWAIT signal when accessing the memory in synchronous mode.
func (r *registerBcr4Type) GetWaiten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldWaitenMask) != 0
}

// SetWaiten Wait enable bit This bit enables/disables wait-state insertion via the NWAIT signal when accessing the memory in synchronous mode.
func (r *registerBcr4Type) SetWaiten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldWaitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldWaitenMask)
	}
}

const (
	RegisterBcr4FieldExtmodShift = 14
	RegisterBcr4FieldExtmodMask  = 0x4000
)

// GetExtmod Extended mode enable. This bit enables the FMC to program the write timings for asynchronous accesses inside the FMC_BWTR register, thus resulting in different timings for read and write operations. Note: When the extended mode is disabled, the FMC can operate in Mode1 or Mode2 as follows: ** Mode 1 is the default mode when the SRAM/PSRAM memory type is selected (MTYP =0x0 or 0x01) ** Mode 2 is the default mode when the NOR memory type is selected (MTYP = 0x10).
func (r *registerBcr4Type) GetExtmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldExtmodMask) != 0
}

// SetExtmod Extended mode enable. This bit enables the FMC to program the write timings for asynchronous accesses inside the FMC_BWTR register, thus resulting in different timings for read and write operations. Note: When the extended mode is disabled, the FMC can operate in Mode1 or Mode2 as follows: ** Mode 1 is the default mode when the SRAM/PSRAM memory type is selected (MTYP =0x0 or 0x01) ** Mode 2 is the default mode when the NOR memory type is selected (MTYP = 0x10).
func (r *registerBcr4Type) SetExtmod(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldExtmodMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldExtmodMask)
	}
}

const (
	RegisterBcr4FieldAsyncwaitShift = 15
	RegisterBcr4FieldAsyncwaitMask  = 0x8000
)

// GetAsyncwait Wait signal during asynchronous transfers This bit enables/disables the FMC to use the wait signal even during an asynchronous protocol.
func (r *registerBcr4Type) GetAsyncwait() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldAsyncwaitMask) != 0
}

// SetAsyncwait Wait signal during asynchronous transfers This bit enables/disables the FMC to use the wait signal even during an asynchronous protocol.
func (r *registerBcr4Type) SetAsyncwait(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldAsyncwaitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldAsyncwaitMask)
	}
}

const (
	RegisterBcr4FieldCpsizeShift = 16
	RegisterBcr4FieldCpsizeMask  = 0x70000
)

// GetCpsize CRAM Page Size These are used for Cellular RAM 1.5 which does not allow burst access to cross the address boundaries between pages. When these bits are configured, the FMC controller splits automatically the burst access when the memory page size is reached (refer to memory datasheet for page size). Other configuration: reserved.
func (r *registerBcr4Type) GetCpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldCpsizeMask) >> RegisterBcr4FieldCpsizeShift)
}

// SetCpsize CRAM Page Size These are used for Cellular RAM 1.5 which does not allow burst access to cross the address boundaries between pages. When these bits are configured, the FMC controller splits automatically the burst access when the memory page size is reached (refer to memory datasheet for page size). Other configuration: reserved.
func (r *registerBcr4Type) SetCpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldCpsizeMask)|(uint32(value)<<RegisterBcr4FieldCpsizeShift))
}

const (
	RegisterBcr4FieldCburstrwShift = 19
	RegisterBcr4FieldCburstrwMask  = 0x80000
)

// GetCburstrw Write burst enable For PSRAM (CRAM) operating in Burst mode, the bit enables synchronous accesses during write operations. The enable bit for synchronous read accesses is the BURSTEN bit in the FMC_BCRx register.
func (r *registerBcr4Type) GetCburstrw() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldCburstrwMask) != 0
}

// SetCburstrw Write burst enable For PSRAM (CRAM) operating in Burst mode, the bit enables synchronous accesses during write operations. The enable bit for synchronous read accesses is the BURSTEN bit in the FMC_BCRx register.
func (r *registerBcr4Type) SetCburstrw(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldCburstrwMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldCburstrwMask)
	}
}

const (
	RegisterBcr4FieldCclkenShift = 20
	RegisterBcr4FieldCclkenMask  = 0x100000
)

// GetCclken Continuous Clock Enable This bit enables the FMC_CLK clock output to external memory devices. Note: The CCLKEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register. Bank 1 must be configured in synchronous mode to generate the FMC_CLK continuous clock. If CCLKEN bit is set, the FMC_CLK clock ratio is specified by CLKDIV value in the FMC_BTR1 register. CLKDIV in FMC_BWTR1 is dont care. If the synchronous mode is used and CCLKEN bit is set, the synchronous memories connected to other banks than Bank 1 are clocked by the same clock (the CLKDIV value in the FMC_BTR2..4 and FMC_BWTR2..4 registers for other banks has no effect.)
func (r *registerBcr4Type) GetCclken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldCclkenMask) != 0
}

// SetCclken Continuous Clock Enable This bit enables the FMC_CLK clock output to external memory devices. Note: The CCLKEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register. Bank 1 must be configured in synchronous mode to generate the FMC_CLK continuous clock. If CCLKEN bit is set, the FMC_CLK clock ratio is specified by CLKDIV value in the FMC_BTR1 register. CLKDIV in FMC_BWTR1 is dont care. If the synchronous mode is used and CCLKEN bit is set, the synchronous memories connected to other banks than Bank 1 are clocked by the same clock (the CLKDIV value in the FMC_BTR2..4 and FMC_BWTR2..4 registers for other banks has no effect.)
func (r *registerBcr4Type) SetCclken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldCclkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldCclkenMask)
	}
}

const (
	RegisterBcr4FieldWfdisShift = 21
	RegisterBcr4FieldWfdisMask  = 0x200000
)

// GetWfdis Write FIFO Disable This bit disables the Write FIFO used by the FMC controller. Note: The WFDIS bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr4Type) GetWfdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldWfdisMask) != 0
}

// SetWfdis Write FIFO Disable This bit disables the Write FIFO used by the FMC controller. Note: The WFDIS bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr4Type) SetWfdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldWfdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldWfdisMask)
	}
}

const (
	RegisterBcr4FieldBmapShift = 24
	RegisterBcr4FieldBmapMask  = 0x3000000
)

// GetBmap FMC bank mapping These bits allows different to remap SDRAM bank2 or swap the FMC NOR/PSRAM and SDRAM banks.Refer to Table 10 for Note: The BMAP bits of the FMC_BCR2..4 registers are dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr4Type) GetBmap() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldBmapMask) >> RegisterBcr4FieldBmapShift)
}

// SetBmap FMC bank mapping These bits allows different to remap SDRAM bank2 or swap the FMC NOR/PSRAM and SDRAM banks.Refer to Table 10 for Note: The BMAP bits of the FMC_BCR2..4 registers are dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr4Type) SetBmap(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldBmapMask)|(uint32(value)<<RegisterBcr4FieldBmapShift))
}

const (
	RegisterBcr4FieldFmcenShift = 31
	RegisterBcr4FieldFmcenMask  = 0x80000000
)

// GetFmcen FMC controller Enable This bit enables/disables the FMC controller. Note: The FMCEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr4Type) GetFmcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBcr4FieldFmcenMask) != 0
}

// SetFmcen FMC controller Enable This bit enables/disables the FMC controller. Note: The FMCEN bit of the FMC_BCR2..4 registers is dont care. It is only enabled through the FMC_BCR1 register.
func (r *registerBcr4Type) SetFmcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBcr4FieldFmcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBcr4FieldFmcenMask)
	}
}

// registerBtr4Type This register contains the control information of each memory bank, used for SRAMs, PSRAM and NOR Flash memories.If the EXTMOD bit is set in the FMC_BCRx register, then this register is partitioned for write and read access, that is, 2 registers are available: one to configure read accesses (this register) and one to configure write accesses (FMC_BWTRx registers).
type registerBtr4Type uint32

const (
	RegisterBtr4FieldAddsetShift = 0
	RegisterBtr4FieldAddsetMask  = 0xf
)

// GetAddset Address setup phase duration These bits are written by software to define the duration of the address setup phase (refer to Figure81 to Figure93), used in SRAMs, ROMs and asynchronous NOR Flash: For each access mode address setup phase duration, please refer to the respective figure (refer to Figure81 to Figure93). Note: In synchronous accesses, this value is dont care. In Muxed mode or Mode D, the minimum value for ADDSET is 1.
func (r *registerBtr4Type) GetAddset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr4FieldAddsetMask) >> RegisterBtr4FieldAddsetShift)
}

// SetAddset Address setup phase duration These bits are written by software to define the duration of the address setup phase (refer to Figure81 to Figure93), used in SRAMs, ROMs and asynchronous NOR Flash: For each access mode address setup phase duration, please refer to the respective figure (refer to Figure81 to Figure93). Note: In synchronous accesses, this value is dont care. In Muxed mode or Mode D, the minimum value for ADDSET is 1.
func (r *registerBtr4Type) SetAddset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr4FieldAddsetMask)|(uint32(value)<<RegisterBtr4FieldAddsetShift))
}

const (
	RegisterBtr4FieldAddhldShift = 4
	RegisterBtr4FieldAddhldMask  = 0xf0
)

// GetAddhld Address-hold phase duration These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in mode D or multiplexed accesses: For each access mode address-hold phase duration, please refer to the respective figure (Figure81 to Figure93). Note: In synchronous accesses, this value is not used, the address hold phase is always 1 memory clock period duration.
func (r *registerBtr4Type) GetAddhld() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr4FieldAddhldMask) >> RegisterBtr4FieldAddhldShift)
}

// SetAddhld Address-hold phase duration These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in mode D or multiplexed accesses: For each access mode address-hold phase duration, please refer to the respective figure (Figure81 to Figure93). Note: In synchronous accesses, this value is not used, the address hold phase is always 1 memory clock period duration.
func (r *registerBtr4Type) SetAddhld(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr4FieldAddhldMask)|(uint32(value)<<RegisterBtr4FieldAddhldShift))
}

const (
	RegisterBtr4FieldDatastShift = 8
	RegisterBtr4FieldDatastMask  = 0xff00
)

// GetDatast Data-phase duration These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous accesses: For each memory type and access mode data-phase duration, please refer to the respective figure (Figure81 to Figure93). Example: Mode1, write access, DATAST=1: Data-phase duration= DATAST+1 = 2 KCK_FMC clock cycles. Note: In synchronous accesses, this value is dont care.
func (r *registerBtr4Type) GetDatast() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr4FieldDatastMask) >> RegisterBtr4FieldDatastShift)
}

// SetDatast Data-phase duration These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous accesses: For each memory type and access mode data-phase duration, please refer to the respective figure (Figure81 to Figure93). Example: Mode1, write access, DATAST=1: Data-phase duration= DATAST+1 = 2 KCK_FMC clock cycles. Note: In synchronous accesses, this value is dont care.
func (r *registerBtr4Type) SetDatast(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr4FieldDatastMask)|(uint32(value)<<RegisterBtr4FieldDatastShift))
}

const (
	RegisterBtr4FieldBusturnShift = 16
	RegisterBtr4FieldBusturnMask  = 0xf0000
)

// GetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write-to-read or read-to write transaction. The programmed bus turnaround delay is inserted between an asynchronous read (in muxed or mode D) or write transaction and any other asynchronous /synchronous read/write from/to a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except in muxed mode and mode D. There is a bus turnaround delay of 1 FMC clock cycle between: Two consecutive asynchronous read transfers to the same static memory bank except for modes muxed and D. An asynchronous read to an asynchronous or synchronous write to any static bank or dynamic bank except in modes muxed and D mode. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank. A synchronous write (burst or single) access and an asynchronous write or read transfer to or from static memory bank (the bank can be the same or a different one in case of a read operation. Two consecutive synchronous read operations (in Burst or Single mode) followed by any synchronous/asynchronous read or write from/to another static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write access (in Burst or Single mode) and a synchronous read from the same or a different bank. The bus turnaround delay allows to match the minimum time between consecutive transactions (tEHEL from NEx high to NEx low) and the maximum time required by the memory to free the data bus after a read access (tEHQZ): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin and (BUSTRUN + 2)KCK_FMC period &#8805; tEHQZmax if EXTMOD = 0 (BUSTRUN + 2)KCK_FMC period &#8805; max (tEHELmin, tEHQZmax) if EXTMOD =1. ...
func (r *registerBtr4Type) GetBusturn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr4FieldBusturnMask) >> RegisterBtr4FieldBusturnShift)
}

// SetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write-to-read or read-to write transaction. The programmed bus turnaround delay is inserted between an asynchronous read (in muxed or mode D) or write transaction and any other asynchronous /synchronous read/write from/to a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except in muxed mode and mode D. There is a bus turnaround delay of 1 FMC clock cycle between: Two consecutive asynchronous read transfers to the same static memory bank except for modes muxed and D. An asynchronous read to an asynchronous or synchronous write to any static bank or dynamic bank except in modes muxed and D mode. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank. A synchronous write (burst or single) access and an asynchronous write or read transfer to or from static memory bank (the bank can be the same or a different one in case of a read operation. Two consecutive synchronous read operations (in Burst or Single mode) followed by any synchronous/asynchronous read or write from/to another static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write access (in Burst or Single mode) and a synchronous read from the same or a different bank. The bus turnaround delay allows to match the minimum time between consecutive transactions (tEHEL from NEx high to NEx low) and the maximum time required by the memory to free the data bus after a read access (tEHQZ): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin and (BUSTRUN + 2)KCK_FMC period &#8805; tEHQZmax if EXTMOD = 0 (BUSTRUN + 2)KCK_FMC period &#8805; max (tEHELmin, tEHQZmax) if EXTMOD =1. ...
func (r *registerBtr4Type) SetBusturn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr4FieldBusturnMask)|(uint32(value)<<RegisterBtr4FieldBusturnShift))
}

const (
	RegisterBtr4FieldClkdivShift = 20
	RegisterBtr4FieldClkdivMask  = 0xf00000
)

// GetClkdiv Clock divide ratio (for FMC_CLK signal) These bits define the period of FMC_CLK clock output signal, expressed in number of KCK_FMC cycles: In asynchronous NOR Flash, SRAM or PSRAM accesses, this value is dont care. Note: Refer to Section20.6.5: Synchronous transactions for FMC_CLK divider ratio formula)
func (r *registerBtr4Type) GetClkdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr4FieldClkdivMask) >> RegisterBtr4FieldClkdivShift)
}

// SetClkdiv Clock divide ratio (for FMC_CLK signal) These bits define the period of FMC_CLK clock output signal, expressed in number of KCK_FMC cycles: In asynchronous NOR Flash, SRAM or PSRAM accesses, this value is dont care. Note: Refer to Section20.6.5: Synchronous transactions for FMC_CLK divider ratio formula)
func (r *registerBtr4Type) SetClkdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr4FieldClkdivMask)|(uint32(value)<<RegisterBtr4FieldClkdivShift))
}

const (
	RegisterBtr4FieldDatlatShift = 24
	RegisterBtr4FieldDatlatMask  = 0xf000000
)

// GetDatlat Data latency for synchronous memory For synchronous access with read write burst mode enabled these bits define the number of memory clock cycles
func (r *registerBtr4Type) GetDatlat() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr4FieldDatlatMask) >> RegisterBtr4FieldDatlatShift)
}

// SetDatlat Data latency for synchronous memory For synchronous access with read write burst mode enabled these bits define the number of memory clock cycles
func (r *registerBtr4Type) SetDatlat(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr4FieldDatlatMask)|(uint32(value)<<RegisterBtr4FieldDatlatShift))
}

const (
	RegisterBtr4FieldAccmodShift = 28
	RegisterBtr4FieldAccmodMask  = 0x30000000
)

// GetAccmod Access mode These bits specify the asynchronous access modes as shown in the timing diagrams. They are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBtr4Type) GetAccmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBtr4FieldAccmodMask) >> RegisterBtr4FieldAccmodShift)
}

// SetAccmod Access mode These bits specify the asynchronous access modes as shown in the timing diagrams. They are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBtr4Type) SetAccmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBtr4FieldAccmodMask)|(uint32(value)<<RegisterBtr4FieldAccmodShift))
}

// registerPcrType NAND Flash control registers
type registerPcrType uint32

const (
	RegisterPcrFieldPwaitenShift = 1
	RegisterPcrFieldPwaitenMask  = 0x2
)

// GetPwaiten Wait feature enable bit. This bit enables the Wait feature for the NAND Flash memory bank:
func (r *registerPcrType) GetPwaiten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPcrFieldPwaitenMask) != 0
}

// SetPwaiten Wait feature enable bit. This bit enables the Wait feature for the NAND Flash memory bank:
func (r *registerPcrType) SetPwaiten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPcrFieldPwaitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPcrFieldPwaitenMask)
	}
}

const (
	RegisterPcrFieldPbkenShift = 2
	RegisterPcrFieldPbkenMask  = 0x4
)

// GetPbken NAND Flash memory bank enable bit. This bit enables the memory bank. Accessing a disabled memory bank causes an ERROR on AXI bus
func (r *registerPcrType) GetPbken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPcrFieldPbkenMask) != 0
}

// SetPbken NAND Flash memory bank enable bit. This bit enables the memory bank. Accessing a disabled memory bank causes an ERROR on AXI bus
func (r *registerPcrType) SetPbken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPcrFieldPbkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPcrFieldPbkenMask)
	}
}

const (
	RegisterPcrFieldPwidShift = 4
	RegisterPcrFieldPwidMask  = 0x30
)

// GetPwid Data bus width. These bits define the external memory device width.
func (r *registerPcrType) GetPwid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPcrFieldPwidMask) >> RegisterPcrFieldPwidShift)
}

// SetPwid Data bus width. These bits define the external memory device width.
func (r *registerPcrType) SetPwid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPcrFieldPwidMask)|(uint32(value)<<RegisterPcrFieldPwidShift))
}

const (
	RegisterPcrFieldEccenShift = 6
	RegisterPcrFieldEccenMask  = 0x40
)

// GetEccen ECC computation logic enable bit
func (r *registerPcrType) GetEccen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPcrFieldEccenMask) != 0
}

// SetEccen ECC computation logic enable bit
func (r *registerPcrType) SetEccen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPcrFieldEccenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPcrFieldEccenMask)
	}
}

const (
	RegisterPcrFieldTclrShift = 9
	RegisterPcrFieldTclrMask  = 0x1e00
)

// GetTclr CLE to RE delay. These bits set time from CLE low to RE low in number of KCK_FMC clock cycles. The time is give by the following formula: t_clr = (TCLR + SET + 2) TKCK_FMC where TKCK_FMC is the KCK_FMC clock period Note: Set is MEMSET or ATTSET according to the addressed space.
func (r *registerPcrType) GetTclr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPcrFieldTclrMask) >> RegisterPcrFieldTclrShift)
}

// SetTclr CLE to RE delay. These bits set time from CLE low to RE low in number of KCK_FMC clock cycles. The time is give by the following formula: t_clr = (TCLR + SET + 2) TKCK_FMC where TKCK_FMC is the KCK_FMC clock period Note: Set is MEMSET or ATTSET according to the addressed space.
func (r *registerPcrType) SetTclr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPcrFieldTclrMask)|(uint32(value)<<RegisterPcrFieldTclrShift))
}

const (
	RegisterPcrFieldTarShift = 13
	RegisterPcrFieldTarMask  = 0x1e000
)

// GetTar ALE to RE delay. These bits set time from ALE low to RE low in number of KCK_FMC clock cycles. Time is: t_ar = (TAR + SET + 2) TKCK_FMC where TKCK_FMC is the FMC clock period Note: Set is MEMSET or ATTSET according to the addressed space.
func (r *registerPcrType) GetTar() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPcrFieldTarMask) >> RegisterPcrFieldTarShift)
}

// SetTar ALE to RE delay. These bits set time from ALE low to RE low in number of KCK_FMC clock cycles. Time is: t_ar = (TAR + SET + 2) TKCK_FMC where TKCK_FMC is the FMC clock period Note: Set is MEMSET or ATTSET according to the addressed space.
func (r *registerPcrType) SetTar(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPcrFieldTarMask)|(uint32(value)<<RegisterPcrFieldTarShift))
}

const (
	RegisterPcrFieldEccpsShift = 17
	RegisterPcrFieldEccpsMask  = 0xe0000
)

// GetEccps ECC page size. These bits define the page size for the extended ECC:
func (r *registerPcrType) GetEccps() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPcrFieldEccpsMask) >> RegisterPcrFieldEccpsShift)
}

// SetEccps ECC page size. These bits define the page size for the extended ECC:
func (r *registerPcrType) SetEccps(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPcrFieldEccpsMask)|(uint32(value)<<RegisterPcrFieldEccpsShift))
}

// registerSrType This register contains information about the FIFO status and interrupt. The FMC features a FIFO that is used when writing to memories to transfer up to 16 words of data.This is used to quickly write to the FIFO and free the AXI bus for transactions to peripherals other than the FMC, while the FMC is draining its FIFO into the memory. One of these register bits indicates the status of the FIFO, for ECC purposes.The ECC is calculated while the data are written to the memory. To read the correct ECC, the software must consequently wait until the FIFO is empty.
type registerSrType uint32

const (
	RegisterSrFieldIrsShift = 0
	RegisterSrFieldIrsMask  = 0x1
)

// GetIrs Interrupt rising edge status The flag is set by hardware and reset by software. Note: If this bit is written by software to 1 it will be set.
func (r *registerSrType) GetIrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIrsMask) != 0
}

// SetIrs Interrupt rising edge status The flag is set by hardware and reset by software. Note: If this bit is written by software to 1 it will be set.
func (r *registerSrType) SetIrs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldIrsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldIrsMask)
	}
}

const (
	RegisterSrFieldIlsShift = 1
	RegisterSrFieldIlsMask  = 0x2
)

// GetIls Interrupt high-level status The flag is set by hardware and reset by software.
func (r *registerSrType) GetIls() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIlsMask) != 0
}

// SetIls Interrupt high-level status The flag is set by hardware and reset by software.
func (r *registerSrType) SetIls(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldIlsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldIlsMask)
	}
}

const (
	RegisterSrFieldIfsShift = 2
	RegisterSrFieldIfsMask  = 0x4
)

// GetIfs Interrupt falling edge status The flag is set by hardware and reset by software. Note: If this bit is written by software to 1 it will be set.
func (r *registerSrType) GetIfs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIfsMask) != 0
}

// SetIfs Interrupt falling edge status The flag is set by hardware and reset by software. Note: If this bit is written by software to 1 it will be set.
func (r *registerSrType) SetIfs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldIfsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldIfsMask)
	}
}

const (
	RegisterSrFieldIrenShift = 3
	RegisterSrFieldIrenMask  = 0x8
)

// GetIren Interrupt rising edge detection enable bit
func (r *registerSrType) GetIren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIrenMask) != 0
}

// SetIren Interrupt rising edge detection enable bit
func (r *registerSrType) SetIren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldIrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldIrenMask)
	}
}

const (
	RegisterSrFieldIlenShift = 4
	RegisterSrFieldIlenMask  = 0x10
)

// GetIlen Interrupt high-level detection enable bit
func (r *registerSrType) GetIlen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIlenMask) != 0
}

// SetIlen Interrupt high-level detection enable bit
func (r *registerSrType) SetIlen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldIlenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldIlenMask)
	}
}

const (
	RegisterSrFieldIfenShift = 5
	RegisterSrFieldIfenMask  = 0x20
)

// GetIfen Interrupt falling edge detection enable bit
func (r *registerSrType) GetIfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIfenMask) != 0
}

// SetIfen Interrupt falling edge detection enable bit
func (r *registerSrType) SetIfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldIfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldIfenMask)
	}
}

const (
	RegisterSrFieldFemptShift = 6
	RegisterSrFieldFemptMask  = 0x40
)

// GetFempt FIFO empty. Read-only bit that provides the status of the FIFO
func (r *registerSrType) GetFempt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldFemptMask) != 0
}

// registerPmemType The FMC_PMEM read/write register contains the timing information for NAND Flash memory bank. This information is used to access either the common memory space of the NAND Flash for command, address write access and data read/write access.
type registerPmemType uint32

const (
	RegisterPmemFieldMemsetShift = 0
	RegisterPmemFieldMemsetMask  = 0xff
)

// GetMemset Common memory x setup time These bits define the number of KCK_FMC (+1) clock cycles to set up the address before the command assertion (NWE, NOE), for NAND Flash read or write access to common memory space:
func (r *registerPmemType) GetMemset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPmemFieldMemsetMask) >> RegisterPmemFieldMemsetShift)
}

// SetMemset Common memory x setup time These bits define the number of KCK_FMC (+1) clock cycles to set up the address before the command assertion (NWE, NOE), for NAND Flash read or write access to common memory space:
func (r *registerPmemType) SetMemset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPmemFieldMemsetMask)|(uint32(value)<<RegisterPmemFieldMemsetShift))
}

const (
	RegisterPmemFieldMemwaitShift = 8
	RegisterPmemFieldMemwaitMask  = 0xff00
)

// GetMemwait Common memory wait time These bits define the minimum number of KCK_FMC (+1) clock cycles to assert the command (NWE, NOE), for NAND Flash read or write access to common memory space. The duration of command assertion is extended if the wait signal (NWAIT) is active (low) at the end of the programmed value of KCK_FMC:
func (r *registerPmemType) GetMemwait() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPmemFieldMemwaitMask) >> RegisterPmemFieldMemwaitShift)
}

// SetMemwait Common memory wait time These bits define the minimum number of KCK_FMC (+1) clock cycles to assert the command (NWE, NOE), for NAND Flash read or write access to common memory space. The duration of command assertion is extended if the wait signal (NWAIT) is active (low) at the end of the programmed value of KCK_FMC:
func (r *registerPmemType) SetMemwait(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPmemFieldMemwaitMask)|(uint32(value)<<RegisterPmemFieldMemwaitShift))
}

const (
	RegisterPmemFieldMemholdShift = 16
	RegisterPmemFieldMemholdMask  = 0xff0000
)

// GetMemhold Common memory hold time These bits define the number of KCK_FMC clock cycles for write accesses and KCK_FMC+1 clock cycles for read accesses during which the address is held (and data for write accesses) after the command is de-asserted (NWE, NOE), for NAND Flash read or write access to common memory space:
func (r *registerPmemType) GetMemhold() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPmemFieldMemholdMask) >> RegisterPmemFieldMemholdShift)
}

// SetMemhold Common memory hold time These bits define the number of KCK_FMC clock cycles for write accesses and KCK_FMC+1 clock cycles for read accesses during which the address is held (and data for write accesses) after the command is de-asserted (NWE, NOE), for NAND Flash read or write access to common memory space:
func (r *registerPmemType) SetMemhold(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPmemFieldMemholdMask)|(uint32(value)<<RegisterPmemFieldMemholdShift))
}

const (
	RegisterPmemFieldMemhizShift = 24
	RegisterPmemFieldMemhizMask  = 0xff000000
)

// GetMemhiz Common memory x data bus Hi-Z time These bits define the number of KCK_FMC clock cycles during which the data bus is kept Hi-Z after the start of a NAND Flash write access to common memory space. This is only valid for write transactions:
func (r *registerPmemType) GetMemhiz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPmemFieldMemhizMask) >> RegisterPmemFieldMemhizShift)
}

// SetMemhiz Common memory x data bus Hi-Z time These bits define the number of KCK_FMC clock cycles during which the data bus is kept Hi-Z after the start of a NAND Flash write access to common memory space. This is only valid for write transactions:
func (r *registerPmemType) SetMemhiz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPmemFieldMemhizMask)|(uint32(value)<<RegisterPmemFieldMemhizShift))
}

// registerPattType The FMC_PATT read/write register contains the timing information for NAND Flash memory bank. It is used for 8-bit accesses to the attribute memory space of the NAND Flash for the last address write access if the timing must differ from that of previous accesses (for Ready/Busy management, refer to Section20.8.5: NAND Flash prewait feature).
type registerPattType uint32

const (
	RegisterPattFieldAttsetShift = 0
	RegisterPattFieldAttsetMask  = 0xff
)

// GetAttset Attribute memory setup time These bits define the number of KCK_FMC (+1) clock cycles to set up address before the command assertion (NWE, NOE), for NAND Flash read or write access to attribute memory space:
func (r *registerPattType) GetAttset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPattFieldAttsetMask) >> RegisterPattFieldAttsetShift)
}

// SetAttset Attribute memory setup time These bits define the number of KCK_FMC (+1) clock cycles to set up address before the command assertion (NWE, NOE), for NAND Flash read or write access to attribute memory space:
func (r *registerPattType) SetAttset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPattFieldAttsetMask)|(uint32(value)<<RegisterPattFieldAttsetShift))
}

const (
	RegisterPattFieldAttwaitShift = 8
	RegisterPattFieldAttwaitMask  = 0xff00
)

// GetAttwait Attribute memory wait time These bits define the minimum number of x KCK_FMC (+1) clock cycles to assert the command (NWE, NOE), for NAND Flash read or write access to attribute memory space. The duration for command assertion is extended if the wait signal (NWAIT) is active (low) at the end of the programmed value of KCK_FMC:
func (r *registerPattType) GetAttwait() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPattFieldAttwaitMask) >> RegisterPattFieldAttwaitShift)
}

// SetAttwait Attribute memory wait time These bits define the minimum number of x KCK_FMC (+1) clock cycles to assert the command (NWE, NOE), for NAND Flash read or write access to attribute memory space. The duration for command assertion is extended if the wait signal (NWAIT) is active (low) at the end of the programmed value of KCK_FMC:
func (r *registerPattType) SetAttwait(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPattFieldAttwaitMask)|(uint32(value)<<RegisterPattFieldAttwaitShift))
}

const (
	RegisterPattFieldAttholdShift = 16
	RegisterPattFieldAttholdMask  = 0xff0000
)

// GetAtthold Attribute memory hold time These bits define the number of KCK_FMC clock cycles during which the address is held (and data for write access) after the command de-assertion (NWE, NOE), for NAND Flash read or write access to attribute memory space:
func (r *registerPattType) GetAtthold() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPattFieldAttholdMask) >> RegisterPattFieldAttholdShift)
}

// SetAtthold Attribute memory hold time These bits define the number of KCK_FMC clock cycles during which the address is held (and data for write access) after the command de-assertion (NWE, NOE), for NAND Flash read or write access to attribute memory space:
func (r *registerPattType) SetAtthold(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPattFieldAttholdMask)|(uint32(value)<<RegisterPattFieldAttholdShift))
}

const (
	RegisterPattFieldAtthizShift = 24
	RegisterPattFieldAtthizMask  = 0xff000000
)

// GetAtthiz Attribute memory data bus Hi-Z time These bits define the number of KCK_FMC clock cycles during which the data bus is kept in Hi-Z after the start of a NAND Flash write access to attribute memory space on socket. Only valid for writ transaction:
func (r *registerPattType) GetAtthiz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPattFieldAtthizMask) >> RegisterPattFieldAtthizShift)
}

// SetAtthiz Attribute memory data bus Hi-Z time These bits define the number of KCK_FMC clock cycles during which the data bus is kept in Hi-Z after the start of a NAND Flash write access to attribute memory space on socket. Only valid for writ transaction:
func (r *registerPattType) SetAtthiz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPattFieldAtthizMask)|(uint32(value)<<RegisterPattFieldAtthizShift))
}

// registerEccrType This register contain the current error correction code value computed by the ECC computation modules of the FMC NAND controller. When the CPU reads/writes the data from a NAND Flash memory page at the correct address (refer to Section20.8.6: Computation of the error correction code (ECC) in NAND Flash memory), the data read/written from/to the NAND Flash memory are processed automatically by the ECC computation module. When X bytes have been read (according to the ECCPS field in the FMC_PCR registers), the CPU must read the computed ECC value from the FMC_ECC registers. It then verifies if these computed parity data are the same as the parity value recorded in the spare area, to determine whether a page is valid, and, to correct it otherwise. The FMC_ECCR register should be cleared after being read by setting the ECCEN bit to 0. To compute a new data block, the ECCEN bit must be set to 1.
type registerEccrType uint32

const (
	RegisterEccrFieldEccShift = 0
	RegisterEccrFieldEccMask  = 0xffffffff
)

// GetEcc ECC result This field contains the value computed by the ECC computation logic. Table167 describes the contents of these bit fields.
func (r *registerEccrType) GetEcc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterEccrFieldEccMask) >> RegisterEccrFieldEccShift)
}

// SetEcc ECC result This field contains the value computed by the ECC computation logic. Table167 describes the contents of these bit fields.
func (r *registerEccrType) SetEcc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEccrFieldEccMask)|(uint32(value)<<RegisterEccrFieldEccShift))
}

// registerBwtr1Type This register contains the control information of each memory bank. It is used for SRAMs, PSRAMs and NOR Flash memories. When the EXTMOD bit is set in the FMC_BCRx register, then this register is active for write access.
type registerBwtr1Type uint32

const (
	RegisterBwtr1FieldAddsetShift = 0
	RegisterBwtr1FieldAddsetMask  = 0xf
)

// GetAddset Address setup phase duration. These bits are written by software to define the duration of the address setup phase in KCK_FMC cycles (refer to Figure81 to Figure93), used in asynchronous accesses: ... Note: In synchronous accesses, this value is not used, the address setup phase is always 1 Flash clock period duration. In muxed mode, the minimum ADDSET value is 1.
func (r *registerBwtr1Type) GetAddset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr1FieldAddsetMask) >> RegisterBwtr1FieldAddsetShift)
}

// SetAddset Address setup phase duration. These bits are written by software to define the duration of the address setup phase in KCK_FMC cycles (refer to Figure81 to Figure93), used in asynchronous accesses: ... Note: In synchronous accesses, this value is not used, the address setup phase is always 1 Flash clock period duration. In muxed mode, the minimum ADDSET value is 1.
func (r *registerBwtr1Type) SetAddset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr1FieldAddsetMask)|(uint32(value)<<RegisterBwtr1FieldAddsetShift))
}

const (
	RegisterBwtr1FieldAddhldShift = 4
	RegisterBwtr1FieldAddhldMask  = 0xf0
)

// GetAddhld Address-hold phase duration. These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in asynchronous multiplexed accesses: ... Note: In synchronous NOR Flash accesses, this value is not used, the address hold phase is always 1 Flash clock period duration.
func (r *registerBwtr1Type) GetAddhld() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr1FieldAddhldMask) >> RegisterBwtr1FieldAddhldShift)
}

// SetAddhld Address-hold phase duration. These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in asynchronous multiplexed accesses: ... Note: In synchronous NOR Flash accesses, this value is not used, the address hold phase is always 1 Flash clock period duration.
func (r *registerBwtr1Type) SetAddhld(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr1FieldAddhldMask)|(uint32(value)<<RegisterBwtr1FieldAddhldShift))
}

const (
	RegisterBwtr1FieldDatastShift = 8
	RegisterBwtr1FieldDatastMask  = 0xff00
)

// GetDatast Data-phase duration. These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous SRAM, PSRAM and NOR Flash memory accesses:
func (r *registerBwtr1Type) GetDatast() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr1FieldDatastMask) >> RegisterBwtr1FieldDatastShift)
}

// SetDatast Data-phase duration. These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous SRAM, PSRAM and NOR Flash memory accesses:
func (r *registerBwtr1Type) SetDatast(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr1FieldDatastMask)|(uint32(value)<<RegisterBwtr1FieldDatastShift))
}

const (
	RegisterBwtr1FieldBusturnShift = 16
	RegisterBwtr1FieldBusturnMask  = 0xf0000
)

// GetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write transaction to match the minimum time between consecutive transactions (tEHEL from ENx high to ENx low): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin. The programmed bus turnaround delay is inserted between a an asynchronous write transfer and any other asynchronous /synchronous read or write transfer to or from a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except for muxed mode and mode D. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank A synchronous write transfer ((in Burst or Single mode) and an asynchronous write or read transfer to or from static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write transfer (in Burst or Single mode) and a synchronous read from the same or a different bank. ...
func (r *registerBwtr1Type) GetBusturn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr1FieldBusturnMask) >> RegisterBwtr1FieldBusturnShift)
}

// SetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write transaction to match the minimum time between consecutive transactions (tEHEL from ENx high to ENx low): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin. The programmed bus turnaround delay is inserted between a an asynchronous write transfer and any other asynchronous /synchronous read or write transfer to or from a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except for muxed mode and mode D. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank A synchronous write transfer ((in Burst or Single mode) and an asynchronous write or read transfer to or from static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write transfer (in Burst or Single mode) and a synchronous read from the same or a different bank. ...
func (r *registerBwtr1Type) SetBusturn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr1FieldBusturnMask)|(uint32(value)<<RegisterBwtr1FieldBusturnShift))
}

const (
	RegisterBwtr1FieldAccmodShift = 28
	RegisterBwtr1FieldAccmodMask  = 0x30000000
)

// GetAccmod Access mode. These bits specify the asynchronous access modes as shown in the next timing diagrams.These bits are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBwtr1Type) GetAccmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr1FieldAccmodMask) >> RegisterBwtr1FieldAccmodShift)
}

// SetAccmod Access mode. These bits specify the asynchronous access modes as shown in the next timing diagrams.These bits are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBwtr1Type) SetAccmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr1FieldAccmodMask)|(uint32(value)<<RegisterBwtr1FieldAccmodShift))
}

// registerBwtr2Type This register contains the control information of each memory bank. It is used for SRAMs, PSRAMs and NOR Flash memories. When the EXTMOD bit is set in the FMC_BCRx register, then this register is active for write access.
type registerBwtr2Type uint32

const (
	RegisterBwtr2FieldAddsetShift = 0
	RegisterBwtr2FieldAddsetMask  = 0xf
)

// GetAddset Address setup phase duration. These bits are written by software to define the duration of the address setup phase in KCK_FMC cycles (refer to Figure81 to Figure93), used in asynchronous accesses: ... Note: In synchronous accesses, this value is not used, the address setup phase is always 1 Flash clock period duration. In muxed mode, the minimum ADDSET value is 1.
func (r *registerBwtr2Type) GetAddset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr2FieldAddsetMask) >> RegisterBwtr2FieldAddsetShift)
}

// SetAddset Address setup phase duration. These bits are written by software to define the duration of the address setup phase in KCK_FMC cycles (refer to Figure81 to Figure93), used in asynchronous accesses: ... Note: In synchronous accesses, this value is not used, the address setup phase is always 1 Flash clock period duration. In muxed mode, the minimum ADDSET value is 1.
func (r *registerBwtr2Type) SetAddset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr2FieldAddsetMask)|(uint32(value)<<RegisterBwtr2FieldAddsetShift))
}

const (
	RegisterBwtr2FieldAddhldShift = 4
	RegisterBwtr2FieldAddhldMask  = 0xf0
)

// GetAddhld Address-hold phase duration. These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in asynchronous multiplexed accesses: ... Note: In synchronous NOR Flash accesses, this value is not used, the address hold phase is always 1 Flash clock period duration.
func (r *registerBwtr2Type) GetAddhld() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr2FieldAddhldMask) >> RegisterBwtr2FieldAddhldShift)
}

// SetAddhld Address-hold phase duration. These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in asynchronous multiplexed accesses: ... Note: In synchronous NOR Flash accesses, this value is not used, the address hold phase is always 1 Flash clock period duration.
func (r *registerBwtr2Type) SetAddhld(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr2FieldAddhldMask)|(uint32(value)<<RegisterBwtr2FieldAddhldShift))
}

const (
	RegisterBwtr2FieldDatastShift = 8
	RegisterBwtr2FieldDatastMask  = 0xff00
)

// GetDatast Data-phase duration. These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous SRAM, PSRAM and NOR Flash memory accesses:
func (r *registerBwtr2Type) GetDatast() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr2FieldDatastMask) >> RegisterBwtr2FieldDatastShift)
}

// SetDatast Data-phase duration. These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous SRAM, PSRAM and NOR Flash memory accesses:
func (r *registerBwtr2Type) SetDatast(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr2FieldDatastMask)|(uint32(value)<<RegisterBwtr2FieldDatastShift))
}

const (
	RegisterBwtr2FieldBusturnShift = 16
	RegisterBwtr2FieldBusturnMask  = 0xf0000
)

// GetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write transaction to match the minimum time between consecutive transactions (tEHEL from ENx high to ENx low): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin. The programmed bus turnaround delay is inserted between a an asynchronous write transfer and any other asynchronous /synchronous read or write transfer to or from a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except for muxed mode and mode D. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank A synchronous write transfer ((in Burst or Single mode) and an asynchronous write or read transfer to or from static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write transfer (in Burst or Single mode) and a synchronous read from the same or a different bank. ...
func (r *registerBwtr2Type) GetBusturn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr2FieldBusturnMask) >> RegisterBwtr2FieldBusturnShift)
}

// SetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write transaction to match the minimum time between consecutive transactions (tEHEL from ENx high to ENx low): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin. The programmed bus turnaround delay is inserted between a an asynchronous write transfer and any other asynchronous /synchronous read or write transfer to or from a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except for muxed mode and mode D. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank A synchronous write transfer ((in Burst or Single mode) and an asynchronous write or read transfer to or from static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write transfer (in Burst or Single mode) and a synchronous read from the same or a different bank. ...
func (r *registerBwtr2Type) SetBusturn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr2FieldBusturnMask)|(uint32(value)<<RegisterBwtr2FieldBusturnShift))
}

const (
	RegisterBwtr2FieldAccmodShift = 28
	RegisterBwtr2FieldAccmodMask  = 0x30000000
)

// GetAccmod Access mode. These bits specify the asynchronous access modes as shown in the next timing diagrams.These bits are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBwtr2Type) GetAccmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr2FieldAccmodMask) >> RegisterBwtr2FieldAccmodShift)
}

// SetAccmod Access mode. These bits specify the asynchronous access modes as shown in the next timing diagrams.These bits are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBwtr2Type) SetAccmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr2FieldAccmodMask)|(uint32(value)<<RegisterBwtr2FieldAccmodShift))
}

// registerBwtr3Type This register contains the control information of each memory bank. It is used for SRAMs, PSRAMs and NOR Flash memories. When the EXTMOD bit is set in the FMC_BCRx register, then this register is active for write access.
type registerBwtr3Type uint32

const (
	RegisterBwtr3FieldAddsetShift = 0
	RegisterBwtr3FieldAddsetMask  = 0xf
)

// GetAddset Address setup phase duration. These bits are written by software to define the duration of the address setup phase in KCK_FMC cycles (refer to Figure81 to Figure93), used in asynchronous accesses: ... Note: In synchronous accesses, this value is not used, the address setup phase is always 1 Flash clock period duration. In muxed mode, the minimum ADDSET value is 1.
func (r *registerBwtr3Type) GetAddset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr3FieldAddsetMask) >> RegisterBwtr3FieldAddsetShift)
}

// SetAddset Address setup phase duration. These bits are written by software to define the duration of the address setup phase in KCK_FMC cycles (refer to Figure81 to Figure93), used in asynchronous accesses: ... Note: In synchronous accesses, this value is not used, the address setup phase is always 1 Flash clock period duration. In muxed mode, the minimum ADDSET value is 1.
func (r *registerBwtr3Type) SetAddset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr3FieldAddsetMask)|(uint32(value)<<RegisterBwtr3FieldAddsetShift))
}

const (
	RegisterBwtr3FieldAddhldShift = 4
	RegisterBwtr3FieldAddhldMask  = 0xf0
)

// GetAddhld Address-hold phase duration. These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in asynchronous multiplexed accesses: ... Note: In synchronous NOR Flash accesses, this value is not used, the address hold phase is always 1 Flash clock period duration.
func (r *registerBwtr3Type) GetAddhld() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr3FieldAddhldMask) >> RegisterBwtr3FieldAddhldShift)
}

// SetAddhld Address-hold phase duration. These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in asynchronous multiplexed accesses: ... Note: In synchronous NOR Flash accesses, this value is not used, the address hold phase is always 1 Flash clock period duration.
func (r *registerBwtr3Type) SetAddhld(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr3FieldAddhldMask)|(uint32(value)<<RegisterBwtr3FieldAddhldShift))
}

const (
	RegisterBwtr3FieldDatastShift = 8
	RegisterBwtr3FieldDatastMask  = 0xff00
)

// GetDatast Data-phase duration. These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous SRAM, PSRAM and NOR Flash memory accesses:
func (r *registerBwtr3Type) GetDatast() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr3FieldDatastMask) >> RegisterBwtr3FieldDatastShift)
}

// SetDatast Data-phase duration. These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous SRAM, PSRAM and NOR Flash memory accesses:
func (r *registerBwtr3Type) SetDatast(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr3FieldDatastMask)|(uint32(value)<<RegisterBwtr3FieldDatastShift))
}

const (
	RegisterBwtr3FieldBusturnShift = 16
	RegisterBwtr3FieldBusturnMask  = 0xf0000
)

// GetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write transaction to match the minimum time between consecutive transactions (tEHEL from ENx high to ENx low): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin. The programmed bus turnaround delay is inserted between a an asynchronous write transfer and any other asynchronous /synchronous read or write transfer to or from a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except for muxed mode and mode D. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank A synchronous write transfer ((in Burst or Single mode) and an asynchronous write or read transfer to or from static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write transfer (in Burst or Single mode) and a synchronous read from the same or a different bank. ...
func (r *registerBwtr3Type) GetBusturn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr3FieldBusturnMask) >> RegisterBwtr3FieldBusturnShift)
}

// SetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write transaction to match the minimum time between consecutive transactions (tEHEL from ENx high to ENx low): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin. The programmed bus turnaround delay is inserted between a an asynchronous write transfer and any other asynchronous /synchronous read or write transfer to or from a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except for muxed mode and mode D. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank A synchronous write transfer ((in Burst or Single mode) and an asynchronous write or read transfer to or from static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write transfer (in Burst or Single mode) and a synchronous read from the same or a different bank. ...
func (r *registerBwtr3Type) SetBusturn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr3FieldBusturnMask)|(uint32(value)<<RegisterBwtr3FieldBusturnShift))
}

const (
	RegisterBwtr3FieldAccmodShift = 28
	RegisterBwtr3FieldAccmodMask  = 0x30000000
)

// GetAccmod Access mode. These bits specify the asynchronous access modes as shown in the next timing diagrams.These bits are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBwtr3Type) GetAccmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr3FieldAccmodMask) >> RegisterBwtr3FieldAccmodShift)
}

// SetAccmod Access mode. These bits specify the asynchronous access modes as shown in the next timing diagrams.These bits are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBwtr3Type) SetAccmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr3FieldAccmodMask)|(uint32(value)<<RegisterBwtr3FieldAccmodShift))
}

// registerBwtr4Type This register contains the control information of each memory bank. It is used for SRAMs, PSRAMs and NOR Flash memories. When the EXTMOD bit is set in the FMC_BCRx register, then this register is active for write access.
type registerBwtr4Type uint32

const (
	RegisterBwtr4FieldAddsetShift = 0
	RegisterBwtr4FieldAddsetMask  = 0xf
)

// GetAddset Address setup phase duration. These bits are written by software to define the duration of the address setup phase in KCK_FMC cycles (refer to Figure81 to Figure93), used in asynchronous accesses: ... Note: In synchronous accesses, this value is not used, the address setup phase is always 1 Flash clock period duration. In muxed mode, the minimum ADDSET value is 1.
func (r *registerBwtr4Type) GetAddset() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr4FieldAddsetMask) >> RegisterBwtr4FieldAddsetShift)
}

// SetAddset Address setup phase duration. These bits are written by software to define the duration of the address setup phase in KCK_FMC cycles (refer to Figure81 to Figure93), used in asynchronous accesses: ... Note: In synchronous accesses, this value is not used, the address setup phase is always 1 Flash clock period duration. In muxed mode, the minimum ADDSET value is 1.
func (r *registerBwtr4Type) SetAddset(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr4FieldAddsetMask)|(uint32(value)<<RegisterBwtr4FieldAddsetShift))
}

const (
	RegisterBwtr4FieldAddhldShift = 4
	RegisterBwtr4FieldAddhldMask  = 0xf0
)

// GetAddhld Address-hold phase duration. These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in asynchronous multiplexed accesses: ... Note: In synchronous NOR Flash accesses, this value is not used, the address hold phase is always 1 Flash clock period duration.
func (r *registerBwtr4Type) GetAddhld() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr4FieldAddhldMask) >> RegisterBwtr4FieldAddhldShift)
}

// SetAddhld Address-hold phase duration. These bits are written by software to define the duration of the address hold phase (refer to Figure81 to Figure93), used in asynchronous multiplexed accesses: ... Note: In synchronous NOR Flash accesses, this value is not used, the address hold phase is always 1 Flash clock period duration.
func (r *registerBwtr4Type) SetAddhld(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr4FieldAddhldMask)|(uint32(value)<<RegisterBwtr4FieldAddhldShift))
}

const (
	RegisterBwtr4FieldDatastShift = 8
	RegisterBwtr4FieldDatastMask  = 0xff00
)

// GetDatast Data-phase duration. These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous SRAM, PSRAM and NOR Flash memory accesses:
func (r *registerBwtr4Type) GetDatast() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr4FieldDatastMask) >> RegisterBwtr4FieldDatastShift)
}

// SetDatast Data-phase duration. These bits are written by software to define the duration of the data phase (refer to Figure81 to Figure93), used in asynchronous SRAM, PSRAM and NOR Flash memory accesses:
func (r *registerBwtr4Type) SetDatast(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr4FieldDatastMask)|(uint32(value)<<RegisterBwtr4FieldDatastShift))
}

const (
	RegisterBwtr4FieldBusturnShift = 16
	RegisterBwtr4FieldBusturnMask  = 0xf0000
)

// GetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write transaction to match the minimum time between consecutive transactions (tEHEL from ENx high to ENx low): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin. The programmed bus turnaround delay is inserted between a an asynchronous write transfer and any other asynchronous /synchronous read or write transfer to or from a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except for muxed mode and mode D. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank A synchronous write transfer ((in Burst or Single mode) and an asynchronous write or read transfer to or from static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write transfer (in Burst or Single mode) and a synchronous read from the same or a different bank. ...
func (r *registerBwtr4Type) GetBusturn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr4FieldBusturnMask) >> RegisterBwtr4FieldBusturnShift)
}

// SetBusturn Bus turnaround phase duration These bits are written by software to add a delay at the end of a write transaction to match the minimum time between consecutive transactions (tEHEL from ENx high to ENx low): (BUSTRUN + 1) KCK_FMC period &#8805; tEHELmin. The programmed bus turnaround delay is inserted between a an asynchronous write transfer and any other asynchronous /synchronous read or write transfer to or from a static bank. If a read operation is performed, the bank can be the same or a different one, whereas it must be different in case of write operation to the bank, except in muxed mode or mode D. In some cases, whatever the programmed BUSTRUN values, the bus turnaround delay is fixed as follows: The bus turnaround delay is not inserted between two consecutive asynchronous write transfers to the same static memory bank except for muxed mode and mode D. There is a bus turnaround delay of 2 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to the same bank A synchronous write transfer ((in Burst or Single mode) and an asynchronous write or read transfer to or from static memory bank. There is a bus turnaround delay of 3 FMC clock cycle between: Two consecutive synchronous write operations (in Burst or Single mode) to different static banks. A synchronous write transfer (in Burst or Single mode) and a synchronous read from the same or a different bank. ...
func (r *registerBwtr4Type) SetBusturn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr4FieldBusturnMask)|(uint32(value)<<RegisterBwtr4FieldBusturnShift))
}

const (
	RegisterBwtr4FieldAccmodShift = 28
	RegisterBwtr4FieldAccmodMask  = 0x30000000
)

// GetAccmod Access mode. These bits specify the asynchronous access modes as shown in the next timing diagrams.These bits are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBwtr4Type) GetAccmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBwtr4FieldAccmodMask) >> RegisterBwtr4FieldAccmodShift)
}

// SetAccmod Access mode. These bits specify the asynchronous access modes as shown in the next timing diagrams.These bits are taken into account only when the EXTMOD bit in the FMC_BCRx register is 1.
func (r *registerBwtr4Type) SetAccmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBwtr4FieldAccmodMask)|(uint32(value)<<RegisterBwtr4FieldAccmodShift))
}

// registerSdcr1Type This register contains the control parameters for each SDRAM memory bank
type registerSdcr1Type uint32

const (
	RegisterSdcr1FieldNcShift = 0
	RegisterSdcr1FieldNcMask  = 0x3
)

// GetNc Number of column address bits These bits define the number of bits of a column address.
func (r *registerSdcr1Type) GetNc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr1FieldNcMask) >> RegisterSdcr1FieldNcShift)
}

// SetNc Number of column address bits These bits define the number of bits of a column address.
func (r *registerSdcr1Type) SetNc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr1FieldNcMask)|(uint32(value)<<RegisterSdcr1FieldNcShift))
}

const (
	RegisterSdcr1FieldNrShift = 2
	RegisterSdcr1FieldNrMask  = 0xc
)

// GetNr Number of row address bits These bits define the number of bits of a row address.
func (r *registerSdcr1Type) GetNr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr1FieldNrMask) >> RegisterSdcr1FieldNrShift)
}

// SetNr Number of row address bits These bits define the number of bits of a row address.
func (r *registerSdcr1Type) SetNr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr1FieldNrMask)|(uint32(value)<<RegisterSdcr1FieldNrShift))
}

const (
	RegisterSdcr1FieldMwidShift = 4
	RegisterSdcr1FieldMwidMask  = 0x30
)

// GetMwid Memory data bus width. These bits define the memory device width.
func (r *registerSdcr1Type) GetMwid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr1FieldMwidMask) >> RegisterSdcr1FieldMwidShift)
}

// SetMwid Memory data bus width. These bits define the memory device width.
func (r *registerSdcr1Type) SetMwid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr1FieldMwidMask)|(uint32(value)<<RegisterSdcr1FieldMwidShift))
}

const (
	RegisterSdcr1FieldNbShift = 6
	RegisterSdcr1FieldNbMask  = 0x40
)

// GetNb Number of internal banks This bit sets the number of internal banks.
func (r *registerSdcr1Type) GetNb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSdcr1FieldNbMask) != 0
}

// SetNb Number of internal banks This bit sets the number of internal banks.
func (r *registerSdcr1Type) SetNb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSdcr1FieldNbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSdcr1FieldNbMask)
	}
}

const (
	RegisterSdcr1FieldCasShift = 7
	RegisterSdcr1FieldCasMask  = 0x180
)

// GetCas CAS Latency This bits sets the SDRAM CAS latency in number of memory clock cycles
func (r *registerSdcr1Type) GetCas() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr1FieldCasMask) >> RegisterSdcr1FieldCasShift)
}

// SetCas CAS Latency This bits sets the SDRAM CAS latency in number of memory clock cycles
func (r *registerSdcr1Type) SetCas(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr1FieldCasMask)|(uint32(value)<<RegisterSdcr1FieldCasShift))
}

const (
	RegisterSdcr1FieldWpShift = 9
	RegisterSdcr1FieldWpMask  = 0x200
)

// GetWp Write protection This bit enables write mode access to the SDRAM bank.
func (r *registerSdcr1Type) GetWp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSdcr1FieldWpMask) != 0
}

// SetWp Write protection This bit enables write mode access to the SDRAM bank.
func (r *registerSdcr1Type) SetWp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSdcr1FieldWpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSdcr1FieldWpMask)
	}
}

const (
	RegisterSdcr1FieldSdclkShift = 10
	RegisterSdcr1FieldSdclkMask  = 0xc00
)

// GetSdclk SDRAM clock configuration These bits define the SDRAM clock period for both SDRAM banks and allow disabling the clock before changing the frequency. In this case the SDRAM must be re-initialized. Note: The corresponding bits in the FMC_SDCR2 register is read only.
func (r *registerSdcr1Type) GetSdclk() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr1FieldSdclkMask) >> RegisterSdcr1FieldSdclkShift)
}

// SetSdclk SDRAM clock configuration These bits define the SDRAM clock period for both SDRAM banks and allow disabling the clock before changing the frequency. In this case the SDRAM must be re-initialized. Note: The corresponding bits in the FMC_SDCR2 register is read only.
func (r *registerSdcr1Type) SetSdclk(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr1FieldSdclkMask)|(uint32(value)<<RegisterSdcr1FieldSdclkShift))
}

const (
	RegisterSdcr1FieldRburstShift = 12
	RegisterSdcr1FieldRburstMask  = 0x1000
)

// GetRburst Burst read This bit enables burst read mode. The SDRAM controller anticipates the next read commands during the CAS latency and stores data in the Read FIFO. Note: The corresponding bit in the FMC_SDCR2 register is read only.
func (r *registerSdcr1Type) GetRburst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSdcr1FieldRburstMask) != 0
}

// SetRburst Burst read This bit enables burst read mode. The SDRAM controller anticipates the next read commands during the CAS latency and stores data in the Read FIFO. Note: The corresponding bit in the FMC_SDCR2 register is read only.
func (r *registerSdcr1Type) SetRburst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSdcr1FieldRburstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSdcr1FieldRburstMask)
	}
}

const (
	RegisterSdcr1FieldRpipeShift = 13
	RegisterSdcr1FieldRpipeMask  = 0x6000
)

// GetRpipe Read pipe These bits define the delay, in KCK_FMC clock cycles, for reading data after CAS latency. Note: The corresponding bits in the FMC_SDCR2 register is read only.
func (r *registerSdcr1Type) GetRpipe() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr1FieldRpipeMask) >> RegisterSdcr1FieldRpipeShift)
}

// SetRpipe Read pipe These bits define the delay, in KCK_FMC clock cycles, for reading data after CAS latency. Note: The corresponding bits in the FMC_SDCR2 register is read only.
func (r *registerSdcr1Type) SetRpipe(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr1FieldRpipeMask)|(uint32(value)<<RegisterSdcr1FieldRpipeShift))
}

// registerSdcr2Type This register contains the control parameters for each SDRAM memory bank
type registerSdcr2Type uint32

const (
	RegisterSdcr2FieldNcShift = 0
	RegisterSdcr2FieldNcMask  = 0x3
)

// GetNc Number of column address bits These bits define the number of bits of a column address.
func (r *registerSdcr2Type) GetNc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr2FieldNcMask) >> RegisterSdcr2FieldNcShift)
}

// SetNc Number of column address bits These bits define the number of bits of a column address.
func (r *registerSdcr2Type) SetNc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr2FieldNcMask)|(uint32(value)<<RegisterSdcr2FieldNcShift))
}

const (
	RegisterSdcr2FieldNrShift = 2
	RegisterSdcr2FieldNrMask  = 0xc
)

// GetNr Number of row address bits These bits define the number of bits of a row address.
func (r *registerSdcr2Type) GetNr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr2FieldNrMask) >> RegisterSdcr2FieldNrShift)
}

// SetNr Number of row address bits These bits define the number of bits of a row address.
func (r *registerSdcr2Type) SetNr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr2FieldNrMask)|(uint32(value)<<RegisterSdcr2FieldNrShift))
}

const (
	RegisterSdcr2FieldMwidShift = 4
	RegisterSdcr2FieldMwidMask  = 0x30
)

// GetMwid Memory data bus width. These bits define the memory device width.
func (r *registerSdcr2Type) GetMwid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr2FieldMwidMask) >> RegisterSdcr2FieldMwidShift)
}

// SetMwid Memory data bus width. These bits define the memory device width.
func (r *registerSdcr2Type) SetMwid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr2FieldMwidMask)|(uint32(value)<<RegisterSdcr2FieldMwidShift))
}

const (
	RegisterSdcr2FieldNbShift = 6
	RegisterSdcr2FieldNbMask  = 0x40
)

// GetNb Number of internal banks This bit sets the number of internal banks.
func (r *registerSdcr2Type) GetNb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSdcr2FieldNbMask) != 0
}

// SetNb Number of internal banks This bit sets the number of internal banks.
func (r *registerSdcr2Type) SetNb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSdcr2FieldNbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSdcr2FieldNbMask)
	}
}

const (
	RegisterSdcr2FieldCasShift = 7
	RegisterSdcr2FieldCasMask  = 0x180
)

// GetCas CAS Latency This bits sets the SDRAM CAS latency in number of memory clock cycles
func (r *registerSdcr2Type) GetCas() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr2FieldCasMask) >> RegisterSdcr2FieldCasShift)
}

// SetCas CAS Latency This bits sets the SDRAM CAS latency in number of memory clock cycles
func (r *registerSdcr2Type) SetCas(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr2FieldCasMask)|(uint32(value)<<RegisterSdcr2FieldCasShift))
}

const (
	RegisterSdcr2FieldWpShift = 9
	RegisterSdcr2FieldWpMask  = 0x200
)

// GetWp Write protection This bit enables write mode access to the SDRAM bank.
func (r *registerSdcr2Type) GetWp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSdcr2FieldWpMask) != 0
}

// SetWp Write protection This bit enables write mode access to the SDRAM bank.
func (r *registerSdcr2Type) SetWp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSdcr2FieldWpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSdcr2FieldWpMask)
	}
}

const (
	RegisterSdcr2FieldSdclkShift = 10
	RegisterSdcr2FieldSdclkMask  = 0xc00
)

// GetSdclk SDRAM clock configuration These bits define the SDRAM clock period for both SDRAM banks and allow disabling the clock before changing the frequency. In this case the SDRAM must be re-initialized. Note: The corresponding bits in the FMC_SDCR2 register is read only.
func (r *registerSdcr2Type) GetSdclk() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr2FieldSdclkMask) >> RegisterSdcr2FieldSdclkShift)
}

// SetSdclk SDRAM clock configuration These bits define the SDRAM clock period for both SDRAM banks and allow disabling the clock before changing the frequency. In this case the SDRAM must be re-initialized. Note: The corresponding bits in the FMC_SDCR2 register is read only.
func (r *registerSdcr2Type) SetSdclk(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr2FieldSdclkMask)|(uint32(value)<<RegisterSdcr2FieldSdclkShift))
}

const (
	RegisterSdcr2FieldRburstShift = 12
	RegisterSdcr2FieldRburstMask  = 0x1000
)

// GetRburst Burst read This bit enables burst read mode. The SDRAM controller anticipates the next read commands during the CAS latency and stores data in the Read FIFO. Note: The corresponding bit in the FMC_SDCR2 register is read only.
func (r *registerSdcr2Type) GetRburst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSdcr2FieldRburstMask) != 0
}

// SetRburst Burst read This bit enables burst read mode. The SDRAM controller anticipates the next read commands during the CAS latency and stores data in the Read FIFO. Note: The corresponding bit in the FMC_SDCR2 register is read only.
func (r *registerSdcr2Type) SetRburst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSdcr2FieldRburstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSdcr2FieldRburstMask)
	}
}

const (
	RegisterSdcr2FieldRpipeShift = 13
	RegisterSdcr2FieldRpipeMask  = 0x6000
)

// GetRpipe Read pipe These bits define the delay, in KCK_FMC clock cycles, for reading data after CAS latency. Note: The corresponding bits in the FMC_SDCR2 register is read only.
func (r *registerSdcr2Type) GetRpipe() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcr2FieldRpipeMask) >> RegisterSdcr2FieldRpipeShift)
}

// SetRpipe Read pipe These bits define the delay, in KCK_FMC clock cycles, for reading data after CAS latency. Note: The corresponding bits in the FMC_SDCR2 register is read only.
func (r *registerSdcr2Type) SetRpipe(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcr2FieldRpipeMask)|(uint32(value)<<RegisterSdcr2FieldRpipeShift))
}

// registerSdtr1Type This register contains the timing parameters of each SDRAM bank
type registerSdtr1Type uint32

const (
	RegisterSdtr1FieldTmrdShift = 0
	RegisterSdtr1FieldTmrdMask  = 0xf
)

// GetTmrd Load Mode Register to Active These bits define the delay between a Load Mode Register command and an Active or Refresh command in number of memory clock cycles. ....
func (r *registerSdtr1Type) GetTmrd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr1FieldTmrdMask) >> RegisterSdtr1FieldTmrdShift)
}

// SetTmrd Load Mode Register to Active These bits define the delay between a Load Mode Register command and an Active or Refresh command in number of memory clock cycles. ....
func (r *registerSdtr1Type) SetTmrd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr1FieldTmrdMask)|(uint32(value)<<RegisterSdtr1FieldTmrdShift))
}

const (
	RegisterSdtr1FieldTxsrShift = 4
	RegisterSdtr1FieldTxsrMask  = 0xf0
)

// GetTxsr Exit Self-refresh delay These bits define the delay from releasing the Self-refresh command to issuing the Activate command in number of memory clock cycles. .... Note: If two SDRAM devices are used, the FMC_SDTR1 and FMC_SDTR2 must be programmed with the same TXSR timing corresponding to the slowest SDRAM device.
func (r *registerSdtr1Type) GetTxsr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr1FieldTxsrMask) >> RegisterSdtr1FieldTxsrShift)
}

// SetTxsr Exit Self-refresh delay These bits define the delay from releasing the Self-refresh command to issuing the Activate command in number of memory clock cycles. .... Note: If two SDRAM devices are used, the FMC_SDTR1 and FMC_SDTR2 must be programmed with the same TXSR timing corresponding to the slowest SDRAM device.
func (r *registerSdtr1Type) SetTxsr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr1FieldTxsrMask)|(uint32(value)<<RegisterSdtr1FieldTxsrShift))
}

const (
	RegisterSdtr1FieldTrasShift = 8
	RegisterSdtr1FieldTrasMask  = 0xf00
)

// GetTras Self refresh time These bits define the minimum Self-refresh period in number of memory clock cycles. ....
func (r *registerSdtr1Type) GetTras() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr1FieldTrasMask) >> RegisterSdtr1FieldTrasShift)
}

// SetTras Self refresh time These bits define the minimum Self-refresh period in number of memory clock cycles. ....
func (r *registerSdtr1Type) SetTras(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr1FieldTrasMask)|(uint32(value)<<RegisterSdtr1FieldTrasShift))
}

const (
	RegisterSdtr1FieldTrcShift = 12
	RegisterSdtr1FieldTrcMask  = 0xf000
)

// GetTrc Row cycle delay These bits define the delay between the Refresh command and the Activate command, as well as the delay between two consecutive Refresh commands. It is expressed in number of memory clock cycles. The TRC timing is only configured in the FMC_SDTR1 register. If two SDRAM devices are used, the TRC must be programmed with the timings of the slowest device. .... Note: TRC must match the TRC and TRFC (Auto Refresh period) timings defined in the SDRAM device datasheet. Note: The corresponding bits in the FMC_SDTR2 register are dont care.
func (r *registerSdtr1Type) GetTrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr1FieldTrcMask) >> RegisterSdtr1FieldTrcShift)
}

// SetTrc Row cycle delay These bits define the delay between the Refresh command and the Activate command, as well as the delay between two consecutive Refresh commands. It is expressed in number of memory clock cycles. The TRC timing is only configured in the FMC_SDTR1 register. If two SDRAM devices are used, the TRC must be programmed with the timings of the slowest device. .... Note: TRC must match the TRC and TRFC (Auto Refresh period) timings defined in the SDRAM device datasheet. Note: The corresponding bits in the FMC_SDTR2 register are dont care.
func (r *registerSdtr1Type) SetTrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr1FieldTrcMask)|(uint32(value)<<RegisterSdtr1FieldTrcShift))
}

const (
	RegisterSdtr1FieldTwrShift = 16
	RegisterSdtr1FieldTwrMask  = 0xf0000
)

// GetTwr Recovery delay These bits define the delay between a Write and a Precharge command in number of memory clock cycles. .... Note: TWR must be programmed to match the write recovery time (tWR) defined in the SDRAM datasheet, and to guarantee that: TWR &#8805; TRAS - TRCD and TWR &#8805;TRC - TRCD - TRP Example: TRAS= 4 cycles, TRCD= 2 cycles. So, TWR &gt;= 2 cycles. TWR must be programmed to 0x1. If two SDRAM devices are used, the FMC_SDTR1 and FMC_SDTR2 must be programmed with the same TWR timing corresponding to the slowest SDRAM device.
func (r *registerSdtr1Type) GetTwr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr1FieldTwrMask) >> RegisterSdtr1FieldTwrShift)
}

// SetTwr Recovery delay These bits define the delay between a Write and a Precharge command in number of memory clock cycles. .... Note: TWR must be programmed to match the write recovery time (tWR) defined in the SDRAM datasheet, and to guarantee that: TWR &#8805; TRAS - TRCD and TWR &#8805;TRC - TRCD - TRP Example: TRAS= 4 cycles, TRCD= 2 cycles. So, TWR &gt;= 2 cycles. TWR must be programmed to 0x1. If two SDRAM devices are used, the FMC_SDTR1 and FMC_SDTR2 must be programmed with the same TWR timing corresponding to the slowest SDRAM device.
func (r *registerSdtr1Type) SetTwr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr1FieldTwrMask)|(uint32(value)<<RegisterSdtr1FieldTwrShift))
}

const (
	RegisterSdtr1FieldTrpShift = 20
	RegisterSdtr1FieldTrpMask  = 0xf00000
)

// GetTrp Row precharge delay These bits define the delay between a Precharge command and another command in number of memory clock cycles. The TRP timing is only configured in the FMC_SDTR1 register. If two SDRAM devices are used, the TRP must be programmed with the timing of the slowest device. .... Note: The corresponding bits in the FMC_SDTR2 register are dont care.
func (r *registerSdtr1Type) GetTrp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr1FieldTrpMask) >> RegisterSdtr1FieldTrpShift)
}

// SetTrp Row precharge delay These bits define the delay between a Precharge command and another command in number of memory clock cycles. The TRP timing is only configured in the FMC_SDTR1 register. If two SDRAM devices are used, the TRP must be programmed with the timing of the slowest device. .... Note: The corresponding bits in the FMC_SDTR2 register are dont care.
func (r *registerSdtr1Type) SetTrp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr1FieldTrpMask)|(uint32(value)<<RegisterSdtr1FieldTrpShift))
}

const (
	RegisterSdtr1FieldTrcdShift = 24
	RegisterSdtr1FieldTrcdMask  = 0xf000000
)

// GetTrcd Row to column delay These bits define the delay between the Activate command and a Read/Write command in number of memory clock cycles. ....
func (r *registerSdtr1Type) GetTrcd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr1FieldTrcdMask) >> RegisterSdtr1FieldTrcdShift)
}

// SetTrcd Row to column delay These bits define the delay between the Activate command and a Read/Write command in number of memory clock cycles. ....
func (r *registerSdtr1Type) SetTrcd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr1FieldTrcdMask)|(uint32(value)<<RegisterSdtr1FieldTrcdShift))
}

// registerSdtr2Type This register contains the timing parameters of each SDRAM bank
type registerSdtr2Type uint32

const (
	RegisterSdtr2FieldTmrdShift = 0
	RegisterSdtr2FieldTmrdMask  = 0xf
)

// GetTmrd Load Mode Register to Active These bits define the delay between a Load Mode Register command and an Active or Refresh command in number of memory clock cycles. ....
func (r *registerSdtr2Type) GetTmrd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr2FieldTmrdMask) >> RegisterSdtr2FieldTmrdShift)
}

// SetTmrd Load Mode Register to Active These bits define the delay between a Load Mode Register command and an Active or Refresh command in number of memory clock cycles. ....
func (r *registerSdtr2Type) SetTmrd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr2FieldTmrdMask)|(uint32(value)<<RegisterSdtr2FieldTmrdShift))
}

const (
	RegisterSdtr2FieldTxsrShift = 4
	RegisterSdtr2FieldTxsrMask  = 0xf0
)

// GetTxsr Exit Self-refresh delay These bits define the delay from releasing the Self-refresh command to issuing the Activate command in number of memory clock cycles. .... Note: If two SDRAM devices are used, the FMC_SDTR1 and FMC_SDTR2 must be programmed with the same TXSR timing corresponding to the slowest SDRAM device.
func (r *registerSdtr2Type) GetTxsr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr2FieldTxsrMask) >> RegisterSdtr2FieldTxsrShift)
}

// SetTxsr Exit Self-refresh delay These bits define the delay from releasing the Self-refresh command to issuing the Activate command in number of memory clock cycles. .... Note: If two SDRAM devices are used, the FMC_SDTR1 and FMC_SDTR2 must be programmed with the same TXSR timing corresponding to the slowest SDRAM device.
func (r *registerSdtr2Type) SetTxsr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr2FieldTxsrMask)|(uint32(value)<<RegisterSdtr2FieldTxsrShift))
}

const (
	RegisterSdtr2FieldTrasShift = 8
	RegisterSdtr2FieldTrasMask  = 0xf00
)

// GetTras Self refresh time These bits define the minimum Self-refresh period in number of memory clock cycles. ....
func (r *registerSdtr2Type) GetTras() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr2FieldTrasMask) >> RegisterSdtr2FieldTrasShift)
}

// SetTras Self refresh time These bits define the minimum Self-refresh period in number of memory clock cycles. ....
func (r *registerSdtr2Type) SetTras(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr2FieldTrasMask)|(uint32(value)<<RegisterSdtr2FieldTrasShift))
}

const (
	RegisterSdtr2FieldTrcShift = 12
	RegisterSdtr2FieldTrcMask  = 0xf000
)

// GetTrc Row cycle delay These bits define the delay between the Refresh command and the Activate command, as well as the delay between two consecutive Refresh commands. It is expressed in number of memory clock cycles. The TRC timing is only configured in the FMC_SDTR1 register. If two SDRAM devices are used, the TRC must be programmed with the timings of the slowest device. .... Note: TRC must match the TRC and TRFC (Auto Refresh period) timings defined in the SDRAM device datasheet. Note: The corresponding bits in the FMC_SDTR2 register are dont care.
func (r *registerSdtr2Type) GetTrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr2FieldTrcMask) >> RegisterSdtr2FieldTrcShift)
}

// SetTrc Row cycle delay These bits define the delay between the Refresh command and the Activate command, as well as the delay between two consecutive Refresh commands. It is expressed in number of memory clock cycles. The TRC timing is only configured in the FMC_SDTR1 register. If two SDRAM devices are used, the TRC must be programmed with the timings of the slowest device. .... Note: TRC must match the TRC and TRFC (Auto Refresh period) timings defined in the SDRAM device datasheet. Note: The corresponding bits in the FMC_SDTR2 register are dont care.
func (r *registerSdtr2Type) SetTrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr2FieldTrcMask)|(uint32(value)<<RegisterSdtr2FieldTrcShift))
}

const (
	RegisterSdtr2FieldTwrShift = 16
	RegisterSdtr2FieldTwrMask  = 0xf0000
)

// GetTwr Recovery delay These bits define the delay between a Write and a Precharge command in number of memory clock cycles. .... Note: TWR must be programmed to match the write recovery time (tWR) defined in the SDRAM datasheet, and to guarantee that: TWR &#8805; TRAS - TRCD and TWR &#8805;TRC - TRCD - TRP Example: TRAS= 4 cycles, TRCD= 2 cycles. So, TWR &gt;= 2 cycles. TWR must be programmed to 0x1. If two SDRAM devices are used, the FMC_SDTR1 and FMC_SDTR2 must be programmed with the same TWR timing corresponding to the slowest SDRAM device.
func (r *registerSdtr2Type) GetTwr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr2FieldTwrMask) >> RegisterSdtr2FieldTwrShift)
}

// SetTwr Recovery delay These bits define the delay between a Write and a Precharge command in number of memory clock cycles. .... Note: TWR must be programmed to match the write recovery time (tWR) defined in the SDRAM datasheet, and to guarantee that: TWR &#8805; TRAS - TRCD and TWR &#8805;TRC - TRCD - TRP Example: TRAS= 4 cycles, TRCD= 2 cycles. So, TWR &gt;= 2 cycles. TWR must be programmed to 0x1. If two SDRAM devices are used, the FMC_SDTR1 and FMC_SDTR2 must be programmed with the same TWR timing corresponding to the slowest SDRAM device.
func (r *registerSdtr2Type) SetTwr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr2FieldTwrMask)|(uint32(value)<<RegisterSdtr2FieldTwrShift))
}

const (
	RegisterSdtr2FieldTrpShift = 20
	RegisterSdtr2FieldTrpMask  = 0xf00000
)

// GetTrp Row precharge delay These bits define the delay between a Precharge command and another command in number of memory clock cycles. The TRP timing is only configured in the FMC_SDTR1 register. If two SDRAM devices are used, the TRP must be programmed with the timing of the slowest device. .... Note: The corresponding bits in the FMC_SDTR2 register are dont care.
func (r *registerSdtr2Type) GetTrp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr2FieldTrpMask) >> RegisterSdtr2FieldTrpShift)
}

// SetTrp Row precharge delay These bits define the delay between a Precharge command and another command in number of memory clock cycles. The TRP timing is only configured in the FMC_SDTR1 register. If two SDRAM devices are used, the TRP must be programmed with the timing of the slowest device. .... Note: The corresponding bits in the FMC_SDTR2 register are dont care.
func (r *registerSdtr2Type) SetTrp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr2FieldTrpMask)|(uint32(value)<<RegisterSdtr2FieldTrpShift))
}

const (
	RegisterSdtr2FieldTrcdShift = 24
	RegisterSdtr2FieldTrcdMask  = 0xf000000
)

// GetTrcd Row to column delay These bits define the delay between the Activate command and a Read/Write command in number of memory clock cycles. ....
func (r *registerSdtr2Type) GetTrcd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdtr2FieldTrcdMask) >> RegisterSdtr2FieldTrcdShift)
}

// SetTrcd Row to column delay These bits define the delay between the Activate command and a Read/Write command in number of memory clock cycles. ....
func (r *registerSdtr2Type) SetTrcd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdtr2FieldTrcdMask)|(uint32(value)<<RegisterSdtr2FieldTrcdShift))
}

// registerSdcmrType This register contains the command issued when the SDRAM device is accessed. This register is used to initialize the SDRAM device, and to activate the Self-refresh and the Power-down modes. As soon as the MODE field is written, the command will be issued only to one or to both SDRAM banks according to CTB1 and CTB2 command bits. This register is the same for both SDRAM banks.
type registerSdcmrType uint32

const (
	RegisterSdcmrFieldModeShift = 0
	RegisterSdcmrFieldModeMask  = 0x7
)

// GetMode Command mode These bits define the command issued to the SDRAM device. Note: When a command is issued, at least one Command Target Bank bit ( CTB1 or CTB2) must be set otherwise the command will be ignored. Note: If two SDRAM banks are used, the Auto-refresh and PALL command must be issued simultaneously to the two devices with CTB1 and CTB2 bits set otherwise the command will be ignored. Note: If only one SDRAM bank is used and a command is issued with its associated CTB bit set, the other CTB bit of the unused bank must be kept to 0.
func (r *registerSdcmrType) GetMode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcmrFieldModeMask) >> RegisterSdcmrFieldModeShift)
}

// SetMode Command mode These bits define the command issued to the SDRAM device. Note: When a command is issued, at least one Command Target Bank bit ( CTB1 or CTB2) must be set otherwise the command will be ignored. Note: If two SDRAM banks are used, the Auto-refresh and PALL command must be issued simultaneously to the two devices with CTB1 and CTB2 bits set otherwise the command will be ignored. Note: If only one SDRAM bank is used and a command is issued with its associated CTB bit set, the other CTB bit of the unused bank must be kept to 0.
func (r *registerSdcmrType) SetMode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcmrFieldModeMask)|(uint32(value)<<RegisterSdcmrFieldModeShift))
}

const (
	RegisterSdcmrFieldCtb2Shift = 3
	RegisterSdcmrFieldCtb2Mask  = 0x8
)

// GetCtb2 Command Target Bank 2 This bit indicates whether the command will be issued to SDRAM Bank 2 or not.
func (r *registerSdcmrType) GetCtb2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSdcmrFieldCtb2Mask) != 0
}

// SetCtb2 Command Target Bank 2 This bit indicates whether the command will be issued to SDRAM Bank 2 or not.
func (r *registerSdcmrType) SetCtb2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSdcmrFieldCtb2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSdcmrFieldCtb2Mask)
	}
}

const (
	RegisterSdcmrFieldCtb1Shift = 4
	RegisterSdcmrFieldCtb1Mask  = 0x10
)

// GetCtb1 Command Target Bank 1 This bit indicates whether the command will be issued to SDRAM Bank 1 or not.
func (r *registerSdcmrType) GetCtb1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSdcmrFieldCtb1Mask) != 0
}

// SetCtb1 Command Target Bank 1 This bit indicates whether the command will be issued to SDRAM Bank 1 or not.
func (r *registerSdcmrType) SetCtb1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSdcmrFieldCtb1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSdcmrFieldCtb1Mask)
	}
}

const (
	RegisterSdcmrFieldNrfsShift = 5
	RegisterSdcmrFieldNrfsMask  = 0x1e0
)

// GetNrfs Number of Auto-refresh These bits define the number of consecutive Auto-refresh commands issued when MODE = 011. ....
func (r *registerSdcmrType) GetNrfs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdcmrFieldNrfsMask) >> RegisterSdcmrFieldNrfsShift)
}

// SetNrfs Number of Auto-refresh These bits define the number of consecutive Auto-refresh commands issued when MODE = 011. ....
func (r *registerSdcmrType) SetNrfs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcmrFieldNrfsMask)|(uint32(value)<<RegisterSdcmrFieldNrfsShift))
}

const (
	RegisterSdcmrFieldMrdShift = 9
	RegisterSdcmrFieldMrdMask  = 0x7ffe00
)

// GetMrd Mode Register definition This 14-bit field defines the SDRAM Mode Register content. The Mode Register is programmed using the Load Mode Register command. The MRD[13:0] bits are also used to program the extended mode register for mobile SDRAM.
func (r *registerSdcmrType) GetMrd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSdcmrFieldMrdMask) >> RegisterSdcmrFieldMrdShift)
}

// SetMrd Mode Register definition This 14-bit field defines the SDRAM Mode Register content. The Mode Register is programmed using the Load Mode Register command. The MRD[13:0] bits are also used to program the extended mode register for mobile SDRAM.
func (r *registerSdcmrType) SetMrd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdcmrFieldMrdMask)|(uint32(value)<<RegisterSdcmrFieldMrdShift))
}

// registerSdrtrType This register sets the refresh rate in number of SDCLK clock cycles between the refresh cycles by configuring the Refresh Timer Count value.Examplewhere 64 ms is the SDRAM refresh period.The refresh rate must be increased by 20 SDRAM clock cycles (as in the above example) to obtain a safe margin if an internal refresh request occurs when a read request has been accepted. It corresponds to a COUNT value of 0000111000000 (448). This 13-bit field is loaded into a timer which is decremented using the SDRAM clock. This timer generates a refresh pulse when zero is reached. The COUNT value must be set at least to 41 SDRAM clock cycles.As soon as the FMC_SDRTR register is programmed, the timer starts counting. If the value programmed in the register is 0, no refresh is carried out. This register must not be reprogrammed after the initialization procedure to avoid modifying the refresh rate.Each time a refresh pulse is generated, this 13-bit COUNT field is reloaded into the counter.If a memory access is in progress, the Auto-refresh request is delayed. However, if the memory access and Auto-refresh requests are generated simultaneously, the Auto-refresh takes precedence. If the memory access occurs during a refresh operation, the request is buffered to be processed when the refresh is complete.This register is common to SDRAM bank 1 and bank 2.
type registerSdrtrType uint32

const (
	RegisterSdrtrFieldCreShift = 0
	RegisterSdrtrFieldCreMask  = 0x1
)

// SetCre Clear Refresh error flag This bit is used to clear the Refresh Error Flag (RE) in the Status Register.
func (r *registerSdrtrType) SetCre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSdrtrFieldCreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSdrtrFieldCreMask)
	}
}

const (
	RegisterSdrtrFieldCountShift = 1
	RegisterSdrtrFieldCountMask  = 0x3ffe
)

// GetCount Refresh Timer Count This 13-bit field defines the refresh rate of the SDRAM device. It is expressed in number of memory clock cycles. It must be set at least to 41 SDRAM clock cycles (0x29). Refresh rate = (COUNT + 1) x SDRAM frequency clock COUNT = (SDRAM refresh period / Number of rows) - 20
func (r *registerSdrtrType) GetCount() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSdrtrFieldCountMask) >> RegisterSdrtrFieldCountShift)
}

// SetCount Refresh Timer Count This 13-bit field defines the refresh rate of the SDRAM device. It is expressed in number of memory clock cycles. It must be set at least to 41 SDRAM clock cycles (0x29). Refresh rate = (COUNT + 1) x SDRAM frequency clock COUNT = (SDRAM refresh period / Number of rows) - 20
func (r *registerSdrtrType) SetCount(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdrtrFieldCountMask)|(uint32(value)<<RegisterSdrtrFieldCountShift))
}

const (
	RegisterSdrtrFieldReieShift = 14
	RegisterSdrtrFieldReieMask  = 0x4000
)

// GetReie RES Interrupt Enable
func (r *registerSdrtrType) GetReie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSdrtrFieldReieMask) != 0
}

// SetReie RES Interrupt Enable
func (r *registerSdrtrType) SetReie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSdrtrFieldReieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSdrtrFieldReieMask)
	}
}

// registerSdsrType SDRAM Status register
type registerSdsrType uint32

const (
	RegisterSdsrFieldReShift = 0
	RegisterSdsrFieldReMask  = 0x1
)

// GetRe Refresh error flag An interrupt is generated if REIE = 1 and RE = 1
func (r *registerSdsrType) GetRe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSdsrFieldReMask) != 0
}

// SetRe Refresh error flag An interrupt is generated if REIE = 1 and RE = 1
func (r *registerSdsrType) SetRe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSdsrFieldReMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSdsrFieldReMask)
	}
}

const (
	RegisterSdsrFieldModes1Shift = 1
	RegisterSdsrFieldModes1Mask  = 0x6
)

// GetModes1 Status Mode for Bank 1 These bits define the Status Mode of SDRAM Bank 1.
func (r *registerSdsrType) GetModes1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdsrFieldModes1Mask) >> RegisterSdsrFieldModes1Shift)
}

// SetModes1 Status Mode for Bank 1 These bits define the Status Mode of SDRAM Bank 1.
func (r *registerSdsrType) SetModes1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdsrFieldModes1Mask)|(uint32(value)<<RegisterSdsrFieldModes1Shift))
}

const (
	RegisterSdsrFieldModes2Shift = 3
	RegisterSdsrFieldModes2Mask  = 0x18
)

// GetModes2 Status Mode for Bank 2 These bits define the Status Mode of SDRAM Bank 2.
func (r *registerSdsrType) GetModes2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSdsrFieldModes2Mask) >> RegisterSdsrFieldModes2Shift)
}

// SetModes2 Status Mode for Bank 2 These bits define the Status Mode of SDRAM Bank 2.
func (r *registerSdsrType) SetModes2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSdsrFieldModes2Mask)|(uint32(value)<<RegisterSdsrFieldModes2Shift))
}
