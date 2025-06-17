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
	Hsemr0     RegisterHsemr0Type
	Hsemr1     RegisterHsemr1Type
	Hsemr2     RegisterHsemr2Type
	Hsemr3     RegisterHsemr3Type
	Hsemr4     RegisterHsemr4Type
	Hsemr5     RegisterHsemr5Type
	Hsemr6     RegisterHsemr6Type
	Hsemr7     RegisterHsemr7Type
	Hsemr8     RegisterHsemr8Type
	Hsemr9     RegisterHsemr9Type
	Hsemr10    RegisterHsemr10Type
	Hsemr11    RegisterHsemr11Type
	Hsemr12    RegisterHsemr12Type
	Hsemr13    RegisterHsemr13Type
	Hsemr14    RegisterHsemr14Type
	Hsemr15    RegisterHsemr15Type
	Hsemr16    RegisterHsemr16Type
	Hsemr17    RegisterHsemr17Type
	Hsemr18    RegisterHsemr18Type
	Hsemr19    RegisterHsemr19Type
	Hsemr20    RegisterHsemr20Type
	Hsemr21    RegisterHsemr21Type
	Hsemr22    RegisterHsemr22Type
	Hsemr23    RegisterHsemr23Type
	Hsemr24    RegisterHsemr24Type
	Hsemr25    RegisterHsemr25Type
	Hsemr26    RegisterHsemr26Type
	Hsemr27    RegisterHsemr27Type
	Hsemr28    RegisterHsemr28Type
	Hsemr29    RegisterHsemr29Type
	Hsemr30    RegisterHsemr30Type
	Hsemr31    RegisterHsemr31Type
	Hsemrlr0   RegisterHsemrlr0Type
	Hsemrlr1   RegisterHsemrlr1Type
	Hsemrlr2   RegisterHsemrlr2Type
	Hsemrlr3   RegisterHsemrlr3Type
	Hsemrlr4   RegisterHsemrlr4Type
	Hsemrlr5   RegisterHsemrlr5Type
	Hsemrlr6   RegisterHsemrlr6Type
	Hsemrlr7   RegisterHsemrlr7Type
	Hsemrlr8   RegisterHsemrlr8Type
	Hsemrlr9   RegisterHsemrlr9Type
	Hsemrlr10  RegisterHsemrlr10Type
	Hsemrlr11  RegisterHsemrlr11Type
	Hsemrlr12  RegisterHsemrlr12Type
	Hsemrlr13  RegisterHsemrlr13Type
	Hsemrlr14  RegisterHsemrlr14Type
	Hsemrlr15  RegisterHsemrlr15Type
	Hsemrlr16  RegisterHsemrlr16Type
	Hsemrlr17  RegisterHsemrlr17Type
	Hsemrlr18  RegisterHsemrlr18Type
	Hsemrlr19  RegisterHsemrlr19Type
	Hsemrlr20  RegisterHsemrlr20Type
	Hsemrlr21  RegisterHsemrlr21Type
	Hsemrlr22  RegisterHsemrlr22Type
	Hsemrlr23  RegisterHsemrlr23Type
	Hsemrlr24  RegisterHsemrlr24Type
	Hsemrlr25  RegisterHsemrlr25Type
	Hsemrlr26  RegisterHsemrlr26Type
	Hsemrlr27  RegisterHsemrlr27Type
	Hsemrlr28  RegisterHsemrlr28Type
	Hsemrlr29  RegisterHsemrlr29Type
	Hsemrlr30  RegisterHsemrlr30Type
	Hsemrlr31  RegisterHsemrlr31Type
	Hsemc1ier  RegisterHsemc1ierType
	Hsemc1icr  RegisterHsemc1icrType
	Hsemc1isr  RegisterHsemc1isrType
	Hsemc1misr RegisterHsemc1misrType
	Hsemc2ier  RegisterHsemc2ierType
	Hsemc2icr  RegisterHsemc2icrType
	Hsemc2isr  RegisterHsemc2isrType
	Hsemc2misr RegisterHsemc2misrType
	_          [32]byte
	Hsemcr     RegisterHsemcrType
	Hsemkeyr   RegisterHsemkeyrType
}

// RegisterHsemr0Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr0Type uint32

func (r *RegisterHsemr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr0FieldProcidShift = 0
	RegisterHsemr0FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr0Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr0FieldProcidMask) >> RegisterHsemr0FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr0Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr0FieldProcidMask)|(uint32(value)<<RegisterHsemr0FieldProcidShift))
}

const (
	RegisterHsemr0FieldCoreidShift = 8
	RegisterHsemr0FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr0Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr0FieldCoreidMask) >> RegisterHsemr0FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr0Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr0FieldCoreidMask)|(uint32(value)<<RegisterHsemr0FieldCoreidShift))
}

const (
	RegisterHsemr0FieldLockShift = 31
	RegisterHsemr0FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr0Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr0FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr0Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr0FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr0FieldLockMask)
	}
}

// RegisterHsemr1Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr1Type uint32

func (r *RegisterHsemr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr1FieldProcidShift = 0
	RegisterHsemr1FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr1Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr1FieldProcidMask) >> RegisterHsemr1FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr1Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr1FieldProcidMask)|(uint32(value)<<RegisterHsemr1FieldProcidShift))
}

const (
	RegisterHsemr1FieldCoreidShift = 8
	RegisterHsemr1FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr1Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr1FieldCoreidMask) >> RegisterHsemr1FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr1Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr1FieldCoreidMask)|(uint32(value)<<RegisterHsemr1FieldCoreidShift))
}

const (
	RegisterHsemr1FieldLockShift = 31
	RegisterHsemr1FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr1Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr1FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr1Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr1FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr1FieldLockMask)
	}
}

// RegisterHsemr2Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr2Type uint32

func (r *RegisterHsemr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr2FieldProcidShift = 0
	RegisterHsemr2FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr2Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr2FieldProcidMask) >> RegisterHsemr2FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr2Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr2FieldProcidMask)|(uint32(value)<<RegisterHsemr2FieldProcidShift))
}

const (
	RegisterHsemr2FieldCoreidShift = 8
	RegisterHsemr2FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr2Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr2FieldCoreidMask) >> RegisterHsemr2FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr2Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr2FieldCoreidMask)|(uint32(value)<<RegisterHsemr2FieldCoreidShift))
}

const (
	RegisterHsemr2FieldLockShift = 31
	RegisterHsemr2FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr2Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr2FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr2Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr2FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr2FieldLockMask)
	}
}

// RegisterHsemr3Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr3Type uint32

func (r *RegisterHsemr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr3FieldProcidShift = 0
	RegisterHsemr3FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr3Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr3FieldProcidMask) >> RegisterHsemr3FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr3Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr3FieldProcidMask)|(uint32(value)<<RegisterHsemr3FieldProcidShift))
}

const (
	RegisterHsemr3FieldCoreidShift = 8
	RegisterHsemr3FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr3Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr3FieldCoreidMask) >> RegisterHsemr3FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr3Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr3FieldCoreidMask)|(uint32(value)<<RegisterHsemr3FieldCoreidShift))
}

const (
	RegisterHsemr3FieldLockShift = 31
	RegisterHsemr3FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr3Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr3FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr3Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr3FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr3FieldLockMask)
	}
}

// RegisterHsemr4Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr4Type uint32

func (r *RegisterHsemr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr4FieldProcidShift = 0
	RegisterHsemr4FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr4Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr4FieldProcidMask) >> RegisterHsemr4FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr4Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr4FieldProcidMask)|(uint32(value)<<RegisterHsemr4FieldProcidShift))
}

const (
	RegisterHsemr4FieldCoreidShift = 8
	RegisterHsemr4FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr4Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr4FieldCoreidMask) >> RegisterHsemr4FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr4Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr4FieldCoreidMask)|(uint32(value)<<RegisterHsemr4FieldCoreidShift))
}

const (
	RegisterHsemr4FieldLockShift = 31
	RegisterHsemr4FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr4Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr4FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr4Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr4FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr4FieldLockMask)
	}
}

// RegisterHsemr5Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr5Type uint32

func (r *RegisterHsemr5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr5FieldProcidShift = 0
	RegisterHsemr5FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr5Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr5FieldProcidMask) >> RegisterHsemr5FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr5Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr5FieldProcidMask)|(uint32(value)<<RegisterHsemr5FieldProcidShift))
}

const (
	RegisterHsemr5FieldCoreidShift = 8
	RegisterHsemr5FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr5Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr5FieldCoreidMask) >> RegisterHsemr5FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr5Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr5FieldCoreidMask)|(uint32(value)<<RegisterHsemr5FieldCoreidShift))
}

const (
	RegisterHsemr5FieldLockShift = 31
	RegisterHsemr5FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr5Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr5FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr5Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr5FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr5FieldLockMask)
	}
}

// RegisterHsemr6Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr6Type uint32

func (r *RegisterHsemr6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr6FieldProcidShift = 0
	RegisterHsemr6FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr6Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr6FieldProcidMask) >> RegisterHsemr6FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr6Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr6FieldProcidMask)|(uint32(value)<<RegisterHsemr6FieldProcidShift))
}

const (
	RegisterHsemr6FieldCoreidShift = 8
	RegisterHsemr6FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr6Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr6FieldCoreidMask) >> RegisterHsemr6FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr6Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr6FieldCoreidMask)|(uint32(value)<<RegisterHsemr6FieldCoreidShift))
}

const (
	RegisterHsemr6FieldLockShift = 31
	RegisterHsemr6FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr6Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr6FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr6Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr6FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr6FieldLockMask)
	}
}

// RegisterHsemr7Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr7Type uint32

func (r *RegisterHsemr7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr7FieldProcidShift = 0
	RegisterHsemr7FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr7Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr7FieldProcidMask) >> RegisterHsemr7FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr7Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr7FieldProcidMask)|(uint32(value)<<RegisterHsemr7FieldProcidShift))
}

const (
	RegisterHsemr7FieldCoreidShift = 8
	RegisterHsemr7FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr7Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr7FieldCoreidMask) >> RegisterHsemr7FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr7Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr7FieldCoreidMask)|(uint32(value)<<RegisterHsemr7FieldCoreidShift))
}

const (
	RegisterHsemr7FieldLockShift = 31
	RegisterHsemr7FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr7Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr7FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr7Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr7FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr7FieldLockMask)
	}
}

// RegisterHsemr8Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr8Type uint32

func (r *RegisterHsemr8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr8FieldProcidShift = 0
	RegisterHsemr8FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr8Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr8FieldProcidMask) >> RegisterHsemr8FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr8Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr8FieldProcidMask)|(uint32(value)<<RegisterHsemr8FieldProcidShift))
}

const (
	RegisterHsemr8FieldCoreidShift = 8
	RegisterHsemr8FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr8Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr8FieldCoreidMask) >> RegisterHsemr8FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr8Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr8FieldCoreidMask)|(uint32(value)<<RegisterHsemr8FieldCoreidShift))
}

const (
	RegisterHsemr8FieldLockShift = 31
	RegisterHsemr8FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr8Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr8FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr8Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr8FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr8FieldLockMask)
	}
}

// RegisterHsemr9Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr9Type uint32

func (r *RegisterHsemr9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr9FieldProcidShift = 0
	RegisterHsemr9FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr9Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr9FieldProcidMask) >> RegisterHsemr9FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr9Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr9FieldProcidMask)|(uint32(value)<<RegisterHsemr9FieldProcidShift))
}

const (
	RegisterHsemr9FieldCoreidShift = 8
	RegisterHsemr9FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr9Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr9FieldCoreidMask) >> RegisterHsemr9FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr9Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr9FieldCoreidMask)|(uint32(value)<<RegisterHsemr9FieldCoreidShift))
}

const (
	RegisterHsemr9FieldLockShift = 31
	RegisterHsemr9FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr9Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr9FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr9Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr9FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr9FieldLockMask)
	}
}

// RegisterHsemr10Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr10Type uint32

func (r *RegisterHsemr10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr10FieldProcidShift = 0
	RegisterHsemr10FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr10Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr10FieldProcidMask) >> RegisterHsemr10FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr10Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr10FieldProcidMask)|(uint32(value)<<RegisterHsemr10FieldProcidShift))
}

const (
	RegisterHsemr10FieldCoreidShift = 8
	RegisterHsemr10FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr10Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr10FieldCoreidMask) >> RegisterHsemr10FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr10Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr10FieldCoreidMask)|(uint32(value)<<RegisterHsemr10FieldCoreidShift))
}

const (
	RegisterHsemr10FieldLockShift = 31
	RegisterHsemr10FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr10Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr10FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr10Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr10FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr10FieldLockMask)
	}
}

// RegisterHsemr11Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr11Type uint32

func (r *RegisterHsemr11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr11FieldProcidShift = 0
	RegisterHsemr11FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr11Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr11FieldProcidMask) >> RegisterHsemr11FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr11Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr11FieldProcidMask)|(uint32(value)<<RegisterHsemr11FieldProcidShift))
}

const (
	RegisterHsemr11FieldCoreidShift = 8
	RegisterHsemr11FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr11Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr11FieldCoreidMask) >> RegisterHsemr11FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr11Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr11FieldCoreidMask)|(uint32(value)<<RegisterHsemr11FieldCoreidShift))
}

const (
	RegisterHsemr11FieldLockShift = 31
	RegisterHsemr11FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr11Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr11FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr11Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr11FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr11FieldLockMask)
	}
}

// RegisterHsemr12Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr12Type uint32

func (r *RegisterHsemr12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr12FieldProcidShift = 0
	RegisterHsemr12FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr12Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr12FieldProcidMask) >> RegisterHsemr12FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr12Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr12FieldProcidMask)|(uint32(value)<<RegisterHsemr12FieldProcidShift))
}

const (
	RegisterHsemr12FieldCoreidShift = 8
	RegisterHsemr12FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr12Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr12FieldCoreidMask) >> RegisterHsemr12FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr12Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr12FieldCoreidMask)|(uint32(value)<<RegisterHsemr12FieldCoreidShift))
}

const (
	RegisterHsemr12FieldLockShift = 31
	RegisterHsemr12FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr12Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr12FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr12Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr12FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr12FieldLockMask)
	}
}

// RegisterHsemr13Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr13Type uint32

func (r *RegisterHsemr13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr13FieldProcidShift = 0
	RegisterHsemr13FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr13Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr13FieldProcidMask) >> RegisterHsemr13FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr13Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr13FieldProcidMask)|(uint32(value)<<RegisterHsemr13FieldProcidShift))
}

const (
	RegisterHsemr13FieldCoreidShift = 8
	RegisterHsemr13FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr13Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr13FieldCoreidMask) >> RegisterHsemr13FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr13Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr13FieldCoreidMask)|(uint32(value)<<RegisterHsemr13FieldCoreidShift))
}

const (
	RegisterHsemr13FieldLockShift = 31
	RegisterHsemr13FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr13Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr13FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr13Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr13FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr13FieldLockMask)
	}
}

// RegisterHsemr14Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr14Type uint32

func (r *RegisterHsemr14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr14FieldProcidShift = 0
	RegisterHsemr14FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr14Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr14FieldProcidMask) >> RegisterHsemr14FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr14Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr14FieldProcidMask)|(uint32(value)<<RegisterHsemr14FieldProcidShift))
}

const (
	RegisterHsemr14FieldCoreidShift = 8
	RegisterHsemr14FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr14Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr14FieldCoreidMask) >> RegisterHsemr14FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr14Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr14FieldCoreidMask)|(uint32(value)<<RegisterHsemr14FieldCoreidShift))
}

const (
	RegisterHsemr14FieldLockShift = 31
	RegisterHsemr14FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr14Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr14FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr14Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr14FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr14FieldLockMask)
	}
}

// RegisterHsemr15Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr15Type uint32

func (r *RegisterHsemr15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr15FieldProcidShift = 0
	RegisterHsemr15FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr15Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr15FieldProcidMask) >> RegisterHsemr15FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr15Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr15FieldProcidMask)|(uint32(value)<<RegisterHsemr15FieldProcidShift))
}

const (
	RegisterHsemr15FieldCoreidShift = 8
	RegisterHsemr15FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr15Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr15FieldCoreidMask) >> RegisterHsemr15FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr15Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr15FieldCoreidMask)|(uint32(value)<<RegisterHsemr15FieldCoreidShift))
}

const (
	RegisterHsemr15FieldLockShift = 31
	RegisterHsemr15FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr15Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr15FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr15Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr15FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr15FieldLockMask)
	}
}

// RegisterHsemr16Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr16Type uint32

func (r *RegisterHsemr16Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr16Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr16Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr16Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr16Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr16FieldProcidShift = 0
	RegisterHsemr16FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr16Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr16FieldProcidMask) >> RegisterHsemr16FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr16Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr16FieldProcidMask)|(uint32(value)<<RegisterHsemr16FieldProcidShift))
}

const (
	RegisterHsemr16FieldCoreidShift = 8
	RegisterHsemr16FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr16Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr16FieldCoreidMask) >> RegisterHsemr16FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr16Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr16FieldCoreidMask)|(uint32(value)<<RegisterHsemr16FieldCoreidShift))
}

const (
	RegisterHsemr16FieldLockShift = 31
	RegisterHsemr16FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr16Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr16FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr16Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr16FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr16FieldLockMask)
	}
}

// RegisterHsemr17Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr17Type uint32

func (r *RegisterHsemr17Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr17Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr17Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr17Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr17Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr17FieldProcidShift = 0
	RegisterHsemr17FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr17Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr17FieldProcidMask) >> RegisterHsemr17FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr17Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr17FieldProcidMask)|(uint32(value)<<RegisterHsemr17FieldProcidShift))
}

const (
	RegisterHsemr17FieldCoreidShift = 8
	RegisterHsemr17FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr17Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr17FieldCoreidMask) >> RegisterHsemr17FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr17Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr17FieldCoreidMask)|(uint32(value)<<RegisterHsemr17FieldCoreidShift))
}

const (
	RegisterHsemr17FieldLockShift = 31
	RegisterHsemr17FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr17Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr17FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr17Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr17FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr17FieldLockMask)
	}
}

// RegisterHsemr18Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr18Type uint32

func (r *RegisterHsemr18Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr18Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr18Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr18Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr18Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr18FieldProcidShift = 0
	RegisterHsemr18FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr18Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr18FieldProcidMask) >> RegisterHsemr18FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr18Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr18FieldProcidMask)|(uint32(value)<<RegisterHsemr18FieldProcidShift))
}

const (
	RegisterHsemr18FieldCoreidShift = 8
	RegisterHsemr18FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr18Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr18FieldCoreidMask) >> RegisterHsemr18FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr18Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr18FieldCoreidMask)|(uint32(value)<<RegisterHsemr18FieldCoreidShift))
}

const (
	RegisterHsemr18FieldLockShift = 31
	RegisterHsemr18FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr18Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr18FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr18Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr18FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr18FieldLockMask)
	}
}

// RegisterHsemr19Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr19Type uint32

func (r *RegisterHsemr19Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr19Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr19Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr19Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr19Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr19FieldProcidShift = 0
	RegisterHsemr19FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr19Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr19FieldProcidMask) >> RegisterHsemr19FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr19Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr19FieldProcidMask)|(uint32(value)<<RegisterHsemr19FieldProcidShift))
}

const (
	RegisterHsemr19FieldCoreidShift = 8
	RegisterHsemr19FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr19Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr19FieldCoreidMask) >> RegisterHsemr19FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr19Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr19FieldCoreidMask)|(uint32(value)<<RegisterHsemr19FieldCoreidShift))
}

const (
	RegisterHsemr19FieldLockShift = 31
	RegisterHsemr19FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr19Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr19FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr19Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr19FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr19FieldLockMask)
	}
}

// RegisterHsemr20Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr20Type uint32

func (r *RegisterHsemr20Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr20Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr20Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr20Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr20Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr20FieldProcidShift = 0
	RegisterHsemr20FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr20Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr20FieldProcidMask) >> RegisterHsemr20FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr20Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr20FieldProcidMask)|(uint32(value)<<RegisterHsemr20FieldProcidShift))
}

const (
	RegisterHsemr20FieldCoreidShift = 8
	RegisterHsemr20FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr20Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr20FieldCoreidMask) >> RegisterHsemr20FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr20Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr20FieldCoreidMask)|(uint32(value)<<RegisterHsemr20FieldCoreidShift))
}

const (
	RegisterHsemr20FieldLockShift = 31
	RegisterHsemr20FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr20Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr20FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr20Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr20FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr20FieldLockMask)
	}
}

// RegisterHsemr21Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr21Type uint32

func (r *RegisterHsemr21Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr21Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr21Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr21Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr21Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr21FieldProcidShift = 0
	RegisterHsemr21FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr21Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr21FieldProcidMask) >> RegisterHsemr21FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr21Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr21FieldProcidMask)|(uint32(value)<<RegisterHsemr21FieldProcidShift))
}

const (
	RegisterHsemr21FieldCoreidShift = 8
	RegisterHsemr21FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr21Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr21FieldCoreidMask) >> RegisterHsemr21FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr21Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr21FieldCoreidMask)|(uint32(value)<<RegisterHsemr21FieldCoreidShift))
}

const (
	RegisterHsemr21FieldLockShift = 31
	RegisterHsemr21FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr21Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr21FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr21Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr21FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr21FieldLockMask)
	}
}

// RegisterHsemr22Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr22Type uint32

func (r *RegisterHsemr22Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr22Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr22Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr22Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr22Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr22FieldProcidShift = 0
	RegisterHsemr22FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr22Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr22FieldProcidMask) >> RegisterHsemr22FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr22Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr22FieldProcidMask)|(uint32(value)<<RegisterHsemr22FieldProcidShift))
}

const (
	RegisterHsemr22FieldCoreidShift = 8
	RegisterHsemr22FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr22Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr22FieldCoreidMask) >> RegisterHsemr22FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr22Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr22FieldCoreidMask)|(uint32(value)<<RegisterHsemr22FieldCoreidShift))
}

const (
	RegisterHsemr22FieldLockShift = 31
	RegisterHsemr22FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr22Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr22FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr22Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr22FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr22FieldLockMask)
	}
}

// RegisterHsemr23Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr23Type uint32

func (r *RegisterHsemr23Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr23Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr23Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr23Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr23Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr23FieldProcidShift = 0
	RegisterHsemr23FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr23Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr23FieldProcidMask) >> RegisterHsemr23FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr23Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr23FieldProcidMask)|(uint32(value)<<RegisterHsemr23FieldProcidShift))
}

const (
	RegisterHsemr23FieldCoreidShift = 8
	RegisterHsemr23FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr23Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr23FieldCoreidMask) >> RegisterHsemr23FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr23Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr23FieldCoreidMask)|(uint32(value)<<RegisterHsemr23FieldCoreidShift))
}

const (
	RegisterHsemr23FieldLockShift = 31
	RegisterHsemr23FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr23Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr23FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr23Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr23FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr23FieldLockMask)
	}
}

// RegisterHsemr24Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr24Type uint32

func (r *RegisterHsemr24Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr24Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr24Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr24Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr24Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr24FieldProcidShift = 0
	RegisterHsemr24FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr24Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr24FieldProcidMask) >> RegisterHsemr24FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr24Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr24FieldProcidMask)|(uint32(value)<<RegisterHsemr24FieldProcidShift))
}

const (
	RegisterHsemr24FieldCoreidShift = 8
	RegisterHsemr24FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr24Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr24FieldCoreidMask) >> RegisterHsemr24FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr24Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr24FieldCoreidMask)|(uint32(value)<<RegisterHsemr24FieldCoreidShift))
}

const (
	RegisterHsemr24FieldLockShift = 31
	RegisterHsemr24FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr24Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr24FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr24Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr24FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr24FieldLockMask)
	}
}

// RegisterHsemr25Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr25Type uint32

func (r *RegisterHsemr25Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr25Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr25Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr25Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr25Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr25FieldProcidShift = 0
	RegisterHsemr25FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr25Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr25FieldProcidMask) >> RegisterHsemr25FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr25Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr25FieldProcidMask)|(uint32(value)<<RegisterHsemr25FieldProcidShift))
}

const (
	RegisterHsemr25FieldCoreidShift = 8
	RegisterHsemr25FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr25Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr25FieldCoreidMask) >> RegisterHsemr25FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr25Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr25FieldCoreidMask)|(uint32(value)<<RegisterHsemr25FieldCoreidShift))
}

const (
	RegisterHsemr25FieldLockShift = 31
	RegisterHsemr25FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr25Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr25FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr25Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr25FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr25FieldLockMask)
	}
}

// RegisterHsemr26Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr26Type uint32

func (r *RegisterHsemr26Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr26Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr26Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr26Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr26Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr26FieldProcidShift = 0
	RegisterHsemr26FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr26Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr26FieldProcidMask) >> RegisterHsemr26FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr26Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr26FieldProcidMask)|(uint32(value)<<RegisterHsemr26FieldProcidShift))
}

const (
	RegisterHsemr26FieldCoreidShift = 8
	RegisterHsemr26FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr26Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr26FieldCoreidMask) >> RegisterHsemr26FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr26Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr26FieldCoreidMask)|(uint32(value)<<RegisterHsemr26FieldCoreidShift))
}

const (
	RegisterHsemr26FieldLockShift = 31
	RegisterHsemr26FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr26Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr26FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr26Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr26FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr26FieldLockMask)
	}
}

// RegisterHsemr27Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr27Type uint32

func (r *RegisterHsemr27Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr27Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr27Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr27Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr27Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr27FieldProcidShift = 0
	RegisterHsemr27FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr27Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr27FieldProcidMask) >> RegisterHsemr27FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr27Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr27FieldProcidMask)|(uint32(value)<<RegisterHsemr27FieldProcidShift))
}

const (
	RegisterHsemr27FieldCoreidShift = 8
	RegisterHsemr27FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr27Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr27FieldCoreidMask) >> RegisterHsemr27FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr27Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr27FieldCoreidMask)|(uint32(value)<<RegisterHsemr27FieldCoreidShift))
}

const (
	RegisterHsemr27FieldLockShift = 31
	RegisterHsemr27FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr27Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr27FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr27Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr27FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr27FieldLockMask)
	}
}

// RegisterHsemr28Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr28Type uint32

func (r *RegisterHsemr28Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr28Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr28Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr28Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr28Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr28FieldProcidShift = 0
	RegisterHsemr28FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr28Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr28FieldProcidMask) >> RegisterHsemr28FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr28Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr28FieldProcidMask)|(uint32(value)<<RegisterHsemr28FieldProcidShift))
}

const (
	RegisterHsemr28FieldCoreidShift = 8
	RegisterHsemr28FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr28Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr28FieldCoreidMask) >> RegisterHsemr28FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr28Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr28FieldCoreidMask)|(uint32(value)<<RegisterHsemr28FieldCoreidShift))
}

const (
	RegisterHsemr28FieldLockShift = 31
	RegisterHsemr28FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr28Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr28FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr28Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr28FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr28FieldLockMask)
	}
}

// RegisterHsemr29Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr29Type uint32

func (r *RegisterHsemr29Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr29Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr29Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr29Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr29Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr29FieldProcidShift = 0
	RegisterHsemr29FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr29Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr29FieldProcidMask) >> RegisterHsemr29FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr29Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr29FieldProcidMask)|(uint32(value)<<RegisterHsemr29FieldProcidShift))
}

const (
	RegisterHsemr29FieldCoreidShift = 8
	RegisterHsemr29FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr29Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr29FieldCoreidMask) >> RegisterHsemr29FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr29Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr29FieldCoreidMask)|(uint32(value)<<RegisterHsemr29FieldCoreidShift))
}

const (
	RegisterHsemr29FieldLockShift = 31
	RegisterHsemr29FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr29Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr29FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr29Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr29FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr29FieldLockMask)
	}
}

// RegisterHsemr30Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr30Type uint32

func (r *RegisterHsemr30Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr30Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr30Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr30Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr30Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr30FieldProcidShift = 0
	RegisterHsemr30FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr30Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr30FieldProcidMask) >> RegisterHsemr30FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr30Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr30FieldProcidMask)|(uint32(value)<<RegisterHsemr30FieldProcidShift))
}

const (
	RegisterHsemr30FieldCoreidShift = 8
	RegisterHsemr30FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr30Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr30FieldCoreidMask) >> RegisterHsemr30FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr30Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr30FieldCoreidMask)|(uint32(value)<<RegisterHsemr30FieldCoreidShift))
}

const (
	RegisterHsemr30FieldLockShift = 31
	RegisterHsemr30FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr30Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr30FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr30Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr30FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr30FieldLockMask)
	}
}

// RegisterHsemr31Type HSEM register HSEM_R0 HSEM_R31
type RegisterHsemr31Type uint32

func (r *RegisterHsemr31Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemr31Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemr31Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemr31Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemr31Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemr31FieldProcidShift = 0
	RegisterHsemr31FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemr31Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr31FieldProcidMask) >> RegisterHsemr31FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemr31Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr31FieldProcidMask)|(uint32(value)<<RegisterHsemr31FieldProcidShift))
}

const (
	RegisterHsemr31FieldCoreidShift = 8
	RegisterHsemr31FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemr31Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemr31FieldCoreidMask) >> RegisterHsemr31FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemr31Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemr31FieldCoreidMask)|(uint32(value)<<RegisterHsemr31FieldCoreidShift))
}

const (
	RegisterHsemr31FieldLockShift = 31
	RegisterHsemr31FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemr31Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemr31FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemr31Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemr31FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemr31FieldLockMask)
	}
}

// RegisterHsemrlr0Type HSEM Read lock register
type RegisterHsemrlr0Type uint32

func (r *RegisterHsemrlr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr0FieldProcidShift = 0
	RegisterHsemrlr0FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr0Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr0FieldProcidMask) >> RegisterHsemrlr0FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr0Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr0FieldProcidMask)|(uint32(value)<<RegisterHsemrlr0FieldProcidShift))
}

const (
	RegisterHsemrlr0FieldCoreidShift = 8
	RegisterHsemrlr0FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr0Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr0FieldCoreidMask) >> RegisterHsemrlr0FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr0Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr0FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr0FieldCoreidShift))
}

const (
	RegisterHsemrlr0FieldLockShift = 31
	RegisterHsemrlr0FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr0Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr0FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr0Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr0FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr0FieldLockMask)
	}
}

// RegisterHsemrlr1Type HSEM Read lock register
type RegisterHsemrlr1Type uint32

func (r *RegisterHsemrlr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr1FieldProcidShift = 0
	RegisterHsemrlr1FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr1Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr1FieldProcidMask) >> RegisterHsemrlr1FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr1Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr1FieldProcidMask)|(uint32(value)<<RegisterHsemrlr1FieldProcidShift))
}

const (
	RegisterHsemrlr1FieldCoreidShift = 8
	RegisterHsemrlr1FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr1Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr1FieldCoreidMask) >> RegisterHsemrlr1FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr1Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr1FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr1FieldCoreidShift))
}

const (
	RegisterHsemrlr1FieldLockShift = 31
	RegisterHsemrlr1FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr1Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr1FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr1Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr1FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr1FieldLockMask)
	}
}

// RegisterHsemrlr2Type HSEM Read lock register
type RegisterHsemrlr2Type uint32

func (r *RegisterHsemrlr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr2FieldProcidShift = 0
	RegisterHsemrlr2FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr2Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr2FieldProcidMask) >> RegisterHsemrlr2FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr2Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr2FieldProcidMask)|(uint32(value)<<RegisterHsemrlr2FieldProcidShift))
}

const (
	RegisterHsemrlr2FieldCoreidShift = 8
	RegisterHsemrlr2FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr2Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr2FieldCoreidMask) >> RegisterHsemrlr2FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr2Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr2FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr2FieldCoreidShift))
}

const (
	RegisterHsemrlr2FieldLockShift = 31
	RegisterHsemrlr2FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr2Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr2FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr2Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr2FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr2FieldLockMask)
	}
}

// RegisterHsemrlr3Type HSEM Read lock register
type RegisterHsemrlr3Type uint32

func (r *RegisterHsemrlr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr3FieldProcidShift = 0
	RegisterHsemrlr3FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr3Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr3FieldProcidMask) >> RegisterHsemrlr3FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr3Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr3FieldProcidMask)|(uint32(value)<<RegisterHsemrlr3FieldProcidShift))
}

const (
	RegisterHsemrlr3FieldCoreidShift = 8
	RegisterHsemrlr3FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr3Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr3FieldCoreidMask) >> RegisterHsemrlr3FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr3Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr3FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr3FieldCoreidShift))
}

const (
	RegisterHsemrlr3FieldLockShift = 31
	RegisterHsemrlr3FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr3Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr3FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr3Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr3FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr3FieldLockMask)
	}
}

// RegisterHsemrlr4Type HSEM Read lock register
type RegisterHsemrlr4Type uint32

func (r *RegisterHsemrlr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr4FieldProcidShift = 0
	RegisterHsemrlr4FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr4Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr4FieldProcidMask) >> RegisterHsemrlr4FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr4Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr4FieldProcidMask)|(uint32(value)<<RegisterHsemrlr4FieldProcidShift))
}

const (
	RegisterHsemrlr4FieldCoreidShift = 8
	RegisterHsemrlr4FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr4Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr4FieldCoreidMask) >> RegisterHsemrlr4FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr4Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr4FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr4FieldCoreidShift))
}

const (
	RegisterHsemrlr4FieldLockShift = 31
	RegisterHsemrlr4FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr4Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr4FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr4Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr4FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr4FieldLockMask)
	}
}

// RegisterHsemrlr5Type HSEM Read lock register
type RegisterHsemrlr5Type uint32

func (r *RegisterHsemrlr5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr5FieldProcidShift = 0
	RegisterHsemrlr5FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr5Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr5FieldProcidMask) >> RegisterHsemrlr5FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr5Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr5FieldProcidMask)|(uint32(value)<<RegisterHsemrlr5FieldProcidShift))
}

const (
	RegisterHsemrlr5FieldCoreidShift = 8
	RegisterHsemrlr5FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr5Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr5FieldCoreidMask) >> RegisterHsemrlr5FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr5Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr5FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr5FieldCoreidShift))
}

const (
	RegisterHsemrlr5FieldLockShift = 31
	RegisterHsemrlr5FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr5Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr5FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr5Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr5FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr5FieldLockMask)
	}
}

// RegisterHsemrlr6Type HSEM Read lock register
type RegisterHsemrlr6Type uint32

func (r *RegisterHsemrlr6Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr6Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr6Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr6Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr6Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr6FieldProcidShift = 0
	RegisterHsemrlr6FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr6Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr6FieldProcidMask) >> RegisterHsemrlr6FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr6Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr6FieldProcidMask)|(uint32(value)<<RegisterHsemrlr6FieldProcidShift))
}

const (
	RegisterHsemrlr6FieldCoreidShift = 8
	RegisterHsemrlr6FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr6Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr6FieldCoreidMask) >> RegisterHsemrlr6FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr6Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr6FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr6FieldCoreidShift))
}

const (
	RegisterHsemrlr6FieldLockShift = 31
	RegisterHsemrlr6FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr6Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr6FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr6Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr6FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr6FieldLockMask)
	}
}

// RegisterHsemrlr7Type HSEM Read lock register
type RegisterHsemrlr7Type uint32

func (r *RegisterHsemrlr7Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr7Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr7Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr7Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr7Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr7FieldProcidShift = 0
	RegisterHsemrlr7FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr7Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr7FieldProcidMask) >> RegisterHsemrlr7FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr7Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr7FieldProcidMask)|(uint32(value)<<RegisterHsemrlr7FieldProcidShift))
}

const (
	RegisterHsemrlr7FieldCoreidShift = 8
	RegisterHsemrlr7FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr7Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr7FieldCoreidMask) >> RegisterHsemrlr7FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr7Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr7FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr7FieldCoreidShift))
}

const (
	RegisterHsemrlr7FieldLockShift = 31
	RegisterHsemrlr7FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr7Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr7FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr7Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr7FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr7FieldLockMask)
	}
}

// RegisterHsemrlr8Type HSEM Read lock register
type RegisterHsemrlr8Type uint32

func (r *RegisterHsemrlr8Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr8Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr8Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr8Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr8Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr8FieldProcidShift = 0
	RegisterHsemrlr8FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr8Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr8FieldProcidMask) >> RegisterHsemrlr8FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr8Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr8FieldProcidMask)|(uint32(value)<<RegisterHsemrlr8FieldProcidShift))
}

const (
	RegisterHsemrlr8FieldCoreidShift = 8
	RegisterHsemrlr8FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr8Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr8FieldCoreidMask) >> RegisterHsemrlr8FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr8Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr8FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr8FieldCoreidShift))
}

const (
	RegisterHsemrlr8FieldLockShift = 31
	RegisterHsemrlr8FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr8Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr8FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr8Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr8FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr8FieldLockMask)
	}
}

// RegisterHsemrlr9Type HSEM Read lock register
type RegisterHsemrlr9Type uint32

func (r *RegisterHsemrlr9Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr9Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr9Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr9Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr9Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr9FieldProcidShift = 0
	RegisterHsemrlr9FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr9Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr9FieldProcidMask) >> RegisterHsemrlr9FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr9Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr9FieldProcidMask)|(uint32(value)<<RegisterHsemrlr9FieldProcidShift))
}

const (
	RegisterHsemrlr9FieldCoreidShift = 8
	RegisterHsemrlr9FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr9Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr9FieldCoreidMask) >> RegisterHsemrlr9FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr9Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr9FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr9FieldCoreidShift))
}

const (
	RegisterHsemrlr9FieldLockShift = 31
	RegisterHsemrlr9FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr9Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr9FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr9Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr9FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr9FieldLockMask)
	}
}

// RegisterHsemrlr10Type HSEM Read lock register
type RegisterHsemrlr10Type uint32

func (r *RegisterHsemrlr10Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr10Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr10Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr10Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr10Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr10FieldProcidShift = 0
	RegisterHsemrlr10FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr10Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr10FieldProcidMask) >> RegisterHsemrlr10FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr10Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr10FieldProcidMask)|(uint32(value)<<RegisterHsemrlr10FieldProcidShift))
}

const (
	RegisterHsemrlr10FieldCoreidShift = 8
	RegisterHsemrlr10FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr10Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr10FieldCoreidMask) >> RegisterHsemrlr10FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr10Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr10FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr10FieldCoreidShift))
}

const (
	RegisterHsemrlr10FieldLockShift = 31
	RegisterHsemrlr10FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr10Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr10FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr10Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr10FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr10FieldLockMask)
	}
}

// RegisterHsemrlr11Type HSEM Read lock register
type RegisterHsemrlr11Type uint32

func (r *RegisterHsemrlr11Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr11Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr11Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr11Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr11Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr11FieldProcidShift = 0
	RegisterHsemrlr11FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr11Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr11FieldProcidMask) >> RegisterHsemrlr11FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr11Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr11FieldProcidMask)|(uint32(value)<<RegisterHsemrlr11FieldProcidShift))
}

const (
	RegisterHsemrlr11FieldCoreidShift = 8
	RegisterHsemrlr11FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr11Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr11FieldCoreidMask) >> RegisterHsemrlr11FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr11Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr11FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr11FieldCoreidShift))
}

const (
	RegisterHsemrlr11FieldLockShift = 31
	RegisterHsemrlr11FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr11Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr11FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr11Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr11FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr11FieldLockMask)
	}
}

// RegisterHsemrlr12Type HSEM Read lock register
type RegisterHsemrlr12Type uint32

func (r *RegisterHsemrlr12Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr12Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr12Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr12Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr12Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr12FieldProcidShift = 0
	RegisterHsemrlr12FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr12Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr12FieldProcidMask) >> RegisterHsemrlr12FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr12Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr12FieldProcidMask)|(uint32(value)<<RegisterHsemrlr12FieldProcidShift))
}

const (
	RegisterHsemrlr12FieldCoreidShift = 8
	RegisterHsemrlr12FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr12Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr12FieldCoreidMask) >> RegisterHsemrlr12FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr12Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr12FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr12FieldCoreidShift))
}

const (
	RegisterHsemrlr12FieldLockShift = 31
	RegisterHsemrlr12FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr12Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr12FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr12Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr12FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr12FieldLockMask)
	}
}

// RegisterHsemrlr13Type HSEM Read lock register
type RegisterHsemrlr13Type uint32

func (r *RegisterHsemrlr13Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr13Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr13Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr13Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr13Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr13FieldProcidShift = 0
	RegisterHsemrlr13FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr13Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr13FieldProcidMask) >> RegisterHsemrlr13FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr13Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr13FieldProcidMask)|(uint32(value)<<RegisterHsemrlr13FieldProcidShift))
}

const (
	RegisterHsemrlr13FieldCoreidShift = 8
	RegisterHsemrlr13FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr13Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr13FieldCoreidMask) >> RegisterHsemrlr13FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr13Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr13FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr13FieldCoreidShift))
}

const (
	RegisterHsemrlr13FieldLockShift = 31
	RegisterHsemrlr13FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr13Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr13FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr13Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr13FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr13FieldLockMask)
	}
}

// RegisterHsemrlr14Type HSEM Read lock register
type RegisterHsemrlr14Type uint32

func (r *RegisterHsemrlr14Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr14Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr14Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr14Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr14Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr14FieldProcidShift = 0
	RegisterHsemrlr14FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr14Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr14FieldProcidMask) >> RegisterHsemrlr14FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr14Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr14FieldProcidMask)|(uint32(value)<<RegisterHsemrlr14FieldProcidShift))
}

const (
	RegisterHsemrlr14FieldCoreidShift = 8
	RegisterHsemrlr14FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr14Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr14FieldCoreidMask) >> RegisterHsemrlr14FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr14Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr14FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr14FieldCoreidShift))
}

const (
	RegisterHsemrlr14FieldLockShift = 31
	RegisterHsemrlr14FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr14Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr14FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr14Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr14FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr14FieldLockMask)
	}
}

// RegisterHsemrlr15Type HSEM Read lock register
type RegisterHsemrlr15Type uint32

func (r *RegisterHsemrlr15Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr15Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr15Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr15Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr15Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr15FieldProcidShift = 0
	RegisterHsemrlr15FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr15Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr15FieldProcidMask) >> RegisterHsemrlr15FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr15Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr15FieldProcidMask)|(uint32(value)<<RegisterHsemrlr15FieldProcidShift))
}

const (
	RegisterHsemrlr15FieldCoreidShift = 8
	RegisterHsemrlr15FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr15Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr15FieldCoreidMask) >> RegisterHsemrlr15FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr15Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr15FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr15FieldCoreidShift))
}

const (
	RegisterHsemrlr15FieldLockShift = 31
	RegisterHsemrlr15FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr15Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr15FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr15Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr15FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr15FieldLockMask)
	}
}

// RegisterHsemrlr16Type HSEM Read lock register
type RegisterHsemrlr16Type uint32

func (r *RegisterHsemrlr16Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr16Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr16Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr16Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr16Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr16FieldProcidShift = 0
	RegisterHsemrlr16FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr16Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr16FieldProcidMask) >> RegisterHsemrlr16FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr16Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr16FieldProcidMask)|(uint32(value)<<RegisterHsemrlr16FieldProcidShift))
}

const (
	RegisterHsemrlr16FieldCoreidShift = 8
	RegisterHsemrlr16FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr16Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr16FieldCoreidMask) >> RegisterHsemrlr16FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr16Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr16FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr16FieldCoreidShift))
}

const (
	RegisterHsemrlr16FieldLockShift = 31
	RegisterHsemrlr16FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr16Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr16FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr16Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr16FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr16FieldLockMask)
	}
}

// RegisterHsemrlr17Type HSEM Read lock register
type RegisterHsemrlr17Type uint32

func (r *RegisterHsemrlr17Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr17Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr17Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr17Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr17Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr17FieldProcidShift = 0
	RegisterHsemrlr17FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr17Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr17FieldProcidMask) >> RegisterHsemrlr17FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr17Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr17FieldProcidMask)|(uint32(value)<<RegisterHsemrlr17FieldProcidShift))
}

const (
	RegisterHsemrlr17FieldCoreidShift = 8
	RegisterHsemrlr17FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr17Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr17FieldCoreidMask) >> RegisterHsemrlr17FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr17Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr17FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr17FieldCoreidShift))
}

const (
	RegisterHsemrlr17FieldLockShift = 31
	RegisterHsemrlr17FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr17Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr17FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr17Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr17FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr17FieldLockMask)
	}
}

// RegisterHsemrlr18Type HSEM Read lock register
type RegisterHsemrlr18Type uint32

func (r *RegisterHsemrlr18Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr18Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr18Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr18Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr18Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr18FieldProcidShift = 0
	RegisterHsemrlr18FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr18Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr18FieldProcidMask) >> RegisterHsemrlr18FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr18Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr18FieldProcidMask)|(uint32(value)<<RegisterHsemrlr18FieldProcidShift))
}

const (
	RegisterHsemrlr18FieldCoreidShift = 8
	RegisterHsemrlr18FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr18Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr18FieldCoreidMask) >> RegisterHsemrlr18FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr18Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr18FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr18FieldCoreidShift))
}

const (
	RegisterHsemrlr18FieldLockShift = 31
	RegisterHsemrlr18FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr18Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr18FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr18Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr18FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr18FieldLockMask)
	}
}

// RegisterHsemrlr19Type HSEM Read lock register
type RegisterHsemrlr19Type uint32

func (r *RegisterHsemrlr19Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr19Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr19Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr19Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr19Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr19FieldProcidShift = 0
	RegisterHsemrlr19FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr19Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr19FieldProcidMask) >> RegisterHsemrlr19FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr19Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr19FieldProcidMask)|(uint32(value)<<RegisterHsemrlr19FieldProcidShift))
}

const (
	RegisterHsemrlr19FieldCoreidShift = 8
	RegisterHsemrlr19FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr19Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr19FieldCoreidMask) >> RegisterHsemrlr19FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr19Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr19FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr19FieldCoreidShift))
}

const (
	RegisterHsemrlr19FieldLockShift = 31
	RegisterHsemrlr19FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr19Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr19FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr19Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr19FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr19FieldLockMask)
	}
}

// RegisterHsemrlr20Type HSEM Read lock register
type RegisterHsemrlr20Type uint32

func (r *RegisterHsemrlr20Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr20Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr20Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr20Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr20Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr20FieldProcidShift = 0
	RegisterHsemrlr20FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr20Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr20FieldProcidMask) >> RegisterHsemrlr20FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr20Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr20FieldProcidMask)|(uint32(value)<<RegisterHsemrlr20FieldProcidShift))
}

const (
	RegisterHsemrlr20FieldCoreidShift = 8
	RegisterHsemrlr20FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr20Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr20FieldCoreidMask) >> RegisterHsemrlr20FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr20Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr20FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr20FieldCoreidShift))
}

const (
	RegisterHsemrlr20FieldLockShift = 31
	RegisterHsemrlr20FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr20Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr20FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr20Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr20FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr20FieldLockMask)
	}
}

// RegisterHsemrlr21Type HSEM Read lock register
type RegisterHsemrlr21Type uint32

func (r *RegisterHsemrlr21Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr21Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr21Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr21Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr21Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr21FieldProcidShift = 0
	RegisterHsemrlr21FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr21Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr21FieldProcidMask) >> RegisterHsemrlr21FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr21Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr21FieldProcidMask)|(uint32(value)<<RegisterHsemrlr21FieldProcidShift))
}

const (
	RegisterHsemrlr21FieldCoreidShift = 8
	RegisterHsemrlr21FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr21Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr21FieldCoreidMask) >> RegisterHsemrlr21FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr21Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr21FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr21FieldCoreidShift))
}

const (
	RegisterHsemrlr21FieldLockShift = 31
	RegisterHsemrlr21FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr21Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr21FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr21Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr21FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr21FieldLockMask)
	}
}

// RegisterHsemrlr22Type HSEM Read lock register
type RegisterHsemrlr22Type uint32

func (r *RegisterHsemrlr22Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr22Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr22Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr22Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr22Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr22FieldProcidShift = 0
	RegisterHsemrlr22FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr22Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr22FieldProcidMask) >> RegisterHsemrlr22FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr22Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr22FieldProcidMask)|(uint32(value)<<RegisterHsemrlr22FieldProcidShift))
}

const (
	RegisterHsemrlr22FieldCoreidShift = 8
	RegisterHsemrlr22FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr22Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr22FieldCoreidMask) >> RegisterHsemrlr22FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr22Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr22FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr22FieldCoreidShift))
}

const (
	RegisterHsemrlr22FieldLockShift = 31
	RegisterHsemrlr22FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr22Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr22FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr22Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr22FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr22FieldLockMask)
	}
}

// RegisterHsemrlr23Type HSEM Read lock register
type RegisterHsemrlr23Type uint32

func (r *RegisterHsemrlr23Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr23Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr23Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr23Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr23Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr23FieldProcidShift = 0
	RegisterHsemrlr23FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr23Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr23FieldProcidMask) >> RegisterHsemrlr23FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr23Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr23FieldProcidMask)|(uint32(value)<<RegisterHsemrlr23FieldProcidShift))
}

const (
	RegisterHsemrlr23FieldCoreidShift = 8
	RegisterHsemrlr23FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr23Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr23FieldCoreidMask) >> RegisterHsemrlr23FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr23Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr23FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr23FieldCoreidShift))
}

const (
	RegisterHsemrlr23FieldLockShift = 31
	RegisterHsemrlr23FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr23Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr23FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr23Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr23FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr23FieldLockMask)
	}
}

// RegisterHsemrlr24Type HSEM Read lock register
type RegisterHsemrlr24Type uint32

func (r *RegisterHsemrlr24Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr24Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr24Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr24Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr24Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr24FieldProcidShift = 0
	RegisterHsemrlr24FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr24Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr24FieldProcidMask) >> RegisterHsemrlr24FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr24Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr24FieldProcidMask)|(uint32(value)<<RegisterHsemrlr24FieldProcidShift))
}

const (
	RegisterHsemrlr24FieldCoreidShift = 8
	RegisterHsemrlr24FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr24Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr24FieldCoreidMask) >> RegisterHsemrlr24FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr24Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr24FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr24FieldCoreidShift))
}

const (
	RegisterHsemrlr24FieldLockShift = 31
	RegisterHsemrlr24FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr24Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr24FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr24Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr24FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr24FieldLockMask)
	}
}

// RegisterHsemrlr25Type HSEM Read lock register
type RegisterHsemrlr25Type uint32

func (r *RegisterHsemrlr25Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr25Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr25Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr25Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr25Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr25FieldProcidShift = 0
	RegisterHsemrlr25FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr25Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr25FieldProcidMask) >> RegisterHsemrlr25FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr25Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr25FieldProcidMask)|(uint32(value)<<RegisterHsemrlr25FieldProcidShift))
}

const (
	RegisterHsemrlr25FieldCoreidShift = 8
	RegisterHsemrlr25FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr25Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr25FieldCoreidMask) >> RegisterHsemrlr25FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr25Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr25FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr25FieldCoreidShift))
}

const (
	RegisterHsemrlr25FieldLockShift = 31
	RegisterHsemrlr25FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr25Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr25FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr25Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr25FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr25FieldLockMask)
	}
}

// RegisterHsemrlr26Type HSEM Read lock register
type RegisterHsemrlr26Type uint32

func (r *RegisterHsemrlr26Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr26Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr26Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr26Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr26Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr26FieldProcidShift = 0
	RegisterHsemrlr26FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr26Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr26FieldProcidMask) >> RegisterHsemrlr26FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr26Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr26FieldProcidMask)|(uint32(value)<<RegisterHsemrlr26FieldProcidShift))
}

const (
	RegisterHsemrlr26FieldCoreidShift = 8
	RegisterHsemrlr26FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr26Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr26FieldCoreidMask) >> RegisterHsemrlr26FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr26Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr26FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr26FieldCoreidShift))
}

const (
	RegisterHsemrlr26FieldLockShift = 31
	RegisterHsemrlr26FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr26Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr26FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr26Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr26FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr26FieldLockMask)
	}
}

// RegisterHsemrlr27Type HSEM Read lock register
type RegisterHsemrlr27Type uint32

func (r *RegisterHsemrlr27Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr27Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr27Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr27Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr27Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr27FieldProcidShift = 0
	RegisterHsemrlr27FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr27Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr27FieldProcidMask) >> RegisterHsemrlr27FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr27Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr27FieldProcidMask)|(uint32(value)<<RegisterHsemrlr27FieldProcidShift))
}

const (
	RegisterHsemrlr27FieldCoreidShift = 8
	RegisterHsemrlr27FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr27Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr27FieldCoreidMask) >> RegisterHsemrlr27FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr27Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr27FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr27FieldCoreidShift))
}

const (
	RegisterHsemrlr27FieldLockShift = 31
	RegisterHsemrlr27FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr27Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr27FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr27Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr27FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr27FieldLockMask)
	}
}

// RegisterHsemrlr28Type HSEM Read lock register
type RegisterHsemrlr28Type uint32

func (r *RegisterHsemrlr28Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr28Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr28Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr28Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr28Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr28FieldProcidShift = 0
	RegisterHsemrlr28FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr28Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr28FieldProcidMask) >> RegisterHsemrlr28FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr28Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr28FieldProcidMask)|(uint32(value)<<RegisterHsemrlr28FieldProcidShift))
}

const (
	RegisterHsemrlr28FieldCoreidShift = 8
	RegisterHsemrlr28FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr28Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr28FieldCoreidMask) >> RegisterHsemrlr28FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr28Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr28FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr28FieldCoreidShift))
}

const (
	RegisterHsemrlr28FieldLockShift = 31
	RegisterHsemrlr28FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr28Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr28FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr28Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr28FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr28FieldLockMask)
	}
}

// RegisterHsemrlr29Type HSEM Read lock register
type RegisterHsemrlr29Type uint32

func (r *RegisterHsemrlr29Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr29Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr29Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr29Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr29Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr29FieldProcidShift = 0
	RegisterHsemrlr29FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr29Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr29FieldProcidMask) >> RegisterHsemrlr29FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr29Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr29FieldProcidMask)|(uint32(value)<<RegisterHsemrlr29FieldProcidShift))
}

const (
	RegisterHsemrlr29FieldCoreidShift = 8
	RegisterHsemrlr29FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr29Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr29FieldCoreidMask) >> RegisterHsemrlr29FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr29Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr29FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr29FieldCoreidShift))
}

const (
	RegisterHsemrlr29FieldLockShift = 31
	RegisterHsemrlr29FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr29Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr29FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr29Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr29FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr29FieldLockMask)
	}
}

// RegisterHsemrlr30Type HSEM Read lock register
type RegisterHsemrlr30Type uint32

func (r *RegisterHsemrlr30Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr30Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr30Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr30Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr30Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr30FieldProcidShift = 0
	RegisterHsemrlr30FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr30Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr30FieldProcidMask) >> RegisterHsemrlr30FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr30Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr30FieldProcidMask)|(uint32(value)<<RegisterHsemrlr30FieldProcidShift))
}

const (
	RegisterHsemrlr30FieldCoreidShift = 8
	RegisterHsemrlr30FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr30Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr30FieldCoreidMask) >> RegisterHsemrlr30FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr30Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr30FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr30FieldCoreidShift))
}

const (
	RegisterHsemrlr30FieldLockShift = 31
	RegisterHsemrlr30FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr30Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr30FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr30Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr30FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr30FieldLockMask)
	}
}

// RegisterHsemrlr31Type HSEM Read lock register
type RegisterHsemrlr31Type uint32

func (r *RegisterHsemrlr31Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemrlr31Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemrlr31Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemrlr31Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemrlr31Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemrlr31FieldProcidShift = 0
	RegisterHsemrlr31FieldProcidMask  = 0xff
)

// GetProcid Semaphore ProcessID
func (r *RegisterHsemrlr31Type) GetProcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr31FieldProcidMask) >> RegisterHsemrlr31FieldProcidShift)
}

// SetProcid Semaphore ProcessID
func (r *RegisterHsemrlr31Type) SetProcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr31FieldProcidMask)|(uint32(value)<<RegisterHsemrlr31FieldProcidShift))
}

const (
	RegisterHsemrlr31FieldCoreidShift = 8
	RegisterHsemrlr31FieldCoreidMask  = 0xf00
)

// GetCoreid Semaphore COREID
func (r *RegisterHsemrlr31Type) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr31FieldCoreidMask) >> RegisterHsemrlr31FieldCoreidShift)
}

// SetCoreid Semaphore COREID
func (r *RegisterHsemrlr31Type) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr31FieldCoreidMask)|(uint32(value)<<RegisterHsemrlr31FieldCoreidShift))
}

const (
	RegisterHsemrlr31FieldLockShift = 31
	RegisterHsemrlr31FieldLockMask  = 0x80000000
)

// GetLock Lock indication
func (r *RegisterHsemrlr31Type) GetLock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterHsemrlr31FieldLockMask) != 0
}

// SetLock Lock indication
func (r *RegisterHsemrlr31Type) SetLock(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterHsemrlr31FieldLockMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterHsemrlr31FieldLockMask)
	}
}

// RegisterHsemc1ierType HSEM Interrupt enable register
type RegisterHsemc1ierType uint32

func (r *RegisterHsemc1ierType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemc1ierType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemc1ierType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemc1ierType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemc1ierType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemc1ierFieldIseShift = 0
	RegisterHsemc1ierFieldIseMask  = 0xffffffff
)

// GetIse Interrupt semaphore x enable bit
func (r *RegisterHsemc1ierType) GetIse() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc1ierFieldIseMask) >> RegisterHsemc1ierFieldIseShift)
}

// SetIse Interrupt semaphore x enable bit
func (r *RegisterHsemc1ierType) SetIse(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc1ierFieldIseMask)|(uint32(value)<<RegisterHsemc1ierFieldIseShift))
}

// RegisterHsemc1icrType HSEM Interrupt clear register
type RegisterHsemc1icrType uint32

func (r *RegisterHsemc1icrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemc1icrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemc1icrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemc1icrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemc1icrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemc1icrFieldIscShift = 0
	RegisterHsemc1icrFieldIscMask  = 0xffffffff
)

// GetIsc Interrupt semaphore x clear bit
func (r *RegisterHsemc1icrType) GetIsc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc1icrFieldIscMask) >> RegisterHsemc1icrFieldIscShift)
}

// SetIsc Interrupt semaphore x clear bit
func (r *RegisterHsemc1icrType) SetIsc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc1icrFieldIscMask)|(uint32(value)<<RegisterHsemc1icrFieldIscShift))
}

// RegisterHsemc1isrType HSEM Interrupt status register
type RegisterHsemc1isrType uint32

func (r *RegisterHsemc1isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemc1isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemc1isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemc1isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemc1isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemc1isrFieldIsfShift = 0
	RegisterHsemc1isrFieldIsfMask  = 0xffffffff
)

// GetIsf Interrupt semaphore x status bit before enable (mask)
func (r *RegisterHsemc1isrType) GetIsf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc1isrFieldIsfMask) >> RegisterHsemc1isrFieldIsfShift)
}

// SetIsf Interrupt semaphore x status bit before enable (mask)
func (r *RegisterHsemc1isrType) SetIsf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc1isrFieldIsfMask)|(uint32(value)<<RegisterHsemc1isrFieldIsfShift))
}

// RegisterHsemc1misrType HSEM Masked interrupt status register
type RegisterHsemc1misrType uint32

func (r *RegisterHsemc1misrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemc1misrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemc1misrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemc1misrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemc1misrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemc1misrFieldMisfShift = 0
	RegisterHsemc1misrFieldMisfMask  = 0xffffffff
)

// GetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *RegisterHsemc1misrType) GetMisf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc1misrFieldMisfMask) >> RegisterHsemc1misrFieldMisfShift)
}

// SetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *RegisterHsemc1misrType) SetMisf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc1misrFieldMisfMask)|(uint32(value)<<RegisterHsemc1misrFieldMisfShift))
}

// RegisterHsemc2ierType HSEM Interrupt enable register
type RegisterHsemc2ierType uint32

func (r *RegisterHsemc2ierType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemc2ierType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemc2ierType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemc2ierType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemc2ierType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemc2ierFieldIseShift = 0
	RegisterHsemc2ierFieldIseMask  = 0xffffffff
)

// GetIse Interrupt semaphore x enable bit
func (r *RegisterHsemc2ierType) GetIse() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc2ierFieldIseMask) >> RegisterHsemc2ierFieldIseShift)
}

// SetIse Interrupt semaphore x enable bit
func (r *RegisterHsemc2ierType) SetIse(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc2ierFieldIseMask)|(uint32(value)<<RegisterHsemc2ierFieldIseShift))
}

// RegisterHsemc2icrType HSEM Interrupt clear register
type RegisterHsemc2icrType uint32

func (r *RegisterHsemc2icrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemc2icrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemc2icrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemc2icrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemc2icrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemc2icrFieldIscShift = 0
	RegisterHsemc2icrFieldIscMask  = 0xffffffff
)

// GetIsc Interrupt semaphore x clear bit
func (r *RegisterHsemc2icrType) GetIsc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc2icrFieldIscMask) >> RegisterHsemc2icrFieldIscShift)
}

// SetIsc Interrupt semaphore x clear bit
func (r *RegisterHsemc2icrType) SetIsc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc2icrFieldIscMask)|(uint32(value)<<RegisterHsemc2icrFieldIscShift))
}

// RegisterHsemc2isrType HSEM Interrupt status register
type RegisterHsemc2isrType uint32

func (r *RegisterHsemc2isrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemc2isrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemc2isrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemc2isrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemc2isrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemc2isrFieldIsfShift = 0
	RegisterHsemc2isrFieldIsfMask  = 0xffffffff
)

// GetIsf Interrupt semaphore x status bit before enable (mask)
func (r *RegisterHsemc2isrType) GetIsf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc2isrFieldIsfMask) >> RegisterHsemc2isrFieldIsfShift)
}

// SetIsf Interrupt semaphore x status bit before enable (mask)
func (r *RegisterHsemc2isrType) SetIsf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc2isrFieldIsfMask)|(uint32(value)<<RegisterHsemc2isrFieldIsfShift))
}

// RegisterHsemc2misrType HSEM Masked interrupt status register
type RegisterHsemc2misrType uint32

func (r *RegisterHsemc2misrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemc2misrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemc2misrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemc2misrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemc2misrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemc2misrFieldMisfShift = 0
	RegisterHsemc2misrFieldMisfMask  = 0xffffffff
)

// GetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *RegisterHsemc2misrType) GetMisf() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterHsemc2misrFieldMisfMask) >> RegisterHsemc2misrFieldMisfShift)
}

// SetMisf masked interrupt semaphore x status bit after enable (mask)
func (r *RegisterHsemc2misrType) SetMisf(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemc2misrFieldMisfMask)|(uint32(value)<<RegisterHsemc2misrFieldMisfShift))
}

// RegisterHsemcrType HSEM Clear register
type RegisterHsemcrType uint32

func (r *RegisterHsemcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemcrFieldCoreidShift = 8
	RegisterHsemcrFieldCoreidMask  = 0xf00
)

// GetCoreid COREID of semaphores to be cleared
func (r *RegisterHsemcrType) GetCoreid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsemcrFieldCoreidMask) >> RegisterHsemcrFieldCoreidShift)
}

// SetCoreid COREID of semaphores to be cleared
func (r *RegisterHsemcrType) SetCoreid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemcrFieldCoreidMask)|(uint32(value)<<RegisterHsemcrFieldCoreidShift))
}

const (
	RegisterHsemcrFieldKeyShift = 16
	RegisterHsemcrFieldKeyMask  = 0xffff0000
)

// GetKey Semaphore clear Key
func (r *RegisterHsemcrType) GetKey() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterHsemcrFieldKeyMask) >> RegisterHsemcrFieldKeyShift)
}

// SetKey Semaphore clear Key
func (r *RegisterHsemcrType) SetKey(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemcrFieldKeyMask)|(uint32(value)<<RegisterHsemcrFieldKeyShift))
}

// RegisterHsemkeyrType HSEM Interrupt clear register
type RegisterHsemkeyrType uint32

func (r *RegisterHsemkeyrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsemkeyrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsemkeyrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsemkeyrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsemkeyrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsemkeyrFieldKeyShift = 16
	RegisterHsemkeyrFieldKeyMask  = 0xffff0000
)

// GetKey Semaphore Clear Key
func (r *RegisterHsemkeyrType) GetKey() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterHsemkeyrFieldKeyMask) >> RegisterHsemkeyrFieldKeyShift)
}

// SetKey Semaphore Clear Key
func (r *RegisterHsemkeyrType) SetKey(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsemkeyrFieldKeyMask)|(uint32(value)<<RegisterHsemkeyrFieldKeyShift))
}
