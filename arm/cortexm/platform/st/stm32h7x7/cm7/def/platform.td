#ifndef _STM32H7x7_TD
#define _STM32H7x7_TD

include "base.td"
include "arm/cortexm/interrupts.td"
include "arm/cortexm/registers.td"
include "arm/cortexm/variant.td"
include "arm/family.td"

include "peripherals/adc.td"
include "peripherals/adc_common.td"
include "peripherals/art.td"
include "peripherals/axi.td"
include "peripherals/bdma.td"
include "peripherals/can_ccu.td"
include "peripherals/cec.td"
include "peripherals/comp.td"
include "peripherals/crc.td"
include "peripherals/crs.td"
include "peripherals/cryp.td"
include "peripherals/dac.td"
include "peripherals/dbgmcu.td"
include "peripherals/dcmi.td"
include "peripherals/delay_block_sdmmc1.td"
include "peripherals/dfsdm.td"
include "peripherals/dma.td"
include "peripherals/dma2d.td"
include "peripherals/dmamux.td"
include "peripherals/dsihost.td"
include "peripherals/ethernet_mac.td"
include "peripherals/exti.td"
include "peripherals/fdcan.td"
include "peripherals/flash.td"
include "peripherals/fmc.td"
include "peripherals/gpio.td"
include "peripherals/hash.td"
include "peripherals/hrtim_common.td"
include "peripherals/hrtim_master.td"
include "peripherals/hrtim_tim.td"
include "peripherals/hsem.td"
include "peripherals/i2c.td"
include "peripherals/iwdg.td"
include "peripherals/jpeg.td"
include "peripherals/lptim.td"
include "peripherals/lpuart.td"
include "peripherals/ltdc.td"
include "peripherals/mdios.td"
include "peripherals/mdma.td"
include "peripherals/opamp.td"
include "peripherals/otg_hs_device.td"
include "peripherals/otg_hs_global.td"
include "peripherals/otg_hs_host.td"
include "peripherals/otg_hs_pwrclk.td"
include "peripherals/pwr.td"
include "peripherals/quadspi.td"
include "peripherals/ramecc1.td"
include "peripherals/ramecc2.td"
include "peripherals/ramecc3.td"
include "peripherals/rcc.td"
include "peripherals/rng.td"
include "peripherals/rtc.td"
include "peripherals/sai.td"
include "peripherals/sdmmc.td"
include "peripherals/spdifrx.td"
include "peripherals/spi.td"
include "peripherals/swpmi.td"
include "peripherals/syscfg.td"
include "peripherals/tim1_8.td"
include "peripherals/tim2_3_4_5.td"
include "peripherals/tim6_7.td"
include "peripherals/tim12.td"
include "peripherals/tim13_14.td"
include "peripherals/tim15.td"
include "peripherals/tim16_17.td"
include "peripherals/usart.td"
include "peripherals/vrefbuf.td"
include "peripherals/wwdg.td"

defvar peripherals = [
  ADCPeripheral,
  ADC_CommonPeripheral,
  ARTPeripheral,
  AXIPeripheral,
  BDMAPeripheral,
  CAN_CCUPeripheral,
  CECPeripheral,
  COMPPeripheral,
  CRCPeripheral,
  CRSPeripheral,
  CRYPPeripheral,
  DACPeripheral,
  DBGMCUPeripheral,
  DCMIPeripheral,
  DELAY_Block_SDMMC1Peripheral,
  DFSDMPeripheral,
  DMAPeripheral,
  DMA2DPeripheral,
  DMAMUXPeripheral,
  DSIHOSTPeripheral,
  EXTIPeripheral,
  Ethernet_MACPeripheral,
  FDCANPeripheral,
  FMCPeripheral,
  FlashPeripheral,
  GPIOPeripheral,
  HASHPeripheral,
  HRTIM_CommonPeripheral,
  HRTIM_MasterPeripheral,
  HRTIM_TIMPeripheral,
  HSEMPeripheral,
  I2CPeripheral,
  IWDGPeripheral,
  JPEGPeripheral,
  LPTIMPeripheral,
  LPUARTPeripheral,
  LTDCPeripheral,
  MDIOSPeripheral,
  MDMAPeripheral,
  OPAMPPeripheral,
  OTG_HS_DEVICEPeripheral,
  OTG_HS_GLOBALPeripheral,
  OTG_HS_HOSTPeripheral,
  OTG_HS_PWRCLKPeripheral,
  PWRPeripheral,
  QUADSPIPeripheral,
  RAMECC1Peripheral,
  RAMECC2Peripheral,
  RAMECC3Peripheral,
  RCCPeripheral,
  RNGPeripheral,
  RTCPeripheral,
  SAIPeripheral,
  SDMMCPeripheral,
  SPDIFRXPeripheral,
  SPIPeripheral,
  SWPMIPeripheral,
  SYSCFGPeripheral,
  TIM1_8Peripheral,
  TIM2_3_4_5Peripheral,
  TIM6_7Peripheral,
  TIM12Peripheral,
  TIM13_14Peripheral,
  TIM15Peripheral,
  TIM15Peripheral,
  USARTPeripheral,
  VREFBUFPeripheral,
  WWDGPeripheral,
];

defvar interrupts = [
  Interrupt<"WWDG1", 0, "Window Watchdog interrupt">,
  Interrupt<"PVD_PVM", 1, "PVD through EXTI line">,
  Interrupt<"RTC_TAMP_STAMP_CSS_LSE", 2, "RTC tamper, timestamp">,
  Interrupt<"RTC_WKUP", 3, "RTC Wakeup interrupt">,
  Interrupt<"FLASH", 4, "Flash memory">,
  Interrupt<"RCC", 5, "RCC global interrupt">,
  Interrupt<"EXTI0", 6, "EXTI Line 0 interrupt">,
  Interrupt<"EXTI1", 7, "EXTI Line 1 interrupt">,
  Interrupt<"EXTI2", 8, "EXTI Line 2 interrupt">,
  Interrupt<"EXTI3", 9, "EXTI Line 3interrupt">,
  Interrupt<"EXTI4", 10, "EXTI Line 4interrupt">,
  Interrupt<"DMA_STR0", 11, "DMA1 Stream0">,
  Interrupt<"DMA_STR1", 12, "DMA1 Stream1">,
  Interrupt<"DMA_STR2", 13, "DMA1 Stream2">,
  Interrupt<"DMA_STR3", 14, "DMA1 Stream3">,
  Interrupt<"DMA_STR4", 15, "DMA1 Stream4">,
  Interrupt<"DMA_STR5", 16, "DMA1 Stream5">,
  Interrupt<"DMA_STR6", 17, "DMA1 Stream6">,
  Interrupt<"ADC1_2", 18, "ADC1 and ADC2">,
  Interrupt<"FDCAN1_IT0", 19, "FDCAN1 Interrupt 0">,
  Interrupt<"FDCAN2_IT0", 20, "FDCAN2 Interrupt 0">,
  Interrupt<"FDCAN1_IT1", 21, "FDCAN1 Interrupt 1">,
  Interrupt<"FDCAN2_IT1", 22, "FDCAN2 Interrupt 1">,
  Interrupt<"EXTI9_5", 23, "EXTI Line[9:5] interrupts">,
  Interrupt<"TIM1_BRK", 24, "TIM1 break interrupt">,
  Interrupt<"TIM1_UP", 25, "TIM1 update interrupt">,
  Interrupt<"TIM1_TRG_COM", 26, "TIM1 trigger and commutation">,
  Interrupt<"TIM_CC", 27, "TIM1 capture / compare">,
  Interrupt<"TIM2", 28, "TIM2 global interrupt">,
  Interrupt<"TIM3", 29, "TIM3 global interrupt">,
  Interrupt<"TIM4", 30, "TIM4 global interrupt">,
  Interrupt<"I2C1_EV", 31, "I2C1 event interrupt">,
  Interrupt<"I2C1_ER", 32, "I2C1 error interrupt">,
  Interrupt<"I2C2_EV", 33, "I2C2 event interrupt">,
  Interrupt<"I2C2_ER", 34, "I2C2 error interrupt">,
  Interrupt<"SPI1", 35, "SPI1 global interrupt">,
  Interrupt<"SPI2", 36, "SPI2 global interrupt">,
  Interrupt<"USART1", 37, "USART1 global interrupt">,
  Interrupt<"USART2", 38, "USART2 global interrupt">,
  Interrupt<"USART3", 39, "USART3 global interrupt">,
  Interrupt<"EXTI15_10", 40, "EXTI Line[15:10] interrupts">,
  Interrupt<"RTC_ALARM", 41, "RTC alarms (A and B)">,
  Interrupt<"TIM8_BRK_TIM12", 43, "TIM8 and 12 break global">,
  Interrupt<"TIM8_UP_TIM13", 44, "TIM8 and 13 update global">,
  Interrupt<"TIM8_TRG_COM_TIM14", 45, "TIM8 and 14 trigger /commutation and global">,
  Interrupt<"TIM8_CC", 46, "TIM8 capture / compare">,
  Interrupt<"DMA1_STR7", 47, "DMA1 Stream7">,
  Interrupt<"FMC", 48, "FMC global interrupt">,
  Interrupt<"SDMMC1", 49, "SDMMC global interrupt">,
  Interrupt<"TIM5", 50, "TIM5 global interrupt">,
  Interrupt<"SPI3", 51, "SPI3 global interrupt">,
  Interrupt<"UART4", 52, "UART4 global interrupt">,
  Interrupt<"UART5", 53, "UART5 global interrupt">,
  Interrupt<"TIM6_DAC", 54, "TIM6 global interrupt">,
  Interrupt<"TIM7", 55, "TIM7 global interrupt">,
  Interrupt<"DMA2_STR0", 56, "DMA2 Stream0 interrupt">,
  Interrupt<"DMA2_STR1", 57, "DMA2 Stream1 interrupt">,
  Interrupt<"DMA2_STR2", 58, "DMA2 Stream2 interrupt">,
  Interrupt<"DMA2_STR3", 59, "DMA2 Stream3 interrupt">,
  Interrupt<"DMA2_STR4", 60, "DMA2 Stream4 interrupt">,
  Interrupt<"ETH", 61, "Ethernet global interrupt">,
  Interrupt<"ETH_WKUP", 62, "Ethernet wakeup through EXTI">,
  Interrupt<"FDCAN_CAL", 63, "CAN2TX interrupts">,
  Interrupt<"CM4_SEV_IT", 65, "Arm Cortex-M4 Send even interrupt">,
  Interrupt<"DMA2_STR5", 68, "DMA2 Stream5 interrupt">,
  Interrupt<"DMA2_STR6", 69, "DMA2 Stream6 interrupt">,
  Interrupt<"DMA2_STR7", 70, "DMA2 Stream7 interrupt">,
  Interrupt<"USART6", 71, "USART6 global interrupt">,
  Interrupt<"I2C3_EV", 72, "I2C3 event interrupt">,
  Interrupt<"I2C3_ER", 73, "I2C3 error interrupt">,
  Interrupt<"OTG_HS_EP1_OUT", 74, "OTG_HS out global interrupt">,
  Interrupt<"OTG_HS_EP1_IN", 75, "OTG_HS in global interrupt">,
  Interrupt<"OTG_HS_WKUP", 76, "OTG_HS wakeup interrupt">,
  Interrupt<"OTG_HS", 77, "OTG_HS global interrupt">,
  Interrupt<"DCMI", 78, "DCMI global interrupt">,
  Interrupt<"CRYP", 79, "CRYP global interrupt">,
  Interrupt<"HASH_RNG", 80, "HASH and RNG">,
  Interrupt<"FPU", 81, "CPU2">,
  Interrupt<"UART7", 82, "UART7 global interrupt">,
  Interrupt<"UART8", 83, "UART8 global interrupt">,
  Interrupt<"SPI4", 84, "SPI4 global interrupt">,
  Interrupt<"SPI5", 85, "SPI5 global interrupt">,
  Interrupt<"SPI6", 86, "SPI6 global interrupt">,
  Interrupt<"SAI1", 87, "SAI1 global interrupt">,
  Interrupt<"LTDC", 88, "LCD-TFT global interrupt">,
  Interrupt<"LTDC_ER", 89, "LCD-TFT error interrupt">,
  Interrupt<"DMA2D", 90, "DMA2D global interrupt">,
  Interrupt<"SAI2", 91, "SAI2 global interrupt">,
  Interrupt<"QUADSPI", 92, "QuadSPI global interrupt">,
  Interrupt<"LPTIM1", 93, "LPTIM1 global interrupt">,
  Interrupt<"CEC", 94, "HDMI-CEC global interrupt">,
  Interrupt<"I2C4_EV", 95, "I2C4 event interrupt">,
  Interrupt<"I2C4_ER", 96, "I2C4 error interrupt">,
  Interrupt<"SPDIF", 97, "SPDIFRX global interrupt">,
  Interrupt<"OTG_FS_EP1_OUT", 98, "OTG_FS out global interrupt">,
  Interrupt<"OTG_FS_EP1_IN", 99, "OTG_FS in global interrupt">,
  Interrupt<"OTG_FS_WKUP", 100, "OTG_FS wakeup">,
  Interrupt<"OTG_FS", 101, "OTG_FS global interrupt">,
  Interrupt<"DMAMUX1_OV", 102, "DMAMUX1 overrun interrupt">,
  Interrupt<"HRTIM1_MST", 103, "HRTIM1 master timer interrupt">,
  Interrupt<"HRTIM1_TIMA", 104, "HRTIM1 timer A interrupt">,
  Interrupt<"HRTIM_TIMB", 105, "HRTIM1 timer B interrupt">,
  Interrupt<"HRTIM1_TIMC", 106, "HRTIM1 timer C interrupt">,
  Interrupt<"HRTIM1_TIMD", 107, "HRTIM1 timer D interrupt">,
  Interrupt<"HRTIM_TIME", 108, "HRTIM1 timer E interrupt">,
  Interrupt<"HRTIM1_FLT", 109, "HRTIM1 fault interrupt">,
  Interrupt<"DFSDM1_FLT0", 110, "DFSDM1 filter 0 interrupt">,
  Interrupt<"DFSDM1_FLT1", 111, "DFSDM1 filter 1 interrupt">,
  Interrupt<"DFSDM1_FLT2", 112, "DFSDM1 filter 2 interrupt">,
  Interrupt<"DFSDM1_FLT3", 113, "DFSDM1 filter 3 interrupt">,
  Interrupt<"SAI3", 114, "SAI3 global interrupt">,
  Interrupt<"SWPMI1", 115, "SWPMI global interrupt">,
  Interrupt<"TIM15", 116, "TIM15 global interrupt">,
  Interrupt<"TIM16", 117, "TIM16 global interrupt">,
  Interrupt<"TIM17", 118, "TIM17 global interrupt">,
  Interrupt<"MDIOS_WKUP", 119, "MDIOS wakeup">,
  Interrupt<"MDIOS", 120, "MDIOS global interrupt">,
  Interrupt<"JPEG", 121, "JPEG global interrupt">,
  Interrupt<"MDMA", 122, "MDMA">,
  Interrupt<"SDMMC", 124, "SDMMC global interrupt">,
  Interrupt<"HSEM0", 125, "HSEM global interrupt 1">,
  Interrupt<"ADC3", 127, "ADC3 global interrupt">,
  Interrupt<"DMAMUX2_OVR", 128, "DMAMUX2 overrun interrupt">,
  Interrupt<"BDMA_CH1", 129, "BDMA channel 1 interrupt">,
  Interrupt<"BDMA_CH2", 130, "BDMA channel 2 interrupt">,
  Interrupt<"BDMA_CH3", 131, "BDMA channel 3 interrupt">,
  Interrupt<"BDMA_CH4", 132, "BDMA channel 4 interrupt">,
  Interrupt<"BDMA_CH5", 133, "BDMA channel 5 interrupt">,
  Interrupt<"BDMA_CH6", 134, "BDMA channel 6 interrupt">,
  Interrupt<"BDMA_CH7", 135, "BDMA channel 7 interrupt">,
  Interrupt<"BDMA_CH8", 136, "BDMA channel 8 interrupt">,
  Interrupt<"COMP", 137, "COMP1 and COMP2">,
  Interrupt<"LPTIM2", 138, "LPTIM2 timer interrupt">,
  Interrupt<"LPTIM3", 139, "LPTIM3 timer interrupt">,
  Interrupt<"LPTIM4", 140, "LPTIM4 timer interrupt">,
  Interrupt<"LPTIM5", 141, "LPTIM5 timer interrupt">,
  Interrupt<"LPUART", 142, "LPUART global interrupt">,
  Interrupt<"WWDG2_RST", 143, "Window Watchdog interrupt">,
  Interrupt<"CRS", 144, "Clock Recovery System globa">,
  Interrupt<"SAI4", 146, "SAI4 global interrupt">,
  Interrupt<"HOLD_CORE", 148, "CPU1 hold">,
  Interrupt<"WKUP", 149, "WKUP1 to WKUP6 pins">,
];

defvar tags = [ "stm32h7x7" ];

def Flash512K_1 : MemoryRange<"FLASH", 0x08000000, KB<512>, true, ReadWrite, "Flash memory bank 1">;
def Flash512K_2 : MemoryRange<"FLASH2", 0x08100000, KB<512>, true, ReadWrite, "Flash memory bank 2">;
def Flash2M : MemoryRange<"FLASH", 0x08000000, MB<2>, true, ReadWrite, "Flash memory">;

def SRAM_AXI : MemoryRange<"SRAM_AXI", 0x24000000, KB<512>, true, ReadWrite, "SRAM mapped on AXI bus">;
def SRAM1_3 : MemoryRange<"RAM", 0x30000000, KB<288>, true, ReadWrite, "[SRAM1 : SRAM3]">;
def SRAMSRAM1 : MemoryRange<"SRAM1", 0x30000000, KB<128>, true, ReadWrite, "SRAM1 (D2 domain)">;
def SRAMSRAM2 : MemoryRange<"SRAM2", 0x30020000, KB<128>, true, ReadWrite, "SRAM2 (D2 domain)">;
def SRAMSRAM3 : MemoryRange<"SRAM3", 0x30040000, KB<32>, true, ReadWrite, "SRAM3 (D2 domain)">;
def SRAMSRAM4 : MemoryRange<"SRAM4", 0x38000000, KB<64>, true, ReadWrite, "SRAM4 (D3 domain)">;
def SRAMITCM : MemoryRange<"ITCM", 0x00000000, KB<64>, true, ReadWrite, "ITCM RAM (instruction)">;
def SRAMDTCM : MemoryRange<"DTCM", 0x20000000, KB<128>, true, ReadWrite, "DTCM RAM (instruction)">;
def SRAMBackup : MemoryRange<"SRAM_BACKUP", 0x38800000, KB<4>, true, ReadWrite, "Backup SRAM (Kbytes)">;

class STM32H7x7Variant<string Name, list<MemoryRange> Memories, MemoryUnit StackSize, list<string> Tags = [],
    string Description = "">
      : CortexMVariant<Name, Memories, interrupts, StackSize, !listconcat(tags, Tags), Description>;

defvar GMemories = [
    Flash512K_1,
    Flash512K_2,
    SRAM_AXI,
    SRAM1_3,
    SRAMSRAM1,
    SRAMSRAM2,
    SRAMSRAM3,
    SRAMSRAM4,
    SRAMITCM,
    SRAMDTCM,
    SRAMBackup,
];

defvar IMemories = [
    Flash2M,
    SRAM_AXI,
    SRAM1_3,
    SRAMSRAM1,
    SRAMSRAM2,
    SRAMSRAM3,
    SRAMSRAM4,
    SRAMITCM,
    SRAMDTCM,
    SRAMBackup,
];

class GVariant<string Name, list<string> Tags = []> : STM32H7x7Variant<Name, GMemories, KB<4>, Tags>;
class IVariant<string Name, list<string> Tags = []> : STM32H7x7Variant<Name, IMemories, KB<4>, Tags>;

defvar allVariants = [
    GVariant<"STM32H747AG">,
    GVariant<"STM32H747IG">,
    GVariant<"STM32H747BG">,
    GVariant<"STM32H747XG">,
    IVariant<"STM32H747ZI">,
    IVariant<"STM32H747AI">,
    IVariant<"STM32H747II">,
    IVariant<"STM32H747BI">,
    IVariant<"STM32H747XI">,

    IVariant<"STM32H757ZI">,
    IVariant<"STM32H757AI">,
    IVariant<"STM32H757II">,
    IVariant<"STM32H757BI">,
    IVariant<"STM32H757XI">,
];

def STM32H7x7 : Series<"STM32H7x7", CortexM7> {
  let variants = allVariants;
  let peripheralTypes = peripherals;
  let runtimePackages = [
      "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/def"
  ];
}

#endif // _STM32H7x7_TD
