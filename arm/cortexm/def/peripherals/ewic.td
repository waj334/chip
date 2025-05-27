#ifndef _PERIPHERALS_EWIC_TD
#define _PERIPHERALS_EWIC_TD

include "base.td"

def EWICPeripheral : PeripheralType<"EWIC", "External Wakeup Interrupt Controller"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"EN", 0, 1, ReadWrite, "Enable">,
    ], "The main External Wakeup Interrupt Controller (EWIC) control register">,
    Register<"ASCR", 0x4, 32, [
      Field<"ASPD", 0, 1, ReadWrite, "The processor must use this value to decide whether any automatic EWIC accesses must be performed on transitioning to a low-power state">,
      Field<"ASPU", 1, 1, ReadWrite, "The processor must use this value to decide whether any automatic EWIC accesses must be performed on transitioning from a low-power state">,
    ], "Determines whether the processor generates APB transactions on entry and exit from Wakeup Interrupt Controller (WIC) sleep to set up the wakeup state in the External Wakeup Interrupt Controller (EWIC)">,
    Register<"CLRMASK", 0x8, 32, [
    ], "Causes EWIC_MASKA and all the EWIC_MASKn registers to be cleared. The write data is ignored. This register is RAZ.">,
    Register<"NUMID", 0xc, 32, [
      Field<"NUMEVENT", 0, 16, ReadWrite, "The number of events supported">,
    ], "Returns the total number of events supported in the External Wakeup Interrupt Controller (EWIC) that have been configured during synthesis">,
    Register<"MASKA", 0x200, 32, [
      Field<"EVENT", 0, 1, ReadWrite, "Mask for Wait For Exception (WFE) wakeup event">,
      Field<"NMI", 1, 1, ReadWrite, "Mask for Non-Maskable Interrupt (NMI)">,
      Field<"EDBGREQ", 2, 1, ReadWrite, "Mask for external debug request">,
    ], "Defines the mask for special events and the EWIC_MASKn registers for external interrupt (IRQ) events">,
    RepeatingRegister<"MASK", 15, 0x204, 32, [
      Field<"IRQ", 0, 32, ReadWrite, "Masks for interrupts (nx32) to ((n+1)x32)-1">,
    ], "Defines the mask for special events and the EWIC_MASKn registers for external interrupt (IRQ) events">,
    Register<"PENDA", 0x400, 32, [
      Field<"EVENT", 0, 1, ReadWrite, "Wait For Exception (WFE) wakeup event is pended">,
      Field<"NMI", 1, 1, ReadWrite, "Non-Maskable Interrupt (NMI) is pended">,
      Field<"EDBGREQ", 2, 1, ReadWrite, "External debug request is pended">,
    ], "Indicate which events have been pended">,
    RepeatingRegister<"PEND", 15, 0x404, 32, [
      Field<"IRQ", 0, 32, ReadWrite, "Interrupts (nx32) to ((n+1)x32)-1 are pended. A write of zero to this field is ignored.">,
    ], "Indicate which events have been pended">,
    Register<"PSR", 0x600, 32, [
      Field<"NZA", 0, 1, ReadWrite, "If EWIC_PSR.NZA set, then EWIC_PENDA is non-zero">,
      Field<"NZ", 1, 15, ReadWrite, "If EWIC_PSR.NZ[n+1] is set, then EWIC_PENDn is non-zero">,
    ], "Indicates which EWIC_PENDn registers are non-zero">,
  ];
}

class EWICInstance<string Name, int Base> : PeripheralInstance<Name, Base, EWICPeripheral>;
def EWIC : EWICInstance<"EWIC", 0xe0047000>;

#endif // _PERIPHERALS_EWIC_TD
