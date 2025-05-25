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
	Dsi_vr      registerDsi_vrType
	Dsi_cr      registerDsi_crType
	Dsi_ccr     registerDsi_ccrType
	Dsi_lvcidr  registerDsi_lvcidrType
	Dsi_lcolcr  registerDsi_lcolcrType
	Dsi_lpcr    registerDsi_lpcrType
	Dsi_lpmcr   registerDsi_lpmcrType
	_           [16]byte
	Dsi_pcr     registerDsi_pcrType
	Dsi_gvcidr  registerDsi_gvcidrType
	Dsi_mcr     registerDsi_mcrType
	Dsi_vmcr    registerDsi_vmcrType
	Dsi_vpcr    registerDsi_vpcrType
	Dsi_vccr    registerDsi_vccrType
	Dsi_vnpcr   registerDsi_vnpcrType
	Dsi_vhsacr  registerDsi_vhsacrType
	Dsi_vhbpcr  registerDsi_vhbpcrType
	Dsi_vlcr    registerDsi_vlcrType
	Dsi_vvsacr  registerDsi_vvsacrType
	Dsi_vvbpcr  registerDsi_vvbpcrType
	Dsi_vvfpcr  registerDsi_vvfpcrType
	Dsi_vvacr   registerDsi_vvacrType
	Dsi_lccr    registerDsi_lccrType
	Dsi_cmcr    registerDsi_cmcrType
	Dsi_ghcr    registerDsi_ghcrType
	Dsi_gpdr    registerDsi_gpdrType
	Dsi_gpsr    registerDsi_gpsrType
	Dsi_tccr0   registerDsi_tccr0Type
	Dsi_tccr1   registerDsi_tccr1Type
	Dsi_tccr2   registerDsi_tccr2Type
	Dsi_tccr3   registerDsi_tccr3Type
	Dsi_tccr4   registerDsi_tccr4Type
	Dsi_tccr5   registerDsi_tccr5Type
	_           [4]byte
	Dsi_clcr    registerDsi_clcrType
	Dsi_cltcr   registerDsi_cltcrType
	Dsi_dltcr   registerDsi_dltcrType
	Dsi_pctlr   registerDsi_pctlrType
	Dsi_pconfr  registerDsi_pconfrType
	Dsi_pucr    registerDsi_pucrType
	Dsi_pttcr   registerDsi_pttcrType
	Dsi_psr     registerDsi_psrType
	_           [8]byte
	Dsi_isr0    registerDsi_isr0Type
	Dsi_isr1    registerDsi_isr1Type
	Dsi_ier0    registerDsi_ier0Type
	Dsi_ier1    registerDsi_ier1Type
	_           [12]byte
	Dsi_fir0    registerDsi_fir0Type
	Dsi_fir1    registerDsi_fir1Type
	_           [32]byte
	Dsi_vscr    registerDsi_vscrType
	_           [8]byte
	Dsi_lcvcidr registerDsi_lcvcidrType
	Dsi_lcccr   registerDsi_lcccrType
	_           [4]byte
	Dsi_lpmccr  registerDsi_lpmccrType
	_           [28]byte
	Dsi_vmccr   registerDsi_vmccrType
	Dsi_vpccr   registerDsi_vpccrType
	Dsi_vcccr   registerDsi_vcccrType
	Dsi_vnpccr  registerDsi_vnpccrType
	Dsi_vhsaccr registerDsi_vhsaccrType
	Dsi_vhbpccr registerDsi_vhbpccrType
	Dsi_vlccr   registerDsi_vlccrType
	Dsi_vvsaccr registerDsi_vvsaccrType
	Dsi_vvbpccr registerDsi_vvbpccrType
	Dsi_vvfpccr registerDsi_vvfpccrType
	Dsi_vvaccr  registerDsi_vvaccrType
	_           [668]byte
	Dsi_wcfgr   registerDsi_wcfgrType
	Dsi_wcr     registerDsi_wcrType
	Dsi_wier    registerDsi_wierType
	Dsi_wisr    registerDsi_wisrType
	Dsi_wifcr   registerDsi_wifcrType
	_           [4]byte
	Dsi_wpcr0   registerDsi_wpcr0Type
	Dsi_wpcr1   registerDsi_wpcr1Type
	Dsi_wpcr2   registerDsi_wpcr2Type
	Dsi_wpcr3   registerDsi_wpcr3Type
	Dsi_wpcr4   registerDsi_wpcr4Type
	_           [4]byte
	Dsi_wrpcr   registerDsi_wrpcrType
}

// registerDsi_vrType DSI Host version register
type registerDsi_vrType uint32

const (
	RegisterDsi_vrFieldVersionShift = 0
	RegisterDsi_vrFieldVersionMask  = 0xffffffff
)

// GetVersion VERSION
func (r *registerDsi_vrType) GetVersion() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vrFieldVersionMask) >> RegisterDsi_vrFieldVersionShift)
}

// SetVersion VERSION
func (r *registerDsi_vrType) SetVersion(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vrFieldVersionMask)|(uint32(value)<<RegisterDsi_vrFieldVersionShift))
}

// registerDsi_crType DSI Host control register
type registerDsi_crType uint32

const (
	RegisterDsi_crFieldEnShift = 0
	RegisterDsi_crFieldEnMask  = 0x1
)

// GetEn EN
func (r *registerDsi_crType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_crFieldEnMask) != 0
}

// SetEn EN
func (r *registerDsi_crType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_crFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_crFieldEnMask)
	}
}

// registerDsi_ccrType DSI Host clock control register
type registerDsi_ccrType uint32

const (
	RegisterDsi_ccrFieldTxeckdivShift = 0
	RegisterDsi_ccrFieldTxeckdivMask  = 0xff
)

// GetTxeckdiv TXECKDIV
func (r *registerDsi_ccrType) GetTxeckdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ccrFieldTxeckdivMask) >> RegisterDsi_ccrFieldTxeckdivShift)
}

// SetTxeckdiv TXECKDIV
func (r *registerDsi_ccrType) SetTxeckdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ccrFieldTxeckdivMask)|(uint32(value)<<RegisterDsi_ccrFieldTxeckdivShift))
}

const (
	RegisterDsi_ccrFieldTockdivShift = 8
	RegisterDsi_ccrFieldTockdivMask  = 0xff00
)

// GetTockdiv TOCKDIV
func (r *registerDsi_ccrType) GetTockdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ccrFieldTockdivMask) >> RegisterDsi_ccrFieldTockdivShift)
}

// SetTockdiv TOCKDIV
func (r *registerDsi_ccrType) SetTockdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ccrFieldTockdivMask)|(uint32(value)<<RegisterDsi_ccrFieldTockdivShift))
}

// registerDsi_lvcidrType DSI Host LTDC VCID register
type registerDsi_lvcidrType uint32

const (
	RegisterDsi_lvcidrFieldVcidShift = 0
	RegisterDsi_lvcidrFieldVcidMask  = 0x3
)

// GetVcid VCID
func (r *registerDsi_lvcidrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lvcidrFieldVcidMask) >> RegisterDsi_lvcidrFieldVcidShift)
}

// SetVcid VCID
func (r *registerDsi_lvcidrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lvcidrFieldVcidMask)|(uint32(value)<<RegisterDsi_lvcidrFieldVcidShift))
}

// registerDsi_lcolcrType DSI Host LTDC color coding register
type registerDsi_lcolcrType uint32

const (
	RegisterDsi_lcolcrFieldColcShift = 0
	RegisterDsi_lcolcrFieldColcMask  = 0xf
)

// GetColc COLC
func (r *registerDsi_lcolcrType) GetColc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lcolcrFieldColcMask) >> RegisterDsi_lcolcrFieldColcShift)
}

// SetColc COLC
func (r *registerDsi_lcolcrType) SetColc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lcolcrFieldColcMask)|(uint32(value)<<RegisterDsi_lcolcrFieldColcShift))
}

const (
	RegisterDsi_lcolcrFieldLpeShift = 8
	RegisterDsi_lcolcrFieldLpeMask  = 0x100
)

// GetLpe LPE
func (r *registerDsi_lcolcrType) GetLpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lcolcrFieldLpeMask) != 0
}

// SetLpe LPE
func (r *registerDsi_lcolcrType) SetLpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_lcolcrFieldLpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lcolcrFieldLpeMask)
	}
}

// registerDsi_lpcrType DSI Host LTDC polarity configuration register
type registerDsi_lpcrType uint32

const (
	RegisterDsi_lpcrFieldDepShift = 0
	RegisterDsi_lpcrFieldDepMask  = 0x1
)

// GetDep DEP
func (r *registerDsi_lpcrType) GetDep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lpcrFieldDepMask) != 0
}

// SetDep DEP
func (r *registerDsi_lpcrType) SetDep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_lpcrFieldDepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lpcrFieldDepMask)
	}
}

const (
	RegisterDsi_lpcrFieldVspShift = 1
	RegisterDsi_lpcrFieldVspMask  = 0x2
)

// GetVsp VSP
func (r *registerDsi_lpcrType) GetVsp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lpcrFieldVspMask) != 0
}

// SetVsp VSP
func (r *registerDsi_lpcrType) SetVsp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_lpcrFieldVspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lpcrFieldVspMask)
	}
}

const (
	RegisterDsi_lpcrFieldHspShift = 2
	RegisterDsi_lpcrFieldHspMask  = 0x4
)

// GetHsp HSP
func (r *registerDsi_lpcrType) GetHsp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lpcrFieldHspMask) != 0
}

// SetHsp HSP
func (r *registerDsi_lpcrType) SetHsp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_lpcrFieldHspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lpcrFieldHspMask)
	}
}

// registerDsi_lpmcrType DSI Host low-power mode configuration register
type registerDsi_lpmcrType uint32

const (
	RegisterDsi_lpmcrFieldVlpsizeShift = 0
	RegisterDsi_lpmcrFieldVlpsizeMask  = 0xff
)

// GetVlpsize VLPSIZE
func (r *registerDsi_lpmcrType) GetVlpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lpmcrFieldVlpsizeMask) >> RegisterDsi_lpmcrFieldVlpsizeShift)
}

// SetVlpsize VLPSIZE
func (r *registerDsi_lpmcrType) SetVlpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lpmcrFieldVlpsizeMask)|(uint32(value)<<RegisterDsi_lpmcrFieldVlpsizeShift))
}

const (
	RegisterDsi_lpmcrFieldLpsizeShift = 16
	RegisterDsi_lpmcrFieldLpsizeMask  = 0xff0000
)

// GetLpsize LPSIZE
func (r *registerDsi_lpmcrType) GetLpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lpmcrFieldLpsizeMask) >> RegisterDsi_lpmcrFieldLpsizeShift)
}

// SetLpsize LPSIZE
func (r *registerDsi_lpmcrType) SetLpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lpmcrFieldLpsizeMask)|(uint32(value)<<RegisterDsi_lpmcrFieldLpsizeShift))
}

// registerDsi_pcrType DSI Host protocol configuration register
type registerDsi_pcrType uint32

const (
	RegisterDsi_pcrFieldEttxeShift = 0
	RegisterDsi_pcrFieldEttxeMask  = 0x1
)

// GetEttxe ETTXE
func (r *registerDsi_pcrType) GetEttxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pcrFieldEttxeMask) != 0
}

// SetEttxe ETTXE
func (r *registerDsi_pcrType) SetEttxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_pcrFieldEttxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pcrFieldEttxeMask)
	}
}

const (
	RegisterDsi_pcrFieldEtrxeShift = 1
	RegisterDsi_pcrFieldEtrxeMask  = 0x2
)

// GetEtrxe ETRXE
func (r *registerDsi_pcrType) GetEtrxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pcrFieldEtrxeMask) != 0
}

// SetEtrxe ETRXE
func (r *registerDsi_pcrType) SetEtrxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_pcrFieldEtrxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pcrFieldEtrxeMask)
	}
}

const (
	RegisterDsi_pcrFieldBtaeShift = 2
	RegisterDsi_pcrFieldBtaeMask  = 0x4
)

// GetBtae BTAE
func (r *registerDsi_pcrType) GetBtae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pcrFieldBtaeMask) != 0
}

// SetBtae BTAE
func (r *registerDsi_pcrType) SetBtae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_pcrFieldBtaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pcrFieldBtaeMask)
	}
}

const (
	RegisterDsi_pcrFieldEccrxeShift = 3
	RegisterDsi_pcrFieldEccrxeMask  = 0x8
)

// GetEccrxe ECCRXE
func (r *registerDsi_pcrType) GetEccrxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pcrFieldEccrxeMask) != 0
}

// SetEccrxe ECCRXE
func (r *registerDsi_pcrType) SetEccrxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_pcrFieldEccrxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pcrFieldEccrxeMask)
	}
}

const (
	RegisterDsi_pcrFieldCrcrxeShift = 4
	RegisterDsi_pcrFieldCrcrxeMask  = 0x10
)

// GetCrcrxe CRCRXE
func (r *registerDsi_pcrType) GetCrcrxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pcrFieldCrcrxeMask) != 0
}

// SetCrcrxe CRCRXE
func (r *registerDsi_pcrType) SetCrcrxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_pcrFieldCrcrxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pcrFieldCrcrxeMask)
	}
}

// registerDsi_gvcidrType DSI Host generic VCID register
type registerDsi_gvcidrType uint32

const (
	RegisterDsi_gvcidrFieldVcidShift = 0
	RegisterDsi_gvcidrFieldVcidMask  = 0x3
)

// GetVcid VCID
func (r *registerDsi_gvcidrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gvcidrFieldVcidMask) >> RegisterDsi_gvcidrFieldVcidShift)
}

// SetVcid VCID
func (r *registerDsi_gvcidrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gvcidrFieldVcidMask)|(uint32(value)<<RegisterDsi_gvcidrFieldVcidShift))
}

// registerDsi_mcrType DSI Host mode configuration register
type registerDsi_mcrType uint32

const (
	RegisterDsi_mcrFieldCmdmShift = 0
	RegisterDsi_mcrFieldCmdmMask  = 0x1
)

// GetCmdm CMDM
func (r *registerDsi_mcrType) GetCmdm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_mcrFieldCmdmMask) != 0
}

// SetCmdm CMDM
func (r *registerDsi_mcrType) SetCmdm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_mcrFieldCmdmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_mcrFieldCmdmMask)
	}
}

// registerDsi_vmcrType DSI Host video mode configuration register
type registerDsi_vmcrType uint32

const (
	RegisterDsi_vmcrFieldVmtShift = 0
	RegisterDsi_vmcrFieldVmtMask  = 0x3
)

// GetVmt VMT
func (r *registerDsi_vmcrType) GetVmt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldVmtMask) >> RegisterDsi_vmcrFieldVmtShift)
}

// SetVmt VMT
func (r *registerDsi_vmcrType) SetVmt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldVmtMask)|(uint32(value)<<RegisterDsi_vmcrFieldVmtShift))
}

const (
	RegisterDsi_vmcrFieldLpvsaeShift = 8
	RegisterDsi_vmcrFieldLpvsaeMask  = 0x100
)

// GetLpvsae LPVSAE
func (r *registerDsi_vmcrType) GetLpvsae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldLpvsaeMask) != 0
}

// SetLpvsae LPVSAE
func (r *registerDsi_vmcrType) SetLpvsae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmcrFieldLpvsaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldLpvsaeMask)
	}
}

const (
	RegisterDsi_vmcrFieldLpvbpeShift = 9
	RegisterDsi_vmcrFieldLpvbpeMask  = 0x200
)

// GetLpvbpe LPVBPE
func (r *registerDsi_vmcrType) GetLpvbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldLpvbpeMask) != 0
}

// SetLpvbpe LPVBPE
func (r *registerDsi_vmcrType) SetLpvbpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmcrFieldLpvbpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldLpvbpeMask)
	}
}

const (
	RegisterDsi_vmcrFieldLpvfpeShift = 10
	RegisterDsi_vmcrFieldLpvfpeMask  = 0x400
)

// GetLpvfpe LPVFPE
func (r *registerDsi_vmcrType) GetLpvfpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldLpvfpeMask) != 0
}

// SetLpvfpe LPVFPE
func (r *registerDsi_vmcrType) SetLpvfpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmcrFieldLpvfpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldLpvfpeMask)
	}
}

const (
	RegisterDsi_vmcrFieldLpvaeShift = 11
	RegisterDsi_vmcrFieldLpvaeMask  = 0x800
)

// GetLpvae LPVAE
func (r *registerDsi_vmcrType) GetLpvae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldLpvaeMask) != 0
}

// SetLpvae LPVAE
func (r *registerDsi_vmcrType) SetLpvae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmcrFieldLpvaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldLpvaeMask)
	}
}

const (
	RegisterDsi_vmcrFieldLphbpeShift = 12
	RegisterDsi_vmcrFieldLphbpeMask  = 0x1000
)

// GetLphbpe LPHBPE
func (r *registerDsi_vmcrType) GetLphbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldLphbpeMask) != 0
}

// SetLphbpe LPHBPE
func (r *registerDsi_vmcrType) SetLphbpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmcrFieldLphbpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldLphbpeMask)
	}
}

const (
	RegisterDsi_vmcrFieldLphfpeShift = 13
	RegisterDsi_vmcrFieldLphfpeMask  = 0x2000
)

// GetLphfpe LPHFPE
func (r *registerDsi_vmcrType) GetLphfpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldLphfpeMask) != 0
}

// SetLphfpe LPHFPE
func (r *registerDsi_vmcrType) SetLphfpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmcrFieldLphfpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldLphfpeMask)
	}
}

const (
	RegisterDsi_vmcrFieldFbtaaeShift = 14
	RegisterDsi_vmcrFieldFbtaaeMask  = 0x4000
)

// GetFbtaae FBTAAE
func (r *registerDsi_vmcrType) GetFbtaae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldFbtaaeMask) != 0
}

// SetFbtaae FBTAAE
func (r *registerDsi_vmcrType) SetFbtaae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmcrFieldFbtaaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldFbtaaeMask)
	}
}

const (
	RegisterDsi_vmcrFieldLpceShift = 15
	RegisterDsi_vmcrFieldLpceMask  = 0x8000
)

// GetLpce LPCE
func (r *registerDsi_vmcrType) GetLpce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldLpceMask) != 0
}

// SetLpce LPCE
func (r *registerDsi_vmcrType) SetLpce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmcrFieldLpceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldLpceMask)
	}
}

const (
	RegisterDsi_vmcrFieldPgeShift = 16
	RegisterDsi_vmcrFieldPgeMask  = 0x10000
)

// GetPge PGE
func (r *registerDsi_vmcrType) GetPge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldPgeMask) != 0
}

// SetPge PGE
func (r *registerDsi_vmcrType) SetPge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmcrFieldPgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldPgeMask)
	}
}

const (
	RegisterDsi_vmcrFieldPgmShift = 20
	RegisterDsi_vmcrFieldPgmMask  = 0x100000
)

// GetPgm PGM
func (r *registerDsi_vmcrType) GetPgm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldPgmMask) != 0
}

// SetPgm PGM
func (r *registerDsi_vmcrType) SetPgm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmcrFieldPgmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldPgmMask)
	}
}

const (
	RegisterDsi_vmcrFieldPgoShift = 24
	RegisterDsi_vmcrFieldPgoMask  = 0x1000000
)

// GetPgo PGO
func (r *registerDsi_vmcrType) GetPgo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmcrFieldPgoMask) != 0
}

// SetPgo PGO
func (r *registerDsi_vmcrType) SetPgo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmcrFieldPgoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmcrFieldPgoMask)
	}
}

// registerDsi_vpcrType DSI Host video packet configuration register
type registerDsi_vpcrType uint32

const (
	RegisterDsi_vpcrFieldVpsizeShift = 0
	RegisterDsi_vpcrFieldVpsizeMask  = 0x3fff
)

// GetVpsize VPSIZE
func (r *registerDsi_vpcrType) GetVpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vpcrFieldVpsizeMask) >> RegisterDsi_vpcrFieldVpsizeShift)
}

// SetVpsize VPSIZE
func (r *registerDsi_vpcrType) SetVpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vpcrFieldVpsizeMask)|(uint32(value)<<RegisterDsi_vpcrFieldVpsizeShift))
}

// registerDsi_vccrType DSI Host video chunks configuration register
type registerDsi_vccrType uint32

const (
	RegisterDsi_vccrFieldNumcShift = 0
	RegisterDsi_vccrFieldNumcMask  = 0x1fff
)

// GetNumc NUMC
func (r *registerDsi_vccrType) GetNumc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vccrFieldNumcMask) >> RegisterDsi_vccrFieldNumcShift)
}

// SetNumc NUMC
func (r *registerDsi_vccrType) SetNumc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vccrFieldNumcMask)|(uint32(value)<<RegisterDsi_vccrFieldNumcShift))
}

// registerDsi_vnpcrType DSI Host video null packet configuration register
type registerDsi_vnpcrType uint32

const (
	RegisterDsi_vnpcrFieldNpsizeShift = 0
	RegisterDsi_vnpcrFieldNpsizeMask  = 0x1fff
)

// GetNpsize NPSIZE
func (r *registerDsi_vnpcrType) GetNpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vnpcrFieldNpsizeMask) >> RegisterDsi_vnpcrFieldNpsizeShift)
}

// SetNpsize NPSIZE
func (r *registerDsi_vnpcrType) SetNpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vnpcrFieldNpsizeMask)|(uint32(value)<<RegisterDsi_vnpcrFieldNpsizeShift))
}

// registerDsi_vhsacrType DSI Host video HSA configuration register
type registerDsi_vhsacrType uint32

const (
	RegisterDsi_vhsacrFieldHsaShift = 0
	RegisterDsi_vhsacrFieldHsaMask  = 0xfff
)

// GetHsa HSA
func (r *registerDsi_vhsacrType) GetHsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vhsacrFieldHsaMask) >> RegisterDsi_vhsacrFieldHsaShift)
}

// SetHsa HSA
func (r *registerDsi_vhsacrType) SetHsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vhsacrFieldHsaMask)|(uint32(value)<<RegisterDsi_vhsacrFieldHsaShift))
}

// registerDsi_vhbpcrType DSI Host video HBP configuration register
type registerDsi_vhbpcrType uint32

const (
	RegisterDsi_vhbpcrFieldHbpShift = 0
	RegisterDsi_vhbpcrFieldHbpMask  = 0xfff
)

// GetHbp HBP
func (r *registerDsi_vhbpcrType) GetHbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vhbpcrFieldHbpMask) >> RegisterDsi_vhbpcrFieldHbpShift)
}

// SetHbp HBP
func (r *registerDsi_vhbpcrType) SetHbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vhbpcrFieldHbpMask)|(uint32(value)<<RegisterDsi_vhbpcrFieldHbpShift))
}

// registerDsi_vlcrType DSI Host video line configuration register
type registerDsi_vlcrType uint32

const (
	RegisterDsi_vlcrFieldHlineShift = 0
	RegisterDsi_vlcrFieldHlineMask  = 0x7fff
)

// GetHline HLINE
func (r *registerDsi_vlcrType) GetHline() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vlcrFieldHlineMask) >> RegisterDsi_vlcrFieldHlineShift)
}

// SetHline HLINE
func (r *registerDsi_vlcrType) SetHline(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vlcrFieldHlineMask)|(uint32(value)<<RegisterDsi_vlcrFieldHlineShift))
}

// registerDsi_vvsacrType DSI Host video VSA configuration register
type registerDsi_vvsacrType uint32

const (
	RegisterDsi_vvsacrFieldVsaShift = 0
	RegisterDsi_vvsacrFieldVsaMask  = 0x3ff
)

// GetVsa VSA
func (r *registerDsi_vvsacrType) GetVsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vvsacrFieldVsaMask) >> RegisterDsi_vvsacrFieldVsaShift)
}

// SetVsa VSA
func (r *registerDsi_vvsacrType) SetVsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vvsacrFieldVsaMask)|(uint32(value)<<RegisterDsi_vvsacrFieldVsaShift))
}

// registerDsi_vvbpcrType DSI Host video VBP configuration register
type registerDsi_vvbpcrType uint32

const (
	RegisterDsi_vvbpcrFieldVbpShift = 0
	RegisterDsi_vvbpcrFieldVbpMask  = 0x3ff
)

// GetVbp VBP
func (r *registerDsi_vvbpcrType) GetVbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vvbpcrFieldVbpMask) >> RegisterDsi_vvbpcrFieldVbpShift)
}

// SetVbp VBP
func (r *registerDsi_vvbpcrType) SetVbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vvbpcrFieldVbpMask)|(uint32(value)<<RegisterDsi_vvbpcrFieldVbpShift))
}

// registerDsi_vvfpcrType DSI Host video VFP configuration register
type registerDsi_vvfpcrType uint32

const (
	RegisterDsi_vvfpcrFieldVfpShift = 0
	RegisterDsi_vvfpcrFieldVfpMask  = 0x3ff
)

// GetVfp VFP
func (r *registerDsi_vvfpcrType) GetVfp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vvfpcrFieldVfpMask) >> RegisterDsi_vvfpcrFieldVfpShift)
}

// SetVfp VFP
func (r *registerDsi_vvfpcrType) SetVfp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vvfpcrFieldVfpMask)|(uint32(value)<<RegisterDsi_vvfpcrFieldVfpShift))
}

// registerDsi_vvacrType DSI Host video VA configuration register
type registerDsi_vvacrType uint32

const (
	RegisterDsi_vvacrFieldVaShift = 0
	RegisterDsi_vvacrFieldVaMask  = 0x3fff
)

// GetVa VA
func (r *registerDsi_vvacrType) GetVa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vvacrFieldVaMask) >> RegisterDsi_vvacrFieldVaShift)
}

// SetVa VA
func (r *registerDsi_vvacrType) SetVa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vvacrFieldVaMask)|(uint32(value)<<RegisterDsi_vvacrFieldVaShift))
}

// registerDsi_lccrType DSI Host LTDC command configuration register
type registerDsi_lccrType uint32

const (
	RegisterDsi_lccrFieldCmdsizeShift = 0
	RegisterDsi_lccrFieldCmdsizeMask  = 0xffff
)

// GetCmdsize CMDSIZE
func (r *registerDsi_lccrType) GetCmdsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lccrFieldCmdsizeMask) >> RegisterDsi_lccrFieldCmdsizeShift)
}

// SetCmdsize CMDSIZE
func (r *registerDsi_lccrType) SetCmdsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lccrFieldCmdsizeMask)|(uint32(value)<<RegisterDsi_lccrFieldCmdsizeShift))
}

// registerDsi_cmcrType DSI Host command mode configuration register
type registerDsi_cmcrType uint32

const (
	RegisterDsi_cmcrFieldTeareShift = 0
	RegisterDsi_cmcrFieldTeareMask  = 0x1
)

// GetTeare TEARE
func (r *registerDsi_cmcrType) GetTeare() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldTeareMask) != 0
}

// SetTeare TEARE
func (r *registerDsi_cmcrType) SetTeare(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldTeareMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldTeareMask)
	}
}

const (
	RegisterDsi_cmcrFieldAreShift = 1
	RegisterDsi_cmcrFieldAreMask  = 0x2
)

// GetAre ARE
func (r *registerDsi_cmcrType) GetAre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldAreMask) != 0
}

// SetAre ARE
func (r *registerDsi_cmcrType) SetAre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldAreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldAreMask)
	}
}

const (
	RegisterDsi_cmcrFieldGsw0txShift = 8
	RegisterDsi_cmcrFieldGsw0txMask  = 0x100
)

// GetGsw0tx GSW0TX
func (r *registerDsi_cmcrType) GetGsw0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldGsw0txMask) != 0
}

// SetGsw0tx GSW0TX
func (r *registerDsi_cmcrType) SetGsw0tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldGsw0txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldGsw0txMask)
	}
}

const (
	RegisterDsi_cmcrFieldGsw1txShift = 9
	RegisterDsi_cmcrFieldGsw1txMask  = 0x200
)

// GetGsw1tx GSW1TX
func (r *registerDsi_cmcrType) GetGsw1tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldGsw1txMask) != 0
}

// SetGsw1tx GSW1TX
func (r *registerDsi_cmcrType) SetGsw1tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldGsw1txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldGsw1txMask)
	}
}

const (
	RegisterDsi_cmcrFieldGsw2txShift = 10
	RegisterDsi_cmcrFieldGsw2txMask  = 0x400
)

// GetGsw2tx GSW2TX
func (r *registerDsi_cmcrType) GetGsw2tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldGsw2txMask) != 0
}

// SetGsw2tx GSW2TX
func (r *registerDsi_cmcrType) SetGsw2tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldGsw2txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldGsw2txMask)
	}
}

const (
	RegisterDsi_cmcrFieldGsr0txShift = 11
	RegisterDsi_cmcrFieldGsr0txMask  = 0x800
)

// GetGsr0tx GSR0TX
func (r *registerDsi_cmcrType) GetGsr0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldGsr0txMask) != 0
}

// SetGsr0tx GSR0TX
func (r *registerDsi_cmcrType) SetGsr0tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldGsr0txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldGsr0txMask)
	}
}

const (
	RegisterDsi_cmcrFieldGsr1txShift = 12
	RegisterDsi_cmcrFieldGsr1txMask  = 0x1000
)

// GetGsr1tx GSR1TX
func (r *registerDsi_cmcrType) GetGsr1tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldGsr1txMask) != 0
}

// SetGsr1tx GSR1TX
func (r *registerDsi_cmcrType) SetGsr1tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldGsr1txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldGsr1txMask)
	}
}

const (
	RegisterDsi_cmcrFieldGsr2txShift = 13
	RegisterDsi_cmcrFieldGsr2txMask  = 0x2000
)

// GetGsr2tx GSR2TX
func (r *registerDsi_cmcrType) GetGsr2tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldGsr2txMask) != 0
}

// SetGsr2tx GSR2TX
func (r *registerDsi_cmcrType) SetGsr2tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldGsr2txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldGsr2txMask)
	}
}

const (
	RegisterDsi_cmcrFieldGlwtxShift = 14
	RegisterDsi_cmcrFieldGlwtxMask  = 0x4000
)

// GetGlwtx GLWTX
func (r *registerDsi_cmcrType) GetGlwtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldGlwtxMask) != 0
}

// SetGlwtx GLWTX
func (r *registerDsi_cmcrType) SetGlwtx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldGlwtxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldGlwtxMask)
	}
}

const (
	RegisterDsi_cmcrFieldDsw0txShift = 16
	RegisterDsi_cmcrFieldDsw0txMask  = 0x10000
)

// GetDsw0tx DSW0TX
func (r *registerDsi_cmcrType) GetDsw0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldDsw0txMask) != 0
}

// SetDsw0tx DSW0TX
func (r *registerDsi_cmcrType) SetDsw0tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldDsw0txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldDsw0txMask)
	}
}

const (
	RegisterDsi_cmcrFieldDsw1txShift = 17
	RegisterDsi_cmcrFieldDsw1txMask  = 0x20000
)

// GetDsw1tx DSW1TX
func (r *registerDsi_cmcrType) GetDsw1tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldDsw1txMask) != 0
}

// SetDsw1tx DSW1TX
func (r *registerDsi_cmcrType) SetDsw1tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldDsw1txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldDsw1txMask)
	}
}

const (
	RegisterDsi_cmcrFieldDsr0txShift = 18
	RegisterDsi_cmcrFieldDsr0txMask  = 0x40000
)

// GetDsr0tx DSR0TX
func (r *registerDsi_cmcrType) GetDsr0tx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldDsr0txMask) != 0
}

// SetDsr0tx DSR0TX
func (r *registerDsi_cmcrType) SetDsr0tx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldDsr0txMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldDsr0txMask)
	}
}

const (
	RegisterDsi_cmcrFieldDlwtxShift = 19
	RegisterDsi_cmcrFieldDlwtxMask  = 0x80000
)

// GetDlwtx DLWTX
func (r *registerDsi_cmcrType) GetDlwtx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldDlwtxMask) != 0
}

// SetDlwtx DLWTX
func (r *registerDsi_cmcrType) SetDlwtx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldDlwtxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldDlwtxMask)
	}
}

const (
	RegisterDsi_cmcrFieldMrdpsShift = 24
	RegisterDsi_cmcrFieldMrdpsMask  = 0x1000000
)

// GetMrdps MRDPS
func (r *registerDsi_cmcrType) GetMrdps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cmcrFieldMrdpsMask) != 0
}

// SetMrdps MRDPS
func (r *registerDsi_cmcrType) SetMrdps(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_cmcrFieldMrdpsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cmcrFieldMrdpsMask)
	}
}

// registerDsi_ghcrType DSI Host generic header configuration register
type registerDsi_ghcrType uint32

const (
	RegisterDsi_ghcrFieldDtShift = 0
	RegisterDsi_ghcrFieldDtMask  = 0x3f
)

// GetDt DT
func (r *registerDsi_ghcrType) GetDt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ghcrFieldDtMask) >> RegisterDsi_ghcrFieldDtShift)
}

// SetDt DT
func (r *registerDsi_ghcrType) SetDt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ghcrFieldDtMask)|(uint32(value)<<RegisterDsi_ghcrFieldDtShift))
}

const (
	RegisterDsi_ghcrFieldVcidShift = 6
	RegisterDsi_ghcrFieldVcidMask  = 0xc0
)

// GetVcid VCID
func (r *registerDsi_ghcrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ghcrFieldVcidMask) >> RegisterDsi_ghcrFieldVcidShift)
}

// SetVcid VCID
func (r *registerDsi_ghcrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ghcrFieldVcidMask)|(uint32(value)<<RegisterDsi_ghcrFieldVcidShift))
}

const (
	RegisterDsi_ghcrFieldWclsbShift = 8
	RegisterDsi_ghcrFieldWclsbMask  = 0xff00
)

// GetWclsb WCLSB
func (r *registerDsi_ghcrType) GetWclsb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ghcrFieldWclsbMask) >> RegisterDsi_ghcrFieldWclsbShift)
}

// SetWclsb WCLSB
func (r *registerDsi_ghcrType) SetWclsb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ghcrFieldWclsbMask)|(uint32(value)<<RegisterDsi_ghcrFieldWclsbShift))
}

const (
	RegisterDsi_ghcrFieldWcmsbShift = 16
	RegisterDsi_ghcrFieldWcmsbMask  = 0xff0000
)

// GetWcmsb WCMSB
func (r *registerDsi_ghcrType) GetWcmsb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ghcrFieldWcmsbMask) >> RegisterDsi_ghcrFieldWcmsbShift)
}

// SetWcmsb WCMSB
func (r *registerDsi_ghcrType) SetWcmsb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ghcrFieldWcmsbMask)|(uint32(value)<<RegisterDsi_ghcrFieldWcmsbShift))
}

// registerDsi_gpdrType DSI Host generic payload data register
type registerDsi_gpdrType uint32

const (
	RegisterDsi_gpdrFieldData1Shift = 0
	RegisterDsi_gpdrFieldData1Mask  = 0xff
)

// GetData1 DATA1
func (r *registerDsi_gpdrType) GetData1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gpdrFieldData1Mask) >> RegisterDsi_gpdrFieldData1Shift)
}

// SetData1 DATA1
func (r *registerDsi_gpdrType) SetData1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gpdrFieldData1Mask)|(uint32(value)<<RegisterDsi_gpdrFieldData1Shift))
}

const (
	RegisterDsi_gpdrFieldData2Shift = 8
	RegisterDsi_gpdrFieldData2Mask  = 0xff00
)

// GetData2 DATA2
func (r *registerDsi_gpdrType) GetData2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gpdrFieldData2Mask) >> RegisterDsi_gpdrFieldData2Shift)
}

// SetData2 DATA2
func (r *registerDsi_gpdrType) SetData2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gpdrFieldData2Mask)|(uint32(value)<<RegisterDsi_gpdrFieldData2Shift))
}

const (
	RegisterDsi_gpdrFieldData3Shift = 16
	RegisterDsi_gpdrFieldData3Mask  = 0xff0000
)

// GetData3 DATA3
func (r *registerDsi_gpdrType) GetData3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gpdrFieldData3Mask) >> RegisterDsi_gpdrFieldData3Shift)
}

// SetData3 DATA3
func (r *registerDsi_gpdrType) SetData3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gpdrFieldData3Mask)|(uint32(value)<<RegisterDsi_gpdrFieldData3Shift))
}

const (
	RegisterDsi_gpdrFieldData4Shift = 24
	RegisterDsi_gpdrFieldData4Mask  = 0xff000000
)

// GetData4 DATA4
func (r *registerDsi_gpdrType) GetData4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gpdrFieldData4Mask) >> RegisterDsi_gpdrFieldData4Shift)
}

// SetData4 DATA4
func (r *registerDsi_gpdrType) SetData4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gpdrFieldData4Mask)|(uint32(value)<<RegisterDsi_gpdrFieldData4Shift))
}

// registerDsi_gpsrType DSI Host generic packet status register
type registerDsi_gpsrType uint32

const (
	RegisterDsi_gpsrFieldCmdfeShift = 0
	RegisterDsi_gpsrFieldCmdfeMask  = 0x1
)

// GetCmdfe CMDFE
func (r *registerDsi_gpsrType) GetCmdfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gpsrFieldCmdfeMask) != 0
}

// SetCmdfe CMDFE
func (r *registerDsi_gpsrType) SetCmdfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_gpsrFieldCmdfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gpsrFieldCmdfeMask)
	}
}

const (
	RegisterDsi_gpsrFieldCmdffShift = 1
	RegisterDsi_gpsrFieldCmdffMask  = 0x2
)

// GetCmdff CMDFF
func (r *registerDsi_gpsrType) GetCmdff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gpsrFieldCmdffMask) != 0
}

// SetCmdff CMDFF
func (r *registerDsi_gpsrType) SetCmdff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_gpsrFieldCmdffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gpsrFieldCmdffMask)
	}
}

const (
	RegisterDsi_gpsrFieldPwrfeShift = 2
	RegisterDsi_gpsrFieldPwrfeMask  = 0x4
)

// GetPwrfe PWRFE
func (r *registerDsi_gpsrType) GetPwrfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gpsrFieldPwrfeMask) != 0
}

// SetPwrfe PWRFE
func (r *registerDsi_gpsrType) SetPwrfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_gpsrFieldPwrfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gpsrFieldPwrfeMask)
	}
}

const (
	RegisterDsi_gpsrFieldPwrffShift = 3
	RegisterDsi_gpsrFieldPwrffMask  = 0x8
)

// GetPwrff PWRFF
func (r *registerDsi_gpsrType) GetPwrff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gpsrFieldPwrffMask) != 0
}

// SetPwrff PWRFF
func (r *registerDsi_gpsrType) SetPwrff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_gpsrFieldPwrffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gpsrFieldPwrffMask)
	}
}

const (
	RegisterDsi_gpsrFieldPrdfeShift = 4
	RegisterDsi_gpsrFieldPrdfeMask  = 0x10
)

// GetPrdfe PRDFE
func (r *registerDsi_gpsrType) GetPrdfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gpsrFieldPrdfeMask) != 0
}

// SetPrdfe PRDFE
func (r *registerDsi_gpsrType) SetPrdfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_gpsrFieldPrdfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gpsrFieldPrdfeMask)
	}
}

const (
	RegisterDsi_gpsrFieldPrdffShift = 5
	RegisterDsi_gpsrFieldPrdffMask  = 0x20
)

// GetPrdff PRDFF
func (r *registerDsi_gpsrType) GetPrdff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gpsrFieldPrdffMask) != 0
}

// SetPrdff PRDFF
func (r *registerDsi_gpsrType) SetPrdff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_gpsrFieldPrdffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gpsrFieldPrdffMask)
	}
}

const (
	RegisterDsi_gpsrFieldRcbShift = 6
	RegisterDsi_gpsrFieldRcbMask  = 0x40
)

// GetRcb RCB
func (r *registerDsi_gpsrType) GetRcb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_gpsrFieldRcbMask) != 0
}

// SetRcb RCB
func (r *registerDsi_gpsrType) SetRcb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_gpsrFieldRcbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_gpsrFieldRcbMask)
	}
}

// registerDsi_tccr0Type DSI Host timeout counter configuration register 0
type registerDsi_tccr0Type uint32

const (
	RegisterDsi_tccr0FieldLprx_tocntShift = 0
	RegisterDsi_tccr0FieldLprx_tocntMask  = 0xffff
)

// GetLprx_tocnt LPRX_TOCNT
func (r *registerDsi_tccr0Type) GetLprx_tocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_tccr0FieldLprx_tocntMask) >> RegisterDsi_tccr0FieldLprx_tocntShift)
}

// SetLprx_tocnt LPRX_TOCNT
func (r *registerDsi_tccr0Type) SetLprx_tocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_tccr0FieldLprx_tocntMask)|(uint32(value)<<RegisterDsi_tccr0FieldLprx_tocntShift))
}

const (
	RegisterDsi_tccr0FieldHstx_tocntShift = 16
	RegisterDsi_tccr0FieldHstx_tocntMask  = 0xffff0000
)

// GetHstx_tocnt HSTX_TOCNT
func (r *registerDsi_tccr0Type) GetHstx_tocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_tccr0FieldHstx_tocntMask) >> RegisterDsi_tccr0FieldHstx_tocntShift)
}

// SetHstx_tocnt HSTX_TOCNT
func (r *registerDsi_tccr0Type) SetHstx_tocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_tccr0FieldHstx_tocntMask)|(uint32(value)<<RegisterDsi_tccr0FieldHstx_tocntShift))
}

// registerDsi_tccr1Type DSI Host timeout counter configuration register 1
type registerDsi_tccr1Type uint32

const (
	RegisterDsi_tccr1FieldHsrd_tocntShift = 0
	RegisterDsi_tccr1FieldHsrd_tocntMask  = 0xffff
)

// GetHsrd_tocnt HSRD_TOCNT
func (r *registerDsi_tccr1Type) GetHsrd_tocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_tccr1FieldHsrd_tocntMask) >> RegisterDsi_tccr1FieldHsrd_tocntShift)
}

// SetHsrd_tocnt HSRD_TOCNT
func (r *registerDsi_tccr1Type) SetHsrd_tocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_tccr1FieldHsrd_tocntMask)|(uint32(value)<<RegisterDsi_tccr1FieldHsrd_tocntShift))
}

// registerDsi_tccr2Type DSI Host timeout counter configuration register 2
type registerDsi_tccr2Type uint32

const (
	RegisterDsi_tccr2FieldLprd_tocntShift = 0
	RegisterDsi_tccr2FieldLprd_tocntMask  = 0xffff
)

// GetLprd_tocnt LPRD_TOCNT
func (r *registerDsi_tccr2Type) GetLprd_tocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_tccr2FieldLprd_tocntMask) >> RegisterDsi_tccr2FieldLprd_tocntShift)
}

// SetLprd_tocnt LPRD_TOCNT
func (r *registerDsi_tccr2Type) SetLprd_tocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_tccr2FieldLprd_tocntMask)|(uint32(value)<<RegisterDsi_tccr2FieldLprd_tocntShift))
}

// registerDsi_tccr3Type DSI Host timeout counter configuration register 3
type registerDsi_tccr3Type uint32

const (
	RegisterDsi_tccr3FieldHswr_tocntShift = 0
	RegisterDsi_tccr3FieldHswr_tocntMask  = 0xffff
)

// GetHswr_tocnt HSWR_TOCNT
func (r *registerDsi_tccr3Type) GetHswr_tocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_tccr3FieldHswr_tocntMask) >> RegisterDsi_tccr3FieldHswr_tocntShift)
}

// SetHswr_tocnt HSWR_TOCNT
func (r *registerDsi_tccr3Type) SetHswr_tocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_tccr3FieldHswr_tocntMask)|(uint32(value)<<RegisterDsi_tccr3FieldHswr_tocntShift))
}

const (
	RegisterDsi_tccr3FieldPmShift = 24
	RegisterDsi_tccr3FieldPmMask  = 0x1000000
)

// GetPm PM
func (r *registerDsi_tccr3Type) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_tccr3FieldPmMask) != 0
}

// SetPm PM
func (r *registerDsi_tccr3Type) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_tccr3FieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_tccr3FieldPmMask)
	}
}

// registerDsi_tccr4Type DSI Host timeout counter configuration register 4
type registerDsi_tccr4Type uint32

const (
	RegisterDsi_tccr4FieldLpwr_tocntShift = 0
	RegisterDsi_tccr4FieldLpwr_tocntMask  = 0xffff
)

// GetLpwr_tocnt LPWR_TOCNT
func (r *registerDsi_tccr4Type) GetLpwr_tocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_tccr4FieldLpwr_tocntMask) >> RegisterDsi_tccr4FieldLpwr_tocntShift)
}

// SetLpwr_tocnt LPWR_TOCNT
func (r *registerDsi_tccr4Type) SetLpwr_tocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_tccr4FieldLpwr_tocntMask)|(uint32(value)<<RegisterDsi_tccr4FieldLpwr_tocntShift))
}

// registerDsi_tccr5Type DSI Host timeout counter configuration register 5
type registerDsi_tccr5Type uint32

const (
	RegisterDsi_tccr5FieldBta_tocntShift = 0
	RegisterDsi_tccr5FieldBta_tocntMask  = 0xffff
)

// GetBta_tocnt BTA_TOCNT
func (r *registerDsi_tccr5Type) GetBta_tocnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_tccr5FieldBta_tocntMask) >> RegisterDsi_tccr5FieldBta_tocntShift)
}

// SetBta_tocnt BTA_TOCNT
func (r *registerDsi_tccr5Type) SetBta_tocnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_tccr5FieldBta_tocntMask)|(uint32(value)<<RegisterDsi_tccr5FieldBta_tocntShift))
}

// registerDsi_clcrType DSI Host clock lane configuration register
type registerDsi_clcrType uint32

const (
	RegisterDsi_clcrFieldDpccShift = 0
	RegisterDsi_clcrFieldDpccMask  = 0x1
)

// GetDpcc DPCC
func (r *registerDsi_clcrType) GetDpcc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_clcrFieldDpccMask) != 0
}

// SetDpcc DPCC
func (r *registerDsi_clcrType) SetDpcc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_clcrFieldDpccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_clcrFieldDpccMask)
	}
}

const (
	RegisterDsi_clcrFieldAcrShift = 1
	RegisterDsi_clcrFieldAcrMask  = 0x2
)

// GetAcr ACR
func (r *registerDsi_clcrType) GetAcr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_clcrFieldAcrMask) != 0
}

// SetAcr ACR
func (r *registerDsi_clcrType) SetAcr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_clcrFieldAcrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_clcrFieldAcrMask)
	}
}

// registerDsi_cltcrType DSI Host clock lane timer configuration register
type registerDsi_cltcrType uint32

const (
	RegisterDsi_cltcrFieldLp2hs_timeShift = 0
	RegisterDsi_cltcrFieldLp2hs_timeMask  = 0x3ff
)

// GetLp2hs_time LP2HS_TIME
func (r *registerDsi_cltcrType) GetLp2hs_time() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cltcrFieldLp2hs_timeMask) >> RegisterDsi_cltcrFieldLp2hs_timeShift)
}

// SetLp2hs_time LP2HS_TIME
func (r *registerDsi_cltcrType) SetLp2hs_time(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cltcrFieldLp2hs_timeMask)|(uint32(value)<<RegisterDsi_cltcrFieldLp2hs_timeShift))
}

const (
	RegisterDsi_cltcrFieldHs2lp_timeShift = 16
	RegisterDsi_cltcrFieldHs2lp_timeMask  = 0x3ff0000
)

// GetHs2lp_time HS2LP_TIME
func (r *registerDsi_cltcrType) GetHs2lp_time() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_cltcrFieldHs2lp_timeMask) >> RegisterDsi_cltcrFieldHs2lp_timeShift)
}

// SetHs2lp_time HS2LP_TIME
func (r *registerDsi_cltcrType) SetHs2lp_time(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_cltcrFieldHs2lp_timeMask)|(uint32(value)<<RegisterDsi_cltcrFieldHs2lp_timeShift))
}

// registerDsi_dltcrType DSI Host data lane timer configuration register
type registerDsi_dltcrType uint32

const (
	RegisterDsi_dltcrFieldMrd_timeShift = 0
	RegisterDsi_dltcrFieldMrd_timeMask  = 0x7fff
)

// GetMrd_time Maximum read time
func (r *registerDsi_dltcrType) GetMrd_time() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_dltcrFieldMrd_timeMask) >> RegisterDsi_dltcrFieldMrd_timeShift)
}

// SetMrd_time Maximum read time
func (r *registerDsi_dltcrType) SetMrd_time(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_dltcrFieldMrd_timeMask)|(uint32(value)<<RegisterDsi_dltcrFieldMrd_timeShift))
}

const (
	RegisterDsi_dltcrFieldLp2hs_timeShift = 16
	RegisterDsi_dltcrFieldLp2hs_timeMask  = 0xff0000
)

// GetLp2hs_time LP2HS_TIME
func (r *registerDsi_dltcrType) GetLp2hs_time() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_dltcrFieldLp2hs_timeMask) >> RegisterDsi_dltcrFieldLp2hs_timeShift)
}

// SetLp2hs_time LP2HS_TIME
func (r *registerDsi_dltcrType) SetLp2hs_time(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_dltcrFieldLp2hs_timeMask)|(uint32(value)<<RegisterDsi_dltcrFieldLp2hs_timeShift))
}

const (
	RegisterDsi_dltcrFieldHs2lp_timeShift = 24
	RegisterDsi_dltcrFieldHs2lp_timeMask  = 0xff000000
)

// GetHs2lp_time HS2LP_TIME
func (r *registerDsi_dltcrType) GetHs2lp_time() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_dltcrFieldHs2lp_timeMask) >> RegisterDsi_dltcrFieldHs2lp_timeShift)
}

// SetHs2lp_time HS2LP_TIME
func (r *registerDsi_dltcrType) SetHs2lp_time(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_dltcrFieldHs2lp_timeMask)|(uint32(value)<<RegisterDsi_dltcrFieldHs2lp_timeShift))
}

// registerDsi_pctlrType DSI Host PHY control register
type registerDsi_pctlrType uint32

const (
	RegisterDsi_pctlrFieldDenShift = 1
	RegisterDsi_pctlrFieldDenMask  = 0x2
)

// GetDen DEN
func (r *registerDsi_pctlrType) GetDen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pctlrFieldDenMask) != 0
}

// SetDen DEN
func (r *registerDsi_pctlrType) SetDen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_pctlrFieldDenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pctlrFieldDenMask)
	}
}

const (
	RegisterDsi_pctlrFieldCkeShift = 2
	RegisterDsi_pctlrFieldCkeMask  = 0x4
)

// GetCke CKE
func (r *registerDsi_pctlrType) GetCke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pctlrFieldCkeMask) != 0
}

// SetCke CKE
func (r *registerDsi_pctlrType) SetCke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_pctlrFieldCkeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pctlrFieldCkeMask)
	}
}

// registerDsi_pconfrType DSI Host PHY configuration register
type registerDsi_pconfrType uint32

const (
	RegisterDsi_pconfrFieldNlShift = 0
	RegisterDsi_pconfrFieldNlMask  = 0x3
)

// GetNl NL
func (r *registerDsi_pconfrType) GetNl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pconfrFieldNlMask) >> RegisterDsi_pconfrFieldNlShift)
}

// SetNl NL
func (r *registerDsi_pconfrType) SetNl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pconfrFieldNlMask)|(uint32(value)<<RegisterDsi_pconfrFieldNlShift))
}

const (
	RegisterDsi_pconfrFieldSw_timeShift = 8
	RegisterDsi_pconfrFieldSw_timeMask  = 0xff00
)

// GetSw_time SW_TIME
func (r *registerDsi_pconfrType) GetSw_time() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pconfrFieldSw_timeMask) >> RegisterDsi_pconfrFieldSw_timeShift)
}

// SetSw_time SW_TIME
func (r *registerDsi_pconfrType) SetSw_time(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pconfrFieldSw_timeMask)|(uint32(value)<<RegisterDsi_pconfrFieldSw_timeShift))
}

// registerDsi_pucrType DSI Host PHY ULPS control register
type registerDsi_pucrType uint32

const (
	RegisterDsi_pucrFieldUrclShift = 0
	RegisterDsi_pucrFieldUrclMask  = 0x1
)

// GetUrcl URCL
func (r *registerDsi_pucrType) GetUrcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pucrFieldUrclMask) != 0
}

// SetUrcl URCL
func (r *registerDsi_pucrType) SetUrcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_pucrFieldUrclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pucrFieldUrclMask)
	}
}

const (
	RegisterDsi_pucrFieldUeclShift = 1
	RegisterDsi_pucrFieldUeclMask  = 0x2
)

// GetUecl UECL
func (r *registerDsi_pucrType) GetUecl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pucrFieldUeclMask) != 0
}

// SetUecl UECL
func (r *registerDsi_pucrType) SetUecl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_pucrFieldUeclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pucrFieldUeclMask)
	}
}

const (
	RegisterDsi_pucrFieldUrdlShift = 2
	RegisterDsi_pucrFieldUrdlMask  = 0x4
)

// GetUrdl URDL
func (r *registerDsi_pucrType) GetUrdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pucrFieldUrdlMask) != 0
}

// SetUrdl URDL
func (r *registerDsi_pucrType) SetUrdl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_pucrFieldUrdlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pucrFieldUrdlMask)
	}
}

const (
	RegisterDsi_pucrFieldUedlShift = 3
	RegisterDsi_pucrFieldUedlMask  = 0x8
)

// GetUedl UEDL
func (r *registerDsi_pucrType) GetUedl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pucrFieldUedlMask) != 0
}

// SetUedl UEDL
func (r *registerDsi_pucrType) SetUedl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_pucrFieldUedlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pucrFieldUedlMask)
	}
}

// registerDsi_pttcrType DSI Host PHY TX triggers configuration register
type registerDsi_pttcrType uint32

const (
	RegisterDsi_pttcrFieldTx_trigShift = 0
	RegisterDsi_pttcrFieldTx_trigMask  = 0xf
)

// GetTx_trig TX_TRIG
func (r *registerDsi_pttcrType) GetTx_trig() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_pttcrFieldTx_trigMask) >> RegisterDsi_pttcrFieldTx_trigShift)
}

// SetTx_trig TX_TRIG
func (r *registerDsi_pttcrType) SetTx_trig(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_pttcrFieldTx_trigMask)|(uint32(value)<<RegisterDsi_pttcrFieldTx_trigShift))
}

// registerDsi_psrType DSI Host PHY status register
type registerDsi_psrType uint32

const (
	RegisterDsi_psrFieldPdShift = 1
	RegisterDsi_psrFieldPdMask  = 0x2
)

// GetPd PD
func (r *registerDsi_psrType) GetPd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_psrFieldPdMask) != 0
}

// SetPd PD
func (r *registerDsi_psrType) SetPd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_psrFieldPdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_psrFieldPdMask)
	}
}

const (
	RegisterDsi_psrFieldPsscShift = 2
	RegisterDsi_psrFieldPsscMask  = 0x4
)

// GetPssc PSSC
func (r *registerDsi_psrType) GetPssc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_psrFieldPsscMask) != 0
}

// SetPssc PSSC
func (r *registerDsi_psrType) SetPssc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_psrFieldPsscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_psrFieldPsscMask)
	}
}

const (
	RegisterDsi_psrFieldUancShift = 3
	RegisterDsi_psrFieldUancMask  = 0x8
)

// GetUanc UANC
func (r *registerDsi_psrType) GetUanc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_psrFieldUancMask) != 0
}

// SetUanc UANC
func (r *registerDsi_psrType) SetUanc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_psrFieldUancMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_psrFieldUancMask)
	}
}

const (
	RegisterDsi_psrFieldPss0Shift = 4
	RegisterDsi_psrFieldPss0Mask  = 0x10
)

// GetPss0 PSS0
func (r *registerDsi_psrType) GetPss0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_psrFieldPss0Mask) != 0
}

// SetPss0 PSS0
func (r *registerDsi_psrType) SetPss0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_psrFieldPss0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_psrFieldPss0Mask)
	}
}

const (
	RegisterDsi_psrFieldUan0Shift = 5
	RegisterDsi_psrFieldUan0Mask  = 0x20
)

// GetUan0 UAN0
func (r *registerDsi_psrType) GetUan0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_psrFieldUan0Mask) != 0
}

// SetUan0 UAN0
func (r *registerDsi_psrType) SetUan0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_psrFieldUan0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_psrFieldUan0Mask)
	}
}

const (
	RegisterDsi_psrFieldRue0Shift = 6
	RegisterDsi_psrFieldRue0Mask  = 0x40
)

// GetRue0 RUE0
func (r *registerDsi_psrType) GetRue0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_psrFieldRue0Mask) != 0
}

// SetRue0 RUE0
func (r *registerDsi_psrType) SetRue0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_psrFieldRue0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_psrFieldRue0Mask)
	}
}

const (
	RegisterDsi_psrFieldPss1Shift = 7
	RegisterDsi_psrFieldPss1Mask  = 0x80
)

// GetPss1 PSS1
func (r *registerDsi_psrType) GetPss1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_psrFieldPss1Mask) != 0
}

// SetPss1 PSS1
func (r *registerDsi_psrType) SetPss1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_psrFieldPss1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_psrFieldPss1Mask)
	}
}

const (
	RegisterDsi_psrFieldUan1Shift = 8
	RegisterDsi_psrFieldUan1Mask  = 0x100
)

// GetUan1 UAN1
func (r *registerDsi_psrType) GetUan1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_psrFieldUan1Mask) != 0
}

// SetUan1 UAN1
func (r *registerDsi_psrType) SetUan1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_psrFieldUan1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_psrFieldUan1Mask)
	}
}

// registerDsi_isr0Type DSI Host interrupt and status register 0
type registerDsi_isr0Type uint32

const (
	RegisterDsi_isr0FieldAe0Shift = 0
	RegisterDsi_isr0FieldAe0Mask  = 0x1
)

// GetAe0 AE0
func (r *registerDsi_isr0Type) GetAe0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe0Mask) != 0
}

// SetAe0 AE0
func (r *registerDsi_isr0Type) SetAe0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe0Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe1Shift = 1
	RegisterDsi_isr0FieldAe1Mask  = 0x2
)

// GetAe1 AE1
func (r *registerDsi_isr0Type) GetAe1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe1Mask) != 0
}

// SetAe1 AE1
func (r *registerDsi_isr0Type) SetAe1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe1Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe2Shift = 2
	RegisterDsi_isr0FieldAe2Mask  = 0x4
)

// GetAe2 AE2
func (r *registerDsi_isr0Type) GetAe2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe2Mask) != 0
}

// SetAe2 AE2
func (r *registerDsi_isr0Type) SetAe2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe2Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe3Shift = 3
	RegisterDsi_isr0FieldAe3Mask  = 0x8
)

// GetAe3 AE3
func (r *registerDsi_isr0Type) GetAe3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe3Mask) != 0
}

// SetAe3 AE3
func (r *registerDsi_isr0Type) SetAe3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe3Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe4Shift = 4
	RegisterDsi_isr0FieldAe4Mask  = 0x10
)

// GetAe4 AE4
func (r *registerDsi_isr0Type) GetAe4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe4Mask) != 0
}

// SetAe4 AE4
func (r *registerDsi_isr0Type) SetAe4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe4Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe5Shift = 5
	RegisterDsi_isr0FieldAe5Mask  = 0x20
)

// GetAe5 AE5
func (r *registerDsi_isr0Type) GetAe5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe5Mask) != 0
}

// SetAe5 AE5
func (r *registerDsi_isr0Type) SetAe5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe5Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe6Shift = 6
	RegisterDsi_isr0FieldAe6Mask  = 0x40
)

// GetAe6 AE6
func (r *registerDsi_isr0Type) GetAe6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe6Mask) != 0
}

// SetAe6 AE6
func (r *registerDsi_isr0Type) SetAe6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe6Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe7Shift = 7
	RegisterDsi_isr0FieldAe7Mask  = 0x80
)

// GetAe7 AE7
func (r *registerDsi_isr0Type) GetAe7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe7Mask) != 0
}

// SetAe7 AE7
func (r *registerDsi_isr0Type) SetAe7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe7Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe8Shift = 8
	RegisterDsi_isr0FieldAe8Mask  = 0x100
)

// GetAe8 AE8
func (r *registerDsi_isr0Type) GetAe8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe8Mask) != 0
}

// SetAe8 AE8
func (r *registerDsi_isr0Type) SetAe8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe8Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe9Shift = 9
	RegisterDsi_isr0FieldAe9Mask  = 0x200
)

// GetAe9 AE9
func (r *registerDsi_isr0Type) GetAe9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe9Mask) != 0
}

// SetAe9 AE9
func (r *registerDsi_isr0Type) SetAe9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe9Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe10Shift = 10
	RegisterDsi_isr0FieldAe10Mask  = 0x400
)

// GetAe10 AE10
func (r *registerDsi_isr0Type) GetAe10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe10Mask) != 0
}

// SetAe10 AE10
func (r *registerDsi_isr0Type) SetAe10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe10Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe11Shift = 11
	RegisterDsi_isr0FieldAe11Mask  = 0x800
)

// GetAe11 AE11
func (r *registerDsi_isr0Type) GetAe11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe11Mask) != 0
}

// SetAe11 AE11
func (r *registerDsi_isr0Type) SetAe11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe11Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe12Shift = 12
	RegisterDsi_isr0FieldAe12Mask  = 0x1000
)

// GetAe12 AE12
func (r *registerDsi_isr0Type) GetAe12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe12Mask) != 0
}

// SetAe12 AE12
func (r *registerDsi_isr0Type) SetAe12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe12Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe13Shift = 13
	RegisterDsi_isr0FieldAe13Mask  = 0x2000
)

// GetAe13 AE13
func (r *registerDsi_isr0Type) GetAe13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe13Mask) != 0
}

// SetAe13 AE13
func (r *registerDsi_isr0Type) SetAe13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe13Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe14Shift = 14
	RegisterDsi_isr0FieldAe14Mask  = 0x4000
)

// GetAe14 AE14
func (r *registerDsi_isr0Type) GetAe14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe14Mask) != 0
}

// SetAe14 AE14
func (r *registerDsi_isr0Type) SetAe14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe14Mask)
	}
}

const (
	RegisterDsi_isr0FieldAe15Shift = 15
	RegisterDsi_isr0FieldAe15Mask  = 0x8000
)

// GetAe15 AE15
func (r *registerDsi_isr0Type) GetAe15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldAe15Mask) != 0
}

// SetAe15 AE15
func (r *registerDsi_isr0Type) SetAe15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldAe15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldAe15Mask)
	}
}

const (
	RegisterDsi_isr0FieldPe0Shift = 16
	RegisterDsi_isr0FieldPe0Mask  = 0x10000
)

// GetPe0 PE0
func (r *registerDsi_isr0Type) GetPe0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldPe0Mask) != 0
}

// SetPe0 PE0
func (r *registerDsi_isr0Type) SetPe0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldPe0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldPe0Mask)
	}
}

const (
	RegisterDsi_isr0FieldPe1Shift = 17
	RegisterDsi_isr0FieldPe1Mask  = 0x20000
)

// GetPe1 PE1
func (r *registerDsi_isr0Type) GetPe1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldPe1Mask) != 0
}

// SetPe1 PE1
func (r *registerDsi_isr0Type) SetPe1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldPe1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldPe1Mask)
	}
}

const (
	RegisterDsi_isr0FieldPe2Shift = 18
	RegisterDsi_isr0FieldPe2Mask  = 0x40000
)

// GetPe2 PE2
func (r *registerDsi_isr0Type) GetPe2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldPe2Mask) != 0
}

// SetPe2 PE2
func (r *registerDsi_isr0Type) SetPe2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldPe2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldPe2Mask)
	}
}

const (
	RegisterDsi_isr0FieldPe3Shift = 19
	RegisterDsi_isr0FieldPe3Mask  = 0x80000
)

// GetPe3 PE3
func (r *registerDsi_isr0Type) GetPe3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldPe3Mask) != 0
}

// SetPe3 PE3
func (r *registerDsi_isr0Type) SetPe3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldPe3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldPe3Mask)
	}
}

const (
	RegisterDsi_isr0FieldPe4Shift = 20
	RegisterDsi_isr0FieldPe4Mask  = 0x100000
)

// GetPe4 PE4
func (r *registerDsi_isr0Type) GetPe4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr0FieldPe4Mask) != 0
}

// SetPe4 PE4
func (r *registerDsi_isr0Type) SetPe4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr0FieldPe4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr0FieldPe4Mask)
	}
}

// registerDsi_isr1Type DSI Host interrupt and status register 1
type registerDsi_isr1Type uint32

const (
	RegisterDsi_isr1FieldTohstxShift = 0
	RegisterDsi_isr1FieldTohstxMask  = 0x1
)

// GetTohstx TOHSTX
func (r *registerDsi_isr1Type) GetTohstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldTohstxMask) != 0
}

// SetTohstx TOHSTX
func (r *registerDsi_isr1Type) SetTohstx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldTohstxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldTohstxMask)
	}
}

const (
	RegisterDsi_isr1FieldTolprxShift = 1
	RegisterDsi_isr1FieldTolprxMask  = 0x2
)

// GetTolprx TOLPRX
func (r *registerDsi_isr1Type) GetTolprx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldTolprxMask) != 0
}

// SetTolprx TOLPRX
func (r *registerDsi_isr1Type) SetTolprx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldTolprxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldTolprxMask)
	}
}

const (
	RegisterDsi_isr1FieldEccseShift = 2
	RegisterDsi_isr1FieldEccseMask  = 0x4
)

// GetEccse ECCSE
func (r *registerDsi_isr1Type) GetEccse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldEccseMask) != 0
}

// SetEccse ECCSE
func (r *registerDsi_isr1Type) SetEccse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldEccseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldEccseMask)
	}
}

const (
	RegisterDsi_isr1FieldEccmeShift = 3
	RegisterDsi_isr1FieldEccmeMask  = 0x8
)

// GetEccme ECCME
func (r *registerDsi_isr1Type) GetEccme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldEccmeMask) != 0
}

// SetEccme ECCME
func (r *registerDsi_isr1Type) SetEccme(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldEccmeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldEccmeMask)
	}
}

const (
	RegisterDsi_isr1FieldCrceShift = 4
	RegisterDsi_isr1FieldCrceMask  = 0x10
)

// GetCrce CRCE
func (r *registerDsi_isr1Type) GetCrce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldCrceMask) != 0
}

// SetCrce CRCE
func (r *registerDsi_isr1Type) SetCrce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldCrceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldCrceMask)
	}
}

const (
	RegisterDsi_isr1FieldPseShift = 5
	RegisterDsi_isr1FieldPseMask  = 0x20
)

// GetPse PSE
func (r *registerDsi_isr1Type) GetPse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldPseMask) != 0
}

// SetPse PSE
func (r *registerDsi_isr1Type) SetPse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldPseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldPseMask)
	}
}

const (
	RegisterDsi_isr1FieldEotpeShift = 6
	RegisterDsi_isr1FieldEotpeMask  = 0x40
)

// GetEotpe EOTPE
func (r *registerDsi_isr1Type) GetEotpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldEotpeMask) != 0
}

// SetEotpe EOTPE
func (r *registerDsi_isr1Type) SetEotpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldEotpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldEotpeMask)
	}
}

const (
	RegisterDsi_isr1FieldLpwreShift = 7
	RegisterDsi_isr1FieldLpwreMask  = 0x80
)

// GetLpwre LPWRE
func (r *registerDsi_isr1Type) GetLpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldLpwreMask) != 0
}

// SetLpwre LPWRE
func (r *registerDsi_isr1Type) SetLpwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldLpwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldLpwreMask)
	}
}

const (
	RegisterDsi_isr1FieldGcwreShift = 8
	RegisterDsi_isr1FieldGcwreMask  = 0x100
)

// GetGcwre GCWRE
func (r *registerDsi_isr1Type) GetGcwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldGcwreMask) != 0
}

// SetGcwre GCWRE
func (r *registerDsi_isr1Type) SetGcwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldGcwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldGcwreMask)
	}
}

const (
	RegisterDsi_isr1FieldGpwreShift = 9
	RegisterDsi_isr1FieldGpwreMask  = 0x200
)

// GetGpwre GPWRE
func (r *registerDsi_isr1Type) GetGpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldGpwreMask) != 0
}

// SetGpwre GPWRE
func (r *registerDsi_isr1Type) SetGpwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldGpwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldGpwreMask)
	}
}

const (
	RegisterDsi_isr1FieldGptxeShift = 10
	RegisterDsi_isr1FieldGptxeMask  = 0x400
)

// GetGptxe GPTXE
func (r *registerDsi_isr1Type) GetGptxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldGptxeMask) != 0
}

// SetGptxe GPTXE
func (r *registerDsi_isr1Type) SetGptxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldGptxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldGptxeMask)
	}
}

const (
	RegisterDsi_isr1FieldGprdeShift = 11
	RegisterDsi_isr1FieldGprdeMask  = 0x800
)

// GetGprde GPRDE
func (r *registerDsi_isr1Type) GetGprde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldGprdeMask) != 0
}

// SetGprde GPRDE
func (r *registerDsi_isr1Type) SetGprde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldGprdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldGprdeMask)
	}
}

const (
	RegisterDsi_isr1FieldGprxeShift = 12
	RegisterDsi_isr1FieldGprxeMask  = 0x1000
)

// GetGprxe GPRXE
func (r *registerDsi_isr1Type) GetGprxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_isr1FieldGprxeMask) != 0
}

// SetGprxe GPRXE
func (r *registerDsi_isr1Type) SetGprxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_isr1FieldGprxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_isr1FieldGprxeMask)
	}
}

// registerDsi_ier0Type DSI Host interrupt enable register 0
type registerDsi_ier0Type uint32

const (
	RegisterDsi_ier0FieldAe0ieShift = 0
	RegisterDsi_ier0FieldAe0ieMask  = 0x1
)

// GetAe0ie AE0IE
func (r *registerDsi_ier0Type) GetAe0ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe0ieMask) != 0
}

// SetAe0ie AE0IE
func (r *registerDsi_ier0Type) SetAe0ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe0ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe0ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe1ieShift = 1
	RegisterDsi_ier0FieldAe1ieMask  = 0x2
)

// GetAe1ie AE1IE
func (r *registerDsi_ier0Type) GetAe1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe1ieMask) != 0
}

// SetAe1ie AE1IE
func (r *registerDsi_ier0Type) SetAe1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe1ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe2ieShift = 2
	RegisterDsi_ier0FieldAe2ieMask  = 0x4
)

// GetAe2ie AE2IE
func (r *registerDsi_ier0Type) GetAe2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe2ieMask) != 0
}

// SetAe2ie AE2IE
func (r *registerDsi_ier0Type) SetAe2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe2ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe3ieShift = 3
	RegisterDsi_ier0FieldAe3ieMask  = 0x8
)

// GetAe3ie AE3IE
func (r *registerDsi_ier0Type) GetAe3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe3ieMask) != 0
}

// SetAe3ie AE3IE
func (r *registerDsi_ier0Type) SetAe3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe3ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe4ieShift = 4
	RegisterDsi_ier0FieldAe4ieMask  = 0x10
)

// GetAe4ie AE4IE
func (r *registerDsi_ier0Type) GetAe4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe4ieMask) != 0
}

// SetAe4ie AE4IE
func (r *registerDsi_ier0Type) SetAe4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe4ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe5ieShift = 5
	RegisterDsi_ier0FieldAe5ieMask  = 0x20
)

// GetAe5ie AE5IE
func (r *registerDsi_ier0Type) GetAe5ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe5ieMask) != 0
}

// SetAe5ie AE5IE
func (r *registerDsi_ier0Type) SetAe5ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe5ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe5ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe6ieShift = 6
	RegisterDsi_ier0FieldAe6ieMask  = 0x40
)

// GetAe6ie AE6IE
func (r *registerDsi_ier0Type) GetAe6ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe6ieMask) != 0
}

// SetAe6ie AE6IE
func (r *registerDsi_ier0Type) SetAe6ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe6ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe6ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe7ieShift = 7
	RegisterDsi_ier0FieldAe7ieMask  = 0x80
)

// GetAe7ie AE7IE
func (r *registerDsi_ier0Type) GetAe7ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe7ieMask) != 0
}

// SetAe7ie AE7IE
func (r *registerDsi_ier0Type) SetAe7ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe7ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe7ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe8ieShift = 8
	RegisterDsi_ier0FieldAe8ieMask  = 0x100
)

// GetAe8ie AE8IE
func (r *registerDsi_ier0Type) GetAe8ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe8ieMask) != 0
}

// SetAe8ie AE8IE
func (r *registerDsi_ier0Type) SetAe8ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe8ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe8ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe9ieShift = 9
	RegisterDsi_ier0FieldAe9ieMask  = 0x200
)

// GetAe9ie AE9IE
func (r *registerDsi_ier0Type) GetAe9ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe9ieMask) != 0
}

// SetAe9ie AE9IE
func (r *registerDsi_ier0Type) SetAe9ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe9ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe9ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe10ieShift = 10
	RegisterDsi_ier0FieldAe10ieMask  = 0x400
)

// GetAe10ie AE10IE
func (r *registerDsi_ier0Type) GetAe10ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe10ieMask) != 0
}

// SetAe10ie AE10IE
func (r *registerDsi_ier0Type) SetAe10ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe10ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe10ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe11ieShift = 11
	RegisterDsi_ier0FieldAe11ieMask  = 0x800
)

// GetAe11ie AE11IE
func (r *registerDsi_ier0Type) GetAe11ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe11ieMask) != 0
}

// SetAe11ie AE11IE
func (r *registerDsi_ier0Type) SetAe11ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe11ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe11ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe12ieShift = 12
	RegisterDsi_ier0FieldAe12ieMask  = 0x1000
)

// GetAe12ie AE12IE
func (r *registerDsi_ier0Type) GetAe12ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe12ieMask) != 0
}

// SetAe12ie AE12IE
func (r *registerDsi_ier0Type) SetAe12ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe12ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe12ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe13ieShift = 13
	RegisterDsi_ier0FieldAe13ieMask  = 0x2000
)

// GetAe13ie AE13IE
func (r *registerDsi_ier0Type) GetAe13ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe13ieMask) != 0
}

// SetAe13ie AE13IE
func (r *registerDsi_ier0Type) SetAe13ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe13ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe13ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe14ieShift = 14
	RegisterDsi_ier0FieldAe14ieMask  = 0x4000
)

// GetAe14ie AE14IE
func (r *registerDsi_ier0Type) GetAe14ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe14ieMask) != 0
}

// SetAe14ie AE14IE
func (r *registerDsi_ier0Type) SetAe14ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe14ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe14ieMask)
	}
}

const (
	RegisterDsi_ier0FieldAe15ieShift = 15
	RegisterDsi_ier0FieldAe15ieMask  = 0x8000
)

// GetAe15ie AE15IE
func (r *registerDsi_ier0Type) GetAe15ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldAe15ieMask) != 0
}

// SetAe15ie AE15IE
func (r *registerDsi_ier0Type) SetAe15ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldAe15ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldAe15ieMask)
	}
}

const (
	RegisterDsi_ier0FieldPe0ieShift = 16
	RegisterDsi_ier0FieldPe0ieMask  = 0x10000
)

// GetPe0ie PE0IE
func (r *registerDsi_ier0Type) GetPe0ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldPe0ieMask) != 0
}

// SetPe0ie PE0IE
func (r *registerDsi_ier0Type) SetPe0ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldPe0ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldPe0ieMask)
	}
}

const (
	RegisterDsi_ier0FieldPe1ieShift = 17
	RegisterDsi_ier0FieldPe1ieMask  = 0x20000
)

// GetPe1ie PE1IE
func (r *registerDsi_ier0Type) GetPe1ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldPe1ieMask) != 0
}

// SetPe1ie PE1IE
func (r *registerDsi_ier0Type) SetPe1ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldPe1ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldPe1ieMask)
	}
}

const (
	RegisterDsi_ier0FieldPe2ieShift = 18
	RegisterDsi_ier0FieldPe2ieMask  = 0x40000
)

// GetPe2ie PE2IE
func (r *registerDsi_ier0Type) GetPe2ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldPe2ieMask) != 0
}

// SetPe2ie PE2IE
func (r *registerDsi_ier0Type) SetPe2ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldPe2ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldPe2ieMask)
	}
}

const (
	RegisterDsi_ier0FieldPe3ieShift = 19
	RegisterDsi_ier0FieldPe3ieMask  = 0x80000
)

// GetPe3ie PE3IE
func (r *registerDsi_ier0Type) GetPe3ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldPe3ieMask) != 0
}

// SetPe3ie PE3IE
func (r *registerDsi_ier0Type) SetPe3ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldPe3ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldPe3ieMask)
	}
}

const (
	RegisterDsi_ier0FieldPe4ieShift = 20
	RegisterDsi_ier0FieldPe4ieMask  = 0x100000
)

// GetPe4ie PE4IE
func (r *registerDsi_ier0Type) GetPe4ie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier0FieldPe4ieMask) != 0
}

// SetPe4ie PE4IE
func (r *registerDsi_ier0Type) SetPe4ie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier0FieldPe4ieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier0FieldPe4ieMask)
	}
}

// registerDsi_ier1Type DSI Host interrupt enable register 1
type registerDsi_ier1Type uint32

const (
	RegisterDsi_ier1FieldTohstxieShift = 0
	RegisterDsi_ier1FieldTohstxieMask  = 0x1
)

// GetTohstxie TOHSTXIE
func (r *registerDsi_ier1Type) GetTohstxie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldTohstxieMask) != 0
}

// SetTohstxie TOHSTXIE
func (r *registerDsi_ier1Type) SetTohstxie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldTohstxieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldTohstxieMask)
	}
}

const (
	RegisterDsi_ier1FieldTolprxieShift = 1
	RegisterDsi_ier1FieldTolprxieMask  = 0x2
)

// GetTolprxie TOLPRXIE
func (r *registerDsi_ier1Type) GetTolprxie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldTolprxieMask) != 0
}

// SetTolprxie TOLPRXIE
func (r *registerDsi_ier1Type) SetTolprxie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldTolprxieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldTolprxieMask)
	}
}

const (
	RegisterDsi_ier1FieldEccseieShift = 2
	RegisterDsi_ier1FieldEccseieMask  = 0x4
)

// GetEccseie ECCSEIE
func (r *registerDsi_ier1Type) GetEccseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldEccseieMask) != 0
}

// SetEccseie ECCSEIE
func (r *registerDsi_ier1Type) SetEccseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldEccseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldEccseieMask)
	}
}

const (
	RegisterDsi_ier1FieldEccmeieShift = 3
	RegisterDsi_ier1FieldEccmeieMask  = 0x8
)

// GetEccmeie ECCMEIE
func (r *registerDsi_ier1Type) GetEccmeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldEccmeieMask) != 0
}

// SetEccmeie ECCMEIE
func (r *registerDsi_ier1Type) SetEccmeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldEccmeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldEccmeieMask)
	}
}

const (
	RegisterDsi_ier1FieldCrceieShift = 4
	RegisterDsi_ier1FieldCrceieMask  = 0x10
)

// GetCrceie CRCEIE
func (r *registerDsi_ier1Type) GetCrceie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldCrceieMask) != 0
}

// SetCrceie CRCEIE
func (r *registerDsi_ier1Type) SetCrceie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldCrceieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldCrceieMask)
	}
}

const (
	RegisterDsi_ier1FieldPseieShift = 5
	RegisterDsi_ier1FieldPseieMask  = 0x20
)

// GetPseie PSEIE
func (r *registerDsi_ier1Type) GetPseie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldPseieMask) != 0
}

// SetPseie PSEIE
func (r *registerDsi_ier1Type) SetPseie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldPseieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldPseieMask)
	}
}

const (
	RegisterDsi_ier1FieldEotpeieShift = 6
	RegisterDsi_ier1FieldEotpeieMask  = 0x40
)

// GetEotpeie EOTPEIE
func (r *registerDsi_ier1Type) GetEotpeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldEotpeieMask) != 0
}

// SetEotpeie EOTPEIE
func (r *registerDsi_ier1Type) SetEotpeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldEotpeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldEotpeieMask)
	}
}

const (
	RegisterDsi_ier1FieldLpwreieShift = 7
	RegisterDsi_ier1FieldLpwreieMask  = 0x80
)

// GetLpwreie LPWREIE
func (r *registerDsi_ier1Type) GetLpwreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldLpwreieMask) != 0
}

// SetLpwreie LPWREIE
func (r *registerDsi_ier1Type) SetLpwreie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldLpwreieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldLpwreieMask)
	}
}

const (
	RegisterDsi_ier1FieldGcwreieShift = 8
	RegisterDsi_ier1FieldGcwreieMask  = 0x100
)

// GetGcwreie GCWREIE
func (r *registerDsi_ier1Type) GetGcwreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldGcwreieMask) != 0
}

// SetGcwreie GCWREIE
func (r *registerDsi_ier1Type) SetGcwreie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldGcwreieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldGcwreieMask)
	}
}

const (
	RegisterDsi_ier1FieldGpwreieShift = 9
	RegisterDsi_ier1FieldGpwreieMask  = 0x200
)

// GetGpwreie GPWREIE
func (r *registerDsi_ier1Type) GetGpwreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldGpwreieMask) != 0
}

// SetGpwreie GPWREIE
func (r *registerDsi_ier1Type) SetGpwreie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldGpwreieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldGpwreieMask)
	}
}

const (
	RegisterDsi_ier1FieldGptxeieShift = 10
	RegisterDsi_ier1FieldGptxeieMask  = 0x400
)

// GetGptxeie GPTXEIE
func (r *registerDsi_ier1Type) GetGptxeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldGptxeieMask) != 0
}

// SetGptxeie GPTXEIE
func (r *registerDsi_ier1Type) SetGptxeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldGptxeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldGptxeieMask)
	}
}

const (
	RegisterDsi_ier1FieldGprdeieShift = 11
	RegisterDsi_ier1FieldGprdeieMask  = 0x800
)

// GetGprdeie GPRDEIE
func (r *registerDsi_ier1Type) GetGprdeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldGprdeieMask) != 0
}

// SetGprdeie GPRDEIE
func (r *registerDsi_ier1Type) SetGprdeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldGprdeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldGprdeieMask)
	}
}

const (
	RegisterDsi_ier1FieldGprxeieShift = 12
	RegisterDsi_ier1FieldGprxeieMask  = 0x1000
)

// GetGprxeie GPRXEIE
func (r *registerDsi_ier1Type) GetGprxeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_ier1FieldGprxeieMask) != 0
}

// SetGprxeie GPRXEIE
func (r *registerDsi_ier1Type) SetGprxeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_ier1FieldGprxeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_ier1FieldGprxeieMask)
	}
}

// registerDsi_fir0Type DSI Host force interrupt register 0
type registerDsi_fir0Type uint32

const (
	RegisterDsi_fir0FieldFae0Shift = 0
	RegisterDsi_fir0FieldFae0Mask  = 0x1
)

// GetFae0 FAE0
func (r *registerDsi_fir0Type) GetFae0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae0Mask) != 0
}

// SetFae0 FAE0
func (r *registerDsi_fir0Type) SetFae0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae0Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae1Shift = 1
	RegisterDsi_fir0FieldFae1Mask  = 0x2
)

// GetFae1 FAE1
func (r *registerDsi_fir0Type) GetFae1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae1Mask) != 0
}

// SetFae1 FAE1
func (r *registerDsi_fir0Type) SetFae1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae1Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae2Shift = 2
	RegisterDsi_fir0FieldFae2Mask  = 0x4
)

// GetFae2 FAE2
func (r *registerDsi_fir0Type) GetFae2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae2Mask) != 0
}

// SetFae2 FAE2
func (r *registerDsi_fir0Type) SetFae2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae2Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae3Shift = 3
	RegisterDsi_fir0FieldFae3Mask  = 0x8
)

// GetFae3 FAE3
func (r *registerDsi_fir0Type) GetFae3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae3Mask) != 0
}

// SetFae3 FAE3
func (r *registerDsi_fir0Type) SetFae3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae3Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae4Shift = 4
	RegisterDsi_fir0FieldFae4Mask  = 0x10
)

// GetFae4 FAE4
func (r *registerDsi_fir0Type) GetFae4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae4Mask) != 0
}

// SetFae4 FAE4
func (r *registerDsi_fir0Type) SetFae4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae4Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae5Shift = 5
	RegisterDsi_fir0FieldFae5Mask  = 0x20
)

// GetFae5 FAE5
func (r *registerDsi_fir0Type) GetFae5() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae5Mask) != 0
}

// SetFae5 FAE5
func (r *registerDsi_fir0Type) SetFae5(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae5Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae5Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae6Shift = 6
	RegisterDsi_fir0FieldFae6Mask  = 0x40
)

// GetFae6 FAE6
func (r *registerDsi_fir0Type) GetFae6() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae6Mask) != 0
}

// SetFae6 FAE6
func (r *registerDsi_fir0Type) SetFae6(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae6Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae6Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae7Shift = 7
	RegisterDsi_fir0FieldFae7Mask  = 0x80
)

// GetFae7 FAE7
func (r *registerDsi_fir0Type) GetFae7() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae7Mask) != 0
}

// SetFae7 FAE7
func (r *registerDsi_fir0Type) SetFae7(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae7Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae7Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae8Shift = 8
	RegisterDsi_fir0FieldFae8Mask  = 0x100
)

// GetFae8 FAE8
func (r *registerDsi_fir0Type) GetFae8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae8Mask) != 0
}

// SetFae8 FAE8
func (r *registerDsi_fir0Type) SetFae8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae8Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae9Shift = 9
	RegisterDsi_fir0FieldFae9Mask  = 0x200
)

// GetFae9 FAE9
func (r *registerDsi_fir0Type) GetFae9() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae9Mask) != 0
}

// SetFae9 FAE9
func (r *registerDsi_fir0Type) SetFae9(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae9Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae9Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae10Shift = 10
	RegisterDsi_fir0FieldFae10Mask  = 0x400
)

// GetFae10 FAE10
func (r *registerDsi_fir0Type) GetFae10() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae10Mask) != 0
}

// SetFae10 FAE10
func (r *registerDsi_fir0Type) SetFae10(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae10Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae10Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae11Shift = 11
	RegisterDsi_fir0FieldFae11Mask  = 0x800
)

// GetFae11 FAE11
func (r *registerDsi_fir0Type) GetFae11() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae11Mask) != 0
}

// SetFae11 FAE11
func (r *registerDsi_fir0Type) SetFae11(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae11Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae11Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae12Shift = 12
	RegisterDsi_fir0FieldFae12Mask  = 0x1000
)

// GetFae12 FAE12
func (r *registerDsi_fir0Type) GetFae12() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae12Mask) != 0
}

// SetFae12 FAE12
func (r *registerDsi_fir0Type) SetFae12(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae12Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae12Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae13Shift = 13
	RegisterDsi_fir0FieldFae13Mask  = 0x2000
)

// GetFae13 FAE13
func (r *registerDsi_fir0Type) GetFae13() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae13Mask) != 0
}

// SetFae13 FAE13
func (r *registerDsi_fir0Type) SetFae13(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae13Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae13Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae14Shift = 14
	RegisterDsi_fir0FieldFae14Mask  = 0x4000
)

// GetFae14 FAE14
func (r *registerDsi_fir0Type) GetFae14() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae14Mask) != 0
}

// SetFae14 FAE14
func (r *registerDsi_fir0Type) SetFae14(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae14Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae14Mask)
	}
}

const (
	RegisterDsi_fir0FieldFae15Shift = 15
	RegisterDsi_fir0FieldFae15Mask  = 0x8000
)

// GetFae15 FAE15
func (r *registerDsi_fir0Type) GetFae15() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFae15Mask) != 0
}

// SetFae15 FAE15
func (r *registerDsi_fir0Type) SetFae15(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFae15Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFae15Mask)
	}
}

const (
	RegisterDsi_fir0FieldFpe0Shift = 16
	RegisterDsi_fir0FieldFpe0Mask  = 0x10000
)

// GetFpe0 FPE0
func (r *registerDsi_fir0Type) GetFpe0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFpe0Mask) != 0
}

// SetFpe0 FPE0
func (r *registerDsi_fir0Type) SetFpe0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFpe0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFpe0Mask)
	}
}

const (
	RegisterDsi_fir0FieldFpe1Shift = 17
	RegisterDsi_fir0FieldFpe1Mask  = 0x20000
)

// GetFpe1 FPE1
func (r *registerDsi_fir0Type) GetFpe1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFpe1Mask) != 0
}

// SetFpe1 FPE1
func (r *registerDsi_fir0Type) SetFpe1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFpe1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFpe1Mask)
	}
}

const (
	RegisterDsi_fir0FieldFpe2Shift = 18
	RegisterDsi_fir0FieldFpe2Mask  = 0x40000
)

// GetFpe2 FPE2
func (r *registerDsi_fir0Type) GetFpe2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFpe2Mask) != 0
}

// SetFpe2 FPE2
func (r *registerDsi_fir0Type) SetFpe2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFpe2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFpe2Mask)
	}
}

const (
	RegisterDsi_fir0FieldFpe3Shift = 19
	RegisterDsi_fir0FieldFpe3Mask  = 0x80000
)

// GetFpe3 FPE3
func (r *registerDsi_fir0Type) GetFpe3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFpe3Mask) != 0
}

// SetFpe3 FPE3
func (r *registerDsi_fir0Type) SetFpe3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFpe3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFpe3Mask)
	}
}

const (
	RegisterDsi_fir0FieldFpe4Shift = 20
	RegisterDsi_fir0FieldFpe4Mask  = 0x100000
)

// GetFpe4 FPE4
func (r *registerDsi_fir0Type) GetFpe4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir0FieldFpe4Mask) != 0
}

// SetFpe4 FPE4
func (r *registerDsi_fir0Type) SetFpe4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir0FieldFpe4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir0FieldFpe4Mask)
	}
}

// registerDsi_fir1Type DSI Host force interrupt register 1
type registerDsi_fir1Type uint32

const (
	RegisterDsi_fir1FieldFtohstxShift = 0
	RegisterDsi_fir1FieldFtohstxMask  = 0x1
)

// GetFtohstx FTOHSTX
func (r *registerDsi_fir1Type) GetFtohstx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFtohstxMask) != 0
}

// SetFtohstx FTOHSTX
func (r *registerDsi_fir1Type) SetFtohstx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFtohstxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFtohstxMask)
	}
}

const (
	RegisterDsi_fir1FieldFtolprxShift = 1
	RegisterDsi_fir1FieldFtolprxMask  = 0x2
)

// GetFtolprx FTOLPRX
func (r *registerDsi_fir1Type) GetFtolprx() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFtolprxMask) != 0
}

// SetFtolprx FTOLPRX
func (r *registerDsi_fir1Type) SetFtolprx(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFtolprxMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFtolprxMask)
	}
}

const (
	RegisterDsi_fir1FieldFeccseShift = 2
	RegisterDsi_fir1FieldFeccseMask  = 0x4
)

// GetFeccse FECCSE
func (r *registerDsi_fir1Type) GetFeccse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFeccseMask) != 0
}

// SetFeccse FECCSE
func (r *registerDsi_fir1Type) SetFeccse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFeccseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFeccseMask)
	}
}

const (
	RegisterDsi_fir1FieldFeccmeShift = 3
	RegisterDsi_fir1FieldFeccmeMask  = 0x8
)

// GetFeccme FECCME
func (r *registerDsi_fir1Type) GetFeccme() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFeccmeMask) != 0
}

// SetFeccme FECCME
func (r *registerDsi_fir1Type) SetFeccme(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFeccmeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFeccmeMask)
	}
}

const (
	RegisterDsi_fir1FieldFcrceShift = 4
	RegisterDsi_fir1FieldFcrceMask  = 0x10
)

// GetFcrce FCRCE
func (r *registerDsi_fir1Type) GetFcrce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFcrceMask) != 0
}

// SetFcrce FCRCE
func (r *registerDsi_fir1Type) SetFcrce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFcrceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFcrceMask)
	}
}

const (
	RegisterDsi_fir1FieldFpseShift = 5
	RegisterDsi_fir1FieldFpseMask  = 0x20
)

// GetFpse FPSE
func (r *registerDsi_fir1Type) GetFpse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFpseMask) != 0
}

// SetFpse FPSE
func (r *registerDsi_fir1Type) SetFpse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFpseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFpseMask)
	}
}

const (
	RegisterDsi_fir1FieldFeotpeShift = 6
	RegisterDsi_fir1FieldFeotpeMask  = 0x40
)

// GetFeotpe FEOTPE
func (r *registerDsi_fir1Type) GetFeotpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFeotpeMask) != 0
}

// SetFeotpe FEOTPE
func (r *registerDsi_fir1Type) SetFeotpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFeotpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFeotpeMask)
	}
}

const (
	RegisterDsi_fir1FieldFlpwreShift = 7
	RegisterDsi_fir1FieldFlpwreMask  = 0x80
)

// GetFlpwre FLPWRE
func (r *registerDsi_fir1Type) GetFlpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFlpwreMask) != 0
}

// SetFlpwre FLPWRE
func (r *registerDsi_fir1Type) SetFlpwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFlpwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFlpwreMask)
	}
}

const (
	RegisterDsi_fir1FieldFgcwreShift = 8
	RegisterDsi_fir1FieldFgcwreMask  = 0x100
)

// GetFgcwre FGCWRE
func (r *registerDsi_fir1Type) GetFgcwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFgcwreMask) != 0
}

// SetFgcwre FGCWRE
func (r *registerDsi_fir1Type) SetFgcwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFgcwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFgcwreMask)
	}
}

const (
	RegisterDsi_fir1FieldFgpwreShift = 9
	RegisterDsi_fir1FieldFgpwreMask  = 0x200
)

// GetFgpwre FGPWRE
func (r *registerDsi_fir1Type) GetFgpwre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFgpwreMask) != 0
}

// SetFgpwre FGPWRE
func (r *registerDsi_fir1Type) SetFgpwre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFgpwreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFgpwreMask)
	}
}

const (
	RegisterDsi_fir1FieldFgptxeShift = 10
	RegisterDsi_fir1FieldFgptxeMask  = 0x400
)

// GetFgptxe FGPTXE
func (r *registerDsi_fir1Type) GetFgptxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFgptxeMask) != 0
}

// SetFgptxe FGPTXE
func (r *registerDsi_fir1Type) SetFgptxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFgptxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFgptxeMask)
	}
}

const (
	RegisterDsi_fir1FieldFgprdeShift = 11
	RegisterDsi_fir1FieldFgprdeMask  = 0x800
)

// GetFgprde FGPRDE
func (r *registerDsi_fir1Type) GetFgprde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFgprdeMask) != 0
}

// SetFgprde FGPRDE
func (r *registerDsi_fir1Type) SetFgprde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFgprdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFgprdeMask)
	}
}

const (
	RegisterDsi_fir1FieldFgprxeShift = 12
	RegisterDsi_fir1FieldFgprxeMask  = 0x1000
)

// GetFgprxe FGPRXE
func (r *registerDsi_fir1Type) GetFgprxe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_fir1FieldFgprxeMask) != 0
}

// SetFgprxe FGPRXE
func (r *registerDsi_fir1Type) SetFgprxe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_fir1FieldFgprxeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_fir1FieldFgprxeMask)
	}
}

// registerDsi_vscrType DSI Host video shadow control register
type registerDsi_vscrType uint32

const (
	RegisterDsi_vscrFieldEnShift = 0
	RegisterDsi_vscrFieldEnMask  = 0x1
)

// GetEn EN
func (r *registerDsi_vscrType) GetEn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vscrFieldEnMask) != 0
}

// SetEn EN
func (r *registerDsi_vscrType) SetEn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vscrFieldEnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vscrFieldEnMask)
	}
}

const (
	RegisterDsi_vscrFieldUrShift = 8
	RegisterDsi_vscrFieldUrMask  = 0x100
)

// GetUr UR
func (r *registerDsi_vscrType) GetUr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vscrFieldUrMask) != 0
}

// SetUr UR
func (r *registerDsi_vscrType) SetUr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vscrFieldUrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vscrFieldUrMask)
	}
}

// registerDsi_lcvcidrType DSI Host LTDC current VCID register
type registerDsi_lcvcidrType uint32

const (
	RegisterDsi_lcvcidrFieldVcidShift = 0
	RegisterDsi_lcvcidrFieldVcidMask  = 0x3
)

// GetVcid VCID
func (r *registerDsi_lcvcidrType) GetVcid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lcvcidrFieldVcidMask) >> RegisterDsi_lcvcidrFieldVcidShift)
}

// SetVcid VCID
func (r *registerDsi_lcvcidrType) SetVcid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lcvcidrFieldVcidMask)|(uint32(value)<<RegisterDsi_lcvcidrFieldVcidShift))
}

// registerDsi_lcccrType DSI Host LTDC current color coding register
type registerDsi_lcccrType uint32

const (
	RegisterDsi_lcccrFieldColcShift = 0
	RegisterDsi_lcccrFieldColcMask  = 0xf
)

// GetColc COLC
func (r *registerDsi_lcccrType) GetColc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lcccrFieldColcMask) >> RegisterDsi_lcccrFieldColcShift)
}

// SetColc COLC
func (r *registerDsi_lcccrType) SetColc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lcccrFieldColcMask)|(uint32(value)<<RegisterDsi_lcccrFieldColcShift))
}

const (
	RegisterDsi_lcccrFieldLpeShift = 8
	RegisterDsi_lcccrFieldLpeMask  = 0x100
)

// GetLpe LPE
func (r *registerDsi_lcccrType) GetLpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lcccrFieldLpeMask) != 0
}

// SetLpe LPE
func (r *registerDsi_lcccrType) SetLpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_lcccrFieldLpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lcccrFieldLpeMask)
	}
}

// registerDsi_lpmccrType DSI Host low-power mode current configuration register
type registerDsi_lpmccrType uint32

const (
	RegisterDsi_lpmccrFieldVlpsizeShift = 0
	RegisterDsi_lpmccrFieldVlpsizeMask  = 0xff
)

// GetVlpsize VLPSIZE
func (r *registerDsi_lpmccrType) GetVlpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lpmccrFieldVlpsizeMask) >> RegisterDsi_lpmccrFieldVlpsizeShift)
}

// SetVlpsize VLPSIZE
func (r *registerDsi_lpmccrType) SetVlpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lpmccrFieldVlpsizeMask)|(uint32(value)<<RegisterDsi_lpmccrFieldVlpsizeShift))
}

const (
	RegisterDsi_lpmccrFieldLpsizeShift = 16
	RegisterDsi_lpmccrFieldLpsizeMask  = 0xff0000
)

// GetLpsize LPSIZE
func (r *registerDsi_lpmccrType) GetLpsize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_lpmccrFieldLpsizeMask) >> RegisterDsi_lpmccrFieldLpsizeShift)
}

// SetLpsize LPSIZE
func (r *registerDsi_lpmccrType) SetLpsize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_lpmccrFieldLpsizeMask)|(uint32(value)<<RegisterDsi_lpmccrFieldLpsizeShift))
}

// registerDsi_vmccrType DSI Host video mode current configuration register
type registerDsi_vmccrType uint32

const (
	RegisterDsi_vmccrFieldVmtShift = 0
	RegisterDsi_vmccrFieldVmtMask  = 0x3
)

// GetVmt VMT
func (r *registerDsi_vmccrType) GetVmt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmccrFieldVmtMask) >> RegisterDsi_vmccrFieldVmtShift)
}

// SetVmt VMT
func (r *registerDsi_vmccrType) SetVmt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmccrFieldVmtMask)|(uint32(value)<<RegisterDsi_vmccrFieldVmtShift))
}

const (
	RegisterDsi_vmccrFieldLpvsaeShift = 2
	RegisterDsi_vmccrFieldLpvsaeMask  = 0x4
)

// GetLpvsae LPVSAE
func (r *registerDsi_vmccrType) GetLpvsae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmccrFieldLpvsaeMask) != 0
}

// SetLpvsae LPVSAE
func (r *registerDsi_vmccrType) SetLpvsae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmccrFieldLpvsaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmccrFieldLpvsaeMask)
	}
}

const (
	RegisterDsi_vmccrFieldLpvbpeShift = 3
	RegisterDsi_vmccrFieldLpvbpeMask  = 0x8
)

// GetLpvbpe LPVBPE
func (r *registerDsi_vmccrType) GetLpvbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmccrFieldLpvbpeMask) != 0
}

// SetLpvbpe LPVBPE
func (r *registerDsi_vmccrType) SetLpvbpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmccrFieldLpvbpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmccrFieldLpvbpeMask)
	}
}

const (
	RegisterDsi_vmccrFieldLpvfpeShift = 4
	RegisterDsi_vmccrFieldLpvfpeMask  = 0x10
)

// GetLpvfpe LPVFPE
func (r *registerDsi_vmccrType) GetLpvfpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmccrFieldLpvfpeMask) != 0
}

// SetLpvfpe LPVFPE
func (r *registerDsi_vmccrType) SetLpvfpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmccrFieldLpvfpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmccrFieldLpvfpeMask)
	}
}

const (
	RegisterDsi_vmccrFieldLpvaeShift = 5
	RegisterDsi_vmccrFieldLpvaeMask  = 0x20
)

// GetLpvae LPVAE
func (r *registerDsi_vmccrType) GetLpvae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmccrFieldLpvaeMask) != 0
}

// SetLpvae LPVAE
func (r *registerDsi_vmccrType) SetLpvae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmccrFieldLpvaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmccrFieldLpvaeMask)
	}
}

const (
	RegisterDsi_vmccrFieldLphbpeShift = 6
	RegisterDsi_vmccrFieldLphbpeMask  = 0x40
)

// GetLphbpe LPHBPE
func (r *registerDsi_vmccrType) GetLphbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmccrFieldLphbpeMask) != 0
}

// SetLphbpe LPHBPE
func (r *registerDsi_vmccrType) SetLphbpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmccrFieldLphbpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmccrFieldLphbpeMask)
	}
}

const (
	RegisterDsi_vmccrFieldLphfeShift = 7
	RegisterDsi_vmccrFieldLphfeMask  = 0x80
)

// GetLphfe LPHFE
func (r *registerDsi_vmccrType) GetLphfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmccrFieldLphfeMask) != 0
}

// SetLphfe LPHFE
func (r *registerDsi_vmccrType) SetLphfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmccrFieldLphfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmccrFieldLphfeMask)
	}
}

const (
	RegisterDsi_vmccrFieldFbtaaeShift = 8
	RegisterDsi_vmccrFieldFbtaaeMask  = 0x100
)

// GetFbtaae FBTAAE
func (r *registerDsi_vmccrType) GetFbtaae() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmccrFieldFbtaaeMask) != 0
}

// SetFbtaae FBTAAE
func (r *registerDsi_vmccrType) SetFbtaae(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmccrFieldFbtaaeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmccrFieldFbtaaeMask)
	}
}

const (
	RegisterDsi_vmccrFieldLpceShift = 9
	RegisterDsi_vmccrFieldLpceMask  = 0x200
)

// GetLpce LPCE
func (r *registerDsi_vmccrType) GetLpce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vmccrFieldLpceMask) != 0
}

// SetLpce LPCE
func (r *registerDsi_vmccrType) SetLpce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_vmccrFieldLpceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vmccrFieldLpceMask)
	}
}

// registerDsi_vpccrType DSI Host video packet current configuration register
type registerDsi_vpccrType uint32

const (
	RegisterDsi_vpccrFieldVpsizeShift = 0
	RegisterDsi_vpccrFieldVpsizeMask  = 0x3fff
)

// GetVpsize VPSIZE
func (r *registerDsi_vpccrType) GetVpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vpccrFieldVpsizeMask) >> RegisterDsi_vpccrFieldVpsizeShift)
}

// SetVpsize VPSIZE
func (r *registerDsi_vpccrType) SetVpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vpccrFieldVpsizeMask)|(uint32(value)<<RegisterDsi_vpccrFieldVpsizeShift))
}

// registerDsi_vcccrType DSI Host video chunks current configuration register
type registerDsi_vcccrType uint32

const (
	RegisterDsi_vcccrFieldNumcShift = 0
	RegisterDsi_vcccrFieldNumcMask  = 0x1fff
)

// GetNumc NUMC
func (r *registerDsi_vcccrType) GetNumc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vcccrFieldNumcMask) >> RegisterDsi_vcccrFieldNumcShift)
}

// SetNumc NUMC
func (r *registerDsi_vcccrType) SetNumc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vcccrFieldNumcMask)|(uint32(value)<<RegisterDsi_vcccrFieldNumcShift))
}

// registerDsi_vnpccrType DSI Host video null packet current configuration register
type registerDsi_vnpccrType uint32

const (
	RegisterDsi_vnpccrFieldNpsizeShift = 0
	RegisterDsi_vnpccrFieldNpsizeMask  = 0x1fff
)

// GetNpsize NPSIZE
func (r *registerDsi_vnpccrType) GetNpsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vnpccrFieldNpsizeMask) >> RegisterDsi_vnpccrFieldNpsizeShift)
}

// SetNpsize NPSIZE
func (r *registerDsi_vnpccrType) SetNpsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vnpccrFieldNpsizeMask)|(uint32(value)<<RegisterDsi_vnpccrFieldNpsizeShift))
}

// registerDsi_vhsaccrType DSI Host video HSA current configuration register
type registerDsi_vhsaccrType uint32

const (
	RegisterDsi_vhsaccrFieldHsaShift = 0
	RegisterDsi_vhsaccrFieldHsaMask  = 0xfff
)

// GetHsa HSA
func (r *registerDsi_vhsaccrType) GetHsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vhsaccrFieldHsaMask) >> RegisterDsi_vhsaccrFieldHsaShift)
}

// SetHsa HSA
func (r *registerDsi_vhsaccrType) SetHsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vhsaccrFieldHsaMask)|(uint32(value)<<RegisterDsi_vhsaccrFieldHsaShift))
}

// registerDsi_vhbpccrType DSI Host video HBP current configuration register
type registerDsi_vhbpccrType uint32

const (
	RegisterDsi_vhbpccrFieldHbpShift = 0
	RegisterDsi_vhbpccrFieldHbpMask  = 0xfff
)

// GetHbp HBP
func (r *registerDsi_vhbpccrType) GetHbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vhbpccrFieldHbpMask) >> RegisterDsi_vhbpccrFieldHbpShift)
}

// SetHbp HBP
func (r *registerDsi_vhbpccrType) SetHbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vhbpccrFieldHbpMask)|(uint32(value)<<RegisterDsi_vhbpccrFieldHbpShift))
}

// registerDsi_vlccrType DSI Host video line current configuration register
type registerDsi_vlccrType uint32

const (
	RegisterDsi_vlccrFieldHlineShift = 0
	RegisterDsi_vlccrFieldHlineMask  = 0x7fff
)

// GetHline HLINE
func (r *registerDsi_vlccrType) GetHline() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vlccrFieldHlineMask) >> RegisterDsi_vlccrFieldHlineShift)
}

// SetHline HLINE
func (r *registerDsi_vlccrType) SetHline(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vlccrFieldHlineMask)|(uint32(value)<<RegisterDsi_vlccrFieldHlineShift))
}

// registerDsi_vvsaccrType DSI Host video VSA current configuration register
type registerDsi_vvsaccrType uint32

const (
	RegisterDsi_vvsaccrFieldVsaShift = 0
	RegisterDsi_vvsaccrFieldVsaMask  = 0x3ff
)

// GetVsa VSA
func (r *registerDsi_vvsaccrType) GetVsa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vvsaccrFieldVsaMask) >> RegisterDsi_vvsaccrFieldVsaShift)
}

// SetVsa VSA
func (r *registerDsi_vvsaccrType) SetVsa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vvsaccrFieldVsaMask)|(uint32(value)<<RegisterDsi_vvsaccrFieldVsaShift))
}

// registerDsi_vvbpccrType DSI Host video VBP current configuration register
type registerDsi_vvbpccrType uint32

const (
	RegisterDsi_vvbpccrFieldVbpShift = 0
	RegisterDsi_vvbpccrFieldVbpMask  = 0x3ff
)

// GetVbp VBP
func (r *registerDsi_vvbpccrType) GetVbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vvbpccrFieldVbpMask) >> RegisterDsi_vvbpccrFieldVbpShift)
}

// SetVbp VBP
func (r *registerDsi_vvbpccrType) SetVbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vvbpccrFieldVbpMask)|(uint32(value)<<RegisterDsi_vvbpccrFieldVbpShift))
}

// registerDsi_vvfpccrType DSI Host video VFP current configuration register
type registerDsi_vvfpccrType uint32

const (
	RegisterDsi_vvfpccrFieldVfpShift = 0
	RegisterDsi_vvfpccrFieldVfpMask  = 0x3ff
)

// GetVfp VFP
func (r *registerDsi_vvfpccrType) GetVfp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vvfpccrFieldVfpMask) >> RegisterDsi_vvfpccrFieldVfpShift)
}

// SetVfp VFP
func (r *registerDsi_vvfpccrType) SetVfp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vvfpccrFieldVfpMask)|(uint32(value)<<RegisterDsi_vvfpccrFieldVfpShift))
}

// registerDsi_vvaccrType DSI Host video VA current configuration register
type registerDsi_vvaccrType uint32

const (
	RegisterDsi_vvaccrFieldVaShift = 0
	RegisterDsi_vvaccrFieldVaMask  = 0x3fff
)

// GetVa VA
func (r *registerDsi_vvaccrType) GetVa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_vvaccrFieldVaMask) >> RegisterDsi_vvaccrFieldVaShift)
}

// SetVa VA
func (r *registerDsi_vvaccrType) SetVa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_vvaccrFieldVaMask)|(uint32(value)<<RegisterDsi_vvaccrFieldVaShift))
}

// registerDsi_wcfgrType DSI wrapper configuration register
type registerDsi_wcfgrType uint32

const (
	RegisterDsi_wcfgrFieldDsimShift = 0
	RegisterDsi_wcfgrFieldDsimMask  = 0x1
)

// GetDsim DSIM
func (r *registerDsi_wcfgrType) GetDsim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wcfgrFieldDsimMask) != 0
}

// SetDsim DSIM
func (r *registerDsi_wcfgrType) SetDsim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wcfgrFieldDsimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wcfgrFieldDsimMask)
	}
}

const (
	RegisterDsi_wcfgrFieldColmuxShift = 1
	RegisterDsi_wcfgrFieldColmuxMask  = 0xe
)

// GetColmux COLMUX
func (r *registerDsi_wcfgrType) GetColmux() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wcfgrFieldColmuxMask) >> RegisterDsi_wcfgrFieldColmuxShift)
}

// SetColmux COLMUX
func (r *registerDsi_wcfgrType) SetColmux(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wcfgrFieldColmuxMask)|(uint32(value)<<RegisterDsi_wcfgrFieldColmuxShift))
}

const (
	RegisterDsi_wcfgrFieldTesrcShift = 4
	RegisterDsi_wcfgrFieldTesrcMask  = 0x10
)

// GetTesrc TESRC
func (r *registerDsi_wcfgrType) GetTesrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wcfgrFieldTesrcMask) != 0
}

// SetTesrc TESRC
func (r *registerDsi_wcfgrType) SetTesrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wcfgrFieldTesrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wcfgrFieldTesrcMask)
	}
}

const (
	RegisterDsi_wcfgrFieldTepolShift = 5
	RegisterDsi_wcfgrFieldTepolMask  = 0x20
)

// GetTepol TEPOL
func (r *registerDsi_wcfgrType) GetTepol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wcfgrFieldTepolMask) != 0
}

// SetTepol TEPOL
func (r *registerDsi_wcfgrType) SetTepol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wcfgrFieldTepolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wcfgrFieldTepolMask)
	}
}

const (
	RegisterDsi_wcfgrFieldArShift = 6
	RegisterDsi_wcfgrFieldArMask  = 0x40
)

// GetAr AR
func (r *registerDsi_wcfgrType) GetAr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wcfgrFieldArMask) != 0
}

// SetAr AR
func (r *registerDsi_wcfgrType) SetAr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wcfgrFieldArMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wcfgrFieldArMask)
	}
}

const (
	RegisterDsi_wcfgrFieldVspolShift = 7
	RegisterDsi_wcfgrFieldVspolMask  = 0x80
)

// GetVspol VSPOL
func (r *registerDsi_wcfgrType) GetVspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wcfgrFieldVspolMask) != 0
}

// SetVspol VSPOL
func (r *registerDsi_wcfgrType) SetVspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wcfgrFieldVspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wcfgrFieldVspolMask)
	}
}

// registerDsi_wcrType DSI wrapper control register
type registerDsi_wcrType uint32

const (
	RegisterDsi_wcrFieldColmShift = 0
	RegisterDsi_wcrFieldColmMask  = 0x1
)

// GetColm COLM
func (r *registerDsi_wcrType) GetColm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wcrFieldColmMask) != 0
}

// SetColm COLM
func (r *registerDsi_wcrType) SetColm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wcrFieldColmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wcrFieldColmMask)
	}
}

const (
	RegisterDsi_wcrFieldShtdnShift = 1
	RegisterDsi_wcrFieldShtdnMask  = 0x2
)

// GetShtdn SHTDN
func (r *registerDsi_wcrType) GetShtdn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wcrFieldShtdnMask) != 0
}

// SetShtdn SHTDN
func (r *registerDsi_wcrType) SetShtdn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wcrFieldShtdnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wcrFieldShtdnMask)
	}
}

const (
	RegisterDsi_wcrFieldLtdcenShift = 2
	RegisterDsi_wcrFieldLtdcenMask  = 0x4
)

// GetLtdcen LTDCEN
func (r *registerDsi_wcrType) GetLtdcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wcrFieldLtdcenMask) != 0
}

// SetLtdcen LTDCEN
func (r *registerDsi_wcrType) SetLtdcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wcrFieldLtdcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wcrFieldLtdcenMask)
	}
}

const (
	RegisterDsi_wcrFieldDsienShift = 3
	RegisterDsi_wcrFieldDsienMask  = 0x8
)

// GetDsien DSIEN
func (r *registerDsi_wcrType) GetDsien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wcrFieldDsienMask) != 0
}

// SetDsien DSIEN
func (r *registerDsi_wcrType) SetDsien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wcrFieldDsienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wcrFieldDsienMask)
	}
}

// registerDsi_wierType DSI wrapper interrupt enable register
type registerDsi_wierType uint32

const (
	RegisterDsi_wierFieldTeieShift = 0
	RegisterDsi_wierFieldTeieMask  = 0x1
)

// GetTeie TEIE
func (r *registerDsi_wierType) GetTeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wierFieldTeieMask) != 0
}

// SetTeie TEIE
func (r *registerDsi_wierType) SetTeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wierFieldTeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wierFieldTeieMask)
	}
}

const (
	RegisterDsi_wierFieldErieShift = 1
	RegisterDsi_wierFieldErieMask  = 0x2
)

// GetErie ERIE
func (r *registerDsi_wierType) GetErie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wierFieldErieMask) != 0
}

// SetErie ERIE
func (r *registerDsi_wierType) SetErie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wierFieldErieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wierFieldErieMask)
	}
}

const (
	RegisterDsi_wierFieldPlllieShift = 9
	RegisterDsi_wierFieldPlllieMask  = 0x200
)

// GetPlllie PLLLIE
func (r *registerDsi_wierType) GetPlllie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wierFieldPlllieMask) != 0
}

// SetPlllie PLLLIE
func (r *registerDsi_wierType) SetPlllie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wierFieldPlllieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wierFieldPlllieMask)
	}
}

const (
	RegisterDsi_wierFieldPlluieShift = 10
	RegisterDsi_wierFieldPlluieMask  = 0x400
)

// GetPlluie PLLUIE
func (r *registerDsi_wierType) GetPlluie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wierFieldPlluieMask) != 0
}

// SetPlluie PLLUIE
func (r *registerDsi_wierType) SetPlluie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wierFieldPlluieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wierFieldPlluieMask)
	}
}

const (
	RegisterDsi_wierFieldRrieShift = 13
	RegisterDsi_wierFieldRrieMask  = 0x2000
)

// GetRrie RRIE
func (r *registerDsi_wierType) GetRrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wierFieldRrieMask) != 0
}

// SetRrie RRIE
func (r *registerDsi_wierType) SetRrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wierFieldRrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wierFieldRrieMask)
	}
}

// registerDsi_wisrType DSI wrapper interrupt and status register
type registerDsi_wisrType uint32

const (
	RegisterDsi_wisrFieldTeifShift = 0
	RegisterDsi_wisrFieldTeifMask  = 0x1
)

// GetTeif TEIF
func (r *registerDsi_wisrType) GetTeif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wisrFieldTeifMask) != 0
}

// SetTeif TEIF
func (r *registerDsi_wisrType) SetTeif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wisrFieldTeifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wisrFieldTeifMask)
	}
}

const (
	RegisterDsi_wisrFieldErifShift = 1
	RegisterDsi_wisrFieldErifMask  = 0x2
)

// GetErif ERIF
func (r *registerDsi_wisrType) GetErif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wisrFieldErifMask) != 0
}

// SetErif ERIF
func (r *registerDsi_wisrType) SetErif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wisrFieldErifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wisrFieldErifMask)
	}
}

const (
	RegisterDsi_wisrFieldBusyShift = 2
	RegisterDsi_wisrFieldBusyMask  = 0x4
)

// GetBusy BUSY
func (r *registerDsi_wisrType) GetBusy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wisrFieldBusyMask) != 0
}

// SetBusy BUSY
func (r *registerDsi_wisrType) SetBusy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wisrFieldBusyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wisrFieldBusyMask)
	}
}

const (
	RegisterDsi_wisrFieldPlllsShift = 8
	RegisterDsi_wisrFieldPlllsMask  = 0x100
)

// GetPllls PLLLS
func (r *registerDsi_wisrType) GetPllls() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wisrFieldPlllsMask) != 0
}

// SetPllls PLLLS
func (r *registerDsi_wisrType) SetPllls(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wisrFieldPlllsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wisrFieldPlllsMask)
	}
}

const (
	RegisterDsi_wisrFieldPlllifShift = 9
	RegisterDsi_wisrFieldPlllifMask  = 0x200
)

// GetPlllif PLLLIF
func (r *registerDsi_wisrType) GetPlllif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wisrFieldPlllifMask) != 0
}

// SetPlllif PLLLIF
func (r *registerDsi_wisrType) SetPlllif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wisrFieldPlllifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wisrFieldPlllifMask)
	}
}

const (
	RegisterDsi_wisrFieldPlluifShift = 10
	RegisterDsi_wisrFieldPlluifMask  = 0x400
)

// GetPlluif PLLUIF
func (r *registerDsi_wisrType) GetPlluif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wisrFieldPlluifMask) != 0
}

// SetPlluif PLLUIF
func (r *registerDsi_wisrType) SetPlluif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wisrFieldPlluifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wisrFieldPlluifMask)
	}
}

const (
	RegisterDsi_wisrFieldRrsShift = 12
	RegisterDsi_wisrFieldRrsMask  = 0x1000
)

// GetRrs RRS
func (r *registerDsi_wisrType) GetRrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wisrFieldRrsMask) != 0
}

// SetRrs RRS
func (r *registerDsi_wisrType) SetRrs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wisrFieldRrsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wisrFieldRrsMask)
	}
}

const (
	RegisterDsi_wisrFieldRrifShift = 13
	RegisterDsi_wisrFieldRrifMask  = 0x2000
)

// GetRrif RRIF
func (r *registerDsi_wisrType) GetRrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wisrFieldRrifMask) != 0
}

// SetRrif RRIF
func (r *registerDsi_wisrType) SetRrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wisrFieldRrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wisrFieldRrifMask)
	}
}

// registerDsi_wifcrType DSI wrapper interrupt flag clear register
type registerDsi_wifcrType uint32

const (
	RegisterDsi_wifcrFieldCteifShift = 0
	RegisterDsi_wifcrFieldCteifMask  = 0x1
)

// GetCteif CTEIF
func (r *registerDsi_wifcrType) GetCteif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wifcrFieldCteifMask) != 0
}

// SetCteif CTEIF
func (r *registerDsi_wifcrType) SetCteif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wifcrFieldCteifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wifcrFieldCteifMask)
	}
}

const (
	RegisterDsi_wifcrFieldCerifShift = 1
	RegisterDsi_wifcrFieldCerifMask  = 0x2
)

// GetCerif CERIF
func (r *registerDsi_wifcrType) GetCerif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wifcrFieldCerifMask) != 0
}

// SetCerif CERIF
func (r *registerDsi_wifcrType) SetCerif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wifcrFieldCerifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wifcrFieldCerifMask)
	}
}

const (
	RegisterDsi_wifcrFieldCplllifShift = 9
	RegisterDsi_wifcrFieldCplllifMask  = 0x200
)

// GetCplllif CPLLLIF
func (r *registerDsi_wifcrType) GetCplllif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wifcrFieldCplllifMask) != 0
}

// SetCplllif CPLLLIF
func (r *registerDsi_wifcrType) SetCplllif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wifcrFieldCplllifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wifcrFieldCplllifMask)
	}
}

const (
	RegisterDsi_wifcrFieldCplluifShift = 10
	RegisterDsi_wifcrFieldCplluifMask  = 0x400
)

// GetCplluif CPLLUIF
func (r *registerDsi_wifcrType) GetCplluif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wifcrFieldCplluifMask) != 0
}

// SetCplluif CPLLUIF
func (r *registerDsi_wifcrType) SetCplluif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wifcrFieldCplluifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wifcrFieldCplluifMask)
	}
}

const (
	RegisterDsi_wifcrFieldCrrifShift = 13
	RegisterDsi_wifcrFieldCrrifMask  = 0x2000
)

// GetCrrif CRRIF
func (r *registerDsi_wifcrType) GetCrrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wifcrFieldCrrifMask) != 0
}

// SetCrrif CRRIF
func (r *registerDsi_wifcrType) SetCrrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wifcrFieldCrrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wifcrFieldCrrifMask)
	}
}

// registerDsi_wpcr0Type DSI wrapper PHY configuration register 0
type registerDsi_wpcr0Type uint32

const (
	RegisterDsi_wpcr0FieldUix4Shift = 0
	RegisterDsi_wpcr0FieldUix4Mask  = 0x3f
)

// GetUix4 UIX4
func (r *registerDsi_wpcr0Type) GetUix4() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldUix4Mask) >> RegisterDsi_wpcr0FieldUix4Shift)
}

// SetUix4 UIX4
func (r *registerDsi_wpcr0Type) SetUix4(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldUix4Mask)|(uint32(value)<<RegisterDsi_wpcr0FieldUix4Shift))
}

const (
	RegisterDsi_wpcr0FieldSwclShift = 6
	RegisterDsi_wpcr0FieldSwclMask  = 0x40
)

// GetSwcl SWCL
func (r *registerDsi_wpcr0Type) GetSwcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldSwclMask) != 0
}

// SetSwcl SWCL
func (r *registerDsi_wpcr0Type) SetSwcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldSwclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldSwclMask)
	}
}

const (
	RegisterDsi_wpcr0FieldSwdl0Shift = 7
	RegisterDsi_wpcr0FieldSwdl0Mask  = 0x80
)

// GetSwdl0 SWDL0
func (r *registerDsi_wpcr0Type) GetSwdl0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldSwdl0Mask) != 0
}

// SetSwdl0 SWDL0
func (r *registerDsi_wpcr0Type) SetSwdl0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldSwdl0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldSwdl0Mask)
	}
}

const (
	RegisterDsi_wpcr0FieldSwdl1Shift = 8
	RegisterDsi_wpcr0FieldSwdl1Mask  = 0x100
)

// GetSwdl1 SWDL1
func (r *registerDsi_wpcr0Type) GetSwdl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldSwdl1Mask) != 0
}

// SetSwdl1 SWDL1
func (r *registerDsi_wpcr0Type) SetSwdl1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldSwdl1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldSwdl1Mask)
	}
}

const (
	RegisterDsi_wpcr0FieldHsiclShift = 9
	RegisterDsi_wpcr0FieldHsiclMask  = 0x200
)

// GetHsicl HSICL
func (r *registerDsi_wpcr0Type) GetHsicl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldHsiclMask) != 0
}

// SetHsicl HSICL
func (r *registerDsi_wpcr0Type) SetHsicl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldHsiclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldHsiclMask)
	}
}

const (
	RegisterDsi_wpcr0FieldHsidl0Shift = 10
	RegisterDsi_wpcr0FieldHsidl0Mask  = 0x400
)

// GetHsidl0 HSIDL0
func (r *registerDsi_wpcr0Type) GetHsidl0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldHsidl0Mask) != 0
}

// SetHsidl0 HSIDL0
func (r *registerDsi_wpcr0Type) SetHsidl0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldHsidl0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldHsidl0Mask)
	}
}

const (
	RegisterDsi_wpcr0FieldHsidl1Shift = 11
	RegisterDsi_wpcr0FieldHsidl1Mask  = 0x800
)

// GetHsidl1 HSIDL1
func (r *registerDsi_wpcr0Type) GetHsidl1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldHsidl1Mask) != 0
}

// SetHsidl1 HSIDL1
func (r *registerDsi_wpcr0Type) SetHsidl1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldHsidl1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldHsidl1Mask)
	}
}

const (
	RegisterDsi_wpcr0FieldFtxsmclShift = 12
	RegisterDsi_wpcr0FieldFtxsmclMask  = 0x1000
)

// GetFtxsmcl FTXSMCL
func (r *registerDsi_wpcr0Type) GetFtxsmcl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldFtxsmclMask) != 0
}

// SetFtxsmcl FTXSMCL
func (r *registerDsi_wpcr0Type) SetFtxsmcl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldFtxsmclMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldFtxsmclMask)
	}
}

const (
	RegisterDsi_wpcr0FieldFtxsmdlShift = 13
	RegisterDsi_wpcr0FieldFtxsmdlMask  = 0x2000
)

// GetFtxsmdl FTXSMDL
func (r *registerDsi_wpcr0Type) GetFtxsmdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldFtxsmdlMask) != 0
}

// SetFtxsmdl FTXSMDL
func (r *registerDsi_wpcr0Type) SetFtxsmdl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldFtxsmdlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldFtxsmdlMask)
	}
}

const (
	RegisterDsi_wpcr0FieldCdoffdlShift = 14
	RegisterDsi_wpcr0FieldCdoffdlMask  = 0x4000
)

// GetCdoffdl CDOFFDL
func (r *registerDsi_wpcr0Type) GetCdoffdl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldCdoffdlMask) != 0
}

// SetCdoffdl CDOFFDL
func (r *registerDsi_wpcr0Type) SetCdoffdl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldCdoffdlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldCdoffdlMask)
	}
}

const (
	RegisterDsi_wpcr0FieldTddlShift = 16
	RegisterDsi_wpcr0FieldTddlMask  = 0x10000
)

// GetTddl TDDL
func (r *registerDsi_wpcr0Type) GetTddl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldTddlMask) != 0
}

// SetTddl TDDL
func (r *registerDsi_wpcr0Type) SetTddl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldTddlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldTddlMask)
	}
}

const (
	RegisterDsi_wpcr0FieldPdenShift = 18
	RegisterDsi_wpcr0FieldPdenMask  = 0x40000
)

// GetPden Pull-down enable
func (r *registerDsi_wpcr0Type) GetPden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldPdenMask) != 0
}

// SetPden Pull-down enable
func (r *registerDsi_wpcr0Type) SetPden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldPdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldPdenMask)
	}
}

const (
	RegisterDsi_wpcr0FieldTclkprepenShift = 19
	RegisterDsi_wpcr0FieldTclkprepenMask  = 0x80000
)

// GetTclkprepen Custom time for tCLK-PREPARE enable
func (r *registerDsi_wpcr0Type) GetTclkprepen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldTclkprepenMask) != 0
}

// SetTclkprepen Custom time for tCLK-PREPARE enable
func (r *registerDsi_wpcr0Type) SetTclkprepen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldTclkprepenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldTclkprepenMask)
	}
}

const (
	RegisterDsi_wpcr0FieldTclkzeroenShift = 20
	RegisterDsi_wpcr0FieldTclkzeroenMask  = 0x100000
)

// GetTclkzeroen Custom time for tCLK-ZERO enable
func (r *registerDsi_wpcr0Type) GetTclkzeroen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldTclkzeroenMask) != 0
}

// SetTclkzeroen Custom time for tCLK-ZERO enable
func (r *registerDsi_wpcr0Type) SetTclkzeroen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldTclkzeroenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldTclkzeroenMask)
	}
}

const (
	RegisterDsi_wpcr0FieldThsprepenShift = 21
	RegisterDsi_wpcr0FieldThsprepenMask  = 0x200000
)

// GetThsprepen Custom time for tHS-PREPARE enable
func (r *registerDsi_wpcr0Type) GetThsprepen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldThsprepenMask) != 0
}

// SetThsprepen Custom time for tHS-PREPARE enable
func (r *registerDsi_wpcr0Type) SetThsprepen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldThsprepenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldThsprepenMask)
	}
}

const (
	RegisterDsi_wpcr0FieldThstrailenShift = 22
	RegisterDsi_wpcr0FieldThstrailenMask  = 0x400000
)

// GetThstrailen Custom time for tHS-TRAIL enable
func (r *registerDsi_wpcr0Type) GetThstrailen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldThstrailenMask) != 0
}

// SetThstrailen Custom time for tHS-TRAIL enable
func (r *registerDsi_wpcr0Type) SetThstrailen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldThstrailenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldThstrailenMask)
	}
}

const (
	RegisterDsi_wpcr0FieldThszeroenShift = 23
	RegisterDsi_wpcr0FieldThszeroenMask  = 0x800000
)

// GetThszeroen Custom time for tHS-ZERO enable
func (r *registerDsi_wpcr0Type) GetThszeroen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldThszeroenMask) != 0
}

// SetThszeroen Custom time for tHS-ZERO enable
func (r *registerDsi_wpcr0Type) SetThszeroen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldThszeroenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldThszeroenMask)
	}
}

const (
	RegisterDsi_wpcr0FieldTlpxdenShift = 24
	RegisterDsi_wpcr0FieldTlpxdenMask  = 0x1000000
)

// GetTlpxden Custom time for tLPX for data lanes enable
func (r *registerDsi_wpcr0Type) GetTlpxden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldTlpxdenMask) != 0
}

// SetTlpxden Custom time for tLPX for data lanes enable
func (r *registerDsi_wpcr0Type) SetTlpxden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldTlpxdenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldTlpxdenMask)
	}
}

const (
	RegisterDsi_wpcr0FieldThsexitenShift = 25
	RegisterDsi_wpcr0FieldThsexitenMask  = 0x2000000
)

// GetThsexiten Custom time for tHS-EXIT enable
func (r *registerDsi_wpcr0Type) GetThsexiten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldThsexitenMask) != 0
}

// SetThsexiten Custom time for tHS-EXIT enable
func (r *registerDsi_wpcr0Type) SetThsexiten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldThsexitenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldThsexitenMask)
	}
}

const (
	RegisterDsi_wpcr0FieldTlpxcenShift = 26
	RegisterDsi_wpcr0FieldTlpxcenMask  = 0x4000000
)

// GetTlpxcen Custom time for tLPX for clock lane enable
func (r *registerDsi_wpcr0Type) GetTlpxcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldTlpxcenMask) != 0
}

// SetTlpxcen Custom time for tLPX for clock lane enable
func (r *registerDsi_wpcr0Type) SetTlpxcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldTlpxcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldTlpxcenMask)
	}
}

const (
	RegisterDsi_wpcr0FieldTclkpostenShift = 27
	RegisterDsi_wpcr0FieldTclkpostenMask  = 0x8000000
)

// GetTclkposten Custom time for tCLK-POST enable
func (r *registerDsi_wpcr0Type) GetTclkposten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr0FieldTclkpostenMask) != 0
}

// SetTclkposten Custom time for tCLK-POST enable
func (r *registerDsi_wpcr0Type) SetTclkposten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr0FieldTclkpostenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr0FieldTclkpostenMask)
	}
}

// registerDsi_wpcr1Type This register shall be programmed only when DSI is stopped (CR. DSIEN=0 and CR.EN = 0).
type registerDsi_wpcr1Type uint32

const (
	RegisterDsi_wpcr1FieldHstxdclShift = 0
	RegisterDsi_wpcr1FieldHstxdclMask  = 0x3
)

// GetHstxdcl High-speed transmission delay on clock lane
func (r *registerDsi_wpcr1Type) GetHstxdcl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr1FieldHstxdclMask) >> RegisterDsi_wpcr1FieldHstxdclShift)
}

// SetHstxdcl High-speed transmission delay on clock lane
func (r *registerDsi_wpcr1Type) SetHstxdcl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr1FieldHstxdclMask)|(uint32(value)<<RegisterDsi_wpcr1FieldHstxdclShift))
}

const (
	RegisterDsi_wpcr1FieldHstxddlShift = 2
	RegisterDsi_wpcr1FieldHstxddlMask  = 0xc
)

// GetHstxddl High-speed transmission delay on data lanes
func (r *registerDsi_wpcr1Type) GetHstxddl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr1FieldHstxddlMask) >> RegisterDsi_wpcr1FieldHstxddlShift)
}

// SetHstxddl High-speed transmission delay on data lanes
func (r *registerDsi_wpcr1Type) SetHstxddl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr1FieldHstxddlMask)|(uint32(value)<<RegisterDsi_wpcr1FieldHstxddlShift))
}

const (
	RegisterDsi_wpcr1FieldLpsrcclShift = 6
	RegisterDsi_wpcr1FieldLpsrcclMask  = 0xc0
)

// GetLpsrccl Low-power transmission slew-rate compensation on clock lane
func (r *registerDsi_wpcr1Type) GetLpsrccl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr1FieldLpsrcclMask) >> RegisterDsi_wpcr1FieldLpsrcclShift)
}

// SetLpsrccl Low-power transmission slew-rate compensation on clock lane
func (r *registerDsi_wpcr1Type) SetLpsrccl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr1FieldLpsrcclMask)|(uint32(value)<<RegisterDsi_wpcr1FieldLpsrcclShift))
}

const (
	RegisterDsi_wpcr1FieldLpsrcdlShift = 8
	RegisterDsi_wpcr1FieldLpsrcdlMask  = 0x300
)

// GetLpsrcdl Low-power transmission slew-rate compensation on data lanes
func (r *registerDsi_wpcr1Type) GetLpsrcdl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr1FieldLpsrcdlMask) >> RegisterDsi_wpcr1FieldLpsrcdlShift)
}

// SetLpsrcdl Low-power transmission slew-rate compensation on data lanes
func (r *registerDsi_wpcr1Type) SetLpsrcdl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr1FieldLpsrcdlMask)|(uint32(value)<<RegisterDsi_wpcr1FieldLpsrcdlShift))
}

const (
	RegisterDsi_wpcr1FieldSddcShift = 12
	RegisterDsi_wpcr1FieldSddcMask  = 0x1000
)

// GetSddc SDD control
func (r *registerDsi_wpcr1Type) GetSddc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr1FieldSddcMask) != 0
}

// SetSddc SDD control
func (r *registerDsi_wpcr1Type) SetSddc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr1FieldSddcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr1FieldSddcMask)
	}
}

const (
	RegisterDsi_wpcr1FieldHstxsrcclShift = 16
	RegisterDsi_wpcr1FieldHstxsrcclMask  = 0x30000
)

// GetHstxsrccl High-speed transmission slew-rate control on clock lane
func (r *registerDsi_wpcr1Type) GetHstxsrccl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr1FieldHstxsrcclMask) >> RegisterDsi_wpcr1FieldHstxsrcclShift)
}

// SetHstxsrccl High-speed transmission slew-rate control on clock lane
func (r *registerDsi_wpcr1Type) SetHstxsrccl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr1FieldHstxsrcclMask)|(uint32(value)<<RegisterDsi_wpcr1FieldHstxsrcclShift))
}

const (
	RegisterDsi_wpcr1FieldHstxsrcdlShift = 18
	RegisterDsi_wpcr1FieldHstxsrcdlMask  = 0xc0000
)

// GetHstxsrcdl High-speed transmission slew-rate control on data lanes
func (r *registerDsi_wpcr1Type) GetHstxsrcdl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr1FieldHstxsrcdlMask) >> RegisterDsi_wpcr1FieldHstxsrcdlShift)
}

// SetHstxsrcdl High-speed transmission slew-rate control on data lanes
func (r *registerDsi_wpcr1Type) SetHstxsrcdl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr1FieldHstxsrcdlMask)|(uint32(value)<<RegisterDsi_wpcr1FieldHstxsrcdlShift))
}

const (
	RegisterDsi_wpcr1FieldFlprxlpmShift = 22
	RegisterDsi_wpcr1FieldFlprxlpmMask  = 0x400000
)

// GetFlprxlpm Forces LP receiver in low-power mode
func (r *registerDsi_wpcr1Type) GetFlprxlpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr1FieldFlprxlpmMask) != 0
}

// SetFlprxlpm Forces LP receiver in low-power mode
func (r *registerDsi_wpcr1Type) SetFlprxlpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wpcr1FieldFlprxlpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr1FieldFlprxlpmMask)
	}
}

const (
	RegisterDsi_wpcr1FieldLprxftShift = 25
	RegisterDsi_wpcr1FieldLprxftMask  = 0x6000000
)

// GetLprxft Low-power RX low-pass filtering tuning
func (r *registerDsi_wpcr1Type) GetLprxft() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr1FieldLprxftMask) >> RegisterDsi_wpcr1FieldLprxftShift)
}

// SetLprxft Low-power RX low-pass filtering tuning
func (r *registerDsi_wpcr1Type) SetLprxft(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr1FieldLprxftMask)|(uint32(value)<<RegisterDsi_wpcr1FieldLprxftShift))
}

// registerDsi_wpcr2Type DSI wrapper PHY configuration register 2
type registerDsi_wpcr2Type uint32

const (
	RegisterDsi_wpcr2FieldTclkprepShift = 0
	RegisterDsi_wpcr2FieldTclkprepMask  = 0xff
)

// GetTclkprep TCLKPREP
func (r *registerDsi_wpcr2Type) GetTclkprep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr2FieldTclkprepMask) >> RegisterDsi_wpcr2FieldTclkprepShift)
}

// SetTclkprep TCLKPREP
func (r *registerDsi_wpcr2Type) SetTclkprep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr2FieldTclkprepMask)|(uint32(value)<<RegisterDsi_wpcr2FieldTclkprepShift))
}

const (
	RegisterDsi_wpcr2FieldTclkzeroShift = 8
	RegisterDsi_wpcr2FieldTclkzeroMask  = 0xff00
)

// GetTclkzero TCLKZERO
func (r *registerDsi_wpcr2Type) GetTclkzero() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr2FieldTclkzeroMask) >> RegisterDsi_wpcr2FieldTclkzeroShift)
}

// SetTclkzero TCLKZERO
func (r *registerDsi_wpcr2Type) SetTclkzero(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr2FieldTclkzeroMask)|(uint32(value)<<RegisterDsi_wpcr2FieldTclkzeroShift))
}

const (
	RegisterDsi_wpcr2FieldThsprepShift = 16
	RegisterDsi_wpcr2FieldThsprepMask  = 0xff0000
)

// GetThsprep THSPREP
func (r *registerDsi_wpcr2Type) GetThsprep() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr2FieldThsprepMask) >> RegisterDsi_wpcr2FieldThsprepShift)
}

// SetThsprep THSPREP
func (r *registerDsi_wpcr2Type) SetThsprep(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr2FieldThsprepMask)|(uint32(value)<<RegisterDsi_wpcr2FieldThsprepShift))
}

const (
	RegisterDsi_wpcr2FieldThstrailShift = 24
	RegisterDsi_wpcr2FieldThstrailMask  = 0xff000000
)

// GetThstrail THSTRAIL
func (r *registerDsi_wpcr2Type) GetThstrail() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr2FieldThstrailMask) >> RegisterDsi_wpcr2FieldThstrailShift)
}

// SetThstrail THSTRAIL
func (r *registerDsi_wpcr2Type) SetThstrail(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr2FieldThstrailMask)|(uint32(value)<<RegisterDsi_wpcr2FieldThstrailShift))
}

// registerDsi_wpcr3Type DSI wrapper PHY configuration register 3
type registerDsi_wpcr3Type uint32

const (
	RegisterDsi_wpcr3FieldThszeroShift = 0
	RegisterDsi_wpcr3FieldThszeroMask  = 0xff
)

// GetThszero THSZERO
func (r *registerDsi_wpcr3Type) GetThszero() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr3FieldThszeroMask) >> RegisterDsi_wpcr3FieldThszeroShift)
}

// SetThszero THSZERO
func (r *registerDsi_wpcr3Type) SetThszero(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr3FieldThszeroMask)|(uint32(value)<<RegisterDsi_wpcr3FieldThszeroShift))
}

const (
	RegisterDsi_wpcr3FieldTlpxdShift = 8
	RegisterDsi_wpcr3FieldTlpxdMask  = 0xff00
)

// GetTlpxd TLPXD
func (r *registerDsi_wpcr3Type) GetTlpxd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr3FieldTlpxdMask) >> RegisterDsi_wpcr3FieldTlpxdShift)
}

// SetTlpxd TLPXD
func (r *registerDsi_wpcr3Type) SetTlpxd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr3FieldTlpxdMask)|(uint32(value)<<RegisterDsi_wpcr3FieldTlpxdShift))
}

const (
	RegisterDsi_wpcr3FieldThsexitShift = 16
	RegisterDsi_wpcr3FieldThsexitMask  = 0xff0000
)

// GetThsexit THSEXIT
func (r *registerDsi_wpcr3Type) GetThsexit() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr3FieldThsexitMask) >> RegisterDsi_wpcr3FieldThsexitShift)
}

// SetThsexit THSEXIT
func (r *registerDsi_wpcr3Type) SetThsexit(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr3FieldThsexitMask)|(uint32(value)<<RegisterDsi_wpcr3FieldThsexitShift))
}

const (
	RegisterDsi_wpcr3FieldTlpxcShift = 24
	RegisterDsi_wpcr3FieldTlpxcMask  = 0xff000000
)

// GetTlpxc TLPXC
func (r *registerDsi_wpcr3Type) GetTlpxc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr3FieldTlpxcMask) >> RegisterDsi_wpcr3FieldTlpxcShift)
}

// SetTlpxc TLPXC
func (r *registerDsi_wpcr3Type) SetTlpxc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr3FieldTlpxcMask)|(uint32(value)<<RegisterDsi_wpcr3FieldTlpxcShift))
}

// registerDsi_wpcr4Type DSI wrapper PHY configuration register 4
type registerDsi_wpcr4Type uint32

const (
	RegisterDsi_wpcr4FieldTclkpostShift = 0
	RegisterDsi_wpcr4FieldTclkpostMask  = 0xff
)

// GetTclkpost TCLKPOST
func (r *registerDsi_wpcr4Type) GetTclkpost() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wpcr4FieldTclkpostMask) >> RegisterDsi_wpcr4FieldTclkpostShift)
}

// SetTclkpost TCLKPOST
func (r *registerDsi_wpcr4Type) SetTclkpost(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wpcr4FieldTclkpostMask)|(uint32(value)<<RegisterDsi_wpcr4FieldTclkpostShift))
}

// registerDsi_wrpcrType DSI wrapper regulator and PLL control register
type registerDsi_wrpcrType uint32

const (
	RegisterDsi_wrpcrFieldPllenShift = 0
	RegisterDsi_wrpcrFieldPllenMask  = 0x1
)

// GetPllen PLLEN
func (r *registerDsi_wrpcrType) GetPllen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wrpcrFieldPllenMask) != 0
}

// SetPllen PLLEN
func (r *registerDsi_wrpcrType) SetPllen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wrpcrFieldPllenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wrpcrFieldPllenMask)
	}
}

const (
	RegisterDsi_wrpcrFieldNdivShift = 2
	RegisterDsi_wrpcrFieldNdivMask  = 0x1fc
)

// GetNdiv NDIV
func (r *registerDsi_wrpcrType) GetNdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wrpcrFieldNdivMask) >> RegisterDsi_wrpcrFieldNdivShift)
}

// SetNdiv NDIV
func (r *registerDsi_wrpcrType) SetNdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wrpcrFieldNdivMask)|(uint32(value)<<RegisterDsi_wrpcrFieldNdivShift))
}

const (
	RegisterDsi_wrpcrFieldIdfShift = 11
	RegisterDsi_wrpcrFieldIdfMask  = 0x7800
)

// GetIdf IDF
func (r *registerDsi_wrpcrType) GetIdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wrpcrFieldIdfMask) >> RegisterDsi_wrpcrFieldIdfShift)
}

// SetIdf IDF
func (r *registerDsi_wrpcrType) SetIdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wrpcrFieldIdfMask)|(uint32(value)<<RegisterDsi_wrpcrFieldIdfShift))
}

const (
	RegisterDsi_wrpcrFieldOdfShift = 16
	RegisterDsi_wrpcrFieldOdfMask  = 0x30000
)

// GetOdf ODF
func (r *registerDsi_wrpcrType) GetOdf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wrpcrFieldOdfMask) >> RegisterDsi_wrpcrFieldOdfShift)
}

// SetOdf ODF
func (r *registerDsi_wrpcrType) SetOdf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wrpcrFieldOdfMask)|(uint32(value)<<RegisterDsi_wrpcrFieldOdfShift))
}

const (
	RegisterDsi_wrpcrFieldRegenShift = 24
	RegisterDsi_wrpcrFieldRegenMask  = 0x1000000
)

// GetRegen REGEN
func (r *registerDsi_wrpcrType) GetRegen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDsi_wrpcrFieldRegenMask) != 0
}

// SetRegen REGEN
func (r *registerDsi_wrpcrType) SetRegen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDsi_wrpcrFieldRegenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDsi_wrpcrFieldRegenMask)
	}
}
