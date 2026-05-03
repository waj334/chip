ENTRY(ResetHandler)

SECTIONS
{
    .text :
    {
        KEEP(*(.isr_vector))
        *(.text)
        *(.text.*)
        *(.rodata)
        *(.rodata.*)
        . = ALIGN(8);

        /* Required for global ctors. */
        PROVIDE_HIDDEN ( __init_array_start = . );
        KEEP (*(SORT_BY_INIT_PRIORITY(.init_array.*) SORT_BY_INIT_PRIORITY(.ctors.*)))
        KEEP (*(.init_array .ctors))
        PROVIDE_HIDDEN ( __init_array_end = . );

    } >FLASH

    .data :
    {
        __start_data = .;
        *(.data)
        *(.data.*)
        . = ALIGN(8);
        __end_data = .;
    } >RAM AT>FLASH
    __data_base_addr = LOADADDR(.data);

    .bss (NOLOAD) :
    {
        __start_bss = .;
        *(.bss)
        *(.bss.*)
        *(COMMON)
        . = ALIGN(8);
        __end_bss = .;
    } >RAM

    /* Heap grows up from end of BSS. */
    __heap_start = __end_bss;
    __heap_end = ORIGIN(RAM) + LENGTH(RAM) - __stack_size;
    __heap_size = __heap_end - __heap_start;

    /* Stack at the very end of RAM. Overflow goes past RAM_END and faults. */
   .stack ORIGIN(RAM) + LENGTH(RAM) - __stack_size (NOLOAD) :
   {
       __stack_bottom = .;
       . += __stack_size;
       . = ALIGN(8);
       __stack_top = .;
   } >RAM

    /* Define the range that the garbage collector
       should scan for heap references. */
    __gc_scan_start = __start_data;
    __gc_scan_end   = __end_bss;

    /* Exception handling is not needed. */
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
__globals_end = __end_bss;
_estack = __stack_top;
__stack = __stack_top;

/* Boehm GC */
__data_start = __gc_scan_start;
end = __gc_scan_end;
