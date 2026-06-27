package qspi

import (
	"sync"
	"unsafe"
	"volatile"

	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/quadspi"
)

var driverState _state

type _state struct {
	mutex  sync.Mutex
	config Config

	data  []byte
	index int32 // bytes processed
	done  int32 // 0 = not done, 1 = done
	dir   int32 // 0 = none, 1 = read, 2 = write

	lastError error
}

//sigo:interrupt quadspiHandler QuadspiHandler
func quadspiHandler() {
	state := &driverState
	sr := quadspi.Quadspi.Sr
	fcr := &quadspi.Quadspi.Fcr
	cr := &quadspi.Quadspi.Cr

	if volatile.LoadInt32(&state.done) != 0 {
		// Stale interrupt — mask everything and bail.
		cr.SetFtie(false)
		cr.SetTcie(false)
		cr.SetSmie(false)
		cr.SetTeie(false)
		fcr.StoreBits(0x1B) // clear TOF | SMF | TCF | TEF
		return
	}

	dir := volatile.LoadInt32(&state.dir)
	data := state.data
	length := len(data)
	index := int(volatile.LoadInt32(&state.index))

	// FIFO service.
	if sr.GetTcf() {
		// For indirect RX, TCF can arrive while the final bytes are still sitting
		// in the FIFO below/around the FIFO threshold. Drain them before done.
		if dir == 1 {
			for index < length && quadspi.Quadspi.Sr.GetFlevel() != 0 {
				// Perform byte reads.
				data[index] = volatile.LoadUint8((*uint8)(unsafe.Pointer(&quadspi.Quadspi.Dr)))
				index++
			}
			volatile.StoreInt32(&state.index, int32(index))
		}

		fcr.SetCtcf(true)

		cr.SetFtie(false)
		cr.SetTcie(false)
		cr.SetTeie(false)

		volatile.StoreInt32(&state.done, 1)
	}

	// Errors.
	if sr.GetTef() {
		fcr.SetCtef(true)
		state.lastError = ErrTransfer
		cr.SetFtie(false)
		cr.SetTcie(false)
		cr.SetSmie(false)
		cr.SetTeie(false)
		volatile.StoreInt32(&state.done, 1)
		return
	}

	// Terminal flags.
	if sr.GetTcf() {
		fcr.SetCtcf(true)
		cr.SetFtie(false)
		cr.SetTcie(false)
		cr.SetTeie(false)
		volatile.StoreInt32(&state.done, 1)
	}

	if sr.GetSmf() {
		fcr.SetCsmf(true)
		cr.SetSmie(false)
		cr.SetTeie(false)
		volatile.StoreInt32(&state.done, 1)
	}

	if sr.GetTof() {
		fcr.SetCtof(true)
		// Memory-mapped idle timeout — informational; nobody's waiting.
	}
}
