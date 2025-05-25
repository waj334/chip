.syntax unified

.section .text.DefaultHandler
.global  DefaultHandler
.type    DefaultHandler, %function
DefaultHandler:
    wfe
    b    DefaultHandler
.size DefaultHandler, .-DefaultHandler

.macro IRQ handler
    .weak  \handler
    .set   \handler, DefaultHandler
.endm

.section .isr_vector, "a", %progbits
.global  __isr_vector
__isr_vector:
    .long __stack
    .long ResetHandler
    .long NMIHandler
    .long HardfaultHandler
    .long MemoryManageHandler
    .long BusFaultHandler
    .long UsageFaultHandler
    .long 0
    .long 0
    .long 0
    .long 0
    .long SVCallHandler
    .long DebugMonitorHandler
    .long 0
    .long PendSVHandlerHandler
    .long SysTickHandlerHandler
    .long WWDG1Handler
    .long PVD_PVMHandler
    .long RTC_TAMP_STAMP_CSS_LSEHandler
    .long RTC_WKUPHandler
    .long FLASHHandler
    .long RCCHandler
    .long EXTI0Handler
    .long EXTI1Handler
    .long EXTI2Handler
    .long EXTI3Handler
    .long EXTI4Handler
    .long DMA_STR0Handler
    .long DMA_STR1Handler
    .long DMA_STR2Handler
    .long DMA_STR3Handler
    .long DMA_STR4Handler
    .long DMA_STR5Handler
    .long DMA_STR6Handler
    .long ADC1_2Handler
    .long FDCAN1_IT0Handler
    .long FDCAN2_IT0Handler
    .long FDCAN1_IT1Handler
    .long FDCAN2_IT1Handler
    .long EXTI9_5Handler
    .long TIM1_BRKHandler
    .long TIM1_UPHandler
    .long TIM1_TRG_COMHandler
    .long TIM_CCHandler
    .long TIM2Handler
    .long TIM3Handler
    .long TIM4Handler
    .long I2C1_EVHandler
    .long I2C1_ERHandler
    .long I2C2_EVHandler
    .long I2C2_ERHandler
    .long SPI1Handler
    .long SPI2Handler
    .long USART1Handler
    .long USART2Handler
    .long USART3Handler
    .long EXTI15_10Handler
    .long RTC_ALARMHandler
    .long 0
    .long TIM8_BRK_TIM12Handler
    .long TIM8_UP_TIM13Handler
    .long TIM8_TRG_COM_TIM14Handler
    .long TIM8_CCHandler
    .long DMA1_STR7Handler
    .long FMCHandler
    .long SDMMC1Handler
    .long TIM5Handler
    .long SPI3Handler
    .long UART4Handler
    .long UART5Handler
    .long TIM6_DACHandler
    .long TIM7Handler
    .long DMA2_STR0Handler
    .long DMA2_STR1Handler
    .long DMA2_STR2Handler
    .long DMA2_STR3Handler
    .long DMA2_STR4Handler
    .long ETHHandler
    .long ETH_WKUPHandler
    .long FDCAN_CALHandler
    .long 0
    .long CM4_SEV_ITHandler
    .long 0
    .long 0
    .long DMA2_STR5Handler
    .long DMA2_STR6Handler
    .long DMA2_STR7Handler
    .long USART6Handler
    .long I2C3_EVHandler
    .long I2C3_ERHandler
    .long OTG_HS_EP1_OUTHandler
    .long OTG_HS_EP1_INHandler
    .long OTG_HS_WKUPHandler
    .long OTG_HSHandler
    .long DCMIHandler
    .long CRYPHandler
    .long HASH_RNGHandler
    .long FPUHandler
    .long UART7Handler
    .long UART8Handler
    .long SPI4Handler
    .long SPI5Handler
    .long SPI6Handler
    .long SAI1Handler
    .long LTDCHandler
    .long LTDC_ERHandler
    .long DMA2DHandler
    .long SAI2Handler
    .long QUADSPIHandler
    .long LPTIM1Handler
    .long CECHandler
    .long I2C4_EVHandler
    .long I2C4_ERHandler
    .long SPDIFHandler
    .long OTG_FS_EP1_OUTHandler
    .long OTG_FS_EP1_INHandler
    .long OTG_FS_WKUPHandler
    .long OTG_FSHandler
    .long DMAMUX1_OVHandler
    .long HRTIM1_MSTHandler
    .long HRTIM1_TIMAHandler
    .long HRTIM_TIMBHandler
    .long HRTIM1_TIMCHandler
    .long HRTIM1_TIMDHandler
    .long HRTIM_TIMEHandler
    .long HRTIM1_FLTHandler
    .long DFSDM1_FLT0Handler
    .long DFSDM1_FLT1Handler
    .long DFSDM1_FLT2Handler
    .long DFSDM1_FLT3Handler
    .long SAI3Handler
    .long SWPMI1Handler
    .long TIM15Handler
    .long TIM16Handler
    .long TIM17Handler
    .long MDIOS_WKUPHandler
    .long MDIOSHandler
    .long JPEGHandler
    .long MDMAHandler
    .long 0
    .long SDMMCHandler
    .long HSEM0Handler
    .long 0
    .long ADC3Handler
    .long DMAMUX2_OVRHandler
    .long BDMA_CH1Handler
    .long BDMA_CH2Handler
    .long BDMA_CH3Handler
    .long BDMA_CH4Handler
    .long BDMA_CH5Handler
    .long BDMA_CH6Handler
    .long BDMA_CH7Handler
    .long BDMA_CH8Handler
    .long COMPHandler
    .long LPTIM2Handler
    .long LPTIM3Handler
    .long LPTIM4Handler
    .long LPTIM5Handler
    .long LPUARTHandler
    .long WWDG2_RSTHandler
    .long CRSHandler
    .long 0
    .long SAI4Handler
    .long 0
    .long HOLD_COREHandler
    .long WKUPHandler

    IRQ ResetHandler                      // Reset Vector, invoked on Power up and warm reset.
    IRQ NMIHandler                        // Non maskable Interrupt, cannot be stopped or preempted.
    IRQ HardfaultHandler                  // Hard Fault, all classes of Fault.
    IRQ MemoryManageHandler               // Memory Management, MPU mismatch, including Access Violation and No Match.
    IRQ BusFaultHandler                   // Bus Fault, Pre-Fetch-, Memory Access Fault, other address/memory related Fault.
    IRQ UsageFaultHandler                 // Usage Fault, i.e. Undef Instruction, Illegal State Transition
    IRQ SVCallHandler                     // System Service Call via SVC instruction.
    IRQ DebugMonitorHandler               // Debug Monitor.
    IRQ PendSVHandlerHandler              // Pendable request for system service.
    IRQ SysTickHandlerHandler             // System Tick Timer.
    IRQ WWDG1Handler                      // Window Watchdog interrupt
    IRQ PVD_PVMHandler                    // PVD through EXTI line
    IRQ RTC_TAMP_STAMP_CSS_LSEHandler     // RTC tamper, timestamp
    IRQ RTC_WKUPHandler                   // RTC Wakeup interrupt
    IRQ FLASHHandler                      // Flash memory
    IRQ RCCHandler                        // RCC global interrupt
    IRQ EXTI0Handler                      // EXTI Line 0 interrupt
    IRQ EXTI1Handler                      // EXTI Line 1 interrupt
    IRQ EXTI2Handler                      // EXTI Line 2 interrupt
    IRQ EXTI3Handler                      // EXTI Line 3interrupt
    IRQ EXTI4Handler                      // EXTI Line 4interrupt
    IRQ DMA_STR0Handler                   // DMA1 Stream0
    IRQ DMA_STR1Handler                   // DMA1 Stream1
    IRQ DMA_STR2Handler                   // DMA1 Stream2
    IRQ DMA_STR3Handler                   // DMA1 Stream3
    IRQ DMA_STR4Handler                   // DMA1 Stream4
    IRQ DMA_STR5Handler                   // DMA1 Stream5
    IRQ DMA_STR6Handler                   // DMA1 Stream6
    IRQ ADC1_2Handler                     // ADC1 and ADC2
    IRQ FDCAN1_IT0Handler                 // FDCAN1 Interrupt 0
    IRQ FDCAN2_IT0Handler                 // FDCAN2 Interrupt 0
    IRQ FDCAN1_IT1Handler                 // FDCAN1 Interrupt 1
    IRQ FDCAN2_IT1Handler                 // FDCAN2 Interrupt 1
    IRQ EXTI9_5Handler                    // EXTI Line[9:5] interrupts
    IRQ TIM1_BRKHandler                   // TIM1 break interrupt
    IRQ TIM1_UPHandler                    // TIM1 update interrupt
    IRQ TIM1_TRG_COMHandler               // TIM1 trigger and commutation
    IRQ TIM_CCHandler                     // TIM1 capture / compare
    IRQ TIM2Handler                       // TIM2 global interrupt
    IRQ TIM3Handler                       // TIM3 global interrupt
    IRQ TIM4Handler                       // TIM4 global interrupt
    IRQ I2C1_EVHandler                    // I2C1 event interrupt
    IRQ I2C1_ERHandler                    // I2C1 error interrupt
    IRQ I2C2_EVHandler                    // I2C2 event interrupt
    IRQ I2C2_ERHandler                    // I2C2 error interrupt
    IRQ SPI1Handler                       // SPI1 global interrupt
    IRQ SPI2Handler                       // SPI2 global interrupt
    IRQ USART1Handler                     // USART1 global interrupt
    IRQ USART2Handler                     // USART2 global interrupt
    IRQ USART3Handler                     // USART3 global interrupt
    IRQ EXTI15_10Handler                  // EXTI Line[15:10] interrupts
    IRQ RTC_ALARMHandler                  // RTC alarms (A and B)
    IRQ TIM8_BRK_TIM12Handler             // TIM8 and 12 break global
    IRQ TIM8_UP_TIM13Handler              // TIM8 and 13 update global
    IRQ TIM8_TRG_COM_TIM14Handler         // TIM8 and 14 trigger /commutation and global
    IRQ TIM8_CCHandler                    // TIM8 capture / compare
    IRQ DMA1_STR7Handler                  // DMA1 Stream7
    IRQ FMCHandler                        // FMC global interrupt
    IRQ SDMMC1Handler                     // SDMMC global interrupt
    IRQ TIM5Handler                       // TIM5 global interrupt
    IRQ SPI3Handler                       // SPI3 global interrupt
    IRQ UART4Handler                      // UART4 global interrupt
    IRQ UART5Handler                      // UART5 global interrupt
    IRQ TIM6_DACHandler                   // TIM6 global interrupt
    IRQ TIM7Handler                       // TIM7 global interrupt
    IRQ DMA2_STR0Handler                  // DMA2 Stream0 interrupt
    IRQ DMA2_STR1Handler                  // DMA2 Stream1 interrupt
    IRQ DMA2_STR2Handler                  // DMA2 Stream2 interrupt
    IRQ DMA2_STR3Handler                  // DMA2 Stream3 interrupt
    IRQ DMA2_STR4Handler                  // DMA2 Stream4 interrupt
    IRQ ETHHandler                        // Ethernet global interrupt
    IRQ ETH_WKUPHandler                   // Ethernet wakeup through EXTI
    IRQ FDCAN_CALHandler                  // CAN2TX interrupts
    IRQ CM4_SEV_ITHandler                 // Arm Cortex-M4 Send even interrupt
    IRQ DMA2_STR5Handler                  // DMA2 Stream5 interrupt
    IRQ DMA2_STR6Handler                  // DMA2 Stream6 interrupt
    IRQ DMA2_STR7Handler                  // DMA2 Stream7 interrupt
    IRQ USART6Handler                     // USART6 global interrupt
    IRQ I2C3_EVHandler                    // I2C3 event interrupt
    IRQ I2C3_ERHandler                    // I2C3 error interrupt
    IRQ OTG_HS_EP1_OUTHandler             // OTG_HS out global interrupt
    IRQ OTG_HS_EP1_INHandler              // OTG_HS in global interrupt
    IRQ OTG_HS_WKUPHandler                // OTG_HS wakeup interrupt
    IRQ OTG_HSHandler                     // OTG_HS global interrupt
    IRQ DCMIHandler                       // DCMI global interrupt
    IRQ CRYPHandler                       // CRYP global interrupt
    IRQ HASH_RNGHandler                   // HASH and RNG
    IRQ FPUHandler                        // CPU2
    IRQ UART7Handler                      // UART7 global interrupt
    IRQ UART8Handler                      // UART8 global interrupt
    IRQ SPI4Handler                       // SPI4 global interrupt
    IRQ SPI5Handler                       // SPI5 global interrupt
    IRQ SPI6Handler                       // SPI6 global interrupt
    IRQ SAI1Handler                       // SAI1 global interrupt
    IRQ LTDCHandler                       // LCD-TFT global interrupt
    IRQ LTDC_ERHandler                    // LCD-TFT error interrupt
    IRQ DMA2DHandler                      // DMA2D global interrupt
    IRQ SAI2Handler                       // SAI2 global interrupt
    IRQ QUADSPIHandler                    // QuadSPI global interrupt
    IRQ LPTIM1Handler                     // LPTIM1 global interrupt
    IRQ CECHandler                        // HDMI-CEC global interrupt
    IRQ I2C4_EVHandler                    // I2C4 event interrupt
    IRQ I2C4_ERHandler                    // I2C4 error interrupt
    IRQ SPDIFHandler                      // SPDIFRX global interrupt
    IRQ OTG_FS_EP1_OUTHandler             // OTG_FS out global interrupt
    IRQ OTG_FS_EP1_INHandler              // OTG_FS in global interrupt
    IRQ OTG_FS_WKUPHandler                // OTG_FS wakeup
    IRQ OTG_FSHandler                     // OTG_FS global interrupt
    IRQ DMAMUX1_OVHandler                 // DMAMUX1 overrun interrupt
    IRQ HRTIM1_MSTHandler                 // HRTIM1 master timer interrupt
    IRQ HRTIM1_TIMAHandler                // HRTIM1 timer A interrupt
    IRQ HRTIM_TIMBHandler                 // HRTIM1 timer B interrupt
    IRQ HRTIM1_TIMCHandler                // HRTIM1 timer C interrupt
    IRQ HRTIM1_TIMDHandler                // HRTIM1 timer D interrupt
    IRQ HRTIM_TIMEHandler                 // HRTIM1 timer E interrupt
    IRQ HRTIM1_FLTHandler                 // HRTIM1 fault interrupt
    IRQ DFSDM1_FLT0Handler                // DFSDM1 filter 0 interrupt
    IRQ DFSDM1_FLT1Handler                // DFSDM1 filter 1 interrupt
    IRQ DFSDM1_FLT2Handler                // DFSDM1 filter 2 interrupt
    IRQ DFSDM1_FLT3Handler                // DFSDM1 filter 3 interrupt
    IRQ SAI3Handler                       // SAI3 global interrupt
    IRQ SWPMI1Handler                     // SWPMI global interrupt
    IRQ TIM15Handler                      // TIM15 global interrupt
    IRQ TIM16Handler                      // TIM16 global interrupt
    IRQ TIM17Handler                      // TIM17 global interrupt
    IRQ MDIOS_WKUPHandler                 // MDIOS wakeup
    IRQ MDIOSHandler                      // MDIOS global interrupt
    IRQ JPEGHandler                       // JPEG global interrupt
    IRQ MDMAHandler                       // MDMA
    IRQ SDMMCHandler                      // SDMMC global interrupt
    IRQ HSEM0Handler                      // HSEM global interrupt 1
    IRQ ADC3Handler                       // ADC3 global interrupt
    IRQ DMAMUX2_OVRHandler                // DMAMUX2 overrun interrupt
    IRQ BDMA_CH1Handler                   // BDMA channel 1 interrupt
    IRQ BDMA_CH2Handler                   // BDMA channel 2 interrupt
    IRQ BDMA_CH3Handler                   // BDMA channel 3 interrupt
    IRQ BDMA_CH4Handler                   // BDMA channel 4 interrupt
    IRQ BDMA_CH5Handler                   // BDMA channel 5 interrupt
    IRQ BDMA_CH6Handler                   // BDMA channel 6 interrupt
    IRQ BDMA_CH7Handler                   // BDMA channel 7 interrupt
    IRQ BDMA_CH8Handler                   // BDMA channel 8 interrupt
    IRQ COMPHandler                       // COMP1 and COMP2
    IRQ LPTIM2Handler                     // LPTIM2 timer interrupt
    IRQ LPTIM3Handler                     // LPTIM2 timer interrupt
    IRQ LPTIM4Handler                     // LPTIM2 timer interrupt
    IRQ LPTIM5Handler                     // LPTIM2 timer interrupt
    IRQ LPUARTHandler                     // LPUART global interrupt
    IRQ WWDG2_RSTHandler                  // Window Watchdog interrupt
    IRQ CRSHandler                        // Clock Recovery System globa
    IRQ SAI4Handler                       // SAI4 global interrupt
    IRQ HOLD_COREHandler                  // CPU1 hold
    IRQ WKUPHandler                       // WKUP1 to WKUP6 pins

