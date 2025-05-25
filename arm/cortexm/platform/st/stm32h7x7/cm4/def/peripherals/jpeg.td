#ifndef _PERIPHERALS_JPEG_TD
#define _PERIPHERALS_JPEG_TD

include "base.td"

def JPEGPeripheral : PeripheralType<"JPEG", "JPEG"> {
  let accessWidth = 32;
  let registers = [
    Register<"CONFR0", 0x0, 32, [
      Field<"START", 0, 1, ReadWrite, "Start This bit start or stop the encoding or decoding process. Read this register always return 0.">,
    ], "JPEG codec control register">,
    Register<"CONFR1", 0x4, 32, [
      Field<"NF", 0, 2, ReadWrite, "Number of color components This field defines the number of color components minus 1.">,
      Field<"DE", 3, 1, ReadWrite, "Decoding Enable This bit selects the coding or decoding process">,
      Field<"COLORSPACE", 4, 2, ReadWrite, "Color Space This filed defines the number of quantization tables minus 1 to insert in the output stream.">,
      Field<"NS", 6, 2, ReadWrite, "Number of components for Scan This field defines the number of components minus 1 for scan header marker segment.">,
      Field<"HDR", 8, 1, ReadWrite, "Header Processing This bit enable the header processing (generation/parsing).">,
      Field<"YSIZE", 16, 16, ReadWrite, "Y Size This field defines the number of lines in source image.">,
    ], "JPEG codec configuration register 1">,
    Register<"CONFR2", 0x8, 32, [
      Field<"NMCU", 0, 26, ReadWrite, "Number of MCU For encoding: this field defines the number of MCU units minus 1 to encode. For decoding: this field indicates the number of complete MCU units minus 1 to be decoded (this field is updated after the JPEG header parsing). If the decoded image size has not a X or Y size multiple of 8 or 16 (depending on the sub-sampling process), the resulting incomplete or empty MCU must be added to this value to get the total number of MCU generated.">,
    ], "JPEG codec configuration register 2">,
    Register<"CONFR3", 0xc, 32, [
      Field<"XSIZE", 16, 16, ReadWrite, "X size This field defines the number of pixels per line.">,
    ], "JPEG codec configuration register 3">,
    Register<"CONFRN1", 0x10, 32, [
      Field<"HD", 0, 1, ReadWrite, "Huffman DC Selects the Huffman table for encoding the DC coefficients.">,
      Field<"HA", 1, 1, ReadWrite, "Huffman AC Selects the Huffman table for encoding the AC coefficients.">,
      Field<"QT", 2, 2, ReadWrite, "Quantization Table Selects quantization table associated with a color component.">,
      Field<"NB", 4, 4, ReadWrite, "Number of Block Number of data units minus 1 that belong to a particular color in the MCU.">,
      Field<"VSF", 8, 4, ReadWrite, "Vertical Sampling Factor Vertical sampling factor for component i.">,
      Field<"HSF", 12, 4, ReadWrite, "Horizontal Sampling Factor Horizontal sampling factor for component i.">,
    ], "JPEG codec configuration register 4-7">,
    Register<"CONFRN2", 0x14, 32, [
      Field<"HD", 0, 1, ReadWrite, "Huffman DC Selects the Huffman table for encoding the DC coefficients.">,
      Field<"HA", 1, 1, ReadWrite, "Huffman AC Selects the Huffman table for encoding the AC coefficients.">,
      Field<"QT", 2, 2, ReadWrite, "Quantization Table Selects quantization table associated with a color component.">,
      Field<"NB", 4, 4, ReadWrite, "Number of Block Number of data units minus 1 that belong to a particular color in the MCU.">,
      Field<"VSF", 8, 4, ReadWrite, "Vertical Sampling Factor Vertical sampling factor for component i.">,
      Field<"HSF", 12, 4, ReadWrite, "Horizontal Sampling Factor Horizontal sampling factor for component i.">,
    ], "JPEG codec configuration register 4-7">,
    Register<"CONFRN3", 0x18, 32, [
      Field<"HD", 0, 1, ReadWrite, "Huffman DC Selects the Huffman table for encoding the DC coefficients.">,
      Field<"HA", 1, 1, ReadWrite, "Huffman AC Selects the Huffman table for encoding the AC coefficients.">,
      Field<"QT", 2, 2, ReadWrite, "Quantization Table Selects quantization table associated with a color component.">,
      Field<"NB", 4, 4, ReadWrite, "Number of Block Number of data units minus 1 that belong to a particular color in the MCU.">,
      Field<"VSF", 8, 4, ReadWrite, "Vertical Sampling Factor Vertical sampling factor for component i.">,
      Field<"HSF", 12, 4, ReadWrite, "Horizontal Sampling Factor Horizontal sampling factor for component i.">,
    ], "JPEG codec configuration register 4-7">,
    Register<"CONFRN4", 0x1c, 32, [
      Field<"HD", 0, 1, ReadWrite, "Huffman DC Selects the Huffman table for encoding the DC coefficients.">,
      Field<"HA", 1, 1, ReadWrite, "Huffman AC Selects the Huffman table for encoding the AC coefficients.">,
      Field<"QT", 2, 2, ReadWrite, "Quantization Table Selects quantization table associated with a color component.">,
      Field<"NB", 4, 4, ReadWrite, "Number of Block Number of data units minus 1 that belong to a particular color in the MCU.">,
      Field<"VSF", 8, 4, ReadWrite, "Vertical Sampling Factor Vertical sampling factor for component i.">,
      Field<"HSF", 12, 4, ReadWrite, "Horizontal Sampling Factor Horizontal sampling factor for component i.">,
    ], "JPEG codec configuration register 4-7">,
    Register<"CR", 0x30, 32, [
      Field<"JCEN", 0, 1, ReadWrite, "JPEG Core Enable Enable the JPEG codec Core.">,
      Field<"IFTIE", 1, 1, ReadWrite, "Input FIFO Threshold Interrupt Enable This bit enables the interrupt generation when input FIFO reach the threshold.">,
      Field<"IFNFIE", 2, 1, ReadWrite, "Input FIFO Not Full Interrupt Enable This bit enables the interrupt generation when input FIFO is not empty.">,
      Field<"OFTIE", 3, 1, ReadWrite, "Output FIFO Threshold Interrupt Enable This bit enables the interrupt generation when output FIFO reach the threshold.">,
      Field<"OFNEIE", 4, 1, ReadWrite, "Output FIFO Not Empty Interrupt Enable This bit enables the interrupt generation when output FIFO is not empty.">,
      Field<"EOCIE", 5, 1, ReadWrite, "End of Conversion Interrupt Enable This bit enables the interrupt generation on the end of conversion.">,
      Field<"HPDIE", 6, 1, ReadWrite, "Header Parsing Done Interrupt Enable This bit enables the interrupt generation on the Header Parsing Operation.">,
      Field<"IDMAEN", 11, 1, ReadWrite, "Input DMA Enable Enable the DMA request generation for the input FIFO.">,
      Field<"ODMAEN", 12, 1, ReadWrite, "Output DMA Enable Enable the DMA request generation for the output FIFO.">,
      Field<"IFF", 13, 1, ReadWrite, "Input FIFO Flush This bit flush the input FIFO. This bit is always read as 0.">,
      Field<"OFF", 14, 1, ReadWrite, "Output FIFO Flush This bit flush the output FIFO. This bit is always read as 0.">,
    ], "JPEG control register">,
    Register<"SR", 0x34, 32, [
      Field<"IFTF", 1, 1, ReadWrite, "Input FIFO Threshold Flag This bit is set when the input FIFO is not full and is bellow its threshold.">,
      Field<"IFNFF", 2, 1, ReadWrite, "Input FIFO Not Full Flag This bit is set when the input FIFO is not full (a data can be written).">,
      Field<"OFTF", 3, 1, ReadWrite, "Output FIFO Threshold Flag This bit is set when the output FIFO is not empty and has reach its threshold.">,
      Field<"OFNEF", 4, 1, ReadWrite, "Output FIFO Not Empty Flag This bit is set when the output FIFO is not empty (a data is available).">,
      Field<"EOCF", 5, 1, ReadWrite, "End of Conversion Flag This bit is set when the JPEG codec core has finished the encoding or the decoding process and than last data has been sent to the output FIFO.">,
      Field<"HPDF", 6, 1, ReadWrite, "Header Parsing Done Flag This bit is set in decode mode when the JPEG codec has finished the parsing of the headers and the internal registers have been updated.">,
      Field<"COF", 7, 1, ReadWrite, "Codec Operation Flag This bit is set when when a JPEG codec operation is on going (encoding or decoding).">,
    ], "JPEG status register">,
    Register<"CFR", 0x38, 32, [
      Field<"CEOCF", 5, 1, ReadWrite, "Clear End of Conversion Flag Writing 1 clears the End of Conversion Flag of the JPEG Status Register.">,
      Field<"CHPDF", 6, 1, ReadWrite, "Clear Header Parsing Done Flag Writing 1 clears the Header Parsing Done Flag of the JPEG Status Register.">,
    ], "JPEG clear flag register">,
    Register<"DIR", 0x40, 32, [
      Field<"DATAIN", 0, 32, ReadWrite, "Data Input FIFO Input FIFO data register.">,
    ], "JPEG data input register">,
    Register<"DOR", 0x44, 32, [
      Field<"DATAOUT", 0, 32, ReadWrite, "Data Output FIFO Output FIFO data register.">,
    ], "JPEG data output register">,
  ];
}

class JPEGInstance<string Name, int Base> : PeripheralInstance<Name, Base, JPEGPeripheral>;
def JPEG : JPEGInstance<"JPEG", 0x52003000>;

#endif // _PERIPHERALS_JPEG_TD
