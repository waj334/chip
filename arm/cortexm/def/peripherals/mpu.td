#ifndef _PERIPHERALS_MPU_TD
#define _PERIPHERALS_MPU_TD

include "base.td"

def MPUPeripheral : PeripheralType<"MPU", "Memory Protection Unit"> {
  let accessWidth = 32;
  let registers = [
    Register<"TYPE", 0x0, 32, [
      Field<"SEPARATE", 0, 1, ReadWrite, "Indicates support for unified or separate instructions and data address regions">,
      Field<"DREGION", 8, 8, ReadWrite, "Number of regions supported by the MPU">,
    ], "Indicates whether the MPU is present, and if so, how many regions it supports">,
    Register<"CTRL", 0x4, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Enables the MPU">,
      Field<"HFNMIENA", 1, 1, ReadWrite, "Enables the operation of MPU during HardFault and NMI handlers">,
      Field<"PRIVDEFENA", 2, 1, ReadWrite, "Enables privileged software access to the default memory map">,
    ], "Enables the MPU">,
    Register<"RNR", 0x8, 32, [
      Field<"REGION", 0, 8, ReadWrite, "Indicates the memory region accessed by MPU_RBAR and PMU_RLAR">,
    ], "Selects the region currently accessed by MPU_RBAR and MPU_RLAR">,
    Register<"RBAR", 0xc, 32, [
      Field<"XN", 0, 1, ReadWrite, "Defines whether code can be executed from this region">,
      Field<"AP", 1, 2, ReadWrite, "Access permissions">,
      Field<"SH", 3, 2, ReadWrite, "Defines the shareability domain of this region for Normal memory">,
      Field<"BASE", 5, 27, ReadWrite, "Contains bits[31:5] of the lower inclusive limit of the selected MPU memory region. This value is zero extended to provide the base address to be checked against.">,
    ], "Defines the base address of the MPU region selected by the MPU_RNR">,
    Register<"RLAR", 0x10, 32, [
      Field<"EN", 0, 1, ReadWrite, "Region enable">,
      Field<"ATTRINDX", 1, 3, ReadWrite, "Associates a set of attributes in the MPU_MAIR0 and MPU_MAIR1 fields">,
      Field<"PXN", 4, 1, ReadWrite, "Defines whether code can be executed from this privileged region">,
      Field<"LIMIT", 5, 27, ReadWrite, "Contains bits[31:5] of the upper inclusive limit of the selected MPU memory region. This value is postfixed with 0x1F to provide the limit address to be checked against.">,
    ], "Defines the limit address of the MPU region selected by the MPU_RNR">,
    Register<"RBAR_A1", 0x14, 32, [
    ], "Provides indirect read and write access to the MPU base address register">,
    Register<"RBAR_A2", 0x18, 32, [
    ], "Provides indirect read and write access to the MPU base address register">,
    Register<"RBAR_A3", 0x1c, 32, [
    ], "Provides indirect read and write access to the MPU base address register">,
    Register<"RLAR_A1", 0x20, 32, [
    ], "Provides indirect read and write access to the MPU limit address register">,
    Register<"RLAR_A2", 0x24, 32, [
    ], "Provides indirect read and write access to the MPU limit address register">,
    Register<"RLAR_A3", 0x28, 32, [
    ], "Provides indirect read and write access to the MPU limit address register">,
    Register<"MAIR0", 0x30, 32, [
      Field<"ATTR0", 0, 8, ReadWrite, "Memory attribute encoding for MPU regions with an AttrIndex of 0">,
      Field<"ATTR1", 8, 8, ReadWrite, "Memory attribute encoding for MPU regions with an AttrIndex of 1">,
      Field<"ATTR2", 16, 8, ReadWrite, "Memory attribute encoding for MPU regions with an AttrIndex of 2">,
      Field<"ATTR3", 24, 8, ReadWrite, "Memory attribute encoding for MPU regions with an AttrIndex of 3">,
    ], "Provides the memory attribute encodings corresponding to the AttrIndex values">,
    Register<"MAIR1", 0x34, 32, [
      Field<"ATTR4", 0, 8, ReadWrite, "Memory attribute encoding for MPU regions with an AttrIndex of 4">,
      Field<"ATTR5", 8, 8, ReadWrite, "Memory attribute encoding for MPU regions with an AttrIndex of 5">,
      Field<"ATTR6", 16, 8, ReadWrite, "Memory attribute encoding for MPU regions with an AttrIndex of 6">,
      Field<"ATTR7", 24, 8, ReadWrite, "Memory attribute encoding for MPU regions with an AttrIndex of 7">,
    ], "Provides the memory attribute encodings corresponding to the AttrIndex values">,
  ];
}

class MPUInstance<string Name, int Base> : PeripheralInstance<Name, Base, MPUPeripheral>;
def MPU : MPUInstance<"MPU", 0xe000ed90>;

#endif // _PERIPHERALS_MPU_TD
