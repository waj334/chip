package nvic

import (
	"unsafe"
	"volatile"
)

var (
	Nvic = (*_nvic)(unsafe.Pointer(uintptr(0xe000e100)))
)

type _nvic struct {
	Iser0 registerIser0Type
	Iser1 registerIser1Type
	Iser2 registerIser2Type
	_     [116]byte
	Icer0 registerIcer0Type
	Icer1 registerIcer1Type
	Icer2 registerIcer2Type
	_     [116]byte
	Ispr0 registerIspr0Type
	Ispr1 registerIspr1Type
	Ispr2 registerIspr2Type
	_     [116]byte
	Icpr0 registerIcpr0Type
	Icpr1 registerIcpr1Type
	Icpr2 registerIcpr2Type
	_     [116]byte
	Iabr0 registerIabr0Type
	Iabr1 registerIabr1Type
	Iabr2 registerIabr2Type
	_     [244]byte
	Ipr0  registerIpr0Type
	Ipr1  registerIpr1Type
	Ipr2  registerIpr2Type
	Ipr3  registerIpr3Type
	Ipr4  registerIpr4Type
	Ipr5  registerIpr5Type
	Ipr6  registerIpr6Type
	Ipr7  registerIpr7Type
	Ipr8  registerIpr8Type
	Ipr9  registerIpr9Type
	Ipr10 registerIpr10Type
	Ipr11 registerIpr11Type
	Ipr12 registerIpr12Type
	Ipr13 registerIpr13Type
	Ipr14 registerIpr14Type
	Ipr15 registerIpr15Type
	Ipr16 registerIpr16Type
	Ipr17 registerIpr17Type
	Ipr18 registerIpr18Type
	Ipr19 registerIpr19Type
	Ipr20 registerIpr20Type
	Ipr21 registerIpr21Type
	Ipr22 registerIpr22Type
	Ipr23 registerIpr23Type
	Ipr24 registerIpr24Type
	Ipr25 registerIpr25Type
	Ipr26 registerIpr26Type
	Ipr27 registerIpr27Type
	Ipr28 registerIpr28Type
	Ipr29 registerIpr29Type
	Ipr30 registerIpr30Type
	Ipr31 registerIpr31Type
	Ipr32 registerIpr32Type
	Ipr33 registerIpr33Type
	Ipr34 registerIpr34Type
	Ipr35 registerIpr35Type
	Ipr36 registerIpr36Type
	Ipr37 registerIpr37Type
	Ipr38 registerIpr38Type
	Ipr39 registerIpr39Type
	Iser3 registerIser3Type
	Iser4 registerIser4Type
	_     [120]byte
	Icer3 registerIcer3Type
	Icer4 registerIcer4Type
	_     [120]byte
	Ispr3 registerIspr3Type
	Ispr4 registerIspr4Type
	_     [172]byte
	Icpr3 registerIcpr3Type
	Icpr4 registerIcpr4Type
	_     [68]byte
	Iabr3 registerIabr3Type
	Iabr4 registerIabr4Type
}

// registerIser0Type Interrupt Set-Enable Register
type registerIser0Type uint32

const (
	RegisterIser0FieldSetenaShift = 0
	RegisterIser0FieldSetenaMask  = 0xffffffff
)

// GetSetena SETENA
func (r *registerIser0Type) GetSetena() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIser0FieldSetenaMask) >> RegisterIser0FieldSetenaShift)
}

// SetSetena SETENA
func (r *registerIser0Type) SetSetena(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIser0FieldSetenaMask)|(uint32(value)<<RegisterIser0FieldSetenaShift))
}

// registerIser1Type Interrupt Set-Enable Register
type registerIser1Type uint32

const (
	RegisterIser1FieldSetenaShift = 0
	RegisterIser1FieldSetenaMask  = 0xffffffff
)

// GetSetena SETENA
func (r *registerIser1Type) GetSetena() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIser1FieldSetenaMask) >> RegisterIser1FieldSetenaShift)
}

// SetSetena SETENA
func (r *registerIser1Type) SetSetena(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIser1FieldSetenaMask)|(uint32(value)<<RegisterIser1FieldSetenaShift))
}

// registerIser2Type Interrupt Set-Enable Register
type registerIser2Type uint32

const (
	RegisterIser2FieldSetenaShift = 0
	RegisterIser2FieldSetenaMask  = 0xffffffff
)

// GetSetena SETENA
func (r *registerIser2Type) GetSetena() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIser2FieldSetenaMask) >> RegisterIser2FieldSetenaShift)
}

// SetSetena SETENA
func (r *registerIser2Type) SetSetena(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIser2FieldSetenaMask)|(uint32(value)<<RegisterIser2FieldSetenaShift))
}

// registerIcer0Type Interrupt Clear-Enable Register
type registerIcer0Type uint32

const (
	RegisterIcer0FieldClrenaShift = 0
	RegisterIcer0FieldClrenaMask  = 0xffffffff
)

// GetClrena CLRENA
func (r *registerIcer0Type) GetClrena() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcer0FieldClrenaMask) >> RegisterIcer0FieldClrenaShift)
}

// SetClrena CLRENA
func (r *registerIcer0Type) SetClrena(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcer0FieldClrenaMask)|(uint32(value)<<RegisterIcer0FieldClrenaShift))
}

// registerIcer1Type Interrupt Clear-Enable Register
type registerIcer1Type uint32

const (
	RegisterIcer1FieldClrenaShift = 0
	RegisterIcer1FieldClrenaMask  = 0xffffffff
)

// GetClrena CLRENA
func (r *registerIcer1Type) GetClrena() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcer1FieldClrenaMask) >> RegisterIcer1FieldClrenaShift)
}

// SetClrena CLRENA
func (r *registerIcer1Type) SetClrena(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcer1FieldClrenaMask)|(uint32(value)<<RegisterIcer1FieldClrenaShift))
}

// registerIcer2Type Interrupt Clear-Enable Register
type registerIcer2Type uint32

const (
	RegisterIcer2FieldClrenaShift = 0
	RegisterIcer2FieldClrenaMask  = 0xffffffff
)

// GetClrena CLRENA
func (r *registerIcer2Type) GetClrena() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcer2FieldClrenaMask) >> RegisterIcer2FieldClrenaShift)
}

// SetClrena CLRENA
func (r *registerIcer2Type) SetClrena(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcer2FieldClrenaMask)|(uint32(value)<<RegisterIcer2FieldClrenaShift))
}

// registerIspr0Type Interrupt Set-Pending Register
type registerIspr0Type uint32

const (
	RegisterIspr0FieldSetpendShift = 0
	RegisterIspr0FieldSetpendMask  = 0xffffffff
)

// GetSetpend SETPEND
func (r *registerIspr0Type) GetSetpend() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIspr0FieldSetpendMask) >> RegisterIspr0FieldSetpendShift)
}

// SetSetpend SETPEND
func (r *registerIspr0Type) SetSetpend(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIspr0FieldSetpendMask)|(uint32(value)<<RegisterIspr0FieldSetpendShift))
}

// registerIspr1Type Interrupt Set-Pending Register
type registerIspr1Type uint32

const (
	RegisterIspr1FieldSetpendShift = 0
	RegisterIspr1FieldSetpendMask  = 0xffffffff
)

// GetSetpend SETPEND
func (r *registerIspr1Type) GetSetpend() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIspr1FieldSetpendMask) >> RegisterIspr1FieldSetpendShift)
}

// SetSetpend SETPEND
func (r *registerIspr1Type) SetSetpend(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIspr1FieldSetpendMask)|(uint32(value)<<RegisterIspr1FieldSetpendShift))
}

// registerIspr2Type Interrupt Set-Pending Register
type registerIspr2Type uint32

const (
	RegisterIspr2FieldSetpendShift = 0
	RegisterIspr2FieldSetpendMask  = 0xffffffff
)

// GetSetpend SETPEND
func (r *registerIspr2Type) GetSetpend() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIspr2FieldSetpendMask) >> RegisterIspr2FieldSetpendShift)
}

// SetSetpend SETPEND
func (r *registerIspr2Type) SetSetpend(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIspr2FieldSetpendMask)|(uint32(value)<<RegisterIspr2FieldSetpendShift))
}

// registerIcpr0Type Interrupt Clear-Pending Register
type registerIcpr0Type uint32

const (
	RegisterIcpr0FieldClrpendShift = 0
	RegisterIcpr0FieldClrpendMask  = 0xffffffff
)

// GetClrpend CLRPEND
func (r *registerIcpr0Type) GetClrpend() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcpr0FieldClrpendMask) >> RegisterIcpr0FieldClrpendShift)
}

// SetClrpend CLRPEND
func (r *registerIcpr0Type) SetClrpend(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcpr0FieldClrpendMask)|(uint32(value)<<RegisterIcpr0FieldClrpendShift))
}

// registerIcpr1Type Interrupt Clear-Pending Register
type registerIcpr1Type uint32

const (
	RegisterIcpr1FieldClrpendShift = 0
	RegisterIcpr1FieldClrpendMask  = 0xffffffff
)

// GetClrpend CLRPEND
func (r *registerIcpr1Type) GetClrpend() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcpr1FieldClrpendMask) >> RegisterIcpr1FieldClrpendShift)
}

// SetClrpend CLRPEND
func (r *registerIcpr1Type) SetClrpend(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcpr1FieldClrpendMask)|(uint32(value)<<RegisterIcpr1FieldClrpendShift))
}

// registerIcpr2Type Interrupt Clear-Pending Register
type registerIcpr2Type uint32

const (
	RegisterIcpr2FieldClrpendShift = 0
	RegisterIcpr2FieldClrpendMask  = 0xffffffff
)

// GetClrpend CLRPEND
func (r *registerIcpr2Type) GetClrpend() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIcpr2FieldClrpendMask) >> RegisterIcpr2FieldClrpendShift)
}

// SetClrpend CLRPEND
func (r *registerIcpr2Type) SetClrpend(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIcpr2FieldClrpendMask)|(uint32(value)<<RegisterIcpr2FieldClrpendShift))
}

// registerIabr0Type Interrupt Active Bit Register
type registerIabr0Type uint32

const (
	RegisterIabr0FieldActiveShift = 0
	RegisterIabr0FieldActiveMask  = 0xffffffff
)

// GetActive ACTIVE
func (r *registerIabr0Type) GetActive() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIabr0FieldActiveMask) >> RegisterIabr0FieldActiveShift)
}

// SetActive ACTIVE
func (r *registerIabr0Type) SetActive(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIabr0FieldActiveMask)|(uint32(value)<<RegisterIabr0FieldActiveShift))
}

// registerIabr1Type Interrupt Active Bit Register
type registerIabr1Type uint32

const (
	RegisterIabr1FieldActiveShift = 0
	RegisterIabr1FieldActiveMask  = 0xffffffff
)

// GetActive ACTIVE
func (r *registerIabr1Type) GetActive() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIabr1FieldActiveMask) >> RegisterIabr1FieldActiveShift)
}

// SetActive ACTIVE
func (r *registerIabr1Type) SetActive(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIabr1FieldActiveMask)|(uint32(value)<<RegisterIabr1FieldActiveShift))
}

// registerIabr2Type Interrupt Active Bit Register
type registerIabr2Type uint32

const (
	RegisterIabr2FieldActiveShift = 0
	RegisterIabr2FieldActiveMask  = 0xffffffff
)

// GetActive ACTIVE
func (r *registerIabr2Type) GetActive() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIabr2FieldActiveMask) >> RegisterIabr2FieldActiveShift)
}

// SetActive ACTIVE
func (r *registerIabr2Type) SetActive(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIabr2FieldActiveMask)|(uint32(value)<<RegisterIabr2FieldActiveShift))
}

// registerIpr0Type Interrupt Priority Register
type registerIpr0Type uint32

const (
	RegisterIpr0FieldIpr_n0Shift = 0
	RegisterIpr0FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr0Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr0FieldIpr_n0Mask) >> RegisterIpr0FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr0Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr0FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr0FieldIpr_n0Shift))
}

const (
	RegisterIpr0FieldIpr_n1Shift = 8
	RegisterIpr0FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr0Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr0FieldIpr_n1Mask) >> RegisterIpr0FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr0Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr0FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr0FieldIpr_n1Shift))
}

const (
	RegisterIpr0FieldIpr_n2Shift = 16
	RegisterIpr0FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr0Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr0FieldIpr_n2Mask) >> RegisterIpr0FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr0Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr0FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr0FieldIpr_n2Shift))
}

const (
	RegisterIpr0FieldIpr_n3Shift = 24
	RegisterIpr0FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr0Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr0FieldIpr_n3Mask) >> RegisterIpr0FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr0Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr0FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr0FieldIpr_n3Shift))
}

// registerIpr1Type Interrupt Priority Register
type registerIpr1Type uint32

const (
	RegisterIpr1FieldIpr_n0Shift = 0
	RegisterIpr1FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr1Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr1FieldIpr_n0Mask) >> RegisterIpr1FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr1Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr1FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr1FieldIpr_n0Shift))
}

const (
	RegisterIpr1FieldIpr_n1Shift = 8
	RegisterIpr1FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr1Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr1FieldIpr_n1Mask) >> RegisterIpr1FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr1Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr1FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr1FieldIpr_n1Shift))
}

const (
	RegisterIpr1FieldIpr_n2Shift = 16
	RegisterIpr1FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr1Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr1FieldIpr_n2Mask) >> RegisterIpr1FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr1Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr1FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr1FieldIpr_n2Shift))
}

const (
	RegisterIpr1FieldIpr_n3Shift = 24
	RegisterIpr1FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr1Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr1FieldIpr_n3Mask) >> RegisterIpr1FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr1Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr1FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr1FieldIpr_n3Shift))
}

// registerIpr2Type Interrupt Priority Register
type registerIpr2Type uint32

const (
	RegisterIpr2FieldIpr_n0Shift = 0
	RegisterIpr2FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr2Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr2FieldIpr_n0Mask) >> RegisterIpr2FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr2Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr2FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr2FieldIpr_n0Shift))
}

const (
	RegisterIpr2FieldIpr_n1Shift = 8
	RegisterIpr2FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr2Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr2FieldIpr_n1Mask) >> RegisterIpr2FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr2Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr2FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr2FieldIpr_n1Shift))
}

const (
	RegisterIpr2FieldIpr_n2Shift = 16
	RegisterIpr2FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr2Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr2FieldIpr_n2Mask) >> RegisterIpr2FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr2Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr2FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr2FieldIpr_n2Shift))
}

const (
	RegisterIpr2FieldIpr_n3Shift = 24
	RegisterIpr2FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr2Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr2FieldIpr_n3Mask) >> RegisterIpr2FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr2Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr2FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr2FieldIpr_n3Shift))
}

// registerIpr3Type Interrupt Priority Register
type registerIpr3Type uint32

const (
	RegisterIpr3FieldIpr_n0Shift = 0
	RegisterIpr3FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr3Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr3FieldIpr_n0Mask) >> RegisterIpr3FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr3Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr3FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr3FieldIpr_n0Shift))
}

const (
	RegisterIpr3FieldIpr_n1Shift = 8
	RegisterIpr3FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr3Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr3FieldIpr_n1Mask) >> RegisterIpr3FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr3Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr3FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr3FieldIpr_n1Shift))
}

const (
	RegisterIpr3FieldIpr_n2Shift = 16
	RegisterIpr3FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr3Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr3FieldIpr_n2Mask) >> RegisterIpr3FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr3Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr3FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr3FieldIpr_n2Shift))
}

const (
	RegisterIpr3FieldIpr_n3Shift = 24
	RegisterIpr3FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr3Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr3FieldIpr_n3Mask) >> RegisterIpr3FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr3Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr3FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr3FieldIpr_n3Shift))
}

// registerIpr4Type Interrupt Priority Register
type registerIpr4Type uint32

const (
	RegisterIpr4FieldIpr_n0Shift = 0
	RegisterIpr4FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr4Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr4FieldIpr_n0Mask) >> RegisterIpr4FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr4Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr4FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr4FieldIpr_n0Shift))
}

const (
	RegisterIpr4FieldIpr_n1Shift = 8
	RegisterIpr4FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr4Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr4FieldIpr_n1Mask) >> RegisterIpr4FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr4Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr4FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr4FieldIpr_n1Shift))
}

const (
	RegisterIpr4FieldIpr_n2Shift = 16
	RegisterIpr4FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr4Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr4FieldIpr_n2Mask) >> RegisterIpr4FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr4Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr4FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr4FieldIpr_n2Shift))
}

const (
	RegisterIpr4FieldIpr_n3Shift = 24
	RegisterIpr4FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr4Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr4FieldIpr_n3Mask) >> RegisterIpr4FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr4Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr4FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr4FieldIpr_n3Shift))
}

// registerIpr5Type Interrupt Priority Register
type registerIpr5Type uint32

const (
	RegisterIpr5FieldIpr_n0Shift = 0
	RegisterIpr5FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr5Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr5FieldIpr_n0Mask) >> RegisterIpr5FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr5Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr5FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr5FieldIpr_n0Shift))
}

const (
	RegisterIpr5FieldIpr_n1Shift = 8
	RegisterIpr5FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr5Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr5FieldIpr_n1Mask) >> RegisterIpr5FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr5Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr5FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr5FieldIpr_n1Shift))
}

const (
	RegisterIpr5FieldIpr_n2Shift = 16
	RegisterIpr5FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr5Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr5FieldIpr_n2Mask) >> RegisterIpr5FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr5Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr5FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr5FieldIpr_n2Shift))
}

const (
	RegisterIpr5FieldIpr_n3Shift = 24
	RegisterIpr5FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr5Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr5FieldIpr_n3Mask) >> RegisterIpr5FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr5Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr5FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr5FieldIpr_n3Shift))
}

// registerIpr6Type Interrupt Priority Register
type registerIpr6Type uint32

const (
	RegisterIpr6FieldIpr_n0Shift = 0
	RegisterIpr6FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr6Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr6FieldIpr_n0Mask) >> RegisterIpr6FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr6Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr6FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr6FieldIpr_n0Shift))
}

const (
	RegisterIpr6FieldIpr_n1Shift = 8
	RegisterIpr6FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr6Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr6FieldIpr_n1Mask) >> RegisterIpr6FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr6Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr6FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr6FieldIpr_n1Shift))
}

const (
	RegisterIpr6FieldIpr_n2Shift = 16
	RegisterIpr6FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr6Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr6FieldIpr_n2Mask) >> RegisterIpr6FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr6Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr6FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr6FieldIpr_n2Shift))
}

const (
	RegisterIpr6FieldIpr_n3Shift = 24
	RegisterIpr6FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr6Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr6FieldIpr_n3Mask) >> RegisterIpr6FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr6Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr6FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr6FieldIpr_n3Shift))
}

// registerIpr7Type Interrupt Priority Register
type registerIpr7Type uint32

const (
	RegisterIpr7FieldIpr_n0Shift = 0
	RegisterIpr7FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr7Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr7FieldIpr_n0Mask) >> RegisterIpr7FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr7Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr7FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr7FieldIpr_n0Shift))
}

const (
	RegisterIpr7FieldIpr_n1Shift = 8
	RegisterIpr7FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr7Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr7FieldIpr_n1Mask) >> RegisterIpr7FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr7Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr7FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr7FieldIpr_n1Shift))
}

const (
	RegisterIpr7FieldIpr_n2Shift = 16
	RegisterIpr7FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr7Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr7FieldIpr_n2Mask) >> RegisterIpr7FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr7Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr7FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr7FieldIpr_n2Shift))
}

const (
	RegisterIpr7FieldIpr_n3Shift = 24
	RegisterIpr7FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr7Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr7FieldIpr_n3Mask) >> RegisterIpr7FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr7Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr7FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr7FieldIpr_n3Shift))
}

// registerIpr8Type Interrupt Priority Register
type registerIpr8Type uint32

const (
	RegisterIpr8FieldIpr_n0Shift = 0
	RegisterIpr8FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr8Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr8FieldIpr_n0Mask) >> RegisterIpr8FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr8Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr8FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr8FieldIpr_n0Shift))
}

const (
	RegisterIpr8FieldIpr_n1Shift = 8
	RegisterIpr8FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr8Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr8FieldIpr_n1Mask) >> RegisterIpr8FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr8Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr8FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr8FieldIpr_n1Shift))
}

const (
	RegisterIpr8FieldIpr_n2Shift = 16
	RegisterIpr8FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr8Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr8FieldIpr_n2Mask) >> RegisterIpr8FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr8Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr8FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr8FieldIpr_n2Shift))
}

const (
	RegisterIpr8FieldIpr_n3Shift = 24
	RegisterIpr8FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr8Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr8FieldIpr_n3Mask) >> RegisterIpr8FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr8Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr8FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr8FieldIpr_n3Shift))
}

// registerIpr9Type Interrupt Priority Register
type registerIpr9Type uint32

const (
	RegisterIpr9FieldIpr_n0Shift = 0
	RegisterIpr9FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr9Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr9FieldIpr_n0Mask) >> RegisterIpr9FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr9Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr9FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr9FieldIpr_n0Shift))
}

const (
	RegisterIpr9FieldIpr_n1Shift = 8
	RegisterIpr9FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr9Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr9FieldIpr_n1Mask) >> RegisterIpr9FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr9Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr9FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr9FieldIpr_n1Shift))
}

const (
	RegisterIpr9FieldIpr_n2Shift = 16
	RegisterIpr9FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr9Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr9FieldIpr_n2Mask) >> RegisterIpr9FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr9Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr9FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr9FieldIpr_n2Shift))
}

const (
	RegisterIpr9FieldIpr_n3Shift = 24
	RegisterIpr9FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr9Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr9FieldIpr_n3Mask) >> RegisterIpr9FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr9Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr9FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr9FieldIpr_n3Shift))
}

// registerIpr10Type Interrupt Priority Register
type registerIpr10Type uint32

const (
	RegisterIpr10FieldIpr_n0Shift = 0
	RegisterIpr10FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr10Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr10FieldIpr_n0Mask) >> RegisterIpr10FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr10Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr10FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr10FieldIpr_n0Shift))
}

const (
	RegisterIpr10FieldIpr_n1Shift = 8
	RegisterIpr10FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr10Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr10FieldIpr_n1Mask) >> RegisterIpr10FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr10Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr10FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr10FieldIpr_n1Shift))
}

const (
	RegisterIpr10FieldIpr_n2Shift = 16
	RegisterIpr10FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr10Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr10FieldIpr_n2Mask) >> RegisterIpr10FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr10Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr10FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr10FieldIpr_n2Shift))
}

const (
	RegisterIpr10FieldIpr_n3Shift = 24
	RegisterIpr10FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr10Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr10FieldIpr_n3Mask) >> RegisterIpr10FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr10Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr10FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr10FieldIpr_n3Shift))
}

// registerIpr11Type Interrupt Priority Register
type registerIpr11Type uint32

const (
	RegisterIpr11FieldIpr_n0Shift = 0
	RegisterIpr11FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr11Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr11FieldIpr_n0Mask) >> RegisterIpr11FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr11Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr11FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr11FieldIpr_n0Shift))
}

const (
	RegisterIpr11FieldIpr_n1Shift = 8
	RegisterIpr11FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr11Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr11FieldIpr_n1Mask) >> RegisterIpr11FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr11Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr11FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr11FieldIpr_n1Shift))
}

const (
	RegisterIpr11FieldIpr_n2Shift = 16
	RegisterIpr11FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr11Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr11FieldIpr_n2Mask) >> RegisterIpr11FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr11Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr11FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr11FieldIpr_n2Shift))
}

const (
	RegisterIpr11FieldIpr_n3Shift = 24
	RegisterIpr11FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr11Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr11FieldIpr_n3Mask) >> RegisterIpr11FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr11Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr11FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr11FieldIpr_n3Shift))
}

// registerIpr12Type Interrupt Priority Register
type registerIpr12Type uint32

const (
	RegisterIpr12FieldIpr_n0Shift = 0
	RegisterIpr12FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr12Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr12FieldIpr_n0Mask) >> RegisterIpr12FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr12Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr12FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr12FieldIpr_n0Shift))
}

const (
	RegisterIpr12FieldIpr_n1Shift = 8
	RegisterIpr12FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr12Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr12FieldIpr_n1Mask) >> RegisterIpr12FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr12Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr12FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr12FieldIpr_n1Shift))
}

const (
	RegisterIpr12FieldIpr_n2Shift = 16
	RegisterIpr12FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr12Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr12FieldIpr_n2Mask) >> RegisterIpr12FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr12Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr12FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr12FieldIpr_n2Shift))
}

const (
	RegisterIpr12FieldIpr_n3Shift = 24
	RegisterIpr12FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr12Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr12FieldIpr_n3Mask) >> RegisterIpr12FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr12Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr12FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr12FieldIpr_n3Shift))
}

// registerIpr13Type Interrupt Priority Register
type registerIpr13Type uint32

const (
	RegisterIpr13FieldIpr_n0Shift = 0
	RegisterIpr13FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr13Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr13FieldIpr_n0Mask) >> RegisterIpr13FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr13Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr13FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr13FieldIpr_n0Shift))
}

const (
	RegisterIpr13FieldIpr_n1Shift = 8
	RegisterIpr13FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr13Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr13FieldIpr_n1Mask) >> RegisterIpr13FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr13Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr13FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr13FieldIpr_n1Shift))
}

const (
	RegisterIpr13FieldIpr_n2Shift = 16
	RegisterIpr13FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr13Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr13FieldIpr_n2Mask) >> RegisterIpr13FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr13Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr13FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr13FieldIpr_n2Shift))
}

const (
	RegisterIpr13FieldIpr_n3Shift = 24
	RegisterIpr13FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr13Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr13FieldIpr_n3Mask) >> RegisterIpr13FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr13Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr13FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr13FieldIpr_n3Shift))
}

// registerIpr14Type Interrupt Priority Register
type registerIpr14Type uint32

const (
	RegisterIpr14FieldIpr_n0Shift = 0
	RegisterIpr14FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr14Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr14FieldIpr_n0Mask) >> RegisterIpr14FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr14Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr14FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr14FieldIpr_n0Shift))
}

const (
	RegisterIpr14FieldIpr_n1Shift = 8
	RegisterIpr14FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr14Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr14FieldIpr_n1Mask) >> RegisterIpr14FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr14Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr14FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr14FieldIpr_n1Shift))
}

const (
	RegisterIpr14FieldIpr_n2Shift = 16
	RegisterIpr14FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr14Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr14FieldIpr_n2Mask) >> RegisterIpr14FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr14Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr14FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr14FieldIpr_n2Shift))
}

const (
	RegisterIpr14FieldIpr_n3Shift = 24
	RegisterIpr14FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr14Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr14FieldIpr_n3Mask) >> RegisterIpr14FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr14Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr14FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr14FieldIpr_n3Shift))
}

// registerIpr15Type Interrupt Priority Register
type registerIpr15Type uint32

const (
	RegisterIpr15FieldIpr_n0Shift = 0
	RegisterIpr15FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr15Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr15FieldIpr_n0Mask) >> RegisterIpr15FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr15Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr15FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr15FieldIpr_n0Shift))
}

const (
	RegisterIpr15FieldIpr_n1Shift = 8
	RegisterIpr15FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr15Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr15FieldIpr_n1Mask) >> RegisterIpr15FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr15Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr15FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr15FieldIpr_n1Shift))
}

const (
	RegisterIpr15FieldIpr_n2Shift = 16
	RegisterIpr15FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr15Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr15FieldIpr_n2Mask) >> RegisterIpr15FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr15Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr15FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr15FieldIpr_n2Shift))
}

const (
	RegisterIpr15FieldIpr_n3Shift = 24
	RegisterIpr15FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr15Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr15FieldIpr_n3Mask) >> RegisterIpr15FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr15Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr15FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr15FieldIpr_n3Shift))
}

// registerIpr16Type Interrupt Priority Register
type registerIpr16Type uint32

const (
	RegisterIpr16FieldIpr_n0Shift = 0
	RegisterIpr16FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr16Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr16FieldIpr_n0Mask) >> RegisterIpr16FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr16Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr16FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr16FieldIpr_n0Shift))
}

const (
	RegisterIpr16FieldIpr_n1Shift = 8
	RegisterIpr16FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr16Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr16FieldIpr_n1Mask) >> RegisterIpr16FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr16Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr16FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr16FieldIpr_n1Shift))
}

const (
	RegisterIpr16FieldIpr_n2Shift = 16
	RegisterIpr16FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr16Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr16FieldIpr_n2Mask) >> RegisterIpr16FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr16Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr16FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr16FieldIpr_n2Shift))
}

const (
	RegisterIpr16FieldIpr_n3Shift = 24
	RegisterIpr16FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr16Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr16FieldIpr_n3Mask) >> RegisterIpr16FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr16Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr16FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr16FieldIpr_n3Shift))
}

// registerIpr17Type Interrupt Priority Register
type registerIpr17Type uint32

const (
	RegisterIpr17FieldIpr_n0Shift = 0
	RegisterIpr17FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr17Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr17FieldIpr_n0Mask) >> RegisterIpr17FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr17Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr17FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr17FieldIpr_n0Shift))
}

const (
	RegisterIpr17FieldIpr_n1Shift = 8
	RegisterIpr17FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr17Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr17FieldIpr_n1Mask) >> RegisterIpr17FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr17Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr17FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr17FieldIpr_n1Shift))
}

const (
	RegisterIpr17FieldIpr_n2Shift = 16
	RegisterIpr17FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr17Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr17FieldIpr_n2Mask) >> RegisterIpr17FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr17Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr17FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr17FieldIpr_n2Shift))
}

const (
	RegisterIpr17FieldIpr_n3Shift = 24
	RegisterIpr17FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr17Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr17FieldIpr_n3Mask) >> RegisterIpr17FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr17Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr17FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr17FieldIpr_n3Shift))
}

// registerIpr18Type Interrupt Priority Register
type registerIpr18Type uint32

const (
	RegisterIpr18FieldIpr_n0Shift = 0
	RegisterIpr18FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr18Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr18FieldIpr_n0Mask) >> RegisterIpr18FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr18Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr18FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr18FieldIpr_n0Shift))
}

const (
	RegisterIpr18FieldIpr_n1Shift = 8
	RegisterIpr18FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr18Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr18FieldIpr_n1Mask) >> RegisterIpr18FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr18Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr18FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr18FieldIpr_n1Shift))
}

const (
	RegisterIpr18FieldIpr_n2Shift = 16
	RegisterIpr18FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr18Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr18FieldIpr_n2Mask) >> RegisterIpr18FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr18Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr18FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr18FieldIpr_n2Shift))
}

const (
	RegisterIpr18FieldIpr_n3Shift = 24
	RegisterIpr18FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr18Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr18FieldIpr_n3Mask) >> RegisterIpr18FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr18Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr18FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr18FieldIpr_n3Shift))
}

// registerIpr19Type Interrupt Priority Register
type registerIpr19Type uint32

const (
	RegisterIpr19FieldIpr_n0Shift = 0
	RegisterIpr19FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr19Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr19FieldIpr_n0Mask) >> RegisterIpr19FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr19Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr19FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr19FieldIpr_n0Shift))
}

const (
	RegisterIpr19FieldIpr_n1Shift = 8
	RegisterIpr19FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr19Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr19FieldIpr_n1Mask) >> RegisterIpr19FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr19Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr19FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr19FieldIpr_n1Shift))
}

const (
	RegisterIpr19FieldIpr_n2Shift = 16
	RegisterIpr19FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr19Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr19FieldIpr_n2Mask) >> RegisterIpr19FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr19Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr19FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr19FieldIpr_n2Shift))
}

const (
	RegisterIpr19FieldIpr_n3Shift = 24
	RegisterIpr19FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr19Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr19FieldIpr_n3Mask) >> RegisterIpr19FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr19Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr19FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr19FieldIpr_n3Shift))
}

// registerIpr20Type Interrupt Priority Register
type registerIpr20Type uint32

const (
	RegisterIpr20FieldIpr_n0Shift = 0
	RegisterIpr20FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr20Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr20FieldIpr_n0Mask) >> RegisterIpr20FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr20Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr20FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr20FieldIpr_n0Shift))
}

const (
	RegisterIpr20FieldIpr_n1Shift = 8
	RegisterIpr20FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr20Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr20FieldIpr_n1Mask) >> RegisterIpr20FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr20Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr20FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr20FieldIpr_n1Shift))
}

const (
	RegisterIpr20FieldIpr_n2Shift = 16
	RegisterIpr20FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr20Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr20FieldIpr_n2Mask) >> RegisterIpr20FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr20Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr20FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr20FieldIpr_n2Shift))
}

const (
	RegisterIpr20FieldIpr_n3Shift = 24
	RegisterIpr20FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr20Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr20FieldIpr_n3Mask) >> RegisterIpr20FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr20Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr20FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr20FieldIpr_n3Shift))
}

// registerIpr21Type Interrupt Priority Register
type registerIpr21Type uint32

const (
	RegisterIpr21FieldIpr_n0Shift = 0
	RegisterIpr21FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr21Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr21FieldIpr_n0Mask) >> RegisterIpr21FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr21Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr21FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr21FieldIpr_n0Shift))
}

const (
	RegisterIpr21FieldIpr_n1Shift = 8
	RegisterIpr21FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr21Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr21FieldIpr_n1Mask) >> RegisterIpr21FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr21Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr21FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr21FieldIpr_n1Shift))
}

const (
	RegisterIpr21FieldIpr_n2Shift = 16
	RegisterIpr21FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr21Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr21FieldIpr_n2Mask) >> RegisterIpr21FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr21Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr21FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr21FieldIpr_n2Shift))
}

const (
	RegisterIpr21FieldIpr_n3Shift = 24
	RegisterIpr21FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr21Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr21FieldIpr_n3Mask) >> RegisterIpr21FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr21Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr21FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr21FieldIpr_n3Shift))
}

// registerIpr22Type Interrupt Priority Register
type registerIpr22Type uint32

const (
	RegisterIpr22FieldIpr_n0Shift = 0
	RegisterIpr22FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr22Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr22FieldIpr_n0Mask) >> RegisterIpr22FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr22Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr22FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr22FieldIpr_n0Shift))
}

const (
	RegisterIpr22FieldIpr_n1Shift = 8
	RegisterIpr22FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr22Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr22FieldIpr_n1Mask) >> RegisterIpr22FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr22Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr22FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr22FieldIpr_n1Shift))
}

const (
	RegisterIpr22FieldIpr_n2Shift = 16
	RegisterIpr22FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr22Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr22FieldIpr_n2Mask) >> RegisterIpr22FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr22Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr22FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr22FieldIpr_n2Shift))
}

const (
	RegisterIpr22FieldIpr_n3Shift = 24
	RegisterIpr22FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr22Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr22FieldIpr_n3Mask) >> RegisterIpr22FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr22Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr22FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr22FieldIpr_n3Shift))
}

// registerIpr23Type Interrupt Priority Register
type registerIpr23Type uint32

const (
	RegisterIpr23FieldIpr_n0Shift = 0
	RegisterIpr23FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr23Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr23FieldIpr_n0Mask) >> RegisterIpr23FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr23Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr23FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr23FieldIpr_n0Shift))
}

const (
	RegisterIpr23FieldIpr_n1Shift = 8
	RegisterIpr23FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr23Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr23FieldIpr_n1Mask) >> RegisterIpr23FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr23Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr23FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr23FieldIpr_n1Shift))
}

const (
	RegisterIpr23FieldIpr_n2Shift = 16
	RegisterIpr23FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr23Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr23FieldIpr_n2Mask) >> RegisterIpr23FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr23Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr23FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr23FieldIpr_n2Shift))
}

const (
	RegisterIpr23FieldIpr_n3Shift = 24
	RegisterIpr23FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr23Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr23FieldIpr_n3Mask) >> RegisterIpr23FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr23Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr23FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr23FieldIpr_n3Shift))
}

// registerIpr24Type Interrupt Priority Register
type registerIpr24Type uint32

const (
	RegisterIpr24FieldIpr_n0Shift = 0
	RegisterIpr24FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr24Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr24FieldIpr_n0Mask) >> RegisterIpr24FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr24Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr24FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr24FieldIpr_n0Shift))
}

const (
	RegisterIpr24FieldIpr_n1Shift = 8
	RegisterIpr24FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr24Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr24FieldIpr_n1Mask) >> RegisterIpr24FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr24Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr24FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr24FieldIpr_n1Shift))
}

const (
	RegisterIpr24FieldIpr_n2Shift = 16
	RegisterIpr24FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr24Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr24FieldIpr_n2Mask) >> RegisterIpr24FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr24Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr24FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr24FieldIpr_n2Shift))
}

const (
	RegisterIpr24FieldIpr_n3Shift = 24
	RegisterIpr24FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr24Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr24FieldIpr_n3Mask) >> RegisterIpr24FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr24Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr24FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr24FieldIpr_n3Shift))
}

// registerIpr25Type Interrupt Priority Register
type registerIpr25Type uint32

const (
	RegisterIpr25FieldIpr_n0Shift = 0
	RegisterIpr25FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr25Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr25FieldIpr_n0Mask) >> RegisterIpr25FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr25Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr25FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr25FieldIpr_n0Shift))
}

const (
	RegisterIpr25FieldIpr_n1Shift = 8
	RegisterIpr25FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr25Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr25FieldIpr_n1Mask) >> RegisterIpr25FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr25Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr25FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr25FieldIpr_n1Shift))
}

const (
	RegisterIpr25FieldIpr_n2Shift = 16
	RegisterIpr25FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr25Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr25FieldIpr_n2Mask) >> RegisterIpr25FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr25Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr25FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr25FieldIpr_n2Shift))
}

const (
	RegisterIpr25FieldIpr_n3Shift = 24
	RegisterIpr25FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr25Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr25FieldIpr_n3Mask) >> RegisterIpr25FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr25Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr25FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr25FieldIpr_n3Shift))
}

// registerIpr26Type Interrupt Priority Register
type registerIpr26Type uint32

const (
	RegisterIpr26FieldIpr_n0Shift = 0
	RegisterIpr26FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr26Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr26FieldIpr_n0Mask) >> RegisterIpr26FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr26Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr26FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr26FieldIpr_n0Shift))
}

const (
	RegisterIpr26FieldIpr_n1Shift = 8
	RegisterIpr26FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr26Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr26FieldIpr_n1Mask) >> RegisterIpr26FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr26Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr26FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr26FieldIpr_n1Shift))
}

const (
	RegisterIpr26FieldIpr_n2Shift = 16
	RegisterIpr26FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr26Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr26FieldIpr_n2Mask) >> RegisterIpr26FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr26Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr26FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr26FieldIpr_n2Shift))
}

const (
	RegisterIpr26FieldIpr_n3Shift = 24
	RegisterIpr26FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr26Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr26FieldIpr_n3Mask) >> RegisterIpr26FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr26Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr26FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr26FieldIpr_n3Shift))
}

// registerIpr27Type Interrupt Priority Register
type registerIpr27Type uint32

const (
	RegisterIpr27FieldIpr_n0Shift = 0
	RegisterIpr27FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr27Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr27FieldIpr_n0Mask) >> RegisterIpr27FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr27Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr27FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr27FieldIpr_n0Shift))
}

const (
	RegisterIpr27FieldIpr_n1Shift = 8
	RegisterIpr27FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr27Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr27FieldIpr_n1Mask) >> RegisterIpr27FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr27Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr27FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr27FieldIpr_n1Shift))
}

const (
	RegisterIpr27FieldIpr_n2Shift = 16
	RegisterIpr27FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr27Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr27FieldIpr_n2Mask) >> RegisterIpr27FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr27Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr27FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr27FieldIpr_n2Shift))
}

const (
	RegisterIpr27FieldIpr_n3Shift = 24
	RegisterIpr27FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr27Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr27FieldIpr_n3Mask) >> RegisterIpr27FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr27Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr27FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr27FieldIpr_n3Shift))
}

// registerIpr28Type Interrupt Priority Register
type registerIpr28Type uint32

const (
	RegisterIpr28FieldIpr_n0Shift = 0
	RegisterIpr28FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr28Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr28FieldIpr_n0Mask) >> RegisterIpr28FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr28Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr28FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr28FieldIpr_n0Shift))
}

const (
	RegisterIpr28FieldIpr_n1Shift = 8
	RegisterIpr28FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr28Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr28FieldIpr_n1Mask) >> RegisterIpr28FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr28Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr28FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr28FieldIpr_n1Shift))
}

const (
	RegisterIpr28FieldIpr_n2Shift = 16
	RegisterIpr28FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr28Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr28FieldIpr_n2Mask) >> RegisterIpr28FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr28Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr28FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr28FieldIpr_n2Shift))
}

const (
	RegisterIpr28FieldIpr_n3Shift = 24
	RegisterIpr28FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr28Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr28FieldIpr_n3Mask) >> RegisterIpr28FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr28Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr28FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr28FieldIpr_n3Shift))
}

// registerIpr29Type Interrupt Priority Register
type registerIpr29Type uint32

const (
	RegisterIpr29FieldIpr_n0Shift = 0
	RegisterIpr29FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr29Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr29FieldIpr_n0Mask) >> RegisterIpr29FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr29Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr29FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr29FieldIpr_n0Shift))
}

const (
	RegisterIpr29FieldIpr_n1Shift = 8
	RegisterIpr29FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr29Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr29FieldIpr_n1Mask) >> RegisterIpr29FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr29Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr29FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr29FieldIpr_n1Shift))
}

const (
	RegisterIpr29FieldIpr_n2Shift = 16
	RegisterIpr29FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr29Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr29FieldIpr_n2Mask) >> RegisterIpr29FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr29Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr29FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr29FieldIpr_n2Shift))
}

const (
	RegisterIpr29FieldIpr_n3Shift = 24
	RegisterIpr29FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr29Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr29FieldIpr_n3Mask) >> RegisterIpr29FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr29Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr29FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr29FieldIpr_n3Shift))
}

// registerIpr30Type Interrupt Priority Register
type registerIpr30Type uint32

const (
	RegisterIpr30FieldIpr_n0Shift = 0
	RegisterIpr30FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr30Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr30FieldIpr_n0Mask) >> RegisterIpr30FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr30Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr30FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr30FieldIpr_n0Shift))
}

const (
	RegisterIpr30FieldIpr_n1Shift = 8
	RegisterIpr30FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr30Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr30FieldIpr_n1Mask) >> RegisterIpr30FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr30Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr30FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr30FieldIpr_n1Shift))
}

const (
	RegisterIpr30FieldIpr_n2Shift = 16
	RegisterIpr30FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr30Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr30FieldIpr_n2Mask) >> RegisterIpr30FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr30Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr30FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr30FieldIpr_n2Shift))
}

const (
	RegisterIpr30FieldIpr_n3Shift = 24
	RegisterIpr30FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr30Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr30FieldIpr_n3Mask) >> RegisterIpr30FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr30Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr30FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr30FieldIpr_n3Shift))
}

// registerIpr31Type Interrupt Priority Register
type registerIpr31Type uint32

const (
	RegisterIpr31FieldIpr_n0Shift = 0
	RegisterIpr31FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr31Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr31FieldIpr_n0Mask) >> RegisterIpr31FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr31Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr31FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr31FieldIpr_n0Shift))
}

const (
	RegisterIpr31FieldIpr_n1Shift = 8
	RegisterIpr31FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr31Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr31FieldIpr_n1Mask) >> RegisterIpr31FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr31Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr31FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr31FieldIpr_n1Shift))
}

const (
	RegisterIpr31FieldIpr_n2Shift = 16
	RegisterIpr31FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr31Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr31FieldIpr_n2Mask) >> RegisterIpr31FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr31Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr31FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr31FieldIpr_n2Shift))
}

const (
	RegisterIpr31FieldIpr_n3Shift = 24
	RegisterIpr31FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr31Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr31FieldIpr_n3Mask) >> RegisterIpr31FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr31Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr31FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr31FieldIpr_n3Shift))
}

// registerIpr32Type Interrupt Priority Register
type registerIpr32Type uint32

const (
	RegisterIpr32FieldIpr_n0Shift = 0
	RegisterIpr32FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr32Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr32FieldIpr_n0Mask) >> RegisterIpr32FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr32Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr32FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr32FieldIpr_n0Shift))
}

const (
	RegisterIpr32FieldIpr_n1Shift = 8
	RegisterIpr32FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr32Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr32FieldIpr_n1Mask) >> RegisterIpr32FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr32Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr32FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr32FieldIpr_n1Shift))
}

const (
	RegisterIpr32FieldIpr_n2Shift = 16
	RegisterIpr32FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr32Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr32FieldIpr_n2Mask) >> RegisterIpr32FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr32Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr32FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr32FieldIpr_n2Shift))
}

const (
	RegisterIpr32FieldIpr_n3Shift = 24
	RegisterIpr32FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr32Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr32FieldIpr_n3Mask) >> RegisterIpr32FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr32Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr32FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr32FieldIpr_n3Shift))
}

// registerIpr33Type Interrupt Priority Register
type registerIpr33Type uint32

const (
	RegisterIpr33FieldIpr_n0Shift = 0
	RegisterIpr33FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr33Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr33FieldIpr_n0Mask) >> RegisterIpr33FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr33Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr33FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr33FieldIpr_n0Shift))
}

const (
	RegisterIpr33FieldIpr_n1Shift = 8
	RegisterIpr33FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr33Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr33FieldIpr_n1Mask) >> RegisterIpr33FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr33Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr33FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr33FieldIpr_n1Shift))
}

const (
	RegisterIpr33FieldIpr_n2Shift = 16
	RegisterIpr33FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr33Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr33FieldIpr_n2Mask) >> RegisterIpr33FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr33Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr33FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr33FieldIpr_n2Shift))
}

const (
	RegisterIpr33FieldIpr_n3Shift = 24
	RegisterIpr33FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr33Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr33FieldIpr_n3Mask) >> RegisterIpr33FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr33Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr33FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr33FieldIpr_n3Shift))
}

// registerIpr34Type Interrupt Priority Register
type registerIpr34Type uint32

const (
	RegisterIpr34FieldIpr_n0Shift = 0
	RegisterIpr34FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr34Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr34FieldIpr_n0Mask) >> RegisterIpr34FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr34Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr34FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr34FieldIpr_n0Shift))
}

const (
	RegisterIpr34FieldIpr_n1Shift = 8
	RegisterIpr34FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr34Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr34FieldIpr_n1Mask) >> RegisterIpr34FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr34Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr34FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr34FieldIpr_n1Shift))
}

const (
	RegisterIpr34FieldIpr_n2Shift = 16
	RegisterIpr34FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr34Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr34FieldIpr_n2Mask) >> RegisterIpr34FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr34Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr34FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr34FieldIpr_n2Shift))
}

const (
	RegisterIpr34FieldIpr_n3Shift = 24
	RegisterIpr34FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr34Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr34FieldIpr_n3Mask) >> RegisterIpr34FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr34Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr34FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr34FieldIpr_n3Shift))
}

// registerIpr35Type Interrupt Priority Register
type registerIpr35Type uint32

const (
	RegisterIpr35FieldIpr_n0Shift = 0
	RegisterIpr35FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr35Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr35FieldIpr_n0Mask) >> RegisterIpr35FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr35Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr35FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr35FieldIpr_n0Shift))
}

const (
	RegisterIpr35FieldIpr_n1Shift = 8
	RegisterIpr35FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr35Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr35FieldIpr_n1Mask) >> RegisterIpr35FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr35Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr35FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr35FieldIpr_n1Shift))
}

const (
	RegisterIpr35FieldIpr_n2Shift = 16
	RegisterIpr35FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr35Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr35FieldIpr_n2Mask) >> RegisterIpr35FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr35Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr35FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr35FieldIpr_n2Shift))
}

const (
	RegisterIpr35FieldIpr_n3Shift = 24
	RegisterIpr35FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr35Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr35FieldIpr_n3Mask) >> RegisterIpr35FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr35Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr35FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr35FieldIpr_n3Shift))
}

// registerIpr36Type Interrupt Priority Register
type registerIpr36Type uint32

const (
	RegisterIpr36FieldIpr_n0Shift = 0
	RegisterIpr36FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr36Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr36FieldIpr_n0Mask) >> RegisterIpr36FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr36Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr36FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr36FieldIpr_n0Shift))
}

const (
	RegisterIpr36FieldIpr_n1Shift = 8
	RegisterIpr36FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr36Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr36FieldIpr_n1Mask) >> RegisterIpr36FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr36Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr36FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr36FieldIpr_n1Shift))
}

const (
	RegisterIpr36FieldIpr_n2Shift = 16
	RegisterIpr36FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr36Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr36FieldIpr_n2Mask) >> RegisterIpr36FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr36Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr36FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr36FieldIpr_n2Shift))
}

const (
	RegisterIpr36FieldIpr_n3Shift = 24
	RegisterIpr36FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr36Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr36FieldIpr_n3Mask) >> RegisterIpr36FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr36Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr36FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr36FieldIpr_n3Shift))
}

// registerIpr37Type Interrupt Priority Register
type registerIpr37Type uint32

const (
	RegisterIpr37FieldIpr_n0Shift = 0
	RegisterIpr37FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr37Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr37FieldIpr_n0Mask) >> RegisterIpr37FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr37Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr37FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr37FieldIpr_n0Shift))
}

const (
	RegisterIpr37FieldIpr_n1Shift = 8
	RegisterIpr37FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr37Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr37FieldIpr_n1Mask) >> RegisterIpr37FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr37Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr37FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr37FieldIpr_n1Shift))
}

const (
	RegisterIpr37FieldIpr_n2Shift = 16
	RegisterIpr37FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr37Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr37FieldIpr_n2Mask) >> RegisterIpr37FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr37Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr37FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr37FieldIpr_n2Shift))
}

const (
	RegisterIpr37FieldIpr_n3Shift = 24
	RegisterIpr37FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr37Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr37FieldIpr_n3Mask) >> RegisterIpr37FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr37Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr37FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr37FieldIpr_n3Shift))
}

// registerIpr38Type Interrupt Priority Register
type registerIpr38Type uint32

const (
	RegisterIpr38FieldIpr_n0Shift = 0
	RegisterIpr38FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr38Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr38FieldIpr_n0Mask) >> RegisterIpr38FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr38Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr38FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr38FieldIpr_n0Shift))
}

const (
	RegisterIpr38FieldIpr_n1Shift = 8
	RegisterIpr38FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr38Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr38FieldIpr_n1Mask) >> RegisterIpr38FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr38Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr38FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr38FieldIpr_n1Shift))
}

const (
	RegisterIpr38FieldIpr_n2Shift = 16
	RegisterIpr38FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr38Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr38FieldIpr_n2Mask) >> RegisterIpr38FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr38Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr38FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr38FieldIpr_n2Shift))
}

const (
	RegisterIpr38FieldIpr_n3Shift = 24
	RegisterIpr38FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr38Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr38FieldIpr_n3Mask) >> RegisterIpr38FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr38Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr38FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr38FieldIpr_n3Shift))
}

// registerIpr39Type Interrupt Priority Register
type registerIpr39Type uint32

const (
	RegisterIpr39FieldIpr_n0Shift = 0
	RegisterIpr39FieldIpr_n0Mask  = 0xff
)

// GetIpr_n0 IPR_N0
func (r *registerIpr39Type) GetIpr_n0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr39FieldIpr_n0Mask) >> RegisterIpr39FieldIpr_n0Shift)
}

// SetIpr_n0 IPR_N0
func (r *registerIpr39Type) SetIpr_n0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr39FieldIpr_n0Mask)|(uint32(value)<<RegisterIpr39FieldIpr_n0Shift))
}

const (
	RegisterIpr39FieldIpr_n1Shift = 8
	RegisterIpr39FieldIpr_n1Mask  = 0xff00
)

// GetIpr_n1 IPR_N1
func (r *registerIpr39Type) GetIpr_n1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr39FieldIpr_n1Mask) >> RegisterIpr39FieldIpr_n1Shift)
}

// SetIpr_n1 IPR_N1
func (r *registerIpr39Type) SetIpr_n1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr39FieldIpr_n1Mask)|(uint32(value)<<RegisterIpr39FieldIpr_n1Shift))
}

const (
	RegisterIpr39FieldIpr_n2Shift = 16
	RegisterIpr39FieldIpr_n2Mask  = 0xff0000
)

// GetIpr_n2 IPR_N2
func (r *registerIpr39Type) GetIpr_n2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr39FieldIpr_n2Mask) >> RegisterIpr39FieldIpr_n2Shift)
}

// SetIpr_n2 IPR_N2
func (r *registerIpr39Type) SetIpr_n2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr39FieldIpr_n2Mask)|(uint32(value)<<RegisterIpr39FieldIpr_n2Shift))
}

const (
	RegisterIpr39FieldIpr_n3Shift = 24
	RegisterIpr39FieldIpr_n3Mask  = 0xff000000
)

// GetIpr_n3 IPR_N3
func (r *registerIpr39Type) GetIpr_n3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIpr39FieldIpr_n3Mask) >> RegisterIpr39FieldIpr_n3Shift)
}

// SetIpr_n3 IPR_N3
func (r *registerIpr39Type) SetIpr_n3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIpr39FieldIpr_n3Mask)|(uint32(value)<<RegisterIpr39FieldIpr_n3Shift))
}

// registerIser3Type Interrupt Set-Enable Register
type registerIser3Type uint32

// registerIser4Type Interrupt Set-Enable Register
type registerIser4Type uint32

// registerIcer3Type Interrupt Clear-Enable Register
type registerIcer3Type uint32

// registerIcer4Type Interrupt Clear-Enable Register
type registerIcer4Type uint32

// registerIspr3Type Interrupt Set-Pending Register
type registerIspr3Type uint32

// registerIspr4Type Interrupt Set-Pending Register
type registerIspr4Type uint32

// registerIcpr3Type Interrupt Clear-Pending Register
type registerIcpr3Type uint32

// registerIcpr4Type Interrupt Clear-Pending Register
type registerIcpr4Type uint32

// registerIabr3Type Interrupt Active Bit Register
type registerIabr3Type uint32

// registerIabr4Type Interrupt Active Bit Register
type registerIabr4Type uint32
