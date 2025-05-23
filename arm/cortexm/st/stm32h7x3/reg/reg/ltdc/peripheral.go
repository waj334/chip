package ltdc

import (
	"unsafe"
	"volatile"
)

var (
	Ltdc = (*_ltdc)(unsafe.Pointer(uintptr(0x50001000)))
)

type _ltdc struct {
	_        [8]byte
	Sscr     registerSscrType
	Bpcr     registerBpcrType
	Awcr     registerAwcrType
	Twcr     registerTwcrType
	Gcr      registerGcrType
	_        [8]byte
	Srcr     registerSrcrType
	_        [4]byte
	Bccr     registerBccrType
	_        [4]byte
	Ier      registerIerType
	Isr      registerIsrType
	Icr      registerIcrType
	Lipcr    registerLipcrType
	Cpsr     registerCpsrType
	Cdsr     registerCdsrType
	_        [56]byte
	L1cr     registerL1crType
	L1whpcr  registerL1whpcrType
	L1wvpcr  registerL1wvpcrType
	L1ckcr   registerL1ckcrType
	L1pfcr   registerL1pfcrType
	L1cacr   registerL1cacrType
	L1dccr   registerL1dccrType
	L1bfcr   registerL1bfcrType
	_        [8]byte
	L1cfbar  registerL1cfbarType
	L1cfblr  registerL1cfblrType
	L1cfblnr registerL1cfblnrType
	_        [12]byte
	L1clutwr registerL1clutwrType
	_        [60]byte
	L2cr     registerL2crType
	L2whpcr  registerL2whpcrType
	L2wvpcr  registerL2wvpcrType
	L2ckcr   registerL2ckcrType
	L2pfcr   registerL2pfcrType
	L2cacr   registerL2cacrType
	L2dccr   registerL2dccrType
	L2bfcr   registerL2bfcrType
	_        [8]byte
	L2cfbar  registerL2cfbarType
	L2cfblr  registerL2cfblrType
	L2cfblnr registerL2cfblnrType
	_        [12]byte
	L2clutwr registerL2clutwrType
}

// registerSscrType Synchronization Size Configuration Register
type registerSscrType uint32

const (
	RegisterSscrFieldHswShift = 16
	RegisterSscrFieldHswMask  = 0x3ff0000
)

// GetHsw Horizontal Synchronization Width (in units of pixel clock period)
func (r *registerSscrType) GetHsw() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSscrFieldHswMask) >> RegisterSscrFieldHswShift)
}

// SetHsw Horizontal Synchronization Width (in units of pixel clock period)
func (r *registerSscrType) SetHsw(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSscrFieldHswMask)|(uint32(value)<<RegisterSscrFieldHswShift))
}

const (
	RegisterSscrFieldVshShift = 0
	RegisterSscrFieldVshMask  = 0x7ff
)

// GetVsh Vertical Synchronization Height (in units of horizontal scan line)
func (r *registerSscrType) GetVsh() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSscrFieldVshMask) >> RegisterSscrFieldVshShift)
}

// SetVsh Vertical Synchronization Height (in units of horizontal scan line)
func (r *registerSscrType) SetVsh(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSscrFieldVshMask)|(uint32(value)<<RegisterSscrFieldVshShift))
}

// registerBpcrType Back Porch Configuration Register
type registerBpcrType uint32

const (
	RegisterBpcrFieldAhbpShift = 16
	RegisterBpcrFieldAhbpMask  = 0xfff0000
)

// GetAhbp Accumulated Horizontal back porch (in units of pixel clock period)
func (r *registerBpcrType) GetAhbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBpcrFieldAhbpMask) >> RegisterBpcrFieldAhbpShift)
}

// SetAhbp Accumulated Horizontal back porch (in units of pixel clock period)
func (r *registerBpcrType) SetAhbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBpcrFieldAhbpMask)|(uint32(value)<<RegisterBpcrFieldAhbpShift))
}

const (
	RegisterBpcrFieldAvbpShift = 0
	RegisterBpcrFieldAvbpMask  = 0x7ff
)

// GetAvbp Accumulated Vertical back porch (in units of horizontal scan line)
func (r *registerBpcrType) GetAvbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBpcrFieldAvbpMask) >> RegisterBpcrFieldAvbpShift)
}

// SetAvbp Accumulated Vertical back porch (in units of horizontal scan line)
func (r *registerBpcrType) SetAvbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBpcrFieldAvbpMask)|(uint32(value)<<RegisterBpcrFieldAvbpShift))
}

// registerAwcrType Active Width Configuration Register
type registerAwcrType uint32

const (
	RegisterAwcrFieldAavShift = 16
	RegisterAwcrFieldAavMask  = 0xfff0000
)

// GetAav AAV
func (r *registerAwcrType) GetAav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterAwcrFieldAavMask) >> RegisterAwcrFieldAavShift)
}

// SetAav AAV
func (r *registerAwcrType) SetAav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAwcrFieldAavMask)|(uint32(value)<<RegisterAwcrFieldAavShift))
}

const (
	RegisterAwcrFieldAahShift = 0
	RegisterAwcrFieldAahMask  = 0x7ff
)

// GetAah Accumulated Active Height (in units of horizontal scan line)
func (r *registerAwcrType) GetAah() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterAwcrFieldAahMask) >> RegisterAwcrFieldAahShift)
}

// SetAah Accumulated Active Height (in units of horizontal scan line)
func (r *registerAwcrType) SetAah(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAwcrFieldAahMask)|(uint32(value)<<RegisterAwcrFieldAahShift))
}

// registerTwcrType Total Width Configuration Register
type registerTwcrType uint32

const (
	RegisterTwcrFieldTotalwShift = 16
	RegisterTwcrFieldTotalwMask  = 0xfff0000
)

// GetTotalw Total Width (in units of pixel clock period)
func (r *registerTwcrType) GetTotalw() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterTwcrFieldTotalwMask) >> RegisterTwcrFieldTotalwShift)
}

// SetTotalw Total Width (in units of pixel clock period)
func (r *registerTwcrType) SetTotalw(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTwcrFieldTotalwMask)|(uint32(value)<<RegisterTwcrFieldTotalwShift))
}

const (
	RegisterTwcrFieldTotalhShift = 0
	RegisterTwcrFieldTotalhMask  = 0x7ff
)

// GetTotalh Total Height (in units of horizontal scan line)
func (r *registerTwcrType) GetTotalh() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterTwcrFieldTotalhMask) >> RegisterTwcrFieldTotalhShift)
}

// SetTotalh Total Height (in units of horizontal scan line)
func (r *registerTwcrType) SetTotalh(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTwcrFieldTotalhMask)|(uint32(value)<<RegisterTwcrFieldTotalhShift))
}

// registerGcrType Global Control Register
type registerGcrType uint32

const (
	RegisterGcrFieldHspolShift = 31
	RegisterGcrFieldHspolMask  = 0x80000000
)

// GetHspol Horizontal Synchronization Polarity
func (r *registerGcrType) GetHspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldHspolMask) != 0
}

// SetHspol Horizontal Synchronization Polarity
func (r *registerGcrType) SetHspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldHspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldHspolMask)
	}
}

const (
	RegisterGcrFieldVspolShift = 30
	RegisterGcrFieldVspolMask  = 0x40000000
)

// GetVspol Vertical Synchronization Polarity
func (r *registerGcrType) GetVspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldVspolMask) != 0
}

// SetVspol Vertical Synchronization Polarity
func (r *registerGcrType) SetVspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldVspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldVspolMask)
	}
}

const (
	RegisterGcrFieldDepolShift = 29
	RegisterGcrFieldDepolMask  = 0x20000000
)

// GetDepol Data Enable Polarity
func (r *registerGcrType) GetDepol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldDepolMask) != 0
}

// SetDepol Data Enable Polarity
func (r *registerGcrType) SetDepol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldDepolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldDepolMask)
	}
}

const (
	RegisterGcrFieldPcpolShift = 28
	RegisterGcrFieldPcpolMask  = 0x10000000
)

// GetPcpol Pixel Clock Polarity
func (r *registerGcrType) GetPcpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldPcpolMask) != 0
}

// SetPcpol Pixel Clock Polarity
func (r *registerGcrType) SetPcpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldPcpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldPcpolMask)
	}
}

const (
	RegisterGcrFieldDenShift = 16
	RegisterGcrFieldDenMask  = 0x10000
)

// GetDen Dither Enable
func (r *registerGcrType) GetDen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldDenMask) != 0
}

// SetDen Dither Enable
func (r *registerGcrType) SetDen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldDenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldDenMask)
	}
}

const (
	RegisterGcrFieldDrwShift = 12
	RegisterGcrFieldDrwMask  = 0x7000
)

// GetDrw Dither Red Width
func (r *registerGcrType) GetDrw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldDrwMask) >> RegisterGcrFieldDrwShift)
}

const (
	RegisterGcrFieldDgwShift = 8
	RegisterGcrFieldDgwMask  = 0x700
)

// GetDgw Dither Green Width
func (r *registerGcrType) GetDgw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldDgwMask) >> RegisterGcrFieldDgwShift)
}

const (
	RegisterGcrFieldDbwShift = 4
	RegisterGcrFieldDbwMask  = 0x70
)

// GetDbw Dither Blue Width
func (r *registerGcrType) GetDbw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldDbwMask) >> RegisterGcrFieldDbwShift)
}

const (
	RegisterGcrFieldLtdcenShift = 0
	RegisterGcrFieldLtdcenMask  = 0x1
)

// GetLtdcen LCD-TFT controller enable bit
func (r *registerGcrType) GetLtdcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldLtdcenMask) != 0
}

// SetLtdcen LCD-TFT controller enable bit
func (r *registerGcrType) SetLtdcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldLtdcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldLtdcenMask)
	}
}

// registerSrcrType Shadow Reload Configuration Register
type registerSrcrType uint32

const (
	RegisterSrcrFieldVbrShift = 1
	RegisterSrcrFieldVbrMask  = 0x2
)

// GetVbr Vertical Blanking Reload
func (r *registerSrcrType) GetVbr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrcrFieldVbrMask) != 0
}

// SetVbr Vertical Blanking Reload
func (r *registerSrcrType) SetVbr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrcrFieldVbrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrcrFieldVbrMask)
	}
}

const (
	RegisterSrcrFieldImrShift = 0
	RegisterSrcrFieldImrMask  = 0x1
)

// GetImr Immediate Reload
func (r *registerSrcrType) GetImr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrcrFieldImrMask) != 0
}

// SetImr Immediate Reload
func (r *registerSrcrType) SetImr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrcrFieldImrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrcrFieldImrMask)
	}
}

// registerBccrType Background Color Configuration Register
type registerBccrType uint32

const (
	RegisterBccrFieldBcblueShift = 0
	RegisterBccrFieldBcblueMask  = 0xff
)

// GetBcblue Background Color Blue value
func (r *registerBccrType) GetBcblue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBccrFieldBcblueMask) >> RegisterBccrFieldBcblueShift)
}

// SetBcblue Background Color Blue value
func (r *registerBccrType) SetBcblue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBccrFieldBcblueMask)|(uint32(value)<<RegisterBccrFieldBcblueShift))
}

const (
	RegisterBccrFieldBcgreenShift = 8
	RegisterBccrFieldBcgreenMask  = 0xff00
)

// GetBcgreen Background Color Green value
func (r *registerBccrType) GetBcgreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBccrFieldBcgreenMask) >> RegisterBccrFieldBcgreenShift)
}

// SetBcgreen Background Color Green value
func (r *registerBccrType) SetBcgreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBccrFieldBcgreenMask)|(uint32(value)<<RegisterBccrFieldBcgreenShift))
}

const (
	RegisterBccrFieldBcredShift = 16
	RegisterBccrFieldBcredMask  = 0xff0000
)

// GetBcred Background Color Red value
func (r *registerBccrType) GetBcred() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBccrFieldBcredMask) >> RegisterBccrFieldBcredShift)
}

// SetBcred Background Color Red value
func (r *registerBccrType) SetBcred(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBccrFieldBcredMask)|(uint32(value)<<RegisterBccrFieldBcredShift))
}

// registerIerType Interrupt Enable Register
type registerIerType uint32

const (
	RegisterIerFieldRrieShift = 3
	RegisterIerFieldRrieMask  = 0x8
)

// GetRrie Register Reload interrupt enable
func (r *registerIerType) GetRrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRrieMask) != 0
}

// SetRrie Register Reload interrupt enable
func (r *registerIerType) SetRrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRrieMask)
	}
}

const (
	RegisterIerFieldTerrieShift = 2
	RegisterIerFieldTerrieMask  = 0x4
)

// GetTerrie Transfer Error Interrupt Enable
func (r *registerIerType) GetTerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTerrieMask) != 0
}

// SetTerrie Transfer Error Interrupt Enable
func (r *registerIerType) SetTerrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTerrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTerrieMask)
	}
}

const (
	RegisterIerFieldFuieShift = 1
	RegisterIerFieldFuieMask  = 0x2
)

// GetFuie FIFO Underrun Interrupt Enable
func (r *registerIerType) GetFuie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldFuieMask) != 0
}

// SetFuie FIFO Underrun Interrupt Enable
func (r *registerIerType) SetFuie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldFuieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldFuieMask)
	}
}

const (
	RegisterIerFieldLieShift = 0
	RegisterIerFieldLieMask  = 0x1
)

// GetLie Line Interrupt Enable
func (r *registerIerType) GetLie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldLieMask) != 0
}

// SetLie Line Interrupt Enable
func (r *registerIerType) SetLie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldLieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldLieMask)
	}
}

// registerIsrType Interrupt Status Register
type registerIsrType uint32

const (
	RegisterIsrFieldRrifShift = 3
	RegisterIsrFieldRrifMask  = 0x8
)

// GetRrif Register Reload Interrupt Flag
func (r *registerIsrType) GetRrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRrifMask) != 0
}

// SetRrif Register Reload Interrupt Flag
func (r *registerIsrType) SetRrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRrifMask)
	}
}

const (
	RegisterIsrFieldTerrifShift = 2
	RegisterIsrFieldTerrifMask  = 0x4
)

// GetTerrif Transfer Error interrupt flag
func (r *registerIsrType) GetTerrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTerrifMask) != 0
}

// SetTerrif Transfer Error interrupt flag
func (r *registerIsrType) SetTerrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTerrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTerrifMask)
	}
}

const (
	RegisterIsrFieldFuifShift = 1
	RegisterIsrFieldFuifMask  = 0x2
)

// GetFuif FIFO Underrun Interrupt flag
func (r *registerIsrType) GetFuif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldFuifMask) != 0
}

// SetFuif FIFO Underrun Interrupt flag
func (r *registerIsrType) SetFuif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldFuifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldFuifMask)
	}
}

const (
	RegisterIsrFieldLifShift = 0
	RegisterIsrFieldLifMask  = 0x1
)

// GetLif Line Interrupt flag
func (r *registerIsrType) GetLif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldLifMask) != 0
}

// SetLif Line Interrupt flag
func (r *registerIsrType) SetLif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldLifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldLifMask)
	}
}

// registerIcrType Interrupt Clear Register
type registerIcrType uint32

const (
	RegisterIcrFieldCrrifShift = 3
	RegisterIcrFieldCrrifMask  = 0x8
)

// GetCrrif Clears Register Reload Interrupt Flag
func (r *registerIcrType) GetCrrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCrrifMask) != 0
}

// SetCrrif Clears Register Reload Interrupt Flag
func (r *registerIcrType) SetCrrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCrrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCrrifMask)
	}
}

const (
	RegisterIcrFieldCterrifShift = 2
	RegisterIcrFieldCterrifMask  = 0x4
)

// GetCterrif Clears the Transfer Error Interrupt Flag
func (r *registerIcrType) GetCterrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCterrifMask) != 0
}

// SetCterrif Clears the Transfer Error Interrupt Flag
func (r *registerIcrType) SetCterrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCterrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCterrifMask)
	}
}

const (
	RegisterIcrFieldCfuifShift = 1
	RegisterIcrFieldCfuifMask  = 0x2
)

// GetCfuif Clears the FIFO Underrun Interrupt flag
func (r *registerIcrType) GetCfuif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCfuifMask) != 0
}

// SetCfuif Clears the FIFO Underrun Interrupt flag
func (r *registerIcrType) SetCfuif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCfuifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCfuifMask)
	}
}

const (
	RegisterIcrFieldClifShift = 0
	RegisterIcrFieldClifMask  = 0x1
)

// GetClif Clears the Line Interrupt Flag
func (r *registerIcrType) GetClif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldClifMask) != 0
}

// SetClif Clears the Line Interrupt Flag
func (r *registerIcrType) SetClif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldClifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldClifMask)
	}
}

// registerLipcrType Line Interrupt Position Configuration Register
type registerLipcrType uint32

const (
	RegisterLipcrFieldLiposShift = 0
	RegisterLipcrFieldLiposMask  = 0x7ff
)

// GetLipos Line Interrupt Position
func (r *registerLipcrType) GetLipos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterLipcrFieldLiposMask) >> RegisterLipcrFieldLiposShift)
}

// SetLipos Line Interrupt Position
func (r *registerLipcrType) SetLipos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterLipcrFieldLiposMask)|(uint32(value)<<RegisterLipcrFieldLiposShift))
}

// registerCpsrType Current Position Status Register
type registerCpsrType uint32

const (
	RegisterCpsrFieldCxposShift = 16
	RegisterCpsrFieldCxposMask  = 0xffff0000
)

// GetCxpos Current X Position
func (r *registerCpsrType) GetCxpos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpsrFieldCxposMask) >> RegisterCpsrFieldCxposShift)
}

// SetCxpos Current X Position
func (r *registerCpsrType) SetCxpos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpsrFieldCxposMask)|(uint32(value)<<RegisterCpsrFieldCxposShift))
}

const (
	RegisterCpsrFieldCyposShift = 0
	RegisterCpsrFieldCyposMask  = 0xffff
)

// GetCypos Current Y Position
func (r *registerCpsrType) GetCypos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpsrFieldCyposMask) >> RegisterCpsrFieldCyposShift)
}

// SetCypos Current Y Position
func (r *registerCpsrType) SetCypos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpsrFieldCyposMask)|(uint32(value)<<RegisterCpsrFieldCyposShift))
}

// registerCdsrType Current Display Status Register
type registerCdsrType uint32

const (
	RegisterCdsrFieldHsyncsShift = 3
	RegisterCdsrFieldHsyncsMask  = 0x8
)

// GetHsyncs Horizontal Synchronization display Status
func (r *registerCdsrType) GetHsyncs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCdsrFieldHsyncsMask) != 0
}

// SetHsyncs Horizontal Synchronization display Status
func (r *registerCdsrType) SetHsyncs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCdsrFieldHsyncsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCdsrFieldHsyncsMask)
	}
}

const (
	RegisterCdsrFieldVsyncsShift = 2
	RegisterCdsrFieldVsyncsMask  = 0x4
)

// GetVsyncs Vertical Synchronization display Status
func (r *registerCdsrType) GetVsyncs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCdsrFieldVsyncsMask) != 0
}

// SetVsyncs Vertical Synchronization display Status
func (r *registerCdsrType) SetVsyncs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCdsrFieldVsyncsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCdsrFieldVsyncsMask)
	}
}

const (
	RegisterCdsrFieldHdesShift = 1
	RegisterCdsrFieldHdesMask  = 0x2
)

// GetHdes Horizontal Data Enable display Status
func (r *registerCdsrType) GetHdes() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCdsrFieldHdesMask) != 0
}

// SetHdes Horizontal Data Enable display Status
func (r *registerCdsrType) SetHdes(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCdsrFieldHdesMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCdsrFieldHdesMask)
	}
}

const (
	RegisterCdsrFieldVdesShift = 0
	RegisterCdsrFieldVdesMask  = 0x1
)

// GetVdes Vertical Data Enable display Status
func (r *registerCdsrType) GetVdes() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCdsrFieldVdesMask) != 0
}

// SetVdes Vertical Data Enable display Status
func (r *registerCdsrType) SetVdes(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCdsrFieldVdesMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCdsrFieldVdesMask)
	}
}

// registerL1crType Layerx Control Register
type registerL1crType uint32

const (
	RegisterL1crFieldClutenShift = 4
	RegisterL1crFieldClutenMask  = 0x10
)

// GetCluten Color Look-Up Table Enable
func (r *registerL1crType) GetCluten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL1crFieldClutenMask) != 0
}

// SetCluten Color Look-Up Table Enable
func (r *registerL1crType) SetCluten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL1crFieldClutenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL1crFieldClutenMask)
	}
}

const (
	RegisterL1crFieldColkenShift = 1
	RegisterL1crFieldColkenMask  = 0x2
)

// GetColken Color Keying Enable
func (r *registerL1crType) GetColken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL1crFieldColkenMask) != 0
}

// SetColken Color Keying Enable
func (r *registerL1crType) SetColken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL1crFieldColkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL1crFieldColkenMask)
	}
}

const (
	RegisterL1crFieldLenShift = 0
	RegisterL1crFieldLenMask  = 0x1
)

// GetLen Layer Enable
func (r *registerL1crType) GetLen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL1crFieldLenMask) != 0
}

// SetLen Layer Enable
func (r *registerL1crType) SetLen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL1crFieldLenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL1crFieldLenMask)
	}
}

// registerL1whpcrType Layerx Window Horizontal Position Configuration Register
type registerL1whpcrType uint32

const (
	RegisterL1whpcrFieldWhspposShift = 16
	RegisterL1whpcrFieldWhspposMask  = 0xfff0000
)

// GetWhsppos Window Horizontal Stop Position
func (r *registerL1whpcrType) GetWhsppos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1whpcrFieldWhspposMask) >> RegisterL1whpcrFieldWhspposShift)
}

// SetWhsppos Window Horizontal Stop Position
func (r *registerL1whpcrType) SetWhsppos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1whpcrFieldWhspposMask)|(uint32(value)<<RegisterL1whpcrFieldWhspposShift))
}

const (
	RegisterL1whpcrFieldWhstposShift = 0
	RegisterL1whpcrFieldWhstposMask  = 0xfff
)

// GetWhstpos Window Horizontal Start Position
func (r *registerL1whpcrType) GetWhstpos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1whpcrFieldWhstposMask) >> RegisterL1whpcrFieldWhstposShift)
}

// SetWhstpos Window Horizontal Start Position
func (r *registerL1whpcrType) SetWhstpos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1whpcrFieldWhstposMask)|(uint32(value)<<RegisterL1whpcrFieldWhstposShift))
}

// registerL1wvpcrType Layerx Window Vertical Position Configuration Register
type registerL1wvpcrType uint32

const (
	RegisterL1wvpcrFieldWvspposShift = 16
	RegisterL1wvpcrFieldWvspposMask  = 0x7ff0000
)

// GetWvsppos Window Vertical Stop Position
func (r *registerL1wvpcrType) GetWvsppos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1wvpcrFieldWvspposMask) >> RegisterL1wvpcrFieldWvspposShift)
}

// SetWvsppos Window Vertical Stop Position
func (r *registerL1wvpcrType) SetWvsppos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1wvpcrFieldWvspposMask)|(uint32(value)<<RegisterL1wvpcrFieldWvspposShift))
}

const (
	RegisterL1wvpcrFieldWvstposShift = 0
	RegisterL1wvpcrFieldWvstposMask  = 0x7ff
)

// GetWvstpos Window Vertical Start Position
func (r *registerL1wvpcrType) GetWvstpos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1wvpcrFieldWvstposMask) >> RegisterL1wvpcrFieldWvstposShift)
}

// SetWvstpos Window Vertical Start Position
func (r *registerL1wvpcrType) SetWvstpos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1wvpcrFieldWvstposMask)|(uint32(value)<<RegisterL1wvpcrFieldWvstposShift))
}

// registerL1ckcrType Layerx Color Keying Configuration Register
type registerL1ckcrType uint32

const (
	RegisterL1ckcrFieldCkredShift = 16
	RegisterL1ckcrFieldCkredMask  = 0xff0000
)

// GetCkred Color Key Red value
func (r *registerL1ckcrType) GetCkred() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1ckcrFieldCkredMask) >> RegisterL1ckcrFieldCkredShift)
}

// SetCkred Color Key Red value
func (r *registerL1ckcrType) SetCkred(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1ckcrFieldCkredMask)|(uint32(value)<<RegisterL1ckcrFieldCkredShift))
}

const (
	RegisterL1ckcrFieldCkgreenShift = 8
	RegisterL1ckcrFieldCkgreenMask  = 0xff00
)

// GetCkgreen Color Key Green value
func (r *registerL1ckcrType) GetCkgreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1ckcrFieldCkgreenMask) >> RegisterL1ckcrFieldCkgreenShift)
}

// SetCkgreen Color Key Green value
func (r *registerL1ckcrType) SetCkgreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1ckcrFieldCkgreenMask)|(uint32(value)<<RegisterL1ckcrFieldCkgreenShift))
}

const (
	RegisterL1ckcrFieldCkblueShift = 0
	RegisterL1ckcrFieldCkblueMask  = 0xff
)

// GetCkblue Color Key Blue value
func (r *registerL1ckcrType) GetCkblue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1ckcrFieldCkblueMask) >> RegisterL1ckcrFieldCkblueShift)
}

// SetCkblue Color Key Blue value
func (r *registerL1ckcrType) SetCkblue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1ckcrFieldCkblueMask)|(uint32(value)<<RegisterL1ckcrFieldCkblueShift))
}

// registerL1pfcrType Layerx Pixel Format Configuration Register
type registerL1pfcrType uint32

const (
	RegisterL1pfcrFieldPfShift = 0
	RegisterL1pfcrFieldPfMask  = 0x7
)

// GetPf Pixel Format
func (r *registerL1pfcrType) GetPf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1pfcrFieldPfMask) >> RegisterL1pfcrFieldPfShift)
}

// SetPf Pixel Format
func (r *registerL1pfcrType) SetPf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1pfcrFieldPfMask)|(uint32(value)<<RegisterL1pfcrFieldPfShift))
}

// registerL1cacrType Layerx Constant Alpha Configuration Register
type registerL1cacrType uint32

const (
	RegisterL1cacrFieldConstaShift = 0
	RegisterL1cacrFieldConstaMask  = 0xff
)

// GetConsta Constant Alpha
func (r *registerL1cacrType) GetConsta() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1cacrFieldConstaMask) >> RegisterL1cacrFieldConstaShift)
}

// SetConsta Constant Alpha
func (r *registerL1cacrType) SetConsta(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1cacrFieldConstaMask)|(uint32(value)<<RegisterL1cacrFieldConstaShift))
}

// registerL1dccrType Layerx Default Color Configuration Register
type registerL1dccrType uint32

const (
	RegisterL1dccrFieldDcalphaShift = 24
	RegisterL1dccrFieldDcalphaMask  = 0xff000000
)

// GetDcalpha Default Color Alpha
func (r *registerL1dccrType) GetDcalpha() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1dccrFieldDcalphaMask) >> RegisterL1dccrFieldDcalphaShift)
}

// SetDcalpha Default Color Alpha
func (r *registerL1dccrType) SetDcalpha(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1dccrFieldDcalphaMask)|(uint32(value)<<RegisterL1dccrFieldDcalphaShift))
}

const (
	RegisterL1dccrFieldDcredShift = 16
	RegisterL1dccrFieldDcredMask  = 0xff0000
)

// GetDcred Default Color Red
func (r *registerL1dccrType) GetDcred() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1dccrFieldDcredMask) >> RegisterL1dccrFieldDcredShift)
}

// SetDcred Default Color Red
func (r *registerL1dccrType) SetDcred(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1dccrFieldDcredMask)|(uint32(value)<<RegisterL1dccrFieldDcredShift))
}

const (
	RegisterL1dccrFieldDcgreenShift = 8
	RegisterL1dccrFieldDcgreenMask  = 0xff00
)

// GetDcgreen Default Color Green
func (r *registerL1dccrType) GetDcgreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1dccrFieldDcgreenMask) >> RegisterL1dccrFieldDcgreenShift)
}

// SetDcgreen Default Color Green
func (r *registerL1dccrType) SetDcgreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1dccrFieldDcgreenMask)|(uint32(value)<<RegisterL1dccrFieldDcgreenShift))
}

const (
	RegisterL1dccrFieldDcblueShift = 0
	RegisterL1dccrFieldDcblueMask  = 0xff
)

// GetDcblue Default Color Blue
func (r *registerL1dccrType) GetDcblue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1dccrFieldDcblueMask) >> RegisterL1dccrFieldDcblueShift)
}

// SetDcblue Default Color Blue
func (r *registerL1dccrType) SetDcblue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1dccrFieldDcblueMask)|(uint32(value)<<RegisterL1dccrFieldDcblueShift))
}

// registerL1bfcrType Layerx Blending Factors Configuration Register
type registerL1bfcrType uint32

const (
	RegisterL1bfcrFieldBf1Shift = 8
	RegisterL1bfcrFieldBf1Mask  = 0x700
)

// GetBf1 Blending Factor 1
func (r *registerL1bfcrType) GetBf1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1bfcrFieldBf1Mask) >> RegisterL1bfcrFieldBf1Shift)
}

// SetBf1 Blending Factor 1
func (r *registerL1bfcrType) SetBf1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1bfcrFieldBf1Mask)|(uint32(value)<<RegisterL1bfcrFieldBf1Shift))
}

const (
	RegisterL1bfcrFieldBf2Shift = 0
	RegisterL1bfcrFieldBf2Mask  = 0x7
)

// GetBf2 Blending Factor 2
func (r *registerL1bfcrType) GetBf2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1bfcrFieldBf2Mask) >> RegisterL1bfcrFieldBf2Shift)
}

// SetBf2 Blending Factor 2
func (r *registerL1bfcrType) SetBf2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1bfcrFieldBf2Mask)|(uint32(value)<<RegisterL1bfcrFieldBf2Shift))
}

// registerL1cfbarType Layerx Color Frame Buffer Address Register
type registerL1cfbarType uint32

const (
	RegisterL1cfbarFieldCfbaddShift = 0
	RegisterL1cfbarFieldCfbaddMask  = 0xffffffff
)

// GetCfbadd Color Frame Buffer Start Address
func (r *registerL1cfbarType) GetCfbadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterL1cfbarFieldCfbaddMask) >> RegisterL1cfbarFieldCfbaddShift)
}

// SetCfbadd Color Frame Buffer Start Address
func (r *registerL1cfbarType) SetCfbadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1cfbarFieldCfbaddMask)|(uint32(value)<<RegisterL1cfbarFieldCfbaddShift))
}

// registerL1cfblrType Layerx Color Frame Buffer Length Register
type registerL1cfblrType uint32

const (
	RegisterL1cfblrFieldCfbpShift = 16
	RegisterL1cfblrFieldCfbpMask  = 0x1fff0000
)

// GetCfbp Color Frame Buffer Pitch in bytes
func (r *registerL1cfblrType) GetCfbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1cfblrFieldCfbpMask) >> RegisterL1cfblrFieldCfbpShift)
}

// SetCfbp Color Frame Buffer Pitch in bytes
func (r *registerL1cfblrType) SetCfbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1cfblrFieldCfbpMask)|(uint32(value)<<RegisterL1cfblrFieldCfbpShift))
}

const (
	RegisterL1cfblrFieldCfbllShift = 0
	RegisterL1cfblrFieldCfbllMask  = 0x1fff
)

// GetCfbll Color Frame Buffer Line Length
func (r *registerL1cfblrType) GetCfbll() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1cfblrFieldCfbllMask) >> RegisterL1cfblrFieldCfbllShift)
}

// SetCfbll Color Frame Buffer Line Length
func (r *registerL1cfblrType) SetCfbll(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1cfblrFieldCfbllMask)|(uint32(value)<<RegisterL1cfblrFieldCfbllShift))
}

// registerL1cfblnrType Layerx ColorFrame Buffer Line Number Register
type registerL1cfblnrType uint32

const (
	RegisterL1cfblnrFieldCfblnbrShift = 0
	RegisterL1cfblnrFieldCfblnbrMask  = 0x7ff
)

// GetCfblnbr Frame Buffer Line Number
func (r *registerL1cfblnrType) GetCfblnbr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1cfblnrFieldCfblnbrMask) >> RegisterL1cfblnrFieldCfblnbrShift)
}

// SetCfblnbr Frame Buffer Line Number
func (r *registerL1cfblnrType) SetCfblnbr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1cfblnrFieldCfblnbrMask)|(uint32(value)<<RegisterL1cfblnrFieldCfblnbrShift))
}

// registerL1clutwrType Layerx CLUT Write Register
type registerL1clutwrType uint32

const (
	RegisterL1clutwrFieldClutaddShift = 24
	RegisterL1clutwrFieldClutaddMask  = 0xff000000
)

// GetClutadd CLUT Address
func (r *registerL1clutwrType) GetClutadd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1clutwrFieldClutaddMask) >> RegisterL1clutwrFieldClutaddShift)
}

// SetClutadd CLUT Address
func (r *registerL1clutwrType) SetClutadd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1clutwrFieldClutaddMask)|(uint32(value)<<RegisterL1clutwrFieldClutaddShift))
}

const (
	RegisterL1clutwrFieldRedShift = 16
	RegisterL1clutwrFieldRedMask  = 0xff0000
)

// GetRed Red value
func (r *registerL1clutwrType) GetRed() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1clutwrFieldRedMask) >> RegisterL1clutwrFieldRedShift)
}

// SetRed Red value
func (r *registerL1clutwrType) SetRed(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1clutwrFieldRedMask)|(uint32(value)<<RegisterL1clutwrFieldRedShift))
}

const (
	RegisterL1clutwrFieldGreenShift = 8
	RegisterL1clutwrFieldGreenMask  = 0xff00
)

// GetGreen Green value
func (r *registerL1clutwrType) GetGreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1clutwrFieldGreenMask) >> RegisterL1clutwrFieldGreenShift)
}

// SetGreen Green value
func (r *registerL1clutwrType) SetGreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1clutwrFieldGreenMask)|(uint32(value)<<RegisterL1clutwrFieldGreenShift))
}

const (
	RegisterL1clutwrFieldBlueShift = 0
	RegisterL1clutwrFieldBlueMask  = 0xff
)

// GetBlue Blue value
func (r *registerL1clutwrType) GetBlue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1clutwrFieldBlueMask) >> RegisterL1clutwrFieldBlueShift)
}

// SetBlue Blue value
func (r *registerL1clutwrType) SetBlue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1clutwrFieldBlueMask)|(uint32(value)<<RegisterL1clutwrFieldBlueShift))
}

// registerL2crType Layerx Control Register
type registerL2crType uint32

const (
	RegisterL2crFieldClutenShift = 4
	RegisterL2crFieldClutenMask  = 0x10
)

// GetCluten Color Look-Up Table Enable
func (r *registerL2crType) GetCluten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL2crFieldClutenMask) != 0
}

// SetCluten Color Look-Up Table Enable
func (r *registerL2crType) SetCluten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL2crFieldClutenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL2crFieldClutenMask)
	}
}

const (
	RegisterL2crFieldColkenShift = 1
	RegisterL2crFieldColkenMask  = 0x2
)

// GetColken Color Keying Enable
func (r *registerL2crType) GetColken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL2crFieldColkenMask) != 0
}

// SetColken Color Keying Enable
func (r *registerL2crType) SetColken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL2crFieldColkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL2crFieldColkenMask)
	}
}

const (
	RegisterL2crFieldLenShift = 0
	RegisterL2crFieldLenMask  = 0x1
)

// GetLen Layer Enable
func (r *registerL2crType) GetLen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL2crFieldLenMask) != 0
}

// SetLen Layer Enable
func (r *registerL2crType) SetLen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL2crFieldLenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL2crFieldLenMask)
	}
}

// registerL2whpcrType Layerx Window Horizontal Position Configuration Register
type registerL2whpcrType uint32

const (
	RegisterL2whpcrFieldWhspposShift = 16
	RegisterL2whpcrFieldWhspposMask  = 0xfff0000
)

// GetWhsppos Window Horizontal Stop Position
func (r *registerL2whpcrType) GetWhsppos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2whpcrFieldWhspposMask) >> RegisterL2whpcrFieldWhspposShift)
}

// SetWhsppos Window Horizontal Stop Position
func (r *registerL2whpcrType) SetWhsppos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2whpcrFieldWhspposMask)|(uint32(value)<<RegisterL2whpcrFieldWhspposShift))
}

const (
	RegisterL2whpcrFieldWhstposShift = 0
	RegisterL2whpcrFieldWhstposMask  = 0xfff
)

// GetWhstpos Window Horizontal Start Position
func (r *registerL2whpcrType) GetWhstpos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2whpcrFieldWhstposMask) >> RegisterL2whpcrFieldWhstposShift)
}

// SetWhstpos Window Horizontal Start Position
func (r *registerL2whpcrType) SetWhstpos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2whpcrFieldWhstposMask)|(uint32(value)<<RegisterL2whpcrFieldWhstposShift))
}

// registerL2wvpcrType Layerx Window Vertical Position Configuration Register
type registerL2wvpcrType uint32

const (
	RegisterL2wvpcrFieldWvspposShift = 16
	RegisterL2wvpcrFieldWvspposMask  = 0x7ff0000
)

// GetWvsppos Window Vertical Stop Position
func (r *registerL2wvpcrType) GetWvsppos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2wvpcrFieldWvspposMask) >> RegisterL2wvpcrFieldWvspposShift)
}

// SetWvsppos Window Vertical Stop Position
func (r *registerL2wvpcrType) SetWvsppos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2wvpcrFieldWvspposMask)|(uint32(value)<<RegisterL2wvpcrFieldWvspposShift))
}

const (
	RegisterL2wvpcrFieldWvstposShift = 0
	RegisterL2wvpcrFieldWvstposMask  = 0x7ff
)

// GetWvstpos Window Vertical Start Position
func (r *registerL2wvpcrType) GetWvstpos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2wvpcrFieldWvstposMask) >> RegisterL2wvpcrFieldWvstposShift)
}

// SetWvstpos Window Vertical Start Position
func (r *registerL2wvpcrType) SetWvstpos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2wvpcrFieldWvstposMask)|(uint32(value)<<RegisterL2wvpcrFieldWvstposShift))
}

// registerL2ckcrType Layerx Color Keying Configuration Register
type registerL2ckcrType uint32

const (
	RegisterL2ckcrFieldCkredShift = 16
	RegisterL2ckcrFieldCkredMask  = 0xff0000
)

// GetCkred Color Key Red value
func (r *registerL2ckcrType) GetCkred() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2ckcrFieldCkredMask) >> RegisterL2ckcrFieldCkredShift)
}

// SetCkred Color Key Red value
func (r *registerL2ckcrType) SetCkred(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2ckcrFieldCkredMask)|(uint32(value)<<RegisterL2ckcrFieldCkredShift))
}

const (
	RegisterL2ckcrFieldCkgreenShift = 8
	RegisterL2ckcrFieldCkgreenMask  = 0xff00
)

// GetCkgreen Color Key Green value
func (r *registerL2ckcrType) GetCkgreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2ckcrFieldCkgreenMask) >> RegisterL2ckcrFieldCkgreenShift)
}

// SetCkgreen Color Key Green value
func (r *registerL2ckcrType) SetCkgreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2ckcrFieldCkgreenMask)|(uint32(value)<<RegisterL2ckcrFieldCkgreenShift))
}

const (
	RegisterL2ckcrFieldCkblueShift = 0
	RegisterL2ckcrFieldCkblueMask  = 0xff
)

// GetCkblue Color Key Blue value
func (r *registerL2ckcrType) GetCkblue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2ckcrFieldCkblueMask) >> RegisterL2ckcrFieldCkblueShift)
}

// SetCkblue Color Key Blue value
func (r *registerL2ckcrType) SetCkblue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2ckcrFieldCkblueMask)|(uint32(value)<<RegisterL2ckcrFieldCkblueShift))
}

// registerL2pfcrType Layerx Pixel Format Configuration Register
type registerL2pfcrType uint32

const (
	RegisterL2pfcrFieldPfShift = 0
	RegisterL2pfcrFieldPfMask  = 0x7
)

// GetPf Pixel Format
func (r *registerL2pfcrType) GetPf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2pfcrFieldPfMask) >> RegisterL2pfcrFieldPfShift)
}

// SetPf Pixel Format
func (r *registerL2pfcrType) SetPf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2pfcrFieldPfMask)|(uint32(value)<<RegisterL2pfcrFieldPfShift))
}

// registerL2cacrType Layerx Constant Alpha Configuration Register
type registerL2cacrType uint32

const (
	RegisterL2cacrFieldConstaShift = 0
	RegisterL2cacrFieldConstaMask  = 0xff
)

// GetConsta Constant Alpha
func (r *registerL2cacrType) GetConsta() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2cacrFieldConstaMask) >> RegisterL2cacrFieldConstaShift)
}

// SetConsta Constant Alpha
func (r *registerL2cacrType) SetConsta(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2cacrFieldConstaMask)|(uint32(value)<<RegisterL2cacrFieldConstaShift))
}

// registerL2dccrType Layerx Default Color Configuration Register
type registerL2dccrType uint32

const (
	RegisterL2dccrFieldDcalphaShift = 24
	RegisterL2dccrFieldDcalphaMask  = 0xff000000
)

// GetDcalpha Default Color Alpha
func (r *registerL2dccrType) GetDcalpha() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2dccrFieldDcalphaMask) >> RegisterL2dccrFieldDcalphaShift)
}

// SetDcalpha Default Color Alpha
func (r *registerL2dccrType) SetDcalpha(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2dccrFieldDcalphaMask)|(uint32(value)<<RegisterL2dccrFieldDcalphaShift))
}

const (
	RegisterL2dccrFieldDcredShift = 16
	RegisterL2dccrFieldDcredMask  = 0xff0000
)

// GetDcred Default Color Red
func (r *registerL2dccrType) GetDcred() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2dccrFieldDcredMask) >> RegisterL2dccrFieldDcredShift)
}

// SetDcred Default Color Red
func (r *registerL2dccrType) SetDcred(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2dccrFieldDcredMask)|(uint32(value)<<RegisterL2dccrFieldDcredShift))
}

const (
	RegisterL2dccrFieldDcgreenShift = 8
	RegisterL2dccrFieldDcgreenMask  = 0xff00
)

// GetDcgreen Default Color Green
func (r *registerL2dccrType) GetDcgreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2dccrFieldDcgreenMask) >> RegisterL2dccrFieldDcgreenShift)
}

// SetDcgreen Default Color Green
func (r *registerL2dccrType) SetDcgreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2dccrFieldDcgreenMask)|(uint32(value)<<RegisterL2dccrFieldDcgreenShift))
}

const (
	RegisterL2dccrFieldDcblueShift = 0
	RegisterL2dccrFieldDcblueMask  = 0xff
)

// GetDcblue Default Color Blue
func (r *registerL2dccrType) GetDcblue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2dccrFieldDcblueMask) >> RegisterL2dccrFieldDcblueShift)
}

// SetDcblue Default Color Blue
func (r *registerL2dccrType) SetDcblue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2dccrFieldDcblueMask)|(uint32(value)<<RegisterL2dccrFieldDcblueShift))
}

// registerL2bfcrType Layerx Blending Factors Configuration Register
type registerL2bfcrType uint32

const (
	RegisterL2bfcrFieldBf1Shift = 8
	RegisterL2bfcrFieldBf1Mask  = 0x700
)

// GetBf1 Blending Factor 1
func (r *registerL2bfcrType) GetBf1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2bfcrFieldBf1Mask) >> RegisterL2bfcrFieldBf1Shift)
}

// SetBf1 Blending Factor 1
func (r *registerL2bfcrType) SetBf1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2bfcrFieldBf1Mask)|(uint32(value)<<RegisterL2bfcrFieldBf1Shift))
}

const (
	RegisterL2bfcrFieldBf2Shift = 0
	RegisterL2bfcrFieldBf2Mask  = 0x7
)

// GetBf2 Blending Factor 2
func (r *registerL2bfcrType) GetBf2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2bfcrFieldBf2Mask) >> RegisterL2bfcrFieldBf2Shift)
}

// SetBf2 Blending Factor 2
func (r *registerL2bfcrType) SetBf2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2bfcrFieldBf2Mask)|(uint32(value)<<RegisterL2bfcrFieldBf2Shift))
}

// registerL2cfbarType Layerx Color Frame Buffer Address Register
type registerL2cfbarType uint32

const (
	RegisterL2cfbarFieldCfbaddShift = 0
	RegisterL2cfbarFieldCfbaddMask  = 0xffffffff
)

// GetCfbadd Color Frame Buffer Start Address
func (r *registerL2cfbarType) GetCfbadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterL2cfbarFieldCfbaddMask) >> RegisterL2cfbarFieldCfbaddShift)
}

// SetCfbadd Color Frame Buffer Start Address
func (r *registerL2cfbarType) SetCfbadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2cfbarFieldCfbaddMask)|(uint32(value)<<RegisterL2cfbarFieldCfbaddShift))
}

// registerL2cfblrType Layerx Color Frame Buffer Length Register
type registerL2cfblrType uint32

const (
	RegisterL2cfblrFieldCfbpShift = 16
	RegisterL2cfblrFieldCfbpMask  = 0x1fff0000
)

// GetCfbp Color Frame Buffer Pitch in bytes
func (r *registerL2cfblrType) GetCfbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2cfblrFieldCfbpMask) >> RegisterL2cfblrFieldCfbpShift)
}

// SetCfbp Color Frame Buffer Pitch in bytes
func (r *registerL2cfblrType) SetCfbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2cfblrFieldCfbpMask)|(uint32(value)<<RegisterL2cfblrFieldCfbpShift))
}

const (
	RegisterL2cfblrFieldCfbllShift = 0
	RegisterL2cfblrFieldCfbllMask  = 0x1fff
)

// GetCfbll Color Frame Buffer Line Length
func (r *registerL2cfblrType) GetCfbll() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2cfblrFieldCfbllMask) >> RegisterL2cfblrFieldCfbllShift)
}

// SetCfbll Color Frame Buffer Line Length
func (r *registerL2cfblrType) SetCfbll(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2cfblrFieldCfbllMask)|(uint32(value)<<RegisterL2cfblrFieldCfbllShift))
}

// registerL2cfblnrType Layerx ColorFrame Buffer Line Number Register
type registerL2cfblnrType uint32

const (
	RegisterL2cfblnrFieldCfblnbrShift = 0
	RegisterL2cfblnrFieldCfblnbrMask  = 0x7ff
)

// GetCfblnbr Frame Buffer Line Number
func (r *registerL2cfblnrType) GetCfblnbr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2cfblnrFieldCfblnbrMask) >> RegisterL2cfblnrFieldCfblnbrShift)
}

// SetCfblnbr Frame Buffer Line Number
func (r *registerL2cfblnrType) SetCfblnbr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2cfblnrFieldCfblnbrMask)|(uint32(value)<<RegisterL2cfblnrFieldCfblnbrShift))
}

// registerL2clutwrType Layerx CLUT Write Register
type registerL2clutwrType uint32

const (
	RegisterL2clutwrFieldClutaddShift = 24
	RegisterL2clutwrFieldClutaddMask  = 0xff000000
)

// GetClutadd CLUT Address
func (r *registerL2clutwrType) GetClutadd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2clutwrFieldClutaddMask) >> RegisterL2clutwrFieldClutaddShift)
}

// SetClutadd CLUT Address
func (r *registerL2clutwrType) SetClutadd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2clutwrFieldClutaddMask)|(uint32(value)<<RegisterL2clutwrFieldClutaddShift))
}

const (
	RegisterL2clutwrFieldRedShift = 16
	RegisterL2clutwrFieldRedMask  = 0xff0000
)

// GetRed Red value
func (r *registerL2clutwrType) GetRed() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2clutwrFieldRedMask) >> RegisterL2clutwrFieldRedShift)
}

// SetRed Red value
func (r *registerL2clutwrType) SetRed(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2clutwrFieldRedMask)|(uint32(value)<<RegisterL2clutwrFieldRedShift))
}

const (
	RegisterL2clutwrFieldGreenShift = 8
	RegisterL2clutwrFieldGreenMask  = 0xff00
)

// GetGreen Green value
func (r *registerL2clutwrType) GetGreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2clutwrFieldGreenMask) >> RegisterL2clutwrFieldGreenShift)
}

// SetGreen Green value
func (r *registerL2clutwrType) SetGreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2clutwrFieldGreenMask)|(uint32(value)<<RegisterL2clutwrFieldGreenShift))
}

const (
	RegisterL2clutwrFieldBlueShift = 0
	RegisterL2clutwrFieldBlueMask  = 0xff
)

// GetBlue Blue value
func (r *registerL2clutwrType) GetBlue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2clutwrFieldBlueMask) >> RegisterL2clutwrFieldBlueShift)
}

// SetBlue Blue value
func (r *registerL2clutwrType) SetBlue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2clutwrFieldBlueMask)|(uint32(value)<<RegisterL2clutwrFieldBlueShift))
}
