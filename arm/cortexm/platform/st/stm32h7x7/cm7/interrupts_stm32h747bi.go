//go:build stm32h7x7 && stm32h747bi

package stm32h7x7

import "pkg.si-go.dev/chip/arm/cortexm"

const (
	// IrqReset Reset Vector, invoked on Power up and warm reset.
	IrqReset cortexm.Interrupt = -15

	// IrqNmi Non maskable Interrupt, cannot be stopped or preempted.
	IrqNmi cortexm.Interrupt = -14

	// IrqHardfault Hard Fault, all classes of Fault.
	IrqHardfault cortexm.Interrupt = -13

	// IrqMemoryManage Memory Management, MPU mismatch, including Access Violation and No Match.
	IrqMemoryManage cortexm.Interrupt = -12

	// IrqBusFault Bus Fault, Pre-Fetch-, Memory Access Fault, other address/memory related Fault.
	IrqBusFault cortexm.Interrupt = -11

	// IrqUsageFault Usage Fault, i.e. Undef Instruction, Illegal State Transition
	IrqUsageFault cortexm.Interrupt = -10

	// IrqSvCall System Service Call via SVC instruction.
	IrqSvCall cortexm.Interrupt = -5

	// IrqDebugMonitor Debug Monitor.
	IrqDebugMonitor cortexm.Interrupt = -4

	// IrqPendSv Pendable request for system service.
	IrqPendSv cortexm.Interrupt = -2

	// IrqSysTick System Tick Timer.
	IrqSysTick cortexm.Interrupt = -1

	// IrqWwdg1 Window Watchdog interrupt
	IrqWwdg1 cortexm.Interrupt = 0

	// IrqPvdPvm PVD through EXTI line
	IrqPvdPvm cortexm.Interrupt = 1

	// IrqRtcTampStampCssLse RTC tamper, timestamp
	IrqRtcTampStampCssLse cortexm.Interrupt = 2

	// IrqRtcWkup RTC Wakeup interrupt
	IrqRtcWkup cortexm.Interrupt = 3

	// IrqFlash Flash memory
	IrqFlash cortexm.Interrupt = 4

	// IrqRcc RCC global interrupt
	IrqRcc cortexm.Interrupt = 5

	// IrqExti0 EXTI Line 0 interrupt
	IrqExti0 cortexm.Interrupt = 6

	// IrqExti1 EXTI Line 1 interrupt
	IrqExti1 cortexm.Interrupt = 7

	// IrqExti2 EXTI Line 2 interrupt
	IrqExti2 cortexm.Interrupt = 8

	// IrqExti3 EXTI Line 3interrupt
	IrqExti3 cortexm.Interrupt = 9

	// IrqExti4 EXTI Line 4interrupt
	IrqExti4 cortexm.Interrupt = 10

	// IrqDmaStr0 DMA1 Stream0
	IrqDmaStr0 cortexm.Interrupt = 11

	// IrqDmaStr1 DMA1 Stream1
	IrqDmaStr1 cortexm.Interrupt = 12

	// IrqDmaStr2 DMA1 Stream2
	IrqDmaStr2 cortexm.Interrupt = 13

	// IrqDmaStr3 DMA1 Stream3
	IrqDmaStr3 cortexm.Interrupt = 14

	// IrqDmaStr4 DMA1 Stream4
	IrqDmaStr4 cortexm.Interrupt = 15

	// IrqDmaStr5 DMA1 Stream5
	IrqDmaStr5 cortexm.Interrupt = 16

	// IrqDmaStr6 DMA1 Stream6
	IrqDmaStr6 cortexm.Interrupt = 17

	// IrqAdc12 ADC1 and ADC2
	IrqAdc12 cortexm.Interrupt = 18

	// IrqFdcan1It0 FDCAN1 Interrupt 0
	IrqFdcan1It0 cortexm.Interrupt = 19

	// IrqFdcan2It0 FDCAN2 Interrupt 0
	IrqFdcan2It0 cortexm.Interrupt = 20

	// IrqFdcan1It1 FDCAN1 Interrupt 1
	IrqFdcan1It1 cortexm.Interrupt = 21

	// IrqFdcan2It1 FDCAN2 Interrupt 1
	IrqFdcan2It1 cortexm.Interrupt = 22

	// IrqExti95 EXTI Line[9:5] interrupts
	IrqExti95 cortexm.Interrupt = 23

	// IrqTim1Brk TIM1 break interrupt
	IrqTim1Brk cortexm.Interrupt = 24

	// IrqTim1Up TIM1 update interrupt
	IrqTim1Up cortexm.Interrupt = 25

	// IrqTim1TrgCom TIM1 trigger and commutation
	IrqTim1TrgCom cortexm.Interrupt = 26

	// IrqTimCc TIM1 capture / compare
	IrqTimCc cortexm.Interrupt = 27

	// IrqTim2 TIM2 global interrupt
	IrqTim2 cortexm.Interrupt = 28

	// IrqTim3 TIM3 global interrupt
	IrqTim3 cortexm.Interrupt = 29

	// IrqTim4 TIM4 global interrupt
	IrqTim4 cortexm.Interrupt = 30

	// IrqI2c1Ev I2C1 event interrupt
	IrqI2c1Ev cortexm.Interrupt = 31

	// IrqI2c1Er I2C1 error interrupt
	IrqI2c1Er cortexm.Interrupt = 32

	// IrqI2c2Ev I2C2 event interrupt
	IrqI2c2Ev cortexm.Interrupt = 33

	// IrqI2c2Er I2C2 error interrupt
	IrqI2c2Er cortexm.Interrupt = 34

	// IrqSpi1 SPI1 global interrupt
	IrqSpi1 cortexm.Interrupt = 35

	// IrqSpi2 SPI2 global interrupt
	IrqSpi2 cortexm.Interrupt = 36

	// IrqUsart1 USART1 global interrupt
	IrqUsart1 cortexm.Interrupt = 37

	// IrqUsart2 USART2 global interrupt
	IrqUsart2 cortexm.Interrupt = 38

	// IrqUsart3 USART3 global interrupt
	IrqUsart3 cortexm.Interrupt = 39

	// IrqExti1510 EXTI Line[15:10] interrupts
	IrqExti1510 cortexm.Interrupt = 40

	// IrqRtcAlarm RTC alarms (A and B)
	IrqRtcAlarm cortexm.Interrupt = 41

	// IrqTim8BrkTim12 TIM8 and 12 break global
	IrqTim8BrkTim12 cortexm.Interrupt = 43

	// IrqTim8UpTim13 TIM8 and 13 update global
	IrqTim8UpTim13 cortexm.Interrupt = 44

	// IrqTim8TrgComTim14 TIM8 and 14 trigger /commutation and global
	IrqTim8TrgComTim14 cortexm.Interrupt = 45

	// IrqTim8Cc TIM8 capture / compare
	IrqTim8Cc cortexm.Interrupt = 46

	// IrqDma1Str7 DMA1 Stream7
	IrqDma1Str7 cortexm.Interrupt = 47

	// IrqFmc FMC global interrupt
	IrqFmc cortexm.Interrupt = 48

	// IrqSdmmc1 SDMMC global interrupt
	IrqSdmmc1 cortexm.Interrupt = 49

	// IrqTim5 TIM5 global interrupt
	IrqTim5 cortexm.Interrupt = 50

	// IrqSpi3 SPI3 global interrupt
	IrqSpi3 cortexm.Interrupt = 51

	// IrqUart4 UART4 global interrupt
	IrqUart4 cortexm.Interrupt = 52

	// IrqUart5 UART5 global interrupt
	IrqUart5 cortexm.Interrupt = 53

	// IrqTim6Dac TIM6 global interrupt
	IrqTim6Dac cortexm.Interrupt = 54

	// IrqTim7 TIM7 global interrupt
	IrqTim7 cortexm.Interrupt = 55

	// IrqDma2Str0 DMA2 Stream0 interrupt
	IrqDma2Str0 cortexm.Interrupt = 56

	// IrqDma2Str1 DMA2 Stream1 interrupt
	IrqDma2Str1 cortexm.Interrupt = 57

	// IrqDma2Str2 DMA2 Stream2 interrupt
	IrqDma2Str2 cortexm.Interrupt = 58

	// IrqDma2Str3 DMA2 Stream3 interrupt
	IrqDma2Str3 cortexm.Interrupt = 59

	// IrqDma2Str4 DMA2 Stream4 interrupt
	IrqDma2Str4 cortexm.Interrupt = 60

	// IrqEth Ethernet global interrupt
	IrqEth cortexm.Interrupt = 61

	// IrqEthWkup Ethernet wakeup through EXTI
	IrqEthWkup cortexm.Interrupt = 62

	// IrqFdcanCal CAN2TX interrupts
	IrqFdcanCal cortexm.Interrupt = 63

	// IrqCm4SevIt Arm Cortex-M4 Send even interrupt
	IrqCm4SevIt cortexm.Interrupt = 65

	// IrqDma2Str5 DMA2 Stream5 interrupt
	IrqDma2Str5 cortexm.Interrupt = 68

	// IrqDma2Str6 DMA2 Stream6 interrupt
	IrqDma2Str6 cortexm.Interrupt = 69

	// IrqDma2Str7 DMA2 Stream7 interrupt
	IrqDma2Str7 cortexm.Interrupt = 70

	// IrqUsart6 USART6 global interrupt
	IrqUsart6 cortexm.Interrupt = 71

	// IrqI2c3Ev I2C3 event interrupt
	IrqI2c3Ev cortexm.Interrupt = 72

	// IrqI2c3Er I2C3 error interrupt
	IrqI2c3Er cortexm.Interrupt = 73

	// IrqOtgHsEp1Out OTG_HS out global interrupt
	IrqOtgHsEp1Out cortexm.Interrupt = 74

	// IrqOtgHsEp1In OTG_HS in global interrupt
	IrqOtgHsEp1In cortexm.Interrupt = 75

	// IrqOtgHsWkup OTG_HS wakeup interrupt
	IrqOtgHsWkup cortexm.Interrupt = 76

	// IrqOtgHs OTG_HS global interrupt
	IrqOtgHs cortexm.Interrupt = 77

	// IrqDcmi DCMI global interrupt
	IrqDcmi cortexm.Interrupt = 78

	// IrqCryp CRYP global interrupt
	IrqCryp cortexm.Interrupt = 79

	// IrqHashRng HASH and RNG
	IrqHashRng cortexm.Interrupt = 80

	// IrqFpu CPU2
	IrqFpu cortexm.Interrupt = 81

	// IrqUart7 UART7 global interrupt
	IrqUart7 cortexm.Interrupt = 82

	// IrqUart8 UART8 global interrupt
	IrqUart8 cortexm.Interrupt = 83

	// IrqSpi4 SPI4 global interrupt
	IrqSpi4 cortexm.Interrupt = 84

	// IrqSpi5 SPI5 global interrupt
	IrqSpi5 cortexm.Interrupt = 85

	// IrqSpi6 SPI6 global interrupt
	IrqSpi6 cortexm.Interrupt = 86

	// IrqSai1 SAI1 global interrupt
	IrqSai1 cortexm.Interrupt = 87

	// IrqLtdc LCD-TFT global interrupt
	IrqLtdc cortexm.Interrupt = 88

	// IrqLtdcEr LCD-TFT error interrupt
	IrqLtdcEr cortexm.Interrupt = 89

	// IrqDma2d DMA2D global interrupt
	IrqDma2d cortexm.Interrupt = 90

	// IrqSai2 SAI2 global interrupt
	IrqSai2 cortexm.Interrupt = 91

	// IrqQuadspi QuadSPI global interrupt
	IrqQuadspi cortexm.Interrupt = 92

	// IrqLptim1 LPTIM1 global interrupt
	IrqLptim1 cortexm.Interrupt = 93

	// IrqCec HDMI-CEC global interrupt
	IrqCec cortexm.Interrupt = 94

	// IrqI2c4Ev I2C4 event interrupt
	IrqI2c4Ev cortexm.Interrupt = 95

	// IrqI2c4Er I2C4 error interrupt
	IrqI2c4Er cortexm.Interrupt = 96

	// IrqSpdif SPDIFRX global interrupt
	IrqSpdif cortexm.Interrupt = 97

	// IrqOtgFsEp1Out OTG_FS out global interrupt
	IrqOtgFsEp1Out cortexm.Interrupt = 98

	// IrqOtgFsEp1In OTG_FS in global interrupt
	IrqOtgFsEp1In cortexm.Interrupt = 99

	// IrqOtgFsWkup OTG_FS wakeup
	IrqOtgFsWkup cortexm.Interrupt = 100

	// IrqOtgFs OTG_FS global interrupt
	IrqOtgFs cortexm.Interrupt = 101

	// IrqDmamux1Ov DMAMUX1 overrun interrupt
	IrqDmamux1Ov cortexm.Interrupt = 102

	// IrqHrtim1Mst HRTIM1 master timer interrupt
	IrqHrtim1Mst cortexm.Interrupt = 103

	// IrqHrtim1Tima HRTIM1 timer A interrupt
	IrqHrtim1Tima cortexm.Interrupt = 104

	// IrqHrtimTimb HRTIM1 timer B interrupt
	IrqHrtimTimb cortexm.Interrupt = 105

	// IrqHrtim1Timc HRTIM1 timer C interrupt
	IrqHrtim1Timc cortexm.Interrupt = 106

	// IrqHrtim1Timd HRTIM1 timer D interrupt
	IrqHrtim1Timd cortexm.Interrupt = 107

	// IrqHrtimTime HRTIM1 timer E interrupt
	IrqHrtimTime cortexm.Interrupt = 108

	// IrqHrtim1Flt HRTIM1 fault interrupt
	IrqHrtim1Flt cortexm.Interrupt = 109

	// IrqDfsdm1Flt0 DFSDM1 filter 0 interrupt
	IrqDfsdm1Flt0 cortexm.Interrupt = 110

	// IrqDfsdm1Flt1 DFSDM1 filter 1 interrupt
	IrqDfsdm1Flt1 cortexm.Interrupt = 111

	// IrqDfsdm1Flt2 DFSDM1 filter 2 interrupt
	IrqDfsdm1Flt2 cortexm.Interrupt = 112

	// IrqDfsdm1Flt3 DFSDM1 filter 3 interrupt
	IrqDfsdm1Flt3 cortexm.Interrupt = 113

	// IrqSai3 SAI3 global interrupt
	IrqSai3 cortexm.Interrupt = 114

	// IrqSwpmi1 SWPMI global interrupt
	IrqSwpmi1 cortexm.Interrupt = 115

	// IrqTim15 TIM15 global interrupt
	IrqTim15 cortexm.Interrupt = 116

	// IrqTim16 TIM16 global interrupt
	IrqTim16 cortexm.Interrupt = 117

	// IrqTim17 TIM17 global interrupt
	IrqTim17 cortexm.Interrupt = 118

	// IrqMdiosWkup MDIOS wakeup
	IrqMdiosWkup cortexm.Interrupt = 119

	// IrqMdios MDIOS global interrupt
	IrqMdios cortexm.Interrupt = 120

	// IrqJpeg JPEG global interrupt
	IrqJpeg cortexm.Interrupt = 121

	// IrqMdma MDMA
	IrqMdma cortexm.Interrupt = 122

	// IrqSdmmc2 SDMMC2 global interrupt
	IrqSdmmc2 cortexm.Interrupt = 124

	// IrqHsem0 HSEM global interrupt 1
	IrqHsem0 cortexm.Interrupt = 125

	// IrqAdc3 ADC3 global interrupt
	IrqAdc3 cortexm.Interrupt = 127

	// IrqDmamux2Ovr DMAMUX2 overrun interrupt
	IrqDmamux2Ovr cortexm.Interrupt = 128

	// IrqBdmaCh1 BDMA channel 1 interrupt
	IrqBdmaCh1 cortexm.Interrupt = 129

	// IrqBdmaCh2 BDMA channel 2 interrupt
	IrqBdmaCh2 cortexm.Interrupt = 130

	// IrqBdmaCh3 BDMA channel 3 interrupt
	IrqBdmaCh3 cortexm.Interrupt = 131

	// IrqBdmaCh4 BDMA channel 4 interrupt
	IrqBdmaCh4 cortexm.Interrupt = 132

	// IrqBdmaCh5 BDMA channel 5 interrupt
	IrqBdmaCh5 cortexm.Interrupt = 133

	// IrqBdmaCh6 BDMA channel 6 interrupt
	IrqBdmaCh6 cortexm.Interrupt = 134

	// IrqBdmaCh7 BDMA channel 7 interrupt
	IrqBdmaCh7 cortexm.Interrupt = 135

	// IrqBdmaCh8 BDMA channel 8 interrupt
	IrqBdmaCh8 cortexm.Interrupt = 136

	// IrqComp COMP1 and COMP2
	IrqComp cortexm.Interrupt = 137

	// IrqLptim2 LPTIM2 timer interrupt
	IrqLptim2 cortexm.Interrupt = 138

	// IrqLptim3 LPTIM3 timer interrupt
	IrqLptim3 cortexm.Interrupt = 139

	// IrqLptim4 LPTIM4 timer interrupt
	IrqLptim4 cortexm.Interrupt = 140

	// IrqLptim5 LPTIM5 timer interrupt
	IrqLptim5 cortexm.Interrupt = 141

	// IrqLpuart LPUART global interrupt
	IrqLpuart cortexm.Interrupt = 142

	// IrqWwdg2Rst Window Watchdog interrupt
	IrqWwdg2Rst cortexm.Interrupt = 143

	// IrqCrs Clock Recovery System globa
	IrqCrs cortexm.Interrupt = 144

	// IrqSai4 SAI4 global interrupt
	IrqSai4 cortexm.Interrupt = 146

	// IrqHoldCore CPU1 hold
	IrqHoldCore cortexm.Interrupt = 148

	// IrqWkup WKUP1 to WKUP6 pins
	IrqWkup cortexm.Interrupt = 149
)
