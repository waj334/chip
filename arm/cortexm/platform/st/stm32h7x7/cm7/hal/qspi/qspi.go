//go:build stm32h7x7

package qspi

import (
	stm32h7x7 "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/quadspi"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	corehal "pkg.si-go.dev/chip/core/hal"
)

var (
	QSPI QSpi
)

type FunctionalMode uint8

const (
	ModeIndirectWrite FunctionalMode = 0
	ModeIndirectRead  FunctionalMode = 1
	ModeAutoPoll      FunctionalMode = 2
	ModeMemoryMapped  FunctionalMode = 3
)

type SignalMode uint8

const (
	SingleMode SignalMode = iota
	DualMode
	QuadMode
)

type FlashSelection uint8

const (
	Flash0 FlashSelection = iota
	Flash1
)

type QSpi struct{}

type Config struct {
	Enable    bool
	Signaling SignalMode
	DualFlash bool
	Selection FlashSelection

	// Prescaler divides the QSPI kernel clock by (Prescaler). 1..255.
	Prescaler uint8

	// Size is the FSIZE field value. Flash capacity = 2^(Size+1) bytes.
	// For a 16 MiB chip: Size = 23.
	Size uint8

	// CSHighCycles is the minimum NCS-high time between transactions, in
	// QSPI clock cycles. 1..8. Stored as (N-1) in DCR.CSHT.
	CSHighCycles uint8

	// ClockMode3: false = mode 0 (CLK idles low), true = mode 3 (CLK idles high).
	ClockMode3 bool

	// FIFOThreshold sets when FTF asserts: FIFO has at least (N+1) bytes
	// available (for reads) or free (for writes). 0..31.
	FIFOThreshold uint8

	// SampleShift shifts the input sampling point by half a clock cycle,
	// improving timing margin at high frequencies. Recommended: true.
	SampleShift bool

	CLK pin.Pin // Clock to FLASH 1 and FLASH 2.

	BK1_NCS pin.Pin // Chip select (active low) for FLASH 1. Can also be used for FLASH 2 if QUADSPI is always used in dual-flash mode.
	BK1_IO0 pin.Pin // Bidirectional I/O in dual/quad modes or serial output in single mode, for FLASH 1.
	BK1_IO1 pin.Pin // Bidirectional I/O in dual/quad modes or serial input in single mode, for FLASH 1.
	BK1_IO2 pin.Pin // Bidirectional I/O in quad mode, for FLASH 1.
	BK1_IO3 pin.Pin // Bidirectional I/O in quad mode, for FLASH 1.

	BK2_NCS pin.Pin // Chip select (active low) for FLASH 2. Can also be used for FLASH 1 if QUADSPI is always used in dual-flash mode.
	BK2_IO0 pin.Pin // Bidirectional I/O in dual/quad modes or serial output in single mode, for FLASH 2.
	BK2_IO1 pin.Pin // Bidirectional I/O in dual/quad modes or serial input in single mode, for FLASH 2.
	BK2_IO2 pin.Pin // Bidirectional I/O in quad mode, for FLASH 2.
	BK2_IO3 pin.Pin // Bidirectional I/O in quad mode, for FLASH 2.
}

func (q QSpi) Configure(config Config) error {
	state := &driverState

	state.mutex.Lock()

	// Disable IRQ and reset the peripheral. Done unconditionally so a
	// Configure(Enable: false) cleanly tears down whatever was there.
	stm32h7x7.IrqQuadspi.Disable()
	rcc.Rcc.Ahb3rstr.SetQspirst(true)
	for !rcc.Rcc.Ahb3rstr.GetQspirst() {
	}
	rcc.Rcc.Ahb3rstr.SetQspirst(false)
	for rcc.Rcc.Ahb3rstr.GetQspirst() {
	}

	if !config.Enable {
		rcc.Rcc.Ahb3enr.SetQspien(false)
		state.mutex.Unlock()
		return nil
	}

	// Enable bus clock so register writes stick.
	rcc.Rcc.Ahb3enr.SetQspien(true)
	for !rcc.Rcc.Ahb3enr.GetQspien() {
	}

	needBank1 := config.Selection == Flash0 || config.DualFlash
	needBank2 := config.Selection == Flash1 || config.DualFlash

	if config.Signaling == QuadMode {
		if needBank1 {
			if config.BK1_IO2 == pin.NoPin || config.BK1_IO3 == pin.NoPin {
				state.mutex.Unlock()
				return corehal.ErrInvalidPinout
			}
		}
		if needBank2 {
			if config.BK2_IO2 == pin.NoPin || config.BK2_IO3 == pin.NoPin {
				state.mutex.Unlock()
				return corehal.ErrInvalidPinout
			}
		}
	}

	if config.CSHighCycles < 1 || config.CSHighCycles > 8 {
		state.mutex.Unlock()
		return corehal.ErrInvalidConfig
	}

	switch config.CLK {
	case pin.PB2, pin.PF10:
		config.CLK.SetMode(pin.AltFunction | pin.Alt9)
	default:
		state.mutex.Unlock()
		return corehal.ErrInvalidPinout
	}

	config.CLK.SetSpeedMode(pin.VeryHighSpeed)
	config.CLK.SetPullMode(pin.NoPull)

	// Bank 1.
	if needBank1 {
		switch config.BK1_NCS {
		case pin.PB10:
			config.BK1_NCS.SetMode(pin.AltFunction | pin.Alt9)
		case pin.PB6, pin.PG6:
			config.BK1_NCS.SetMode(pin.AltFunction | pin.Alt10)
		default:
			state.mutex.Unlock()
			return corehal.ErrInvalidPinout
		}

		config.BK1_NCS.SetSpeedMode(pin.VeryHighSpeed)
		config.BK1_NCS.SetPullMode(pin.NoPull)

		switch config.BK1_IO0 {
		case pin.PC9, pin.PD11:
			config.BK1_IO0.SetMode(pin.AltFunction | pin.Alt9)
		case pin.PF8:
			config.BK1_IO0.SetMode(pin.AltFunction | pin.Alt10)
		default:
			state.mutex.Unlock()
			return corehal.ErrInvalidPinout
		}

		config.BK1_IO0.SetSpeedMode(pin.VeryHighSpeed)
		config.BK1_IO0.SetPullMode(pin.NoPull)

		switch config.BK1_IO1 {
		case pin.PC10, pin.PD12:
			config.BK1_IO1.SetMode(pin.AltFunction | pin.Alt9)
		case pin.PF9:
			config.BK1_IO1.SetMode(pin.AltFunction | pin.Alt10)
		default:
			state.mutex.Unlock()
			return corehal.ErrInvalidPinout
		}

		config.BK1_IO1.SetSpeedMode(pin.VeryHighSpeed)
		config.BK1_IO1.SetPullMode(pin.NoPull)

		if config.BK1_IO2 != pin.NoPin {
			switch config.BK1_IO2 {
			case pin.PE2, pin.PF7:
				config.BK1_IO2.SetMode(pin.AltFunction | pin.Alt9)
			default:
				state.mutex.Unlock()
				return corehal.ErrInvalidPinout
			}

			config.BK1_IO2.SetSpeedMode(pin.VeryHighSpeed)
			config.BK1_IO2.SetPullMode(pin.NoPull)
		}

		if config.BK1_IO3 != pin.NoPin {
			switch config.BK1_IO3 {
			case pin.PA1, pin.PD13, pin.PF6:
				config.BK1_IO3.SetMode(pin.AltFunction | pin.Alt9)
			default:
				state.mutex.Unlock()
				return corehal.ErrInvalidPinout
			}

			config.BK1_IO3.SetSpeedMode(pin.VeryHighSpeed)
			config.BK1_IO3.SetPullMode(pin.NoPull)
		}
	}

	// Bank 2.
	if needBank2 {
		switch config.BK2_NCS {
		case pin.PC11:
			config.BK2_NCS.SetMode(pin.AltFunction | pin.Alt9)
		default:
			state.mutex.Unlock()
			return corehal.ErrInvalidPinout
		}

		config.BK2_NCS.SetSpeedMode(pin.VeryHighSpeed)
		config.BK2_NCS.SetPullMode(pin.NoPull)

		switch config.BK2_IO0 {
		case pin.PH2:
			config.BK2_IO0.SetMode(pin.AltFunction | pin.Alt9)
		case pin.PE7:
			config.BK2_IO0.SetMode(pin.AltFunction | pin.Alt10)
		default:
			state.mutex.Unlock()
			return corehal.ErrInvalidPinout
		}

		config.BK2_IO0.SetSpeedMode(pin.VeryHighSpeed)
		config.BK2_IO0.SetPullMode(pin.NoPull)

		switch config.BK2_IO1 {
		case pin.PH3:
			config.BK2_IO1.SetMode(pin.AltFunction | pin.Alt9)
		case pin.PE8:
			config.BK2_IO1.SetMode(pin.AltFunction | pin.Alt10)
		default:
			state.mutex.Unlock()
			return corehal.ErrInvalidPinout
		}

		config.BK2_IO1.SetSpeedMode(pin.VeryHighSpeed)
		config.BK2_IO1.SetPullMode(pin.NoPull)

		if config.BK2_IO2 != pin.NoPin {
			switch config.BK2_IO2 {
			case pin.PG9:
				config.BK2_IO2.SetMode(pin.AltFunction | pin.Alt9)
			case pin.PE9:
				config.BK2_IO2.SetMode(pin.AltFunction | pin.Alt10)
			default:
				state.mutex.Unlock()
				return corehal.ErrInvalidPinout
			}

			config.BK2_IO2.SetSpeedMode(pin.VeryHighSpeed)
			config.BK2_IO2.SetPullMode(pin.NoPull)
		}

		if config.BK2_IO3 != pin.NoPin {
			switch config.BK2_IO3 {
			case pin.PG14:
				config.BK2_IO3.SetMode(pin.AltFunction | pin.Alt9)
			case pin.PE10:
				config.BK2_IO3.SetMode(pin.AltFunction | pin.Alt10)
			default:
				state.mutex.Unlock()
				return corehal.ErrInvalidPinout
			}

			config.BK2_IO3.SetSpeedMode(pin.VeryHighSpeed)
			config.BK2_IO3.SetPullMode(pin.NoPull)
		}
	}

	// Refuse to reconfigure while a transaction is in flight. Reading SR
	// directly avoids any shadow state.
	if quadspi.Quadspi.Sr.GetBusy() {
		state.mutex.Unlock()
		return corehal.ErrBusy
	}

	// Register config. EN must be 0 when changing FSIZE/PRESCALER/DFM/FSEL,
	// and it is 0 here because we just reset the peripheral.
	quadspi.Quadspi.Dcr.SetFsize(config.Size)
	quadspi.Quadspi.Dcr.SetCsht(config.CSHighCycles - 1)
	quadspi.Quadspi.Dcr.SetCkmode(config.ClockMode3)

	quadspi.Quadspi.Cr.SetPrescaler(min(255, max(0, config.Prescaler-1)))
	quadspi.Quadspi.Cr.SetFthres(config.FIFOThreshold)
	quadspi.Quadspi.Cr.SetSshift(config.SampleShift)
	quadspi.Quadspi.Cr.SetDfm(config.DualFlash)
	quadspi.Quadspi.Cr.SetFsel(config.Selection == Flash1)

	// Bring the peripheral online and arm the IRQ. Per-transaction interrupt
	// sources (TCIE/FTIE/SMIE/TEIE) are toggled inside Execute / AutoPoll;
	// here we only enable the line at the NVIC.
	quadspi.Quadspi.Cr.SetEn(true)
	stm32h7x7.IrqQuadspi.Enable()

	state.config = config
	state.mutex.Unlock()
	return nil
}

// Enabled reports whether the QSPI peripheral is currently enabled.
func (q QSpi) Enabled() bool {
	return quadspi.Quadspi.Cr.GetEn()
}

// Busy reports whether a transaction is currently in progress.
func (q QSpi) Busy() bool {
	return quadspi.Quadspi.Sr.GetBusy()
}

// FIFOLevel returns the current number of bytes in the FIFO (0..32).
func (q QSpi) FIFOLevel() uint8 {
	return quadspi.Quadspi.Sr.GetFlevel()
}

// FlashSizeLog2 returns the configured FSIZE value. Capacity in bytes is
// 1 << (FlashSizeLog2()+1).
func (q QSpi) FlashSizeLog2() uint8 {
	return quadspi.Quadspi.Dcr.GetFsize()
}

// Capacity returns the configured flash capacity in bytes.
func (q QSpi) Capacity() uint64 {
	return uint64(1) << (uint64(q.FlashSizeLog2()) + 1)
}

// Prescaler returns the current clock prescaler. Effective QSPI clock is
// kernel_clock / (Prescaler()).
func (q QSpi) Prescaler() uint8 {
	return quadspi.Quadspi.Cr.GetPrescaler()
}

// SetPrescaler updates the clock prescaler. Must not be called while Busy().
func (q QSpi) SetPrescaler(p uint8) error {
	if q.Busy() {
		return corehal.ErrBusy
	}
	quadspi.Quadspi.Cr.SetPrescaler(min(255, max(0, p-1)))
	return nil
}

// FunctionalMode returns the FMODE of the last-configured transaction.
// Mostly useful as IsMemoryMapped() — Execute() uses it to decide whether
// it needs to abort the peripheral before issuing a new indirect command.
func (q QSpi) FunctionalMode() FunctionalMode {
	return FunctionalMode(quadspi.Quadspi.Ccr.GetFmode())
}

// IsMemoryMapped reports whether the peripheral is currently in memory-mapped
// mode (FMODE == 11). When true, the flash is accessible at 0x9000_0000.
func (q QSpi) IsMemoryMapped() bool {
	return q.FunctionalMode() == ModeMemoryMapped
}
