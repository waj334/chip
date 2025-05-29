#ifndef _PERIPHERALS_FLASH_TD
#define _PERIPHERALS_FLASH_TD

include "base.td"

def FlashPeripheral : RepeatingPeripheralType<"Flash", 2, "Flash"> {
  let accessWidth = 32;
  let count = 2;
  let arrayLabel = "Bank";
  let registers = [
    Register<"ACR", 0x0, 32, [
      Field<"LATENCY", 0, 3, ReadWrite, "Read latency">,
      Field<"WRHIGHFREQ", 4, 2, ReadWrite, "Flash signal delay">,
    ], "Access control register">,
    Register<"KEYR", 0x4, 32, [
      Field<"KEYR", 0, 32, ReadWrite, "Bank n access configuration unlock key">,
    ], "FLASH key register for bank n">,
    Register<"OPTKEYR", 0x8, 32, [
      Field<"OPTKEYR", 0, 32, ReadWrite, "Unlock key option bytes">,
    ], "FLASH option key register">,
    Register<"CR", 0xc, 32, [
      Field<"LOCK", 0, 1, ReadWrite, "Bank n configuration lock bit">,
      Field<"PG", 1, 1, ReadWrite, "Bank n program enable bit">,
      Field<"SER", 2, 1, ReadWrite, "Bank n sector erase request">,
      Field<"BER", 3, 1, ReadWrite, "Bank n erase request">,
      Field<"PSIZE", 4, 2, ReadWrite, "Bank n program size">,
      Field<"FW", 6, 1, ReadWrite, "Bank n write forcing control bit">,
      Field<"START", 7, 1, ReadWrite, "Bank n bank or sector erase start control bit">,
      Field<"SNB", 8, 3, ReadWrite, "Bank n sector erase selection number">,
      Field<"CRC_EN", 15, 1, ReadWrite, "Bank n CRC control bit">,
      Field<"EOPIE", 16, 1, ReadWrite, "Bank n end-of-program interrupt control bit">,
      Field<"WRPERRIE", 17, 1, ReadWrite, "Bank n write protection error interrupt enable bit">,
      Field<"PGSERRIE", 18, 1, ReadWrite, "Bank n programming sequence error interrupt enable bit">,
      Field<"STRBERRIE", 19, 1, ReadWrite, "Bank n strobe error interrupt enable bit">,
      Field<"INCERRIE", 21, 1, ReadWrite, "Bank n inconsistency error interrupt enable bit">,
      Field<"OPERRIE", 22, 1, ReadWrite, "Bank n write/erase error interrupt enable bit">,
      Field<"RDPERRIE", 23, 1, ReadWrite, "Bank n read protection error interrupt enable bit">,
      Field<"RDSERRIE", 24, 1, ReadWrite, "Bank n secure error interrupt enable bit">,
      Field<"SNECCERRIE", 25, 1, ReadWrite, "Bank n ECC single correction error interrupt enable bit">,
      Field<"DBECCERRIE", 26, 1, ReadWrite, "Bank n ECC double detection error interrupt enable bit">,
      Field<"CRCENDIE", 27, 1, ReadWrite, "Bank n end of CRC calculation interrupt enable bit">,
    ], "FLASH control register for bank n">,
    Register<"SR", 0x10, 32, [
      Field<"BSY", 0, 1, ReadWrite, "Bank n ongoing program flag">,
      Field<"WBNE", 1, 1, ReadWrite, "Bank n write buffer not empty flag">,
      Field<"QW", 2, 1, ReadWrite, "Bank n wait queue flag">,
      Field<"CRC_BUSY", 3, 1, ReadWrite, "Bank n CRC busy flag">,
      Field<"EOP", 16, 1, ReadWrite, "Bank n end-of-program flag">,
      Field<"WRPERR", 17, 1, ReadWrite, "Bank n write protection error flag">,
      Field<"PGSERR", 18, 1, ReadWrite, "Bank n programming sequence error flag">,
      Field<"STRBERR", 19, 1, ReadWrite, "Bank n strobe error flag">,
      Field<"INCERR", 21, 1, ReadWrite, "Bank n inconsistency error flag">,
      Field<"OPERR", 22, 1, ReadWrite, "Bank n write/erase error flag">,
      Field<"RDPERR", 23, 1, ReadWrite, "Bank n read protection error flag">,
      Field<"RDSERR", 24, 1, ReadWrite, "Bank n secure error flag">,
      Field<"SNECCERR1", 25, 1, ReadWrite, "Bank n single correction error flag">,
      Field<"DBECCERR", 26, 1, ReadWrite, "Bank n ECC double detection error flag">,
      Field<"CRCEND", 27, 1, ReadWrite, "Bank n CRC-complete flag">,
    ], "FLASH status register for bank n">,
    Register<"CCR", 0x14, 32, [
      Field<"CLR_EOP", 16, 1, ReadWrite, "Bank n EOP1 flag clear bit">,
      Field<"CLR_WRPERR", 17, 1, ReadWrite, "Bank n WRPERR1 flag clear bit">,
      Field<"CLR_PGSERR", 18, 1, ReadWrite, "Bank n PGSERR1 flag clear bi">,
      Field<"CLR_STRBERR", 19, 1, ReadWrite, "Bank n STRBERR1 flag clear bit">,
      Field<"CLR_INCERR", 21, 1, ReadWrite, "Bank n INCERR1 flag clear bit">,
      Field<"CLR_OPERR", 22, 1, ReadWrite, "Bank n OPERR1 flag clear bit">,
      Field<"CLR_RDPERR", 23, 1, ReadWrite, "Bank n RDPERR1 flag clear bit">,
      Field<"CLR_RDSERR", 24, 1, ReadWrite, "Bank n RDSERR1 flag clear bit">,
      Field<"CLR_SNECCERR", 25, 1, ReadWrite, "Bank n SNECCERR1 flag clear bit">,
      Field<"CLR_DBECCERR", 26, 1, ReadWrite, "Bank n DBECCERR1 flag clear bit">,
      Field<"CLR_CRCEND", 27, 1, ReadWrite, "Bank n CRCEND1 flag clear bit">,
    ], "FLASH clear control register for bank n">,
    Register<"OPTCR", 0x18, 32, [
      Field<"OPTLOCK", 0, 1, ReadWrite, "FLASH_OPTCR lock option configuration bit">,
      Field<"OPTSTART", 1, 1, ReadWrite, "Option byte start change option configuration bit">,
      Field<"MER", 4, 1, ReadWrite, "Flash mass erase enable bit">,
      Field<"OPTCHANGEERRIE", 30, 1, ReadWrite, "Option byte change error interrupt enable bit">,
      Field<"SWAP_BANK", 31, 1, ReadWrite, "Bank swapping configuration bit">,
    ], "FLASH option control register">,
    Register<"OPTSR_CUR", 0x1c, 32, [
      Field<"OPT_BUSY", 0, 1, ReadWrite, "Option byte change ongoing flag">,
      Field<"BOR_LEV", 2, 2, ReadWrite, "Brownout level option status bit">,
      Field<"IWDG1_HW", 4, 1, ReadWrite, "IWDG1 control option status bit">,
      Field<"NRST_STOP_D", 6, 1, ReadWrite, "D1 DStop entry reset option status bit">,
      Field<"NRST_STBY_D", 7, 1, ReadWrite, "D1 DStandby entry reset option status bit">,
      Field<"RDP", 8, 8, ReadWrite, "Readout protection level option status byte">,
      Field<"FZ_IWDG_STOP", 17, 1, ReadWrite, "IWDG Stop mode freeze option status bit">,
      Field<"FZ_IWDG_SDBY", 18, 1, ReadWrite, "IWDG Standby mode freeze option status bit">,
      Field<"ST_RAM_SIZE", 19, 2, ReadWrite, "DTCM RAM size option status">,
      Field<"SECURITY", 21, 1, ReadWrite, "Security enable option status bit">,
      Field<"RSS1", 26, 1, ReadWrite, "User option bit 1">,
      Field<"PERSO_OK", 28, 1, ReadWrite, "Device personalization status bit">,
      Field<"IO_HSLV", 29, 1, ReadWrite, "I/O high-speed at low-voltage status bit (PRODUCT_BELOW_25V)">,
      Field<"OPTCHANGEERR", 30, 1, ReadWrite, "Option byte change error flag">,
      Field<"SWAP_BANK_OPT", 31, 1, ReadWrite, "Bank swapping option status bit">,
    ], "FLASH option status register">,
    Register<"OPTSR_PRG", 0x20, 32, [
      Field<"BOR_LEV", 2, 2, ReadWrite, "BOR reset level option configuration bits">,
      Field<"IWDG1_HW", 4, 1, ReadWrite, "IWDG1 option configuration bit">,
      Field<"NRST_STOP_D1", 6, 1, ReadWrite, "Option byte erase after D1 DStop option configuration bit">,
      Field<"NRST_STBY_D1", 7, 1, ReadWrite, "Option byte erase after D1 DStandby option configuration bit">,
      Field<"RDP", 8, 8, ReadWrite, "Readout protection level option configuration byte">,
      Field<"FZ_IWDG_STOP", 17, 1, ReadWrite, "IWDG Stop mode freeze option configuration bit">,
      Field<"FZ_IWDG_SDBY", 18, 1, ReadWrite, "IWDG Standby mode freeze option configuration bit">,
      Field<"ST_RAM_SIZE", 19, 2, ReadWrite, "DTCM size select option configuration bits">,
      Field<"SECURITY", 21, 1, ReadWrite, "Security option configuration bit">,
      Field<"RSS1", 26, 1, ReadWrite, "User option configuration bit 1">,
      Field<"RSS2", 27, 1, ReadWrite, "User option configuration bit 2">,
      Field<"IO_HSLV", 29, 1, ReadWrite, "I/O high-speed at low-voltage (PRODUCT_BELOW_25V)">,
      Field<"SWAP_BANK_OPT", 31, 1, ReadWrite, "Bank swapping option configuration bit">,
    ], "FLASH option status register">,
    Register<"OPTCCR", 0x24, 32, [
      Field<"CLR_OPTCHANGEERR", 30, 1, ReadWrite, "OPTCHANGEERR reset bit">,
    ], "FLASH option clear control register">,
    Register<"PRAR_CUR", 0x28, 32, [
      Field<"PROT_AREA_START", 0, 12, ReadWrite, "Bank n lowest PCROP protected address">,
      Field<"PROT_AREA_END", 16, 12, ReadWrite, "Bank n highest PCROP protected address">,
      Field<"DMEP", 31, 1, ReadWrite, "Bank n PCROP protected erase enable option status bit">,
    ], "FLASH protection address for bank n">,
    Register<"PRAR_PRG2", 0x2c, 32, [
      Field<"PROT_AREA_START2", 0, 12, ReadWrite, "Bank 2 lowest PCROP protected address configuration">,
      Field<"PROT_AREA_END2", 16, 12, ReadWrite, "Bank 2 highest PCROP protected address configuration">,
      Field<"DMEP2", 31, 1, ReadWrite, "Bank 2 PCROP protected erase enable option configuration bit">,
    ], "FLASH protection address for bank 2">,
    Register<"PRAR_PRG", 0x2c, 32, [
      Field<"PROT_AREA_START", 0, 12, ReadWrite, "Bank n lowest PCROP protected address configuration">,
      Field<"PROT_AREA_END", 16, 12, ReadWrite, "Bank n highest PCROP protected address configuration">,
      Field<"DMEP", 31, 1, ReadWrite, "Bank n PCROP protected erase enable option configuration bit">,
    ], "FLASH protection address for bank n">,
    Register<"SCAR_CUR", 0x30, 32, [
      Field<"SEC_AREA_START", 0, 12, ReadWrite, "Bank n lowest secure protected address">,
      Field<"SEC_AREA_END", 16, 12, ReadWrite, "Bank n highest secure protected address">,
      Field<"DMES", 31, 1, ReadWrite, "Bank n secure protected erase enable option status bit">,
    ], "FLASH secure address for bank n">,
    Register<"SCAR_PRG", 0x34, 32, [
      Field<"SEC_AREA_START", 0, 12, ReadWrite, "Bank n lowest secure protected address configuration">,
      Field<"SEC_AREA_END", 16, 12, ReadWrite, "Bank n highest secure protected address configuration">,
      Field<"DMES", 31, 1, ReadWrite, "Bank n secure protected erase enable option configuration bit">,
    ], "FLASH secure address for bank n">,
    Register<"WPSN_CUR1R", 0x38, 32, [
      Field<"WRPSN", 0, 8, ReadWrite, "Bank n sector write protection option status byte">,
    ], "FLASH write sector protection for bank n">,
    Register<"WPSN_PRG1R", 0x3c, 32, [
      Field<"WRPSN", 0, 8, ReadWrite, "Bank n sector write protection configuration byte">,
    ], "FLASH write sector protection for bank n">,
    Register<"BOOT_CURR", 0x40, 32, [
      Field<"BOOT_ADD0", 0, 16, ReadWrite, "Boot address 0">,
      Field<"BOOT_ADD1", 16, 16, ReadWrite, "Boot address 1">,
    ], "FLASH register with boot address">,
    Register<"BOOT_PRGR", 0x44, 32, [
      Field<"BOOT_ADD0", 0, 16, ReadWrite, "Boot address 0">,
      Field<"BOOT_ADD1", 16, 16, ReadWrite, "Boot address 1">,
    ], "FLASH register with boot address">,
    Register<"CRCCR", 0x50, 32, [
      Field<"CRC_SECT", 0, 3, ReadWrite, "Bank n CRC sector number">,
      Field<"ALL_BANK", 7, 1, ReadWrite, "Bank n CRC select bit">,
      Field<"CRC_BY_SECT", 8, 1, ReadWrite, "Bank n CRC sector mode select bit">,
      Field<"ADD_SECT", 9, 1, ReadWrite, "Bank n CRC sector select bit">,
      Field<"CLEAN_SECT", 10, 1, ReadWrite, "Bank n CRC sector list clear bit">,
      Field<"START_CRC", 16, 1, ReadWrite, "Bank n CRC start bit">,
      Field<"CLEAN_CRC", 17, 1, ReadWrite, "Bank n CRC clear bit">,
      Field<"CRC_BURST", 20, 2, ReadWrite, "Bank n CRC burst size">,
    ], "FLASH CRC control register for bank n">,
    Register<"CRCSADD1R", 0x54, 32, [
      Field<"CRC_START_ADDR", 0, 32, ReadWrite, "CRC start address on bank n">,
    ], "FLASH CRC start address register for bank n">,
    Register<"CRCEADD1R", 0x58, 32, [
      Field<"CRC_END_ADDR", 0, 32, ReadWrite, "CRC end address on bank n">,
    ], "FLASH CRC end address register for bank n">,
    Register<"CRCDATAR", 0x5c, 32, [
      Field<"CRC_DATA", 0, 32, ReadWrite, "CRC result">,
    ], "FLASH CRC data register">,
    Register<"ECC_FA1R", 0x60, 32, [
      Field<"FAIL_ECC_ADDR", 0, 15, ReadWrite, "Bank n ECC error address">,
    ], "FLASH ECC fail address for bank n">,
  ];
}

class FlashInstance<string Name, int Base> : PeripheralInstance<Name, Base, FlashPeripheral>;
def Flash : FlashInstance<"Flash", 0x52002000>;

#endif // _PERIPHERALS_FLASH_TD
