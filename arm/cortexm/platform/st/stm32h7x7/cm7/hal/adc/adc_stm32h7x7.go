//go:build stm32h7x7

package adc

import (
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/adc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/pwr"
	"pkg.si-go.dev/chip/core/hal"
	"runtime"
	"sync"
	"time"
)

const (
	ADC1 _adc = iota
	ADC2
	ADC3
)

type Resolution uint8

const (
	Resolution8Bit Resolution = iota + 1
	Resolution10Bit
	Resolution12Bit
	Resolution14Bit
	Resolution16Bit
)

type Config struct {
	Enable     bool
	Resolution Resolution
}

var mutex [3]sync.Mutex

type _adc uint8

func (a _adc) Configure(config Config) error {
	instance := adc.Instances[a]

	// Validate configuration.
	if config.Resolution > Resolution16Bit {
		return hal.ErrInvalidConfig
	}

	mutex[a].Lock()
	if instance.Cr.GetAden() {
		// Disable the instance.
		instance.Cr.SetAden(false)
		for instance.Cr.GetAden() {
			runtime.Gosched()
		}

		// Disable the ADC voltage regulator.
		instance.Cr.SetAdvregen(false)
		for instance.Cr.GetAdvregen() {
			runtime.Gosched()
		}

		// Enable deep power down.
		instance.Cr.SetDeeppwd(false)
	}

	if config.Enable {
		// Disable deep power down.
		instance.Cr.SetDeeppwd(false)

		// Start ADC calibration
		instance.Cr.SetAdcal(true)
		for instance.Cr.GetAdcal() {
			runtime.Gosched()
		}

		// Enable the ADC voltage regulator.
		instance.Cr.SetAdvregen(true)
		for !instance.Cr.GetAdvregen() {
			runtime.Gosched()
		}

		// Enable reference voltage.
		pwr.Pwr.Cr3.SetVbe(true)

		// Wait for the ADC to stabilize.
		time.Sleep(time.Second)

		if config.Resolution > 0 {
			// Set the data resolution.
			switch config.Resolution {
			case Resolution8Bit:
				instance.Cfgr.SetRes(adc.RegisterCfgrFieldResEnumResolution8bit)
			case Resolution10Bit:
				instance.Cfgr.SetRes(adc.RegisterCfgrFieldResEnumResolution10bit)
			case Resolution12Bit:
				instance.Cfgr.SetRes(adc.RegisterCfgrFieldResEnumResolution12bit)
			case Resolution14Bit:
				instance.Cfgr.SetRes(adc.RegisterCfgrFieldResEnumResolution14bit)
			case Resolution16Bit:
				instance.Cfgr.SetRes(adc.RegisterCfgrFieldResEnumResolution16bit)
			default:
				panic("unreachable")
			}
		}

		// Sample all channels.
		instance.Pcsel.SetPcsel(0xF_FFFF)

		// Enable the ADC and then wait for it to become ready.
		instance.Cr.SetAden(true)
		for !instance.Isr.GetAdrdy() {
			runtime.Gosched()
		}
	}

	mutex[a].Unlock()
	return nil
}

func (a _adc) Read(channel int) (uint, error) {
	instance := adc.Instances[a]
	mutex[a].Lock()
	// The _adc must be enabled.
	if !instance.Cr.GetAden() {
		mutex[a].Unlock()
		return 0, hal.ErrInvalidState
	}

	if instance.Isr.GetEoc() {
		// Read and discard stale data.
		_ = instance.Dr.GetRdata()
	}

	instance.Sqr1.SetSq1(uint8(channel))
	instance.Cr.SetAdstart(true)
	for !instance.Isr.GetEoc() {
		// Yield to another goroutine while waiting on the conversion to complete.
		runtime.Gosched()
	}
	value := instance.Dr.GetRdata()
	mutex[a].Unlock()

	return uint(value), nil
}
