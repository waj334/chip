#ifndef _PERIPHERALS_NVIC_TD
#define _PERIPHERALS_NVIC_TD

include "base.td"

def NVICPeripheral : PeripheralType<"NVIC", "Nested Vectored Interrupt Controller"> {
  let accessWidth = 32;
  let registers = [
    RepeatingRegister<"ISER", 16, 0x0, 32, [
      Field<"SETENA", 0, 32, ReadWrite, "Interrupt set-enable bits. For SETENA[m] in NVIC_ISERn, allows interrupt 32n+m to be accessed.">,
    ], "Enables, or reads the enable state of a group of interrupts">,
    RepeatingRegister<"ICER", 16, 0x80, 32, [
      Field<"CLRENA", 0, 32, ReadWrite, "Interrupt clear-enable bits. For SETENA[m] in NVIC_ICERn, allows interrupt 32n+m to be accessed.">,
    ], "Clears, or reads the enable state of, a group of registers">,
    RepeatingRegister<"ISPR", 16, 0x100, 32, [
      Field<"SETPEND", 0, 32, ReadWrite, "Interrupt set-pending bits. For SETPEND[m] in NVIC_ISPRn, allows interrupt 32n + m to be accessed.">,
    ], "For a group of interrupts, changes interrupt status to pending, or shows the current pending status">,
    RepeatingRegister<"ICPR", 16, 0x180, 32, [
      Field<"CLRPEND", 0, 32, ReadWrite, "Interrupt clear-pending bits. Clears the pending state of interrupt (m+(32*n)), or shows whether the state of the interrupt is pending">,
    ], "For a group of interrupts, clears the interrupt pending status, or shows the current pending status">,
    RepeatingRegister<"IABR", 16, 0x200, 32, [
      Field<"ACTIVE", 0, 32, ReadWrite, "For ACTIVE[m] in NVIC_IABRn, indicates the active state for interrupt 32n+m.">,
    ], "For a group of interrupts, shows whether each interrupt is active">,
    RepeatingRegister<"ITNS", 16, 0x280, 32, [
      Field<"ITNS", 0, 32, ReadWrite, "For ITNS[m] in NVIC_ITNSn, this field indicates and allows modification of the target Security state for interrupt 32n+m.">,
    ], "For a group of interrupts, specifies or reads whether each interrupt targets non-secure or secure state">,
    RepeatingRegister<"IPR", 120, 0x300, 32, [
      Field<"PRI_N0", 0, 8, ReadWrite, "For register NVIC_IPRn, priority of interrupt number 4n">,
      Field<"PRI_N1", 8, 8, ReadWrite, "For register NVIC_IPRn, priority of interrupt number 4n+1">,
      Field<"PRI_N2", 16, 8, ReadWrite, "For register NVIC_IPRn, priority of interrupt number 4n+2">,
      Field<"PRI_N3", 24, 8, ReadWrite, "For register NVIC_IPRn, priority of interrupt number 4n+3">,
    ], "Sets or reads interrupt priorities">,
    Register<"STIR", 0xe00, 32, [
      Field<"INTID", 0, 9, ReadWrite, "Interrupt ID of the interrupt to trigger, in the range 0-479.">,
    ], "Write to the STIR to generate an interrupt from software.">,
  ];
}

class NVICInstance<string Name, int Base> : PeripheralInstance<Name, Base, NVICPeripheral>;
def NVIC : NVICInstance<"NVIC", 0xe000e100>;

#endif // _PERIPHERALS_NVIC_TD
