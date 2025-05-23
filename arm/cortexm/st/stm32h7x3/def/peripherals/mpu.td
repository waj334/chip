#ifndef _PERIPHERALS_MPU_TD
#define _PERIPHERALS_MPU_TD

include "base.td"

def MPUPeripheral : PeripheralType<"MPU", "Memory protection unit"> {
  let accessWidth = 32;
  let registers = [
    Register<"MPU_TYPER", 0x0, 32, [
      Field<"SEPARATE", 0, 1, ReadWrite, "Separate flag">,
      Field<"DREGION", 8, 8, ReadWrite, "Number of MPU data regions">,
      Field<"IREGION", 16, 8, ReadWrite, "Number of MPU instruction regions">,
    ], "MPU type register">,
    Register<"MPU_CTRL", 0x4, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Enables the MPU">,
      Field<"HFNMIENA", 1, 1, ReadWrite, "Enables the operation of MPU during hard fault">,
      Field<"PRIVDEFENA", 2, 1, ReadWrite, "Enable priviliged software access to default memory map">,
    ], "MPU control register">,
    Register<"MPU_RNR", 0x8, 32, [
      Field<"REGION", 0, 8, ReadWrite, "MPU region">,
    ], "MPU region number register">,
    Register<"MPU_RBAR", 0xc, 32, [
      Field<"REGION", 0, 4, ReadWrite, "MPU region field">,
      Field<"VALID", 4, 1, ReadWrite, "MPU region number valid">,
      Field<"ADDR", 5, 27, ReadWrite, "Region base address field">,
    ], "MPU region base address register">,
    Register<"MPU_RASR", 0x10, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Region enable bit.">,
      Field<"SIZE", 1, 5, ReadWrite, "Size of the MPU protection region">,
      Field<"SRD", 8, 8, ReadWrite, "Subregion disable bits">,
      Field<"B", 16, 1, ReadWrite, "memory attribute">,
      Field<"C", 17, 1, ReadWrite, "memory attribute">,
      Field<"S", 18, 1, ReadWrite, "Shareable memory attribute">,
      Field<"TEX", 19, 3, ReadWrite, "memory attribute">,
      Field<"AP", 24, 3, ReadWrite, "Access permission">,
      Field<"XN", 28, 1, ReadWrite, "Instruction access disable bit">,
    ], "MPU region attribute and size register">,
  ];
}

class MPUInstance<string Name, int Base> : PeripheralInstance<Name, Base, MPUPeripheral>;
def MPU : MPUInstance<"MPU", 0xe000ed90>;

#endif // _PERIPHERALS_MPU_TD
