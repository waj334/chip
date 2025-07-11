//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package rcc

import (
	"unsafe"
	"volatile"
)

var (
	Rcc = (*_rcc)(unsafe.Pointer(uintptr(0x58024400)))
)

type _rcc struct {
	Cr           RegisterCrType
	Hsicfgr      RegisterHsicfgrType
	Crrcr        RegisterCrrcrType
	_            [4]byte
	Cfgr         RegisterCfgrType
	_            [4]byte
	D1cfgr       RegisterD1cfgrType
	D2cfgr       RegisterD2cfgrType
	D3cfgr       RegisterD3cfgrType
	_            [4]byte
	Pllckselr    RegisterPllckselrType
	Pllcfgr      RegisterPllcfgrType
	Pll1divr     RegisterPll1divrType
	Pll1fracr    RegisterPll1fracrType
	Pll2divr     RegisterPll2divrType
	Pll2fracr    RegisterPll2fracrType
	Pll3divr     RegisterPll3divrType
	Pll3fracr    RegisterPll3fracrType
	_            [4]byte
	D1ccipr      RegisterD1cciprType
	D2ccip1r     RegisterD2ccip1rType
	D2ccip2r     RegisterD2ccip2rType
	D3ccipr      RegisterD3cciprType
	_            [4]byte
	Cier         RegisterCierType
	Cifr         RegisterCifrType
	Cicr         RegisterCicrType
	_            [4]byte
	Bdcr         RegisterBdcrType
	Csr          RegisterCsrType
	_            [4]byte
	Ahb3rstr     RegisterAhb3rstrType
	Ahb1rstr     RegisterAhb1rstrType
	Ahb2rstr     RegisterAhb2rstrType
	Ahb4rstr     RegisterAhb4rstrType
	Apb3rstr     RegisterApb3rstrType
	Apb1lrstr    RegisterApb1lrstrType
	Apb1hrstr    RegisterApb1hrstrType
	Apb2rstr     RegisterApb2rstrType
	Apb4rstr     RegisterApb4rstrType
	Gcr          RegisterGcrType
	_            [4]byte
	D3amr        RegisterD3amrType
	_            [36]byte
	Rsr          RegisterRsrType
	Ahb3enr      RegisterAhb3enrType
	Ahb1enr      RegisterAhb1enrType
	Ahb2enr      RegisterAhb2enrType
	Ahb4enr      RegisterAhb4enrType
	Apb3enr      RegisterApb3enrType
	Apb1lenr     RegisterApb1lenrType
	Apb1henr     RegisterApb1henrType
	Apb2enr      RegisterApb2enrType
	Apb4enr      RegisterApb4enrType
	_            [4]byte
	Ahb3lpenr    RegisterAhb3lpenrType
	Ahb1lpenr    RegisterAhb1lpenrType
	Ahb2lpenr    RegisterAhb2lpenrType
	Ahb4lpenr    RegisterAhb4lpenrType
	Apb3lpenr    RegisterApb3lpenrType
	Apb1llpenr   RegisterApb1llpenrType
	Apb1hlpenr   RegisterApb1hlpenrType
	Apb2lpenr    RegisterApb2lpenrType
	Apb4lpenr    RegisterApb4lpenrType
	_            [16]byte
	C1rsr        RegisterC1rsrType
	C1ahb3enr    RegisterC1ahb3enrType
	C1ahb1enr    RegisterC1ahb1enrType
	C1ahb2enr    RegisterC1ahb2enrType
	C1ahb4enr    RegisterC1ahb4enrType
	C1apb3enr    RegisterC1apb3enrType
	C1apb1lenr   RegisterC1apb1lenrType
	C1apb1henr   RegisterC1apb1henrType
	C1apb2enr    RegisterC1apb2enrType
	C1apb4enr    RegisterC1apb4enrType
	_            [4]byte
	C1ahb3lpenr  RegisterC1ahb3lpenrType
	C1ahb1lpenr  RegisterC1ahb1lpenrType
	C1ahb2lpenr  RegisterC1ahb2lpenrType
	C1ahb4lpenr  RegisterC1ahb4lpenrType
	C1apb3lpenr  RegisterC1apb3lpenrType
	C1apb1llpenr RegisterC1apb1llpenrType
	C1apb1hlpenr RegisterC1apb1hlpenrType
	C1apb2lpenr  RegisterC1apb2lpenrType
	C1apb4lpenr  RegisterC1apb4lpenrType
}

// RegisterCrType clock control register
type RegisterCrType uint32

func (r *RegisterCrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrFieldHsionShift = 0
	RegisterCrFieldHsionMask  = 0x1
)

// GetHsion Internal high-speed clock enable
func (r *RegisterCrType) GetHsion() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHsionMask) != 0
}

// SetHsion Internal high-speed clock enable
func (r *RegisterCrType) SetHsion(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldHsionMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldHsionMask)
	}
}

const (
	RegisterCrFieldHsikeronShift = 1
	RegisterCrFieldHsikeronMask  = 0x2
)

// GetHsikeron High Speed Internal clock enable in Stop mode
func (r *RegisterCrType) GetHsikeron() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHsikeronMask) != 0
}

// SetHsikeron High Speed Internal clock enable in Stop mode
func (r *RegisterCrType) SetHsikeron(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldHsikeronMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldHsikeronMask)
	}
}

const (
	RegisterCrFieldHsirdyShift = 2
	RegisterCrFieldHsirdyMask  = 0x4
)

// GetHsirdy HSI clock ready flag
func (r *RegisterCrType) GetHsirdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHsirdyMask) != 0
}

// SetHsirdy HSI clock ready flag
func (r *RegisterCrType) SetHsirdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldHsirdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldHsirdyMask)
	}
}

const (
	RegisterCrFieldHsidivShift = 3
	RegisterCrFieldHsidivMask  = 0x18
)

// GetHsidiv HSI clock divider
func (r *RegisterCrType) GetHsidiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHsidivMask) >> RegisterCrFieldHsidivShift)
}

// SetHsidiv HSI clock divider
func (r *RegisterCrType) SetHsidiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldHsidivMask)|(uint32(value)<<RegisterCrFieldHsidivShift))
}

const (
	RegisterCrFieldHsidivfShift = 5
	RegisterCrFieldHsidivfMask  = 0x20
)

// GetHsidivf HSI divider flag
func (r *RegisterCrType) GetHsidivf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHsidivfMask) != 0
}

// SetHsidivf HSI divider flag
func (r *RegisterCrType) SetHsidivf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldHsidivfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldHsidivfMask)
	}
}

const (
	RegisterCrFieldCsionShift = 7
	RegisterCrFieldCsionMask  = 0x80
)

// GetCsion CSI clock enable
func (r *RegisterCrType) GetCsion() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCsionMask) != 0
}

// SetCsion CSI clock enable
func (r *RegisterCrType) SetCsion(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCsionMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCsionMask)
	}
}

const (
	RegisterCrFieldCsirdyShift = 8
	RegisterCrFieldCsirdyMask  = 0x100
)

// GetCsirdy CSI clock ready flag
func (r *RegisterCrType) GetCsirdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCsirdyMask) != 0
}

// SetCsirdy CSI clock ready flag
func (r *RegisterCrType) SetCsirdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCsirdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCsirdyMask)
	}
}

const (
	RegisterCrFieldCsikeronShift = 9
	RegisterCrFieldCsikeronMask  = 0x200
)

// GetCsikeron CSI clock enable in Stop mode
func (r *RegisterCrType) GetCsikeron() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCsikeronMask) != 0
}

// SetCsikeron CSI clock enable in Stop mode
func (r *RegisterCrType) SetCsikeron(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCsikeronMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCsikeronMask)
	}
}

const (
	RegisterCrFieldRc48onShift = 12
	RegisterCrFieldRc48onMask  = 0x1000
)

// GetRc48on RC48 clock enable
func (r *RegisterCrType) GetRc48on() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRc48onMask) != 0
}

// SetRc48on RC48 clock enable
func (r *RegisterCrType) SetRc48on(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRc48onMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRc48onMask)
	}
}

const (
	RegisterCrFieldRc48rdyShift = 13
	RegisterCrFieldRc48rdyMask  = 0x2000
)

// GetRc48rdy RC48 clock ready flag
func (r *RegisterCrType) GetRc48rdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRc48rdyMask) != 0
}

// SetRc48rdy RC48 clock ready flag
func (r *RegisterCrType) SetRc48rdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRc48rdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRc48rdyMask)
	}
}

const (
	RegisterCrFieldD1ckrdyShift = 14
	RegisterCrFieldD1ckrdyMask  = 0x4000
)

// GetD1ckrdy D1 domain clocks ready flag
func (r *RegisterCrType) GetD1ckrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldD1ckrdyMask) != 0
}

// SetD1ckrdy D1 domain clocks ready flag
func (r *RegisterCrType) SetD1ckrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldD1ckrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldD1ckrdyMask)
	}
}

const (
	RegisterCrFieldD2ckrdyShift = 15
	RegisterCrFieldD2ckrdyMask  = 0x8000
)

// GetD2ckrdy D2 domain clocks ready flag
func (r *RegisterCrType) GetD2ckrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldD2ckrdyMask) != 0
}

// SetD2ckrdy D2 domain clocks ready flag
func (r *RegisterCrType) SetD2ckrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldD2ckrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldD2ckrdyMask)
	}
}

const (
	RegisterCrFieldHseonShift = 16
	RegisterCrFieldHseonMask  = 0x10000
)

// GetHseon HSE clock enable
func (r *RegisterCrType) GetHseon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHseonMask) != 0
}

// SetHseon HSE clock enable
func (r *RegisterCrType) SetHseon(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldHseonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldHseonMask)
	}
}

const (
	RegisterCrFieldHserdyShift = 17
	RegisterCrFieldHserdyMask  = 0x20000
)

// GetHserdy HSE clock ready flag
func (r *RegisterCrType) GetHserdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHserdyMask) != 0
}

// SetHserdy HSE clock ready flag
func (r *RegisterCrType) SetHserdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldHserdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldHserdyMask)
	}
}

const (
	RegisterCrFieldHsebypShift = 18
	RegisterCrFieldHsebypMask  = 0x40000
)

// GetHsebyp HSE clock bypass
func (r *RegisterCrType) GetHsebyp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHsebypMask) != 0
}

// SetHsebyp HSE clock bypass
func (r *RegisterCrType) SetHsebyp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldHsebypMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldHsebypMask)
	}
}

const (
	RegisterCrFieldHsecssonShift = 19
	RegisterCrFieldHsecssonMask  = 0x80000
)

// GetHsecsson HSE Clock Security System enable
func (r *RegisterCrType) GetHsecsson() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHsecssonMask) != 0
}

// SetHsecsson HSE Clock Security System enable
func (r *RegisterCrType) SetHsecsson(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldHsecssonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldHsecssonMask)
	}
}

const (
	RegisterCrFieldPll1onShift = 24
	RegisterCrFieldPll1onMask  = 0x1000000
)

// GetPll1on PLL1 enable
func (r *RegisterCrType) GetPll1on() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPll1onMask) != 0
}

// SetPll1on PLL1 enable
func (r *RegisterCrType) SetPll1on(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPll1onMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPll1onMask)
	}
}

const (
	RegisterCrFieldPll1rdyShift = 25
	RegisterCrFieldPll1rdyMask  = 0x2000000
)

// GetPll1rdy PLL1 clock ready flag
func (r *RegisterCrType) GetPll1rdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPll1rdyMask) != 0
}

// SetPll1rdy PLL1 clock ready flag
func (r *RegisterCrType) SetPll1rdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPll1rdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPll1rdyMask)
	}
}

const (
	RegisterCrFieldPll2onShift = 26
	RegisterCrFieldPll2onMask  = 0x4000000
)

// GetPll2on PLL2 enable
func (r *RegisterCrType) GetPll2on() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPll2onMask) != 0
}

// SetPll2on PLL2 enable
func (r *RegisterCrType) SetPll2on(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPll2onMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPll2onMask)
	}
}

const (
	RegisterCrFieldPll2rdyShift = 27
	RegisterCrFieldPll2rdyMask  = 0x8000000
)

// GetPll2rdy PLL2 clock ready flag
func (r *RegisterCrType) GetPll2rdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPll2rdyMask) != 0
}

// SetPll2rdy PLL2 clock ready flag
func (r *RegisterCrType) SetPll2rdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPll2rdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPll2rdyMask)
	}
}

const (
	RegisterCrFieldPll3onShift = 28
	RegisterCrFieldPll3onMask  = 0x10000000
)

// GetPll3on PLL3 enable
func (r *RegisterCrType) GetPll3on() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPll3onMask) != 0
}

// SetPll3on PLL3 enable
func (r *RegisterCrType) SetPll3on(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPll3onMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPll3onMask)
	}
}

const (
	RegisterCrFieldPll3rdyShift = 29
	RegisterCrFieldPll3rdyMask  = 0x20000000
)

// GetPll3rdy PLL3 clock ready flag
func (r *RegisterCrType) GetPll3rdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldPll3rdyMask) != 0
}

// SetPll3rdy PLL3 clock ready flag
func (r *RegisterCrType) SetPll3rdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldPll3rdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldPll3rdyMask)
	}
}

// RegisterHsicfgrType RCC HSI configuration register
type RegisterHsicfgrType uint32

func (r *RegisterHsicfgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterHsicfgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterHsicfgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterHsicfgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterHsicfgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterHsicfgrFieldHsicalShift = 0
	RegisterHsicfgrFieldHsicalMask  = 0xfff
)

// GetHsical HSI clock calibration
func (r *RegisterHsicfgrType) GetHsical() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterHsicfgrFieldHsicalMask) >> RegisterHsicfgrFieldHsicalShift)
}

const (
	RegisterHsicfgrFieldHsitrimShift = 24
	RegisterHsicfgrFieldHsitrimMask  = 0x7f000000
)

// GetHsitrim HSI clock trimming
func (r *RegisterHsicfgrType) GetHsitrim() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterHsicfgrFieldHsitrimMask) >> RegisterHsicfgrFieldHsitrimShift)
}

// SetHsitrim HSI clock trimming
func (r *RegisterHsicfgrType) SetHsitrim(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterHsicfgrFieldHsitrimMask)|(uint32(value)<<RegisterHsicfgrFieldHsitrimShift))
}

// RegisterCrrcrType RCC Clock Recovery RC Register
type RegisterCrrcrType uint32

func (r *RegisterCrrcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCrrcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCrrcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCrrcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCrrcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCrrcrFieldRc48calShift = 0
	RegisterCrrcrFieldRc48calMask  = 0x3ff
)

// GetRc48cal Internal RC 48 MHz clock calibration
func (r *RegisterCrrcrType) GetRc48cal() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCrrcrFieldRc48calMask) >> RegisterCrrcrFieldRc48calShift)
}

// SetRc48cal Internal RC 48 MHz clock calibration
func (r *RegisterCrrcrType) SetRc48cal(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrrcrFieldRc48calMask)|(uint32(value)<<RegisterCrrcrFieldRc48calShift))
}

// RegisterCfgrType RCC Clock Configuration Register
type RegisterCfgrType uint32

func (r *RegisterCfgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

type RegisterCfgrFieldSwEnumType uint8

const (
	// RegisterCfgrFieldSwEnumHsi HSI selected as system clock (hsi_ck) (default after reset)
	RegisterCfgrFieldSwEnumHsi RegisterCfgrFieldSwEnumType = 0x0

	// RegisterCfgrFieldSwEnumCsi CSI selected as system clock (csi_ck)
	RegisterCfgrFieldSwEnumCsi RegisterCfgrFieldSwEnumType = 0x1

	// RegisterCfgrFieldSwEnumHse HSE selected as system clock (hse_ck)
	RegisterCfgrFieldSwEnumHse RegisterCfgrFieldSwEnumType = 0x2

	// RegisterCfgrFieldSwEnumPll1 PLL1 selected as system clock (pll1_p_ck)
	RegisterCfgrFieldSwEnumPll1 RegisterCfgrFieldSwEnumType = 0x3

	RegisterCfgrFieldSwShift = 0
	RegisterCfgrFieldSwMask  = 0x7
)

// GetSw System clock switch
func (r *RegisterCfgrType) GetSw() RegisterCfgrFieldSwEnumType {
	return RegisterCfgrFieldSwEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSwMask) >> RegisterCfgrFieldSwShift)
}

// SetSw System clock switch
func (r *RegisterCfgrType) SetSw(value RegisterCfgrFieldSwEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSwMask)|(uint32(value)<<RegisterCfgrFieldSwShift))
}

type RegisterCfgrFieldSwsEnumType uint8

const (
	// RegisterCfgrFieldSwsEnumHsi HSI selected as system clock (hsi_ck) (default after reset)
	RegisterCfgrFieldSwsEnumHsi RegisterCfgrFieldSwsEnumType = 0x0

	// RegisterCfgrFieldSwsEnumCsi CSI selected as system clock (csi_ck)
	RegisterCfgrFieldSwsEnumCsi RegisterCfgrFieldSwsEnumType = 0x1

	// RegisterCfgrFieldSwsEnumHse HSE selected as system clock (hse_ck)
	RegisterCfgrFieldSwsEnumHse RegisterCfgrFieldSwsEnumType = 0x2

	// RegisterCfgrFieldSwsEnumPll1 PLL1 selected as system clock (pll1_p_ck)
	RegisterCfgrFieldSwsEnumPll1 RegisterCfgrFieldSwsEnumType = 0x3

	RegisterCfgrFieldSwsShift = 3
	RegisterCfgrFieldSwsMask  = 0x38
)

// GetSws System clock switch status
func (r *RegisterCfgrType) GetSws() RegisterCfgrFieldSwsEnumType {
	return RegisterCfgrFieldSwsEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSwsMask) >> RegisterCfgrFieldSwsShift)
}

// SetSws System clock switch status
func (r *RegisterCfgrType) SetSws(value RegisterCfgrFieldSwsEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSwsMask)|(uint32(value)<<RegisterCfgrFieldSwsShift))
}

const (
	RegisterCfgrFieldStopwuckShift = 6
	RegisterCfgrFieldStopwuckMask  = 0x40
)

// GetStopwuck System clock selection after a wake up from system Stop
func (r *RegisterCfgrType) GetStopwuck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldStopwuckMask) != 0
}

// SetStopwuck System clock selection after a wake up from system Stop
func (r *RegisterCfgrType) SetStopwuck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldStopwuckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldStopwuckMask)
	}
}

const (
	RegisterCfgrFieldStopkerwuckShift = 7
	RegisterCfgrFieldStopkerwuckMask  = 0x80
)

// GetStopkerwuck Kernel clock selection after a wake up from system Stop
func (r *RegisterCfgrType) GetStopkerwuck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldStopkerwuckMask) != 0
}

// SetStopkerwuck Kernel clock selection after a wake up from system Stop
func (r *RegisterCfgrType) SetStopkerwuck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldStopkerwuckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldStopkerwuckMask)
	}
}

const (
	RegisterCfgrFieldRtcpreShift = 8
	RegisterCfgrFieldRtcpreMask  = 0x3f00
)

// GetRtcpre HSE division factor for RTC clock
func (r *RegisterCfgrType) GetRtcpre() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldRtcpreMask) >> RegisterCfgrFieldRtcpreShift)
}

// SetRtcpre HSE division factor for RTC clock
func (r *RegisterCfgrType) SetRtcpre(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldRtcpreMask)|(uint32(value)<<RegisterCfgrFieldRtcpreShift))
}

const (
	RegisterCfgrFieldHrtimselShift = 14
	RegisterCfgrFieldHrtimselMask  = 0x4000
)

// GetHrtimsel High Resolution Timer clock prescaler selection
func (r *RegisterCfgrType) GetHrtimsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldHrtimselMask) != 0
}

// SetHrtimsel High Resolution Timer clock prescaler selection
func (r *RegisterCfgrType) SetHrtimsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldHrtimselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldHrtimselMask)
	}
}

const (
	RegisterCfgrFieldTimpreShift = 15
	RegisterCfgrFieldTimpreMask  = 0x8000
)

// GetTimpre Timers clocks prescaler selection
func (r *RegisterCfgrType) GetTimpre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldTimpreMask) != 0
}

// SetTimpre Timers clocks prescaler selection
func (r *RegisterCfgrType) SetTimpre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldTimpreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldTimpreMask)
	}
}

const (
	RegisterCfgrFieldMco1preShift = 18
	RegisterCfgrFieldMco1preMask  = 0x3c0000
)

// GetMco1pre MCO1 prescaler
func (r *RegisterCfgrType) GetMco1pre() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldMco1preMask) >> RegisterCfgrFieldMco1preShift)
}

// SetMco1pre MCO1 prescaler
func (r *RegisterCfgrType) SetMco1pre(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldMco1preMask)|(uint32(value)<<RegisterCfgrFieldMco1preShift))
}

const (
	RegisterCfgrFieldMco1selShift = 22
	RegisterCfgrFieldMco1selMask  = 0x1c00000
)

// GetMco1sel Micro-controller clock output 1
func (r *RegisterCfgrType) GetMco1sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldMco1selMask) >> RegisterCfgrFieldMco1selShift)
}

// SetMco1sel Micro-controller clock output 1
func (r *RegisterCfgrType) SetMco1sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldMco1selMask)|(uint32(value)<<RegisterCfgrFieldMco1selShift))
}

const (
	RegisterCfgrFieldMco2preShift = 25
	RegisterCfgrFieldMco2preMask  = 0x1e000000
)

// GetMco2pre MCO2 prescaler
func (r *RegisterCfgrType) GetMco2pre() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldMco2preMask) >> RegisterCfgrFieldMco2preShift)
}

// SetMco2pre MCO2 prescaler
func (r *RegisterCfgrType) SetMco2pre(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldMco2preMask)|(uint32(value)<<RegisterCfgrFieldMco2preShift))
}

const (
	RegisterCfgrFieldMco2selShift = 29
	RegisterCfgrFieldMco2selMask  = 0xe0000000
)

// GetMco2sel Micro-controller clock output 2
func (r *RegisterCfgrType) GetMco2sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldMco2selMask) >> RegisterCfgrFieldMco2selShift)
}

// SetMco2sel Micro-controller clock output 2
func (r *RegisterCfgrType) SetMco2sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldMco2selMask)|(uint32(value)<<RegisterCfgrFieldMco2selShift))
}

// RegisterD1cfgrType RCC Domain 1 Clock Configuration Register
type RegisterD1cfgrType uint32

func (r *RegisterD1cfgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterD1cfgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterD1cfgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterD1cfgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterD1cfgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

type RegisterD1cfgrFieldHpreEnumType uint8

const (
	// RegisterD1cfgrFieldHpreEnumDiv1 rcc_hclk3 = sys_d1cpre_ck / 1
	RegisterD1cfgrFieldHpreEnumDiv1 RegisterD1cfgrFieldHpreEnumType = 0x7

	// RegisterD1cfgrFieldHpreEnumDiv2 rcc_hclk3 = sys_d1cpre_ck / 2
	RegisterD1cfgrFieldHpreEnumDiv2 RegisterD1cfgrFieldHpreEnumType = 0x8

	// RegisterD1cfgrFieldHpreEnumDiv4 rcc_hclk3 = sys_d1cpre_ck / 4
	RegisterD1cfgrFieldHpreEnumDiv4 RegisterD1cfgrFieldHpreEnumType = 0x9

	// RegisterD1cfgrFieldHpreEnumDiv8 rcc_hclk3 = sys_d1cpre_ck / 8
	RegisterD1cfgrFieldHpreEnumDiv8 RegisterD1cfgrFieldHpreEnumType = 0xa

	// RegisterD1cfgrFieldHpreEnumDiv16 rcc_hclk3 = sys_d1cpre_ck / 16
	RegisterD1cfgrFieldHpreEnumDiv16 RegisterD1cfgrFieldHpreEnumType = 0xb

	// RegisterD1cfgrFieldHpreEnumDiv32 rcc_hclk3 = sys_d1cpre_ck / 32
	RegisterD1cfgrFieldHpreEnumDiv32 RegisterD1cfgrFieldHpreEnumType = 0xc

	// RegisterD1cfgrFieldHpreEnumDiv64 rcc_hclk3 = sys_d1cpre_ck / 64
	RegisterD1cfgrFieldHpreEnumDiv64 RegisterD1cfgrFieldHpreEnumType = 0xd

	// RegisterD1cfgrFieldHpreEnumDiv128 rcc_hclk3 = sys_d1cpre_ck / 128
	RegisterD1cfgrFieldHpreEnumDiv128 RegisterD1cfgrFieldHpreEnumType = 0xe

	// RegisterD1cfgrFieldHpreEnumDiv256 rcc_hclk3 = sys_d1cpre_ck / 256
	RegisterD1cfgrFieldHpreEnumDiv256 RegisterD1cfgrFieldHpreEnumType = 0xf

	// RegisterD1cfgrFieldHpreEnumDiv512 rcc_hclk3 = sys_d1cpre_ck / 512
	RegisterD1cfgrFieldHpreEnumDiv512 RegisterD1cfgrFieldHpreEnumType = 0x10

	RegisterD1cfgrFieldHpreShift = 0
	RegisterD1cfgrFieldHpreMask  = 0xf
)

// GetHpre D1 domain AHB prescaler
func (r *RegisterD1cfgrType) GetHpre() RegisterD1cfgrFieldHpreEnumType {
	return RegisterD1cfgrFieldHpreEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterD1cfgrFieldHpreMask) >> RegisterD1cfgrFieldHpreShift)
}

// SetHpre D1 domain AHB prescaler
func (r *RegisterD1cfgrType) SetHpre(value RegisterD1cfgrFieldHpreEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD1cfgrFieldHpreMask)|(uint32(value)<<RegisterD1cfgrFieldHpreShift))
}

type RegisterD1cfgrFieldD1ppreEnumType uint8

const (
	// RegisterD1cfgrFieldD1ppreEnumDiv1 rcc_pclk3 = rcc_hclk3 / 1
	RegisterD1cfgrFieldD1ppreEnumDiv1 RegisterD1cfgrFieldD1ppreEnumType = 0x3

	// RegisterD1cfgrFieldD1ppreEnumDiv2 rcc_pclk3 = rcc_hclk3 / 2
	RegisterD1cfgrFieldD1ppreEnumDiv2 RegisterD1cfgrFieldD1ppreEnumType = 0x4

	// RegisterD1cfgrFieldD1ppreEnumDiv4 rcc_pclk3 = rcc_hclk3 / 4
	RegisterD1cfgrFieldD1ppreEnumDiv4 RegisterD1cfgrFieldD1ppreEnumType = 0x5

	// RegisterD1cfgrFieldD1ppreEnumDiv8 rcc_pclk3 = rcc_hclk3 / 8
	RegisterD1cfgrFieldD1ppreEnumDiv8 RegisterD1cfgrFieldD1ppreEnumType = 0x6

	// RegisterD1cfgrFieldD1ppreEnumDiv16 rcc_pclk3 = rcc_hclk3 / 16
	RegisterD1cfgrFieldD1ppreEnumDiv16 RegisterD1cfgrFieldD1ppreEnumType = 0x7

	RegisterD1cfgrFieldD1ppreShift = 4
	RegisterD1cfgrFieldD1ppreMask  = 0x70
)

// GetD1ppre D1 domain APB3 prescaler
func (r *RegisterD1cfgrType) GetD1ppre() RegisterD1cfgrFieldD1ppreEnumType {
	return RegisterD1cfgrFieldD1ppreEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterD1cfgrFieldD1ppreMask) >> RegisterD1cfgrFieldD1ppreShift)
}

// SetD1ppre D1 domain APB3 prescaler
func (r *RegisterD1cfgrType) SetD1ppre(value RegisterD1cfgrFieldD1ppreEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD1cfgrFieldD1ppreMask)|(uint32(value)<<RegisterD1cfgrFieldD1ppreShift))
}

type RegisterD1cfgrFieldD1cpreEnumType uint8

const (
	// RegisterD1cfgrFieldD1cpreEnumDiv1 sys_ck / 1
	RegisterD1cfgrFieldD1cpreEnumDiv1 RegisterD1cfgrFieldD1cpreEnumType = 0x7

	// RegisterD1cfgrFieldD1cpreEnumDiv2 sys_ck / 2
	RegisterD1cfgrFieldD1cpreEnumDiv2 RegisterD1cfgrFieldD1cpreEnumType = 0x8

	// RegisterD1cfgrFieldD1cpreEnumDiv4 sys_ck / 4
	RegisterD1cfgrFieldD1cpreEnumDiv4 RegisterD1cfgrFieldD1cpreEnumType = 0x9

	// RegisterD1cfgrFieldD1cpreEnumDiv8 sys_ck / 8
	RegisterD1cfgrFieldD1cpreEnumDiv8 RegisterD1cfgrFieldD1cpreEnumType = 0xa

	// RegisterD1cfgrFieldD1cpreEnumDiv16 sys_ck / 16
	RegisterD1cfgrFieldD1cpreEnumDiv16 RegisterD1cfgrFieldD1cpreEnumType = 0xb

	// RegisterD1cfgrFieldD1cpreEnumDiv32 sys_ck / 32
	RegisterD1cfgrFieldD1cpreEnumDiv32 RegisterD1cfgrFieldD1cpreEnumType = 0xc

	// RegisterD1cfgrFieldD1cpreEnumDiv64 sys_ck / 64
	RegisterD1cfgrFieldD1cpreEnumDiv64 RegisterD1cfgrFieldD1cpreEnumType = 0xd

	// RegisterD1cfgrFieldD1cpreEnumDiv128 sys_ck / 128
	RegisterD1cfgrFieldD1cpreEnumDiv128 RegisterD1cfgrFieldD1cpreEnumType = 0xe

	// RegisterD1cfgrFieldD1cpreEnumDiv256 sys_ck / 256
	RegisterD1cfgrFieldD1cpreEnumDiv256 RegisterD1cfgrFieldD1cpreEnumType = 0xf

	// RegisterD1cfgrFieldD1cpreEnumDiv512 sys_ck / 512
	RegisterD1cfgrFieldD1cpreEnumDiv512 RegisterD1cfgrFieldD1cpreEnumType = 0x10

	RegisterD1cfgrFieldD1cpreShift = 8
	RegisterD1cfgrFieldD1cpreMask  = 0xf00
)

// GetD1cpre D1 domain Core prescaler
func (r *RegisterD1cfgrType) GetD1cpre() RegisterD1cfgrFieldD1cpreEnumType {
	return RegisterD1cfgrFieldD1cpreEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterD1cfgrFieldD1cpreMask) >> RegisterD1cfgrFieldD1cpreShift)
}

// SetD1cpre D1 domain Core prescaler
func (r *RegisterD1cfgrType) SetD1cpre(value RegisterD1cfgrFieldD1cpreEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD1cfgrFieldD1cpreMask)|(uint32(value)<<RegisterD1cfgrFieldD1cpreShift))
}

// RegisterD2cfgrType RCC Domain 2 Clock Configuration Register
type RegisterD2cfgrType uint32

func (r *RegisterD2cfgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterD2cfgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterD2cfgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterD2cfgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterD2cfgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

type RegisterD2cfgrFieldD2ppre1EnumType uint8

const (
	// RegisterD2cfgrFieldD2ppre1EnumDiv1 rcc_pclk2 = rcc_hclk1 / 1
	RegisterD2cfgrFieldD2ppre1EnumDiv1 RegisterD2cfgrFieldD2ppre1EnumType = 0x3

	// RegisterD2cfgrFieldD2ppre1EnumDiv2 rcc_pclk2 = rcc_hclk1 / 2
	RegisterD2cfgrFieldD2ppre1EnumDiv2 RegisterD2cfgrFieldD2ppre1EnumType = 0x4

	// RegisterD2cfgrFieldD2ppre1EnumDiv4 rcc_pclk2 = rcc_hclk1 / 4
	RegisterD2cfgrFieldD2ppre1EnumDiv4 RegisterD2cfgrFieldD2ppre1EnumType = 0x5

	// RegisterD2cfgrFieldD2ppre1EnumDiv8 rcc_pclk2 = rcc_hclk1 / 8
	RegisterD2cfgrFieldD2ppre1EnumDiv8 RegisterD2cfgrFieldD2ppre1EnumType = 0x6

	// RegisterD2cfgrFieldD2ppre1EnumDiv16 rcc_pclk2 = rcc_hclk1 / 16
	RegisterD2cfgrFieldD2ppre1EnumDiv16 RegisterD2cfgrFieldD2ppre1EnumType = 0x7

	RegisterD2cfgrFieldD2ppre1Shift = 4
	RegisterD2cfgrFieldD2ppre1Mask  = 0x70
)

// GetD2ppre1 D2 domain APB1 prescaler
func (r *RegisterD2cfgrType) GetD2ppre1() RegisterD2cfgrFieldD2ppre1EnumType {
	return RegisterD2cfgrFieldD2ppre1EnumType((volatile.LoadUint32((*uint32)(r)) & RegisterD2cfgrFieldD2ppre1Mask) >> RegisterD2cfgrFieldD2ppre1Shift)
}

// SetD2ppre1 D2 domain APB1 prescaler
func (r *RegisterD2cfgrType) SetD2ppre1(value RegisterD2cfgrFieldD2ppre1EnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2cfgrFieldD2ppre1Mask)|(uint32(value)<<RegisterD2cfgrFieldD2ppre1Shift))
}

type RegisterD2cfgrFieldD2ppre2EnumType uint8

const (
	// RegisterD2cfgrFieldD2ppre2EnumDiv1 rcc_pclk1 = rcc_hclk1 / 1
	RegisterD2cfgrFieldD2ppre2EnumDiv1 RegisterD2cfgrFieldD2ppre2EnumType = 0x3

	// RegisterD2cfgrFieldD2ppre2EnumDiv2 rcc_pclk1 = rcc_hclk1 / 2
	RegisterD2cfgrFieldD2ppre2EnumDiv2 RegisterD2cfgrFieldD2ppre2EnumType = 0x4

	// RegisterD2cfgrFieldD2ppre2EnumDiv4 rcc_pclk1 = rcc_hclk1 / 4
	RegisterD2cfgrFieldD2ppre2EnumDiv4 RegisterD2cfgrFieldD2ppre2EnumType = 0x5

	// RegisterD2cfgrFieldD2ppre2EnumDiv8 rcc_pclk1 = rcc_hclk1 / 8
	RegisterD2cfgrFieldD2ppre2EnumDiv8 RegisterD2cfgrFieldD2ppre2EnumType = 0x6

	// RegisterD2cfgrFieldD2ppre2EnumDiv16 rcc_pclk1 = rcc_hclk1 / 16
	RegisterD2cfgrFieldD2ppre2EnumDiv16 RegisterD2cfgrFieldD2ppre2EnumType = 0x7

	RegisterD2cfgrFieldD2ppre2Shift = 8
	RegisterD2cfgrFieldD2ppre2Mask  = 0x700
)

// GetD2ppre2 D2 domain APB2 prescaler
func (r *RegisterD2cfgrType) GetD2ppre2() RegisterD2cfgrFieldD2ppre2EnumType {
	return RegisterD2cfgrFieldD2ppre2EnumType((volatile.LoadUint32((*uint32)(r)) & RegisterD2cfgrFieldD2ppre2Mask) >> RegisterD2cfgrFieldD2ppre2Shift)
}

// SetD2ppre2 D2 domain APB2 prescaler
func (r *RegisterD2cfgrType) SetD2ppre2(value RegisterD2cfgrFieldD2ppre2EnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2cfgrFieldD2ppre2Mask)|(uint32(value)<<RegisterD2cfgrFieldD2ppre2Shift))
}

// RegisterD3cfgrType RCC Domain 3 Clock Configuration Register
type RegisterD3cfgrType uint32

func (r *RegisterD3cfgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterD3cfgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterD3cfgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterD3cfgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterD3cfgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

type RegisterD3cfgrFieldD3ppreEnumType uint8

const (
	// RegisterD3cfgrFieldD3ppreEnumDiv1 rcc_pclk4 = rcc_hclk4 / 1
	RegisterD3cfgrFieldD3ppreEnumDiv1 RegisterD3cfgrFieldD3ppreEnumType = 0x3

	// RegisterD3cfgrFieldD3ppreEnumDiv2 rcc_pclk4 = rcc_hclk4 / 2
	RegisterD3cfgrFieldD3ppreEnumDiv2 RegisterD3cfgrFieldD3ppreEnumType = 0x4

	// RegisterD3cfgrFieldD3ppreEnumDiv4 rcc_pclk4 = rcc_hclk4 / 4
	RegisterD3cfgrFieldD3ppreEnumDiv4 RegisterD3cfgrFieldD3ppreEnumType = 0x5

	// RegisterD3cfgrFieldD3ppreEnumDiv8 rcc_pclk4 = rcc_hclk4 / 8
	RegisterD3cfgrFieldD3ppreEnumDiv8 RegisterD3cfgrFieldD3ppreEnumType = 0x6

	// RegisterD3cfgrFieldD3ppreEnumDiv16 rcc_pclk4 = rcc_hclk4 / 16
	RegisterD3cfgrFieldD3ppreEnumDiv16 RegisterD3cfgrFieldD3ppreEnumType = 0x7

	RegisterD3cfgrFieldD3ppreShift = 4
	RegisterD3cfgrFieldD3ppreMask  = 0x70
)

// GetD3ppre D3 domain APB4 prescaler
func (r *RegisterD3cfgrType) GetD3ppre() RegisterD3cfgrFieldD3ppreEnumType {
	return RegisterD3cfgrFieldD3ppreEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterD3cfgrFieldD3ppreMask) >> RegisterD3cfgrFieldD3ppreShift)
}

// SetD3ppre D3 domain APB4 prescaler
func (r *RegisterD3cfgrType) SetD3ppre(value RegisterD3cfgrFieldD3ppreEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3cfgrFieldD3ppreMask)|(uint32(value)<<RegisterD3cfgrFieldD3ppreShift))
}

// RegisterPllckselrType RCC PLLs Clock Source Selection Register
type RegisterPllckselrType uint32

func (r *RegisterPllckselrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPllckselrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPllckselrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPllckselrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPllckselrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

type RegisterPllckselrFieldPllsrcEnumType uint8

const (
	// RegisterPllckselrFieldPllsrcEnumHsi HSI selected as PLL clock (hsi_ck) (default after reset)
	RegisterPllckselrFieldPllsrcEnumHsi RegisterPllckselrFieldPllsrcEnumType = 0x0

	// RegisterPllckselrFieldPllsrcEnumCsi CSI selected as PLL clock (csi_ck)
	RegisterPllckselrFieldPllsrcEnumCsi RegisterPllckselrFieldPllsrcEnumType = 0x1

	// RegisterPllckselrFieldPllsrcEnumHse HSE selected as PLL clock (hse_ck)
	RegisterPllckselrFieldPllsrcEnumHse RegisterPllckselrFieldPllsrcEnumType = 0x2

	// RegisterPllckselrFieldPllsrcEnumNoclock No clock send to DIVMx divider and PLLs
	RegisterPllckselrFieldPllsrcEnumNoclock RegisterPllckselrFieldPllsrcEnumType = 0x3

	RegisterPllckselrFieldPllsrcShift = 0
	RegisterPllckselrFieldPllsrcMask  = 0x3
)

// GetPllsrc DIVMx and PLLs clock source selection
func (r *RegisterPllckselrType) GetPllsrc() RegisterPllckselrFieldPllsrcEnumType {
	return RegisterPllckselrFieldPllsrcEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterPllckselrFieldPllsrcMask) >> RegisterPllckselrFieldPllsrcShift)
}

// SetPllsrc DIVMx and PLLs clock source selection
func (r *RegisterPllckselrType) SetPllsrc(value RegisterPllckselrFieldPllsrcEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPllckselrFieldPllsrcMask)|(uint32(value)<<RegisterPllckselrFieldPllsrcShift))
}

const (
	RegisterPllckselrFieldDivm1Shift = 4
	RegisterPllckselrFieldDivm1Mask  = 0x3f0
)

// GetDivm1 Prescaler for PLL1
func (r *RegisterPllckselrType) GetDivm1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPllckselrFieldDivm1Mask) >> RegisterPllckselrFieldDivm1Shift)
}

// SetDivm1 Prescaler for PLL1
func (r *RegisterPllckselrType) SetDivm1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPllckselrFieldDivm1Mask)|(uint32(value)<<RegisterPllckselrFieldDivm1Shift))
}

const (
	RegisterPllckselrFieldDivm2Shift = 12
	RegisterPllckselrFieldDivm2Mask  = 0x3f000
)

// GetDivm2 Prescaler for PLL2
func (r *RegisterPllckselrType) GetDivm2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPllckselrFieldDivm2Mask) >> RegisterPllckselrFieldDivm2Shift)
}

// SetDivm2 Prescaler for PLL2
func (r *RegisterPllckselrType) SetDivm2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPllckselrFieldDivm2Mask)|(uint32(value)<<RegisterPllckselrFieldDivm2Shift))
}

const (
	RegisterPllckselrFieldDivm3Shift = 20
	RegisterPllckselrFieldDivm3Mask  = 0x3f00000
)

// GetDivm3 Prescaler for PLL3
func (r *RegisterPllckselrType) GetDivm3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPllckselrFieldDivm3Mask) >> RegisterPllckselrFieldDivm3Shift)
}

// SetDivm3 Prescaler for PLL3
func (r *RegisterPllckselrType) SetDivm3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPllckselrFieldDivm3Mask)|(uint32(value)<<RegisterPllckselrFieldDivm3Shift))
}

// RegisterPllcfgrType RCC PLLs Configuration Register
type RegisterPllcfgrType uint32

func (r *RegisterPllcfgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPllcfgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPllcfgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPllcfgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPllcfgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPllcfgrFieldPll1fracenShift = 0
	RegisterPllcfgrFieldPll1fracenMask  = 0x1
)

// GetPll1fracen PLL1 fractional latch enable
func (r *RegisterPllcfgrType) GetPll1fracen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldPll1fracenMask) != 0
}

// SetPll1fracen PLL1 fractional latch enable
func (r *RegisterPllcfgrType) SetPll1fracen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldPll1fracenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldPll1fracenMask)
	}
}

const (
	RegisterPllcfgrFieldPll1vcoselShift = 1
	RegisterPllcfgrFieldPll1vcoselMask  = 0x2
)

// GetPll1vcosel PLL1 VCO selection
func (r *RegisterPllcfgrType) GetPll1vcosel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldPll1vcoselMask) != 0
}

// SetPll1vcosel PLL1 VCO selection
func (r *RegisterPllcfgrType) SetPll1vcosel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldPll1vcoselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldPll1vcoselMask)
	}
}

type RegisterPllcfgrFieldPll1rgeEnumType uint8

const (
	// RegisterPllcfgrFieldPll1rgeEnum1to2mhz The PLLn input (refn_ck) clock range frequency is between 1 and 2 MHz (default after reset)
	RegisterPllcfgrFieldPll1rgeEnum1to2mhz RegisterPllcfgrFieldPll1rgeEnumType = 0x0

	// RegisterPllcfgrFieldPll1rgeEnum2to4mhz The PLLn input (refn_ck) clock range frequency is between 2 and 4 MHz
	RegisterPllcfgrFieldPll1rgeEnum2to4mhz RegisterPllcfgrFieldPll1rgeEnumType = 0x1

	// RegisterPllcfgrFieldPll1rgeEnum4to8mhz The PLLn input (refn_ck) clock range frequency is between 4 and 8 MHz
	RegisterPllcfgrFieldPll1rgeEnum4to8mhz RegisterPllcfgrFieldPll1rgeEnumType = 0x2

	// RegisterPllcfgrFieldPll1rgeEnum8to16mhz The PLLn input (refn_ck) clock range frequency is between 8 and 16 MHz
	RegisterPllcfgrFieldPll1rgeEnum8to16mhz RegisterPllcfgrFieldPll1rgeEnumType = 0x3

	RegisterPllcfgrFieldPll1rgeShift = 2
	RegisterPllcfgrFieldPll1rgeMask  = 0xc
)

// GetPll1rge PLL1 input frequency range
func (r *RegisterPllcfgrType) GetPll1rge() RegisterPllcfgrFieldPll1rgeEnumType {
	return RegisterPllcfgrFieldPll1rgeEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldPll1rgeMask) >> RegisterPllcfgrFieldPll1rgeShift)
}

// SetPll1rge PLL1 input frequency range
func (r *RegisterPllcfgrType) SetPll1rge(value RegisterPllcfgrFieldPll1rgeEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldPll1rgeMask)|(uint32(value)<<RegisterPllcfgrFieldPll1rgeShift))
}

const (
	RegisterPllcfgrFieldPll2fracenShift = 4
	RegisterPllcfgrFieldPll2fracenMask  = 0x10
)

// GetPll2fracen PLL2 fractional latch enable
func (r *RegisterPllcfgrType) GetPll2fracen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldPll2fracenMask) != 0
}

// SetPll2fracen PLL2 fractional latch enable
func (r *RegisterPllcfgrType) SetPll2fracen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldPll2fracenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldPll2fracenMask)
	}
}

const (
	RegisterPllcfgrFieldPll2vcoselShift = 5
	RegisterPllcfgrFieldPll2vcoselMask  = 0x20
)

// GetPll2vcosel PLL2 VCO selection
func (r *RegisterPllcfgrType) GetPll2vcosel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldPll2vcoselMask) != 0
}

// SetPll2vcosel PLL2 VCO selection
func (r *RegisterPllcfgrType) SetPll2vcosel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldPll2vcoselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldPll2vcoselMask)
	}
}

type RegisterPllcfgrFieldPll2rgeEnumType uint8

const (
	// RegisterPllcfgrFieldPll2rgeEnum1to2mhz The PLLn input (refn_ck) clock range frequency is between 1 and 2 MHz (default after reset)
	RegisterPllcfgrFieldPll2rgeEnum1to2mhz RegisterPllcfgrFieldPll2rgeEnumType = 0x0

	// RegisterPllcfgrFieldPll2rgeEnum2to4mhz The PLLn input (refn_ck) clock range frequency is between 2 and 4 MHz
	RegisterPllcfgrFieldPll2rgeEnum2to4mhz RegisterPllcfgrFieldPll2rgeEnumType = 0x1

	// RegisterPllcfgrFieldPll2rgeEnum4to8mhz The PLLn input (refn_ck) clock range frequency is between 4 and 8 MHz
	RegisterPllcfgrFieldPll2rgeEnum4to8mhz RegisterPllcfgrFieldPll2rgeEnumType = 0x2

	// RegisterPllcfgrFieldPll2rgeEnum8to16mhz The PLLn input (refn_ck) clock range frequency is between 8 and 16 MHz
	RegisterPllcfgrFieldPll2rgeEnum8to16mhz RegisterPllcfgrFieldPll2rgeEnumType = 0x3

	RegisterPllcfgrFieldPll2rgeShift = 6
	RegisterPllcfgrFieldPll2rgeMask  = 0xc0
)

// GetPll2rge PLL2 input frequency range
func (r *RegisterPllcfgrType) GetPll2rge() RegisterPllcfgrFieldPll2rgeEnumType {
	return RegisterPllcfgrFieldPll2rgeEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldPll2rgeMask) >> RegisterPllcfgrFieldPll2rgeShift)
}

// SetPll2rge PLL2 input frequency range
func (r *RegisterPllcfgrType) SetPll2rge(value RegisterPllcfgrFieldPll2rgeEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldPll2rgeMask)|(uint32(value)<<RegisterPllcfgrFieldPll2rgeShift))
}

const (
	RegisterPllcfgrFieldPll3fracenShift = 8
	RegisterPllcfgrFieldPll3fracenMask  = 0x100
)

// GetPll3fracen PLL3 fractional latch enable
func (r *RegisterPllcfgrType) GetPll3fracen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldPll3fracenMask) != 0
}

// SetPll3fracen PLL3 fractional latch enable
func (r *RegisterPllcfgrType) SetPll3fracen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldPll3fracenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldPll3fracenMask)
	}
}

const (
	RegisterPllcfgrFieldPll3vcoselShift = 9
	RegisterPllcfgrFieldPll3vcoselMask  = 0x200
)

// GetPll3vcosel PLL3 VCO selection
func (r *RegisterPllcfgrType) GetPll3vcosel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldPll3vcoselMask) != 0
}

// SetPll3vcosel PLL3 VCO selection
func (r *RegisterPllcfgrType) SetPll3vcosel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldPll3vcoselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldPll3vcoselMask)
	}
}

type RegisterPllcfgrFieldPll3rgeEnumType uint8

const (
	// RegisterPllcfgrFieldPll3rgeEnum1to2mhz The PLLn input (refn_ck) clock range frequency is between 1 and 2 MHz (default after reset)
	RegisterPllcfgrFieldPll3rgeEnum1to2mhz RegisterPllcfgrFieldPll3rgeEnumType = 0x0

	// RegisterPllcfgrFieldPll3rgeEnum2to4mhz The PLLn input (refn_ck) clock range frequency is between 2 and 4 MHz
	RegisterPllcfgrFieldPll3rgeEnum2to4mhz RegisterPllcfgrFieldPll3rgeEnumType = 0x1

	// RegisterPllcfgrFieldPll3rgeEnum4to8mhz The PLLn input (refn_ck) clock range frequency is between 4 and 8 MHz
	RegisterPllcfgrFieldPll3rgeEnum4to8mhz RegisterPllcfgrFieldPll3rgeEnumType = 0x2

	// RegisterPllcfgrFieldPll3rgeEnum8to16mhz The PLLn input (refn_ck) clock range frequency is between 8 and 16 MHz
	RegisterPllcfgrFieldPll3rgeEnum8to16mhz RegisterPllcfgrFieldPll3rgeEnumType = 0x3

	RegisterPllcfgrFieldPll3rgeShift = 10
	RegisterPllcfgrFieldPll3rgeMask  = 0xc00
)

// GetPll3rge PLL3 input frequency range
func (r *RegisterPllcfgrType) GetPll3rge() RegisterPllcfgrFieldPll3rgeEnumType {
	return RegisterPllcfgrFieldPll3rgeEnumType((volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldPll3rgeMask) >> RegisterPllcfgrFieldPll3rgeShift)
}

// SetPll3rge PLL3 input frequency range
func (r *RegisterPllcfgrType) SetPll3rge(value RegisterPllcfgrFieldPll3rgeEnumType) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldPll3rgeMask)|(uint32(value)<<RegisterPllcfgrFieldPll3rgeShift))
}

const (
	RegisterPllcfgrFieldDivp1enShift = 16
	RegisterPllcfgrFieldDivp1enMask  = 0x10000
)

// GetDivp1en PLL1 DIVP divider output enable
func (r *RegisterPllcfgrType) GetDivp1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldDivp1enMask) != 0
}

// SetDivp1en PLL1 DIVP divider output enable
func (r *RegisterPllcfgrType) SetDivp1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldDivp1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldDivp1enMask)
	}
}

const (
	RegisterPllcfgrFieldDivq1enShift = 17
	RegisterPllcfgrFieldDivq1enMask  = 0x20000
)

// GetDivq1en PLL1 DIVQ divider output enable
func (r *RegisterPllcfgrType) GetDivq1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldDivq1enMask) != 0
}

// SetDivq1en PLL1 DIVQ divider output enable
func (r *RegisterPllcfgrType) SetDivq1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldDivq1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldDivq1enMask)
	}
}

const (
	RegisterPllcfgrFieldDivr1enShift = 18
	RegisterPllcfgrFieldDivr1enMask  = 0x40000
)

// GetDivr1en PLL1 DIVR divider output enable
func (r *RegisterPllcfgrType) GetDivr1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldDivr1enMask) != 0
}

// SetDivr1en PLL1 DIVR divider output enable
func (r *RegisterPllcfgrType) SetDivr1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldDivr1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldDivr1enMask)
	}
}

const (
	RegisterPllcfgrFieldDivp2enShift = 19
	RegisterPllcfgrFieldDivp2enMask  = 0x80000
)

// GetDivp2en PLL2 DIVP divider output enable
func (r *RegisterPllcfgrType) GetDivp2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldDivp2enMask) != 0
}

// SetDivp2en PLL2 DIVP divider output enable
func (r *RegisterPllcfgrType) SetDivp2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldDivp2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldDivp2enMask)
	}
}

const (
	RegisterPllcfgrFieldDivq2enShift = 20
	RegisterPllcfgrFieldDivq2enMask  = 0x100000
)

// GetDivq2en PLL2 DIVQ divider output enable
func (r *RegisterPllcfgrType) GetDivq2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldDivq2enMask) != 0
}

// SetDivq2en PLL2 DIVQ divider output enable
func (r *RegisterPllcfgrType) SetDivq2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldDivq2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldDivq2enMask)
	}
}

const (
	RegisterPllcfgrFieldDivr2enShift = 21
	RegisterPllcfgrFieldDivr2enMask  = 0x200000
)

// GetDivr2en PLL2 DIVR divider output enable
func (r *RegisterPllcfgrType) GetDivr2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldDivr2enMask) != 0
}

// SetDivr2en PLL2 DIVR divider output enable
func (r *RegisterPllcfgrType) SetDivr2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldDivr2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldDivr2enMask)
	}
}

const (
	RegisterPllcfgrFieldDivp3enShift = 22
	RegisterPllcfgrFieldDivp3enMask  = 0x400000
)

// GetDivp3en PLL3 DIVP divider output enable
func (r *RegisterPllcfgrType) GetDivp3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldDivp3enMask) != 0
}

// SetDivp3en PLL3 DIVP divider output enable
func (r *RegisterPllcfgrType) SetDivp3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldDivp3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldDivp3enMask)
	}
}

const (
	RegisterPllcfgrFieldDivq3enShift = 23
	RegisterPllcfgrFieldDivq3enMask  = 0x800000
)

// GetDivq3en PLL3 DIVQ divider output enable
func (r *RegisterPllcfgrType) GetDivq3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldDivq3enMask) != 0
}

// SetDivq3en PLL3 DIVQ divider output enable
func (r *RegisterPllcfgrType) SetDivq3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldDivq3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldDivq3enMask)
	}
}

const (
	RegisterPllcfgrFieldDivr3enShift = 24
	RegisterPllcfgrFieldDivr3enMask  = 0x1000000
)

// GetDivr3en PLL3 DIVR divider output enable
func (r *RegisterPllcfgrType) GetDivr3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPllcfgrFieldDivr3enMask) != 0
}

// SetDivr3en PLL3 DIVR divider output enable
func (r *RegisterPllcfgrType) SetDivr3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPllcfgrFieldDivr3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPllcfgrFieldDivr3enMask)
	}
}

// RegisterPll1divrType RCC PLL1 Dividers Configuration Register
type RegisterPll1divrType uint32

func (r *RegisterPll1divrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPll1divrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPll1divrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPll1divrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPll1divrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPll1divrFieldDivn1Shift = 0
	RegisterPll1divrFieldDivn1Mask  = 0x1ff
)

// GetDivn1 Multiplication factor for PLL1 VCO
func (r *RegisterPll1divrType) GetDivn1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPll1divrFieldDivn1Mask) >> RegisterPll1divrFieldDivn1Shift)
}

// SetDivn1 Multiplication factor for PLL1 VCO
func (r *RegisterPll1divrType) SetDivn1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll1divrFieldDivn1Mask)|(uint32(value)<<RegisterPll1divrFieldDivn1Shift))
}

const (
	RegisterPll1divrFieldDivp1Shift = 9
	RegisterPll1divrFieldDivp1Mask  = 0xfe00
)

// GetDivp1 PLL1 DIVP division factor
func (r *RegisterPll1divrType) GetDivp1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPll1divrFieldDivp1Mask) >> RegisterPll1divrFieldDivp1Shift)
}

// SetDivp1 PLL1 DIVP division factor
func (r *RegisterPll1divrType) SetDivp1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll1divrFieldDivp1Mask)|(uint32(value)<<RegisterPll1divrFieldDivp1Shift))
}

const (
	RegisterPll1divrFieldDivq1Shift = 16
	RegisterPll1divrFieldDivq1Mask  = 0x7f0000
)

// GetDivq1 PLL1 DIVQ division factor
func (r *RegisterPll1divrType) GetDivq1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPll1divrFieldDivq1Mask) >> RegisterPll1divrFieldDivq1Shift)
}

// SetDivq1 PLL1 DIVQ division factor
func (r *RegisterPll1divrType) SetDivq1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll1divrFieldDivq1Mask)|(uint32(value)<<RegisterPll1divrFieldDivq1Shift))
}

const (
	RegisterPll1divrFieldDivr1Shift = 24
	RegisterPll1divrFieldDivr1Mask  = 0x7f000000
)

// GetDivr1 PLL1 DIVR division factor
func (r *RegisterPll1divrType) GetDivr1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPll1divrFieldDivr1Mask) >> RegisterPll1divrFieldDivr1Shift)
}

// SetDivr1 PLL1 DIVR division factor
func (r *RegisterPll1divrType) SetDivr1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll1divrFieldDivr1Mask)|(uint32(value)<<RegisterPll1divrFieldDivr1Shift))
}

// RegisterPll1fracrType RCC PLL1 Fractional Divider Register
type RegisterPll1fracrType uint32

func (r *RegisterPll1fracrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPll1fracrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPll1fracrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPll1fracrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPll1fracrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPll1fracrFieldFracn1Shift = 3
	RegisterPll1fracrFieldFracn1Mask  = 0xfff8
)

// GetFracn1 Fractional part of the multiplication factor for PLL1 VCO
func (r *RegisterPll1fracrType) GetFracn1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPll1fracrFieldFracn1Mask) >> RegisterPll1fracrFieldFracn1Shift)
}

// SetFracn1 Fractional part of the multiplication factor for PLL1 VCO
func (r *RegisterPll1fracrType) SetFracn1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll1fracrFieldFracn1Mask)|(uint32(value)<<RegisterPll1fracrFieldFracn1Shift))
}

// RegisterPll2divrType RCC PLL2 Dividers Configuration Register
type RegisterPll2divrType uint32

func (r *RegisterPll2divrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPll2divrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPll2divrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPll2divrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPll2divrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPll2divrFieldDivn2Shift = 0
	RegisterPll2divrFieldDivn2Mask  = 0x1ff
)

// GetDivn2 Multiplication factor for PLL1 VCO
func (r *RegisterPll2divrType) GetDivn2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPll2divrFieldDivn2Mask) >> RegisterPll2divrFieldDivn2Shift)
}

// SetDivn2 Multiplication factor for PLL1 VCO
func (r *RegisterPll2divrType) SetDivn2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll2divrFieldDivn2Mask)|(uint32(value)<<RegisterPll2divrFieldDivn2Shift))
}

const (
	RegisterPll2divrFieldDivp2Shift = 9
	RegisterPll2divrFieldDivp2Mask  = 0xfe00
)

// GetDivp2 PLL1 DIVP division factor
func (r *RegisterPll2divrType) GetDivp2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPll2divrFieldDivp2Mask) >> RegisterPll2divrFieldDivp2Shift)
}

// SetDivp2 PLL1 DIVP division factor
func (r *RegisterPll2divrType) SetDivp2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll2divrFieldDivp2Mask)|(uint32(value)<<RegisterPll2divrFieldDivp2Shift))
}

const (
	RegisterPll2divrFieldDivq2Shift = 16
	RegisterPll2divrFieldDivq2Mask  = 0x7f0000
)

// GetDivq2 PLL1 DIVQ division factor
func (r *RegisterPll2divrType) GetDivq2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPll2divrFieldDivq2Mask) >> RegisterPll2divrFieldDivq2Shift)
}

// SetDivq2 PLL1 DIVQ division factor
func (r *RegisterPll2divrType) SetDivq2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll2divrFieldDivq2Mask)|(uint32(value)<<RegisterPll2divrFieldDivq2Shift))
}

const (
	RegisterPll2divrFieldDivr2Shift = 24
	RegisterPll2divrFieldDivr2Mask  = 0x7f000000
)

// GetDivr2 PLL1 DIVR division factor
func (r *RegisterPll2divrType) GetDivr2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPll2divrFieldDivr2Mask) >> RegisterPll2divrFieldDivr2Shift)
}

// SetDivr2 PLL1 DIVR division factor
func (r *RegisterPll2divrType) SetDivr2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll2divrFieldDivr2Mask)|(uint32(value)<<RegisterPll2divrFieldDivr2Shift))
}

// RegisterPll2fracrType RCC PLL2 Fractional Divider Register
type RegisterPll2fracrType uint32

func (r *RegisterPll2fracrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPll2fracrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPll2fracrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPll2fracrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPll2fracrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPll2fracrFieldFracn2Shift = 3
	RegisterPll2fracrFieldFracn2Mask  = 0xfff8
)

// GetFracn2 Fractional part of the multiplication factor for PLL VCO
func (r *RegisterPll2fracrType) GetFracn2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPll2fracrFieldFracn2Mask) >> RegisterPll2fracrFieldFracn2Shift)
}

// SetFracn2 Fractional part of the multiplication factor for PLL VCO
func (r *RegisterPll2fracrType) SetFracn2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll2fracrFieldFracn2Mask)|(uint32(value)<<RegisterPll2fracrFieldFracn2Shift))
}

// RegisterPll3divrType RCC PLL3 Dividers Configuration Register
type RegisterPll3divrType uint32

func (r *RegisterPll3divrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPll3divrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPll3divrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPll3divrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPll3divrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPll3divrFieldDivn3Shift = 0
	RegisterPll3divrFieldDivn3Mask  = 0x1ff
)

// GetDivn3 Multiplication factor for PLL1 VCO
func (r *RegisterPll3divrType) GetDivn3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPll3divrFieldDivn3Mask) >> RegisterPll3divrFieldDivn3Shift)
}

// SetDivn3 Multiplication factor for PLL1 VCO
func (r *RegisterPll3divrType) SetDivn3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll3divrFieldDivn3Mask)|(uint32(value)<<RegisterPll3divrFieldDivn3Shift))
}

const (
	RegisterPll3divrFieldDivp3Shift = 9
	RegisterPll3divrFieldDivp3Mask  = 0xfe00
)

// GetDivp3 PLL DIVP division factor
func (r *RegisterPll3divrType) GetDivp3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPll3divrFieldDivp3Mask) >> RegisterPll3divrFieldDivp3Shift)
}

// SetDivp3 PLL DIVP division factor
func (r *RegisterPll3divrType) SetDivp3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll3divrFieldDivp3Mask)|(uint32(value)<<RegisterPll3divrFieldDivp3Shift))
}

const (
	RegisterPll3divrFieldDivq3Shift = 16
	RegisterPll3divrFieldDivq3Mask  = 0x7f0000
)

// GetDivq3 PLL DIVQ division factor
func (r *RegisterPll3divrType) GetDivq3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPll3divrFieldDivq3Mask) >> RegisterPll3divrFieldDivq3Shift)
}

// SetDivq3 PLL DIVQ division factor
func (r *RegisterPll3divrType) SetDivq3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll3divrFieldDivq3Mask)|(uint32(value)<<RegisterPll3divrFieldDivq3Shift))
}

const (
	RegisterPll3divrFieldDivr3Shift = 24
	RegisterPll3divrFieldDivr3Mask  = 0x7f000000
)

// GetDivr3 PLL DIVR division factor
func (r *RegisterPll3divrType) GetDivr3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPll3divrFieldDivr3Mask) >> RegisterPll3divrFieldDivr3Shift)
}

// SetDivr3 PLL DIVR division factor
func (r *RegisterPll3divrType) SetDivr3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll3divrFieldDivr3Mask)|(uint32(value)<<RegisterPll3divrFieldDivr3Shift))
}

// RegisterPll3fracrType RCC PLL3 Fractional Divider Register
type RegisterPll3fracrType uint32

func (r *RegisterPll3fracrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPll3fracrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPll3fracrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPll3fracrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPll3fracrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPll3fracrFieldFracn3Shift = 3
	RegisterPll3fracrFieldFracn3Mask  = 0xfff8
)

// GetFracn3 Fractional part of the multiplication factor for PLL3 VCO
func (r *RegisterPll3fracrType) GetFracn3() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPll3fracrFieldFracn3Mask) >> RegisterPll3fracrFieldFracn3Shift)
}

// SetFracn3 Fractional part of the multiplication factor for PLL3 VCO
func (r *RegisterPll3fracrType) SetFracn3(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPll3fracrFieldFracn3Mask)|(uint32(value)<<RegisterPll3fracrFieldFracn3Shift))
}

// RegisterD1cciprType RCC Domain 1 Kernel Clock Configuration Register
type RegisterD1cciprType uint32

func (r *RegisterD1cciprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterD1cciprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterD1cciprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterD1cciprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterD1cciprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterD1cciprFieldFmcsrcShift = 0
	RegisterD1cciprFieldFmcsrcMask  = 0x3
)

// GetFmcsrc FMC kernel clock source selection
func (r *RegisterD1cciprType) GetFmcsrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD1cciprFieldFmcsrcMask) >> RegisterD1cciprFieldFmcsrcShift)
}

// SetFmcsrc FMC kernel clock source selection
func (r *RegisterD1cciprType) SetFmcsrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD1cciprFieldFmcsrcMask)|(uint32(value)<<RegisterD1cciprFieldFmcsrcShift))
}

const (
	RegisterD1cciprFieldQspisrcShift = 4
	RegisterD1cciprFieldQspisrcMask  = 0x30
)

// GetQspisrc QUADSPI kernel clock source selection
func (r *RegisterD1cciprType) GetQspisrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD1cciprFieldQspisrcMask) >> RegisterD1cciprFieldQspisrcShift)
}

// SetQspisrc QUADSPI kernel clock source selection
func (r *RegisterD1cciprType) SetQspisrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD1cciprFieldQspisrcMask)|(uint32(value)<<RegisterD1cciprFieldQspisrcShift))
}

const (
	RegisterD1cciprFieldSdmmcsrcShift = 16
	RegisterD1cciprFieldSdmmcsrcMask  = 0x10000
)

// GetSdmmcsrc SDMMC kernel clock source selection
func (r *RegisterD1cciprType) GetSdmmcsrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD1cciprFieldSdmmcsrcMask) != 0
}

// SetSdmmcsrc SDMMC kernel clock source selection
func (r *RegisterD1cciprType) SetSdmmcsrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD1cciprFieldSdmmcsrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD1cciprFieldSdmmcsrcMask)
	}
}

const (
	RegisterD1cciprFieldCkpersrcShift = 28
	RegisterD1cciprFieldCkpersrcMask  = 0x30000000
)

// GetCkpersrc per_ck clock source selection
func (r *RegisterD1cciprType) GetCkpersrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD1cciprFieldCkpersrcMask) >> RegisterD1cciprFieldCkpersrcShift)
}

// SetCkpersrc per_ck clock source selection
func (r *RegisterD1cciprType) SetCkpersrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD1cciprFieldCkpersrcMask)|(uint32(value)<<RegisterD1cciprFieldCkpersrcShift))
}

// RegisterD2ccip1rType RCC Domain 2 Kernel Clock Configuration Register
type RegisterD2ccip1rType uint32

func (r *RegisterD2ccip1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterD2ccip1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterD2ccip1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterD2ccip1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterD2ccip1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterD2ccip1rFieldSai1srcShift = 0
	RegisterD2ccip1rFieldSai1srcMask  = 0x7
)

// GetSai1src SAI1 and DFSDM1 kernel Aclk clock source selection
func (r *RegisterD2ccip1rType) GetSai1src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip1rFieldSai1srcMask) >> RegisterD2ccip1rFieldSai1srcShift)
}

// SetSai1src SAI1 and DFSDM1 kernel Aclk clock source selection
func (r *RegisterD2ccip1rType) SetSai1src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip1rFieldSai1srcMask)|(uint32(value)<<RegisterD2ccip1rFieldSai1srcShift))
}

const (
	RegisterD2ccip1rFieldSai23srcShift = 6
	RegisterD2ccip1rFieldSai23srcMask  = 0x1c0
)

// GetSai23src SAI2 and SAI3 kernel clock source selection
func (r *RegisterD2ccip1rType) GetSai23src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip1rFieldSai23srcMask) >> RegisterD2ccip1rFieldSai23srcShift)
}

// SetSai23src SAI2 and SAI3 kernel clock source selection
func (r *RegisterD2ccip1rType) SetSai23src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip1rFieldSai23srcMask)|(uint32(value)<<RegisterD2ccip1rFieldSai23srcShift))
}

const (
	RegisterD2ccip1rFieldSpi123srcShift = 12
	RegisterD2ccip1rFieldSpi123srcMask  = 0x7000
)

// GetSpi123src SPI/I2S1,2 and 3 kernel clock source selection
func (r *RegisterD2ccip1rType) GetSpi123src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip1rFieldSpi123srcMask) >> RegisterD2ccip1rFieldSpi123srcShift)
}

// SetSpi123src SPI/I2S1,2 and 3 kernel clock source selection
func (r *RegisterD2ccip1rType) SetSpi123src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip1rFieldSpi123srcMask)|(uint32(value)<<RegisterD2ccip1rFieldSpi123srcShift))
}

const (
	RegisterD2ccip1rFieldSpi45srcShift = 16
	RegisterD2ccip1rFieldSpi45srcMask  = 0x70000
)

// GetSpi45src SPI4 and 5 kernel clock source selection
func (r *RegisterD2ccip1rType) GetSpi45src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip1rFieldSpi45srcMask) >> RegisterD2ccip1rFieldSpi45srcShift)
}

// SetSpi45src SPI4 and 5 kernel clock source selection
func (r *RegisterD2ccip1rType) SetSpi45src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip1rFieldSpi45srcMask)|(uint32(value)<<RegisterD2ccip1rFieldSpi45srcShift))
}

const (
	RegisterD2ccip1rFieldSpdifsrcShift = 20
	RegisterD2ccip1rFieldSpdifsrcMask  = 0x300000
)

// GetSpdifsrc SPDIFRX kernel clock source selection
func (r *RegisterD2ccip1rType) GetSpdifsrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip1rFieldSpdifsrcMask) >> RegisterD2ccip1rFieldSpdifsrcShift)
}

// SetSpdifsrc SPDIFRX kernel clock source selection
func (r *RegisterD2ccip1rType) SetSpdifsrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip1rFieldSpdifsrcMask)|(uint32(value)<<RegisterD2ccip1rFieldSpdifsrcShift))
}

const (
	RegisterD2ccip1rFieldDfsdm1srcShift = 24
	RegisterD2ccip1rFieldDfsdm1srcMask  = 0x1000000
)

// GetDfsdm1src DFSDM1 kernel Clk clock source selection
func (r *RegisterD2ccip1rType) GetDfsdm1src() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip1rFieldDfsdm1srcMask) != 0
}

// SetDfsdm1src DFSDM1 kernel Clk clock source selection
func (r *RegisterD2ccip1rType) SetDfsdm1src(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD2ccip1rFieldDfsdm1srcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip1rFieldDfsdm1srcMask)
	}
}

const (
	RegisterD2ccip1rFieldFdcansrcShift = 28
	RegisterD2ccip1rFieldFdcansrcMask  = 0x30000000
)

// GetFdcansrc FDCAN kernel clock source selection
func (r *RegisterD2ccip1rType) GetFdcansrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip1rFieldFdcansrcMask) >> RegisterD2ccip1rFieldFdcansrcShift)
}

// SetFdcansrc FDCAN kernel clock source selection
func (r *RegisterD2ccip1rType) SetFdcansrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip1rFieldFdcansrcMask)|(uint32(value)<<RegisterD2ccip1rFieldFdcansrcShift))
}

const (
	RegisterD2ccip1rFieldSwpsrcShift = 31
	RegisterD2ccip1rFieldSwpsrcMask  = 0x80000000
)

// GetSwpsrc SWPMI kernel clock source selection
func (r *RegisterD2ccip1rType) GetSwpsrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip1rFieldSwpsrcMask) != 0
}

// SetSwpsrc SWPMI kernel clock source selection
func (r *RegisterD2ccip1rType) SetSwpsrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD2ccip1rFieldSwpsrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip1rFieldSwpsrcMask)
	}
}

// RegisterD2ccip2rType RCC Domain 2 Kernel Clock Configuration Register
type RegisterD2ccip2rType uint32

func (r *RegisterD2ccip2rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterD2ccip2rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterD2ccip2rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterD2ccip2rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterD2ccip2rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterD2ccip2rFieldUsart234578srcShift = 0
	RegisterD2ccip2rFieldUsart234578srcMask  = 0x7
)

// GetUsart234578src USART2/3, UART4,5, 7/8 (APB1) kernel clock source selection
func (r *RegisterD2ccip2rType) GetUsart234578src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip2rFieldUsart234578srcMask) >> RegisterD2ccip2rFieldUsart234578srcShift)
}

// SetUsart234578src USART2/3, UART4,5, 7/8 (APB1) kernel clock source selection
func (r *RegisterD2ccip2rType) SetUsart234578src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip2rFieldUsart234578srcMask)|(uint32(value)<<RegisterD2ccip2rFieldUsart234578srcShift))
}

const (
	RegisterD2ccip2rFieldUsart16srcShift = 3
	RegisterD2ccip2rFieldUsart16srcMask  = 0x38
)

// GetUsart16src USART1 and 6 kernel clock source selection
func (r *RegisterD2ccip2rType) GetUsart16src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip2rFieldUsart16srcMask) >> RegisterD2ccip2rFieldUsart16srcShift)
}

// SetUsart16src USART1 and 6 kernel clock source selection
func (r *RegisterD2ccip2rType) SetUsart16src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip2rFieldUsart16srcMask)|(uint32(value)<<RegisterD2ccip2rFieldUsart16srcShift))
}

const (
	RegisterD2ccip2rFieldRngsrcShift = 8
	RegisterD2ccip2rFieldRngsrcMask  = 0x300
)

// GetRngsrc RNG kernel clock source selection
func (r *RegisterD2ccip2rType) GetRngsrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip2rFieldRngsrcMask) >> RegisterD2ccip2rFieldRngsrcShift)
}

// SetRngsrc RNG kernel clock source selection
func (r *RegisterD2ccip2rType) SetRngsrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip2rFieldRngsrcMask)|(uint32(value)<<RegisterD2ccip2rFieldRngsrcShift))
}

const (
	RegisterD2ccip2rFieldI2c123srcShift = 12
	RegisterD2ccip2rFieldI2c123srcMask  = 0x3000
)

// GetI2c123src I2C1,2,3 kernel clock source selection
func (r *RegisterD2ccip2rType) GetI2c123src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip2rFieldI2c123srcMask) >> RegisterD2ccip2rFieldI2c123srcShift)
}

// SetI2c123src I2C1,2,3 kernel clock source selection
func (r *RegisterD2ccip2rType) SetI2c123src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip2rFieldI2c123srcMask)|(uint32(value)<<RegisterD2ccip2rFieldI2c123srcShift))
}

const (
	RegisterD2ccip2rFieldUsbsrcShift = 20
	RegisterD2ccip2rFieldUsbsrcMask  = 0x300000
)

// GetUsbsrc USBOTG 1 and 2 kernel clock source selection
func (r *RegisterD2ccip2rType) GetUsbsrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip2rFieldUsbsrcMask) >> RegisterD2ccip2rFieldUsbsrcShift)
}

// SetUsbsrc USBOTG 1 and 2 kernel clock source selection
func (r *RegisterD2ccip2rType) SetUsbsrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip2rFieldUsbsrcMask)|(uint32(value)<<RegisterD2ccip2rFieldUsbsrcShift))
}

const (
	RegisterD2ccip2rFieldCecsrcShift = 22
	RegisterD2ccip2rFieldCecsrcMask  = 0xc00000
)

// GetCecsrc HDMI-CEC kernel clock source selection
func (r *RegisterD2ccip2rType) GetCecsrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip2rFieldCecsrcMask) >> RegisterD2ccip2rFieldCecsrcShift)
}

// SetCecsrc HDMI-CEC kernel clock source selection
func (r *RegisterD2ccip2rType) SetCecsrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip2rFieldCecsrcMask)|(uint32(value)<<RegisterD2ccip2rFieldCecsrcShift))
}

const (
	RegisterD2ccip2rFieldLptim1srcShift = 28
	RegisterD2ccip2rFieldLptim1srcMask  = 0x70000000
)

// GetLptim1src LPTIM1 kernel clock source selection
func (r *RegisterD2ccip2rType) GetLptim1src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD2ccip2rFieldLptim1srcMask) >> RegisterD2ccip2rFieldLptim1srcShift)
}

// SetLptim1src LPTIM1 kernel clock source selection
func (r *RegisterD2ccip2rType) SetLptim1src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD2ccip2rFieldLptim1srcMask)|(uint32(value)<<RegisterD2ccip2rFieldLptim1srcShift))
}

// RegisterD3cciprType RCC Domain 3 Kernel Clock Configuration Register
type RegisterD3cciprType uint32

func (r *RegisterD3cciprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterD3cciprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterD3cciprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterD3cciprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterD3cciprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterD3cciprFieldLpuart1srcShift = 0
	RegisterD3cciprFieldLpuart1srcMask  = 0x7
)

// GetLpuart1src LPUART1 kernel clock source selection
func (r *RegisterD3cciprType) GetLpuart1src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3cciprFieldLpuart1srcMask) >> RegisterD3cciprFieldLpuart1srcShift)
}

// SetLpuart1src LPUART1 kernel clock source selection
func (r *RegisterD3cciprType) SetLpuart1src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3cciprFieldLpuart1srcMask)|(uint32(value)<<RegisterD3cciprFieldLpuart1srcShift))
}

const (
	RegisterD3cciprFieldI2c4srcShift = 8
	RegisterD3cciprFieldI2c4srcMask  = 0x300
)

// GetI2c4src I2C4 kernel clock source selection
func (r *RegisterD3cciprType) GetI2c4src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3cciprFieldI2c4srcMask) >> RegisterD3cciprFieldI2c4srcShift)
}

// SetI2c4src I2C4 kernel clock source selection
func (r *RegisterD3cciprType) SetI2c4src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3cciprFieldI2c4srcMask)|(uint32(value)<<RegisterD3cciprFieldI2c4srcShift))
}

const (
	RegisterD3cciprFieldLptim2srcShift = 10
	RegisterD3cciprFieldLptim2srcMask  = 0x1c00
)

// GetLptim2src LPTIM2 kernel clock source selection
func (r *RegisterD3cciprType) GetLptim2src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3cciprFieldLptim2srcMask) >> RegisterD3cciprFieldLptim2srcShift)
}

// SetLptim2src LPTIM2 kernel clock source selection
func (r *RegisterD3cciprType) SetLptim2src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3cciprFieldLptim2srcMask)|(uint32(value)<<RegisterD3cciprFieldLptim2srcShift))
}

const (
	RegisterD3cciprFieldLptim345srcShift = 13
	RegisterD3cciprFieldLptim345srcMask  = 0xe000
)

// GetLptim345src LPTIM3,4,5 kernel clock source selection
func (r *RegisterD3cciprType) GetLptim345src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3cciprFieldLptim345srcMask) >> RegisterD3cciprFieldLptim345srcShift)
}

// SetLptim345src LPTIM3,4,5 kernel clock source selection
func (r *RegisterD3cciprType) SetLptim345src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3cciprFieldLptim345srcMask)|(uint32(value)<<RegisterD3cciprFieldLptim345srcShift))
}

const (
	RegisterD3cciprFieldAdcsrcShift = 16
	RegisterD3cciprFieldAdcsrcMask  = 0x30000
)

// GetAdcsrc SAR ADC kernel clock source selection
func (r *RegisterD3cciprType) GetAdcsrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3cciprFieldAdcsrcMask) >> RegisterD3cciprFieldAdcsrcShift)
}

// SetAdcsrc SAR ADC kernel clock source selection
func (r *RegisterD3cciprType) SetAdcsrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3cciprFieldAdcsrcMask)|(uint32(value)<<RegisterD3cciprFieldAdcsrcShift))
}

const (
	RegisterD3cciprFieldSai4asrcShift = 21
	RegisterD3cciprFieldSai4asrcMask  = 0xe00000
)

// GetSai4asrc Sub-Block A of SAI4 kernel clock source selection
func (r *RegisterD3cciprType) GetSai4asrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3cciprFieldSai4asrcMask) >> RegisterD3cciprFieldSai4asrcShift)
}

// SetSai4asrc Sub-Block A of SAI4 kernel clock source selection
func (r *RegisterD3cciprType) SetSai4asrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3cciprFieldSai4asrcMask)|(uint32(value)<<RegisterD3cciprFieldSai4asrcShift))
}

const (
	RegisterD3cciprFieldSai4bsrcShift = 24
	RegisterD3cciprFieldSai4bsrcMask  = 0x7000000
)

// GetSai4bsrc Sub-Block B of SAI4 kernel clock source selection
func (r *RegisterD3cciprType) GetSai4bsrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3cciprFieldSai4bsrcMask) >> RegisterD3cciprFieldSai4bsrcShift)
}

// SetSai4bsrc Sub-Block B of SAI4 kernel clock source selection
func (r *RegisterD3cciprType) SetSai4bsrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3cciprFieldSai4bsrcMask)|(uint32(value)<<RegisterD3cciprFieldSai4bsrcShift))
}

const (
	RegisterD3cciprFieldSpi6srcShift = 28
	RegisterD3cciprFieldSpi6srcMask  = 0x70000000
)

// GetSpi6src SPI6 kernel clock source selection
func (r *RegisterD3cciprType) GetSpi6src() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterD3cciprFieldSpi6srcMask) >> RegisterD3cciprFieldSpi6srcShift)
}

// SetSpi6src SPI6 kernel clock source selection
func (r *RegisterD3cciprType) SetSpi6src(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterD3cciprFieldSpi6srcMask)|(uint32(value)<<RegisterD3cciprFieldSpi6srcShift))
}

// RegisterCierType RCC Clock Source Interrupt Enable Register
type RegisterCierType uint32

func (r *RegisterCierType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCierType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCierType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCierType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCierType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCierFieldLsirdyieShift = 0
	RegisterCierFieldLsirdyieMask  = 0x1
)

// GetLsirdyie LSI ready Interrupt Enable
func (r *RegisterCierType) GetLsirdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCierFieldLsirdyieMask) != 0
}

// SetLsirdyie LSI ready Interrupt Enable
func (r *RegisterCierType) SetLsirdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCierFieldLsirdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCierFieldLsirdyieMask)
	}
}

const (
	RegisterCierFieldLserdyieShift = 1
	RegisterCierFieldLserdyieMask  = 0x2
)

// GetLserdyie LSE ready Interrupt Enable
func (r *RegisterCierType) GetLserdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCierFieldLserdyieMask) != 0
}

// SetLserdyie LSE ready Interrupt Enable
func (r *RegisterCierType) SetLserdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCierFieldLserdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCierFieldLserdyieMask)
	}
}

const (
	RegisterCierFieldHsirdyieShift = 2
	RegisterCierFieldHsirdyieMask  = 0x4
)

// GetHsirdyie HSI ready Interrupt Enable
func (r *RegisterCierType) GetHsirdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCierFieldHsirdyieMask) != 0
}

// SetHsirdyie HSI ready Interrupt Enable
func (r *RegisterCierType) SetHsirdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCierFieldHsirdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCierFieldHsirdyieMask)
	}
}

const (
	RegisterCierFieldHserdyieShift = 3
	RegisterCierFieldHserdyieMask  = 0x8
)

// GetHserdyie HSE ready Interrupt Enable
func (r *RegisterCierType) GetHserdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCierFieldHserdyieMask) != 0
}

// SetHserdyie HSE ready Interrupt Enable
func (r *RegisterCierType) SetHserdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCierFieldHserdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCierFieldHserdyieMask)
	}
}

const (
	RegisterCierFieldCsirdyieShift = 4
	RegisterCierFieldCsirdyieMask  = 0x10
)

// GetCsirdyie CSI ready Interrupt Enable
func (r *RegisterCierType) GetCsirdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCierFieldCsirdyieMask) != 0
}

// SetCsirdyie CSI ready Interrupt Enable
func (r *RegisterCierType) SetCsirdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCierFieldCsirdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCierFieldCsirdyieMask)
	}
}

const (
	RegisterCierFieldRc48rdyieShift = 5
	RegisterCierFieldRc48rdyieMask  = 0x20
)

// GetRc48rdyie RC48 ready Interrupt Enable
func (r *RegisterCierType) GetRc48rdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCierFieldRc48rdyieMask) != 0
}

// SetRc48rdyie RC48 ready Interrupt Enable
func (r *RegisterCierType) SetRc48rdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCierFieldRc48rdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCierFieldRc48rdyieMask)
	}
}

const (
	RegisterCierFieldPll1rdyieShift = 6
	RegisterCierFieldPll1rdyieMask  = 0x40
)

// GetPll1rdyie PLL1 ready Interrupt Enable
func (r *RegisterCierType) GetPll1rdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCierFieldPll1rdyieMask) != 0
}

// SetPll1rdyie PLL1 ready Interrupt Enable
func (r *RegisterCierType) SetPll1rdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCierFieldPll1rdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCierFieldPll1rdyieMask)
	}
}

const (
	RegisterCierFieldPll2rdyieShift = 7
	RegisterCierFieldPll2rdyieMask  = 0x80
)

// GetPll2rdyie PLL2 ready Interrupt Enable
func (r *RegisterCierType) GetPll2rdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCierFieldPll2rdyieMask) != 0
}

// SetPll2rdyie PLL2 ready Interrupt Enable
func (r *RegisterCierType) SetPll2rdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCierFieldPll2rdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCierFieldPll2rdyieMask)
	}
}

const (
	RegisterCierFieldPll3rdyieShift = 8
	RegisterCierFieldPll3rdyieMask  = 0x100
)

// GetPll3rdyie PLL3 ready Interrupt Enable
func (r *RegisterCierType) GetPll3rdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCierFieldPll3rdyieMask) != 0
}

// SetPll3rdyie PLL3 ready Interrupt Enable
func (r *RegisterCierType) SetPll3rdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCierFieldPll3rdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCierFieldPll3rdyieMask)
	}
}

const (
	RegisterCierFieldLsecssieShift = 9
	RegisterCierFieldLsecssieMask  = 0x200
)

// GetLsecssie LSE clock security system Interrupt Enable
func (r *RegisterCierType) GetLsecssie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCierFieldLsecssieMask) != 0
}

// SetLsecssie LSE clock security system Interrupt Enable
func (r *RegisterCierType) SetLsecssie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCierFieldLsecssieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCierFieldLsecssieMask)
	}
}

// RegisterCifrType RCC Clock Source Interrupt Flag Register
type RegisterCifrType uint32

func (r *RegisterCifrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCifrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCifrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCifrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCifrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCifrFieldLsirdyfShift = 0
	RegisterCifrFieldLsirdyfMask  = 0x1
)

// GetLsirdyf LSI ready Interrupt Flag
func (r *RegisterCifrType) GetLsirdyf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCifrFieldLsirdyfMask) != 0
}

// SetLsirdyf LSI ready Interrupt Flag
func (r *RegisterCifrType) SetLsirdyf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCifrFieldLsirdyfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCifrFieldLsirdyfMask)
	}
}

const (
	RegisterCifrFieldLserdyfShift = 1
	RegisterCifrFieldLserdyfMask  = 0x2
)

// GetLserdyf LSE ready Interrupt Flag
func (r *RegisterCifrType) GetLserdyf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCifrFieldLserdyfMask) != 0
}

// SetLserdyf LSE ready Interrupt Flag
func (r *RegisterCifrType) SetLserdyf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCifrFieldLserdyfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCifrFieldLserdyfMask)
	}
}

const (
	RegisterCifrFieldHsirdyfShift = 2
	RegisterCifrFieldHsirdyfMask  = 0x4
)

// GetHsirdyf HSI ready Interrupt Flag
func (r *RegisterCifrType) GetHsirdyf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCifrFieldHsirdyfMask) != 0
}

// SetHsirdyf HSI ready Interrupt Flag
func (r *RegisterCifrType) SetHsirdyf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCifrFieldHsirdyfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCifrFieldHsirdyfMask)
	}
}

const (
	RegisterCifrFieldHserdyfShift = 3
	RegisterCifrFieldHserdyfMask  = 0x8
)

// GetHserdyf HSE ready Interrupt Flag
func (r *RegisterCifrType) GetHserdyf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCifrFieldHserdyfMask) != 0
}

// SetHserdyf HSE ready Interrupt Flag
func (r *RegisterCifrType) SetHserdyf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCifrFieldHserdyfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCifrFieldHserdyfMask)
	}
}

const (
	RegisterCifrFieldCsirdyShift = 4
	RegisterCifrFieldCsirdyMask  = 0x10
)

// GetCsirdy CSI ready Interrupt Flag
func (r *RegisterCifrType) GetCsirdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCifrFieldCsirdyMask) != 0
}

// SetCsirdy CSI ready Interrupt Flag
func (r *RegisterCifrType) SetCsirdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCifrFieldCsirdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCifrFieldCsirdyMask)
	}
}

const (
	RegisterCifrFieldRc48rdyfShift = 5
	RegisterCifrFieldRc48rdyfMask  = 0x20
)

// GetRc48rdyf RC48 ready Interrupt Flag
func (r *RegisterCifrType) GetRc48rdyf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCifrFieldRc48rdyfMask) != 0
}

// SetRc48rdyf RC48 ready Interrupt Flag
func (r *RegisterCifrType) SetRc48rdyf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCifrFieldRc48rdyfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCifrFieldRc48rdyfMask)
	}
}

const (
	RegisterCifrFieldPll1rdyfShift = 6
	RegisterCifrFieldPll1rdyfMask  = 0x40
)

// GetPll1rdyf PLL1 ready Interrupt Flag
func (r *RegisterCifrType) GetPll1rdyf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCifrFieldPll1rdyfMask) != 0
}

// SetPll1rdyf PLL1 ready Interrupt Flag
func (r *RegisterCifrType) SetPll1rdyf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCifrFieldPll1rdyfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCifrFieldPll1rdyfMask)
	}
}

const (
	RegisterCifrFieldPll2rdyfShift = 7
	RegisterCifrFieldPll2rdyfMask  = 0x80
)

// GetPll2rdyf PLL2 ready Interrupt Flag
func (r *RegisterCifrType) GetPll2rdyf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCifrFieldPll2rdyfMask) != 0
}

// SetPll2rdyf PLL2 ready Interrupt Flag
func (r *RegisterCifrType) SetPll2rdyf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCifrFieldPll2rdyfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCifrFieldPll2rdyfMask)
	}
}

const (
	RegisterCifrFieldPll3rdyfShift = 8
	RegisterCifrFieldPll3rdyfMask  = 0x100
)

// GetPll3rdyf PLL3 ready Interrupt Flag
func (r *RegisterCifrType) GetPll3rdyf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCifrFieldPll3rdyfMask) != 0
}

// SetPll3rdyf PLL3 ready Interrupt Flag
func (r *RegisterCifrType) SetPll3rdyf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCifrFieldPll3rdyfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCifrFieldPll3rdyfMask)
	}
}

const (
	RegisterCifrFieldLsecssfShift = 9
	RegisterCifrFieldLsecssfMask  = 0x200
)

// GetLsecssf LSE clock security system Interrupt Flag
func (r *RegisterCifrType) GetLsecssf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCifrFieldLsecssfMask) != 0
}

// SetLsecssf LSE clock security system Interrupt Flag
func (r *RegisterCifrType) SetLsecssf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCifrFieldLsecssfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCifrFieldLsecssfMask)
	}
}

const (
	RegisterCifrFieldHsecssfShift = 10
	RegisterCifrFieldHsecssfMask  = 0x400
)

// GetHsecssf HSE clock security system Interrupt Flag
func (r *RegisterCifrType) GetHsecssf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCifrFieldHsecssfMask) != 0
}

// SetHsecssf HSE clock security system Interrupt Flag
func (r *RegisterCifrType) SetHsecssf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCifrFieldHsecssfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCifrFieldHsecssfMask)
	}
}

// RegisterCicrType RCC Clock Source Interrupt Clear Register
type RegisterCicrType uint32

func (r *RegisterCicrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCicrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCicrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCicrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCicrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCicrFieldLsirdycShift = 0
	RegisterCicrFieldLsirdycMask  = 0x1
)

// GetLsirdyc LSI ready Interrupt Clear
func (r *RegisterCicrType) GetLsirdyc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCicrFieldLsirdycMask) != 0
}

// SetLsirdyc LSI ready Interrupt Clear
func (r *RegisterCicrType) SetLsirdyc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCicrFieldLsirdycMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCicrFieldLsirdycMask)
	}
}

const (
	RegisterCicrFieldLserdycShift = 1
	RegisterCicrFieldLserdycMask  = 0x2
)

// GetLserdyc LSE ready Interrupt Clear
func (r *RegisterCicrType) GetLserdyc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCicrFieldLserdycMask) != 0
}

// SetLserdyc LSE ready Interrupt Clear
func (r *RegisterCicrType) SetLserdyc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCicrFieldLserdycMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCicrFieldLserdycMask)
	}
}

const (
	RegisterCicrFieldHsirdycShift = 2
	RegisterCicrFieldHsirdycMask  = 0x4
)

// GetHsirdyc HSI ready Interrupt Clear
func (r *RegisterCicrType) GetHsirdyc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCicrFieldHsirdycMask) != 0
}

// SetHsirdyc HSI ready Interrupt Clear
func (r *RegisterCicrType) SetHsirdyc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCicrFieldHsirdycMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCicrFieldHsirdycMask)
	}
}

const (
	RegisterCicrFieldHserdycShift = 3
	RegisterCicrFieldHserdycMask  = 0x8
)

// GetHserdyc HSE ready Interrupt Clear
func (r *RegisterCicrType) GetHserdyc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCicrFieldHserdycMask) != 0
}

// SetHserdyc HSE ready Interrupt Clear
func (r *RegisterCicrType) SetHserdyc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCicrFieldHserdycMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCicrFieldHserdycMask)
	}
}

const (
	RegisterCicrFieldHsereadyinterruptclearShift = 4
	RegisterCicrFieldHsereadyinterruptclearMask  = 0x10
)

// GetHsereadyinterruptclear CSI ready Interrupt Clear
func (r *RegisterCicrType) GetHsereadyinterruptclear() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCicrFieldHsereadyinterruptclearMask) != 0
}

// SetHsereadyinterruptclear CSI ready Interrupt Clear
func (r *RegisterCicrType) SetHsereadyinterruptclear(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCicrFieldHsereadyinterruptclearMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCicrFieldHsereadyinterruptclearMask)
	}
}

const (
	RegisterCicrFieldRc48rdycShift = 5
	RegisterCicrFieldRc48rdycMask  = 0x20
)

// GetRc48rdyc RC48 ready Interrupt Clear
func (r *RegisterCicrType) GetRc48rdyc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCicrFieldRc48rdycMask) != 0
}

// SetRc48rdyc RC48 ready Interrupt Clear
func (r *RegisterCicrType) SetRc48rdyc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCicrFieldRc48rdycMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCicrFieldRc48rdycMask)
	}
}

const (
	RegisterCicrFieldPll1rdycShift = 6
	RegisterCicrFieldPll1rdycMask  = 0x40
)

// GetPll1rdyc PLL1 ready Interrupt Clear
func (r *RegisterCicrType) GetPll1rdyc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCicrFieldPll1rdycMask) != 0
}

// SetPll1rdyc PLL1 ready Interrupt Clear
func (r *RegisterCicrType) SetPll1rdyc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCicrFieldPll1rdycMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCicrFieldPll1rdycMask)
	}
}

const (
	RegisterCicrFieldPll2rdycShift = 7
	RegisterCicrFieldPll2rdycMask  = 0x80
)

// GetPll2rdyc PLL2 ready Interrupt Clear
func (r *RegisterCicrType) GetPll2rdyc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCicrFieldPll2rdycMask) != 0
}

// SetPll2rdyc PLL2 ready Interrupt Clear
func (r *RegisterCicrType) SetPll2rdyc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCicrFieldPll2rdycMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCicrFieldPll2rdycMask)
	}
}

const (
	RegisterCicrFieldPll3rdycShift = 8
	RegisterCicrFieldPll3rdycMask  = 0x100
)

// GetPll3rdyc PLL3 ready Interrupt Clear
func (r *RegisterCicrType) GetPll3rdyc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCicrFieldPll3rdycMask) != 0
}

// SetPll3rdyc PLL3 ready Interrupt Clear
func (r *RegisterCicrType) SetPll3rdyc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCicrFieldPll3rdycMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCicrFieldPll3rdycMask)
	}
}

const (
	RegisterCicrFieldLsecsscShift = 9
	RegisterCicrFieldLsecsscMask  = 0x200
)

// GetLsecssc LSE clock security system Interrupt Clear
func (r *RegisterCicrType) GetLsecssc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCicrFieldLsecsscMask) != 0
}

// SetLsecssc LSE clock security system Interrupt Clear
func (r *RegisterCicrType) SetLsecssc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCicrFieldLsecsscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCicrFieldLsecsscMask)
	}
}

const (
	RegisterCicrFieldHsecsscShift = 10
	RegisterCicrFieldHsecsscMask  = 0x400
)

// GetHsecssc HSE clock security system Interrupt Clear
func (r *RegisterCicrType) GetHsecssc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCicrFieldHsecsscMask) != 0
}

// SetHsecssc HSE clock security system Interrupt Clear
func (r *RegisterCicrType) SetHsecssc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCicrFieldHsecsscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCicrFieldHsecsscMask)
	}
}

// RegisterBdcrType RCC Backup Domain Control Register
type RegisterBdcrType uint32

func (r *RegisterBdcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterBdcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterBdcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterBdcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterBdcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterBdcrFieldLseonShift = 0
	RegisterBdcrFieldLseonMask  = 0x1
)

// GetLseon LSE oscillator enabled
func (r *RegisterBdcrType) GetLseon() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdcrFieldLseonMask) != 0
}

// SetLseon LSE oscillator enabled
func (r *RegisterBdcrType) SetLseon(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdcrFieldLseonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdcrFieldLseonMask)
	}
}

const (
	RegisterBdcrFieldLserdyShift = 1
	RegisterBdcrFieldLserdyMask  = 0x2
)

// GetLserdy LSE oscillator ready
func (r *RegisterBdcrType) GetLserdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdcrFieldLserdyMask) != 0
}

// SetLserdy LSE oscillator ready
func (r *RegisterBdcrType) SetLserdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdcrFieldLserdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdcrFieldLserdyMask)
	}
}

const (
	RegisterBdcrFieldLsebypShift = 2
	RegisterBdcrFieldLsebypMask  = 0x4
)

// GetLsebyp LSE oscillator bypass
func (r *RegisterBdcrType) GetLsebyp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdcrFieldLsebypMask) != 0
}

// SetLsebyp LSE oscillator bypass
func (r *RegisterBdcrType) SetLsebyp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdcrFieldLsebypMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdcrFieldLsebypMask)
	}
}

const (
	RegisterBdcrFieldLsedrvShift = 3
	RegisterBdcrFieldLsedrvMask  = 0x18
)

// GetLsedrv LSE oscillator driving capability
func (r *RegisterBdcrType) GetLsedrv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBdcrFieldLsedrvMask) >> RegisterBdcrFieldLsedrvShift)
}

// SetLsedrv LSE oscillator driving capability
func (r *RegisterBdcrType) SetLsedrv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBdcrFieldLsedrvMask)|(uint32(value)<<RegisterBdcrFieldLsedrvShift))
}

const (
	RegisterBdcrFieldLsecssonShift = 5
	RegisterBdcrFieldLsecssonMask  = 0x20
)

// GetLsecsson LSE clock security system enable
func (r *RegisterBdcrType) GetLsecsson() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdcrFieldLsecssonMask) != 0
}

// SetLsecsson LSE clock security system enable
func (r *RegisterBdcrType) SetLsecsson(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdcrFieldLsecssonMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdcrFieldLsecssonMask)
	}
}

const (
	RegisterBdcrFieldLsecssdShift = 6
	RegisterBdcrFieldLsecssdMask  = 0x40
)

// GetLsecssd LSE clock security system failure detection
func (r *RegisterBdcrType) GetLsecssd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdcrFieldLsecssdMask) != 0
}

// SetLsecssd LSE clock security system failure detection
func (r *RegisterBdcrType) SetLsecssd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdcrFieldLsecssdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdcrFieldLsecssdMask)
	}
}

const (
	RegisterBdcrFieldRtcsrcShift = 8
	RegisterBdcrFieldRtcsrcMask  = 0x300
)

// GetRtcsrc RTC clock source selection
func (r *RegisterBdcrType) GetRtcsrc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterBdcrFieldRtcsrcMask) >> RegisterBdcrFieldRtcsrcShift)
}

// SetRtcsrc RTC clock source selection
func (r *RegisterBdcrType) SetRtcsrc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterBdcrFieldRtcsrcMask)|(uint32(value)<<RegisterBdcrFieldRtcsrcShift))
}

const (
	RegisterBdcrFieldRtcenShift = 15
	RegisterBdcrFieldRtcenMask  = 0x8000
)

// GetRtcen RTC clock enable
func (r *RegisterBdcrType) GetRtcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdcrFieldRtcenMask) != 0
}

// SetRtcen RTC clock enable
func (r *RegisterBdcrType) SetRtcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdcrFieldRtcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdcrFieldRtcenMask)
	}
}

const (
	RegisterBdcrFieldVswrstShift = 16
	RegisterBdcrFieldVswrstMask  = 0x10000
)

// GetVswrst VSwitch domain software reset
func (r *RegisterBdcrType) GetVswrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterBdcrFieldVswrstMask) != 0
}

// SetVswrst VSwitch domain software reset
func (r *RegisterBdcrType) SetVswrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterBdcrFieldVswrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterBdcrFieldVswrstMask)
	}
}

// RegisterCsrType RCC Clock Control and Status Register
type RegisterCsrType uint32

func (r *RegisterCsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCsrFieldLsionShift = 0
	RegisterCsrFieldLsionMask  = 0x1
)

// GetLsion LSI oscillator enable
func (r *RegisterCsrType) GetLsion() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldLsionMask) != 0
}

// SetLsion LSI oscillator enable
func (r *RegisterCsrType) SetLsion(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldLsionMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldLsionMask)
	}
}

const (
	RegisterCsrFieldLsirdyShift = 1
	RegisterCsrFieldLsirdyMask  = 0x2
)

// GetLsirdy LSI oscillator ready
func (r *RegisterCsrType) GetLsirdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCsrFieldLsirdyMask) != 0
}

// SetLsirdy LSI oscillator ready
func (r *RegisterCsrType) SetLsirdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCsrFieldLsirdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCsrFieldLsirdyMask)
	}
}

// RegisterAhb3rstrType RCC AHB3 Reset Register
type RegisterAhb3rstrType uint32

func (r *RegisterAhb3rstrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb3rstrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb3rstrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb3rstrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb3rstrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb3rstrFieldMdmarstShift = 0
	RegisterAhb3rstrFieldMdmarstMask  = 0x1
)

// GetMdmarst MDMA block reset
func (r *RegisterAhb3rstrType) GetMdmarst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3rstrFieldMdmarstMask) != 0
}

// SetMdmarst MDMA block reset
func (r *RegisterAhb3rstrType) SetMdmarst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3rstrFieldMdmarstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3rstrFieldMdmarstMask)
	}
}

const (
	RegisterAhb3rstrFieldDma2drstShift = 4
	RegisterAhb3rstrFieldDma2drstMask  = 0x10
)

// GetDma2drst DMA2D block reset
func (r *RegisterAhb3rstrType) GetDma2drst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3rstrFieldDma2drstMask) != 0
}

// SetDma2drst DMA2D block reset
func (r *RegisterAhb3rstrType) SetDma2drst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3rstrFieldDma2drstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3rstrFieldDma2drstMask)
	}
}

const (
	RegisterAhb3rstrFieldJpgdecrstShift = 5
	RegisterAhb3rstrFieldJpgdecrstMask  = 0x20
)

// GetJpgdecrst JPGDEC block reset
func (r *RegisterAhb3rstrType) GetJpgdecrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3rstrFieldJpgdecrstMask) != 0
}

// SetJpgdecrst JPGDEC block reset
func (r *RegisterAhb3rstrType) SetJpgdecrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3rstrFieldJpgdecrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3rstrFieldJpgdecrstMask)
	}
}

const (
	RegisterAhb3rstrFieldFmcrstShift = 12
	RegisterAhb3rstrFieldFmcrstMask  = 0x1000
)

// GetFmcrst FMC block reset
func (r *RegisterAhb3rstrType) GetFmcrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3rstrFieldFmcrstMask) != 0
}

// SetFmcrst FMC block reset
func (r *RegisterAhb3rstrType) SetFmcrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3rstrFieldFmcrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3rstrFieldFmcrstMask)
	}
}

const (
	RegisterAhb3rstrFieldQspirstShift = 14
	RegisterAhb3rstrFieldQspirstMask  = 0x4000
)

// GetQspirst QUADSPI and QUADSPI delay block reset
func (r *RegisterAhb3rstrType) GetQspirst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3rstrFieldQspirstMask) != 0
}

// SetQspirst QUADSPI and QUADSPI delay block reset
func (r *RegisterAhb3rstrType) SetQspirst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3rstrFieldQspirstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3rstrFieldQspirstMask)
	}
}

const (
	RegisterAhb3rstrFieldSdmmc1rstShift = 16
	RegisterAhb3rstrFieldSdmmc1rstMask  = 0x10000
)

// GetSdmmc1rst SDMMC1 and SDMMC1 delay block reset
func (r *RegisterAhb3rstrType) GetSdmmc1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3rstrFieldSdmmc1rstMask) != 0
}

// SetSdmmc1rst SDMMC1 and SDMMC1 delay block reset
func (r *RegisterAhb3rstrType) SetSdmmc1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3rstrFieldSdmmc1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3rstrFieldSdmmc1rstMask)
	}
}

const (
	RegisterAhb3rstrFieldCpurstShift = 31
	RegisterAhb3rstrFieldCpurstMask  = 0x80000000
)

// GetCpurst CPU reset
func (r *RegisterAhb3rstrType) GetCpurst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3rstrFieldCpurstMask) != 0
}

// SetCpurst CPU reset
func (r *RegisterAhb3rstrType) SetCpurst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3rstrFieldCpurstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3rstrFieldCpurstMask)
	}
}

// RegisterAhb1rstrType RCC AHB1 Peripheral Reset Register
type RegisterAhb1rstrType uint32

func (r *RegisterAhb1rstrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb1rstrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb1rstrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb1rstrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb1rstrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb1rstrFieldDma1rstShift = 0
	RegisterAhb1rstrFieldDma1rstMask  = 0x1
)

// GetDma1rst DMA1 block reset
func (r *RegisterAhb1rstrType) GetDma1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1rstrFieldDma1rstMask) != 0
}

// SetDma1rst DMA1 block reset
func (r *RegisterAhb1rstrType) SetDma1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1rstrFieldDma1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1rstrFieldDma1rstMask)
	}
}

const (
	RegisterAhb1rstrFieldDma2rstShift = 1
	RegisterAhb1rstrFieldDma2rstMask  = 0x2
)

// GetDma2rst DMA2 block reset
func (r *RegisterAhb1rstrType) GetDma2rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1rstrFieldDma2rstMask) != 0
}

// SetDma2rst DMA2 block reset
func (r *RegisterAhb1rstrType) SetDma2rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1rstrFieldDma2rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1rstrFieldDma2rstMask)
	}
}

const (
	RegisterAhb1rstrFieldAdc12rstShift = 5
	RegisterAhb1rstrFieldAdc12rstMask  = 0x20
)

// GetAdc12rst ADC1&2 block reset
func (r *RegisterAhb1rstrType) GetAdc12rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1rstrFieldAdc12rstMask) != 0
}

// SetAdc12rst ADC1&2 block reset
func (r *RegisterAhb1rstrType) SetAdc12rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1rstrFieldAdc12rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1rstrFieldAdc12rstMask)
	}
}

const (
	RegisterAhb1rstrFieldEth1macrstShift = 15
	RegisterAhb1rstrFieldEth1macrstMask  = 0x8000
)

// GetEth1macrst ETH1MAC block reset
func (r *RegisterAhb1rstrType) GetEth1macrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1rstrFieldEth1macrstMask) != 0
}

// SetEth1macrst ETH1MAC block reset
func (r *RegisterAhb1rstrType) SetEth1macrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1rstrFieldEth1macrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1rstrFieldEth1macrstMask)
	}
}

const (
	RegisterAhb1rstrFieldUsb1otgrstShift = 25
	RegisterAhb1rstrFieldUsb1otgrstMask  = 0x2000000
)

// GetUsb1otgrst USB1OTG block reset
func (r *RegisterAhb1rstrType) GetUsb1otgrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1rstrFieldUsb1otgrstMask) != 0
}

// SetUsb1otgrst USB1OTG block reset
func (r *RegisterAhb1rstrType) SetUsb1otgrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1rstrFieldUsb1otgrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1rstrFieldUsb1otgrstMask)
	}
}

const (
	RegisterAhb1rstrFieldUsb2otgrstShift = 27
	RegisterAhb1rstrFieldUsb2otgrstMask  = 0x8000000
)

// GetUsb2otgrst USB2OTG block reset
func (r *RegisterAhb1rstrType) GetUsb2otgrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1rstrFieldUsb2otgrstMask) != 0
}

// SetUsb2otgrst USB2OTG block reset
func (r *RegisterAhb1rstrType) SetUsb2otgrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1rstrFieldUsb2otgrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1rstrFieldUsb2otgrstMask)
	}
}

// RegisterAhb2rstrType RCC AHB2 Peripheral Reset Register
type RegisterAhb2rstrType uint32

func (r *RegisterAhb2rstrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb2rstrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb2rstrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb2rstrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb2rstrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb2rstrFieldCamitfrstShift = 0
	RegisterAhb2rstrFieldCamitfrstMask  = 0x1
)

// GetCamitfrst CAMITF block reset
func (r *RegisterAhb2rstrType) GetCamitfrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2rstrFieldCamitfrstMask) != 0
}

// SetCamitfrst CAMITF block reset
func (r *RegisterAhb2rstrType) SetCamitfrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2rstrFieldCamitfrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2rstrFieldCamitfrstMask)
	}
}

const (
	RegisterAhb2rstrFieldCryptrstShift = 4
	RegisterAhb2rstrFieldCryptrstMask  = 0x10
)

// GetCryptrst Cryptography block reset
func (r *RegisterAhb2rstrType) GetCryptrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2rstrFieldCryptrstMask) != 0
}

// SetCryptrst Cryptography block reset
func (r *RegisterAhb2rstrType) SetCryptrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2rstrFieldCryptrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2rstrFieldCryptrstMask)
	}
}

const (
	RegisterAhb2rstrFieldHashrstShift = 5
	RegisterAhb2rstrFieldHashrstMask  = 0x20
)

// GetHashrst Hash block reset
func (r *RegisterAhb2rstrType) GetHashrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2rstrFieldHashrstMask) != 0
}

// SetHashrst Hash block reset
func (r *RegisterAhb2rstrType) SetHashrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2rstrFieldHashrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2rstrFieldHashrstMask)
	}
}

const (
	RegisterAhb2rstrFieldRngrstShift = 6
	RegisterAhb2rstrFieldRngrstMask  = 0x40
)

// GetRngrst Random Number Generator block reset
func (r *RegisterAhb2rstrType) GetRngrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2rstrFieldRngrstMask) != 0
}

// SetRngrst Random Number Generator block reset
func (r *RegisterAhb2rstrType) SetRngrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2rstrFieldRngrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2rstrFieldRngrstMask)
	}
}

const (
	RegisterAhb2rstrFieldSdmmc2rstShift = 9
	RegisterAhb2rstrFieldSdmmc2rstMask  = 0x200
)

// GetSdmmc2rst SDMMC2 and SDMMC2 Delay block reset
func (r *RegisterAhb2rstrType) GetSdmmc2rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2rstrFieldSdmmc2rstMask) != 0
}

// SetSdmmc2rst SDMMC2 and SDMMC2 Delay block reset
func (r *RegisterAhb2rstrType) SetSdmmc2rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2rstrFieldSdmmc2rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2rstrFieldSdmmc2rstMask)
	}
}

// RegisterAhb4rstrType RCC AHB4 Peripheral Reset Register
type RegisterAhb4rstrType uint32

func (r *RegisterAhb4rstrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb4rstrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb4rstrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb4rstrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb4rstrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb4rstrFieldGpioarstShift = 0
	RegisterAhb4rstrFieldGpioarstMask  = 0x1
)

// GetGpioarst GPIO block reset
func (r *RegisterAhb4rstrType) GetGpioarst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldGpioarstMask) != 0
}

// SetGpioarst GPIO block reset
func (r *RegisterAhb4rstrType) SetGpioarst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldGpioarstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldGpioarstMask)
	}
}

const (
	RegisterAhb4rstrFieldGpiobrstShift = 1
	RegisterAhb4rstrFieldGpiobrstMask  = 0x2
)

// GetGpiobrst GPIO block reset
func (r *RegisterAhb4rstrType) GetGpiobrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldGpiobrstMask) != 0
}

// SetGpiobrst GPIO block reset
func (r *RegisterAhb4rstrType) SetGpiobrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldGpiobrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldGpiobrstMask)
	}
}

const (
	RegisterAhb4rstrFieldGpiocrstShift = 2
	RegisterAhb4rstrFieldGpiocrstMask  = 0x4
)

// GetGpiocrst GPIO block reset
func (r *RegisterAhb4rstrType) GetGpiocrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldGpiocrstMask) != 0
}

// SetGpiocrst GPIO block reset
func (r *RegisterAhb4rstrType) SetGpiocrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldGpiocrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldGpiocrstMask)
	}
}

const (
	RegisterAhb4rstrFieldGpiodrstShift = 3
	RegisterAhb4rstrFieldGpiodrstMask  = 0x8
)

// GetGpiodrst GPIO block reset
func (r *RegisterAhb4rstrType) GetGpiodrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldGpiodrstMask) != 0
}

// SetGpiodrst GPIO block reset
func (r *RegisterAhb4rstrType) SetGpiodrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldGpiodrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldGpiodrstMask)
	}
}

const (
	RegisterAhb4rstrFieldGpioerstShift = 4
	RegisterAhb4rstrFieldGpioerstMask  = 0x10
)

// GetGpioerst GPIO block reset
func (r *RegisterAhb4rstrType) GetGpioerst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldGpioerstMask) != 0
}

// SetGpioerst GPIO block reset
func (r *RegisterAhb4rstrType) SetGpioerst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldGpioerstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldGpioerstMask)
	}
}

const (
	RegisterAhb4rstrFieldGpiofrstShift = 5
	RegisterAhb4rstrFieldGpiofrstMask  = 0x20
)

// GetGpiofrst GPIO block reset
func (r *RegisterAhb4rstrType) GetGpiofrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldGpiofrstMask) != 0
}

// SetGpiofrst GPIO block reset
func (r *RegisterAhb4rstrType) SetGpiofrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldGpiofrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldGpiofrstMask)
	}
}

const (
	RegisterAhb4rstrFieldGpiogrstShift = 6
	RegisterAhb4rstrFieldGpiogrstMask  = 0x40
)

// GetGpiogrst GPIO block reset
func (r *RegisterAhb4rstrType) GetGpiogrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldGpiogrstMask) != 0
}

// SetGpiogrst GPIO block reset
func (r *RegisterAhb4rstrType) SetGpiogrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldGpiogrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldGpiogrstMask)
	}
}

const (
	RegisterAhb4rstrFieldGpiohrstShift = 7
	RegisterAhb4rstrFieldGpiohrstMask  = 0x80
)

// GetGpiohrst GPIO block reset
func (r *RegisterAhb4rstrType) GetGpiohrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldGpiohrstMask) != 0
}

// SetGpiohrst GPIO block reset
func (r *RegisterAhb4rstrType) SetGpiohrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldGpiohrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldGpiohrstMask)
	}
}

const (
	RegisterAhb4rstrFieldGpioirstShift = 8
	RegisterAhb4rstrFieldGpioirstMask  = 0x100
)

// GetGpioirst GPIO block reset
func (r *RegisterAhb4rstrType) GetGpioirst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldGpioirstMask) != 0
}

// SetGpioirst GPIO block reset
func (r *RegisterAhb4rstrType) SetGpioirst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldGpioirstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldGpioirstMask)
	}
}

const (
	RegisterAhb4rstrFieldGpiojrstShift = 9
	RegisterAhb4rstrFieldGpiojrstMask  = 0x200
)

// GetGpiojrst GPIO block reset
func (r *RegisterAhb4rstrType) GetGpiojrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldGpiojrstMask) != 0
}

// SetGpiojrst GPIO block reset
func (r *RegisterAhb4rstrType) SetGpiojrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldGpiojrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldGpiojrstMask)
	}
}

const (
	RegisterAhb4rstrFieldGpiokrstShift = 10
	RegisterAhb4rstrFieldGpiokrstMask  = 0x400
)

// GetGpiokrst GPIO block reset
func (r *RegisterAhb4rstrType) GetGpiokrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldGpiokrstMask) != 0
}

// SetGpiokrst GPIO block reset
func (r *RegisterAhb4rstrType) SetGpiokrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldGpiokrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldGpiokrstMask)
	}
}

const (
	RegisterAhb4rstrFieldCrcrstShift = 19
	RegisterAhb4rstrFieldCrcrstMask  = 0x80000
)

// GetCrcrst CRC block reset
func (r *RegisterAhb4rstrType) GetCrcrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldCrcrstMask) != 0
}

// SetCrcrst CRC block reset
func (r *RegisterAhb4rstrType) SetCrcrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldCrcrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldCrcrstMask)
	}
}

const (
	RegisterAhb4rstrFieldBdmarstShift = 21
	RegisterAhb4rstrFieldBdmarstMask  = 0x200000
)

// GetBdmarst BDMA block reset
func (r *RegisterAhb4rstrType) GetBdmarst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldBdmarstMask) != 0
}

// SetBdmarst BDMA block reset
func (r *RegisterAhb4rstrType) SetBdmarst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldBdmarstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldBdmarstMask)
	}
}

const (
	RegisterAhb4rstrFieldAdc3rstShift = 24
	RegisterAhb4rstrFieldAdc3rstMask  = 0x1000000
)

// GetAdc3rst ADC3 block reset
func (r *RegisterAhb4rstrType) GetAdc3rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldAdc3rstMask) != 0
}

// SetAdc3rst ADC3 block reset
func (r *RegisterAhb4rstrType) SetAdc3rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldAdc3rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldAdc3rstMask)
	}
}

const (
	RegisterAhb4rstrFieldHsemrstShift = 25
	RegisterAhb4rstrFieldHsemrstMask  = 0x2000000
)

// GetHsemrst HSEM block reset
func (r *RegisterAhb4rstrType) GetHsemrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4rstrFieldHsemrstMask) != 0
}

// SetHsemrst HSEM block reset
func (r *RegisterAhb4rstrType) SetHsemrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4rstrFieldHsemrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4rstrFieldHsemrstMask)
	}
}

// RegisterApb3rstrType RCC APB3 Peripheral Reset Register
type RegisterApb3rstrType uint32

func (r *RegisterApb3rstrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb3rstrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb3rstrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb3rstrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb3rstrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb3rstrFieldLtdcrstShift = 3
	RegisterApb3rstrFieldLtdcrstMask  = 0x8
)

// GetLtdcrst LTDC block reset
func (r *RegisterApb3rstrType) GetLtdcrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb3rstrFieldLtdcrstMask) != 0
}

// SetLtdcrst LTDC block reset
func (r *RegisterApb3rstrType) SetLtdcrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb3rstrFieldLtdcrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb3rstrFieldLtdcrstMask)
	}
}

// RegisterApb1lrstrType RCC APB1 Peripheral Reset Register
type RegisterApb1lrstrType uint32

func (r *RegisterApb1lrstrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb1lrstrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb1lrstrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb1lrstrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb1lrstrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb1lrstrFieldTim2rstShift = 0
	RegisterApb1lrstrFieldTim2rstMask  = 0x1
)

// GetTim2rst TIM block reset
func (r *RegisterApb1lrstrType) GetTim2rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldTim2rstMask) != 0
}

// SetTim2rst TIM block reset
func (r *RegisterApb1lrstrType) SetTim2rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldTim2rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldTim2rstMask)
	}
}

const (
	RegisterApb1lrstrFieldTim3rstShift = 1
	RegisterApb1lrstrFieldTim3rstMask  = 0x2
)

// GetTim3rst TIM block reset
func (r *RegisterApb1lrstrType) GetTim3rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldTim3rstMask) != 0
}

// SetTim3rst TIM block reset
func (r *RegisterApb1lrstrType) SetTim3rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldTim3rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldTim3rstMask)
	}
}

const (
	RegisterApb1lrstrFieldTim4rstShift = 2
	RegisterApb1lrstrFieldTim4rstMask  = 0x4
)

// GetTim4rst TIM block reset
func (r *RegisterApb1lrstrType) GetTim4rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldTim4rstMask) != 0
}

// SetTim4rst TIM block reset
func (r *RegisterApb1lrstrType) SetTim4rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldTim4rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldTim4rstMask)
	}
}

const (
	RegisterApb1lrstrFieldTim5rstShift = 3
	RegisterApb1lrstrFieldTim5rstMask  = 0x8
)

// GetTim5rst TIM block reset
func (r *RegisterApb1lrstrType) GetTim5rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldTim5rstMask) != 0
}

// SetTim5rst TIM block reset
func (r *RegisterApb1lrstrType) SetTim5rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldTim5rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldTim5rstMask)
	}
}

const (
	RegisterApb1lrstrFieldTim6rstShift = 4
	RegisterApb1lrstrFieldTim6rstMask  = 0x10
)

// GetTim6rst TIM block reset
func (r *RegisterApb1lrstrType) GetTim6rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldTim6rstMask) != 0
}

// SetTim6rst TIM block reset
func (r *RegisterApb1lrstrType) SetTim6rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldTim6rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldTim6rstMask)
	}
}

const (
	RegisterApb1lrstrFieldTim7rstShift = 5
	RegisterApb1lrstrFieldTim7rstMask  = 0x20
)

// GetTim7rst TIM block reset
func (r *RegisterApb1lrstrType) GetTim7rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldTim7rstMask) != 0
}

// SetTim7rst TIM block reset
func (r *RegisterApb1lrstrType) SetTim7rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldTim7rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldTim7rstMask)
	}
}

const (
	RegisterApb1lrstrFieldTim12rstShift = 6
	RegisterApb1lrstrFieldTim12rstMask  = 0x40
)

// GetTim12rst TIM block reset
func (r *RegisterApb1lrstrType) GetTim12rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldTim12rstMask) != 0
}

// SetTim12rst TIM block reset
func (r *RegisterApb1lrstrType) SetTim12rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldTim12rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldTim12rstMask)
	}
}

const (
	RegisterApb1lrstrFieldTim13rstShift = 7
	RegisterApb1lrstrFieldTim13rstMask  = 0x80
)

// GetTim13rst TIM block reset
func (r *RegisterApb1lrstrType) GetTim13rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldTim13rstMask) != 0
}

// SetTim13rst TIM block reset
func (r *RegisterApb1lrstrType) SetTim13rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldTim13rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldTim13rstMask)
	}
}

const (
	RegisterApb1lrstrFieldTim14rstShift = 8
	RegisterApb1lrstrFieldTim14rstMask  = 0x100
)

// GetTim14rst TIM block reset
func (r *RegisterApb1lrstrType) GetTim14rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldTim14rstMask) != 0
}

// SetTim14rst TIM block reset
func (r *RegisterApb1lrstrType) SetTim14rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldTim14rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldTim14rstMask)
	}
}

const (
	RegisterApb1lrstrFieldLptim1rstShift = 9
	RegisterApb1lrstrFieldLptim1rstMask  = 0x200
)

// GetLptim1rst TIM block reset
func (r *RegisterApb1lrstrType) GetLptim1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldLptim1rstMask) != 0
}

// SetLptim1rst TIM block reset
func (r *RegisterApb1lrstrType) SetLptim1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldLptim1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldLptim1rstMask)
	}
}

const (
	RegisterApb1lrstrFieldSpi2rstShift = 14
	RegisterApb1lrstrFieldSpi2rstMask  = 0x4000
)

// GetSpi2rst SPI2 block reset
func (r *RegisterApb1lrstrType) GetSpi2rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldSpi2rstMask) != 0
}

// SetSpi2rst SPI2 block reset
func (r *RegisterApb1lrstrType) SetSpi2rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldSpi2rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldSpi2rstMask)
	}
}

const (
	RegisterApb1lrstrFieldSpi3rstShift = 15
	RegisterApb1lrstrFieldSpi3rstMask  = 0x8000
)

// GetSpi3rst SPI3 block reset
func (r *RegisterApb1lrstrType) GetSpi3rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldSpi3rstMask) != 0
}

// SetSpi3rst SPI3 block reset
func (r *RegisterApb1lrstrType) SetSpi3rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldSpi3rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldSpi3rstMask)
	}
}

const (
	RegisterApb1lrstrFieldSpdifrxrstShift = 16
	RegisterApb1lrstrFieldSpdifrxrstMask  = 0x10000
)

// GetSpdifrxrst SPDIFRX block reset
func (r *RegisterApb1lrstrType) GetSpdifrxrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldSpdifrxrstMask) != 0
}

// SetSpdifrxrst SPDIFRX block reset
func (r *RegisterApb1lrstrType) SetSpdifrxrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldSpdifrxrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldSpdifrxrstMask)
	}
}

const (
	RegisterApb1lrstrFieldUsart2rstShift = 17
	RegisterApb1lrstrFieldUsart2rstMask  = 0x20000
)

// GetUsart2rst USART2 block reset
func (r *RegisterApb1lrstrType) GetUsart2rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldUsart2rstMask) != 0
}

// SetUsart2rst USART2 block reset
func (r *RegisterApb1lrstrType) SetUsart2rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldUsart2rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldUsart2rstMask)
	}
}

const (
	RegisterApb1lrstrFieldUsart3rstShift = 18
	RegisterApb1lrstrFieldUsart3rstMask  = 0x40000
)

// GetUsart3rst USART3 block reset
func (r *RegisterApb1lrstrType) GetUsart3rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldUsart3rstMask) != 0
}

// SetUsart3rst USART3 block reset
func (r *RegisterApb1lrstrType) SetUsart3rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldUsart3rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldUsart3rstMask)
	}
}

const (
	RegisterApb1lrstrFieldUart4rstShift = 19
	RegisterApb1lrstrFieldUart4rstMask  = 0x80000
)

// GetUart4rst UART4 block reset
func (r *RegisterApb1lrstrType) GetUart4rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldUart4rstMask) != 0
}

// SetUart4rst UART4 block reset
func (r *RegisterApb1lrstrType) SetUart4rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldUart4rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldUart4rstMask)
	}
}

const (
	RegisterApb1lrstrFieldUart5rstShift = 20
	RegisterApb1lrstrFieldUart5rstMask  = 0x100000
)

// GetUart5rst UART5 block reset
func (r *RegisterApb1lrstrType) GetUart5rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldUart5rstMask) != 0
}

// SetUart5rst UART5 block reset
func (r *RegisterApb1lrstrType) SetUart5rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldUart5rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldUart5rstMask)
	}
}

const (
	RegisterApb1lrstrFieldI2c1rstShift = 21
	RegisterApb1lrstrFieldI2c1rstMask  = 0x200000
)

// GetI2c1rst I2C1 block reset
func (r *RegisterApb1lrstrType) GetI2c1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldI2c1rstMask) != 0
}

// SetI2c1rst I2C1 block reset
func (r *RegisterApb1lrstrType) SetI2c1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldI2c1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldI2c1rstMask)
	}
}

const (
	RegisterApb1lrstrFieldI2c2rstShift = 22
	RegisterApb1lrstrFieldI2c2rstMask  = 0x400000
)

// GetI2c2rst I2C2 block reset
func (r *RegisterApb1lrstrType) GetI2c2rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldI2c2rstMask) != 0
}

// SetI2c2rst I2C2 block reset
func (r *RegisterApb1lrstrType) SetI2c2rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldI2c2rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldI2c2rstMask)
	}
}

const (
	RegisterApb1lrstrFieldI2c3rstShift = 23
	RegisterApb1lrstrFieldI2c3rstMask  = 0x800000
)

// GetI2c3rst I2C3 block reset
func (r *RegisterApb1lrstrType) GetI2c3rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldI2c3rstMask) != 0
}

// SetI2c3rst I2C3 block reset
func (r *RegisterApb1lrstrType) SetI2c3rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldI2c3rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldI2c3rstMask)
	}
}

const (
	RegisterApb1lrstrFieldHdmicecrstShift = 27
	RegisterApb1lrstrFieldHdmicecrstMask  = 0x8000000
)

// GetHdmicecrst HDMI-CEC block reset
func (r *RegisterApb1lrstrType) GetHdmicecrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldHdmicecrstMask) != 0
}

// SetHdmicecrst HDMI-CEC block reset
func (r *RegisterApb1lrstrType) SetHdmicecrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldHdmicecrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldHdmicecrstMask)
	}
}

const (
	RegisterApb1lrstrFieldDac12rstShift = 29
	RegisterApb1lrstrFieldDac12rstMask  = 0x20000000
)

// GetDac12rst DAC1 and 2 Blocks Reset
func (r *RegisterApb1lrstrType) GetDac12rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldDac12rstMask) != 0
}

// SetDac12rst DAC1 and 2 Blocks Reset
func (r *RegisterApb1lrstrType) SetDac12rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldDac12rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldDac12rstMask)
	}
}

const (
	RegisterApb1lrstrFieldUsart7rstShift = 30
	RegisterApb1lrstrFieldUsart7rstMask  = 0x40000000
)

// GetUsart7rst USART7 block reset
func (r *RegisterApb1lrstrType) GetUsart7rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldUsart7rstMask) != 0
}

// SetUsart7rst USART7 block reset
func (r *RegisterApb1lrstrType) SetUsart7rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldUsart7rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldUsart7rstMask)
	}
}

const (
	RegisterApb1lrstrFieldUsart8rstShift = 31
	RegisterApb1lrstrFieldUsart8rstMask  = 0x80000000
)

// GetUsart8rst USART8 block reset
func (r *RegisterApb1lrstrType) GetUsart8rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lrstrFieldUsart8rstMask) != 0
}

// SetUsart8rst USART8 block reset
func (r *RegisterApb1lrstrType) SetUsart8rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lrstrFieldUsart8rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lrstrFieldUsart8rstMask)
	}
}

// RegisterApb1hrstrType RCC APB1 Peripheral Reset Register
type RegisterApb1hrstrType uint32

func (r *RegisterApb1hrstrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb1hrstrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb1hrstrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb1hrstrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb1hrstrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb1hrstrFieldCrsrstShift = 1
	RegisterApb1hrstrFieldCrsrstMask  = 0x2
)

// GetCrsrst Clock Recovery System reset
func (r *RegisterApb1hrstrType) GetCrsrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1hrstrFieldCrsrstMask) != 0
}

// SetCrsrst Clock Recovery System reset
func (r *RegisterApb1hrstrType) SetCrsrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1hrstrFieldCrsrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1hrstrFieldCrsrstMask)
	}
}

const (
	RegisterApb1hrstrFieldSwprstShift = 2
	RegisterApb1hrstrFieldSwprstMask  = 0x4
)

// GetSwprst SWPMI block reset
func (r *RegisterApb1hrstrType) GetSwprst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1hrstrFieldSwprstMask) != 0
}

// SetSwprst SWPMI block reset
func (r *RegisterApb1hrstrType) SetSwprst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1hrstrFieldSwprstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1hrstrFieldSwprstMask)
	}
}

const (
	RegisterApb1hrstrFieldOpamprstShift = 4
	RegisterApb1hrstrFieldOpamprstMask  = 0x10
)

// GetOpamprst OPAMP block reset
func (r *RegisterApb1hrstrType) GetOpamprst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1hrstrFieldOpamprstMask) != 0
}

// SetOpamprst OPAMP block reset
func (r *RegisterApb1hrstrType) SetOpamprst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1hrstrFieldOpamprstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1hrstrFieldOpamprstMask)
	}
}

const (
	RegisterApb1hrstrFieldMdiosrstShift = 5
	RegisterApb1hrstrFieldMdiosrstMask  = 0x20
)

// GetMdiosrst MDIOS block reset
func (r *RegisterApb1hrstrType) GetMdiosrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1hrstrFieldMdiosrstMask) != 0
}

// SetMdiosrst MDIOS block reset
func (r *RegisterApb1hrstrType) SetMdiosrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1hrstrFieldMdiosrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1hrstrFieldMdiosrstMask)
	}
}

const (
	RegisterApb1hrstrFieldFdcanrstShift = 8
	RegisterApb1hrstrFieldFdcanrstMask  = 0x100
)

// GetFdcanrst FDCAN block reset
func (r *RegisterApb1hrstrType) GetFdcanrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1hrstrFieldFdcanrstMask) != 0
}

// SetFdcanrst FDCAN block reset
func (r *RegisterApb1hrstrType) SetFdcanrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1hrstrFieldFdcanrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1hrstrFieldFdcanrstMask)
	}
}

// RegisterApb2rstrType RCC APB2 Peripheral Reset Register
type RegisterApb2rstrType uint32

func (r *RegisterApb2rstrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb2rstrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb2rstrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb2rstrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb2rstrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb2rstrFieldTim1rstShift = 0
	RegisterApb2rstrFieldTim1rstMask  = 0x1
)

// GetTim1rst TIM1 block reset
func (r *RegisterApb2rstrType) GetTim1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldTim1rstMask) != 0
}

// SetTim1rst TIM1 block reset
func (r *RegisterApb2rstrType) SetTim1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldTim1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldTim1rstMask)
	}
}

const (
	RegisterApb2rstrFieldTim8rstShift = 1
	RegisterApb2rstrFieldTim8rstMask  = 0x2
)

// GetTim8rst TIM8 block reset
func (r *RegisterApb2rstrType) GetTim8rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldTim8rstMask) != 0
}

// SetTim8rst TIM8 block reset
func (r *RegisterApb2rstrType) SetTim8rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldTim8rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldTim8rstMask)
	}
}

const (
	RegisterApb2rstrFieldUsart1rstShift = 4
	RegisterApb2rstrFieldUsart1rstMask  = 0x10
)

// GetUsart1rst USART1 block reset
func (r *RegisterApb2rstrType) GetUsart1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldUsart1rstMask) != 0
}

// SetUsart1rst USART1 block reset
func (r *RegisterApb2rstrType) SetUsart1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldUsart1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldUsart1rstMask)
	}
}

const (
	RegisterApb2rstrFieldUsart6rstShift = 5
	RegisterApb2rstrFieldUsart6rstMask  = 0x20
)

// GetUsart6rst USART6 block reset
func (r *RegisterApb2rstrType) GetUsart6rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldUsart6rstMask) != 0
}

// SetUsart6rst USART6 block reset
func (r *RegisterApb2rstrType) SetUsart6rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldUsart6rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldUsart6rstMask)
	}
}

const (
	RegisterApb2rstrFieldSpi1rstShift = 12
	RegisterApb2rstrFieldSpi1rstMask  = 0x1000
)

// GetSpi1rst SPI1 block reset
func (r *RegisterApb2rstrType) GetSpi1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldSpi1rstMask) != 0
}

// SetSpi1rst SPI1 block reset
func (r *RegisterApb2rstrType) SetSpi1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldSpi1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldSpi1rstMask)
	}
}

const (
	RegisterApb2rstrFieldSpi4rstShift = 13
	RegisterApb2rstrFieldSpi4rstMask  = 0x2000
)

// GetSpi4rst SPI4 block reset
func (r *RegisterApb2rstrType) GetSpi4rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldSpi4rstMask) != 0
}

// SetSpi4rst SPI4 block reset
func (r *RegisterApb2rstrType) SetSpi4rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldSpi4rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldSpi4rstMask)
	}
}

const (
	RegisterApb2rstrFieldTim15rstShift = 16
	RegisterApb2rstrFieldTim15rstMask  = 0x10000
)

// GetTim15rst TIM15 block reset
func (r *RegisterApb2rstrType) GetTim15rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldTim15rstMask) != 0
}

// SetTim15rst TIM15 block reset
func (r *RegisterApb2rstrType) SetTim15rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldTim15rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldTim15rstMask)
	}
}

const (
	RegisterApb2rstrFieldTim16rstShift = 17
	RegisterApb2rstrFieldTim16rstMask  = 0x20000
)

// GetTim16rst TIM16 block reset
func (r *RegisterApb2rstrType) GetTim16rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldTim16rstMask) != 0
}

// SetTim16rst TIM16 block reset
func (r *RegisterApb2rstrType) SetTim16rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldTim16rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldTim16rstMask)
	}
}

const (
	RegisterApb2rstrFieldTim17rstShift = 18
	RegisterApb2rstrFieldTim17rstMask  = 0x40000
)

// GetTim17rst TIM17 block reset
func (r *RegisterApb2rstrType) GetTim17rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldTim17rstMask) != 0
}

// SetTim17rst TIM17 block reset
func (r *RegisterApb2rstrType) SetTim17rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldTim17rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldTim17rstMask)
	}
}

const (
	RegisterApb2rstrFieldSpi5rstShift = 20
	RegisterApb2rstrFieldSpi5rstMask  = 0x100000
)

// GetSpi5rst SPI5 block reset
func (r *RegisterApb2rstrType) GetSpi5rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldSpi5rstMask) != 0
}

// SetSpi5rst SPI5 block reset
func (r *RegisterApb2rstrType) SetSpi5rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldSpi5rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldSpi5rstMask)
	}
}

const (
	RegisterApb2rstrFieldSai1rstShift = 22
	RegisterApb2rstrFieldSai1rstMask  = 0x400000
)

// GetSai1rst SAI1 block reset
func (r *RegisterApb2rstrType) GetSai1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldSai1rstMask) != 0
}

// SetSai1rst SAI1 block reset
func (r *RegisterApb2rstrType) SetSai1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldSai1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldSai1rstMask)
	}
}

const (
	RegisterApb2rstrFieldSai2rstShift = 23
	RegisterApb2rstrFieldSai2rstMask  = 0x800000
)

// GetSai2rst SAI2 block reset
func (r *RegisterApb2rstrType) GetSai2rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldSai2rstMask) != 0
}

// SetSai2rst SAI2 block reset
func (r *RegisterApb2rstrType) SetSai2rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldSai2rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldSai2rstMask)
	}
}

const (
	RegisterApb2rstrFieldSai3rstShift = 24
	RegisterApb2rstrFieldSai3rstMask  = 0x1000000
)

// GetSai3rst SAI3 block reset
func (r *RegisterApb2rstrType) GetSai3rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldSai3rstMask) != 0
}

// SetSai3rst SAI3 block reset
func (r *RegisterApb2rstrType) SetSai3rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldSai3rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldSai3rstMask)
	}
}

const (
	RegisterApb2rstrFieldDfsdm1rstShift = 28
	RegisterApb2rstrFieldDfsdm1rstMask  = 0x10000000
)

// GetDfsdm1rst DFSDM1 block reset
func (r *RegisterApb2rstrType) GetDfsdm1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldDfsdm1rstMask) != 0
}

// SetDfsdm1rst DFSDM1 block reset
func (r *RegisterApb2rstrType) SetDfsdm1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldDfsdm1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldDfsdm1rstMask)
	}
}

const (
	RegisterApb2rstrFieldHrtimrstShift = 29
	RegisterApb2rstrFieldHrtimrstMask  = 0x20000000
)

// GetHrtimrst HRTIM block reset
func (r *RegisterApb2rstrType) GetHrtimrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2rstrFieldHrtimrstMask) != 0
}

// SetHrtimrst HRTIM block reset
func (r *RegisterApb2rstrType) SetHrtimrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2rstrFieldHrtimrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2rstrFieldHrtimrstMask)
	}
}

// RegisterApb4rstrType RCC APB4 Peripheral Reset Register
type RegisterApb4rstrType uint32

func (r *RegisterApb4rstrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb4rstrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb4rstrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb4rstrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb4rstrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb4rstrFieldSyscfgrstShift = 1
	RegisterApb4rstrFieldSyscfgrstMask  = 0x2
)

// GetSyscfgrst SYSCFG block reset
func (r *RegisterApb4rstrType) GetSyscfgrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4rstrFieldSyscfgrstMask) != 0
}

// SetSyscfgrst SYSCFG block reset
func (r *RegisterApb4rstrType) SetSyscfgrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4rstrFieldSyscfgrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4rstrFieldSyscfgrstMask)
	}
}

const (
	RegisterApb4rstrFieldLpuart1rstShift = 3
	RegisterApb4rstrFieldLpuart1rstMask  = 0x8
)

// GetLpuart1rst LPUART1 block reset
func (r *RegisterApb4rstrType) GetLpuart1rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4rstrFieldLpuart1rstMask) != 0
}

// SetLpuart1rst LPUART1 block reset
func (r *RegisterApb4rstrType) SetLpuart1rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4rstrFieldLpuart1rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4rstrFieldLpuart1rstMask)
	}
}

const (
	RegisterApb4rstrFieldSpi6rstShift = 5
	RegisterApb4rstrFieldSpi6rstMask  = 0x20
)

// GetSpi6rst SPI6 block reset
func (r *RegisterApb4rstrType) GetSpi6rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4rstrFieldSpi6rstMask) != 0
}

// SetSpi6rst SPI6 block reset
func (r *RegisterApb4rstrType) SetSpi6rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4rstrFieldSpi6rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4rstrFieldSpi6rstMask)
	}
}

const (
	RegisterApb4rstrFieldI2c4rstShift = 7
	RegisterApb4rstrFieldI2c4rstMask  = 0x80
)

// GetI2c4rst I2C4 block reset
func (r *RegisterApb4rstrType) GetI2c4rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4rstrFieldI2c4rstMask) != 0
}

// SetI2c4rst I2C4 block reset
func (r *RegisterApb4rstrType) SetI2c4rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4rstrFieldI2c4rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4rstrFieldI2c4rstMask)
	}
}

const (
	RegisterApb4rstrFieldLptim2rstShift = 9
	RegisterApb4rstrFieldLptim2rstMask  = 0x200
)

// GetLptim2rst LPTIM2 block reset
func (r *RegisterApb4rstrType) GetLptim2rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4rstrFieldLptim2rstMask) != 0
}

// SetLptim2rst LPTIM2 block reset
func (r *RegisterApb4rstrType) SetLptim2rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4rstrFieldLptim2rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4rstrFieldLptim2rstMask)
	}
}

const (
	RegisterApb4rstrFieldLptim3rstShift = 10
	RegisterApb4rstrFieldLptim3rstMask  = 0x400
)

// GetLptim3rst LPTIM3 block reset
func (r *RegisterApb4rstrType) GetLptim3rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4rstrFieldLptim3rstMask) != 0
}

// SetLptim3rst LPTIM3 block reset
func (r *RegisterApb4rstrType) SetLptim3rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4rstrFieldLptim3rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4rstrFieldLptim3rstMask)
	}
}

const (
	RegisterApb4rstrFieldLptim4rstShift = 11
	RegisterApb4rstrFieldLptim4rstMask  = 0x800
)

// GetLptim4rst LPTIM4 block reset
func (r *RegisterApb4rstrType) GetLptim4rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4rstrFieldLptim4rstMask) != 0
}

// SetLptim4rst LPTIM4 block reset
func (r *RegisterApb4rstrType) SetLptim4rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4rstrFieldLptim4rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4rstrFieldLptim4rstMask)
	}
}

const (
	RegisterApb4rstrFieldLptim5rstShift = 12
	RegisterApb4rstrFieldLptim5rstMask  = 0x1000
)

// GetLptim5rst LPTIM5 block reset
func (r *RegisterApb4rstrType) GetLptim5rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4rstrFieldLptim5rstMask) != 0
}

// SetLptim5rst LPTIM5 block reset
func (r *RegisterApb4rstrType) SetLptim5rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4rstrFieldLptim5rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4rstrFieldLptim5rstMask)
	}
}

const (
	RegisterApb4rstrFieldComp12rstShift = 14
	RegisterApb4rstrFieldComp12rstMask  = 0x4000
)

// GetComp12rst COMP12 Blocks Reset
func (r *RegisterApb4rstrType) GetComp12rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4rstrFieldComp12rstMask) != 0
}

// SetComp12rst COMP12 Blocks Reset
func (r *RegisterApb4rstrType) SetComp12rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4rstrFieldComp12rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4rstrFieldComp12rstMask)
	}
}

const (
	RegisterApb4rstrFieldVrefrstShift = 15
	RegisterApb4rstrFieldVrefrstMask  = 0x8000
)

// GetVrefrst VREF block reset
func (r *RegisterApb4rstrType) GetVrefrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4rstrFieldVrefrstMask) != 0
}

// SetVrefrst VREF block reset
func (r *RegisterApb4rstrType) SetVrefrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4rstrFieldVrefrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4rstrFieldVrefrstMask)
	}
}

const (
	RegisterApb4rstrFieldSai4rstShift = 21
	RegisterApb4rstrFieldSai4rstMask  = 0x200000
)

// GetSai4rst SAI4 block reset
func (r *RegisterApb4rstrType) GetSai4rst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4rstrFieldSai4rstMask) != 0
}

// SetSai4rst SAI4 block reset
func (r *RegisterApb4rstrType) SetSai4rst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4rstrFieldSai4rstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4rstrFieldSai4rstMask)
	}
}

// RegisterGcrType RCC Global Control Register
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
	RegisterGcrFieldWw1rscShift = 0
	RegisterGcrFieldWw1rscMask  = 0x1
)

// GetWw1rsc WWDG1 reset scope control
func (r *RegisterGcrType) GetWw1rsc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterGcrFieldWw1rscMask) != 0
}

// SetWw1rsc WWDG1 reset scope control
func (r *RegisterGcrType) SetWw1rsc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterGcrFieldWw1rscMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterGcrFieldWw1rscMask)
	}
}

// RegisterD3amrType RCC D3 Autonomous mode Register
type RegisterD3amrType uint32

func (r *RegisterD3amrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterD3amrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterD3amrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterD3amrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterD3amrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterD3amrFieldBdmaamenShift = 0
	RegisterD3amrFieldBdmaamenMask  = 0x1
)

// GetBdmaamen BDMA and DMAMUX Autonomous mode enable
func (r *RegisterD3amrType) GetBdmaamen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldBdmaamenMask) != 0
}

// SetBdmaamen BDMA and DMAMUX Autonomous mode enable
func (r *RegisterD3amrType) SetBdmaamen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldBdmaamenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldBdmaamenMask)
	}
}

const (
	RegisterD3amrFieldLpuart1amenShift = 3
	RegisterD3amrFieldLpuart1amenMask  = 0x8
)

// GetLpuart1amen LPUART1 Autonomous mode enable
func (r *RegisterD3amrType) GetLpuart1amen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldLpuart1amenMask) != 0
}

// SetLpuart1amen LPUART1 Autonomous mode enable
func (r *RegisterD3amrType) SetLpuart1amen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldLpuart1amenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldLpuart1amenMask)
	}
}

const (
	RegisterD3amrFieldSpi6amenShift = 5
	RegisterD3amrFieldSpi6amenMask  = 0x20
)

// GetSpi6amen SPI6 Autonomous mode enable
func (r *RegisterD3amrType) GetSpi6amen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldSpi6amenMask) != 0
}

// SetSpi6amen SPI6 Autonomous mode enable
func (r *RegisterD3amrType) SetSpi6amen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldSpi6amenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldSpi6amenMask)
	}
}

const (
	RegisterD3amrFieldI2c4amenShift = 7
	RegisterD3amrFieldI2c4amenMask  = 0x80
)

// GetI2c4amen I2C4 Autonomous mode enable
func (r *RegisterD3amrType) GetI2c4amen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldI2c4amenMask) != 0
}

// SetI2c4amen I2C4 Autonomous mode enable
func (r *RegisterD3amrType) SetI2c4amen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldI2c4amenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldI2c4amenMask)
	}
}

const (
	RegisterD3amrFieldLptim2amenShift = 9
	RegisterD3amrFieldLptim2amenMask  = 0x200
)

// GetLptim2amen LPTIM2 Autonomous mode enable
func (r *RegisterD3amrType) GetLptim2amen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldLptim2amenMask) != 0
}

// SetLptim2amen LPTIM2 Autonomous mode enable
func (r *RegisterD3amrType) SetLptim2amen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldLptim2amenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldLptim2amenMask)
	}
}

const (
	RegisterD3amrFieldLptim3amenShift = 10
	RegisterD3amrFieldLptim3amenMask  = 0x400
)

// GetLptim3amen LPTIM3 Autonomous mode enable
func (r *RegisterD3amrType) GetLptim3amen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldLptim3amenMask) != 0
}

// SetLptim3amen LPTIM3 Autonomous mode enable
func (r *RegisterD3amrType) SetLptim3amen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldLptim3amenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldLptim3amenMask)
	}
}

const (
	RegisterD3amrFieldLptim4amenShift = 11
	RegisterD3amrFieldLptim4amenMask  = 0x800
)

// GetLptim4amen LPTIM4 Autonomous mode enable
func (r *RegisterD3amrType) GetLptim4amen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldLptim4amenMask) != 0
}

// SetLptim4amen LPTIM4 Autonomous mode enable
func (r *RegisterD3amrType) SetLptim4amen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldLptim4amenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldLptim4amenMask)
	}
}

const (
	RegisterD3amrFieldLptim5amenShift = 12
	RegisterD3amrFieldLptim5amenMask  = 0x1000
)

// GetLptim5amen LPTIM5 Autonomous mode enable
func (r *RegisterD3amrType) GetLptim5amen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldLptim5amenMask) != 0
}

// SetLptim5amen LPTIM5 Autonomous mode enable
func (r *RegisterD3amrType) SetLptim5amen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldLptim5amenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldLptim5amenMask)
	}
}

const (
	RegisterD3amrFieldComp12amenShift = 14
	RegisterD3amrFieldComp12amenMask  = 0x4000
)

// GetComp12amen COMP12 Autonomous mode enable
func (r *RegisterD3amrType) GetComp12amen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldComp12amenMask) != 0
}

// SetComp12amen COMP12 Autonomous mode enable
func (r *RegisterD3amrType) SetComp12amen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldComp12amenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldComp12amenMask)
	}
}

const (
	RegisterD3amrFieldVrefamenShift = 15
	RegisterD3amrFieldVrefamenMask  = 0x8000
)

// GetVrefamen VREF Autonomous mode enable
func (r *RegisterD3amrType) GetVrefamen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldVrefamenMask) != 0
}

// SetVrefamen VREF Autonomous mode enable
func (r *RegisterD3amrType) SetVrefamen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldVrefamenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldVrefamenMask)
	}
}

const (
	RegisterD3amrFieldRtcamenShift = 16
	RegisterD3amrFieldRtcamenMask  = 0x10000
)

// GetRtcamen RTC Autonomous mode enable
func (r *RegisterD3amrType) GetRtcamen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldRtcamenMask) != 0
}

// SetRtcamen RTC Autonomous mode enable
func (r *RegisterD3amrType) SetRtcamen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldRtcamenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldRtcamenMask)
	}
}

const (
	RegisterD3amrFieldCrcamenShift = 19
	RegisterD3amrFieldCrcamenMask  = 0x80000
)

// GetCrcamen CRC Autonomous mode enable
func (r *RegisterD3amrType) GetCrcamen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldCrcamenMask) != 0
}

// SetCrcamen CRC Autonomous mode enable
func (r *RegisterD3amrType) SetCrcamen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldCrcamenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldCrcamenMask)
	}
}

const (
	RegisterD3amrFieldSai4amenShift = 21
	RegisterD3amrFieldSai4amenMask  = 0x200000
)

// GetSai4amen SAI4 Autonomous mode enable
func (r *RegisterD3amrType) GetSai4amen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldSai4amenMask) != 0
}

// SetSai4amen SAI4 Autonomous mode enable
func (r *RegisterD3amrType) SetSai4amen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldSai4amenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldSai4amenMask)
	}
}

const (
	RegisterD3amrFieldAdc3amenShift = 24
	RegisterD3amrFieldAdc3amenMask  = 0x1000000
)

// GetAdc3amen ADC3 Autonomous mode enable
func (r *RegisterD3amrType) GetAdc3amen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldAdc3amenMask) != 0
}

// SetAdc3amen ADC3 Autonomous mode enable
func (r *RegisterD3amrType) SetAdc3amen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldAdc3amenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldAdc3amenMask)
	}
}

const (
	RegisterD3amrFieldBkpsramamenShift = 28
	RegisterD3amrFieldBkpsramamenMask  = 0x10000000
)

// GetBkpsramamen Backup RAM Autonomous mode enable
func (r *RegisterD3amrType) GetBkpsramamen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldBkpsramamenMask) != 0
}

// SetBkpsramamen Backup RAM Autonomous mode enable
func (r *RegisterD3amrType) SetBkpsramamen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldBkpsramamenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldBkpsramamenMask)
	}
}

const (
	RegisterD3amrFieldSram4amenShift = 29
	RegisterD3amrFieldSram4amenMask  = 0x20000000
)

// GetSram4amen SRAM4 Autonomous mode enable
func (r *RegisterD3amrType) GetSram4amen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterD3amrFieldSram4amenMask) != 0
}

// SetSram4amen SRAM4 Autonomous mode enable
func (r *RegisterD3amrType) SetSram4amen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterD3amrFieldSram4amenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterD3amrFieldSram4amenMask)
	}
}

// RegisterRsrType RCC Reset Status Register
type RegisterRsrType uint32

func (r *RegisterRsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRsrFieldRmvfShift = 16
	RegisterRsrFieldRmvfMask  = 0x10000
)

// GetRmvf Remove reset flag
func (r *RegisterRsrType) GetRmvf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsrFieldRmvfMask) != 0
}

// SetRmvf Remove reset flag
func (r *RegisterRsrType) SetRmvf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsrFieldRmvfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsrFieldRmvfMask)
	}
}

const (
	RegisterRsrFieldCpurstfShift = 17
	RegisterRsrFieldCpurstfMask  = 0x20000
)

// GetCpurstf CPU reset flag
func (r *RegisterRsrType) GetCpurstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsrFieldCpurstfMask) != 0
}

// SetCpurstf CPU reset flag
func (r *RegisterRsrType) SetCpurstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsrFieldCpurstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsrFieldCpurstfMask)
	}
}

const (
	RegisterRsrFieldD1rstfShift = 19
	RegisterRsrFieldD1rstfMask  = 0x80000
)

// GetD1rstf D1 domain power switch reset flag
func (r *RegisterRsrType) GetD1rstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsrFieldD1rstfMask) != 0
}

// SetD1rstf D1 domain power switch reset flag
func (r *RegisterRsrType) SetD1rstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsrFieldD1rstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsrFieldD1rstfMask)
	}
}

const (
	RegisterRsrFieldD2rstfShift = 20
	RegisterRsrFieldD2rstfMask  = 0x100000
)

// GetD2rstf D2 domain power switch reset flag
func (r *RegisterRsrType) GetD2rstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsrFieldD2rstfMask) != 0
}

// SetD2rstf D2 domain power switch reset flag
func (r *RegisterRsrType) SetD2rstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsrFieldD2rstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsrFieldD2rstfMask)
	}
}

const (
	RegisterRsrFieldBorrstfShift = 21
	RegisterRsrFieldBorrstfMask  = 0x200000
)

// GetBorrstf BOR reset flag
func (r *RegisterRsrType) GetBorrstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsrFieldBorrstfMask) != 0
}

// SetBorrstf BOR reset flag
func (r *RegisterRsrType) SetBorrstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsrFieldBorrstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsrFieldBorrstfMask)
	}
}

const (
	RegisterRsrFieldPinrstfShift = 22
	RegisterRsrFieldPinrstfMask  = 0x400000
)

// GetPinrstf Pin reset flag (NRST)
func (r *RegisterRsrType) GetPinrstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsrFieldPinrstfMask) != 0
}

// SetPinrstf Pin reset flag (NRST)
func (r *RegisterRsrType) SetPinrstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsrFieldPinrstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsrFieldPinrstfMask)
	}
}

const (
	RegisterRsrFieldPorrstfShift = 23
	RegisterRsrFieldPorrstfMask  = 0x800000
)

// GetPorrstf POR/PDR reset flag
func (r *RegisterRsrType) GetPorrstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsrFieldPorrstfMask) != 0
}

// SetPorrstf POR/PDR reset flag
func (r *RegisterRsrType) SetPorrstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsrFieldPorrstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsrFieldPorrstfMask)
	}
}

const (
	RegisterRsrFieldSftrstfShift = 24
	RegisterRsrFieldSftrstfMask  = 0x1000000
)

// GetSftrstf System reset from CPU reset flag
func (r *RegisterRsrType) GetSftrstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsrFieldSftrstfMask) != 0
}

// SetSftrstf System reset from CPU reset flag
func (r *RegisterRsrType) SetSftrstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsrFieldSftrstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsrFieldSftrstfMask)
	}
}

const (
	RegisterRsrFieldIwdg1rstfShift = 26
	RegisterRsrFieldIwdg1rstfMask  = 0x4000000
)

// GetIwdg1rstf Independent Watchdog reset flag
func (r *RegisterRsrType) GetIwdg1rstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsrFieldIwdg1rstfMask) != 0
}

// SetIwdg1rstf Independent Watchdog reset flag
func (r *RegisterRsrType) SetIwdg1rstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsrFieldIwdg1rstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsrFieldIwdg1rstfMask)
	}
}

const (
	RegisterRsrFieldWwdg1rstfShift = 28
	RegisterRsrFieldWwdg1rstfMask  = 0x10000000
)

// GetWwdg1rstf Window Watchdog reset flag
func (r *RegisterRsrType) GetWwdg1rstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsrFieldWwdg1rstfMask) != 0
}

// SetWwdg1rstf Window Watchdog reset flag
func (r *RegisterRsrType) SetWwdg1rstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsrFieldWwdg1rstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsrFieldWwdg1rstfMask)
	}
}

const (
	RegisterRsrFieldLpwrrstfShift = 30
	RegisterRsrFieldLpwrrstfMask  = 0x40000000
)

// GetLpwrrstf Reset due to illegal D1 DStandby or CPU CStop flag
func (r *RegisterRsrType) GetLpwrrstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterRsrFieldLpwrrstfMask) != 0
}

// SetLpwrrstf Reset due to illegal D1 DStandby or CPU CStop flag
func (r *RegisterRsrType) SetLpwrrstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterRsrFieldLpwrrstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterRsrFieldLpwrrstfMask)
	}
}

// RegisterAhb3enrType RCC AHB3 Clock Register
type RegisterAhb3enrType uint32

func (r *RegisterAhb3enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb3enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb3enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb3enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb3enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb3enrFieldMdmaenShift = 0
	RegisterAhb3enrFieldMdmaenMask  = 0x1
)

// GetMdmaen MDMA Peripheral Clock Enable
func (r *RegisterAhb3enrType) GetMdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3enrFieldMdmaenMask) != 0
}

// SetMdmaen MDMA Peripheral Clock Enable
func (r *RegisterAhb3enrType) SetMdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3enrFieldMdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3enrFieldMdmaenMask)
	}
}

const (
	RegisterAhb3enrFieldDma2denShift = 4
	RegisterAhb3enrFieldDma2denMask  = 0x10
)

// GetDma2den DMA2D Peripheral Clock Enable
func (r *RegisterAhb3enrType) GetDma2den() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3enrFieldDma2denMask) != 0
}

// SetDma2den DMA2D Peripheral Clock Enable
func (r *RegisterAhb3enrType) SetDma2den(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3enrFieldDma2denMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3enrFieldDma2denMask)
	}
}

const (
	RegisterAhb3enrFieldJpgdecenShift = 5
	RegisterAhb3enrFieldJpgdecenMask  = 0x20
)

// GetJpgdecen JPGDEC Peripheral Clock Enable
func (r *RegisterAhb3enrType) GetJpgdecen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3enrFieldJpgdecenMask) != 0
}

// SetJpgdecen JPGDEC Peripheral Clock Enable
func (r *RegisterAhb3enrType) SetJpgdecen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3enrFieldJpgdecenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3enrFieldJpgdecenMask)
	}
}

const (
	RegisterAhb3enrFieldFmcenShift = 12
	RegisterAhb3enrFieldFmcenMask  = 0x1000
)

// GetFmcen FMC Peripheral Clocks Enable
func (r *RegisterAhb3enrType) GetFmcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3enrFieldFmcenMask) != 0
}

// SetFmcen FMC Peripheral Clocks Enable
func (r *RegisterAhb3enrType) SetFmcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3enrFieldFmcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3enrFieldFmcenMask)
	}
}

const (
	RegisterAhb3enrFieldQspienShift = 14
	RegisterAhb3enrFieldQspienMask  = 0x4000
)

// GetQspien QUADSPI and QUADSPI Delay Clock Enable
func (r *RegisterAhb3enrType) GetQspien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3enrFieldQspienMask) != 0
}

// SetQspien QUADSPI and QUADSPI Delay Clock Enable
func (r *RegisterAhb3enrType) SetQspien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3enrFieldQspienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3enrFieldQspienMask)
	}
}

const (
	RegisterAhb3enrFieldSdmmc1enShift = 16
	RegisterAhb3enrFieldSdmmc1enMask  = 0x10000
)

// GetSdmmc1en SDMMC1 and SDMMC1 Delay Clock Enable
func (r *RegisterAhb3enrType) GetSdmmc1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3enrFieldSdmmc1enMask) != 0
}

// SetSdmmc1en SDMMC1 and SDMMC1 Delay Clock Enable
func (r *RegisterAhb3enrType) SetSdmmc1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3enrFieldSdmmc1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3enrFieldSdmmc1enMask)
	}
}

// RegisterAhb1enrType RCC AHB1 Clock Register
type RegisterAhb1enrType uint32

func (r *RegisterAhb1enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb1enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb1enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb1enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb1enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb1enrFieldDma1enShift = 0
	RegisterAhb1enrFieldDma1enMask  = 0x1
)

// GetDma1en DMA1 Clock Enable
func (r *RegisterAhb1enrType) GetDma1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1enrFieldDma1enMask) != 0
}

// SetDma1en DMA1 Clock Enable
func (r *RegisterAhb1enrType) SetDma1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1enrFieldDma1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1enrFieldDma1enMask)
	}
}

const (
	RegisterAhb1enrFieldDma2enShift = 1
	RegisterAhb1enrFieldDma2enMask  = 0x2
)

// GetDma2en DMA2 Clock Enable
func (r *RegisterAhb1enrType) GetDma2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1enrFieldDma2enMask) != 0
}

// SetDma2en DMA2 Clock Enable
func (r *RegisterAhb1enrType) SetDma2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1enrFieldDma2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1enrFieldDma2enMask)
	}
}

const (
	RegisterAhb1enrFieldAdc12enShift = 5
	RegisterAhb1enrFieldAdc12enMask  = 0x20
)

// GetAdc12en ADC1/2 Peripheral Clocks Enable
func (r *RegisterAhb1enrType) GetAdc12en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1enrFieldAdc12enMask) != 0
}

// SetAdc12en ADC1/2 Peripheral Clocks Enable
func (r *RegisterAhb1enrType) SetAdc12en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1enrFieldAdc12enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1enrFieldAdc12enMask)
	}
}

const (
	RegisterAhb1enrFieldEth1macenShift = 15
	RegisterAhb1enrFieldEth1macenMask  = 0x8000
)

// GetEth1macen Ethernet MAC bus interface Clock Enable
func (r *RegisterAhb1enrType) GetEth1macen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1enrFieldEth1macenMask) != 0
}

// SetEth1macen Ethernet MAC bus interface Clock Enable
func (r *RegisterAhb1enrType) SetEth1macen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1enrFieldEth1macenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1enrFieldEth1macenMask)
	}
}

const (
	RegisterAhb1enrFieldEth1txenShift = 16
	RegisterAhb1enrFieldEth1txenMask  = 0x10000
)

// GetEth1txen Ethernet Transmission Clock Enable
func (r *RegisterAhb1enrType) GetEth1txen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1enrFieldEth1txenMask) != 0
}

// SetEth1txen Ethernet Transmission Clock Enable
func (r *RegisterAhb1enrType) SetEth1txen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1enrFieldEth1txenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1enrFieldEth1txenMask)
	}
}

const (
	RegisterAhb1enrFieldEth1rxenShift = 17
	RegisterAhb1enrFieldEth1rxenMask  = 0x20000
)

// GetEth1rxen Ethernet Reception Clock Enable
func (r *RegisterAhb1enrType) GetEth1rxen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1enrFieldEth1rxenMask) != 0
}

// SetEth1rxen Ethernet Reception Clock Enable
func (r *RegisterAhb1enrType) SetEth1rxen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1enrFieldEth1rxenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1enrFieldEth1rxenMask)
	}
}

const (
	RegisterAhb1enrFieldUsb1otgenShift = 25
	RegisterAhb1enrFieldUsb1otgenMask  = 0x2000000
)

// GetUsb1otgen USB1OTG Peripheral Clocks Enable
func (r *RegisterAhb1enrType) GetUsb1otgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1enrFieldUsb1otgenMask) != 0
}

// SetUsb1otgen USB1OTG Peripheral Clocks Enable
func (r *RegisterAhb1enrType) SetUsb1otgen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1enrFieldUsb1otgenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1enrFieldUsb1otgenMask)
	}
}

const (
	RegisterAhb1enrFieldUsb1ulpienShift = 26
	RegisterAhb1enrFieldUsb1ulpienMask  = 0x4000000
)

// GetUsb1ulpien USB_PHY1 Clocks Enable
func (r *RegisterAhb1enrType) GetUsb1ulpien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1enrFieldUsb1ulpienMask) != 0
}

// SetUsb1ulpien USB_PHY1 Clocks Enable
func (r *RegisterAhb1enrType) SetUsb1ulpien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1enrFieldUsb1ulpienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1enrFieldUsb1ulpienMask)
	}
}

const (
	RegisterAhb1enrFieldUsb2otgenShift = 27
	RegisterAhb1enrFieldUsb2otgenMask  = 0x8000000
)

// GetUsb2otgen USB2OTG Peripheral Clocks Enable
func (r *RegisterAhb1enrType) GetUsb2otgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1enrFieldUsb2otgenMask) != 0
}

// SetUsb2otgen USB2OTG Peripheral Clocks Enable
func (r *RegisterAhb1enrType) SetUsb2otgen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1enrFieldUsb2otgenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1enrFieldUsb2otgenMask)
	}
}

const (
	RegisterAhb1enrFieldUsb2ulpienShift = 28
	RegisterAhb1enrFieldUsb2ulpienMask  = 0x10000000
)

// GetUsb2ulpien USB_PHY2 Clocks Enable
func (r *RegisterAhb1enrType) GetUsb2ulpien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1enrFieldUsb2ulpienMask) != 0
}

// SetUsb2ulpien USB_PHY2 Clocks Enable
func (r *RegisterAhb1enrType) SetUsb2ulpien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1enrFieldUsb2ulpienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1enrFieldUsb2ulpienMask)
	}
}

// RegisterAhb2enrType RCC AHB2 Clock Register
type RegisterAhb2enrType uint32

func (r *RegisterAhb2enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb2enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb2enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb2enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb2enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb2enrFieldCamitfenShift = 0
	RegisterAhb2enrFieldCamitfenMask  = 0x1
)

// GetCamitfen CAMITF peripheral clock enable
func (r *RegisterAhb2enrType) GetCamitfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2enrFieldCamitfenMask) != 0
}

// SetCamitfen CAMITF peripheral clock enable
func (r *RegisterAhb2enrType) SetCamitfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2enrFieldCamitfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2enrFieldCamitfenMask)
	}
}

const (
	RegisterAhb2enrFieldCryptenShift = 4
	RegisterAhb2enrFieldCryptenMask  = 0x10
)

// GetCrypten CRYPT peripheral clock enable
func (r *RegisterAhb2enrType) GetCrypten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2enrFieldCryptenMask) != 0
}

// SetCrypten CRYPT peripheral clock enable
func (r *RegisterAhb2enrType) SetCrypten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2enrFieldCryptenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2enrFieldCryptenMask)
	}
}

const (
	RegisterAhb2enrFieldHashenShift = 5
	RegisterAhb2enrFieldHashenMask  = 0x20
)

// GetHashen HASH peripheral clock enable
func (r *RegisterAhb2enrType) GetHashen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2enrFieldHashenMask) != 0
}

// SetHashen HASH peripheral clock enable
func (r *RegisterAhb2enrType) SetHashen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2enrFieldHashenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2enrFieldHashenMask)
	}
}

const (
	RegisterAhb2enrFieldRngenShift = 6
	RegisterAhb2enrFieldRngenMask  = 0x40
)

// GetRngen RNG peripheral clocks enable
func (r *RegisterAhb2enrType) GetRngen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2enrFieldRngenMask) != 0
}

// SetRngen RNG peripheral clocks enable
func (r *RegisterAhb2enrType) SetRngen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2enrFieldRngenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2enrFieldRngenMask)
	}
}

const (
	RegisterAhb2enrFieldSdmmc2enShift = 9
	RegisterAhb2enrFieldSdmmc2enMask  = 0x200
)

// GetSdmmc2en SDMMC2 and SDMMC2 delay clock enable
func (r *RegisterAhb2enrType) GetSdmmc2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2enrFieldSdmmc2enMask) != 0
}

// SetSdmmc2en SDMMC2 and SDMMC2 delay clock enable
func (r *RegisterAhb2enrType) SetSdmmc2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2enrFieldSdmmc2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2enrFieldSdmmc2enMask)
	}
}

const (
	RegisterAhb2enrFieldSram1enShift = 29
	RegisterAhb2enrFieldSram1enMask  = 0x20000000
)

// GetSram1en SRAM1 block enable
func (r *RegisterAhb2enrType) GetSram1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2enrFieldSram1enMask) != 0
}

// SetSram1en SRAM1 block enable
func (r *RegisterAhb2enrType) SetSram1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2enrFieldSram1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2enrFieldSram1enMask)
	}
}

const (
	RegisterAhb2enrFieldSram2enShift = 30
	RegisterAhb2enrFieldSram2enMask  = 0x40000000
)

// GetSram2en SRAM2 block enable
func (r *RegisterAhb2enrType) GetSram2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2enrFieldSram2enMask) != 0
}

// SetSram2en SRAM2 block enable
func (r *RegisterAhb2enrType) SetSram2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2enrFieldSram2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2enrFieldSram2enMask)
	}
}

const (
	RegisterAhb2enrFieldSram3enShift = 31
	RegisterAhb2enrFieldSram3enMask  = 0x80000000
)

// GetSram3en SRAM3 block enable
func (r *RegisterAhb2enrType) GetSram3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2enrFieldSram3enMask) != 0
}

// SetSram3en SRAM3 block enable
func (r *RegisterAhb2enrType) SetSram3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2enrFieldSram3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2enrFieldSram3enMask)
	}
}

// RegisterAhb4enrType RCC AHB4 Clock Register
type RegisterAhb4enrType uint32

func (r *RegisterAhb4enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb4enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb4enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb4enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb4enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb4enrFieldGpioaenShift = 0
	RegisterAhb4enrFieldGpioaenMask  = 0x1
)

// GetGpioaen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) GetGpioaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldGpioaenMask) != 0
}

// SetGpioaen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) SetGpioaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldGpioaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldGpioaenMask)
	}
}

const (
	RegisterAhb4enrFieldGpiobenShift = 1
	RegisterAhb4enrFieldGpiobenMask  = 0x2
)

// GetGpioben 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) GetGpioben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldGpiobenMask) != 0
}

// SetGpioben 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) SetGpioben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldGpiobenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldGpiobenMask)
	}
}

const (
	RegisterAhb4enrFieldGpiocenShift = 2
	RegisterAhb4enrFieldGpiocenMask  = 0x4
)

// GetGpiocen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) GetGpiocen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldGpiocenMask) != 0
}

// SetGpiocen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) SetGpiocen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldGpiocenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldGpiocenMask)
	}
}

const (
	RegisterAhb4enrFieldGpiodenShift = 3
	RegisterAhb4enrFieldGpiodenMask  = 0x8
)

// GetGpioden 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) GetGpioden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldGpiodenMask) != 0
}

// SetGpioden 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) SetGpioden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldGpiodenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldGpiodenMask)
	}
}

const (
	RegisterAhb4enrFieldGpioeenShift = 4
	RegisterAhb4enrFieldGpioeenMask  = 0x10
)

// GetGpioeen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) GetGpioeen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldGpioeenMask) != 0
}

// SetGpioeen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) SetGpioeen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldGpioeenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldGpioeenMask)
	}
}

const (
	RegisterAhb4enrFieldGpiofenShift = 5
	RegisterAhb4enrFieldGpiofenMask  = 0x20
)

// GetGpiofen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) GetGpiofen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldGpiofenMask) != 0
}

// SetGpiofen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) SetGpiofen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldGpiofenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldGpiofenMask)
	}
}

const (
	RegisterAhb4enrFieldGpiogenShift = 6
	RegisterAhb4enrFieldGpiogenMask  = 0x40
)

// GetGpiogen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) GetGpiogen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldGpiogenMask) != 0
}

// SetGpiogen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) SetGpiogen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldGpiogenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldGpiogenMask)
	}
}

const (
	RegisterAhb4enrFieldGpiohenShift = 7
	RegisterAhb4enrFieldGpiohenMask  = 0x80
)

// GetGpiohen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) GetGpiohen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldGpiohenMask) != 0
}

// SetGpiohen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) SetGpiohen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldGpiohenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldGpiohenMask)
	}
}

const (
	RegisterAhb4enrFieldGpioienShift = 8
	RegisterAhb4enrFieldGpioienMask  = 0x100
)

// GetGpioien 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) GetGpioien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldGpioienMask) != 0
}

// SetGpioien 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) SetGpioien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldGpioienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldGpioienMask)
	}
}

const (
	RegisterAhb4enrFieldGpiojenShift = 9
	RegisterAhb4enrFieldGpiojenMask  = 0x200
)

// GetGpiojen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) GetGpiojen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldGpiojenMask) != 0
}

// SetGpiojen 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) SetGpiojen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldGpiojenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldGpiojenMask)
	}
}

const (
	RegisterAhb4enrFieldGpiokenShift = 10
	RegisterAhb4enrFieldGpiokenMask  = 0x400
)

// GetGpioken 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) GetGpioken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldGpiokenMask) != 0
}

// SetGpioken 0GPIO peripheral clock enable
func (r *RegisterAhb4enrType) SetGpioken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldGpiokenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldGpiokenMask)
	}
}

const (
	RegisterAhb4enrFieldCrcenShift = 19
	RegisterAhb4enrFieldCrcenMask  = 0x80000
)

// GetCrcen CRC peripheral clock enable
func (r *RegisterAhb4enrType) GetCrcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldCrcenMask) != 0
}

// SetCrcen CRC peripheral clock enable
func (r *RegisterAhb4enrType) SetCrcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldCrcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldCrcenMask)
	}
}

const (
	RegisterAhb4enrFieldBdmaenShift = 21
	RegisterAhb4enrFieldBdmaenMask  = 0x200000
)

// GetBdmaen BDMA and DMAMUX2 Clock Enable
func (r *RegisterAhb4enrType) GetBdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldBdmaenMask) != 0
}

// SetBdmaen BDMA and DMAMUX2 Clock Enable
func (r *RegisterAhb4enrType) SetBdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldBdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldBdmaenMask)
	}
}

const (
	RegisterAhb4enrFieldAdc3enShift = 24
	RegisterAhb4enrFieldAdc3enMask  = 0x1000000
)

// GetAdc3en ADC3 Peripheral Clocks Enable
func (r *RegisterAhb4enrType) GetAdc3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldAdc3enMask) != 0
}

// SetAdc3en ADC3 Peripheral Clocks Enable
func (r *RegisterAhb4enrType) SetAdc3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldAdc3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldAdc3enMask)
	}
}

const (
	RegisterAhb4enrFieldHsemenShift = 25
	RegisterAhb4enrFieldHsemenMask  = 0x2000000
)

// GetHsemen HSEM peripheral clock enable
func (r *RegisterAhb4enrType) GetHsemen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldHsemenMask) != 0
}

// SetHsemen HSEM peripheral clock enable
func (r *RegisterAhb4enrType) SetHsemen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldHsemenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldHsemenMask)
	}
}

const (
	RegisterAhb4enrFieldBkpramenShift = 28
	RegisterAhb4enrFieldBkpramenMask  = 0x10000000
)

// GetBkpramen Backup RAM Clock Enable
func (r *RegisterAhb4enrType) GetBkpramen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4enrFieldBkpramenMask) != 0
}

// SetBkpramen Backup RAM Clock Enable
func (r *RegisterAhb4enrType) SetBkpramen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4enrFieldBkpramenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4enrFieldBkpramenMask)
	}
}

// RegisterApb3enrType RCC APB3 Clock Register
type RegisterApb3enrType uint32

func (r *RegisterApb3enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb3enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb3enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb3enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb3enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb3enrFieldLtdcenShift = 3
	RegisterApb3enrFieldLtdcenMask  = 0x8
)

// GetLtdcen LTDC peripheral clock enable
func (r *RegisterApb3enrType) GetLtdcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb3enrFieldLtdcenMask) != 0
}

// SetLtdcen LTDC peripheral clock enable
func (r *RegisterApb3enrType) SetLtdcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb3enrFieldLtdcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb3enrFieldLtdcenMask)
	}
}

const (
	RegisterApb3enrFieldWwdg1enShift = 6
	RegisterApb3enrFieldWwdg1enMask  = 0x40
)

// GetWwdg1en WWDG1 Clock Enable
func (r *RegisterApb3enrType) GetWwdg1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb3enrFieldWwdg1enMask) != 0
}

// SetWwdg1en WWDG1 Clock Enable
func (r *RegisterApb3enrType) SetWwdg1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb3enrFieldWwdg1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb3enrFieldWwdg1enMask)
	}
}

// RegisterApb1lenrType RCC APB1 Clock Register
type RegisterApb1lenrType uint32

func (r *RegisterApb1lenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb1lenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb1lenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb1lenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb1lenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb1lenrFieldTim2enShift = 0
	RegisterApb1lenrFieldTim2enMask  = 0x1
)

// GetTim2en TIM peripheral clock enable
func (r *RegisterApb1lenrType) GetTim2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldTim2enMask) != 0
}

// SetTim2en TIM peripheral clock enable
func (r *RegisterApb1lenrType) SetTim2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldTim2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldTim2enMask)
	}
}

const (
	RegisterApb1lenrFieldTim3enShift = 1
	RegisterApb1lenrFieldTim3enMask  = 0x2
)

// GetTim3en TIM peripheral clock enable
func (r *RegisterApb1lenrType) GetTim3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldTim3enMask) != 0
}

// SetTim3en TIM peripheral clock enable
func (r *RegisterApb1lenrType) SetTim3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldTim3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldTim3enMask)
	}
}

const (
	RegisterApb1lenrFieldTim4enShift = 2
	RegisterApb1lenrFieldTim4enMask  = 0x4
)

// GetTim4en TIM peripheral clock enable
func (r *RegisterApb1lenrType) GetTim4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldTim4enMask) != 0
}

// SetTim4en TIM peripheral clock enable
func (r *RegisterApb1lenrType) SetTim4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldTim4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldTim4enMask)
	}
}

const (
	RegisterApb1lenrFieldTim5enShift = 3
	RegisterApb1lenrFieldTim5enMask  = 0x8
)

// GetTim5en TIM peripheral clock enable
func (r *RegisterApb1lenrType) GetTim5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldTim5enMask) != 0
}

// SetTim5en TIM peripheral clock enable
func (r *RegisterApb1lenrType) SetTim5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldTim5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldTim5enMask)
	}
}

const (
	RegisterApb1lenrFieldTim6enShift = 4
	RegisterApb1lenrFieldTim6enMask  = 0x10
)

// GetTim6en TIM peripheral clock enable
func (r *RegisterApb1lenrType) GetTim6en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldTim6enMask) != 0
}

// SetTim6en TIM peripheral clock enable
func (r *RegisterApb1lenrType) SetTim6en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldTim6enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldTim6enMask)
	}
}

const (
	RegisterApb1lenrFieldTim7enShift = 5
	RegisterApb1lenrFieldTim7enMask  = 0x20
)

// GetTim7en TIM peripheral clock enable
func (r *RegisterApb1lenrType) GetTim7en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldTim7enMask) != 0
}

// SetTim7en TIM peripheral clock enable
func (r *RegisterApb1lenrType) SetTim7en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldTim7enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldTim7enMask)
	}
}

const (
	RegisterApb1lenrFieldTim12enShift = 6
	RegisterApb1lenrFieldTim12enMask  = 0x40
)

// GetTim12en TIM peripheral clock enable
func (r *RegisterApb1lenrType) GetTim12en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldTim12enMask) != 0
}

// SetTim12en TIM peripheral clock enable
func (r *RegisterApb1lenrType) SetTim12en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldTim12enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldTim12enMask)
	}
}

const (
	RegisterApb1lenrFieldTim13enShift = 7
	RegisterApb1lenrFieldTim13enMask  = 0x80
)

// GetTim13en TIM peripheral clock enable
func (r *RegisterApb1lenrType) GetTim13en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldTim13enMask) != 0
}

// SetTim13en TIM peripheral clock enable
func (r *RegisterApb1lenrType) SetTim13en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldTim13enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldTim13enMask)
	}
}

const (
	RegisterApb1lenrFieldTim14enShift = 8
	RegisterApb1lenrFieldTim14enMask  = 0x100
)

// GetTim14en TIM peripheral clock enable
func (r *RegisterApb1lenrType) GetTim14en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldTim14enMask) != 0
}

// SetTim14en TIM peripheral clock enable
func (r *RegisterApb1lenrType) SetTim14en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldTim14enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldTim14enMask)
	}
}

const (
	RegisterApb1lenrFieldLptim1enShift = 9
	RegisterApb1lenrFieldLptim1enMask  = 0x200
)

// GetLptim1en LPTIM1 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetLptim1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldLptim1enMask) != 0
}

// SetLptim1en LPTIM1 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetLptim1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldLptim1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldLptim1enMask)
	}
}

const (
	RegisterApb1lenrFieldSpi2enShift = 14
	RegisterApb1lenrFieldSpi2enMask  = 0x4000
)

// GetSpi2en SPI2 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetSpi2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldSpi2enMask) != 0
}

// SetSpi2en SPI2 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetSpi2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldSpi2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldSpi2enMask)
	}
}

const (
	RegisterApb1lenrFieldSpi3enShift = 15
	RegisterApb1lenrFieldSpi3enMask  = 0x8000
)

// GetSpi3en SPI3 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetSpi3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldSpi3enMask) != 0
}

// SetSpi3en SPI3 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetSpi3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldSpi3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldSpi3enMask)
	}
}

const (
	RegisterApb1lenrFieldSpdifrxenShift = 16
	RegisterApb1lenrFieldSpdifrxenMask  = 0x10000
)

// GetSpdifrxen SPDIFRX Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetSpdifrxen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldSpdifrxenMask) != 0
}

// SetSpdifrxen SPDIFRX Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetSpdifrxen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldSpdifrxenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldSpdifrxenMask)
	}
}

const (
	RegisterApb1lenrFieldUsart2enShift = 17
	RegisterApb1lenrFieldUsart2enMask  = 0x20000
)

// GetUsart2en USART2 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetUsart2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldUsart2enMask) != 0
}

// SetUsart2en USART2 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetUsart2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldUsart2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldUsart2enMask)
	}
}

const (
	RegisterApb1lenrFieldUsart3enShift = 18
	RegisterApb1lenrFieldUsart3enMask  = 0x40000
)

// GetUsart3en USART3 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetUsart3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldUsart3enMask) != 0
}

// SetUsart3en USART3 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetUsart3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldUsart3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldUsart3enMask)
	}
}

const (
	RegisterApb1lenrFieldUart4enShift = 19
	RegisterApb1lenrFieldUart4enMask  = 0x80000
)

// GetUart4en UART4 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetUart4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldUart4enMask) != 0
}

// SetUart4en UART4 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetUart4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldUart4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldUart4enMask)
	}
}

const (
	RegisterApb1lenrFieldUart5enShift = 20
	RegisterApb1lenrFieldUart5enMask  = 0x100000
)

// GetUart5en UART5 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetUart5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldUart5enMask) != 0
}

// SetUart5en UART5 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetUart5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldUart5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldUart5enMask)
	}
}

const (
	RegisterApb1lenrFieldI2c1enShift = 21
	RegisterApb1lenrFieldI2c1enMask  = 0x200000
)

// GetI2c1en I2C1 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetI2c1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldI2c1enMask) != 0
}

// SetI2c1en I2C1 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetI2c1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldI2c1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldI2c1enMask)
	}
}

const (
	RegisterApb1lenrFieldI2c2enShift = 22
	RegisterApb1lenrFieldI2c2enMask  = 0x400000
)

// GetI2c2en I2C2 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetI2c2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldI2c2enMask) != 0
}

// SetI2c2en I2C2 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetI2c2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldI2c2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldI2c2enMask)
	}
}

const (
	RegisterApb1lenrFieldI2c3enShift = 23
	RegisterApb1lenrFieldI2c3enMask  = 0x800000
)

// GetI2c3en I2C3 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetI2c3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldI2c3enMask) != 0
}

// SetI2c3en I2C3 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetI2c3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldI2c3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldI2c3enMask)
	}
}

const (
	RegisterApb1lenrFieldHdmicecenShift = 27
	RegisterApb1lenrFieldHdmicecenMask  = 0x8000000
)

// GetHdmicecen HDMI-CEC peripheral clock enable
func (r *RegisterApb1lenrType) GetHdmicecen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldHdmicecenMask) != 0
}

// SetHdmicecen HDMI-CEC peripheral clock enable
func (r *RegisterApb1lenrType) SetHdmicecen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldHdmicecenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldHdmicecenMask)
	}
}

const (
	RegisterApb1lenrFieldDac12enShift = 29
	RegisterApb1lenrFieldDac12enMask  = 0x20000000
)

// GetDac12en DAC1&2 peripheral clock enable
func (r *RegisterApb1lenrType) GetDac12en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldDac12enMask) != 0
}

// SetDac12en DAC1&2 peripheral clock enable
func (r *RegisterApb1lenrType) SetDac12en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldDac12enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldDac12enMask)
	}
}

const (
	RegisterApb1lenrFieldUsart7enShift = 30
	RegisterApb1lenrFieldUsart7enMask  = 0x40000000
)

// GetUsart7en USART7 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetUsart7en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldUsart7enMask) != 0
}

// SetUsart7en USART7 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetUsart7en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldUsart7enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldUsart7enMask)
	}
}

const (
	RegisterApb1lenrFieldUsart8enShift = 31
	RegisterApb1lenrFieldUsart8enMask  = 0x80000000
)

// GetUsart8en USART8 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) GetUsart8en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1lenrFieldUsart8enMask) != 0
}

// SetUsart8en USART8 Peripheral Clocks Enable
func (r *RegisterApb1lenrType) SetUsart8en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1lenrFieldUsart8enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1lenrFieldUsart8enMask)
	}
}

// RegisterApb1henrType RCC APB1 Clock Register
type RegisterApb1henrType uint32

func (r *RegisterApb1henrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb1henrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb1henrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb1henrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb1henrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb1henrFieldCrsenShift = 1
	RegisterApb1henrFieldCrsenMask  = 0x2
)

// GetCrsen Clock Recovery System peripheral clock enable
func (r *RegisterApb1henrType) GetCrsen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1henrFieldCrsenMask) != 0
}

// SetCrsen Clock Recovery System peripheral clock enable
func (r *RegisterApb1henrType) SetCrsen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1henrFieldCrsenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1henrFieldCrsenMask)
	}
}

const (
	RegisterApb1henrFieldSwpenShift = 2
	RegisterApb1henrFieldSwpenMask  = 0x4
)

// GetSwpen SWPMI Peripheral Clocks Enable
func (r *RegisterApb1henrType) GetSwpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1henrFieldSwpenMask) != 0
}

// SetSwpen SWPMI Peripheral Clocks Enable
func (r *RegisterApb1henrType) SetSwpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1henrFieldSwpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1henrFieldSwpenMask)
	}
}

const (
	RegisterApb1henrFieldOpampenShift = 4
	RegisterApb1henrFieldOpampenMask  = 0x10
)

// GetOpampen OPAMP peripheral clock enable
func (r *RegisterApb1henrType) GetOpampen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1henrFieldOpampenMask) != 0
}

// SetOpampen OPAMP peripheral clock enable
func (r *RegisterApb1henrType) SetOpampen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1henrFieldOpampenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1henrFieldOpampenMask)
	}
}

const (
	RegisterApb1henrFieldMdiosenShift = 5
	RegisterApb1henrFieldMdiosenMask  = 0x20
)

// GetMdiosen MDIOS peripheral clock enable
func (r *RegisterApb1henrType) GetMdiosen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1henrFieldMdiosenMask) != 0
}

// SetMdiosen MDIOS peripheral clock enable
func (r *RegisterApb1henrType) SetMdiosen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1henrFieldMdiosenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1henrFieldMdiosenMask)
	}
}

const (
	RegisterApb1henrFieldFdcanenShift = 8
	RegisterApb1henrFieldFdcanenMask  = 0x100
)

// GetFdcanen FDCAN Peripheral Clocks Enable
func (r *RegisterApb1henrType) GetFdcanen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1henrFieldFdcanenMask) != 0
}

// SetFdcanen FDCAN Peripheral Clocks Enable
func (r *RegisterApb1henrType) SetFdcanen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1henrFieldFdcanenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1henrFieldFdcanenMask)
	}
}

// RegisterApb2enrType RCC APB2 Clock Register
type RegisterApb2enrType uint32

func (r *RegisterApb2enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb2enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb2enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb2enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb2enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb2enrFieldTim1enShift = 0
	RegisterApb2enrFieldTim1enMask  = 0x1
)

// GetTim1en TIM1 peripheral clock enable
func (r *RegisterApb2enrType) GetTim1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldTim1enMask) != 0
}

// SetTim1en TIM1 peripheral clock enable
func (r *RegisterApb2enrType) SetTim1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldTim1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldTim1enMask)
	}
}

const (
	RegisterApb2enrFieldTim8enShift = 1
	RegisterApb2enrFieldTim8enMask  = 0x2
)

// GetTim8en TIM8 peripheral clock enable
func (r *RegisterApb2enrType) GetTim8en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldTim8enMask) != 0
}

// SetTim8en TIM8 peripheral clock enable
func (r *RegisterApb2enrType) SetTim8en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldTim8enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldTim8enMask)
	}
}

const (
	RegisterApb2enrFieldUsart1enShift = 4
	RegisterApb2enrFieldUsart1enMask  = 0x10
)

// GetUsart1en USART1 Peripheral Clocks Enable
func (r *RegisterApb2enrType) GetUsart1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldUsart1enMask) != 0
}

// SetUsart1en USART1 Peripheral Clocks Enable
func (r *RegisterApb2enrType) SetUsart1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldUsart1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldUsart1enMask)
	}
}

const (
	RegisterApb2enrFieldUsart6enShift = 5
	RegisterApb2enrFieldUsart6enMask  = 0x20
)

// GetUsart6en USART6 Peripheral Clocks Enable
func (r *RegisterApb2enrType) GetUsart6en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldUsart6enMask) != 0
}

// SetUsart6en USART6 Peripheral Clocks Enable
func (r *RegisterApb2enrType) SetUsart6en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldUsart6enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldUsart6enMask)
	}
}

const (
	RegisterApb2enrFieldSpi1enShift = 12
	RegisterApb2enrFieldSpi1enMask  = 0x1000
)

// GetSpi1en SPI1 Peripheral Clocks Enable
func (r *RegisterApb2enrType) GetSpi1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldSpi1enMask) != 0
}

// SetSpi1en SPI1 Peripheral Clocks Enable
func (r *RegisterApb2enrType) SetSpi1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldSpi1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldSpi1enMask)
	}
}

const (
	RegisterApb2enrFieldSpi4enShift = 13
	RegisterApb2enrFieldSpi4enMask  = 0x2000
)

// GetSpi4en SPI4 Peripheral Clocks Enable
func (r *RegisterApb2enrType) GetSpi4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldSpi4enMask) != 0
}

// SetSpi4en SPI4 Peripheral Clocks Enable
func (r *RegisterApb2enrType) SetSpi4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldSpi4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldSpi4enMask)
	}
}

const (
	RegisterApb2enrFieldTim15enShift = 16
	RegisterApb2enrFieldTim15enMask  = 0x10000
)

// GetTim15en TIM15 peripheral clock enable
func (r *RegisterApb2enrType) GetTim15en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldTim15enMask) != 0
}

// SetTim15en TIM15 peripheral clock enable
func (r *RegisterApb2enrType) SetTim15en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldTim15enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldTim15enMask)
	}
}

const (
	RegisterApb2enrFieldTim16enShift = 17
	RegisterApb2enrFieldTim16enMask  = 0x20000
)

// GetTim16en TIM16 peripheral clock enable
func (r *RegisterApb2enrType) GetTim16en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldTim16enMask) != 0
}

// SetTim16en TIM16 peripheral clock enable
func (r *RegisterApb2enrType) SetTim16en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldTim16enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldTim16enMask)
	}
}

const (
	RegisterApb2enrFieldTim17enShift = 18
	RegisterApb2enrFieldTim17enMask  = 0x40000
)

// GetTim17en TIM17 peripheral clock enable
func (r *RegisterApb2enrType) GetTim17en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldTim17enMask) != 0
}

// SetTim17en TIM17 peripheral clock enable
func (r *RegisterApb2enrType) SetTim17en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldTim17enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldTim17enMask)
	}
}

const (
	RegisterApb2enrFieldSpi5enShift = 20
	RegisterApb2enrFieldSpi5enMask  = 0x100000
)

// GetSpi5en SPI5 Peripheral Clocks Enable
func (r *RegisterApb2enrType) GetSpi5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldSpi5enMask) != 0
}

// SetSpi5en SPI5 Peripheral Clocks Enable
func (r *RegisterApb2enrType) SetSpi5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldSpi5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldSpi5enMask)
	}
}

const (
	RegisterApb2enrFieldSai1enShift = 22
	RegisterApb2enrFieldSai1enMask  = 0x400000
)

// GetSai1en SAI1 Peripheral Clocks Enable
func (r *RegisterApb2enrType) GetSai1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldSai1enMask) != 0
}

// SetSai1en SAI1 Peripheral Clocks Enable
func (r *RegisterApb2enrType) SetSai1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldSai1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldSai1enMask)
	}
}

const (
	RegisterApb2enrFieldSai2enShift = 23
	RegisterApb2enrFieldSai2enMask  = 0x800000
)

// GetSai2en SAI2 Peripheral Clocks Enable
func (r *RegisterApb2enrType) GetSai2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldSai2enMask) != 0
}

// SetSai2en SAI2 Peripheral Clocks Enable
func (r *RegisterApb2enrType) SetSai2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldSai2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldSai2enMask)
	}
}

const (
	RegisterApb2enrFieldSai3enShift = 24
	RegisterApb2enrFieldSai3enMask  = 0x1000000
)

// GetSai3en SAI3 Peripheral Clocks Enable
func (r *RegisterApb2enrType) GetSai3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldSai3enMask) != 0
}

// SetSai3en SAI3 Peripheral Clocks Enable
func (r *RegisterApb2enrType) SetSai3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldSai3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldSai3enMask)
	}
}

const (
	RegisterApb2enrFieldDfsdm1enShift = 28
	RegisterApb2enrFieldDfsdm1enMask  = 0x10000000
)

// GetDfsdm1en DFSDM1 Peripheral Clocks Enable
func (r *RegisterApb2enrType) GetDfsdm1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldDfsdm1enMask) != 0
}

// SetDfsdm1en DFSDM1 Peripheral Clocks Enable
func (r *RegisterApb2enrType) SetDfsdm1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldDfsdm1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldDfsdm1enMask)
	}
}

const (
	RegisterApb2enrFieldHrtimenShift = 29
	RegisterApb2enrFieldHrtimenMask  = 0x20000000
)

// GetHrtimen HRTIM peripheral clock enable
func (r *RegisterApb2enrType) GetHrtimen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2enrFieldHrtimenMask) != 0
}

// SetHrtimen HRTIM peripheral clock enable
func (r *RegisterApb2enrType) SetHrtimen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2enrFieldHrtimenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2enrFieldHrtimenMask)
	}
}

// RegisterApb4enrType RCC APB4 Clock Register
type RegisterApb4enrType uint32

func (r *RegisterApb4enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb4enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb4enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb4enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb4enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb4enrFieldSyscfgenShift = 1
	RegisterApb4enrFieldSyscfgenMask  = 0x2
)

// GetSyscfgen SYSCFG peripheral clock enable
func (r *RegisterApb4enrType) GetSyscfgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldSyscfgenMask) != 0
}

// SetSyscfgen SYSCFG peripheral clock enable
func (r *RegisterApb4enrType) SetSyscfgen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldSyscfgenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldSyscfgenMask)
	}
}

const (
	RegisterApb4enrFieldLpuart1enShift = 3
	RegisterApb4enrFieldLpuart1enMask  = 0x8
)

// GetLpuart1en LPUART1 Peripheral Clocks Enable
func (r *RegisterApb4enrType) GetLpuart1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldLpuart1enMask) != 0
}

// SetLpuart1en LPUART1 Peripheral Clocks Enable
func (r *RegisterApb4enrType) SetLpuart1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldLpuart1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldLpuart1enMask)
	}
}

const (
	RegisterApb4enrFieldSpi6enShift = 5
	RegisterApb4enrFieldSpi6enMask  = 0x20
)

// GetSpi6en SPI6 Peripheral Clocks Enable
func (r *RegisterApb4enrType) GetSpi6en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldSpi6enMask) != 0
}

// SetSpi6en SPI6 Peripheral Clocks Enable
func (r *RegisterApb4enrType) SetSpi6en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldSpi6enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldSpi6enMask)
	}
}

const (
	RegisterApb4enrFieldI2c4enShift = 7
	RegisterApb4enrFieldI2c4enMask  = 0x80
)

// GetI2c4en I2C4 Peripheral Clocks Enable
func (r *RegisterApb4enrType) GetI2c4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldI2c4enMask) != 0
}

// SetI2c4en I2C4 Peripheral Clocks Enable
func (r *RegisterApb4enrType) SetI2c4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldI2c4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldI2c4enMask)
	}
}

const (
	RegisterApb4enrFieldLptim2enShift = 9
	RegisterApb4enrFieldLptim2enMask  = 0x200
)

// GetLptim2en LPTIM2 Peripheral Clocks Enable
func (r *RegisterApb4enrType) GetLptim2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldLptim2enMask) != 0
}

// SetLptim2en LPTIM2 Peripheral Clocks Enable
func (r *RegisterApb4enrType) SetLptim2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldLptim2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldLptim2enMask)
	}
}

const (
	RegisterApb4enrFieldLptim3enShift = 10
	RegisterApb4enrFieldLptim3enMask  = 0x400
)

// GetLptim3en LPTIM3 Peripheral Clocks Enable
func (r *RegisterApb4enrType) GetLptim3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldLptim3enMask) != 0
}

// SetLptim3en LPTIM3 Peripheral Clocks Enable
func (r *RegisterApb4enrType) SetLptim3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldLptim3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldLptim3enMask)
	}
}

const (
	RegisterApb4enrFieldLptim4enShift = 11
	RegisterApb4enrFieldLptim4enMask  = 0x800
)

// GetLptim4en LPTIM4 Peripheral Clocks Enable
func (r *RegisterApb4enrType) GetLptim4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldLptim4enMask) != 0
}

// SetLptim4en LPTIM4 Peripheral Clocks Enable
func (r *RegisterApb4enrType) SetLptim4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldLptim4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldLptim4enMask)
	}
}

const (
	RegisterApb4enrFieldLptim5enShift = 12
	RegisterApb4enrFieldLptim5enMask  = 0x1000
)

// GetLptim5en LPTIM5 Peripheral Clocks Enable
func (r *RegisterApb4enrType) GetLptim5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldLptim5enMask) != 0
}

// SetLptim5en LPTIM5 Peripheral Clocks Enable
func (r *RegisterApb4enrType) SetLptim5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldLptim5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldLptim5enMask)
	}
}

const (
	RegisterApb4enrFieldComp12enShift = 14
	RegisterApb4enrFieldComp12enMask  = 0x4000
)

// GetComp12en COMP1/2 peripheral clock enable
func (r *RegisterApb4enrType) GetComp12en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldComp12enMask) != 0
}

// SetComp12en COMP1/2 peripheral clock enable
func (r *RegisterApb4enrType) SetComp12en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldComp12enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldComp12enMask)
	}
}

const (
	RegisterApb4enrFieldVrefenShift = 15
	RegisterApb4enrFieldVrefenMask  = 0x8000
)

// GetVrefen VREF peripheral clock enable
func (r *RegisterApb4enrType) GetVrefen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldVrefenMask) != 0
}

// SetVrefen VREF peripheral clock enable
func (r *RegisterApb4enrType) SetVrefen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldVrefenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldVrefenMask)
	}
}

const (
	RegisterApb4enrFieldRtcapbenShift = 16
	RegisterApb4enrFieldRtcapbenMask  = 0x10000
)

// GetRtcapben RTC APB Clock Enable
func (r *RegisterApb4enrType) GetRtcapben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldRtcapbenMask) != 0
}

// SetRtcapben RTC APB Clock Enable
func (r *RegisterApb4enrType) SetRtcapben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldRtcapbenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldRtcapbenMask)
	}
}

const (
	RegisterApb4enrFieldSai4enShift = 21
	RegisterApb4enrFieldSai4enMask  = 0x200000
)

// GetSai4en SAI4 Peripheral Clocks Enable
func (r *RegisterApb4enrType) GetSai4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4enrFieldSai4enMask) != 0
}

// SetSai4en SAI4 Peripheral Clocks Enable
func (r *RegisterApb4enrType) SetSai4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4enrFieldSai4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4enrFieldSai4enMask)
	}
}

// RegisterAhb3lpenrType RCC AHB3 Sleep Clock Register
type RegisterAhb3lpenrType uint32

func (r *RegisterAhb3lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb3lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb3lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb3lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb3lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb3lpenrFieldMdmalpenShift = 0
	RegisterAhb3lpenrFieldMdmalpenMask  = 0x1
)

// GetMdmalpen MDMA Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) GetMdmalpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3lpenrFieldMdmalpenMask) != 0
}

// SetMdmalpen MDMA Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) SetMdmalpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3lpenrFieldMdmalpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3lpenrFieldMdmalpenMask)
	}
}

const (
	RegisterAhb3lpenrFieldDma2dlpenShift = 4
	RegisterAhb3lpenrFieldDma2dlpenMask  = 0x10
)

// GetDma2dlpen DMA2D Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) GetDma2dlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3lpenrFieldDma2dlpenMask) != 0
}

// SetDma2dlpen DMA2D Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) SetDma2dlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3lpenrFieldDma2dlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3lpenrFieldDma2dlpenMask)
	}
}

const (
	RegisterAhb3lpenrFieldJpgdeclpenShift = 5
	RegisterAhb3lpenrFieldJpgdeclpenMask  = 0x20
)

// GetJpgdeclpen JPGDEC Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) GetJpgdeclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3lpenrFieldJpgdeclpenMask) != 0
}

// SetJpgdeclpen JPGDEC Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) SetJpgdeclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3lpenrFieldJpgdeclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3lpenrFieldJpgdeclpenMask)
	}
}

const (
	RegisterAhb3lpenrFieldFlitflpenShift = 8
	RegisterAhb3lpenrFieldFlitflpenMask  = 0x100
)

// GetFlitflpen FLITF Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) GetFlitflpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3lpenrFieldFlitflpenMask) != 0
}

// SetFlitflpen FLITF Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) SetFlitflpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3lpenrFieldFlitflpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3lpenrFieldFlitflpenMask)
	}
}

const (
	RegisterAhb3lpenrFieldFmclpenShift = 12
	RegisterAhb3lpenrFieldFmclpenMask  = 0x1000
)

// GetFmclpen FMC Peripheral Clocks Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) GetFmclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3lpenrFieldFmclpenMask) != 0
}

// SetFmclpen FMC Peripheral Clocks Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) SetFmclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3lpenrFieldFmclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3lpenrFieldFmclpenMask)
	}
}

const (
	RegisterAhb3lpenrFieldQspilpenShift = 14
	RegisterAhb3lpenrFieldQspilpenMask  = 0x4000
)

// GetQspilpen QUADSPI and QUADSPI Delay Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) GetQspilpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3lpenrFieldQspilpenMask) != 0
}

// SetQspilpen QUADSPI and QUADSPI Delay Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) SetQspilpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3lpenrFieldQspilpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3lpenrFieldQspilpenMask)
	}
}

const (
	RegisterAhb3lpenrFieldSdmmc1lpenShift = 16
	RegisterAhb3lpenrFieldSdmmc1lpenMask  = 0x10000
)

// GetSdmmc1lpen SDMMC1 and SDMMC1 Delay Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) GetSdmmc1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3lpenrFieldSdmmc1lpenMask) != 0
}

// SetSdmmc1lpen SDMMC1 and SDMMC1 Delay Clock Enable During CSleep Mode
func (r *RegisterAhb3lpenrType) SetSdmmc1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3lpenrFieldSdmmc1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3lpenrFieldSdmmc1lpenMask)
	}
}

const (
	RegisterAhb3lpenrFieldD1dtcm1lpenShift = 28
	RegisterAhb3lpenrFieldD1dtcm1lpenMask  = 0x10000000
)

// GetD1dtcm1lpen D1DTCM1 Block Clock Enable During CSleep mode
func (r *RegisterAhb3lpenrType) GetD1dtcm1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3lpenrFieldD1dtcm1lpenMask) != 0
}

// SetD1dtcm1lpen D1DTCM1 Block Clock Enable During CSleep mode
func (r *RegisterAhb3lpenrType) SetD1dtcm1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3lpenrFieldD1dtcm1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3lpenrFieldD1dtcm1lpenMask)
	}
}

const (
	RegisterAhb3lpenrFieldDtcm2lpenShift = 29
	RegisterAhb3lpenrFieldDtcm2lpenMask  = 0x20000000
)

// GetDtcm2lpen D1 DTCM2 Block Clock Enable During CSleep mode
func (r *RegisterAhb3lpenrType) GetDtcm2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3lpenrFieldDtcm2lpenMask) != 0
}

// SetDtcm2lpen D1 DTCM2 Block Clock Enable During CSleep mode
func (r *RegisterAhb3lpenrType) SetDtcm2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3lpenrFieldDtcm2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3lpenrFieldDtcm2lpenMask)
	}
}

const (
	RegisterAhb3lpenrFieldItcmlpenShift = 30
	RegisterAhb3lpenrFieldItcmlpenMask  = 0x40000000
)

// GetItcmlpen D1ITCM Block Clock Enable During CSleep mode
func (r *RegisterAhb3lpenrType) GetItcmlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3lpenrFieldItcmlpenMask) != 0
}

// SetItcmlpen D1ITCM Block Clock Enable During CSleep mode
func (r *RegisterAhb3lpenrType) SetItcmlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3lpenrFieldItcmlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3lpenrFieldItcmlpenMask)
	}
}

const (
	RegisterAhb3lpenrFieldAxisramlpenShift = 31
	RegisterAhb3lpenrFieldAxisramlpenMask  = 0x80000000
)

// GetAxisramlpen AXISRAM Block Clock Enable During CSleep mode
func (r *RegisterAhb3lpenrType) GetAxisramlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb3lpenrFieldAxisramlpenMask) != 0
}

// SetAxisramlpen AXISRAM Block Clock Enable During CSleep mode
func (r *RegisterAhb3lpenrType) SetAxisramlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb3lpenrFieldAxisramlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb3lpenrFieldAxisramlpenMask)
	}
}

// RegisterAhb1lpenrType RCC AHB1 Sleep Clock Register
type RegisterAhb1lpenrType uint32

func (r *RegisterAhb1lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb1lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb1lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb1lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb1lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb1lpenrFieldDma1lpenShift = 0
	RegisterAhb1lpenrFieldDma1lpenMask  = 0x1
)

// GetDma1lpen DMA1 Clock Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) GetDma1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1lpenrFieldDma1lpenMask) != 0
}

// SetDma1lpen DMA1 Clock Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) SetDma1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1lpenrFieldDma1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1lpenrFieldDma1lpenMask)
	}
}

const (
	RegisterAhb1lpenrFieldDma2lpenShift = 1
	RegisterAhb1lpenrFieldDma2lpenMask  = 0x2
)

// GetDma2lpen DMA2 Clock Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) GetDma2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1lpenrFieldDma2lpenMask) != 0
}

// SetDma2lpen DMA2 Clock Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) SetDma2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1lpenrFieldDma2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1lpenrFieldDma2lpenMask)
	}
}

const (
	RegisterAhb1lpenrFieldAdc12lpenShift = 5
	RegisterAhb1lpenrFieldAdc12lpenMask  = 0x20
)

// GetAdc12lpen ADC1/2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) GetAdc12lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1lpenrFieldAdc12lpenMask) != 0
}

// SetAdc12lpen ADC1/2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) SetAdc12lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1lpenrFieldAdc12lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1lpenrFieldAdc12lpenMask)
	}
}

const (
	RegisterAhb1lpenrFieldEth1maclpenShift = 15
	RegisterAhb1lpenrFieldEth1maclpenMask  = 0x8000
)

// GetEth1maclpen Ethernet MAC bus interface Clock Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) GetEth1maclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1lpenrFieldEth1maclpenMask) != 0
}

// SetEth1maclpen Ethernet MAC bus interface Clock Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) SetEth1maclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1lpenrFieldEth1maclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1lpenrFieldEth1maclpenMask)
	}
}

const (
	RegisterAhb1lpenrFieldEth1txlpenShift = 16
	RegisterAhb1lpenrFieldEth1txlpenMask  = 0x10000
)

// GetEth1txlpen Ethernet Transmission Clock Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) GetEth1txlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1lpenrFieldEth1txlpenMask) != 0
}

// SetEth1txlpen Ethernet Transmission Clock Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) SetEth1txlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1lpenrFieldEth1txlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1lpenrFieldEth1txlpenMask)
	}
}

const (
	RegisterAhb1lpenrFieldEth1rxlpenShift = 17
	RegisterAhb1lpenrFieldEth1rxlpenMask  = 0x20000
)

// GetEth1rxlpen Ethernet Reception Clock Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) GetEth1rxlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1lpenrFieldEth1rxlpenMask) != 0
}

// SetEth1rxlpen Ethernet Reception Clock Enable During CSleep Mode
func (r *RegisterAhb1lpenrType) SetEth1rxlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1lpenrFieldEth1rxlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1lpenrFieldEth1rxlpenMask)
	}
}

const (
	RegisterAhb1lpenrFieldUsb1otglpenShift = 25
	RegisterAhb1lpenrFieldUsb1otglpenMask  = 0x2000000
)

// GetUsb1otglpen USB1OTG peripheral clock enable during CSleep mode
func (r *RegisterAhb1lpenrType) GetUsb1otglpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1lpenrFieldUsb1otglpenMask) != 0
}

// SetUsb1otglpen USB1OTG peripheral clock enable during CSleep mode
func (r *RegisterAhb1lpenrType) SetUsb1otglpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1lpenrFieldUsb1otglpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1lpenrFieldUsb1otglpenMask)
	}
}

const (
	RegisterAhb1lpenrFieldUsb1ulpilpenShift = 26
	RegisterAhb1lpenrFieldUsb1ulpilpenMask  = 0x4000000
)

// GetUsb1ulpilpen USB_PHY1 clock enable during CSleep mode
func (r *RegisterAhb1lpenrType) GetUsb1ulpilpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1lpenrFieldUsb1ulpilpenMask) != 0
}

// SetUsb1ulpilpen USB_PHY1 clock enable during CSleep mode
func (r *RegisterAhb1lpenrType) SetUsb1ulpilpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1lpenrFieldUsb1ulpilpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1lpenrFieldUsb1ulpilpenMask)
	}
}

const (
	RegisterAhb1lpenrFieldUsb2otglpenShift = 27
	RegisterAhb1lpenrFieldUsb2otglpenMask  = 0x8000000
)

// GetUsb2otglpen USB2OTG peripheral clock enable during CSleep mode
func (r *RegisterAhb1lpenrType) GetUsb2otglpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1lpenrFieldUsb2otglpenMask) != 0
}

// SetUsb2otglpen USB2OTG peripheral clock enable during CSleep mode
func (r *RegisterAhb1lpenrType) SetUsb2otglpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1lpenrFieldUsb2otglpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1lpenrFieldUsb2otglpenMask)
	}
}

const (
	RegisterAhb1lpenrFieldUsb2ulpilpenShift = 28
	RegisterAhb1lpenrFieldUsb2ulpilpenMask  = 0x10000000
)

// GetUsb2ulpilpen USB_PHY2 clocks enable during CSleep mode
func (r *RegisterAhb1lpenrType) GetUsb2ulpilpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb1lpenrFieldUsb2ulpilpenMask) != 0
}

// SetUsb2ulpilpen USB_PHY2 clocks enable during CSleep mode
func (r *RegisterAhb1lpenrType) SetUsb2ulpilpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb1lpenrFieldUsb2ulpilpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb1lpenrFieldUsb2ulpilpenMask)
	}
}

// RegisterAhb2lpenrType RCC AHB2 Sleep Clock Register
type RegisterAhb2lpenrType uint32

func (r *RegisterAhb2lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb2lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb2lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb2lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb2lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb2lpenrFieldCamitflpenShift = 0
	RegisterAhb2lpenrFieldCamitflpenMask  = 0x1
)

// GetCamitflpen CAMITF peripheral clock enable during CSleep mode
func (r *RegisterAhb2lpenrType) GetCamitflpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2lpenrFieldCamitflpenMask) != 0
}

// SetCamitflpen CAMITF peripheral clock enable during CSleep mode
func (r *RegisterAhb2lpenrType) SetCamitflpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2lpenrFieldCamitflpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2lpenrFieldCamitflpenMask)
	}
}

const (
	RegisterAhb2lpenrFieldCryptlpenShift = 4
	RegisterAhb2lpenrFieldCryptlpenMask  = 0x10
)

// GetCryptlpen CRYPT peripheral clock enable during CSleep mode
func (r *RegisterAhb2lpenrType) GetCryptlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2lpenrFieldCryptlpenMask) != 0
}

// SetCryptlpen CRYPT peripheral clock enable during CSleep mode
func (r *RegisterAhb2lpenrType) SetCryptlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2lpenrFieldCryptlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2lpenrFieldCryptlpenMask)
	}
}

const (
	RegisterAhb2lpenrFieldHashlpenShift = 5
	RegisterAhb2lpenrFieldHashlpenMask  = 0x20
)

// GetHashlpen HASH peripheral clock enable during CSleep mode
func (r *RegisterAhb2lpenrType) GetHashlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2lpenrFieldHashlpenMask) != 0
}

// SetHashlpen HASH peripheral clock enable during CSleep mode
func (r *RegisterAhb2lpenrType) SetHashlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2lpenrFieldHashlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2lpenrFieldHashlpenMask)
	}
}

const (
	RegisterAhb2lpenrFieldRnglpenShift = 6
	RegisterAhb2lpenrFieldRnglpenMask  = 0x40
)

// GetRnglpen RNG peripheral clock enable during CSleep mode
func (r *RegisterAhb2lpenrType) GetRnglpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2lpenrFieldRnglpenMask) != 0
}

// SetRnglpen RNG peripheral clock enable during CSleep mode
func (r *RegisterAhb2lpenrType) SetRnglpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2lpenrFieldRnglpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2lpenrFieldRnglpenMask)
	}
}

const (
	RegisterAhb2lpenrFieldSdmmc2lpenShift = 9
	RegisterAhb2lpenrFieldSdmmc2lpenMask  = 0x200
)

// GetSdmmc2lpen SDMMC2 and SDMMC2 Delay Clock Enable During CSleep Mode
func (r *RegisterAhb2lpenrType) GetSdmmc2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2lpenrFieldSdmmc2lpenMask) != 0
}

// SetSdmmc2lpen SDMMC2 and SDMMC2 Delay Clock Enable During CSleep Mode
func (r *RegisterAhb2lpenrType) SetSdmmc2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2lpenrFieldSdmmc2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2lpenrFieldSdmmc2lpenMask)
	}
}

const (
	RegisterAhb2lpenrFieldSram1lpenShift = 29
	RegisterAhb2lpenrFieldSram1lpenMask  = 0x20000000
)

// GetSram1lpen SRAM1 Clock Enable During CSleep Mode
func (r *RegisterAhb2lpenrType) GetSram1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2lpenrFieldSram1lpenMask) != 0
}

// SetSram1lpen SRAM1 Clock Enable During CSleep Mode
func (r *RegisterAhb2lpenrType) SetSram1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2lpenrFieldSram1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2lpenrFieldSram1lpenMask)
	}
}

const (
	RegisterAhb2lpenrFieldSram2lpenShift = 30
	RegisterAhb2lpenrFieldSram2lpenMask  = 0x40000000
)

// GetSram2lpen SRAM2 Clock Enable During CSleep Mode
func (r *RegisterAhb2lpenrType) GetSram2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2lpenrFieldSram2lpenMask) != 0
}

// SetSram2lpen SRAM2 Clock Enable During CSleep Mode
func (r *RegisterAhb2lpenrType) SetSram2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2lpenrFieldSram2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2lpenrFieldSram2lpenMask)
	}
}

const (
	RegisterAhb2lpenrFieldSram3lpenShift = 31
	RegisterAhb2lpenrFieldSram3lpenMask  = 0x80000000
)

// GetSram3lpen SRAM3 Clock Enable During CSleep Mode
func (r *RegisterAhb2lpenrType) GetSram3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb2lpenrFieldSram3lpenMask) != 0
}

// SetSram3lpen SRAM3 Clock Enable During CSleep Mode
func (r *RegisterAhb2lpenrType) SetSram3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb2lpenrFieldSram3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb2lpenrFieldSram3lpenMask)
	}
}

// RegisterAhb4lpenrType RCC AHB4 Sleep Clock Register
type RegisterAhb4lpenrType uint32

func (r *RegisterAhb4lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterAhb4lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterAhb4lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterAhb4lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterAhb4lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterAhb4lpenrFieldGpioalpenShift = 0
	RegisterAhb4lpenrFieldGpioalpenMask  = 0x1
)

// GetGpioalpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetGpioalpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldGpioalpenMask) != 0
}

// SetGpioalpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetGpioalpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldGpioalpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldGpioalpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldGpioblpenShift = 1
	RegisterAhb4lpenrFieldGpioblpenMask  = 0x2
)

// GetGpioblpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetGpioblpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldGpioblpenMask) != 0
}

// SetGpioblpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetGpioblpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldGpioblpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldGpioblpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldGpioclpenShift = 2
	RegisterAhb4lpenrFieldGpioclpenMask  = 0x4
)

// GetGpioclpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetGpioclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldGpioclpenMask) != 0
}

// SetGpioclpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetGpioclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldGpioclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldGpioclpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldGpiodlpenShift = 3
	RegisterAhb4lpenrFieldGpiodlpenMask  = 0x8
)

// GetGpiodlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetGpiodlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldGpiodlpenMask) != 0
}

// SetGpiodlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetGpiodlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldGpiodlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldGpiodlpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldGpioelpenShift = 4
	RegisterAhb4lpenrFieldGpioelpenMask  = 0x10
)

// GetGpioelpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetGpioelpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldGpioelpenMask) != 0
}

// SetGpioelpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetGpioelpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldGpioelpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldGpioelpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldGpioflpenShift = 5
	RegisterAhb4lpenrFieldGpioflpenMask  = 0x20
)

// GetGpioflpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetGpioflpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldGpioflpenMask) != 0
}

// SetGpioflpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetGpioflpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldGpioflpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldGpioflpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldGpioglpenShift = 6
	RegisterAhb4lpenrFieldGpioglpenMask  = 0x40
)

// GetGpioglpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetGpioglpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldGpioglpenMask) != 0
}

// SetGpioglpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetGpioglpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldGpioglpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldGpioglpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldGpiohlpenShift = 7
	RegisterAhb4lpenrFieldGpiohlpenMask  = 0x80
)

// GetGpiohlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetGpiohlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldGpiohlpenMask) != 0
}

// SetGpiohlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetGpiohlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldGpiohlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldGpiohlpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldGpioilpenShift = 8
	RegisterAhb4lpenrFieldGpioilpenMask  = 0x100
)

// GetGpioilpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetGpioilpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldGpioilpenMask) != 0
}

// SetGpioilpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetGpioilpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldGpioilpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldGpioilpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldGpiojlpenShift = 9
	RegisterAhb4lpenrFieldGpiojlpenMask  = 0x200
)

// GetGpiojlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetGpiojlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldGpiojlpenMask) != 0
}

// SetGpiojlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetGpiojlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldGpiojlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldGpiojlpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldGpioklpenShift = 10
	RegisterAhb4lpenrFieldGpioklpenMask  = 0x400
)

// GetGpioklpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetGpioklpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldGpioklpenMask) != 0
}

// SetGpioklpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetGpioklpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldGpioklpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldGpioklpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldCrclpenShift = 19
	RegisterAhb4lpenrFieldCrclpenMask  = 0x80000
)

// GetCrclpen CRC peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) GetCrclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldCrclpenMask) != 0
}

// SetCrclpen CRC peripheral clock enable during CSleep mode
func (r *RegisterAhb4lpenrType) SetCrclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldCrclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldCrclpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldBdmalpenShift = 21
	RegisterAhb4lpenrFieldBdmalpenMask  = 0x200000
)

// GetBdmalpen BDMA Clock Enable During CSleep Mode
func (r *RegisterAhb4lpenrType) GetBdmalpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldBdmalpenMask) != 0
}

// SetBdmalpen BDMA Clock Enable During CSleep Mode
func (r *RegisterAhb4lpenrType) SetBdmalpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldBdmalpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldBdmalpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldAdc3lpenShift = 24
	RegisterAhb4lpenrFieldAdc3lpenMask  = 0x1000000
)

// GetAdc3lpen ADC3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterAhb4lpenrType) GetAdc3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldAdc3lpenMask) != 0
}

// SetAdc3lpen ADC3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterAhb4lpenrType) SetAdc3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldAdc3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldAdc3lpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldBkpramlpenShift = 28
	RegisterAhb4lpenrFieldBkpramlpenMask  = 0x10000000
)

// GetBkpramlpen Backup RAM Clock Enable During CSleep Mode
func (r *RegisterAhb4lpenrType) GetBkpramlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldBkpramlpenMask) != 0
}

// SetBkpramlpen Backup RAM Clock Enable During CSleep Mode
func (r *RegisterAhb4lpenrType) SetBkpramlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldBkpramlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldBkpramlpenMask)
	}
}

const (
	RegisterAhb4lpenrFieldSram4lpenShift = 29
	RegisterAhb4lpenrFieldSram4lpenMask  = 0x20000000
)

// GetSram4lpen SRAM4 Clock Enable During CSleep Mode
func (r *RegisterAhb4lpenrType) GetSram4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterAhb4lpenrFieldSram4lpenMask) != 0
}

// SetSram4lpen SRAM4 Clock Enable During CSleep Mode
func (r *RegisterAhb4lpenrType) SetSram4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterAhb4lpenrFieldSram4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterAhb4lpenrFieldSram4lpenMask)
	}
}

// RegisterApb3lpenrType RCC APB3 Sleep Clock Register
type RegisterApb3lpenrType uint32

func (r *RegisterApb3lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb3lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb3lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb3lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb3lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb3lpenrFieldLtdclpenShift = 3
	RegisterApb3lpenrFieldLtdclpenMask  = 0x8
)

// GetLtdclpen LTDC peripheral clock enable during CSleep mode
func (r *RegisterApb3lpenrType) GetLtdclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb3lpenrFieldLtdclpenMask) != 0
}

// SetLtdclpen LTDC peripheral clock enable during CSleep mode
func (r *RegisterApb3lpenrType) SetLtdclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb3lpenrFieldLtdclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb3lpenrFieldLtdclpenMask)
	}
}

const (
	RegisterApb3lpenrFieldWwdg1lpenShift = 6
	RegisterApb3lpenrFieldWwdg1lpenMask  = 0x40
)

// GetWwdg1lpen WWDG1 Clock Enable During CSleep Mode
func (r *RegisterApb3lpenrType) GetWwdg1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb3lpenrFieldWwdg1lpenMask) != 0
}

// SetWwdg1lpen WWDG1 Clock Enable During CSleep Mode
func (r *RegisterApb3lpenrType) SetWwdg1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb3lpenrFieldWwdg1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb3lpenrFieldWwdg1lpenMask)
	}
}

// RegisterApb1llpenrType RCC APB1 Low Sleep Clock Register
type RegisterApb1llpenrType uint32

func (r *RegisterApb1llpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb1llpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb1llpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb1llpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb1llpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb1llpenrFieldTim2lpenShift = 0
	RegisterApb1llpenrFieldTim2lpenMask  = 0x1
)

// GetTim2lpen TIM2 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) GetTim2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldTim2lpenMask) != 0
}

// SetTim2lpen TIM2 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) SetTim2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldTim2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldTim2lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldTim3lpenShift = 1
	RegisterApb1llpenrFieldTim3lpenMask  = 0x2
)

// GetTim3lpen TIM3 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) GetTim3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldTim3lpenMask) != 0
}

// SetTim3lpen TIM3 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) SetTim3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldTim3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldTim3lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldTim4lpenShift = 2
	RegisterApb1llpenrFieldTim4lpenMask  = 0x4
)

// GetTim4lpen TIM4 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) GetTim4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldTim4lpenMask) != 0
}

// SetTim4lpen TIM4 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) SetTim4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldTim4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldTim4lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldTim5lpenShift = 3
	RegisterApb1llpenrFieldTim5lpenMask  = 0x8
)

// GetTim5lpen TIM5 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) GetTim5lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldTim5lpenMask) != 0
}

// SetTim5lpen TIM5 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) SetTim5lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldTim5lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldTim5lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldTim6lpenShift = 4
	RegisterApb1llpenrFieldTim6lpenMask  = 0x10
)

// GetTim6lpen TIM6 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) GetTim6lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldTim6lpenMask) != 0
}

// SetTim6lpen TIM6 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) SetTim6lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldTim6lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldTim6lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldTim7lpenShift = 5
	RegisterApb1llpenrFieldTim7lpenMask  = 0x20
)

// GetTim7lpen TIM7 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) GetTim7lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldTim7lpenMask) != 0
}

// SetTim7lpen TIM7 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) SetTim7lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldTim7lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldTim7lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldTim12lpenShift = 6
	RegisterApb1llpenrFieldTim12lpenMask  = 0x40
)

// GetTim12lpen TIM12 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) GetTim12lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldTim12lpenMask) != 0
}

// SetTim12lpen TIM12 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) SetTim12lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldTim12lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldTim12lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldTim13lpenShift = 7
	RegisterApb1llpenrFieldTim13lpenMask  = 0x80
)

// GetTim13lpen TIM13 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) GetTim13lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldTim13lpenMask) != 0
}

// SetTim13lpen TIM13 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) SetTim13lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldTim13lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldTim13lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldTim14lpenShift = 8
	RegisterApb1llpenrFieldTim14lpenMask  = 0x100
)

// GetTim14lpen TIM14 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) GetTim14lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldTim14lpenMask) != 0
}

// SetTim14lpen TIM14 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) SetTim14lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldTim14lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldTim14lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldLptim1lpenShift = 9
	RegisterApb1llpenrFieldLptim1lpenMask  = 0x200
)

// GetLptim1lpen LPTIM1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetLptim1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldLptim1lpenMask) != 0
}

// SetLptim1lpen LPTIM1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetLptim1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldLptim1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldLptim1lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldSpi2lpenShift = 14
	RegisterApb1llpenrFieldSpi2lpenMask  = 0x4000
)

// GetSpi2lpen SPI2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetSpi2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldSpi2lpenMask) != 0
}

// SetSpi2lpen SPI2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetSpi2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldSpi2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldSpi2lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldSpi3lpenShift = 15
	RegisterApb1llpenrFieldSpi3lpenMask  = 0x8000
)

// GetSpi3lpen SPI3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetSpi3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldSpi3lpenMask) != 0
}

// SetSpi3lpen SPI3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetSpi3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldSpi3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldSpi3lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldSpdifrxlpenShift = 16
	RegisterApb1llpenrFieldSpdifrxlpenMask  = 0x10000
)

// GetSpdifrxlpen SPDIFRX Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetSpdifrxlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldSpdifrxlpenMask) != 0
}

// SetSpdifrxlpen SPDIFRX Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetSpdifrxlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldSpdifrxlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldSpdifrxlpenMask)
	}
}

const (
	RegisterApb1llpenrFieldUsart2lpenShift = 17
	RegisterApb1llpenrFieldUsart2lpenMask  = 0x20000
)

// GetUsart2lpen USART2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetUsart2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldUsart2lpenMask) != 0
}

// SetUsart2lpen USART2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetUsart2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldUsart2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldUsart2lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldUsart3lpenShift = 18
	RegisterApb1llpenrFieldUsart3lpenMask  = 0x40000
)

// GetUsart3lpen USART3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetUsart3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldUsart3lpenMask) != 0
}

// SetUsart3lpen USART3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetUsart3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldUsart3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldUsart3lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldUart4lpenShift = 19
	RegisterApb1llpenrFieldUart4lpenMask  = 0x80000
)

// GetUart4lpen UART4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetUart4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldUart4lpenMask) != 0
}

// SetUart4lpen UART4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetUart4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldUart4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldUart4lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldUart5lpenShift = 20
	RegisterApb1llpenrFieldUart5lpenMask  = 0x100000
)

// GetUart5lpen UART5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetUart5lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldUart5lpenMask) != 0
}

// SetUart5lpen UART5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetUart5lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldUart5lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldUart5lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldI2c1lpenShift = 21
	RegisterApb1llpenrFieldI2c1lpenMask  = 0x200000
)

// GetI2c1lpen I2C1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetI2c1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldI2c1lpenMask) != 0
}

// SetI2c1lpen I2C1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetI2c1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldI2c1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldI2c1lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldI2c2lpenShift = 22
	RegisterApb1llpenrFieldI2c2lpenMask  = 0x400000
)

// GetI2c2lpen I2C2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetI2c2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldI2c2lpenMask) != 0
}

// SetI2c2lpen I2C2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetI2c2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldI2c2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldI2c2lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldI2c3lpenShift = 23
	RegisterApb1llpenrFieldI2c3lpenMask  = 0x800000
)

// GetI2c3lpen I2C3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetI2c3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldI2c3lpenMask) != 0
}

// SetI2c3lpen I2C3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetI2c3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldI2c3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldI2c3lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldHdmiceclpenShift = 27
	RegisterApb1llpenrFieldHdmiceclpenMask  = 0x8000000
)

// GetHdmiceclpen HDMI-CEC Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetHdmiceclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldHdmiceclpenMask) != 0
}

// SetHdmiceclpen HDMI-CEC Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetHdmiceclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldHdmiceclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldHdmiceclpenMask)
	}
}

const (
	RegisterApb1llpenrFieldDac12lpenShift = 29
	RegisterApb1llpenrFieldDac12lpenMask  = 0x20000000
)

// GetDac12lpen DAC1/2 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) GetDac12lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldDac12lpenMask) != 0
}

// SetDac12lpen DAC1/2 peripheral clock enable during CSleep mode
func (r *RegisterApb1llpenrType) SetDac12lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldDac12lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldDac12lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldUsart7lpenShift = 30
	RegisterApb1llpenrFieldUsart7lpenMask  = 0x40000000
)

// GetUsart7lpen USART7 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetUsart7lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldUsart7lpenMask) != 0
}

// SetUsart7lpen USART7 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetUsart7lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldUsart7lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldUsart7lpenMask)
	}
}

const (
	RegisterApb1llpenrFieldUsart8lpenShift = 31
	RegisterApb1llpenrFieldUsart8lpenMask  = 0x80000000
)

// GetUsart8lpen USART8 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) GetUsart8lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1llpenrFieldUsart8lpenMask) != 0
}

// SetUsart8lpen USART8 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1llpenrType) SetUsart8lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1llpenrFieldUsart8lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1llpenrFieldUsart8lpenMask)
	}
}

// RegisterApb1hlpenrType RCC APB1 High Sleep Clock Register
type RegisterApb1hlpenrType uint32

func (r *RegisterApb1hlpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb1hlpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb1hlpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb1hlpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb1hlpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb1hlpenrFieldCrslpenShift = 1
	RegisterApb1hlpenrFieldCrslpenMask  = 0x2
)

// GetCrslpen Clock Recovery System peripheral clock enable during CSleep mode
func (r *RegisterApb1hlpenrType) GetCrslpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1hlpenrFieldCrslpenMask) != 0
}

// SetCrslpen Clock Recovery System peripheral clock enable during CSleep mode
func (r *RegisterApb1hlpenrType) SetCrslpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1hlpenrFieldCrslpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1hlpenrFieldCrslpenMask)
	}
}

const (
	RegisterApb1hlpenrFieldSwplpenShift = 2
	RegisterApb1hlpenrFieldSwplpenMask  = 0x4
)

// GetSwplpen SWPMI Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1hlpenrType) GetSwplpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1hlpenrFieldSwplpenMask) != 0
}

// SetSwplpen SWPMI Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1hlpenrType) SetSwplpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1hlpenrFieldSwplpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1hlpenrFieldSwplpenMask)
	}
}

const (
	RegisterApb1hlpenrFieldOpamplpenShift = 4
	RegisterApb1hlpenrFieldOpamplpenMask  = 0x10
)

// GetOpamplpen OPAMP peripheral clock enable during CSleep mode
func (r *RegisterApb1hlpenrType) GetOpamplpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1hlpenrFieldOpamplpenMask) != 0
}

// SetOpamplpen OPAMP peripheral clock enable during CSleep mode
func (r *RegisterApb1hlpenrType) SetOpamplpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1hlpenrFieldOpamplpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1hlpenrFieldOpamplpenMask)
	}
}

const (
	RegisterApb1hlpenrFieldMdioslpenShift = 5
	RegisterApb1hlpenrFieldMdioslpenMask  = 0x20
)

// GetMdioslpen MDIOS peripheral clock enable during CSleep mode
func (r *RegisterApb1hlpenrType) GetMdioslpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1hlpenrFieldMdioslpenMask) != 0
}

// SetMdioslpen MDIOS peripheral clock enable during CSleep mode
func (r *RegisterApb1hlpenrType) SetMdioslpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1hlpenrFieldMdioslpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1hlpenrFieldMdioslpenMask)
	}
}

const (
	RegisterApb1hlpenrFieldFdcanlpenShift = 8
	RegisterApb1hlpenrFieldFdcanlpenMask  = 0x100
)

// GetFdcanlpen FDCAN Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1hlpenrType) GetFdcanlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb1hlpenrFieldFdcanlpenMask) != 0
}

// SetFdcanlpen FDCAN Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb1hlpenrType) SetFdcanlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb1hlpenrFieldFdcanlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb1hlpenrFieldFdcanlpenMask)
	}
}

// RegisterApb2lpenrType RCC APB2 Sleep Clock Register
type RegisterApb2lpenrType uint32

func (r *RegisterApb2lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb2lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb2lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb2lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb2lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb2lpenrFieldTim1lpenShift = 0
	RegisterApb2lpenrFieldTim1lpenMask  = 0x1
)

// GetTim1lpen TIM1 peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) GetTim1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldTim1lpenMask) != 0
}

// SetTim1lpen TIM1 peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) SetTim1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldTim1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldTim1lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldTim8lpenShift = 1
	RegisterApb2lpenrFieldTim8lpenMask  = 0x2
)

// GetTim8lpen TIM8 peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) GetTim8lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldTim8lpenMask) != 0
}

// SetTim8lpen TIM8 peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) SetTim8lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldTim8lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldTim8lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldUsart1lpenShift = 4
	RegisterApb2lpenrFieldUsart1lpenMask  = 0x10
)

// GetUsart1lpen USART1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) GetUsart1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldUsart1lpenMask) != 0
}

// SetUsart1lpen USART1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) SetUsart1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldUsart1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldUsart1lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldUsart6lpenShift = 5
	RegisterApb2lpenrFieldUsart6lpenMask  = 0x20
)

// GetUsart6lpen USART6 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) GetUsart6lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldUsart6lpenMask) != 0
}

// SetUsart6lpen USART6 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) SetUsart6lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldUsart6lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldUsart6lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldSpi1lpenShift = 12
	RegisterApb2lpenrFieldSpi1lpenMask  = 0x1000
)

// GetSpi1lpen SPI1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) GetSpi1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldSpi1lpenMask) != 0
}

// SetSpi1lpen SPI1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) SetSpi1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldSpi1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldSpi1lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldSpi4lpenShift = 13
	RegisterApb2lpenrFieldSpi4lpenMask  = 0x2000
)

// GetSpi4lpen SPI4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) GetSpi4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldSpi4lpenMask) != 0
}

// SetSpi4lpen SPI4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) SetSpi4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldSpi4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldSpi4lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldTim15lpenShift = 16
	RegisterApb2lpenrFieldTim15lpenMask  = 0x10000
)

// GetTim15lpen TIM15 peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) GetTim15lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldTim15lpenMask) != 0
}

// SetTim15lpen TIM15 peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) SetTim15lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldTim15lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldTim15lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldTim16lpenShift = 17
	RegisterApb2lpenrFieldTim16lpenMask  = 0x20000
)

// GetTim16lpen TIM16 peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) GetTim16lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldTim16lpenMask) != 0
}

// SetTim16lpen TIM16 peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) SetTim16lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldTim16lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldTim16lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldTim17lpenShift = 18
	RegisterApb2lpenrFieldTim17lpenMask  = 0x40000
)

// GetTim17lpen TIM17 peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) GetTim17lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldTim17lpenMask) != 0
}

// SetTim17lpen TIM17 peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) SetTim17lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldTim17lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldTim17lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldSpi5lpenShift = 20
	RegisterApb2lpenrFieldSpi5lpenMask  = 0x100000
)

// GetSpi5lpen SPI5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) GetSpi5lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldSpi5lpenMask) != 0
}

// SetSpi5lpen SPI5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) SetSpi5lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldSpi5lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldSpi5lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldSai1lpenShift = 22
	RegisterApb2lpenrFieldSai1lpenMask  = 0x400000
)

// GetSai1lpen SAI1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) GetSai1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldSai1lpenMask) != 0
}

// SetSai1lpen SAI1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) SetSai1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldSai1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldSai1lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldSai2lpenShift = 23
	RegisterApb2lpenrFieldSai2lpenMask  = 0x800000
)

// GetSai2lpen SAI2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) GetSai2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldSai2lpenMask) != 0
}

// SetSai2lpen SAI2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) SetSai2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldSai2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldSai2lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldSai3lpenShift = 24
	RegisterApb2lpenrFieldSai3lpenMask  = 0x1000000
)

// GetSai3lpen SAI3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) GetSai3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldSai3lpenMask) != 0
}

// SetSai3lpen SAI3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) SetSai3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldSai3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldSai3lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldDfsdm1lpenShift = 28
	RegisterApb2lpenrFieldDfsdm1lpenMask  = 0x10000000
)

// GetDfsdm1lpen DFSDM1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) GetDfsdm1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldDfsdm1lpenMask) != 0
}

// SetDfsdm1lpen DFSDM1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb2lpenrType) SetDfsdm1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldDfsdm1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldDfsdm1lpenMask)
	}
}

const (
	RegisterApb2lpenrFieldHrtimlpenShift = 29
	RegisterApb2lpenrFieldHrtimlpenMask  = 0x20000000
)

// GetHrtimlpen HRTIM peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) GetHrtimlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb2lpenrFieldHrtimlpenMask) != 0
}

// SetHrtimlpen HRTIM peripheral clock enable during CSleep mode
func (r *RegisterApb2lpenrType) SetHrtimlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb2lpenrFieldHrtimlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb2lpenrFieldHrtimlpenMask)
	}
}

// RegisterApb4lpenrType RCC APB4 Sleep Clock Register
type RegisterApb4lpenrType uint32

func (r *RegisterApb4lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterApb4lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterApb4lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterApb4lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterApb4lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterApb4lpenrFieldSyscfglpenShift = 1
	RegisterApb4lpenrFieldSyscfglpenMask  = 0x2
)

// GetSyscfglpen SYSCFG peripheral clock enable during CSleep mode
func (r *RegisterApb4lpenrType) GetSyscfglpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldSyscfglpenMask) != 0
}

// SetSyscfglpen SYSCFG peripheral clock enable during CSleep mode
func (r *RegisterApb4lpenrType) SetSyscfglpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldSyscfglpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldSyscfglpenMask)
	}
}

const (
	RegisterApb4lpenrFieldLpuart1lpenShift = 3
	RegisterApb4lpenrFieldLpuart1lpenMask  = 0x8
)

// GetLpuart1lpen LPUART1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) GetLpuart1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldLpuart1lpenMask) != 0
}

// SetLpuart1lpen LPUART1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) SetLpuart1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldLpuart1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldLpuart1lpenMask)
	}
}

const (
	RegisterApb4lpenrFieldSpi6lpenShift = 5
	RegisterApb4lpenrFieldSpi6lpenMask  = 0x20
)

// GetSpi6lpen SPI6 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) GetSpi6lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldSpi6lpenMask) != 0
}

// SetSpi6lpen SPI6 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) SetSpi6lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldSpi6lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldSpi6lpenMask)
	}
}

const (
	RegisterApb4lpenrFieldI2c4lpenShift = 7
	RegisterApb4lpenrFieldI2c4lpenMask  = 0x80
)

// GetI2c4lpen I2C4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) GetI2c4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldI2c4lpenMask) != 0
}

// SetI2c4lpen I2C4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) SetI2c4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldI2c4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldI2c4lpenMask)
	}
}

const (
	RegisterApb4lpenrFieldLptim2lpenShift = 9
	RegisterApb4lpenrFieldLptim2lpenMask  = 0x200
)

// GetLptim2lpen LPTIM2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) GetLptim2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldLptim2lpenMask) != 0
}

// SetLptim2lpen LPTIM2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) SetLptim2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldLptim2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldLptim2lpenMask)
	}
}

const (
	RegisterApb4lpenrFieldLptim3lpenShift = 10
	RegisterApb4lpenrFieldLptim3lpenMask  = 0x400
)

// GetLptim3lpen LPTIM3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) GetLptim3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldLptim3lpenMask) != 0
}

// SetLptim3lpen LPTIM3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) SetLptim3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldLptim3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldLptim3lpenMask)
	}
}

const (
	RegisterApb4lpenrFieldLptim4lpenShift = 11
	RegisterApb4lpenrFieldLptim4lpenMask  = 0x800
)

// GetLptim4lpen LPTIM4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) GetLptim4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldLptim4lpenMask) != 0
}

// SetLptim4lpen LPTIM4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) SetLptim4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldLptim4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldLptim4lpenMask)
	}
}

const (
	RegisterApb4lpenrFieldLptim5lpenShift = 12
	RegisterApb4lpenrFieldLptim5lpenMask  = 0x1000
)

// GetLptim5lpen LPTIM5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) GetLptim5lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldLptim5lpenMask) != 0
}

// SetLptim5lpen LPTIM5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) SetLptim5lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldLptim5lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldLptim5lpenMask)
	}
}

const (
	RegisterApb4lpenrFieldComp12lpenShift = 14
	RegisterApb4lpenrFieldComp12lpenMask  = 0x4000
)

// GetComp12lpen COMP1/2 peripheral clock enable during CSleep mode
func (r *RegisterApb4lpenrType) GetComp12lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldComp12lpenMask) != 0
}

// SetComp12lpen COMP1/2 peripheral clock enable during CSleep mode
func (r *RegisterApb4lpenrType) SetComp12lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldComp12lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldComp12lpenMask)
	}
}

const (
	RegisterApb4lpenrFieldVreflpenShift = 15
	RegisterApb4lpenrFieldVreflpenMask  = 0x8000
)

// GetVreflpen VREF peripheral clock enable during CSleep mode
func (r *RegisterApb4lpenrType) GetVreflpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldVreflpenMask) != 0
}

// SetVreflpen VREF peripheral clock enable during CSleep mode
func (r *RegisterApb4lpenrType) SetVreflpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldVreflpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldVreflpenMask)
	}
}

const (
	RegisterApb4lpenrFieldRtcapblpenShift = 16
	RegisterApb4lpenrFieldRtcapblpenMask  = 0x10000
)

// GetRtcapblpen RTC APB Clock Enable During CSleep Mode
func (r *RegisterApb4lpenrType) GetRtcapblpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldRtcapblpenMask) != 0
}

// SetRtcapblpen RTC APB Clock Enable During CSleep Mode
func (r *RegisterApb4lpenrType) SetRtcapblpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldRtcapblpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldRtcapblpenMask)
	}
}

const (
	RegisterApb4lpenrFieldSai4lpenShift = 21
	RegisterApb4lpenrFieldSai4lpenMask  = 0x200000
)

// GetSai4lpen SAI4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) GetSai4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterApb4lpenrFieldSai4lpenMask) != 0
}

// SetSai4lpen SAI4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterApb4lpenrType) SetSai4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterApb4lpenrFieldSai4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterApb4lpenrFieldSai4lpenMask)
	}
}

// RegisterC1rsrType RCC Reset Status Register
type RegisterC1rsrType uint32

func (r *RegisterC1rsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1rsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1rsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1rsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1rsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1rsrFieldRmvfShift = 16
	RegisterC1rsrFieldRmvfMask  = 0x10000
)

// GetRmvf Remove reset flag
func (r *RegisterC1rsrType) GetRmvf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1rsrFieldRmvfMask) != 0
}

// SetRmvf Remove reset flag
func (r *RegisterC1rsrType) SetRmvf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1rsrFieldRmvfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1rsrFieldRmvfMask)
	}
}

const (
	RegisterC1rsrFieldCpurstfShift = 17
	RegisterC1rsrFieldCpurstfMask  = 0x20000
)

// GetCpurstf CPU reset flag
func (r *RegisterC1rsrType) GetCpurstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1rsrFieldCpurstfMask) != 0
}

// SetCpurstf CPU reset flag
func (r *RegisterC1rsrType) SetCpurstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1rsrFieldCpurstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1rsrFieldCpurstfMask)
	}
}

const (
	RegisterC1rsrFieldD1rstfShift = 19
	RegisterC1rsrFieldD1rstfMask  = 0x80000
)

// GetD1rstf D1 domain power switch reset flag
func (r *RegisterC1rsrType) GetD1rstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1rsrFieldD1rstfMask) != 0
}

// SetD1rstf D1 domain power switch reset flag
func (r *RegisterC1rsrType) SetD1rstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1rsrFieldD1rstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1rsrFieldD1rstfMask)
	}
}

const (
	RegisterC1rsrFieldD2rstfShift = 20
	RegisterC1rsrFieldD2rstfMask  = 0x100000
)

// GetD2rstf D2 domain power switch reset flag
func (r *RegisterC1rsrType) GetD2rstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1rsrFieldD2rstfMask) != 0
}

// SetD2rstf D2 domain power switch reset flag
func (r *RegisterC1rsrType) SetD2rstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1rsrFieldD2rstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1rsrFieldD2rstfMask)
	}
}

const (
	RegisterC1rsrFieldBorrstfShift = 21
	RegisterC1rsrFieldBorrstfMask  = 0x200000
)

// GetBorrstf BOR reset flag
func (r *RegisterC1rsrType) GetBorrstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1rsrFieldBorrstfMask) != 0
}

// SetBorrstf BOR reset flag
func (r *RegisterC1rsrType) SetBorrstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1rsrFieldBorrstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1rsrFieldBorrstfMask)
	}
}

const (
	RegisterC1rsrFieldPinrstfShift = 22
	RegisterC1rsrFieldPinrstfMask  = 0x400000
)

// GetPinrstf Pin reset flag (NRST)
func (r *RegisterC1rsrType) GetPinrstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1rsrFieldPinrstfMask) != 0
}

// SetPinrstf Pin reset flag (NRST)
func (r *RegisterC1rsrType) SetPinrstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1rsrFieldPinrstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1rsrFieldPinrstfMask)
	}
}

const (
	RegisterC1rsrFieldPorrstfShift = 23
	RegisterC1rsrFieldPorrstfMask  = 0x800000
)

// GetPorrstf POR/PDR reset flag
func (r *RegisterC1rsrType) GetPorrstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1rsrFieldPorrstfMask) != 0
}

// SetPorrstf POR/PDR reset flag
func (r *RegisterC1rsrType) SetPorrstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1rsrFieldPorrstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1rsrFieldPorrstfMask)
	}
}

const (
	RegisterC1rsrFieldSftrstfShift = 24
	RegisterC1rsrFieldSftrstfMask  = 0x1000000
)

// GetSftrstf System reset from CPU reset flag
func (r *RegisterC1rsrType) GetSftrstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1rsrFieldSftrstfMask) != 0
}

// SetSftrstf System reset from CPU reset flag
func (r *RegisterC1rsrType) SetSftrstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1rsrFieldSftrstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1rsrFieldSftrstfMask)
	}
}

const (
	RegisterC1rsrFieldIwdg1rstfShift = 26
	RegisterC1rsrFieldIwdg1rstfMask  = 0x4000000
)

// GetIwdg1rstf Independent Watchdog reset flag
func (r *RegisterC1rsrType) GetIwdg1rstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1rsrFieldIwdg1rstfMask) != 0
}

// SetIwdg1rstf Independent Watchdog reset flag
func (r *RegisterC1rsrType) SetIwdg1rstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1rsrFieldIwdg1rstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1rsrFieldIwdg1rstfMask)
	}
}

const (
	RegisterC1rsrFieldWwdg1rstfShift = 28
	RegisterC1rsrFieldWwdg1rstfMask  = 0x10000000
)

// GetWwdg1rstf Window Watchdog reset flag
func (r *RegisterC1rsrType) GetWwdg1rstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1rsrFieldWwdg1rstfMask) != 0
}

// SetWwdg1rstf Window Watchdog reset flag
func (r *RegisterC1rsrType) SetWwdg1rstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1rsrFieldWwdg1rstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1rsrFieldWwdg1rstfMask)
	}
}

const (
	RegisterC1rsrFieldLpwrrstfShift = 30
	RegisterC1rsrFieldLpwrrstfMask  = 0x40000000
)

// GetLpwrrstf Reset due to illegal D1 DStandby or CPU CStop flag
func (r *RegisterC1rsrType) GetLpwrrstf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1rsrFieldLpwrrstfMask) != 0
}

// SetLpwrrstf Reset due to illegal D1 DStandby or CPU CStop flag
func (r *RegisterC1rsrType) SetLpwrrstf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1rsrFieldLpwrrstfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1rsrFieldLpwrrstfMask)
	}
}

// RegisterC1ahb3enrType RCC AHB3 Clock Register
type RegisterC1ahb3enrType uint32

func (r *RegisterC1ahb3enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1ahb3enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1ahb3enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1ahb3enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1ahb3enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1ahb3enrFieldMdmaenShift = 0
	RegisterC1ahb3enrFieldMdmaenMask  = 0x1
)

// GetMdmaen MDMA Peripheral Clock Enable
func (r *RegisterC1ahb3enrType) GetMdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3enrFieldMdmaenMask) != 0
}

// SetMdmaen MDMA Peripheral Clock Enable
func (r *RegisterC1ahb3enrType) SetMdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3enrFieldMdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3enrFieldMdmaenMask)
	}
}

const (
	RegisterC1ahb3enrFieldDma2denShift = 4
	RegisterC1ahb3enrFieldDma2denMask  = 0x10
)

// GetDma2den DMA2D Peripheral Clock Enable
func (r *RegisterC1ahb3enrType) GetDma2den() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3enrFieldDma2denMask) != 0
}

// SetDma2den DMA2D Peripheral Clock Enable
func (r *RegisterC1ahb3enrType) SetDma2den(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3enrFieldDma2denMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3enrFieldDma2denMask)
	}
}

const (
	RegisterC1ahb3enrFieldJpgdecenShift = 5
	RegisterC1ahb3enrFieldJpgdecenMask  = 0x20
)

// GetJpgdecen JPGDEC Peripheral Clock Enable
func (r *RegisterC1ahb3enrType) GetJpgdecen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3enrFieldJpgdecenMask) != 0
}

// SetJpgdecen JPGDEC Peripheral Clock Enable
func (r *RegisterC1ahb3enrType) SetJpgdecen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3enrFieldJpgdecenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3enrFieldJpgdecenMask)
	}
}

const (
	RegisterC1ahb3enrFieldFmcenShift = 12
	RegisterC1ahb3enrFieldFmcenMask  = 0x1000
)

// GetFmcen FMC Peripheral Clocks Enable
func (r *RegisterC1ahb3enrType) GetFmcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3enrFieldFmcenMask) != 0
}

// SetFmcen FMC Peripheral Clocks Enable
func (r *RegisterC1ahb3enrType) SetFmcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3enrFieldFmcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3enrFieldFmcenMask)
	}
}

const (
	RegisterC1ahb3enrFieldQspienShift = 14
	RegisterC1ahb3enrFieldQspienMask  = 0x4000
)

// GetQspien QUADSPI and QUADSPI Delay Clock Enable
func (r *RegisterC1ahb3enrType) GetQspien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3enrFieldQspienMask) != 0
}

// SetQspien QUADSPI and QUADSPI Delay Clock Enable
func (r *RegisterC1ahb3enrType) SetQspien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3enrFieldQspienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3enrFieldQspienMask)
	}
}

const (
	RegisterC1ahb3enrFieldSdmmc1enShift = 16
	RegisterC1ahb3enrFieldSdmmc1enMask  = 0x10000
)

// GetSdmmc1en SDMMC1 and SDMMC1 Delay Clock Enable
func (r *RegisterC1ahb3enrType) GetSdmmc1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3enrFieldSdmmc1enMask) != 0
}

// SetSdmmc1en SDMMC1 and SDMMC1 Delay Clock Enable
func (r *RegisterC1ahb3enrType) SetSdmmc1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3enrFieldSdmmc1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3enrFieldSdmmc1enMask)
	}
}

// RegisterC1ahb1enrType RCC AHB1 Clock Register
type RegisterC1ahb1enrType uint32

func (r *RegisterC1ahb1enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1ahb1enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1ahb1enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1ahb1enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1ahb1enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1ahb1enrFieldDma1enShift = 0
	RegisterC1ahb1enrFieldDma1enMask  = 0x1
)

// GetDma1en DMA1 Clock Enable
func (r *RegisterC1ahb1enrType) GetDma1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1enrFieldDma1enMask) != 0
}

// SetDma1en DMA1 Clock Enable
func (r *RegisterC1ahb1enrType) SetDma1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1enrFieldDma1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1enrFieldDma1enMask)
	}
}

const (
	RegisterC1ahb1enrFieldDma2enShift = 1
	RegisterC1ahb1enrFieldDma2enMask  = 0x2
)

// GetDma2en DMA2 Clock Enable
func (r *RegisterC1ahb1enrType) GetDma2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1enrFieldDma2enMask) != 0
}

// SetDma2en DMA2 Clock Enable
func (r *RegisterC1ahb1enrType) SetDma2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1enrFieldDma2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1enrFieldDma2enMask)
	}
}

const (
	RegisterC1ahb1enrFieldAdc12enShift = 5
	RegisterC1ahb1enrFieldAdc12enMask  = 0x20
)

// GetAdc12en ADC1/2 Peripheral Clocks Enable
func (r *RegisterC1ahb1enrType) GetAdc12en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1enrFieldAdc12enMask) != 0
}

// SetAdc12en ADC1/2 Peripheral Clocks Enable
func (r *RegisterC1ahb1enrType) SetAdc12en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1enrFieldAdc12enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1enrFieldAdc12enMask)
	}
}

const (
	RegisterC1ahb1enrFieldEth1macenShift = 15
	RegisterC1ahb1enrFieldEth1macenMask  = 0x8000
)

// GetEth1macen Ethernet MAC bus interface Clock Enable
func (r *RegisterC1ahb1enrType) GetEth1macen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1enrFieldEth1macenMask) != 0
}

// SetEth1macen Ethernet MAC bus interface Clock Enable
func (r *RegisterC1ahb1enrType) SetEth1macen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1enrFieldEth1macenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1enrFieldEth1macenMask)
	}
}

const (
	RegisterC1ahb1enrFieldEth1txenShift = 16
	RegisterC1ahb1enrFieldEth1txenMask  = 0x10000
)

// GetEth1txen Ethernet Transmission Clock Enable
func (r *RegisterC1ahb1enrType) GetEth1txen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1enrFieldEth1txenMask) != 0
}

// SetEth1txen Ethernet Transmission Clock Enable
func (r *RegisterC1ahb1enrType) SetEth1txen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1enrFieldEth1txenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1enrFieldEth1txenMask)
	}
}

const (
	RegisterC1ahb1enrFieldEth1rxenShift = 17
	RegisterC1ahb1enrFieldEth1rxenMask  = 0x20000
)

// GetEth1rxen Ethernet Reception Clock Enable
func (r *RegisterC1ahb1enrType) GetEth1rxen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1enrFieldEth1rxenMask) != 0
}

// SetEth1rxen Ethernet Reception Clock Enable
func (r *RegisterC1ahb1enrType) SetEth1rxen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1enrFieldEth1rxenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1enrFieldEth1rxenMask)
	}
}

const (
	RegisterC1ahb1enrFieldUsb1otgenShift = 25
	RegisterC1ahb1enrFieldUsb1otgenMask  = 0x2000000
)

// GetUsb1otgen USB1OTG Peripheral Clocks Enable
func (r *RegisterC1ahb1enrType) GetUsb1otgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1enrFieldUsb1otgenMask) != 0
}

// SetUsb1otgen USB1OTG Peripheral Clocks Enable
func (r *RegisterC1ahb1enrType) SetUsb1otgen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1enrFieldUsb1otgenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1enrFieldUsb1otgenMask)
	}
}

const (
	RegisterC1ahb1enrFieldUsb1ulpienShift = 26
	RegisterC1ahb1enrFieldUsb1ulpienMask  = 0x4000000
)

// GetUsb1ulpien USB_PHY1 Clocks Enable
func (r *RegisterC1ahb1enrType) GetUsb1ulpien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1enrFieldUsb1ulpienMask) != 0
}

// SetUsb1ulpien USB_PHY1 Clocks Enable
func (r *RegisterC1ahb1enrType) SetUsb1ulpien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1enrFieldUsb1ulpienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1enrFieldUsb1ulpienMask)
	}
}

const (
	RegisterC1ahb1enrFieldUsb2otgenShift = 27
	RegisterC1ahb1enrFieldUsb2otgenMask  = 0x8000000
)

// GetUsb2otgen USB2OTG Peripheral Clocks Enable
func (r *RegisterC1ahb1enrType) GetUsb2otgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1enrFieldUsb2otgenMask) != 0
}

// SetUsb2otgen USB2OTG Peripheral Clocks Enable
func (r *RegisterC1ahb1enrType) SetUsb2otgen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1enrFieldUsb2otgenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1enrFieldUsb2otgenMask)
	}
}

const (
	RegisterC1ahb1enrFieldUsb2ulpienShift = 28
	RegisterC1ahb1enrFieldUsb2ulpienMask  = 0x10000000
)

// GetUsb2ulpien USB_PHY2 Clocks Enable
func (r *RegisterC1ahb1enrType) GetUsb2ulpien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1enrFieldUsb2ulpienMask) != 0
}

// SetUsb2ulpien USB_PHY2 Clocks Enable
func (r *RegisterC1ahb1enrType) SetUsb2ulpien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1enrFieldUsb2ulpienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1enrFieldUsb2ulpienMask)
	}
}

// RegisterC1ahb2enrType RCC AHB2 Clock Register
type RegisterC1ahb2enrType uint32

func (r *RegisterC1ahb2enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1ahb2enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1ahb2enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1ahb2enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1ahb2enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1ahb2enrFieldCamitfenShift = 0
	RegisterC1ahb2enrFieldCamitfenMask  = 0x1
)

// GetCamitfen CAMITF peripheral clock enable
func (r *RegisterC1ahb2enrType) GetCamitfen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2enrFieldCamitfenMask) != 0
}

// SetCamitfen CAMITF peripheral clock enable
func (r *RegisterC1ahb2enrType) SetCamitfen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2enrFieldCamitfenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2enrFieldCamitfenMask)
	}
}

const (
	RegisterC1ahb2enrFieldCryptenShift = 4
	RegisterC1ahb2enrFieldCryptenMask  = 0x10
)

// GetCrypten CRYPT peripheral clock enable
func (r *RegisterC1ahb2enrType) GetCrypten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2enrFieldCryptenMask) != 0
}

// SetCrypten CRYPT peripheral clock enable
func (r *RegisterC1ahb2enrType) SetCrypten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2enrFieldCryptenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2enrFieldCryptenMask)
	}
}

const (
	RegisterC1ahb2enrFieldHashenShift = 5
	RegisterC1ahb2enrFieldHashenMask  = 0x20
)

// GetHashen HASH peripheral clock enable
func (r *RegisterC1ahb2enrType) GetHashen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2enrFieldHashenMask) != 0
}

// SetHashen HASH peripheral clock enable
func (r *RegisterC1ahb2enrType) SetHashen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2enrFieldHashenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2enrFieldHashenMask)
	}
}

const (
	RegisterC1ahb2enrFieldRngenShift = 6
	RegisterC1ahb2enrFieldRngenMask  = 0x40
)

// GetRngen RNG peripheral clocks enable
func (r *RegisterC1ahb2enrType) GetRngen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2enrFieldRngenMask) != 0
}

// SetRngen RNG peripheral clocks enable
func (r *RegisterC1ahb2enrType) SetRngen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2enrFieldRngenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2enrFieldRngenMask)
	}
}

const (
	RegisterC1ahb2enrFieldSdmmc2enShift = 9
	RegisterC1ahb2enrFieldSdmmc2enMask  = 0x200
)

// GetSdmmc2en SDMMC2 and SDMMC2 delay clock enable
func (r *RegisterC1ahb2enrType) GetSdmmc2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2enrFieldSdmmc2enMask) != 0
}

// SetSdmmc2en SDMMC2 and SDMMC2 delay clock enable
func (r *RegisterC1ahb2enrType) SetSdmmc2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2enrFieldSdmmc2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2enrFieldSdmmc2enMask)
	}
}

const (
	RegisterC1ahb2enrFieldSram1enShift = 29
	RegisterC1ahb2enrFieldSram1enMask  = 0x20000000
)

// GetSram1en SRAM1 block enable
func (r *RegisterC1ahb2enrType) GetSram1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2enrFieldSram1enMask) != 0
}

// SetSram1en SRAM1 block enable
func (r *RegisterC1ahb2enrType) SetSram1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2enrFieldSram1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2enrFieldSram1enMask)
	}
}

const (
	RegisterC1ahb2enrFieldSram2enShift = 30
	RegisterC1ahb2enrFieldSram2enMask  = 0x40000000
)

// GetSram2en SRAM2 block enable
func (r *RegisterC1ahb2enrType) GetSram2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2enrFieldSram2enMask) != 0
}

// SetSram2en SRAM2 block enable
func (r *RegisterC1ahb2enrType) SetSram2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2enrFieldSram2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2enrFieldSram2enMask)
	}
}

const (
	RegisterC1ahb2enrFieldSram3enShift = 31
	RegisterC1ahb2enrFieldSram3enMask  = 0x80000000
)

// GetSram3en SRAM3 block enable
func (r *RegisterC1ahb2enrType) GetSram3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2enrFieldSram3enMask) != 0
}

// SetSram3en SRAM3 block enable
func (r *RegisterC1ahb2enrType) SetSram3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2enrFieldSram3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2enrFieldSram3enMask)
	}
}

// RegisterC1ahb4enrType RCC AHB4 Clock Register
type RegisterC1ahb4enrType uint32

func (r *RegisterC1ahb4enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1ahb4enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1ahb4enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1ahb4enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1ahb4enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1ahb4enrFieldGpioaenShift = 0
	RegisterC1ahb4enrFieldGpioaenMask  = 0x1
)

// GetGpioaen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) GetGpioaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldGpioaenMask) != 0
}

// SetGpioaen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) SetGpioaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldGpioaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldGpioaenMask)
	}
}

const (
	RegisterC1ahb4enrFieldGpiobenShift = 1
	RegisterC1ahb4enrFieldGpiobenMask  = 0x2
)

// GetGpioben 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) GetGpioben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldGpiobenMask) != 0
}

// SetGpioben 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) SetGpioben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldGpiobenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldGpiobenMask)
	}
}

const (
	RegisterC1ahb4enrFieldGpiocenShift = 2
	RegisterC1ahb4enrFieldGpiocenMask  = 0x4
)

// GetGpiocen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) GetGpiocen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldGpiocenMask) != 0
}

// SetGpiocen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) SetGpiocen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldGpiocenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldGpiocenMask)
	}
}

const (
	RegisterC1ahb4enrFieldGpiodenShift = 3
	RegisterC1ahb4enrFieldGpiodenMask  = 0x8
)

// GetGpioden 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) GetGpioden() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldGpiodenMask) != 0
}

// SetGpioden 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) SetGpioden(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldGpiodenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldGpiodenMask)
	}
}

const (
	RegisterC1ahb4enrFieldGpioeenShift = 4
	RegisterC1ahb4enrFieldGpioeenMask  = 0x10
)

// GetGpioeen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) GetGpioeen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldGpioeenMask) != 0
}

// SetGpioeen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) SetGpioeen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldGpioeenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldGpioeenMask)
	}
}

const (
	RegisterC1ahb4enrFieldGpiofenShift = 5
	RegisterC1ahb4enrFieldGpiofenMask  = 0x20
)

// GetGpiofen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) GetGpiofen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldGpiofenMask) != 0
}

// SetGpiofen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) SetGpiofen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldGpiofenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldGpiofenMask)
	}
}

const (
	RegisterC1ahb4enrFieldGpiogenShift = 6
	RegisterC1ahb4enrFieldGpiogenMask  = 0x40
)

// GetGpiogen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) GetGpiogen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldGpiogenMask) != 0
}

// SetGpiogen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) SetGpiogen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldGpiogenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldGpiogenMask)
	}
}

const (
	RegisterC1ahb4enrFieldGpiohenShift = 7
	RegisterC1ahb4enrFieldGpiohenMask  = 0x80
)

// GetGpiohen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) GetGpiohen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldGpiohenMask) != 0
}

// SetGpiohen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) SetGpiohen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldGpiohenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldGpiohenMask)
	}
}

const (
	RegisterC1ahb4enrFieldGpioienShift = 8
	RegisterC1ahb4enrFieldGpioienMask  = 0x100
)

// GetGpioien 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) GetGpioien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldGpioienMask) != 0
}

// SetGpioien 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) SetGpioien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldGpioienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldGpioienMask)
	}
}

const (
	RegisterC1ahb4enrFieldGpiojenShift = 9
	RegisterC1ahb4enrFieldGpiojenMask  = 0x200
)

// GetGpiojen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) GetGpiojen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldGpiojenMask) != 0
}

// SetGpiojen 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) SetGpiojen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldGpiojenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldGpiojenMask)
	}
}

const (
	RegisterC1ahb4enrFieldGpiokenShift = 10
	RegisterC1ahb4enrFieldGpiokenMask  = 0x400
)

// GetGpioken 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) GetGpioken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldGpiokenMask) != 0
}

// SetGpioken 0GPIO peripheral clock enable
func (r *RegisterC1ahb4enrType) SetGpioken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldGpiokenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldGpiokenMask)
	}
}

const (
	RegisterC1ahb4enrFieldCrcenShift = 19
	RegisterC1ahb4enrFieldCrcenMask  = 0x80000
)

// GetCrcen CRC peripheral clock enable
func (r *RegisterC1ahb4enrType) GetCrcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldCrcenMask) != 0
}

// SetCrcen CRC peripheral clock enable
func (r *RegisterC1ahb4enrType) SetCrcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldCrcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldCrcenMask)
	}
}

const (
	RegisterC1ahb4enrFieldBdmaenShift = 21
	RegisterC1ahb4enrFieldBdmaenMask  = 0x200000
)

// GetBdmaen BDMA and DMAMUX2 Clock Enable
func (r *RegisterC1ahb4enrType) GetBdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldBdmaenMask) != 0
}

// SetBdmaen BDMA and DMAMUX2 Clock Enable
func (r *RegisterC1ahb4enrType) SetBdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldBdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldBdmaenMask)
	}
}

const (
	RegisterC1ahb4enrFieldAdc3enShift = 24
	RegisterC1ahb4enrFieldAdc3enMask  = 0x1000000
)

// GetAdc3en ADC3 Peripheral Clocks Enable
func (r *RegisterC1ahb4enrType) GetAdc3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldAdc3enMask) != 0
}

// SetAdc3en ADC3 Peripheral Clocks Enable
func (r *RegisterC1ahb4enrType) SetAdc3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldAdc3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldAdc3enMask)
	}
}

const (
	RegisterC1ahb4enrFieldHsemenShift = 25
	RegisterC1ahb4enrFieldHsemenMask  = 0x2000000
)

// GetHsemen HSEM peripheral clock enable
func (r *RegisterC1ahb4enrType) GetHsemen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldHsemenMask) != 0
}

// SetHsemen HSEM peripheral clock enable
func (r *RegisterC1ahb4enrType) SetHsemen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldHsemenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldHsemenMask)
	}
}

const (
	RegisterC1ahb4enrFieldBkpramenShift = 28
	RegisterC1ahb4enrFieldBkpramenMask  = 0x10000000
)

// GetBkpramen Backup RAM Clock Enable
func (r *RegisterC1ahb4enrType) GetBkpramen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4enrFieldBkpramenMask) != 0
}

// SetBkpramen Backup RAM Clock Enable
func (r *RegisterC1ahb4enrType) SetBkpramen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4enrFieldBkpramenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4enrFieldBkpramenMask)
	}
}

// RegisterC1apb3enrType RCC APB3 Clock Register
type RegisterC1apb3enrType uint32

func (r *RegisterC1apb3enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1apb3enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1apb3enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1apb3enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1apb3enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1apb3enrFieldLtdcenShift = 3
	RegisterC1apb3enrFieldLtdcenMask  = 0x8
)

// GetLtdcen LTDC peripheral clock enable
func (r *RegisterC1apb3enrType) GetLtdcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb3enrFieldLtdcenMask) != 0
}

// SetLtdcen LTDC peripheral clock enable
func (r *RegisterC1apb3enrType) SetLtdcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb3enrFieldLtdcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb3enrFieldLtdcenMask)
	}
}

const (
	RegisterC1apb3enrFieldWwdg1enShift = 6
	RegisterC1apb3enrFieldWwdg1enMask  = 0x40
)

// GetWwdg1en WWDG1 Clock Enable
func (r *RegisterC1apb3enrType) GetWwdg1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb3enrFieldWwdg1enMask) != 0
}

// SetWwdg1en WWDG1 Clock Enable
func (r *RegisterC1apb3enrType) SetWwdg1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb3enrFieldWwdg1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb3enrFieldWwdg1enMask)
	}
}

// RegisterC1apb1lenrType RCC APB1 Clock Register
type RegisterC1apb1lenrType uint32

func (r *RegisterC1apb1lenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1apb1lenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1apb1lenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1apb1lenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1apb1lenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1apb1lenrFieldTim2enShift = 0
	RegisterC1apb1lenrFieldTim2enMask  = 0x1
)

// GetTim2en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) GetTim2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldTim2enMask) != 0
}

// SetTim2en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) SetTim2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldTim2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldTim2enMask)
	}
}

const (
	RegisterC1apb1lenrFieldTim3enShift = 1
	RegisterC1apb1lenrFieldTim3enMask  = 0x2
)

// GetTim3en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) GetTim3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldTim3enMask) != 0
}

// SetTim3en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) SetTim3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldTim3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldTim3enMask)
	}
}

const (
	RegisterC1apb1lenrFieldTim4enShift = 2
	RegisterC1apb1lenrFieldTim4enMask  = 0x4
)

// GetTim4en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) GetTim4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldTim4enMask) != 0
}

// SetTim4en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) SetTim4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldTim4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldTim4enMask)
	}
}

const (
	RegisterC1apb1lenrFieldTim5enShift = 3
	RegisterC1apb1lenrFieldTim5enMask  = 0x8
)

// GetTim5en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) GetTim5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldTim5enMask) != 0
}

// SetTim5en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) SetTim5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldTim5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldTim5enMask)
	}
}

const (
	RegisterC1apb1lenrFieldTim6enShift = 4
	RegisterC1apb1lenrFieldTim6enMask  = 0x10
)

// GetTim6en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) GetTim6en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldTim6enMask) != 0
}

// SetTim6en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) SetTim6en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldTim6enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldTim6enMask)
	}
}

const (
	RegisterC1apb1lenrFieldTim7enShift = 5
	RegisterC1apb1lenrFieldTim7enMask  = 0x20
)

// GetTim7en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) GetTim7en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldTim7enMask) != 0
}

// SetTim7en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) SetTim7en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldTim7enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldTim7enMask)
	}
}

const (
	RegisterC1apb1lenrFieldTim12enShift = 6
	RegisterC1apb1lenrFieldTim12enMask  = 0x40
)

// GetTim12en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) GetTim12en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldTim12enMask) != 0
}

// SetTim12en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) SetTim12en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldTim12enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldTim12enMask)
	}
}

const (
	RegisterC1apb1lenrFieldTim13enShift = 7
	RegisterC1apb1lenrFieldTim13enMask  = 0x80
)

// GetTim13en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) GetTim13en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldTim13enMask) != 0
}

// SetTim13en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) SetTim13en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldTim13enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldTim13enMask)
	}
}

const (
	RegisterC1apb1lenrFieldTim14enShift = 8
	RegisterC1apb1lenrFieldTim14enMask  = 0x100
)

// GetTim14en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) GetTim14en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldTim14enMask) != 0
}

// SetTim14en TIM peripheral clock enable
func (r *RegisterC1apb1lenrType) SetTim14en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldTim14enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldTim14enMask)
	}
}

const (
	RegisterC1apb1lenrFieldLptim1enShift = 9
	RegisterC1apb1lenrFieldLptim1enMask  = 0x200
)

// GetLptim1en LPTIM1 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetLptim1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldLptim1enMask) != 0
}

// SetLptim1en LPTIM1 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetLptim1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldLptim1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldLptim1enMask)
	}
}

const (
	RegisterC1apb1lenrFieldSpi2enShift = 14
	RegisterC1apb1lenrFieldSpi2enMask  = 0x4000
)

// GetSpi2en SPI2 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetSpi2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldSpi2enMask) != 0
}

// SetSpi2en SPI2 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetSpi2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldSpi2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldSpi2enMask)
	}
}

const (
	RegisterC1apb1lenrFieldSpi3enShift = 15
	RegisterC1apb1lenrFieldSpi3enMask  = 0x8000
)

// GetSpi3en SPI3 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetSpi3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldSpi3enMask) != 0
}

// SetSpi3en SPI3 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetSpi3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldSpi3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldSpi3enMask)
	}
}

const (
	RegisterC1apb1lenrFieldSpdifrxenShift = 16
	RegisterC1apb1lenrFieldSpdifrxenMask  = 0x10000
)

// GetSpdifrxen SPDIFRX Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetSpdifrxen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldSpdifrxenMask) != 0
}

// SetSpdifrxen SPDIFRX Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetSpdifrxen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldSpdifrxenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldSpdifrxenMask)
	}
}

const (
	RegisterC1apb1lenrFieldUsart2enShift = 17
	RegisterC1apb1lenrFieldUsart2enMask  = 0x20000
)

// GetUsart2en USART2 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetUsart2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldUsart2enMask) != 0
}

// SetUsart2en USART2 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetUsart2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldUsart2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldUsart2enMask)
	}
}

const (
	RegisterC1apb1lenrFieldUsart3enShift = 18
	RegisterC1apb1lenrFieldUsart3enMask  = 0x40000
)

// GetUsart3en USART3 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetUsart3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldUsart3enMask) != 0
}

// SetUsart3en USART3 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetUsart3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldUsart3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldUsart3enMask)
	}
}

const (
	RegisterC1apb1lenrFieldUart4enShift = 19
	RegisterC1apb1lenrFieldUart4enMask  = 0x80000
)

// GetUart4en UART4 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetUart4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldUart4enMask) != 0
}

// SetUart4en UART4 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetUart4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldUart4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldUart4enMask)
	}
}

const (
	RegisterC1apb1lenrFieldUart5enShift = 20
	RegisterC1apb1lenrFieldUart5enMask  = 0x100000
)

// GetUart5en UART5 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetUart5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldUart5enMask) != 0
}

// SetUart5en UART5 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetUart5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldUart5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldUart5enMask)
	}
}

const (
	RegisterC1apb1lenrFieldI2c1enShift = 21
	RegisterC1apb1lenrFieldI2c1enMask  = 0x200000
)

// GetI2c1en I2C1 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetI2c1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldI2c1enMask) != 0
}

// SetI2c1en I2C1 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetI2c1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldI2c1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldI2c1enMask)
	}
}

const (
	RegisterC1apb1lenrFieldI2c2enShift = 22
	RegisterC1apb1lenrFieldI2c2enMask  = 0x400000
)

// GetI2c2en I2C2 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetI2c2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldI2c2enMask) != 0
}

// SetI2c2en I2C2 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetI2c2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldI2c2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldI2c2enMask)
	}
}

const (
	RegisterC1apb1lenrFieldI2c3enShift = 23
	RegisterC1apb1lenrFieldI2c3enMask  = 0x800000
)

// GetI2c3en I2C3 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetI2c3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldI2c3enMask) != 0
}

// SetI2c3en I2C3 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetI2c3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldI2c3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldI2c3enMask)
	}
}

const (
	RegisterC1apb1lenrFieldHdmicecenShift = 27
	RegisterC1apb1lenrFieldHdmicecenMask  = 0x8000000
)

// GetHdmicecen HDMI-CEC peripheral clock enable
func (r *RegisterC1apb1lenrType) GetHdmicecen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldHdmicecenMask) != 0
}

// SetHdmicecen HDMI-CEC peripheral clock enable
func (r *RegisterC1apb1lenrType) SetHdmicecen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldHdmicecenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldHdmicecenMask)
	}
}

const (
	RegisterC1apb1lenrFieldDac12enShift = 29
	RegisterC1apb1lenrFieldDac12enMask  = 0x20000000
)

// GetDac12en DAC1&2 peripheral clock enable
func (r *RegisterC1apb1lenrType) GetDac12en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldDac12enMask) != 0
}

// SetDac12en DAC1&2 peripheral clock enable
func (r *RegisterC1apb1lenrType) SetDac12en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldDac12enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldDac12enMask)
	}
}

const (
	RegisterC1apb1lenrFieldUsart7enShift = 30
	RegisterC1apb1lenrFieldUsart7enMask  = 0x40000000
)

// GetUsart7en USART7 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetUsart7en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldUsart7enMask) != 0
}

// SetUsart7en USART7 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetUsart7en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldUsart7enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldUsart7enMask)
	}
}

const (
	RegisterC1apb1lenrFieldUsart8enShift = 31
	RegisterC1apb1lenrFieldUsart8enMask  = 0x80000000
)

// GetUsart8en USART8 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) GetUsart8en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1lenrFieldUsart8enMask) != 0
}

// SetUsart8en USART8 Peripheral Clocks Enable
func (r *RegisterC1apb1lenrType) SetUsart8en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1lenrFieldUsart8enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1lenrFieldUsart8enMask)
	}
}

// RegisterC1apb1henrType RCC APB1 Clock Register
type RegisterC1apb1henrType uint32

func (r *RegisterC1apb1henrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1apb1henrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1apb1henrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1apb1henrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1apb1henrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1apb1henrFieldCrsenShift = 1
	RegisterC1apb1henrFieldCrsenMask  = 0x2
)

// GetCrsen Clock Recovery System peripheral clock enable
func (r *RegisterC1apb1henrType) GetCrsen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1henrFieldCrsenMask) != 0
}

// SetCrsen Clock Recovery System peripheral clock enable
func (r *RegisterC1apb1henrType) SetCrsen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1henrFieldCrsenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1henrFieldCrsenMask)
	}
}

const (
	RegisterC1apb1henrFieldSwpenShift = 2
	RegisterC1apb1henrFieldSwpenMask  = 0x4
)

// GetSwpen SWPMI Peripheral Clocks Enable
func (r *RegisterC1apb1henrType) GetSwpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1henrFieldSwpenMask) != 0
}

// SetSwpen SWPMI Peripheral Clocks Enable
func (r *RegisterC1apb1henrType) SetSwpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1henrFieldSwpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1henrFieldSwpenMask)
	}
}

const (
	RegisterC1apb1henrFieldOpampenShift = 4
	RegisterC1apb1henrFieldOpampenMask  = 0x10
)

// GetOpampen OPAMP peripheral clock enable
func (r *RegisterC1apb1henrType) GetOpampen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1henrFieldOpampenMask) != 0
}

// SetOpampen OPAMP peripheral clock enable
func (r *RegisterC1apb1henrType) SetOpampen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1henrFieldOpampenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1henrFieldOpampenMask)
	}
}

const (
	RegisterC1apb1henrFieldMdiosenShift = 5
	RegisterC1apb1henrFieldMdiosenMask  = 0x20
)

// GetMdiosen MDIOS peripheral clock enable
func (r *RegisterC1apb1henrType) GetMdiosen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1henrFieldMdiosenMask) != 0
}

// SetMdiosen MDIOS peripheral clock enable
func (r *RegisterC1apb1henrType) SetMdiosen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1henrFieldMdiosenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1henrFieldMdiosenMask)
	}
}

const (
	RegisterC1apb1henrFieldFdcanenShift = 8
	RegisterC1apb1henrFieldFdcanenMask  = 0x100
)

// GetFdcanen FDCAN Peripheral Clocks Enable
func (r *RegisterC1apb1henrType) GetFdcanen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1henrFieldFdcanenMask) != 0
}

// SetFdcanen FDCAN Peripheral Clocks Enable
func (r *RegisterC1apb1henrType) SetFdcanen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1henrFieldFdcanenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1henrFieldFdcanenMask)
	}
}

// RegisterC1apb2enrType RCC APB2 Clock Register
type RegisterC1apb2enrType uint32

func (r *RegisterC1apb2enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1apb2enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1apb2enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1apb2enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1apb2enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1apb2enrFieldTim1enShift = 0
	RegisterC1apb2enrFieldTim1enMask  = 0x1
)

// GetTim1en TIM1 peripheral clock enable
func (r *RegisterC1apb2enrType) GetTim1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldTim1enMask) != 0
}

// SetTim1en TIM1 peripheral clock enable
func (r *RegisterC1apb2enrType) SetTim1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldTim1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldTim1enMask)
	}
}

const (
	RegisterC1apb2enrFieldTim8enShift = 1
	RegisterC1apb2enrFieldTim8enMask  = 0x2
)

// GetTim8en TIM8 peripheral clock enable
func (r *RegisterC1apb2enrType) GetTim8en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldTim8enMask) != 0
}

// SetTim8en TIM8 peripheral clock enable
func (r *RegisterC1apb2enrType) SetTim8en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldTim8enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldTim8enMask)
	}
}

const (
	RegisterC1apb2enrFieldUsart1enShift = 4
	RegisterC1apb2enrFieldUsart1enMask  = 0x10
)

// GetUsart1en USART1 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) GetUsart1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldUsart1enMask) != 0
}

// SetUsart1en USART1 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) SetUsart1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldUsart1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldUsart1enMask)
	}
}

const (
	RegisterC1apb2enrFieldUsart6enShift = 5
	RegisterC1apb2enrFieldUsart6enMask  = 0x20
)

// GetUsart6en USART6 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) GetUsart6en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldUsart6enMask) != 0
}

// SetUsart6en USART6 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) SetUsart6en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldUsart6enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldUsart6enMask)
	}
}

const (
	RegisterC1apb2enrFieldSpi1enShift = 12
	RegisterC1apb2enrFieldSpi1enMask  = 0x1000
)

// GetSpi1en SPI1 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) GetSpi1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldSpi1enMask) != 0
}

// SetSpi1en SPI1 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) SetSpi1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldSpi1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldSpi1enMask)
	}
}

const (
	RegisterC1apb2enrFieldSpi4enShift = 13
	RegisterC1apb2enrFieldSpi4enMask  = 0x2000
)

// GetSpi4en SPI4 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) GetSpi4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldSpi4enMask) != 0
}

// SetSpi4en SPI4 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) SetSpi4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldSpi4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldSpi4enMask)
	}
}

const (
	RegisterC1apb2enrFieldTim15enShift = 16
	RegisterC1apb2enrFieldTim15enMask  = 0x10000
)

// GetTim15en TIM15 peripheral clock enable
func (r *RegisterC1apb2enrType) GetTim15en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldTim15enMask) != 0
}

// SetTim15en TIM15 peripheral clock enable
func (r *RegisterC1apb2enrType) SetTim15en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldTim15enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldTim15enMask)
	}
}

const (
	RegisterC1apb2enrFieldTim16enShift = 17
	RegisterC1apb2enrFieldTim16enMask  = 0x20000
)

// GetTim16en TIM16 peripheral clock enable
func (r *RegisterC1apb2enrType) GetTim16en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldTim16enMask) != 0
}

// SetTim16en TIM16 peripheral clock enable
func (r *RegisterC1apb2enrType) SetTim16en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldTim16enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldTim16enMask)
	}
}

const (
	RegisterC1apb2enrFieldTim17enShift = 18
	RegisterC1apb2enrFieldTim17enMask  = 0x40000
)

// GetTim17en TIM17 peripheral clock enable
func (r *RegisterC1apb2enrType) GetTim17en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldTim17enMask) != 0
}

// SetTim17en TIM17 peripheral clock enable
func (r *RegisterC1apb2enrType) SetTim17en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldTim17enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldTim17enMask)
	}
}

const (
	RegisterC1apb2enrFieldSpi5enShift = 20
	RegisterC1apb2enrFieldSpi5enMask  = 0x100000
)

// GetSpi5en SPI5 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) GetSpi5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldSpi5enMask) != 0
}

// SetSpi5en SPI5 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) SetSpi5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldSpi5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldSpi5enMask)
	}
}

const (
	RegisterC1apb2enrFieldSai1enShift = 22
	RegisterC1apb2enrFieldSai1enMask  = 0x400000
)

// GetSai1en SAI1 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) GetSai1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldSai1enMask) != 0
}

// SetSai1en SAI1 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) SetSai1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldSai1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldSai1enMask)
	}
}

const (
	RegisterC1apb2enrFieldSai2enShift = 23
	RegisterC1apb2enrFieldSai2enMask  = 0x800000
)

// GetSai2en SAI2 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) GetSai2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldSai2enMask) != 0
}

// SetSai2en SAI2 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) SetSai2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldSai2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldSai2enMask)
	}
}

const (
	RegisterC1apb2enrFieldSai3enShift = 24
	RegisterC1apb2enrFieldSai3enMask  = 0x1000000
)

// GetSai3en SAI3 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) GetSai3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldSai3enMask) != 0
}

// SetSai3en SAI3 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) SetSai3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldSai3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldSai3enMask)
	}
}

const (
	RegisterC1apb2enrFieldDfsdm1enShift = 28
	RegisterC1apb2enrFieldDfsdm1enMask  = 0x10000000
)

// GetDfsdm1en DFSDM1 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) GetDfsdm1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldDfsdm1enMask) != 0
}

// SetDfsdm1en DFSDM1 Peripheral Clocks Enable
func (r *RegisterC1apb2enrType) SetDfsdm1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldDfsdm1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldDfsdm1enMask)
	}
}

const (
	RegisterC1apb2enrFieldHrtimenShift = 29
	RegisterC1apb2enrFieldHrtimenMask  = 0x20000000
)

// GetHrtimen HRTIM peripheral clock enable
func (r *RegisterC1apb2enrType) GetHrtimen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2enrFieldHrtimenMask) != 0
}

// SetHrtimen HRTIM peripheral clock enable
func (r *RegisterC1apb2enrType) SetHrtimen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2enrFieldHrtimenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2enrFieldHrtimenMask)
	}
}

// RegisterC1apb4enrType RCC APB4 Clock Register
type RegisterC1apb4enrType uint32

func (r *RegisterC1apb4enrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1apb4enrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1apb4enrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1apb4enrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1apb4enrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1apb4enrFieldSyscfgenShift = 1
	RegisterC1apb4enrFieldSyscfgenMask  = 0x2
)

// GetSyscfgen SYSCFG peripheral clock enable
func (r *RegisterC1apb4enrType) GetSyscfgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldSyscfgenMask) != 0
}

// SetSyscfgen SYSCFG peripheral clock enable
func (r *RegisterC1apb4enrType) SetSyscfgen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldSyscfgenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldSyscfgenMask)
	}
}

const (
	RegisterC1apb4enrFieldLpuart1enShift = 3
	RegisterC1apb4enrFieldLpuart1enMask  = 0x8
)

// GetLpuart1en LPUART1 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) GetLpuart1en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldLpuart1enMask) != 0
}

// SetLpuart1en LPUART1 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) SetLpuart1en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldLpuart1enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldLpuart1enMask)
	}
}

const (
	RegisterC1apb4enrFieldSpi6enShift = 5
	RegisterC1apb4enrFieldSpi6enMask  = 0x20
)

// GetSpi6en SPI6 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) GetSpi6en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldSpi6enMask) != 0
}

// SetSpi6en SPI6 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) SetSpi6en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldSpi6enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldSpi6enMask)
	}
}

const (
	RegisterC1apb4enrFieldI2c4enShift = 7
	RegisterC1apb4enrFieldI2c4enMask  = 0x80
)

// GetI2c4en I2C4 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) GetI2c4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldI2c4enMask) != 0
}

// SetI2c4en I2C4 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) SetI2c4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldI2c4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldI2c4enMask)
	}
}

const (
	RegisterC1apb4enrFieldLptim2enShift = 9
	RegisterC1apb4enrFieldLptim2enMask  = 0x200
)

// GetLptim2en LPTIM2 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) GetLptim2en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldLptim2enMask) != 0
}

// SetLptim2en LPTIM2 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) SetLptim2en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldLptim2enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldLptim2enMask)
	}
}

const (
	RegisterC1apb4enrFieldLptim3enShift = 10
	RegisterC1apb4enrFieldLptim3enMask  = 0x400
)

// GetLptim3en LPTIM3 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) GetLptim3en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldLptim3enMask) != 0
}

// SetLptim3en LPTIM3 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) SetLptim3en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldLptim3enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldLptim3enMask)
	}
}

const (
	RegisterC1apb4enrFieldLptim4enShift = 11
	RegisterC1apb4enrFieldLptim4enMask  = 0x800
)

// GetLptim4en LPTIM4 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) GetLptim4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldLptim4enMask) != 0
}

// SetLptim4en LPTIM4 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) SetLptim4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldLptim4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldLptim4enMask)
	}
}

const (
	RegisterC1apb4enrFieldLptim5enShift = 12
	RegisterC1apb4enrFieldLptim5enMask  = 0x1000
)

// GetLptim5en LPTIM5 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) GetLptim5en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldLptim5enMask) != 0
}

// SetLptim5en LPTIM5 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) SetLptim5en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldLptim5enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldLptim5enMask)
	}
}

const (
	RegisterC1apb4enrFieldComp12enShift = 14
	RegisterC1apb4enrFieldComp12enMask  = 0x4000
)

// GetComp12en COMP1/2 peripheral clock enable
func (r *RegisterC1apb4enrType) GetComp12en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldComp12enMask) != 0
}

// SetComp12en COMP1/2 peripheral clock enable
func (r *RegisterC1apb4enrType) SetComp12en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldComp12enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldComp12enMask)
	}
}

const (
	RegisterC1apb4enrFieldVrefenShift = 15
	RegisterC1apb4enrFieldVrefenMask  = 0x8000
)

// GetVrefen VREF peripheral clock enable
func (r *RegisterC1apb4enrType) GetVrefen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldVrefenMask) != 0
}

// SetVrefen VREF peripheral clock enable
func (r *RegisterC1apb4enrType) SetVrefen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldVrefenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldVrefenMask)
	}
}

const (
	RegisterC1apb4enrFieldRtcapbenShift = 16
	RegisterC1apb4enrFieldRtcapbenMask  = 0x10000
)

// GetRtcapben RTC APB Clock Enable
func (r *RegisterC1apb4enrType) GetRtcapben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldRtcapbenMask) != 0
}

// SetRtcapben RTC APB Clock Enable
func (r *RegisterC1apb4enrType) SetRtcapben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldRtcapbenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldRtcapbenMask)
	}
}

const (
	RegisterC1apb4enrFieldSai4enShift = 21
	RegisterC1apb4enrFieldSai4enMask  = 0x200000
)

// GetSai4en SAI4 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) GetSai4en() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4enrFieldSai4enMask) != 0
}

// SetSai4en SAI4 Peripheral Clocks Enable
func (r *RegisterC1apb4enrType) SetSai4en(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4enrFieldSai4enMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4enrFieldSai4enMask)
	}
}

// RegisterC1ahb3lpenrType RCC AHB3 Sleep Clock Register
type RegisterC1ahb3lpenrType uint32

func (r *RegisterC1ahb3lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1ahb3lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1ahb3lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1ahb3lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1ahb3lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1ahb3lpenrFieldMdmalpenShift = 0
	RegisterC1ahb3lpenrFieldMdmalpenMask  = 0x1
)

// GetMdmalpen MDMA Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) GetMdmalpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3lpenrFieldMdmalpenMask) != 0
}

// SetMdmalpen MDMA Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) SetMdmalpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3lpenrFieldMdmalpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3lpenrFieldMdmalpenMask)
	}
}

const (
	RegisterC1ahb3lpenrFieldDma2dlpenShift = 4
	RegisterC1ahb3lpenrFieldDma2dlpenMask  = 0x10
)

// GetDma2dlpen DMA2D Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) GetDma2dlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3lpenrFieldDma2dlpenMask) != 0
}

// SetDma2dlpen DMA2D Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) SetDma2dlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3lpenrFieldDma2dlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3lpenrFieldDma2dlpenMask)
	}
}

const (
	RegisterC1ahb3lpenrFieldJpgdeclpenShift = 5
	RegisterC1ahb3lpenrFieldJpgdeclpenMask  = 0x20
)

// GetJpgdeclpen JPGDEC Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) GetJpgdeclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3lpenrFieldJpgdeclpenMask) != 0
}

// SetJpgdeclpen JPGDEC Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) SetJpgdeclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3lpenrFieldJpgdeclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3lpenrFieldJpgdeclpenMask)
	}
}

const (
	RegisterC1ahb3lpenrFieldFlitflpenShift = 8
	RegisterC1ahb3lpenrFieldFlitflpenMask  = 0x100
)

// GetFlitflpen FLITF Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) GetFlitflpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3lpenrFieldFlitflpenMask) != 0
}

// SetFlitflpen FLITF Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) SetFlitflpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3lpenrFieldFlitflpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3lpenrFieldFlitflpenMask)
	}
}

const (
	RegisterC1ahb3lpenrFieldFmclpenShift = 12
	RegisterC1ahb3lpenrFieldFmclpenMask  = 0x1000
)

// GetFmclpen FMC Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) GetFmclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3lpenrFieldFmclpenMask) != 0
}

// SetFmclpen FMC Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) SetFmclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3lpenrFieldFmclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3lpenrFieldFmclpenMask)
	}
}

const (
	RegisterC1ahb3lpenrFieldQspilpenShift = 14
	RegisterC1ahb3lpenrFieldQspilpenMask  = 0x4000
)

// GetQspilpen QUADSPI and QUADSPI Delay Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) GetQspilpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3lpenrFieldQspilpenMask) != 0
}

// SetQspilpen QUADSPI and QUADSPI Delay Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) SetQspilpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3lpenrFieldQspilpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3lpenrFieldQspilpenMask)
	}
}

const (
	RegisterC1ahb3lpenrFieldSdmmc1lpenShift = 16
	RegisterC1ahb3lpenrFieldSdmmc1lpenMask  = 0x10000
)

// GetSdmmc1lpen SDMMC1 and SDMMC1 Delay Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) GetSdmmc1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3lpenrFieldSdmmc1lpenMask) != 0
}

// SetSdmmc1lpen SDMMC1 and SDMMC1 Delay Clock Enable During CSleep Mode
func (r *RegisterC1ahb3lpenrType) SetSdmmc1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3lpenrFieldSdmmc1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3lpenrFieldSdmmc1lpenMask)
	}
}

const (
	RegisterC1ahb3lpenrFieldD1dtcm1lpenShift = 28
	RegisterC1ahb3lpenrFieldD1dtcm1lpenMask  = 0x10000000
)

// GetD1dtcm1lpen D1DTCM1 Block Clock Enable During CSleep mode
func (r *RegisterC1ahb3lpenrType) GetD1dtcm1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3lpenrFieldD1dtcm1lpenMask) != 0
}

// SetD1dtcm1lpen D1DTCM1 Block Clock Enable During CSleep mode
func (r *RegisterC1ahb3lpenrType) SetD1dtcm1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3lpenrFieldD1dtcm1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3lpenrFieldD1dtcm1lpenMask)
	}
}

const (
	RegisterC1ahb3lpenrFieldDtcm2lpenShift = 29
	RegisterC1ahb3lpenrFieldDtcm2lpenMask  = 0x20000000
)

// GetDtcm2lpen D1 DTCM2 Block Clock Enable During CSleep mode
func (r *RegisterC1ahb3lpenrType) GetDtcm2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3lpenrFieldDtcm2lpenMask) != 0
}

// SetDtcm2lpen D1 DTCM2 Block Clock Enable During CSleep mode
func (r *RegisterC1ahb3lpenrType) SetDtcm2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3lpenrFieldDtcm2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3lpenrFieldDtcm2lpenMask)
	}
}

const (
	RegisterC1ahb3lpenrFieldItcmlpenShift = 30
	RegisterC1ahb3lpenrFieldItcmlpenMask  = 0x40000000
)

// GetItcmlpen D1ITCM Block Clock Enable During CSleep mode
func (r *RegisterC1ahb3lpenrType) GetItcmlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3lpenrFieldItcmlpenMask) != 0
}

// SetItcmlpen D1ITCM Block Clock Enable During CSleep mode
func (r *RegisterC1ahb3lpenrType) SetItcmlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3lpenrFieldItcmlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3lpenrFieldItcmlpenMask)
	}
}

const (
	RegisterC1ahb3lpenrFieldAxisramlpenShift = 31
	RegisterC1ahb3lpenrFieldAxisramlpenMask  = 0x80000000
)

// GetAxisramlpen AXISRAM Block Clock Enable During CSleep mode
func (r *RegisterC1ahb3lpenrType) GetAxisramlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb3lpenrFieldAxisramlpenMask) != 0
}

// SetAxisramlpen AXISRAM Block Clock Enable During CSleep mode
func (r *RegisterC1ahb3lpenrType) SetAxisramlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb3lpenrFieldAxisramlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb3lpenrFieldAxisramlpenMask)
	}
}

// RegisterC1ahb1lpenrType RCC AHB1 Sleep Clock Register
type RegisterC1ahb1lpenrType uint32

func (r *RegisterC1ahb1lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1ahb1lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1ahb1lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1ahb1lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1ahb1lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1ahb1lpenrFieldDma1lpenShift = 0
	RegisterC1ahb1lpenrFieldDma1lpenMask  = 0x1
)

// GetDma1lpen DMA1 Clock Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) GetDma1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1lpenrFieldDma1lpenMask) != 0
}

// SetDma1lpen DMA1 Clock Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) SetDma1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1lpenrFieldDma1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1lpenrFieldDma1lpenMask)
	}
}

const (
	RegisterC1ahb1lpenrFieldDma2lpenShift = 1
	RegisterC1ahb1lpenrFieldDma2lpenMask  = 0x2
)

// GetDma2lpen DMA2 Clock Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) GetDma2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1lpenrFieldDma2lpenMask) != 0
}

// SetDma2lpen DMA2 Clock Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) SetDma2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1lpenrFieldDma2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1lpenrFieldDma2lpenMask)
	}
}

const (
	RegisterC1ahb1lpenrFieldAdc12lpenShift = 5
	RegisterC1ahb1lpenrFieldAdc12lpenMask  = 0x20
)

// GetAdc12lpen ADC1/2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) GetAdc12lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1lpenrFieldAdc12lpenMask) != 0
}

// SetAdc12lpen ADC1/2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) SetAdc12lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1lpenrFieldAdc12lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1lpenrFieldAdc12lpenMask)
	}
}

const (
	RegisterC1ahb1lpenrFieldEth1maclpenShift = 15
	RegisterC1ahb1lpenrFieldEth1maclpenMask  = 0x8000
)

// GetEth1maclpen Ethernet MAC bus interface Clock Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) GetEth1maclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1lpenrFieldEth1maclpenMask) != 0
}

// SetEth1maclpen Ethernet MAC bus interface Clock Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) SetEth1maclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1lpenrFieldEth1maclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1lpenrFieldEth1maclpenMask)
	}
}

const (
	RegisterC1ahb1lpenrFieldEth1txlpenShift = 16
	RegisterC1ahb1lpenrFieldEth1txlpenMask  = 0x10000
)

// GetEth1txlpen Ethernet Transmission Clock Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) GetEth1txlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1lpenrFieldEth1txlpenMask) != 0
}

// SetEth1txlpen Ethernet Transmission Clock Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) SetEth1txlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1lpenrFieldEth1txlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1lpenrFieldEth1txlpenMask)
	}
}

const (
	RegisterC1ahb1lpenrFieldEth1rxlpenShift = 17
	RegisterC1ahb1lpenrFieldEth1rxlpenMask  = 0x20000
)

// GetEth1rxlpen Ethernet Reception Clock Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) GetEth1rxlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1lpenrFieldEth1rxlpenMask) != 0
}

// SetEth1rxlpen Ethernet Reception Clock Enable During CSleep Mode
func (r *RegisterC1ahb1lpenrType) SetEth1rxlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1lpenrFieldEth1rxlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1lpenrFieldEth1rxlpenMask)
	}
}

const (
	RegisterC1ahb1lpenrFieldUsb1otglpenShift = 25
	RegisterC1ahb1lpenrFieldUsb1otglpenMask  = 0x2000000
)

// GetUsb1otglpen USB1OTG peripheral clock enable during CSleep mode
func (r *RegisterC1ahb1lpenrType) GetUsb1otglpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1lpenrFieldUsb1otglpenMask) != 0
}

// SetUsb1otglpen USB1OTG peripheral clock enable during CSleep mode
func (r *RegisterC1ahb1lpenrType) SetUsb1otglpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1lpenrFieldUsb1otglpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1lpenrFieldUsb1otglpenMask)
	}
}

const (
	RegisterC1ahb1lpenrFieldUsb1ulpilpenShift = 26
	RegisterC1ahb1lpenrFieldUsb1ulpilpenMask  = 0x4000000
)

// GetUsb1ulpilpen USB_PHY1 clock enable during CSleep mode
func (r *RegisterC1ahb1lpenrType) GetUsb1ulpilpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1lpenrFieldUsb1ulpilpenMask) != 0
}

// SetUsb1ulpilpen USB_PHY1 clock enable during CSleep mode
func (r *RegisterC1ahb1lpenrType) SetUsb1ulpilpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1lpenrFieldUsb1ulpilpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1lpenrFieldUsb1ulpilpenMask)
	}
}

const (
	RegisterC1ahb1lpenrFieldUsb2otglpenShift = 27
	RegisterC1ahb1lpenrFieldUsb2otglpenMask  = 0x8000000
)

// GetUsb2otglpen USB2OTG peripheral clock enable during CSleep mode
func (r *RegisterC1ahb1lpenrType) GetUsb2otglpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1lpenrFieldUsb2otglpenMask) != 0
}

// SetUsb2otglpen USB2OTG peripheral clock enable during CSleep mode
func (r *RegisterC1ahb1lpenrType) SetUsb2otglpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1lpenrFieldUsb2otglpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1lpenrFieldUsb2otglpenMask)
	}
}

const (
	RegisterC1ahb1lpenrFieldUsb2ulpilpenShift = 28
	RegisterC1ahb1lpenrFieldUsb2ulpilpenMask  = 0x10000000
)

// GetUsb2ulpilpen USB_PHY2 clocks enable during CSleep mode
func (r *RegisterC1ahb1lpenrType) GetUsb2ulpilpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb1lpenrFieldUsb2ulpilpenMask) != 0
}

// SetUsb2ulpilpen USB_PHY2 clocks enable during CSleep mode
func (r *RegisterC1ahb1lpenrType) SetUsb2ulpilpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb1lpenrFieldUsb2ulpilpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb1lpenrFieldUsb2ulpilpenMask)
	}
}

// RegisterC1ahb2lpenrType RCC AHB2 Sleep Clock Register
type RegisterC1ahb2lpenrType uint32

func (r *RegisterC1ahb2lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1ahb2lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1ahb2lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1ahb2lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1ahb2lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1ahb2lpenrFieldCamitflpenShift = 0
	RegisterC1ahb2lpenrFieldCamitflpenMask  = 0x1
)

// GetCamitflpen CAMITF peripheral clock enable during CSleep mode
func (r *RegisterC1ahb2lpenrType) GetCamitflpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2lpenrFieldCamitflpenMask) != 0
}

// SetCamitflpen CAMITF peripheral clock enable during CSleep mode
func (r *RegisterC1ahb2lpenrType) SetCamitflpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2lpenrFieldCamitflpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2lpenrFieldCamitflpenMask)
	}
}

const (
	RegisterC1ahb2lpenrFieldCryptlpenShift = 4
	RegisterC1ahb2lpenrFieldCryptlpenMask  = 0x10
)

// GetCryptlpen CRYPT peripheral clock enable during CSleep mode
func (r *RegisterC1ahb2lpenrType) GetCryptlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2lpenrFieldCryptlpenMask) != 0
}

// SetCryptlpen CRYPT peripheral clock enable during CSleep mode
func (r *RegisterC1ahb2lpenrType) SetCryptlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2lpenrFieldCryptlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2lpenrFieldCryptlpenMask)
	}
}

const (
	RegisterC1ahb2lpenrFieldHashlpenShift = 5
	RegisterC1ahb2lpenrFieldHashlpenMask  = 0x20
)

// GetHashlpen HASH peripheral clock enable during CSleep mode
func (r *RegisterC1ahb2lpenrType) GetHashlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2lpenrFieldHashlpenMask) != 0
}

// SetHashlpen HASH peripheral clock enable during CSleep mode
func (r *RegisterC1ahb2lpenrType) SetHashlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2lpenrFieldHashlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2lpenrFieldHashlpenMask)
	}
}

const (
	RegisterC1ahb2lpenrFieldRnglpenShift = 6
	RegisterC1ahb2lpenrFieldRnglpenMask  = 0x40
)

// GetRnglpen RNG peripheral clock enable during CSleep mode
func (r *RegisterC1ahb2lpenrType) GetRnglpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2lpenrFieldRnglpenMask) != 0
}

// SetRnglpen RNG peripheral clock enable during CSleep mode
func (r *RegisterC1ahb2lpenrType) SetRnglpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2lpenrFieldRnglpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2lpenrFieldRnglpenMask)
	}
}

const (
	RegisterC1ahb2lpenrFieldSdmmc2lpenShift = 9
	RegisterC1ahb2lpenrFieldSdmmc2lpenMask  = 0x200
)

// GetSdmmc2lpen SDMMC2 and SDMMC2 Delay Clock Enable During CSleep Mode
func (r *RegisterC1ahb2lpenrType) GetSdmmc2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2lpenrFieldSdmmc2lpenMask) != 0
}

// SetSdmmc2lpen SDMMC2 and SDMMC2 Delay Clock Enable During CSleep Mode
func (r *RegisterC1ahb2lpenrType) SetSdmmc2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2lpenrFieldSdmmc2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2lpenrFieldSdmmc2lpenMask)
	}
}

const (
	RegisterC1ahb2lpenrFieldSram1lpenShift = 29
	RegisterC1ahb2lpenrFieldSram1lpenMask  = 0x20000000
)

// GetSram1lpen SRAM1 Clock Enable During CSleep Mode
func (r *RegisterC1ahb2lpenrType) GetSram1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2lpenrFieldSram1lpenMask) != 0
}

// SetSram1lpen SRAM1 Clock Enable During CSleep Mode
func (r *RegisterC1ahb2lpenrType) SetSram1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2lpenrFieldSram1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2lpenrFieldSram1lpenMask)
	}
}

const (
	RegisterC1ahb2lpenrFieldSram2lpenShift = 30
	RegisterC1ahb2lpenrFieldSram2lpenMask  = 0x40000000
)

// GetSram2lpen SRAM2 Clock Enable During CSleep Mode
func (r *RegisterC1ahb2lpenrType) GetSram2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2lpenrFieldSram2lpenMask) != 0
}

// SetSram2lpen SRAM2 Clock Enable During CSleep Mode
func (r *RegisterC1ahb2lpenrType) SetSram2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2lpenrFieldSram2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2lpenrFieldSram2lpenMask)
	}
}

const (
	RegisterC1ahb2lpenrFieldSram3lpenShift = 31
	RegisterC1ahb2lpenrFieldSram3lpenMask  = 0x80000000
)

// GetSram3lpen SRAM3 Clock Enable During CSleep Mode
func (r *RegisterC1ahb2lpenrType) GetSram3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb2lpenrFieldSram3lpenMask) != 0
}

// SetSram3lpen SRAM3 Clock Enable During CSleep Mode
func (r *RegisterC1ahb2lpenrType) SetSram3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb2lpenrFieldSram3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb2lpenrFieldSram3lpenMask)
	}
}

// RegisterC1ahb4lpenrType RCC AHB4 Sleep Clock Register
type RegisterC1ahb4lpenrType uint32

func (r *RegisterC1ahb4lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1ahb4lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1ahb4lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1ahb4lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1ahb4lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1ahb4lpenrFieldGpioalpenShift = 0
	RegisterC1ahb4lpenrFieldGpioalpenMask  = 0x1
)

// GetGpioalpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetGpioalpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldGpioalpenMask) != 0
}

// SetGpioalpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetGpioalpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldGpioalpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldGpioalpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldGpioblpenShift = 1
	RegisterC1ahb4lpenrFieldGpioblpenMask  = 0x2
)

// GetGpioblpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetGpioblpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldGpioblpenMask) != 0
}

// SetGpioblpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetGpioblpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldGpioblpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldGpioblpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldGpioclpenShift = 2
	RegisterC1ahb4lpenrFieldGpioclpenMask  = 0x4
)

// GetGpioclpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetGpioclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldGpioclpenMask) != 0
}

// SetGpioclpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetGpioclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldGpioclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldGpioclpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldGpiodlpenShift = 3
	RegisterC1ahb4lpenrFieldGpiodlpenMask  = 0x8
)

// GetGpiodlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetGpiodlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldGpiodlpenMask) != 0
}

// SetGpiodlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetGpiodlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldGpiodlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldGpiodlpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldGpioelpenShift = 4
	RegisterC1ahb4lpenrFieldGpioelpenMask  = 0x10
)

// GetGpioelpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetGpioelpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldGpioelpenMask) != 0
}

// SetGpioelpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetGpioelpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldGpioelpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldGpioelpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldGpioflpenShift = 5
	RegisterC1ahb4lpenrFieldGpioflpenMask  = 0x20
)

// GetGpioflpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetGpioflpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldGpioflpenMask) != 0
}

// SetGpioflpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetGpioflpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldGpioflpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldGpioflpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldGpioglpenShift = 6
	RegisterC1ahb4lpenrFieldGpioglpenMask  = 0x40
)

// GetGpioglpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetGpioglpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldGpioglpenMask) != 0
}

// SetGpioglpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetGpioglpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldGpioglpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldGpioglpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldGpiohlpenShift = 7
	RegisterC1ahb4lpenrFieldGpiohlpenMask  = 0x80
)

// GetGpiohlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetGpiohlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldGpiohlpenMask) != 0
}

// SetGpiohlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetGpiohlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldGpiohlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldGpiohlpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldGpioilpenShift = 8
	RegisterC1ahb4lpenrFieldGpioilpenMask  = 0x100
)

// GetGpioilpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetGpioilpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldGpioilpenMask) != 0
}

// SetGpioilpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetGpioilpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldGpioilpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldGpioilpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldGpiojlpenShift = 9
	RegisterC1ahb4lpenrFieldGpiojlpenMask  = 0x200
)

// GetGpiojlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetGpiojlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldGpiojlpenMask) != 0
}

// SetGpiojlpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetGpiojlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldGpiojlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldGpiojlpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldGpioklpenShift = 10
	RegisterC1ahb4lpenrFieldGpioklpenMask  = 0x400
)

// GetGpioklpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetGpioklpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldGpioklpenMask) != 0
}

// SetGpioklpen GPIO peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetGpioklpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldGpioklpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldGpioklpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldCrclpenShift = 19
	RegisterC1ahb4lpenrFieldCrclpenMask  = 0x80000
)

// GetCrclpen CRC peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) GetCrclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldCrclpenMask) != 0
}

// SetCrclpen CRC peripheral clock enable during CSleep mode
func (r *RegisterC1ahb4lpenrType) SetCrclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldCrclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldCrclpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldBdmalpenShift = 21
	RegisterC1ahb4lpenrFieldBdmalpenMask  = 0x200000
)

// GetBdmalpen BDMA Clock Enable During CSleep Mode
func (r *RegisterC1ahb4lpenrType) GetBdmalpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldBdmalpenMask) != 0
}

// SetBdmalpen BDMA Clock Enable During CSleep Mode
func (r *RegisterC1ahb4lpenrType) SetBdmalpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldBdmalpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldBdmalpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldAdc3lpenShift = 24
	RegisterC1ahb4lpenrFieldAdc3lpenMask  = 0x1000000
)

// GetAdc3lpen ADC3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1ahb4lpenrType) GetAdc3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldAdc3lpenMask) != 0
}

// SetAdc3lpen ADC3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1ahb4lpenrType) SetAdc3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldAdc3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldAdc3lpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldBkpramlpenShift = 28
	RegisterC1ahb4lpenrFieldBkpramlpenMask  = 0x10000000
)

// GetBkpramlpen Backup RAM Clock Enable During CSleep Mode
func (r *RegisterC1ahb4lpenrType) GetBkpramlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldBkpramlpenMask) != 0
}

// SetBkpramlpen Backup RAM Clock Enable During CSleep Mode
func (r *RegisterC1ahb4lpenrType) SetBkpramlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldBkpramlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldBkpramlpenMask)
	}
}

const (
	RegisterC1ahb4lpenrFieldSram4lpenShift = 29
	RegisterC1ahb4lpenrFieldSram4lpenMask  = 0x20000000
)

// GetSram4lpen SRAM4 Clock Enable During CSleep Mode
func (r *RegisterC1ahb4lpenrType) GetSram4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1ahb4lpenrFieldSram4lpenMask) != 0
}

// SetSram4lpen SRAM4 Clock Enable During CSleep Mode
func (r *RegisterC1ahb4lpenrType) SetSram4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1ahb4lpenrFieldSram4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1ahb4lpenrFieldSram4lpenMask)
	}
}

// RegisterC1apb3lpenrType RCC APB3 Sleep Clock Register
type RegisterC1apb3lpenrType uint32

func (r *RegisterC1apb3lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1apb3lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1apb3lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1apb3lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1apb3lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1apb3lpenrFieldLtdclpenShift = 3
	RegisterC1apb3lpenrFieldLtdclpenMask  = 0x8
)

// GetLtdclpen LTDC peripheral clock enable during CSleep mode
func (r *RegisterC1apb3lpenrType) GetLtdclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb3lpenrFieldLtdclpenMask) != 0
}

// SetLtdclpen LTDC peripheral clock enable during CSleep mode
func (r *RegisterC1apb3lpenrType) SetLtdclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb3lpenrFieldLtdclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb3lpenrFieldLtdclpenMask)
	}
}

const (
	RegisterC1apb3lpenrFieldWwdg1lpenShift = 6
	RegisterC1apb3lpenrFieldWwdg1lpenMask  = 0x40
)

// GetWwdg1lpen WWDG1 Clock Enable During CSleep Mode
func (r *RegisterC1apb3lpenrType) GetWwdg1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb3lpenrFieldWwdg1lpenMask) != 0
}

// SetWwdg1lpen WWDG1 Clock Enable During CSleep Mode
func (r *RegisterC1apb3lpenrType) SetWwdg1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb3lpenrFieldWwdg1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb3lpenrFieldWwdg1lpenMask)
	}
}

// RegisterC1apb1llpenrType RCC APB1 Low Sleep Clock Register
type RegisterC1apb1llpenrType uint32

func (r *RegisterC1apb1llpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1apb1llpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1apb1llpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1apb1llpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1apb1llpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1apb1llpenrFieldTim2lpenShift = 0
	RegisterC1apb1llpenrFieldTim2lpenMask  = 0x1
)

// GetTim2lpen TIM2 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) GetTim2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldTim2lpenMask) != 0
}

// SetTim2lpen TIM2 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) SetTim2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldTim2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldTim2lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldTim3lpenShift = 1
	RegisterC1apb1llpenrFieldTim3lpenMask  = 0x2
)

// GetTim3lpen TIM3 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) GetTim3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldTim3lpenMask) != 0
}

// SetTim3lpen TIM3 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) SetTim3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldTim3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldTim3lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldTim4lpenShift = 2
	RegisterC1apb1llpenrFieldTim4lpenMask  = 0x4
)

// GetTim4lpen TIM4 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) GetTim4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldTim4lpenMask) != 0
}

// SetTim4lpen TIM4 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) SetTim4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldTim4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldTim4lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldTim5lpenShift = 3
	RegisterC1apb1llpenrFieldTim5lpenMask  = 0x8
)

// GetTim5lpen TIM5 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) GetTim5lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldTim5lpenMask) != 0
}

// SetTim5lpen TIM5 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) SetTim5lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldTim5lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldTim5lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldTim6lpenShift = 4
	RegisterC1apb1llpenrFieldTim6lpenMask  = 0x10
)

// GetTim6lpen TIM6 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) GetTim6lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldTim6lpenMask) != 0
}

// SetTim6lpen TIM6 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) SetTim6lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldTim6lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldTim6lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldTim7lpenShift = 5
	RegisterC1apb1llpenrFieldTim7lpenMask  = 0x20
)

// GetTim7lpen TIM7 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) GetTim7lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldTim7lpenMask) != 0
}

// SetTim7lpen TIM7 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) SetTim7lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldTim7lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldTim7lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldTim12lpenShift = 6
	RegisterC1apb1llpenrFieldTim12lpenMask  = 0x40
)

// GetTim12lpen TIM12 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) GetTim12lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldTim12lpenMask) != 0
}

// SetTim12lpen TIM12 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) SetTim12lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldTim12lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldTim12lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldTim13lpenShift = 7
	RegisterC1apb1llpenrFieldTim13lpenMask  = 0x80
)

// GetTim13lpen TIM13 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) GetTim13lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldTim13lpenMask) != 0
}

// SetTim13lpen TIM13 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) SetTim13lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldTim13lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldTim13lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldTim14lpenShift = 8
	RegisterC1apb1llpenrFieldTim14lpenMask  = 0x100
)

// GetTim14lpen TIM14 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) GetTim14lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldTim14lpenMask) != 0
}

// SetTim14lpen TIM14 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) SetTim14lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldTim14lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldTim14lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldLptim1lpenShift = 9
	RegisterC1apb1llpenrFieldLptim1lpenMask  = 0x200
)

// GetLptim1lpen LPTIM1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetLptim1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldLptim1lpenMask) != 0
}

// SetLptim1lpen LPTIM1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetLptim1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldLptim1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldLptim1lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldSpi2lpenShift = 14
	RegisterC1apb1llpenrFieldSpi2lpenMask  = 0x4000
)

// GetSpi2lpen SPI2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetSpi2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldSpi2lpenMask) != 0
}

// SetSpi2lpen SPI2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetSpi2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldSpi2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldSpi2lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldSpi3lpenShift = 15
	RegisterC1apb1llpenrFieldSpi3lpenMask  = 0x8000
)

// GetSpi3lpen SPI3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetSpi3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldSpi3lpenMask) != 0
}

// SetSpi3lpen SPI3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetSpi3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldSpi3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldSpi3lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldSpdifrxlpenShift = 16
	RegisterC1apb1llpenrFieldSpdifrxlpenMask  = 0x10000
)

// GetSpdifrxlpen SPDIFRX Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetSpdifrxlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldSpdifrxlpenMask) != 0
}

// SetSpdifrxlpen SPDIFRX Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetSpdifrxlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldSpdifrxlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldSpdifrxlpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldUsart2lpenShift = 17
	RegisterC1apb1llpenrFieldUsart2lpenMask  = 0x20000
)

// GetUsart2lpen USART2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetUsart2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldUsart2lpenMask) != 0
}

// SetUsart2lpen USART2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetUsart2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldUsart2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldUsart2lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldUsart3lpenShift = 18
	RegisterC1apb1llpenrFieldUsart3lpenMask  = 0x40000
)

// GetUsart3lpen USART3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetUsart3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldUsart3lpenMask) != 0
}

// SetUsart3lpen USART3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetUsart3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldUsart3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldUsart3lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldUart4lpenShift = 19
	RegisterC1apb1llpenrFieldUart4lpenMask  = 0x80000
)

// GetUart4lpen UART4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetUart4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldUart4lpenMask) != 0
}

// SetUart4lpen UART4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetUart4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldUart4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldUart4lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldUart5lpenShift = 20
	RegisterC1apb1llpenrFieldUart5lpenMask  = 0x100000
)

// GetUart5lpen UART5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetUart5lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldUart5lpenMask) != 0
}

// SetUart5lpen UART5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetUart5lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldUart5lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldUart5lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldI2c1lpenShift = 21
	RegisterC1apb1llpenrFieldI2c1lpenMask  = 0x200000
)

// GetI2c1lpen I2C1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetI2c1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldI2c1lpenMask) != 0
}

// SetI2c1lpen I2C1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetI2c1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldI2c1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldI2c1lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldI2c2lpenShift = 22
	RegisterC1apb1llpenrFieldI2c2lpenMask  = 0x400000
)

// GetI2c2lpen I2C2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetI2c2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldI2c2lpenMask) != 0
}

// SetI2c2lpen I2C2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetI2c2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldI2c2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldI2c2lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldI2c3lpenShift = 23
	RegisterC1apb1llpenrFieldI2c3lpenMask  = 0x800000
)

// GetI2c3lpen I2C3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetI2c3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldI2c3lpenMask) != 0
}

// SetI2c3lpen I2C3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetI2c3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldI2c3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldI2c3lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldHdmiceclpenShift = 27
	RegisterC1apb1llpenrFieldHdmiceclpenMask  = 0x8000000
)

// GetHdmiceclpen HDMI-CEC Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetHdmiceclpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldHdmiceclpenMask) != 0
}

// SetHdmiceclpen HDMI-CEC Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetHdmiceclpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldHdmiceclpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldHdmiceclpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldDac12lpenShift = 29
	RegisterC1apb1llpenrFieldDac12lpenMask  = 0x20000000
)

// GetDac12lpen DAC1/2 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) GetDac12lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldDac12lpenMask) != 0
}

// SetDac12lpen DAC1/2 peripheral clock enable during CSleep mode
func (r *RegisterC1apb1llpenrType) SetDac12lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldDac12lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldDac12lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldUsart7lpenShift = 30
	RegisterC1apb1llpenrFieldUsart7lpenMask  = 0x40000000
)

// GetUsart7lpen USART7 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetUsart7lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldUsart7lpenMask) != 0
}

// SetUsart7lpen USART7 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetUsart7lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldUsart7lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldUsart7lpenMask)
	}
}

const (
	RegisterC1apb1llpenrFieldUsart8lpenShift = 31
	RegisterC1apb1llpenrFieldUsart8lpenMask  = 0x80000000
)

// GetUsart8lpen USART8 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) GetUsart8lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1llpenrFieldUsart8lpenMask) != 0
}

// SetUsart8lpen USART8 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1llpenrType) SetUsart8lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1llpenrFieldUsart8lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1llpenrFieldUsart8lpenMask)
	}
}

// RegisterC1apb1hlpenrType RCC APB1 High Sleep Clock Register
type RegisterC1apb1hlpenrType uint32

func (r *RegisterC1apb1hlpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1apb1hlpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1apb1hlpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1apb1hlpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1apb1hlpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1apb1hlpenrFieldCrslpenShift = 1
	RegisterC1apb1hlpenrFieldCrslpenMask  = 0x2
)

// GetCrslpen Clock Recovery System peripheral clock enable during CSleep mode
func (r *RegisterC1apb1hlpenrType) GetCrslpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1hlpenrFieldCrslpenMask) != 0
}

// SetCrslpen Clock Recovery System peripheral clock enable during CSleep mode
func (r *RegisterC1apb1hlpenrType) SetCrslpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1hlpenrFieldCrslpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1hlpenrFieldCrslpenMask)
	}
}

const (
	RegisterC1apb1hlpenrFieldSwplpenShift = 2
	RegisterC1apb1hlpenrFieldSwplpenMask  = 0x4
)

// GetSwplpen SWPMI Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1hlpenrType) GetSwplpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1hlpenrFieldSwplpenMask) != 0
}

// SetSwplpen SWPMI Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1hlpenrType) SetSwplpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1hlpenrFieldSwplpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1hlpenrFieldSwplpenMask)
	}
}

const (
	RegisterC1apb1hlpenrFieldOpamplpenShift = 4
	RegisterC1apb1hlpenrFieldOpamplpenMask  = 0x10
)

// GetOpamplpen OPAMP peripheral clock enable during CSleep mode
func (r *RegisterC1apb1hlpenrType) GetOpamplpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1hlpenrFieldOpamplpenMask) != 0
}

// SetOpamplpen OPAMP peripheral clock enable during CSleep mode
func (r *RegisterC1apb1hlpenrType) SetOpamplpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1hlpenrFieldOpamplpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1hlpenrFieldOpamplpenMask)
	}
}

const (
	RegisterC1apb1hlpenrFieldMdioslpenShift = 5
	RegisterC1apb1hlpenrFieldMdioslpenMask  = 0x20
)

// GetMdioslpen MDIOS peripheral clock enable during CSleep mode
func (r *RegisterC1apb1hlpenrType) GetMdioslpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1hlpenrFieldMdioslpenMask) != 0
}

// SetMdioslpen MDIOS peripheral clock enable during CSleep mode
func (r *RegisterC1apb1hlpenrType) SetMdioslpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1hlpenrFieldMdioslpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1hlpenrFieldMdioslpenMask)
	}
}

const (
	RegisterC1apb1hlpenrFieldFdcanlpenShift = 8
	RegisterC1apb1hlpenrFieldFdcanlpenMask  = 0x100
)

// GetFdcanlpen FDCAN Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1hlpenrType) GetFdcanlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb1hlpenrFieldFdcanlpenMask) != 0
}

// SetFdcanlpen FDCAN Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb1hlpenrType) SetFdcanlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb1hlpenrFieldFdcanlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb1hlpenrFieldFdcanlpenMask)
	}
}

// RegisterC1apb2lpenrType RCC APB2 Sleep Clock Register
type RegisterC1apb2lpenrType uint32

func (r *RegisterC1apb2lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1apb2lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1apb2lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1apb2lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1apb2lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1apb2lpenrFieldTim1lpenShift = 0
	RegisterC1apb2lpenrFieldTim1lpenMask  = 0x1
)

// GetTim1lpen TIM1 peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) GetTim1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldTim1lpenMask) != 0
}

// SetTim1lpen TIM1 peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) SetTim1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldTim1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldTim1lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldTim8lpenShift = 1
	RegisterC1apb2lpenrFieldTim8lpenMask  = 0x2
)

// GetTim8lpen TIM8 peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) GetTim8lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldTim8lpenMask) != 0
}

// SetTim8lpen TIM8 peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) SetTim8lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldTim8lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldTim8lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldUsart1lpenShift = 4
	RegisterC1apb2lpenrFieldUsart1lpenMask  = 0x10
)

// GetUsart1lpen USART1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) GetUsart1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldUsart1lpenMask) != 0
}

// SetUsart1lpen USART1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) SetUsart1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldUsart1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldUsart1lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldUsart6lpenShift = 5
	RegisterC1apb2lpenrFieldUsart6lpenMask  = 0x20
)

// GetUsart6lpen USART6 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) GetUsart6lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldUsart6lpenMask) != 0
}

// SetUsart6lpen USART6 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) SetUsart6lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldUsart6lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldUsart6lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldSpi1lpenShift = 12
	RegisterC1apb2lpenrFieldSpi1lpenMask  = 0x1000
)

// GetSpi1lpen SPI1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) GetSpi1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldSpi1lpenMask) != 0
}

// SetSpi1lpen SPI1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) SetSpi1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldSpi1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldSpi1lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldSpi4lpenShift = 13
	RegisterC1apb2lpenrFieldSpi4lpenMask  = 0x2000
)

// GetSpi4lpen SPI4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) GetSpi4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldSpi4lpenMask) != 0
}

// SetSpi4lpen SPI4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) SetSpi4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldSpi4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldSpi4lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldTim15lpenShift = 16
	RegisterC1apb2lpenrFieldTim15lpenMask  = 0x10000
)

// GetTim15lpen TIM15 peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) GetTim15lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldTim15lpenMask) != 0
}

// SetTim15lpen TIM15 peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) SetTim15lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldTim15lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldTim15lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldTim16lpenShift = 17
	RegisterC1apb2lpenrFieldTim16lpenMask  = 0x20000
)

// GetTim16lpen TIM16 peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) GetTim16lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldTim16lpenMask) != 0
}

// SetTim16lpen TIM16 peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) SetTim16lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldTim16lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldTim16lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldTim17lpenShift = 18
	RegisterC1apb2lpenrFieldTim17lpenMask  = 0x40000
)

// GetTim17lpen TIM17 peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) GetTim17lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldTim17lpenMask) != 0
}

// SetTim17lpen TIM17 peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) SetTim17lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldTim17lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldTim17lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldSpi5lpenShift = 20
	RegisterC1apb2lpenrFieldSpi5lpenMask  = 0x100000
)

// GetSpi5lpen SPI5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) GetSpi5lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldSpi5lpenMask) != 0
}

// SetSpi5lpen SPI5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) SetSpi5lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldSpi5lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldSpi5lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldSai1lpenShift = 22
	RegisterC1apb2lpenrFieldSai1lpenMask  = 0x400000
)

// GetSai1lpen SAI1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) GetSai1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldSai1lpenMask) != 0
}

// SetSai1lpen SAI1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) SetSai1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldSai1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldSai1lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldSai2lpenShift = 23
	RegisterC1apb2lpenrFieldSai2lpenMask  = 0x800000
)

// GetSai2lpen SAI2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) GetSai2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldSai2lpenMask) != 0
}

// SetSai2lpen SAI2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) SetSai2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldSai2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldSai2lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldSai3lpenShift = 24
	RegisterC1apb2lpenrFieldSai3lpenMask  = 0x1000000
)

// GetSai3lpen SAI3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) GetSai3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldSai3lpenMask) != 0
}

// SetSai3lpen SAI3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) SetSai3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldSai3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldSai3lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldDfsdm1lpenShift = 28
	RegisterC1apb2lpenrFieldDfsdm1lpenMask  = 0x10000000
)

// GetDfsdm1lpen DFSDM1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) GetDfsdm1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldDfsdm1lpenMask) != 0
}

// SetDfsdm1lpen DFSDM1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb2lpenrType) SetDfsdm1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldDfsdm1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldDfsdm1lpenMask)
	}
}

const (
	RegisterC1apb2lpenrFieldHrtimlpenShift = 29
	RegisterC1apb2lpenrFieldHrtimlpenMask  = 0x20000000
)

// GetHrtimlpen HRTIM peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) GetHrtimlpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb2lpenrFieldHrtimlpenMask) != 0
}

// SetHrtimlpen HRTIM peripheral clock enable during CSleep mode
func (r *RegisterC1apb2lpenrType) SetHrtimlpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb2lpenrFieldHrtimlpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb2lpenrFieldHrtimlpenMask)
	}
}

// RegisterC1apb4lpenrType RCC APB4 Sleep Clock Register
type RegisterC1apb4lpenrType uint32

func (r *RegisterC1apb4lpenrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterC1apb4lpenrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterC1apb4lpenrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterC1apb4lpenrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterC1apb4lpenrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterC1apb4lpenrFieldSyscfglpenShift = 1
	RegisterC1apb4lpenrFieldSyscfglpenMask  = 0x2
)

// GetSyscfglpen SYSCFG peripheral clock enable during CSleep mode
func (r *RegisterC1apb4lpenrType) GetSyscfglpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldSyscfglpenMask) != 0
}

// SetSyscfglpen SYSCFG peripheral clock enable during CSleep mode
func (r *RegisterC1apb4lpenrType) SetSyscfglpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldSyscfglpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldSyscfglpenMask)
	}
}

const (
	RegisterC1apb4lpenrFieldLpuart1lpenShift = 3
	RegisterC1apb4lpenrFieldLpuart1lpenMask  = 0x8
)

// GetLpuart1lpen LPUART1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) GetLpuart1lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldLpuart1lpenMask) != 0
}

// SetLpuart1lpen LPUART1 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) SetLpuart1lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldLpuart1lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldLpuart1lpenMask)
	}
}

const (
	RegisterC1apb4lpenrFieldSpi6lpenShift = 5
	RegisterC1apb4lpenrFieldSpi6lpenMask  = 0x20
)

// GetSpi6lpen SPI6 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) GetSpi6lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldSpi6lpenMask) != 0
}

// SetSpi6lpen SPI6 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) SetSpi6lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldSpi6lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldSpi6lpenMask)
	}
}

const (
	RegisterC1apb4lpenrFieldI2c4lpenShift = 7
	RegisterC1apb4lpenrFieldI2c4lpenMask  = 0x80
)

// GetI2c4lpen I2C4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) GetI2c4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldI2c4lpenMask) != 0
}

// SetI2c4lpen I2C4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) SetI2c4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldI2c4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldI2c4lpenMask)
	}
}

const (
	RegisterC1apb4lpenrFieldLptim2lpenShift = 9
	RegisterC1apb4lpenrFieldLptim2lpenMask  = 0x200
)

// GetLptim2lpen LPTIM2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) GetLptim2lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldLptim2lpenMask) != 0
}

// SetLptim2lpen LPTIM2 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) SetLptim2lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldLptim2lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldLptim2lpenMask)
	}
}

const (
	RegisterC1apb4lpenrFieldLptim3lpenShift = 10
	RegisterC1apb4lpenrFieldLptim3lpenMask  = 0x400
)

// GetLptim3lpen LPTIM3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) GetLptim3lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldLptim3lpenMask) != 0
}

// SetLptim3lpen LPTIM3 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) SetLptim3lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldLptim3lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldLptim3lpenMask)
	}
}

const (
	RegisterC1apb4lpenrFieldLptim4lpenShift = 11
	RegisterC1apb4lpenrFieldLptim4lpenMask  = 0x800
)

// GetLptim4lpen LPTIM4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) GetLptim4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldLptim4lpenMask) != 0
}

// SetLptim4lpen LPTIM4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) SetLptim4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldLptim4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldLptim4lpenMask)
	}
}

const (
	RegisterC1apb4lpenrFieldLptim5lpenShift = 12
	RegisterC1apb4lpenrFieldLptim5lpenMask  = 0x1000
)

// GetLptim5lpen LPTIM5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) GetLptim5lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldLptim5lpenMask) != 0
}

// SetLptim5lpen LPTIM5 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) SetLptim5lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldLptim5lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldLptim5lpenMask)
	}
}

const (
	RegisterC1apb4lpenrFieldComp12lpenShift = 14
	RegisterC1apb4lpenrFieldComp12lpenMask  = 0x4000
)

// GetComp12lpen COMP1/2 peripheral clock enable during CSleep mode
func (r *RegisterC1apb4lpenrType) GetComp12lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldComp12lpenMask) != 0
}

// SetComp12lpen COMP1/2 peripheral clock enable during CSleep mode
func (r *RegisterC1apb4lpenrType) SetComp12lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldComp12lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldComp12lpenMask)
	}
}

const (
	RegisterC1apb4lpenrFieldVreflpenShift = 15
	RegisterC1apb4lpenrFieldVreflpenMask  = 0x8000
)

// GetVreflpen VREF peripheral clock enable during CSleep mode
func (r *RegisterC1apb4lpenrType) GetVreflpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldVreflpenMask) != 0
}

// SetVreflpen VREF peripheral clock enable during CSleep mode
func (r *RegisterC1apb4lpenrType) SetVreflpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldVreflpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldVreflpenMask)
	}
}

const (
	RegisterC1apb4lpenrFieldRtcapblpenShift = 16
	RegisterC1apb4lpenrFieldRtcapblpenMask  = 0x10000
)

// GetRtcapblpen RTC APB Clock Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) GetRtcapblpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldRtcapblpenMask) != 0
}

// SetRtcapblpen RTC APB Clock Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) SetRtcapblpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldRtcapblpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldRtcapblpenMask)
	}
}

const (
	RegisterC1apb4lpenrFieldSai4lpenShift = 21
	RegisterC1apb4lpenrFieldSai4lpenMask  = 0x200000
)

// GetSai4lpen SAI4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) GetSai4lpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterC1apb4lpenrFieldSai4lpenMask) != 0
}

// SetSai4lpen SAI4 Peripheral Clocks Enable During CSleep Mode
func (r *RegisterC1apb4lpenrType) SetSai4lpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterC1apb4lpenrFieldSai4lpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterC1apb4lpenrFieldSai4lpenMask)
	}
}
