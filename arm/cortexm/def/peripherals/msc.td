#ifndef _PERIPHERALS_MSC_TD
#define _PERIPHERALS_MSC_TD

include "base.td"

def MSCPeripheral : PeripheralType<"MSC", "Memory System Control Registers"> {
  let accessWidth = 32;
  let registers = [
    Register<"MSCR", 0x0, 32, [
      Field<"ECCEN", 1, 1, ReadWrite, "Indicates whether Error Correcting Code (ECC) is present and enabled">,
      Field<"FORCEWT", 2, 1, ReadWrite, "Enables Forced Write-Through in the L1 data cache">,
      Field<"EVECCFAULT", 3, 1, ReadWrite, "Enables asynchronous BusFault exceptions when data is lost on evictions">,
      Field<"DCACTIVE", 12, 1, ReadWrite, "Indicates whether the L1 data cache is active">,
      Field<"ICACTIVE", 13, 1, ReadWrite, "Indicates whether the L1 instruction cache is active">,
      Field<"DCCLEAN", 16, 1, ReadWrite, "Indicates whether the data cache contains any dirty lines">,
      Field<"CPWRDN", 17, 1, ReadWrite, "Indicates when the data and instruction caches are not accessible because they are either being powered down or being initialized using the automatic invalidation sequence">,
    ], "Controls the memory system features specific to the processor">,
    Register<"PFCR", 0x4, 32, [
      Field<"ENABLE", 0, 1, ReadWrite, "Prefetcher enable">,
      Field<"DIS_NLP", 7, 1, ReadWrite, "Disables Next Line Prefetch mode">,
    ], "Controls the prefetcher">,
    Register<"ITCMCR", 0x10, 32, [
      Field<"EN", 0, 1, ReadWrite, "TCM enable">,
      Field<"SZ", 3, 4, ReadWrite, "Indicates the size of the relevant TCM">,
    ], "Enable access to the Tightly Coupled Memories (TCMs) by software running on the processor">,
    Register<"DTCMCR", 0x14, 32, [
      Field<"EN", 0, 1, ReadWrite, "TCM enable">,
      Field<"SZ", 3, 4, ReadWrite, "Indicates the size of the relevant TCM">,
    ], "Enable access to the Tightly Coupled Memories (TCMs) by software running on the processor">,
    Register<"PAHBCR", 0x18, 32, [
      Field<"EN", 0, 1, ReadWrite, "P-AHB enable">,
      Field<"SZ", 1, 3, ReadWrite, "P-AHB size">,
    ], "Enables accesses to Peripheral AHB (P-AHB) interface from software running on the processor">,
  ];
}

class MSCInstance<string Name, int Base> : PeripheralInstance<Name, Base, MSCPeripheral>;
def MSC : MSCInstance<"MSC", 0xe001e000>;

#endif // _PERIPHERALS_MSC_TD
