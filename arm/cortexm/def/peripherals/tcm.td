#ifndef _PERIPHERALS_TCM_TD
#define _PERIPHERALS_TCM_TD

include "base.td"

def TCMPeripheral : PeripheralType<"TCM", "TCM Security Gate"> {
  let accessWidth = 32;
  let registers = [
    Register<"ITGU_CTRL", 0x0, 32, [
      Field<"DBFEN", 0, 1, ReadWrite, "Enable data side BusFault for TGU fault">,
      Field<"DEREN", 1, 1, ReadWrite, "Enable Slave AHB (S-AHB) error response for TGU fault">,
    ], "Main TCM Gate Unit (TGU) control registers for the ITCM">,
    Register<"ITGU_CFG", 0x4, 32, [
      Field<"BLKSZ", 0, 4, ReadWrite, "TGU block size in bytes">,
      Field<"NUMBLKS", 8, 4, ReadWrite, "This value is used to calculate the number of TCM blocks">,
      Field<"PRESENT", 31, 1, ReadWrite, "Determines if the TGU is present">,
    ], "Provides access to configuration values for the ITGU">,
    RepeatingRegister<"ITGU_LUT", 16, 0x10, 32, [
    ], "Identifying the ITGU blocks as being Secure or Non-secure">,
    Register<"DTGU_CTRL", 0x100, 32, [
      Field<"DBFEN", 0, 1, ReadWrite, "Enable data side BusFault for TGU fault">,
      Field<"DEREN", 1, 1, ReadWrite, "Enable Slave AHB (S-AHB) error response for TGU fault">,
    ], "Main TCM Gate Unit (TGU) control registers for the DTCM">,
    Register<"DTGU_CFG", 0x104, 32, [
      Field<"BLKSZ", 0, 4, ReadWrite, "TGU block size in bytes">,
      Field<"NUMBLKS", 8, 4, ReadWrite, "This value is used to calculate the number of TCM blocks">,
      Field<"PRESENT", 31, 1, ReadWrite, "Determines if the TGU is present">,
    ], "Provides access to configuration values for the DTGU">,
    RepeatingRegister<"DTGU_LUT", 16, 0x110, 32, [
    ], "Identifying the DTGU blocks as being Secure or Non-secure">,
  ];
}

class TCMInstance<string Name, int Base> : PeripheralInstance<Name, Base, TCMPeripheral>;
def TCM : TCMInstance<"TCM", 0xe001e500>;

#endif // _PERIPHERALS_TCM_TD
