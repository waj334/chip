#ifndef _PERIPHERALS_ADC3_COMMON_TD
#define _PERIPHERALS_ADC3_COMMON_TD

include "base.td"

def ADC3_CommonPeripheral : PeripheralType<"ADC3_Common", "Analog-to-Digital Converter"> {
  let accessWidth = 32;
  let registers = [
    Register<"CSR", 0x0, 32, [
      Field<"ADRDY_MST", 0, 1, ReadWrite, "Master ADC ready">,
      Field<"EOSMP_MST", 1, 1, ReadWrite, "End of Sampling phase flag of the master ADC">,
      Field<"EOC_MST", 2, 1, ReadWrite, "End of regular conversion of the master ADC">,
      Field<"EOS_MST", 3, 1, ReadWrite, "End of regular sequence flag of the master ADC">,
      Field<"OVR_MST", 4, 1, ReadWrite, "Overrun flag of the master ADC">,
      Field<"JEOC_MST", 5, 1, ReadWrite, "End of injected conversion flag of the master ADC">,
      Field<"JEOS_MST", 6, 1, ReadWrite, "End of injected sequence flag of the master ADC">,
      Field<"AWD1_MST", 7, 1, ReadWrite, "Analog watchdog 1 flag of the master ADC">,
      Field<"AWD2_MST", 8, 1, ReadWrite, "Analog watchdog 2 flag of the master ADC">,
      Field<"AWD3_MST", 9, 1, ReadWrite, "Analog watchdog 3 flag of the master ADC">,
      Field<"JQOVF_MST", 10, 1, ReadWrite, "Injected Context Queue Overflow flag of the master ADC">,
      Field<"ADRDY_SLV", 16, 1, ReadWrite, "Slave ADC ready">,
      Field<"EOSMP_SLV", 17, 1, ReadWrite, "End of Sampling phase flag of the slave ADC">,
      Field<"EOC_SLV", 18, 1, ReadWrite, "End of regular conversion of the slave ADC">,
      Field<"EOS_SLV", 19, 1, ReadWrite, "End of regular sequence flag of the slave ADC">,
      Field<"OVR_SLV", 20, 1, ReadWrite, "Overrun flag of the slave ADC">,
      Field<"JEOC_SLV", 21, 1, ReadWrite, "End of injected conversion flag of the slave ADC">,
      Field<"JEOS_SLV", 22, 1, ReadWrite, "End of injected sequence flag of the slave ADC">,
      Field<"AWD1_SLV", 23, 1, ReadWrite, "Analog watchdog 1 flag of the slave ADC">,
      Field<"AWD2_SLV", 24, 1, ReadWrite, "Analog watchdog 2 flag of the slave ADC">,
      Field<"AWD3_SLV", 25, 1, ReadWrite, "Analog watchdog 3 flag of the slave ADC">,
      Field<"JQOVF_SLV", 26, 1, ReadWrite, "Injected Context Queue Overflow flag of the slave ADC">,
    ], "ADC Common status register">,
    Register<"CCR", 0x8, 32, [
      Field<"DUAL", 0, 5, ReadWrite, "Dual ADC mode selection">,
      Field<"DELAY", 8, 4, ReadWrite, "Delay between 2 sampling phases">,
      Field<"DAMDF", 14, 2, ReadWrite, "Dual ADC Mode Data Format">,
      Field<"CKMODE", 16, 2, ReadWrite, "ADC clock mode">,
      Field<"PRESC", 18, 4, ReadWrite, "ADC prescaler">,
      Field<"VREFEN", 22, 1, ReadWrite, "VREFINT enable">,
      Field<"TSEN", 23, 1, ReadWrite, "Temperature sensor enable">,
      Field<"VBATEN", 24, 1, ReadWrite, "VBAT enable">,
    ], "ADC common control register">,
    Register<"CDR", 0xc, 32, [
      Field<"RDATA_SLV", 16, 16, ReadWrite, "Regular data of the slave ADC">,
      Field<"RDATA_MST", 0, 16, ReadWrite, "Regular data of the master ADC">,
    ], "ADC common regular data register for dual and triple modes">,
    Register<"CDR2", 0x10, 32, [
      Field<"RDATA_ALT", 0, 32, ReadWrite, "Regular data of the master/slave alternated ADCs">,
    ], "ADC x common regular data register for 32-bit dual mode">,
  ];
}

class ADC3_CommonInstance<string Name, int Base> : PeripheralInstance<Name, Base, ADC3_CommonPeripheral>;
def ADC3_Common : ADC3_CommonInstance<"ADC3_Common", 0x58026300>;
def ADC12_Common : ADC3_CommonInstance<"ADC12_Common", 0x40022300>;

#endif // _PERIPHERALS_ADC3_COMMON_TD
