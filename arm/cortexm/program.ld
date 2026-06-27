ENTRY(ResetHandler)

SECTIONS
{
    /* Code, read-only data, init arrays. */
    .text :
    {
        KEEP(*(.isr_vector))

        *(.text)
        *(.text.*)

        *(.rodata)
        *(.rodata.*)

        /*
         * Picolibc constructor arrays.
         *
         * __bothinit_array_* covers the contiguous preinit and init
         * arrays for Picolibc builds without _init/_fini functions.
         */
        . = ALIGN(8);

        PROVIDE_HIDDEN(__preinit_array_start = .);
        PROVIDE_HIDDEN(__bothinit_array_start = .);

        KEEP(*(.preinit_array))

        PROVIDE_HIDDEN(__preinit_array_end = .);
        PROVIDE_HIDDEN(__init_array_start = .);

        KEEP(*(SORT_BY_INIT_PRIORITY(.init_array.*)
               SORT_BY_INIT_PRIORITY(.ctors.*)))
        KEEP(*(.init_array .ctors))

        PROVIDE_HIDDEN(__init_array_end = .);
        PROVIDE_HIDDEN(__bothinit_array_end = .);
    } > TEXT

    /* Initialized data. VMA in RAM, LMA in TEXT (startup copies). */
    .data :
    {
        __start_data = .;
        *(.data)
        *(.data.*)
        . = ALIGN(8);
        __end_data = .;
    } > RAM AT> TEXT

    __data_base_addr = LOADADDR(.data);

    /* Zero-initialized data. */
    .bss (NOLOAD) :
    {
        __start_bss = .;
        *(.bss)
        *(.bss.*)
        *(COMMON)
        . = ALIGN(8);
        __end_bss = .;
    } > RAM

    /*
     * Explicit sections that live in the heap region must be emitted before
     * .heap consumes the rest of HEAP.
     *
     * Board/linker wrappers can define HEAP_RESERVED_SECTIONS to place fixed
     * framebuffers, DMA pools, external RAM objects, etc. before the heap.
     */
    #ifdef HEAP_RESERVED_SECTIONS
        HEAP_RESERVED_SECTIONS
    #endif

    .heap (NOLOAD) :
    {
        . = ALIGN(8);
        __heap_start = .;

        . = (ORIGIN(STACK) == ORIGIN(HEAP))
            ? (ORIGIN(HEAP) + LENGTH(HEAP) - __stack_size)
            : (ORIGIN(HEAP) + LENGTH(HEAP));

        . = ALIGN(8);
        __heap_end = .;
    } > HEAP

    .stack ((ORIGIN(STACK) + LENGTH(STACK) - __stack_size) & ~7) (NOLOAD) :
    {
        . = ALIGN(8);
        __stack_bottom = .;

        . += __stack_size;
        . = ALIGN(8);
        __stack_top = .;
    } > STACK

    ASSERT((__stack_size & 7) == 0, "__stack_size must be 8-byte aligned")
    ASSERT((__stack_bottom & 7) == 0, "stack bottom is not 8-byte aligned")
    ASSERT((__stack_top & 7) == 0, "stack top is not 8-byte aligned")
    ASSERT(__heap_end >= __heap_start, "heap region overflowed by explicit sections")
    ASSERT(__stack_top <= ORIGIN(STACK) + LENGTH(STACK), "stack exceeds stack region")
    ASSERT(__heap_end <= __stack_bottom || ORIGIN(STACK) != ORIGIN(HEAP), "heap overlaps stack")

    /* GC scans the contiguous globals range covering .data and .bss. */
    __gc_scan_start = __start_data;
    __gc_scan_end   = __end_bss;

    /* Exception handling not needed on bare metal. */
    /DISCARD/ :
    {
        *(.note .note.*)
        *(.eh_frame .eh_frame.*)
        *(.ARM.extab* .gnu.linkonce.armextab.*)
        *(.ARM.exidx*)
        *(.fini_array .fini_array.*)
        *(.dtors .dtors.*)
    }
}

__globals_start = __start_data;
__globals_end   = __end_bss;
_estack         = __stack_top;
__stack         = __stack_top;
