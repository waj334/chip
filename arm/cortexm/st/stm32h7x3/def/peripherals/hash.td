#ifndef _PERIPHERALS_HASH_TD
#define _PERIPHERALS_HASH_TD

include "base.td"

def HASHPeripheral : PeripheralType<"HASH", "Hash processor"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"INIT", 2, 1, Write, "Initialize message digest calculation">,
      Field<"DMAE", 3, 1, ReadWrite, "DMA enable">,
      Field<"DATATYPE", 4, 2, ReadWrite, "Data type selection">,
      Field<"MODE", 6, 1, ReadWrite, "Mode selection">,
      Field<"ALGO0", 7, 1, ReadWrite, "Algorithm selection">,
      Field<"NBW", 8, 4, Read, "Number of words already pushed">,
      Field<"DINNE", 12, 1, Read, "DIN not empty">,
      Field<"MDMAT", 13, 1, ReadWrite, "Multiple DMA Transfers">,
      Field<"LKEY", 16, 1, ReadWrite, "Long key selection">,
      Field<"ALGO1", 18, 1, ReadWrite, "ALGO">,
    ], "control register">,
    Register<"DIN", 0x4, 32, [
      Field<"DATAIN", 0, 32, ReadWrite, "Data input">,
    ], "data input register">,
    Register<"STR", 0x8, 32, [
      Field<"DCAL", 8, 1, Write, "Digest calculation">,
      Field<"NBLW", 0, 5, ReadWrite, "Number of valid bits in the last word of the message">,
    ], "start register">,
    Register<"HR0", 0xc, 32, [
      Field<"H0", 0, 32, ReadWrite, "H0">,
    ], "digest registers">,
    Register<"HR1", 0x10, 32, [
      Field<"H1", 0, 32, ReadWrite, "H1">,
    ], "digest registers">,
    Register<"HR2", 0x14, 32, [
      Field<"H2", 0, 32, ReadWrite, "H2">,
    ], "digest registers">,
    Register<"HR3", 0x18, 32, [
      Field<"H3", 0, 32, ReadWrite, "H3">,
    ], "digest registers">,
    Register<"HR4", 0x1c, 32, [
      Field<"H4", 0, 32, ReadWrite, "H4">,
    ], "digest registers">,
    Register<"IMR", 0x20, 32, [
      Field<"DCIE", 1, 1, ReadWrite, "Digest calculation completion interrupt enable">,
      Field<"DINIE", 0, 1, ReadWrite, "Data input interrupt enable">,
    ], "interrupt enable register">,
    Register<"SR", 0x24, 32, [
      Field<"BUSY", 3, 1, Read, "Busy bit">,
      Field<"DMAS", 2, 1, Read, "DMA Status">,
      Field<"DCIS", 1, 1, ReadWrite, "Digest calculation completion interrupt status">,
      Field<"DINIS", 0, 1, ReadWrite, "Data input interrupt status">,
    ], "status register">,
    Register<"CSR0", 0xf8, 32, [
      Field<"CSR0", 0, 32, ReadWrite, "CSR0">,
    ], "context swap registers">,
    Register<"CSR1", 0xfc, 32, [
      Field<"CSR1", 0, 32, ReadWrite, "CSR1">,
    ], "context swap registers">,
    Register<"CSR2", 0x100, 32, [
      Field<"CSR2", 0, 32, ReadWrite, "CSR2">,
    ], "context swap registers">,
    Register<"CSR3", 0x104, 32, [
      Field<"CSR3", 0, 32, ReadWrite, "CSR3">,
    ], "context swap registers">,
    Register<"CSR4", 0x108, 32, [
      Field<"CSR4", 0, 32, ReadWrite, "CSR4">,
    ], "context swap registers">,
    Register<"CSR5", 0x10c, 32, [
      Field<"CSR5", 0, 32, ReadWrite, "CSR5">,
    ], "context swap registers">,
    Register<"CSR6", 0x110, 32, [
      Field<"CSR6", 0, 32, ReadWrite, "CSR6">,
    ], "context swap registers">,
    Register<"CSR7", 0x114, 32, [
      Field<"CSR7", 0, 32, ReadWrite, "CSR7">,
    ], "context swap registers">,
    Register<"CSR8", 0x118, 32, [
      Field<"CSR8", 0, 32, ReadWrite, "CSR8">,
    ], "context swap registers">,
    Register<"CSR9", 0x11c, 32, [
      Field<"CSR9", 0, 32, ReadWrite, "CSR9">,
    ], "context swap registers">,
    Register<"CSR10", 0x120, 32, [
      Field<"CSR10", 0, 32, ReadWrite, "CSR10">,
    ], "context swap registers">,
    Register<"CSR11", 0x124, 32, [
      Field<"CSR11", 0, 32, ReadWrite, "CSR11">,
    ], "context swap registers">,
    Register<"CSR12", 0x128, 32, [
      Field<"CSR12", 0, 32, ReadWrite, "CSR12">,
    ], "context swap registers">,
    Register<"CSR13", 0x12c, 32, [
      Field<"CSR13", 0, 32, ReadWrite, "CSR13">,
    ], "context swap registers">,
    Register<"CSR14", 0x130, 32, [
      Field<"CSR14", 0, 32, ReadWrite, "CSR14">,
    ], "context swap registers">,
    Register<"CSR15", 0x134, 32, [
      Field<"CSR15", 0, 32, ReadWrite, "CSR15">,
    ], "context swap registers">,
    Register<"CSR16", 0x138, 32, [
      Field<"CSR16", 0, 32, ReadWrite, "CSR16">,
    ], "context swap registers">,
    Register<"CSR17", 0x13c, 32, [
      Field<"CSR17", 0, 32, ReadWrite, "CSR17">,
    ], "context swap registers">,
    Register<"CSR18", 0x140, 32, [
      Field<"CSR18", 0, 32, ReadWrite, "CSR18">,
    ], "context swap registers">,
    Register<"CSR19", 0x144, 32, [
      Field<"CSR19", 0, 32, ReadWrite, "CSR19">,
    ], "context swap registers">,
    Register<"CSR20", 0x148, 32, [
      Field<"CSR20", 0, 32, ReadWrite, "CSR20">,
    ], "context swap registers">,
    Register<"CSR21", 0x14c, 32, [
      Field<"CSR21", 0, 32, ReadWrite, "CSR21">,
    ], "context swap registers">,
    Register<"CSR22", 0x150, 32, [
      Field<"CSR22", 0, 32, ReadWrite, "CSR22">,
    ], "context swap registers">,
    Register<"CSR23", 0x154, 32, [
      Field<"CSR23", 0, 32, ReadWrite, "CSR23">,
    ], "context swap registers">,
    Register<"CSR24", 0x158, 32, [
      Field<"CSR24", 0, 32, ReadWrite, "CSR24">,
    ], "context swap registers">,
    Register<"CSR25", 0x15c, 32, [
      Field<"CSR25", 0, 32, ReadWrite, "CSR25">,
    ], "context swap registers">,
    Register<"CSR26", 0x160, 32, [
      Field<"CSR26", 0, 32, ReadWrite, "CSR26">,
    ], "context swap registers">,
    Register<"CSR27", 0x164, 32, [
      Field<"CSR27", 0, 32, ReadWrite, "CSR27">,
    ], "context swap registers">,
    Register<"CSR28", 0x168, 32, [
      Field<"CSR28", 0, 32, ReadWrite, "CSR28">,
    ], "context swap registers">,
    Register<"CSR29", 0x16c, 32, [
      Field<"CSR29", 0, 32, ReadWrite, "CSR29">,
    ], "context swap registers">,
    Register<"CSR30", 0x170, 32, [
      Field<"CSR30", 0, 32, ReadWrite, "CSR30">,
    ], "context swap registers">,
    Register<"CSR31", 0x174, 32, [
      Field<"CSR31", 0, 32, ReadWrite, "CSR31">,
    ], "context swap registers">,
    Register<"CSR32", 0x178, 32, [
      Field<"CSR32", 0, 32, ReadWrite, "CSR32">,
    ], "context swap registers">,
    Register<"CSR33", 0x17c, 32, [
      Field<"CSR33", 0, 32, ReadWrite, "CSR33">,
    ], "context swap registers">,
    Register<"CSR34", 0x180, 32, [
      Field<"CSR34", 0, 32, ReadWrite, "CSR34">,
    ], "context swap registers">,
    Register<"CSR35", 0x184, 32, [
      Field<"CSR35", 0, 32, ReadWrite, "CSR35">,
    ], "context swap registers">,
    Register<"CSR36", 0x188, 32, [
      Field<"CSR36", 0, 32, ReadWrite, "CSR36">,
    ], "context swap registers">,
    Register<"CSR37", 0x18c, 32, [
      Field<"CSR37", 0, 32, ReadWrite, "CSR37">,
    ], "context swap registers">,
    Register<"CSR38", 0x190, 32, [
      Field<"CSR38", 0, 32, ReadWrite, "CSR38">,
    ], "context swap registers">,
    Register<"CSR39", 0x194, 32, [
      Field<"CSR39", 0, 32, ReadWrite, "CSR39">,
    ], "context swap registers">,
    Register<"CSR40", 0x198, 32, [
      Field<"CSR40", 0, 32, ReadWrite, "CSR40">,
    ], "context swap registers">,
    Register<"CSR41", 0x19c, 32, [
      Field<"CSR41", 0, 32, ReadWrite, "CSR41">,
    ], "context swap registers">,
    Register<"CSR42", 0x1a0, 32, [
      Field<"CSR42", 0, 32, ReadWrite, "CSR42">,
    ], "context swap registers">,
    Register<"CSR43", 0x1a4, 32, [
      Field<"CSR43", 0, 32, ReadWrite, "CSR43">,
    ], "context swap registers">,
    Register<"CSR44", 0x1a8, 32, [
      Field<"CSR44", 0, 32, ReadWrite, "CSR44">,
    ], "context swap registers">,
    Register<"CSR45", 0x1ac, 32, [
      Field<"CSR45", 0, 32, ReadWrite, "CSR45">,
    ], "context swap registers">,
    Register<"CSR46", 0x1b0, 32, [
      Field<"CSR46", 0, 32, ReadWrite, "CSR46">,
    ], "context swap registers">,
    Register<"CSR47", 0x1b4, 32, [
      Field<"CSR47", 0, 32, ReadWrite, "CSR47">,
    ], "context swap registers">,
    Register<"CSR48", 0x1b8, 32, [
      Field<"CSR48", 0, 32, ReadWrite, "CSR48">,
    ], "context swap registers">,
    Register<"CSR49", 0x1bc, 32, [
      Field<"CSR49", 0, 32, ReadWrite, "CSR49">,
    ], "context swap registers">,
    Register<"CSR50", 0x1c0, 32, [
      Field<"CSR50", 0, 32, ReadWrite, "CSR50">,
    ], "context swap registers">,
    Register<"CSR51", 0x1c4, 32, [
      Field<"CSR51", 0, 32, ReadWrite, "CSR51">,
    ], "context swap registers">,
    Register<"CSR52", 0x1c8, 32, [
      Field<"CSR52", 0, 32, ReadWrite, "CSR52">,
    ], "context swap registers">,
    Register<"CSR53", 0x1cc, 32, [
      Field<"CSR53", 0, 32, ReadWrite, "CSR53">,
    ], "context swap registers">,
    Register<"HASH_HR0", 0x310, 32, [
      Field<"H0", 0, 32, ReadWrite, "H0">,
    ], "HASH digest register">,
    Register<"HASH_HR1", 0x314, 32, [
      Field<"H1", 0, 32, ReadWrite, "H1">,
    ], "read-only">,
    Register<"HASH_HR2", 0x318, 32, [
      Field<"H2", 0, 32, ReadWrite, "H2">,
    ], "read-only">,
    Register<"HASH_HR3", 0x31c, 32, [
      Field<"H3", 0, 32, ReadWrite, "H3">,
    ], "read-only">,
    Register<"HASH_HR4", 0x320, 32, [
      Field<"H4", 0, 32, ReadWrite, "H4">,
    ], "read-only">,
    Register<"HASH_HR5", 0x324, 32, [
      Field<"H5", 0, 32, ReadWrite, "H5">,
    ], "read-only">,
    Register<"HASH_HR6", 0x328, 32, [
      Field<"H6", 0, 32, ReadWrite, "H6">,
    ], "read-only">,
    Register<"HASH_HR7", 0x32c, 32, [
      Field<"H7", 0, 32, ReadWrite, "H7">,
    ], "read-only">,
  ];
}

class HASHInstance<string Name, int Base> : PeripheralInstance<Name, Base, HASHPeripheral>;
def HASH : HASHInstance<"HASH", 0x48021400>;

#endif // _PERIPHERALS_HASH_TD
