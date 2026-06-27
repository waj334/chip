package cortexm

import (
	"sync/atomic"

	"pkg.si-go.dev/chip/arm/cortexm/reg/scb"
	"pkg.si-go.dev/chip/arm/cortexm/reg/systick"
)

var _tickCount uint32
var SysTickFrequency uint32 = 100_000_000
var SysTickCanWake = true

//sigo:extern wake runtime.wake
func wake(t uint64)

type SysTickSource struct{}

func (s SysTickSource) Now() uint64 {
	// NOTE: The tick count is incremented at a frequency of 1ms
	// Return the tick count in nanoseconds
	return uint64(atomic.LoadUint32(&_tickCount)) * 1_000_000
}

//sigo:export nanotime runtime.nanotime
//sigo:linkage nanotime weak
func nanotime() uint64 {
	state := DisableInterrupts()
	value := uint64(atomic.LoadUint32(&_tickCount)) * 1_000_000
	EnableInterrupts(state)
	return value
}

func initSysTick() {
	systick.Systick.Csr.SetEnable(false)

	// STM32H747 implements 4 of the 8 architectural priority bits.
	const priorityShift = 4

	// Lowest configurable priority.
	scb.Scb.Shpr3.SetPri14(15 << priorityShift)

	// One priority level above PendSV.
	scb.Scb.Shpr3.SetPri15(14 << priorityShift)

	UpdateSysTickFrequency(SysTickFrequency)

	systick.Systick.Csr.SetTickint(true)
	systick.Systick.Csr.SetClksource(true)
	systick.Systick.Csr.SetEnable(true)

	for !systick.Systick.Csr.GetEnable() {
	}
}

func UpdateSysTickFrequency(value uint32) {
	SysTickFrequency = value
	systick.Systick.Rvr.SetReload(SysTickFrequency / 1000)
}

//sigo:interrupt sysTickHandler SysTickHandler
func sysTickHandler() {
	// Atomically increment the tick counter
	atomic.AddUint32(&_tickCount, 1)

	if SysTickCanWake {
		wake(nanotime())
	}

	// Trigger a PendSV interrupt to run the scheduler
	gosched()
}

//go:export gosched runtime.gosched
func gosched() {
	// Set the PendSV flag
	scb.Scb.Icsr.SetPendsvset(true)
}
