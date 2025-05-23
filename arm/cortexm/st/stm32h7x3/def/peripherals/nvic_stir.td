#ifndef _PERIPHERALS_NVIC_STIR_TD
#define _PERIPHERALS_NVIC_STIR_TD

include "base.td"

def NVIC_STIRPeripheral : PeripheralType<"NVIC_STIR", "Nested vectored interrupt controller"> {
  let accessWidth = 32;
  let registers = [
    Register<"STIR", 0x0, 32, [
      Field<"INTID", 0, 9, ReadWrite, "Software generated interrupt ID">,
    ], "Software trigger interrupt register">,
  ];
}

class NVIC_STIRInstance<string Name, int Base> : PeripheralInstance<Name, Base, NVIC_STIRPeripheral>;
def NVIC_STIR : NVIC_STIRInstance<"NVIC_STIR", 0xe000ef00>;

#endif // _PERIPHERALS_NVIC_STIR_TD
