#ifndef _PERIPHERALS_QUADSPI_TD
#define _PERIPHERALS_QUADSPI_TD

include "base.td"

def QUADSPIPeripheral : PeripheralType<"QUADSPI", "QUADSPI"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"EN", 0, 1, ReadWrite, "Enable Enable the QUADSPI.">,
      Field<"ABORT", 1, 1, ReadWrite, "Abort request This bit aborts the on-going command sequence. It is automatically reset once the abort is complete. This bit stops the current transfer. In polling mode or memory-mapped mode, this bit also reset the APM bit or the DM bit.">,
      Field<"DMAEN", 2, 1, ReadWrite, "DMA enable In indirect mode, DMA can be used to input or output data via the QUADSPI_DR register. DMA transfers are initiated when the FIFO threshold flag, FTF, is set.">,
      Field<"TCEN", 3, 1, ReadWrite, "Timeout counter enable This bit is valid only when memory-mapped mode (FMODE = 11) is selected. Activating this bit causes the chip select (nCS) to be released (and thus reduces consumption) if there has not been an access after a certain amount of time, where this time is defined by TIMEOUT[15:0] (QUADSPI_LPTR). Enable the timeout counter. By default, the QUADSPI never stops its prefetch operation, keeping the previous read operation active with nCS maintained low, even if no access to the Flash memory occurs for a long time. Since Flash memories tend to consume more when nCS is held low, the application might want to activate the timeout counter (TCEN = 1, QUADSPI_CR[3]) so that nCS is released after a period of TIMEOUT[15:0] (QUADSPI_LPTR) cycles have elapsed without an access since when the FIFO becomes full with prefetch data. This bit can be modified only when BUSY = 0.">,
      Field<"SSHIFT", 4, 1, ReadWrite, "Sample shift By default, the QUADSPI samples data 1/2 of a CLK cycle after the data is driven by the Flash memory. This bit allows the data is to be sampled later in order to account for external signal delays. Firmware must assure that SSHIFT = 0 when in DDR mode (when DDRM = 1). This field can be modified only when BUSY = 0.">,
      Field<"DFM", 6, 1, ReadWrite, "Dual-flash mode This bit activates dual-flash mode, where two external Flash memories are used simultaneously to double throughput and capacity. This bit can be modified only when BUSY = 0.">,
      Field<"FSEL", 7, 1, ReadWrite, "Flash memory selection This bit selects the Flash memory to be addressed in single flash mode (when DFM = 0). This bit can be modified only when BUSY = 0. This bit is ignored when DFM = 1.">,
      Field<"FTHRES", 8, 5, ReadWrite, "FIFO threshold level Defines, in indirect mode, the threshold number of bytes in the FIFO that will cause the FIFO threshold flag (FTF, QUADSPI_SR[2]) to be set. In indirect write mode (FMODE = 00): ... In indirect read mode (FMODE = 01): ... If DMAEN = 1, then the DMA controller for the corresponding channel must be disabled before changing the FTHRES value.">,
      Field<"TEIE", 16, 1, ReadWrite, "Transfer error interrupt enable This bit enables the transfer error interrupt.">,
      Field<"TCIE", 17, 1, ReadWrite, "Transfer complete interrupt enable This bit enables the transfer complete interrupt.">,
      Field<"FTIE", 18, 1, ReadWrite, "FIFO threshold interrupt enable This bit enables the FIFO threshold interrupt.">,
      Field<"SMIE", 19, 1, ReadWrite, "Status match interrupt enable This bit enables the status match interrupt.">,
      Field<"TOIE", 20, 1, ReadWrite, "TimeOut interrupt enable This bit enables the TimeOut interrupt.">,
      Field<"APMS", 22, 1, ReadWrite, "Automatic poll mode stop This bit determines if automatic polling is stopped after a match. This bit can be modified only when BUSY = 0.">,
      Field<"PMM", 23, 1, ReadWrite, "Polling match mode This bit indicates which method should be used for determining a match during automatic polling mode. This bit can be modified only when BUSY = 0.">,
      Field<"PRESCALER", 24, 8, ReadWrite, "clock prescaler">,
    ], "QUADSPI control register">,
    Register<"DCR", 0x4, 32, [
      Field<"CKMODE", 0, 1, ReadWrite, "indicates the level that clk takes between command">,
      Field<"CSHT", 8, 3, ReadWrite, "Chip select high time CSHT+1 defines the minimum number of CLK cycles which the chip select (nCS) must remain high between commands issued to the Flash memory. ... This field can be modified only when BUSY = 0.">,
      Field<"FSIZE", 16, 5, ReadWrite, "Flash memory size This field defines the size of external memory using the following formula: Number of bytes in Flash memory = 2[FSIZE+1] FSIZE+1 is effectively the number of address bits required to address the Flash memory. The Flash memory capacity can be up to 4GB (addressed using 32 bits) in indirect mode, but the addressable space in memory-mapped mode is limited to 256MB. If DFM = 1, FSIZE indicates the total capacity of the two Flash memories together. This field can be modified only when BUSY = 0.">,
    ], "QUADSPI device configuration register">,
    Register<"SR", 0x8, 32, [
      Field<"TEF", 0, 1, ReadWrite, "Transfer error flag This bit is set in indirect mode when an invalid address is being accessed in indirect mode. It is cleared by writing 1 to CTEF.">,
      Field<"TCF", 1, 1, ReadWrite, "Transfer complete flag This bit is set in indirect mode when the programmed number of data has been transferred or in any mode when the transfer has been aborted.It is cleared by writing 1 to CTCF.">,
      Field<"FTF", 2, 1, ReadWrite, "FIFO threshold flag In indirect mode, this bit is set when the FIFO threshold has been reached, or if there is any data left in the FIFO after reads from the Flash memory are complete. It is cleared automatically as soon as threshold condition is no longer true. In automatic polling mode this bit is set every time the status register is read, and the bit is cleared when the data register is read.">,
      Field<"SMF", 3, 1, ReadWrite, "Status match flag This bit is set in automatic polling mode when the unmasked received data matches the corresponding bits in the match register (QUADSPI_PSMAR). It is cleared by writing 1 to CSMF.">,
      Field<"TOF", 4, 1, ReadWrite, "Timeout flag This bit is set when timeout occurs. It is cleared by writing 1 to CTOF.">,
      Field<"BUSY", 5, 1, ReadWrite, "Busy This bit is set when an operation is on going. This bit clears automatically when the operation with the Flash memory is finished and the FIFO is empty.">,
      Field<"FLEVEL", 8, 6, ReadWrite, "FIFO level This field gives the number of valid bytes which are being held in the FIFO. FLEVEL = 0 when the FIFO is empty, and 16 when it is full. In memory-mapped mode and in automatic status polling mode, FLEVEL is zero.">,
    ], "QUADSPI status register">,
    Register<"FCR", 0xc, 32, [
      Field<"CTEF", 0, 1, ReadWrite, "Clear transfer error flag Writing 1 clears the TEF flag in the QUADSPI_SR register">,
      Field<"CTCF", 1, 1, ReadWrite, "Clear transfer complete flag Writing 1 clears the TCF flag in the QUADSPI_SR register">,
      Field<"CSMF", 3, 1, ReadWrite, "Clear status match flag Writing 1 clears the SMF flag in the QUADSPI_SR register">,
      Field<"CTOF", 4, 1, ReadWrite, "Clear timeout flag Writing 1 clears the TOF flag in the QUADSPI_SR register">,
    ], "QUADSPI flag clear register">,
    Register<"DLR", 0x10, 32, [
      Field<"DL", 0, 32, ReadWrite, "Data length Number of data to be retrieved (value+1) in indirect and status-polling modes. A value no greater than 3 (indicating 4 bytes) should be used for status-polling mode. All 1s in indirect mode means undefined length, where QUADSPI will continue until the end of memory, as defined by FSIZE. 0x0000_0000: 1 byte is to be transferred 0x0000_0001: 2 bytes are to be transferred 0x0000_0002: 3 bytes are to be transferred 0x0000_0003: 4 bytes are to be transferred ... 0xFFFF_FFFD: 4,294,967,294 (4G-2) bytes are to be transferred 0xFFFF_FFFE: 4,294,967,295 (4G-1) bytes are to be transferred 0xFFFF_FFFF: undefined length -- all bytes until the end of Flash memory (as defined by FSIZE) are to be transferred. Continue reading indefinitely if FSIZE = 0x1F. DL[0] is stuck at 1 in dual-flash mode (DFM = 1) even when 0 is written to this bit, thus assuring that each access transfers an even number of bytes. This field has no effect when in memory-mapped mode (FMODE = 10). This field can be written only when BUSY = 0.">,
    ], "QUADSPI data length register">,
    Register<"CCR", 0x14, 32, [
      Field<"INSTRUCTION", 0, 8, ReadWrite, "Instruction Instruction to be send to the external SPI device. This field can be written only when BUSY = 0.">,
      Field<"IMODE", 8, 2, ReadWrite, "Instruction mode This field defines the instruction phase mode of operation: This field can be written only when BUSY = 0.">,
      Field<"ADMODE", 10, 2, ReadWrite, "Address mode This field defines the address phase mode of operation: This field can be written only when BUSY = 0.">,
      Field<"ADSIZE", 12, 2, ReadWrite, "Address size This bit defines address size: This field can be written only when BUSY = 0.">,
      Field<"ABMODE", 14, 2, ReadWrite, "Alternate bytes mode This field defines the alternate-bytes phase mode of operation: This field can be written only when BUSY = 0.">,
      Field<"ABSIZE", 16, 2, ReadWrite, "Alternate bytes size This bit defines alternate bytes size: This field can be written only when BUSY = 0.">,
      Field<"DCYC", 18, 5, ReadWrite, "Number of dummy cycles This field defines the duration of the dummy phase. In both SDR and DDR modes, it specifies a number of CLK cycles (0-31). This field can be written only when BUSY = 0.">,
      Field<"DMODE", 24, 2, ReadWrite, "Data mode This field defines the data phases mode of operation: This field also determines the dummy phase mode of operation. This field can be written only when BUSY = 0.">,
      Field<"FMODE", 26, 2, ReadWrite, "Functional mode This field defines the QUADSPI functional mode of operation. If DMAEN = 1 already, then the DMA controller for the corresponding channel must be disabled before changing the FMODE value. This field can be written only when BUSY = 0.">,
      Field<"SIOO", 28, 1, ReadWrite, "Send instruction only once mode See Section15.3.11: Sending the instruction only once on page13. This bit has no effect when IMODE = 00. This field can be written only when BUSY = 0.">,
      Field<"DHHC", 30, 1, ReadWrite, "DDR hold Delay the data output by 1/4 of the QUADSPI output clock cycle in DDR mode: This feature is only active in DDR mode. This field can be written only when BUSY = 0.">,
      Field<"DDRM", 31, 1, ReadWrite, "Double data rate mode This bit sets the DDR mode for the address, alternate byte and data phase: This field can be written only when BUSY = 0.">,
    ], "QUADSPI communication configuration register">,
    Register<"AR", 0x18, 32, [
      Field<"ADDRESS", 0, 32, ReadWrite, "[31 0]: Address Address to be send to the external Flash memory Writes to this field are ignored when BUSY = 0 or when FMODE = 11 (memory-mapped mode). In dual flash mode, ADDRESS[0] is automatically stuck to 0 as the address should always be even">,
    ], "QUADSPI address register">,
    Register<"ABR", 0x1c, 32, [
      Field<"ALTERNATE", 0, 32, ReadWrite, "Alternate Bytes Optional data to be send to the external SPI device right after the address. This field can be written only when BUSY = 0.">,
    ], "QUADSPI alternate bytes registers">,
    Register<"DR", 0x20, 32, [
      Field<"DATA", 0, 32, ReadWrite, "Data Data to be sent/received to/from the external SPI device. In indirect write mode, data written to this register is stored on the FIFO before it is sent to the Flash memory during the data phase. If the FIFO is too full, a write operation is stalled until the FIFO has enough space to accept the amount of data being written. In indirect read mode, reading this register gives (via the FIFO) the data which was received from the Flash memory. If the FIFO does not have as many bytes as requested by the read operation and if BUSY=1, the read operation is stalled until enough data is present or until the transfer is complete, whichever happens first. In automatic polling mode, this register contains the last data read from the Flash memory (without masking). Word, halfword, and byte accesses to this register are supported. In indirect write mode, a byte write adds 1 byte to the FIFO, a halfword write 2, and a word write 4. Similarly, in indirect read mode, a byte read removes 1 byte from the FIFO, a halfword read 2, and a word read 4. Accesses in indirect mode must be aligned to the bottom of this register: a byte read must read DATA[7:0] and a halfword read must read DATA[15:0].">,
    ], "QUADSPI data register">,
    Register<"PSMKR", 0x24, 32, [
      Field<"MASK", 0, 32, ReadWrite, "Status mask Mask to be applied to the status bytes received in polling mode. For bit n: This field can be written only when BUSY = 0.">,
    ], "QUADSPI polling status mask register">,
    Register<"PSMAR", 0x28, 32, [
      Field<"MATCH", 0, 32, ReadWrite, "Status match Value to be compared with the masked status register to get a match. This field can be written only when BUSY = 0.">,
    ], "QUADSPI polling status match register">,
    Register<"PIR", 0x2c, 32, [
      Field<"INTERVAL", 0, 16, ReadWrite, "Polling interval Number of CLK cycles between to read during automatic polling phases. This field can be written only when BUSY = 0.">,
    ], "QUADSPI polling interval register">,
    Register<"LPTR", 0x30, 32, [
      Field<"TIMEOUT", 0, 16, ReadWrite, "Timeout period After each access in memory-mapped mode, the QUADSPI prefetches the subsequent bytes and holds these bytes in the FIFO. This field indicates how many CLK cycles the QUADSPI waits after the FIFO becomes full until it raises nCS, putting the Flash memory in a lower-consumption state. This field can be written only when BUSY = 0.">,
    ], "QUADSPI low-power timeout register">,
  ];
}

class QUADSPIInstance<string Name, int Base> : PeripheralInstance<Name, Base, QUADSPIPeripheral>;
def QUADSPI : QUADSPIInstance<"QUADSPI", 0x52005000>;

#endif // _PERIPHERALS_QUADSPI_TD
