#ifndef __TD
#define __TD

include "base.td"
include "arm/cortexm/interrupts.td"
include "arm/cortexm/registers.td"
include "arm/cortexm/variant.td"
include "arm/family.td"

include "peripherals/cmo.td"
include "peripherals/dca.td"
include "peripherals/errbnk.td"
include "peripherals/ewic.td"
include "peripherals/ewiccs.td"
include "peripherals/ewicis.td"
include "peripherals/fpu.td"
include "peripherals/icb.td"
include "peripherals/mpu.td"
include "peripherals/msc.td"
include "peripherals/nvic.td"
include "peripherals/pci.td"
include "peripherals/pmc.td"
include "peripherals/pmu.td"
include "peripherals/ras.td"
include "peripherals/sau.td"
include "peripherals/scb.td"
include "peripherals/systick.td"
include "peripherals/tcm.td"

defvar peripherals = [
  CMOPeripheral,
  DCAPeripheral,
  ERRBNKPeripheral,
  EWICPeripheral,
  EWICCSPeripheral,
  EWICISPeripheral,
  FPUPeripheral,
  ICBPeripheral,
  MPUPeripheral,
  MSCPeripheral,
  NVICPeripheral,
  PCIPeripheral,
  PMCPeripheral,
  PMUPeripheral,
  RASPeripheral,
  SAUPeripheral,
  SCBPeripheral,
  SYSTICKPeripheral,
  TCMPeripheral,
];

def CortexMBase : ARMArchitecture<"", "", "", 4> {
    let tags = [ "arm", "cortexm" ];
}

def CortexM : Series<"CortexM", CortexMBase> {
  let peripheralTypes = peripherals;
  let internal = true;
}

#endif // __TD
