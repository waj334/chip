package main

import (
	_ "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/uart"
)

var (
	UART = uart.UART1
)

func main() {
	hal.ConfigureClocks()

	// Configure UART
	UART.Configure(uart.Config{
		Enable: true,
		TX:     pin.PA9,
		RX:     pin.PB7,
		// FrameFormat:     uart.UsartFrame,
		Baud:            115_200,
		CharacterSize:   8,
		NumStopBits:     1,
		ReceiveEnabled:  true,
		TransmitEnabled: true,
	})

	UART.WriteString("Hello, World!\n\n")
	UART.WriteString("Any message received will be echoed below:\n")

	b := make([]byte, 32)
	for {
		n, err := UART.Read(b)
		if err != nil {
			panic(err)
		}

		if n > 0 {
			UART.Write(b[:n])
		}
	}
}
