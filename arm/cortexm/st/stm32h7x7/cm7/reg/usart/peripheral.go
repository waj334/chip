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
	Cr1   registerCr1Type
	Cr2   registerCr2Type
	Cr3   registerCr3Type
	Brr   registerBrrType
	Gtpr  registerGtprType
	Rtor  registerRtorType
	Rqr   registerRqrType
	Isr   registerIsrType
	Icr   registerIcrType
	Rdr   registerRdrType
	Tdr   registerTdrType
	Presc registerPrescType
}

// registerCr1Type Control register 1
type registerCr1Type uint32

const (
	RegisterCr1FieldUeShift = 0
	RegisterCr1FieldUeMask  = 0x1
)

// GetUe USART enable
func (r *registerCr1Type) GetUe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldUeMask) != 0
}

// SetUe USART enable
func (r *registerCr1Type) SetUe(value bool) {
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
func (r *registerCr1Type) GetUesm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldUesmMask) != 0
}

// SetUesm USART enable in Stop mode
func (r *registerCr1Type) SetUesm(value bool) {
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
func (r *registerCr1Type) GetRe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldReMask) != 0
}

// SetRe Receiver enable
func (r *registerCr1Type) SetRe(value bool) {
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
func (r *registerCr1Type) GetTe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTeMask) != 0
}

// SetTe Transmitter enable
func (r *registerCr1Type) SetTe(value bool) {
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
func (r *registerCr1Type) GetIdleie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldIdleieMask) != 0
}

// SetIdleie IDLE interrupt enable
func (r *registerCr1Type) SetIdleie(value bool) {
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
func (r *registerCr1Type) GetRxneie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRxneieMask) != 0
}

// SetRxneie RXNE interrupt enable
func (r *registerCr1Type) SetRxneie(value bool) {
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
func (r *registerCr1Type) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTcieMask) != 0
}

// SetTcie Transmission complete interrupt enable
func (r *registerCr1Type) SetTcie(value bool) {
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
func (r *registerCr1Type) GetTxeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTxeieMask) != 0
}

// SetTxeie interrupt enable
func (r *registerCr1Type) SetTxeie(value bool) {
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
func (r *registerCr1Type) GetPeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPeieMask) != 0
}

// SetPeie PE interrupt enable
func (r *registerCr1Type) SetPeie(value bool) {
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
func (r *registerCr1Type) GetPs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPsMask) != 0
}

// SetPs Parity selection
func (r *registerCr1Type) SetPs(value bool) {
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
func (r *registerCr1Type) GetPce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldPceMask) != 0
}

// SetPce Parity control enable
func (r *registerCr1Type) SetPce(value bool) {
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
func (r *registerCr1Type) GetWake() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldWakeMask) != 0
}

// SetWake Receiver wakeup method
func (r *registerCr1Type) SetWake(value bool) {
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
func (r *registerCr1Type) GetM0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldM0Mask) != 0
}

// SetM0 Word length
func (r *registerCr1Type) SetM0(value bool) {
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
func (r *registerCr1Type) GetMme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldMmeMask) != 0
}

// SetMme Mute mode enable
func (r *registerCr1Type) SetMme(value bool) {
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
func (r *registerCr1Type) GetCmie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCmieMask) != 0
}

// SetCmie Character match interrupt enable
func (r *registerCr1Type) SetCmie(value bool) {
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
func (r *registerCr1Type) GetOver8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldOver8Mask) != 0
}

// SetOver8 Oversampling mode
func (r *registerCr1Type) SetOver8(value bool) {
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
func (r *registerCr1Type) GetDedt0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDedt0Mask) != 0
}

// SetDedt0 DEDT0
func (r *registerCr1Type) SetDedt0(value bool) {
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
func (r *registerCr1Type) GetDedt1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDedt1Mask) != 0
}

// SetDedt1 DEDT1
func (r *registerCr1Type) SetDedt1(value bool) {
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
func (r *registerCr1Type) GetDedt2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDedt2Mask) != 0
}

// SetDedt2 DEDT2
func (r *registerCr1Type) SetDedt2(value bool) {
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
func (r *registerCr1Type) GetDedt3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDedt3Mask) != 0
}

// SetDedt3 DEDT3
func (r *registerCr1Type) SetDedt3(value bool) {
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
func (r *registerCr1Type) GetDedt4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDedt4Mask) != 0
}

// SetDedt4 Driver Enable de-assertion time
func (r *registerCr1Type) SetDedt4(value bool) {
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
func (r *registerCr1Type) GetDeat0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDeat0Mask) != 0
}

// SetDeat0 DEAT0
func (r *registerCr1Type) SetDeat0(value bool) {
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
func (r *registerCr1Type) GetDeat1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDeat1Mask) != 0
}

// SetDeat1 DEAT1
func (r *registerCr1Type) SetDeat1(value bool) {
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
func (r *registerCr1Type) GetDeat2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDeat2Mask) != 0
}

// SetDeat2 DEAT2
func (r *registerCr1Type) SetDeat2(value bool) {
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
func (r *registerCr1Type) GetDeat3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDeat3Mask) != 0
}

// SetDeat3 DEAT3
func (r *registerCr1Type) SetDeat3(value bool) {
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
func (r *registerCr1Type) GetDeat4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldDeat4Mask) != 0
}

// SetDeat4 Driver Enable assertion time
func (r *registerCr1Type) SetDeat4(value bool) {
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
func (r *registerCr1Type) GetRtoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRtoieMask) != 0
}

// SetRtoie Receiver timeout interrupt enable
func (r *registerCr1Type) SetRtoie(value bool) {
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
func (r *registerCr1Type) GetEobie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldEobieMask) != 0
}

// SetEobie End of Block interrupt enable
func (r *registerCr1Type) SetEobie(value bool) {
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
func (r *registerCr1Type) GetM1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldM1Mask) != 0
}

// SetM1 Word length
func (r *registerCr1Type) SetM1(value bool) {
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
func (r *registerCr1Type) GetFifoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldFifoenMask) != 0
}

// SetFifoen FIFO mode enable
func (r *registerCr1Type) SetFifoen(value bool) {
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
func (r *registerCr1Type) GetTxfeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTxfeieMask) != 0
}

// SetTxfeie TXFIFO empty interrupt enable
func (r *registerCr1Type) SetTxfeie(value bool) {
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
func (r *registerCr1Type) GetRxffie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRxffieMask) != 0
}

// SetRxffie RXFIFO Full interrupt enable
func (r *registerCr1Type) SetRxffie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldRxffieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldRxffieMask)
	}
}

// registerCr2Type Control register 2
type registerCr2Type uint32

const (
	RegisterCr2FieldSlvenShift = 0
	RegisterCr2FieldSlvenMask  = 0x1
)

// GetSlven Synchronous Slave mode enable
func (r *registerCr2Type) GetSlven() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSlvenMask) != 0
}

// SetSlven Synchronous Slave mode enable
func (r *registerCr2Type) SetSlven(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldSlvenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldSlvenMask)
	}
}

const (
	RegisterCr2FieldDis_nssShift = 3
	RegisterCr2FieldDis_nssMask  = 0x8
)

// GetDis_nss When the DSI_NSS bit is set, the NSS pin input is ignored
func (r *registerCr2Type) GetDis_nss() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldDis_nssMask) != 0
}

// SetDis_nss When the DSI_NSS bit is set, the NSS pin input is ignored
func (r *registerCr2Type) SetDis_nss(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldDis_nssMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldDis_nssMask)
	}
}

const (
	RegisterCr2FieldAddm7Shift = 4
	RegisterCr2FieldAddm7Mask  = 0x10
)

// GetAddm7 7-bit Address Detection/4-bit Address Detection
func (r *registerCr2Type) GetAddm7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAddm7Mask) != 0
}

// SetAddm7 7-bit Address Detection/4-bit Address Detection
func (r *registerCr2Type) SetAddm7(value bool) {
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
func (r *registerCr2Type) GetLbdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldLbdlMask) != 0
}

// SetLbdl LIN break detection length
func (r *registerCr2Type) SetLbdl(value bool) {
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
func (r *registerCr2Type) GetLbdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldLbdieMask) != 0
}

// SetLbdie LIN break detection interrupt enable
func (r *registerCr2Type) SetLbdie(value bool) {
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
func (r *registerCr2Type) GetLbcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldLbclMask) != 0
}

// SetLbcl Last bit clock pulse
func (r *registerCr2Type) SetLbcl(value bool) {
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
func (r *registerCr2Type) GetCpha() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCphaMask) != 0
}

// SetCpha Clock phase
func (r *registerCr2Type) SetCpha(value bool) {
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
func (r *registerCr2Type) GetCpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldCpolMask) != 0
}

// SetCpol Clock polarity
func (r *registerCr2Type) SetCpol(value bool) {
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
func (r *registerCr2Type) GetClken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldClkenMask) != 0
}

// SetClken Clock enable
func (r *registerCr2Type) SetClken(value bool) {
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
func (r *registerCr2Type) GetStop() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldStopMask) >> RegisterCr2FieldStopShift)
}

// SetStop STOP bits
func (r *registerCr2Type) SetStop(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldStopMask)|(uint32(value)<<RegisterCr2FieldStopShift))
}

const (
	RegisterCr2FieldLinenShift = 14
	RegisterCr2FieldLinenMask  = 0x4000
)

// GetLinen LIN mode enable
func (r *registerCr2Type) GetLinen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldLinenMask) != 0
}

// SetLinen LIN mode enable
func (r *registerCr2Type) SetLinen(value bool) {
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
func (r *registerCr2Type) GetSwap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldSwapMask) != 0
}

// SetSwap Swap TX/RX pins
func (r *registerCr2Type) SetSwap(value bool) {
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
func (r *registerCr2Type) GetRxinv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldRxinvMask) != 0
}

// SetRxinv RX pin active level inversion
func (r *registerCr2Type) SetRxinv(value bool) {
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
func (r *registerCr2Type) GetTxinv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTxinvMask) != 0
}

// SetTxinv TX pin active level inversion
func (r *registerCr2Type) SetTxinv(value bool) {
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
func (r *registerCr2Type) GetTainv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTainvMask) != 0
}

// SetTainv Binary data inversion
func (r *registerCr2Type) SetTainv(value bool) {
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
func (r *registerCr2Type) GetMsbfirst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldMsbfirstMask) != 0
}

// SetMsbfirst Most significant bit first
func (r *registerCr2Type) SetMsbfirst(value bool) {
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
func (r *registerCr2Type) GetAbren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAbrenMask) != 0
}

// SetAbren Auto baud rate enable
func (r *registerCr2Type) SetAbren(value bool) {
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
func (r *registerCr2Type) GetAbrmod0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAbrmod0Mask) != 0
}

// SetAbrmod0 ABRMOD0
func (r *registerCr2Type) SetAbrmod0(value bool) {
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
func (r *registerCr2Type) GetAbrmod1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAbrmod1Mask) != 0
}

// SetAbrmod1 Auto baud rate mode
func (r *registerCr2Type) SetAbrmod1(value bool) {
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
func (r *registerCr2Type) GetRtoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldRtoenMask) != 0
}

// SetRtoen Receiver timeout enable
func (r *registerCr2Type) SetRtoen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr2FieldRtoenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldRtoenMask)
	}
}

const (
	RegisterCr2FieldAdd0_3Shift = 24
	RegisterCr2FieldAdd0_3Mask  = 0xf000000
)

// GetAdd0_3 Address of the USART node
func (r *registerCr2Type) GetAdd0_3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAdd0_3Mask) >> RegisterCr2FieldAdd0_3Shift)
}

// SetAdd0_3 Address of the USART node
func (r *registerCr2Type) SetAdd0_3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldAdd0_3Mask)|(uint32(value)<<RegisterCr2FieldAdd0_3Shift))
}

const (
	RegisterCr2FieldAdd4_7Shift = 28
	RegisterCr2FieldAdd4_7Mask  = 0xf0000000
)

// GetAdd4_7 Address of the USART node
func (r *registerCr2Type) GetAdd4_7() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldAdd4_7Mask) >> RegisterCr2FieldAdd4_7Shift)
}

// SetAdd4_7 Address of the USART node
func (r *registerCr2Type) SetAdd4_7(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldAdd4_7Mask)|(uint32(value)<<RegisterCr2FieldAdd4_7Shift))
}

// registerCr3Type Control register 3
type registerCr3Type uint32

const (
	RegisterCr3FieldEieShift = 0
	RegisterCr3FieldEieMask  = 0x1
)

// GetEie Error interrupt enable
func (r *registerCr3Type) GetEie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldEieMask) != 0
}

// SetEie Error interrupt enable
func (r *registerCr3Type) SetEie(value bool) {
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
func (r *registerCr3Type) GetIren() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldIrenMask) != 0
}

// SetIren Ir mode enable
func (r *registerCr3Type) SetIren(value bool) {
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
func (r *registerCr3Type) GetIrlp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldIrlpMask) != 0
}

// SetIrlp Ir low-power
func (r *registerCr3Type) SetIrlp(value bool) {
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
func (r *registerCr3Type) GetHdsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldHdselMask) != 0
}

// SetHdsel Half-duplex selection
func (r *registerCr3Type) SetHdsel(value bool) {
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
func (r *registerCr3Type) GetNack() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldNackMask) != 0
}

// SetNack Smartcard NACK enable
func (r *registerCr3Type) SetNack(value bool) {
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
func (r *registerCr3Type) GetScen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldScenMask) != 0
}

// SetScen Smartcard mode enable
func (r *registerCr3Type) SetScen(value bool) {
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
func (r *registerCr3Type) GetDmar() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldDmarMask) != 0
}

// SetDmar DMA enable receiver
func (r *registerCr3Type) SetDmar(value bool) {
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
func (r *registerCr3Type) GetDmat() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldDmatMask) != 0
}

// SetDmat DMA enable transmitter
func (r *registerCr3Type) SetDmat(value bool) {
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
func (r *registerCr3Type) GetRtse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldRtseMask) != 0
}

// SetRtse RTS enable
func (r *registerCr3Type) SetRtse(value bool) {
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
func (r *registerCr3Type) GetCtse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldCtseMask) != 0
}

// SetCtse CTS enable
func (r *registerCr3Type) SetCtse(value bool) {
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
func (r *registerCr3Type) GetCtsie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldCtsieMask) != 0
}

// SetCtsie CTS interrupt enable
func (r *registerCr3Type) SetCtsie(value bool) {
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
func (r *registerCr3Type) GetOnebit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldOnebitMask) != 0
}

// SetOnebit One sample bit method enable
func (r *registerCr3Type) SetOnebit(value bool) {
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
func (r *registerCr3Type) GetOvrdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldOvrdisMask) != 0
}

// SetOvrdis Overrun Disable
func (r *registerCr3Type) SetOvrdis(value bool) {
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
func (r *registerCr3Type) GetDdre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldDdreMask) != 0
}

// SetDdre DMA Disable on Reception Error
func (r *registerCr3Type) SetDdre(value bool) {
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
func (r *registerCr3Type) GetDem() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldDemMask) != 0
}

// SetDem Driver enable mode
func (r *registerCr3Type) SetDem(value bool) {
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
func (r *registerCr3Type) GetDep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldDepMask) != 0
}

// SetDep Driver enable polarity selection
func (r *registerCr3Type) SetDep(value bool) {
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
func (r *registerCr3Type) GetScarcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldScarcntMask) >> RegisterCr3FieldScarcntShift)
}

// SetScarcnt Smartcard auto-retry count
func (r *registerCr3Type) SetScarcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldScarcntMask)|(uint32(value)<<RegisterCr3FieldScarcntShift))
}

const (
	RegisterCr3FieldWusShift = 20
	RegisterCr3FieldWusMask  = 0x300000
)

// GetWus Wakeup from Stop mode interrupt flag selection
func (r *registerCr3Type) GetWus() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldWusMask) >> RegisterCr3FieldWusShift)
}

// SetWus Wakeup from Stop mode interrupt flag selection
func (r *registerCr3Type) SetWus(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldWusMask)|(uint32(value)<<RegisterCr3FieldWusShift))
}

const (
	RegisterCr3FieldWufieShift = 22
	RegisterCr3FieldWufieMask  = 0x400000
)

// GetWufie Wakeup from Stop mode interrupt enable
func (r *registerCr3Type) GetWufie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldWufieMask) != 0
}

// SetWufie Wakeup from Stop mode interrupt enable
func (r *registerCr3Type) SetWufie(value bool) {
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
func (r *registerCr3Type) GetTxftie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldTxftieMask) != 0
}

// SetTxftie TXFIFO threshold interrupt enable
func (r *registerCr3Type) SetTxftie(value bool) {
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
func (r *registerCr3Type) GetTcbgtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldTcbgtieMask) != 0
}

// SetTcbgtie Transmission Complete before guard time, interrupt enable
func (r *registerCr3Type) SetTcbgtie(value bool) {
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
func (r *registerCr3Type) GetRxftcfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldRxftcfgMask) >> RegisterCr3FieldRxftcfgShift)
}

// SetRxftcfg Receive FIFO threshold configuration
func (r *registerCr3Type) SetRxftcfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldRxftcfgMask)|(uint32(value)<<RegisterCr3FieldRxftcfgShift))
}

const (
	RegisterCr3FieldRxftieShift = 28
	RegisterCr3FieldRxftieMask  = 0x10000000
)

// GetRxftie RXFIFO threshold interrupt enable
func (r *registerCr3Type) GetRxftie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldRxftieMask) != 0
}

// SetRxftie RXFIFO threshold interrupt enable
func (r *registerCr3Type) SetRxftie(value bool) {
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
func (r *registerCr3Type) GetTxftcfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCr3FieldTxftcfgMask) >> RegisterCr3FieldTxftcfgShift)
}

// SetTxftcfg TXFIFO threshold configuration
func (r *registerCr3Type) SetTxftcfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr3FieldTxftcfgMask)|(uint32(value)<<RegisterCr3FieldTxftcfgShift))
}

// registerBrrType Baud rate register
type registerBrrType uint32

const (
	RegisterBrrFieldBrr_0_3Shift = 0
	RegisterBrrFieldBrr_0_3Mask  = 0xf
)

// GetBrr_0_3 DIV_Fraction
func (r *registerBrrType) GetBrr_0_3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBrrFieldBrr_0_3Mask) >> RegisterBrrFieldBrr_0_3Shift)
}

// SetBrr_0_3 DIV_Fraction
func (r *registerBrrType) SetBrr_0_3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBrrFieldBrr_0_3Mask)|(uint32(value)<<RegisterBrrFieldBrr_0_3Shift))
}

const (
	RegisterBrrFieldBrr_4_15Shift = 4
	RegisterBrrFieldBrr_4_15Mask  = 0xfff0
)

// GetBrr_4_15 DIV_Mantissa
func (r *registerBrrType) GetBrr_4_15() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBrrFieldBrr_4_15Mask) >> RegisterBrrFieldBrr_4_15Shift)
}

// SetBrr_4_15 DIV_Mantissa
func (r *registerBrrType) SetBrr_4_15(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBrrFieldBrr_4_15Mask)|(uint32(value)<<RegisterBrrFieldBrr_4_15Shift))
}

// registerGtprType Guard time and prescaler register
type registerGtprType uint32

const (
	RegisterGtprFieldPscShift = 0
	RegisterGtprFieldPscMask  = 0xff
)

// GetPsc Prescaler value
func (r *registerGtprType) GetPsc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterGtprFieldPscMask) >> RegisterGtprFieldPscShift)
}

// SetPsc Prescaler value
func (r *registerGtprType) SetPsc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterGtprFieldPscMask)|(uint32(value)<<RegisterGtprFieldPscShift))
}

const (
	RegisterGtprFieldGtShift = 8
	RegisterGtprFieldGtMask  = 0xff00
)

// GetGt Guard time value
func (r *registerGtprType) GetGt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterGtprFieldGtMask) >> RegisterGtprFieldGtShift)
}

// SetGt Guard time value
func (r *registerGtprType) SetGt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterGtprFieldGtMask)|(uint32(value)<<RegisterGtprFieldGtShift))
}

// registerRtorType Receiver timeout register
type registerRtorType uint32

const (
	RegisterRtorFieldRtoShift = 0
	RegisterRtorFieldRtoMask  = 0xffffff
)

// GetRto Receiver timeout value
func (r *registerRtorType) GetRto() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRtorFieldRtoMask) >> RegisterRtorFieldRtoShift)
}

// SetRto Receiver timeout value
func (r *registerRtorType) SetRto(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtorFieldRtoMask)|(uint32(value)<<RegisterRtorFieldRtoShift))
}

const (
	RegisterRtorFieldBlenShift = 24
	RegisterRtorFieldBlenMask  = 0xff000000
)

// GetBlen Block Length
func (r *registerRtorType) GetBlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRtorFieldBlenMask) >> RegisterRtorFieldBlenShift)
}

// SetBlen Block Length
func (r *registerRtorType) SetBlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRtorFieldBlenMask)|(uint32(value)<<RegisterRtorFieldBlenShift))
}

// registerRqrType Request register
type registerRqrType uint32

const (
	RegisterRqrFieldAbrrqShift = 0
	RegisterRqrFieldAbrrqMask  = 0x1
)

// GetAbrrq Auto baud rate request
func (r *registerRqrType) GetAbrrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRqrFieldAbrrqMask) != 0
}

// SetAbrrq Auto baud rate request
func (r *registerRqrType) SetAbrrq(value bool) {
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
func (r *registerRqrType) GetSbkrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRqrFieldSbkrqMask) != 0
}

// SetSbkrq Send break request
func (r *registerRqrType) SetSbkrq(value bool) {
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
func (r *registerRqrType) GetMmrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRqrFieldMmrqMask) != 0
}

// SetMmrq Mute mode request
func (r *registerRqrType) SetMmrq(value bool) {
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
func (r *registerRqrType) GetRxfrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRqrFieldRxfrqMask) != 0
}

// SetRxfrq Receive data flush request
func (r *registerRqrType) SetRxfrq(value bool) {
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
func (r *registerRqrType) GetTxfrq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRqrFieldTxfrqMask) != 0
}

// SetTxfrq Transmit data flush request
func (r *registerRqrType) SetTxfrq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRqrFieldTxfrqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRqrFieldTxfrqMask)
	}
}

// registerIsrType Interrupt & status register
type registerIsrType uint32

const (
	RegisterIsrFieldPeShift = 0
	RegisterIsrFieldPeMask  = 0x1
)

// GetPe PE
func (r *registerIsrType) GetPe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldPeMask) != 0
}

// SetPe PE
func (r *registerIsrType) SetPe(value bool) {
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
func (r *registerIsrType) GetFe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldFeMask) != 0
}

// SetFe FE
func (r *registerIsrType) SetFe(value bool) {
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
func (r *registerIsrType) GetNf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldNfMask) != 0
}

// SetNf NF
func (r *registerIsrType) SetNf(value bool) {
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
func (r *registerIsrType) GetOre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldOreMask) != 0
}

// SetOre ORE
func (r *registerIsrType) SetOre(value bool) {
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
func (r *registerIsrType) GetIdle() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldIdleMask) != 0
}

// SetIdle IDLE
func (r *registerIsrType) SetIdle(value bool) {
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
func (r *registerIsrType) GetRxne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxneMask) != 0
}

// SetRxne RXNE
func (r *registerIsrType) SetRxne(value bool) {
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
func (r *registerIsrType) GetTc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcMask) != 0
}

// SetTc TC
func (r *registerIsrType) SetTc(value bool) {
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
func (r *registerIsrType) GetTxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxeMask) != 0
}

// SetTxe TXE
func (r *registerIsrType) SetTxe(value bool) {
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
func (r *registerIsrType) GetLbdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldLbdfMask) != 0
}

// SetLbdf LBDF
func (r *registerIsrType) SetLbdf(value bool) {
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
func (r *registerIsrType) GetCtsif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCtsifMask) != 0
}

// SetCtsif CTSIF
func (r *registerIsrType) SetCtsif(value bool) {
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
func (r *registerIsrType) GetCts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCtsMask) != 0
}

// SetCts CTS
func (r *registerIsrType) SetCts(value bool) {
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
func (r *registerIsrType) GetRtof() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRtofMask) != 0
}

// SetRtof RTOF
func (r *registerIsrType) SetRtof(value bool) {
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
func (r *registerIsrType) GetEobf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldEobfMask) != 0
}

// SetEobf EOBF
func (r *registerIsrType) SetEobf(value bool) {
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
func (r *registerIsrType) GetUdr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldUdrMask) != 0
}

// SetUdr SPI slave underrun error flag
func (r *registerIsrType) SetUdr(value bool) {
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
func (r *registerIsrType) GetAbre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldAbreMask) != 0
}

// SetAbre ABRE
func (r *registerIsrType) SetAbre(value bool) {
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
func (r *registerIsrType) GetAbrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldAbrfMask) != 0
}

// SetAbrf ABRF
func (r *registerIsrType) SetAbrf(value bool) {
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
func (r *registerIsrType) GetBusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldBusyMask) != 0
}

// SetBusy BUSY
func (r *registerIsrType) SetBusy(value bool) {
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
func (r *registerIsrType) GetCmf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCmfMask) != 0
}

// SetCmf CMF
func (r *registerIsrType) SetCmf(value bool) {
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
func (r *registerIsrType) GetSbkf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldSbkfMask) != 0
}

// SetSbkf SBKF
func (r *registerIsrType) SetSbkf(value bool) {
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
func (r *registerIsrType) GetRwu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRwuMask) != 0
}

// SetRwu RWU
func (r *registerIsrType) SetRwu(value bool) {
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
func (r *registerIsrType) GetWuf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldWufMask) != 0
}

// SetWuf WUF
func (r *registerIsrType) SetWuf(value bool) {
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
func (r *registerIsrType) GetTeack() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTeackMask) != 0
}

// SetTeack TEACK
func (r *registerIsrType) SetTeack(value bool) {
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
func (r *registerIsrType) GetReack() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldReackMask) != 0
}

// SetReack REACK
func (r *registerIsrType) SetReack(value bool) {
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
func (r *registerIsrType) GetTxfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxfeMask) != 0
}

// SetTxfe TXFIFO Empty
func (r *registerIsrType) SetTxfe(value bool) {
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
func (r *registerIsrType) GetRxff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxffMask) != 0
}

// SetRxff RXFIFO Full
func (r *registerIsrType) SetRxff(value bool) {
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
func (r *registerIsrType) GetTcbgt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcbgtMask) != 0
}

// SetTcbgt Transmission complete before guard time flag
func (r *registerIsrType) SetTcbgt(value bool) {
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
func (r *registerIsrType) GetRxft() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxftMask) != 0
}

// SetRxft RXFIFO threshold flag
func (r *registerIsrType) SetRxft(value bool) {
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
func (r *registerIsrType) GetTxft() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxftMask) != 0
}

// SetTxft TXFIFO threshold flag
func (r *registerIsrType) SetTxft(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxftMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxftMask)
	}
}

// registerIcrType Interrupt flag clear register
type registerIcrType uint32

const (
	RegisterIcrFieldPecfShift = 0
	RegisterIcrFieldPecfMask  = 0x1
)

// GetPecf Parity error clear flag
func (r *registerIcrType) GetPecf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldPecfMask) != 0
}

// SetPecf Parity error clear flag
func (r *registerIcrType) SetPecf(value bool) {
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
func (r *registerIcrType) GetFecf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldFecfMask) != 0
}

// SetFecf Framing error clear flag
func (r *registerIcrType) SetFecf(value bool) {
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
func (r *registerIcrType) GetNcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldNcfMask) != 0
}

// SetNcf Noise detected clear flag
func (r *registerIcrType) SetNcf(value bool) {
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
func (r *registerIcrType) GetOrecf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldOrecfMask) != 0
}

// SetOrecf Overrun error clear flag
func (r *registerIcrType) SetOrecf(value bool) {
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
func (r *registerIcrType) GetIdlecf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldIdlecfMask) != 0
}

// SetIdlecf Idle line detected clear flag
func (r *registerIcrType) SetIdlecf(value bool) {
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
func (r *registerIcrType) GetTxfecf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldTxfecfMask) != 0
}

// SetTxfecf TXFIFO empty clear flag
func (r *registerIcrType) SetTxfecf(value bool) {
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
func (r *registerIcrType) GetTccf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldTccfMask) != 0
}

// SetTccf Transmission complete clear flag
func (r *registerIcrType) SetTccf(value bool) {
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
func (r *registerIcrType) GetTcbgtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldTcbgtcMask) != 0
}

// SetTcbgtc Transmission complete before Guard time clear flag
func (r *registerIcrType) SetTcbgtc(value bool) {
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
func (r *registerIcrType) GetLbdcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldLbdcfMask) != 0
}

// SetLbdcf LIN break detection clear flag
func (r *registerIcrType) SetLbdcf(value bool) {
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
func (r *registerIcrType) GetCtscf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCtscfMask) != 0
}

// SetCtscf CTS clear flag
func (r *registerIcrType) SetCtscf(value bool) {
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
func (r *registerIcrType) GetRtocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldRtocfMask) != 0
}

// SetRtocf Receiver timeout clear flag
func (r *registerIcrType) SetRtocf(value bool) {
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
func (r *registerIcrType) GetEobcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldEobcfMask) != 0
}

// SetEobcf End of block clear flag
func (r *registerIcrType) SetEobcf(value bool) {
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
func (r *registerIcrType) GetUdrcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldUdrcfMask) != 0
}

// SetUdrcf SPI slave underrun clear flag
func (r *registerIcrType) SetUdrcf(value bool) {
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
func (r *registerIcrType) GetCmcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCmcfMask) != 0
}

// SetCmcf Character match clear flag
func (r *registerIcrType) SetCmcf(value bool) {
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
func (r *registerIcrType) GetWucf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldWucfMask) != 0
}

// SetWucf Wakeup from Stop mode clear flag
func (r *registerIcrType) SetWucf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldWucfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldWucfMask)
	}
}

// registerRdrType Receive data register
type registerRdrType uint32

const (
	RegisterRdrFieldRdrShift = 0
	RegisterRdrFieldRdrMask  = 0x1ff
)

// GetRdr Receive data value
func (r *registerRdrType) GetRdr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRdrFieldRdrMask) >> RegisterRdrFieldRdrShift)
}

// SetRdr Receive data value
func (r *registerRdrType) SetRdr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRdrFieldRdrMask)|(uint32(value)<<RegisterRdrFieldRdrShift))
}

// registerTdrType Transmit data register
type registerTdrType uint32

const (
	RegisterTdrFieldTdrShift = 0
	RegisterTdrFieldTdrMask  = 0x1ff
)

// GetTdr Transmit data value
func (r *registerTdrType) GetTdr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterTdrFieldTdrMask) >> RegisterTdrFieldTdrShift)
}

// SetTdr Transmit data value
func (r *registerTdrType) SetTdr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTdrFieldTdrMask)|(uint32(value)<<RegisterTdrFieldTdrShift))
}

// registerPrescType USART prescaler register
type registerPrescType uint32

const (
	RegisterPrescFieldPrescalerShift = 0
	RegisterPrescFieldPrescalerMask  = 0xf
)

// GetPrescaler Clock prescaler
func (r *registerPrescType) GetPrescaler() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPrescFieldPrescalerMask) >> RegisterPrescFieldPrescalerShift)
}

// SetPrescaler Clock prescaler
func (r *registerPrescType) SetPrescaler(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrescFieldPrescalerMask)|(uint32(value)<<RegisterPrescFieldPrescalerShift))
}
