#ifndef _PERIPHERALS_ETHERNET_DMA_TD
#define _PERIPHERALS_ETHERNET_DMA_TD

include "base.td"

def Ethernet_DMAPeripheral : PeripheralType<"Ethernet_DMA", "Ethernet: DMA mode register (DMA)"> {
  let accessWidth = 32;
  let registers = [
    Register<"DMAMR", 0x0, 32, [
      Field<"SWR", 0, 1, ReadWrite, "Software Reset">,
      Field<"DA", 1, 1, Read, "DMA Tx or Rx Arbitration Scheme">,
      Field<"TXPR", 11, 1, Read, "Transmit priority">,
      Field<"PR", 12, 3, Read, "Priority ratio">,
      Field<"INTM", 16, 1, ReadWrite, "Interrupt Mode">,
    ], "DMA mode register">,
    Register<"DMASBMR", 0x4, 32, [
      Field<"FB", 0, 1, ReadWrite, "Fixed Burst Length">,
      Field<"AAL", 12, 1, ReadWrite, "Address-Aligned Beats">,
      Field<"MB", 14, 1, Read, "Mixed Burst">,
      Field<"RB", 15, 1, Read, "Rebuild INCRx Burst">,
    ], "System bus mode register">,
    Register<"DMAISR", 0x8, 32, [
      Field<"DC0IS", 0, 1, ReadWrite, "DMA Channel Interrupt Status">,
      Field<"MTLIS", 16, 1, ReadWrite, "MTL Interrupt Status">,
      Field<"MACIS", 17, 1, ReadWrite, "MAC Interrupt Status">,
    ], "Interrupt status register">,
    Register<"DMADSR", 0xc, 32, [
      Field<"AXWHSTS", 0, 1, ReadWrite, "AHB Master Write Channel">,
      Field<"RPS0", 8, 4, ReadWrite, "DMA Channel Receive Process State">,
      Field<"TPS0", 12, 4, ReadWrite, "DMA Channel Transmit Process State">,
    ], "Debug status register">,
    Register<"DMACCR", 0x100, 32, [
      Field<"MSS", 0, 14, ReadWrite, "Maximum Segment Size">,
      Field<"PBLX8", 16, 1, ReadWrite, "8xPBL mode">,
      Field<"DSL", 18, 3, ReadWrite, "Descriptor Skip Length">,
    ], "Channel control register">,
    Register<"DMACTxCR", 0x104, 32, [
      Field<"ST", 0, 1, ReadWrite, "Start or Stop Transmission Command">,
      Field<"OSF", 4, 1, ReadWrite, "Operate on Second Packet">,
      Field<"TSE", 12, 1, ReadWrite, "TCP Segmentation Enabled">,
      Field<"TXPBL", 16, 6, ReadWrite, "Transmit Programmable Burst Length">,
    ], "Channel transmit control register">,
    Register<"DMACRxCR", 0x108, 32, [
      Field<"SR", 0, 1, ReadWrite, "Start or Stop Receive Command">,
      Field<"RBSZ", 1, 14, ReadWrite, "Receive Buffer size">,
      Field<"RXPBL", 16, 6, ReadWrite, "RXPBL">,
      Field<"RPF", 31, 1, ReadWrite, "DMA Rx Channel Packet Flush">,
    ], "Channel receive control register">,
    Register<"DMACTxDLAR", 0x114, 32, [
      Field<"TDESLA", 2, 30, ReadWrite, "Start of Transmit List">,
    ], "Channel Tx descriptor list address register">,
    Register<"DMACRxDLAR", 0x11c, 32, [
      Field<"RDESLA", 2, 30, ReadWrite, "Start of Receive List">,
    ], "Channel Rx descriptor list address register">,
    Register<"DMACTxDTPR", 0x120, 32, [
      Field<"TDT", 2, 30, ReadWrite, "Transmit Descriptor Tail Pointer">,
    ], "Channel Tx descriptor tail pointer register">,
    Register<"DMACRxDTPR", 0x128, 32, [
      Field<"RDT", 2, 30, ReadWrite, "Receive Descriptor Tail Pointer">,
    ], "Channel Rx descriptor tail pointer register">,
    Register<"DMACTxRLR", 0x12c, 32, [
      Field<"TDRL", 0, 10, ReadWrite, "Transmit Descriptor Ring Length">,
    ], "Channel Tx descriptor ring length register">,
    Register<"DMACRxRLR", 0x130, 32, [
      Field<"RDRL", 0, 10, ReadWrite, "Receive Descriptor Ring Length">,
    ], "Channel Rx descriptor ring length register">,
    Register<"DMACIER", 0x134, 32, [
      Field<"TIE", 0, 1, ReadWrite, "Transmit Interrupt Enable">,
      Field<"TXSE", 1, 1, ReadWrite, "Transmit Stopped Enable">,
      Field<"TBUE", 2, 1, ReadWrite, "Transmit Buffer Unavailable Enable">,
      Field<"RIE", 6, 1, ReadWrite, "Receive Interrupt Enable">,
      Field<"RBUE", 7, 1, ReadWrite, "Receive Buffer Unavailable Enable">,
      Field<"RSE", 8, 1, ReadWrite, "Receive Stopped Enable">,
      Field<"RWTE", 9, 1, ReadWrite, "Receive Watchdog Timeout Enable">,
      Field<"ETIE", 10, 1, ReadWrite, "Early Transmit Interrupt Enable">,
      Field<"ERIE", 11, 1, ReadWrite, "Early Receive Interrupt Enable">,
      Field<"FBEE", 12, 1, ReadWrite, "Fatal Bus Error Enable">,
      Field<"CDEE", 13, 1, ReadWrite, "Context Descriptor Error Enable">,
      Field<"AIE", 14, 1, ReadWrite, "Abnormal Interrupt Summary Enable">,
      Field<"NIE", 15, 1, ReadWrite, "Normal Interrupt Summary Enable">,
    ], "Channel interrupt enable register">,
    Register<"DMACRxIWTR", 0x138, 32, [
      Field<"RWT", 0, 8, ReadWrite, "Receive Interrupt Watchdog Timer Count">,
    ], "Channel Rx interrupt watchdog timer register">,
    Register<"DMACCATxDR", 0x144, 32, [
      Field<"CURTDESAPTR", 0, 32, ReadWrite, "Application Transmit Descriptor Address Pointer">,
    ], "Channel current application transmit descriptor register">,
    Register<"DMACCARxDR", 0x14c, 32, [
      Field<"CURRDESAPTR", 0, 32, ReadWrite, "Application Receive Descriptor Address Pointer">,
    ], "Channel current application receive descriptor register">,
    Register<"DMACCATxBR", 0x154, 32, [
      Field<"CURTBUFAPTR", 0, 32, ReadWrite, "Application Transmit Buffer Address Pointer">,
    ], "Channel current application transmit buffer register">,
    Register<"DMACCARxBR", 0x15c, 32, [
      Field<"CURRBUFAPTR", 0, 32, ReadWrite, "Application Receive Buffer Address Pointer">,
    ], "Channel current application receive buffer register">,
    Register<"DMACSR", 0x160, 32, [
      Field<"TI", 0, 1, ReadWrite, "Transmit Interrupt">,
      Field<"TPS", 1, 1, ReadWrite, "Transmit Process Stopped">,
      Field<"TBU", 2, 1, ReadWrite, "Transmit Buffer Unavailable">,
      Field<"RI", 6, 1, ReadWrite, "Receive Interrupt">,
      Field<"RBU", 7, 1, ReadWrite, "Receive Buffer Unavailable">,
      Field<"RPS", 8, 1, ReadWrite, "Receive Process Stopped">,
      Field<"RWT", 9, 1, ReadWrite, "Receive Watchdog Timeout">,
      Field<"ET", 10, 1, ReadWrite, "Early Transmit Interrupt">,
      Field<"ER", 11, 1, ReadWrite, "Early Receive Interrupt">,
      Field<"FBE", 12, 1, ReadWrite, "Fatal Bus Error">,
      Field<"CDE", 13, 1, ReadWrite, "Context Descriptor Error">,
      Field<"AIS", 14, 1, ReadWrite, "Abnormal Interrupt Summary">,
      Field<"NIS", 15, 1, ReadWrite, "Normal Interrupt Summary">,
      Field<"TEB", 16, 3, Read, "Tx DMA Error Bits">,
      Field<"REB", 19, 3, Read, "Rx DMA Error Bits">,
    ], "Channel status register">,
    Register<"DMACMFCR", 0x16c, 32, [
      Field<"MFC", 0, 11, ReadWrite, "Dropped Packet Counters">,
      Field<"MFCO", 15, 1, ReadWrite, "Overflow status of the MFC Counter">,
    ], "Channel missed frame count register">,
  ];
}

class Ethernet_DMAInstance<string Name, int Base> : PeripheralInstance<Name, Base, Ethernet_DMAPeripheral>;
def Ethernet_DMA : Ethernet_DMAInstance<"Ethernet_DMA", 0x40029000>;

#endif // _PERIPHERALS_ETHERNET_DMA_TD
