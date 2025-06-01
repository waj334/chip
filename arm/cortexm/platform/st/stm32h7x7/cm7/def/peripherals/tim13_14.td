#ifndef _PERIPHERALS_TIM2_TD
#define _PERIPHERALS_TIM2_TD

include "base.td"

def TIM13_14Peripheral : PeripheralType<"TIM13_14", "General purpose timers"> {
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
    Register<"DIER", 0xc, 32, [
      Field<"UIE", 0, 1, ReadWrite, "Update interrupt enable">,
      Field<"CCIE", 1, 1, ReadWrite, "Capture/Compare interrupt enable">,
    ], "DMA/Interrupt enable register">,
    Register<"SR", 0x10, 32, [
      Field<"UIF", 0, 1, ReadWrite, "Update interrupt flag">,
      Field<"CCIF", 1, 1, ReadWrite, "Capture/compare interrupt flag">,
    ], "status register">,
    Register<"EGR", 0x14, 32, [
      Field<"UG", 0, 1, ReadWrite, "Update generation">,
      Field<"CC1G", 1, 1, ReadWrite, "Capture/compare generation">,
    ], "event generation register">,
    Register<"CCMR", 0x18, 32, [
      Field<"CCS", 0, 2, ReadWrite, "Output CCS">,
      Field<"OCFE", 2, 1, ReadWrite, "Output OCFE">,
      Field<"ICPCS", 2, 2, ReadWrite, "Input capture prescaler">,
      Field<"OCPE", 3, 1, ReadWrite, "Output OCPE">,
      Field<"OCM", 4, 3, ReadWrite, "Output OCM">,
      Field<"ICF", 4, 4, ReadWrite, "Input capture filter">,
      Field<"OCM_3", 16, 1, ReadWrite, "Output Compare mode - bit 3">,
    ], "capture/compare mode register 1 (output mode)">,
    Register<"CCER", 0x20, 32, [
      Field<"CCE", 0, 1, ReadWrite, "Capture/Compare output enable">,
      Field<"CCP", 1, 1, ReadWrite, "Capture/Compare output Polarity">,
      Field<"CCNP", 3, 1, ReadWrite, "Capture/Compare output Polarity">,
    ], "capture/compare enable register">,
    Register<"CNT", 0x24, 32, [
      Field<"CNT", 0, 16, ReadWrite, "low counter value">,
      Field<"UIFCPY", 31, 1, ReadWrite, "UIF Copy - This bit is a read-only copy of the UIF bit in the TIMx_ISR register.">,
    ], "counter">,
    Register<"PSC", 0x28, 32, [
      Field<"PSC", 0, 16, ReadWrite, "Prescaler value">,
    ], "prescaler">,
    Register<"ARR", 0x2c, 32, [
      Field<"ARR", 0, 16, ReadWrite, "Auto-reload value">,
    ], "auto-reload register">,
    Register<"CCR", 0x34, 32, [
      Field<"CCR", 0, 16, ReadWrite, "Capture/Compare value">,
    ], "capture/compare register 1">,
    Register<"TISEL", 0x68, 32, [
      Field<"TI1SEL", 0, 4, ReadWrite, "TI1[0] to TI1[15] input selection">,
    ], "timer input selection register">,
  ];
}

class TIM13_14Instance<string Name, int Base> : PeripheralInstance<Name, Base, TIM13_14Peripheral>;
def TIM13 : TIM13_14Instance<"TIM13", 0x40001c00>;
def TIM14 : TIM13_14Instance<"TIM14", 0x40002000>;

def TIM13_14Group : PeripheralGroup<[TIM13, TIM14]>;

#endif // _PERIPHERALS_TIM13_14_TD
