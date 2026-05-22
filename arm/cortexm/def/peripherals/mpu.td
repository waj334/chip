#ifndef _PERIPHERALS_MPU_TD
#define _PERIPHERALS_MPU_TD

include "base.td"

def MPUPeripheral : PeripheralType<"MPU", "Memory Protection Unit"> {
  let accessWidth = 32;
  let registers = [
    Register<"TYPE", 0x0, 32, [
      Field<"SEPARATE", 0, 1, ReadOnly, "Indicates support for unified or separate instruction and data address regions">,
      Field<"DREGION", 8, 8, ReadOnly, "Number of supported MPU data regions">
    ], "Indicates whether the MPU is present, and how many regions it supports">,

    Register<"CTRL", 0x4, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Enables the MPU">,
      Field<"HFNMIENA", 1, 1, ReadWrite, "Enables MPU operation during HardFault and NMI handlers">,
      Field<"PRIVDEFENA", 2, 1, ReadWrite, "Enables privileged software access to the default memory map">
    ], "MPU control register">,

    Register<"RNR", 0x8, 32, [
      Field<"REGION", 0, 8, ReadWrite, "Selects the MPU region accessed by MPU_RBAR and MPU_RASR">
    ], "MPU region number register">,

    Register<"RBAR", 0xc, 32, [
      Field<"REGION", 0, 4, ReadWrite, "MPU region number field used when VALID is set">,
      Field<"VALID", 4, 1, ReadWrite, "Updates MPU_RNR with REGION when this bit is set">,
      Field<"ADDR", 5, 27, ReadWrite, "Region base address bits[31:5]">
    ], "MPU region base address register">,

    Register<"RASR", 0x10, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Enables the selected MPU region">,
      Field<"SIZE", 1, 5, ReadWrite, "Region size encoding; region size is 2^(SIZE + 1)">,
      Field<"SRD", 8, 8, ReadWrite, "Subregion disable bits">,
      Field<"B", 16, 1, ReadWrite, "Bufferable memory attribute">,
      Field<"C", 17, 1, ReadWrite, "Cacheable memory attribute">,
      Field<"S", 18, 1, ReadWrite, "Shareable memory attribute">,
      Field<"TEX", 19, 3, ReadWrite, "Type extension memory attribute">,
      Field<"AP", 24, 3, ReadWrite, "Access permission field">,
      Field<"XN", 28, 1, ReadWrite, "Execute-never attribute">
    ], "MPU region attribute and size register">,

    Register<"RBAR_A1", 0x14, 32, [
      Field<"REGION", 0, 4, ReadWrite, "MPU region number field used when VALID is set">,
      Field<"VALID", 4, 1, ReadWrite, "Updates MPU_RNR with REGION when this bit is set">,
      Field<"ADDR", 5, 27, ReadWrite, "Region base address bits[31:5]">
    ], "MPU region base address alias register 1">,

    Register<"RASR_A1", 0x18, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Enables the selected MPU region">,
      Field<"SIZE", 1, 5, ReadWrite, "Region size encoding; region size is 2^(SIZE + 1)">,
      Field<"SRD", 8, 8, ReadWrite, "Subregion disable bits">,
      Field<"B", 16, 1, ReadWrite, "Bufferable memory attribute">,
      Field<"C", 17, 1, ReadWrite, "Cacheable memory attribute">,
      Field<"S", 18, 1, ReadWrite, "Shareable memory attribute">,
      Field<"TEX", 19, 3, ReadWrite, "Type extension memory attribute">,
      Field<"AP", 24, 3, ReadWrite, "Access permission field">,
      Field<"XN", 28, 1, ReadWrite, "Execute-never attribute">
    ], "MPU region attribute and size alias register 1">,

    Register<"RBAR_A2", 0x1c, 32, [
      Field<"REGION", 0, 4, ReadWrite, "MPU region number field used when VALID is set">,
      Field<"VALID", 4, 1, ReadWrite, "Updates MPU_RNR with REGION when this bit is set">,
      Field<"ADDR", 5, 27, ReadWrite, "Region base address bits[31:5]">
    ], "MPU region base address alias register 2">,

    Register<"RASR_A2", 0x20, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Enables the selected MPU region">,
      Field<"SIZE", 1, 5, ReadWrite, "Region size encoding; region size is 2^(SIZE + 1)">,
      Field<"SRD", 8, 8, ReadWrite, "Subregion disable bits">,
      Field<"B", 16, 1, ReadWrite, "Bufferable memory attribute">,
      Field<"C", 17, 1, ReadWrite, "Cacheable memory attribute">,
      Field<"S", 18, 1, ReadWrite, "Shareable memory attribute">,
      Field<"TEX", 19, 3, ReadWrite, "Type extension memory attribute">,
      Field<"AP", 24, 3, ReadWrite, "Access permission field">,
      Field<"XN", 28, 1, ReadWrite, "Execute-never attribute">
    ], "MPU region attribute and size alias register 2">,

    Register<"RBAR_A3", 0x24, 32, [
      Field<"REGION", 0, 4, ReadWrite, "MPU region number field used when VALID is set">,
      Field<"VALID", 4, 1, ReadWrite, "Updates MPU_RNR with REGION when this bit is set">,
      Field<"ADDR", 5, 27, ReadWrite, "Region base address bits[31:5]">
    ], "MPU region base address alias register 3">,

    Register<"RASR_A3", 0x28, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Enables the selected MPU region">,
      Field<"SIZE", 1, 5, ReadWrite, "Region size encoding; region size is 2^(SIZE + 1)">,
      Field<"SRD", 8, 8, ReadWrite, "Subregion disable bits">,
      Field<"B", 16, 1, ReadWrite, "Bufferable memory attribute">,
      Field<"C", 17, 1, ReadWrite, "Cacheable memory attribute">,
      Field<"S", 18, 1, ReadWrite, "Shareable memory attribute">,
      Field<"TEX", 19, 3, ReadWrite, "Type extension memory attribute">,
      Field<"AP", 24, 3, ReadWrite, "Access permission field">,
      Field<"XN", 28, 1, ReadWrite, "Execute-never attribute">
    ], "MPU region attribute and size alias register 3">
  ];
}

class MPUInstance<string Name, int Base, list<string> Tags = []> : PeripheralInstance<Name, Base, MPUPeripheral, "", Tags>;

def MPU : MPUInstance<"MPU", 0xe000ed90, [ "armv7m"] >;

#endif // _PERIPHERALS_MPU_TD