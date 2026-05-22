#ifndef _ARM_CORTEXM_VARIANT_TD
#define _ARM_CORTEXM_VARIANT_TD

include "base.td"
include "arm/cortexm/interrupts.td"

defvar CortexMTags = [ "arm", "cortexm" ];

class CortexMVariant<string Name, list<MemoryRange> Memories = [], list<Interrupt> Interrupts = [],
    MemoryUnit StackSize = KB<4>, list<string> Tags, string Description = "",
    MemoryRange DefaultTextRegion = UnsetMemoryRange,
    MemoryRange DefaultRAMRegion = UnsetMemoryRange,
    MemoryRange DefaultHeapRegion = UnsetMemoryRange,
    MemoryRange DefaultStackRegion = UnsetMemoryRange>
        : Variant<Name, Memories, !listconcat(CortexMInterrupts, Interrupts), StackSize, !listconcat(CortexMTags, Tags),
            Description, DefaultTextRegion, DefaultRAMRegion, DefaultHeapRegion, DefaultStackRegion>;

#endif // _ARM_CORTEXM_VARIANT_TD
