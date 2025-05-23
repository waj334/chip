#ifndef _PERIPHERALS_LTDC_TD
#define _PERIPHERALS_LTDC_TD

include "base.td"

def LTDCPeripheral : PeripheralType<"LTDC", "LCD-TFT Controller"> {
  let accessWidth = 32;
  let registers = [
    Register<"SSCR", 0x8, 32, [
      Field<"VSH", 0, 11, ReadWrite, "Vertical Synchronization Height (in units of horizontal scan line)">,
      Field<"HSW", 16, 10, ReadWrite, "Horizontal Synchronization Width (in units of pixel clock period)">,
    ], "Synchronization Size Configuration Register">,
    Register<"BPCR", 0xc, 32, [
      Field<"AVBP", 0, 11, ReadWrite, "Accumulated Vertical back porch (in units of horizontal scan line)">,
      Field<"AHBP", 16, 12, ReadWrite, "Accumulated Horizontal back porch (in units of pixel clock period)">,
    ], "Back Porch Configuration Register">,
    Register<"AWCR", 0x10, 32, [
      Field<"AAH", 0, 11, ReadWrite, "Accumulated Active Height (in units of horizontal scan line)">,
      Field<"AAV", 16, 12, ReadWrite, "AAV">,
    ], "Active Width Configuration Register">,
    Register<"TWCR", 0x14, 32, [
      Field<"TOTALH", 0, 11, ReadWrite, "Total Height (in units of horizontal scan line)">,
      Field<"TOTALW", 16, 12, ReadWrite, "Total Width (in units of pixel clock period)">,
    ], "Total Width Configuration Register">,
    Register<"GCR", 0x18, 32, [
      Field<"LTDCEN", 0, 1, ReadWrite, "LCD-TFT controller enable bit">,
      Field<"DBW", 4, 3, Read, "Dither Blue Width">,
      Field<"DGW", 8, 3, Read, "Dither Green Width">,
      Field<"DRW", 12, 3, Read, "Dither Red Width">,
      Field<"DEN", 16, 1, ReadWrite, "Dither Enable">,
      Field<"PCPOL", 28, 1, ReadWrite, "Pixel Clock Polarity">,
      Field<"DEPOL", 29, 1, ReadWrite, "Data Enable Polarity">,
      Field<"VSPOL", 30, 1, ReadWrite, "Vertical Synchronization Polarity">,
      Field<"HSPOL", 31, 1, ReadWrite, "Horizontal Synchronization Polarity">,
    ], "Global Control Register">,
    Register<"SRCR", 0x24, 32, [
      Field<"IMR", 0, 1, ReadWrite, "Immediate Reload">,
      Field<"VBR", 1, 1, ReadWrite, "Vertical Blanking Reload">,
    ], "Shadow Reload Configuration Register">,
    Register<"BCCR", 0x2c, 32, [
      Field<"BCBLUE", 0, 8, ReadWrite, "Background Color Blue value">,
      Field<"BCGREEN", 8, 8, ReadWrite, "Background Color Green value">,
      Field<"BCRED", 16, 8, ReadWrite, "Background Color Red value">,
    ], "Background Color Configuration Register">,
    Register<"IER", 0x34, 32, [
      Field<"LIE", 0, 1, ReadWrite, "Line Interrupt Enable">,
      Field<"FUIE", 1, 1, ReadWrite, "FIFO Underrun Interrupt Enable">,
      Field<"TERRIE", 2, 1, ReadWrite, "Transfer Error Interrupt Enable">,
      Field<"RRIE", 3, 1, ReadWrite, "Register Reload interrupt enable">,
    ], "Interrupt Enable Register">,
    Register<"ISR", 0x38, 32, [
      Field<"LIF", 0, 1, ReadWrite, "Line Interrupt flag">,
      Field<"FUIF", 1, 1, ReadWrite, "FIFO Underrun Interrupt flag">,
      Field<"TERRIF", 2, 1, ReadWrite, "Transfer Error interrupt flag">,
      Field<"RRIF", 3, 1, ReadWrite, "Register Reload Interrupt Flag">,
    ], "Interrupt Status Register">,
    Register<"ICR", 0x3c, 32, [
      Field<"CLIF", 0, 1, ReadWrite, "Clears the Line Interrupt Flag">,
      Field<"CFUIF", 1, 1, ReadWrite, "Clears the FIFO Underrun Interrupt flag">,
      Field<"CTERRIF", 2, 1, ReadWrite, "Clears the Transfer Error Interrupt Flag">,
      Field<"CRRIF", 3, 1, ReadWrite, "Clears Register Reload Interrupt Flag">,
    ], "Interrupt Clear Register">,
    Register<"LIPCR", 0x40, 32, [
      Field<"LIPOS", 0, 11, ReadWrite, "Line Interrupt Position">,
    ], "Line Interrupt Position Configuration Register">,
    Register<"CPSR", 0x44, 32, [
      Field<"CYPOS", 0, 16, ReadWrite, "Current Y Position">,
      Field<"CXPOS", 16, 16, ReadWrite, "Current X Position">,
    ], "Current Position Status Register">,
    Register<"CDSR", 0x48, 32, [
      Field<"VDES", 0, 1, ReadWrite, "Vertical Data Enable display Status">,
      Field<"HDES", 1, 1, ReadWrite, "Horizontal Data Enable display Status">,
      Field<"VSYNCS", 2, 1, ReadWrite, "Vertical Synchronization display Status">,
      Field<"HSYNCS", 3, 1, ReadWrite, "Horizontal Synchronization display Status">,
    ], "Current Display Status Register">,
    Register<"L1CR", 0x84, 32, [
      Field<"LEN", 0, 1, ReadWrite, "Layer Enable">,
      Field<"COLKEN", 1, 1, ReadWrite, "Color Keying Enable">,
      Field<"CLUTEN", 4, 1, ReadWrite, "Color Look-Up Table Enable">,
    ], "Layerx Control Register">,
    Register<"L1WHPCR", 0x88, 32, [
      Field<"WHSTPOS", 0, 12, ReadWrite, "Window Horizontal Start Position">,
      Field<"WHSPPOS", 16, 12, ReadWrite, "Window Horizontal Stop Position">,
    ], "Layerx Window Horizontal Position Configuration Register">,
    Register<"L1WVPCR", 0x8c, 32, [
      Field<"WVSTPOS", 0, 11, ReadWrite, "Window Vertical Start Position">,
      Field<"WVSPPOS", 16, 11, ReadWrite, "Window Vertical Stop Position">,
    ], "Layerx Window Vertical Position Configuration Register">,
    Register<"L1CKCR", 0x90, 32, [
      Field<"CKBLUE", 0, 8, ReadWrite, "Color Key Blue value">,
      Field<"CKGREEN", 8, 8, ReadWrite, "Color Key Green value">,
      Field<"CKRED", 16, 8, ReadWrite, "Color Key Red value">,
    ], "Layerx Color Keying Configuration Register">,
    Register<"L1PFCR", 0x94, 32, [
      Field<"PF", 0, 3, ReadWrite, "Pixel Format">,
    ], "Layerx Pixel Format Configuration Register">,
    Register<"L1CACR", 0x98, 32, [
      Field<"CONSTA", 0, 8, ReadWrite, "Constant Alpha">,
    ], "Layerx Constant Alpha Configuration Register">,
    Register<"L1DCCR", 0x9c, 32, [
      Field<"DCBLUE", 0, 8, ReadWrite, "Default Color Blue">,
      Field<"DCGREEN", 8, 8, ReadWrite, "Default Color Green">,
      Field<"DCRED", 16, 8, ReadWrite, "Default Color Red">,
      Field<"DCALPHA", 24, 8, ReadWrite, "Default Color Alpha">,
    ], "Layerx Default Color Configuration Register">,
    Register<"L1BFCR", 0xa0, 32, [
      Field<"BF2", 0, 3, ReadWrite, "Blending Factor 2">,
      Field<"BF1", 8, 3, ReadWrite, "Blending Factor 1">,
    ], "Layerx Blending Factors Configuration Register">,
    Register<"L1CFBAR", 0xac, 32, [
      Field<"CFBADD", 0, 32, ReadWrite, "Color Frame Buffer Start Address">,
    ], "Layerx Color Frame Buffer Address Register">,
    Register<"L1CFBLR", 0xb0, 32, [
      Field<"CFBLL", 0, 13, ReadWrite, "Color Frame Buffer Line Length">,
      Field<"CFBP", 16, 13, ReadWrite, "Color Frame Buffer Pitch in bytes">,
    ], "Layerx Color Frame Buffer Length Register">,
    Register<"L1CFBLNR", 0xb4, 32, [
      Field<"CFBLNBR", 0, 11, ReadWrite, "Frame Buffer Line Number">,
    ], "Layerx ColorFrame Buffer Line Number Register">,
    Register<"L1CLUTWR", 0xc4, 32, [
      Field<"BLUE", 0, 8, ReadWrite, "Blue value">,
      Field<"GREEN", 8, 8, ReadWrite, "Green value">,
      Field<"RED", 16, 8, ReadWrite, "Red value">,
      Field<"CLUTADD", 24, 8, ReadWrite, "CLUT Address">,
    ], "Layerx CLUT Write Register">,
    Register<"L2CR", 0x104, 32, [
      Field<"LEN", 0, 1, ReadWrite, "Layer Enable">,
      Field<"COLKEN", 1, 1, ReadWrite, "Color Keying Enable">,
      Field<"CLUTEN", 4, 1, ReadWrite, "Color Look-Up Table Enable">,
    ], "Layerx Control Register">,
    Register<"L2WHPCR", 0x108, 32, [
      Field<"WHSTPOS", 0, 12, ReadWrite, "Window Horizontal Start Position">,
      Field<"WHSPPOS", 16, 12, ReadWrite, "Window Horizontal Stop Position">,
    ], "Layerx Window Horizontal Position Configuration Register">,
    Register<"L2WVPCR", 0x10c, 32, [
      Field<"WVSTPOS", 0, 11, ReadWrite, "Window Vertical Start Position">,
      Field<"WVSPPOS", 16, 11, ReadWrite, "Window Vertical Stop Position">,
    ], "Layerx Window Vertical Position Configuration Register">,
    Register<"L2CKCR", 0x110, 32, [
      Field<"CKBLUE", 0, 8, ReadWrite, "Color Key Blue value">,
      Field<"CKGREEN", 8, 8, ReadWrite, "Color Key Green value">,
      Field<"CKRED", 16, 8, ReadWrite, "Color Key Red value">,
    ], "Layerx Color Keying Configuration Register">,
    Register<"L2PFCR", 0x114, 32, [
      Field<"PF", 0, 3, ReadWrite, "Pixel Format">,
    ], "Layerx Pixel Format Configuration Register">,
    Register<"L2CACR", 0x118, 32, [
      Field<"CONSTA", 0, 8, ReadWrite, "Constant Alpha">,
    ], "Layerx Constant Alpha Configuration Register">,
    Register<"L2DCCR", 0x11c, 32, [
      Field<"DCBLUE", 0, 8, ReadWrite, "Default Color Blue">,
      Field<"DCGREEN", 8, 8, ReadWrite, "Default Color Green">,
      Field<"DCRED", 16, 8, ReadWrite, "Default Color Red">,
      Field<"DCALPHA", 24, 8, ReadWrite, "Default Color Alpha">,
    ], "Layerx Default Color Configuration Register">,
    Register<"L2BFCR", 0x120, 32, [
      Field<"BF2", 0, 3, ReadWrite, "Blending Factor 2">,
      Field<"BF1", 8, 3, ReadWrite, "Blending Factor 1">,
    ], "Layerx Blending Factors Configuration Register">,
    Register<"L2CFBAR", 0x12c, 32, [
      Field<"CFBADD", 0, 32, ReadWrite, "Color Frame Buffer Start Address">,
    ], "Layerx Color Frame Buffer Address Register">,
    Register<"L2CFBLR", 0x130, 32, [
      Field<"CFBLL", 0, 13, ReadWrite, "Color Frame Buffer Line Length">,
      Field<"CFBP", 16, 13, ReadWrite, "Color Frame Buffer Pitch in bytes">,
    ], "Layerx Color Frame Buffer Length Register">,
    Register<"L2CFBLNR", 0x134, 32, [
      Field<"CFBLNBR", 0, 11, ReadWrite, "Frame Buffer Line Number">,
    ], "Layerx ColorFrame Buffer Line Number Register">,
    Register<"L2CLUTWR", 0x144, 32, [
      Field<"BLUE", 0, 8, ReadWrite, "Blue value">,
      Field<"GREEN", 8, 8, ReadWrite, "Green value">,
      Field<"RED", 16, 8, ReadWrite, "Red value">,
      Field<"CLUTADD", 24, 8, ReadWrite, "CLUT Address">,
    ], "Layerx CLUT Write Register">,
  ];
}

class LTDCInstance<string Name, int Base> : PeripheralInstance<Name, Base, LTDCPeripheral>;
def LTDC : LTDCInstance<"LTDC", 0x50001000>;

#endif // _PERIPHERALS_LTDC_TD
