//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package pwr

import (
	"unsafe"
	"volatile"
)

var (
	Pwr = (*_pwr)(unsafe.Pointer(uintptr(0x58024800)))
)

type _pwr struct {
	Cr1     registerCr1Type
	Csr1    registerCsr1Type
	Cr2     registerCr2Type
	Cr3     registerCr3Type
	Cpucr   registerCpucrType
	_       [4]byte
	D3cr    registerD3crType
	_       [4]byte
	Wkupcr  registerWkupcrType
	Wkupfr  registerWkupfrType
	Wkupepr registerWkupeprType
}

// registerCr1Type PWR control register 1
type registerCr1Type uint32

const (
	RegisterCr1FieldLpdsShift = 0
	RegisterCr1FieldLpdsMask  = 0x1
)

// GetLpds Low-power Deepsleep with SVOS3 (SVOS4 and SVOS5 always use low-power, regardless of the setting of this bit)
func (r *registerCr1Type) GetLpds() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldLpdsMask) != 0
}

// SetLpds Low-power Deepsleep with SVOS3 (SVOS4 and SVOS5 always use low-power, regardless of the setting of this bit)
func (r *registerCr1Type) SetLpds(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldLpdsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldLpdsMask)
	}
}

const (
	RegisterCr1FieldPvdeShift = 4
	RegisterCr1FieldPvdeMask  = 0x10
)

// GetPvde Programmable voltage detector enable
func (r *registerCr1Type) GetPvde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPvdeMask) != 0
}

// SetPvde Programmable voltage detector enable
func (r *registerCr1Type) SetPvde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldPvdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldPvdeMask)
	}
}

const (
	RegisterCr1FieldPlsShift = 5
	RegisterCr1FieldPlsMask  = 0xe0
)

// GetPls Programmable voltage detector level selection These bits select the voltage threshold detected by the PVD. Note: Refer to Section Electrical characteristics of the product datasheet for more details.
func (r *registerCr1Type) GetPls() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPlsMask) >> RegisterCr1FieldPlsShift)
}

// SetPls Programmable voltage detector level selection These bits select the voltage threshold detected by the PVD. Note: Refer to Section Electrical characteristics of the product datasheet for more details.
func (r *registerCr1Type) SetPls(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldPlsMask)|(uint32(value)<<RegisterCr1FieldPlsShift))
}

const (
	RegisterCr1FieldDbpShift = 8
	RegisterCr1FieldDbpMask  = 0x100
)

// GetDbp Disable backup domain write protection In reset state, the RCC_BDCR register, the RTC registers (including the backup registers), BREN and MOEN bits in PWR_CR2 register, are protected against parasitic write access. This bit must be set to enable write access to these registers.
func (r *registerCr1Type) GetDbp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDbpMask) != 0
}

// SetDbp Disable backup domain write protection In reset state, the RCC_BDCR register, the RTC registers (including the backup registers), BREN and MOEN bits in PWR_CR2 register, are protected against parasitic write access. This bit must be set to enable write access to these registers.
func (r *registerCr1Type) SetDbp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDbpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDbpMask)
	}
}

const (
	RegisterCr1FieldFlpsShift = 9
	RegisterCr1FieldFlpsMask  = 0x200
)

// GetFlps Flash low-power mode in DStop mode This bit allows to obtain the best trade-off between low-power consumption and restart time when exiting from DStop mode. When it is set, the Flash memory enters low-power mode when D1 domain is in DStop mode.
func (r *registerCr1Type) GetFlps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldFlpsMask) != 0
}

// SetFlps Flash low-power mode in DStop mode This bit allows to obtain the best trade-off between low-power consumption and restart time when exiting from DStop mode. When it is set, the Flash memory enters low-power mode when D1 domain is in DStop mode.
func (r *registerCr1Type) SetFlps(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldFlpsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldFlpsMask)
	}
}

const (
	RegisterCr1FieldSvosShift = 14
	RegisterCr1FieldSvosMask  = 0xc000
)

// GetSvos System Stop mode voltage scaling selection These bits control the VCORE voltage level in system Stop mode, to obtain the best trade-off between power consumption and performance.
func (r *registerCr1Type) GetSvos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldSvosMask) >> RegisterCr1FieldSvosShift)
}

// SetSvos System Stop mode voltage scaling selection These bits control the VCORE voltage level in system Stop mode, to obtain the best trade-off between power consumption and performance.
func (r *registerCr1Type) SetSvos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldSvosMask)|(uint32(value)<<RegisterCr1FieldSvosShift))
}

const (
	RegisterCr1FieldAvdenShift = 16
	RegisterCr1FieldAvdenMask  = 0x10000
)

// GetAvden Peripheral voltage monitor on VDDA enable
func (r *registerCr1Type) GetAvden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldAvdenMask) != 0
}

// SetAvden Peripheral voltage monitor on VDDA enable
func (r *registerCr1Type) SetAvden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldAvdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldAvdenMask)
	}
}

const (
	RegisterCr1FieldAlsShift = 17
	RegisterCr1FieldAlsMask  = 0x60000
)

// GetAls Analog voltage detector level selection These bits select the voltage threshold detected by the AVD.
func (r *registerCr1Type) GetAls() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldAlsMask) >> RegisterCr1FieldAlsShift)
}

// SetAls Analog voltage detector level selection These bits select the voltage threshold detected by the AVD.
func (r *registerCr1Type) SetAls(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldAlsMask)|(uint32(value)<<RegisterCr1FieldAlsShift))
}

// registerCsr1Type PWR control status register 1
type registerCsr1Type uint32

const (
	RegisterCsr1FieldPvdoShift = 4
	RegisterCsr1FieldPvdoMask  = 0x10
)

// GetPvdo Programmable voltage detect output This bit is set and cleared by hardware. It is valid only if the PVD has been enabled by the PVDE bit. Note: since the PVD is disabled in Standby mode, this bit is equal to 0 after Standby or reset until the PVDE bit is set.
func (r *registerCsr1Type) GetPvdo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsr1FieldPvdoMask) != 0
}

// SetPvdo Programmable voltage detect output This bit is set and cleared by hardware. It is valid only if the PVD has been enabled by the PVDE bit. Note: since the PVD is disabled in Standby mode, this bit is equal to 0 after Standby or reset until the PVDE bit is set.
func (r *registerCsr1Type) SetPvdo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsr1FieldPvdoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsr1FieldPvdoMask)
	}
}

const (
	RegisterCsr1FieldActvosrdyShift = 13
	RegisterCsr1FieldActvosrdyMask  = 0x2000
)

// GetActvosrdy Voltage levels ready bit for currently used VOS and SDLEVEL This bit is set to 1 by hardware when the voltage regulator and the SD converter are both disabled and Bypass mode is selected in PWR control register 3 (PWR_CR3).
func (r *registerCsr1Type) GetActvosrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsr1FieldActvosrdyMask) != 0
}

// SetActvosrdy Voltage levels ready bit for currently used VOS and SDLEVEL This bit is set to 1 by hardware when the voltage regulator and the SD converter are both disabled and Bypass mode is selected in PWR control register 3 (PWR_CR3).
func (r *registerCsr1Type) SetActvosrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsr1FieldActvosrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsr1FieldActvosrdyMask)
	}
}

const (
	RegisterCsr1FieldActvosShift = 14
	RegisterCsr1FieldActvosMask  = 0xc000
)

// GetActvos VOS currently applied for VCORE voltage scaling selection. These bits reflect the last VOS value applied to the PMU.
func (r *registerCsr1Type) GetActvos() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCsr1FieldActvosMask) >> RegisterCsr1FieldActvosShift)
}

// SetActvos VOS currently applied for VCORE voltage scaling selection. These bits reflect the last VOS value applied to the PMU.
func (r *registerCsr1Type) SetActvos(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCsr1FieldActvosMask)|(uint32(value)<<RegisterCsr1FieldActvosShift))
}

const (
	RegisterCsr1FieldAvdoShift = 16
	RegisterCsr1FieldAvdoMask  = 0x10000
)

// GetAvdo Analog voltage detector output on VDDA This bit is set and cleared by hardware. It is valid only if AVD on VDDA is enabled by the AVDEN bit. Note: Since the AVD is disabled in Standby mode, this bit is equal to 0 after Standby or reset until the AVDEN bit is set.
func (r *registerCsr1Type) GetAvdo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsr1FieldAvdoMask) != 0
}

// SetAvdo Analog voltage detector output on VDDA This bit is set and cleared by hardware. It is valid only if AVD on VDDA is enabled by the AVDEN bit. Note: Since the AVD is disabled in Standby mode, this bit is equal to 0 after Standby or reset until the AVDEN bit is set.
func (r *registerCsr1Type) SetAvdo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsr1FieldAvdoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsr1FieldAvdoMask)
	}
}

// registerCr2Type This register is not reset by wakeup from Standby mode, RESET signal and VDD POR. It is only reset by VSW POR and VSWRST reset. This register shall not be accessed when VSWRST bit in RCC_BDCR register resets the VSW domain.After reset, PWR_CR2 register is write-protected. Prior to modifying its content, the DBP bit in PWR_CR1 register must be set to disable the write protection.
type registerCr2Type uint32

const (
	RegisterCr2FieldBrenShift = 0
	RegisterCr2FieldBrenMask  = 0x1
)

// GetBren Backup regulator enable When set, the Backup regulator (used to maintain the backup RAM content in Standby and VBAT modes) is enabled. If BREN is reset, the backup regulator is switched off. The backup RAM can still be used in Run and Stop modes. However, its content will be lost in Standby and VBAT modes. If BREN is set, the application must wait till the Backup Regulator Ready flag (BRRDY) is set to indicate that the data written into the SRAM will be maintained in Standby and VBAT modes.
func (r *registerCr2Type) GetBren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldBrenMask) != 0
}

// SetBren Backup regulator enable When set, the Backup regulator (used to maintain the backup RAM content in Standby and VBAT modes) is enabled. If BREN is reset, the backup regulator is switched off. The backup RAM can still be used in Run and Stop modes. However, its content will be lost in Standby and VBAT modes. If BREN is set, the application must wait till the Backup Regulator Ready flag (BRRDY) is set to indicate that the data written into the SRAM will be maintained in Standby and VBAT modes.
func (r *registerCr2Type) SetBren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldBrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldBrenMask)
	}
}

const (
	RegisterCr2FieldMonenShift = 4
	RegisterCr2FieldMonenMask  = 0x10
)

// GetMonen VBAT and temperature monitoring enable When set, the VBAT supply and temperature monitoring is enabled.
func (r *registerCr2Type) GetMonen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldMonenMask) != 0
}

// SetMonen VBAT and temperature monitoring enable When set, the VBAT supply and temperature monitoring is enabled.
func (r *registerCr2Type) SetMonen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldMonenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldMonenMask)
	}
}

const (
	RegisterCr2FieldBrrdyShift = 16
	RegisterCr2FieldBrrdyMask  = 0x10000
)

// GetBrrdy Backup regulator ready This bit is set by hardware to indicate that the Backup regulator is ready.
func (r *registerCr2Type) GetBrrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldBrrdyMask) != 0
}

const (
	RegisterCr2FieldVbatlShift = 20
	RegisterCr2FieldVbatlMask  = 0x100000
)

// GetVbatl VBAT level monitoring versus low threshold
func (r *registerCr2Type) GetVbatl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldVbatlMask) != 0
}

const (
	RegisterCr2FieldVbathShift = 21
	RegisterCr2FieldVbathMask  = 0x200000
)

// GetVbath VBAT level monitoring versus high threshold
func (r *registerCr2Type) GetVbath() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldVbathMask) != 0
}

const (
	RegisterCr2FieldTemplShift = 22
	RegisterCr2FieldTemplMask  = 0x400000
)

// GetTempl Temperature level monitoring versus low threshold
func (r *registerCr2Type) GetTempl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTemplMask) != 0
}

const (
	RegisterCr2FieldTemphShift = 23
	RegisterCr2FieldTemphMask  = 0x800000
)

// GetTemph Temperature level monitoring versus high threshold
func (r *registerCr2Type) GetTemph() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTemphMask) != 0
}

// registerCr3Type Reset only by POR only, not reset by wakeup from Standby mode and RESET pad. The lower byte of this register is written once after POR and shall be written before changing VOS level or ck_sys clock frequency. No limitation applies to the upper bytes.Programming data corresponding to an invalid combination of SDLEVEL, SDEXTHP, SDEN, LDOEN and BYPASS bits (see Table9) will be ignored: data will not be written, the written-once mechanism will lock the register and any further write access will be ignored. The default supply configuration will be kept and the ACTVOSRDY bit in PWR control status register 1 (PWR_CSR1) will go on indicating invalid voltage levels. The system shall be power cycled before writing a new value.
type registerCr3Type uint32

const (
	RegisterCr3FieldBypassShift = 0
	RegisterCr3FieldBypassMask  = 0x1
)

// GetBypass Power management unit bypass
func (r *registerCr3Type) GetBypass() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldBypassMask) != 0
}

// SetBypass Power management unit bypass
func (r *registerCr3Type) SetBypass(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldBypassMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldBypassMask)
	}
}

const (
	RegisterCr3FieldLdoenShift = 1
	RegisterCr3FieldLdoenMask  = 0x2
)

// GetLdoen Low drop-out regulator enable
func (r *registerCr3Type) GetLdoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldLdoenMask) != 0
}

// SetLdoen Low drop-out regulator enable
func (r *registerCr3Type) SetLdoen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldLdoenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldLdoenMask)
	}
}

const (
	RegisterCr3FieldSdenShift = 2
	RegisterCr3FieldSdenMask  = 0x4
)

// GetSden SD converter Enable
func (r *registerCr3Type) GetSden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldSdenMask) != 0
}

// SetSden SD converter Enable
func (r *registerCr3Type) SetSden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldSdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldSdenMask)
	}
}

const (
	RegisterCr3FieldVbeShift = 8
	RegisterCr3FieldVbeMask  = 0x100
)

// GetVbe VBAT charging enable
func (r *registerCr3Type) GetVbe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldVbeMask) != 0
}

// SetVbe VBAT charging enable
func (r *registerCr3Type) SetVbe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldVbeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldVbeMask)
	}
}

const (
	RegisterCr3FieldVbrsShift = 9
	RegisterCr3FieldVbrsMask  = 0x200
)

// GetVbrs VBAT charging resistor selection
func (r *registerCr3Type) GetVbrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldVbrsMask) != 0
}

// SetVbrs VBAT charging resistor selection
func (r *registerCr3Type) SetVbrs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldVbrsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldVbrsMask)
	}
}

const (
	RegisterCr3FieldUsb33denShift = 24
	RegisterCr3FieldUsb33denMask  = 0x1000000
)

// SetUsb33den VDD33USB voltage level detector enable.
func (r *registerCr3Type) SetUsb33den(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldUsb33denMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldUsb33denMask)
	}
}

const (
	RegisterCr3FieldUsbregenShift = 25
	RegisterCr3FieldUsbregenMask  = 0x2000000
)

// GetUsbregen USB regulator enable.
func (r *registerCr3Type) GetUsbregen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldUsbregenMask) != 0
}

// SetUsbregen USB regulator enable.
func (r *registerCr3Type) SetUsbregen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldUsbregenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldUsbregenMask)
	}
}

const (
	RegisterCr3FieldUsb33rdyShift = 26
	RegisterCr3FieldUsb33rdyMask  = 0x4000000
)

// GetUsb33rdy USB supply ready.
func (r *registerCr3Type) GetUsb33rdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldUsb33rdyMask) != 0
}

// registerCpucrType This register allows controlling CPU1 power.
type registerCpucrType uint32

const (
	RegisterCpucrFieldPddsd1Shift = 0
	RegisterCpucrFieldPddsd1Mask  = 0x1
)

// GetPddsd1 D1 domain Power Down Deepsleep selection. This bit allows CPU1 to define the Deepsleep mode for D1 domain.
func (r *registerCpucrType) GetPddsd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpucrFieldPddsd1Mask) != 0
}

// SetPddsd1 D1 domain Power Down Deepsleep selection. This bit allows CPU1 to define the Deepsleep mode for D1 domain.
func (r *registerCpucrType) SetPddsd1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpucrFieldPddsd1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpucrFieldPddsd1Mask)
	}
}

const (
	RegisterCpucrFieldPddsd2Shift = 1
	RegisterCpucrFieldPddsd2Mask  = 0x2
)

// GetPddsd2 D2 domain Power Down Deepsleep. This bit allows CPU1 to define the Deepsleep mode for D2 domain.
func (r *registerCpucrType) GetPddsd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpucrFieldPddsd2Mask) != 0
}

// SetPddsd2 D2 domain Power Down Deepsleep. This bit allows CPU1 to define the Deepsleep mode for D2 domain.
func (r *registerCpucrType) SetPddsd2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpucrFieldPddsd2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpucrFieldPddsd2Mask)
	}
}

const (
	RegisterCpucrFieldPddsd3Shift = 2
	RegisterCpucrFieldPddsd3Mask  = 0x4
)

// GetPddsd3 System D3 domain Power Down Deepsleep. This bit allows CPU1 to define the Deepsleep mode for System D3 domain.
func (r *registerCpucrType) GetPddsd3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpucrFieldPddsd3Mask) != 0
}

// SetPddsd3 System D3 domain Power Down Deepsleep. This bit allows CPU1 to define the Deepsleep mode for System D3 domain.
func (r *registerCpucrType) SetPddsd3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpucrFieldPddsd3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpucrFieldPddsd3Mask)
	}
}

const (
	RegisterCpucrFieldStopfShift = 5
	RegisterCpucrFieldStopfMask  = 0x20
)

// GetStopf STOP flag This bit is set by hardware and cleared only by any reset or by setting the CPU1 CSSF bit.
func (r *registerCpucrType) GetStopf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpucrFieldStopfMask) != 0
}

const (
	RegisterCpucrFieldSbfShift = 6
	RegisterCpucrFieldSbfMask  = 0x40
)

// GetSbf System Standby flag This bit is set by hardware and cleared only by a POR (Power-on Reset) or by setting the CPU1 CSSF bit
func (r *registerCpucrType) GetSbf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpucrFieldSbfMask) != 0
}

const (
	RegisterCpucrFieldSbfd1Shift = 7
	RegisterCpucrFieldSbfd1Mask  = 0x80
)

// GetSbfd1 D1 domain DStandby flag This bit is set by hardware and cleared by any system reset or by setting the CPU1 CSSF bit. Once set, this bit can be cleared only when the D1 domain is no longer in DStandby mode.
func (r *registerCpucrType) GetSbfd1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpucrFieldSbfd1Mask) != 0
}

const (
	RegisterCpucrFieldSbfd2Shift = 8
	RegisterCpucrFieldSbfd2Mask  = 0x100
)

// GetSbfd2 D2 domain DStandby flag This bit is set by hardware and cleared by any system reset or by setting the CPU1 CSSF bit. Once set, this bit can be cleared only when the D2 domain is no longer in DStandby mode.
func (r *registerCpucrType) GetSbfd2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpucrFieldSbfd2Mask) != 0
}

const (
	RegisterCpucrFieldCssfShift = 9
	RegisterCpucrFieldCssfMask  = 0x200
)

// GetCssf Clear D1 domain CPU1 Standby, Stop and HOLD flags (always read as 0) This bit is cleared to 0 by hardware.
func (r *registerCpucrType) GetCssf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpucrFieldCssfMask) != 0
}

// SetCssf Clear D1 domain CPU1 Standby, Stop and HOLD flags (always read as 0) This bit is cleared to 0 by hardware.
func (r *registerCpucrType) SetCssf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpucrFieldCssfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpucrFieldCssfMask)
	}
}

const (
	RegisterCpucrFieldRund3Shift = 11
	RegisterCpucrFieldRund3Mask  = 0x800
)

// GetRund3 Keep system D3 domain in Run mode regardless of the CPU sub-systems modes
func (r *registerCpucrType) GetRund3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCpucrFieldRund3Mask) != 0
}

// SetRund3 Keep system D3 domain in Run mode regardless of the CPU sub-systems modes
func (r *registerCpucrType) SetRund3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCpucrFieldRund3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCpucrFieldRund3Mask)
	}
}

// registerD3crType This register allows controlling D3 domain power.Following reset VOSRDY will be read 1 by software
type registerD3crType uint32

const (
	RegisterD3crFieldVosrdyShift = 13
	RegisterD3crFieldVosrdyMask  = 0x2000
)

// GetVosrdy VOS Ready bit for VCORE voltage scaling output selection. This bit is set to 1 by hardware when Bypass mode is selected in PWR control register 3 (PWR_CR3).
func (r *registerD3crType) GetVosrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3crFieldVosrdyMask) != 0
}

type RegisterD3crFieldVosEnumType uint8

const (
	// RegisterD3crFieldVosEnumScale3 (default)
	RegisterD3crFieldVosEnumScale3 RegisterD3crFieldVosEnumType = 0x1

	RegisterD3crFieldVosEnumScale2 RegisterD3crFieldVosEnumType = 0x2
	RegisterD3crFieldVosEnumScale1 RegisterD3crFieldVosEnumType = 0x3

	RegisterD3crFieldVosShift = 14
	RegisterD3crFieldVosMask  = 0xc000
)

// GetVos Voltage scaling selection according to performance These bits control the VCORE voltage level and allow to obtains the best trade-off between power consumption and performance: When increasing the performance, the voltage scaling shall be changed before increasing the system frequency. When decreasing performance, the system frequency shall first be decreased before changing the voltage scaling.
func (r *registerD3crType) GetVos() RegisterD3crFieldVosEnumType {
	return RegisterD3crFieldVosEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterD3crFieldVosMask) >> RegisterD3crFieldVosShift)
}

// SetVos Voltage scaling selection according to performance These bits control the VCORE voltage level and allow to obtains the best trade-off between power consumption and performance: When increasing the performance, the voltage scaling shall be changed before increasing the system frequency. When decreasing performance, the system frequency shall first be decreased before changing the voltage scaling.
func (r *registerD3crType) SetVos(value RegisterD3crFieldVosEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3crFieldVosMask)|(uint32(value)<<RegisterD3crFieldVosShift))
}

// registerWkupcrType reset only by system reset, not reset by wakeup from Standby mode5 wait states are required when writing this register (when clearing a WKUPF bit in PWR_WKUPFR, the AHB write access will complete after the WKUPF has been cleared).
type registerWkupcrType uint32

const (
	RegisterWkupcrFieldWkupcShift = 0
	RegisterWkupcrFieldWkupcMask  = 0x3f
)

// GetWkupc Clear Wakeup pin flag for WKUP. These bits are always read as 0.
func (r *registerWkupcrType) GetWkupc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWkupcrFieldWkupcMask) >> RegisterWkupcrFieldWkupcShift)
}

// SetWkupc Clear Wakeup pin flag for WKUP. These bits are always read as 0.
func (r *registerWkupcrType) SetWkupc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWkupcrFieldWkupcMask)|(uint32(value)<<RegisterWkupcrFieldWkupcShift))
}

// registerWkupfrType reset only by system reset, not reset by wakeup from Standby mode
type registerWkupfrType uint32

const (
	RegisterWkupfrFieldWkupf1Shift = 0
	RegisterWkupfrFieldWkupf1Mask  = 0x1
)

// GetWkupf1 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) GetWkupf1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupfrFieldWkupf1Mask) != 0
}

// SetWkupf1 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) SetWkupf1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupfrFieldWkupf1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupfrFieldWkupf1Mask)
	}
}

const (
	RegisterWkupfrFieldWkupf2Shift = 1
	RegisterWkupfrFieldWkupf2Mask  = 0x2
)

// GetWkupf2 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) GetWkupf2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupfrFieldWkupf2Mask) != 0
}

// SetWkupf2 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) SetWkupf2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupfrFieldWkupf2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupfrFieldWkupf2Mask)
	}
}

const (
	RegisterWkupfrFieldWkupf3Shift = 2
	RegisterWkupfrFieldWkupf3Mask  = 0x4
)

// GetWkupf3 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) GetWkupf3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupfrFieldWkupf3Mask) != 0
}

// SetWkupf3 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) SetWkupf3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupfrFieldWkupf3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupfrFieldWkupf3Mask)
	}
}

const (
	RegisterWkupfrFieldWkupf4Shift = 3
	RegisterWkupfrFieldWkupf4Mask  = 0x8
)

// GetWkupf4 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) GetWkupf4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupfrFieldWkupf4Mask) != 0
}

// SetWkupf4 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) SetWkupf4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupfrFieldWkupf4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupfrFieldWkupf4Mask)
	}
}

const (
	RegisterWkupfrFieldWkupf5Shift = 4
	RegisterWkupfrFieldWkupf5Mask  = 0x10
)

// GetWkupf5 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) GetWkupf5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupfrFieldWkupf5Mask) != 0
}

// SetWkupf5 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) SetWkupf5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupfrFieldWkupf5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupfrFieldWkupf5Mask)
	}
}

const (
	RegisterWkupfrFieldWkupf6Shift = 5
	RegisterWkupfrFieldWkupf6Mask  = 0x20
)

// GetWkupf6 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) GetWkupf6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupfrFieldWkupf6Mask) != 0
}

// SetWkupf6 Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).
func (r *registerWkupfrType) SetWkupf6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupfrFieldWkupf6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupfrFieldWkupf6Mask)
	}
}

// registerWkupeprType Reset only by system reset, not reset by wakeup from Standby mode
type registerWkupeprType uint32

const (
	RegisterWkupeprFieldWkupen1Shift = 0
	RegisterWkupeprFieldWkupen1Mask  = 0x1
)

// GetWkupen1 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) GetWkupen1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupen1Mask) != 0
}

// SetWkupen1 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) SetWkupen1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupen1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupen1Mask)
	}
}

const (
	RegisterWkupeprFieldWkupen2Shift = 1
	RegisterWkupeprFieldWkupen2Mask  = 0x2
)

// GetWkupen2 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) GetWkupen2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupen2Mask) != 0
}

// SetWkupen2 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) SetWkupen2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupen2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupen2Mask)
	}
}

const (
	RegisterWkupeprFieldWkupen3Shift = 2
	RegisterWkupeprFieldWkupen3Mask  = 0x4
)

// GetWkupen3 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) GetWkupen3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupen3Mask) != 0
}

// SetWkupen3 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) SetWkupen3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupen3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupen3Mask)
	}
}

const (
	RegisterWkupeprFieldWkupen4Shift = 3
	RegisterWkupeprFieldWkupen4Mask  = 0x8
)

// GetWkupen4 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) GetWkupen4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupen4Mask) != 0
}

// SetWkupen4 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) SetWkupen4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupen4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupen4Mask)
	}
}

const (
	RegisterWkupeprFieldWkupen5Shift = 4
	RegisterWkupeprFieldWkupen5Mask  = 0x10
)

// GetWkupen5 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) GetWkupen5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupen5Mask) != 0
}

// SetWkupen5 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) SetWkupen5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupen5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupen5Mask)
	}
}

const (
	RegisterWkupeprFieldWkupen6Shift = 5
	RegisterWkupeprFieldWkupen6Mask  = 0x20
)

// GetWkupen6 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) GetWkupen6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupen6Mask) != 0
}

// SetWkupen6 Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.
func (r *registerWkupeprType) SetWkupen6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupen6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupen6Mask)
	}
}

const (
	RegisterWkupeprFieldWkupp1Shift = 8
	RegisterWkupeprFieldWkupp1Mask  = 0x100
)

// GetWkupp1 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) GetWkupp1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupp1Mask) != 0
}

// SetWkupp1 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) SetWkupp1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupp1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupp1Mask)
	}
}

const (
	RegisterWkupeprFieldWkupp2Shift = 9
	RegisterWkupeprFieldWkupp2Mask  = 0x200
)

// GetWkupp2 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) GetWkupp2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupp2Mask) != 0
}

// SetWkupp2 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) SetWkupp2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupp2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupp2Mask)
	}
}

const (
	RegisterWkupeprFieldWkupp3Shift = 10
	RegisterWkupeprFieldWkupp3Mask  = 0x400
)

// GetWkupp3 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) GetWkupp3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupp3Mask) != 0
}

// SetWkupp3 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) SetWkupp3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupp3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupp3Mask)
	}
}

const (
	RegisterWkupeprFieldWkupp4Shift = 11
	RegisterWkupeprFieldWkupp4Mask  = 0x800
)

// GetWkupp4 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) GetWkupp4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupp4Mask) != 0
}

// SetWkupp4 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) SetWkupp4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupp4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupp4Mask)
	}
}

const (
	RegisterWkupeprFieldWkupp5Shift = 12
	RegisterWkupeprFieldWkupp5Mask  = 0x1000
)

// GetWkupp5 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) GetWkupp5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupp5Mask) != 0
}

// SetWkupp5 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) SetWkupp5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupp5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupp5Mask)
	}
}

const (
	RegisterWkupeprFieldWkupp6Shift = 13
	RegisterWkupeprFieldWkupp6Mask  = 0x2000
)

// GetWkupp6 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) GetWkupp6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkupp6Mask) != 0
}

// SetWkupp6 Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.
func (r *registerWkupeprType) SetWkupp6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterWkupeprFieldWkupp6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkupp6Mask)
	}
}

const (
	RegisterWkupeprFieldWkuppupd1Shift = 16
	RegisterWkupeprFieldWkuppupd1Mask  = 0x30000
)

// GetWkuppupd1 Wakeup pin pull configuration
func (r *registerWkupeprType) GetWkuppupd1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkuppupd1Mask) >> RegisterWkupeprFieldWkuppupd1Shift)
}

// SetWkuppupd1 Wakeup pin pull configuration
func (r *registerWkupeprType) SetWkuppupd1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkuppupd1Mask)|(uint32(value)<<RegisterWkupeprFieldWkuppupd1Shift))
}

const (
	RegisterWkupeprFieldWkuppupd2Shift = 18
	RegisterWkupeprFieldWkuppupd2Mask  = 0xc0000
)

// GetWkuppupd2 Wakeup pin pull configuration
func (r *registerWkupeprType) GetWkuppupd2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkuppupd2Mask) >> RegisterWkupeprFieldWkuppupd2Shift)
}

// SetWkuppupd2 Wakeup pin pull configuration
func (r *registerWkupeprType) SetWkuppupd2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkuppupd2Mask)|(uint32(value)<<RegisterWkupeprFieldWkuppupd2Shift))
}

const (
	RegisterWkupeprFieldWkuppupd3Shift = 20
	RegisterWkupeprFieldWkuppupd3Mask  = 0x300000
)

// GetWkuppupd3 Wakeup pin pull configuration
func (r *registerWkupeprType) GetWkuppupd3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkuppupd3Mask) >> RegisterWkupeprFieldWkuppupd3Shift)
}

// SetWkuppupd3 Wakeup pin pull configuration
func (r *registerWkupeprType) SetWkuppupd3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkuppupd3Mask)|(uint32(value)<<RegisterWkupeprFieldWkuppupd3Shift))
}

const (
	RegisterWkupeprFieldWkuppupd4Shift = 22
	RegisterWkupeprFieldWkuppupd4Mask  = 0xc00000
)

// GetWkuppupd4 Wakeup pin pull configuration
func (r *registerWkupeprType) GetWkuppupd4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkuppupd4Mask) >> RegisterWkupeprFieldWkuppupd4Shift)
}

// SetWkuppupd4 Wakeup pin pull configuration
func (r *registerWkupeprType) SetWkuppupd4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkuppupd4Mask)|(uint32(value)<<RegisterWkupeprFieldWkuppupd4Shift))
}

const (
	RegisterWkupeprFieldWkuppupd5Shift = 24
	RegisterWkupeprFieldWkuppupd5Mask  = 0x3000000
)

// GetWkuppupd5 Wakeup pin pull configuration
func (r *registerWkupeprType) GetWkuppupd5() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkuppupd5Mask) >> RegisterWkupeprFieldWkuppupd5Shift)
}

// SetWkuppupd5 Wakeup pin pull configuration
func (r *registerWkupeprType) SetWkuppupd5(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkuppupd5Mask)|(uint32(value)<<RegisterWkupeprFieldWkuppupd5Shift))
}

const (
	RegisterWkupeprFieldWkuppupd6Shift = 26
	RegisterWkupeprFieldWkuppupd6Mask  = 0xc000000
)

// GetWkuppupd6 Wakeup pin pull configuration for WKUP(truncate(n/2)-7) These bits define the I/O pad pull configuration used when WKUPEN(truncate(n/2)-7) = 1. The associated GPIO port pull configuration shall be set to the same value or to 00. The Wakeup pin pull configuration is kept in Standby mode.
func (r *registerWkupeprType) GetWkuppupd6() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterWkupeprFieldWkuppupd6Mask) >> RegisterWkupeprFieldWkuppupd6Shift)
}

// SetWkuppupd6 Wakeup pin pull configuration for WKUP(truncate(n/2)-7) These bits define the I/O pad pull configuration used when WKUPEN(truncate(n/2)-7) = 1. The associated GPIO port pull configuration shall be set to the same value or to 00. The Wakeup pin pull configuration is kept in Standby mode.
func (r *registerWkupeprType) SetWkuppupd6(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWkupeprFieldWkuppupd6Mask)|(uint32(value)<<RegisterWkupeprFieldWkuppupd6Shift))
}
