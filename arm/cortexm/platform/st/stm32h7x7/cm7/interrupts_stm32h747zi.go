//go:build stm32h7x7 && stm32h747zi

package stm32h7x7

import "pkg.si-go.dev/chip/arm/cortexm/runtime"

const (
	// IrqReset Reset Vector, invoked on Power up and warm reset.
	IrqReset runtime.Interrupt = -15

	// IrqNmi Non maskable Interrupt, cannot be stopped or preempted.
	IrqNmi runtime.Interrupt = -14

	// IrqHardfault Hard Fault, all classes of Fault.
	IrqHardfault runtime.Interrupt = -13

	// IrqMemoryManage Memory Management, MPU mismatch, including Access Violation and No Match.
	IrqMemoryManage runtime.Interrupt = -12

	// IrqBusFault Bus Fault, Pre-Fetch-, Memory Access Fault, other address/memory related Fault.
	IrqBusFault runtime.Interrupt = -11

	// IrqUsageFault Usage Fault, i.e. Undef Instruction, Illegal State Transition
	IrqUsageFault runtime.Interrupt = -10

	// IrqSvCall System Service Call via SVC instruction.
	IrqSvCall runtime.Interrupt = -5

	// IrqDebugMonitor Debug Monitor.
	IrqDebugMonitor runtime.Interrupt = -4

	// IrqPendSv Pendable request for system service.
	IrqPendSv runtime.Interrupt = -2

	// IrqSysTick System Tick Timer.
	IrqSysTick runtime.Interrupt = -1

	// IrqWwdg1 Window Watchdog interrupt
	IrqWwdg1 runtime.Interrupt = 0

	// IrqPvdPvm PVD through EXTI line
	IrqPvdPvm runtime.Interrupt = 1

	// IrqRtcTampStampCssLse RTC tamper, timestamp
	IrqRtcTampStampCssLse runtime.Interrupt = 2

	// IrqRtcWkup RTC Wakeup interrupt
	IrqRtcWkup runtime.Interrupt = 3

	// IrqFlash Flash memory
	IrqFlash runtime.Interrupt = 4

	// IrqRcc RCC global interrupt
	IrqRcc runtime.Interrupt = 5

	// IrqExti0 EXTI Line 0 interrupt
	IrqExti0 runtime.Interrupt = 6

	// IrqExti1 EXTI Line 1 interrupt
	IrqExti1 runtime.Interrupt = 7

	// IrqExti2 EXTI Line 2 interrupt
	IrqExti2 runtime.Interrupt = 8

	// IrqExti3 EXTI Line 3interrupt
	IrqExti3 runtime.Interrupt = 9

	// IrqExti4 EXTI Line 4interrupt
	IrqExti4 runtime.Interrupt = 10

	// IrqDmaStr0 DMA1 Stream0
	IrqDmaStr0 runtime.Interrupt = 11

	// IrqDmaStr1 DMA1 Stream1
	IrqDmaStr1 runtime.Interrupt = 12

	// IrqDmaStr2 DMA1 Stream2
	IrqDmaStr2 runtime.Interrupt = 13

	// IrqDmaStr3 DMA1 Stream3
	IrqDmaStr3 runtime.Interrupt = 14

	// IrqDmaStr4 DMA1 Stream4
	IrqDmaStr4 runtime.Interrupt = 15

	// IrqDmaStr5 DMA1 Stream5
	IrqDmaStr5 runtime.Interrupt = 16

	// IrqDmaStr6 DMA1 Stream6
	IrqDmaStr6 runtime.Interrupt = 17

	// IrqAdc12 ADC1 and ADC2
	IrqAdc12 runtime.Interrupt = 18

	// IrqFdcan1It0 FDCAN1 Interrupt 0
	IrqFdcan1It0 runtime.Interrupt = 19

	// IrqFdcan2It0 FDCAN2 Interrupt 0
	IrqFdcan2It0 runtime.Interrupt = 20

	// IrqFdcan1It1 FDCAN1 Interrupt 1
	IrqFdcan1It1 runtime.Interrupt = 21

	// IrqFdcan2It1 FDCAN2 Interrupt 1
	IrqFdcan2It1 runtime.Interrupt = 22

	// IrqExti95 EXTI Line[9:5] interrupts
	IrqExti95 runtime.Interrupt = 23

	// IrqTim1Brk TIM1 break interrupt
	IrqTim1Brk runtime.Interrupt = 24

	// IrqTim1Up TIM1 update interrupt
	IrqTim1Up runtime.Interrupt = 25

	// IrqTim1TrgCom TIM1 trigger and commutation
	IrqTim1TrgCom runtime.Interrupt = 26

	// IrqTimCc TIM1 capture / compare
	IrqTimCc runtime.Interrupt = 27

	// IrqTim2 TIM2 global interrupt
	IrqTim2 runtime.Interrupt = 28

	// IrqTim3 TIM3 global interrupt
	IrqTim3 runtime.Interrupt = 29

	// IrqTim4 TIM4 global interrupt
	IrqTim4 runtime.Interrupt = 30

	// IrqI2c1Ev I2C1 event interrupt
	IrqI2c1Ev runtime.Interrupt = 31

	// IrqI2c1Er I2C1 error interrupt
	IrqI2c1Er runtime.Interrupt = 32

	// IrqI2c2Ev I2C2 event interrupt
	IrqI2c2Ev runtime.Interrupt = 33

	// IrqI2c2Er I2C2 error interrupt
	IrqI2c2Er runtime.Interrupt = 34

	// IrqSpi1 SPI1 global interrupt
	IrqSpi1 runtime.Interrupt = 35

	// IrqSpi2 SPI2 global interrupt
	IrqSpi2 runtime.Interrupt = 36

	// IrqUsart1 USART1 global interrupt
	IrqUsart1 runtime.Interrupt = 37

	// IrqUsart2 USART2 global interrupt
	IrqUsart2 runtime.Interrupt = 38

	// IrqUsart3 USART3 global interrupt
	IrqUsart3 runtime.Interrupt = 39

	// IrqExti1510 EXTI Line[15:10] interrupts
	IrqExti1510 runtime.Interrupt = 40

	// IrqRtcAlarm RTC alarms (A and B)
	IrqRtcAlarm runtime.Interrupt = 41

	// IrqTim8BrkTim12 TIM8 and 12 break global
	IrqTim8BrkTim12 runtime.Interrupt = 43

	// IrqTim8UpTim13 TIM8 and 13 update global
	IrqTim8UpTim13 runtime.Interrupt = 44

	// IrqTim8TrgComTim14 TIM8 and 14 trigger /commutation and global
	IrqTim8TrgComTim14 runtime.Interrupt = 45

	// IrqTim8Cc TIM8 capture / compare
	IrqTim8Cc runtime.Interrupt = 46

	// IrqDma1Str7 DMA1 Stream7
	IrqDma1Str7 runtime.Interrupt = 47

	// IrqFmc FMC global interrupt
	IrqFmc runtime.Interrupt = 48

	// IrqSdmmc1 SDMMC global interrupt
	IrqSdmmc1 runtime.Interrupt = 49

	// IrqTim5 TIM5 global interrupt
	IrqTim5 runtime.Interrupt = 50

	// IrqSpi3 SPI3 global interrupt
	IrqSpi3 runtime.Interrupt = 51

	// IrqUart4 UART4 global interrupt
	IrqUart4 runtime.Interrupt = 52

	// IrqUart5 UART5 global interrupt
	IrqUart5 runtime.Interrupt = 53

	// IrqTim6Dac TIM6 global interrupt
	IrqTim6Dac runtime.Interrupt = 54

	// IrqTim7 TIM7 global interrupt
	IrqTim7 runtime.Interrupt = 55

	// IrqDma2Str0 DMA2 Stream0 interrupt
	IrqDma2Str0 runtime.Interrupt = 56

	// IrqDma2Str1 DMA2 Stream1 interrupt
	IrqDma2Str1 runtime.Interrupt = 57

	// IrqDma2Str2 DMA2 Stream2 interrupt
	IrqDma2Str2 runtime.Interrupt = 58

	// IrqDma2Str3 DMA2 Stream3 interrupt
	IrqDma2Str3 runtime.Interrupt = 59

	// IrqDma2Str4 DMA2 Stream4 interrupt
	IrqDma2Str4 runtime.Interrupt = 60

	// IrqEth Ethernet global interrupt
	IrqEth runtime.Interrupt = 61

	// IrqEthWkup Ethernet wakeup through EXTI
	IrqEthWkup runtime.Interrupt = 62

	// IrqFdcanCal CAN2TX interrupts
	IrqFdcanCal runtime.Interrupt = 63

	// IrqCm4SevIt Arm Cortex-M4 Send even interrupt
	IrqCm4SevIt runtime.Interrupt = 65

	// IrqDma2Str5 DMA2 Stream5 interrupt
	IrqDma2Str5 runtime.Interrupt = 68

	// IrqDma2Str6 DMA2 Stream6 interrupt
	IrqDma2Str6 runtime.Interrupt = 69

	// IrqDma2Str7 DMA2 Stream7 interrupt
	IrqDma2Str7 runtime.Interrupt = 70

	// IrqUsart6 USART6 global interrupt
	IrqUsart6 runtime.Interrupt = 71

	// IrqI2c3Ev I2C3 event interrupt
	IrqI2c3Ev runtime.Interrupt = 72

	// IrqI2c3Er I2C3 error interrupt
	IrqI2c3Er runtime.Interrupt = 73

	// IrqOtgHsEp1Out OTG_HS out global interrupt
	IrqOtgHsEp1Out runtime.Interrupt = 74

	// IrqOtgHsEp1In OTG_HS in global interrupt
	IrqOtgHsEp1In runtime.Interrupt = 75

	// IrqOtgHsWkup OTG_HS wakeup interrupt
	IrqOtgHsWkup runtime.Interrupt = 76

	// IrqOtgHs OTG_HS global interrupt
	IrqOtgHs runtime.Interrupt = 77

	// IrqDcmi DCMI global interrupt
	IrqDcmi runtime.Interrupt = 78

	// IrqCryp CRYP global interrupt
	IrqCryp runtime.Interrupt = 79

	// IrqHashRng HASH and RNG
	IrqHashRng runtime.Interrupt = 80

	// IrqFpu CPU2
	IrqFpu runtime.Interrupt = 81

	// IrqUart7 UART7 global interrupt
	IrqUart7 runtime.Interrupt = 82

	// IrqUart8 UART8 global interrupt
	IrqUart8 runtime.Interrupt = 83

	// IrqSpi4 SPI4 global interrupt
	IrqSpi4 runtime.Interrupt = 84

	// IrqSpi5 SPI5 global interrupt
	IrqSpi5 runtime.Interrupt = 85

	// IrqSpi6 SPI6 global interrupt
	IrqSpi6 runtime.Interrupt = 86

	// IrqSai1 SAI1 global interrupt
	IrqSai1 runtime.Interrupt = 87

	// IrqLtdc LCD-TFT global interrupt
	IrqLtdc runtime.Interrupt = 88

	// IrqLtdcEr LCD-TFT error interrupt
	IrqLtdcEr runtime.Interrupt = 89

	// IrqDma2d DMA2D global interrupt
	IrqDma2d runtime.Interrupt = 90

	// IrqSai2 SAI2 global interrupt
	IrqSai2 runtime.Interrupt = 91

	// IrqQuadspi QuadSPI global interrupt
	IrqQuadspi runtime.Interrupt = 92

	// IrqLptim1 LPTIM1 global interrupt
	IrqLptim1 runtime.Interrupt = 93

	// IrqCec HDMI-CEC global interrupt
	IrqCec runtime.Interrupt = 94

	// IrqI2c4Ev I2C4 event interrupt
	IrqI2c4Ev runtime.Interrupt = 95

	// IrqI2c4Er I2C4 error interrupt
	IrqI2c4Er runtime.Interrupt = 96

	// IrqSpdif SPDIFRX global interrupt
	IrqSpdif runtime.Interrupt = 97

	// IrqOtgFsEp1Out OTG_FS out global interrupt
	IrqOtgFsEp1Out runtime.Interrupt = 98

	// IrqOtgFsEp1In OTG_FS in global interrupt
	IrqOtgFsEp1In runtime.Interrupt = 99

	// IrqOtgFsWkup OTG_FS wakeup
	IrqOtgFsWkup runtime.Interrupt = 100

	// IrqOtgFs OTG_FS global interrupt
	IrqOtgFs runtime.Interrupt = 101

	// IrqDmamux1Ov DMAMUX1 overrun interrupt
	IrqDmamux1Ov runtime.Interrupt = 102

	// IrqHrtim1Mst HRTIM1 master timer interrupt
	IrqHrtim1Mst runtime.Interrupt = 103

	// IrqHrtim1Tima HRTIM1 timer A interrupt
	IrqHrtim1Tima runtime.Interrupt = 104

	// IrqHrtimTimb HRTIM1 timer B interrupt
	IrqHrtimTimb runtime.Interrupt = 105

	// IrqHrtim1Timc HRTIM1 timer C interrupt
	IrqHrtim1Timc runtime.Interrupt = 106

	// IrqHrtim1Timd HRTIM1 timer D interrupt
	IrqHrtim1Timd runtime.Interrupt = 107

	// IrqHrtimTime HRTIM1 timer E interrupt
	IrqHrtimTime runtime.Interrupt = 108

	// IrqHrtim1Flt HRTIM1 fault interrupt
	IrqHrtim1Flt runtime.Interrupt = 109

	// IrqDfsdm1Flt0 DFSDM1 filter 0 interrupt
	IrqDfsdm1Flt0 runtime.Interrupt = 110

	// IrqDfsdm1Flt1 DFSDM1 filter 1 interrupt
	IrqDfsdm1Flt1 runtime.Interrupt = 111

	// IrqDfsdm1Flt2 DFSDM1 filter 2 interrupt
	IrqDfsdm1Flt2 runtime.Interrupt = 112

	// IrqDfsdm1Flt3 DFSDM1 filter 3 interrupt
	IrqDfsdm1Flt3 runtime.Interrupt = 113

	// IrqSai3 SAI3 global interrupt
	IrqSai3 runtime.Interrupt = 114

	// IrqSwpmi1 SWPMI global interrupt
	IrqSwpmi1 runtime.Interrupt = 115

	// IrqTim15 TIM15 global interrupt
	IrqTim15 runtime.Interrupt = 116

	// IrqTim16 TIM16 global interrupt
	IrqTim16 runtime.Interrupt = 117

	// IrqTim17 TIM17 global interrupt
	IrqTim17 runtime.Interrupt = 118

	// IrqMdiosWkup MDIOS wakeup
	IrqMdiosWkup runtime.Interrupt = 119

	// IrqMdios MDIOS global interrupt
	IrqMdios runtime.Interrupt = 120

	// IrqJpeg JPEG global interrupt
	IrqJpeg runtime.Interrupt = 121

	// IrqMdma MDMA
	IrqMdma runtime.Interrupt = 122

	// IrqSdmmc2 SDMMC2 global interrupt
	IrqSdmmc2 runtime.Interrupt = 124

	// IrqHsem0 HSEM global interrupt 1
	IrqHsem0 runtime.Interrupt = 125

	// IrqAdc3 ADC3 global interrupt
	IrqAdc3 runtime.Interrupt = 127

	// IrqDmamux2Ovr DMAMUX2 overrun interrupt
	IrqDmamux2Ovr runtime.Interrupt = 128

	// IrqBdmaCh1 BDMA channel 1 interrupt
	IrqBdmaCh1 runtime.Interrupt = 129

	// IrqBdmaCh2 BDMA channel 2 interrupt
	IrqBdmaCh2 runtime.Interrupt = 130

	// IrqBdmaCh3 BDMA channel 3 interrupt
	IrqBdmaCh3 runtime.Interrupt = 131

	// IrqBdmaCh4 BDMA channel 4 interrupt
	IrqBdmaCh4 runtime.Interrupt = 132

	// IrqBdmaCh5 BDMA channel 5 interrupt
	IrqBdmaCh5 runtime.Interrupt = 133

	// IrqBdmaCh6 BDMA channel 6 interrupt
	IrqBdmaCh6 runtime.Interrupt = 134

	// IrqBdmaCh7 BDMA channel 7 interrupt
	IrqBdmaCh7 runtime.Interrupt = 135

	// IrqBdmaCh8 BDMA channel 8 interrupt
	IrqBdmaCh8 runtime.Interrupt = 136

	// IrqComp COMP1 and COMP2
	IrqComp runtime.Interrupt = 137

	// IrqLptim2 LPTIM2 timer interrupt
	IrqLptim2 runtime.Interrupt = 138

	// IrqLptim3 LPTIM3 timer interrupt
	IrqLptim3 runtime.Interrupt = 139

	// IrqLptim4 LPTIM4 timer interrupt
	IrqLptim4 runtime.Interrupt = 140

	// IrqLptim5 LPTIM5 timer interrupt
	IrqLptim5 runtime.Interrupt = 141

	// IrqLpuart LPUART global interrupt
	IrqLpuart runtime.Interrupt = 142

	// IrqWwdg2Rst Window Watchdog interrupt
	IrqWwdg2Rst runtime.Interrupt = 143

	// IrqCrs Clock Recovery System globa
	IrqCrs runtime.Interrupt = 144

	// IrqSai4 SAI4 global interrupt
	IrqSai4 runtime.Interrupt = 146

	// IrqHoldCore CPU1 hold
	IrqHoldCore runtime.Interrupt = 148

	// IrqWkup WKUP1 to WKUP6 pins
	IrqWkup runtime.Interrupt = 149
)
