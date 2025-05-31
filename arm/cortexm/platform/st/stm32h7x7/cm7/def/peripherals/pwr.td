#ifndef _PERIPHERALS_PWR_TD
#define _PERIPHERALS_PWR_TD

include "base.td"

def PWRPeripheral : PeripheralType<"PWR", "PWR"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR1", 0x0, 32, [
      Field<"LPDS", 0, 1, ReadWrite, "Low-power Deepsleep with SVOS3 (SVOS4 and SVOS5 always use low-power, regardless of the setting of this bit)">,
      Field<"PVDE", 4, 1, ReadWrite, "Programmable voltage detector enable">,
      Field<"PLS", 5, 3, ReadWrite, "Programmable voltage detector level selection These bits select the voltage threshold detected by the PVD. Note: Refer to Section Electrical characteristics of the product datasheet for more details.">,
      Field<"DBP", 8, 1, ReadWrite, "Disable backup domain write protection In reset state, the RCC_BDCR register, the RTC registers (including the backup registers), BREN and MOEN bits in PWR_CR2 register, are protected against parasitic write access. This bit must be set to enable write access to these registers.">,
      Field<"FLPS", 9, 1, ReadWrite, "Flash low-power mode in DStop mode This bit allows to obtain the best trade-off between low-power consumption and restart time when exiting from DStop mode. When it is set, the Flash memory enters low-power mode when D1 domain is in DStop mode.">,
      Field<"SVOS", 14, 2, ReadWrite, "System Stop mode voltage scaling selection These bits control the VCORE voltage level in system Stop mode, to obtain the best trade-off between power consumption and performance.">,
      Field<"AVDEN", 16, 1, ReadWrite, "Peripheral voltage monitor on VDDA enable">,
      Field<"ALS", 17, 2, ReadWrite, "Analog voltage detector level selection These bits select the voltage threshold detected by the AVD.">,
    ], "PWR control register 1">,
    Register<"CSR1", 0x4, 32, [
      Field<"PVDO", 4, 1, ReadWrite, "Programmable voltage detect output This bit is set and cleared by hardware. It is valid only if the PVD has been enabled by the PVDE bit. Note: since the PVD is disabled in Standby mode, this bit is equal to 0 after Standby or reset until the PVDE bit is set.">,
      Field<"ACTVOSRDY", 13, 1, ReadWrite, "Voltage levels ready bit for currently used VOS and SDLEVEL This bit is set to 1 by hardware when the voltage regulator and the SD converter are both disabled and Bypass mode is selected in PWR control register 3 (PWR_CR3).">,
      Field<"ACTVOS", 14, 2, ReadWrite, "VOS currently applied for VCORE voltage scaling selection. These bits reflect the last VOS value applied to the PMU.">,
      Field<"AVDO", 16, 1, ReadWrite, "Analog voltage detector output on VDDA This bit is set and cleared by hardware. It is valid only if AVD on VDDA is enabled by the AVDEN bit. Note: Since the AVD is disabled in Standby mode, this bit is equal to 0 after Standby or reset until the AVDEN bit is set.">,
    ], "PWR control status register 1">,
    Register<"CR2", 0x8, 32, [
      Field<"BREN", 0, 1, ReadWrite, "Backup regulator enable When set, the Backup regulator (used to maintain the backup RAM content in Standby and VBAT modes) is enabled. If BREN is reset, the backup regulator is switched off. The backup RAM can still be used in Run and Stop modes. However, its content will be lost in Standby and VBAT modes. If BREN is set, the application must wait till the Backup Regulator Ready flag (BRRDY) is set to indicate that the data written into the SRAM will be maintained in Standby and VBAT modes.">,
      Field<"MONEN", 4, 1, ReadWrite, "VBAT and temperature monitoring enable When set, the VBAT supply and temperature monitoring is enabled.">,
      Field<"BRRDY", 16, 1, Read, "Backup regulator ready This bit is set by hardware to indicate that the Backup regulator is ready.">,
      Field<"VBATL", 20, 1, Read, "VBAT level monitoring versus low threshold">,
      Field<"VBATH", 21, 1, Read, "VBAT level monitoring versus high threshold">,
      Field<"TEMPL", 22, 1, Read, "Temperature level monitoring versus low threshold">,
      Field<"TEMPH", 23, 1, Read, "Temperature level monitoring versus high threshold">,
    ], "This register is not reset by wakeup from Standby mode, RESET signal and VDD POR. It is only reset by VSW POR and VSWRST reset. This register shall not be accessed when VSWRST bit in RCC_BDCR register resets the VSW domain.After reset, PWR_CR2 register is write-protected. Prior to modifying its content, the DBP bit in PWR_CR1 register must be set to disable the write protection.">,
    Register<"CR3", 0xc, 32, [
      Field<"BYPASS", 0, 1, ReadWrite, "Power management unit bypass">,
      Field<"LDOEN", 1, 1, ReadWrite, "Low drop-out regulator enable">,
      Field<"SDEN", 2, 1, ReadWrite, "SD converter Enable">,
      Field<"VBE", 8, 1, ReadWrite, "VBAT charging enable">,
      Field<"VBRS", 9, 1, ReadWrite, "VBAT charging resistor selection">,
      Field<"USB33DEN", 24, 1, Write, "VDD33USB voltage level detector enable.">,
      Field<"USBREGEN", 25, 1, ReadWrite, "USB regulator enable.">,
      Field<"USB33RDY", 26, 1, Read, "USB supply ready.">,
    ], "Reset only by POR only, not reset by wakeup from Standby mode and RESET pad. The lower byte of this register is written once after POR and shall be written before changing VOS level or ck_sys clock frequency. No limitation applies to the upper bytes.Programming data corresponding to an invalid combination of SDLEVEL, SDEXTHP, SDEN, LDOEN and BYPASS bits (see Table9) will be ignored: data will not be written, the written-once mechanism will lock the register and any further write access will be ignored. The default supply configuration will be kept and the ACTVOSRDY bit in PWR control status register 1 (PWR_CSR1) will go on indicating invalid voltage levels. The system shall be power cycled before writing a new value.">,
    Register<"CPUCR", 0x10, 32, [
      Field<"PDDS_D1", 0, 1, ReadWrite, "D1 domain Power Down Deepsleep selection. This bit allows CPU1 to define the Deepsleep mode for D1 domain.">,
      Field<"PDDS_D2", 1, 1, ReadWrite, "D2 domain Power Down Deepsleep. This bit allows CPU1 to define the Deepsleep mode for D2 domain.">,
      Field<"PDDS_D3", 2, 1, ReadWrite, "System D3 domain Power Down Deepsleep. This bit allows CPU1 to define the Deepsleep mode for System D3 domain.">,
      Field<"STOPF", 5, 1, Read, "STOP flag This bit is set by hardware and cleared only by any reset or by setting the CPU1 CSSF bit.">,
      Field<"SBF", 6, 1, Read, "System Standby flag This bit is set by hardware and cleared only by a POR (Power-on Reset) or by setting the CPU1 CSSF bit">,
      Field<"SBF_D1", 7, 1, Read, "D1 domain DStandby flag This bit is set by hardware and cleared by any system reset or by setting the CPU1 CSSF bit. Once set, this bit can be cleared only when the D1 domain is no longer in DStandby mode.">,
      Field<"SBF_D2", 8, 1, Read, "D2 domain DStandby flag This bit is set by hardware and cleared by any system reset or by setting the CPU1 CSSF bit. Once set, this bit can be cleared only when the D2 domain is no longer in DStandby mode.">,
      Field<"CSSF", 9, 1, ReadWrite, "Clear D1 domain CPU1 Standby, Stop and HOLD flags (always read as 0) This bit is cleared to 0 by hardware.">,
      Field<"RUN_D3", 11, 1, ReadWrite, "Keep system D3 domain in Run mode regardless of the CPU sub-systems modes">,
    ], "This register allows controlling CPU1 power.">,
    Register<"D3CR", 0x18, 32, [
      Field<"VOSRDY", 13, 1, Read, "VOS Ready bit for VCORE voltage scaling output selection. This bit is set to 1 by hardware when Bypass mode is selected in PWR control register 3 (PWR_CR3).">,
      Field<"VOS", 14, 2, ReadWrite, "Voltage scaling selection according to performance These bits control the VCORE voltage level and allow to obtains the best trade-off between power consumption and performance: When increasing the performance, the voltage scaling shall be changed before increasing the system frequency. When decreasing performance, the system frequency shall first be decreased before changing the voltage scaling.", [
        Enum<"Scale3", 1, "(default)">,
        Enum<"Scale2", 2>,
        Enum<"Scale1", 3>,
      ]>,
    ], "This register allows controlling D3 domain power.Following reset VOSRDY will be read 1 by software">,
    Register<"WKUPCR", 0x20, 32, [
      Field<"WKUPC", 0, 6, ReadWrite, "Clear Wakeup pin flag for WKUP. These bits are always read as 0.">,
    ], "reset only by system reset, not reset by wakeup from Standby mode5 wait states are required when writing this register (when clearing a WKUPF bit in PWR_WKUPFR, the AHB write access will complete after the WKUPF has been cleared).">,
    Register<"WKUPFR", 0x24, 32, [
      Field<"WKUPF1", 0, 1, ReadWrite, "Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).">,
      Field<"WKUPF2", 1, 1, ReadWrite, "Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).">,
      Field<"WKUPF3", 2, 1, ReadWrite, "Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).">,
      Field<"WKUPF4", 3, 1, ReadWrite, "Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).">,
      Field<"WKUPF5", 4, 1, ReadWrite, "Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).">,
      Field<"WKUPF6", 5, 1, ReadWrite, "Wakeup pin WKUPF flag. This bit is set by hardware and cleared only by a Reset pin or by setting the WKUPCn+1 bit in the PWR wakeup clear register (PWR_WKUPCR).">,
    ], "reset only by system reset, not reset by wakeup from Standby mode">,
    Register<"WKUPEPR", 0x28, 32, [
      Field<"WKUPEN1", 0, 1, ReadWrite, "Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.">,
      Field<"WKUPEN2", 1, 1, ReadWrite, "Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.">,
      Field<"WKUPEN3", 2, 1, ReadWrite, "Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.">,
      Field<"WKUPEN4", 3, 1, ReadWrite, "Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.">,
      Field<"WKUPEN5", 4, 1, ReadWrite, "Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.">,
      Field<"WKUPEN6", 5, 1, ReadWrite, "Enable Wakeup Pin WKUPn+1 Each bit is set and cleared by software. Note: An additional wakeup event is detected if WKUPn+1 pin is enabled (by setting the WKUPENn+1 bit) when WKUPn+1 pin level is already high when WKUPPn+1 selects rising edge, or low when WKUPPn+1 selects falling edge.">,
      Field<"WKUPP1", 8, 1, ReadWrite, "Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.">,
      Field<"WKUPP2", 9, 1, ReadWrite, "Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.">,
      Field<"WKUPP3", 10, 1, ReadWrite, "Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.">,
      Field<"WKUPP4", 11, 1, ReadWrite, "Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.">,
      Field<"WKUPP5", 12, 1, ReadWrite, "Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.">,
      Field<"WKUPP6", 13, 1, ReadWrite, "Wakeup pin polarity bit for WKUPn-7 These bits define the polarity used for event detection on WKUPn-7 external wakeup pin.">,
      Field<"WKUPPUPD1", 16, 2, ReadWrite, "Wakeup pin pull configuration">,
      Field<"WKUPPUPD2", 18, 2, ReadWrite, "Wakeup pin pull configuration">,
      Field<"WKUPPUPD3", 20, 2, ReadWrite, "Wakeup pin pull configuration">,
      Field<"WKUPPUPD4", 22, 2, ReadWrite, "Wakeup pin pull configuration">,
      Field<"WKUPPUPD5", 24, 2, ReadWrite, "Wakeup pin pull configuration">,
      Field<"WKUPPUPD6", 26, 2, ReadWrite, "Wakeup pin pull configuration for WKUP(truncate(n/2)-7) These bits define the I/O pad pull configuration used when WKUPEN(truncate(n/2)-7) = 1. The associated GPIO port pull configuration shall be set to the same value or to 00. The Wakeup pin pull configuration is kept in Standby mode.">,
    ], "Reset only by system reset, not reset by wakeup from Standby mode">,
  ];
}

class PWRInstance<string Name, int Base> : PeripheralInstance<Name, Base, PWRPeripheral>;
def PWR : PWRInstance<"PWR", 0x58024800>;

#endif // _PERIPHERALS_PWR_TD
