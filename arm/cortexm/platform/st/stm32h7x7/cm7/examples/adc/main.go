package main

import (
	_ "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/adc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
)

const (
	LEDR = pin.PI12
	LEDB = pin.PE3
	LEDG = pin.PJ13

	A0 = pin.PC4
)

func init() {
	hal.ConfigureClocks()
}

func main() {
	LEDR.SetMode(pin.Output)
	LEDB.SetMode(pin.Output)
	LEDG.SetMode(pin.Output)
	A0.SetMode(pin.Analog)

	LEDR.Low()
	LEDG.Low()
	LEDB.Low()

	adc.ADC1.Configure(adc.Config{
		Enable:     true,
		Resolution: adc.Resolution16Bit,
	})

	for {
		value, err := A0.Value()
		if err != nil {
			panic(err)
		}

		switch {
		case value > 0xAAAA:
			LEDR.Low()
			LEDG.Low()
			LEDB.High()
		case value > 0x5555:
			LEDR.Low()
			LEDG.High()
			LEDB.Low()
		case value > 2000:
			LEDR.High()
			LEDG.Low()
			LEDB.Low()
		default:
			LEDR.Low()
			LEDG.Low()
			LEDB.Low()
		}
	}
}
