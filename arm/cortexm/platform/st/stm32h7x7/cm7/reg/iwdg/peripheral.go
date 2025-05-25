//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package iwdg

import (
	"unsafe"
	"volatile"
)

var (
	Iwdg1 = (*_iwdg)(unsafe.Pointer(uintptr(0x58004800)))
	Iwdg2 = (*_iwdg)(unsafe.Pointer(uintptr(0x58004c00)))

	Instances = [2]*_iwdg{
		Iwdg1,
		Iwdg2,
	}
)

type _iwdg struct {
	Kr   registerKrType
	Pr   registerPrType
	Rlr  registerRlrType
	Sr   registerSrType
	Winr registerWinrType
}

// registerKrType Key register
type registerKrType uint32

const (
	RegisterKrFieldKeyShift = 0
	RegisterKrFieldKeyMask  = 0xffff
)

// GetKey Key value (write only, read 0x0000) These bits must be written by software at regular intervals with the key value 0xAAAA, otherwise the watchdog generates a reset when the counter reaches 0. Writing the key value 0x5555 to enable access to the IWDG_PR, IWDG_RLR and IWDG_WINR registers (see Section23.3.6: Register access protection) Writing the key value CCCCh starts the watchdog (except if the hardware watchdog option is selected)
func (r *registerKrType) GetKey() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterKrFieldKeyMask) >> RegisterKrFieldKeyShift)
}

// SetKey Key value (write only, read 0x0000) These bits must be written by software at regular intervals with the key value 0xAAAA, otherwise the watchdog generates a reset when the counter reaches 0. Writing the key value 0x5555 to enable access to the IWDG_PR, IWDG_RLR and IWDG_WINR registers (see Section23.3.6: Register access protection) Writing the key value CCCCh starts the watchdog (except if the hardware watchdog option is selected)
func (r *registerKrType) SetKey(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterKrFieldKeyMask)|(uint32(value)<<RegisterKrFieldKeyShift))
}

// registerPrType Prescaler register
type registerPrType uint32

const (
	RegisterPrFieldPrShift = 0
	RegisterPrFieldPrMask  = 0x7
)

// GetPr Prescaler divider These bits are write access protected see Section23.3.6: Register access protection. They are written by software to select the prescaler divider feeding the counter clock. PVU bit of IWDG_SR must be reset in order to be able to change the prescaler divider. Note: Reading this register returns the prescaler value from the VDD voltage domain. This value may not be up to date/valid if a write operation to this register is ongoing. For this reason the value read from this register is valid only when the PVU bit in the IWDG_SR register is reset.
func (r *registerPrType) GetPr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPrFieldPrMask) >> RegisterPrFieldPrShift)
}

// SetPr Prescaler divider These bits are write access protected see Section23.3.6: Register access protection. They are written by software to select the prescaler divider feeding the counter clock. PVU bit of IWDG_SR must be reset in order to be able to change the prescaler divider. Note: Reading this register returns the prescaler value from the VDD voltage domain. This value may not be up to date/valid if a write operation to this register is ongoing. For this reason the value read from this register is valid only when the PVU bit in the IWDG_SR register is reset.
func (r *registerPrType) SetPr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPrFieldPrMask)|(uint32(value)<<RegisterPrFieldPrShift))
}

// registerRlrType Reload register
type registerRlrType uint32

const (
	RegisterRlrFieldRlShift = 0
	RegisterRlrFieldRlMask  = 0xfff
)

// GetRl Watchdog counter reload value These bits are write access protected see Section23.3.6. They are written by software to define the value to be loaded in the watchdog counter each time the value 0xAAAA is written in the IWDG_KR register. The watchdog counter counts down from this value. The timeout period is a function of this value and the clock prescaler. Refer to the datasheet for the timeout information. The RVU bit in the IWDG_SR register must be reset in order to be able to change the reload value. Note: Reading this register returns the reload value from the VDD voltage domain. This value may not be up to date/valid if a write operation to this register is ongoing on this register. For this reason the value read from this register is valid only when the RVU bit in the IWDG_SR register is reset.
func (r *registerRlrType) GetRl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterRlrFieldRlMask) >> RegisterRlrFieldRlShift)
}

// SetRl Watchdog counter reload value These bits are write access protected see Section23.3.6. They are written by software to define the value to be loaded in the watchdog counter each time the value 0xAAAA is written in the IWDG_KR register. The watchdog counter counts down from this value. The timeout period is a function of this value and the clock prescaler. Refer to the datasheet for the timeout information. The RVU bit in the IWDG_SR register must be reset in order to be able to change the reload value. Note: Reading this register returns the reload value from the VDD voltage domain. This value may not be up to date/valid if a write operation to this register is ongoing on this register. For this reason the value read from this register is valid only when the RVU bit in the IWDG_SR register is reset.
func (r *registerRlrType) SetRl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRlrFieldRlMask)|(uint32(value)<<RegisterRlrFieldRlShift))
}

// registerSrType Status register
type registerSrType uint32

const (
	RegisterSrFieldPvuShift = 0
	RegisterSrFieldPvuMask  = 0x1
)

// GetPvu Watchdog prescaler value update This bit is set by hardware to indicate that an update of the prescaler value is ongoing. It is reset by hardware when the prescaler update operation is completed in the VDD voltage domain (takes up to 5 RC 40 kHz cycles). Prescaler value can be updated only when PVU bit is reset.
func (r *registerSrType) GetPvu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldPvuMask) != 0
}

// SetPvu Watchdog prescaler value update This bit is set by hardware to indicate that an update of the prescaler value is ongoing. It is reset by hardware when the prescaler update operation is completed in the VDD voltage domain (takes up to 5 RC 40 kHz cycles). Prescaler value can be updated only when PVU bit is reset.
func (r *registerSrType) SetPvu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldPvuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldPvuMask)
	}
}

const (
	RegisterSrFieldRvuShift = 1
	RegisterSrFieldRvuMask  = 0x2
)

// GetRvu Watchdog counter reload value update This bit is set by hardware to indicate that an update of the reload value is ongoing. It is reset by hardware when the reload value update operation is completed in the VDD voltage domain (takes up to 5 RC 40 kHz cycles). Reload value can be updated only when RVU bit is reset.
func (r *registerSrType) GetRvu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldRvuMask) != 0
}

// SetRvu Watchdog counter reload value update This bit is set by hardware to indicate that an update of the reload value is ongoing. It is reset by hardware when the reload value update operation is completed in the VDD voltage domain (takes up to 5 RC 40 kHz cycles). Reload value can be updated only when RVU bit is reset.
func (r *registerSrType) SetRvu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldRvuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldRvuMask)
	}
}

const (
	RegisterSrFieldWvuShift = 2
	RegisterSrFieldWvuMask  = 0x4
)

// GetWvu Watchdog counter window value update This bit is set by hardware to indicate that an update of the window value is ongoing. It is reset by hardware when the reload value update operation is completed in the VDD voltage domain (takes up to 5 RC 40 kHz cycles). Window value can be updated only when WVU bit is reset. This bit is generated only if generic window = 1
func (r *registerSrType) GetWvu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldWvuMask) != 0
}

// SetWvu Watchdog counter window value update This bit is set by hardware to indicate that an update of the window value is ongoing. It is reset by hardware when the reload value update operation is completed in the VDD voltage domain (takes up to 5 RC 40 kHz cycles). Window value can be updated only when WVU bit is reset. This bit is generated only if generic window = 1
func (r *registerSrType) SetWvu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldWvuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldWvuMask)
	}
}

// registerWinrType Window register
type registerWinrType uint32

const (
	RegisterWinrFieldWinShift = 0
	RegisterWinrFieldWinMask  = 0xfff
)

// GetWin Watchdog counter window value These bits are write access protected see Section23.3.6. These bits contain the high limit of the window value to be compared to the downcounter. To prevent a reset, the downcounter must be reloaded when its value is lower than the window register value and greater than 0x0 The WVU bit in the IWDG_SR register must be reset in order to be able to change the reload value. Note: Reading this register returns the reload value from the VDD voltage domain. This value may not be valid if a write operation to this register is ongoing. For this reason the value read from this register is valid only when the WVU bit in the IWDG_SR register is reset.
func (r *registerWinrType) GetWin() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterWinrFieldWinMask) >> RegisterWinrFieldWinShift)
}

// SetWin Watchdog counter window value These bits are write access protected see Section23.3.6. These bits contain the high limit of the window value to be compared to the downcounter. To prevent a reset, the downcounter must be reloaded when its value is lower than the window register value and greater than 0x0 The WVU bit in the IWDG_SR register must be reset in order to be able to change the reload value. Note: Reading this register returns the reload value from the VDD voltage domain. This value may not be valid if a write operation to this register is ongoing. For this reason the value read from this register is valid only when the WVU bit in the IWDG_SR register is reset.
func (r *registerWinrType) SetWin(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterWinrFieldWinMask)|(uint32(value)<<RegisterWinrFieldWinShift))
}
