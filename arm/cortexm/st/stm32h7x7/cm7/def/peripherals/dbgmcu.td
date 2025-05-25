#ifndef _PERIPHERALS_DBGMCU_TD
#define _PERIPHERALS_DBGMCU_TD

include "base.td"

def DBGMCUPeripheral : PeripheralType<"DBGMCU", "Microcontroller Debug Unit"> {
  let accessWidth = 32;
  let registers = [
    Register<"IDC", 0x0, 32, [
      Field<"DEV_ID", 0, 12, ReadWrite, "Device ID">,
      Field<"REV_ID", 16, 16, ReadWrite, "Revision">,
    ], "DBGMCU Identity Code Register">,
    Register<"CR", 0x4, 32, [
      Field<"DBGSLPD1", 0, 1, ReadWrite, "Allow D1 domain debug in Sleep mode">,
      Field<"DBGSTPD1", 1, 1, ReadWrite, "Allow D1 domain debug in Stop mode">,
      Field<"DBGSTBD1", 2, 1, ReadWrite, "Allow D1 domain debug in Standby mode">,
      Field<"DBGSLPD2", 3, 1, ReadWrite, "Allow D2 domain debug in Sleep mode">,
      Field<"DBGSTPD2", 4, 1, ReadWrite, "Allow D2 domain debug in Stop mode">,
      Field<"DBGSTBD2", 5, 1, ReadWrite, "Allow D2 domain debug in Standby mode">,
      Field<"DBGSTPD3", 7, 1, ReadWrite, "Allow debug in D3 Stop mode">,
      Field<"DBGSTBD3", 8, 1, ReadWrite, "Allow debug in D3 Standby mode">,
      Field<"TRACECLKEN", 20, 1, ReadWrite, "Trace port clock enable">,
      Field<"D1DBGCKEN", 21, 1, ReadWrite, "D1 debug clock enable">,
      Field<"D3DBGCKEN", 22, 1, ReadWrite, "D3 debug clock enable">,
      Field<"TRGOEN", 28, 1, ReadWrite, "External trigger output enable">,
    ], "DBGMCU Configuration Register">,
    Register<"APB3FZ1", 0x34, 32, [
      Field<"WWDG1", 6, 1, ReadWrite, "WWDG1 stop in debug">,
    ], "DBGMCU APB3 peripheral freeze register CPU1">,
    Register<"APB3FZ2", 0x38, 32, [
      Field<"WWDG1", 6, 1, ReadWrite, "WWDG1 stop in debug">,
    ], "DBGMCU APB3 peripheral freeze register CPU2">,
    Register<"APB1LFZ1", 0x3c, 32, [
      Field<"DBG_TIM2", 0, 1, ReadWrite, "TIM2 stop in debug">,
      Field<"DBG_TIM3", 1, 1, ReadWrite, "TIM3 stop in debug">,
      Field<"DBG_TIM4", 2, 1, ReadWrite, "TIM4 stop in debug">,
      Field<"DBG_TIM5", 3, 1, ReadWrite, "TIM5 stop in debug">,
      Field<"DBG_TIM6", 4, 1, ReadWrite, "TIM6 stop in debug">,
      Field<"DBG_TIM7", 5, 1, ReadWrite, "TIM7 stop in debug">,
      Field<"DBG_TIM12", 6, 1, ReadWrite, "TIM12 stop in debug">,
      Field<"DBG_TIM13", 7, 1, ReadWrite, "TIM13 stop in debug">,
      Field<"DBG_TIM14", 8, 1, ReadWrite, "TIM14 stop in debug">,
      Field<"DBG_LPTIM1", 9, 1, ReadWrite, "LPTIM1 stop in debug">,
      Field<"DBG_WWDG2", 11, 1, ReadWrite, "WWDG2 stop in debug">,
      Field<"DBG_I2C1", 21, 1, ReadWrite, "I2C1 SMBUS timeout stop in debug">,
      Field<"DBG_I2C2", 22, 1, ReadWrite, "I2C2 SMBUS timeout stop in debug">,
      Field<"DBG_I2C3", 23, 1, ReadWrite, "I2C3 SMBUS timeout stop in debug">,
    ], "DBGMCU APB1L peripheral freeze register">,
    Register<"APB1LFZ2", 0x40, 32, [
      Field<"DBG_TIM2", 0, 1, ReadWrite, "TIM2 stop in debug">,
      Field<"DBG_TIM3", 1, 1, ReadWrite, "TIM3 stop in debug">,
      Field<"DBG_TIM4", 2, 1, ReadWrite, "TIM4 stop in debug">,
      Field<"DBG_TIM5", 3, 1, ReadWrite, "TIM5 stop in debug">,
      Field<"DBG_TIM6", 4, 1, ReadWrite, "TIM6 stop in debug">,
      Field<"DBG_TIM7", 5, 1, ReadWrite, "TIM4 stop in debug">,
      Field<"DBG_TIM12", 6, 1, ReadWrite, "TIM12 stop in debug">,
      Field<"DBG_TIM13", 7, 1, ReadWrite, "TIM13 stop in debug">,
      Field<"DBG_TIM14", 8, 1, ReadWrite, "TIM14 stop in debug">,
      Field<"DBG_LPTIM1", 9, 1, ReadWrite, "LPTIM1 stop in debug">,
      Field<"DBG_WWDG2", 11, 1, ReadWrite, "WWDG2 stop in debug">,
      Field<"DBG_I2C1", 21, 1, ReadWrite, "I2C1 SMBUS timeout stop in debug">,
      Field<"DBG_I2C2", 22, 1, ReadWrite, "I2C2 SMBUS timeout stop in debug">,
      Field<"DBG_I2C3", 23, 1, ReadWrite, "I2C3 SMBUS timeout stop in debug">,
    ], "DBGMCU APB1L peripheral freeze register CPU2">,
    Register<"APB2FZ1", 0x4c, 32, [
      Field<"DBG_TIM1", 0, 1, ReadWrite, "TIM1 stop in debug">,
      Field<"DBG_TIM8", 1, 1, ReadWrite, "TIM8 stop in debug">,
      Field<"DBG_TIM15", 16, 1, ReadWrite, "TIM15 stop in debug">,
      Field<"DBG_TIM16", 17, 1, ReadWrite, "TIM16 stop in debug">,
      Field<"DBG_TIM17", 18, 1, ReadWrite, "TIM17 stop in debug">,
      Field<"DBG_HRTIM", 29, 1, ReadWrite, "HRTIM stop in debug">,
    ], "DBGMCU APB2 peripheral freeze register">,
    Register<"APB2FZ2", 0x50, 32, [
      Field<"DBG_TIM1", 0, 1, ReadWrite, "TIM1 stop in debug">,
      Field<"DBG_TIM8", 1, 1, ReadWrite, "TIM8 stop in debug">,
      Field<"DBG_TIM15", 16, 1, ReadWrite, "TIM15 stop in debug">,
      Field<"DBG_TIM16", 17, 1, ReadWrite, "TIM16 stop in debug">,
      Field<"DBG_TIM17", 18, 1, ReadWrite, "TIM17 stop in debug">,
      Field<"DBG_HRTIM", 29, 1, ReadWrite, "HRTIM stop in debug">,
    ], "DBGMCU APB2 peripheral freeze register CPU2">,
    Register<"APB4FZ1", 0x54, 32, [
      Field<"DBG_I2C4", 7, 1, ReadWrite, "I2C4 SMBUS timeout stop in debug">,
      Field<"DBG_LPTIM2", 9, 1, ReadWrite, "LPTIM2 stop in debug">,
      Field<"DBG_LPTIM3", 10, 1, ReadWrite, "LPTIM2 stop in debug">,
      Field<"DBG_LPTIM4", 11, 1, ReadWrite, "LPTIM4 stop in debug">,
      Field<"DBG_LPTIM5", 12, 1, ReadWrite, "LPTIM5 stop in debug">,
      Field<"DBG_RTC", 16, 1, ReadWrite, "RTC stop in debug">,
      Field<"DBG_WDGLSD1", 18, 1, ReadWrite, "Independent watchdog for D1 stop in debug">,
      Field<"DBG_WDGLSD2", 19, 1, ReadWrite, "Independent watchdog for D2 stop in debug">,
    ], "DBGMCU APB4 peripheral freeze register">,
    Register<"APB4FZ2", 0x58, 32, [
      Field<"DBG_I2C4", 7, 1, ReadWrite, "I2C4 SMBUS timeout stop in debug">,
      Field<"DBG_LPTIM2", 9, 1, ReadWrite, "LPTIM2 stop in debug">,
      Field<"DBG_LPTIM3", 10, 1, ReadWrite, "LPTIM2 stop in debug">,
      Field<"DBG_LPTIM4", 11, 1, ReadWrite, "LPTIM4 stop in debug">,
      Field<"DBG_LPTIM5", 12, 1, ReadWrite, "LPTIM5 stop in debug">,
      Field<"DBG_RTC", 16, 1, ReadWrite, "RTC stop in debug">,
      Field<"DBG_WDGLSD1", 18, 1, ReadWrite, "LS watchdog for D1 stop in debug">,
      Field<"DBG_WDGLSD2", 19, 1, ReadWrite, "LS watchdog for D2 stop in debug">,
    ], "DBGMCU APB4 peripheral freeze register CPU2">,
  ];
}

class DBGMCUInstance<string Name, int Base> : PeripheralInstance<Name, Base, DBGMCUPeripheral>;
def DBGMCU : DBGMCUInstance<"DBGMCU", 0x5c001000>;

#endif // _PERIPHERALS_DBGMCU_TD
