#ifndef _PERIPHERALS_AC_TD
#define _PERIPHERALS_AC_TD

include "base.td"

def ACPeripheral : PeripheralType<"AC", "Access control"> {
  let accessWidth = 32;
  let registers = [
    Register<"ITCMCR", 0x0, 32, [
      Field<"EN", 0, 1, ReadWrite, "EN">,
      Field<"RMW", 1, 1, ReadWrite, "RMW">,
      Field<"RETEN", 2, 1, ReadWrite, "RETEN">,
      Field<"SZ", 3, 4, ReadWrite, "SZ">,
    ], "Instruction and Data Tightly-Coupled Memory Control Registers">,
    Register<"DTCMCR", 0x4, 32, [
      Field<"EN", 0, 1, ReadWrite, "EN">,
      Field<"RMW", 1, 1, ReadWrite, "RMW">,
      Field<"RETEN", 2, 1, ReadWrite, "RETEN">,
      Field<"SZ", 3, 4, ReadWrite, "SZ">,
    ], "Instruction and Data Tightly-Coupled Memory Control Registers">,
    Register<"AHBPCR", 0x8, 32, [
      Field<"EN", 0, 1, ReadWrite, "EN">,
      Field<"SZ", 1, 3, ReadWrite, "SZ">,
    ], "AHBP Control register">,
    Register<"CACR", 0xc, 32, [
      Field<"SIWT", 0, 1, ReadWrite, "SIWT">,
      Field<"ECCEN", 1, 1, ReadWrite, "ECCEN">,
      Field<"FORCEWT", 2, 1, ReadWrite, "FORCEWT">,
    ], "Auxiliary Cache Control register">,
    Register<"AHBSCR", 0x10, 32, [
      Field<"CTL", 0, 2, ReadWrite, "CTL">,
      Field<"TPRI", 2, 9, ReadWrite, "TPRI">,
      Field<"INITCOUNT", 11, 5, ReadWrite, "INITCOUNT">,
    ], "AHB Slave Control register">,
    Register<"ABFSR", 0x18, 32, [
      Field<"ITCM", 0, 1, ReadWrite, "ITCM">,
      Field<"DTCM", 1, 1, ReadWrite, "DTCM">,
      Field<"AHBP", 2, 1, ReadWrite, "AHBP">,
      Field<"AXIM", 3, 1, ReadWrite, "AXIM">,
      Field<"EPPB", 4, 1, ReadWrite, "EPPB">,
      Field<"AXIMTYPE", 8, 2, ReadWrite, "AXIMTYPE">,
    ], "Auxiliary Bus Fault Status register">,
  ];
}

class ACInstance<string Name, int Base> : PeripheralInstance<Name, Base, ACPeripheral>;
def AC : ACInstance<"AC", 0xe000ef90>;

#endif // _PERIPHERALS_AC_TD
