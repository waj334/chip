package cec

import (
	"unsafe"
	"volatile"
)

var (
	Cec = (*_cec)(unsafe.Pointer(uintptr(0x40006c00)))
)

type _cec struct {
	Cr   registerCrType
	Cfgr registerCfgrType
	Txdr registerTxdrType
	Rxdr registerRxdrType
	Isr  registerIsrType
	Ier  registerIerType
}

// registerCrType CEC control register
type registerCrType uint32

const (
	RegisterCrFieldCecenShift = 0
	RegisterCrFieldCecenMask  = 0x1
)

// GetCecen CEC Enable The CECEN bit is set and cleared by software. CECEN=1 starts message reception and enables the TXSOM control. CECEN=0 disables the CEC peripheral, clears all bits of CEC_CR register and aborts any on-going reception or transmission.
func (r *registerCrType) GetCecen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldCecenMask) != 0
}

// SetCecen CEC Enable The CECEN bit is set and cleared by software. CECEN=1 starts message reception and enables the TXSOM control. CECEN=0 disables the CEC peripheral, clears all bits of CEC_CR register and aborts any on-going reception or transmission.
func (r *registerCrType) SetCecen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldCecenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldCecenMask)
	}
}

const (
	RegisterCrFieldTxsomShift = 1
	RegisterCrFieldTxsomMask  = 0x2
)

// GetTxsom Tx Start Of Message TXSOM is set by software to command transmission of the first byte of a CEC message. If the CEC message consists of only one byte, TXEOM must be set before of TXSOM. Start-Bit is effectively started on the CEC line after SFT is counted. If TXSOM is set while a message reception is ongoing, transmission will start after the end of reception. TXSOM is cleared by hardware after the last byte of the message is sent with a positive acknowledge (TXEND=1), in case of transmission underrun (TXUDR=1), negative acknowledge (TXACKE=1), and transmission error (TXERR=1). It is also cleared by CECEN=0. It is not cleared and transmission is automatically retried in case of arbitration lost (ARBLST=1). TXSOM can be also used as a status bit informing application whether any transmission request is pending or under execution. The application can abort a transmission request at any time by clearing the CECEN bit. Note: TXSOM must be set when CECEN=1 TXSOM must be set when transmission data is available into TXDR HEADERs first four bits containing own peripheral address are taken from TXDR[7:4], not from CEC_CFGR.OAR which is used only for reception
func (r *registerCrType) GetTxsom() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTxsomMask) != 0
}

// SetTxsom Tx Start Of Message TXSOM is set by software to command transmission of the first byte of a CEC message. If the CEC message consists of only one byte, TXEOM must be set before of TXSOM. Start-Bit is effectively started on the CEC line after SFT is counted. If TXSOM is set while a message reception is ongoing, transmission will start after the end of reception. TXSOM is cleared by hardware after the last byte of the message is sent with a positive acknowledge (TXEND=1), in case of transmission underrun (TXUDR=1), negative acknowledge (TXACKE=1), and transmission error (TXERR=1). It is also cleared by CECEN=0. It is not cleared and transmission is automatically retried in case of arbitration lost (ARBLST=1). TXSOM can be also used as a status bit informing application whether any transmission request is pending or under execution. The application can abort a transmission request at any time by clearing the CECEN bit. Note: TXSOM must be set when CECEN=1 TXSOM must be set when transmission data is available into TXDR HEADERs first four bits containing own peripheral address are taken from TXDR[7:4], not from CEC_CFGR.OAR which is used only for reception
func (r *registerCrType) SetTxsom(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTxsomMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTxsomMask)
	}
}

const (
	RegisterCrFieldTxeomShift = 2
	RegisterCrFieldTxeomMask  = 0x4
)

// GetTxeom Tx End Of Message The TXEOM bit is set by software to command transmission of the last byte of a CEC message. TXEOM is cleared by hardware at the same time and under the same conditions as for TXSOM. Note: TXEOM must be set when CECEN=1 TXEOM must be set before writing transmission data to TXDR If TXEOM is set when TXSOM=0, transmitted message will consist of 1 byte (HEADER) only (PING message)
func (r *registerCrType) GetTxeom() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldTxeomMask) != 0
}

// SetTxeom Tx End Of Message The TXEOM bit is set by software to command transmission of the last byte of a CEC message. TXEOM is cleared by hardware at the same time and under the same conditions as for TXSOM. Note: TXEOM must be set when CECEN=1 TXEOM must be set before writing transmission data to TXDR If TXEOM is set when TXSOM=0, transmitted message will consist of 1 byte (HEADER) only (PING message)
func (r *registerCrType) SetTxeom(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldTxeomMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldTxeomMask)
	}
}

// registerCfgrType This register is used to configure the HDMI-CEC controller. It is mandatory to write CEC_CFGR only when CECEN=0.
type registerCfgrType uint32

const (
	RegisterCfgrFieldSftShift = 0
	RegisterCfgrFieldSftMask  = 0x7
)

// GetSft Signal Free Time SFT bits are set by software. In the SFT=0x0 configuration the number of nominal data bit periods waited before transmission is ruled by hardware according to the transmission history. In all the other configurations the SFT number is determined by software. * 0x0 ** 2.5 Data-Bit periods if CEC is the last bus initiator with unsuccessful transmission (ARBLST=1, TXERR=1, TXUDR=1 or TXACKE= 1) ** 4 Data-Bit periods if CEC is the new bus initiator ** 6 Data-Bit periods if CEC is the last bus initiator with successful transmission (TXEOM=1) * 0x1: 0.5 nominal data bit periods * 0x2: 1.5 nominal data bit periods * 0x3: 2.5 nominal data bit periods * 0x4: 3.5 nominal data bit periods * 0x5: 4.5 nominal data bit periods * 0x6: 5.5 nominal data bit periods * 0x7: 6.5 nominal data bit periods
func (r *registerCfgrType) GetSft() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSftMask) >> RegisterCfgrFieldSftShift)
}

// SetSft Signal Free Time SFT bits are set by software. In the SFT=0x0 configuration the number of nominal data bit periods waited before transmission is ruled by hardware according to the transmission history. In all the other configurations the SFT number is determined by software. * 0x0 ** 2.5 Data-Bit periods if CEC is the last bus initiator with unsuccessful transmission (ARBLST=1, TXERR=1, TXUDR=1 or TXACKE= 1) ** 4 Data-Bit periods if CEC is the new bus initiator ** 6 Data-Bit periods if CEC is the last bus initiator with successful transmission (TXEOM=1) * 0x1: 0.5 nominal data bit periods * 0x2: 1.5 nominal data bit periods * 0x3: 2.5 nominal data bit periods * 0x4: 3.5 nominal data bit periods * 0x5: 4.5 nominal data bit periods * 0x6: 5.5 nominal data bit periods * 0x7: 6.5 nominal data bit periods
func (r *registerCfgrType) SetSft(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSftMask)|(uint32(value)<<RegisterCfgrFieldSftShift))
}

const (
	RegisterCfgrFieldRxtolShift = 3
	RegisterCfgrFieldRxtolMask  = 0x8
)

// GetRxtol Rx-Tolerance The RXTOL bit is set and cleared by software. ** Start-Bit, +/- 200 s rise, +/- 200 s fall. ** Data-Bit: +/- 200 s rise. +/- 350 s fall. ** Start-Bit: +/- 400 s rise, +/- 400 s fall ** Data-Bit: +/-300 s rise, +/- 500 s fall
func (r *registerCfgrType) GetRxtol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldRxtolMask) != 0
}

// SetRxtol Rx-Tolerance The RXTOL bit is set and cleared by software. ** Start-Bit, +/- 200 s rise, +/- 200 s fall. ** Data-Bit: +/- 200 s rise. +/- 350 s fall. ** Start-Bit: +/- 400 s rise, +/- 400 s fall ** Data-Bit: +/-300 s rise, +/- 500 s fall
func (r *registerCfgrType) SetRxtol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldRxtolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldRxtolMask)
	}
}

const (
	RegisterCfgrFieldBrestpShift = 4
	RegisterCfgrFieldBrestpMask  = 0x10
)

// GetBrestp Rx-Stop on Bit Rising Error The BRESTP bit is set and cleared by software.
func (r *registerCfgrType) GetBrestp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldBrestpMask) != 0
}

// SetBrestp Rx-Stop on Bit Rising Error The BRESTP bit is set and cleared by software.
func (r *registerCfgrType) SetBrestp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldBrestpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldBrestpMask)
	}
}

const (
	RegisterCfgrFieldBregenShift = 5
	RegisterCfgrFieldBregenMask  = 0x20
)

// GetBregen Generate Error-Bit on Bit Rising Error The BREGEN bit is set and cleared by software. Note: If BRDNOGEN=0, an Error-bit is generated upon BRE detection with BRESTP=1 in broadcast even if BREGEN=0
func (r *registerCfgrType) GetBregen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldBregenMask) != 0
}

// SetBregen Generate Error-Bit on Bit Rising Error The BREGEN bit is set and cleared by software. Note: If BRDNOGEN=0, an Error-bit is generated upon BRE detection with BRESTP=1 in broadcast even if BREGEN=0
func (r *registerCfgrType) SetBregen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldBregenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldBregenMask)
	}
}

const (
	RegisterCfgrFieldLbpegenShift = 6
	RegisterCfgrFieldLbpegenMask  = 0x40
)

// GetLbpegen Generate Error-Bit on Long Bit Period Error The LBPEGEN bit is set and cleared by software. Note: If BRDNOGEN=0, an Error-bit is generated upon LBPE detection in broadcast even if LBPEGEN=0
func (r *registerCfgrType) GetLbpegen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldLbpegenMask) != 0
}

// SetLbpegen Generate Error-Bit on Long Bit Period Error The LBPEGEN bit is set and cleared by software. Note: If BRDNOGEN=0, an Error-bit is generated upon LBPE detection in broadcast even if LBPEGEN=0
func (r *registerCfgrType) SetLbpegen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldLbpegenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldLbpegenMask)
	}
}

const (
	RegisterCfgrFieldBrdnogenShift = 7
	RegisterCfgrFieldBrdnogenMask  = 0x80
)

// GetBrdnogen Avoid Error-Bit Generation in Broadcast The BRDNOGEN bit is set and cleared by software.
func (r *registerCfgrType) GetBrdnogen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldBrdnogenMask) != 0
}

// SetBrdnogen Avoid Error-Bit Generation in Broadcast The BRDNOGEN bit is set and cleared by software.
func (r *registerCfgrType) SetBrdnogen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldBrdnogenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldBrdnogenMask)
	}
}

const (
	RegisterCfgrFieldSftoptShift = 8
	RegisterCfgrFieldSftoptMask  = 0x100
)

// GetSftopt SFT Option Bit The SFTOPT bit is set and cleared by software.
func (r *registerCfgrType) GetSftopt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldSftoptMask) != 0
}

// SetSftopt SFT Option Bit The SFTOPT bit is set and cleared by software.
func (r *registerCfgrType) SetSftopt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldSftoptMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldSftoptMask)
	}
}

const (
	RegisterCfgrFieldOarShift = 16
	RegisterCfgrFieldOarMask  = 0x7fff0000
)

// GetOar Own addresses configuration The OAR bits are set by software to select which destination logical addresses has to be considered in receive mode. Each bit, when set, enables the CEC logical address identified by the given bit position. At the end of HEADER reception, the received destination address is compared with the enabled addresses. In case of matching address, the incoming message is acknowledged and received. In case of non-matching address, the incoming message is received only in listen mode (LSTN=1), but without acknowledge sent. Broadcast messages are always received. Example: OAR = 0b000 0000 0010 0001 means that CEC acknowledges addresses 0x0 and 0x5. Consequently, each message directed to one of these addresses is received.
func (r *registerCfgrType) GetOar() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldOarMask) >> RegisterCfgrFieldOarShift)
}

// SetOar Own addresses configuration The OAR bits are set by software to select which destination logical addresses has to be considered in receive mode. Each bit, when set, enables the CEC logical address identified by the given bit position. At the end of HEADER reception, the received destination address is compared with the enabled addresses. In case of matching address, the incoming message is acknowledged and received. In case of non-matching address, the incoming message is received only in listen mode (LSTN=1), but without acknowledge sent. Broadcast messages are always received. Example: OAR = 0b000 0000 0010 0001 means that CEC acknowledges addresses 0x0 and 0x5. Consequently, each message directed to one of these addresses is received.
func (r *registerCfgrType) SetOar(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldOarMask)|(uint32(value)<<RegisterCfgrFieldOarShift))
}

const (
	RegisterCfgrFieldLstnShift = 31
	RegisterCfgrFieldLstnMask  = 0x80000000
)

// GetLstn Listen mode LSTN bit is set and cleared by software.
func (r *registerCfgrType) GetLstn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfgrFieldLstnMask) != 0
}

// SetLstn Listen mode LSTN bit is set and cleared by software.
func (r *registerCfgrType) SetLstn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfgrFieldLstnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfgrFieldLstnMask)
	}
}

// registerTxdrType CEC Tx data register
type registerTxdrType uint32

const (
	RegisterTxdrFieldTxdShift = 0
	RegisterTxdrFieldTxdMask  = 0xff
)

// GetTxd Tx Data register. TXD is a write-only register containing the data byte to be transmitted. Note: TXD must be written when TXSTART=1
func (r *registerTxdrType) GetTxd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTxdrFieldTxdMask) >> RegisterTxdrFieldTxdShift)
}

// SetTxd Tx Data register. TXD is a write-only register containing the data byte to be transmitted. Note: TXD must be written when TXSTART=1
func (r *registerTxdrType) SetTxd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxdrFieldTxdMask)|(uint32(value)<<RegisterTxdrFieldTxdShift))
}

// registerRxdrType CEC Rx Data Register
type registerRxdrType uint32

const (
	RegisterRxdrFieldRxdShift = 0
	RegisterRxdrFieldRxdMask  = 0xff
)

// GetRxd Rx Data register. RXD is read-only and contains the last data byte which has been received from the CEC line.
func (r *registerRxdrType) GetRxd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterRxdrFieldRxdMask) >> RegisterRxdrFieldRxdShift)
}

// SetRxd Rx Data register. RXD is read-only and contains the last data byte which has been received from the CEC line.
func (r *registerRxdrType) SetRxd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxdrFieldRxdMask)|(uint32(value)<<RegisterRxdrFieldRxdShift))
}

// registerIsrType CEC Interrupt and Status Register
type registerIsrType uint32

const (
	RegisterIsrFieldRxbrShift = 0
	RegisterIsrFieldRxbrMask  = 0x1
)

// GetRxbr Rx-Byte Received The RXBR bit is set by hardware to inform application that a new byte has been received from the CEC line and stored into the RXD buffer. RXBR is cleared by software write at 1.
func (r *registerIsrType) GetRxbr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxbrMask) != 0
}

// SetRxbr Rx-Byte Received The RXBR bit is set by hardware to inform application that a new byte has been received from the CEC line and stored into the RXD buffer. RXBR is cleared by software write at 1.
func (r *registerIsrType) SetRxbr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRxbrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRxbrMask)
	}
}

const (
	RegisterIsrFieldRxendShift = 1
	RegisterIsrFieldRxendMask  = 0x2
)

// GetRxend End Of Reception RXEND is set by hardware to inform application that the last byte of a CEC message is received from the CEC line and stored into the RXD buffer. RXEND is set at the same time of RXBR. RXEND is cleared by software write at 1.
func (r *registerIsrType) GetRxend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxendMask) != 0
}

// SetRxend End Of Reception RXEND is set by hardware to inform application that the last byte of a CEC message is received from the CEC line and stored into the RXD buffer. RXEND is set at the same time of RXBR. RXEND is cleared by software write at 1.
func (r *registerIsrType) SetRxend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRxendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRxendMask)
	}
}

const (
	RegisterIsrFieldRxovrShift = 2
	RegisterIsrFieldRxovrMask  = 0x4
)

// GetRxovr Rx-Overrun RXOVR is set by hardware if RXBR is not yet cleared at the time a new byte is received on the CEC line and stored into RXD. RXOVR assertion stops message reception so that no acknowledge is sent. In case of broadcast, a negative acknowledge is sent. RXOVR is cleared by software write at 1.
func (r *registerIsrType) GetRxovr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxovrMask) != 0
}

// SetRxovr Rx-Overrun RXOVR is set by hardware if RXBR is not yet cleared at the time a new byte is received on the CEC line and stored into RXD. RXOVR assertion stops message reception so that no acknowledge is sent. In case of broadcast, a negative acknowledge is sent. RXOVR is cleared by software write at 1.
func (r *registerIsrType) SetRxovr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRxovrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRxovrMask)
	}
}

const (
	RegisterIsrFieldBreShift = 3
	RegisterIsrFieldBreMask  = 0x8
)

// GetBre Rx-Bit Rising Error BRE is set by hardware in case a Data-Bit waveform is detected with Bit Rising Error. BRE is set either at the time the misplaced rising edge occurs, or at the end of the maximum BRE tolerance allowed by RXTOL, in case rising edge is still longing. BRE stops message reception if BRESTP=1. BRE generates an Error-Bit on the CEC line if BREGEN=1. BRE is cleared by software write at 1.
func (r *registerIsrType) GetBre() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldBreMask) != 0
}

// SetBre Rx-Bit Rising Error BRE is set by hardware in case a Data-Bit waveform is detected with Bit Rising Error. BRE is set either at the time the misplaced rising edge occurs, or at the end of the maximum BRE tolerance allowed by RXTOL, in case rising edge is still longing. BRE stops message reception if BRESTP=1. BRE generates an Error-Bit on the CEC line if BREGEN=1. BRE is cleared by software write at 1.
func (r *registerIsrType) SetBre(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldBreMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldBreMask)
	}
}

const (
	RegisterIsrFieldSbpeShift = 4
	RegisterIsrFieldSbpeMask  = 0x10
)

// GetSbpe Rx-Short Bit Period Error SBPE is set by hardware in case a Data-Bit waveform is detected with Short Bit Period Error. SBPE is set at the time the anticipated falling edge occurs. SBPE generates an Error-Bit on the CEC line. SBPE is cleared by software write at 1.
func (r *registerIsrType) GetSbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldSbpeMask) != 0
}

// SetSbpe Rx-Short Bit Period Error SBPE is set by hardware in case a Data-Bit waveform is detected with Short Bit Period Error. SBPE is set at the time the anticipated falling edge occurs. SBPE generates an Error-Bit on the CEC line. SBPE is cleared by software write at 1.
func (r *registerIsrType) SetSbpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldSbpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldSbpeMask)
	}
}

const (
	RegisterIsrFieldLbpeShift = 5
	RegisterIsrFieldLbpeMask  = 0x20
)

// GetLbpe Rx-Long Bit Period Error LBPE is set by hardware in case a Data-Bit waveform is detected with Long Bit Period Error. LBPE is set at the end of the maximum bit-extension tolerance allowed by RXTOL, in case falling edge is still longing. LBPE always stops reception of the CEC message. LBPE generates an Error-Bit on the CEC line if LBPEGEN=1. In case of broadcast, Error-Bit is generated even in case of LBPEGEN=0. LBPE is cleared by software write at 1.
func (r *registerIsrType) GetLbpe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldLbpeMask) != 0
}

// SetLbpe Rx-Long Bit Period Error LBPE is set by hardware in case a Data-Bit waveform is detected with Long Bit Period Error. LBPE is set at the end of the maximum bit-extension tolerance allowed by RXTOL, in case falling edge is still longing. LBPE always stops reception of the CEC message. LBPE generates an Error-Bit on the CEC line if LBPEGEN=1. In case of broadcast, Error-Bit is generated even in case of LBPEGEN=0. LBPE is cleared by software write at 1.
func (r *registerIsrType) SetLbpe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldLbpeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldLbpeMask)
	}
}

const (
	RegisterIsrFieldRxackeShift = 6
	RegisterIsrFieldRxackeMask  = 0x40
)

// GetRxacke Rx-Missing Acknowledge In receive mode, RXACKE is set by hardware to inform application that no acknowledge was seen on the CEC line. RXACKE applies only for broadcast messages and in listen mode also for not directly addressed messages (destination address not enabled in OAR). RXACKE aborts message reception. RXACKE is cleared by software write at 1.
func (r *registerIsrType) GetRxacke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldRxackeMask) != 0
}

// SetRxacke Rx-Missing Acknowledge In receive mode, RXACKE is set by hardware to inform application that no acknowledge was seen on the CEC line. RXACKE applies only for broadcast messages and in listen mode also for not directly addressed messages (destination address not enabled in OAR). RXACKE aborts message reception. RXACKE is cleared by software write at 1.
func (r *registerIsrType) SetRxacke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldRxackeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldRxackeMask)
	}
}

const (
	RegisterIsrFieldArblstShift = 7
	RegisterIsrFieldArblstMask  = 0x80
)

// GetArblst Arbitration Lost ARBLST is set by hardware to inform application that CEC device is switching to reception due to arbitration lost event following the TXSOM command. ARBLST can be due either to a contending CEC device starting earlier or starting at the same time but with higher HEADER priority. After ARBLST assertion TXSOM bit keeps pending for next transmission attempt. ARBLST is cleared by software write at 1.
func (r *registerIsrType) GetArblst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldArblstMask) != 0
}

// SetArblst Arbitration Lost ARBLST is set by hardware to inform application that CEC device is switching to reception due to arbitration lost event following the TXSOM command. ARBLST can be due either to a contending CEC device starting earlier or starting at the same time but with higher HEADER priority. After ARBLST assertion TXSOM bit keeps pending for next transmission attempt. ARBLST is cleared by software write at 1.
func (r *registerIsrType) SetArblst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldArblstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldArblstMask)
	}
}

const (
	RegisterIsrFieldTxbrShift = 8
	RegisterIsrFieldTxbrMask  = 0x100
)

// GetTxbr Tx-Byte Request TXBR is set by hardware to inform application that the next transmission data has to be written to TXDR. TXBR is set when the 4th bit of currently transmitted byte is sent. Application must write the next byte to TXDR within 6 nominal data-bit periods before transmission underrun error occurs (TXUDR). TXBR is cleared by software write at 1.
func (r *registerIsrType) GetTxbr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxbrMask) != 0
}

// SetTxbr Tx-Byte Request TXBR is set by hardware to inform application that the next transmission data has to be written to TXDR. TXBR is set when the 4th bit of currently transmitted byte is sent. Application must write the next byte to TXDR within 6 nominal data-bit periods before transmission underrun error occurs (TXUDR). TXBR is cleared by software write at 1.
func (r *registerIsrType) SetTxbr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxbrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxbrMask)
	}
}

const (
	RegisterIsrFieldTxendShift = 9
	RegisterIsrFieldTxendMask  = 0x200
)

// GetTxend End of Transmission TXEND is set by hardware to inform application that the last byte of the CEC message has been successfully transmitted. TXEND clears the TXSOM and TXEOM control bits. TXEND is cleared by software write at 1.
func (r *registerIsrType) GetTxend() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxendMask) != 0
}

// SetTxend End of Transmission TXEND is set by hardware to inform application that the last byte of the CEC message has been successfully transmitted. TXEND clears the TXSOM and TXEOM control bits. TXEND is cleared by software write at 1.
func (r *registerIsrType) SetTxend(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxendMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxendMask)
	}
}

const (
	RegisterIsrFieldTxudrShift = 10
	RegisterIsrFieldTxudrMask  = 0x400
)

// GetTxudr Tx-Buffer Underrun In transmission mode, TXUDR is set by hardware if application was not in time to load TXDR before of next byte transmission. TXUDR aborts message transmission and clears TXSOM and TXEOM control bits. TXUDR is cleared by software write at 1
func (r *registerIsrType) GetTxudr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxudrMask) != 0
}

// SetTxudr Tx-Buffer Underrun In transmission mode, TXUDR is set by hardware if application was not in time to load TXDR before of next byte transmission. TXUDR aborts message transmission and clears TXSOM and TXEOM control bits. TXUDR is cleared by software write at 1
func (r *registerIsrType) SetTxudr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxudrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxudrMask)
	}
}

const (
	RegisterIsrFieldTxerrShift = 11
	RegisterIsrFieldTxerrMask  = 0x800
)

// GetTxerr Tx-Error In transmission mode, TXERR is set by hardware if the CEC initiator detects low impedance on the CEC line while it is released. TXERR aborts message transmission and clears TXSOM and TXEOM controls. TXERR is cleared by software write at 1.
func (r *registerIsrType) GetTxerr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxerrMask) != 0
}

// SetTxerr Tx-Error In transmission mode, TXERR is set by hardware if the CEC initiator detects low impedance on the CEC line while it is released. TXERR aborts message transmission and clears TXSOM and TXEOM controls. TXERR is cleared by software write at 1.
func (r *registerIsrType) SetTxerr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxerrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxerrMask)
	}
}

const (
	RegisterIsrFieldTxackeShift = 12
	RegisterIsrFieldTxackeMask  = 0x1000
)

// GetTxacke Tx-Missing Acknowledge Error In transmission mode, TXACKE is set by hardware to inform application that no acknowledge was received. In case of broadcast transmission, TXACKE informs application that a negative acknowledge was received. TXACKE aborts message transmission and clears TXSOM and TXEOM controls. TXACKE is cleared by software write at 1.
func (r *registerIsrType) GetTxacke() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIsrFieldTxackeMask) != 0
}

// SetTxacke Tx-Missing Acknowledge Error In transmission mode, TXACKE is set by hardware to inform application that no acknowledge was received. In case of broadcast transmission, TXACKE informs application that a negative acknowledge was received. TXACKE aborts message transmission and clears TXSOM and TXEOM controls. TXACKE is cleared by software write at 1.
func (r *registerIsrType) SetTxacke(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIsrFieldTxackeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIsrFieldTxackeMask)
	}
}

// registerIerType CEC interrupt enable register
type registerIerType uint32

const (
	RegisterIerFieldRxbrieShift = 0
	RegisterIerFieldRxbrieMask  = 0x1
)

// GetRxbrie Rx-Byte Received Interrupt Enable The RXBRIE bit is set and cleared by software.
func (r *registerIerType) GetRxbrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRxbrieMask) != 0
}

// SetRxbrie Rx-Byte Received Interrupt Enable The RXBRIE bit is set and cleared by software.
func (r *registerIerType) SetRxbrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRxbrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRxbrieMask)
	}
}

const (
	RegisterIerFieldRxendieShift = 1
	RegisterIerFieldRxendieMask  = 0x2
)

// GetRxendie End Of Reception Interrupt Enable The RXENDIE bit is set and cleared by software.
func (r *registerIerType) GetRxendie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRxendieMask) != 0
}

// SetRxendie End Of Reception Interrupt Enable The RXENDIE bit is set and cleared by software.
func (r *registerIerType) SetRxendie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRxendieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRxendieMask)
	}
}

const (
	RegisterIerFieldRxovrieShift = 2
	RegisterIerFieldRxovrieMask  = 0x4
)

// GetRxovrie Rx-Buffer Overrun Interrupt Enable The RXOVRIE bit is set and cleared by software.
func (r *registerIerType) GetRxovrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRxovrieMask) != 0
}

// SetRxovrie Rx-Buffer Overrun Interrupt Enable The RXOVRIE bit is set and cleared by software.
func (r *registerIerType) SetRxovrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRxovrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRxovrieMask)
	}
}

const (
	RegisterIerFieldBreieShift = 3
	RegisterIerFieldBreieMask  = 0x8
)

// GetBreie Bit Rising Error Interrupt Enable The BREIE bit is set and cleared by software.
func (r *registerIerType) GetBreie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldBreieMask) != 0
}

// SetBreie Bit Rising Error Interrupt Enable The BREIE bit is set and cleared by software.
func (r *registerIerType) SetBreie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldBreieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldBreieMask)
	}
}

const (
	RegisterIerFieldSbpeieShift = 4
	RegisterIerFieldSbpeieMask  = 0x10
)

// GetSbpeie Short Bit Period Error Interrupt Enable The SBPEIE bit is set and cleared by software.
func (r *registerIerType) GetSbpeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldSbpeieMask) != 0
}

// SetSbpeie Short Bit Period Error Interrupt Enable The SBPEIE bit is set and cleared by software.
func (r *registerIerType) SetSbpeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldSbpeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldSbpeieMask)
	}
}

const (
	RegisterIerFieldLbpeieShift = 5
	RegisterIerFieldLbpeieMask  = 0x20
)

// GetLbpeie Long Bit Period Error Interrupt Enable The LBPEIE bit is set and cleared by software.
func (r *registerIerType) GetLbpeie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldLbpeieMask) != 0
}

// SetLbpeie Long Bit Period Error Interrupt Enable The LBPEIE bit is set and cleared by software.
func (r *registerIerType) SetLbpeie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldLbpeieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldLbpeieMask)
	}
}

const (
	RegisterIerFieldRxackieShift = 6
	RegisterIerFieldRxackieMask  = 0x40
)

// GetRxackie Rx-Missing Acknowledge Error Interrupt Enable The RXACKIE bit is set and cleared by software.
func (r *registerIerType) GetRxackie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldRxackieMask) != 0
}

// SetRxackie Rx-Missing Acknowledge Error Interrupt Enable The RXACKIE bit is set and cleared by software.
func (r *registerIerType) SetRxackie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldRxackieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldRxackieMask)
	}
}

const (
	RegisterIerFieldArblstieShift = 7
	RegisterIerFieldArblstieMask  = 0x80
)

// GetArblstie Arbitration Lost Interrupt Enable The ARBLSTIE bit is set and cleared by software.
func (r *registerIerType) GetArblstie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldArblstieMask) != 0
}

// SetArblstie Arbitration Lost Interrupt Enable The ARBLSTIE bit is set and cleared by software.
func (r *registerIerType) SetArblstie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldArblstieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldArblstieMask)
	}
}

const (
	RegisterIerFieldTxbrieShift = 8
	RegisterIerFieldTxbrieMask  = 0x100
)

// GetTxbrie Tx-Byte Request Interrupt Enable The TXBRIE bit is set and cleared by software.
func (r *registerIerType) GetTxbrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTxbrieMask) != 0
}

// SetTxbrie Tx-Byte Request Interrupt Enable The TXBRIE bit is set and cleared by software.
func (r *registerIerType) SetTxbrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTxbrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTxbrieMask)
	}
}

const (
	RegisterIerFieldTxendieShift = 9
	RegisterIerFieldTxendieMask  = 0x200
)

// GetTxendie Tx-End Of Message Interrupt Enable The TXENDIE bit is set and cleared by software.
func (r *registerIerType) GetTxendie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTxendieMask) != 0
}

// SetTxendie Tx-End Of Message Interrupt Enable The TXENDIE bit is set and cleared by software.
func (r *registerIerType) SetTxendie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTxendieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTxendieMask)
	}
}

const (
	RegisterIerFieldTxudrieShift = 10
	RegisterIerFieldTxudrieMask  = 0x400
)

// GetTxudrie Tx-Underrun Interrupt Enable The TXUDRIE bit is set and cleared by software.
func (r *registerIerType) GetTxudrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTxudrieMask) != 0
}

// SetTxudrie Tx-Underrun Interrupt Enable The TXUDRIE bit is set and cleared by software.
func (r *registerIerType) SetTxudrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTxudrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTxudrieMask)
	}
}

const (
	RegisterIerFieldTxerrieShift = 11
	RegisterIerFieldTxerrieMask  = 0x800
)

// GetTxerrie Tx-Error Interrupt Enable The TXERRIE bit is set and cleared by software.
func (r *registerIerType) GetTxerrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTxerrieMask) != 0
}

// SetTxerrie Tx-Error Interrupt Enable The TXERRIE bit is set and cleared by software.
func (r *registerIerType) SetTxerrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTxerrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTxerrieMask)
	}
}

const (
	RegisterIerFieldTxackieShift = 12
	RegisterIerFieldTxackieMask  = 0x1000
)

// GetTxackie Tx-Missing Acknowledge Error Interrupt Enable The TXACKEIE bit is set and cleared by software.
func (r *registerIerType) GetTxackie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterIerFieldTxackieMask) != 0
}

// SetTxackie Tx-Missing Acknowledge Error Interrupt Enable The TXACKEIE bit is set and cleared by software.
func (r *registerIerType) SetTxackie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterIerFieldTxackieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterIerFieldTxackieMask)
	}
}
