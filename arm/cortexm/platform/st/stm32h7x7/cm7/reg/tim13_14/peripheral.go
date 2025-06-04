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
	Cr    registerCrType
	_     [8]byte
	Dier  registerDierType
	Sr    registerSrType
	Egr   registerEgrType
	Ccmr  registerCcmrType
	_     [4]byte
	Ccer  registerCcerType
	Cnt   registerCntType
	Psc   registerPscType
	Arr   registerArrType
	_     [4]byte
	Ccr   registerCcrType
	_     [48]byte
	Tisel registerTiselType
}

// registerCrType control register 1
type registerCrType uint32

const (
	RegisterCrFieldCenShift = 0
	RegisterCrFieldCenMask  = 0x1
)

// GetCen Counter enable
func (r *registerCrType) GetCen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCenMask) != 0
}

// SetCen Counter enable
func (r *registerCrType) SetCen(value bool) {
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
func (r *registerCrType) GetUdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldUdisMask) != 0
}

// SetUdis Update disable
func (r *registerCrType) SetUdis(value bool) {
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
func (r *registerCrType) GetUrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldUrsMask) != 0
}

// SetUrs Update request source
func (r *registerCrType) SetUrs(value bool) {
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
func (r *registerCrType) GetOpm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOpmMask) != 0
}

// SetOpm One-pulse mode
func (r *registerCrType) SetOpm(value bool) {
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
func (r *registerCrType) GetArpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldArpeMask) != 0
}

// SetArpe Auto-reload preload enable
func (r *registerCrType) SetArpe(value bool) {
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
func (r *registerCrType) GetCkd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCkdMask) >> RegisterCrFieldCkdShift)
}

// SetCkd Clock division
func (r *registerCrType) SetCkd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCkdMask)|(uint32(value)<<RegisterCrFieldCkdShift))
}

const (
	RegisterCrFieldUifremapShift = 11
	RegisterCrFieldUifremapMask  = 0x800
)

// GetUifremap UIF status bit remapping
func (r *registerCrType) GetUifremap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldUifremapMask) != 0
}

// SetUifremap UIF status bit remapping
func (r *registerCrType) SetUifremap(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldUifremapMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldUifremapMask)
	}
}

// registerDierType DMA/Interrupt enable register
type registerDierType uint32

const (
	RegisterDierFieldUieShift = 0
	RegisterDierFieldUieMask  = 0x1
)

// GetUie Update interrupt enable
func (r *registerDierType) GetUie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldUieMask) != 0
}

// SetUie Update interrupt enable
func (r *registerDierType) SetUie(value bool) {
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
func (r *registerDierType) GetCcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDierFieldCcieMask) != 0
}

// SetCcie Capture/Compare interrupt enable
func (r *registerDierType) SetCcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDierFieldCcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDierFieldCcieMask)
	}
}

// registerSrType status register
type registerSrType uint32

const (
	RegisterSrFieldUifShift = 0
	RegisterSrFieldUifMask  = 0x1
)

// GetUif Update interrupt flag
func (r *registerSrType) GetUif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldUifMask) != 0
}

// SetUif Update interrupt flag
func (r *registerSrType) SetUif(value bool) {
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
func (r *registerSrType) GetCcif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCcifMask) != 0
}

// SetCcif Capture/compare interrupt flag
func (r *registerSrType) SetCcif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCcifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCcifMask)
	}
}

// registerEgrType event generation register
type registerEgrType uint32

const (
	RegisterEgrFieldUgShift = 0
	RegisterEgrFieldUgMask  = 0x1
)

// GetUg Update generation
func (r *registerEgrType) GetUg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldUgMask) != 0
}

// SetUg Update generation
func (r *registerEgrType) SetUg(value bool) {
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
func (r *registerEgrType) GetCc1g() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterEgrFieldCc1gMask) != 0
}

// SetCc1g Capture/compare generation
func (r *registerEgrType) SetCc1g(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterEgrFieldCc1gMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterEgrFieldCc1gMask)
	}
}

// registerCcmrType capture/compare mode register 1 (output mode)
type registerCcmrType uint32

const (
	RegisterCcmrFieldCcsShift = 0
	RegisterCcmrFieldCcsMask  = 0x3
)

// GetCcs Output CCS
func (r *registerCcmrType) GetCcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldCcsMask) >> RegisterCcmrFieldCcsShift)
}

// SetCcs Output CCS
func (r *registerCcmrType) SetCcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldCcsMask)|(uint32(value)<<RegisterCcmrFieldCcsShift))
}

const (
	RegisterCcmrFieldOcfeShift = 2
	RegisterCcmrFieldOcfeMask  = 0x4
)

// GetOcfe Output OCFE
func (r *registerCcmrType) GetOcfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldOcfeMask) != 0
}

// SetOcfe Output OCFE
func (r *registerCcmrType) SetOcfe(value bool) {
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
func (r *registerCcmrType) GetIcpcs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldIcpcsMask) >> RegisterCcmrFieldIcpcsShift)
}

// SetIcpcs Input capture prescaler
func (r *registerCcmrType) SetIcpcs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldIcpcsMask)|(uint32(value)<<RegisterCcmrFieldIcpcsShift))
}

const (
	RegisterCcmrFieldOcpeShift = 3
	RegisterCcmrFieldOcpeMask  = 0x8
)

// GetOcpe Output OCPE
func (r *registerCcmrType) GetOcpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldOcpeMask) != 0
}

// SetOcpe Output OCPE
func (r *registerCcmrType) SetOcpe(value bool) {
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
func (r *registerCcmrType) GetOcm() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldOcmMask) >> RegisterCcmrFieldOcmShift)
}

// SetOcm Output OCM
func (r *registerCcmrType) SetOcm(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldOcmMask)|(uint32(value)<<RegisterCcmrFieldOcmShift))
}

const (
	RegisterCcmrFieldIcfShift = 4
	RegisterCcmrFieldIcfMask  = 0xf0
)

// GetIcf Input capture filter
func (r *registerCcmrType) GetIcf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldIcfMask) >> RegisterCcmrFieldIcfShift)
}

// SetIcf Input capture filter
func (r *registerCcmrType) SetIcf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldIcfMask)|(uint32(value)<<RegisterCcmrFieldIcfShift))
}

const (
	RegisterCcmrFieldOcm3Shift = 16
	RegisterCcmrFieldOcm3Mask  = 0x10000
)

// GetOcm3 Output Compare mode - bit 3
func (r *registerCcmrType) GetOcm3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcmrFieldOcm3Mask) != 0
}

// SetOcm3 Output Compare mode - bit 3
func (r *registerCcmrType) SetOcm3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcmrFieldOcm3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcmrFieldOcm3Mask)
	}
}

// registerCcerType capture/compare enable register
type registerCcerType uint32

const (
	RegisterCcerFieldCceShift = 0
	RegisterCcerFieldCceMask  = 0x1
)

// GetCce Capture/Compare output enable
func (r *registerCcerType) GetCce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCceMask) != 0
}

// SetCce Capture/Compare output enable
func (r *registerCcerType) SetCce(value bool) {
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
func (r *registerCcerType) GetCcp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCcpMask) != 0
}

// SetCcp Capture/Compare output Polarity
func (r *registerCcerType) SetCcp(value bool) {
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
func (r *registerCcerType) GetCcnp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCcerFieldCcnpMask) != 0
}

// SetCcnp Capture/Compare output Polarity
func (r *registerCcerType) SetCcnp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCcerFieldCcnpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCcerFieldCcnpMask)
	}
}

// registerCntType counter
type registerCntType uint32

const (
	RegisterCntFieldCntShift = 0
	RegisterCntFieldCntMask  = 0xffff
)

// GetCnt low counter value
func (r *registerCntType) GetCnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldCntMask) >> RegisterCntFieldCntShift)
}

// SetCnt low counter value
func (r *registerCntType) SetCnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldCntMask)|(uint32(value)<<RegisterCntFieldCntShift))
}

const (
	RegisterCntFieldUifcpyShift = 31
	RegisterCntFieldUifcpyMask  = 0x80000000
)

// GetUifcpy UIF Copy - This bit is a read-only copy of the UIF bit in the TIMx_ISR register.
func (r *registerCntType) GetUifcpy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCntFieldUifcpyMask) != 0
}

// SetUifcpy UIF Copy - This bit is a read-only copy of the UIF bit in the TIMx_ISR register.
func (r *registerCntType) SetUifcpy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCntFieldUifcpyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCntFieldUifcpyMask)
	}
}

// registerPscType prescaler
type registerPscType uint32

const (
	RegisterPscFieldPscShift = 0
	RegisterPscFieldPscMask  = 0xffff
)

// GetPsc Prescaler value
func (r *registerPscType) GetPsc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterPscFieldPscMask) >> RegisterPscFieldPscShift)
}

// SetPsc Prescaler value
func (r *registerPscType) SetPsc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPscFieldPscMask)|(uint32(value)<<RegisterPscFieldPscShift))
}

// registerArrType auto-reload register
type registerArrType uint32

const (
	RegisterArrFieldArrShift = 0
	RegisterArrFieldArrMask  = 0xffff
)

// GetArr Auto-reload value
func (r *registerArrType) GetArr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterArrFieldArrMask) >> RegisterArrFieldArrShift)
}

// SetArr Auto-reload value
func (r *registerArrType) SetArr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArrFieldArrMask)|(uint32(value)<<RegisterArrFieldArrShift))
}

// registerCcrType capture/compare register 1
type registerCcrType uint32

const (
	RegisterCcrFieldCcrShift = 0
	RegisterCcrFieldCcrMask  = 0xffff
)

// GetCcr Capture/Compare value
func (r *registerCcrType) GetCcr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCcrFieldCcrMask) >> RegisterCcrFieldCcrShift)
}

// SetCcr Capture/Compare value
func (r *registerCcrType) SetCcr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCcrFieldCcrMask)|(uint32(value)<<RegisterCcrFieldCcrShift))
}

// registerTiselType timer input selection register
type registerTiselType uint32

const (
	RegisterTiselFieldTi1selShift = 0
	RegisterTiselFieldTi1selMask  = 0xf
)

// GetTi1sel TI1[0] to TI1[15] input selection
func (r *registerTiselType) GetTi1sel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTiselFieldTi1selMask) >> RegisterTiselFieldTi1selShift)
}

// SetTi1sel TI1[0] to TI1[15] input selection
func (r *registerTiselType) SetTi1sel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTiselFieldTi1selMask)|(uint32(value)<<RegisterTiselFieldTi1selShift))
}
