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
	Hsemr0     registerHsemr0Type
	Hsemr1     registerHsemr1Type
	Hsemr2     registerHsemr2Type
	Hsemr3     registerHsemr3Type
	Hsemr4     registerHsemr4Type
	Hsemr5     registerHsemr5Type
	Hsemr6     registerHsemr6Type
	Hsemr7     registerHsemr7Type
	Hsemr8     registerHsemr8Type
	Hsemr9     registerHsemr9Type
	Hsemr10    registerHsemr10Type
	Hsemr11    registerHsemr11Type
	Hsemr12    registerHsemr12Type
	Hsemr13    registerHsemr13Type
	Hsemr14    registerHsemr14Type
	Hsemr15    registerHsemr15Type
	Hsemr16    registerHsemr16Type
	Hsemr17    registerHsemr17Type
	Hsemr18    registerHsemr18Type
	Hsemr19    registerHsemr19Type
	Hsemr20    registerHsemr20Type
	Hsemr21    registerHsemr21Type
	Hsemr22    registerHsemr22Type
	Hsemr23    registerHsemr23Type
	Hsemr24    registerHsemr24Type
	Hsemr25    registerHsemr25Type
	Hsemr26    registerHsemr26Type
	Hsemr27    registerHsemr27Type
	Hsemr28    registerHsemr28Type
	Hsemr29    registerHsemr29Type
	Hsemr30    registerHsemr30Type
	Hsemr31    registerHsemr31Type
	Hsemrlr0   registerHsemrlr0Type
	Hsemrlr1   registerHsemrlr1Type
	Hsemrlr2   registerHsemrlr2Type
	Hsemrlr3   registerHsemrlr3Type
	Hsemrlr4   registerHsemrlr4Type
	Hsemrlr5   registerHsemrlr5Type
	Hsemrlr6   registerHsemrlr6Type
	Hsemrlr7   registerHsemrlr7Type
	Hsemrlr8   registerHsemrlr8Type
	Hsemrlr9   registerHsemrlr9Type
	Hsemrlr10  registerHsemrlr10Type
	Hsemrlr11  registerHsemrlr11Type
	Hsemrlr12  registerHsemrlr12Type
	Hsemrlr13  registerHsemrlr13Type
	Hsemrlr14  registerHsemrlr14Type
	Hsemrlr15  registerHsemrlr15Type
	Hsemrlr16  registerHsemrlr16Type
	Hsemrlr17  registerHsemrlr17Type
	Hsemrlr18  registerHsemrlr18Type
	Hsemrlr19  registerHsemrlr19Type
	Hsemrlr20  registerHsemrlr20Type
	Hsemrlr21  registerHsemrlr21Type
	Hsemrlr22  registerHsemrlr22Type
	Hsemrlr23  registerHsemrlr23Type
	Hsemrlr24  registerHsemrlr24Type
	Hsemrlr25  registerHsemrlr25Type
	Hsemrlr26  registerHsemrlr26Type
	Hsemrlr27  registerHsemrlr27Type
	Hsemrlr28  registerHsemrlr28Type
	Hsemrlr29  registerHsemrlr29Type
	Hsemrlr30  registerHsemrlr30Type
	Hsemrlr31  registerHsemrlr31Type
	Hsemc1ier  registerHsemc1ierType
	Hsemc1icr  registerHsemc1icrType
	Hsemc1isr  registerHsemc1isrType
	Hsemc1misr registerHsemc1misrType
	Hsemc2ier  registerHsemc2ierType
	Hsemc2icr  registerHsemc2icrType
	Hsemc2isr  registerHsemc2isrType
	Hsemc2misr registerHsemc2misrType
	_          [32]byte
	Hsemcr     registerHsemcrType
	Hsemkeyr   registerHsemkeyrType
}

// registerHsemr0Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr0Type uint32

const (
	RegisterHsemr0FieldProcidShift = 0
	RegisterHsemr0FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr0Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr0FieldProcidMask) >> RegisterHsemr0FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr0Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr0FieldProcidMask)|(uint32(value)<<RegisterHsemr0FieldProcidShift))
}

const (
	RegisterHsemr0FieldCoreidShift = 8
	RegisterHsemr0FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr0Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr0FieldCoreidMask) >> RegisterHsemr0FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr0Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr0FieldCoreidMask)|(uint32(value)<<RegisterHsemr0FieldCoreidShift))
}

const (
	RegisterHsemr0FieldLockShift = 31
	RegisterHsemr0FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr0Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr0FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr0Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr0FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr0FieldLockMask)
	}
}

// registerHsemr1Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr1Type uint32

const (
	RegisterHsemr1FieldProcidShift = 0
	RegisterHsemr1FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr1Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr1FieldProcidMask) >> RegisterHsemr1FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr1Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr1FieldProcidMask)|(uint32(value)<<RegisterHsemr1FieldProcidShift))
}

const (
	RegisterHsemr1FieldCoreidShift = 8
	RegisterHsemr1FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr1Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr1FieldCoreidMask) >> RegisterHsemr1FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr1Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr1FieldCoreidMask)|(uint32(value)<<RegisterHsemr1FieldCoreidShift))
}

const (
	RegisterHsemr1FieldLockShift = 31
	RegisterHsemr1FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr1Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr1FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr1Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr1FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr1FieldLockMask)
	}
}

// registerHsemr2Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr2Type uint32

const (
	RegisterHsemr2FieldProcidShift = 0
	RegisterHsemr2FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr2Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr2FieldProcidMask) >> RegisterHsemr2FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr2Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr2FieldProcidMask)|(uint32(value)<<RegisterHsemr2FieldProcidShift))
}

const (
	RegisterHsemr2FieldCoreidShift = 8
	RegisterHsemr2FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr2Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr2FieldCoreidMask) >> RegisterHsemr2FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr2Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr2FieldCoreidMask)|(uint32(value)<<RegisterHsemr2FieldCoreidShift))
}

const (
	RegisterHsemr2FieldLockShift = 31
	RegisterHsemr2FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr2Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr2FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr2Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr2FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr2FieldLockMask)
	}
}

// registerHsemr3Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr3Type uint32

const (
	RegisterHsemr3FieldProcidShift = 0
	RegisterHsemr3FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr3Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr3FieldProcidMask) >> RegisterHsemr3FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr3Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr3FieldProcidMask)|(uint32(value)<<RegisterHsemr3FieldProcidShift))
}

const (
	RegisterHsemr3FieldCoreidShift = 8
	RegisterHsemr3FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr3Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr3FieldCoreidMask) >> RegisterHsemr3FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr3Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr3FieldCoreidMask)|(uint32(value)<<RegisterHsemr3FieldCoreidShift))
}

const (
	RegisterHsemr3FieldLockShift = 31
	RegisterHsemr3FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr3Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr3FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr3Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr3FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr3FieldLockMask)
	}
}

// registerHsemr4Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr4Type uint32

const (
	RegisterHsemr4FieldProcidShift = 0
	RegisterHsemr4FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr4Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr4FieldProcidMask) >> RegisterHsemr4FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr4Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr4FieldProcidMask)|(uint32(value)<<RegisterHsemr4FieldProcidShift))
}

const (
	RegisterHsemr4FieldCoreidShift = 8
	RegisterHsemr4FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr4Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr4FieldCoreidMask) >> RegisterHsemr4FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr4Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr4FieldCoreidMask)|(uint32(value)<<RegisterHsemr4FieldCoreidShift))
}

const (
	RegisterHsemr4FieldLockShift = 31
	RegisterHsemr4FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr4Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr4FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr4Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr4FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr4FieldLockMask)
	}
}

// registerHsemr5Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr5Type uint32

const (
	RegisterHsemr5FieldProcidShift = 0
	RegisterHsemr5FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr5Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr5FieldProcidMask) >> RegisterHsemr5FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr5Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr5FieldProcidMask)|(uint32(value)<<RegisterHsemr5FieldProcidShift))
}

const (
	RegisterHsemr5FieldCoreidShift = 8
	RegisterHsemr5FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr5Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr5FieldCoreidMask) >> RegisterHsemr5FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr5Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr5FieldCoreidMask)|(uint32(value)<<RegisterHsemr5FieldCoreidShift))
}

const (
	RegisterHsemr5FieldLockShift = 31
	RegisterHsemr5FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr5Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr5FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr5Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr5FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr5FieldLockMask)
	}
}

// registerHsemr6Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr6Type uint32

const (
	RegisterHsemr6FieldProcidShift = 0
	RegisterHsemr6FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr6Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr6FieldProcidMask) >> RegisterHsemr6FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr6Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr6FieldProcidMask)|(uint32(value)<<RegisterHsemr6FieldProcidShift))
}

const (
	RegisterHsemr6FieldCoreidShift = 8
	RegisterHsemr6FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr6Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr6FieldCoreidMask) >> RegisterHsemr6FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr6Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr6FieldCoreidMask)|(uint32(value)<<RegisterHsemr6FieldCoreidShift))
}

const (
	RegisterHsemr6FieldLockShift = 31
	RegisterHsemr6FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr6Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr6FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr6Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr6FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr6FieldLockMask)
	}
}

// registerHsemr7Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr7Type uint32

const (
	RegisterHsemr7FieldProcidShift = 0
	RegisterHsemr7FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr7Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr7FieldProcidMask) >> RegisterHsemr7FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr7Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr7FieldProcidMask)|(uint32(value)<<RegisterHsemr7FieldProcidShift))
}

const (
	RegisterHsemr7FieldCoreidShift = 8
	RegisterHsemr7FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr7Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr7FieldCoreidMask) >> RegisterHsemr7FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr7Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr7FieldCoreidMask)|(uint32(value)<<RegisterHsemr7FieldCoreidShift))
}

const (
	RegisterHsemr7FieldLockShift = 31
	RegisterHsemr7FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr7Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr7FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr7Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr7FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr7FieldLockMask)
	}
}

// registerHsemr8Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr8Type uint32

const (
	RegisterHsemr8FieldProcidShift = 0
	RegisterHsemr8FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr8Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr8FieldProcidMask) >> RegisterHsemr8FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr8Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr8FieldProcidMask)|(uint32(value)<<RegisterHsemr8FieldProcidShift))
}

const (
	RegisterHsemr8FieldCoreidShift = 8
	RegisterHsemr8FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr8Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr8FieldCoreidMask) >> RegisterHsemr8FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr8Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr8FieldCoreidMask)|(uint32(value)<<RegisterHsemr8FieldCoreidShift))
}

const (
	RegisterHsemr8FieldLockShift = 31
	RegisterHsemr8FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr8Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr8FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr8Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr8FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr8FieldLockMask)
	}
}

// registerHsemr9Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr9Type uint32

const (
	RegisterHsemr9FieldProcidShift = 0
	RegisterHsemr9FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr9Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr9FieldProcidMask) >> RegisterHsemr9FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr9Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr9FieldProcidMask)|(uint32(value)<<RegisterHsemr9FieldProcidShift))
}

const (
	RegisterHsemr9FieldCoreidShift = 8
	RegisterHsemr9FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr9Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr9FieldCoreidMask) >> RegisterHsemr9FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr9Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr9FieldCoreidMask)|(uint32(value)<<RegisterHsemr9FieldCoreidShift))
}

const (
	RegisterHsemr9FieldLockShift = 31
	RegisterHsemr9FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr9Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr9FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr9Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr9FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr9FieldLockMask)
	}
}

// registerHsemr10Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr10Type uint32

const (
	RegisterHsemr10FieldProcidShift = 0
	RegisterHsemr10FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr10Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr10FieldProcidMask) >> RegisterHsemr10FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr10Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr10FieldProcidMask)|(uint32(value)<<RegisterHsemr10FieldProcidShift))
}

const (
	RegisterHsemr10FieldCoreidShift = 8
	RegisterHsemr10FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr10Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr10FieldCoreidMask) >> RegisterHsemr10FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr10Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr10FieldCoreidMask)|(uint32(value)<<RegisterHsemr10FieldCoreidShift))
}

const (
	RegisterHsemr10FieldLockShift = 31
	RegisterHsemr10FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr10Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr10FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr10Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr10FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr10FieldLockMask)
	}
}

// registerHsemr11Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr11Type uint32

const (
	RegisterHsemr11FieldProcidShift = 0
	RegisterHsemr11FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr11Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr11FieldProcidMask) >> RegisterHsemr11FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr11Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr11FieldProcidMask)|(uint32(value)<<RegisterHsemr11FieldProcidShift))
}

const (
	RegisterHsemr11FieldCoreidShift = 8
	RegisterHsemr11FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr11Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr11FieldCoreidMask) >> RegisterHsemr11FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr11Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr11FieldCoreidMask)|(uint32(value)<<RegisterHsemr11FieldCoreidShift))
}

const (
	RegisterHsemr11FieldLockShift = 31
	RegisterHsemr11FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr11Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr11FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr11Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr11FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr11FieldLockMask)
	}
}

// registerHsemr12Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr12Type uint32

const (
	RegisterHsemr12FieldProcidShift = 0
	RegisterHsemr12FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr12Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr12FieldProcidMask) >> RegisterHsemr12FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr12Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr12FieldProcidMask)|(uint32(value)<<RegisterHsemr12FieldProcidShift))
}

const (
	RegisterHsemr12FieldCoreidShift = 8
	RegisterHsemr12FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr12Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr12FieldCoreidMask) >> RegisterHsemr12FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr12Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr12FieldCoreidMask)|(uint32(value)<<RegisterHsemr12FieldCoreidShift))
}

const (
	RegisterHsemr12FieldLockShift = 31
	RegisterHsemr12FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr12Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr12FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr12Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr12FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr12FieldLockMask)
	}
}

// registerHsemr13Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr13Type uint32

const (
	RegisterHsemr13FieldProcidShift = 0
	RegisterHsemr13FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr13Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr13FieldProcidMask) >> RegisterHsemr13FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr13Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr13FieldProcidMask)|(uint32(value)<<RegisterHsemr13FieldProcidShift))
}

const (
	RegisterHsemr13FieldCoreidShift = 8
	RegisterHsemr13FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr13Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr13FieldCoreidMask) >> RegisterHsemr13FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr13Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr13FieldCoreidMask)|(uint32(value)<<RegisterHsemr13FieldCoreidShift))
}

const (
	RegisterHsemr13FieldLockShift = 31
	RegisterHsemr13FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr13Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr13FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr13Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr13FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr13FieldLockMask)
	}
}

// registerHsemr14Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr14Type uint32

const (
	RegisterHsemr14FieldProcidShift = 0
	RegisterHsemr14FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr14Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr14FieldProcidMask) >> RegisterHsemr14FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr14Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr14FieldProcidMask)|(uint32(value)<<RegisterHsemr14FieldProcidShift))
}

const (
	RegisterHsemr14FieldCoreidShift = 8
	RegisterHsemr14FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr14Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr14FieldCoreidMask) >> RegisterHsemr14FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr14Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr14FieldCoreidMask)|(uint32(value)<<RegisterHsemr14FieldCoreidShift))
}

const (
	RegisterHsemr14FieldLockShift = 31
	RegisterHsemr14FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr14Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr14FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr14Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr14FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr14FieldLockMask)
	}
}

// registerHsemr15Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr15Type uint32

const (
	RegisterHsemr15FieldProcidShift = 0
	RegisterHsemr15FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr15Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr15FieldProcidMask) >> RegisterHsemr15FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr15Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr15FieldProcidMask)|(uint32(value)<<RegisterHsemr15FieldProcidShift))
}

const (
	RegisterHsemr15FieldCoreidShift = 8
	RegisterHsemr15FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr15Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr15FieldCoreidMask) >> RegisterHsemr15FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr15Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr15FieldCoreidMask)|(uint32(value)<<RegisterHsemr15FieldCoreidShift))
}

const (
	RegisterHsemr15FieldLockShift = 31
	RegisterHsemr15FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr15Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr15FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr15Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr15FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr15FieldLockMask)
	}
}

// registerHsemr16Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr16Type uint32

const (
	RegisterHsemr16FieldProcidShift = 0
	RegisterHsemr16FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr16Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr16FieldProcidMask) >> RegisterHsemr16FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr16Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr16FieldProcidMask)|(uint32(value)<<RegisterHsemr16FieldProcidShift))
}

const (
	RegisterHsemr16FieldCoreidShift = 8
	RegisterHsemr16FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr16Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr16FieldCoreidMask) >> RegisterHsemr16FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr16Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr16FieldCoreidMask)|(uint32(value)<<RegisterHsemr16FieldCoreidShift))
}

const (
	RegisterHsemr16FieldLockShift = 31
	RegisterHsemr16FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr16Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr16FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr16Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr16FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr16FieldLockMask)
	}
}

// registerHsemr17Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr17Type uint32

const (
	RegisterHsemr17FieldProcidShift = 0
	RegisterHsemr17FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr17Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr17FieldProcidMask) >> RegisterHsemr17FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr17Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr17FieldProcidMask)|(uint32(value)<<RegisterHsemr17FieldProcidShift))
}

const (
	RegisterHsemr17FieldCoreidShift = 8
	RegisterHsemr17FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr17Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr17FieldCoreidMask) >> RegisterHsemr17FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr17Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr17FieldCoreidMask)|(uint32(value)<<RegisterHsemr17FieldCoreidShift))
}

const (
	RegisterHsemr17FieldLockShift = 31
	RegisterHsemr17FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr17Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr17FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr17Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr17FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr17FieldLockMask)
	}
}

// registerHsemr18Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr18Type uint32

const (
	RegisterHsemr18FieldProcidShift = 0
	RegisterHsemr18FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr18Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr18FieldProcidMask) >> RegisterHsemr18FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr18Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr18FieldProcidMask)|(uint32(value)<<RegisterHsemr18FieldProcidShift))
}

const (
	RegisterHsemr18FieldCoreidShift = 8
	RegisterHsemr18FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr18Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr18FieldCoreidMask) >> RegisterHsemr18FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr18Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr18FieldCoreidMask)|(uint32(value)<<RegisterHsemr18FieldCoreidShift))
}

const (
	RegisterHsemr18FieldLockShift = 31
	RegisterHsemr18FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr18Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr18FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr18Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr18FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr18FieldLockMask)
	}
}

// registerHsemr19Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr19Type uint32

const (
	RegisterHsemr19FieldProcidShift = 0
	RegisterHsemr19FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr19Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr19FieldProcidMask) >> RegisterHsemr19FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr19Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr19FieldProcidMask)|(uint32(value)<<RegisterHsemr19FieldProcidShift))
}

const (
	RegisterHsemr19FieldCoreidShift = 8
	RegisterHsemr19FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr19Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr19FieldCoreidMask) >> RegisterHsemr19FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr19Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr19FieldCoreidMask)|(uint32(value)<<RegisterHsemr19FieldCoreidShift))
}

const (
	RegisterHsemr19FieldLockShift = 31
	RegisterHsemr19FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr19Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr19FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr19Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr19FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr19FieldLockMask)
	}
}

// registerHsemr20Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr20Type uint32

const (
	RegisterHsemr20FieldProcidShift = 0
	RegisterHsemr20FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr20Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr20FieldProcidMask) >> RegisterHsemr20FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr20Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr20FieldProcidMask)|(uint32(value)<<RegisterHsemr20FieldProcidShift))
}

const (
	RegisterHsemr20FieldCoreidShift = 8
	RegisterHsemr20FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr20Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr20FieldCoreidMask) >> RegisterHsemr20FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr20Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr20FieldCoreidMask)|(uint32(value)<<RegisterHsemr20FieldCoreidShift))
}

const (
	RegisterHsemr20FieldLockShift = 31
	RegisterHsemr20FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr20Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr20FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr20Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr20FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr20FieldLockMask)
	}
}

// registerHsemr21Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr21Type uint32

const (
	RegisterHsemr21FieldProcidShift = 0
	RegisterHsemr21FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr21Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr21FieldProcidMask) >> RegisterHsemr21FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr21Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr21FieldProcidMask)|(uint32(value)<<RegisterHsemr21FieldProcidShift))
}

const (
	RegisterHsemr21FieldCoreidShift = 8
	RegisterHsemr21FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr21Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr21FieldCoreidMask) >> RegisterHsemr21FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr21Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr21FieldCoreidMask)|(uint32(value)<<RegisterHsemr21FieldCoreidShift))
}

const (
	RegisterHsemr21FieldLockShift = 31
	RegisterHsemr21FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr21Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr21FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr21Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr21FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr21FieldLockMask)
	}
}

// registerHsemr22Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr22Type uint32

const (
	RegisterHsemr22FieldProcidShift = 0
	RegisterHsemr22FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr22Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr22FieldProcidMask) >> RegisterHsemr22FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr22Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr22FieldProcidMask)|(uint32(value)<<RegisterHsemr22FieldProcidShift))
}

const (
	RegisterHsemr22FieldCoreidShift = 8
	RegisterHsemr22FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr22Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr22FieldCoreidMask) >> RegisterHsemr22FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr22Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr22FieldCoreidMask)|(uint32(value)<<RegisterHsemr22FieldCoreidShift))
}

const (
	RegisterHsemr22FieldLockShift = 31
	RegisterHsemr22FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr22Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr22FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr22Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr22FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr22FieldLockMask)
	}
}

// registerHsemr23Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr23Type uint32

const (
	RegisterHsemr23FieldProcidShift = 0
	RegisterHsemr23FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr23Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr23FieldProcidMask) >> RegisterHsemr23FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr23Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr23FieldProcidMask)|(uint32(value)<<RegisterHsemr23FieldProcidShift))
}

const (
	RegisterHsemr23FieldCoreidShift = 8
	RegisterHsemr23FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr23Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr23FieldCoreidMask) >> RegisterHsemr23FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr23Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr23FieldCoreidMask)|(uint32(value)<<RegisterHsemr23FieldCoreidShift))
}

const (
	RegisterHsemr23FieldLockShift = 31
	RegisterHsemr23FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr23Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr23FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr23Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr23FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr23FieldLockMask)
	}
}

// registerHsemr24Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr24Type uint32

const (
	RegisterHsemr24FieldProcidShift = 0
	RegisterHsemr24FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr24Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr24FieldProcidMask) >> RegisterHsemr24FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr24Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr24FieldProcidMask)|(uint32(value)<<RegisterHsemr24FieldProcidShift))
}

const (
	RegisterHsemr24FieldCoreidShift = 8
	RegisterHsemr24FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr24Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr24FieldCoreidMask) >> RegisterHsemr24FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr24Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr24FieldCoreidMask)|(uint32(value)<<RegisterHsemr24FieldCoreidShift))
}

const (
	RegisterHsemr24FieldLockShift = 31
	RegisterHsemr24FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr24Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr24FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr24Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr24FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr24FieldLockMask)
	}
}

// registerHsemr25Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr25Type uint32

const (
	RegisterHsemr25FieldProcidShift = 0
	RegisterHsemr25FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr25Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr25FieldProcidMask) >> RegisterHsemr25FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr25Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr25FieldProcidMask)|(uint32(value)<<RegisterHsemr25FieldProcidShift))
}

const (
	RegisterHsemr25FieldCoreidShift = 8
	RegisterHsemr25FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr25Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr25FieldCoreidMask) >> RegisterHsemr25FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr25Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr25FieldCoreidMask)|(uint32(value)<<RegisterHsemr25FieldCoreidShift))
}

const (
	RegisterHsemr25FieldLockShift = 31
	RegisterHsemr25FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr25Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr25FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr25Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr25FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr25FieldLockMask)
	}
}

// registerHsemr26Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr26Type uint32

const (
	RegisterHsemr26FieldProcidShift = 0
	RegisterHsemr26FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr26Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr26FieldProcidMask) >> RegisterHsemr26FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr26Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr26FieldProcidMask)|(uint32(value)<<RegisterHsemr26FieldProcidShift))
}

const (
	RegisterHsemr26FieldCoreidShift = 8
	RegisterHsemr26FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr26Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr26FieldCoreidMask) >> RegisterHsemr26FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr26Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr26FieldCoreidMask)|(uint32(value)<<RegisterHsemr26FieldCoreidShift))
}

const (
	RegisterHsemr26FieldLockShift = 31
	RegisterHsemr26FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr26Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr26FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr26Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr26FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr26FieldLockMask)
	}
}

// registerHsemr27Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr27Type uint32

const (
	RegisterHsemr27FieldProcidShift = 0
	RegisterHsemr27FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr27Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr27FieldProcidMask) >> RegisterHsemr27FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr27Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr27FieldProcidMask)|(uint32(value)<<RegisterHsemr27FieldProcidShift))
}

const (
	RegisterHsemr27FieldCoreidShift = 8
	RegisterHsemr27FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr27Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr27FieldCoreidMask) >> RegisterHsemr27FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr27Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr27FieldCoreidMask)|(uint32(value)<<RegisterHsemr27FieldCoreidShift))
}

const (
	RegisterHsemr27FieldLockShift = 31
	RegisterHsemr27FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr27Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr27FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr27Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr27FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr27FieldLockMask)
	}
}

// registerHsemr28Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr28Type uint32

const (
	RegisterHsemr28FieldProcidShift = 0
	RegisterHsemr28FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr28Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr28FieldProcidMask) >> RegisterHsemr28FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr28Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr28FieldProcidMask)|(uint32(value)<<RegisterHsemr28FieldProcidShift))
}

const (
	RegisterHsemr28FieldCoreidShift = 8
	RegisterHsemr28FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr28Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr28FieldCoreidMask) >> RegisterHsemr28FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr28Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr28FieldCoreidMask)|(uint32(value)<<RegisterHsemr28FieldCoreidShift))
}

const (
	RegisterHsemr28FieldLockShift = 31
	RegisterHsemr28FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr28Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr28FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr28Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr28FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr28FieldLockMask)
	}
}

// registerHsemr29Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr29Type uint32

const (
	RegisterHsemr29FieldProcidShift = 0
	RegisterHsemr29FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr29Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr29FieldProcidMask) >> RegisterHsemr29FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr29Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr29FieldProcidMask)|(uint32(value)<<RegisterHsemr29FieldProcidShift))
}

const (
	RegisterHsemr29FieldCoreidShift = 8
	RegisterHsemr29FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr29Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr29FieldCoreidMask) >> RegisterHsemr29FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr29Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr29FieldCoreidMask)|(uint32(value)<<RegisterHsemr29FieldCoreidShift))
}

const (
	RegisterHsemr29FieldLockShift = 31
	RegisterHsemr29FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr29Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr29FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr29Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr29FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr29FieldLockMask)
	}
}

// registerHsemr30Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr30Type uint32

const (
	RegisterHsemr30FieldProcidShift = 0
	RegisterHsemr30FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr30Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr30FieldProcidMask) >> RegisterHsemr30FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr30Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr30FieldProcidMask)|(uint32(value)<<RegisterHsemr30FieldProcidShift))
}

const (
	RegisterHsemr30FieldCoreidShift = 8
	RegisterHsemr30FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr30Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr30FieldCoreidMask) >> RegisterHsemr30FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr30Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr30FieldCoreidMask)|(uint32(value)<<RegisterHsemr30FieldCoreidShift))
}

const (
	RegisterHsemr30FieldLockShift = 31
	RegisterHsemr30FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr30Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr30FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr30Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr30FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr30FieldLockMask)
	}
}

// registerHsemr31Type HSEM register HSEM_R0 HSEM_R31
type registerHsemr31Type uint32

const (
	RegisterHsemr31FieldProcidShift = 0
	RegisterHsemr31FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemr31Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr31FieldProcidMask) >> RegisterHsemr31FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemr31Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr31FieldProcidMask)|(uint32(value)<<RegisterHsemr31FieldProcidShift))
}

const (
	RegisterHsemr31FieldCoreidShift = 8
	RegisterHsemr31FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemr31Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr31FieldCoreidMask) >> RegisterHsemr31FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemr31Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr31FieldCoreidMask)|(uint32(value)<<RegisterHsemr31FieldCoreidShift))
}

const (
	RegisterHsemr31FieldLockShift = 31
	RegisterHsemr31FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemr31Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr31FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemr31Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr31FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr31FieldLockMask)
	}
}

// registerHsemrlr0Type HSEM Read lock register
type registerHsemrlr0Type uint32

const (
	RegisterHsemrlr0FieldProcidShift = 0
	RegisterHsemrlr0FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr0Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr0FieldProcidMask) >> RegisterHsemrlr0FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr0Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr0FieldProcidMask)|(uint32(value)<<RegisterHsemrlr0FieldProcidShift))
}

const (
	RegisterHsemrlr0FieldCoreidShift = 8
	RegisterHsemrlr0FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr0Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr0FieldCoreidMask) >> RegisterHsemrlr0FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr0Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr0FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr0FieldCoreidShift))
}

const (
	RegisterHsemrlr0FieldLockShift = 31
	RegisterHsemrlr0FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr0Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr0FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr0Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr0FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr0FieldLockMask)
	}
}

// registerHsemrlr1Type HSEM Read lock register
type registerHsemrlr1Type uint32

const (
	RegisterHsemrlr1FieldProcidShift = 0
	RegisterHsemrlr1FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr1Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr1FieldProcidMask) >> RegisterHsemrlr1FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr1Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr1FieldProcidMask)|(uint32(value)<<RegisterHsemrlr1FieldProcidShift))
}

const (
	RegisterHsemrlr1FieldCoreidShift = 8
	RegisterHsemrlr1FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr1Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr1FieldCoreidMask) >> RegisterHsemrlr1FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr1Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr1FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr1FieldCoreidShift))
}

const (
	RegisterHsemrlr1FieldLockShift = 31
	RegisterHsemrlr1FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr1Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr1FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr1Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr1FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr1FieldLockMask)
	}
}

// registerHsemrlr2Type HSEM Read lock register
type registerHsemrlr2Type uint32

const (
	RegisterHsemrlr2FieldProcidShift = 0
	RegisterHsemrlr2FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr2Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr2FieldProcidMask) >> RegisterHsemrlr2FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr2Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr2FieldProcidMask)|(uint32(value)<<RegisterHsemrlr2FieldProcidShift))
}

const (
	RegisterHsemrlr2FieldCoreidShift = 8
	RegisterHsemrlr2FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr2Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr2FieldCoreidMask) >> RegisterHsemrlr2FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr2Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr2FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr2FieldCoreidShift))
}

const (
	RegisterHsemrlr2FieldLockShift = 31
	RegisterHsemrlr2FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr2Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr2FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr2Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr2FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr2FieldLockMask)
	}
}

// registerHsemrlr3Type HSEM Read lock register
type registerHsemrlr3Type uint32

const (
	RegisterHsemrlr3FieldProcidShift = 0
	RegisterHsemrlr3FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr3Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr3FieldProcidMask) >> RegisterHsemrlr3FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr3Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr3FieldProcidMask)|(uint32(value)<<RegisterHsemrlr3FieldProcidShift))
}

const (
	RegisterHsemrlr3FieldCoreidShift = 8
	RegisterHsemrlr3FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr3Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr3FieldCoreidMask) >> RegisterHsemrlr3FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr3Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr3FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr3FieldCoreidShift))
}

const (
	RegisterHsemrlr3FieldLockShift = 31
	RegisterHsemrlr3FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr3Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr3FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr3Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr3FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr3FieldLockMask)
	}
}

// registerHsemrlr4Type HSEM Read lock register
type registerHsemrlr4Type uint32

const (
	RegisterHsemrlr4FieldProcidShift = 0
	RegisterHsemrlr4FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr4Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr4FieldProcidMask) >> RegisterHsemrlr4FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr4Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr4FieldProcidMask)|(uint32(value)<<RegisterHsemrlr4FieldProcidShift))
}

const (
	RegisterHsemrlr4FieldCoreidShift = 8
	RegisterHsemrlr4FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr4Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr4FieldCoreidMask) >> RegisterHsemrlr4FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr4Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr4FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr4FieldCoreidShift))
}

const (
	RegisterHsemrlr4FieldLockShift = 31
	RegisterHsemrlr4FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr4Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr4FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr4Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr4FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr4FieldLockMask)
	}
}

// registerHsemrlr5Type HSEM Read lock register
type registerHsemrlr5Type uint32

const (
	RegisterHsemrlr5FieldProcidShift = 0
	RegisterHsemrlr5FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr5Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr5FieldProcidMask) >> RegisterHsemrlr5FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr5Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr5FieldProcidMask)|(uint32(value)<<RegisterHsemrlr5FieldProcidShift))
}

const (
	RegisterHsemrlr5FieldCoreidShift = 8
	RegisterHsemrlr5FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr5Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr5FieldCoreidMask) >> RegisterHsemrlr5FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr5Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr5FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr5FieldCoreidShift))
}

const (
	RegisterHsemrlr5FieldLockShift = 31
	RegisterHsemrlr5FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr5Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr5FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr5Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr5FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr5FieldLockMask)
	}
}

// registerHsemrlr6Type HSEM Read lock register
type registerHsemrlr6Type uint32

const (
	RegisterHsemrlr6FieldProcidShift = 0
	RegisterHsemrlr6FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr6Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr6FieldProcidMask) >> RegisterHsemrlr6FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr6Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr6FieldProcidMask)|(uint32(value)<<RegisterHsemrlr6FieldProcidShift))
}

const (
	RegisterHsemrlr6FieldCoreidShift = 8
	RegisterHsemrlr6FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr6Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr6FieldCoreidMask) >> RegisterHsemrlr6FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr6Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr6FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr6FieldCoreidShift))
}

const (
	RegisterHsemrlr6FieldLockShift = 31
	RegisterHsemrlr6FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr6Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr6FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr6Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr6FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr6FieldLockMask)
	}
}

// registerHsemrlr7Type HSEM Read lock register
type registerHsemrlr7Type uint32

const (
	RegisterHsemrlr7FieldProcidShift = 0
	RegisterHsemrlr7FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr7Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr7FieldProcidMask) >> RegisterHsemrlr7FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr7Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr7FieldProcidMask)|(uint32(value)<<RegisterHsemrlr7FieldProcidShift))
}

const (
	RegisterHsemrlr7FieldCoreidShift = 8
	RegisterHsemrlr7FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr7Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr7FieldCoreidMask) >> RegisterHsemrlr7FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr7Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr7FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr7FieldCoreidShift))
}

const (
	RegisterHsemrlr7FieldLockShift = 31
	RegisterHsemrlr7FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr7Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr7FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr7Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr7FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr7FieldLockMask)
	}
}

// registerHsemrlr8Type HSEM Read lock register
type registerHsemrlr8Type uint32

const (
	RegisterHsemrlr8FieldProcidShift = 0
	RegisterHsemrlr8FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr8Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr8FieldProcidMask) >> RegisterHsemrlr8FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr8Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr8FieldProcidMask)|(uint32(value)<<RegisterHsemrlr8FieldProcidShift))
}

const (
	RegisterHsemrlr8FieldCoreidShift = 8
	RegisterHsemrlr8FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr8Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr8FieldCoreidMask) >> RegisterHsemrlr8FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr8Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr8FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr8FieldCoreidShift))
}

const (
	RegisterHsemrlr8FieldLockShift = 31
	RegisterHsemrlr8FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr8Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr8FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr8Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr8FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr8FieldLockMask)
	}
}

// registerHsemrlr9Type HSEM Read lock register
type registerHsemrlr9Type uint32

const (
	RegisterHsemrlr9FieldProcidShift = 0
	RegisterHsemrlr9FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr9Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr9FieldProcidMask) >> RegisterHsemrlr9FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr9Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr9FieldProcidMask)|(uint32(value)<<RegisterHsemrlr9FieldProcidShift))
}

const (
	RegisterHsemrlr9FieldCoreidShift = 8
	RegisterHsemrlr9FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr9Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr9FieldCoreidMask) >> RegisterHsemrlr9FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr9Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr9FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr9FieldCoreidShift))
}

const (
	RegisterHsemrlr9FieldLockShift = 31
	RegisterHsemrlr9FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr9Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr9FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr9Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr9FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr9FieldLockMask)
	}
}

// registerHsemrlr10Type HSEM Read lock register
type registerHsemrlr10Type uint32

const (
	RegisterHsemrlr10FieldProcidShift = 0
	RegisterHsemrlr10FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr10Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr10FieldProcidMask) >> RegisterHsemrlr10FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr10Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr10FieldProcidMask)|(uint32(value)<<RegisterHsemrlr10FieldProcidShift))
}

const (
	RegisterHsemrlr10FieldCoreidShift = 8
	RegisterHsemrlr10FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr10Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr10FieldCoreidMask) >> RegisterHsemrlr10FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr10Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr10FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr10FieldCoreidShift))
}

const (
	RegisterHsemrlr10FieldLockShift = 31
	RegisterHsemrlr10FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr10Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr10FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr10Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr10FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr10FieldLockMask)
	}
}

// registerHsemrlr11Type HSEM Read lock register
type registerHsemrlr11Type uint32

const (
	RegisterHsemrlr11FieldProcidShift = 0
	RegisterHsemrlr11FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr11Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr11FieldProcidMask) >> RegisterHsemrlr11FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr11Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr11FieldProcidMask)|(uint32(value)<<RegisterHsemrlr11FieldProcidShift))
}

const (
	RegisterHsemrlr11FieldCoreidShift = 8
	RegisterHsemrlr11FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr11Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr11FieldCoreidMask) >> RegisterHsemrlr11FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr11Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr11FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr11FieldCoreidShift))
}

const (
	RegisterHsemrlr11FieldLockShift = 31
	RegisterHsemrlr11FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr11Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr11FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr11Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr11FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr11FieldLockMask)
	}
}

// registerHsemrlr12Type HSEM Read lock register
type registerHsemrlr12Type uint32

const (
	RegisterHsemrlr12FieldProcidShift = 0
	RegisterHsemrlr12FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr12Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr12FieldProcidMask) >> RegisterHsemrlr12FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr12Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr12FieldProcidMask)|(uint32(value)<<RegisterHsemrlr12FieldProcidShift))
}

const (
	RegisterHsemrlr12FieldCoreidShift = 8
	RegisterHsemrlr12FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr12Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr12FieldCoreidMask) >> RegisterHsemrlr12FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr12Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr12FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr12FieldCoreidShift))
}

const (
	RegisterHsemrlr12FieldLockShift = 31
	RegisterHsemrlr12FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr12Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr12FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr12Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr12FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr12FieldLockMask)
	}
}

// registerHsemrlr13Type HSEM Read lock register
type registerHsemrlr13Type uint32

const (
	RegisterHsemrlr13FieldProcidShift = 0
	RegisterHsemrlr13FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr13Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr13FieldProcidMask) >> RegisterHsemrlr13FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr13Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr13FieldProcidMask)|(uint32(value)<<RegisterHsemrlr13FieldProcidShift))
}

const (
	RegisterHsemrlr13FieldCoreidShift = 8
	RegisterHsemrlr13FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr13Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr13FieldCoreidMask) >> RegisterHsemrlr13FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr13Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr13FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr13FieldCoreidShift))
}

const (
	RegisterHsemrlr13FieldLockShift = 31
	RegisterHsemrlr13FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr13Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr13FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr13Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr13FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr13FieldLockMask)
	}
}

// registerHsemrlr14Type HSEM Read lock register
type registerHsemrlr14Type uint32

const (
	RegisterHsemrlr14FieldProcidShift = 0
	RegisterHsemrlr14FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr14Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr14FieldProcidMask) >> RegisterHsemrlr14FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr14Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr14FieldProcidMask)|(uint32(value)<<RegisterHsemrlr14FieldProcidShift))
}

const (
	RegisterHsemrlr14FieldCoreidShift = 8
	RegisterHsemrlr14FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr14Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr14FieldCoreidMask) >> RegisterHsemrlr14FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr14Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr14FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr14FieldCoreidShift))
}

const (
	RegisterHsemrlr14FieldLockShift = 31
	RegisterHsemrlr14FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr14Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr14FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr14Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr14FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr14FieldLockMask)
	}
}

// registerHsemrlr15Type HSEM Read lock register
type registerHsemrlr15Type uint32

const (
	RegisterHsemrlr15FieldProcidShift = 0
	RegisterHsemrlr15FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr15Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr15FieldProcidMask) >> RegisterHsemrlr15FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr15Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr15FieldProcidMask)|(uint32(value)<<RegisterHsemrlr15FieldProcidShift))
}

const (
	RegisterHsemrlr15FieldCoreidShift = 8
	RegisterHsemrlr15FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr15Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr15FieldCoreidMask) >> RegisterHsemrlr15FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr15Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr15FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr15FieldCoreidShift))
}

const (
	RegisterHsemrlr15FieldLockShift = 31
	RegisterHsemrlr15FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr15Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr15FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr15Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr15FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr15FieldLockMask)
	}
}

// registerHsemrlr16Type HSEM Read lock register
type registerHsemrlr16Type uint32

const (
	RegisterHsemrlr16FieldProcidShift = 0
	RegisterHsemrlr16FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr16Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr16FieldProcidMask) >> RegisterHsemrlr16FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr16Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr16FieldProcidMask)|(uint32(value)<<RegisterHsemrlr16FieldProcidShift))
}

const (
	RegisterHsemrlr16FieldCoreidShift = 8
	RegisterHsemrlr16FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr16Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr16FieldCoreidMask) >> RegisterHsemrlr16FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr16Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr16FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr16FieldCoreidShift))
}

const (
	RegisterHsemrlr16FieldLockShift = 31
	RegisterHsemrlr16FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr16Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr16FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr16Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr16FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr16FieldLockMask)
	}
}

// registerHsemrlr17Type HSEM Read lock register
type registerHsemrlr17Type uint32

const (
	RegisterHsemrlr17FieldProcidShift = 0
	RegisterHsemrlr17FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr17Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr17FieldProcidMask) >> RegisterHsemrlr17FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr17Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr17FieldProcidMask)|(uint32(value)<<RegisterHsemrlr17FieldProcidShift))
}

const (
	RegisterHsemrlr17FieldCoreidShift = 8
	RegisterHsemrlr17FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr17Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr17FieldCoreidMask) >> RegisterHsemrlr17FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr17Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr17FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr17FieldCoreidShift))
}

const (
	RegisterHsemrlr17FieldLockShift = 31
	RegisterHsemrlr17FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr17Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr17FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr17Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr17FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr17FieldLockMask)
	}
}

// registerHsemrlr18Type HSEM Read lock register
type registerHsemrlr18Type uint32

const (
	RegisterHsemrlr18FieldProcidShift = 0
	RegisterHsemrlr18FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr18Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr18FieldProcidMask) >> RegisterHsemrlr18FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr18Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr18FieldProcidMask)|(uint32(value)<<RegisterHsemrlr18FieldProcidShift))
}

const (
	RegisterHsemrlr18FieldCoreidShift = 8
	RegisterHsemrlr18FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr18Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr18FieldCoreidMask) >> RegisterHsemrlr18FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr18Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr18FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr18FieldCoreidShift))
}

const (
	RegisterHsemrlr18FieldLockShift = 31
	RegisterHsemrlr18FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr18Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr18FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr18Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr18FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr18FieldLockMask)
	}
}

// registerHsemrlr19Type HSEM Read lock register
type registerHsemrlr19Type uint32

const (
	RegisterHsemrlr19FieldProcidShift = 0
	RegisterHsemrlr19FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr19Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr19FieldProcidMask) >> RegisterHsemrlr19FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr19Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr19FieldProcidMask)|(uint32(value)<<RegisterHsemrlr19FieldProcidShift))
}

const (
	RegisterHsemrlr19FieldCoreidShift = 8
	RegisterHsemrlr19FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr19Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr19FieldCoreidMask) >> RegisterHsemrlr19FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr19Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr19FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr19FieldCoreidShift))
}

const (
	RegisterHsemrlr19FieldLockShift = 31
	RegisterHsemrlr19FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr19Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr19FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr19Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr19FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr19FieldLockMask)
	}
}

// registerHsemrlr20Type HSEM Read lock register
type registerHsemrlr20Type uint32

const (
	RegisterHsemrlr20FieldProcidShift = 0
	RegisterHsemrlr20FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr20Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr20FieldProcidMask) >> RegisterHsemrlr20FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr20Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr20FieldProcidMask)|(uint32(value)<<RegisterHsemrlr20FieldProcidShift))
}

const (
	RegisterHsemrlr20FieldCoreidShift = 8
	RegisterHsemrlr20FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr20Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr20FieldCoreidMask) >> RegisterHsemrlr20FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr20Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr20FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr20FieldCoreidShift))
}

const (
	RegisterHsemrlr20FieldLockShift = 31
	RegisterHsemrlr20FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr20Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr20FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr20Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr20FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr20FieldLockMask)
	}
}

// registerHsemrlr21Type HSEM Read lock register
type registerHsemrlr21Type uint32

const (
	RegisterHsemrlr21FieldProcidShift = 0
	RegisterHsemrlr21FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr21Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr21FieldProcidMask) >> RegisterHsemrlr21FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr21Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr21FieldProcidMask)|(uint32(value)<<RegisterHsemrlr21FieldProcidShift))
}

const (
	RegisterHsemrlr21FieldCoreidShift = 8
	RegisterHsemrlr21FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr21Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr21FieldCoreidMask) >> RegisterHsemrlr21FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr21Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr21FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr21FieldCoreidShift))
}

const (
	RegisterHsemrlr21FieldLockShift = 31
	RegisterHsemrlr21FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr21Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr21FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr21Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr21FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr21FieldLockMask)
	}
}

// registerHsemrlr22Type HSEM Read lock register
type registerHsemrlr22Type uint32

const (
	RegisterHsemrlr22FieldProcidShift = 0
	RegisterHsemrlr22FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr22Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr22FieldProcidMask) >> RegisterHsemrlr22FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr22Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr22FieldProcidMask)|(uint32(value)<<RegisterHsemrlr22FieldProcidShift))
}

const (
	RegisterHsemrlr22FieldCoreidShift = 8
	RegisterHsemrlr22FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr22Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr22FieldCoreidMask) >> RegisterHsemrlr22FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr22Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr22FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr22FieldCoreidShift))
}

const (
	RegisterHsemrlr22FieldLockShift = 31
	RegisterHsemrlr22FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr22Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr22FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr22Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr22FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr22FieldLockMask)
	}
}

// registerHsemrlr23Type HSEM Read lock register
type registerHsemrlr23Type uint32

const (
	RegisterHsemrlr23FieldProcidShift = 0
	RegisterHsemrlr23FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr23Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr23FieldProcidMask) >> RegisterHsemrlr23FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr23Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr23FieldProcidMask)|(uint32(value)<<RegisterHsemrlr23FieldProcidShift))
}

const (
	RegisterHsemrlr23FieldCoreidShift = 8
	RegisterHsemrlr23FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr23Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr23FieldCoreidMask) >> RegisterHsemrlr23FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr23Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr23FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr23FieldCoreidShift))
}

const (
	RegisterHsemrlr23FieldLockShift = 31
	RegisterHsemrlr23FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr23Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr23FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr23Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr23FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr23FieldLockMask)
	}
}

// registerHsemrlr24Type HSEM Read lock register
type registerHsemrlr24Type uint32

const (
	RegisterHsemrlr24FieldProcidShift = 0
	RegisterHsemrlr24FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr24Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr24FieldProcidMask) >> RegisterHsemrlr24FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr24Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr24FieldProcidMask)|(uint32(value)<<RegisterHsemrlr24FieldProcidShift))
}

const (
	RegisterHsemrlr24FieldCoreidShift = 8
	RegisterHsemrlr24FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr24Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr24FieldCoreidMask) >> RegisterHsemrlr24FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr24Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr24FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr24FieldCoreidShift))
}

const (
	RegisterHsemrlr24FieldLockShift = 31
	RegisterHsemrlr24FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr24Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr24FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr24Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr24FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr24FieldLockMask)
	}
}

// registerHsemrlr25Type HSEM Read lock register
type registerHsemrlr25Type uint32

const (
	RegisterHsemrlr25FieldProcidShift = 0
	RegisterHsemrlr25FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr25Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr25FieldProcidMask) >> RegisterHsemrlr25FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr25Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr25FieldProcidMask)|(uint32(value)<<RegisterHsemrlr25FieldProcidShift))
}

const (
	RegisterHsemrlr25FieldCoreidShift = 8
	RegisterHsemrlr25FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr25Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr25FieldCoreidMask) >> RegisterHsemrlr25FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr25Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr25FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr25FieldCoreidShift))
}

const (
	RegisterHsemrlr25FieldLockShift = 31
	RegisterHsemrlr25FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr25Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr25FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr25Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr25FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr25FieldLockMask)
	}
}

// registerHsemrlr26Type HSEM Read lock register
type registerHsemrlr26Type uint32

const (
	RegisterHsemrlr26FieldProcidShift = 0
	RegisterHsemrlr26FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr26Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr26FieldProcidMask) >> RegisterHsemrlr26FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr26Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr26FieldProcidMask)|(uint32(value)<<RegisterHsemrlr26FieldProcidShift))
}

const (
	RegisterHsemrlr26FieldCoreidShift = 8
	RegisterHsemrlr26FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr26Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr26FieldCoreidMask) >> RegisterHsemrlr26FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr26Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr26FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr26FieldCoreidShift))
}

const (
	RegisterHsemrlr26FieldLockShift = 31
	RegisterHsemrlr26FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr26Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr26FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr26Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr26FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr26FieldLockMask)
	}
}

// registerHsemrlr27Type HSEM Read lock register
type registerHsemrlr27Type uint32

const (
	RegisterHsemrlr27FieldProcidShift = 0
	RegisterHsemrlr27FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr27Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr27FieldProcidMask) >> RegisterHsemrlr27FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr27Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr27FieldProcidMask)|(uint32(value)<<RegisterHsemrlr27FieldProcidShift))
}

const (
	RegisterHsemrlr27FieldCoreidShift = 8
	RegisterHsemrlr27FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr27Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr27FieldCoreidMask) >> RegisterHsemrlr27FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr27Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr27FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr27FieldCoreidShift))
}

const (
	RegisterHsemrlr27FieldLockShift = 31
	RegisterHsemrlr27FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr27Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr27FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr27Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr27FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr27FieldLockMask)
	}
}

// registerHsemrlr28Type HSEM Read lock register
type registerHsemrlr28Type uint32

const (
	RegisterHsemrlr28FieldProcidShift = 0
	RegisterHsemrlr28FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr28Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr28FieldProcidMask) >> RegisterHsemrlr28FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr28Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr28FieldProcidMask)|(uint32(value)<<RegisterHsemrlr28FieldProcidShift))
}

const (
	RegisterHsemrlr28FieldCoreidShift = 8
	RegisterHsemrlr28FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr28Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr28FieldCoreidMask) >> RegisterHsemrlr28FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr28Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr28FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr28FieldCoreidShift))
}

const (
	RegisterHsemrlr28FieldLockShift = 31
	RegisterHsemrlr28FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr28Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr28FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr28Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr28FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr28FieldLockMask)
	}
}

// registerHsemrlr29Type HSEM Read lock register
type registerHsemrlr29Type uint32

const (
	RegisterHsemrlr29FieldProcidShift = 0
	RegisterHsemrlr29FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr29Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr29FieldProcidMask) >> RegisterHsemrlr29FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr29Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr29FieldProcidMask)|(uint32(value)<<RegisterHsemrlr29FieldProcidShift))
}

const (
	RegisterHsemrlr29FieldCoreidShift = 8
	RegisterHsemrlr29FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr29Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr29FieldCoreidMask) >> RegisterHsemrlr29FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr29Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr29FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr29FieldCoreidShift))
}

const (
	RegisterHsemrlr29FieldLockShift = 31
	RegisterHsemrlr29FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr29Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr29FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr29Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr29FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr29FieldLockMask)
	}
}

// registerHsemrlr30Type HSEM Read lock register
type registerHsemrlr30Type uint32

const (
	RegisterHsemrlr30FieldProcidShift = 0
	RegisterHsemrlr30FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr30Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr30FieldProcidMask) >> RegisterHsemrlr30FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr30Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr30FieldProcidMask)|(uint32(value)<<RegisterHsemrlr30FieldProcidShift))
}

const (
	RegisterHsemrlr30FieldCoreidShift = 8
	RegisterHsemrlr30FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr30Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr30FieldCoreidMask) >> RegisterHsemrlr30FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr30Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr30FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr30FieldCoreidShift))
}

const (
	RegisterHsemrlr30FieldLockShift = 31
	RegisterHsemrlr30FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr30Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr30FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr30Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr30FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr30FieldLockMask)
	}
}

// registerHsemrlr31Type HSEM Read lock register
type registerHsemrlr31Type uint32

const (
	RegisterHsemrlr31FieldProcidShift = 0
	RegisterHsemrlr31FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *registerHsemrlr31Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr31FieldProcidMask) >> RegisterHsemrlr31FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *registerHsemrlr31Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr31FieldProcidMask)|(uint32(value)<<RegisterHsemrlr31FieldProcidShift))
}

const (
	RegisterHsemrlr31FieldCoreidShift = 8
	RegisterHsemrlr31FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *registerHsemrlr31Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr31FieldCoreidMask) >> RegisterHsemrlr31FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *registerHsemrlr31Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr31FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr31FieldCoreidShift))
}

const (
	RegisterHsemrlr31FieldLockShift = 31
	RegisterHsemrlr31FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *registerHsemrlr31Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr31FieldLockMask) != 0
}

// SetLock Lock indication
func (r *registerHsemrlr31Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr31FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr31FieldLockMask)
	}
}

// registerHsemc1ierType HSEM Interrupt enable register
type registerHsemc1ierType uint32

const (
	RegisterHsemc1ierFieldIseShift = 0
	RegisterHsemc1ierFieldIseMask  = 0xffffffff
)

// GetIse Interrupt semaphore x enable bit
func (r *registerHsemc1ierType) GetIse() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc1ierFieldIseMask) >> RegisterHsemc1ierFieldIseShift)
}

// SetIse Interrupt semaphore x enable bit
func (r *registerHsemc1ierType) SetIse(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc1ierFieldIseMask)|(uint32(value)<<RegisterHsemc1ierFieldIseShift))
}

// registerHsemc1icrType HSEM Interrupt clear register
type registerHsemc1icrType uint32

const (
	RegisterHsemc1icrFieldIscShift = 0
	RegisterHsemc1icrFieldIscMask  = 0xffffffff
)

// GetIsc Interrupt semaphore x clear bit
func (r *registerHsemc1icrType) GetIsc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc1icrFieldIscMask) >> RegisterHsemc1icrFieldIscShift)
}

// SetIsc Interrupt semaphore x clear bit
func (r *registerHsemc1icrType) SetIsc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc1icrFieldIscMask)|(uint32(value)<<RegisterHsemc1icrFieldIscShift))
}

// registerHsemc1isrType HSEM Interrupt status register
type registerHsemc1isrType uint32

const (
	RegisterHsemc1isrFieldIsfShift = 0
	RegisterHsemc1isrFieldIsfMask  = 0xffffffff
)

// GetIsf Interrupt semaphore x status bit before enable (mask)
func (r *registerHsemc1isrType) GetIsf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc1isrFieldIsfMask) >> RegisterHsemc1isrFieldIsfShift)
}

// SetIsf Interrupt semaphore x status bit before enable (mask)
func (r *registerHsemc1isrType) SetIsf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc1isrFieldIsfMask)|(uint32(value)<<RegisterHsemc1isrFieldIsfShift))
}

// registerHsemc1misrType HSEM Masked interrupt status register
type registerHsemc1misrType uint32

const (
	RegisterHsemc1misrFieldMisfShift = 0
	RegisterHsemc1misrFieldMisfMask  = 0xffffffff
)

// GetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *registerHsemc1misrType) GetMisf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc1misrFieldMisfMask) >> RegisterHsemc1misrFieldMisfShift)
}

// SetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *registerHsemc1misrType) SetMisf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc1misrFieldMisfMask)|(uint32(value)<<RegisterHsemc1misrFieldMisfShift))
}

// registerHsemc2ierType HSEM Interrupt enable register
type registerHsemc2ierType uint32

const (
	RegisterHsemc2ierFieldIseShift = 0
	RegisterHsemc2ierFieldIseMask  = 0xffffffff
)

// GetIse Interrupt semaphore x enable bit
func (r *registerHsemc2ierType) GetIse() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc2ierFieldIseMask) >> RegisterHsemc2ierFieldIseShift)
}

// SetIse Interrupt semaphore x enable bit
func (r *registerHsemc2ierType) SetIse(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc2ierFieldIseMask)|(uint32(value)<<RegisterHsemc2ierFieldIseShift))
}

// registerHsemc2icrType HSEM Interrupt clear register
type registerHsemc2icrType uint32

const (
	RegisterHsemc2icrFieldIscShift = 0
	RegisterHsemc2icrFieldIscMask  = 0xffffffff
)

// GetIsc Interrupt semaphore x clear bit
func (r *registerHsemc2icrType) GetIsc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc2icrFieldIscMask) >> RegisterHsemc2icrFieldIscShift)
}

// SetIsc Interrupt semaphore x clear bit
func (r *registerHsemc2icrType) SetIsc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc2icrFieldIscMask)|(uint32(value)<<RegisterHsemc2icrFieldIscShift))
}

// registerHsemc2isrType HSEM Interrupt status register
type registerHsemc2isrType uint32

const (
	RegisterHsemc2isrFieldIsfShift = 0
	RegisterHsemc2isrFieldIsfMask  = 0xffffffff
)

// GetIsf Interrupt semaphore x status bit before enable (mask)
func (r *registerHsemc2isrType) GetIsf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc2isrFieldIsfMask) >> RegisterHsemc2isrFieldIsfShift)
}

// SetIsf Interrupt semaphore x status bit before enable (mask)
func (r *registerHsemc2isrType) SetIsf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc2isrFieldIsfMask)|(uint32(value)<<RegisterHsemc2isrFieldIsfShift))
}

// registerHsemc2misrType HSEM Masked interrupt status register
type registerHsemc2misrType uint32

const (
	RegisterHsemc2misrFieldMisfShift = 0
	RegisterHsemc2misrFieldMisfMask  = 0xffffffff
)

// GetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *registerHsemc2misrType) GetMisf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc2misrFieldMisfMask) >> RegisterHsemc2misrFieldMisfShift)
}

// SetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *registerHsemc2misrType) SetMisf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc2misrFieldMisfMask)|(uint32(value)<<RegisterHsemc2misrFieldMisfShift))
}

// registerHsemcrType HSEM Clear register
type registerHsemcrType uint32

const (
	RegisterHsemcrFieldCoreidShift = 8
	RegisterHsemcrFieldCoreidMask  = 0xf00
)

// GetCoreid COREID of semaphores to be cleared
func (r *registerHsemcrType) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemcrFieldCoreidMask) >> RegisterHsemcrFieldCoreidShift)
}

// SetCoreid COREID of semaphores to be cleared
func (r *registerHsemcrType) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemcrFieldCoreidMask)|(uint32(value)<<RegisterHsemcrFieldCoreidShift))
}

const (
	RegisterHsemcrFieldKeyShift = 16
	RegisterHsemcrFieldKeyMask  = 0xffff0000
)

// GetKey Semaphore clear Key
func (r *registerHsemcrType) GetKey() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterHsemcrFieldKeyMask) >> RegisterHsemcrFieldKeyShift)
}

// SetKey Semaphore clear Key
func (r *registerHsemcrType) SetKey(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemcrFieldKeyMask)|(uint32(value)<<RegisterHsemcrFieldKeyShift))
}

// registerHsemkeyrType HSEM Interrupt clear register
type registerHsemkeyrType uint32

const (
	RegisterHsemkeyrFieldKeyShift = 16
	RegisterHsemkeyrFieldKeyMask  = 0xffff0000
)

// GetKey Semaphore Clear Key
func (r *registerHsemkeyrType) GetKey() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterHsemkeyrFieldKeyMask) >> RegisterHsemkeyrFieldKeyShift)
}

// SetKey Semaphore Clear Key
func (r *registerHsemkeyrType) SetKey(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemkeyrFieldKeyMask)|(uint32(value)<<RegisterHsemkeyrFieldKeyShift))
}
