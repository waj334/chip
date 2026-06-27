package dsihost

import (
	"pkg.si-go.dev/chip/arm/cortexm"
	regdsi "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/dsihost"
	corehal "pkg.si-go.dev/chip/core/hal"
)

func waitRegulatorReady() error {
	for i := uint32(0); i < defaultPollLimit; i++ {
		if regdsi.Dsihost.Dsiwisr.GetRrs() {
			return nil
		}
	}
	return corehal.ErrTimeout
}

func waitWrapperPLLLock() error {
	for i := uint32(0); i < defaultPollLimit; i++ {
		if regdsi.Dsihost.Dsiwisr.GetPllls() {
			return nil
		}
	}
	return corehal.ErrTimeout
}

// waitPHYStopState waits for the D-PHY to reach Stop state on the clock lane
// and the relevant data lanes. Matches the equivalent wait in ST HAL_DSI_Init.
func waitPHYStopState(lanes LaneCount) error {
	for i := uint32(0); i < defaultPollLimit; i++ {
		clockReady := regdsi.Dsihost.Dsipsr.GetPssc()
		data0Ready := regdsi.Dsihost.Dsipsr.GetPss0()
		data1Ready := true
		if lanes == TwoDataLanes {
			data1Ready = regdsi.Dsihost.Dsipsr.GetPss1()
		}

		if clockReady && data0Ready && data1Ready {
			return nil
		}
	}
	return corehal.ErrTimeout
}

func waitCommandWritable() error {
	for i := uint32(0); i < defaultPollLimit; i++ {
		if !regdsi.Dsihost.Dsigpsr.GetCmdff() {
			return nil
		}
	}
	return corehal.ErrTimeout
}

func waitPayloadWritable() error {
	for i := uint32(0); i < defaultPollLimit; i++ {
		if !regdsi.Dsihost.Dsigpsr.GetPwrff() {
			return nil
		}
	}
	return corehal.ErrTimeout
}

// waitCommandEmpty blocks until DSI_GPSR.CMDFE is set, i.e. the command FIFO
// has fully drained. ST HAL waits on this at the start of a long write.
func waitCommandEmpty() error {
	for i := uint32(0); i < defaultPollLimit; i++ {
		if regdsi.Dsihost.Dsigpsr.GetCmdfe() {
			return nil
		}
	}
	return corehal.ErrTimeout
}

func waitPayloadWriteEmpty() error {
	for i := uint32(0); i < defaultPollLimit; i++ {
		if regdsi.Dsihost.Dsigpsr.GetPwrfe() {
			return nil
		}
	}
	return corehal.ErrTimeout
}

func writePayloadWord(data []byte) {
	word := uint32(data[0]) |
		uint32(data[1])<<8 |
		uint32(data[2])<<16 |
		uint32(data[3])<<24

	regdsi.Dsihost.Dsigpdr.Store(word)
}

func makeHeader(dt PacketDataType, channelID uint8, data0 uint8, data1 uint8) uint32 {
	return uint32(dt)&0x3f |
		(uint32(channelID)&0x3)<<6 |
		uint32(data0)<<8 |
		uint32(data1)<<16
}

func writeHeader(dt PacketDataType, channelID uint8, wordCount uint16) error {
	if channelID > 3 {
		return corehal.ErrInvalidConfig
	}

	if err := waitCommandEmpty(); err != nil {
		return err
	}

	regdsi.Dsihost.Dsighcr.Store(makeHeader(
		dt,
		channelID,
		uint8(wordCount&0xff),
		uint8(wordCount>>8),
	))
	return nil
}

func writeHeaderWithParams(dt PacketDataType, channelID uint8, data0 uint8, data1 uint8) error {
	if channelID > 3 {
		return corehal.ErrInvalidConfig
	}

	if err := waitCommandEmpty(); err != nil {
		return err
	}

	regdsi.Dsihost.Dsighcr.Store(makeHeader(dt, channelID, data0, data1))
	cortexm.DSB()
	return nil
}

func configureTimeouts(t Timeouts) {
	if t.LowPowerRX == 0 {
		t.LowPowerRX = 0xffff
	}
	if t.HighSpeedTX == 0 {
		t.HighSpeedTX = 0xffff
	}
	if t.HighSpeedRead == 0 {
		t.HighSpeedRead = 0xffff
	}
	if t.LowPowerRead == 0 {
		t.LowPowerRead = 0xffff
	}
	if t.HighSpeedWrite == 0 {
		t.HighSpeedWrite = 0xffff
	}
	if t.LowPowerWrite == 0 {
		t.LowPowerWrite = 0xffff
	}
	if t.BTATimeout == 0 {
		t.BTATimeout = 0xffff
	}

	regdsi.Dsihost.Dsitccr0.SetLprxtocnt(t.LowPowerRX)
	regdsi.Dsihost.Dsitccr0.SetHstxtocnt(t.HighSpeedTX)

	regdsi.Dsihost.Dsitccr1.SetHsrdtocnt(t.HighSpeedRead)
	regdsi.Dsihost.Dsitccr2.SetLprdtocnt(t.LowPowerRead)

	regdsi.Dsihost.Dsitccr3.SetHswrtocnt(t.HighSpeedWrite)
	regdsi.Dsihost.Dsitccr3.SetPm(false)

	regdsi.Dsihost.Dsitccr4.SetLpwrtocnt(t.LowPowerWrite)
	regdsi.Dsihost.Dsitccr5.SetBtatocnt(t.BTATimeout)
}

func disableHostAndWrapper() {
	regdsi.Dsihost.Dsiwcr.SetLtdcen(false)
	regdsi.Dsihost.Dsiwcr.SetDsien(false)

	regdsi.Dsihost.Dsipctlr.SetCke(false)
	regdsi.Dsihost.Dsipctlr.SetDen(false)

	regdsi.Dsihost.Dsicr.SetEn(false)
	cortexm.DSB()
}

func clearWrapperInterruptFlags() {
	regdsi.Dsihost.Dsiwifcr.SetCteif(true)
	regdsi.Dsihost.Dsiwifcr.SetCerif(true)
	regdsi.Dsihost.Dsiwifcr.SetCplllif(true)
	regdsi.Dsihost.Dsiwifcr.SetCplluif(true)
	regdsi.Dsihost.Dsiwifcr.SetCrrif(true)
}
