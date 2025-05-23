package lptim3

import (
	"unsafe"
	"volatile"
)

var (
	Lptim3 = (*_lptim3)(unsafe.Pointer(uintptr(0x58002800)))
	Lptim4 = (*_lptim3)(unsafe.Pointer(uintptr(0x58002c00)))
	Lptim5 = (*_lptim3)(unsafe.Pointer(uintptr(0x58003000)))
)

type _lptim3 struct {
	Isr   registerIsrType
	Icr   registerIcrType
	Ier   registerIerType
	Cfgr  registerCfgrType
	Cr    registerCrType
	Cmp   registerCmpType
	Arr   registerArrType
	Cnt   registerCntType
	_     [4]byte
	Cfgr2 registerCfgr2Type
}

// registerIsrType Interrupt and Status Register
type registerIsrType uint32

const (
	RegisterIsrFieldDownShift = 6
	RegisterIsrFieldDownMask  = 0x40
)

// GetDown Counter direction change up to down
func (r *registerIsrType) GetDown() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldDownMask) != 0
}

// SetDown Counter direction change up to down
func (r *registerIsrType) SetDown(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldDownMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldDownMask)
	}
}

const (
	RegisterIsrFieldUpShift = 5
	RegisterIsrFieldUpMask  = 0x20
)

// GetUp Counter direction change down to up
func (r *registerIsrType) GetUp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldUpMask) != 0
}

// SetUp Counter direction change down to up
func (r *registerIsrType) SetUp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldUpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldUpMask)
	}
}

const (
	RegisterIsrFieldArrokShift = 4
	RegisterIsrFieldArrokMask  = 0x10
)

// GetArrok Autoreload register update OK
func (r *registerIsrType) GetArrok() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldArrokMask) != 0
}

// SetArrok Autoreload register update OK
func (r *registerIsrType) SetArrok(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldArrokMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldArrokMask)
	}
}

const (
	RegisterIsrFieldCmpokShift = 3
	RegisterIsrFieldCmpokMask  = 0x8
)

// GetCmpok Compare register update OK
func (r *registerIsrType) GetCmpok() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCmpokMask) != 0
}

// SetCmpok Compare register update OK
func (r *registerIsrType) SetCmpok(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldCmpokMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldCmpokMask)
	}
}

const (
	RegisterIsrFieldExttrigShift = 2
	RegisterIsrFieldExttrigMask  = 0x4
)

// GetExttrig External trigger edge event
func (r *registerIsrType) GetExttrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldExttrigMask) != 0
}

// SetExttrig External trigger edge event
func (r *registerIsrType) SetExttrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldExttrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldExttrigMask)
	}
}

const (
	RegisterIsrFieldArrmShift = 1
	RegisterIsrFieldArrmMask  = 0x2
)

// GetArrm Autoreload match
func (r *registerIsrType) GetArrm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldArrmMask) != 0
}

// SetArrm Autoreload match
func (r *registerIsrType) SetArrm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldArrmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldArrmMask)
	}
}

const (
	RegisterIsrFieldCmpmShift = 0
	RegisterIsrFieldCmpmMask  = 0x1
)

// GetCmpm Compare match
func (r *registerIsrType) GetCmpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldCmpmMask) != 0
}

// SetCmpm Compare match
func (r *registerIsrType) SetCmpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldCmpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldCmpmMask)
	}
}

// registerIcrType Interrupt Clear Register
type registerIcrType uint32

const (
	RegisterIcrFieldDowncfShift = 6
	RegisterIcrFieldDowncfMask  = 0x40
)

// GetDowncf Direction change to down Clear Flag
func (r *registerIcrType) GetDowncf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldDowncfMask) != 0
}

// SetDowncf Direction change to down Clear Flag
func (r *registerIcrType) SetDowncf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldDowncfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldDowncfMask)
	}
}

const (
	RegisterIcrFieldUpcfShift = 5
	RegisterIcrFieldUpcfMask  = 0x20
)

// GetUpcf Direction change to UP Clear Flag
func (r *registerIcrType) GetUpcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldUpcfMask) != 0
}

// SetUpcf Direction change to UP Clear Flag
func (r *registerIcrType) SetUpcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldUpcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldUpcfMask)
	}
}

const (
	RegisterIcrFieldArrokcfShift = 4
	RegisterIcrFieldArrokcfMask  = 0x10
)

// GetArrokcf Autoreload register update OK Clear Flag
func (r *registerIcrType) GetArrokcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldArrokcfMask) != 0
}

// SetArrokcf Autoreload register update OK Clear Flag
func (r *registerIcrType) SetArrokcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldArrokcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldArrokcfMask)
	}
}

const (
	RegisterIcrFieldCmpokcfShift = 3
	RegisterIcrFieldCmpokcfMask  = 0x8
)

// GetCmpokcf Compare register update OK Clear Flag
func (r *registerIcrType) GetCmpokcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCmpokcfMask) != 0
}

// SetCmpokcf Compare register update OK Clear Flag
func (r *registerIcrType) SetCmpokcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCmpokcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCmpokcfMask)
	}
}

const (
	RegisterIcrFieldExttrigcfShift = 2
	RegisterIcrFieldExttrigcfMask  = 0x4
)

// GetExttrigcf External trigger valid edge Clear Flag
func (r *registerIcrType) GetExttrigcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldExttrigcfMask) != 0
}

// SetExttrigcf External trigger valid edge Clear Flag
func (r *registerIcrType) SetExttrigcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldExttrigcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldExttrigcfMask)
	}
}

const (
	RegisterIcrFieldArrmcfShift = 1
	RegisterIcrFieldArrmcfMask  = 0x2
)

// GetArrmcf Autoreload match Clear Flag
func (r *registerIcrType) GetArrmcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldArrmcfMask) != 0
}

// SetArrmcf Autoreload match Clear Flag
func (r *registerIcrType) SetArrmcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldArrmcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldArrmcfMask)
	}
}

const (
	RegisterIcrFieldCmpmcfShift = 0
	RegisterIcrFieldCmpmcfMask  = 0x1
)

// GetCmpmcf compare match Clear Flag
func (r *registerIcrType) GetCmpmcf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCmpmcfMask) != 0
}

// SetCmpmcf compare match Clear Flag
func (r *registerIcrType) SetCmpmcf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCmpmcfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCmpmcfMask)
	}
}

// registerIerType Interrupt Enable Register
type registerIerType uint32

const (
	RegisterIerFieldDownieShift = 6
	RegisterIerFieldDownieMask  = 0x40
)

// GetDownie Direction change to down Interrupt Enable
func (r *registerIerType) GetDownie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldDownieMask) != 0
}

// SetDownie Direction change to down Interrupt Enable
func (r *registerIerType) SetDownie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldDownieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldDownieMask)
	}
}

const (
	RegisterIerFieldUpieShift = 5
	RegisterIerFieldUpieMask  = 0x20
)

// GetUpie Direction change to UP Interrupt Enable
func (r *registerIerType) GetUpie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldUpieMask) != 0
}

// SetUpie Direction change to UP Interrupt Enable
func (r *registerIerType) SetUpie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldUpieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldUpieMask)
	}
}

const (
	RegisterIerFieldArrokieShift = 4
	RegisterIerFieldArrokieMask  = 0x10
)

// GetArrokie Autoreload register update OK Interrupt Enable
func (r *registerIerType) GetArrokie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldArrokieMask) != 0
}

// SetArrokie Autoreload register update OK Interrupt Enable
func (r *registerIerType) SetArrokie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldArrokieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldArrokieMask)
	}
}

const (
	RegisterIerFieldCmpokieShift = 3
	RegisterIerFieldCmpokieMask  = 0x8
)

// GetCmpokie Compare register update OK Interrupt Enable
func (r *registerIerType) GetCmpokie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldCmpokieMask) != 0
}

// SetCmpokie Compare register update OK Interrupt Enable
func (r *registerIerType) SetCmpokie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldCmpokieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldCmpokieMask)
	}
}

const (
	RegisterIerFieldExttrigieShift = 2
	RegisterIerFieldExttrigieMask  = 0x4
)

// GetExttrigie External trigger valid edge Interrupt Enable
func (r *registerIerType) GetExttrigie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldExttrigieMask) != 0
}

// SetExttrigie External trigger valid edge Interrupt Enable
func (r *registerIerType) SetExttrigie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldExttrigieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldExttrigieMask)
	}
}

const (
	RegisterIerFieldArrmieShift = 1
	RegisterIerFieldArrmieMask  = 0x2
)

// GetArrmie Autoreload match Interrupt Enable
func (r *registerIerType) GetArrmie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldArrmieMask) != 0
}

// SetArrmie Autoreload match Interrupt Enable
func (r *registerIerType) SetArrmie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldArrmieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldArrmieMask)
	}
}

const (
	RegisterIerFieldCmpmieShift = 0
	RegisterIerFieldCmpmieMask  = 0x1
)

// GetCmpmie Compare match Interrupt Enable
func (r *registerIerType) GetCmpmie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldCmpmieMask) != 0
}

// SetCmpmie Compare match Interrupt Enable
func (r *registerIerType) SetCmpmie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldCmpmieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldCmpmieMask)
	}
}

// registerCfgrType Configuration Register
type registerCfgrType uint32

const (
	RegisterCfgrFieldEncShift = 24
	RegisterCfgrFieldEncMask  = 0x1000000
)

// GetEnc Encoder mode enable
func (r *registerCfgrType) GetEnc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldEncMask) != 0
}

// SetEnc Encoder mode enable
func (r *registerCfgrType) SetEnc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldEncMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldEncMask)
	}
}

const (
	RegisterCfgrFieldCountmodeShift = 23
	RegisterCfgrFieldCountmodeMask  = 0x800000
)

// GetCountmode counter mode enabled
func (r *registerCfgrType) GetCountmode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCountmodeMask) != 0
}

// SetCountmode counter mode enabled
func (r *registerCfgrType) SetCountmode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldCountmodeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldCountmodeMask)
	}
}

const (
	RegisterCfgrFieldPreloadShift = 22
	RegisterCfgrFieldPreloadMask  = 0x400000
)

// GetPreload Registers update mode
func (r *registerCfgrType) GetPreload() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldPreloadMask) != 0
}

// SetPreload Registers update mode
func (r *registerCfgrType) SetPreload(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldPreloadMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldPreloadMask)
	}
}

const (
	RegisterCfgrFieldWavpolShift = 21
	RegisterCfgrFieldWavpolMask  = 0x200000
)

// GetWavpol Waveform shape polarity
func (r *registerCfgrType) GetWavpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldWavpolMask) != 0
}

// SetWavpol Waveform shape polarity
func (r *registerCfgrType) SetWavpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldWavpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldWavpolMask)
	}
}

const (
	RegisterCfgrFieldWaveShift = 20
	RegisterCfgrFieldWaveMask  = 0x100000
)

// GetWave Waveform shape
func (r *registerCfgrType) GetWave() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldWaveMask) != 0
}

// SetWave Waveform shape
func (r *registerCfgrType) SetWave(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldWaveMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldWaveMask)
	}
}

const (
	RegisterCfgrFieldTimoutShift = 19
	RegisterCfgrFieldTimoutMask  = 0x80000
)

// GetTimout Timeout enable
func (r *registerCfgrType) GetTimout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldTimoutMask) != 0
}

// SetTimout Timeout enable
func (r *registerCfgrType) SetTimout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldTimoutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldTimoutMask)
	}
}

const (
	RegisterCfgrFieldTrigenShift = 17
	RegisterCfgrFieldTrigenMask  = 0x60000
)

// GetTrigen Trigger enable and polarity
func (r *registerCfgrType) GetTrigen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldTrigenMask) >> RegisterCfgrFieldTrigenShift)
}

// SetTrigen Trigger enable and polarity
func (r *registerCfgrType) SetTrigen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldTrigenMask)|(uint32(value)<<RegisterCfgrFieldTrigenShift))
}

const (
	RegisterCfgrFieldTrigselShift = 13
	RegisterCfgrFieldTrigselMask  = 0xe000
)

// GetTrigsel Trigger selector
func (r *registerCfgrType) GetTrigsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldTrigselMask) >> RegisterCfgrFieldTrigselShift)
}

// SetTrigsel Trigger selector
func (r *registerCfgrType) SetTrigsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldTrigselMask)|(uint32(value)<<RegisterCfgrFieldTrigselShift))
}

const (
	RegisterCfgrFieldPrescShift = 9
	RegisterCfgrFieldPrescMask  = 0xe00
)

// GetPresc Clock prescaler
func (r *registerCfgrType) GetPresc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldPrescMask) >> RegisterCfgrFieldPrescShift)
}

// SetPresc Clock prescaler
func (r *registerCfgrType) SetPresc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldPrescMask)|(uint32(value)<<RegisterCfgrFieldPrescShift))
}

const (
	RegisterCfgrFieldTrgfltShift = 6
	RegisterCfgrFieldTrgfltMask  = 0xc0
)

// GetTrgflt Configurable digital filter for trigger
func (r *registerCfgrType) GetTrgflt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldTrgfltMask) >> RegisterCfgrFieldTrgfltShift)
}

// SetTrgflt Configurable digital filter for trigger
func (r *registerCfgrType) SetTrgflt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldTrgfltMask)|(uint32(value)<<RegisterCfgrFieldTrgfltShift))
}

const (
	RegisterCfgrFieldCkfltShift = 3
	RegisterCfgrFieldCkfltMask  = 0x18
)

// GetCkflt Configurable digital filter for external clock
func (r *registerCfgrType) GetCkflt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCkfltMask) >> RegisterCfgrFieldCkfltShift)
}

// SetCkflt Configurable digital filter for external clock
func (r *registerCfgrType) SetCkflt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldCkfltMask)|(uint32(value)<<RegisterCfgrFieldCkfltShift))
}

const (
	RegisterCfgrFieldCkpolShift = 1
	RegisterCfgrFieldCkpolMask  = 0x6
)

// GetCkpol Clock Polarity
func (r *registerCfgrType) GetCkpol() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCkpolMask) >> RegisterCfgrFieldCkpolShift)
}

// SetCkpol Clock Polarity
func (r *registerCfgrType) SetCkpol(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldCkpolMask)|(uint32(value)<<RegisterCfgrFieldCkpolShift))
}

const (
	RegisterCfgrFieldCkselShift = 0
	RegisterCfgrFieldCkselMask  = 0x1
)

// GetCksel Clock selector
func (r *registerCfgrType) GetCksel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldCkselMask) != 0
}

// SetCksel Clock selector
func (r *registerCfgrType) SetCksel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldCkselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldCkselMask)
	}
}

// registerCrType Control Register
type registerCrType uint32

const (
	RegisterCrFieldEnableShift = 0
	RegisterCrFieldEnableMask  = 0x1
)

// GetEnable LPTIM Enable
func (r *registerCrType) GetEnable() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEnableMask) != 0
}

// SetEnable LPTIM Enable
func (r *registerCrType) SetEnable(value bool) {
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
func (r *registerCrType) GetSngstrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldSngstrtMask) != 0
}

// SetSngstrt LPTIM start in single mode
func (r *registerCrType) SetSngstrt(value bool) {
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
func (r *registerCrType) GetCntstrt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCntstrtMask) != 0
}

// SetCntstrt Timer start in continuous mode
func (r *registerCrType) SetCntstrt(value bool) {
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
func (r *registerCrType) GetCountrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCountrstMask) != 0
}

// SetCountrst Counter reset
func (r *registerCrType) SetCountrst(value bool) {
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
func (r *registerCrType) GetRstare() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldRstareMask) != 0
}

// SetRstare Reset after read enable
func (r *registerCrType) SetRstare(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldRstareMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldRstareMask)
	}
}

// registerCmpType Compare Register
type registerCmpType uint32

const (
	RegisterCmpFieldCmpShift = 0
	RegisterCmpFieldCmpMask  = 0xffff
)

// GetCmp Compare value
func (r *registerCmpType) GetCmp() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCmpFieldCmpMask) >> RegisterCmpFieldCmpShift)
}

// SetCmp Compare value
func (r *registerCmpType) SetCmp(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmpFieldCmpMask)|(uint32(value)<<RegisterCmpFieldCmpShift))
}

// registerArrType Autoreload Register
type registerArrType uint32

const (
	RegisterArrFieldArrShift = 0
	RegisterArrFieldArrMask  = 0xffff
)

// GetArr Auto reload value
func (r *registerArrType) GetArr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterArrFieldArrMask) >> RegisterArrFieldArrShift)
}

// SetArr Auto reload value
func (r *registerArrType) SetArr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArrFieldArrMask)|(uint32(value)<<RegisterArrFieldArrShift))
}

// registerCntType Counter Register
type registerCntType uint32

const (
	RegisterCntFieldCntShift = 0
	RegisterCntFieldCntMask  = 0xffff
)

// GetCnt Counter value
func (r *registerCntType) GetCnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldCntMask) >> RegisterCntFieldCntShift)
}

// SetCnt Counter value
func (r *registerCntType) SetCnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldCntMask)|(uint32(value)<<RegisterCntFieldCntShift))
}

// registerCfgr2Type LPTIM configuration register 2
type registerCfgr2Type uint32

const (
	RegisterCfgr2FieldIn1selShift = 0
	RegisterCfgr2FieldIn1selMask  = 0x3
)

// GetIn1sel LPTIM Input 1 selection
func (r *registerCfgr2Type) GetIn1sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgr2FieldIn1selMask) >> RegisterCfgr2FieldIn1selShift)
}

// SetIn1sel LPTIM Input 1 selection
func (r *registerCfgr2Type) SetIn1sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgr2FieldIn1selMask)|(uint32(value)<<RegisterCfgr2FieldIn1selShift))
}
