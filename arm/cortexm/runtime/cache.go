//go:build cortexm7 || cortexm55 || cortexm85

package runtime

import (
	"unsafe"

	"pkg.si-go.dev/chip/arm/cortexm/reg/cmo"
)

type DCache struct{}

func (DCache) Clean(addr unsafe.Pointer, size uintptr) {
	const lineSize = 32
	start := uintptr(addr) &^ (lineSize - 1)
	end := (uintptr(addr) + size + lineSize - 1) &^ (lineSize - 1)

	// DSB ensures prior buffer writes complete before cache lines are
	// written back, preventing stale data from reaching memory.
	DSB()
	for p := start; p < end; p += lineSize {
		cmo.Cmo.Dccmvac.SetAddress(uint32(p))
	}
	// DSB ensures all writebacks complete before subsequent code (e.g.,
	// a DMA start) executes; ISB flushes the pipeline.
	DSB()
	ISB()
}

func (DCache) Invalidate(addr unsafe.Pointer, size uintptr) {
	const lineSize = 32
	start := uintptr(addr) &^ (lineSize - 1)
	end := (uintptr(addr) + size + lineSize - 1) &^ (lineSize - 1)
	// DSB ensures prior accesses complete before cache lines are
	// invalidated, so no in-flight load returns stale data after
	// invalidation.
	DSB()
	for p := start; p < end; p += lineSize {
		cmo.Cmo.Dcimvac.SetAddress(uint32(p))
	}
	// DSB ensures all invalidations complete before subsequent code
	// (e.g., reading the buffer after DMA wrote it) executes; ISB
	// flushes the pipeline.
	DSB()
	ISB()
}

func (DCache) CleanInvalidate(addr unsafe.Pointer, size uintptr) {
	const lineSize = 32
	start := uintptr(addr) &^ (lineSize - 1)
	end := (uintptr(addr) + size + lineSize - 1) &^ (lineSize - 1)
	// DSB ensures prior buffer writes complete before cache lines are
	// written back and invalidated.
	DSB()
	for p := start; p < end; p += lineSize {
		cmo.Cmo.Dccimvac.SetAddress(uint32(p))
	}
	// DSB ensures all clean+invalidate operations complete before
	// subsequent code executes; ISB flushes the pipeline.
	DSB()
	ISB()
}
