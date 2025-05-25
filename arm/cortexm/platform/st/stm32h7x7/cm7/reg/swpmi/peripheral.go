//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package swpmi

import (
	"unsafe"
	"volatile"
)

var (
	Swpmi = (*_swpmi)(unsafe.Pointer(uintptr(0x40008800)))
)

type _swpmi struct {
	Cr  registerCrType
	Brr registerBrrType
	_   [4]byte
	Isr registerIsrType
	Icr registerIcrType
	Ier registerIerType
	Rfl registerRflType
	Tdr registerTdrType
	Rdr registerRdrType
	Or  registerOrType
}

// registerCrType SWPMI Configuration/Control register
type registerCrType uint32

const (
	RegisterCrFieldRxdmaShift = 0
	RegisterCrFieldRxdmaMask  = 0x1
)

// GetRxdma Reception DMA enable
func (r *registerCrType) GetRxdma() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRxdmaMask) != 0
}

// SetRxdma Reception DMA enable
func (r *registerCrType) SetRxdma(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRxdmaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRxdmaMask)
	}
}

const (
	RegisterCrFieldTxdmaShift = 1
	RegisterCrFieldTxdmaMask  = 0x2
)

// GetTxdma Transmission DMA enable
func (r *registerCrType) GetTxdma() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTxdmaMask) != 0
}

// SetTxdma Transmission DMA enable
func (r *registerCrType) SetTxdma(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTxdmaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTxdmaMask)
	}
}

const (
	RegisterCrFieldRxmodeShift = 2
	RegisterCrFieldRxmodeMask  = 0x4
)

// GetRxmode Reception buffering mode
func (r *registerCrType) GetRxmode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRxmodeMask) != 0
}

// SetRxmode Reception buffering mode
func (r *registerCrType) SetRxmode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRxmodeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRxmodeMask)
	}
}

const (
	RegisterCrFieldTxmodeShift = 3
	RegisterCrFieldTxmodeMask  = 0x8
)

// GetTxmode Transmission buffering mode
func (r *registerCrType) GetTxmode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTxmodeMask) != 0
}

// SetTxmode Transmission buffering mode
func (r *registerCrType) SetTxmode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTxmodeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTxmodeMask)
	}
}

const (
	RegisterCrFieldLpbkShift = 4
	RegisterCrFieldLpbkMask  = 0x10
)

// GetLpbk Loopback mode enable
func (r *registerCrType) GetLpbk() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldLpbkMask) != 0
}

// SetLpbk Loopback mode enable
func (r *registerCrType) SetLpbk(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldLpbkMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldLpbkMask)
	}
}

const (
	RegisterCrFieldSwpactShift = 5
	RegisterCrFieldSwpactMask  = 0x20
)

// GetSwpact Single wire protocol master interface activate
func (r *registerCrType) GetSwpact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSwpactMask) != 0
}

// SetSwpact Single wire protocol master interface activate
func (r *registerCrType) SetSwpact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSwpactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSwpactMask)
	}
}

const (
	RegisterCrFieldDeactShift = 10
	RegisterCrFieldDeactMask  = 0x400
)

// GetDeact Single wire protocol master interface deactivate
func (r *registerCrType) GetDeact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldDeactMask) != 0
}

// SetDeact Single wire protocol master interface deactivate
func (r *registerCrType) SetDeact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldDeactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldDeactMask)
	}
}

const (
	RegisterCrFieldSwptenShift = 11
	RegisterCrFieldSwptenMask  = 0x800
)

// GetSwpten Single wire protocol master transceiver enable
func (r *registerCrType) GetSwpten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSwptenMask) != 0
}

// SetSwpten Single wire protocol master transceiver enable
func (r *registerCrType) SetSwpten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSwptenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSwptenMask)
	}
}

// registerBrrType SWPMI Bitrate register
type registerBrrType uint32

const (
	RegisterBrrFieldBrShift = 0
	RegisterBrrFieldBrMask  = 0xff
)

// GetBr Bitrate prescaler
func (r *registerBrrType) GetBr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBrrFieldBrMask) >> RegisterBrrFieldBrShift)
}

// SetBr Bitrate prescaler
func (r *registerBrrType) SetBr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBrrFieldBrMask)|(uint32(value)<<RegisterBrrFieldBrShift))
}

// registerIsrType SWPMI Interrupt and Status register
type registerIsrType uint32

const (
	RegisterIsrFieldRxbffShift = 0
	RegisterIsrFieldRxbffMask  = 0x1
)

// GetRxbff Receive buffer full flag
func (r *registerIsrType) GetRxbff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxbffMask) != 0
}

// SetRxbff Receive buffer full flag
func (r *registerIsrType) SetRxbff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRxbffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRxbffMask)
	}
}

const (
	RegisterIsrFieldTxbefShift = 1
	RegisterIsrFieldTxbefMask  = 0x2
)

// GetTxbef Transmit buffer empty flag
func (r *registerIsrType) GetTxbef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxbefMask) != 0
}

// SetTxbef Transmit buffer empty flag
func (r *registerIsrType) SetTxbef(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxbefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxbefMask)
	}
}

const (
	RegisterIsrFieldRxberfShift = 2
	RegisterIsrFieldRxberfMask  = 0x4
)

// GetRxberf Receive CRC error flag
func (r *registerIsrType) GetRxberf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxberfMask) != 0
}

// SetRxberf Receive CRC error flag
func (r *registerIsrType) SetRxberf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRxberfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRxberfMask)
	}
}

const (
	RegisterIsrFieldRxovrfShift = 3
	RegisterIsrFieldRxovrfMask  = 0x8
)

// GetRxovrf Receive overrun error flag
func (r *registerIsrType) GetRxovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxovrfMask) != 0
}

// SetRxovrf Receive overrun error flag
func (r *registerIsrType) SetRxovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRxovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRxovrfMask)
	}
}

const (
	RegisterIsrFieldTxunrfShift = 4
	RegisterIsrFieldTxunrfMask  = 0x10
)

// GetTxunrf Transmit underrun error flag
func (r *registerIsrType) GetTxunrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxunrfMask) != 0
}

// SetTxunrf Transmit underrun error flag
func (r *registerIsrType) SetTxunrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxunrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxunrfMask)
	}
}

const (
	RegisterIsrFieldRxneShift = 5
	RegisterIsrFieldRxneMask  = 0x20
)

// GetRxne Receive data register not empty
func (r *registerIsrType) GetRxne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxneMask) != 0
}

// SetRxne Receive data register not empty
func (r *registerIsrType) SetRxne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRxneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRxneMask)
	}
}

const (
	RegisterIsrFieldTxeShift = 6
	RegisterIsrFieldTxeMask  = 0x40
)

// GetTxe Transmit data register empty
func (r *registerIsrType) GetTxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxeMask) != 0
}

// SetTxe Transmit data register empty
func (r *registerIsrType) SetTxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxeMask)
	}
}

const (
	RegisterIsrFieldTcfShift = 7
	RegisterIsrFieldTcfMask  = 0x80
)

// GetTcf Transfer complete flag
func (r *registerIsrType) GetTcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTcfMask) != 0
}

// SetTcf Transfer complete flag
func (r *registerIsrType) SetTcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTcfMask)
	}
}

const (
	RegisterIsrFieldSrfShift = 8
	RegisterIsrFieldSrfMask  = 0x100
)

// GetSrf Slave resume flag
func (r *registerIsrType) GetSrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldSrfMask) != 0
}

// SetSrf Slave resume flag
func (r *registerIsrType) SetSrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldSrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldSrfMask)
	}
}

const (
	RegisterIsrFieldSuspShift = 9
	RegisterIsrFieldSuspMask  = 0x200
)

// GetSusp SUSPEND flag
func (r *registerIsrType) GetSusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldSuspMask) != 0
}

// SetSusp SUSPEND flag
func (r *registerIsrType) SetSusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldSuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldSuspMask)
	}
}

const (
	RegisterIsrFieldDeactfShift = 10
	RegisterIsrFieldDeactfMask  = 0x400
)

// GetDeactf DEACTIVATED flag
func (r *registerIsrType) GetDeactf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldDeactfMask) != 0
}

// SetDeactf DEACTIVATED flag
func (r *registerIsrType) SetDeactf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldDeactfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldDeactfMask)
	}
}

const (
	RegisterIsrFieldRdyfShift = 11
	RegisterIsrFieldRdyfMask  = 0x800
)

// GetRdyf transceiver ready flag
func (r *registerIsrType) GetRdyf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRdyfMask) != 0
}

// SetRdyf transceiver ready flag
func (r *registerIsrType) SetRdyf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRdyfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRdyfMask)
	}
}

// registerIcrType SWPMI Interrupt Flag Clear register
type registerIcrType uint32

const (
	RegisterIcrFieldCrxbffShift = 0
	RegisterIcrFieldCrxbffMask  = 0x1
)

// GetCrxbff Clear receive buffer full flag
func (r *registerIcrType) GetCrxbff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCrxbffMask) != 0
}

// SetCrxbff Clear receive buffer full flag
func (r *registerIcrType) SetCrxbff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCrxbffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCrxbffMask)
	}
}

const (
	RegisterIcrFieldCtxbefShift = 1
	RegisterIcrFieldCtxbefMask  = 0x2
)

// GetCtxbef Clear transmit buffer empty flag
func (r *registerIcrType) GetCtxbef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCtxbefMask) != 0
}

// SetCtxbef Clear transmit buffer empty flag
func (r *registerIcrType) SetCtxbef(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCtxbefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCtxbefMask)
	}
}

const (
	RegisterIcrFieldCrxberfShift = 2
	RegisterIcrFieldCrxberfMask  = 0x4
)

// GetCrxberf Clear receive CRC error flag
func (r *registerIcrType) GetCrxberf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCrxberfMask) != 0
}

// SetCrxberf Clear receive CRC error flag
func (r *registerIcrType) SetCrxberf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCrxberfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCrxberfMask)
	}
}

const (
	RegisterIcrFieldCrxovrfShift = 3
	RegisterIcrFieldCrxovrfMask  = 0x8
)

// GetCrxovrf Clear receive overrun error flag
func (r *registerIcrType) GetCrxovrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCrxovrfMask) != 0
}

// SetCrxovrf Clear receive overrun error flag
func (r *registerIcrType) SetCrxovrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCrxovrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCrxovrfMask)
	}
}

const (
	RegisterIcrFieldCtxunrfShift = 4
	RegisterIcrFieldCtxunrfMask  = 0x10
)

// GetCtxunrf Clear transmit underrun error flag
func (r *registerIcrType) GetCtxunrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCtxunrfMask) != 0
}

// SetCtxunrf Clear transmit underrun error flag
func (r *registerIcrType) SetCtxunrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCtxunrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCtxunrfMask)
	}
}

const (
	RegisterIcrFieldCtcfShift = 7
	RegisterIcrFieldCtcfMask  = 0x80
)

// GetCtcf Clear transfer complete flag
func (r *registerIcrType) GetCtcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCtcfMask) != 0
}

// SetCtcf Clear transfer complete flag
func (r *registerIcrType) SetCtcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCtcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCtcfMask)
	}
}

const (
	RegisterIcrFieldCsrfShift = 8
	RegisterIcrFieldCsrfMask  = 0x100
)

// GetCsrf Clear slave resume flag
func (r *registerIcrType) GetCsrf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCsrfMask) != 0
}

// SetCsrf Clear slave resume flag
func (r *registerIcrType) SetCsrf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCsrfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCsrfMask)
	}
}

const (
	RegisterIcrFieldCrdyfShift = 11
	RegisterIcrFieldCrdyfMask  = 0x800
)

// GetCrdyf Clear transceiver ready flag
func (r *registerIcrType) GetCrdyf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCrdyfMask) != 0
}

// SetCrdyf Clear transceiver ready flag
func (r *registerIcrType) SetCrdyf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCrdyfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCrdyfMask)
	}
}

// registerIerType SWPMI Interrupt Enable register
type registerIerType uint32

const (
	RegisterIerFieldRxbfieShift = 0
	RegisterIerFieldRxbfieMask  = 0x1
)

// GetRxbfie Receive buffer full interrupt enable
func (r *registerIerType) GetRxbfie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRxbfieMask) != 0
}

// SetRxbfie Receive buffer full interrupt enable
func (r *registerIerType) SetRxbfie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRxbfieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRxbfieMask)
	}
}

const (
	RegisterIerFieldTxbeieShift = 1
	RegisterIerFieldTxbeieMask  = 0x2
)

// GetTxbeie Transmit buffer empty interrupt enable
func (r *registerIerType) GetTxbeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTxbeieMask) != 0
}

// SetTxbeie Transmit buffer empty interrupt enable
func (r *registerIerType) SetTxbeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTxbeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTxbeieMask)
	}
}

const (
	RegisterIerFieldRxberieShift = 2
	RegisterIerFieldRxberieMask  = 0x4
)

// GetRxberie Receive CRC error interrupt enable
func (r *registerIerType) GetRxberie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRxberieMask) != 0
}

// SetRxberie Receive CRC error interrupt enable
func (r *registerIerType) SetRxberie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRxberieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRxberieMask)
	}
}

const (
	RegisterIerFieldRxovrieShift = 3
	RegisterIerFieldRxovrieMask  = 0x8
)

// GetRxovrie Receive overrun error interrupt enable
func (r *registerIerType) GetRxovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRxovrieMask) != 0
}

// SetRxovrie Receive overrun error interrupt enable
func (r *registerIerType) SetRxovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRxovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRxovrieMask)
	}
}

const (
	RegisterIerFieldTxunrieShift = 4
	RegisterIerFieldTxunrieMask  = 0x10
)

// GetTxunrie Transmit underrun error interrupt enable
func (r *registerIerType) GetTxunrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTxunrieMask) != 0
}

// SetTxunrie Transmit underrun error interrupt enable
func (r *registerIerType) SetTxunrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTxunrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTxunrieMask)
	}
}

const (
	RegisterIerFieldRieShift = 5
	RegisterIerFieldRieMask  = 0x20
)

// GetRie Receive interrupt enable
func (r *registerIerType) GetRie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRieMask) != 0
}

// SetRie Receive interrupt enable
func (r *registerIerType) SetRie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRieMask)
	}
}

const (
	RegisterIerFieldTieShift = 6
	RegisterIerFieldTieMask  = 0x40
)

// GetTie Transmit interrupt enable
func (r *registerIerType) GetTie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTieMask) != 0
}

// SetTie Transmit interrupt enable
func (r *registerIerType) SetTie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTieMask)
	}
}

const (
	RegisterIerFieldTcieShift = 7
	RegisterIerFieldTcieMask  = 0x80
)

// GetTcie Transmit complete interrupt enable
func (r *registerIerType) GetTcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTcieMask) != 0
}

// SetTcie Transmit complete interrupt enable
func (r *registerIerType) SetTcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTcieMask)
	}
}

const (
	RegisterIerFieldSrieShift = 8
	RegisterIerFieldSrieMask  = 0x100
)

// GetSrie Slave resume interrupt enable
func (r *registerIerType) GetSrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldSrieMask) != 0
}

// SetSrie Slave resume interrupt enable
func (r *registerIerType) SetSrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldSrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldSrieMask)
	}
}

const (
	RegisterIerFieldRdyieShift = 11
	RegisterIerFieldRdyieMask  = 0x800
)

// GetRdyie Transceiver ready interrupt enable
func (r *registerIerType) GetRdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRdyieMask) != 0
}

// SetRdyie Transceiver ready interrupt enable
func (r *registerIerType) SetRdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRdyieMask)
	}
}

// registerRflType SWPMI Receive Frame Length register
type registerRflType uint32

const (
	RegisterRflFieldRflShift = 0
	RegisterRflFieldRflMask  = 0x1f
)

// GetRfl Receive frame length
func (r *registerRflType) GetRfl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRflFieldRflMask) >> RegisterRflFieldRflShift)
}

// SetRfl Receive frame length
func (r *registerRflType) SetRfl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRflFieldRflMask)|(uint32(value)<<RegisterRflFieldRflShift))
}

// registerTdrType SWPMI Transmit data register
type registerTdrType uint32

const (
	RegisterTdrFieldTdShift = 0
	RegisterTdrFieldTdMask  = 0xffffffff
)

// GetTd Transmit data
func (r *registerTdrType) GetTd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTdrFieldTdMask) >> RegisterTdrFieldTdShift)
}

// SetTd Transmit data
func (r *registerTdrType) SetTd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTdrFieldTdMask)|(uint32(value)<<RegisterTdrFieldTdShift))
}

// registerRdrType SWPMI Receive data register
type registerRdrType uint32

const (
	RegisterRdrFieldRdShift = 0
	RegisterRdrFieldRdMask  = 0xffffffff
)

// GetRd received data
func (r *registerRdrType) GetRd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRdrFieldRdMask) >> RegisterRdrFieldRdShift)
}

// SetRd received data
func (r *registerRdrType) SetRd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRdrFieldRdMask)|(uint32(value)<<RegisterRdrFieldRdShift))
}

// registerOrType SWPMI Option register
type registerOrType uint32

const (
	RegisterOrFieldSwp_tbypShift = 0
	RegisterOrFieldSwp_tbypMask  = 0x1
)

// GetSwp_tbyp SWP transceiver bypass
func (r *registerOrType) GetSwp_tbyp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOrFieldSwp_tbypMask) != 0
}

// SetSwp_tbyp SWP transceiver bypass
func (r *registerOrType) SetSwp_tbyp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOrFieldSwp_tbypMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOrFieldSwp_tbypMask)
	}
}

const (
	RegisterOrFieldSwp_classShift = 1
	RegisterOrFieldSwp_classMask  = 0x2
)

// GetSwp_class SWP class selection
func (r *registerOrType) GetSwp_class() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOrFieldSwp_classMask) != 0
}

// SetSwp_class SWP class selection
func (r *registerOrType) SetSwp_class(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOrFieldSwp_classMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOrFieldSwp_classMask)
	}
}
