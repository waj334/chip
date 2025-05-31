#ifndef _PERIPHERALS_DCA_TD
#define _PERIPHERALS_DCA_TD

include "base.td"

def DCAPeripheral : PeripheralType<"DCA", "Direct Cache Access"> {
  let accessWidth = 32;
  let registers = [
    Register<"DCADCRR", 0x0, 32, [
      Field<"TAG", 0, 21, ReadWrite, "Tag address. The number of significant bits of TAG depends on the cache size.">,
      Field<"VALID", 21, 1, ReadWrite, "Valid state of the cache line">,
    ], "Read the data from the Level 1 (L1) data cache from the location that is determined by the DCADCLR registers">,
    Register<"DCAICRR", 0x4, 32, [
      Field<"TAG", 0, 21, ReadWrite, "Tag address. The number of significant bits of TAG depends on the cache size.">,
      Field<"VALID", 21, 1, ReadWrite, "Valid state of the cache line">,
    ], "Read the data from the Level 1 (L1) instruction cache from the location that is determined by the DCAICLR registers">,
    Register<"DCADCLR", 0x10, 32, [
      Field<"RAMTYPE", 0, 1, ReadWrite, "RAM type">,
      Field<"OFFSET", 2, 3, ReadWrite, "Data offset">,
      Field<"SET", 5, 10, ReadWrite, "Set index. The number of significant bits of SET depends on the cache size.">,
      Field<"WAY", 30, 1, ReadWrite, "Cache way">,
    ], "Set the location to be read from the Level 1 (L1) data cache">,
    Register<"DCAICLR", 0x14, 32, [
      Field<"RAMTYPE", 0, 1, ReadWrite, "RAM type">,
      Field<"OFFSET", 2, 3, ReadWrite, "Data offset">,
      Field<"SET", 5, 10, ReadWrite, "Set index. The number of significant bits of SET depends on the cache size.">,
      Field<"WAY", 30, 1, ReadWrite, "Cache way">,
    ], "Set the location to be read from the Level 1 (L1) instruction cache">,
  ];
}

class DCAInstance<string Name, int Base> : PeripheralInstance<Name, Base, DCAPeripheral>;
def DCA : DCAInstance<"DCA", 0xe001e200>;

#endif // _PERIPHERALS_DCA_TD
