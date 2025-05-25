#ifndef _PERIPHERALS_WWDG_TD
#define _PERIPHERALS_WWDG_TD

include "base.td"

def WWDGPeripheral : PeripheralType<"WWDG", "WWDG"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"T", 0, 7, ReadWrite, "7-bit counter (MSB to LSB) These bits contain the value of the watchdog counter. It is decremented every (4096 x 2WDGTB[1:0]) PCLK cycles. A reset is produced when it is decremented from 0x40 to 0x3F (T6 becomes cleared).">,
      Field<"WDGA", 7, 1, ReadWrite, "Activation bit This bit is set by software and only cleared by hardware after a reset. When WDGA=1, the watchdog can generate a reset.">,
    ], "Control register">,
    Register<"CFR", 0x4, 32, [
      Field<"W", 0, 7, ReadWrite, "7-bit window value These bits contain the window value to be compared to the downcounter.">,
      Field<"EWI", 9, 1, ReadWrite, "Early wakeup interrupt When set, an interrupt occurs whenever the counter reaches the value 0x40. This interrupt is only cleared by hardware after a reset.">,
      Field<"WDGTB", 11, 2, ReadWrite, "Timer base The time base of the prescaler can be modified as follows:">,
    ], "Configuration register">,
    Register<"SR", 0x8, 32, [
      Field<"EWIF", 0, 1, ReadWrite, "Early wakeup interrupt flag This bit is set by hardware when the counter has reached the value 0x40. It must be cleared by software by writing 0. A write of 1 has no effect. This bit is also set if the interrupt is not enabled.">,
    ], "Status register">,
  ];
}

class WWDGInstance<string Name, int Base> : PeripheralInstance<Name, Base, WWDGPeripheral>;
def WWDG2 : WWDGInstance<"WWDG2", 0x40002c00>;
def WWDG1 : WWDGInstance<"WWDG1", 0x50003000>;

def WWDGGroup : PeripheralGroup<[ WWDG1, WWDG2 ]>;

#endif // _PERIPHERALS_WWDG_TD
