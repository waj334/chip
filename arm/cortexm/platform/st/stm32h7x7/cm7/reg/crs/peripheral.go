//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package crs

import (
	"unsafe"
	"volatile"
)

var (
	Crs = (*_crs)(unsafe.Pointer(uintptr(0x40008400)))
)

type _crs struct {
	Cr   RegisterCrType
	Cfgr RegisterCfgrType
	Isr  RegisterIsrType
	Icr  RegisterIcrType
}

// RegisterCrType CRS control register
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
	RegisterCrFieldSyncokieShift = 0
	RegisterCrFieldSyncokieMask  = 0x1
)

// GetSyncokie SYNC event OK interrupt enable
func (r *RegisterCrType) GetSyncokie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSyncokieMask) != 0
}

// SetSyncokie SYNC event OK interrupt enable
func (r *RegisterCrType) SetSyncokie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSyncokieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSyncokieMask)
	}
}

const (
	RegisterCrFieldSyncwarnieShift = 1
	RegisterCrFieldSyncwarnieMask  = 0x2
)

// GetSyncwarnie SYNC warning interrupt enable
func (r *RegisterCrType) GetSyncwarnie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSyncwarnieMask) != 0
}

// SetSyncwarnie SYNC warning interrupt enable
func (r *RegisterCrType) SetSyncwarnie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSyncwarnieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSyncwarnieMask)
	}
}

const (
	RegisterCrFieldErrieShift = 2
	RegisterCrFieldErrieMask  = 0x4
)

// GetErrie Synchronization or trimming error interrupt enable
func (r *RegisterCrType) GetErrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldErrieMask) != 0
}

// SetErrie Synchronization or trimming error interrupt enable
func (r *RegisterCrType) SetErrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldErrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldErrieMask)
	}
}

const (
	RegisterCrFieldEsyncieShift = 3
	RegisterCrFieldEsyncieMask  = 0x8
)

// GetEsyncie Expected SYNC interrupt enable
func (r *RegisterCrType) GetEsyncie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEsyncieMask) != 0
}

// SetEsyncie Expected SYNC interrupt enable
func (r *RegisterCrType) SetEsyncie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEsyncieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEsyncieMask)
	}
}

const (
	RegisterCrFieldCenShift = 5
	RegisterCrFieldCenMask  = 0x20
)

// GetCen Frequency error counter enable This bit enables the oscillator clock for the frequency error counter. When this bit is set, the CRS_CFGR register is write-protected and cannot be modified.
func (r *RegisterCrType) GetCen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCenMask) != 0
}

// SetCen Frequency error counter enable This bit enables the oscillator clock for the frequency error counter. When this bit is set, the CRS_CFGR register is write-protected and cannot be modified.
func (r *RegisterCrType) SetCen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCenMask)
	}
}

const (
	RegisterCrFieldAutotrimenShift = 6
	RegisterCrFieldAutotrimenMask  = 0x40
)

// GetAutotrimen Automatic trimming enable This bit enables the automatic hardware adjustment of TRIM bits according to the measured frequency error between two SYNC events. If this bit is set, the TRIM bits are read-only. The TRIM value can be adjusted by hardware by one or two steps at a time, depending on the measured frequency error value. Refer to Section7.3.4: Frequency error evaluation and automatic trimming for more details.
func (r *RegisterCrType) GetAutotrimen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldAutotrimenMask) != 0
}

// SetAutotrimen Automatic trimming enable This bit enables the automatic hardware adjustment of TRIM bits according to the measured frequency error between two SYNC events. If this bit is set, the TRIM bits are read-only. The TRIM value can be adjusted by hardware by one or two steps at a time, depending on the measured frequency error value. Refer to Section7.3.4: Frequency error evaluation and automatic trimming for more details.
func (r *RegisterCrType) SetAutotrimen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldAutotrimenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldAutotrimenMask)
	}
}

const (
	RegisterCrFieldSwsyncShift = 7
	RegisterCrFieldSwsyncMask  = 0x80
)

// GetSwsync Generate software SYNC event This bit is set by software in order to generate a software SYNC event. It is automatically cleared by hardware.
func (r *RegisterCrType) GetSwsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSwsyncMask) != 0
}

const (
	RegisterCrFieldTrimShift = 8
	RegisterCrFieldTrimMask  = 0x3f00
)

// GetTrim HSI48 oscillator smooth trimming These bits provide a user-programmable trimming value to the HSI48 oscillator. They can be programmed to adjust to variations in voltage and temperature that influence the frequency of the HSI48. The default value is 32, which corresponds to the middle of the trimming interval. The trimming step is around 67 kHz between two consecutive TRIM steps. A higher TRIM value corresponds to a higher output frequency. When the AUTOTRIMEN bit is set, this field is controlled by hardware and is read-only.
func (r *RegisterCrType) GetTrim() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTrimMask) >> RegisterCrFieldTrimShift)
}

// SetTrim HSI48 oscillator smooth trimming These bits provide a user-programmable trimming value to the HSI48 oscillator. They can be programmed to adjust to variations in voltage and temperature that influence the frequency of the HSI48. The default value is 32, which corresponds to the middle of the trimming interval. The trimming step is around 67 kHz between two consecutive TRIM steps. A higher TRIM value corresponds to a higher output frequency. When the AUTOTRIMEN bit is set, this field is controlled by hardware and is read-only.
func (r *RegisterCrType) SetTrim(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTrimMask)|(uint32(value)<<RegisterCrFieldTrimShift))
}

// RegisterCfgrType This register can be written only when the frequency error counter is disabled (CEN bit is cleared in CRS_CR). When the counter is enabled, this register is write-protected.
type RegisterCfgrType uint32

func (r *RegisterCfgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCfgrFieldReloadShift = 0
	RegisterCfgrFieldReloadMask  = 0xffff
)

// GetReload Counter reload value RELOAD is the value to be loaded in the frequency error counter with each SYNC event. Refer to Section7.3.3: Frequency error measurement for more details about counter behavior.
func (r *RegisterCfgrType) GetReload() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldReloadMask) >> RegisterCfgrFieldReloadShift)
}

// SetReload Counter reload value RELOAD is the value to be loaded in the frequency error counter with each SYNC event. Refer to Section7.3.3: Frequency error measurement for more details about counter behavior.
func (r *RegisterCfgrType) SetReload(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldReloadMask)|(uint32(value)<<RegisterCfgrFieldReloadShift))
}

const (
	RegisterCfgrFieldFelimShift = 16
	RegisterCfgrFieldFelimMask  = 0xff0000
)

// GetFelim Frequency error limit FELIM contains the value to be used to evaluate the captured frequency error value latched in the FECAP[15:0] bits of the CRS_ISR register. Refer to Section7.3.4: Frequency error evaluation and automatic trimming for more details about FECAP evaluation.
func (r *RegisterCfgrType) GetFelim() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldFelimMask) >> RegisterCfgrFieldFelimShift)
}

// SetFelim Frequency error limit FELIM contains the value to be used to evaluate the captured frequency error value latched in the FECAP[15:0] bits of the CRS_ISR register. Refer to Section7.3.4: Frequency error evaluation and automatic trimming for more details about FECAP evaluation.
func (r *RegisterCfgrType) SetFelim(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldFelimMask)|(uint32(value)<<RegisterCfgrFieldFelimShift))
}

const (
	RegisterCfgrFieldSyncdivShift = 24
	RegisterCfgrFieldSyncdivMask  = 0x7000000
)

// GetSyncdiv SYNC divider These bits are set and cleared by software to control the division factor of the SYNC signal.
func (r *RegisterCfgrType) GetSyncdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSyncdivMask) >> RegisterCfgrFieldSyncdivShift)
}

// SetSyncdiv SYNC divider These bits are set and cleared by software to control the division factor of the SYNC signal.
func (r *RegisterCfgrType) SetSyncdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSyncdivMask)|(uint32(value)<<RegisterCfgrFieldSyncdivShift))
}

const (
	RegisterCfgrFieldSyncsrcShift = 28
	RegisterCfgrFieldSyncsrcMask  = 0x30000000
)

// GetSyncsrc SYNC signal source selection These bits are set and cleared by software to select the SYNC signal source. Note: When using USB LPM (Link Power Management) and the device is in Sleep mode, the periodic USB SOF will not be generated by the host. No SYNC signal will therefore be provided to the CRS to calibrate the HSI48 on the run. To guarantee the required clock precision after waking up from Sleep mode, the LSE or reference clock on the GPIOs should be used as SYNC signal.
func (r *RegisterCfgrType) GetSyncsrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSyncsrcMask) >> RegisterCfgrFieldSyncsrcShift)
}

// SetSyncsrc SYNC signal source selection These bits are set and cleared by software to select the SYNC signal source. Note: When using USB LPM (Link Power Management) and the device is in Sleep mode, the periodic USB SOF will not be generated by the host. No SYNC signal will therefore be provided to the CRS to calibrate the HSI48 on the run. To guarantee the required clock precision after waking up from Sleep mode, the LSE or reference clock on the GPIOs should be used as SYNC signal.
func (r *RegisterCfgrType) SetSyncsrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSyncsrcMask)|(uint32(value)<<RegisterCfgrFieldSyncsrcShift))
}

const (
	RegisterCfgrFieldSyncpolShift = 31
	RegisterCfgrFieldSyncpolMask  = 0x80000000
)

// GetSyncpol SYNC polarity selection This bit is set and cleared by software to select the input polarity for the SYNC signal source.
func (r *RegisterCfgrType) GetSyncpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSyncpolMask) != 0
}

// SetSyncpol SYNC polarity selection This bit is set and cleared by software to select the input polarity for the SYNC signal source.
func (r *RegisterCfgrType) SetSyncpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldSyncpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSyncpolMask)
	}
}

// RegisterIsrType CRS interrupt and status register
type RegisterIsrType uint32

func (r *RegisterIsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIsrFieldSyncokfShift = 0
	RegisterIsrFieldSyncokfMask  = 0x1
)

// GetSyncokf SYNC event OK flag This flag is set by hardware when the measured frequency error is smaller than FELIM * 3. This means that either no adjustment of the TRIM value is needed or that an adjustment by one trimming step is enough to compensate the frequency error. An interrupt is generated if the SYNCOKIE bit is set in the CRS_CR register. It is cleared by software by setting the SYNCOKC bit in the CRS_ICR register.
func (r *RegisterIsrType) GetSyncokf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldSyncokfMask) != 0
}

// SetSyncokf SYNC event OK flag This flag is set by hardware when the measured frequency error is smaller than FELIM * 3. This means that either no adjustment of the TRIM value is needed or that an adjustment by one trimming step is enough to compensate the frequency error. An interrupt is generated if the SYNCOKIE bit is set in the CRS_CR register. It is cleared by software by setting the SYNCOKC bit in the CRS_ICR register.
func (r *RegisterIsrType) SetSyncokf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldSyncokfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldSyncokfMask)
	}
}

const (
	RegisterIsrFieldSyncwarnfShift = 1
	RegisterIsrFieldSyncwarnfMask  = 0x2
)

// GetSyncwarnf SYNC warning flag This flag is set by hardware when the measured frequency error is greater than or equal to FELIM * 3, but smaller than FELIM * 128. This means that to compensate the frequency error, the TRIM value must be adjusted by two steps or more. An interrupt is generated if the SYNCWARNIE bit is set in the CRS_CR register. It is cleared by software by setting the SYNCWARNC bit in the CRS_ICR register.
func (r *RegisterIsrType) GetSyncwarnf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldSyncwarnfMask) != 0
}

// SetSyncwarnf SYNC warning flag This flag is set by hardware when the measured frequency error is greater than or equal to FELIM * 3, but smaller than FELIM * 128. This means that to compensate the frequency error, the TRIM value must be adjusted by two steps or more. An interrupt is generated if the SYNCWARNIE bit is set in the CRS_CR register. It is cleared by software by setting the SYNCWARNC bit in the CRS_ICR register.
func (r *RegisterIsrType) SetSyncwarnf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldSyncwarnfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldSyncwarnfMask)
	}
}

const (
	RegisterIsrFieldErrfShift = 2
	RegisterIsrFieldErrfMask  = 0x4
)

// GetErrf Error flag This flag is set by hardware in case of any synchronization or trimming error. It is the logical OR of the TRIMOVF, SYNCMISS and SYNCERR bits. An interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software in reaction to setting the ERRC bit in the CRS_ICR register, which clears the TRIMOVF, SYNCMISS and SYNCERR bits.
func (r *RegisterIsrType) GetErrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldErrfMask) != 0
}

// SetErrf Error flag This flag is set by hardware in case of any synchronization or trimming error. It is the logical OR of the TRIMOVF, SYNCMISS and SYNCERR bits. An interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software in reaction to setting the ERRC bit in the CRS_ICR register, which clears the TRIMOVF, SYNCMISS and SYNCERR bits.
func (r *RegisterIsrType) SetErrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldErrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldErrfMask)
	}
}

const (
	RegisterIsrFieldEsyncfShift = 3
	RegisterIsrFieldEsyncfMask  = 0x8
)

// GetEsyncf Expected SYNC flag This flag is set by hardware when the frequency error counter reached a zero value. An interrupt is generated if the ESYNCIE bit is set in the CRS_CR register. It is cleared by software by setting the ESYNCC bit in the CRS_ICR register.
func (r *RegisterIsrType) GetEsyncf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldEsyncfMask) != 0
}

// SetEsyncf Expected SYNC flag This flag is set by hardware when the frequency error counter reached a zero value. An interrupt is generated if the ESYNCIE bit is set in the CRS_CR register. It is cleared by software by setting the ESYNCC bit in the CRS_ICR register.
func (r *RegisterIsrType) SetEsyncf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldEsyncfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldEsyncfMask)
	}
}

const (
	RegisterIsrFieldSyncerrShift = 8
	RegisterIsrFieldSyncerrMask  = 0x100
)

// GetSyncerr SYNC error This flag is set by hardware when the SYNC pulse arrives before the ESYNC event and the measured frequency error is greater than or equal to FELIM * 128. This means that the frequency error is too big (internal frequency too low) to be compensated by adjusting the TRIM value, and that some other action should be taken. An interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software by setting the ERRC bit in the CRS_ICR register.
func (r *RegisterIsrType) GetSyncerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldSyncerrMask) != 0
}

// SetSyncerr SYNC error This flag is set by hardware when the SYNC pulse arrives before the ESYNC event and the measured frequency error is greater than or equal to FELIM * 128. This means that the frequency error is too big (internal frequency too low) to be compensated by adjusting the TRIM value, and that some other action should be taken. An interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software by setting the ERRC bit in the CRS_ICR register.
func (r *RegisterIsrType) SetSyncerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldSyncerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldSyncerrMask)
	}
}

const (
	RegisterIsrFieldSyncmissShift = 9
	RegisterIsrFieldSyncmissMask  = 0x200
)

// GetSyncmiss SYNC missed This flag is set by hardware when the frequency error counter reached value FELIM * 128 and no SYNC was detected, meaning either that a SYNC pulse was missed or that the frequency error is too big (internal frequency too high) to be compensated by adjusting the TRIM value, and that some other action should be taken. At this point, the frequency error counter is stopped (waiting for a next SYNC) and an interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software by setting the ERRC bit in the CRS_ICR register.
func (r *RegisterIsrType) GetSyncmiss() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldSyncmissMask) != 0
}

// SetSyncmiss SYNC missed This flag is set by hardware when the frequency error counter reached value FELIM * 128 and no SYNC was detected, meaning either that a SYNC pulse was missed or that the frequency error is too big (internal frequency too high) to be compensated by adjusting the TRIM value, and that some other action should be taken. At this point, the frequency error counter is stopped (waiting for a next SYNC) and an interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software by setting the ERRC bit in the CRS_ICR register.
func (r *RegisterIsrType) SetSyncmiss(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldSyncmissMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldSyncmissMask)
	}
}

const (
	RegisterIsrFieldTrimovfShift = 10
	RegisterIsrFieldTrimovfMask  = 0x400
)

// GetTrimovf Trimming overflow or underflow This flag is set by hardware when the automatic trimming tries to over- or under-flow the TRIM value. An interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software by setting the ERRC bit in the CRS_ICR register.
func (r *RegisterIsrType) GetTrimovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTrimovfMask) != 0
}

// SetTrimovf Trimming overflow or underflow This flag is set by hardware when the automatic trimming tries to over- or under-flow the TRIM value. An interrupt is generated if the ERRIE bit is set in the CRS_CR register. It is cleared by software by setting the ERRC bit in the CRS_ICR register.
func (r *RegisterIsrType) SetTrimovf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTrimovfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTrimovfMask)
	}
}

const (
	RegisterIsrFieldFedirShift = 15
	RegisterIsrFieldFedirMask  = 0x8000
)

// GetFedir Frequency error direction FEDIR is the counting direction of the frequency error counter latched in the time of the last SYNC event. It shows whether the actual frequency is below or above the target.
func (r *RegisterIsrType) GetFedir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldFedirMask) != 0
}

// SetFedir Frequency error direction FEDIR is the counting direction of the frequency error counter latched in the time of the last SYNC event. It shows whether the actual frequency is below or above the target.
func (r *RegisterIsrType) SetFedir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldFedirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldFedirMask)
	}
}

const (
	RegisterIsrFieldFecapShift = 16
	RegisterIsrFieldFecapMask  = 0xffff0000
)

// GetFecap Frequency error capture FECAP is the frequency error counter value latched in the time of the last SYNC event. Refer to Section7.3.4: Frequency error evaluation and automatic trimming for more details about FECAP usage.
func (r *RegisterIsrType) GetFecap() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldFecapMask) >> RegisterIsrFieldFecapShift)
}

// SetFecap Frequency error capture FECAP is the frequency error counter value latched in the time of the last SYNC event. Refer to Section7.3.4: Frequency error evaluation and automatic trimming for more details about FECAP usage.
func (r *RegisterIsrType) SetFecap(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldFecapMask)|(uint32(value)<<RegisterIsrFieldFecapShift))
}

// RegisterIcrType CRS interrupt flag clear register
type RegisterIcrType uint32

func (r *RegisterIcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIcrFieldSyncokcShift = 0
	RegisterIcrFieldSyncokcMask  = 0x1
)

// GetSyncokc SYNC event OK clear flag Writing 1 to this bit clears the SYNCOKF flag in the CRS_ISR register.
func (r *RegisterIcrType) GetSyncokc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldSyncokcMask) != 0
}

// SetSyncokc SYNC event OK clear flag Writing 1 to this bit clears the SYNCOKF flag in the CRS_ISR register.
func (r *RegisterIcrType) SetSyncokc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldSyncokcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldSyncokcMask)
	}
}

const (
	RegisterIcrFieldSyncwarncShift = 1
	RegisterIcrFieldSyncwarncMask  = 0x2
)

// GetSyncwarnc SYNC warning clear flag Writing 1 to this bit clears the SYNCWARNF flag in the CRS_ISR register.
func (r *RegisterIcrType) GetSyncwarnc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldSyncwarncMask) != 0
}

// SetSyncwarnc SYNC warning clear flag Writing 1 to this bit clears the SYNCWARNF flag in the CRS_ISR register.
func (r *RegisterIcrType) SetSyncwarnc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldSyncwarncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldSyncwarncMask)
	}
}

const (
	RegisterIcrFieldErrcShift = 2
	RegisterIcrFieldErrcMask  = 0x4
)

// GetErrc Error clear flag Writing 1 to this bit clears TRIMOVF, SYNCMISS and SYNCERR bits and consequently also the ERRF flag in the CRS_ISR register.
func (r *RegisterIcrType) GetErrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldErrcMask) != 0
}

// SetErrc Error clear flag Writing 1 to this bit clears TRIMOVF, SYNCMISS and SYNCERR bits and consequently also the ERRF flag in the CRS_ISR register.
func (r *RegisterIcrType) SetErrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldErrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldErrcMask)
	}
}

const (
	RegisterIcrFieldEsynccShift = 3
	RegisterIcrFieldEsynccMask  = 0x8
)

// GetEsyncc Expected SYNC clear flag Writing 1 to this bit clears the ESYNCF flag in the CRS_ISR register.
func (r *RegisterIcrType) GetEsyncc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldEsynccMask) != 0
}

// SetEsyncc Expected SYNC clear flag Writing 1 to this bit clears the ESYNCF flag in the CRS_ISR register.
func (r *RegisterIcrType) SetEsyncc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldEsynccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldEsynccMask)
	}
}
