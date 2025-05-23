#ifndef _PERIPHERALS_SPDIFRX_TD
#define _PERIPHERALS_SPDIFRX_TD

include "base.td"

def SPDIFRXPeripheral : PeripheralType<"SPDIFRX", "Receiver Interface"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"SPDIFRXEN", 0, 2, ReadWrite, "Peripheral Block Enable">,
      Field<"RXDMAEN", 2, 1, ReadWrite, "Receiver DMA ENable for data flow">,
      Field<"RXSTEO", 3, 1, ReadWrite, "STerEO Mode">,
      Field<"DRFMT", 4, 2, ReadWrite, "RX Data format">,
      Field<"PMSK", 6, 1, ReadWrite, "Mask Parity error bit">,
      Field<"VMSK", 7, 1, ReadWrite, "Mask of Validity bit">,
      Field<"CUMSK", 8, 1, ReadWrite, "Mask of channel status and user bits">,
      Field<"PTMSK", 9, 1, ReadWrite, "Mask of Preamble Type bits">,
      Field<"CBDMAEN", 10, 1, ReadWrite, "Control Buffer DMA ENable for control flow">,
      Field<"CHSEL", 11, 1, ReadWrite, "Channel Selection">,
      Field<"NBTR", 12, 2, ReadWrite, "Maximum allowed re-tries during synchronization phase">,
      Field<"WFA", 14, 1, ReadWrite, "Wait For Activity">,
      Field<"INSEL", 16, 3, ReadWrite, "input selection">,
      Field<"CKSEN", 20, 1, ReadWrite, "Symbol Clock Enable">,
      Field<"CKSBKPEN", 21, 1, ReadWrite, "Backup Symbol Clock Enable">,
    ], "Control register">,
    Register<"IMR", 0x4, 32, [
      Field<"RXNEIE", 0, 1, ReadWrite, "RXNE interrupt enable">,
      Field<"CSRNEIE", 1, 1, ReadWrite, "Control Buffer Ready Interrupt Enable">,
      Field<"PERRIE", 2, 1, ReadWrite, "Parity error interrupt enable">,
      Field<"OVRIE", 3, 1, ReadWrite, "Overrun error Interrupt Enable">,
      Field<"SBLKIE", 4, 1, ReadWrite, "Synchronization Block Detected Interrupt Enable">,
      Field<"SYNCDIE", 5, 1, ReadWrite, "Synchronization Done">,
      Field<"IFEIE", 6, 1, ReadWrite, "Serial Interface Error Interrupt Enable">,
    ], "Interrupt mask register">,
    Register<"SR", 0x8, 32, [
      Field<"RXNE", 0, 1, ReadWrite, "Read data register not empty">,
      Field<"CSRNE", 1, 1, ReadWrite, "Control Buffer register is not empty">,
      Field<"PERR", 2, 1, ReadWrite, "Parity error">,
      Field<"OVR", 3, 1, ReadWrite, "Overrun error">,
      Field<"SBD", 4, 1, ReadWrite, "Synchronization Block Detected">,
      Field<"SYNCD", 5, 1, ReadWrite, "Synchronization Done">,
      Field<"FERR", 6, 1, ReadWrite, "Framing error">,
      Field<"SERR", 7, 1, ReadWrite, "Synchronization error">,
      Field<"TERR", 8, 1, ReadWrite, "Time-out error">,
      Field<"WIDTH5", 16, 15, ReadWrite, "Duration of 5 symbols counted with SPDIF_CLK">,
    ], "Status register">,
    Register<"IFCR", 0xc, 32, [
      Field<"PERRCF", 2, 1, ReadWrite, "Clears the Parity error flag">,
      Field<"OVRCF", 3, 1, ReadWrite, "Clears the Overrun error flag">,
      Field<"SBDCF", 4, 1, ReadWrite, "Clears the Synchronization Block Detected flag">,
      Field<"SYNCDCF", 5, 1, ReadWrite, "Clears the Synchronization Done flag">,
    ], "Interrupt Flag Clear register">,
    Register<"DR_00", 0x10, 32, [
      Field<"DR", 0, 24, ReadWrite, "Parity Error bit">,
      Field<"PE", 24, 1, ReadWrite, "Parity Error bit">,
      Field<"V", 25, 1, ReadWrite, "Validity bit">,
      Field<"U", 26, 1, ReadWrite, "User bit">,
      Field<"C", 27, 1, ReadWrite, "Channel Status bit">,
      Field<"PT", 28, 2, ReadWrite, "Preamble Type">,
    ], "Data input register">,
    Register<"CSR", 0x14, 32, [
      Field<"USR", 0, 16, ReadWrite, "User data information">,
      Field<"CS", 16, 8, ReadWrite, "Channel A status information">,
      Field<"SOB", 24, 1, ReadWrite, "Start Of Block">,
    ], "Channel Status register">,
    Register<"DIR", 0x18, 32, [
      Field<"THI", 0, 13, ReadWrite, "Threshold HIGH">,
      Field<"TLO", 16, 13, ReadWrite, "Threshold LOW">,
    ], "Debug Information register">,
    Register<"VERR", 0x3f4, 32, [
      Field<"MINREV", 0, 4, ReadWrite, "Minor revision">,
      Field<"MAJREV", 4, 4, ReadWrite, "Major revision">,
    ], "SPDIFRX version register">,
    Register<"IDR", 0x3f8, 32, [
      Field<"ID", 0, 32, ReadWrite, "SPDIFRX identifier">,
    ], "SPDIFRX identification register">,
    Register<"SIDR", 0x3fc, 32, [
      Field<"SID", 0, 32, ReadWrite, "Size identification">,
    ], "SPDIFRX size identification register">,
    Register<"DR_01", 0x10, 32, [
      Field<"PE", 0, 1, ReadWrite, "Parity Error bit">,
      Field<"V", 1, 1, ReadWrite, "Validity bit">,
      Field<"U", 2, 1, ReadWrite, "User bit">,
      Field<"C", 3, 1, ReadWrite, "Channel Status bit">,
      Field<"PT", 4, 2, ReadWrite, "Preamble Type">,
      Field<"DR", 8, 24, ReadWrite, "Data value">,
    ], "Data input register">,
    Register<"DR_10", 0x10, 32, [
      Field<"DRNL1", 0, 16, ReadWrite, "Data value">,
      Field<"DRNL2", 16, 16, ReadWrite, "Data value">,
    ], "Data input register">,
  ];
}

class SPDIFRXInstance<string Name, int Base> : PeripheralInstance<Name, Base, SPDIFRXPeripheral>;
def SPDIFRX : SPDIFRXInstance<"SPDIFRX", 0x40004000>;

#endif // _PERIPHERALS_SPDIFRX_TD
