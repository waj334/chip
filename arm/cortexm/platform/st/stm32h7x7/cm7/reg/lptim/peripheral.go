//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package lptim

import (
	"unsafe"
	"volatile"
)

var (
	Lptim1 = (*_lptim)(unsafe.Pointer(uintptr(0x40002400)))
	Lptim2 = (*_lptim)(unsafe.Pointer(uintptr(0x58002400)))
	Lptim3 = (*_lptim)(unsafe.Pointer(uintptr(0x58002800)))
	Lptim4 = (*_lptim)(unsafe.Pointer(uintptr(0x58002c00)))
	Lptim5 = (*_lptim)(unsafe.Pointer(uintptr(0x58003000)))

	Instances = [5]*_lptim{
		Lptim1,
		Lptim2,
		Lptim3,
		Lptim4,
		Lptim5,
	}
)

type _lptim struct {
	Isr   RegisterIsrType
	Icr   RegisterIcrType
	Ier   RegisterIerType
	Cfgr  RegisterCfgrType
	Cr    RegisterCrType
	Cmp   RegisterCmpType
	Arr   RegisterArrType
	Cnt   RegisterCntType
	_     [4]byte
	Cfgr2 RegisterCfgr2Type
}

// RegisterIsrType Interrupt and Status Register
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
	RegisterIsrFieldCmpmShift = 0
	RegisterIsrFieldCmpmMask  = 0x1
)

// GetCmpm Compare match
func (r *RegisterIsrType) GetCmpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCmpmMask) != 0
}

// SetCmpm Compare match
func (r *RegisterIsrType) SetCmpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldCmpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldCmpmMask)
	}
}

const (
	RegisterIsrFieldArrmShift = 1
	RegisterIsrFieldArrmMask  = 0x2
)

// GetArrm Autoreload match
func (r *RegisterIsrType) GetArrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldArrmMask) != 0
}

// SetArrm Autoreload match
func (r *RegisterIsrType) SetArrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldArrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldArrmMask)
	}
}

const (
	RegisterIsrFieldExttrigShift = 2
	RegisterIsrFieldExttrigMask  = 0x4
)

// GetExttrig External trigger edge event
func (r *RegisterIsrType) GetExttrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldExttrigMask) != 0
}

// SetExttrig External trigger edge event
func (r *RegisterIsrType) SetExttrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldExttrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldExttrigMask)
	}
}

const (
	RegisterIsrFieldCmpokShift = 3
	RegisterIsrFieldCmpokMask  = 0x8
)

// GetCmpok Compare register update OK
func (r *RegisterIsrType) GetCmpok() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCmpokMask) != 0
}

// SetCmpok Compare register update OK
func (r *RegisterIsrType) SetCmpok(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldCmpokMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldCmpokMask)
	}
}

const (
	RegisterIsrFieldArrokShift = 4
	RegisterIsrFieldArrokMask  = 0x10
)

// GetArrok Autoreload register update OK
func (r *RegisterIsrType) GetArrok() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldArrokMask) != 0
}

// SetArrok Autoreload register update OK
func (r *RegisterIsrType) SetArrok(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldArrokMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldArrokMask)
	}
}

const (
	RegisterIsrFieldUpShift = 5
	RegisterIsrFieldUpMask  = 0x20
)

// GetUp Counter direction change down to up
func (r *RegisterIsrType) GetUp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldUpMask) != 0
}

// SetUp Counter direction change down to up
func (r *RegisterIsrType) SetUp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldUpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldUpMask)
	}
}

const (
	RegisterIsrFieldDownShift = 6
	RegisterIsrFieldDownMask  = 0x40
)

// GetDown Counter direction change up to down
func (r *RegisterIsrType) GetDown() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldDownMask) != 0
}

// SetDown Counter direction change up to down
func (r *RegisterIsrType) SetDown(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldDownMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldDownMask)
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
	RegisterIcrFieldCmpmcfShift = 0
	RegisterIcrFieldCmpmcfMask  = 0x1
)

// GetCmpmcf compare match Clear Flag
func (r *RegisterIcrType) GetCmpmcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCmpmcfMask) != 0
}

// SetCmpmcf compare match Clear Flag
func (r *RegisterIcrType) SetCmpmcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCmpmcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCmpmcfMask)
	}
}

const (
	RegisterIcrFieldArrmcfShift = 1
	RegisterIcrFieldArrmcfMask  = 0x2
)

// GetArrmcf Autoreload match Clear Flag
func (r *RegisterIcrType) GetArrmcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldArrmcfMask) != 0
}

// SetArrmcf Autoreload match Clear Flag
func (r *RegisterIcrType) SetArrmcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldArrmcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldArrmcfMask)
	}
}

const (
	RegisterIcrFieldExttrigcfShift = 2
	RegisterIcrFieldExttrigcfMask  = 0x4
)

// GetExttrigcf External trigger valid edge Clear Flag
func (r *RegisterIcrType) GetExttrigcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldExttrigcfMask) != 0
}

// SetExttrigcf External trigger valid edge Clear Flag
func (r *RegisterIcrType) SetExttrigcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldExttrigcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldExttrigcfMask)
	}
}

const (
	RegisterIcrFieldCmpokcfShift = 3
	RegisterIcrFieldCmpokcfMask  = 0x8
)

// GetCmpokcf Compare register update OK Clear Flag
func (r *RegisterIcrType) GetCmpokcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCmpokcfMask) != 0
}

// SetCmpokcf Compare register update OK Clear Flag
func (r *RegisterIcrType) SetCmpokcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCmpokcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCmpokcfMask)
	}
}

const (
	RegisterIcrFieldArrokcfShift = 4
	RegisterIcrFieldArrokcfMask  = 0x10
)

// GetArrokcf Autoreload register update OK Clear Flag
func (r *RegisterIcrType) GetArrokcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldArrokcfMask) != 0
}

// SetArrokcf Autoreload register update OK Clear Flag
func (r *RegisterIcrType) SetArrokcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldArrokcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldArrokcfMask)
	}
}

const (
	RegisterIcrFieldUpcfShift = 5
	RegisterIcrFieldUpcfMask  = 0x20
)

// GetUpcf Direction change to UP Clear Flag
func (r *RegisterIcrType) GetUpcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldUpcfMask) != 0
}

// SetUpcf Direction change to UP Clear Flag
func (r *RegisterIcrType) SetUpcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldUpcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldUpcfMask)
	}
}

const (
	RegisterIcrFieldDowncfShift = 6
	RegisterIcrFieldDowncfMask  = 0x40
)

// GetDowncf Direction change to down Clear Flag
func (r *RegisterIcrType) GetDowncf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldDowncfMask) != 0
}

// SetDowncf Direction change to down Clear Flag
func (r *RegisterIcrType) SetDowncf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldDowncfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldDowncfMask)
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
	RegisterIerFieldCmpmieShift = 0
	RegisterIerFieldCmpmieMask  = 0x1
)

// GetCmpmie Compare match Interrupt Enable
func (r *RegisterIerType) GetCmpmie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldCmpmieMask) != 0
}

// SetCmpmie Compare match Interrupt Enable
func (r *RegisterIerType) SetCmpmie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldCmpmieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldCmpmieMask)
	}
}

const (
	RegisterIerFieldArrmieShift = 1
	RegisterIerFieldArrmieMask  = 0x2
)

// GetArrmie Autoreload match Interrupt Enable
func (r *RegisterIerType) GetArrmie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldArrmieMask) != 0
}

// SetArrmie Autoreload match Interrupt Enable
func (r *RegisterIerType) SetArrmie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldArrmieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldArrmieMask)
	}
}

const (
	RegisterIerFieldExttrigieShift = 2
	RegisterIerFieldExttrigieMask  = 0x4
)

// GetExttrigie External trigger valid edge Interrupt Enable
func (r *RegisterIerType) GetExttrigie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldExttrigieMask) != 0
}

// SetExttrigie External trigger valid edge Interrupt Enable
func (r *RegisterIerType) SetExttrigie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldExttrigieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldExttrigieMask)
	}
}

const (
	RegisterIerFieldCmpokieShift = 3
	RegisterIerFieldCmpokieMask  = 0x8
)

// GetCmpokie Compare register update OK Interrupt Enable
func (r *RegisterIerType) GetCmpokie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldCmpokieMask) != 0
}

// SetCmpokie Compare register update OK Interrupt Enable
func (r *RegisterIerType) SetCmpokie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldCmpokieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldCmpokieMask)
	}
}

const (
	RegisterIerFieldArrokieShift = 4
	RegisterIerFieldArrokieMask  = 0x10
)

// GetArrokie Autoreload register update OK Interrupt Enable
func (r *RegisterIerType) GetArrokie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldArrokieMask) != 0
}

// SetArrokie Autoreload register update OK Interrupt Enable
func (r *RegisterIerType) SetArrokie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldArrokieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldArrokieMask)
	}
}

const (
	RegisterIerFieldUpieShift = 5
	RegisterIerFieldUpieMask  = 0x20
)

// GetUpie Direction change to UP Interrupt Enable
func (r *RegisterIerType) GetUpie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldUpieMask) != 0
}

// SetUpie Direction change to UP Interrupt Enable
func (r *RegisterIerType) SetUpie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldUpieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldUpieMask)
	}
}

const (
	RegisterIerFieldDownieShift = 6
	RegisterIerFieldDownieMask  = 0x40
)

// GetDownie Direction change to down Interrupt Enable
func (r *RegisterIerType) GetDownie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldDownieMask) != 0
}

// SetDownie Direction change to down Interrupt Enable
func (r *RegisterIerType) SetDownie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldDownieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldDownieMask)
	}
}

// RegisterCfgrType Configuration Register
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

const (
	RegisterCfgrFieldCkselShift = 0
	RegisterCfgrFieldCkselMask  = 0x1
)

// GetCksel Clock selector
func (r *RegisterCfgrType) GetCksel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCkselMask) != 0
}

// SetCksel Clock selector
func (r *RegisterCfgrType) SetCksel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldCkselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldCkselMask)
	}
}

const (
	RegisterCfgrFieldCkpolShift = 1
	RegisterCfgrFieldCkpolMask  = 0x6
)

// GetCkpol Clock Polarity
func (r *RegisterCfgrType) GetCkpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCkpolMask) >> RegisterCfgrFieldCkpolShift)
}

// SetCkpol Clock Polarity
func (r *RegisterCfgrType) SetCkpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldCkpolMask)|(uint32(value)<<RegisterCfgrFieldCkpolShift))
}

const (
	RegisterCfgrFieldCkfltShift = 3
	RegisterCfgrFieldCkfltMask  = 0x18
)

// GetCkflt Configurable digital filter for external clock
func (r *RegisterCfgrType) GetCkflt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCkfltMask) >> RegisterCfgrFieldCkfltShift)
}

// SetCkflt Configurable digital filter for external clock
func (r *RegisterCfgrType) SetCkflt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldCkfltMask)|(uint32(value)<<RegisterCfgrFieldCkfltShift))
}

const (
	RegisterCfgrFieldTrgfltShift = 6
	RegisterCfgrFieldTrgfltMask  = 0xc0
)

// GetTrgflt Configurable digital filter for trigger
func (r *RegisterCfgrType) GetTrgflt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldTrgfltMask) >> RegisterCfgrFieldTrgfltShift)
}

// SetTrgflt Configurable digital filter for trigger
func (r *RegisterCfgrType) SetTrgflt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldTrgfltMask)|(uint32(value)<<RegisterCfgrFieldTrgfltShift))
}

const (
	RegisterCfgrFieldPrescShift = 9
	RegisterCfgrFieldPrescMask  = 0xe00
)

// GetPresc Clock prescaler
func (r *RegisterCfgrType) GetPresc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldPrescMask) >> RegisterCfgrFieldPrescShift)
}

// SetPresc Clock prescaler
func (r *RegisterCfgrType) SetPresc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldPrescMask)|(uint32(value)<<RegisterCfgrFieldPrescShift))
}

const (
	RegisterCfgrFieldTrigselShift = 13
	RegisterCfgrFieldTrigselMask  = 0xe000
)

// GetTrigsel Trigger selector
func (r *RegisterCfgrType) GetTrigsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldTrigselMask) >> RegisterCfgrFieldTrigselShift)
}

// SetTrigsel Trigger selector
func (r *RegisterCfgrType) SetTrigsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldTrigselMask)|(uint32(value)<<RegisterCfgrFieldTrigselShift))
}

const (
	RegisterCfgrFieldTrigenShift = 17
	RegisterCfgrFieldTrigenMask  = 0x60000
)

// GetTrigen Trigger enable and polarity
func (r *RegisterCfgrType) GetTrigen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldTrigenMask) >> RegisterCfgrFieldTrigenShift)
}

// SetTrigen Trigger enable and polarity
func (r *RegisterCfgrType) SetTrigen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldTrigenMask)|(uint32(value)<<RegisterCfgrFieldTrigenShift))
}

const (
	RegisterCfgrFieldTimoutShift = 19
	RegisterCfgrFieldTimoutMask  = 0x80000
)

// GetTimout Timeout enable
func (r *RegisterCfgrType) GetTimout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldTimoutMask) != 0
}

// SetTimout Timeout enable
func (r *RegisterCfgrType) SetTimout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldTimoutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldTimoutMask)
	}
}

const (
	RegisterCfgrFieldWaveShift = 20
	RegisterCfgrFieldWaveMask  = 0x100000
)

// GetWave Waveform shape
func (r *RegisterCfgrType) GetWave() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldWaveMask) != 0
}

// SetWave Waveform shape
func (r *RegisterCfgrType) SetWave(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldWaveMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldWaveMask)
	}
}

const (
	RegisterCfgrFieldWavpolShift = 21
	RegisterCfgrFieldWavpolMask  = 0x200000
)

// GetWavpol Waveform shape polarity
func (r *RegisterCfgrType) GetWavpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldWavpolMask) != 0
}

// SetWavpol Waveform shape polarity
func (r *RegisterCfgrType) SetWavpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldWavpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldWavpolMask)
	}
}

const (
	RegisterCfgrFieldPreloadShift = 22
	RegisterCfgrFieldPreloadMask  = 0x400000
)

// GetPreload Registers update mode
func (r *RegisterCfgrType) GetPreload() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldPreloadMask) != 0
}

// SetPreload Registers update mode
func (r *RegisterCfgrType) SetPreload(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldPreloadMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldPreloadMask)
	}
}

const (
	RegisterCfgrFieldCountmodeShift = 23
	RegisterCfgrFieldCountmodeMask  = 0x800000
)

// GetCountmode counter mode enabled
func (r *RegisterCfgrType) GetCountmode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCountmodeMask) != 0
}

// SetCountmode counter mode enabled
func (r *RegisterCfgrType) SetCountmode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldCountmodeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldCountmodeMask)
	}
}

const (
	RegisterCfgrFieldEncShift = 24
	RegisterCfgrFieldEncMask  = 0x1000000
)

// GetEnc Encoder mode enable
func (r *RegisterCfgrType) GetEnc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldEncMask) != 0
}

// SetEnc Encoder mode enable
func (r *RegisterCfgrType) SetEnc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldEncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldEncMask)
	}
}

// RegisterCrType Control Register
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
	RegisterCrFieldEnableShift = 0
	RegisterCrFieldEnableMask  = 0x1
)

// GetEnable LPTIM Enable
func (r *RegisterCrType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEnableMask) != 0
}

// SetEnable LPTIM Enable
func (r *RegisterCrType) SetEnable(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEnableMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEnableMask)
	}
}

const (
	RegisterCrFieldSngstrtShift = 1
	RegisterCrFieldSngstrtMask  = 0x2
)

// GetSngstrt LPTIM start in single mode
func (r *RegisterCrType) GetSngstrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSngstrtMask) != 0
}

// SetSngstrt LPTIM start in single mode
func (r *RegisterCrType) SetSngstrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldSngstrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldSngstrtMask)
	}
}

const (
	RegisterCrFieldCntstrtShift = 2
	RegisterCrFieldCntstrtMask  = 0x4
)

// GetCntstrt Timer start in continuous mode
func (r *RegisterCrType) GetCntstrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCntstrtMask) != 0
}

// SetCntstrt Timer start in continuous mode
func (r *RegisterCrType) SetCntstrt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCntstrtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCntstrtMask)
	}
}

const (
	RegisterCrFieldCountrstShift = 3
	RegisterCrFieldCountrstMask  = 0x8
)

// GetCountrst Counter reset
func (r *RegisterCrType) GetCountrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCountrstMask) != 0
}

// SetCountrst Counter reset
func (r *RegisterCrType) SetCountrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCountrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCountrstMask)
	}
}

const (
	RegisterCrFieldRstareShift = 4
	RegisterCrFieldRstareMask  = 0x10
)

// GetRstare Reset after read enable
func (r *RegisterCrType) GetRstare() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRstareMask) != 0
}

// SetRstare Reset after read enable
func (r *RegisterCrType) SetRstare(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRstareMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRstareMask)
	}
}

// RegisterCmpType Compare Register
type RegisterCmpType uint32

func (r *RegisterCmpType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCmpType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCmpType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCmpType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCmpType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCmpFieldCmpShift = 0
	RegisterCmpFieldCmpMask  = 0xffff
)

// GetCmp Compare value
func (r *RegisterCmpType) GetCmp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmpFieldCmpMask) >> RegisterCmpFieldCmpShift)
}

// SetCmp Compare value
func (r *RegisterCmpType) SetCmp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmpFieldCmpMask)|(uint32(value)<<RegisterCmpFieldCmpShift))
}

// RegisterArrType Autoreload Register
type RegisterArrType uint32

func (r *RegisterArrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterArrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterArrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterArrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterArrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterArrFieldArrShift = 0
	RegisterArrFieldArrMask  = 0xffff
)

// GetArr Auto reload value
func (r *RegisterArrType) GetArr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterArrFieldArrMask) >> RegisterArrFieldArrShift)
}

// SetArr Auto reload value
func (r *RegisterArrType) SetArr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArrFieldArrMask)|(uint32(value)<<RegisterArrFieldArrShift))
}

// RegisterCntType Counter Register
type RegisterCntType uint32

func (r *RegisterCntType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCntType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCntType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCntType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCntType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCntFieldCntShift = 0
	RegisterCntFieldCntMask  = 0xffff
)

// GetCnt Counter value
func (r *RegisterCntType) GetCnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldCntMask) >> RegisterCntFieldCntShift)
}

// SetCnt Counter value
func (r *RegisterCntType) SetCnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldCntMask)|(uint32(value)<<RegisterCntFieldCntShift))
}

// RegisterCfgr2Type LPTIM configuration register 2
type RegisterCfgr2Type uint32

func (r *RegisterCfgr2Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCfgr2Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCfgr2Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCfgr2Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCfgr2Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCfgr2FieldIn1selShift = 0
	RegisterCfgr2FieldIn1selMask  = 0x3
)

// GetIn1sel LPTIM Input 1 selection
func (r *RegisterCfgr2Type) GetIn1sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldIn1selMask) >> RegisterCfgr2FieldIn1selShift)
}

// SetIn1sel LPTIM Input 1 selection
func (r *RegisterCfgr2Type) SetIn1sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldIn1selMask)|(uint32(value)<<RegisterCfgr2FieldIn1selShift))
}

const (
	RegisterCfgr2FieldIn2selShift = 4
	RegisterCfgr2FieldIn2selMask  = 0x30
)

// GetIn2sel LPTIM Input 2 selection
func (r *RegisterCfgr2Type) GetIn2sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldIn2selMask) >> RegisterCfgr2FieldIn2selShift)
}

// SetIn2sel LPTIM Input 2 selection
func (r *RegisterCfgr2Type) SetIn2sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldIn2selMask)|(uint32(value)<<RegisterCfgr2FieldIn2selShift))
}
