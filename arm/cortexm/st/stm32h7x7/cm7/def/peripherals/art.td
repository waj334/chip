#ifndef _PERIPHERALS_ART_TD
#define _PERIPHERALS_ART_TD

include "base.td"

def ARTPeripheral : PeripheralType<"ART", "accelerator - control register (ART_CTR)"> {
  let accessWidth = 32;
  let registers = [
    Register<"CTR", 0x0, 32, [
      Field<"EN", 0, 1, ReadWrite, "Cache enable">,
      Field<"PCACHEADDR", 8, 12, ReadWrite, "Cacheable page index">,
    ], "control register">,
  ];
}

class ARTInstance<string Name, int Base> : PeripheralInstance<Name, Base, ARTPeripheral>;
def ART : ARTInstance<"ART", 0x40024400>;

#endif // _PERIPHERALS_ART_TD
