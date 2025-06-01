#ifndef _PERIPHERALS_TIM16_17_TD
#define _PERIPHERALS_TIM16_17_TD

include "base.td"

def TIM16_17Peripheral : PeripheralType<"TIM16_17", "General-purpose-timers"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
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
      Field<"OIS1", 8, 1, ReadWrite, "Output Idle state 1">,
      Field<"OIS1N", 9, 1, ReadWrite, "Output Idle state 1">,
    ], "control register 2">,
    Register<"DIER", 0xc, 32, [
      Field<"UIE", 0, 1, ReadWrite, "Update interrupt enable">,
      Field<"CCIE", 1, 1, ReadWrite, "Capture/Compare interrupt enable">,
      Field<"COMIE", 5, 1, ReadWrite, "COM interrupt enable">,
      Field<"BIE", 7, 1, ReadWrite, "Break interrupt enable">,
      Field<"UDE", 8, 1, ReadWrite, "Update DMA request enable">,
      Field<"CCDE", 9, 1, ReadWrite, "Capture/Compare DMA request enable">,
      Field<"COMDE", 13, 1, ReadWrite, "COM DMA request enable">,
    ], "DMA/Interrupt enable register">,
    Register<"SR", 0x10, 32, [
      Field<"UIF", 0, 1, ReadWrite, "Update interrupt flag">,
      Field<"CCIF", 1, 1, ReadWrite, "Capture/compare interrupt flag">,
      Field<"COMIF", 5, 1, ReadWrite, "COM interrupt flag">,
      Field<"BIF", 7, 1, ReadWrite, "Break interrupt flag">,
      Field<"CCOF", 9, 1, ReadWrite, "Capture/Compare overcapture flag">,
    ], "status register">,
    Register<"EGR", 0x14, 32, [
      Field<"UG", 0, 1, ReadWrite, "Update generation">,
      Field<"CCG", 1, 1, ReadWrite, "Capture/compare generation">,
      Field<"COMG", 5, 1, ReadWrite, "Capture/Compare control update generation">,
      Field<"BG", 7, 1, ReadWrite, "Break generation">,
    ], "event generation register">,
    Register<"CCMR", 0x18, 32, [
      Field<"CCS", 0, 2, ReadWrite, "Capture/Compare selection">,
      Field<"OCFE", 2, 1, ReadWrite, "Output Compare fast enable">,
      Field<"OCPE", 3, 1, ReadWrite, "Output Compare preload enable">,
      Field<"OCM", 4, 3, ReadWrite, "Output Compare mode">,
      Field<"OCM_3", 16, 1, ReadWrite, "Output Compare mode">,
      Field<"CCS", 0, 2, ReadWrite, "Capture/Compare selection">,
      Field<"ICPSC", 2, 2, ReadWrite, "Input capture prescaler">,
      Field<"ICF", 4, 4, ReadWrite, "Input capture filter">,
    ], "capture/compare mode register">,
    Register<"CCER", 0x20, 32, [
      Field<"CCE", 0, 1, ReadWrite, "Capture/Compare output enable">,
      Field<"CCP", 1, 1, ReadWrite, "Capture/Compare output Polarity">,
      Field<"CCNE", 2, 1, ReadWrite, "Capture/Compare complementary output enable">,
      Field<"CCNP", 3, 1, ReadWrite, "Capture/Compare output Polarity">,
    ], "capture/compare enable register">,
    Register<"CNT", 0x24, 32, [
      Field<"CNT", 0, 16, ReadWrite, "counter value">,
      Field<"UIFCPY", 31, 1, Read, "UIF Copy">,
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
    Register<"CCR", 0x34, 32, [
      Field<"CCR", 0, 16, ReadWrite, "Capture/Compare value">,
    ], "capture/compare register 1">,
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
      Field<"BKDFBK1E", 8, 1, ReadWrite, "BRK dfsdm1_break[1] enable">,
      Field<"BKINP", 9, 1, ReadWrite, "BRK BKIN input polarity">,
      Field<"BKCMP1P", 10, 1, ReadWrite, "BRK COMP1 input polarity">,
      Field<"BKCMP2P", 11, 1, ReadWrite, "BRK COMP2 input polarity">,
    ], "alternate function register 1">,
    Register<"TISEL", 0x68, 32, [
      Field<"TI1SEL", 0, 4, ReadWrite, "selects TI1[0] to TI1[15] input">,
    ], "input selection register">,
  ];
}

class TIM16_17Instance<string Name, int Base> : PeripheralInstance<Name, Base, TIM16_17Peripheral>;
def TIM16 : TIM16_17Instance<"TIM16", 0x40014400>;
def TIM17 : TIM16_17Instance<"TIM17", 0x40014800>;

def TIM16_17Group : PeripheralGroup<[ TIM16, TIM17 ]>;

#endif // _PERIPHERALS_TIM16_17_TD
