MEMORY
{
 ITCM        (rwx) : ORIGIN = 0x0000000000, LENGTH = 64K  /* ITCM RAM (instruction) */
 FLASH1      (rwx) : ORIGIN = 0x0008000000, LENGTH = 1M   /* Flash bank 1 (1MB) */
 FLASH       (rwx) : ORIGIN = 0x0008000000, LENGTH = 2M   /* Flash bank [1:2] (2MB) */
 FLASH2      (rwx) : ORIGIN = 0x0008100000, LENGTH = 1M   /* Flash bank 2 (1MB) */
 DTCM        (rwx) : ORIGIN = 0x0020000000, LENGTH = 128K /* DTCM RAM (instruction) */
 SRAM_AXI    (rwx) : ORIGIN = 0x0024000000, LENGTH = 512K /* SRAM mapped on AXI bus */
 SRAM1       (rwx) : ORIGIN = 0x0030000000, LENGTH = 128K /* SRAM1 (D2 domain) */
 SRAM2       (rwx) : ORIGIN = 0x0030020000, LENGTH = 128K /* SRAM2 (D2 domain) */
 SRAM3       (rwx) : ORIGIN = 0x0030040000, LENGTH = 32K  /* SRAM3 (D2 domain) */
 SRAM4       (rwx) : ORIGIN = 0x0038000000, LENGTH = 64K  /* SRAM4 (D3 domain) */
 SRAM_BACKUP (rwx) : ORIGIN = 0x0038800000, LENGTH = 4K   /* Backup SRAM (Kbytes) */
#ifdef EXTRA_MEMORY_REGIONS
 EXTRA_MEMORY_REGIONS
#endif
}

#ifndef TEXT_REGION
#define TEXT_REGION FLASH
#endif
#ifndef RAM_REGION
#define RAM_REGION SRAM_AXI
#endif
#ifndef HEAP_REGION
#define HEAP_REGION SRAM_AXI
#endif
#ifndef STACK_REGION
#define STACK_REGION SRAM_AXI
#endif

REGION_ALIAS("TEXT",  TEXT_REGION);
REGION_ALIAS("RAM",   RAM_REGION);
REGION_ALIAS("HEAP",  HEAP_REGION);
REGION_ALIAS("STACK", STACK_REGION);

SECTIONS
{
	.itcm (NOLOAD) : ALIGN(32)
	{
		__itcm_start = .;
		KEEP(*(.itcm))
		KEEP(*(.itcm.*))
		. = ALIGN(32);
		__itcm_end = .;
	} > ITCM

	.flash1 (NOLOAD) : ALIGN(32)
	{
		__flash1_start = .;
		KEEP(*(.flash1))
		KEEP(*(.flash1.*))
		. = ALIGN(32);
		__flash1_end = .;
	} > FLASH1

	.flash (NOLOAD) : ALIGN(32)
	{
		__flash_start = .;
		KEEP(*(.flash))
		KEEP(*(.flash.*))
		. = ALIGN(32);
		__flash_end = .;
	} > FLASH

	.flash2 (NOLOAD) : ALIGN(32)
	{
		__flash2_start = .;
		KEEP(*(.flash2))
		KEEP(*(.flash2.*))
		. = ALIGN(32);
		__flash2_end = .;
	} > FLASH2

	.dtcm (NOLOAD) : ALIGN(32)
	{
		__dtcm_start = .;
		KEEP(*(.dtcm))
		KEEP(*(.dtcm.*))
		. = ALIGN(32);
		__dtcm_end = .;
	} > DTCM

	.sram_axi (NOLOAD) : ALIGN(32)
	{
		__sram_axi_start = .;
		KEEP(*(.sram_axi))
		KEEP(*(.sram_axi.*))
		. = ALIGN(32);
		__sram_axi_end = .;
	} > SRAM_AXI

	.sram1 (NOLOAD) : ALIGN(32)
	{
		__sram1_start = .;
		KEEP(*(.sram1))
		KEEP(*(.sram1.*))
		. = ALIGN(32);
		__sram1_end = .;
	} > SRAM1

	.sram2 (NOLOAD) : ALIGN(32)
	{
		__sram2_start = .;
		KEEP(*(.sram2))
		KEEP(*(.sram2.*))
		. = ALIGN(32);
		__sram2_end = .;
	} > SRAM2

	.sram3 (NOLOAD) : ALIGN(32)
	{
		__sram3_start = .;
		KEEP(*(.sram3))
		KEEP(*(.sram3.*))
		. = ALIGN(32);
		__sram3_end = .;
	} > SRAM3

	.sram4 (NOLOAD) : ALIGN(32)
	{
		__sram4_start = .;
		KEEP(*(.sram4))
		KEEP(*(.sram4.*))
		. = ALIGN(32);
		__sram4_end = .;
	} > SRAM4

	.sram_backup (NOLOAD) : ALIGN(32)
	{
		__sram_backup_start = .;
		KEEP(*(.sram_backup))
		KEEP(*(.sram_backup.*))
		. = ALIGN(32);
		__sram_backup_end = .;
	} > SRAM_BACKUP
}

__stack_size = 4K;
#include <program.ld>
