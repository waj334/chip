package adc_common

import (
	"unsafe"
	"volatile"
)

var (
	Adc3_common  = (*_adc_common)(unsafe.Pointer(uintptr(0x58026300)))
	Adc12_common = (*_adc_common)(unsafe.Pointer(uintptr(0x40022300)))

	Instances = [2]*_adc_common{
		Adc12_common,
		Adc3_common,
	}
)

type _adc_common struct {
	Csr  registerCsrType
	_    [4]byte
	Ccr  registerCcrType
	Cdr  registerCdrType
	Cdr2 registerCdr2Type
}

// registerCsrType ADC Common status register
type registerCsrType uint32

const (
	RegisterCsrFieldAdrdy_mstShift = 0
	RegisterCsrFieldAdrdy_mstMask  = 0x1
)

// GetAdrdy_mst Master ADC ready
func (r *registerCsrType) GetAdrdy_mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAdrdy_mstMask) != 0
}

// SetAdrdy_mst Master ADC ready
func (r *registerCsrType) SetAdrdy_mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAdrdy_mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAdrdy_mstMask)
	}
}

const (
	RegisterCsrFieldEosmp_mstShift = 1
	RegisterCsrFieldEosmp_mstMask  = 0x2
)

// GetEosmp_mst End of Sampling phase flag of the master ADC
func (r *registerCsrType) GetEosmp_mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEosmp_mstMask) != 0
}

// SetEosmp_mst End of Sampling phase flag of the master ADC
func (r *registerCsrType) SetEosmp_mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEosmp_mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEosmp_mstMask)
	}
}

const (
	RegisterCsrFieldEoc_mstShift = 2
	RegisterCsrFieldEoc_mstMask  = 0x4
)

// GetEoc_mst End of regular conversion of the master ADC
func (r *registerCsrType) GetEoc_mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEoc_mstMask) != 0
}

// SetEoc_mst End of regular conversion of the master ADC
func (r *registerCsrType) SetEoc_mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEoc_mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEoc_mstMask)
	}
}

const (
	RegisterCsrFieldEos_mstShift = 3
	RegisterCsrFieldEos_mstMask  = 0x8
)

// GetEos_mst End of regular sequence flag of the master ADC
func (r *registerCsrType) GetEos_mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEos_mstMask) != 0
}

// SetEos_mst End of regular sequence flag of the master ADC
func (r *registerCsrType) SetEos_mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEos_mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEos_mstMask)
	}
}

const (
	RegisterCsrFieldOvr_mstShift = 4
	RegisterCsrFieldOvr_mstMask  = 0x10
)

// GetOvr_mst Overrun flag of the master ADC
func (r *registerCsrType) GetOvr_mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldOvr_mstMask) != 0
}

// SetOvr_mst Overrun flag of the master ADC
func (r *registerCsrType) SetOvr_mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldOvr_mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldOvr_mstMask)
	}
}

const (
	RegisterCsrFieldJeoc_mstShift = 5
	RegisterCsrFieldJeoc_mstMask  = 0x20
)

// GetJeoc_mst End of injected conversion flag of the master ADC
func (r *registerCsrType) GetJeoc_mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJeoc_mstMask) != 0
}

// SetJeoc_mst End of injected conversion flag of the master ADC
func (r *registerCsrType) SetJeoc_mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJeoc_mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJeoc_mstMask)
	}
}

const (
	RegisterCsrFieldJeos_mstShift = 6
	RegisterCsrFieldJeos_mstMask  = 0x40
)

// GetJeos_mst End of injected sequence flag of the master ADC
func (r *registerCsrType) GetJeos_mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJeos_mstMask) != 0
}

// SetJeos_mst End of injected sequence flag of the master ADC
func (r *registerCsrType) SetJeos_mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJeos_mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJeos_mstMask)
	}
}

const (
	RegisterCsrFieldAwd1_mstShift = 7
	RegisterCsrFieldAwd1_mstMask  = 0x80
)

// GetAwd1_mst Analog watchdog 1 flag of the master ADC
func (r *registerCsrType) GetAwd1_mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd1_mstMask) != 0
}

// SetAwd1_mst Analog watchdog 1 flag of the master ADC
func (r *registerCsrType) SetAwd1_mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd1_mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd1_mstMask)
	}
}

const (
	RegisterCsrFieldAwd2_mstShift = 8
	RegisterCsrFieldAwd2_mstMask  = 0x100
)

// GetAwd2_mst Analog watchdog 2 flag of the master ADC
func (r *registerCsrType) GetAwd2_mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd2_mstMask) != 0
}

// SetAwd2_mst Analog watchdog 2 flag of the master ADC
func (r *registerCsrType) SetAwd2_mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd2_mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd2_mstMask)
	}
}

const (
	RegisterCsrFieldAwd3_mstShift = 9
	RegisterCsrFieldAwd3_mstMask  = 0x200
)

// GetAwd3_mst Analog watchdog 3 flag of the master ADC
func (r *registerCsrType) GetAwd3_mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd3_mstMask) != 0
}

// SetAwd3_mst Analog watchdog 3 flag of the master ADC
func (r *registerCsrType) SetAwd3_mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd3_mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd3_mstMask)
	}
}

const (
	RegisterCsrFieldJqovf_mstShift = 10
	RegisterCsrFieldJqovf_mstMask  = 0x400
)

// GetJqovf_mst Injected Context Queue Overflow flag of the master ADC
func (r *registerCsrType) GetJqovf_mst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJqovf_mstMask) != 0
}

// SetJqovf_mst Injected Context Queue Overflow flag of the master ADC
func (r *registerCsrType) SetJqovf_mst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJqovf_mstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJqovf_mstMask)
	}
}

const (
	RegisterCsrFieldAdrdy_slvShift = 16
	RegisterCsrFieldAdrdy_slvMask  = 0x10000
)

// GetAdrdy_slv Slave ADC ready
func (r *registerCsrType) GetAdrdy_slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAdrdy_slvMask) != 0
}

// SetAdrdy_slv Slave ADC ready
func (r *registerCsrType) SetAdrdy_slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAdrdy_slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAdrdy_slvMask)
	}
}

const (
	RegisterCsrFieldEosmp_slvShift = 17
	RegisterCsrFieldEosmp_slvMask  = 0x20000
)

// GetEosmp_slv End of Sampling phase flag of the slave ADC
func (r *registerCsrType) GetEosmp_slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEosmp_slvMask) != 0
}

// SetEosmp_slv End of Sampling phase flag of the slave ADC
func (r *registerCsrType) SetEosmp_slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEosmp_slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEosmp_slvMask)
	}
}

const (
	RegisterCsrFieldEoc_slvShift = 18
	RegisterCsrFieldEoc_slvMask  = 0x40000
)

// GetEoc_slv End of regular conversion of the slave ADC
func (r *registerCsrType) GetEoc_slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEoc_slvMask) != 0
}

// SetEoc_slv End of regular conversion of the slave ADC
func (r *registerCsrType) SetEoc_slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEoc_slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEoc_slvMask)
	}
}

const (
	RegisterCsrFieldEos_slvShift = 19
	RegisterCsrFieldEos_slvMask  = 0x80000
)

// GetEos_slv End of regular sequence flag of the slave ADC
func (r *registerCsrType) GetEos_slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldEos_slvMask) != 0
}

// SetEos_slv End of regular sequence flag of the slave ADC
func (r *registerCsrType) SetEos_slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldEos_slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldEos_slvMask)
	}
}

const (
	RegisterCsrFieldOvr_slvShift = 20
	RegisterCsrFieldOvr_slvMask  = 0x100000
)

// GetOvr_slv Overrun flag of the slave ADC
func (r *registerCsrType) GetOvr_slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldOvr_slvMask) != 0
}

// SetOvr_slv Overrun flag of the slave ADC
func (r *registerCsrType) SetOvr_slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldOvr_slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldOvr_slvMask)
	}
}

const (
	RegisterCsrFieldJeoc_slvShift = 21
	RegisterCsrFieldJeoc_slvMask  = 0x200000
)

// GetJeoc_slv End of injected conversion flag of the slave ADC
func (r *registerCsrType) GetJeoc_slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJeoc_slvMask) != 0
}

// SetJeoc_slv End of injected conversion flag of the slave ADC
func (r *registerCsrType) SetJeoc_slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJeoc_slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJeoc_slvMask)
	}
}

const (
	RegisterCsrFieldJeos_slvShift = 22
	RegisterCsrFieldJeos_slvMask  = 0x400000
)

// GetJeos_slv End of injected sequence flag of the slave ADC
func (r *registerCsrType) GetJeos_slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJeos_slvMask) != 0
}

// SetJeos_slv End of injected sequence flag of the slave ADC
func (r *registerCsrType) SetJeos_slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJeos_slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJeos_slvMask)
	}
}

const (
	RegisterCsrFieldAwd1_slvShift = 23
	RegisterCsrFieldAwd1_slvMask  = 0x800000
)

// GetAwd1_slv Analog watchdog 1 flag of the slave ADC
func (r *registerCsrType) GetAwd1_slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd1_slvMask) != 0
}

// SetAwd1_slv Analog watchdog 1 flag of the slave ADC
func (r *registerCsrType) SetAwd1_slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd1_slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd1_slvMask)
	}
}

const (
	RegisterCsrFieldAwd2_slvShift = 24
	RegisterCsrFieldAwd2_slvMask  = 0x1000000
)

// GetAwd2_slv Analog watchdog 2 flag of the slave ADC
func (r *registerCsrType) GetAwd2_slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd2_slvMask) != 0
}

// SetAwd2_slv Analog watchdog 2 flag of the slave ADC
func (r *registerCsrType) SetAwd2_slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd2_slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd2_slvMask)
	}
}

const (
	RegisterCsrFieldAwd3_slvShift = 25
	RegisterCsrFieldAwd3_slvMask  = 0x2000000
)

// GetAwd3_slv Analog watchdog 3 flag of the slave ADC
func (r *registerCsrType) GetAwd3_slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldAwd3_slvMask) != 0
}

// SetAwd3_slv Analog watchdog 3 flag of the slave ADC
func (r *registerCsrType) SetAwd3_slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldAwd3_slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldAwd3_slvMask)
	}
}

const (
	RegisterCsrFieldJqovf_slvShift = 26
	RegisterCsrFieldJqovf_slvMask  = 0x4000000
)

// GetJqovf_slv Injected Context Queue Overflow flag of the slave ADC
func (r *registerCsrType) GetJqovf_slv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldJqovf_slvMask) != 0
}

// SetJqovf_slv Injected Context Queue Overflow flag of the slave ADC
func (r *registerCsrType) SetJqovf_slv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldJqovf_slvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldJqovf_slvMask)
	}
}

// registerCcrType ADC common control register
type registerCcrType uint32

const (
	RegisterCcrFieldDualShift = 0
	RegisterCcrFieldDualMask  = 0x1f
)

// GetDual Dual ADC mode selection
func (r *registerCcrType) GetDual() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDualMask) >> RegisterCcrFieldDualShift)
}

// SetDual Dual ADC mode selection
func (r *registerCcrType) SetDual(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDualMask)|(uint32(value)<<RegisterCcrFieldDualShift))
}

const (
	RegisterCcrFieldDelayShift = 8
	RegisterCcrFieldDelayMask  = 0xf00
)

// GetDelay Delay between 2 sampling phases
func (r *registerCcrType) GetDelay() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDelayMask) >> RegisterCcrFieldDelayShift)
}

// SetDelay Delay between 2 sampling phases
func (r *registerCcrType) SetDelay(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDelayMask)|(uint32(value)<<RegisterCcrFieldDelayShift))
}

const (
	RegisterCcrFieldDamdfShift = 14
	RegisterCcrFieldDamdfMask  = 0xc000
)

// GetDamdf Dual ADC Mode Data Format
func (r *registerCcrType) GetDamdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldDamdfMask) >> RegisterCcrFieldDamdfShift)
}

// SetDamdf Dual ADC Mode Data Format
func (r *registerCcrType) SetDamdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldDamdfMask)|(uint32(value)<<RegisterCcrFieldDamdfShift))
}

const (
	RegisterCcrFieldCkmodeShift = 16
	RegisterCcrFieldCkmodeMask  = 0x30000
)

// GetCkmode ADC clock mode
func (r *registerCcrType) GetCkmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldCkmodeMask) >> RegisterCcrFieldCkmodeShift)
}

// SetCkmode ADC clock mode
func (r *registerCcrType) SetCkmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldCkmodeMask)|(uint32(value)<<RegisterCcrFieldCkmodeShift))
}

const (
	RegisterCcrFieldPrescShift = 18
	RegisterCcrFieldPrescMask  = 0x3c0000
)

// GetPresc ADC prescaler
func (r *registerCcrType) GetPresc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldPrescMask) >> RegisterCcrFieldPrescShift)
}

// SetPresc ADC prescaler
func (r *registerCcrType) SetPresc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldPrescMask)|(uint32(value)<<RegisterCcrFieldPrescShift))
}

const (
	RegisterCcrFieldVrefenShift = 22
	RegisterCcrFieldVrefenMask  = 0x400000
)

// GetVrefen VREFINT enable
func (r *registerCcrType) GetVrefen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldVrefenMask) != 0
}

// SetVrefen VREFINT enable
func (r *registerCcrType) SetVrefen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldVrefenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldVrefenMask)
	}
}

const (
	RegisterCcrFieldVsenseenShift = 23
	RegisterCcrFieldVsenseenMask  = 0x800000
)

// GetVsenseen Temperature sensor enable
func (r *registerCcrType) GetVsenseen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldVsenseenMask) != 0
}

// SetVsenseen Temperature sensor enable
func (r *registerCcrType) SetVsenseen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldVsenseenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldVsenseenMask)
	}
}

const (
	RegisterCcrFieldVbatenShift = 24
	RegisterCcrFieldVbatenMask  = 0x1000000
)

// GetVbaten VBAT enable
func (r *registerCcrType) GetVbaten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldVbatenMask) != 0
}

// SetVbaten VBAT enable
func (r *registerCcrType) SetVbaten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcrFieldVbatenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldVbatenMask)
	}
}

// registerCdrType ADC common regular data register for dual and triple modes
type registerCdrType uint32

const (
	RegisterCdrFieldRdata_mstShift = 0
	RegisterCdrFieldRdata_mstMask  = 0xffff
)

// GetRdata_mst Regular data of the master ADC
func (r *registerCdrType) GetRdata_mst() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCdrFieldRdata_mstMask) >> RegisterCdrFieldRdata_mstShift)
}

// SetRdata_mst Regular data of the master ADC
func (r *registerCdrType) SetRdata_mst(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCdrFieldRdata_mstMask)|(uint32(value)<<RegisterCdrFieldRdata_mstShift))
}

const (
	RegisterCdrFieldRdata_slvShift = 16
	RegisterCdrFieldRdata_slvMask  = 0xffff0000
)

// GetRdata_slv Regular data of the slave ADC
func (r *registerCdrType) GetRdata_slv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCdrFieldRdata_slvMask) >> RegisterCdrFieldRdata_slvShift)
}

// SetRdata_slv Regular data of the slave ADC
func (r *registerCdrType) SetRdata_slv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCdrFieldRdata_slvMask)|(uint32(value)<<RegisterCdrFieldRdata_slvShift))
}

// registerCdr2Type ADC x common regular data register for 32-bit dual mode
type registerCdr2Type uint32

const (
	RegisterCdr2FieldRdata_altShift = 0
	RegisterCdr2FieldRdata_altMask  = 0xffffffff
)

// GetRdata_alt Regular data of the master/slave alternated ADCs
func (r *registerCdr2Type) GetRdata_alt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCdr2FieldRdata_altMask) >> RegisterCdr2FieldRdata_altShift)
}

// SetRdata_alt Regular data of the master/slave alternated ADCs
func (r *registerCdr2Type) SetRdata_alt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCdr2FieldRdata_altMask)|(uint32(value)<<RegisterCdr2FieldRdata_altShift))
}
