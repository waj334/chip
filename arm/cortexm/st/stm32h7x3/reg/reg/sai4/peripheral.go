package sai4

import (
	"unsafe"
	"volatile"
)

var (
	Sai1 = (*_sai4)(unsafe.Pointer(uintptr(0x40015800)))
	Sai2 = (*_sai4)(unsafe.Pointer(uintptr(0x40015c00)))
	Sai3 = (*_sai4)(unsafe.Pointer(uintptr(0x40016000)))
	Sai4 = (*_sai4)(unsafe.Pointer(uintptr(0x58005400)))
)

type _sai4 struct {
	Sai_gcr    registerSai_gcrType
	Sai_acr1   registerSai_acr1Type
	Sai_acr2   registerSai_acr2Type
	Sai_afrcr  registerSai_afrcrType
	Sai_aslotr registerSai_aslotrType
	Sai_aim    registerSai_aimType
	Sai_asr    registerSai_asrType
	Sai_aclrfr registerSai_aclrfrType
	Sai_adr    registerSai_adrType
	Sai_bcr1   registerSai_bcr1Type
	Sai_bcr2   registerSai_bcr2Type
	Sai_bfrcr  registerSai_bfrcrType
	Sai_bslotr registerSai_bslotrType
	Sai_bim    registerSai_bimType
	Sai_bsr    registerSai_bsrType
	Sai_bclrfr registerSai_bclrfrType
	Sai_bdr    registerSai_bdrType
	Sai_pdmcr  registerSai_pdmcrType
	Sai_pdmdly registerSai_pdmdlyType
}

// registerSai_gcrType Global configuration register
type registerSai_gcrType uint32

const (
	RegisterSai_gcrFieldSyncoutShift = 4
	RegisterSai_gcrFieldSyncoutMask  = 0x30
)

// GetSyncout Synchronization outputs These bits are set and cleared by software.
func (r *registerSai_gcrType) GetSyncout() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_gcrFieldSyncoutMask) >> RegisterSai_gcrFieldSyncoutShift)
}

// SetSyncout Synchronization outputs These bits are set and cleared by software.
func (r *registerSai_gcrType) SetSyncout(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_gcrFieldSyncoutMask)|(uint32(value)<<RegisterSai_gcrFieldSyncoutShift))
}

const (
	RegisterSai_gcrFieldSyncinShift = 0
	RegisterSai_gcrFieldSyncinMask  = 0x3
)

// GetSyncin Synchronization inputs
func (r *registerSai_gcrType) GetSyncin() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_gcrFieldSyncinMask) >> RegisterSai_gcrFieldSyncinShift)
}

// SetSyncin Synchronization inputs
func (r *registerSai_gcrType) SetSyncin(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_gcrFieldSyncinMask)|(uint32(value)<<RegisterSai_gcrFieldSyncinShift))
}

// registerSai_acr1Type Configuration register 1
type registerSai_acr1Type uint32

const (
	RegisterSai_acr1FieldModeShift = 0
	RegisterSai_acr1FieldModeMask  = 0x3
)

// GetMode SAIx audio block mode immediately
func (r *registerSai_acr1Type) GetMode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldModeMask) >> RegisterSai_acr1FieldModeShift)
}

// SetMode SAIx audio block mode immediately
func (r *registerSai_acr1Type) SetMode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldModeMask)|(uint32(value)<<RegisterSai_acr1FieldModeShift))
}

const (
	RegisterSai_acr1FieldPrtcfgShift = 2
	RegisterSai_acr1FieldPrtcfgMask  = 0xc
)

// GetPrtcfg Protocol configuration. These bits are set and cleared by software. These bits have to be configured when the audio block is disabled.
func (r *registerSai_acr1Type) GetPrtcfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldPrtcfgMask) >> RegisterSai_acr1FieldPrtcfgShift)
}

// SetPrtcfg Protocol configuration. These bits are set and cleared by software. These bits have to be configured when the audio block is disabled.
func (r *registerSai_acr1Type) SetPrtcfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldPrtcfgMask)|(uint32(value)<<RegisterSai_acr1FieldPrtcfgShift))
}

const (
	RegisterSai_acr1FieldDsShift = 5
	RegisterSai_acr1FieldDsMask  = 0xe0
)

// GetDs Data size. These bits are set and cleared by software. These bits are ignored when the SPDIF protocols are selected (bit PRTCFG[1:0]), because the frame and the data size are fixed in such case. When the companding mode is selected through COMP[1:0] bits, DS[1:0] are ignored since the data size is fixed to 8 bits by the algorithm. These bits must be configured when the audio block is disabled.
func (r *registerSai_acr1Type) GetDs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldDsMask) >> RegisterSai_acr1FieldDsShift)
}

// SetDs Data size. These bits are set and cleared by software. These bits are ignored when the SPDIF protocols are selected (bit PRTCFG[1:0]), because the frame and the data size are fixed in such case. When the companding mode is selected through COMP[1:0] bits, DS[1:0] are ignored since the data size is fixed to 8 bits by the algorithm. These bits must be configured when the audio block is disabled.
func (r *registerSai_acr1Type) SetDs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldDsMask)|(uint32(value)<<RegisterSai_acr1FieldDsShift))
}

const (
	RegisterSai_acr1FieldLsbfirstShift = 8
	RegisterSai_acr1FieldLsbfirstMask  = 0x100
)

// GetLsbfirst Least significant bit first. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in AC97 audio protocol since AC97 data are always transferred with the MSB first. This bit has no meaning in SPDIF audio protocol since in SPDIF data are always transferred with LSB first.
func (r *registerSai_acr1Type) GetLsbfirst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldLsbfirstMask) != 0
}

// SetLsbfirst Least significant bit first. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in AC97 audio protocol since AC97 data are always transferred with the MSB first. This bit has no meaning in SPDIF audio protocol since in SPDIF data are always transferred with LSB first.
func (r *registerSai_acr1Type) SetLsbfirst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr1FieldLsbfirstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldLsbfirstMask)
	}
}

const (
	RegisterSai_acr1FieldCkstrShift = 9
	RegisterSai_acr1FieldCkstrMask  = 0x200
)

// GetCkstr Clock strobing edge. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in SPDIF audio protocol.
func (r *registerSai_acr1Type) GetCkstr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldCkstrMask) != 0
}

// SetCkstr Clock strobing edge. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in SPDIF audio protocol.
func (r *registerSai_acr1Type) SetCkstr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr1FieldCkstrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldCkstrMask)
	}
}

const (
	RegisterSai_acr1FieldSyncenShift = 10
	RegisterSai_acr1FieldSyncenMask  = 0xc00
)

// GetSyncen Synchronization enable. These bits are set and cleared by software. They must be configured when the audio sub-block is disabled. Note: The audio sub-block should be configured as asynchronous when SPDIF mode is enabled.
func (r *registerSai_acr1Type) GetSyncen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldSyncenMask) >> RegisterSai_acr1FieldSyncenShift)
}

// SetSyncen Synchronization enable. These bits are set and cleared by software. They must be configured when the audio sub-block is disabled. Note: The audio sub-block should be configured as asynchronous when SPDIF mode is enabled.
func (r *registerSai_acr1Type) SetSyncen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldSyncenMask)|(uint32(value)<<RegisterSai_acr1FieldSyncenShift))
}

const (
	RegisterSai_acr1FieldMonoShift = 12
	RegisterSai_acr1FieldMonoMask  = 0x1000
)

// GetMono Mono mode. This bit is set and cleared by software. It is meaningful only when the number of slots is equal to 2. When the mono mode is selected, slot 0 data are duplicated on slot 1 when the audio block operates as a transmitter. In reception mode, the slot1 is discarded and only the data received from slot 0 are stored. Refer to Section: Mono/stereo mode for more details.
func (r *registerSai_acr1Type) GetMono() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldMonoMask) != 0
}

// SetMono Mono mode. This bit is set and cleared by software. It is meaningful only when the number of slots is equal to 2. When the mono mode is selected, slot 0 data are duplicated on slot 1 when the audio block operates as a transmitter. In reception mode, the slot1 is discarded and only the data received from slot 0 are stored. Refer to Section: Mono/stereo mode for more details.
func (r *registerSai_acr1Type) SetMono(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr1FieldMonoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldMonoMask)
	}
}

const (
	RegisterSai_acr1FieldOutdrivShift = 13
	RegisterSai_acr1FieldOutdrivMask  = 0x2000
)

// GetOutdriv Output drive. This bit is set and cleared by software. Note: This bit has to be set before enabling the audio block and after the audio block configuration.
func (r *registerSai_acr1Type) GetOutdriv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldOutdrivMask) != 0
}

// SetOutdriv Output drive. This bit is set and cleared by software. Note: This bit has to be set before enabling the audio block and after the audio block configuration.
func (r *registerSai_acr1Type) SetOutdriv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr1FieldOutdrivMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldOutdrivMask)
	}
}

const (
	RegisterSai_acr1FieldSaixenShift = 16
	RegisterSai_acr1FieldSaixenMask  = 0x10000
)

// GetSaixen Audio block enable where x is A or B. This bit is set by software. To switch off the audio block, the application software must program this bit to 0 and poll the bit till it reads back 0, meaning that the block is completely disabled. Before setting this bit to 1, check that it is set to 0, otherwise the enable command will not be taken into account. This bit allows to control the state of SAIx audio block. If it is disabled when an audio frame transfer is ongoing, the ongoing transfer completes and the cell is fully disabled at the end of this audio frame transfer. Note: When SAIx block is configured in master mode, the clock must be present on the input of SAIx before setting SAIXEN bit.
func (r *registerSai_acr1Type) GetSaixen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldSaixenMask) != 0
}

// SetSaixen Audio block enable where x is A or B. This bit is set by software. To switch off the audio block, the application software must program this bit to 0 and poll the bit till it reads back 0, meaning that the block is completely disabled. Before setting this bit to 1, check that it is set to 0, otherwise the enable command will not be taken into account. This bit allows to control the state of SAIx audio block. If it is disabled when an audio frame transfer is ongoing, the ongoing transfer completes and the cell is fully disabled at the end of this audio frame transfer. Note: When SAIx block is configured in master mode, the clock must be present on the input of SAIx before setting SAIXEN bit.
func (r *registerSai_acr1Type) SetSaixen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr1FieldSaixenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldSaixenMask)
	}
}

const (
	RegisterSai_acr1FieldDmaenShift = 17
	RegisterSai_acr1FieldDmaenMask  = 0x20000
)

// GetDmaen DMA enable. This bit is set and cleared by software. Note: Since the audio block defaults to operate as a transmitter after reset, the MODE[1:0] bits must be configured before setting DMAEN to avoid a DMA request in receiver mode.
func (r *registerSai_acr1Type) GetDmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldDmaenMask) != 0
}

// SetDmaen DMA enable. This bit is set and cleared by software. Note: Since the audio block defaults to operate as a transmitter after reset, the MODE[1:0] bits must be configured before setting DMAEN to avoid a DMA request in receiver mode.
func (r *registerSai_acr1Type) SetDmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr1FieldDmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldDmaenMask)
	}
}

const (
	RegisterSai_acr1FieldNomckShift = 19
	RegisterSai_acr1FieldNomckMask  = 0x80000
)

// GetNomck No divider
func (r *registerSai_acr1Type) GetNomck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldNomckMask) != 0
}

// SetNomck No divider
func (r *registerSai_acr1Type) SetNomck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr1FieldNomckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldNomckMask)
	}
}

const (
	RegisterSai_acr1FieldMckdivShift = 20
	RegisterSai_acr1FieldMckdivMask  = 0xf00000
)

// GetMckdiv Master clock divider. These bits are set and cleared by software. These bits are meaningless when the audio block operates in slave mode. They have to be configured when the audio block is disabled. Others: the master clock frequency is calculated accordingly to the following formula:
func (r *registerSai_acr1Type) GetMckdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldMckdivMask) >> RegisterSai_acr1FieldMckdivShift)
}

// SetMckdiv Master clock divider. These bits are set and cleared by software. These bits are meaningless when the audio block operates in slave mode. They have to be configured when the audio block is disabled. Others: the master clock frequency is calculated accordingly to the following formula:
func (r *registerSai_acr1Type) SetMckdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldMckdivMask)|(uint32(value)<<RegisterSai_acr1FieldMckdivShift))
}

const (
	RegisterSai_acr1FieldOsrShift = 26
	RegisterSai_acr1FieldOsrMask  = 0x4000000
)

// GetOsr Oversampling ratio for master clock
func (r *registerSai_acr1Type) GetOsr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr1FieldOsrMask) != 0
}

// SetOsr Oversampling ratio for master clock
func (r *registerSai_acr1Type) SetOsr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr1FieldOsrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr1FieldOsrMask)
	}
}

// registerSai_acr2Type Configuration register 2
type registerSai_acr2Type uint32

const (
	RegisterSai_acr2FieldFthShift = 0
	RegisterSai_acr2FieldFthMask  = 0x7
)

// GetFth FIFO threshold. This bit is set and cleared by software.
func (r *registerSai_acr2Type) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr2FieldFthMask) >> RegisterSai_acr2FieldFthShift)
}

// SetFth FIFO threshold. This bit is set and cleared by software.
func (r *registerSai_acr2Type) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr2FieldFthMask)|(uint32(value)<<RegisterSai_acr2FieldFthShift))
}

const (
	RegisterSai_acr2FieldFflushShift = 3
	RegisterSai_acr2FieldFflushMask  = 0x8
)

// SetFflush FIFO flush. This bit is set by software. It is always read as 0. This bit should be configured when the SAI is disabled.
func (r *registerSai_acr2Type) SetFflush(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr2FieldFflushMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr2FieldFflushMask)
	}
}

const (
	RegisterSai_acr2FieldTrisShift = 4
	RegisterSai_acr2FieldTrisMask  = 0x10
)

// GetTris Tristate management on data line. This bit is set and cleared by software. It is meaningful only if the audio block is configured as a transmitter. This bit is not used when the audio block is configured in SPDIF mode. It should be configured when SAI is disabled. Refer to Section: Output data line management on an inactive slot for more details.
func (r *registerSai_acr2Type) GetTris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr2FieldTrisMask) != 0
}

// SetTris Tristate management on data line. This bit is set and cleared by software. It is meaningful only if the audio block is configured as a transmitter. This bit is not used when the audio block is configured in SPDIF mode. It should be configured when SAI is disabled. Refer to Section: Output data line management on an inactive slot for more details.
func (r *registerSai_acr2Type) SetTris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr2FieldTrisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr2FieldTrisMask)
	}
}

const (
	RegisterSai_acr2FieldMuteShift = 5
	RegisterSai_acr2FieldMuteMask  = 0x20
)

// GetMute Mute. This bit is set and cleared by software. It is meaningful only when the audio block operates as a transmitter. The MUTE value is linked to value of MUTEVAL if the number of slots is lower or equal to 2, or equal to 0 if it is greater than 2. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSai_acr2Type) GetMute() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr2FieldMuteMask) != 0
}

// SetMute Mute. This bit is set and cleared by software. It is meaningful only when the audio block operates as a transmitter. The MUTE value is linked to value of MUTEVAL if the number of slots is lower or equal to 2, or equal to 0 if it is greater than 2. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSai_acr2Type) SetMute(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr2FieldMuteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr2FieldMuteMask)
	}
}

const (
	RegisterSai_acr2FieldMutevalShift = 6
	RegisterSai_acr2FieldMutevalMask  = 0x40
)

// GetMuteval Mute value. This bit is set and cleared by software.It must be written before enabling the audio block: SAIXEN. This bit is meaningful only when the audio block operates as a transmitter, the number of slots is lower or equal to 2 and the MUTE bit is set. If more slots are declared, the bit value sent during the transmission in mute mode is equal to 0, whatever the value of MUTEVAL. if the number of slot is lower or equal to 2 and MUTEVAL = 1, the MUTE value transmitted for each slot is the one sent during the previous frame. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSai_acr2Type) GetMuteval() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr2FieldMutevalMask) != 0
}

// SetMuteval Mute value. This bit is set and cleared by software.It must be written before enabling the audio block: SAIXEN. This bit is meaningful only when the audio block operates as a transmitter, the number of slots is lower or equal to 2 and the MUTE bit is set. If more slots are declared, the bit value sent during the transmission in mute mode is equal to 0, whatever the value of MUTEVAL. if the number of slot is lower or equal to 2 and MUTEVAL = 1, the MUTE value transmitted for each slot is the one sent during the previous frame. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSai_acr2Type) SetMuteval(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr2FieldMutevalMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr2FieldMutevalMask)
	}
}

const (
	RegisterSai_acr2FieldMutecntShift = 7
	RegisterSai_acr2FieldMutecntMask  = 0x1f80
)

// GetMutecnt Mute counter. These bits are set and cleared by software. They are used only in reception mode. The value set in these bits is compared to the number of consecutive mute frames detected in reception. When the number of mute frames is equal to this value, the flag MUTEDET will be set and an interrupt will be generated if bit MUTEDETIE is set. Refer to Section: Mute mode for more details.
func (r *registerSai_acr2Type) GetMutecnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr2FieldMutecntMask) >> RegisterSai_acr2FieldMutecntShift)
}

// SetMutecnt Mute counter. These bits are set and cleared by software. They are used only in reception mode. The value set in these bits is compared to the number of consecutive mute frames detected in reception. When the number of mute frames is equal to this value, the flag MUTEDET will be set and an interrupt will be generated if bit MUTEDETIE is set. Refer to Section: Mute mode for more details.
func (r *registerSai_acr2Type) SetMutecnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr2FieldMutecntMask)|(uint32(value)<<RegisterSai_acr2FieldMutecntShift))
}

const (
	RegisterSai_acr2FieldCplShift = 13
	RegisterSai_acr2FieldCplMask  = 0x2000
)

// GetCpl Complement bit. This bit is set and cleared by software. It defines the type of complement to be used for companding mode Note: This bit has effect only when the companding mode is -Law algorithm or A-Law algorithm.
func (r *registerSai_acr2Type) GetCpl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr2FieldCplMask) != 0
}

// SetCpl Complement bit. This bit is set and cleared by software. It defines the type of complement to be used for companding mode Note: This bit has effect only when the companding mode is -Law algorithm or A-Law algorithm.
func (r *registerSai_acr2Type) SetCpl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_acr2FieldCplMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr2FieldCplMask)
	}
}

const (
	RegisterSai_acr2FieldCompShift = 14
	RegisterSai_acr2FieldCompMask  = 0xc000
)

// GetComp Companding mode. These bits are set and cleared by software. The -Law and the A-Law log are a part of the CCITT G.711 recommendation, the type of complement that will be used depends on CPL bit. The data expansion or data compression are determined by the state of bit MODE[0]. The data compression is applied if the audio block is configured as a transmitter. The data expansion is automatically applied when the audio block is configured as a receiver. Refer to Section: Companding mode for more details. Note: Companding mode is applicable only when TDM is selected.
func (r *registerSai_acr2Type) GetComp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_acr2FieldCompMask) >> RegisterSai_acr2FieldCompShift)
}

// SetComp Companding mode. These bits are set and cleared by software. The -Law and the A-Law log are a part of the CCITT G.711 recommendation, the type of complement that will be used depends on CPL bit. The data expansion or data compression are determined by the state of bit MODE[0]. The data compression is applied if the audio block is configured as a transmitter. The data expansion is automatically applied when the audio block is configured as a receiver. Refer to Section: Companding mode for more details. Note: Companding mode is applicable only when TDM is selected.
func (r *registerSai_acr2Type) SetComp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_acr2FieldCompMask)|(uint32(value)<<RegisterSai_acr2FieldCompShift))
}

// registerSai_afrcrType This register has no meaning in AC97 and SPDIF audio protocol
type registerSai_afrcrType uint32

const (
	RegisterSai_afrcrFieldFrlShift = 0
	RegisterSai_afrcrFieldFrlMask  = 0xff
)

// GetFrl Frame length. These bits are set and cleared by software. They define the audio frame length expressed in number of SCK clock cycles: the number of bits in the frame is equal to FRL[7:0] + 1. The minimum number of bits to transfer in an audio frame must be equal to 8, otherwise the audio block will behaves in an unexpected way. This is the case when the data size is 8 bits and only one slot 0 is defined in NBSLOT[4:0] of SAI_xSLOTR register (NBSLOT[3:0] = 0000). In master mode, if the master clock (available on MCLK_x pin) is used, the frame length should be aligned with a number equal to a power of 2, ranging from 8 to 256. When the master clock is not used (NODIV = 1), it is recommended to program the frame length to an value ranging from 8 to 256. These bits are meaningless and are not used in AC97 or SPDIF audio block configuration.
func (r *registerSai_afrcrType) GetFrl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_afrcrFieldFrlMask) >> RegisterSai_afrcrFieldFrlShift)
}

// SetFrl Frame length. These bits are set and cleared by software. They define the audio frame length expressed in number of SCK clock cycles: the number of bits in the frame is equal to FRL[7:0] + 1. The minimum number of bits to transfer in an audio frame must be equal to 8, otherwise the audio block will behaves in an unexpected way. This is the case when the data size is 8 bits and only one slot 0 is defined in NBSLOT[4:0] of SAI_xSLOTR register (NBSLOT[3:0] = 0000). In master mode, if the master clock (available on MCLK_x pin) is used, the frame length should be aligned with a number equal to a power of 2, ranging from 8 to 256. When the master clock is not used (NODIV = 1), it is recommended to program the frame length to an value ranging from 8 to 256. These bits are meaningless and are not used in AC97 or SPDIF audio block configuration.
func (r *registerSai_afrcrType) SetFrl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_afrcrFieldFrlMask)|(uint32(value)<<RegisterSai_afrcrFieldFrlShift))
}

const (
	RegisterSai_afrcrFieldFsallShift = 8
	RegisterSai_afrcrFieldFsallMask  = 0x7f00
)

// GetFsall Frame synchronization active level length. These bits are set and cleared by software. They specify the length in number of bit clock (SCK) + 1 (FSALL[6:0] + 1) of the active level of the FS signal in the audio frame These bits are meaningless and are not used in AC97 or SPDIF audio block configuration. They must be configured when the audio block is disabled.
func (r *registerSai_afrcrType) GetFsall() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_afrcrFieldFsallMask) >> RegisterSai_afrcrFieldFsallShift)
}

// SetFsall Frame synchronization active level length. These bits are set and cleared by software. They specify the length in number of bit clock (SCK) + 1 (FSALL[6:0] + 1) of the active level of the FS signal in the audio frame These bits are meaningless and are not used in AC97 or SPDIF audio block configuration. They must be configured when the audio block is disabled.
func (r *registerSai_afrcrType) SetFsall(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_afrcrFieldFsallMask)|(uint32(value)<<RegisterSai_afrcrFieldFsallShift))
}

const (
	RegisterSai_afrcrFieldFsdefShift = 16
	RegisterSai_afrcrFieldFsdefMask  = 0x10000
)

// GetFsdef Frame synchronization definition. This bit is set and cleared by software. When the bit is set, the number of slots defined in the SAI_xSLOTR register has to be even. It means that half of this number of slots will be dedicated to the left channel and the other slots for the right channel (e.g: this bit has to be set for I2S or MSB/LSB-justified protocols...). This bit is meaningless and is not used in AC97 or SPDIF audio block configuration. It must be configured when the audio block is disabled.
func (r *registerSai_afrcrType) GetFsdef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_afrcrFieldFsdefMask) != 0
}

const (
	RegisterSai_afrcrFieldFspolShift = 17
	RegisterSai_afrcrFieldFspolMask  = 0x20000
)

// GetFspol Frame synchronization polarity. This bit is set and cleared by software. It is used to configure the level of the start of frame on the FS signal. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSai_afrcrType) GetFspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_afrcrFieldFspolMask) != 0
}

// SetFspol Frame synchronization polarity. This bit is set and cleared by software. It is used to configure the level of the start of frame on the FS signal. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSai_afrcrType) SetFspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_afrcrFieldFspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_afrcrFieldFspolMask)
	}
}

const (
	RegisterSai_afrcrFieldFsoffShift = 18
	RegisterSai_afrcrFieldFsoffMask  = 0x40000
)

// GetFsoff Frame synchronization offset. This bit is set and cleared by software. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSai_afrcrType) GetFsoff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_afrcrFieldFsoffMask) != 0
}

// SetFsoff Frame synchronization offset. This bit is set and cleared by software. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSai_afrcrType) SetFsoff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_afrcrFieldFsoffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_afrcrFieldFsoffMask)
	}
}

// registerSai_aslotrType This register has no meaning in AC97 and SPDIF audio protocol
type registerSai_aslotrType uint32

const (
	RegisterSai_aslotrFieldFboffShift = 0
	RegisterSai_aslotrFieldFboffMask  = 0x1f
)

// GetFboff First bit offset These bits are set and cleared by software. The value set in this bitfield defines the position of the first data transfer bit in the slot. It represents an offset value. In transmission mode, the bits outside the data field are forced to 0. In reception mode, the extra received bits are discarded. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_aslotrType) GetFboff() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_aslotrFieldFboffMask) >> RegisterSai_aslotrFieldFboffShift)
}

// SetFboff First bit offset These bits are set and cleared by software. The value set in this bitfield defines the position of the first data transfer bit in the slot. It represents an offset value. In transmission mode, the bits outside the data field are forced to 0. In reception mode, the extra received bits are discarded. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_aslotrType) SetFboff(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_aslotrFieldFboffMask)|(uint32(value)<<RegisterSai_aslotrFieldFboffShift))
}

const (
	RegisterSai_aslotrFieldSlotszShift = 6
	RegisterSai_aslotrFieldSlotszMask  = 0xc0
)

// GetSlotsz Slot size This bits is set and cleared by software. The slot size must be higher or equal to the data size. If this condition is not respected, the behavior of the SAI will be undetermined. Refer to Section: Output data line management on an inactive slot for information on how to drive SD line. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_aslotrType) GetSlotsz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_aslotrFieldSlotszMask) >> RegisterSai_aslotrFieldSlotszShift)
}

// SetSlotsz Slot size This bits is set and cleared by software. The slot size must be higher or equal to the data size. If this condition is not respected, the behavior of the SAI will be undetermined. Refer to Section: Output data line management on an inactive slot for information on how to drive SD line. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_aslotrType) SetSlotsz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_aslotrFieldSlotszMask)|(uint32(value)<<RegisterSai_aslotrFieldSlotszShift))
}

const (
	RegisterSai_aslotrFieldNbslotShift = 8
	RegisterSai_aslotrFieldNbslotMask  = 0xf00
)

// GetNbslot Number of slots in an audio frame. These bits are set and cleared by software. The value set in this bitfield represents the number of slots + 1 in the audio frame (including the number of inactive slots). The maximum number of slots is 16. The number of slots should be even if FSDEF bit in the SAI_xFRCR register is set. The number of slots must be configured when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_aslotrType) GetNbslot() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_aslotrFieldNbslotMask) >> RegisterSai_aslotrFieldNbslotShift)
}

// SetNbslot Number of slots in an audio frame. These bits are set and cleared by software. The value set in this bitfield represents the number of slots + 1 in the audio frame (including the number of inactive slots). The maximum number of slots is 16. The number of slots should be even if FSDEF bit in the SAI_xFRCR register is set. The number of slots must be configured when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_aslotrType) SetNbslot(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_aslotrFieldNbslotMask)|(uint32(value)<<RegisterSai_aslotrFieldNbslotShift))
}

const (
	RegisterSai_aslotrFieldSlotenShift = 16
	RegisterSai_aslotrFieldSlotenMask  = 0xffff0000
)

// GetSloten Slot enable. These bits are set and cleared by software. Each SLOTEN bit corresponds to a slot position from 0 to 15 (maximum 16 slots). The slot must be enabled when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_aslotrType) GetSloten() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSai_aslotrFieldSlotenMask) >> RegisterSai_aslotrFieldSlotenShift)
}

// SetSloten Slot enable. These bits are set and cleared by software. Each SLOTEN bit corresponds to a slot position from 0 to 15 (maximum 16 slots). The slot must be enabled when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_aslotrType) SetSloten(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_aslotrFieldSlotenMask)|(uint32(value)<<RegisterSai_aslotrFieldSlotenShift))
}

// registerSai_aimType Interrupt mask register 2
type registerSai_aimType uint32

const (
	RegisterSai_aimFieldOvrudrieShift = 0
	RegisterSai_aimFieldOvrudrieMask  = 0x1
)

// GetOvrudrie Overrun/underrun interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the OVRUDR bit in the SAI_xSR register is set.
func (r *registerSai_aimType) GetOvrudrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aimFieldOvrudrieMask) != 0
}

// SetOvrudrie Overrun/underrun interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the OVRUDR bit in the SAI_xSR register is set.
func (r *registerSai_aimType) SetOvrudrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aimFieldOvrudrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aimFieldOvrudrieMask)
	}
}

const (
	RegisterSai_aimFieldMutedetieShift = 1
	RegisterSai_aimFieldMutedetieMask  = 0x2
)

// GetMutedetie Mute detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the MUTEDET bit in the SAI_xSR register is set. This bit has a meaning only if the audio block is configured in receiver mode.
func (r *registerSai_aimType) GetMutedetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aimFieldMutedetieMask) != 0
}

// SetMutedetie Mute detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the MUTEDET bit in the SAI_xSR register is set. This bit has a meaning only if the audio block is configured in receiver mode.
func (r *registerSai_aimType) SetMutedetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aimFieldMutedetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aimFieldMutedetieMask)
	}
}

const (
	RegisterSai_aimFieldWckcfgieShift = 2
	RegisterSai_aimFieldWckcfgieMask  = 0x4
)

// GetWckcfgie Wrong clock configuration interrupt enable. This bit is set and cleared by software. This bit is taken into account only if the audio block is configured as a master (MODE[1] = 0) and NODIV = 0. It generates an interrupt if the WCKCFG flag in the SAI_xSR register is set. Note: This bit is used only in TDM mode and is meaningless in other modes.
func (r *registerSai_aimType) GetWckcfgie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aimFieldWckcfgieMask) != 0
}

// SetWckcfgie Wrong clock configuration interrupt enable. This bit is set and cleared by software. This bit is taken into account only if the audio block is configured as a master (MODE[1] = 0) and NODIV = 0. It generates an interrupt if the WCKCFG flag in the SAI_xSR register is set. Note: This bit is used only in TDM mode and is meaningless in other modes.
func (r *registerSai_aimType) SetWckcfgie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aimFieldWckcfgieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aimFieldWckcfgieMask)
	}
}

const (
	RegisterSai_aimFieldFreqieShift = 3
	RegisterSai_aimFieldFreqieMask  = 0x8
)

// GetFreqie FIFO request interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the FREQ bit in the SAI_xSR register is set. Since the audio block defaults to operate as a transmitter after reset, the MODE bit must be configured before setting FREQIE to avoid a parasitic interruption in receiver mode,
func (r *registerSai_aimType) GetFreqie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aimFieldFreqieMask) != 0
}

// SetFreqie FIFO request interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the FREQ bit in the SAI_xSR register is set. Since the audio block defaults to operate as a transmitter after reset, the MODE bit must be configured before setting FREQIE to avoid a parasitic interruption in receiver mode,
func (r *registerSai_aimType) SetFreqie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aimFieldFreqieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aimFieldFreqieMask)
	}
}

const (
	RegisterSai_aimFieldCnrdyieShift = 4
	RegisterSai_aimFieldCnrdyieMask  = 0x10
)

// GetCnrdyie Codec not ready interrupt enable (AC97). This bit is set and cleared by software. When the interrupt is enabled, the audio block detects in the slot 0 (tag0) of the AC97 frame if the Codec connected to this line is ready or not. If it is not ready, the CNRDY flag in the SAI_xSR register is set and an interruption i generated. This bit has a meaning only if the AC97 mode is selected through PRTCFG[1:0] bits and the audio block is operates as a receiver.
func (r *registerSai_aimType) GetCnrdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aimFieldCnrdyieMask) != 0
}

// SetCnrdyie Codec not ready interrupt enable (AC97). This bit is set and cleared by software. When the interrupt is enabled, the audio block detects in the slot 0 (tag0) of the AC97 frame if the Codec connected to this line is ready or not. If it is not ready, the CNRDY flag in the SAI_xSR register is set and an interruption i generated. This bit has a meaning only if the AC97 mode is selected through PRTCFG[1:0] bits and the audio block is operates as a receiver.
func (r *registerSai_aimType) SetCnrdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aimFieldCnrdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aimFieldCnrdyieMask)
	}
}

const (
	RegisterSai_aimFieldAfsdetieShift = 5
	RegisterSai_aimFieldAfsdetieMask  = 0x20
)

// GetAfsdetie Anticipated frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the AFSDET bit in the SAI_xSR register is set. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSai_aimType) GetAfsdetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aimFieldAfsdetieMask) != 0
}

// SetAfsdetie Anticipated frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the AFSDET bit in the SAI_xSR register is set. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSai_aimType) SetAfsdetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aimFieldAfsdetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aimFieldAfsdetieMask)
	}
}

const (
	RegisterSai_aimFieldLfsdetieShift = 6
	RegisterSai_aimFieldLfsdetieMask  = 0x40
)

// GetLfsdetie Late frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the LFSDET bit is set in the SAI_xSR register. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSai_aimType) GetLfsdetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aimFieldLfsdetieMask) != 0
}

// SetLfsdetie Late frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the LFSDET bit is set in the SAI_xSR register. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSai_aimType) SetLfsdetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aimFieldLfsdetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aimFieldLfsdetieMask)
	}
}

// registerSai_asrType Status register
type registerSai_asrType uint32

const (
	RegisterSai_asrFieldOvrudrShift = 0
	RegisterSai_asrFieldOvrudrMask  = 0x1
)

// GetOvrudr Overrun / underrun. This bit is read only. The overrun and underrun conditions can occur only when the audio block is configured as a receiver and a transmitter, respectively. It can generate an interrupt if OVRUDRIE bit is set in SAI_xIM register. This flag is cleared when the software sets COVRUDR bit in SAI_xCLRFR register.
func (r *registerSai_asrType) GetOvrudr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_asrFieldOvrudrMask) != 0
}

// SetOvrudr Overrun / underrun. This bit is read only. The overrun and underrun conditions can occur only when the audio block is configured as a receiver and a transmitter, respectively. It can generate an interrupt if OVRUDRIE bit is set in SAI_xIM register. This flag is cleared when the software sets COVRUDR bit in SAI_xCLRFR register.
func (r *registerSai_asrType) SetOvrudr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_asrFieldOvrudrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_asrFieldOvrudrMask)
	}
}

const (
	RegisterSai_asrFieldMutedetShift = 1
	RegisterSai_asrFieldMutedetMask  = 0x2
)

// GetMutedet Mute detection. This bit is read only. This flag is set if consecutive 0 values are received in each slot of a given audio frame and for a consecutive number of audio frames (set in the MUTECNT bit in the SAI_xCR2 register). It can generate an interrupt if MUTEDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets bit CMUTEDET in the SAI_xCLRFR register.
func (r *registerSai_asrType) GetMutedet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_asrFieldMutedetMask) != 0
}

// SetMutedet Mute detection. This bit is read only. This flag is set if consecutive 0 values are received in each slot of a given audio frame and for a consecutive number of audio frames (set in the MUTECNT bit in the SAI_xCR2 register). It can generate an interrupt if MUTEDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets bit CMUTEDET in the SAI_xCLRFR register.
func (r *registerSai_asrType) SetMutedet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_asrFieldMutedetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_asrFieldMutedetMask)
	}
}

const (
	RegisterSai_asrFieldWckcfgShift = 2
	RegisterSai_asrFieldWckcfgMask  = 0x4
)

// GetWckcfg Wrong clock configuration flag. This bit is read only. This bit is used only when the audio block operates in master mode (MODE[1] = 0) and NODIV = 0. It can generate an interrupt if WCKCFGIE bit is set in SAI_xIM register. This flag is cleared when the software sets CWCKCFG bit in SAI_xCLRFR register.
func (r *registerSai_asrType) GetWckcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_asrFieldWckcfgMask) != 0
}

// SetWckcfg Wrong clock configuration flag. This bit is read only. This bit is used only when the audio block operates in master mode (MODE[1] = 0) and NODIV = 0. It can generate an interrupt if WCKCFGIE bit is set in SAI_xIM register. This flag is cleared when the software sets CWCKCFG bit in SAI_xCLRFR register.
func (r *registerSai_asrType) SetWckcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_asrFieldWckcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_asrFieldWckcfgMask)
	}
}

const (
	RegisterSai_asrFieldFreqShift = 3
	RegisterSai_asrFieldFreqMask  = 0x8
)

// GetFreq FIFO request. This bit is read only. The request depends on the audio block configuration: If the block is configured in transmission mode, the FIFO request is related to a write request operation in the SAI_xDR. If the block configured in reception, the FIFO request related to a read request operation from the SAI_xDR. This flag can generate an interrupt if FREQIE bit is set in SAI_xIM register.
func (r *registerSai_asrType) GetFreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_asrFieldFreqMask) != 0
}

// SetFreq FIFO request. This bit is read only. The request depends on the audio block configuration: If the block is configured in transmission mode, the FIFO request is related to a write request operation in the SAI_xDR. If the block configured in reception, the FIFO request related to a read request operation from the SAI_xDR. This flag can generate an interrupt if FREQIE bit is set in SAI_xIM register.
func (r *registerSai_asrType) SetFreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_asrFieldFreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_asrFieldFreqMask)
	}
}

const (
	RegisterSai_asrFieldCnrdyShift = 4
	RegisterSai_asrFieldCnrdyMask  = 0x10
)

// GetCnrdy Codec not ready. This bit is read only. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register and configured in receiver mode. It can generate an interrupt if CNRDYIE bit is set in SAI_xIM register. This flag is cleared when the software sets CCNRDY bit in SAI_xCLRFR register.
func (r *registerSai_asrType) GetCnrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_asrFieldCnrdyMask) != 0
}

// SetCnrdy Codec not ready. This bit is read only. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register and configured in receiver mode. It can generate an interrupt if CNRDYIE bit is set in SAI_xIM register. This flag is cleared when the software sets CCNRDY bit in SAI_xCLRFR register.
func (r *registerSai_asrType) SetCnrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_asrFieldCnrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_asrFieldCnrdyMask)
	}
}

const (
	RegisterSai_asrFieldAfsdetShift = 5
	RegisterSai_asrFieldAfsdetMask  = 0x20
)

// GetAfsdet Anticipated frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97or SPDIF mode. It can generate an interrupt if AFSDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets CAFSDET bit in SAI_xCLRFR register.
func (r *registerSai_asrType) GetAfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_asrFieldAfsdetMask) != 0
}

// SetAfsdet Anticipated frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97or SPDIF mode. It can generate an interrupt if AFSDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets CAFSDET bit in SAI_xCLRFR register.
func (r *registerSai_asrType) SetAfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_asrFieldAfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_asrFieldAfsdetMask)
	}
}

const (
	RegisterSai_asrFieldLfsdetShift = 6
	RegisterSai_asrFieldLfsdetMask  = 0x40
)

// GetLfsdet Late frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97 or SPDIF mode. It can generate an interrupt if LFSDETIE bit is set in the SAI_xIM register. This flag is cleared when the software sets bit CLFSDET in SAI_xCLRFR register
func (r *registerSai_asrType) GetLfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_asrFieldLfsdetMask) != 0
}

// SetLfsdet Late frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97 or SPDIF mode. It can generate an interrupt if LFSDETIE bit is set in the SAI_xIM register. This flag is cleared when the software sets bit CLFSDET in SAI_xCLRFR register
func (r *registerSai_asrType) SetLfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_asrFieldLfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_asrFieldLfsdetMask)
	}
}

const (
	RegisterSai_asrFieldFlvlShift = 16
	RegisterSai_asrFieldFlvlMask  = 0x70000
)

// GetFlvl FIFO level threshold. This bit is read only. The FIFO level threshold flag is managed only by hardware and its setting depends on SAI block configuration (transmitter or receiver mode). If the SAI block is configured as transmitter: If SAI block is configured as receiver:
func (r *registerSai_asrType) GetFlvl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_asrFieldFlvlMask) >> RegisterSai_asrFieldFlvlShift)
}

// SetFlvl FIFO level threshold. This bit is read only. The FIFO level threshold flag is managed only by hardware and its setting depends on SAI block configuration (transmitter or receiver mode). If the SAI block is configured as transmitter: If SAI block is configured as receiver:
func (r *registerSai_asrType) SetFlvl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_asrFieldFlvlMask)|(uint32(value)<<RegisterSai_asrFieldFlvlShift))
}

// registerSai_aclrfrType Clear flag register
type registerSai_aclrfrType uint32

const (
	RegisterSai_aclrfrFieldCovrudrShift = 0
	RegisterSai_aclrfrFieldCovrudrMask  = 0x1
)

// GetCovrudr Clear overrun / underrun. This bit is write only. Programming this bit to 1 clears the OVRUDR flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) GetCovrudr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aclrfrFieldCovrudrMask) != 0
}

// SetCovrudr Clear overrun / underrun. This bit is write only. Programming this bit to 1 clears the OVRUDR flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) SetCovrudr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aclrfrFieldCovrudrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aclrfrFieldCovrudrMask)
	}
}

const (
	RegisterSai_aclrfrFieldCmutedetShift = 1
	RegisterSai_aclrfrFieldCmutedetMask  = 0x2
)

// GetCmutedet Mute detection flag. This bit is write only. Programming this bit to 1 clears the MUTEDET flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) GetCmutedet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aclrfrFieldCmutedetMask) != 0
}

// SetCmutedet Mute detection flag. This bit is write only. Programming this bit to 1 clears the MUTEDET flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) SetCmutedet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aclrfrFieldCmutedetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aclrfrFieldCmutedetMask)
	}
}

const (
	RegisterSai_aclrfrFieldCwckcfgShift = 2
	RegisterSai_aclrfrFieldCwckcfgMask  = 0x4
)

// GetCwckcfg Clear wrong clock configuration flag. This bit is write only. Programming this bit to 1 clears the WCKCFG flag in the SAI_xSR register. This bit is used only when the audio block is set as master (MODE[1] = 0) and NODIV = 0 in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) GetCwckcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aclrfrFieldCwckcfgMask) != 0
}

// SetCwckcfg Clear wrong clock configuration flag. This bit is write only. Programming this bit to 1 clears the WCKCFG flag in the SAI_xSR register. This bit is used only when the audio block is set as master (MODE[1] = 0) and NODIV = 0 in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) SetCwckcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aclrfrFieldCwckcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aclrfrFieldCwckcfgMask)
	}
}

const (
	RegisterSai_aclrfrFieldCcnrdyShift = 4
	RegisterSai_aclrfrFieldCcnrdyMask  = 0x10
)

// GetCcnrdy Clear Codec not ready flag. This bit is write only. Programming this bit to 1 clears the CNRDY flag in the SAI_xSR register. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) GetCcnrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aclrfrFieldCcnrdyMask) != 0
}

// SetCcnrdy Clear Codec not ready flag. This bit is write only. Programming this bit to 1 clears the CNRDY flag in the SAI_xSR register. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) SetCcnrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aclrfrFieldCcnrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aclrfrFieldCcnrdyMask)
	}
}

const (
	RegisterSai_aclrfrFieldCafsdetShift = 5
	RegisterSai_aclrfrFieldCafsdetMask  = 0x20
)

// GetCafsdet Clear anticipated frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the AFSDET flag in the SAI_xSR register. It is not used in AC97or SPDIF mode. Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) GetCafsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aclrfrFieldCafsdetMask) != 0
}

// SetCafsdet Clear anticipated frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the AFSDET flag in the SAI_xSR register. It is not used in AC97or SPDIF mode. Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) SetCafsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aclrfrFieldCafsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aclrfrFieldCafsdetMask)
	}
}

const (
	RegisterSai_aclrfrFieldClfsdetShift = 6
	RegisterSai_aclrfrFieldClfsdetMask  = 0x40
)

// GetClfsdet Clear late frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the LFSDET flag in the SAI_xSR register. This bit is not used in AC97or SPDIF mode Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) GetClfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_aclrfrFieldClfsdetMask) != 0
}

// SetClfsdet Clear late frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the LFSDET flag in the SAI_xSR register. This bit is not used in AC97or SPDIF mode Reading this bit always returns the value 0.
func (r *registerSai_aclrfrType) SetClfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_aclrfrFieldClfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_aclrfrFieldClfsdetMask)
	}
}

// registerSai_adrType Data register
type registerSai_adrType uint32

const (
	RegisterSai_adrFieldDataShift = 0
	RegisterSai_adrFieldDataMask  = 0xffffffff
)

// GetData Data A write to this register loads the FIFO provided the FIFO is not full. A read from this register empties the FIFO if the FIFO is not empty.
func (r *registerSai_adrType) GetData() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterSai_adrFieldDataMask) >> RegisterSai_adrFieldDataShift)
}

// SetData Data A write to this register loads the FIFO provided the FIFO is not full. A read from this register empties the FIFO if the FIFO is not empty.
func (r *registerSai_adrType) SetData(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_adrFieldDataMask)|(uint32(value)<<RegisterSai_adrFieldDataShift))
}

// registerSai_bcr1Type Configuration register 1
type registerSai_bcr1Type uint32

const (
	RegisterSai_bcr1FieldModeShift = 0
	RegisterSai_bcr1FieldModeMask  = 0x3
)

// GetMode SAIx audio block mode immediately
func (r *registerSai_bcr1Type) GetMode() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldModeMask) >> RegisterSai_bcr1FieldModeShift)
}

// SetMode SAIx audio block mode immediately
func (r *registerSai_bcr1Type) SetMode(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldModeMask)|(uint32(value)<<RegisterSai_bcr1FieldModeShift))
}

const (
	RegisterSai_bcr1FieldPrtcfgShift = 2
	RegisterSai_bcr1FieldPrtcfgMask  = 0xc
)

// GetPrtcfg Protocol configuration. These bits are set and cleared by software. These bits have to be configured when the audio block is disabled.
func (r *registerSai_bcr1Type) GetPrtcfg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldPrtcfgMask) >> RegisterSai_bcr1FieldPrtcfgShift)
}

// SetPrtcfg Protocol configuration. These bits are set and cleared by software. These bits have to be configured when the audio block is disabled.
func (r *registerSai_bcr1Type) SetPrtcfg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldPrtcfgMask)|(uint32(value)<<RegisterSai_bcr1FieldPrtcfgShift))
}

const (
	RegisterSai_bcr1FieldDsShift = 5
	RegisterSai_bcr1FieldDsMask  = 0xe0
)

// GetDs Data size. These bits are set and cleared by software. These bits are ignored when the SPDIF protocols are selected (bit PRTCFG[1:0]), because the frame and the data size are fixed in such case. When the companding mode is selected through COMP[1:0] bits, DS[1:0] are ignored since the data size is fixed to 8 bits by the algorithm. These bits must be configured when the audio block is disabled.
func (r *registerSai_bcr1Type) GetDs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldDsMask) >> RegisterSai_bcr1FieldDsShift)
}

// SetDs Data size. These bits are set and cleared by software. These bits are ignored when the SPDIF protocols are selected (bit PRTCFG[1:0]), because the frame and the data size are fixed in such case. When the companding mode is selected through COMP[1:0] bits, DS[1:0] are ignored since the data size is fixed to 8 bits by the algorithm. These bits must be configured when the audio block is disabled.
func (r *registerSai_bcr1Type) SetDs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldDsMask)|(uint32(value)<<RegisterSai_bcr1FieldDsShift))
}

const (
	RegisterSai_bcr1FieldLsbfirstShift = 8
	RegisterSai_bcr1FieldLsbfirstMask  = 0x100
)

// GetLsbfirst Least significant bit first. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in AC97 audio protocol since AC97 data are always transferred with the MSB first. This bit has no meaning in SPDIF audio protocol since in SPDIF data are always transferred with LSB first.
func (r *registerSai_bcr1Type) GetLsbfirst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldLsbfirstMask) != 0
}

// SetLsbfirst Least significant bit first. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in AC97 audio protocol since AC97 data are always transferred with the MSB first. This bit has no meaning in SPDIF audio protocol since in SPDIF data are always transferred with LSB first.
func (r *registerSai_bcr1Type) SetLsbfirst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr1FieldLsbfirstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldLsbfirstMask)
	}
}

const (
	RegisterSai_bcr1FieldCkstrShift = 9
	RegisterSai_bcr1FieldCkstrMask  = 0x200
)

// GetCkstr Clock strobing edge. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in SPDIF audio protocol.
func (r *registerSai_bcr1Type) GetCkstr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldCkstrMask) != 0
}

// SetCkstr Clock strobing edge. This bit is set and cleared by software. It must be configured when the audio block is disabled. This bit has no meaning in SPDIF audio protocol.
func (r *registerSai_bcr1Type) SetCkstr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr1FieldCkstrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldCkstrMask)
	}
}

const (
	RegisterSai_bcr1FieldSyncenShift = 10
	RegisterSai_bcr1FieldSyncenMask  = 0xc00
)

// GetSyncen Synchronization enable. These bits are set and cleared by software. They must be configured when the audio sub-block is disabled. Note: The audio sub-block should be configured as asynchronous when SPDIF mode is enabled.
func (r *registerSai_bcr1Type) GetSyncen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldSyncenMask) >> RegisterSai_bcr1FieldSyncenShift)
}

// SetSyncen Synchronization enable. These bits are set and cleared by software. They must be configured when the audio sub-block is disabled. Note: The audio sub-block should be configured as asynchronous when SPDIF mode is enabled.
func (r *registerSai_bcr1Type) SetSyncen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldSyncenMask)|(uint32(value)<<RegisterSai_bcr1FieldSyncenShift))
}

const (
	RegisterSai_bcr1FieldMonoShift = 12
	RegisterSai_bcr1FieldMonoMask  = 0x1000
)

// GetMono Mono mode. This bit is set and cleared by software. It is meaningful only when the number of slots is equal to 2. When the mono mode is selected, slot 0 data are duplicated on slot 1 when the audio block operates as a transmitter. In reception mode, the slot1 is discarded and only the data received from slot 0 are stored. Refer to Section: Mono/stereo mode for more details.
func (r *registerSai_bcr1Type) GetMono() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldMonoMask) != 0
}

// SetMono Mono mode. This bit is set and cleared by software. It is meaningful only when the number of slots is equal to 2. When the mono mode is selected, slot 0 data are duplicated on slot 1 when the audio block operates as a transmitter. In reception mode, the slot1 is discarded and only the data received from slot 0 are stored. Refer to Section: Mono/stereo mode for more details.
func (r *registerSai_bcr1Type) SetMono(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr1FieldMonoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldMonoMask)
	}
}

const (
	RegisterSai_bcr1FieldOutdrivShift = 13
	RegisterSai_bcr1FieldOutdrivMask  = 0x2000
)

// GetOutdriv Output drive. This bit is set and cleared by software. Note: This bit has to be set before enabling the audio block and after the audio block configuration.
func (r *registerSai_bcr1Type) GetOutdriv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldOutdrivMask) != 0
}

// SetOutdriv Output drive. This bit is set and cleared by software. Note: This bit has to be set before enabling the audio block and after the audio block configuration.
func (r *registerSai_bcr1Type) SetOutdriv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr1FieldOutdrivMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldOutdrivMask)
	}
}

const (
	RegisterSai_bcr1FieldSaixenShift = 16
	RegisterSai_bcr1FieldSaixenMask  = 0x10000
)

// GetSaixen Audio block enable where x is A or B. This bit is set by software. To switch off the audio block, the application software must program this bit to 0 and poll the bit till it reads back 0, meaning that the block is completely disabled. Before setting this bit to 1, check that it is set to 0, otherwise the enable command will not be taken into account. This bit allows to control the state of SAIx audio block. If it is disabled when an audio frame transfer is ongoing, the ongoing transfer completes and the cell is fully disabled at the end of this audio frame transfer. Note: When SAIx block is configured in master mode, the clock must be present on the input of SAIx before setting SAIXEN bit.
func (r *registerSai_bcr1Type) GetSaixen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldSaixenMask) != 0
}

// SetSaixen Audio block enable where x is A or B. This bit is set by software. To switch off the audio block, the application software must program this bit to 0 and poll the bit till it reads back 0, meaning that the block is completely disabled. Before setting this bit to 1, check that it is set to 0, otherwise the enable command will not be taken into account. This bit allows to control the state of SAIx audio block. If it is disabled when an audio frame transfer is ongoing, the ongoing transfer completes and the cell is fully disabled at the end of this audio frame transfer. Note: When SAIx block is configured in master mode, the clock must be present on the input of SAIx before setting SAIXEN bit.
func (r *registerSai_bcr1Type) SetSaixen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr1FieldSaixenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldSaixenMask)
	}
}

const (
	RegisterSai_bcr1FieldDmaenShift = 17
	RegisterSai_bcr1FieldDmaenMask  = 0x20000
)

// GetDmaen DMA enable. This bit is set and cleared by software. Note: Since the audio block defaults to operate as a transmitter after reset, the MODE[1:0] bits must be configured before setting DMAEN to avoid a DMA request in receiver mode.
func (r *registerSai_bcr1Type) GetDmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldDmaenMask) != 0
}

// SetDmaen DMA enable. This bit is set and cleared by software. Note: Since the audio block defaults to operate as a transmitter after reset, the MODE[1:0] bits must be configured before setting DMAEN to avoid a DMA request in receiver mode.
func (r *registerSai_bcr1Type) SetDmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr1FieldDmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldDmaenMask)
	}
}

const (
	RegisterSai_bcr1FieldNomckShift = 19
	RegisterSai_bcr1FieldNomckMask  = 0x80000
)

// GetNomck No divider
func (r *registerSai_bcr1Type) GetNomck() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldNomckMask) != 0
}

// SetNomck No divider
func (r *registerSai_bcr1Type) SetNomck(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr1FieldNomckMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldNomckMask)
	}
}

const (
	RegisterSai_bcr1FieldMckdivShift = 20
	RegisterSai_bcr1FieldMckdivMask  = 0xf00000
)

// GetMckdiv Master clock divider. These bits are set and cleared by software. These bits are meaningless when the audio block operates in slave mode. They have to be configured when the audio block is disabled. Others: the master clock frequency is calculated accordingly to the following formula:
func (r *registerSai_bcr1Type) GetMckdiv() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldMckdivMask) >> RegisterSai_bcr1FieldMckdivShift)
}

// SetMckdiv Master clock divider. These bits are set and cleared by software. These bits are meaningless when the audio block operates in slave mode. They have to be configured when the audio block is disabled. Others: the master clock frequency is calculated accordingly to the following formula:
func (r *registerSai_bcr1Type) SetMckdiv(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldMckdivMask)|(uint32(value)<<RegisterSai_bcr1FieldMckdivShift))
}

const (
	RegisterSai_bcr1FieldOsrShift = 26
	RegisterSai_bcr1FieldOsrMask  = 0x4000000
)

// GetOsr Oversampling ratio for master clock
func (r *registerSai_bcr1Type) GetOsr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr1FieldOsrMask) != 0
}

// SetOsr Oversampling ratio for master clock
func (r *registerSai_bcr1Type) SetOsr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr1FieldOsrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr1FieldOsrMask)
	}
}

// registerSai_bcr2Type Configuration register 2
type registerSai_bcr2Type uint32

const (
	RegisterSai_bcr2FieldFthShift = 0
	RegisterSai_bcr2FieldFthMask  = 0x7
)

// GetFth FIFO threshold. This bit is set and cleared by software.
func (r *registerSai_bcr2Type) GetFth() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr2FieldFthMask) >> RegisterSai_bcr2FieldFthShift)
}

// SetFth FIFO threshold. This bit is set and cleared by software.
func (r *registerSai_bcr2Type) SetFth(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr2FieldFthMask)|(uint32(value)<<RegisterSai_bcr2FieldFthShift))
}

const (
	RegisterSai_bcr2FieldFflushShift = 3
	RegisterSai_bcr2FieldFflushMask  = 0x8
)

// SetFflush FIFO flush. This bit is set by software. It is always read as 0. This bit should be configured when the SAI is disabled.
func (r *registerSai_bcr2Type) SetFflush(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr2FieldFflushMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr2FieldFflushMask)
	}
}

const (
	RegisterSai_bcr2FieldTrisShift = 4
	RegisterSai_bcr2FieldTrisMask  = 0x10
)

// GetTris Tristate management on data line. This bit is set and cleared by software. It is meaningful only if the audio block is configured as a transmitter. This bit is not used when the audio block is configured in SPDIF mode. It should be configured when SAI is disabled. Refer to Section: Output data line management on an inactive slot for more details.
func (r *registerSai_bcr2Type) GetTris() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr2FieldTrisMask) != 0
}

// SetTris Tristate management on data line. This bit is set and cleared by software. It is meaningful only if the audio block is configured as a transmitter. This bit is not used when the audio block is configured in SPDIF mode. It should be configured when SAI is disabled. Refer to Section: Output data line management on an inactive slot for more details.
func (r *registerSai_bcr2Type) SetTris(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr2FieldTrisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr2FieldTrisMask)
	}
}

const (
	RegisterSai_bcr2FieldMuteShift = 5
	RegisterSai_bcr2FieldMuteMask  = 0x20
)

// GetMute Mute. This bit is set and cleared by software. It is meaningful only when the audio block operates as a transmitter. The MUTE value is linked to value of MUTEVAL if the number of slots is lower or equal to 2, or equal to 0 if it is greater than 2. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSai_bcr2Type) GetMute() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr2FieldMuteMask) != 0
}

// SetMute Mute. This bit is set and cleared by software. It is meaningful only when the audio block operates as a transmitter. The MUTE value is linked to value of MUTEVAL if the number of slots is lower or equal to 2, or equal to 0 if it is greater than 2. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSai_bcr2Type) SetMute(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr2FieldMuteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr2FieldMuteMask)
	}
}

const (
	RegisterSai_bcr2FieldMutevalShift = 6
	RegisterSai_bcr2FieldMutevalMask  = 0x40
)

// GetMuteval Mute value. This bit is set and cleared by software.It must be written before enabling the audio block: SAIXEN. This bit is meaningful only when the audio block operates as a transmitter, the number of slots is lower or equal to 2 and the MUTE bit is set. If more slots are declared, the bit value sent during the transmission in mute mode is equal to 0, whatever the value of MUTEVAL. if the number of slot is lower or equal to 2 and MUTEVAL = 1, the MUTE value transmitted for each slot is the one sent during the previous frame. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSai_bcr2Type) GetMuteval() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr2FieldMutevalMask) != 0
}

// SetMuteval Mute value. This bit is set and cleared by software.It must be written before enabling the audio block: SAIXEN. This bit is meaningful only when the audio block operates as a transmitter, the number of slots is lower or equal to 2 and the MUTE bit is set. If more slots are declared, the bit value sent during the transmission in mute mode is equal to 0, whatever the value of MUTEVAL. if the number of slot is lower or equal to 2 and MUTEVAL = 1, the MUTE value transmitted for each slot is the one sent during the previous frame. Refer to Section: Mute mode for more details. Note: This bit is meaningless and should not be used for SPDIF audio blocks.
func (r *registerSai_bcr2Type) SetMuteval(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr2FieldMutevalMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr2FieldMutevalMask)
	}
}

const (
	RegisterSai_bcr2FieldMutecntShift = 7
	RegisterSai_bcr2FieldMutecntMask  = 0x1f80
)

// GetMutecnt Mute counter. These bits are set and cleared by software. They are used only in reception mode. The value set in these bits is compared to the number of consecutive mute frames detected in reception. When the number of mute frames is equal to this value, the flag MUTEDET will be set and an interrupt will be generated if bit MUTEDETIE is set. Refer to Section: Mute mode for more details.
func (r *registerSai_bcr2Type) GetMutecnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr2FieldMutecntMask) >> RegisterSai_bcr2FieldMutecntShift)
}

// SetMutecnt Mute counter. These bits are set and cleared by software. They are used only in reception mode. The value set in these bits is compared to the number of consecutive mute frames detected in reception. When the number of mute frames is equal to this value, the flag MUTEDET will be set and an interrupt will be generated if bit MUTEDETIE is set. Refer to Section: Mute mode for more details.
func (r *registerSai_bcr2Type) SetMutecnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr2FieldMutecntMask)|(uint32(value)<<RegisterSai_bcr2FieldMutecntShift))
}

const (
	RegisterSai_bcr2FieldCplShift = 13
	RegisterSai_bcr2FieldCplMask  = 0x2000
)

// GetCpl Complement bit. This bit is set and cleared by software. It defines the type of complement to be used for companding mode Note: This bit has effect only when the companding mode is -Law algorithm or A-Law algorithm.
func (r *registerSai_bcr2Type) GetCpl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr2FieldCplMask) != 0
}

// SetCpl Complement bit. This bit is set and cleared by software. It defines the type of complement to be used for companding mode Note: This bit has effect only when the companding mode is -Law algorithm or A-Law algorithm.
func (r *registerSai_bcr2Type) SetCpl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bcr2FieldCplMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr2FieldCplMask)
	}
}

const (
	RegisterSai_bcr2FieldCompShift = 14
	RegisterSai_bcr2FieldCompMask  = 0xc000
)

// GetComp Companding mode. These bits are set and cleared by software. The -Law and the A-Law log are a part of the CCITT G.711 recommendation, the type of complement that will be used depends on CPL bit. The data expansion or data compression are determined by the state of bit MODE[0]. The data compression is applied if the audio block is configured as a transmitter. The data expansion is automatically applied when the audio block is configured as a receiver. Refer to Section: Companding mode for more details. Note: Companding mode is applicable only when TDM is selected.
func (r *registerSai_bcr2Type) GetComp() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bcr2FieldCompMask) >> RegisterSai_bcr2FieldCompShift)
}

// SetComp Companding mode. These bits are set and cleared by software. The -Law and the A-Law log are a part of the CCITT G.711 recommendation, the type of complement that will be used depends on CPL bit. The data expansion or data compression are determined by the state of bit MODE[0]. The data compression is applied if the audio block is configured as a transmitter. The data expansion is automatically applied when the audio block is configured as a receiver. Refer to Section: Companding mode for more details. Note: Companding mode is applicable only when TDM is selected.
func (r *registerSai_bcr2Type) SetComp(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bcr2FieldCompMask)|(uint32(value)<<RegisterSai_bcr2FieldCompShift))
}

// registerSai_bfrcrType This register has no meaning in AC97 and SPDIF audio protocol
type registerSai_bfrcrType uint32

const (
	RegisterSai_bfrcrFieldFrlShift = 0
	RegisterSai_bfrcrFieldFrlMask  = 0xff
)

// GetFrl Frame length. These bits are set and cleared by software. They define the audio frame length expressed in number of SCK clock cycles: the number of bits in the frame is equal to FRL[7:0] + 1. The minimum number of bits to transfer in an audio frame must be equal to 8, otherwise the audio block will behaves in an unexpected way. This is the case when the data size is 8 bits and only one slot 0 is defined in NBSLOT[4:0] of SAI_xSLOTR register (NBSLOT[3:0] = 0000). In master mode, if the master clock (available on MCLK_x pin) is used, the frame length should be aligned with a number equal to a power of 2, ranging from 8 to 256. When the master clock is not used (NODIV = 1), it is recommended to program the frame length to an value ranging from 8 to 256. These bits are meaningless and are not used in AC97 or SPDIF audio block configuration.
func (r *registerSai_bfrcrType) GetFrl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bfrcrFieldFrlMask) >> RegisterSai_bfrcrFieldFrlShift)
}

// SetFrl Frame length. These bits are set and cleared by software. They define the audio frame length expressed in number of SCK clock cycles: the number of bits in the frame is equal to FRL[7:0] + 1. The minimum number of bits to transfer in an audio frame must be equal to 8, otherwise the audio block will behaves in an unexpected way. This is the case when the data size is 8 bits and only one slot 0 is defined in NBSLOT[4:0] of SAI_xSLOTR register (NBSLOT[3:0] = 0000). In master mode, if the master clock (available on MCLK_x pin) is used, the frame length should be aligned with a number equal to a power of 2, ranging from 8 to 256. When the master clock is not used (NODIV = 1), it is recommended to program the frame length to an value ranging from 8 to 256. These bits are meaningless and are not used in AC97 or SPDIF audio block configuration.
func (r *registerSai_bfrcrType) SetFrl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bfrcrFieldFrlMask)|(uint32(value)<<RegisterSai_bfrcrFieldFrlShift))
}

const (
	RegisterSai_bfrcrFieldFsallShift = 8
	RegisterSai_bfrcrFieldFsallMask  = 0x7f00
)

// GetFsall Frame synchronization active level length. These bits are set and cleared by software. They specify the length in number of bit clock (SCK) + 1 (FSALL[6:0] + 1) of the active level of the FS signal in the audio frame These bits are meaningless and are not used in AC97 or SPDIF audio block configuration. They must be configured when the audio block is disabled.
func (r *registerSai_bfrcrType) GetFsall() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bfrcrFieldFsallMask) >> RegisterSai_bfrcrFieldFsallShift)
}

// SetFsall Frame synchronization active level length. These bits are set and cleared by software. They specify the length in number of bit clock (SCK) + 1 (FSALL[6:0] + 1) of the active level of the FS signal in the audio frame These bits are meaningless and are not used in AC97 or SPDIF audio block configuration. They must be configured when the audio block is disabled.
func (r *registerSai_bfrcrType) SetFsall(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bfrcrFieldFsallMask)|(uint32(value)<<RegisterSai_bfrcrFieldFsallShift))
}

const (
	RegisterSai_bfrcrFieldFsdefShift = 16
	RegisterSai_bfrcrFieldFsdefMask  = 0x10000
)

// GetFsdef Frame synchronization definition. This bit is set and cleared by software. When the bit is set, the number of slots defined in the SAI_xSLOTR register has to be even. It means that half of this number of slots will be dedicated to the left channel and the other slots for the right channel (e.g: this bit has to be set for I2S or MSB/LSB-justified protocols...). This bit is meaningless and is not used in AC97 or SPDIF audio block configuration. It must be configured when the audio block is disabled.
func (r *registerSai_bfrcrType) GetFsdef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bfrcrFieldFsdefMask) != 0
}

const (
	RegisterSai_bfrcrFieldFspolShift = 17
	RegisterSai_bfrcrFieldFspolMask  = 0x20000
)

// GetFspol Frame synchronization polarity. This bit is set and cleared by software. It is used to configure the level of the start of frame on the FS signal. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSai_bfrcrType) GetFspol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bfrcrFieldFspolMask) != 0
}

// SetFspol Frame synchronization polarity. This bit is set and cleared by software. It is used to configure the level of the start of frame on the FS signal. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSai_bfrcrType) SetFspol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bfrcrFieldFspolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bfrcrFieldFspolMask)
	}
}

const (
	RegisterSai_bfrcrFieldFsoffShift = 18
	RegisterSai_bfrcrFieldFsoffMask  = 0x40000
)

// GetFsoff Frame synchronization offset. This bit is set and cleared by software. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSai_bfrcrType) GetFsoff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bfrcrFieldFsoffMask) != 0
}

// SetFsoff Frame synchronization offset. This bit is set and cleared by software. It is meaningless and is not used in AC97 or SPDIF audio block configuration. This bit must be configured when the audio block is disabled.
func (r *registerSai_bfrcrType) SetFsoff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bfrcrFieldFsoffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bfrcrFieldFsoffMask)
	}
}

// registerSai_bslotrType This register has no meaning in AC97 and SPDIF audio protocol
type registerSai_bslotrType uint32

const (
	RegisterSai_bslotrFieldFboffShift = 0
	RegisterSai_bslotrFieldFboffMask  = 0x1f
)

// GetFboff First bit offset These bits are set and cleared by software. The value set in this bitfield defines the position of the first data transfer bit in the slot. It represents an offset value. In transmission mode, the bits outside the data field are forced to 0. In reception mode, the extra received bits are discarded. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_bslotrType) GetFboff() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bslotrFieldFboffMask) >> RegisterSai_bslotrFieldFboffShift)
}

// SetFboff First bit offset These bits are set and cleared by software. The value set in this bitfield defines the position of the first data transfer bit in the slot. It represents an offset value. In transmission mode, the bits outside the data field are forced to 0. In reception mode, the extra received bits are discarded. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_bslotrType) SetFboff(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bslotrFieldFboffMask)|(uint32(value)<<RegisterSai_bslotrFieldFboffShift))
}

const (
	RegisterSai_bslotrFieldSlotszShift = 6
	RegisterSai_bslotrFieldSlotszMask  = 0xc0
)

// GetSlotsz Slot size This bits is set and cleared by software. The slot size must be higher or equal to the data size. If this condition is not respected, the behavior of the SAI will be undetermined. Refer to Section: Output data line management on an inactive slot for information on how to drive SD line. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_bslotrType) GetSlotsz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bslotrFieldSlotszMask) >> RegisterSai_bslotrFieldSlotszShift)
}

// SetSlotsz Slot size This bits is set and cleared by software. The slot size must be higher or equal to the data size. If this condition is not respected, the behavior of the SAI will be undetermined. Refer to Section: Output data line management on an inactive slot for information on how to drive SD line. These bits must be set when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_bslotrType) SetSlotsz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bslotrFieldSlotszMask)|(uint32(value)<<RegisterSai_bslotrFieldSlotszShift))
}

const (
	RegisterSai_bslotrFieldNbslotShift = 8
	RegisterSai_bslotrFieldNbslotMask  = 0xf00
)

// GetNbslot Number of slots in an audio frame. These bits are set and cleared by software. The value set in this bitfield represents the number of slots + 1 in the audio frame (including the number of inactive slots). The maximum number of slots is 16. The number of slots should be even if FSDEF bit in the SAI_xFRCR register is set. The number of slots must be configured when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_bslotrType) GetNbslot() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bslotrFieldNbslotMask) >> RegisterSai_bslotrFieldNbslotShift)
}

// SetNbslot Number of slots in an audio frame. These bits are set and cleared by software. The value set in this bitfield represents the number of slots + 1 in the audio frame (including the number of inactive slots). The maximum number of slots is 16. The number of slots should be even if FSDEF bit in the SAI_xFRCR register is set. The number of slots must be configured when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_bslotrType) SetNbslot(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bslotrFieldNbslotMask)|(uint32(value)<<RegisterSai_bslotrFieldNbslotShift))
}

const (
	RegisterSai_bslotrFieldSlotenShift = 16
	RegisterSai_bslotrFieldSlotenMask  = 0xffff0000
)

// GetSloten Slot enable. These bits are set and cleared by software. Each SLOTEN bit corresponds to a slot position from 0 to 15 (maximum 16 slots). The slot must be enabled when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_bslotrType) GetSloten() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bslotrFieldSlotenMask) >> RegisterSai_bslotrFieldSlotenShift)
}

// SetSloten Slot enable. These bits are set and cleared by software. Each SLOTEN bit corresponds to a slot position from 0 to 15 (maximum 16 slots). The slot must be enabled when the audio block is disabled. They are ignored in AC97 or SPDIF mode.
func (r *registerSai_bslotrType) SetSloten(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bslotrFieldSlotenMask)|(uint32(value)<<RegisterSai_bslotrFieldSlotenShift))
}

// registerSai_bimType Interrupt mask register 2
type registerSai_bimType uint32

const (
	RegisterSai_bimFieldOvrudrieShift = 0
	RegisterSai_bimFieldOvrudrieMask  = 0x1
)

// GetOvrudrie Overrun/underrun interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the OVRUDR bit in the SAI_xSR register is set.
func (r *registerSai_bimType) GetOvrudrie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bimFieldOvrudrieMask) != 0
}

// SetOvrudrie Overrun/underrun interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the OVRUDR bit in the SAI_xSR register is set.
func (r *registerSai_bimType) SetOvrudrie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bimFieldOvrudrieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bimFieldOvrudrieMask)
	}
}

const (
	RegisterSai_bimFieldMutedetieShift = 1
	RegisterSai_bimFieldMutedetieMask  = 0x2
)

// GetMutedetie Mute detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the MUTEDET bit in the SAI_xSR register is set. This bit has a meaning only if the audio block is configured in receiver mode.
func (r *registerSai_bimType) GetMutedetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bimFieldMutedetieMask) != 0
}

// SetMutedetie Mute detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the MUTEDET bit in the SAI_xSR register is set. This bit has a meaning only if the audio block is configured in receiver mode.
func (r *registerSai_bimType) SetMutedetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bimFieldMutedetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bimFieldMutedetieMask)
	}
}

const (
	RegisterSai_bimFieldWckcfgieShift = 2
	RegisterSai_bimFieldWckcfgieMask  = 0x4
)

// GetWckcfgie Wrong clock configuration interrupt enable. This bit is set and cleared by software. This bit is taken into account only if the audio block is configured as a master (MODE[1] = 0) and NODIV = 0. It generates an interrupt if the WCKCFG flag in the SAI_xSR register is set. Note: This bit is used only in TDM mode and is meaningless in other modes.
func (r *registerSai_bimType) GetWckcfgie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bimFieldWckcfgieMask) != 0
}

// SetWckcfgie Wrong clock configuration interrupt enable. This bit is set and cleared by software. This bit is taken into account only if the audio block is configured as a master (MODE[1] = 0) and NODIV = 0. It generates an interrupt if the WCKCFG flag in the SAI_xSR register is set. Note: This bit is used only in TDM mode and is meaningless in other modes.
func (r *registerSai_bimType) SetWckcfgie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bimFieldWckcfgieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bimFieldWckcfgieMask)
	}
}

const (
	RegisterSai_bimFieldFreqieShift = 3
	RegisterSai_bimFieldFreqieMask  = 0x8
)

// GetFreqie FIFO request interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the FREQ bit in the SAI_xSR register is set. Since the audio block defaults to operate as a transmitter after reset, the MODE bit must be configured before setting FREQIE to avoid a parasitic interruption in receiver mode,
func (r *registerSai_bimType) GetFreqie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bimFieldFreqieMask) != 0
}

// SetFreqie FIFO request interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt is generated if the FREQ bit in the SAI_xSR register is set. Since the audio block defaults to operate as a transmitter after reset, the MODE bit must be configured before setting FREQIE to avoid a parasitic interruption in receiver mode,
func (r *registerSai_bimType) SetFreqie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bimFieldFreqieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bimFieldFreqieMask)
	}
}

const (
	RegisterSai_bimFieldCnrdyieShift = 4
	RegisterSai_bimFieldCnrdyieMask  = 0x10
)

// GetCnrdyie Codec not ready interrupt enable (AC97). This bit is set and cleared by software. When the interrupt is enabled, the audio block detects in the slot 0 (tag0) of the AC97 frame if the Codec connected to this line is ready or not. If it is not ready, the CNRDY flag in the SAI_xSR register is set and an interruption i generated. This bit has a meaning only if the AC97 mode is selected through PRTCFG[1:0] bits and the audio block is operates as a receiver.
func (r *registerSai_bimType) GetCnrdyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bimFieldCnrdyieMask) != 0
}

// SetCnrdyie Codec not ready interrupt enable (AC97). This bit is set and cleared by software. When the interrupt is enabled, the audio block detects in the slot 0 (tag0) of the AC97 frame if the Codec connected to this line is ready or not. If it is not ready, the CNRDY flag in the SAI_xSR register is set and an interruption i generated. This bit has a meaning only if the AC97 mode is selected through PRTCFG[1:0] bits and the audio block is operates as a receiver.
func (r *registerSai_bimType) SetCnrdyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bimFieldCnrdyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bimFieldCnrdyieMask)
	}
}

const (
	RegisterSai_bimFieldAfsdetieShift = 5
	RegisterSai_bimFieldAfsdetieMask  = 0x20
)

// GetAfsdetie Anticipated frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the AFSDET bit in the SAI_xSR register is set. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSai_bimType) GetAfsdetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bimFieldAfsdetieMask) != 0
}

// SetAfsdetie Anticipated frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the AFSDET bit in the SAI_xSR register is set. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSai_bimType) SetAfsdetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bimFieldAfsdetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bimFieldAfsdetieMask)
	}
}

const (
	RegisterSai_bimFieldLfsdetieShift = 6
	RegisterSai_bimFieldLfsdetieMask  = 0x40
)

// GetLfsdetie Late frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the LFSDET bit is set in the SAI_xSR register. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSai_bimType) GetLfsdetie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bimFieldLfsdetieMask) != 0
}

// SetLfsdetie Late frame synchronization detection interrupt enable. This bit is set and cleared by software. When this bit is set, an interrupt will be generated if the LFSDET bit is set in the SAI_xSR register. This bit is meaningless in AC97, SPDIF mode or when the audio block operates as a master.
func (r *registerSai_bimType) SetLfsdetie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bimFieldLfsdetieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bimFieldLfsdetieMask)
	}
}

// registerSai_bsrType Status register
type registerSai_bsrType uint32

const (
	RegisterSai_bsrFieldOvrudrShift = 0
	RegisterSai_bsrFieldOvrudrMask  = 0x1
)

// GetOvrudr Overrun / underrun. This bit is read only. The overrun and underrun conditions can occur only when the audio block is configured as a receiver and a transmitter, respectively. It can generate an interrupt if OVRUDRIE bit is set in SAI_xIM register. This flag is cleared when the software sets COVRUDR bit in SAI_xCLRFR register.
func (r *registerSai_bsrType) GetOvrudr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bsrFieldOvrudrMask) != 0
}

// SetOvrudr Overrun / underrun. This bit is read only. The overrun and underrun conditions can occur only when the audio block is configured as a receiver and a transmitter, respectively. It can generate an interrupt if OVRUDRIE bit is set in SAI_xIM register. This flag is cleared when the software sets COVRUDR bit in SAI_xCLRFR register.
func (r *registerSai_bsrType) SetOvrudr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bsrFieldOvrudrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bsrFieldOvrudrMask)
	}
}

const (
	RegisterSai_bsrFieldMutedetShift = 1
	RegisterSai_bsrFieldMutedetMask  = 0x2
)

// GetMutedet Mute detection. This bit is read only. This flag is set if consecutive 0 values are received in each slot of a given audio frame and for a consecutive number of audio frames (set in the MUTECNT bit in the SAI_xCR2 register). It can generate an interrupt if MUTEDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets bit CMUTEDET in the SAI_xCLRFR register.
func (r *registerSai_bsrType) GetMutedet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bsrFieldMutedetMask) != 0
}

// SetMutedet Mute detection. This bit is read only. This flag is set if consecutive 0 values are received in each slot of a given audio frame and for a consecutive number of audio frames (set in the MUTECNT bit in the SAI_xCR2 register). It can generate an interrupt if MUTEDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets bit CMUTEDET in the SAI_xCLRFR register.
func (r *registerSai_bsrType) SetMutedet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bsrFieldMutedetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bsrFieldMutedetMask)
	}
}

const (
	RegisterSai_bsrFieldWckcfgShift = 2
	RegisterSai_bsrFieldWckcfgMask  = 0x4
)

// GetWckcfg Wrong clock configuration flag. This bit is read only. This bit is used only when the audio block operates in master mode (MODE[1] = 0) and NODIV = 0. It can generate an interrupt if WCKCFGIE bit is set in SAI_xIM register. This flag is cleared when the software sets CWCKCFG bit in SAI_xCLRFR register.
func (r *registerSai_bsrType) GetWckcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bsrFieldWckcfgMask) != 0
}

// SetWckcfg Wrong clock configuration flag. This bit is read only. This bit is used only when the audio block operates in master mode (MODE[1] = 0) and NODIV = 0. It can generate an interrupt if WCKCFGIE bit is set in SAI_xIM register. This flag is cleared when the software sets CWCKCFG bit in SAI_xCLRFR register.
func (r *registerSai_bsrType) SetWckcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bsrFieldWckcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bsrFieldWckcfgMask)
	}
}

const (
	RegisterSai_bsrFieldFreqShift = 3
	RegisterSai_bsrFieldFreqMask  = 0x8
)

// GetFreq FIFO request. This bit is read only. The request depends on the audio block configuration: If the block is configured in transmission mode, the FIFO request is related to a write request operation in the SAI_xDR. If the block configured in reception, the FIFO request related to a read request operation from the SAI_xDR. This flag can generate an interrupt if FREQIE bit is set in SAI_xIM register.
func (r *registerSai_bsrType) GetFreq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bsrFieldFreqMask) != 0
}

// SetFreq FIFO request. This bit is read only. The request depends on the audio block configuration: If the block is configured in transmission mode, the FIFO request is related to a write request operation in the SAI_xDR. If the block configured in reception, the FIFO request related to a read request operation from the SAI_xDR. This flag can generate an interrupt if FREQIE bit is set in SAI_xIM register.
func (r *registerSai_bsrType) SetFreq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bsrFieldFreqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bsrFieldFreqMask)
	}
}

const (
	RegisterSai_bsrFieldCnrdyShift = 4
	RegisterSai_bsrFieldCnrdyMask  = 0x10
)

// GetCnrdy Codec not ready. This bit is read only. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register and configured in receiver mode. It can generate an interrupt if CNRDYIE bit is set in SAI_xIM register. This flag is cleared when the software sets CCNRDY bit in SAI_xCLRFR register.
func (r *registerSai_bsrType) GetCnrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bsrFieldCnrdyMask) != 0
}

// SetCnrdy Codec not ready. This bit is read only. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register and configured in receiver mode. It can generate an interrupt if CNRDYIE bit is set in SAI_xIM register. This flag is cleared when the software sets CCNRDY bit in SAI_xCLRFR register.
func (r *registerSai_bsrType) SetCnrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bsrFieldCnrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bsrFieldCnrdyMask)
	}
}

const (
	RegisterSai_bsrFieldAfsdetShift = 5
	RegisterSai_bsrFieldAfsdetMask  = 0x20
)

// GetAfsdet Anticipated frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97or SPDIF mode. It can generate an interrupt if AFSDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets CAFSDET bit in SAI_xCLRFR register.
func (r *registerSai_bsrType) GetAfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bsrFieldAfsdetMask) != 0
}

// SetAfsdet Anticipated frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97or SPDIF mode. It can generate an interrupt if AFSDETIE bit is set in SAI_xIM register. This flag is cleared when the software sets CAFSDET bit in SAI_xCLRFR register.
func (r *registerSai_bsrType) SetAfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bsrFieldAfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bsrFieldAfsdetMask)
	}
}

const (
	RegisterSai_bsrFieldLfsdetShift = 6
	RegisterSai_bsrFieldLfsdetMask  = 0x40
)

// GetLfsdet Late frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97 or SPDIF mode. It can generate an interrupt if LFSDETIE bit is set in the SAI_xIM register. This flag is cleared when the software sets bit CLFSDET in SAI_xCLRFR register
func (r *registerSai_bsrType) GetLfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bsrFieldLfsdetMask) != 0
}

// SetLfsdet Late frame synchronization detection. This bit is read only. This flag can be set only if the audio block is configured in slave mode. It is not used in AC97 or SPDIF mode. It can generate an interrupt if LFSDETIE bit is set in the SAI_xIM register. This flag is cleared when the software sets bit CLFSDET in SAI_xCLRFR register
func (r *registerSai_bsrType) SetLfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bsrFieldLfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bsrFieldLfsdetMask)
	}
}

const (
	RegisterSai_bsrFieldFlvlShift = 16
	RegisterSai_bsrFieldFlvlMask  = 0x70000
)

// GetFlvl FIFO level threshold. This bit is read only. The FIFO level threshold flag is managed only by hardware and its setting depends on SAI block configuration (transmitter or receiver mode). If the SAI block is configured as transmitter: If SAI block is configured as receiver:
func (r *registerSai_bsrType) GetFlvl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bsrFieldFlvlMask) >> RegisterSai_bsrFieldFlvlShift)
}

// SetFlvl FIFO level threshold. This bit is read only. The FIFO level threshold flag is managed only by hardware and its setting depends on SAI block configuration (transmitter or receiver mode). If the SAI block is configured as transmitter: If SAI block is configured as receiver:
func (r *registerSai_bsrType) SetFlvl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bsrFieldFlvlMask)|(uint32(value)<<RegisterSai_bsrFieldFlvlShift))
}

// registerSai_bclrfrType Clear flag register
type registerSai_bclrfrType uint32

const (
	RegisterSai_bclrfrFieldCovrudrShift = 0
	RegisterSai_bclrfrFieldCovrudrMask  = 0x1
)

// GetCovrudr Clear overrun / underrun. This bit is write only. Programming this bit to 1 clears the OVRUDR flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) GetCovrudr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bclrfrFieldCovrudrMask) != 0
}

// SetCovrudr Clear overrun / underrun. This bit is write only. Programming this bit to 1 clears the OVRUDR flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) SetCovrudr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bclrfrFieldCovrudrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bclrfrFieldCovrudrMask)
	}
}

const (
	RegisterSai_bclrfrFieldCmutedetShift = 1
	RegisterSai_bclrfrFieldCmutedetMask  = 0x2
)

// GetCmutedet Mute detection flag. This bit is write only. Programming this bit to 1 clears the MUTEDET flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) GetCmutedet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bclrfrFieldCmutedetMask) != 0
}

// SetCmutedet Mute detection flag. This bit is write only. Programming this bit to 1 clears the MUTEDET flag in the SAI_xSR register. Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) SetCmutedet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bclrfrFieldCmutedetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bclrfrFieldCmutedetMask)
	}
}

const (
	RegisterSai_bclrfrFieldCwckcfgShift = 2
	RegisterSai_bclrfrFieldCwckcfgMask  = 0x4
)

// GetCwckcfg Clear wrong clock configuration flag. This bit is write only. Programming this bit to 1 clears the WCKCFG flag in the SAI_xSR register. This bit is used only when the audio block is set as master (MODE[1] = 0) and NODIV = 0 in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) GetCwckcfg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bclrfrFieldCwckcfgMask) != 0
}

// SetCwckcfg Clear wrong clock configuration flag. This bit is write only. Programming this bit to 1 clears the WCKCFG flag in the SAI_xSR register. This bit is used only when the audio block is set as master (MODE[1] = 0) and NODIV = 0 in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) SetCwckcfg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bclrfrFieldCwckcfgMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bclrfrFieldCwckcfgMask)
	}
}

const (
	RegisterSai_bclrfrFieldCcnrdyShift = 4
	RegisterSai_bclrfrFieldCcnrdyMask  = 0x10
)

// GetCcnrdy Clear Codec not ready flag. This bit is write only. Programming this bit to 1 clears the CNRDY flag in the SAI_xSR register. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) GetCcnrdy() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bclrfrFieldCcnrdyMask) != 0
}

// SetCcnrdy Clear Codec not ready flag. This bit is write only. Programming this bit to 1 clears the CNRDY flag in the SAI_xSR register. This bit is used only when the AC97 audio protocol is selected in the SAI_xCR1 register. Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) SetCcnrdy(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bclrfrFieldCcnrdyMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bclrfrFieldCcnrdyMask)
	}
}

const (
	RegisterSai_bclrfrFieldCafsdetShift = 5
	RegisterSai_bclrfrFieldCafsdetMask  = 0x20
)

// GetCafsdet Clear anticipated frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the AFSDET flag in the SAI_xSR register. It is not used in AC97or SPDIF mode. Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) GetCafsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bclrfrFieldCafsdetMask) != 0
}

// SetCafsdet Clear anticipated frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the AFSDET flag in the SAI_xSR register. It is not used in AC97or SPDIF mode. Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) SetCafsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bclrfrFieldCafsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bclrfrFieldCafsdetMask)
	}
}

const (
	RegisterSai_bclrfrFieldClfsdetShift = 6
	RegisterSai_bclrfrFieldClfsdetMask  = 0x40
)

// GetClfsdet Clear late frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the LFSDET flag in the SAI_xSR register. This bit is not used in AC97or SPDIF mode Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) GetClfsdet() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_bclrfrFieldClfsdetMask) != 0
}

// SetClfsdet Clear late frame synchronization detection flag. This bit is write only. Programming this bit to 1 clears the LFSDET flag in the SAI_xSR register. This bit is not used in AC97or SPDIF mode Reading this bit always returns the value 0.
func (r *registerSai_bclrfrType) SetClfsdet(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_bclrfrFieldClfsdetMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_bclrfrFieldClfsdetMask)
	}
}

// registerSai_bdrType Data register
type registerSai_bdrType uint32

const (
	RegisterSai_bdrFieldDataShift = 0
	RegisterSai_bdrFieldDataMask  = 0xffffffff
)

// GetData Data A write to this register loads the FIFO provided the FIFO is not full. A read from this register empties the FIFO if the FIFO is not empty.
func (r *registerSai_bdrType) GetData() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterSai_bdrFieldDataMask) >> RegisterSai_bdrFieldDataShift)
}

// SetData Data A write to this register loads the FIFO provided the FIFO is not full. A read from this register empties the FIFO if the FIFO is not empty.
func (r *registerSai_bdrType) SetData(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_bdrFieldDataMask)|(uint32(value)<<RegisterSai_bdrFieldDataShift))
}

// registerSai_pdmcrType PDM control register
type registerSai_pdmcrType uint32

const (
	RegisterSai_pdmcrFieldPdmenShift = 0
	RegisterSai_pdmcrFieldPdmenMask  = 0x1
)

// GetPdmen PDM enable
func (r *registerSai_pdmcrType) GetPdmen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmcrFieldPdmenMask) != 0
}

// SetPdmen PDM enable
func (r *registerSai_pdmcrType) SetPdmen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_pdmcrFieldPdmenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmcrFieldPdmenMask)
	}
}

const (
	RegisterSai_pdmcrFieldMicnbrShift = 4
	RegisterSai_pdmcrFieldMicnbrMask  = 0x30
)

// GetMicnbr Number of microphones
func (r *registerSai_pdmcrType) GetMicnbr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmcrFieldMicnbrMask) >> RegisterSai_pdmcrFieldMicnbrShift)
}

// SetMicnbr Number of microphones
func (r *registerSai_pdmcrType) SetMicnbr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmcrFieldMicnbrMask)|(uint32(value)<<RegisterSai_pdmcrFieldMicnbrShift))
}

const (
	RegisterSai_pdmcrFieldCken1Shift = 8
	RegisterSai_pdmcrFieldCken1Mask  = 0x100
)

// GetCken1 Clock enable of bitstream clock number 1
func (r *registerSai_pdmcrType) GetCken1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmcrFieldCken1Mask) != 0
}

// SetCken1 Clock enable of bitstream clock number 1
func (r *registerSai_pdmcrType) SetCken1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_pdmcrFieldCken1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmcrFieldCken1Mask)
	}
}

const (
	RegisterSai_pdmcrFieldCken2Shift = 9
	RegisterSai_pdmcrFieldCken2Mask  = 0x200
)

// GetCken2 Clock enable of bitstream clock number 2
func (r *registerSai_pdmcrType) GetCken2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmcrFieldCken2Mask) != 0
}

// SetCken2 Clock enable of bitstream clock number 2
func (r *registerSai_pdmcrType) SetCken2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_pdmcrFieldCken2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmcrFieldCken2Mask)
	}
}

const (
	RegisterSai_pdmcrFieldCken3Shift = 10
	RegisterSai_pdmcrFieldCken3Mask  = 0x400
)

// GetCken3 Clock enable of bitstream clock number 3
func (r *registerSai_pdmcrType) GetCken3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmcrFieldCken3Mask) != 0
}

// SetCken3 Clock enable of bitstream clock number 3
func (r *registerSai_pdmcrType) SetCken3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_pdmcrFieldCken3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmcrFieldCken3Mask)
	}
}

const (
	RegisterSai_pdmcrFieldCken4Shift = 11
	RegisterSai_pdmcrFieldCken4Mask  = 0x800
)

// GetCken4 Clock enable of bitstream clock number 4
func (r *registerSai_pdmcrType) GetCken4() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmcrFieldCken4Mask) != 0
}

// SetCken4 Clock enable of bitstream clock number 4
func (r *registerSai_pdmcrType) SetCken4(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSai_pdmcrFieldCken4Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmcrFieldCken4Mask)
	}
}

// registerSai_pdmdlyType PDM delay register
type registerSai_pdmdlyType uint32

const (
	RegisterSai_pdmdlyFieldDlym1lShift = 0
	RegisterSai_pdmdlyFieldDlym1lMask  = 0x7
)

// GetDlym1l Delay line adjust for first microphone of pair 1
func (r *registerSai_pdmdlyType) GetDlym1l() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmdlyFieldDlym1lMask) >> RegisterSai_pdmdlyFieldDlym1lShift)
}

// SetDlym1l Delay line adjust for first microphone of pair 1
func (r *registerSai_pdmdlyType) SetDlym1l(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmdlyFieldDlym1lMask)|(uint32(value)<<RegisterSai_pdmdlyFieldDlym1lShift))
}

const (
	RegisterSai_pdmdlyFieldDlym1rShift = 4
	RegisterSai_pdmdlyFieldDlym1rMask  = 0x70
)

// GetDlym1r Delay line adjust for second microphone of pair 1
func (r *registerSai_pdmdlyType) GetDlym1r() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmdlyFieldDlym1rMask) >> RegisterSai_pdmdlyFieldDlym1rShift)
}

// SetDlym1r Delay line adjust for second microphone of pair 1
func (r *registerSai_pdmdlyType) SetDlym1r(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmdlyFieldDlym1rMask)|(uint32(value)<<RegisterSai_pdmdlyFieldDlym1rShift))
}

const (
	RegisterSai_pdmdlyFieldDlym2lShift = 8
	RegisterSai_pdmdlyFieldDlym2lMask  = 0x700
)

// GetDlym2l Delay line for first microphone of pair 2
func (r *registerSai_pdmdlyType) GetDlym2l() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmdlyFieldDlym2lMask) >> RegisterSai_pdmdlyFieldDlym2lShift)
}

// SetDlym2l Delay line for first microphone of pair 2
func (r *registerSai_pdmdlyType) SetDlym2l(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmdlyFieldDlym2lMask)|(uint32(value)<<RegisterSai_pdmdlyFieldDlym2lShift))
}

const (
	RegisterSai_pdmdlyFieldDlym2rShift = 12
	RegisterSai_pdmdlyFieldDlym2rMask  = 0x7000
)

// GetDlym2r Delay line for second microphone of pair 2
func (r *registerSai_pdmdlyType) GetDlym2r() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmdlyFieldDlym2rMask) >> RegisterSai_pdmdlyFieldDlym2rShift)
}

// SetDlym2r Delay line for second microphone of pair 2
func (r *registerSai_pdmdlyType) SetDlym2r(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmdlyFieldDlym2rMask)|(uint32(value)<<RegisterSai_pdmdlyFieldDlym2rShift))
}

const (
	RegisterSai_pdmdlyFieldDlym3lShift = 16
	RegisterSai_pdmdlyFieldDlym3lMask  = 0x70000
)

// GetDlym3l Delay line for first microphone of pair 3
func (r *registerSai_pdmdlyType) GetDlym3l() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmdlyFieldDlym3lMask) >> RegisterSai_pdmdlyFieldDlym3lShift)
}

// SetDlym3l Delay line for first microphone of pair 3
func (r *registerSai_pdmdlyType) SetDlym3l(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmdlyFieldDlym3lMask)|(uint32(value)<<RegisterSai_pdmdlyFieldDlym3lShift))
}

const (
	RegisterSai_pdmdlyFieldDlym3rShift = 20
	RegisterSai_pdmdlyFieldDlym3rMask  = 0x700000
)

// GetDlym3r Delay line for second microphone of pair 3
func (r *registerSai_pdmdlyType) GetDlym3r() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmdlyFieldDlym3rMask) >> RegisterSai_pdmdlyFieldDlym3rShift)
}

// SetDlym3r Delay line for second microphone of pair 3
func (r *registerSai_pdmdlyType) SetDlym3r(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmdlyFieldDlym3rMask)|(uint32(value)<<RegisterSai_pdmdlyFieldDlym3rShift))
}

const (
	RegisterSai_pdmdlyFieldDlym4lShift = 24
	RegisterSai_pdmdlyFieldDlym4lMask  = 0x7000000
)

// GetDlym4l Delay line for first microphone of pair 4
func (r *registerSai_pdmdlyType) GetDlym4l() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmdlyFieldDlym4lMask) >> RegisterSai_pdmdlyFieldDlym4lShift)
}

// SetDlym4l Delay line for first microphone of pair 4
func (r *registerSai_pdmdlyType) SetDlym4l(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmdlyFieldDlym4lMask)|(uint32(value)<<RegisterSai_pdmdlyFieldDlym4lShift))
}

const (
	RegisterSai_pdmdlyFieldDlym4rShift = 28
	RegisterSai_pdmdlyFieldDlym4rMask  = 0x70000000
)

// GetDlym4r Delay line for second microphone of pair 4
func (r *registerSai_pdmdlyType) GetDlym4r() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSai_pdmdlyFieldDlym4rMask) >> RegisterSai_pdmdlyFieldDlym4rShift)
}

// SetDlym4r Delay line for second microphone of pair 4
func (r *registerSai_pdmdlyType) SetDlym4r(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSai_pdmdlyFieldDlym4rMask)|(uint32(value)<<RegisterSai_pdmdlyFieldDlym4rShift))
}
