package hsem

import (
	"unsafe"
	"volatile"
)

var (
	Hsem = (*_hsem)(unsafe.Pointer(uintptr(0x58026400)))
)

type _hsem struct {
	Hsem_r0    registerHsem_r0Type
	Hsem_r1    registerHsem_r1Type
	Hsem_r2    registerHsem_r2Type
	Hsem_r3    registerHsem_r3Type
	Hsem_r4    registerHsem_r4Type
	Hsem_r5    registerHsem_r5Type
	Hsem_r6    registerHsem_r6Type
	Hsem_r7    registerHsem_r7Type
	Hsem_r8    registerHsem_r8Type
	Hsem_r9    registerHsem_r9Type
	Hsem_r10   registerHsem_r10Type
	Hsem_r11   registerHsem_r11Type
	Hsem_r12   registerHsem_r12Type
	Hsem_r13   registerHsem_r13Type
	Hsem_r14   registerHsem_r14Type
	Hsem_r15   registerHsem_r15Type
	Hsem_r16   registerHsem_r16Type
	Hsem_r17   registerHsem_r17Type
	Hsem_r18   registerHsem_r18Type
	Hsem_r19   registerHsem_r19Type
	Hsem_r20   registerHsem_r20Type
	Hsem_r21   registerHsem_r21Type
	Hsem_r22   registerHsem_r22Type
	Hsem_r23   registerHsem_r23Type
	Hsem_r24   registerHsem_r24Type
	Hsem_r25   registerHsem_r25Type
	Hsem_r26   registerHsem_r26Type
	Hsem_r27   registerHsem_r27Type
	Hsem_r28   registerHsem_r28Type
	Hsem_r29   registerHsem_r29Type
	Hsem_r30   registerHsem_r30Type
	Hsem_r31   registerHsem_r31Type
	Hsem_rlr0  registerHsem_rlr0Type
	Hsem_rlr1  registerHsem_rlr1Type
	Hsem_rlr2  registerHsem_rlr2Type
	Hsem_rlr3  registerHsem_rlr3Type
	Hsem_rlr4  registerHsem_rlr4Type
	Hsem_rlr5  registerHsem_rlr5Type
	Hsem_rlr6  registerHsem_rlr6Type
	Hsem_rlr7  registerHsem_rlr7Type
	Hsem_rlr8  registerHsem_rlr8Type
	Hsem_rlr9  registerHsem_rlr9Type
	Hsem_rlr10 registerHsem_rlr10Type
	Hsem_rlr11 registerHsem_rlr11Type
	Hsem_rlr12 registerHsem_rlr12Type
	Hsem_rlr13 registerHsem_rlr13Type
	Hsem_rlr14 registerHsem_rlr14Type
	Hsem_rlr15 registerHsem_rlr15Type
	Hsem_rlr16 registerHsem_rlr16Type
	Hsem_rlr17 registerHsem_rlr17Type
	Hsem_rlr18 registerHsem_rlr18Type
	Hsem_rlr19 registerHsem_rlr19Type
	Hsem_rlr20 registerHsem_rlr20Type
	Hsem_rlr21 registerHsem_rlr21Type
	Hsem_rlr22 registerHsem_rlr22Type
	Hsem_rlr23 registerHsem_rlr23Type
	Hsem_rlr24 registerHsem_rlr24Type
	Hsem_rlr25 registerHsem_rlr25Type
	Hsem_rlr26 registerHsem_rlr26Type
	Hsem_rlr27 registerHsem_rlr27Type
	Hsem_rlr28 registerHsem_rlr28Type
	Hsem_rlr29 registerHsem_rlr29Type
	Hsem_rlr30 registerHsem_rlr30Type
	Hsem_rlr31 registerHsem_rlr31Type
	Hsem_ier   registerHsem_ierType
	Hsem_icr   registerHsem_icrType
	Hsem_isr   registerHsem_isrType
	Hsem_misr  registerHsem_misrType
	_          [48]byte
	Hsem_cr    registerHsem_crType
	Hsem_keyr  registerHsem_keyrType
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
	RegisterHsem_r0FieldMasteridShift = 8
	RegisterHsem_r0FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r0Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r0FieldMasteridMask) >> RegisterHsem_r0FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r0Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r0FieldMasteridMask)|(uint32(value)<<RegisterHsem_r0FieldMasteridShift))
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
	RegisterHsem_r1FieldMasteridShift = 8
	RegisterHsem_r1FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r1Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r1FieldMasteridMask) >> RegisterHsem_r1FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r1Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r1FieldMasteridMask)|(uint32(value)<<RegisterHsem_r1FieldMasteridShift))
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
	RegisterHsem_r2FieldMasteridShift = 8
	RegisterHsem_r2FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r2Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r2FieldMasteridMask) >> RegisterHsem_r2FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r2Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r2FieldMasteridMask)|(uint32(value)<<RegisterHsem_r2FieldMasteridShift))
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
	RegisterHsem_r3FieldMasteridShift = 8
	RegisterHsem_r3FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r3Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r3FieldMasteridMask) >> RegisterHsem_r3FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r3Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r3FieldMasteridMask)|(uint32(value)<<RegisterHsem_r3FieldMasteridShift))
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
	RegisterHsem_r4FieldMasteridShift = 8
	RegisterHsem_r4FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r4Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r4FieldMasteridMask) >> RegisterHsem_r4FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r4Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r4FieldMasteridMask)|(uint32(value)<<RegisterHsem_r4FieldMasteridShift))
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
	RegisterHsem_r5FieldMasteridShift = 8
	RegisterHsem_r5FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r5Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r5FieldMasteridMask) >> RegisterHsem_r5FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r5Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r5FieldMasteridMask)|(uint32(value)<<RegisterHsem_r5FieldMasteridShift))
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
	RegisterHsem_r6FieldMasteridShift = 8
	RegisterHsem_r6FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r6Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r6FieldMasteridMask) >> RegisterHsem_r6FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r6Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r6FieldMasteridMask)|(uint32(value)<<RegisterHsem_r6FieldMasteridShift))
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
	RegisterHsem_r7FieldMasteridShift = 8
	RegisterHsem_r7FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r7Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r7FieldMasteridMask) >> RegisterHsem_r7FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r7Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r7FieldMasteridMask)|(uint32(value)<<RegisterHsem_r7FieldMasteridShift))
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
	RegisterHsem_r8FieldMasteridShift = 8
	RegisterHsem_r8FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r8Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r8FieldMasteridMask) >> RegisterHsem_r8FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r8Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r8FieldMasteridMask)|(uint32(value)<<RegisterHsem_r8FieldMasteridShift))
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
	RegisterHsem_r9FieldMasteridShift = 8
	RegisterHsem_r9FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r9Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r9FieldMasteridMask) >> RegisterHsem_r9FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r9Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r9FieldMasteridMask)|(uint32(value)<<RegisterHsem_r9FieldMasteridShift))
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
	RegisterHsem_r10FieldMasteridShift = 8
	RegisterHsem_r10FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r10Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r10FieldMasteridMask) >> RegisterHsem_r10FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r10Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r10FieldMasteridMask)|(uint32(value)<<RegisterHsem_r10FieldMasteridShift))
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
	RegisterHsem_r11FieldMasteridShift = 8
	RegisterHsem_r11FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r11Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r11FieldMasteridMask) >> RegisterHsem_r11FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r11Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r11FieldMasteridMask)|(uint32(value)<<RegisterHsem_r11FieldMasteridShift))
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
	RegisterHsem_r12FieldMasteridShift = 8
	RegisterHsem_r12FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r12Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r12FieldMasteridMask) >> RegisterHsem_r12FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r12Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r12FieldMasteridMask)|(uint32(value)<<RegisterHsem_r12FieldMasteridShift))
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
	RegisterHsem_r13FieldMasteridShift = 8
	RegisterHsem_r13FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r13Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r13FieldMasteridMask) >> RegisterHsem_r13FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r13Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r13FieldMasteridMask)|(uint32(value)<<RegisterHsem_r13FieldMasteridShift))
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
	RegisterHsem_r14FieldMasteridShift = 8
	RegisterHsem_r14FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r14Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r14FieldMasteridMask) >> RegisterHsem_r14FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r14Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r14FieldMasteridMask)|(uint32(value)<<RegisterHsem_r14FieldMasteridShift))
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
	RegisterHsem_r15FieldMasteridShift = 8
	RegisterHsem_r15FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r15Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r15FieldMasteridMask) >> RegisterHsem_r15FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r15Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r15FieldMasteridMask)|(uint32(value)<<RegisterHsem_r15FieldMasteridShift))
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
	RegisterHsem_r16FieldMasteridShift = 8
	RegisterHsem_r16FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r16Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r16FieldMasteridMask) >> RegisterHsem_r16FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r16Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r16FieldMasteridMask)|(uint32(value)<<RegisterHsem_r16FieldMasteridShift))
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
	RegisterHsem_r17FieldMasteridShift = 8
	RegisterHsem_r17FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r17Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r17FieldMasteridMask) >> RegisterHsem_r17FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r17Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r17FieldMasteridMask)|(uint32(value)<<RegisterHsem_r17FieldMasteridShift))
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
	RegisterHsem_r18FieldMasteridShift = 8
	RegisterHsem_r18FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r18Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r18FieldMasteridMask) >> RegisterHsem_r18FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r18Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r18FieldMasteridMask)|(uint32(value)<<RegisterHsem_r18FieldMasteridShift))
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
	RegisterHsem_r19FieldMasteridShift = 8
	RegisterHsem_r19FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r19Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r19FieldMasteridMask) >> RegisterHsem_r19FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r19Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r19FieldMasteridMask)|(uint32(value)<<RegisterHsem_r19FieldMasteridShift))
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
	RegisterHsem_r20FieldMasteridShift = 8
	RegisterHsem_r20FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r20Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r20FieldMasteridMask) >> RegisterHsem_r20FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r20Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r20FieldMasteridMask)|(uint32(value)<<RegisterHsem_r20FieldMasteridShift))
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
	RegisterHsem_r21FieldMasteridShift = 8
	RegisterHsem_r21FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r21Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r21FieldMasteridMask) >> RegisterHsem_r21FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r21Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r21FieldMasteridMask)|(uint32(value)<<RegisterHsem_r21FieldMasteridShift))
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
	RegisterHsem_r22FieldMasteridShift = 8
	RegisterHsem_r22FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r22Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r22FieldMasteridMask) >> RegisterHsem_r22FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r22Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r22FieldMasteridMask)|(uint32(value)<<RegisterHsem_r22FieldMasteridShift))
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
	RegisterHsem_r23FieldMasteridShift = 8
	RegisterHsem_r23FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r23Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r23FieldMasteridMask) >> RegisterHsem_r23FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r23Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r23FieldMasteridMask)|(uint32(value)<<RegisterHsem_r23FieldMasteridShift))
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
	RegisterHsem_r24FieldMasteridShift = 8
	RegisterHsem_r24FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r24Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r24FieldMasteridMask) >> RegisterHsem_r24FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r24Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r24FieldMasteridMask)|(uint32(value)<<RegisterHsem_r24FieldMasteridShift))
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
	RegisterHsem_r25FieldMasteridShift = 8
	RegisterHsem_r25FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r25Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r25FieldMasteridMask) >> RegisterHsem_r25FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r25Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r25FieldMasteridMask)|(uint32(value)<<RegisterHsem_r25FieldMasteridShift))
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
	RegisterHsem_r26FieldMasteridShift = 8
	RegisterHsem_r26FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r26Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r26FieldMasteridMask) >> RegisterHsem_r26FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r26Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r26FieldMasteridMask)|(uint32(value)<<RegisterHsem_r26FieldMasteridShift))
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
	RegisterHsem_r27FieldMasteridShift = 8
	RegisterHsem_r27FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r27Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r27FieldMasteridMask) >> RegisterHsem_r27FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r27Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r27FieldMasteridMask)|(uint32(value)<<RegisterHsem_r27FieldMasteridShift))
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
	RegisterHsem_r28FieldMasteridShift = 8
	RegisterHsem_r28FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r28Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r28FieldMasteridMask) >> RegisterHsem_r28FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r28Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r28FieldMasteridMask)|(uint32(value)<<RegisterHsem_r28FieldMasteridShift))
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
	RegisterHsem_r29FieldMasteridShift = 8
	RegisterHsem_r29FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r29Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r29FieldMasteridMask) >> RegisterHsem_r29FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r29Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r29FieldMasteridMask)|(uint32(value)<<RegisterHsem_r29FieldMasteridShift))
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
	RegisterHsem_r30FieldMasteridShift = 8
	RegisterHsem_r30FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r30Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r30FieldMasteridMask) >> RegisterHsem_r30FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r30Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r30FieldMasteridMask)|(uint32(value)<<RegisterHsem_r30FieldMasteridShift))
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
	RegisterHsem_r31FieldMasteridShift = 8
	RegisterHsem_r31FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_r31Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_r31FieldMasteridMask) >> RegisterHsem_r31FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_r31Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_r31FieldMasteridMask)|(uint32(value)<<RegisterHsem_r31FieldMasteridShift))
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
	RegisterHsem_rlr0FieldMasteridShift = 8
	RegisterHsem_rlr0FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr0Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr0FieldMasteridMask) >> RegisterHsem_rlr0FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr0Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr0FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr0FieldMasteridShift))
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
	RegisterHsem_rlr1FieldMasteridShift = 8
	RegisterHsem_rlr1FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr1Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr1FieldMasteridMask) >> RegisterHsem_rlr1FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr1Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr1FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr1FieldMasteridShift))
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
	RegisterHsem_rlr2FieldMasteridShift = 8
	RegisterHsem_rlr2FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr2Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr2FieldMasteridMask) >> RegisterHsem_rlr2FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr2Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr2FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr2FieldMasteridShift))
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
	RegisterHsem_rlr3FieldMasteridShift = 8
	RegisterHsem_rlr3FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr3Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr3FieldMasteridMask) >> RegisterHsem_rlr3FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr3Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr3FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr3FieldMasteridShift))
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
	RegisterHsem_rlr4FieldMasteridShift = 8
	RegisterHsem_rlr4FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr4Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr4FieldMasteridMask) >> RegisterHsem_rlr4FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr4Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr4FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr4FieldMasteridShift))
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
	RegisterHsem_rlr5FieldMasteridShift = 8
	RegisterHsem_rlr5FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr5Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr5FieldMasteridMask) >> RegisterHsem_rlr5FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr5Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr5FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr5FieldMasteridShift))
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
	RegisterHsem_rlr6FieldMasteridShift = 8
	RegisterHsem_rlr6FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr6Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr6FieldMasteridMask) >> RegisterHsem_rlr6FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr6Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr6FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr6FieldMasteridShift))
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
	RegisterHsem_rlr7FieldMasteridShift = 8
	RegisterHsem_rlr7FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr7Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr7FieldMasteridMask) >> RegisterHsem_rlr7FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr7Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr7FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr7FieldMasteridShift))
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
	RegisterHsem_rlr8FieldMasteridShift = 8
	RegisterHsem_rlr8FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr8Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr8FieldMasteridMask) >> RegisterHsem_rlr8FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr8Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr8FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr8FieldMasteridShift))
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
	RegisterHsem_rlr9FieldMasteridShift = 8
	RegisterHsem_rlr9FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr9Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr9FieldMasteridMask) >> RegisterHsem_rlr9FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr9Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr9FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr9FieldMasteridShift))
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
	RegisterHsem_rlr10FieldMasteridShift = 8
	RegisterHsem_rlr10FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr10Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr10FieldMasteridMask) >> RegisterHsem_rlr10FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr10Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr10FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr10FieldMasteridShift))
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
	RegisterHsem_rlr11FieldMasteridShift = 8
	RegisterHsem_rlr11FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr11Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr11FieldMasteridMask) >> RegisterHsem_rlr11FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr11Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr11FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr11FieldMasteridShift))
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
	RegisterHsem_rlr12FieldMasteridShift = 8
	RegisterHsem_rlr12FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr12Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr12FieldMasteridMask) >> RegisterHsem_rlr12FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr12Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr12FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr12FieldMasteridShift))
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
	RegisterHsem_rlr13FieldMasteridShift = 8
	RegisterHsem_rlr13FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr13Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr13FieldMasteridMask) >> RegisterHsem_rlr13FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr13Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr13FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr13FieldMasteridShift))
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
	RegisterHsem_rlr14FieldMasteridShift = 8
	RegisterHsem_rlr14FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr14Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr14FieldMasteridMask) >> RegisterHsem_rlr14FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr14Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr14FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr14FieldMasteridShift))
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
	RegisterHsem_rlr15FieldMasteridShift = 8
	RegisterHsem_rlr15FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr15Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr15FieldMasteridMask) >> RegisterHsem_rlr15FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr15Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr15FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr15FieldMasteridShift))
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
	RegisterHsem_rlr16FieldMasteridShift = 8
	RegisterHsem_rlr16FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr16Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr16FieldMasteridMask) >> RegisterHsem_rlr16FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr16Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr16FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr16FieldMasteridShift))
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
	RegisterHsem_rlr17FieldMasteridShift = 8
	RegisterHsem_rlr17FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr17Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr17FieldMasteridMask) >> RegisterHsem_rlr17FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr17Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr17FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr17FieldMasteridShift))
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
	RegisterHsem_rlr18FieldMasteridShift = 8
	RegisterHsem_rlr18FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr18Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr18FieldMasteridMask) >> RegisterHsem_rlr18FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr18Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr18FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr18FieldMasteridShift))
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
	RegisterHsem_rlr19FieldMasteridShift = 8
	RegisterHsem_rlr19FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr19Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr19FieldMasteridMask) >> RegisterHsem_rlr19FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr19Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr19FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr19FieldMasteridShift))
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
	RegisterHsem_rlr20FieldMasteridShift = 8
	RegisterHsem_rlr20FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr20Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr20FieldMasteridMask) >> RegisterHsem_rlr20FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr20Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr20FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr20FieldMasteridShift))
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
	RegisterHsem_rlr21FieldMasteridShift = 8
	RegisterHsem_rlr21FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr21Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr21FieldMasteridMask) >> RegisterHsem_rlr21FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr21Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr21FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr21FieldMasteridShift))
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
	RegisterHsem_rlr22FieldMasteridShift = 8
	RegisterHsem_rlr22FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr22Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr22FieldMasteridMask) >> RegisterHsem_rlr22FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr22Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr22FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr22FieldMasteridShift))
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
	RegisterHsem_rlr23FieldMasteridShift = 8
	RegisterHsem_rlr23FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr23Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr23FieldMasteridMask) >> RegisterHsem_rlr23FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr23Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr23FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr23FieldMasteridShift))
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
	RegisterHsem_rlr24FieldMasteridShift = 8
	RegisterHsem_rlr24FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr24Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr24FieldMasteridMask) >> RegisterHsem_rlr24FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr24Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr24FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr24FieldMasteridShift))
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
	RegisterHsem_rlr25FieldMasteridShift = 8
	RegisterHsem_rlr25FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr25Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr25FieldMasteridMask) >> RegisterHsem_rlr25FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr25Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr25FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr25FieldMasteridShift))
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
	RegisterHsem_rlr26FieldMasteridShift = 8
	RegisterHsem_rlr26FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr26Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr26FieldMasteridMask) >> RegisterHsem_rlr26FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr26Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr26FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr26FieldMasteridShift))
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
	RegisterHsem_rlr27FieldMasteridShift = 8
	RegisterHsem_rlr27FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr27Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr27FieldMasteridMask) >> RegisterHsem_rlr27FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr27Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr27FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr27FieldMasteridShift))
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
	RegisterHsem_rlr28FieldMasteridShift = 8
	RegisterHsem_rlr28FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr28Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr28FieldMasteridMask) >> RegisterHsem_rlr28FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr28Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr28FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr28FieldMasteridShift))
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
	RegisterHsem_rlr29FieldMasteridShift = 8
	RegisterHsem_rlr29FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr29Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr29FieldMasteridMask) >> RegisterHsem_rlr29FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr29Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr29FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr29FieldMasteridShift))
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
	RegisterHsem_rlr30FieldMasteridShift = 8
	RegisterHsem_rlr30FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr30Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr30FieldMasteridMask) >> RegisterHsem_rlr30FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr30Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr30FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr30FieldMasteridShift))
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
	RegisterHsem_rlr31FieldMasteridShift = 8
	RegisterHsem_rlr31FieldMasteridMask  = 0xff00
)

// GetMasterid Semaphore MasterID
func (r *registerHsem_rlr31Type) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_rlr31FieldMasteridMask) >> RegisterHsem_rlr31FieldMasteridShift)
}

// SetMasterid Semaphore MasterID
func (r *registerHsem_rlr31Type) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_rlr31FieldMasteridMask)|(uint32(value)<<RegisterHsem_rlr31FieldMasteridShift))
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

// registerHsem_ierType HSEM Interrupt enable register
type registerHsem_ierType uint32

const (
	RegisterHsem_ierFieldIsem0Shift = 0
	RegisterHsem_ierFieldIsem0Mask  = 0x1
)

// GetIsem0 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem0Mask) != 0
}

// SetIsem0 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem0Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem1Shift = 1
	RegisterHsem_ierFieldIsem1Mask  = 0x2
)

// GetIsem1 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem1Mask) != 0
}

// SetIsem1 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem1Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem2Shift = 2
	RegisterHsem_ierFieldIsem2Mask  = 0x4
)

// GetIsem2 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem2Mask) != 0
}

// SetIsem2 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem2Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem3Shift = 3
	RegisterHsem_ierFieldIsem3Mask  = 0x8
)

// GetIsem3 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem3Mask) != 0
}

// SetIsem3 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem3Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem4Shift = 4
	RegisterHsem_ierFieldIsem4Mask  = 0x10
)

// GetIsem4 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem4Mask) != 0
}

// SetIsem4 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem4Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem5Shift = 5
	RegisterHsem_ierFieldIsem5Mask  = 0x20
)

// GetIsem5 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem5Mask) != 0
}

// SetIsem5 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem5Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem6Shift = 6
	RegisterHsem_ierFieldIsem6Mask  = 0x40
)

// GetIsem6 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem6Mask) != 0
}

// SetIsem6 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem6Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem7Shift = 7
	RegisterHsem_ierFieldIsem7Mask  = 0x80
)

// GetIsem7 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem7Mask) != 0
}

// SetIsem7 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem7Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem8Shift = 8
	RegisterHsem_ierFieldIsem8Mask  = 0x100
)

// GetIsem8 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem8Mask) != 0
}

// SetIsem8 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem8Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem9Shift = 9
	RegisterHsem_ierFieldIsem9Mask  = 0x200
)

// GetIsem9 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem9Mask) != 0
}

// SetIsem9 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem9Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem10Shift = 10
	RegisterHsem_ierFieldIsem10Mask  = 0x400
)

// GetIsem10 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem10Mask) != 0
}

// SetIsem10 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem10Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem11Shift = 11
	RegisterHsem_ierFieldIsem11Mask  = 0x800
)

// GetIsem11 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem11Mask) != 0
}

// SetIsem11 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem11Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem12Shift = 12
	RegisterHsem_ierFieldIsem12Mask  = 0x1000
)

// GetIsem12 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem12Mask) != 0
}

// SetIsem12 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem12Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem13Shift = 13
	RegisterHsem_ierFieldIsem13Mask  = 0x2000
)

// GetIsem13 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem13Mask) != 0
}

// SetIsem13 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem13Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem14Shift = 14
	RegisterHsem_ierFieldIsem14Mask  = 0x4000
)

// GetIsem14 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem14Mask) != 0
}

// SetIsem14 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem14Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem15Shift = 15
	RegisterHsem_ierFieldIsem15Mask  = 0x8000
)

// GetIsem15 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem15Mask) != 0
}

// SetIsem15 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem15Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem16Shift = 16
	RegisterHsem_ierFieldIsem16Mask  = 0x10000
)

// GetIsem16 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem16Mask) != 0
}

// SetIsem16 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem16Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem17Shift = 17
	RegisterHsem_ierFieldIsem17Mask  = 0x20000
)

// GetIsem17 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem17Mask) != 0
}

// SetIsem17 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem17Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem18Shift = 18
	RegisterHsem_ierFieldIsem18Mask  = 0x40000
)

// GetIsem18 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem18Mask) != 0
}

// SetIsem18 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem18Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem19Shift = 19
	RegisterHsem_ierFieldIsem19Mask  = 0x80000
)

// GetIsem19 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem19Mask) != 0
}

// SetIsem19 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem19Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem20Shift = 20
	RegisterHsem_ierFieldIsem20Mask  = 0x100000
)

// GetIsem20 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem20Mask) != 0
}

// SetIsem20 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem20Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem21Shift = 21
	RegisterHsem_ierFieldIsem21Mask  = 0x200000
)

// GetIsem21 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem21Mask) != 0
}

// SetIsem21 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem21Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem22Shift = 22
	RegisterHsem_ierFieldIsem22Mask  = 0x400000
)

// GetIsem22 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem22Mask) != 0
}

// SetIsem22 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem22Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem23Shift = 23
	RegisterHsem_ierFieldIsem23Mask  = 0x800000
)

// GetIsem23 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem23Mask) != 0
}

// SetIsem23 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem23Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem24Shift = 24
	RegisterHsem_ierFieldIsem24Mask  = 0x1000000
)

// GetIsem24 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem24Mask) != 0
}

// SetIsem24 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem24Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem25Shift = 25
	RegisterHsem_ierFieldIsem25Mask  = 0x2000000
)

// GetIsem25 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem25Mask) != 0
}

// SetIsem25 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem25Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem26Shift = 26
	RegisterHsem_ierFieldIsem26Mask  = 0x4000000
)

// GetIsem26 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem26Mask) != 0
}

// SetIsem26 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem26Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem27Shift = 27
	RegisterHsem_ierFieldIsem27Mask  = 0x8000000
)

// GetIsem27 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem27Mask) != 0
}

// SetIsem27 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem27Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem28Shift = 28
	RegisterHsem_ierFieldIsem28Mask  = 0x10000000
)

// GetIsem28 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem28Mask) != 0
}

// SetIsem28 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem28Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem29Shift = 29
	RegisterHsem_ierFieldIsem29Mask  = 0x20000000
)

// GetIsem29 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem29Mask) != 0
}

// SetIsem29 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem29Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem30Shift = 30
	RegisterHsem_ierFieldIsem30Mask  = 0x40000000
)

// GetIsem30 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) GetIsem30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem30Mask) != 0
}

// SetIsem30 Interrupt semaphore n enable bit
func (r *registerHsem_ierType) SetIsem30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem30Mask)
	}
}

const (
	RegisterHsem_ierFieldIsem31Shift = 31
	RegisterHsem_ierFieldIsem31Mask  = 0x80000000
)

// GetIsem31 Interrupt(N) semaphore n enable bit.
func (r *registerHsem_ierType) GetIsem31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_ierFieldIsem31Mask) != 0
}

// SetIsem31 Interrupt(N) semaphore n enable bit.
func (r *registerHsem_ierType) SetIsem31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_ierFieldIsem31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_ierFieldIsem31Mask)
	}
}

// registerHsem_icrType HSEM Interrupt clear register
type registerHsem_icrType uint32

const (
	RegisterHsem_icrFieldIsem0Shift = 0
	RegisterHsem_icrFieldIsem0Mask  = 0x1
)

// GetIsem0 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem0Mask) != 0
}

// SetIsem0 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem0Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem1Shift = 1
	RegisterHsem_icrFieldIsem1Mask  = 0x2
)

// GetIsem1 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem1Mask) != 0
}

// SetIsem1 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem1Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem2Shift = 2
	RegisterHsem_icrFieldIsem2Mask  = 0x4
)

// GetIsem2 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem2Mask) != 0
}

// SetIsem2 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem2Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem3Shift = 3
	RegisterHsem_icrFieldIsem3Mask  = 0x8
)

// GetIsem3 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem3Mask) != 0
}

// SetIsem3 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem3Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem4Shift = 4
	RegisterHsem_icrFieldIsem4Mask  = 0x10
)

// GetIsem4 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem4Mask) != 0
}

// SetIsem4 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem4Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem5Shift = 5
	RegisterHsem_icrFieldIsem5Mask  = 0x20
)

// GetIsem5 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem5Mask) != 0
}

// SetIsem5 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem5Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem6Shift = 6
	RegisterHsem_icrFieldIsem6Mask  = 0x40
)

// GetIsem6 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem6Mask) != 0
}

// SetIsem6 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem6Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem7Shift = 7
	RegisterHsem_icrFieldIsem7Mask  = 0x80
)

// GetIsem7 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem7Mask) != 0
}

// SetIsem7 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem7Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem8Shift = 8
	RegisterHsem_icrFieldIsem8Mask  = 0x100
)

// GetIsem8 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem8Mask) != 0
}

// SetIsem8 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem8Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem9Shift = 9
	RegisterHsem_icrFieldIsem9Mask  = 0x200
)

// GetIsem9 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem9Mask) != 0
}

// SetIsem9 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem9Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem10Shift = 10
	RegisterHsem_icrFieldIsem10Mask  = 0x400
)

// GetIsem10 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem10Mask) != 0
}

// SetIsem10 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem10Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem11Shift = 11
	RegisterHsem_icrFieldIsem11Mask  = 0x800
)

// GetIsem11 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem11Mask) != 0
}

// SetIsem11 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem11Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem12Shift = 12
	RegisterHsem_icrFieldIsem12Mask  = 0x1000
)

// GetIsem12 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem12Mask) != 0
}

// SetIsem12 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem12Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem13Shift = 13
	RegisterHsem_icrFieldIsem13Mask  = 0x2000
)

// GetIsem13 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem13Mask) != 0
}

// SetIsem13 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem13Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem14Shift = 14
	RegisterHsem_icrFieldIsem14Mask  = 0x4000
)

// GetIsem14 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem14Mask) != 0
}

// SetIsem14 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem14Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem15Shift = 15
	RegisterHsem_icrFieldIsem15Mask  = 0x8000
)

// GetIsem15 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem15Mask) != 0
}

// SetIsem15 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem15Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem16Shift = 16
	RegisterHsem_icrFieldIsem16Mask  = 0x10000
)

// GetIsem16 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem16Mask) != 0
}

// SetIsem16 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem16Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem17Shift = 17
	RegisterHsem_icrFieldIsem17Mask  = 0x20000
)

// GetIsem17 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem17Mask) != 0
}

// SetIsem17 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem17Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem18Shift = 18
	RegisterHsem_icrFieldIsem18Mask  = 0x40000
)

// GetIsem18 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem18Mask) != 0
}

// SetIsem18 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem18Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem19Shift = 19
	RegisterHsem_icrFieldIsem19Mask  = 0x80000
)

// GetIsem19 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem19Mask) != 0
}

// SetIsem19 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem19Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem20Shift = 20
	RegisterHsem_icrFieldIsem20Mask  = 0x100000
)

// GetIsem20 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem20Mask) != 0
}

// SetIsem20 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem20Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem21Shift = 21
	RegisterHsem_icrFieldIsem21Mask  = 0x200000
)

// GetIsem21 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem21Mask) != 0
}

// SetIsem21 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem21Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem22Shift = 22
	RegisterHsem_icrFieldIsem22Mask  = 0x400000
)

// GetIsem22 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem22Mask) != 0
}

// SetIsem22 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem22Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem23Shift = 23
	RegisterHsem_icrFieldIsem23Mask  = 0x800000
)

// GetIsem23 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem23Mask) != 0
}

// SetIsem23 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem23Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem24Shift = 24
	RegisterHsem_icrFieldIsem24Mask  = 0x1000000
)

// GetIsem24 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem24Mask) != 0
}

// SetIsem24 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem24Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem25Shift = 25
	RegisterHsem_icrFieldIsem25Mask  = 0x2000000
)

// GetIsem25 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem25Mask) != 0
}

// SetIsem25 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem25Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem26Shift = 26
	RegisterHsem_icrFieldIsem26Mask  = 0x4000000
)

// GetIsem26 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem26Mask) != 0
}

// SetIsem26 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem26Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem27Shift = 27
	RegisterHsem_icrFieldIsem27Mask  = 0x8000000
)

// GetIsem27 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem27Mask) != 0
}

// SetIsem27 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem27Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem28Shift = 28
	RegisterHsem_icrFieldIsem28Mask  = 0x10000000
)

// GetIsem28 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem28Mask) != 0
}

// SetIsem28 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem28Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem29Shift = 29
	RegisterHsem_icrFieldIsem29Mask  = 0x20000000
)

// GetIsem29 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem29Mask) != 0
}

// SetIsem29 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem29Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem30Shift = 30
	RegisterHsem_icrFieldIsem30Mask  = 0x40000000
)

// GetIsem30 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem30Mask) != 0
}

// SetIsem30 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem30Mask)
	}
}

const (
	RegisterHsem_icrFieldIsem31Shift = 31
	RegisterHsem_icrFieldIsem31Mask  = 0x80000000
)

// GetIsem31 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) GetIsem31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_icrFieldIsem31Mask) != 0
}

// SetIsem31 Interrupt(N) semaphore n clear bit
func (r *registerHsem_icrType) SetIsem31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_icrFieldIsem31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_icrFieldIsem31Mask)
	}
}

// registerHsem_isrType HSEM Interrupt status register
type registerHsem_isrType uint32

const (
	RegisterHsem_isrFieldIsem0Shift = 0
	RegisterHsem_isrFieldIsem0Mask  = 0x1
)

// GetIsem0 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem0Mask) != 0
}

// SetIsem0 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem0Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem1Shift = 1
	RegisterHsem_isrFieldIsem1Mask  = 0x2
)

// GetIsem1 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem1Mask) != 0
}

// SetIsem1 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem1Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem2Shift = 2
	RegisterHsem_isrFieldIsem2Mask  = 0x4
)

// GetIsem2 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem2Mask) != 0
}

// SetIsem2 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem2Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem3Shift = 3
	RegisterHsem_isrFieldIsem3Mask  = 0x8
)

// GetIsem3 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem3Mask) != 0
}

// SetIsem3 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem3Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem4Shift = 4
	RegisterHsem_isrFieldIsem4Mask  = 0x10
)

// GetIsem4 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem4Mask) != 0
}

// SetIsem4 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem4Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem5Shift = 5
	RegisterHsem_isrFieldIsem5Mask  = 0x20
)

// GetIsem5 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem5Mask) != 0
}

// SetIsem5 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem5Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem6Shift = 6
	RegisterHsem_isrFieldIsem6Mask  = 0x40
)

// GetIsem6 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem6Mask) != 0
}

// SetIsem6 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem6Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem7Shift = 7
	RegisterHsem_isrFieldIsem7Mask  = 0x80
)

// GetIsem7 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem7Mask) != 0
}

// SetIsem7 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem7Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem8Shift = 8
	RegisterHsem_isrFieldIsem8Mask  = 0x100
)

// GetIsem8 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem8Mask) != 0
}

// SetIsem8 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem8Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem9Shift = 9
	RegisterHsem_isrFieldIsem9Mask  = 0x200
)

// GetIsem9 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem9Mask) != 0
}

// SetIsem9 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem9Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem10Shift = 10
	RegisterHsem_isrFieldIsem10Mask  = 0x400
)

// GetIsem10 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem10Mask) != 0
}

// SetIsem10 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem10Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem11Shift = 11
	RegisterHsem_isrFieldIsem11Mask  = 0x800
)

// GetIsem11 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem11Mask) != 0
}

// SetIsem11 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem11Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem12Shift = 12
	RegisterHsem_isrFieldIsem12Mask  = 0x1000
)

// GetIsem12 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem12Mask) != 0
}

// SetIsem12 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem12Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem13Shift = 13
	RegisterHsem_isrFieldIsem13Mask  = 0x2000
)

// GetIsem13 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem13Mask) != 0
}

// SetIsem13 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem13Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem14Shift = 14
	RegisterHsem_isrFieldIsem14Mask  = 0x4000
)

// GetIsem14 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem14Mask) != 0
}

// SetIsem14 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem14Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem15Shift = 15
	RegisterHsem_isrFieldIsem15Mask  = 0x8000
)

// GetIsem15 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem15Mask) != 0
}

// SetIsem15 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem15Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem16Shift = 16
	RegisterHsem_isrFieldIsem16Mask  = 0x10000
)

// GetIsem16 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem16Mask) != 0
}

// SetIsem16 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem16Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem17Shift = 17
	RegisterHsem_isrFieldIsem17Mask  = 0x20000
)

// GetIsem17 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem17Mask) != 0
}

// SetIsem17 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem17Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem18Shift = 18
	RegisterHsem_isrFieldIsem18Mask  = 0x40000
)

// GetIsem18 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem18Mask) != 0
}

// SetIsem18 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem18Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem19Shift = 19
	RegisterHsem_isrFieldIsem19Mask  = 0x80000
)

// GetIsem19 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem19Mask) != 0
}

// SetIsem19 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem19Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem20Shift = 20
	RegisterHsem_isrFieldIsem20Mask  = 0x100000
)

// GetIsem20 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem20Mask) != 0
}

// SetIsem20 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem20Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem21Shift = 21
	RegisterHsem_isrFieldIsem21Mask  = 0x200000
)

// GetIsem21 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem21Mask) != 0
}

// SetIsem21 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem21Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem22Shift = 22
	RegisterHsem_isrFieldIsem22Mask  = 0x400000
)

// GetIsem22 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem22Mask) != 0
}

// SetIsem22 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem22Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem23Shift = 23
	RegisterHsem_isrFieldIsem23Mask  = 0x800000
)

// GetIsem23 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem23Mask) != 0
}

// SetIsem23 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem23Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem24Shift = 24
	RegisterHsem_isrFieldIsem24Mask  = 0x1000000
)

// GetIsem24 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem24Mask) != 0
}

// SetIsem24 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem24Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem25Shift = 25
	RegisterHsem_isrFieldIsem25Mask  = 0x2000000
)

// GetIsem25 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem25Mask) != 0
}

// SetIsem25 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem25Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem26Shift = 26
	RegisterHsem_isrFieldIsem26Mask  = 0x4000000
)

// GetIsem26 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem26Mask) != 0
}

// SetIsem26 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem26Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem27Shift = 27
	RegisterHsem_isrFieldIsem27Mask  = 0x8000000
)

// GetIsem27 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem27Mask) != 0
}

// SetIsem27 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem27Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem28Shift = 28
	RegisterHsem_isrFieldIsem28Mask  = 0x10000000
)

// GetIsem28 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem28Mask) != 0
}

// SetIsem28 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem28Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem29Shift = 29
	RegisterHsem_isrFieldIsem29Mask  = 0x20000000
)

// GetIsem29 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem29Mask) != 0
}

// SetIsem29 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem29Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem30Shift = 30
	RegisterHsem_isrFieldIsem30Mask  = 0x40000000
)

// GetIsem30 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem30Mask) != 0
}

// SetIsem30 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem30Mask)
	}
}

const (
	RegisterHsem_isrFieldIsem31Shift = 31
	RegisterHsem_isrFieldIsem31Mask  = 0x80000000
)

// GetIsem31 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) GetIsem31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_isrFieldIsem31Mask) != 0
}

// SetIsem31 Interrupt(N) semaphore n status bit before enable (mask)
func (r *registerHsem_isrType) SetIsem31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_isrFieldIsem31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_isrFieldIsem31Mask)
	}
}

// registerHsem_misrType HSEM Masked interrupt status register
type registerHsem_misrType uint32

const (
	RegisterHsem_misrFieldIsem0Shift = 0
	RegisterHsem_misrFieldIsem0Mask  = 0x1
)

// GetIsem0 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem0Mask) != 0
}

// SetIsem0 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem0Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem1Shift = 1
	RegisterHsem_misrFieldIsem1Mask  = 0x2
)

// GetIsem1 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem1Mask) != 0
}

// SetIsem1 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem1Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem2Shift = 2
	RegisterHsem_misrFieldIsem2Mask  = 0x4
)

// GetIsem2 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem2Mask) != 0
}

// SetIsem2 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem2Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem3Shift = 3
	RegisterHsem_misrFieldIsem3Mask  = 0x8
)

// GetIsem3 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem3Mask) != 0
}

// SetIsem3 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem3Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem4Shift = 4
	RegisterHsem_misrFieldIsem4Mask  = 0x10
)

// GetIsem4 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem4Mask) != 0
}

// SetIsem4 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem4Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem5Shift = 5
	RegisterHsem_misrFieldIsem5Mask  = 0x20
)

// GetIsem5 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem5Mask) != 0
}

// SetIsem5 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem5Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem6Shift = 6
	RegisterHsem_misrFieldIsem6Mask  = 0x40
)

// GetIsem6 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem6Mask) != 0
}

// SetIsem6 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem6Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem7Shift = 7
	RegisterHsem_misrFieldIsem7Mask  = 0x80
)

// GetIsem7 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem7Mask) != 0
}

// SetIsem7 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem7Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem8Shift = 8
	RegisterHsem_misrFieldIsem8Mask  = 0x100
)

// GetIsem8 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem8Mask) != 0
}

// SetIsem8 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem8Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem9Shift = 9
	RegisterHsem_misrFieldIsem9Mask  = 0x200
)

// GetIsem9 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem9Mask) != 0
}

// SetIsem9 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem9Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem10Shift = 10
	RegisterHsem_misrFieldIsem10Mask  = 0x400
)

// GetIsem10 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem10Mask) != 0
}

// SetIsem10 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem10Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem11Shift = 11
	RegisterHsem_misrFieldIsem11Mask  = 0x800
)

// GetIsem11 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem11Mask) != 0
}

// SetIsem11 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem11Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem12Shift = 12
	RegisterHsem_misrFieldIsem12Mask  = 0x1000
)

// GetIsem12 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem12Mask) != 0
}

// SetIsem12 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem12Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem13Shift = 13
	RegisterHsem_misrFieldIsem13Mask  = 0x2000
)

// GetIsem13 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem13Mask) != 0
}

// SetIsem13 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem13Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem14Shift = 14
	RegisterHsem_misrFieldIsem14Mask  = 0x4000
)

// GetIsem14 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem14Mask) != 0
}

// SetIsem14 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem14Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem15Shift = 15
	RegisterHsem_misrFieldIsem15Mask  = 0x8000
)

// GetIsem15 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem15Mask) != 0
}

// SetIsem15 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem15Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem16Shift = 16
	RegisterHsem_misrFieldIsem16Mask  = 0x10000
)

// GetIsem16 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem16() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem16Mask) != 0
}

// SetIsem16 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem16(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem16Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem16Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem17Shift = 17
	RegisterHsem_misrFieldIsem17Mask  = 0x20000
)

// GetIsem17 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem17Mask) != 0
}

// SetIsem17 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem17Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem18Shift = 18
	RegisterHsem_misrFieldIsem18Mask  = 0x40000
)

// GetIsem18 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem18() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem18Mask) != 0
}

// SetIsem18 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem18(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem18Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem18Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem19Shift = 19
	RegisterHsem_misrFieldIsem19Mask  = 0x80000
)

// GetIsem19 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem19() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem19Mask) != 0
}

// SetIsem19 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem19(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem19Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem19Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem20Shift = 20
	RegisterHsem_misrFieldIsem20Mask  = 0x100000
)

// GetIsem20 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem20() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem20Mask) != 0
}

// SetIsem20 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem20(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem20Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem20Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem21Shift = 21
	RegisterHsem_misrFieldIsem21Mask  = 0x200000
)

// GetIsem21 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem21() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem21Mask) != 0
}

// SetIsem21 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem21(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem21Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem21Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem22Shift = 22
	RegisterHsem_misrFieldIsem22Mask  = 0x400000
)

// GetIsem22 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem22() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem22Mask) != 0
}

// SetIsem22 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem22(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem22Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem22Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem23Shift = 23
	RegisterHsem_misrFieldIsem23Mask  = 0x800000
)

// GetIsem23 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem23() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem23Mask) != 0
}

// SetIsem23 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem23(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem23Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem23Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem24Shift = 24
	RegisterHsem_misrFieldIsem24Mask  = 0x1000000
)

// GetIsem24 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem24() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem24Mask) != 0
}

// SetIsem24 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem24(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem24Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem24Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem25Shift = 25
	RegisterHsem_misrFieldIsem25Mask  = 0x2000000
)

// GetIsem25 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem25() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem25Mask) != 0
}

// SetIsem25 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem25(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem25Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem25Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem26Shift = 26
	RegisterHsem_misrFieldIsem26Mask  = 0x4000000
)

// GetIsem26 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem26() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem26Mask) != 0
}

// SetIsem26 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem26(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem26Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem26Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem27Shift = 27
	RegisterHsem_misrFieldIsem27Mask  = 0x8000000
)

// GetIsem27 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem27() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem27Mask) != 0
}

// SetIsem27 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem27(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem27Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem27Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem28Shift = 28
	RegisterHsem_misrFieldIsem28Mask  = 0x10000000
)

// GetIsem28 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem28() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem28Mask) != 0
}

// SetIsem28 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem28(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem28Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem28Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem29Shift = 29
	RegisterHsem_misrFieldIsem29Mask  = 0x20000000
)

// GetIsem29 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem29() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem29Mask) != 0
}

// SetIsem29 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem29(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem29Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem29Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem30Shift = 30
	RegisterHsem_misrFieldIsem30Mask  = 0x40000000
)

// GetIsem30 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem30() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem30Mask) != 0
}

// SetIsem30 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem30(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem30Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem30Mask)
	}
}

const (
	RegisterHsem_misrFieldIsem31Shift = 31
	RegisterHsem_misrFieldIsem31Mask  = 0x80000000
)

// GetIsem31 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) GetIsem31() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsem_misrFieldIsem31Mask) != 0
}

// SetIsem31 masked interrupt(N) semaphore n status bit after enable (mask)
func (r *registerHsem_misrType) SetIsem31(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsem_misrFieldIsem31Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsem_misrFieldIsem31Mask)
	}
}

// registerHsem_crType HSEM Clear register
type registerHsem_crType uint32

const (
	RegisterHsem_crFieldMasteridShift = 8
	RegisterHsem_crFieldMasteridMask  = 0xff00
)

// GetMasterid MasterID of semaphores to be cleared
func (r *registerHsem_crType) GetMasterid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsem_crFieldMasteridMask) >> RegisterHsem_crFieldMasteridShift)
}

// SetMasterid MasterID of semaphores to be cleared
func (r *registerHsem_crType) SetMasterid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsem_crFieldMasteridMask)|(uint32(value)<<RegisterHsem_crFieldMasteridShift))
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
