#ifndef _PERIPHERALS_LTDC_TD
#define _PERIPHERALS_LTDC_TD

include "base.td"

def LTDCPeripheral : PeripheralType<"LTDC", "LCD-TFT Controller"> {
  let accessWidth = 32;
  let registers = [
    Register<"SSCR", 0x8, 32, [
      Field<"HSW", 16, 10, ReadWrite, "Horizontal Synchronization Width (in units of pixel clock period)">,
      Field<"VSH", 0, 11, ReadWrite, "Vertical Synchronization Height (in units of horizontal scan line)">,
    ], "Synchronization Size Configuration Register">,
    Register<"BPCR", 0xc, 32, [
      Field<"AHBP", 16, 12, ReadWrite, "Accumulated Horizontal back porch (in units of pixel clock period)">,
      Field<"AVBP", 0, 11, ReadWrite, "Accumulated Vertical back porch (in units of horizontal scan line)">,
    ], "Back Porch Configuration Register">,
    Register<"AWCR", 0x10, 32, [
      Field<"AAV", 16, 12, ReadWrite, "AAV">,
      Field<"AAH", 0, 11, ReadWrite, "Accumulated Active Height (in units of horizontal scan line)">,
    ], "Active Width Configuration Register">,
    Register<"TWCR", 0x14, 32, [
      Field<"TOTALW", 16, 12, ReadWrite, "Total Width (in units of pixel clock period)">,
      Field<"TOTALH", 0, 11, ReadWrite, "Total Height (in units of horizontal scan line)">,
    ], "Total Width Configuration Register">,
    Register<"GCR", 0x18, 32, [
      Field<"HSPOL", 31, 1, ReadWrite, "Horizontal Synchronization Polarity">,
      Field<"VSPOL", 30, 1, ReadWrite, "Vertical Synchronization Polarity">,
      Field<"DEPOL", 29, 1, ReadWrite, "Data Enable Polarity">,
      Field<"PCPOL", 28, 1, ReadWrite, "Pixel Clock Polarity">,
      Field<"DEN", 16, 1, ReadWrite, "Dither Enable">,
      Field<"DRW", 12, 3, Read, "Dither Red Width">,
      Field<"DGW", 8, 3, Read, "Dither Green Width">,
      Field<"DBW", 4, 3, Read, "Dither Blue Width">,
      Field<"LTDCEN", 0, 1, ReadWrite, "LCD-TFT controller enable bit">,
    ], "Global Control Register">,
    Register<"SRCR", 0x24, 32, [
      Field<"VBR", 1, 1, ReadWrite, "Vertical Blanking Reload">,
      Field<"IMR", 0, 1, ReadWrite, "Immediate Reload">,
    ], "Shadow Reload Configuration Register">,
    Register<"BCCR", 0x2c, 32, [
      Field<"BCBLUE", 0, 8, ReadWrite, "Background Color Blue value">,
      Field<"BCGREEN", 8, 8, ReadWrite, "Background Color Green value">,
      Field<"BCRED", 16, 8, ReadWrite, "Background Color Red value">,
    ], "Background Color Configuration Register">,
    Register<"IER", 0x34, 32, [
      Field<"RRIE", 3, 1, ReadWrite, "Register Reload interrupt enable">,
      Field<"TERRIE", 2, 1, ReadWrite, "Transfer Error Interrupt Enable">,
      Field<"FUIE", 1, 1, ReadWrite, "FIFO Underrun Interrupt Enable">,
      Field<"LIE", 0, 1, ReadWrite, "Line Interrupt Enable">,
    ], "Interrupt Enable Register">,
    Register<"ISR", 0x38, 32, [
      Field<"RRIF", 3, 1, ReadWrite, "Register Reload Interrupt Flag">,
      Field<"TERRIF", 2, 1, ReadWrite, "Transfer Error interrupt flag">,
      Field<"FUIF", 1, 1, ReadWrite, "FIFO Underrun Interrupt flag">,
      Field<"LIF", 0, 1, ReadWrite, "Line Interrupt flag">,
    ], "Interrupt Status Register">,
    Register<"ICR", 0x3c, 32, [
      Field<"CRRIF", 3, 1, ReadWrite, "Clears Register Reload Interrupt Flag">,
      Field<"CTERRIF", 2, 1, ReadWrite, "Clears the Transfer Error Interrupt Flag">,
      Field<"CFUIF", 1, 1, ReadWrite, "Clears the FIFO Underrun Interrupt flag">,
      Field<"CLIF", 0, 1, ReadWrite, "Clears the Line Interrupt Flag">,
    ], "Interrupt Clear Register">,
    Register<"LIPCR", 0x40, 32, [
      Field<"LIPOS", 0, 11, ReadWrite, "Line Interrupt Position">,
    ], "Line Interrupt Position Configuration Register">,
    Register<"CPSR", 0x44, 32, [
      Field<"CXPOS", 16, 16, ReadWrite, "Current X Position">,
      Field<"CYPOS", 0, 16, ReadWrite, "Current Y Position">,
    ], "Current Position Status Register">,
    Register<"CDSR", 0x48, 32, [
      Field<"HSYNCS", 3, 1, ReadWrite, "Horizontal Synchronization display Status">,
      Field<"VSYNCS", 2, 1, ReadWrite, "Vertical Synchronization display Status">,
      Field<"HDES", 1, 1, ReadWrite, "Horizontal Data Enable display Status">,
      Field<"VDES", 0, 1, ReadWrite, "Vertical Data Enable display Status">,
    ], "Current Display Status Register">,
    Register<"L1CR", 0x84, 32, [
      Field<"CLUTEN", 4, 1, ReadWrite, "Color Look-Up Table Enable">,
      Field<"COLKEN", 1, 1, ReadWrite, "Color Keying Enable">,
      Field<"LEN", 0, 1, ReadWrite, "Layer Enable">,
    ], "Layerx Control Register">,
    Register<"L1WHPCR", 0x88, 32, [
      Field<"WHSPPOS", 16, 12, ReadWrite, "Window Horizontal Stop Position">,
      Field<"WHSTPOS", 0, 12, ReadWrite, "Window Horizontal Start Position">,
    ], "Layerx Window Horizontal Position Configuration Register">,
    Register<"L1WVPCR", 0x8c, 32, [
      Field<"WVSPPOS", 16, 11, ReadWrite, "Window Vertical Stop Position">,
      Field<"WVSTPOS", 0, 11, ReadWrite, "Window Vertical Start Position">,
    ], "Layerx Window Vertical Position Configuration Register">,
    Register<"L1CKCR", 0x90, 32, [
      Field<"CKRED", 16, 8, ReadWrite, "Color Key Red value">,
      Field<"CKGREEN", 8, 8, ReadWrite, "Color Key Green value">,
      Field<"CKBLUE", 0, 8, ReadWrite, "Color Key Blue value">,
    ], "Layerx Color Keying Configuration Register">,
    Register<"L1PFCR", 0x94, 32, [
      Field<"PF", 0, 3, ReadWrite, "Pixel Format">,
    ], "Layerx Pixel Format Configuration Register">,
    Register<"L1CACR", 0x98, 32, [
      Field<"CONSTA", 0, 8, ReadWrite, "Constant Alpha">,
    ], "Layerx Constant Alpha Configuration Register">,
    Register<"L1DCCR", 0x9c, 32, [
      Field<"DCALPHA", 24, 8, ReadWrite, "Default Color Alpha">,
      Field<"DCRED", 16, 8, ReadWrite, "Default Color Red">,
      Field<"DCGREEN", 8, 8, ReadWrite, "Default Color Green">,
      Field<"DCBLUE", 0, 8, ReadWrite, "Default Color Blue">,
    ], "Layerx Default Color Configuration Register">,
    Register<"L1BFCR", 0xa0, 32, [
      Field<"BF1", 8, 3, ReadWrite, "Blending Factor 1">,
      Field<"BF2", 0, 3, ReadWrite, "Blending Factor 2">,
    ], "Layerx Blending Factors Configuration Register">,
    Register<"L1CFBAR", 0xac, 32, [
      Field<"CFBADD", 0, 32, ReadWrite, "Color Frame Buffer Start Address">,
    ], "Layerx Color Frame Buffer Address Register">,
    Register<"L1CFBLR", 0xb0, 32, [
      Field<"CFBP", 16, 13, ReadWrite, "Color Frame Buffer Pitch in bytes">,
      Field<"CFBLL", 0, 13, ReadWrite, "Color Frame Buffer Line Length">,
    ], "Layerx Color Frame Buffer Length Register">,
    Register<"L1CFBLNR", 0xb4, 32, [
      Field<"CFBLNBR", 0, 11, ReadWrite, "Frame Buffer Line Number">,
    ], "Layerx ColorFrame Buffer Line Number Register">,
    Register<"L1CLUTWR", 0xc4, 32, [
      Field<"CLUTADD", 24, 8, ReadWrite, "CLUT Address">,
      Field<"RED", 16, 8, ReadWrite, "Red value">,
      Field<"GREEN", 8, 8, ReadWrite, "Green value">,
      Field<"BLUE", 0, 8, ReadWrite, "Blue value">,
    ], "Layerx CLUT Write Register">,
    Register<"L2CR", 0x104, 32, [
      Field<"CLUTEN", 4, 1, ReadWrite, "Color Look-Up Table Enable">,
      Field<"COLKEN", 1, 1, ReadWrite, "Color Keying Enable">,
      Field<"LEN", 0, 1, ReadWrite, "Layer Enable">,
    ], "Layerx Control Register">,
    Register<"L2WHPCR", 0x108, 32, [
      Field<"WHSPPOS", 16, 12, ReadWrite, "Window Horizontal Stop Position">,
      Field<"WHSTPOS", 0, 12, ReadWrite, "Window Horizontal Start Position">,
    ], "Layerx Window Horizontal Position Configuration Register">,
    Register<"L2WVPCR", 0x10c, 32, [
      Field<"WVSPPOS", 16, 11, ReadWrite, "Window Vertical Stop Position">,
      Field<"WVSTPOS", 0, 11, ReadWrite, "Window Vertical Start Position">,
    ], "Layerx Window Vertical Position Configuration Register">,
    Register<"L2CKCR", 0x110, 32, [
      Field<"CKRED", 16, 8, ReadWrite, "Color Key Red value">,
      Field<"CKGREEN", 8, 8, ReadWrite, "Color Key Green value">,
      Field<"CKBLUE", 0, 8, ReadWrite, "Color Key Blue value">,
    ], "Layerx Color Keying Configuration Register">,
    Register<"L2PFCR", 0x114, 32, [
      Field<"PF", 0, 3, ReadWrite, "Pixel Format">,
    ], "Layerx Pixel Format Configuration Register">,
    Register<"L2CACR", 0x118, 32, [
      Field<"CONSTA", 0, 8, ReadWrite, "Constant Alpha">,
    ], "Layerx Constant Alpha Configuration Register">,
    Register<"L2DCCR", 0x11c, 32, [
      Field<"DCALPHA", 24, 8, ReadWrite, "Default Color Alpha">,
      Field<"DCRED", 16, 8, ReadWrite, "Default Color Red">,
      Field<"DCGREEN", 8, 8, ReadWrite, "Default Color Green">,
      Field<"DCBLUE", 0, 8, ReadWrite, "Default Color Blue">,
    ], "Layerx Default Color Configuration Register">,
    Register<"L2BFCR", 0x120, 32, [
      Field<"BF1", 8, 3, ReadWrite, "Blending Factor 1">,
      Field<"BF2", 0, 3, ReadWrite, "Blending Factor 2">,
    ], "Layerx Blending Factors Configuration Register">,
    Register<"L2CFBAR", 0x12c, 32, [
      Field<"CFBADD", 0, 32, ReadWrite, "Color Frame Buffer Start Address">,
    ], "Layerx Color Frame Buffer Address Register">,
    Register<"L2CFBLR", 0x130, 32, [
      Field<"CFBP", 16, 13, ReadWrite, "Color Frame Buffer Pitch in bytes">,
      Field<"CFBLL", 0, 13, ReadWrite, "Color Frame Buffer Line Length">,
    ], "Layerx Color Frame Buffer Length Register">,
    Register<"L2CFBLNR", 0x134, 32, [
      Field<"CFBLNBR", 0, 11, ReadWrite, "Frame Buffer Line Number">,
    ], "Layerx ColorFrame Buffer Line Number Register">,
    Register<"L2CLUTWR", 0x144, 32, [
      Field<"CLUTADD", 24, 8, ReadWrite, "CLUT Address">,
      Field<"RED", 16, 8, ReadWrite, "Red value">,
      Field<"GREEN", 8, 8, ReadWrite, "Green value">,
      Field<"BLUE", 0, 8, ReadWrite, "Blue value">,
    ], "Layerx CLUT Write Register">,
  ];
}

class LTDCInstance<string Name, int Base> : PeripheralInstance<Name, Base, LTDCPeripheral>;
def LTDC : LTDCInstance<"LTDC", 0x50001000>;

#endif // _PERIPHERALS_LTDC_TD
