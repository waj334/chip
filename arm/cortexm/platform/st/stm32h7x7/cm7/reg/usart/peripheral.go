//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package usart

import (
	"unsafe"
	"volatile"
)

var (
	Uart4  = (*_usart)(unsafe.Pointer(uintptr(0x40004c00)))
	Uart5  = (*_usart)(unsafe.Pointer(uintptr(0x40005000)))
	Uart7  = (*_usart)(unsafe.Pointer(uintptr(0x40007800)))
	Uart8  = (*_usart)(unsafe.Pointer(uintptr(0x40007c00)))
	Usart1 = (*_usart)(unsafe.Pointer(uintptr(0x40011000)))
	Usart2 = (*_usart)(unsafe.Pointer(uintptr(0x40004400)))
	Usart3 = (*_usart)(unsafe.Pointer(uintptr(0x40004800)))
	Usart6 = (*_usart)(unsafe.Pointer(uintptr(0x40011400)))

	Instances = [8]*_usart{
		Usart1,
		Usart2,
		Usart3,
		Uart4,
		Uart5,
		Usart6,
		Uart7,
		Uart8,
	}
)

type _usart struct {
	Cr1   RegisterCr1Type
	Cr2   RegisterCr2Type
	Cr3   RegisterCr3Type
	Brr   RegisterBrrType
	Gtpr  RegisterGtprType
	Rtor  RegisterRtorType
	Rqr   RegisterRqrType
	Isr   RegisterIsrType
	Icr   RegisterIcrType
	Rdr   RegisterRdrType
	Tdr   RegisterTdrType
	Presc RegisterPrescType
}

// RegisterCr1Type Control register 1
type RegisterCr1Type uint32

func (r *RegisterCr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCr1FieldUeShift = 0
	RegisterCr1FieldUeMask  = 0x1
)

// GetUe USART enable
func (r *RegisterCr1Type) GetUe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldUeMask) != 0
}

// SetUe USART enable
func (r *RegisterCr1Type) SetUe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldUeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldUeMask)
	}
}

const (
	RegisterCr1FieldUesmShift = 1
	RegisterCr1FieldUesmMask  = 0x2
)

// GetUesm USART enable in Stop mode
func (r *RegisterCr1Type) GetUesm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldUesmMask) != 0
}

// SetUesm USART enable in Stop mode
func (r *RegisterCr1Type) SetUesm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldUesmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldUesmMask)
	}
}

const (
	RegisterCr1FieldReShift = 2
	RegisterCr1FieldReMask  = 0x4
)

// GetRe Receiver enable
func (r *RegisterCr1Type) GetRe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldReMask) != 0
}

// SetRe Receiver enable
func (r *RegisterCr1Type) SetRe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldReMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldReMask)
	}
}

const (
	RegisterCr1FieldTeShift = 3
	RegisterCr1FieldTeMask  = 0x8
)

// GetTe Transmitter enable
func (r *RegisterCr1Type) GetTe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTeMask) != 0
}

// SetTe Transmitter enable
func (r *RegisterCr1Type) SetTe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTeMask)
	}
}

const (
	RegisterCr1FieldIdleieShift = 4
	RegisterCr1FieldIdleieMask  = 0x10
)

// GetIdleie IDLE interrupt enable
func (r *RegisterCr1Type) GetIdleie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldIdleieMask) != 0
}

// SetIdleie IDLE interrupt enable
func (r *RegisterCr1Type) SetIdleie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldIdleieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldIdleieMask)
	}
}

const (
	RegisterCr1FieldRxneieShift = 5
	RegisterCr1FieldRxneieMask  = 0x20
)

// GetRxneie RXNE interrupt enable
func (r *RegisterCr1Type) GetRxneie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRxneieMask) != 0
}

// SetRxneie RXNE interrupt enable
func (r *RegisterCr1Type) SetRxneie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldRxneieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldRxneieMask)
	}
}

const (
	RegisterCr1FieldTcieShift = 6
	RegisterCr1FieldTcieMask  = 0x40
)

// GetTcie Transmission complete interrupt enable
func (r *RegisterCr1Type) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTcieMask) != 0
}

// SetTcie Transmission complete interrupt enable
func (r *RegisterCr1Type) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTcieMask)
	}
}

const (
	RegisterCr1FieldTxeieShift = 7
	RegisterCr1FieldTxeieMask  = 0x80
)

// GetTxeie interrupt enable
func (r *RegisterCr1Type) GetTxeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTxeieMask) != 0
}

// SetTxeie interrupt enable
func (r *RegisterCr1Type) SetTxeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTxeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTxeieMask)
	}
}

const (
	RegisterCr1FieldPeieShift = 8
	RegisterCr1FieldPeieMask  = 0x100
)

// GetPeie PE interrupt enable
func (r *RegisterCr1Type) GetPeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPeieMask) != 0
}

// SetPeie PE interrupt enable
func (r *RegisterCr1Type) SetPeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldPeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldPeieMask)
	}
}

const (
	RegisterCr1FieldPsShift = 9
	RegisterCr1FieldPsMask  = 0x200
)

// GetPs Parity selection
func (r *RegisterCr1Type) GetPs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPsMask) != 0
}

// SetPs Parity selection
func (r *RegisterCr1Type) SetPs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldPsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldPsMask)
	}
}

const (
	RegisterCr1FieldPceShift = 10
	RegisterCr1FieldPceMask  = 0x400
)

// GetPce Parity control enable
func (r *RegisterCr1Type) GetPce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPceMask) != 0
}

// SetPce Parity control enable
func (r *RegisterCr1Type) SetPce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldPceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldPceMask)
	}
}

const (
	RegisterCr1FieldWakeShift = 11
	RegisterCr1FieldWakeMask  = 0x800
)

// GetWake Receiver wakeup method
func (r *RegisterCr1Type) GetWake() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldWakeMask) != 0
}

// SetWake Receiver wakeup method
func (r *RegisterCr1Type) SetWake(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldWakeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldWakeMask)
	}
}

const (
	RegisterCr1FieldM0Shift = 12
	RegisterCr1FieldM0Mask  = 0x1000
)

// GetM0 Word length
func (r *RegisterCr1Type) GetM0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldM0Mask) != 0
}

// SetM0 Word length
func (r *RegisterCr1Type) SetM0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldM0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldM0Mask)
	}
}

const (
	RegisterCr1FieldMmeShift = 13
	RegisterCr1FieldMmeMask  = 0x2000
)

// GetMme Mute mode enable
func (r *RegisterCr1Type) GetMme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldMmeMask) != 0
}

// SetMme Mute mode enable
func (r *RegisterCr1Type) SetMme(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldMmeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldMmeMask)
	}
}

const (
	RegisterCr1FieldCmieShift = 14
	RegisterCr1FieldCmieMask  = 0x4000
)

// GetCmie Character match interrupt enable
func (r *RegisterCr1Type) GetCmie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCmieMask) != 0
}

// SetCmie Character match interrupt enable
func (r *RegisterCr1Type) SetCmie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldCmieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCmieMask)
	}
}

const (
	RegisterCr1FieldOver8Shift = 15
	RegisterCr1FieldOver8Mask  = 0x8000
)

// GetOver8 Oversampling mode
func (r *RegisterCr1Type) GetOver8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldOver8Mask) != 0
}

// SetOver8 Oversampling mode
func (r *RegisterCr1Type) SetOver8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldOver8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldOver8Mask)
	}
}

const (
	RegisterCr1FieldDedt0Shift = 16
	RegisterCr1FieldDedt0Mask  = 0x10000
)

// GetDedt0 DEDT0
func (r *RegisterCr1Type) GetDedt0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDedt0Mask) != 0
}

// SetDedt0 DEDT0
func (r *RegisterCr1Type) SetDedt0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDedt0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDedt0Mask)
	}
}

const (
	RegisterCr1FieldDedt1Shift = 17
	RegisterCr1FieldDedt1Mask  = 0x20000
)

// GetDedt1 DEDT1
func (r *RegisterCr1Type) GetDedt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDedt1Mask) != 0
}

// SetDedt1 DEDT1
func (r *RegisterCr1Type) SetDedt1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDedt1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDedt1Mask)
	}
}

const (
	RegisterCr1FieldDedt2Shift = 18
	RegisterCr1FieldDedt2Mask  = 0x40000
)

// GetDedt2 DEDT2
func (r *RegisterCr1Type) GetDedt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDedt2Mask) != 0
}

// SetDedt2 DEDT2
func (r *RegisterCr1Type) SetDedt2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDedt2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDedt2Mask)
	}
}

const (
	RegisterCr1FieldDedt3Shift = 19
	RegisterCr1FieldDedt3Mask  = 0x80000
)

// GetDedt3 DEDT3
func (r *RegisterCr1Type) GetDedt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDedt3Mask) != 0
}

// SetDedt3 DEDT3
func (r *RegisterCr1Type) SetDedt3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDedt3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDedt3Mask)
	}
}

const (
	RegisterCr1FieldDedt4Shift = 20
	RegisterCr1FieldDedt4Mask  = 0x100000
)

// GetDedt4 Driver Enable de-assertion time
func (r *RegisterCr1Type) GetDedt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDedt4Mask) != 0
}

// SetDedt4 Driver Enable de-assertion time
func (r *RegisterCr1Type) SetDedt4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDedt4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDedt4Mask)
	}
}

const (
	RegisterCr1FieldDeat0Shift = 21
	RegisterCr1FieldDeat0Mask  = 0x200000
)

// GetDeat0 DEAT0
func (r *RegisterCr1Type) GetDeat0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDeat0Mask) != 0
}

// SetDeat0 DEAT0
func (r *RegisterCr1Type) SetDeat0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDeat0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDeat0Mask)
	}
}

const (
	RegisterCr1FieldDeat1Shift = 22
	RegisterCr1FieldDeat1Mask  = 0x400000
)

// GetDeat1 DEAT1
func (r *RegisterCr1Type) GetDeat1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDeat1Mask) != 0
}

// SetDeat1 DEAT1
func (r *RegisterCr1Type) SetDeat1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDeat1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDeat1Mask)
	}
}

const (
	RegisterCr1FieldDeat2Shift = 23
	RegisterCr1FieldDeat2Mask  = 0x800000
)

// GetDeat2 DEAT2
func (r *RegisterCr1Type) GetDeat2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDeat2Mask) != 0
}

// SetDeat2 DEAT2
func (r *RegisterCr1Type) SetDeat2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDeat2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDeat2Mask)
	}
}

const (
	RegisterCr1FieldDeat3Shift = 24
	RegisterCr1FieldDeat3Mask  = 0x1000000
)

// GetDeat3 DEAT3
func (r *RegisterCr1Type) GetDeat3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDeat3Mask) != 0
}

// SetDeat3 DEAT3
func (r *RegisterCr1Type) SetDeat3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDeat3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDeat3Mask)
	}
}

const (
	RegisterCr1FieldDeat4Shift = 25
	RegisterCr1FieldDeat4Mask  = 0x2000000
)

// GetDeat4 Driver Enable assertion time
func (r *RegisterCr1Type) GetDeat4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDeat4Mask) != 0
}

// SetDeat4 Driver Enable assertion time
func (r *RegisterCr1Type) SetDeat4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldDeat4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldDeat4Mask)
	}
}

const (
	RegisterCr1FieldRtoieShift = 26
	RegisterCr1FieldRtoieMask  = 0x4000000
)

// GetRtoie Receiver timeout interrupt enable
func (r *RegisterCr1Type) GetRtoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRtoieMask) != 0
}

// SetRtoie Receiver timeout interrupt enable
func (r *RegisterCr1Type) SetRtoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldRtoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldRtoieMask)
	}
}

const (
	RegisterCr1FieldEobieShift = 27
	RegisterCr1FieldEobieMask  = 0x8000000
)

// GetEobie End of Block interrupt enable
func (r *RegisterCr1Type) GetEobie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldEobieMask) != 0
}

// SetEobie End of Block interrupt enable
func (r *RegisterCr1Type) SetEobie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldEobieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldEobieMask)
	}
}

const (
	RegisterCr1FieldM1Shift = 28
	RegisterCr1FieldM1Mask  = 0x10000000
)

// GetM1 Word length
func (r *RegisterCr1Type) GetM1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldM1Mask) != 0
}

// SetM1 Word length
func (r *RegisterCr1Type) SetM1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldM1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldM1Mask)
	}
}

const (
	RegisterCr1FieldFifoenShift = 29
	RegisterCr1FieldFifoenMask  = 0x20000000
)

// GetFifoen FIFO mode enable
func (r *RegisterCr1Type) GetFifoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldFifoenMask) != 0
}

// SetFifoen FIFO mode enable
func (r *RegisterCr1Type) SetFifoen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldFifoenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldFifoenMask)
	}
}

const (
	RegisterCr1FieldTxfeieShift = 30
	RegisterCr1FieldTxfeieMask  = 0x40000000
)

// GetTxfeie TXFIFO empty interrupt enable
func (r *RegisterCr1Type) GetTxfeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTxfeieMask) != 0
}

// SetTxfeie TXFIFO empty interrupt enable
func (r *RegisterCr1Type) SetTxfeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTxfeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTxfeieMask)
	}
}

const (
	RegisterCr1FieldRxffieShift = 31
	RegisterCr1FieldRxffieMask  = 0x80000000
)

// GetRxffie RXFIFO Full interrupt enable
func (r *RegisterCr1Type) GetRxffie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRxffieMask) != 0
}

// SetRxffie RXFIFO Full interrupt enable
func (r *RegisterCr1Type) SetRxffie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldRxffieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldRxffieMask)
	}
}

// RegisterCr2Type Control register 2
type RegisterCr2Type uint32

func (r *RegisterCr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCr2FieldSlvenShift = 0
	RegisterCr2FieldSlvenMask  = 0x1
)

// GetSlven Synchronous Slave mode enable
func (r *RegisterCr2Type) GetSlven() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSlvenMask) != 0
}

// SetSlven Synchronous Slave mode enable
func (r *RegisterCr2Type) SetSlven(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSlvenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSlvenMask)
	}
}

const (
	RegisterCr2FieldDisnssShift = 3
	RegisterCr2FieldDisnssMask  = 0x8
)

// GetDisnss When the DSI_NSS bit is set, the NSS pin input is ignored
func (r *RegisterCr2Type) GetDisnss() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldDisnssMask) != 0
}

// SetDisnss When the DSI_NSS bit is set, the NSS pin input is ignored
func (r *RegisterCr2Type) SetDisnss(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldDisnssMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldDisnssMask)
	}
}

const (
	RegisterCr2FieldAddm7Shift = 4
	RegisterCr2FieldAddm7Mask  = 0x10
)

// GetAddm7 7-bit Address Detection/4-bit Address Detection
func (r *RegisterCr2Type) GetAddm7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAddm7Mask) != 0
}

// SetAddm7 7-bit Address Detection/4-bit Address Detection
func (r *RegisterCr2Type) SetAddm7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldAddm7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldAddm7Mask)
	}
}

const (
	RegisterCr2FieldLbdlShift = 5
	RegisterCr2FieldLbdlMask  = 0x20
)

// GetLbdl LIN break detection length
func (r *RegisterCr2Type) GetLbdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldLbdlMask) != 0
}

// SetLbdl LIN break detection length
func (r *RegisterCr2Type) SetLbdl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldLbdlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldLbdlMask)
	}
}

const (
	RegisterCr2FieldLbdieShift = 6
	RegisterCr2FieldLbdieMask  = 0x40
)

// GetLbdie LIN break detection interrupt enable
func (r *RegisterCr2Type) GetLbdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldLbdieMask) != 0
}

// SetLbdie LIN break detection interrupt enable
func (r *RegisterCr2Type) SetLbdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldLbdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldLbdieMask)
	}
}

const (
	RegisterCr2FieldLbclShift = 8
	RegisterCr2FieldLbclMask  = 0x100
)

// GetLbcl Last bit clock pulse
func (r *RegisterCr2Type) GetLbcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldLbclMask) != 0
}

// SetLbcl Last bit clock pulse
func (r *RegisterCr2Type) SetLbcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldLbclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldLbclMask)
	}
}

const (
	RegisterCr2FieldCphaShift = 9
	RegisterCr2FieldCphaMask  = 0x200
)

// GetCpha Clock phase
func (r *RegisterCr2Type) GetCpha() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCphaMask) != 0
}

// SetCpha Clock phase
func (r *RegisterCr2Type) SetCpha(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldCphaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldCphaMask)
	}
}

const (
	RegisterCr2FieldCpolShift = 10
	RegisterCr2FieldCpolMask  = 0x400
)

// GetCpol Clock polarity
func (r *RegisterCr2Type) GetCpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCpolMask) != 0
}

// SetCpol Clock polarity
func (r *RegisterCr2Type) SetCpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldCpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldCpolMask)
	}
}

const (
	RegisterCr2FieldClkenShift = 11
	RegisterCr2FieldClkenMask  = 0x800
)

// GetClken Clock enable
func (r *RegisterCr2Type) GetClken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldClkenMask) != 0
}

// SetClken Clock enable
func (r *RegisterCr2Type) SetClken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldClkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldClkenMask)
	}
}

const (
	RegisterCr2FieldStopShift = 12
	RegisterCr2FieldStopMask  = 0x3000
)

// GetStop STOP bits
func (r *RegisterCr2Type) GetStop() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldStopMask) >> RegisterCr2FieldStopShift)
}

// SetStop STOP bits
func (r *RegisterCr2Type) SetStop(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldStopMask)|(uint32(value)<<RegisterCr2FieldStopShift))
}

const (
	RegisterCr2FieldLinenShift = 14
	RegisterCr2FieldLinenMask  = 0x4000
)

// GetLinen LIN mode enable
func (r *RegisterCr2Type) GetLinen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldLinenMask) != 0
}

// SetLinen LIN mode enable
func (r *RegisterCr2Type) SetLinen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldLinenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldLinenMask)
	}
}

const (
	RegisterCr2FieldSwapShift = 15
	RegisterCr2FieldSwapMask  = 0x8000
)

// GetSwap Swap TX/RX pins
func (r *RegisterCr2Type) GetSwap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSwapMask) != 0
}

// SetSwap Swap TX/RX pins
func (r *RegisterCr2Type) SetSwap(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSwapMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSwapMask)
	}
}

const (
	RegisterCr2FieldRxinvShift = 16
	RegisterCr2FieldRxinvMask  = 0x10000
)

// GetRxinv RX pin active level inversion
func (r *RegisterCr2Type) GetRxinv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldRxinvMask) != 0
}

// SetRxinv RX pin active level inversion
func (r *RegisterCr2Type) SetRxinv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldRxinvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldRxinvMask)
	}
}

const (
	RegisterCr2FieldTxinvShift = 17
	RegisterCr2FieldTxinvMask  = 0x20000
)

// GetTxinv TX pin active level inversion
func (r *RegisterCr2Type) GetTxinv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTxinvMask) != 0
}

// SetTxinv TX pin active level inversion
func (r *RegisterCr2Type) SetTxinv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTxinvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTxinvMask)
	}
}

const (
	RegisterCr2FieldTainvShift = 18
	RegisterCr2FieldTainvMask  = 0x40000
)

// GetTainv Binary data inversion
func (r *RegisterCr2Type) GetTainv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTainvMask) != 0
}

// SetTainv Binary data inversion
func (r *RegisterCr2Type) SetTainv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldTainvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTainvMask)
	}
}

const (
	RegisterCr2FieldMsbfirstShift = 19
	RegisterCr2FieldMsbfirstMask  = 0x80000
)

// GetMsbfirst Most significant bit first
func (r *RegisterCr2Type) GetMsbfirst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldMsbfirstMask) != 0
}

// SetMsbfirst Most significant bit first
func (r *RegisterCr2Type) SetMsbfirst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldMsbfirstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldMsbfirstMask)
	}
}

const (
	RegisterCr2FieldAbrenShift = 20
	RegisterCr2FieldAbrenMask  = 0x100000
)

// GetAbren Auto baud rate enable
func (r *RegisterCr2Type) GetAbren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAbrenMask) != 0
}

// SetAbren Auto baud rate enable
func (r *RegisterCr2Type) SetAbren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldAbrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldAbrenMask)
	}
}

const (
	RegisterCr2FieldAbrmod0Shift = 21
	RegisterCr2FieldAbrmod0Mask  = 0x200000
)

// GetAbrmod0 ABRMOD0
func (r *RegisterCr2Type) GetAbrmod0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAbrmod0Mask) != 0
}

// SetAbrmod0 ABRMOD0
func (r *RegisterCr2Type) SetAbrmod0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldAbrmod0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldAbrmod0Mask)
	}
}

const (
	RegisterCr2FieldAbrmod1Shift = 22
	RegisterCr2FieldAbrmod1Mask  = 0x400000
)

// GetAbrmod1 Auto baud rate mode
func (r *RegisterCr2Type) GetAbrmod1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAbrmod1Mask) != 0
}

// SetAbrmod1 Auto baud rate mode
func (r *RegisterCr2Type) SetAbrmod1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldAbrmod1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldAbrmod1Mask)
	}
}

const (
	RegisterCr2FieldRtoenShift = 23
	RegisterCr2FieldRtoenMask  = 0x800000
)

// GetRtoen Receiver timeout enable
func (r *RegisterCr2Type) GetRtoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldRtoenMask) != 0
}

// SetRtoen Receiver timeout enable
func (r *RegisterCr2Type) SetRtoen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldRtoenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldRtoenMask)
	}
}

const (
	RegisterCr2FieldAdd03Shift = 24
	RegisterCr2FieldAdd03Mask  = 0xf000000
)

// GetAdd03 Address of the USART node
func (r *RegisterCr2Type) GetAdd03() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAdd03Mask) >> RegisterCr2FieldAdd03Shift)
}

// SetAdd03 Address of the USART node
func (r *RegisterCr2Type) SetAdd03(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldAdd03Mask)|(uint32(value)<<RegisterCr2FieldAdd03Shift))
}

const (
	RegisterCr2FieldAdd47Shift = 28
	RegisterCr2FieldAdd47Mask  = 0xf0000000
)

// GetAdd47 Address of the USART node
func (r *RegisterCr2Type) GetAdd47() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAdd47Mask) >> RegisterCr2FieldAdd47Shift)
}

// SetAdd47 Address of the USART node
func (r *RegisterCr2Type) SetAdd47(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldAdd47Mask)|(uint32(value)<<RegisterCr2FieldAdd47Shift))
}

// RegisterCr3Type Control register 3
type RegisterCr3Type uint32

func (r *RegisterCr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCr3FieldEieShift = 0
	RegisterCr3FieldEieMask  = 0x1
)

// GetEie Error interrupt enable
func (r *RegisterCr3Type) GetEie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldEieMask) != 0
}

// SetEie Error interrupt enable
func (r *RegisterCr3Type) SetEie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldEieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldEieMask)
	}
}

const (
	RegisterCr3FieldIrenShift = 1
	RegisterCr3FieldIrenMask  = 0x2
)

// GetIren Ir mode enable
func (r *RegisterCr3Type) GetIren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldIrenMask) != 0
}

// SetIren Ir mode enable
func (r *RegisterCr3Type) SetIren(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldIrenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldIrenMask)
	}
}

const (
	RegisterCr3FieldIrlpShift = 2
	RegisterCr3FieldIrlpMask  = 0x4
)

// GetIrlp Ir low-power
func (r *RegisterCr3Type) GetIrlp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldIrlpMask) != 0
}

// SetIrlp Ir low-power
func (r *RegisterCr3Type) SetIrlp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldIrlpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldIrlpMask)
	}
}

const (
	RegisterCr3FieldHdselShift = 3
	RegisterCr3FieldHdselMask  = 0x8
)

// GetHdsel Half-duplex selection
func (r *RegisterCr3Type) GetHdsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldHdselMask) != 0
}

// SetHdsel Half-duplex selection
func (r *RegisterCr3Type) SetHdsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldHdselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldHdselMask)
	}
}

const (
	RegisterCr3FieldNackShift = 4
	RegisterCr3FieldNackMask  = 0x10
)

// GetNack Smartcard NACK enable
func (r *RegisterCr3Type) GetNack() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldNackMask) != 0
}

// SetNack Smartcard NACK enable
func (r *RegisterCr3Type) SetNack(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldNackMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldNackMask)
	}
}

const (
	RegisterCr3FieldScenShift = 5
	RegisterCr3FieldScenMask  = 0x20
)

// GetScen Smartcard mode enable
func (r *RegisterCr3Type) GetScen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldScenMask) != 0
}

// SetScen Smartcard mode enable
func (r *RegisterCr3Type) SetScen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldScenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldScenMask)
	}
}

const (
	RegisterCr3FieldDmarShift = 6
	RegisterCr3FieldDmarMask  = 0x40
)

// GetDmar DMA enable receiver
func (r *RegisterCr3Type) GetDmar() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldDmarMask) != 0
}

// SetDmar DMA enable receiver
func (r *RegisterCr3Type) SetDmar(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldDmarMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldDmarMask)
	}
}

const (
	RegisterCr3FieldDmatShift = 7
	RegisterCr3FieldDmatMask  = 0x80
)

// GetDmat DMA enable transmitter
func (r *RegisterCr3Type) GetDmat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldDmatMask) != 0
}

// SetDmat DMA enable transmitter
func (r *RegisterCr3Type) SetDmat(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldDmatMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldDmatMask)
	}
}

const (
	RegisterCr3FieldRtseShift = 8
	RegisterCr3FieldRtseMask  = 0x100
)

// GetRtse RTS enable
func (r *RegisterCr3Type) GetRtse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldRtseMask) != 0
}

// SetRtse RTS enable
func (r *RegisterCr3Type) SetRtse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldRtseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldRtseMask)
	}
}

const (
	RegisterCr3FieldCtseShift = 9
	RegisterCr3FieldCtseMask  = 0x200
)

// GetCtse CTS enable
func (r *RegisterCr3Type) GetCtse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldCtseMask) != 0
}

// SetCtse CTS enable
func (r *RegisterCr3Type) SetCtse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldCtseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldCtseMask)
	}
}

const (
	RegisterCr3FieldCtsieShift = 10
	RegisterCr3FieldCtsieMask  = 0x400
)

// GetCtsie CTS interrupt enable
func (r *RegisterCr3Type) GetCtsie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldCtsieMask) != 0
}

// SetCtsie CTS interrupt enable
func (r *RegisterCr3Type) SetCtsie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldCtsieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldCtsieMask)
	}
}

const (
	RegisterCr3FieldOnebitShift = 11
	RegisterCr3FieldOnebitMask  = 0x800
)

// GetOnebit One sample bit method enable
func (r *RegisterCr3Type) GetOnebit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldOnebitMask) != 0
}

// SetOnebit One sample bit method enable
func (r *RegisterCr3Type) SetOnebit(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldOnebitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldOnebitMask)
	}
}

const (
	RegisterCr3FieldOvrdisShift = 12
	RegisterCr3FieldOvrdisMask  = 0x1000
)

// GetOvrdis Overrun Disable
func (r *RegisterCr3Type) GetOvrdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldOvrdisMask) != 0
}

// SetOvrdis Overrun Disable
func (r *RegisterCr3Type) SetOvrdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldOvrdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldOvrdisMask)
	}
}

const (
	RegisterCr3FieldDdreShift = 13
	RegisterCr3FieldDdreMask  = 0x2000
)

// GetDdre DMA Disable on Reception Error
func (r *RegisterCr3Type) GetDdre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldDdreMask) != 0
}

// SetDdre DMA Disable on Reception Error
func (r *RegisterCr3Type) SetDdre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldDdreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldDdreMask)
	}
}

const (
	RegisterCr3FieldDemShift = 14
	RegisterCr3FieldDemMask  = 0x4000
)

// GetDem Driver enable mode
func (r *RegisterCr3Type) GetDem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldDemMask) != 0
}

// SetDem Driver enable mode
func (r *RegisterCr3Type) SetDem(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldDemMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldDemMask)
	}
}

const (
	RegisterCr3FieldDepShift = 15
	RegisterCr3FieldDepMask  = 0x8000
)

// GetDep Driver enable polarity selection
func (r *RegisterCr3Type) GetDep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldDepMask) != 0
}

// SetDep Driver enable polarity selection
func (r *RegisterCr3Type) SetDep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldDepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldDepMask)
	}
}

const (
	RegisterCr3FieldScarcntShift = 17
	RegisterCr3FieldScarcntMask  = 0xe0000
)

// GetScarcnt Smartcard auto-retry count
func (r *RegisterCr3Type) GetScarcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldScarcntMask) >> RegisterCr3FieldScarcntShift)
}

// SetScarcnt Smartcard auto-retry count
func (r *RegisterCr3Type) SetScarcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldScarcntMask)|(uint32(value)<<RegisterCr3FieldScarcntShift))
}

const (
	RegisterCr3FieldWusShift = 20
	RegisterCr3FieldWusMask  = 0x300000
)

// GetWus Wakeup from Stop mode interrupt flag selection
func (r *RegisterCr3Type) GetWus() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldWusMask) >> RegisterCr3FieldWusShift)
}

// SetWus Wakeup from Stop mode interrupt flag selection
func (r *RegisterCr3Type) SetWus(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldWusMask)|(uint32(value)<<RegisterCr3FieldWusShift))
}

const (
	RegisterCr3FieldWufieShift = 22
	RegisterCr3FieldWufieMask  = 0x400000
)

// GetWufie Wakeup from Stop mode interrupt enable
func (r *RegisterCr3Type) GetWufie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldWufieMask) != 0
}

// SetWufie Wakeup from Stop mode interrupt enable
func (r *RegisterCr3Type) SetWufie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldWufieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldWufieMask)
	}
}

const (
	RegisterCr3FieldTxftieShift = 23
	RegisterCr3FieldTxftieMask  = 0x800000
)

// GetTxftie TXFIFO threshold interrupt enable
func (r *RegisterCr3Type) GetTxftie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldTxftieMask) != 0
}

// SetTxftie TXFIFO threshold interrupt enable
func (r *RegisterCr3Type) SetTxftie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldTxftieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldTxftieMask)
	}
}

const (
	RegisterCr3FieldTcbgtieShift = 24
	RegisterCr3FieldTcbgtieMask  = 0x1000000
)

// GetTcbgtie Transmission Complete before guard time, interrupt enable
func (r *RegisterCr3Type) GetTcbgtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldTcbgtieMask) != 0
}

// SetTcbgtie Transmission Complete before guard time, interrupt enable
func (r *RegisterCr3Type) SetTcbgtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldTcbgtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldTcbgtieMask)
	}
}

const (
	RegisterCr3FieldRxftcfgShift = 25
	RegisterCr3FieldRxftcfgMask  = 0xe000000
)

// GetRxftcfg Receive FIFO threshold configuration
func (r *RegisterCr3Type) GetRxftcfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldRxftcfgMask) >> RegisterCr3FieldRxftcfgShift)
}

// SetRxftcfg Receive FIFO threshold configuration
func (r *RegisterCr3Type) SetRxftcfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldRxftcfgMask)|(uint32(value)<<RegisterCr3FieldRxftcfgShift))
}

const (
	RegisterCr3FieldRxftieShift = 28
	RegisterCr3FieldRxftieMask  = 0x10000000
)

// GetRxftie RXFIFO threshold interrupt enable
func (r *RegisterCr3Type) GetRxftie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldRxftieMask) != 0
}

// SetRxftie RXFIFO threshold interrupt enable
func (r *RegisterCr3Type) SetRxftie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr3FieldRxftieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldRxftieMask)
	}
}

const (
	RegisterCr3FieldTxftcfgShift = 29
	RegisterCr3FieldTxftcfgMask  = 0xe0000000
)

// GetTxftcfg TXFIFO threshold configuration
func (r *RegisterCr3Type) GetTxftcfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldTxftcfgMask) >> RegisterCr3FieldTxftcfgShift)
}

// SetTxftcfg TXFIFO threshold configuration
func (r *RegisterCr3Type) SetTxftcfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldTxftcfgMask)|(uint32(value)<<RegisterCr3FieldTxftcfgShift))
}

// RegisterBrrType Baud rate register
type RegisterBrrType uint32

func (r *RegisterBrrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBrrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBrrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBrrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBrrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBrrFieldBrr03Shift = 0
	RegisterBrrFieldBrr03Mask  = 0xf
)

// GetBrr03 DIV_Fraction
func (r *RegisterBrrType) GetBrr03() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBrrFieldBrr03Mask) >> RegisterBrrFieldBrr03Shift)
}

// SetBrr03 DIV_Fraction
func (r *RegisterBrrType) SetBrr03(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBrrFieldBrr03Mask)|(uint32(value)<<RegisterBrrFieldBrr03Shift))
}

const (
	RegisterBrrFieldBrr415Shift = 4
	RegisterBrrFieldBrr415Mask  = 0xfff0
)

// GetBrr415 DIV_Mantissa
func (r *RegisterBrrType) GetBrr415() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBrrFieldBrr415Mask) >> RegisterBrrFieldBrr415Shift)
}

// SetBrr415 DIV_Mantissa
func (r *RegisterBrrType) SetBrr415(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBrrFieldBrr415Mask)|(uint32(value)<<RegisterBrrFieldBrr415Shift))
}

// RegisterGtprType Guard time and prescaler register
type RegisterGtprType uint32

func (r *RegisterGtprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterGtprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterGtprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterGtprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterGtprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterGtprFieldPscShift = 0
	RegisterGtprFieldPscMask  = 0xff
)

// GetPsc Prescaler value
func (r *RegisterGtprType) GetPsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterGtprFieldPscMask) >> RegisterGtprFieldPscShift)
}

// SetPsc Prescaler value
func (r *RegisterGtprType) SetPsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterGtprFieldPscMask)|(uint32(value)<<RegisterGtprFieldPscShift))
}

const (
	RegisterGtprFieldGtShift = 8
	RegisterGtprFieldGtMask  = 0xff00
)

// GetGt Guard time value
func (r *RegisterGtprType) GetGt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterGtprFieldGtMask) >> RegisterGtprFieldGtShift)
}

// SetGt Guard time value
func (r *RegisterGtprType) SetGt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterGtprFieldGtMask)|(uint32(value)<<RegisterGtprFieldGtShift))
}

// RegisterRtorType Receiver timeout register
type RegisterRtorType uint32

func (r *RegisterRtorType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRtorType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRtorType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRtorType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRtorType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRtorFieldRtoShift = 0
	RegisterRtorFieldRtoMask  = 0xffffff
)

// GetRto Receiver timeout value
func (r *RegisterRtorType) GetRto() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtorFieldRtoMask) >> RegisterRtorFieldRtoShift)
}

// SetRto Receiver timeout value
func (r *RegisterRtorType) SetRto(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtorFieldRtoMask)|(uint32(value)<<RegisterRtorFieldRtoShift))
}

const (
	RegisterRtorFieldBlenShift = 24
	RegisterRtorFieldBlenMask  = 0xff000000
)

// GetBlen Block Length
func (r *RegisterRtorType) GetBlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtorFieldBlenMask) >> RegisterRtorFieldBlenShift)
}

// SetBlen Block Length
func (r *RegisterRtorType) SetBlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtorFieldBlenMask)|(uint32(value)<<RegisterRtorFieldBlenShift))
}

// RegisterRqrType Request register
type RegisterRqrType uint32

func (r *RegisterRqrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRqrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRqrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRqrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRqrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRqrFieldAbrrqShift = 0
	RegisterRqrFieldAbrrqMask  = 0x1
)

// GetAbrrq Auto baud rate request
func (r *RegisterRqrType) GetAbrrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRqrFieldAbrrqMask) != 0
}

// SetAbrrq Auto baud rate request
func (r *RegisterRqrType) SetAbrrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRqrFieldAbrrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRqrFieldAbrrqMask)
	}
}

const (
	RegisterRqrFieldSbkrqShift = 1
	RegisterRqrFieldSbkrqMask  = 0x2
)

// GetSbkrq Send break request
func (r *RegisterRqrType) GetSbkrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRqrFieldSbkrqMask) != 0
}

// SetSbkrq Send break request
func (r *RegisterRqrType) SetSbkrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRqrFieldSbkrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRqrFieldSbkrqMask)
	}
}

const (
	RegisterRqrFieldMmrqShift = 2
	RegisterRqrFieldMmrqMask  = 0x4
)

// GetMmrq Mute mode request
func (r *RegisterRqrType) GetMmrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRqrFieldMmrqMask) != 0
}

// SetMmrq Mute mode request
func (r *RegisterRqrType) SetMmrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRqrFieldMmrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRqrFieldMmrqMask)
	}
}

const (
	RegisterRqrFieldRxfrqShift = 3
	RegisterRqrFieldRxfrqMask  = 0x8
)

// GetRxfrq Receive data flush request
func (r *RegisterRqrType) GetRxfrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRqrFieldRxfrqMask) != 0
}

// SetRxfrq Receive data flush request
func (r *RegisterRqrType) SetRxfrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRqrFieldRxfrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRqrFieldRxfrqMask)
	}
}

const (
	RegisterRqrFieldTxfrqShift = 4
	RegisterRqrFieldTxfrqMask  = 0x10
)

// GetTxfrq Transmit data flush request
func (r *RegisterRqrType) GetTxfrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRqrFieldTxfrqMask) != 0
}

// SetTxfrq Transmit data flush request
func (r *RegisterRqrType) SetTxfrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRqrFieldTxfrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRqrFieldTxfrqMask)
	}
}

// RegisterIsrType Interrupt & status register
type RegisterIsrType uint32

func (r *RegisterIsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIsrFieldPeShift = 0
	RegisterIsrFieldPeMask  = 0x1
)

// GetPe PE
func (r *RegisterIsrType) GetPe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldPeMask) != 0
}

// SetPe PE
func (r *RegisterIsrType) SetPe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldPeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldPeMask)
	}
}

const (
	RegisterIsrFieldFeShift = 1
	RegisterIsrFieldFeMask  = 0x2
)

// GetFe FE
func (r *RegisterIsrType) GetFe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldFeMask) != 0
}

// SetFe FE
func (r *RegisterIsrType) SetFe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldFeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldFeMask)
	}
}

const (
	RegisterIsrFieldNfShift = 2
	RegisterIsrFieldNfMask  = 0x4
)

// GetNf NF
func (r *RegisterIsrType) GetNf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldNfMask) != 0
}

// SetNf NF
func (r *RegisterIsrType) SetNf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldNfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldNfMask)
	}
}

const (
	RegisterIsrFieldOreShift = 3
	RegisterIsrFieldOreMask  = 0x8
)

// GetOre ORE
func (r *RegisterIsrType) GetOre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldOreMask) != 0
}

// SetOre ORE
func (r *RegisterIsrType) SetOre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldOreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldOreMask)
	}
}

const (
	RegisterIsrFieldIdleShift = 4
	RegisterIsrFieldIdleMask  = 0x10
)

// GetIdle IDLE
func (r *RegisterIsrType) GetIdle() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldIdleMask) != 0
}

// SetIdle IDLE
func (r *RegisterIsrType) SetIdle(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldIdleMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldIdleMask)
	}
}

const (
	RegisterIsrFieldRxneShift = 5
	RegisterIsrFieldRxneMask  = 0x20
)

// GetRxne RXNE
func (r *RegisterIsrType) GetRxne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxneMask) != 0
}

// SetRxne RXNE
func (r *RegisterIsrType) SetRxne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRxneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRxneMask)
	}
}

const (
	RegisterIsrFieldTcShift = 6
	RegisterIsrFieldTcMask  = 0x40
)

// GetTc TC
func (r *RegisterIsrType) GetTc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcMask) != 0
}

// SetTc TC
func (r *RegisterIsrType) SetTc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcMask)
	}
}

const (
	RegisterIsrFieldTxeShift = 7
	RegisterIsrFieldTxeMask  = 0x80
)

// GetTxe TXE
func (r *RegisterIsrType) GetTxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxeMask) != 0
}

// SetTxe TXE
func (r *RegisterIsrType) SetTxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxeMask)
	}
}

const (
	RegisterIsrFieldLbdfShift = 8
	RegisterIsrFieldLbdfMask  = 0x100
)

// GetLbdf LBDF
func (r *RegisterIsrType) GetLbdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldLbdfMask) != 0
}

// SetLbdf LBDF
func (r *RegisterIsrType) SetLbdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldLbdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldLbdfMask)
	}
}

const (
	RegisterIsrFieldCtsifShift = 9
	RegisterIsrFieldCtsifMask  = 0x200
)

// GetCtsif CTSIF
func (r *RegisterIsrType) GetCtsif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCtsifMask) != 0
}

// SetCtsif CTSIF
func (r *RegisterIsrType) SetCtsif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldCtsifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldCtsifMask)
	}
}

const (
	RegisterIsrFieldCtsShift = 10
	RegisterIsrFieldCtsMask  = 0x400
)

// GetCts CTS
func (r *RegisterIsrType) GetCts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCtsMask) != 0
}

// SetCts CTS
func (r *RegisterIsrType) SetCts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldCtsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldCtsMask)
	}
}

const (
	RegisterIsrFieldRtofShift = 11
	RegisterIsrFieldRtofMask  = 0x800
)

// GetRtof RTOF
func (r *RegisterIsrType) GetRtof() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRtofMask) != 0
}

// SetRtof RTOF
func (r *RegisterIsrType) SetRtof(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRtofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRtofMask)
	}
}

const (
	RegisterIsrFieldEobfShift = 12
	RegisterIsrFieldEobfMask  = 0x1000
)

// GetEobf EOBF
func (r *RegisterIsrType) GetEobf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldEobfMask) != 0
}

// SetEobf EOBF
func (r *RegisterIsrType) SetEobf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldEobfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldEobfMask)
	}
}

const (
	RegisterIsrFieldUdrShift = 13
	RegisterIsrFieldUdrMask  = 0x2000
)

// GetUdr SPI slave underrun error flag
func (r *RegisterIsrType) GetUdr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldUdrMask) != 0
}

// SetUdr SPI slave underrun error flag
func (r *RegisterIsrType) SetUdr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldUdrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldUdrMask)
	}
}

const (
	RegisterIsrFieldAbreShift = 14
	RegisterIsrFieldAbreMask  = 0x4000
)

// GetAbre ABRE
func (r *RegisterIsrType) GetAbre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldAbreMask) != 0
}

// SetAbre ABRE
func (r *RegisterIsrType) SetAbre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldAbreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldAbreMask)
	}
}

const (
	RegisterIsrFieldAbrfShift = 15
	RegisterIsrFieldAbrfMask  = 0x8000
)

// GetAbrf ABRF
func (r *RegisterIsrType) GetAbrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldAbrfMask) != 0
}

// SetAbrf ABRF
func (r *RegisterIsrType) SetAbrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldAbrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldAbrfMask)
	}
}

const (
	RegisterIsrFieldBusyShift = 16
	RegisterIsrFieldBusyMask  = 0x10000
)

// GetBusy BUSY
func (r *RegisterIsrType) GetBusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldBusyMask) != 0
}

// SetBusy BUSY
func (r *RegisterIsrType) SetBusy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldBusyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldBusyMask)
	}
}

const (
	RegisterIsrFieldCmfShift = 17
	RegisterIsrFieldCmfMask  = 0x20000
)

// GetCmf CMF
func (r *RegisterIsrType) GetCmf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCmfMask) != 0
}

// SetCmf CMF
func (r *RegisterIsrType) SetCmf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldCmfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldCmfMask)
	}
}

const (
	RegisterIsrFieldSbkfShift = 18
	RegisterIsrFieldSbkfMask  = 0x40000
)

// GetSbkf SBKF
func (r *RegisterIsrType) GetSbkf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldSbkfMask) != 0
}

// SetSbkf SBKF
func (r *RegisterIsrType) SetSbkf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldSbkfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldSbkfMask)
	}
}

const (
	RegisterIsrFieldRwuShift = 19
	RegisterIsrFieldRwuMask  = 0x80000
)

// GetRwu RWU
func (r *RegisterIsrType) GetRwu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRwuMask) != 0
}

// SetRwu RWU
func (r *RegisterIsrType) SetRwu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRwuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRwuMask)
	}
}

const (
	RegisterIsrFieldWufShift = 20
	RegisterIsrFieldWufMask  = 0x100000
)

// GetWuf WUF
func (r *RegisterIsrType) GetWuf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldWufMask) != 0
}

// SetWuf WUF
func (r *RegisterIsrType) SetWuf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldWufMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldWufMask)
	}
}

const (
	RegisterIsrFieldTeackShift = 21
	RegisterIsrFieldTeackMask  = 0x200000
)

// GetTeack TEACK
func (r *RegisterIsrType) GetTeack() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTeackMask) != 0
}

// SetTeack TEACK
func (r *RegisterIsrType) SetTeack(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTeackMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTeackMask)
	}
}

const (
	RegisterIsrFieldReackShift = 22
	RegisterIsrFieldReackMask  = 0x400000
)

// GetReack REACK
func (r *RegisterIsrType) GetReack() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldReackMask) != 0
}

// SetReack REACK
func (r *RegisterIsrType) SetReack(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldReackMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldReackMask)
	}
}

const (
	RegisterIsrFieldTxfeShift = 23
	RegisterIsrFieldTxfeMask  = 0x800000
)

// GetTxfe TXFIFO Empty
func (r *RegisterIsrType) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxfeMask) != 0
}

// SetTxfe TXFIFO Empty
func (r *RegisterIsrType) SetTxfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxfeMask)
	}
}

const (
	RegisterIsrFieldRxffShift = 24
	RegisterIsrFieldRxffMask  = 0x1000000
)

// GetRxff RXFIFO Full
func (r *RegisterIsrType) GetRxff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxffMask) != 0
}

// SetRxff RXFIFO Full
func (r *RegisterIsrType) SetRxff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRxffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRxffMask)
	}
}

const (
	RegisterIsrFieldTcbgtShift = 25
	RegisterIsrFieldTcbgtMask  = 0x2000000
)

// GetTcbgt Transmission complete before guard time flag
func (r *RegisterIsrType) GetTcbgt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcbgtMask) != 0
}

// SetTcbgt Transmission complete before guard time flag
func (r *RegisterIsrType) SetTcbgt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcbgtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcbgtMask)
	}
}

const (
	RegisterIsrFieldRxftShift = 26
	RegisterIsrFieldRxftMask  = 0x4000000
)

// GetRxft RXFIFO threshold flag
func (r *RegisterIsrType) GetRxft() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxftMask) != 0
}

// SetRxft RXFIFO threshold flag
func (r *RegisterIsrType) SetRxft(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRxftMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRxftMask)
	}
}

const (
	RegisterIsrFieldTxftShift = 27
	RegisterIsrFieldTxftMask  = 0x8000000
)

// GetTxft TXFIFO threshold flag
func (r *RegisterIsrType) GetTxft() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxftMask) != 0
}

// SetTxft TXFIFO threshold flag
func (r *RegisterIsrType) SetTxft(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxftMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxftMask)
	}
}

// RegisterIcrType Interrupt flag clear register
type RegisterIcrType uint32

func (r *RegisterIcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIcrFieldPecfShift = 0
	RegisterIcrFieldPecfMask  = 0x1
)

// GetPecf Parity error clear flag
func (r *RegisterIcrType) GetPecf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldPecfMask) != 0
}

// SetPecf Parity error clear flag
func (r *RegisterIcrType) SetPecf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldPecfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldPecfMask)
	}
}

const (
	RegisterIcrFieldFecfShift = 1
	RegisterIcrFieldFecfMask  = 0x2
)

// GetFecf Framing error clear flag
func (r *RegisterIcrType) GetFecf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldFecfMask) != 0
}

// SetFecf Framing error clear flag
func (r *RegisterIcrType) SetFecf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldFecfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldFecfMask)
	}
}

const (
	RegisterIcrFieldNcfShift = 2
	RegisterIcrFieldNcfMask  = 0x4
)

// GetNcf Noise detected clear flag
func (r *RegisterIcrType) GetNcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldNcfMask) != 0
}

// SetNcf Noise detected clear flag
func (r *RegisterIcrType) SetNcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldNcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldNcfMask)
	}
}

const (
	RegisterIcrFieldOrecfShift = 3
	RegisterIcrFieldOrecfMask  = 0x8
)

// GetOrecf Overrun error clear flag
func (r *RegisterIcrType) GetOrecf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldOrecfMask) != 0
}

// SetOrecf Overrun error clear flag
func (r *RegisterIcrType) SetOrecf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldOrecfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldOrecfMask)
	}
}

const (
	RegisterIcrFieldIdlecfShift = 4
	RegisterIcrFieldIdlecfMask  = 0x10
)

// GetIdlecf Idle line detected clear flag
func (r *RegisterIcrType) GetIdlecf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldIdlecfMask) != 0
}

// SetIdlecf Idle line detected clear flag
func (r *RegisterIcrType) SetIdlecf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldIdlecfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldIdlecfMask)
	}
}

const (
	RegisterIcrFieldTxfecfShift = 5
	RegisterIcrFieldTxfecfMask  = 0x20
)

// GetTxfecf TXFIFO empty clear flag
func (r *RegisterIcrType) GetTxfecf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldTxfecfMask) != 0
}

// SetTxfecf TXFIFO empty clear flag
func (r *RegisterIcrType) SetTxfecf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldTxfecfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldTxfecfMask)
	}
}

const (
	RegisterIcrFieldTccfShift = 6
	RegisterIcrFieldTccfMask  = 0x40
)

// GetTccf Transmission complete clear flag
func (r *RegisterIcrType) GetTccf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldTccfMask) != 0
}

// SetTccf Transmission complete clear flag
func (r *RegisterIcrType) SetTccf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldTccfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldTccfMask)
	}
}

const (
	RegisterIcrFieldTcbgtcShift = 7
	RegisterIcrFieldTcbgtcMask  = 0x80
)

// GetTcbgtc Transmission complete before Guard time clear flag
func (r *RegisterIcrType) GetTcbgtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldTcbgtcMask) != 0
}

// SetTcbgtc Transmission complete before Guard time clear flag
func (r *RegisterIcrType) SetTcbgtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldTcbgtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldTcbgtcMask)
	}
}

const (
	RegisterIcrFieldLbdcfShift = 8
	RegisterIcrFieldLbdcfMask  = 0x100
)

// GetLbdcf LIN break detection clear flag
func (r *RegisterIcrType) GetLbdcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldLbdcfMask) != 0
}

// SetLbdcf LIN break detection clear flag
func (r *RegisterIcrType) SetLbdcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldLbdcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldLbdcfMask)
	}
}

const (
	RegisterIcrFieldCtscfShift = 9
	RegisterIcrFieldCtscfMask  = 0x200
)

// GetCtscf CTS clear flag
func (r *RegisterIcrType) GetCtscf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCtscfMask) != 0
}

// SetCtscf CTS clear flag
func (r *RegisterIcrType) SetCtscf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCtscfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCtscfMask)
	}
}

const (
	RegisterIcrFieldRtocfShift = 11
	RegisterIcrFieldRtocfMask  = 0x800
)

// GetRtocf Receiver timeout clear flag
func (r *RegisterIcrType) GetRtocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldRtocfMask) != 0
}

// SetRtocf Receiver timeout clear flag
func (r *RegisterIcrType) SetRtocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldRtocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldRtocfMask)
	}
}

const (
	RegisterIcrFieldEobcfShift = 12
	RegisterIcrFieldEobcfMask  = 0x1000
)

// GetEobcf End of block clear flag
func (r *RegisterIcrType) GetEobcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldEobcfMask) != 0
}

// SetEobcf End of block clear flag
func (r *RegisterIcrType) SetEobcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldEobcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldEobcfMask)
	}
}

const (
	RegisterIcrFieldUdrcfShift = 13
	RegisterIcrFieldUdrcfMask  = 0x2000
)

// GetUdrcf SPI slave underrun clear flag
func (r *RegisterIcrType) GetUdrcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldUdrcfMask) != 0
}

// SetUdrcf SPI slave underrun clear flag
func (r *RegisterIcrType) SetUdrcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldUdrcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldUdrcfMask)
	}
}

const (
	RegisterIcrFieldCmcfShift = 17
	RegisterIcrFieldCmcfMask  = 0x20000
)

// GetCmcf Character match clear flag
func (r *RegisterIcrType) GetCmcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCmcfMask) != 0
}

// SetCmcf Character match clear flag
func (r *RegisterIcrType) SetCmcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCmcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCmcfMask)
	}
}

const (
	RegisterIcrFieldWucfShift = 20
	RegisterIcrFieldWucfMask  = 0x100000
)

// GetWucf Wakeup from Stop mode clear flag
func (r *RegisterIcrType) GetWucf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldWucfMask) != 0
}

// SetWucf Wakeup from Stop mode clear flag
func (r *RegisterIcrType) SetWucf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldWucfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldWucfMask)
	}
}

// RegisterRdrType Receive data register
type RegisterRdrType uint32

func (r *RegisterRdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRdrFieldRdrShift = 0
	RegisterRdrFieldRdrMask  = 0x1ff
)

// GetRdr Receive data value
func (r *RegisterRdrType) GetRdr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRdrFieldRdrMask) >> RegisterRdrFieldRdrShift)
}

// SetRdr Receive data value
func (r *RegisterRdrType) SetRdr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRdrFieldRdrMask)|(uint32(value)<<RegisterRdrFieldRdrShift))
}

// RegisterTdrType Transmit data register
type RegisterTdrType uint32

func (r *RegisterTdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTdrFieldTdrShift = 0
	RegisterTdrFieldTdrMask  = 0x1ff
)

// GetTdr Transmit data value
func (r *RegisterTdrType) GetTdr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterTdrFieldTdrMask) >> RegisterTdrFieldTdrShift)
}

// SetTdr Transmit data value
func (r *RegisterTdrType) SetTdr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTdrFieldTdrMask)|(uint32(value)<<RegisterTdrFieldTdrShift))
}

// RegisterPrescType USART prescaler register
type RegisterPrescType uint32

func (r *RegisterPrescType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPrescType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPrescType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPrescType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPrescType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPrescFieldPrescalerShift = 0
	RegisterPrescFieldPrescalerMask  = 0xf
)

// GetPrescaler Clock prescaler
func (r *RegisterPrescType) GetPrescaler() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPrescFieldPrescalerMask) >> RegisterPrescFieldPrescalerShift)
}

// SetPrescaler Clock prescaler
func (r *RegisterPrescType) SetPrescaler(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrescFieldPrescalerMask)|(uint32(value)<<RegisterPrescFieldPrescalerShift))
}
