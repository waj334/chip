#ifndef _PERIPHERALS_FPU_TD
#define _PERIPHERALS_FPU_TD

include "base.td"

def FPUPeripheral : PeripheralType<"FPU", "Floting point unit"> {
  let accessWidth = 32;
  let registers = [
    Register<"FPCCR", 0x0, 32, [
      Field<"LSPACT", 0, 1, ReadWrite, "LSPACT">,
      Field<"USER", 1, 1, ReadWrite, "USER">,
      Field<"THREAD", 3, 1, ReadWrite, "THREAD">,
      Field<"HFRDY", 4, 1, ReadWrite, "HFRDY">,
      Field<"MMRDY", 5, 1, ReadWrite, "MMRDY">,
      Field<"BFRDY", 6, 1, ReadWrite, "BFRDY">,
      Field<"MONRDY", 8, 1, ReadWrite, "MONRDY">,
      Field<"LSPEN", 30, 1, ReadWrite, "LSPEN">,
      Field<"ASPEN", 31, 1, ReadWrite, "ASPEN">,
    ], "Floating-point context control register">,
    Register<"FPCAR", 0x4, 32, [
      Field<"ADDRESS", 3, 29, ReadWrite, "Location of unpopulated floating-point">,
    ], "Floating-point context address register">,
    Register<"FPSCR", 0x8, 32, [
      Field<"IOC", 0, 1, ReadWrite, "Invalid operation cumulative exception bit">,
      Field<"DZC", 1, 1, ReadWrite, "Division by zero cumulative exception bit.">,
      Field<"OFC", 2, 1, ReadWrite, "Overflow cumulative exception bit">,
      Field<"UFC", 3, 1, ReadWrite, "Underflow cumulative exception bit">,
      Field<"IXC", 4, 1, ReadWrite, "Inexact cumulative exception bit">,
      Field<"IDC", 7, 1, ReadWrite, "Input denormal cumulative exception bit.">,
      Field<"RMode", 22, 2, ReadWrite, "Rounding Mode control field">,
      Field<"FZ", 24, 1, ReadWrite, "Flush-to-zero mode control bit:">,
      Field<"DN", 25, 1, ReadWrite, "Default NaN mode control bit">,
      Field<"AHP", 26, 1, ReadWrite, "Alternative half-precision control bit">,
      Field<"V", 28, 1, ReadWrite, "Overflow condition code flag">,
      Field<"C", 29, 1, ReadWrite, "Carry condition code flag">,
      Field<"Z", 30, 1, ReadWrite, "Zero condition code flag">,
      Field<"N", 31, 1, ReadWrite, "Negative condition code flag">,
    ], "Floating-point status control register">,
  ];
}

class FPUInstance<string Name, int Base> : PeripheralInstance<Name, Base, FPUPeripheral>;
def FPU : FPUInstance<"FPU", 0xe000ef34>;

#endif // _PERIPHERALS_FPU_TD
