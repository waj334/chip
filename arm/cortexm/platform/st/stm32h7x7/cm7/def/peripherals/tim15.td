#ifndef _PERIPHERALS_TIM15_TD
#define _PERIPHERALS_TIM15_TD

include "base.td"

def TIM15Peripheral : PeripheralType<"TIM15", "General purpose timers"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR1", 0x0, 32, [
      Field<"CEN", 0, 1, ReadWrite, "Counter enable">,
      Field<"UDIS", 1, 1, ReadWrite, "Update disable">,
      Field<"URS", 2, 1, ReadWrite, "Update request source">,
      Field<"OPM", 3, 1, ReadWrite, "One-pulse mode">,
      Field<"ARPE", 7, 1, ReadWrite, "Auto-reload preload enable">,
      Field<"CKD", 8, 2, ReadWrite, "Clock division">,
      Field<"UIFREMAP", 11, 1, ReadWrite, "UIF status bit remapping">,
    ], "control register 1">,
    Register<"CR2", 0x4, 32, [
      Field<"CCPC", 0, 1, ReadWrite, "Capture/compare preloaded control">,
      Field<"CCUS", 2, 1, ReadWrite, "Capture/compare control update selection">,
      Field<"CCDS", 3, 1, ReadWrite, "Capture/compare DMA selection">,
      Field<"MMS", 4, 3, ReadWrite, "Master mode selection">,
      Field<"TI1S", 7, 1, ReadWrite, "TI1 selection">,
      Field<"OIS1", 8, 1, ReadWrite, "Output Idle state 1">,
      Field<"OIS1N", 9, 1, ReadWrite, "Output Idle state 1">,
      Field<"OIS2", 10, 1, ReadWrite, "Output Idle state 2">,
    ], "control register 2">,
    Register<"SMCR", 0x8, 32, [
      Field<"SMS", 0, 3, ReadWrite, "Slave mode selection">,
      Field<"TS_2_0", 4, 3, ReadWrite, "Trigger selection">,
      Field<"MSM", 7, 1, ReadWrite, "Master/Slave mode">,
      Field<"SMS_3", 16, 1, ReadWrite, "Slave mode selection bit 3">,
      Field<"TS_4_3", 20, 2, ReadWrite, "Trigger selection - bit 4:3">,
    ], "slave mode control register">,
    Register<"DIER", 0xc, 32, [
      Field<"UIE", 0, 1, ReadWrite, "Update interrupt enable">,
      Field<"CC1IE", 1, 1, ReadWrite, "Capture/Compare 1 interrupt enable">,
      Field<"CC2IE", 2, 1, ReadWrite, "Capture/Compare 2 interrupt enable">,
      Field<"COMIE", 5, 1, ReadWrite, "COM interrupt enable">,
      Field<"TIE", 6, 1, ReadWrite, "Trigger interrupt enable">,
      Field<"BIE", 7, 1, ReadWrite, "Break interrupt enable">,
      Field<"UDE", 8, 1, ReadWrite, "Update DMA request enable">,
      Field<"CC1DE", 9, 1, ReadWrite, "Capture/Compare 1 DMA request enable">,
      Field<"CC2DE", 10, 1, ReadWrite, "Capture/Compare 2 DMA request enable">,
      Field<"COMDE", 13, 1, ReadWrite, "COM DMA request enable">,
      Field<"TDE", 14, 1, ReadWrite, "Trigger DMA request enable">,
    ], "DMA/Interrupt enable register">,
    Register<"SR", 0x10, 32, [
      Field<"UIF", 0, 1, ReadWrite, "Update interrupt flag">,
      Field<"CC1IF", 1, 1, ReadWrite, "Capture/compare 1 interrupt flag">,
      Field<"CC2IF", 2, 1, ReadWrite, "Capture/Compare 2 interrupt flag">,
      Field<"COMIF", 5, 1, ReadWrite, "COM interrupt flag">,
      Field<"TIF", 6, 1, ReadWrite, "Trigger interrupt flag">,
      Field<"BIF", 7, 1, ReadWrite, "Break interrupt flag">,
      Field<"CC1OF", 9, 1, ReadWrite, "Capture/Compare 1 overcapture flag">,
      Field<"CC2OF", 10, 1, ReadWrite, "Capture/compare 2 overcapture flag">,
    ], "status register">,
    Register<"EGR", 0x14, 32, [
      Field<"UG", 0, 1, ReadWrite, "Update generation">,
      Field<"CC1G", 1, 1, ReadWrite, "Capture/compare 1 generation">,
      Field<"CC2G", 2, 1, ReadWrite, "Capture/compare 2 generation">,
      Field<"COMG", 5, 1, ReadWrite, "Capture/Compare control update generation">,
      Field<"TG", 6, 1, ReadWrite, "Trigger generation">,
      Field<"BG", 7, 1, ReadWrite, "Break generation">,
    ], "event generation register">,
    Register<"CCMR1_OUTPUT", 0x18, 32, [
      Field<"CC1S", 0, 2, ReadWrite, "Capture/Compare 1 selection">,
      Field<"OC1FE", 2, 1, ReadWrite, "Output Compare 1 fast enable">,
      Field<"OC1PE", 3, 1, ReadWrite, "Output Compare 1 preload enable">,
      Field<"OC1M", 4, 3, ReadWrite, "Output Compare 1 mode">,
      Field<"CC2S", 8, 2, ReadWrite, "Capture/Compare 2 selection">,
      Field<"OC2FE", 10, 1, ReadWrite, "Output Compare 2 fast enable">,
      Field<"OC2PE", 11, 1, ReadWrite, "Output Compare 2 preload enable">,
      Field<"OC2M", 12, 3, ReadWrite, "Output Compare 2 mode">,
      Field<"OC1M_3", 16, 1, ReadWrite, "Output Compare 1 mode bit 3">,
      Field<"OC2M_3", 24, 1, ReadWrite, "Output Compare 2 mode bit 3">,
    ], "capture/compare mode register (output mode)">,
    Register<"CCMR1_INPUT", 0x18, 32, [
      Field<"CC1S", 0, 2, ReadWrite, "Capture/Compare 1 selection">,
      Field<"IC1PSC", 2, 2, ReadWrite, "Input capture 1 prescaler">,
      Field<"IC1F", 4, 4, ReadWrite, "Input capture 1 filter">,
      Field<"CC2S", 8, 2, ReadWrite, "Capture/Compare 2 selection">,
      Field<"IC2PSC", 10, 2, ReadWrite, "Input capture 2 prescaler">,
      Field<"IC2F", 12, 4, ReadWrite, "Input capture 2 filter">,
    ], "capture/compare mode register 1 (input mode)">,
    Register<"CCER", 0x20, 32, [
      Field<"CC1E", 0, 1, ReadWrite, "Capture/Compare 1 output enable">,
      Field<"CC1P", 1, 1, ReadWrite, "Capture/Compare 1 output Polarity">,
      Field<"CC1NE", 2, 1, ReadWrite, "Capture/Compare 1 complementary output enable">,
      Field<"CC1NP", 3, 1, ReadWrite, "Capture/Compare 1 output Polarity">,
      Field<"CC2E", 4, 1, ReadWrite, "Capture/Compare 2 output enable">,
      Field<"CC2P", 5, 1, ReadWrite, "Capture/Compare 2 output Polarity">,
      Field<"CC2NP", 7, 1, ReadWrite, "Capture/Compare 2 output Polarity">,
    ], "capture/compare enable register">,
    Register<"CNT", 0x24, 32, [
      Field<"CNT", 0, 16, ReadWrite, "counter value">,
      Field<"UIFCPY", 31, 1, Read, "UIF copy">,
    ], "counter">,
    Register<"PSC", 0x28, 32, [
      Field<"PSC", 0, 16, ReadWrite, "Prescaler value">,
    ], "prescaler">,
    Register<"ARR", 0x2c, 32, [
      Field<"ARR", 0, 16, ReadWrite, "Auto-reload value">,
    ], "auto-reload register">,
    Register<"RCR", 0x30, 32, [
      Field<"REP", 0, 8, ReadWrite, "Repetition counter value">,
    ], "repetition counter register">,
    Register<"CCR1", 0x34, 32, [
      Field<"CCR1", 0, 16, ReadWrite, "Capture/Compare 1 value">,
    ], "capture/compare register 1">,
    Register<"CCR2", 0x38, 32, [
      Field<"CCR2", 0, 16, ReadWrite, "Capture/Compare 2 value">,
    ], "capture/compare register 2">,
    Register<"BDTR", 0x44, 32, [
      Field<"DTG", 0, 8, ReadWrite, "Dead-time generator setup">,
      Field<"LOCK", 8, 2, ReadWrite, "Lock configuration">,
      Field<"OSSI", 10, 1, ReadWrite, "Off-state selection for Idle mode">,
      Field<"OSSR", 11, 1, ReadWrite, "Off-state selection for Run mode">,
      Field<"BKE", 12, 1, ReadWrite, "Break enable">,
      Field<"BKP", 13, 1, ReadWrite, "Break polarity">,
      Field<"AOE", 14, 1, ReadWrite, "Automatic output enable">,
      Field<"MOE", 15, 1, ReadWrite, "Main output enable">,
      Field<"BKF", 16, 4, ReadWrite, "Break filter">,
    ], "break and dead-time register">,
    Register<"DCR", 0x48, 32, [
      Field<"DBA", 0, 5, ReadWrite, "DMA base address">,
      Field<"DBL", 8, 5, ReadWrite, "DMA burst length">,
    ], "DMA control register">,
    Register<"DMAR", 0x4c, 32, [
      Field<"DMAB", 0, 16, ReadWrite, "DMA register for burst accesses">,
    ], "DMA address for full transfer">,
    Register<"AF", 0x60, 32, [
      Field<"BKINE", 0, 1, ReadWrite, "BRK BKIN input enable">,
      Field<"BKCMP1E", 1, 1, ReadWrite, "BRK COMP1 enable">,
      Field<"BKCMP2E", 2, 1, ReadWrite, "BRK COMP2 enable">,
      Field<"BKDF1BK0E", 8, 1, ReadWrite, "BRK dfsdm1_break[0] enable">,
      Field<"BKINP", 9, 1, ReadWrite, "BRK BKIN input polarity">,
      Field<"BKCMP1P", 10, 1, ReadWrite, "BRK COMP1 input polarity">,
      Field<"BKCMP2P", 11, 1, ReadWrite, "BRK COMP2 input polarity">,
    ], "alternate fdfsdm1_breakon register 1">,
    Register<"TISEL", 0x68, 32, [
      Field<"TI1SEL", 0, 4, ReadWrite, "selects TI1[0] to TI1[15] input">,
      Field<"TI2SEL", 8, 4, ReadWrite, "selects TI2[0] to TI2[15] input">,
    ], "input selection register">,
  ];
}

class TIM15Instance<string Name, int Base> : PeripheralInstance<Name, Base, TIM15Peripheral>;
def TIM15 : TIM15Instance<"TIM15", 0x40014000>;

#endif // _PERIPHERALS_TIM15_TD
