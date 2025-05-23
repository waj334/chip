#ifndef _PERIPHERALS_STK_TD
#define _PERIPHERALS_STK_TD

include "base.td"

def STKPeripheral : PeripheralType<"STK", "SysTick timer"> {
  let accessWidth = 32;
  let registers = [
    Register<"CSR", 0x0, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Counter enable">,
      Field<"TICKINT", 1, 1, ReadWrite, "SysTick exception request enable">,
      Field<"CLKSOURCE", 2, 1, ReadWrite, "Clock source selection">,
      Field<"COUNTFLAG", 16, 1, ReadWrite, "COUNTFLAG">,
    ], "SysTick control and status register">,
    Register<"RVR", 0x4, 32, [
      Field<"RELOAD", 0, 24, ReadWrite, "RELOAD value">,
    ], "SysTick reload value register">,
    Register<"CVR", 0x8, 32, [
      Field<"CURRENT", 0, 24, ReadWrite, "Current counter value">,
    ], "SysTick current value register">,
    Register<"CALIB", 0xc, 32, [
      Field<"TENMS", 0, 24, ReadWrite, "Calibration value">,
      Field<"SKEW", 30, 1, ReadWrite, "SKEW flag: Indicates whether the TENMS value is exact">,
      Field<"NOREF", 31, 1, ReadWrite, "NOREF flag. Reads as zero">,
    ], "SysTick calibration value register">,
  ];
}

class STKInstance<string Name, int Base> : PeripheralInstance<Name, Base, STKPeripheral>;
def STK : STKInstance<"STK", 0xe000e010>;

#endif // _PERIPHERALS_STK_TD
