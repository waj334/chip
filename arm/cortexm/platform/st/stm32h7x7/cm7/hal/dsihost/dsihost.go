//go:build stm32h7x7

package dsihost

import (
	"pkg.si-go.dev/chip/arm/cortexm"
	regdsi "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/reg/dsihost"
	corehal "pkg.si-go.dev/chip/core/hal"
)

const (
	defaultPollLimit = 5_000_000
)

// HSEFrequencyHz must match the HSE clock used by the DSI wrapper PLL.
// Used to compute the UIX4 bit period for the D-PHY (ST HAL_DSI_Init logic).
const HSEFrequencyHz = 16_000_000

var (
	DSIHOST DSIHost
)

type DSIHost struct{}

type LaneCount uint8

const (
	OneDataLane  LaneCount = 1
	TwoDataLanes LaneCount = 2
)

type ColorCoding uint8

const (
	ColorRGB565 ColorCoding = 0x0
	ColorRGB666 ColorCoding = 0x3
	ColorRGB888 ColorCoding = 0x5
)

type ColorMux uint8

const (
	ColorMux16Bit1 = 0x0 // 16-bit configuration 1
	ColorMux16Bit2 = 0x1 // 16-bit configuration 2
	ColorMux16Bit3 = 0x2 // 16-bit configuration 3
	ColorMux18Bit1 = 0x3 // 18-bit configuration 1
	ColorMux18Bit2 = 0x4 // 18-bit configuration 2
	ColorMux24Bit  = 0x5 // 24-bit
)

type VideoMode uint8

const (
	VideoModeNonBurstSyncPulse VideoMode = 0
	VideoModeNonBurstSyncEvent VideoMode = 1
	VideoModeBurst             VideoMode = 2
)

type PLLIDF uint8

const (
	PLLIDF1 PLLIDF = 1
	PLLIDF2 PLLIDF = 2
	PLLIDF3 PLLIDF = 3
	PLLIDF4 PLLIDF = 4
	PLLIDF5 PLLIDF = 5
	PLLIDF6 PLLIDF = 6
	PLLIDF7 PLLIDF = 7
)

// PLLODF encoding: 0 = /1, 1 = /2, 2 = /4, 3 = /8.
type PLLODF uint8

const (
	PLLODF1 PLLODF = 0
	PLLODF2 PLLODF = 1
	PLLODF4 PLLODF = 2
	PLLODF8 PLLODF = 3
)

type PLLConfig struct {
	Enable bool

	NDIV uint8
	IDF  PLLIDF
	ODF  PLLODF

	RegulatorEnable bool
}

type Config struct {
	Enable bool
	PLL    PLLConfig
	Lanes  LaneCount

	// Must be >= 2. For 62.5 MHz lane-byte clock, 4 gives ~15.6 MHz escape clock.
	TXEscapeCkdiv uint8

	AutomaticClockLaneControl bool

	VirtualChannelID uint8
	ColorCoding      ColorCoding
	ColorMux         ColorMux
	LooselyPacked    bool

	// Internal LTDC<->DSI bridge polarities. Should match the panel.
	// Note: LTDC side polarities are INVERTED by the bridge for DE (see ST HAL
	// HAL_LTDCEx_StructInitFromVideoConfig). Set these to match what the
	// physical panel expects on its DSI link.
	HSPolarityHigh bool
	VSPolarityHigh bool
	DEPolarityHigh bool

	VideoMode VideoMode

	PacketSize     uint16
	NumberOfChunks uint16
	NullPacketSize uint16

	LaneByteClockHz uint32
	PixelClockHz    uint32

	// DSI video horizontal timing values are in lane-byte clock cycles.
	HorizontalSyncActive uint16
	HorizontalBackPorch  uint16
	HorizontalLine       uint16

	// DSI video vertical timing values are in lines.
	VerticalSyncActive uint16
	VerticalBackPorch  uint16
	VerticalFrontPorch uint16
	VerticalActive     uint16

	// LP-during-video controls. Enabling these allows sending LP commands
	// during the indicated periods of the active video stream — this is how
	// panel init commands are delivered without stopping the video stream
	// (matches ST HAL / Arduino behavior).
	LPCommandEnable bool
	LPHFPEnable     bool
	LPHBPEnable     bool
	LPVACTEnable    bool
	LPVFPEnable     bool
	LPVBPEnable     bool
	LPVSAEnable     bool

	FrameBTAAcknowledge bool

	LPLargestPacketSize     uint8
	LPVACTLargestPacketSize uint8

	Timeouts  Timeouts
	PhyTiming PhyTiming

	// LPCommandMode controls whether DCS/generic command packets are sent in
	// LP mode (true) or HS mode (false). For panel init this must be true.
	LPCommandMode bool
}

type Timeouts struct {
	LowPowerRX     uint16
	HighSpeedTX    uint16
	HighSpeedRead  uint16
	LowPowerRead   uint16
	HighSpeedWrite uint16
	LowPowerWrite  uint16
	BTATimeout     uint16
}

type PhyTiming struct {
	Enable bool

	ClockLP2HS uint16
	ClockHS2LP uint16

	DataLP2HS uint8
	DataHS2LP uint8

	MaxReadTime  uint16
	StopWaitTime uint8
}

type PacketDataType uint8

const (
	GenericShortWrite0 PacketDataType = 0x03
	GenericShortWrite1 PacketDataType = 0x13
	GenericShortWrite2 PacketDataType = 0x23
	GenericLongWrite   PacketDataType = 0x29

	DCSShortWrite0 PacketDataType = 0x05
	DCSShortWrite1 PacketDataType = 0x15
	DCSLongWrite   PacketDataType = 0x39

	DCSRead             PacketDataType = 0x06
	GenericRead0        PacketDataType = 0x04
	GenericRead1        PacketDataType = 0x14
	GenericRead2        PacketDataType = 0x24
	SetMaxReturnPktSize PacketDataType = 0x37
)

// Configure performs the equivalent of ST HAL_DSI_Init + HAL_DSI_ConfigVideoMode
// + HAL_DSI_ConfigPhyTimer. At the end of Configure, both the DSI host (DSI_CR.EN)
// and the wrapper (DSI_WCR.DSIEN/LTDCEN) are DISABLED. Call Start() and
// EnableLTDC() to begin operation.
func (d DSIHost) Configure(cfg Config) error {
	if !cfg.Enable {
		d.Disable()
		return nil
	}

	if err := validateConfig(cfg); err != nil {
		return err
	}

	enableClock()
	resetPeripheral()

	// ===== HAL_DSI_DeInit equivalent =====
	// Quiet state before configuring.
	disableHostAndWrapper()
	clearWrapperInterruptFlags()

	// ===== HAL_DSI_Init body =====

	// Turn on regulator, wait for RRS.
	if err := configurePLL(cfg.PLL); err != nil {
		return err
	}

	// Briefly enable the host so the PHY can come up.
	regdsi.Dsihost.Dsicr.SetEn(true)
	cortexm.DSB()

	// Set TX escape clock divider.
	configureClocking(cfg)

	// D-PHY digital + clock enable. ST HAL order is DEN then CKE.
	regdsi.Dsihost.Dsipctlr.SetDen(true)
	regdsi.Dsihost.Dsipctlr.SetCke(true)
	cortexm.DSB()

	// Number of data lanes.
	// Hardware encoding: NL=0 means 1 data lane, NL=1 means 2 data lanes.
	regdsi.Dsihost.Dsipconfr.SetNl(uint8(cfg.Lanes - 1))

	// Wait for D-PHY to reach stop state on the relevant lanes.
	if err := waitPHYStopState(cfg.Lanes); err != nil {
		return err
	}

	// Compute and set UIX4 — bit period in HS mode in 0.25 ns units.
	// ST HAL formula:
	//   UIX4 = (4_000_000 * IDF * (1 << ODF)) / ((HSE_MHz * 1000) * NDIV)
	if err := setUIX4(cfg.PLL); err != nil {
		return err
	}

	// Disable all error interrupts during init.
	clearInterruptEnables()

	// HAL_DSI_Init then disables EN. Match that.
	regdsi.Dsihost.Dsicr.SetEn(false)
	cortexm.DSB()

	// Clock lane: DPCC=1 (clock lane data lane provided by PHY), ACR=auto.
	regdsi.Dsihost.Dsiclcr.SetDpcc(true)
	regdsi.Dsihost.Dsiclcr.SetAcr(cfg.AutomaticClockLaneControl)

	// PHY stop-wait time.
	regdsi.Dsihost.Dsipconfr.SetSwtime(cfg.PhyTiming.StopWaitTime)

	// ===== HAL_DSI_ConfigVideoMode =====
	configureVideoMode(cfg)

	// ===== HAL_DSI_ConfigPhyTimer =====
	configurePHYTimer(cfg)

	// Protocol config (flow control bits in PCR), VC IDs.
	configureProtocol(cfg)

	// Timeouts (TCCR0..TCCR5).
	configureTimeouts(cfg.Timeouts)

	// Command transmission mode (LP vs HS for each packet type).
	configureCommandMode(cfg)

	return nil
}

// Start enables the DSI host and wrapper. This is the equivalent of
// ST HAL_DSI_Start: __HAL_DSI_ENABLE + __HAL_DSI_WRAPPER_ENABLE.
// LTDC streaming is NOT enabled here — call EnableLTDC for that.
func (d DSIHost) Start() error {
	regdsi.Dsihost.Dsicr.SetEn(true)
	cortexm.DSB()

	regdsi.Dsihost.Dsiwcr.SetDsien(true)
	cortexm.DSB()
	return nil
}

// Stop disables wrapper and host. Equivalent to ST HAL_DSI_Stop.
func (d DSIHost) Stop() {
	regdsi.Dsihost.Dsiwcr.SetLtdcen(false)
	regdsi.Dsihost.Dsiwcr.SetDsien(false)
	regdsi.Dsihost.Dsicr.SetEn(false)
	cortexm.DSB()
}

// Disable powers everything down: wrapper, host, PHY, PLL, regulator.
func (d DSIHost) Disable() {
	enableClock()
	disableHostAndWrapper()
	regdsi.Dsihost.Dsiwrpcr.SetPllen(false)
	regdsi.Dsihost.Dsiwrpcr.SetRegen(false)
	cortexm.DSB()
}

func (d DSIHost) Status0() uint32 {
	return regdsi.Dsihost.Dsiisr0.Load()
}

func (d DSIHost) Status1() uint32 {
	return regdsi.Dsihost.Dsiisr1.Load()
}

func (d DSIHost) WrapperStatus() uint32 {
	return regdsi.Dsihost.Dsiwisr.Load()
}

func (d DSIHost) PHYStatus() uint32 {
	return regdsi.Dsihost.Dsipsr.Load()
}

func (d DSIHost) ClearWrapperInterrupts() {
	clearWrapperInterruptFlags()
}

func (d DSIHost) ClearHostInterrupts() {
}

func (d DSIHost) EnableLTDC() {
	regdsi.Dsihost.Dsiwcr.SetLtdcen(true)
	cortexm.DSB()
}

func (d DSIHost) DisableLTDC() {
	regdsi.Dsihost.Dsiwcr.SetLtdcen(false)
	cortexm.DSB()
}

// ===== Packet send =====

func (d DSIHost) DCSShortWrite(channelID uint8, cmd uint8) error {
	return d.shortWrite(DCSShortWrite0, channelID, cmd, 0)
}

func (d DSIHost) DCSShortWriteParam(channelID uint8, cmd uint8, param uint8) error {
	return d.shortWrite(DCSShortWrite1, channelID, cmd, param)
}

// DCSLongWrite sends a DCS long packet. The first payload byte is the DCS
// command, followed by parameters. Matches ST HAL_DSI_LongWrite layout.
func (d DSIHost) DCSLongWrite(channelID uint8, cmd uint8, data []byte) error {
	return d.longWrite(DCSLongWrite, channelID, cmd, data)
}

func (d DSIHost) GenericShortWrite0(channelID uint8) error {
	return d.shortWrite(GenericShortWrite0, channelID, 0, 0)
}

func (d DSIHost) GenericShortWrite1(channelID uint8, p0 uint8) error {
	return d.shortWrite(GenericShortWrite1, channelID, p0, 0)
}

// GenericShortWrite1P sends DT=0x13 with two data bytes, matching ST HAL
// DSI_GEN_SHORT_PKT_WRITE_P1 semantics used by Arduino's
// Generic_Short_Write_1P.
func (d DSIHost) GenericShortWrite1P(channelID uint8, p0 uint8, p1 uint8) error {
	return d.shortWrite(GenericShortWrite1, channelID, p0, p1)
}

func (d DSIHost) GenericShortWrite2(channelID uint8, p0 uint8, p1 uint8) error {
	return d.shortWrite(GenericShortWrite2, channelID, p0, p1)
}

// GenericLongWrite sends a generic long packet. data is the full payload.
func (d DSIHost) GenericLongWrite(channelID uint8, data []byte) error {
	if len(data) == 0 {
		return corehal.ErrInvalidConfig
	}
	return d.longWrite(GenericLongWrite, channelID, data[0], data[1:])
}

func (d DSIHost) shortWrite(dt PacketDataType, channelID uint8, data0 uint8, data1 uint8) error {
	return writeHeaderWithParams(dt, channelID, data0, data1)
}

// longWrite layout matches ST HAL_DSI_LongWrite:
//   - First GPDR word: byte0=cmd, byte1=data[0], byte2=data[1], byte3=data[2]
//   - Subsequent words: 4 bytes of data per word
//   - WCount = len(data) + 1 (including the cmd byte)
//
// ST HAL waits for CMDFE at the start (prior command fully drained) then writes
// without per-word stall. We keep per-word PWRFE checks for extra safety.
func (d DSIHost) longWrite(dt PacketDataType, channelID uint8, cmd uint8, data []byte) error {
	totalBytes := len(data) + 1
	if totalBytes > 0xffff {
		return corehal.ErrInvalidConfig
	}

	// Match ST HAL: wait for prior command header to drain.
	if err := waitCommandEmpty(); err != nil {
		return err
	}

	var word [4]byte
	word[0] = cmd
	count := 1
	dataIdx := 0

	for count < 4 && dataIdx < len(data) {
		word[count] = data[dataIdx]
		count++
		dataIdx++
	}

	if err := waitPayloadWritable(); err != nil {
		return err
	}
	writePayloadWord(word[:])

	for dataIdx < len(data) {
		word = [4]byte{}
		count = 0

		for count < 4 && dataIdx < len(data) {
			word[count] = data[dataIdx]
			count++
			dataIdx++
		}

		if err := waitPayloadWritable(); err != nil {
			return err
		}
		writePayloadWord(word[:])
	}

	// Launch packet.
	return writeHeader(dt, channelID, uint16(totalBytes))
}

// DCSRead issues a DCS read (DT 0x06) for cmd and fills dst with up to len(dst)
// bytes returned by the panel. Requires DSI_PCR.BTAE=1 so the host can turn the
// bus around. Returns the number of bytes actually read.
func (d DSIHost) DCSRead(channelID uint8, cmd uint8, dst []byte) (int, error) {
	return d.read(DCSRead, channelID, cmd, 0, dst)
}

func (d DSIHost) read(dt PacketDataType, channelID, data0, data1 uint8, dst []byte) (int, error) {
	if channelID > 3 {
		return 0, corehal.ErrInvalidConfig
	}
	if len(dst) == 0 {
		return 0, nil
	}

	// Reads longer than 2 bytes must announce the max return size first.
	if len(dst) > 2 {
		if err := writeHeaderWithParams(
			SetMaxReturnPktSize, channelID,
			uint8(len(dst)&0xff), uint8(len(dst)>>8),
		); err != nil {
			return 0, err
		}
	}

	// Issue the read. writeHeaderWithParams already waits for the command FIFO
	// to drain and DSBs after the store, so the BTA is launched here.
	if err := writeHeaderWithParams(dt, channelID, data0, data1); err != nil {
		return 0, err
	}

	// Block until the read payload FIFO has data (PRDFE clears when non-empty).
	if err := waitReadNotEmpty(); err != nil {
		return 0, err
	}

	// Drain whole 32-bit words; the host has already stripped the response header.
	n := 0
	for n < len(dst) {
		if regdsi.Dsihost.Dsigpsr.GetPrdfe() {
			break
		}
		word := regdsi.Dsihost.Dsigpdr.Load()
		for shift := uint(0); shift < 32 && n < len(dst); shift += 8 {
			dst[n] = uint8(word >> shift)
			n++
		}
	}
	return n, nil
}

func waitReadNotEmpty() error {
	for i := uint32(0); i < defaultPollLimit; i++ {
		if !regdsi.Dsihost.Dsigpsr.GetPrdfe() {
			return nil
		}
	}
	return corehal.ErrTimeout
}
