#ifndef _PERIPHERALS_OPAMP_TD
#define _PERIPHERALS_OPAMP_TD

include "base.td"

def OPAMPPeripheral : PeripheralType<"OPAMP", "Operational amplifiers"> {
  let accessWidth = 32;
  let registers = [
    Register<"OPAMP1_CSR", 0x0, 32, [
      Field<"OPAEN", 0, 1, ReadWrite, "Operational amplifier Enable">,
      Field<"FORCE_VP", 1, 1, ReadWrite, "Force internal reference on VP (reserved for test">,
      Field<"VP_SEL", 2, 2, ReadWrite, "Operational amplifier PGA mode">,
      Field<"VM_SEL", 5, 2, ReadWrite, "Inverting input selection">,
      Field<"OPAHSM", 8, 1, ReadWrite, "Operational amplifier high-speed mode">,
      Field<"CALON", 11, 1, ReadWrite, "Calibration mode enabled">,
      Field<"CALSEL", 12, 2, ReadWrite, "Calibration selection">,
      Field<"PGA_GAIN", 14, 4, ReadWrite, "allows to switch from AOP offset trimmed values to AOP offset">,
      Field<"USERTRIM", 18, 1, ReadWrite, "User trimming enable">,
      Field<"TSTREF", 29, 1, ReadWrite, "OPAMP calibration reference voltage output control (reserved for test)">,
      Field<"CALOUT", 30, 1, ReadWrite, "Operational amplifier calibration output">,
    ], "OPAMP1 control/status register">,
    Register<"OPAMP1_OTR", 0x4, 32, [
      Field<"TRIMOFFSETN", 0, 5, ReadWrite, "Trim for NMOS differential pairs">,
      Field<"TRIMOFFSETP", 8, 5, ReadWrite, "Trim for PMOS differential pairs">,
    ], "OPAMP1 offset trimming register in normal mode">,
    Register<"OPAMP1_HSOTR", 0x8, 32, [
      Field<"TRIMLPOFFSETN", 0, 5, ReadWrite, "Trim for NMOS differential pairs">,
      Field<"TRIMLPOFFSETP", 8, 5, ReadWrite, "Trim for PMOS differential pairs">,
    ], "OPAMP1 offset trimming register in low-power mode">,
    Register<"OPAMP2_CSR", 0x10, 32, [
      Field<"OPAEN", 0, 1, ReadWrite, "Operational amplifier Enable">,
      Field<"FORCE_VP", 1, 1, ReadWrite, "Force internal reference on VP (reserved for test)">,
      Field<"VM_SEL", 5, 2, ReadWrite, "Inverting input selection">,
      Field<"OPAHSM", 8, 1, ReadWrite, "Operational amplifier high-speed mode">,
      Field<"CALON", 11, 1, ReadWrite, "Calibration mode enabled">,
      Field<"CALSEL", 12, 2, ReadWrite, "Calibration selection">,
      Field<"PGA_GAIN", 14, 4, ReadWrite, "Operational amplifier Programmable amplifier gain value">,
      Field<"USERTRIM", 18, 1, ReadWrite, "User trimming enable">,
      Field<"TSTREF", 29, 1, ReadWrite, "OPAMP calibration reference voltage output control (reserved for test)">,
      Field<"CALOUT", 30, 1, ReadWrite, "Operational amplifier calibration output">,
    ], "OPAMP2 control/status register">,
    Register<"OPAMP2_OTR", 0x14, 32, [
      Field<"TRIMOFFSETN", 0, 5, ReadWrite, "Trim for NMOS differential pairs">,
      Field<"TRIMOFFSETP", 8, 5, ReadWrite, "Trim for PMOS differential pairs">,
    ], "OPAMP2 offset trimming register in normal mode">,
    Register<"OPAMP2_HSOTR", 0x18, 32, [
      Field<"TRIMLPOFFSETN", 0, 5, ReadWrite, "Trim for NMOS differential pairs">,
      Field<"TRIMLPOFFSETP", 8, 5, ReadWrite, "Trim for PMOS differential pairs">,
    ], "OPAMP2 offset trimming register in low-power mode">,
  ];
}

class OPAMPInstance<string Name, int Base> : PeripheralInstance<Name, Base, OPAMPPeripheral>;
def OPAMP : OPAMPInstance<"OPAMP", 0x40009000>;

#endif // _PERIPHERALS_OPAMP_TD
