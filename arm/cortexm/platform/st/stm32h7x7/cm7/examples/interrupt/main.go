package main

import (
	_ "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
)

const (
	LEDR = pin.PI12
	LEDB = pin.PE3
	LEDG = pin.PJ13
	D2   = pin.PA3
)

func init() {
	hal.ConfigureClocks()
}

func main() {
	LEDR.SetMode(pin.Output)
	LEDB.SetMode(pin.Output)
	LEDG.SetMode(pin.Output)

	LEDR.Low()
	LEDG.Low()
	LEDB.Low()

	D2.SetMode(pin.Input)
	D2.SetInterrupt(pin.RisingEdge, func(pin pin.Pin) {
		pin.Toggle()
	})

	// Block forever...
	select {}
}
