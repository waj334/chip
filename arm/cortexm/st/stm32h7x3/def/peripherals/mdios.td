#ifndef _PERIPHERALS_MDIOS_TD
#define _PERIPHERALS_MDIOS_TD

include "base.td"

def MDIOSPeripheral : PeripheralType<"MDIOS", "Management data input/output slave"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"EN", 0, 1, ReadWrite, "Peripheral enable">,
      Field<"WRIE", 1, 1, ReadWrite, "Register write interrupt enable">,
      Field<"RDIE", 2, 1, ReadWrite, "Register Read Interrupt Enable">,
      Field<"EIE", 3, 1, ReadWrite, "Error interrupt enable">,
      Field<"DPC", 7, 1, ReadWrite, "Disable Preamble Check">,
      Field<"PORT_ADDRESS", 8, 5, ReadWrite, "Slaves's address">,
    ], "MDIOS configuration register">,
    Register<"WRFR", 0x4, 32, [
      Field<"WRF", 0, 32, ReadWrite, "Write flags for MDIO registers 0 to 31">,
    ], "MDIOS write flag register">,
    Register<"CWRFR", 0x8, 32, [
      Field<"CWRF", 0, 32, ReadWrite, "Clear the write flag">,
    ], "MDIOS clear write flag register">,
    Register<"RDFR", 0xc, 32, [
      Field<"RDF", 0, 32, ReadWrite, "Read flags for MDIO registers 0 to 31">,
    ], "MDIOS read flag register">,
    Register<"CRDFR", 0x10, 32, [
      Field<"CRDF", 0, 32, ReadWrite, "Clear the read flag">,
    ], "MDIOS clear read flag register">,
    Register<"SR", 0x14, 32, [
      Field<"PERF", 0, 1, ReadWrite, "Preamble error flag">,
      Field<"SERF", 1, 1, ReadWrite, "Start error flag">,
      Field<"TERF", 2, 1, ReadWrite, "Turnaround error flag">,
    ], "MDIOS status register">,
    Register<"CLRFR", 0x18, 32, [
      Field<"CPERF", 0, 1, ReadWrite, "Clear the preamble error flag">,
      Field<"CSERF", 1, 1, ReadWrite, "Clear the start error flag">,
      Field<"CTERF", 2, 1, ReadWrite, "Clear the turnaround error flag">,
    ], "MDIOS clear flag register">,
    Register<"DINR0", 0x1c, 32, [
      Field<"DIN0", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 0">,
    Register<"DINR1", 0x20, 32, [
      Field<"DIN1", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 1">,
    Register<"DINR2", 0x24, 32, [
      Field<"DIN2", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 2">,
    Register<"DINR3", 0x28, 32, [
      Field<"DIN3", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 3">,
    Register<"DINR4", 0x2c, 32, [
      Field<"DIN4", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 4">,
    Register<"DINR5", 0x30, 32, [
      Field<"DIN5", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 5">,
    Register<"DINR6", 0x34, 32, [
      Field<"DIN6", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 6">,
    Register<"DINR7", 0x38, 32, [
      Field<"DIN7", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 7">,
    Register<"DINR8", 0x3c, 32, [
      Field<"DIN8", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 8">,
    Register<"DINR9", 0x40, 32, [
      Field<"DIN9", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 9">,
    Register<"DINR10", 0x44, 32, [
      Field<"DIN10", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 10">,
    Register<"DINR11", 0x48, 32, [
      Field<"DIN11", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 11">,
    Register<"DINR12", 0x4c, 32, [
      Field<"DIN12", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 12">,
    Register<"DINR13", 0x50, 32, [
      Field<"DIN13", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 13">,
    Register<"DINR14", 0x54, 32, [
      Field<"DIN14", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 14">,
    Register<"DINR15", 0x58, 32, [
      Field<"DIN15", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 15">,
    Register<"DINR16", 0x5c, 32, [
      Field<"DIN16", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 16">,
    Register<"DINR17", 0x60, 32, [
      Field<"DIN17", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 17">,
    Register<"DINR18", 0x64, 32, [
      Field<"DIN18", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 18">,
    Register<"DINR19", 0x68, 32, [
      Field<"DIN19", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 19">,
    Register<"DINR20", 0x6c, 32, [
      Field<"DIN20", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 20">,
    Register<"DINR21", 0x70, 32, [
      Field<"DIN21", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 21">,
    Register<"DINR22", 0x74, 32, [
      Field<"DIN22", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 22">,
    Register<"DINR23", 0x78, 32, [
      Field<"DIN23", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 23">,
    Register<"DINR24", 0x7c, 32, [
      Field<"DIN24", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 24">,
    Register<"DINR25", 0x80, 32, [
      Field<"DIN25", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 25">,
    Register<"DINR26", 0x84, 32, [
      Field<"DIN26", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 26">,
    Register<"DINR27", 0x88, 32, [
      Field<"DIN27", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 27">,
    Register<"DINR28", 0x8c, 32, [
      Field<"DIN28", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 28">,
    Register<"DINR29", 0x90, 32, [
      Field<"DIN29", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 29">,
    Register<"DINR30", 0x94, 32, [
      Field<"DIN30", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 30">,
    Register<"DINR31", 0x98, 32, [
      Field<"DIN31", 0, 16, ReadWrite, "Input data received from MDIO Master during write frames">,
    ], "MDIOS input data register 31">,
    Register<"DOUTR0", 0x9c, 32, [
      Field<"DOUT0", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 0">,
    Register<"DOUTR1", 0xa0, 32, [
      Field<"DOUT1", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 1">,
    Register<"DOUTR2", 0xa4, 32, [
      Field<"DOUT2", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 2">,
    Register<"DOUTR3", 0xa8, 32, [
      Field<"DOUT3", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 3">,
    Register<"DOUTR4", 0xac, 32, [
      Field<"DOUT4", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 4">,
    Register<"DOUTR5", 0xb0, 32, [
      Field<"DOUT5", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 5">,
    Register<"DOUTR6", 0xb4, 32, [
      Field<"DOUT6", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 6">,
    Register<"DOUTR7", 0xb8, 32, [
      Field<"DOUT7", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 7">,
    Register<"DOUTR8", 0xbc, 32, [
      Field<"DOUT8", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 8">,
    Register<"DOUTR9", 0xc0, 32, [
      Field<"DOUT9", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 9">,
    Register<"DOUTR10", 0xc4, 32, [
      Field<"DOUT10", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 10">,
    Register<"DOUTR11", 0xc8, 32, [
      Field<"DOUT11", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 11">,
    Register<"DOUTR12", 0xcc, 32, [
      Field<"DOUT12", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 12">,
    Register<"DOUTR13", 0xd0, 32, [
      Field<"DOUT13", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 13">,
    Register<"DOUTR14", 0xd4, 32, [
      Field<"DOUT14", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 14">,
    Register<"DOUTR15", 0xd8, 32, [
      Field<"DOUT15", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 15">,
    Register<"DOUTR16", 0xdc, 32, [
      Field<"DOUT16", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 16">,
    Register<"DOUTR17", 0xe0, 32, [
      Field<"DOUT17", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 17">,
    Register<"DOUTR18", 0xe4, 32, [
      Field<"DOUT18", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 18">,
    Register<"DOUTR19", 0xe8, 32, [
      Field<"DOUT19", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 19">,
    Register<"DOUTR20", 0xec, 32, [
      Field<"DOUT20", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 20">,
    Register<"DOUTR21", 0xf0, 32, [
      Field<"DOUT21", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 21">,
    Register<"DOUTR22", 0xf4, 32, [
      Field<"DOUT22", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 22">,
    Register<"DOUTR23", 0xf8, 32, [
      Field<"DOUT23", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 23">,
    Register<"DOUTR24", 0xfc, 32, [
      Field<"DOUT24", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 24">,
    Register<"DOUTR25", 0x100, 32, [
      Field<"DOUT25", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 25">,
    Register<"DOUTR26", 0x104, 32, [
      Field<"DOUT26", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 26">,
    Register<"DOUTR27", 0x108, 32, [
      Field<"DOUT27", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 27">,
    Register<"DOUTR28", 0x10c, 32, [
      Field<"DOUT28", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 28">,
    Register<"DOUTR29", 0x110, 32, [
      Field<"DOUT29", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 29">,
    Register<"DOUTR30", 0x114, 32, [
      Field<"DOUT30", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 30">,
    Register<"DOUTR31", 0x118, 32, [
      Field<"DOUT31", 0, 16, ReadWrite, "Output data sent to MDIO Master during read frames">,
    ], "MDIOS output data register 31">,
  ];
}

class MDIOSInstance<string Name, int Base> : PeripheralInstance<Name, Base, MDIOSPeripheral>;
def MDIOS : MDIOSInstance<"MDIOS", 0x40009400>;

#endif // _PERIPHERALS_MDIOS_TD
