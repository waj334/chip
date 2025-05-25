#ifndef _PERIPHERALS_RAMECC3_TD
#define _PERIPHERALS_RAMECC3_TD

include "base.td"

def RAMECC3Peripheral : PeripheralType<"RAMECC3", "ECC controller is associated to each RAM area"> {
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
  ];
}

class RAMECC3Instance<string Name, int Base> : PeripheralInstance<Name, Base, RAMECC3Peripheral>;
def RAMECC3 : RAMECC3Instance<"RAMECC3", 0x58027000>;

#endif // _PERIPHERALS_RAMECC3_TD
