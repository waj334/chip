//go:build arm && !fpu

package cortexm

type extendedExceptionFrame struct {
	// Empty when the FPU is disabled: no FP context is stacked, so
	// stackFrame == basicFrame in size.
}

func (e extendedExceptionFrame) Dump() {
}

type extendedGoroutineContext struct {
	// Empty when the FPU is disabled: no FP callee-saved regs to preserve,
	// so goroutineContext == basicGoroutineContext in size.
}

func initFPU() {
	// Do nothing.
}
