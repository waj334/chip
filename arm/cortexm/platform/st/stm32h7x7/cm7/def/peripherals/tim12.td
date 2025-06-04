#ifndef _PERIPHERALS_TIM12_TD
#define _PERIPHERALS_TIM12_TD

include "base.td"

def TIM12Peripheral : PeripheralType<"TIM12", "General purpose timers"> {
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
      Field<"MMS", 4, 3, ReadWrite, "Master mode selection">,
      Field<"TI1S", 7, 1, ReadWrite, "TI1 selection">,
    ], "control register 2">,
    Register<"SMCR", 0x8, 32, [
      Field<"SMS", 0, 3, ReadWrite, "Slave mode selection">,
      Field<"TS", 4, 3, ReadWrite, "Trigger selection">,
      Field<"MSM", 7, 1, ReadWrite, "Master/Slave mode">,
      Field<"SMS_3", 16, 1, ReadWrite, "Slave mode selection - bit 3">,
      Field<"TS_4_3", 20, 2, ReadWrite, "Trigger selection">,
    ], "slave mode control register">,
    Register<"DIER", 0xc, 32, [
      Field<"UIE", 0, 1, ReadWrite, "Update interrupt enable">,
      Field<"CC1IE", 1, 1, ReadWrite, "Capture/Compare 1 interrupt enable">,
      Field<"CC2IE", 2, 1, ReadWrite, "Capture/Compare 2 interrupt enable">,
      Field<"TIE", 6, 1, ReadWrite, "Trigger interrupt enable">,
    ], "DMA/Interrupt enable register">,
    Register<"SR", 0x10, 32, [
      Field<"UIF", 0, 1, ReadWrite, "Update interrupt flag">,
      Field<"CC1IF", 1, 1, ReadWrite, "Capture/compare 1 interrupt flag">,
      Field<"CC2IF", 2, 1, ReadWrite, "Capture/Compare 2 interrupt flag">,
      Field<"TIF", 6, 1, ReadWrite, "Trigger interrupt flag">,
    ], "status register">,
    Register<"EGR", 0x14, 32, [
      Field<"UG", 0, 1, ReadWrite, "Update generation">,
      Field<"CC1G", 1, 1, ReadWrite, "Capture/compare 1 generation">,
      Field<"CC2G", 2, 1, ReadWrite, "Capture/compare 2 generation">,
      Field<"TG", 6, 1, ReadWrite, "Trigger generation">,
    ], "event generation register">,
    Register<"CCMR1", 0x18, 32, [
      Field<"CC1S", 0, 2, ReadWrite, "CC1S">,
      Field<"OC1FE", 2, 1, ReadWrite, "OC1FE">,
      Field<"OC1PE", 3, 1, ReadWrite, "OC1PE">,
      Field<"OC1M", 4, 3, ReadWrite, "OC1M">,
      Field<"CC2S", 8, 2, ReadWrite, "CC2S">,
      Field<"OC2FE", 10, 1, ReadWrite, "OC2FE">,
      Field<"OC2PE", 11, 1, ReadWrite, "OC2PE">,
      Field<"OC2M", 12, 3, ReadWrite, "OC2M">,
      Field<"OC1M_3", 16, 1, ReadWrite, "Output Compare 1 mode - bit 3">,
      Field<"OC2M_3", 24, 1, ReadWrite, "Output Compare 2 mode - bit 3">,
      Field<"ICPCS", 2, 2, ReadWrite, "Input capture 1 prescaler">,
      Field<"IC1F", 4, 4, ReadWrite, "Input capture 1 filter">,
      Field<"IC2PCS", 10, 2, ReadWrite, "Input capture 2 prescaler">,
      Field<"IC2F", 12, 4, ReadWrite, "Input capture 2 filter">,
    ], "capture/compare mode register 1 (output mode)">,
    Register<"CCER", 0x20, 32, [
      Field<"CC1E", 0, 1, ReadWrite, "Capture/Compare 1 output enable">,
      Field<"CC1P", 1, 1, ReadWrite, "Capture/Compare 1 output Polarity">,
      Field<"CC1NP", 3, 1, ReadWrite, "Capture/Compare 1 output Polarity">,
      Field<"CC2E", 4, 1, ReadWrite, "Capture/Compare 2 output enable">,
      Field<"CC2P", 5, 1, ReadWrite, "Capture/Compare 2 output Polarity">,
      Field<"CC2NP", 7, 1, ReadWrite, "Capture/Compare 2 output Polarity">,
    ], "capture/compare enable register">,
    Register<"CNT", 0x24, 32, [
      Field<"CNT", 0, 16, ReadWrite, "low counter value">,
    ], "counter">,
    Register<"PSC", 0x28, 32, [
      Field<"PSC", 0, 16, ReadWrite, "Prescaler value">,
    ], "prescaler">,
    Register<"ARR", 0x2c, 32, [
      Field<"ARR", 0, 16, ReadWrite, "Low Auto-reload value">,
    ], "auto-reload register">,
    Register<"CCR1", 0x34, 32, [
      Field<"CCR1", 0, 16, ReadWrite, "Low Capture/Compare 1 value">,
    ], "capture/compare register 1">,
    Register<"CCR2", 0x38, 32, [
      Field<"CCR2", 0, 16, ReadWrite, "Low Capture/Compare 2 value">,
    ], "capture/compare register 2">,
    Register<"TISEL", 0x68, 32, [
      Field<"TI1SEL", 0, 4, ReadWrite, "TI1[0] to TI1[15] input selection">,
      Field<"TI2SEL", 8, 4, ReadWrite, "TI2[0] to TI2[15] input selection">,
    ], "timer input selection register">,
  ];
}

class TIM12Instance<string Name, int Base> : PeripheralInstance<Name, Base, TIM12Peripheral>;
def TIM12 : TIM12Instance<"TIM12", 0x40001800>;


#endif // _PERIPHERALS_TIM12_TD
