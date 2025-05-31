#ifndef _PERIPHERALS_EWICIS_TD
#define _PERIPHERALS_EWICIS_TD

include "base.td"

def EWICISPeripheral : PeripheralType<"EWICIS", "External Wakeup Interrupt Controller Interrupt Status Access"> {
  let accessWidth = 32;
  let registers = [
    Register<"EVENTSPR", 0x0, 32, [
      Field<"NMI", 1, 1, ReadWrite, "Causes the processor to behave like a Non-maskable Interrupt (NMI) has occurred">,
      Field<"EVENT", 1, 1, ReadWrite, "Causes the processor to behave like an RXEV event has occurred">,
      Field<"EDBGREQ", 2, 1, ReadWrite, "Causes the processor to behave like an external debug request has occurred">,
    ], "Set pending events at wakeup that cannot be directly set in the Nested Vectored Interrupt Controller (NVIC) using the architecture programming model">,
    Register<"EVENTMASKA", 0x80, 32, [
      Field<"EVENT", 0, 1, ReadWrite, "Sensitive to RXEV when in WFE sleep">,
      Field<"NMI", 1, 1, ReadWrite, "Mask for Non-Maskable Interrupt (NMI)">,
      Field<"EDBGREQ", 2, 1, ReadWrite, "Mask for external debug request">,
    ], "Provides the events on sleep entry which cause the processor to wake up, including information about internal events">,
    RepeatingRegister<"EVENTMASK", 32, 0x84, 32, [
      Field<"IRQ", 0, 32, ReadWrite, "Masks for interrupts ((n-1)x32) to (nx32)-1.">,
    ], "Provides the events on sleep entry which cause the processor to wake up, including information about external events">,
  ];
}

class EWICISInstance<string Name, int Base> : PeripheralInstance<Name, Base, EWICISPeripheral>;
def EWICIS : EWICISInstance<"EWICIS", 0xe001e400>;

#endif // _PERIPHERALS_EWICIS_TD
