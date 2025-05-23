package i2c1

import (
	"unsafe"
	"volatile"
)

var (
	I2c1 = (*_i2c1)(unsafe.Pointer(uintptr(0x40005400)))
	I2c2 = (*_i2c1)(unsafe.Pointer(uintptr(0x40005800)))
	I2c3 = (*_i2c1)(unsafe.Pointer(uintptr(0x40005c00)))
	I2c4 = (*_i2c1)(unsafe.Pointer(uintptr(0x58001c00)))
)

type _i2c1 struct {
	Cr1      registerCr1Type
	Cr2      registerCr2Type
	Oar1     registerOar1Type
	Oar2     registerOar2Type
	Timingr  registerTimingrType
	Timeoutr registerTimeoutrType
	Isr      registerIsrType
	Icr      registerIcrType
	Pecr     registerPecrType
	Rxdr     registerRxdrType
	Txdr     registerTxdrType
}

// registerCr1Type Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.
type registerCr1Type uint32

const (
	RegisterCr1FieldPeShift = 0
	RegisterCr1FieldPeMask  = 0x1
)

// GetPe Peripheral enable Note: When PE=0, the I2C SCL and SDA lines are released. Internal state machines and status bits are put back to their reset value. When cleared, PE must be kept low for at least 3 APB clock cycles.
func (r *registerCr1Type) GetPe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPeMask) != 0
}

// SetPe Peripheral enable Note: When PE=0, the I2C SCL and SDA lines are released. Internal state machines and status bits are put back to their reset value. When cleared, PE must be kept low for at least 3 APB clock cycles.
func (r *registerCr1Type) SetPe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldPeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldPeMask)
	}
}

const (
	RegisterCr1FieldTxieShift = 1
	RegisterCr1FieldTxieMask  = 0x2
)

// GetTxie TX Interrupt enable
func (r *registerCr1Type) GetTxie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTxieMask) != 0
}

// SetTxie TX Interrupt enable
func (r *registerCr1Type) SetTxie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTxieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTxieMask)
	}
}

const (
	RegisterCr1FieldRxieShift = 2
	RegisterCr1FieldRxieMask  = 0x4
)

// GetRxie RX Interrupt enable
func (r *registerCr1Type) GetRxie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRxieMask) != 0
}

// SetRxie RX Interrupt enable
func (r *registerCr1Type) SetRxie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldRxieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldRxieMask)
	}
}

const (
	RegisterCr1FieldAddrieShift = 3
	RegisterCr1FieldAddrieMask  = 0x8
)

// GetAddrie Address match Interrupt enable (slave only)
func (r *registerCr1Type) GetAddrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldAddrieMask) != 0
}

// SetAddrie Address match Interrupt enable (slave only)
func (r *registerCr1Type) SetAddrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldAddrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldAddrieMask)
	}
}

const (
	RegisterCr1FieldNackieShift = 4
	RegisterCr1FieldNackieMask  = 0x10
)

// GetNackie Not acknowledge received Interrupt enable
func (r *registerCr1Type) GetNackie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldNackieMask) != 0
}

// SetNackie Not acknowledge received Interrupt enable
func (r *registerCr1Type) SetNackie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldNackieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldNackieMask)
	}
}

const (
	RegisterCr1FieldStopieShift = 5
	RegisterCr1FieldStopieMask  = 0x20
)

// GetStopie STOP detection Interrupt enable
func (r *registerCr1Type) GetStopie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldStopieMask) != 0
}

// SetStopie STOP detection Interrupt enable
func (r *registerCr1Type) SetStopie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldStopieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldStopieMask)
	}
}

const (
	RegisterCr1FieldTcieShift = 6
	RegisterCr1FieldTcieMask  = 0x40
)

// GetTcie Transfer Complete interrupt enable Note: Any of these events will generate an interrupt: Transfer Complete (TC) Transfer Complete Reload (TCR)
func (r *registerCr1Type) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTcieMask) != 0
}

// SetTcie Transfer Complete interrupt enable Note: Any of these events will generate an interrupt: Transfer Complete (TC) Transfer Complete Reload (TCR)
func (r *registerCr1Type) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTcieMask)
	}
}

const (
	RegisterCr1FieldErrieShift = 7
	RegisterCr1FieldErrieMask  = 0x80
)

// GetErrie Error interrupts enable Note: Any of these errors generate an interrupt: Arbitration Loss (ARLO) Bus Error detection (BERR) Overrun/Underrun (OVR) Timeout detection (TIMEOUT) PEC error detection (PECERR) Alert pin event detection (ALERT)
func (r *registerCr1Type) GetErrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldErrieMask) != 0
}

// SetErrie Error interrupts enable Note: Any of these errors generate an interrupt: Arbitration Loss (ARLO) Bus Error detection (BERR) Overrun/Underrun (OVR) Timeout detection (TIMEOUT) PEC error detection (PECERR) Alert pin event detection (ALERT)
func (r *registerCr1Type) SetErrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldErrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldErrieMask)
	}
}

const (
	RegisterCr1FieldDnfShift = 8
	RegisterCr1FieldDnfMask  = 0xf00
)

// GetDnf Digital noise filter These bits are used to configure the digital noise filter on SDA and SCL input. The digital filter will filter spikes with a length of up to DNF[3:0] * tI2CCLK ... Note: If the analog filter is also enabled, the digital filter is added to the analog filter. This filter can only be programmed when the I2C is disabled (PE = 0).
func (r *registerCr1Type) GetDnf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDnfMask) >> RegisterCr1FieldDnfShift)
}

// SetDnf Digital noise filter These bits are used to configure the digital noise filter on SDA and SCL input. The digital filter will filter spikes with a length of up to DNF[3:0] * tI2CCLK ... Note: If the analog filter is also enabled, the digital filter is added to the analog filter. This filter can only be programmed when the I2C is disabled (PE = 0).
func (r *registerCr1Type) SetDnf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDnfMask)|(uint32(value)<<RegisterCr1FieldDnfShift))
}

const (
	RegisterCr1FieldAnfoffShift = 12
	RegisterCr1FieldAnfoffMask  = 0x1000
)

// GetAnfoff Analog noise filter OFF Note: This bit can only be programmed when the I2C is disabled (PE = 0).
func (r *registerCr1Type) GetAnfoff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldAnfoffMask) != 0
}

// SetAnfoff Analog noise filter OFF Note: This bit can only be programmed when the I2C is disabled (PE = 0).
func (r *registerCr1Type) SetAnfoff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldAnfoffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldAnfoffMask)
	}
}

const (
	RegisterCr1FieldTxdmaenShift = 14
	RegisterCr1FieldTxdmaenMask  = 0x4000
)

// GetTxdmaen DMA transmission requests enable
func (r *registerCr1Type) GetTxdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTxdmaenMask) != 0
}

// SetTxdmaen DMA transmission requests enable
func (r *registerCr1Type) SetTxdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTxdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTxdmaenMask)
	}
}

const (
	RegisterCr1FieldRxdmaenShift = 15
	RegisterCr1FieldRxdmaenMask  = 0x8000
)

// GetRxdmaen DMA reception requests enable
func (r *registerCr1Type) GetRxdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRxdmaenMask) != 0
}

// SetRxdmaen DMA reception requests enable
func (r *registerCr1Type) SetRxdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldRxdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldRxdmaenMask)
	}
}

const (
	RegisterCr1FieldSbcShift = 16
	RegisterCr1FieldSbcMask  = 0x10000
)

// GetSbc Slave byte control This bit is used to enable hardware byte control in slave mode.
func (r *registerCr1Type) GetSbc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldSbcMask) != 0
}

// SetSbc Slave byte control This bit is used to enable hardware byte control in slave mode.
func (r *registerCr1Type) SetSbc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldSbcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldSbcMask)
	}
}

const (
	RegisterCr1FieldNostretchShift = 17
	RegisterCr1FieldNostretchMask  = 0x20000
)

// GetNostretch Clock stretching disable This bit is used to disable clock stretching in slave mode. It must be kept cleared in master mode. Note: This bit can only be programmed when the I2C is disabled (PE = 0).
func (r *registerCr1Type) GetNostretch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldNostretchMask) != 0
}

// SetNostretch Clock stretching disable This bit is used to disable clock stretching in slave mode. It must be kept cleared in master mode. Note: This bit can only be programmed when the I2C is disabled (PE = 0).
func (r *registerCr1Type) SetNostretch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldNostretchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldNostretchMask)
	}
}

const (
	RegisterCr1FieldWupenShift = 18
	RegisterCr1FieldWupenMask  = 0x40000
)

// GetWupen Wakeup from Stop mode enable Note: If the Wakeup from Stop mode feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation. Note: WUPEN can be set only when DNF = 0000
func (r *registerCr1Type) GetWupen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldWupenMask) != 0
}

// SetWupen Wakeup from Stop mode enable Note: If the Wakeup from Stop mode feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation. Note: WUPEN can be set only when DNF = 0000
func (r *registerCr1Type) SetWupen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldWupenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldWupenMask)
	}
}

const (
	RegisterCr1FieldGcenShift = 19
	RegisterCr1FieldGcenMask  = 0x80000
)

// GetGcen General call enable
func (r *registerCr1Type) GetGcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldGcenMask) != 0
}

// SetGcen General call enable
func (r *registerCr1Type) SetGcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldGcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldGcenMask)
	}
}

const (
	RegisterCr1FieldSmbhenShift = 20
	RegisterCr1FieldSmbhenMask  = 0x100000
)

// GetSmbhen SMBus Host address enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerCr1Type) GetSmbhen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldSmbhenMask) != 0
}

// SetSmbhen SMBus Host address enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerCr1Type) SetSmbhen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldSmbhenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldSmbhenMask)
	}
}

const (
	RegisterCr1FieldSmbdenShift = 21
	RegisterCr1FieldSmbdenMask  = 0x200000
)

// GetSmbden SMBus Device Default address enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerCr1Type) GetSmbden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldSmbdenMask) != 0
}

// SetSmbden SMBus Device Default address enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerCr1Type) SetSmbden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldSmbdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldSmbdenMask)
	}
}

const (
	RegisterCr1FieldAlertenShift = 22
	RegisterCr1FieldAlertenMask  = 0x400000
)

// GetAlerten SMBus alert enable Device mode (SMBHEN=0): Host mode (SMBHEN=1): Note: When ALERTEN=0, the SMBA pin can be used as a standard GPIO. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerCr1Type) GetAlerten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldAlertenMask) != 0
}

// SetAlerten SMBus alert enable Device mode (SMBHEN=0): Host mode (SMBHEN=1): Note: When ALERTEN=0, the SMBA pin can be used as a standard GPIO. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerCr1Type) SetAlerten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldAlertenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldAlertenMask)
	}
}

const (
	RegisterCr1FieldPecenShift = 23
	RegisterCr1FieldPecenMask  = 0x800000
)

// GetPecen PEC enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerCr1Type) GetPecen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPecenMask) != 0
}

// SetPecen PEC enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerCr1Type) SetPecen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldPecenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldPecenMask)
	}
}

// registerCr2Type Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.
type registerCr2Type uint32

const (
	RegisterCr2FieldSadd0Shift = 0
	RegisterCr2FieldSadd0Mask  = 0x1
)

// GetSadd0 Slave address bit 0 (master mode) In 7-bit addressing mode (ADD10 = 0): This bit is dont care In 10-bit addressing mode (ADD10 = 1): This bit should be written with bit 0 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) GetSadd0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSadd0Mask) != 0
}

// SetSadd0 Slave address bit 0 (master mode) In 7-bit addressing mode (ADD10 = 0): This bit is dont care In 10-bit addressing mode (ADD10 = 1): This bit should be written with bit 0 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) SetSadd0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSadd0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSadd0Mask)
	}
}

const (
	RegisterCr2FieldSadd1Shift = 1
	RegisterCr2FieldSadd1Mask  = 0x2
)

// GetSadd1 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) GetSadd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSadd1Mask) != 0
}

// SetSadd1 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) SetSadd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSadd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSadd1Mask)
	}
}

const (
	RegisterCr2FieldSadd2Shift = 2
	RegisterCr2FieldSadd2Mask  = 0x4
)

// GetSadd2 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) GetSadd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSadd2Mask) != 0
}

// SetSadd2 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) SetSadd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSadd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSadd2Mask)
	}
}

const (
	RegisterCr2FieldSadd3Shift = 3
	RegisterCr2FieldSadd3Mask  = 0x8
)

// GetSadd3 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) GetSadd3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSadd3Mask) != 0
}

// SetSadd3 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) SetSadd3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSadd3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSadd3Mask)
	}
}

const (
	RegisterCr2FieldSadd4Shift = 4
	RegisterCr2FieldSadd4Mask  = 0x10
)

// GetSadd4 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) GetSadd4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSadd4Mask) != 0
}

// SetSadd4 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) SetSadd4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSadd4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSadd4Mask)
	}
}

const (
	RegisterCr2FieldSadd5Shift = 5
	RegisterCr2FieldSadd5Mask  = 0x20
)

// GetSadd5 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) GetSadd5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSadd5Mask) != 0
}

// SetSadd5 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) SetSadd5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSadd5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSadd5Mask)
	}
}

const (
	RegisterCr2FieldSadd6Shift = 6
	RegisterCr2FieldSadd6Mask  = 0x40
)

// GetSadd6 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) GetSadd6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSadd6Mask) != 0
}

// SetSadd6 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) SetSadd6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSadd6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSadd6Mask)
	}
}

const (
	RegisterCr2FieldSadd7Shift = 7
	RegisterCr2FieldSadd7Mask  = 0x80
)

// GetSadd7 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) GetSadd7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSadd7Mask) != 0
}

// SetSadd7 Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) SetSadd7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSadd7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSadd7Mask)
	}
}

const (
	RegisterCr2FieldSadd8Shift = 8
	RegisterCr2FieldSadd8Mask  = 0x100
)

// GetSadd8 Slave address bit 9:8 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits are dont care In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 9:8 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) GetSadd8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSadd8Mask) != 0
}

// SetSadd8 Slave address bit 9:8 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits are dont care In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 9:8 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) SetSadd8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSadd8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSadd8Mask)
	}
}

const (
	RegisterCr2FieldSadd9Shift = 9
	RegisterCr2FieldSadd9Mask  = 0x200
)

// GetSadd9 Slave address bit 9:8 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits are dont care In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 9:8 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) GetSadd9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSadd9Mask) != 0
}

// SetSadd9 Slave address bit 9:8 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits are dont care In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 9:8 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) SetSadd9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSadd9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSadd9Mask)
	}
}

const (
	RegisterCr2FieldRd_wrnShift = 10
	RegisterCr2FieldRd_wrnMask  = 0x400
)

// GetRd_wrn Transfer direction (master mode) Note: Changing this bit when the START bit is set is not allowed.
func (r *registerCr2Type) GetRd_wrn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldRd_wrnMask) != 0
}

// SetRd_wrn Transfer direction (master mode) Note: Changing this bit when the START bit is set is not allowed.
func (r *registerCr2Type) SetRd_wrn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldRd_wrnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldRd_wrnMask)
	}
}

const (
	RegisterCr2FieldAdd10Shift = 11
	RegisterCr2FieldAdd10Mask  = 0x800
)

// GetAdd10 10-bit addressing mode (master mode) Note: Changing this bit when the START bit is set is not allowed.
func (r *registerCr2Type) GetAdd10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAdd10Mask) != 0
}

// SetAdd10 10-bit addressing mode (master mode) Note: Changing this bit when the START bit is set is not allowed.
func (r *registerCr2Type) SetAdd10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldAdd10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldAdd10Mask)
	}
}

const (
	RegisterCr2FieldHead10rShift = 12
	RegisterCr2FieldHead10rMask  = 0x1000
)

// GetHead10r 10-bit address header only read direction (master receiver mode) Note: Changing this bit when the START bit is set is not allowed.
func (r *registerCr2Type) GetHead10r() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldHead10rMask) != 0
}

// SetHead10r 10-bit address header only read direction (master receiver mode) Note: Changing this bit when the START bit is set is not allowed.
func (r *registerCr2Type) SetHead10r(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldHead10rMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldHead10rMask)
	}
}

const (
	RegisterCr2FieldStartShift = 13
	RegisterCr2FieldStartMask  = 0x2000
)

// GetStart Start generation This bit is set by software, and cleared by hardware after the Start followed by the address sequence is sent, by an arbitration loss, by a timeout error detection, or when PE = 0. It can also be cleared by software by writing 1 to the ADDRCF bit in the I2C_ICR register. If the I2C is already in master mode with AUTOEND = 0, setting this bit generates a Repeated Start condition when RELOAD=0, after the end of the NBYTES transfer. Otherwise setting this bit will generate a START condition once the bus is free. Note: Writing 0 to this bit has no effect. The START bit can be set even if the bus is BUSY or I2C is in slave mode. This bit has no effect when RELOAD is set.
func (r *registerCr2Type) GetStart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldStartMask) != 0
}

// SetStart Start generation This bit is set by software, and cleared by hardware after the Start followed by the address sequence is sent, by an arbitration loss, by a timeout error detection, or when PE = 0. It can also be cleared by software by writing 1 to the ADDRCF bit in the I2C_ICR register. If the I2C is already in master mode with AUTOEND = 0, setting this bit generates a Repeated Start condition when RELOAD=0, after the end of the NBYTES transfer. Otherwise setting this bit will generate a START condition once the bus is free. Note: Writing 0 to this bit has no effect. The START bit can be set even if the bus is BUSY or I2C is in slave mode. This bit has no effect when RELOAD is set.
func (r *registerCr2Type) SetStart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldStartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldStartMask)
	}
}

const (
	RegisterCr2FieldStopShift = 14
	RegisterCr2FieldStopMask  = 0x4000
)

// GetStop Stop generation (master mode) The bit is set by software, cleared by hardware when a Stop condition is detected, or when PE = 0. In Master Mode: Note: Writing 0 to this bit has no effect.
func (r *registerCr2Type) GetStop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldStopMask) != 0
}

// SetStop Stop generation (master mode) The bit is set by software, cleared by hardware when a Stop condition is detected, or when PE = 0. In Master Mode: Note: Writing 0 to this bit has no effect.
func (r *registerCr2Type) SetStop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldStopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldStopMask)
	}
}

const (
	RegisterCr2FieldNackShift = 15
	RegisterCr2FieldNackMask  = 0x8000
)

// GetNack NACK generation (slave mode) The bit is set by software, cleared by hardware when the NACK is sent, or when a STOP condition or an Address matched is received, or when PE=0. Note: Writing 0 to this bit has no effect. This bit is used in slave mode only: in master receiver mode, NACK is automatically generated after last byte preceding STOP or RESTART condition, whatever the NACK bit value. When an overrun occurs in slave receiver NOSTRETCH mode, a NACK is automatically generated whatever the NACK bit value. When hardware PEC checking is enabled (PECBYTE=1), the PEC acknowledge value does not depend on the NACK value.
func (r *registerCr2Type) GetNack() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldNackMask) != 0
}

// SetNack NACK generation (slave mode) The bit is set by software, cleared by hardware when the NACK is sent, or when a STOP condition or an Address matched is received, or when PE=0. Note: Writing 0 to this bit has no effect. This bit is used in slave mode only: in master receiver mode, NACK is automatically generated after last byte preceding STOP or RESTART condition, whatever the NACK bit value. When an overrun occurs in slave receiver NOSTRETCH mode, a NACK is automatically generated whatever the NACK bit value. When hardware PEC checking is enabled (PECBYTE=1), the PEC acknowledge value does not depend on the NACK value.
func (r *registerCr2Type) SetNack(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldNackMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldNackMask)
	}
}

const (
	RegisterCr2FieldNbytesShift = 16
	RegisterCr2FieldNbytesMask  = 0xff0000
)

// GetNbytes Number of bytes The number of bytes to be transmitted/received is programmed there. This field is dont care in slave mode with SBC=0. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) GetNbytes() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldNbytesMask) >> RegisterCr2FieldNbytesShift)
}

// SetNbytes Number of bytes The number of bytes to be transmitted/received is programmed there. This field is dont care in slave mode with SBC=0. Note: Changing these bits when the START bit is set is not allowed.
func (r *registerCr2Type) SetNbytes(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldNbytesMask)|(uint32(value)<<RegisterCr2FieldNbytesShift))
}

const (
	RegisterCr2FieldReloadShift = 24
	RegisterCr2FieldReloadMask  = 0x1000000
)

// GetReload NBYTES reload mode This bit is set and cleared by software.
func (r *registerCr2Type) GetReload() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldReloadMask) != 0
}

// SetReload NBYTES reload mode This bit is set and cleared by software.
func (r *registerCr2Type) SetReload(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldReloadMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldReloadMask)
	}
}

const (
	RegisterCr2FieldAutoendShift = 25
	RegisterCr2FieldAutoendMask  = 0x2000000
)

// GetAutoend Automatic end mode (master mode) This bit is set and cleared by software. Note: This bit has no effect in slave mode or when the RELOAD bit is set.
func (r *registerCr2Type) GetAutoend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAutoendMask) != 0
}

// SetAutoend Automatic end mode (master mode) This bit is set and cleared by software. Note: This bit has no effect in slave mode or when the RELOAD bit is set.
func (r *registerCr2Type) SetAutoend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldAutoendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldAutoendMask)
	}
}

const (
	RegisterCr2FieldPecbyteShift = 26
	RegisterCr2FieldPecbyteMask  = 0x4000000
)

// GetPecbyte Packet error checking byte This bit is set by software, and cleared by hardware when the PEC is transferred, or when a STOP condition or an Address matched is received, also when PE=0. Note: Writing 0 to this bit has no effect. This bit has no effect when RELOAD is set. This bit has no effect is slave mode when SBC=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerCr2Type) GetPecbyte() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldPecbyteMask) != 0
}

// SetPecbyte Packet error checking byte This bit is set by software, and cleared by hardware when the PEC is transferred, or when a STOP condition or an Address matched is received, also when PE=0. Note: Writing 0 to this bit has no effect. This bit has no effect when RELOAD is set. This bit has no effect is slave mode when SBC=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerCr2Type) SetPecbyte(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldPecbyteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldPecbyteMask)
	}
}

// registerOar1Type Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.
type registerOar1Type uint32

const (
	RegisterOar1FieldOa1Shift = 0
	RegisterOar1FieldOa1Mask  = 0x3ff
)

// GetOa1 Interface address 7-bit addressing mode: dont care 10-bit addressing mode: bits 9:8 of address Note: These bits can be written only when OA1EN=0. OA1[7:1]: Interface address Bits 7:1 of address Note: These bits can be written only when OA1EN=0. OA1[0]: Interface address 7-bit addressing mode: dont care 10-bit addressing mode: bit 0 of address Note: This bit can be written only when OA1EN=0.
func (r *registerOar1Type) GetOa1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterOar1FieldOa1Mask) >> RegisterOar1FieldOa1Shift)
}

// SetOa1 Interface address 7-bit addressing mode: dont care 10-bit addressing mode: bits 9:8 of address Note: These bits can be written only when OA1EN=0. OA1[7:1]: Interface address Bits 7:1 of address Note: These bits can be written only when OA1EN=0. OA1[0]: Interface address 7-bit addressing mode: dont care 10-bit addressing mode: bit 0 of address Note: This bit can be written only when OA1EN=0.
func (r *registerOar1Type) SetOa1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOar1FieldOa1Mask)|(uint32(value)<<RegisterOar1FieldOa1Shift))
}

const (
	RegisterOar1FieldOa1modeShift = 10
	RegisterOar1FieldOa1modeMask  = 0x400
)

// GetOa1mode Own Address 1 10-bit mode Note: This bit can be written only when OA1EN=0.
func (r *registerOar1Type) GetOa1mode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOar1FieldOa1modeMask) != 0
}

// SetOa1mode Own Address 1 10-bit mode Note: This bit can be written only when OA1EN=0.
func (r *registerOar1Type) SetOa1mode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOar1FieldOa1modeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOar1FieldOa1modeMask)
	}
}

const (
	RegisterOar1FieldOa1enShift = 15
	RegisterOar1FieldOa1enMask  = 0x8000
)

// GetOa1en Own Address 1 enable
func (r *registerOar1Type) GetOa1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOar1FieldOa1enMask) != 0
}

// SetOa1en Own Address 1 enable
func (r *registerOar1Type) SetOa1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOar1FieldOa1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOar1FieldOa1enMask)
	}
}

// registerOar2Type Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.
type registerOar2Type uint32

const (
	RegisterOar2FieldOa2Shift = 1
	RegisterOar2FieldOa2Mask  = 0xfe
)

// GetOa2 Interface address bits 7:1 of address Note: These bits can be written only when OA2EN=0.
func (r *registerOar2Type) GetOa2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOar2FieldOa2Mask) >> RegisterOar2FieldOa2Shift)
}

// SetOa2 Interface address bits 7:1 of address Note: These bits can be written only when OA2EN=0.
func (r *registerOar2Type) SetOa2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOar2FieldOa2Mask)|(uint32(value)<<RegisterOar2FieldOa2Shift))
}

const (
	RegisterOar2FieldOa2mskShift = 8
	RegisterOar2FieldOa2mskMask  = 0x700
)

// GetOa2msk Own Address 2 masks Note: These bits can be written only when OA2EN=0. As soon as OA2MSK is not equal to 0, the reserved I2C addresses (0b0000xxx and 0b1111xxx) are not acknowledged even if the comparison matches.
func (r *registerOar2Type) GetOa2msk() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterOar2FieldOa2mskMask) >> RegisterOar2FieldOa2mskShift)
}

// SetOa2msk Own Address 2 masks Note: These bits can be written only when OA2EN=0. As soon as OA2MSK is not equal to 0, the reserved I2C addresses (0b0000xxx and 0b1111xxx) are not acknowledged even if the comparison matches.
func (r *registerOar2Type) SetOa2msk(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOar2FieldOa2mskMask)|(uint32(value)<<RegisterOar2FieldOa2mskShift))
}

const (
	RegisterOar2FieldOa2enShift = 15
	RegisterOar2FieldOa2enMask  = 0x8000
)

// GetOa2en Own Address 2 enable
func (r *registerOar2Type) GetOa2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOar2FieldOa2enMask) != 0
}

// SetOa2en Own Address 2 enable
func (r *registerOar2Type) SetOa2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOar2FieldOa2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOar2FieldOa2enMask)
	}
}

// registerTimingrType Access: No wait states
type registerTimingrType uint32

const (
	RegisterTimingrFieldScllShift = 0
	RegisterTimingrFieldScllMask  = 0xff
)

// GetScll SCL low period (master mode) This field is used to generate the SCL low period in master mode. tSCLL = (SCLL+1) x tPRESC Note: SCLL is also used to generate tBUF and tSU:STA timings.
func (r *registerTimingrType) GetScll() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimingrFieldScllMask) >> RegisterTimingrFieldScllShift)
}

// SetScll SCL low period (master mode) This field is used to generate the SCL low period in master mode. tSCLL = (SCLL+1) x tPRESC Note: SCLL is also used to generate tBUF and tSU:STA timings.
func (r *registerTimingrType) SetScll(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimingrFieldScllMask)|(uint32(value)<<RegisterTimingrFieldScllShift))
}

const (
	RegisterTimingrFieldSclhShift = 8
	RegisterTimingrFieldSclhMask  = 0xff00
)

// GetSclh SCL high period (master mode) This field is used to generate the SCL high period in master mode. tSCLH = (SCLH+1) x tPRESC Note: SCLH is also used to generate tSU:STO and tHD:STA timing.
func (r *registerTimingrType) GetSclh() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimingrFieldSclhMask) >> RegisterTimingrFieldSclhShift)
}

// SetSclh SCL high period (master mode) This field is used to generate the SCL high period in master mode. tSCLH = (SCLH+1) x tPRESC Note: SCLH is also used to generate tSU:STO and tHD:STA timing.
func (r *registerTimingrType) SetSclh(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimingrFieldSclhMask)|(uint32(value)<<RegisterTimingrFieldSclhShift))
}

const (
	RegisterTimingrFieldSdadelShift = 16
	RegisterTimingrFieldSdadelMask  = 0xf0000
)

// GetSdadel Data hold time This field is used to generate the delay tSDADEL between SCL falling edge and SDA edge. In master mode and in slave mode with NOSTRETCH = 0, the SCL line is stretched low during tSDADEL. tSDADEL= SDADEL x tPRESC Note: SDADEL is used to generate tHD:DAT timing.
func (r *registerTimingrType) GetSdadel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimingrFieldSdadelMask) >> RegisterTimingrFieldSdadelShift)
}

// SetSdadel Data hold time This field is used to generate the delay tSDADEL between SCL falling edge and SDA edge. In master mode and in slave mode with NOSTRETCH = 0, the SCL line is stretched low during tSDADEL. tSDADEL= SDADEL x tPRESC Note: SDADEL is used to generate tHD:DAT timing.
func (r *registerTimingrType) SetSdadel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimingrFieldSdadelMask)|(uint32(value)<<RegisterTimingrFieldSdadelShift))
}

const (
	RegisterTimingrFieldScldelShift = 20
	RegisterTimingrFieldScldelMask  = 0xf00000
)

// GetScldel Data setup time This field is used to generate a delay tSCLDEL between SDA edge and SCL rising edge. In master mode and in slave mode with NOSTRETCH = 0, the SCL line is stretched low during tSCLDEL. tSCLDEL = (SCLDEL+1) x tPRESC Note: tSCLDEL is used to generate tSU:DAT timing.
func (r *registerTimingrType) GetScldel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimingrFieldScldelMask) >> RegisterTimingrFieldScldelShift)
}

// SetScldel Data setup time This field is used to generate a delay tSCLDEL between SDA edge and SCL rising edge. In master mode and in slave mode with NOSTRETCH = 0, the SCL line is stretched low during tSCLDEL. tSCLDEL = (SCLDEL+1) x tPRESC Note: tSCLDEL is used to generate tSU:DAT timing.
func (r *registerTimingrType) SetScldel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimingrFieldScldelMask)|(uint32(value)<<RegisterTimingrFieldScldelShift))
}

const (
	RegisterTimingrFieldPrescShift = 28
	RegisterTimingrFieldPrescMask  = 0xf0000000
)

// GetPresc Timing prescaler This field is used to prescale I2CCLK in order to generate the clock period tPRESC used for data setup and hold counters (refer to I2C timings on page9) and for SCL high and low level counters (refer to I2C master initialization on page24). tPRESC = (PRESC+1) x tI2CCLK
func (r *registerTimingrType) GetPresc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTimingrFieldPrescMask) >> RegisterTimingrFieldPrescShift)
}

// SetPresc Timing prescaler This field is used to prescale I2CCLK in order to generate the clock period tPRESC used for data setup and hold counters (refer to I2C timings on page9) and for SCL high and low level counters (refer to I2C master initialization on page24). tPRESC = (PRESC+1) x tI2CCLK
func (r *registerTimingrType) SetPresc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimingrFieldPrescMask)|(uint32(value)<<RegisterTimingrFieldPrescShift))
}

// registerTimeoutrType Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.
type registerTimeoutrType uint32

const (
	RegisterTimeoutrFieldTimeoutaShift = 0
	RegisterTimeoutrFieldTimeoutaMask  = 0xfff
)

// GetTimeouta Bus Timeout A This field is used to configure: The SCL low timeout condition tTIMEOUT when TIDLE=0 tTIMEOUT= (TIMEOUTA+1) x 2048 x tI2CCLK The bus idle condition (both SCL and SDA high) when TIDLE=1 tIDLE= (TIMEOUTA+1) x 4 x tI2CCLK Note: These bits can be written only when TIMOUTEN=0.
func (r *registerTimeoutrType) GetTimeouta() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterTimeoutrFieldTimeoutaMask) >> RegisterTimeoutrFieldTimeoutaShift)
}

// SetTimeouta Bus Timeout A This field is used to configure: The SCL low timeout condition tTIMEOUT when TIDLE=0 tTIMEOUT= (TIMEOUTA+1) x 2048 x tI2CCLK The bus idle condition (both SCL and SDA high) when TIDLE=1 tIDLE= (TIMEOUTA+1) x 4 x tI2CCLK Note: These bits can be written only when TIMOUTEN=0.
func (r *registerTimeoutrType) SetTimeouta(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimeoutrFieldTimeoutaMask)|(uint32(value)<<RegisterTimeoutrFieldTimeoutaShift))
}

const (
	RegisterTimeoutrFieldTidleShift = 12
	RegisterTimeoutrFieldTidleMask  = 0x1000
)

// GetTidle Idle clock timeout detection Note: This bit can be written only when TIMOUTEN=0.
func (r *registerTimeoutrType) GetTidle() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeoutrFieldTidleMask) != 0
}

// SetTidle Idle clock timeout detection Note: This bit can be written only when TIMOUTEN=0.
func (r *registerTimeoutrType) SetTidle(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeoutrFieldTidleMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeoutrFieldTidleMask)
	}
}

const (
	RegisterTimeoutrFieldTimoutenShift = 15
	RegisterTimeoutrFieldTimoutenMask  = 0x8000
)

// GetTimouten Clock timeout enable
func (r *registerTimeoutrType) GetTimouten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeoutrFieldTimoutenMask) != 0
}

// SetTimouten Clock timeout enable
func (r *registerTimeoutrType) SetTimouten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeoutrFieldTimoutenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeoutrFieldTimoutenMask)
	}
}

const (
	RegisterTimeoutrFieldTimeoutbShift = 16
	RegisterTimeoutrFieldTimeoutbMask  = 0xfff0000
)

// GetTimeoutb Bus timeout B This field is used to configure the cumulative clock extension timeout: In master mode, the master cumulative clock low extend time (tLOW:MEXT) is detected In slave mode, the slave cumulative clock low extend time (tLOW:SEXT) is detected tLOW:EXT= (TIMEOUTB+1) x 2048 x tI2CCLK Note: These bits can be written only when TEXTEN=0.
func (r *registerTimeoutrType) GetTimeoutb() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterTimeoutrFieldTimeoutbMask) >> RegisterTimeoutrFieldTimeoutbShift)
}

// SetTimeoutb Bus timeout B This field is used to configure the cumulative clock extension timeout: In master mode, the master cumulative clock low extend time (tLOW:MEXT) is detected In slave mode, the slave cumulative clock low extend time (tLOW:SEXT) is detected tLOW:EXT= (TIMEOUTB+1) x 2048 x tI2CCLK Note: These bits can be written only when TEXTEN=0.
func (r *registerTimeoutrType) SetTimeoutb(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTimeoutrFieldTimeoutbMask)|(uint32(value)<<RegisterTimeoutrFieldTimeoutbShift))
}

const (
	RegisterTimeoutrFieldTextenShift = 31
	RegisterTimeoutrFieldTextenMask  = 0x80000000
)

// GetTexten Extended clock timeout enable
func (r *registerTimeoutrType) GetTexten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTimeoutrFieldTextenMask) != 0
}

// SetTexten Extended clock timeout enable
func (r *registerTimeoutrType) SetTexten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTimeoutrFieldTextenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTimeoutrFieldTextenMask)
	}
}

// registerIsrType Access: No wait states
type registerIsrType uint32

const (
	RegisterIsrFieldTxeShift = 0
	RegisterIsrFieldTxeMask  = 0x1
)

// GetTxe Transmit data register empty (transmitters) This bit is set by hardware when the I2C_TXDR register is empty. It is cleared when the next data to be sent is written in the I2C_TXDR register. This bit can be written to 1 by software in order to flush the transmit data register I2C_TXDR. Note: This bit is set by hardware when PE=0.
func (r *registerIsrType) GetTxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxeMask) != 0
}

// SetTxe Transmit data register empty (transmitters) This bit is set by hardware when the I2C_TXDR register is empty. It is cleared when the next data to be sent is written in the I2C_TXDR register. This bit can be written to 1 by software in order to flush the transmit data register I2C_TXDR. Note: This bit is set by hardware when PE=0.
func (r *registerIsrType) SetTxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxeMask)
	}
}

const (
	RegisterIsrFieldTxisShift = 1
	RegisterIsrFieldTxisMask  = 0x2
)

// GetTxis Transmit interrupt status (transmitters) This bit is set by hardware when the I2C_TXDR register is empty and the data to be transmitted must be written in the I2C_TXDR register. It is cleared when the next data to be sent is written in the I2C_TXDR register. This bit can be written to 1 by software when NOSTRETCH=1 only, in order to generate a TXIS event (interrupt if TXIE=1 or DMA request if TXDMAEN=1). Note: This bit is cleared by hardware when PE=0.
func (r *registerIsrType) GetTxis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxisMask) != 0
}

// SetTxis Transmit interrupt status (transmitters) This bit is set by hardware when the I2C_TXDR register is empty and the data to be transmitted must be written in the I2C_TXDR register. It is cleared when the next data to be sent is written in the I2C_TXDR register. This bit can be written to 1 by software when NOSTRETCH=1 only, in order to generate a TXIS event (interrupt if TXIE=1 or DMA request if TXDMAEN=1). Note: This bit is cleared by hardware when PE=0.
func (r *registerIsrType) SetTxis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxisMask)
	}
}

const (
	RegisterIsrFieldRxneShift = 2
	RegisterIsrFieldRxneMask  = 0x4
)

// GetRxne Receive data register not empty (receivers) This bit is set by hardware when the received data is copied into the I2C_RXDR register, and is ready to be read. It is cleared when I2C_RXDR is read. Note: This bit is cleared by hardware when PE=0.
func (r *registerIsrType) GetRxne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxneMask) != 0
}

const (
	RegisterIsrFieldAddrShift = 3
	RegisterIsrFieldAddrMask  = 0x8
)

// GetAddr Address matched (slave mode) This bit is set by hardware as soon as the received slave address matched with one of the enabled slave addresses. It is cleared by software by setting ADDRCF bit. Note: This bit is cleared by hardware when PE=0.
func (r *registerIsrType) GetAddr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldAddrMask) != 0
}

const (
	RegisterIsrFieldNackfShift = 4
	RegisterIsrFieldNackfMask  = 0x10
)

// GetNackf Not Acknowledge received flag This flag is set by hardware when a NACK is received after a byte transmission. It is cleared by software by setting the NACKCF bit. Note: This bit is cleared by hardware when PE=0.
func (r *registerIsrType) GetNackf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldNackfMask) != 0
}

const (
	RegisterIsrFieldStopfShift = 5
	RegisterIsrFieldStopfMask  = 0x20
)

// GetStopf Stop detection flag This flag is set by hardware when a Stop condition is detected on the bus and the peripheral is involved in this transfer: either as a master, provided that the STOP condition is generated by the peripheral. or as a slave, provided that the peripheral has been addressed previously during this transfer. It is cleared by software by setting the STOPCF bit. Note: This bit is cleared by hardware when PE=0.
func (r *registerIsrType) GetStopf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldStopfMask) != 0
}

const (
	RegisterIsrFieldTcShift = 6
	RegisterIsrFieldTcMask  = 0x40
)

// GetTc Transfer Complete (master mode) This flag is set by hardware when RELOAD=0, AUTOEND=0 and NBYTES data have been transferred. It is cleared by software when START bit or STOP bit is set. Note: This bit is cleared by hardware when PE=0.
func (r *registerIsrType) GetTc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcMask) != 0
}

const (
	RegisterIsrFieldTcrShift = 7
	RegisterIsrFieldTcrMask  = 0x80
)

// GetTcr Transfer Complete Reload This flag is set by hardware when RELOAD=1 and NBYTES data have been transferred. It is cleared by software when NBYTES is written to a non-zero value. Note: This bit is cleared by hardware when PE=0. This flag is only for master mode, or for slave mode when the SBC bit is set.
func (r *registerIsrType) GetTcr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcrMask) != 0
}

const (
	RegisterIsrFieldBerrShift = 8
	RegisterIsrFieldBerrMask  = 0x100
)

// GetBerr Bus error This flag is set by hardware when a misplaced Start or Stop condition is detected whereas the peripheral is involved in the transfer. The flag is not set during the address phase in slave mode. It is cleared by software by setting BERRCF bit. Note: This bit is cleared by hardware when PE=0.
func (r *registerIsrType) GetBerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldBerrMask) != 0
}

const (
	RegisterIsrFieldArloShift = 9
	RegisterIsrFieldArloMask  = 0x200
)

// GetArlo Arbitration lost This flag is set by hardware in case of arbitration loss. It is cleared by software by setting the ARLOCF bit. Note: This bit is cleared by hardware when PE=0.
func (r *registerIsrType) GetArlo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldArloMask) != 0
}

const (
	RegisterIsrFieldOvrShift = 10
	RegisterIsrFieldOvrMask  = 0x400
)

// GetOvr Overrun/Underrun (slave mode) This flag is set by hardware in slave mode with NOSTRETCH=1, when an overrun/underrun error occurs. It is cleared by software by setting the OVRCF bit. Note: This bit is cleared by hardware when PE=0.
func (r *registerIsrType) GetOvr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldOvrMask) != 0
}

const (
	RegisterIsrFieldPecerrShift = 11
	RegisterIsrFieldPecerrMask  = 0x800
)

// GetPecerr PEC Error in reception This flag is set by hardware when the received PEC does not match with the PEC register content. A NACK is automatically sent after the wrong PEC reception. It is cleared by software by setting the PECCF bit. Note: This bit is cleared by hardware when PE=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerIsrType) GetPecerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldPecerrMask) != 0
}

const (
	RegisterIsrFieldTimeoutShift = 12
	RegisterIsrFieldTimeoutMask  = 0x1000
)

// GetTimeout Timeout or tLOW detection flag This flag is set by hardware when a timeout or extended clock timeout occurred. It is cleared by software by setting the TIMEOUTCF bit. Note: This bit is cleared by hardware when PE=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerIsrType) GetTimeout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTimeoutMask) != 0
}

const (
	RegisterIsrFieldAlertShift = 13
	RegisterIsrFieldAlertMask  = 0x2000
)

// GetAlert SMBus alert This flag is set by hardware when SMBHEN=1 (SMBus host configuration), ALERTEN=1 and a SMBALERT event (falling edge) is detected on SMBA pin. It is cleared by software by setting the ALERTCF bit. Note: This bit is cleared by hardware when PE=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerIsrType) GetAlert() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldAlertMask) != 0
}

const (
	RegisterIsrFieldBusyShift = 15
	RegisterIsrFieldBusyMask  = 0x8000
)

// GetBusy Bus busy This flag indicates that a communication is in progress on the bus. It is set by hardware when a START condition is detected. It is cleared by hardware when a Stop condition is detected, or when PE=0.
func (r *registerIsrType) GetBusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldBusyMask) != 0
}

const (
	RegisterIsrFieldDirShift = 16
	RegisterIsrFieldDirMask  = 0x10000
)

// GetDir Transfer direction (Slave mode) This flag is updated when an address match event occurs (ADDR=1).
func (r *registerIsrType) GetDir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldDirMask) != 0
}

const (
	RegisterIsrFieldAddcodeShift = 17
	RegisterIsrFieldAddcodeMask  = 0xfe0000
)

// GetAddcode Address match code (Slave mode) These bits are updated with the received address when an address match event occurs (ADDR = 1). In the case of a 10-bit address, ADDCODE provides the 10-bit header followed by the 2 MSBs of the address.
func (r *registerIsrType) GetAddcode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldAddcodeMask) >> RegisterIsrFieldAddcodeShift)
}

// registerIcrType Access: No wait states
type registerIcrType uint32

const (
	RegisterIcrFieldAddrcfShift = 3
	RegisterIcrFieldAddrcfMask  = 0x8
)

// GetAddrcf Address matched flag clear Writing 1 to this bit clears the ADDR flag in the I2C_ISR register. Writing 1 to this bit also clears the START bit in the I2C_CR2 register.
func (r *registerIcrType) GetAddrcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldAddrcfMask) != 0
}

// SetAddrcf Address matched flag clear Writing 1 to this bit clears the ADDR flag in the I2C_ISR register. Writing 1 to this bit also clears the START bit in the I2C_CR2 register.
func (r *registerIcrType) SetAddrcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldAddrcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldAddrcfMask)
	}
}

const (
	RegisterIcrFieldNackcfShift = 4
	RegisterIcrFieldNackcfMask  = 0x10
)

// GetNackcf Not Acknowledge flag clear Writing 1 to this bit clears the ACKF flag in I2C_ISR register.
func (r *registerIcrType) GetNackcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldNackcfMask) != 0
}

// SetNackcf Not Acknowledge flag clear Writing 1 to this bit clears the ACKF flag in I2C_ISR register.
func (r *registerIcrType) SetNackcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldNackcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldNackcfMask)
	}
}

const (
	RegisterIcrFieldStopcfShift = 5
	RegisterIcrFieldStopcfMask  = 0x20
)

// GetStopcf Stop detection flag clear Writing 1 to this bit clears the STOPF flag in the I2C_ISR register.
func (r *registerIcrType) GetStopcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldStopcfMask) != 0
}

// SetStopcf Stop detection flag clear Writing 1 to this bit clears the STOPF flag in the I2C_ISR register.
func (r *registerIcrType) SetStopcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldStopcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldStopcfMask)
	}
}

const (
	RegisterIcrFieldBerrcfShift = 8
	RegisterIcrFieldBerrcfMask  = 0x100
)

// GetBerrcf Bus error flag clear Writing 1 to this bit clears the BERRF flag in the I2C_ISR register.
func (r *registerIcrType) GetBerrcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldBerrcfMask) != 0
}

// SetBerrcf Bus error flag clear Writing 1 to this bit clears the BERRF flag in the I2C_ISR register.
func (r *registerIcrType) SetBerrcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldBerrcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldBerrcfMask)
	}
}

const (
	RegisterIcrFieldArlocfShift = 9
	RegisterIcrFieldArlocfMask  = 0x200
)

// GetArlocf Arbitration Lost flag clear Writing 1 to this bit clears the ARLO flag in the I2C_ISR register.
func (r *registerIcrType) GetArlocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldArlocfMask) != 0
}

// SetArlocf Arbitration Lost flag clear Writing 1 to this bit clears the ARLO flag in the I2C_ISR register.
func (r *registerIcrType) SetArlocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldArlocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldArlocfMask)
	}
}

const (
	RegisterIcrFieldOvrcfShift = 10
	RegisterIcrFieldOvrcfMask  = 0x400
)

// GetOvrcf Overrun/Underrun flag clear Writing 1 to this bit clears the OVR flag in the I2C_ISR register.
func (r *registerIcrType) GetOvrcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldOvrcfMask) != 0
}

// SetOvrcf Overrun/Underrun flag clear Writing 1 to this bit clears the OVR flag in the I2C_ISR register.
func (r *registerIcrType) SetOvrcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldOvrcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldOvrcfMask)
	}
}

const (
	RegisterIcrFieldPeccfShift = 11
	RegisterIcrFieldPeccfMask  = 0x800
)

// GetPeccf PEC Error flag clear Writing 1 to this bit clears the PECERR flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerIcrType) GetPeccf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldPeccfMask) != 0
}

// SetPeccf PEC Error flag clear Writing 1 to this bit clears the PECERR flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerIcrType) SetPeccf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldPeccfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldPeccfMask)
	}
}

const (
	RegisterIcrFieldTimoutcfShift = 12
	RegisterIcrFieldTimoutcfMask  = 0x1000
)

// GetTimoutcf Timeout detection flag clear Writing 1 to this bit clears the TIMEOUT flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerIcrType) GetTimoutcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldTimoutcfMask) != 0
}

// SetTimoutcf Timeout detection flag clear Writing 1 to this bit clears the TIMEOUT flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerIcrType) SetTimoutcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldTimoutcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldTimoutcfMask)
	}
}

const (
	RegisterIcrFieldAlertcfShift = 13
	RegisterIcrFieldAlertcfMask  = 0x2000
)

// GetAlertcf Alert flag clear Writing 1 to this bit clears the ALERT flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerIcrType) GetAlertcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldAlertcfMask) != 0
}

// SetAlertcf Alert flag clear Writing 1 to this bit clears the ALERT flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.
func (r *registerIcrType) SetAlertcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldAlertcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldAlertcfMask)
	}
}

// registerPecrType Access: No wait states
type registerPecrType uint32

const (
	RegisterPecrFieldPecShift = 0
	RegisterPecrFieldPecMask  = 0xff
)

// GetPec Packet error checking register This field contains the internal PEC when PECEN=1. The PEC is cleared by hardware when PE=0.
func (r *registerPecrType) GetPec() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPecrFieldPecMask) >> RegisterPecrFieldPecShift)
}

// SetPec Packet error checking register This field contains the internal PEC when PECEN=1. The PEC is cleared by hardware when PE=0.
func (r *registerPecrType) SetPec(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPecrFieldPecMask)|(uint32(value)<<RegisterPecrFieldPecShift))
}

// registerRxdrType Access: No wait states
type registerRxdrType uint32

const (
	RegisterRxdrFieldRxdataShift = 0
	RegisterRxdrFieldRxdataMask  = 0xff
)

// GetRxdata 8-bit receive data Data byte received from the I2C bus.
func (r *registerRxdrType) GetRxdata() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRxdrFieldRxdataMask) >> RegisterRxdrFieldRxdataShift)
}

// SetRxdata 8-bit receive data Data byte received from the I2C bus.
func (r *registerRxdrType) SetRxdata(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxdrFieldRxdataMask)|(uint32(value)<<RegisterRxdrFieldRxdataShift))
}

// registerTxdrType Access: No wait states
type registerTxdrType uint32

const (
	RegisterTxdrFieldTxdataShift = 0
	RegisterTxdrFieldTxdataMask  = 0xff
)

// GetTxdata 8-bit transmit data Data byte to be transmitted to the I2C bus. Note: These bits can be written only when TXE=1.
func (r *registerTxdrType) GetTxdata() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTxdrFieldTxdataMask) >> RegisterTxdrFieldTxdataShift)
}

// SetTxdata 8-bit transmit data Data byte to be transmitted to the I2C bus. Note: These bits can be written only when TXE=1.
func (r *registerTxdrType) SetTxdata(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxdrFieldTxdataMask)|(uint32(value)<<RegisterTxdrFieldTxdataShift))
}
