#ifndef _PERIPHERALS_DCMI_TD
#define _PERIPHERALS_DCMI_TD

include "base.td"

def DCMIPeripheral : PeripheralType<"DCMI", "Digital camera interface"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"CAPTURE", 0, 1, ReadWrite, "Capture enable">,
      Field<"CM", 1, 1, ReadWrite, "Capture mode">,
      Field<"CROP", 2, 1, ReadWrite, "Crop feature">,
      Field<"JPEG", 3, 1, ReadWrite, "JPEG format">,
      Field<"ESS", 4, 1, ReadWrite, "Embedded synchronization select">,
      Field<"PCKPOL", 5, 1, ReadWrite, "Pixel clock polarity">,
      Field<"HSPOL", 6, 1, ReadWrite, "Horizontal synchronization polarity">,
      Field<"VSPOL", 7, 1, ReadWrite, "Vertical synchronization polarity">,
      Field<"FCRC", 8, 2, ReadWrite, "Frame capture rate control">,
      Field<"EDM", 10, 2, ReadWrite, "Extended data mode">,
      Field<"ENABLE", 14, 1, ReadWrite, "DCMI enable">,
      Field<"BSM", 16, 2, ReadWrite, "Byte Select mode">,
      Field<"OEBS", 18, 1, ReadWrite, "Odd/Even Byte Select (Byte Select Start)">,
      Field<"LSM", 19, 1, ReadWrite, "Line Select mode">,
      Field<"OELS", 20, 1, ReadWrite, "Odd/Even Line Select (Line Select Start)">,
    ], "control register 1">,
    Register<"SR", 0x4, 32, [
      Field<"HSYNC", 0, 1, ReadWrite, "HSYNC">,
      Field<"VSYNC", 1, 1, ReadWrite, "VSYNC">,
      Field<"FNE", 2, 1, ReadWrite, "FIFO not empty">,
    ], "status register">,
    Register<"RIS", 0x8, 32, [
      Field<"FRAME_RIS", 0, 1, ReadWrite, "Capture complete raw interrupt status">,
      Field<"OVR_RIS", 1, 1, ReadWrite, "Overrun raw interrupt status">,
      Field<"ERR_RIS", 2, 1, ReadWrite, "Synchronization error raw interrupt status">,
      Field<"VSYNC_RIS", 3, 1, ReadWrite, "VSYNC raw interrupt status">,
      Field<"LINE_RIS", 4, 1, ReadWrite, "Line raw interrupt status">,
    ], "raw interrupt status register">,
    Register<"IER", 0xc, 32, [
      Field<"FRAME_IE", 0, 1, ReadWrite, "Capture complete interrupt enable">,
      Field<"OVR_IE", 1, 1, ReadWrite, "Overrun interrupt enable">,
      Field<"ERR_IE", 2, 1, ReadWrite, "Synchronization error interrupt enable">,
      Field<"VSYNC_IE", 3, 1, ReadWrite, "VSYNC interrupt enable">,
      Field<"LINE_IE", 4, 1, ReadWrite, "Line interrupt enable">,
    ], "interrupt enable register">,
    Register<"MIS", 0x10, 32, [
      Field<"FRAME_MIS", 0, 1, ReadWrite, "Capture complete masked interrupt status">,
      Field<"OVR_MIS", 1, 1, ReadWrite, "Overrun masked interrupt status">,
      Field<"ERR_MIS", 2, 1, ReadWrite, "Synchronization error masked interrupt status">,
      Field<"VSYNC_MIS", 3, 1, ReadWrite, "VSYNC masked interrupt status">,
      Field<"LINE_MIS", 4, 1, ReadWrite, "Line masked interrupt status">,
    ], "masked interrupt status register">,
    Register<"ICR", 0x14, 32, [
      Field<"FRAME_ISC", 0, 1, ReadWrite, "Capture complete interrupt status clear">,
      Field<"OVR_ISC", 1, 1, ReadWrite, "Overrun interrupt status clear">,
      Field<"ERR_ISC", 2, 1, ReadWrite, "Synchronization error interrupt status clear">,
      Field<"VSYNC_ISC", 3, 1, ReadWrite, "Vertical synch interrupt status clear">,
      Field<"LINE_ISC", 4, 1, ReadWrite, "line interrupt status clear">,
    ], "interrupt clear register">,
    Register<"ESCR", 0x18, 32, [
      Field<"FSC", 0, 8, ReadWrite, "Frame start delimiter code">,
      Field<"LSC", 8, 8, ReadWrite, "Line start delimiter code">,
      Field<"LEC", 16, 8, ReadWrite, "Line end delimiter code">,
      Field<"FEC", 24, 8, ReadWrite, "Frame end delimiter code">,
    ], "embedded synchronization code register">,
    Register<"ESUR", 0x1c, 32, [
      Field<"FSU", 0, 8, ReadWrite, "Frame start delimiter unmask">,
      Field<"LSU", 8, 8, ReadWrite, "Line start delimiter unmask">,
      Field<"LEU", 16, 8, ReadWrite, "Line end delimiter unmask">,
      Field<"FEU", 24, 8, ReadWrite, "Frame end delimiter unmask">,
    ], "embedded synchronization unmask register">,
    Register<"CWSTRT", 0x20, 32, [
      Field<"HOFFCNT", 0, 14, ReadWrite, "Horizontal offset count">,
      Field<"VST", 16, 13, ReadWrite, "Vertical start line count">,
    ], "crop window start">,
    Register<"CWSIZE", 0x24, 32, [
      Field<"CAPCNT", 0, 14, ReadWrite, "Capture count">,
      Field<"VLINE", 16, 14, ReadWrite, "Vertical line count">,
    ], "crop window size">,
    Register<"DR", 0x28, 32, [
      Field<"BYTE0", 0, 8, ReadWrite, "Data byte 0">,
      Field<"BYTE1", 8, 8, ReadWrite, "Data byte 1">,
      Field<"BYTE2", 16, 8, ReadWrite, "Data byte 2">,
      Field<"BYTE3", 24, 8, ReadWrite, "Data byte 3">,
    ], "data register">,
  ];
}

class DCMIInstance<string Name, int Base> : PeripheralInstance<Name, Base, DCMIPeripheral>;
def DCMI : DCMIInstance<"DCMI", 0x48020000>;

#endif // _PERIPHERALS_DCMI_TD
