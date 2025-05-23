#ifndef _PERIPHERALS_LPTIM3_TD
#define _PERIPHERALS_LPTIM3_TD

include "base.td"

def LPTIM3Peripheral : PeripheralType<"LPTIM3", "Low power timer"> {
  let accessWidth = 32;
  let registers = [
    Register<"ISR", 0x0, 32, [
      Field<"DOWN", 6, 1, ReadWrite, "Counter direction change up to down">,
      Field<"UP", 5, 1, ReadWrite, "Counter direction change down to up">,
      Field<"ARROK", 4, 1, ReadWrite, "Autoreload register update OK">,
      Field<"CMPOK", 3, 1, ReadWrite, "Compare register update OK">,
      Field<"EXTTRIG", 2, 1, ReadWrite, "External trigger edge event">,
      Field<"ARRM", 1, 1, ReadWrite, "Autoreload match">,
      Field<"CMPM", 0, 1, ReadWrite, "Compare match">,
    ], "Interrupt and Status Register">,
    Register<"ICR", 0x4, 32, [
      Field<"DOWNCF", 6, 1, ReadWrite, "Direction change to down Clear Flag">,
      Field<"UPCF", 5, 1, ReadWrite, "Direction change to UP Clear Flag">,
      Field<"ARROKCF", 4, 1, ReadWrite, "Autoreload register update OK Clear Flag">,
      Field<"CMPOKCF", 3, 1, ReadWrite, "Compare register update OK Clear Flag">,
      Field<"EXTTRIGCF", 2, 1, ReadWrite, "External trigger valid edge Clear Flag">,
      Field<"ARRMCF", 1, 1, ReadWrite, "Autoreload match Clear Flag">,
      Field<"CMPMCF", 0, 1, ReadWrite, "compare match Clear Flag">,
    ], "Interrupt Clear Register">,
    Register<"IER", 0x8, 32, [
      Field<"DOWNIE", 6, 1, ReadWrite, "Direction change to down Interrupt Enable">,
      Field<"UPIE", 5, 1, ReadWrite, "Direction change to UP Interrupt Enable">,
      Field<"ARROKIE", 4, 1, ReadWrite, "Autoreload register update OK Interrupt Enable">,
      Field<"CMPOKIE", 3, 1, ReadWrite, "Compare register update OK Interrupt Enable">,
      Field<"EXTTRIGIE", 2, 1, ReadWrite, "External trigger valid edge Interrupt Enable">,
      Field<"ARRMIE", 1, 1, ReadWrite, "Autoreload match Interrupt Enable">,
      Field<"CMPMIE", 0, 1, ReadWrite, "Compare match Interrupt Enable">,
    ], "Interrupt Enable Register">,
    Register<"CFGR", 0xc, 32, [
      Field<"ENC", 24, 1, ReadWrite, "Encoder mode enable">,
      Field<"COUNTMODE", 23, 1, ReadWrite, "counter mode enabled">,
      Field<"PRELOAD", 22, 1, ReadWrite, "Registers update mode">,
      Field<"WAVPOL", 21, 1, ReadWrite, "Waveform shape polarity">,
      Field<"WAVE", 20, 1, ReadWrite, "Waveform shape">,
      Field<"TIMOUT", 19, 1, ReadWrite, "Timeout enable">,
      Field<"TRIGEN", 17, 2, ReadWrite, "Trigger enable and polarity">,
      Field<"TRIGSEL", 13, 3, ReadWrite, "Trigger selector">,
      Field<"PRESC", 9, 3, ReadWrite, "Clock prescaler">,
      Field<"TRGFLT", 6, 2, ReadWrite, "Configurable digital filter for trigger">,
      Field<"CKFLT", 3, 2, ReadWrite, "Configurable digital filter for external clock">,
      Field<"CKPOL", 1, 2, ReadWrite, "Clock Polarity">,
      Field<"CKSEL", 0, 1, ReadWrite, "Clock selector">,
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
    ], "LPTIM configuration register 2">,
  ];
}

class LPTIM3Instance<string Name, int Base> : PeripheralInstance<Name, Base, LPTIM3Peripheral>;
def LPTIM3 : LPTIM3Instance<"LPTIM3", 0x58002800>;
def LPTIM4 : LPTIM3Instance<"LPTIM4", 0x58002c00>;
def LPTIM5 : LPTIM3Instance<"LPTIM5", 0x58003000>;

#endif // _PERIPHERALS_LPTIM3_TD
