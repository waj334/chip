MEMORY
{
	ITCM (rwx) : ORIGIN = 0x0, LENGTH = 64K				/* ITCM RAM (instruction) */
	FLASH (rwx) : ORIGIN = 0x8000000, LENGTH = 512K				/* Flash memory bank 1 */
	FLASH2 (rwx) : ORIGIN = 0x8100000, LENGTH = 512K				/* Flash memory bank 2 */
	DTCM (rwx) : ORIGIN = 0x20000000, LENGTH = 128K				/* DTCM RAM (instruction) */
	SRAM_AXI (rwx) : ORIGIN = 0x24000000, LENGTH = 512K				/* SRAM mapped on AXI bus */
	RAM (rwx) : ORIGIN = 0x30000000, LENGTH = 288K				/* [SRAM1 : SRAM3] */
	SRAM1 (rwx) : ORIGIN = 0x30000000, LENGTH = 128K				/* SRAM1 (D2 domain) */
	SRAM2 (rwx) : ORIGIN = 0x30020000, LENGTH = 128K				/* SRAM2 (D2 domain) */
	SRAM3 (rwx) : ORIGIN = 0x30040000, LENGTH = 32K				/* SRAM3 (D2 domain) */
	SRAM4 (rwx) : ORIGIN = 0x38000000, LENGTH = 64K				/* SRAM4 (D3 domain) */
	SRAM_BACKUP (rwx) : ORIGIN = 0x38800000, LENGTH = 4K				/* Backup SRAM (Kbytes) */
}

__stack_size = 4K;
INCLUDE program.ld

