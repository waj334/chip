//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package sdmmc

import (
	"unsafe"
	"volatile"
)

var (
	Sdmmc1 = (*_sdmmc)(unsafe.Pointer(uintptr(0x52007000)))
	Sdmmc2 = (*_sdmmc)(unsafe.Pointer(uintptr(0x48022400)))

	Instances = [2]*_sdmmc{
		Sdmmc1,
		Sdmmc2,
	}
)

type _sdmmc struct {
	Power      registerPowerType
	Clkcr      registerClkcrType
	Argr       registerArgrType
	Cmdr       registerCmdrType
	Respcmdr   registerRespcmdrType
	Resp1r     registerResp1rType
	Resp2r     registerResp2rType
	Resp3r     registerResp3rType
	Resp4r     registerResp4rType
	Dtimer     registerDtimerType
	Dlenr      registerDlenrType
	Dctrl      registerDctrlType
	Dcntr      registerDcntrType
	Star       registerStarType
	Icr        registerIcrType
	Maskr      registerMaskrType
	Acktimer   registerAcktimerType
	_          [12]byte
	Idmactrlr  registerIdmactrlrType
	Idmabsizer registerIdmabsizerType
	Idmabase0r registerIdmabase0rType
	Idmabase1r registerIdmabase1rType
	_          [32]byte
	Fifor      registerFiforType
}

// registerPowerType SDMMC power control register
type registerPowerType uint32

const (
	RegisterPowerFieldPwrctrlShift = 0
	RegisterPowerFieldPwrctrlMask  = 0x3
)

// GetPwrctrl SDMMC state control bits. These bits can only be written when the SDMMC is not in the power-on state (PWRCTRL?11). These bits are used to define the functional state of the SDMMC signals: Any further write will be ignored, PWRCTRL value will keep 11.
func (r *registerPowerType) GetPwrctrl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPowerFieldPwrctrlMask) >> RegisterPowerFieldPwrctrlShift)
}

// SetPwrctrl SDMMC state control bits. These bits can only be written when the SDMMC is not in the power-on state (PWRCTRL?11). These bits are used to define the functional state of the SDMMC signals: Any further write will be ignored, PWRCTRL value will keep 11.
func (r *registerPowerType) SetPwrctrl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPowerFieldPwrctrlMask)|(uint32(value)<<RegisterPowerFieldPwrctrlShift))
}

const (
	RegisterPowerFieldVswitchShift = 2
	RegisterPowerFieldVswitchMask  = 0x4
)

// GetVswitch Voltage switch sequence start. This bit is used to start the timing critical section of the voltage switch sequence:
func (r *registerPowerType) GetVswitch() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPowerFieldVswitchMask) != 0
}

// SetVswitch Voltage switch sequence start. This bit is used to start the timing critical section of the voltage switch sequence:
func (r *registerPowerType) SetVswitch(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPowerFieldVswitchMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPowerFieldVswitchMask)
	}
}

const (
	RegisterPowerFieldVswitchenShift = 3
	RegisterPowerFieldVswitchenMask  = 0x8
)

// GetVswitchen Voltage switch procedure enable. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). This bit is used to stop the SDMMC_CK after the voltage switch command response:
func (r *registerPowerType) GetVswitchen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPowerFieldVswitchenMask) != 0
}

// SetVswitchen Voltage switch procedure enable. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). This bit is used to stop the SDMMC_CK after the voltage switch command response:
func (r *registerPowerType) SetVswitchen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPowerFieldVswitchenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPowerFieldVswitchenMask)
	}
}

const (
	RegisterPowerFieldDirpolShift = 4
	RegisterPowerFieldDirpolMask  = 0x10
)

// GetDirpol Data and command direction signals polarity selection. This bit can only be written when the SDMMC is in the power-off state (PWRCTRL = 00).
func (r *registerPowerType) GetDirpol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPowerFieldDirpolMask) != 0
}

// SetDirpol Data and command direction signals polarity selection. This bit can only be written when the SDMMC is in the power-off state (PWRCTRL = 00).
func (r *registerPowerType) SetDirpol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPowerFieldDirpolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPowerFieldDirpolMask)
	}
}

// registerClkcrType The SDMMC_CLKCR register controls the SDMMC_CK output clock, the SDMMC_RX_CLK receive clock, and the bus width.
type registerClkcrType uint32

const (
	RegisterClkcrFieldClkdivShift = 0
	RegisterClkcrFieldClkdivMask  = 0x3ff
)

// GetClkdiv Clock divide factor This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). This field defines the divide factor between the input clock (SDMMCCLK) and the output clock (SDMMC_CK): SDMMC_CK frequency = SDMMCCLK / [2 * CLKDIV]. 0xx: etc.. xxx: etc..
func (r *registerClkcrType) GetClkdiv() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterClkcrFieldClkdivMask) >> RegisterClkcrFieldClkdivShift)
}

// SetClkdiv Clock divide factor This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). This field defines the divide factor between the input clock (SDMMCCLK) and the output clock (SDMMC_CK): SDMMC_CK frequency = SDMMCCLK / [2 * CLKDIV]. 0xx: etc.. xxx: etc..
func (r *registerClkcrType) SetClkdiv(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClkcrFieldClkdivMask)|(uint32(value)<<RegisterClkcrFieldClkdivShift))
}

const (
	RegisterClkcrFieldPwrsavShift = 12
	RegisterClkcrFieldPwrsavMask  = 0x1000
)

// GetPwrsav Power saving configuration bit This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) For power saving, the SDMMC_CK clock output can be disabled when the bus is idle by setting PWRSAV:
func (r *registerClkcrType) GetPwrsav() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterClkcrFieldPwrsavMask) != 0
}

// SetPwrsav Power saving configuration bit This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) For power saving, the SDMMC_CK clock output can be disabled when the bus is idle by setting PWRSAV:
func (r *registerClkcrType) SetPwrsav(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterClkcrFieldPwrsavMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterClkcrFieldPwrsavMask)
	}
}

const (
	RegisterClkcrFieldWidbusShift = 14
	RegisterClkcrFieldWidbusMask  = 0xc000
)

// GetWidbus Wide bus mode enable bit This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)
func (r *registerClkcrType) GetWidbus() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClkcrFieldWidbusMask) >> RegisterClkcrFieldWidbusShift)
}

// SetWidbus Wide bus mode enable bit This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)
func (r *registerClkcrType) SetWidbus(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClkcrFieldWidbusMask)|(uint32(value)<<RegisterClkcrFieldWidbusShift))
}

const (
	RegisterClkcrFieldNegedgeShift = 16
	RegisterClkcrFieldNegedgeMask  = 0x10000
)

// GetNegedge SDMMC_CK dephasing selection bit for data and Command. This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). When clock division = 1 (CLKDIV = 0), this bit has no effect. Data and Command change on SDMMC_CK falling edge. When clock division &gt;1 (CLKDIV &gt; 0) &amp; DDR = 0: - SDMMC_CK edge occurs on SDMMCCLK rising edge. When clock division >1 (CLKDIV > 0) & DDR = 1: - Data changed on the SDMMCCLK falling edge succeeding a SDMMC_CK edge. - SDMMC_CK edge occurs on SDMMCCLK rising edge. - Data changed on the SDMMC_CK falling edge succeeding a SDMMC_CK edge. - SDMMC_CK edge occurs on SDMMCCLK rising edge.
func (r *registerClkcrType) GetNegedge() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterClkcrFieldNegedgeMask) != 0
}

// SetNegedge SDMMC_CK dephasing selection bit for data and Command. This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). When clock division = 1 (CLKDIV = 0), this bit has no effect. Data and Command change on SDMMC_CK falling edge. When clock division &gt;1 (CLKDIV &gt; 0) &amp; DDR = 0: - SDMMC_CK edge occurs on SDMMCCLK rising edge. When clock division >1 (CLKDIV > 0) & DDR = 1: - Data changed on the SDMMCCLK falling edge succeeding a SDMMC_CK edge. - SDMMC_CK edge occurs on SDMMCCLK rising edge. - Data changed on the SDMMC_CK falling edge succeeding a SDMMC_CK edge. - SDMMC_CK edge occurs on SDMMCCLK rising edge.
func (r *registerClkcrType) SetNegedge(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterClkcrFieldNegedgeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterClkcrFieldNegedgeMask)
	}
}

const (
	RegisterClkcrFieldHwfcenShift = 17
	RegisterClkcrFieldHwfcenMask  = 0x20000
)

// GetHwfcen Hardware flow control enable This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) When Hardware flow control is enabled, the meaning of the TXFIFOE and RXFIFOF flags change, please see SDMMC status register definition in Section56.8.11.
func (r *registerClkcrType) GetHwfcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterClkcrFieldHwfcenMask) != 0
}

// SetHwfcen Hardware flow control enable This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) When Hardware flow control is enabled, the meaning of the TXFIFOE and RXFIFOF flags change, please see SDMMC status register definition in Section56.8.11.
func (r *registerClkcrType) SetHwfcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterClkcrFieldHwfcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterClkcrFieldHwfcenMask)
	}
}

const (
	RegisterClkcrFieldDdrShift = 18
	RegisterClkcrFieldDdrMask  = 0x40000
)

// GetDdr Data rate signaling selection This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) DDR rate shall only be selected with 4-bit or 8-bit wide bus mode. (WIDBUS &gt; 00). DDR = 1 has no effect when WIDBUS = 00 (1-bit wide bus). DDR rate shall only be selected with clock division &gt;1. (CLKDIV &gt; 0)
func (r *registerClkcrType) GetDdr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterClkcrFieldDdrMask) != 0
}

// SetDdr Data rate signaling selection This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) DDR rate shall only be selected with 4-bit or 8-bit wide bus mode. (WIDBUS &gt; 00). DDR = 1 has no effect when WIDBUS = 00 (1-bit wide bus). DDR rate shall only be selected with clock division &gt;1. (CLKDIV &gt; 0)
func (r *registerClkcrType) SetDdr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterClkcrFieldDdrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterClkcrFieldDdrMask)
	}
}

const (
	RegisterClkcrFieldBusspeedShift = 19
	RegisterClkcrFieldBusspeedMask  = 0x80000
)

// GetBusspeed Bus speed mode selection between DS, HS, SDR12, SDR25 and SDR50, DDR50, SDR104. This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)
func (r *registerClkcrType) GetBusspeed() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterClkcrFieldBusspeedMask) != 0
}

// SetBusspeed Bus speed mode selection between DS, HS, SDR12, SDR25 and SDR50, DDR50, SDR104. This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)
func (r *registerClkcrType) SetBusspeed(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterClkcrFieldBusspeedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterClkcrFieldBusspeedMask)
	}
}

const (
	RegisterClkcrFieldSelclkrxShift = 20
	RegisterClkcrFieldSelclkrxMask  = 0x300000
)

// GetSelclkrx Receive clock selection. These bits can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)
func (r *registerClkcrType) GetSelclkrx() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterClkcrFieldSelclkrxMask) >> RegisterClkcrFieldSelclkrxShift)
}

// SetSelclkrx Receive clock selection. These bits can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)
func (r *registerClkcrType) SetSelclkrx(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterClkcrFieldSelclkrxMask)|(uint32(value)<<RegisterClkcrFieldSelclkrxShift))
}

// registerArgrType The SDMMC_ARGR register contains a 32-bit command argument, which is sent to a card as part of a command message.
type registerArgrType uint32

const (
	RegisterArgrFieldCmdargShift = 0
	RegisterArgrFieldCmdargMask  = 0xffffffff
)

// GetCmdarg Command argument. These bits can only be written by firmware when CPSM is disabled (CPSMEN = 0). Command argument sent to a card as part of a command message. If a command contains an argument, it must be loaded into this register before writing a command to the command register.
func (r *registerArgrType) GetCmdarg() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterArgrFieldCmdargMask) >> RegisterArgrFieldCmdargShift)
}

// SetCmdarg Command argument. These bits can only be written by firmware when CPSM is disabled (CPSMEN = 0). Command argument sent to a card as part of a command message. If a command contains an argument, it must be loaded into this register before writing a command to the command register.
func (r *registerArgrType) SetCmdarg(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterArgrFieldCmdargMask)|(uint32(value)<<RegisterArgrFieldCmdargShift))
}

// registerCmdrType The SDMMC_CMDR register contains the command index and command type bits. The command index is sent to a card as part of a command message. The command type bits control the command path state machine (CPSM).
type registerCmdrType uint32

const (
	RegisterCmdrFieldCmdindexShift = 0
	RegisterCmdrFieldCmdindexMask  = 0x3f
)

// GetCmdindex Command index. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). The command index is sent to the card as part of a command message.
func (r *registerCmdrType) GetCmdindex() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCmdrFieldCmdindexMask) >> RegisterCmdrFieldCmdindexShift)
}

// SetCmdindex Command index. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). The command index is sent to the card as part of a command message.
func (r *registerCmdrType) SetCmdindex(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmdrFieldCmdindexMask)|(uint32(value)<<RegisterCmdrFieldCmdindexShift))
}

const (
	RegisterCmdrFieldCmdtransShift = 6
	RegisterCmdrFieldCmdtransMask  = 0x40
)

// GetCmdtrans The CPSM treats the command as a data transfer command, stops the interrupt period, and signals DataEnable to the DPSM This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). If this bit is set, the CPSM issues an end of interrupt period and issues DataEnable signal to the DPSM when the command is sent.
func (r *registerCmdrType) GetCmdtrans() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCmdrFieldCmdtransMask) != 0
}

// SetCmdtrans The CPSM treats the command as a data transfer command, stops the interrupt period, and signals DataEnable to the DPSM This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). If this bit is set, the CPSM issues an end of interrupt period and issues DataEnable signal to the DPSM when the command is sent.
func (r *registerCmdrType) SetCmdtrans(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCmdrFieldCmdtransMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCmdrFieldCmdtransMask)
	}
}

const (
	RegisterCmdrFieldCmdstopShift = 7
	RegisterCmdrFieldCmdstopMask  = 0x80
)

// GetCmdstop The CPSM treats the command as a Stop Transmission command and signals Abort to the DPSM. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). If this bit is set, the CPSM issues the Abort signal to the DPSM when the command is sent.
func (r *registerCmdrType) GetCmdstop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCmdrFieldCmdstopMask) != 0
}

// SetCmdstop The CPSM treats the command as a Stop Transmission command and signals Abort to the DPSM. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). If this bit is set, the CPSM issues the Abort signal to the DPSM when the command is sent.
func (r *registerCmdrType) SetCmdstop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCmdrFieldCmdstopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCmdrFieldCmdstopMask)
	}
}

const (
	RegisterCmdrFieldWaitrespShift = 8
	RegisterCmdrFieldWaitrespMask  = 0x300
)

// GetWaitresp Wait for response bits. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). They are used to configure whether the CPSM is to wait for a response, and if yes, which kind of response.
func (r *registerCmdrType) GetWaitresp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCmdrFieldWaitrespMask) >> RegisterCmdrFieldWaitrespShift)
}

// SetWaitresp Wait for response bits. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). They are used to configure whether the CPSM is to wait for a response, and if yes, which kind of response.
func (r *registerCmdrType) SetWaitresp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCmdrFieldWaitrespMask)|(uint32(value)<<RegisterCmdrFieldWaitrespShift))
}

const (
	RegisterCmdrFieldWaitintShift = 10
	RegisterCmdrFieldWaitintMask  = 0x400
)

// GetWaitint CPSM waits for interrupt request. If this bit is set, the CPSM disables command timeout and waits for an card interrupt request (Response). If this bit is cleared in the CPSM Wait state, will cause the abort of the interrupt mode.
func (r *registerCmdrType) GetWaitint() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCmdrFieldWaitintMask) != 0
}

// SetWaitint CPSM waits for interrupt request. If this bit is set, the CPSM disables command timeout and waits for an card interrupt request (Response). If this bit is cleared in the CPSM Wait state, will cause the abort of the interrupt mode.
func (r *registerCmdrType) SetWaitint(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCmdrFieldWaitintMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCmdrFieldWaitintMask)
	}
}

const (
	RegisterCmdrFieldWaitpendShift = 11
	RegisterCmdrFieldWaitpendMask  = 0x800
)

// GetWaitpend CPSM Waits for end of data transfer (CmdPend internal signal) from DPSM. This bit when set, the CPSM waits for the end of data transfer trigger before it starts sending a command. WAITPEND is only taken into account when DTMODE = MMC stream data transfer, WIDBUS = 1-bit wide bus mode, DPSMACT = 1 and DTDIR = from host to card.
func (r *registerCmdrType) GetWaitpend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCmdrFieldWaitpendMask) != 0
}

// SetWaitpend CPSM Waits for end of data transfer (CmdPend internal signal) from DPSM. This bit when set, the CPSM waits for the end of data transfer trigger before it starts sending a command. WAITPEND is only taken into account when DTMODE = MMC stream data transfer, WIDBUS = 1-bit wide bus mode, DPSMACT = 1 and DTDIR = from host to card.
func (r *registerCmdrType) SetWaitpend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCmdrFieldWaitpendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCmdrFieldWaitpendMask)
	}
}

const (
	RegisterCmdrFieldCpsmenShift = 12
	RegisterCmdrFieldCpsmenMask  = 0x1000
)

// GetCpsmen Command path state machine (CPSM) Enable bit This bit is written 1 by firmware, and cleared by hardware when the CPSM enters the Idle state. If this bit is set, the CPSM is enabled. When DTEN = 1, no command will be transfered nor boot procedure will be started. CPSMEN is cleared to 0.
func (r *registerCmdrType) GetCpsmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCmdrFieldCpsmenMask) != 0
}

// SetCpsmen Command path state machine (CPSM) Enable bit This bit is written 1 by firmware, and cleared by hardware when the CPSM enters the Idle state. If this bit is set, the CPSM is enabled. When DTEN = 1, no command will be transfered nor boot procedure will be started. CPSMEN is cleared to 0.
func (r *registerCmdrType) SetCpsmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCmdrFieldCpsmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCmdrFieldCpsmenMask)
	}
}

const (
	RegisterCmdrFieldDtholdShift = 13
	RegisterCmdrFieldDtholdMask  = 0x2000
)

// GetDthold Hold new data block transmission and reception in the DPSM. If this bit is set, the DPSM will not move from the Wait_S state to the Send state or from the Wait_R state to the Receive state.
func (r *registerCmdrType) GetDthold() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCmdrFieldDtholdMask) != 0
}

// SetDthold Hold new data block transmission and reception in the DPSM. If this bit is set, the DPSM will not move from the Wait_S state to the Send state or from the Wait_R state to the Receive state.
func (r *registerCmdrType) SetDthold(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCmdrFieldDtholdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCmdrFieldDtholdMask)
	}
}

const (
	RegisterCmdrFieldBootmodeShift = 14
	RegisterCmdrFieldBootmodeMask  = 0x4000
)

// GetBootmode Select the boot mode procedure to be used. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0)
func (r *registerCmdrType) GetBootmode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCmdrFieldBootmodeMask) != 0
}

// SetBootmode Select the boot mode procedure to be used. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0)
func (r *registerCmdrType) SetBootmode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCmdrFieldBootmodeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCmdrFieldBootmodeMask)
	}
}

const (
	RegisterCmdrFieldBootenShift = 15
	RegisterCmdrFieldBootenMask  = 0x8000
)

// GetBooten Enable boot mode procedure.
func (r *registerCmdrType) GetBooten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCmdrFieldBootenMask) != 0
}

// SetBooten Enable boot mode procedure.
func (r *registerCmdrType) SetBooten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCmdrFieldBootenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCmdrFieldBootenMask)
	}
}

const (
	RegisterCmdrFieldCmdsuspendShift = 16
	RegisterCmdrFieldCmdsuspendMask  = 0x10000
)

// GetCmdsuspend The CPSM treats the command as a Suspend or Resume command and signals interrupt period start/end. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). CMDSUSPEND = 1 and CMDTRANS = 0 Suspend command, start interrupt period when response bit BS=0. CMDSUSPEND = 1 and CMDTRANS = 1 Resume command with data, end interrupt period when response bit DF=1.
func (r *registerCmdrType) GetCmdsuspend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCmdrFieldCmdsuspendMask) != 0
}

// SetCmdsuspend The CPSM treats the command as a Suspend or Resume command and signals interrupt period start/end. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). CMDSUSPEND = 1 and CMDTRANS = 0 Suspend command, start interrupt period when response bit BS=0. CMDSUSPEND = 1 and CMDTRANS = 1 Resume command with data, end interrupt period when response bit DF=1.
func (r *registerCmdrType) SetCmdsuspend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCmdrFieldCmdsuspendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCmdrFieldCmdsuspendMask)
	}
}

// registerRespcmdrType SDMMC command response register
type registerRespcmdrType uint32

const (
	RegisterRespcmdrFieldRespcmdShift = 0
	RegisterRespcmdrFieldRespcmdMask  = 0x3f
)

// GetRespcmd Response command index
func (r *registerRespcmdrType) GetRespcmd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRespcmdrFieldRespcmdMask) >> RegisterRespcmdrFieldRespcmdShift)
}

// SetRespcmd Response command index
func (r *registerRespcmdrType) SetRespcmd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRespcmdrFieldRespcmdMask)|(uint32(value)<<RegisterRespcmdrFieldRespcmdShift))
}

// registerResp1rType The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.
type registerResp1rType uint32

const (
	RegisterResp1rFieldCardstatus1Shift = 0
	RegisterResp1rFieldCardstatus1Mask  = 0xffffffff
)

// GetCardstatus1 see Table 432
func (r *registerResp1rType) GetCardstatus1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterResp1rFieldCardstatus1Mask) >> RegisterResp1rFieldCardstatus1Shift)
}

// SetCardstatus1 see Table 432
func (r *registerResp1rType) SetCardstatus1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterResp1rFieldCardstatus1Mask)|(uint32(value)<<RegisterResp1rFieldCardstatus1Shift))
}

// registerResp2rType The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.
type registerResp2rType uint32

const (
	RegisterResp2rFieldCardstatus2Shift = 0
	RegisterResp2rFieldCardstatus2Mask  = 0xffffffff
)

// GetCardstatus2 see Table404.
func (r *registerResp2rType) GetCardstatus2() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterResp2rFieldCardstatus2Mask) >> RegisterResp2rFieldCardstatus2Shift)
}

// SetCardstatus2 see Table404.
func (r *registerResp2rType) SetCardstatus2(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterResp2rFieldCardstatus2Mask)|(uint32(value)<<RegisterResp2rFieldCardstatus2Shift))
}

// registerResp3rType The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.
type registerResp3rType uint32

const (
	RegisterResp3rFieldCardstatus3Shift = 0
	RegisterResp3rFieldCardstatus3Mask  = 0xffffffff
)

// GetCardstatus3 see Table404.
func (r *registerResp3rType) GetCardstatus3() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterResp3rFieldCardstatus3Mask) >> RegisterResp3rFieldCardstatus3Shift)
}

// SetCardstatus3 see Table404.
func (r *registerResp3rType) SetCardstatus3(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterResp3rFieldCardstatus3Mask)|(uint32(value)<<RegisterResp3rFieldCardstatus3Shift))
}

// registerResp4rType The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.
type registerResp4rType uint32

const (
	RegisterResp4rFieldCardstatus4Shift = 0
	RegisterResp4rFieldCardstatus4Mask  = 0xffffffff
)

// GetCardstatus4 see Table404.
func (r *registerResp4rType) GetCardstatus4() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterResp4rFieldCardstatus4Mask) >> RegisterResp4rFieldCardstatus4Shift)
}

// SetCardstatus4 see Table404.
func (r *registerResp4rType) SetCardstatus4(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterResp4rFieldCardstatus4Mask)|(uint32(value)<<RegisterResp4rFieldCardstatus4Shift))
}

// registerDtimerType The SDMMC_DTIMER register contains the data timeout period, in card bus clock periods. A counter loads the value from the SDMMC_DTIMER register, and starts decrementing when the data path state machine (DPSM) enters the Wait_R or Busy state. If the timer reaches 0 while the DPSM is in either of these states, the timeout status flag is set.
type registerDtimerType uint32

const (
	RegisterDtimerFieldDatatimeShift = 0
	RegisterDtimerFieldDatatimeMask  = 0xffffffff
)

// GetDatatime Data and R1b busy timeout period This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). Data and R1b busy timeout period expressed in card bus clock periods.
func (r *registerDtimerType) GetDatatime() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDtimerFieldDatatimeMask) >> RegisterDtimerFieldDatatimeShift)
}

// SetDatatime Data and R1b busy timeout period This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). Data and R1b busy timeout period expressed in card bus clock periods.
func (r *registerDtimerType) SetDatatime(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDtimerFieldDatatimeMask)|(uint32(value)<<RegisterDtimerFieldDatatimeShift))
}

// registerDlenrType The SDMMC_DLENR register contains the number of data bytes to be transferred. The value is loaded into the data counter when data transfer starts.
type registerDlenrType uint32

const (
	RegisterDlenrFieldDatalengthShift = 0
	RegisterDlenrFieldDatalengthMask  = 0x1ffffff
)

// GetDatalength Data length value This register can only be written by firmware when DPSM is inactive (DPSMACT = 0). Number of data bytes to be transferred. When DDR = 1 DATALENGTH is truncated to a multiple of 2. (The last odd byte is not transfered) When DATALENGTH = 0 no data will be transfered, when requested by a CPSMEN and CMDTRANS = 1 also no command will be transfered. DTEN and CPSMEN are cleared to 0.
func (r *registerDlenrType) GetDatalength() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDlenrFieldDatalengthMask) >> RegisterDlenrFieldDatalengthShift)
}

// SetDatalength Data length value This register can only be written by firmware when DPSM is inactive (DPSMACT = 0). Number of data bytes to be transferred. When DDR = 1 DATALENGTH is truncated to a multiple of 2. (The last odd byte is not transfered) When DATALENGTH = 0 no data will be transfered, when requested by a CPSMEN and CMDTRANS = 1 also no command will be transfered. DTEN and CPSMEN are cleared to 0.
func (r *registerDlenrType) SetDatalength(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDlenrFieldDatalengthMask)|(uint32(value)<<RegisterDlenrFieldDatalengthShift))
}

// registerDctrlType The SDMMC_DCTRL register control the data path state machine (DPSM).
type registerDctrlType uint32

const (
	RegisterDctrlFieldDtenShift = 0
	RegisterDctrlFieldDtenMask  = 0x1
)

// GetDten Data transfer enable bit This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). This bit is cleared by Hardware when data transfer completes. This bit shall only be used to transfer data when no associated data transfer command is used, i.e. shall not be used with SD or eMMC cards.
func (r *registerDctrlType) GetDten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDctrlFieldDtenMask) != 0
}

// SetDten Data transfer enable bit This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). This bit is cleared by Hardware when data transfer completes. This bit shall only be used to transfer data when no associated data transfer command is used, i.e. shall not be used with SD or eMMC cards.
func (r *registerDctrlType) SetDten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDctrlFieldDtenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDctrlFieldDtenMask)
	}
}

const (
	RegisterDctrlFieldDtdirShift = 1
	RegisterDctrlFieldDtdirMask  = 0x2
)

// GetDtdir Data transfer direction selection This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerDctrlType) GetDtdir() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDctrlFieldDtdirMask) != 0
}

// SetDtdir Data transfer direction selection This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerDctrlType) SetDtdir(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDctrlFieldDtdirMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDctrlFieldDtdirMask)
	}
}

const (
	RegisterDctrlFieldDtmodeShift = 2
	RegisterDctrlFieldDtmodeMask  = 0xc
)

// GetDtmode Data transfer mode selection. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerDctrlType) GetDtmode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDctrlFieldDtmodeMask) >> RegisterDctrlFieldDtmodeShift)
}

// SetDtmode Data transfer mode selection. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerDctrlType) SetDtmode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDctrlFieldDtmodeMask)|(uint32(value)<<RegisterDctrlFieldDtmodeShift))
}

const (
	RegisterDctrlFieldDblocksizeShift = 4
	RegisterDctrlFieldDblocksizeMask  = 0xf0
)

// GetDblocksize Data block size This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). Define the data block length when the block data transfer mode is selected: When DATALENGTH is not a multiple of DBLOCKSIZE, the transfered data is truncated at a multiple of DBLOCKSIZE. (Any remain data will not be transfered.) When DDR = 1, DBLOCKSIZE = 0000 shall not be used. (No data will be transfered)
func (r *registerDctrlType) GetDblocksize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDctrlFieldDblocksizeMask) >> RegisterDctrlFieldDblocksizeShift)
}

// SetDblocksize Data block size This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). Define the data block length when the block data transfer mode is selected: When DATALENGTH is not a multiple of DBLOCKSIZE, the transfered data is truncated at a multiple of DBLOCKSIZE. (Any remain data will not be transfered.) When DDR = 1, DBLOCKSIZE = 0000 shall not be used. (No data will be transfered)
func (r *registerDctrlType) SetDblocksize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDctrlFieldDblocksizeMask)|(uint32(value)<<RegisterDctrlFieldDblocksizeShift))
}

const (
	RegisterDctrlFieldRwstartShift = 8
	RegisterDctrlFieldRwstartMask  = 0x100
)

// GetRwstart Read wait start. If this bit is set, read wait operation starts.
func (r *registerDctrlType) GetRwstart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDctrlFieldRwstartMask) != 0
}

// SetRwstart Read wait start. If this bit is set, read wait operation starts.
func (r *registerDctrlType) SetRwstart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDctrlFieldRwstartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDctrlFieldRwstartMask)
	}
}

const (
	RegisterDctrlFieldRwstopShift = 9
	RegisterDctrlFieldRwstopMask  = 0x200
)

// GetRwstop Read wait stop This bit is written by firmware and auto cleared by hardware when the DPSM moves from the READ_WAIT state to the WAIT_R or IDLE state.
func (r *registerDctrlType) GetRwstop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDctrlFieldRwstopMask) != 0
}

// SetRwstop Read wait stop This bit is written by firmware and auto cleared by hardware when the DPSM moves from the READ_WAIT state to the WAIT_R or IDLE state.
func (r *registerDctrlType) SetRwstop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDctrlFieldRwstopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDctrlFieldRwstopMask)
	}
}

const (
	RegisterDctrlFieldRwmodShift = 10
	RegisterDctrlFieldRwmodMask  = 0x400
)

// GetRwmod Read wait mode. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerDctrlType) GetRwmod() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDctrlFieldRwmodMask) != 0
}

// SetRwmod Read wait mode. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerDctrlType) SetRwmod(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDctrlFieldRwmodMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDctrlFieldRwmodMask)
	}
}

const (
	RegisterDctrlFieldSdioenShift = 11
	RegisterDctrlFieldSdioenMask  = 0x800
)

// GetSdioen SD I/O interrupt enable functions This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). If this bit is set, the DPSM enables the SD I/O card specific interrupt operation.
func (r *registerDctrlType) GetSdioen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDctrlFieldSdioenMask) != 0
}

// SetSdioen SD I/O interrupt enable functions This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). If this bit is set, the DPSM enables the SD I/O card specific interrupt operation.
func (r *registerDctrlType) SetSdioen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDctrlFieldSdioenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDctrlFieldSdioenMask)
	}
}

const (
	RegisterDctrlFieldBootackenShift = 12
	RegisterDctrlFieldBootackenMask  = 0x1000
)

// GetBootacken Enable the reception of the boot acknowledgment. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerDctrlType) GetBootacken() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDctrlFieldBootackenMask) != 0
}

// SetBootacken Enable the reception of the boot acknowledgment. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerDctrlType) SetBootacken(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDctrlFieldBootackenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDctrlFieldBootackenMask)
	}
}

const (
	RegisterDctrlFieldFiforstShift = 13
	RegisterDctrlFieldFiforstMask  = 0x2000
)

// GetFiforst FIFO reset, will flush any remaining data. This bit can only be written by firmware when IDMAEN= 0 and DPSM is active (DPSMACT = 1). This bit will only take effect when a transfer error or transfer hold occurs.
func (r *registerDctrlType) GetFiforst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDctrlFieldFiforstMask) != 0
}

// SetFiforst FIFO reset, will flush any remaining data. This bit can only be written by firmware when IDMAEN= 0 and DPSM is active (DPSMACT = 1). This bit will only take effect when a transfer error or transfer hold occurs.
func (r *registerDctrlType) SetFiforst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDctrlFieldFiforstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDctrlFieldFiforstMask)
	}
}

// registerDcntrType The SDMMC_DCNTR register loads the value from the data length register (see SDMMC_DLENR) when the DPSM moves from the Idle state to the Wait_R or Wait_S state. As data is transferred, the counter decrements the value until it reaches 0. The DPSM then moves to the Idle state and when there has been no error, the data status end flag (DATAEND) is set.
type registerDcntrType uint32

const (
	RegisterDcntrFieldDatacountShift = 0
	RegisterDcntrFieldDatacountMask  = 0x1ffffff
)

// GetDatacount Data count value When read, the number of remaining data bytes to be transferred is returned. Write has no effect.
func (r *registerDcntrType) GetDatacount() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDcntrFieldDatacountMask) >> RegisterDcntrFieldDatacountShift)
}

// SetDatacount Data count value When read, the number of remaining data bytes to be transferred is returned. Write has no effect.
func (r *registerDcntrType) SetDatacount(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDcntrFieldDatacountMask)|(uint32(value)<<RegisterDcntrFieldDatacountShift))
}

// registerStarType The SDMMC_STAR register is a read-only register. It contains two types of flag:Static flags (bits [29,21,11:0]): these bits remain asserted until they are cleared by writing to the SDMMC interrupt Clear register (see SDMMC_ICR)Dynamic flags (bits [20:12]): these bits change state depending on the state of the underlying logic (for example, FIFO full and empty flags are asserted and de-asserted as data while written to the FIFO)
type registerStarType uint32

const (
	RegisterStarFieldCcrcfailShift = 0
	RegisterStarFieldCcrcfailMask  = 0x1
)

// GetCcrcfail Command response received (CRC check failed). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetCcrcfail() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldCcrcfailMask) != 0
}

// SetCcrcfail Command response received (CRC check failed). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetCcrcfail(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldCcrcfailMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldCcrcfailMask)
	}
}

const (
	RegisterStarFieldDcrcfailShift = 1
	RegisterStarFieldDcrcfailMask  = 0x2
)

// GetDcrcfail Data block sent/received (CRC check failed). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetDcrcfail() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldDcrcfailMask) != 0
}

// SetDcrcfail Data block sent/received (CRC check failed). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetDcrcfail(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldDcrcfailMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldDcrcfailMask)
	}
}

const (
	RegisterStarFieldCtimeoutShift = 2
	RegisterStarFieldCtimeoutMask  = 0x4
)

// GetCtimeout Command response timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR. The Command Timeout period has a fixed value of 64 SDMMC_CK clock periods.
func (r *registerStarType) GetCtimeout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldCtimeoutMask) != 0
}

// SetCtimeout Command response timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR. The Command Timeout period has a fixed value of 64 SDMMC_CK clock periods.
func (r *registerStarType) SetCtimeout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldCtimeoutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldCtimeoutMask)
	}
}

const (
	RegisterStarFieldDtimeoutShift = 3
	RegisterStarFieldDtimeoutMask  = 0x8
)

// GetDtimeout Data timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetDtimeout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldDtimeoutMask) != 0
}

// SetDtimeout Data timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetDtimeout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldDtimeoutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldDtimeoutMask)
	}
}

const (
	RegisterStarFieldTxunderrShift = 4
	RegisterStarFieldTxunderrMask  = 0x10
)

// GetTxunderr Transmit FIFO underrun error or IDMA read transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetTxunderr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldTxunderrMask) != 0
}

// SetTxunderr Transmit FIFO underrun error or IDMA read transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetTxunderr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldTxunderrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldTxunderrMask)
	}
}

const (
	RegisterStarFieldRxoverrShift = 5
	RegisterStarFieldRxoverrMask  = 0x20
)

// GetRxoverr Received FIFO overrun error or IDMA write transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetRxoverr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldRxoverrMask) != 0
}

// SetRxoverr Received FIFO overrun error or IDMA write transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetRxoverr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldRxoverrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldRxoverrMask)
	}
}

const (
	RegisterStarFieldCmdrendShift = 6
	RegisterStarFieldCmdrendMask  = 0x40
)

// GetCmdrend Command response received (CRC check passed, or no CRC). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetCmdrend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldCmdrendMask) != 0
}

// SetCmdrend Command response received (CRC check passed, or no CRC). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetCmdrend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldCmdrendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldCmdrendMask)
	}
}

const (
	RegisterStarFieldCmdsentShift = 7
	RegisterStarFieldCmdsentMask  = 0x80
)

// GetCmdsent Command sent (no response required). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetCmdsent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldCmdsentMask) != 0
}

// SetCmdsent Command sent (no response required). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetCmdsent(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldCmdsentMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldCmdsentMask)
	}
}

const (
	RegisterStarFieldDataendShift = 8
	RegisterStarFieldDataendMask  = 0x100
)

// GetDataend Data transfer ended correctly. (data counter, DATACOUNT is zero and no errors occur). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetDataend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldDataendMask) != 0
}

// SetDataend Data transfer ended correctly. (data counter, DATACOUNT is zero and no errors occur). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetDataend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldDataendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldDataendMask)
	}
}

const (
	RegisterStarFieldDholdShift = 9
	RegisterStarFieldDholdMask  = 0x200
)

// GetDhold Data transfer Hold. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetDhold() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldDholdMask) != 0
}

// SetDhold Data transfer Hold. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetDhold(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldDholdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldDholdMask)
	}
}

const (
	RegisterStarFieldDbckendShift = 10
	RegisterStarFieldDbckendMask  = 0x400
)

// GetDbckend Data block sent/received. (CRC check passed) and DPSM moves to the READWAIT state. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetDbckend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldDbckendMask) != 0
}

// SetDbckend Data block sent/received. (CRC check passed) and DPSM moves to the READWAIT state. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetDbckend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldDbckendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldDbckendMask)
	}
}

const (
	RegisterStarFieldDabortShift = 11
	RegisterStarFieldDabortMask  = 0x800
)

// GetDabort Data transfer aborted by CMD12. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetDabort() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldDabortMask) != 0
}

// SetDabort Data transfer aborted by CMD12. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetDabort(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldDabortMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldDabortMask)
	}
}

const (
	RegisterStarFieldDpsmactShift = 12
	RegisterStarFieldDpsmactMask  = 0x1000
)

// GetDpsmact Data path state machine active, i.e. not in Idle state. This is a hardware status flag only, does not generate an interrupt.
func (r *registerStarType) GetDpsmact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldDpsmactMask) != 0
}

// SetDpsmact Data path state machine active, i.e. not in Idle state. This is a hardware status flag only, does not generate an interrupt.
func (r *registerStarType) SetDpsmact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldDpsmactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldDpsmactMask)
	}
}

const (
	RegisterStarFieldCpsmactShift = 13
	RegisterStarFieldCpsmactMask  = 0x2000
)

// GetCpsmact Command path state machine active, i.e. not in Idle state. This is a hardware status flag only, does not generate an interrupt.
func (r *registerStarType) GetCpsmact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldCpsmactMask) != 0
}

// SetCpsmact Command path state machine active, i.e. not in Idle state. This is a hardware status flag only, does not generate an interrupt.
func (r *registerStarType) SetCpsmact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldCpsmactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldCpsmactMask)
	}
}

const (
	RegisterStarFieldTxfifoheShift = 14
	RegisterStarFieldTxfifoheMask  = 0x4000
)

// GetTxfifohe Transmit FIFO half empty At least half the number of words can be written into the FIFO. This bit is cleared when the FIFO becomes half+1 full.
func (r *registerStarType) GetTxfifohe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldTxfifoheMask) != 0
}

// SetTxfifohe Transmit FIFO half empty At least half the number of words can be written into the FIFO. This bit is cleared when the FIFO becomes half+1 full.
func (r *registerStarType) SetTxfifohe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldTxfifoheMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldTxfifoheMask)
	}
}

const (
	RegisterStarFieldRxfifohfShift = 15
	RegisterStarFieldRxfifohfMask  = 0x8000
)

// GetRxfifohf Receive FIFO half full There are at least half the number of words in the FIFO. This bit is cleared when the FIFO becomes half+1 empty.
func (r *registerStarType) GetRxfifohf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldRxfifohfMask) != 0
}

// SetRxfifohf Receive FIFO half full There are at least half the number of words in the FIFO. This bit is cleared when the FIFO becomes half+1 empty.
func (r *registerStarType) SetRxfifohf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldRxfifohfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldRxfifohfMask)
	}
}

const (
	RegisterStarFieldTxfifofShift = 16
	RegisterStarFieldTxfifofMask  = 0x10000
)

// GetTxfifof Transmit FIFO full This is a hardware status flag only, does not generate an interrupt. This bit is cleared when one FIFO location becomes empty.
func (r *registerStarType) GetTxfifof() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldTxfifofMask) != 0
}

// SetTxfifof Transmit FIFO full This is a hardware status flag only, does not generate an interrupt. This bit is cleared when one FIFO location becomes empty.
func (r *registerStarType) SetTxfifof(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldTxfifofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldTxfifofMask)
	}
}

const (
	RegisterStarFieldRxfifofShift = 17
	RegisterStarFieldRxfifofMask  = 0x20000
)

// GetRxfifof Receive FIFO full This bit is cleared when one FIFO location becomes empty.
func (r *registerStarType) GetRxfifof() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldRxfifofMask) != 0
}

// SetRxfifof Receive FIFO full This bit is cleared when one FIFO location becomes empty.
func (r *registerStarType) SetRxfifof(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldRxfifofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldRxfifofMask)
	}
}

const (
	RegisterStarFieldTxfifoeShift = 18
	RegisterStarFieldTxfifoeMask  = 0x40000
)

// GetTxfifoe Transmit FIFO empty This bit is cleared when one FIFO location becomes full.
func (r *registerStarType) GetTxfifoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldTxfifoeMask) != 0
}

// SetTxfifoe Transmit FIFO empty This bit is cleared when one FIFO location becomes full.
func (r *registerStarType) SetTxfifoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldTxfifoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldTxfifoeMask)
	}
}

const (
	RegisterStarFieldRxfifoeShift = 19
	RegisterStarFieldRxfifoeMask  = 0x80000
)

// GetRxfifoe Receive FIFO empty This is a hardware status flag only, does not generate an interrupt. This bit is cleared when one FIFO location becomes full.
func (r *registerStarType) GetRxfifoe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldRxfifoeMask) != 0
}

// SetRxfifoe Receive FIFO empty This is a hardware status flag only, does not generate an interrupt. This bit is cleared when one FIFO location becomes full.
func (r *registerStarType) SetRxfifoe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldRxfifoeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldRxfifoeMask)
	}
}

const (
	RegisterStarFieldBusyd0Shift = 20
	RegisterStarFieldBusyd0Mask  = 0x100000
)

// GetBusyd0 Inverted value of SDMMC_D0 line (Busy), sampled at the end of a CMD response and a second time 2 SDMMC_CK cycles after the CMD response. This bit is reset to not busy when the SDMMCD0 line changes from busy to not busy. This bit does not signal busy due to data transfer. This is a hardware status flag only, it does not generate an interrupt.
func (r *registerStarType) GetBusyd0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldBusyd0Mask) != 0
}

// SetBusyd0 Inverted value of SDMMC_D0 line (Busy), sampled at the end of a CMD response and a second time 2 SDMMC_CK cycles after the CMD response. This bit is reset to not busy when the SDMMCD0 line changes from busy to not busy. This bit does not signal busy due to data transfer. This is a hardware status flag only, it does not generate an interrupt.
func (r *registerStarType) SetBusyd0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldBusyd0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldBusyd0Mask)
	}
}

const (
	RegisterStarFieldBusyd0endShift = 21
	RegisterStarFieldBusyd0endMask  = 0x200000
)

// GetBusyd0end end of SDMMC_D0 Busy following a CMD response detected. This indicates only end of busy following a CMD response. This bit does not signal busy due to data transfer. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetBusyd0end() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldBusyd0endMask) != 0
}

// SetBusyd0end end of SDMMC_D0 Busy following a CMD response detected. This indicates only end of busy following a CMD response. This bit does not signal busy due to data transfer. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetBusyd0end(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldBusyd0endMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldBusyd0endMask)
	}
}

const (
	RegisterStarFieldSdioitShift = 22
	RegisterStarFieldSdioitMask  = 0x400000
)

// GetSdioit SDIO interrupt received. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetSdioit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldSdioitMask) != 0
}

// SetSdioit SDIO interrupt received. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetSdioit(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldSdioitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldSdioitMask)
	}
}

const (
	RegisterStarFieldAckfailShift = 23
	RegisterStarFieldAckfailMask  = 0x800000
)

// GetAckfail Boot acknowledgment received (boot acknowledgment check fail). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetAckfail() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldAckfailMask) != 0
}

// SetAckfail Boot acknowledgment received (boot acknowledgment check fail). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetAckfail(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldAckfailMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldAckfailMask)
	}
}

const (
	RegisterStarFieldAcktimeoutShift = 24
	RegisterStarFieldAcktimeoutMask  = 0x1000000
)

// GetAcktimeout Boot acknowledgment timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetAcktimeout() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldAcktimeoutMask) != 0
}

// SetAcktimeout Boot acknowledgment timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetAcktimeout(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldAcktimeoutMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldAcktimeoutMask)
	}
}

const (
	RegisterStarFieldVswendShift = 25
	RegisterStarFieldVswendMask  = 0x2000000
)

// GetVswend Voltage switch critical timing section completion. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetVswend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldVswendMask) != 0
}

// SetVswend Voltage switch critical timing section completion. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetVswend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldVswendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldVswendMask)
	}
}

const (
	RegisterStarFieldCkstopShift = 26
	RegisterStarFieldCkstopMask  = 0x4000000
)

// GetCkstop SDMMC_CK stopped in Voltage switch procedure. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetCkstop() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldCkstopMask) != 0
}

// SetCkstop SDMMC_CK stopped in Voltage switch procedure. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetCkstop(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldCkstopMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldCkstopMask)
	}
}

const (
	RegisterStarFieldIdmateShift = 27
	RegisterStarFieldIdmateMask  = 0x8000000
)

// GetIdmate IDMA transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetIdmate() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldIdmateMask) != 0
}

// SetIdmate IDMA transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetIdmate(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldIdmateMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldIdmateMask)
	}
}

const (
	RegisterStarFieldIdmabtcShift = 28
	RegisterStarFieldIdmabtcMask  = 0x10000000
)

// GetIdmabtc IDMA buffer transfer complete. interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) GetIdmabtc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterStarFieldIdmabtcMask) != 0
}

// SetIdmabtc IDMA buffer transfer complete. interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.
func (r *registerStarType) SetIdmabtc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterStarFieldIdmabtcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterStarFieldIdmabtcMask)
	}
}

// registerIcrType The SDMMC_ICR register is a write-only register. Writing a bit with 1 clears the corresponding bit in the SDMMC_STAR status register.
type registerIcrType uint32

const (
	RegisterIcrFieldCcrcfailcShift = 0
	RegisterIcrFieldCcrcfailcMask  = 0x1
)

// GetCcrcfailc CCRCFAIL flag clear bit Set by software to clear the CCRCFAIL flag.
func (r *registerIcrType) GetCcrcfailc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCcrcfailcMask) != 0
}

// SetCcrcfailc CCRCFAIL flag clear bit Set by software to clear the CCRCFAIL flag.
func (r *registerIcrType) SetCcrcfailc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCcrcfailcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCcrcfailcMask)
	}
}

const (
	RegisterIcrFieldDcrcfailcShift = 1
	RegisterIcrFieldDcrcfailcMask  = 0x2
)

// GetDcrcfailc DCRCFAIL flag clear bit Set by software to clear the DCRCFAIL flag.
func (r *registerIcrType) GetDcrcfailc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldDcrcfailcMask) != 0
}

// SetDcrcfailc DCRCFAIL flag clear bit Set by software to clear the DCRCFAIL flag.
func (r *registerIcrType) SetDcrcfailc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldDcrcfailcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldDcrcfailcMask)
	}
}

const (
	RegisterIcrFieldCtimeoutcShift = 2
	RegisterIcrFieldCtimeoutcMask  = 0x4
)

// GetCtimeoutc CTIMEOUT flag clear bit Set by software to clear the CTIMEOUT flag.
func (r *registerIcrType) GetCtimeoutc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCtimeoutcMask) != 0
}

// SetCtimeoutc CTIMEOUT flag clear bit Set by software to clear the CTIMEOUT flag.
func (r *registerIcrType) SetCtimeoutc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCtimeoutcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCtimeoutcMask)
	}
}

const (
	RegisterIcrFieldDtimeoutcShift = 3
	RegisterIcrFieldDtimeoutcMask  = 0x8
)

// GetDtimeoutc DTIMEOUT flag clear bit Set by software to clear the DTIMEOUT flag.
func (r *registerIcrType) GetDtimeoutc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldDtimeoutcMask) != 0
}

// SetDtimeoutc DTIMEOUT flag clear bit Set by software to clear the DTIMEOUT flag.
func (r *registerIcrType) SetDtimeoutc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldDtimeoutcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldDtimeoutcMask)
	}
}

const (
	RegisterIcrFieldTxunderrcShift = 4
	RegisterIcrFieldTxunderrcMask  = 0x10
)

// GetTxunderrc TXUNDERR flag clear bit Set by software to clear TXUNDERR flag.
func (r *registerIcrType) GetTxunderrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldTxunderrcMask) != 0
}

// SetTxunderrc TXUNDERR flag clear bit Set by software to clear TXUNDERR flag.
func (r *registerIcrType) SetTxunderrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldTxunderrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldTxunderrcMask)
	}
}

const (
	RegisterIcrFieldRxoverrcShift = 5
	RegisterIcrFieldRxoverrcMask  = 0x20
)

// GetRxoverrc RXOVERR flag clear bit Set by software to clear the RXOVERR flag.
func (r *registerIcrType) GetRxoverrc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldRxoverrcMask) != 0
}

// SetRxoverrc RXOVERR flag clear bit Set by software to clear the RXOVERR flag.
func (r *registerIcrType) SetRxoverrc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldRxoverrcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldRxoverrcMask)
	}
}

const (
	RegisterIcrFieldCmdrendcShift = 6
	RegisterIcrFieldCmdrendcMask  = 0x40
)

// GetCmdrendc CMDREND flag clear bit Set by software to clear the CMDREND flag.
func (r *registerIcrType) GetCmdrendc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCmdrendcMask) != 0
}

// SetCmdrendc CMDREND flag clear bit Set by software to clear the CMDREND flag.
func (r *registerIcrType) SetCmdrendc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCmdrendcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCmdrendcMask)
	}
}

const (
	RegisterIcrFieldCmdsentcShift = 7
	RegisterIcrFieldCmdsentcMask  = 0x80
)

// GetCmdsentc CMDSENT flag clear bit Set by software to clear the CMDSENT flag.
func (r *registerIcrType) GetCmdsentc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCmdsentcMask) != 0
}

// SetCmdsentc CMDSENT flag clear bit Set by software to clear the CMDSENT flag.
func (r *registerIcrType) SetCmdsentc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCmdsentcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCmdsentcMask)
	}
}

const (
	RegisterIcrFieldDataendcShift = 8
	RegisterIcrFieldDataendcMask  = 0x100
)

// GetDataendc DATAEND flag clear bit Set by software to clear the DATAEND flag.
func (r *registerIcrType) GetDataendc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldDataendcMask) != 0
}

// SetDataendc DATAEND flag clear bit Set by software to clear the DATAEND flag.
func (r *registerIcrType) SetDataendc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldDataendcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldDataendcMask)
	}
}

const (
	RegisterIcrFieldDholdcShift = 9
	RegisterIcrFieldDholdcMask  = 0x200
)

// GetDholdc DHOLD flag clear bit Set by software to clear the DHOLD flag.
func (r *registerIcrType) GetDholdc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldDholdcMask) != 0
}

// SetDholdc DHOLD flag clear bit Set by software to clear the DHOLD flag.
func (r *registerIcrType) SetDholdc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldDholdcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldDholdcMask)
	}
}

const (
	RegisterIcrFieldDbckendcShift = 10
	RegisterIcrFieldDbckendcMask  = 0x400
)

// GetDbckendc DBCKEND flag clear bit Set by software to clear the DBCKEND flag.
func (r *registerIcrType) GetDbckendc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldDbckendcMask) != 0
}

// SetDbckendc DBCKEND flag clear bit Set by software to clear the DBCKEND flag.
func (r *registerIcrType) SetDbckendc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldDbckendcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldDbckendcMask)
	}
}

const (
	RegisterIcrFieldDabortcShift = 11
	RegisterIcrFieldDabortcMask  = 0x800
)

// GetDabortc DABORT flag clear bit Set by software to clear the DABORT flag.
func (r *registerIcrType) GetDabortc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldDabortcMask) != 0
}

// SetDabortc DABORT flag clear bit Set by software to clear the DABORT flag.
func (r *registerIcrType) SetDabortc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldDabortcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldDabortcMask)
	}
}

const (
	RegisterIcrFieldBusyd0endcShift = 21
	RegisterIcrFieldBusyd0endcMask  = 0x200000
)

// GetBusyd0endc BUSYD0END flag clear bit Set by software to clear the BUSYD0END flag.
func (r *registerIcrType) GetBusyd0endc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldBusyd0endcMask) != 0
}

// SetBusyd0endc BUSYD0END flag clear bit Set by software to clear the BUSYD0END flag.
func (r *registerIcrType) SetBusyd0endc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldBusyd0endcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldBusyd0endcMask)
	}
}

const (
	RegisterIcrFieldSdioitcShift = 22
	RegisterIcrFieldSdioitcMask  = 0x400000
)

// GetSdioitc SDIOIT flag clear bit Set by software to clear the SDIOIT flag.
func (r *registerIcrType) GetSdioitc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldSdioitcMask) != 0
}

// SetSdioitc SDIOIT flag clear bit Set by software to clear the SDIOIT flag.
func (r *registerIcrType) SetSdioitc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldSdioitcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldSdioitcMask)
	}
}

const (
	RegisterIcrFieldAckfailcShift = 23
	RegisterIcrFieldAckfailcMask  = 0x800000
)

// GetAckfailc ACKFAIL flag clear bit Set by software to clear the ACKFAIL flag.
func (r *registerIcrType) GetAckfailc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldAckfailcMask) != 0
}

// SetAckfailc ACKFAIL flag clear bit Set by software to clear the ACKFAIL flag.
func (r *registerIcrType) SetAckfailc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldAckfailcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldAckfailcMask)
	}
}

const (
	RegisterIcrFieldAcktimeoutcShift = 24
	RegisterIcrFieldAcktimeoutcMask  = 0x1000000
)

// GetAcktimeoutc ACKTIMEOUT flag clear bit Set by software to clear the ACKTIMEOUT flag.
func (r *registerIcrType) GetAcktimeoutc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldAcktimeoutcMask) != 0
}

// SetAcktimeoutc ACKTIMEOUT flag clear bit Set by software to clear the ACKTIMEOUT flag.
func (r *registerIcrType) SetAcktimeoutc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldAcktimeoutcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldAcktimeoutcMask)
	}
}

const (
	RegisterIcrFieldVswendcShift = 25
	RegisterIcrFieldVswendcMask  = 0x2000000
)

// GetVswendc VSWEND flag clear bit Set by software to clear the VSWEND flag.
func (r *registerIcrType) GetVswendc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldVswendcMask) != 0
}

// SetVswendc VSWEND flag clear bit Set by software to clear the VSWEND flag.
func (r *registerIcrType) SetVswendc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldVswendcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldVswendcMask)
	}
}

const (
	RegisterIcrFieldCkstopcShift = 26
	RegisterIcrFieldCkstopcMask  = 0x4000000
)

// GetCkstopc CKSTOP flag clear bit Set by software to clear the CKSTOP flag.
func (r *registerIcrType) GetCkstopc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldCkstopcMask) != 0
}

// SetCkstopc CKSTOP flag clear bit Set by software to clear the CKSTOP flag.
func (r *registerIcrType) SetCkstopc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldCkstopcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldCkstopcMask)
	}
}

const (
	RegisterIcrFieldIdmatecShift = 27
	RegisterIcrFieldIdmatecMask  = 0x8000000
)

// GetIdmatec IDMA transfer error clear bit Set by software to clear the IDMATE flag.
func (r *registerIcrType) GetIdmatec() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldIdmatecMask) != 0
}

// SetIdmatec IDMA transfer error clear bit Set by software to clear the IDMATE flag.
func (r *registerIcrType) SetIdmatec(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldIdmatecMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldIdmatecMask)
	}
}

const (
	RegisterIcrFieldIdmabtccShift = 28
	RegisterIcrFieldIdmabtccMask  = 0x10000000
)

// GetIdmabtcc IDMA buffer transfer complete clear bit Set by software to clear the IDMABTC flag.
func (r *registerIcrType) GetIdmabtcc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIcrFieldIdmabtccMask) != 0
}

// SetIdmabtcc IDMA buffer transfer complete clear bit Set by software to clear the IDMABTC flag.
func (r *registerIcrType) SetIdmabtcc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIcrFieldIdmabtccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIcrFieldIdmabtccMask)
	}
}

// registerMaskrType The interrupt mask register determines which status flags generate an interrupt request by setting the corresponding bit to 1.
type registerMaskrType uint32

const (
	RegisterMaskrFieldCcrcfailieShift = 0
	RegisterMaskrFieldCcrcfailieMask  = 0x1
)

// GetCcrcfailie Command CRC fail interrupt enable Set and cleared by software to enable/disable interrupt caused by command CRC failure.
func (r *registerMaskrType) GetCcrcfailie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldCcrcfailieMask) != 0
}

// SetCcrcfailie Command CRC fail interrupt enable Set and cleared by software to enable/disable interrupt caused by command CRC failure.
func (r *registerMaskrType) SetCcrcfailie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldCcrcfailieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldCcrcfailieMask)
	}
}

const (
	RegisterMaskrFieldDcrcfailieShift = 1
	RegisterMaskrFieldDcrcfailieMask  = 0x2
)

// GetDcrcfailie Data CRC fail interrupt enable Set and cleared by software to enable/disable interrupt caused by data CRC failure.
func (r *registerMaskrType) GetDcrcfailie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldDcrcfailieMask) != 0
}

// SetDcrcfailie Data CRC fail interrupt enable Set and cleared by software to enable/disable interrupt caused by data CRC failure.
func (r *registerMaskrType) SetDcrcfailie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldDcrcfailieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldDcrcfailieMask)
	}
}

const (
	RegisterMaskrFieldCtimeoutieShift = 2
	RegisterMaskrFieldCtimeoutieMask  = 0x4
)

// GetCtimeoutie Command timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by command timeout.
func (r *registerMaskrType) GetCtimeoutie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldCtimeoutieMask) != 0
}

// SetCtimeoutie Command timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by command timeout.
func (r *registerMaskrType) SetCtimeoutie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldCtimeoutieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldCtimeoutieMask)
	}
}

const (
	RegisterMaskrFieldDtimeoutieShift = 3
	RegisterMaskrFieldDtimeoutieMask  = 0x8
)

// GetDtimeoutie Data timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by data timeout.
func (r *registerMaskrType) GetDtimeoutie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldDtimeoutieMask) != 0
}

// SetDtimeoutie Data timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by data timeout.
func (r *registerMaskrType) SetDtimeoutie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldDtimeoutieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldDtimeoutieMask)
	}
}

const (
	RegisterMaskrFieldTxunderrieShift = 4
	RegisterMaskrFieldTxunderrieMask  = 0x10
)

// GetTxunderrie Tx FIFO underrun error interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO underrun error.
func (r *registerMaskrType) GetTxunderrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldTxunderrieMask) != 0
}

// SetTxunderrie Tx FIFO underrun error interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO underrun error.
func (r *registerMaskrType) SetTxunderrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldTxunderrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldTxunderrieMask)
	}
}

const (
	RegisterMaskrFieldRxoverrieShift = 5
	RegisterMaskrFieldRxoverrieMask  = 0x20
)

// GetRxoverrie Rx FIFO overrun error interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO overrun error.
func (r *registerMaskrType) GetRxoverrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldRxoverrieMask) != 0
}

// SetRxoverrie Rx FIFO overrun error interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO overrun error.
func (r *registerMaskrType) SetRxoverrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldRxoverrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldRxoverrieMask)
	}
}

const (
	RegisterMaskrFieldCmdrendieShift = 6
	RegisterMaskrFieldCmdrendieMask  = 0x40
)

// GetCmdrendie Command response received interrupt enable Set and cleared by software to enable/disable interrupt caused by receiving command response.
func (r *registerMaskrType) GetCmdrendie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldCmdrendieMask) != 0
}

// SetCmdrendie Command response received interrupt enable Set and cleared by software to enable/disable interrupt caused by receiving command response.
func (r *registerMaskrType) SetCmdrendie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldCmdrendieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldCmdrendieMask)
	}
}

const (
	RegisterMaskrFieldCmdsentieShift = 7
	RegisterMaskrFieldCmdsentieMask  = 0x80
)

// GetCmdsentie Command sent interrupt enable Set and cleared by software to enable/disable interrupt caused by sending command.
func (r *registerMaskrType) GetCmdsentie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldCmdsentieMask) != 0
}

// SetCmdsentie Command sent interrupt enable Set and cleared by software to enable/disable interrupt caused by sending command.
func (r *registerMaskrType) SetCmdsentie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldCmdsentieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldCmdsentieMask)
	}
}

const (
	RegisterMaskrFieldDataendieShift = 8
	RegisterMaskrFieldDataendieMask  = 0x100
)

// GetDataendie Data end interrupt enable Set and cleared by software to enable/disable interrupt caused by data end.
func (r *registerMaskrType) GetDataendie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldDataendieMask) != 0
}

// SetDataendie Data end interrupt enable Set and cleared by software to enable/disable interrupt caused by data end.
func (r *registerMaskrType) SetDataendie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldDataendieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldDataendieMask)
	}
}

const (
	RegisterMaskrFieldDholdieShift = 9
	RegisterMaskrFieldDholdieMask  = 0x200
)

// GetDholdie Data hold interrupt enable Set and cleared by software to enable/disable the interrupt generated when sending new data is hold in the DPSM Wait_S state.
func (r *registerMaskrType) GetDholdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldDholdieMask) != 0
}

// SetDholdie Data hold interrupt enable Set and cleared by software to enable/disable the interrupt generated when sending new data is hold in the DPSM Wait_S state.
func (r *registerMaskrType) SetDholdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldDholdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldDholdieMask)
	}
}

const (
	RegisterMaskrFieldDbckendieShift = 10
	RegisterMaskrFieldDbckendieMask  = 0x400
)

// GetDbckendie Data block end interrupt enable Set and cleared by software to enable/disable interrupt caused by data block end.
func (r *registerMaskrType) GetDbckendie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldDbckendieMask) != 0
}

// SetDbckendie Data block end interrupt enable Set and cleared by software to enable/disable interrupt caused by data block end.
func (r *registerMaskrType) SetDbckendie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldDbckendieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldDbckendieMask)
	}
}

const (
	RegisterMaskrFieldDabortieShift = 11
	RegisterMaskrFieldDabortieMask  = 0x800
)

// GetDabortie Data transfer aborted interrupt enable Set and cleared by software to enable/disable interrupt caused by a data transfer being aborted.
func (r *registerMaskrType) GetDabortie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldDabortieMask) != 0
}

// SetDabortie Data transfer aborted interrupt enable Set and cleared by software to enable/disable interrupt caused by a data transfer being aborted.
func (r *registerMaskrType) SetDabortie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldDabortieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldDabortieMask)
	}
}

const (
	RegisterMaskrFieldTxfifoheieShift = 14
	RegisterMaskrFieldTxfifoheieMask  = 0x4000
)

// GetTxfifoheie Tx FIFO half empty interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO half empty.
func (r *registerMaskrType) GetTxfifoheie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldTxfifoheieMask) != 0
}

// SetTxfifoheie Tx FIFO half empty interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO half empty.
func (r *registerMaskrType) SetTxfifoheie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldTxfifoheieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldTxfifoheieMask)
	}
}

const (
	RegisterMaskrFieldRxfifohfieShift = 15
	RegisterMaskrFieldRxfifohfieMask  = 0x8000
)

// GetRxfifohfie Rx FIFO half full interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO half full.
func (r *registerMaskrType) GetRxfifohfie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldRxfifohfieMask) != 0
}

// SetRxfifohfie Rx FIFO half full interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO half full.
func (r *registerMaskrType) SetRxfifohfie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldRxfifohfieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldRxfifohfieMask)
	}
}

const (
	RegisterMaskrFieldRxfifofieShift = 17
	RegisterMaskrFieldRxfifofieMask  = 0x20000
)

// GetRxfifofie Rx FIFO full interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO full.
func (r *registerMaskrType) GetRxfifofie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldRxfifofieMask) != 0
}

// SetRxfifofie Rx FIFO full interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO full.
func (r *registerMaskrType) SetRxfifofie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldRxfifofieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldRxfifofieMask)
	}
}

const (
	RegisterMaskrFieldTxfifoeieShift = 18
	RegisterMaskrFieldTxfifoeieMask  = 0x40000
)

// GetTxfifoeie Tx FIFO empty interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO empty.
func (r *registerMaskrType) GetTxfifoeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldTxfifoeieMask) != 0
}

// SetTxfifoeie Tx FIFO empty interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO empty.
func (r *registerMaskrType) SetTxfifoeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldTxfifoeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldTxfifoeieMask)
	}
}

const (
	RegisterMaskrFieldBusyd0endieShift = 21
	RegisterMaskrFieldBusyd0endieMask  = 0x200000
)

// GetBusyd0endie BUSYD0END interrupt enable Set and cleared by software to enable/disable the interrupt generated when SDMMC_D0 signal changes from busy to NOT busy following a CMD response.
func (r *registerMaskrType) GetBusyd0endie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldBusyd0endieMask) != 0
}

// SetBusyd0endie BUSYD0END interrupt enable Set and cleared by software to enable/disable the interrupt generated when SDMMC_D0 signal changes from busy to NOT busy following a CMD response.
func (r *registerMaskrType) SetBusyd0endie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldBusyd0endieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldBusyd0endieMask)
	}
}

const (
	RegisterMaskrFieldSdioitieShift = 22
	RegisterMaskrFieldSdioitieMask  = 0x400000
)

// GetSdioitie SDIO mode interrupt received interrupt enable Set and cleared by software to enable/disable the interrupt generated when receiving the SDIO mode interrupt.
func (r *registerMaskrType) GetSdioitie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldSdioitieMask) != 0
}

// SetSdioitie SDIO mode interrupt received interrupt enable Set and cleared by software to enable/disable the interrupt generated when receiving the SDIO mode interrupt.
func (r *registerMaskrType) SetSdioitie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldSdioitieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldSdioitieMask)
	}
}

const (
	RegisterMaskrFieldAckfailieShift = 23
	RegisterMaskrFieldAckfailieMask  = 0x800000
)

// GetAckfailie Acknowledgment Fail interrupt enable Set and cleared by software to enable/disable interrupt caused by acknowledgment Fail.
func (r *registerMaskrType) GetAckfailie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldAckfailieMask) != 0
}

// SetAckfailie Acknowledgment Fail interrupt enable Set and cleared by software to enable/disable interrupt caused by acknowledgment Fail.
func (r *registerMaskrType) SetAckfailie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldAckfailieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldAckfailieMask)
	}
}

const (
	RegisterMaskrFieldAcktimeoutieShift = 24
	RegisterMaskrFieldAcktimeoutieMask  = 0x1000000
)

// GetAcktimeoutie Acknowledgment timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by acknowledgment timeout.
func (r *registerMaskrType) GetAcktimeoutie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldAcktimeoutieMask) != 0
}

// SetAcktimeoutie Acknowledgment timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by acknowledgment timeout.
func (r *registerMaskrType) SetAcktimeoutie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldAcktimeoutieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldAcktimeoutieMask)
	}
}

const (
	RegisterMaskrFieldVswendieShift = 25
	RegisterMaskrFieldVswendieMask  = 0x2000000
)

// GetVswendie Voltage switch critical timing section completion interrupt enable Set and cleared by software to enable/disable the interrupt generated when voltage switch critical timing section completion.
func (r *registerMaskrType) GetVswendie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldVswendieMask) != 0
}

// SetVswendie Voltage switch critical timing section completion interrupt enable Set and cleared by software to enable/disable the interrupt generated when voltage switch critical timing section completion.
func (r *registerMaskrType) SetVswendie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldVswendieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldVswendieMask)
	}
}

const (
	RegisterMaskrFieldCkstopieShift = 26
	RegisterMaskrFieldCkstopieMask  = 0x4000000
)

// GetCkstopie Voltage Switch clock stopped interrupt enable Set and cleared by software to enable/disable interrupt caused by Voltage Switch clock stopped.
func (r *registerMaskrType) GetCkstopie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldCkstopieMask) != 0
}

// SetCkstopie Voltage Switch clock stopped interrupt enable Set and cleared by software to enable/disable interrupt caused by Voltage Switch clock stopped.
func (r *registerMaskrType) SetCkstopie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldCkstopieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldCkstopieMask)
	}
}

const (
	RegisterMaskrFieldIdmabtcieShift = 28
	RegisterMaskrFieldIdmabtcieMask  = 0x10000000
)

// GetIdmabtcie IDMA buffer transfer complete interrupt enable Set and cleared by software to enable/disable the interrupt generated when the IDMA has transferred all data belonging to a memory buffer.
func (r *registerMaskrType) GetIdmabtcie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaskrFieldIdmabtcieMask) != 0
}

// SetIdmabtcie IDMA buffer transfer complete interrupt enable Set and cleared by software to enable/disable the interrupt generated when the IDMA has transferred all data belonging to a memory buffer.
func (r *registerMaskrType) SetIdmabtcie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaskrFieldIdmabtcieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaskrFieldIdmabtcieMask)
	}
}

// registerAcktimerType The SDMMC_ACKTIMER register contains the acknowledgment timeout period, in SDMMC_CK bus clock periods. A counter loads the value from the SDMMC_ACKTIMER register, and starts decrementing when the data path state machine (DPSM) enters the Wait_Ack state. If the timer reaches 0 while the DPSM is in this states, the acknowledgment timeout status flag is set.
type registerAcktimerType uint32

const (
	RegisterAcktimerFieldAcktimeShift = 0
	RegisterAcktimerFieldAcktimeMask  = 0x1ffffff
)

// GetAcktime Boot acknowledgment timeout period This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). Boot acknowledgment timeout period expressed in card bus clock periods.
func (r *registerAcktimerType) GetAcktime() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterAcktimerFieldAcktimeMask) >> RegisterAcktimerFieldAcktimeShift)
}

// SetAcktime Boot acknowledgment timeout period This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). Boot acknowledgment timeout period expressed in card bus clock periods.
func (r *registerAcktimerType) SetAcktime(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAcktimerFieldAcktimeMask)|(uint32(value)<<RegisterAcktimerFieldAcktimeShift))
}

// registerIdmactrlrType The receive and transmit FIFOs can be read or written as 32-bit wide registers. The FIFOs contain 32 entries on 32 sequential addresses. This allows the CPU to use its load and store multiple operands to read from/write to the FIFO.
type registerIdmactrlrType uint32

const (
	RegisterIdmactrlrFieldIdmaenShift = 0
	RegisterIdmactrlrFieldIdmaenMask  = 0x1
)

// GetIdmaen IDMA enable This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerIdmactrlrType) GetIdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdmactrlrFieldIdmaenMask) != 0
}

// SetIdmaen IDMA enable This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerIdmactrlrType) SetIdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdmactrlrFieldIdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdmactrlrFieldIdmaenMask)
	}
}

const (
	RegisterIdmactrlrFieldIdmabmodeShift = 1
	RegisterIdmactrlrFieldIdmabmodeMask  = 0x2
)

// GetIdmabmode Buffer mode selection. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerIdmactrlrType) GetIdmabmode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdmactrlrFieldIdmabmodeMask) != 0
}

// SetIdmabmode Buffer mode selection. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerIdmactrlrType) SetIdmabmode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdmactrlrFieldIdmabmodeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdmactrlrFieldIdmabmodeMask)
	}
}

const (
	RegisterIdmactrlrFieldIdmabactShift = 2
	RegisterIdmactrlrFieldIdmabactMask  = 0x4
)

// GetIdmabact Double buffer mode active buffer indication This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). When IDMA is enabled this bit is toggled by hardware.
func (r *registerIdmactrlrType) GetIdmabact() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIdmactrlrFieldIdmabactMask) != 0
}

// SetIdmabact Double buffer mode active buffer indication This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). When IDMA is enabled this bit is toggled by hardware.
func (r *registerIdmactrlrType) SetIdmabact(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIdmactrlrFieldIdmabactMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIdmactrlrFieldIdmabactMask)
	}
}

// registerIdmabsizerType The SDMMC_IDMABSIZER register contains the buffers size when in double buffer configuration.
type registerIdmabsizerType uint32

const (
	RegisterIdmabsizerFieldIdmabndtShift = 5
	RegisterIdmabsizerFieldIdmabndtMask  = 0x1fe0
)

// GetIdmabndt Number of transfers per buffer. This 8-bit value shall be multiplied by 8 to get the size of the buffer in 32-bit words and by 32 to get the size of the buffer in bytes. Example: IDMABNDT = 0x01: buffer size = 8 words = 32 bytes. These bits can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerIdmabsizerType) GetIdmabndt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIdmabsizerFieldIdmabndtMask) >> RegisterIdmabsizerFieldIdmabndtShift)
}

// SetIdmabndt Number of transfers per buffer. This 8-bit value shall be multiplied by 8 to get the size of the buffer in 32-bit words and by 32 to get the size of the buffer in bytes. Example: IDMABNDT = 0x01: buffer size = 8 words = 32 bytes. These bits can only be written by firmware when DPSM is inactive (DPSMACT = 0).
func (r *registerIdmabsizerType) SetIdmabndt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmabsizerFieldIdmabndtMask)|(uint32(value)<<RegisterIdmabsizerFieldIdmabndtShift))
}

// registerIdmabase0rType The SDMMC_IDMABASE0R register contains the memory buffer base address in single buffer configuration and the buffer 0 base address in double buffer configuration.
type registerIdmabase0rType uint32

const (
	RegisterIdmabase0rFieldIdmabase0Shift = 0
	RegisterIdmabase0rFieldIdmabase0Mask  = 0xffffffff
)

// GetIdmabase0 Buffer 0 memory base address bits [31:2], shall be word aligned (bit [1:0] are always 0 and read only). This register can be written by firmware when DPSM is inactive (DPSMACT = 0), and can dynamically be written by firmware when DPSM active (DPSMACT = 1) and memory buffer 0 is inactive (IDMABACT = 1).
func (r *registerIdmabase0rType) GetIdmabase0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIdmabase0rFieldIdmabase0Mask) >> RegisterIdmabase0rFieldIdmabase0Shift)
}

// SetIdmabase0 Buffer 0 memory base address bits [31:2], shall be word aligned (bit [1:0] are always 0 and read only). This register can be written by firmware when DPSM is inactive (DPSMACT = 0), and can dynamically be written by firmware when DPSM active (DPSMACT = 1) and memory buffer 0 is inactive (IDMABACT = 1).
func (r *registerIdmabase0rType) SetIdmabase0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmabase0rFieldIdmabase0Mask)|(uint32(value)<<RegisterIdmabase0rFieldIdmabase0Shift))
}

// registerIdmabase1rType The SDMMC_IDMABASE1R register contains the double buffer configuration second buffer memory base address.
type registerIdmabase1rType uint32

const (
	RegisterIdmabase1rFieldIdmabase1Shift = 0
	RegisterIdmabase1rFieldIdmabase1Mask  = 0xffffffff
)

// GetIdmabase1 Buffer 1 memory base address, shall be word aligned (bit [1:0] are always 0 and read only). This register can be written by firmware when DPSM is inactive (DPSMACT = 0), and can dynamically be written by firmware when DPSM active (DPSMACT = 1) and memory buffer 1 is inactive (IDMABACT = 0).
func (r *registerIdmabase1rType) GetIdmabase1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIdmabase1rFieldIdmabase1Mask) >> RegisterIdmabase1rFieldIdmabase1Shift)
}

// SetIdmabase1 Buffer 1 memory base address, shall be word aligned (bit [1:0] are always 0 and read only). This register can be written by firmware when DPSM is inactive (DPSMACT = 0), and can dynamically be written by firmware when DPSM active (DPSMACT = 1) and memory buffer 1 is inactive (IDMABACT = 0).
func (r *registerIdmabase1rType) SetIdmabase1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIdmabase1rFieldIdmabase1Mask)|(uint32(value)<<RegisterIdmabase1rFieldIdmabase1Shift))
}

// registerFiforType The receive and transmit FIFOs can be only read or written as word (32-bit) wide registers. The FIFOs contain 16 entries on sequential addresses. This allows the CPU to use its load and store multiple operands to read from/write to the FIFO.When accessing SDMMC_FIFOR with half word or byte access an AHB bus fault is generated.
type registerFiforType uint32

const (
	RegisterFiforFieldFifodataShift = 0
	RegisterFiforFieldFifodataMask  = 0xffffffff
)

// GetFifodata Receive and transmit FIFO data This register can only be read or written by firmware when the DPSM is active (DPSMACT=1). The FIFO data occupies 16 entries of 32-bit words.
func (r *registerFiforType) GetFifodata() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterFiforFieldFifodataMask) >> RegisterFiforFieldFifodataShift)
}

// SetFifodata Receive and transmit FIFO data This register can only be read or written by firmware when the DPSM is active (DPSMACT=1). The FIFO data occupies 16 entries of 32-bit words.
func (r *registerFiforType) SetFifodata(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterFiforFieldFifodataMask)|(uint32(value)<<RegisterFiforFieldFifodataShift))
}
