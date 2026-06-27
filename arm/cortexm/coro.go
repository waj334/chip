package cortexm

import (
	"nonstandard"
	"unsafe"
)

//sigo:extern _coroStackSize runtime._coroStackSize
var _coroStackSize uintptr

// _coro mirrors the first two fields of runtime.coro in SiGo. Only sp and
// stack needs to be reachable from the chip-support side. Order MUST match
// runtime.coro in SiGo: sp first, stack second.
type _coro struct {
	sp    unsafe.Pointer
	stack unsafe.Pointer
}

//sigo:extern coroEntry runtime.coroEntry
func coroEntry(c unsafe.Pointer)

// initCoro lays out the initial stack frame for a fresh coroutine, so the
// first coroswitch_inner pop+bx jumps to coroEntry with r0 holding the
// coro pointer.
//
// Frame layout (low → high addresses on the coro stack):
//
//	[extendedGoroutineContext = s16...s31] (only when FPU enabled)
//	[basicGoroutineContext = r4...r11, lr=coroEntry]
//
// coroswitch_inner's pop sequence reads the FPU regs first (vldmia sp!),
// then the integer regs (pop {r4-r11, lr}), then bx lr. After bx, control
// lands in coroEntry. r0 was never touched by the asm, so it still holds
// the coro pointer the parent passed when it called coroswitch_inner.
//
//go:export initCoro runtime.initCoro
func initCoro(cptr unsafe.Pointer) {
	c := (*_coro)(cptr)

	frameSize := baseGoroutineContextSize + unsafe.Sizeof(extendedGoroutineContext{})
	top := uintptr(c.stack) + _coroStackSize
	frameAddr := top - frameSize

	// Zero the entire frame (FPU regs + integer regs) defensively.
	for i := uintptr(0); i < frameSize; i += 4 {
		*(*uint32)(unsafe.Pointer(frameAddr + i)) = 0
	}

	// Integer context starts after the FPU context (FPU at low end of frame).
	basicAddr := frameAddr + uintptr(unsafe.Sizeof(extendedGoroutineContext{}))
	basic := (*baseGoroutineContext)(unsafe.Pointer(basicAddr))
	basic.LR = uintptr(nonstandard.PointerOf(coroEntry)) | thumbEnabledBit

	c.sp = unsafe.Pointer(frameAddr)
}
