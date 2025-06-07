//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

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
	Sscr     RegisterSscrType
	Bpcr     RegisterBpcrType
	Awcr     RegisterAwcrType
	Twcr     RegisterTwcrType
	Gcr      RegisterGcrType
	_        [8]byte
	Srcr     RegisterSrcrType
	_        [4]byte
	Bccr     RegisterBccrType
	_        [4]byte
	Ier      RegisterIerType
	Isr      RegisterIsrType
	Icr      RegisterIcrType
	Lipcr    RegisterLipcrType
	Cpsr     RegisterCpsrType
	Cdsr     RegisterCdsrType
	_        [56]byte
	L1cr     RegisterL1crType
	L1whpcr  RegisterL1whpcrType
	L1wvpcr  RegisterL1wvpcrType
	L1ckcr   RegisterL1ckcrType
	L1pfcr   RegisterL1pfcrType
	L1cacr   RegisterL1cacrType
	L1dccr   RegisterL1dccrType
	L1bfcr   RegisterL1bfcrType
	_        [8]byte
	L1cfbar  RegisterL1cfbarType
	L1cfblr  RegisterL1cfblrType
	L1cfblnr RegisterL1cfblnrType
	_        [12]byte
	L1clutwr RegisterL1clutwrType
	_        [60]byte
	L2cr     RegisterL2crType
	L2whpcr  RegisterL2whpcrType
	L2wvpcr  RegisterL2wvpcrType
	L2ckcr   RegisterL2ckcrType
	L2pfcr   RegisterL2pfcrType
	L2cacr   RegisterL2cacrType
	L2dccr   RegisterL2dccrType
	L2bfcr   RegisterL2bfcrType
	_        [8]byte
	L2cfbar  RegisterL2cfbarType
	L2cfblr  RegisterL2cfblrType
	L2cfblnr RegisterL2cfblnrType
	_        [12]byte
	L2clutwr RegisterL2clutwrType
}

// RegisterSscrType Synchronization Size Configuration Register
type RegisterSscrType uint32

func (r *RegisterSscrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSscrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSscrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSscrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSscrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSscrFieldVshShift = 0
	RegisterSscrFieldVshMask  = 0x7ff
)

// GetVsh Vertical Synchronization Height (in units of horizontal scan line)
func (r *RegisterSscrType) GetVsh() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSscrFieldVshMask) >> RegisterSscrFieldVshShift)
}

// SetVsh Vertical Synchronization Height (in units of horizontal scan line)
func (r *RegisterSscrType) SetVsh(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSscrFieldVshMask)|(uint32(value)<<RegisterSscrFieldVshShift))
}

const (
	RegisterSscrFieldHswShift = 16
	RegisterSscrFieldHswMask  = 0x3ff0000
)

// GetHsw Horizontal Synchronization Width (in units of pixel clock period)
func (r *RegisterSscrType) GetHsw() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSscrFieldHswMask) >> RegisterSscrFieldHswShift)
}

// SetHsw Horizontal Synchronization Width (in units of pixel clock period)
func (r *RegisterSscrType) SetHsw(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSscrFieldHswMask)|(uint32(value)<<RegisterSscrFieldHswShift))
}

// RegisterBpcrType Back Porch Configuration Register
type RegisterBpcrType uint32

func (r *RegisterBpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBpcrFieldAvbpShift = 0
	RegisterBpcrFieldAvbpMask  = 0x7ff
)

// GetAvbp Accumulated Vertical back porch (in units of horizontal scan line)
func (r *RegisterBpcrType) GetAvbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBpcrFieldAvbpMask) >> RegisterBpcrFieldAvbpShift)
}

// SetAvbp Accumulated Vertical back porch (in units of horizontal scan line)
func (r *RegisterBpcrType) SetAvbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBpcrFieldAvbpMask)|(uint32(value)<<RegisterBpcrFieldAvbpShift))
}

const (
	RegisterBpcrFieldAhbpShift = 16
	RegisterBpcrFieldAhbpMask  = 0xfff0000
)

// GetAhbp Accumulated Horizontal back porch (in units of pixel clock period)
func (r *RegisterBpcrType) GetAhbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterBpcrFieldAhbpMask) >> RegisterBpcrFieldAhbpShift)
}

// SetAhbp Accumulated Horizontal back porch (in units of pixel clock period)
func (r *RegisterBpcrType) SetAhbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBpcrFieldAhbpMask)|(uint32(value)<<RegisterBpcrFieldAhbpShift))
}

// RegisterAwcrType Active Width Configuration Register
type RegisterAwcrType uint32

func (r *RegisterAwcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAwcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAwcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAwcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAwcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAwcrFieldAahShift = 0
	RegisterAwcrFieldAahMask  = 0x7ff
)

// GetAah Accumulated Active Height (in units of horizontal scan line)
func (r *RegisterAwcrType) GetAah() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterAwcrFieldAahMask) >> RegisterAwcrFieldAahShift)
}

// SetAah Accumulated Active Height (in units of horizontal scan line)
func (r *RegisterAwcrType) SetAah(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAwcrFieldAahMask)|(uint32(value)<<RegisterAwcrFieldAahShift))
}

const (
	RegisterAwcrFieldAavShift = 16
	RegisterAwcrFieldAavMask  = 0xfff0000
)

// GetAav AAV
func (r *RegisterAwcrType) GetAav() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterAwcrFieldAavMask) >> RegisterAwcrFieldAavShift)
}

// SetAav AAV
func (r *RegisterAwcrType) SetAav(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAwcrFieldAavMask)|(uint32(value)<<RegisterAwcrFieldAavShift))
}

// RegisterTwcrType Total Width Configuration Register
type RegisterTwcrType uint32

func (r *RegisterTwcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTwcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTwcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTwcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTwcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTwcrFieldTotalhShift = 0
	RegisterTwcrFieldTotalhMask  = 0x7ff
)

// GetTotalh Total Height (in units of horizontal scan line)
func (r *RegisterTwcrType) GetTotalh() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterTwcrFieldTotalhMask) >> RegisterTwcrFieldTotalhShift)
}

// SetTotalh Total Height (in units of horizontal scan line)
func (r *RegisterTwcrType) SetTotalh(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTwcrFieldTotalhMask)|(uint32(value)<<RegisterTwcrFieldTotalhShift))
}

const (
	RegisterTwcrFieldTotalwShift = 16
	RegisterTwcrFieldTotalwMask  = 0xfff0000
)

// GetTotalw Total Width (in units of pixel clock period)
func (r *RegisterTwcrType) GetTotalw() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterTwcrFieldTotalwMask) >> RegisterTwcrFieldTotalwShift)
}

// SetTotalw Total Width (in units of pixel clock period)
func (r *RegisterTwcrType) SetTotalw(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTwcrFieldTotalwMask)|(uint32(value)<<RegisterTwcrFieldTotalwShift))
}

// RegisterGcrType Global Control Register
type RegisterGcrType uint32

func (r *RegisterGcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterGcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterGcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterGcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterGcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterGcrFieldLtdcenShift = 0
	RegisterGcrFieldLtdcenMask  = 0x1
)

// GetLtdcen LCD-TFT controller enable bit
func (r *RegisterGcrType) GetLtdcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldLtdcenMask) != 0
}

// SetLtdcen LCD-TFT controller enable bit
func (r *RegisterGcrType) SetLtdcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldLtdcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldLtdcenMask)
	}
}

const (
	RegisterGcrFieldDbwShift = 4
	RegisterGcrFieldDbwMask  = 0x70
)

// GetDbw Dither Blue Width
func (r *RegisterGcrType) GetDbw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldDbwMask) >> RegisterGcrFieldDbwShift)
}

const (
	RegisterGcrFieldDgwShift = 8
	RegisterGcrFieldDgwMask  = 0x700
)

// GetDgw Dither Green Width
func (r *RegisterGcrType) GetDgw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldDgwMask) >> RegisterGcrFieldDgwShift)
}

const (
	RegisterGcrFieldDrwShift = 12
	RegisterGcrFieldDrwMask  = 0x7000
)

// GetDrw Dither Red Width
func (r *RegisterGcrType) GetDrw() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldDrwMask) >> RegisterGcrFieldDrwShift)
}

const (
	RegisterGcrFieldDenShift = 16
	RegisterGcrFieldDenMask  = 0x10000
)

// GetDen Dither Enable
func (r *RegisterGcrType) GetDen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldDenMask) != 0
}

// SetDen Dither Enable
func (r *RegisterGcrType) SetDen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldDenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldDenMask)
	}
}

const (
	RegisterGcrFieldPcpolShift = 28
	RegisterGcrFieldPcpolMask  = 0x10000000
)

// GetPcpol Pixel Clock Polarity
func (r *RegisterGcrType) GetPcpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldPcpolMask) != 0
}

// SetPcpol Pixel Clock Polarity
func (r *RegisterGcrType) SetPcpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldPcpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldPcpolMask)
	}
}

const (
	RegisterGcrFieldDepolShift = 29
	RegisterGcrFieldDepolMask  = 0x20000000
)

// GetDepol Data Enable Polarity
func (r *RegisterGcrType) GetDepol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldDepolMask) != 0
}

// SetDepol Data Enable Polarity
func (r *RegisterGcrType) SetDepol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldDepolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldDepolMask)
	}
}

const (
	RegisterGcrFieldVspolShift = 30
	RegisterGcrFieldVspolMask  = 0x40000000
)

// GetVspol Vertical Synchronization Polarity
func (r *RegisterGcrType) GetVspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldVspolMask) != 0
}

// SetVspol Vertical Synchronization Polarity
func (r *RegisterGcrType) SetVspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldVspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldVspolMask)
	}
}

const (
	RegisterGcrFieldHspolShift = 31
	RegisterGcrFieldHspolMask  = 0x80000000
)

// GetHspol Horizontal Synchronization Polarity
func (r *RegisterGcrType) GetHspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldHspolMask) != 0
}

// SetHspol Horizontal Synchronization Polarity
func (r *RegisterGcrType) SetHspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldHspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldHspolMask)
	}
}

// RegisterSrcrType Shadow Reload Configuration Register
type RegisterSrcrType uint32

func (r *RegisterSrcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterSrcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterSrcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterSrcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterSrcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterSrcrFieldImrShift = 0
	RegisterSrcrFieldImrMask  = 0x1
)

// GetImr Immediate Reload
func (r *RegisterSrcrType) GetImr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrcrFieldImrMask) != 0
}

// SetImr Immediate Reload
func (r *RegisterSrcrType) SetImr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrcrFieldImrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrcrFieldImrMask)
	}
}

const (
	RegisterSrcrFieldVbrShift = 1
	RegisterSrcrFieldVbrMask  = 0x2
)

// GetVbr Vertical Blanking Reload
func (r *RegisterSrcrType) GetVbr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrcrFieldVbrMask) != 0
}

// SetVbr Vertical Blanking Reload
func (r *RegisterSrcrType) SetVbr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrcrFieldVbrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrcrFieldVbrMask)
	}
}

// RegisterBccrType Background Color Configuration Register
type RegisterBccrType uint32

func (r *RegisterBccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBccrFieldBcblueShift = 0
	RegisterBccrFieldBcblueMask  = 0xff
)

// GetBcblue Background Color Blue value
func (r *RegisterBccrType) GetBcblue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBccrFieldBcblueMask) >> RegisterBccrFieldBcblueShift)
}

// SetBcblue Background Color Blue value
func (r *RegisterBccrType) SetBcblue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBccrFieldBcblueMask)|(uint32(value)<<RegisterBccrFieldBcblueShift))
}

const (
	RegisterBccrFieldBcgreenShift = 8
	RegisterBccrFieldBcgreenMask  = 0xff00
)

// GetBcgreen Background Color Green value
func (r *RegisterBccrType) GetBcgreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBccrFieldBcgreenMask) >> RegisterBccrFieldBcgreenShift)
}

// SetBcgreen Background Color Green value
func (r *RegisterBccrType) SetBcgreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBccrFieldBcgreenMask)|(uint32(value)<<RegisterBccrFieldBcgreenShift))
}

const (
	RegisterBccrFieldBcredShift = 16
	RegisterBccrFieldBcredMask  = 0xff0000
)

// GetBcred Background Color Red value
func (r *RegisterBccrType) GetBcred() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBccrFieldBcredMask) >> RegisterBccrFieldBcredShift)
}

// SetBcred Background Color Red value
func (r *RegisterBccrType) SetBcred(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBccrFieldBcredMask)|(uint32(value)<<RegisterBccrFieldBcredShift))
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
	RegisterIerFieldLieShift = 0
	RegisterIerFieldLieMask  = 0x1
)

// GetLie Line Interrupt Enable
func (r *RegisterIerType) GetLie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldLieMask) != 0
}

// SetLie Line Interrupt Enable
func (r *RegisterIerType) SetLie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldLieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldLieMask)
	}
}

const (
	RegisterIerFieldFuieShift = 1
	RegisterIerFieldFuieMask  = 0x2
)

// GetFuie FIFO Underrun Interrupt Enable
func (r *RegisterIerType) GetFuie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldFuieMask) != 0
}

// SetFuie FIFO Underrun Interrupt Enable
func (r *RegisterIerType) SetFuie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldFuieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldFuieMask)
	}
}

const (
	RegisterIerFieldTerrieShift = 2
	RegisterIerFieldTerrieMask  = 0x4
)

// GetTerrie Transfer Error Interrupt Enable
func (r *RegisterIerType) GetTerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTerrieMask) != 0
}

// SetTerrie Transfer Error Interrupt Enable
func (r *RegisterIerType) SetTerrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTerrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTerrieMask)
	}
}

const (
	RegisterIerFieldRrieShift = 3
	RegisterIerFieldRrieMask  = 0x8
)

// GetRrie Register Reload interrupt enable
func (r *RegisterIerType) GetRrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRrieMask) != 0
}

// SetRrie Register Reload interrupt enable
func (r *RegisterIerType) SetRrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRrieMask)
	}
}

// RegisterIsrType Interrupt Status Register
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
	RegisterIsrFieldLifShift = 0
	RegisterIsrFieldLifMask  = 0x1
)

// GetLif Line Interrupt flag
func (r *RegisterIsrType) GetLif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldLifMask) != 0
}

// SetLif Line Interrupt flag
func (r *RegisterIsrType) SetLif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldLifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldLifMask)
	}
}

const (
	RegisterIsrFieldFuifShift = 1
	RegisterIsrFieldFuifMask  = 0x2
)

// GetFuif FIFO Underrun Interrupt flag
func (r *RegisterIsrType) GetFuif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldFuifMask) != 0
}

// SetFuif FIFO Underrun Interrupt flag
func (r *RegisterIsrType) SetFuif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldFuifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldFuifMask)
	}
}

const (
	RegisterIsrFieldTerrifShift = 2
	RegisterIsrFieldTerrifMask  = 0x4
)

// GetTerrif Transfer Error interrupt flag
func (r *RegisterIsrType) GetTerrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTerrifMask) != 0
}

// SetTerrif Transfer Error interrupt flag
func (r *RegisterIsrType) SetTerrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTerrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTerrifMask)
	}
}

const (
	RegisterIsrFieldRrifShift = 3
	RegisterIsrFieldRrifMask  = 0x8
)

// GetRrif Register Reload Interrupt Flag
func (r *RegisterIsrType) GetRrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRrifMask) != 0
}

// SetRrif Register Reload Interrupt Flag
func (r *RegisterIsrType) SetRrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRrifMask)
	}
}

// RegisterIcrType Interrupt Clear Register
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
	RegisterIcrFieldClifShift = 0
	RegisterIcrFieldClifMask  = 0x1
)

// GetClif Clears the Line Interrupt Flag
func (r *RegisterIcrType) GetClif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldClifMask) != 0
}

// SetClif Clears the Line Interrupt Flag
func (r *RegisterIcrType) SetClif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldClifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldClifMask)
	}
}

const (
	RegisterIcrFieldCfuifShift = 1
	RegisterIcrFieldCfuifMask  = 0x2
)

// GetCfuif Clears the FIFO Underrun Interrupt flag
func (r *RegisterIcrType) GetCfuif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCfuifMask) != 0
}

// SetCfuif Clears the FIFO Underrun Interrupt flag
func (r *RegisterIcrType) SetCfuif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCfuifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCfuifMask)
	}
}

const (
	RegisterIcrFieldCterrifShift = 2
	RegisterIcrFieldCterrifMask  = 0x4
)

// GetCterrif Clears the Transfer Error Interrupt Flag
func (r *RegisterIcrType) GetCterrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCterrifMask) != 0
}

// SetCterrif Clears the Transfer Error Interrupt Flag
func (r *RegisterIcrType) SetCterrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCterrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCterrifMask)
	}
}

const (
	RegisterIcrFieldCrrifShift = 3
	RegisterIcrFieldCrrifMask  = 0x8
)

// GetCrrif Clears Register Reload Interrupt Flag
func (r *RegisterIcrType) GetCrrif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCrrifMask) != 0
}

// SetCrrif Clears Register Reload Interrupt Flag
func (r *RegisterIcrType) SetCrrif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCrrifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCrrifMask)
	}
}

// RegisterLipcrType Line Interrupt Position Configuration Register
type RegisterLipcrType uint32

func (r *RegisterLipcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterLipcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterLipcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterLipcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterLipcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterLipcrFieldLiposShift = 0
	RegisterLipcrFieldLiposMask  = 0x7ff
)

// GetLipos Line Interrupt Position
func (r *RegisterLipcrType) GetLipos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterLipcrFieldLiposMask) >> RegisterLipcrFieldLiposShift)
}

// SetLipos Line Interrupt Position
func (r *RegisterLipcrType) SetLipos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterLipcrFieldLiposMask)|(uint32(value)<<RegisterLipcrFieldLiposShift))
}

// RegisterCpsrType Current Position Status Register
type RegisterCpsrType uint32

func (r *RegisterCpsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCpsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCpsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCpsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCpsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCpsrFieldCyposShift = 0
	RegisterCpsrFieldCyposMask  = 0xffff
)

// GetCypos Current Y Position
func (r *RegisterCpsrType) GetCypos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpsrFieldCyposMask) >> RegisterCpsrFieldCyposShift)
}

// SetCypos Current Y Position
func (r *RegisterCpsrType) SetCypos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpsrFieldCyposMask)|(uint32(value)<<RegisterCpsrFieldCyposShift))
}

const (
	RegisterCpsrFieldCxposShift = 16
	RegisterCpsrFieldCxposMask  = 0xffff0000
)

// GetCxpos Current X Position
func (r *RegisterCpsrType) GetCxpos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCpsrFieldCxposMask) >> RegisterCpsrFieldCxposShift)
}

// SetCxpos Current X Position
func (r *RegisterCpsrType) SetCxpos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCpsrFieldCxposMask)|(uint32(value)<<RegisterCpsrFieldCxposShift))
}

// RegisterCdsrType Current Display Status Register
type RegisterCdsrType uint32

func (r *RegisterCdsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCdsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCdsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCdsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCdsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCdsrFieldVdesShift = 0
	RegisterCdsrFieldVdesMask  = 0x1
)

// GetVdes Vertical Data Enable display Status
func (r *RegisterCdsrType) GetVdes() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCdsrFieldVdesMask) != 0
}

// SetVdes Vertical Data Enable display Status
func (r *RegisterCdsrType) SetVdes(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCdsrFieldVdesMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCdsrFieldVdesMask)
	}
}

const (
	RegisterCdsrFieldHdesShift = 1
	RegisterCdsrFieldHdesMask  = 0x2
)

// GetHdes Horizontal Data Enable display Status
func (r *RegisterCdsrType) GetHdes() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCdsrFieldHdesMask) != 0
}

// SetHdes Horizontal Data Enable display Status
func (r *RegisterCdsrType) SetHdes(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCdsrFieldHdesMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCdsrFieldHdesMask)
	}
}

const (
	RegisterCdsrFieldVsyncsShift = 2
	RegisterCdsrFieldVsyncsMask  = 0x4
)

// GetVsyncs Vertical Synchronization display Status
func (r *RegisterCdsrType) GetVsyncs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCdsrFieldVsyncsMask) != 0
}

// SetVsyncs Vertical Synchronization display Status
func (r *RegisterCdsrType) SetVsyncs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCdsrFieldVsyncsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCdsrFieldVsyncsMask)
	}
}

const (
	RegisterCdsrFieldHsyncsShift = 3
	RegisterCdsrFieldHsyncsMask  = 0x8
)

// GetHsyncs Horizontal Synchronization display Status
func (r *RegisterCdsrType) GetHsyncs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCdsrFieldHsyncsMask) != 0
}

// SetHsyncs Horizontal Synchronization display Status
func (r *RegisterCdsrType) SetHsyncs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCdsrFieldHsyncsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCdsrFieldHsyncsMask)
	}
}

// RegisterL1crType Layerx Control Register
type RegisterL1crType uint32

func (r *RegisterL1crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1crFieldLenShift = 0
	RegisterL1crFieldLenMask  = 0x1
)

// GetLen Layer Enable
func (r *RegisterL1crType) GetLen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL1crFieldLenMask) != 0
}

// SetLen Layer Enable
func (r *RegisterL1crType) SetLen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL1crFieldLenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL1crFieldLenMask)
	}
}

const (
	RegisterL1crFieldColkenShift = 1
	RegisterL1crFieldColkenMask  = 0x2
)

// GetColken Color Keying Enable
func (r *RegisterL1crType) GetColken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL1crFieldColkenMask) != 0
}

// SetColken Color Keying Enable
func (r *RegisterL1crType) SetColken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL1crFieldColkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL1crFieldColkenMask)
	}
}

const (
	RegisterL1crFieldClutenShift = 4
	RegisterL1crFieldClutenMask  = 0x10
)

// GetCluten Color Look-Up Table Enable
func (r *RegisterL1crType) GetCluten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL1crFieldClutenMask) != 0
}

// SetCluten Color Look-Up Table Enable
func (r *RegisterL1crType) SetCluten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL1crFieldClutenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL1crFieldClutenMask)
	}
}

// RegisterL1whpcrType Layerx Window Horizontal Position Configuration Register
type RegisterL1whpcrType uint32

func (r *RegisterL1whpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1whpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1whpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1whpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1whpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1whpcrFieldWhstposShift = 0
	RegisterL1whpcrFieldWhstposMask  = 0xfff
)

// GetWhstpos Window Horizontal Start Position
func (r *RegisterL1whpcrType) GetWhstpos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1whpcrFieldWhstposMask) >> RegisterL1whpcrFieldWhstposShift)
}

// SetWhstpos Window Horizontal Start Position
func (r *RegisterL1whpcrType) SetWhstpos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1whpcrFieldWhstposMask)|(uint32(value)<<RegisterL1whpcrFieldWhstposShift))
}

const (
	RegisterL1whpcrFieldWhspposShift = 16
	RegisterL1whpcrFieldWhspposMask  = 0xfff0000
)

// GetWhsppos Window Horizontal Stop Position
func (r *RegisterL1whpcrType) GetWhsppos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1whpcrFieldWhspposMask) >> RegisterL1whpcrFieldWhspposShift)
}

// SetWhsppos Window Horizontal Stop Position
func (r *RegisterL1whpcrType) SetWhsppos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1whpcrFieldWhspposMask)|(uint32(value)<<RegisterL1whpcrFieldWhspposShift))
}

// RegisterL1wvpcrType Layerx Window Vertical Position Configuration Register
type RegisterL1wvpcrType uint32

func (r *RegisterL1wvpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1wvpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1wvpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1wvpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1wvpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1wvpcrFieldWvstposShift = 0
	RegisterL1wvpcrFieldWvstposMask  = 0x7ff
)

// GetWvstpos Window Vertical Start Position
func (r *RegisterL1wvpcrType) GetWvstpos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1wvpcrFieldWvstposMask) >> RegisterL1wvpcrFieldWvstposShift)
}

// SetWvstpos Window Vertical Start Position
func (r *RegisterL1wvpcrType) SetWvstpos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1wvpcrFieldWvstposMask)|(uint32(value)<<RegisterL1wvpcrFieldWvstposShift))
}

const (
	RegisterL1wvpcrFieldWvspposShift = 16
	RegisterL1wvpcrFieldWvspposMask  = 0x7ff0000
)

// GetWvsppos Window Vertical Stop Position
func (r *RegisterL1wvpcrType) GetWvsppos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1wvpcrFieldWvspposMask) >> RegisterL1wvpcrFieldWvspposShift)
}

// SetWvsppos Window Vertical Stop Position
func (r *RegisterL1wvpcrType) SetWvsppos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1wvpcrFieldWvspposMask)|(uint32(value)<<RegisterL1wvpcrFieldWvspposShift))
}

// RegisterL1ckcrType Layerx Color Keying Configuration Register
type RegisterL1ckcrType uint32

func (r *RegisterL1ckcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1ckcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1ckcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1ckcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1ckcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1ckcrFieldCkblueShift = 0
	RegisterL1ckcrFieldCkblueMask  = 0xff
)

// GetCkblue Color Key Blue value
func (r *RegisterL1ckcrType) GetCkblue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1ckcrFieldCkblueMask) >> RegisterL1ckcrFieldCkblueShift)
}

// SetCkblue Color Key Blue value
func (r *RegisterL1ckcrType) SetCkblue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1ckcrFieldCkblueMask)|(uint32(value)<<RegisterL1ckcrFieldCkblueShift))
}

const (
	RegisterL1ckcrFieldCkgreenShift = 8
	RegisterL1ckcrFieldCkgreenMask  = 0xff00
)

// GetCkgreen Color Key Green value
func (r *RegisterL1ckcrType) GetCkgreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1ckcrFieldCkgreenMask) >> RegisterL1ckcrFieldCkgreenShift)
}

// SetCkgreen Color Key Green value
func (r *RegisterL1ckcrType) SetCkgreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1ckcrFieldCkgreenMask)|(uint32(value)<<RegisterL1ckcrFieldCkgreenShift))
}

const (
	RegisterL1ckcrFieldCkredShift = 16
	RegisterL1ckcrFieldCkredMask  = 0xff0000
)

// GetCkred Color Key Red value
func (r *RegisterL1ckcrType) GetCkred() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1ckcrFieldCkredMask) >> RegisterL1ckcrFieldCkredShift)
}

// SetCkred Color Key Red value
func (r *RegisterL1ckcrType) SetCkred(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1ckcrFieldCkredMask)|(uint32(value)<<RegisterL1ckcrFieldCkredShift))
}

// RegisterL1pfcrType Layerx Pixel Format Configuration Register
type RegisterL1pfcrType uint32

func (r *RegisterL1pfcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1pfcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1pfcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1pfcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1pfcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1pfcrFieldPfShift = 0
	RegisterL1pfcrFieldPfMask  = 0x7
)

// GetPf Pixel Format
func (r *RegisterL1pfcrType) GetPf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1pfcrFieldPfMask) >> RegisterL1pfcrFieldPfShift)
}

// SetPf Pixel Format
func (r *RegisterL1pfcrType) SetPf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1pfcrFieldPfMask)|(uint32(value)<<RegisterL1pfcrFieldPfShift))
}

// RegisterL1cacrType Layerx Constant Alpha Configuration Register
type RegisterL1cacrType uint32

func (r *RegisterL1cacrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1cacrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1cacrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1cacrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1cacrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1cacrFieldConstaShift = 0
	RegisterL1cacrFieldConstaMask  = 0xff
)

// GetConsta Constant Alpha
func (r *RegisterL1cacrType) GetConsta() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1cacrFieldConstaMask) >> RegisterL1cacrFieldConstaShift)
}

// SetConsta Constant Alpha
func (r *RegisterL1cacrType) SetConsta(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1cacrFieldConstaMask)|(uint32(value)<<RegisterL1cacrFieldConstaShift))
}

// RegisterL1dccrType Layerx Default Color Configuration Register
type RegisterL1dccrType uint32

func (r *RegisterL1dccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1dccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1dccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1dccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1dccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1dccrFieldDcblueShift = 0
	RegisterL1dccrFieldDcblueMask  = 0xff
)

// GetDcblue Default Color Blue
func (r *RegisterL1dccrType) GetDcblue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1dccrFieldDcblueMask) >> RegisterL1dccrFieldDcblueShift)
}

// SetDcblue Default Color Blue
func (r *RegisterL1dccrType) SetDcblue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1dccrFieldDcblueMask)|(uint32(value)<<RegisterL1dccrFieldDcblueShift))
}

const (
	RegisterL1dccrFieldDcgreenShift = 8
	RegisterL1dccrFieldDcgreenMask  = 0xff00
)

// GetDcgreen Default Color Green
func (r *RegisterL1dccrType) GetDcgreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1dccrFieldDcgreenMask) >> RegisterL1dccrFieldDcgreenShift)
}

// SetDcgreen Default Color Green
func (r *RegisterL1dccrType) SetDcgreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1dccrFieldDcgreenMask)|(uint32(value)<<RegisterL1dccrFieldDcgreenShift))
}

const (
	RegisterL1dccrFieldDcredShift = 16
	RegisterL1dccrFieldDcredMask  = 0xff0000
)

// GetDcred Default Color Red
func (r *RegisterL1dccrType) GetDcred() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1dccrFieldDcredMask) >> RegisterL1dccrFieldDcredShift)
}

// SetDcred Default Color Red
func (r *RegisterL1dccrType) SetDcred(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1dccrFieldDcredMask)|(uint32(value)<<RegisterL1dccrFieldDcredShift))
}

const (
	RegisterL1dccrFieldDcalphaShift = 24
	RegisterL1dccrFieldDcalphaMask  = 0xff000000
)

// GetDcalpha Default Color Alpha
func (r *RegisterL1dccrType) GetDcalpha() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1dccrFieldDcalphaMask) >> RegisterL1dccrFieldDcalphaShift)
}

// SetDcalpha Default Color Alpha
func (r *RegisterL1dccrType) SetDcalpha(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1dccrFieldDcalphaMask)|(uint32(value)<<RegisterL1dccrFieldDcalphaShift))
}

// RegisterL1bfcrType Layerx Blending Factors Configuration Register
type RegisterL1bfcrType uint32

func (r *RegisterL1bfcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1bfcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1bfcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1bfcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1bfcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1bfcrFieldBf2Shift = 0
	RegisterL1bfcrFieldBf2Mask  = 0x7
)

// GetBf2 Blending Factor 2
func (r *RegisterL1bfcrType) GetBf2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1bfcrFieldBf2Mask) >> RegisterL1bfcrFieldBf2Shift)
}

// SetBf2 Blending Factor 2
func (r *RegisterL1bfcrType) SetBf2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1bfcrFieldBf2Mask)|(uint32(value)<<RegisterL1bfcrFieldBf2Shift))
}

const (
	RegisterL1bfcrFieldBf1Shift = 8
	RegisterL1bfcrFieldBf1Mask  = 0x700
)

// GetBf1 Blending Factor 1
func (r *RegisterL1bfcrType) GetBf1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1bfcrFieldBf1Mask) >> RegisterL1bfcrFieldBf1Shift)
}

// SetBf1 Blending Factor 1
func (r *RegisterL1bfcrType) SetBf1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1bfcrFieldBf1Mask)|(uint32(value)<<RegisterL1bfcrFieldBf1Shift))
}

// RegisterL1cfbarType Layerx Color Frame Buffer Address Register
type RegisterL1cfbarType uint32

func (r *RegisterL1cfbarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1cfbarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1cfbarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1cfbarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1cfbarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1cfbarFieldCfbaddShift = 0
	RegisterL1cfbarFieldCfbaddMask  = 0xffffffff
)

// GetCfbadd Color Frame Buffer Start Address
func (r *RegisterL1cfbarType) GetCfbadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterL1cfbarFieldCfbaddMask) >> RegisterL1cfbarFieldCfbaddShift)
}

// SetCfbadd Color Frame Buffer Start Address
func (r *RegisterL1cfbarType) SetCfbadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1cfbarFieldCfbaddMask)|(uint32(value)<<RegisterL1cfbarFieldCfbaddShift))
}

// RegisterL1cfblrType Layerx Color Frame Buffer Length Register
type RegisterL1cfblrType uint32

func (r *RegisterL1cfblrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1cfblrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1cfblrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1cfblrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1cfblrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1cfblrFieldCfbllShift = 0
	RegisterL1cfblrFieldCfbllMask  = 0x1fff
)

// GetCfbll Color Frame Buffer Line Length
func (r *RegisterL1cfblrType) GetCfbll() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1cfblrFieldCfbllMask) >> RegisterL1cfblrFieldCfbllShift)
}

// SetCfbll Color Frame Buffer Line Length
func (r *RegisterL1cfblrType) SetCfbll(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1cfblrFieldCfbllMask)|(uint32(value)<<RegisterL1cfblrFieldCfbllShift))
}

const (
	RegisterL1cfblrFieldCfbpShift = 16
	RegisterL1cfblrFieldCfbpMask  = 0x1fff0000
)

// GetCfbp Color Frame Buffer Pitch in bytes
func (r *RegisterL1cfblrType) GetCfbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1cfblrFieldCfbpMask) >> RegisterL1cfblrFieldCfbpShift)
}

// SetCfbp Color Frame Buffer Pitch in bytes
func (r *RegisterL1cfblrType) SetCfbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1cfblrFieldCfbpMask)|(uint32(value)<<RegisterL1cfblrFieldCfbpShift))
}

// RegisterL1cfblnrType Layerx ColorFrame Buffer Line Number Register
type RegisterL1cfblnrType uint32

func (r *RegisterL1cfblnrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1cfblnrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1cfblnrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1cfblnrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1cfblnrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1cfblnrFieldCfblnbrShift = 0
	RegisterL1cfblnrFieldCfblnbrMask  = 0x7ff
)

// GetCfblnbr Frame Buffer Line Number
func (r *RegisterL1cfblnrType) GetCfblnbr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL1cfblnrFieldCfblnbrMask) >> RegisterL1cfblnrFieldCfblnbrShift)
}

// SetCfblnbr Frame Buffer Line Number
func (r *RegisterL1cfblnrType) SetCfblnbr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1cfblnrFieldCfblnbrMask)|(uint32(value)<<RegisterL1cfblnrFieldCfblnbrShift))
}

// RegisterL1clutwrType Layerx CLUT Write Register
type RegisterL1clutwrType uint32

func (r *RegisterL1clutwrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL1clutwrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL1clutwrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL1clutwrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL1clutwrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL1clutwrFieldBlueShift = 0
	RegisterL1clutwrFieldBlueMask  = 0xff
)

// GetBlue Blue value
func (r *RegisterL1clutwrType) GetBlue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1clutwrFieldBlueMask) >> RegisterL1clutwrFieldBlueShift)
}

// SetBlue Blue value
func (r *RegisterL1clutwrType) SetBlue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1clutwrFieldBlueMask)|(uint32(value)<<RegisterL1clutwrFieldBlueShift))
}

const (
	RegisterL1clutwrFieldGreenShift = 8
	RegisterL1clutwrFieldGreenMask  = 0xff00
)

// GetGreen Green value
func (r *RegisterL1clutwrType) GetGreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1clutwrFieldGreenMask) >> RegisterL1clutwrFieldGreenShift)
}

// SetGreen Green value
func (r *RegisterL1clutwrType) SetGreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1clutwrFieldGreenMask)|(uint32(value)<<RegisterL1clutwrFieldGreenShift))
}

const (
	RegisterL1clutwrFieldRedShift = 16
	RegisterL1clutwrFieldRedMask  = 0xff0000
)

// GetRed Red value
func (r *RegisterL1clutwrType) GetRed() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1clutwrFieldRedMask) >> RegisterL1clutwrFieldRedShift)
}

// SetRed Red value
func (r *RegisterL1clutwrType) SetRed(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1clutwrFieldRedMask)|(uint32(value)<<RegisterL1clutwrFieldRedShift))
}

const (
	RegisterL1clutwrFieldClutaddShift = 24
	RegisterL1clutwrFieldClutaddMask  = 0xff000000
)

// GetClutadd CLUT Address
func (r *RegisterL1clutwrType) GetClutadd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL1clutwrFieldClutaddMask) >> RegisterL1clutwrFieldClutaddShift)
}

// SetClutadd CLUT Address
func (r *RegisterL1clutwrType) SetClutadd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL1clutwrFieldClutaddMask)|(uint32(value)<<RegisterL1clutwrFieldClutaddShift))
}

// RegisterL2crType Layerx Control Register
type RegisterL2crType uint32

func (r *RegisterL2crType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2crType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2crType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2crType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2crType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2crFieldLenShift = 0
	RegisterL2crFieldLenMask  = 0x1
)

// GetLen Layer Enable
func (r *RegisterL2crType) GetLen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL2crFieldLenMask) != 0
}

// SetLen Layer Enable
func (r *RegisterL2crType) SetLen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL2crFieldLenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL2crFieldLenMask)
	}
}

const (
	RegisterL2crFieldColkenShift = 1
	RegisterL2crFieldColkenMask  = 0x2
)

// GetColken Color Keying Enable
func (r *RegisterL2crType) GetColken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL2crFieldColkenMask) != 0
}

// SetColken Color Keying Enable
func (r *RegisterL2crType) SetColken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL2crFieldColkenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL2crFieldColkenMask)
	}
}

const (
	RegisterL2crFieldClutenShift = 4
	RegisterL2crFieldClutenMask  = 0x10
)

// GetCluten Color Look-Up Table Enable
func (r *RegisterL2crType) GetCluten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterL2crFieldClutenMask) != 0
}

// SetCluten Color Look-Up Table Enable
func (r *RegisterL2crType) SetCluten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterL2crFieldClutenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterL2crFieldClutenMask)
	}
}

// RegisterL2whpcrType Layerx Window Horizontal Position Configuration Register
type RegisterL2whpcrType uint32

func (r *RegisterL2whpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2whpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2whpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2whpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2whpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2whpcrFieldWhstposShift = 0
	RegisterL2whpcrFieldWhstposMask  = 0xfff
)

// GetWhstpos Window Horizontal Start Position
func (r *RegisterL2whpcrType) GetWhstpos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2whpcrFieldWhstposMask) >> RegisterL2whpcrFieldWhstposShift)
}

// SetWhstpos Window Horizontal Start Position
func (r *RegisterL2whpcrType) SetWhstpos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2whpcrFieldWhstposMask)|(uint32(value)<<RegisterL2whpcrFieldWhstposShift))
}

const (
	RegisterL2whpcrFieldWhspposShift = 16
	RegisterL2whpcrFieldWhspposMask  = 0xfff0000
)

// GetWhsppos Window Horizontal Stop Position
func (r *RegisterL2whpcrType) GetWhsppos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2whpcrFieldWhspposMask) >> RegisterL2whpcrFieldWhspposShift)
}

// SetWhsppos Window Horizontal Stop Position
func (r *RegisterL2whpcrType) SetWhsppos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2whpcrFieldWhspposMask)|(uint32(value)<<RegisterL2whpcrFieldWhspposShift))
}

// RegisterL2wvpcrType Layerx Window Vertical Position Configuration Register
type RegisterL2wvpcrType uint32

func (r *RegisterL2wvpcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2wvpcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2wvpcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2wvpcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2wvpcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2wvpcrFieldWvstposShift = 0
	RegisterL2wvpcrFieldWvstposMask  = 0x7ff
)

// GetWvstpos Window Vertical Start Position
func (r *RegisterL2wvpcrType) GetWvstpos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2wvpcrFieldWvstposMask) >> RegisterL2wvpcrFieldWvstposShift)
}

// SetWvstpos Window Vertical Start Position
func (r *RegisterL2wvpcrType) SetWvstpos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2wvpcrFieldWvstposMask)|(uint32(value)<<RegisterL2wvpcrFieldWvstposShift))
}

const (
	RegisterL2wvpcrFieldWvspposShift = 16
	RegisterL2wvpcrFieldWvspposMask  = 0x7ff0000
)

// GetWvsppos Window Vertical Stop Position
func (r *RegisterL2wvpcrType) GetWvsppos() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2wvpcrFieldWvspposMask) >> RegisterL2wvpcrFieldWvspposShift)
}

// SetWvsppos Window Vertical Stop Position
func (r *RegisterL2wvpcrType) SetWvsppos(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2wvpcrFieldWvspposMask)|(uint32(value)<<RegisterL2wvpcrFieldWvspposShift))
}

// RegisterL2ckcrType Layerx Color Keying Configuration Register
type RegisterL2ckcrType uint32

func (r *RegisterL2ckcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2ckcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2ckcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2ckcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2ckcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2ckcrFieldCkblueShift = 0
	RegisterL2ckcrFieldCkblueMask  = 0xff
)

// GetCkblue Color Key Blue value
func (r *RegisterL2ckcrType) GetCkblue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2ckcrFieldCkblueMask) >> RegisterL2ckcrFieldCkblueShift)
}

// SetCkblue Color Key Blue value
func (r *RegisterL2ckcrType) SetCkblue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2ckcrFieldCkblueMask)|(uint32(value)<<RegisterL2ckcrFieldCkblueShift))
}

const (
	RegisterL2ckcrFieldCkgreenShift = 8
	RegisterL2ckcrFieldCkgreenMask  = 0xff00
)

// GetCkgreen Color Key Green value
func (r *RegisterL2ckcrType) GetCkgreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2ckcrFieldCkgreenMask) >> RegisterL2ckcrFieldCkgreenShift)
}

// SetCkgreen Color Key Green value
func (r *RegisterL2ckcrType) SetCkgreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2ckcrFieldCkgreenMask)|(uint32(value)<<RegisterL2ckcrFieldCkgreenShift))
}

const (
	RegisterL2ckcrFieldCkredShift = 16
	RegisterL2ckcrFieldCkredMask  = 0xff0000
)

// GetCkred Color Key Red value
func (r *RegisterL2ckcrType) GetCkred() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2ckcrFieldCkredMask) >> RegisterL2ckcrFieldCkredShift)
}

// SetCkred Color Key Red value
func (r *RegisterL2ckcrType) SetCkred(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2ckcrFieldCkredMask)|(uint32(value)<<RegisterL2ckcrFieldCkredShift))
}

// RegisterL2pfcrType Layerx Pixel Format Configuration Register
type RegisterL2pfcrType uint32

func (r *RegisterL2pfcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2pfcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2pfcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2pfcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2pfcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2pfcrFieldPfShift = 0
	RegisterL2pfcrFieldPfMask  = 0x7
)

// GetPf Pixel Format
func (r *RegisterL2pfcrType) GetPf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2pfcrFieldPfMask) >> RegisterL2pfcrFieldPfShift)
}

// SetPf Pixel Format
func (r *RegisterL2pfcrType) SetPf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2pfcrFieldPfMask)|(uint32(value)<<RegisterL2pfcrFieldPfShift))
}

// RegisterL2cacrType Layerx Constant Alpha Configuration Register
type RegisterL2cacrType uint32

func (r *RegisterL2cacrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2cacrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2cacrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2cacrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2cacrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2cacrFieldConstaShift = 0
	RegisterL2cacrFieldConstaMask  = 0xff
)

// GetConsta Constant Alpha
func (r *RegisterL2cacrType) GetConsta() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2cacrFieldConstaMask) >> RegisterL2cacrFieldConstaShift)
}

// SetConsta Constant Alpha
func (r *RegisterL2cacrType) SetConsta(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2cacrFieldConstaMask)|(uint32(value)<<RegisterL2cacrFieldConstaShift))
}

// RegisterL2dccrType Layerx Default Color Configuration Register
type RegisterL2dccrType uint32

func (r *RegisterL2dccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2dccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2dccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2dccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2dccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2dccrFieldDcblueShift = 0
	RegisterL2dccrFieldDcblueMask  = 0xff
)

// GetDcblue Default Color Blue
func (r *RegisterL2dccrType) GetDcblue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2dccrFieldDcblueMask) >> RegisterL2dccrFieldDcblueShift)
}

// SetDcblue Default Color Blue
func (r *RegisterL2dccrType) SetDcblue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2dccrFieldDcblueMask)|(uint32(value)<<RegisterL2dccrFieldDcblueShift))
}

const (
	RegisterL2dccrFieldDcgreenShift = 8
	RegisterL2dccrFieldDcgreenMask  = 0xff00
)

// GetDcgreen Default Color Green
func (r *RegisterL2dccrType) GetDcgreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2dccrFieldDcgreenMask) >> RegisterL2dccrFieldDcgreenShift)
}

// SetDcgreen Default Color Green
func (r *RegisterL2dccrType) SetDcgreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2dccrFieldDcgreenMask)|(uint32(value)<<RegisterL2dccrFieldDcgreenShift))
}

const (
	RegisterL2dccrFieldDcredShift = 16
	RegisterL2dccrFieldDcredMask  = 0xff0000
)

// GetDcred Default Color Red
func (r *RegisterL2dccrType) GetDcred() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2dccrFieldDcredMask) >> RegisterL2dccrFieldDcredShift)
}

// SetDcred Default Color Red
func (r *RegisterL2dccrType) SetDcred(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2dccrFieldDcredMask)|(uint32(value)<<RegisterL2dccrFieldDcredShift))
}

const (
	RegisterL2dccrFieldDcalphaShift = 24
	RegisterL2dccrFieldDcalphaMask  = 0xff000000
)

// GetDcalpha Default Color Alpha
func (r *RegisterL2dccrType) GetDcalpha() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2dccrFieldDcalphaMask) >> RegisterL2dccrFieldDcalphaShift)
}

// SetDcalpha Default Color Alpha
func (r *RegisterL2dccrType) SetDcalpha(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2dccrFieldDcalphaMask)|(uint32(value)<<RegisterL2dccrFieldDcalphaShift))
}

// RegisterL2bfcrType Layerx Blending Factors Configuration Register
type RegisterL2bfcrType uint32

func (r *RegisterL2bfcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2bfcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2bfcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2bfcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2bfcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2bfcrFieldBf2Shift = 0
	RegisterL2bfcrFieldBf2Mask  = 0x7
)

// GetBf2 Blending Factor 2
func (r *RegisterL2bfcrType) GetBf2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2bfcrFieldBf2Mask) >> RegisterL2bfcrFieldBf2Shift)
}

// SetBf2 Blending Factor 2
func (r *RegisterL2bfcrType) SetBf2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2bfcrFieldBf2Mask)|(uint32(value)<<RegisterL2bfcrFieldBf2Shift))
}

const (
	RegisterL2bfcrFieldBf1Shift = 8
	RegisterL2bfcrFieldBf1Mask  = 0x700
)

// GetBf1 Blending Factor 1
func (r *RegisterL2bfcrType) GetBf1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2bfcrFieldBf1Mask) >> RegisterL2bfcrFieldBf1Shift)
}

// SetBf1 Blending Factor 1
func (r *RegisterL2bfcrType) SetBf1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2bfcrFieldBf1Mask)|(uint32(value)<<RegisterL2bfcrFieldBf1Shift))
}

// RegisterL2cfbarType Layerx Color Frame Buffer Address Register
type RegisterL2cfbarType uint32

func (r *RegisterL2cfbarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2cfbarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2cfbarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2cfbarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2cfbarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2cfbarFieldCfbaddShift = 0
	RegisterL2cfbarFieldCfbaddMask  = 0xffffffff
)

// GetCfbadd Color Frame Buffer Start Address
func (r *RegisterL2cfbarType) GetCfbadd() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterL2cfbarFieldCfbaddMask) >> RegisterL2cfbarFieldCfbaddShift)
}

// SetCfbadd Color Frame Buffer Start Address
func (r *RegisterL2cfbarType) SetCfbadd(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2cfbarFieldCfbaddMask)|(uint32(value)<<RegisterL2cfbarFieldCfbaddShift))
}

// RegisterL2cfblrType Layerx Color Frame Buffer Length Register
type RegisterL2cfblrType uint32

func (r *RegisterL2cfblrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2cfblrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2cfblrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2cfblrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2cfblrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2cfblrFieldCfbllShift = 0
	RegisterL2cfblrFieldCfbllMask  = 0x1fff
)

// GetCfbll Color Frame Buffer Line Length
func (r *RegisterL2cfblrType) GetCfbll() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2cfblrFieldCfbllMask) >> RegisterL2cfblrFieldCfbllShift)
}

// SetCfbll Color Frame Buffer Line Length
func (r *RegisterL2cfblrType) SetCfbll(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2cfblrFieldCfbllMask)|(uint32(value)<<RegisterL2cfblrFieldCfbllShift))
}

const (
	RegisterL2cfblrFieldCfbpShift = 16
	RegisterL2cfblrFieldCfbpMask  = 0x1fff0000
)

// GetCfbp Color Frame Buffer Pitch in bytes
func (r *RegisterL2cfblrType) GetCfbp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2cfblrFieldCfbpMask) >> RegisterL2cfblrFieldCfbpShift)
}

// SetCfbp Color Frame Buffer Pitch in bytes
func (r *RegisterL2cfblrType) SetCfbp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2cfblrFieldCfbpMask)|(uint32(value)<<RegisterL2cfblrFieldCfbpShift))
}

// RegisterL2cfblnrType Layerx ColorFrame Buffer Line Number Register
type RegisterL2cfblnrType uint32

func (r *RegisterL2cfblnrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2cfblnrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2cfblnrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2cfblnrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2cfblnrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2cfblnrFieldCfblnbrShift = 0
	RegisterL2cfblnrFieldCfblnbrMask  = 0x7ff
)

// GetCfblnbr Frame Buffer Line Number
func (r *RegisterL2cfblnrType) GetCfblnbr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterL2cfblnrFieldCfblnbrMask) >> RegisterL2cfblnrFieldCfblnbrShift)
}

// SetCfblnbr Frame Buffer Line Number
func (r *RegisterL2cfblnrType) SetCfblnbr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2cfblnrFieldCfblnbrMask)|(uint32(value)<<RegisterL2cfblnrFieldCfblnbrShift))
}

// RegisterL2clutwrType Layerx CLUT Write Register
type RegisterL2clutwrType uint32

func (r *RegisterL2clutwrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterL2clutwrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterL2clutwrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterL2clutwrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterL2clutwrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterL2clutwrFieldBlueShift = 0
	RegisterL2clutwrFieldBlueMask  = 0xff
)

// GetBlue Blue value
func (r *RegisterL2clutwrType) GetBlue() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2clutwrFieldBlueMask) >> RegisterL2clutwrFieldBlueShift)
}

// SetBlue Blue value
func (r *RegisterL2clutwrType) SetBlue(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2clutwrFieldBlueMask)|(uint32(value)<<RegisterL2clutwrFieldBlueShift))
}

const (
	RegisterL2clutwrFieldGreenShift = 8
	RegisterL2clutwrFieldGreenMask  = 0xff00
)

// GetGreen Green value
func (r *RegisterL2clutwrType) GetGreen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2clutwrFieldGreenMask) >> RegisterL2clutwrFieldGreenShift)
}

// SetGreen Green value
func (r *RegisterL2clutwrType) SetGreen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2clutwrFieldGreenMask)|(uint32(value)<<RegisterL2clutwrFieldGreenShift))
}

const (
	RegisterL2clutwrFieldRedShift = 16
	RegisterL2clutwrFieldRedMask  = 0xff0000
)

// GetRed Red value
func (r *RegisterL2clutwrType) GetRed() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2clutwrFieldRedMask) >> RegisterL2clutwrFieldRedShift)
}

// SetRed Red value
func (r *RegisterL2clutwrType) SetRed(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2clutwrFieldRedMask)|(uint32(value)<<RegisterL2clutwrFieldRedShift))
}

const (
	RegisterL2clutwrFieldClutaddShift = 24
	RegisterL2clutwrFieldClutaddMask  = 0xff000000
)

// GetClutadd CLUT Address
func (r *RegisterL2clutwrType) GetClutadd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterL2clutwrFieldClutaddMask) >> RegisterL2clutwrFieldClutaddShift)
}

// SetClutadd CLUT Address
func (r *RegisterL2clutwrType) SetClutadd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterL2clutwrFieldClutaddMask)|(uint32(value)<<RegisterL2clutwrFieldClutaddShift))
}
