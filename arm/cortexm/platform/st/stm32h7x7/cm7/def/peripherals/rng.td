#ifndef _PERIPHERALS_RNG_TD
#define _PERIPHERALS_RNG_TD

include "base.td"

def RNGPeripheral : PeripheralType<"RNG", "RNG"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR", 0x0, 32, [
      Field<"RNGEN", 2, 1, ReadWrite, "Random number generator enable">,
      Field<"IE", 3, 1, ReadWrite, "Interrupt enable">,
      Field<"CED", 5, 1, ReadWrite, "Clock error detection Note: The clock error detection can be used only when ck_rc48 or ck_pll1_q (ck_pll1_q = 48MHz) source is selected otherwise, CED bit must be equal to 1. The clock error detection cannot be enabled nor disabled on the fly when RNG peripheral is enabled, to enable or disable CED the RNG must be disabled.">,
    ], "RNG control register">,
    Register<"SR", 0x4, 32, [
      Field<"DRDY", 0, 1, Read, "Data ready Note: If IE=1 in RNG_CR, an interrupt is generated when DRDY=1. It can rise when the peripheral is disabled. When the output buffer becomes empty (after reading RNG_DR), this bit returns to 0 until a new random value is generated.">,
      Field<"CECS", 1, 1, Read, "Clock error current status Note: This bit is meaningless if CED (Clock error detection) bit in RNG_CR is equal to 1.">,
      Field<"SECS", 2, 1, Read, "Seed error current status ** More than 64 consecutive bits at the same value (0 or 1) ** More than 32 consecutive alternances of 0 and 1 (0101010101...01)">,
      Field<"CEIS", 5, 1, ReadWrite, "Clock error interrupt status This bit is set at the same time as CECS. It is cleared by writing it to 0. An interrupt is pending if IE = 1 in the RNG_CR register. Note: This bit is meaningless if CED (Clock error detection) bit in RNG_CR is equal to 1.">,
      Field<"SEIS", 6, 1, ReadWrite, "Seed error interrupt status This bit is set at the same time as SECS. It is cleared by writing it to 0. ** More than 64 consecutive bits at the same value (0 or 1) ** More than 32 consecutive alternances of 0 and 1 (0101010101...01) An interrupt is pending if IE = 1 in the RNG_CR register.">,
    ], "RNG status register">,
    Register<"DR", 0x8, 32, [
      Field<"RNDATA", 0, 32, ReadWrite, "Random data 32-bit random data which are valid when DRDY=1.">,
    ], "The RNG_DR register is a read-only register that delivers a 32-bit random value when read. The content of this register is valid when DRDY= 1, even if RNGEN=0.">,
  ];
}

class RNGInstance<string Name, int Base> : PeripheralInstance<Name, Base, RNGPeripheral>;
def RNG : RNGInstance<"RNG", 0x48021800>;

#endif // _PERIPHERALS_RNG_TD
