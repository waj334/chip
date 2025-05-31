#ifndef _PERIPHERALS_EWICCS_TD
#define _PERIPHERALS_EWICCS_TD

include "base.td"

def EWICCSPeripheral : PeripheralType<"EWICCS", "External Wakeup Interrupt Controller CoreSight registers"> {
  let accessWidth = 32;
  let registers = [
    Register<"ITCTRL", 0x0, 32, [
    ], "Used to dynamically switch between functional mode and integration mode. The EWIC does not support integration mode, and this register is RAZ.">,
    Register<"CLAIMSET", 0xa0, 32, [
    ], "The EWIC_CLAIMSET register is used to set whether functionality is in use by a debug agent. The EWIC does not have any associated debug functionality, and this register is 0xF.">,
    Register<"CLAIMCLR", 0xa4, 32, [
    ], "The EWIC_CLAIMCLR register is used to set whether functionality is in use by a debug agent. The EWIC does not have any associated debug functionality, and this register is 0x0.">,
    Register<"DEVAFF0", 0xa8, 32, [
    ], "The EWIC_DEVAFF0 register enables a debugger to determine whether the EWIC and the processor have an affinity with each other. The EWIC does not have any associated debug functionality, and this register is 0x80000000.">,
    Register<"DEVAFF1", 0xac, 32, [
    ], "The EWIC_DEVAFF1 register enables a debugger to determine whether the EWIC and the processor have an affinity with each other. The EWIC does not have any associated debug functionality, and this register is 0x0.">,
    Register<"LAR", 0xb0, 32, [
    ], "The EWIC_LAR register controls software access to CoreSight components to reduce the likelihood of accidental access to the EWIC. The EWIC does not support software locking, and writing to this register has no affect.">,
    Register<"LSR", 0xb4, 32, [
    ], "The EWIC_LSR register controls software access to CoreSight components to reduce the likelihood of accidental access to the EWIC. The EWIC does not support software locking, and this register is RAZ.">,
    Register<"AUTHSTATUS", 0xb8, 32, [
    ], "The EWIC_AUTHSTATUS register indicates whether the EWIC includes debug functionality. The EWIC is not a debug component, therefore, the EWIC_AUTHSTATIS register is RAZ.">,
    Register<"DEVARCH", 0xbc, 32, [
      Field<"ARCHID", 0, 16, ReadWrite, "Returns a value that identifies the architecture of the component. This value is 0x0A07 corresponding to EWIC architecture.">,
      Field<"REVISION", 16, 4, ReadWrite, "Architecture revision. This value is 0b0000.">,
      Field<"PRESENT", 20, 1, ReadWrite, "Indicates the presence of this register. This value is 0b1.">,
      Field<"ARCHITECT_JEP106I", 21, 7, ReadWrite, "Defines the architect of the component, JEP106 identification code">,
      Field<"ARCHITECT_JEP106C", 28, 4, ReadWrite, "Defines the architect of the component, JEP106 continuation code">,
    ], "Identifies the architecture of the EWIC">,
    Register<"DEVID2", 0xc0, 32, [
    ], "The EWIC_DEVID2 indicates some capabilities of the EWIC. This register is RAZ/WI.">,
    Register<"DEVID1", 0xc4, 32, [
    ], "The EWIC_DEVID1 indicates some capabilities of the EWIC. This register is RAZ/WI.">,
    Register<"DEVID", 0xc8, 32, [
    ], "The EWIC_DEVID indicates some capabilities of the EWIC. This register is RAZ/WI.">,
    Register<"DEVTYPE", 0xcc, 32, [
      Field<"MAJOR", 0, 4, ReadWrite, "Major type for the component device type. This field is set to 0b0000.">,
      Field<"SUB", 4, 4, ReadWrite, "Sub type for the component device type. This field is set to 0b0000.">,
    ], "Provides part number information about the EWIC component to the debugger.">,
    Register<"PIDR4", 0xd0, 32, [
      Field<"DES_2", 0, 4, ReadWrite, "JEP106 continuation code. Together with EWIC_PIDR2.DES_1 and EWIC_PIDR1.DES_0, they indicate the designer of the component, not the implementer, except where the two are the same. The reset value of this field is 0x4.">,
      Field<"SIZE", 4, 4, ReadWrite, "This field indicates the memory size that the EWIC uses. This field returns 0x0 indicating that the component uses an UNKNOWN number of 4KB blocks.">,
    ], "provides information about the memory size and JEP106 continuation code that the External Wakeup Interrupt Controller (EWIC) component uses">,
    Register<"PIDR5", 0xd4, 32, [
    ], "The register is reserved, RES0">,
    Register<"PIDR6", 0xd8, 32, [
    ], "The register is reserved, RES0">,
    Register<"PIDR7", 0xdc, 32, [
    ], "The register is reserved, RES0">,
    Register<"PIDR0", 0xe0, 32, [
      Field<"PART_0", 0, 8, ReadWrite, "This field indicates the part number. When taken together with EWIC_PIDR1.PART_1, it indicates the component. The part number is selected by the designer of the component.">,
    ], "Indicates the EWIC component part number">,
    Register<"PIDR1", 0xe4, 32, [
      Field<"PART_1", 0, 4, ReadWrite, "This field indicates the part number, bits[11:8]. Taken together with EWIC_PIDR0.PART_0 it indicates the component. The part number is selected by the designer of the component. The reset value is 0xD.">,
      Field<"DES_0", 4, 4, ReadWrite, "This field indicates the JEP106 identification code, bits[3:0]. Together, with EWIC_PIDR4.DES_2 and EWIC_PIDR2.DES_1, they indicate the designer of the component and not the implementer, except where the two are the same. The reset value is 0xB.">,
    ], "Indicates the EWIC component JEP106 continuation code and part number">,
    Register<"PIDR2", 0xe8, 32, [
      Field<"DES_1", 0, 3, ReadWrite, "This field is the JEP106 identification code, bits[6:4]. Together, with CTI_PIDR4.DES_2 and CTI_PIDR1.DES_0, they indicate the designer of the component and not the implementer, except where the two are the same. The reset value is 0b011.">,
      Field<"JEDEC", 3, 1, ReadWrite, "This field is always 1, indicating that a JEDEC assigned value is used.">,
      Field<"REVISION", 4, 4, ReadWrite, "This field indicates the revision number of the EWIC component. It is an incremental value starting at 0x0 for the first design. The reset value is 0x0.">,
    ], "Indicates the EWIC component revision number, JEDEC value, and part of the JEP106 continuation code">,
    Register<"PIDR3", 0xec, 32, [
      Field<"CMOD", 0, 4, ReadWrite, "Customer modified. Where the component is reusable IP, this value indicates whether you have modified the behavior of the component. In most cases, this field is 0x0.">,
      Field<"REVAND", 4, 4, ReadWrite, "This field indicates minor errata fixes specific to this design, for example metal fixes after implementation. In most cases this field is 0x0.">,
    ], "Indicates minor errata fixes of the Cross Trigger Interface (CTI) component and if you have modified the behavior of the component">,
    Register<"CIDR0", 0xf0, 32, [
      Field<"PRMBL_0", 0, 8, ReadWrite, "Preamble. This field returns 0x0D.">,
    ], "Indicates the preamble">,
    Register<"CIDR1", 0xf4, 32, [
      Field<"CLASS", 0, 8, ReadWrite, "Component class. Returns 0x9, indicating this is a CoreSight component.">,
      Field<"PRMBL_1", 0, 8, ReadWrite, "Preamble. This field returns 0x0.">,
    ], "indicates the component class and preamble">,
    Register<"CIDR2", 0xf8, 32, [
      Field<"PRMBL_2", 0, 8, ReadWrite, "Preamble. This field returns 0x05.">,
    ], "Indicates the preamble">,
    Register<"CIDR3", 0xfc, 32, [
      Field<"PRMBL_3", 0, 8, ReadWrite, "Preamble. This field returns 0xB1.">,
    ], "Indicates the preamble">,
  ];
}

class EWICCSInstance<string Name, int Base> : PeripheralInstance<Name, Base, EWICCSPeripheral>;
def EWICCS : EWICCSInstance<"EWICCS", 0xe0047f00>;

#endif // _PERIPHERALS_EWICCS_TD
