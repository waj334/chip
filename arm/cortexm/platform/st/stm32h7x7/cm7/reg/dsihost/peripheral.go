//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package dsihost

import (
	"unsafe"
	"volatile"
)

var (
	Dsihost = (*_dsihost)(unsafe.Pointer(uintptr(0x50000000)))
)

type _dsihost struct {
	Dsivr      RegisterDsivrType
	Dsicr      RegisterDsicrType
	Dsiccr     RegisterDsiccrType
	Dsilvcidr  RegisterDsilvcidrType
	Dsilcolcr  RegisterDsilcolcrType
	Dsilpcr    RegisterDsilpcrType
	Dsilpmcr   RegisterDsilpmcrType
	_          [16]byte
	Dsipcr     RegisterDsipcrType
	Dsigvcidr  RegisterDsigvcidrType
	Dsimcr     RegisterDsimcrType
	Dsivmcr    RegisterDsivmcrType
	Dsivpcr    RegisterDsivpcrType
	Dsivccr    RegisterDsivccrType
	Dsivnpcr   RegisterDsivnpcrType
	Dsivhsacr  RegisterDsivhsacrType
	Dsivhbpcr  RegisterDsivhbpcrType
	Dsivlcr    RegisterDsivlcrType
	Dsivvsacr  RegisterDsivvsacrType
	Dsivvbpcr  RegisterDsivvbpcrType
	Dsivvfpcr  RegisterDsivvfpcrType
	Dsivvacr   RegisterDsivvacrType
	Dsilccr    RegisterDsilccrType
	Dsicmcr    RegisterDsicmcrType
	Dsighcr    RegisterDsighcrType
	Dsigpdr    RegisterDsigpdrType
	Dsigpsr    RegisterDsigpsrType
	Dsitccr0   RegisterDsitccr0Type
	Dsitccr1   RegisterDsitccr1Type
	Dsitccr2   RegisterDsitccr2Type
	Dsitccr3   RegisterDsitccr3Type
	Dsitccr4   RegisterDsitccr4Type
	Dsitccr5   RegisterDsitccr5Type
	_          [4]byte
	Dsiclcr    RegisterDsiclcrType
	Dsicltcr   RegisterDsicltcrType
	Dsidltcr   RegisterDsidltcrType
	Dsipctlr   RegisterDsipctlrType
	Dsipconfr  RegisterDsipconfrType
	Dsipucr    RegisterDsipucrType
	Dsipttcr   RegisterDsipttcrType
	Dsipsr     RegisterDsipsrType
	_          [8]byte
	Dsiisr0    RegisterDsiisr0Type
	Dsiisr1    RegisterDsiisr1Type
	Dsiier0    RegisterDsiier0Type
	Dsiier1    RegisterDsiier1Type
	_          [12]byte
	Dsifir0    RegisterDsifir0Type
	Dsifir1    RegisterDsifir1Type
	_          [32]byte
	Dsivscr    RegisterDsivscrType
	_          [8]byte
	Dsilcvcidr RegisterDsilcvcidrType
	Dsilcccr   RegisterDsilcccrType
	_          [4]byte
	Dsilpmccr  RegisterDsilpmccrType
	_          [28]byte
	Dsivmccr   RegisterDsivmccrType
	Dsivpccr   RegisterDsivpccrType
	Dsivcccr   RegisterDsivcccrType
	Dsivnpccr  RegisterDsivnpccrType
	Dsivhsaccr RegisterDsivhsaccrType
	Dsivhbpccr RegisterDsivhbpccrType
	Dsivlccr   RegisterDsivlccrType
	Dsivvsaccr RegisterDsivvsaccrType
	Dsivvbpccr RegisterDsivvbpccrType
	Dsivvfpccr RegisterDsivvfpccrType
	Dsivvaccr  RegisterDsivvaccrType
	_          [668]byte
	Dsiwcfgr   RegisterDsiwcfgrType
	Dsiwcr     RegisterDsiwcrType
	Dsiwier    RegisterDsiwierType
	Dsiwisr    RegisterDsiwisrType
	Dsiwifcr   RegisterDsiwifcrType
	_          [4]byte
	Dsiwpcr0   RegisterDsiwpcr0Type
	Dsiwpcr1   RegisterDsiwpcr1Type
	Dsiwpcr2   RegisterDsiwpcr2Type
	Dsiwpcr3   RegisterDsiwpcr3Type
	Dsiwpcr4   RegisterDsiwpcr4Type
	_          [4]byte
	Dsiwrpcr   RegisterDsiwrpcrType
}

// RegisterDsivrType DSI Host version register
type RegisterDsivrType uint32

func (r *RegisterDsivrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivrFieldVersionShift = 0
	RegisterDsivrFieldVersionMask  = 0xffffffff
)

// GetVersion VERSION
func (r *RegisterDsivrType) GetVersion() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDsivrFieldVersionMask) >> RegisterDsivrFieldVersionShift)
}

// SetVersion VERSION
func (r *RegisterDsivrType) SetVersion(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivrFieldVersionMask)|(uint32(value)<<RegisterDsivrFieldVersionShift))
}

// RegisterDsicrType DSI Host control register
type RegisterDsicrType uint32

func (r *RegisterDsicrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsicrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsicrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsicrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsicrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsicrFieldEnShift = 0
	RegisterDsicrFieldEnMask  = 0x1
)

// GetEn EN
func (r *RegisterDsicrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicrFieldEnMask) != 0
}

// SetEn EN
func (r *RegisterDsicrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicrFieldEnMask)
	}
}

// RegisterDsiccrType DSI Host clock control register
type RegisterDsiccrType uint32

func (r *RegisterDsiccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiccrFieldTxeckdivShift = 0
	RegisterDsiccrFieldTxeckdivMask  = 0xff
)

// GetTxeckdiv TXECKDIV
func (r *RegisterDsiccrType) GetTxeckdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiccrFieldTxeckdivMask) >> RegisterDsiccrFieldTxeckdivShift)
}

// SetTxeckdiv TXECKDIV
func (r *RegisterDsiccrType) SetTxeckdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiccrFieldTxeckdivMask)|(uint32(value)<<RegisterDsiccrFieldTxeckdivShift))
}

const (
	RegisterDsiccrFieldTockdivShift = 8
	RegisterDsiccrFieldTockdivMask  = 0xff00
)

// GetTockdiv TOCKDIV
func (r *RegisterDsiccrType) GetTockdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiccrFieldTockdivMask) >> RegisterDsiccrFieldTockdivShift)
}

// SetTockdiv TOCKDIV
func (r *RegisterDsiccrType) SetTockdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiccrFieldTockdivMask)|(uint32(value)<<RegisterDsiccrFieldTockdivShift))
}

// RegisterDsilvcidrType DSI Host LTDC VCID register
type RegisterDsilvcidrType uint32

func (r *RegisterDsilvcidrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsilvcidrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsilvcidrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsilvcidrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsilvcidrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsilvcidrFieldVcidShift = 0
	RegisterDsilvcidrFieldVcidMask  = 0x3
)

// GetVcid VCID
func (r *RegisterDsilvcidrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilvcidrFieldVcidMask) >> RegisterDsilvcidrFieldVcidShift)
}

// SetVcid VCID
func (r *RegisterDsilvcidrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilvcidrFieldVcidMask)|(uint32(value)<<RegisterDsilvcidrFieldVcidShift))
}

// RegisterDsilcolcrType DSI Host LTDC color coding register
type RegisterDsilcolcrType uint32

func (r *RegisterDsilcolcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsilcolcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsilcolcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsilcolcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsilcolcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsilcolcrFieldColcShift = 0
	RegisterDsilcolcrFieldColcMask  = 0xf
)

// GetColc COLC
func (r *RegisterDsilcolcrType) GetColc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilcolcrFieldColcMask) >> RegisterDsilcolcrFieldColcShift)
}

// SetColc COLC
func (r *RegisterDsilcolcrType) SetColc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilcolcrFieldColcMask)|(uint32(value)<<RegisterDsilcolcrFieldColcShift))
}

const (
	RegisterDsilcolcrFieldLpeShift = 8
	RegisterDsilcolcrFieldLpeMask  = 0x100
)

// GetLpe LPE
func (r *RegisterDsilcolcrType) GetLpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsilcolcrFieldLpeMask) != 0
}

// SetLpe LPE
func (r *RegisterDsilcolcrType) SetLpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsilcolcrFieldLpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsilcolcrFieldLpeMask)
	}
}

// RegisterDsilpcrType DSI Host LTDC polarity configuration register
type RegisterDsilpcrType uint32

func (r *RegisterDsilpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsilpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsilpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsilpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsilpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsilpcrFieldDepShift = 0
	RegisterDsilpcrFieldDepMask  = 0x1
)

// GetDep DEP
func (r *RegisterDsilpcrType) GetDep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsilpcrFieldDepMask) != 0
}

// SetDep DEP
func (r *RegisterDsilpcrType) SetDep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsilpcrFieldDepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsilpcrFieldDepMask)
	}
}

const (
	RegisterDsilpcrFieldVspShift = 1
	RegisterDsilpcrFieldVspMask  = 0x2
)

// GetVsp VSP
func (r *RegisterDsilpcrType) GetVsp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsilpcrFieldVspMask) != 0
}

// SetVsp VSP
func (r *RegisterDsilpcrType) SetVsp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsilpcrFieldVspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsilpcrFieldVspMask)
	}
}

const (
	RegisterDsilpcrFieldHspShift = 2
	RegisterDsilpcrFieldHspMask  = 0x4
)

// GetHsp HSP
func (r *RegisterDsilpcrType) GetHsp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsilpcrFieldHspMask) != 0
}

// SetHsp HSP
func (r *RegisterDsilpcrType) SetHsp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsilpcrFieldHspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsilpcrFieldHspMask)
	}
}

// RegisterDsilpmcrType DSI Host low-power mode configuration register
type RegisterDsilpmcrType uint32

func (r *RegisterDsilpmcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsilpmcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsilpmcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsilpmcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsilpmcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsilpmcrFieldVlpsizeShift = 0
	RegisterDsilpmcrFieldVlpsizeMask  = 0xff
)

// GetVlpsize VLPSIZE
func (r *RegisterDsilpmcrType) GetVlpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilpmcrFieldVlpsizeMask) >> RegisterDsilpmcrFieldVlpsizeShift)
}

// SetVlpsize VLPSIZE
func (r *RegisterDsilpmcrType) SetVlpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilpmcrFieldVlpsizeMask)|(uint32(value)<<RegisterDsilpmcrFieldVlpsizeShift))
}

const (
	RegisterDsilpmcrFieldLpsizeShift = 16
	RegisterDsilpmcrFieldLpsizeMask  = 0xff0000
)

// GetLpsize LPSIZE
func (r *RegisterDsilpmcrType) GetLpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilpmcrFieldLpsizeMask) >> RegisterDsilpmcrFieldLpsizeShift)
}

// SetLpsize LPSIZE
func (r *RegisterDsilpmcrType) SetLpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilpmcrFieldLpsizeMask)|(uint32(value)<<RegisterDsilpmcrFieldLpsizeShift))
}

// RegisterDsipcrType DSI Host protocol configuration register
type RegisterDsipcrType uint32

func (r *RegisterDsipcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsipcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsipcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsipcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsipcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsipcrFieldEttxeShift = 0
	RegisterDsipcrFieldEttxeMask  = 0x1
)

// GetEttxe ETTXE
func (r *RegisterDsipcrType) GetEttxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipcrFieldEttxeMask) != 0
}

// SetEttxe ETTXE
func (r *RegisterDsipcrType) SetEttxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipcrFieldEttxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipcrFieldEttxeMask)
	}
}

const (
	RegisterDsipcrFieldEtrxeShift = 1
	RegisterDsipcrFieldEtrxeMask  = 0x2
)

// GetEtrxe ETRXE
func (r *RegisterDsipcrType) GetEtrxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipcrFieldEtrxeMask) != 0
}

// SetEtrxe ETRXE
func (r *RegisterDsipcrType) SetEtrxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipcrFieldEtrxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipcrFieldEtrxeMask)
	}
}

const (
	RegisterDsipcrFieldBtaeShift = 2
	RegisterDsipcrFieldBtaeMask  = 0x4
)

// GetBtae BTAE
func (r *RegisterDsipcrType) GetBtae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipcrFieldBtaeMask) != 0
}

// SetBtae BTAE
func (r *RegisterDsipcrType) SetBtae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipcrFieldBtaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipcrFieldBtaeMask)
	}
}

const (
	RegisterDsipcrFieldEccrxeShift = 3
	RegisterDsipcrFieldEccrxeMask  = 0x8
)

// GetEccrxe ECCRXE
func (r *RegisterDsipcrType) GetEccrxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipcrFieldEccrxeMask) != 0
}

// SetEccrxe ECCRXE
func (r *RegisterDsipcrType) SetEccrxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipcrFieldEccrxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipcrFieldEccrxeMask)
	}
}

const (
	RegisterDsipcrFieldCrcrxeShift = 4
	RegisterDsipcrFieldCrcrxeMask  = 0x10
)

// GetCrcrxe CRCRXE
func (r *RegisterDsipcrType) GetCrcrxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipcrFieldCrcrxeMask) != 0
}

// SetCrcrxe CRCRXE
func (r *RegisterDsipcrType) SetCrcrxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipcrFieldCrcrxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipcrFieldCrcrxeMask)
	}
}

// RegisterDsigvcidrType DSI Host generic VCID register
type RegisterDsigvcidrType uint32

func (r *RegisterDsigvcidrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsigvcidrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsigvcidrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsigvcidrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsigvcidrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsigvcidrFieldVcidShift = 0
	RegisterDsigvcidrFieldVcidMask  = 0x3
)

// GetVcid VCID
func (r *RegisterDsigvcidrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsigvcidrFieldVcidMask) >> RegisterDsigvcidrFieldVcidShift)
}

// SetVcid VCID
func (r *RegisterDsigvcidrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsigvcidrFieldVcidMask)|(uint32(value)<<RegisterDsigvcidrFieldVcidShift))
}

// RegisterDsimcrType DSI Host mode configuration register
type RegisterDsimcrType uint32

func (r *RegisterDsimcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsimcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsimcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsimcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsimcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsimcrFieldCmdmShift = 0
	RegisterDsimcrFieldCmdmMask  = 0x1
)

// GetCmdm CMDM
func (r *RegisterDsimcrType) GetCmdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsimcrFieldCmdmMask) != 0
}

// SetCmdm CMDM
func (r *RegisterDsimcrType) SetCmdm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsimcrFieldCmdmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsimcrFieldCmdmMask)
	}
}

// RegisterDsivmcrType DSI Host video mode configuration register
type RegisterDsivmcrType uint32

func (r *RegisterDsivmcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivmcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivmcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivmcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivmcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivmcrFieldVmtShift = 0
	RegisterDsivmcrFieldVmtMask  = 0x3
)

// GetVmt VMT
func (r *RegisterDsivmcrType) GetVmt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldVmtMask) >> RegisterDsivmcrFieldVmtShift)
}

// SetVmt VMT
func (r *RegisterDsivmcrType) SetVmt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldVmtMask)|(uint32(value)<<RegisterDsivmcrFieldVmtShift))
}

const (
	RegisterDsivmcrFieldLpvsaeShift = 8
	RegisterDsivmcrFieldLpvsaeMask  = 0x100
)

// GetLpvsae LPVSAE
func (r *RegisterDsivmcrType) GetLpvsae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLpvsaeMask) != 0
}

// SetLpvsae LPVSAE
func (r *RegisterDsivmcrType) SetLpvsae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldLpvsaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldLpvsaeMask)
	}
}

const (
	RegisterDsivmcrFieldLpvbpeShift = 9
	RegisterDsivmcrFieldLpvbpeMask  = 0x200
)

// GetLpvbpe LPVBPE
func (r *RegisterDsivmcrType) GetLpvbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLpvbpeMask) != 0
}

// SetLpvbpe LPVBPE
func (r *RegisterDsivmcrType) SetLpvbpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldLpvbpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldLpvbpeMask)
	}
}

const (
	RegisterDsivmcrFieldLpvfpeShift = 10
	RegisterDsivmcrFieldLpvfpeMask  = 0x400
)

// GetLpvfpe LPVFPE
func (r *RegisterDsivmcrType) GetLpvfpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLpvfpeMask) != 0
}

// SetLpvfpe LPVFPE
func (r *RegisterDsivmcrType) SetLpvfpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldLpvfpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldLpvfpeMask)
	}
}

const (
	RegisterDsivmcrFieldLpvaeShift = 11
	RegisterDsivmcrFieldLpvaeMask  = 0x800
)

// GetLpvae LPVAE
func (r *RegisterDsivmcrType) GetLpvae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLpvaeMask) != 0
}

// SetLpvae LPVAE
func (r *RegisterDsivmcrType) SetLpvae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldLpvaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldLpvaeMask)
	}
}

const (
	RegisterDsivmcrFieldLphbpeShift = 12
	RegisterDsivmcrFieldLphbpeMask  = 0x1000
)

// GetLphbpe LPHBPE
func (r *RegisterDsivmcrType) GetLphbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLphbpeMask) != 0
}

// SetLphbpe LPHBPE
func (r *RegisterDsivmcrType) SetLphbpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldLphbpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldLphbpeMask)
	}
}

const (
	RegisterDsivmcrFieldLphfpeShift = 13
	RegisterDsivmcrFieldLphfpeMask  = 0x2000
)

// GetLphfpe LPHFPE
func (r *RegisterDsivmcrType) GetLphfpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLphfpeMask) != 0
}

// SetLphfpe LPHFPE
func (r *RegisterDsivmcrType) SetLphfpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldLphfpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldLphfpeMask)
	}
}

const (
	RegisterDsivmcrFieldFbtaaeShift = 14
	RegisterDsivmcrFieldFbtaaeMask  = 0x4000
)

// GetFbtaae FBTAAE
func (r *RegisterDsivmcrType) GetFbtaae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldFbtaaeMask) != 0
}

// SetFbtaae FBTAAE
func (r *RegisterDsivmcrType) SetFbtaae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldFbtaaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldFbtaaeMask)
	}
}

const (
	RegisterDsivmcrFieldLpceShift = 15
	RegisterDsivmcrFieldLpceMask  = 0x8000
)

// GetLpce LPCE
func (r *RegisterDsivmcrType) GetLpce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLpceMask) != 0
}

// SetLpce LPCE
func (r *RegisterDsivmcrType) SetLpce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldLpceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldLpceMask)
	}
}

const (
	RegisterDsivmcrFieldPgeShift = 16
	RegisterDsivmcrFieldPgeMask  = 0x10000
)

// GetPge PGE
func (r *RegisterDsivmcrType) GetPge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldPgeMask) != 0
}

// SetPge PGE
func (r *RegisterDsivmcrType) SetPge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldPgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldPgeMask)
	}
}

const (
	RegisterDsivmcrFieldPgmShift = 20
	RegisterDsivmcrFieldPgmMask  = 0x100000
)

// GetPgm PGM
func (r *RegisterDsivmcrType) GetPgm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldPgmMask) != 0
}

// SetPgm PGM
func (r *RegisterDsivmcrType) SetPgm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldPgmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldPgmMask)
	}
}

const (
	RegisterDsivmcrFieldPgoShift = 24
	RegisterDsivmcrFieldPgoMask  = 0x1000000
)

// GetPgo PGO
func (r *RegisterDsivmcrType) GetPgo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldPgoMask) != 0
}

// SetPgo PGO
func (r *RegisterDsivmcrType) SetPgo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldPgoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldPgoMask)
	}
}

// RegisterDsivpcrType DSI Host video packet configuration register
type RegisterDsivpcrType uint32

func (r *RegisterDsivpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivpcrFieldVpsizeShift = 0
	RegisterDsivpcrFieldVpsizeMask  = 0x3fff
)

// GetVpsize VPSIZE
func (r *RegisterDsivpcrType) GetVpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivpcrFieldVpsizeMask) >> RegisterDsivpcrFieldVpsizeShift)
}

// SetVpsize VPSIZE
func (r *RegisterDsivpcrType) SetVpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivpcrFieldVpsizeMask)|(uint32(value)<<RegisterDsivpcrFieldVpsizeShift))
}

// RegisterDsivccrType DSI Host video chunks configuration register
type RegisterDsivccrType uint32

func (r *RegisterDsivccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivccrFieldNumcShift = 0
	RegisterDsivccrFieldNumcMask  = 0x1fff
)

// GetNumc NUMC
func (r *RegisterDsivccrType) GetNumc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivccrFieldNumcMask) >> RegisterDsivccrFieldNumcShift)
}

// SetNumc NUMC
func (r *RegisterDsivccrType) SetNumc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivccrFieldNumcMask)|(uint32(value)<<RegisterDsivccrFieldNumcShift))
}

// RegisterDsivnpcrType DSI Host video null packet configuration register
type RegisterDsivnpcrType uint32

func (r *RegisterDsivnpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivnpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivnpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivnpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivnpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivnpcrFieldNpsizeShift = 0
	RegisterDsivnpcrFieldNpsizeMask  = 0x1fff
)

// GetNpsize NPSIZE
func (r *RegisterDsivnpcrType) GetNpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivnpcrFieldNpsizeMask) >> RegisterDsivnpcrFieldNpsizeShift)
}

// SetNpsize NPSIZE
func (r *RegisterDsivnpcrType) SetNpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivnpcrFieldNpsizeMask)|(uint32(value)<<RegisterDsivnpcrFieldNpsizeShift))
}

// RegisterDsivhsacrType DSI Host video HSA configuration register
type RegisterDsivhsacrType uint32

func (r *RegisterDsivhsacrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivhsacrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivhsacrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivhsacrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivhsacrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivhsacrFieldHsaShift = 0
	RegisterDsivhsacrFieldHsaMask  = 0xfff
)

// GetHsa HSA
func (r *RegisterDsivhsacrType) GetHsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivhsacrFieldHsaMask) >> RegisterDsivhsacrFieldHsaShift)
}

// SetHsa HSA
func (r *RegisterDsivhsacrType) SetHsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivhsacrFieldHsaMask)|(uint32(value)<<RegisterDsivhsacrFieldHsaShift))
}

// RegisterDsivhbpcrType DSI Host video HBP configuration register
type RegisterDsivhbpcrType uint32

func (r *RegisterDsivhbpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivhbpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivhbpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivhbpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivhbpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivhbpcrFieldHbpShift = 0
	RegisterDsivhbpcrFieldHbpMask  = 0xfff
)

// GetHbp HBP
func (r *RegisterDsivhbpcrType) GetHbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivhbpcrFieldHbpMask) >> RegisterDsivhbpcrFieldHbpShift)
}

// SetHbp HBP
func (r *RegisterDsivhbpcrType) SetHbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivhbpcrFieldHbpMask)|(uint32(value)<<RegisterDsivhbpcrFieldHbpShift))
}

// RegisterDsivlcrType DSI Host video line configuration register
type RegisterDsivlcrType uint32

func (r *RegisterDsivlcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivlcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivlcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivlcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivlcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivlcrFieldHlineShift = 0
	RegisterDsivlcrFieldHlineMask  = 0x7fff
)

// GetHline HLINE
func (r *RegisterDsivlcrType) GetHline() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivlcrFieldHlineMask) >> RegisterDsivlcrFieldHlineShift)
}

// SetHline HLINE
func (r *RegisterDsivlcrType) SetHline(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivlcrFieldHlineMask)|(uint32(value)<<RegisterDsivlcrFieldHlineShift))
}

// RegisterDsivvsacrType DSI Host video VSA configuration register
type RegisterDsivvsacrType uint32

func (r *RegisterDsivvsacrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivvsacrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivvsacrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivvsacrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivvsacrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivvsacrFieldVsaShift = 0
	RegisterDsivvsacrFieldVsaMask  = 0x3ff
)

// GetVsa VSA
func (r *RegisterDsivvsacrType) GetVsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvsacrFieldVsaMask) >> RegisterDsivvsacrFieldVsaShift)
}

// SetVsa VSA
func (r *RegisterDsivvsacrType) SetVsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvsacrFieldVsaMask)|(uint32(value)<<RegisterDsivvsacrFieldVsaShift))
}

// RegisterDsivvbpcrType DSI Host video VBP configuration register
type RegisterDsivvbpcrType uint32

func (r *RegisterDsivvbpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivvbpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivvbpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivvbpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivvbpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivvbpcrFieldVbpShift = 0
	RegisterDsivvbpcrFieldVbpMask  = 0x3ff
)

// GetVbp VBP
func (r *RegisterDsivvbpcrType) GetVbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvbpcrFieldVbpMask) >> RegisterDsivvbpcrFieldVbpShift)
}

// SetVbp VBP
func (r *RegisterDsivvbpcrType) SetVbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvbpcrFieldVbpMask)|(uint32(value)<<RegisterDsivvbpcrFieldVbpShift))
}

// RegisterDsivvfpcrType DSI Host video VFP configuration register
type RegisterDsivvfpcrType uint32

func (r *RegisterDsivvfpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivvfpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivvfpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivvfpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivvfpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivvfpcrFieldVfpShift = 0
	RegisterDsivvfpcrFieldVfpMask  = 0x3ff
)

// GetVfp VFP
func (r *RegisterDsivvfpcrType) GetVfp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvfpcrFieldVfpMask) >> RegisterDsivvfpcrFieldVfpShift)
}

// SetVfp VFP
func (r *RegisterDsivvfpcrType) SetVfp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvfpcrFieldVfpMask)|(uint32(value)<<RegisterDsivvfpcrFieldVfpShift))
}

// RegisterDsivvacrType DSI Host video VA configuration register
type RegisterDsivvacrType uint32

func (r *RegisterDsivvacrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivvacrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivvacrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivvacrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivvacrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivvacrFieldVaShift = 0
	RegisterDsivvacrFieldVaMask  = 0x3fff
)

// GetVa VA
func (r *RegisterDsivvacrType) GetVa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvacrFieldVaMask) >> RegisterDsivvacrFieldVaShift)
}

// SetVa VA
func (r *RegisterDsivvacrType) SetVa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvacrFieldVaMask)|(uint32(value)<<RegisterDsivvacrFieldVaShift))
}

// RegisterDsilccrType DSI Host LTDC command configuration register
type RegisterDsilccrType uint32

func (r *RegisterDsilccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsilccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsilccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsilccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsilccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsilccrFieldCmdsizeShift = 0
	RegisterDsilccrFieldCmdsizeMask  = 0xffff
)

// GetCmdsize CMDSIZE
func (r *RegisterDsilccrType) GetCmdsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsilccrFieldCmdsizeMask) >> RegisterDsilccrFieldCmdsizeShift)
}

// SetCmdsize CMDSIZE
func (r *RegisterDsilccrType) SetCmdsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilccrFieldCmdsizeMask)|(uint32(value)<<RegisterDsilccrFieldCmdsizeShift))
}

// RegisterDsicmcrType DSI Host command mode configuration register
type RegisterDsicmcrType uint32

func (r *RegisterDsicmcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsicmcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsicmcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsicmcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsicmcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsicmcrFieldTeareShift = 0
	RegisterDsicmcrFieldTeareMask  = 0x1
)

// GetTeare TEARE
func (r *RegisterDsicmcrType) GetTeare() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldTeareMask) != 0
}

// SetTeare TEARE
func (r *RegisterDsicmcrType) SetTeare(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldTeareMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldTeareMask)
	}
}

const (
	RegisterDsicmcrFieldAreShift = 1
	RegisterDsicmcrFieldAreMask  = 0x2
)

// GetAre ARE
func (r *RegisterDsicmcrType) GetAre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldAreMask) != 0
}

// SetAre ARE
func (r *RegisterDsicmcrType) SetAre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldAreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldAreMask)
	}
}

const (
	RegisterDsicmcrFieldGsw0txShift = 8
	RegisterDsicmcrFieldGsw0txMask  = 0x100
)

// GetGsw0tx GSW0TX
func (r *RegisterDsicmcrType) GetGsw0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsw0txMask) != 0
}

// SetGsw0tx GSW0TX
func (r *RegisterDsicmcrType) SetGsw0tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldGsw0txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldGsw0txMask)
	}
}

const (
	RegisterDsicmcrFieldGsw1txShift = 9
	RegisterDsicmcrFieldGsw1txMask  = 0x200
)

// GetGsw1tx GSW1TX
func (r *RegisterDsicmcrType) GetGsw1tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsw1txMask) != 0
}

// SetGsw1tx GSW1TX
func (r *RegisterDsicmcrType) SetGsw1tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldGsw1txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldGsw1txMask)
	}
}

const (
	RegisterDsicmcrFieldGsw2txShift = 10
	RegisterDsicmcrFieldGsw2txMask  = 0x400
)

// GetGsw2tx GSW2TX
func (r *RegisterDsicmcrType) GetGsw2tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsw2txMask) != 0
}

// SetGsw2tx GSW2TX
func (r *RegisterDsicmcrType) SetGsw2tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldGsw2txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldGsw2txMask)
	}
}

const (
	RegisterDsicmcrFieldGsr0txShift = 11
	RegisterDsicmcrFieldGsr0txMask  = 0x800
)

// GetGsr0tx GSR0TX
func (r *RegisterDsicmcrType) GetGsr0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsr0txMask) != 0
}

// SetGsr0tx GSR0TX
func (r *RegisterDsicmcrType) SetGsr0tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldGsr0txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldGsr0txMask)
	}
}

const (
	RegisterDsicmcrFieldGsr1txShift = 12
	RegisterDsicmcrFieldGsr1txMask  = 0x1000
)

// GetGsr1tx GSR1TX
func (r *RegisterDsicmcrType) GetGsr1tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsr1txMask) != 0
}

// SetGsr1tx GSR1TX
func (r *RegisterDsicmcrType) SetGsr1tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldGsr1txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldGsr1txMask)
	}
}

const (
	RegisterDsicmcrFieldGsr2txShift = 13
	RegisterDsicmcrFieldGsr2txMask  = 0x2000
)

// GetGsr2tx GSR2TX
func (r *RegisterDsicmcrType) GetGsr2tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsr2txMask) != 0
}

// SetGsr2tx GSR2TX
func (r *RegisterDsicmcrType) SetGsr2tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldGsr2txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldGsr2txMask)
	}
}

const (
	RegisterDsicmcrFieldGlwtxShift = 14
	RegisterDsicmcrFieldGlwtxMask  = 0x4000
)

// GetGlwtx GLWTX
func (r *RegisterDsicmcrType) GetGlwtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGlwtxMask) != 0
}

// SetGlwtx GLWTX
func (r *RegisterDsicmcrType) SetGlwtx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldGlwtxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldGlwtxMask)
	}
}

const (
	RegisterDsicmcrFieldDsw0txShift = 16
	RegisterDsicmcrFieldDsw0txMask  = 0x10000
)

// GetDsw0tx DSW0TX
func (r *RegisterDsicmcrType) GetDsw0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldDsw0txMask) != 0
}

// SetDsw0tx DSW0TX
func (r *RegisterDsicmcrType) SetDsw0tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldDsw0txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldDsw0txMask)
	}
}

const (
	RegisterDsicmcrFieldDsw1txShift = 17
	RegisterDsicmcrFieldDsw1txMask  = 0x20000
)

// GetDsw1tx DSW1TX
func (r *RegisterDsicmcrType) GetDsw1tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldDsw1txMask) != 0
}

// SetDsw1tx DSW1TX
func (r *RegisterDsicmcrType) SetDsw1tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldDsw1txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldDsw1txMask)
	}
}

const (
	RegisterDsicmcrFieldDsr0txShift = 18
	RegisterDsicmcrFieldDsr0txMask  = 0x40000
)

// GetDsr0tx DSR0TX
func (r *RegisterDsicmcrType) GetDsr0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldDsr0txMask) != 0
}

// SetDsr0tx DSR0TX
func (r *RegisterDsicmcrType) SetDsr0tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldDsr0txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldDsr0txMask)
	}
}

const (
	RegisterDsicmcrFieldDlwtxShift = 19
	RegisterDsicmcrFieldDlwtxMask  = 0x80000
)

// GetDlwtx DLWTX
func (r *RegisterDsicmcrType) GetDlwtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldDlwtxMask) != 0
}

// SetDlwtx DLWTX
func (r *RegisterDsicmcrType) SetDlwtx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldDlwtxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldDlwtxMask)
	}
}

const (
	RegisterDsicmcrFieldMrdpsShift = 24
	RegisterDsicmcrFieldMrdpsMask  = 0x1000000
)

// GetMrdps MRDPS
func (r *RegisterDsicmcrType) GetMrdps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldMrdpsMask) != 0
}

// SetMrdps MRDPS
func (r *RegisterDsicmcrType) SetMrdps(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldMrdpsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldMrdpsMask)
	}
}

// RegisterDsighcrType DSI Host generic header configuration register
type RegisterDsighcrType uint32

func (r *RegisterDsighcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsighcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsighcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsighcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsighcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsighcrFieldDtShift = 0
	RegisterDsighcrFieldDtMask  = 0x3f
)

// GetDt DT
func (r *RegisterDsighcrType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsighcrFieldDtMask) >> RegisterDsighcrFieldDtShift)
}

// SetDt DT
func (r *RegisterDsighcrType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsighcrFieldDtMask)|(uint32(value)<<RegisterDsighcrFieldDtShift))
}

const (
	RegisterDsighcrFieldVcidShift = 6
	RegisterDsighcrFieldVcidMask  = 0xc0
)

// GetVcid VCID
func (r *RegisterDsighcrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsighcrFieldVcidMask) >> RegisterDsighcrFieldVcidShift)
}

// SetVcid VCID
func (r *RegisterDsighcrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsighcrFieldVcidMask)|(uint32(value)<<RegisterDsighcrFieldVcidShift))
}

const (
	RegisterDsighcrFieldWclsbShift = 8
	RegisterDsighcrFieldWclsbMask  = 0xff00
)

// GetWclsb WCLSB
func (r *RegisterDsighcrType) GetWclsb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsighcrFieldWclsbMask) >> RegisterDsighcrFieldWclsbShift)
}

// SetWclsb WCLSB
func (r *RegisterDsighcrType) SetWclsb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsighcrFieldWclsbMask)|(uint32(value)<<RegisterDsighcrFieldWclsbShift))
}

const (
	RegisterDsighcrFieldWcmsbShift = 16
	RegisterDsighcrFieldWcmsbMask  = 0xff0000
)

// GetWcmsb WCMSB
func (r *RegisterDsighcrType) GetWcmsb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsighcrFieldWcmsbMask) >> RegisterDsighcrFieldWcmsbShift)
}

// SetWcmsb WCMSB
func (r *RegisterDsighcrType) SetWcmsb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsighcrFieldWcmsbMask)|(uint32(value)<<RegisterDsighcrFieldWcmsbShift))
}

// RegisterDsigpdrType DSI Host generic payload data register
type RegisterDsigpdrType uint32

func (r *RegisterDsigpdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsigpdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsigpdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsigpdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsigpdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsigpdrFieldData1Shift = 0
	RegisterDsigpdrFieldData1Mask  = 0xff
)

// GetData1 DATA1
func (r *RegisterDsigpdrType) GetData1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsigpdrFieldData1Mask) >> RegisterDsigpdrFieldData1Shift)
}

// SetData1 DATA1
func (r *RegisterDsigpdrType) SetData1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsigpdrFieldData1Mask)|(uint32(value)<<RegisterDsigpdrFieldData1Shift))
}

const (
	RegisterDsigpdrFieldData2Shift = 8
	RegisterDsigpdrFieldData2Mask  = 0xff00
)

// GetData2 DATA2
func (r *RegisterDsigpdrType) GetData2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsigpdrFieldData2Mask) >> RegisterDsigpdrFieldData2Shift)
}

// SetData2 DATA2
func (r *RegisterDsigpdrType) SetData2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsigpdrFieldData2Mask)|(uint32(value)<<RegisterDsigpdrFieldData2Shift))
}

const (
	RegisterDsigpdrFieldData3Shift = 16
	RegisterDsigpdrFieldData3Mask  = 0xff0000
)

// GetData3 DATA3
func (r *RegisterDsigpdrType) GetData3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsigpdrFieldData3Mask) >> RegisterDsigpdrFieldData3Shift)
}

// SetData3 DATA3
func (r *RegisterDsigpdrType) SetData3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsigpdrFieldData3Mask)|(uint32(value)<<RegisterDsigpdrFieldData3Shift))
}

const (
	RegisterDsigpdrFieldData4Shift = 24
	RegisterDsigpdrFieldData4Mask  = 0xff000000
)

// GetData4 DATA4
func (r *RegisterDsigpdrType) GetData4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsigpdrFieldData4Mask) >> RegisterDsigpdrFieldData4Shift)
}

// SetData4 DATA4
func (r *RegisterDsigpdrType) SetData4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsigpdrFieldData4Mask)|(uint32(value)<<RegisterDsigpdrFieldData4Shift))
}

// RegisterDsigpsrType DSI Host generic packet status register
type RegisterDsigpsrType uint32

func (r *RegisterDsigpsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsigpsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsigpsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsigpsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsigpsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsigpsrFieldCmdfeShift = 0
	RegisterDsigpsrFieldCmdfeMask  = 0x1
)

// GetCmdfe CMDFE
func (r *RegisterDsigpsrType) GetCmdfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldCmdfeMask) != 0
}

// SetCmdfe CMDFE
func (r *RegisterDsigpsrType) SetCmdfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsigpsrFieldCmdfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsigpsrFieldCmdfeMask)
	}
}

const (
	RegisterDsigpsrFieldCmdffShift = 1
	RegisterDsigpsrFieldCmdffMask  = 0x2
)

// GetCmdff CMDFF
func (r *RegisterDsigpsrType) GetCmdff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldCmdffMask) != 0
}

// SetCmdff CMDFF
func (r *RegisterDsigpsrType) SetCmdff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsigpsrFieldCmdffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsigpsrFieldCmdffMask)
	}
}

const (
	RegisterDsigpsrFieldPwrfeShift = 2
	RegisterDsigpsrFieldPwrfeMask  = 0x4
)

// GetPwrfe PWRFE
func (r *RegisterDsigpsrType) GetPwrfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldPwrfeMask) != 0
}

// SetPwrfe PWRFE
func (r *RegisterDsigpsrType) SetPwrfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsigpsrFieldPwrfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsigpsrFieldPwrfeMask)
	}
}

const (
	RegisterDsigpsrFieldPwrffShift = 3
	RegisterDsigpsrFieldPwrffMask  = 0x8
)

// GetPwrff PWRFF
func (r *RegisterDsigpsrType) GetPwrff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldPwrffMask) != 0
}

// SetPwrff PWRFF
func (r *RegisterDsigpsrType) SetPwrff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsigpsrFieldPwrffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsigpsrFieldPwrffMask)
	}
}

const (
	RegisterDsigpsrFieldPrdfeShift = 4
	RegisterDsigpsrFieldPrdfeMask  = 0x10
)

// GetPrdfe PRDFE
func (r *RegisterDsigpsrType) GetPrdfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldPrdfeMask) != 0
}

// SetPrdfe PRDFE
func (r *RegisterDsigpsrType) SetPrdfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsigpsrFieldPrdfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsigpsrFieldPrdfeMask)
	}
}

const (
	RegisterDsigpsrFieldPrdffShift = 5
	RegisterDsigpsrFieldPrdffMask  = 0x20
)

// GetPrdff PRDFF
func (r *RegisterDsigpsrType) GetPrdff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldPrdffMask) != 0
}

// SetPrdff PRDFF
func (r *RegisterDsigpsrType) SetPrdff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsigpsrFieldPrdffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsigpsrFieldPrdffMask)
	}
}

const (
	RegisterDsigpsrFieldRcbShift = 6
	RegisterDsigpsrFieldRcbMask  = 0x40
)

// GetRcb RCB
func (r *RegisterDsigpsrType) GetRcb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldRcbMask) != 0
}

// SetRcb RCB
func (r *RegisterDsigpsrType) SetRcb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsigpsrFieldRcbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsigpsrFieldRcbMask)
	}
}

// RegisterDsitccr0Type DSI Host timeout counter configuration register 0
type RegisterDsitccr0Type uint32

func (r *RegisterDsitccr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsitccr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsitccr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsitccr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsitccr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsitccr0FieldLprxtocntShift = 0
	RegisterDsitccr0FieldLprxtocntMask  = 0xffff
)

// GetLprxtocnt LPRX_TOCNT
func (r *RegisterDsitccr0Type) GetLprxtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr0FieldLprxtocntMask) >> RegisterDsitccr0FieldLprxtocntShift)
}

// SetLprxtocnt LPRX_TOCNT
func (r *RegisterDsitccr0Type) SetLprxtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr0FieldLprxtocntMask)|(uint32(value)<<RegisterDsitccr0FieldLprxtocntShift))
}

const (
	RegisterDsitccr0FieldHstxtocntShift = 16
	RegisterDsitccr0FieldHstxtocntMask  = 0xffff0000
)

// GetHstxtocnt HSTX_TOCNT
func (r *RegisterDsitccr0Type) GetHstxtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr0FieldHstxtocntMask) >> RegisterDsitccr0FieldHstxtocntShift)
}

// SetHstxtocnt HSTX_TOCNT
func (r *RegisterDsitccr0Type) SetHstxtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr0FieldHstxtocntMask)|(uint32(value)<<RegisterDsitccr0FieldHstxtocntShift))
}

// RegisterDsitccr1Type DSI Host timeout counter configuration register 1
type RegisterDsitccr1Type uint32

func (r *RegisterDsitccr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsitccr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsitccr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsitccr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsitccr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsitccr1FieldHsrdtocntShift = 0
	RegisterDsitccr1FieldHsrdtocntMask  = 0xffff
)

// GetHsrdtocnt HSRD_TOCNT
func (r *RegisterDsitccr1Type) GetHsrdtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr1FieldHsrdtocntMask) >> RegisterDsitccr1FieldHsrdtocntShift)
}

// SetHsrdtocnt HSRD_TOCNT
func (r *RegisterDsitccr1Type) SetHsrdtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr1FieldHsrdtocntMask)|(uint32(value)<<RegisterDsitccr1FieldHsrdtocntShift))
}

// RegisterDsitccr2Type DSI Host timeout counter configuration register 2
type RegisterDsitccr2Type uint32

func (r *RegisterDsitccr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsitccr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsitccr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsitccr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsitccr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsitccr2FieldLprdtocntShift = 0
	RegisterDsitccr2FieldLprdtocntMask  = 0xffff
)

// GetLprdtocnt LPRD_TOCNT
func (r *RegisterDsitccr2Type) GetLprdtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr2FieldLprdtocntMask) >> RegisterDsitccr2FieldLprdtocntShift)
}

// SetLprdtocnt LPRD_TOCNT
func (r *RegisterDsitccr2Type) SetLprdtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr2FieldLprdtocntMask)|(uint32(value)<<RegisterDsitccr2FieldLprdtocntShift))
}

// RegisterDsitccr3Type DSI Host timeout counter configuration register 3
type RegisterDsitccr3Type uint32

func (r *RegisterDsitccr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsitccr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsitccr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsitccr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsitccr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsitccr3FieldHswrtocntShift = 0
	RegisterDsitccr3FieldHswrtocntMask  = 0xffff
)

// GetHswrtocnt HSWR_TOCNT
func (r *RegisterDsitccr3Type) GetHswrtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr3FieldHswrtocntMask) >> RegisterDsitccr3FieldHswrtocntShift)
}

// SetHswrtocnt HSWR_TOCNT
func (r *RegisterDsitccr3Type) SetHswrtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr3FieldHswrtocntMask)|(uint32(value)<<RegisterDsitccr3FieldHswrtocntShift))
}

const (
	RegisterDsitccr3FieldPmShift = 24
	RegisterDsitccr3FieldPmMask  = 0x1000000
)

// GetPm PM
func (r *RegisterDsitccr3Type) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr3FieldPmMask) != 0
}

// SetPm PM
func (r *RegisterDsitccr3Type) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsitccr3FieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr3FieldPmMask)
	}
}

// RegisterDsitccr4Type DSI Host timeout counter configuration register 4
type RegisterDsitccr4Type uint32

func (r *RegisterDsitccr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsitccr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsitccr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsitccr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsitccr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsitccr4FieldLpwrtocntShift = 0
	RegisterDsitccr4FieldLpwrtocntMask  = 0xffff
)

// GetLpwrtocnt LPWR_TOCNT
func (r *RegisterDsitccr4Type) GetLpwrtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr4FieldLpwrtocntMask) >> RegisterDsitccr4FieldLpwrtocntShift)
}

// SetLpwrtocnt LPWR_TOCNT
func (r *RegisterDsitccr4Type) SetLpwrtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr4FieldLpwrtocntMask)|(uint32(value)<<RegisterDsitccr4FieldLpwrtocntShift))
}

// RegisterDsitccr5Type DSI Host timeout counter configuration register 5
type RegisterDsitccr5Type uint32

func (r *RegisterDsitccr5Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsitccr5Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsitccr5Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsitccr5Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsitccr5Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsitccr5FieldBtatocntShift = 0
	RegisterDsitccr5FieldBtatocntMask  = 0xffff
)

// GetBtatocnt BTA_TOCNT
func (r *RegisterDsitccr5Type) GetBtatocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr5FieldBtatocntMask) >> RegisterDsitccr5FieldBtatocntShift)
}

// SetBtatocnt BTA_TOCNT
func (r *RegisterDsitccr5Type) SetBtatocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr5FieldBtatocntMask)|(uint32(value)<<RegisterDsitccr5FieldBtatocntShift))
}

// RegisterDsiclcrType DSI Host clock lane configuration register
type RegisterDsiclcrType uint32

func (r *RegisterDsiclcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiclcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiclcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiclcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiclcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiclcrFieldDpccShift = 0
	RegisterDsiclcrFieldDpccMask  = 0x1
)

// GetDpcc DPCC
func (r *RegisterDsiclcrType) GetDpcc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiclcrFieldDpccMask) != 0
}

// SetDpcc DPCC
func (r *RegisterDsiclcrType) SetDpcc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiclcrFieldDpccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiclcrFieldDpccMask)
	}
}

const (
	RegisterDsiclcrFieldAcrShift = 1
	RegisterDsiclcrFieldAcrMask  = 0x2
)

// GetAcr ACR
func (r *RegisterDsiclcrType) GetAcr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiclcrFieldAcrMask) != 0
}

// SetAcr ACR
func (r *RegisterDsiclcrType) SetAcr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiclcrFieldAcrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiclcrFieldAcrMask)
	}
}

// RegisterDsicltcrType DSI Host clock lane timer configuration register
type RegisterDsicltcrType uint32

func (r *RegisterDsicltcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsicltcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsicltcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsicltcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsicltcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsicltcrFieldLp2hstimeShift = 0
	RegisterDsicltcrFieldLp2hstimeMask  = 0x3ff
)

// GetLp2hstime LP2HS_TIME
func (r *RegisterDsicltcrType) GetLp2hstime() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsicltcrFieldLp2hstimeMask) >> RegisterDsicltcrFieldLp2hstimeShift)
}

// SetLp2hstime LP2HS_TIME
func (r *RegisterDsicltcrType) SetLp2hstime(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsicltcrFieldLp2hstimeMask)|(uint32(value)<<RegisterDsicltcrFieldLp2hstimeShift))
}

const (
	RegisterDsicltcrFieldHs2lptimeShift = 16
	RegisterDsicltcrFieldHs2lptimeMask  = 0x3ff0000
)

// GetHs2lptime HS2LP_TIME
func (r *RegisterDsicltcrType) GetHs2lptime() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsicltcrFieldHs2lptimeMask) >> RegisterDsicltcrFieldHs2lptimeShift)
}

// SetHs2lptime HS2LP_TIME
func (r *RegisterDsicltcrType) SetHs2lptime(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsicltcrFieldHs2lptimeMask)|(uint32(value)<<RegisterDsicltcrFieldHs2lptimeShift))
}

// RegisterDsidltcrType DSI Host data lane timer configuration register
type RegisterDsidltcrType uint32

func (r *RegisterDsidltcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsidltcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsidltcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsidltcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsidltcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsidltcrFieldMrdtimeShift = 0
	RegisterDsidltcrFieldMrdtimeMask  = 0x7fff
)

// GetMrdtime Maximum read time
func (r *RegisterDsidltcrType) GetMrdtime() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsidltcrFieldMrdtimeMask) >> RegisterDsidltcrFieldMrdtimeShift)
}

// SetMrdtime Maximum read time
func (r *RegisterDsidltcrType) SetMrdtime(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsidltcrFieldMrdtimeMask)|(uint32(value)<<RegisterDsidltcrFieldMrdtimeShift))
}

const (
	RegisterDsidltcrFieldLp2hstimeShift = 16
	RegisterDsidltcrFieldLp2hstimeMask  = 0xff0000
)

// GetLp2hstime LP2HS_TIME
func (r *RegisterDsidltcrType) GetLp2hstime() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsidltcrFieldLp2hstimeMask) >> RegisterDsidltcrFieldLp2hstimeShift)
}

// SetLp2hstime LP2HS_TIME
func (r *RegisterDsidltcrType) SetLp2hstime(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsidltcrFieldLp2hstimeMask)|(uint32(value)<<RegisterDsidltcrFieldLp2hstimeShift))
}

const (
	RegisterDsidltcrFieldHs2lptimeShift = 24
	RegisterDsidltcrFieldHs2lptimeMask  = 0xff000000
)

// GetHs2lptime HS2LP_TIME
func (r *RegisterDsidltcrType) GetHs2lptime() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsidltcrFieldHs2lptimeMask) >> RegisterDsidltcrFieldHs2lptimeShift)
}

// SetHs2lptime HS2LP_TIME
func (r *RegisterDsidltcrType) SetHs2lptime(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsidltcrFieldHs2lptimeMask)|(uint32(value)<<RegisterDsidltcrFieldHs2lptimeShift))
}

// RegisterDsipctlrType DSI Host PHY control register
type RegisterDsipctlrType uint32

func (r *RegisterDsipctlrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsipctlrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsipctlrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsipctlrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsipctlrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsipctlrFieldDenShift = 1
	RegisterDsipctlrFieldDenMask  = 0x2
)

// GetDen DEN
func (r *RegisterDsipctlrType) GetDen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipctlrFieldDenMask) != 0
}

// SetDen DEN
func (r *RegisterDsipctlrType) SetDen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipctlrFieldDenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipctlrFieldDenMask)
	}
}

const (
	RegisterDsipctlrFieldCkeShift = 2
	RegisterDsipctlrFieldCkeMask  = 0x4
)

// GetCke CKE
func (r *RegisterDsipctlrType) GetCke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipctlrFieldCkeMask) != 0
}

// SetCke CKE
func (r *RegisterDsipctlrType) SetCke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipctlrFieldCkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipctlrFieldCkeMask)
	}
}

// RegisterDsipconfrType DSI Host PHY configuration register
type RegisterDsipconfrType uint32

func (r *RegisterDsipconfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsipconfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsipconfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsipconfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsipconfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsipconfrFieldNlShift = 0
	RegisterDsipconfrFieldNlMask  = 0x3
)

// GetNl NL
func (r *RegisterDsipconfrType) GetNl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsipconfrFieldNlMask) >> RegisterDsipconfrFieldNlShift)
}

// SetNl NL
func (r *RegisterDsipconfrType) SetNl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsipconfrFieldNlMask)|(uint32(value)<<RegisterDsipconfrFieldNlShift))
}

const (
	RegisterDsipconfrFieldSwtimeShift = 8
	RegisterDsipconfrFieldSwtimeMask  = 0xff00
)

// GetSwtime SW_TIME
func (r *RegisterDsipconfrType) GetSwtime() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsipconfrFieldSwtimeMask) >> RegisterDsipconfrFieldSwtimeShift)
}

// SetSwtime SW_TIME
func (r *RegisterDsipconfrType) SetSwtime(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsipconfrFieldSwtimeMask)|(uint32(value)<<RegisterDsipconfrFieldSwtimeShift))
}

// RegisterDsipucrType DSI Host PHY ULPS control register
type RegisterDsipucrType uint32

func (r *RegisterDsipucrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsipucrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsipucrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsipucrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsipucrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsipucrFieldUrclShift = 0
	RegisterDsipucrFieldUrclMask  = 0x1
)

// GetUrcl URCL
func (r *RegisterDsipucrType) GetUrcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipucrFieldUrclMask) != 0
}

// SetUrcl URCL
func (r *RegisterDsipucrType) SetUrcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipucrFieldUrclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipucrFieldUrclMask)
	}
}

const (
	RegisterDsipucrFieldUeclShift = 1
	RegisterDsipucrFieldUeclMask  = 0x2
)

// GetUecl UECL
func (r *RegisterDsipucrType) GetUecl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipucrFieldUeclMask) != 0
}

// SetUecl UECL
func (r *RegisterDsipucrType) SetUecl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipucrFieldUeclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipucrFieldUeclMask)
	}
}

const (
	RegisterDsipucrFieldUrdlShift = 2
	RegisterDsipucrFieldUrdlMask  = 0x4
)

// GetUrdl URDL
func (r *RegisterDsipucrType) GetUrdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipucrFieldUrdlMask) != 0
}

// SetUrdl URDL
func (r *RegisterDsipucrType) SetUrdl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipucrFieldUrdlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipucrFieldUrdlMask)
	}
}

const (
	RegisterDsipucrFieldUedlShift = 3
	RegisterDsipucrFieldUedlMask  = 0x8
)

// GetUedl UEDL
func (r *RegisterDsipucrType) GetUedl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipucrFieldUedlMask) != 0
}

// SetUedl UEDL
func (r *RegisterDsipucrType) SetUedl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipucrFieldUedlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipucrFieldUedlMask)
	}
}

// RegisterDsipttcrType DSI Host PHY TX triggers configuration register
type RegisterDsipttcrType uint32

func (r *RegisterDsipttcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsipttcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsipttcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsipttcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsipttcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsipttcrFieldTxtrigShift = 0
	RegisterDsipttcrFieldTxtrigMask  = 0xf
)

// GetTxtrig TX_TRIG
func (r *RegisterDsipttcrType) GetTxtrig() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsipttcrFieldTxtrigMask) >> RegisterDsipttcrFieldTxtrigShift)
}

// SetTxtrig TX_TRIG
func (r *RegisterDsipttcrType) SetTxtrig(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsipttcrFieldTxtrigMask)|(uint32(value)<<RegisterDsipttcrFieldTxtrigShift))
}

// RegisterDsipsrType DSI Host PHY status register
type RegisterDsipsrType uint32

func (r *RegisterDsipsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsipsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsipsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsipsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsipsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsipsrFieldPdShift = 1
	RegisterDsipsrFieldPdMask  = 0x2
)

// GetPd PD
func (r *RegisterDsipsrType) GetPd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldPdMask) != 0
}

// SetPd PD
func (r *RegisterDsipsrType) SetPd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipsrFieldPdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipsrFieldPdMask)
	}
}

const (
	RegisterDsipsrFieldPsscShift = 2
	RegisterDsipsrFieldPsscMask  = 0x4
)

// GetPssc PSSC
func (r *RegisterDsipsrType) GetPssc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldPsscMask) != 0
}

// SetPssc PSSC
func (r *RegisterDsipsrType) SetPssc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipsrFieldPsscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipsrFieldPsscMask)
	}
}

const (
	RegisterDsipsrFieldUancShift = 3
	RegisterDsipsrFieldUancMask  = 0x8
)

// GetUanc UANC
func (r *RegisterDsipsrType) GetUanc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldUancMask) != 0
}

// SetUanc UANC
func (r *RegisterDsipsrType) SetUanc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipsrFieldUancMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipsrFieldUancMask)
	}
}

const (
	RegisterDsipsrFieldPss0Shift = 4
	RegisterDsipsrFieldPss0Mask  = 0x10
)

// GetPss0 PSS0
func (r *RegisterDsipsrType) GetPss0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldPss0Mask) != 0
}

// SetPss0 PSS0
func (r *RegisterDsipsrType) SetPss0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipsrFieldPss0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipsrFieldPss0Mask)
	}
}

const (
	RegisterDsipsrFieldUan0Shift = 5
	RegisterDsipsrFieldUan0Mask  = 0x20
)

// GetUan0 UAN0
func (r *RegisterDsipsrType) GetUan0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldUan0Mask) != 0
}

// SetUan0 UAN0
func (r *RegisterDsipsrType) SetUan0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipsrFieldUan0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipsrFieldUan0Mask)
	}
}

const (
	RegisterDsipsrFieldRue0Shift = 6
	RegisterDsipsrFieldRue0Mask  = 0x40
)

// GetRue0 RUE0
func (r *RegisterDsipsrType) GetRue0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldRue0Mask) != 0
}

// SetRue0 RUE0
func (r *RegisterDsipsrType) SetRue0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipsrFieldRue0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipsrFieldRue0Mask)
	}
}

const (
	RegisterDsipsrFieldPss1Shift = 7
	RegisterDsipsrFieldPss1Mask  = 0x80
)

// GetPss1 PSS1
func (r *RegisterDsipsrType) GetPss1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldPss1Mask) != 0
}

// SetPss1 PSS1
func (r *RegisterDsipsrType) SetPss1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipsrFieldPss1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipsrFieldPss1Mask)
	}
}

const (
	RegisterDsipsrFieldUan1Shift = 8
	RegisterDsipsrFieldUan1Mask  = 0x100
)

// GetUan1 UAN1
func (r *RegisterDsipsrType) GetUan1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldUan1Mask) != 0
}

// SetUan1 UAN1
func (r *RegisterDsipsrType) SetUan1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipsrFieldUan1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipsrFieldUan1Mask)
	}
}

// RegisterDsiisr0Type DSI Host interrupt and status register 0
type RegisterDsiisr0Type uint32

func (r *RegisterDsiisr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiisr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiisr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiisr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiisr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiisr0FieldAe0Shift = 0
	RegisterDsiisr0FieldAe0Mask  = 0x1
)

// GetAe0 AE0
func (r *RegisterDsiisr0Type) GetAe0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe0Mask) != 0
}

// SetAe0 AE0
func (r *RegisterDsiisr0Type) SetAe0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe0Mask)
	}
}

const (
	RegisterDsiisr0FieldAe1Shift = 1
	RegisterDsiisr0FieldAe1Mask  = 0x2
)

// GetAe1 AE1
func (r *RegisterDsiisr0Type) GetAe1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe1Mask) != 0
}

// SetAe1 AE1
func (r *RegisterDsiisr0Type) SetAe1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe1Mask)
	}
}

const (
	RegisterDsiisr0FieldAe2Shift = 2
	RegisterDsiisr0FieldAe2Mask  = 0x4
)

// GetAe2 AE2
func (r *RegisterDsiisr0Type) GetAe2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe2Mask) != 0
}

// SetAe2 AE2
func (r *RegisterDsiisr0Type) SetAe2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe2Mask)
	}
}

const (
	RegisterDsiisr0FieldAe3Shift = 3
	RegisterDsiisr0FieldAe3Mask  = 0x8
)

// GetAe3 AE3
func (r *RegisterDsiisr0Type) GetAe3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe3Mask) != 0
}

// SetAe3 AE3
func (r *RegisterDsiisr0Type) SetAe3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe3Mask)
	}
}

const (
	RegisterDsiisr0FieldAe4Shift = 4
	RegisterDsiisr0FieldAe4Mask  = 0x10
)

// GetAe4 AE4
func (r *RegisterDsiisr0Type) GetAe4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe4Mask) != 0
}

// SetAe4 AE4
func (r *RegisterDsiisr0Type) SetAe4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe4Mask)
	}
}

const (
	RegisterDsiisr0FieldAe5Shift = 5
	RegisterDsiisr0FieldAe5Mask  = 0x20
)

// GetAe5 AE5
func (r *RegisterDsiisr0Type) GetAe5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe5Mask) != 0
}

// SetAe5 AE5
func (r *RegisterDsiisr0Type) SetAe5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe5Mask)
	}
}

const (
	RegisterDsiisr0FieldAe6Shift = 6
	RegisterDsiisr0FieldAe6Mask  = 0x40
)

// GetAe6 AE6
func (r *RegisterDsiisr0Type) GetAe6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe6Mask) != 0
}

// SetAe6 AE6
func (r *RegisterDsiisr0Type) SetAe6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe6Mask)
	}
}

const (
	RegisterDsiisr0FieldAe7Shift = 7
	RegisterDsiisr0FieldAe7Mask  = 0x80
)

// GetAe7 AE7
func (r *RegisterDsiisr0Type) GetAe7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe7Mask) != 0
}

// SetAe7 AE7
func (r *RegisterDsiisr0Type) SetAe7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe7Mask)
	}
}

const (
	RegisterDsiisr0FieldAe8Shift = 8
	RegisterDsiisr0FieldAe8Mask  = 0x100
)

// GetAe8 AE8
func (r *RegisterDsiisr0Type) GetAe8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe8Mask) != 0
}

// SetAe8 AE8
func (r *RegisterDsiisr0Type) SetAe8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe8Mask)
	}
}

const (
	RegisterDsiisr0FieldAe9Shift = 9
	RegisterDsiisr0FieldAe9Mask  = 0x200
)

// GetAe9 AE9
func (r *RegisterDsiisr0Type) GetAe9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe9Mask) != 0
}

// SetAe9 AE9
func (r *RegisterDsiisr0Type) SetAe9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe9Mask)
	}
}

const (
	RegisterDsiisr0FieldAe10Shift = 10
	RegisterDsiisr0FieldAe10Mask  = 0x400
)

// GetAe10 AE10
func (r *RegisterDsiisr0Type) GetAe10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe10Mask) != 0
}

// SetAe10 AE10
func (r *RegisterDsiisr0Type) SetAe10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe10Mask)
	}
}

const (
	RegisterDsiisr0FieldAe11Shift = 11
	RegisterDsiisr0FieldAe11Mask  = 0x800
)

// GetAe11 AE11
func (r *RegisterDsiisr0Type) GetAe11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe11Mask) != 0
}

// SetAe11 AE11
func (r *RegisterDsiisr0Type) SetAe11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe11Mask)
	}
}

const (
	RegisterDsiisr0FieldAe12Shift = 12
	RegisterDsiisr0FieldAe12Mask  = 0x1000
)

// GetAe12 AE12
func (r *RegisterDsiisr0Type) GetAe12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe12Mask) != 0
}

// SetAe12 AE12
func (r *RegisterDsiisr0Type) SetAe12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe12Mask)
	}
}

const (
	RegisterDsiisr0FieldAe13Shift = 13
	RegisterDsiisr0FieldAe13Mask  = 0x2000
)

// GetAe13 AE13
func (r *RegisterDsiisr0Type) GetAe13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe13Mask) != 0
}

// SetAe13 AE13
func (r *RegisterDsiisr0Type) SetAe13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe13Mask)
	}
}

const (
	RegisterDsiisr0FieldAe14Shift = 14
	RegisterDsiisr0FieldAe14Mask  = 0x4000
)

// GetAe14 AE14
func (r *RegisterDsiisr0Type) GetAe14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe14Mask) != 0
}

// SetAe14 AE14
func (r *RegisterDsiisr0Type) SetAe14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe14Mask)
	}
}

const (
	RegisterDsiisr0FieldAe15Shift = 15
	RegisterDsiisr0FieldAe15Mask  = 0x8000
)

// GetAe15 AE15
func (r *RegisterDsiisr0Type) GetAe15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe15Mask) != 0
}

// SetAe15 AE15
func (r *RegisterDsiisr0Type) SetAe15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldAe15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldAe15Mask)
	}
}

const (
	RegisterDsiisr0FieldPe0Shift = 16
	RegisterDsiisr0FieldPe0Mask  = 0x10000
)

// GetPe0 PE0
func (r *RegisterDsiisr0Type) GetPe0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldPe0Mask) != 0
}

// SetPe0 PE0
func (r *RegisterDsiisr0Type) SetPe0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldPe0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldPe0Mask)
	}
}

const (
	RegisterDsiisr0FieldPe1Shift = 17
	RegisterDsiisr0FieldPe1Mask  = 0x20000
)

// GetPe1 PE1
func (r *RegisterDsiisr0Type) GetPe1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldPe1Mask) != 0
}

// SetPe1 PE1
func (r *RegisterDsiisr0Type) SetPe1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldPe1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldPe1Mask)
	}
}

const (
	RegisterDsiisr0FieldPe2Shift = 18
	RegisterDsiisr0FieldPe2Mask  = 0x40000
)

// GetPe2 PE2
func (r *RegisterDsiisr0Type) GetPe2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldPe2Mask) != 0
}

// SetPe2 PE2
func (r *RegisterDsiisr0Type) SetPe2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldPe2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldPe2Mask)
	}
}

const (
	RegisterDsiisr0FieldPe3Shift = 19
	RegisterDsiisr0FieldPe3Mask  = 0x80000
)

// GetPe3 PE3
func (r *RegisterDsiisr0Type) GetPe3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldPe3Mask) != 0
}

// SetPe3 PE3
func (r *RegisterDsiisr0Type) SetPe3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldPe3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldPe3Mask)
	}
}

const (
	RegisterDsiisr0FieldPe4Shift = 20
	RegisterDsiisr0FieldPe4Mask  = 0x100000
)

// GetPe4 PE4
func (r *RegisterDsiisr0Type) GetPe4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldPe4Mask) != 0
}

// SetPe4 PE4
func (r *RegisterDsiisr0Type) SetPe4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldPe4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldPe4Mask)
	}
}

// RegisterDsiisr1Type DSI Host interrupt and status register 1
type RegisterDsiisr1Type uint32

func (r *RegisterDsiisr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiisr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiisr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiisr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiisr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiisr1FieldTohstxShift = 0
	RegisterDsiisr1FieldTohstxMask  = 0x1
)

// GetTohstx TOHSTX
func (r *RegisterDsiisr1Type) GetTohstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldTohstxMask) != 0
}

// SetTohstx TOHSTX
func (r *RegisterDsiisr1Type) SetTohstx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldTohstxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldTohstxMask)
	}
}

const (
	RegisterDsiisr1FieldTolprxShift = 1
	RegisterDsiisr1FieldTolprxMask  = 0x2
)

// GetTolprx TOLPRX
func (r *RegisterDsiisr1Type) GetTolprx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldTolprxMask) != 0
}

// SetTolprx TOLPRX
func (r *RegisterDsiisr1Type) SetTolprx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldTolprxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldTolprxMask)
	}
}

const (
	RegisterDsiisr1FieldEccseShift = 2
	RegisterDsiisr1FieldEccseMask  = 0x4
)

// GetEccse ECCSE
func (r *RegisterDsiisr1Type) GetEccse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldEccseMask) != 0
}

// SetEccse ECCSE
func (r *RegisterDsiisr1Type) SetEccse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldEccseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldEccseMask)
	}
}

const (
	RegisterDsiisr1FieldEccmeShift = 3
	RegisterDsiisr1FieldEccmeMask  = 0x8
)

// GetEccme ECCME
func (r *RegisterDsiisr1Type) GetEccme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldEccmeMask) != 0
}

// SetEccme ECCME
func (r *RegisterDsiisr1Type) SetEccme(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldEccmeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldEccmeMask)
	}
}

const (
	RegisterDsiisr1FieldCrceShift = 4
	RegisterDsiisr1FieldCrceMask  = 0x10
)

// GetCrce CRCE
func (r *RegisterDsiisr1Type) GetCrce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldCrceMask) != 0
}

// SetCrce CRCE
func (r *RegisterDsiisr1Type) SetCrce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldCrceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldCrceMask)
	}
}

const (
	RegisterDsiisr1FieldPseShift = 5
	RegisterDsiisr1FieldPseMask  = 0x20
)

// GetPse PSE
func (r *RegisterDsiisr1Type) GetPse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldPseMask) != 0
}

// SetPse PSE
func (r *RegisterDsiisr1Type) SetPse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldPseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldPseMask)
	}
}

const (
	RegisterDsiisr1FieldEotpeShift = 6
	RegisterDsiisr1FieldEotpeMask  = 0x40
)

// GetEotpe EOTPE
func (r *RegisterDsiisr1Type) GetEotpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldEotpeMask) != 0
}

// SetEotpe EOTPE
func (r *RegisterDsiisr1Type) SetEotpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldEotpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldEotpeMask)
	}
}

const (
	RegisterDsiisr1FieldLpwreShift = 7
	RegisterDsiisr1FieldLpwreMask  = 0x80
)

// GetLpwre LPWRE
func (r *RegisterDsiisr1Type) GetLpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldLpwreMask) != 0
}

// SetLpwre LPWRE
func (r *RegisterDsiisr1Type) SetLpwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldLpwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldLpwreMask)
	}
}

const (
	RegisterDsiisr1FieldGcwreShift = 8
	RegisterDsiisr1FieldGcwreMask  = 0x100
)

// GetGcwre GCWRE
func (r *RegisterDsiisr1Type) GetGcwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldGcwreMask) != 0
}

// SetGcwre GCWRE
func (r *RegisterDsiisr1Type) SetGcwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldGcwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldGcwreMask)
	}
}

const (
	RegisterDsiisr1FieldGpwreShift = 9
	RegisterDsiisr1FieldGpwreMask  = 0x200
)

// GetGpwre GPWRE
func (r *RegisterDsiisr1Type) GetGpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldGpwreMask) != 0
}

// SetGpwre GPWRE
func (r *RegisterDsiisr1Type) SetGpwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldGpwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldGpwreMask)
	}
}

const (
	RegisterDsiisr1FieldGptxeShift = 10
	RegisterDsiisr1FieldGptxeMask  = 0x400
)

// GetGptxe GPTXE
func (r *RegisterDsiisr1Type) GetGptxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldGptxeMask) != 0
}

// SetGptxe GPTXE
func (r *RegisterDsiisr1Type) SetGptxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldGptxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldGptxeMask)
	}
}

const (
	RegisterDsiisr1FieldGprdeShift = 11
	RegisterDsiisr1FieldGprdeMask  = 0x800
)

// GetGprde GPRDE
func (r *RegisterDsiisr1Type) GetGprde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldGprdeMask) != 0
}

// SetGprde GPRDE
func (r *RegisterDsiisr1Type) SetGprde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldGprdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldGprdeMask)
	}
}

const (
	RegisterDsiisr1FieldGprxeShift = 12
	RegisterDsiisr1FieldGprxeMask  = 0x1000
)

// GetGprxe GPRXE
func (r *RegisterDsiisr1Type) GetGprxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldGprxeMask) != 0
}

// SetGprxe GPRXE
func (r *RegisterDsiisr1Type) SetGprxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldGprxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldGprxeMask)
	}
}

// RegisterDsiier0Type DSI Host interrupt enable register 0
type RegisterDsiier0Type uint32

func (r *RegisterDsiier0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiier0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiier0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiier0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiier0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiier0FieldAe0ieShift = 0
	RegisterDsiier0FieldAe0ieMask  = 0x1
)

// GetAe0ie AE0IE
func (r *RegisterDsiier0Type) GetAe0ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe0ieMask) != 0
}

// SetAe0ie AE0IE
func (r *RegisterDsiier0Type) SetAe0ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe0ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe0ieMask)
	}
}

const (
	RegisterDsiier0FieldAe1ieShift = 1
	RegisterDsiier0FieldAe1ieMask  = 0x2
)

// GetAe1ie AE1IE
func (r *RegisterDsiier0Type) GetAe1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe1ieMask) != 0
}

// SetAe1ie AE1IE
func (r *RegisterDsiier0Type) SetAe1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe1ieMask)
	}
}

const (
	RegisterDsiier0FieldAe2ieShift = 2
	RegisterDsiier0FieldAe2ieMask  = 0x4
)

// GetAe2ie AE2IE
func (r *RegisterDsiier0Type) GetAe2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe2ieMask) != 0
}

// SetAe2ie AE2IE
func (r *RegisterDsiier0Type) SetAe2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe2ieMask)
	}
}

const (
	RegisterDsiier0FieldAe3ieShift = 3
	RegisterDsiier0FieldAe3ieMask  = 0x8
)

// GetAe3ie AE3IE
func (r *RegisterDsiier0Type) GetAe3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe3ieMask) != 0
}

// SetAe3ie AE3IE
func (r *RegisterDsiier0Type) SetAe3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe3ieMask)
	}
}

const (
	RegisterDsiier0FieldAe4ieShift = 4
	RegisterDsiier0FieldAe4ieMask  = 0x10
)

// GetAe4ie AE4IE
func (r *RegisterDsiier0Type) GetAe4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe4ieMask) != 0
}

// SetAe4ie AE4IE
func (r *RegisterDsiier0Type) SetAe4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe4ieMask)
	}
}

const (
	RegisterDsiier0FieldAe5ieShift = 5
	RegisterDsiier0FieldAe5ieMask  = 0x20
)

// GetAe5ie AE5IE
func (r *RegisterDsiier0Type) GetAe5ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe5ieMask) != 0
}

// SetAe5ie AE5IE
func (r *RegisterDsiier0Type) SetAe5ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe5ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe5ieMask)
	}
}

const (
	RegisterDsiier0FieldAe6ieShift = 6
	RegisterDsiier0FieldAe6ieMask  = 0x40
)

// GetAe6ie AE6IE
func (r *RegisterDsiier0Type) GetAe6ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe6ieMask) != 0
}

// SetAe6ie AE6IE
func (r *RegisterDsiier0Type) SetAe6ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe6ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe6ieMask)
	}
}

const (
	RegisterDsiier0FieldAe7ieShift = 7
	RegisterDsiier0FieldAe7ieMask  = 0x80
)

// GetAe7ie AE7IE
func (r *RegisterDsiier0Type) GetAe7ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe7ieMask) != 0
}

// SetAe7ie AE7IE
func (r *RegisterDsiier0Type) SetAe7ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe7ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe7ieMask)
	}
}

const (
	RegisterDsiier0FieldAe8ieShift = 8
	RegisterDsiier0FieldAe8ieMask  = 0x100
)

// GetAe8ie AE8IE
func (r *RegisterDsiier0Type) GetAe8ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe8ieMask) != 0
}

// SetAe8ie AE8IE
func (r *RegisterDsiier0Type) SetAe8ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe8ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe8ieMask)
	}
}

const (
	RegisterDsiier0FieldAe9ieShift = 9
	RegisterDsiier0FieldAe9ieMask  = 0x200
)

// GetAe9ie AE9IE
func (r *RegisterDsiier0Type) GetAe9ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe9ieMask) != 0
}

// SetAe9ie AE9IE
func (r *RegisterDsiier0Type) SetAe9ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe9ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe9ieMask)
	}
}

const (
	RegisterDsiier0FieldAe10ieShift = 10
	RegisterDsiier0FieldAe10ieMask  = 0x400
)

// GetAe10ie AE10IE
func (r *RegisterDsiier0Type) GetAe10ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe10ieMask) != 0
}

// SetAe10ie AE10IE
func (r *RegisterDsiier0Type) SetAe10ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe10ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe10ieMask)
	}
}

const (
	RegisterDsiier0FieldAe11ieShift = 11
	RegisterDsiier0FieldAe11ieMask  = 0x800
)

// GetAe11ie AE11IE
func (r *RegisterDsiier0Type) GetAe11ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe11ieMask) != 0
}

// SetAe11ie AE11IE
func (r *RegisterDsiier0Type) SetAe11ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe11ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe11ieMask)
	}
}

const (
	RegisterDsiier0FieldAe12ieShift = 12
	RegisterDsiier0FieldAe12ieMask  = 0x1000
)

// GetAe12ie AE12IE
func (r *RegisterDsiier0Type) GetAe12ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe12ieMask) != 0
}

// SetAe12ie AE12IE
func (r *RegisterDsiier0Type) SetAe12ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe12ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe12ieMask)
	}
}

const (
	RegisterDsiier0FieldAe13ieShift = 13
	RegisterDsiier0FieldAe13ieMask  = 0x2000
)

// GetAe13ie AE13IE
func (r *RegisterDsiier0Type) GetAe13ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe13ieMask) != 0
}

// SetAe13ie AE13IE
func (r *RegisterDsiier0Type) SetAe13ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe13ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe13ieMask)
	}
}

const (
	RegisterDsiier0FieldAe14ieShift = 14
	RegisterDsiier0FieldAe14ieMask  = 0x4000
)

// GetAe14ie AE14IE
func (r *RegisterDsiier0Type) GetAe14ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe14ieMask) != 0
}

// SetAe14ie AE14IE
func (r *RegisterDsiier0Type) SetAe14ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe14ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe14ieMask)
	}
}

const (
	RegisterDsiier0FieldAe15ieShift = 15
	RegisterDsiier0FieldAe15ieMask  = 0x8000
)

// GetAe15ie AE15IE
func (r *RegisterDsiier0Type) GetAe15ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe15ieMask) != 0
}

// SetAe15ie AE15IE
func (r *RegisterDsiier0Type) SetAe15ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldAe15ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldAe15ieMask)
	}
}

const (
	RegisterDsiier0FieldPe0ieShift = 16
	RegisterDsiier0FieldPe0ieMask  = 0x10000
)

// GetPe0ie PE0IE
func (r *RegisterDsiier0Type) GetPe0ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldPe0ieMask) != 0
}

// SetPe0ie PE0IE
func (r *RegisterDsiier0Type) SetPe0ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldPe0ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldPe0ieMask)
	}
}

const (
	RegisterDsiier0FieldPe1ieShift = 17
	RegisterDsiier0FieldPe1ieMask  = 0x20000
)

// GetPe1ie PE1IE
func (r *RegisterDsiier0Type) GetPe1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldPe1ieMask) != 0
}

// SetPe1ie PE1IE
func (r *RegisterDsiier0Type) SetPe1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldPe1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldPe1ieMask)
	}
}

const (
	RegisterDsiier0FieldPe2ieShift = 18
	RegisterDsiier0FieldPe2ieMask  = 0x40000
)

// GetPe2ie PE2IE
func (r *RegisterDsiier0Type) GetPe2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldPe2ieMask) != 0
}

// SetPe2ie PE2IE
func (r *RegisterDsiier0Type) SetPe2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldPe2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldPe2ieMask)
	}
}

const (
	RegisterDsiier0FieldPe3ieShift = 19
	RegisterDsiier0FieldPe3ieMask  = 0x80000
)

// GetPe3ie PE3IE
func (r *RegisterDsiier0Type) GetPe3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldPe3ieMask) != 0
}

// SetPe3ie PE3IE
func (r *RegisterDsiier0Type) SetPe3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldPe3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldPe3ieMask)
	}
}

const (
	RegisterDsiier0FieldPe4ieShift = 20
	RegisterDsiier0FieldPe4ieMask  = 0x100000
)

// GetPe4ie PE4IE
func (r *RegisterDsiier0Type) GetPe4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldPe4ieMask) != 0
}

// SetPe4ie PE4IE
func (r *RegisterDsiier0Type) SetPe4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldPe4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldPe4ieMask)
	}
}

// RegisterDsiier1Type DSI Host interrupt enable register 1
type RegisterDsiier1Type uint32

func (r *RegisterDsiier1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiier1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiier1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiier1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiier1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiier1FieldTohstxieShift = 0
	RegisterDsiier1FieldTohstxieMask  = 0x1
)

// GetTohstxie TOHSTXIE
func (r *RegisterDsiier1Type) GetTohstxie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldTohstxieMask) != 0
}

// SetTohstxie TOHSTXIE
func (r *RegisterDsiier1Type) SetTohstxie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldTohstxieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldTohstxieMask)
	}
}

const (
	RegisterDsiier1FieldTolprxieShift = 1
	RegisterDsiier1FieldTolprxieMask  = 0x2
)

// GetTolprxie TOLPRXIE
func (r *RegisterDsiier1Type) GetTolprxie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldTolprxieMask) != 0
}

// SetTolprxie TOLPRXIE
func (r *RegisterDsiier1Type) SetTolprxie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldTolprxieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldTolprxieMask)
	}
}

const (
	RegisterDsiier1FieldEccseieShift = 2
	RegisterDsiier1FieldEccseieMask  = 0x4
)

// GetEccseie ECCSEIE
func (r *RegisterDsiier1Type) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldEccseieMask) != 0
}

// SetEccseie ECCSEIE
func (r *RegisterDsiier1Type) SetEccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldEccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldEccseieMask)
	}
}

const (
	RegisterDsiier1FieldEccmeieShift = 3
	RegisterDsiier1FieldEccmeieMask  = 0x8
)

// GetEccmeie ECCMEIE
func (r *RegisterDsiier1Type) GetEccmeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldEccmeieMask) != 0
}

// SetEccmeie ECCMEIE
func (r *RegisterDsiier1Type) SetEccmeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldEccmeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldEccmeieMask)
	}
}

const (
	RegisterDsiier1FieldCrceieShift = 4
	RegisterDsiier1FieldCrceieMask  = 0x10
)

// GetCrceie CRCEIE
func (r *RegisterDsiier1Type) GetCrceie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldCrceieMask) != 0
}

// SetCrceie CRCEIE
func (r *RegisterDsiier1Type) SetCrceie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldCrceieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldCrceieMask)
	}
}

const (
	RegisterDsiier1FieldPseieShift = 5
	RegisterDsiier1FieldPseieMask  = 0x20
)

// GetPseie PSEIE
func (r *RegisterDsiier1Type) GetPseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldPseieMask) != 0
}

// SetPseie PSEIE
func (r *RegisterDsiier1Type) SetPseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldPseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldPseieMask)
	}
}

const (
	RegisterDsiier1FieldEotpeieShift = 6
	RegisterDsiier1FieldEotpeieMask  = 0x40
)

// GetEotpeie EOTPEIE
func (r *RegisterDsiier1Type) GetEotpeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldEotpeieMask) != 0
}

// SetEotpeie EOTPEIE
func (r *RegisterDsiier1Type) SetEotpeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldEotpeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldEotpeieMask)
	}
}

const (
	RegisterDsiier1FieldLpwreieShift = 7
	RegisterDsiier1FieldLpwreieMask  = 0x80
)

// GetLpwreie LPWREIE
func (r *RegisterDsiier1Type) GetLpwreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldLpwreieMask) != 0
}

// SetLpwreie LPWREIE
func (r *RegisterDsiier1Type) SetLpwreie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldLpwreieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldLpwreieMask)
	}
}

const (
	RegisterDsiier1FieldGcwreieShift = 8
	RegisterDsiier1FieldGcwreieMask  = 0x100
)

// GetGcwreie GCWREIE
func (r *RegisterDsiier1Type) GetGcwreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldGcwreieMask) != 0
}

// SetGcwreie GCWREIE
func (r *RegisterDsiier1Type) SetGcwreie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldGcwreieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldGcwreieMask)
	}
}

const (
	RegisterDsiier1FieldGpwreieShift = 9
	RegisterDsiier1FieldGpwreieMask  = 0x200
)

// GetGpwreie GPWREIE
func (r *RegisterDsiier1Type) GetGpwreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldGpwreieMask) != 0
}

// SetGpwreie GPWREIE
func (r *RegisterDsiier1Type) SetGpwreie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldGpwreieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldGpwreieMask)
	}
}

const (
	RegisterDsiier1FieldGptxeieShift = 10
	RegisterDsiier1FieldGptxeieMask  = 0x400
)

// GetGptxeie GPTXEIE
func (r *RegisterDsiier1Type) GetGptxeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldGptxeieMask) != 0
}

// SetGptxeie GPTXEIE
func (r *RegisterDsiier1Type) SetGptxeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldGptxeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldGptxeieMask)
	}
}

const (
	RegisterDsiier1FieldGprdeieShift = 11
	RegisterDsiier1FieldGprdeieMask  = 0x800
)

// GetGprdeie GPRDEIE
func (r *RegisterDsiier1Type) GetGprdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldGprdeieMask) != 0
}

// SetGprdeie GPRDEIE
func (r *RegisterDsiier1Type) SetGprdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldGprdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldGprdeieMask)
	}
}

const (
	RegisterDsiier1FieldGprxeieShift = 12
	RegisterDsiier1FieldGprxeieMask  = 0x1000
)

// GetGprxeie GPRXEIE
func (r *RegisterDsiier1Type) GetGprxeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldGprxeieMask) != 0
}

// SetGprxeie GPRXEIE
func (r *RegisterDsiier1Type) SetGprxeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldGprxeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldGprxeieMask)
	}
}

// RegisterDsifir0Type DSI Host force interrupt register 0
type RegisterDsifir0Type uint32

func (r *RegisterDsifir0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsifir0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsifir0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsifir0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsifir0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsifir0FieldFae0Shift = 0
	RegisterDsifir0FieldFae0Mask  = 0x1
)

// GetFae0 FAE0
func (r *RegisterDsifir0Type) GetFae0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae0Mask) != 0
}

// SetFae0 FAE0
func (r *RegisterDsifir0Type) SetFae0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae0Mask)
	}
}

const (
	RegisterDsifir0FieldFae1Shift = 1
	RegisterDsifir0FieldFae1Mask  = 0x2
)

// GetFae1 FAE1
func (r *RegisterDsifir0Type) GetFae1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae1Mask) != 0
}

// SetFae1 FAE1
func (r *RegisterDsifir0Type) SetFae1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae1Mask)
	}
}

const (
	RegisterDsifir0FieldFae2Shift = 2
	RegisterDsifir0FieldFae2Mask  = 0x4
)

// GetFae2 FAE2
func (r *RegisterDsifir0Type) GetFae2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae2Mask) != 0
}

// SetFae2 FAE2
func (r *RegisterDsifir0Type) SetFae2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae2Mask)
	}
}

const (
	RegisterDsifir0FieldFae3Shift = 3
	RegisterDsifir0FieldFae3Mask  = 0x8
)

// GetFae3 FAE3
func (r *RegisterDsifir0Type) GetFae3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae3Mask) != 0
}

// SetFae3 FAE3
func (r *RegisterDsifir0Type) SetFae3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae3Mask)
	}
}

const (
	RegisterDsifir0FieldFae4Shift = 4
	RegisterDsifir0FieldFae4Mask  = 0x10
)

// GetFae4 FAE4
func (r *RegisterDsifir0Type) GetFae4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae4Mask) != 0
}

// SetFae4 FAE4
func (r *RegisterDsifir0Type) SetFae4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae4Mask)
	}
}

const (
	RegisterDsifir0FieldFae5Shift = 5
	RegisterDsifir0FieldFae5Mask  = 0x20
)

// GetFae5 FAE5
func (r *RegisterDsifir0Type) GetFae5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae5Mask) != 0
}

// SetFae5 FAE5
func (r *RegisterDsifir0Type) SetFae5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae5Mask)
	}
}

const (
	RegisterDsifir0FieldFae6Shift = 6
	RegisterDsifir0FieldFae6Mask  = 0x40
)

// GetFae6 FAE6
func (r *RegisterDsifir0Type) GetFae6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae6Mask) != 0
}

// SetFae6 FAE6
func (r *RegisterDsifir0Type) SetFae6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae6Mask)
	}
}

const (
	RegisterDsifir0FieldFae7Shift = 7
	RegisterDsifir0FieldFae7Mask  = 0x80
)

// GetFae7 FAE7
func (r *RegisterDsifir0Type) GetFae7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae7Mask) != 0
}

// SetFae7 FAE7
func (r *RegisterDsifir0Type) SetFae7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae7Mask)
	}
}

const (
	RegisterDsifir0FieldFae8Shift = 8
	RegisterDsifir0FieldFae8Mask  = 0x100
)

// GetFae8 FAE8
func (r *RegisterDsifir0Type) GetFae8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae8Mask) != 0
}

// SetFae8 FAE8
func (r *RegisterDsifir0Type) SetFae8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae8Mask)
	}
}

const (
	RegisterDsifir0FieldFae9Shift = 9
	RegisterDsifir0FieldFae9Mask  = 0x200
)

// GetFae9 FAE9
func (r *RegisterDsifir0Type) GetFae9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae9Mask) != 0
}

// SetFae9 FAE9
func (r *RegisterDsifir0Type) SetFae9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae9Mask)
	}
}

const (
	RegisterDsifir0FieldFae10Shift = 10
	RegisterDsifir0FieldFae10Mask  = 0x400
)

// GetFae10 FAE10
func (r *RegisterDsifir0Type) GetFae10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae10Mask) != 0
}

// SetFae10 FAE10
func (r *RegisterDsifir0Type) SetFae10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae10Mask)
	}
}

const (
	RegisterDsifir0FieldFae11Shift = 11
	RegisterDsifir0FieldFae11Mask  = 0x800
)

// GetFae11 FAE11
func (r *RegisterDsifir0Type) GetFae11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae11Mask) != 0
}

// SetFae11 FAE11
func (r *RegisterDsifir0Type) SetFae11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae11Mask)
	}
}

const (
	RegisterDsifir0FieldFae12Shift = 12
	RegisterDsifir0FieldFae12Mask  = 0x1000
)

// GetFae12 FAE12
func (r *RegisterDsifir0Type) GetFae12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae12Mask) != 0
}

// SetFae12 FAE12
func (r *RegisterDsifir0Type) SetFae12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae12Mask)
	}
}

const (
	RegisterDsifir0FieldFae13Shift = 13
	RegisterDsifir0FieldFae13Mask  = 0x2000
)

// GetFae13 FAE13
func (r *RegisterDsifir0Type) GetFae13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae13Mask) != 0
}

// SetFae13 FAE13
func (r *RegisterDsifir0Type) SetFae13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae13Mask)
	}
}

const (
	RegisterDsifir0FieldFae14Shift = 14
	RegisterDsifir0FieldFae14Mask  = 0x4000
)

// GetFae14 FAE14
func (r *RegisterDsifir0Type) GetFae14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae14Mask) != 0
}

// SetFae14 FAE14
func (r *RegisterDsifir0Type) SetFae14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae14Mask)
	}
}

const (
	RegisterDsifir0FieldFae15Shift = 15
	RegisterDsifir0FieldFae15Mask  = 0x8000
)

// GetFae15 FAE15
func (r *RegisterDsifir0Type) GetFae15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae15Mask) != 0
}

// SetFae15 FAE15
func (r *RegisterDsifir0Type) SetFae15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFae15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFae15Mask)
	}
}

const (
	RegisterDsifir0FieldFpe0Shift = 16
	RegisterDsifir0FieldFpe0Mask  = 0x10000
)

// GetFpe0 FPE0
func (r *RegisterDsifir0Type) GetFpe0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFpe0Mask) != 0
}

// SetFpe0 FPE0
func (r *RegisterDsifir0Type) SetFpe0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFpe0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFpe0Mask)
	}
}

const (
	RegisterDsifir0FieldFpe1Shift = 17
	RegisterDsifir0FieldFpe1Mask  = 0x20000
)

// GetFpe1 FPE1
func (r *RegisterDsifir0Type) GetFpe1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFpe1Mask) != 0
}

// SetFpe1 FPE1
func (r *RegisterDsifir0Type) SetFpe1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFpe1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFpe1Mask)
	}
}

const (
	RegisterDsifir0FieldFpe2Shift = 18
	RegisterDsifir0FieldFpe2Mask  = 0x40000
)

// GetFpe2 FPE2
func (r *RegisterDsifir0Type) GetFpe2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFpe2Mask) != 0
}

// SetFpe2 FPE2
func (r *RegisterDsifir0Type) SetFpe2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFpe2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFpe2Mask)
	}
}

const (
	RegisterDsifir0FieldFpe3Shift = 19
	RegisterDsifir0FieldFpe3Mask  = 0x80000
)

// GetFpe3 FPE3
func (r *RegisterDsifir0Type) GetFpe3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFpe3Mask) != 0
}

// SetFpe3 FPE3
func (r *RegisterDsifir0Type) SetFpe3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFpe3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFpe3Mask)
	}
}

const (
	RegisterDsifir0FieldFpe4Shift = 20
	RegisterDsifir0FieldFpe4Mask  = 0x100000
)

// GetFpe4 FPE4
func (r *RegisterDsifir0Type) GetFpe4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFpe4Mask) != 0
}

// SetFpe4 FPE4
func (r *RegisterDsifir0Type) SetFpe4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFpe4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFpe4Mask)
	}
}

// RegisterDsifir1Type DSI Host force interrupt register 1
type RegisterDsifir1Type uint32

func (r *RegisterDsifir1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsifir1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsifir1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsifir1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsifir1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsifir1FieldFtohstxShift = 0
	RegisterDsifir1FieldFtohstxMask  = 0x1
)

// GetFtohstx FTOHSTX
func (r *RegisterDsifir1Type) GetFtohstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFtohstxMask) != 0
}

// SetFtohstx FTOHSTX
func (r *RegisterDsifir1Type) SetFtohstx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFtohstxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFtohstxMask)
	}
}

const (
	RegisterDsifir1FieldFtolprxShift = 1
	RegisterDsifir1FieldFtolprxMask  = 0x2
)

// GetFtolprx FTOLPRX
func (r *RegisterDsifir1Type) GetFtolprx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFtolprxMask) != 0
}

// SetFtolprx FTOLPRX
func (r *RegisterDsifir1Type) SetFtolprx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFtolprxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFtolprxMask)
	}
}

const (
	RegisterDsifir1FieldFeccseShift = 2
	RegisterDsifir1FieldFeccseMask  = 0x4
)

// GetFeccse FECCSE
func (r *RegisterDsifir1Type) GetFeccse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFeccseMask) != 0
}

// SetFeccse FECCSE
func (r *RegisterDsifir1Type) SetFeccse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFeccseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFeccseMask)
	}
}

const (
	RegisterDsifir1FieldFeccmeShift = 3
	RegisterDsifir1FieldFeccmeMask  = 0x8
)

// GetFeccme FECCME
func (r *RegisterDsifir1Type) GetFeccme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFeccmeMask) != 0
}

// SetFeccme FECCME
func (r *RegisterDsifir1Type) SetFeccme(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFeccmeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFeccmeMask)
	}
}

const (
	RegisterDsifir1FieldFcrceShift = 4
	RegisterDsifir1FieldFcrceMask  = 0x10
)

// GetFcrce FCRCE
func (r *RegisterDsifir1Type) GetFcrce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFcrceMask) != 0
}

// SetFcrce FCRCE
func (r *RegisterDsifir1Type) SetFcrce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFcrceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFcrceMask)
	}
}

const (
	RegisterDsifir1FieldFpseShift = 5
	RegisterDsifir1FieldFpseMask  = 0x20
)

// GetFpse FPSE
func (r *RegisterDsifir1Type) GetFpse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFpseMask) != 0
}

// SetFpse FPSE
func (r *RegisterDsifir1Type) SetFpse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFpseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFpseMask)
	}
}

const (
	RegisterDsifir1FieldFeotpeShift = 6
	RegisterDsifir1FieldFeotpeMask  = 0x40
)

// GetFeotpe FEOTPE
func (r *RegisterDsifir1Type) GetFeotpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFeotpeMask) != 0
}

// SetFeotpe FEOTPE
func (r *RegisterDsifir1Type) SetFeotpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFeotpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFeotpeMask)
	}
}

const (
	RegisterDsifir1FieldFlpwreShift = 7
	RegisterDsifir1FieldFlpwreMask  = 0x80
)

// GetFlpwre FLPWRE
func (r *RegisterDsifir1Type) GetFlpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFlpwreMask) != 0
}

// SetFlpwre FLPWRE
func (r *RegisterDsifir1Type) SetFlpwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFlpwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFlpwreMask)
	}
}

const (
	RegisterDsifir1FieldFgcwreShift = 8
	RegisterDsifir1FieldFgcwreMask  = 0x100
)

// GetFgcwre FGCWRE
func (r *RegisterDsifir1Type) GetFgcwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFgcwreMask) != 0
}

// SetFgcwre FGCWRE
func (r *RegisterDsifir1Type) SetFgcwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFgcwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFgcwreMask)
	}
}

const (
	RegisterDsifir1FieldFgpwreShift = 9
	RegisterDsifir1FieldFgpwreMask  = 0x200
)

// GetFgpwre FGPWRE
func (r *RegisterDsifir1Type) GetFgpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFgpwreMask) != 0
}

// SetFgpwre FGPWRE
func (r *RegisterDsifir1Type) SetFgpwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFgpwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFgpwreMask)
	}
}

const (
	RegisterDsifir1FieldFgptxeShift = 10
	RegisterDsifir1FieldFgptxeMask  = 0x400
)

// GetFgptxe FGPTXE
func (r *RegisterDsifir1Type) GetFgptxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFgptxeMask) != 0
}

// SetFgptxe FGPTXE
func (r *RegisterDsifir1Type) SetFgptxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFgptxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFgptxeMask)
	}
}

const (
	RegisterDsifir1FieldFgprdeShift = 11
	RegisterDsifir1FieldFgprdeMask  = 0x800
)

// GetFgprde FGPRDE
func (r *RegisterDsifir1Type) GetFgprde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFgprdeMask) != 0
}

// SetFgprde FGPRDE
func (r *RegisterDsifir1Type) SetFgprde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFgprdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFgprdeMask)
	}
}

const (
	RegisterDsifir1FieldFgprxeShift = 12
	RegisterDsifir1FieldFgprxeMask  = 0x1000
)

// GetFgprxe FGPRXE
func (r *RegisterDsifir1Type) GetFgprxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFgprxeMask) != 0
}

// SetFgprxe FGPRXE
func (r *RegisterDsifir1Type) SetFgprxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFgprxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFgprxeMask)
	}
}

// RegisterDsivscrType DSI Host video shadow control register
type RegisterDsivscrType uint32

func (r *RegisterDsivscrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivscrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivscrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivscrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivscrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivscrFieldEnShift = 0
	RegisterDsivscrFieldEnMask  = 0x1
)

// GetEn EN
func (r *RegisterDsivscrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivscrFieldEnMask) != 0
}

// SetEn EN
func (r *RegisterDsivscrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivscrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivscrFieldEnMask)
	}
}

const (
	RegisterDsivscrFieldUrShift = 8
	RegisterDsivscrFieldUrMask  = 0x100
)

// GetUr UR
func (r *RegisterDsivscrType) GetUr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivscrFieldUrMask) != 0
}

// SetUr UR
func (r *RegisterDsivscrType) SetUr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivscrFieldUrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivscrFieldUrMask)
	}
}

// RegisterDsilcvcidrType DSI Host LTDC current VCID register
type RegisterDsilcvcidrType uint32

func (r *RegisterDsilcvcidrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsilcvcidrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsilcvcidrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsilcvcidrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsilcvcidrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsilcvcidrFieldVcidShift = 0
	RegisterDsilcvcidrFieldVcidMask  = 0x3
)

// GetVcid VCID
func (r *RegisterDsilcvcidrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilcvcidrFieldVcidMask) >> RegisterDsilcvcidrFieldVcidShift)
}

// SetVcid VCID
func (r *RegisterDsilcvcidrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilcvcidrFieldVcidMask)|(uint32(value)<<RegisterDsilcvcidrFieldVcidShift))
}

// RegisterDsilcccrType DSI Host LTDC current color coding register
type RegisterDsilcccrType uint32

func (r *RegisterDsilcccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsilcccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsilcccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsilcccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsilcccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsilcccrFieldColcShift = 0
	RegisterDsilcccrFieldColcMask  = 0xf
)

// GetColc COLC
func (r *RegisterDsilcccrType) GetColc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilcccrFieldColcMask) >> RegisterDsilcccrFieldColcShift)
}

// SetColc COLC
func (r *RegisterDsilcccrType) SetColc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilcccrFieldColcMask)|(uint32(value)<<RegisterDsilcccrFieldColcShift))
}

const (
	RegisterDsilcccrFieldLpeShift = 8
	RegisterDsilcccrFieldLpeMask  = 0x100
)

// GetLpe LPE
func (r *RegisterDsilcccrType) GetLpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsilcccrFieldLpeMask) != 0
}

// SetLpe LPE
func (r *RegisterDsilcccrType) SetLpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsilcccrFieldLpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsilcccrFieldLpeMask)
	}
}

// RegisterDsilpmccrType DSI Host low-power mode current configuration register
type RegisterDsilpmccrType uint32

func (r *RegisterDsilpmccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsilpmccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsilpmccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsilpmccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsilpmccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsilpmccrFieldVlpsizeShift = 0
	RegisterDsilpmccrFieldVlpsizeMask  = 0xff
)

// GetVlpsize VLPSIZE
func (r *RegisterDsilpmccrType) GetVlpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilpmccrFieldVlpsizeMask) >> RegisterDsilpmccrFieldVlpsizeShift)
}

// SetVlpsize VLPSIZE
func (r *RegisterDsilpmccrType) SetVlpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilpmccrFieldVlpsizeMask)|(uint32(value)<<RegisterDsilpmccrFieldVlpsizeShift))
}

const (
	RegisterDsilpmccrFieldLpsizeShift = 16
	RegisterDsilpmccrFieldLpsizeMask  = 0xff0000
)

// GetLpsize LPSIZE
func (r *RegisterDsilpmccrType) GetLpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilpmccrFieldLpsizeMask) >> RegisterDsilpmccrFieldLpsizeShift)
}

// SetLpsize LPSIZE
func (r *RegisterDsilpmccrType) SetLpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilpmccrFieldLpsizeMask)|(uint32(value)<<RegisterDsilpmccrFieldLpsizeShift))
}

// RegisterDsivmccrType DSI Host video mode current configuration register
type RegisterDsivmccrType uint32

func (r *RegisterDsivmccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivmccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivmccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivmccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivmccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivmccrFieldVmtShift = 0
	RegisterDsivmccrFieldVmtMask  = 0x3
)

// GetVmt VMT
func (r *RegisterDsivmccrType) GetVmt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldVmtMask) >> RegisterDsivmccrFieldVmtShift)
}

// SetVmt VMT
func (r *RegisterDsivmccrType) SetVmt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivmccrFieldVmtMask)|(uint32(value)<<RegisterDsivmccrFieldVmtShift))
}

const (
	RegisterDsivmccrFieldLpvsaeShift = 2
	RegisterDsivmccrFieldLpvsaeMask  = 0x4
)

// GetLpvsae LPVSAE
func (r *RegisterDsivmccrType) GetLpvsae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLpvsaeMask) != 0
}

// SetLpvsae LPVSAE
func (r *RegisterDsivmccrType) SetLpvsae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmccrFieldLpvsaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmccrFieldLpvsaeMask)
	}
}

const (
	RegisterDsivmccrFieldLpvbpeShift = 3
	RegisterDsivmccrFieldLpvbpeMask  = 0x8
)

// GetLpvbpe LPVBPE
func (r *RegisterDsivmccrType) GetLpvbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLpvbpeMask) != 0
}

// SetLpvbpe LPVBPE
func (r *RegisterDsivmccrType) SetLpvbpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmccrFieldLpvbpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmccrFieldLpvbpeMask)
	}
}

const (
	RegisterDsivmccrFieldLpvfpeShift = 4
	RegisterDsivmccrFieldLpvfpeMask  = 0x10
)

// GetLpvfpe LPVFPE
func (r *RegisterDsivmccrType) GetLpvfpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLpvfpeMask) != 0
}

// SetLpvfpe LPVFPE
func (r *RegisterDsivmccrType) SetLpvfpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmccrFieldLpvfpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmccrFieldLpvfpeMask)
	}
}

const (
	RegisterDsivmccrFieldLpvaeShift = 5
	RegisterDsivmccrFieldLpvaeMask  = 0x20
)

// GetLpvae LPVAE
func (r *RegisterDsivmccrType) GetLpvae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLpvaeMask) != 0
}

// SetLpvae LPVAE
func (r *RegisterDsivmccrType) SetLpvae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmccrFieldLpvaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmccrFieldLpvaeMask)
	}
}

const (
	RegisterDsivmccrFieldLphbpeShift = 6
	RegisterDsivmccrFieldLphbpeMask  = 0x40
)

// GetLphbpe LPHBPE
func (r *RegisterDsivmccrType) GetLphbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLphbpeMask) != 0
}

// SetLphbpe LPHBPE
func (r *RegisterDsivmccrType) SetLphbpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmccrFieldLphbpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmccrFieldLphbpeMask)
	}
}

const (
	RegisterDsivmccrFieldLphfeShift = 7
	RegisterDsivmccrFieldLphfeMask  = 0x80
)

// GetLphfe LPHFE
func (r *RegisterDsivmccrType) GetLphfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLphfeMask) != 0
}

// SetLphfe LPHFE
func (r *RegisterDsivmccrType) SetLphfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmccrFieldLphfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmccrFieldLphfeMask)
	}
}

const (
	RegisterDsivmccrFieldFbtaaeShift = 8
	RegisterDsivmccrFieldFbtaaeMask  = 0x100
)

// GetFbtaae FBTAAE
func (r *RegisterDsivmccrType) GetFbtaae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldFbtaaeMask) != 0
}

// SetFbtaae FBTAAE
func (r *RegisterDsivmccrType) SetFbtaae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmccrFieldFbtaaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmccrFieldFbtaaeMask)
	}
}

const (
	RegisterDsivmccrFieldLpceShift = 9
	RegisterDsivmccrFieldLpceMask  = 0x200
)

// GetLpce LPCE
func (r *RegisterDsivmccrType) GetLpce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLpceMask) != 0
}

// SetLpce LPCE
func (r *RegisterDsivmccrType) SetLpce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmccrFieldLpceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmccrFieldLpceMask)
	}
}

// RegisterDsivpccrType DSI Host video packet current configuration register
type RegisterDsivpccrType uint32

func (r *RegisterDsivpccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivpccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivpccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivpccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivpccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivpccrFieldVpsizeShift = 0
	RegisterDsivpccrFieldVpsizeMask  = 0x3fff
)

// GetVpsize VPSIZE
func (r *RegisterDsivpccrType) GetVpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivpccrFieldVpsizeMask) >> RegisterDsivpccrFieldVpsizeShift)
}

// SetVpsize VPSIZE
func (r *RegisterDsivpccrType) SetVpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivpccrFieldVpsizeMask)|(uint32(value)<<RegisterDsivpccrFieldVpsizeShift))
}

// RegisterDsivcccrType DSI Host video chunks current configuration register
type RegisterDsivcccrType uint32

func (r *RegisterDsivcccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivcccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivcccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivcccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivcccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivcccrFieldNumcShift = 0
	RegisterDsivcccrFieldNumcMask  = 0x1fff
)

// GetNumc NUMC
func (r *RegisterDsivcccrType) GetNumc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivcccrFieldNumcMask) >> RegisterDsivcccrFieldNumcShift)
}

// SetNumc NUMC
func (r *RegisterDsivcccrType) SetNumc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivcccrFieldNumcMask)|(uint32(value)<<RegisterDsivcccrFieldNumcShift))
}

// RegisterDsivnpccrType DSI Host video null packet current configuration register
type RegisterDsivnpccrType uint32

func (r *RegisterDsivnpccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivnpccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivnpccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivnpccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivnpccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivnpccrFieldNpsizeShift = 0
	RegisterDsivnpccrFieldNpsizeMask  = 0x1fff
)

// GetNpsize NPSIZE
func (r *RegisterDsivnpccrType) GetNpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivnpccrFieldNpsizeMask) >> RegisterDsivnpccrFieldNpsizeShift)
}

// SetNpsize NPSIZE
func (r *RegisterDsivnpccrType) SetNpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivnpccrFieldNpsizeMask)|(uint32(value)<<RegisterDsivnpccrFieldNpsizeShift))
}

// RegisterDsivhsaccrType DSI Host video HSA current configuration register
type RegisterDsivhsaccrType uint32

func (r *RegisterDsivhsaccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivhsaccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivhsaccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivhsaccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivhsaccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivhsaccrFieldHsaShift = 0
	RegisterDsivhsaccrFieldHsaMask  = 0xfff
)

// GetHsa HSA
func (r *RegisterDsivhsaccrType) GetHsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivhsaccrFieldHsaMask) >> RegisterDsivhsaccrFieldHsaShift)
}

// SetHsa HSA
func (r *RegisterDsivhsaccrType) SetHsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivhsaccrFieldHsaMask)|(uint32(value)<<RegisterDsivhsaccrFieldHsaShift))
}

// RegisterDsivhbpccrType DSI Host video HBP current configuration register
type RegisterDsivhbpccrType uint32

func (r *RegisterDsivhbpccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivhbpccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivhbpccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivhbpccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivhbpccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivhbpccrFieldHbpShift = 0
	RegisterDsivhbpccrFieldHbpMask  = 0xfff
)

// GetHbp HBP
func (r *RegisterDsivhbpccrType) GetHbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivhbpccrFieldHbpMask) >> RegisterDsivhbpccrFieldHbpShift)
}

// SetHbp HBP
func (r *RegisterDsivhbpccrType) SetHbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivhbpccrFieldHbpMask)|(uint32(value)<<RegisterDsivhbpccrFieldHbpShift))
}

// RegisterDsivlccrType DSI Host video line current configuration register
type RegisterDsivlccrType uint32

func (r *RegisterDsivlccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivlccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivlccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivlccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivlccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivlccrFieldHlineShift = 0
	RegisterDsivlccrFieldHlineMask  = 0x7fff
)

// GetHline HLINE
func (r *RegisterDsivlccrType) GetHline() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivlccrFieldHlineMask) >> RegisterDsivlccrFieldHlineShift)
}

// SetHline HLINE
func (r *RegisterDsivlccrType) SetHline(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivlccrFieldHlineMask)|(uint32(value)<<RegisterDsivlccrFieldHlineShift))
}

// RegisterDsivvsaccrType DSI Host video VSA current configuration register
type RegisterDsivvsaccrType uint32

func (r *RegisterDsivvsaccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivvsaccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivvsaccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivvsaccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivvsaccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivvsaccrFieldVsaShift = 0
	RegisterDsivvsaccrFieldVsaMask  = 0x3ff
)

// GetVsa VSA
func (r *RegisterDsivvsaccrType) GetVsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvsaccrFieldVsaMask) >> RegisterDsivvsaccrFieldVsaShift)
}

// SetVsa VSA
func (r *RegisterDsivvsaccrType) SetVsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvsaccrFieldVsaMask)|(uint32(value)<<RegisterDsivvsaccrFieldVsaShift))
}

// RegisterDsivvbpccrType DSI Host video VBP current configuration register
type RegisterDsivvbpccrType uint32

func (r *RegisterDsivvbpccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivvbpccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivvbpccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivvbpccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivvbpccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivvbpccrFieldVbpShift = 0
	RegisterDsivvbpccrFieldVbpMask  = 0x3ff
)

// GetVbp VBP
func (r *RegisterDsivvbpccrType) GetVbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvbpccrFieldVbpMask) >> RegisterDsivvbpccrFieldVbpShift)
}

// SetVbp VBP
func (r *RegisterDsivvbpccrType) SetVbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvbpccrFieldVbpMask)|(uint32(value)<<RegisterDsivvbpccrFieldVbpShift))
}

// RegisterDsivvfpccrType DSI Host video VFP current configuration register
type RegisterDsivvfpccrType uint32

func (r *RegisterDsivvfpccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivvfpccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivvfpccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivvfpccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivvfpccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivvfpccrFieldVfpShift = 0
	RegisterDsivvfpccrFieldVfpMask  = 0x3ff
)

// GetVfp VFP
func (r *RegisterDsivvfpccrType) GetVfp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvfpccrFieldVfpMask) >> RegisterDsivvfpccrFieldVfpShift)
}

// SetVfp VFP
func (r *RegisterDsivvfpccrType) SetVfp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvfpccrFieldVfpMask)|(uint32(value)<<RegisterDsivvfpccrFieldVfpShift))
}

// RegisterDsivvaccrType DSI Host video VA current configuration register
type RegisterDsivvaccrType uint32

func (r *RegisterDsivvaccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsivvaccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsivvaccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsivvaccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsivvaccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsivvaccrFieldVaShift = 0
	RegisterDsivvaccrFieldVaMask  = 0x3fff
)

// GetVa VA
func (r *RegisterDsivvaccrType) GetVa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvaccrFieldVaMask) >> RegisterDsivvaccrFieldVaShift)
}

// SetVa VA
func (r *RegisterDsivvaccrType) SetVa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvaccrFieldVaMask)|(uint32(value)<<RegisterDsivvaccrFieldVaShift))
}

// RegisterDsiwcfgrType DSI wrapper configuration register
type RegisterDsiwcfgrType uint32

func (r *RegisterDsiwcfgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiwcfgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiwcfgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiwcfgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiwcfgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiwcfgrFieldDsimShift = 0
	RegisterDsiwcfgrFieldDsimMask  = 0x1
)

// GetDsim DSIM
func (r *RegisterDsiwcfgrType) GetDsim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldDsimMask) != 0
}

// SetDsim DSIM
func (r *RegisterDsiwcfgrType) SetDsim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwcfgrFieldDsimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcfgrFieldDsimMask)
	}
}

const (
	RegisterDsiwcfgrFieldColmuxShift = 1
	RegisterDsiwcfgrFieldColmuxMask  = 0xe
)

// GetColmux COLMUX
func (r *RegisterDsiwcfgrType) GetColmux() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldColmuxMask) >> RegisterDsiwcfgrFieldColmuxShift)
}

// SetColmux COLMUX
func (r *RegisterDsiwcfgrType) SetColmux(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcfgrFieldColmuxMask)|(uint32(value)<<RegisterDsiwcfgrFieldColmuxShift))
}

const (
	RegisterDsiwcfgrFieldTesrcShift = 4
	RegisterDsiwcfgrFieldTesrcMask  = 0x10
)

// GetTesrc TESRC
func (r *RegisterDsiwcfgrType) GetTesrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldTesrcMask) != 0
}

// SetTesrc TESRC
func (r *RegisterDsiwcfgrType) SetTesrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwcfgrFieldTesrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcfgrFieldTesrcMask)
	}
}

const (
	RegisterDsiwcfgrFieldTepolShift = 5
	RegisterDsiwcfgrFieldTepolMask  = 0x20
)

// GetTepol TEPOL
func (r *RegisterDsiwcfgrType) GetTepol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldTepolMask) != 0
}

// SetTepol TEPOL
func (r *RegisterDsiwcfgrType) SetTepol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwcfgrFieldTepolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcfgrFieldTepolMask)
	}
}

const (
	RegisterDsiwcfgrFieldArShift = 6
	RegisterDsiwcfgrFieldArMask  = 0x40
)

// GetAr AR
func (r *RegisterDsiwcfgrType) GetAr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldArMask) != 0
}

// SetAr AR
func (r *RegisterDsiwcfgrType) SetAr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwcfgrFieldArMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcfgrFieldArMask)
	}
}

const (
	RegisterDsiwcfgrFieldVspolShift = 7
	RegisterDsiwcfgrFieldVspolMask  = 0x80
)

// GetVspol VSPOL
func (r *RegisterDsiwcfgrType) GetVspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldVspolMask) != 0
}

// SetVspol VSPOL
func (r *RegisterDsiwcfgrType) SetVspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwcfgrFieldVspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcfgrFieldVspolMask)
	}
}

// RegisterDsiwcrType DSI wrapper control register
type RegisterDsiwcrType uint32

func (r *RegisterDsiwcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiwcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiwcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiwcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiwcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiwcrFieldColmShift = 0
	RegisterDsiwcrFieldColmMask  = 0x1
)

// GetColm COLM
func (r *RegisterDsiwcrType) GetColm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcrFieldColmMask) != 0
}

// SetColm COLM
func (r *RegisterDsiwcrType) SetColm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwcrFieldColmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcrFieldColmMask)
	}
}

const (
	RegisterDsiwcrFieldShtdnShift = 1
	RegisterDsiwcrFieldShtdnMask  = 0x2
)

// GetShtdn SHTDN
func (r *RegisterDsiwcrType) GetShtdn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcrFieldShtdnMask) != 0
}

// SetShtdn SHTDN
func (r *RegisterDsiwcrType) SetShtdn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwcrFieldShtdnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcrFieldShtdnMask)
	}
}

const (
	RegisterDsiwcrFieldLtdcenShift = 2
	RegisterDsiwcrFieldLtdcenMask  = 0x4
)

// GetLtdcen LTDCEN
func (r *RegisterDsiwcrType) GetLtdcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcrFieldLtdcenMask) != 0
}

// SetLtdcen LTDCEN
func (r *RegisterDsiwcrType) SetLtdcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwcrFieldLtdcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcrFieldLtdcenMask)
	}
}

const (
	RegisterDsiwcrFieldDsienShift = 3
	RegisterDsiwcrFieldDsienMask  = 0x8
)

// GetDsien DSIEN
func (r *RegisterDsiwcrType) GetDsien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcrFieldDsienMask) != 0
}

// SetDsien DSIEN
func (r *RegisterDsiwcrType) SetDsien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwcrFieldDsienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcrFieldDsienMask)
	}
}

// RegisterDsiwierType DSI wrapper interrupt enable register
type RegisterDsiwierType uint32

func (r *RegisterDsiwierType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiwierType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiwierType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiwierType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiwierType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiwierFieldTeieShift = 0
	RegisterDsiwierFieldTeieMask  = 0x1
)

// GetTeie TEIE
func (r *RegisterDsiwierType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwierFieldTeieMask) != 0
}

// SetTeie TEIE
func (r *RegisterDsiwierType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwierFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwierFieldTeieMask)
	}
}

const (
	RegisterDsiwierFieldErieShift = 1
	RegisterDsiwierFieldErieMask  = 0x2
)

// GetErie ERIE
func (r *RegisterDsiwierType) GetErie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwierFieldErieMask) != 0
}

// SetErie ERIE
func (r *RegisterDsiwierType) SetErie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwierFieldErieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwierFieldErieMask)
	}
}

const (
	RegisterDsiwierFieldPlllieShift = 9
	RegisterDsiwierFieldPlllieMask  = 0x200
)

// GetPlllie PLLLIE
func (r *RegisterDsiwierType) GetPlllie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwierFieldPlllieMask) != 0
}

// SetPlllie PLLLIE
func (r *RegisterDsiwierType) SetPlllie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwierFieldPlllieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwierFieldPlllieMask)
	}
}

const (
	RegisterDsiwierFieldPlluieShift = 10
	RegisterDsiwierFieldPlluieMask  = 0x400
)

// GetPlluie PLLUIE
func (r *RegisterDsiwierType) GetPlluie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwierFieldPlluieMask) != 0
}

// SetPlluie PLLUIE
func (r *RegisterDsiwierType) SetPlluie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwierFieldPlluieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwierFieldPlluieMask)
	}
}

const (
	RegisterDsiwierFieldRrieShift = 13
	RegisterDsiwierFieldRrieMask  = 0x2000
)

// GetRrie RRIE
func (r *RegisterDsiwierType) GetRrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwierFieldRrieMask) != 0
}

// SetRrie RRIE
func (r *RegisterDsiwierType) SetRrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwierFieldRrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwierFieldRrieMask)
	}
}

// RegisterDsiwisrType DSI wrapper interrupt and status register
type RegisterDsiwisrType uint32

func (r *RegisterDsiwisrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiwisrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiwisrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiwisrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiwisrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiwisrFieldTeifShift = 0
	RegisterDsiwisrFieldTeifMask  = 0x1
)

// GetTeif TEIF
func (r *RegisterDsiwisrType) GetTeif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldTeifMask) != 0
}

// SetTeif TEIF
func (r *RegisterDsiwisrType) SetTeif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwisrFieldTeifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwisrFieldTeifMask)
	}
}

const (
	RegisterDsiwisrFieldErifShift = 1
	RegisterDsiwisrFieldErifMask  = 0x2
)

// GetErif ERIF
func (r *RegisterDsiwisrType) GetErif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldErifMask) != 0
}

// SetErif ERIF
func (r *RegisterDsiwisrType) SetErif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwisrFieldErifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwisrFieldErifMask)
	}
}

const (
	RegisterDsiwisrFieldBusyShift = 2
	RegisterDsiwisrFieldBusyMask  = 0x4
)

// GetBusy BUSY
func (r *RegisterDsiwisrType) GetBusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldBusyMask) != 0
}

// SetBusy BUSY
func (r *RegisterDsiwisrType) SetBusy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwisrFieldBusyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwisrFieldBusyMask)
	}
}

const (
	RegisterDsiwisrFieldPlllsShift = 8
	RegisterDsiwisrFieldPlllsMask  = 0x100
)

// GetPllls PLLLS
func (r *RegisterDsiwisrType) GetPllls() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldPlllsMask) != 0
}

// SetPllls PLLLS
func (r *RegisterDsiwisrType) SetPllls(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwisrFieldPlllsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwisrFieldPlllsMask)
	}
}

const (
	RegisterDsiwisrFieldPlllifShift = 9
	RegisterDsiwisrFieldPlllifMask  = 0x200
)

// GetPlllif PLLLIF
func (r *RegisterDsiwisrType) GetPlllif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldPlllifMask) != 0
}

// SetPlllif PLLLIF
func (r *RegisterDsiwisrType) SetPlllif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwisrFieldPlllifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwisrFieldPlllifMask)
	}
}

const (
	RegisterDsiwisrFieldPlluifShift = 10
	RegisterDsiwisrFieldPlluifMask  = 0x400
)

// GetPlluif PLLUIF
func (r *RegisterDsiwisrType) GetPlluif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldPlluifMask) != 0
}

// SetPlluif PLLUIF
func (r *RegisterDsiwisrType) SetPlluif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwisrFieldPlluifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwisrFieldPlluifMask)
	}
}

const (
	RegisterDsiwisrFieldRrsShift = 12
	RegisterDsiwisrFieldRrsMask  = 0x1000
)

// GetRrs RRS
func (r *RegisterDsiwisrType) GetRrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldRrsMask) != 0
}

// SetRrs RRS
func (r *RegisterDsiwisrType) SetRrs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwisrFieldRrsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwisrFieldRrsMask)
	}
}

const (
	RegisterDsiwisrFieldRrifShift = 13
	RegisterDsiwisrFieldRrifMask  = 0x2000
)

// GetRrif RRIF
func (r *RegisterDsiwisrType) GetRrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldRrifMask) != 0
}

// SetRrif RRIF
func (r *RegisterDsiwisrType) SetRrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwisrFieldRrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwisrFieldRrifMask)
	}
}

// RegisterDsiwifcrType DSI wrapper interrupt flag clear register
type RegisterDsiwifcrType uint32

func (r *RegisterDsiwifcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiwifcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiwifcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiwifcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiwifcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiwifcrFieldCteifShift = 0
	RegisterDsiwifcrFieldCteifMask  = 0x1
)

// GetCteif CTEIF
func (r *RegisterDsiwifcrType) GetCteif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwifcrFieldCteifMask) != 0
}

// SetCteif CTEIF
func (r *RegisterDsiwifcrType) SetCteif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwifcrFieldCteifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwifcrFieldCteifMask)
	}
}

const (
	RegisterDsiwifcrFieldCerifShift = 1
	RegisterDsiwifcrFieldCerifMask  = 0x2
)

// GetCerif CERIF
func (r *RegisterDsiwifcrType) GetCerif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwifcrFieldCerifMask) != 0
}

// SetCerif CERIF
func (r *RegisterDsiwifcrType) SetCerif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwifcrFieldCerifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwifcrFieldCerifMask)
	}
}

const (
	RegisterDsiwifcrFieldCplllifShift = 9
	RegisterDsiwifcrFieldCplllifMask  = 0x200
)

// GetCplllif CPLLLIF
func (r *RegisterDsiwifcrType) GetCplllif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwifcrFieldCplllifMask) != 0
}

// SetCplllif CPLLLIF
func (r *RegisterDsiwifcrType) SetCplllif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwifcrFieldCplllifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwifcrFieldCplllifMask)
	}
}

const (
	RegisterDsiwifcrFieldCplluifShift = 10
	RegisterDsiwifcrFieldCplluifMask  = 0x400
)

// GetCplluif CPLLUIF
func (r *RegisterDsiwifcrType) GetCplluif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwifcrFieldCplluifMask) != 0
}

// SetCplluif CPLLUIF
func (r *RegisterDsiwifcrType) SetCplluif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwifcrFieldCplluifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwifcrFieldCplluifMask)
	}
}

const (
	RegisterDsiwifcrFieldCrrifShift = 13
	RegisterDsiwifcrFieldCrrifMask  = 0x2000
)

// GetCrrif CRRIF
func (r *RegisterDsiwifcrType) GetCrrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwifcrFieldCrrifMask) != 0
}

// SetCrrif CRRIF
func (r *RegisterDsiwifcrType) SetCrrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwifcrFieldCrrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwifcrFieldCrrifMask)
	}
}

// RegisterDsiwpcr0Type DSI wrapper PHY configuration register 0
type RegisterDsiwpcr0Type uint32

func (r *RegisterDsiwpcr0Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiwpcr0Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiwpcr0Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiwpcr0Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiwpcr0Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiwpcr0FieldUix4Shift = 0
	RegisterDsiwpcr0FieldUix4Mask  = 0x3f
)

// GetUix4 UIX4
func (r *RegisterDsiwpcr0Type) GetUix4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldUix4Mask) >> RegisterDsiwpcr0FieldUix4Shift)
}

// SetUix4 UIX4
func (r *RegisterDsiwpcr0Type) SetUix4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldUix4Mask)|(uint32(value)<<RegisterDsiwpcr0FieldUix4Shift))
}

const (
	RegisterDsiwpcr0FieldSwclShift = 6
	RegisterDsiwpcr0FieldSwclMask  = 0x40
)

// GetSwcl SWCL
func (r *RegisterDsiwpcr0Type) GetSwcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldSwclMask) != 0
}

// SetSwcl SWCL
func (r *RegisterDsiwpcr0Type) SetSwcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldSwclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldSwclMask)
	}
}

const (
	RegisterDsiwpcr0FieldSwdl0Shift = 7
	RegisterDsiwpcr0FieldSwdl0Mask  = 0x80
)

// GetSwdl0 SWDL0
func (r *RegisterDsiwpcr0Type) GetSwdl0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldSwdl0Mask) != 0
}

// SetSwdl0 SWDL0
func (r *RegisterDsiwpcr0Type) SetSwdl0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldSwdl0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldSwdl0Mask)
	}
}

const (
	RegisterDsiwpcr0FieldSwdl1Shift = 8
	RegisterDsiwpcr0FieldSwdl1Mask  = 0x100
)

// GetSwdl1 SWDL1
func (r *RegisterDsiwpcr0Type) GetSwdl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldSwdl1Mask) != 0
}

// SetSwdl1 SWDL1
func (r *RegisterDsiwpcr0Type) SetSwdl1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldSwdl1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldSwdl1Mask)
	}
}

const (
	RegisterDsiwpcr0FieldHsiclShift = 9
	RegisterDsiwpcr0FieldHsiclMask  = 0x200
)

// GetHsicl HSICL
func (r *RegisterDsiwpcr0Type) GetHsicl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldHsiclMask) != 0
}

// SetHsicl HSICL
func (r *RegisterDsiwpcr0Type) SetHsicl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldHsiclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldHsiclMask)
	}
}

const (
	RegisterDsiwpcr0FieldHsidl0Shift = 10
	RegisterDsiwpcr0FieldHsidl0Mask  = 0x400
)

// GetHsidl0 HSIDL0
func (r *RegisterDsiwpcr0Type) GetHsidl0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldHsidl0Mask) != 0
}

// SetHsidl0 HSIDL0
func (r *RegisterDsiwpcr0Type) SetHsidl0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldHsidl0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldHsidl0Mask)
	}
}

const (
	RegisterDsiwpcr0FieldHsidl1Shift = 11
	RegisterDsiwpcr0FieldHsidl1Mask  = 0x800
)

// GetHsidl1 HSIDL1
func (r *RegisterDsiwpcr0Type) GetHsidl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldHsidl1Mask) != 0
}

// SetHsidl1 HSIDL1
func (r *RegisterDsiwpcr0Type) SetHsidl1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldHsidl1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldHsidl1Mask)
	}
}

const (
	RegisterDsiwpcr0FieldFtxsmclShift = 12
	RegisterDsiwpcr0FieldFtxsmclMask  = 0x1000
)

// GetFtxsmcl FTXSMCL
func (r *RegisterDsiwpcr0Type) GetFtxsmcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldFtxsmclMask) != 0
}

// SetFtxsmcl FTXSMCL
func (r *RegisterDsiwpcr0Type) SetFtxsmcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldFtxsmclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldFtxsmclMask)
	}
}

const (
	RegisterDsiwpcr0FieldFtxsmdlShift = 13
	RegisterDsiwpcr0FieldFtxsmdlMask  = 0x2000
)

// GetFtxsmdl FTXSMDL
func (r *RegisterDsiwpcr0Type) GetFtxsmdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldFtxsmdlMask) != 0
}

// SetFtxsmdl FTXSMDL
func (r *RegisterDsiwpcr0Type) SetFtxsmdl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldFtxsmdlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldFtxsmdlMask)
	}
}

const (
	RegisterDsiwpcr0FieldCdoffdlShift = 14
	RegisterDsiwpcr0FieldCdoffdlMask  = 0x4000
)

// GetCdoffdl CDOFFDL
func (r *RegisterDsiwpcr0Type) GetCdoffdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldCdoffdlMask) != 0
}

// SetCdoffdl CDOFFDL
func (r *RegisterDsiwpcr0Type) SetCdoffdl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldCdoffdlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldCdoffdlMask)
	}
}

const (
	RegisterDsiwpcr0FieldTddlShift = 16
	RegisterDsiwpcr0FieldTddlMask  = 0x10000
)

// GetTddl TDDL
func (r *RegisterDsiwpcr0Type) GetTddl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTddlMask) != 0
}

// SetTddl TDDL
func (r *RegisterDsiwpcr0Type) SetTddl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldTddlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldTddlMask)
	}
}

const (
	RegisterDsiwpcr0FieldPdenShift = 18
	RegisterDsiwpcr0FieldPdenMask  = 0x40000
)

// GetPden Pull-down enable
func (r *RegisterDsiwpcr0Type) GetPden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldPdenMask) != 0
}

// SetPden Pull-down enable
func (r *RegisterDsiwpcr0Type) SetPden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldPdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldPdenMask)
	}
}

const (
	RegisterDsiwpcr0FieldTclkprepenShift = 19
	RegisterDsiwpcr0FieldTclkprepenMask  = 0x80000
)

// GetTclkprepen Custom time for tCLK-PREPARE enable
func (r *RegisterDsiwpcr0Type) GetTclkprepen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTclkprepenMask) != 0
}

// SetTclkprepen Custom time for tCLK-PREPARE enable
func (r *RegisterDsiwpcr0Type) SetTclkprepen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldTclkprepenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldTclkprepenMask)
	}
}

const (
	RegisterDsiwpcr0FieldTclkzeroenShift = 20
	RegisterDsiwpcr0FieldTclkzeroenMask  = 0x100000
)

// GetTclkzeroen Custom time for tCLK-ZERO enable
func (r *RegisterDsiwpcr0Type) GetTclkzeroen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTclkzeroenMask) != 0
}

// SetTclkzeroen Custom time for tCLK-ZERO enable
func (r *RegisterDsiwpcr0Type) SetTclkzeroen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldTclkzeroenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldTclkzeroenMask)
	}
}

const (
	RegisterDsiwpcr0FieldThsprepenShift = 21
	RegisterDsiwpcr0FieldThsprepenMask  = 0x200000
)

// GetThsprepen Custom time for tHS-PREPARE enable
func (r *RegisterDsiwpcr0Type) GetThsprepen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldThsprepenMask) != 0
}

// SetThsprepen Custom time for tHS-PREPARE enable
func (r *RegisterDsiwpcr0Type) SetThsprepen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldThsprepenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldThsprepenMask)
	}
}

const (
	RegisterDsiwpcr0FieldThstrailenShift = 22
	RegisterDsiwpcr0FieldThstrailenMask  = 0x400000
)

// GetThstrailen Custom time for tHS-TRAIL enable
func (r *RegisterDsiwpcr0Type) GetThstrailen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldThstrailenMask) != 0
}

// SetThstrailen Custom time for tHS-TRAIL enable
func (r *RegisterDsiwpcr0Type) SetThstrailen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldThstrailenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldThstrailenMask)
	}
}

const (
	RegisterDsiwpcr0FieldThszeroenShift = 23
	RegisterDsiwpcr0FieldThszeroenMask  = 0x800000
)

// GetThszeroen Custom time for tHS-ZERO enable
func (r *RegisterDsiwpcr0Type) GetThszeroen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldThszeroenMask) != 0
}

// SetThszeroen Custom time for tHS-ZERO enable
func (r *RegisterDsiwpcr0Type) SetThszeroen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldThszeroenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldThszeroenMask)
	}
}

const (
	RegisterDsiwpcr0FieldTlpxdenShift = 24
	RegisterDsiwpcr0FieldTlpxdenMask  = 0x1000000
)

// GetTlpxden Custom time for tLPX for data lanes enable
func (r *RegisterDsiwpcr0Type) GetTlpxden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTlpxdenMask) != 0
}

// SetTlpxden Custom time for tLPX for data lanes enable
func (r *RegisterDsiwpcr0Type) SetTlpxden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldTlpxdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldTlpxdenMask)
	}
}

const (
	RegisterDsiwpcr0FieldThsexitenShift = 25
	RegisterDsiwpcr0FieldThsexitenMask  = 0x2000000
)

// GetThsexiten Custom time for tHS-EXIT enable
func (r *RegisterDsiwpcr0Type) GetThsexiten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldThsexitenMask) != 0
}

// SetThsexiten Custom time for tHS-EXIT enable
func (r *RegisterDsiwpcr0Type) SetThsexiten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldThsexitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldThsexitenMask)
	}
}

const (
	RegisterDsiwpcr0FieldTlpxcenShift = 26
	RegisterDsiwpcr0FieldTlpxcenMask  = 0x4000000
)

// GetTlpxcen Custom time for tLPX for clock lane enable
func (r *RegisterDsiwpcr0Type) GetTlpxcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTlpxcenMask) != 0
}

// SetTlpxcen Custom time for tLPX for clock lane enable
func (r *RegisterDsiwpcr0Type) SetTlpxcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldTlpxcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldTlpxcenMask)
	}
}

const (
	RegisterDsiwpcr0FieldTclkpostenShift = 27
	RegisterDsiwpcr0FieldTclkpostenMask  = 0x8000000
)

// GetTclkposten Custom time for tCLK-POST enable
func (r *RegisterDsiwpcr0Type) GetTclkposten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTclkpostenMask) != 0
}

// SetTclkposten Custom time for tCLK-POST enable
func (r *RegisterDsiwpcr0Type) SetTclkposten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldTclkpostenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldTclkpostenMask)
	}
}

// RegisterDsiwpcr1Type This register shall be programmed only when DSI is stopped (CR. DSIEN=0 and CR.EN = 0).
type RegisterDsiwpcr1Type uint32

func (r *RegisterDsiwpcr1Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiwpcr1Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiwpcr1Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiwpcr1Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiwpcr1Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiwpcr1FieldHstxdclShift = 0
	RegisterDsiwpcr1FieldHstxdclMask  = 0x3
)

// GetHstxdcl High-speed transmission delay on clock lane
func (r *RegisterDsiwpcr1Type) GetHstxdcl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldHstxdclMask) >> RegisterDsiwpcr1FieldHstxdclShift)
}

// SetHstxdcl High-speed transmission delay on clock lane
func (r *RegisterDsiwpcr1Type) SetHstxdcl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldHstxdclMask)|(uint32(value)<<RegisterDsiwpcr1FieldHstxdclShift))
}

const (
	RegisterDsiwpcr1FieldHstxddlShift = 2
	RegisterDsiwpcr1FieldHstxddlMask  = 0xc
)

// GetHstxddl High-speed transmission delay on data lanes
func (r *RegisterDsiwpcr1Type) GetHstxddl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldHstxddlMask) >> RegisterDsiwpcr1FieldHstxddlShift)
}

// SetHstxddl High-speed transmission delay on data lanes
func (r *RegisterDsiwpcr1Type) SetHstxddl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldHstxddlMask)|(uint32(value)<<RegisterDsiwpcr1FieldHstxddlShift))
}

const (
	RegisterDsiwpcr1FieldLpsrcclShift = 6
	RegisterDsiwpcr1FieldLpsrcclMask  = 0xc0
)

// GetLpsrccl Low-power transmission slew-rate compensation on clock lane
func (r *RegisterDsiwpcr1Type) GetLpsrccl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldLpsrcclMask) >> RegisterDsiwpcr1FieldLpsrcclShift)
}

// SetLpsrccl Low-power transmission slew-rate compensation on clock lane
func (r *RegisterDsiwpcr1Type) SetLpsrccl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldLpsrcclMask)|(uint32(value)<<RegisterDsiwpcr1FieldLpsrcclShift))
}

const (
	RegisterDsiwpcr1FieldLpsrcdlShift = 8
	RegisterDsiwpcr1FieldLpsrcdlMask  = 0x300
)

// GetLpsrcdl Low-power transmission slew-rate compensation on data lanes
func (r *RegisterDsiwpcr1Type) GetLpsrcdl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldLpsrcdlMask) >> RegisterDsiwpcr1FieldLpsrcdlShift)
}

// SetLpsrcdl Low-power transmission slew-rate compensation on data lanes
func (r *RegisterDsiwpcr1Type) SetLpsrcdl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldLpsrcdlMask)|(uint32(value)<<RegisterDsiwpcr1FieldLpsrcdlShift))
}

const (
	RegisterDsiwpcr1FieldSddcShift = 12
	RegisterDsiwpcr1FieldSddcMask  = 0x1000
)

// GetSddc SDD control
func (r *RegisterDsiwpcr1Type) GetSddc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldSddcMask) != 0
}

// SetSddc SDD control
func (r *RegisterDsiwpcr1Type) SetSddc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr1FieldSddcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldSddcMask)
	}
}

const (
	RegisterDsiwpcr1FieldHstxsrcclShift = 16
	RegisterDsiwpcr1FieldHstxsrcclMask  = 0x30000
)

// GetHstxsrccl High-speed transmission slew-rate control on clock lane
func (r *RegisterDsiwpcr1Type) GetHstxsrccl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldHstxsrcclMask) >> RegisterDsiwpcr1FieldHstxsrcclShift)
}

// SetHstxsrccl High-speed transmission slew-rate control on clock lane
func (r *RegisterDsiwpcr1Type) SetHstxsrccl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldHstxsrcclMask)|(uint32(value)<<RegisterDsiwpcr1FieldHstxsrcclShift))
}

const (
	RegisterDsiwpcr1FieldHstxsrcdlShift = 18
	RegisterDsiwpcr1FieldHstxsrcdlMask  = 0xc0000
)

// GetHstxsrcdl High-speed transmission slew-rate control on data lanes
func (r *RegisterDsiwpcr1Type) GetHstxsrcdl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldHstxsrcdlMask) >> RegisterDsiwpcr1FieldHstxsrcdlShift)
}

// SetHstxsrcdl High-speed transmission slew-rate control on data lanes
func (r *RegisterDsiwpcr1Type) SetHstxsrcdl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldHstxsrcdlMask)|(uint32(value)<<RegisterDsiwpcr1FieldHstxsrcdlShift))
}

const (
	RegisterDsiwpcr1FieldFlprxlpmShift = 22
	RegisterDsiwpcr1FieldFlprxlpmMask  = 0x400000
)

// GetFlprxlpm Forces LP receiver in low-power mode
func (r *RegisterDsiwpcr1Type) GetFlprxlpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldFlprxlpmMask) != 0
}

// SetFlprxlpm Forces LP receiver in low-power mode
func (r *RegisterDsiwpcr1Type) SetFlprxlpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr1FieldFlprxlpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldFlprxlpmMask)
	}
}

const (
	RegisterDsiwpcr1FieldLprxftShift = 25
	RegisterDsiwpcr1FieldLprxftMask  = 0x6000000
)

// GetLprxft Low-power RX low-pass filtering tuning
func (r *RegisterDsiwpcr1Type) GetLprxft() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldLprxftMask) >> RegisterDsiwpcr1FieldLprxftShift)
}

// SetLprxft Low-power RX low-pass filtering tuning
func (r *RegisterDsiwpcr1Type) SetLprxft(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldLprxftMask)|(uint32(value)<<RegisterDsiwpcr1FieldLprxftShift))
}

// RegisterDsiwpcr2Type DSI wrapper PHY configuration register 2
type RegisterDsiwpcr2Type uint32

func (r *RegisterDsiwpcr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiwpcr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiwpcr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiwpcr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiwpcr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiwpcr2FieldTclkprepShift = 0
	RegisterDsiwpcr2FieldTclkprepMask  = 0xff
)

// GetTclkprep TCLKPREP
func (r *RegisterDsiwpcr2Type) GetTclkprep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr2FieldTclkprepMask) >> RegisterDsiwpcr2FieldTclkprepShift)
}

// SetTclkprep TCLKPREP
func (r *RegisterDsiwpcr2Type) SetTclkprep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr2FieldTclkprepMask)|(uint32(value)<<RegisterDsiwpcr2FieldTclkprepShift))
}

const (
	RegisterDsiwpcr2FieldTclkzeroShift = 8
	RegisterDsiwpcr2FieldTclkzeroMask  = 0xff00
)

// GetTclkzero TCLKZERO
func (r *RegisterDsiwpcr2Type) GetTclkzero() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr2FieldTclkzeroMask) >> RegisterDsiwpcr2FieldTclkzeroShift)
}

// SetTclkzero TCLKZERO
func (r *RegisterDsiwpcr2Type) SetTclkzero(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr2FieldTclkzeroMask)|(uint32(value)<<RegisterDsiwpcr2FieldTclkzeroShift))
}

const (
	RegisterDsiwpcr2FieldThsprepShift = 16
	RegisterDsiwpcr2FieldThsprepMask  = 0xff0000
)

// GetThsprep THSPREP
func (r *RegisterDsiwpcr2Type) GetThsprep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr2FieldThsprepMask) >> RegisterDsiwpcr2FieldThsprepShift)
}

// SetThsprep THSPREP
func (r *RegisterDsiwpcr2Type) SetThsprep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr2FieldThsprepMask)|(uint32(value)<<RegisterDsiwpcr2FieldThsprepShift))
}

const (
	RegisterDsiwpcr2FieldThstrailShift = 24
	RegisterDsiwpcr2FieldThstrailMask  = 0xff000000
)

// GetThstrail THSTRAIL
func (r *RegisterDsiwpcr2Type) GetThstrail() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr2FieldThstrailMask) >> RegisterDsiwpcr2FieldThstrailShift)
}

// SetThstrail THSTRAIL
func (r *RegisterDsiwpcr2Type) SetThstrail(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr2FieldThstrailMask)|(uint32(value)<<RegisterDsiwpcr2FieldThstrailShift))
}

// RegisterDsiwpcr3Type DSI wrapper PHY configuration register 3
type RegisterDsiwpcr3Type uint32

func (r *RegisterDsiwpcr3Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiwpcr3Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiwpcr3Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiwpcr3Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiwpcr3Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiwpcr3FieldThszeroShift = 0
	RegisterDsiwpcr3FieldThszeroMask  = 0xff
)

// GetThszero THSZERO
func (r *RegisterDsiwpcr3Type) GetThszero() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr3FieldThszeroMask) >> RegisterDsiwpcr3FieldThszeroShift)
}

// SetThszero THSZERO
func (r *RegisterDsiwpcr3Type) SetThszero(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr3FieldThszeroMask)|(uint32(value)<<RegisterDsiwpcr3FieldThszeroShift))
}

const (
	RegisterDsiwpcr3FieldTlpxdShift = 8
	RegisterDsiwpcr3FieldTlpxdMask  = 0xff00
)

// GetTlpxd TLPXD
func (r *RegisterDsiwpcr3Type) GetTlpxd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr3FieldTlpxdMask) >> RegisterDsiwpcr3FieldTlpxdShift)
}

// SetTlpxd TLPXD
func (r *RegisterDsiwpcr3Type) SetTlpxd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr3FieldTlpxdMask)|(uint32(value)<<RegisterDsiwpcr3FieldTlpxdShift))
}

const (
	RegisterDsiwpcr3FieldThsexitShift = 16
	RegisterDsiwpcr3FieldThsexitMask  = 0xff0000
)

// GetThsexit THSEXIT
func (r *RegisterDsiwpcr3Type) GetThsexit() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr3FieldThsexitMask) >> RegisterDsiwpcr3FieldThsexitShift)
}

// SetThsexit THSEXIT
func (r *RegisterDsiwpcr3Type) SetThsexit(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr3FieldThsexitMask)|(uint32(value)<<RegisterDsiwpcr3FieldThsexitShift))
}

const (
	RegisterDsiwpcr3FieldTlpxcShift = 24
	RegisterDsiwpcr3FieldTlpxcMask  = 0xff000000
)

// GetTlpxc TLPXC
func (r *RegisterDsiwpcr3Type) GetTlpxc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr3FieldTlpxcMask) >> RegisterDsiwpcr3FieldTlpxcShift)
}

// SetTlpxc TLPXC
func (r *RegisterDsiwpcr3Type) SetTlpxc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr3FieldTlpxcMask)|(uint32(value)<<RegisterDsiwpcr3FieldTlpxcShift))
}

// RegisterDsiwpcr4Type DSI wrapper PHY configuration register 4
type RegisterDsiwpcr4Type uint32

func (r *RegisterDsiwpcr4Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiwpcr4Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiwpcr4Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiwpcr4Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiwpcr4Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiwpcr4FieldTclkpostShift = 0
	RegisterDsiwpcr4FieldTclkpostMask  = 0xff
)

// GetTclkpost TCLKPOST
func (r *RegisterDsiwpcr4Type) GetTclkpost() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr4FieldTclkpostMask) >> RegisterDsiwpcr4FieldTclkpostShift)
}

// SetTclkpost TCLKPOST
func (r *RegisterDsiwpcr4Type) SetTclkpost(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr4FieldTclkpostMask)|(uint32(value)<<RegisterDsiwpcr4FieldTclkpostShift))
}

// RegisterDsiwrpcrType DSI wrapper regulator and PLL control register
type RegisterDsiwrpcrType uint32

func (r *RegisterDsiwrpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDsiwrpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDsiwrpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDsiwrpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDsiwrpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDsiwrpcrFieldPllenShift = 0
	RegisterDsiwrpcrFieldPllenMask  = 0x1
)

// GetPllen PLLEN
func (r *RegisterDsiwrpcrType) GetPllen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwrpcrFieldPllenMask) != 0
}

// SetPllen PLLEN
func (r *RegisterDsiwrpcrType) SetPllen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwrpcrFieldPllenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwrpcrFieldPllenMask)
	}
}

const (
	RegisterDsiwrpcrFieldNdivShift = 2
	RegisterDsiwrpcrFieldNdivMask  = 0x1fc
)

// GetNdiv NDIV
func (r *RegisterDsiwrpcrType) GetNdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwrpcrFieldNdivMask) >> RegisterDsiwrpcrFieldNdivShift)
}

// SetNdiv NDIV
func (r *RegisterDsiwrpcrType) SetNdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwrpcrFieldNdivMask)|(uint32(value)<<RegisterDsiwrpcrFieldNdivShift))
}

const (
	RegisterDsiwrpcrFieldIdfShift = 11
	RegisterDsiwrpcrFieldIdfMask  = 0x7800
)

// GetIdf IDF
func (r *RegisterDsiwrpcrType) GetIdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwrpcrFieldIdfMask) >> RegisterDsiwrpcrFieldIdfShift)
}

// SetIdf IDF
func (r *RegisterDsiwrpcrType) SetIdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwrpcrFieldIdfMask)|(uint32(value)<<RegisterDsiwrpcrFieldIdfShift))
}

const (
	RegisterDsiwrpcrFieldOdfShift = 16
	RegisterDsiwrpcrFieldOdfMask  = 0x30000
)

// GetOdf ODF
func (r *RegisterDsiwrpcrType) GetOdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwrpcrFieldOdfMask) >> RegisterDsiwrpcrFieldOdfShift)
}

// SetOdf ODF
func (r *RegisterDsiwrpcrType) SetOdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwrpcrFieldOdfMask)|(uint32(value)<<RegisterDsiwrpcrFieldOdfShift))
}

const (
	RegisterDsiwrpcrFieldRegenShift = 24
	RegisterDsiwrpcrFieldRegenMask  = 0x1000000
)

// GetRegen REGEN
func (r *RegisterDsiwrpcrType) GetRegen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwrpcrFieldRegenMask) != 0
}

// SetRegen REGEN
func (r *RegisterDsiwrpcrType) SetRegen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwrpcrFieldRegenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwrpcrFieldRegenMask)
	}
}
