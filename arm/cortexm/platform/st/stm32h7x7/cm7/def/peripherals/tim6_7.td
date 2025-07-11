#ifndef _PERIPHERALS_TIM6_7_TD
#define _PERIPHERALS_TIM6_7_TD

include "base.td"

def TIM6_7Peripheral : PeripheralType<"TIM6_7", "Basic timers"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR1", 0x0, 32, [
      Field<"CEN", 0, 1, ReadWrite, "Counter enable">,
      Field<"UDIS", 1, 1, ReadWrite, "Update disable">,
      Field<"URS", 2, 1, ReadWrite, "Update request source">,
      Field<"OPM", 3, 1, ReadWrite, "One-pulse mode">,
      Field<"ARPE", 7, 1, ReadWrite, "Auto-reload preload enable">,
      Field<"UIFREMAP", 11, 1, ReadWrite, "UIF status bit remapping">,
    ], "control register 1">,
    Register<"CR2", 0x4, 32, [
      Field<"MMS", 4, 3, ReadWrite, "Master mode selection">,
    ], "control register 2">,
    Register<"DIER", 0xc, 32, [
      Field<"UIE", 0, 1, ReadWrite, "Update interrupt enable">,
      Field<"UDE", 8, 1, ReadWrite, "Update DMA request enable">,
    ], "DMA/Interrupt enable register">,
    Register<"SR", 0x10, 32, [
      Field<"UIF", 0, 1, ReadWrite, "Update interrupt flag">,
    ], "status register">,
    Register<"EGR", 0x14, 32, [
      Field<"UG", 0, 1, ReadWrite, "Update generation">,
    ], "event generation register">,
    Register<"CNT", 0x24, 32, [
      Field<"CNT", 0, 16, ReadWrite, "Low counter value">,
      Field<"UIFCPY", 31, 1, ReadWrite, "UIF Copy">,
    ], "counter">,
    Register<"PSC", 0x28, 32, [
      Field<"PSC", 0, 16, ReadWrite, "Prescaler value">,
    ], "prescaler">,
    Register<"ARR", 0x2c, 32, [
      Field<"ARR", 0, 16, ReadWrite, "Low Auto-reload value">,
    ], "auto-reload register">,
  ];
}

class TIM6_7Instance<string Name, int Base> : PeripheralInstance<Name, Base, TIM6_7Peripheral>;
def TIM6 : TIM6_7Instance<"TIM6", 0x40001000>;
def TIM7 : TIM6_7Instance<"TIM7", 0x40001400>;

def TIM6_7PeripheralGroup : PeripheralGroup<[ TIM6, TIM7 ]>;

#endif // _PERIPHERALS_TIM6_7_TD
