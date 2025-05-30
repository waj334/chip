package main

import (
	_ "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/adc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/dac"
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

	dac.DAC1.Configure(dac.Config{
		Enable:     true,
		Resolution: dac.Resolution12Bit,
	})

	dac.DAC2.Configure(dac.Config{
		Enable:     true,
		Resolution: dac.Resolution12Bit,
	})

	for {
		// Read the level of the input.
		value, err := A0.Value()
		if err != nil {
			panic(err)
		}

		// Interpolate the 16-bit input value to a 12-bit output value for the DAC.
		v := (value * 0xFFF) / 0xFFFF

		// Write the output value to both DAC1 and DAC2.
		if err := dac.DAC1.Write(v); err != nil {
			panic(err)
		}

		if err := dac.DAC2.Write(v); err != nil {
			panic(err)
		}
	}
}
