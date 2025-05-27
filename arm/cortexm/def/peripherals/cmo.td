#ifndef _PERIPHERALS_CMO_TD
#define _PERIPHERALS_CMO_TD

include "base.td"

def CMOPeripheral : PeripheralType<"CMO", "Cache Maintenance Operations"> {
  let accessWidth = 32;
  let registers = [
    Register<"ICIALLU", 0x0, 32, [
      Field<"IGNORED", 0, 32, ReadWrite, "The value written to this field is ignored">,
    ], "Invalidates all instruction caches to Point of Unification (PoU)">,
    Register<"ICIMVAU", 0x8, 32, [
      Field<"ADDRESS", 0, 32, ReadWrite, "Writing to this field initiates the maintenance operation for the address that is written">,
    ], "Invalidates instruction cache lines by address to Point of Unification (PoU)">,
    Register<"DCIMVAC", 0xc, 32, [
      Field<"ADDRESS", 0, 32, ReadWrite, "Writing to this field initiates the maintenance operation for the address that is written">,
    ], "Invalidates data or unified cache lines by address to Point of Coherency (PoC)">,
    Register<"DCISW", 0x10, 32, [
      Field<"LEVEL", 1, 3, ReadWrite, "Cache level. This field is 0b000.">,
      Field<"SETWAY", 4, 28, ReadWrite, "Cache set/way. Contains two fields: Way, bits[31:32-A], the number of the way to operate on. Set, bits[B-1:L], the number of the set to operate on.">,
    ], "Invalidates instruction cache lines by address to Point of Unification (PoU)">,
    Register<"DCCMVAU", 0x14, 32, [
      Field<"ADDRESS", 0, 32, ReadWrite, "Writing to this field initiates the maintenance operation for the address that is written">,
    ], "Cleans data or unified cache lines by address to Point of Unification (PoU)">,
    Register<"DCCMVAC", 0x18, 32, [
      Field<"ADDRESS", 0, 32, ReadWrite, "Writing to this field initiates the maintenance operation for the address that is written">,
    ], "Cleans data or unified cache lines by address to Point of Unification (PoC)">,
    Register<"DCCSW", 0x1c, 32, [
      Field<"LEVEL", 1, 3, ReadWrite, "Cache level. This field is 0b000">,
      Field<"SETWAY", 4, 28, ReadWrite, "Contains two fields: Way, bits [31:32-A], the number of the way to operate on. Set, bits [B-1:L], the number of the set to operate on.">,
    ], "Cleans data or unified cache line by set/way">,
    Register<"DCCIMVAC", 0x20, 32, [
      Field<"ADDRESS", 0, 32, ReadWrite, "Writing to this field initiates the maintenance operation for the address that is written">,
    ], "Cleans data or unified cache lines by address to Point of Coherency (PoC)">,
    Register<"DCCISW", 0x24, 32, [
      Field<"LEVEL", 1, 3, ReadWrite, "Cache level. This field is 0b000.">,
      Field<"SETWAY", 4, 28, ReadWrite, "Contains two fields: Way, bits[31:32-A], the number of the way to operate on. Set, bits[B-1:L], the number of the set to operate on.">,
    ], "Cleans and invalidates data or unified cache line by set/way">,
    Register<"BPIALL", 0x28, 32, [
      Field<"IGNORED", 0, 32, ReadWrite, "The value written to this field is ignored">,
    ], "Invalidates all entries from branch predictors">,
  ];
}

class CMOInstance<string Name, int Base> : PeripheralInstance<Name, Base, CMOPeripheral>;
def CMO : CMOInstance<"CMO", 0xe000ef50>;

#endif // _PERIPHERALS_CMO_TD
