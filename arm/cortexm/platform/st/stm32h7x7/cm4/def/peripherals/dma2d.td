#ifndef _PERIPHERALS_DMA2D_TD
#define _PERIPHERALS_DMA2D_TD

include "base.td"

def DMA2DPeripheral : PeripheralType<"DMA2D", "DMA2D"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"START", 0, 1, ReadWrite, "Start This bit can be used to launch the DMA2D according to the parameters loaded in the various configuration registers">,
      Field<"SUSP", 1, 1, ReadWrite, "Suspend This bit can be used to suspend the current transfer. This bit is set and reset by software. It is automatically reset by hardware when the START bit is reset.">,
      Field<"ABORT", 2, 1, ReadWrite, "Abort This bit can be used to abort the current transfer. This bit is set by software and is automatically reset by hardware when the START bit is reset.">,
      Field<"TEIE", 8, 1, ReadWrite, "Transfer error interrupt enable This bit is set and cleared by software.">,
      Field<"TCIE", 9, 1, ReadWrite, "Transfer complete interrupt enable This bit is set and cleared by software.">,
      Field<"TWIE", 10, 1, ReadWrite, "Transfer watermark interrupt enable This bit is set and cleared by software.">,
      Field<"CAEIE", 11, 1, ReadWrite, "CLUT access error interrupt enable This bit is set and cleared by software.">,
      Field<"CTCIE", 12, 1, ReadWrite, "CLUT transfer complete interrupt enable This bit is set and cleared by software.">,
      Field<"CEIE", 13, 1, ReadWrite, "Configuration Error Interrupt Enable This bit is set and cleared by software.">,
      Field<"MODE", 16, 2, ReadWrite, "DMA2D mode This bit is set and cleared by software. It cannot be modified while a transfer is ongoing.">,
    ], "DMA2D control register">,
    Register<"ISR", 0x4, 32, [
      Field<"TEIF", 0, 1, ReadWrite, "Transfer error interrupt flag This bit is set when an error occurs during a DMA transfer (data transfer or automatic CLUT loading).">,
      Field<"TCIF", 1, 1, ReadWrite, "Transfer complete interrupt flag This bit is set when a DMA2D transfer operation is complete (data transfer only).">,
      Field<"TWIF", 2, 1, ReadWrite, "Transfer watermark interrupt flag This bit is set when the last pixel of the watermarked line has been transferred.">,
      Field<"CAEIF", 3, 1, ReadWrite, "CLUT access error interrupt flag This bit is set when the CPU accesses the CLUT while the CLUT is being automatically copied from a system memory to the internal DMA2D.">,
      Field<"CTCIF", 4, 1, ReadWrite, "CLUT transfer complete interrupt flag This bit is set when the CLUT copy from a system memory area to the internal DMA2D memory is complete.">,
      Field<"CEIF", 5, 1, ReadWrite, "Configuration error interrupt flag This bit is set when the START bit of DMA2D_CR, DMA2DFGPFCCR or DMA2D_BGPFCCR is set and a wrong configuration has been programmed.">,
    ], "DMA2D Interrupt Status Register">,
    Register<"IFCR", 0x8, 32, [
      Field<"CTEIF", 0, 1, ReadWrite, "Clear Transfer error interrupt flag Programming this bit to 1 clears the TEIF flag in the DMA2D_ISR register">,
      Field<"CTCIF", 1, 1, ReadWrite, "Clear transfer complete interrupt flag Programming this bit to 1 clears the TCIF flag in the DMA2D_ISR register">,
      Field<"CTWIF", 2, 1, ReadWrite, "Clear transfer watermark interrupt flag Programming this bit to 1 clears the TWIF flag in the DMA2D_ISR register">,
      Field<"CAECIF", 3, 1, ReadWrite, "Clear CLUT access error interrupt flag Programming this bit to 1 clears the CAEIF flag in the DMA2D_ISR register">,
      Field<"CCTCIF", 4, 1, ReadWrite, "Clear CLUT transfer complete interrupt flag Programming this bit to 1 clears the CTCIF flag in the DMA2D_ISR register">,
      Field<"CCEIF", 5, 1, ReadWrite, "Clear configuration error interrupt flag Programming this bit to 1 clears the CEIF flag in the DMA2D_ISR register">,
    ], "DMA2D interrupt flag clear register">,
    Register<"FGMAR", 0xc, 32, [
      Field<"MA", 0, 32, ReadWrite, "Memory address Address of the data used for the foreground image. This register can only be written when data transfers are disabled. Once the data transfer has started, this register is read-only. The address alignment must match the image format selected e.g. a 32-bit per pixel format must be 32-bit aligned, a 16-bit per pixel format must be 16-bit aligned and a 4-bit per pixel format must be 8-bit aligned.">,
    ], "DMA2D foreground memory address register">,
    Register<"FGOR", 0x10, 32, [
      Field<"LO", 0, 14, ReadWrite, "Line offset Line offset used for the foreground expressed in pixel. This value is used to generate the address. It is added at the end of each line to determine the starting address of the next line. These bits can only be written when data transfers are disabled. Once a data transfer has started, they become read-only. If the image format is 4-bit per pixel, the line offset must be even.">,
    ], "DMA2D foreground offset register">,
    Register<"BGMAR", 0x14, 32, [
      Field<"MA", 0, 32, ReadWrite, "Memory address Address of the data used for the background image. This register can only be written when data transfers are disabled. Once a data transfer has started, this register is read-only. The address alignment must match the image format selected e.g. a 32-bit per pixel format must be 32-bit aligned, a 16-bit per pixel format must be 16-bit aligned and a 4-bit per pixel format must be 8-bit aligned.">,
    ], "DMA2D background memory address register">,
    Register<"BGOR", 0x18, 32, [
      Field<"LO", 0, 14, ReadWrite, "Line offset Line offset used for the background image (expressed in pixel). This value is used for the address generation. It is added at the end of each line to determine the starting address of the next line. These bits can only be written when data transfers are disabled. Once data transfer has started, they become read-only. If the image format is 4-bit per pixel, the line offset must be even.">,
    ], "DMA2D background offset register">,
    Register<"FGPFCCR", 0x1c, 32, [
      Field<"CM", 0, 4, ReadWrite, "Color mode These bits defines the color format of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless">,
      Field<"CCM", 4, 1, ReadWrite, "CLUT color mode This bit defines the color format of the CLUT. It can only be written when the transfer is disabled. Once the CLUT transfer has started, this bit is read-only.">,
      Field<"START", 5, 1, ReadWrite, "Start This bit can be set to start the automatic loading of the CLUT. It is automatically reset: ** at the end of the transfer ** when the transfer is aborted by the user application by setting the ABORT bit in DMA2D_CR ** when a transfer error occurs ** when the transfer has not started due to a configuration error or another transfer operation already ongoing (data transfer or automatic background CLUT transfer).">,
      Field<"CS", 8, 8, ReadWrite, "CLUT size These bits define the size of the CLUT used for the foreground image. Once the CLUT transfer has started, this field is read-only. The number of CLUT entries is equal to CS[7:0] + 1.">,
      Field<"AM", 16, 2, ReadWrite, "Alpha mode These bits select the alpha channel value to be used for the foreground image. They can only be written data the transfer are disabled. Once the transfer has started, they become read-only. other configurations are meaningless">,
      Field<"CSS", 18, 2, ReadWrite, "Chroma Sub-Sampling These bits define the chroma sub-sampling mode for YCbCr color mode. Once the transfer has started, these bits are read-only. others: meaningless">,
      Field<"AI", 20, 1, ReadWrite, "Alpha Inverted This bit inverts the alpha value. Once the transfer has started, this bit is read-only.">,
      Field<"RBS", 21, 1, ReadWrite, "Red Blue Swap This bit allows to swap the R &amp; B to support BGR or ABGR color formats. Once the transfer has started, this bit is read-only.">,
      Field<"ALPHA", 24, 8, ReadWrite, "Alpha value These bits define a fixed alpha channel value which can replace the original alpha value or be multiplied by the original alpha value according to the alpha mode selected through the AM[1:0] bits. These bits can only be written when data transfers are disabled. Once a transfer has started, they become read-only.">,
    ], "DMA2D foreground PFC control register">,
    Register<"FGCOLR", 0x20, 32, [
      Field<"BLUE", 0, 8, ReadWrite, "Blue Value These bits defines the blue value for the A4 or A8 mode of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, They are read-only.">,
      Field<"GREEN", 8, 8, ReadWrite, "Green Value These bits defines the green value for the A4 or A8 mode of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, They are read-only.">,
      Field<"RED", 16, 8, ReadWrite, "Red Value These bits defines the red value for the A4 or A8 mode of the foreground image. They can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
    ], "DMA2D foreground color register">,
    Register<"BGPFCCR", 0x24, 32, [
      Field<"CM", 0, 4, ReadWrite, "Color mode These bits define the color format of the foreground image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless">,
      Field<"CCM", 4, 1, ReadWrite, "CLUT Color mode These bits define the color format of the CLUT. This register can only be written when the transfer is disabled. Once the CLUT transfer has started, this bit is read-only.">,
      Field<"START", 5, 1, ReadWrite, "Start This bit is set to start the automatic loading of the CLUT. This bit is automatically reset: ** at the end of the transfer ** when the transfer is aborted by the user application by setting the ABORT bit in the DMA2D_CR ** when a transfer error occurs ** when the transfer has not started due to a configuration error or another transfer operation already on going (data transfer or automatic BackGround CLUT transfer).">,
      Field<"CS", 8, 8, ReadWrite, "CLUT size These bits define the size of the CLUT used for the BG. Once the CLUT transfer has started, this field is read-only. The number of CLUT entries is equal to CS[7:0] + 1.">,
      Field<"AM", 16, 2, ReadWrite, "Alpha mode These bits define which alpha channel value to be used for the background image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless">,
      Field<"AI", 20, 1, ReadWrite, "Alpha Inverted This bit inverts the alpha value. Once the transfer has started, this bit is read-only.">,
      Field<"RBS", 21, 1, ReadWrite, "Red Blue Swap This bit allows to swap the R &amp; B to support BGR or ABGR color formats. Once the transfer has started, this bit is read-only.">,
      Field<"ALPHA", 24, 8, ReadWrite, "Alpha value These bits define a fixed alpha channel value which can replace the original alpha value or be multiplied with the original alpha value according to the alpha mode selected with bits AM[1: 0]. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
    ], "DMA2D background PFC control register">,
    Register<"BGCOLR", 0x28, 32, [
      Field<"BLUE", 0, 8, ReadWrite, "Blue Value These bits define the blue value for the A4 or A8 mode of the background. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
      Field<"GREEN", 8, 8, ReadWrite, "Green Value These bits define the green value for the A4 or A8 mode of the background. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
      Field<"RED", 16, 8, ReadWrite, "Red Value These bits define the red value for the A4 or A8 mode of the background. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
    ], "DMA2D background color register">,
    Register<"FGCMAR", 0x2c, 32, [
      Field<"MA", 0, 32, ReadWrite, "Memory Address Address of the data used for the CLUT address dedicated to the foreground image. This register can only be written when no transfer is ongoing. Once the CLUT transfer has started, this register is read-only. If the foreground CLUT format is 32-bit, the address must be 32-bit aligned.">,
    ], "DMA2D foreground CLUT memory address register">,
    Register<"BGCMAR", 0x30, 32, [
      Field<"MA", 0, 32, ReadWrite, "Memory address Address of the data used for the CLUT address dedicated to the background image. This register can only be written when no transfer is on going. Once the CLUT transfer has started, this register is read-only. If the background CLUT format is 32-bit, the address must be 32-bit aligned.">,
    ], "DMA2D background CLUT memory address register">,
    Register<"OPFCCR", 0x34, 32, [
      Field<"CM", 0, 3, ReadWrite, "Color mode These bits define the color format of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. others: meaningless">,
      Field<"AI", 20, 1, ReadWrite, "Alpha Inverted This bit inverts the alpha value. Once the transfer has started, this bit is read-only.">,
      Field<"RBS", 21, 1, ReadWrite, "Red Blue Swap This bit allows to swap the R &amp; B to support BGR or ABGR color formats. Once the transfer has started, this bit is read-only.">,
    ], "DMA2D output PFC control register">,
    Register<"OCOLR", 0x38, 32, [
      Field<"BLUE", 0, 8, ReadWrite, "Blue Value These bits define the blue value of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
      Field<"GREEN", 8, 8, ReadWrite, "Green Value These bits define the green value of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
      Field<"RED", 16, 8, ReadWrite, "Red Value These bits define the red value of the output image. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
      Field<"ALPHA", 24, 8, ReadWrite, "Alpha Channel Value These bits define the alpha channel of the output color. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
    ], "DMA2D output color register">,
    Register<"OMAR", 0x3c, 32, [
      Field<"MA", 0, 32, ReadWrite, "Memory Address Address of the data used for the output FIFO. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. The address alignment must match the image format selected e.g. a 32-bit per pixel format must be 32-bit aligned and a 16-bit per pixel format must be 16-bit aligned.">,
    ], "DMA2D output memory address register">,
    Register<"OOR", 0x40, 32, [
      Field<"LO", 0, 14, ReadWrite, "Line Offset Line offset used for the output (expressed in pixels). This value is used for the address generation. It is added at the end of each line to determine the starting address of the next line. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
    ], "DMA2D output offset register">,
    Register<"NLR", 0x44, 32, [
      Field<"NL", 0, 16, ReadWrite, "Number of lines Number of lines of the area to be transferred. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
      Field<"PL", 16, 14, ReadWrite, "Pixel per lines Number of pixels per lines of the area to be transferred. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only. If any of the input image format is 4-bit per pixel, pixel per lines must be even.">,
    ], "DMA2D number of line register">,
    Register<"LWR", 0x48, 32, [
      Field<"LW", 0, 16, ReadWrite, "Line watermark These bits allow to configure the line watermark for interrupt generation. An interrupt is raised when the last pixel of the watermarked line has been transferred. These bits can only be written when data transfers are disabled. Once the transfer has started, they are read-only.">,
    ], "DMA2D line watermark register">,
    Register<"AMTCR", 0x4c, 32, [
      Field<"EN", 0, 1, ReadWrite, "Enable Enables the dead time functionality.">,
      Field<"DT", 8, 8, ReadWrite, "Dead Time Dead time value in the AXI clock cycle inserted between two consecutive accesses on the AXI master port. These bits represent the minimum guaranteed number of cycles between two consecutive AXI accesses.">,
    ], "DMA2D AXI master timer configuration register">,
  ];
}

class DMA2DInstance<string Name, int Base> : PeripheralInstance<Name, Base, DMA2DPeripheral>;
def DMA2D : DMA2DInstance<"DMA2D", 0x52001000>;

#endif // _PERIPHERALS_DMA2D_TD
