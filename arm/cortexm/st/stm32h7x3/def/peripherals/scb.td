#ifndef _PERIPHERALS_SCB_TD
#define _PERIPHERALS_SCB_TD

include "base.td"

def SCBPeripheral : PeripheralType<"SCB", "System control block"> {
  let accessWidth = 32;
  let registers = [
    Register<"CPUID", 0x0, 32, [
      Field<"Revision", 0, 4, ReadWrite, "Revision number">,
      Field<"PartNo", 4, 12, ReadWrite, "Part number of the processor">,
      Field<"Constant", 16, 4, ReadWrite, "Reads as 0xF">,
      Field<"Variant", 20, 4, ReadWrite, "Variant number">,
      Field<"Implementer", 24, 8, ReadWrite, "Implementer code">,
    ], "CPUID base register">,
    Register<"ICSR", 0x4, 32, [
      Field<"VECTACTIVE", 0, 9, ReadWrite, "Active vector">,
      Field<"RETTOBASE", 11, 1, ReadWrite, "Return to base level">,
      Field<"VECTPENDING", 12, 7, ReadWrite, "Pending vector">,
      Field<"ISRPENDING", 22, 1, ReadWrite, "Interrupt pending flag">,
      Field<"PENDSTCLR", 25, 1, ReadWrite, "SysTick exception clear-pending bit">,
      Field<"PENDSTSET", 26, 1, ReadWrite, "SysTick exception set-pending bit">,
      Field<"PENDSVCLR", 27, 1, ReadWrite, "PendSV clear-pending bit">,
      Field<"PENDSVSET", 28, 1, ReadWrite, "PendSV set-pending bit">,
      Field<"NMIPENDSET", 31, 1, ReadWrite, "NMI set-pending bit.">,
    ], "Interrupt control and state register">,
    Register<"VTOR", 0x8, 32, [
      Field<"TBLOFF", 9, 21, ReadWrite, "Vector table base offset field">,
    ], "Vector table offset register">,
    Register<"AIRCR", 0xc, 32, [
      Field<"VECTRESET", 0, 1, ReadWrite, "VECTRESET">,
      Field<"VECTCLRACTIVE", 1, 1, ReadWrite, "VECTCLRACTIVE">,
      Field<"SYSRESETREQ", 2, 1, ReadWrite, "SYSRESETREQ">,
      Field<"PRIGROUP", 8, 3, ReadWrite, "PRIGROUP">,
      Field<"ENDIANESS", 15, 1, ReadWrite, "ENDIANESS">,
      Field<"VECTKEYSTAT", 16, 16, ReadWrite, "Register key">,
    ], "Application interrupt and reset control register">,
    Register<"SCR", 0x10, 32, [
      Field<"SLEEPONEXIT", 1, 1, ReadWrite, "SLEEPONEXIT">,
      Field<"SLEEPDEEP", 2, 1, ReadWrite, "SLEEPDEEP">,
      Field<"SEVEONPEND", 4, 1, ReadWrite, "Send Event on Pending bit">,
    ], "System control register">,
    Register<"CCR", 0x14, 32, [
      Field<"NONBASETHRDENA", 0, 1, ReadWrite, "Configures how the processor enters Thread mode">,
      Field<"USERSETMPEND", 1, 1, ReadWrite, "USERSETMPEND">,
      Field<"UNALIGN__TRP", 3, 1, ReadWrite, "UNALIGN_ TRP">,
      Field<"DIV_0_TRP", 4, 1, ReadWrite, "DIV_0_TRP">,
      Field<"BFHFNMIGN", 8, 1, ReadWrite, "BFHFNMIGN">,
      Field<"STKALIGN", 9, 1, ReadWrite, "STKALIGN">,
      Field<"DC", 16, 1, ReadWrite, "DC">,
      Field<"IC", 17, 1, ReadWrite, "IC">,
      Field<"BP", 18, 1, ReadWrite, "BP">,
    ], "Configuration and control register">,
    Register<"SHPR1", 0x18, 32, [
      Field<"PRI_4", 0, 8, ReadWrite, "Priority of system handler 4">,
      Field<"PRI_5", 8, 8, ReadWrite, "Priority of system handler 5">,
      Field<"PRI_6", 16, 8, ReadWrite, "Priority of system handler 6">,
    ], "System handler priority registers">,
    Register<"SHPR2", 0x1c, 32, [
      Field<"PRI_11", 24, 8, ReadWrite, "Priority of system handler 11">,
    ], "System handler priority registers">,
    Register<"SHPR3", 0x20, 32, [
      Field<"PRI_14", 16, 8, ReadWrite, "Priority of system handler 14">,
      Field<"PRI_15", 24, 8, ReadWrite, "Priority of system handler 15">,
    ], "System handler priority registers">,
    Register<"SHCRS", 0x24, 32, [
      Field<"MEMFAULTACT", 0, 1, ReadWrite, "Memory management fault exception active bit">,
      Field<"BUSFAULTACT", 1, 1, ReadWrite, "Bus fault exception active bit">,
      Field<"USGFAULTACT", 3, 1, ReadWrite, "Usage fault exception active bit">,
      Field<"SVCALLACT", 7, 1, ReadWrite, "SVC call active bit">,
      Field<"MONITORACT", 8, 1, ReadWrite, "Debug monitor active bit">,
      Field<"PENDSVACT", 10, 1, ReadWrite, "PendSV exception active bit">,
      Field<"SYSTICKACT", 11, 1, ReadWrite, "SysTick exception active bit">,
      Field<"USGFAULTPENDED", 12, 1, ReadWrite, "Usage fault exception pending bit">,
      Field<"MEMFAULTPENDED", 13, 1, ReadWrite, "Memory management fault exception pending bit">,
      Field<"BUSFAULTPENDED", 14, 1, ReadWrite, "Bus fault exception pending bit">,
      Field<"SVCALLPENDED", 15, 1, ReadWrite, "SVC call pending bit">,
      Field<"MEMFAULTENA", 16, 1, ReadWrite, "Memory management fault enable bit">,
      Field<"BUSFAULTENA", 17, 1, ReadWrite, "Bus fault enable bit">,
      Field<"USGFAULTENA", 18, 1, ReadWrite, "Usage fault enable bit">,
    ], "System handler control and state register">,
    Register<"CFSR_UFSR_BFSR_MMFSR", 0x28, 32, [
      Field<"IACCVIOL", 0, 1, ReadWrite, "IACCVIOL">,
      Field<"DACCVIOL", 1, 1, ReadWrite, "DACCVIOL">,
      Field<"MUNSTKERR", 3, 1, ReadWrite, "MUNSTKERR">,
      Field<"MSTKERR", 4, 1, ReadWrite, "MSTKERR">,
      Field<"MLSPERR", 5, 1, ReadWrite, "MLSPERR">,
      Field<"MMARVALID", 7, 1, ReadWrite, "MMARVALID">,
      Field<"IBUSERR", 8, 1, ReadWrite, "Instruction bus error">,
      Field<"PRECISERR", 9, 1, ReadWrite, "Precise data bus error">,
      Field<"IMPRECISERR", 10, 1, ReadWrite, "Imprecise data bus error">,
      Field<"UNSTKERR", 11, 1, ReadWrite, "Bus fault on unstacking for a return from exception">,
      Field<"STKERR", 12, 1, ReadWrite, "Bus fault on stacking for exception entry">,
      Field<"LSPERR", 13, 1, ReadWrite, "Bus fault on floating-point lazy state preservation">,
      Field<"BFARVALID", 15, 1, ReadWrite, "Bus Fault Address Register (BFAR) valid flag">,
      Field<"UNDEFINSTR", 16, 1, ReadWrite, "Undefined instruction usage fault">,
      Field<"INVSTATE", 17, 1, ReadWrite, "Invalid state usage fault">,
      Field<"INVPC", 18, 1, ReadWrite, "Invalid PC load usage fault">,
      Field<"NOCP", 19, 1, ReadWrite, "No coprocessor usage fault.">,
      Field<"UNALIGNED", 24, 1, ReadWrite, "Unaligned access usage fault">,
      Field<"DIVBYZERO", 25, 1, ReadWrite, "Divide by zero usage fault">,
    ], "Configurable fault status register">,
    Register<"HFSR", 0x2c, 32, [
      Field<"VECTTBL", 1, 1, ReadWrite, "Vector table hard fault">,
      Field<"FORCED", 30, 1, ReadWrite, "Forced hard fault">,
      Field<"DEBUG_VT", 31, 1, ReadWrite, "Reserved for Debug use">,
    ], "Hard fault status register">,
    Register<"MMFAR", 0x34, 32, [
      Field<"ADDRESS", 0, 32, ReadWrite, "Memory management fault address">,
    ], "Memory management fault address register">,
    Register<"BFAR", 0x38, 32, [
      Field<"ADDRESS", 0, 32, ReadWrite, "Bus fault address">,
    ], "Bus fault address register">,
  ];
}

class SCBInstance<string Name, int Base> : PeripheralInstance<Name, Base, SCBPeripheral>;
def SCB : SCBInstance<"SCB", 0xe000ed00>;

#endif // _PERIPHERALS_SCB_TD
