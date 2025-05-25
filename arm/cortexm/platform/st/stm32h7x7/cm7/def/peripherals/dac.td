#ifndef _PERIPHERALS_DAC_TD
#define _PERIPHERALS_DAC_TD

include "base.td"

def DACPeripheral : PeripheralType<"DAC", "DAC"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"EN1", 0, 1, ReadWrite, "DAC channel1 enable This bit is set and cleared by software to enable/disable DAC channel1.">,
      Field<"TEN1", 1, 1, ReadWrite, "DAC channel1 trigger enable">,
      Field<"TSEL1", 2, 3, ReadWrite, "DAC channel1 trigger selection These bits select the external event used to trigger DAC channel1. Note: Only used if bit TEN1 = 1 (DAC channel1 trigger enabled).">,
      Field<"WAVE1", 6, 2, ReadWrite, "DAC channel1 noise/triangle wave generation enable These bits are set and cleared by software. Note: Only used if bit TEN1 = 1 (DAC channel1 trigger enabled).">,
      Field<"MAMP1", 8, 4, ReadWrite, "DAC channel1 mask/amplitude selector These bits are written by software to select mask in wave generation mode or amplitude in triangle generation mode. = 1011: Unmask bits[11:0] of LFSR/ triangle amplitude equal to 4095">,
      Field<"DMAEN1", 12, 1, ReadWrite, "DAC channel1 DMA enable This bit is set and cleared by software.">,
      Field<"DMAUDRIE1", 13, 1, ReadWrite, "DAC channel1 DMA Underrun Interrupt enable This bit is set and cleared by software.">,
      Field<"CEN1", 14, 1, ReadWrite, "DAC Channel 1 calibration enable This bit is set and cleared by software to enable/disable DAC channel 1 calibration, it can be written only if bit EN1=0 into DAC_CR (the calibration mode can be entered/exit only when the DAC channel is disabled) Otherwise, the write operation is ignored.">,
      Field<"EN2", 16, 1, ReadWrite, "DAC channel2 enable This bit is set and cleared by software to enable/disable DAC channel2.">,
      Field<"TEN2", 17, 1, ReadWrite, "DAC channel2 trigger enable">,
      Field<"TSEL2", 18, 3, ReadWrite, "DAC channel2 trigger selection These bits select the external event used to trigger DAC channel2 Note: Only used if bit TEN2 = 1 (DAC channel2 trigger enabled).">,
      Field<"WAVE2", 22, 2, ReadWrite, "DAC channel2 noise/triangle wave generation enable These bits are set/reset by software. 1x: Triangle wave generation enabled Note: Only used if bit TEN2 = 1 (DAC channel2 trigger enabled)">,
      Field<"MAMP2", 24, 4, ReadWrite, "DAC channel2 mask/amplitude selector These bits are written by software to select mask in wave generation mode or amplitude in triangle generation mode. = 1011: Unmask bits[11:0] of LFSR/ triangle amplitude equal to 4095">,
      Field<"DMAEN2", 28, 1, ReadWrite, "DAC channel2 DMA enable This bit is set and cleared by software.">,
      Field<"DMAUDRIE2", 29, 1, ReadWrite, "DAC channel2 DMA underrun interrupt enable This bit is set and cleared by software.">,
      Field<"CEN2", 30, 1, ReadWrite, "DAC Channel 2 calibration enable This bit is set and cleared by software to enable/disable DAC channel 2 calibration, it can be written only if bit EN2=0 into DAC_CR (the calibration mode can be entered/exit only when the DAC channel is disabled) Otherwise, the write operation is ignored.">,
    ], "DAC control register">,
    Register<"SWTRGR", 0x4, 32, [
      Field<"SWTRIG1", 0, 1, ReadWrite, "DAC channel1 software trigger This bit is set by software to trigger the DAC in software trigger mode. Note: This bit is cleared by hardware (one APB1 clock cycle later) once the DAC_DHR1 register value has been loaded into the DAC_DOR1 register.">,
      Field<"SWTRIG2", 1, 1, ReadWrite, "DAC channel2 software trigger This bit is set by software to trigger the DAC in software trigger mode. Note: This bit is cleared by hardware (one APB1 clock cycle later) once the DAC_DHR2 register value has been loaded into the DAC_DOR2 register.">,
    ], "DAC software trigger register">,
    Register<"DHR12R1", 0x8, 32, [
      Field<"DACC1DHR", 0, 12, ReadWrite, "DAC channel1 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.">,
    ], "DAC channel1 12-bit right-aligned data holding register">,
    Register<"DHR12L1", 0xc, 32, [
      Field<"DACC1DHR", 4, 12, ReadWrite, "DAC channel1 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.">,
    ], "DAC channel1 12-bit left aligned data holding register">,
    Register<"DHR8R1", 0x10, 32, [
      Field<"DACC1DHR", 0, 8, ReadWrite, "DAC channel1 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel1.">,
    ], "DAC channel1 8-bit right aligned data holding register">,
    Register<"DHR12R2", 0x14, 32, [
      Field<"DACC2DHR", 0, 12, ReadWrite, "DAC channel2 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.">,
    ], "DAC channel2 12-bit right aligned data holding register">,
    Register<"DHR12L2", 0x18, 32, [
      Field<"DACC2DHR", 4, 12, ReadWrite, "DAC channel2 12-bit left-aligned data These bits are written by software which specify 12-bit data for DAC channel2.">,
    ], "DAC channel2 12-bit left aligned data holding register">,
    Register<"DHR8R2", 0x1c, 32, [
      Field<"DACC2DHR", 0, 8, ReadWrite, "DAC channel2 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel2.">,
    ], "DAC channel2 8-bit right-aligned data holding register">,
    Register<"DHR12RD", 0x20, 32, [
      Field<"DACC1DHR", 0, 12, ReadWrite, "DAC channel1 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.">,
      Field<"DACC2DHR", 16, 12, ReadWrite, "DAC channel2 12-bit right-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.">,
    ], "Dual DAC 12-bit right-aligned data holding register">,
    Register<"DHR12LD", 0x24, 32, [
      Field<"DACC1DHR", 4, 12, ReadWrite, "DAC channel1 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel1.">,
      Field<"DACC2DHR", 20, 12, ReadWrite, "DAC channel2 12-bit left-aligned data These bits are written by software which specifies 12-bit data for DAC channel2.">,
    ], "DUAL DAC 12-bit left aligned data holding register">,
    Register<"DHR8RD", 0x28, 32, [
      Field<"DACC1DHR", 0, 8, ReadWrite, "DAC channel1 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel1.">,
      Field<"DACC2DHR", 8, 8, ReadWrite, "DAC channel2 8-bit right-aligned data These bits are written by software which specifies 8-bit data for DAC channel2.">,
    ], "DUAL DAC 8-bit right aligned data holding register">,
    Register<"DOR1", 0x2c, 32, [
      Field<"DACC1DOR", 0, 12, ReadWrite, "DAC channel1 data output These bits are read-only, they contain data output for DAC channel1.">,
    ], "DAC channel1 data output register">,
    Register<"DOR2", 0x30, 32, [
      Field<"DACC2DOR", 0, 12, ReadWrite, "DAC channel2 data output These bits are read-only, they contain data output for DAC channel2.">,
    ], "DAC channel2 data output register">,
    Register<"SR", 0x34, 32, [
      Field<"DMAUDR1", 13, 1, ReadWrite, "DAC channel1 DMA underrun flag This bit is set by hardware and cleared by software (by writing it to 1).">,
      Field<"CAL_FLAG1", 14, 1, Read, "DAC Channel 1 calibration offset status This bit is set and cleared by hardware">,
      Field<"BWST1", 15, 1, Read, "DAC Channel 1 busy writing sample time flag This bit is systematically set just after Sample & Hold mode enable and is set each time the software writes the register DAC_SHSR1, It is cleared by hardware when the write operation of DAC_SHSR1 is complete. (It takes about 3LSI periods of synchronization).">,
      Field<"DMAUDR2", 29, 1, ReadWrite, "DAC channel2 DMA underrun flag This bit is set by hardware and cleared by software (by writing it to 1).">,
      Field<"CAL_FLAG2", 30, 1, Read, "DAC Channel 2 calibration offset status This bit is set and cleared by hardware">,
      Field<"BWST2", 31, 1, Read, "DAC Channel 2 busy writing sample time flag This bit is systematically set just after Sample & Hold mode enable and is set each time the software writes the register DAC_SHSR2, It is cleared by hardware when the write operation of DAC_SHSR2 is complete. (It takes about 3 LSI periods of synchronization).">,
    ], "DAC status register">,
    Register<"CCR", 0x38, 32, [
      Field<"OTRIM1", 0, 5, ReadWrite, "DAC Channel 1 offset trimming value">,
      Field<"OTRIM2", 16, 5, ReadWrite, "DAC Channel 2 offset trimming value">,
    ], "DAC calibration control register">,
    Register<"MCR", 0x3c, 32, [
      Field<"MODE1", 0, 3, ReadWrite, "DAC Channel 1 mode These bits can be written only when the DAC is disabled and not in the calibration mode (when bit EN1=0 and bit CEN1 =0 in the DAC_CR register). If EN1=1 or CEN1 =1 the write operation is ignored. They can be set and cleared by software to select the DAC Channel 1 mode: DAC Channel 1 in normal Mode DAC Channel 1 in sample &amp; hold mode">,
      Field<"MODE2", 16, 3, ReadWrite, "DAC Channel 2 mode These bits can be written only when the DAC is disabled and not in the calibration mode (when bit EN2=0 and bit CEN2 =0 in the DAC_CR register). If EN2=1 or CEN2 =1 the write operation is ignored. They can be set and cleared by software to select the DAC Channel 2 mode: DAC Channel 2 in normal Mode DAC Channel 2 in sample &amp; hold mode">,
    ], "DAC mode control register">,
    Register<"SHSR1", 0x40, 32, [
      Field<"TSAMPLE1", 0, 10, ReadWrite, "DAC Channel 1 sample Time (only valid in sample &amp; hold mode) These bits can be written when the DAC channel1 is disabled or also during normal operation. in the latter case, the write can be done only when BWSTx of DAC_SR register is low, If BWSTx=1, the write operation is ignored.">,
    ], "DAC Sample and Hold sample time register 1">,
    Register<"SHSR2", 0x44, 32, [
      Field<"TSAMPLE2", 0, 10, ReadWrite, "DAC Channel 2 sample Time (only valid in sample &amp; hold mode) These bits can be written when the DAC channel2 is disabled or also during normal operation. in the latter case, the write can be done only when BWSTx of DAC_SR register is low, if BWSTx=1, the write operation is ignored.">,
    ], "DAC Sample and Hold sample time register 2">,
    Register<"SHHR", 0x48, 32, [
      Field<"THOLD1", 0, 10, ReadWrite, "DAC Channel 1 hold Time (only valid in sample &amp; hold mode) Hold time= (THOLD[9:0]) x T LSI">,
      Field<"THOLD2", 16, 10, ReadWrite, "DAC Channel 2 hold time (only valid in sample &amp; hold mode). Hold time= (THOLD[9:0]) x T LSI">,
    ], "DAC Sample and Hold hold time register">,
    Register<"SHRR", 0x4c, 32, [
      Field<"TREFRESH1", 0, 8, ReadWrite, "DAC Channel 1 refresh Time (only valid in sample &amp; hold mode) Refresh time= (TREFRESH[7:0]) x T LSI">,
      Field<"TREFRESH2", 16, 8, ReadWrite, "DAC Channel 2 refresh Time (only valid in sample &amp; hold mode) Refresh time= (TREFRESH[7:0]) x T LSI">,
    ], "DAC Sample and Hold refresh time register">,
  ];
}

class DACInstance<string Name, int Base> : PeripheralInstance<Name, Base, DACPeripheral>;
def DAC : DACInstance<"DAC", 0x40007400>;

#endif // _PERIPHERALS_DAC_TD
