#ifndef _PERIPHERALS_PMC_TD
#define _PERIPHERALS_PMC_TD

include "base.td"

def PMCPeripheral : PeripheralType<"PMC", "Implementation defined power mode control"> {
  let accessWidth = 32;
  let registers = [
    Register<"CPDLPSTATE", 0x0, 32, [
      Field<"CLPSTATE", 0, 2, ReadWrite, "Type of low-power state for PDCORE">,
      Field<"ELPSTATE", 4, 2, ReadWrite, "Type of low-power state for PDEPU">,
      Field<"RLPSTATE", 8, 2, ReadWrite, "Power-on state for PDRAMS power domain">,
    ], "Specifies the desired low-power states for core (PDCORE), Extension Processing Unit (PDEPU), and RAM (PDRAMS) power domains">,
    Register<"DPDLPSTATE", 0x4, 32, [
      Field<"DLPSTATE", 0, 2, ReadWrite, "Type of low-power state for PDDEBUG">,
    ], "Specifies the desired low-power states for the debug (PDDEBUG) power domain">,
  ];
}

class PMCInstance<string Name, int Base> : PeripheralInstance<Name, Base, PMCPeripheral>;
def PMC : PMCInstance<"PMC", 0xe001e300>;

#endif // _PERIPHERALS_PMC_TD
