package main

import (
	_ "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"time"
)

const (
	LEDR = pin.PI12
	LEDB = pin.PE3
	LEDG = pin.PJ13
)

func init() {
	hal.ConfigureClocks()
}

func main() {
	LEDR.SetMode(pin.Output)
	LEDB.SetMode(pin.Output)
	LEDG.SetMode(pin.Output)
	for {
		LEDR.Toggle()
		time.Sleep(time.Millisecond * 500)
		LEDB.Toggle()
		time.Sleep(time.Millisecond * 500)
		LEDG.Toggle()
		time.Sleep(time.Millisecond * 500)
	}
}
