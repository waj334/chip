//go:build arm && cortexm

package nvic

import (
	"unsafe"
	"volatile"
)

var (
	Nvic = (*_nvic)(unsafe.Pointer(uintptr(0xe000e100)))
)

type _nvic struct {
	Iser [16]RegisterIserType
	_    [64]byte
	Icer [16]RegisterIcerType
	_    [64]byte
	Ispr [16]RegisterIsprType
	_    [64]byte
	Icpr [16]RegisterIcprType
	_    [64]byte
	Iabr [16]RegisterIabrType
	_    [64]byte
	Itns [16]RegisterItnsType
	_    [64]byte
	Ipr  [120]RegisterIprType
	_    [2336]byte
	Stir RegisterStirType
}

// RegisterIserType Enables, or reads the enable state of a group of interrupts
type RegisterIserType uint32

func (r *RegisterIserType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIserType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIserType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIserType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIserType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIserFieldSetenaShift = 0
	RegisterIserFieldSetenaMask  = 0xffffffff
)

// GetSetena Interrupt set-enable bits. For SETENA[m] in NVIC_ISERn, allows interrupt 32n+m to be accessed.
func (r *RegisterIserType) GetSetena() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIserFieldSetenaMask) >> RegisterIserFieldSetenaShift)
}

// SetSetena Interrupt set-enable bits. For SETENA[m] in NVIC_ISERn, allows interrupt 32n+m to be accessed.
func (r *RegisterIserType) SetSetena(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIserFieldSetenaMask)|(uint32(value)<<RegisterIserFieldSetenaShift))
}

// RegisterIcerType Clears, or reads the enable state of, a group of registers
type RegisterIcerType uint32

func (r *RegisterIcerType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIcerType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIcerType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIcerType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIcerType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIcerFieldClrenaShift = 0
	RegisterIcerFieldClrenaMask  = 0xffffffff
)

// GetClrena Interrupt clear-enable bits. For SETENA[m] in NVIC_ICERn, allows interrupt 32n+m to be accessed.
func (r *RegisterIcerType) GetClrena() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcerFieldClrenaMask) >> RegisterIcerFieldClrenaShift)
}

// SetClrena Interrupt clear-enable bits. For SETENA[m] in NVIC_ICERn, allows interrupt 32n+m to be accessed.
func (r *RegisterIcerType) SetClrena(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcerFieldClrenaMask)|(uint32(value)<<RegisterIcerFieldClrenaShift))
}

// RegisterIsprType For a group of interrupts, changes interrupt status to pending, or shows the current pending status
type RegisterIsprType uint32

func (r *RegisterIsprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIsprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIsprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIsprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIsprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIsprFieldSetpendShift = 0
	RegisterIsprFieldSetpendMask  = 0xffffffff
)

// GetSetpend Interrupt set-pending bits. For SETPEND[m] in NVIC_ISPRn, allows interrupt 32n + m to be accessed.
func (r *RegisterIsprType) GetSetpend() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIsprFieldSetpendMask) >> RegisterIsprFieldSetpendShift)
}

// SetSetpend Interrupt set-pending bits. For SETPEND[m] in NVIC_ISPRn, allows interrupt 32n + m to be accessed.
func (r *RegisterIsprType) SetSetpend(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIsprFieldSetpendMask)|(uint32(value)<<RegisterIsprFieldSetpendShift))
}

// RegisterIcprType For a group of interrupts, clears the interrupt pending status, or shows the current pending status
type RegisterIcprType uint32

func (r *RegisterIcprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIcprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIcprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIcprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIcprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIcprFieldClrpendShift = 0
	RegisterIcprFieldClrpendMask  = 0xffffffff
)

// GetClrpend Interrupt clear-pending bits. Clears the pending state of interrupt (m+(32*n)), or shows whether the state of the interrupt is pending
func (r *RegisterIcprType) GetClrpend() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcprFieldClrpendMask) >> RegisterIcprFieldClrpendShift)
}

// SetClrpend Interrupt clear-pending bits. Clears the pending state of interrupt (m+(32*n)), or shows whether the state of the interrupt is pending
func (r *RegisterIcprType) SetClrpend(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcprFieldClrpendMask)|(uint32(value)<<RegisterIcprFieldClrpendShift))
}

// RegisterIabrType For a group of interrupts, shows whether each interrupt is active
type RegisterIabrType uint32

func (r *RegisterIabrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIabrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIabrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIabrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIabrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIabrFieldActiveShift = 0
	RegisterIabrFieldActiveMask  = 0xffffffff
)

// GetActive For ACTIVE[m] in NVIC_IABRn, indicates the active state for interrupt 32n+m.
func (r *RegisterIabrType) GetActive() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIabrFieldActiveMask) >> RegisterIabrFieldActiveShift)
}

// SetActive For ACTIVE[m] in NVIC_IABRn, indicates the active state for interrupt 32n+m.
func (r *RegisterIabrType) SetActive(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIabrFieldActiveMask)|(uint32(value)<<RegisterIabrFieldActiveShift))
}

// RegisterItnsType For a group of interrupts, specifies or reads whether each interrupt targets non-secure or secure state
type RegisterItnsType uint32

func (r *RegisterItnsType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterItnsType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterItnsType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterItnsType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterItnsType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterItnsFieldItnsShift = 0
	RegisterItnsFieldItnsMask  = 0xffffffff
)

// GetItns For ITNS[m] in NVIC_ITNSn, this field indicates and allows modification of the target Security state for interrupt 32n+m.
func (r *RegisterItnsType) GetItns() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterItnsFieldItnsMask) >> RegisterItnsFieldItnsShift)
}

// SetItns For ITNS[m] in NVIC_ITNSn, this field indicates and allows modification of the target Security state for interrupt 32n+m.
func (r *RegisterItnsType) SetItns(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterItnsFieldItnsMask)|(uint32(value)<<RegisterItnsFieldItnsShift))
}

// RegisterIprType Sets or reads interrupt priorities
type RegisterIprType uint32

func (r *RegisterIprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIprFieldPrin0Shift = 0
	RegisterIprFieldPrin0Mask  = 0xff
)

// GetPrin0 For register NVIC_IPRn, priority of interrupt number 4n
func (r *RegisterIprType) GetPrin0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIprFieldPrin0Mask) >> RegisterIprFieldPrin0Shift)
}

// SetPrin0 For register NVIC_IPRn, priority of interrupt number 4n
func (r *RegisterIprType) SetPrin0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIprFieldPrin0Mask)|(uint32(value)<<RegisterIprFieldPrin0Shift))
}

const (
	RegisterIprFieldPrin1Shift = 8
	RegisterIprFieldPrin1Mask  = 0xff00
)

// GetPrin1 For register NVIC_IPRn, priority of interrupt number 4n+1
func (r *RegisterIprType) GetPrin1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIprFieldPrin1Mask) >> RegisterIprFieldPrin1Shift)
}

// SetPrin1 For register NVIC_IPRn, priority of interrupt number 4n+1
func (r *RegisterIprType) SetPrin1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIprFieldPrin1Mask)|(uint32(value)<<RegisterIprFieldPrin1Shift))
}

const (
	RegisterIprFieldPrin2Shift = 16
	RegisterIprFieldPrin2Mask  = 0xff0000
)

// GetPrin2 For register NVIC_IPRn, priority of interrupt number 4n+2
func (r *RegisterIprType) GetPrin2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIprFieldPrin2Mask) >> RegisterIprFieldPrin2Shift)
}

// SetPrin2 For register NVIC_IPRn, priority of interrupt number 4n+2
func (r *RegisterIprType) SetPrin2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIprFieldPrin2Mask)|(uint32(value)<<RegisterIprFieldPrin2Shift))
}

const (
	RegisterIprFieldPrin3Shift = 24
	RegisterIprFieldPrin3Mask  = 0xff000000
)

// GetPrin3 For register NVIC_IPRn, priority of interrupt number 4n+3
func (r *RegisterIprType) GetPrin3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIprFieldPrin3Mask) >> RegisterIprFieldPrin3Shift)
}

// SetPrin3 For register NVIC_IPRn, priority of interrupt number 4n+3
func (r *RegisterIprType) SetPrin3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIprFieldPrin3Mask)|(uint32(value)<<RegisterIprFieldPrin3Shift))
}

// RegisterStirType Write to the STIR to generate an interrupt from software.
type RegisterStirType uint32

func (r *RegisterStirType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterStirType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterStirType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterStirType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterStirType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterStirFieldIntidShift = 0
	RegisterStirFieldIntidMask  = 0x1ff
)

// GetIntid Interrupt ID of the interrupt to trigger, in the range 0-479.
func (r *RegisterStirType) GetIntid() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterStirFieldIntidMask) >> RegisterStirFieldIntidShift)
}

// SetIntid Interrupt ID of the interrupt to trigger, in the range 0-479.
func (r *RegisterStirType) SetIntid(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterStirFieldIntidMask)|(uint32(value)<<RegisterStirFieldIntidShift))
}
