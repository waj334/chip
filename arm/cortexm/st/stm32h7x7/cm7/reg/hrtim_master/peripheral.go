package hrtim_master

import (
	"unsafe"
	"volatile"
)

var (
	Hrtim_master = (*_hrtim_master)(unsafe.Pointer(uintptr(0x40017400)))
)

type _hrtim_master struct {
	Mcr    registerMcrType
	Misr   registerMisrType
	Micr   registerMicrType
	Mdier4 registerMdier4Type
	Mcntr  registerMcntrType
	Mper   registerMperType
	Mrep   registerMrepType
	Mcmp1r registerMcmp1rType
	_      [4]byte
	Mcmp2r registerMcmp2rType
	Mcmp3r registerMcmp3rType
	Mcmp4r registerMcmp4rType
}

// registerMcrType Master Timer Control Register
type registerMcrType uint32

const (
	RegisterMcrFieldCk_pscShift = 0
	RegisterMcrFieldCk_pscMask  = 0x7
)

// GetCk_psc HRTIM Master Clock prescaler
func (r *registerMcrType) GetCk_psc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldCk_pscMask) >> RegisterMcrFieldCk_pscShift)
}

// SetCk_psc HRTIM Master Clock prescaler
func (r *registerMcrType) SetCk_psc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldCk_pscMask)|(uint32(value)<<RegisterMcrFieldCk_pscShift))
}

const (
	RegisterMcrFieldContShift = 3
	RegisterMcrFieldContMask  = 0x8
)

// GetCont Master Continuous mode
func (r *registerMcrType) GetCont() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldContMask) != 0
}

// SetCont Master Continuous mode
func (r *registerMcrType) SetCont(value bool) {
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
func (r *registerMcrType) GetRetrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldRetrigMask) != 0
}

// SetRetrig Master Re-triggerable mode
func (r *registerMcrType) SetRetrig(value bool) {
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
func (r *registerMcrType) GetHalf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldHalfMask) != 0
}

// SetHalf Half mode enable
func (r *registerMcrType) SetHalf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldHalfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldHalfMask)
	}
}

const (
	RegisterMcrFieldSync_inShift = 8
	RegisterMcrFieldSync_inMask  = 0x300
)

// GetSync_in ynchronization input
func (r *registerMcrType) GetSync_in() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldSync_inMask) >> RegisterMcrFieldSync_inShift)
}

// SetSync_in ynchronization input
func (r *registerMcrType) SetSync_in(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldSync_inMask)|(uint32(value)<<RegisterMcrFieldSync_inShift))
}

const (
	RegisterMcrFieldSyncrstmShift = 10
	RegisterMcrFieldSyncrstmMask  = 0x400
)

// GetSyncrstm Synchronization Resets Master
func (r *registerMcrType) GetSyncrstm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldSyncrstmMask) != 0
}

// SetSyncrstm Synchronization Resets Master
func (r *registerMcrType) SetSyncrstm(value bool) {
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
func (r *registerMcrType) GetSyncstrtm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldSyncstrtmMask) != 0
}

// SetSyncstrtm Synchronization Starts Master
func (r *registerMcrType) SetSyncstrtm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMcrFieldSyncstrtmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldSyncstrtmMask)
	}
}

const (
	RegisterMcrFieldSync_outShift = 12
	RegisterMcrFieldSync_outMask  = 0x3000
)

// GetSync_out Synchronization output
func (r *registerMcrType) GetSync_out() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldSync_outMask) >> RegisterMcrFieldSync_outShift)
}

// SetSync_out Synchronization output
func (r *registerMcrType) SetSync_out(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldSync_outMask)|(uint32(value)<<RegisterMcrFieldSync_outShift))
}

const (
	RegisterMcrFieldSync_srcShift = 14
	RegisterMcrFieldSync_srcMask  = 0xc000
)

// GetSync_src Synchronization source
func (r *registerMcrType) GetSync_src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldSync_srcMask) >> RegisterMcrFieldSync_srcShift)
}

// SetSync_src Synchronization source
func (r *registerMcrType) SetSync_src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldSync_srcMask)|(uint32(value)<<RegisterMcrFieldSync_srcShift))
}

const (
	RegisterMcrFieldMcenShift = 16
	RegisterMcrFieldMcenMask  = 0x10000
)

// GetMcen Master Counter enable
func (r *registerMcrType) GetMcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldMcenMask) != 0
}

// SetMcen Master Counter enable
func (r *registerMcrType) SetMcen(value bool) {
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
func (r *registerMcrType) GetTacen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldTacenMask) != 0
}

// SetTacen Timer A counter enable
func (r *registerMcrType) SetTacen(value bool) {
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
func (r *registerMcrType) GetTbcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldTbcenMask) != 0
}

// SetTbcen Timer B counter enable
func (r *registerMcrType) SetTbcen(value bool) {
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
func (r *registerMcrType) GetTccen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldTccenMask) != 0
}

// SetTccen Timer C counter enable
func (r *registerMcrType) SetTccen(value bool) {
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
func (r *registerMcrType) GetTdcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldTdcenMask) != 0
}

// SetTdcen Timer D counter enable
func (r *registerMcrType) SetTdcen(value bool) {
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
func (r *registerMcrType) GetTecen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldTecenMask) != 0
}

// SetTecen Timer E counter enable
func (r *registerMcrType) SetTecen(value bool) {
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
func (r *registerMcrType) GetDacsync() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldDacsyncMask) >> RegisterMcrFieldDacsyncShift)
}

// SetDacsync AC Synchronization
func (r *registerMcrType) SetDacsync(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldDacsyncMask)|(uint32(value)<<RegisterMcrFieldDacsyncShift))
}

const (
	RegisterMcrFieldPreenShift = 27
	RegisterMcrFieldPreenMask  = 0x8000000
)

// GetPreen Preload enable
func (r *registerMcrType) GetPreen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldPreenMask) != 0
}

// SetPreen Preload enable
func (r *registerMcrType) SetPreen(value bool) {
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
func (r *registerMcrType) GetMrepu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldMrepuMask) != 0
}

// SetMrepu Master Timer Repetition update
func (r *registerMcrType) SetMrepu(value bool) {
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
func (r *registerMcrType) GetBrstdma() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMcrFieldBrstdmaMask) >> RegisterMcrFieldBrstdmaShift)
}

// SetBrstdma Burst DMA Update
func (r *registerMcrType) SetBrstdma(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcrFieldBrstdmaMask)|(uint32(value)<<RegisterMcrFieldBrstdmaShift))
}

// registerMisrType Master Timer Interrupt Status Register
type registerMisrType uint32

const (
	RegisterMisrFieldMcmp1Shift = 0
	RegisterMisrFieldMcmp1Mask  = 0x1
)

// GetMcmp1 Master Compare 1 Interrupt Flag
func (r *registerMisrType) GetMcmp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMcmp1Mask) != 0
}

// SetMcmp1 Master Compare 1 Interrupt Flag
func (r *registerMisrType) SetMcmp1(value bool) {
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
func (r *registerMisrType) GetMcmp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMcmp2Mask) != 0
}

// SetMcmp2 Master Compare 2 Interrupt Flag
func (r *registerMisrType) SetMcmp2(value bool) {
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
func (r *registerMisrType) GetMcmp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMcmp3Mask) != 0
}

// SetMcmp3 Master Compare 3 Interrupt Flag
func (r *registerMisrType) SetMcmp3(value bool) {
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
func (r *registerMisrType) GetMcmp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMcmp4Mask) != 0
}

// SetMcmp4 Master Compare 4 Interrupt Flag
func (r *registerMisrType) SetMcmp4(value bool) {
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
func (r *registerMisrType) GetMrep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMrepMask) != 0
}

// SetMrep Master Repetition Interrupt Flag
func (r *registerMisrType) SetMrep(value bool) {
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
func (r *registerMisrType) GetSync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldSyncMask) != 0
}

// SetSync Sync Input Interrupt Flag
func (r *registerMisrType) SetSync(value bool) {
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
func (r *registerMisrType) GetMupd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMisrFieldMupdMask) != 0
}

// SetMupd Master Update Interrupt Flag
func (r *registerMisrType) SetMupd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMisrFieldMupdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMisrFieldMupdMask)
	}
}

// registerMicrType Master Timer Interrupt Clear Register
type registerMicrType uint32

const (
	RegisterMicrFieldMcmp1cShift = 0
	RegisterMicrFieldMcmp1cMask  = 0x1
)

// GetMcmp1c Master Compare 1 Interrupt flag clear
func (r *registerMicrType) GetMcmp1c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMcmp1cMask) != 0
}

// SetMcmp1c Master Compare 1 Interrupt flag clear
func (r *registerMicrType) SetMcmp1c(value bool) {
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
func (r *registerMicrType) GetMcmp2c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMcmp2cMask) != 0
}

// SetMcmp2c Master Compare 2 Interrupt flag clear
func (r *registerMicrType) SetMcmp2c(value bool) {
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
func (r *registerMicrType) GetMcmp3c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMcmp3cMask) != 0
}

// SetMcmp3c Master Compare 3 Interrupt flag clear
func (r *registerMicrType) SetMcmp3c(value bool) {
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
func (r *registerMicrType) GetMcmp4c() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMcmp4cMask) != 0
}

// SetMcmp4c Master Compare 4 Interrupt flag clear
func (r *registerMicrType) SetMcmp4c(value bool) {
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
func (r *registerMicrType) GetMrepc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMrepcMask) != 0
}

// SetMrepc Repetition Interrupt flag clear
func (r *registerMicrType) SetMrepc(value bool) {
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
func (r *registerMicrType) GetSyncc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldSynccMask) != 0
}

// SetSyncc Sync Input Interrupt flag clear
func (r *registerMicrType) SetSyncc(value bool) {
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
func (r *registerMicrType) GetMupdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMicrFieldMupdcMask) != 0
}

// SetMupdc Master update Interrupt flag clear
func (r *registerMicrType) SetMupdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMicrFieldMupdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMicrFieldMupdcMask)
	}
}

// registerMdier4Type MDIER4
type registerMdier4Type uint32

const (
	RegisterMdier4FieldMcmp1ieShift = 0
	RegisterMdier4FieldMcmp1ieMask  = 0x1
)

// GetMcmp1ie MCMP1IE
func (r *registerMdier4Type) GetMcmp1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp1ieMask) != 0
}

// SetMcmp1ie MCMP1IE
func (r *registerMdier4Type) SetMcmp1ie(value bool) {
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
func (r *registerMdier4Type) GetMcmp2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp2ieMask) != 0
}

// SetMcmp2ie MCMP2IE
func (r *registerMdier4Type) SetMcmp2ie(value bool) {
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
func (r *registerMdier4Type) GetMcmp3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp3ieMask) != 0
}

// SetMcmp3ie MCMP3IE
func (r *registerMdier4Type) SetMcmp3ie(value bool) {
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
func (r *registerMdier4Type) GetMcmp4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp4ieMask) != 0
}

// SetMcmp4ie MCMP4IE
func (r *registerMdier4Type) SetMcmp4ie(value bool) {
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
func (r *registerMdier4Type) GetMrepie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMrepieMask) != 0
}

// SetMrepie MREPIE
func (r *registerMdier4Type) SetMrepie(value bool) {
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
func (r *registerMdier4Type) GetSyncie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldSyncieMask) != 0
}

// SetSyncie SYNCIE
func (r *registerMdier4Type) SetSyncie(value bool) {
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
func (r *registerMdier4Type) GetMupdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMupdieMask) != 0
}

// SetMupdie MUPDIE
func (r *registerMdier4Type) SetMupdie(value bool) {
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
func (r *registerMdier4Type) GetMcmp1de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp1deMask) != 0
}

// SetMcmp1de MCMP1DE
func (r *registerMdier4Type) SetMcmp1de(value bool) {
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
func (r *registerMdier4Type) GetMcmp2de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp2deMask) != 0
}

// SetMcmp2de MCMP2DE
func (r *registerMdier4Type) SetMcmp2de(value bool) {
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
func (r *registerMdier4Type) GetMcmp3de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp3deMask) != 0
}

// SetMcmp3de MCMP3DE
func (r *registerMdier4Type) SetMcmp3de(value bool) {
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
func (r *registerMdier4Type) GetMcmp4de() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMcmp4deMask) != 0
}

// SetMcmp4de MCMP4DE
func (r *registerMdier4Type) SetMcmp4de(value bool) {
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
func (r *registerMdier4Type) GetMrepde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMrepdeMask) != 0
}

// SetMrepde MREPDE
func (r *registerMdier4Type) SetMrepde(value bool) {
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
func (r *registerMdier4Type) GetSyncde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldSyncdeMask) != 0
}

// SetSyncde SYNCDE
func (r *registerMdier4Type) SetSyncde(value bool) {
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
func (r *registerMdier4Type) GetMupdde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMdier4FieldMupddeMask) != 0
}

// SetMupdde MUPDDE
func (r *registerMdier4Type) SetMupdde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMdier4FieldMupddeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMdier4FieldMupddeMask)
	}
}

// registerMcntrType Master Timer Counter Register
type registerMcntrType uint32

const (
	RegisterMcntrFieldMcntShift = 0
	RegisterMcntrFieldMcntMask  = 0xffff
)

// GetMcnt Counter value
func (r *registerMcntrType) GetMcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMcntrFieldMcntMask) >> RegisterMcntrFieldMcntShift)
}

// SetMcnt Counter value
func (r *registerMcntrType) SetMcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcntrFieldMcntMask)|(uint32(value)<<RegisterMcntrFieldMcntShift))
}

// registerMperType Master Timer Period Register
type registerMperType uint32

const (
	RegisterMperFieldMperShift = 0
	RegisterMperFieldMperMask  = 0xffff
)

// GetMper Master Timer Period value
func (r *registerMperType) GetMper() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMperFieldMperMask) >> RegisterMperFieldMperShift)
}

// SetMper Master Timer Period value
func (r *registerMperType) SetMper(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMperFieldMperMask)|(uint32(value)<<RegisterMperFieldMperShift))
}

// registerMrepType Master Timer Repetition Register
type registerMrepType uint32

const (
	RegisterMrepFieldMrepShift = 0
	RegisterMrepFieldMrepMask  = 0xff
)

// GetMrep Master Timer Repetition counter value
func (r *registerMrepType) GetMrep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMrepFieldMrepMask) >> RegisterMrepFieldMrepShift)
}

// SetMrep Master Timer Repetition counter value
func (r *registerMrepType) SetMrep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMrepFieldMrepMask)|(uint32(value)<<RegisterMrepFieldMrepShift))
}

// registerMcmp1rType Master Timer Compare 1 Register
type registerMcmp1rType uint32

const (
	RegisterMcmp1rFieldMcmp1Shift = 0
	RegisterMcmp1rFieldMcmp1Mask  = 0xffff
)

// GetMcmp1 Master Timer Compare 1 value
func (r *registerMcmp1rType) GetMcmp1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMcmp1rFieldMcmp1Mask) >> RegisterMcmp1rFieldMcmp1Shift)
}

// SetMcmp1 Master Timer Compare 1 value
func (r *registerMcmp1rType) SetMcmp1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcmp1rFieldMcmp1Mask)|(uint32(value)<<RegisterMcmp1rFieldMcmp1Shift))
}

// registerMcmp2rType Master Timer Compare 2 Register
type registerMcmp2rType uint32

const (
	RegisterMcmp2rFieldMcmp2Shift = 0
	RegisterMcmp2rFieldMcmp2Mask  = 0xffff
)

// GetMcmp2 Master Timer Compare 2 value
func (r *registerMcmp2rType) GetMcmp2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMcmp2rFieldMcmp2Mask) >> RegisterMcmp2rFieldMcmp2Shift)
}

// SetMcmp2 Master Timer Compare 2 value
func (r *registerMcmp2rType) SetMcmp2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcmp2rFieldMcmp2Mask)|(uint32(value)<<RegisterMcmp2rFieldMcmp2Shift))
}

// registerMcmp3rType Master Timer Compare 3 Register
type registerMcmp3rType uint32

const (
	RegisterMcmp3rFieldMcmp3Shift = 0
	RegisterMcmp3rFieldMcmp3Mask  = 0xffff
)

// GetMcmp3 Master Timer Compare 3 value
func (r *registerMcmp3rType) GetMcmp3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMcmp3rFieldMcmp3Mask) >> RegisterMcmp3rFieldMcmp3Shift)
}

// SetMcmp3 Master Timer Compare 3 value
func (r *registerMcmp3rType) SetMcmp3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcmp3rFieldMcmp3Mask)|(uint32(value)<<RegisterMcmp3rFieldMcmp3Shift))
}

// registerMcmp4rType Master Timer Compare 4 Register
type registerMcmp4rType uint32

const (
	RegisterMcmp4rFieldMcmp4Shift = 0
	RegisterMcmp4rFieldMcmp4Mask  = 0xffff
)

// GetMcmp4 Master Timer Compare 4 value
func (r *registerMcmp4rType) GetMcmp4() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMcmp4rFieldMcmp4Mask) >> RegisterMcmp4rFieldMcmp4Shift)
}

// SetMcmp4 Master Timer Compare 4 value
func (r *registerMcmp4rType) SetMcmp4(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMcmp4rFieldMcmp4Mask)|(uint32(value)<<RegisterMcmp4rFieldMcmp4Shift))
}
