//go:build arm && fpu

package cortexm

import (
	"fmt"

	"pkg.si-go.dev/chip/arm/cortexm/reg/fpu"
	"pkg.si-go.dev/chip/arm/cortexm/reg/scb"
)

type extendedExceptionFrame struct {
	Sn    [16]float32
	FPSCR uintptr
	_     uintptr // Reserved
}

func (e extendedExceptionFrame) Dump() {
	for i := range e.Sn {
		fmt.Printf("S%d: %e\n", i, e.Sn[i])
	}
	fmt.Printf("FPSCR: %#x\n", e.FPSCR)
}

type extendedGoroutineContext struct {
	Sn [16]float32
}

//sigo:extern _fpuEnabled runtime._fpuEnabled
var _fpuEnabled bool

func initFPU() {
	if _fpuEnabled {
		scb.Scb.Cpacr.SetCp10(scb.RegisterCpacrFieldCp10EnumFull)
		scb.Scb.Cpacr.SetCp11(scb.RegisterCpacrFieldCp11EnumFull)

		// ASPEN=1: hardware automatically sets CONTROL.FPCA on the first FP
		// instruction and automatically stacks the extended exception frame
		// (and selects the FP EXC_RETURN) whenever FPCA is set at exception
		// entry. LSPEN=1: lazy FP stacking. Together these let the hardware
		// drive the basic<->extended transition per goroutine, so EXC_RETURN
		// always agrees with FPCA and the exception-return integrity check
		// never fails. We do NOT force FPCA: forcing it desynchronizes
		// EXC_RETURN (which the hardware latches at entry) from CONTROL.FPCA
		// and causes INVPC. Fresh goroutines start basic (FPCA=0, 0xFFFFFFFD)
		// and become FP-active naturally; goroutine.S saves/restores s16-s31
		// conditionally on EXC_RETURN bit 4 to match.
		fpu.Fpu.Fpccr.SetAspen(true)
		fpu.Fpu.Fpccr.SetLspen(true)
	}
}
