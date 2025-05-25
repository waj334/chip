#ifndef _PERIPHERALS_CAN_CCU_TD
#define _PERIPHERALS_CAN_CCU_TD

include "base.td"

def CAN_CCUPeripheral : PeripheralType<"CAN_CCU", "CCU registers"> {
  let accessWidth = 32;
  let registers = [
    Register<"CREL", 0x0, 32, [
      Field<"DAY", 0, 8, ReadWrite, "Time Stamp Day">,
      Field<"MON", 8, 8, ReadWrite, "Time Stamp Month">,
      Field<"YEAR", 16, 4, ReadWrite, "Time Stamp Year">,
      Field<"SUBSTEP", 20, 4, ReadWrite, "Sub-step of Core Release">,
      Field<"STEP", 24, 4, ReadWrite, "Step of Core Release">,
      Field<"REL", 28, 4, ReadWrite, "Core Release">,
    ], "Clock Calibration Unit Core Release Register">,
    Register<"CCFG", 0x4, 32, [
      Field<"TQBT", 0, 5, ReadWrite, "Time Quanta per Bit Time">,
      Field<"BCC", 6, 1, ReadWrite, "Bypass Clock Calibration">,
      Field<"CFL", 7, 1, ReadWrite, "Calibration Field Length">,
      Field<"OCPM", 8, 8, ReadWrite, "Oscillator Clock Periods Minimum">,
      Field<"CDIV", 16, 4, ReadWrite, "Clock Divider">,
      Field<"SWR", 31, 1, ReadWrite, "Software Reset">,
    ], "Calibration Configuration Register">,
    Register<"CSTAT", 0x8, 32, [
      Field<"OCPC", 0, 18, ReadWrite, "Oscillator Clock Period Counter">,
      Field<"TQC", 18, 11, ReadWrite, "Time Quanta Counter">,
      Field<"CALS", 30, 2, ReadWrite, "Calibration State">,
    ], "Calibration Status Register">,
    Register<"CWD", 0xc, 32, [
      Field<"WDC", 0, 16, ReadWrite, "WDC">,
      Field<"WDV", 16, 16, ReadWrite, "WDV">,
    ], "Calibration Watchdog Register">,
    Register<"IR", 0x10, 32, [
      Field<"CWE", 0, 1, ReadWrite, "Calibration Watchdog Event">,
      Field<"CSC", 1, 1, ReadWrite, "Calibration State Changed">,
    ], "Clock Calibration Unit Interrupt Register">,
    Register<"IE", 0x14, 32, [
      Field<"CWEE", 0, 1, ReadWrite, "Calibration Watchdog Event Enable">,
      Field<"CSCE", 1, 1, ReadWrite, "Calibration State Changed Enable">,
    ], "Clock Calibration Unit Interrupt Enable Register">,
  ];
}

class CAN_CCUInstance<string Name, int Base> : PeripheralInstance<Name, Base, CAN_CCUPeripheral>;
def CAN_CCU : CAN_CCUInstance<"CAN_CCU", 0x4000a800>;

#endif // _PERIPHERALS_CAN_CCU_TD
