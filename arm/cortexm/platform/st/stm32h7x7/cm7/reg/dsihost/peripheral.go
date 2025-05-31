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
	Dsivr      registerDsivrType
	Dsicr      registerDsicrType
	Dsiccr     registerDsiccrType
	Dsilvcidr  registerDsilvcidrType
	Dsilcolcr  registerDsilcolcrType
	Dsilpcr    registerDsilpcrType
	Dsilpmcr   registerDsilpmcrType
	_          [16]byte
	Dsipcr     registerDsipcrType
	Dsigvcidr  registerDsigvcidrType
	Dsimcr     registerDsimcrType
	Dsivmcr    registerDsivmcrType
	Dsivpcr    registerDsivpcrType
	Dsivccr    registerDsivccrType
	Dsivnpcr   registerDsivnpcrType
	Dsivhsacr  registerDsivhsacrType
	Dsivhbpcr  registerDsivhbpcrType
	Dsivlcr    registerDsivlcrType
	Dsivvsacr  registerDsivvsacrType
	Dsivvbpcr  registerDsivvbpcrType
	Dsivvfpcr  registerDsivvfpcrType
	Dsivvacr   registerDsivvacrType
	Dsilccr    registerDsilccrType
	Dsicmcr    registerDsicmcrType
	Dsighcr    registerDsighcrType
	Dsigpdr    registerDsigpdrType
	Dsigpsr    registerDsigpsrType
	Dsitccr0   registerDsitccr0Type
	Dsitccr1   registerDsitccr1Type
	Dsitccr2   registerDsitccr2Type
	Dsitccr3   registerDsitccr3Type
	Dsitccr4   registerDsitccr4Type
	Dsitccr5   registerDsitccr5Type
	_          [4]byte
	Dsiclcr    registerDsiclcrType
	Dsicltcr   registerDsicltcrType
	Dsidltcr   registerDsidltcrType
	Dsipctlr   registerDsipctlrType
	Dsipconfr  registerDsipconfrType
	Dsipucr    registerDsipucrType
	Dsipttcr   registerDsipttcrType
	Dsipsr     registerDsipsrType
	_          [8]byte
	Dsiisr0    registerDsiisr0Type
	Dsiisr1    registerDsiisr1Type
	Dsiier0    registerDsiier0Type
	Dsiier1    registerDsiier1Type
	_          [12]byte
	Dsifir0    registerDsifir0Type
	Dsifir1    registerDsifir1Type
	_          [32]byte
	Dsivscr    registerDsivscrType
	_          [8]byte
	Dsilcvcidr registerDsilcvcidrType
	Dsilcccr   registerDsilcccrType
	_          [4]byte
	Dsilpmccr  registerDsilpmccrType
	_          [28]byte
	Dsivmccr   registerDsivmccrType
	Dsivpccr   registerDsivpccrType
	Dsivcccr   registerDsivcccrType
	Dsivnpccr  registerDsivnpccrType
	Dsivhsaccr registerDsivhsaccrType
	Dsivhbpccr registerDsivhbpccrType
	Dsivlccr   registerDsivlccrType
	Dsivvsaccr registerDsivvsaccrType
	Dsivvbpccr registerDsivvbpccrType
	Dsivvfpccr registerDsivvfpccrType
	Dsivvaccr  registerDsivvaccrType
	_          [668]byte
	Dsiwcfgr   registerDsiwcfgrType
	Dsiwcr     registerDsiwcrType
	Dsiwier    registerDsiwierType
	Dsiwisr    registerDsiwisrType
	Dsiwifcr   registerDsiwifcrType
	_          [4]byte
	Dsiwpcr0   registerDsiwpcr0Type
	Dsiwpcr1   registerDsiwpcr1Type
	Dsiwpcr2   registerDsiwpcr2Type
	Dsiwpcr3   registerDsiwpcr3Type
	Dsiwpcr4   registerDsiwpcr4Type
	_          [4]byte
	Dsiwrpcr   registerDsiwrpcrType
}

// registerDsivrType DSI Host version register
type registerDsivrType uint32

const (
	RegisterDsivrFieldVersionShift = 0
	RegisterDsivrFieldVersionMask  = 0xffffffff
)

// GetVersion VERSION
func (r *registerDsivrType) GetVersion() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDsivrFieldVersionMask) >> RegisterDsivrFieldVersionShift)
}

// SetVersion VERSION
func (r *registerDsivrType) SetVersion(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivrFieldVersionMask)|(uint32(value)<<RegisterDsivrFieldVersionShift))
}

// registerDsicrType DSI Host control register
type registerDsicrType uint32

const (
	RegisterDsicrFieldEnShift = 0
	RegisterDsicrFieldEnMask  = 0x1
)

// GetEn EN
func (r *registerDsicrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicrFieldEnMask) != 0
}

// SetEn EN
func (r *registerDsicrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicrFieldEnMask)
	}
}

// registerDsiccrType DSI Host clock control register
type registerDsiccrType uint32

const (
	RegisterDsiccrFieldTxeckdivShift = 0
	RegisterDsiccrFieldTxeckdivMask  = 0xff
)

// GetTxeckdiv TXECKDIV
func (r *registerDsiccrType) GetTxeckdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiccrFieldTxeckdivMask) >> RegisterDsiccrFieldTxeckdivShift)
}

// SetTxeckdiv TXECKDIV
func (r *registerDsiccrType) SetTxeckdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiccrFieldTxeckdivMask)|(uint32(value)<<RegisterDsiccrFieldTxeckdivShift))
}

const (
	RegisterDsiccrFieldTockdivShift = 8
	RegisterDsiccrFieldTockdivMask  = 0xff00
)

// GetTockdiv TOCKDIV
func (r *registerDsiccrType) GetTockdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiccrFieldTockdivMask) >> RegisterDsiccrFieldTockdivShift)
}

// SetTockdiv TOCKDIV
func (r *registerDsiccrType) SetTockdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiccrFieldTockdivMask)|(uint32(value)<<RegisterDsiccrFieldTockdivShift))
}

// registerDsilvcidrType DSI Host LTDC VCID register
type registerDsilvcidrType uint32

const (
	RegisterDsilvcidrFieldVcidShift = 0
	RegisterDsilvcidrFieldVcidMask  = 0x3
)

// GetVcid VCID
func (r *registerDsilvcidrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilvcidrFieldVcidMask) >> RegisterDsilvcidrFieldVcidShift)
}

// SetVcid VCID
func (r *registerDsilvcidrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilvcidrFieldVcidMask)|(uint32(value)<<RegisterDsilvcidrFieldVcidShift))
}

// registerDsilcolcrType DSI Host LTDC color coding register
type registerDsilcolcrType uint32

const (
	RegisterDsilcolcrFieldColcShift = 0
	RegisterDsilcolcrFieldColcMask  = 0xf
)

// GetColc COLC
func (r *registerDsilcolcrType) GetColc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilcolcrFieldColcMask) >> RegisterDsilcolcrFieldColcShift)
}

// SetColc COLC
func (r *registerDsilcolcrType) SetColc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilcolcrFieldColcMask)|(uint32(value)<<RegisterDsilcolcrFieldColcShift))
}

const (
	RegisterDsilcolcrFieldLpeShift = 8
	RegisterDsilcolcrFieldLpeMask  = 0x100
)

// GetLpe LPE
func (r *registerDsilcolcrType) GetLpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsilcolcrFieldLpeMask) != 0
}

// SetLpe LPE
func (r *registerDsilcolcrType) SetLpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsilcolcrFieldLpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsilcolcrFieldLpeMask)
	}
}

// registerDsilpcrType DSI Host LTDC polarity configuration register
type registerDsilpcrType uint32

const (
	RegisterDsilpcrFieldDepShift = 0
	RegisterDsilpcrFieldDepMask  = 0x1
)

// GetDep DEP
func (r *registerDsilpcrType) GetDep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsilpcrFieldDepMask) != 0
}

// SetDep DEP
func (r *registerDsilpcrType) SetDep(value bool) {
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
func (r *registerDsilpcrType) GetVsp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsilpcrFieldVspMask) != 0
}

// SetVsp VSP
func (r *registerDsilpcrType) SetVsp(value bool) {
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
func (r *registerDsilpcrType) GetHsp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsilpcrFieldHspMask) != 0
}

// SetHsp HSP
func (r *registerDsilpcrType) SetHsp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsilpcrFieldHspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsilpcrFieldHspMask)
	}
}

// registerDsilpmcrType DSI Host low-power mode configuration register
type registerDsilpmcrType uint32

const (
	RegisterDsilpmcrFieldVlpsizeShift = 0
	RegisterDsilpmcrFieldVlpsizeMask  = 0xff
)

// GetVlpsize VLPSIZE
func (r *registerDsilpmcrType) GetVlpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilpmcrFieldVlpsizeMask) >> RegisterDsilpmcrFieldVlpsizeShift)
}

// SetVlpsize VLPSIZE
func (r *registerDsilpmcrType) SetVlpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilpmcrFieldVlpsizeMask)|(uint32(value)<<RegisterDsilpmcrFieldVlpsizeShift))
}

const (
	RegisterDsilpmcrFieldLpsizeShift = 16
	RegisterDsilpmcrFieldLpsizeMask  = 0xff0000
)

// GetLpsize LPSIZE
func (r *registerDsilpmcrType) GetLpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilpmcrFieldLpsizeMask) >> RegisterDsilpmcrFieldLpsizeShift)
}

// SetLpsize LPSIZE
func (r *registerDsilpmcrType) SetLpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilpmcrFieldLpsizeMask)|(uint32(value)<<RegisterDsilpmcrFieldLpsizeShift))
}

// registerDsipcrType DSI Host protocol configuration register
type registerDsipcrType uint32

const (
	RegisterDsipcrFieldEttxeShift = 0
	RegisterDsipcrFieldEttxeMask  = 0x1
)

// GetEttxe ETTXE
func (r *registerDsipcrType) GetEttxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipcrFieldEttxeMask) != 0
}

// SetEttxe ETTXE
func (r *registerDsipcrType) SetEttxe(value bool) {
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
func (r *registerDsipcrType) GetEtrxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipcrFieldEtrxeMask) != 0
}

// SetEtrxe ETRXE
func (r *registerDsipcrType) SetEtrxe(value bool) {
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
func (r *registerDsipcrType) GetBtae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipcrFieldBtaeMask) != 0
}

// SetBtae BTAE
func (r *registerDsipcrType) SetBtae(value bool) {
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
func (r *registerDsipcrType) GetEccrxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipcrFieldEccrxeMask) != 0
}

// SetEccrxe ECCRXE
func (r *registerDsipcrType) SetEccrxe(value bool) {
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
func (r *registerDsipcrType) GetCrcrxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipcrFieldCrcrxeMask) != 0
}

// SetCrcrxe CRCRXE
func (r *registerDsipcrType) SetCrcrxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipcrFieldCrcrxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipcrFieldCrcrxeMask)
	}
}

// registerDsigvcidrType DSI Host generic VCID register
type registerDsigvcidrType uint32

const (
	RegisterDsigvcidrFieldVcidShift = 0
	RegisterDsigvcidrFieldVcidMask  = 0x3
)

// GetVcid VCID
func (r *registerDsigvcidrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsigvcidrFieldVcidMask) >> RegisterDsigvcidrFieldVcidShift)
}

// SetVcid VCID
func (r *registerDsigvcidrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsigvcidrFieldVcidMask)|(uint32(value)<<RegisterDsigvcidrFieldVcidShift))
}

// registerDsimcrType DSI Host mode configuration register
type registerDsimcrType uint32

const (
	RegisterDsimcrFieldCmdmShift = 0
	RegisterDsimcrFieldCmdmMask  = 0x1
)

// GetCmdm CMDM
func (r *registerDsimcrType) GetCmdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsimcrFieldCmdmMask) != 0
}

// SetCmdm CMDM
func (r *registerDsimcrType) SetCmdm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsimcrFieldCmdmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsimcrFieldCmdmMask)
	}
}

// registerDsivmcrType DSI Host video mode configuration register
type registerDsivmcrType uint32

const (
	RegisterDsivmcrFieldVmtShift = 0
	RegisterDsivmcrFieldVmtMask  = 0x3
)

// GetVmt VMT
func (r *registerDsivmcrType) GetVmt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldVmtMask) >> RegisterDsivmcrFieldVmtShift)
}

// SetVmt VMT
func (r *registerDsivmcrType) SetVmt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldVmtMask)|(uint32(value)<<RegisterDsivmcrFieldVmtShift))
}

const (
	RegisterDsivmcrFieldLpvsaeShift = 8
	RegisterDsivmcrFieldLpvsaeMask  = 0x100
)

// GetLpvsae LPVSAE
func (r *registerDsivmcrType) GetLpvsae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLpvsaeMask) != 0
}

// SetLpvsae LPVSAE
func (r *registerDsivmcrType) SetLpvsae(value bool) {
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
func (r *registerDsivmcrType) GetLpvbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLpvbpeMask) != 0
}

// SetLpvbpe LPVBPE
func (r *registerDsivmcrType) SetLpvbpe(value bool) {
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
func (r *registerDsivmcrType) GetLpvfpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLpvfpeMask) != 0
}

// SetLpvfpe LPVFPE
func (r *registerDsivmcrType) SetLpvfpe(value bool) {
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
func (r *registerDsivmcrType) GetLpvae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLpvaeMask) != 0
}

// SetLpvae LPVAE
func (r *registerDsivmcrType) SetLpvae(value bool) {
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
func (r *registerDsivmcrType) GetLphbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLphbpeMask) != 0
}

// SetLphbpe LPHBPE
func (r *registerDsivmcrType) SetLphbpe(value bool) {
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
func (r *registerDsivmcrType) GetLphfpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLphfpeMask) != 0
}

// SetLphfpe LPHFPE
func (r *registerDsivmcrType) SetLphfpe(value bool) {
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
func (r *registerDsivmcrType) GetFbtaae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldFbtaaeMask) != 0
}

// SetFbtaae FBTAAE
func (r *registerDsivmcrType) SetFbtaae(value bool) {
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
func (r *registerDsivmcrType) GetLpce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldLpceMask) != 0
}

// SetLpce LPCE
func (r *registerDsivmcrType) SetLpce(value bool) {
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
func (r *registerDsivmcrType) GetPge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldPgeMask) != 0
}

// SetPge PGE
func (r *registerDsivmcrType) SetPge(value bool) {
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
func (r *registerDsivmcrType) GetPgm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldPgmMask) != 0
}

// SetPgm PGM
func (r *registerDsivmcrType) SetPgm(value bool) {
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
func (r *registerDsivmcrType) GetPgo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmcrFieldPgoMask) != 0
}

// SetPgo PGO
func (r *registerDsivmcrType) SetPgo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmcrFieldPgoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmcrFieldPgoMask)
	}
}

// registerDsivpcrType DSI Host video packet configuration register
type registerDsivpcrType uint32

const (
	RegisterDsivpcrFieldVpsizeShift = 0
	RegisterDsivpcrFieldVpsizeMask  = 0x3fff
)

// GetVpsize VPSIZE
func (r *registerDsivpcrType) GetVpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivpcrFieldVpsizeMask) >> RegisterDsivpcrFieldVpsizeShift)
}

// SetVpsize VPSIZE
func (r *registerDsivpcrType) SetVpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivpcrFieldVpsizeMask)|(uint32(value)<<RegisterDsivpcrFieldVpsizeShift))
}

// registerDsivccrType DSI Host video chunks configuration register
type registerDsivccrType uint32

const (
	RegisterDsivccrFieldNumcShift = 0
	RegisterDsivccrFieldNumcMask  = 0x1fff
)

// GetNumc NUMC
func (r *registerDsivccrType) GetNumc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivccrFieldNumcMask) >> RegisterDsivccrFieldNumcShift)
}

// SetNumc NUMC
func (r *registerDsivccrType) SetNumc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivccrFieldNumcMask)|(uint32(value)<<RegisterDsivccrFieldNumcShift))
}

// registerDsivnpcrType DSI Host video null packet configuration register
type registerDsivnpcrType uint32

const (
	RegisterDsivnpcrFieldNpsizeShift = 0
	RegisterDsivnpcrFieldNpsizeMask  = 0x1fff
)

// GetNpsize NPSIZE
func (r *registerDsivnpcrType) GetNpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivnpcrFieldNpsizeMask) >> RegisterDsivnpcrFieldNpsizeShift)
}

// SetNpsize NPSIZE
func (r *registerDsivnpcrType) SetNpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivnpcrFieldNpsizeMask)|(uint32(value)<<RegisterDsivnpcrFieldNpsizeShift))
}

// registerDsivhsacrType DSI Host video HSA configuration register
type registerDsivhsacrType uint32

const (
	RegisterDsivhsacrFieldHsaShift = 0
	RegisterDsivhsacrFieldHsaMask  = 0xfff
)

// GetHsa HSA
func (r *registerDsivhsacrType) GetHsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivhsacrFieldHsaMask) >> RegisterDsivhsacrFieldHsaShift)
}

// SetHsa HSA
func (r *registerDsivhsacrType) SetHsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivhsacrFieldHsaMask)|(uint32(value)<<RegisterDsivhsacrFieldHsaShift))
}

// registerDsivhbpcrType DSI Host video HBP configuration register
type registerDsivhbpcrType uint32

const (
	RegisterDsivhbpcrFieldHbpShift = 0
	RegisterDsivhbpcrFieldHbpMask  = 0xfff
)

// GetHbp HBP
func (r *registerDsivhbpcrType) GetHbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivhbpcrFieldHbpMask) >> RegisterDsivhbpcrFieldHbpShift)
}

// SetHbp HBP
func (r *registerDsivhbpcrType) SetHbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivhbpcrFieldHbpMask)|(uint32(value)<<RegisterDsivhbpcrFieldHbpShift))
}

// registerDsivlcrType DSI Host video line configuration register
type registerDsivlcrType uint32

const (
	RegisterDsivlcrFieldHlineShift = 0
	RegisterDsivlcrFieldHlineMask  = 0x7fff
)

// GetHline HLINE
func (r *registerDsivlcrType) GetHline() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivlcrFieldHlineMask) >> RegisterDsivlcrFieldHlineShift)
}

// SetHline HLINE
func (r *registerDsivlcrType) SetHline(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivlcrFieldHlineMask)|(uint32(value)<<RegisterDsivlcrFieldHlineShift))
}

// registerDsivvsacrType DSI Host video VSA configuration register
type registerDsivvsacrType uint32

const (
	RegisterDsivvsacrFieldVsaShift = 0
	RegisterDsivvsacrFieldVsaMask  = 0x3ff
)

// GetVsa VSA
func (r *registerDsivvsacrType) GetVsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvsacrFieldVsaMask) >> RegisterDsivvsacrFieldVsaShift)
}

// SetVsa VSA
func (r *registerDsivvsacrType) SetVsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvsacrFieldVsaMask)|(uint32(value)<<RegisterDsivvsacrFieldVsaShift))
}

// registerDsivvbpcrType DSI Host video VBP configuration register
type registerDsivvbpcrType uint32

const (
	RegisterDsivvbpcrFieldVbpShift = 0
	RegisterDsivvbpcrFieldVbpMask  = 0x3ff
)

// GetVbp VBP
func (r *registerDsivvbpcrType) GetVbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvbpcrFieldVbpMask) >> RegisterDsivvbpcrFieldVbpShift)
}

// SetVbp VBP
func (r *registerDsivvbpcrType) SetVbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvbpcrFieldVbpMask)|(uint32(value)<<RegisterDsivvbpcrFieldVbpShift))
}

// registerDsivvfpcrType DSI Host video VFP configuration register
type registerDsivvfpcrType uint32

const (
	RegisterDsivvfpcrFieldVfpShift = 0
	RegisterDsivvfpcrFieldVfpMask  = 0x3ff
)

// GetVfp VFP
func (r *registerDsivvfpcrType) GetVfp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvfpcrFieldVfpMask) >> RegisterDsivvfpcrFieldVfpShift)
}

// SetVfp VFP
func (r *registerDsivvfpcrType) SetVfp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvfpcrFieldVfpMask)|(uint32(value)<<RegisterDsivvfpcrFieldVfpShift))
}

// registerDsivvacrType DSI Host video VA configuration register
type registerDsivvacrType uint32

const (
	RegisterDsivvacrFieldVaShift = 0
	RegisterDsivvacrFieldVaMask  = 0x3fff
)

// GetVa VA
func (r *registerDsivvacrType) GetVa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvacrFieldVaMask) >> RegisterDsivvacrFieldVaShift)
}

// SetVa VA
func (r *registerDsivvacrType) SetVa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvacrFieldVaMask)|(uint32(value)<<RegisterDsivvacrFieldVaShift))
}

// registerDsilccrType DSI Host LTDC command configuration register
type registerDsilccrType uint32

const (
	RegisterDsilccrFieldCmdsizeShift = 0
	RegisterDsilccrFieldCmdsizeMask  = 0xffff
)

// GetCmdsize CMDSIZE
func (r *registerDsilccrType) GetCmdsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsilccrFieldCmdsizeMask) >> RegisterDsilccrFieldCmdsizeShift)
}

// SetCmdsize CMDSIZE
func (r *registerDsilccrType) SetCmdsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilccrFieldCmdsizeMask)|(uint32(value)<<RegisterDsilccrFieldCmdsizeShift))
}

// registerDsicmcrType DSI Host command mode configuration register
type registerDsicmcrType uint32

const (
	RegisterDsicmcrFieldTeareShift = 0
	RegisterDsicmcrFieldTeareMask  = 0x1
)

// GetTeare TEARE
func (r *registerDsicmcrType) GetTeare() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldTeareMask) != 0
}

// SetTeare TEARE
func (r *registerDsicmcrType) SetTeare(value bool) {
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
func (r *registerDsicmcrType) GetAre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldAreMask) != 0
}

// SetAre ARE
func (r *registerDsicmcrType) SetAre(value bool) {
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
func (r *registerDsicmcrType) GetGsw0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsw0txMask) != 0
}

// SetGsw0tx GSW0TX
func (r *registerDsicmcrType) SetGsw0tx(value bool) {
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
func (r *registerDsicmcrType) GetGsw1tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsw1txMask) != 0
}

// SetGsw1tx GSW1TX
func (r *registerDsicmcrType) SetGsw1tx(value bool) {
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
func (r *registerDsicmcrType) GetGsw2tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsw2txMask) != 0
}

// SetGsw2tx GSW2TX
func (r *registerDsicmcrType) SetGsw2tx(value bool) {
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
func (r *registerDsicmcrType) GetGsr0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsr0txMask) != 0
}

// SetGsr0tx GSR0TX
func (r *registerDsicmcrType) SetGsr0tx(value bool) {
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
func (r *registerDsicmcrType) GetGsr1tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsr1txMask) != 0
}

// SetGsr1tx GSR1TX
func (r *registerDsicmcrType) SetGsr1tx(value bool) {
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
func (r *registerDsicmcrType) GetGsr2tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGsr2txMask) != 0
}

// SetGsr2tx GSR2TX
func (r *registerDsicmcrType) SetGsr2tx(value bool) {
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
func (r *registerDsicmcrType) GetGlwtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldGlwtxMask) != 0
}

// SetGlwtx GLWTX
func (r *registerDsicmcrType) SetGlwtx(value bool) {
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
func (r *registerDsicmcrType) GetDsw0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldDsw0txMask) != 0
}

// SetDsw0tx DSW0TX
func (r *registerDsicmcrType) SetDsw0tx(value bool) {
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
func (r *registerDsicmcrType) GetDsw1tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldDsw1txMask) != 0
}

// SetDsw1tx DSW1TX
func (r *registerDsicmcrType) SetDsw1tx(value bool) {
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
func (r *registerDsicmcrType) GetDsr0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldDsr0txMask) != 0
}

// SetDsr0tx DSR0TX
func (r *registerDsicmcrType) SetDsr0tx(value bool) {
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
func (r *registerDsicmcrType) GetDlwtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldDlwtxMask) != 0
}

// SetDlwtx DLWTX
func (r *registerDsicmcrType) SetDlwtx(value bool) {
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
func (r *registerDsicmcrType) GetMrdps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsicmcrFieldMrdpsMask) != 0
}

// SetMrdps MRDPS
func (r *registerDsicmcrType) SetMrdps(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsicmcrFieldMrdpsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsicmcrFieldMrdpsMask)
	}
}

// registerDsighcrType DSI Host generic header configuration register
type registerDsighcrType uint32

const (
	RegisterDsighcrFieldDtShift = 0
	RegisterDsighcrFieldDtMask  = 0x3f
)

// GetDt DT
func (r *registerDsighcrType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsighcrFieldDtMask) >> RegisterDsighcrFieldDtShift)
}

// SetDt DT
func (r *registerDsighcrType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsighcrFieldDtMask)|(uint32(value)<<RegisterDsighcrFieldDtShift))
}

const (
	RegisterDsighcrFieldVcidShift = 6
	RegisterDsighcrFieldVcidMask  = 0xc0
)

// GetVcid VCID
func (r *registerDsighcrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsighcrFieldVcidMask) >> RegisterDsighcrFieldVcidShift)
}

// SetVcid VCID
func (r *registerDsighcrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsighcrFieldVcidMask)|(uint32(value)<<RegisterDsighcrFieldVcidShift))
}

const (
	RegisterDsighcrFieldWclsbShift = 8
	RegisterDsighcrFieldWclsbMask  = 0xff00
)

// GetWclsb WCLSB
func (r *registerDsighcrType) GetWclsb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsighcrFieldWclsbMask) >> RegisterDsighcrFieldWclsbShift)
}

// SetWclsb WCLSB
func (r *registerDsighcrType) SetWclsb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsighcrFieldWclsbMask)|(uint32(value)<<RegisterDsighcrFieldWclsbShift))
}

const (
	RegisterDsighcrFieldWcmsbShift = 16
	RegisterDsighcrFieldWcmsbMask  = 0xff0000
)

// GetWcmsb WCMSB
func (r *registerDsighcrType) GetWcmsb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsighcrFieldWcmsbMask) >> RegisterDsighcrFieldWcmsbShift)
}

// SetWcmsb WCMSB
func (r *registerDsighcrType) SetWcmsb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsighcrFieldWcmsbMask)|(uint32(value)<<RegisterDsighcrFieldWcmsbShift))
}

// registerDsigpdrType DSI Host generic payload data register
type registerDsigpdrType uint32

const (
	RegisterDsigpdrFieldData1Shift = 0
	RegisterDsigpdrFieldData1Mask  = 0xff
)

// GetData1 DATA1
func (r *registerDsigpdrType) GetData1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsigpdrFieldData1Mask) >> RegisterDsigpdrFieldData1Shift)
}

// SetData1 DATA1
func (r *registerDsigpdrType) SetData1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsigpdrFieldData1Mask)|(uint32(value)<<RegisterDsigpdrFieldData1Shift))
}

const (
	RegisterDsigpdrFieldData2Shift = 8
	RegisterDsigpdrFieldData2Mask  = 0xff00
)

// GetData2 DATA2
func (r *registerDsigpdrType) GetData2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsigpdrFieldData2Mask) >> RegisterDsigpdrFieldData2Shift)
}

// SetData2 DATA2
func (r *registerDsigpdrType) SetData2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsigpdrFieldData2Mask)|(uint32(value)<<RegisterDsigpdrFieldData2Shift))
}

const (
	RegisterDsigpdrFieldData3Shift = 16
	RegisterDsigpdrFieldData3Mask  = 0xff0000
)

// GetData3 DATA3
func (r *registerDsigpdrType) GetData3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsigpdrFieldData3Mask) >> RegisterDsigpdrFieldData3Shift)
}

// SetData3 DATA3
func (r *registerDsigpdrType) SetData3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsigpdrFieldData3Mask)|(uint32(value)<<RegisterDsigpdrFieldData3Shift))
}

const (
	RegisterDsigpdrFieldData4Shift = 24
	RegisterDsigpdrFieldData4Mask  = 0xff000000
)

// GetData4 DATA4
func (r *registerDsigpdrType) GetData4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsigpdrFieldData4Mask) >> RegisterDsigpdrFieldData4Shift)
}

// SetData4 DATA4
func (r *registerDsigpdrType) SetData4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsigpdrFieldData4Mask)|(uint32(value)<<RegisterDsigpdrFieldData4Shift))
}

// registerDsigpsrType DSI Host generic packet status register
type registerDsigpsrType uint32

const (
	RegisterDsigpsrFieldCmdfeShift = 0
	RegisterDsigpsrFieldCmdfeMask  = 0x1
)

// GetCmdfe CMDFE
func (r *registerDsigpsrType) GetCmdfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldCmdfeMask) != 0
}

// SetCmdfe CMDFE
func (r *registerDsigpsrType) SetCmdfe(value bool) {
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
func (r *registerDsigpsrType) GetCmdff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldCmdffMask) != 0
}

// SetCmdff CMDFF
func (r *registerDsigpsrType) SetCmdff(value bool) {
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
func (r *registerDsigpsrType) GetPwrfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldPwrfeMask) != 0
}

// SetPwrfe PWRFE
func (r *registerDsigpsrType) SetPwrfe(value bool) {
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
func (r *registerDsigpsrType) GetPwrff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldPwrffMask) != 0
}

// SetPwrff PWRFF
func (r *registerDsigpsrType) SetPwrff(value bool) {
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
func (r *registerDsigpsrType) GetPrdfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldPrdfeMask) != 0
}

// SetPrdfe PRDFE
func (r *registerDsigpsrType) SetPrdfe(value bool) {
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
func (r *registerDsigpsrType) GetPrdff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldPrdffMask) != 0
}

// SetPrdff PRDFF
func (r *registerDsigpsrType) SetPrdff(value bool) {
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
func (r *registerDsigpsrType) GetRcb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsigpsrFieldRcbMask) != 0
}

// SetRcb RCB
func (r *registerDsigpsrType) SetRcb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsigpsrFieldRcbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsigpsrFieldRcbMask)
	}
}

// registerDsitccr0Type DSI Host timeout counter configuration register 0
type registerDsitccr0Type uint32

const (
	RegisterDsitccr0FieldLprxtocntShift = 0
	RegisterDsitccr0FieldLprxtocntMask  = 0xffff
)

// GetLprxtocnt LPRX_TOCNT
func (r *registerDsitccr0Type) GetLprxtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr0FieldLprxtocntMask) >> RegisterDsitccr0FieldLprxtocntShift)
}

// SetLprxtocnt LPRX_TOCNT
func (r *registerDsitccr0Type) SetLprxtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr0FieldLprxtocntMask)|(uint32(value)<<RegisterDsitccr0FieldLprxtocntShift))
}

const (
	RegisterDsitccr0FieldHstxtocntShift = 16
	RegisterDsitccr0FieldHstxtocntMask  = 0xffff0000
)

// GetHstxtocnt HSTX_TOCNT
func (r *registerDsitccr0Type) GetHstxtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr0FieldHstxtocntMask) >> RegisterDsitccr0FieldHstxtocntShift)
}

// SetHstxtocnt HSTX_TOCNT
func (r *registerDsitccr0Type) SetHstxtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr0FieldHstxtocntMask)|(uint32(value)<<RegisterDsitccr0FieldHstxtocntShift))
}

// registerDsitccr1Type DSI Host timeout counter configuration register 1
type registerDsitccr1Type uint32

const (
	RegisterDsitccr1FieldHsrdtocntShift = 0
	RegisterDsitccr1FieldHsrdtocntMask  = 0xffff
)

// GetHsrdtocnt HSRD_TOCNT
func (r *registerDsitccr1Type) GetHsrdtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr1FieldHsrdtocntMask) >> RegisterDsitccr1FieldHsrdtocntShift)
}

// SetHsrdtocnt HSRD_TOCNT
func (r *registerDsitccr1Type) SetHsrdtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr1FieldHsrdtocntMask)|(uint32(value)<<RegisterDsitccr1FieldHsrdtocntShift))
}

// registerDsitccr2Type DSI Host timeout counter configuration register 2
type registerDsitccr2Type uint32

const (
	RegisterDsitccr2FieldLprdtocntShift = 0
	RegisterDsitccr2FieldLprdtocntMask  = 0xffff
)

// GetLprdtocnt LPRD_TOCNT
func (r *registerDsitccr2Type) GetLprdtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr2FieldLprdtocntMask) >> RegisterDsitccr2FieldLprdtocntShift)
}

// SetLprdtocnt LPRD_TOCNT
func (r *registerDsitccr2Type) SetLprdtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr2FieldLprdtocntMask)|(uint32(value)<<RegisterDsitccr2FieldLprdtocntShift))
}

// registerDsitccr3Type DSI Host timeout counter configuration register 3
type registerDsitccr3Type uint32

const (
	RegisterDsitccr3FieldHswrtocntShift = 0
	RegisterDsitccr3FieldHswrtocntMask  = 0xffff
)

// GetHswrtocnt HSWR_TOCNT
func (r *registerDsitccr3Type) GetHswrtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr3FieldHswrtocntMask) >> RegisterDsitccr3FieldHswrtocntShift)
}

// SetHswrtocnt HSWR_TOCNT
func (r *registerDsitccr3Type) SetHswrtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr3FieldHswrtocntMask)|(uint32(value)<<RegisterDsitccr3FieldHswrtocntShift))
}

const (
	RegisterDsitccr3FieldPmShift = 24
	RegisterDsitccr3FieldPmMask  = 0x1000000
)

// GetPm PM
func (r *registerDsitccr3Type) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr3FieldPmMask) != 0
}

// SetPm PM
func (r *registerDsitccr3Type) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsitccr3FieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr3FieldPmMask)
	}
}

// registerDsitccr4Type DSI Host timeout counter configuration register 4
type registerDsitccr4Type uint32

const (
	RegisterDsitccr4FieldLpwrtocntShift = 0
	RegisterDsitccr4FieldLpwrtocntMask  = 0xffff
)

// GetLpwrtocnt LPWR_TOCNT
func (r *registerDsitccr4Type) GetLpwrtocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr4FieldLpwrtocntMask) >> RegisterDsitccr4FieldLpwrtocntShift)
}

// SetLpwrtocnt LPWR_TOCNT
func (r *registerDsitccr4Type) SetLpwrtocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr4FieldLpwrtocntMask)|(uint32(value)<<RegisterDsitccr4FieldLpwrtocntShift))
}

// registerDsitccr5Type DSI Host timeout counter configuration register 5
type registerDsitccr5Type uint32

const (
	RegisterDsitccr5FieldBtatocntShift = 0
	RegisterDsitccr5FieldBtatocntMask  = 0xffff
)

// GetBtatocnt BTA_TOCNT
func (r *registerDsitccr5Type) GetBtatocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsitccr5FieldBtatocntMask) >> RegisterDsitccr5FieldBtatocntShift)
}

// SetBtatocnt BTA_TOCNT
func (r *registerDsitccr5Type) SetBtatocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsitccr5FieldBtatocntMask)|(uint32(value)<<RegisterDsitccr5FieldBtatocntShift))
}

// registerDsiclcrType DSI Host clock lane configuration register
type registerDsiclcrType uint32

const (
	RegisterDsiclcrFieldDpccShift = 0
	RegisterDsiclcrFieldDpccMask  = 0x1
)

// GetDpcc DPCC
func (r *registerDsiclcrType) GetDpcc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiclcrFieldDpccMask) != 0
}

// SetDpcc DPCC
func (r *registerDsiclcrType) SetDpcc(value bool) {
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
func (r *registerDsiclcrType) GetAcr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiclcrFieldAcrMask) != 0
}

// SetAcr ACR
func (r *registerDsiclcrType) SetAcr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiclcrFieldAcrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiclcrFieldAcrMask)
	}
}

// registerDsicltcrType DSI Host clock lane timer configuration register
type registerDsicltcrType uint32

const (
	RegisterDsicltcrFieldLp2hstimeShift = 0
	RegisterDsicltcrFieldLp2hstimeMask  = 0x3ff
)

// GetLp2hstime LP2HS_TIME
func (r *registerDsicltcrType) GetLp2hstime() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsicltcrFieldLp2hstimeMask) >> RegisterDsicltcrFieldLp2hstimeShift)
}

// SetLp2hstime LP2HS_TIME
func (r *registerDsicltcrType) SetLp2hstime(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsicltcrFieldLp2hstimeMask)|(uint32(value)<<RegisterDsicltcrFieldLp2hstimeShift))
}

const (
	RegisterDsicltcrFieldHs2lptimeShift = 16
	RegisterDsicltcrFieldHs2lptimeMask  = 0x3ff0000
)

// GetHs2lptime HS2LP_TIME
func (r *registerDsicltcrType) GetHs2lptime() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsicltcrFieldHs2lptimeMask) >> RegisterDsicltcrFieldHs2lptimeShift)
}

// SetHs2lptime HS2LP_TIME
func (r *registerDsicltcrType) SetHs2lptime(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsicltcrFieldHs2lptimeMask)|(uint32(value)<<RegisterDsicltcrFieldHs2lptimeShift))
}

// registerDsidltcrType DSI Host data lane timer configuration register
type registerDsidltcrType uint32

const (
	RegisterDsidltcrFieldMrdtimeShift = 0
	RegisterDsidltcrFieldMrdtimeMask  = 0x7fff
)

// GetMrdtime Maximum read time
func (r *registerDsidltcrType) GetMrdtime() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsidltcrFieldMrdtimeMask) >> RegisterDsidltcrFieldMrdtimeShift)
}

// SetMrdtime Maximum read time
func (r *registerDsidltcrType) SetMrdtime(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsidltcrFieldMrdtimeMask)|(uint32(value)<<RegisterDsidltcrFieldMrdtimeShift))
}

const (
	RegisterDsidltcrFieldLp2hstimeShift = 16
	RegisterDsidltcrFieldLp2hstimeMask  = 0xff0000
)

// GetLp2hstime LP2HS_TIME
func (r *registerDsidltcrType) GetLp2hstime() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsidltcrFieldLp2hstimeMask) >> RegisterDsidltcrFieldLp2hstimeShift)
}

// SetLp2hstime LP2HS_TIME
func (r *registerDsidltcrType) SetLp2hstime(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsidltcrFieldLp2hstimeMask)|(uint32(value)<<RegisterDsidltcrFieldLp2hstimeShift))
}

const (
	RegisterDsidltcrFieldHs2lptimeShift = 24
	RegisterDsidltcrFieldHs2lptimeMask  = 0xff000000
)

// GetHs2lptime HS2LP_TIME
func (r *registerDsidltcrType) GetHs2lptime() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsidltcrFieldHs2lptimeMask) >> RegisterDsidltcrFieldHs2lptimeShift)
}

// SetHs2lptime HS2LP_TIME
func (r *registerDsidltcrType) SetHs2lptime(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsidltcrFieldHs2lptimeMask)|(uint32(value)<<RegisterDsidltcrFieldHs2lptimeShift))
}

// registerDsipctlrType DSI Host PHY control register
type registerDsipctlrType uint32

const (
	RegisterDsipctlrFieldDenShift = 1
	RegisterDsipctlrFieldDenMask  = 0x2
)

// GetDen DEN
func (r *registerDsipctlrType) GetDen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipctlrFieldDenMask) != 0
}

// SetDen DEN
func (r *registerDsipctlrType) SetDen(value bool) {
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
func (r *registerDsipctlrType) GetCke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipctlrFieldCkeMask) != 0
}

// SetCke CKE
func (r *registerDsipctlrType) SetCke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipctlrFieldCkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipctlrFieldCkeMask)
	}
}

// registerDsipconfrType DSI Host PHY configuration register
type registerDsipconfrType uint32

const (
	RegisterDsipconfrFieldNlShift = 0
	RegisterDsipconfrFieldNlMask  = 0x3
)

// GetNl NL
func (r *registerDsipconfrType) GetNl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsipconfrFieldNlMask) >> RegisterDsipconfrFieldNlShift)
}

// SetNl NL
func (r *registerDsipconfrType) SetNl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsipconfrFieldNlMask)|(uint32(value)<<RegisterDsipconfrFieldNlShift))
}

const (
	RegisterDsipconfrFieldSwtimeShift = 8
	RegisterDsipconfrFieldSwtimeMask  = 0xff00
)

// GetSwtime SW_TIME
func (r *registerDsipconfrType) GetSwtime() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsipconfrFieldSwtimeMask) >> RegisterDsipconfrFieldSwtimeShift)
}

// SetSwtime SW_TIME
func (r *registerDsipconfrType) SetSwtime(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsipconfrFieldSwtimeMask)|(uint32(value)<<RegisterDsipconfrFieldSwtimeShift))
}

// registerDsipucrType DSI Host PHY ULPS control register
type registerDsipucrType uint32

const (
	RegisterDsipucrFieldUrclShift = 0
	RegisterDsipucrFieldUrclMask  = 0x1
)

// GetUrcl URCL
func (r *registerDsipucrType) GetUrcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipucrFieldUrclMask) != 0
}

// SetUrcl URCL
func (r *registerDsipucrType) SetUrcl(value bool) {
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
func (r *registerDsipucrType) GetUecl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipucrFieldUeclMask) != 0
}

// SetUecl UECL
func (r *registerDsipucrType) SetUecl(value bool) {
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
func (r *registerDsipucrType) GetUrdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipucrFieldUrdlMask) != 0
}

// SetUrdl URDL
func (r *registerDsipucrType) SetUrdl(value bool) {
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
func (r *registerDsipucrType) GetUedl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipucrFieldUedlMask) != 0
}

// SetUedl UEDL
func (r *registerDsipucrType) SetUedl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipucrFieldUedlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipucrFieldUedlMask)
	}
}

// registerDsipttcrType DSI Host PHY TX triggers configuration register
type registerDsipttcrType uint32

const (
	RegisterDsipttcrFieldTxtrigShift = 0
	RegisterDsipttcrFieldTxtrigMask  = 0xf
)

// GetTxtrig TX_TRIG
func (r *registerDsipttcrType) GetTxtrig() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsipttcrFieldTxtrigMask) >> RegisterDsipttcrFieldTxtrigShift)
}

// SetTxtrig TX_TRIG
func (r *registerDsipttcrType) SetTxtrig(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsipttcrFieldTxtrigMask)|(uint32(value)<<RegisterDsipttcrFieldTxtrigShift))
}

// registerDsipsrType DSI Host PHY status register
type registerDsipsrType uint32

const (
	RegisterDsipsrFieldPdShift = 1
	RegisterDsipsrFieldPdMask  = 0x2
)

// GetPd PD
func (r *registerDsipsrType) GetPd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldPdMask) != 0
}

// SetPd PD
func (r *registerDsipsrType) SetPd(value bool) {
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
func (r *registerDsipsrType) GetPssc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldPsscMask) != 0
}

// SetPssc PSSC
func (r *registerDsipsrType) SetPssc(value bool) {
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
func (r *registerDsipsrType) GetUanc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldUancMask) != 0
}

// SetUanc UANC
func (r *registerDsipsrType) SetUanc(value bool) {
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
func (r *registerDsipsrType) GetPss0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldPss0Mask) != 0
}

// SetPss0 PSS0
func (r *registerDsipsrType) SetPss0(value bool) {
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
func (r *registerDsipsrType) GetUan0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldUan0Mask) != 0
}

// SetUan0 UAN0
func (r *registerDsipsrType) SetUan0(value bool) {
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
func (r *registerDsipsrType) GetRue0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldRue0Mask) != 0
}

// SetRue0 RUE0
func (r *registerDsipsrType) SetRue0(value bool) {
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
func (r *registerDsipsrType) GetPss1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldPss1Mask) != 0
}

// SetPss1 PSS1
func (r *registerDsipsrType) SetPss1(value bool) {
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
func (r *registerDsipsrType) GetUan1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsipsrFieldUan1Mask) != 0
}

// SetUan1 UAN1
func (r *registerDsipsrType) SetUan1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsipsrFieldUan1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsipsrFieldUan1Mask)
	}
}

// registerDsiisr0Type DSI Host interrupt and status register 0
type registerDsiisr0Type uint32

const (
	RegisterDsiisr0FieldAe0Shift = 0
	RegisterDsiisr0FieldAe0Mask  = 0x1
)

// GetAe0 AE0
func (r *registerDsiisr0Type) GetAe0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe0Mask) != 0
}

// SetAe0 AE0
func (r *registerDsiisr0Type) SetAe0(value bool) {
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
func (r *registerDsiisr0Type) GetAe1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe1Mask) != 0
}

// SetAe1 AE1
func (r *registerDsiisr0Type) SetAe1(value bool) {
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
func (r *registerDsiisr0Type) GetAe2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe2Mask) != 0
}

// SetAe2 AE2
func (r *registerDsiisr0Type) SetAe2(value bool) {
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
func (r *registerDsiisr0Type) GetAe3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe3Mask) != 0
}

// SetAe3 AE3
func (r *registerDsiisr0Type) SetAe3(value bool) {
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
func (r *registerDsiisr0Type) GetAe4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe4Mask) != 0
}

// SetAe4 AE4
func (r *registerDsiisr0Type) SetAe4(value bool) {
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
func (r *registerDsiisr0Type) GetAe5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe5Mask) != 0
}

// SetAe5 AE5
func (r *registerDsiisr0Type) SetAe5(value bool) {
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
func (r *registerDsiisr0Type) GetAe6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe6Mask) != 0
}

// SetAe6 AE6
func (r *registerDsiisr0Type) SetAe6(value bool) {
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
func (r *registerDsiisr0Type) GetAe7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe7Mask) != 0
}

// SetAe7 AE7
func (r *registerDsiisr0Type) SetAe7(value bool) {
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
func (r *registerDsiisr0Type) GetAe8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe8Mask) != 0
}

// SetAe8 AE8
func (r *registerDsiisr0Type) SetAe8(value bool) {
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
func (r *registerDsiisr0Type) GetAe9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe9Mask) != 0
}

// SetAe9 AE9
func (r *registerDsiisr0Type) SetAe9(value bool) {
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
func (r *registerDsiisr0Type) GetAe10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe10Mask) != 0
}

// SetAe10 AE10
func (r *registerDsiisr0Type) SetAe10(value bool) {
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
func (r *registerDsiisr0Type) GetAe11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe11Mask) != 0
}

// SetAe11 AE11
func (r *registerDsiisr0Type) SetAe11(value bool) {
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
func (r *registerDsiisr0Type) GetAe12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe12Mask) != 0
}

// SetAe12 AE12
func (r *registerDsiisr0Type) SetAe12(value bool) {
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
func (r *registerDsiisr0Type) GetAe13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe13Mask) != 0
}

// SetAe13 AE13
func (r *registerDsiisr0Type) SetAe13(value bool) {
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
func (r *registerDsiisr0Type) GetAe14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe14Mask) != 0
}

// SetAe14 AE14
func (r *registerDsiisr0Type) SetAe14(value bool) {
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
func (r *registerDsiisr0Type) GetAe15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldAe15Mask) != 0
}

// SetAe15 AE15
func (r *registerDsiisr0Type) SetAe15(value bool) {
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
func (r *registerDsiisr0Type) GetPe0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldPe0Mask) != 0
}

// SetPe0 PE0
func (r *registerDsiisr0Type) SetPe0(value bool) {
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
func (r *registerDsiisr0Type) GetPe1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldPe1Mask) != 0
}

// SetPe1 PE1
func (r *registerDsiisr0Type) SetPe1(value bool) {
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
func (r *registerDsiisr0Type) GetPe2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldPe2Mask) != 0
}

// SetPe2 PE2
func (r *registerDsiisr0Type) SetPe2(value bool) {
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
func (r *registerDsiisr0Type) GetPe3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldPe3Mask) != 0
}

// SetPe3 PE3
func (r *registerDsiisr0Type) SetPe3(value bool) {
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
func (r *registerDsiisr0Type) GetPe4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr0FieldPe4Mask) != 0
}

// SetPe4 PE4
func (r *registerDsiisr0Type) SetPe4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr0FieldPe4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr0FieldPe4Mask)
	}
}

// registerDsiisr1Type DSI Host interrupt and status register 1
type registerDsiisr1Type uint32

const (
	RegisterDsiisr1FieldTohstxShift = 0
	RegisterDsiisr1FieldTohstxMask  = 0x1
)

// GetTohstx TOHSTX
func (r *registerDsiisr1Type) GetTohstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldTohstxMask) != 0
}

// SetTohstx TOHSTX
func (r *registerDsiisr1Type) SetTohstx(value bool) {
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
func (r *registerDsiisr1Type) GetTolprx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldTolprxMask) != 0
}

// SetTolprx TOLPRX
func (r *registerDsiisr1Type) SetTolprx(value bool) {
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
func (r *registerDsiisr1Type) GetEccse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldEccseMask) != 0
}

// SetEccse ECCSE
func (r *registerDsiisr1Type) SetEccse(value bool) {
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
func (r *registerDsiisr1Type) GetEccme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldEccmeMask) != 0
}

// SetEccme ECCME
func (r *registerDsiisr1Type) SetEccme(value bool) {
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
func (r *registerDsiisr1Type) GetCrce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldCrceMask) != 0
}

// SetCrce CRCE
func (r *registerDsiisr1Type) SetCrce(value bool) {
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
func (r *registerDsiisr1Type) GetPse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldPseMask) != 0
}

// SetPse PSE
func (r *registerDsiisr1Type) SetPse(value bool) {
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
func (r *registerDsiisr1Type) GetEotpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldEotpeMask) != 0
}

// SetEotpe EOTPE
func (r *registerDsiisr1Type) SetEotpe(value bool) {
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
func (r *registerDsiisr1Type) GetLpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldLpwreMask) != 0
}

// SetLpwre LPWRE
func (r *registerDsiisr1Type) SetLpwre(value bool) {
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
func (r *registerDsiisr1Type) GetGcwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldGcwreMask) != 0
}

// SetGcwre GCWRE
func (r *registerDsiisr1Type) SetGcwre(value bool) {
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
func (r *registerDsiisr1Type) GetGpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldGpwreMask) != 0
}

// SetGpwre GPWRE
func (r *registerDsiisr1Type) SetGpwre(value bool) {
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
func (r *registerDsiisr1Type) GetGptxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldGptxeMask) != 0
}

// SetGptxe GPTXE
func (r *registerDsiisr1Type) SetGptxe(value bool) {
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
func (r *registerDsiisr1Type) GetGprde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldGprdeMask) != 0
}

// SetGprde GPRDE
func (r *registerDsiisr1Type) SetGprde(value bool) {
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
func (r *registerDsiisr1Type) GetGprxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiisr1FieldGprxeMask) != 0
}

// SetGprxe GPRXE
func (r *registerDsiisr1Type) SetGprxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiisr1FieldGprxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiisr1FieldGprxeMask)
	}
}

// registerDsiier0Type DSI Host interrupt enable register 0
type registerDsiier0Type uint32

const (
	RegisterDsiier0FieldAe0ieShift = 0
	RegisterDsiier0FieldAe0ieMask  = 0x1
)

// GetAe0ie AE0IE
func (r *registerDsiier0Type) GetAe0ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe0ieMask) != 0
}

// SetAe0ie AE0IE
func (r *registerDsiier0Type) SetAe0ie(value bool) {
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
func (r *registerDsiier0Type) GetAe1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe1ieMask) != 0
}

// SetAe1ie AE1IE
func (r *registerDsiier0Type) SetAe1ie(value bool) {
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
func (r *registerDsiier0Type) GetAe2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe2ieMask) != 0
}

// SetAe2ie AE2IE
func (r *registerDsiier0Type) SetAe2ie(value bool) {
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
func (r *registerDsiier0Type) GetAe3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe3ieMask) != 0
}

// SetAe3ie AE3IE
func (r *registerDsiier0Type) SetAe3ie(value bool) {
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
func (r *registerDsiier0Type) GetAe4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe4ieMask) != 0
}

// SetAe4ie AE4IE
func (r *registerDsiier0Type) SetAe4ie(value bool) {
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
func (r *registerDsiier0Type) GetAe5ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe5ieMask) != 0
}

// SetAe5ie AE5IE
func (r *registerDsiier0Type) SetAe5ie(value bool) {
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
func (r *registerDsiier0Type) GetAe6ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe6ieMask) != 0
}

// SetAe6ie AE6IE
func (r *registerDsiier0Type) SetAe6ie(value bool) {
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
func (r *registerDsiier0Type) GetAe7ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe7ieMask) != 0
}

// SetAe7ie AE7IE
func (r *registerDsiier0Type) SetAe7ie(value bool) {
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
func (r *registerDsiier0Type) GetAe8ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe8ieMask) != 0
}

// SetAe8ie AE8IE
func (r *registerDsiier0Type) SetAe8ie(value bool) {
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
func (r *registerDsiier0Type) GetAe9ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe9ieMask) != 0
}

// SetAe9ie AE9IE
func (r *registerDsiier0Type) SetAe9ie(value bool) {
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
func (r *registerDsiier0Type) GetAe10ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe10ieMask) != 0
}

// SetAe10ie AE10IE
func (r *registerDsiier0Type) SetAe10ie(value bool) {
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
func (r *registerDsiier0Type) GetAe11ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe11ieMask) != 0
}

// SetAe11ie AE11IE
func (r *registerDsiier0Type) SetAe11ie(value bool) {
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
func (r *registerDsiier0Type) GetAe12ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe12ieMask) != 0
}

// SetAe12ie AE12IE
func (r *registerDsiier0Type) SetAe12ie(value bool) {
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
func (r *registerDsiier0Type) GetAe13ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe13ieMask) != 0
}

// SetAe13ie AE13IE
func (r *registerDsiier0Type) SetAe13ie(value bool) {
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
func (r *registerDsiier0Type) GetAe14ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe14ieMask) != 0
}

// SetAe14ie AE14IE
func (r *registerDsiier0Type) SetAe14ie(value bool) {
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
func (r *registerDsiier0Type) GetAe15ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldAe15ieMask) != 0
}

// SetAe15ie AE15IE
func (r *registerDsiier0Type) SetAe15ie(value bool) {
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
func (r *registerDsiier0Type) GetPe0ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldPe0ieMask) != 0
}

// SetPe0ie PE0IE
func (r *registerDsiier0Type) SetPe0ie(value bool) {
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
func (r *registerDsiier0Type) GetPe1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldPe1ieMask) != 0
}

// SetPe1ie PE1IE
func (r *registerDsiier0Type) SetPe1ie(value bool) {
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
func (r *registerDsiier0Type) GetPe2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldPe2ieMask) != 0
}

// SetPe2ie PE2IE
func (r *registerDsiier0Type) SetPe2ie(value bool) {
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
func (r *registerDsiier0Type) GetPe3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldPe3ieMask) != 0
}

// SetPe3ie PE3IE
func (r *registerDsiier0Type) SetPe3ie(value bool) {
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
func (r *registerDsiier0Type) GetPe4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier0FieldPe4ieMask) != 0
}

// SetPe4ie PE4IE
func (r *registerDsiier0Type) SetPe4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier0FieldPe4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier0FieldPe4ieMask)
	}
}

// registerDsiier1Type DSI Host interrupt enable register 1
type registerDsiier1Type uint32

const (
	RegisterDsiier1FieldTohstxieShift = 0
	RegisterDsiier1FieldTohstxieMask  = 0x1
)

// GetTohstxie TOHSTXIE
func (r *registerDsiier1Type) GetTohstxie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldTohstxieMask) != 0
}

// SetTohstxie TOHSTXIE
func (r *registerDsiier1Type) SetTohstxie(value bool) {
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
func (r *registerDsiier1Type) GetTolprxie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldTolprxieMask) != 0
}

// SetTolprxie TOLPRXIE
func (r *registerDsiier1Type) SetTolprxie(value bool) {
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
func (r *registerDsiier1Type) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldEccseieMask) != 0
}

// SetEccseie ECCSEIE
func (r *registerDsiier1Type) SetEccseie(value bool) {
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
func (r *registerDsiier1Type) GetEccmeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldEccmeieMask) != 0
}

// SetEccmeie ECCMEIE
func (r *registerDsiier1Type) SetEccmeie(value bool) {
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
func (r *registerDsiier1Type) GetCrceie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldCrceieMask) != 0
}

// SetCrceie CRCEIE
func (r *registerDsiier1Type) SetCrceie(value bool) {
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
func (r *registerDsiier1Type) GetPseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldPseieMask) != 0
}

// SetPseie PSEIE
func (r *registerDsiier1Type) SetPseie(value bool) {
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
func (r *registerDsiier1Type) GetEotpeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldEotpeieMask) != 0
}

// SetEotpeie EOTPEIE
func (r *registerDsiier1Type) SetEotpeie(value bool) {
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
func (r *registerDsiier1Type) GetLpwreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldLpwreieMask) != 0
}

// SetLpwreie LPWREIE
func (r *registerDsiier1Type) SetLpwreie(value bool) {
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
func (r *registerDsiier1Type) GetGcwreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldGcwreieMask) != 0
}

// SetGcwreie GCWREIE
func (r *registerDsiier1Type) SetGcwreie(value bool) {
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
func (r *registerDsiier1Type) GetGpwreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldGpwreieMask) != 0
}

// SetGpwreie GPWREIE
func (r *registerDsiier1Type) SetGpwreie(value bool) {
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
func (r *registerDsiier1Type) GetGptxeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldGptxeieMask) != 0
}

// SetGptxeie GPTXEIE
func (r *registerDsiier1Type) SetGptxeie(value bool) {
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
func (r *registerDsiier1Type) GetGprdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldGprdeieMask) != 0
}

// SetGprdeie GPRDEIE
func (r *registerDsiier1Type) SetGprdeie(value bool) {
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
func (r *registerDsiier1Type) GetGprxeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiier1FieldGprxeieMask) != 0
}

// SetGprxeie GPRXEIE
func (r *registerDsiier1Type) SetGprxeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiier1FieldGprxeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiier1FieldGprxeieMask)
	}
}

// registerDsifir0Type DSI Host force interrupt register 0
type registerDsifir0Type uint32

const (
	RegisterDsifir0FieldFae0Shift = 0
	RegisterDsifir0FieldFae0Mask  = 0x1
)

// GetFae0 FAE0
func (r *registerDsifir0Type) GetFae0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae0Mask) != 0
}

// SetFae0 FAE0
func (r *registerDsifir0Type) SetFae0(value bool) {
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
func (r *registerDsifir0Type) GetFae1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae1Mask) != 0
}

// SetFae1 FAE1
func (r *registerDsifir0Type) SetFae1(value bool) {
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
func (r *registerDsifir0Type) GetFae2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae2Mask) != 0
}

// SetFae2 FAE2
func (r *registerDsifir0Type) SetFae2(value bool) {
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
func (r *registerDsifir0Type) GetFae3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae3Mask) != 0
}

// SetFae3 FAE3
func (r *registerDsifir0Type) SetFae3(value bool) {
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
func (r *registerDsifir0Type) GetFae4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae4Mask) != 0
}

// SetFae4 FAE4
func (r *registerDsifir0Type) SetFae4(value bool) {
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
func (r *registerDsifir0Type) GetFae5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae5Mask) != 0
}

// SetFae5 FAE5
func (r *registerDsifir0Type) SetFae5(value bool) {
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
func (r *registerDsifir0Type) GetFae6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae6Mask) != 0
}

// SetFae6 FAE6
func (r *registerDsifir0Type) SetFae6(value bool) {
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
func (r *registerDsifir0Type) GetFae7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae7Mask) != 0
}

// SetFae7 FAE7
func (r *registerDsifir0Type) SetFae7(value bool) {
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
func (r *registerDsifir0Type) GetFae8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae8Mask) != 0
}

// SetFae8 FAE8
func (r *registerDsifir0Type) SetFae8(value bool) {
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
func (r *registerDsifir0Type) GetFae9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae9Mask) != 0
}

// SetFae9 FAE9
func (r *registerDsifir0Type) SetFae9(value bool) {
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
func (r *registerDsifir0Type) GetFae10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae10Mask) != 0
}

// SetFae10 FAE10
func (r *registerDsifir0Type) SetFae10(value bool) {
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
func (r *registerDsifir0Type) GetFae11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae11Mask) != 0
}

// SetFae11 FAE11
func (r *registerDsifir0Type) SetFae11(value bool) {
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
func (r *registerDsifir0Type) GetFae12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae12Mask) != 0
}

// SetFae12 FAE12
func (r *registerDsifir0Type) SetFae12(value bool) {
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
func (r *registerDsifir0Type) GetFae13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae13Mask) != 0
}

// SetFae13 FAE13
func (r *registerDsifir0Type) SetFae13(value bool) {
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
func (r *registerDsifir0Type) GetFae14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae14Mask) != 0
}

// SetFae14 FAE14
func (r *registerDsifir0Type) SetFae14(value bool) {
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
func (r *registerDsifir0Type) GetFae15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFae15Mask) != 0
}

// SetFae15 FAE15
func (r *registerDsifir0Type) SetFae15(value bool) {
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
func (r *registerDsifir0Type) GetFpe0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFpe0Mask) != 0
}

// SetFpe0 FPE0
func (r *registerDsifir0Type) SetFpe0(value bool) {
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
func (r *registerDsifir0Type) GetFpe1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFpe1Mask) != 0
}

// SetFpe1 FPE1
func (r *registerDsifir0Type) SetFpe1(value bool) {
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
func (r *registerDsifir0Type) GetFpe2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFpe2Mask) != 0
}

// SetFpe2 FPE2
func (r *registerDsifir0Type) SetFpe2(value bool) {
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
func (r *registerDsifir0Type) GetFpe3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFpe3Mask) != 0
}

// SetFpe3 FPE3
func (r *registerDsifir0Type) SetFpe3(value bool) {
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
func (r *registerDsifir0Type) GetFpe4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir0FieldFpe4Mask) != 0
}

// SetFpe4 FPE4
func (r *registerDsifir0Type) SetFpe4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir0FieldFpe4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir0FieldFpe4Mask)
	}
}

// registerDsifir1Type DSI Host force interrupt register 1
type registerDsifir1Type uint32

const (
	RegisterDsifir1FieldFtohstxShift = 0
	RegisterDsifir1FieldFtohstxMask  = 0x1
)

// GetFtohstx FTOHSTX
func (r *registerDsifir1Type) GetFtohstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFtohstxMask) != 0
}

// SetFtohstx FTOHSTX
func (r *registerDsifir1Type) SetFtohstx(value bool) {
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
func (r *registerDsifir1Type) GetFtolprx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFtolprxMask) != 0
}

// SetFtolprx FTOLPRX
func (r *registerDsifir1Type) SetFtolprx(value bool) {
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
func (r *registerDsifir1Type) GetFeccse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFeccseMask) != 0
}

// SetFeccse FECCSE
func (r *registerDsifir1Type) SetFeccse(value bool) {
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
func (r *registerDsifir1Type) GetFeccme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFeccmeMask) != 0
}

// SetFeccme FECCME
func (r *registerDsifir1Type) SetFeccme(value bool) {
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
func (r *registerDsifir1Type) GetFcrce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFcrceMask) != 0
}

// SetFcrce FCRCE
func (r *registerDsifir1Type) SetFcrce(value bool) {
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
func (r *registerDsifir1Type) GetFpse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFpseMask) != 0
}

// SetFpse FPSE
func (r *registerDsifir1Type) SetFpse(value bool) {
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
func (r *registerDsifir1Type) GetFeotpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFeotpeMask) != 0
}

// SetFeotpe FEOTPE
func (r *registerDsifir1Type) SetFeotpe(value bool) {
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
func (r *registerDsifir1Type) GetFlpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFlpwreMask) != 0
}

// SetFlpwre FLPWRE
func (r *registerDsifir1Type) SetFlpwre(value bool) {
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
func (r *registerDsifir1Type) GetFgcwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFgcwreMask) != 0
}

// SetFgcwre FGCWRE
func (r *registerDsifir1Type) SetFgcwre(value bool) {
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
func (r *registerDsifir1Type) GetFgpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFgpwreMask) != 0
}

// SetFgpwre FGPWRE
func (r *registerDsifir1Type) SetFgpwre(value bool) {
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
func (r *registerDsifir1Type) GetFgptxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFgptxeMask) != 0
}

// SetFgptxe FGPTXE
func (r *registerDsifir1Type) SetFgptxe(value bool) {
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
func (r *registerDsifir1Type) GetFgprde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFgprdeMask) != 0
}

// SetFgprde FGPRDE
func (r *registerDsifir1Type) SetFgprde(value bool) {
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
func (r *registerDsifir1Type) GetFgprxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsifir1FieldFgprxeMask) != 0
}

// SetFgprxe FGPRXE
func (r *registerDsifir1Type) SetFgprxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsifir1FieldFgprxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsifir1FieldFgprxeMask)
	}
}

// registerDsivscrType DSI Host video shadow control register
type registerDsivscrType uint32

const (
	RegisterDsivscrFieldEnShift = 0
	RegisterDsivscrFieldEnMask  = 0x1
)

// GetEn EN
func (r *registerDsivscrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivscrFieldEnMask) != 0
}

// SetEn EN
func (r *registerDsivscrType) SetEn(value bool) {
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
func (r *registerDsivscrType) GetUr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivscrFieldUrMask) != 0
}

// SetUr UR
func (r *registerDsivscrType) SetUr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivscrFieldUrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivscrFieldUrMask)
	}
}

// registerDsilcvcidrType DSI Host LTDC current VCID register
type registerDsilcvcidrType uint32

const (
	RegisterDsilcvcidrFieldVcidShift = 0
	RegisterDsilcvcidrFieldVcidMask  = 0x3
)

// GetVcid VCID
func (r *registerDsilcvcidrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilcvcidrFieldVcidMask) >> RegisterDsilcvcidrFieldVcidShift)
}

// SetVcid VCID
func (r *registerDsilcvcidrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilcvcidrFieldVcidMask)|(uint32(value)<<RegisterDsilcvcidrFieldVcidShift))
}

// registerDsilcccrType DSI Host LTDC current color coding register
type registerDsilcccrType uint32

const (
	RegisterDsilcccrFieldColcShift = 0
	RegisterDsilcccrFieldColcMask  = 0xf
)

// GetColc COLC
func (r *registerDsilcccrType) GetColc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilcccrFieldColcMask) >> RegisterDsilcccrFieldColcShift)
}

// SetColc COLC
func (r *registerDsilcccrType) SetColc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilcccrFieldColcMask)|(uint32(value)<<RegisterDsilcccrFieldColcShift))
}

const (
	RegisterDsilcccrFieldLpeShift = 8
	RegisterDsilcccrFieldLpeMask  = 0x100
)

// GetLpe LPE
func (r *registerDsilcccrType) GetLpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsilcccrFieldLpeMask) != 0
}

// SetLpe LPE
func (r *registerDsilcccrType) SetLpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsilcccrFieldLpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsilcccrFieldLpeMask)
	}
}

// registerDsilpmccrType DSI Host low-power mode current configuration register
type registerDsilpmccrType uint32

const (
	RegisterDsilpmccrFieldVlpsizeShift = 0
	RegisterDsilpmccrFieldVlpsizeMask  = 0xff
)

// GetVlpsize VLPSIZE
func (r *registerDsilpmccrType) GetVlpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilpmccrFieldVlpsizeMask) >> RegisterDsilpmccrFieldVlpsizeShift)
}

// SetVlpsize VLPSIZE
func (r *registerDsilpmccrType) SetVlpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilpmccrFieldVlpsizeMask)|(uint32(value)<<RegisterDsilpmccrFieldVlpsizeShift))
}

const (
	RegisterDsilpmccrFieldLpsizeShift = 16
	RegisterDsilpmccrFieldLpsizeMask  = 0xff0000
)

// GetLpsize LPSIZE
func (r *registerDsilpmccrType) GetLpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsilpmccrFieldLpsizeMask) >> RegisterDsilpmccrFieldLpsizeShift)
}

// SetLpsize LPSIZE
func (r *registerDsilpmccrType) SetLpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsilpmccrFieldLpsizeMask)|(uint32(value)<<RegisterDsilpmccrFieldLpsizeShift))
}

// registerDsivmccrType DSI Host video mode current configuration register
type registerDsivmccrType uint32

const (
	RegisterDsivmccrFieldVmtShift = 0
	RegisterDsivmccrFieldVmtMask  = 0x3
)

// GetVmt VMT
func (r *registerDsivmccrType) GetVmt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldVmtMask) >> RegisterDsivmccrFieldVmtShift)
}

// SetVmt VMT
func (r *registerDsivmccrType) SetVmt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivmccrFieldVmtMask)|(uint32(value)<<RegisterDsivmccrFieldVmtShift))
}

const (
	RegisterDsivmccrFieldLpvsaeShift = 2
	RegisterDsivmccrFieldLpvsaeMask  = 0x4
)

// GetLpvsae LPVSAE
func (r *registerDsivmccrType) GetLpvsae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLpvsaeMask) != 0
}

// SetLpvsae LPVSAE
func (r *registerDsivmccrType) SetLpvsae(value bool) {
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
func (r *registerDsivmccrType) GetLpvbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLpvbpeMask) != 0
}

// SetLpvbpe LPVBPE
func (r *registerDsivmccrType) SetLpvbpe(value bool) {
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
func (r *registerDsivmccrType) GetLpvfpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLpvfpeMask) != 0
}

// SetLpvfpe LPVFPE
func (r *registerDsivmccrType) SetLpvfpe(value bool) {
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
func (r *registerDsivmccrType) GetLpvae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLpvaeMask) != 0
}

// SetLpvae LPVAE
func (r *registerDsivmccrType) SetLpvae(value bool) {
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
func (r *registerDsivmccrType) GetLphbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLphbpeMask) != 0
}

// SetLphbpe LPHBPE
func (r *registerDsivmccrType) SetLphbpe(value bool) {
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
func (r *registerDsivmccrType) GetLphfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLphfeMask) != 0
}

// SetLphfe LPHFE
func (r *registerDsivmccrType) SetLphfe(value bool) {
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
func (r *registerDsivmccrType) GetFbtaae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldFbtaaeMask) != 0
}

// SetFbtaae FBTAAE
func (r *registerDsivmccrType) SetFbtaae(value bool) {
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
func (r *registerDsivmccrType) GetLpce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsivmccrFieldLpceMask) != 0
}

// SetLpce LPCE
func (r *registerDsivmccrType) SetLpce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsivmccrFieldLpceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsivmccrFieldLpceMask)
	}
}

// registerDsivpccrType DSI Host video packet current configuration register
type registerDsivpccrType uint32

const (
	RegisterDsivpccrFieldVpsizeShift = 0
	RegisterDsivpccrFieldVpsizeMask  = 0x3fff
)

// GetVpsize VPSIZE
func (r *registerDsivpccrType) GetVpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivpccrFieldVpsizeMask) >> RegisterDsivpccrFieldVpsizeShift)
}

// SetVpsize VPSIZE
func (r *registerDsivpccrType) SetVpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivpccrFieldVpsizeMask)|(uint32(value)<<RegisterDsivpccrFieldVpsizeShift))
}

// registerDsivcccrType DSI Host video chunks current configuration register
type registerDsivcccrType uint32

const (
	RegisterDsivcccrFieldNumcShift = 0
	RegisterDsivcccrFieldNumcMask  = 0x1fff
)

// GetNumc NUMC
func (r *registerDsivcccrType) GetNumc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivcccrFieldNumcMask) >> RegisterDsivcccrFieldNumcShift)
}

// SetNumc NUMC
func (r *registerDsivcccrType) SetNumc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivcccrFieldNumcMask)|(uint32(value)<<RegisterDsivcccrFieldNumcShift))
}

// registerDsivnpccrType DSI Host video null packet current configuration register
type registerDsivnpccrType uint32

const (
	RegisterDsivnpccrFieldNpsizeShift = 0
	RegisterDsivnpccrFieldNpsizeMask  = 0x1fff
)

// GetNpsize NPSIZE
func (r *registerDsivnpccrType) GetNpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivnpccrFieldNpsizeMask) >> RegisterDsivnpccrFieldNpsizeShift)
}

// SetNpsize NPSIZE
func (r *registerDsivnpccrType) SetNpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivnpccrFieldNpsizeMask)|(uint32(value)<<RegisterDsivnpccrFieldNpsizeShift))
}

// registerDsivhsaccrType DSI Host video HSA current configuration register
type registerDsivhsaccrType uint32

const (
	RegisterDsivhsaccrFieldHsaShift = 0
	RegisterDsivhsaccrFieldHsaMask  = 0xfff
)

// GetHsa HSA
func (r *registerDsivhsaccrType) GetHsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivhsaccrFieldHsaMask) >> RegisterDsivhsaccrFieldHsaShift)
}

// SetHsa HSA
func (r *registerDsivhsaccrType) SetHsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivhsaccrFieldHsaMask)|(uint32(value)<<RegisterDsivhsaccrFieldHsaShift))
}

// registerDsivhbpccrType DSI Host video HBP current configuration register
type registerDsivhbpccrType uint32

const (
	RegisterDsivhbpccrFieldHbpShift = 0
	RegisterDsivhbpccrFieldHbpMask  = 0xfff
)

// GetHbp HBP
func (r *registerDsivhbpccrType) GetHbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivhbpccrFieldHbpMask) >> RegisterDsivhbpccrFieldHbpShift)
}

// SetHbp HBP
func (r *registerDsivhbpccrType) SetHbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivhbpccrFieldHbpMask)|(uint32(value)<<RegisterDsivhbpccrFieldHbpShift))
}

// registerDsivlccrType DSI Host video line current configuration register
type registerDsivlccrType uint32

const (
	RegisterDsivlccrFieldHlineShift = 0
	RegisterDsivlccrFieldHlineMask  = 0x7fff
)

// GetHline HLINE
func (r *registerDsivlccrType) GetHline() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivlccrFieldHlineMask) >> RegisterDsivlccrFieldHlineShift)
}

// SetHline HLINE
func (r *registerDsivlccrType) SetHline(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivlccrFieldHlineMask)|(uint32(value)<<RegisterDsivlccrFieldHlineShift))
}

// registerDsivvsaccrType DSI Host video VSA current configuration register
type registerDsivvsaccrType uint32

const (
	RegisterDsivvsaccrFieldVsaShift = 0
	RegisterDsivvsaccrFieldVsaMask  = 0x3ff
)

// GetVsa VSA
func (r *registerDsivvsaccrType) GetVsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvsaccrFieldVsaMask) >> RegisterDsivvsaccrFieldVsaShift)
}

// SetVsa VSA
func (r *registerDsivvsaccrType) SetVsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvsaccrFieldVsaMask)|(uint32(value)<<RegisterDsivvsaccrFieldVsaShift))
}

// registerDsivvbpccrType DSI Host video VBP current configuration register
type registerDsivvbpccrType uint32

const (
	RegisterDsivvbpccrFieldVbpShift = 0
	RegisterDsivvbpccrFieldVbpMask  = 0x3ff
)

// GetVbp VBP
func (r *registerDsivvbpccrType) GetVbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvbpccrFieldVbpMask) >> RegisterDsivvbpccrFieldVbpShift)
}

// SetVbp VBP
func (r *registerDsivvbpccrType) SetVbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvbpccrFieldVbpMask)|(uint32(value)<<RegisterDsivvbpccrFieldVbpShift))
}

// registerDsivvfpccrType DSI Host video VFP current configuration register
type registerDsivvfpccrType uint32

const (
	RegisterDsivvfpccrFieldVfpShift = 0
	RegisterDsivvfpccrFieldVfpMask  = 0x3ff
)

// GetVfp VFP
func (r *registerDsivvfpccrType) GetVfp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvfpccrFieldVfpMask) >> RegisterDsivvfpccrFieldVfpShift)
}

// SetVfp VFP
func (r *registerDsivvfpccrType) SetVfp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvfpccrFieldVfpMask)|(uint32(value)<<RegisterDsivvfpccrFieldVfpShift))
}

// registerDsivvaccrType DSI Host video VA current configuration register
type registerDsivvaccrType uint32

const (
	RegisterDsivvaccrFieldVaShift = 0
	RegisterDsivvaccrFieldVaMask  = 0x3fff
)

// GetVa VA
func (r *registerDsivvaccrType) GetVa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsivvaccrFieldVaMask) >> RegisterDsivvaccrFieldVaShift)
}

// SetVa VA
func (r *registerDsivvaccrType) SetVa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsivvaccrFieldVaMask)|(uint32(value)<<RegisterDsivvaccrFieldVaShift))
}

// registerDsiwcfgrType DSI wrapper configuration register
type registerDsiwcfgrType uint32

const (
	RegisterDsiwcfgrFieldDsimShift = 0
	RegisterDsiwcfgrFieldDsimMask  = 0x1
)

// GetDsim DSIM
func (r *registerDsiwcfgrType) GetDsim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldDsimMask) != 0
}

// SetDsim DSIM
func (r *registerDsiwcfgrType) SetDsim(value bool) {
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
func (r *registerDsiwcfgrType) GetColmux() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldColmuxMask) >> RegisterDsiwcfgrFieldColmuxShift)
}

// SetColmux COLMUX
func (r *registerDsiwcfgrType) SetColmux(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcfgrFieldColmuxMask)|(uint32(value)<<RegisterDsiwcfgrFieldColmuxShift))
}

const (
	RegisterDsiwcfgrFieldTesrcShift = 4
	RegisterDsiwcfgrFieldTesrcMask  = 0x10
)

// GetTesrc TESRC
func (r *registerDsiwcfgrType) GetTesrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldTesrcMask) != 0
}

// SetTesrc TESRC
func (r *registerDsiwcfgrType) SetTesrc(value bool) {
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
func (r *registerDsiwcfgrType) GetTepol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldTepolMask) != 0
}

// SetTepol TEPOL
func (r *registerDsiwcfgrType) SetTepol(value bool) {
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
func (r *registerDsiwcfgrType) GetAr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldArMask) != 0
}

// SetAr AR
func (r *registerDsiwcfgrType) SetAr(value bool) {
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
func (r *registerDsiwcfgrType) GetVspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcfgrFieldVspolMask) != 0
}

// SetVspol VSPOL
func (r *registerDsiwcfgrType) SetVspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwcfgrFieldVspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcfgrFieldVspolMask)
	}
}

// registerDsiwcrType DSI wrapper control register
type registerDsiwcrType uint32

const (
	RegisterDsiwcrFieldColmShift = 0
	RegisterDsiwcrFieldColmMask  = 0x1
)

// GetColm COLM
func (r *registerDsiwcrType) GetColm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcrFieldColmMask) != 0
}

// SetColm COLM
func (r *registerDsiwcrType) SetColm(value bool) {
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
func (r *registerDsiwcrType) GetShtdn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcrFieldShtdnMask) != 0
}

// SetShtdn SHTDN
func (r *registerDsiwcrType) SetShtdn(value bool) {
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
func (r *registerDsiwcrType) GetLtdcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcrFieldLtdcenMask) != 0
}

// SetLtdcen LTDCEN
func (r *registerDsiwcrType) SetLtdcen(value bool) {
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
func (r *registerDsiwcrType) GetDsien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwcrFieldDsienMask) != 0
}

// SetDsien DSIEN
func (r *registerDsiwcrType) SetDsien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwcrFieldDsienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwcrFieldDsienMask)
	}
}

// registerDsiwierType DSI wrapper interrupt enable register
type registerDsiwierType uint32

const (
	RegisterDsiwierFieldTeieShift = 0
	RegisterDsiwierFieldTeieMask  = 0x1
)

// GetTeie TEIE
func (r *registerDsiwierType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwierFieldTeieMask) != 0
}

// SetTeie TEIE
func (r *registerDsiwierType) SetTeie(value bool) {
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
func (r *registerDsiwierType) GetErie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwierFieldErieMask) != 0
}

// SetErie ERIE
func (r *registerDsiwierType) SetErie(value bool) {
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
func (r *registerDsiwierType) GetPlllie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwierFieldPlllieMask) != 0
}

// SetPlllie PLLLIE
func (r *registerDsiwierType) SetPlllie(value bool) {
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
func (r *registerDsiwierType) GetPlluie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwierFieldPlluieMask) != 0
}

// SetPlluie PLLUIE
func (r *registerDsiwierType) SetPlluie(value bool) {
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
func (r *registerDsiwierType) GetRrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwierFieldRrieMask) != 0
}

// SetRrie RRIE
func (r *registerDsiwierType) SetRrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwierFieldRrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwierFieldRrieMask)
	}
}

// registerDsiwisrType DSI wrapper interrupt and status register
type registerDsiwisrType uint32

const (
	RegisterDsiwisrFieldTeifShift = 0
	RegisterDsiwisrFieldTeifMask  = 0x1
)

// GetTeif TEIF
func (r *registerDsiwisrType) GetTeif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldTeifMask) != 0
}

// SetTeif TEIF
func (r *registerDsiwisrType) SetTeif(value bool) {
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
func (r *registerDsiwisrType) GetErif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldErifMask) != 0
}

// SetErif ERIF
func (r *registerDsiwisrType) SetErif(value bool) {
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
func (r *registerDsiwisrType) GetBusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldBusyMask) != 0
}

// SetBusy BUSY
func (r *registerDsiwisrType) SetBusy(value bool) {
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
func (r *registerDsiwisrType) GetPllls() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldPlllsMask) != 0
}

// SetPllls PLLLS
func (r *registerDsiwisrType) SetPllls(value bool) {
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
func (r *registerDsiwisrType) GetPlllif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldPlllifMask) != 0
}

// SetPlllif PLLLIF
func (r *registerDsiwisrType) SetPlllif(value bool) {
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
func (r *registerDsiwisrType) GetPlluif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldPlluifMask) != 0
}

// SetPlluif PLLUIF
func (r *registerDsiwisrType) SetPlluif(value bool) {
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
func (r *registerDsiwisrType) GetRrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldRrsMask) != 0
}

// SetRrs RRS
func (r *registerDsiwisrType) SetRrs(value bool) {
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
func (r *registerDsiwisrType) GetRrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwisrFieldRrifMask) != 0
}

// SetRrif RRIF
func (r *registerDsiwisrType) SetRrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwisrFieldRrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwisrFieldRrifMask)
	}
}

// registerDsiwifcrType DSI wrapper interrupt flag clear register
type registerDsiwifcrType uint32

const (
	RegisterDsiwifcrFieldCteifShift = 0
	RegisterDsiwifcrFieldCteifMask  = 0x1
)

// GetCteif CTEIF
func (r *registerDsiwifcrType) GetCteif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwifcrFieldCteifMask) != 0
}

// SetCteif CTEIF
func (r *registerDsiwifcrType) SetCteif(value bool) {
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
func (r *registerDsiwifcrType) GetCerif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwifcrFieldCerifMask) != 0
}

// SetCerif CERIF
func (r *registerDsiwifcrType) SetCerif(value bool) {
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
func (r *registerDsiwifcrType) GetCplllif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwifcrFieldCplllifMask) != 0
}

// SetCplllif CPLLLIF
func (r *registerDsiwifcrType) SetCplllif(value bool) {
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
func (r *registerDsiwifcrType) GetCplluif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwifcrFieldCplluifMask) != 0
}

// SetCplluif CPLLUIF
func (r *registerDsiwifcrType) SetCplluif(value bool) {
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
func (r *registerDsiwifcrType) GetCrrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwifcrFieldCrrifMask) != 0
}

// SetCrrif CRRIF
func (r *registerDsiwifcrType) SetCrrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwifcrFieldCrrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwifcrFieldCrrifMask)
	}
}

// registerDsiwpcr0Type DSI wrapper PHY configuration register 0
type registerDsiwpcr0Type uint32

const (
	RegisterDsiwpcr0FieldUix4Shift = 0
	RegisterDsiwpcr0FieldUix4Mask  = 0x3f
)

// GetUix4 UIX4
func (r *registerDsiwpcr0Type) GetUix4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldUix4Mask) >> RegisterDsiwpcr0FieldUix4Shift)
}

// SetUix4 UIX4
func (r *registerDsiwpcr0Type) SetUix4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldUix4Mask)|(uint32(value)<<RegisterDsiwpcr0FieldUix4Shift))
}

const (
	RegisterDsiwpcr0FieldSwclShift = 6
	RegisterDsiwpcr0FieldSwclMask  = 0x40
)

// GetSwcl SWCL
func (r *registerDsiwpcr0Type) GetSwcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldSwclMask) != 0
}

// SetSwcl SWCL
func (r *registerDsiwpcr0Type) SetSwcl(value bool) {
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
func (r *registerDsiwpcr0Type) GetSwdl0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldSwdl0Mask) != 0
}

// SetSwdl0 SWDL0
func (r *registerDsiwpcr0Type) SetSwdl0(value bool) {
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
func (r *registerDsiwpcr0Type) GetSwdl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldSwdl1Mask) != 0
}

// SetSwdl1 SWDL1
func (r *registerDsiwpcr0Type) SetSwdl1(value bool) {
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
func (r *registerDsiwpcr0Type) GetHsicl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldHsiclMask) != 0
}

// SetHsicl HSICL
func (r *registerDsiwpcr0Type) SetHsicl(value bool) {
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
func (r *registerDsiwpcr0Type) GetHsidl0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldHsidl0Mask) != 0
}

// SetHsidl0 HSIDL0
func (r *registerDsiwpcr0Type) SetHsidl0(value bool) {
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
func (r *registerDsiwpcr0Type) GetHsidl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldHsidl1Mask) != 0
}

// SetHsidl1 HSIDL1
func (r *registerDsiwpcr0Type) SetHsidl1(value bool) {
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
func (r *registerDsiwpcr0Type) GetFtxsmcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldFtxsmclMask) != 0
}

// SetFtxsmcl FTXSMCL
func (r *registerDsiwpcr0Type) SetFtxsmcl(value bool) {
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
func (r *registerDsiwpcr0Type) GetFtxsmdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldFtxsmdlMask) != 0
}

// SetFtxsmdl FTXSMDL
func (r *registerDsiwpcr0Type) SetFtxsmdl(value bool) {
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
func (r *registerDsiwpcr0Type) GetCdoffdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldCdoffdlMask) != 0
}

// SetCdoffdl CDOFFDL
func (r *registerDsiwpcr0Type) SetCdoffdl(value bool) {
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
func (r *registerDsiwpcr0Type) GetTddl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTddlMask) != 0
}

// SetTddl TDDL
func (r *registerDsiwpcr0Type) SetTddl(value bool) {
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
func (r *registerDsiwpcr0Type) GetPden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldPdenMask) != 0
}

// SetPden Pull-down enable
func (r *registerDsiwpcr0Type) SetPden(value bool) {
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
func (r *registerDsiwpcr0Type) GetTclkprepen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTclkprepenMask) != 0
}

// SetTclkprepen Custom time for tCLK-PREPARE enable
func (r *registerDsiwpcr0Type) SetTclkprepen(value bool) {
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
func (r *registerDsiwpcr0Type) GetTclkzeroen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTclkzeroenMask) != 0
}

// SetTclkzeroen Custom time for tCLK-ZERO enable
func (r *registerDsiwpcr0Type) SetTclkzeroen(value bool) {
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
func (r *registerDsiwpcr0Type) GetThsprepen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldThsprepenMask) != 0
}

// SetThsprepen Custom time for tHS-PREPARE enable
func (r *registerDsiwpcr0Type) SetThsprepen(value bool) {
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
func (r *registerDsiwpcr0Type) GetThstrailen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldThstrailenMask) != 0
}

// SetThstrailen Custom time for tHS-TRAIL enable
func (r *registerDsiwpcr0Type) SetThstrailen(value bool) {
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
func (r *registerDsiwpcr0Type) GetThszeroen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldThszeroenMask) != 0
}

// SetThszeroen Custom time for tHS-ZERO enable
func (r *registerDsiwpcr0Type) SetThszeroen(value bool) {
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
func (r *registerDsiwpcr0Type) GetTlpxden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTlpxdenMask) != 0
}

// SetTlpxden Custom time for tLPX for data lanes enable
func (r *registerDsiwpcr0Type) SetTlpxden(value bool) {
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
func (r *registerDsiwpcr0Type) GetThsexiten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldThsexitenMask) != 0
}

// SetThsexiten Custom time for tHS-EXIT enable
func (r *registerDsiwpcr0Type) SetThsexiten(value bool) {
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
func (r *registerDsiwpcr0Type) GetTlpxcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTlpxcenMask) != 0
}

// SetTlpxcen Custom time for tLPX for clock lane enable
func (r *registerDsiwpcr0Type) SetTlpxcen(value bool) {
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
func (r *registerDsiwpcr0Type) GetTclkposten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr0FieldTclkpostenMask) != 0
}

// SetTclkposten Custom time for tCLK-POST enable
func (r *registerDsiwpcr0Type) SetTclkposten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwpcr0FieldTclkpostenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr0FieldTclkpostenMask)
	}
}

// registerDsiwpcr1Type This register shall be programmed only when DSI is stopped (CR. DSIEN=0 and CR.EN = 0).
type registerDsiwpcr1Type uint32

const (
	RegisterDsiwpcr1FieldHstxdclShift = 0
	RegisterDsiwpcr1FieldHstxdclMask  = 0x3
)

// GetHstxdcl High-speed transmission delay on clock lane
func (r *registerDsiwpcr1Type) GetHstxdcl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldHstxdclMask) >> RegisterDsiwpcr1FieldHstxdclShift)
}

// SetHstxdcl High-speed transmission delay on clock lane
func (r *registerDsiwpcr1Type) SetHstxdcl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldHstxdclMask)|(uint32(value)<<RegisterDsiwpcr1FieldHstxdclShift))
}

const (
	RegisterDsiwpcr1FieldHstxddlShift = 2
	RegisterDsiwpcr1FieldHstxddlMask  = 0xc
)

// GetHstxddl High-speed transmission delay on data lanes
func (r *registerDsiwpcr1Type) GetHstxddl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldHstxddlMask) >> RegisterDsiwpcr1FieldHstxddlShift)
}

// SetHstxddl High-speed transmission delay on data lanes
func (r *registerDsiwpcr1Type) SetHstxddl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldHstxddlMask)|(uint32(value)<<RegisterDsiwpcr1FieldHstxddlShift))
}

const (
	RegisterDsiwpcr1FieldLpsrcclShift = 6
	RegisterDsiwpcr1FieldLpsrcclMask  = 0xc0
)

// GetLpsrccl Low-power transmission slew-rate compensation on clock lane
func (r *registerDsiwpcr1Type) GetLpsrccl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldLpsrcclMask) >> RegisterDsiwpcr1FieldLpsrcclShift)
}

// SetLpsrccl Low-power transmission slew-rate compensation on clock lane
func (r *registerDsiwpcr1Type) SetLpsrccl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldLpsrcclMask)|(uint32(value)<<RegisterDsiwpcr1FieldLpsrcclShift))
}

const (
	RegisterDsiwpcr1FieldLpsrcdlShift = 8
	RegisterDsiwpcr1FieldLpsrcdlMask  = 0x300
)

// GetLpsrcdl Low-power transmission slew-rate compensation on data lanes
func (r *registerDsiwpcr1Type) GetLpsrcdl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldLpsrcdlMask) >> RegisterDsiwpcr1FieldLpsrcdlShift)
}

// SetLpsrcdl Low-power transmission slew-rate compensation on data lanes
func (r *registerDsiwpcr1Type) SetLpsrcdl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldLpsrcdlMask)|(uint32(value)<<RegisterDsiwpcr1FieldLpsrcdlShift))
}

const (
	RegisterDsiwpcr1FieldSddcShift = 12
	RegisterDsiwpcr1FieldSddcMask  = 0x1000
)

// GetSddc SDD control
func (r *registerDsiwpcr1Type) GetSddc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldSddcMask) != 0
}

// SetSddc SDD control
func (r *registerDsiwpcr1Type) SetSddc(value bool) {
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
func (r *registerDsiwpcr1Type) GetHstxsrccl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldHstxsrcclMask) >> RegisterDsiwpcr1FieldHstxsrcclShift)
}

// SetHstxsrccl High-speed transmission slew-rate control on clock lane
func (r *registerDsiwpcr1Type) SetHstxsrccl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldHstxsrcclMask)|(uint32(value)<<RegisterDsiwpcr1FieldHstxsrcclShift))
}

const (
	RegisterDsiwpcr1FieldHstxsrcdlShift = 18
	RegisterDsiwpcr1FieldHstxsrcdlMask  = 0xc0000
)

// GetHstxsrcdl High-speed transmission slew-rate control on data lanes
func (r *registerDsiwpcr1Type) GetHstxsrcdl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldHstxsrcdlMask) >> RegisterDsiwpcr1FieldHstxsrcdlShift)
}

// SetHstxsrcdl High-speed transmission slew-rate control on data lanes
func (r *registerDsiwpcr1Type) SetHstxsrcdl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldHstxsrcdlMask)|(uint32(value)<<RegisterDsiwpcr1FieldHstxsrcdlShift))
}

const (
	RegisterDsiwpcr1FieldFlprxlpmShift = 22
	RegisterDsiwpcr1FieldFlprxlpmMask  = 0x400000
)

// GetFlprxlpm Forces LP receiver in low-power mode
func (r *registerDsiwpcr1Type) GetFlprxlpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldFlprxlpmMask) != 0
}

// SetFlprxlpm Forces LP receiver in low-power mode
func (r *registerDsiwpcr1Type) SetFlprxlpm(value bool) {
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
func (r *registerDsiwpcr1Type) GetLprxft() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr1FieldLprxftMask) >> RegisterDsiwpcr1FieldLprxftShift)
}

// SetLprxft Low-power RX low-pass filtering tuning
func (r *registerDsiwpcr1Type) SetLprxft(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr1FieldLprxftMask)|(uint32(value)<<RegisterDsiwpcr1FieldLprxftShift))
}

// registerDsiwpcr2Type DSI wrapper PHY configuration register 2
type registerDsiwpcr2Type uint32

const (
	RegisterDsiwpcr2FieldTclkprepShift = 0
	RegisterDsiwpcr2FieldTclkprepMask  = 0xff
)

// GetTclkprep TCLKPREP
func (r *registerDsiwpcr2Type) GetTclkprep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr2FieldTclkprepMask) >> RegisterDsiwpcr2FieldTclkprepShift)
}

// SetTclkprep TCLKPREP
func (r *registerDsiwpcr2Type) SetTclkprep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr2FieldTclkprepMask)|(uint32(value)<<RegisterDsiwpcr2FieldTclkprepShift))
}

const (
	RegisterDsiwpcr2FieldTclkzeroShift = 8
	RegisterDsiwpcr2FieldTclkzeroMask  = 0xff00
)

// GetTclkzero TCLKZERO
func (r *registerDsiwpcr2Type) GetTclkzero() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr2FieldTclkzeroMask) >> RegisterDsiwpcr2FieldTclkzeroShift)
}

// SetTclkzero TCLKZERO
func (r *registerDsiwpcr2Type) SetTclkzero(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr2FieldTclkzeroMask)|(uint32(value)<<RegisterDsiwpcr2FieldTclkzeroShift))
}

const (
	RegisterDsiwpcr2FieldThsprepShift = 16
	RegisterDsiwpcr2FieldThsprepMask  = 0xff0000
)

// GetThsprep THSPREP
func (r *registerDsiwpcr2Type) GetThsprep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr2FieldThsprepMask) >> RegisterDsiwpcr2FieldThsprepShift)
}

// SetThsprep THSPREP
func (r *registerDsiwpcr2Type) SetThsprep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr2FieldThsprepMask)|(uint32(value)<<RegisterDsiwpcr2FieldThsprepShift))
}

const (
	RegisterDsiwpcr2FieldThstrailShift = 24
	RegisterDsiwpcr2FieldThstrailMask  = 0xff000000
)

// GetThstrail THSTRAIL
func (r *registerDsiwpcr2Type) GetThstrail() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr2FieldThstrailMask) >> RegisterDsiwpcr2FieldThstrailShift)
}

// SetThstrail THSTRAIL
func (r *registerDsiwpcr2Type) SetThstrail(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr2FieldThstrailMask)|(uint32(value)<<RegisterDsiwpcr2FieldThstrailShift))
}

// registerDsiwpcr3Type DSI wrapper PHY configuration register 3
type registerDsiwpcr3Type uint32

const (
	RegisterDsiwpcr3FieldThszeroShift = 0
	RegisterDsiwpcr3FieldThszeroMask  = 0xff
)

// GetThszero THSZERO
func (r *registerDsiwpcr3Type) GetThszero() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr3FieldThszeroMask) >> RegisterDsiwpcr3FieldThszeroShift)
}

// SetThszero THSZERO
func (r *registerDsiwpcr3Type) SetThszero(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr3FieldThszeroMask)|(uint32(value)<<RegisterDsiwpcr3FieldThszeroShift))
}

const (
	RegisterDsiwpcr3FieldTlpxdShift = 8
	RegisterDsiwpcr3FieldTlpxdMask  = 0xff00
)

// GetTlpxd TLPXD
func (r *registerDsiwpcr3Type) GetTlpxd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr3FieldTlpxdMask) >> RegisterDsiwpcr3FieldTlpxdShift)
}

// SetTlpxd TLPXD
func (r *registerDsiwpcr3Type) SetTlpxd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr3FieldTlpxdMask)|(uint32(value)<<RegisterDsiwpcr3FieldTlpxdShift))
}

const (
	RegisterDsiwpcr3FieldThsexitShift = 16
	RegisterDsiwpcr3FieldThsexitMask  = 0xff0000
)

// GetThsexit THSEXIT
func (r *registerDsiwpcr3Type) GetThsexit() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr3FieldThsexitMask) >> RegisterDsiwpcr3FieldThsexitShift)
}

// SetThsexit THSEXIT
func (r *registerDsiwpcr3Type) SetThsexit(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr3FieldThsexitMask)|(uint32(value)<<RegisterDsiwpcr3FieldThsexitShift))
}

const (
	RegisterDsiwpcr3FieldTlpxcShift = 24
	RegisterDsiwpcr3FieldTlpxcMask  = 0xff000000
)

// GetTlpxc TLPXC
func (r *registerDsiwpcr3Type) GetTlpxc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr3FieldTlpxcMask) >> RegisterDsiwpcr3FieldTlpxcShift)
}

// SetTlpxc TLPXC
func (r *registerDsiwpcr3Type) SetTlpxc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr3FieldTlpxcMask)|(uint32(value)<<RegisterDsiwpcr3FieldTlpxcShift))
}

// registerDsiwpcr4Type DSI wrapper PHY configuration register 4
type registerDsiwpcr4Type uint32

const (
	RegisterDsiwpcr4FieldTclkpostShift = 0
	RegisterDsiwpcr4FieldTclkpostMask  = 0xff
)

// GetTclkpost TCLKPOST
func (r *registerDsiwpcr4Type) GetTclkpost() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwpcr4FieldTclkpostMask) >> RegisterDsiwpcr4FieldTclkpostShift)
}

// SetTclkpost TCLKPOST
func (r *registerDsiwpcr4Type) SetTclkpost(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwpcr4FieldTclkpostMask)|(uint32(value)<<RegisterDsiwpcr4FieldTclkpostShift))
}

// registerDsiwrpcrType DSI wrapper regulator and PLL control register
type registerDsiwrpcrType uint32

const (
	RegisterDsiwrpcrFieldPllenShift = 0
	RegisterDsiwrpcrFieldPllenMask  = 0x1
)

// GetPllen PLLEN
func (r *registerDsiwrpcrType) GetPllen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwrpcrFieldPllenMask) != 0
}

// SetPllen PLLEN
func (r *registerDsiwrpcrType) SetPllen(value bool) {
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
func (r *registerDsiwrpcrType) GetNdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwrpcrFieldNdivMask) >> RegisterDsiwrpcrFieldNdivShift)
}

// SetNdiv NDIV
func (r *registerDsiwrpcrType) SetNdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwrpcrFieldNdivMask)|(uint32(value)<<RegisterDsiwrpcrFieldNdivShift))
}

const (
	RegisterDsiwrpcrFieldIdfShift = 11
	RegisterDsiwrpcrFieldIdfMask  = 0x7800
)

// GetIdf IDF
func (r *registerDsiwrpcrType) GetIdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwrpcrFieldIdfMask) >> RegisterDsiwrpcrFieldIdfShift)
}

// SetIdf IDF
func (r *registerDsiwrpcrType) SetIdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwrpcrFieldIdfMask)|(uint32(value)<<RegisterDsiwrpcrFieldIdfShift))
}

const (
	RegisterDsiwrpcrFieldOdfShift = 16
	RegisterDsiwrpcrFieldOdfMask  = 0x30000
)

// GetOdf ODF
func (r *registerDsiwrpcrType) GetOdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsiwrpcrFieldOdfMask) >> RegisterDsiwrpcrFieldOdfShift)
}

// SetOdf ODF
func (r *registerDsiwrpcrType) SetOdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsiwrpcrFieldOdfMask)|(uint32(value)<<RegisterDsiwrpcrFieldOdfShift))
}

const (
	RegisterDsiwrpcrFieldRegenShift = 24
	RegisterDsiwrpcrFieldRegenMask  = 0x1000000
)

// GetRegen REGEN
func (r *registerDsiwrpcrType) GetRegen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsiwrpcrFieldRegenMask) != 0
}

// SetRegen REGEN
func (r *registerDsiwrpcrType) SetRegen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsiwrpcrFieldRegenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsiwrpcrFieldRegenMask)
	}
}
