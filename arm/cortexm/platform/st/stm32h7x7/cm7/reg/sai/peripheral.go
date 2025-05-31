//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package sai

import (
	"unsafe"
	"volatile"
)

var (
	Sai1 = (*_sai)(unsafe.Pointer(uintptr(0x40015800)))
	Sai2 = (*_sai)(unsafe.Pointer(uintptr(0x40015c00)))
	Sai3 = (*_sai)(unsafe.Pointer(uintptr(0x40016000)))
	Sai4 = (*_sai)(unsafe.Pointer(uintptr(0x58005400)))

	Instances = [4]*_sai{
		Sai1,
		Sai2,
		Sai3,
		Sai4,
	}
)

type _sai struct {
	Saigcr    registerSaigcrType
	Saiacr1   registerSaiacr1Type
	Saiacr2   registerSaiacr2Type
	Saiafrcr  registerSaiafrcrType
	Saiaslotr registerSaiaslotrType
	Saiaim    registerSaiaimType
	Saiasr    registerSaiasrType
	Saiaclrfr registerSaiaclrfrType
	Saiadr    registerSaiadrType
	Saibcr1   registerSaibcr1Type
	Saibcr2   registerSaibcr2Type
	Saibfrcr  registerSaibfrcrType
	Saibslotr registerSaibslotrType
	Saibim    registerSaibimType
	Saibsr    registerSaibsrType
	Saibclrfr registerSaibclrfrType
	Saibdr    registerSaibdrType
	Saipdmcr  registerSaipdmcrType
	Saipdmdly registerSaipdmdlyType
}

// registerSaigcrType Global configuration register
type registerSaigcrType uint32

const (
	RegisterSaigcrFieldSyncinShift = 0
	RegisterSaigcrFieldSyncinMask  = 0x3
)

// GetSyncin Synchronization inputs
func (r *registerSaigcrType) GetSyncin() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaigcrFieldSyncinMask) >> RegisterSaigcrFieldSyncinShift)
}

// SetSyncin Synchronization inputs
func (r *registerSaigcrType) SetSyncin(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaigcrFieldSyncinMask)|(uint32(value)<<RegisterSaigcrFieldSyncinShift))
}

const (
	RegisterSaigcrFieldSyncoutShift = 4
	RegisterSaigcrFieldSyncoutMask  = 0x30
)

// GetSyncout Synchronization outputs These bits are set and cleared by software.
func (r *registerSaigcrType) GetSyncout() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaigcrFieldSyncoutMask) >> RegisterSaigcrFieldSyncoutShift)
}

// SetSyncout Synchronization outputs These bits are set and cleared by software.
func (r *registerSaigcrType) SetSyncout(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaigcrFieldSyncoutMask)|(uint32(value)<<RegisterSaigcrFieldSyncoutShift))
}

// registerSaiacr1Type Configuration register 1
type registerSaiacr1Type uint32

const (
	RegisterSaiacr1FieldModeShift = 0
	RegisterSaiacr1FieldModeMask  = 0x3
)

// GetMode SAIx audio block mode immediately
func (r *registerSaiacr1Type) GetMode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldModeMask) >> RegisterSaiacr1FieldModeShift)
}

// SetMode SAIx audio block mode immediately
func (r *registerSaiacr1Type) SetMode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldModeMask)|(uint32(value)<<RegisterSaiacr1FieldModeShift))
}

const (
	RegisterSaiacr1FieldPrtcfgShift = 2
	RegisterSaiacr1FieldPrtcfgMask  = 0xc
)

// GetPrtcfg Protocol configuration. These bits are set and cleared by software. These bits have to be configured when the audio block is disabled.
func (r *registerSaiacr1Type) GetPrtcfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldPrtcfgMask) >> RegisterSaiacr1FieldPrtcfgShift)
}

// SetPrtcfg Protocol configuration. These bits are set and cleared by software. These bits have to be configured when the audio block is disabled.
func (r *registerSaiacr1Type) SetPrtcfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldPrtcfgMask)|(uint32(value)<<RegisterSaiacr1FieldPrtcfgShift))
}

const (
	RegisterSaiacr1FieldDsShift = 5
	RegisterSaiacr1FieldDsMask  = 0xe0
)

// GetDs Data size. These bits are set and cleared by software. These bits are ignored when the SPDIF protocols are selected (bit PRTCFG[1:0]), because the frame and the data size are fixed in such case. When the companding mode is selected through COMP[1:0] bits, DS[1:0] are ignored since the data size is fixed to 8 bits by the algorithm. These bits must be configured when the audio block is disabled.
func (r *registerSaiacr1Type) GetDs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldDsMask) >> RegisterSaiacr1FieldDsShift)
}

// SetDs Data size. These bits are set and cleared by software. These bits are ignored when the SPDIF protocols are selected (bit PRTCFG[1:0]), because the frame and the data size are fixed in such case. When the companding mode is selected through COMP[1:0] bits, DS[1:0] are ignored since the data size is fixed to 8 bits by the algorithm. These bits must be configured when the audio block is disabled.
func (r *registerSaiacr1Type) SetDs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldDsMask)|(uint32(value)<<RegisterSaiacr1FieldDsShift))
}

const (
	RegisterSaiacr1FieldLsbfirstShift = 8
	RegisterSaiacr1FieldLsbfirstMask  = 0x100
)

// GetLsbfirst Least significant bit first. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in AC97 audio protocol since AC97 data are always transferred with the MSB first. This bit has no meaning in SPDIF audio protocol since in SPDIF data are always transferred with LSB first.
func (r *registerSaiacr1Type) GetLsbfirst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldLsbfirstMask) != 0
}

// SetLsbfirst Least significant bit first. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in AC97 audio protocol since AC97 data are always transferred with the MSB first. This bit has no meaning in SPDIF audio protocol since in SPDIF data are always transferred with LSB first.
func (r *registerSaiacr1Type) SetLsbfirst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr1FieldLsbfirstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldLsbfirstMask)
	}
}

const (
	RegisterSaiacr1FieldCkstrShift = 9
	RegisterSaiacr1FieldCkstrMask  = 0x200
)

// GetCkstr Clock strobing edge. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in SPDIF audio protocol.
func (r *registerSaiacr1Type) GetCkstr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldCkstrMask) != 0
}

// SetCkstr Clock strobing edge. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in SPDIF audio protocol.
func (r *registerSaiacr1Type) SetCkstr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr1FieldCkstrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldCkstrMask)
	}
}

const (
	RegisterSaiacr1FieldSyncenShift = 10
	RegisterSaiacr1FieldSyncenMask  = 0xc00
)

// GetSyncen Synchronization enable. These bits are set and cleared by software. They must be configured when the audio sub-block is disabled. Note: The audio sub-block should be configured as asynchronous when SPDIF mode is enabled.
func (r *registerSaiacr1Type) GetSyncen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldSyncenMask) >> RegisterSaiacr1FieldSyncenShift)
}

// SetSyncen Synchronization enable. These bits are set and cleared by software. They must be configured when the audio sub-block is disabled. Note: The audio sub-block should be configured as asynchronous when SPDIF mode is enabled.
func (r *registerSaiacr1Type) SetSyncen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldSyncenMask)|(uint32(value)<<RegisterSaiacr1FieldSyncenShift))
}

const (
	RegisterSaiacr1FieldMonoShift = 12
	RegisterSaiacr1FieldMonoMask  = 0x1000
)

// GetMono Mono mode. This bit is set and cleared by software. It is meaningful only when the number of slots is equal to 2. When the mono mode is selected, slot 0 data are duplicated on slot 1 when the audio block operates as a transmitter. In reception mode, the slot1 is discarded and only the data received from slot 0 are stored. Refer to Section: Mono/stereo mode for more details.
func (r *registerSaiacr1Type) GetMono() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldMonoMask) != 0
}

// SetMono Mono mode. This bit is set and cleared by software. It is meaningful only when the number of slots is equal to 2. When the mono mode is selected, slot 0 data are duplicated on slot 1 when the audio block operates as a transmitter. In reception mode, the slot1 is discarded and only the data received from slot 0 are stored. Refer to Section: Mono/stereo mode for more details.
func (r *registerSaiacr1Type) SetMono(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr1FieldMonoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldMonoMask)
	}
}

const (
	RegisterSaiacr1FieldOutdrivShift = 13
	RegisterSaiacr1FieldOutdrivMask  = 0x2000
)

// GetOutdriv Output drive. This bit is set and cleared by software. Note: This bit has to be set before enabling the audio block and after the audio block configuration.
func (r *registerSaiacr1Type) GetOutdriv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldOutdrivMask) != 0
}

// SetOutdriv Output drive. This bit is set and cleared by software. Note: This bit has to be set before enabling the audio block and after the audio block configuration.
func (r *registerSaiacr1Type) SetOutdriv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr1FieldOutdrivMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldOutdrivMask)
	}
}

const (
	RegisterSaiacr1FieldSaixenShift = 16
	RegisterSaiacr1FieldSaixenMask  = 0x10000
)

// GetSaixen Audio block enable where x is A or B. This bit is set by software. To switch off the audio block, the application software must program this bit to 0 and poll the bit till it reads back 0, meaning that the block is completely disabled. Before setting this bit to 1, check that it is set to 0, otherwise the enable command will not be taken into account. This bit allows to control the state of SAIx audio block. If it is disabled when an audio frame transfer is ongoing, the ongoing transfer completes and the cell is fully disabled at the end of this audio frame transfer. Note: When SAIx block is configured in master mode, the clock must be present on the input of SAIx before setting SAIXEN bit.
func (r *registerSaiacr1Type) GetSaixen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldSaixenMask) != 0
}

// SetSaixen Audio block enable where x is A or B. This bit is set by software. To switch off the audio block, the application software must program this bit to 0 and poll the bit till it reads back 0, meaning that the block is completely disabled. Before setting this bit to 1, check that it is set to 0, otherwise the enable command will not be taken into account. This bit allows to control the state of SAIx audio block. If it is disabled when an audio frame transfer is ongoing, the ongoing transfer completes and the cell is fully disabled at the end of this audio frame transfer. Note: When SAIx block is configured in master mode, the clock must be present on the input of SAIx before setting SAIXEN bit.
func (r *registerSaiacr1Type) SetSaixen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr1FieldSaixenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldSaixenMask)
	}
}

const (
	RegisterSaiacr1FieldDmaenShift = 17
	RegisterSaiacr1FieldDmaenMask  = 0x20000
)

// GetDmaen DMA enable. This bit is set and cleared by software. Note: Since the audio block defaults to operate as a transmitter after reset, the MODE[1:0] bits must be configured before setting DMAEN to avoid a DMA request in receiver mode.
func (r *registerSaiacr1Type) GetDmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldDmaenMask) != 0
}

// SetDmaen DMA enable. This bit is set and cleared by software. Note: Since the audio block defaults to operate as a transmitter after reset, the MODE[1:0] bits must be configured before setting DMAEN to avoid a DMA request in receiver mode.
func (r *registerSaiacr1Type) SetDmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr1FieldDmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldDmaenMask)
	}
}

const (
	RegisterSaiacr1FieldNomckShift = 19
	RegisterSaiacr1FieldNomckMask  = 0x80000
)

// GetNomck No divider
func (r *registerSaiacr1Type) GetNomck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldNomckMask) != 0
}

// SetNomck No divider
func (r *registerSaiacr1Type) SetNomck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr1FieldNomckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldNomckMask)
	}
}

const (
	RegisterSaiacr1FieldMckdivShift = 20
	RegisterSaiacr1FieldMckdivMask  = 0xf00000
)

// GetMckdiv Master clock divider. These bits are set and cleared by software. These bits are meaningless when the audio block operates in slave mode. They have to be configured when the audio block is disabled. Others: the master clock frequency is calculated accordingly to the following formula:
func (r *registerSaiacr1Type) GetMckdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldMckdivMask) >> RegisterSaiacr1FieldMckdivShift)
}

// SetMckdiv Master clock divider. These bits are set and cleared by software. These bits are meaningless when the audio block operates in slave mode. They have to be configured when the audio block is disabled. Others: the master clock frequency is calculated accordingly to the following formula:
func (r *registerSaiacr1Type) SetMckdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldMckdivMask)|(uint32(value)<<RegisterSaiacr1FieldMckdivShift))
}

const (
	RegisterSaiacr1FieldOsrShift = 26
	RegisterSaiacr1FieldOsrMask  = 0x4000000
)

// GetOsr Oversampling ratio for master clock
func (r *registerSaiacr1Type) GetOsr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr1FieldOsrMask) != 0
}

// SetOsr Oversampling ratio for master clock
func (r *registerSaiacr1Type) SetOsr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr1FieldOsrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr1FieldOsrMask)
	}
}

// registerSaiacr2Type Configuration register 2
type registerSaiacr2Type uint32

const (
	RegisterSaiacr2FieldFthShift = 0
	RegisterSaiacr2FieldFthMask  = 0x7
)

// GetFth FIFO threshold. This bit is set and cleared by software.
func (r *registerSaiacr2Type) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr2FieldFthMask) >> RegisterSaiacr2FieldFthShift)
}

// SetFth FIFO threshold. This bit is set and cleared by software.
func (r *registerSaiacr2Type) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr2FieldFthMask)|(uint32(value)<<RegisterSaiacr2FieldFthShift))
}

const (
	RegisterSaiacr2FieldFflushShift = 3
	RegisterSaiacr2FieldFflushMask  = 0x8
)

// SetFflush FIFO flush. This bit is set by software. It is always read as 0. This bit should be configured when the SAI is disabled.
func (r *registerSaiacr2Type) SetFflush(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr2FieldFflushMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr2FieldFflushMask)
	}
}

const (
	RegisterSaiacr2FieldTrisShift = 4
	RegisterSaiacr2FieldTrisMask  = 0x10
)

// GetTris Tristate management on data line. This bit is set and cleared by software. It is meaningful only if the audio block is configured as a transmitter. This bit is not used when the audio block is configured in SPDIF mode. It should be configured when SAI is disabled. Refer to Section: Output data line management on an inactive slot for more details.
func (r *registerSaiacr2Type) GetTris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr2FieldTrisMask) != 0
}

// SetTris Tristate management on data line. This bit is set and cleared by software. It is meaningful only if the audio block is configured as a transmitter. This bit is not used when the audio block is configured in SPDIF mode. It should be configured when SAI is disabled. Refer to Section: Output data line management on an inactive slot for more details.
func (r *registerSaiacr2Type) SetTris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr2FieldTrisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr2FieldTrisMask)
	}
}

const (
	RegisterSaiacr2FieldMuteShift = 5
	RegisterSaiacr2FieldMuteMask  = 0x20
)

// GetMute Mute. This bit is set and cleared by software. It is meaningful only when the audio block operates as a transmitter. The MUTE value is linked to value of MUTEVAL if the number of slots is lower or equal to 2, or equal to 0 if it is greater than 2. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSaiacr2Type) GetMute() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr2FieldMuteMask) != 0
}

// SetMute Mute. This bit is set and cleared by software. It is meaningful only when the audio block operates as a transmitter. The MUTE value is linked to value of MUTEVAL if the number of slots is lower or equal to 2, or equal to 0 if it is greater than 2. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSaiacr2Type) SetMute(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr2FieldMuteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr2FieldMuteMask)
	}
}

const (
	RegisterSaiacr2FieldMutevalShift = 6
	RegisterSaiacr2FieldMutevalMask  = 0x40
)

// GetMuteval Mute value. This bit is set and cleared by software.It must be written before enabling the audio block: SAIXEN. This bit is meaningful only when the audio block operates as a transmitter, the number of slots is lower or equal to 2 and the MUTE bit is set. If more slots are declared, the bit value sent during the transmission in mute mode is equal to 0, whatever the value of MUTEVAL. if the number of slot is lower or equal to 2 and MUTEVAL = 1, the MUTE value transmitted for each slot is the one sent during the previous frame. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSaiacr2Type) GetMuteval() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr2FieldMutevalMask) != 0
}

// SetMuteval Mute value. This bit is set and cleared by software.It must be written before enabling the audio block: SAIXEN. This bit is meaningful only when the audio block operates as a transmitter, the number of slots is lower or equal to 2 and the MUTE bit is set. If more slots are declared, the bit value sent during the transmission in mute mode is equal to 0, whatever the value of MUTEVAL. if the number of slot is lower or equal to 2 and MUTEVAL = 1, the MUTE value transmitted for each slot is the one sent during the previous frame. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSaiacr2Type) SetMuteval(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr2FieldMutevalMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr2FieldMutevalMask)
	}
}

const (
	RegisterSaiacr2FieldMutecntShift = 7
	RegisterSaiacr2FieldMutecntMask  = 0x1f80
)

// GetMutecnt Mute counter. These bits are set and cleared by software. They are used only in reception mode. The value set in these bits is compared to the number of consecutive mute frames detected in reception. When the number of mute frames is equal to this value, the flag MUTEDET will be set and an interrupt will be generated if bit MUTEDETIE is set. Refer to Section: Mute mode for more details.
func (r *registerSaiacr2Type) GetMutecnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr2FieldMutecntMask) >> RegisterSaiacr2FieldMutecntShift)
}

// SetMutecnt Mute counter. These bits are set and cleared by software. They are used only in reception mode. The value set in these bits is compared to the number of consecutive mute frames detected in reception. When the number of mute frames is equal to this value, the flag MUTEDET will be set and an interrupt will be generated if bit MUTEDETIE is set. Refer to Section: Mute mode for more details.
func (r *registerSaiacr2Type) SetMutecnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr2FieldMutecntMask)|(uint32(value)<<RegisterSaiacr2FieldMutecntShift))
}

const (
	RegisterSaiacr2FieldCplShift = 13
	RegisterSaiacr2FieldCplMask  = 0x2000
)

// GetCpl Complement bit. This bit is set and cleared by software. It defines the type of complement to be used for companding mode Note: This bit has effect only when the companding mode is -Law algorithm or A-Law algorithm.
func (r *registerSaiacr2Type) GetCpl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr2FieldCplMask) != 0
}

// SetCpl Complement bit. This bit is set and cleared by software. It defines the type of complement to be used for companding mode Note: This bit has effect only when the companding mode is -Law algorithm or A-Law algorithm.
func (r *registerSaiacr2Type) SetCpl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiacr2FieldCplMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr2FieldCplMask)
	}
}

const (
	RegisterSaiacr2FieldCompShift = 14
	RegisterSaiacr2FieldCompMask  = 0xc000
)

// GetComp Companding mode. These bits are set and cleared by software. The -Law and the A-Law log are a part of the CCITT G.711 recommendation, the type of complement that will be used depends on CPL bit. The data expansion or data compression are determined by the state of bit MODE[0]. The data compression is applied if the audio block is configured as a transmitter. The data expansion is automatically applied when the audio block is configured as a receiver. Refer to Section: Companding mode for more details. Note: Companding mode is applicable only when TDM is selected.
func (r *registerSaiacr2Type) GetComp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiacr2FieldCompMask) >> RegisterSaiacr2FieldCompShift)
}

// SetComp Companding mode. These bits are set and cleared by software. The -Law and the A-Law log are a part of the CCITT G.711 recommendation, the type of complement that will be used depends on CPL bit. The data expansion or data compression are determined by the state of bit MODE[0]. The data compression is applied if the audio block is configured as a transmitter. The data expansion is automatically applied when the audio block is configured as a receiver. Refer to Section: Companding mode for more details. Note: Companding mode is applicable only when TDM is selected.
func (r *registerSaiacr2Type) SetComp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiacr2FieldCompMask)|(uint32(value)<<RegisterSaiacr2FieldCompShift))
}

// registerSaiafrcrType This register has no meaning in AC97 and SPDIF audio protocol
type registerSaiafrcrType uint32

const (
	RegisterSaiafrcrFieldFrlShift = 0
	RegisterSaiafrcrFieldFrlMask  = 0xff
)

// GetFrl Frame length. These bits are set and cleared by software. They define the audio frame length expressed in number of SCK clock cycles: the number of bits in the frame is equal to FRL[7:0] + 1. The minimum number of bits to transfer in an audio frame must be equal to 8, otherwise the audio block will behaves in an unexpected way. This is the case when the data size is 8 bits and only one slot 0 is defined in NBSLOT[4:0] of SAI_xSLOTR register (NBSLOT[3:0] = 0000). In master mode, if the master clock (available on MCLK_x pin) is used, the frame length should be aligned with a number equal to a power of 2, ranging from 8 to 256. When the master clock is not used (NODIV = 1), it is recommended to program the frame length to an value ranging from 8 to 256. These bits are meaningless and are not used in AC97 or SPDIF audio block configuration.
func (r *registerSaiafrcrType) GetFrl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiafrcrFieldFrlMask) >> RegisterSaiafrcrFieldFrlShift)
}

// SetFrl Frame length. These bits are set and cleared by software. They define the audio frame length expressed in number of SCK clock cycles: the number of bits in the frame is equal to FRL[7:0] + 1. The minimum number of bits to transfer in an audio frame must be equal to 8, otherwise the audio block will behaves in an unexpected way. This is the case when the data size is 8 bits and only one slot 0 is defined in NBSLOT[4:0] of SAI_xSLOTR register (NBSLOT[3:0] = 0000). In master mode, if the master clock (available on MCLK_x pin) is used, the frame length should be aligned with a number equal to a power of 2, ranging from 8 to 256. When the master clock is not used (NODIV = 1), it is recommended to program the frame length to an value ranging from 8 to 256. These bits are meaningless and are not used in AC97 or SPDIF audio block configuration.
func (r *registerSaiafrcrType) SetFrl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiafrcrFieldFrlMask)|(uint32(value)<<RegisterSaiafrcrFieldFrlShift))
}

const (
	RegisterSaiafrcrFieldFsallShift = 8
	RegisterSaiafrcrFieldFsallMask  = 0x7f00
)

// GetFsall Frame synchronization active level length. These bits are set and cleared by software. They specify the length in number of bit clock (SCK) + 1 (FSALL[6:0] + 1) of the active level of the FS signal in the audio frame These bits are meaningless and are not used in AC97 or SPDIF audio block configuration. They must be configured when the audio block is disabled.
func (r *registerSaiafrcrType) GetFsall() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiafrcrFieldFsallMask) >> RegisterSaiafrcrFieldFsallShift)
}

// SetFsall Frame synchronization active level length. These bits are set and cleared by software. They specify the length in number of bit clock (SCK) + 1 (FSALL[6:0] + 1) of the active level of the FS signal in the audio frame These bits are meaningless and are not used in AC97 or SPDIF audio block configuration. They must be configured when the audio block is disabled.
func (r *registerSaiafrcrType) SetFsall(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiafrcrFieldFsallMask)|(uint32(value)<<RegisterSaiafrcrFieldFsallShift))
}

const (
	RegisterSaiafrcrFieldFsdefShift = 16
	RegisterSaiafrcrFieldFsdefMask  = 0x10000
)

// GetFsdef Frame synchronization definition. This bit is set and cleared by software. When the bit is set, the number of slots defined in the SAI_xSLOTR register has to be even. It means that half of this number of slots will be dedicated to the left channel and the other slots for the right channel (e.g: this bit has to be set for I2S or MSB/LSB-justified protocols...). This bit is meaningless and is not used in AC97 or SPDIF audio block configuration. It must be configured when the audio block is disabled.
func (r *registerSaiafrcrType) GetFsdef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiafrcrFieldFsdefMask) != 0
}

const (
	RegisterSaiafrcrFieldFspolShift = 17
	RegisterSaiafrcrFieldFspolMask  = 0x20000
)

// GetFspol Frame synchronization polarity. This bit is set and cleared by software. It is used to configure the level of the start of frame on the FS signal. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSaiafrcrType) GetFspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiafrcrFieldFspolMask) != 0
}

// SetFspol Frame synchronization polarity. This bit is set and cleared by software. It is used to configure the level of the start of frame on the FS signal. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSaiafrcrType) SetFspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiafrcrFieldFspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiafrcrFieldFspolMask)
	}
}

const (
	RegisterSaiafrcrFieldFsoffShift = 18
	RegisterSaiafrcrFieldFsoffMask  = 0x40000
)

// GetFsoff Frame synchronization offset. This bit is set and cleared by software. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSaiafrcrType) GetFsoff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiafrcrFieldFsoffMask) != 0
}

// SetFsoff Frame synchronization offset. This bit is set and cleared by software. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSaiafrcrType) SetFsoff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiafrcrFieldFsoffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiafrcrFieldFsoffMask)
	}
}

// registerSaiaslotrType This register has no meaning in AC97 and SPDIF audio protocol
type registerSaiaslotrType uint32

const (
	RegisterSaiaslotrFieldFboffShift = 0
	RegisterSaiaslotrFieldFboffMask  = 0x1f
)

// GetFboff First bit offset These bits are set and cleared by software. The value set in this bitfield defines the position of the first data transfer bit in the slot. It represents an offset value. In transmission mode, the bits outside the data field are forced to 0. In reception mode, the extra received bits are discarded. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaiaslotrType) GetFboff() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiaslotrFieldFboffMask) >> RegisterSaiaslotrFieldFboffShift)
}

// SetFboff First bit offset These bits are set and cleared by software. The value set in this bitfield defines the position of the first data transfer bit in the slot. It represents an offset value. In transmission mode, the bits outside the data field are forced to 0. In reception mode, the extra received bits are discarded. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaiaslotrType) SetFboff(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiaslotrFieldFboffMask)|(uint32(value)<<RegisterSaiaslotrFieldFboffShift))
}

const (
	RegisterSaiaslotrFieldSlotszShift = 6
	RegisterSaiaslotrFieldSlotszMask  = 0xc0
)

// GetSlotsz Slot size This bits is set and cleared by software. The slot size must be higher or equal to the data size. If this condition is not respected, the behavior of the SAI will be undetermined. Refer to Section: Output data line management on an inactive slot for information on how to drive SD line. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaiaslotrType) GetSlotsz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiaslotrFieldSlotszMask) >> RegisterSaiaslotrFieldSlotszShift)
}

// SetSlotsz Slot size This bits is set and cleared by software. The slot size must be higher or equal to the data size. If this condition is not respected, the behavior of the SAI will be undetermined. Refer to Section: Output data line management on an inactive slot for information on how to drive SD line. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaiaslotrType) SetSlotsz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiaslotrFieldSlotszMask)|(uint32(value)<<RegisterSaiaslotrFieldSlotszShift))
}

const (
	RegisterSaiaslotrFieldNbslotShift = 8
	RegisterSaiaslotrFieldNbslotMask  = 0xf00
)

// GetNbslot Number of slots in an audio frame. These bits are set and cleared by software. The value set in this bitfield represents the number of slots + 1 in the audio frame (including the number of inactive slots). The maximum number of slots is 16. The number of slots should be even if FSDEF bit in the SAI_xFRCR register is set. The number of slots must be configured when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaiaslotrType) GetNbslot() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiaslotrFieldNbslotMask) >> RegisterSaiaslotrFieldNbslotShift)
}

// SetNbslot Number of slots in an audio frame. These bits are set and cleared by software. The value set in this bitfield represents the number of slots + 1 in the audio frame (including the number of inactive slots). The maximum number of slots is 16. The number of slots should be even if FSDEF bit in the SAI_xFRCR register is set. The number of slots must be configured when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaiaslotrType) SetNbslot(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiaslotrFieldNbslotMask)|(uint32(value)<<RegisterSaiaslotrFieldNbslotShift))
}

const (
	RegisterSaiaslotrFieldSlotenShift = 16
	RegisterSaiaslotrFieldSlotenMask  = 0xffff0000
)

// GetSloten Slot enable. These bits are set and cleared by software. Each SLOTEN bit corresponds to a slot position from 0 to 15 (maximum 16 slots). The slot must be enabled when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaiaslotrType) GetSloten() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSaiaslotrFieldSlotenMask) >> RegisterSaiaslotrFieldSlotenShift)
}

// SetSloten Slot enable. These bits are set and cleared by software. Each SLOTEN bit corresponds to a slot position from 0 to 15 (maximum 16 slots). The slot must be enabled when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaiaslotrType) SetSloten(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiaslotrFieldSlotenMask)|(uint32(value)<<RegisterSaiaslotrFieldSlotenShift))
}

// registerSaiaimType Interrupt mask register 2
type registerSaiaimType uint32

const (
	RegisterSaiaimFieldOvrudrieShift = 0
	RegisterSaiaimFieldOvrudrieMask  = 0x1
)

// GetOvrudrie Overrun/underrun interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the OVRUDR bit in the SAI_xSR register is set.
func (r *registerSaiaimType) GetOvrudrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaimFieldOvrudrieMask) != 0
}

// SetOvrudrie Overrun/underrun interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the OVRUDR bit in the SAI_xSR register is set.
func (r *registerSaiaimType) SetOvrudrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaimFieldOvrudrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaimFieldOvrudrieMask)
	}
}

const (
	RegisterSaiaimFieldMutedetieShift = 1
	RegisterSaiaimFieldMutedetieMask  = 0x2
)

// GetMutedetie Mute detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the MUTEDET bit in the SAI_xSR register is set. This bit has a meaning only if the audio block is configured in receiver mode.
func (r *registerSaiaimType) GetMutedetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaimFieldMutedetieMask) != 0
}

// SetMutedetie Mute detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the MUTEDET bit in the SAI_xSR register is set. This bit has a meaning only if the audio block is configured in receiver mode.
func (r *registerSaiaimType) SetMutedetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaimFieldMutedetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaimFieldMutedetieMask)
	}
}

const (
	RegisterSaiaimFieldWckcfgieShift = 2
	RegisterSaiaimFieldWckcfgieMask  = 0x4
)

// GetWckcfgie Wrong clock configuration interrupt enable. This bit is set and cleared by software. This bit is taken into account only if the audio block is configured as a master (MODE[1] = 0) and NODIV = 0. It generates an interrupt if the WCKCFG flag in the SAI_xSR register is set. Note: This bit is used only in TDM mode and is meaningless in other modes.
func (r *registerSaiaimType) GetWckcfgie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaimFieldWckcfgieMask) != 0
}

// SetWckcfgie Wrong clock configuration interrupt enable. This bit is set and cleared by software. This bit is taken into account only if the audio block is configured as a master (MODE[1] = 0) and NODIV = 0. It generates an interrupt if the WCKCFG flag in the SAI_xSR register is set. Note: This bit is used only in TDM mode and is meaningless in other modes.
func (r *registerSaiaimType) SetWckcfgie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaimFieldWckcfgieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaimFieldWckcfgieMask)
	}
}

const (
	RegisterSaiaimFieldFreqieShift = 3
	RegisterSaiaimFieldFreqieMask  = 0x8
)

// GetFreqie FIFO request interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the FREQ bit in the SAI_xSR register is set. Since the audio block defaults to operate as a transmitter after reset, the MODE bit must be configured before setting FREQIE to avoid a parasitic interruption in receiver mode,
func (r *registerSaiaimType) GetFreqie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaimFieldFreqieMask) != 0
}

// SetFreqie FIFO request interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the FREQ bit in the SAI_xSR register is set. Since the audio block defaults to operate as a transmitter after reset, the MODE bit must be configured before setting FREQIE to avoid a parasitic interruption in receiver mode,
func (r *registerSaiaimType) SetFreqie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaimFieldFreqieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaimFieldFreqieMask)
	}
}

const (
	RegisterSaiaimFieldCnrdyieShift = 4
	RegisterSaiaimFieldCnrdyieMask  = 0x10
)

// GetCnrdyie Codec not ready interrupt enable (AC97). This bit is set and cleared by software. When the interrupt is enabled, the audio block detects in the slot 0 (tag0) of the AC97 frame if the Codec connected to this line is ready or not. If it is not ready, the CNRDY flag in the SAI_xSR register is set and an interruption i generated. This bit has a meaning only if the AC97 mode is selected through PRTCFG[1:0] bits and the audio block is operates as a receiver.
func (r *registerSaiaimType) GetCnrdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaimFieldCnrdyieMask) != 0
}

// SetCnrdyie Codec not ready interrupt enable (AC97). This bit is set and cleared by software. When the interrupt is enabled, the audio block detects in the slot 0 (tag0) of the AC97 frame if the Codec connected to this line is ready or not. If it is not ready, the CNRDY flag in the SAI_xSR register is set and an interruption i generated. This bit has a meaning only if the AC97 mode is selected through PRTCFG[1:0] bits and the audio block is operates as a receiver.
func (r *registerSaiaimType) SetCnrdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaimFieldCnrdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaimFieldCnrdyieMask)
	}
}

const (
	RegisterSaiaimFieldAfsdetieShift = 5
	RegisterSaiaimFieldAfsdetieMask  = 0x20
)

// GetAfsdetie Anticipated frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the AFSDET bit in the SAI_xSR register is set. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSaiaimType) GetAfsdetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaimFieldAfsdetieMask) != 0
}

// SetAfsdetie Anticipated frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the AFSDET bit in the SAI_xSR register is set. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSaiaimType) SetAfsdetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaimFieldAfsdetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaimFieldAfsdetieMask)
	}
}

const (
	RegisterSaiaimFieldLfsdetieShift = 6
	RegisterSaiaimFieldLfsdetieMask  = 0x40
)

// GetLfsdetie Late frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the LFSDET bit is set in the SAI_xSR register. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSaiaimType) GetLfsdetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaimFieldLfsdetieMask) != 0
}

// SetLfsdetie Late frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the LFSDET bit is set in the SAI_xSR register. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSaiaimType) SetLfsdetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaimFieldLfsdetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaimFieldLfsdetieMask)
	}
}

// registerSaiasrType Status register
type registerSaiasrType uint32

const (
	RegisterSaiasrFieldOvrudrShift = 0
	RegisterSaiasrFieldOvrudrMask  = 0x1
)

// GetOvrudr Overrun / underrun. This bit is read only. The overrun and underrun conditions can occur only when the audio block is configured as a receiver and a transmitter, respectively. It can generate an interrupt if OVRUDRIE bit is set in SAI_xIM register. This flag is cleared when the software sets COVRUDR bit in SAI_xCLRFR register.
func (r *registerSaiasrType) GetOvrudr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiasrFieldOvrudrMask) != 0
}

// SetOvrudr Overrun / underrun. This bit is read only. The overrun and underrun conditions can occur only when the audio block is configured as a receiver and a transmitter, respectively. It can generate an interrupt if OVRUDRIE bit is set in SAI_xIM register. This flag is cleared when the software sets COVRUDR bit in SAI_xCLRFR register.
func (r *registerSaiasrType) SetOvrudr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiasrFieldOvrudrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiasrFieldOvrudrMask)
	}
}

const (
	RegisterSaiasrFieldMutedetShift = 1
	RegisterSaiasrFieldMutedetMask  = 0x2
)

// GetMutedet Mute detection. This bit is read only. This flag is set if consecutive 0 values are received in each slot of a given audio frame and for a consecutive number of audio frames (set in the MUTECNT bit in the SAI_xCR2 register). It can generate an interrupt if MUTEDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets bit CMUTEDET in the SAI_xCLRFR register.
func (r *registerSaiasrType) GetMutedet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiasrFieldMutedetMask) != 0
}

// SetMutedet Mute detection. This bit is read only. This flag is set if consecutive 0 values are received in each slot of a given audio frame and for a consecutive number of audio frames (set in the MUTECNT bit in the SAI_xCR2 register). It can generate an interrupt if MUTEDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets bit CMUTEDET in the SAI_xCLRFR register.
func (r *registerSaiasrType) SetMutedet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiasrFieldMutedetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiasrFieldMutedetMask)
	}
}

const (
	RegisterSaiasrFieldWckcfgShift = 2
	RegisterSaiasrFieldWckcfgMask  = 0x4
)

// GetWckcfg Wrong clock configuration flag. This bit is read only. This bit is used only when the audio block operates in master mode (MODE[1] = 0) and NODIV = 0. It can generate an interrupt if WCKCFGIE bit is set in SAI_xIM register. This flag is cleared when the software sets CWCKCFG bit in SAI_xCLRFR register.
func (r *registerSaiasrType) GetWckcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiasrFieldWckcfgMask) != 0
}

// SetWckcfg Wrong clock configuration flag. This bit is read only. This bit is used only when the audio block operates in master mode (MODE[1] = 0) and NODIV = 0. It can generate an interrupt if WCKCFGIE bit is set in SAI_xIM register. This flag is cleared when the software sets CWCKCFG bit in SAI_xCLRFR register.
func (r *registerSaiasrType) SetWckcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiasrFieldWckcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiasrFieldWckcfgMask)
	}
}

const (
	RegisterSaiasrFieldFreqShift = 3
	RegisterSaiasrFieldFreqMask  = 0x8
)

// GetFreq FIFO request. This bit is read only. The request depends on the audio block configuration: If the block is configured in transmission mode, the FIFO request is related to a write request operation in the SAI_xDR. If the block configured in reception, the FIFO request related to a read request operation from the SAI_xDR. This flag can generate an interrupt if FREQIE bit is set in SAI_xIM register.
func (r *registerSaiasrType) GetFreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiasrFieldFreqMask) != 0
}

// SetFreq FIFO request. This bit is read only. The request depends on the audio block configuration: If the block is configured in transmission mode, the FIFO request is related to a write request operation in the SAI_xDR. If the block configured in reception, the FIFO request related to a read request operation from the SAI_xDR. This flag can generate an interrupt if FREQIE bit is set in SAI_xIM register.
func (r *registerSaiasrType) SetFreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiasrFieldFreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiasrFieldFreqMask)
	}
}

const (
	RegisterSaiasrFieldCnrdyShift = 4
	RegisterSaiasrFieldCnrdyMask  = 0x10
)

// GetCnrdy Codec not ready. This bit is read only. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register and configured in receiver mode. It can generate an interrupt if CNRDYIE bit is set in SAI_xIM register. This flag is cleared when the software sets CCNRDY bit in SAI_xCLRFR register.
func (r *registerSaiasrType) GetCnrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiasrFieldCnrdyMask) != 0
}

// SetCnrdy Codec not ready. This bit is read only. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register and configured in receiver mode. It can generate an interrupt if CNRDYIE bit is set in SAI_xIM register. This flag is cleared when the software sets CCNRDY bit in SAI_xCLRFR register.
func (r *registerSaiasrType) SetCnrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiasrFieldCnrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiasrFieldCnrdyMask)
	}
}

const (
	RegisterSaiasrFieldAfsdetShift = 5
	RegisterSaiasrFieldAfsdetMask  = 0x20
)

// GetAfsdet Anticipated frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97or SPDIF mode. It can generate an interrupt if AFSDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets CAFSDET bit in SAI_xCLRFR register.
func (r *registerSaiasrType) GetAfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiasrFieldAfsdetMask) != 0
}

// SetAfsdet Anticipated frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97or SPDIF mode. It can generate an interrupt if AFSDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets CAFSDET bit in SAI_xCLRFR register.
func (r *registerSaiasrType) SetAfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiasrFieldAfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiasrFieldAfsdetMask)
	}
}

const (
	RegisterSaiasrFieldLfsdetShift = 6
	RegisterSaiasrFieldLfsdetMask  = 0x40
)

// GetLfsdet Late frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97 or SPDIF mode. It can generate an interrupt if LFSDETIE bit is set in the SAI_xIM register. This flag is cleared when the software sets bit CLFSDET in SAI_xCLRFR register
func (r *registerSaiasrType) GetLfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiasrFieldLfsdetMask) != 0
}

// SetLfsdet Late frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97 or SPDIF mode. It can generate an interrupt if LFSDETIE bit is set in the SAI_xIM register. This flag is cleared when the software sets bit CLFSDET in SAI_xCLRFR register
func (r *registerSaiasrType) SetLfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiasrFieldLfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiasrFieldLfsdetMask)
	}
}

const (
	RegisterSaiasrFieldFlvlShift = 16
	RegisterSaiasrFieldFlvlMask  = 0x70000
)

// GetFlvl FIFO level threshold. This bit is read only. The FIFO level threshold flag is managed only by hardware and its setting depends on SAI block configuration (transmitter or receiver mode). If the SAI block is configured as transmitter: If SAI block is configured as receiver:
func (r *registerSaiasrType) GetFlvl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaiasrFieldFlvlMask) >> RegisterSaiasrFieldFlvlShift)
}

// SetFlvl FIFO level threshold. This bit is read only. The FIFO level threshold flag is managed only by hardware and its setting depends on SAI block configuration (transmitter or receiver mode). If the SAI block is configured as transmitter: If SAI block is configured as receiver:
func (r *registerSaiasrType) SetFlvl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiasrFieldFlvlMask)|(uint32(value)<<RegisterSaiasrFieldFlvlShift))
}

// registerSaiaclrfrType Clear flag register
type registerSaiaclrfrType uint32

const (
	RegisterSaiaclrfrFieldCovrudrShift = 0
	RegisterSaiaclrfrFieldCovrudrMask  = 0x1
)

// GetCovrudr Clear overrun / underrun. This bit is write only. Programming this bit to 1 clears the OVRUDR flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) GetCovrudr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaclrfrFieldCovrudrMask) != 0
}

// SetCovrudr Clear overrun / underrun. This bit is write only. Programming this bit to 1 clears the OVRUDR flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) SetCovrudr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaclrfrFieldCovrudrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaclrfrFieldCovrudrMask)
	}
}

const (
	RegisterSaiaclrfrFieldCmutedetShift = 1
	RegisterSaiaclrfrFieldCmutedetMask  = 0x2
)

// GetCmutedet Mute detection flag. This bit is write only. Programming this bit to 1 clears the MUTEDET flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) GetCmutedet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaclrfrFieldCmutedetMask) != 0
}

// SetCmutedet Mute detection flag. This bit is write only. Programming this bit to 1 clears the MUTEDET flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) SetCmutedet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaclrfrFieldCmutedetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaclrfrFieldCmutedetMask)
	}
}

const (
	RegisterSaiaclrfrFieldCwckcfgShift = 2
	RegisterSaiaclrfrFieldCwckcfgMask  = 0x4
)

// GetCwckcfg Clear wrong clock configuration flag. This bit is write only. Programming this bit to 1 clears the WCKCFG flag in the SAI_xSR register. This bit is used only when the audio block is set as master (MODE[1] = 0) and NODIV = 0 in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) GetCwckcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaclrfrFieldCwckcfgMask) != 0
}

// SetCwckcfg Clear wrong clock configuration flag. This bit is write only. Programming this bit to 1 clears the WCKCFG flag in the SAI_xSR register. This bit is used only when the audio block is set as master (MODE[1] = 0) and NODIV = 0 in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) SetCwckcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaclrfrFieldCwckcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaclrfrFieldCwckcfgMask)
	}
}

const (
	RegisterSaiaclrfrFieldCcnrdyShift = 4
	RegisterSaiaclrfrFieldCcnrdyMask  = 0x10
)

// GetCcnrdy Clear Codec not ready flag. This bit is write only. Programming this bit to 1 clears the CNRDY flag in the SAI_xSR register. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) GetCcnrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaclrfrFieldCcnrdyMask) != 0
}

// SetCcnrdy Clear Codec not ready flag. This bit is write only. Programming this bit to 1 clears the CNRDY flag in the SAI_xSR register. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) SetCcnrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaclrfrFieldCcnrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaclrfrFieldCcnrdyMask)
	}
}

const (
	RegisterSaiaclrfrFieldCafsdetShift = 5
	RegisterSaiaclrfrFieldCafsdetMask  = 0x20
)

// GetCafsdet Clear anticipated frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the AFSDET flag in the SAI_xSR register. It is not used in AC97or SPDIF mode. Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) GetCafsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaclrfrFieldCafsdetMask) != 0
}

// SetCafsdet Clear anticipated frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the AFSDET flag in the SAI_xSR register. It is not used in AC97or SPDIF mode. Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) SetCafsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaclrfrFieldCafsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaclrfrFieldCafsdetMask)
	}
}

const (
	RegisterSaiaclrfrFieldClfsdetShift = 6
	RegisterSaiaclrfrFieldClfsdetMask  = 0x40
)

// GetClfsdet Clear late frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the LFSDET flag in the SAI_xSR register. This bit is not used in AC97or SPDIF mode Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) GetClfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaiaclrfrFieldClfsdetMask) != 0
}

// SetClfsdet Clear late frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the LFSDET flag in the SAI_xSR register. This bit is not used in AC97or SPDIF mode Reading this bit always returns the value 0.
func (r *registerSaiaclrfrType) SetClfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaiaclrfrFieldClfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaiaclrfrFieldClfsdetMask)
	}
}

// registerSaiadrType Data register
type registerSaiadrType uint32

const (
	RegisterSaiadrFieldDataShift = 0
	RegisterSaiadrFieldDataMask  = 0xffffffff
)

// GetData Data A write to this register loads the FIFO provided the FIFO is not full. A read from this register empties the FIFO if the FIFO is not empty.
func (r *registerSaiadrType) GetData() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterSaiadrFieldDataMask) >> RegisterSaiadrFieldDataShift)
}

// SetData Data A write to this register loads the FIFO provided the FIFO is not full. A read from this register empties the FIFO if the FIFO is not empty.
func (r *registerSaiadrType) SetData(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaiadrFieldDataMask)|(uint32(value)<<RegisterSaiadrFieldDataShift))
}

// registerSaibcr1Type Configuration register 1
type registerSaibcr1Type uint32

const (
	RegisterSaibcr1FieldModeShift = 0
	RegisterSaibcr1FieldModeMask  = 0x3
)

// GetMode SAIx audio block mode immediately
func (r *registerSaibcr1Type) GetMode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldModeMask) >> RegisterSaibcr1FieldModeShift)
}

// SetMode SAIx audio block mode immediately
func (r *registerSaibcr1Type) SetMode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldModeMask)|(uint32(value)<<RegisterSaibcr1FieldModeShift))
}

const (
	RegisterSaibcr1FieldPrtcfgShift = 2
	RegisterSaibcr1FieldPrtcfgMask  = 0xc
)

// GetPrtcfg Protocol configuration. These bits are set and cleared by software. These bits have to be configured when the audio block is disabled.
func (r *registerSaibcr1Type) GetPrtcfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldPrtcfgMask) >> RegisterSaibcr1FieldPrtcfgShift)
}

// SetPrtcfg Protocol configuration. These bits are set and cleared by software. These bits have to be configured when the audio block is disabled.
func (r *registerSaibcr1Type) SetPrtcfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldPrtcfgMask)|(uint32(value)<<RegisterSaibcr1FieldPrtcfgShift))
}

const (
	RegisterSaibcr1FieldDsShift = 5
	RegisterSaibcr1FieldDsMask  = 0xe0
)

// GetDs Data size. These bits are set and cleared by software. These bits are ignored when the SPDIF protocols are selected (bit PRTCFG[1:0]), because the frame and the data size are fixed in such case. When the companding mode is selected through COMP[1:0] bits, DS[1:0] are ignored since the data size is fixed to 8 bits by the algorithm. These bits must be configured when the audio block is disabled.
func (r *registerSaibcr1Type) GetDs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldDsMask) >> RegisterSaibcr1FieldDsShift)
}

// SetDs Data size. These bits are set and cleared by software. These bits are ignored when the SPDIF protocols are selected (bit PRTCFG[1:0]), because the frame and the data size are fixed in such case. When the companding mode is selected through COMP[1:0] bits, DS[1:0] are ignored since the data size is fixed to 8 bits by the algorithm. These bits must be configured when the audio block is disabled.
func (r *registerSaibcr1Type) SetDs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldDsMask)|(uint32(value)<<RegisterSaibcr1FieldDsShift))
}

const (
	RegisterSaibcr1FieldLsbfirstShift = 8
	RegisterSaibcr1FieldLsbfirstMask  = 0x100
)

// GetLsbfirst Least significant bit first. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in AC97 audio protocol since AC97 data are always transferred with the MSB first. This bit has no meaning in SPDIF audio protocol since in SPDIF data are always transferred with LSB first.
func (r *registerSaibcr1Type) GetLsbfirst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldLsbfirstMask) != 0
}

// SetLsbfirst Least significant bit first. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in AC97 audio protocol since AC97 data are always transferred with the MSB first. This bit has no meaning in SPDIF audio protocol since in SPDIF data are always transferred with LSB first.
func (r *registerSaibcr1Type) SetLsbfirst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr1FieldLsbfirstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldLsbfirstMask)
	}
}

const (
	RegisterSaibcr1FieldCkstrShift = 9
	RegisterSaibcr1FieldCkstrMask  = 0x200
)

// GetCkstr Clock strobing edge. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in SPDIF audio protocol.
func (r *registerSaibcr1Type) GetCkstr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldCkstrMask) != 0
}

// SetCkstr Clock strobing edge. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in SPDIF audio protocol.
func (r *registerSaibcr1Type) SetCkstr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr1FieldCkstrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldCkstrMask)
	}
}

const (
	RegisterSaibcr1FieldSyncenShift = 10
	RegisterSaibcr1FieldSyncenMask  = 0xc00
)

// GetSyncen Synchronization enable. These bits are set and cleared by software. They must be configured when the audio sub-block is disabled. Note: The audio sub-block should be configured as asynchronous when SPDIF mode is enabled.
func (r *registerSaibcr1Type) GetSyncen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldSyncenMask) >> RegisterSaibcr1FieldSyncenShift)
}

// SetSyncen Synchronization enable. These bits are set and cleared by software. They must be configured when the audio sub-block is disabled. Note: The audio sub-block should be configured as asynchronous when SPDIF mode is enabled.
func (r *registerSaibcr1Type) SetSyncen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldSyncenMask)|(uint32(value)<<RegisterSaibcr1FieldSyncenShift))
}

const (
	RegisterSaibcr1FieldMonoShift = 12
	RegisterSaibcr1FieldMonoMask  = 0x1000
)

// GetMono Mono mode. This bit is set and cleared by software. It is meaningful only when the number of slots is equal to 2. When the mono mode is selected, slot 0 data are duplicated on slot 1 when the audio block operates as a transmitter. In reception mode, the slot1 is discarded and only the data received from slot 0 are stored. Refer to Section: Mono/stereo mode for more details.
func (r *registerSaibcr1Type) GetMono() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldMonoMask) != 0
}

// SetMono Mono mode. This bit is set and cleared by software. It is meaningful only when the number of slots is equal to 2. When the mono mode is selected, slot 0 data are duplicated on slot 1 when the audio block operates as a transmitter. In reception mode, the slot1 is discarded and only the data received from slot 0 are stored. Refer to Section: Mono/stereo mode for more details.
func (r *registerSaibcr1Type) SetMono(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr1FieldMonoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldMonoMask)
	}
}

const (
	RegisterSaibcr1FieldOutdrivShift = 13
	RegisterSaibcr1FieldOutdrivMask  = 0x2000
)

// GetOutdriv Output drive. This bit is set and cleared by software. Note: This bit has to be set before enabling the audio block and after the audio block configuration.
func (r *registerSaibcr1Type) GetOutdriv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldOutdrivMask) != 0
}

// SetOutdriv Output drive. This bit is set and cleared by software. Note: This bit has to be set before enabling the audio block and after the audio block configuration.
func (r *registerSaibcr1Type) SetOutdriv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr1FieldOutdrivMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldOutdrivMask)
	}
}

const (
	RegisterSaibcr1FieldSaixenShift = 16
	RegisterSaibcr1FieldSaixenMask  = 0x10000
)

// GetSaixen Audio block enable where x is A or B. This bit is set by software. To switch off the audio block, the application software must program this bit to 0 and poll the bit till it reads back 0, meaning that the block is completely disabled. Before setting this bit to 1, check that it is set to 0, otherwise the enable command will not be taken into account. This bit allows to control the state of SAIx audio block. If it is disabled when an audio frame transfer is ongoing, the ongoing transfer completes and the cell is fully disabled at the end of this audio frame transfer. Note: When SAIx block is configured in master mode, the clock must be present on the input of SAIx before setting SAIXEN bit.
func (r *registerSaibcr1Type) GetSaixen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldSaixenMask) != 0
}

// SetSaixen Audio block enable where x is A or B. This bit is set by software. To switch off the audio block, the application software must program this bit to 0 and poll the bit till it reads back 0, meaning that the block is completely disabled. Before setting this bit to 1, check that it is set to 0, otherwise the enable command will not be taken into account. This bit allows to control the state of SAIx audio block. If it is disabled when an audio frame transfer is ongoing, the ongoing transfer completes and the cell is fully disabled at the end of this audio frame transfer. Note: When SAIx block is configured in master mode, the clock must be present on the input of SAIx before setting SAIXEN bit.
func (r *registerSaibcr1Type) SetSaixen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr1FieldSaixenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldSaixenMask)
	}
}

const (
	RegisterSaibcr1FieldDmaenShift = 17
	RegisterSaibcr1FieldDmaenMask  = 0x20000
)

// GetDmaen DMA enable. This bit is set and cleared by software. Note: Since the audio block defaults to operate as a transmitter after reset, the MODE[1:0] bits must be configured before setting DMAEN to avoid a DMA request in receiver mode.
func (r *registerSaibcr1Type) GetDmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldDmaenMask) != 0
}

// SetDmaen DMA enable. This bit is set and cleared by software. Note: Since the audio block defaults to operate as a transmitter after reset, the MODE[1:0] bits must be configured before setting DMAEN to avoid a DMA request in receiver mode.
func (r *registerSaibcr1Type) SetDmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr1FieldDmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldDmaenMask)
	}
}

const (
	RegisterSaibcr1FieldNomckShift = 19
	RegisterSaibcr1FieldNomckMask  = 0x80000
)

// GetNomck No divider
func (r *registerSaibcr1Type) GetNomck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldNomckMask) != 0
}

// SetNomck No divider
func (r *registerSaibcr1Type) SetNomck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr1FieldNomckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldNomckMask)
	}
}

const (
	RegisterSaibcr1FieldMckdivShift = 20
	RegisterSaibcr1FieldMckdivMask  = 0xf00000
)

// GetMckdiv Master clock divider. These bits are set and cleared by software. These bits are meaningless when the audio block operates in slave mode. They have to be configured when the audio block is disabled. Others: the master clock frequency is calculated accordingly to the following formula:
func (r *registerSaibcr1Type) GetMckdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldMckdivMask) >> RegisterSaibcr1FieldMckdivShift)
}

// SetMckdiv Master clock divider. These bits are set and cleared by software. These bits are meaningless when the audio block operates in slave mode. They have to be configured when the audio block is disabled. Others: the master clock frequency is calculated accordingly to the following formula:
func (r *registerSaibcr1Type) SetMckdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldMckdivMask)|(uint32(value)<<RegisterSaibcr1FieldMckdivShift))
}

const (
	RegisterSaibcr1FieldOsrShift = 26
	RegisterSaibcr1FieldOsrMask  = 0x4000000
)

// GetOsr Oversampling ratio for master clock
func (r *registerSaibcr1Type) GetOsr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr1FieldOsrMask) != 0
}

// SetOsr Oversampling ratio for master clock
func (r *registerSaibcr1Type) SetOsr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr1FieldOsrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr1FieldOsrMask)
	}
}

// registerSaibcr2Type Configuration register 2
type registerSaibcr2Type uint32

const (
	RegisterSaibcr2FieldFthShift = 0
	RegisterSaibcr2FieldFthMask  = 0x7
)

// GetFth FIFO threshold. This bit is set and cleared by software.
func (r *registerSaibcr2Type) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr2FieldFthMask) >> RegisterSaibcr2FieldFthShift)
}

// SetFth FIFO threshold. This bit is set and cleared by software.
func (r *registerSaibcr2Type) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr2FieldFthMask)|(uint32(value)<<RegisterSaibcr2FieldFthShift))
}

const (
	RegisterSaibcr2FieldFflushShift = 3
	RegisterSaibcr2FieldFflushMask  = 0x8
)

// SetFflush FIFO flush. This bit is set by software. It is always read as 0. This bit should be configured when the SAI is disabled.
func (r *registerSaibcr2Type) SetFflush(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr2FieldFflushMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr2FieldFflushMask)
	}
}

const (
	RegisterSaibcr2FieldTrisShift = 4
	RegisterSaibcr2FieldTrisMask  = 0x10
)

// GetTris Tristate management on data line. This bit is set and cleared by software. It is meaningful only if the audio block is configured as a transmitter. This bit is not used when the audio block is configured in SPDIF mode. It should be configured when SAI is disabled. Refer to Section: Output data line management on an inactive slot for more details.
func (r *registerSaibcr2Type) GetTris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr2FieldTrisMask) != 0
}

// SetTris Tristate management on data line. This bit is set and cleared by software. It is meaningful only if the audio block is configured as a transmitter. This bit is not used when the audio block is configured in SPDIF mode. It should be configured when SAI is disabled. Refer to Section: Output data line management on an inactive slot for more details.
func (r *registerSaibcr2Type) SetTris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr2FieldTrisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr2FieldTrisMask)
	}
}

const (
	RegisterSaibcr2FieldMuteShift = 5
	RegisterSaibcr2FieldMuteMask  = 0x20
)

// GetMute Mute. This bit is set and cleared by software. It is meaningful only when the audio block operates as a transmitter. The MUTE value is linked to value of MUTEVAL if the number of slots is lower or equal to 2, or equal to 0 if it is greater than 2. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSaibcr2Type) GetMute() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr2FieldMuteMask) != 0
}

// SetMute Mute. This bit is set and cleared by software. It is meaningful only when the audio block operates as a transmitter. The MUTE value is linked to value of MUTEVAL if the number of slots is lower or equal to 2, or equal to 0 if it is greater than 2. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSaibcr2Type) SetMute(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr2FieldMuteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr2FieldMuteMask)
	}
}

const (
	RegisterSaibcr2FieldMutevalShift = 6
	RegisterSaibcr2FieldMutevalMask  = 0x40
)

// GetMuteval Mute value. This bit is set and cleared by software.It must be written before enabling the audio block: SAIXEN. This bit is meaningful only when the audio block operates as a transmitter, the number of slots is lower or equal to 2 and the MUTE bit is set. If more slots are declared, the bit value sent during the transmission in mute mode is equal to 0, whatever the value of MUTEVAL. if the number of slot is lower or equal to 2 and MUTEVAL = 1, the MUTE value transmitted for each slot is the one sent during the previous frame. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSaibcr2Type) GetMuteval() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr2FieldMutevalMask) != 0
}

// SetMuteval Mute value. This bit is set and cleared by software.It must be written before enabling the audio block: SAIXEN. This bit is meaningful only when the audio block operates as a transmitter, the number of slots is lower or equal to 2 and the MUTE bit is set. If more slots are declared, the bit value sent during the transmission in mute mode is equal to 0, whatever the value of MUTEVAL. if the number of slot is lower or equal to 2 and MUTEVAL = 1, the MUTE value transmitted for each slot is the one sent during the previous frame. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSaibcr2Type) SetMuteval(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr2FieldMutevalMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr2FieldMutevalMask)
	}
}

const (
	RegisterSaibcr2FieldMutecntShift = 7
	RegisterSaibcr2FieldMutecntMask  = 0x1f80
)

// GetMutecnt Mute counter. These bits are set and cleared by software. They are used only in reception mode. The value set in these bits is compared to the number of consecutive mute frames detected in reception. When the number of mute frames is equal to this value, the flag MUTEDET will be set and an interrupt will be generated if bit MUTEDETIE is set. Refer to Section: Mute mode for more details.
func (r *registerSaibcr2Type) GetMutecnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr2FieldMutecntMask) >> RegisterSaibcr2FieldMutecntShift)
}

// SetMutecnt Mute counter. These bits are set and cleared by software. They are used only in reception mode. The value set in these bits is compared to the number of consecutive mute frames detected in reception. When the number of mute frames is equal to this value, the flag MUTEDET will be set and an interrupt will be generated if bit MUTEDETIE is set. Refer to Section: Mute mode for more details.
func (r *registerSaibcr2Type) SetMutecnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr2FieldMutecntMask)|(uint32(value)<<RegisterSaibcr2FieldMutecntShift))
}

const (
	RegisterSaibcr2FieldCplShift = 13
	RegisterSaibcr2FieldCplMask  = 0x2000
)

// GetCpl Complement bit. This bit is set and cleared by software. It defines the type of complement to be used for companding mode Note: This bit has effect only when the companding mode is -Law algorithm or A-Law algorithm.
func (r *registerSaibcr2Type) GetCpl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr2FieldCplMask) != 0
}

// SetCpl Complement bit. This bit is set and cleared by software. It defines the type of complement to be used for companding mode Note: This bit has effect only when the companding mode is -Law algorithm or A-Law algorithm.
func (r *registerSaibcr2Type) SetCpl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibcr2FieldCplMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr2FieldCplMask)
	}
}

const (
	RegisterSaibcr2FieldCompShift = 14
	RegisterSaibcr2FieldCompMask  = 0xc000
)

// GetComp Companding mode. These bits are set and cleared by software. The -Law and the A-Law log are a part of the CCITT G.711 recommendation, the type of complement that will be used depends on CPL bit. The data expansion or data compression are determined by the state of bit MODE[0]. The data compression is applied if the audio block is configured as a transmitter. The data expansion is automatically applied when the audio block is configured as a receiver. Refer to Section: Companding mode for more details. Note: Companding mode is applicable only when TDM is selected.
func (r *registerSaibcr2Type) GetComp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibcr2FieldCompMask) >> RegisterSaibcr2FieldCompShift)
}

// SetComp Companding mode. These bits are set and cleared by software. The -Law and the A-Law log are a part of the CCITT G.711 recommendation, the type of complement that will be used depends on CPL bit. The data expansion or data compression are determined by the state of bit MODE[0]. The data compression is applied if the audio block is configured as a transmitter. The data expansion is automatically applied when the audio block is configured as a receiver. Refer to Section: Companding mode for more details. Note: Companding mode is applicable only when TDM is selected.
func (r *registerSaibcr2Type) SetComp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibcr2FieldCompMask)|(uint32(value)<<RegisterSaibcr2FieldCompShift))
}

// registerSaibfrcrType This register has no meaning in AC97 and SPDIF audio protocol
type registerSaibfrcrType uint32

const (
	RegisterSaibfrcrFieldFrlShift = 0
	RegisterSaibfrcrFieldFrlMask  = 0xff
)

// GetFrl Frame length. These bits are set and cleared by software. They define the audio frame length expressed in number of SCK clock cycles: the number of bits in the frame is equal to FRL[7:0] + 1. The minimum number of bits to transfer in an audio frame must be equal to 8, otherwise the audio block will behaves in an unexpected way. This is the case when the data size is 8 bits and only one slot 0 is defined in NBSLOT[4:0] of SAI_xSLOTR register (NBSLOT[3:0] = 0000). In master mode, if the master clock (available on MCLK_x pin) is used, the frame length should be aligned with a number equal to a power of 2, ranging from 8 to 256. When the master clock is not used (NODIV = 1), it is recommended to program the frame length to an value ranging from 8 to 256. These bits are meaningless and are not used in AC97 or SPDIF audio block configuration.
func (r *registerSaibfrcrType) GetFrl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibfrcrFieldFrlMask) >> RegisterSaibfrcrFieldFrlShift)
}

// SetFrl Frame length. These bits are set and cleared by software. They define the audio frame length expressed in number of SCK clock cycles: the number of bits in the frame is equal to FRL[7:0] + 1. The minimum number of bits to transfer in an audio frame must be equal to 8, otherwise the audio block will behaves in an unexpected way. This is the case when the data size is 8 bits and only one slot 0 is defined in NBSLOT[4:0] of SAI_xSLOTR register (NBSLOT[3:0] = 0000). In master mode, if the master clock (available on MCLK_x pin) is used, the frame length should be aligned with a number equal to a power of 2, ranging from 8 to 256. When the master clock is not used (NODIV = 1), it is recommended to program the frame length to an value ranging from 8 to 256. These bits are meaningless and are not used in AC97 or SPDIF audio block configuration.
func (r *registerSaibfrcrType) SetFrl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibfrcrFieldFrlMask)|(uint32(value)<<RegisterSaibfrcrFieldFrlShift))
}

const (
	RegisterSaibfrcrFieldFsallShift = 8
	RegisterSaibfrcrFieldFsallMask  = 0x7f00
)

// GetFsall Frame synchronization active level length. These bits are set and cleared by software. They specify the length in number of bit clock (SCK) + 1 (FSALL[6:0] + 1) of the active level of the FS signal in the audio frame These bits are meaningless and are not used in AC97 or SPDIF audio block configuration. They must be configured when the audio block is disabled.
func (r *registerSaibfrcrType) GetFsall() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibfrcrFieldFsallMask) >> RegisterSaibfrcrFieldFsallShift)
}

// SetFsall Frame synchronization active level length. These bits are set and cleared by software. They specify the length in number of bit clock (SCK) + 1 (FSALL[6:0] + 1) of the active level of the FS signal in the audio frame These bits are meaningless and are not used in AC97 or SPDIF audio block configuration. They must be configured when the audio block is disabled.
func (r *registerSaibfrcrType) SetFsall(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibfrcrFieldFsallMask)|(uint32(value)<<RegisterSaibfrcrFieldFsallShift))
}

const (
	RegisterSaibfrcrFieldFsdefShift = 16
	RegisterSaibfrcrFieldFsdefMask  = 0x10000
)

// GetFsdef Frame synchronization definition. This bit is set and cleared by software. When the bit is set, the number of slots defined in the SAI_xSLOTR register has to be even. It means that half of this number of slots will be dedicated to the left channel and the other slots for the right channel (e.g: this bit has to be set for I2S or MSB/LSB-justified protocols...). This bit is meaningless and is not used in AC97 or SPDIF audio block configuration. It must be configured when the audio block is disabled.
func (r *registerSaibfrcrType) GetFsdef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibfrcrFieldFsdefMask) != 0
}

const (
	RegisterSaibfrcrFieldFspolShift = 17
	RegisterSaibfrcrFieldFspolMask  = 0x20000
)

// GetFspol Frame synchronization polarity. This bit is set and cleared by software. It is used to configure the level of the start of frame on the FS signal. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSaibfrcrType) GetFspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibfrcrFieldFspolMask) != 0
}

// SetFspol Frame synchronization polarity. This bit is set and cleared by software. It is used to configure the level of the start of frame on the FS signal. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSaibfrcrType) SetFspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibfrcrFieldFspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibfrcrFieldFspolMask)
	}
}

const (
	RegisterSaibfrcrFieldFsoffShift = 18
	RegisterSaibfrcrFieldFsoffMask  = 0x40000
)

// GetFsoff Frame synchronization offset. This bit is set and cleared by software. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSaibfrcrType) GetFsoff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibfrcrFieldFsoffMask) != 0
}

// SetFsoff Frame synchronization offset. This bit is set and cleared by software. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSaibfrcrType) SetFsoff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibfrcrFieldFsoffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibfrcrFieldFsoffMask)
	}
}

// registerSaibslotrType This register has no meaning in AC97 and SPDIF audio protocol
type registerSaibslotrType uint32

const (
	RegisterSaibslotrFieldFboffShift = 0
	RegisterSaibslotrFieldFboffMask  = 0x1f
)

// GetFboff First bit offset These bits are set and cleared by software. The value set in this bitfield defines the position of the first data transfer bit in the slot. It represents an offset value. In transmission mode, the bits outside the data field are forced to 0. In reception mode, the extra received bits are discarded. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaibslotrType) GetFboff() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibslotrFieldFboffMask) >> RegisterSaibslotrFieldFboffShift)
}

// SetFboff First bit offset These bits are set and cleared by software. The value set in this bitfield defines the position of the first data transfer bit in the slot. It represents an offset value. In transmission mode, the bits outside the data field are forced to 0. In reception mode, the extra received bits are discarded. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaibslotrType) SetFboff(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibslotrFieldFboffMask)|(uint32(value)<<RegisterSaibslotrFieldFboffShift))
}

const (
	RegisterSaibslotrFieldSlotszShift = 6
	RegisterSaibslotrFieldSlotszMask  = 0xc0
)

// GetSlotsz Slot size This bits is set and cleared by software. The slot size must be higher or equal to the data size. If this condition is not respected, the behavior of the SAI will be undetermined. Refer to Section: Output data line management on an inactive slot for information on how to drive SD line. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaibslotrType) GetSlotsz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibslotrFieldSlotszMask) >> RegisterSaibslotrFieldSlotszShift)
}

// SetSlotsz Slot size This bits is set and cleared by software. The slot size must be higher or equal to the data size. If this condition is not respected, the behavior of the SAI will be undetermined. Refer to Section: Output data line management on an inactive slot for information on how to drive SD line. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaibslotrType) SetSlotsz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibslotrFieldSlotszMask)|(uint32(value)<<RegisterSaibslotrFieldSlotszShift))
}

const (
	RegisterSaibslotrFieldNbslotShift = 8
	RegisterSaibslotrFieldNbslotMask  = 0xf00
)

// GetNbslot Number of slots in an audio frame. These bits are set and cleared by software. The value set in this bitfield represents the number of slots + 1 in the audio frame (including the number of inactive slots). The maximum number of slots is 16. The number of slots should be even if FSDEF bit in the SAI_xFRCR register is set. The number of slots must be configured when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaibslotrType) GetNbslot() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibslotrFieldNbslotMask) >> RegisterSaibslotrFieldNbslotShift)
}

// SetNbslot Number of slots in an audio frame. These bits are set and cleared by software. The value set in this bitfield represents the number of slots + 1 in the audio frame (including the number of inactive slots). The maximum number of slots is 16. The number of slots should be even if FSDEF bit in the SAI_xFRCR register is set. The number of slots must be configured when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaibslotrType) SetNbslot(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibslotrFieldNbslotMask)|(uint32(value)<<RegisterSaibslotrFieldNbslotShift))
}

const (
	RegisterSaibslotrFieldSlotenShift = 16
	RegisterSaibslotrFieldSlotenMask  = 0xffff0000
)

// GetSloten Slot enable. These bits are set and cleared by software. Each SLOTEN bit corresponds to a slot position from 0 to 15 (maximum 16 slots). The slot must be enabled when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaibslotrType) GetSloten() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSaibslotrFieldSlotenMask) >> RegisterSaibslotrFieldSlotenShift)
}

// SetSloten Slot enable. These bits are set and cleared by software. Each SLOTEN bit corresponds to a slot position from 0 to 15 (maximum 16 slots). The slot must be enabled when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSaibslotrType) SetSloten(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibslotrFieldSlotenMask)|(uint32(value)<<RegisterSaibslotrFieldSlotenShift))
}

// registerSaibimType Interrupt mask register 2
type registerSaibimType uint32

const (
	RegisterSaibimFieldOvrudrieShift = 0
	RegisterSaibimFieldOvrudrieMask  = 0x1
)

// GetOvrudrie Overrun/underrun interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the OVRUDR bit in the SAI_xSR register is set.
func (r *registerSaibimType) GetOvrudrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibimFieldOvrudrieMask) != 0
}

// SetOvrudrie Overrun/underrun interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the OVRUDR bit in the SAI_xSR register is set.
func (r *registerSaibimType) SetOvrudrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibimFieldOvrudrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibimFieldOvrudrieMask)
	}
}

const (
	RegisterSaibimFieldMutedetieShift = 1
	RegisterSaibimFieldMutedetieMask  = 0x2
)

// GetMutedetie Mute detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the MUTEDET bit in the SAI_xSR register is set. This bit has a meaning only if the audio block is configured in receiver mode.
func (r *registerSaibimType) GetMutedetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibimFieldMutedetieMask) != 0
}

// SetMutedetie Mute detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the MUTEDET bit in the SAI_xSR register is set. This bit has a meaning only if the audio block is configured in receiver mode.
func (r *registerSaibimType) SetMutedetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibimFieldMutedetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibimFieldMutedetieMask)
	}
}

const (
	RegisterSaibimFieldWckcfgieShift = 2
	RegisterSaibimFieldWckcfgieMask  = 0x4
)

// GetWckcfgie Wrong clock configuration interrupt enable. This bit is set and cleared by software. This bit is taken into account only if the audio block is configured as a master (MODE[1] = 0) and NODIV = 0. It generates an interrupt if the WCKCFG flag in the SAI_xSR register is set. Note: This bit is used only in TDM mode and is meaningless in other modes.
func (r *registerSaibimType) GetWckcfgie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibimFieldWckcfgieMask) != 0
}

// SetWckcfgie Wrong clock configuration interrupt enable. This bit is set and cleared by software. This bit is taken into account only if the audio block is configured as a master (MODE[1] = 0) and NODIV = 0. It generates an interrupt if the WCKCFG flag in the SAI_xSR register is set. Note: This bit is used only in TDM mode and is meaningless in other modes.
func (r *registerSaibimType) SetWckcfgie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibimFieldWckcfgieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibimFieldWckcfgieMask)
	}
}

const (
	RegisterSaibimFieldFreqieShift = 3
	RegisterSaibimFieldFreqieMask  = 0x8
)

// GetFreqie FIFO request interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the FREQ bit in the SAI_xSR register is set. Since the audio block defaults to operate as a transmitter after reset, the MODE bit must be configured before setting FREQIE to avoid a parasitic interruption in receiver mode,
func (r *registerSaibimType) GetFreqie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibimFieldFreqieMask) != 0
}

// SetFreqie FIFO request interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the FREQ bit in the SAI_xSR register is set. Since the audio block defaults to operate as a transmitter after reset, the MODE bit must be configured before setting FREQIE to avoid a parasitic interruption in receiver mode,
func (r *registerSaibimType) SetFreqie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibimFieldFreqieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibimFieldFreqieMask)
	}
}

const (
	RegisterSaibimFieldCnrdyieShift = 4
	RegisterSaibimFieldCnrdyieMask  = 0x10
)

// GetCnrdyie Codec not ready interrupt enable (AC97). This bit is set and cleared by software. When the interrupt is enabled, the audio block detects in the slot 0 (tag0) of the AC97 frame if the Codec connected to this line is ready or not. If it is not ready, the CNRDY flag in the SAI_xSR register is set and an interruption i generated. This bit has a meaning only if the AC97 mode is selected through PRTCFG[1:0] bits and the audio block is operates as a receiver.
func (r *registerSaibimType) GetCnrdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibimFieldCnrdyieMask) != 0
}

// SetCnrdyie Codec not ready interrupt enable (AC97). This bit is set and cleared by software. When the interrupt is enabled, the audio block detects in the slot 0 (tag0) of the AC97 frame if the Codec connected to this line is ready or not. If it is not ready, the CNRDY flag in the SAI_xSR register is set and an interruption i generated. This bit has a meaning only if the AC97 mode is selected through PRTCFG[1:0] bits and the audio block is operates as a receiver.
func (r *registerSaibimType) SetCnrdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibimFieldCnrdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibimFieldCnrdyieMask)
	}
}

const (
	RegisterSaibimFieldAfsdetieShift = 5
	RegisterSaibimFieldAfsdetieMask  = 0x20
)

// GetAfsdetie Anticipated frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the AFSDET bit in the SAI_xSR register is set. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSaibimType) GetAfsdetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibimFieldAfsdetieMask) != 0
}

// SetAfsdetie Anticipated frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the AFSDET bit in the SAI_xSR register is set. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSaibimType) SetAfsdetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibimFieldAfsdetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibimFieldAfsdetieMask)
	}
}

const (
	RegisterSaibimFieldLfsdetieShift = 6
	RegisterSaibimFieldLfsdetieMask  = 0x40
)

// GetLfsdetie Late frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the LFSDET bit is set in the SAI_xSR register. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSaibimType) GetLfsdetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibimFieldLfsdetieMask) != 0
}

// SetLfsdetie Late frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the LFSDET bit is set in the SAI_xSR register. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSaibimType) SetLfsdetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibimFieldLfsdetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibimFieldLfsdetieMask)
	}
}

// registerSaibsrType Status register
type registerSaibsrType uint32

const (
	RegisterSaibsrFieldOvrudrShift = 0
	RegisterSaibsrFieldOvrudrMask  = 0x1
)

// GetOvrudr Overrun / underrun. This bit is read only. The overrun and underrun conditions can occur only when the audio block is configured as a receiver and a transmitter, respectively. It can generate an interrupt if OVRUDRIE bit is set in SAI_xIM register. This flag is cleared when the software sets COVRUDR bit in SAI_xCLRFR register.
func (r *registerSaibsrType) GetOvrudr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibsrFieldOvrudrMask) != 0
}

// SetOvrudr Overrun / underrun. This bit is read only. The overrun and underrun conditions can occur only when the audio block is configured as a receiver and a transmitter, respectively. It can generate an interrupt if OVRUDRIE bit is set in SAI_xIM register. This flag is cleared when the software sets COVRUDR bit in SAI_xCLRFR register.
func (r *registerSaibsrType) SetOvrudr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibsrFieldOvrudrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibsrFieldOvrudrMask)
	}
}

const (
	RegisterSaibsrFieldMutedetShift = 1
	RegisterSaibsrFieldMutedetMask  = 0x2
)

// GetMutedet Mute detection. This bit is read only. This flag is set if consecutive 0 values are received in each slot of a given audio frame and for a consecutive number of audio frames (set in the MUTECNT bit in the SAI_xCR2 register). It can generate an interrupt if MUTEDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets bit CMUTEDET in the SAI_xCLRFR register.
func (r *registerSaibsrType) GetMutedet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibsrFieldMutedetMask) != 0
}

// SetMutedet Mute detection. This bit is read only. This flag is set if consecutive 0 values are received in each slot of a given audio frame and for a consecutive number of audio frames (set in the MUTECNT bit in the SAI_xCR2 register). It can generate an interrupt if MUTEDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets bit CMUTEDET in the SAI_xCLRFR register.
func (r *registerSaibsrType) SetMutedet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibsrFieldMutedetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibsrFieldMutedetMask)
	}
}

const (
	RegisterSaibsrFieldWckcfgShift = 2
	RegisterSaibsrFieldWckcfgMask  = 0x4
)

// GetWckcfg Wrong clock configuration flag. This bit is read only. This bit is used only when the audio block operates in master mode (MODE[1] = 0) and NODIV = 0. It can generate an interrupt if WCKCFGIE bit is set in SAI_xIM register. This flag is cleared when the software sets CWCKCFG bit in SAI_xCLRFR register.
func (r *registerSaibsrType) GetWckcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibsrFieldWckcfgMask) != 0
}

// SetWckcfg Wrong clock configuration flag. This bit is read only. This bit is used only when the audio block operates in master mode (MODE[1] = 0) and NODIV = 0. It can generate an interrupt if WCKCFGIE bit is set in SAI_xIM register. This flag is cleared when the software sets CWCKCFG bit in SAI_xCLRFR register.
func (r *registerSaibsrType) SetWckcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibsrFieldWckcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibsrFieldWckcfgMask)
	}
}

const (
	RegisterSaibsrFieldFreqShift = 3
	RegisterSaibsrFieldFreqMask  = 0x8
)

// GetFreq FIFO request. This bit is read only. The request depends on the audio block configuration: If the block is configured in transmission mode, the FIFO request is related to a write request operation in the SAI_xDR. If the block configured in reception, the FIFO request related to a read request operation from the SAI_xDR. This flag can generate an interrupt if FREQIE bit is set in SAI_xIM register.
func (r *registerSaibsrType) GetFreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibsrFieldFreqMask) != 0
}

// SetFreq FIFO request. This bit is read only. The request depends on the audio block configuration: If the block is configured in transmission mode, the FIFO request is related to a write request operation in the SAI_xDR. If the block configured in reception, the FIFO request related to a read request operation from the SAI_xDR. This flag can generate an interrupt if FREQIE bit is set in SAI_xIM register.
func (r *registerSaibsrType) SetFreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibsrFieldFreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibsrFieldFreqMask)
	}
}

const (
	RegisterSaibsrFieldCnrdyShift = 4
	RegisterSaibsrFieldCnrdyMask  = 0x10
)

// GetCnrdy Codec not ready. This bit is read only. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register and configured in receiver mode. It can generate an interrupt if CNRDYIE bit is set in SAI_xIM register. This flag is cleared when the software sets CCNRDY bit in SAI_xCLRFR register.
func (r *registerSaibsrType) GetCnrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibsrFieldCnrdyMask) != 0
}

// SetCnrdy Codec not ready. This bit is read only. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register and configured in receiver mode. It can generate an interrupt if CNRDYIE bit is set in SAI_xIM register. This flag is cleared when the software sets CCNRDY bit in SAI_xCLRFR register.
func (r *registerSaibsrType) SetCnrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibsrFieldCnrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibsrFieldCnrdyMask)
	}
}

const (
	RegisterSaibsrFieldAfsdetShift = 5
	RegisterSaibsrFieldAfsdetMask  = 0x20
)

// GetAfsdet Anticipated frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97or SPDIF mode. It can generate an interrupt if AFSDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets CAFSDET bit in SAI_xCLRFR register.
func (r *registerSaibsrType) GetAfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibsrFieldAfsdetMask) != 0
}

// SetAfsdet Anticipated frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97or SPDIF mode. It can generate an interrupt if AFSDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets CAFSDET bit in SAI_xCLRFR register.
func (r *registerSaibsrType) SetAfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibsrFieldAfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibsrFieldAfsdetMask)
	}
}

const (
	RegisterSaibsrFieldLfsdetShift = 6
	RegisterSaibsrFieldLfsdetMask  = 0x40
)

// GetLfsdet Late frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97 or SPDIF mode. It can generate an interrupt if LFSDETIE bit is set in the SAI_xIM register. This flag is cleared when the software sets bit CLFSDET in SAI_xCLRFR register
func (r *registerSaibsrType) GetLfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibsrFieldLfsdetMask) != 0
}

// SetLfsdet Late frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97 or SPDIF mode. It can generate an interrupt if LFSDETIE bit is set in the SAI_xIM register. This flag is cleared when the software sets bit CLFSDET in SAI_xCLRFR register
func (r *registerSaibsrType) SetLfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibsrFieldLfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibsrFieldLfsdetMask)
	}
}

const (
	RegisterSaibsrFieldFlvlShift = 16
	RegisterSaibsrFieldFlvlMask  = 0x70000
)

// GetFlvl FIFO level threshold. This bit is read only. The FIFO level threshold flag is managed only by hardware and its setting depends on SAI block configuration (transmitter or receiver mode). If the SAI block is configured as transmitter: If SAI block is configured as receiver:
func (r *registerSaibsrType) GetFlvl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaibsrFieldFlvlMask) >> RegisterSaibsrFieldFlvlShift)
}

// SetFlvl FIFO level threshold. This bit is read only. The FIFO level threshold flag is managed only by hardware and its setting depends on SAI block configuration (transmitter or receiver mode). If the SAI block is configured as transmitter: If SAI block is configured as receiver:
func (r *registerSaibsrType) SetFlvl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibsrFieldFlvlMask)|(uint32(value)<<RegisterSaibsrFieldFlvlShift))
}

// registerSaibclrfrType Clear flag register
type registerSaibclrfrType uint32

const (
	RegisterSaibclrfrFieldCovrudrShift = 0
	RegisterSaibclrfrFieldCovrudrMask  = 0x1
)

// GetCovrudr Clear overrun / underrun. This bit is write only. Programming this bit to 1 clears the OVRUDR flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) GetCovrudr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibclrfrFieldCovrudrMask) != 0
}

// SetCovrudr Clear overrun / underrun. This bit is write only. Programming this bit to 1 clears the OVRUDR flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) SetCovrudr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibclrfrFieldCovrudrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibclrfrFieldCovrudrMask)
	}
}

const (
	RegisterSaibclrfrFieldCmutedetShift = 1
	RegisterSaibclrfrFieldCmutedetMask  = 0x2
)

// GetCmutedet Mute detection flag. This bit is write only. Programming this bit to 1 clears the MUTEDET flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) GetCmutedet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibclrfrFieldCmutedetMask) != 0
}

// SetCmutedet Mute detection flag. This bit is write only. Programming this bit to 1 clears the MUTEDET flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) SetCmutedet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibclrfrFieldCmutedetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibclrfrFieldCmutedetMask)
	}
}

const (
	RegisterSaibclrfrFieldCwckcfgShift = 2
	RegisterSaibclrfrFieldCwckcfgMask  = 0x4
)

// GetCwckcfg Clear wrong clock configuration flag. This bit is write only. Programming this bit to 1 clears the WCKCFG flag in the SAI_xSR register. This bit is used only when the audio block is set as master (MODE[1] = 0) and NODIV = 0 in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) GetCwckcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibclrfrFieldCwckcfgMask) != 0
}

// SetCwckcfg Clear wrong clock configuration flag. This bit is write only. Programming this bit to 1 clears the WCKCFG flag in the SAI_xSR register. This bit is used only when the audio block is set as master (MODE[1] = 0) and NODIV = 0 in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) SetCwckcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibclrfrFieldCwckcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibclrfrFieldCwckcfgMask)
	}
}

const (
	RegisterSaibclrfrFieldCcnrdyShift = 4
	RegisterSaibclrfrFieldCcnrdyMask  = 0x10
)

// GetCcnrdy Clear Codec not ready flag. This bit is write only. Programming this bit to 1 clears the CNRDY flag in the SAI_xSR register. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) GetCcnrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibclrfrFieldCcnrdyMask) != 0
}

// SetCcnrdy Clear Codec not ready flag. This bit is write only. Programming this bit to 1 clears the CNRDY flag in the SAI_xSR register. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) SetCcnrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibclrfrFieldCcnrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibclrfrFieldCcnrdyMask)
	}
}

const (
	RegisterSaibclrfrFieldCafsdetShift = 5
	RegisterSaibclrfrFieldCafsdetMask  = 0x20
)

// GetCafsdet Clear anticipated frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the AFSDET flag in the SAI_xSR register. It is not used in AC97or SPDIF mode. Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) GetCafsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibclrfrFieldCafsdetMask) != 0
}

// SetCafsdet Clear anticipated frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the AFSDET flag in the SAI_xSR register. It is not used in AC97or SPDIF mode. Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) SetCafsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibclrfrFieldCafsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibclrfrFieldCafsdetMask)
	}
}

const (
	RegisterSaibclrfrFieldClfsdetShift = 6
	RegisterSaibclrfrFieldClfsdetMask  = 0x40
)

// GetClfsdet Clear late frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the LFSDET flag in the SAI_xSR register. This bit is not used in AC97or SPDIF mode Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) GetClfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaibclrfrFieldClfsdetMask) != 0
}

// SetClfsdet Clear late frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the LFSDET flag in the SAI_xSR register. This bit is not used in AC97or SPDIF mode Reading this bit always returns the value 0.
func (r *registerSaibclrfrType) SetClfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaibclrfrFieldClfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaibclrfrFieldClfsdetMask)
	}
}

// registerSaibdrType Data register
type registerSaibdrType uint32

const (
	RegisterSaibdrFieldDataShift = 0
	RegisterSaibdrFieldDataMask  = 0xffffffff
)

// GetData Data A write to this register loads the FIFO provided the FIFO is not full. A read from this register empties the FIFO if the FIFO is not empty.
func (r *registerSaibdrType) GetData() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterSaibdrFieldDataMask) >> RegisterSaibdrFieldDataShift)
}

// SetData Data A write to this register loads the FIFO provided the FIFO is not full. A read from this register empties the FIFO if the FIFO is not empty.
func (r *registerSaibdrType) SetData(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaibdrFieldDataMask)|(uint32(value)<<RegisterSaibdrFieldDataShift))
}

// registerSaipdmcrType PDM control register
type registerSaipdmcrType uint32

const (
	RegisterSaipdmcrFieldPdmenShift = 0
	RegisterSaipdmcrFieldPdmenMask  = 0x1
)

// GetPdmen PDM enable
func (r *registerSaipdmcrType) GetPdmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmcrFieldPdmenMask) != 0
}

// SetPdmen PDM enable
func (r *registerSaipdmcrType) SetPdmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaipdmcrFieldPdmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmcrFieldPdmenMask)
	}
}

const (
	RegisterSaipdmcrFieldMicnbrShift = 4
	RegisterSaipdmcrFieldMicnbrMask  = 0x30
)

// GetMicnbr Number of microphones
func (r *registerSaipdmcrType) GetMicnbr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmcrFieldMicnbrMask) >> RegisterSaipdmcrFieldMicnbrShift)
}

// SetMicnbr Number of microphones
func (r *registerSaipdmcrType) SetMicnbr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmcrFieldMicnbrMask)|(uint32(value)<<RegisterSaipdmcrFieldMicnbrShift))
}

const (
	RegisterSaipdmcrFieldCken1Shift = 8
	RegisterSaipdmcrFieldCken1Mask  = 0x100
)

// GetCken1 Clock enable of bitstream clock number 1
func (r *registerSaipdmcrType) GetCken1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmcrFieldCken1Mask) != 0
}

// SetCken1 Clock enable of bitstream clock number 1
func (r *registerSaipdmcrType) SetCken1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaipdmcrFieldCken1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmcrFieldCken1Mask)
	}
}

const (
	RegisterSaipdmcrFieldCken2Shift = 9
	RegisterSaipdmcrFieldCken2Mask  = 0x200
)

// GetCken2 Clock enable of bitstream clock number 2
func (r *registerSaipdmcrType) GetCken2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmcrFieldCken2Mask) != 0
}

// SetCken2 Clock enable of bitstream clock number 2
func (r *registerSaipdmcrType) SetCken2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaipdmcrFieldCken2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmcrFieldCken2Mask)
	}
}

const (
	RegisterSaipdmcrFieldCken3Shift = 10
	RegisterSaipdmcrFieldCken3Mask  = 0x400
)

// GetCken3 Clock enable of bitstream clock number 3
func (r *registerSaipdmcrType) GetCken3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmcrFieldCken3Mask) != 0
}

// SetCken3 Clock enable of bitstream clock number 3
func (r *registerSaipdmcrType) SetCken3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaipdmcrFieldCken3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmcrFieldCken3Mask)
	}
}

const (
	RegisterSaipdmcrFieldCken4Shift = 11
	RegisterSaipdmcrFieldCken4Mask  = 0x800
)

// GetCken4 Clock enable of bitstream clock number 4
func (r *registerSaipdmcrType) GetCken4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmcrFieldCken4Mask) != 0
}

// SetCken4 Clock enable of bitstream clock number 4
func (r *registerSaipdmcrType) SetCken4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSaipdmcrFieldCken4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmcrFieldCken4Mask)
	}
}

// registerSaipdmdlyType PDM delay register
type registerSaipdmdlyType uint32

const (
	RegisterSaipdmdlyFieldDlym1lShift = 0
	RegisterSaipdmdlyFieldDlym1lMask  = 0x7
)

// GetDlym1l Delay line adjust for first microphone of pair 1
func (r *registerSaipdmdlyType) GetDlym1l() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmdlyFieldDlym1lMask) >> RegisterSaipdmdlyFieldDlym1lShift)
}

// SetDlym1l Delay line adjust for first microphone of pair 1
func (r *registerSaipdmdlyType) SetDlym1l(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmdlyFieldDlym1lMask)|(uint32(value)<<RegisterSaipdmdlyFieldDlym1lShift))
}

const (
	RegisterSaipdmdlyFieldDlym1rShift = 4
	RegisterSaipdmdlyFieldDlym1rMask  = 0x70
)

// GetDlym1r Delay line adjust for second microphone of pair 1
func (r *registerSaipdmdlyType) GetDlym1r() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmdlyFieldDlym1rMask) >> RegisterSaipdmdlyFieldDlym1rShift)
}

// SetDlym1r Delay line adjust for second microphone of pair 1
func (r *registerSaipdmdlyType) SetDlym1r(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmdlyFieldDlym1rMask)|(uint32(value)<<RegisterSaipdmdlyFieldDlym1rShift))
}

const (
	RegisterSaipdmdlyFieldDlym2lShift = 8
	RegisterSaipdmdlyFieldDlym2lMask  = 0x700
)

// GetDlym2l Delay line for first microphone of pair 2
func (r *registerSaipdmdlyType) GetDlym2l() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmdlyFieldDlym2lMask) >> RegisterSaipdmdlyFieldDlym2lShift)
}

// SetDlym2l Delay line for first microphone of pair 2
func (r *registerSaipdmdlyType) SetDlym2l(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmdlyFieldDlym2lMask)|(uint32(value)<<RegisterSaipdmdlyFieldDlym2lShift))
}

const (
	RegisterSaipdmdlyFieldDlym2rShift = 12
	RegisterSaipdmdlyFieldDlym2rMask  = 0x7000
)

// GetDlym2r Delay line for second microphone of pair 2
func (r *registerSaipdmdlyType) GetDlym2r() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmdlyFieldDlym2rMask) >> RegisterSaipdmdlyFieldDlym2rShift)
}

// SetDlym2r Delay line for second microphone of pair 2
func (r *registerSaipdmdlyType) SetDlym2r(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmdlyFieldDlym2rMask)|(uint32(value)<<RegisterSaipdmdlyFieldDlym2rShift))
}

const (
	RegisterSaipdmdlyFieldDlym3lShift = 16
	RegisterSaipdmdlyFieldDlym3lMask  = 0x70000
)

// GetDlym3l Delay line for first microphone of pair 3
func (r *registerSaipdmdlyType) GetDlym3l() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmdlyFieldDlym3lMask) >> RegisterSaipdmdlyFieldDlym3lShift)
}

// SetDlym3l Delay line for first microphone of pair 3
func (r *registerSaipdmdlyType) SetDlym3l(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmdlyFieldDlym3lMask)|(uint32(value)<<RegisterSaipdmdlyFieldDlym3lShift))
}

const (
	RegisterSaipdmdlyFieldDlym3rShift = 20
	RegisterSaipdmdlyFieldDlym3rMask  = 0x700000
)

// GetDlym3r Delay line for second microphone of pair 3
func (r *registerSaipdmdlyType) GetDlym3r() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmdlyFieldDlym3rMask) >> RegisterSaipdmdlyFieldDlym3rShift)
}

// SetDlym3r Delay line for second microphone of pair 3
func (r *registerSaipdmdlyType) SetDlym3r(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmdlyFieldDlym3rMask)|(uint32(value)<<RegisterSaipdmdlyFieldDlym3rShift))
}

const (
	RegisterSaipdmdlyFieldDlym4lShift = 24
	RegisterSaipdmdlyFieldDlym4lMask  = 0x7000000
)

// GetDlym4l Delay line for first microphone of pair 4
func (r *registerSaipdmdlyType) GetDlym4l() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmdlyFieldDlym4lMask) >> RegisterSaipdmdlyFieldDlym4lShift)
}

// SetDlym4l Delay line for first microphone of pair 4
func (r *registerSaipdmdlyType) SetDlym4l(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmdlyFieldDlym4lMask)|(uint32(value)<<RegisterSaipdmdlyFieldDlym4lShift))
}

const (
	RegisterSaipdmdlyFieldDlym4rShift = 28
	RegisterSaipdmdlyFieldDlym4rMask  = 0x70000000
)

// GetDlym4r Delay line for second microphone of pair 4
func (r *registerSaipdmdlyType) GetDlym4r() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSaipdmdlyFieldDlym4rMask) >> RegisterSaipdmdlyFieldDlym4rShift)
}

// SetDlym4r Delay line for second microphone of pair 4
func (r *registerSaipdmdlyType) SetDlym4r(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSaipdmdlyFieldDlym4rMask)|(uint32(value)<<RegisterSaipdmdlyFieldDlym4rShift))
}
