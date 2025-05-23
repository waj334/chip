#ifndef _PERIPHERALS_ETHERNET_MTL_TD
#define _PERIPHERALS_ETHERNET_MTL_TD

include "base.td"

def Ethernet_MTLPeripheral : PeripheralType<"Ethernet_MTL", "Ethernet: MTL mode register (MTL)"> {
  let accessWidth = 32;
  let registers = [
    Register<"MTLOMR", 0x0, 32, [
      Field<"DTXSTS", 1, 1, ReadWrite, "DTXSTS">,
      Field<"CNTPRST", 8, 1, ReadWrite, "CNTPRST">,
      Field<"CNTCLR", 9, 1, ReadWrite, "CNTCLR">,
    ], "Operating mode Register">,
    Register<"MTLISR", 0x20, 32, [
      Field<"Q0IS", 0, 1, ReadWrite, "Queue interrupt status">,
    ], "Interrupt status Register">,
    Register<"MTLTxQOMR", 0x100, 32, [
      Field<"FTQ", 0, 1, ReadWrite, "Flush Transmit Queue">,
      Field<"TSF", 1, 1, ReadWrite, "Transmit Store and Forward">,
      Field<"TXQEN", 2, 2, Read, "Transmit Queue Enable">,
      Field<"TTC", 4, 3, ReadWrite, "Transmit Threshold Control">,
      Field<"TQS", 16, 3, ReadWrite, "Transmit Queue Size">,
    ], "Tx queue operating mode Register">,
    Register<"MTLTxQUR", 0x104, 32, [
      Field<"UFFRMCNT", 0, 11, ReadWrite, "Underflow Packet Counter">,
      Field<"UFCNTOVF", 11, 1, ReadWrite, "UFCNTOVF">,
    ], "Tx queue underflow register">,
    Register<"MTLTxQDR", 0x108, 32, [
      Field<"TXQPAUSED", 0, 1, ReadWrite, "TXQPAUSED">,
      Field<"TRCSTS", 1, 2, ReadWrite, "TRCSTS">,
      Field<"TWCSTS", 3, 1, ReadWrite, "TWCSTS">,
      Field<"TXQSTS", 4, 1, ReadWrite, "TXQSTS">,
      Field<"TXSTSFSTS", 5, 1, ReadWrite, "TXSTSFSTS">,
      Field<"PTXQ", 16, 3, ReadWrite, "PTXQ">,
      Field<"STXSTSF", 20, 3, ReadWrite, "STXSTSF">,
    ], "Tx queue debug Register">,
    Register<"MTLQICSR", 0x12c, 32, [
      Field<"TXUNFIS", 0, 1, ReadWrite, "TXUNFIS">,
      Field<"TXUIE", 8, 1, ReadWrite, "TXUIE">,
      Field<"RXOVFIS", 16, 1, ReadWrite, "RXOVFIS">,
      Field<"RXOIE", 24, 1, ReadWrite, "RXOIE">,
    ], "Queue interrupt control status Register">,
    Register<"MTLRxQOMR", 0x130, 32, [
      Field<"RTC", 0, 2, ReadWrite, "RTC">,
      Field<"FUP", 3, 1, ReadWrite, "FUP">,
      Field<"FEP", 4, 1, ReadWrite, "FEP">,
      Field<"RSF", 5, 1, ReadWrite, "RSF">,
      Field<"DIS_TCP_EF", 6, 1, ReadWrite, "DIS_TCP_EF">,
      Field<"EHFC", 7, 1, ReadWrite, "EHFC">,
      Field<"RFA", 8, 3, ReadWrite, "RFA">,
      Field<"RFD", 14, 3, ReadWrite, "RFD">,
      Field<"RQS", 20, 3, Read, "RQS">,
    ], "Rx queue operating mode register">,
    Register<"MTLRxQMPOCR", 0x134, 32, [
      Field<"OVFPKTCNT", 0, 11, ReadWrite, "OVFPKTCNT">,
      Field<"OVFCNTOVF", 11, 1, ReadWrite, "OVFCNTOVF">,
      Field<"MISPKTCNT", 16, 11, ReadWrite, "MISPKTCNT">,
      Field<"MISCNTOVF", 27, 1, ReadWrite, "MISCNTOVF">,
    ], "Rx queue missed packet and overflow counter register">,
    Register<"MTLRxQDR", 0x138, 32, [
      Field<"RWCSTS", 0, 1, ReadWrite, "RWCSTS">,
      Field<"RRCSTS", 1, 2, ReadWrite, "RRCSTS">,
      Field<"RXQSTS", 4, 2, ReadWrite, "RXQSTS">,
      Field<"PRXQ", 16, 14, ReadWrite, "PRXQ">,
    ], "Rx queue debug register">,
  ];
}

class Ethernet_MTLInstance<string Name, int Base> : PeripheralInstance<Name, Base, Ethernet_MTLPeripheral>;
def Ethernet_MTL : Ethernet_MTLInstance<"Ethernet_MTL", 0x40028c00>;

#endif // _PERIPHERALS_ETHERNET_MTL_TD
