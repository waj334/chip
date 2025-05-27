#ifndef _PERIPHERALS_ERRBNK_TD
#define _PERIPHERALS_ERRBNK_TD

include "base.td"

def ERRBNKPeripheral : PeripheralType<"ERRBNK", "Implementation defined error banking"> {
  let accessWidth = 32;
  let registers = [
    Register<"IEBR0", 0x0, 32, [
      Field<"VALID", 0, 1, ReadWrite, "Indicates whether the entry is valid or not">,
      Field<"LOCKED", 1, 1, ReadWrite, "Indicates whether the location is locked or not">,
      Field<"LOCATION_LWO", 2, 3, ReadWrite, "Indicates the location in the L1 instruction cache RAM, line word offset">,
      Field<"LOCATION_INDEX", 5, 10, ReadWrite, "Indicates the location in the L1 instruction cache RAM, index">,
      Field<"LOCATION_WAY", 15, 1, ReadWrite, "Indicates the location in the L1 instruction cache RAM, way">,
      Field<"BANK", 16, 1, ReadWrite, "Indicates which RAM bank to use">,
      Field<"SWDEF", 30, 2, ReadWrite, "Error detection logic sets this field to 0b00 on a new allocation and on Cold reset">,
    ], "Record errors that occur during memory accesses to the L1 instruction cache">,
    Register<"IEBR1", 0x4, 32, [
      Field<"VALID", 0, 1, ReadWrite, "Indicates whether the entry is valid or not">,
      Field<"LOCKED", 1, 1, ReadWrite, "Indicates whether the location is locked or not">,
      Field<"LOCATION_LWO", 2, 3, ReadWrite, "Indicates the location in the L1 instruction cache RAM, line word offset">,
      Field<"LOCATION_INDEX", 5, 10, ReadWrite, "Indicates the location in the L1 instruction cache RAM, index">,
      Field<"LOCATION_WAY", 15, 1, ReadWrite, "Indicates the location in the L1 instruction cache RAM, way">,
      Field<"BANK", 16, 1, ReadWrite, "Indicates which RAM bank to use">,
      Field<"SWDEF", 30, 2, ReadWrite, "Error detection logic sets this field to 0b00 on a new allocation and on Cold reset">,
    ], "Record errors that occur during memory accesses to the L1 instruction cache">,
    Register<"DEBR0", 0x10, 32, [
      Field<"VALID", 0, 1, ReadWrite, "Indicates whether the entry is valid or not">,
      Field<"LOCKED", 1, 1, ReadWrite, "Indicates whether the location is locked or not">,
      Field<"LOCATION_LWO", 2, 3, ReadWrite, "Indicates the location in the L1 instruction cache RAM, line doubleword offset">,
      Field<"LOCATION_INDEX", 5, 9, ReadWrite, "Indicates the location in the L1 instruction cache RAM, index">,
      Field<"LOCATION_WAY", 14, 2, ReadWrite, "Indicates the location in the L1 instruction cache RAM, way">,
      Field<"BANK", 16, 1, ReadWrite, "Indicates which RAM bank to use">,
      Field<"TYPE", 17, 1, ReadWrite, "Indicates the error type">,
      Field<"SWDEF", 30, 2, ReadWrite, "Error detection logic sets this field to 0b00 on a new allocation and on Cold reset">,
    ], "Record errors that occur during memory accesses to the L1 data cache">,
    Register<"DEBR1", 0x14, 32, [
      Field<"VALID", 0, 1, ReadWrite, "Indicates whether the entry is valid or not">,
      Field<"LOCKED", 1, 1, ReadWrite, "Indicates whether the location is locked or not">,
      Field<"LOCATION_LWO", 2, 3, ReadWrite, "Indicates the location in the L1 instruction cache RAM, line doubleword offset">,
      Field<"LOCATION_INDEX", 5, 9, ReadWrite, "Indicates the location in the L1 instruction cache RAM, index">,
      Field<"LOCATION_WAY", 14, 2, ReadWrite, "Indicates the location in the L1 instruction cache RAM, way">,
      Field<"BANK", 16, 1, ReadWrite, "Indicates which RAM bank to use">,
      Field<"TYPE", 17, 1, ReadWrite, "Indicates the error type">,
      Field<"SWDEF", 30, 2, ReadWrite, "Error detection logic sets this field to 0b00 on a new allocation and on Cold reset">,
    ], "Record errors that occur during memory accesses to the L1 data cache">,
    Register<"TEBR0", 0x20, 32, [
      Field<"VALID", 0, 1, ReadWrite, "Indicates whether the entry is valid or not">,
      Field<"LOCKED", 1, 1, ReadWrite, "Indicates whether the location is locked or not">,
      Field<"LOCATION", 2, 22, ReadWrite, "Indicates the physical location in the data cache RAM">,
      Field<"BANK", 24, 3, ReadWrite, "Indicates which RAM bank to use">,
      Field<"TYPE", 27, 1, ReadWrite, "Indicates the error type">,
      Field<"POISON", 28, 1, ReadWrite, "Indicates whether a BusFault is generated or not">,
      Field<"SWDEF", 30, 2, ReadWrite, "Error detection logic sets this field to 0b00 on a new allocation and on Cold reset">,
    ], "Record the location of errors in the TCM">,
    Register<"TEBR1", 0x28, 32, [
      Field<"VALID", 0, 1, ReadWrite, "Indicates whether the entry is valid or not">,
      Field<"LOCKED", 1, 1, ReadWrite, "Indicates whether the location is locked or not">,
      Field<"LOCATION", 2, 22, ReadWrite, "Indicates the physical location in the data cache RAM">,
      Field<"BANK", 24, 3, ReadWrite, "Indicates which RAM bank to use">,
      Field<"TYPE", 27, 1, ReadWrite, "Indicates the error type">,
      Field<"POISON", 28, 1, ReadWrite, "Indicates whether a BusFault is generated or not">,
      Field<"SWDEF", 30, 2, ReadWrite, "Error detection logic sets this field to 0b00 on a new allocation and on Cold reset">,
    ], "Record the location of errors in the TCM">,
  ];
}

class ERRBNKInstance<string Name, int Base> : PeripheralInstance<Name, Base, ERRBNKPeripheral>;
def ERRBNK : ERRBNKInstance<"ERRBNK", 0xe001e100>;

#endif // _PERIPHERALS_ERRBNK_TD
