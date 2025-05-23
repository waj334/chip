#ifndef _PERIPHERALS_NVIC_TD
#define _PERIPHERALS_NVIC_TD

include "base.td"

def NVICPeripheral : PeripheralType<"NVIC", "Nested Vectored Interrupt Controller"> {
  let accessWidth = 32;
  let registers = [
    Register<"ISER0", 0x0, 32, [
      Field<"SETENA", 0, 32, ReadWrite, "SETENA">,
    ], "Interrupt Set-Enable Register">,
    Register<"ISER1", 0x4, 32, [
      Field<"SETENA", 0, 32, ReadWrite, "SETENA">,
    ], "Interrupt Set-Enable Register">,
    Register<"ISER2", 0x8, 32, [
      Field<"SETENA", 0, 32, ReadWrite, "SETENA">,
    ], "Interrupt Set-Enable Register">,
    Register<"ICER0", 0x80, 32, [
      Field<"CLRENA", 0, 32, ReadWrite, "CLRENA">,
    ], "Interrupt Clear-Enable Register">,
    Register<"ICER1", 0x84, 32, [
      Field<"CLRENA", 0, 32, ReadWrite, "CLRENA">,
    ], "Interrupt Clear-Enable Register">,
    Register<"ICER2", 0x88, 32, [
      Field<"CLRENA", 0, 32, ReadWrite, "CLRENA">,
    ], "Interrupt Clear-Enable Register">,
    Register<"ISPR0", 0x100, 32, [
      Field<"SETPEND", 0, 32, ReadWrite, "SETPEND">,
    ], "Interrupt Set-Pending Register">,
    Register<"ISPR1", 0x104, 32, [
      Field<"SETPEND", 0, 32, ReadWrite, "SETPEND">,
    ], "Interrupt Set-Pending Register">,
    Register<"ISPR2", 0x108, 32, [
      Field<"SETPEND", 0, 32, ReadWrite, "SETPEND">,
    ], "Interrupt Set-Pending Register">,
    Register<"ICPR0", 0x180, 32, [
      Field<"CLRPEND", 0, 32, ReadWrite, "CLRPEND">,
    ], "Interrupt Clear-Pending Register">,
    Register<"ICPR1", 0x184, 32, [
      Field<"CLRPEND", 0, 32, ReadWrite, "CLRPEND">,
    ], "Interrupt Clear-Pending Register">,
    Register<"ICPR2", 0x188, 32, [
      Field<"CLRPEND", 0, 32, ReadWrite, "CLRPEND">,
    ], "Interrupt Clear-Pending Register">,
    Register<"IABR0", 0x200, 32, [
      Field<"ACTIVE", 0, 32, ReadWrite, "ACTIVE">,
    ], "Interrupt Active Bit Register">,
    Register<"IABR1", 0x204, 32, [
      Field<"ACTIVE", 0, 32, ReadWrite, "ACTIVE">,
    ], "Interrupt Active Bit Register">,
    Register<"IABR2", 0x208, 32, [
      Field<"ACTIVE", 0, 32, ReadWrite, "ACTIVE">,
    ], "Interrupt Active Bit Register">,
    Register<"IPR0", 0x300, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR1", 0x304, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR2", 0x308, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR3", 0x30c, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR4", 0x310, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR5", 0x314, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR6", 0x318, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR7", 0x31c, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR8", 0x320, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR9", 0x324, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR10", 0x328, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR11", 0x32c, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR12", 0x330, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR13", 0x334, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR14", 0x338, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR15", 0x33c, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR16", 0x340, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR17", 0x344, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR18", 0x348, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR19", 0x34c, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR20", 0x350, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR21", 0x354, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR22", 0x358, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR23", 0x35c, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR24", 0x360, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR25", 0x364, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR26", 0x368, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR27", 0x36c, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR28", 0x370, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR29", 0x374, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR30", 0x378, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR31", 0x37c, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR32", 0x380, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR33", 0x384, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR34", 0x388, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR35", 0x38c, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR36", 0x390, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR37", 0x394, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR38", 0x398, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"IPR39", 0x39c, 32, [
      Field<"IPR_N0", 0, 8, ReadWrite, "IPR_N0">,
      Field<"IPR_N1", 8, 8, ReadWrite, "IPR_N1">,
      Field<"IPR_N2", 16, 8, ReadWrite, "IPR_N2">,
      Field<"IPR_N3", 24, 8, ReadWrite, "IPR_N3">,
    ], "Interrupt Priority Register">,
    Register<"ISER3", 0xc, 32, [
    ], "Interrupt Set-Enable Register">,
    Register<"ISER4", 0x10, 32, [
    ], "Interrupt Set-Enable Register">,
    Register<"ICER3", 0x8c, 32, [
    ], "Interrupt Clear-Enable Register">,
    Register<"ICER4", 0x90, 32, [
    ], "Interrupt Clear-Enable Register">,
    Register<"ISPR3", 0x10c, 32, [
    ], "Interrupt Set-Pending Register">,
    Register<"ISPR4", 0x110, 32, [
    ], "Interrupt Set-Pending Register">,
    Register<"ICPR3", 0x1c0, 32, [
    ], "Interrupt Clear-Pending Register">,
    Register<"ICPR4", 0x1c4, 32, [
    ], "Interrupt Clear-Pending Register">,
    Register<"IABR3", 0x20c, 32, [
    ], "Interrupt Active Bit Register">,
    Register<"IABR4", 0x210, 32, [
    ], "Interrupt Active Bit Register">,
  ];
}

class NVICInstance<string Name, int Base> : PeripheralInstance<Name, Base, NVICPeripheral>;
def NVIC : NVICInstance<"NVIC", 0xe000e100>;

#endif // _PERIPHERALS_NVIC_TD
