#ifndef _PERIPHERALS_USART_TD
#define _PERIPHERALS_USART_TD

include "base.td"

def USARTPeripheral : PeripheralType<"USART", "Universal synchronous asynchronous receiver transmitter"> {
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
      Field<"OVER8", 15, 1, ReadWrite, "Oversampling mode">,
      Field<"DEDT0", 16, 1, ReadWrite, "DEDT0">,
      Field<"DEDT1", 17, 1, ReadWrite, "DEDT1">,
      Field<"DEDT2", 18, 1, ReadWrite, "DEDT2">,
      Field<"DEDT3", 19, 1, ReadWrite, "DEDT3">,
      Field<"DEDT4", 20, 1, ReadWrite, "Driver Enable de-assertion time">,
      Field<"DEAT0", 21, 1, ReadWrite, "DEAT0">,
      Field<"DEAT1", 22, 1, ReadWrite, "DEAT1">,
      Field<"DEAT2", 23, 1, ReadWrite, "DEAT2">,
      Field<"DEAT3", 24, 1, ReadWrite, "DEAT3">,
      Field<"DEAT4", 25, 1, ReadWrite, "Driver Enable assertion time">,
      Field<"RTOIE", 26, 1, ReadWrite, "Receiver timeout interrupt enable">,
      Field<"EOBIE", 27, 1, ReadWrite, "End of Block interrupt enable">,
      Field<"M1", 28, 1, ReadWrite, "Word length">,
      Field<"FIFOEN", 29, 1, ReadWrite, "FIFO mode enable">,
      Field<"TXFEIE", 30, 1, ReadWrite, "TXFIFO empty interrupt enable">,
      Field<"RXFFIE", 31, 1, ReadWrite, "RXFIFO Full interrupt enable">,
    ], "Control register 1">,
    Register<"CR2", 0x4, 32, [
      Field<"SLVEN", 0, 1, ReadWrite, "Synchronous Slave mode enable">,
      Field<"DIS_NSS", 3, 1, ReadWrite, "When the DSI_NSS bit is set, the NSS pin input is ignored">,
      Field<"ADDM7", 4, 1, ReadWrite, "7-bit Address Detection/4-bit Address Detection">,
      Field<"LBDL", 5, 1, ReadWrite, "LIN break detection length">,
      Field<"LBDIE", 6, 1, ReadWrite, "LIN break detection interrupt enable">,
      Field<"LBCL", 8, 1, ReadWrite, "Last bit clock pulse">,
      Field<"CPHA", 9, 1, ReadWrite, "Clock phase">,
      Field<"CPOL", 10, 1, ReadWrite, "Clock polarity">,
      Field<"CLKEN", 11, 1, ReadWrite, "Clock enable">,
      Field<"STOP", 12, 2, ReadWrite, "STOP bits">,
      Field<"LINEN", 14, 1, ReadWrite, "LIN mode enable">,
      Field<"SWAP", 15, 1, ReadWrite, "Swap TX/RX pins">,
      Field<"RXINV", 16, 1, ReadWrite, "RX pin active level inversion">,
      Field<"TXINV", 17, 1, ReadWrite, "TX pin active level inversion">,
      Field<"TAINV", 18, 1, ReadWrite, "Binary data inversion">,
      Field<"MSBFIRST", 19, 1, ReadWrite, "Most significant bit first">,
      Field<"ABREN", 20, 1, ReadWrite, "Auto baud rate enable">,
      Field<"ABRMOD0", 21, 1, ReadWrite, "ABRMOD0">,
      Field<"ABRMOD1", 22, 1, ReadWrite, "Auto baud rate mode">,
      Field<"RTOEN", 23, 1, ReadWrite, "Receiver timeout enable">,
      Field<"ADD0_3", 24, 4, ReadWrite, "Address of the USART node">,
      Field<"ADD4_7", 28, 4, ReadWrite, "Address of the USART node">,
    ], "Control register 2">,
    Register<"CR3", 0x8, 32, [
      Field<"EIE", 0, 1, ReadWrite, "Error interrupt enable">,
      Field<"IREN", 1, 1, ReadWrite, "Ir mode enable">,
      Field<"IRLP", 2, 1, ReadWrite, "Ir low-power">,
      Field<"HDSEL", 3, 1, ReadWrite, "Half-duplex selection">,
      Field<"NACK", 4, 1, ReadWrite, "Smartcard NACK enable">,
      Field<"SCEN", 5, 1, ReadWrite, "Smartcard mode enable">,
      Field<"DMAR", 6, 1, ReadWrite, "DMA enable receiver">,
      Field<"DMAT", 7, 1, ReadWrite, "DMA enable transmitter">,
      Field<"RTSE", 8, 1, ReadWrite, "RTS enable">,
      Field<"CTSE", 9, 1, ReadWrite, "CTS enable">,
      Field<"CTSIE", 10, 1, ReadWrite, "CTS interrupt enable">,
      Field<"ONEBIT", 11, 1, ReadWrite, "One sample bit method enable">,
      Field<"OVRDIS", 12, 1, ReadWrite, "Overrun Disable">,
      Field<"DDRE", 13, 1, ReadWrite, "DMA Disable on Reception Error">,
      Field<"DEM", 14, 1, ReadWrite, "Driver enable mode">,
      Field<"DEP", 15, 1, ReadWrite, "Driver enable polarity selection">,
      Field<"SCARCNT", 17, 3, ReadWrite, "Smartcard auto-retry count">,
      Field<"WUS", 20, 2, ReadWrite, "Wakeup from Stop mode interrupt flag selection">,
      Field<"WUFIE", 22, 1, ReadWrite, "Wakeup from Stop mode interrupt enable">,
      Field<"TXFTIE", 23, 1, ReadWrite, "TXFIFO threshold interrupt enable">,
      Field<"TCBGTIE", 24, 1, ReadWrite, "Transmission Complete before guard time, interrupt enable">,
      Field<"RXFTCFG", 25, 3, ReadWrite, "Receive FIFO threshold configuration">,
      Field<"RXFTIE", 28, 1, ReadWrite, "RXFIFO threshold interrupt enable">,
      Field<"TXFTCFG", 29, 3, ReadWrite, "TXFIFO threshold configuration">,
    ], "Control register 3">,
    Register<"BRR", 0xc, 32, [
      Field<"BRR_0_3", 0, 4, ReadWrite, "DIV_Fraction">,
      Field<"BRR_4_15", 4, 12, ReadWrite, "DIV_Mantissa">,
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
      Field<"NF", 2, 1, ReadWrite, "NF">,
      Field<"ORE", 3, 1, ReadWrite, "ORE">,
      Field<"IDLE", 4, 1, ReadWrite, "IDLE">,
      Field<"RXNE", 5, 1, ReadWrite, "RXNE">,
      Field<"TC", 6, 1, ReadWrite, "TC">,
      Field<"TXE", 7, 1, ReadWrite, "TXE">,
      Field<"LBDF", 8, 1, ReadWrite, "LBDF">,
      Field<"CTSIF", 9, 1, ReadWrite, "CTSIF">,
      Field<"CTS", 10, 1, ReadWrite, "CTS">,
      Field<"RTOF", 11, 1, ReadWrite, "RTOF">,
      Field<"EOBF", 12, 1, ReadWrite, "EOBF">,
      Field<"UDR", 13, 1, ReadWrite, "SPI slave underrun error flag">,
      Field<"ABRE", 14, 1, ReadWrite, "ABRE">,
      Field<"ABRF", 15, 1, ReadWrite, "ABRF">,
      Field<"BUSY", 16, 1, ReadWrite, "BUSY">,
      Field<"CMF", 17, 1, ReadWrite, "CMF">,
      Field<"SBKF", 18, 1, ReadWrite, "SBKF">,
      Field<"RWU", 19, 1, ReadWrite, "RWU">,
      Field<"WUF", 20, 1, ReadWrite, "WUF">,
      Field<"TEACK", 21, 1, ReadWrite, "TEACK">,
      Field<"REACK", 22, 1, ReadWrite, "REACK">,
      Field<"TXFE", 23, 1, ReadWrite, "TXFIFO Empty">,
      Field<"RXFF", 24, 1, ReadWrite, "RXFIFO Full">,
      Field<"TCBGT", 25, 1, ReadWrite, "Transmission complete before guard time flag">,
      Field<"RXFT", 26, 1, ReadWrite, "RXFIFO threshold flag">,
      Field<"TXFT", 27, 1, ReadWrite, "TXFIFO threshold flag">,
    ], "Interrupt & status register">,
    Register<"ICR", 0x20, 32, [
      Field<"PECF", 0, 1, ReadWrite, "Parity error clear flag">,
      Field<"FECF", 1, 1, ReadWrite, "Framing error clear flag">,
      Field<"NCF", 2, 1, ReadWrite, "Noise detected clear flag">,
      Field<"ORECF", 3, 1, ReadWrite, "Overrun error clear flag">,
      Field<"IDLECF", 4, 1, ReadWrite, "Idle line detected clear flag">,
      Field<"TXFECF", 5, 1, ReadWrite, "TXFIFO empty clear flag">,
      Field<"TCCF", 6, 1, ReadWrite, "Transmission complete clear flag">,
      Field<"TCBGTC", 7, 1, ReadWrite, "Transmission complete before Guard time clear flag">,
      Field<"LBDCF", 8, 1, ReadWrite, "LIN break detection clear flag">,
      Field<"CTSCF", 9, 1, ReadWrite, "CTS clear flag">,
      Field<"RTOCF", 11, 1, ReadWrite, "Receiver timeout clear flag">,
      Field<"EOBCF", 12, 1, ReadWrite, "End of block clear flag">,
      Field<"UDRCF", 13, 1, ReadWrite, "SPI slave underrun clear flag">,
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
    ], "USART prescaler register">,
  ];
}

class USARTInstance<string Name, int Base> : PeripheralInstance<Name, Base, USARTPeripheral>;
def USART2 : USARTInstance<"USART2", 0x40004400>;
def USART3 : USARTInstance<"USART3", 0x40004800>;
def UART4 : USARTInstance<"UART4", 0x40004c00>;
def UART5 : USARTInstance<"UART5", 0x40005000>;
def UART7 : USARTInstance<"UART7", 0x40007800>;
def UART8 : USARTInstance<"UART8", 0x40007c00>;
def USART1 : USARTInstance<"USART1", 0x40011000>;
def USART6 : USARTInstance<"USART6", 0x40011400>;

def USARTGroup : PeripheralGroup<[ USART1, USART2, USART3, UART4, UART5, USART6, UART7, UART8 ]>;

#endif // _PERIPHERALS_USART_TD
