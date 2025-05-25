#ifndef _PERIPHERALS_DFSDM_TD
#define _PERIPHERALS_DFSDM_TD

include "base.td"

def DFSDMPeripheral : PeripheralType<"DFSDM", "Digital filter for sigma delta modulators"> {
  let accessWidth = 32;
  let registers = [
    Register<"CH0CFGR1", 0x0, 32, [
      Field<"SITP", 0, 2, ReadWrite, "SITP">,
      Field<"SPICKSEL", 2, 2, ReadWrite, "SPICKSEL">,
      Field<"SCDEN", 5, 1, ReadWrite, "SCDEN">,
      Field<"CKABEN", 6, 1, ReadWrite, "CKABEN">,
      Field<"CHEN", 7, 1, ReadWrite, "CHEN">,
      Field<"CHINSEL", 8, 1, ReadWrite, "CHINSEL">,
      Field<"DATMPX", 12, 2, ReadWrite, "DATMPX">,
      Field<"DATPACK", 14, 2, ReadWrite, "DATPACK">,
      Field<"CKOUTDIV", 16, 8, ReadWrite, "CKOUTDIV">,
      Field<"CKOUTSRC", 30, 1, ReadWrite, "CKOUTSRC">,
      Field<"DFSDMEN", 31, 1, ReadWrite, "DFSDMEN">,
    ], "channel configuration y register">,
    Register<"CH0CFGR2", 0x4, 32, [
      Field<"DTRBS", 3, 5, ReadWrite, "DTRBS">,
      Field<"OFFSET", 8, 24, ReadWrite, "OFFSET">,
    ], "channel configuration y register">,
    Register<"CH0AWSCDR", 0x8, 32, [
      Field<"SCDT", 0, 8, ReadWrite, "SCDT">,
      Field<"BKSCD", 12, 4, ReadWrite, "BKSCD">,
      Field<"AWFOSR", 16, 5, ReadWrite, "AWFOSR">,
      Field<"AWFORD", 22, 2, ReadWrite, "AWFORD">,
    ], "analog watchdog and short-circuit detector register">,
    Register<"CH0WDATR", 0xc, 32, [
      Field<"WDATA", 0, 16, ReadWrite, "WDATA">,
    ], "channel watchdog filter data register">,
    Register<"CH0DATINR", 0x10, 32, [
      Field<"INDAT0", 0, 16, ReadWrite, "INDAT0">,
      Field<"INDAT1", 16, 16, ReadWrite, "INDAT1">,
    ], "channel data input register">,
    Register<"CH0DLYR", 0x14, 32, [
      Field<"PLSSKP", 0, 6, ReadWrite, "PLSSKP">,
    ], "channel y delay register">,
    Register<"CH1CFGR1", 0x20, 32, [
      Field<"SITP", 0, 2, ReadWrite, "SITP">,
      Field<"SPICKSEL", 2, 2, ReadWrite, "SPICKSEL">,
      Field<"SCDEN", 5, 1, ReadWrite, "SCDEN">,
      Field<"CKABEN", 6, 1, ReadWrite, "CKABEN">,
      Field<"CHEN", 7, 1, ReadWrite, "CHEN">,
      Field<"CHINSEL", 8, 1, ReadWrite, "CHINSEL">,
      Field<"DATMPX", 12, 2, ReadWrite, "DATMPX">,
      Field<"DATPACK", 14, 2, ReadWrite, "DATPACK">,
    ], "CH1CFGR1">,
    Register<"CH1CFGR2", 0x24, 32, [
      Field<"DTRBS", 3, 5, ReadWrite, "DTRBS">,
      Field<"OFFSET", 8, 24, ReadWrite, "OFFSET">,
    ], "CH1CFGR2">,
    Register<"CH1AWSCDR", 0x28, 32, [
      Field<"SCDT", 0, 8, ReadWrite, "SCDT">,
      Field<"BKSCD", 12, 4, ReadWrite, "BKSCD">,
      Field<"AWFOSR", 16, 5, ReadWrite, "AWFOSR">,
      Field<"AWFORD", 22, 2, ReadWrite, "AWFORD">,
    ], "CH1AWSCDR">,
    Register<"CH1WDATR", 0x2c, 32, [
      Field<"WDATA", 0, 16, ReadWrite, "WDATA">,
    ], "CH1WDATR">,
    Register<"CH1DATINR", 0x30, 32, [
      Field<"INDAT0", 0, 16, ReadWrite, "INDAT0">,
      Field<"INDAT1", 16, 16, ReadWrite, "INDAT1">,
    ], "CH1DATINR">,
    Register<"CH1DLYR", 0x34, 32, [
      Field<"PLSSKP", 0, 6, ReadWrite, "PLSSKP">,
    ], "channel y delay register">,
    Register<"CH2CFGR1", 0x40, 32, [
      Field<"SITP", 0, 2, ReadWrite, "SITP">,
      Field<"SPICKSEL", 2, 2, ReadWrite, "SPICKSEL">,
      Field<"SCDEN", 5, 1, ReadWrite, "SCDEN">,
      Field<"CKABEN", 6, 1, ReadWrite, "CKABEN">,
      Field<"CHEN", 7, 1, ReadWrite, "CHEN">,
      Field<"CHINSEL", 8, 1, ReadWrite, "CHINSEL">,
      Field<"DATMPX", 12, 2, ReadWrite, "DATMPX">,
      Field<"DATPACK", 14, 2, ReadWrite, "DATPACK">,
    ], "CH2CFGR1">,
    Register<"CH2CFGR2", 0x44, 32, [
      Field<"DTRBS", 3, 5, ReadWrite, "DTRBS">,
      Field<"OFFSET", 8, 24, ReadWrite, "OFFSET">,
    ], "CH2CFGR2">,
    Register<"CH2AWSCDR", 0x48, 32, [
      Field<"SCDT", 0, 8, ReadWrite, "SCDT">,
      Field<"BKSCD", 12, 4, ReadWrite, "BKSCD">,
      Field<"AWFOSR", 16, 5, ReadWrite, "AWFOSR">,
      Field<"AWFORD", 22, 2, ReadWrite, "AWFORD">,
    ], "CH2AWSCDR">,
    Register<"CH2WDATR", 0x4c, 32, [
      Field<"WDATA", 0, 16, ReadWrite, "WDATA">,
    ], "CH2WDATR">,
    Register<"CH2DATINR", 0x50, 32, [
      Field<"INDAT0", 0, 16, ReadWrite, "INDAT0">,
      Field<"INDAT1", 16, 16, ReadWrite, "INDAT1">,
    ], "CH2DATINR">,
    Register<"CH2DLYR", 0x54, 32, [
      Field<"PLSSKP", 0, 6, ReadWrite, "PLSSKP">,
    ], "channel y delay register">,
    Register<"CH3CFGR1", 0x60, 32, [
      Field<"SITP", 0, 2, ReadWrite, "SITP">,
      Field<"SPICKSEL", 2, 2, ReadWrite, "SPICKSEL">,
      Field<"SCDEN", 5, 1, ReadWrite, "SCDEN">,
      Field<"CKABEN", 6, 1, ReadWrite, "CKABEN">,
      Field<"CHEN", 7, 1, ReadWrite, "CHEN">,
      Field<"CHINSEL", 8, 1, ReadWrite, "CHINSEL">,
      Field<"DATMPX", 12, 2, ReadWrite, "DATMPX">,
      Field<"DATPACK", 14, 2, ReadWrite, "DATPACK">,
    ], "CH3CFGR1">,
    Register<"CH3CFGR2", 0x64, 32, [
      Field<"DTRBS", 3, 5, ReadWrite, "DTRBS">,
      Field<"OFFSET", 8, 24, ReadWrite, "OFFSET">,
    ], "CH3CFGR2">,
    Register<"CH3AWSCDR", 0x68, 32, [
      Field<"SCDT", 0, 8, ReadWrite, "SCDT">,
      Field<"BKSCD", 12, 4, ReadWrite, "BKSCD">,
      Field<"AWFOSR", 16, 5, ReadWrite, "AWFOSR">,
      Field<"AWFORD", 22, 2, ReadWrite, "AWFORD">,
    ], "CH3AWSCDR">,
    Register<"CH3WDATR", 0x6c, 32, [
      Field<"WDATA", 0, 16, ReadWrite, "WDATA">,
    ], "CH3WDATR">,
    Register<"CH3DATINR", 0x70, 32, [
      Field<"INDAT0", 0, 16, ReadWrite, "INDAT0">,
      Field<"INDAT1", 16, 16, ReadWrite, "INDAT1">,
    ], "CH3DATINR">,
    Register<"CH3DLYR", 0x74, 32, [
      Field<"PLSSKP", 0, 6, ReadWrite, "PLSSKP">,
    ], "channel y delay register">,
    Register<"CH4CFGR1", 0x80, 32, [
      Field<"SITP", 0, 2, ReadWrite, "SITP">,
      Field<"SPICKSEL", 2, 2, ReadWrite, "SPICKSEL">,
      Field<"SCDEN", 5, 1, ReadWrite, "SCDEN">,
      Field<"CKABEN", 6, 1, ReadWrite, "CKABEN">,
      Field<"CHEN", 7, 1, ReadWrite, "CHEN">,
      Field<"CHINSEL", 8, 1, ReadWrite, "CHINSEL">,
      Field<"DATMPX", 12, 2, ReadWrite, "DATMPX">,
      Field<"DATPACK", 14, 2, ReadWrite, "DATPACK">,
    ], "CH4CFGR1">,
    Register<"CH4CFGR2", 0x84, 32, [
      Field<"DTRBS", 3, 5, ReadWrite, "DTRBS">,
      Field<"OFFSET", 8, 24, ReadWrite, "OFFSET">,
    ], "CH4CFGR2">,
    Register<"CH4AWSCDR", 0x88, 32, [
      Field<"SCDT", 0, 8, ReadWrite, "SCDT">,
      Field<"BKSCD", 12, 4, ReadWrite, "BKSCD">,
      Field<"AWFOSR", 16, 5, ReadWrite, "AWFOSR">,
      Field<"AWFORD", 22, 2, ReadWrite, "AWFORD">,
    ], "CH4AWSCDR">,
    Register<"CH4WDATR", 0x8c, 32, [
      Field<"WDATA", 0, 16, ReadWrite, "WDATA">,
    ], "CH4WDATR">,
    Register<"CH4DATINR", 0x90, 32, [
      Field<"INDAT0", 0, 16, ReadWrite, "INDAT0">,
      Field<"INDAT1", 16, 16, ReadWrite, "INDAT1">,
    ], "CH4DATINR">,
    Register<"CH4DLYR", 0x94, 32, [
      Field<"PLSSKP", 0, 6, ReadWrite, "PLSSKP">,
    ], "channel y delay register">,
    Register<"CH5CFGR1", 0xa0, 32, [
      Field<"SITP", 0, 2, ReadWrite, "SITP">,
      Field<"SPICKSEL", 2, 2, ReadWrite, "SPICKSEL">,
      Field<"SCDEN", 5, 1, ReadWrite, "SCDEN">,
      Field<"CKABEN", 6, 1, ReadWrite, "CKABEN">,
      Field<"CHEN", 7, 1, ReadWrite, "CHEN">,
      Field<"CHINSEL", 8, 1, ReadWrite, "CHINSEL">,
      Field<"DATMPX", 12, 2, ReadWrite, "DATMPX">,
      Field<"DATPACK", 14, 2, ReadWrite, "DATPACK">,
    ], "CH5CFGR1">,
    Register<"CH5CFGR2", 0xa4, 32, [
      Field<"DTRBS", 3, 5, ReadWrite, "DTRBS">,
      Field<"OFFSET", 8, 24, ReadWrite, "OFFSET">,
    ], "CH5CFGR2">,
    Register<"CH5AWSCDR", 0xa8, 32, [
      Field<"SCDT", 0, 8, ReadWrite, "SCDT">,
      Field<"BKSCD", 12, 4, ReadWrite, "BKSCD">,
      Field<"AWFOSR", 16, 5, ReadWrite, "AWFOSR">,
      Field<"AWFORD", 22, 2, ReadWrite, "AWFORD">,
    ], "CH5AWSCDR">,
    Register<"CH5WDATR", 0xac, 32, [
      Field<"WDATA", 0, 16, ReadWrite, "WDATA">,
    ], "CH5WDATR">,
    Register<"CH5DATINR", 0xb0, 32, [
      Field<"INDAT0", 0, 16, ReadWrite, "INDAT0">,
      Field<"INDAT1", 16, 16, ReadWrite, "INDAT1">,
    ], "CH5DATINR">,
    Register<"CH5DLYR", 0xb4, 32, [
      Field<"PLSSKP", 0, 6, ReadWrite, "PLSSKP">,
    ], "channel y delay register">,
    Register<"CH6CFGR1", 0xc0, 32, [
      Field<"SITP", 0, 2, ReadWrite, "SITP">,
      Field<"SPICKSEL", 2, 2, ReadWrite, "SPICKSEL">,
      Field<"SCDEN", 5, 1, ReadWrite, "SCDEN">,
      Field<"CKABEN", 6, 1, ReadWrite, "CKABEN">,
      Field<"CHEN", 7, 1, ReadWrite, "CHEN">,
      Field<"CHINSEL", 8, 1, ReadWrite, "CHINSEL">,
      Field<"DATMPX", 12, 2, ReadWrite, "DATMPX">,
      Field<"DATPACK", 14, 2, ReadWrite, "DATPACK">,
    ], "CH6CFGR1">,
    Register<"CH6CFGR2", 0xc4, 32, [
      Field<"DTRBS", 3, 5, ReadWrite, "DTRBS">,
      Field<"OFFSET", 8, 24, ReadWrite, "OFFSET">,
    ], "CH6CFGR2">,
    Register<"CH6AWSCDR", 0xc8, 32, [
      Field<"SCDT", 0, 8, ReadWrite, "SCDT">,
      Field<"BKSCD", 12, 4, ReadWrite, "BKSCD">,
      Field<"AWFOSR", 16, 5, ReadWrite, "AWFOSR">,
      Field<"AWFORD", 22, 2, ReadWrite, "AWFORD">,
    ], "CH6AWSCDR">,
    Register<"CH6WDATR", 0xcc, 32, [
      Field<"WDATA", 0, 16, ReadWrite, "WDATA">,
    ], "CH6WDATR">,
    Register<"CH6DATINR", 0xd0, 32, [
      Field<"INDAT0", 0, 16, ReadWrite, "INDAT0">,
      Field<"INDAT1", 16, 16, ReadWrite, "INDAT1">,
    ], "CH6DATINR">,
    Register<"CH6DLYR", 0xd4, 32, [
      Field<"PLSSKP", 0, 6, ReadWrite, "PLSSKP">,
    ], "channel y delay register">,
    Register<"CH7CFGR1", 0xe0, 32, [
      Field<"SITP", 0, 2, ReadWrite, "SITP">,
      Field<"SPICKSEL", 2, 2, ReadWrite, "SPICKSEL">,
      Field<"SCDEN", 5, 1, ReadWrite, "SCDEN">,
      Field<"CKABEN", 6, 1, ReadWrite, "CKABEN">,
      Field<"CHEN", 7, 1, ReadWrite, "CHEN">,
      Field<"CHINSEL", 8, 1, ReadWrite, "CHINSEL">,
      Field<"DATMPX", 12, 2, ReadWrite, "DATMPX">,
      Field<"DATPACK", 14, 2, ReadWrite, "DATPACK">,
    ], "CH7CFGR1">,
    Register<"CH7CFGR2", 0xe4, 32, [
      Field<"DTRBS", 3, 5, ReadWrite, "DTRBS">,
      Field<"OFFSET", 8, 24, ReadWrite, "OFFSET">,
    ], "CH7CFGR2">,
    Register<"CH7AWSCDR", 0xe8, 32, [
      Field<"SCDT", 0, 8, ReadWrite, "SCDT">,
      Field<"BKSCD", 12, 4, ReadWrite, "BKSCD">,
      Field<"AWFOSR", 16, 5, ReadWrite, "AWFOSR">,
      Field<"AWFORD", 22, 2, ReadWrite, "AWFORD">,
    ], "CH7AWSCDR">,
    Register<"CH7WDATR", 0xec, 32, [
      Field<"WDATA", 0, 16, ReadWrite, "WDATA">,
    ], "CH7WDATR">,
    Register<"CH7DATINR", 0xf0, 32, [
      Field<"INDAT0", 0, 16, ReadWrite, "INDAT0">,
      Field<"INDAT1", 16, 16, ReadWrite, "INDAT1">,
    ], "CH7DATINR">,
    Register<"CH7DLYR", 0xf4, 32, [
      Field<"PLSSKP", 0, 6, ReadWrite, "PLSSKP">,
    ], "channel y delay register">,
    Register<"DFSDM_FLT0CR1", 0x100, 32, [
      Field<"DFEN", 0, 1, ReadWrite, "DFSDM enable">,
      Field<"JSWSTART", 1, 1, ReadWrite, "Start a conversion of the injected group of channels">,
      Field<"JSYNC", 3, 1, ReadWrite, "Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger">,
      Field<"JSCAN", 4, 1, ReadWrite, "Scanning conversion mode for injected conversions">,
      Field<"JDMAEN", 5, 1, ReadWrite, "DMA channel enabled to read data for the injected channel group">,
      Field<"JEXTSEL", 8, 3, ReadWrite, "Trigger signal selection for launching injected conversions">,
      Field<"JEXTEN", 13, 2, ReadWrite, "Trigger enable and trigger edge selection for injected conversions">,
      Field<"RSWSTART", 17, 1, ReadWrite, "Software start of a conversion on the regular channel">,
      Field<"RCONT", 18, 1, ReadWrite, "Continuous mode selection for regular conversions">,
      Field<"RSYNC", 19, 1, ReadWrite, "Launch regular conversion synchronously with DFSDM0">,
      Field<"RDMAEN", 21, 1, ReadWrite, "DMA channel enabled to read data for the regular conversion">,
      Field<"RCH", 24, 3, ReadWrite, "Regular channel selection">,
      Field<"FAST", 29, 1, ReadWrite, "Fast conversion mode selection for regular conversions">,
      Field<"AWFSEL", 30, 1, ReadWrite, "Analog watchdog fast mode select">,
    ], "control register 1">,
    Register<"DFSDM_FLT0CR2", 0x104, 32, [
      Field<"JEOCIE", 0, 1, ReadWrite, "Injected end of conversion interrupt enable">,
      Field<"REOCIE", 1, 1, ReadWrite, "Regular end of conversion interrupt enable">,
      Field<"JOVRIE", 2, 1, ReadWrite, "Injected data overrun interrupt enable">,
      Field<"ROVRIE", 3, 1, ReadWrite, "Regular data overrun interrupt enable">,
      Field<"AWDIE", 4, 1, ReadWrite, "Analog watchdog interrupt enable">,
      Field<"SCDIE", 5, 1, ReadWrite, "Short-circuit detector interrupt enable">,
      Field<"CKABIE", 6, 1, ReadWrite, "Clock absence interrupt enable">,
      Field<"EXCH", 8, 8, ReadWrite, "Extremes detector channel selection">,
      Field<"AWDCH", 16, 8, ReadWrite, "Analog watchdog channel selection">,
    ], "control register 2">,
    Register<"DFSDM_FLT0ISR", 0x108, 32, [
      Field<"JEOCF", 0, 1, ReadWrite, "End of injected conversion flag">,
      Field<"REOCF", 1, 1, ReadWrite, "End of regular conversion flag">,
      Field<"JOVRF", 2, 1, ReadWrite, "Injected conversion overrun flag">,
      Field<"ROVRF", 3, 1, ReadWrite, "Regular conversion overrun flag">,
      Field<"AWDF", 4, 1, ReadWrite, "Analog watchdog">,
      Field<"JCIP", 13, 1, ReadWrite, "Injected conversion in progress status">,
      Field<"RCIP", 14, 1, ReadWrite, "Regular conversion in progress status">,
      Field<"CKABF", 16, 8, ReadWrite, "Clock absence flag">,
      Field<"SCDF", 24, 8, ReadWrite, "short-circuit detector flag">,
    ], "interrupt and status register">,
    Register<"DFSDM_FLT0ICR", 0x10c, 32, [
      Field<"CLRJOVRF", 2, 1, ReadWrite, "Clear the injected conversion overrun flag">,
      Field<"CLRROVRF", 3, 1, ReadWrite, "Clear the regular conversion overrun flag">,
      Field<"CLRCKABF", 16, 8, ReadWrite, "Clear the clock absence flag">,
      Field<"CLRSCDF", 24, 8, ReadWrite, "Clear the short-circuit detector flag">,
    ], "interrupt flag clear register">,
    Register<"DFSDM_FLT0JCHGR", 0x110, 32, [
      Field<"JCHG", 0, 8, ReadWrite, "Injected channel group selection">,
    ], "injected channel group selection register">,
    Register<"DFSDM_FLT0FCR", 0x114, 32, [
      Field<"IOSR", 0, 8, ReadWrite, "Integrator oversampling ratio (averaging length)">,
      Field<"FOSR", 16, 10, ReadWrite, "Sinc filter oversampling ratio (decimation rate)">,
      Field<"FORD", 29, 3, ReadWrite, "Sinc filter order">,
    ], "filter control register">,
    Register<"DFSDM_FLT0JDATAR", 0x118, 32, [
      Field<"JDATACH", 0, 3, ReadWrite, "Injected channel most recently converted">,
      Field<"JDATA", 8, 24, ReadWrite, "Injected group conversion data">,
    ], "data register for injected group">,
    Register<"DFSDM_FLT0RDATAR", 0x11c, 32, [
      Field<"RDATACH", 0, 3, ReadWrite, "Regular channel most recently converted">,
      Field<"RPEND", 4, 1, ReadWrite, "Regular channel pending data">,
      Field<"RDATA", 8, 24, ReadWrite, "Regular channel conversion data">,
    ], "data register for the regular channel">,
    Register<"DFSDM_FLT0AWHTR", 0x120, 32, [
      Field<"BKAWH", 0, 4, ReadWrite, "Break signal assignment to analog watchdog high threshold event">,
      Field<"AWHT", 8, 24, ReadWrite, "Analog watchdog high threshold">,
    ], "analog watchdog high threshold register">,
    Register<"DFSDM_FLT0AWLTR", 0x124, 32, [
      Field<"BKAWL", 0, 4, ReadWrite, "Break signal assignment to analog watchdog low threshold event">,
      Field<"AWLT", 8, 24, ReadWrite, "Analog watchdog low threshold">,
    ], "analog watchdog low threshold register">,
    Register<"DFSDM_FLT0AWSR", 0x128, 32, [
      Field<"AWLTF", 0, 8, ReadWrite, "Analog watchdog low threshold flag">,
      Field<"AWHTF", 8, 8, ReadWrite, "Analog watchdog high threshold flag">,
    ], "analog watchdog status register">,
    Register<"DFSDM_FLT0AWCFR", 0x12c, 32, [
      Field<"CLRAWLTF", 0, 8, ReadWrite, "Clear the analog watchdog low threshold flag">,
      Field<"CLRAWHTF", 8, 8, ReadWrite, "Clear the analog watchdog high threshold flag">,
    ], "analog watchdog clear flag register">,
    Register<"DFSDM_FLT0EXMAX", 0x130, 32, [
      Field<"EXMAXCH", 0, 3, ReadWrite, "Extremes detector maximum data channel">,
      Field<"EXMAX", 8, 24, ReadWrite, "Extremes detector maximum value">,
    ], "Extremes detector maximum register">,
    Register<"DFSDM_FLT0EXMIN", 0x134, 32, [
      Field<"EXMINCH", 0, 3, ReadWrite, "Extremes detector minimum data channel">,
      Field<"EXMIN", 8, 24, ReadWrite, "EXMIN">,
    ], "Extremes detector minimum register">,
    Register<"DFSDM_FLT0CNVTIMR", 0x138, 32, [
      Field<"CNVCNT", 4, 28, ReadWrite, "28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN">,
    ], "conversion timer register">,
    Register<"DFSDM_FLT1CR1", 0x180, 32, [
      Field<"DFEN", 0, 1, ReadWrite, "DFSDM enable">,
      Field<"JSWSTART", 1, 1, ReadWrite, "Start a conversion of the injected group of channels">,
      Field<"JSYNC", 3, 1, ReadWrite, "Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger">,
      Field<"JSCAN", 4, 1, ReadWrite, "Scanning conversion mode for injected conversions">,
      Field<"JDMAEN", 5, 1, ReadWrite, "DMA channel enabled to read data for the injected channel group">,
      Field<"JEXTSEL", 8, 3, ReadWrite, "Trigger signal selection for launching injected conversions">,
      Field<"JEXTEN", 13, 2, ReadWrite, "Trigger enable and trigger edge selection for injected conversions">,
      Field<"RSWSTART", 17, 1, ReadWrite, "Software start of a conversion on the regular channel">,
      Field<"RCONT", 18, 1, ReadWrite, "Continuous mode selection for regular conversions">,
      Field<"RSYNC", 19, 1, ReadWrite, "Launch regular conversion synchronously with DFSDM0">,
      Field<"RDMAEN", 21, 1, ReadWrite, "DMA channel enabled to read data for the regular conversion">,
      Field<"RCH", 24, 3, ReadWrite, "Regular channel selection">,
      Field<"FAST", 29, 1, ReadWrite, "Fast conversion mode selection for regular conversions">,
      Field<"AWFSEL", 30, 1, ReadWrite, "Analog watchdog fast mode select">,
    ], "control register 1">,
    Register<"DFSDM_FLT1CR2", 0x184, 32, [
      Field<"JEOCIE", 0, 1, ReadWrite, "Injected end of conversion interrupt enable">,
      Field<"REOCIE", 1, 1, ReadWrite, "Regular end of conversion interrupt enable">,
      Field<"JOVRIE", 2, 1, ReadWrite, "Injected data overrun interrupt enable">,
      Field<"ROVRIE", 3, 1, ReadWrite, "Regular data overrun interrupt enable">,
      Field<"AWDIE", 4, 1, ReadWrite, "Analog watchdog interrupt enable">,
      Field<"SCDIE", 5, 1, ReadWrite, "Short-circuit detector interrupt enable">,
      Field<"CKABIE", 6, 1, ReadWrite, "Clock absence interrupt enable">,
      Field<"EXCH", 8, 8, ReadWrite, "Extremes detector channel selection">,
      Field<"AWDCH", 16, 8, ReadWrite, "Analog watchdog channel selection">,
    ], "control register 2">,
    Register<"DFSDM_FLT1ISR", 0x188, 32, [
      Field<"JEOCF", 0, 1, ReadWrite, "End of injected conversion flag">,
      Field<"REOCF", 1, 1, ReadWrite, "End of regular conversion flag">,
      Field<"JOVRF", 2, 1, ReadWrite, "Injected conversion overrun flag">,
      Field<"ROVRF", 3, 1, ReadWrite, "Regular conversion overrun flag">,
      Field<"AWDF", 4, 1, ReadWrite, "Analog watchdog">,
      Field<"JCIP", 13, 1, ReadWrite, "Injected conversion in progress status">,
      Field<"RCIP", 14, 1, ReadWrite, "Regular conversion in progress status">,
      Field<"CKABF", 16, 8, ReadWrite, "Clock absence flag">,
      Field<"SCDF", 24, 8, ReadWrite, "short-circuit detector flag">,
    ], "interrupt and status register">,
    Register<"DFSDM1_ICR", 0x18c, 32, [
      Field<"CLRJOVRF", 2, 1, ReadWrite, "Clear the injected conversion overrun flag">,
      Field<"CLRROVRF", 3, 1, ReadWrite, "Clear the regular conversion overrun flag">,
      Field<"CLRCKABF", 16, 8, ReadWrite, "Clear the clock absence flag">,
      Field<"CLRSCDF", 24, 8, ReadWrite, "Clear the short-circuit detector flag">,
    ], "interrupt flag clear register">,
    Register<"DFSDM_FLT1JCHGR", 0x190, 32, [
      Field<"JCHG", 0, 8, ReadWrite, "Injected channel group selection">,
    ], "injected channel group selection register">,
    Register<"DFSDM1_FCR", 0x194, 32, [
      Field<"IOSR", 0, 8, ReadWrite, "Integrator oversampling ratio (averaging length)">,
      Field<"FOSR", 16, 10, ReadWrite, "Sinc filter oversampling ratio (decimation rate)">,
      Field<"FORD", 29, 3, ReadWrite, "Sinc filter order">,
    ], "filter control register">,
    Register<"DFSDM_FLT1JDATAR", 0x198, 32, [
      Field<"JDATACH", 0, 3, ReadWrite, "Injected channel most recently converted">,
      Field<"JDATA", 8, 24, ReadWrite, "Injected group conversion data">,
    ], "data register for injected group">,
    Register<"DFSDM_FLT1RDATAR", 0x19c, 32, [
      Field<"RDATACH", 0, 3, ReadWrite, "Regular channel most recently converted">,
      Field<"RPEND", 4, 1, ReadWrite, "Regular channel pending data">,
      Field<"RDATA", 8, 24, ReadWrite, "Regular channel conversion data">,
    ], "data register for the regular channel">,
    Register<"DFSDM_FLT1AWHTR", 0x1a0, 32, [
      Field<"BKAWH", 0, 4, ReadWrite, "Break signal assignment to analog watchdog high threshold event">,
      Field<"AWHT", 8, 24, ReadWrite, "Analog watchdog high threshold">,
    ], "analog watchdog high threshold register">,
    Register<"DFSDM_FLT1AWLTR", 0x1a4, 32, [
      Field<"BKAWL", 0, 4, ReadWrite, "Break signal assignment to analog watchdog low threshold event">,
      Field<"AWLT", 8, 24, ReadWrite, "Analog watchdog low threshold">,
    ], "analog watchdog low threshold register">,
    Register<"DFSDM_FLT1AWSR", 0x1a8, 32, [
      Field<"AWLTF", 0, 8, ReadWrite, "Analog watchdog low threshold flag">,
      Field<"AWHTF", 8, 8, ReadWrite, "Analog watchdog high threshold flag">,
    ], "analog watchdog status register">,
    Register<"DFSDM_FLT1AWCFR", 0x1ac, 32, [
      Field<"CLRAWLTF", 0, 8, ReadWrite, "Clear the analog watchdog low threshold flag">,
      Field<"CLRAWHTF", 8, 8, ReadWrite, "Clear the analog watchdog high threshold flag">,
    ], "analog watchdog clear flag register">,
    Register<"DFSDM_FLT1EXMAX", 0x1b0, 32, [
      Field<"EXMAXCH", 0, 3, ReadWrite, "Extremes detector maximum data channel">,
      Field<"EXMAX", 8, 24, ReadWrite, "Extremes detector maximum value">,
    ], "Extremes detector maximum register">,
    Register<"DFSDM_FLT1EXMIN", 0x1b4, 32, [
      Field<"EXMINCH", 0, 3, ReadWrite, "Extremes detector minimum data channel">,
      Field<"EXMIN", 8, 24, ReadWrite, "EXMIN">,
    ], "Extremes detector minimum register">,
    Register<"DFSDM_FLT1CNVTIMR", 0x1b8, 32, [
      Field<"CNVCNT", 4, 28, ReadWrite, "28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN">,
    ], "conversion timer register">,
    Register<"DFSDM_FLT2CR1", 0x200, 32, [
      Field<"DFEN", 0, 1, ReadWrite, "DFSDM enable">,
      Field<"JSWSTART", 1, 1, ReadWrite, "Start a conversion of the injected group of channels">,
      Field<"JSYNC", 3, 1, ReadWrite, "Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger">,
      Field<"JSCAN", 4, 1, ReadWrite, "Scanning conversion mode for injected conversions">,
      Field<"JDMAEN", 5, 1, ReadWrite, "DMA channel enabled to read data for the injected channel group">,
      Field<"JEXTSEL", 8, 3, ReadWrite, "Trigger signal selection for launching injected conversions">,
      Field<"JEXTEN", 13, 2, ReadWrite, "Trigger enable and trigger edge selection for injected conversions">,
      Field<"RSWSTART", 17, 1, ReadWrite, "Software start of a conversion on the regular channel">,
      Field<"RCONT", 18, 1, ReadWrite, "Continuous mode selection for regular conversions">,
      Field<"RSYNC", 19, 1, ReadWrite, "Launch regular conversion synchronously with DFSDM0">,
      Field<"RDMAEN", 21, 1, ReadWrite, "DMA channel enabled to read data for the regular conversion">,
      Field<"RCH", 24, 3, ReadWrite, "Regular channel selection">,
      Field<"FAST", 29, 1, ReadWrite, "Fast conversion mode selection for regular conversions">,
      Field<"AWFSEL", 30, 1, ReadWrite, "Analog watchdog fast mode select">,
    ], "control register 1">,
    Register<"DFSDM_FLT2CR2", 0x204, 32, [
      Field<"JEOCIE", 0, 1, ReadWrite, "Injected end of conversion interrupt enable">,
      Field<"REOCIE", 1, 1, ReadWrite, "Regular end of conversion interrupt enable">,
      Field<"JOVRIE", 2, 1, ReadWrite, "Injected data overrun interrupt enable">,
      Field<"ROVRIE", 3, 1, ReadWrite, "Regular data overrun interrupt enable">,
      Field<"AWDIE", 4, 1, ReadWrite, "Analog watchdog interrupt enable">,
      Field<"SCDIE", 5, 1, ReadWrite, "Short-circuit detector interrupt enable">,
      Field<"CKABIE", 6, 1, ReadWrite, "Clock absence interrupt enable">,
      Field<"EXCH", 8, 8, ReadWrite, "Extremes detector channel selection">,
      Field<"AWDCH", 16, 8, ReadWrite, "Analog watchdog channel selection">,
    ], "control register 2">,
    Register<"DFSDM_FLT2ISR", 0x208, 32, [
      Field<"JEOCF", 0, 1, ReadWrite, "End of injected conversion flag">,
      Field<"REOCF", 1, 1, ReadWrite, "End of regular conversion flag">,
      Field<"JOVRF", 2, 1, ReadWrite, "Injected conversion overrun flag">,
      Field<"ROVRF", 3, 1, ReadWrite, "Regular conversion overrun flag">,
      Field<"AWDF", 4, 1, ReadWrite, "Analog watchdog">,
      Field<"JCIP", 13, 1, ReadWrite, "Injected conversion in progress status">,
      Field<"RCIP", 14, 1, ReadWrite, "Regular conversion in progress status">,
      Field<"CKABF", 16, 8, ReadWrite, "Clock absence flag">,
      Field<"SCDF", 24, 8, ReadWrite, "short-circuit detector flag">,
    ], "interrupt and status register">,
    Register<"DFSDM_FLT2ICR", 0x20c, 32, [
      Field<"CLRJOVRF", 2, 1, ReadWrite, "Clear the injected conversion overrun flag">,
      Field<"CLRROVRF", 3, 1, ReadWrite, "Clear the regular conversion overrun flag">,
      Field<"CLRCKABF", 16, 8, ReadWrite, "Clear the clock absence flag">,
      Field<"CLRSCDF", 24, 8, ReadWrite, "Clear the short-circuit detector flag">,
    ], "interrupt flag clear register">,
    Register<"DFSDM_FLT2JCHGR", 0x210, 32, [
      Field<"JCHG", 0, 8, ReadWrite, "Injected channel group selection">,
    ], "injected channel group selection register">,
    Register<"DFSDM_FLT2FCR", 0x214, 32, [
      Field<"IOSR", 0, 8, ReadWrite, "Integrator oversampling ratio (averaging length)">,
      Field<"FOSR", 16, 10, ReadWrite, "Sinc filter oversampling ratio (decimation rate)">,
      Field<"FORD", 29, 3, ReadWrite, "Sinc filter order">,
    ], "filter control register">,
    Register<"DFSDM_FLT2JDATAR", 0x218, 32, [
      Field<"JDATACH", 0, 3, ReadWrite, "Injected channel most recently converted">,
      Field<"JDATA", 8, 24, ReadWrite, "Injected group conversion data">,
    ], "data register for injected group">,
    Register<"DFSDM_FLT2RDATAR", 0x21c, 32, [
      Field<"RDATACH", 0, 3, ReadWrite, "Regular channel most recently converted">,
      Field<"RPEND", 4, 1, ReadWrite, "Regular channel pending data">,
      Field<"RDATA", 8, 24, ReadWrite, "Regular channel conversion data">,
    ], "data register for the regular channel">,
    Register<"DFSDM_FLT2AWHTR", 0x220, 32, [
      Field<"BKAWH", 0, 4, ReadWrite, "Break signal assignment to analog watchdog high threshold event">,
      Field<"AWHT", 8, 24, ReadWrite, "Analog watchdog high threshold">,
    ], "analog watchdog high threshold register">,
    Register<"DFSDM_FLT2AWLTR", 0x224, 32, [
      Field<"BKAWL", 0, 4, ReadWrite, "Break signal assignment to analog watchdog low threshold event">,
      Field<"AWLT", 8, 24, ReadWrite, "Analog watchdog low threshold">,
    ], "analog watchdog low threshold register">,
    Register<"DFSDM_FLT2AWSR", 0x228, 32, [
      Field<"AWLTF", 0, 8, ReadWrite, "Analog watchdog low threshold flag">,
      Field<"AWHTF", 8, 8, ReadWrite, "Analog watchdog high threshold flag">,
    ], "analog watchdog status register">,
    Register<"DFSDM_FLT2AWCFR", 0x22c, 32, [
      Field<"CLRAWLTF", 0, 8, ReadWrite, "Clear the analog watchdog low threshold flag">,
      Field<"CLRAWHTF", 8, 8, ReadWrite, "Clear the analog watchdog high threshold flag">,
    ], "analog watchdog clear flag register">,
    Register<"DFSDM_FLT2EXMAX", 0x230, 32, [
      Field<"EXMAXCH", 0, 3, ReadWrite, "Extremes detector maximum data channel">,
      Field<"EXMAX", 8, 24, ReadWrite, "Extremes detector maximum value">,
    ], "Extremes detector maximum register">,
    Register<"DFSDM_FLT2EXMIN", 0x234, 32, [
      Field<"EXMINCH", 0, 3, ReadWrite, "Extremes detector minimum data channel">,
      Field<"EXMIN", 8, 24, ReadWrite, "EXMIN">,
    ], "Extremes detector minimum register">,
    Register<"DFSDM_FLT2CNVTIMR", 0x238, 32, [
      Field<"CNVCNT", 4, 28, ReadWrite, "28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN">,
    ], "conversion timer register">,
    Register<"DFSDM_FLT3CR1", 0x280, 32, [
      Field<"DFEN", 0, 1, ReadWrite, "DFSDM enable">,
      Field<"JSWSTART", 1, 1, ReadWrite, "Start a conversion of the injected group of channels">,
      Field<"JSYNC", 3, 1, ReadWrite, "Launch an injected conversion synchronously with the DFSDM0 JSWSTART trigger">,
      Field<"JSCAN", 4, 1, ReadWrite, "Scanning conversion mode for injected conversions">,
      Field<"JDMAEN", 5, 1, ReadWrite, "DMA channel enabled to read data for the injected channel group">,
      Field<"JEXTSEL", 8, 3, ReadWrite, "Trigger signal selection for launching injected conversions">,
      Field<"JEXTEN", 13, 2, ReadWrite, "Trigger enable and trigger edge selection for injected conversions">,
      Field<"RSWSTART", 17, 1, ReadWrite, "Software start of a conversion on the regular channel">,
      Field<"RCONT", 18, 1, ReadWrite, "Continuous mode selection for regular conversions">,
      Field<"RSYNC", 19, 1, ReadWrite, "Launch regular conversion synchronously with DFSDM0">,
      Field<"RDMAEN", 21, 1, ReadWrite, "DMA channel enabled to read data for the regular conversion">,
      Field<"RCH", 24, 3, ReadWrite, "Regular channel selection">,
      Field<"FAST", 29, 1, ReadWrite, "Fast conversion mode selection for regular conversions">,
      Field<"AWFSEL", 30, 1, ReadWrite, "Analog watchdog fast mode select">,
    ], "control register 1">,
    Register<"DFSDM_FLT3CR2", 0x284, 32, [
      Field<"JEOCIE", 0, 1, ReadWrite, "Injected end of conversion interrupt enable">,
      Field<"REOCIE", 1, 1, ReadWrite, "Regular end of conversion interrupt enable">,
      Field<"JOVRIE", 2, 1, ReadWrite, "Injected data overrun interrupt enable">,
      Field<"ROVRIE", 3, 1, ReadWrite, "Regular data overrun interrupt enable">,
      Field<"AWDIE", 4, 1, ReadWrite, "Analog watchdog interrupt enable">,
      Field<"SCDIE", 5, 1, ReadWrite, "Short-circuit detector interrupt enable">,
      Field<"CKABIE", 6, 1, ReadWrite, "Clock absence interrupt enable">,
      Field<"EXCH", 8, 8, ReadWrite, "Extremes detector channel selection">,
      Field<"AWDCH", 16, 8, ReadWrite, "Analog watchdog channel selection">,
    ], "control register 2">,
    Register<"DFSDM_FLT3ISR", 0x288, 32, [
      Field<"JEOCF", 0, 1, ReadWrite, "End of injected conversion flag">,
      Field<"REOCF", 1, 1, ReadWrite, "End of regular conversion flag">,
      Field<"JOVRF", 2, 1, ReadWrite, "Injected conversion overrun flag">,
      Field<"ROVRF", 3, 1, ReadWrite, "Regular conversion overrun flag">,
      Field<"AWDF", 4, 1, ReadWrite, "Analog watchdog">,
      Field<"JCIP", 13, 1, ReadWrite, "Injected conversion in progress status">,
      Field<"RCIP", 14, 1, ReadWrite, "Regular conversion in progress status">,
      Field<"CKABF", 16, 8, ReadWrite, "Clock absence flag">,
      Field<"SCDF", 24, 8, ReadWrite, "short-circuit detector flag">,
    ], "interrupt and status register">,
    Register<"DFSDM_FLT3ICR", 0x28c, 32, [
      Field<"CLRJOVRF", 2, 1, ReadWrite, "Clear the injected conversion overrun flag">,
      Field<"CLRROVRF", 3, 1, ReadWrite, "Clear the regular conversion overrun flag">,
      Field<"CLRCKABF", 16, 8, ReadWrite, "Clear the clock absence flag">,
      Field<"CLRSCDF", 24, 8, ReadWrite, "Clear the short-circuit detector flag">,
    ], "interrupt flag clear register">,
    Register<"DFSDM_FLT3JCHGR", 0x290, 32, [
      Field<"JCHG", 0, 8, ReadWrite, "Injected channel group selection">,
    ], "injected channel group selection register">,
    Register<"DFSDM_FLT3FCR", 0x294, 32, [
      Field<"IOSR", 0, 8, ReadWrite, "Integrator oversampling ratio (averaging length)">,
      Field<"FOSR", 16, 10, ReadWrite, "Sinc filter oversampling ratio (decimation rate)">,
      Field<"FORD", 29, 3, ReadWrite, "Sinc filter order">,
    ], "filter control register">,
    Register<"DFSDM_FLT3JDATAR", 0x298, 32, [
      Field<"JDATACH", 0, 3, ReadWrite, "Injected channel most recently converted">,
      Field<"JDATA", 8, 24, ReadWrite, "Injected group conversion data">,
    ], "data register for injected group">,
    Register<"DFSDM_FLT3RDATAR", 0x29c, 32, [
      Field<"RDATACH", 0, 3, ReadWrite, "Regular channel most recently converted">,
      Field<"RPEND", 4, 1, ReadWrite, "Regular channel pending data">,
      Field<"RDATA", 8, 24, ReadWrite, "Regular channel conversion data">,
    ], "data register for the regular channel">,
    Register<"DFSDM_FLT3AWHTR", 0x2a0, 32, [
      Field<"BKAWH", 0, 4, ReadWrite, "Break signal assignment to analog watchdog high threshold event">,
      Field<"AWHT", 8, 24, ReadWrite, "Analog watchdog high threshold">,
    ], "analog watchdog high threshold register">,
    Register<"DFSDM_FLT3AWLTR", 0x2a4, 32, [
      Field<"BKAWL", 0, 4, ReadWrite, "Break signal assignment to analog watchdog low threshold event">,
      Field<"AWLT", 8, 24, ReadWrite, "Analog watchdog low threshold">,
    ], "analog watchdog low threshold register">,
    Register<"DFSDM_FLT3AWSR", 0x2a8, 32, [
      Field<"AWLTF", 0, 8, ReadWrite, "Analog watchdog low threshold flag">,
      Field<"AWHTF", 8, 8, ReadWrite, "Analog watchdog high threshold flag">,
    ], "analog watchdog status register">,
    Register<"DFSDM_FLT3AWCFR", 0x2ac, 32, [
      Field<"CLRAWLTF", 0, 8, ReadWrite, "Clear the analog watchdog low threshold flag">,
      Field<"CLRAWHTF", 8, 8, ReadWrite, "Clear the analog watchdog high threshold flag">,
    ], "analog watchdog clear flag register">,
    Register<"DFSDM_FLT3EXMAX", 0x2b0, 32, [
      Field<"EXMAXCH", 0, 3, ReadWrite, "Extremes detector maximum data channel">,
      Field<"EXMAX", 8, 24, ReadWrite, "Extremes detector maximum value">,
    ], "Extremes detector maximum register">,
    Register<"DFSDM_FLT3EXMIN", 0x2b4, 32, [
      Field<"EXMINCH", 0, 3, ReadWrite, "Extremes detector minimum data channel">,
      Field<"EXMIN", 8, 24, ReadWrite, "EXMIN">,
    ], "Extremes detector minimum register">,
    Register<"DFSDM_FLT3CNVTIMR", 0x2b8, 32, [
      Field<"CNVCNT", 4, 28, ReadWrite, "28-bit timer counting conversion time t = CNVCNT[27:0] / fDFSDM_CKIN">,
    ], "conversion timer register">,
  ];
}

class DFSDMInstance<string Name, int Base> : PeripheralInstance<Name, Base, DFSDMPeripheral>;
def DFSDM : DFSDMInstance<"DFSDM", 0x40017000>;

#endif // _PERIPHERALS_DFSDM_TD
