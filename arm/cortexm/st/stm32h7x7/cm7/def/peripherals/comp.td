#ifndef _PERIPHERALS_COMP_TD
#define _PERIPHERALS_COMP_TD

include "base.td"

def COMPPeripheral : PeripheralType<"COMP", "COMP"> {
  let accessWidth = 32;
  let registers = [
    Register<"SR", 0x0, 32, [
      Field<"C1VAL", 0, 1, ReadWrite, "COMP channel 1 output status bit">,
      Field<"C2VAL", 1, 1, ReadWrite, "COMP channel 2 output status bit">,
      Field<"C1IF", 16, 1, ReadWrite, "COMP channel 1 Interrupt Flag">,
      Field<"C2IF", 17, 1, ReadWrite, "COMP channel 2 Interrupt Flag">,
    ], "Comparator status register">,
    Register<"ICFR", 0x4, 32, [
      Field<"CC1IF", 16, 1, ReadWrite, "Clear COMP channel 1 Interrupt Flag">,
      Field<"CC2IF", 17, 1, ReadWrite, "Clear COMP channel 2 Interrupt Flag">,
    ], "Comparator interrupt clear flag register">,
    Register<"OR", 0x8, 32, [
      Field<"AFOP", 0, 11, ReadWrite, "Selection of source for alternate function of output ports">,
      Field<"OR", 11, 21, ReadWrite, "Option Register">,
    ], "Comparator option register">,
    Register<"CFGR1", 0xc, 32, [
      Field<"EN", 0, 1, ReadWrite, "COMP channel 1 enable bit">,
      Field<"BRGEN", 1, 1, ReadWrite, "Scaler bridge enable">,
      Field<"SCALEN", 2, 1, ReadWrite, "Voltage scaler enable bit">,
      Field<"POLARITY", 3, 1, ReadWrite, "COMP channel 1 polarity selection bit">,
      Field<"ITEN", 6, 1, ReadWrite, "COMP channel 1 interrupt enable">,
      Field<"HYST", 8, 2, ReadWrite, "COMP channel 1 hysteresis selection bits">,
      Field<"PWRMODE", 12, 2, ReadWrite, "Power Mode of the COMP channel 1">,
      Field<"INMSEL", 16, 3, ReadWrite, "COMP channel 1 inverting input selection field">,
      Field<"INPSEL", 20, 1, ReadWrite, "COMP channel 1 non-inverting input selection bit">,
      Field<"BLANKING", 24, 4, ReadWrite, "COMP channel 1 blanking source selection bits">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock bit">,
    ], "Comparator configuration register 1">,
    Register<"CFGR2", 0x10, 32, [
      Field<"EN", 0, 1, ReadWrite, "COMP channel 1 enable bit">,
      Field<"BRGEN", 1, 1, ReadWrite, "Scaler bridge enable">,
      Field<"SCALEN", 2, 1, ReadWrite, "Voltage scaler enable bit">,
      Field<"POLARITY", 3, 1, ReadWrite, "COMP channel 1 polarity selection bit">,
      Field<"WINMODE", 4, 1, ReadWrite, "Window comparator mode selection bit">,
      Field<"ITEN", 6, 1, ReadWrite, "COMP channel 1 interrupt enable">,
      Field<"HYST", 8, 2, ReadWrite, "COMP channel 1 hysteresis selection bits">,
      Field<"PWRMODE", 12, 2, ReadWrite, "Power Mode of the COMP channel 1">,
      Field<"INMSEL", 16, 3, ReadWrite, "COMP channel 1 inverting input selection field">,
      Field<"INPSEL", 20, 1, ReadWrite, "COMP channel 1 non-inverting input selection bit">,
      Field<"BLANKING", 24, 4, ReadWrite, "COMP channel 1 blanking source selection bits">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock bit">,
    ], "Comparator configuration register 2">,
  ];
}

class COMPInstance<string Name, int Base> : PeripheralInstance<Name, Base, COMPPeripheral>;
def COMP : COMPInstance<"COMP", 0x58003800>;

#endif // _PERIPHERALS_COMP_TD
