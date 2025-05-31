#ifndef _PERIPHERALS_SCB_TD
#define _PERIPHERALS_SCB_TD

include "base.td"

def SCBPeripheral : PeripheralType<"SCB", "System Control Block"> {
  let accessWidth = 32;
  let registers = [
    Register<"REVIDR", 0x0, 32, [
      Field<"IMPLEMENTATION_SPECIFIC", 0, 32, ReadWrite, "Implementation-specific minor revision information that can be interpreted with the CPUID register.">,
    ], "Provides additional implementation-specific minor revision that can be interpreted with the CPUID register">,
    Register<"CPUID", 0x4, 32, [
      Field<"REVISION", 0, 4, ReadWrite, "Revision number, the m value in the rnpm product revision identifier">,
      Field<"PARTNO", 4, 12, ReadWrite, "Part number of the processor">,
      Field<"ARCHITECTURE", 16, 4, ReadWrite, "Reads as 0xF">,
      Field<"VARIANT", 20, 4, ReadWrite, "Variant number, the n value in the rnpm product revision identifier">,
      Field<"IMPLEMENTER", 24, 8, ReadWrite, "Implementer code">,
    ], "Contains the processor part number, version, and implementation information">,
    Register<"ICSR", 0x8, 32, [
      Field<"VECTACTIVE", 0, 9, ReadWrite, "Contains the active exception number">,
      Field<"RETTOBASE", 11, 1, ReadWrite, "Indicates whether there are pre-empted active exceptions">,
      Field<"VECTPENDING", 12, 9, ReadWrite, "Indicates the exception number of the highest priority pending enabled exception">,
      Field<"ISRPENDING", 22, 1, ReadWrite, "Interrupt pending flag, excluding NMI and Faults">,
      Field<"ISRPREEMPT", 23, 1, ReadWrite, "Indicates whether a pending exception is handled on exit from debug state.">,
      Field<"PENDSTCLR", 25, 1, ReadWrite, "SysTick exception clear-pending bit">,
      Field<"PENDSTSET", 26, 1, ReadWrite, "SysTick exception set-pending bit">,
      Field<"PENDSVCLR", 27, 1, ReadWrite, "PendSV clear-pending bit (write-only)">,
      Field<"PENDSVSET", 28, 1, ReadWrite, "PendSV set-pending bit">,
      Field<"PENDNMICLR", 30, 1, ReadWrite, "Pend NMI clear bit (write-only)">,
      Field<"PENDNMISET", 31, 1, ReadWrite, "On writes, makes the NMI exception pending">,
    ], "Provides a set-pending bit for the non-maskable interrupt exception, and set-pending and clear-pending bits for the PendSV and SysTick exceptions">,
    Register<"VTOR", 0xc, 32, [
      Field<"TBLOFF", 7, 25, ReadWrite, "Bits[31:7] of the vector table address">,
    ], "Indicates the offset of the vector table base address from memory address 0x00000000">,
    Register<"AIRCR", 0x10, 32, [
      Field<"VECTCLRACTIVE", 1, 1, ReadWrite, "Reserved for Debug use. This bit reads as 0. When writing to the register you must write 0 to this bit, otherwise behavior is UNPREDICTABLE">,
      Field<"SYSRESETREQ", 2, 1, ReadWrite, "System Reset Request">,
      Field<"SYSRESETREQS", 3, 1, ReadWrite, "System reset request, Secure state only">,
      Field<"PRIGROUP", 8, 3, ReadWrite, "Interrupt priority grouping">,
      Field<"BFHFNMINS", 13, 1, ReadWrite, "BusFault, HardFault, and NMI Non-secure enable">,
      Field<"PRIS", 14, 1, ReadWrite, "Prioritize Secure exceptions">,
      Field<"ENDIANNESS", 15, 1, ReadWrite, "Data endianness">,
      Field<"VECTKEY", 16, 16, ReadWrite, "Register key. Reads as 0x0FA05, On writes, write 0x5FA to VECTKEY, otherwise the write is ignored.">,
    ], "Sets or returns interrupt control data">,
    Register<"SCR", 0x14, 32, [
      Field<"SLEEPONEXIT", 1, 1, ReadWrite, "Indicates sleep-on-exit when returning from Handler mode to Thread mode">,
      Field<"SLEEPDEEP", 2, 1, ReadWrite, "Indicates whether the processor uses sleep or deep sleep as its low-power mode">,
      Field<"SLEEPDEEPS", 3, 1, ReadWrite, "Controls whether the SLEEPDEEP bit is only accessible from the Secure state">,
      Field<"SEVONPEND", 4, 1, ReadWrite, "Determines whether an interrupt assigned to the same Security state as the SEVONPEND bit transitioning from inactive state to pending state generates a wakeup event">,
    ], "Controls features of entry to and exit from low-power state">,
    Register<"CCR", 0x18, 32, [
      Field<"USERSETMPEND", 1, 1, ReadWrite, "User set main pending. Determines whether unprivileged accesses are permitted to pend interrupts from the STIR.">,
      Field<"UNALIGN_TRP", 3, 1, ReadWrite, "Controls the trapping of unaligned word or halfword accesses">,
      Field<"DIV_0_TRP", 4, 1, ReadWrite, "Divide by zero trap. Controls the generation of a DIVBYZERO UsageFault when attempting to perform integer division by zero.">,
      Field<"BFHFNMIGN", 8, 1, ReadWrite, "Determines the effect of precise bus faults on handlers running at a requested priority less than 0">,
      Field<"STKOFHFNMIGN", 10, 1, ReadWrite, "Controls the effect of a stack limit violation while executing at a requested priority less than 0">,
      Field<"DC", 16, 1, ReadWrite, "Data cache enable">,
      Field<"IC", 17, 1, ReadWrite, "Instruction cache enable">,
      Field<"LOB", 19, 1, ReadWrite, "Loop and branch info cache enable">,
      Field<"TRD", 20, 1, ReadWrite, "Thread reentrancy disabled">,
    ], "indicates some aspects of the behavior of the processor">,
    Register<"SHPR1", 0x1c, 32, [
      Field<"PRI_4", 0, 8, ReadWrite, "Priority of system handler 4, MemManage">,
      Field<"PRI_5", 8, 8, ReadWrite, "Priority of system handler 5, BusFault">,
      Field<"PRI_6", 16, 8, ReadWrite, "Priority of system handler 6, UsageFault">,
      Field<"PRI_7", 24, 8, ReadWrite, "Priority of system handler 7, SecureFault">,
    ], "Sets or returns priority for system handlers 4-7">,
    Register<"SHPR2", 0x20, 32, [
      Field<"PRI_11", 24, 8, ReadWrite, "Priority of system handler 11, SVCall">,
    ], "Sets or returns priority for system handler 8-11">,
    Register<"SHPR3", 0x24, 32, [
      Field<"PRI_14", 16, 8, ReadWrite, "Priority of system handler 14, PendSV">,
      Field<"PRI_15", 24, 8, ReadWrite, "Priority of system handler 15, SysTick">,
    ], "Sets or returns priority for system handlers 12-15">,
    Register<"SHCSR", 0x28, 32, [
      Field<"MEMFAULTACT", 0, 1, ReadWrite, "MemManage exception active">,
      Field<"BUSFAULTACT", 1, 1, ReadWrite, "BusFault exception active">,
      Field<"HARDFAULTACT", 2, 1, ReadWrite, "HardFault exception active">,
      Field<"USGFAULTACT", 3, 1, ReadWrite, "UsageFault exception active">,
      Field<"SECUREFAULTACT", 4, 1, ReadWrite, "SecureFault exception active state">,
      Field<"NMIACT", 5, 1, ReadWrite, "NMI exception active">,
      Field<"SVCALLACT", 7, 1, ReadWrite, "SVCall active">,
      Field<"MONITORACT", 8, 1, ReadWrite, "Debug monitor active">,
      Field<"PENDSVACT", 10, 1, ReadWrite, "PendSV exception active">,
      Field<"SYSTICKACT", 11, 1, ReadWrite, "SysTick exception active">,
      Field<"USGFAULTPENDED", 12, 1, ReadWrite, "UsageFault exception pending">,
      Field<"MEMFAULTPENDED", 13, 1, ReadWrite, "MemManage exception pending">,
      Field<"BUSFAULTPENDED", 14, 1, ReadWrite, "BusFault pending">,
      Field<"SVCALLPENDED", 15, 1, ReadWrite, "SVCall pending">,
      Field<"MEMFAULTENA", 16, 1, ReadWrite, "MemManage enable">,
      Field<"BUSFAULTENA", 17, 1, ReadWrite, "BusFault enable">,
      Field<"USGFAULTENA", 18, 1, ReadWrite, "UsageFault enable">,
      Field<"SECUREFAULTENA", 19, 1, ReadWrite, "RES0">,
      Field<"SECUREFAULTPENDED", 20, 1, ReadWrite, "SecureFault exception pended state bit, set to 1 to allow exception modification">,
      Field<"HARDFAULTPENDED", 21, 1, ReadWrite, "HardFault exception pended state bit, set to 1 to allow exception modification">,
    ], "Enables the system handlers and indicates the pending status of the BusFault, MemManage fault, and SVC exceptions, and indicates the active status of the system handlers">,
    Register<"CFSR", 0x2c, 32, [
      Field<"MEMMANAGE", 0, 8, ReadWrite, "Provides information on MemManage exceptions">,
      Field<"BUSFAULT", 8, 8, ReadWrite, "Provides information on BusFault exceptions">,
      Field<"USAGEFAULT", 16, 16, ReadWrite, "Provides information on UsageFault exceptions">,
    ], "Indicates the cause of a MemManage fault, BusFault, or UsageFault">,
    Register<"HFSR", 0x30, 32, [
      Field<"VECTTBL", 1, 1, ReadWrite, "Indicates a HardFault on a vector table read during exception processing">,
      Field<"FORCED", 30, 1, ReadWrite, "Indicates a forced HardFault, generated by escalation of a fault with configurable priority that cannot be handled, either because of priority or because it is disabled">,
      Field<"DEBUGEVT", 31, 1, ReadWrite, "Reserved for Debug use. When writing to the register you must write 1 to this bit, otherwise behavior is unpredictable.">,
    ], "Gives information about events that activate the HardFault handler">,
    Register<"MMFAR", 0x38, 32, [
    ], "Shows the address of the memory location that caused a Memory Protection Unit (MPU) fault">,
    Register<"BFAR", 0x3c, 32, [
    ], "Shows the address associated with a precise data access BusFault">,
    Register<"AFSR", 0x40, 32, [
      Field<"IITCM", 0, 1, ReadWrite, "Imprecise fault on ITCM interface">,
      Field<"IDTCM", 1, 1, ReadWrite, "Imprecise fault on DTCM interface">,
      Field<"IPAHB", 2, 1, ReadWrite, "Imprecise fault on P-AHB interface">,
      Field<"IMAXI", 3, 1, ReadWrite, "Imprecise fault on M-AXI interface">,
      Field<"IEPPB", 4, 1, ReadWrite, "Imprecise fault on EPPB interface">,
      Field<"IMAXITYPE", 6, 1, ReadWrite, "AXI response that caused the imprecise fault. Only valid when AFSR.IMAXI is 1.">,
      Field<"IECC", 7, 1, ReadWrite, "Imprecise fault caused by uncorrectable ECC error">,
      Field<"IPOISON", 9, 1, ReadWrite, "Imprecise BusFault because of RPOISON">,
      Field<"PITCM", 10, 1, ReadWrite, "Precise fault on ITCM interface">,
      Field<"PDTCM", 11, 1, ReadWrite, "Precise fault on DTCM interface">,
      Field<"PPAHB", 12, 1, ReadWrite, "Precise fault on Peripheral AHB (P-AHB) interface">,
      Field<"PMAXI", 13, 1, ReadWrite, "Precise fault on M-AXI interface">,
      Field<"PEPPB", 14, 1, ReadWrite, "Precise fault on External Private Peripheral Bus (EPPB) interface">,
      Field<"PIPPB", 15, 1, ReadWrite, "Precise fault on Internal Private Peripheral Bus (EPPB) interface">,
      Field<"PMAXITYPE", 16, 1, ReadWrite, "AXI response that caused the precise fault. Only valid when AFSR.PMAXI is 1.">,
      Field<"PECC", 17, 1, ReadWrite, "Precise fault caused by uncorrectable ECC error">,
      Field<"PTGU", 18, 1, ReadWrite, "Precise fault caused by TGU security violation">,
      Field<"PPOISON", 19, 1, ReadWrite, "Precise fault caused by RPOISON or TEBRx.POISON">,
      Field<"FITCM", 21, 1, ReadWrite, "Fetch fault on Instruction Tightly Coupled Memory (ITCM) interface">,
      Field<"FDTCM", 22, 1, ReadWrite, "Fetch fault on Data Tightly Coupled Memory (DTCM) interface">,
      Field<"FMAXI", 24, 1, ReadWrite, "Fetch fault on Master AXI (M-AXI) interface">,
      Field<"FMAXITYPE", 27, 1, ReadWrite, "AXI response that caused the fetch fault. Only valid when AFSR.FMAXI is 1.">,
      Field<"FECC", 28, 1, ReadWrite, "Fetch fault caused by uncorrectable Error Correcting Code (ECC) error">,
      Field<"FTGU", 29, 1, ReadWrite, "Fetch fault caused by TCM Gate Unit (TGU) security violation">,
      Field<"FPOISON", 30, 1, ReadWrite, "Fetch fault caused by RPOISON or TEBRx.POISON">,
    ], "provides fault status information">,
    Register<"ID_PFR0", 0x44, 32, [
      Field<"STATE0", 0, 4, ReadWrite, "A32 instruction set support">,
      Field<"STATE1", 4, 4, ReadWrite, "T32 instruction set support">,
      Field<"RAS", 28, 4, ReadWrite, "Identifies which version of the RAS architecture is implemented">,
    ], "Indicates the version of the Reliability, Availability, and Serviceability (RAS) extension supported">,
    Register<"ID_PFR1", 0x48, 32, [
      Field<"SECURITY", 4, 4, ReadWrite, "Identifies whether the Security Extension in implemented">,
      Field<"MPROGMOD", 8, 4, ReadWrite, "M-profile programmers' model">,
    ], "Gives information about the programmers' model and Extensions support">,
    Register<"ID_DFR0", 0x4c, 32, [
      Field<"MPROFDBG", 20, 4, ReadWrite, "Indicates the supported M-Proﬁle debug architecture">,
      Field<"UDE", 28, 4, ReadWrite, "Indicates support for the Unprivileged Debug Extension">,
    ], "no information available">,
    Register<"ID_AFR0", 0x50, 32, [
    ], "The ID_AFR0 is fully Reserved, RES0">,
    Register<"ID_MMFR0", 0x54, 32, [
      Field<"PMSA", 4, 4, ReadWrite, "Indicates support for the protected memory system architecture (PMSA)">,
      Field<"OUTERSHR", 8, 4, ReadWrite, "Indicates the outermost Shareability domain implemented">,
      Field<"SHARELVL", 12, 4, ReadWrite, "Indicates the number of Shareability levels implemented">,
      Field<"TCM", 16, 4, ReadWrite, "Indicates support for TCMs">,
      Field<"AUXREG", 20, 4, ReadWrite, "Indicates support for Auxiliary Control registers">,
    ], "Gives information about the implemented memory model and memory management support">,
    Register<"ID_MMFR1", 0x58, 32, [
    ], "Gives information about the IMPLEMENTATION DEFINED memory model and memory management support. This register is Reserved, RES0.">,
    Register<"ID_MMFR2", 0x5c, 32, [
      Field<"WFISTALL", 24, 4, ReadWrite, "Indicates the support for Wait For Interrupt (WFI) stalling">,
    ], "Gives information about the implemented memory model and memory management support">,
    Register<"ID_MMFR3", 0x60, 32, [
      Field<"CMAINTVA", 0, 4, ReadWrite, "Indicates the supported cache maintenance operations by address">,
      Field<"CMAINTSW", 4, 4, ReadWrite, "Indicates the supported cache maintenance operations by set/way">,
      Field<"BPMAINT", 8, 4, ReadWrite, "Indicates the supported branch predictor maintenance">,
    ], "Gives information about the implemented memory model and memory management support">,
    Register<"ID_ISAR0", 0x64, 32, [
      Field<"BITCOUNT", 4, 4, ReadWrite, "Indicates the supported bit count instructions">,
      Field<"BITFIELD", 8, 4, ReadWrite, "Indicates the supported bit field instructions">,
      Field<"CMPBRANCH", 12, 4, ReadWrite, "Indicates the supported combined Compare and Branch instructions">,
      Field<"COPROC", 16, 4, ReadWrite, "Indicates the supported coprocessor instructions">,
      Field<"DEBUG", 20, 4, ReadWrite, "Indicates the implemented debug instructions">,
      Field<"DIVIDE", 24, 4, ReadWrite, "Indicates the supported divide instructions">,
    ], "Gives information about the implemented instruction set">,
    Register<"ID_ISAR1", 0x68, 32, [
      Field<"EXTEND", 12, 4, ReadWrite, "Indicates the implemented Extend instructions">,
      Field<"IFTHEN", 16, 4, ReadWrite, "Indicates the implemented If-Then instructions">,
      Field<"IMMEDIATE", 20, 4, ReadWrite, "Indicates the implemented data-processing instructions with long immediates">,
      Field<"INTERWORK", 24, 4, ReadWrite, "Indicates the implemented interworking instructions">,
    ], "Gives information about the implemented instruction set">,
    Register<"ID_ISAR2", 0x6c, 32, [
      Field<"LOADSTORE", 0, 4, ReadWrite, "Indicates the implemented additional load/store instructions">,
      Field<"MEMHINT", 4, 4, ReadWrite, "Memory hint. Indicates the implemented memory hint instructions">,
      Field<"MULTIACCESSINT", 8, 4, ReadWrite, "Multi-access instructions. Indicates the support for interruptible multi-access instructions">,
      Field<"MULT", 12, 4, ReadWrite, "Indicates the implemented additional Multiply instructions">,
      Field<"MULTS", 16, 4, ReadWrite, "Indicates the implemented Advanced Signed Multiple instructions">,
      Field<"MULTU", 20, 4, ReadWrite, "Indicates the implemented Advanced Unsigned Multiple instructions">,
      Field<"REVERSAL", 28, 4, ReadWrite, "Indicates the implemented reversal instructions">,
    ], "Gives information about the implemented instruction set">,
    Register<"ID_ISAR3", 0x70, 32, [
      Field<"SATURATE", 0, 4, ReadWrite, "Indicates the implemented saturating instructions">,
      Field<"SIMD", 4, 4, ReadWrite, "Indicates the implemented SIMD instructions">,
      Field<"SVC", 8, 4, ReadWrite, "Indicates the implemented SVC instructions">,
      Field<"SYNCHPRIM", 12, 4, ReadWrite, "Used with ID_ISAR4.SynchPrim_frac to indicate the implemented synchronization primitive instructions">,
      Field<"TABBRANCH", 16, 4, ReadWrite, "Indicates the implemented table branch instructions">,
      Field<"T32COPY", 20, 4, ReadWrite, "Indicates the support for T32 non-flag setting MOV instructions">,
      Field<"TRUENOP", 24, 4, ReadWrite, "Indicates the implemented true NOP instructions">,
    ], "Gives information about the implemented instruction set">,
    Register<"ID_ISAR4", 0x74, 32, [
      Field<"UNPRIV", 0, 4, ReadWrite, "Indicates the implemented unprivileged instructions">,
      Field<"WITHSHIFTS", 4, 4, ReadWrite, "Indicates the support for writeback addressing modes">,
      Field<"WRITEBACK", 8, 4, ReadWrite, "Indicates the support for writeback addressing modes">,
      Field<"BARRIER", 16, 4, ReadWrite, "Indicates the implemented Barrier instructions">,
      Field<"SYNCHPRIM_FRAC", 20, 4, ReadWrite, "Used in conjunction with ID_ISAR3.SynchPrim to indicate the implemented synchronization primitive instructions">,
      Field<"PSR_M", 24, 4, ReadWrite, "Indicates the implemented M-profile instructions to modify the PSRs">,
    ], "Gives information about the implemented instruction set">,
    Register<"CLIDR", 0x78, 32, [
      Field<"CTYPE_1", 0, 3, ReadWrite, "L1 cache type">,
      Field<"LOUIS", 21, 3, ReadWrite, "Level of Unification Inner Shareable">,
      Field<"LOC", 24, 3, ReadWrite, "Level of Coherency">,
      Field<"LOUU", 27, 3, ReadWrite, "Level of Unification Uniprocessor">,
      Field<"ICB", 30, 2, ReadWrite, "Inner cache boundary. The Cortex-M55 processor supports inner cacheability on the bus. Therefore, this field cannot disclose any information.">,
    ], "The CLIDR identifies the type of caches implemented and the level of coherency and unification">,
    Register<"CTR", 0x80, 32, [
      Field<"IMINLINE", 0, 4, ReadWrite, "Instruction cache minimum line length. Log2 of the number of words in the smallest cache line of all the instruction caches that are controlled by the processor.">,
      Field<"DMINLINE", 16, 4, ReadWrite, "Data cache minimum line length. Log2 of the number of words in the smallest cache line of all the data caches and unified caches that are controlled by the processor.">,
      Field<"ERG", 20, 4, ReadWrite, "Exclusives Reservation Granule">,
      Field<"CWG", 24, 4, ReadWrite, "Cache Write-Back Granule">,
      Field<"FORMAT", 29, 3, ReadWrite, "Cache type register format">,
    ], "The CTR provides information about the architecture of the caches">,
    Register<"CCSIDR", 0x84, 32, [
      Field<"WT", 0, 0, ReadWrite, "Indicates support available for Write-Through">,
      Field<"LINESIZE", 0, 3, ReadWrite, "Indicates the number of words in each cache line">,
      Field<"ASSOCIATIVITY", 3, 10, ReadWrite, "Indicates associativity">,
      Field<"NUMSET", 13, 15, ReadWrite, "Indicates the number of sets">,
      Field<"WA", 28, 1, ReadWrite, "Indicates support available for write allocation">,
      Field<"RA", 29, 1, ReadWrite, "Indicates support available for read allocation">,
      Field<"WB", 30, 1, ReadWrite, "Indicates support available for Write-Back">,
    ], "Provides information about the architecture of the Level 1 (L1) instruction or data cache that the CSSELR selects.">,
    Register<"CSSELR", 0x88, 32, [
      Field<"IND", 0, 1, ReadWrite, "Instruction not data bit">,
      Field<"LEVEL", 1, 3, ReadWrite, "Cache level of required cache">,
    ], "The CSSELR selects the current CCSIDR by specifying the cache level and the type of cache">,
    Register<"CPACR", 0x8c, 32, [
      Field<"CP", 0, 2, ReadWrite, "Coprocessor %s privilege">,
      Field<"CP_10", 20, 2, ReadWrite, "CP10 Privilege. Defines the access rights for the floating-point functionality.", [
        Enum<"Disabled", 0, "All accesses to the FP Extension result in NOCP UsageFault.">,
        Enum<"Unprivileged", 1, "Unprivileged accesses to the FP Extension result in NOCP UsageFault.">,
        Enum<"Full", 3, "Full access to the FP Extension.">,
      ]>,
      Field<"CP_11", 22, 2, ReadWrite, "CP11 Privilege. The value in this field is ignored.", [
        Enum<"Disabled", 0, "All accesses to the FP Extension result in NOCP UsageFault.">,
        Enum<"Unprivileged", 1, "Unprivileged accesses to the FP Extension result in NOCP UsageFault.">,
        Enum<"Full", 3, "Full access to the FP Extension.">,
      ]>,
    ], "Specifies the access privileges for coprocessors">,
    Register<"NSACR", 0x90, 32, [
      Field<"CP", 0, 1, ReadWrite, "Enables Non-secure access to this coprocessor">,
      Field<"CP_11", 1, 11, ReadWrite, "Enables non-secure access to the Floating-point and MVE. Programming with a different value other than that used for CP10 is unpredictable.">,
      Field<"CP_10", 10, 1, ReadWrite, "Enables non-secure access to the Floating-point and MVE">,
    ], "In an implementation with the Security Extension, the NSACR register defines the Non-secure access permissions for both the Floating-point and MVE and coprocessors CP m, bit[ m], for m = 0-7. If MVE is implemented, this register specifies the Non-secure permissions for MVE">,
  ];
}

class SCBInstance<string Name, int Base> : PeripheralInstance<Name, Base, SCBPeripheral>;
def SCB : SCBInstance<"SCB", 0xe000ecfc>;

#endif // _PERIPHERALS_SCB_TD
