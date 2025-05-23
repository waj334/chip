#ifndef _PERIPHERALS_FPU_CPACR_TD
#define _PERIPHERALS_FPU_CPACR_TD

include "base.td"

def FPU_CPACRPeripheral : PeripheralType<"FPU_CPACR", "Floating point unit CPACR"> {
  let accessWidth = 32;
  let registers = [
    Register<"CPACR", 0x0, 32, [
      Field<"CP", 20, 4, ReadWrite, "CP">,
    ], "Coprocessor access control register">,
  ];
}

class FPU_CPACRInstance<string Name, int Base> : PeripheralInstance<Name, Base, FPU_CPACRPeripheral>;
def FPU_CPACR : FPU_CPACRInstance<"FPU_CPACR", 0xe000ed88>;

#endif // _PERIPHERALS_FPU_CPACR_TD
