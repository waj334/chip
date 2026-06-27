package main

import (
	"time"

	"pkg.si-go.dev/chip/arm/cortexm"
	_ "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/timer"
)

const (
	LEDR = pin.PI12
	LEDB = pin.PE3
	LEDG = pin.PJ13
	TIM2 = timer.TIM2
)

func init() {
	// Prevent SysTick from driving timers.
	cortexm.SysTickCanWake = false

	hal.ConfigureClocks()

	err := TIM2.Configure(timer.Config{Enable: true})
	if err != nil {
		panic(err)
	}
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
