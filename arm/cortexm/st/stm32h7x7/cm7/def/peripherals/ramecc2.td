#ifndef _PERIPHERALS_RAMECC2_TD
#define _PERIPHERALS_RAMECC2_TD

include "base.td"

def RAMECC2Peripheral : PeripheralType<"RAMECC2", "ECC controller is associated to each RAM area"> {
  let accessWidth = 32;
  let registers = [
    Register<"IER", 0x0, 32, [
      Field<"GIE", 0, 1, ReadWrite, "Global interrupt enable">,
      Field<"GECCSEIE", 1, 1, ReadWrite, "Global ECC single error interrupt enable">,
      Field<"GECCDEIE", 2, 1, ReadWrite, "Global ECC double error interrupt enable">,
      Field<"GECCDEBWIE", 3, 1, ReadWrite, "Global ECC double error on byte write (BW) interrupt enable">,
    ], "RAMECC interrupt enable register">,
    Register<"M1CR", 0x20, 32, [
      Field<"ECCSEIE", 2, 1, ReadWrite, "ECC single error interrupt enable">,
      Field<"ECCDEIE", 3, 1, ReadWrite, "ECC double error interrupt enable">,
      Field<"ECCDEBWIE", 4, 1, ReadWrite, "ECC double error on byte write (BW) interrupt enable">,
      Field<"ECCELEN", 5, 1, ReadWrite, "ECC error latching enable">,
    ], "RAMECC monitor x configuration register">,
    Register<"M1SR", 0x24, 32, [
      Field<"SEDCF", 0, 1, ReadWrite, "ECC single error detected and corrected flag">,
      Field<"DEDF", 1, 1, ReadWrite, "ECC double error detected flag">,
      Field<"DEBWDF", 2, 1, ReadWrite, "ECC double error on byte write (BW) detected flag">,
    ], "RAMECC monitor x status register">,
    Register<"M1FAR", 0x28, 32, [
      Field<"FADD", 0, 32, ReadWrite, "ECC error failing address">,
    ], "RAMECC monitor x failing address register">,
    Register<"M1FDRL", 0x2c, 32, [
      Field<"FDATAL", 0, 32, ReadWrite, "Failing data low">,
    ], "RAMECC monitor x failing data low register">,
    Register<"M1FDRH", 0x30, 32, [
      Field<"FDATAH", 0, 32, ReadWrite, "Failing data high (64-bit memory)">,
    ], "RAMECC monitor x failing data high register">,
    Register<"M1FECR", 0x34, 32, [
      Field<"FEC", 0, 32, ReadWrite, "Failing error code">,
    ], "RAMECC monitor x failing ECC error code register">,
    Register<"M2CR", 0x40, 32, [
      Field<"ECCSEIE", 2, 1, ReadWrite, "ECC single error interrupt enable">,
      Field<"ECCDEIE", 3, 1, ReadWrite, "ECC double error interrupt enable">,
      Field<"ECCDEBWIE", 4, 1, ReadWrite, "ECC double error on byte write (BW) interrupt enable">,
      Field<"ECCELEN", 5, 1, ReadWrite, "ECC error latching enable">,
    ], "RAMECC monitor x configuration register">,
    Register<"M2SR", 0x44, 32, [
      Field<"SEDCF", 0, 1, ReadWrite, "ECC single error detected and corrected flag">,
      Field<"DEDF", 1, 1, ReadWrite, "ECC double error detected flag">,
      Field<"DEBWDF", 2, 1, ReadWrite, "ECC double error on byte write (BW) detected flag">,
    ], "RAMECC monitor x status register">,
    Register<"M2FAR", 0x48, 32, [
      Field<"FADD", 0, 32, ReadWrite, "ECC error failing address">,
    ], "RAMECC monitor x failing address register">,
    Register<"M2FDRL", 0x4c, 32, [
      Field<"FDATAL", 0, 32, ReadWrite, "Failing data low">,
    ], "RAMECC monitor x failing data low register">,
    Register<"M2FDRH", 0x50, 32, [
      Field<"FDATAH", 0, 32, ReadWrite, "Failing data high (64-bit memory)">,
    ], "RAMECC monitor x failing data high register">,
    Register<"M2FECR", 0x58, 32, [
      Field<"FEC", 0, 32, ReadWrite, "Failing error code">,
    ], "RAMECC monitor x failing ECC error code register">,
    Register<"M3CR", 0x60, 32, [
      Field<"ECCSEIE", 2, 1, ReadWrite, "ECC single error interrupt enable">,
      Field<"ECCDEIE", 3, 1, ReadWrite, "ECC double error interrupt enable">,
      Field<"ECCDEBWIE", 4, 1, ReadWrite, "ECC double error on byte write (BW) interrupt enable">,
      Field<"ECCELEN", 5, 1, ReadWrite, "ECC error latching enable">,
    ], "RAMECC monitor x configuration register">,
    Register<"M3SR", 0x64, 32, [
      Field<"SEDCF", 0, 1, ReadWrite, "ECC single error detected and corrected flag">,
      Field<"DEDF", 1, 1, ReadWrite, "ECC double error detected flag">,
      Field<"DEBWDF", 2, 1, ReadWrite, "ECC double error on byte write (BW) detected flag">,
    ], "RAMECC monitor x status register">,
    Register<"M3FAR", 0x68, 32, [
      Field<"FADD", 0, 32, ReadWrite, "ECC error failing address">,
    ], "RAMECC monitor x failing address register">,
    Register<"M3FDRL", 0x6c, 32, [
      Field<"FDATAL", 0, 32, ReadWrite, "Failing data low">,
    ], "RAMECC monitor x failing data low register">,
    Register<"M3FDRH", 0x70, 32, [
      Field<"FDATAH", 0, 32, ReadWrite, "Failing data high (64-bit memory)">,
    ], "RAMECC monitor x failing data high register">,
    Register<"M3FECR", 0x7c, 32, [
      Field<"FEC", 0, 32, ReadWrite, "Failing error code">,
    ], "RAMECC monitor x failing ECC error code register">,
    Register<"M4CR", 0x80, 32, [
      Field<"ECCSEIE", 2, 1, ReadWrite, "ECC single error interrupt enable">,
      Field<"ECCDEIE", 3, 1, ReadWrite, "ECC double error interrupt enable">,
      Field<"ECCDEBWIE", 4, 1, ReadWrite, "ECC double error on byte write (BW) interrupt enable">,
      Field<"ECCELEN", 5, 1, ReadWrite, "ECC error latching enable">,
    ], "RAMECC monitor x configuration register">,
    Register<"M4SR", 0x84, 32, [
      Field<"SEDCF", 0, 1, ReadWrite, "ECC single error detected and corrected flag">,
      Field<"DEDF", 1, 1, ReadWrite, "ECC double error detected flag">,
      Field<"DEBWDF", 2, 1, ReadWrite, "ECC double error on byte write (BW) detected flag">,
    ], "RAMECC monitor x status register">,
    Register<"M4FAR", 0x88, 32, [
      Field<"FADD", 0, 32, ReadWrite, "ECC error failing address">,
    ], "RAMECC monitor x failing address register">,
    Register<"M4FDRL", 0x8c, 32, [
      Field<"FDATAL", 0, 32, ReadWrite, "Failing data low">,
    ], "RAMECC monitor x failing data low register">,
    Register<"M4FDRH", 0x90, 32, [
      Field<"FDATAH", 0, 32, ReadWrite, "Failing data high (64-bit memory)">,
    ], "RAMECC monitor x failing data high register">,
    Register<"M4FECR", 0x90, 32, [
      Field<"FEC", 0, 32, ReadWrite, "Failing error code">,
    ], "RAMECC monitor x failing ECC error code register">,
    Register<"M5CR", 0xa0, 32, [
      Field<"ECCSEIE", 2, 1, ReadWrite, "ECC single error interrupt enable">,
      Field<"ECCDEIE", 3, 1, ReadWrite, "ECC double error interrupt enable">,
      Field<"ECCDEBWIE", 4, 1, ReadWrite, "ECC double error on byte write (BW) interrupt enable">,
      Field<"ECCELEN", 5, 1, ReadWrite, "ECC error latching enable">,
    ], "RAMECC monitor x configuration register">,
    Register<"M5SR", 0xa4, 32, [
      Field<"SEDCF", 0, 1, ReadWrite, "ECC single error detected and corrected flag">,
      Field<"DEDF", 1, 1, ReadWrite, "ECC double error detected flag">,
      Field<"DEBWDF", 2, 1, ReadWrite, "ECC double error on byte write (BW) detected flag">,
    ], "RAMECC monitor x status register">,
    Register<"M5FAR", 0xa8, 32, [
      Field<"FADD", 0, 32, ReadWrite, "ECC error failing address">,
    ], "RAMECC monitor x failing address register">,
    Register<"M5FDRL", 0xac, 32, [
      Field<"FDATAL", 0, 32, ReadWrite, "Failing data low">,
    ], "RAMECC monitor x failing data low register">,
    Register<"M5FDRH", 0xb0, 32, [
      Field<"FEC", 0, 32, ReadWrite, "Failing error code">,
    ], "RAMECC monitor x failing data high register">,
    Register<"M5FECR", 0xb4, 32, [
      Field<"FEC", 0, 32, ReadWrite, "Failing error code">,
    ], "RAMECC monitor x failing ECC error code register">,
  ];
}

class RAMECC2Instance<string Name, int Base> : PeripheralInstance<Name, Base, RAMECC2Peripheral>;
def RAMECC2 : RAMECC2Instance<"RAMECC2", 0x48023000>;

#endif // _PERIPHERALS_RAMECC2_TD
