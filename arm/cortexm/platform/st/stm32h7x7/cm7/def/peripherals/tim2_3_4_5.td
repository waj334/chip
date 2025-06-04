#ifndef _PERIPHERALS_TIM2_3_4_5_TD
#define _PERIPHERALS_TIM2_3_4_5_TD

include "base.td"

def TIM2_3_4_5Peripheral : PeripheralType<"TIM2_3_4_5", "General purpose timers"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR1", 0x0, 32, [
      Field<"CEN", 0, 1, ReadWrite, "Counter enable">,
      Field<"UDIS", 1, 1, ReadWrite, "Update disable">,
      Field<"URS", 2, 1, ReadWrite, "Update request source">,
      Field<"OPM", 3, 1, ReadWrite, "One-pulse mode">,
      Field<"DIR", 4, 1, ReadWrite, "Direction">,
      Field<"CMS", 5, 2, ReadWrite, "Center-aligned mode selection">,
      Field<"ARPE", 7, 1, ReadWrite, "Auto-reload preload enable">,
      Field<"CKD", 8, 2, ReadWrite, "Clock division">,
      Field<"UIFREMAP", 11, 1, ReadWrite, "UIF status bit remapping">,
    ], "control register 1">,
    Register<"CR2", 0x4, 32, [
      Field<"CCDS", 3, 1, ReadWrite, "Capture/compare DMA selection">,
      Field<"MMS", 4, 3, ReadWrite, "Master mode selection">,
      Field<"TI1S", 7, 1, ReadWrite, "TI1 selection">,
    ], "control register 2">,
    Register<"SMCR", 0x8, 32, [
      Field<"SMS", 0, 3, ReadWrite, "Slave mode selection">,
      Field<"TS", 4, 3, ReadWrite, "Trigger selection">,
      Field<"MSM", 7, 1, ReadWrite, "Master/Slave mode">,
      Field<"ETF", 8, 4, ReadWrite, "External trigger filter">,
      Field<"ETPS", 12, 2, ReadWrite, "External trigger prescaler">,
      Field<"ECE", 14, 1, ReadWrite, "External clock enable">,
      Field<"ETP", 15, 1, ReadWrite, "External trigger polarity">,
      Field<"SMS_3", 16, 1, ReadWrite, "Slave mode selection - bit 3">,
      Field<"TS_4_3", 20, 2, ReadWrite, "Trigger selection">,
    ], "slave mode control register">,
    Register<"DIER", 0xc, 32, [
      Field<"UIE", 0, 1, ReadWrite, "Update interrupt enable">,
      Field<"CC1IE", 1, 1, ReadWrite, "Capture/Compare 1 interrupt enable">,
      Field<"CC2IE", 2, 1, ReadWrite, "Capture/Compare 2 interrupt enable">,
      Field<"CC3IE", 3, 1, ReadWrite, "Capture/Compare 3 interrupt enable">,
      Field<"CC4IE", 4, 1, ReadWrite, "Capture/Compare 4 interrupt enable">,
      Field<"TIE", 6, 1, ReadWrite, "Trigger interrupt enable">,
      Field<"UDE", 8, 1, ReadWrite, "Update DMA request enable">,
      Field<"CC1DE", 9, 1, ReadWrite, "Capture/Compare 1 DMA request enable">,
      Field<"CC2DE", 10, 1, ReadWrite, "Capture/Compare 2 DMA request enable">,
      Field<"CC3DE", 11, 1, ReadWrite, "Capture/Compare 3 DMA request enable">,
      Field<"CC4DE", 12, 1, ReadWrite, "Capture/Compare 4 DMA request enable">,
      Field<"TDE", 14, 1, ReadWrite, "Trigger DMA request enable">,
    ], "DMA/Interrupt enable register">,
    Register<"SR", 0x10, 32, [
      Field<"UIF", 0, 1, ReadWrite, "Update interrupt flag">,
      Field<"CC1IF", 1, 1, ReadWrite, "Capture/compare 1 interrupt flag">,
      Field<"CC2IF", 2, 1, ReadWrite, "Capture/Compare 2 interrupt flag">,
      Field<"CC3IF", 3, 1, ReadWrite, "Capture/Compare 3 interrupt flag">,
      Field<"CC4IF", 4, 1, ReadWrite, "Capture/Compare 4 interrupt flag">,
      Field<"TIF", 6, 1, ReadWrite, "Trigger interrupt flag">,
      Field<"CC1OF", 9, 1, ReadWrite, "Capture/Compare 1 overcapture flag">,
      Field<"CC2OF", 10, 1, ReadWrite, "Capture/compare 2 overcapture flag">,
      Field<"CC3OF", 11, 1, ReadWrite, "Capture/Compare 3 overcapture flag">,
      Field<"CC4OF", 12, 1, ReadWrite, "Capture/Compare 4 overcapture flag">,
    ], "status register">,
    Register<"EGR", 0x14, 32, [
      Field<"UG", 0, 1, ReadWrite, "Update generation">,
      Field<"CC1G", 1, 1, ReadWrite, "Capture/compare 1 generation">,
      Field<"CC2G", 2, 1, ReadWrite, "Capture/compare 2 generation">,
      Field<"CC3G", 3, 1, ReadWrite, "Capture/compare 3 generation">,
      Field<"CC4G", 4, 1, ReadWrite, "Capture/compare 4 generation">,
      Field<"TG", 6, 1, ReadWrite, "Trigger generation">,
    ], "event generation register">,
    Register<"CCMR1", 0x18, 32, [
      Field<"CC1S", 0, 2, ReadWrite, "CC1S">,
      Field<"OC1FE", 2, 1, ReadWrite, "OC1FE">,
      Field<"OC1PE", 3, 1, ReadWrite, "OC1PE">,
      Field<"OC1M", 4, 3, ReadWrite, "OC1M">,
      Field<"OC1CE", 7, 1, ReadWrite, "OC1CE">,
      Field<"CC2S", 8, 2, ReadWrite, "CC2S">,
      Field<"OC2FE", 10, 1, ReadWrite, "OC2FE">,
      Field<"OC2PE", 11, 1, ReadWrite, "OC2PE">,
      Field<"OC2M", 12, 3, ReadWrite, "OC2M">,
      Field<"OC2CE", 15, 1, ReadWrite, "OC2CE">,
      Field<"OC1M_3", 16, 1, ReadWrite, "Output Compare 1 mode - bit 3">,
      Field<"OC2M_3", 24, 1, ReadWrite, "Output Compare 2 mode - bit 3">,
      Field<"ICPCS", 2, 2, ReadWrite, "Input capture 1 prescaler">,
      Field<"IC1F", 4, 4, ReadWrite, "Input capture 1 filter">,
      Field<"IC2PCS", 10, 2, ReadWrite, "Input capture 2 prescaler">,
      Field<"IC2F", 12, 4, ReadWrite, "Input capture 2 filter">,
    ], "capture/compare mode register 1">,
    Register<"CCMR2", 0x1c, 32, [
      Field<"CC3S", 0, 2, ReadWrite, "CC3S">,
      Field<"OC3FE", 2, 1, ReadWrite, "OC3FE">,
      Field<"OC3PE", 3, 1, ReadWrite, "OC3PE">,
      Field<"OC3M", 4, 3, ReadWrite, "OC3M">,
      Field<"OC3CE", 7, 1, ReadWrite, "OC3CE">,
      Field<"CC4S", 8, 2, ReadWrite, "CC4S">,
      Field<"OC4FE", 10, 1, ReadWrite, "OC4FE">,
      Field<"OC4PE", 11, 1, ReadWrite, "OC4PE">,
      Field<"OC4M", 12, 3, ReadWrite, "OC4M">,
      Field<"OC4CE", 15, 1, ReadWrite, "OC4CE">,
      Field<"OC3M_3", 16, 1, ReadWrite, "Output Compare 1 mode - bit 3">,
      Field<"OC4M_3", 24, 1, ReadWrite, "Output Compare 2 mode - bit 3">,
      Field<"IC3PSC", 2, 2, ReadWrite, "Input capture 3 prescaler">,
      Field<"IC3F", 4, 4, ReadWrite, "Input capture 3 filter">,
      Field<"IC4PSC", 10, 2, ReadWrite, "Input capture 4 prescaler">,
      Field<"IC4F", 12, 4, ReadWrite, "Input capture 4 filter">,
    ], "capture/compare mode register 2">,
    Register<"CCER", 0x20, 32, [
      Field<"CC1E", 0, 1, ReadWrite, "Capture/Compare 1 output enable">,
      Field<"CC1P", 1, 1, ReadWrite, "Capture/Compare 1 output Polarity">,
      Field<"CC1NP", 3, 1, ReadWrite, "Capture/Compare 1 output Polarity">,
      Field<"CC2E", 4, 1, ReadWrite, "Capture/Compare 2 output enable">,
      Field<"CC2P", 5, 1, ReadWrite, "Capture/Compare 2 output Polarity">,
      Field<"CC2NP", 7, 1, ReadWrite, "Capture/Compare 2 output Polarity">,
      Field<"CC3E", 8, 1, ReadWrite, "Capture/Compare 3 output enable">,
      Field<"CC3P", 9, 1, ReadWrite, "Capture/Compare 3 output Polarity">,
      Field<"CC3NP", 11, 1, ReadWrite, "Capture/Compare 3 output Polarity">,
      Field<"CC4E", 12, 1, ReadWrite, "Capture/Compare 4 output enable">,
      Field<"CC4P", 13, 1, ReadWrite, "Capture/Compare 3 output Polarity">,
      Field<"CC4NP", 15, 1, ReadWrite, "Capture/Compare 4 output Polarity">,
    ], "capture/compare enable register">,
    Register<"CNT", 0x24, 32, [
      Field<"CNT_L", 0, 16, ReadWrite, "low counter value">,
      Field<"CNT_H", 16, 16, ReadWrite, "High counter value">,
    ], "counter">,
    Register<"PSC", 0x28, 32, [
      Field<"PSC", 0, 16, ReadWrite, "Prescaler value">,
    ], "prescaler">,
    Register<"ARR", 0x2c, 32, [
      Field<"ARR_L", 0, 16, ReadWrite, "Low Auto-reload value">,
      Field<"ARR_H", 16, 16, ReadWrite, "High Auto-reload value">,
    ], "auto-reload register">,
    Register<"CCR1", 0x34, 32, [
      Field<"CCR1_L", 0, 16, ReadWrite, "Low Capture/Compare 1 value">,
      Field<"CCR1_H", 16, 16, ReadWrite, "High Capture/Compare 1 value">,
    ], "capture/compare register 1">,
    Register<"CCR2", 0x38, 32, [
      Field<"CCR2_L", 0, 16, ReadWrite, "Low Capture/Compare 2 value">,
      Field<"CCR2_H", 16, 16, ReadWrite, "High Capture/Compare 2 value">,
    ], "capture/compare register 2">,
    Register<"CCR3", 0x3c, 32, [
      Field<"CCR3_L", 0, 16, ReadWrite, "Low Capture/Compare value">,
      Field<"CCR3_H", 16, 16, ReadWrite, "High Capture/Compare value">,
    ], "capture/compare register 3">,
    Register<"CCR4", 0x40, 32, [
      Field<"CCR4_L", 0, 16, ReadWrite, "Low Capture/Compare value">,
      Field<"CCR4_H", 16, 16, ReadWrite, "High Capture/Compare value">,
    ], "capture/compare register 4">,
    Register<"DCR", 0x48, 32, [
      Field<"DBA", 0, 5, ReadWrite, "DMA base address">,
      Field<"DBL", 8, 5, ReadWrite, "DMA burst length">,
    ], "DMA control register">,
    Register<"DMAR", 0x4c, 32, [
      Field<"DMAB", 0, 16, ReadWrite, "DMA register for burst accesses">,
    ], "DMA address for full transfer">,
    Register<"AF1", 0x60, 32, [
      Field<"ETRSEL", 14, 4, ReadWrite, "ETR source selection">,
    ], "alternate function option register 1">,
    Register<"TISEL", 0x68, 32, [
      Field<"TI1SEL", 0, 4, ReadWrite, "TI1[0] to TI1[15] input selection">,
      Field<"TI2SEL", 8, 4, ReadWrite, "TI2[0] to TI2[15] input selection">,
      Field<"TI3SEL", 16, 4, ReadWrite, "TI3[0] to TI3[15] input selection">,
      Field<"TI4SEL", 24, 4, ReadWrite, "TI4[0] to TI4[15] input selection">,
    ], "timer input selection register">,
  ];
}

class TIM2_3_4_5Instance<string Name, int Base> : PeripheralInstance<Name, Base, TIM2_3_4_5Peripheral>;
def TIM2 : TIM2_3_4_5Instance<"TIM2", 0x40000000>;
def TIM3 : TIM2_3_4_5Instance<"TIM3", 0x40000400>;
def TIM4 : TIM2_3_4_5Instance<"TIM4", 0x40000800>;
def TIM5 : TIM2_3_4_5Instance<"TIM5", 0x40000c00>;

def TIM2_3_4_5PeripheralGroup : PeripheralGroup<[ TIM2, TIM3, TIM4,TIM5 ]>;

#endif // _PERIPHERALS_TIM2_3_4_5_TD
