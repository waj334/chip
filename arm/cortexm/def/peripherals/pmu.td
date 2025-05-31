#ifndef _PERIPHERALS_PMU_TD
#define _PERIPHERALS_PMU_TD

include "base.td"

def PMUPeripheral : PeripheralType<"PMU", "Performance Monitoring Unit Extension"> {
  let accessWidth = 32;
  let registers = [
    RepeatingRegister<"EVCNTR", 8, 0x0, 32, [
      Field<"COUNTER", 0, 16, ReadWrite, "The counter counts whenever the selected event occurs, and either of the following conditions are met: SecureNoninvasiveDebugAllowed()==TRUE; The source of the event is Non-secure and NoninvasiveDebugAllowed()==TRUE.">,
    ], "Holds performance counter, which count events">,
    Register<"CCNTR", 0x7c, 32, [
    ], "Holds the value of the cycle counter, which counts processor clock cycles">,
    RepeatingRegister<"EVTYPER", 8, 0x400, 32, [
      Field<"EVTCOUNT", 0, 16, ReadWrite, "The event number of the event that is counted by event counter PMU_EVCNTR0-7">,
    ], "Configures the event counter">,
    Register<"CCFILTR", 0x47c, 32, [
    ], "This register is RES0">,
    Register<"CNTENSET", 0xc00, 32, [
      Field<"PN", 0, 8, ReadWrite, "Event counter PMU_EVCNTR0-7 enable bit">,
      Field<"C", 8, 24, ReadWrite, "CCNTR enable bit">,
    ], "Enables the PMU_CCNTR register and PMU_EVCNTR0-7 registers">,
    Register<"CNTENCLR", 0xc20, 32, [
      Field<"PN", 0, 8, ReadWrite, "Event counter PMU_EVCNTR0-7 enable bit">,
      Field<"C", 8, 24, ReadWrite, "CCNTR disable bit">,
    ], "Disables the PMU_CCNTR register and PMU_EVCNTR0-7 registers">,
    Register<"INTENSET", 0xc40, 32, [
      Field<"PN", 0, 8, ReadWrite, "Event counter overflow interrupt request disable bit for PMU_EVCNTR0-7">,
      Field<"C", 8, 24, ReadWrite, "CCNTR interrupt request enable bit">,
    ], "Enables the generation of interrupt requests on overflows from the PMU_CCNTR, and the event counter, PMU_EVCNTR.">,
    Register<"INTENCLR", 0xc60, 32, [
      Field<"PN", 0, 8, ReadWrite, "Event counter overflow interrupt request enable bit for PMU_EVCNTRn. Enable the overflow interrupt for PMU_EVCNTR0-7">,
      Field<"C", 8, 24, ReadWrite, "CCNTR overflow interrupt request disable bit">,
    ], "Disables the generation of interrupt requests on overflows from the PMU_CCNTR, and the event counters, PMU_EVCNTR">,
    Register<"OVSCLR", 0xc80, 32, [
      Field<"PN", 0, 31, ReadWrite, "Event counter overflow clear bit for PMU_EVCNTRn. Clears the PMU_EVCNTRn overflow bit">,
      Field<"C", 31, 1, ReadWrite, "CCNTR overflow bit. Clears the PMU_CCNTR overflow bit.">,
    ], "Contains the state of the overflow bit for the PMU_CCNTR, and each of the implemented event counters, PMU_EVCNTRn. Writing to this register clears these bits.">,
    Register<"SWINC", 0xca0, 32, [
      Field<"PN", 0, 8, ReadWrite, "Event counter software increment bit for PMU_EVCNTRn. An event counter n, configured for SW_INCR events, increments on every write to bit n of this field.">,
    ], "Increments a counter that is configured to count the software increment event, event 0x00">,
    Register<"OVSSET", 0xcc0, 32, [
      Field<"PN", 0, 31, ReadWrite, "Event counter overflow set bit for PMU_EVCNTRn. Sets the overflow status for PMU_EVCNTRn.">,
      Field<"C", 31, 1, ReadWrite, "CCNTR overflow bit. Sets the overflow status for PMU_CCNTR.">,
    ], "Performance Monitoring Unit Overflow Flag Status Set Register">,
    Register<"TYPE", 0xe00, 32, [
      Field<"N", 0, 8, ReadWrite, "Number of counters implemented in addition to the cycle counter, PMU_CCNTR. This field is set to 0b00001000, indicating that 8 16-bit event counters are present in addition to PMU_CCNTR.">,
      Field<"SIZE", 8, 6, ReadWrite, "Size of counters. This field determines the spacing of counters in the memory-map. This field is 0b011111.">,
      Field<"CC", 14, 1, ReadWrite, "A dedicated cycle counter is present. This bit is 0b1.">,
      Field<"FZO", 21, 1, ReadWrite, "Freeze-on-overflow is supported. This bit is 0b1.">,
      Field<"TRO", 23, 1, ReadWrite, "Trace-on-overflow not supported. This bit is 0b0.">,
    ], "Contains information regarding what the PMU supports">,
    Register<"CTRL", 0xe04, 32, [
      Field<"E", 0, 1, ReadWrite, "Enable. Enable the event counters.">,
      Field<"P", 1, 1, ReadWrite, "Event counter reset">,
      Field<"C", 2, 1, ReadWrite, "Cycle counter reset. Reset the PMU_CCNTR counter.">,
      Field<"DP", 5, 1, ReadWrite, "Disable cycle counter when event counting is prohibited">,
      Field<"FZO", 9, 1, ReadWrite, "Freeze-on-overflow. Stops events being counted once PMU_OVSCLR or PMU_OVSSET is non-zero.">,
      Field<"TRO", 11, 1, ReadWrite, "Trace-on-overflow not supported in Cortex-M55. Therefore, this bit is 0b0.">,
    ], "Configures and controls the PMU">,
    Register<"AUTHSTATUS", 0xfb8, 32, [
      Field<"NSID", 0, 2, ReadWrite, "Non-secure Invasive Debug. Indicates whether Non-secure invasive debug is allowed.">,
      Field<"NSNID", 2, 2, ReadWrite, "Non-secure Non-invasive Debug. Indicates whether Non-secure Non-invasive debug is allowed.">,
      Field<"SID", 4, 2, ReadWrite, "Secure Invasive Debug. Indicates whether Secure invasive debug is implemented and allowed.">,
      Field<"SNID", 6, 2, ReadWrite, "Secure Non-invasive Debug. Indicates whether Secure Non-invasive debug is implemented and allowed.">,
      Field<"NSUID", 16, 2, ReadWrite, "Non-secure Unprivileged Invasive Debug Allowed. Indicates that Unprivileged Halting Debug is allowed for the Non-secure state.">,
      Field<"NSUNID", 18, 2, ReadWrite, "Non-secure Unprivileged Non-invasive Debug Allowed. Indicates that Unprivileged Non-invasive Debug is allowed for the Non-secure state.">,
      Field<"SUID", 20, 2, ReadWrite, "Secure Unprivileged Invasive Debug Allowed. Indicates that Unprivileged Invasive Debug is allowed for the Secure state.">,
      Field<"SUNID", 22, 2, ReadWrite, "Secure Unprivileged Non-invasive Debug Allowed. Indicates that Unprivileged Non-invasive Debug is allowed for the Secure state.">,
    ], "Provides information about the state of the implementation defined authentication interface for the PMU">,
    Register<"DEVARCH", 0xfbc, 32, [
      Field<"ARCHID", 0, 16, ReadWrite, "Defines this part to be an Armv8‑M debug component. For the PMU, bits [15:12] are 0x0, bits [11:0] are 0xA06.">,
      Field<"REVISION", 16, 4, ReadWrite, "Defines the architecture revision. For the PMU, the revision defined by Armv8.1‑M is 0x2.">,
      Field<"PRESENT", 20, 1, ReadWrite, "Determines the presence of DEVARCH. When set to 1, indicates that the DEVARCH is present. This bit reads as 0x1.">,
      Field<"ARCHITECT", 21, 11, ReadWrite, "Defines the architecture of the component. For the PMU, it is Arm Limited. Bits[31:28] are the JEP 106 continuation code, 0x4, and bits [27:21] are the JEP 106 ID code, 0x3B.">,
    ], "Identifies the programmers model architecture of the PMU">,
    Register<"DEVTYPE", 0xfcc, 32, [
      Field<"MAJOR", 0, 4, ReadWrite, "Major type. This field reads as 0x6.">,
      Field<"SUB", 4, 4, ReadWrite, "Subtype. This field reads as 0x1.">,
    ], "Indicates to a debugger that the PMU is a part of the processor interface">,
    Register<"PIDR4", 0xfd0, 32, [
      Field<"DES_2", 0, 4, ReadWrite, "Designer, JEP 106 continuation code, least significant nibble. For Arm Limited, this field is 0x4.">,
      Field<"SIZE", 4, 4, ReadWrite, "Size of the component. This field is RAZ.">,
    ], "Provides information to identify the PMU.">,
    Register<"PIDR0", 0xfe0, 32, [
      Field<"PART_0", 4, 4, ReadWrite, "Part number, least significant byte. This field reads 0x22.">,
    ], "Provides information to identify the PMU">,
    Register<"PIDR1", 0xfe4, 32, [
      Field<"PART_1", 0, 4, ReadWrite, "Part number, most significant nibble. This field is 0xD.">,
      Field<"DES_0", 4, 4, ReadWrite, "Designer, least significant nibble of JEP 106 ID code. For Arm Limited, this field 0xB.">,
    ], "Provides information to identify the PMU">,
    Register<"PIDR2", 0xfe8, 32, [
      Field<"DES_1", 0, 3, ReadWrite, "Designer, most significant bits of JEP 106 ID code. For Arm Limited, this field reas as 0b011.">,
      Field<"JEDEC", 3, 1, ReadWrite, "Indicates a JEP10 identity code. This is RAO.">,
      Field<"REVISION", 4, 4, ReadWrite, "Part major revision. Parts can also use this field to extend Part number to 16-bits. This field reads as 0x0000.">,
    ], "Provides information to identify the PMU.">,
    Register<"PIDR3", 0xfec, 32, [
      Field<"CMOD", 0, 3, ReadWrite, "Customer modified. This field is 0x0000.">,
      Field<"REVAND", 4, 4, ReadWrite, "Part minor revision. This field is 0x0000.">,
    ], "Provides information to identify the PMU">,
    Register<"CIDR0", 0xff0, 32, [
      Field<"PRMBL_0", 0, 8, ReadWrite, "Preamble. This field reads 0x0D.">,
    ], "Provides information to identify a PMU component">,
    Register<"CIDR1", 0xff4, 32, [
      Field<"PRMBL_1", 0, 4, ReadWrite, "Preamble. This field reads 0x0.">,
      Field<"CLASS", 4, 4, ReadWrite, "Component class. This field reads 0x0.">,
    ], "Provides information to identify a PMU component">,
    Register<"CIDR2", 0xff8, 32, [
      Field<"PRMBL_2", 0, 4, ReadWrite, "Preamble. This field reads 0x05.">,
    ], "Provides information to identify a PMU component">,
    Register<"CIDR3", 0xffc, 32, [
      Field<"PRMBL_3", 0, 4, ReadWrite, "Preamble. This field reads 0xB1.">,
    ], "Provides information to identify a PMU component">,
  ];
}

class PMUInstance<string Name, int Base> : PeripheralInstance<Name, Base, PMUPeripheral>;
def PMU : PMUInstance<"PMU", 0xe0003000>;

#endif // _PERIPHERALS_PMU_TD
