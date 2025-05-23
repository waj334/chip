#ifndef _PERIPHERALS_LPUART1_TD
#define _PERIPHERALS_LPUART1_TD

include "base.td"

def LPUART1Peripheral : PeripheralType<"LPUART1", "LPUART1"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR1", 0x0, 32, [
      Field<"RXFFIE", 31, 1, ReadWrite, "RXFIFO Full interrupt enable">,
      Field<"TXFEIE", 30, 1, ReadWrite, "TXFIFO empty interrupt enable">,
      Field<"FIFOEN", 29, 1, ReadWrite, "FIFO mode enable">,
      Field<"M1", 28, 1, ReadWrite, "Word length">,
      Field<"DEAT", 21, 5, ReadWrite, "Driver Enable assertion time">,
      Field<"DEDT", 16, 5, ReadWrite, "Driver Enable deassertion time">,
      Field<"CMIE", 14, 1, ReadWrite, "Character match interrupt enable">,
      Field<"MME", 13, 1, ReadWrite, "Mute mode enable">,
      Field<"M0", 12, 1, ReadWrite, "Word length">,
      Field<"WAKE", 11, 1, ReadWrite, "Receiver wakeup method">,
      Field<"PCE", 10, 1, ReadWrite, "Parity control enable">,
      Field<"PS", 9, 1, ReadWrite, "Parity selection">,
      Field<"PEIE", 8, 1, ReadWrite, "PE interrupt enable">,
      Field<"TXEIE", 7, 1, ReadWrite, "interrupt enable">,
      Field<"TCIE", 6, 1, ReadWrite, "Transmission complete interrupt enable">,
      Field<"RXNEIE", 5, 1, ReadWrite, "RXNE interrupt enable">,
      Field<"IDLEIE", 4, 1, ReadWrite, "IDLE interrupt enable">,
      Field<"TE", 3, 1, ReadWrite, "Transmitter enable">,
      Field<"RE", 2, 1, ReadWrite, "Receiver enable">,
      Field<"UESM", 1, 1, ReadWrite, "USART enable in Stop mode">,
      Field<"UE", 0, 1, ReadWrite, "USART enable">,
    ], "Control register 1">,
    Register<"CR2", 0x4, 32, [
      Field<"ADD", 24, 8, ReadWrite, "Address of the USART node">,
      Field<"MSBFIRST", 19, 1, ReadWrite, "Most significant bit first">,
      Field<"DATAINV", 18, 1, ReadWrite, "Binary data inversion">,
      Field<"TXINV", 17, 1, ReadWrite, "TX pin active level inversion">,
      Field<"RXINV", 16, 1, ReadWrite, "RX pin active level inversion">,
      Field<"SWAP", 15, 1, ReadWrite, "Swap TX/RX pins">,
      Field<"STOP", 12, 2, ReadWrite, "STOP bits">,
      Field<"ADDM7", 4, 1, ReadWrite, "7-bit Address Detection/4-bit Address Detection">,
    ], "Control register 2">,
    Register<"CR3", 0x8, 32, [
      Field<"TXFTCFG", 29, 3, ReadWrite, "TXFIFO threshold configuration">,
      Field<"RXFTIE", 28, 1, ReadWrite, "RXFIFO threshold interrupt enable">,
      Field<"RXFTCFG", 25, 3, ReadWrite, "Receive FIFO threshold configuration">,
      Field<"TXFTIE", 23, 1, ReadWrite, "TXFIFO threshold interrupt enable">,
      Field<"WUFIE", 22, 1, ReadWrite, "Wakeup from Stop mode interrupt enable">,
      Field<"WUS", 20, 2, ReadWrite, "Wakeup from Stop mode interrupt flag selection">,
      Field<"DEP", 15, 1, ReadWrite, "Driver enable polarity selection">,
      Field<"DEM", 14, 1, ReadWrite, "Driver enable mode">,
      Field<"DDRE", 13, 1, ReadWrite, "DMA Disable on Reception Error">,
      Field<"OVRDIS", 12, 1, ReadWrite, "Overrun Disable">,
      Field<"CTSIE", 10, 1, ReadWrite, "CTS interrupt enable">,
      Field<"CTSE", 9, 1, ReadWrite, "CTS enable">,
      Field<"RTSE", 8, 1, ReadWrite, "RTS enable">,
      Field<"DMAT", 7, 1, ReadWrite, "DMA enable transmitter">,
      Field<"DMAR", 6, 1, ReadWrite, "DMA enable receiver">,
      Field<"HDSEL", 3, 1, ReadWrite, "Half-duplex selection">,
      Field<"EIE", 0, 1, ReadWrite, "Error interrupt enable">,
    ], "Control register 3">,
    Register<"BRR", 0xc, 32, [
      Field<"BRR", 0, 20, ReadWrite, "BRR">,
    ], "Baud rate register">,
    Register<"GTPR", 0x10, 32, [
      Field<"GT", 8, 8, ReadWrite, "Guard time value">,
      Field<"PSC", 0, 8, ReadWrite, "Prescaler value">,
    ], "Guard time and prescaler register">,
    Register<"RTOR", 0x14, 32, [
      Field<"BLEN", 24, 8, ReadWrite, "Block Length">,
      Field<"RTO", 0, 24, ReadWrite, "Receiver timeout value">,
    ], "Receiver timeout register">,
    Register<"RQR", 0x18, 32, [
      Field<"TXFRQ", 4, 1, ReadWrite, "Transmit data flush request">,
      Field<"RXFRQ", 3, 1, ReadWrite, "Receive data flush request">,
      Field<"MMRQ", 2, 1, ReadWrite, "Mute mode request">,
      Field<"SBKRQ", 1, 1, ReadWrite, "Send break request">,
      Field<"ABRRQ", 0, 1, ReadWrite, "Auto baud rate request">,
    ], "Request register">,
    Register<"ISR", 0x1c, 32, [
      Field<"TXFT", 27, 1, ReadWrite, "TXFIFO threshold flag">,
      Field<"RXFT", 26, 1, ReadWrite, "RXFIFO threshold flag">,
      Field<"RXFF", 24, 1, ReadWrite, "RXFIFO Full">,
      Field<"TXFE", 23, 1, ReadWrite, "TXFIFO Empty">,
      Field<"REACK", 22, 1, ReadWrite, "REACK">,
      Field<"TEACK", 21, 1, ReadWrite, "TEACK">,
      Field<"WUF", 20, 1, ReadWrite, "WUF">,
      Field<"RWU", 19, 1, ReadWrite, "RWU">,
      Field<"SBKF", 18, 1, ReadWrite, "SBKF">,
      Field<"CMF", 17, 1, ReadWrite, "CMF">,
      Field<"BUSY", 16, 1, ReadWrite, "BUSY">,
      Field<"CTS", 10, 1, ReadWrite, "CTS">,
      Field<"CTSIF", 9, 1, ReadWrite, "CTSIF">,
      Field<"TXE", 7, 1, ReadWrite, "TXE">,
      Field<"TC", 6, 1, ReadWrite, "TC">,
      Field<"RXNE", 5, 1, ReadWrite, "RXNE">,
      Field<"IDLE", 4, 1, ReadWrite, "IDLE">,
      Field<"ORE", 3, 1, ReadWrite, "ORE">,
      Field<"NE", 2, 1, ReadWrite, "NE">,
      Field<"FE", 1, 1, ReadWrite, "FE">,
      Field<"PE", 0, 1, ReadWrite, "PE">,
    ], "Interrupt & status register">,
    Register<"ICR", 0x20, 32, [
      Field<"WUCF", 20, 1, ReadWrite, "Wakeup from Stop mode clear flag">,
      Field<"CMCF", 17, 1, ReadWrite, "Character match clear flag">,
      Field<"CTSCF", 9, 1, ReadWrite, "CTS clear flag">,
      Field<"TCCF", 6, 1, ReadWrite, "Transmission complete clear flag">,
      Field<"IDLECF", 4, 1, ReadWrite, "Idle line detected clear flag">,
      Field<"ORECF", 3, 1, ReadWrite, "Overrun error clear flag">,
      Field<"NCF", 2, 1, ReadWrite, "Noise detected clear flag">,
      Field<"FECF", 1, 1, ReadWrite, "Framing error clear flag">,
      Field<"PECF", 0, 1, ReadWrite, "Parity error clear flag">,
    ], "Interrupt flag clear register">,
    Register<"RDR", 0x24, 32, [
      Field<"RDR", 0, 9, ReadWrite, "Receive data value">,
    ], "Receive data register">,
    Register<"TDR", 0x28, 32, [
      Field<"TDR", 0, 9, ReadWrite, "Transmit data value">,
    ], "Transmit data register">,
    Register<"PRESC", 0x2c, 32, [
      Field<"PRESCALER", 0, 4, ReadWrite, "Clock prescaler">,
    ], "Prescaler register">,
  ];
}

class LPUART1Instance<string Name, int Base> : PeripheralInstance<Name, Base, LPUART1Peripheral>;
def LPUART1 : LPUART1Instance<"LPUART1", 0x58000c00>;

#endif // _PERIPHERALS_LPUART1_TD
