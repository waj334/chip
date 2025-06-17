//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package tim13_14

import (
	"unsafe"
	"volatile"
)

var (
	Tim13 = (*_tim13_14)(unsafe.Pointer(uintptr(0x40001c00)))
	Tim14 = (*_tim13_14)(unsafe.Pointer(uintptr(0x40002000)))

	Instances = [2]*_tim13_14{
		Tim13,
		Tim14,
	}
)

type _tim13_14 struct {
	Cr    RegisterCrType
	_     [8]byte
	Dier  RegisterDierType
	Sr    RegisterSrType
	Egr   RegisterEgrType
	Ccmr  RegisterCcmrType
	_     [4]byte
	Ccer  RegisterCcerType
	Cnt   RegisterCntType
	Psc   RegisterPscType
	Arr   RegisterArrType
	_     [4]byte
	Ccr   RegisterCcrType
	_     [48]byte
	Tisel RegisterTiselType
}

// RegisterCrType control register 1
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
	RegisterCrFieldCenShift = 0
	RegisterCrFieldCenMask  = 0x1
)

// GetCen Counter enable
func (r *RegisterCrType) GetCen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCenMask) != 0
}

// SetCen Counter enable
func (r *RegisterCrType) SetCen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCenMask)
	}
}

const (
	RegisterCrFieldUdisShift = 1
	RegisterCrFieldUdisMask  = 0x2
)

// GetUdis Update disable
func (r *RegisterCrType) GetUdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldUdisMask) != 0
}

// SetUdis Update disable
func (r *RegisterCrType) SetUdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldUdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldUdisMask)
	}
}

const (
	RegisterCrFieldUrsShift = 2
	RegisterCrFieldUrsMask  = 0x4
)

// GetUrs Update request source
func (r *RegisterCrType) GetUrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldUrsMask) != 0
}

// SetUrs Update request source
func (r *RegisterCrType) SetUrs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldUrsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldUrsMask)
	}
}

const (
	RegisterCrFieldOpmShift = 3
	RegisterCrFieldOpmMask  = 0x8
)

// GetOpm One-pulse mode
func (r *RegisterCrType) GetOpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOpmMask) != 0
}

// SetOpm One-pulse mode
func (r *RegisterCrType) SetOpm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldOpmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldOpmMask)
	}
}

const (
	RegisterCrFieldArpeShift = 7
	RegisterCrFieldArpeMask  = 0x80
)

// GetArpe Auto-reload preload enable
func (r *RegisterCrType) GetArpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldArpeMask) != 0
}

// SetArpe Auto-reload preload enable
func (r *RegisterCrType) SetArpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldArpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldArpeMask)
	}
}

const (
	RegisterCrFieldCkdShift = 8
	RegisterCrFieldCkdMask  = 0x300
)

// GetCkd Clock division
func (r *RegisterCrType) GetCkd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCkdMask) >> RegisterCrFieldCkdShift)
}

// SetCkd Clock division
func (r *RegisterCrType) SetCkd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCkdMask)|(uint32(value)<<RegisterCrFieldCkdShift))
}

const (
	RegisterCrFieldUifremapShift = 11
	RegisterCrFieldUifremapMask  = 0x800
)

// GetUifremap UIF status bit remapping
func (r *RegisterCrType) GetUifremap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldUifremapMask) != 0
}

// SetUifremap UIF status bit remapping
func (r *RegisterCrType) SetUifremap(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldUifremapMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldUifremapMask)
	}
}

// RegisterDierType DMA/Interrupt enable register
type RegisterDierType uint32

func (r *RegisterDierType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDierType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDierType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDierType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDierType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDierFieldUieShift = 0
	RegisterDierFieldUieMask  = 0x1
)

// GetUie Update interrupt enable
func (r *RegisterDierType) GetUie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldUieMask) != 0
}

// SetUie Update interrupt enable
func (r *RegisterDierType) SetUie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldUieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldUieMask)
	}
}

const (
	RegisterDierFieldCcieShift = 1
	RegisterDierFieldCcieMask  = 0x2
)

// GetCcie Capture/Compare interrupt enable
func (r *RegisterDierType) GetCcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCcieMask) != 0
}

// SetCcie Capture/Compare interrupt enable
func (r *RegisterDierType) SetCcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldCcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldCcieMask)
	}
}

// RegisterSrType status register
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
	RegisterSrFieldUifShift = 0
	RegisterSrFieldUifMask  = 0x1
)

// GetUif Update interrupt flag
func (r *RegisterSrType) GetUif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldUifMask) != 0
}

// SetUif Update interrupt flag
func (r *RegisterSrType) SetUif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldUifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldUifMask)
	}
}

const (
	RegisterSrFieldCcifShift = 1
	RegisterSrFieldCcifMask  = 0x2
)

// GetCcif Capture/compare interrupt flag
func (r *RegisterSrType) GetCcif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCcifMask) != 0
}

// SetCcif Capture/compare interrupt flag
func (r *RegisterSrType) SetCcif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCcifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCcifMask)
	}
}

// RegisterEgrType event generation register
type RegisterEgrType uint32

func (r *RegisterEgrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterEgrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterEgrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterEgrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterEgrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterEgrFieldUgShift = 0
	RegisterEgrFieldUgMask  = 0x1
)

// GetUg Update generation
func (r *RegisterEgrType) GetUg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldUgMask) != 0
}

// SetUg Update generation
func (r *RegisterEgrType) SetUg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldUgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldUgMask)
	}
}

const (
	RegisterEgrFieldCc1gShift = 1
	RegisterEgrFieldCc1gMask  = 0x2
)

// GetCc1g Capture/compare generation
func (r *RegisterEgrType) GetCc1g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldCc1gMask) != 0
}

// SetCc1g Capture/compare generation
func (r *RegisterEgrType) SetCc1g(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldCc1gMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldCc1gMask)
	}
}

// RegisterCcmrType capture/compare mode register 1 (output mode)
type RegisterCcmrType uint32

func (r *RegisterCcmrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcmrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcmrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcmrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcmrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcmrFieldCcsShift = 0
	RegisterCcmrFieldCcsMask  = 0x3
)

// GetCcs Output CCS
func (r *RegisterCcmrType) GetCcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldCcsMask) >> RegisterCcmrFieldCcsShift)
}

// SetCcs Output CCS
func (r *RegisterCcmrType) SetCcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldCcsMask)|(uint32(value)<<RegisterCcmrFieldCcsShift))
}

const (
	RegisterCcmrFieldOcfeShift = 2
	RegisterCcmrFieldOcfeMask  = 0x4
)

// GetOcfe Output OCFE
func (r *RegisterCcmrType) GetOcfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldOcfeMask) != 0
}

// SetOcfe Output OCFE
func (r *RegisterCcmrType) SetOcfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmrFieldOcfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldOcfeMask)
	}
}

const (
	RegisterCcmrFieldIcpcsShift = 2
	RegisterCcmrFieldIcpcsMask  = 0xc
)

// GetIcpcs Input capture prescaler
func (r *RegisterCcmrType) GetIcpcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldIcpcsMask) >> RegisterCcmrFieldIcpcsShift)
}

// SetIcpcs Input capture prescaler
func (r *RegisterCcmrType) SetIcpcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldIcpcsMask)|(uint32(value)<<RegisterCcmrFieldIcpcsShift))
}

const (
	RegisterCcmrFieldOcpeShift = 3
	RegisterCcmrFieldOcpeMask  = 0x8
)

// GetOcpe Output OCPE
func (r *RegisterCcmrType) GetOcpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldOcpeMask) != 0
}

// SetOcpe Output OCPE
func (r *RegisterCcmrType) SetOcpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmrFieldOcpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldOcpeMask)
	}
}

const (
	RegisterCcmrFieldOcmShift = 4
	RegisterCcmrFieldOcmMask  = 0x70
)

// GetOcm Output OCM
func (r *RegisterCcmrType) GetOcm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldOcmMask) >> RegisterCcmrFieldOcmShift)
}

// SetOcm Output OCM
func (r *RegisterCcmrType) SetOcm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldOcmMask)|(uint32(value)<<RegisterCcmrFieldOcmShift))
}

const (
	RegisterCcmrFieldIcfShift = 4
	RegisterCcmrFieldIcfMask  = 0xf0
)

// GetIcf Input capture filter
func (r *RegisterCcmrType) GetIcf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldIcfMask) >> RegisterCcmrFieldIcfShift)
}

// SetIcf Input capture filter
func (r *RegisterCcmrType) SetIcf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldIcfMask)|(uint32(value)<<RegisterCcmrFieldIcfShift))
}

const (
	RegisterCcmrFieldOcm3Shift = 16
	RegisterCcmrFieldOcm3Mask  = 0x10000
)

// GetOcm3 Output Compare mode - bit 3
func (r *RegisterCcmrType) GetOcm3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldOcm3Mask) != 0
}

// SetOcm3 Output Compare mode - bit 3
func (r *RegisterCcmrType) SetOcm3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmrFieldOcm3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldOcm3Mask)
	}
}

// RegisterCcerType capture/compare enable register
type RegisterCcerType uint32

func (r *RegisterCcerType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcerType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcerType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcerType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcerType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcerFieldCceShift = 0
	RegisterCcerFieldCceMask  = 0x1
)

// GetCce Capture/Compare output enable
func (r *RegisterCcerType) GetCce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCceMask) != 0
}

// SetCce Capture/Compare output enable
func (r *RegisterCcerType) SetCce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCceMask)
	}
}

const (
	RegisterCcerFieldCcpShift = 1
	RegisterCcerFieldCcpMask  = 0x2
)

// GetCcp Capture/Compare output Polarity
func (r *RegisterCcerType) GetCcp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCcpMask) != 0
}

// SetCcp Capture/Compare output Polarity
func (r *RegisterCcerType) SetCcp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCcpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCcpMask)
	}
}

const (
	RegisterCcerFieldCcnpShift = 3
	RegisterCcerFieldCcnpMask  = 0x8
)

// GetCcnp Capture/Compare output Polarity
func (r *RegisterCcerType) GetCcnp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCcnpMask) != 0
}

// SetCcnp Capture/Compare output Polarity
func (r *RegisterCcerType) SetCcnp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCcnpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCcnpMask)
	}
}

// RegisterCntType counter
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

// GetCnt low counter value
func (r *RegisterCntType) GetCnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldCntMask) >> RegisterCntFieldCntShift)
}

// SetCnt low counter value
func (r *RegisterCntType) SetCnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldCntMask)|(uint32(value)<<RegisterCntFieldCntShift))
}

const (
	RegisterCntFieldUifcpyShift = 31
	RegisterCntFieldUifcpyMask  = 0x80000000
)

// GetUifcpy UIF Copy - This bit is a read-only copy of the UIF bit in the TIMx_ISR register.
func (r *RegisterCntType) GetUifcpy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldUifcpyMask) != 0
}

// SetUifcpy UIF Copy - This bit is a read-only copy of the UIF bit in the TIMx_ISR register.
func (r *RegisterCntType) SetUifcpy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCntFieldUifcpyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldUifcpyMask)
	}
}

// RegisterPscType prescaler
type RegisterPscType uint32

func (r *RegisterPscType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterPscType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterPscType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterPscType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterPscType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterPscFieldPscShift = 0
	RegisterPscFieldPscMask  = 0xffff
)

// GetPsc Prescaler value
func (r *RegisterPscType) GetPsc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPscFieldPscMask) >> RegisterPscFieldPscShift)
}

// SetPsc Prescaler value
func (r *RegisterPscType) SetPsc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPscFieldPscMask)|(uint32(value)<<RegisterPscFieldPscShift))
}

// RegisterArrType auto-reload register
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

// GetArr Auto-reload value
func (r *RegisterArrType) GetArr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterArrFieldArrMask) >> RegisterArrFieldArrShift)
}

// SetArr Auto-reload value
func (r *RegisterArrType) SetArr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArrFieldArrMask)|(uint32(value)<<RegisterArrFieldArrShift))
}

// RegisterCcrType capture/compare register 1
type RegisterCcrType uint32

func (r *RegisterCcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterCcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterCcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterCcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterCcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterCcrFieldCcrShift = 0
	RegisterCcrFieldCcrMask  = 0xffff
)

// GetCcr Capture/Compare value
func (r *RegisterCcrType) GetCcr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldCcrMask) >> RegisterCcrFieldCcrShift)
}

// SetCcr Capture/Compare value
func (r *RegisterCcrType) SetCcr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldCcrMask)|(uint32(value)<<RegisterCcrFieldCcrShift))
}

// RegisterTiselType timer input selection register
type RegisterTiselType uint32

func (r *RegisterTiselType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTiselType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTiselType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTiselType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTiselType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTiselFieldTi1selShift = 0
	RegisterTiselFieldTi1selMask  = 0xf
)

// GetTi1sel TI1[0] to TI1[15] input selection
func (r *RegisterTiselType) GetTi1sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi1selMask) >> RegisterTiselFieldTi1selShift)
}

// SetTi1sel TI1[0] to TI1[15] input selection
func (r *RegisterTiselType) SetTi1sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi1selMask)|(uint32(value)<<RegisterTiselFieldTi1selShift))
}
