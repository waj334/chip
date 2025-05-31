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
    .long NmiHandler
    .long HardfaultHandler
    .long MemoryManageHandler
    .long BusFaultHandler
    .long UsageFaultHandler
    .long 0
    .long 0
    .long 0
    .long 0
    .long SvCallHandler
    .long DebugMonitorHandler
    .long 0
    .long PendSvHandler
    .long SysTickHandler
    .long Wwdg1Handler
    .long PvdPvmHandler
    .long RtcTampStampCssLseHandler
    .long RtcWkupHandler
    .long FlashHandler
    .long RccHandler
    .long Exti0Handler
    .long Exti1Handler
    .long Exti2Handler
    .long Exti3Handler
    .long Exti4Handler
    .long DmaStr0Handler
    .long DmaStr1Handler
    .long DmaStr2Handler
    .long DmaStr3Handler
    .long DmaStr4Handler
    .long DmaStr5Handler
    .long DmaStr6Handler
    .long Adc12Handler
    .long Fdcan1It0Handler
    .long Fdcan2It0Handler
    .long Fdcan1It1Handler
    .long Fdcan2It1Handler
    .long Exti95Handler
    .long Tim1BrkHandler
    .long Tim1UpHandler
    .long Tim1TrgComHandler
    .long TimCcHandler
    .long Tim2Handler
    .long Tim3Handler
    .long Tim4Handler
    .long I2c1EvHandler
    .long I2c1ErHandler
    .long I2c2EvHandler
    .long I2c2ErHandler
    .long Spi1Handler
    .long Spi2Handler
    .long Usart1Handler
    .long Usart2Handler
    .long Usart3Handler
    .long Exti1510Handler
    .long RtcAlarmHandler
    .long 0
    .long Tim8BrkTim12Handler
    .long Tim8UpTim13Handler
    .long Tim8TrgComTim14Handler
    .long Tim8CcHandler
    .long Dma1Str7Handler
    .long FmcHandler
    .long Sdmmc1Handler
    .long Tim5Handler
    .long Spi3Handler
    .long Uart4Handler
    .long Uart5Handler
    .long Tim6DacHandler
    .long Tim7Handler
    .long Dma2Str0Handler
    .long Dma2Str1Handler
    .long Dma2Str2Handler
    .long Dma2Str3Handler
    .long Dma2Str4Handler
    .long EthHandler
    .long EthWkupHandler
    .long FdcanCalHandler
    .long 0
    .long Cm4SevItHandler
    .long 0
    .long 0
    .long Dma2Str5Handler
    .long Dma2Str6Handler
    .long Dma2Str7Handler
    .long Usart6Handler
    .long I2c3EvHandler
    .long I2c3ErHandler
    .long OtgHsEp1OutHandler
    .long OtgHsEp1InHandler
    .long OtgHsWkupHandler
    .long OtgHsHandler
    .long DcmiHandler
    .long CrypHandler
    .long HashRngHandler
    .long FpuHandler
    .long Uart7Handler
    .long Uart8Handler
    .long Spi4Handler
    .long Spi5Handler
    .long Spi6Handler
    .long Sai1Handler
    .long LtdcHandler
    .long LtdcErHandler
    .long Dma2dHandler
    .long Sai2Handler
    .long QuadspiHandler
    .long Lptim1Handler
    .long CecHandler
    .long I2c4EvHandler
    .long I2c4ErHandler
    .long SpdifHandler
    .long OtgFsEp1OutHandler
    .long OtgFsEp1InHandler
    .long OtgFsWkupHandler
    .long OtgFsHandler
    .long Dmamux1OvHandler
    .long Hrtim1MstHandler
    .long Hrtim1TimaHandler
    .long HrtimTimbHandler
    .long Hrtim1TimcHandler
    .long Hrtim1TimdHandler
    .long HrtimTimeHandler
    .long Hrtim1FltHandler
    .long Dfsdm1Flt0Handler
    .long Dfsdm1Flt1Handler
    .long Dfsdm1Flt2Handler
    .long Dfsdm1Flt3Handler
    .long Sai3Handler
    .long Swpmi1Handler
    .long Tim15Handler
    .long Tim16Handler
    .long Tim17Handler
    .long MdiosWkupHandler
    .long MdiosHandler
    .long JpegHandler
    .long MdmaHandler
    .long 0
    .long SdmmcHandler
    .long Hsem0Handler
    .long 0
    .long Adc3Handler
    .long Dmamux2OvrHandler
    .long BdmaCh1Handler
    .long BdmaCh2Handler
    .long BdmaCh3Handler
    .long BdmaCh4Handler
    .long BdmaCh5Handler
    .long BdmaCh6Handler
    .long BdmaCh7Handler
    .long BdmaCh8Handler
    .long CompHandler
    .long Lptim2Handler
    .long Lptim3Handler
    .long Lptim4Handler
    .long Lptim5Handler
    .long LpuartHandler
    .long Wwdg2RstHandler
    .long CrsHandler
    .long 0
    .long Sai4Handler
    .long 0
    .long HoldCoreHandler
    .long WkupHandler

    IRQ ResetHandler                      // Reset Vector, invoked on Power up and warm reset.
    IRQ NmiHandler                        // Non maskable Interrupt, cannot be stopped or preempted.
    IRQ HardfaultHandler                  // Hard Fault, all classes of Fault.
    IRQ MemoryManageHandler               // Memory Management, MPU mismatch, including Access Violation and No Match.
    IRQ BusFaultHandler                   // Bus Fault, Pre-Fetch-, Memory Access Fault, other address/memory related Fault.
    IRQ UsageFaultHandler                 // Usage Fault, i.e. Undef Instruction, Illegal State Transition
    IRQ SvCallHandler                     // System Service Call via SVC instruction.
    IRQ DebugMonitorHandler               // Debug Monitor.
    IRQ PendSvHandler                     // Pendable request for system service.
    IRQ SysTickHandler                    // System Tick Timer.
    IRQ Wwdg1Handler                      // Window Watchdog interrupt
    IRQ PvdPvmHandler                     // PVD through EXTI line
    IRQ RtcTampStampCssLseHandler         // RTC tamper, timestamp
    IRQ RtcWkupHandler                    // RTC Wakeup interrupt
    IRQ FlashHandler                      // Flash memory
    IRQ RccHandler                        // RCC global interrupt
    IRQ Exti0Handler                      // EXTI Line 0 interrupt
    IRQ Exti1Handler                      // EXTI Line 1 interrupt
    IRQ Exti2Handler                      // EXTI Line 2 interrupt
    IRQ Exti3Handler                      // EXTI Line 3interrupt
    IRQ Exti4Handler                      // EXTI Line 4interrupt
    IRQ DmaStr0Handler                    // DMA1 Stream0
    IRQ DmaStr1Handler                    // DMA1 Stream1
    IRQ DmaStr2Handler                    // DMA1 Stream2
    IRQ DmaStr3Handler                    // DMA1 Stream3
    IRQ DmaStr4Handler                    // DMA1 Stream4
    IRQ DmaStr5Handler                    // DMA1 Stream5
    IRQ DmaStr6Handler                    // DMA1 Stream6
    IRQ Adc12Handler                      // ADC1 and ADC2
    IRQ Fdcan1It0Handler                  // FDCAN1 Interrupt 0
    IRQ Fdcan2It0Handler                  // FDCAN2 Interrupt 0
    IRQ Fdcan1It1Handler                  // FDCAN1 Interrupt 1
    IRQ Fdcan2It1Handler                  // FDCAN2 Interrupt 1
    IRQ Exti95Handler                     // EXTI Line[9:5] interrupts
    IRQ Tim1BrkHandler                    // TIM1 break interrupt
    IRQ Tim1UpHandler                     // TIM1 update interrupt
    IRQ Tim1TrgComHandler                 // TIM1 trigger and commutation
    IRQ TimCcHandler                      // TIM1 capture / compare
    IRQ Tim2Handler                       // TIM2 global interrupt
    IRQ Tim3Handler                       // TIM3 global interrupt
    IRQ Tim4Handler                       // TIM4 global interrupt
    IRQ I2c1EvHandler                     // I2C1 event interrupt
    IRQ I2c1ErHandler                     // I2C1 error interrupt
    IRQ I2c2EvHandler                     // I2C2 event interrupt
    IRQ I2c2ErHandler                     // I2C2 error interrupt
    IRQ Spi1Handler                       // SPI1 global interrupt
    IRQ Spi2Handler                       // SPI2 global interrupt
    IRQ Usart1Handler                     // USART1 global interrupt
    IRQ Usart2Handler                     // USART2 global interrupt
    IRQ Usart3Handler                     // USART3 global interrupt
    IRQ Exti1510Handler                   // EXTI Line[15:10] interrupts
    IRQ RtcAlarmHandler                   // RTC alarms (A and B)
    IRQ Tim8BrkTim12Handler               // TIM8 and 12 break global
    IRQ Tim8UpTim13Handler                // TIM8 and 13 update global
    IRQ Tim8TrgComTim14Handler            // TIM8 and 14 trigger /commutation and global
    IRQ Tim8CcHandler                     // TIM8 capture / compare
    IRQ Dma1Str7Handler                   // DMA1 Stream7
    IRQ FmcHandler                        // FMC global interrupt
    IRQ Sdmmc1Handler                     // SDMMC global interrupt
    IRQ Tim5Handler                       // TIM5 global interrupt
    IRQ Spi3Handler                       // SPI3 global interrupt
    IRQ Uart4Handler                      // UART4 global interrupt
    IRQ Uart5Handler                      // UART5 global interrupt
    IRQ Tim6DacHandler                    // TIM6 global interrupt
    IRQ Tim7Handler                       // TIM7 global interrupt
    IRQ Dma2Str0Handler                   // DMA2 Stream0 interrupt
    IRQ Dma2Str1Handler                   // DMA2 Stream1 interrupt
    IRQ Dma2Str2Handler                   // DMA2 Stream2 interrupt
    IRQ Dma2Str3Handler                   // DMA2 Stream3 interrupt
    IRQ Dma2Str4Handler                   // DMA2 Stream4 interrupt
    IRQ EthHandler                        // Ethernet global interrupt
    IRQ EthWkupHandler                    // Ethernet wakeup through EXTI
    IRQ FdcanCalHandler                   // CAN2TX interrupts
    IRQ Cm4SevItHandler                   // Arm Cortex-M4 Send even interrupt
    IRQ Dma2Str5Handler                   // DMA2 Stream5 interrupt
    IRQ Dma2Str6Handler                   // DMA2 Stream6 interrupt
    IRQ Dma2Str7Handler                   // DMA2 Stream7 interrupt
    IRQ Usart6Handler                     // USART6 global interrupt
    IRQ I2c3EvHandler                     // I2C3 event interrupt
    IRQ I2c3ErHandler                     // I2C3 error interrupt
    IRQ OtgHsEp1OutHandler                // OTG_HS out global interrupt
    IRQ OtgHsEp1InHandler                 // OTG_HS in global interrupt
    IRQ OtgHsWkupHandler                  // OTG_HS wakeup interrupt
    IRQ OtgHsHandler                      // OTG_HS global interrupt
    IRQ DcmiHandler                       // DCMI global interrupt
    IRQ CrypHandler                       // CRYP global interrupt
    IRQ HashRngHandler                    // HASH and RNG
    IRQ FpuHandler                        // CPU2
    IRQ Uart7Handler                      // UART7 global interrupt
    IRQ Uart8Handler                      // UART8 global interrupt
    IRQ Spi4Handler                       // SPI4 global interrupt
    IRQ Spi5Handler                       // SPI5 global interrupt
    IRQ Spi6Handler                       // SPI6 global interrupt
    IRQ Sai1Handler                       // SAI1 global interrupt
    IRQ LtdcHandler                       // LCD-TFT global interrupt
    IRQ LtdcErHandler                     // LCD-TFT error interrupt
    IRQ Dma2dHandler                      // DMA2D global interrupt
    IRQ Sai2Handler                       // SAI2 global interrupt
    IRQ QuadspiHandler                    // QuadSPI global interrupt
    IRQ Lptim1Handler                     // LPTIM1 global interrupt
    IRQ CecHandler                        // HDMI-CEC global interrupt
    IRQ I2c4EvHandler                     // I2C4 event interrupt
    IRQ I2c4ErHandler                     // I2C4 error interrupt
    IRQ SpdifHandler                      // SPDIFRX global interrupt
    IRQ OtgFsEp1OutHandler                // OTG_FS out global interrupt
    IRQ OtgFsEp1InHandler                 // OTG_FS in global interrupt
    IRQ OtgFsWkupHandler                  // OTG_FS wakeup
    IRQ OtgFsHandler                      // OTG_FS global interrupt
    IRQ Dmamux1OvHandler                  // DMAMUX1 overrun interrupt
    IRQ Hrtim1MstHandler                  // HRTIM1 master timer interrupt
    IRQ Hrtim1TimaHandler                 // HRTIM1 timer A interrupt
    IRQ HrtimTimbHandler                  // HRTIM1 timer B interrupt
    IRQ Hrtim1TimcHandler                 // HRTIM1 timer C interrupt
    IRQ Hrtim1TimdHandler                 // HRTIM1 timer D interrupt
    IRQ HrtimTimeHandler                  // HRTIM1 timer E interrupt
    IRQ Hrtim1FltHandler                  // HRTIM1 fault interrupt
    IRQ Dfsdm1Flt0Handler                 // DFSDM1 filter 0 interrupt
    IRQ Dfsdm1Flt1Handler                 // DFSDM1 filter 1 interrupt
    IRQ Dfsdm1Flt2Handler                 // DFSDM1 filter 2 interrupt
    IRQ Dfsdm1Flt3Handler                 // DFSDM1 filter 3 interrupt
    IRQ Sai3Handler                       // SAI3 global interrupt
    IRQ Swpmi1Handler                     // SWPMI global interrupt
    IRQ Tim15Handler                      // TIM15 global interrupt
    IRQ Tim16Handler                      // TIM16 global interrupt
    IRQ Tim17Handler                      // TIM17 global interrupt
    IRQ MdiosWkupHandler                  // MDIOS wakeup
    IRQ MdiosHandler                      // MDIOS global interrupt
    IRQ JpegHandler                       // JPEG global interrupt
    IRQ MdmaHandler                       // MDMA
    IRQ SdmmcHandler                      // SDMMC global interrupt
    IRQ Hsem0Handler                      // HSEM global interrupt 1
    IRQ Adc3Handler                       // ADC3 global interrupt
    IRQ Dmamux2OvrHandler                 // DMAMUX2 overrun interrupt
    IRQ BdmaCh1Handler                    // BDMA channel 1 interrupt
    IRQ BdmaCh2Handler                    // BDMA channel 2 interrupt
    IRQ BdmaCh3Handler                    // BDMA channel 3 interrupt
    IRQ BdmaCh4Handler                    // BDMA channel 4 interrupt
    IRQ BdmaCh5Handler                    // BDMA channel 5 interrupt
    IRQ BdmaCh6Handler                    // BDMA channel 6 interrupt
    IRQ BdmaCh7Handler                    // BDMA channel 7 interrupt
    IRQ BdmaCh8Handler                    // BDMA channel 8 interrupt
    IRQ CompHandler                       // COMP1 and COMP2
    IRQ Lptim2Handler                     // LPTIM2 timer interrupt
    IRQ Lptim3Handler                     // LPTIM3 timer interrupt
    IRQ Lptim4Handler                     // LPTIM4 timer interrupt
    IRQ Lptim5Handler                     // LPTIM5 timer interrupt
    IRQ LpuartHandler                     // LPUART global interrupt
    IRQ Wwdg2RstHandler                   // Window Watchdog interrupt
    IRQ CrsHandler                        // Clock Recovery System globa
    IRQ Sai4Handler                       // SAI4 global interrupt
    IRQ HoldCoreHandler                   // CPU1 hold
    IRQ WkupHandler                       // WKUP1 to WKUP6 pins

