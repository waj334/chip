#ifndef _PERIPHERALS_PF_TD
#define _PERIPHERALS_PF_TD

include "base.td"

def PFPeripheral : PeripheralType<"PF", "Processor features"> {
  let accessWidth = 32;
  let registers = [
    Register<"CLIDR", 0x0, 32, [
      Field<"CL1", 0, 3, ReadWrite, "CL1">,
      Field<"CL2", 3, 3, ReadWrite, "CL2">,
      Field<"CL3", 6, 3, ReadWrite, "CL3">,
      Field<"CL4", 9, 3, ReadWrite, "CL4">,
      Field<"CL5", 12, 3, ReadWrite, "CL5">,
      Field<"CL6", 15, 3, ReadWrite, "CL6">,
      Field<"CL7", 18, 3, ReadWrite, "CL7">,
      Field<"LoUIS", 21, 3, ReadWrite, "LoUIS">,
      Field<"LoC", 24, 3, ReadWrite, "LoC">,
      Field<"LoU", 27, 3, ReadWrite, "LoU">,
    ], "Cache Level ID register">,
    Register<"CTR", 0x4, 32, [
      Field<"_IminLine", 0, 4, ReadWrite, "IminLine">,
      Field<"DMinLine", 16, 4, ReadWrite, "DMinLine">,
      Field<"ERG", 20, 4, ReadWrite, "ERG">,
      Field<"CWG", 24, 4, ReadWrite, "CWG">,
      Field<"Format", 29, 3, ReadWrite, "Format">,
    ], "Cache Type register">,
    Register<"CCSIDR", 0x8, 32, [
      Field<"LineSize", 0, 3, ReadWrite, "LineSize">,
      Field<"Associativity", 3, 10, ReadWrite, "Associativity">,
      Field<"NumSets", 13, 15, ReadWrite, "NumSets">,
      Field<"WA", 28, 1, ReadWrite, "WA">,
      Field<"RA", 29, 1, ReadWrite, "RA">,
      Field<"WB", 30, 1, ReadWrite, "WB">,
      Field<"WT", 31, 1, ReadWrite, "WT">,
    ], "Cache Size ID register">,
  ];
}

class PFInstance<string Name, int Base> : PeripheralInstance<Name, Base, PFPeripheral>;
def PF : PFInstance<"PF", 0xe000ed78>;

#endif // _PERIPHERALS_PF_TD
