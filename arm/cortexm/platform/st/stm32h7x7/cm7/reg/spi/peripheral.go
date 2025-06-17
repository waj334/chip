//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

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
	Cr1     RegisterCr1Type
	Cr2     RegisterCr2Type
	Cfg1    RegisterCfg1Type
	Cfg2    RegisterCfg2Type
	Ier     RegisterIerType
	Sr      RegisterSrType
	Ifcr    RegisterIfcrType
	_       [4]byte
	Txdr    RegisterTxdrType
	_       [12]byte
	Rxdr    RegisterRxdrType
	_       [12]byte
	Crcpoly RegisterCrcpolyType
	Txcrc   RegisterTxcrcType
	Rxcrc   RegisterRxcrcType
	Udrdr   RegisterUdrdrType
	Cgfr    RegisterCgfrType
}

// RegisterCr1Type control register 1
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
	RegisterCr1FieldSpeShift = 0
	RegisterCr1FieldSpeMask  = 0x1
)

// GetSpe Serial Peripheral Enable
func (r *RegisterCr1Type) GetSpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldSpeMask) != 0
}

// SetSpe Serial Peripheral Enable
func (r *RegisterCr1Type) SetSpe(value bool) {
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
func (r *RegisterCr1Type) GetMasrx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldMasrxMask) != 0
}

// SetMasrx Master automatic SUSP in Receive mode
func (r *RegisterCr1Type) SetMasrx(value bool) {
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
func (r *RegisterCr1Type) GetCstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCstartMask) != 0
}

const (
	RegisterCr1FieldCsuspShift = 10
	RegisterCr1FieldCsuspMask  = 0x400
)

// SetCsusp Master SUSPend request
func (r *RegisterCr1Type) SetCsusp(value bool) {
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
func (r *RegisterCr1Type) GetHddir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldHddirMask) != 0
}

// SetHddir Rx/Tx direction at Half-duplex mode
func (r *RegisterCr1Type) SetHddir(value bool) {
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
func (r *RegisterCr1Type) GetSsi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldSsiMask) != 0
}

// SetSsi Internal SS signal input level
func (r *RegisterCr1Type) SetSsi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldSsiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldSsiMask)
	}
}

const (
	RegisterCr1FieldCrc3317Shift = 13
	RegisterCr1FieldCrc3317Mask  = 0x2000
)

// GetCrc3317 32-bit CRC polynomial configuration
func (r *RegisterCr1Type) GetCrc3317() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldCrc3317Mask) != 0
}

// SetCrc3317 32-bit CRC polynomial configuration
func (r *RegisterCr1Type) SetCrc3317(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCr1FieldCrc3317Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCr1FieldCrc3317Mask)
	}
}

const (
	RegisterCr1FieldRcrciShift = 14
	RegisterCr1FieldRcrciMask  = 0x4000
)

// GetRcrci CRC calculation initialization pattern control for receiver
func (r *RegisterCr1Type) GetRcrci() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldRcrciMask) != 0
}

// SetRcrci CRC calculation initialization pattern control for receiver
func (r *RegisterCr1Type) SetRcrci(value bool) {
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
func (r *RegisterCr1Type) GetTcrci() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldTcrciMask) != 0
}

// SetTcrci CRC calculation initialization pattern control for transmitter
func (r *RegisterCr1Type) SetTcrci(value bool) {
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
func (r *RegisterCr1Type) GetIolock() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCr1FieldIolockMask) != 0
}

// RegisterCr2Type control register 2
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
	RegisterCr2FieldTsizeShift = 0
	RegisterCr2FieldTsizeMask  = 0xffff
)

// GetTsize Number of data at current transfer
func (r *RegisterCr2Type) GetTsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTsizeMask) >> RegisterCr2FieldTsizeShift)
}

// SetTsize Number of data at current transfer
func (r *RegisterCr2Type) SetTsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCr2FieldTsizeMask)|(uint32(value)<<RegisterCr2FieldTsizeShift))
}

const (
	RegisterCr2FieldTserShift = 16
	RegisterCr2FieldTserMask  = 0xffff0000
)

// GetTser Number of data transfer extension to be reload into TSIZE just when a previous
func (r *RegisterCr2Type) GetTser() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCr2FieldTserMask) >> RegisterCr2FieldTserShift)
}

// RegisterCfg1Type configuration register 1
type RegisterCfg1Type uint32

func (r *RegisterCfg1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfg1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfg1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfg1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfg1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCfg1FieldDsizeShift = 0
	RegisterCfg1FieldDsizeMask  = 0x1f
)

// GetDsize Number of bits in at single SPI data frame
func (r *RegisterCfg1Type) GetDsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldDsizeMask) >> RegisterCfg1FieldDsizeShift)
}

// SetDsize Number of bits in at single SPI data frame
func (r *RegisterCfg1Type) SetDsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldDsizeMask)|(uint32(value)<<RegisterCfg1FieldDsizeShift))
}

const (
	RegisterCfg1FieldFthvlShift = 5
	RegisterCfg1FieldFthvlMask  = 0x1e0
)

// GetFthvl threshold level
func (r *RegisterCfg1Type) GetFthvl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldFthvlMask) >> RegisterCfg1FieldFthvlShift)
}

// SetFthvl threshold level
func (r *RegisterCfg1Type) SetFthvl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldFthvlMask)|(uint32(value)<<RegisterCfg1FieldFthvlShift))
}

const (
	RegisterCfg1FieldUdrcfgShift = 9
	RegisterCfg1FieldUdrcfgMask  = 0x600
)

// GetUdrcfg Behavior of slave transmitter at underrun condition
func (r *RegisterCfg1Type) GetUdrcfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldUdrcfgMask) >> RegisterCfg1FieldUdrcfgShift)
}

// SetUdrcfg Behavior of slave transmitter at underrun condition
func (r *RegisterCfg1Type) SetUdrcfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldUdrcfgMask)|(uint32(value)<<RegisterCfg1FieldUdrcfgShift))
}

const (
	RegisterCfg1FieldUdrdetShift = 11
	RegisterCfg1FieldUdrdetMask  = 0x1800
)

// GetUdrdet Detection of underrun condition at slave transmitter
func (r *RegisterCfg1Type) GetUdrdet() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldUdrdetMask) >> RegisterCfg1FieldUdrdetShift)
}

// SetUdrdet Detection of underrun condition at slave transmitter
func (r *RegisterCfg1Type) SetUdrdet(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldUdrdetMask)|(uint32(value)<<RegisterCfg1FieldUdrdetShift))
}

const (
	RegisterCfg1FieldRxdmaenShift = 14
	RegisterCfg1FieldRxdmaenMask  = 0x4000
)

// GetRxdmaen Rx DMA stream enable
func (r *RegisterCfg1Type) GetRxdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldRxdmaenMask) != 0
}

// SetRxdmaen Rx DMA stream enable
func (r *RegisterCfg1Type) SetRxdmaen(value bool) {
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
func (r *RegisterCfg1Type) GetTxdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldTxdmaenMask) != 0
}

// SetTxdmaen Tx DMA stream enable
func (r *RegisterCfg1Type) SetTxdmaen(value bool) {
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
func (r *RegisterCfg1Type) GetCrcsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldCrcsizeMask) >> RegisterCfg1FieldCrcsizeShift)
}

// SetCrcsize Length of CRC frame to be transacted and compared
func (r *RegisterCfg1Type) SetCrcsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldCrcsizeMask)|(uint32(value)<<RegisterCfg1FieldCrcsizeShift))
}

const (
	RegisterCfg1FieldCrcenShift = 22
	RegisterCfg1FieldCrcenMask  = 0x400000
)

// GetCrcen Hardware CRC computation enable
func (r *RegisterCfg1Type) GetCrcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldCrcenMask) != 0
}

// SetCrcen Hardware CRC computation enable
func (r *RegisterCfg1Type) SetCrcen(value bool) {
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
func (r *RegisterCfg1Type) GetMbr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg1FieldMbrMask) >> RegisterCfg1FieldMbrShift)
}

// SetMbr Master baud rate
func (r *RegisterCfg1Type) SetMbr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg1FieldMbrMask)|(uint32(value)<<RegisterCfg1FieldMbrShift))
}

// RegisterCfg2Type configuration register 2
type RegisterCfg2Type uint32

func (r *RegisterCfg2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfg2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfg2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfg2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfg2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCfg2FieldMssiShift = 0
	RegisterCfg2FieldMssiMask  = 0xf
)

// GetMssi Master SS Idleness
func (r *RegisterCfg2Type) GetMssi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldMssiMask) >> RegisterCfg2FieldMssiShift)
}

// SetMssi Master SS Idleness
func (r *RegisterCfg2Type) SetMssi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldMssiMask)|(uint32(value)<<RegisterCfg2FieldMssiShift))
}

const (
	RegisterCfg2FieldMidiShift = 4
	RegisterCfg2FieldMidiMask  = 0xf0
)

// GetMidi Master Inter-Data Idleness
func (r *RegisterCfg2Type) GetMidi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldMidiMask) >> RegisterCfg2FieldMidiShift)
}

// SetMidi Master Inter-Data Idleness
func (r *RegisterCfg2Type) SetMidi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldMidiMask)|(uint32(value)<<RegisterCfg2FieldMidiShift))
}

const (
	RegisterCfg2FieldIoswpShift = 15
	RegisterCfg2FieldIoswpMask  = 0x8000
)

// GetIoswp Swap functionality of MISO and MOSI pins
func (r *RegisterCfg2Type) GetIoswp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldIoswpMask) != 0
}

// SetIoswp Swap functionality of MISO and MOSI pins
func (r *RegisterCfg2Type) SetIoswp(value bool) {
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
func (r *RegisterCfg2Type) GetComm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldCommMask) >> RegisterCfg2FieldCommShift)
}

// SetComm SPI Communication Mode
func (r *RegisterCfg2Type) SetComm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldCommMask)|(uint32(value)<<RegisterCfg2FieldCommShift))
}

const (
	RegisterCfg2FieldSpShift = 19
	RegisterCfg2FieldSpMask  = 0x380000
)

// GetSp Serial Protocol
func (r *RegisterCfg2Type) GetSp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldSpMask) >> RegisterCfg2FieldSpShift)
}

// SetSp Serial Protocol
func (r *RegisterCfg2Type) SetSp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldSpMask)|(uint32(value)<<RegisterCfg2FieldSpShift))
}

const (
	RegisterCfg2FieldMasterShift = 22
	RegisterCfg2FieldMasterMask  = 0x400000
)

// GetMaster SPI Master
func (r *RegisterCfg2Type) GetMaster() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldMasterMask) != 0
}

// SetMaster SPI Master
func (r *RegisterCfg2Type) SetMaster(value bool) {
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
func (r *RegisterCfg2Type) GetLsbfrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldLsbfrstMask) != 0
}

// SetLsbfrst Data frame format
func (r *RegisterCfg2Type) SetLsbfrst(value bool) {
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
func (r *RegisterCfg2Type) GetCpha() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldCphaMask) != 0
}

// SetCpha Clock phase
func (r *RegisterCfg2Type) SetCpha(value bool) {
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
func (r *RegisterCfg2Type) GetCpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldCpolMask) != 0
}

// SetCpol Clock polarity
func (r *RegisterCfg2Type) SetCpol(value bool) {
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
func (r *RegisterCfg2Type) GetSsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldSsmMask) != 0
}

// SetSsm Software management of SS signal input
func (r *RegisterCfg2Type) SetSsm(value bool) {
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
func (r *RegisterCfg2Type) GetSsiop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldSsiopMask) != 0
}

// SetSsiop SS input/output polarity
func (r *RegisterCfg2Type) SetSsiop(value bool) {
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
func (r *RegisterCfg2Type) GetSsoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldSsoeMask) != 0
}

// SetSsoe SS output enable
func (r *RegisterCfg2Type) SetSsoe(value bool) {
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
func (r *RegisterCfg2Type) GetSsom() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldSsomMask) != 0
}

// SetSsom SS output management in master mode
func (r *RegisterCfg2Type) SetSsom(value bool) {
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
func (r *RegisterCfg2Type) GetAfcntr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfg2FieldAfcntrMask) != 0
}

// SetAfcntr Alternate function GPIOs control
func (r *RegisterCfg2Type) SetAfcntr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfg2FieldAfcntrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfg2FieldAfcntrMask)
	}
}

// RegisterIerType Interrupt Enable Register
type RegisterIerType uint32

func (r *RegisterIerType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIerType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIerType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIerType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIerType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIerFieldRxpieShift = 0
	RegisterIerFieldRxpieMask  = 0x1
)

// GetRxpie RXP Interrupt Enable
func (r *RegisterIerType) GetRxpie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRxpieMask) != 0
}

// SetRxpie RXP Interrupt Enable
func (r *RegisterIerType) SetRxpie(value bool) {
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
func (r *RegisterIerType) GetTxpie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTxpieMask) != 0
}

const (
	RegisterIerFieldDpxpieShift = 2
	RegisterIerFieldDpxpieMask  = 0x4
)

// GetDpxpie DXP interrupt enabled
func (r *RegisterIerType) GetDpxpie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldDpxpieMask) != 0
}

const (
	RegisterIerFieldEotieShift = 3
	RegisterIerFieldEotieMask  = 0x8
)

// GetEotie EOT, SUSP and TXC interrupt enable
func (r *RegisterIerType) GetEotie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldEotieMask) != 0
}

// SetEotie EOT, SUSP and TXC interrupt enable
func (r *RegisterIerType) SetEotie(value bool) {
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
func (r *RegisterIerType) GetTxtfie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTxtfieMask) != 0
}

// SetTxtfie TXTFIE interrupt enable
func (r *RegisterIerType) SetTxtfie(value bool) {
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
func (r *RegisterIerType) GetUdrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldUdrieMask) != 0
}

// SetUdrie UDR interrupt enable
func (r *RegisterIerType) SetUdrie(value bool) {
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
func (r *RegisterIerType) GetOvrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldOvrieMask) != 0
}

// SetOvrie OVR interrupt enable
func (r *RegisterIerType) SetOvrie(value bool) {
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
func (r *RegisterIerType) GetCrceie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldCrceieMask) != 0
}

// SetCrceie CRC Interrupt enable
func (r *RegisterIerType) SetCrceie(value bool) {
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
func (r *RegisterIerType) GetTifreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTifreieMask) != 0
}

// SetTifreie TIFRE interrupt enable
func (r *RegisterIerType) SetTifreie(value bool) {
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
func (r *RegisterIerType) GetModfie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldModfieMask) != 0
}

// SetModfie Mode Fault interrupt enable
func (r *RegisterIerType) SetModfie(value bool) {
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
func (r *RegisterIerType) GetTserfie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTserfieMask) != 0
}

// SetTserfie Additional number of transactions reload interrupt enable
func (r *RegisterIerType) SetTserfie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTserfieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTserfieMask)
	}
}

// RegisterSrType Status Register
type RegisterSrType uint32

func (r *RegisterSrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSrFieldRxpShift = 0
	RegisterSrFieldRxpMask  = 0x1
)

// GetRxp Rx-Packet available
func (r *RegisterSrType) GetRxp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRxpMask) != 0
}

// SetRxp Rx-Packet available
func (r *RegisterSrType) SetRxp(value bool) {
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
func (r *RegisterSrType) GetTxp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTxpMask) != 0
}

// SetTxp Tx-Packet space available
func (r *RegisterSrType) SetTxp(value bool) {
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
func (r *RegisterSrType) GetDxp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldDxpMask) != 0
}

// SetDxp Duplex Packet
func (r *RegisterSrType) SetDxp(value bool) {
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
func (r *RegisterSrType) GetEot() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldEotMask) != 0
}

// SetEot End Of Transfer
func (r *RegisterSrType) SetEot(value bool) {
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
func (r *RegisterSrType) GetTxtf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTxtfMask) != 0
}

// SetTxtf Transmission Transfer Filled
func (r *RegisterSrType) SetTxtf(value bool) {
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
func (r *RegisterSrType) GetUdr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldUdrMask) != 0
}

// SetUdr Underrun at slave transmission mode
func (r *RegisterSrType) SetUdr(value bool) {
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
func (r *RegisterSrType) GetOvr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldOvrMask) != 0
}

// SetOvr Overrun
func (r *RegisterSrType) SetOvr(value bool) {
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
func (r *RegisterSrType) GetCrce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCrceMask) != 0
}

// SetCrce CRC Error
func (r *RegisterSrType) SetCrce(value bool) {
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
func (r *RegisterSrType) GetTifre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTifreMask) != 0
}

// SetTifre TI frame format error
func (r *RegisterSrType) SetTifre(value bool) {
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
func (r *RegisterSrType) GetModf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldModfMask) != 0
}

// SetModf Mode Fault
func (r *RegisterSrType) SetModf(value bool) {
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
func (r *RegisterSrType) GetTserf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTserfMask) != 0
}

// SetTserf Additional number of SPI data to be transacted was reload
func (r *RegisterSrType) SetTserf(value bool) {
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
func (r *RegisterSrType) GetSusp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldSuspMask) != 0
}

// SetSusp SUSPend
func (r *RegisterSrType) SetSusp(value bool) {
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
func (r *RegisterSrType) GetTxc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldTxcMask) != 0
}

// SetTxc TxFIFO transmission complete
func (r *RegisterSrType) SetTxc(value bool) {
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
func (r *RegisterSrType) GetRxplvl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRxplvlMask) >> RegisterSrFieldRxplvlShift)
}

// SetRxplvl RxFIFO Packing LeVeL
func (r *RegisterSrType) SetRxplvl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldRxplvlMask)|(uint32(value)<<RegisterSrFieldRxplvlShift))
}

const (
	RegisterSrFieldRxwneShift = 15
	RegisterSrFieldRxwneMask  = 0x8000
)

// GetRxwne RxFIFO Word Not Empty
func (r *RegisterSrType) GetRxwne() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRxwneMask) != 0
}

// SetRxwne RxFIFO Word Not Empty
func (r *RegisterSrType) SetRxwne(value bool) {
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
func (r *RegisterSrType) GetCtsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCtsizeMask) >> RegisterSrFieldCtsizeShift)
}

// SetCtsize Number of data frames remaining in current TSIZE session
func (r *RegisterSrType) SetCtsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCtsizeMask)|(uint32(value)<<RegisterSrFieldCtsizeShift))
}

// RegisterIfcrType Interrupt/Status Flags Clear Register
type RegisterIfcrType uint32

func (r *RegisterIfcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterIfcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterIfcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterIfcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterIfcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterIfcrFieldEotcShift = 3
	RegisterIfcrFieldEotcMask  = 0x8
)

// GetEotc End Of Transfer flag clear
func (r *RegisterIfcrType) GetEotc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldEotcMask) != 0
}

// SetEotc End Of Transfer flag clear
func (r *RegisterIfcrType) SetEotc(value bool) {
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
func (r *RegisterIfcrType) GetTxtfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldTxtfcMask) != 0
}

// SetTxtfc Transmission Transfer Filled flag clear
func (r *RegisterIfcrType) SetTxtfc(value bool) {
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
func (r *RegisterIfcrType) GetUdrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldUdrcMask) != 0
}

// SetUdrc Underrun flag clear
func (r *RegisterIfcrType) SetUdrc(value bool) {
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
func (r *RegisterIfcrType) GetOvrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldOvrcMask) != 0
}

// SetOvrc Overrun flag clear
func (r *RegisterIfcrType) SetOvrc(value bool) {
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
func (r *RegisterIfcrType) GetCrcec() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldCrcecMask) != 0
}

// SetCrcec CRC Error flag clear
func (r *RegisterIfcrType) SetCrcec(value bool) {
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
func (r *RegisterIfcrType) GetTifrec() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldTifrecMask) != 0
}

// SetTifrec TI frame format error flag clear
func (r *RegisterIfcrType) SetTifrec(value bool) {
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
func (r *RegisterIfcrType) GetModfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldModfcMask) != 0
}

// SetModfc Mode Fault flag clear
func (r *RegisterIfcrType) SetModfc(value bool) {
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
func (r *RegisterIfcrType) GetTserfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldTserfcMask) != 0
}

// SetTserfc TSERFC flag clear
func (r *RegisterIfcrType) SetTserfc(value bool) {
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
func (r *RegisterIfcrType) GetSuspc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIfcrFieldSuspcMask) != 0
}

// SetSuspc SUSPend flag clear
func (r *RegisterIfcrType) SetSuspc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIfcrFieldSuspcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIfcrFieldSuspcMask)
	}
}

// RegisterTxdrType Transmit Data Register
type RegisterTxdrType uint32

func (r *RegisterTxdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTxdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTxdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTxdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTxdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTxdrFieldTxdrShift = 0
	RegisterTxdrFieldTxdrMask  = 0xffffffff
)

// GetTxdr Transmit data register
func (r *RegisterTxdrType) GetTxdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxdrFieldTxdrMask) >> RegisterTxdrFieldTxdrShift)
}

// SetTxdr Transmit data register
func (r *RegisterTxdrType) SetTxdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxdrFieldTxdrMask)|(uint32(value)<<RegisterTxdrFieldTxdrShift))
}

// RegisterRxdrType Receive Data Register
type RegisterRxdrType uint32

func (r *RegisterRxdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRxdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRxdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRxdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRxdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRxdrFieldRxdrShift = 0
	RegisterRxdrFieldRxdrMask  = 0xffffffff
)

// GetRxdr Receive data register
func (r *RegisterRxdrType) GetRxdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxdrFieldRxdrMask) >> RegisterRxdrFieldRxdrShift)
}

// SetRxdr Receive data register
func (r *RegisterRxdrType) SetRxdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxdrFieldRxdrMask)|(uint32(value)<<RegisterRxdrFieldRxdrShift))
}

// RegisterCrcpolyType Polynomial Register
type RegisterCrcpolyType uint32

func (r *RegisterCrcpolyType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrcpolyType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrcpolyType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrcpolyType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrcpolyType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrcpolyFieldCrcpolyShift = 0
	RegisterCrcpolyFieldCrcpolyMask  = 0xffffffff
)

// GetCrcpoly CRC polynomial register
func (r *RegisterCrcpolyType) GetCrcpoly() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCrcpolyFieldCrcpolyMask) >> RegisterCrcpolyFieldCrcpolyShift)
}

// SetCrcpoly CRC polynomial register
func (r *RegisterCrcpolyType) SetCrcpoly(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrcpolyFieldCrcpolyMask)|(uint32(value)<<RegisterCrcpolyFieldCrcpolyShift))
}

// RegisterTxcrcType Transmitter CRC Register
type RegisterTxcrcType uint32

func (r *RegisterTxcrcType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTxcrcType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTxcrcType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTxcrcType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTxcrcType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTxcrcFieldTxcrcShift = 0
	RegisterTxcrcFieldTxcrcMask  = 0xffffffff
)

// GetTxcrc CRC register for transmitter
func (r *RegisterTxcrcType) GetTxcrc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxcrcFieldTxcrcMask) >> RegisterTxcrcFieldTxcrcShift)
}

// SetTxcrc CRC register for transmitter
func (r *RegisterTxcrcType) SetTxcrc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxcrcFieldTxcrcMask)|(uint32(value)<<RegisterTxcrcFieldTxcrcShift))
}

// RegisterRxcrcType Receiver CRC Register
type RegisterRxcrcType uint32

func (r *RegisterRxcrcType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRxcrcType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRxcrcType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRxcrcType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRxcrcType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRxcrcFieldRxcrcShift = 0
	RegisterRxcrcFieldRxcrcMask  = 0xffffffff
)

// GetRxcrc CRC register for receiver
func (r *RegisterRxcrcType) GetRxcrc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxcrcFieldRxcrcMask) >> RegisterRxcrcFieldRxcrcShift)
}

// SetRxcrc CRC register for receiver
func (r *RegisterRxcrcType) SetRxcrc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxcrcFieldRxcrcMask)|(uint32(value)<<RegisterRxcrcFieldRxcrcShift))
}

// RegisterUdrdrType Underrun Data Register
type RegisterUdrdrType uint32

func (r *RegisterUdrdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterUdrdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterUdrdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterUdrdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterUdrdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterUdrdrFieldUdrdrShift = 0
	RegisterUdrdrFieldUdrdrMask  = 0xffffffff
)

// GetUdrdr Data at slave underrun condition
func (r *RegisterUdrdrType) GetUdrdr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterUdrdrFieldUdrdrMask) >> RegisterUdrdrFieldUdrdrShift)
}

// SetUdrdr Data at slave underrun condition
func (r *RegisterUdrdrType) SetUdrdr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterUdrdrFieldUdrdrMask)|(uint32(value)<<RegisterUdrdrFieldUdrdrShift))
}

// RegisterCgfrType configuration register
type RegisterCgfrType uint32

func (r *RegisterCgfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCgfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCgfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCgfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCgfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCgfrFieldI2smodShift = 0
	RegisterCgfrFieldI2smodMask  = 0x1
)

// GetI2smod I2S mode selection
func (r *RegisterCgfrType) GetI2smod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldI2smodMask) != 0
}

// SetI2smod I2S mode selection
func (r *RegisterCgfrType) SetI2smod(value bool) {
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
func (r *RegisterCgfrType) GetI2scfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldI2scfgMask) >> RegisterCgfrFieldI2scfgShift)
}

// SetI2scfg I2S configuration mode
func (r *RegisterCgfrType) SetI2scfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldI2scfgMask)|(uint32(value)<<RegisterCgfrFieldI2scfgShift))
}

const (
	RegisterCgfrFieldI2sstdShift = 4
	RegisterCgfrFieldI2sstdMask  = 0x30
)

// GetI2sstd I2S standard selection
func (r *RegisterCgfrType) GetI2sstd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldI2sstdMask) >> RegisterCgfrFieldI2sstdShift)
}

// SetI2sstd I2S standard selection
func (r *RegisterCgfrType) SetI2sstd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldI2sstdMask)|(uint32(value)<<RegisterCgfrFieldI2sstdShift))
}

const (
	RegisterCgfrFieldPcmsyncShift = 7
	RegisterCgfrFieldPcmsyncMask  = 0x80
)

// GetPcmsync PCM frame synchronization
func (r *RegisterCgfrType) GetPcmsync() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldPcmsyncMask) != 0
}

// SetPcmsync PCM frame synchronization
func (r *RegisterCgfrType) SetPcmsync(value bool) {
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
func (r *RegisterCgfrType) GetDatlen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldDatlenMask) >> RegisterCgfrFieldDatlenShift)
}

// SetDatlen Data length to be transferred
func (r *RegisterCgfrType) SetDatlen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldDatlenMask)|(uint32(value)<<RegisterCgfrFieldDatlenShift))
}

const (
	RegisterCgfrFieldChlenShift = 10
	RegisterCgfrFieldChlenMask  = 0x400
)

// GetChlen Channel length (number of bits per audio channel)
func (r *RegisterCgfrType) GetChlen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldChlenMask) != 0
}

// SetChlen Channel length (number of bits per audio channel)
func (r *RegisterCgfrType) SetChlen(value bool) {
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
func (r *RegisterCgfrType) GetCkpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldCkpolMask) != 0
}

// SetCkpol Serial audio clock polarity
func (r *RegisterCgfrType) SetCkpol(value bool) {
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
func (r *RegisterCgfrType) GetFixch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldFixchMask) != 0
}

// SetFixch Word select inversion
func (r *RegisterCgfrType) SetFixch(value bool) {
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
func (r *RegisterCgfrType) GetWsinv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldWsinvMask) != 0
}

// SetWsinv Fixed channel length in SLAVE
func (r *RegisterCgfrType) SetWsinv(value bool) {
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
func (r *RegisterCgfrType) GetDatfmt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldDatfmtMask) != 0
}

// SetDatfmt Data format
func (r *RegisterCgfrType) SetDatfmt(value bool) {
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
func (r *RegisterCgfrType) GetI2sdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldI2sdivMask) >> RegisterCgfrFieldI2sdivShift)
}

// SetI2sdiv I2S linear prescaler
func (r *RegisterCgfrType) SetI2sdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldI2sdivMask)|(uint32(value)<<RegisterCgfrFieldI2sdivShift))
}

const (
	RegisterCgfrFieldOddShift = 24
	RegisterCgfrFieldOddMask  = 0x1000000
)

// GetOdd Odd factor for the prescaler
func (r *RegisterCgfrType) GetOdd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldOddMask) != 0
}

// SetOdd Odd factor for the prescaler
func (r *RegisterCgfrType) SetOdd(value bool) {
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
func (r *RegisterCgfrType) GetMckoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCgfrFieldMckoeMask) != 0
}

// SetMckoe Master clock output enable
func (r *RegisterCgfrType) SetMckoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCgfrFieldMckoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCgfrFieldMckoeMask)
	}
}
