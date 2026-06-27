package qspi

import (
	"errors"
	"runtime"
	"time"
	"volatile"

	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/quadspi"
	corehal "pkg.si-go.dev/chip/core/hal"
)

const (
	defaultTimeout = time.Second
)

type Lines uint8

const (
	LinesNone   Lines = 0
	LinesSingle Lines = 1
	LinesDual   Lines = 2
	LinesQuad   Lines = 3
)

type AddrSize uint8

const (
	AddrSize8  AddrSize = 0
	AddrSize16 AddrSize = 1
	AddrSize24 AddrSize = 2
	AddrSize32 AddrSize = 3
)

type Direction uint8

const (
	DirNone Direction = iota
	DirRead
	DirWrite
)

// Command describes one QSPI transaction. A phase is included only when its
// Lines field is non-zero (LinesNone == 0). The peripheral runs phases in
// fixed order: instruction, address, alternate bytes, dummy, data.
type Command struct {
	Instruction uint8
	InstrLines  Lines

	Address   uint32
	AddrLines Lines
	AddrSize  AddrSize

	AltBytes uint32
	AltLines Lines
	AltSize  AddrSize

	DummyCycles uint8 // 0..31

	DataLines Lines
	Direction Direction
	Data      []byte // nil for DirNone; length sets DLR

	// SIOO ("send instruction only once") is honored by the peripheral in
	// memory-mapped mode only. Ignored on indirect transactions.
	SIOO bool

	// DDR enables double-data-rate clocking. Not supported on the AT25SF128A;
	// included for completeness.
	DDR bool

	// Timeout bounds the whole transaction. Zero means use a sensible default.
	Timeout time.Duration
}

var (
	ErrTimeout  = errors.New("qspi: timeout")
	ErrTransfer = errors.New("qspi: transfer error")
)

// Execute runs cmd synchronously. The peripheral must be enabled. If the
// peripheral is currently in memory-mapped mode, Execute aborts it first.
func (q QSpi) Execute(cmd *Command) error {
	state := &driverState

	if !q.Enabled() {
		return corehal.ErrNotReady
	}

	state.mutex.Lock()
	defer state.mutex.Unlock()

	if q.IsMemoryMapped() {
		if err := q.abortImpl(); err != nil {
			return err
		}
	}
	if err := q.waitNotBusy(cmd.timeout()); err != nil {
		return err
	}

	// Reset transaction state.
	state.data = cmd.Data
	state.lastError = nil
	volatile.StoreInt32(&state.index, 0)
	volatile.StoreInt32(&state.done, 0)
	switch cmd.Direction {
	case DirRead:
		volatile.StoreInt32(&state.dir, 1)
	case DirWrite:
		volatile.StoreInt32(&state.dir, 2)
	default:
		volatile.StoreInt32(&state.dir, 0)
	}

	q.clearFlags()

	if cmd.Direction != DirNone && len(cmd.Data) > 0 {
		quadspi.Quadspi.Dlr.SetDl(uint32(len(cmd.Data) - 1))
	}
	if cmd.AltLines != LinesNone {
		quadspi.Quadspi.Abr.SetAlternate(cmd.AltBytes)
	}

	fmode := uint8(ModeIndirectWrite)
	if cmd.Direction == DirRead {
		fmode = uint8(ModeIndirectRead)
	}

	// Arm interrupts before triggering the transaction.
	quadspi.Quadspi.Cr.SetFtie(false)
	quadspi.Quadspi.Cr.SetTcie(false)
	quadspi.Quadspi.Cr.SetTeie(false)

	q.clearFlags()

	quadspi.Quadspi.Cr.SetTeie(true)
	quadspi.Quadspi.Cr.SetTcie(true)

	q.writeCCR(cmd, fmode)

	if cmd.AddrLines != LinesNone {
		quadspi.Quadspi.Ar.SetAddress(cmd.Address)
	}

	if cmd.Direction != DirNone && len(cmd.Data) > 0 {
		quadspi.Quadspi.Cr.SetFtie(true)
	}

	// Wait for ISR completion.
	deadline := time.Now().Add(cmd.timeout())
	for volatile.LoadInt32(&state.done) == 0 {
		if time.Now().After(deadline) {
			quadspi.Quadspi.Cr.SetFtie(false)
			quadspi.Quadspi.Cr.SetTcie(false)
			quadspi.Quadspi.Cr.SetTeie(false)
			_ = q.abortImpl()
			volatile.StoreInt32(&state.done, 1)
			return ErrTimeout
		}
		runtime.Gosched()
	}

	return state.lastError
}

// MapMemory configures the peripheral for memory-mapped reads using template
// as the auto-issued read command. The Data field of template is ignored
// (memory-mapped reads pull whatever length the CPU requests). After this
// returns successfully, the flash is readable at 0x9000_0000.
//
// SIOO in the template is honored: when true, the instruction is sent only
// for the first access and skipped for subsequent contiguous reads, which
// is the standard XIP-friendly optimization.
//
// TEIE is armed so the ISR can record an illegal-access error (e.g. a write
// into 0x9000_xxxx) into state.lastError. Memory-mapped mode never asserts
// TCF, so only TEF can wake the ISR after this returns.
func (q QSpi) MapMemory(template *Command) error {
	if !q.Enabled() {
		return corehal.ErrNotReady
	}
	if template.Direction != DirRead {
		return corehal.ErrInvalidConfig
	}

	state := &driverState
	state.mutex.Lock()
	defer state.mutex.Unlock()

	if q.IsMemoryMapped() {
		if err := q.abortImpl(); err != nil {
			return err
		}
	}
	if err := q.waitNotBusy(template.timeout()); err != nil {
		return err
	}

	// Reset transaction state so the ISR's "done != 0" early-return path
	// doesn't suppress the first TEF interrupt after we enter memory-mapped
	// mode. lastError is cleared so a stale error from the previous Execute
	// doesn't get reported as a memory-mapping failure.
	state.data = nil
	state.lastError = nil
	volatile.StoreInt32(&state.index, 0)
	volatile.StoreInt32(&state.done, 0)
	volatile.StoreInt32(&state.dir, 0)

	q.clearFlags()

	if template.AltLines != LinesNone {
		quadspi.Quadspi.Abr.SetAlternate(template.AltBytes)
	}

	// Only TEIE is armed. No FTIE (no FIFO servicing — the CPU consumes data
	// via AHB loads at 0x9000_0000). No TCIE (memory-mapped never finishes).
	// No SMIE (this isn't auto-poll).
	quadspi.Quadspi.Cr.SetFtie(false)
	quadspi.Quadspi.Cr.SetTcie(false)
	quadspi.Quadspi.Cr.SetSmie(false)
	quadspi.Quadspi.Cr.SetTeie(true)

	q.writeCCR(template, uint8(ModeMemoryMapped))
	return nil
}

// Unmap exits memory-mapped mode. Safe to call when not memory-mapped.
func (q QSpi) Unmap() error {
	if !q.IsMemoryMapped() {
		return nil
	}
	return q.Abort()
}

// AutoPoll repeatedly issues cmd and waits for (status & mask) == match. The
// peripheral does the polling in hardware, so the goroutine only yields on
// state.done. intervalCycles is the gap between polls in QSPI clock cycles.
//
// cmd is typically a status-register read (e.g. SR1 with mask=BUSY, match=0).
// cmd.Data length sets the number of status bytes per poll (1..4); the bytes
// are concatenated into a 32-bit value compared against match.
func (q QSpi) AutoPoll(cmd *Command, match, mask uint32, intervalCycles uint16, timeout time.Duration) error {
	state := &driverState

	if !q.Enabled() {
		return corehal.ErrNotReady
	}
	if cmd.Direction != DirRead {
		return corehal.ErrInvalidConfig
	}
	if len(cmd.Data) < 1 || len(cmd.Data) > 4 {
		return corehal.ErrInvalidConfig
	}

	state.mutex.Lock()

	if q.IsMemoryMapped() {
		if err := q.abortImpl(); err != nil {
			state.mutex.Unlock()
			return err
		}
	}
	if err := q.waitNotBusy(timeout); err != nil {
		state.mutex.Unlock()
		return err
	}

	// Reset transaction state. Auto-poll consumes status bytes inside the
	// peripheral — no FIFO servicing from the ISR, so dir stays DirNone.
	state.data = nil
	state.lastError = nil
	volatile.StoreInt32(&state.index, 0)
	volatile.StoreInt32(&state.done, 0)
	volatile.StoreInt32(&state.dir, 0)

	q.clearFlags()

	quadspi.Quadspi.Psmar.SetMatch(match)
	quadspi.Quadspi.Psmkr.SetMask(mask)
	quadspi.Quadspi.Pir.SetInterval(intervalCycles)
	quadspi.Quadspi.Dlr.SetDl(uint32(len(cmd.Data) - 1))

	// APMS=1 (stop on match), PMM=0 (AND polarity / match-when-equal).
	quadspi.Quadspi.Cr.SetApms(true)
	quadspi.Quadspi.Cr.SetPmm(false)

	if cmd.AltLines != LinesNone {
		quadspi.Quadspi.Abr.SetAlternate(cmd.AltBytes)
	}

	// Arm SMIE + TEIE before triggering. FTIE/TCIE stay disabled — auto-poll
	// doesn't move data through the FIFO and doesn't assert TCF on match.
	quadspi.Quadspi.Cr.SetFtie(false)
	quadspi.Quadspi.Cr.SetTcie(false)
	quadspi.Quadspi.Cr.SetSmie(true)
	quadspi.Quadspi.Cr.SetTeie(true)

	q.writeCCR(cmd, uint8(ModeAutoPoll))

	if cmd.AddrLines != LinesNone {
		quadspi.Quadspi.Ar.SetAddress(cmd.Address)
	}

	// Wait for ISR completion.
	if timeout == 0 {
		timeout = defaultTimeout
	}
	deadline := time.Now().Add(timeout)
	for volatile.LoadInt32(&state.done) == 0 {
		if time.Now().After(deadline) {
			quadspi.Quadspi.Cr.SetSmie(false)
			quadspi.Quadspi.Cr.SetTeie(false)
			_ = q.abortImpl()
			volatile.StoreInt32(&state.done, 1)
			state.mutex.Unlock()
			return ErrTimeout
		}
		runtime.Gosched()
	}

	if cmd.Direction != DirNone && int(volatile.LoadInt32(&state.index)) != len(cmd.Data) {
		state.mutex.Unlock()
		return ErrTransfer
	}

	state.mutex.Unlock()
	return state.lastError
}

// Abort halts any in-flight transaction and exits memory-mapped mode.
func (q QSpi) Abort() error {
	driverState.mutex.Lock()
	err := q.abortImpl()
	driverState.mutex.Unlock()
	return err
}

func (q QSpi) abortImpl() error {
	quadspi.Quadspi.Cr.SetAbort(true)

	deadline := time.Now().Add(defaultTimeout)
	for quadspi.Quadspi.Cr.GetAbort() {
		if time.Now().After(deadline) {
			return ErrTimeout
		}
	}
	// BUSY drops shortly after ABORT clears.
	return q.waitNotBusy(defaultTimeout)
}

func (q QSpi) writeCCR(cmd *Command, fmode uint8) {
	var ccr quadspi.RegisterCcrType

	ccr.SetFmode(fmode)

	if cmd.InstrLines != LinesNone {
		ccr.SetImode(uint8(cmd.InstrLines))
		ccr.SetInstruction(cmd.Instruction)
	}
	if cmd.AddrLines != LinesNone {
		ccr.SetAdmode(uint8(cmd.AddrLines))
		ccr.SetAdsize(uint8(cmd.AddrSize))
	}
	if cmd.AltLines != LinesNone {
		ccr.SetAbmode(uint8(cmd.AltLines))
		ccr.SetAbsize(uint8(cmd.AltSize))
	}
	ccr.SetDcyc(cmd.DummyCycles & 0x1F)
	if cmd.DataLines != LinesNone && cmd.Direction != DirNone {
		ccr.SetDmode(uint8(cmd.DataLines))
	}
	ccr.SetSioo(cmd.SIOO)
	ccr.SetDdrm(cmd.DDR)

	quadspi.Quadspi.Ccr.Store(uint32(ccr))
}

func (q QSpi) waitNotBusy(timeout time.Duration) error {
	if timeout == 0 {
		timeout = defaultTimeout
	}
	deadline := time.Now().Add(timeout)
	for quadspi.Quadspi.Sr.GetBusy() {
		if time.Now().After(deadline) {
			return ErrTimeout
		}
	}
	return nil
}

func (q QSpi) clearFlags() {
	// Clear TOF, SMF, TCF, TEF. (FTF is FIFO-driven, not clearable.)
	quadspi.Quadspi.Fcr.StoreBits(0x1B)
}

func (c *Command) timeout() time.Duration {
	if c.Timeout == 0 {
		return defaultTimeout
	}
	return c.Timeout
}
