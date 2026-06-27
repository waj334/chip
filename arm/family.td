#ifndef _ARM_FAMILY_TD
#define _ARM_FAMILY_TD

include "base.td"

// FPU definitions.
def VFP2 : FPU<"vfpv2", [ "vfp2" ]>;
def VFPV3 : FPU<"vfpv3", [ "vfp3" ]>;
def VFPV3_D16 : FPU<"vfpv3-d16", [ "vfp3d16" ]>;
def VFPV3_SP_D16 : FPU<"vfpv3-d16", [ "vfp3d16sp" ]>;
def VFPV4 : FPU<"vfpv4", [ "vfp4" ]>;
def VFPV4_D16 : FPU<"vfpv4-d16", [ "vfp4d16" ]>;
def FPV4_SP_D16 : FPU<"fpv4-sp-d16", [ "vfp4d16sp" ]>;
def FPV4_DP_D16 : FPU<"vfpv4-d16", [ "vfp4d16" ]>;
def FPV5 : FPU<"vfpv5", [ "fp-armv8" ]>;
def FPV5_SP_D16 : FPU<"fpv5-sp-d16", [ "fp-armv8d16sp" ]>;
def FPV5_D16 : FPU<"fpv5-d16", [ "fp-armv8d16" ]>;
def NEON : FPU<"neon", [ "neon" ]>;

// Define supported ARM family.
def ARM : Family<"arm">;
class ARMArchitecture<
    string Name,
    string Arch,
    string Triple,
    int PointerSize = 4,
    int PointerAlignment = 4,
    int StackAlignment = 8,
    int MinHeapAlignment = 8,
    FPU Fpu = NoFPU> : Architecture<Name, ARM, Arch, Triple,
      PointerSize, PointerAlignment, StackAlignment, MinHeapAlignment, Fpu>;

class CortexMArchitecture<string Name, string Arch, string Triple, list<string> Tags, FPU Fpu = NoFPU>
        : ARMArchitecture<Name, Arch, Triple, 4, 4, 8, 8, Fpu> {
    let tags = !listconcat([ "arm", "cortexm" ], Tags);
    let runtimePackages = [ "pkg.si-go.dev/chip/arm/cortexm" ];
}

def CortexM0 : CortexMArchitecture<"cortexm-m0", "arm", "armv6m-none-eabi", [ "cortexm0","armv6m" ]>;
def CortexM0Plus : CortexMArchitecture<"cortexm-m0plus", "arm", "armv6m-none-eabi", [ "cortexm0", "cortexm0plus", "armv6m" ]>;
def CortexM1 : CortexMArchitecture<"cortexm-m1", "arm", "armv6m-none-eabi", [ "cortexm1", "armv6m" ]>;
def CortexM23 : CortexMArchitecture<"cortex-m23", "arm", "armv8m.base-none-eabihf", [ "cortexm23", "armv8m.base" ]>;
def CortexM3 : CortexMArchitecture<"cortex-m3", "thumb2", "armv7m-none-eabi", [ "cortexm3", "armv7m", "thumb" ]>;
def CortexM33 : CortexMArchitecture<"cortex-m33", "thumb2", "armv8m.main-none-eabihf", [ "cortexm33", "armv8m.main", "thumb" ], FPV5_SP_D16>;
def CortexM35P : CortexMArchitecture<"cortex-m35p", "thumb2", "armv8m.main-none-eabihf", [ "cortexm35p", "armv8m.main", "thumb" ], FPV5_SP_D16>;
def CortexM4 : CortexMArchitecture<"cortex-m4", "thumb2", "armv7em-none-eabihf", [ "cortexm4", "armv7m", "armv7em", "thumb" ], FPV4_SP_D16>;
def CortexM52 : CortexMArchitecture<"cortex-m52", "thumb2", "armv81m.main-none-eabihf", [ "cortexm52", "armv81m.main", "thumb" ], FPV5_D16>;
def CortexM55 : CortexMArchitecture<"cortex-m55", "thumb2", "armv81m.main-none-eabihf", [ "cortexm55", "armv81m.main", "thumb" ], FPV5_D16>;
def CortexM7 : CortexMArchitecture<"cortex-m7", "thumb2", "armv7em-none-eabihf", [ "cortexm7", "armv7m", "armv7em", "thumb" ], FPV5_D16>;
def CortexM85 : CortexMArchitecture<"cortex-m85", "thumb2", "armv81m.main-none-eabihf", [ "cortexm85", "armv81m.main", "thumb" ], FPV5_D16>;

#endif // _ARM_FAMILY_TD
