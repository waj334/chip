#ifndef _PERIPHERALS_ICB_TD
#define _PERIPHERALS_ICB_TD

include "base.td"

def ICBPeripheral : PeripheralType<"ICB", "Implementation Control Block"> {
  let accessWidth = 32;
  let registers = [
    Register<"ICTR", 0x0, 32, [
      Field<"INTLINESNUM", 0, 4, ReadWrite, "The total number of interrupt lines supported by an implementation, defined in groups of 32">,
    ], "Provides information about the interrupt controller">,
    Register<"ACTLR", 0x4, 32, [
      Field<"FPEXCODIS", 10, 1, ReadWrite, "Determines whether floating-point exception outputs are disabled">,
      Field<"DISNWAMODE", 11, 1, ReadWrite, "Determines whether no write allocate mode is disabled">,
      Field<"DISITMATBFLUSH", 12, 1, ReadWrite, "Determines whether Instrumentation Trace Macrocell (ITM) or Data Watchpoint and Trace (DWT) ATB flush is disabled">,
      Field<"EVENTBUSEN_S", 13, 1, ReadWrite, "Accessibility of EVENTBUSEN">,
      Field<"EVENTBUSEN", 14, 1, ReadWrite, "Activate EVENTBUS output">,
      Field<"DISCRITAXIRUR", 15, 1, ReadWrite, "Disable critical AXI Read-Under-Read">,
      Field<"DISCRITAXIRUW", 27, 1, ReadWrite, "Disable-Critical-AXI-Read-Under-Write">,
    ], "Allow software to control the processor features and functionality">,
    Register<"CPPWR", 0x8, 32, [
      Field<"SU0", 0, 1, ReadWrite, "This field indicates and allows modification of whether the state associated with coprocessor 0 is permitted to become unknown">,
      Field<"SUS0", 1, 1, ReadWrite, "This field indicates and allows modification of whether the SU0 field can be modified from Non-secure state">,
      Field<"SU1", 2, 1, ReadWrite, "This field indicates and allows modification of whether the state associated with coprocessor 1 is permitted to become unknown">,
      Field<"SUS1", 3, 1, ReadWrite, "This field indicates and allows modification of whether the SU1 field can be modified from Non-secure state">,
      Field<"SU2", 4, 1, ReadWrite, "This field indicates and allows modification of whether the state associated with coprocessor 2 is permitted to become unknown">,
      Field<"SUS2", 5, 1, ReadWrite, "This field indicates and allows modification of whether the SU2 field can be modified from Non-secure state">,
      Field<"SU3", 6, 1, ReadWrite, "This field indicates and allows modification of whether the state associated with coprocessor 3 is permitted to become unknown">,
      Field<"SUS3", 7, 1, ReadWrite, "This field indicates and allows modification of whether the SU3 field can be modified from Non-secure state">,
      Field<"SU4", 8, 1, ReadWrite, "This field indicates and allows modification of whether the state associated with coprocessor 4 is permitted to become unknown">,
      Field<"SUS4", 9, 1, ReadWrite, "This field indicates and allows modification of whether the SU4 field can be modified from Non-secure state">,
      Field<"SU5", 10, 1, ReadWrite, "This field indicates and allows modification of whether the state associated with coprocessor 5 is permitted to become unknown">,
      Field<"SUS5", 11, 1, ReadWrite, "This field indicates and allows modification of whether the SU5 field can be modified from Non-secure state">,
      Field<"SU6", 12, 1, ReadWrite, "This field indicates and allows modification of whether the state associated with coprocessor 6 is permitted to become unknown">,
      Field<"SUS6", 13, 1, ReadWrite, "This field indicates and allows modification of whether the SU6 field can be modified from Non-secure state">,
      Field<"SU7", 14, 1, ReadWrite, "This field indicates and allows modification of whether the state associated with coprocessor 7 is permitted to become unknown">,
      Field<"SUS7", 15, 1, ReadWrite, "This field indicates and allows modification of whether the SU7 field can be modified from Non-secure state">,
      Field<"SU10", 20, 1, ReadWrite, "This bit indicates and allows modification of whether the state associated with the floating-point and M-profile Vector Extension (MVE) functionality is permitted to become UNKNOWN.">,
      Field<"SUS10", 21, 1, ReadWrite, "This bit indicates and allows modification of whether the SU10 field can be modified from Non-secure state.">,
      Field<"SU11", 22, 1, ReadWrite, "If the value of this bit is not programmed to the same value as the SU10 field, then the value is unknown.">,
      Field<"SUS11", 23, 1, ReadWrite, "If the value of this bit is not programmed to the same value as the SUS10 field, then the value is unknown.">,
    ], "Specifies whether coprocessors are permitted to enter a non-retentive power state">,
  ];
}

class ICBInstance<string Name, int Base> : PeripheralInstance<Name, Base, ICBPeripheral>;
def ICB : ICBInstance<"ICB", 0xe000e004>;

#endif // _PERIPHERALS_ICB_TD
