#ifndef _PERIPHERALS_PCI_TD
#define _PERIPHERALS_PCI_TD

include "base.td"

def PCIPeripheral : PeripheralType<"PCI", "Implementation defined processor configuration information"> {
  let accessWidth = 32;
  let registers = [
    Register<"CFGINFOSEL", 0x0, 32, [
    ], "Selects the configuration information which can then be read back using CFGINFORD">,
    Register<"CFGINFORD", 0x4, 32, [
    ], "Displays the configuration information that the CFGINFOSEL register selects">,
  ];
}

class PCIInstance<string Name, int Base> : PeripheralInstance<Name, Base, PCIPeripheral>;
def PCI : PCIInstance<"PCI", 0xe001e700>;

#endif // _PERIPHERALS_PCI_TD
