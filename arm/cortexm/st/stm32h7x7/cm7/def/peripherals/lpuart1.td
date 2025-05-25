#ifndef _PERIPHERALS_LPUART1_TD
#define _PERIPHERALS_LPUART1_TD

include "base.td"

def LPUART1Peripheral : PeripheralType<"LPUART1", "LPUART1"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR1", 0x0, 32, [
      Field<"UE", 0, 1, ReadWrite, "USART enable">,
      Field<"UESM", 1, 1, ReadWrite, "USART enable in Stop mode">,
      Field<"RE", 2, 1, ReadWrite, "Receiver enable">,
      Field<"TE", 3, 1, ReadWrite, "Transmitter enable">,
      Field<"IDLEIE", 4, 1, ReadWrite, "IDLE interrupt enable">,
      Field<"RXNEIE", 5, 1, ReadWrite, "RXNE interrupt enable">,
      Field<"TCIE", 6, 1, ReadWrite, "Transmission complete interrupt enable">,
      Field<"TXEIE", 7, 1, ReadWrite, "interrupt enable">,
      Field<"PEIE", 8, 1, ReadWrite, "PE interrupt enable">,
      Field<"PS", 9, 1, ReadWrite, "Parity selection">,
      Field<"PCE", 10, 1, ReadWrite, "Parity control enable">,
      Field<"WAKE", 11, 1, ReadWrite, "Receiver wakeup method">,
      Field<"M0", 12, 1, ReadWrite, "Word length">,
      Field<"MME", 13, 1, ReadWrite, "Mute mode enable">,
      Field<"CMIE", 14, 1, ReadWrite, "Character match interrupt enable">,
      Field<"DEDT", 16, 5, ReadWrite, "Driver Enable deassertion time">,
      Field<"DEAT", 21, 5, ReadWrite, "Driver Enable assertion time">,
      Field<"M1", 28, 1, ReadWrite, "Word length">,
      Field<"FIFOEN", 29, 1, ReadWrite, "FIFO mode enable">,
      Field<"TXFEIE", 30, 1, ReadWrite, "TXFIFO empty interrupt enable">,
      Field<"RXFFIE", 31, 1, ReadWrite, "RXFIFO Full interrupt enable">,
    ], "Control register 1">,
    Register<"CR2", 0x4, 32, [
      Field<"ADDM7", 4, 1, ReadWrite, "7-bit Address Detection/4-bit Address Detection">,
      Field<"STOP", 12, 2, ReadWrite, "STOP bits">,
      Field<"SWAP", 15, 1, ReadWrite, "Swap TX/RX pins">,
      Field<"RXINV", 16, 1, ReadWrite, "RX pin active level inversion">,
      Field<"TXINV", 17, 1, ReadWrite, "TX pin active level inversion">,
      Field<"DATAINV", 18, 1, ReadWrite, "Binary data inversion">,
      Field<"MSBFIRST", 19, 1, ReadWrite, "Most significant bit first">,
      Field<"ADD", 24, 8, ReadWrite, "Address of the USART node">,
    ], "Control register 2">,
    Register<"CR3", 0x8, 32, [
      Field<"EIE", 0, 1, ReadWrite, "Error interrupt enable">,
      Field<"HDSEL", 3, 1, ReadWrite, "Half-duplex selection">,
      Field<"DMAR", 6, 1, ReadWrite, "DMA enable receiver">,
      Field<"DMAT", 7, 1, ReadWrite, "DMA enable transmitter">,
      Field<"RTSE", 8, 1, ReadWrite, "RTS enable">,
      Field<"CTSE", 9, 1, ReadWrite, "CTS enable">,
      Field<"CTSIE", 10, 1, ReadWrite, "CTS interrupt enable">,
      Field<"OVRDIS", 12, 1, ReadWrite, "Overrun Disable">,
      Field<"DDRE", 13, 1, ReadWrite, "DMA Disable on Reception Error">,
      Field<"DEM", 14, 1, ReadWrite, "Driver enable mode">,
      Field<"DEP", 15, 1, ReadWrite, "Driver enable polarity selection">,
      Field<"WUS", 20, 2, ReadWrite, "Wakeup from Stop mode interrupt flag selection">,
      Field<"WUFIE", 22, 1, ReadWrite, "Wakeup from Stop mode interrupt enable">,
      Field<"TXFTIE", 23, 1, ReadWrite, "TXFIFO threshold interrupt enable">,
      Field<"RXFTCFG", 25, 3, ReadWrite, "Receive FIFO threshold configuration">,
      Field<"RXFTIE", 28, 1, ReadWrite, "RXFIFO threshold interrupt enable">,
      Field<"TXFTCFG", 29, 3, ReadWrite, "TXFIFO threshold configuration">,
    ], "Control register 3">,
    Register<"BRR", 0xc, 32, [
      Field<"BRR", 0, 20, ReadWrite, "BRR">,
    ], "Baud rate register">,
    Register<"GTPR", 0x10, 32, [
      Field<"PSC", 0, 8, ReadWrite, "Prescaler value">,
      Field<"GT", 8, 8, ReadWrite, "Guard time value">,
    ], "Guard time and prescaler register">,
    Register<"RTOR", 0x14, 32, [
      Field<"RTO", 0, 24, ReadWrite, "Receiver timeout value">,
      Field<"BLEN", 24, 8, ReadWrite, "Block Length">,
    ], "Receiver timeout register">,
    Register<"RQR", 0x18, 32, [
      Field<"ABRRQ", 0, 1, ReadWrite, "Auto baud rate request">,
      Field<"SBKRQ", 1, 1, ReadWrite, "Send break request">,
      Field<"MMRQ", 2, 1, ReadWrite, "Mute mode request">,
      Field<"RXFRQ", 3, 1, ReadWrite, "Receive data flush request">,
      Field<"TXFRQ", 4, 1, ReadWrite, "Transmit data flush request">,
    ], "Request register">,
    Register<"ISR", 0x1c, 32, [
      Field<"PE", 0, 1, ReadWrite, "PE">,
      Field<"FE", 1, 1, ReadWrite, "FE">,
      Field<"NE", 2, 1, ReadWrite, "NE">,
      Field<"ORE", 3, 1, ReadWrite, "ORE">,
      Field<"IDLE", 4, 1, ReadWrite, "IDLE">,
      Field<"RXNE", 5, 1, ReadWrite, "RXNE">,
      Field<"TC", 6, 1, ReadWrite, "TC">,
      Field<"TXE", 7, 1, ReadWrite, "TXE">,
      Field<"CTSIF", 9, 1, ReadWrite, "CTSIF">,
      Field<"CTS", 10, 1, ReadWrite, "CTS">,
      Field<"BUSY", 16, 1, ReadWrite, "BUSY">,
      Field<"CMF", 17, 1, ReadWrite, "CMF">,
      Field<"SBKF", 18, 1, ReadWrite, "SBKF">,
      Field<"RWU", 19, 1, ReadWrite, "RWU">,
      Field<"WUF", 20, 1, ReadWrite, "WUF">,
      Field<"TEACK", 21, 1, ReadWrite, "TEACK">,
      Field<"REACK", 22, 1, ReadWrite, "REACK">,
      Field<"TXFE", 23, 1, ReadWrite, "TXFIFO Empty">,
      Field<"RXFF", 24, 1, ReadWrite, "RXFIFO Full">,
      Field<"RXFT", 26, 1, ReadWrite, "RXFIFO threshold flag">,
      Field<"TXFT", 27, 1, ReadWrite, "TXFIFO threshold flag">,
    ], "Interrupt & status register">,
    Register<"ICR", 0x20, 32, [
      Field<"PECF", 0, 1, ReadWrite, "Parity error clear flag">,
      Field<"FECF", 1, 1, ReadWrite, "Framing error clear flag">,
      Field<"NCF", 2, 1, ReadWrite, "Noise detected clear flag">,
      Field<"ORECF", 3, 1, ReadWrite, "Overrun error clear flag">,
      Field<"IDLECF", 4, 1, ReadWrite, "Idle line detected clear flag">,
      Field<"TCCF", 6, 1, ReadWrite, "Transmission complete clear flag">,
      Field<"CTSCF", 9, 1, ReadWrite, "CTS clear flag">,
      Field<"CMCF", 17, 1, ReadWrite, "Character match clear flag">,
      Field<"WUCF", 20, 1, ReadWrite, "Wakeup from Stop mode clear flag">,
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
