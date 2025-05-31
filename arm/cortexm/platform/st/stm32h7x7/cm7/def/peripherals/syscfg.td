#ifndef _PERIPHERALS_SYSCFG_TD
#define _PERIPHERALS_SYSCFG_TD

include "base.td"

def SYSCFGPeripheral : PeripheralType<"SYSCFG", "System configuration controller"> {
  let accessWidth = 32;
  let registers = [
    Register<"PMCR", 0x4, 32, [
      Field<"I2C1FMP", 0, 1, ReadWrite, "I2C1 Fm+">,
      Field<"I2C2FMP", 1, 1, ReadWrite, "I2C2 Fm+">,
      Field<"I2C3FMP", 2, 1, ReadWrite, "I2C3 Fm+">,
      Field<"I2C4FMP", 3, 1, ReadWrite, "I2C4 Fm+">,
      Field<"PB6FMP", 4, 1, ReadWrite, "PB(6) Fm+">,
      Field<"PB7FMP", 5, 1, ReadWrite, "PB(7) Fast Mode Plus">,
      Field<"PB8FMP", 6, 1, ReadWrite, "PB(8) Fast Mode Plus">,
      Field<"PB9FMP", 7, 1, ReadWrite, "PB(9) Fm+">,
      Field<"BOOSTE", 8, 1, ReadWrite, "Booster Enable">,
      Field<"BOOSTVDDSEL", 9, 1, ReadWrite, "Analog switch supply voltage selection">,
      Field<"EPIS", 21, 3, ReadWrite, "Ethernet PHY Interface Selection">,
      Field<"PA0SO", 24, 1, ReadWrite, "PA0 Switch Open">,
      Field<"PA1SO", 25, 1, ReadWrite, "PA1 Switch Open">,
      Field<"PC2SO", 26, 1, ReadWrite, "PC2 Switch Open">,
      Field<"PC3SO", 27, 1, ReadWrite, "PC3 Switch Open">,
    ], "peripheral mode configuration register">,
    Register<"EXTICR1", 0x8, 32, [
      Field<"EXTI0", 0, 4, ReadWrite, "EXTI x configuration (x = 0 to 3)">,
      Field<"EXTI1", 4, 4, ReadWrite, "EXTI x configuration (x = 0 to 3)">,
      Field<"EXTI2", 8, 4, ReadWrite, "EXTI x configuration (x = 0 to 3)">,
      Field<"EXTI3", 12, 4, ReadWrite, "EXTI x configuration (x = 0 to 3)">,
    ], "external interrupt configuration register 1">,
    Register<"EXTICR2", 0xc, 32, [
      Field<"EXTI4", 0, 4, ReadWrite, "EXTI x configuration (x = 4 to 7)">,
      Field<"EXTI5", 4, 4, ReadWrite, "EXTI x configuration (x = 4 to 7)">,
      Field<"EXTI6", 8, 4, ReadWrite, "EXTI x configuration (x = 4 to 7)">,
      Field<"EXTI7", 12, 4, ReadWrite, "EXTI x configuration (x = 4 to 7)">,
    ], "external interrupt configuration register 2">,
    Register<"EXTICR3", 0x10, 32, [
      Field<"EXTI8", 0, 4, ReadWrite, "EXTI x configuration (x = 8 to 11)">,
      Field<"EXTI9", 4, 4, ReadWrite, "EXTI x configuration (x = 8 to 11)">,
      Field<"EXTI10", 8, 4, ReadWrite, "EXTI10">,
      Field<"EXTI11", 12, 4, ReadWrite, "EXTI x configuration (x = 8 to 11)">,
    ], "external interrupt configuration register 3">,
    Register<"EXTICR4", 0x14, 32, [
      Field<"EXTI12", 0, 4, ReadWrite, "EXTI x configuration (x = 12 to 15)">,
      Field<"EXTI13", 4, 4, ReadWrite, "EXTI x configuration (x = 12 to 15)">,
      Field<"EXTI14", 8, 4, ReadWrite, "EXTI x configuration (x = 12 to 15)">,
      Field<"EXTI15", 12, 4, ReadWrite, "EXTI x configuration (x = 12 to 15)">,
    ], "external interrupt configuration register 4">,
    Register<"CFGR", 0x18, 32, [
      Field<"CM4L", 0, 1, ReadWrite, "CM4L">,
      Field<"PVDL", 2, 1, ReadWrite, "PVDL">,
      Field<"FLASHL", 3, 1, ReadWrite, "FLASHL">,
      Field<"CM7L", 6, 1, ReadWrite, "CM7L">,
      Field<"BKRAML", 7, 1, ReadWrite, "BKRAML">,
      Field<"SRAM4L", 9, 1, ReadWrite, "SRAM4L">,
      Field<"SRAM3L", 10, 1, ReadWrite, "SRAM3L">,
      Field<"SRAM2L", 11, 1, ReadWrite, "SRAM2L">,
      Field<"SRAM1L", 12, 1, ReadWrite, "SRAM1L">,
      Field<"DTCML", 13, 1, ReadWrite, "DTCML">,
      Field<"ITCML", 14, 1, ReadWrite, "ITCML">,
      Field<"AXISRAML", 15, 1, ReadWrite, "AXISRAML">,
    ], "configuration register">,
    Register<"CCCSR", 0x20, 32, [
      Field<"EN", 0, 1, ReadWrite, "enable">,
      Field<"CS", 1, 1, ReadWrite, "Code selection">,
      Field<"READY", 8, 1, ReadWrite, "Compensation cell ready flag">,
      Field<"HSLV", 16, 1, ReadWrite, "High-speed at low-voltage">,
    ], "compensation cell control/status register">,
    Register<"CCVR", 0x24, 32, [
      Field<"NCV", 0, 4, ReadWrite, "NMOS compensation value">,
      Field<"PCV", 4, 4, ReadWrite, "PMOS compensation value">,
    ], "SYSCFG compensation cell value register">,
    Register<"CCCR", 0x28, 32, [
      Field<"NCC", 0, 4, ReadWrite, "NMOS compensation code">,
      Field<"PCC", 4, 4, ReadWrite, "PMOS compensation code">,
    ], "SYSCFG compensation cell code register">,
    Register<"PWRCR", 0x2c, 32, [
      Field<"ODEN", 0, 1, ReadWrite, "Overdrive enable">,
    ], "SYSCFG power control register">,
    Register<"PKGR", 0x124, 32, [
      Field<"PKG", 0, 4, ReadWrite, "Package">,
    ], "SYSCFG package register">,
    Register<"UR0", 0x300, 32, [
      Field<"BKS", 0, 1, ReadWrite, "Bank Swap">,
      Field<"RDP", 16, 8, ReadWrite, "Readout protection">,
    ], "SYSCFG user register 0">,
    Register<"UR1", 0x304, 32, [
      Field<"BCM4", 0, 1, ReadWrite, "Boot Cortex-M4">,
      Field<"BCM7", 16, 1, ReadWrite, "Boot Cortex-M7">,
    ], "SYSCFG user register 1">,
    Register<"UR2", 0x308, 32, [
      Field<"BORH", 0, 2, Read, "BOR_LVL Brownout Reset Threshold Level">,
      Field<"BCM7_ADD0", 16, 16, ReadWrite, "Cortex-M7 Boot Address 0">,
    ], "SYSCFG user register 2">,
    Register<"UR3", 0x30c, 32, [
      Field<"BCM4_ADD1", 0, 16, ReadWrite, "Cortex-M4 Boot Address 0">,
      Field<"BCM7_ADD1", 16, 16, ReadWrite, "Cortex-M7 Boot Address 1">,
    ], "SYSCFG user register 3">,
    Register<"UR4", 0x310, 32, [
      Field<"BCM4_ADD1", 0, 16, ReadWrite, "Mass Erase Protected Area Disabled for bank 1">,
      Field<"MEPAD_1", 16, 1, Read, "Boot Cortex-M4 Address 1">,
    ], "SYSCFG user register 4">,
    Register<"UR5", 0x314, 32, [
      Field<"MESAD_1", 0, 1, ReadWrite, "Mass erase secured area disabled for bank 1">,
      Field<"WRPS_1", 16, 8, ReadWrite, "Write protection for flash bank 1">,
    ], "SYSCFG user register 5">,
    Register<"UR6", 0x318, 32, [
      Field<"PA_BEG_1", 0, 12, ReadWrite, "Protected area start address for bank 1">,
      Field<"PA_END_1", 16, 12, ReadWrite, "Protected area end address for bank 1">,
    ], "SYSCFG user register 6">,
    Register<"UR7", 0x31c, 32, [
      Field<"SA_BEG_1", 0, 12, ReadWrite, "Secured area start address for bank 1">,
      Field<"SA_END_1", 16, 12, ReadWrite, "Secured area end address for bank 1">,
    ], "SYSCFG user register 7">,
    Register<"UR8", 0x320, 32, [
      Field<"MEPAD_2", 0, 1, ReadWrite, "Mass erase protected area disabled for bank 2">,
      Field<"MESAD_2", 16, 1, ReadWrite, "Mass erase secured area disabled for bank 2">,
    ], "SYSCFG user register 8">,
    Register<"UR9", 0x324, 32, [
      Field<"WRPS_2", 0, 8, ReadWrite, "Write protection for flash bank 2">,
      Field<"PA_BEG_2", 16, 12, ReadWrite, "Protected area start address for bank 2">,
    ], "SYSCFG user register 9">,
    Register<"UR10", 0x328, 32, [
      Field<"PA_END_2", 0, 12, ReadWrite, "Protected area end address for bank 2">,
      Field<"SA_BEG_2", 16, 12, ReadWrite, "Secured area start address for bank 2">,
    ], "SYSCFG user register 10">,
    Register<"UR11", 0x32c, 32, [
      Field<"SA_END_2", 0, 12, ReadWrite, "Secured area end address for bank 2">,
      Field<"IWDG1M", 16, 1, ReadWrite, "Independent Watchdog 1 mode">,
    ], "SYSCFG user register 11">,
    Register<"UR12", 0x330, 32, [
      Field<"IWDG2M", 0, 1, ReadWrite, "Independent Watchdog 2 mode">,
      Field<"SECURE", 16, 1, ReadWrite, "Secure mode">,
    ], "SYSCFG user register 12">,
    Register<"UR13", 0x334, 32, [
      Field<"SDRS", 0, 2, ReadWrite, "Secured DTCM RAM Size">,
      Field<"D1SBRST", 16, 1, ReadWrite, "D1 Standby reset">,
    ], "SYSCFG user register 13">,
    Register<"UR14", 0x338, 32, [
      Field<"D1STPRST", 0, 1, ReadWrite, "D1 Stop Reset">,
      Field<"D2SBRST", 16, 1, ReadWrite, "D2 Standby Reset">,
    ], "SYSCFG user register 14">,
    Register<"UR15", 0x33c, 32, [
      Field<"D2STPRST", 0, 1, ReadWrite, "D2 Stop Reset">,
      Field<"FZIWDGSTB", 16, 1, Read, "Freeze independent watchdog in Standby mode">,
    ], "SYSCFG user register 15">,
    Register<"UR16", 0x340, 32, [
      Field<"FZIWDGSTP", 0, 1, ReadWrite, "Freeze independent watchdog in Stop mode">,
      Field<"PKP", 16, 1, ReadWrite, "Private key programmed">,
    ], "SYSCFG user register 16">,
    Register<"UR17", 0x344, 32, [
      Field<"IO_HSLV", 0, 1, ReadWrite, "I/O high speed / low voltage">,
    ], "SYSCFG user register 17">,
  ];
}

class SYSCFGInstance<string Name, int Base> : PeripheralInstance<Name, Base, SYSCFGPeripheral>;
def SYSCFG : SYSCFGInstance<"SYSCFG", 0x58000400>;

#endif // _PERIPHERALS_SYSCFG_TD
