#ifndef _PERIPHERALS_FPU_TD
#define _PERIPHERALS_FPU_TD

include "base.td"

def FPUPeripheral : PeripheralType<"FPU", "Floating-point and MVE Extension"> {
  let accessWidth = 32;
  let registers = [
    Register<"FPCCR", 0x0, 32, [
      Field<"LSPACT", 0, 1, ReadWrite, "Lazy state preservation active">,
      Field<"USER", 1, 1, ReadWrite, "Indicates the privilege level of the software executing, when the processor allocated the floating point stack">,
      Field<"S", 2, 1, ReadWrite, "Security status of the floating point context">,
      Field<"THREAD", 3, 1, ReadWrite, "Thread mode">,
      Field<"HFRDY", 4, 1, ReadWrite, "HardFault ready">,
      Field<"MMRDY", 5, 1, ReadWrite, "MemManage ready">,
      Field<"BFRDY", 6, 1, ReadWrite, "BusFault ready">,
      Field<"SFRDY", 7, 1, ReadWrite, "SecureFault ready">,
      Field<"MONRDY", 8, 1, ReadWrite, "DebugMonitor ready">,
      Field<"SPLIMVIOL", 9, 1, ReadWrite, "Stack pointer limit violation">,
      Field<"UFRDY", 10, 1, ReadWrite, "UsageFault ready">,
      Field<"TS", 26, 1, ReadWrite, "Treat as Secure">,
      Field<"CLRONRETS", 27, 1, ReadWrite, "Clear on return Secure only">,
      Field<"CLRONRET", 28, 1, ReadWrite, "Clear on return">,
      Field<"LSPENS", 29, 1, ReadWrite, "Lazy state preservation enable Secure only">,
      Field<"LSPEN", 30, 1, ReadWrite, "Automatic state preservation enable">,
      Field<"ASPEN", 31, 1, ReadWrite, "Automatic state preservation enable">,
    ], "Holds control data for the Floating Point Unit">,
    Register<"FPCAR", 0x4, 32, [
    ], "Holds the location of the unpopulated floating-point register space allocated on an exception stack frame">,
    Register<"FPDSCR", 0x8, 32, [
      Field<"LTPSIZE", 16, 3, ReadWrite, "Default value for FPSCR.LTPSIZE">,
      Field<"FZ16", 19, 1, ReadWrite, "Default value for FPSCR.FZ16">,
      Field<"RMODE", 22, 2, ReadWrite, "Default value for FPSCR.RMode">,
      Field<"FZ", 24, 1, ReadWrite, "Default value for FPSCR.FZ">,
      Field<"DN", 25, 1, ReadWrite, "Default value for FPSCR.DN">,
      Field<"AHP", 26, 1, ReadWrite, "Default value for FPSCR.AHP">,
    ], "Holds the default values for the floating-point status control data that the processor assigns to the FPSCR when it creates a new floating-point context">,
    Register<"MVFR0", 0xc, 32, [
      Field<"SIMDREG", 0, 4, ReadWrite, "SIMD registers">,
      Field<"FPSP", 4, 4, ReadWrite, "Floating-point single-precision">,
      Field<"FPDP", 8, 4, ReadWrite, "Floating-point double-precision">,
      Field<"FPDIVIDE", 16, 4, ReadWrite, "Floating-point divide">,
      Field<"FPSQRT", 20, 4, ReadWrite, "Floating-point square root">,
      Field<"FPROUND", 28, 4, ReadWrite, "Floating-point rounding modes">,
    ], "Describes the features provided by the Floating-point extension">,
    Register<"MVFR1", 0x10, 32, [
      Field<"FPFTZ", 0, 4, ReadWrite, "Floating-point flush-to-zero">,
      Field<"FPDNAN", 4, 4, ReadWrite, "Floating-point default NaN">,
      Field<"MVE", 8, 4, ReadWrite, "Indicates support for MVE">,
      Field<"FP16", 20, 4, ReadWrite, "Floating-point half-precision data processing">,
      Field<"FPHP", 24, 4, ReadWrite, "Floating-point half-precision conversion">,
      Field<"FMAC", 28, 4, ReadWrite, "Fused multiply accumulate instructions">,
    ], "Describes the features provided by the Floating-point extension">,
    Register<"MVFR2", 0x14, 32, [
      Field<"FPMISC", 4, 4, ReadWrite, "Floating-point miscellaneous">,
    ], "Describes the features provided by the floating-point extension">,
  ];
}

class FPUInstance<string Name, int Base> : PeripheralInstance<Name, Base, FPUPeripheral>;
def FPU : FPUInstance<"FPU", 0xe000ef34>;

#endif // _PERIPHERALS_FPU_TD
