#ifndef _PERIPHERALS_SYSTICK_TD
#define _PERIPHERALS_SYSTICK_TD

include "base.td"

def SYSTICKPeripheral : PeripheralType<"SYSTICK", "SysTick Timer"> {
  let accessWidth = 32;
  let registers = [
    Register<"CSR", 0x0, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Enables the counter">,
      Field<"TICKINT", 1, 1, ReadWrite, "Enables SysTick exception request">,
      Field<"CLKSOURCE", 2, 1, ReadWrite, "Selects the SysTick timer clock source">,
      Field<"COUNTFLAG", 16, 1, ReadWrite, "Returns 1 if timer counted to 0 since the last read of this register">,
    ], "Controls and provides status date for the SysTick timer">,
    Register<"RVR", 0x4, 32, [
      Field<"RELOAD", 0, 24, ReadWrite, "Value to load into the SYST_CVR when the counter is enabled and when it reaches 0">,
    ], "Specifies the SysTick timer counter reload value">,
    Register<"CVR", 0x8, 32, [
      Field<"CURRENT", 0, 24, ReadWrite, "Current value of the SysTick counter. A write of any value clears the field to 0, and also clears the SYST_CSR.COUNTFLAG bit to 0.">,
    ], "Contains the current value of the SysTick counter">,
    Register<"CALIB", 0xc, 32, [
      Field<"TENMS", 0, 24, ReadWrite, "Reload value for 10ms (100Hz) timing, subject to system clock skew errors. If the value reads as zero, the calibration value is not known.">,
      Field<"SKEW", 30, 1, ReadWrite, "Indicates whether the TENMS value is exact">,
      Field<"NOREF", 31, 1, ReadWrite, "Indicates whether the device provides a reference clock to the processor">,
    ], "Indicates the SysTick calibration value and parameters for the selected Security state">,
  ];
}

class SYSTICKInstance<string Name, int Base> : PeripheralInstance<Name, Base, SYSTICKPeripheral>;
def SYSTICK : SYSTICKInstance<"SYSTICK", 0xe000e010>;

#endif // _PERIPHERALS_SYSTICK_TD
