package rng

import (
	regs "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rng"
)

var (
	RNG rng
)

type rng struct{}

func (r rng) Init() {
	// Disable the RNG.
	regs.Rng.Cr.Store(0)

	// Configure the RNG.
	regs.Rng.Cr.SetCed(false)

	// Enable the RNG.
	regs.Rng.Cr.SetRngen(true)
}

func (r rng) Uint64() uint64 {
	for !regs.Rng.Sr.GetDrdy() {
		if regs.Rng.Sr.GetSeis() {
			// Seed error: clear, disable, re-enable per RM recovery.
			regs.Rng.Sr.SetSeis(false)
			regs.Rng.Cr.SetRngen(false)
			regs.Rng.Cr.SetRngen(true)
			return 0
		}
		if regs.Rng.Sr.GetCeis() {
			regs.Rng.Sr.SetCeis(false)
		}
	}
	return uint64(regs.Rng.Dr.GetRndata())
}
