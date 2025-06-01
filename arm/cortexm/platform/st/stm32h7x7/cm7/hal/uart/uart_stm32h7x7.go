package uart

import (
	"runtime"
	"sync"
	"time"
	"volatile"

	corehal "pkg.si-go.dev/chip/core/hal"
	corepin "pkg.si-go.dev/chip/core/hal/pin"

	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/pin"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/rcc"
	"pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/usart"
)

type _error int

const (
	UART1 _uart = 0
	UART2 _uart = 1
	UART3 _uart = 2
	UART4 _uart = 3
	UART5 _uart = 4
	UART6 _uart = 5
	UART7 _uart = 6
	UART8 _uart = 7

	errTimeout     _error = -1
	defaultTimeout        = 100 * time.Millisecond
)

func (e _error) Error() string {
	switch e {
	case 0:
		return "uart: no error"
	case errTimeout:
		return "uart: timeout"
	default:
		return "uart: unknown error"
	}
}

var (
	mutex [8]sync.Mutex
)

type _uart uint8

type Config struct {
	Enable          bool
	ReceiveEnabled  bool
	TransmitEnabled bool
	CharacterSize   uint8
	Baud            uint32
	TX              pin.Pin
	RX              pin.Pin
	CTS             pin.Pin
	RTS             pin.Pin
	CK              pin.Pin
	NumStopBits     uint8
}

const (
	_TX = iota + 1
	_RX
	_CTS
	_RTS
	_CK
)

func altFunction(p pin.Pin, t int, instance _uart) (mode corepin.Mode, err error) {
	tt := 0
	switch instance {
	case UART1:
		switch p {
		case pin.PA8:
			tt, mode = _CK, pin.Alt7
		case pin.PA9:
			tt, mode = _TX, pin.Alt7
		case pin.PA10:
			tt, mode = _RX, pin.Alt7
		case pin.PA11:
			tt, mode = _CTS, pin.Alt7
		case pin.PA12:
			tt, mode = _RTS, pin.Alt7
		case pin.PB6:
			tt, mode = _TX, pin.Alt7
		case pin.PB7:
			tt, mode = _RX, pin.Alt7
		case pin.PB14:
			tt, mode = _TX, pin.Alt4
		case pin.PB15:
			tt, mode = _RX, pin.Alt4
		}
	case UART2:
		switch p {
		case pin.PA0:
			tt, mode = _CTS, pin.Alt7
		case pin.PA1:
			tt, mode = _RTS, pin.Alt7
		case pin.PA2:
			tt, mode = _TX, pin.Alt7
		case pin.PA3:
			tt, mode = _RX, pin.Alt7
		case pin.PA4:
			tt, mode = _CK, pin.Alt7
		case pin.PD3:
			tt, mode = _CTS, pin.Alt7
		case pin.PD4:
			tt, mode = _RTS, pin.Alt7
		case pin.PD5:
			tt, mode = _TX, pin.Alt7
		case pin.PD6:
			tt, mode = _RX, pin.Alt7
		case pin.PD7:
			tt, mode = _CK, pin.Alt7
		}
	case UART3:
		switch p {
		case pin.PB10:
			tt, mode = _TX, pin.Alt7
		case pin.PB11:
			tt, mode = _RX, pin.Alt7
		case pin.PB12:
			tt, mode = _CK, pin.Alt7
		case pin.PB13:
			tt, mode = _CTS, pin.Alt7
		case pin.PB14:
			tt, mode = _RTS, pin.Alt7
		case pin.PC10:
			tt, mode = _TX, pin.Alt7
		case pin.PC11:
			tt, mode = _RX, pin.Alt7
		case pin.PC12:
			tt, mode = _CK, pin.Alt7
		case pin.PD8:
			tt, mode = _TX, pin.Alt7
		case pin.PD9:
			tt, mode = _RX, pin.Alt7
		case pin.PD10:
			tt, mode = _CK, pin.Alt7
		case pin.PD11:
			tt, mode = _CTS, pin.Alt7
		case pin.PD12:
			tt, mode = _RTS, pin.Alt7
		}
	case UART4:
		switch p {
		case pin.PA0:
			tt, mode = _TX, pin.Alt8
		case pin.PA1:
			tt, mode = _RX, pin.Alt8
		case pin.PA11:
			tt, mode = _RX, pin.Alt6
		case pin.PA12:
			tt, mode = _TX, pin.Alt6
		case pin.PA14:
			tt, mode = _RTS, pin.Alt8
		case pin.PB0:
			tt, mode = _CTS, pin.Alt8
		case pin.PB8:
			tt, mode = _RX, pin.Alt8
		case pin.PB9:
			tt, mode = _TX, pin.Alt8
		case pin.PB14:
			tt, mode = _RTS, pin.Alt8
		case pin.PB15:
			tt, mode = _CTS, pin.Alt8
		case pin.PC10:
			tt, mode = _TX, pin.Alt8
		case pin.PC11:
			tt, mode = _RX, pin.Alt8
		case pin.PD0:
			tt, mode = _RX, pin.Alt8
		case pin.PD1:
			tt, mode = _TX, pin.Alt8
		}
	case UART5:
		switch p {
		case pin.PB5:
			tt, mode = _RX, pin.Alt14
		case pin.PB6:
			tt, mode = _TX, pin.Alt14
		case pin.PB12:
			tt, mode = _RX, pin.Alt14
		case pin.PB13:
			tt, mode = _TX, pin.Alt14
		case pin.PC8:
			tt, mode = _RTS, pin.Alt8
		case pin.PC9:
			tt, mode = _CTS, pin.Alt8
		case pin.PC12:
			tt, mode = _TX, pin.Alt8
		case pin.PD2:
			tt, mode = _RX, pin.Alt8
		}
	case UART6:
		switch p {
		case pin.PC6:
			tt, mode = _TX, pin.Alt7
		case pin.PC7:
			tt, mode = _RX, pin.Alt7
		case pin.PC8:
			tt, mode = _CK, pin.Alt7
		case pin.PG7:
			tt, mode = _CK, pin.Alt7
		case pin.PG8:
			tt, mode = _RTS, pin.Alt7
		case pin.PG9:
			tt, mode = _RX, pin.Alt7
		case pin.PG12:
			tt, mode = _RTS, pin.Alt7
		case pin.PG13:
			tt, mode = _CTS, pin.Alt7
		case pin.PG14:
			tt, mode = _TX, pin.Alt7
		case pin.PG15:
			tt, mode = _CTS, pin.Alt7
		}
	case UART7:
		switch p {
		case pin.PA8:
			tt, mode = _RX, pin.Alt11
		case pin.PA15:
			tt, mode = _TX, pin.Alt11
		case pin.PB3:
			tt, mode = _RX, pin.Alt11
		case pin.PB4:
			tt, mode = _TX, pin.Alt11
		case pin.PE7:
			tt, mode = _RX, pin.Alt7
		case pin.PE8:
			tt, mode = _TX, pin.Alt7
		case pin.PE9:
			tt, mode = _RTS, pin.Alt7
		case pin.PE10:
			tt, mode = _CTS, pin.Alt7
		case pin.PF6:
			tt, mode = _RX, pin.Alt7
		case pin.PF7:
			tt, mode = _TX, pin.Alt7
		case pin.PF8:
			tt, mode = _RTS, pin.Alt7
		case pin.PF9:
			tt, mode = _CTS, pin.Alt7
		}
	case UART8:
		switch p {
		case pin.PD14:
			tt, mode = _CTS, pin.Alt8
		case pin.PD15:
			tt, mode = _RTS, pin.Alt8
		case pin.PE0:
			tt, mode = _RX, pin.Alt8
		case pin.PE1:
			tt, mode = _TX, pin.Alt8
		}
	}

	if mode == 0 || tt != t {
		err = corehal.ErrInvalidPinout
	}
	return
}

func (u _uart) Configure(cfg Config) error {
	p := usart.Instances[u]

	// Disable the peripheral first.
	p.Cr1.SetUe(false)
	for p.Cr1.GetUe() {
	}

	var altTX, altRX, altCTS, altRTS, altCK corepin.Mode
	var err error

	if cfg.TX != 0 {
		altTX, err = altFunction(cfg.TX, _TX, u)
		if err != nil {
			return err
		}
	}

	if cfg.RX != 0 {
		altRX, err = altFunction(cfg.RX, _RX, u)
		if err != nil {
			return err
		}
	}

	if cfg.CTS != 0 {
		altCTS, err = altFunction(cfg.CTS, _CTS, u)
		if err != nil {
			return err
		}
	}

	if cfg.RTS != 0 {
		altRTS, err = altFunction(cfg.RX, _RTS, u)
		if err != nil {
			return err
		}
	}

	if cfg.CK != 0 {
		altCK, err = altFunction(cfg.CK, _CK, u)
		if err != nil {
			return err
		}
	}

	// Enable/disable the respective USART peripheral clock.
	switch u {
	case 0:
		rcc.Rcc.Apb2enr.SetUsart1en(cfg.Enable)
	case 1:
		rcc.Rcc.Apb1lenr.SetUsart2en(cfg.Enable)
	case 2:
		rcc.Rcc.Apb1lenr.SetUsart3en(cfg.Enable)
	case 3:
		rcc.Rcc.Apb1lenr.SetUart4en(cfg.Enable)
	case 4:
		rcc.Rcc.Apb1lenr.SetUart5en(cfg.Enable)
	case 5:
		rcc.Rcc.Apb2enr.SetUsart6en(cfg.Enable)
	case 6:
		rcc.Rcc.Apb1lenr.SetUsart7en(cfg.Enable)
	case 7:
		rcc.Rcc.Apb1lenr.SetUsart8en(cfg.Enable)
	}

	if !cfg.Enable {
		return nil
	}

	// Configure the GPIO pins.
	cfg.TX.SetMode(pin.AltFunction | altTX)
	cfg.RX.SetMode(pin.AltFunction | altRX)
	cfg.CTS.SetMode(pin.AltFunction | altCTS)
	cfg.RTS.SetMode(pin.AltFunction | altRTS)
	cfg.CK.SetMode(pin.AltFunction | altCK)

	// Set word length.
	switch cfg.CharacterSize {
	case 7:
		p.Cr1.SetM0(false)
		p.Cr1.SetM1(true)
	case 8:
		p.Cr1.SetM0(false)
		p.Cr1.SetM1(false)
	case 9:
		p.Cr1.SetM0(true)
		p.Cr1.SetM1(false)
	}

	// Derive the clock divider to achieve the specified baud rate.
	var sourceClockHz uint64
	switch u {
	case UART1, UART6:
		sourceClockHz = hal.Usart16SourceFrequencyHz
	case UART2, UART3, UART4, UART5, UART7, UART8:
		sourceClockHz = hal.Usart234578SourceFrequencyHz
	}
	div := sourceClockHz / uint64(cfg.Baud)

	// Set the baud rate.
	volatile.StoreUint32((*uint32)(&p.Brr), uint32(div))

	// Set the number of stop bits.
	p.Cr2.SetStop(cfg.NumStopBits)

	// Enable/disable the transmitter and receiver.
	p.Cr1.SetTe(cfg.TransmitEnabled)
	p.Cr1.SetRe(cfg.ReceiveEnabled)

	// Enable the FIFO.
	p.Cr1.SetFifoen(true)

	// Enable the USART.
	p.Cr1.SetUe(true)
	for !p.Cr1.GetUe() {
	}

	return nil
}

func (u _uart) Read(s []byte) (n int, err error) {
	mutex := &mutex[u]
	mutex.Lock()

	p := usart.Instances[u]
	remaining := len(s)

	if remaining == 0 {
		mutex.Unlock()
		return 0, nil
	}

	// Read as many bytes as currently available
	for remaining > 0 && p.Isr.GetRxne() {
		s[n] = byte(p.Rdr.GetRdr())
		n++
		remaining--
	}

	mutex.Unlock()
	return n, nil
}

func (u _uart) Write(s []byte) (n int, err error) {
	mutex := &mutex[u]
	mutex.Lock()

	p := usart.Instances[u]
	remaining := len(s)

	if remaining == 0 {
		mutex.Unlock()
		return 0, nil
	}

	for remaining > 0 {
		// Fill the FIFO while TXE is set.
		for remaining > 0 && p.Isr.GetTxe() {
			// Write the next character into TDR.
			p.Tdr.SetTdr(uint16(s[n]) & 0x1FF)
			n++
			remaining--
		}
	}

	// Wait for the TC flag.
	deadline := time.Now().Add(defaultTimeout)
	for !p.Isr.GetTc() {
		if time.Now().After(deadline) {
			mutex.Unlock()
			return n, errTimeout
		}
		runtime.Gosched()
	}

	mutex.Unlock()
	return n, nil
}

func (u _uart) WriteString(s string) (n int, err error) {
	return u.Write([]byte(s))
}
