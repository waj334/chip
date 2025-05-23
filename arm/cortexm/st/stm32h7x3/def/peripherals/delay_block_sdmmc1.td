#ifndef _PERIPHERALS_DELAY_BLOCK_SDMMC1_TD
#define _PERIPHERALS_DELAY_BLOCK_SDMMC1_TD

include "base.td"

def DELAY_Block_SDMMC1Peripheral : PeripheralType<"DELAY_Block_SDMMC1", "DELAY_Block_SDMMC1"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"DEN", 0, 1, ReadWrite, "Delay block enable bit">,
      Field<"SEN", 1, 1, ReadWrite, "Sampler length enable bit">,
    ], "DLYB control register">,
    Register<"CFGR", 0x4, 32, [
      Field<"SEL", 0, 4, ReadWrite, "Select the phase for the Output clock">,
      Field<"UNIT", 8, 7, ReadWrite, "Delay Defines the delay of a Unit delay cell">,
      Field<"LNG", 16, 12, ReadWrite, "Delay line length value">,
      Field<"LNGF", 31, 1, ReadWrite, "Length valid flag">,
    ], "DLYB configuration register">,
  ];
}

class DELAY_Block_SDMMC1Instance<string Name, int Base> : PeripheralInstance<Name, Base, DELAY_Block_SDMMC1Peripheral>;
def DELAY_Block_SDMMC1 : DELAY_Block_SDMMC1Instance<"DELAY_Block_SDMMC1", 0x52008000>;
def DELAY_Block_QUADSPI : DELAY_Block_SDMMC1Instance<"DELAY_Block_QUADSPI", 0x52006000>;
def DELAY_Block_SDMMC2 : DELAY_Block_SDMMC1Instance<"DELAY_Block_SDMMC2", 0x48022800>;

#endif // _PERIPHERALS_DELAY_BLOCK_SDMMC1_TD
