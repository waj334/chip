//go:build stm32h7x7

package sdio

import (
	"sync"

	stm32h7x7 "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/sdmmc"

	corehal "pkg.si-go.dev/chip/core/hal"
	corepin "pkg.si-go.dev/chip/core/hal/pin"
)

const (
	SDIO1 _sdio = 0
	SDIO2 _sdio = 1
)

const (
	_ALT_CKIN = iota + 1
	_ALT_CDIR
	_ALT_D0DIR
	_ALT_D123DIR
	_ALT_D0
	_ALT_D1
	_ALT_D2
	_ALT_D3
	_ALT_D4
	_ALT_D5
	_ALT_D6
	_ALT_D7
	_ALT_CK
	_ALT_CMD
)

var instances = [2]sdio{}

type (
	_sdio uint8

	sdio struct {
		mutex sync.Mutex
	}

	Config struct {
		Enable  bool
		CKIN    pin.Pin
		CDIR    pin.Pin
		D0DIR   pin.Pin
		D123DIR pin.Pin
		Dn      [8]pin.Pin
		CK      pin.Pin
		CMD     pin.Pin

		Edge                bool
		BusWidth            uint8
		PowerSave           bool
		HardwareFlowControl bool
		FrequencyHz         uint16
	}
)

type altKey struct {
	pin      pin.Pin
	function int
	instance _sdio
	mode     corepin.Mode
}

var altFunctionTable = [...]altKey{
	// SDIO1
	{pin.PB8, _ALT_CKIN, SDIO1, pin.Alt7},
	{pin.PB8, _ALT_D4, SDIO1, pin.Alt12},
	{pin.PB9, _ALT_CDIR, SDIO1, pin.Alt7},
	{pin.PB9, _ALT_D5, SDIO1, pin.Alt12},
	{pin.PC6, _ALT_D0DIR, SDIO1, pin.Alt8},
	{pin.PC6, _ALT_D6, SDIO1, pin.Alt12},
	{pin.PC7, _ALT_D123DIR, SDIO1, pin.Alt8},
	{pin.PC7, _ALT_D7, SDIO1, pin.Alt12},
	{pin.PC8, _ALT_D0, SDIO1, pin.Alt12},
	{pin.PC9, _ALT_D1, SDIO1, pin.Alt12},
	{pin.PC10, _ALT_D2, SDIO1, pin.Alt12},
	{pin.PC11, _ALT_D3, SDIO1, pin.Alt12},
	{pin.PC12, _ALT_CK, SDIO1, pin.Alt12},
	{pin.PD2, _ALT_CMD, SDIO1, pin.Alt12},

	// SDIO2
	{pin.PA0, _ALT_CMD, SDIO2, pin.Alt9},
	{pin.PB3, _ALT_D2, SDIO2, pin.Alt9},
	{pin.PB4, _ALT_D3, SDIO2, pin.Alt9},
	{pin.PB8, _ALT_D4, SDIO2, pin.Alt10},
	{pin.PB9, _ALT_D5, SDIO2, pin.Alt10},
	{pin.PC14, _ALT_D0, SDIO2, pin.Alt9},
	{pin.PC15, _ALT_D1, SDIO2, pin.Alt9},
	{pin.PC1, _ALT_CK, SDIO2, pin.Alt9},
	{pin.PC6, _ALT_D6, SDIO2, pin.Alt10},
	{pin.PC7, _ALT_D7, SDIO2, pin.Alt10},
	{pin.PD6, _ALT_CK, SDIO2, pin.Alt11},
	{pin.PD7, _ALT_CMD, SDIO2, pin.Alt11},
	{pin.PG11, _ALT_D2, SDIO2, pin.Alt10},
}

func altFunction(p pin.Pin, alt int, instance _sdio) (corepin.Mode, error) {
	for _, entry := range altFunctionTable {
		if entry.pin == p && entry.function == alt && entry.instance == instance {
			return entry.mode, nil
		}
	}
	return 0, corehal.ErrInvalidPinout
}

/*
func altFunction(p pin.Pin, alt int, instance _sdio) (corepin.Mode, error) {
	var mode corepin.Mode = 0xFF

	switch instance {
	case SDIO1:
		switch p {
		case pin.PB8:
			switch alt {
			case _ALT_CKIN:
				mode = pin.Alt7
			case _ALT_D4:
				mode = pin.Alt12
			}
		case pin.PB9:
			switch alt {
			case _ALT_CDIR:
				mode = pin.Alt7
			case _ALT_D5:
				mode = pin.Alt12
			}
		case pin.PC6:
			switch alt {
			case _ALT_D0DIR:
				mode = pin.Alt8
			case _ALT_D6:
				mode = pin.Alt12
			}
		case pin.PC7:
			switch alt {
			case _ALT_D123DIR:
				mode = pin.Alt8
			case _ALT_D7:
				mode = pin.Alt12
			}
		case pin.PC8:
			switch alt {
			case _ALT_D0:
				mode = pin.Alt12
			}
		case pin.PC9:
			switch alt {
			case _ALT_D1:
				mode = pin.Alt12
			}
		case pin.PC10:
			switch alt {
			case _ALT_D2:
				mode = pin.Alt12
			}
		case pin.PC11:
			switch alt {
			case _ALT_D3:
				mode = pin.Alt12
			}
		case pin.PC12:
			switch alt {
			case _ALT_CK:
				mode = pin.Alt12
			}
		case pin.PD2:
			switch alt {
			case _ALT_CMD:
				mode = pin.Alt12
			}
		}
	case SDIO2:
		switch p {
		case pin.PA0:
			switch alt {
			case _ALT_CMD:
				mode = pin.Alt9
			}
		case pin.PB3:
			switch alt {
			case _ALT_D2:
				mode = pin.Alt9
			}
		case pin.PB4:
			switch alt {
			case _ALT_D3:
				mode = pin.Alt9
			}
		case pin.PB8:
			switch alt {
			case _ALT_D4:
				mode = pin.Alt10
			}
		case pin.PB9:
			switch alt {
			case _ALT_D5:
				mode = pin.Alt10
			}
		case pin.PC14:
			switch alt {
			case _ALT_D0:
				mode = pin.Alt9
			}
		case pin.PC15:
			switch alt {
			case _ALT_D1:
				mode = pin.Alt9
			}
		case pin.PC1:
			switch alt {
			case _ALT_CK:
				mode = pin.Alt9
			}
		case pin.PC6:
			switch alt {
			case _ALT_D6:
				mode = pin.Alt10
			}
		case pin.PC7:
			switch alt {
			case _ALT_D7:
				mode = pin.Alt10
			}
		case pin.PD6:
			switch alt {
			case _ALT_CK:
				mode = pin.Alt11
			}
		case pin.PD7:
			switch alt {
			case _ALT_CMD:
				mode = pin.Alt11
			}
		case pin.PG11:
			switch alt {
			case _ALT_D2:
				mode = pin.Alt10
			}
		}
	}

	if mode == 0xFF {
		return 0, corehal.ErrInvalidPinout
	}
	return mode, nil
}
*/

func (s _sdio) Configure(config Config) error {
	regs := sdmmc.Instances[s]
	_s := &instances[s]

	_s.mutex.Lock()

	switch s {
	case SDIO1:
		stm32h7x7.IrqSdmmc1.Disable()
		rcc.Rcc.Ahb3rstr.SetSdmmc1rst(true)
		for !rcc.Rcc.Ahb3rstr.GetSdmmc1rst() {
		}

		rcc.Rcc.Ahb3rstr.SetSdmmc1rst(false)
		for rcc.Rcc.Ahb3rstr.GetSdmmc1rst() {
		}
	case SDIO2:
		stm32h7x7.IrqSdmmc2.Disable()
		rcc.Rcc.Ahb2rstr.SetSdmmc2rst(true)
		for !rcc.Rcc.Ahb2rstr.GetSdmmc2rst() {
		}

		rcc.Rcc.Ahb2rstr.SetSdmmc2rst(false)
		for rcc.Rcc.Ahb2rstr.GetSdmmc2rst() {
		}
	}

	if !config.Enable {
		_s.mutex.Unlock()
		return nil
	}

	// Enable SDIO in RCC.
	switch s {
	case SDIO1:
		rcc.Rcc.Ahb3enr.SetSdmmc1en(true)
		for !rcc.Rcc.Ahb3enr.GetSdmmc1en() {
		}
	case SDIO2:
		rcc.Rcc.Ahb2enr.SetSdmmc2en(true)
		for !rcc.Rcc.Ahb2enr.GetSdmmc2en() {
		}
	}

	// Set up pins.
	pinCfg := func(p pin.Pin, alt int) error {
		if p != 0 {
			mode, err := altFunction(p, alt, s)
			if err != nil {
				return err
			}
			config.CKIN.SetMode(pin.AltFunction | mode)
		}
		return nil
	}

	if err := pinCfg(config.CKIN, _ALT_CKIN); err != nil {
		return err
	}

	if err := pinCfg(config.CDIR, _ALT_CDIR); err != nil {
		return err
	}

	if err := pinCfg(config.D0DIR, _ALT_D0DIR); err != nil {
		return err
	}

	if err := pinCfg(config.D123DIR, _ALT_D123DIR); err != nil {
		return err
	}

	if err := pinCfg(config.CK, _ALT_CK); err != nil {
		return err
	}

	if err := pinCfg(config.CMD, _ALT_CMD); err != nil {
		return err
	}

	for i, Dn := range config.Dn {
		if err := pinCfg(Dn, _ALT_D0+i); err != nil {
			return err
		}
	}

	// Set the SDMMC configuration parameters.
	regs.Clkcr.SetNegedge(config.Edge)
	regs.Clkcr.SetBusspeed(false)
	regs.Clkcr.SetWidbus(config.BusWidth)
	regs.Clkcr.SetHwfcen(config.HardwareFlowControl)

	// Calculate clock divider.
	div := uint16(hal.SdmmcSourceFrequencyHz / uint64(2*config.FrequencyHz))
	regs.Clkcr.SetClkdiv(div)

	// Soft reset the SDIO interface.
	regs.Power.SetPwrctrl(sdmmc.RegisterPowerFieldPwrctrlEnumPoweron)

	_s.mutex.Unlock()
	return nil
}
