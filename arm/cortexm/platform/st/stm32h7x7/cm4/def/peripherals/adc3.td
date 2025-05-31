#ifndef _PERIPHERALS_ADC3_TD
#define _PERIPHERALS_ADC3_TD

include "base.td"

def ADC3Peripheral : PeripheralType<"ADC3", "Analog to Digital Converter"> {
  let accessWidth = 32;
  let registers = [
    Register<"ISR", 0x0, 32, [
      Field<"ADRDY", 0, 1, ReadWrite, "ADC ready flag">,
      Field<"EOSMP", 1, 1, ReadWrite, "ADC group regular end of sampling flag">,
      Field<"EOC", 2, 1, ReadWrite, "ADC group regular end of unitary conversion flag">,
      Field<"EOS", 3, 1, ReadWrite, "ADC group regular end of sequence conversions flag">,
      Field<"OVR", 4, 1, ReadWrite, "ADC group regular overrun flag">,
      Field<"JEOC", 5, 1, ReadWrite, "ADC group injected end of unitary conversion flag">,
      Field<"JEOS", 6, 1, ReadWrite, "ADC group injected end of sequence conversions flag">,
      Field<"AWD1", 7, 1, ReadWrite, "ADC analog watchdog 1 flag">,
      Field<"AWD2", 8, 1, ReadWrite, "ADC analog watchdog 2 flag">,
      Field<"AWD3", 9, 1, ReadWrite, "ADC analog watchdog 3 flag">,
      Field<"JQOVF", 10, 1, ReadWrite, "ADC group injected contexts queue overflow flag">,
    ], "ADC interrupt and status register">,
    Register<"IER", 0x4, 32, [
      Field<"ADRDYIE", 0, 1, ReadWrite, "ADC ready interrupt">,
      Field<"EOSMPIE", 1, 1, ReadWrite, "ADC group regular end of sampling interrupt">,
      Field<"EOCIE", 2, 1, ReadWrite, "ADC group regular end of unitary conversion interrupt">,
      Field<"EOSIE", 3, 1, ReadWrite, "ADC group regular end of sequence conversions interrupt">,
      Field<"OVRIE", 4, 1, ReadWrite, "ADC group regular overrun interrupt">,
      Field<"JEOCIE", 5, 1, ReadWrite, "ADC group injected end of unitary conversion interrupt">,
      Field<"JEOSIE", 6, 1, ReadWrite, "ADC group injected end of sequence conversions interrupt">,
      Field<"AWD1IE", 7, 1, ReadWrite, "ADC analog watchdog 1 interrupt">,
      Field<"AWD2IE", 8, 1, ReadWrite, "ADC analog watchdog 2 interrupt">,
      Field<"AWD3IE", 9, 1, ReadWrite, "ADC analog watchdog 3 interrupt">,
      Field<"JQOVFIE", 10, 1, ReadWrite, "ADC group injected contexts queue overflow interrupt">,
    ], "ADC interrupt enable register">,
    Register<"CR", 0x8, 32, [
      Field<"ADEN", 0, 1, ReadWrite, "ADC enable">,
      Field<"ADDIS", 1, 1, ReadWrite, "ADC disable">,
      Field<"ADSTART", 2, 1, ReadWrite, "ADC group regular conversion start">,
      Field<"JADSTART", 3, 1, ReadWrite, "ADC group injected conversion start">,
      Field<"ADSTP", 4, 1, ReadWrite, "ADC group regular conversion stop">,
      Field<"JADSTP", 5, 1, ReadWrite, "ADC group injected conversion stop">,
      Field<"BOOST", 8, 2, ReadWrite, "Boost mode control">,
      Field<"ADCALLIN", 16, 1, ReadWrite, "Linearity calibration">,
      Field<"LINCALRDYW1", 22, 1, ReadWrite, "Linearity calibration ready Word 1">,
      Field<"LINCALRDYW2", 23, 1, ReadWrite, "Linearity calibration ready Word 2">,
      Field<"LINCALRDYW3", 24, 1, ReadWrite, "Linearity calibration ready Word 3">,
      Field<"LINCALRDYW4", 25, 1, ReadWrite, "Linearity calibration ready Word 4">,
      Field<"LINCALRDYW5", 26, 1, ReadWrite, "Linearity calibration ready Word 5">,
      Field<"LINCALRDYW6", 27, 1, ReadWrite, "Linearity calibration ready Word 6">,
      Field<"ADVREGEN", 28, 1, ReadWrite, "ADC voltage regulator enable">,
      Field<"DEEPPWD", 29, 1, ReadWrite, "ADC deep power down enable">,
      Field<"ADCALDIF", 30, 1, ReadWrite, "ADC differential mode for calibration">,
      Field<"ADCAL", 31, 1, ReadWrite, "ADC calibration">,
    ], "ADC control register">,
    Register<"CFGR", 0xc, 32, [
      Field<"DMNGT", 0, 2, ReadWrite, "ADC DMA transfer enable">,
      Field<"RES", 2, 3, ReadWrite, "ADC data resolution">,
      Field<"EXTSEL", 5, 5, ReadWrite, "ADC group regular external trigger source">,
      Field<"EXTEN", 10, 2, ReadWrite, "ADC group regular external trigger polarity">,
      Field<"OVRMOD", 12, 1, ReadWrite, "ADC group regular overrun configuration">,
      Field<"CONT", 13, 1, ReadWrite, "ADC group regular continuous conversion mode">,
      Field<"AUTDLY", 14, 1, ReadWrite, "ADC low power auto wait">,
      Field<"DISCEN", 16, 1, ReadWrite, "ADC group regular sequencer discontinuous mode">,
      Field<"DISCNUM", 17, 3, ReadWrite, "ADC group regular sequencer discontinuous number of ranks">,
      Field<"JDISCEN", 20, 1, ReadWrite, "ADC group injected sequencer discontinuous mode">,
      Field<"JQM", 21, 1, ReadWrite, "ADC group injected contexts queue mode">,
      Field<"AWD1SGL", 22, 1, ReadWrite, "ADC analog watchdog 1 monitoring a single channel or all channels">,
      Field<"AWD1EN", 23, 1, ReadWrite, "ADC analog watchdog 1 enable on scope ADC group regular">,
      Field<"JAWD1EN", 24, 1, ReadWrite, "ADC analog watchdog 1 enable on scope ADC group injected">,
      Field<"JAUTO", 25, 1, ReadWrite, "ADC group injected automatic trigger mode">,
      Field<"AWDCH1CH", 26, 5, ReadWrite, "ADC analog watchdog 1 monitored channel selection">,
      Field<"JQDIS", 31, 1, ReadWrite, "ADC group injected contexts queue disable">,
    ], "ADC configuration register 1">,
    Register<"CFGR2", 0x10, 32, [
      Field<"ROVSE", 0, 1, ReadWrite, "ADC oversampler enable on scope ADC group regular">,
      Field<"JOVSE", 1, 1, ReadWrite, "ADC oversampler enable on scope ADC group injected">,
      Field<"OVSS", 5, 4, ReadWrite, "ADC oversampling shift">,
      Field<"TROVS", 9, 1, ReadWrite, "ADC oversampling discontinuous mode (triggered mode) for ADC group regular">,
      Field<"ROVSM", 10, 1, ReadWrite, "Regular Oversampling mode">,
      Field<"RSHIFT1", 11, 1, ReadWrite, "Right-shift data after Offset 1 correction">,
      Field<"RSHIFT2", 12, 1, ReadWrite, "Right-shift data after Offset 2 correction">,
      Field<"RSHIFT3", 13, 1, ReadWrite, "Right-shift data after Offset 3 correction">,
      Field<"RSHIFT4", 14, 1, ReadWrite, "Right-shift data after Offset 4 correction">,
      Field<"OSR", 16, 10, ReadWrite, "Oversampling ratio">,
      Field<"LSHIFT", 28, 4, ReadWrite, "Left shift factor">,
    ], "ADC configuration register 2">,
    Register<"SMPR1", 0x14, 32, [
      Field<"SMP1", 3, 3, ReadWrite, "ADC channel 1 sampling time selection">,
      Field<"SMP2", 6, 3, ReadWrite, "ADC channel 2 sampling time selection">,
      Field<"SMP3", 9, 3, ReadWrite, "ADC channel 3 sampling time selection">,
      Field<"SMP4", 12, 3, ReadWrite, "ADC channel 4 sampling time selection">,
      Field<"SMP5", 15, 3, ReadWrite, "ADC channel 5 sampling time selection">,
      Field<"SMP6", 18, 3, ReadWrite, "ADC channel 6 sampling time selection">,
      Field<"SMP7", 21, 3, ReadWrite, "ADC channel 7 sampling time selection">,
      Field<"SMP8", 24, 3, ReadWrite, "ADC channel 8 sampling time selection">,
      Field<"SMP9", 27, 3, ReadWrite, "ADC channel 9 sampling time selection">,
    ], "ADC sampling time register 1">,
    Register<"SMPR2", 0x18, 32, [
      Field<"SMP10", 0, 3, ReadWrite, "ADC channel 10 sampling time selection">,
      Field<"SMP11", 3, 3, ReadWrite, "ADC channel 11 sampling time selection">,
      Field<"SMP12", 6, 3, ReadWrite, "ADC channel 12 sampling time selection">,
      Field<"SMP13", 9, 3, ReadWrite, "ADC channel 13 sampling time selection">,
      Field<"SMP14", 12, 3, ReadWrite, "ADC channel 14 sampling time selection">,
      Field<"SMP15", 15, 3, ReadWrite, "ADC channel 15 sampling time selection">,
      Field<"SMP16", 18, 3, ReadWrite, "ADC channel 16 sampling time selection">,
      Field<"SMP17", 21, 3, ReadWrite, "ADC channel 17 sampling time selection">,
      Field<"SMP18", 24, 3, ReadWrite, "ADC channel 18 sampling time selection">,
      Field<"SMP19", 27, 3, ReadWrite, "ADC channel 18 sampling time selection">,
    ], "ADC sampling time register 2">,
    Register<"PCSEL", 0x1c, 32, [
      Field<"PCSEL", 0, 20, ReadWrite, "Channel x (VINP[i]) pre selection">,
    ], "ADC pre channel selection register">,
    Register<"LTR1", 0x20, 32, [
      Field<"LTR1", 0, 26, ReadWrite, "ADC analog watchdog 1 threshold low">,
    ], "ADC analog watchdog 1 threshold register">,
    Register<"LHTR1", 0x24, 32, [
      Field<"LHTR1", 0, 26, ReadWrite, "ADC analog watchdog 2 threshold low">,
    ], "ADC analog watchdog 2 threshold register">,
    Register<"SQR1", 0x30, 32, [
      Field<"L3", 0, 4, ReadWrite, "L3">,
      Field<"SQ1", 6, 5, ReadWrite, "ADC group regular sequencer rank 1">,
      Field<"SQ2", 12, 5, ReadWrite, "ADC group regular sequencer rank 2">,
      Field<"SQ3", 18, 5, ReadWrite, "ADC group regular sequencer rank 3">,
      Field<"SQ4", 24, 5, ReadWrite, "ADC group regular sequencer rank 4">,
    ], "ADC group regular sequencer ranks register 1">,
    Register<"SQR2", 0x34, 32, [
      Field<"SQ5", 0, 5, ReadWrite, "ADC group regular sequencer rank 5">,
      Field<"SQ6", 6, 5, ReadWrite, "ADC group regular sequencer rank 6">,
      Field<"SQ7", 12, 5, ReadWrite, "ADC group regular sequencer rank 7">,
      Field<"SQ8", 18, 5, ReadWrite, "ADC group regular sequencer rank 8">,
      Field<"SQ9", 24, 5, ReadWrite, "ADC group regular sequencer rank 9">,
    ], "ADC group regular sequencer ranks register 2">,
    Register<"SQR3", 0x38, 32, [
      Field<"SQ10", 0, 5, ReadWrite, "ADC group regular sequencer rank 10">,
      Field<"SQ11", 6, 5, ReadWrite, "ADC group regular sequencer rank 11">,
      Field<"SQ12", 12, 5, ReadWrite, "ADC group regular sequencer rank 12">,
      Field<"SQ13", 18, 5, ReadWrite, "ADC group regular sequencer rank 13">,
      Field<"SQ14", 24, 5, ReadWrite, "ADC group regular sequencer rank 14">,
    ], "ADC group regular sequencer ranks register 3">,
    Register<"SQR4", 0x3c, 32, [
      Field<"SQ15", 0, 5, ReadWrite, "ADC group regular sequencer rank 15">,
      Field<"SQ16", 6, 5, ReadWrite, "ADC group regular sequencer rank 16">,
    ], "ADC group regular sequencer ranks register 4">,
    Register<"DR", 0x40, 32, [
      Field<"RDATA", 0, 16, ReadWrite, "ADC group regular conversion data">,
    ], "ADC group regular conversion data register">,
    Register<"JSQR", 0x4c, 32, [
      Field<"JL", 0, 2, ReadWrite, "ADC group injected sequencer scan length">,
      Field<"JEXTSEL", 2, 5, ReadWrite, "ADC group injected external trigger source">,
      Field<"JEXTEN", 7, 2, ReadWrite, "ADC group injected external trigger polarity">,
      Field<"JSQ1", 9, 5, ReadWrite, "ADC group injected sequencer rank 1">,
      Field<"JSQ2", 15, 5, ReadWrite, "ADC group injected sequencer rank 2">,
      Field<"JSQ3", 21, 5, ReadWrite, "ADC group injected sequencer rank 3">,
      Field<"JSQ4", 27, 5, ReadWrite, "ADC group injected sequencer rank 4">,
    ], "ADC group injected sequencer register">,
    Register<"OFR1", 0x60, 32, [
      Field<"OFFSET1", 0, 26, ReadWrite, "ADC offset number 1 offset level">,
      Field<"OFFSET1_CH", 26, 5, ReadWrite, "ADC offset number 1 channel selection">,
      Field<"SSATE", 31, 1, ReadWrite, "ADC offset number 1 enable">,
    ], "ADC offset number 1 register">,
    Register<"OFR2", 0x64, 32, [
      Field<"OFFSET1", 0, 26, ReadWrite, "ADC offset number 1 offset level">,
      Field<"OFFSET1_CH", 26, 5, ReadWrite, "ADC offset number 1 channel selection">,
      Field<"SSATE", 31, 1, ReadWrite, "ADC offset number 1 enable">,
    ], "ADC offset number 2 register">,
    Register<"OFR3", 0x68, 32, [
      Field<"OFFSET1", 0, 26, ReadWrite, "ADC offset number 1 offset level">,
      Field<"OFFSET1_CH", 26, 5, ReadWrite, "ADC offset number 1 channel selection">,
      Field<"SSATE", 31, 1, ReadWrite, "ADC offset number 1 enable">,
    ], "ADC offset number 3 register">,
    Register<"OFR4", 0x6c, 32, [
      Field<"OFFSET1", 0, 26, ReadWrite, "ADC offset number 1 offset level">,
      Field<"OFFSET1_CH", 26, 5, ReadWrite, "ADC offset number 1 channel selection">,
      Field<"SSATE", 31, 1, ReadWrite, "ADC offset number 1 enable">,
    ], "ADC offset number 4 register">,
    Register<"JDR1", 0x80, 32, [
      Field<"JDATA1", 0, 32, ReadWrite, "ADC group injected sequencer rank 1 conversion data">,
    ], "ADC group injected sequencer rank 1 register">,
    Register<"JDR2", 0x84, 32, [
      Field<"JDATA2", 0, 32, ReadWrite, "ADC group injected sequencer rank 2 conversion data">,
    ], "ADC group injected sequencer rank 2 register">,
    Register<"JDR3", 0x88, 32, [
      Field<"JDATA3", 0, 32, ReadWrite, "ADC group injected sequencer rank 3 conversion data">,
    ], "ADC group injected sequencer rank 3 register">,
    Register<"JDR4", 0x8c, 32, [
      Field<"JDATA4", 0, 32, ReadWrite, "ADC group injected sequencer rank 4 conversion data">,
    ], "ADC group injected sequencer rank 4 register">,
    Register<"AWD2CR", 0xa0, 32, [
      Field<"AWD2CH", 0, 20, ReadWrite, "ADC analog watchdog 2 monitored channel selection">,
    ], "ADC analog watchdog 2 configuration register">,
    Register<"AWD3CR", 0xa4, 32, [
      Field<"AWD3CH", 1, 20, ReadWrite, "ADC analog watchdog 3 monitored channel selection">,
    ], "ADC analog watchdog 3 configuration register">,
    Register<"LTR2", 0xb0, 32, [
      Field<"LTR2", 0, 26, ReadWrite, "Analog watchdog 2 lower threshold">,
    ], "ADC watchdog lower threshold register 2">,
    Register<"HTR2", 0xb4, 32, [
      Field<"HTR2", 0, 26, ReadWrite, "Analog watchdog 2 higher threshold">,
    ], "ADC watchdog higher threshold register 2">,
    Register<"LTR3", 0xb8, 32, [
      Field<"LTR3", 0, 26, ReadWrite, "Analog watchdog 3 lower threshold">,
    ], "ADC watchdog lower threshold register 3">,
    Register<"HTR3", 0xbc, 32, [
      Field<"HTR3", 0, 26, ReadWrite, "Analog watchdog 3 higher threshold">,
    ], "ADC watchdog higher threshold register 3">,
    Register<"DIFSEL", 0xc0, 32, [
      Field<"DIFSEL", 0, 20, ReadWrite, "ADC channel differential or single-ended mode for channel">,
    ], "ADC channel differential or single-ended mode selection register">,
    Register<"CALFACT", 0xc4, 32, [
      Field<"CALFACT_S", 0, 11, ReadWrite, "ADC calibration factor in single-ended mode">,
      Field<"CALFACT_D", 16, 11, ReadWrite, "ADC calibration factor in differential mode">,
    ], "ADC calibration factors register">,
    Register<"CALFACT2", 0xc8, 32, [
      Field<"LINCALFACT", 0, 30, ReadWrite, "Linearity Calibration Factor">,
    ], "ADC Calibration Factor register 2">,
  ];
}

class ADC3Instance<string Name, int Base> : PeripheralInstance<Name, Base, ADC3Peripheral>;
def ADC1 : ADC3Instance<"ADC1", 0x40022000>;
def ADC2 : ADC3Instance<"ADC2", 0x40022100>;
def ADC3 : ADC3Instance<"ADC3", 0x58026000>;

#endif // _PERIPHERALS_ADC3_TD
