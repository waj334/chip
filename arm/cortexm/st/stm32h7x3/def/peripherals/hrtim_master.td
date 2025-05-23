#ifndef _PERIPHERALS_HRTIM_MASTER_TD
#define _PERIPHERALS_HRTIM_MASTER_TD

include "base.td"

def HRTIM_MasterPeripheral : PeripheralType<"HRTIM_Master", "High Resolution Timer: Master Timers"> {
  let accessWidth = 32;
  let registers = [
    Register<"MCR", 0x0, 32, [
      Field<"BRSTDMA", 30, 2, ReadWrite, "Burst DMA Update">,
      Field<"MREPU", 29, 1, ReadWrite, "Master Timer Repetition update">,
      Field<"PREEN", 27, 1, ReadWrite, "Preload enable">,
      Field<"DACSYNC", 25, 2, ReadWrite, "AC Synchronization">,
      Field<"TECEN", 21, 1, ReadWrite, "Timer E counter enable">,
      Field<"TDCEN", 20, 1, ReadWrite, "Timer D counter enable">,
      Field<"TCCEN", 19, 1, ReadWrite, "Timer C counter enable">,
      Field<"TBCEN", 18, 1, ReadWrite, "Timer B counter enable">,
      Field<"TACEN", 17, 1, ReadWrite, "Timer A counter enable">,
      Field<"MCEN", 16, 1, ReadWrite, "Master Counter enable">,
      Field<"SYNC_SRC", 14, 2, ReadWrite, "Synchronization source">,
      Field<"SYNC_OUT", 12, 2, ReadWrite, "Synchronization output">,
      Field<"SYNCSTRTM", 11, 1, ReadWrite, "Synchronization Starts Master">,
      Field<"SYNCRSTM", 10, 1, ReadWrite, "Synchronization Resets Master">,
      Field<"SYNC_IN", 8, 2, ReadWrite, "ynchronization input">,
      Field<"HALF", 5, 1, ReadWrite, "Half mode enable">,
      Field<"RETRIG", 4, 1, ReadWrite, "Master Re-triggerable mode">,
      Field<"CONT", 3, 1, ReadWrite, "Master Continuous mode">,
      Field<"CK_PSC", 0, 3, ReadWrite, "HRTIM Master Clock prescaler">,
    ], "Master Timer Control Register">,
    Register<"MISR", 0x4, 32, [
      Field<"MUPD", 6, 1, ReadWrite, "Master Update Interrupt Flag">,
      Field<"SYNC", 5, 1, ReadWrite, "Sync Input Interrupt Flag">,
      Field<"MREP", 4, 1, ReadWrite, "Master Repetition Interrupt Flag">,
      Field<"MCMP4", 3, 1, ReadWrite, "Master Compare 4 Interrupt Flag">,
      Field<"MCMP3", 2, 1, ReadWrite, "Master Compare 3 Interrupt Flag">,
      Field<"MCMP2", 1, 1, ReadWrite, "Master Compare 2 Interrupt Flag">,
      Field<"MCMP1", 0, 1, ReadWrite, "Master Compare 1 Interrupt Flag">,
    ], "Master Timer Interrupt Status Register">,
    Register<"MICR", 0x8, 32, [
      Field<"MUPDC", 6, 1, ReadWrite, "Master update Interrupt flag clear">,
      Field<"SYNCC", 5, 1, ReadWrite, "Sync Input Interrupt flag clear">,
      Field<"MREPC", 4, 1, ReadWrite, "Repetition Interrupt flag clear">,
      Field<"MCMP4C", 3, 1, ReadWrite, "Master Compare 4 Interrupt flag clear">,
      Field<"MCMP3C", 2, 1, ReadWrite, "Master Compare 3 Interrupt flag clear">,
      Field<"MCMP2C", 1, 1, ReadWrite, "Master Compare 2 Interrupt flag clear">,
      Field<"MCMP1C", 0, 1, ReadWrite, "Master Compare 1 Interrupt flag clear">,
    ], "Master Timer Interrupt Clear Register">,
    Register<"MDIER4", 0xc, 32, [
      Field<"MUPDDE", 22, 1, ReadWrite, "MUPDDE">,
      Field<"SYNCDE", 21, 1, ReadWrite, "SYNCDE">,
      Field<"MREPDE", 20, 1, ReadWrite, "MREPDE">,
      Field<"MCMP4DE", 19, 1, ReadWrite, "MCMP4DE">,
      Field<"MCMP3DE", 18, 1, ReadWrite, "MCMP3DE">,
      Field<"MCMP2DE", 17, 1, ReadWrite, "MCMP2DE">,
      Field<"MCMP1DE", 16, 1, ReadWrite, "MCMP1DE">,
      Field<"MUPDIE", 6, 1, ReadWrite, "MUPDIE">,
      Field<"SYNCIE", 5, 1, ReadWrite, "SYNCIE">,
      Field<"MREPIE", 4, 1, ReadWrite, "MREPIE">,
      Field<"MCMP4IE", 3, 1, ReadWrite, "MCMP4IE">,
      Field<"MCMP3IE", 2, 1, ReadWrite, "MCMP3IE">,
      Field<"MCMP2IE", 1, 1, ReadWrite, "MCMP2IE">,
      Field<"MCMP1IE", 0, 1, ReadWrite, "MCMP1IE">,
    ], "MDIER4">,
    Register<"MCNTR", 0x10, 32, [
      Field<"MCNT", 0, 16, ReadWrite, "Counter value">,
    ], "Master Timer Counter Register">,
    Register<"MPER", 0x14, 32, [
      Field<"MPER", 0, 16, ReadWrite, "Master Timer Period value">,
    ], "Master Timer Period Register">,
    Register<"MREP", 0x18, 32, [
      Field<"MREP", 0, 8, ReadWrite, "Master Timer Repetition counter value">,
    ], "Master Timer Repetition Register">,
    Register<"MCMP1R", 0x1c, 32, [
      Field<"MCMP1", 0, 16, ReadWrite, "Master Timer Compare 1 value">,
    ], "Master Timer Compare 1 Register">,
    Register<"MCMP2R", 0x24, 32, [
      Field<"MCMP2", 0, 16, ReadWrite, "Master Timer Compare 2 value">,
    ], "Master Timer Compare 2 Register">,
    Register<"MCMP3R", 0x28, 32, [
      Field<"MCMP3", 0, 16, ReadWrite, "Master Timer Compare 3 value">,
    ], "Master Timer Compare 3 Register">,
    Register<"MCMP4R", 0x2c, 32, [
      Field<"MCMP4", 0, 16, ReadWrite, "Master Timer Compare 4 value">,
    ], "Master Timer Compare 4 Register">,
  ];
}

class HRTIM_MasterInstance<string Name, int Base> : PeripheralInstance<Name, Base, HRTIM_MasterPeripheral>;
def HRTIM_Master : HRTIM_MasterInstance<"HRTIM_Master", 0x40017400>;

#endif // _PERIPHERALS_HRTIM_MASTER_TD
