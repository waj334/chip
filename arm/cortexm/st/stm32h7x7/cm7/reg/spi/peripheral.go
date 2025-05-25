package spi

import (
	"unsafe"
	"volatile"
)

var (
	Spi1 = (*_spi)(unsafe.Pointer(uintptr(0x40013000)))
	Spi2 = (*_spi)(unsafe.Pointer(uintptr(0x40003800)))
	Spi3 = (*_spi)(unsafe.Pointer(uintptr(0x40003c00)))
	Spi4 = (*_spi)(unsafe.Pointer(uintptr(0x40013400)))
	Spi5 = (*_spi)(unsafe.Pointer(uintptr(0x40015000)))
	Spi6 = (*_spi)(unsafe.Pointer(uintptr(0x58001400)))

	Instances = [6]*_spi{
		Spi1,
		Spi2,
		Spi3,
		Spi4,
		Spi5,
		Spi6,
	}
)

type _spi struct {
	Cr1     registerCr1Type
	Cr2     registerCr2Type
	Cfg1    registerCfg1Type
	Cfg2    registerCfg2Type
	Ier     registerIerType
	Sr      registerSrType
	Ifcr    registerIfcrType
	_       [4]byte
	Txdr    registerTxdrType
	_       [12]byte
	Rxdr    registerRxdrType
	_       [12]byte
	Crcpoly registerCrcpolyType
	Txcrc   registerTxcrcType
	Rxcrc   registerRxcrcType
	Udrdr   registerUdrdrType
	Cgfr    registerCgfrType
}

// registerCr1Type control register 1
type registerCr1Type uint32

const (
	RegisterCr1FieldSpeShift = 0
	RegisterCr1FieldSpeMask  = 0x1
)

// GetSpe Serial Peripheral Enable
func (r *registerCr1Type) GetSpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldSpeMask) != 0
}

// SetSpe Serial Peripheral Enable
func (r *registerCr1Type) SetSpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldSpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldSpeMask)
	}
}

const (
	RegisterCr1FieldMasrxShift = 8
	RegisterCr1FieldMasrxMask  = 0x100
)

// GetMasrx Master automatic SUSP in Receive mode
func (r *registerCr1Type) GetMasrx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldMasrxMask) != 0
}

// SetMasrx Master automatic SUSP in Receive mode
func (r *registerCr1Type) SetMasrx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldMasrxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldMasrxMask)
	}
}

const (
	RegisterCr1FieldCstartShift = 9
	RegisterCr1FieldCstartMask  = 0x200
)

// GetCstart Master transfer start
func (r *registerCr1Type) GetCstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCstartMask) != 0
}

const (
	RegisterCr1FieldCsuspShift = 10
	RegisterCr1FieldCsuspMask  = 0x400
)

// SetCsusp Master SUSPend request
func (r *registerCr1Type) SetCsusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldCsuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCsuspMask)
	}
}

const (
	RegisterCr1FieldHddirShift = 11
	RegisterCr1FieldHddirMask  = 0x800
)

// GetHddir Rx/Tx direction at Half-duplex mode
func (r *registerCr1Type) GetHddir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldHddirMask) != 0
}

// SetHddir Rx/Tx direction at Half-duplex mode
func (r *registerCr1Type) SetHddir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldHddirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldHddirMask)
	}
}

const (
	RegisterCr1FieldSsiShift = 12
	RegisterCr1FieldSsiMask  = 0x1000
)

// GetSsi Internal SS signal input level
func (r *registerCr1Type) GetSsi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldSsiMask) != 0
}

// SetSsi Internal SS signal input level
func (r *registerCr1Type) SetSsi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldSsiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldSsiMask)
	}
}

const (
	RegisterCr1FieldCrc33_17Shift = 13
	RegisterCr1FieldCrc33_17Mask  = 0x2000
)

// GetCrc33_17 32-bit CRC polynomial configuration
func (r *registerCr1Type) GetCrc33_17() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCrc33_17Mask) != 0
}

// SetCrc33_17 32-bit CRC polynomial configuration
func (r *registerCr1Type) SetCrc33_17(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldCrc33_17Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCrc33_17Mask)
	}
}

const (
	RegisterCr1FieldRcrciShift = 14
	RegisterCr1FieldRcrciMask  = 0x4000
)

// GetRcrci CRC calculation initialization pattern control for receiver
func (r *registerCr1Type) GetRcrci() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRcrciMask) != 0
}

// SetRcrci CRC calculation initialization pattern control for receiver
func (r *registerCr1Type) SetRcrci(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldRcrciMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldRcrciMask)
	}
}

const (
	RegisterCr1FieldTcrciShift = 15
	RegisterCr1FieldTcrciMask  = 0x8000
)

// GetTcrci CRC calculation initialization pattern control for transmitter
func (r *registerCr1Type) GetTcrci() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTcrciMask) != 0
}

// SetTcrci CRC calculation initialization pattern control for transmitter
func (r *registerCr1Type) SetTcrci(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldTcrciMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldTcrciMask)
	}
}

const (
	RegisterCr1FieldIolockShift = 16
	RegisterCr1FieldIolockMask  = 0x10000
)

// GetIolock Locking the AF configuration of associated IOs
func (r *registerCr1Type) GetIolock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldIolockMask) != 0
}

// registerCr2Type control register 2
type registerCr2Type uint32

const (
	RegisterCr2FieldTsizeShift = 0
	RegisterCr2FieldTsizeMask  = 0xffff
)

// GetTsize Number of data at current transfer
func (r *registerCr2Type) GetTsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTsizeMask) >> RegisterCr2FieldTsizeShift)
}

// SetTsize Number of data at current transfer
func (r *registerCr2Type) SetTsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTsizeMask)|(uint32(value)<<RegisterCr2FieldTsizeShift))
}

const (
	RegisterCr2FieldTserShift = 16
	RegisterCr2FieldTserMask  = 0xffff0000
)

// GetTser Number of data transfer extension to be reload into TSIZE just when a previous
func (r *registerCr2Type) GetTser() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTserMask) >> RegisterCr2FieldTserShift)
}

// registerCfg1Type configuration register 1
type registerCfg1Type uint32

const (
	RegisterCfg1FieldDsizeShift = 0
	RegisterCfg1FieldDsizeMask  = 0x1f
)

// GetDsize Number of bits in at single SPI data frame
func (r *registerCfg1Type) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldDsizeMask) >> RegisterCfg1FieldDsizeShift)
}

// SetDsize Number of bits in at single SPI data frame
func (r *registerCfg1Type) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldDsizeMask)|(uint32(value)<<RegisterCfg1FieldDsizeShift))
}

const (
	RegisterCfg1FieldFthvlShift = 5
	RegisterCfg1FieldFthvlMask  = 0x1e0
)

// GetFthvl threshold level
func (r *registerCfg1Type) GetFthvl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldFthvlMask) >> RegisterCfg1FieldFthvlShift)
}

// SetFthvl threshold level
func (r *registerCfg1Type) SetFthvl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldFthvlMask)|(uint32(value)<<RegisterCfg1FieldFthvlShift))
}

const (
	RegisterCfg1FieldUdrcfgShift = 9
	RegisterCfg1FieldUdrcfgMask  = 0x600
)

// GetUdrcfg Behavior of slave transmitter at underrun condition
func (r *registerCfg1Type) GetUdrcfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldUdrcfgMask) >> RegisterCfg1FieldUdrcfgShift)
}

// SetUdrcfg Behavior of slave transmitter at underrun condition
func (r *registerCfg1Type) SetUdrcfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldUdrcfgMask)|(uint32(value)<<RegisterCfg1FieldUdrcfgShift))
}

const (
	RegisterCfg1FieldUdrdetShift = 11
	RegisterCfg1FieldUdrdetMask  = 0x1800
)

// GetUdrdet Detection of underrun condition at slave transmitter
func (r *registerCfg1Type) GetUdrdet() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldUdrdetMask) >> RegisterCfg1FieldUdrdetShift)
}

// SetUdrdet Detection of underrun condition at slave transmitter
func (r *registerCfg1Type) SetUdrdet(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldUdrdetMask)|(uint32(value)<<RegisterCfg1FieldUdrdetShift))
}

const (
	RegisterCfg1FieldRxdmaenShift = 14
	RegisterCfg1FieldRxdmaenMask  = 0x4000
)

// GetRxdmaen Rx DMA stream enable
func (r *registerCfg1Type) GetRxdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldRxdmaenMask) != 0
}

// SetRxdmaen Rx DMA stream enable
func (r *registerCfg1Type) SetRxdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg1FieldRxdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldRxdmaenMask)
	}
}

const (
	RegisterCfg1FieldTxdmaenShift = 15
	RegisterCfg1FieldTxdmaenMask  = 0x8000
)

// GetTxdmaen Tx DMA stream enable
func (r *registerCfg1Type) GetTxdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldTxdmaenMask) != 0
}

// SetTxdmaen Tx DMA stream enable
func (r *registerCfg1Type) SetTxdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg1FieldTxdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldTxdmaenMask)
	}
}

const (
	RegisterCfg1FieldCrcsizeShift = 16
	RegisterCfg1FieldCrcsizeMask  = 0x1f0000
)

// GetCrcsize Length of CRC frame to be transacted and compared
func (r *registerCfg1Type) GetCrcsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldCrcsizeMask) >> RegisterCfg1FieldCrcsizeShift)
}

// SetCrcsize Length of CRC frame to be transacted and compared
func (r *registerCfg1Type) SetCrcsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldCrcsizeMask)|(uint32(value)<<RegisterCfg1FieldCrcsizeShift))
}

const (
	RegisterCfg1FieldCrcenShift = 22
	RegisterCfg1FieldCrcenMask  = 0x400000
)

// GetCrcen Hardware CRC computation enable
func (r *registerCfg1Type) GetCrcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldCrcenMask) != 0
}

// SetCrcen Hardware CRC computation enable
func (r *registerCfg1Type) SetCrcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg1FieldCrcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldCrcenMask)
	}
}

const (
	RegisterCfg1FieldMbrShift = 28
	RegisterCfg1FieldMbrMask  = 0x70000000
)

// GetMbr Master baud rate
func (r *registerCfg1Type) GetMbr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldMbrMask) >> RegisterCfg1FieldMbrShift)
}

// SetMbr Master baud rate
func (r *registerCfg1Type) SetMbr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldMbrMask)|(uint32(value)<<RegisterCfg1FieldMbrShift))
}

// registerCfg2Type configuration register 2
type registerCfg2Type uint32

const (
	RegisterCfg2FieldMssiShift = 0
	RegisterCfg2FieldMssiMask  = 0xf
)

// GetMssi Master SS Idleness
func (r *registerCfg2Type) GetMssi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldMssiMask) >> RegisterCfg2FieldMssiShift)
}

// SetMssi Master SS Idleness
func (r *registerCfg2Type) SetMssi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldMssiMask)|(uint32(value)<<RegisterCfg2FieldMssiShift))
}

const (
	RegisterCfg2FieldMidiShift = 4
	RegisterCfg2FieldMidiMask  = 0xf0
)

// GetMidi Master Inter-Data Idleness
func (r *registerCfg2Type) GetMidi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldMidiMask) >> RegisterCfg2FieldMidiShift)
}

// SetMidi Master Inter-Data Idleness
func (r *registerCfg2Type) SetMidi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldMidiMask)|(uint32(value)<<RegisterCfg2FieldMidiShift))
}

const (
	RegisterCfg2FieldIoswpShift = 15
	RegisterCfg2FieldIoswpMask  = 0x8000
)

// GetIoswp Swap functionality of MISO and MOSI pins
func (r *registerCfg2Type) GetIoswp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldIoswpMask) != 0
}

// SetIoswp Swap functionality of MISO and MOSI pins
func (r *registerCfg2Type) SetIoswp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg2FieldIoswpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldIoswpMask)
	}
}

const (
	RegisterCfg2FieldCommShift = 17
	RegisterCfg2FieldCommMask  = 0x60000
)

// GetComm SPI Communication Mode
func (r *registerCfg2Type) GetComm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldCommMask) >> RegisterCfg2FieldCommShift)
}

// SetComm SPI Communication Mode
func (r *registerCfg2Type) SetComm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldCommMask)|(uint32(value)<<RegisterCfg2FieldCommShift))
}

const (
	RegisterCfg2FieldSpShift = 19
	RegisterCfg2FieldSpMask  = 0x380000
)

// GetSp Serial Protocol
func (r *registerCfg2Type) GetSp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldSpMask) >> RegisterCfg2FieldSpShift)
}

// SetSp Serial Protocol
func (r *registerCfg2Type) SetSp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldSpMask)|(uint32(value)<<RegisterCfg2FieldSpShift))
}

const (
	RegisterCfg2FieldMasterShift = 22
	RegisterCfg2FieldMasterMask  = 0x400000
)

// GetMaster SPI Master
func (r *registerCfg2Type) GetMaster() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldMasterMask) != 0
}

// SetMaster SPI Master
func (r *registerCfg2Type) SetMaster(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg2FieldMasterMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldMasterMask)
	}
}

const (
	RegisterCfg2FieldLsbfrstShift = 23
	RegisterCfg2FieldLsbfrstMask  = 0x800000
)

// GetLsbfrst Data frame format
func (r *registerCfg2Type) GetLsbfrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldLsbfrstMask) != 0
}

// SetLsbfrst Data frame format
func (r *registerCfg2Type) SetLsbfrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg2FieldLsbfrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldLsbfrstMask)
	}
}

const (
	RegisterCfg2FieldCphaShift = 24
	RegisterCfg2FieldCphaMask  = 0x1000000
)

// GetCpha Clock phase
func (r *registerCfg2Type) GetCpha() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldCphaMask) != 0
}

// SetCpha Clock phase
func (r *registerCfg2Type) SetCpha(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg2FieldCphaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldCphaMask)
	}
}

const (
	RegisterCfg2FieldCpolShift = 25
	RegisterCfg2FieldCpolMask  = 0x2000000
)

// GetCpol Clock polarity
func (r *registerCfg2Type) GetCpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldCpolMask) != 0
}

// SetCpol Clock polarity
func (r *registerCfg2Type) SetCpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg2FieldCpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldCpolMask)
	}
}

const (
	RegisterCfg2FieldSsmShift = 26
	RegisterCfg2FieldSsmMask  = 0x4000000
)

// GetSsm Software management of SS signal input
func (r *registerCfg2Type) GetSsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldSsmMask) != 0
}

// SetSsm Software management of SS signal input
func (r *registerCfg2Type) SetSsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg2FieldSsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldSsmMask)
	}
}

const (
	RegisterCfg2FieldSsiopShift = 28
	RegisterCfg2FieldSsiopMask  = 0x10000000
)

// GetSsiop SS input/output polarity
func (r *registerCfg2Type) GetSsiop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldSsiopMask) != 0
}

// SetSsiop SS input/output polarity
func (r *registerCfg2Type) SetSsiop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg2FieldSsiopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldSsiopMask)
	}
}

const (
	RegisterCfg2FieldSsoeShift = 29
	RegisterCfg2FieldSsoeMask  = 0x20000000
)

// GetSsoe SS output enable
func (r *registerCfg2Type) GetSsoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldSsoeMask) != 0
}

// SetSsoe SS output enable
func (r *registerCfg2Type) SetSsoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg2FieldSsoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldSsoeMask)
	}
}

const (
	RegisterCfg2FieldSsomShift = 30
	RegisterCfg2FieldSsomMask  = 0x40000000
)

// GetSsom SS output management in master mode
func (r *registerCfg2Type) GetSsom() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldSsomMask) != 0
}

// SetSsom SS output management in master mode
func (r *registerCfg2Type) SetSsom(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg2FieldSsomMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldSsomMask)
	}
}

const (
	RegisterCfg2FieldAfcntrShift = 31
	RegisterCfg2FieldAfcntrMask  = 0x80000000
)

// GetAfcntr Alternate function GPIOs control
func (r *registerCfg2Type) GetAfcntr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldAfcntrMask) != 0
}

// SetAfcntr Alternate function GPIOs control
func (r *registerCfg2Type) SetAfcntr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg2FieldAfcntrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldAfcntrMask)
	}
}

// registerIerType Interrupt Enable Register
type registerIerType uint32

const (
	RegisterIerFieldRxpieShift = 0
	RegisterIerFieldRxpieMask  = 0x1
)

// GetRxpie RXP Interrupt Enable
func (r *registerIerType) GetRxpie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRxpieMask) != 0
}

// SetRxpie RXP Interrupt Enable
func (r *registerIerType) SetRxpie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRxpieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRxpieMask)
	}
}

const (
	RegisterIerFieldTxpieShift = 1
	RegisterIerFieldTxpieMask  = 0x2
)

// GetTxpie TXP interrupt enable
func (r *registerIerType) GetTxpie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTxpieMask) != 0
}

const (
	RegisterIerFieldDpxpieShift = 2
	RegisterIerFieldDpxpieMask  = 0x4
)

// GetDpxpie DXP interrupt enabled
func (r *registerIerType) GetDpxpie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldDpxpieMask) != 0
}

const (
	RegisterIerFieldEotieShift = 3
	RegisterIerFieldEotieMask  = 0x8
)

// GetEotie EOT, SUSP and TXC interrupt enable
func (r *registerIerType) GetEotie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldEotieMask) != 0
}

// SetEotie EOT, SUSP and TXC interrupt enable
func (r *registerIerType) SetEotie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldEotieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldEotieMask)
	}
}

const (
	RegisterIerFieldTxtfieShift = 4
	RegisterIerFieldTxtfieMask  = 0x10
)

// GetTxtfie TXTFIE interrupt enable
func (r *registerIerType) GetTxtfie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTxtfieMask) != 0
}

// SetTxtfie TXTFIE interrupt enable
func (r *registerIerType) SetTxtfie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTxtfieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTxtfieMask)
	}
}

const (
	RegisterIerFieldUdrieShift = 5
	RegisterIerFieldUdrieMask  = 0x20
)

// GetUdrie UDR interrupt enable
func (r *registerIerType) GetUdrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldUdrieMask) != 0
}

// SetUdrie UDR interrupt enable
func (r *registerIerType) SetUdrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldUdrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldUdrieMask)
	}
}

const (
	RegisterIerFieldOvrieShift = 6
	RegisterIerFieldOvrieMask  = 0x40
)

// GetOvrie OVR interrupt enable
func (r *registerIerType) GetOvrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldOvrieMask) != 0
}

// SetOvrie OVR interrupt enable
func (r *registerIerType) SetOvrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldOvrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldOvrieMask)
	}
}

const (
	RegisterIerFieldCrceieShift = 7
	RegisterIerFieldCrceieMask  = 0x80
)

// GetCrceie CRC Interrupt enable
func (r *registerIerType) GetCrceie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldCrceieMask) != 0
}

// SetCrceie CRC Interrupt enable
func (r *registerIerType) SetCrceie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldCrceieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldCrceieMask)
	}
}

const (
	RegisterIerFieldTifreieShift = 8
	RegisterIerFieldTifreieMask  = 0x100
)

// GetTifreie TIFRE interrupt enable
func (r *registerIerType) GetTifreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTifreieMask) != 0
}

// SetTifreie TIFRE interrupt enable
func (r *registerIerType) SetTifreie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTifreieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTifreieMask)
	}
}

const (
	RegisterIerFieldModfieShift = 9
	RegisterIerFieldModfieMask  = 0x200
)

// GetModfie Mode Fault interrupt enable
func (r *registerIerType) GetModfie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldModfieMask) != 0
}

// SetModfie Mode Fault interrupt enable
func (r *registerIerType) SetModfie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldModfieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldModfieMask)
	}
}

const (
	RegisterIerFieldTserfieShift = 10
	RegisterIerFieldTserfieMask  = 0x400
)

// GetTserfie Additional number of transactions reload interrupt enable
func (r *registerIerType) GetTserfie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTserfieMask) != 0
}

// SetTserfie Additional number of transactions reload interrupt enable
func (r *registerIerType) SetTserfie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTserfieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTserfieMask)
	}
}

// registerSrType Status Register
type registerSrType uint32

const (
	RegisterSrFieldRxpShift = 0
	RegisterSrFieldRxpMask  = 0x1
)

// GetRxp Rx-Packet available
func (r *registerSrType) GetRxp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRxpMask) != 0
}

// SetRxp Rx-Packet available
func (r *registerSrType) SetRxp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldRxpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldRxpMask)
	}
}

const (
	RegisterSrFieldTxpShift = 1
	RegisterSrFieldTxpMask  = 0x2
)

// GetTxp Tx-Packet space available
func (r *registerSrType) GetTxp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTxpMask) != 0
}

// SetTxp Tx-Packet space available
func (r *registerSrType) SetTxp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldTxpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldTxpMask)
	}
}

const (
	RegisterSrFieldDxpShift = 2
	RegisterSrFieldDxpMask  = 0x4
)

// GetDxp Duplex Packet
func (r *registerSrType) GetDxp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDxpMask) != 0
}

// SetDxp Duplex Packet
func (r *registerSrType) SetDxp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldDxpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldDxpMask)
	}
}

const (
	RegisterSrFieldEotShift = 3
	RegisterSrFieldEotMask  = 0x8
)

// GetEot End Of Transfer
func (r *registerSrType) GetEot() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldEotMask) != 0
}

// SetEot End Of Transfer
func (r *registerSrType) SetEot(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldEotMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldEotMask)
	}
}

const (
	RegisterSrFieldTxtfShift = 4
	RegisterSrFieldTxtfMask  = 0x10
)

// GetTxtf Transmission Transfer Filled
func (r *registerSrType) GetTxtf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTxtfMask) != 0
}

// SetTxtf Transmission Transfer Filled
func (r *registerSrType) SetTxtf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldTxtfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldTxtfMask)
	}
}

const (
	RegisterSrFieldUdrShift = 5
	RegisterSrFieldUdrMask  = 0x20
)

// GetUdr Underrun at slave transmission mode
func (r *registerSrType) GetUdr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldUdrMask) != 0
}

// SetUdr Underrun at slave transmission mode
func (r *registerSrType) SetUdr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldUdrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldUdrMask)
	}
}

const (
	RegisterSrFieldOvrShift = 6
	RegisterSrFieldOvrMask  = 0x40
)

// GetOvr Overrun
func (r *registerSrType) GetOvr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldOvrMask) != 0
}

// SetOvr Overrun
func (r *registerSrType) SetOvr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldOvrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldOvrMask)
	}
}

const (
	RegisterSrFieldCrceShift = 7
	RegisterSrFieldCrceMask  = 0x80
)

// GetCrce CRC Error
func (r *registerSrType) GetCrce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCrceMask) != 0
}

// SetCrce CRC Error
func (r *registerSrType) SetCrce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCrceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCrceMask)
	}
}

const (
	RegisterSrFieldTifreShift = 8
	RegisterSrFieldTifreMask  = 0x100
)

// GetTifre TI frame format error
func (r *registerSrType) GetTifre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTifreMask) != 0
}

// SetTifre TI frame format error
func (r *registerSrType) SetTifre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldTifreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldTifreMask)
	}
}

const (
	RegisterSrFieldModfShift = 9
	RegisterSrFieldModfMask  = 0x200
)

// GetModf Mode Fault
func (r *registerSrType) GetModf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldModfMask) != 0
}

// SetModf Mode Fault
func (r *registerSrType) SetModf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldModfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldModfMask)
	}
}

const (
	RegisterSrFieldTserfShift = 10
	RegisterSrFieldTserfMask  = 0x400
)

// GetTserf Additional number of SPI data to be transacted was reload
func (r *registerSrType) GetTserf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTserfMask) != 0
}

// SetTserf Additional number of SPI data to be transacted was reload
func (r *registerSrType) SetTserf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldTserfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldTserfMask)
	}
}

const (
	RegisterSrFieldSuspShift = 11
	RegisterSrFieldSuspMask  = 0x800
)

// GetSusp SUSPend
func (r *registerSrType) GetSusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSuspMask) != 0
}

// SetSusp SUSPend
func (r *registerSrType) SetSusp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldSuspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldSuspMask)
	}
}

const (
	RegisterSrFieldTxcShift = 12
	RegisterSrFieldTxcMask  = 0x1000
)

// GetTxc TxFIFO transmission complete
func (r *registerSrType) GetTxc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTxcMask) != 0
}

// SetTxc TxFIFO transmission complete
func (r *registerSrType) SetTxc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldTxcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldTxcMask)
	}
}

const (
	RegisterSrFieldRxplvlShift = 13
	RegisterSrFieldRxplvlMask  = 0x6000
)

// GetRxplvl RxFIFO Packing LeVeL
func (r *registerSrType) GetRxplvl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRxplvlMask) >> RegisterSrFieldRxplvlShift)
}

// SetRxplvl RxFIFO Packing LeVeL
func (r *registerSrType) SetRxplvl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldRxplvlMask)|(uint32(value)<<RegisterSrFieldRxplvlShift))
}

const (
	RegisterSrFieldRxwneShift = 15
	RegisterSrFieldRxwneMask  = 0x8000
)

// GetRxwne RxFIFO Word Not Empty
func (r *registerSrType) GetRxwne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRxwneMask) != 0
}

// SetRxwne RxFIFO Word Not Empty
func (r *registerSrType) SetRxwne(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldRxwneMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldRxwneMask)
	}
}

const (
	RegisterSrFieldCtsizeShift = 16
	RegisterSrFieldCtsizeMask  = 0xffff0000
)

// GetCtsize Number of data frames remaining in current TSIZE session
func (r *registerSrType) GetCtsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCtsizeMask) >> RegisterSrFieldCtsizeShift)
}

// SetCtsize Number of data frames remaining in current TSIZE session
func (r *registerSrType) SetCtsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCtsizeMask)|(uint32(value)<<RegisterSrFieldCtsizeShift))
}

// registerIfcrType Interrupt/Status Flags Clear Register
type registerIfcrType uint32

const (
	RegisterIfcrFieldEotcShift = 3
	RegisterIfcrFieldEotcMask  = 0x8
)

// GetEotc End Of Transfer flag clear
func (r *registerIfcrType) GetEotc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldEotcMask) != 0
}

// SetEotc End Of Transfer flag clear
func (r *registerIfcrType) SetEotc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldEotcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldEotcMask)
	}
}

const (
	RegisterIfcrFieldTxtfcShift = 4
	RegisterIfcrFieldTxtfcMask  = 0x10
)

// GetTxtfc Transmission Transfer Filled flag clear
func (r *registerIfcrType) GetTxtfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldTxtfcMask) != 0
}

// SetTxtfc Transmission Transfer Filled flag clear
func (r *registerIfcrType) SetTxtfc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldTxtfcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldTxtfcMask)
	}
}

const (
	RegisterIfcrFieldUdrcShift = 5
	RegisterIfcrFieldUdrcMask  = 0x20
)

// GetUdrc Underrun flag clear
func (r *registerIfcrType) GetUdrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldUdrcMask) != 0
}

// SetUdrc Underrun flag clear
func (r *registerIfcrType) SetUdrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldUdrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldUdrcMask)
	}
}

const (
	RegisterIfcrFieldOvrcShift = 6
	RegisterIfcrFieldOvrcMask  = 0x40
)

// GetOvrc Overrun flag clear
func (r *registerIfcrType) GetOvrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldOvrcMask) != 0
}

// SetOvrc Overrun flag clear
func (r *registerIfcrType) SetOvrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldOvrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldOvrcMask)
	}
}

const (
	RegisterIfcrFieldCrcecShift = 7
	RegisterIfcrFieldCrcecMask  = 0x80
)

// GetCrcec CRC Error flag clear
func (r *registerIfcrType) GetCrcec() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCrcecMask) != 0
}

// SetCrcec CRC Error flag clear
func (r *registerIfcrType) SetCrcec(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldCrcecMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldCrcecMask)
	}
}

const (
	RegisterIfcrFieldTifrecShift = 8
	RegisterIfcrFieldTifrecMask  = 0x100
)

// GetTifrec TI frame format error flag clear
func (r *registerIfcrType) GetTifrec() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldTifrecMask) != 0
}

// SetTifrec TI frame format error flag clear
func (r *registerIfcrType) SetTifrec(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldTifrecMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldTifrecMask)
	}
}

const (
	RegisterIfcrFieldModfcShift = 9
	RegisterIfcrFieldModfcMask  = 0x200
)

// GetModfc Mode Fault flag clear
func (r *registerIfcrType) GetModfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldModfcMask) != 0
}

// SetModfc Mode Fault flag clear
func (r *registerIfcrType) SetModfc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldModfcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldModfcMask)
	}
}

const (
	RegisterIfcrFieldTserfcShift = 10
	RegisterIfcrFieldTserfcMask  = 0x400
)

// GetTserfc TSERFC flag clear
func (r *registerIfcrType) GetTserfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldTserfcMask) != 0
}

// SetTserfc TSERFC flag clear
func (r *registerIfcrType) SetTserfc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldTserfcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldTserfcMask)
	}
}

const (
	RegisterIfcrFieldSuspcShift = 11
	RegisterIfcrFieldSuspcMask  = 0x800
)

// GetSuspc SUSPend flag clear
func (r *registerIfcrType) GetSuspc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldSuspcMask) != 0
}

// SetSuspc SUSPend flag clear
func (r *registerIfcrType) SetSuspc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldSuspcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldSuspcMask)
	}
}

// registerTxdrType Transmit Data Register
type registerTxdrType uint32

const (
	RegisterTxdrFieldTxdrShift = 0
	RegisterTxdrFieldTxdrMask  = 0xffffffff
)

// GetTxdr Transmit data register
func (r *registerTxdrType) GetTxdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxdrFieldTxdrMask) >> RegisterTxdrFieldTxdrShift)
}

// SetTxdr Transmit data register
func (r *registerTxdrType) SetTxdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxdrFieldTxdrMask)|(uint32(value)<<RegisterTxdrFieldTxdrShift))
}

// registerRxdrType Receive Data Register
type registerRxdrType uint32

const (
	RegisterRxdrFieldRxdrShift = 0
	RegisterRxdrFieldRxdrMask  = 0xffffffff
)

// GetRxdr Receive data register
func (r *registerRxdrType) GetRxdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxdrFieldRxdrMask) >> RegisterRxdrFieldRxdrShift)
}

// SetRxdr Receive data register
func (r *registerRxdrType) SetRxdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxdrFieldRxdrMask)|(uint32(value)<<RegisterRxdrFieldRxdrShift))
}

// registerCrcpolyType Polynomial Register
type registerCrcpolyType uint32

const (
	RegisterCrcpolyFieldCrcpolyShift = 0
	RegisterCrcpolyFieldCrcpolyMask  = 0xffffffff
)

// GetCrcpoly CRC polynomial register
func (r *registerCrcpolyType) GetCrcpoly() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrcpolyFieldCrcpolyMask) >> RegisterCrcpolyFieldCrcpolyShift)
}

// SetCrcpoly CRC polynomial register
func (r *registerCrcpolyType) SetCrcpoly(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrcpolyFieldCrcpolyMask)|(uint32(value)<<RegisterCrcpolyFieldCrcpolyShift))
}

// registerTxcrcType Transmitter CRC Register
type registerTxcrcType uint32

const (
	RegisterTxcrcFieldTxcrcShift = 0
	RegisterTxcrcFieldTxcrcMask  = 0xffffffff
)

// GetTxcrc CRC register for transmitter
func (r *registerTxcrcType) GetTxcrc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxcrcFieldTxcrcMask) >> RegisterTxcrcFieldTxcrcShift)
}

// SetTxcrc CRC register for transmitter
func (r *registerTxcrcType) SetTxcrc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxcrcFieldTxcrcMask)|(uint32(value)<<RegisterTxcrcFieldTxcrcShift))
}

// registerRxcrcType Receiver CRC Register
type registerRxcrcType uint32

const (
	RegisterRxcrcFieldRxcrcShift = 0
	RegisterRxcrcFieldRxcrcMask  = 0xffffffff
)

// GetRxcrc CRC register for receiver
func (r *registerRxcrcType) GetRxcrc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxcrcFieldRxcrcMask) >> RegisterRxcrcFieldRxcrcShift)
}

// SetRxcrc CRC register for receiver
func (r *registerRxcrcType) SetRxcrc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxcrcFieldRxcrcMask)|(uint32(value)<<RegisterRxcrcFieldRxcrcShift))
}

// registerUdrdrType Underrun Data Register
type registerUdrdrType uint32

const (
	RegisterUdrdrFieldUdrdrShift = 0
	RegisterUdrdrFieldUdrdrMask  = 0xffffffff
)

// GetUdrdr Data at slave underrun condition
func (r *registerUdrdrType) GetUdrdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterUdrdrFieldUdrdrMask) >> RegisterUdrdrFieldUdrdrShift)
}

// SetUdrdr Data at slave underrun condition
func (r *registerUdrdrType) SetUdrdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUdrdrFieldUdrdrMask)|(uint32(value)<<RegisterUdrdrFieldUdrdrShift))
}

// registerCgfrType configuration register
type registerCgfrType uint32

const (
	RegisterCgfrFieldI2smodShift = 0
	RegisterCgfrFieldI2smodMask  = 0x1
)

// GetI2smod I2S mode selection
func (r *registerCgfrType) GetI2smod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldI2smodMask) != 0
}

// SetI2smod I2S mode selection
func (r *registerCgfrType) SetI2smod(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCgfrFieldI2smodMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldI2smodMask)
	}
}

const (
	RegisterCgfrFieldI2scfgShift = 1
	RegisterCgfrFieldI2scfgMask  = 0xe
)

// GetI2scfg I2S configuration mode
func (r *registerCgfrType) GetI2scfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldI2scfgMask) >> RegisterCgfrFieldI2scfgShift)
}

// SetI2scfg I2S configuration mode
func (r *registerCgfrType) SetI2scfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldI2scfgMask)|(uint32(value)<<RegisterCgfrFieldI2scfgShift))
}

const (
	RegisterCgfrFieldI2sstdShift = 4
	RegisterCgfrFieldI2sstdMask  = 0x30
)

// GetI2sstd I2S standard selection
func (r *registerCgfrType) GetI2sstd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldI2sstdMask) >> RegisterCgfrFieldI2sstdShift)
}

// SetI2sstd I2S standard selection
func (r *registerCgfrType) SetI2sstd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldI2sstdMask)|(uint32(value)<<RegisterCgfrFieldI2sstdShift))
}

const (
	RegisterCgfrFieldPcmsyncShift = 7
	RegisterCgfrFieldPcmsyncMask  = 0x80
)

// GetPcmsync PCM frame synchronization
func (r *registerCgfrType) GetPcmsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldPcmsyncMask) != 0
}

// SetPcmsync PCM frame synchronization
func (r *registerCgfrType) SetPcmsync(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCgfrFieldPcmsyncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldPcmsyncMask)
	}
}

const (
	RegisterCgfrFieldDatlenShift = 8
	RegisterCgfrFieldDatlenMask  = 0x300
)

// GetDatlen Data length to be transferred
func (r *registerCgfrType) GetDatlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldDatlenMask) >> RegisterCgfrFieldDatlenShift)
}

// SetDatlen Data length to be transferred
func (r *registerCgfrType) SetDatlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldDatlenMask)|(uint32(value)<<RegisterCgfrFieldDatlenShift))
}

const (
	RegisterCgfrFieldChlenShift = 10
	RegisterCgfrFieldChlenMask  = 0x400
)

// GetChlen Channel length (number of bits per audio channel)
func (r *registerCgfrType) GetChlen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldChlenMask) != 0
}

// SetChlen Channel length (number of bits per audio channel)
func (r *registerCgfrType) SetChlen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCgfrFieldChlenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldChlenMask)
	}
}

const (
	RegisterCgfrFieldCkpolShift = 11
	RegisterCgfrFieldCkpolMask  = 0x800
)

// GetCkpol Serial audio clock polarity
func (r *registerCgfrType) GetCkpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldCkpolMask) != 0
}

// SetCkpol Serial audio clock polarity
func (r *registerCgfrType) SetCkpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCgfrFieldCkpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldCkpolMask)
	}
}

const (
	RegisterCgfrFieldFixchShift = 12
	RegisterCgfrFieldFixchMask  = 0x1000
)

// GetFixch Word select inversion
func (r *registerCgfrType) GetFixch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldFixchMask) != 0
}

// SetFixch Word select inversion
func (r *registerCgfrType) SetFixch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCgfrFieldFixchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldFixchMask)
	}
}

const (
	RegisterCgfrFieldWsinvShift = 13
	RegisterCgfrFieldWsinvMask  = 0x2000
)

// GetWsinv Fixed channel length in SLAVE
func (r *registerCgfrType) GetWsinv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldWsinvMask) != 0
}

// SetWsinv Fixed channel length in SLAVE
func (r *registerCgfrType) SetWsinv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCgfrFieldWsinvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldWsinvMask)
	}
}

const (
	RegisterCgfrFieldDatfmtShift = 14
	RegisterCgfrFieldDatfmtMask  = 0x4000
)

// GetDatfmt Data format
func (r *registerCgfrType) GetDatfmt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldDatfmtMask) != 0
}

// SetDatfmt Data format
func (r *registerCgfrType) SetDatfmt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCgfrFieldDatfmtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldDatfmtMask)
	}
}

const (
	RegisterCgfrFieldI2sdivShift = 16
	RegisterCgfrFieldI2sdivMask  = 0xff0000
)

// GetI2sdiv I2S linear prescaler
func (r *registerCgfrType) GetI2sdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldI2sdivMask) >> RegisterCgfrFieldI2sdivShift)
}

// SetI2sdiv I2S linear prescaler
func (r *registerCgfrType) SetI2sdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldI2sdivMask)|(uint32(value)<<RegisterCgfrFieldI2sdivShift))
}

const (
	RegisterCgfrFieldOddShift = 24
	RegisterCgfrFieldOddMask  = 0x1000000
)

// GetOdd Odd factor for the prescaler
func (r *registerCgfrType) GetOdd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldOddMask) != 0
}

// SetOdd Odd factor for the prescaler
func (r *registerCgfrType) SetOdd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCgfrFieldOddMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldOddMask)
	}
}

const (
	RegisterCgfrFieldMckoeShift = 25
	RegisterCgfrFieldMckoeMask  = 0x2000000
)

// GetMckoe Master clock output enable
func (r *registerCgfrType) GetMckoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldMckoeMask) != 0
}

// SetMckoe Master clock output enable
func (r *registerCgfrType) SetMckoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCgfrFieldMckoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldMckoeMask)
	}
}
