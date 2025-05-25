#ifndef _PERIPHERALS_SPI_TD
#define _PERIPHERALS_SPI_TD

include "base.td"

def SPIPeripheral : PeripheralType<"SPI", "Serial peripheral interface"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR1", 0x0, 32, [
      Field<"SPE", 0, 1, ReadWrite, "Serial Peripheral Enable">,
      Field<"MASRX", 8, 1, ReadWrite, "Master automatic SUSP in Receive mode">,
      Field<"CSTART", 9, 1, Read, "Master transfer start">,
      Field<"CSUSP", 10, 1, Write, "Master SUSPend request">,
      Field<"HDDIR", 11, 1, ReadWrite, "Rx/Tx direction at Half-duplex mode">,
      Field<"SSI", 12, 1, ReadWrite, "Internal SS signal input level">,
      Field<"CRC33_17", 13, 1, ReadWrite, "32-bit CRC polynomial configuration">,
      Field<"RCRCI", 14, 1, ReadWrite, "CRC calculation initialization pattern control for receiver">,
      Field<"TCRCI", 15, 1, ReadWrite, "CRC calculation initialization pattern control for transmitter">,
      Field<"IOLOCK", 16, 1, Read, "Locking the AF configuration of associated IOs">,
    ], "control register 1">,
    Register<"CR2", 0x4, 32, [
      Field<"TSIZE", 0, 16, ReadWrite, "Number of data at current transfer">,
      Field<"TSER", 16, 16, Read, "Number of data transfer extension to be reload into TSIZE just when a previous">,
    ], "control register 2">,
    Register<"CFG1", 0x8, 32, [
      Field<"DSIZE", 0, 5, ReadWrite, "Number of bits in at single SPI data frame">,
      Field<"FTHVL", 5, 4, ReadWrite, "threshold level">,
      Field<"UDRCFG", 9, 2, ReadWrite, "Behavior of slave transmitter at underrun condition">,
      Field<"UDRDET", 11, 2, ReadWrite, "Detection of underrun condition at slave transmitter">,
      Field<"RXDMAEN", 14, 1, ReadWrite, "Rx DMA stream enable">,
      Field<"TXDMAEN", 15, 1, ReadWrite, "Tx DMA stream enable">,
      Field<"CRCSIZE", 16, 5, ReadWrite, "Length of CRC frame to be transacted and compared">,
      Field<"CRCEN", 22, 1, ReadWrite, "Hardware CRC computation enable">,
      Field<"MBR", 28, 3, ReadWrite, "Master baud rate">,
    ], "configuration register 1">,
    Register<"CFG2", 0xc, 32, [
      Field<"MSSI", 0, 4, ReadWrite, "Master SS Idleness">,
      Field<"MIDI", 4, 4, ReadWrite, "Master Inter-Data Idleness">,
      Field<"IOSWP", 15, 1, ReadWrite, "Swap functionality of MISO and MOSI pins">,
      Field<"COMM", 17, 2, ReadWrite, "SPI Communication Mode">,
      Field<"SP", 19, 3, ReadWrite, "Serial Protocol">,
      Field<"MASTER", 22, 1, ReadWrite, "SPI Master">,
      Field<"LSBFRST", 23, 1, ReadWrite, "Data frame format">,
      Field<"CPHA", 24, 1, ReadWrite, "Clock phase">,
      Field<"CPOL", 25, 1, ReadWrite, "Clock polarity">,
      Field<"SSM", 26, 1, ReadWrite, "Software management of SS signal input">,
      Field<"SSIOP", 28, 1, ReadWrite, "SS input/output polarity">,
      Field<"SSOE", 29, 1, ReadWrite, "SS output enable">,
      Field<"SSOM", 30, 1, ReadWrite, "SS output management in master mode">,
      Field<"AFCNTR", 31, 1, ReadWrite, "Alternate function GPIOs control">,
    ], "configuration register 2">,
    Register<"IER", 0x10, 32, [
      Field<"RXPIE", 0, 1, ReadWrite, "RXP Interrupt Enable">,
      Field<"TXPIE", 1, 1, Read, "TXP interrupt enable">,
      Field<"DPXPIE", 2, 1, Read, "DXP interrupt enabled">,
      Field<"EOTIE", 3, 1, ReadWrite, "EOT, SUSP and TXC interrupt enable">,
      Field<"TXTFIE", 4, 1, ReadWrite, "TXTFIE interrupt enable">,
      Field<"UDRIE", 5, 1, ReadWrite, "UDR interrupt enable">,
      Field<"OVRIE", 6, 1, ReadWrite, "OVR interrupt enable">,
      Field<"CRCEIE", 7, 1, ReadWrite, "CRC Interrupt enable">,
      Field<"TIFREIE", 8, 1, ReadWrite, "TIFRE interrupt enable">,
      Field<"MODFIE", 9, 1, ReadWrite, "Mode Fault interrupt enable">,
      Field<"TSERFIE", 10, 1, ReadWrite, "Additional number of transactions reload interrupt enable">,
    ], "Interrupt Enable Register">,
    Register<"SR", 0x14, 32, [
      Field<"RXP", 0, 1, ReadWrite, "Rx-Packet available">,
      Field<"TXP", 1, 1, ReadWrite, "Tx-Packet space available">,
      Field<"DXP", 2, 1, ReadWrite, "Duplex Packet">,
      Field<"EOT", 3, 1, ReadWrite, "End Of Transfer">,
      Field<"TXTF", 4, 1, ReadWrite, "Transmission Transfer Filled">,
      Field<"UDR", 5, 1, ReadWrite, "Underrun at slave transmission mode">,
      Field<"OVR", 6, 1, ReadWrite, "Overrun">,
      Field<"CRCE", 7, 1, ReadWrite, "CRC Error">,
      Field<"TIFRE", 8, 1, ReadWrite, "TI frame format error">,
      Field<"MODF", 9, 1, ReadWrite, "Mode Fault">,
      Field<"TSERF", 10, 1, ReadWrite, "Additional number of SPI data to be transacted was reload">,
      Field<"SUSP", 11, 1, ReadWrite, "SUSPend">,
      Field<"TXC", 12, 1, ReadWrite, "TxFIFO transmission complete">,
      Field<"RXPLVL", 13, 2, ReadWrite, "RxFIFO Packing LeVeL">,
      Field<"RXWNE", 15, 1, ReadWrite, "RxFIFO Word Not Empty">,
      Field<"CTSIZE", 16, 16, ReadWrite, "Number of data frames remaining in current TSIZE session">,
    ], "Status Register">,
    Register<"IFCR", 0x18, 32, [
      Field<"EOTC", 3, 1, ReadWrite, "End Of Transfer flag clear">,
      Field<"TXTFC", 4, 1, ReadWrite, "Transmission Transfer Filled flag clear">,
      Field<"UDRC", 5, 1, ReadWrite, "Underrun flag clear">,
      Field<"OVRC", 6, 1, ReadWrite, "Overrun flag clear">,
      Field<"CRCEC", 7, 1, ReadWrite, "CRC Error flag clear">,
      Field<"TIFREC", 8, 1, ReadWrite, "TI frame format error flag clear">,
      Field<"MODFC", 9, 1, ReadWrite, "Mode Fault flag clear">,
      Field<"TSERFC", 10, 1, ReadWrite, "TSERFC flag clear">,
      Field<"SUSPC", 11, 1, ReadWrite, "SUSPend flag clear">,
    ], "Interrupt/Status Flags Clear Register">,
    Register<"TXDR", 0x20, 32, [
      Field<"TXDR", 0, 32, ReadWrite, "Transmit data register">,
    ], "Transmit Data Register">,
    Register<"RXDR", 0x30, 32, [
      Field<"RXDR", 0, 32, ReadWrite, "Receive data register">,
    ], "Receive Data Register">,
    Register<"CRCPOLY", 0x40, 32, [
      Field<"CRCPOLY", 0, 32, ReadWrite, "CRC polynomial register">,
    ], "Polynomial Register">,
    Register<"TXCRC", 0x44, 32, [
      Field<"TXCRC", 0, 32, ReadWrite, "CRC register for transmitter">,
    ], "Transmitter CRC Register">,
    Register<"RXCRC", 0x48, 32, [
      Field<"RXCRC", 0, 32, ReadWrite, "CRC register for receiver">,
    ], "Receiver CRC Register">,
    Register<"UDRDR", 0x4c, 32, [
      Field<"UDRDR", 0, 32, ReadWrite, "Data at slave underrun condition">,
    ], "Underrun Data Register">,
    Register<"CGFR", 0x50, 32, [
      Field<"I2SMOD", 0, 1, ReadWrite, "I2S mode selection">,
      Field<"I2SCFG", 1, 3, ReadWrite, "I2S configuration mode">,
      Field<"I2SSTD", 4, 2, ReadWrite, "I2S standard selection">,
      Field<"PCMSYNC", 7, 1, ReadWrite, "PCM frame synchronization">,
      Field<"DATLEN", 8, 2, ReadWrite, "Data length to be transferred">,
      Field<"CHLEN", 10, 1, ReadWrite, "Channel length (number of bits per audio channel)">,
      Field<"CKPOL", 11, 1, ReadWrite, "Serial audio clock polarity">,
      Field<"FIXCH", 12, 1, ReadWrite, "Word select inversion">,
      Field<"WSINV", 13, 1, ReadWrite, "Fixed channel length in SLAVE">,
      Field<"DATFMT", 14, 1, ReadWrite, "Data format">,
      Field<"I2SDIV", 16, 8, ReadWrite, "I2S linear prescaler">,
      Field<"ODD", 24, 1, ReadWrite, "Odd factor for the prescaler">,
      Field<"MCKOE", 25, 1, ReadWrite, "Master clock output enable">,
    ], "configuration register">,
  ];
}

class SPIInstance<string Name, int Base> : PeripheralInstance<Name, Base, SPIPeripheral>;
def SPI2 : SPIInstance<"SPI2", 0x40003800>;
def SPI3 : SPIInstance<"SPI3", 0x40003c00>;
def SPI1 : SPIInstance<"SPI1", 0x40013000>;
def SPI4 : SPIInstance<"SPI4", 0x40013400>;
def SPI5 : SPIInstance<"SPI5", 0x40015000>;
def SPI6 : SPIInstance<"SPI6", 0x58001400>;

def SPIGroup : PeripheralGroup<[ SPI1, SPI2, SPI3, SPI4, SPI5, SPI6 ]>;

#endif // _PERIPHERALS_SPI_TD
