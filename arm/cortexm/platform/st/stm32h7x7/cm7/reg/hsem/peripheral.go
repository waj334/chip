//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package hsem

import (
	"unsafe"
	"volatile"
)

var (
	Hsem = (*_hsem)(unsafe.Pointer(uintptr(0x58026400)))
)

type _hsem struct {
	Hsem_r0     registerHsem_r0Type
	Hsem_r1     registerHsem_r1Type
	Hsem_r2     registerHsem_r2Type
	Hsem_r3     registerHsem_r3Type
	Hsem_r4     registerHsem_r4Type
	Hsem_r5     registerHsem_r5Type
	Hsem_r6     registerHsem_r6Type
	Hsem_r7     registerHsem_r7Type
	Hsem_r8     registerHsem_r8Type
	Hsem_r9     registerHsem_r9Type
	Hsem_r10    registerHsem_r10Type
	Hsem_r11    registerHsem_r11Type
	Hsem_r12    registerHsem_r12Type
	Hsem_r13    registerHsem_r13Type
	Hsem_r14    registerHsem_r14Type
	Hsem_r15    registerHsem_r15Type
	Hsem_r16    registerHsem_r16Type
	Hsem_r17    registerHsem_r17Type
	Hsem_r18    registerHsem_r18Type
	Hsem_r19    registerHsem_r19Type
	Hsem_r20    registerHsem_r20Type
	Hsem_r21    registerHsem_r21Type
	Hsem_r22    registerHsem_r22Type
	Hsem_r23    registerHsem_r23Type
	Hsem_r24    registerHsem_r24Type
	Hsem_r25    registerHsem_r25Type
	Hsem_r26    registerHsem_r26Type
	Hsem_r27    registerHsem_r27Type
	Hsem_r28    registerHsem_r28Type
	Hsem_r29    registerHsem_r29Type
	Hsem_r30    registerHsem_r30Type
	Hsem_r31    registerHsem_r31Type
	Hsem_rlr0   registerHsem_rlr0Type
	Hsem_rlr1   registerHsem_rlr1Type
	Hsem_rlr2   registerHsem_rlr2Type
	Hsem_rlr3   registerHsem_rlr3Type
	Hsem_rlr4   registerHsem_rlr4Type
	Hsem_rlr5   registerHsem_rlr5Type
	Hsem_rlr6   registerHsem_rlr6Type
	Hsem_rlr7   registerHsem_rlr7Type
	Hsem_rlr8   registerHsem_rlr8Type
	Hsem_rlr9   registerHsem_rlr9Type
	Hsem_rlr10  registerHsem_rlr10Type
	Hsem_rlr11  registerHsem_rlr11Type
	Hsem_rlr12  registerHsem_rlr12Type
	Hsem_rlr13  registerHsem_rlr13Type
	Hsem_rlr14  registerHsem_rlr14Type
	Hsem_rlr15  registerHsem_rlr15Type
	Hsem_rlr16  registerHsem_rlr16Type
	Hsem_rlr17  registerHsem_rlr17Type
	Hsem_rlr18  registerHsem_rlr18Type
	Hsem_rlr19  registerHsem_rlr19Type
	Hsem_rlr20  registerHsem_rlr20Type
	Hsem_rlr21  registerHsem_rlr21Type
	Hsem_rlr22  registerHsem_rlr22Type
	Hsem_rlr23  registerHsem_rlr23Type
	Hsem_rlr24  registerHsem_rlr24Type
	Hsem_rlr25  registerHsem_rlr25Type
	Hsem_rlr26  registerHsem_rlr26Type
	Hsem_rlr27  registerHsem_rlr27Type
	Hsem_rlr28  registerHsem_rlr28Type
	Hsem_rlr29  registerHsem_rlr29Type
	Hsem_rlr30  registerHsem_rlr30Type
	Hsem_rlr31  registerHsem_rlr31Type
	Hsem_c1ier  registerHsem_c1ierType
	Hsem_c1icr  registerHsem_c1icrType
	Hsem_c1isr  registerHsem_c1isrType
	Hsem_c1misr registerHsem_c1misrType
	Hsem_c2ier  registerHsem_c2ierType
	Hsem_c2icr  registerHsem_c2icrType
	Hsem_c2isr  registerHsem_c2isrType
	Hsem_c2misr registerHsem_c2misrType
	_           [32]byte
	Hsem_cr     registerHsem_crType
	Hsem_keyr   registerHsem_keyrType
}

// registerHsem_r0Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r0Type uint32

const (
	RegisterHsem_r0FieldProcidShift = 0
	RegisterHsem_r0FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r0Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r0FieldProcidMask) >> RegisterHsem_r0FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r0Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r0FieldProcidMask)|(uint32(value)<<RegisterHsem_r0FieldProcidShift))
}

const (
	RegisterHsem_r0FieldCoreidShift = 8
	RegisterHsem_r0FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r0Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r0FieldCoreidMask) >> RegisterHsem_r0FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r0Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r0FieldCoreidMask)|(uint32(value)<<RegisterHsem_r0FieldCoreidShift))
}

const (
	RegisterHsem_r0FieldLockShift = 31
	RegisterHsem_r0FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r0Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r0FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r0Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r0FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r0FieldLockMask)
	}
}

// registerHsem_r1Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r1Type uint32

const (
	RegisterHsem_r1FieldProcidShift = 0
	RegisterHsem_r1FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r1Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r1FieldProcidMask) >> RegisterHsem_r1FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r1Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r1FieldProcidMask)|(uint32(value)<<RegisterHsem_r1FieldProcidShift))
}

const (
	RegisterHsem_r1FieldCoreidShift = 8
	RegisterHsem_r1FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r1Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r1FieldCoreidMask) >> RegisterHsem_r1FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r1Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r1FieldCoreidMask)|(uint32(value)<<RegisterHsem_r1FieldCoreidShift))
}

const (
	RegisterHsem_r1FieldLockShift = 31
	RegisterHsem_r1FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r1Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r1FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r1Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r1FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r1FieldLockMask)
	}
}

// registerHsem_r2Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r2Type uint32

const (
	RegisterHsem_r2FieldProcidShift = 0
	RegisterHsem_r2FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r2Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r2FieldProcidMask) >> RegisterHsem_r2FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r2Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r2FieldProcidMask)|(uint32(value)<<RegisterHsem_r2FieldProcidShift))
}

const (
	RegisterHsem_r2FieldCoreidShift = 8
	RegisterHsem_r2FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r2Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r2FieldCoreidMask) >> RegisterHsem_r2FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r2Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r2FieldCoreidMask)|(uint32(value)<<RegisterHsem_r2FieldCoreidShift))
}

const (
	RegisterHsem_r2FieldLockShift = 31
	RegisterHsem_r2FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r2Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r2FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r2Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r2FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r2FieldLockMask)
	}
}

// registerHsem_r3Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r3Type uint32

const (
	RegisterHsem_r3FieldProcidShift = 0
	RegisterHsem_r3FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r3Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r3FieldProcidMask) >> RegisterHsem_r3FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r3Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r3FieldProcidMask)|(uint32(value)<<RegisterHsem_r3FieldProcidShift))
}

const (
	RegisterHsem_r3FieldCoreidShift = 8
	RegisterHsem_r3FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r3Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r3FieldCoreidMask) >> RegisterHsem_r3FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r3Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r3FieldCoreidMask)|(uint32(value)<<RegisterHsem_r3FieldCoreidShift))
}

const (
	RegisterHsem_r3FieldLockShift = 31
	RegisterHsem_r3FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r3Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r3FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r3Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r3FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r3FieldLockMask)
	}
}

// registerHsem_r4Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r4Type uint32

const (
	RegisterHsem_r4FieldProcidShift = 0
	RegisterHsem_r4FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r4Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r4FieldProcidMask) >> RegisterHsem_r4FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r4Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r4FieldProcidMask)|(uint32(value)<<RegisterHsem_r4FieldProcidShift))
}

const (
	RegisterHsem_r4FieldCoreidShift = 8
	RegisterHsem_r4FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r4Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r4FieldCoreidMask) >> RegisterHsem_r4FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r4Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r4FieldCoreidMask)|(uint32(value)<<RegisterHsem_r4FieldCoreidShift))
}

const (
	RegisterHsem_r4FieldLockShift = 31
	RegisterHsem_r4FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r4Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r4FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r4Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r4FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r4FieldLockMask)
	}
}

// registerHsem_r5Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r5Type uint32

const (
	RegisterHsem_r5FieldProcidShift = 0
	RegisterHsem_r5FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r5Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r5FieldProcidMask) >> RegisterHsem_r5FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r5Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r5FieldProcidMask)|(uint32(value)<<RegisterHsem_r5FieldProcidShift))
}

const (
	RegisterHsem_r5FieldCoreidShift = 8
	RegisterHsem_r5FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r5Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r5FieldCoreidMask) >> RegisterHsem_r5FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r5Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r5FieldCoreidMask)|(uint32(value)<<RegisterHsem_r5FieldCoreidShift))
}

const (
	RegisterHsem_r5FieldLockShift = 31
	RegisterHsem_r5FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r5Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r5FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r5Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r5FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r5FieldLockMask)
	}
}

// registerHsem_r6Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r6Type uint32

const (
	RegisterHsem_r6FieldProcidShift = 0
	RegisterHsem_r6FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r6Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r6FieldProcidMask) >> RegisterHsem_r6FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r6Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r6FieldProcidMask)|(uint32(value)<<RegisterHsem_r6FieldProcidShift))
}

const (
	RegisterHsem_r6FieldCoreidShift = 8
	RegisterHsem_r6FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r6Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r6FieldCoreidMask) >> RegisterHsem_r6FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r6Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r6FieldCoreidMask)|(uint32(value)<<RegisterHsem_r6FieldCoreidShift))
}

const (
	RegisterHsem_r6FieldLockShift = 31
	RegisterHsem_r6FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r6Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r6FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r6Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r6FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r6FieldLockMask)
	}
}

// registerHsem_r7Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r7Type uint32

const (
	RegisterHsem_r7FieldProcidShift = 0
	RegisterHsem_r7FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r7Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r7FieldProcidMask) >> RegisterHsem_r7FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r7Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r7FieldProcidMask)|(uint32(value)<<RegisterHsem_r7FieldProcidShift))
}

const (
	RegisterHsem_r7FieldCoreidShift = 8
	RegisterHsem_r7FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r7Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r7FieldCoreidMask) >> RegisterHsem_r7FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r7Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r7FieldCoreidMask)|(uint32(value)<<RegisterHsem_r7FieldCoreidShift))
}

const (
	RegisterHsem_r7FieldLockShift = 31
	RegisterHsem_r7FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r7Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r7FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r7Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r7FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r7FieldLockMask)
	}
}

// registerHsem_r8Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r8Type uint32

const (
	RegisterHsem_r8FieldProcidShift = 0
	RegisterHsem_r8FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r8Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r8FieldProcidMask) >> RegisterHsem_r8FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r8Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r8FieldProcidMask)|(uint32(value)<<RegisterHsem_r8FieldProcidShift))
}

const (
	RegisterHsem_r8FieldCoreidShift = 8
	RegisterHsem_r8FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r8Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r8FieldCoreidMask) >> RegisterHsem_r8FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r8Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r8FieldCoreidMask)|(uint32(value)<<RegisterHsem_r8FieldCoreidShift))
}

const (
	RegisterHsem_r8FieldLockShift = 31
	RegisterHsem_r8FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r8Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r8FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r8Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r8FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r8FieldLockMask)
	}
}

// registerHsem_r9Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r9Type uint32

const (
	RegisterHsem_r9FieldProcidShift = 0
	RegisterHsem_r9FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r9Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r9FieldProcidMask) >> RegisterHsem_r9FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r9Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r9FieldProcidMask)|(uint32(value)<<RegisterHsem_r9FieldProcidShift))
}

const (
	RegisterHsem_r9FieldCoreidShift = 8
	RegisterHsem_r9FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r9Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r9FieldCoreidMask) >> RegisterHsem_r9FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r9Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r9FieldCoreidMask)|(uint32(value)<<RegisterHsem_r9FieldCoreidShift))
}

const (
	RegisterHsem_r9FieldLockShift = 31
	RegisterHsem_r9FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r9Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r9FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r9Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r9FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r9FieldLockMask)
	}
}

// registerHsem_r10Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r10Type uint32

const (
	RegisterHsem_r10FieldProcidShift = 0
	RegisterHsem_r10FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r10Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r10FieldProcidMask) >> RegisterHsem_r10FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r10Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r10FieldProcidMask)|(uint32(value)<<RegisterHsem_r10FieldProcidShift))
}

const (
	RegisterHsem_r10FieldCoreidShift = 8
	RegisterHsem_r10FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r10Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r10FieldCoreidMask) >> RegisterHsem_r10FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r10Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r10FieldCoreidMask)|(uint32(value)<<RegisterHsem_r10FieldCoreidShift))
}

const (
	RegisterHsem_r10FieldLockShift = 31
	RegisterHsem_r10FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r10Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r10FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r10Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r10FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r10FieldLockMask)
	}
}

// registerHsem_r11Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r11Type uint32

const (
	RegisterHsem_r11FieldProcidShift = 0
	RegisterHsem_r11FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r11Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r11FieldProcidMask) >> RegisterHsem_r11FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r11Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r11FieldProcidMask)|(uint32(value)<<RegisterHsem_r11FieldProcidShift))
}

const (
	RegisterHsem_r11FieldCoreidShift = 8
	RegisterHsem_r11FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r11Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r11FieldCoreidMask) >> RegisterHsem_r11FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r11Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r11FieldCoreidMask)|(uint32(value)<<RegisterHsem_r11FieldCoreidShift))
}

const (
	RegisterHsem_r11FieldLockShift = 31
	RegisterHsem_r11FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r11Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r11FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r11Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r11FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r11FieldLockMask)
	}
}

// registerHsem_r12Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r12Type uint32

const (
	RegisterHsem_r12FieldProcidShift = 0
	RegisterHsem_r12FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r12Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r12FieldProcidMask) >> RegisterHsem_r12FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r12Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r12FieldProcidMask)|(uint32(value)<<RegisterHsem_r12FieldProcidShift))
}

const (
	RegisterHsem_r12FieldCoreidShift = 8
	RegisterHsem_r12FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r12Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r12FieldCoreidMask) >> RegisterHsem_r12FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r12Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r12FieldCoreidMask)|(uint32(value)<<RegisterHsem_r12FieldCoreidShift))
}

const (
	RegisterHsem_r12FieldLockShift = 31
	RegisterHsem_r12FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r12Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r12FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r12Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r12FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r12FieldLockMask)
	}
}

// registerHsem_r13Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r13Type uint32

const (
	RegisterHsem_r13FieldProcidShift = 0
	RegisterHsem_r13FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r13Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r13FieldProcidMask) >> RegisterHsem_r13FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r13Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r13FieldProcidMask)|(uint32(value)<<RegisterHsem_r13FieldProcidShift))
}

const (
	RegisterHsem_r13FieldCoreidShift = 8
	RegisterHsem_r13FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r13Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r13FieldCoreidMask) >> RegisterHsem_r13FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r13Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r13FieldCoreidMask)|(uint32(value)<<RegisterHsem_r13FieldCoreidShift))
}

const (
	RegisterHsem_r13FieldLockShift = 31
	RegisterHsem_r13FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r13Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r13FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r13Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r13FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r13FieldLockMask)
	}
}

// registerHsem_r14Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r14Type uint32

const (
	RegisterHsem_r14FieldProcidShift = 0
	RegisterHsem_r14FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r14Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r14FieldProcidMask) >> RegisterHsem_r14FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r14Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r14FieldProcidMask)|(uint32(value)<<RegisterHsem_r14FieldProcidShift))
}

const (
	RegisterHsem_r14FieldCoreidShift = 8
	RegisterHsem_r14FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r14Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r14FieldCoreidMask) >> RegisterHsem_r14FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r14Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r14FieldCoreidMask)|(uint32(value)<<RegisterHsem_r14FieldCoreidShift))
}

const (
	RegisterHsem_r14FieldLockShift = 31
	RegisterHsem_r14FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r14Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r14FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r14Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r14FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r14FieldLockMask)
	}
}

// registerHsem_r15Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r15Type uint32

const (
	RegisterHsem_r15FieldProcidShift = 0
	RegisterHsem_r15FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r15Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r15FieldProcidMask) >> RegisterHsem_r15FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r15Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r15FieldProcidMask)|(uint32(value)<<RegisterHsem_r15FieldProcidShift))
}

const (
	RegisterHsem_r15FieldCoreidShift = 8
	RegisterHsem_r15FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r15Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r15FieldCoreidMask) >> RegisterHsem_r15FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r15Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r15FieldCoreidMask)|(uint32(value)<<RegisterHsem_r15FieldCoreidShift))
}

const (
	RegisterHsem_r15FieldLockShift = 31
	RegisterHsem_r15FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r15Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r15FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r15Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r15FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r15FieldLockMask)
	}
}

// registerHsem_r16Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r16Type uint32

const (
	RegisterHsem_r16FieldProcidShift = 0
	RegisterHsem_r16FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r16Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r16FieldProcidMask) >> RegisterHsem_r16FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r16Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r16FieldProcidMask)|(uint32(value)<<RegisterHsem_r16FieldProcidShift))
}

const (
	RegisterHsem_r16FieldCoreidShift = 8
	RegisterHsem_r16FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r16Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r16FieldCoreidMask) >> RegisterHsem_r16FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r16Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r16FieldCoreidMask)|(uint32(value)<<RegisterHsem_r16FieldCoreidShift))
}

const (
	RegisterHsem_r16FieldLockShift = 31
	RegisterHsem_r16FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r16Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r16FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r16Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r16FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r16FieldLockMask)
	}
}

// registerHsem_r17Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r17Type uint32

const (
	RegisterHsem_r17FieldProcidShift = 0
	RegisterHsem_r17FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r17Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r17FieldProcidMask) >> RegisterHsem_r17FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r17Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r17FieldProcidMask)|(uint32(value)<<RegisterHsem_r17FieldProcidShift))
}

const (
	RegisterHsem_r17FieldCoreidShift = 8
	RegisterHsem_r17FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r17Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r17FieldCoreidMask) >> RegisterHsem_r17FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r17Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r17FieldCoreidMask)|(uint32(value)<<RegisterHsem_r17FieldCoreidShift))
}

const (
	RegisterHsem_r17FieldLockShift = 31
	RegisterHsem_r17FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r17Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r17FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r17Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r17FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r17FieldLockMask)
	}
}

// registerHsem_r18Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r18Type uint32

const (
	RegisterHsem_r18FieldProcidShift = 0
	RegisterHsem_r18FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r18Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r18FieldProcidMask) >> RegisterHsem_r18FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r18Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r18FieldProcidMask)|(uint32(value)<<RegisterHsem_r18FieldProcidShift))
}

const (
	RegisterHsem_r18FieldCoreidShift = 8
	RegisterHsem_r18FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r18Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r18FieldCoreidMask) >> RegisterHsem_r18FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r18Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r18FieldCoreidMask)|(uint32(value)<<RegisterHsem_r18FieldCoreidShift))
}

const (
	RegisterHsem_r18FieldLockShift = 31
	RegisterHsem_r18FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r18Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r18FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r18Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r18FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r18FieldLockMask)
	}
}

// registerHsem_r19Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r19Type uint32

const (
	RegisterHsem_r19FieldProcidShift = 0
	RegisterHsem_r19FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r19Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r19FieldProcidMask) >> RegisterHsem_r19FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r19Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r19FieldProcidMask)|(uint32(value)<<RegisterHsem_r19FieldProcidShift))
}

const (
	RegisterHsem_r19FieldCoreidShift = 8
	RegisterHsem_r19FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r19Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r19FieldCoreidMask) >> RegisterHsem_r19FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r19Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r19FieldCoreidMask)|(uint32(value)<<RegisterHsem_r19FieldCoreidShift))
}

const (
	RegisterHsem_r19FieldLockShift = 31
	RegisterHsem_r19FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r19Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r19FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r19Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r19FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r19FieldLockMask)
	}
}

// registerHsem_r20Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r20Type uint32

const (
	RegisterHsem_r20FieldProcidShift = 0
	RegisterHsem_r20FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r20Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r20FieldProcidMask) >> RegisterHsem_r20FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r20Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r20FieldProcidMask)|(uint32(value)<<RegisterHsem_r20FieldProcidShift))
}

const (
	RegisterHsem_r20FieldCoreidShift = 8
	RegisterHsem_r20FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r20Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r20FieldCoreidMask) >> RegisterHsem_r20FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r20Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r20FieldCoreidMask)|(uint32(value)<<RegisterHsem_r20FieldCoreidShift))
}

const (
	RegisterHsem_r20FieldLockShift = 31
	RegisterHsem_r20FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r20Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r20FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r20Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r20FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r20FieldLockMask)
	}
}

// registerHsem_r21Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r21Type uint32

const (
	RegisterHsem_r21FieldProcidShift = 0
	RegisterHsem_r21FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r21Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r21FieldProcidMask) >> RegisterHsem_r21FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r21Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r21FieldProcidMask)|(uint32(value)<<RegisterHsem_r21FieldProcidShift))
}

const (
	RegisterHsem_r21FieldCoreidShift = 8
	RegisterHsem_r21FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r21Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r21FieldCoreidMask) >> RegisterHsem_r21FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r21Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r21FieldCoreidMask)|(uint32(value)<<RegisterHsem_r21FieldCoreidShift))
}

const (
	RegisterHsem_r21FieldLockShift = 31
	RegisterHsem_r21FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r21Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r21FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r21Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r21FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r21FieldLockMask)
	}
}

// registerHsem_r22Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r22Type uint32

const (
	RegisterHsem_r22FieldProcidShift = 0
	RegisterHsem_r22FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r22Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r22FieldProcidMask) >> RegisterHsem_r22FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r22Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r22FieldProcidMask)|(uint32(value)<<RegisterHsem_r22FieldProcidShift))
}

const (
	RegisterHsem_r22FieldCoreidShift = 8
	RegisterHsem_r22FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r22Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r22FieldCoreidMask) >> RegisterHsem_r22FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r22Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r22FieldCoreidMask)|(uint32(value)<<RegisterHsem_r22FieldCoreidShift))
}

const (
	RegisterHsem_r22FieldLockShift = 31
	RegisterHsem_r22FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r22Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r22FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r22Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r22FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r22FieldLockMask)
	}
}

// registerHsem_r23Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r23Type uint32

const (
	RegisterHsem_r23FieldProcidShift = 0
	RegisterHsem_r23FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r23Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r23FieldProcidMask) >> RegisterHsem_r23FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r23Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r23FieldProcidMask)|(uint32(value)<<RegisterHsem_r23FieldProcidShift))
}

const (
	RegisterHsem_r23FieldCoreidShift = 8
	RegisterHsem_r23FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r23Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r23FieldCoreidMask) >> RegisterHsem_r23FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r23Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r23FieldCoreidMask)|(uint32(value)<<RegisterHsem_r23FieldCoreidShift))
}

const (
	RegisterHsem_r23FieldLockShift = 31
	RegisterHsem_r23FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r23Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r23FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r23Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r23FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r23FieldLockMask)
	}
}

// registerHsem_r24Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r24Type uint32

const (
	RegisterHsem_r24FieldProcidShift = 0
	RegisterHsem_r24FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r24Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r24FieldProcidMask) >> RegisterHsem_r24FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r24Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r24FieldProcidMask)|(uint32(value)<<RegisterHsem_r24FieldProcidShift))
}

const (
	RegisterHsem_r24FieldCoreidShift = 8
	RegisterHsem_r24FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r24Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r24FieldCoreidMask) >> RegisterHsem_r24FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r24Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r24FieldCoreidMask)|(uint32(value)<<RegisterHsem_r24FieldCoreidShift))
}

const (
	RegisterHsem_r24FieldLockShift = 31
	RegisterHsem_r24FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r24Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r24FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r24Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r24FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r24FieldLockMask)
	}
}

// registerHsem_r25Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r25Type uint32

const (
	RegisterHsem_r25FieldProcidShift = 0
	RegisterHsem_r25FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r25Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r25FieldProcidMask) >> RegisterHsem_r25FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r25Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r25FieldProcidMask)|(uint32(value)<<RegisterHsem_r25FieldProcidShift))
}

const (
	RegisterHsem_r25FieldCoreidShift = 8
	RegisterHsem_r25FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r25Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r25FieldCoreidMask) >> RegisterHsem_r25FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r25Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r25FieldCoreidMask)|(uint32(value)<<RegisterHsem_r25FieldCoreidShift))
}

const (
	RegisterHsem_r25FieldLockShift = 31
	RegisterHsem_r25FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r25Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r25FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r25Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r25FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r25FieldLockMask)
	}
}

// registerHsem_r26Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r26Type uint32

const (
	RegisterHsem_r26FieldProcidShift = 0
	RegisterHsem_r26FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r26Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r26FieldProcidMask) >> RegisterHsem_r26FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r26Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r26FieldProcidMask)|(uint32(value)<<RegisterHsem_r26FieldProcidShift))
}

const (
	RegisterHsem_r26FieldCoreidShift = 8
	RegisterHsem_r26FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r26Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r26FieldCoreidMask) >> RegisterHsem_r26FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r26Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r26FieldCoreidMask)|(uint32(value)<<RegisterHsem_r26FieldCoreidShift))
}

const (
	RegisterHsem_r26FieldLockShift = 31
	RegisterHsem_r26FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r26Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r26FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r26Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r26FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r26FieldLockMask)
	}
}

// registerHsem_r27Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r27Type uint32

const (
	RegisterHsem_r27FieldProcidShift = 0
	RegisterHsem_r27FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r27Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r27FieldProcidMask) >> RegisterHsem_r27FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r27Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r27FieldProcidMask)|(uint32(value)<<RegisterHsem_r27FieldProcidShift))
}

const (
	RegisterHsem_r27FieldCoreidShift = 8
	RegisterHsem_r27FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r27Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r27FieldCoreidMask) >> RegisterHsem_r27FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r27Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r27FieldCoreidMask)|(uint32(value)<<RegisterHsem_r27FieldCoreidShift))
}

const (
	RegisterHsem_r27FieldLockShift = 31
	RegisterHsem_r27FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r27Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r27FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r27Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r27FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r27FieldLockMask)
	}
}

// registerHsem_r28Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r28Type uint32

const (
	RegisterHsem_r28FieldProcidShift = 0
	RegisterHsem_r28FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r28Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r28FieldProcidMask) >> RegisterHsem_r28FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r28Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r28FieldProcidMask)|(uint32(value)<<RegisterHsem_r28FieldProcidShift))
}

const (
	RegisterHsem_r28FieldCoreidShift = 8
	RegisterHsem_r28FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r28Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r28FieldCoreidMask) >> RegisterHsem_r28FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r28Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r28FieldCoreidMask)|(uint32(value)<<RegisterHsem_r28FieldCoreidShift))
}

const (
	RegisterHsem_r28FieldLockShift = 31
	RegisterHsem_r28FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r28Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r28FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r28Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r28FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r28FieldLockMask)
	}
}

// registerHsem_r29Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r29Type uint32

const (
	RegisterHsem_r29FieldProcidShift = 0
	RegisterHsem_r29FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r29Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r29FieldProcidMask) >> RegisterHsem_r29FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r29Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r29FieldProcidMask)|(uint32(value)<<RegisterHsem_r29FieldProcidShift))
}

const (
	RegisterHsem_r29FieldCoreidShift = 8
	RegisterHsem_r29FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r29Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r29FieldCoreidMask) >> RegisterHsem_r29FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r29Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r29FieldCoreidMask)|(uint32(value)<<RegisterHsem_r29FieldCoreidShift))
}

const (
	RegisterHsem_r29FieldLockShift = 31
	RegisterHsem_r29FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r29Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r29FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r29Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r29FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r29FieldLockMask)
	}
}

// registerHsem_r30Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r30Type uint32

const (
	RegisterHsem_r30FieldProcidShift = 0
	RegisterHsem_r30FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r30Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r30FieldProcidMask) >> RegisterHsem_r30FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r30Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r30FieldProcidMask)|(uint32(value)<<RegisterHsem_r30FieldProcidShift))
}

const (
	RegisterHsem_r30FieldCoreidShift = 8
	RegisterHsem_r30FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r30Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r30FieldCoreidMask) >> RegisterHsem_r30FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r30Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r30FieldCoreidMask)|(uint32(value)<<RegisterHsem_r30FieldCoreidShift))
}

const (
	RegisterHsem_r30FieldLockShift = 31
	RegisterHsem_r30FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r30Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r30FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r30Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r30FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r30FieldLockMask)
	}
}

// registerHsem_r31Type HSEM register HSEM_R0 HSEM_R31
type registerHsem_r31Type uint32

const (
	RegisterHsem_r31FieldProcidShift = 0
	RegisterHsem_r31FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_r31Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r31FieldProcidMask) >> RegisterHsem_r31FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_r31Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r31FieldProcidMask)|(uint32(value)<<RegisterHsem_r31FieldProcidShift))
}

const (
	RegisterHsem_r31FieldCoreidShift = 8
	RegisterHsem_r31FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_r31Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r31FieldCoreidMask) >> RegisterHsem_r31FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_r31Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r31FieldCoreidMask)|(uint32(value)<<RegisterHsem_r31FieldCoreidShift))
}

const (
	RegisterHsem_r31FieldLockShift = 31
	RegisterHsem_r31FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_r31Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r31FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_r31Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_r31FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r31FieldLockMask)
	}
}

// registerHsem_rlr0Type HSEM Read lock register
type registerHsem_rlr0Type uint32

const (
	RegisterHsem_rlr0FieldProcidShift = 0
	RegisterHsem_rlr0FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr0Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr0FieldProcidMask) >> RegisterHsem_rlr0FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr0Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr0FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr0FieldProcidShift))
}

const (
	RegisterHsem_rlr0FieldCoreidShift = 8
	RegisterHsem_rlr0FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr0Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr0FieldCoreidMask) >> RegisterHsem_rlr0FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr0Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr0FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr0FieldCoreidShift))
}

const (
	RegisterHsem_rlr0FieldLockShift = 31
	RegisterHsem_rlr0FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr0Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr0FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr0Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr0FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr0FieldLockMask)
	}
}

// registerHsem_rlr1Type HSEM Read lock register
type registerHsem_rlr1Type uint32

const (
	RegisterHsem_rlr1FieldProcidShift = 0
	RegisterHsem_rlr1FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr1Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr1FieldProcidMask) >> RegisterHsem_rlr1FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr1Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr1FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr1FieldProcidShift))
}

const (
	RegisterHsem_rlr1FieldCoreidShift = 8
	RegisterHsem_rlr1FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr1Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr1FieldCoreidMask) >> RegisterHsem_rlr1FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr1Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr1FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr1FieldCoreidShift))
}

const (
	RegisterHsem_rlr1FieldLockShift = 31
	RegisterHsem_rlr1FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr1Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr1FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr1Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr1FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr1FieldLockMask)
	}
}

// registerHsem_rlr2Type HSEM Read lock register
type registerHsem_rlr2Type uint32

const (
	RegisterHsem_rlr2FieldProcidShift = 0
	RegisterHsem_rlr2FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr2Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr2FieldProcidMask) >> RegisterHsem_rlr2FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr2Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr2FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr2FieldProcidShift))
}

const (
	RegisterHsem_rlr2FieldCoreidShift = 8
	RegisterHsem_rlr2FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr2Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr2FieldCoreidMask) >> RegisterHsem_rlr2FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr2Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr2FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr2FieldCoreidShift))
}

const (
	RegisterHsem_rlr2FieldLockShift = 31
	RegisterHsem_rlr2FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr2Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr2FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr2Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr2FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr2FieldLockMask)
	}
}

// registerHsem_rlr3Type HSEM Read lock register
type registerHsem_rlr3Type uint32

const (
	RegisterHsem_rlr3FieldProcidShift = 0
	RegisterHsem_rlr3FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr3Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr3FieldProcidMask) >> RegisterHsem_rlr3FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr3Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr3FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr3FieldProcidShift))
}

const (
	RegisterHsem_rlr3FieldCoreidShift = 8
	RegisterHsem_rlr3FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr3Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr3FieldCoreidMask) >> RegisterHsem_rlr3FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr3Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr3FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr3FieldCoreidShift))
}

const (
	RegisterHsem_rlr3FieldLockShift = 31
	RegisterHsem_rlr3FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr3Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr3FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr3Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr3FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr3FieldLockMask)
	}
}

// registerHsem_rlr4Type HSEM Read lock register
type registerHsem_rlr4Type uint32

const (
	RegisterHsem_rlr4FieldProcidShift = 0
	RegisterHsem_rlr4FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr4Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr4FieldProcidMask) >> RegisterHsem_rlr4FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr4Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr4FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr4FieldProcidShift))
}

const (
	RegisterHsem_rlr4FieldCoreidShift = 8
	RegisterHsem_rlr4FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr4Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr4FieldCoreidMask) >> RegisterHsem_rlr4FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr4Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr4FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr4FieldCoreidShift))
}

const (
	RegisterHsem_rlr4FieldLockShift = 31
	RegisterHsem_rlr4FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr4Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr4FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr4Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr4FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr4FieldLockMask)
	}
}

// registerHsem_rlr5Type HSEM Read lock register
type registerHsem_rlr5Type uint32

const (
	RegisterHsem_rlr5FieldProcidShift = 0
	RegisterHsem_rlr5FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr5Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr5FieldProcidMask) >> RegisterHsem_rlr5FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr5Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr5FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr5FieldProcidShift))
}

const (
	RegisterHsem_rlr5FieldCoreidShift = 8
	RegisterHsem_rlr5FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr5Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr5FieldCoreidMask) >> RegisterHsem_rlr5FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr5Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr5FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr5FieldCoreidShift))
}

const (
	RegisterHsem_rlr5FieldLockShift = 31
	RegisterHsem_rlr5FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr5Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr5FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr5Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr5FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr5FieldLockMask)
	}
}

// registerHsem_rlr6Type HSEM Read lock register
type registerHsem_rlr6Type uint32

const (
	RegisterHsem_rlr6FieldProcidShift = 0
	RegisterHsem_rlr6FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr6Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr6FieldProcidMask) >> RegisterHsem_rlr6FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr6Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr6FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr6FieldProcidShift))
}

const (
	RegisterHsem_rlr6FieldCoreidShift = 8
	RegisterHsem_rlr6FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr6Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr6FieldCoreidMask) >> RegisterHsem_rlr6FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr6Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr6FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr6FieldCoreidShift))
}

const (
	RegisterHsem_rlr6FieldLockShift = 31
	RegisterHsem_rlr6FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr6Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr6FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr6Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr6FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr6FieldLockMask)
	}
}

// registerHsem_rlr7Type HSEM Read lock register
type registerHsem_rlr7Type uint32

const (
	RegisterHsem_rlr7FieldProcidShift = 0
	RegisterHsem_rlr7FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr7Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr7FieldProcidMask) >> RegisterHsem_rlr7FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr7Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr7FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr7FieldProcidShift))
}

const (
	RegisterHsem_rlr7FieldCoreidShift = 8
	RegisterHsem_rlr7FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr7Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr7FieldCoreidMask) >> RegisterHsem_rlr7FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr7Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr7FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr7FieldCoreidShift))
}

const (
	RegisterHsem_rlr7FieldLockShift = 31
	RegisterHsem_rlr7FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr7Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr7FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr7Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr7FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr7FieldLockMask)
	}
}

// registerHsem_rlr8Type HSEM Read lock register
type registerHsem_rlr8Type uint32

const (
	RegisterHsem_rlr8FieldProcidShift = 0
	RegisterHsem_rlr8FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr8Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr8FieldProcidMask) >> RegisterHsem_rlr8FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr8Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr8FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr8FieldProcidShift))
}

const (
	RegisterHsem_rlr8FieldCoreidShift = 8
	RegisterHsem_rlr8FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr8Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr8FieldCoreidMask) >> RegisterHsem_rlr8FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr8Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr8FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr8FieldCoreidShift))
}

const (
	RegisterHsem_rlr8FieldLockShift = 31
	RegisterHsem_rlr8FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr8Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr8FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr8Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr8FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr8FieldLockMask)
	}
}

// registerHsem_rlr9Type HSEM Read lock register
type registerHsem_rlr9Type uint32

const (
	RegisterHsem_rlr9FieldProcidShift = 0
	RegisterHsem_rlr9FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr9Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr9FieldProcidMask) >> RegisterHsem_rlr9FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr9Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr9FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr9FieldProcidShift))
}

const (
	RegisterHsem_rlr9FieldCoreidShift = 8
	RegisterHsem_rlr9FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr9Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr9FieldCoreidMask) >> RegisterHsem_rlr9FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr9Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr9FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr9FieldCoreidShift))
}

const (
	RegisterHsem_rlr9FieldLockShift = 31
	RegisterHsem_rlr9FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr9Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr9FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr9Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr9FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr9FieldLockMask)
	}
}

// registerHsem_rlr10Type HSEM Read lock register
type registerHsem_rlr10Type uint32

const (
	RegisterHsem_rlr10FieldProcidShift = 0
	RegisterHsem_rlr10FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr10Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr10FieldProcidMask) >> RegisterHsem_rlr10FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr10Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr10FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr10FieldProcidShift))
}

const (
	RegisterHsem_rlr10FieldCoreidShift = 8
	RegisterHsem_rlr10FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr10Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr10FieldCoreidMask) >> RegisterHsem_rlr10FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr10Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr10FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr10FieldCoreidShift))
}

const (
	RegisterHsem_rlr10FieldLockShift = 31
	RegisterHsem_rlr10FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr10Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr10FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr10Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr10FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr10FieldLockMask)
	}
}

// registerHsem_rlr11Type HSEM Read lock register
type registerHsem_rlr11Type uint32

const (
	RegisterHsem_rlr11FieldProcidShift = 0
	RegisterHsem_rlr11FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr11Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr11FieldProcidMask) >> RegisterHsem_rlr11FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr11Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr11FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr11FieldProcidShift))
}

const (
	RegisterHsem_rlr11FieldCoreidShift = 8
	RegisterHsem_rlr11FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr11Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr11FieldCoreidMask) >> RegisterHsem_rlr11FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr11Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr11FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr11FieldCoreidShift))
}

const (
	RegisterHsem_rlr11FieldLockShift = 31
	RegisterHsem_rlr11FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr11Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr11FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr11Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr11FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr11FieldLockMask)
	}
}

// registerHsem_rlr12Type HSEM Read lock register
type registerHsem_rlr12Type uint32

const (
	RegisterHsem_rlr12FieldProcidShift = 0
	RegisterHsem_rlr12FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr12Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr12FieldProcidMask) >> RegisterHsem_rlr12FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr12Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr12FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr12FieldProcidShift))
}

const (
	RegisterHsem_rlr12FieldCoreidShift = 8
	RegisterHsem_rlr12FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr12Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr12FieldCoreidMask) >> RegisterHsem_rlr12FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr12Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr12FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr12FieldCoreidShift))
}

const (
	RegisterHsem_rlr12FieldLockShift = 31
	RegisterHsem_rlr12FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr12Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr12FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr12Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr12FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr12FieldLockMask)
	}
}

// registerHsem_rlr13Type HSEM Read lock register
type registerHsem_rlr13Type uint32

const (
	RegisterHsem_rlr13FieldProcidShift = 0
	RegisterHsem_rlr13FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr13Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr13FieldProcidMask) >> RegisterHsem_rlr13FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr13Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr13FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr13FieldProcidShift))
}

const (
	RegisterHsem_rlr13FieldCoreidShift = 8
	RegisterHsem_rlr13FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr13Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr13FieldCoreidMask) >> RegisterHsem_rlr13FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr13Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr13FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr13FieldCoreidShift))
}

const (
	RegisterHsem_rlr13FieldLockShift = 31
	RegisterHsem_rlr13FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr13Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr13FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr13Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr13FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr13FieldLockMask)
	}
}

// registerHsem_rlr14Type HSEM Read lock register
type registerHsem_rlr14Type uint32

const (
	RegisterHsem_rlr14FieldProcidShift = 0
	RegisterHsem_rlr14FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr14Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr14FieldProcidMask) >> RegisterHsem_rlr14FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr14Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr14FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr14FieldProcidShift))
}

const (
	RegisterHsem_rlr14FieldCoreidShift = 8
	RegisterHsem_rlr14FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr14Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr14FieldCoreidMask) >> RegisterHsem_rlr14FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr14Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr14FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr14FieldCoreidShift))
}

const (
	RegisterHsem_rlr14FieldLockShift = 31
	RegisterHsem_rlr14FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr14Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr14FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr14Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr14FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr14FieldLockMask)
	}
}

// registerHsem_rlr15Type HSEM Read lock register
type registerHsem_rlr15Type uint32

const (
	RegisterHsem_rlr15FieldProcidShift = 0
	RegisterHsem_rlr15FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr15Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr15FieldProcidMask) >> RegisterHsem_rlr15FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr15Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr15FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr15FieldProcidShift))
}

const (
	RegisterHsem_rlr15FieldCoreidShift = 8
	RegisterHsem_rlr15FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr15Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr15FieldCoreidMask) >> RegisterHsem_rlr15FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr15Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr15FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr15FieldCoreidShift))
}

const (
	RegisterHsem_rlr15FieldLockShift = 31
	RegisterHsem_rlr15FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr15Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr15FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr15Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr15FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr15FieldLockMask)
	}
}

// registerHsem_rlr16Type HSEM Read lock register
type registerHsem_rlr16Type uint32

const (
	RegisterHsem_rlr16FieldProcidShift = 0
	RegisterHsem_rlr16FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr16Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr16FieldProcidMask) >> RegisterHsem_rlr16FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr16Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr16FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr16FieldProcidShift))
}

const (
	RegisterHsem_rlr16FieldCoreidShift = 8
	RegisterHsem_rlr16FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr16Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr16FieldCoreidMask) >> RegisterHsem_rlr16FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr16Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr16FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr16FieldCoreidShift))
}

const (
	RegisterHsem_rlr16FieldLockShift = 31
	RegisterHsem_rlr16FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr16Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr16FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr16Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr16FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr16FieldLockMask)
	}
}

// registerHsem_rlr17Type HSEM Read lock register
type registerHsem_rlr17Type uint32

const (
	RegisterHsem_rlr17FieldProcidShift = 0
	RegisterHsem_rlr17FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr17Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr17FieldProcidMask) >> RegisterHsem_rlr17FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr17Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr17FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr17FieldProcidShift))
}

const (
	RegisterHsem_rlr17FieldCoreidShift = 8
	RegisterHsem_rlr17FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr17Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr17FieldCoreidMask) >> RegisterHsem_rlr17FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr17Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr17FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr17FieldCoreidShift))
}

const (
	RegisterHsem_rlr17FieldLockShift = 31
	RegisterHsem_rlr17FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr17Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr17FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr17Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr17FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr17FieldLockMask)
	}
}

// registerHsem_rlr18Type HSEM Read lock register
type registerHsem_rlr18Type uint32

const (
	RegisterHsem_rlr18FieldProcidShift = 0
	RegisterHsem_rlr18FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr18Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr18FieldProcidMask) >> RegisterHsem_rlr18FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr18Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr18FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr18FieldProcidShift))
}

const (
	RegisterHsem_rlr18FieldCoreidShift = 8
	RegisterHsem_rlr18FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr18Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr18FieldCoreidMask) >> RegisterHsem_rlr18FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr18Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr18FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr18FieldCoreidShift))
}

const (
	RegisterHsem_rlr18FieldLockShift = 31
	RegisterHsem_rlr18FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr18Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr18FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr18Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr18FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr18FieldLockMask)
	}
}

// registerHsem_rlr19Type HSEM Read lock register
type registerHsem_rlr19Type uint32

const (
	RegisterHsem_rlr19FieldProcidShift = 0
	RegisterHsem_rlr19FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr19Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr19FieldProcidMask) >> RegisterHsem_rlr19FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr19Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr19FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr19FieldProcidShift))
}

const (
	RegisterHsem_rlr19FieldCoreidShift = 8
	RegisterHsem_rlr19FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr19Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr19FieldCoreidMask) >> RegisterHsem_rlr19FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr19Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr19FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr19FieldCoreidShift))
}

const (
	RegisterHsem_rlr19FieldLockShift = 31
	RegisterHsem_rlr19FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr19Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr19FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr19Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr19FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr19FieldLockMask)
	}
}

// registerHsem_rlr20Type HSEM Read lock register
type registerHsem_rlr20Type uint32

const (
	RegisterHsem_rlr20FieldProcidShift = 0
	RegisterHsem_rlr20FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr20Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr20FieldProcidMask) >> RegisterHsem_rlr20FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr20Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr20FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr20FieldProcidShift))
}

const (
	RegisterHsem_rlr20FieldCoreidShift = 8
	RegisterHsem_rlr20FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr20Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr20FieldCoreidMask) >> RegisterHsem_rlr20FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr20Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr20FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr20FieldCoreidShift))
}

const (
	RegisterHsem_rlr20FieldLockShift = 31
	RegisterHsem_rlr20FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr20Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr20FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr20Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr20FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr20FieldLockMask)
	}
}

// registerHsem_rlr21Type HSEM Read lock register
type registerHsem_rlr21Type uint32

const (
	RegisterHsem_rlr21FieldProcidShift = 0
	RegisterHsem_rlr21FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr21Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr21FieldProcidMask) >> RegisterHsem_rlr21FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr21Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr21FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr21FieldProcidShift))
}

const (
	RegisterHsem_rlr21FieldCoreidShift = 8
	RegisterHsem_rlr21FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr21Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr21FieldCoreidMask) >> RegisterHsem_rlr21FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr21Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr21FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr21FieldCoreidShift))
}

const (
	RegisterHsem_rlr21FieldLockShift = 31
	RegisterHsem_rlr21FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr21Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr21FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr21Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr21FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr21FieldLockMask)
	}
}

// registerHsem_rlr22Type HSEM Read lock register
type registerHsem_rlr22Type uint32

const (
	RegisterHsem_rlr22FieldProcidShift = 0
	RegisterHsem_rlr22FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr22Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr22FieldProcidMask) >> RegisterHsem_rlr22FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr22Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr22FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr22FieldProcidShift))
}

const (
	RegisterHsem_rlr22FieldCoreidShift = 8
	RegisterHsem_rlr22FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr22Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr22FieldCoreidMask) >> RegisterHsem_rlr22FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr22Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr22FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr22FieldCoreidShift))
}

const (
	RegisterHsem_rlr22FieldLockShift = 31
	RegisterHsem_rlr22FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr22Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr22FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr22Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr22FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr22FieldLockMask)
	}
}

// registerHsem_rlr23Type HSEM Read lock register
type registerHsem_rlr23Type uint32

const (
	RegisterHsem_rlr23FieldProcidShift = 0
	RegisterHsem_rlr23FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr23Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr23FieldProcidMask) >> RegisterHsem_rlr23FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr23Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr23FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr23FieldProcidShift))
}

const (
	RegisterHsem_rlr23FieldCoreidShift = 8
	RegisterHsem_rlr23FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr23Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr23FieldCoreidMask) >> RegisterHsem_rlr23FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr23Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr23FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr23FieldCoreidShift))
}

const (
	RegisterHsem_rlr23FieldLockShift = 31
	RegisterHsem_rlr23FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr23Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr23FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr23Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr23FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr23FieldLockMask)
	}
}

// registerHsem_rlr24Type HSEM Read lock register
type registerHsem_rlr24Type uint32

const (
	RegisterHsem_rlr24FieldProcidShift = 0
	RegisterHsem_rlr24FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr24Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr24FieldProcidMask) >> RegisterHsem_rlr24FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr24Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr24FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr24FieldProcidShift))
}

const (
	RegisterHsem_rlr24FieldCoreidShift = 8
	RegisterHsem_rlr24FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr24Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr24FieldCoreidMask) >> RegisterHsem_rlr24FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr24Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr24FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr24FieldCoreidShift))
}

const (
	RegisterHsem_rlr24FieldLockShift = 31
	RegisterHsem_rlr24FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr24Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr24FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr24Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr24FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr24FieldLockMask)
	}
}

// registerHsem_rlr25Type HSEM Read lock register
type registerHsem_rlr25Type uint32

const (
	RegisterHsem_rlr25FieldProcidShift = 0
	RegisterHsem_rlr25FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr25Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr25FieldProcidMask) >> RegisterHsem_rlr25FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr25Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr25FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr25FieldProcidShift))
}

const (
	RegisterHsem_rlr25FieldCoreidShift = 8
	RegisterHsem_rlr25FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr25Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr25FieldCoreidMask) >> RegisterHsem_rlr25FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr25Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr25FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr25FieldCoreidShift))
}

const (
	RegisterHsem_rlr25FieldLockShift = 31
	RegisterHsem_rlr25FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr25Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr25FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr25Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr25FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr25FieldLockMask)
	}
}

// registerHsem_rlr26Type HSEM Read lock register
type registerHsem_rlr26Type uint32

const (
	RegisterHsem_rlr26FieldProcidShift = 0
	RegisterHsem_rlr26FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr26Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr26FieldProcidMask) >> RegisterHsem_rlr26FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr26Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr26FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr26FieldProcidShift))
}

const (
	RegisterHsem_rlr26FieldCoreidShift = 8
	RegisterHsem_rlr26FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr26Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr26FieldCoreidMask) >> RegisterHsem_rlr26FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr26Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr26FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr26FieldCoreidShift))
}

const (
	RegisterHsem_rlr26FieldLockShift = 31
	RegisterHsem_rlr26FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr26Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr26FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr26Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr26FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr26FieldLockMask)
	}
}

// registerHsem_rlr27Type HSEM Read lock register
type registerHsem_rlr27Type uint32

const (
	RegisterHsem_rlr27FieldProcidShift = 0
	RegisterHsem_rlr27FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr27Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr27FieldProcidMask) >> RegisterHsem_rlr27FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr27Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr27FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr27FieldProcidShift))
}

const (
	RegisterHsem_rlr27FieldCoreidShift = 8
	RegisterHsem_rlr27FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr27Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr27FieldCoreidMask) >> RegisterHsem_rlr27FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr27Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr27FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr27FieldCoreidShift))
}

const (
	RegisterHsem_rlr27FieldLockShift = 31
	RegisterHsem_rlr27FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr27Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr27FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr27Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr27FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr27FieldLockMask)
	}
}

// registerHsem_rlr28Type HSEM Read lock register
type registerHsem_rlr28Type uint32

const (
	RegisterHsem_rlr28FieldProcidShift = 0
	RegisterHsem_rlr28FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr28Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr28FieldProcidMask) >> RegisterHsem_rlr28FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr28Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr28FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr28FieldProcidShift))
}

const (
	RegisterHsem_rlr28FieldCoreidShift = 8
	RegisterHsem_rlr28FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr28Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr28FieldCoreidMask) >> RegisterHsem_rlr28FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr28Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr28FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr28FieldCoreidShift))
}

const (
	RegisterHsem_rlr28FieldLockShift = 31
	RegisterHsem_rlr28FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr28Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr28FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr28Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr28FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr28FieldLockMask)
	}
}

// registerHsem_rlr29Type HSEM Read lock register
type registerHsem_rlr29Type uint32

const (
	RegisterHsem_rlr29FieldProcidShift = 0
	RegisterHsem_rlr29FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr29Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr29FieldProcidMask) >> RegisterHsem_rlr29FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr29Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr29FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr29FieldProcidShift))
}

const (
	RegisterHsem_rlr29FieldCoreidShift = 8
	RegisterHsem_rlr29FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr29Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr29FieldCoreidMask) >> RegisterHsem_rlr29FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr29Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr29FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr29FieldCoreidShift))
}

const (
	RegisterHsem_rlr29FieldLockShift = 31
	RegisterHsem_rlr29FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr29Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr29FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr29Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr29FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr29FieldLockMask)
	}
}

// registerHsem_rlr30Type HSEM Read lock register
type registerHsem_rlr30Type uint32

const (
	RegisterHsem_rlr30FieldProcidShift = 0
	RegisterHsem_rlr30FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr30Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr30FieldProcidMask) >> RegisterHsem_rlr30FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr30Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr30FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr30FieldProcidShift))
}

const (
	RegisterHsem_rlr30FieldCoreidShift = 8
	RegisterHsem_rlr30FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr30Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr30FieldCoreidMask) >> RegisterHsem_rlr30FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr30Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr30FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr30FieldCoreidShift))
}

const (
	RegisterHsem_rlr30FieldLockShift = 31
	RegisterHsem_rlr30FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr30Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr30FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr30Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr30FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr30FieldLockMask)
	}
}

// registerHsem_rlr31Type HSEM Read lock register
type registerHsem_rlr31Type uint32

const (
	RegisterHsem_rlr31FieldProcidShift = 0
	RegisterHsem_rlr31FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsem_rlr31Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr31FieldProcidMask) >> RegisterHsem_rlr31FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsem_rlr31Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr31FieldProcidMask)|(uint32(value)<<RegisterHsem_rlr31FieldProcidShift))
}

const (
	RegisterHsem_rlr31FieldCoreidShift = 8
	RegisterHsem_rlr31FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsem_rlr31Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr31FieldCoreidMask) >> RegisterHsem_rlr31FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsem_rlr31Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr31FieldCoreidMask)|(uint32(value)<<RegisterHsem_rlr31FieldCoreidShift))
}

const (
	RegisterHsem_rlr31FieldLockShift = 31
	RegisterHsem_rlr31FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsem_rlr31Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr31FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsem_rlr31Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_rlr31FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr31FieldLockMask)
	}
}

// registerHsem_c1ierType HSEM Interrupt enable register
type registerHsem_c1ierType uint32

const (
	RegisterHsem_c1ierFieldIseShift = 0
	RegisterHsem_c1ierFieldIseMask  = 0xffffffff
)

// GetIse Interrupt semaphore x enable bit
func (r *registerHsem_c1ierType) GetIse() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_c1ierFieldIseMask) >> RegisterHsem_c1ierFieldIseShift)
}

// SetIse Interrupt semaphore x enable bit
func (r *registerHsem_c1ierType) SetIse(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_c1ierFieldIseMask)|(uint32(value)<<RegisterHsem_c1ierFieldIseShift))
}

// registerHsem_c1icrType HSEM Interrupt clear register
type registerHsem_c1icrType uint32

const (
	RegisterHsem_c1icrFieldIscShift = 0
	RegisterHsem_c1icrFieldIscMask  = 0xffffffff
)

// GetIsc Interrupt semaphore x clear bit
func (r *registerHsem_c1icrType) GetIsc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_c1icrFieldIscMask) >> RegisterHsem_c1icrFieldIscShift)
}

// SetIsc Interrupt semaphore x clear bit
func (r *registerHsem_c1icrType) SetIsc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_c1icrFieldIscMask)|(uint32(value)<<RegisterHsem_c1icrFieldIscShift))
}

// registerHsem_c1isrType HSEM Interrupt status register
type registerHsem_c1isrType uint32

const (
	RegisterHsem_c1isrFieldIsfShift = 0
	RegisterHsem_c1isrFieldIsfMask  = 0xffffffff
)

// GetIsf Interrupt semaphore x status bit before enable (mask)
func (r *registerHsem_c1isrType) GetIsf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_c1isrFieldIsfMask) >> RegisterHsem_c1isrFieldIsfShift)
}

// SetIsf Interrupt semaphore x status bit before enable (mask)
func (r *registerHsem_c1isrType) SetIsf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_c1isrFieldIsfMask)|(uint32(value)<<RegisterHsem_c1isrFieldIsfShift))
}

// registerHsem_c1misrType HSEM Masked interrupt status register
type registerHsem_c1misrType uint32

const (
	RegisterHsem_c1misrFieldMisfShift = 0
	RegisterHsem_c1misrFieldMisfMask  = 0xffffffff
)

// GetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *registerHsem_c1misrType) GetMisf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_c1misrFieldMisfMask) >> RegisterHsem_c1misrFieldMisfShift)
}

// SetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *registerHsem_c1misrType) SetMisf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_c1misrFieldMisfMask)|(uint32(value)<<RegisterHsem_c1misrFieldMisfShift))
}

// registerHsem_c2ierType HSEM Interrupt enable register
type registerHsem_c2ierType uint32

const (
	RegisterHsem_c2ierFieldIseShift = 0
	RegisterHsem_c2ierFieldIseMask  = 0xffffffff
)

// GetIse Interrupt semaphore x enable bit
func (r *registerHsem_c2ierType) GetIse() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_c2ierFieldIseMask) >> RegisterHsem_c2ierFieldIseShift)
}

// SetIse Interrupt semaphore x enable bit
func (r *registerHsem_c2ierType) SetIse(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_c2ierFieldIseMask)|(uint32(value)<<RegisterHsem_c2ierFieldIseShift))
}

// registerHsem_c2icrType HSEM Interrupt clear register
type registerHsem_c2icrType uint32

const (
	RegisterHsem_c2icrFieldIscShift = 0
	RegisterHsem_c2icrFieldIscMask  = 0xffffffff
)

// GetIsc Interrupt semaphore x clear bit
func (r *registerHsem_c2icrType) GetIsc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_c2icrFieldIscMask) >> RegisterHsem_c2icrFieldIscShift)
}

// SetIsc Interrupt semaphore x clear bit
func (r *registerHsem_c2icrType) SetIsc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_c2icrFieldIscMask)|(uint32(value)<<RegisterHsem_c2icrFieldIscShift))
}

// registerHsem_c2isrType HSEM Interrupt status register
type registerHsem_c2isrType uint32

const (
	RegisterHsem_c2isrFieldIsfShift = 0
	RegisterHsem_c2isrFieldIsfMask  = 0xffffffff
)

// GetIsf Interrupt semaphore x status bit before enable (mask)
func (r *registerHsem_c2isrType) GetIsf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_c2isrFieldIsfMask) >> RegisterHsem_c2isrFieldIsfShift)
}

// SetIsf Interrupt semaphore x status bit before enable (mask)
func (r *registerHsem_c2isrType) SetIsf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_c2isrFieldIsfMask)|(uint32(value)<<RegisterHsem_c2isrFieldIsfShift))
}

// registerHsem_c2misrType HSEM Masked interrupt status register
type registerHsem_c2misrType uint32

const (
	RegisterHsem_c2misrFieldMisfShift = 0
	RegisterHsem_c2misrFieldMisfMask  = 0xffffffff
)

// GetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *registerHsem_c2misrType) GetMisf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_c2misrFieldMisfMask) >> RegisterHsem_c2misrFieldMisfShift)
}

// SetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *registerHsem_c2misrType) SetMisf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_c2misrFieldMisfMask)|(uint32(value)<<RegisterHsem_c2misrFieldMisfShift))
}

// registerHsem_crType HSEM Clear register
type registerHsem_crType uint32

const (
	RegisterHsem_crFieldCoreidShift = 8
	RegisterHsem_crFieldCoreidMask  = 0xf00
)

// GetCoreid COREID of semaphores to be cleared
func (r *registerHsem_crType) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_crFieldCoreidMask) >> RegisterHsem_crFieldCoreidShift)
}

// SetCoreid COREID of semaphores to be cleared
func (r *registerHsem_crType) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_crFieldCoreidMask)|(uint32(value)<<RegisterHsem_crFieldCoreidShift))
}

const (
	RegisterHsem_crFieldKeyShift = 16
	RegisterHsem_crFieldKeyMask  = 0xffff0000
)

// GetKey Semaphore clear Key
func (r *registerHsem_crType) GetKey() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_crFieldKeyMask) >> RegisterHsem_crFieldKeyShift)
}

// SetKey Semaphore clear Key
func (r *registerHsem_crType) SetKey(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_crFieldKeyMask)|(uint32(value)<<RegisterHsem_crFieldKeyShift))
}

// registerHsem_keyrType HSEM Interrupt clear register
type registerHsem_keyrType uint32

const (
	RegisterHsem_keyrFieldKeyShift = 16
	RegisterHsem_keyrFieldKeyMask  = 0xffff0000
)

// GetKey Semaphore Clear Key
func (r *registerHsem_keyrType) GetKey() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_keyrFieldKeyMask) >> RegisterHsem_keyrFieldKeyShift)
}

// SetKey Semaphore Clear Key
func (r *registerHsem_keyrType) SetKey(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_keyrFieldKeyMask)|(uint32(value)<<RegisterHsem_keyrFieldKeyShift))
}
