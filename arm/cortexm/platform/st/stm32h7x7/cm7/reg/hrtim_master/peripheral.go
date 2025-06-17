//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package hrtim_master

import (
	"unsafe"
	"volatile"
)

var (
	Hrtim_master = (*_hrtim_master)(unsafe.Pointer(uintptr(0x40017400)))
)

type _hrtim_master struct {
	Mcr    RegisterMcrType
	Misr   RegisterMisrType
	Micr   RegisterMicrType
	Mdier4 RegisterMdier4Type
	Mcntr  RegisterMcntrType
	Mper   RegisterMperType
	Mrep   RegisterMrepType
	Mcmp1r RegisterMcmp1rType
	_      [4]byte
	Mcmp2r RegisterMcmp2rType
	Mcmp3r RegisterMcmp3rType
	Mcmp4r RegisterMcmp4rType
}

// RegisterMcrType Master Timer Control Register
type RegisterMcrType uint32

func (r *RegisterMcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMcrFieldCkpscShift = 0
	RegisterMcrFieldCkpscMask  = 0x7
)

// GetCkpsc HRTIM Master Clock prescaler
func (r *RegisterMcrType) GetCkpsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldCkpscMask) >> RegisterMcrFieldCkpscShift)
}

// SetCkpsc HRTIM Master Clock prescaler
func (r *RegisterMcrType) SetCkpsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldCkpscMask)|(uint32(value)<<RegisterMcrFieldCkpscShift))
}

const (
	RegisterMcrFieldContShift = 3
	RegisterMcrFieldContMask  = 0x8
)

// GetCont Master Continuous mode
func (r *RegisterMcrType) GetCont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldContMask) != 0
}

// SetCont Master Continuous mode
func (r *RegisterMcrType) SetCont(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldContMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldContMask)
	}
}

const (
	RegisterMcrFieldRetrigShift = 4
	RegisterMcrFieldRetrigMask  = 0x10
)

// GetRetrig Master Re-triggerable mode
func (r *RegisterMcrType) GetRetrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldRetrigMask) != 0
}

// SetRetrig Master Re-triggerable mode
func (r *RegisterMcrType) SetRetrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldRetrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldRetrigMask)
	}
}

const (
	RegisterMcrFieldHalfShift = 5
	RegisterMcrFieldHalfMask  = 0x20
)

// GetHalf Half mode enable
func (r *RegisterMcrType) GetHalf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldHalfMask) != 0
}

// SetHalf Half mode enable
func (r *RegisterMcrType) SetHalf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldHalfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldHalfMask)
	}
}

const (
	RegisterMcrFieldSyncinShift = 8
	RegisterMcrFieldSyncinMask  = 0x300
)

// GetSyncin ynchronization input
func (r *RegisterMcrType) GetSyncin() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldSyncinMask) >> RegisterMcrFieldSyncinShift)
}

// SetSyncin ynchronization input
func (r *RegisterMcrType) SetSyncin(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldSyncinMask)|(uint32(value)<<RegisterMcrFieldSyncinShift))
}

const (
	RegisterMcrFieldSyncrstmShift = 10
	RegisterMcrFieldSyncrstmMask  = 0x400
)

// GetSyncrstm Synchronization Resets Master
func (r *RegisterMcrType) GetSyncrstm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldSyncrstmMask) != 0
}

// SetSyncrstm Synchronization Resets Master
func (r *RegisterMcrType) SetSyncrstm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldSyncrstmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldSyncrstmMask)
	}
}

const (
	RegisterMcrFieldSyncstrtmShift = 11
	RegisterMcrFieldSyncstrtmMask  = 0x800
)

// GetSyncstrtm Synchronization Starts Master
func (r *RegisterMcrType) GetSyncstrtm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldSyncstrtmMask) != 0
}

// SetSyncstrtm Synchronization Starts Master
func (r *RegisterMcrType) SetSyncstrtm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldSyncstrtmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldSyncstrtmMask)
	}
}

const (
	RegisterMcrFieldSyncoutShift = 12
	RegisterMcrFieldSyncoutMask  = 0x3000
)

// GetSyncout Synchronization output
func (r *RegisterMcrType) GetSyncout() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldSyncoutMask) >> RegisterMcrFieldSyncoutShift)
}

// SetSyncout Synchronization output
func (r *RegisterMcrType) SetSyncout(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldSyncoutMask)|(uint32(value)<<RegisterMcrFieldSyncoutShift))
}

const (
	RegisterMcrFieldSyncsrcShift = 14
	RegisterMcrFieldSyncsrcMask  = 0xc000
)

// GetSyncsrc Synchronization source
func (r *RegisterMcrType) GetSyncsrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldSyncsrcMask) >> RegisterMcrFieldSyncsrcShift)
}

// SetSyncsrc Synchronization source
func (r *RegisterMcrType) SetSyncsrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldSyncsrcMask)|(uint32(value)<<RegisterMcrFieldSyncsrcShift))
}

const (
	RegisterMcrFieldMcenShift = 16
	RegisterMcrFieldMcenMask  = 0x10000
)

// GetMcen Master Counter enable
func (r *RegisterMcrType) GetMcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldMcenMask) != 0
}

// SetMcen Master Counter enable
func (r *RegisterMcrType) SetMcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldMcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldMcenMask)
	}
}

const (
	RegisterMcrFieldTacenShift = 17
	RegisterMcrFieldTacenMask  = 0x20000
)

// GetTacen Timer A counter enable
func (r *RegisterMcrType) GetTacen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldTacenMask) != 0
}

// SetTacen Timer A counter enable
func (r *RegisterMcrType) SetTacen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldTacenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldTacenMask)
	}
}

const (
	RegisterMcrFieldTbcenShift = 18
	RegisterMcrFieldTbcenMask  = 0x40000
)

// GetTbcen Timer B counter enable
func (r *RegisterMcrType) GetTbcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldTbcenMask) != 0
}

// SetTbcen Timer B counter enable
func (r *RegisterMcrType) SetTbcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldTbcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldTbcenMask)
	}
}

const (
	RegisterMcrFieldTccenShift = 19
	RegisterMcrFieldTccenMask  = 0x80000
)

// GetTccen Timer C counter enable
func (r *RegisterMcrType) GetTccen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldTccenMask) != 0
}

// SetTccen Timer C counter enable
func (r *RegisterMcrType) SetTccen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldTccenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldTccenMask)
	}
}

const (
	RegisterMcrFieldTdcenShift = 20
	RegisterMcrFieldTdcenMask  = 0x100000
)

// GetTdcen Timer D counter enable
func (r *RegisterMcrType) GetTdcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldTdcenMask) != 0
}

// SetTdcen Timer D counter enable
func (r *RegisterMcrType) SetTdcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldTdcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldTdcenMask)
	}
}

const (
	RegisterMcrFieldTecenShift = 21
	RegisterMcrFieldTecenMask  = 0x200000
)

// GetTecen Timer E counter enable
func (r *RegisterMcrType) GetTecen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldTecenMask) != 0
}

// SetTecen Timer E counter enable
func (r *RegisterMcrType) SetTecen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldTecenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldTecenMask)
	}
}

const (
	RegisterMcrFieldDacsyncShift = 25
	RegisterMcrFieldDacsyncMask  = 0x6000000
)

// GetDacsync AC Synchronization
func (r *RegisterMcrType) GetDacsync() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldDacsyncMask) >> RegisterMcrFieldDacsyncShift)
}

// SetDacsync AC Synchronization
func (r *RegisterMcrType) SetDacsync(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldDacsyncMask)|(uint32(value)<<RegisterMcrFieldDacsyncShift))
}

const (
	RegisterMcrFieldPreenShift = 27
	RegisterMcrFieldPreenMask  = 0x8000000
)

// GetPreen Preload enable
func (r *RegisterMcrType) GetPreen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldPreenMask) != 0
}

// SetPreen Preload enable
func (r *RegisterMcrType) SetPreen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldPreenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldPreenMask)
	}
}

const (
	RegisterMcrFieldMrepuShift = 29
	RegisterMcrFieldMrepuMask  = 0x20000000
)

// GetMrepu Master Timer Repetition update
func (r *RegisterMcrType) GetMrepu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldMrepuMask) != 0
}

// SetMrepu Master Timer Repetition update
func (r *RegisterMcrType) SetMrepu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldMrepuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldMrepuMask)
	}
}

const (
	RegisterMcrFieldBrstdmaShift = 30
	RegisterMcrFieldBrstdmaMask  = 0xc0000000
)

// GetBrstdma Burst DMA Update
func (r *RegisterMcrType) GetBrstdma() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldBrstdmaMask) >> RegisterMcrFieldBrstdmaShift)
}

// SetBrstdma Burst DMA Update
func (r *RegisterMcrType) SetBrstdma(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldBrstdmaMask)|(uint32(value)<<RegisterMcrFieldBrstdmaShift))
}

// RegisterMisrType Master Timer Interrupt Status Register
type RegisterMisrType uint32

func (r *RegisterMisrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMisrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMisrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMisrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMisrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMisrFieldMcmp1Shift = 0
	RegisterMisrFieldMcmp1Mask  = 0x1
)

// GetMcmp1 Master Compare 1 Interrupt Flag
func (r *RegisterMisrType) GetMcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMcmp1Mask) != 0
}

// SetMcmp1 Master Compare 1 Interrupt Flag
func (r *RegisterMisrType) SetMcmp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisrFieldMcmp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisrFieldMcmp1Mask)
	}
}

const (
	RegisterMisrFieldMcmp2Shift = 1
	RegisterMisrFieldMcmp2Mask  = 0x2
)

// GetMcmp2 Master Compare 2 Interrupt Flag
func (r *RegisterMisrType) GetMcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMcmp2Mask) != 0
}

// SetMcmp2 Master Compare 2 Interrupt Flag
func (r *RegisterMisrType) SetMcmp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisrFieldMcmp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisrFieldMcmp2Mask)
	}
}

const (
	RegisterMisrFieldMcmp3Shift = 2
	RegisterMisrFieldMcmp3Mask  = 0x4
)

// GetMcmp3 Master Compare 3 Interrupt Flag
func (r *RegisterMisrType) GetMcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMcmp3Mask) != 0
}

// SetMcmp3 Master Compare 3 Interrupt Flag
func (r *RegisterMisrType) SetMcmp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisrFieldMcmp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisrFieldMcmp3Mask)
	}
}

const (
	RegisterMisrFieldMcmp4Shift = 3
	RegisterMisrFieldMcmp4Mask  = 0x8
)

// GetMcmp4 Master Compare 4 Interrupt Flag
func (r *RegisterMisrType) GetMcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMcmp4Mask) != 0
}

// SetMcmp4 Master Compare 4 Interrupt Flag
func (r *RegisterMisrType) SetMcmp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisrFieldMcmp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisrFieldMcmp4Mask)
	}
}

const (
	RegisterMisrFieldMrepShift = 4
	RegisterMisrFieldMrepMask  = 0x10
)

// GetMrep Master Repetition Interrupt Flag
func (r *RegisterMisrType) GetMrep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMrepMask) != 0
}

// SetMrep Master Repetition Interrupt Flag
func (r *RegisterMisrType) SetMrep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisrFieldMrepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisrFieldMrepMask)
	}
}

const (
	RegisterMisrFieldSyncShift = 5
	RegisterMisrFieldSyncMask  = 0x20
)

// GetSync Sync Input Interrupt Flag
func (r *RegisterMisrType) GetSync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldSyncMask) != 0
}

// SetSync Sync Input Interrupt Flag
func (r *RegisterMisrType) SetSync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisrFieldSyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisrFieldSyncMask)
	}
}

const (
	RegisterMisrFieldMupdShift = 6
	RegisterMisrFieldMupdMask  = 0x40
)

// GetMupd Master Update Interrupt Flag
func (r *RegisterMisrType) GetMupd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMupdMask) != 0
}

// SetMupd Master Update Interrupt Flag
func (r *RegisterMisrType) SetMupd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisrFieldMupdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisrFieldMupdMask)
	}
}

// RegisterMicrType Master Timer Interrupt Clear Register
type RegisterMicrType uint32

func (r *RegisterMicrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMicrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMicrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMicrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMicrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMicrFieldMcmp1cShift = 0
	RegisterMicrFieldMcmp1cMask  = 0x1
)

// GetMcmp1c Master Compare 1 Interrupt flag clear
func (r *RegisterMicrType) GetMcmp1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMcmp1cMask) != 0
}

// SetMcmp1c Master Compare 1 Interrupt flag clear
func (r *RegisterMicrType) SetMcmp1c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMicrFieldMcmp1cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMicrFieldMcmp1cMask)
	}
}

const (
	RegisterMicrFieldMcmp2cShift = 1
	RegisterMicrFieldMcmp2cMask  = 0x2
)

// GetMcmp2c Master Compare 2 Interrupt flag clear
func (r *RegisterMicrType) GetMcmp2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMcmp2cMask) != 0
}

// SetMcmp2c Master Compare 2 Interrupt flag clear
func (r *RegisterMicrType) SetMcmp2c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMicrFieldMcmp2cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMicrFieldMcmp2cMask)
	}
}

const (
	RegisterMicrFieldMcmp3cShift = 2
	RegisterMicrFieldMcmp3cMask  = 0x4
)

// GetMcmp3c Master Compare 3 Interrupt flag clear
func (r *RegisterMicrType) GetMcmp3c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMcmp3cMask) != 0
}

// SetMcmp3c Master Compare 3 Interrupt flag clear
func (r *RegisterMicrType) SetMcmp3c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMicrFieldMcmp3cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMicrFieldMcmp3cMask)
	}
}

const (
	RegisterMicrFieldMcmp4cShift = 3
	RegisterMicrFieldMcmp4cMask  = 0x8
)

// GetMcmp4c Master Compare 4 Interrupt flag clear
func (r *RegisterMicrType) GetMcmp4c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMcmp4cMask) != 0
}

// SetMcmp4c Master Compare 4 Interrupt flag clear
func (r *RegisterMicrType) SetMcmp4c(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMicrFieldMcmp4cMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMicrFieldMcmp4cMask)
	}
}

const (
	RegisterMicrFieldMrepcShift = 4
	RegisterMicrFieldMrepcMask  = 0x10
)

// GetMrepc Repetition Interrupt flag clear
func (r *RegisterMicrType) GetMrepc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMrepcMask) != 0
}

// SetMrepc Repetition Interrupt flag clear
func (r *RegisterMicrType) SetMrepc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMicrFieldMrepcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMicrFieldMrepcMask)
	}
}

const (
	RegisterMicrFieldSynccShift = 5
	RegisterMicrFieldSynccMask  = 0x20
)

// GetSyncc Sync Input Interrupt flag clear
func (r *RegisterMicrType) GetSyncc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldSynccMask) != 0
}

// SetSyncc Sync Input Interrupt flag clear
func (r *RegisterMicrType) SetSyncc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMicrFieldSynccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMicrFieldSynccMask)
	}
}

const (
	RegisterMicrFieldMupdcShift = 6
	RegisterMicrFieldMupdcMask  = 0x40
)

// GetMupdc Master update Interrupt flag clear
func (r *RegisterMicrType) GetMupdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMupdcMask) != 0
}

// SetMupdc Master update Interrupt flag clear
func (r *RegisterMicrType) SetMupdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMicrFieldMupdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMicrFieldMupdcMask)
	}
}

// RegisterMdier4Type MDIER4
type RegisterMdier4Type uint32

func (r *RegisterMdier4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMdier4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMdier4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMdier4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMdier4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMdier4FieldMcmp1ieShift = 0
	RegisterMdier4FieldMcmp1ieMask  = 0x1
)

// GetMcmp1ie MCMP1IE
func (r *RegisterMdier4Type) GetMcmp1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp1ieMask) != 0
}

// SetMcmp1ie MCMP1IE
func (r *RegisterMdier4Type) SetMcmp1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMcmp1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMcmp1ieMask)
	}
}

const (
	RegisterMdier4FieldMcmp2ieShift = 1
	RegisterMdier4FieldMcmp2ieMask  = 0x2
)

// GetMcmp2ie MCMP2IE
func (r *RegisterMdier4Type) GetMcmp2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp2ieMask) != 0
}

// SetMcmp2ie MCMP2IE
func (r *RegisterMdier4Type) SetMcmp2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMcmp2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMcmp2ieMask)
	}
}

const (
	RegisterMdier4FieldMcmp3ieShift = 2
	RegisterMdier4FieldMcmp3ieMask  = 0x4
)

// GetMcmp3ie MCMP3IE
func (r *RegisterMdier4Type) GetMcmp3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp3ieMask) != 0
}

// SetMcmp3ie MCMP3IE
func (r *RegisterMdier4Type) SetMcmp3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMcmp3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMcmp3ieMask)
	}
}

const (
	RegisterMdier4FieldMcmp4ieShift = 3
	RegisterMdier4FieldMcmp4ieMask  = 0x8
)

// GetMcmp4ie MCMP4IE
func (r *RegisterMdier4Type) GetMcmp4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp4ieMask) != 0
}

// SetMcmp4ie MCMP4IE
func (r *RegisterMdier4Type) SetMcmp4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMcmp4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMcmp4ieMask)
	}
}

const (
	RegisterMdier4FieldMrepieShift = 4
	RegisterMdier4FieldMrepieMask  = 0x10
)

// GetMrepie MREPIE
func (r *RegisterMdier4Type) GetMrepie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMrepieMask) != 0
}

// SetMrepie MREPIE
func (r *RegisterMdier4Type) SetMrepie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMrepieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMrepieMask)
	}
}

const (
	RegisterMdier4FieldSyncieShift = 5
	RegisterMdier4FieldSyncieMask  = 0x20
)

// GetSyncie SYNCIE
func (r *RegisterMdier4Type) GetSyncie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldSyncieMask) != 0
}

// SetSyncie SYNCIE
func (r *RegisterMdier4Type) SetSyncie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldSyncieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldSyncieMask)
	}
}

const (
	RegisterMdier4FieldMupdieShift = 6
	RegisterMdier4FieldMupdieMask  = 0x40
)

// GetMupdie MUPDIE
func (r *RegisterMdier4Type) GetMupdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMupdieMask) != 0
}

// SetMupdie MUPDIE
func (r *RegisterMdier4Type) SetMupdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMupdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMupdieMask)
	}
}

const (
	RegisterMdier4FieldMcmp1deShift = 16
	RegisterMdier4FieldMcmp1deMask  = 0x10000
)

// GetMcmp1de MCMP1DE
func (r *RegisterMdier4Type) GetMcmp1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp1deMask) != 0
}

// SetMcmp1de MCMP1DE
func (r *RegisterMdier4Type) SetMcmp1de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMcmp1deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMcmp1deMask)
	}
}

const (
	RegisterMdier4FieldMcmp2deShift = 17
	RegisterMdier4FieldMcmp2deMask  = 0x20000
)

// GetMcmp2de MCMP2DE
func (r *RegisterMdier4Type) GetMcmp2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp2deMask) != 0
}

// SetMcmp2de MCMP2DE
func (r *RegisterMdier4Type) SetMcmp2de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMcmp2deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMcmp2deMask)
	}
}

const (
	RegisterMdier4FieldMcmp3deShift = 18
	RegisterMdier4FieldMcmp3deMask  = 0x40000
)

// GetMcmp3de MCMP3DE
func (r *RegisterMdier4Type) GetMcmp3de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp3deMask) != 0
}

// SetMcmp3de MCMP3DE
func (r *RegisterMdier4Type) SetMcmp3de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMcmp3deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMcmp3deMask)
	}
}

const (
	RegisterMdier4FieldMcmp4deShift = 19
	RegisterMdier4FieldMcmp4deMask  = 0x80000
)

// GetMcmp4de MCMP4DE
func (r *RegisterMdier4Type) GetMcmp4de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp4deMask) != 0
}

// SetMcmp4de MCMP4DE
func (r *RegisterMdier4Type) SetMcmp4de(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMcmp4deMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMcmp4deMask)
	}
}

const (
	RegisterMdier4FieldMrepdeShift = 20
	RegisterMdier4FieldMrepdeMask  = 0x100000
)

// GetMrepde MREPDE
func (r *RegisterMdier4Type) GetMrepde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMrepdeMask) != 0
}

// SetMrepde MREPDE
func (r *RegisterMdier4Type) SetMrepde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMrepdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMrepdeMask)
	}
}

const (
	RegisterMdier4FieldSyncdeShift = 21
	RegisterMdier4FieldSyncdeMask  = 0x200000
)

// GetSyncde SYNCDE
func (r *RegisterMdier4Type) GetSyncde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldSyncdeMask) != 0
}

// SetSyncde SYNCDE
func (r *RegisterMdier4Type) SetSyncde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldSyncdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldSyncdeMask)
	}
}

const (
	RegisterMdier4FieldMupddeShift = 22
	RegisterMdier4FieldMupddeMask  = 0x400000
)

// GetMupdde MUPDDE
func (r *RegisterMdier4Type) GetMupdde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMupddeMask) != 0
}

// SetMupdde MUPDDE
func (r *RegisterMdier4Type) SetMupdde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMupddeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMupddeMask)
	}
}

// RegisterMcntrType Master Timer Counter Register
type RegisterMcntrType uint32

func (r *RegisterMcntrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMcntrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMcntrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMcntrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMcntrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMcntrFieldMcntShift = 0
	RegisterMcntrFieldMcntMask  = 0xffff
)

// GetMcnt Counter value
func (r *RegisterMcntrType) GetMcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMcntrFieldMcntMask) >> RegisterMcntrFieldMcntShift)
}

// SetMcnt Counter value
func (r *RegisterMcntrType) SetMcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcntrFieldMcntMask)|(uint32(value)<<RegisterMcntrFieldMcntShift))
}

// RegisterMperType Master Timer Period Register
type RegisterMperType uint32

func (r *RegisterMperType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMperType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMperType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMperType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMperType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMperFieldMperShift = 0
	RegisterMperFieldMperMask  = 0xffff
)

// GetMper Master Timer Period value
func (r *RegisterMperType) GetMper() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMperFieldMperMask) >> RegisterMperFieldMperShift)
}

// SetMper Master Timer Period value
func (r *RegisterMperType) SetMper(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMperFieldMperMask)|(uint32(value)<<RegisterMperFieldMperShift))
}

// RegisterMrepType Master Timer Repetition Register
type RegisterMrepType uint32

func (r *RegisterMrepType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMrepType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMrepType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMrepType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMrepType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMrepFieldMrepShift = 0
	RegisterMrepFieldMrepMask  = 0xff
)

// GetMrep Master Timer Repetition counter value
func (r *RegisterMrepType) GetMrep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMrepFieldMrepMask) >> RegisterMrepFieldMrepShift)
}

// SetMrep Master Timer Repetition counter value
func (r *RegisterMrepType) SetMrep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMrepFieldMrepMask)|(uint32(value)<<RegisterMrepFieldMrepShift))
}

// RegisterMcmp1rType Master Timer Compare 1 Register
type RegisterMcmp1rType uint32

func (r *RegisterMcmp1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMcmp1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMcmp1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMcmp1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMcmp1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMcmp1rFieldMcmp1Shift = 0
	RegisterMcmp1rFieldMcmp1Mask  = 0xffff
)

// GetMcmp1 Master Timer Compare 1 value
func (r *RegisterMcmp1rType) GetMcmp1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMcmp1rFieldMcmp1Mask) >> RegisterMcmp1rFieldMcmp1Shift)
}

// SetMcmp1 Master Timer Compare 1 value
func (r *RegisterMcmp1rType) SetMcmp1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcmp1rFieldMcmp1Mask)|(uint32(value)<<RegisterMcmp1rFieldMcmp1Shift))
}

// RegisterMcmp2rType Master Timer Compare 2 Register
type RegisterMcmp2rType uint32

func (r *RegisterMcmp2rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMcmp2rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMcmp2rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMcmp2rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMcmp2rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMcmp2rFieldMcmp2Shift = 0
	RegisterMcmp2rFieldMcmp2Mask  = 0xffff
)

// GetMcmp2 Master Timer Compare 2 value
func (r *RegisterMcmp2rType) GetMcmp2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMcmp2rFieldMcmp2Mask) >> RegisterMcmp2rFieldMcmp2Shift)
}

// SetMcmp2 Master Timer Compare 2 value
func (r *RegisterMcmp2rType) SetMcmp2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcmp2rFieldMcmp2Mask)|(uint32(value)<<RegisterMcmp2rFieldMcmp2Shift))
}

// RegisterMcmp3rType Master Timer Compare 3 Register
type RegisterMcmp3rType uint32

func (r *RegisterMcmp3rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMcmp3rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMcmp3rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMcmp3rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMcmp3rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMcmp3rFieldMcmp3Shift = 0
	RegisterMcmp3rFieldMcmp3Mask  = 0xffff
)

// GetMcmp3 Master Timer Compare 3 value
func (r *RegisterMcmp3rType) GetMcmp3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMcmp3rFieldMcmp3Mask) >> RegisterMcmp3rFieldMcmp3Shift)
}

// SetMcmp3 Master Timer Compare 3 value
func (r *RegisterMcmp3rType) SetMcmp3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcmp3rFieldMcmp3Mask)|(uint32(value)<<RegisterMcmp3rFieldMcmp3Shift))
}

// RegisterMcmp4rType Master Timer Compare 4 Register
type RegisterMcmp4rType uint32

func (r *RegisterMcmp4rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMcmp4rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMcmp4rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMcmp4rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMcmp4rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMcmp4rFieldMcmp4Shift = 0
	RegisterMcmp4rFieldMcmp4Mask  = 0xffff
)

// GetMcmp4 Master Timer Compare 4 value
func (r *RegisterMcmp4rType) GetMcmp4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMcmp4rFieldMcmp4Mask) >> RegisterMcmp4rFieldMcmp4Shift)
}

// SetMcmp4 Master Timer Compare 4 value
func (r *RegisterMcmp4rType) SetMcmp4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcmp4rFieldMcmp4Mask)|(uint32(value)<<RegisterMcmp4rFieldMcmp4Shift))
}
