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

        . = ALIGN(8);
        /* Required for global ctors. */
        PROVIDE_HIDDEN(__init_array_start = .);
        KEEP(*(SORT_BY_INIT_PRIORITY(.init_array.*) SORT_BY_INIT_PRIORITY(.ctors.*)))
        KEEP(*(.init_array .ctors))
        PROVIDE_HIDDEN(__init_array_end = .);
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

    /* Heap grows up from the end of .bss when HEAP and RAM are the same
       region; otherwise the heap fills its own region.

       The boundary computation handles both cases: when HEAP == RAM, the
       heap starts after .bss and ends before the stack reservation.
       When HEAP is a separate region, the heap fills it entirely. */
    .heap (NOLOAD) :
    {
        . = ALIGN(8);
        __heap_start = .;

        /* When heap shares region with stack, leave room for stack */
        /* When heap is in a separate region, use the whole region */
        . = (ORIGIN(STACK) == ORIGIN(HEAP))
            ? (ORIGIN(HEAP) + LENGTH(HEAP) - __stack_size)
            : (ORIGIN(HEAP) + LENGTH(HEAP));

        __heap_end = .;
    } > HEAP

    /* Stack at the top of its region. Stack overflow runs past the region
       end and faults via the MPU or just an unmapped-address bus error. */
    .stack ORIGIN(STACK) + LENGTH(STACK) - __stack_size (NOLOAD) :
    {
        __stack_bottom = .;
        . += __stack_size;
        . = ALIGN(8);
        __stack_top = .;
    } > STACK

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
