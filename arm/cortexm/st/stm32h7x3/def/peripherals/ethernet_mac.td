#ifndef _PERIPHERALS_ETHERNET_MAC_TD
#define _PERIPHERALS_ETHERNET_MAC_TD

include "base.td"

def Ethernet_MACPeripheral : PeripheralType<"Ethernet_MAC", "Ethernet: media access control (MAC)"> {
  let accessWidth = 32;
  let registers = [
    Register<"MACCR", 0x0, 32, [
      Field<"RE", 0, 1, ReadWrite, "Receiver Enable">,
      Field<"TE", 1, 1, ReadWrite, "TE">,
      Field<"PRELEN", 2, 2, ReadWrite, "PRELEN">,
      Field<"DC", 4, 1, ReadWrite, "DC">,
      Field<"BL", 5, 2, ReadWrite, "BL">,
      Field<"DR", 8, 1, ReadWrite, "DR">,
      Field<"DCRS", 9, 1, ReadWrite, "DCRS">,
      Field<"DO", 10, 1, ReadWrite, "DO">,
      Field<"ECRSFD", 11, 1, ReadWrite, "ECRSFD">,
      Field<"LM", 12, 1, ReadWrite, "LM">,
      Field<"DM", 13, 1, ReadWrite, "DM">,
      Field<"FES", 14, 1, ReadWrite, "FES">,
      Field<"JE", 16, 1, ReadWrite, "JE">,
      Field<"JD", 17, 1, ReadWrite, "JD">,
      Field<"WD", 19, 1, ReadWrite, "WD">,
      Field<"ACS", 20, 1, ReadWrite, "ACS">,
      Field<"CST", 21, 1, ReadWrite, "CST">,
      Field<"S2KP", 22, 1, ReadWrite, "S2KP">,
      Field<"GPSLCE", 23, 1, ReadWrite, "GPSLCE">,
      Field<"IPG", 24, 3, ReadWrite, "IPG">,
      Field<"IPC", 27, 1, ReadWrite, "IPC">,
      Field<"SARC", 28, 3, ReadWrite, "SARC">,
      Field<"ARPEN", 31, 1, ReadWrite, "ARPEN">,
    ], "Operating mode configuration register">,
    Register<"MACECR", 0x4, 32, [
      Field<"GPSL", 0, 14, ReadWrite, "GPSL">,
      Field<"DCRCC", 16, 1, ReadWrite, "DCRCC">,
      Field<"SPEN", 17, 1, ReadWrite, "SPEN">,
      Field<"USP", 18, 1, ReadWrite, "USP">,
      Field<"EIPGEN", 24, 1, ReadWrite, "EIPGEN">,
      Field<"EIPG", 25, 5, ReadWrite, "EIPG">,
    ], "Extended operating mode configuration register">,
    Register<"MACPFR", 0x8, 32, [
      Field<"PR", 0, 1, ReadWrite, "PR">,
      Field<"HUC", 1, 1, ReadWrite, "HUC">,
      Field<"HMC", 2, 1, ReadWrite, "HMC">,
      Field<"DAIF", 3, 1, ReadWrite, "DAIF">,
      Field<"PM", 4, 1, ReadWrite, "PM">,
      Field<"DBF", 5, 1, ReadWrite, "DBF">,
      Field<"PCF", 6, 2, ReadWrite, "PCF">,
      Field<"SAIF", 8, 1, ReadWrite, "SAIF">,
      Field<"SAF", 9, 1, ReadWrite, "SAF">,
      Field<"HPF", 10, 1, ReadWrite, "HPF">,
      Field<"VTFE", 16, 1, ReadWrite, "VTFE">,
      Field<"IPFE", 20, 1, ReadWrite, "IPFE">,
      Field<"DNTU", 21, 1, ReadWrite, "DNTU">,
      Field<"RA", 31, 1, ReadWrite, "RA">,
    ], "Packet filtering control register">,
    Register<"MACWTR", 0xc, 32, [
      Field<"WTO", 0, 4, ReadWrite, "WTO">,
      Field<"PWE", 8, 1, ReadWrite, "PWE">,
    ], "Watchdog timeout register">,
    Register<"MACHT0R", 0x10, 32, [
      Field<"HT31T0", 0, 32, ReadWrite, "HT31T0">,
    ], "Hash Table 0 register">,
    Register<"MACHT1R", 0x14, 32, [
      Field<"HT63T32", 0, 32, ReadWrite, "HT63T32">,
    ], "Hash Table 1 register">,
    Register<"MACVTR", 0x50, 32, [
      Field<"VL", 0, 16, ReadWrite, "VL">,
      Field<"ETV", 16, 1, ReadWrite, "ETV">,
      Field<"VTIM", 17, 1, ReadWrite, "VTIM">,
      Field<"ESVL", 18, 1, ReadWrite, "ESVL">,
      Field<"ERSVLM", 19, 1, ReadWrite, "ERSVLM">,
      Field<"DOVLTC", 20, 1, ReadWrite, "DOVLTC">,
      Field<"EVLS", 21, 2, ReadWrite, "EVLS">,
      Field<"EVLRXS", 24, 1, ReadWrite, "EVLRXS">,
      Field<"VTHM", 25, 1, ReadWrite, "VTHM">,
      Field<"EDVLP", 26, 1, ReadWrite, "EDVLP">,
      Field<"ERIVLT", 27, 1, ReadWrite, "ERIVLT">,
      Field<"EIVLS", 28, 2, ReadWrite, "EIVLS">,
      Field<"EIVLRXS", 31, 1, ReadWrite, "EIVLRXS">,
    ], "VLAN tag register">,
    Register<"MACVHTR", 0x58, 32, [
      Field<"VLHT", 0, 16, ReadWrite, "VLHT">,
    ], "VLAN Hash table register">,
    Register<"MACVIR", 0x60, 32, [
      Field<"VLT", 0, 16, ReadWrite, "VLT">,
      Field<"VLC", 16, 2, ReadWrite, "VLC">,
      Field<"VLP", 18, 1, ReadWrite, "VLP">,
      Field<"CSVL", 19, 1, ReadWrite, "CSVL">,
      Field<"VLTI", 20, 1, ReadWrite, "VLTI">,
    ], "VLAN inclusion register">,
    Register<"MACIVIR", 0x64, 32, [
      Field<"VLT", 0, 16, ReadWrite, "VLT">,
      Field<"VLC", 16, 2, ReadWrite, "VLC">,
      Field<"VLP", 18, 1, ReadWrite, "VLP">,
      Field<"CSVL", 19, 1, ReadWrite, "CSVL">,
      Field<"VLTI", 20, 1, ReadWrite, "VLTI">,
    ], "Inner VLAN inclusion register">,
    Register<"MACQTxFCR", 0x70, 32, [
      Field<"FCB_BPA", 0, 1, ReadWrite, "FCB_BPA">,
      Field<"TFE", 1, 1, ReadWrite, "TFE">,
      Field<"PLT", 4, 3, ReadWrite, "PLT">,
      Field<"DZPQ", 7, 1, ReadWrite, "DZPQ">,
      Field<"PT", 16, 16, ReadWrite, "PT">,
    ], "Tx Queue flow control register">,
    Register<"MACRxFCR", 0x90, 32, [
      Field<"RFE", 0, 1, ReadWrite, "RFE">,
      Field<"UP", 1, 1, ReadWrite, "UP">,
    ], "Rx flow control register">,
    Register<"MACISR", 0xb0, 32, [
      Field<"PHYIS", 3, 1, ReadWrite, "PHYIS">,
      Field<"PMTIS", 4, 1, ReadWrite, "PMTIS">,
      Field<"LPIIS", 5, 1, ReadWrite, "LPIIS">,
      Field<"MMCIS", 8, 1, ReadWrite, "MMCIS">,
      Field<"MMCRXIS", 9, 1, ReadWrite, "MMCRXIS">,
      Field<"MMCTXIS", 10, 1, ReadWrite, "MMCTXIS">,
      Field<"TSIS", 12, 1, ReadWrite, "TSIS">,
      Field<"TXSTSIS", 13, 1, ReadWrite, "TXSTSIS">,
      Field<"RXSTSIS", 14, 1, ReadWrite, "RXSTSIS">,
    ], "Interrupt status register">,
    Register<"MACIER", 0xb4, 32, [
      Field<"PHYIE", 3, 1, ReadWrite, "PHYIE">,
      Field<"PMTIE", 4, 1, ReadWrite, "PMTIE">,
      Field<"LPIIE", 5, 1, ReadWrite, "LPIIE">,
      Field<"TSIE", 12, 1, ReadWrite, "TSIE">,
      Field<"TXSTSIE", 13, 1, ReadWrite, "TXSTSIE">,
      Field<"RXSTSIE", 14, 1, ReadWrite, "RXSTSIE">,
    ], "Interrupt enable register">,
    Register<"MACRxTxSR", 0xb8, 32, [
      Field<"TJT", 0, 1, ReadWrite, "TJT">,
      Field<"NCARR", 1, 1, ReadWrite, "NCARR">,
      Field<"LCARR", 2, 1, ReadWrite, "LCARR">,
      Field<"EXDEF", 3, 1, ReadWrite, "EXDEF">,
      Field<"LCOL", 4, 1, ReadWrite, "LCOL">,
      Field<"EXCOL", 5, 1, ReadWrite, "LCOL">,
      Field<"RWT", 8, 1, ReadWrite, "RWT">,
    ], "Rx Tx status register">,
    Register<"MACPCSR", 0xc0, 32, [
      Field<"PWRDWN", 0, 1, ReadWrite, "PWRDWN">,
      Field<"MGKPKTEN", 1, 1, ReadWrite, "MGKPKTEN">,
      Field<"RWKPKTEN", 2, 1, ReadWrite, "RWKPKTEN">,
      Field<"MGKPRCVD", 5, 1, Read, "MGKPRCVD">,
      Field<"RWKPRCVD", 6, 1, Read, "RWKPRCVD">,
      Field<"GLBLUCAST", 9, 1, ReadWrite, "GLBLUCAST">,
      Field<"RWKPFE", 10, 1, ReadWrite, "RWKPFE">,
      Field<"RWKPTR", 24, 5, ReadWrite, "RWKPTR">,
      Field<"RWKFILTRST", 31, 1, ReadWrite, "RWKFILTRST">,
    ], "PMT control status register">,
    Register<"MACRWKPFR", 0xc4, 32, [
      Field<"MACRWKPFR", 0, 32, ReadWrite, "MACRWKPFR">,
    ], "Remove wakeup packet filter register">,
    Register<"MACLCSR", 0xd0, 32, [
      Field<"TLPIEN", 0, 1, Read, "TLPIEN">,
      Field<"TLPIEX", 1, 1, Read, "TLPIEX">,
      Field<"RLPIEN", 2, 1, Read, "RLPIEN">,
      Field<"RLPIEX", 3, 1, Read, "RLPIEX">,
      Field<"TLPIST", 8, 1, Read, "TLPIST">,
      Field<"RLPIST", 9, 1, Read, "RLPIST">,
      Field<"LPIEN", 16, 1, ReadWrite, "LPIEN">,
      Field<"PLS", 17, 1, ReadWrite, "PLS">,
      Field<"PLSEN", 18, 1, ReadWrite, "PLSEN">,
      Field<"LPITXA", 19, 1, ReadWrite, "LPITXA">,
      Field<"LPITE", 20, 1, ReadWrite, "LPITE">,
    ], "LPI control status register">,
    Register<"MACLTCR", 0xd4, 32, [
      Field<"TWT", 0, 16, ReadWrite, "TWT">,
      Field<"LST", 16, 10, ReadWrite, "LST">,
    ], "LPI timers control register">,
    Register<"MACLETR", 0xd8, 32, [
      Field<"LPIET", 0, 17, ReadWrite, "LPIET">,
    ], "LPI entry timer register">,
    Register<"MAC1USTCR", 0xdc, 32, [
      Field<"TIC_1US_CNTR", 0, 12, ReadWrite, "TIC_1US_CNTR">,
    ], "1-microsecond-tick counter register">,
    Register<"MACVR", 0x110, 32, [
      Field<"SNPSVER", 0, 8, ReadWrite, "SNPSVER">,
      Field<"USERVER", 8, 8, ReadWrite, "USERVER">,
    ], "Version register">,
    Register<"MACDR", 0x114, 32, [
      Field<"RPESTS", 0, 1, ReadWrite, "RPESTS">,
      Field<"RFCFCSTS", 1, 2, ReadWrite, "RFCFCSTS">,
      Field<"TPESTS", 16, 1, ReadWrite, "TPESTS">,
      Field<"TFCSTS", 17, 2, ReadWrite, "TFCSTS">,
    ], "Debug register">,
    Register<"MACHWF1R", 0x120, 32, [
      Field<"RXFIFOSIZE", 0, 5, ReadWrite, "RXFIFOSIZE">,
      Field<"TXFIFOSIZE", 6, 5, ReadWrite, "TXFIFOSIZE">,
      Field<"OSTEN", 11, 1, ReadWrite, "OSTEN">,
      Field<"PTOEN", 12, 1, ReadWrite, "PTOEN">,
      Field<"ADVTHWORD", 13, 1, ReadWrite, "ADVTHWORD">,
      Field<"ADDR64", 14, 2, ReadWrite, "ADDR64">,
      Field<"DCBEN", 16, 1, ReadWrite, "DCBEN">,
      Field<"SPHEN", 17, 1, ReadWrite, "SPHEN">,
      Field<"TSOEN", 18, 1, ReadWrite, "TSOEN">,
      Field<"DBGMEMA", 19, 1, ReadWrite, "DBGMEMA">,
      Field<"AVSEL", 20, 1, ReadWrite, "AVSEL">,
      Field<"HASHTBLSZ", 24, 2, ReadWrite, "HASHTBLSZ">,
      Field<"L3L4FNUM", 27, 4, ReadWrite, "L3L4FNUM">,
    ], "HW feature 1 register">,
    Register<"MACHWF2R", 0x124, 32, [
      Field<"RXQCNT", 0, 4, ReadWrite, "RXQCNT">,
      Field<"TXQCNT", 6, 4, ReadWrite, "TXQCNT">,
      Field<"RXCHCNT", 12, 4, ReadWrite, "RXCHCNT">,
      Field<"TXCHCNT", 18, 4, ReadWrite, "TXCHCNT">,
      Field<"PPSOUTNUM", 24, 3, ReadWrite, "PPSOUTNUM">,
      Field<"AUXSNAPNUM", 28, 3, ReadWrite, "AUXSNAPNUM">,
    ], "HW feature 2 register">,
    Register<"MACMDIOAR", 0x200, 32, [
      Field<"MB", 0, 1, ReadWrite, "MB">,
      Field<"C45E", 1, 1, ReadWrite, "C45E">,
      Field<"GOC", 2, 2, ReadWrite, "GOC">,
      Field<"SKAP", 4, 1, ReadWrite, "SKAP">,
      Field<"CR", 8, 4, ReadWrite, "CR">,
      Field<"NTC", 12, 3, ReadWrite, "NTC">,
      Field<"RDA", 16, 5, ReadWrite, "RDA">,
      Field<"PA", 21, 5, ReadWrite, "PA">,
      Field<"BTB", 26, 1, ReadWrite, "BTB">,
      Field<"PSE", 27, 1, ReadWrite, "PSE">,
    ], "MDIO address register">,
    Register<"MACMDIODR", 0x204, 32, [
      Field<"MD", 0, 16, ReadWrite, "MD">,
      Field<"RA", 16, 16, ReadWrite, "RA">,
    ], "MDIO data register">,
    Register<"MACARPAR", 0xae0, 32, [
      Field<"ARPPA", 0, 32, ReadWrite, "ARPPA">,
    ], "ARP address register">,
    Register<"MACA0HR", 0x300, 32, [
      Field<"ADDRHI", 0, 16, ReadWrite, "ADDRHI">,
      Field<"AE", 31, 1, Read, "AE">,
    ], "Address 0 high register">,
    Register<"MACA0LR", 0x304, 32, [
      Field<"ADDRLO", 0, 32, ReadWrite, "ADDRLO">,
    ], "Address 0 low register">,
    Register<"MACA1LR", 0x30c, 32, [
      Field<"ADDRLO", 0, 32, ReadWrite, "ADDRLO">,
    ], "Address 1 low register">,
    Register<"MACA2LR", 0x314, 32, [
      Field<"ADDRLO", 0, 32, ReadWrite, "ADDRLO">,
    ], "Address 2 low register">,
    Register<"MACA1HR", 0x308, 32, [
      Field<"ADDRHI", 0, 16, ReadWrite, "ADDRHI">,
      Field<"MBC", 24, 6, ReadWrite, "MBC">,
      Field<"SA", 30, 1, ReadWrite, "SA">,
      Field<"AE", 31, 1, ReadWrite, "AE">,
    ], "Address 1 high register">,
    Register<"MACA2HR", 0x310, 32, [
      Field<"ADDRHI", 0, 16, ReadWrite, "ADDRHI">,
      Field<"MBC", 24, 6, ReadWrite, "MBC">,
      Field<"SA", 30, 1, ReadWrite, "SA">,
      Field<"AE", 31, 1, ReadWrite, "AE">,
    ], "Address 2 high register">,
    Register<"MACA3HR", 0x318, 32, [
      Field<"ADDRHI", 0, 16, ReadWrite, "ADDRHI">,
      Field<"MBC", 24, 6, ReadWrite, "MBC">,
      Field<"SA", 30, 1, ReadWrite, "SA">,
      Field<"AE", 31, 1, ReadWrite, "AE">,
    ], "Address 3 high register">,
    Register<"MACA3LR", 0x31c, 32, [
      Field<"ADDRLO", 0, 32, ReadWrite, "ADDRLO">,
    ], "Address 3 low register">,
    Register<"MMC_CONTROL", 0x700, 32, [
      Field<"CNTRST", 0, 1, ReadWrite, "CNTRST">,
      Field<"CNTSTOPRO", 1, 1, ReadWrite, "CNTSTOPRO">,
      Field<"RSTONRD", 2, 1, ReadWrite, "RSTONRD">,
      Field<"CNTFREEZ", 3, 1, ReadWrite, "CNTFREEZ">,
      Field<"CNTPRST", 4, 1, ReadWrite, "CNTPRST">,
      Field<"CNTPRSTLVL", 5, 1, ReadWrite, "CNTPRSTLVL">,
      Field<"UCDBC", 8, 1, ReadWrite, "UCDBC">,
    ], "MMC control register">,
    Register<"MMC_RX_INTERRUPT", 0x704, 32, [
      Field<"RXCRCERPIS", 5, 1, ReadWrite, "RXCRCERPIS">,
      Field<"RXALGNERPIS", 6, 1, ReadWrite, "RXALGNERPIS">,
      Field<"RXUCGPIS", 17, 1, ReadWrite, "RXUCGPIS">,
      Field<"RXLPIUSCIS", 26, 1, ReadWrite, "RXLPIUSCIS">,
      Field<"RXLPITRCIS", 27, 1, ReadWrite, "RXLPITRCIS">,
    ], "MMC Rx interrupt register">,
    Register<"MMC_TX_INTERRUPT", 0x708, 32, [
      Field<"TXSCOLGPIS", 14, 1, ReadWrite, "TXSCOLGPIS">,
      Field<"TXMCOLGPIS", 15, 1, ReadWrite, "TXMCOLGPIS">,
      Field<"TXGPKTIS", 21, 1, ReadWrite, "TXGPKTIS">,
      Field<"TXLPIUSCIS", 26, 1, ReadWrite, "TXLPIUSCIS">,
      Field<"TXLPITRCIS", 27, 1, ReadWrite, "TXLPITRCIS">,
    ], "MMC Tx interrupt register">,
    Register<"MMC_RX_INTERRUPT_MASK", 0x70c, 32, [
      Field<"RXCRCERPIM", 5, 1, ReadWrite, "RXCRCERPIM">,
      Field<"RXALGNERPIM", 6, 1, ReadWrite, "RXALGNERPIM">,
      Field<"RXUCGPIM", 17, 1, ReadWrite, "RXUCGPIM">,
      Field<"RXLPIUSCIM", 26, 1, ReadWrite, "RXLPIUSCIM">,
      Field<"RXLPITRCIM", 27, 1, Read, "RXLPITRCIM">,
    ], "MMC Rx interrupt mask register">,
    Register<"MMC_TX_INTERRUPT_MASK", 0x710, 32, [
      Field<"TXSCOLGPIM", 14, 1, ReadWrite, "TXSCOLGPIM">,
      Field<"TXMCOLGPIM", 15, 1, ReadWrite, "TXMCOLGPIM">,
      Field<"TXGPKTIM", 21, 1, ReadWrite, "TXGPKTIM">,
      Field<"TXLPIUSCIM", 26, 1, ReadWrite, "TXLPIUSCIM">,
      Field<"TXLPITRCIM", 27, 1, Read, "TXLPITRCIM">,
    ], "MMC Tx interrupt mask register">,
    Register<"TX_SINGLE_COLLISION_GOOD_PACKETS", 0x74c, 32, [
      Field<"TXSNGLCOLG", 0, 32, ReadWrite, "TXSNGLCOLG">,
    ], "Tx single collision good packets register">,
    Register<"TX_MULTIPLE_COLLISION_GOOD_PACKETS", 0x750, 32, [
      Field<"TXMULTCOLG", 0, 32, ReadWrite, "TXMULTCOLG">,
    ], "Tx multiple collision good packets register">,
    Register<"TX_PACKET_COUNT_GOOD", 0x768, 32, [
      Field<"TXPKTG", 0, 32, ReadWrite, "TXPKTG">,
    ], "Tx packet count good register">,
    Register<"RX_CRC_ERROR_PACKETS", 0x794, 32, [
      Field<"RXCRCERR", 0, 32, ReadWrite, "RXCRCERR">,
    ], "Rx CRC error packets register">,
    Register<"RX_ALIGNMENT_ERROR_PACKETS", 0x798, 32, [
      Field<"RXALGNERR", 0, 32, ReadWrite, "RXALGNERR">,
    ], "Rx alignment error packets register">,
    Register<"RX_UNICAST_PACKETS_GOOD", 0x7c4, 32, [
      Field<"RXUCASTG", 0, 32, ReadWrite, "RXUCASTG">,
    ], "Rx unicast packets good register">,
    Register<"TX_LPI_USEC_CNTR", 0x7ec, 32, [
      Field<"TXLPIUSC", 0, 32, ReadWrite, "TXLPIUSC">,
    ], "Tx LPI microsecond timer register">,
    Register<"TX_LPI_TRAN_CNTR", 0x7f0, 32, [
      Field<"TXLPITRC", 0, 32, ReadWrite, "TXLPITRC">,
    ], "Tx LPI transition counter register">,
    Register<"RX_LPI_USEC_CNTR", 0x7f4, 32, [
      Field<"RXLPIUSC", 0, 32, ReadWrite, "RXLPIUSC">,
    ], "Rx LPI microsecond counter register">,
    Register<"RX_LPI_TRAN_CNTR", 0x7f8, 32, [
      Field<"RXLPITRC", 0, 32, ReadWrite, "RXLPITRC">,
    ], "Rx LPI transition counter register">,
    Register<"MACL3L4C0R", 0x900, 32, [
      Field<"L3PEN0", 0, 1, ReadWrite, "L3PEN0">,
      Field<"L3SAM0", 2, 1, ReadWrite, "L3SAM0">,
      Field<"L3SAIM0", 3, 1, ReadWrite, "L3SAIM0">,
      Field<"L3DAM0", 4, 1, ReadWrite, "L3DAM0">,
      Field<"L3DAIM0", 5, 1, ReadWrite, "L3DAIM0">,
      Field<"L3HSBM0", 6, 5, ReadWrite, "L3HSBM0">,
      Field<"L3HDBM0", 11, 5, ReadWrite, "L3HDBM0">,
      Field<"L4PEN0", 16, 1, ReadWrite, "L4PEN0">,
      Field<"L4SPM0", 18, 1, ReadWrite, "L4SPM0">,
      Field<"L4SPIM0", 19, 1, ReadWrite, "L4SPIM0">,
      Field<"L4DPM0", 20, 1, ReadWrite, "L4DPM0">,
      Field<"L4DPIM0", 21, 1, ReadWrite, "L4DPIM0">,
    ], "L3 and L4 control 0 register">,
    Register<"MACL4A0R", 0x904, 32, [
      Field<"L4SP0", 0, 16, ReadWrite, "L4SP0">,
      Field<"L4DP0", 16, 16, ReadWrite, "L4DP0">,
    ], "Layer4 address filter 0 register">,
    Register<"MACL3A00R", 0x910, 32, [
      Field<"L3A00", 0, 32, ReadWrite, "L3A00">,
    ], "MACL3A00R">,
    Register<"MACL3A10R", 0x914, 32, [
      Field<"L3A10", 0, 32, ReadWrite, "L3A10">,
    ], "Layer3 address 1 filter 0 register">,
    Register<"MACL3A20", 0x918, 32, [
      Field<"L3A20", 0, 32, ReadWrite, "L3A20">,
    ], "Layer3 Address 2 filter 0 register">,
    Register<"MACL3A30", 0x91c, 32, [
      Field<"L3A30", 0, 32, ReadWrite, "L3A30">,
    ], "Layer3 Address 3 filter 0 register">,
    Register<"MACL3L4C1R", 0x930, 32, [
      Field<"L3PEN1", 0, 1, ReadWrite, "L3PEN1">,
      Field<"L3SAM1", 2, 1, ReadWrite, "L3SAM1">,
      Field<"L3SAIM1", 3, 1, ReadWrite, "L3SAIM1">,
      Field<"L3DAM1", 4, 1, ReadWrite, "L3DAM1">,
      Field<"L3DAIM1", 5, 1, ReadWrite, "L3DAIM1">,
      Field<"L3HSBM1", 6, 5, ReadWrite, "L3HSBM1">,
      Field<"L3HDBM1", 11, 5, ReadWrite, "L3HDBM1">,
      Field<"L4PEN1", 16, 1, ReadWrite, "L4PEN1">,
      Field<"L4SPM1", 18, 1, ReadWrite, "L4SPM1">,
      Field<"L4SPIM1", 19, 1, ReadWrite, "L4SPIM1">,
      Field<"L4DPM1", 20, 1, ReadWrite, "L4DPM1">,
      Field<"L4DPIM1", 21, 1, ReadWrite, "L4DPIM1">,
    ], "L3 and L4 control 1 register">,
    Register<"MACL4A1R", 0x934, 32, [
      Field<"L4SP1", 0, 16, ReadWrite, "L4SP1">,
      Field<"L4DP1", 16, 16, ReadWrite, "L4DP1">,
    ], "Layer 4 address filter 1 register">,
    Register<"MACL3A01R", 0x940, 32, [
      Field<"L3A01", 0, 32, ReadWrite, "L3A01">,
    ], "Layer3 address 0 filter 1 Register">,
    Register<"MACL3A11R", 0x944, 32, [
      Field<"L3A11", 0, 32, ReadWrite, "L3A11">,
    ], "Layer3 address 1 filter 1 register">,
    Register<"MACL3A21R", 0x948, 32, [
      Field<"L3A21", 0, 32, ReadWrite, "L3A21">,
    ], "Layer3 address 2 filter 1 Register">,
    Register<"MACL3A31R", 0x94c, 32, [
      Field<"L3A31", 0, 32, ReadWrite, "L3A31">,
    ], "Layer3 address 3 filter 1 register">,
    Register<"MACTSCR", 0xb00, 32, [
      Field<"TSENA", 0, 1, ReadWrite, "TSENA">,
      Field<"TSCFUPDT", 1, 1, ReadWrite, "TSCFUPDT">,
      Field<"TSINIT", 2, 1, ReadWrite, "TSINIT">,
      Field<"TSUPDT", 3, 1, ReadWrite, "TSUPDT">,
      Field<"TSADDREG", 5, 1, ReadWrite, "TSADDREG">,
      Field<"TSENALL", 8, 1, ReadWrite, "TSENALL">,
      Field<"TSCTRLSSR", 9, 1, ReadWrite, "TSCTRLSSR">,
      Field<"TSVER2ENA", 10, 1, ReadWrite, "TSVER2ENA">,
      Field<"TSIPENA", 11, 1, ReadWrite, "TSIPENA">,
      Field<"TSIPV6ENA", 12, 1, ReadWrite, "TSIPV6ENA">,
      Field<"TSIPV4ENA", 13, 1, ReadWrite, "TSIPV4ENA">,
      Field<"TSEVNTENA", 14, 1, ReadWrite, "TSEVNTENA">,
      Field<"TSMSTRENA", 15, 1, ReadWrite, "TSMSTRENA">,
      Field<"SNAPTYPSEL", 16, 2, ReadWrite, "SNAPTYPSEL">,
      Field<"TSENMACADDR", 18, 1, ReadWrite, "TSENMACADDR">,
      Field<"CSC", 19, 1, Read, "CSC">,
      Field<"TXTSSTSM", 24, 1, ReadWrite, "TXTSSTSM">,
    ], "Timestamp control Register">,
    Register<"MACSSIR", 0xb04, 32, [
      Field<"SNSINC", 8, 8, ReadWrite, "SNSINC">,
      Field<"SSINC", 16, 8, ReadWrite, "SSINC">,
    ], "Sub-second increment register">,
    Register<"MACSTSR", 0xb08, 32, [
      Field<"TSS", 0, 32, ReadWrite, "TSS">,
    ], "System time seconds register">,
    Register<"MACSTNR", 0xb0c, 32, [
      Field<"TSSS", 0, 31, ReadWrite, "TSSS">,
    ], "System time nanoseconds register">,
    Register<"MACSTSUR", 0xb10, 32, [
      Field<"TSS", 0, 32, ReadWrite, "TSS">,
    ], "System time seconds update register">,
    Register<"MACSTNUR", 0xb14, 32, [
      Field<"TSSS", 0, 31, ReadWrite, "TSSS">,
      Field<"ADDSUB", 31, 1, ReadWrite, "ADDSUB">,
    ], "System time nanoseconds update register">,
    Register<"MACTSAR", 0xb18, 32, [
      Field<"TSAR", 0, 32, ReadWrite, "TSAR">,
    ], "Timestamp addend register">,
    Register<"MACTSSR", 0xb20, 32, [
      Field<"TSSOVF", 0, 1, ReadWrite, "TSSOVF">,
      Field<"TSTARGT0", 1, 1, ReadWrite, "TSTARGT0">,
      Field<"AUXTSTRIG", 2, 1, ReadWrite, "AUXTSTRIG">,
      Field<"TSTRGTERR0", 3, 1, ReadWrite, "TSTRGTERR0">,
      Field<"TXTSSIS", 15, 1, ReadWrite, "TXTSSIS">,
      Field<"ATSSTN", 16, 4, ReadWrite, "ATSSTN">,
      Field<"ATSSTM", 24, 1, ReadWrite, "ATSSTM">,
      Field<"ATSNS", 25, 5, ReadWrite, "ATSNS">,
    ], "Timestamp status register">,
    Register<"MACTxTSSNR", 0xb30, 32, [
      Field<"TXTSSLO", 0, 31, ReadWrite, "TXTSSLO">,
      Field<"TXTSSMIS", 31, 1, ReadWrite, "TXTSSMIS">,
    ], "Tx timestamp status nanoseconds register">,
    Register<"MACTxTSSSR", 0xb34, 32, [
      Field<"TXTSSHI", 0, 32, ReadWrite, "TXTSSHI">,
    ], "Tx timestamp status seconds register">,
    Register<"MACACR", 0xb40, 32, [
      Field<"ATSFC", 0, 1, ReadWrite, "ATSFC">,
      Field<"ATSEN0", 4, 1, ReadWrite, "ATSEN0">,
      Field<"ATSEN1", 5, 1, ReadWrite, "ATSEN1">,
      Field<"ATSEN2", 6, 1, ReadWrite, "ATSEN2">,
      Field<"ATSEN3", 7, 1, ReadWrite, "ATSEN3">,
    ], "Auxiliary control register">,
    Register<"MACATSNR", 0xb48, 32, [
      Field<"AUXTSLO", 0, 31, ReadWrite, "AUXTSLO">,
    ], "Auxiliary timestamp nanoseconds register">,
    Register<"MACATSSR", 0xb4c, 32, [
      Field<"AUXTSHI", 0, 32, ReadWrite, "AUXTSHI">,
    ], "Auxiliary timestamp seconds register">,
    Register<"MACTSIACR", 0xb50, 32, [
      Field<"OSTIAC", 0, 32, ReadWrite, "OSTIAC">,
    ], "Timestamp Ingress asymmetric correction register">,
    Register<"MACTSEACR", 0xb54, 32, [
      Field<"OSTEAC", 0, 32, ReadWrite, "OSTEAC">,
    ], "Timestamp Egress asymmetric correction register">,
    Register<"MACTSICNR", 0xb58, 32, [
      Field<"TSIC", 0, 32, ReadWrite, "TSIC">,
    ], "Timestamp Ingress correction nanosecond register">,
    Register<"MACTSECNR", 0xb5c, 32, [
      Field<"TSEC", 0, 32, ReadWrite, "TSEC">,
    ], "Timestamp Egress correction nanosecond register">,
    Register<"MACPPSCR", 0xb70, 32, [
      Field<"PPSCTRL", 0, 4, ReadWrite, "PPSCTRL">,
      Field<"PPSEN0", 4, 1, ReadWrite, "PPSEN0">,
      Field<"TRGTMODSEL0", 5, 2, ReadWrite, "TRGTMODSEL0">,
    ], "PPS control register">,
    Register<"MACPPSTTSR", 0xb80, 32, [
      Field<"TSTRH0", 0, 31, ReadWrite, "TSTRH0">,
    ], "PPS target time seconds register">,
    Register<"MACPPSTTNR", 0xb84, 32, [
      Field<"TTSL0", 0, 31, ReadWrite, "TTSL0">,
      Field<"TRGTBUSY0", 31, 1, ReadWrite, "TRGTBUSY0">,
    ], "PPS target time nanoseconds register">,
    Register<"MACPPSIR", 0xb88, 32, [
      Field<"PPSINT0", 0, 32, ReadWrite, "PPSINT0">,
    ], "PPS interval register">,
    Register<"MACPPSWR", 0xb8c, 32, [
      Field<"PPSWIDTH0", 0, 32, ReadWrite, "PPSWIDTH0">,
    ], "PPS width register">,
    Register<"MACPOCR", 0xbc0, 32, [
      Field<"PTOEN", 0, 1, ReadWrite, "PTOEN">,
      Field<"ASYNCEN", 1, 1, ReadWrite, "ASYNCEN">,
      Field<"APDREQEN", 2, 1, ReadWrite, "APDREQEN">,
      Field<"ASYNCTRIG", 4, 1, ReadWrite, "ASYNCTRIG">,
      Field<"APDREQTRIG", 5, 1, ReadWrite, "APDREQTRIG">,
      Field<"DRRDIS", 6, 1, ReadWrite, "DRRDIS">,
      Field<"DN", 8, 8, ReadWrite, "DN">,
    ], "PTP Offload control register">,
    Register<"MACSPI0R", 0xbc4, 32, [
      Field<"SPI0", 0, 32, ReadWrite, "SPI0">,
    ], "PTP Source Port Identity 0 Register">,
    Register<"MACSPI1R", 0xbc8, 32, [
      Field<"SPI1", 0, 32, ReadWrite, "SPI1">,
    ], "PTP Source port identity 1 register">,
    Register<"MACSPI2R", 0xbcc, 32, [
      Field<"SPI2", 0, 16, ReadWrite, "SPI2">,
    ], "PTP Source port identity 2 register">,
    Register<"MACLMIR", 0xbd0, 32, [
      Field<"LSI", 0, 8, ReadWrite, "LSI">,
      Field<"DRSYNCR", 8, 3, ReadWrite, "DRSYNCR">,
      Field<"LMPDRI", 24, 8, ReadWrite, "LMPDRI">,
    ], "Log message interval register">,
  ];
}

class Ethernet_MACInstance<string Name, int Base> : PeripheralInstance<Name, Base, Ethernet_MACPeripheral>;
def Ethernet_MAC : Ethernet_MACInstance<"Ethernet_MAC", 0x40028000>;

#endif // _PERIPHERALS_ETHERNET_MAC_TD
