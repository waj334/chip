//go:build stm32h7x7

package timer

import (
	"sync"
	"time"

	stm32h7x7 "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	tim "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/tim2_3_4_5"
)

var base [4]uint64
var alarms [4][4]*alarmEntry // [timer_instance][channel]
var alarmQueues [4]*alarmEntry
var mutex [4]sync.Mutex

const (
	TIM2 _tim32 = iota
	TIM3
	TIM4
	TIM5
)

type _tim32 uint8

// AlarmFunc is the callback type for timer alarms.
// The parameter is the current tick count when the alarm fired.
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
	mutex := &mutex[t]
	cs := sync.NewCriticalSection(mutex)
	cs.Begin()

	_t := tim.Instances[t]

	// Reset and enable/disable the timer in RCC.
	switch t {
	case 0:
		// Disable the IRQ.
		stm32h7x7.IrqTim2.Disable()

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
		stm32h7x7.IrqTim3.Disable()

		// Perform reset sequence.
		rcc.Rcc.Apb1lrstr.SetTim3rst(true)
		for !rcc.Rcc.Apb1lrstr.GetTim3rst() {
		}

		rcc.Rcc.Apb1lrstr.SetTim3rst(false)
		for rcc.Rcc.Apb1lrstr.GetTim3rst() {
		}

		rcc.Rcc.Apb1lenr.SetTim3en(config.Enable)
	case 2:
		stm32h7x7.IrqTim4.Disable()

		// Perform reset sequence.
		rcc.Rcc.Apb1lrstr.SetTim4rst(true)
		for !rcc.Rcc.Apb1lrstr.GetTim4rst() {
		}

		rcc.Rcc.Apb1lrstr.SetTim4rst(false)
		for rcc.Rcc.Apb1lrstr.GetTim4rst() {
		}

		rcc.Rcc.Apb1lenr.SetTim4en(config.Enable)
	case 3:
		stm32h7x7.IrqTim5.Disable()

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
		cs.End()
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
		stm32h7x7.IrqTim2.Enable()
	case 1:
		stm32h7x7.IrqTim3.Enable()
	case 2:
		stm32h7x7.IrqTim4.Enable()
	case 3:
		stm32h7x7.IrqTim5.Enable()
	}

	// Start the timer.
	_t.Cr1.SetCen(true)
	for !_t.Cr1.GetCen() {
	}

	cs.End()
	return nil
}

func (t _tim32) Reset(tick uint64) {
	cs := sync.NewCriticalSection(&mutex[t])
	cs.Begin()
	_t := tim.Instances[t]
	_t.Cnt.SetCntl(uint16(tick))
	_t.Cnt.SetCnth(uint16(tick >> 16))

	// Start the timer.
	_t.Cr1.SetCen(true)
	for !_t.Cr1.GetCen() {
	}
	cs.End()
}

func (t _tim32) Stop() {
	cs := sync.NewCriticalSection(&mutex[t])
	cs.Begin()
	_t := tim.Instances[t]
	_t.Cr1.SetCen(false)
	for _t.Cr1.GetCen() {
	}
	cs.End()
}

func (t _tim32) Tick() uint64 {
	cs := sync.NewCriticalSection(&mutex[t])
	cs.Begin()
	result := t.tick()
	cs.End()
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
	criticalSection := sync.NewCriticalSection(&mutex[t])
	criticalSection.Begin()

	now := t.tick()
	if fn == nil || deadline <= now {
		criticalSection.End()
		fn(now)
		return
	}

	// Try to assign to an available channel
	for ch := 0; ch < 4; ch++ {
		if alarms[t][ch] == nil {
			entry := &alarmEntry{deadline: deadline, callback: fn, channel: ch}
			alarms[t][ch] = entry
			armCompare(t, ch, deadline)
			criticalSection.End()
			return
		}
	}

	// No free channel. Sweep stale alarms whose deadlines are past.
	// CC flags can be lost due to read-modify-write races on the rc_w0 SR
	// register, leaving channels permanently occupied. Reclaim them here.
	var staleCallbacks [4]AlarmFunc
	staleCount := 0
	for ch := 0; ch < 4; ch++ {
		if alarms[t][ch] != nil && now >= alarms[t][ch].deadline {
			staleCallbacks[ch] = alarms[t][ch].callback
			disableCompareInterrupt(t, ch)
			alarms[t][ch] = nil
			staleCount++
		}
	}

	if staleCount > 0 {
		// Assign our new alarm to the first freed channel.
		for ch := 0; ch < 4; ch++ {
			if alarms[t][ch] == nil {
				newEntry := &alarmEntry{deadline: deadline, callback: fn, channel: ch}
				alarms[t][ch] = newEntry
				armCompare(t, ch, deadline)

				// Promote queued alarms into remaining freed channels.
				for ch2 := ch + 1; ch2 < 4; ch2++ {
					if alarms[t][ch2] == nil && alarmQueues[t] != nil {
						promoted := alarmQueues[t]
						alarmQueues[t] = promoted.next
						promoted.next = nil
						promoted.channel = ch2
						alarms[t][ch2] = promoted
						armCompare(t, ch2, promoted.deadline)
					}
				}

				criticalSection.End()

				// Fire stale callbacks outside the critical section.
				for i := 0; i < 4; i++ {
					if staleCallbacks[i] != nil {
						staleCallbacks[i](now)
					}
				}
				return
			}
		}
	}

	// No channels available. Queue it.
	entry := &alarmEntry{deadline: deadline, callback: fn, channel: -1}
	insertQueuedAlarm(t, entry)
	criticalSection.End()
	return
}

func (t _tim32) tick() uint64 {
	_t := tim.Instances[t]
	b := base[t]
	tick := uint64(uint32(_t.Cnt.GetCntl()) | (uint32(_t.Cnt.GetCnth()) << 16))
	result := b + tick
	return result
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
	criticalSection := sync.NewCriticalSection(nil)
	criticalSection.Begin()

	_t := tim.Instances[instance]

	if _t.Sr.GetUif() {
		_t.Sr.SetUif(false)
		base[instance] += 0x1_0000_0000
	}

	now := base[instance] + uint64(uint32(_t.Cnt.GetCntl())|(uint32(_t.Cnt.GetCnth())<<16))

	// Phase 1: Collect fired alarms and promote queued alarms into freed
	// channels. No callbacks are called in this phase, preventing re-entrant
	// SetAlarm calls from racing with the channel promotion logic.
	var fired [4]*alarmEntry

	for ch := 0; ch < 4; ch++ {
		entry := alarms[instance][ch]
		if entry == nil {
			continue
		}

		ccFired := false
		switch ch {
		case 0:
			ccFired = _t.Sr.GetCc1if()
			if ccFired {
				_t.Sr.SetCc1if(false)
				_t.Dier.SetCc1ie(false)
			}
		case 1:
			ccFired = _t.Sr.GetCc2if()
			if ccFired {
				_t.Sr.SetCc2if(false)
				_t.Dier.SetCc2ie(false)
			}
		case 2:
			ccFired = _t.Sr.GetCc3if()
			if ccFired {
				_t.Sr.SetCc3if(false)
				_t.Dier.SetCc3ie(false)
			}
		case 3:
			ccFired = _t.Sr.GetCc4if()
			if ccFired {
				_t.Sr.SetCc4if(false)
				_t.Dier.SetCc4ie(false)
			}
		}

		// Fire if deadline is past — whether or not the CC flag was seen.
		// CC flags can be lost via read-modify-write races on the rc_w0 SR
		// register, leaving channels permanently stuck. The deadline check
		// ensures stale channels are always reclaimed.
		if now >= entry.deadline {
			if !ccFired {
				disableCompareInterrupt(instance, ch)
			}

			alarms[instance][ch] = nil
			fired[ch] = entry

			// Promote next queued alarm to this channel.
			if alarmQueues[instance] != nil {
				promoted := alarmQueues[instance]
				alarmQueues[instance] = promoted.next
				promoted.next = nil
				promoted.channel = ch
				alarms[instance][ch] = promoted
				armCompare(instance, ch, promoted.deadline)
			}
		} else if ccFired {
			// CC fired but deadline not yet reached (spurious or early). Reschedule.
			armCompare(instance, ch, entry.deadline)
		}
	}

	criticalSection.End()

	// Phase 2: Fire callbacks outside the critical section. This allows
	// callbacks to safely interact with goroutine scheduling (goresume, etc.)
	// without holding any locks.
	for ch := 0; ch < 4; ch++ {
		if fired[ch] != nil {
			fired[ch].callback(now)
		}
	}
}

func disableCompareInterrupt(t _tim32, ch int) {
	_t := tim.Instances[t]
	switch ch {
	case 0:
		_t.Sr.SetCc1if(false)
		_t.Dier.SetCc1ie(false)
	case 1:
		_t.Sr.SetCc2if(false)
		_t.Dier.SetCc2ie(false)
	case 2:
		_t.Sr.SetCc3if(false)
		_t.Dier.SetCc3ie(false)
	case 3:
		_t.Sr.SetCc4if(false)
		_t.Dier.SetCc4ie(false)
	}
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

	// If the deadline is already in the past, the compare match will never
	// fire (STM32 only matches on CNT == CCR, not CNT > CCR). Force a
	// capture/compare event via EGR to set the CCxIF flag. SR bits are
	// rc_w0 (writing 1 has no effect), so we must use EGR to generate
	// the event from software.
	cnt := uint64(uint32(_t.Cnt.GetCntl()) | (uint32(_t.Cnt.GetCnth()) << 16))
	if deadline <= base[t]+cnt {
		switch ch {
		case 0:
			_t.Egr.SetCc1g(true)
		case 1:
			_t.Egr.SetCc2g(true)
		case 2:
			_t.Egr.SetCc3g(true)
		case 3:
			_t.Egr.SetCc4g(true)
		}
	}
}
