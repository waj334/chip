#ifndef _PERIPHERALS_SWPMI_TD
#define _PERIPHERALS_SWPMI_TD

include "base.td"

def SWPMIPeripheral : PeripheralType<"SWPMI", "Single Wire Protocol Master Interface"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"RXDMA", 0, 1, ReadWrite, "Reception DMA enable">,
      Field<"TXDMA", 1, 1, ReadWrite, "Transmission DMA enable">,
      Field<"RXMODE", 2, 1, ReadWrite, "Reception buffering mode">,
      Field<"TXMODE", 3, 1, ReadWrite, "Transmission buffering mode">,
      Field<"LPBK", 4, 1, ReadWrite, "Loopback mode enable">,
      Field<"SWPACT", 5, 1, ReadWrite, "Single wire protocol master interface activate">,
      Field<"DEACT", 10, 1, ReadWrite, "Single wire protocol master interface deactivate">,
      Field<"SWPTEN", 11, 1, ReadWrite, "Single wire protocol master transceiver enable">,
    ], "SWPMI Configuration/Control register">,
    Register<"BRR", 0x4, 32, [
      Field<"BR", 0, 8, ReadWrite, "Bitrate prescaler">,
    ], "SWPMI Bitrate register">,
    Register<"ISR", 0xc, 32, [
      Field<"RXBFF", 0, 1, ReadWrite, "Receive buffer full flag">,
      Field<"TXBEF", 1, 1, ReadWrite, "Transmit buffer empty flag">,
      Field<"RXBERF", 2, 1, ReadWrite, "Receive CRC error flag">,
      Field<"RXOVRF", 3, 1, ReadWrite, "Receive overrun error flag">,
      Field<"TXUNRF", 4, 1, ReadWrite, "Transmit underrun error flag">,
      Field<"RXNE", 5, 1, ReadWrite, "Receive data register not empty">,
      Field<"TXE", 6, 1, ReadWrite, "Transmit data register empty">,
      Field<"TCF", 7, 1, ReadWrite, "Transfer complete flag">,
      Field<"SRF", 8, 1, ReadWrite, "Slave resume flag">,
      Field<"SUSP", 9, 1, ReadWrite, "SUSPEND flag">,
      Field<"DEACTF", 10, 1, ReadWrite, "DEACTIVATED flag">,
      Field<"RDYF", 11, 1, ReadWrite, "transceiver ready flag">,
    ], "SWPMI Interrupt and Status register">,
    Register<"ICR", 0x10, 32, [
      Field<"CRXBFF", 0, 1, ReadWrite, "Clear receive buffer full flag">,
      Field<"CTXBEF", 1, 1, ReadWrite, "Clear transmit buffer empty flag">,
      Field<"CRXBERF", 2, 1, ReadWrite, "Clear receive CRC error flag">,
      Field<"CRXOVRF", 3, 1, ReadWrite, "Clear receive overrun error flag">,
      Field<"CTXUNRF", 4, 1, ReadWrite, "Clear transmit underrun error flag">,
      Field<"CTCF", 7, 1, ReadWrite, "Clear transfer complete flag">,
      Field<"CSRF", 8, 1, ReadWrite, "Clear slave resume flag">,
      Field<"CRDYF", 11, 1, ReadWrite, "Clear transceiver ready flag">,
    ], "SWPMI Interrupt Flag Clear register">,
    Register<"IER", 0x14, 32, [
      Field<"RXBFIE", 0, 1, ReadWrite, "Receive buffer full interrupt enable">,
      Field<"TXBEIE", 1, 1, ReadWrite, "Transmit buffer empty interrupt enable">,
      Field<"RXBERIE", 2, 1, ReadWrite, "Receive CRC error interrupt enable">,
      Field<"RXOVRIE", 3, 1, ReadWrite, "Receive overrun error interrupt enable">,
      Field<"TXUNRIE", 4, 1, ReadWrite, "Transmit underrun error interrupt enable">,
      Field<"RIE", 5, 1, ReadWrite, "Receive interrupt enable">,
      Field<"TIE", 6, 1, ReadWrite, "Transmit interrupt enable">,
      Field<"TCIE", 7, 1, ReadWrite, "Transmit complete interrupt enable">,
      Field<"SRIE", 8, 1, ReadWrite, "Slave resume interrupt enable">,
      Field<"RDYIE", 11, 1, ReadWrite, "Transceiver ready interrupt enable">,
    ], "SWPMI Interrupt Enable register">,
    Register<"RFL", 0x18, 32, [
      Field<"RFL", 0, 5, ReadWrite, "Receive frame length">,
    ], "SWPMI Receive Frame Length register">,
    Register<"TDR", 0x1c, 32, [
      Field<"TD", 0, 32, ReadWrite, "Transmit data">,
    ], "SWPMI Transmit data register">,
    Register<"RDR", 0x20, 32, [
      Field<"RD", 0, 32, ReadWrite, "received data">,
    ], "SWPMI Receive data register">,
    Register<"OR", 0x24, 32, [
      Field<"SWP_TBYP", 0, 1, ReadWrite, "SWP transceiver bypass">,
      Field<"SWP_CLASS", 1, 1, ReadWrite, "SWP class selection">,
    ], "SWPMI Option register">,
  ];
}

class SWPMIInstance<string Name, int Base> : PeripheralInstance<Name, Base, SWPMIPeripheral>;
def SWPMI : SWPMIInstance<"SWPMI", 0x40008800>;

#endif // _PERIPHERALS_SWPMI_TD
