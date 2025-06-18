package main

import (
	_ "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/timer"
	"pkg.si-go.dev/chip/arm/cortexm/runtime"
	"time"
)

const (
	LEDR = pin.PI12
	LEDB = pin.PE3
	LEDG = pin.PJ13
	TIM2 = timer.TIM2

	timescale = uint64(time.Microsecond)
)

//sigo:export wake runtime.wake
func wake(t uint64)

func alarm(t uint64) {
	wake(t)
}

//sigo:export nanotime runtime.nanotime
func nanotime() uint64 {
	// The timer resolution is 1uS per tick.
	return TIM2.Tick() * timescale
}

//sigo:export addsleep runtime.addsleep
func addsleep(deadline uint64) bool {
	return !TIM2.SetAlarm(deadline/timescale, alarm)
}

func init() {
	// Prevent SysTick from driving timers.
	runtime.SysTickCanWake = false

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
