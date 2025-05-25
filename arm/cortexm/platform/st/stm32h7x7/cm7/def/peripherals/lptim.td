#ifndef _PERIPHERALS_LPTIM_TD
#define _PERIPHERALS_LPTIM_TD

include "base.td"

def LPTIMPeripheral : PeripheralType<"LPTIM", "Low power timer"> {
  let accessWidth = 32;
  let registers = [
    Register<"ISR", 0x0, 32, [
      Field<"CMPM", 0, 1, ReadWrite, "Compare match">,
      Field<"ARRM", 1, 1, ReadWrite, "Autoreload match">,
      Field<"EXTTRIG", 2, 1, ReadWrite, "External trigger edge event">,
      Field<"CMPOK", 3, 1, ReadWrite, "Compare register update OK">,
      Field<"ARROK", 4, 1, ReadWrite, "Autoreload register update OK">,
      Field<"UP", 5, 1, ReadWrite, "Counter direction change down to up">,
      Field<"DOWN", 6, 1, ReadWrite, "Counter direction change up to down">,
    ], "Interrupt and Status Register">,
    Register<"ICR", 0x4, 32, [
      Field<"CMPMCF", 0, 1, ReadWrite, "compare match Clear Flag">,
      Field<"ARRMCF", 1, 1, ReadWrite, "Autoreload match Clear Flag">,
      Field<"EXTTRIGCF", 2, 1, ReadWrite, "External trigger valid edge Clear Flag">,
      Field<"CMPOKCF", 3, 1, ReadWrite, "Compare register update OK Clear Flag">,
      Field<"ARROKCF", 4, 1, ReadWrite, "Autoreload register update OK Clear Flag">,
      Field<"UPCF", 5, 1, ReadWrite, "Direction change to UP Clear Flag">,
      Field<"DOWNCF", 6, 1, ReadWrite, "Direction change to down Clear Flag">,
    ], "Interrupt Clear Register">,
    Register<"IER", 0x8, 32, [
      Field<"CMPMIE", 0, 1, ReadWrite, "Compare match Interrupt Enable">,
      Field<"ARRMIE", 1, 1, ReadWrite, "Autoreload match Interrupt Enable">,
      Field<"EXTTRIGIE", 2, 1, ReadWrite, "External trigger valid edge Interrupt Enable">,
      Field<"CMPOKIE", 3, 1, ReadWrite, "Compare register update OK Interrupt Enable">,
      Field<"ARROKIE", 4, 1, ReadWrite, "Autoreload register update OK Interrupt Enable">,
      Field<"UPIE", 5, 1, ReadWrite, "Direction change to UP Interrupt Enable">,
      Field<"DOWNIE", 6, 1, ReadWrite, "Direction change to down Interrupt Enable">,
    ], "Interrupt Enable Register">,
    Register<"CFGR", 0xc, 32, [
      Field<"CKSEL", 0, 1, ReadWrite, "Clock selector">,
      Field<"CKPOL", 1, 2, ReadWrite, "Clock Polarity">,
      Field<"CKFLT", 3, 2, ReadWrite, "Configurable digital filter for external clock">,
      Field<"TRGFLT", 6, 2, ReadWrite, "Configurable digital filter for trigger">,
      Field<"PRESC", 9, 3, ReadWrite, "Clock prescaler">,
      Field<"TRIGSEL", 13, 3, ReadWrite, "Trigger selector">,
      Field<"TRIGEN", 17, 2, ReadWrite, "Trigger enable and polarity">,
      Field<"TIMOUT", 19, 1, ReadWrite, "Timeout enable">,
      Field<"WAVE", 20, 1, ReadWrite, "Waveform shape">,
      Field<"WAVPOL", 21, 1, ReadWrite, "Waveform shape polarity">,
      Field<"PRELOAD", 22, 1, ReadWrite, "Registers update mode">,
      Field<"COUNTMODE", 23, 1, ReadWrite, "counter mode enabled">,
      Field<"ENC", 24, 1, ReadWrite, "Encoder mode enable">,
    ], "Configuration Register">,
    Register<"CR", 0x10, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "LPTIM Enable">,
      Field<"SNGSTRT", 1, 1, ReadWrite, "LPTIM start in single mode">,
      Field<"CNTSTRT", 2, 1, ReadWrite, "Timer start in continuous mode">,
      Field<"COUNTRST", 3, 1, ReadWrite, "Counter reset">,
      Field<"RSTARE", 4, 1, ReadWrite, "Reset after read enable">,
    ], "Control Register">,
    Register<"CMP", 0x14, 32, [
      Field<"CMP", 0, 16, ReadWrite, "Compare value">,
    ], "Compare Register">,
    Register<"ARR", 0x18, 32, [
      Field<"ARR", 0, 16, ReadWrite, "Auto reload value">,
    ], "Autoreload Register">,
    Register<"CNT", 0x1c, 32, [
      Field<"CNT", 0, 16, ReadWrite, "Counter value">,
    ], "Counter Register">,
    Register<"CFGR2", 0x24, 32, [
      Field<"IN1SEL", 0, 2, ReadWrite, "LPTIM Input 1 selection">,
      Field<"IN2SEL", 4, 2, ReadWrite, "LPTIM Input 2 selection">,
    ], "LPTIM configuration register 2">,
  ];
}

class LPTIMInstance<string Name, int Base> : PeripheralInstance<Name, Base, LPTIMPeripheral>;
def LPTIM1 : LPTIMInstance<"LPTIM1", 0x40002400>;
def LPTIM2 : LPTIMInstance<"LPTIM2", 0x58002400>;
def LPTIM3 : LPTIMInstance<"LPTIM3", 0x58002800>;
def LPTIM4 : LPTIMInstance<"LPTIM4", 0x58002c00>;
def LPTIM5 : LPTIMInstance<"LPTIM5", 0x58003000>;

def LPTIMGroup : PeripheralGroup<[ LPTIM1, LPTIM2, LPTIM3, LPTIM4, LPTIM5 ]>;

#endif // _PERIPHERALS_LPTIM1_TD
