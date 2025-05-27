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
	Iser [16]registerIserType
	_    [64]byte
	Icer [16]registerIcerType
	_    [64]byte
	Ispr [16]registerIsprType
	_    [64]byte
	Icpr [16]registerIcprType
	_    [64]byte
	Iabr [16]registerIabrType
	_    [64]byte
	Itns [16]registerItnsType
	_    [64]byte
	Ipr  [120]registerIprType
	_    [2336]byte
	Stir registerStirType
}

// registerIserType Enables, or reads the enable state of a group of interrupts
type registerIserType uint32

const (
	RegisterIserFieldSetenaShift = 0
	RegisterIserFieldSetenaMask  = 0xffffffff
)

// GetSetena Interrupt set-enable bits. For SETENA[m] in NVIC_ISERn, allows interrupt 32n+m to be accessed.
func (r *registerIserType) GetSetena() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIserFieldSetenaMask) >> RegisterIserFieldSetenaShift)
}

// SetSetena Interrupt set-enable bits. For SETENA[m] in NVIC_ISERn, allows interrupt 32n+m to be accessed.
func (r *registerIserType) SetSetena(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIserFieldSetenaMask)|(uint32(value)<<RegisterIserFieldSetenaShift))
}

// registerIcerType Clears, or reads the enable state of, a group of registers
type registerIcerType uint32

const (
	RegisterIcerFieldClrenaShift = 0
	RegisterIcerFieldClrenaMask  = 0xffffffff
)

// GetClrena Interrupt clear-enable bits. For SETENA[m] in NVIC_ICERn, allows interrupt 32n+m to be accessed.
func (r *registerIcerType) GetClrena() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcerFieldClrenaMask) >> RegisterIcerFieldClrenaShift)
}

// SetClrena Interrupt clear-enable bits. For SETENA[m] in NVIC_ICERn, allows interrupt 32n+m to be accessed.
func (r *registerIcerType) SetClrena(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcerFieldClrenaMask)|(uint32(value)<<RegisterIcerFieldClrenaShift))
}

// registerIsprType For a group of interrupts, changes interrupt status to pending, or shows the current pending status
type registerIsprType uint32

const (
	RegisterIsprFieldSetpendShift = 0
	RegisterIsprFieldSetpendMask  = 0xffffffff
)

// GetSetpend Interrupt set-pending bits. For SETPEND[m] in NVIC_ISPRn, allows interrupt 32n + m to be accessed.
func (r *registerIsprType) GetSetpend() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIsprFieldSetpendMask) >> RegisterIsprFieldSetpendShift)
}

// SetSetpend Interrupt set-pending bits. For SETPEND[m] in NVIC_ISPRn, allows interrupt 32n + m to be accessed.
func (r *registerIsprType) SetSetpend(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIsprFieldSetpendMask)|(uint32(value)<<RegisterIsprFieldSetpendShift))
}

// registerIcprType For a group of interrupts, clears the interrupt pending status, or shows the current pending status
type registerIcprType uint32

const (
	RegisterIcprFieldClrpendShift = 0
	RegisterIcprFieldClrpendMask  = 0xffffffff
)

// GetClrpend Interrupt clear-pending bits. Clears the pending state of interrupt (m+(32*n)), or shows whether the state of the interrupt is pending
func (r *registerIcprType) GetClrpend() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcprFieldClrpendMask) >> RegisterIcprFieldClrpendShift)
}

// SetClrpend Interrupt clear-pending bits. Clears the pending state of interrupt (m+(32*n)), or shows whether the state of the interrupt is pending
func (r *registerIcprType) SetClrpend(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcprFieldClrpendMask)|(uint32(value)<<RegisterIcprFieldClrpendShift))
}

// registerIabrType For a group of interrupts, shows whether each interrupt is active
type registerIabrType uint32

const (
	RegisterIabrFieldActiveShift = 0
	RegisterIabrFieldActiveMask  = 0xffffffff
)

// GetActive For ACTIVE[m] in NVIC_IABRn, indicates the active state for interrupt 32n+m.
func (r *registerIabrType) GetActive() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIabrFieldActiveMask) >> RegisterIabrFieldActiveShift)
}

// SetActive For ACTIVE[m] in NVIC_IABRn, indicates the active state for interrupt 32n+m.
func (r *registerIabrType) SetActive(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIabrFieldActiveMask)|(uint32(value)<<RegisterIabrFieldActiveShift))
}

// registerItnsType For a group of interrupts, specifies or reads whether each interrupt targets non-secure or secure state
type registerItnsType uint32

const (
	RegisterItnsFieldItnsShift = 0
	RegisterItnsFieldItnsMask  = 0xffffffff
)

// GetItns For ITNS[m] in NVIC_ITNSn, this field indicates and allows modification of the target Security state for interrupt 32n+m.
func (r *registerItnsType) GetItns() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterItnsFieldItnsMask) >> RegisterItnsFieldItnsShift)
}

// SetItns For ITNS[m] in NVIC_ITNSn, this field indicates and allows modification of the target Security state for interrupt 32n+m.
func (r *registerItnsType) SetItns(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterItnsFieldItnsMask)|(uint32(value)<<RegisterItnsFieldItnsShift))
}

// registerIprType Sets or reads interrupt priorities
type registerIprType uint32

const (
	RegisterIprFieldPrin0Shift = 0
	RegisterIprFieldPrin0Mask  = 0xff
)

// GetPrin0 For register NVIC_IPRn, priority of interrupt number 4n
func (r *registerIprType) GetPrin0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIprFieldPrin0Mask) >> RegisterIprFieldPrin0Shift)
}

// SetPrin0 For register NVIC_IPRn, priority of interrupt number 4n
func (r *registerIprType) SetPrin0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIprFieldPrin0Mask)|(uint32(value)<<RegisterIprFieldPrin0Shift))
}

const (
	RegisterIprFieldPrin1Shift = 8
	RegisterIprFieldPrin1Mask  = 0xff00
)

// GetPrin1 For register NVIC_IPRn, priority of interrupt number 4n+1
func (r *registerIprType) GetPrin1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIprFieldPrin1Mask) >> RegisterIprFieldPrin1Shift)
}

// SetPrin1 For register NVIC_IPRn, priority of interrupt number 4n+1
func (r *registerIprType) SetPrin1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIprFieldPrin1Mask)|(uint32(value)<<RegisterIprFieldPrin1Shift))
}

const (
	RegisterIprFieldPrin2Shift = 16
	RegisterIprFieldPrin2Mask  = 0xff0000
)

// GetPrin2 For register NVIC_IPRn, priority of interrupt number 4n+2
func (r *registerIprType) GetPrin2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIprFieldPrin2Mask) >> RegisterIprFieldPrin2Shift)
}

// SetPrin2 For register NVIC_IPRn, priority of interrupt number 4n+2
func (r *registerIprType) SetPrin2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIprFieldPrin2Mask)|(uint32(value)<<RegisterIprFieldPrin2Shift))
}

const (
	RegisterIprFieldPrin3Shift = 24
	RegisterIprFieldPrin3Mask  = 0xff000000
)

// GetPrin3 For register NVIC_IPRn, priority of interrupt number 4n+3
func (r *registerIprType) GetPrin3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIprFieldPrin3Mask) >> RegisterIprFieldPrin3Shift)
}

// SetPrin3 For register NVIC_IPRn, priority of interrupt number 4n+3
func (r *registerIprType) SetPrin3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIprFieldPrin3Mask)|(uint32(value)<<RegisterIprFieldPrin3Shift))
}

// registerStirType Write to the STIR to generate an interrupt from software.
type registerStirType uint32

const (
	RegisterStirFieldIntidShift = 0
	RegisterStirFieldIntidMask  = 0x1ff
)

// GetIntid Interrupt ID of the interrupt to trigger, in the range 0-479.
func (r *registerStirType) GetIntid() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterStirFieldIntidMask) >> RegisterStirFieldIntidShift)
}

// SetIntid Interrupt ID of the interrupt to trigger, in the range 0-479.
func (r *registerStirType) SetIntid(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterStirFieldIntidMask)|(uint32(value)<<RegisterStirFieldIntidShift))
}
