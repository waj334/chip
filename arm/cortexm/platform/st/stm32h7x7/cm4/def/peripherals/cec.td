#ifndef _PERIPHERALS_CEC_TD
#define _PERIPHERALS_CEC_TD

include "base.td"

def CECPeripheral : PeripheralType<"CEC", "CEC"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"CECEN", 0, 1, ReadWrite, "CEC Enable The CECEN bit is set and cleared by software. CECEN=1 starts message reception and enables the TXSOM control. CECEN=0 disables the CEC peripheral, clears all bits of CEC_CR register and aborts any on-going reception or transmission.">,
      Field<"TXSOM", 1, 1, ReadWrite, "Tx Start Of Message TXSOM is set by software to command transmission of the first byte of a CEC message. If the CEC message consists of only one byte, TXEOM must be set before of TXSOM. Start-Bit is effectively started on the CEC line after SFT is counted. If TXSOM is set while a message reception is ongoing, transmission will start after the end of reception. TXSOM is cleared by hardware after the last byte of the message is sent with a positive acknowledge (TXEND=1), in case of transmission underrun (TXUDR=1), negative acknowledge (TXACKE=1), and transmission error (TXERR=1). It is also cleared by CECEN=0. It is not cleared and transmission is automatically retried in case of arbitration lost (ARBLST=1). TXSOM can be also used as a status bit informing application whether any transmission request is pending or under execution. The application can abort a transmission request at any time by clearing the CECEN bit. Note: TXSOM must be set when CECEN=1 TXSOM must be set when transmission data is available into TXDR HEADERs first four bits containing own peripheral address are taken from TXDR[7:4], not from CEC_CFGR.OAR which is used only for reception">,
      Field<"TXEOM", 2, 1, ReadWrite, "Tx End Of Message The TXEOM bit is set by software to command transmission of the last byte of a CEC message. TXEOM is cleared by hardware at the same time and under the same conditions as for TXSOM. Note: TXEOM must be set when CECEN=1 TXEOM must be set before writing transmission data to TXDR If TXEOM is set when TXSOM=0, transmitted message will consist of 1 byte (HEADER) only (PING message)">,
    ], "CEC control register">,
    Register<"CFGR", 0x4, 32, [
      Field<"SFT", 0, 3, ReadWrite, "Signal Free Time SFT bits are set by software. In the SFT=0x0 configuration the number of nominal data bit periods waited before transmission is ruled by hardware according to the transmission history. In all the other configurations the SFT number is determined by software. * 0x0 ** 2.5 Data-Bit periods if CEC is the last bus initiator with unsuccessful transmission (ARBLST=1, TXERR=1, TXUDR=1 or TXACKE= 1) ** 4 Data-Bit periods if CEC is the new bus initiator ** 6 Data-Bit periods if CEC is the last bus initiator with successful transmission (TXEOM=1) * 0x1: 0.5 nominal data bit periods * 0x2: 1.5 nominal data bit periods * 0x3: 2.5 nominal data bit periods * 0x4: 3.5 nominal data bit periods * 0x5: 4.5 nominal data bit periods * 0x6: 5.5 nominal data bit periods * 0x7: 6.5 nominal data bit periods">,
      Field<"RXTOL", 3, 1, ReadWrite, "Rx-Tolerance The RXTOL bit is set and cleared by software. ** Start-Bit, +/- 200 s rise, +/- 200 s fall. ** Data-Bit: +/- 200 s rise. +/- 350 s fall. ** Start-Bit: +/- 400 s rise, +/- 400 s fall ** Data-Bit: +/-300 s rise, +/- 500 s fall">,
      Field<"BRESTP", 4, 1, ReadWrite, "Rx-Stop on Bit Rising Error The BRESTP bit is set and cleared by software.">,
      Field<"BREGEN", 5, 1, ReadWrite, "Generate Error-Bit on Bit Rising Error The BREGEN bit is set and cleared by software. Note: If BRDNOGEN=0, an Error-bit is generated upon BRE detection with BRESTP=1 in broadcast even if BREGEN=0">,
      Field<"LBPEGEN", 6, 1, ReadWrite, "Generate Error-Bit on Long Bit Period Error The LBPEGEN bit is set and cleared by software. Note: If BRDNOGEN=0, an Error-bit is generated upon LBPE detection in broadcast even if LBPEGEN=0">,
      Field<"BRDNOGEN", 7, 1, ReadWrite, "Avoid Error-Bit Generation in Broadcast The BRDNOGEN bit is set and cleared by software.">,
      Field<"SFTOPT", 8, 1, ReadWrite, "SFT Option Bit The SFTOPT bit is set and cleared by software.">,
      Field<"OAR", 16, 15, ReadWrite, "Own addresses configuration The OAR bits are set by software to select which destination logical addresses has to be considered in receive mode. Each bit, when set, enables the CEC logical address identified by the given bit position. At the end of HEADER reception, the received destination address is compared with the enabled addresses. In case of matching address, the incoming message is acknowledged and received. In case of non-matching address, the incoming message is received only in listen mode (LSTN=1), but without acknowledge sent. Broadcast messages are always received. Example: OAR = 0b000 0000 0010 0001 means that CEC acknowledges addresses 0x0 and 0x5. Consequently, each message directed to one of these addresses is received.">,
      Field<"LSTN", 31, 1, ReadWrite, "Listen mode LSTN bit is set and cleared by software.">,
    ], "This register is used to configure the HDMI-CEC controller. It is mandatory to write CEC_CFGR only when CECEN=0.">,
    Register<"TXDR", 0x8, 32, [
      Field<"TXD", 0, 8, ReadWrite, "Tx Data register. TXD is a write-only register containing the data byte to be transmitted. Note: TXD must be written when TXSTART=1">,
    ], "CEC Tx data register">,
    Register<"RXDR", 0xc, 32, [
      Field<"RXD", 0, 8, ReadWrite, "Rx Data register. RXD is read-only and contains the last data byte which has been received from the CEC line.">,
    ], "CEC Rx Data Register">,
    Register<"ISR", 0x10, 32, [
      Field<"RXBR", 0, 1, ReadWrite, "Rx-Byte Received The RXBR bit is set by hardware to inform application that a new byte has been received from the CEC line and stored into the RXD buffer. RXBR is cleared by software write at 1.">,
      Field<"RXEND", 1, 1, ReadWrite, "End Of Reception RXEND is set by hardware to inform application that the last byte of a CEC message is received from the CEC line and stored into the RXD buffer. RXEND is set at the same time of RXBR. RXEND is cleared by software write at 1.">,
      Field<"RXOVR", 2, 1, ReadWrite, "Rx-Overrun RXOVR is set by hardware if RXBR is not yet cleared at the time a new byte is received on the CEC line and stored into RXD. RXOVR assertion stops message reception so that no acknowledge is sent. In case of broadcast, a negative acknowledge is sent. RXOVR is cleared by software write at 1.">,
      Field<"BRE", 3, 1, ReadWrite, "Rx-Bit Rising Error BRE is set by hardware in case a Data-Bit waveform is detected with Bit Rising Error. BRE is set either at the time the misplaced rising edge occurs, or at the end of the maximum BRE tolerance allowed by RXTOL, in case rising edge is still longing. BRE stops message reception if BRESTP=1. BRE generates an Error-Bit on the CEC line if BREGEN=1. BRE is cleared by software write at 1.">,
      Field<"SBPE", 4, 1, ReadWrite, "Rx-Short Bit Period Error SBPE is set by hardware in case a Data-Bit waveform is detected with Short Bit Period Error. SBPE is set at the time the anticipated falling edge occurs. SBPE generates an Error-Bit on the CEC line. SBPE is cleared by software write at 1.">,
      Field<"LBPE", 5, 1, ReadWrite, "Rx-Long Bit Period Error LBPE is set by hardware in case a Data-Bit waveform is detected with Long Bit Period Error. LBPE is set at the end of the maximum bit-extension tolerance allowed by RXTOL, in case falling edge is still longing. LBPE always stops reception of the CEC message. LBPE generates an Error-Bit on the CEC line if LBPEGEN=1. In case of broadcast, Error-Bit is generated even in case of LBPEGEN=0. LBPE is cleared by software write at 1.">,
      Field<"RXACKE", 6, 1, ReadWrite, "Rx-Missing Acknowledge In receive mode, RXACKE is set by hardware to inform application that no acknowledge was seen on the CEC line. RXACKE applies only for broadcast messages and in listen mode also for not directly addressed messages (destination address not enabled in OAR). RXACKE aborts message reception. RXACKE is cleared by software write at 1.">,
      Field<"ARBLST", 7, 1, ReadWrite, "Arbitration Lost ARBLST is set by hardware to inform application that CEC device is switching to reception due to arbitration lost event following the TXSOM command. ARBLST can be due either to a contending CEC device starting earlier or starting at the same time but with higher HEADER priority. After ARBLST assertion TXSOM bit keeps pending for next transmission attempt. ARBLST is cleared by software write at 1.">,
      Field<"TXBR", 8, 1, ReadWrite, "Tx-Byte Request TXBR is set by hardware to inform application that the next transmission data has to be written to TXDR. TXBR is set when the 4th bit of currently transmitted byte is sent. Application must write the next byte to TXDR within 6 nominal data-bit periods before transmission underrun error occurs (TXUDR). TXBR is cleared by software write at 1.">,
      Field<"TXEND", 9, 1, ReadWrite, "End of Transmission TXEND is set by hardware to inform application that the last byte of the CEC message has been successfully transmitted. TXEND clears the TXSOM and TXEOM control bits. TXEND is cleared by software write at 1.">,
      Field<"TXUDR", 10, 1, ReadWrite, "Tx-Buffer Underrun In transmission mode, TXUDR is set by hardware if application was not in time to load TXDR before of next byte transmission. TXUDR aborts message transmission and clears TXSOM and TXEOM control bits. TXUDR is cleared by software write at 1">,
      Field<"TXERR", 11, 1, ReadWrite, "Tx-Error In transmission mode, TXERR is set by hardware if the CEC initiator detects low impedance on the CEC line while it is released. TXERR aborts message transmission and clears TXSOM and TXEOM controls. TXERR is cleared by software write at 1.">,
      Field<"TXACKE", 12, 1, ReadWrite, "Tx-Missing Acknowledge Error In transmission mode, TXACKE is set by hardware to inform application that no acknowledge was received. In case of broadcast transmission, TXACKE informs application that a negative acknowledge was received. TXACKE aborts message transmission and clears TXSOM and TXEOM controls. TXACKE is cleared by software write at 1.">,
    ], "CEC Interrupt and Status Register">,
    Register<"IER", 0x14, 32, [
      Field<"RXBRIE", 0, 1, ReadWrite, "Rx-Byte Received Interrupt Enable The RXBRIE bit is set and cleared by software.">,
      Field<"RXENDIE", 1, 1, ReadWrite, "End Of Reception Interrupt Enable The RXENDIE bit is set and cleared by software.">,
      Field<"RXOVRIE", 2, 1, ReadWrite, "Rx-Buffer Overrun Interrupt Enable The RXOVRIE bit is set and cleared by software.">,
      Field<"BREIE", 3, 1, ReadWrite, "Bit Rising Error Interrupt Enable The BREIE bit is set and cleared by software.">,
      Field<"SBPEIE", 4, 1, ReadWrite, "Short Bit Period Error Interrupt Enable The SBPEIE bit is set and cleared by software.">,
      Field<"LBPEIE", 5, 1, ReadWrite, "Long Bit Period Error Interrupt Enable The LBPEIE bit is set and cleared by software.">,
      Field<"RXACKIE", 6, 1, ReadWrite, "Rx-Missing Acknowledge Error Interrupt Enable The RXACKIE bit is set and cleared by software.">,
      Field<"ARBLSTIE", 7, 1, ReadWrite, "Arbitration Lost Interrupt Enable The ARBLSTIE bit is set and cleared by software.">,
      Field<"TXBRIE", 8, 1, ReadWrite, "Tx-Byte Request Interrupt Enable The TXBRIE bit is set and cleared by software.">,
      Field<"TXENDIE", 9, 1, ReadWrite, "Tx-End Of Message Interrupt Enable The TXENDIE bit is set and cleared by software.">,
      Field<"TXUDRIE", 10, 1, ReadWrite, "Tx-Underrun Interrupt Enable The TXUDRIE bit is set and cleared by software.">,
      Field<"TXERRIE", 11, 1, ReadWrite, "Tx-Error Interrupt Enable The TXERRIE bit is set and cleared by software.">,
      Field<"TXACKIE", 12, 1, ReadWrite, "Tx-Missing Acknowledge Error Interrupt Enable The TXACKEIE bit is set and cleared by software.">,
    ], "CEC interrupt enable register">,
  ];
}

class CECInstance<string Name, int Base> : PeripheralInstance<Name, Base, CECPeripheral>;
def CEC : CECInstance<"CEC", 0x40006c00>;

#endif // _PERIPHERALS_CEC_TD
