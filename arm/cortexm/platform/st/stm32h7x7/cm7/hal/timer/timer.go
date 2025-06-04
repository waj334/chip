package timer

import (
	stm32h7x7 "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	tim "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/tim2_3_4_5"
	"pkg.si-go.dev/chip/arm/cortexm/runtime"
	"time"
)

var base [4]uint64
var alarms [4]*alarmEntry
var alarmQueues [4]*alarmEntry

const (
	TIM2 _tim32 = iota
	TIM3
	TIM4
	TIM5
)

type _tim32 uint8

type AlarmFunc func(t uint64)

type alarmEntry struct {
	deadline uint64
	callback AlarmFunc
	channel  int         // 0 to 3 for CCR1 to CCR4
	next     *alarmEntry // Optional for future queuing
}

type Direction bool

type Config struct {
	Enable bool
}

func (t _tim32) Configure(config Config) error {
	state := runtime.DisableInterrupts()
	_t := tim.Instances[t]

	// Reset and enable/disable the timer in RCC.
	switch t {
	case 0:
		// Disable the IRQ.
		stm32h7x7.IrqTim2.DisableIRQ()

		// Perform reset sequence.
		rcc.Rcc.Apb1lrstr.SetTim2rst(true)
		for !rcc.Rcc.Apb1lrstr.GetTim2rst() {
		}

		rcc.Rcc.Apb1lrstr.SetTim2rst(false)
		for rcc.Rcc.Apb1lrstr.GetTim2rst() {
		}

		// Enable or disable the timer.
		rcc.Rcc.Apb1lenr.SetTim2en(config.Enable)
	case 1:
		stm32h7x7.IrqTim3.DisableIRQ()

		// Perform reset sequence.
		rcc.Rcc.Apb1lrstr.SetTim3rst(true)
		for !rcc.Rcc.Apb1lrstr.GetTim3rst() {
		}

		rcc.Rcc.Apb1lrstr.SetTim3rst(false)
		for rcc.Rcc.Apb1lrstr.GetTim3rst() {
		}

		rcc.Rcc.Apb1lenr.SetTim3en(config.Enable)
	case 2:
		stm32h7x7.IrqTim4.DisableIRQ()

		// Perform reset sequence.
		rcc.Rcc.Apb1lrstr.SetTim4rst(true)
		for !rcc.Rcc.Apb1lrstr.GetTim4rst() {
		}

		rcc.Rcc.Apb1lrstr.SetTim4rst(false)
		for rcc.Rcc.Apb1lrstr.GetTim4rst() {
		}

		rcc.Rcc.Apb1lenr.SetTim4en(config.Enable)
	case 3:
		stm32h7x7.IrqTim5.DisableIRQ()

		// Perform reset sequence.
		rcc.Rcc.Apb1lrstr.SetTim5rst(true)
		for !rcc.Rcc.Apb1lrstr.GetTim5rst() {
		}

		rcc.Rcc.Apb1lrstr.SetTim5rst(false)
		for rcc.Rcc.Apb1lrstr.GetTim5rst() {
		}

		rcc.Rcc.Apb1lenr.SetTim5en(config.Enable)
	}

	if !config.Enable {
		runtime.EnableInterrupts(state)
		return nil
	}

	// Set the prescaler for a 1 MHz timer.
	presc := (hal.Pclk1FrequencyHz * hal.TimerMultiplier(hal.D2ppre1)) / 1_000_000
	_t.Psc.SetPsc(uint16(presc - 1))

	// Immediately apply the PSC value by generating an update event.
	_t.Egr.SetUg(true)

	// Initialize the reload value to max value.
	_t.Arr.SetArrl(0xFFFF)
	_t.Arr.SetArrh(0xFFFF)

	// Initialize the counter to zero.
	_t.Cnt.SetCntl(0)
	_t.Cnt.SetCnth(0)

	// Set the compare value to max.
	_t.Ccr1.SetCcr1l(0xFFFF)
	_t.Ccr1.SetCcr1h(0xFFFF)

	_t.Cr1.SetArpe(true)

	// Enable the interrupt.
	_t.Dier.SetUie(true)

	switch t {
	case 0:
		stm32h7x7.IrqTim2.EnableIRQ()
	case 1:
		stm32h7x7.IrqTim3.EnableIRQ()
	case 2:
		stm32h7x7.IrqTim4.EnableIRQ()
	case 3:
		stm32h7x7.IrqTim5.EnableIRQ()
	}

	// Start the timer.
	_t.Cr1.SetCen(true)
	for !_t.Cr1.GetCen() {
	}

	runtime.EnableInterrupts(state)
	return nil
}

func (t _tim32) Reset(tick uint64) {
	_t := tim.Instances[t]
	_t.Cnt.SetCntl(uint16(tick))
	_t.Cnt.SetCnth(uint16(tick >> 16))

	// Start the timer.
	_t.Cr1.SetCen(true)
	for !_t.Cr1.GetCen() {
	}
}

func (t _tim32) Stop() {
	_t := tim.Instances[t]
	_t.Cr1.SetCen(false)
	for _t.Cr1.GetCen() {
	}
}

func (t _tim32) Tick() uint64 {
	_t := tim.Instances[t]
	b := base[t]
	tick := uint64(uint32(_t.Cnt.GetCntl()) | (uint32(_t.Cnt.GetCnth()) << 16))
	result := b + tick
	return result
}

func (t _tim32) Now() time.Time {
	tick := t.Tick()
	return (time.Time{}).Add(t.Resolution() * time.Duration(tick))
}

func (t _tim32) Resolution() time.Duration {
	return time.Microsecond
}

func (t _tim32) SetAlarm(deadline uint64, fn AlarmFunc) {
	state := runtime.DisableInterrupts()
	now := t.Tick()
	if fn == nil || deadline <= now {
		runtime.EnableInterrupts(state)
		fn(now)
		return
	}

	// Try to assign to an available channel
	for ch := 0; ch < 4; ch++ {
		if alarms[ch] == nil {
			entry := &alarmEntry{deadline: deadline, callback: fn, channel: ch}
			alarms[ch] = entry
			armCompare(t, ch, deadline)
			runtime.EnableInterrupts(state)
			return
		}
	}

	// No channels available. Queue it.
	entry := &alarmEntry{deadline: deadline, callback: fn, channel: -1}
	insertQueuedAlarm(t, entry)
	runtime.EnableInterrupts(state)
}

//sigo:interrupt tim2Handler Tim2Handler
func tim2Handler() {
	interrupt(TIM2)
}

//sigo:interrupt tim3Handler Tim3Handler
func tim3Handler() {
	interrupt(TIM3)
}

//sigo:interrupt tim4Handler Tim4Handler
func tim4Handler() {
	interrupt(TIM4)
}

//sigo:interrupt tim5Handler Tim5Handler
func tim5Handler() {
	interrupt(TIM5)
}

func interrupt(instance _tim32) {
	state := runtime.DisableInterrupts()
	_t := tim.Instances[instance]
	alarmQueue := alarmQueues[instance]

	if _t.Sr.GetUif() {
		_t.Sr.SetUif(false)
		base[instance] += 0x1_0000_0000
	}

	now := base[instance] + uint64(uint32(_t.Cnt.GetCntl())|(uint32(_t.Cnt.GetCnth())<<16))

	for ch := 0; ch < 4; ch++ {
		entry := alarms[ch]
		if entry == nil {
			continue
		}

		fired := false
		switch ch {
		case 0:
			fired = _t.Sr.GetCc1if()
			if fired {
				_t.Sr.SetCc1if(false)
				_t.Dier.SetCc1ie(false)
			}
		case 1:
			fired = _t.Sr.GetCc2if()
			if fired {
				_t.Sr.SetCc2if(false)
				_t.Dier.SetCc2ie(false)
			}
		case 2:
			fired = _t.Sr.GetCc3if()
			if fired {
				_t.Sr.SetCc3if(false)
				_t.Dier.SetCc3ie(false)
			}
		case 3:
			fired = _t.Sr.GetCc4if()
			if fired {
				_t.Sr.SetCc4if(false)
				_t.Dier.SetCc4ie(false)
			}
		}

		if fired {
			if now >= entry.deadline {
				alarms[ch] = nil
				entry.callback(now)

				// Promote next alarm from queue to this channel
				if alarmQueue != nil {
					promoted := alarmQueue
					alarmQueues[instance] = alarmQueue.next
					promoted.next = nil
					promoted.channel = ch
					alarms[ch] = promoted
					armCompare(instance, ch, promoted.deadline)
				}
			} else {
				// Reschedule this alarm.
				alarms[ch] = nil
				armCompare(instance, ch, entry.deadline)
			}
		}
	}
	runtime.EnableInterrupts(state)
}

func insertQueuedAlarm(instance _tim32, entry *alarmEntry) {
	alarmQueue := alarmQueues[instance]
	if alarmQueue == nil || entry.deadline < alarmQueue.deadline {
		entry.next = alarmQueues[instance]
		alarmQueues[instance] = entry
		return
	}

	curr := alarmQueue
	for curr.next != nil && curr.next.deadline < entry.deadline {
		curr = curr.next
	}
	entry.next = curr.next
	curr.next = entry
}

func armCompare(t _tim32, ch int, deadline uint64) {
	_t := tim.Instances[t]
	cc := uint32(deadline)
	switch ch {
	case 0:
		_t.Ccr1.SetCcr1l(uint16(cc))
		_t.Ccr1.SetCcr1h(uint16(cc >> 16))
		_t.Sr.SetCc1if(false)
		_t.Dier.SetCc1ie(true)
	case 1:
		_t.Ccr2.SetCcr2l(uint16(cc))
		_t.Ccr2.SetCcr2h(uint16(cc >> 16))
		_t.Sr.SetCc2if(false)
		_t.Dier.SetCc2ie(true)
	case 2:
		_t.Ccr3.SetCcr3l(uint16(cc))
		_t.Ccr3.SetCcr3h(uint16(cc >> 16))
		_t.Sr.SetCc3if(false)
		_t.Dier.SetCc3ie(true)
	case 3:
		_t.Ccr4.SetCcr4l(uint16(cc))
		_t.Ccr4.SetCcr4h(uint16(cc >> 16))
		_t.Sr.SetCc4if(false)
		_t.Dier.SetCc4ie(true)
	}
}
