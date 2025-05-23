#ifndef _PERIPHERALS_OTG1_HS_PWRCLK_TD
#define _PERIPHERALS_OTG1_HS_PWRCLK_TD

include "base.td"

def OTG1_HS_PWRCLKPeripheral : PeripheralType<"OTG1_HS_PWRCLK", "USB 1 on the go high speed"> {
  let accessWidth = 32;
  let registers = [
    Register<"OTG_HS_PCGCR", 0x0, 32, [
      Field<"STPPCLK", 0, 1, ReadWrite, "Stop PHY clock">,
      Field<"GATEHCLK", 1, 1, ReadWrite, "Gate HCLK">,
      Field<"PHYSUSP", 4, 1, ReadWrite, "PHY suspended">,
    ], "Power and clock gating control register">,
  ];
}

class OTG1_HS_PWRCLKInstance<string Name, int Base> : PeripheralInstance<Name, Base, OTG1_HS_PWRCLKPeripheral>;
def OTG1_HS_PWRCLK : OTG1_HS_PWRCLKInstance<"OTG1_HS_PWRCLK", 0x40040e00>;
def OTG2_HS_PWRCLK : OTG1_HS_PWRCLKInstance<"OTG2_HS_PWRCLK", 0x40080e00>;

#endif // _PERIPHERALS_OTG1_HS_PWRCLK_TD
