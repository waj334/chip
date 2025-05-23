#ifndef _PERIPHERALS_SCB_ACTRL_TD
#define _PERIPHERALS_SCB_ACTRL_TD

include "base.td"

def SCB_ACTRLPeripheral : PeripheralType<"SCB_ACTRL", "System control block ACTLR"> {
  let accessWidth = 32;
  let registers = [
    Register<"ACTRL", 0x0, 32, [
      Field<"DISFOLD", 2, 1, ReadWrite, "DISFOLD">,
      Field<"FPEXCODIS", 10, 1, ReadWrite, "FPEXCODIS">,
      Field<"DISRAMODE", 11, 1, ReadWrite, "DISRAMODE">,
      Field<"DISITMATBFLUSH", 12, 1, ReadWrite, "DISITMATBFLUSH">,
    ], "Auxiliary control register">,
  ];
}

class SCB_ACTRLInstance<string Name, int Base> : PeripheralInstance<Name, Base, SCB_ACTRLPeripheral>;
def SCB_ACTRL : SCB_ACTRLInstance<"SCB_ACTRL", 0xe000e008>;

#endif // _PERIPHERALS_SCB_ACTRL_TD
