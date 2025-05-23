#ifndef _PERIPHERALS_AXI_TD
#define _PERIPHERALS_AXI_TD

include "base.td"

def AXIPeripheral : PeripheralType<"AXI", "AXI interconnect registers"> {
  let accessWidth = 32;
  let registers = [
    Register<"AXI_PERIPH_ID_4", 0x1fd0, 32, [
      Field<"JEP106CON", 0, 4, ReadWrite, "JEP106 continuation code">,
      Field<"KCOUNT4", 4, 4, ReadWrite, "Register file size">,
    ], "AXI interconnect - peripheral ID4 register">,
    Register<"AXI_PERIPH_ID_0", 0x1fe0, 32, [
      Field<"PARTNUM", 0, 8, ReadWrite, "Peripheral part number bits 0 to 7">,
    ], "AXI interconnect - peripheral ID0 register">,
    Register<"AXI_PERIPH_ID_1", 0x1fe4, 32, [
      Field<"PARTNUM", 0, 4, ReadWrite, "Peripheral part number bits 8 to 11">,
      Field<"JEP106I", 4, 4, ReadWrite, "JEP106 identity bits 0 to 3">,
    ], "AXI interconnect - peripheral ID1 register">,
    Register<"AXI_PERIPH_ID_2", 0x1fe8, 32, [
      Field<"JEP106ID", 0, 3, ReadWrite, "JEP106 Identity bits 4 to 6">,
      Field<"JEDEC", 3, 1, ReadWrite, "JEP106 code flag">,
      Field<"REVISION", 4, 4, ReadWrite, "Peripheral revision number">,
    ], "AXI interconnect - peripheral ID2 register">,
    Register<"AXI_PERIPH_ID_3", 0x1fec, 32, [
      Field<"CUST_MOD_NUM", 0, 4, ReadWrite, "Customer modification">,
      Field<"REV_AND", 4, 4, ReadWrite, "Customer version">,
    ], "AXI interconnect - peripheral ID3 register">,
    Register<"AXI_COMP_ID_0", 0x1ff0, 32, [
      Field<"PREAMBLE", 0, 8, ReadWrite, "Preamble bits 0 to 7">,
    ], "AXI interconnect - component ID0 register">,
    Register<"AXI_COMP_ID_1", 0x1ff4, 32, [
      Field<"PREAMBLE", 0, 4, ReadWrite, "Preamble bits 8 to 11">,
      Field<"CLASS", 4, 4, ReadWrite, "Component class">,
    ], "AXI interconnect - component ID1 register">,
    Register<"AXI_COMP_ID_2", 0x1ff8, 32, [
      Field<"PREAMBLE", 0, 8, ReadWrite, "Preamble bits 12 to 19">,
    ], "AXI interconnect - component ID2 register">,
    Register<"AXI_COMP_ID_3", 0x1ffc, 32, [
      Field<"PREAMBLE", 0, 8, ReadWrite, "Preamble bits 20 to 27">,
    ], "AXI interconnect - component ID3 register">,
    Register<"AXI_TARG1_FN_MOD_ISS_BM", 0x2008, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "READ_ISS_OVERRIDE">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Switch matrix write issuing override for target">,
    ], "AXI interconnect - TARG x bus matrix issuing functionality register">,
    Register<"AXI_TARG2_FN_MOD_ISS_BM", 0x3008, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "READ_ISS_OVERRIDE">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Switch matrix write issuing override for target">,
    ], "AXI interconnect - TARG x bus matrix issuing functionality register">,
    Register<"AXI_TARG3_FN_MOD_ISS_BM", 0x4008, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "READ_ISS_OVERRIDE">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Switch matrix write issuing override for target">,
    ], "AXI interconnect - TARG x bus matrix issuing functionality register">,
    Register<"AXI_TARG4_FN_MOD_ISS_BM", 0x5008, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "READ_ISS_OVERRIDE">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Switch matrix write issuing override for target">,
    ], "AXI interconnect - TARG x bus matrix issuing functionality register">,
    Register<"AXI_TARG5_FN_MOD_ISS_BM", 0x6008, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "READ_ISS_OVERRIDE">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Switch matrix write issuing override for target">,
    ], "AXI interconnect - TARG x bus matrix issuing functionality register">,
    Register<"AXI_TARG6_FN_MOD_ISS_BM", 0x7008, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "READ_ISS_OVERRIDE">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Switch matrix write issuing override for target">,
    ], "AXI interconnect - TARG x bus matrix issuing functionality register">,
    Register<"AXI_TARG7_FN_MOD_ISS_BM", 0x800c, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "READ_ISS_OVERRIDE">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Switch matrix write issuing override for target">,
    ], "AXI interconnect - TARG x bus matrix issuing functionality register">,
    Register<"AXI_TARG1_FN_MOD2", 0x2024, 32, [
      Field<"BYPASS_MERGE", 0, 1, ReadWrite, "Disable packing of beats to match the output data width">,
    ], "AXI interconnect - TARG x bus matrix functionality 2 register">,
    Register<"AXI_TARG2_FN_MOD2", 0x3024, 32, [
      Field<"BYPASS_MERGE", 0, 1, ReadWrite, "Disable packing of beats to match the output data width">,
    ], "AXI interconnect - TARG x bus matrix functionality 2 register">,
    Register<"AXI_TARG7_FN_MOD2", 0x8024, 32, [
      Field<"BYPASS_MERGE", 0, 1, ReadWrite, "Disable packing of beats to match the output data width">,
    ], "AXI interconnect - TARG x bus matrix functionality 2 register">,
    Register<"AXI_TARG1_FN_MOD_LB", 0x202c, 32, [
      Field<"FN_MOD_LB", 0, 1, ReadWrite, "Controls burst breaking of long bursts">,
    ], "AXI interconnect - TARG x long burst functionality modification">,
    Register<"AXI_TARG2_FN_MOD_LB", 0x302c, 32, [
      Field<"FN_MOD_LB", 0, 1, ReadWrite, "Controls burst breaking of long bursts">,
    ], "AXI interconnect - TARG x long burst functionality modification">,
    Register<"AXI_TARG1_FN_MOD", 0x2108, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "Override AMIB read issuing capability">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Override AMIB write issuing capability">,
    ], "AXI interconnect - TARG x long burst functionality modification">,
    Register<"AXI_TARG2_FN_MOD", 0x3108, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "Override AMIB read issuing capability">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Override AMIB write issuing capability">,
    ], "AXI interconnect - TARG x long burst functionality modification">,
    Register<"AXI_TARG7_FN_MOD", 0x8108, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "Override AMIB read issuing capability">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Override AMIB write issuing capability">,
    ], "AXI interconnect - TARG x long burst functionality modification">,
    Register<"AXI_INI1_FN_MOD2", 0x42024, 32, [
      Field<"BYPASS_MERGE", 0, 1, ReadWrite, "Disables alteration of transactions by the up-sizer unless required by the protocol">,
    ], "AXI interconnect - INI x functionality modification 2 register">,
    Register<"AXI_INI3_FN_MOD2", 0x44024, 32, [
      Field<"BYPASS_MERGE", 0, 1, ReadWrite, "Disables alteration of transactions by the up-sizer unless required by the protocol">,
    ], "AXI interconnect - INI x functionality modification 2 register">,
    Register<"AXI_INI1_FN_MOD_AHB", 0x42028, 32, [
      Field<"RD_INC_OVERRIDE", 0, 1, ReadWrite, "Converts all AHB-Lite write transactions to a series of single beat AXI">,
      Field<"WR_INC_OVERRIDE", 1, 1, ReadWrite, "Converts all AHB-Lite read transactions to a series of single beat AXI">,
    ], "AXI interconnect - INI x AHB functionality modification register">,
    Register<"AXI_INI3_FN_MOD_AHB", 0x44028, 32, [
      Field<"RD_INC_OVERRIDE", 0, 1, ReadWrite, "Converts all AHB-Lite write transactions to a series of single beat AXI">,
      Field<"WR_INC_OVERRIDE", 1, 1, ReadWrite, "Converts all AHB-Lite read transactions to a series of single beat AXI">,
    ], "AXI interconnect - INI x AHB functionality modification register">,
    Register<"AXI_INI1_READ_QOS", 0x42100, 32, [
      Field<"AR_QOS", 0, 4, ReadWrite, "Read channel QoS setting">,
    ], "AXI interconnect - INI x read QoS register">,
    Register<"AXI_INI2_READ_QOS", 0x43100, 32, [
      Field<"AR_QOS", 0, 4, ReadWrite, "Read channel QoS setting">,
    ], "AXI interconnect - INI x read QoS register">,
    Register<"AXI_INI3_READ_QOS", 0x44100, 32, [
      Field<"AR_QOS", 0, 4, ReadWrite, "Read channel QoS setting">,
    ], "AXI interconnect - INI x read QoS register">,
    Register<"AXI_INI4_READ_QOS", 0x45100, 32, [
      Field<"AR_QOS", 0, 4, ReadWrite, "Read channel QoS setting">,
    ], "AXI interconnect - INI x read QoS register">,
    Register<"AXI_INI5_READ_QOS", 0x46100, 32, [
      Field<"AR_QOS", 0, 4, ReadWrite, "Read channel QoS setting">,
    ], "AXI interconnect - INI x read QoS register">,
    Register<"AXI_INI6_READ_QOS", 0x47100, 32, [
      Field<"AR_QOS", 0, 4, ReadWrite, "Read channel QoS setting">,
    ], "AXI interconnect - INI x read QoS register">,
    Register<"AXI_INI1_WRITE_QOS", 0x42104, 32, [
      Field<"AW_QOS", 0, 4, ReadWrite, "Write channel QoS setting">,
    ], "AXI interconnect - INI x write QoS register">,
    Register<"AXI_INI2_WRITE_QOS", 0x43104, 32, [
      Field<"AW_QOS", 0, 4, ReadWrite, "Write channel QoS setting">,
    ], "AXI interconnect - INI x write QoS register">,
    Register<"AXI_INI3_WRITE_QOS", 0x44104, 32, [
      Field<"AW_QOS", 0, 4, ReadWrite, "Write channel QoS setting">,
    ], "AXI interconnect - INI x write QoS register">,
    Register<"AXI_INI4_WRITE_QOS", 0x45104, 32, [
      Field<"AW_QOS", 0, 4, ReadWrite, "Write channel QoS setting">,
    ], "AXI interconnect - INI x write QoS register">,
    Register<"AXI_INI5_WRITE_QOS", 0x46104, 32, [
      Field<"AW_QOS", 0, 4, ReadWrite, "Write channel QoS setting">,
    ], "AXI interconnect - INI x write QoS register">,
    Register<"AXI_INI6_WRITE_QOS", 0x47104, 32, [
      Field<"AW_QOS", 0, 4, ReadWrite, "Write channel QoS setting">,
    ], "AXI interconnect - INI x write QoS register">,
    Register<"AXI_INI1_FN_MOD", 0x42108, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "Override ASIB read issuing capability">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Override ASIB write issuing capability">,
    ], "AXI interconnect - INI x issuing functionality modification register">,
    Register<"AXI_INI2_FN_MOD", 0x43108, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "Override ASIB read issuing capability">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Override ASIB write issuing capability">,
    ], "AXI interconnect - INI x issuing functionality modification register">,
    Register<"AXI_INI3_FN_MOD", 0x44108, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "Override ASIB read issuing capability">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Override ASIB write issuing capability">,
    ], "AXI interconnect - INI x issuing functionality modification register">,
    Register<"AXI_INI4_FN_MOD", 0x45108, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "Override ASIB read issuing capability">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Override ASIB write issuing capability">,
    ], "AXI interconnect - INI x issuing functionality modification register">,
    Register<"AXI_INI5_FN_MOD", 0x46108, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "Override ASIB read issuing capability">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Override ASIB write issuing capability">,
    ], "AXI interconnect - INI x issuing functionality modification register">,
    Register<"AXI_INI6_FN_MOD", 0x47108, 32, [
      Field<"READ_ISS_OVERRIDE", 0, 1, ReadWrite, "Override ASIB read issuing capability">,
      Field<"WRITE_ISS_OVERRIDE", 1, 1, ReadWrite, "Override ASIB write issuing capability">,
    ], "AXI interconnect - INI x issuing functionality modification register">,
  ];
}

class AXIInstance<string Name, int Base> : PeripheralInstance<Name, Base, AXIPeripheral>;
def AXI : AXIInstance<"AXI", 0x51000000>;

#endif // _PERIPHERALS_AXI_TD
