#ifndef _PERIPHERALS_CRC_TD
#define _PERIPHERALS_CRC_TD

include "base.td"

def CRCPeripheral : PeripheralType<"CRC", "Cryptographic processor"> {
  let accessWidth = 32;
  let registers = [
    Register<"DR", 0x0, 32, [
      Field<"DR", 0, 32, ReadWrite, "Data Register">,
    ], "Data register">,
    Register<"IDR", 0x4, 32, [
      Field<"IDR", 0, 32, ReadWrite, "Independent Data register">,
    ], "Independent Data register">,
    Register<"CR", 0x8, 32, [
      Field<"RESET", 0, 1, Write, "RESET bit">,
      Field<"POLYSIZE", 3, 2, ReadWrite, "Polynomial size">,
      Field<"REV_IN", 5, 2, ReadWrite, "Reverse input data">,
      Field<"REV_OUT", 7, 1, ReadWrite, "Reverse output data">,
    ], "Control register">,
    Register<"INIT", 0xc, 32, [
      Field<"CRC_INIT", 0, 32, ReadWrite, "Programmable initial CRC value">,
    ], "Initial CRC value">,
    Register<"POL", 0x10, 32, [
      Field<"POL", 0, 32, ReadWrite, "Programmable polynomial">,
    ], "CRC polynomial">,
  ];
}

class CRCInstance<string Name, int Base> : PeripheralInstance<Name, Base, CRCPeripheral>;
def CRC : CRCInstance<"CRC", 0x58024c00>;

#endif // _PERIPHERALS_CRC_TD
