package cortexm

import (
	"fmt"
	"nonstandard"
	"unsafe"
)

const basicExceptionFrameSize = unsafe.Sizeof(basicExceptionFrame{})
const baseGoroutineContextSize = unsafe.Sizeof(baseGoroutineContext{})
const stackCanary = 0xA55AA55A
const stackGuardWords = 16

//go:extern removeGoroutine  runtime.removeGoroutine
func removeGoroutine(unsafe.Pointer)

//go:export lastExcReturn runtime.lastExcReturn
var lastExcReturn uintptr = 0

//go:export lastRestorePSP runtime.lastRestorePSP
var lastRestorePSP uintptr = 0

type basicExceptionFrame struct {
	R0  uintptr
	R1  uintptr
	R2  uintptr
	R3  uintptr
	R12 uintptr
	LR  uintptr
	PC  uintptr
	PSR uintptr
}

func (e basicExceptionFrame) Dump() {
	fmt.Printf("R0: %#x\n", e.R0)
	fmt.Printf("R1: %#x\n", e.R1)
	fmt.Printf("R2: %#x\n", e.R2)
	fmt.Printf("R3: %#x\n", e.R3)
	fmt.Printf("R12: %#x\n", e.R12)
	fmt.Printf("LR: %#x\n", e.LR)
	fmt.Printf("PC: %#x\n", e.PC)
	fmt.Printf("PSR: %#x\n", e.PSR)
}

type exceptionFrame struct {
	basicExceptionFrame
	extendedExceptionFrame
}

func (e exceptionFrame) Dump() {
	e.basicExceptionFrame.Dump()
	e.extendedExceptionFrame.Dump()
}

type baseGoroutineContext struct {
	R4  uintptr
	R5  uintptr
	R6  uintptr
	R7  uintptr
	R8  uintptr
	R9  uintptr
	R10 uintptr
	R11 uintptr
	LR  uintptr
}

type _goroutine struct {
	stackTop unsafe.Pointer
	fn       struct {
		ptr       unsafe.Pointer
		args      unsafe.Pointer
		stackSize uintptr
	}
	stack unsafe.Pointer
}

//go:export initGoroutine runtime.initGoroutine
func initGoroutine(gptr unsafe.Pointer) {
	g := (*_goroutine)(gptr)
	stackSize := g.fn.stackSize

	// Build the initial stack as a BASIC (non-FP) frame.
	//
	// This matches the FreeRTOS port model and, crucially, matches the actual
	// hardware FP state of a brand-new goroutine: it has NOT executed any FP
	// instruction, so CONTROL.FPCA is 0 for it. The exception-return integrity
	// check requires that EXC_RETURN bit 4 (FP frame present) agree with FPCA.
	// A fresh goroutine must therefore return with 0xFFFFFFFD (basic, no FP),
	// NOT 0xFFFFFFED (extended). Building an extended/0xFFFFFFED initial frame
	// while FPCA=0 is exactly what caused the first switch into a new goroutine
	// to fault INVPC.
	//
	// FP context becomes relevant only AFTER a goroutine executes an FP
	// instruction: the hardware sets FPCA, the next exception stacks an
	// extended frame and enters PendSV with 0xFFFFFFED, and the conditional
	// save/restore in goroutine.S (keyed on EXC_RETURN bit 4) preserves
	// s16-s31 from then on. No frame is forced; the hardware drives the
	// basic/extended transition, and EXC_RETURN always reflects it.
	//
	// The stack grows DOWN on Cortex-M (high -> low address).

	// Add guard to the beginning of the stack.
	guard := unsafe.Slice((*uintptr)(g.stack), stackGuardWords)
	for i := range guard {
		guard[i] = uintptr(stackCanary)
	}

	// Determine the address at which the top of the stack should reside. On ARM,
	// the stack must be 8-byte aligned.
	const stackAlignment = uintptr(8)

	stackBase := uintptr(g.stackTop)
	stackEnd := stackBase + stackSize
	alignedStackEnd := stackEnd &^ (stackAlignment - 1)

	if alignedStackEnd < stackBase+basicExceptionFrameSize {
		panic("goroutine stack too small")
	}

	// Reserve the basic hardware exception frame (8 words) at the top.
	estack := (*basicExceptionFrame)(
		unsafe.Pointer(alignedStackEnd - basicExceptionFrameSize),
	)

	// Set up the call to startGoroutine on first exception-return entry.
	estack.PSR = defaultPsrValue // T-bit set on armv7-m
	estack.PC = uintptr(nonstandard.PointerOf(startGoroutine)) | thumbEnabledBit
	estack.R0 = uintptr(gptr) // startGoroutine(g)

	// Reserve the basic software-saved context (r4-r11 + EXC_RETURN, 9 words)
	// immediately below the hardware frame. No FP (s16-s31) slots: a fresh
	// goroutine has no FP state to restore.
	g.stackTop = unsafe.Add(unsafe.Pointer(estack), -int32(baseGoroutineContextSize))
	ctx := (*baseGoroutineContext)(g.stackTop)
	ctx.LR = 0xFFFF_FFFD // basic EXC_RETURN: thread mode, PSP, no FP frame
}

//go:export align runtime.alignStack
func align(n uintptr) uintptr {
	return (n + 7) &^ uintptr(7)
}

//sigo:export startGoroutine
//sigo:attribute startGoroutine noreturn
func startGoroutine(g *_goroutine) {
	// Start the goroutine.
	exec(g.fn.args, g.fn.ptr)

	// Remove the goroutine from the scheduler.
	removeGoroutine(unsafe.Pointer(g))

	// Busy loop until another task can be run.
	for {
		// Trigger a context switch.
		gosched()
	}
}
