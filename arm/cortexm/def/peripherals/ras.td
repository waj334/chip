#ifndef _PERIPHERALS_RAS_TD
#define _PERIPHERALS_RAS_TD

include "base.td"

def RASPeripheral : PeripheralType<"RAS", "Reliability, Availability, and Serviceability (RAS) Extension"> {
  let accessWidth = 32;
  let registers = [
    Register<"ERRFR0", 0x0, 32, [
      Field<"ED", 0, 2, ReadWrite, "Error reporting and logging">,
      Field<"UE", 8, 2, ReadWrite, "Enable Uncorrected error (UE) reporting as an external abort">,
    ], "Describes the RAS features that are supported">,
    Register<"ERRCTRL0", 0x8, 32, [
    ], "This register is RES0">,
    Register<"ERRSTATUS0", 0x10, 32, [
      Field<"SERR", 0, 8, ReadWrite, "Architecturally-defined primary error code">,
      Field<"UET", 20, 2, ReadWrite, "Uncorrectable error type">,
      Field<"POISON", 22, 1, ReadWrite, "BusFault due to RPOISON or TEBRn.POISON">,
      Field<"DE", 23, 1, ReadWrite, "Deferred errors">,
      Field<"CE", 24, 2, ReadWrite, "Corrected errors">,
      Field<"MV", 26, 1, ReadWrite, "Miscellaneous registers valid">,
      Field<"OF", 27, 1, ReadWrite, "RAS events occurred since the last time ERRSTATUS0.V was cleared">,
      Field<"ER", 28, 1, ReadWrite, "BusFault caused by RAS event">,
      Field<"UE", 29, 1, ReadWrite, "Uncorrected errors (UEs)">,
      Field<"V", 30, 1, ReadWrite, "Status valid">,
      Field<"AV", 31, 1, ReadWrite, "Address valid">,
    ], "Contains information about the Reliability, Availability, and Serviceability (RAS) event that is currently logged in record 0">,
    Register<"ERRADDR0", 0x18, 32, [
    ], "Contains information about the address of the Reliability, Availability, and Serviceability (RAS) event in record 0">,
    Register<"ERRADDR20", 0x1c, 32, [
      Field<"AI", 29, 1, ReadWrite, "Address incorrect">,
      Field<"SI", 30, 1, ReadWrite, "Security information incorrect">,
    ], "Contains information about the address of the Reliability, Availability, and Serviceability (RAS) event in record 0">,
    Register<"ERRMISC00", 0x20, 32, [
    ], "This register is RES0">,
    Register<"ERRMISC10", 0x24, 32, [
      Field<"TYPE", 0, 2, ReadWrite, "type of Reliability, Availability, and Serviceability (RAS) event logged">,
    ], "Error syndrome register for the event in record 0">,
    Register<"ERRMISC20", 0x28, 32, [
    ], "This register is RES0">,
    Register<"ERRMISC30", 0x2c, 32, [
    ], "This register is RES0">,
    Register<"ERRMISC40", 0x30, 32, [
    ], "This register is RES0">,
    Register<"ERRMISC50", 0x34, 32, [
    ], "This register is RES0">,
    Register<"ERRMISC60", 0x38, 32, [
    ], "This register is RES0">,
    Register<"ERRMISC70", 0x3c, 32, [
    ], "This register is RES0">,
    Register<"ERRGSR0", 0xe00, 32, [
      Field<"ERR0", 0, 1, ReadWrite, "Error record 0 is valid">,
    ], "Summarizes the valid error records">,
    Register<"ERRDEVID", 0xfc8, 32, [
      Field<"NUM", 0, 16, ReadWrite, "Maximum Error Record Index+1">,
    ], "Contains the number of error records that an implementation supports">,
    Register<"RFSR", 0x9f04, 32, [
      Field<"UET", 0, 2, ReadWrite, "Error type">,
      Field<"IS", 16, 15, ReadWrite, "Indicates the type of RAS exception that has occurred">,
      Field<"VALID", 31, 1, ReadWrite, "Indicates whether the register is valid">,
    ], "Reports the fault status of Reliability, Availability, and Serviceability (RAS) related faults from Error Correcting Code (ECC) errors that are detected in the Level 1 (L1) instruction cache, data cache, and TCM">,
  ];
}

class RASInstance<string Name, int Base> : PeripheralInstance<Name, Base, RASPeripheral>;
def RAS : RASInstance<"RAS", 0xe0005000>;

#endif // _PERIPHERALS_RAS_TD
