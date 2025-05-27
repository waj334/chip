#ifndef _PERIPHERALS_SAU_TD
#define _PERIPHERALS_SAU_TD

include "base.td"

def SAUPeripheral : PeripheralType<"SAU", "Security Attribution Unit"> {
  let accessWidth = 32;
  let registers = [
    Register<"CTRL", 0x0, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Enables the (SAU)">,
      Field<"ALLNS", 1, 1, ReadWrite, "When SAU_CTRL.ENABLE is 0 this bit controls whether the memory is marked as Non-secure or Secure">,
    ], "Allows enabling of the Security Attribution Unit (SAU)">,
    Register<"TYPE", 0x4, 32, [
      Field<"SREGION", 0, 8, ReadWrite, "The number of implemented (SAU) regions">,
    ], "Indicates the number of regions implemented by the Security Attribution Unit (SAU)">,
    Register<"RNR", 0x8, 32, [
      Field<"REGION", 0, 8, ReadWrite, "Indicates the Security Attribution Unit (SAU) region accessed by SAU_RBAR and SAU_RLAR">,
    ], "Selects the region currently accessed by SAU_RBAR and SAU_RLAR">,
    Register<"RBAR", 0xc, 32, [
      Field<"BADDR", 5, 27, ReadWrite, "Holds bits[31:5] of the base address for the selected (SAU) region">,
    ], "Provides indirect read and write access to the base address of the currently selected Security Attribution Unit (SAU) region">,
    Register<"RLAR", 0x10, 32, [
      Field<"LADDR", 5, 27, ReadWrite, "Holds bits[31:5] of the limit address for the selected (SAU) region">,
      Field<"NSC", 5, 27, ReadWrite, "Controls whether Non-secure state is permitted to execute an SG instruction from this region">,
      Field<"ENABLE", 5, 27, ReadWrite, "(SAU) region enable">,
    ], "Provides indirect read and write access to the limit address of the currently selected Security Attribution Unit (SAU) region">,
    Register<"SFSR", 0x14, 32, [
      Field<"INVEP", 0, 1, ReadWrite, "Inidcates a function call from the Non-secure state or exception targets a non-SG instruction in the Secure state">,
      Field<"INVIS", 1, 1, ReadWrite, "Indicates the integrity signature in an exception stack frame is found to be invalid during the unstacking operation">,
      Field<"INVER", 2, 1, ReadWrite, "Indicates EXC_RETURN.DCRS being set to 0 when returning from an exception in the Non-secure state, or EXC_RETURN.ES being set to 1 when returning from an exception in the Non-secure state">,
      Field<"AUVIOL", 3, 1, ReadWrite, "Sticky flag indicating that an attempt was made to access parts of the address space that are marked as Secure with NS-Req for the transaction set to Non-secure">,
      Field<"INVTRAN", 4, 1, ReadWrite, "Sticky flag indicating that an exception was raised due to a branch that was not flagged as being domain crossing causing a transition from Secure to Non-secure memory">,
      Field<"LSPERR", 5, 1, ReadWrite, "Sticky flag indicating that an Security Attribution Unit (SAU) or Implementation Defined Attribution Unit (IDAU) violation occurred during the lazy preservation of floating-point state">,
      Field<"SFARVALID", 6, 1, ReadWrite, "This bit is set when the SFAR register contains a valid value">,
      Field<"LSERR", 7, 1, ReadWrite, "Sticky flag indicating that an error occurred during lazy state activation or deactivation">,
    ], "Provides information about any security related faults">,
    Register<"SFAR", 0x18, 32, [
    ], "Shows the address of the memory location that caused a security violation">,
  ];
}

class SAUInstance<string Name, int Base> : PeripheralInstance<Name, Base, SAUPeripheral>;
def SAU : SAUInstance<"SAU", 0xe000edd0>;

#endif // _PERIPHERALS_SAU_TD
