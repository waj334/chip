//go:build arm && fpu

package runtime

import (
	"pkg.si-go.dev/chip/arm/cortexm/reg/fpu"
	"pkg.si-go.dev/chip/arm/cortexm/reg/scb"
)

type extendedFrame struct {
	Sn    [16]float32
	FPSCR uintptr
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
		fpu.Fpu.Fpccr.SetAspen(true)
		fpu.Fpu.Fpccr.SetLspen(true)
	}
}
