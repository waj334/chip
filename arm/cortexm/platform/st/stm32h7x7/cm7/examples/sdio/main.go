package main

import (
	"time"

	_ "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/sdio"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/timer"
	"pkg.si-go.dev/chip/arm/cortexm/runtime"
)

const (
	LEDR  = pin.PI12
	LEDB  = pin.PE3
	LEDG  = pin.PJ13
	TIM2  = timer.TIM2
	SDIO1 = sdio.SDIO1

	BLEON = pin.PA10

	WIFION       = pin.PB10
	WIFIHOSTWAKE = pin.PI8

	SDIO_D1  = pin.PC8
	SDIO_D2  = pin.PC9
	SDIO_D3  = pin.PC10
	SDIO_D4  = pin.PC11
	SDIO_CLK = pin.PC12
	SDIO_CMD = pin.PD2

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
func addsleep(deadline uint64) {
	TIM2.SetAlarm(deadline/timescale, alarm)
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

	LEDR.High()
	LEDB.Low()
	LEDG.High()

	SDIO_D1.SetSpeedMode(pin.VeryHighSpeed)
	SDIO_D2.SetSpeedMode(pin.VeryHighSpeed)
	SDIO_D3.SetSpeedMode(pin.VeryHighSpeed)
	SDIO_D4.SetSpeedMode(pin.VeryHighSpeed)
	SDIO_CLK.SetSpeedMode(pin.VeryHighSpeed)
	SDIO_CMD.SetSpeedMode(pin.VeryHighSpeed)

	WIFIHOSTWAKE.SetMode(pin.Input)
	WIFION.SetMode(pin.Output)
	WIFION.High()

	BLEON.SetMode(pin.Output)
	BLEON.High()

	time.Sleep(time.Millisecond * 250)

	err := SDIO1.Configure(sdio.Config{
		Enable: true,
		CK:     SDIO_CLK,
		Dn: [8]pin.Pin{
			SDIO_D1,
			SDIO_D2,
			SDIO_D3,
			SDIO_D4,
		},
		CMD: SDIO_CMD,
	})

	if err != nil {
		errorState()
		busyLoop()
	}

	var resp sdio.Response
	ready := false

	// Send CMD0 to reset the card
	if _, err = SDIO1.SendCommand(sdio.Command{Index: sdio.CMD0, Argument: 0}); err != nil {
		errorState()
		busyLoop()
	}

	for retry := 0; retry < 1000; retry++ {
		// Send CMD5 to get the ready status of the card.
		if resp, err = SDIO1.SendCommand(sdio.Command{Index: sdio.CMD5, Argument: 0x00FF8000}); err != nil {
			time.Sleep(time.Millisecond * 5)
			continue
		} else {
			if resp[0]>>31 == 0 {
				// The card is not ready. Try again.
				time.Sleep(time.Millisecond * 5)
				continue
			}
			ready = true
		}
		break
	}

	if !ready {
		errorState()
		busyLoop()
	}

	// Send CMD3 to get the address of the card.
	resp, err = SDIO1.SendCommand(sdio.Command{Index: sdio.CMD3, Argument: 0})
	if err != nil {
		errorState()
		busyLoop()
	}

	// Send CMD7 with the returned RCA to select the card.
	_, err = SDIO1.SendCommand(sdio.Command{Index: sdio.CMD7, Argument: resp[0] & 0xFFFF0000})
	if err != nil {
		errorState()
		busyLoop()
	}

	goodState()
	busyLoop()
}

func errorState() {
	LEDR.Low()
	LEDB.High()
	LEDG.High()
}

func goodState() {
	LEDR.High()
	LEDB.High()
	LEDG.Low()
}

func busyLoop() {
	select {}
}
