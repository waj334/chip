//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package ethernet_mac

import (
	"unsafe"
	"volatile"
)

var (
	Ethernet_mac = (*_ethernet_mac)(unsafe.Pointer(uintptr(0x40028000)))
)

type _ethernet_mac struct {
	Maccr                          registerMaccrType
	Macecr                         registerMacecrType
	Macpfr                         registerMacpfrType
	Macwtr                         registerMacwtrType
	Macht0r                        registerMacht0rType
	Macht1r                        registerMacht1rType
	_                              [56]byte
	Macvtr                         registerMacvtrType
	_                              [4]byte
	Macvhtr                        registerMacvhtrType
	_                              [4]byte
	Macvir                         registerMacvirType
	Macivir                        registerMacivirType
	_                              [8]byte
	Macqtxfcr                      registerMacqtxfcrType
	_                              [28]byte
	Macrxfcr                       registerMacrxfcrType
	_                              [28]byte
	Macisr                         registerMacisrType
	Macier                         registerMacierType
	Macrxtxsr                      registerMacrxtxsrType
	_                              [4]byte
	Macpcsr                        registerMacpcsrType
	Macrwkpfr                      registerMacrwkpfrType
	_                              [8]byte
	Maclcsr                        registerMaclcsrType
	Macltcr                        registerMacltcrType
	Macletr                        registerMacletrType
	Mac1ustcr                      registerMac1ustcrType
	_                              [48]byte
	Macvr                          registerMacvrType
	Macdr                          registerMacdrType
	_                              [8]byte
	Machwf1r                       registerMachwf1rType
	Machwf2r                       registerMachwf2rType
	_                              [216]byte
	Macmdioar                      registerMacmdioarType
	Macmdiodr                      registerMacmdiodrType
	_                              [248]byte
	Maca0hr                        registerMaca0hrType
	Maca0lr                        registerMaca0lrType
	Maca1hr                        registerMaca1hrType
	Maca1lr                        registerMaca1lrType
	Maca2hr                        registerMaca2hrType
	Maca2lr                        registerMaca2lrType
	Maca3hr                        registerMaca3hrType
	Maca3lr                        registerMaca3lrType
	_                              [992]byte
	Mmccontrol                     registerMmccontrolType
	Mmcrxinterrupt                 registerMmcrxinterruptType
	Mmctxinterrupt                 registerMmctxinterruptType
	Mmcrxinterruptmask             registerMmcrxinterruptmaskType
	Mmctxinterruptmask             registerMmctxinterruptmaskType
	_                              [56]byte
	Txsinglecollisiongoodpackets   registerTxsinglecollisiongoodpacketsType
	Txmultiplecollisiongoodpackets registerTxmultiplecollisiongoodpacketsType
	_                              [20]byte
	Txpacketcountgood              registerTxpacketcountgoodType
	_                              [40]byte
	Rxcrcerrorpackets              registerRxcrcerrorpacketsType
	Rxalignmenterrorpackets        registerRxalignmenterrorpacketsType
	_                              [40]byte
	Rxunicastpacketsgood           registerRxunicastpacketsgoodType
	_                              [36]byte
	Txlpiuseccntr                  registerTxlpiuseccntrType
	Txlpitrancntr                  registerTxlpitrancntrType
	Rxlpiuseccntr                  registerRxlpiuseccntrType
	Rxlpitrancntr                  registerRxlpitrancntrType
	_                              [260]byte
	Macl3l4c0r                     registerMacl3l4c0rType
	Macl4a0r                       registerMacl4a0rType
	_                              [8]byte
	Macl3a00r                      registerMacl3a00rType
	Macl3a10r                      registerMacl3a10rType
	Macl3a20                       registerMacl3a20Type
	Macl3a30                       registerMacl3a30Type
	_                              [16]byte
	Macl3l4c1r                     registerMacl3l4c1rType
	Macl4a1r                       registerMacl4a1rType
	_                              [8]byte
	Macl3a01r                      registerMacl3a01rType
	Macl3a11r                      registerMacl3a11rType
	Macl3a21r                      registerMacl3a21rType
	Macl3a31r                      registerMacl3a31rType
	_                              [400]byte
	Macarpar                       registerMacarparType
	_                              [28]byte
	Mactscr                        registerMactscrType
	Macssir                        registerMacssirType
	Macstsr                        registerMacstsrType
	Macstnr                        registerMacstnrType
	Macstsur                       registerMacstsurType
	Macstnur                       registerMacstnurType
	Mactsar                        registerMactsarType
	_                              [4]byte
	Mactssr                        registerMactssrType
	_                              [12]byte
	Mactxtssnr                     registerMactxtssnrType
	Mactxtsssr                     registerMactxtsssrType
	_                              [8]byte
	Macacr                         registerMacacrType
	_                              [4]byte
	Macatsnr                       registerMacatsnrType
	Macatssr                       registerMacatssrType
	Mactsiacr                      registerMactsiacrType
	Mactseacr                      registerMactseacrType
	Mactsicnr                      registerMactsicnrType
	Mactsecnr                      registerMactsecnrType
	_                              [16]byte
	Macppscr                       registerMacppscrType
	_                              [12]byte
	Macppsttsr                     registerMacppsttsrType
	Macppsttnr                     registerMacppsttnrType
	Macppsir                       registerMacppsirType
	Macppswr                       registerMacppswrType
	_                              [48]byte
	Macpocr                        registerMacpocrType
	Macspi0r                       registerMacspi0rType
	Macspi1r                       registerMacspi1rType
	Macspi2r                       registerMacspi2rType
	Maclmir                        registerMaclmirType
	_                              [44]byte
	Mtlomr                         registerMtlomrType
	_                              [28]byte
	Mtlisr                         registerMtlisrType
	_                              [220]byte
	Mtltxqomr                      registerMtltxqomrType
	Mtltxqur                       registerMtltxqurType
	Mtltxqdr                       registerMtltxqdrType
	_                              [32]byte
	Mtlqicsr                       registerMtlqicsrType
	Mtlrxqomr                      registerMtlrxqomrType
	Mtlrxqmpocr                    registerMtlrxqmpocrType
	Mtlrxqdr                       registerMtlrxqdrType
	_                              [708]byte
	Dmamr                          registerDmamrType
	Dmasbmr                        registerDmasbmrType
	Dmaisr                         registerDmaisrType
	Dmadsr                         registerDmadsrType
	_                              [240]byte
	Dmaccr                         registerDmaccrType
	Dmactxcr                       registerDmactxcrType
	Dmacrxcr                       registerDmacrxcrType
	_                              [8]byte
	Dmactxdlar                     registerDmactxdlarType
	_                              [4]byte
	Dmacrxdlar                     registerDmacrxdlarType
	Dmactxdtpr                     registerDmactxdtprType
	_                              [4]byte
	Dmacrxdtpr                     registerDmacrxdtprType
	Dmactxrlr                      registerDmactxrlrType
	Dmacrxrlr                      registerDmacrxrlrType
	Dmacier                        registerDmacierType
	Dmacrxiwtr                     registerDmacrxiwtrType
	_                              [8]byte
	Dmaccatxdr                     registerDmaccatxdrType
	_                              [4]byte
	Dmaccarxdr                     registerDmaccarxdrType
	_                              [4]byte
	Dmaccatxbr                     registerDmaccatxbrType
	_                              [4]byte
	Dmaccarxbr                     registerDmaccarxbrType
	Dmacsr                         registerDmacsrType
	_                              [8]byte
	Dmacmfcr                       registerDmacmfcrType
}

// registerMaccrType Operating mode configuration register
type registerMaccrType uint32

const (
	RegisterMaccrFieldReShift = 0
	RegisterMaccrFieldReMask  = 0x1
)

// GetRe Receiver Enable
func (r *registerMaccrType) GetRe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldReMask) != 0
}

// SetRe Receiver Enable
func (r *registerMaccrType) SetRe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldReMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldReMask)
	}
}

const (
	RegisterMaccrFieldTeShift = 1
	RegisterMaccrFieldTeMask  = 0x2
)

// GetTe TE
func (r *registerMaccrType) GetTe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldTeMask) != 0
}

// SetTe TE
func (r *registerMaccrType) SetTe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldTeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldTeMask)
	}
}

const (
	RegisterMaccrFieldPrelenShift = 2
	RegisterMaccrFieldPrelenMask  = 0xc
)

// GetPrelen PRELEN
func (r *registerMaccrType) GetPrelen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldPrelenMask) >> RegisterMaccrFieldPrelenShift)
}

// SetPrelen PRELEN
func (r *registerMaccrType) SetPrelen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldPrelenMask)|(uint32(value)<<RegisterMaccrFieldPrelenShift))
}

const (
	RegisterMaccrFieldDcShift = 4
	RegisterMaccrFieldDcMask  = 0x10
)

// GetDc DC
func (r *registerMaccrType) GetDc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldDcMask) != 0
}

// SetDc DC
func (r *registerMaccrType) SetDc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldDcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldDcMask)
	}
}

const (
	RegisterMaccrFieldBlShift = 5
	RegisterMaccrFieldBlMask  = 0x60
)

// GetBl BL
func (r *registerMaccrType) GetBl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldBlMask) >> RegisterMaccrFieldBlShift)
}

// SetBl BL
func (r *registerMaccrType) SetBl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldBlMask)|(uint32(value)<<RegisterMaccrFieldBlShift))
}

const (
	RegisterMaccrFieldDrShift = 8
	RegisterMaccrFieldDrMask  = 0x100
)

// GetDr DR
func (r *registerMaccrType) GetDr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldDrMask) != 0
}

// SetDr DR
func (r *registerMaccrType) SetDr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldDrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldDrMask)
	}
}

const (
	RegisterMaccrFieldDcrsShift = 9
	RegisterMaccrFieldDcrsMask  = 0x200
)

// GetDcrs DCRS
func (r *registerMaccrType) GetDcrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldDcrsMask) != 0
}

// SetDcrs DCRS
func (r *registerMaccrType) SetDcrs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldDcrsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldDcrsMask)
	}
}

const (
	RegisterMaccrFieldDoShift = 10
	RegisterMaccrFieldDoMask  = 0x400
)

// GetDo DO
func (r *registerMaccrType) GetDo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldDoMask) != 0
}

// SetDo DO
func (r *registerMaccrType) SetDo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldDoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldDoMask)
	}
}

const (
	RegisterMaccrFieldEcrsfdShift = 11
	RegisterMaccrFieldEcrsfdMask  = 0x800
)

// GetEcrsfd ECRSFD
func (r *registerMaccrType) GetEcrsfd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldEcrsfdMask) != 0
}

// SetEcrsfd ECRSFD
func (r *registerMaccrType) SetEcrsfd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldEcrsfdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldEcrsfdMask)
	}
}

const (
	RegisterMaccrFieldLmShift = 12
	RegisterMaccrFieldLmMask  = 0x1000
)

// GetLm LM
func (r *registerMaccrType) GetLm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldLmMask) != 0
}

// SetLm LM
func (r *registerMaccrType) SetLm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldLmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldLmMask)
	}
}

const (
	RegisterMaccrFieldDmShift = 13
	RegisterMaccrFieldDmMask  = 0x2000
)

// GetDm DM
func (r *registerMaccrType) GetDm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldDmMask) != 0
}

// SetDm DM
func (r *registerMaccrType) SetDm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldDmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldDmMask)
	}
}

const (
	RegisterMaccrFieldFesShift = 14
	RegisterMaccrFieldFesMask  = 0x4000
)

// GetFes FES
func (r *registerMaccrType) GetFes() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldFesMask) != 0
}

// SetFes FES
func (r *registerMaccrType) SetFes(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldFesMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldFesMask)
	}
}

const (
	RegisterMaccrFieldJeShift = 16
	RegisterMaccrFieldJeMask  = 0x10000
)

// GetJe JE
func (r *registerMaccrType) GetJe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldJeMask) != 0
}

// SetJe JE
func (r *registerMaccrType) SetJe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldJeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldJeMask)
	}
}

const (
	RegisterMaccrFieldJdShift = 17
	RegisterMaccrFieldJdMask  = 0x20000
)

// GetJd JD
func (r *registerMaccrType) GetJd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldJdMask) != 0
}

// SetJd JD
func (r *registerMaccrType) SetJd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldJdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldJdMask)
	}
}

const (
	RegisterMaccrFieldWdShift = 19
	RegisterMaccrFieldWdMask  = 0x80000
)

// GetWd WD
func (r *registerMaccrType) GetWd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldWdMask) != 0
}

// SetWd WD
func (r *registerMaccrType) SetWd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldWdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldWdMask)
	}
}

const (
	RegisterMaccrFieldAcsShift = 20
	RegisterMaccrFieldAcsMask  = 0x100000
)

// GetAcs ACS
func (r *registerMaccrType) GetAcs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldAcsMask) != 0
}

// SetAcs ACS
func (r *registerMaccrType) SetAcs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldAcsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldAcsMask)
	}
}

const (
	RegisterMaccrFieldCstShift = 21
	RegisterMaccrFieldCstMask  = 0x200000
)

// GetCst CST
func (r *registerMaccrType) GetCst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldCstMask) != 0
}

// SetCst CST
func (r *registerMaccrType) SetCst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldCstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldCstMask)
	}
}

const (
	RegisterMaccrFieldS2kpShift = 22
	RegisterMaccrFieldS2kpMask  = 0x400000
)

// GetS2kp S2KP
func (r *registerMaccrType) GetS2kp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldS2kpMask) != 0
}

// SetS2kp S2KP
func (r *registerMaccrType) SetS2kp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldS2kpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldS2kpMask)
	}
}

const (
	RegisterMaccrFieldGpslceShift = 23
	RegisterMaccrFieldGpslceMask  = 0x800000
)

// GetGpslce GPSLCE
func (r *registerMaccrType) GetGpslce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldGpslceMask) != 0
}

// SetGpslce GPSLCE
func (r *registerMaccrType) SetGpslce(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldGpslceMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldGpslceMask)
	}
}

const (
	RegisterMaccrFieldIpgShift = 24
	RegisterMaccrFieldIpgMask  = 0x7000000
)

// GetIpg IPG
func (r *registerMaccrType) GetIpg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldIpgMask) >> RegisterMaccrFieldIpgShift)
}

// SetIpg IPG
func (r *registerMaccrType) SetIpg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldIpgMask)|(uint32(value)<<RegisterMaccrFieldIpgShift))
}

const (
	RegisterMaccrFieldIpcShift = 27
	RegisterMaccrFieldIpcMask  = 0x8000000
)

// GetIpc IPC
func (r *registerMaccrType) GetIpc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldIpcMask) != 0
}

// SetIpc IPC
func (r *registerMaccrType) SetIpc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldIpcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldIpcMask)
	}
}

const (
	RegisterMaccrFieldSarcShift = 28
	RegisterMaccrFieldSarcMask  = 0x70000000
)

// GetSarc SARC
func (r *registerMaccrType) GetSarc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldSarcMask) >> RegisterMaccrFieldSarcShift)
}

// SetSarc SARC
func (r *registerMaccrType) SetSarc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldSarcMask)|(uint32(value)<<RegisterMaccrFieldSarcShift))
}

const (
	RegisterMaccrFieldArpenShift = 31
	RegisterMaccrFieldArpenMask  = 0x80000000
)

// GetArpen ARPEN
func (r *registerMaccrType) GetArpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldArpenMask) != 0
}

// SetArpen ARPEN
func (r *registerMaccrType) SetArpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldArpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldArpenMask)
	}
}

// registerMacecrType Extended operating mode configuration register
type registerMacecrType uint32

const (
	RegisterMacecrFieldGpslShift = 0
	RegisterMacecrFieldGpslMask  = 0x3fff
)

// GetGpsl GPSL
func (r *registerMacecrType) GetGpsl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldGpslMask) >> RegisterMacecrFieldGpslShift)
}

// SetGpsl GPSL
func (r *registerMacecrType) SetGpsl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacecrFieldGpslMask)|(uint32(value)<<RegisterMacecrFieldGpslShift))
}

const (
	RegisterMacecrFieldDcrccShift = 16
	RegisterMacecrFieldDcrccMask  = 0x10000
)

// GetDcrcc DCRCC
func (r *registerMacecrType) GetDcrcc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldDcrccMask) != 0
}

// SetDcrcc DCRCC
func (r *registerMacecrType) SetDcrcc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacecrFieldDcrccMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacecrFieldDcrccMask)
	}
}

const (
	RegisterMacecrFieldSpenShift = 17
	RegisterMacecrFieldSpenMask  = 0x20000
)

// GetSpen SPEN
func (r *registerMacecrType) GetSpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldSpenMask) != 0
}

// SetSpen SPEN
func (r *registerMacecrType) SetSpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacecrFieldSpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacecrFieldSpenMask)
	}
}

const (
	RegisterMacecrFieldUspShift = 18
	RegisterMacecrFieldUspMask  = 0x40000
)

// GetUsp USP
func (r *registerMacecrType) GetUsp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldUspMask) != 0
}

// SetUsp USP
func (r *registerMacecrType) SetUsp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacecrFieldUspMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacecrFieldUspMask)
	}
}

const (
	RegisterMacecrFieldEipgenShift = 24
	RegisterMacecrFieldEipgenMask  = 0x1000000
)

// GetEipgen EIPGEN
func (r *registerMacecrType) GetEipgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldEipgenMask) != 0
}

// SetEipgen EIPGEN
func (r *registerMacecrType) SetEipgen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacecrFieldEipgenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacecrFieldEipgenMask)
	}
}

const (
	RegisterMacecrFieldEipgShift = 25
	RegisterMacecrFieldEipgMask  = 0x3e000000
)

// GetEipg EIPG
func (r *registerMacecrType) GetEipg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldEipgMask) >> RegisterMacecrFieldEipgShift)
}

// SetEipg EIPG
func (r *registerMacecrType) SetEipg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacecrFieldEipgMask)|(uint32(value)<<RegisterMacecrFieldEipgShift))
}

// registerMacpfrType Packet filtering control register
type registerMacpfrType uint32

const (
	RegisterMacpfrFieldPrShift = 0
	RegisterMacpfrFieldPrMask  = 0x1
)

// GetPr PR
func (r *registerMacpfrType) GetPr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldPrMask) != 0
}

// SetPr PR
func (r *registerMacpfrType) SetPr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldPrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldPrMask)
	}
}

const (
	RegisterMacpfrFieldHucShift = 1
	RegisterMacpfrFieldHucMask  = 0x2
)

// GetHuc HUC
func (r *registerMacpfrType) GetHuc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldHucMask) != 0
}

// SetHuc HUC
func (r *registerMacpfrType) SetHuc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldHucMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldHucMask)
	}
}

const (
	RegisterMacpfrFieldHmcShift = 2
	RegisterMacpfrFieldHmcMask  = 0x4
)

// GetHmc HMC
func (r *registerMacpfrType) GetHmc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldHmcMask) != 0
}

// SetHmc HMC
func (r *registerMacpfrType) SetHmc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldHmcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldHmcMask)
	}
}

const (
	RegisterMacpfrFieldDaifShift = 3
	RegisterMacpfrFieldDaifMask  = 0x8
)

// GetDaif DAIF
func (r *registerMacpfrType) GetDaif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldDaifMask) != 0
}

// SetDaif DAIF
func (r *registerMacpfrType) SetDaif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldDaifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldDaifMask)
	}
}

const (
	RegisterMacpfrFieldPmShift = 4
	RegisterMacpfrFieldPmMask  = 0x10
)

// GetPm PM
func (r *registerMacpfrType) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldPmMask) != 0
}

// SetPm PM
func (r *registerMacpfrType) SetPm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldPmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldPmMask)
	}
}

const (
	RegisterMacpfrFieldDbfShift = 5
	RegisterMacpfrFieldDbfMask  = 0x20
)

// GetDbf DBF
func (r *registerMacpfrType) GetDbf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldDbfMask) != 0
}

// SetDbf DBF
func (r *registerMacpfrType) SetDbf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldDbfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldDbfMask)
	}
}

const (
	RegisterMacpfrFieldPcfShift = 6
	RegisterMacpfrFieldPcfMask  = 0xc0
)

// GetPcf PCF
func (r *registerMacpfrType) GetPcf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldPcfMask) >> RegisterMacpfrFieldPcfShift)
}

// SetPcf PCF
func (r *registerMacpfrType) SetPcf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldPcfMask)|(uint32(value)<<RegisterMacpfrFieldPcfShift))
}

const (
	RegisterMacpfrFieldSaifShift = 8
	RegisterMacpfrFieldSaifMask  = 0x100
)

// GetSaif SAIF
func (r *registerMacpfrType) GetSaif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldSaifMask) != 0
}

// SetSaif SAIF
func (r *registerMacpfrType) SetSaif(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldSaifMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldSaifMask)
	}
}

const (
	RegisterMacpfrFieldSafShift = 9
	RegisterMacpfrFieldSafMask  = 0x200
)

// GetSaf SAF
func (r *registerMacpfrType) GetSaf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldSafMask) != 0
}

// SetSaf SAF
func (r *registerMacpfrType) SetSaf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldSafMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldSafMask)
	}
}

const (
	RegisterMacpfrFieldHpfShift = 10
	RegisterMacpfrFieldHpfMask  = 0x400
)

// GetHpf HPF
func (r *registerMacpfrType) GetHpf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldHpfMask) != 0
}

// SetHpf HPF
func (r *registerMacpfrType) SetHpf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldHpfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldHpfMask)
	}
}

const (
	RegisterMacpfrFieldVtfeShift = 16
	RegisterMacpfrFieldVtfeMask  = 0x10000
)

// GetVtfe VTFE
func (r *registerMacpfrType) GetVtfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldVtfeMask) != 0
}

// SetVtfe VTFE
func (r *registerMacpfrType) SetVtfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldVtfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldVtfeMask)
	}
}

const (
	RegisterMacpfrFieldIpfeShift = 20
	RegisterMacpfrFieldIpfeMask  = 0x100000
)

// GetIpfe IPFE
func (r *registerMacpfrType) GetIpfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldIpfeMask) != 0
}

// SetIpfe IPFE
func (r *registerMacpfrType) SetIpfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldIpfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldIpfeMask)
	}
}

const (
	RegisterMacpfrFieldDntuShift = 21
	RegisterMacpfrFieldDntuMask  = 0x200000
)

// GetDntu DNTU
func (r *registerMacpfrType) GetDntu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldDntuMask) != 0
}

// SetDntu DNTU
func (r *registerMacpfrType) SetDntu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldDntuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldDntuMask)
	}
}

const (
	RegisterMacpfrFieldRaShift = 31
	RegisterMacpfrFieldRaMask  = 0x80000000
)

// GetRa RA
func (r *registerMacpfrType) GetRa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldRaMask) != 0
}

// SetRa RA
func (r *registerMacpfrType) SetRa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldRaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldRaMask)
	}
}

// registerMacwtrType Watchdog timeout register
type registerMacwtrType uint32

const (
	RegisterMacwtrFieldWtoShift = 0
	RegisterMacwtrFieldWtoMask  = 0xf
)

// GetWto WTO
func (r *registerMacwtrType) GetWto() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacwtrFieldWtoMask) >> RegisterMacwtrFieldWtoShift)
}

// SetWto WTO
func (r *registerMacwtrType) SetWto(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacwtrFieldWtoMask)|(uint32(value)<<RegisterMacwtrFieldWtoShift))
}

const (
	RegisterMacwtrFieldPweShift = 8
	RegisterMacwtrFieldPweMask  = 0x100
)

// GetPwe PWE
func (r *registerMacwtrType) GetPwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacwtrFieldPweMask) != 0
}

// SetPwe PWE
func (r *registerMacwtrType) SetPwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacwtrFieldPweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacwtrFieldPweMask)
	}
}

// registerMacht0rType Hash Table 0 register
type registerMacht0rType uint32

const (
	RegisterMacht0rFieldHt31t0Shift = 0
	RegisterMacht0rFieldHt31t0Mask  = 0xffffffff
)

// GetHt31t0 HT31T0
func (r *registerMacht0rType) GetHt31t0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacht0rFieldHt31t0Mask) >> RegisterMacht0rFieldHt31t0Shift)
}

// SetHt31t0 HT31T0
func (r *registerMacht0rType) SetHt31t0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacht0rFieldHt31t0Mask)|(uint32(value)<<RegisterMacht0rFieldHt31t0Shift))
}

// registerMacht1rType Hash Table 1 register
type registerMacht1rType uint32

const (
	RegisterMacht1rFieldHt63t32Shift = 0
	RegisterMacht1rFieldHt63t32Mask  = 0xffffffff
)

// GetHt63t32 HT63T32
func (r *registerMacht1rType) GetHt63t32() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacht1rFieldHt63t32Mask) >> RegisterMacht1rFieldHt63t32Shift)
}

// SetHt63t32 HT63T32
func (r *registerMacht1rType) SetHt63t32(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacht1rFieldHt63t32Mask)|(uint32(value)<<RegisterMacht1rFieldHt63t32Shift))
}

// registerMacvtrType VLAN tag register
type registerMacvtrType uint32

const (
	RegisterMacvtrFieldVlShift = 0
	RegisterMacvtrFieldVlMask  = 0xffff
)

// GetVl VL
func (r *registerMacvtrType) GetVl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldVlMask) >> RegisterMacvtrFieldVlShift)
}

// SetVl VL
func (r *registerMacvtrType) SetVl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldVlMask)|(uint32(value)<<RegisterMacvtrFieldVlShift))
}

const (
	RegisterMacvtrFieldEtvShift = 16
	RegisterMacvtrFieldEtvMask  = 0x10000
)

// GetEtv ETV
func (r *registerMacvtrType) GetEtv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEtvMask) != 0
}

// SetEtv ETV
func (r *registerMacvtrType) SetEtv(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvtrFieldEtvMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldEtvMask)
	}
}

const (
	RegisterMacvtrFieldVtimShift = 17
	RegisterMacvtrFieldVtimMask  = 0x20000
)

// GetVtim VTIM
func (r *registerMacvtrType) GetVtim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldVtimMask) != 0
}

// SetVtim VTIM
func (r *registerMacvtrType) SetVtim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvtrFieldVtimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldVtimMask)
	}
}

const (
	RegisterMacvtrFieldEsvlShift = 18
	RegisterMacvtrFieldEsvlMask  = 0x40000
)

// GetEsvl ESVL
func (r *registerMacvtrType) GetEsvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEsvlMask) != 0
}

// SetEsvl ESVL
func (r *registerMacvtrType) SetEsvl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvtrFieldEsvlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldEsvlMask)
	}
}

const (
	RegisterMacvtrFieldErsvlmShift = 19
	RegisterMacvtrFieldErsvlmMask  = 0x80000
)

// GetErsvlm ERSVLM
func (r *registerMacvtrType) GetErsvlm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldErsvlmMask) != 0
}

// SetErsvlm ERSVLM
func (r *registerMacvtrType) SetErsvlm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvtrFieldErsvlmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldErsvlmMask)
	}
}

const (
	RegisterMacvtrFieldDovltcShift = 20
	RegisterMacvtrFieldDovltcMask  = 0x100000
)

// GetDovltc DOVLTC
func (r *registerMacvtrType) GetDovltc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldDovltcMask) != 0
}

// SetDovltc DOVLTC
func (r *registerMacvtrType) SetDovltc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvtrFieldDovltcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldDovltcMask)
	}
}

const (
	RegisterMacvtrFieldEvlsShift = 21
	RegisterMacvtrFieldEvlsMask  = 0x600000
)

// GetEvls EVLS
func (r *registerMacvtrType) GetEvls() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEvlsMask) >> RegisterMacvtrFieldEvlsShift)
}

// SetEvls EVLS
func (r *registerMacvtrType) SetEvls(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldEvlsMask)|(uint32(value)<<RegisterMacvtrFieldEvlsShift))
}

const (
	RegisterMacvtrFieldEvlrxsShift = 24
	RegisterMacvtrFieldEvlrxsMask  = 0x1000000
)

// GetEvlrxs EVLRXS
func (r *registerMacvtrType) GetEvlrxs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEvlrxsMask) != 0
}

// SetEvlrxs EVLRXS
func (r *registerMacvtrType) SetEvlrxs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvtrFieldEvlrxsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldEvlrxsMask)
	}
}

const (
	RegisterMacvtrFieldVthmShift = 25
	RegisterMacvtrFieldVthmMask  = 0x2000000
)

// GetVthm VTHM
func (r *registerMacvtrType) GetVthm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldVthmMask) != 0
}

// SetVthm VTHM
func (r *registerMacvtrType) SetVthm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvtrFieldVthmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldVthmMask)
	}
}

const (
	RegisterMacvtrFieldEdvlpShift = 26
	RegisterMacvtrFieldEdvlpMask  = 0x4000000
)

// GetEdvlp EDVLP
func (r *registerMacvtrType) GetEdvlp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEdvlpMask) != 0
}

// SetEdvlp EDVLP
func (r *registerMacvtrType) SetEdvlp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvtrFieldEdvlpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldEdvlpMask)
	}
}

const (
	RegisterMacvtrFieldErivltShift = 27
	RegisterMacvtrFieldErivltMask  = 0x8000000
)

// GetErivlt ERIVLT
func (r *registerMacvtrType) GetErivlt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldErivltMask) != 0
}

// SetErivlt ERIVLT
func (r *registerMacvtrType) SetErivlt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvtrFieldErivltMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldErivltMask)
	}
}

const (
	RegisterMacvtrFieldEivlsShift = 28
	RegisterMacvtrFieldEivlsMask  = 0x30000000
)

// GetEivls EIVLS
func (r *registerMacvtrType) GetEivls() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEivlsMask) >> RegisterMacvtrFieldEivlsShift)
}

// SetEivls EIVLS
func (r *registerMacvtrType) SetEivls(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldEivlsMask)|(uint32(value)<<RegisterMacvtrFieldEivlsShift))
}

const (
	RegisterMacvtrFieldEivlrxsShift = 31
	RegisterMacvtrFieldEivlrxsMask  = 0x80000000
)

// GetEivlrxs EIVLRXS
func (r *registerMacvtrType) GetEivlrxs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEivlrxsMask) != 0
}

// SetEivlrxs EIVLRXS
func (r *registerMacvtrType) SetEivlrxs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvtrFieldEivlrxsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldEivlrxsMask)
	}
}

// registerMacvhtrType VLAN Hash table register
type registerMacvhtrType uint32

const (
	RegisterMacvhtrFieldVlhtShift = 0
	RegisterMacvhtrFieldVlhtMask  = 0xffff
)

// GetVlht VLHT
func (r *registerMacvhtrType) GetVlht() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacvhtrFieldVlhtMask) >> RegisterMacvhtrFieldVlhtShift)
}

// SetVlht VLHT
func (r *registerMacvhtrType) SetVlht(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvhtrFieldVlhtMask)|(uint32(value)<<RegisterMacvhtrFieldVlhtShift))
}

// registerMacvirType VLAN inclusion register
type registerMacvirType uint32

const (
	RegisterMacvirFieldVltShift = 0
	RegisterMacvirFieldVltMask  = 0xffff
)

// GetVlt VLT
func (r *registerMacvirType) GetVlt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacvirFieldVltMask) >> RegisterMacvirFieldVltShift)
}

// SetVlt VLT
func (r *registerMacvirType) SetVlt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvirFieldVltMask)|(uint32(value)<<RegisterMacvirFieldVltShift))
}

const (
	RegisterMacvirFieldVlcShift = 16
	RegisterMacvirFieldVlcMask  = 0x30000
)

// GetVlc VLC
func (r *registerMacvirType) GetVlc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacvirFieldVlcMask) >> RegisterMacvirFieldVlcShift)
}

// SetVlc VLC
func (r *registerMacvirType) SetVlc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvirFieldVlcMask)|(uint32(value)<<RegisterMacvirFieldVlcShift))
}

const (
	RegisterMacvirFieldVlpShift = 18
	RegisterMacvirFieldVlpMask  = 0x40000
)

// GetVlp VLP
func (r *registerMacvirType) GetVlp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvirFieldVlpMask) != 0
}

// SetVlp VLP
func (r *registerMacvirType) SetVlp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvirFieldVlpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvirFieldVlpMask)
	}
}

const (
	RegisterMacvirFieldCsvlShift = 19
	RegisterMacvirFieldCsvlMask  = 0x80000
)

// GetCsvl CSVL
func (r *registerMacvirType) GetCsvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvirFieldCsvlMask) != 0
}

// SetCsvl CSVL
func (r *registerMacvirType) SetCsvl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvirFieldCsvlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvirFieldCsvlMask)
	}
}

const (
	RegisterMacvirFieldVltiShift = 20
	RegisterMacvirFieldVltiMask  = 0x100000
)

// GetVlti VLTI
func (r *registerMacvirType) GetVlti() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvirFieldVltiMask) != 0
}

// SetVlti VLTI
func (r *registerMacvirType) SetVlti(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvirFieldVltiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvirFieldVltiMask)
	}
}

// registerMacivirType Inner VLAN inclusion register
type registerMacivirType uint32

const (
	RegisterMacivirFieldVltShift = 0
	RegisterMacivirFieldVltMask  = 0xffff
)

// GetVlt VLT
func (r *registerMacivirType) GetVlt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacivirFieldVltMask) >> RegisterMacivirFieldVltShift)
}

// SetVlt VLT
func (r *registerMacivirType) SetVlt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacivirFieldVltMask)|(uint32(value)<<RegisterMacivirFieldVltShift))
}

const (
	RegisterMacivirFieldVlcShift = 16
	RegisterMacivirFieldVlcMask  = 0x30000
)

// GetVlc VLC
func (r *registerMacivirType) GetVlc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacivirFieldVlcMask) >> RegisterMacivirFieldVlcShift)
}

// SetVlc VLC
func (r *registerMacivirType) SetVlc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacivirFieldVlcMask)|(uint32(value)<<RegisterMacivirFieldVlcShift))
}

const (
	RegisterMacivirFieldVlpShift = 18
	RegisterMacivirFieldVlpMask  = 0x40000
)

// GetVlp VLP
func (r *registerMacivirType) GetVlp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacivirFieldVlpMask) != 0
}

// SetVlp VLP
func (r *registerMacivirType) SetVlp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacivirFieldVlpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacivirFieldVlpMask)
	}
}

const (
	RegisterMacivirFieldCsvlShift = 19
	RegisterMacivirFieldCsvlMask  = 0x80000
)

// GetCsvl CSVL
func (r *registerMacivirType) GetCsvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacivirFieldCsvlMask) != 0
}

// SetCsvl CSVL
func (r *registerMacivirType) SetCsvl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacivirFieldCsvlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacivirFieldCsvlMask)
	}
}

const (
	RegisterMacivirFieldVltiShift = 20
	RegisterMacivirFieldVltiMask  = 0x100000
)

// GetVlti VLTI
func (r *registerMacivirType) GetVlti() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacivirFieldVltiMask) != 0
}

// SetVlti VLTI
func (r *registerMacivirType) SetVlti(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacivirFieldVltiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacivirFieldVltiMask)
	}
}

// registerMacqtxfcrType Tx Queue flow control register
type registerMacqtxfcrType uint32

const (
	RegisterMacqtxfcrFieldFcbbpaShift = 0
	RegisterMacqtxfcrFieldFcbbpaMask  = 0x1
)

// GetFcbbpa FCB_BPA
func (r *registerMacqtxfcrType) GetFcbbpa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacqtxfcrFieldFcbbpaMask) != 0
}

// SetFcbbpa FCB_BPA
func (r *registerMacqtxfcrType) SetFcbbpa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacqtxfcrFieldFcbbpaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacqtxfcrFieldFcbbpaMask)
	}
}

const (
	RegisterMacqtxfcrFieldTfeShift = 1
	RegisterMacqtxfcrFieldTfeMask  = 0x2
)

// GetTfe TFE
func (r *registerMacqtxfcrType) GetTfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacqtxfcrFieldTfeMask) != 0
}

// SetTfe TFE
func (r *registerMacqtxfcrType) SetTfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacqtxfcrFieldTfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacqtxfcrFieldTfeMask)
	}
}

const (
	RegisterMacqtxfcrFieldPltShift = 4
	RegisterMacqtxfcrFieldPltMask  = 0x70
)

// GetPlt PLT
func (r *registerMacqtxfcrType) GetPlt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacqtxfcrFieldPltMask) >> RegisterMacqtxfcrFieldPltShift)
}

// SetPlt PLT
func (r *registerMacqtxfcrType) SetPlt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacqtxfcrFieldPltMask)|(uint32(value)<<RegisterMacqtxfcrFieldPltShift))
}

const (
	RegisterMacqtxfcrFieldDzpqShift = 7
	RegisterMacqtxfcrFieldDzpqMask  = 0x80
)

// GetDzpq DZPQ
func (r *registerMacqtxfcrType) GetDzpq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacqtxfcrFieldDzpqMask) != 0
}

// SetDzpq DZPQ
func (r *registerMacqtxfcrType) SetDzpq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacqtxfcrFieldDzpqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacqtxfcrFieldDzpqMask)
	}
}

const (
	RegisterMacqtxfcrFieldPtShift = 16
	RegisterMacqtxfcrFieldPtMask  = 0xffff0000
)

// GetPt PT
func (r *registerMacqtxfcrType) GetPt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacqtxfcrFieldPtMask) >> RegisterMacqtxfcrFieldPtShift)
}

// SetPt PT
func (r *registerMacqtxfcrType) SetPt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacqtxfcrFieldPtMask)|(uint32(value)<<RegisterMacqtxfcrFieldPtShift))
}

// registerMacrxfcrType Rx flow control register
type registerMacrxfcrType uint32

const (
	RegisterMacrxfcrFieldRfeShift = 0
	RegisterMacrxfcrFieldRfeMask  = 0x1
)

// GetRfe RFE
func (r *registerMacrxfcrType) GetRfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxfcrFieldRfeMask) != 0
}

// SetRfe RFE
func (r *registerMacrxfcrType) SetRfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacrxfcrFieldRfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacrxfcrFieldRfeMask)
	}
}

const (
	RegisterMacrxfcrFieldUpShift = 1
	RegisterMacrxfcrFieldUpMask  = 0x2
)

// GetUp UP
func (r *registerMacrxfcrType) GetUp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxfcrFieldUpMask) != 0
}

// SetUp UP
func (r *registerMacrxfcrType) SetUp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacrxfcrFieldUpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacrxfcrFieldUpMask)
	}
}

// registerMacisrType Interrupt status register
type registerMacisrType uint32

const (
	RegisterMacisrFieldPhyisShift = 3
	RegisterMacisrFieldPhyisMask  = 0x8
)

// GetPhyis PHYIS
func (r *registerMacisrType) GetPhyis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldPhyisMask) != 0
}

// SetPhyis PHYIS
func (r *registerMacisrType) SetPhyis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacisrFieldPhyisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacisrFieldPhyisMask)
	}
}

const (
	RegisterMacisrFieldPmtisShift = 4
	RegisterMacisrFieldPmtisMask  = 0x10
)

// GetPmtis PMTIS
func (r *registerMacisrType) GetPmtis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldPmtisMask) != 0
}

// SetPmtis PMTIS
func (r *registerMacisrType) SetPmtis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacisrFieldPmtisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacisrFieldPmtisMask)
	}
}

const (
	RegisterMacisrFieldLpiisShift = 5
	RegisterMacisrFieldLpiisMask  = 0x20
)

// GetLpiis LPIIS
func (r *registerMacisrType) GetLpiis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldLpiisMask) != 0
}

// SetLpiis LPIIS
func (r *registerMacisrType) SetLpiis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacisrFieldLpiisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacisrFieldLpiisMask)
	}
}

const (
	RegisterMacisrFieldMmcisShift = 8
	RegisterMacisrFieldMmcisMask  = 0x100
)

// GetMmcis MMCIS
func (r *registerMacisrType) GetMmcis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldMmcisMask) != 0
}

// SetMmcis MMCIS
func (r *registerMacisrType) SetMmcis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacisrFieldMmcisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacisrFieldMmcisMask)
	}
}

const (
	RegisterMacisrFieldMmcrxisShift = 9
	RegisterMacisrFieldMmcrxisMask  = 0x200
)

// GetMmcrxis MMCRXIS
func (r *registerMacisrType) GetMmcrxis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldMmcrxisMask) != 0
}

// SetMmcrxis MMCRXIS
func (r *registerMacisrType) SetMmcrxis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacisrFieldMmcrxisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacisrFieldMmcrxisMask)
	}
}

const (
	RegisterMacisrFieldMmctxisShift = 10
	RegisterMacisrFieldMmctxisMask  = 0x400
)

// GetMmctxis MMCTXIS
func (r *registerMacisrType) GetMmctxis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldMmctxisMask) != 0
}

// SetMmctxis MMCTXIS
func (r *registerMacisrType) SetMmctxis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacisrFieldMmctxisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacisrFieldMmctxisMask)
	}
}

const (
	RegisterMacisrFieldTsisShift = 12
	RegisterMacisrFieldTsisMask  = 0x1000
)

// GetTsis TSIS
func (r *registerMacisrType) GetTsis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldTsisMask) != 0
}

// SetTsis TSIS
func (r *registerMacisrType) SetTsis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacisrFieldTsisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacisrFieldTsisMask)
	}
}

const (
	RegisterMacisrFieldTxstsisShift = 13
	RegisterMacisrFieldTxstsisMask  = 0x2000
)

// GetTxstsis TXSTSIS
func (r *registerMacisrType) GetTxstsis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldTxstsisMask) != 0
}

// SetTxstsis TXSTSIS
func (r *registerMacisrType) SetTxstsis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacisrFieldTxstsisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacisrFieldTxstsisMask)
	}
}

const (
	RegisterMacisrFieldRxstsisShift = 14
	RegisterMacisrFieldRxstsisMask  = 0x4000
)

// GetRxstsis RXSTSIS
func (r *registerMacisrType) GetRxstsis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldRxstsisMask) != 0
}

// SetRxstsis RXSTSIS
func (r *registerMacisrType) SetRxstsis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacisrFieldRxstsisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacisrFieldRxstsisMask)
	}
}

// registerMacierType Interrupt enable register
type registerMacierType uint32

const (
	RegisterMacierFieldPhyieShift = 3
	RegisterMacierFieldPhyieMask  = 0x8
)

// GetPhyie PHYIE
func (r *registerMacierType) GetPhyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldPhyieMask) != 0
}

// SetPhyie PHYIE
func (r *registerMacierType) SetPhyie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacierFieldPhyieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacierFieldPhyieMask)
	}
}

const (
	RegisterMacierFieldPmtieShift = 4
	RegisterMacierFieldPmtieMask  = 0x10
)

// GetPmtie PMTIE
func (r *registerMacierType) GetPmtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldPmtieMask) != 0
}

// SetPmtie PMTIE
func (r *registerMacierType) SetPmtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacierFieldPmtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacierFieldPmtieMask)
	}
}

const (
	RegisterMacierFieldLpiieShift = 5
	RegisterMacierFieldLpiieMask  = 0x20
)

// GetLpiie LPIIE
func (r *registerMacierType) GetLpiie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldLpiieMask) != 0
}

// SetLpiie LPIIE
func (r *registerMacierType) SetLpiie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacierFieldLpiieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacierFieldLpiieMask)
	}
}

const (
	RegisterMacierFieldTsieShift = 12
	RegisterMacierFieldTsieMask  = 0x1000
)

// GetTsie TSIE
func (r *registerMacierType) GetTsie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldTsieMask) != 0
}

// SetTsie TSIE
func (r *registerMacierType) SetTsie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacierFieldTsieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacierFieldTsieMask)
	}
}

const (
	RegisterMacierFieldTxstsieShift = 13
	RegisterMacierFieldTxstsieMask  = 0x2000
)

// GetTxstsie TXSTSIE
func (r *registerMacierType) GetTxstsie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldTxstsieMask) != 0
}

// SetTxstsie TXSTSIE
func (r *registerMacierType) SetTxstsie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacierFieldTxstsieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacierFieldTxstsieMask)
	}
}

const (
	RegisterMacierFieldRxstsieShift = 14
	RegisterMacierFieldRxstsieMask  = 0x4000
)

// GetRxstsie RXSTSIE
func (r *registerMacierType) GetRxstsie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldRxstsieMask) != 0
}

// SetRxstsie RXSTSIE
func (r *registerMacierType) SetRxstsie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacierFieldRxstsieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacierFieldRxstsieMask)
	}
}

// registerMacrxtxsrType Rx Tx status register
type registerMacrxtxsrType uint32

const (
	RegisterMacrxtxsrFieldTjtShift = 0
	RegisterMacrxtxsrFieldTjtMask  = 0x1
)

// GetTjt TJT
func (r *registerMacrxtxsrType) GetTjt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldTjtMask) != 0
}

// SetTjt TJT
func (r *registerMacrxtxsrType) SetTjt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacrxtxsrFieldTjtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacrxtxsrFieldTjtMask)
	}
}

const (
	RegisterMacrxtxsrFieldNcarrShift = 1
	RegisterMacrxtxsrFieldNcarrMask  = 0x2
)

// GetNcarr NCARR
func (r *registerMacrxtxsrType) GetNcarr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldNcarrMask) != 0
}

// SetNcarr NCARR
func (r *registerMacrxtxsrType) SetNcarr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacrxtxsrFieldNcarrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacrxtxsrFieldNcarrMask)
	}
}

const (
	RegisterMacrxtxsrFieldLcarrShift = 2
	RegisterMacrxtxsrFieldLcarrMask  = 0x4
)

// GetLcarr LCARR
func (r *registerMacrxtxsrType) GetLcarr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldLcarrMask) != 0
}

// SetLcarr LCARR
func (r *registerMacrxtxsrType) SetLcarr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacrxtxsrFieldLcarrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacrxtxsrFieldLcarrMask)
	}
}

const (
	RegisterMacrxtxsrFieldExdefShift = 3
	RegisterMacrxtxsrFieldExdefMask  = 0x8
)

// GetExdef EXDEF
func (r *registerMacrxtxsrType) GetExdef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldExdefMask) != 0
}

// SetExdef EXDEF
func (r *registerMacrxtxsrType) SetExdef(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacrxtxsrFieldExdefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacrxtxsrFieldExdefMask)
	}
}

const (
	RegisterMacrxtxsrFieldLcolShift = 4
	RegisterMacrxtxsrFieldLcolMask  = 0x10
)

// GetLcol LCOL
func (r *registerMacrxtxsrType) GetLcol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldLcolMask) != 0
}

// SetLcol LCOL
func (r *registerMacrxtxsrType) SetLcol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacrxtxsrFieldLcolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacrxtxsrFieldLcolMask)
	}
}

const (
	RegisterMacrxtxsrFieldExcolShift = 5
	RegisterMacrxtxsrFieldExcolMask  = 0x20
)

// GetExcol LCOL
func (r *registerMacrxtxsrType) GetExcol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldExcolMask) != 0
}

// SetExcol LCOL
func (r *registerMacrxtxsrType) SetExcol(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacrxtxsrFieldExcolMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacrxtxsrFieldExcolMask)
	}
}

const (
	RegisterMacrxtxsrFieldRwtShift = 8
	RegisterMacrxtxsrFieldRwtMask  = 0x100
)

// GetRwt RWT
func (r *registerMacrxtxsrType) GetRwt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldRwtMask) != 0
}

// SetRwt RWT
func (r *registerMacrxtxsrType) SetRwt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacrxtxsrFieldRwtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacrxtxsrFieldRwtMask)
	}
}

// registerMacpcsrType PMT control status register
type registerMacpcsrType uint32

const (
	RegisterMacpcsrFieldPwrdwnShift = 0
	RegisterMacpcsrFieldPwrdwnMask  = 0x1
)

// GetPwrdwn PWRDWN
func (r *registerMacpcsrType) GetPwrdwn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldPwrdwnMask) != 0
}

// SetPwrdwn PWRDWN
func (r *registerMacpcsrType) SetPwrdwn(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpcsrFieldPwrdwnMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpcsrFieldPwrdwnMask)
	}
}

const (
	RegisterMacpcsrFieldMgkpktenShift = 1
	RegisterMacpcsrFieldMgkpktenMask  = 0x2
)

// GetMgkpkten MGKPKTEN
func (r *registerMacpcsrType) GetMgkpkten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldMgkpktenMask) != 0
}

// SetMgkpkten MGKPKTEN
func (r *registerMacpcsrType) SetMgkpkten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpcsrFieldMgkpktenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpcsrFieldMgkpktenMask)
	}
}

const (
	RegisterMacpcsrFieldRwkpktenShift = 2
	RegisterMacpcsrFieldRwkpktenMask  = 0x4
)

// GetRwkpkten RWKPKTEN
func (r *registerMacpcsrType) GetRwkpkten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldRwkpktenMask) != 0
}

// SetRwkpkten RWKPKTEN
func (r *registerMacpcsrType) SetRwkpkten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpcsrFieldRwkpktenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpcsrFieldRwkpktenMask)
	}
}

const (
	RegisterMacpcsrFieldMgkprcvdShift = 5
	RegisterMacpcsrFieldMgkprcvdMask  = 0x20
)

// GetMgkprcvd MGKPRCVD
func (r *registerMacpcsrType) GetMgkprcvd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldMgkprcvdMask) != 0
}

const (
	RegisterMacpcsrFieldRwkprcvdShift = 6
	RegisterMacpcsrFieldRwkprcvdMask  = 0x40
)

// GetRwkprcvd RWKPRCVD
func (r *registerMacpcsrType) GetRwkprcvd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldRwkprcvdMask) != 0
}

const (
	RegisterMacpcsrFieldGlblucastShift = 9
	RegisterMacpcsrFieldGlblucastMask  = 0x200
)

// GetGlblucast GLBLUCAST
func (r *registerMacpcsrType) GetGlblucast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldGlblucastMask) != 0
}

// SetGlblucast GLBLUCAST
func (r *registerMacpcsrType) SetGlblucast(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpcsrFieldGlblucastMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpcsrFieldGlblucastMask)
	}
}

const (
	RegisterMacpcsrFieldRwkpfeShift = 10
	RegisterMacpcsrFieldRwkpfeMask  = 0x400
)

// GetRwkpfe RWKPFE
func (r *registerMacpcsrType) GetRwkpfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldRwkpfeMask) != 0
}

// SetRwkpfe RWKPFE
func (r *registerMacpcsrType) SetRwkpfe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpcsrFieldRwkpfeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpcsrFieldRwkpfeMask)
	}
}

const (
	RegisterMacpcsrFieldRwkptrShift = 24
	RegisterMacpcsrFieldRwkptrMask  = 0x1f000000
)

// GetRwkptr RWKPTR
func (r *registerMacpcsrType) GetRwkptr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldRwkptrMask) >> RegisterMacpcsrFieldRwkptrShift)
}

// SetRwkptr RWKPTR
func (r *registerMacpcsrType) SetRwkptr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacpcsrFieldRwkptrMask)|(uint32(value)<<RegisterMacpcsrFieldRwkptrShift))
}

const (
	RegisterMacpcsrFieldRwkfiltrstShift = 31
	RegisterMacpcsrFieldRwkfiltrstMask  = 0x80000000
)

// GetRwkfiltrst RWKFILTRST
func (r *registerMacpcsrType) GetRwkfiltrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldRwkfiltrstMask) != 0
}

// SetRwkfiltrst RWKFILTRST
func (r *registerMacpcsrType) SetRwkfiltrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpcsrFieldRwkfiltrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpcsrFieldRwkfiltrstMask)
	}
}

// registerMacrwkpfrType Remove wakeup packet filter register
type registerMacrwkpfrType uint32

const (
	RegisterMacrwkpfrFieldWkupfrmftrShift = 0
	RegisterMacrwkpfrFieldWkupfrmftrMask  = 0xffffffff
)

// GetWkupfrmftr WKUPFRMFTR
func (r *registerMacrwkpfrType) GetWkupfrmftr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacrwkpfrFieldWkupfrmftrMask) >> RegisterMacrwkpfrFieldWkupfrmftrShift)
}

// SetWkupfrmftr WKUPFRMFTR
func (r *registerMacrwkpfrType) SetWkupfrmftr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacrwkpfrFieldWkupfrmftrMask)|(uint32(value)<<RegisterMacrwkpfrFieldWkupfrmftrShift))
}

// registerMaclcsrType LPI control status register
type registerMaclcsrType uint32

const (
	RegisterMaclcsrFieldTlpienShift = 0
	RegisterMaclcsrFieldTlpienMask  = 0x1
)

// GetTlpien TLPIEN
func (r *registerMaclcsrType) GetTlpien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldTlpienMask) != 0
}

const (
	RegisterMaclcsrFieldTlpiexShift = 1
	RegisterMaclcsrFieldTlpiexMask  = 0x2
)

// GetTlpiex TLPIEX
func (r *registerMaclcsrType) GetTlpiex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldTlpiexMask) != 0
}

const (
	RegisterMaclcsrFieldRlpienShift = 2
	RegisterMaclcsrFieldRlpienMask  = 0x4
)

// GetRlpien RLPIEN
func (r *registerMaclcsrType) GetRlpien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldRlpienMask) != 0
}

const (
	RegisterMaclcsrFieldRlpiexShift = 3
	RegisterMaclcsrFieldRlpiexMask  = 0x8
)

// GetRlpiex RLPIEX
func (r *registerMaclcsrType) GetRlpiex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldRlpiexMask) != 0
}

const (
	RegisterMaclcsrFieldTlpistShift = 8
	RegisterMaclcsrFieldTlpistMask  = 0x100
)

// GetTlpist TLPIST
func (r *registerMaclcsrType) GetTlpist() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldTlpistMask) != 0
}

const (
	RegisterMaclcsrFieldRlpistShift = 9
	RegisterMaclcsrFieldRlpistMask  = 0x200
)

// GetRlpist RLPIST
func (r *registerMaclcsrType) GetRlpist() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldRlpistMask) != 0
}

const (
	RegisterMaclcsrFieldLpienShift = 16
	RegisterMaclcsrFieldLpienMask  = 0x10000
)

// GetLpien LPIEN
func (r *registerMaclcsrType) GetLpien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldLpienMask) != 0
}

// SetLpien LPIEN
func (r *registerMaclcsrType) SetLpien(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaclcsrFieldLpienMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaclcsrFieldLpienMask)
	}
}

const (
	RegisterMaclcsrFieldPlsShift = 17
	RegisterMaclcsrFieldPlsMask  = 0x20000
)

// GetPls PLS
func (r *registerMaclcsrType) GetPls() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldPlsMask) != 0
}

// SetPls PLS
func (r *registerMaclcsrType) SetPls(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaclcsrFieldPlsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaclcsrFieldPlsMask)
	}
}

const (
	RegisterMaclcsrFieldPlsenShift = 18
	RegisterMaclcsrFieldPlsenMask  = 0x40000
)

// GetPlsen PLSEN
func (r *registerMaclcsrType) GetPlsen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldPlsenMask) != 0
}

// SetPlsen PLSEN
func (r *registerMaclcsrType) SetPlsen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaclcsrFieldPlsenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaclcsrFieldPlsenMask)
	}
}

const (
	RegisterMaclcsrFieldLpitxaShift = 19
	RegisterMaclcsrFieldLpitxaMask  = 0x80000
)

// GetLpitxa LPITXA
func (r *registerMaclcsrType) GetLpitxa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldLpitxaMask) != 0
}

// SetLpitxa LPITXA
func (r *registerMaclcsrType) SetLpitxa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaclcsrFieldLpitxaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaclcsrFieldLpitxaMask)
	}
}

const (
	RegisterMaclcsrFieldLpiteShift = 20
	RegisterMaclcsrFieldLpiteMask  = 0x100000
)

// GetLpite LPITE
func (r *registerMaclcsrType) GetLpite() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldLpiteMask) != 0
}

// SetLpite LPITE
func (r *registerMaclcsrType) SetLpite(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaclcsrFieldLpiteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaclcsrFieldLpiteMask)
	}
}

const (
	RegisterMaclcsrFieldLpitcseShift = 21
	RegisterMaclcsrFieldLpitcseMask  = 0x200000
)

// GetLpitcse LPITCSE
func (r *registerMaclcsrType) GetLpitcse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldLpitcseMask) != 0
}

// SetLpitcse LPITCSE
func (r *registerMaclcsrType) SetLpitcse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaclcsrFieldLpitcseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaclcsrFieldLpitcseMask)
	}
}

// registerMacltcrType LPI timers control register
type registerMacltcrType uint32

const (
	RegisterMacltcrFieldTwtShift = 0
	RegisterMacltcrFieldTwtMask  = 0xffff
)

// GetTwt TWT
func (r *registerMacltcrType) GetTwt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacltcrFieldTwtMask) >> RegisterMacltcrFieldTwtShift)
}

// SetTwt TWT
func (r *registerMacltcrType) SetTwt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacltcrFieldTwtMask)|(uint32(value)<<RegisterMacltcrFieldTwtShift))
}

const (
	RegisterMacltcrFieldLstShift = 16
	RegisterMacltcrFieldLstMask  = 0x3ff0000
)

// GetLst LST
func (r *registerMacltcrType) GetLst() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacltcrFieldLstMask) >> RegisterMacltcrFieldLstShift)
}

// SetLst LST
func (r *registerMacltcrType) SetLst(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacltcrFieldLstMask)|(uint32(value)<<RegisterMacltcrFieldLstShift))
}

// registerMacletrType LPI entry timer register
type registerMacletrType uint32

const (
	RegisterMacletrFieldLpietShift = 0
	RegisterMacletrFieldLpietMask  = 0x1ffff
)

// GetLpiet LPIET
func (r *registerMacletrType) GetLpiet() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacletrFieldLpietMask) >> RegisterMacletrFieldLpietShift)
}

// SetLpiet LPIET
func (r *registerMacletrType) SetLpiet(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacletrFieldLpietMask)|(uint32(value)<<RegisterMacletrFieldLpietShift))
}

// registerMac1ustcrType 1-microsecond-tick counter register
type registerMac1ustcrType uint32

const (
	RegisterMac1ustcrFieldTic1uscntrShift = 0
	RegisterMac1ustcrFieldTic1uscntrMask  = 0xfff
)

// GetTic1uscntr TIC_1US_CNTR
func (r *registerMac1ustcrType) GetTic1uscntr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMac1ustcrFieldTic1uscntrMask) >> RegisterMac1ustcrFieldTic1uscntrShift)
}

// SetTic1uscntr TIC_1US_CNTR
func (r *registerMac1ustcrType) SetTic1uscntr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMac1ustcrFieldTic1uscntrMask)|(uint32(value)<<RegisterMac1ustcrFieldTic1uscntrShift))
}

// registerMacvrType Version register
type registerMacvrType uint32

const (
	RegisterMacvrFieldSnpsverShift = 0
	RegisterMacvrFieldSnpsverMask  = 0xff
)

// GetSnpsver SNPSVER
func (r *registerMacvrType) GetSnpsver() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacvrFieldSnpsverMask) >> RegisterMacvrFieldSnpsverShift)
}

// SetSnpsver SNPSVER
func (r *registerMacvrType) SetSnpsver(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvrFieldSnpsverMask)|(uint32(value)<<RegisterMacvrFieldSnpsverShift))
}

const (
	RegisterMacvrFieldUserverShift = 8
	RegisterMacvrFieldUserverMask  = 0xff00
)

// GetUserver USERVER
func (r *registerMacvrType) GetUserver() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacvrFieldUserverMask) >> RegisterMacvrFieldUserverShift)
}

// SetUserver USERVER
func (r *registerMacvrType) SetUserver(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvrFieldUserverMask)|(uint32(value)<<RegisterMacvrFieldUserverShift))
}

// registerMacdrType Debug register
type registerMacdrType uint32

const (
	RegisterMacdrFieldRpestsShift = 0
	RegisterMacdrFieldRpestsMask  = 0x1
)

// GetRpests RPESTS
func (r *registerMacdrType) GetRpests() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacdrFieldRpestsMask) != 0
}

// SetRpests RPESTS
func (r *registerMacdrType) SetRpests(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacdrFieldRpestsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacdrFieldRpestsMask)
	}
}

const (
	RegisterMacdrFieldRfcfcstsShift = 1
	RegisterMacdrFieldRfcfcstsMask  = 0x6
)

// GetRfcfcsts RFCFCSTS
func (r *registerMacdrType) GetRfcfcsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacdrFieldRfcfcstsMask) >> RegisterMacdrFieldRfcfcstsShift)
}

// SetRfcfcsts RFCFCSTS
func (r *registerMacdrType) SetRfcfcsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacdrFieldRfcfcstsMask)|(uint32(value)<<RegisterMacdrFieldRfcfcstsShift))
}

const (
	RegisterMacdrFieldTpestsShift = 16
	RegisterMacdrFieldTpestsMask  = 0x10000
)

// GetTpests TPESTS
func (r *registerMacdrType) GetTpests() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacdrFieldTpestsMask) != 0
}

// SetTpests TPESTS
func (r *registerMacdrType) SetTpests(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacdrFieldTpestsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacdrFieldTpestsMask)
	}
}

const (
	RegisterMacdrFieldTfcstsShift = 17
	RegisterMacdrFieldTfcstsMask  = 0x60000
)

// GetTfcsts TFCSTS
func (r *registerMacdrType) GetTfcsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacdrFieldTfcstsMask) >> RegisterMacdrFieldTfcstsShift)
}

// SetTfcsts TFCSTS
func (r *registerMacdrType) SetTfcsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacdrFieldTfcstsMask)|(uint32(value)<<RegisterMacdrFieldTfcstsShift))
}

// registerMachwf1rType HW feature 1 register
type registerMachwf1rType uint32

const (
	RegisterMachwf1rFieldRxfifosizeShift = 0
	RegisterMachwf1rFieldRxfifosizeMask  = 0x1f
)

// GetRxfifosize RXFIFOSIZE
func (r *registerMachwf1rType) GetRxfifosize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldRxfifosizeMask) >> RegisterMachwf1rFieldRxfifosizeShift)
}

// SetRxfifosize RXFIFOSIZE
func (r *registerMachwf1rType) SetRxfifosize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldRxfifosizeMask)|(uint32(value)<<RegisterMachwf1rFieldRxfifosizeShift))
}

const (
	RegisterMachwf1rFieldTxfifosizeShift = 6
	RegisterMachwf1rFieldTxfifosizeMask  = 0x7c0
)

// GetTxfifosize TXFIFOSIZE
func (r *registerMachwf1rType) GetTxfifosize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldTxfifosizeMask) >> RegisterMachwf1rFieldTxfifosizeShift)
}

// SetTxfifosize TXFIFOSIZE
func (r *registerMachwf1rType) SetTxfifosize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldTxfifosizeMask)|(uint32(value)<<RegisterMachwf1rFieldTxfifosizeShift))
}

const (
	RegisterMachwf1rFieldOstenShift = 11
	RegisterMachwf1rFieldOstenMask  = 0x800
)

// GetOsten OSTEN
func (r *registerMachwf1rType) GetOsten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldOstenMask) != 0
}

// SetOsten OSTEN
func (r *registerMachwf1rType) SetOsten(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMachwf1rFieldOstenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldOstenMask)
	}
}

const (
	RegisterMachwf1rFieldPtoenShift = 12
	RegisterMachwf1rFieldPtoenMask  = 0x1000
)

// GetPtoen PTOEN
func (r *registerMachwf1rType) GetPtoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldPtoenMask) != 0
}

// SetPtoen PTOEN
func (r *registerMachwf1rType) SetPtoen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMachwf1rFieldPtoenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldPtoenMask)
	}
}

const (
	RegisterMachwf1rFieldAdvthwordShift = 13
	RegisterMachwf1rFieldAdvthwordMask  = 0x2000
)

// GetAdvthword ADVTHWORD
func (r *registerMachwf1rType) GetAdvthword() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldAdvthwordMask) != 0
}

// SetAdvthword ADVTHWORD
func (r *registerMachwf1rType) SetAdvthword(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMachwf1rFieldAdvthwordMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldAdvthwordMask)
	}
}

const (
	RegisterMachwf1rFieldDcbenShift = 16
	RegisterMachwf1rFieldDcbenMask  = 0x10000
)

// GetDcben DCBEN
func (r *registerMachwf1rType) GetDcben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldDcbenMask) != 0
}

// SetDcben DCBEN
func (r *registerMachwf1rType) SetDcben(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMachwf1rFieldDcbenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldDcbenMask)
	}
}

const (
	RegisterMachwf1rFieldSphenShift = 17
	RegisterMachwf1rFieldSphenMask  = 0x20000
)

// GetSphen SPHEN
func (r *registerMachwf1rType) GetSphen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldSphenMask) != 0
}

// SetSphen SPHEN
func (r *registerMachwf1rType) SetSphen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMachwf1rFieldSphenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldSphenMask)
	}
}

const (
	RegisterMachwf1rFieldTsoenShift = 18
	RegisterMachwf1rFieldTsoenMask  = 0x40000
)

// GetTsoen TSOEN
func (r *registerMachwf1rType) GetTsoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldTsoenMask) != 0
}

// SetTsoen TSOEN
func (r *registerMachwf1rType) SetTsoen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMachwf1rFieldTsoenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldTsoenMask)
	}
}

const (
	RegisterMachwf1rFieldDbgmemaShift = 19
	RegisterMachwf1rFieldDbgmemaMask  = 0x80000
)

// GetDbgmema DBGMEMA
func (r *registerMachwf1rType) GetDbgmema() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldDbgmemaMask) != 0
}

// SetDbgmema DBGMEMA
func (r *registerMachwf1rType) SetDbgmema(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMachwf1rFieldDbgmemaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldDbgmemaMask)
	}
}

const (
	RegisterMachwf1rFieldAvselShift = 20
	RegisterMachwf1rFieldAvselMask  = 0x100000
)

// GetAvsel AVSEL
func (r *registerMachwf1rType) GetAvsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldAvselMask) != 0
}

// SetAvsel AVSEL
func (r *registerMachwf1rType) SetAvsel(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMachwf1rFieldAvselMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldAvselMask)
	}
}

const (
	RegisterMachwf1rFieldHashtblszShift = 24
	RegisterMachwf1rFieldHashtblszMask  = 0x3000000
)

// GetHashtblsz HASHTBLSZ
func (r *registerMachwf1rType) GetHashtblsz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldHashtblszMask) >> RegisterMachwf1rFieldHashtblszShift)
}

// SetHashtblsz HASHTBLSZ
func (r *registerMachwf1rType) SetHashtblsz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldHashtblszMask)|(uint32(value)<<RegisterMachwf1rFieldHashtblszShift))
}

const (
	RegisterMachwf1rFieldL3l4fnumShift = 27
	RegisterMachwf1rFieldL3l4fnumMask  = 0x78000000
)

// GetL3l4fnum L3L4FNUM
func (r *registerMachwf1rType) GetL3l4fnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldL3l4fnumMask) >> RegisterMachwf1rFieldL3l4fnumShift)
}

// SetL3l4fnum L3L4FNUM
func (r *registerMachwf1rType) SetL3l4fnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldL3l4fnumMask)|(uint32(value)<<RegisterMachwf1rFieldL3l4fnumShift))
}

// registerMachwf2rType HW feature 2 register
type registerMachwf2rType uint32

const (
	RegisterMachwf2rFieldRxqcntShift = 0
	RegisterMachwf2rFieldRxqcntMask  = 0xf
)

// GetRxqcnt RXQCNT
func (r *registerMachwf2rType) GetRxqcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldRxqcntMask) >> RegisterMachwf2rFieldRxqcntShift)
}

// SetRxqcnt RXQCNT
func (r *registerMachwf2rType) SetRxqcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldRxqcntMask)|(uint32(value)<<RegisterMachwf2rFieldRxqcntShift))
}

const (
	RegisterMachwf2rFieldTxqcntShift = 6
	RegisterMachwf2rFieldTxqcntMask  = 0x3c0
)

// GetTxqcnt TXQCNT
func (r *registerMachwf2rType) GetTxqcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldTxqcntMask) >> RegisterMachwf2rFieldTxqcntShift)
}

// SetTxqcnt TXQCNT
func (r *registerMachwf2rType) SetTxqcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldTxqcntMask)|(uint32(value)<<RegisterMachwf2rFieldTxqcntShift))
}

const (
	RegisterMachwf2rFieldRxchcntShift = 12
	RegisterMachwf2rFieldRxchcntMask  = 0xf000
)

// GetRxchcnt RXCHCNT
func (r *registerMachwf2rType) GetRxchcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldRxchcntMask) >> RegisterMachwf2rFieldRxchcntShift)
}

// SetRxchcnt RXCHCNT
func (r *registerMachwf2rType) SetRxchcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldRxchcntMask)|(uint32(value)<<RegisterMachwf2rFieldRxchcntShift))
}

const (
	RegisterMachwf2rFieldTxchcntShift = 18
	RegisterMachwf2rFieldTxchcntMask  = 0x3c0000
)

// GetTxchcnt TXCHCNT
func (r *registerMachwf2rType) GetTxchcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldTxchcntMask) >> RegisterMachwf2rFieldTxchcntShift)
}

// SetTxchcnt TXCHCNT
func (r *registerMachwf2rType) SetTxchcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldTxchcntMask)|(uint32(value)<<RegisterMachwf2rFieldTxchcntShift))
}

const (
	RegisterMachwf2rFieldPpsoutnumShift = 24
	RegisterMachwf2rFieldPpsoutnumMask  = 0x7000000
)

// GetPpsoutnum PPSOUTNUM
func (r *registerMachwf2rType) GetPpsoutnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldPpsoutnumMask) >> RegisterMachwf2rFieldPpsoutnumShift)
}

// SetPpsoutnum PPSOUTNUM
func (r *registerMachwf2rType) SetPpsoutnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldPpsoutnumMask)|(uint32(value)<<RegisterMachwf2rFieldPpsoutnumShift))
}

const (
	RegisterMachwf2rFieldAuxsnapnumShift = 28
	RegisterMachwf2rFieldAuxsnapnumMask  = 0x70000000
)

// GetAuxsnapnum AUXSNAPNUM
func (r *registerMachwf2rType) GetAuxsnapnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldAuxsnapnumMask) >> RegisterMachwf2rFieldAuxsnapnumShift)
}

// SetAuxsnapnum AUXSNAPNUM
func (r *registerMachwf2rType) SetAuxsnapnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldAuxsnapnumMask)|(uint32(value)<<RegisterMachwf2rFieldAuxsnapnumShift))
}

// registerMacmdioarType MDIO address register
type registerMacmdioarType uint32

const (
	RegisterMacmdioarFieldMbShift = 0
	RegisterMacmdioarFieldMbMask  = 0x1
)

// GetMb MB
func (r *registerMacmdioarType) GetMb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldMbMask) != 0
}

// SetMb MB
func (r *registerMacmdioarType) SetMb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacmdioarFieldMbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldMbMask)
	}
}

const (
	RegisterMacmdioarFieldC45eShift = 1
	RegisterMacmdioarFieldC45eMask  = 0x2
)

// GetC45e C45E
func (r *registerMacmdioarType) GetC45e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldC45eMask) != 0
}

// SetC45e C45E
func (r *registerMacmdioarType) SetC45e(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacmdioarFieldC45eMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldC45eMask)
	}
}

const (
	RegisterMacmdioarFieldGocShift = 2
	RegisterMacmdioarFieldGocMask  = 0xc
)

// GetGoc GOC
func (r *registerMacmdioarType) GetGoc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldGocMask) >> RegisterMacmdioarFieldGocShift)
}

// SetGoc GOC
func (r *registerMacmdioarType) SetGoc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldGocMask)|(uint32(value)<<RegisterMacmdioarFieldGocShift))
}

const (
	RegisterMacmdioarFieldSkapShift = 4
	RegisterMacmdioarFieldSkapMask  = 0x10
)

// GetSkap SKAP
func (r *registerMacmdioarType) GetSkap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldSkapMask) != 0
}

// SetSkap SKAP
func (r *registerMacmdioarType) SetSkap(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacmdioarFieldSkapMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldSkapMask)
	}
}

const (
	RegisterMacmdioarFieldCrShift = 8
	RegisterMacmdioarFieldCrMask  = 0xf00
)

// GetCr CR
func (r *registerMacmdioarType) GetCr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldCrMask) >> RegisterMacmdioarFieldCrShift)
}

// SetCr CR
func (r *registerMacmdioarType) SetCr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldCrMask)|(uint32(value)<<RegisterMacmdioarFieldCrShift))
}

const (
	RegisterMacmdioarFieldNtcShift = 12
	RegisterMacmdioarFieldNtcMask  = 0x7000
)

// GetNtc NTC
func (r *registerMacmdioarType) GetNtc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldNtcMask) >> RegisterMacmdioarFieldNtcShift)
}

// SetNtc NTC
func (r *registerMacmdioarType) SetNtc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldNtcMask)|(uint32(value)<<RegisterMacmdioarFieldNtcShift))
}

const (
	RegisterMacmdioarFieldRdaShift = 16
	RegisterMacmdioarFieldRdaMask  = 0x1f0000
)

// GetRda RDA
func (r *registerMacmdioarType) GetRda() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldRdaMask) >> RegisterMacmdioarFieldRdaShift)
}

// SetRda RDA
func (r *registerMacmdioarType) SetRda(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldRdaMask)|(uint32(value)<<RegisterMacmdioarFieldRdaShift))
}

const (
	RegisterMacmdioarFieldPaShift = 21
	RegisterMacmdioarFieldPaMask  = 0x3e00000
)

// GetPa PA
func (r *registerMacmdioarType) GetPa() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldPaMask) >> RegisterMacmdioarFieldPaShift)
}

// SetPa PA
func (r *registerMacmdioarType) SetPa(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldPaMask)|(uint32(value)<<RegisterMacmdioarFieldPaShift))
}

const (
	RegisterMacmdioarFieldBtbShift = 26
	RegisterMacmdioarFieldBtbMask  = 0x4000000
)

// GetBtb BTB
func (r *registerMacmdioarType) GetBtb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldBtbMask) != 0
}

// SetBtb BTB
func (r *registerMacmdioarType) SetBtb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacmdioarFieldBtbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldBtbMask)
	}
}

const (
	RegisterMacmdioarFieldPseShift = 27
	RegisterMacmdioarFieldPseMask  = 0x8000000
)

// GetPse PSE
func (r *registerMacmdioarType) GetPse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldPseMask) != 0
}

// SetPse PSE
func (r *registerMacmdioarType) SetPse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacmdioarFieldPseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldPseMask)
	}
}

// registerMacmdiodrType MDIO data register
type registerMacmdiodrType uint32

const (
	RegisterMacmdiodrFieldMdShift = 0
	RegisterMacmdiodrFieldMdMask  = 0xffff
)

// GetMd MD
func (r *registerMacmdiodrType) GetMd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdiodrFieldMdMask) >> RegisterMacmdiodrFieldMdShift)
}

// SetMd MD
func (r *registerMacmdiodrType) SetMd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdiodrFieldMdMask)|(uint32(value)<<RegisterMacmdiodrFieldMdShift))
}

const (
	RegisterMacmdiodrFieldRaShift = 16
	RegisterMacmdiodrFieldRaMask  = 0xffff0000
)

// GetRa RA
func (r *registerMacmdiodrType) GetRa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdiodrFieldRaMask) >> RegisterMacmdiodrFieldRaShift)
}

// SetRa RA
func (r *registerMacmdiodrType) SetRa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdiodrFieldRaMask)|(uint32(value)<<RegisterMacmdiodrFieldRaShift))
}

// registerMaca0hrType Address 0 high register
type registerMaca0hrType uint32

const (
	RegisterMaca0hrFieldAddrhiShift = 0
	RegisterMaca0hrFieldAddrhiMask  = 0xffff
)

// GetAddrhi ADDRHI
func (r *registerMaca0hrType) GetAddrhi() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMaca0hrFieldAddrhiMask) >> RegisterMaca0hrFieldAddrhiShift)
}

// SetAddrhi ADDRHI
func (r *registerMaca0hrType) SetAddrhi(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca0hrFieldAddrhiMask)|(uint32(value)<<RegisterMaca0hrFieldAddrhiShift))
}

const (
	RegisterMaca0hrFieldAeShift = 31
	RegisterMaca0hrFieldAeMask  = 0x80000000
)

// GetAe AE
func (r *registerMaca0hrType) GetAe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca0hrFieldAeMask) != 0
}

// registerMaca0lrType Address 0 low register
type registerMaca0lrType uint32

const (
	RegisterMaca0lrFieldAddrloShift = 0
	RegisterMaca0lrFieldAddrloMask  = 0xffffffff
)

// GetAddrlo ADDRLO
func (r *registerMaca0lrType) GetAddrlo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMaca0lrFieldAddrloMask) >> RegisterMaca0lrFieldAddrloShift)
}

// SetAddrlo ADDRLO
func (r *registerMaca0lrType) SetAddrlo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca0lrFieldAddrloMask)|(uint32(value)<<RegisterMaca0lrFieldAddrloShift))
}

// registerMaca1hrType Address 1 high register
type registerMaca1hrType uint32

const (
	RegisterMaca1hrFieldAddrhiShift = 0
	RegisterMaca1hrFieldAddrhiMask  = 0xffff
)

// GetAddrhi ADDRHI
func (r *registerMaca1hrType) GetAddrhi() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMaca1hrFieldAddrhiMask) >> RegisterMaca1hrFieldAddrhiShift)
}

// SetAddrhi ADDRHI
func (r *registerMaca1hrType) SetAddrhi(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca1hrFieldAddrhiMask)|(uint32(value)<<RegisterMaca1hrFieldAddrhiShift))
}

const (
	RegisterMaca1hrFieldMbcShift = 24
	RegisterMaca1hrFieldMbcMask  = 0x3f000000
)

// GetMbc MBC
func (r *registerMaca1hrType) GetMbc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaca1hrFieldMbcMask) >> RegisterMaca1hrFieldMbcShift)
}

// SetMbc MBC
func (r *registerMaca1hrType) SetMbc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca1hrFieldMbcMask)|(uint32(value)<<RegisterMaca1hrFieldMbcShift))
}

const (
	RegisterMaca1hrFieldSaShift = 30
	RegisterMaca1hrFieldSaMask  = 0x40000000
)

// GetSa SA
func (r *registerMaca1hrType) GetSa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca1hrFieldSaMask) != 0
}

// SetSa SA
func (r *registerMaca1hrType) SetSa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaca1hrFieldSaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaca1hrFieldSaMask)
	}
}

const (
	RegisterMaca1hrFieldAeShift = 31
	RegisterMaca1hrFieldAeMask  = 0x80000000
)

// GetAe AE
func (r *registerMaca1hrType) GetAe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca1hrFieldAeMask) != 0
}

// SetAe AE
func (r *registerMaca1hrType) SetAe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaca1hrFieldAeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaca1hrFieldAeMask)
	}
}

// registerMaca1lrType Address 1 low register
type registerMaca1lrType uint32

const (
	RegisterMaca1lrFieldAddrloShift = 0
	RegisterMaca1lrFieldAddrloMask  = 0xffffffff
)

// GetAddrlo ADDRLO
func (r *registerMaca1lrType) GetAddrlo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMaca1lrFieldAddrloMask) >> RegisterMaca1lrFieldAddrloShift)
}

// SetAddrlo ADDRLO
func (r *registerMaca1lrType) SetAddrlo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca1lrFieldAddrloMask)|(uint32(value)<<RegisterMaca1lrFieldAddrloShift))
}

// registerMaca2hrType Address 2 high register
type registerMaca2hrType uint32

const (
	RegisterMaca2hrFieldAddrhiShift = 0
	RegisterMaca2hrFieldAddrhiMask  = 0xffff
)

// GetAddrhi ADDRHI
func (r *registerMaca2hrType) GetAddrhi() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMaca2hrFieldAddrhiMask) >> RegisterMaca2hrFieldAddrhiShift)
}

// SetAddrhi ADDRHI
func (r *registerMaca2hrType) SetAddrhi(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca2hrFieldAddrhiMask)|(uint32(value)<<RegisterMaca2hrFieldAddrhiShift))
}

const (
	RegisterMaca2hrFieldMbcShift = 24
	RegisterMaca2hrFieldMbcMask  = 0x3f000000
)

// GetMbc MBC
func (r *registerMaca2hrType) GetMbc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaca2hrFieldMbcMask) >> RegisterMaca2hrFieldMbcShift)
}

// SetMbc MBC
func (r *registerMaca2hrType) SetMbc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca2hrFieldMbcMask)|(uint32(value)<<RegisterMaca2hrFieldMbcShift))
}

const (
	RegisterMaca2hrFieldSaShift = 30
	RegisterMaca2hrFieldSaMask  = 0x40000000
)

// GetSa SA
func (r *registerMaca2hrType) GetSa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca2hrFieldSaMask) != 0
}

// SetSa SA
func (r *registerMaca2hrType) SetSa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaca2hrFieldSaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaca2hrFieldSaMask)
	}
}

const (
	RegisterMaca2hrFieldAeShift = 31
	RegisterMaca2hrFieldAeMask  = 0x80000000
)

// GetAe AE
func (r *registerMaca2hrType) GetAe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca2hrFieldAeMask) != 0
}

// SetAe AE
func (r *registerMaca2hrType) SetAe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaca2hrFieldAeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaca2hrFieldAeMask)
	}
}

// registerMaca2lrType Address 2 low register
type registerMaca2lrType uint32

const (
	RegisterMaca2lrFieldAddrloShift = 0
	RegisterMaca2lrFieldAddrloMask  = 0xffffffff
)

// GetAddrlo ADDRLO
func (r *registerMaca2lrType) GetAddrlo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMaca2lrFieldAddrloMask) >> RegisterMaca2lrFieldAddrloShift)
}

// SetAddrlo ADDRLO
func (r *registerMaca2lrType) SetAddrlo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca2lrFieldAddrloMask)|(uint32(value)<<RegisterMaca2lrFieldAddrloShift))
}

// registerMaca3hrType Address 3 high register
type registerMaca3hrType uint32

const (
	RegisterMaca3hrFieldAddrhiShift = 0
	RegisterMaca3hrFieldAddrhiMask  = 0xffff
)

// GetAddrhi ADDRHI
func (r *registerMaca3hrType) GetAddrhi() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMaca3hrFieldAddrhiMask) >> RegisterMaca3hrFieldAddrhiShift)
}

// SetAddrhi ADDRHI
func (r *registerMaca3hrType) SetAddrhi(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca3hrFieldAddrhiMask)|(uint32(value)<<RegisterMaca3hrFieldAddrhiShift))
}

const (
	RegisterMaca3hrFieldMbcShift = 24
	RegisterMaca3hrFieldMbcMask  = 0x3f000000
)

// GetMbc MBC
func (r *registerMaca3hrType) GetMbc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaca3hrFieldMbcMask) >> RegisterMaca3hrFieldMbcShift)
}

// SetMbc MBC
func (r *registerMaca3hrType) SetMbc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca3hrFieldMbcMask)|(uint32(value)<<RegisterMaca3hrFieldMbcShift))
}

const (
	RegisterMaca3hrFieldSaShift = 30
	RegisterMaca3hrFieldSaMask  = 0x40000000
)

// GetSa SA
func (r *registerMaca3hrType) GetSa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca3hrFieldSaMask) != 0
}

// SetSa SA
func (r *registerMaca3hrType) SetSa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaca3hrFieldSaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaca3hrFieldSaMask)
	}
}

const (
	RegisterMaca3hrFieldAeShift = 31
	RegisterMaca3hrFieldAeMask  = 0x80000000
)

// GetAe AE
func (r *registerMaca3hrType) GetAe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca3hrFieldAeMask) != 0
}

// SetAe AE
func (r *registerMaca3hrType) SetAe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaca3hrFieldAeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaca3hrFieldAeMask)
	}
}

// registerMaca3lrType Address 3 low register
type registerMaca3lrType uint32

const (
	RegisterMaca3lrFieldAddrloShift = 0
	RegisterMaca3lrFieldAddrloMask  = 0xffffffff
)

// GetAddrlo ADDRLO
func (r *registerMaca3lrType) GetAddrlo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMaca3lrFieldAddrloMask) >> RegisterMaca3lrFieldAddrloShift)
}

// SetAddrlo ADDRLO
func (r *registerMaca3lrType) SetAddrlo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca3lrFieldAddrloMask)|(uint32(value)<<RegisterMaca3lrFieldAddrloShift))
}

// registerMmccontrolType MMC control register
type registerMmccontrolType uint32

const (
	RegisterMmccontrolFieldCntrstShift = 0
	RegisterMmccontrolFieldCntrstMask  = 0x1
)

// GetCntrst CNTRST
func (r *registerMmccontrolType) GetCntrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldCntrstMask) != 0
}

// SetCntrst CNTRST
func (r *registerMmccontrolType) SetCntrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmccontrolFieldCntrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmccontrolFieldCntrstMask)
	}
}

const (
	RegisterMmccontrolFieldCntstoproShift = 1
	RegisterMmccontrolFieldCntstoproMask  = 0x2
)

// GetCntstopro CNTSTOPRO
func (r *registerMmccontrolType) GetCntstopro() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldCntstoproMask) != 0
}

// SetCntstopro CNTSTOPRO
func (r *registerMmccontrolType) SetCntstopro(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmccontrolFieldCntstoproMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmccontrolFieldCntstoproMask)
	}
}

const (
	RegisterMmccontrolFieldRstonrdShift = 2
	RegisterMmccontrolFieldRstonrdMask  = 0x4
)

// GetRstonrd RSTONRD
func (r *registerMmccontrolType) GetRstonrd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldRstonrdMask) != 0
}

// SetRstonrd RSTONRD
func (r *registerMmccontrolType) SetRstonrd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmccontrolFieldRstonrdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmccontrolFieldRstonrdMask)
	}
}

const (
	RegisterMmccontrolFieldCntfreezShift = 3
	RegisterMmccontrolFieldCntfreezMask  = 0x8
)

// GetCntfreez CNTFREEZ
func (r *registerMmccontrolType) GetCntfreez() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldCntfreezMask) != 0
}

// SetCntfreez CNTFREEZ
func (r *registerMmccontrolType) SetCntfreez(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmccontrolFieldCntfreezMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmccontrolFieldCntfreezMask)
	}
}

const (
	RegisterMmccontrolFieldCntprstShift = 4
	RegisterMmccontrolFieldCntprstMask  = 0x10
)

// GetCntprst CNTPRST
func (r *registerMmccontrolType) GetCntprst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldCntprstMask) != 0
}

// SetCntprst CNTPRST
func (r *registerMmccontrolType) SetCntprst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmccontrolFieldCntprstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmccontrolFieldCntprstMask)
	}
}

const (
	RegisterMmccontrolFieldCntprstlvlShift = 5
	RegisterMmccontrolFieldCntprstlvlMask  = 0x20
)

// GetCntprstlvl CNTPRSTLVL
func (r *registerMmccontrolType) GetCntprstlvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldCntprstlvlMask) != 0
}

// SetCntprstlvl CNTPRSTLVL
func (r *registerMmccontrolType) SetCntprstlvl(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmccontrolFieldCntprstlvlMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmccontrolFieldCntprstlvlMask)
	}
}

const (
	RegisterMmccontrolFieldUcdbcShift = 8
	RegisterMmccontrolFieldUcdbcMask  = 0x100
)

// GetUcdbc UCDBC
func (r *registerMmccontrolType) GetUcdbc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldUcdbcMask) != 0
}

// SetUcdbc UCDBC
func (r *registerMmccontrolType) SetUcdbc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmccontrolFieldUcdbcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmccontrolFieldUcdbcMask)
	}
}

// registerMmcrxinterruptType MMC Rx interrupt register
type registerMmcrxinterruptType uint32

const (
	RegisterMmcrxinterruptFieldRxcrcerpisShift = 5
	RegisterMmcrxinterruptFieldRxcrcerpisMask  = 0x20
)

// GetRxcrcerpis RXCRCERPIS
func (r *registerMmcrxinterruptType) GetRxcrcerpis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptFieldRxcrcerpisMask) != 0
}

// SetRxcrcerpis RXCRCERPIS
func (r *registerMmcrxinterruptType) SetRxcrcerpis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmcrxinterruptFieldRxcrcerpisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmcrxinterruptFieldRxcrcerpisMask)
	}
}

const (
	RegisterMmcrxinterruptFieldRxalgnerpisShift = 6
	RegisterMmcrxinterruptFieldRxalgnerpisMask  = 0x40
)

// GetRxalgnerpis RXALGNERPIS
func (r *registerMmcrxinterruptType) GetRxalgnerpis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptFieldRxalgnerpisMask) != 0
}

// SetRxalgnerpis RXALGNERPIS
func (r *registerMmcrxinterruptType) SetRxalgnerpis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmcrxinterruptFieldRxalgnerpisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmcrxinterruptFieldRxalgnerpisMask)
	}
}

const (
	RegisterMmcrxinterruptFieldRxucgpisShift = 17
	RegisterMmcrxinterruptFieldRxucgpisMask  = 0x20000
)

// GetRxucgpis RXUCGPIS
func (r *registerMmcrxinterruptType) GetRxucgpis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptFieldRxucgpisMask) != 0
}

// SetRxucgpis RXUCGPIS
func (r *registerMmcrxinterruptType) SetRxucgpis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmcrxinterruptFieldRxucgpisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmcrxinterruptFieldRxucgpisMask)
	}
}

const (
	RegisterMmcrxinterruptFieldRxlpiuscisShift = 26
	RegisterMmcrxinterruptFieldRxlpiuscisMask  = 0x4000000
)

// GetRxlpiuscis RXLPIUSCIS
func (r *registerMmcrxinterruptType) GetRxlpiuscis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptFieldRxlpiuscisMask) != 0
}

// SetRxlpiuscis RXLPIUSCIS
func (r *registerMmcrxinterruptType) SetRxlpiuscis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmcrxinterruptFieldRxlpiuscisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmcrxinterruptFieldRxlpiuscisMask)
	}
}

const (
	RegisterMmcrxinterruptFieldRxlpitrcisShift = 27
	RegisterMmcrxinterruptFieldRxlpitrcisMask  = 0x8000000
)

// GetRxlpitrcis RXLPITRCIS
func (r *registerMmcrxinterruptType) GetRxlpitrcis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptFieldRxlpitrcisMask) != 0
}

// SetRxlpitrcis RXLPITRCIS
func (r *registerMmcrxinterruptType) SetRxlpitrcis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmcrxinterruptFieldRxlpitrcisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmcrxinterruptFieldRxlpitrcisMask)
	}
}

// registerMmctxinterruptType MMC Tx interrupt register
type registerMmctxinterruptType uint32

const (
	RegisterMmctxinterruptFieldTxscolgpisShift = 14
	RegisterMmctxinterruptFieldTxscolgpisMask  = 0x4000
)

// GetTxscolgpis TXSCOLGPIS
func (r *registerMmctxinterruptType) GetTxscolgpis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptFieldTxscolgpisMask) != 0
}

// SetTxscolgpis TXSCOLGPIS
func (r *registerMmctxinterruptType) SetTxscolgpis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmctxinterruptFieldTxscolgpisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmctxinterruptFieldTxscolgpisMask)
	}
}

const (
	RegisterMmctxinterruptFieldTxmcolgpisShift = 15
	RegisterMmctxinterruptFieldTxmcolgpisMask  = 0x8000
)

// GetTxmcolgpis TXMCOLGPIS
func (r *registerMmctxinterruptType) GetTxmcolgpis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptFieldTxmcolgpisMask) != 0
}

// SetTxmcolgpis TXMCOLGPIS
func (r *registerMmctxinterruptType) SetTxmcolgpis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmctxinterruptFieldTxmcolgpisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmctxinterruptFieldTxmcolgpisMask)
	}
}

const (
	RegisterMmctxinterruptFieldTxgpktisShift = 21
	RegisterMmctxinterruptFieldTxgpktisMask  = 0x200000
)

// GetTxgpktis TXGPKTIS
func (r *registerMmctxinterruptType) GetTxgpktis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptFieldTxgpktisMask) != 0
}

// SetTxgpktis TXGPKTIS
func (r *registerMmctxinterruptType) SetTxgpktis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmctxinterruptFieldTxgpktisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmctxinterruptFieldTxgpktisMask)
	}
}

const (
	RegisterMmctxinterruptFieldTxlpiuscisShift = 26
	RegisterMmctxinterruptFieldTxlpiuscisMask  = 0x4000000
)

// GetTxlpiuscis TXLPIUSCIS
func (r *registerMmctxinterruptType) GetTxlpiuscis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptFieldTxlpiuscisMask) != 0
}

// SetTxlpiuscis TXLPIUSCIS
func (r *registerMmctxinterruptType) SetTxlpiuscis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmctxinterruptFieldTxlpiuscisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmctxinterruptFieldTxlpiuscisMask)
	}
}

const (
	RegisterMmctxinterruptFieldTxlpitrcisShift = 27
	RegisterMmctxinterruptFieldTxlpitrcisMask  = 0x8000000
)

// GetTxlpitrcis TXLPITRCIS
func (r *registerMmctxinterruptType) GetTxlpitrcis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptFieldTxlpitrcisMask) != 0
}

// SetTxlpitrcis TXLPITRCIS
func (r *registerMmctxinterruptType) SetTxlpitrcis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmctxinterruptFieldTxlpitrcisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmctxinterruptFieldTxlpitrcisMask)
	}
}

// registerMmcrxinterruptmaskType MMC Rx interrupt mask register
type registerMmcrxinterruptmaskType uint32

const (
	RegisterMmcrxinterruptmaskFieldRxcrcerpimShift = 5
	RegisterMmcrxinterruptmaskFieldRxcrcerpimMask  = 0x20
)

// GetRxcrcerpim RXCRCERPIM
func (r *registerMmcrxinterruptmaskType) GetRxcrcerpim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptmaskFieldRxcrcerpimMask) != 0
}

// SetRxcrcerpim RXCRCERPIM
func (r *registerMmcrxinterruptmaskType) SetRxcrcerpim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmcrxinterruptmaskFieldRxcrcerpimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmcrxinterruptmaskFieldRxcrcerpimMask)
	}
}

const (
	RegisterMmcrxinterruptmaskFieldRxalgnerpimShift = 6
	RegisterMmcrxinterruptmaskFieldRxalgnerpimMask  = 0x40
)

// GetRxalgnerpim RXALGNERPIM
func (r *registerMmcrxinterruptmaskType) GetRxalgnerpim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptmaskFieldRxalgnerpimMask) != 0
}

// SetRxalgnerpim RXALGNERPIM
func (r *registerMmcrxinterruptmaskType) SetRxalgnerpim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmcrxinterruptmaskFieldRxalgnerpimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmcrxinterruptmaskFieldRxalgnerpimMask)
	}
}

const (
	RegisterMmcrxinterruptmaskFieldRxucgpimShift = 17
	RegisterMmcrxinterruptmaskFieldRxucgpimMask  = 0x20000
)

// GetRxucgpim RXUCGPIM
func (r *registerMmcrxinterruptmaskType) GetRxucgpim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptmaskFieldRxucgpimMask) != 0
}

// SetRxucgpim RXUCGPIM
func (r *registerMmcrxinterruptmaskType) SetRxucgpim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmcrxinterruptmaskFieldRxucgpimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmcrxinterruptmaskFieldRxucgpimMask)
	}
}

const (
	RegisterMmcrxinterruptmaskFieldRxlpiuscimShift = 26
	RegisterMmcrxinterruptmaskFieldRxlpiuscimMask  = 0x4000000
)

// GetRxlpiuscim RXLPIUSCIM
func (r *registerMmcrxinterruptmaskType) GetRxlpiuscim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptmaskFieldRxlpiuscimMask) != 0
}

// SetRxlpiuscim RXLPIUSCIM
func (r *registerMmcrxinterruptmaskType) SetRxlpiuscim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmcrxinterruptmaskFieldRxlpiuscimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmcrxinterruptmaskFieldRxlpiuscimMask)
	}
}

const (
	RegisterMmcrxinterruptmaskFieldRxlpitrcimShift = 27
	RegisterMmcrxinterruptmaskFieldRxlpitrcimMask  = 0x8000000
)

// GetRxlpitrcim RXLPITRCIM
func (r *registerMmcrxinterruptmaskType) GetRxlpitrcim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptmaskFieldRxlpitrcimMask) != 0
}

// registerMmctxinterruptmaskType MMC Tx interrupt mask register
type registerMmctxinterruptmaskType uint32

const (
	RegisterMmctxinterruptmaskFieldTxscolgpimShift = 14
	RegisterMmctxinterruptmaskFieldTxscolgpimMask  = 0x4000
)

// GetTxscolgpim TXSCOLGPIM
func (r *registerMmctxinterruptmaskType) GetTxscolgpim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptmaskFieldTxscolgpimMask) != 0
}

// SetTxscolgpim TXSCOLGPIM
func (r *registerMmctxinterruptmaskType) SetTxscolgpim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmctxinterruptmaskFieldTxscolgpimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmctxinterruptmaskFieldTxscolgpimMask)
	}
}

const (
	RegisterMmctxinterruptmaskFieldTxmcolgpimShift = 15
	RegisterMmctxinterruptmaskFieldTxmcolgpimMask  = 0x8000
)

// GetTxmcolgpim TXMCOLGPIM
func (r *registerMmctxinterruptmaskType) GetTxmcolgpim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptmaskFieldTxmcolgpimMask) != 0
}

// SetTxmcolgpim TXMCOLGPIM
func (r *registerMmctxinterruptmaskType) SetTxmcolgpim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmctxinterruptmaskFieldTxmcolgpimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmctxinterruptmaskFieldTxmcolgpimMask)
	}
}

const (
	RegisterMmctxinterruptmaskFieldTxgpktimShift = 21
	RegisterMmctxinterruptmaskFieldTxgpktimMask  = 0x200000
)

// GetTxgpktim TXGPKTIM
func (r *registerMmctxinterruptmaskType) GetTxgpktim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptmaskFieldTxgpktimMask) != 0
}

// SetTxgpktim TXGPKTIM
func (r *registerMmctxinterruptmaskType) SetTxgpktim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmctxinterruptmaskFieldTxgpktimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmctxinterruptmaskFieldTxgpktimMask)
	}
}

const (
	RegisterMmctxinterruptmaskFieldTxlpiuscimShift = 26
	RegisterMmctxinterruptmaskFieldTxlpiuscimMask  = 0x4000000
)

// GetTxlpiuscim TXLPIUSCIM
func (r *registerMmctxinterruptmaskType) GetTxlpiuscim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptmaskFieldTxlpiuscimMask) != 0
}

// SetTxlpiuscim TXLPIUSCIM
func (r *registerMmctxinterruptmaskType) SetTxlpiuscim(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmctxinterruptmaskFieldTxlpiuscimMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmctxinterruptmaskFieldTxlpiuscimMask)
	}
}

const (
	RegisterMmctxinterruptmaskFieldTxlpitrcimShift = 27
	RegisterMmctxinterruptmaskFieldTxlpitrcimMask  = 0x8000000
)

// GetTxlpitrcim TXLPITRCIM
func (r *registerMmctxinterruptmaskType) GetTxlpitrcim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptmaskFieldTxlpitrcimMask) != 0
}

// registerTxsinglecollisiongoodpacketsType Tx single collision good packets register
type registerTxsinglecollisiongoodpacketsType uint32

const (
	RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgShift = 0
	RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgMask  = 0xffffffff
)

// GetTxsnglcolg TXSNGLCOLG
func (r *registerTxsinglecollisiongoodpacketsType) GetTxsnglcolg() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgMask) >> RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgShift)
}

// SetTxsnglcolg TXSNGLCOLG
func (r *registerTxsinglecollisiongoodpacketsType) SetTxsnglcolg(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgMask)|(uint32(value)<<RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgShift))
}

// registerTxmultiplecollisiongoodpacketsType Tx multiple collision good packets register
type registerTxmultiplecollisiongoodpacketsType uint32

const (
	RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgShift = 0
	RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgMask  = 0xffffffff
)

// GetTxmultcolg TXMULTCOLG
func (r *registerTxmultiplecollisiongoodpacketsType) GetTxmultcolg() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgMask) >> RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgShift)
}

// SetTxmultcolg TXMULTCOLG
func (r *registerTxmultiplecollisiongoodpacketsType) SetTxmultcolg(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgMask)|(uint32(value)<<RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgShift))
}

// registerTxpacketcountgoodType Tx packet count good register
type registerTxpacketcountgoodType uint32

const (
	RegisterTxpacketcountgoodFieldTxpktgShift = 0
	RegisterTxpacketcountgoodFieldTxpktgMask  = 0xffffffff
)

// GetTxpktg TXPKTG
func (r *registerTxpacketcountgoodType) GetTxpktg() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxpacketcountgoodFieldTxpktgMask) >> RegisterTxpacketcountgoodFieldTxpktgShift)
}

// SetTxpktg TXPKTG
func (r *registerTxpacketcountgoodType) SetTxpktg(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxpacketcountgoodFieldTxpktgMask)|(uint32(value)<<RegisterTxpacketcountgoodFieldTxpktgShift))
}

// registerRxcrcerrorpacketsType Rx CRC error packets register
type registerRxcrcerrorpacketsType uint32

const (
	RegisterRxcrcerrorpacketsFieldRxcrcerrShift = 0
	RegisterRxcrcerrorpacketsFieldRxcrcerrMask  = 0xffffffff
)

// GetRxcrcerr RXCRCERR
func (r *registerRxcrcerrorpacketsType) GetRxcrcerr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxcrcerrorpacketsFieldRxcrcerrMask) >> RegisterRxcrcerrorpacketsFieldRxcrcerrShift)
}

// SetRxcrcerr RXCRCERR
func (r *registerRxcrcerrorpacketsType) SetRxcrcerr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxcrcerrorpacketsFieldRxcrcerrMask)|(uint32(value)<<RegisterRxcrcerrorpacketsFieldRxcrcerrShift))
}

// registerRxalignmenterrorpacketsType Rx alignment error packets register
type registerRxalignmenterrorpacketsType uint32

const (
	RegisterRxalignmenterrorpacketsFieldRxalgnerrShift = 0
	RegisterRxalignmenterrorpacketsFieldRxalgnerrMask  = 0xffffffff
)

// GetRxalgnerr RXALGNERR
func (r *registerRxalignmenterrorpacketsType) GetRxalgnerr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxalignmenterrorpacketsFieldRxalgnerrMask) >> RegisterRxalignmenterrorpacketsFieldRxalgnerrShift)
}

// SetRxalgnerr RXALGNERR
func (r *registerRxalignmenterrorpacketsType) SetRxalgnerr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxalignmenterrorpacketsFieldRxalgnerrMask)|(uint32(value)<<RegisterRxalignmenterrorpacketsFieldRxalgnerrShift))
}

// registerRxunicastpacketsgoodType Rx unicast packets good register
type registerRxunicastpacketsgoodType uint32

const (
	RegisterRxunicastpacketsgoodFieldRxucastgShift = 0
	RegisterRxunicastpacketsgoodFieldRxucastgMask  = 0xffffffff
)

// GetRxucastg RXUCASTG
func (r *registerRxunicastpacketsgoodType) GetRxucastg() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxunicastpacketsgoodFieldRxucastgMask) >> RegisterRxunicastpacketsgoodFieldRxucastgShift)
}

// SetRxucastg RXUCASTG
func (r *registerRxunicastpacketsgoodType) SetRxucastg(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxunicastpacketsgoodFieldRxucastgMask)|(uint32(value)<<RegisterRxunicastpacketsgoodFieldRxucastgShift))
}

// registerTxlpiuseccntrType Tx LPI microsecond timer register
type registerTxlpiuseccntrType uint32

const (
	RegisterTxlpiuseccntrFieldTxlpiuscShift = 0
	RegisterTxlpiuseccntrFieldTxlpiuscMask  = 0xffffffff
)

// GetTxlpiusc TXLPIUSC
func (r *registerTxlpiuseccntrType) GetTxlpiusc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxlpiuseccntrFieldTxlpiuscMask) >> RegisterTxlpiuseccntrFieldTxlpiuscShift)
}

// SetTxlpiusc TXLPIUSC
func (r *registerTxlpiuseccntrType) SetTxlpiusc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxlpiuseccntrFieldTxlpiuscMask)|(uint32(value)<<RegisterTxlpiuseccntrFieldTxlpiuscShift))
}

// registerTxlpitrancntrType Tx LPI transition counter register
type registerTxlpitrancntrType uint32

const (
	RegisterTxlpitrancntrFieldTxlpitrcShift = 0
	RegisterTxlpitrancntrFieldTxlpitrcMask  = 0xffffffff
)

// GetTxlpitrc TXLPITRC
func (r *registerTxlpitrancntrType) GetTxlpitrc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxlpitrancntrFieldTxlpitrcMask) >> RegisterTxlpitrancntrFieldTxlpitrcShift)
}

// SetTxlpitrc TXLPITRC
func (r *registerTxlpitrancntrType) SetTxlpitrc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxlpitrancntrFieldTxlpitrcMask)|(uint32(value)<<RegisterTxlpitrancntrFieldTxlpitrcShift))
}

// registerRxlpiuseccntrType Rx LPI microsecond counter register
type registerRxlpiuseccntrType uint32

const (
	RegisterRxlpiuseccntrFieldRxlpiuscShift = 0
	RegisterRxlpiuseccntrFieldRxlpiuscMask  = 0xffffffff
)

// GetRxlpiusc RXLPIUSC
func (r *registerRxlpiuseccntrType) GetRxlpiusc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxlpiuseccntrFieldRxlpiuscMask) >> RegisterRxlpiuseccntrFieldRxlpiuscShift)
}

// SetRxlpiusc RXLPIUSC
func (r *registerRxlpiuseccntrType) SetRxlpiusc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxlpiuseccntrFieldRxlpiuscMask)|(uint32(value)<<RegisterRxlpiuseccntrFieldRxlpiuscShift))
}

// registerRxlpitrancntrType Rx LPI transition counter register
type registerRxlpitrancntrType uint32

const (
	RegisterRxlpitrancntrFieldRxlpitrcShift = 0
	RegisterRxlpitrancntrFieldRxlpitrcMask  = 0xffffffff
)

// GetRxlpitrc RXLPITRC
func (r *registerRxlpitrancntrType) GetRxlpitrc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxlpitrancntrFieldRxlpitrcMask) >> RegisterRxlpitrancntrFieldRxlpitrcShift)
}

// SetRxlpitrc RXLPITRC
func (r *registerRxlpitrancntrType) SetRxlpitrc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxlpitrancntrFieldRxlpitrcMask)|(uint32(value)<<RegisterRxlpitrancntrFieldRxlpitrcShift))
}

// registerMacl3l4c0rType L3 and L4 control 0 register
type registerMacl3l4c0rType uint32

const (
	RegisterMacl3l4c0rFieldL3pen0Shift = 0
	RegisterMacl3l4c0rFieldL3pen0Mask  = 0x1
)

// GetL3pen0 L3PEN0
func (r *registerMacl3l4c0rType) GetL3pen0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3pen0Mask) != 0
}

// SetL3pen0 L3PEN0
func (r *registerMacl3l4c0rType) SetL3pen0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c0rFieldL3pen0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL3pen0Mask)
	}
}

const (
	RegisterMacl3l4c0rFieldL3sam0Shift = 2
	RegisterMacl3l4c0rFieldL3sam0Mask  = 0x4
)

// GetL3sam0 L3SAM0
func (r *registerMacl3l4c0rType) GetL3sam0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3sam0Mask) != 0
}

// SetL3sam0 L3SAM0
func (r *registerMacl3l4c0rType) SetL3sam0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c0rFieldL3sam0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL3sam0Mask)
	}
}

const (
	RegisterMacl3l4c0rFieldL3saim0Shift = 3
	RegisterMacl3l4c0rFieldL3saim0Mask  = 0x8
)

// GetL3saim0 L3SAIM0
func (r *registerMacl3l4c0rType) GetL3saim0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3saim0Mask) != 0
}

// SetL3saim0 L3SAIM0
func (r *registerMacl3l4c0rType) SetL3saim0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c0rFieldL3saim0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL3saim0Mask)
	}
}

const (
	RegisterMacl3l4c0rFieldL3dam0Shift = 4
	RegisterMacl3l4c0rFieldL3dam0Mask  = 0x10
)

// GetL3dam0 L3DAM0
func (r *registerMacl3l4c0rType) GetL3dam0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3dam0Mask) != 0
}

// SetL3dam0 L3DAM0
func (r *registerMacl3l4c0rType) SetL3dam0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c0rFieldL3dam0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL3dam0Mask)
	}
}

const (
	RegisterMacl3l4c0rFieldL3daim0Shift = 5
	RegisterMacl3l4c0rFieldL3daim0Mask  = 0x20
)

// GetL3daim0 L3DAIM0
func (r *registerMacl3l4c0rType) GetL3daim0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3daim0Mask) != 0
}

// SetL3daim0 L3DAIM0
func (r *registerMacl3l4c0rType) SetL3daim0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c0rFieldL3daim0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL3daim0Mask)
	}
}

const (
	RegisterMacl3l4c0rFieldL3hsbm0Shift = 6
	RegisterMacl3l4c0rFieldL3hsbm0Mask  = 0x7c0
)

// GetL3hsbm0 L3HSBM0
func (r *registerMacl3l4c0rType) GetL3hsbm0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3hsbm0Mask) >> RegisterMacl3l4c0rFieldL3hsbm0Shift)
}

// SetL3hsbm0 L3HSBM0
func (r *registerMacl3l4c0rType) SetL3hsbm0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL3hsbm0Mask)|(uint32(value)<<RegisterMacl3l4c0rFieldL3hsbm0Shift))
}

const (
	RegisterMacl3l4c0rFieldL3hdbm0Shift = 11
	RegisterMacl3l4c0rFieldL3hdbm0Mask  = 0xf800
)

// GetL3hdbm0 L3HDBM0
func (r *registerMacl3l4c0rType) GetL3hdbm0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3hdbm0Mask) >> RegisterMacl3l4c0rFieldL3hdbm0Shift)
}

// SetL3hdbm0 L3HDBM0
func (r *registerMacl3l4c0rType) SetL3hdbm0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL3hdbm0Mask)|(uint32(value)<<RegisterMacl3l4c0rFieldL3hdbm0Shift))
}

const (
	RegisterMacl3l4c0rFieldL4pen0Shift = 16
	RegisterMacl3l4c0rFieldL4pen0Mask  = 0x10000
)

// GetL4pen0 L4PEN0
func (r *registerMacl3l4c0rType) GetL4pen0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL4pen0Mask) != 0
}

// SetL4pen0 L4PEN0
func (r *registerMacl3l4c0rType) SetL4pen0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c0rFieldL4pen0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL4pen0Mask)
	}
}

const (
	RegisterMacl3l4c0rFieldL4spm0Shift = 18
	RegisterMacl3l4c0rFieldL4spm0Mask  = 0x40000
)

// GetL4spm0 L4SPM0
func (r *registerMacl3l4c0rType) GetL4spm0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL4spm0Mask) != 0
}

// SetL4spm0 L4SPM0
func (r *registerMacl3l4c0rType) SetL4spm0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c0rFieldL4spm0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL4spm0Mask)
	}
}

const (
	RegisterMacl3l4c0rFieldL4spim0Shift = 19
	RegisterMacl3l4c0rFieldL4spim0Mask  = 0x80000
)

// GetL4spim0 L4SPIM0
func (r *registerMacl3l4c0rType) GetL4spim0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL4spim0Mask) != 0
}

// SetL4spim0 L4SPIM0
func (r *registerMacl3l4c0rType) SetL4spim0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c0rFieldL4spim0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL4spim0Mask)
	}
}

const (
	RegisterMacl3l4c0rFieldL4dpm0Shift = 20
	RegisterMacl3l4c0rFieldL4dpm0Mask  = 0x100000
)

// GetL4dpm0 L4DPM0
func (r *registerMacl3l4c0rType) GetL4dpm0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL4dpm0Mask) != 0
}

// SetL4dpm0 L4DPM0
func (r *registerMacl3l4c0rType) SetL4dpm0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c0rFieldL4dpm0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL4dpm0Mask)
	}
}

const (
	RegisterMacl3l4c0rFieldL4dpim0Shift = 21
	RegisterMacl3l4c0rFieldL4dpim0Mask  = 0x200000
)

// GetL4dpim0 L4DPIM0
func (r *registerMacl3l4c0rType) GetL4dpim0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL4dpim0Mask) != 0
}

// SetL4dpim0 L4DPIM0
func (r *registerMacl3l4c0rType) SetL4dpim0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c0rFieldL4dpim0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL4dpim0Mask)
	}
}

// registerMacl4a0rType Layer4 address filter 0 register
type registerMacl4a0rType uint32

const (
	RegisterMacl4a0rFieldL4sp0Shift = 0
	RegisterMacl4a0rFieldL4sp0Mask  = 0xffff
)

// GetL4sp0 L4SP0
func (r *registerMacl4a0rType) GetL4sp0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacl4a0rFieldL4sp0Mask) >> RegisterMacl4a0rFieldL4sp0Shift)
}

// SetL4sp0 L4SP0
func (r *registerMacl4a0rType) SetL4sp0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl4a0rFieldL4sp0Mask)|(uint32(value)<<RegisterMacl4a0rFieldL4sp0Shift))
}

const (
	RegisterMacl4a0rFieldL4dp0Shift = 16
	RegisterMacl4a0rFieldL4dp0Mask  = 0xffff0000
)

// GetL4dp0 L4DP0
func (r *registerMacl4a0rType) GetL4dp0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacl4a0rFieldL4dp0Mask) >> RegisterMacl4a0rFieldL4dp0Shift)
}

// SetL4dp0 L4DP0
func (r *registerMacl4a0rType) SetL4dp0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl4a0rFieldL4dp0Mask)|(uint32(value)<<RegisterMacl4a0rFieldL4dp0Shift))
}

// registerMacl3a00rType MACL3A00R
type registerMacl3a00rType uint32

const (
	RegisterMacl3a00rFieldL3a00Shift = 0
	RegisterMacl3a00rFieldL3a00Mask  = 0xffffffff
)

// GetL3a00 L3A00
func (r *registerMacl3a00rType) GetL3a00() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a00rFieldL3a00Mask) >> RegisterMacl3a00rFieldL3a00Shift)
}

// SetL3a00 L3A00
func (r *registerMacl3a00rType) SetL3a00(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a00rFieldL3a00Mask)|(uint32(value)<<RegisterMacl3a00rFieldL3a00Shift))
}

// registerMacl3a10rType Layer3 address 1 filter 0 register
type registerMacl3a10rType uint32

const (
	RegisterMacl3a10rFieldL3a10Shift = 0
	RegisterMacl3a10rFieldL3a10Mask  = 0xffffffff
)

// GetL3a10 L3A10
func (r *registerMacl3a10rType) GetL3a10() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a10rFieldL3a10Mask) >> RegisterMacl3a10rFieldL3a10Shift)
}

// SetL3a10 L3A10
func (r *registerMacl3a10rType) SetL3a10(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a10rFieldL3a10Mask)|(uint32(value)<<RegisterMacl3a10rFieldL3a10Shift))
}

// registerMacl3a20Type Layer3 Address 2 filter 0 register
type registerMacl3a20Type uint32

const (
	RegisterMacl3a20FieldL3a20Shift = 0
	RegisterMacl3a20FieldL3a20Mask  = 0xffffffff
)

// GetL3a20 L3A20
func (r *registerMacl3a20Type) GetL3a20() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a20FieldL3a20Mask) >> RegisterMacl3a20FieldL3a20Shift)
}

// SetL3a20 L3A20
func (r *registerMacl3a20Type) SetL3a20(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a20FieldL3a20Mask)|(uint32(value)<<RegisterMacl3a20FieldL3a20Shift))
}

// registerMacl3a30Type Layer3 Address 3 filter 0 register
type registerMacl3a30Type uint32

const (
	RegisterMacl3a30FieldL3a30Shift = 0
	RegisterMacl3a30FieldL3a30Mask  = 0xffffffff
)

// GetL3a30 L3A30
func (r *registerMacl3a30Type) GetL3a30() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a30FieldL3a30Mask) >> RegisterMacl3a30FieldL3a30Shift)
}

// SetL3a30 L3A30
func (r *registerMacl3a30Type) SetL3a30(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a30FieldL3a30Mask)|(uint32(value)<<RegisterMacl3a30FieldL3a30Shift))
}

// registerMacl3l4c1rType L3 and L4 control 1 register
type registerMacl3l4c1rType uint32

const (
	RegisterMacl3l4c1rFieldL3pen1Shift = 0
	RegisterMacl3l4c1rFieldL3pen1Mask  = 0x1
)

// GetL3pen1 L3PEN1
func (r *registerMacl3l4c1rType) GetL3pen1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3pen1Mask) != 0
}

// SetL3pen1 L3PEN1
func (r *registerMacl3l4c1rType) SetL3pen1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c1rFieldL3pen1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL3pen1Mask)
	}
}

const (
	RegisterMacl3l4c1rFieldL3sam1Shift = 2
	RegisterMacl3l4c1rFieldL3sam1Mask  = 0x4
)

// GetL3sam1 L3SAM1
func (r *registerMacl3l4c1rType) GetL3sam1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3sam1Mask) != 0
}

// SetL3sam1 L3SAM1
func (r *registerMacl3l4c1rType) SetL3sam1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c1rFieldL3sam1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL3sam1Mask)
	}
}

const (
	RegisterMacl3l4c1rFieldL3saim1Shift = 3
	RegisterMacl3l4c1rFieldL3saim1Mask  = 0x8
)

// GetL3saim1 L3SAIM1
func (r *registerMacl3l4c1rType) GetL3saim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3saim1Mask) != 0
}

// SetL3saim1 L3SAIM1
func (r *registerMacl3l4c1rType) SetL3saim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c1rFieldL3saim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL3saim1Mask)
	}
}

const (
	RegisterMacl3l4c1rFieldL3dam1Shift = 4
	RegisterMacl3l4c1rFieldL3dam1Mask  = 0x10
)

// GetL3dam1 L3DAM1
func (r *registerMacl3l4c1rType) GetL3dam1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3dam1Mask) != 0
}

// SetL3dam1 L3DAM1
func (r *registerMacl3l4c1rType) SetL3dam1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c1rFieldL3dam1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL3dam1Mask)
	}
}

const (
	RegisterMacl3l4c1rFieldL3daim1Shift = 5
	RegisterMacl3l4c1rFieldL3daim1Mask  = 0x20
)

// GetL3daim1 L3DAIM1
func (r *registerMacl3l4c1rType) GetL3daim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3daim1Mask) != 0
}

// SetL3daim1 L3DAIM1
func (r *registerMacl3l4c1rType) SetL3daim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c1rFieldL3daim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL3daim1Mask)
	}
}

const (
	RegisterMacl3l4c1rFieldL3hsbm1Shift = 6
	RegisterMacl3l4c1rFieldL3hsbm1Mask  = 0x7c0
)

// GetL3hsbm1 L3HSBM1
func (r *registerMacl3l4c1rType) GetL3hsbm1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3hsbm1Mask) >> RegisterMacl3l4c1rFieldL3hsbm1Shift)
}

// SetL3hsbm1 L3HSBM1
func (r *registerMacl3l4c1rType) SetL3hsbm1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL3hsbm1Mask)|(uint32(value)<<RegisterMacl3l4c1rFieldL3hsbm1Shift))
}

const (
	RegisterMacl3l4c1rFieldL3hdbm1Shift = 11
	RegisterMacl3l4c1rFieldL3hdbm1Mask  = 0xf800
)

// GetL3hdbm1 L3HDBM1
func (r *registerMacl3l4c1rType) GetL3hdbm1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3hdbm1Mask) >> RegisterMacl3l4c1rFieldL3hdbm1Shift)
}

// SetL3hdbm1 L3HDBM1
func (r *registerMacl3l4c1rType) SetL3hdbm1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL3hdbm1Mask)|(uint32(value)<<RegisterMacl3l4c1rFieldL3hdbm1Shift))
}

const (
	RegisterMacl3l4c1rFieldL4pen1Shift = 16
	RegisterMacl3l4c1rFieldL4pen1Mask  = 0x10000
)

// GetL4pen1 L4PEN1
func (r *registerMacl3l4c1rType) GetL4pen1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL4pen1Mask) != 0
}

// SetL4pen1 L4PEN1
func (r *registerMacl3l4c1rType) SetL4pen1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c1rFieldL4pen1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL4pen1Mask)
	}
}

const (
	RegisterMacl3l4c1rFieldL4spm1Shift = 18
	RegisterMacl3l4c1rFieldL4spm1Mask  = 0x40000
)

// GetL4spm1 L4SPM1
func (r *registerMacl3l4c1rType) GetL4spm1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL4spm1Mask) != 0
}

// SetL4spm1 L4SPM1
func (r *registerMacl3l4c1rType) SetL4spm1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c1rFieldL4spm1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL4spm1Mask)
	}
}

const (
	RegisterMacl3l4c1rFieldL4spim1Shift = 19
	RegisterMacl3l4c1rFieldL4spim1Mask  = 0x80000
)

// GetL4spim1 L4SPIM1
func (r *registerMacl3l4c1rType) GetL4spim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL4spim1Mask) != 0
}

// SetL4spim1 L4SPIM1
func (r *registerMacl3l4c1rType) SetL4spim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c1rFieldL4spim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL4spim1Mask)
	}
}

const (
	RegisterMacl3l4c1rFieldL4dpm1Shift = 20
	RegisterMacl3l4c1rFieldL4dpm1Mask  = 0x100000
)

// GetL4dpm1 L4DPM1
func (r *registerMacl3l4c1rType) GetL4dpm1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL4dpm1Mask) != 0
}

// SetL4dpm1 L4DPM1
func (r *registerMacl3l4c1rType) SetL4dpm1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c1rFieldL4dpm1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL4dpm1Mask)
	}
}

const (
	RegisterMacl3l4c1rFieldL4dpim1Shift = 21
	RegisterMacl3l4c1rFieldL4dpim1Mask  = 0x200000
)

// GetL4dpim1 L4DPIM1
func (r *registerMacl3l4c1rType) GetL4dpim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL4dpim1Mask) != 0
}

// SetL4dpim1 L4DPIM1
func (r *registerMacl3l4c1rType) SetL4dpim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c1rFieldL4dpim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL4dpim1Mask)
	}
}

// registerMacl4a1rType Layer 4 address filter 1 register
type registerMacl4a1rType uint32

const (
	RegisterMacl4a1rFieldL4sp1Shift = 0
	RegisterMacl4a1rFieldL4sp1Mask  = 0xffff
)

// GetL4sp1 L4SP1
func (r *registerMacl4a1rType) GetL4sp1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacl4a1rFieldL4sp1Mask) >> RegisterMacl4a1rFieldL4sp1Shift)
}

// SetL4sp1 L4SP1
func (r *registerMacl4a1rType) SetL4sp1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl4a1rFieldL4sp1Mask)|(uint32(value)<<RegisterMacl4a1rFieldL4sp1Shift))
}

const (
	RegisterMacl4a1rFieldL4dp1Shift = 16
	RegisterMacl4a1rFieldL4dp1Mask  = 0xffff0000
)

// GetL4dp1 L4DP1
func (r *registerMacl4a1rType) GetL4dp1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacl4a1rFieldL4dp1Mask) >> RegisterMacl4a1rFieldL4dp1Shift)
}

// SetL4dp1 L4DP1
func (r *registerMacl4a1rType) SetL4dp1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl4a1rFieldL4dp1Mask)|(uint32(value)<<RegisterMacl4a1rFieldL4dp1Shift))
}

// registerMacl3a01rType Layer3 address 0 filter 1 Register
type registerMacl3a01rType uint32

const (
	RegisterMacl3a01rFieldL3a01Shift = 0
	RegisterMacl3a01rFieldL3a01Mask  = 0xffffffff
)

// GetL3a01 L3A01
func (r *registerMacl3a01rType) GetL3a01() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a01rFieldL3a01Mask) >> RegisterMacl3a01rFieldL3a01Shift)
}

// SetL3a01 L3A01
func (r *registerMacl3a01rType) SetL3a01(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a01rFieldL3a01Mask)|(uint32(value)<<RegisterMacl3a01rFieldL3a01Shift))
}

// registerMacl3a11rType Layer3 address 1 filter 1 register
type registerMacl3a11rType uint32

const (
	RegisterMacl3a11rFieldL3a11Shift = 0
	RegisterMacl3a11rFieldL3a11Mask  = 0xffffffff
)

// GetL3a11 L3A11
func (r *registerMacl3a11rType) GetL3a11() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a11rFieldL3a11Mask) >> RegisterMacl3a11rFieldL3a11Shift)
}

// SetL3a11 L3A11
func (r *registerMacl3a11rType) SetL3a11(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a11rFieldL3a11Mask)|(uint32(value)<<RegisterMacl3a11rFieldL3a11Shift))
}

// registerMacl3a21rType Layer3 address 2 filter 1 Register
type registerMacl3a21rType uint32

const (
	RegisterMacl3a21rFieldL3a21Shift = 0
	RegisterMacl3a21rFieldL3a21Mask  = 0xffffffff
)

// GetL3a21 L3A21
func (r *registerMacl3a21rType) GetL3a21() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a21rFieldL3a21Mask) >> RegisterMacl3a21rFieldL3a21Shift)
}

// SetL3a21 L3A21
func (r *registerMacl3a21rType) SetL3a21(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a21rFieldL3a21Mask)|(uint32(value)<<RegisterMacl3a21rFieldL3a21Shift))
}

// registerMacl3a31rType Layer3 address 3 filter 1 register
type registerMacl3a31rType uint32

const (
	RegisterMacl3a31rFieldL3a31Shift = 0
	RegisterMacl3a31rFieldL3a31Mask  = 0xffffffff
)

// GetL3a31 L3A31
func (r *registerMacl3a31rType) GetL3a31() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a31rFieldL3a31Mask) >> RegisterMacl3a31rFieldL3a31Shift)
}

// SetL3a31 L3A31
func (r *registerMacl3a31rType) SetL3a31(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a31rFieldL3a31Mask)|(uint32(value)<<RegisterMacl3a31rFieldL3a31Shift))
}

// registerMacarparType ARP address register
type registerMacarparType uint32

const (
	RegisterMacarparFieldArppaShift = 0
	RegisterMacarparFieldArppaMask  = 0xffffffff
)

// GetArppa ARPPA
func (r *registerMacarparType) GetArppa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacarparFieldArppaMask) >> RegisterMacarparFieldArppaShift)
}

// SetArppa ARPPA
func (r *registerMacarparType) SetArppa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacarparFieldArppaMask)|(uint32(value)<<RegisterMacarparFieldArppaShift))
}

// registerMactscrType Timestamp control Register
type registerMactscrType uint32

const (
	RegisterMactscrFieldTsenaShift = 0
	RegisterMactscrFieldTsenaMask  = 0x1
)

// GetTsena TSENA
func (r *registerMactscrType) GetTsena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsenaMask) != 0
}

// SetTsena TSENA
func (r *registerMactscrType) SetTsena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsenaMask)
	}
}

const (
	RegisterMactscrFieldTscfupdtShift = 1
	RegisterMactscrFieldTscfupdtMask  = 0x2
)

// GetTscfupdt TSCFUPDT
func (r *registerMactscrType) GetTscfupdt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTscfupdtMask) != 0
}

// SetTscfupdt TSCFUPDT
func (r *registerMactscrType) SetTscfupdt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTscfupdtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTscfupdtMask)
	}
}

const (
	RegisterMactscrFieldTsinitShift = 2
	RegisterMactscrFieldTsinitMask  = 0x4
)

// GetTsinit TSINIT
func (r *registerMactscrType) GetTsinit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsinitMask) != 0
}

// SetTsinit TSINIT
func (r *registerMactscrType) SetTsinit(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsinitMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsinitMask)
	}
}

const (
	RegisterMactscrFieldTsupdtShift = 3
	RegisterMactscrFieldTsupdtMask  = 0x8
)

// GetTsupdt TSUPDT
func (r *registerMactscrType) GetTsupdt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsupdtMask) != 0
}

// SetTsupdt TSUPDT
func (r *registerMactscrType) SetTsupdt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsupdtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsupdtMask)
	}
}

const (
	RegisterMactscrFieldTsaddregShift = 5
	RegisterMactscrFieldTsaddregMask  = 0x20
)

// GetTsaddreg TSADDREG
func (r *registerMactscrType) GetTsaddreg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsaddregMask) != 0
}

// SetTsaddreg TSADDREG
func (r *registerMactscrType) SetTsaddreg(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsaddregMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsaddregMask)
	}
}

const (
	RegisterMactscrFieldTsenallShift = 8
	RegisterMactscrFieldTsenallMask  = 0x100
)

// GetTsenall TSENALL
func (r *registerMactscrType) GetTsenall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsenallMask) != 0
}

// SetTsenall TSENALL
func (r *registerMactscrType) SetTsenall(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsenallMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsenallMask)
	}
}

const (
	RegisterMactscrFieldTsctrlssrShift = 9
	RegisterMactscrFieldTsctrlssrMask  = 0x200
)

// GetTsctrlssr TSCTRLSSR
func (r *registerMactscrType) GetTsctrlssr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsctrlssrMask) != 0
}

// SetTsctrlssr TSCTRLSSR
func (r *registerMactscrType) SetTsctrlssr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsctrlssrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsctrlssrMask)
	}
}

const (
	RegisterMactscrFieldTsver2enaShift = 10
	RegisterMactscrFieldTsver2enaMask  = 0x400
)

// GetTsver2ena TSVER2ENA
func (r *registerMactscrType) GetTsver2ena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsver2enaMask) != 0
}

// SetTsver2ena TSVER2ENA
func (r *registerMactscrType) SetTsver2ena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsver2enaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsver2enaMask)
	}
}

const (
	RegisterMactscrFieldTsipenaShift = 11
	RegisterMactscrFieldTsipenaMask  = 0x800
)

// GetTsipena TSIPENA
func (r *registerMactscrType) GetTsipena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsipenaMask) != 0
}

// SetTsipena TSIPENA
func (r *registerMactscrType) SetTsipena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsipenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsipenaMask)
	}
}

const (
	RegisterMactscrFieldTsipv6enaShift = 12
	RegisterMactscrFieldTsipv6enaMask  = 0x1000
)

// GetTsipv6ena TSIPV6ENA
func (r *registerMactscrType) GetTsipv6ena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsipv6enaMask) != 0
}

// SetTsipv6ena TSIPV6ENA
func (r *registerMactscrType) SetTsipv6ena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsipv6enaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsipv6enaMask)
	}
}

const (
	RegisterMactscrFieldTsipv4enaShift = 13
	RegisterMactscrFieldTsipv4enaMask  = 0x2000
)

// GetTsipv4ena TSIPV4ENA
func (r *registerMactscrType) GetTsipv4ena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsipv4enaMask) != 0
}

// SetTsipv4ena TSIPV4ENA
func (r *registerMactscrType) SetTsipv4ena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsipv4enaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsipv4enaMask)
	}
}

const (
	RegisterMactscrFieldTsevntenaShift = 14
	RegisterMactscrFieldTsevntenaMask  = 0x4000
)

// GetTsevntena TSEVNTENA
func (r *registerMactscrType) GetTsevntena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsevntenaMask) != 0
}

// SetTsevntena TSEVNTENA
func (r *registerMactscrType) SetTsevntena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsevntenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsevntenaMask)
	}
}

const (
	RegisterMactscrFieldTsmstrenaShift = 15
	RegisterMactscrFieldTsmstrenaMask  = 0x8000
)

// GetTsmstrena TSMSTRENA
func (r *registerMactscrType) GetTsmstrena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsmstrenaMask) != 0
}

// SetTsmstrena TSMSTRENA
func (r *registerMactscrType) SetTsmstrena(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsmstrenaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsmstrenaMask)
	}
}

const (
	RegisterMactscrFieldSnaptypselShift = 16
	RegisterMactscrFieldSnaptypselMask  = 0x30000
)

// GetSnaptypsel SNAPTYPSEL
func (r *registerMactscrType) GetSnaptypsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldSnaptypselMask) >> RegisterMactscrFieldSnaptypselShift)
}

// SetSnaptypsel SNAPTYPSEL
func (r *registerMactscrType) SetSnaptypsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldSnaptypselMask)|(uint32(value)<<RegisterMactscrFieldSnaptypselShift))
}

const (
	RegisterMactscrFieldTsenmacaddrShift = 18
	RegisterMactscrFieldTsenmacaddrMask  = 0x40000
)

// GetTsenmacaddr TSENMACADDR
func (r *registerMactscrType) GetTsenmacaddr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsenmacaddrMask) != 0
}

// SetTsenmacaddr TSENMACADDR
func (r *registerMactscrType) SetTsenmacaddr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTsenmacaddrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTsenmacaddrMask)
	}
}

const (
	RegisterMactscrFieldCscShift = 19
	RegisterMactscrFieldCscMask  = 0x80000
)

// GetCsc CSC
func (r *registerMactscrType) GetCsc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldCscMask) != 0
}

const (
	RegisterMactscrFieldTxtsstsmShift = 24
	RegisterMactscrFieldTxtsstsmMask  = 0x1000000
)

// GetTxtsstsm TXTSSTSM
func (r *registerMactscrType) GetTxtsstsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTxtsstsmMask) != 0
}

// SetTxtsstsm TXTSSTSM
func (r *registerMactscrType) SetTxtsstsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTxtsstsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTxtsstsmMask)
	}
}

// registerMacssirType Sub-second increment register
type registerMacssirType uint32

const (
	RegisterMacssirFieldSnsincShift = 8
	RegisterMacssirFieldSnsincMask  = 0xff00
)

// GetSnsinc SNSINC
func (r *registerMacssirType) GetSnsinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacssirFieldSnsincMask) >> RegisterMacssirFieldSnsincShift)
}

// SetSnsinc SNSINC
func (r *registerMacssirType) SetSnsinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacssirFieldSnsincMask)|(uint32(value)<<RegisterMacssirFieldSnsincShift))
}

const (
	RegisterMacssirFieldSsincShift = 16
	RegisterMacssirFieldSsincMask  = 0xff0000
)

// GetSsinc SSINC
func (r *registerMacssirType) GetSsinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacssirFieldSsincMask) >> RegisterMacssirFieldSsincShift)
}

// SetSsinc SSINC
func (r *registerMacssirType) SetSsinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacssirFieldSsincMask)|(uint32(value)<<RegisterMacssirFieldSsincShift))
}

// registerMacstsrType System time seconds register
type registerMacstsrType uint32

const (
	RegisterMacstsrFieldTssShift = 0
	RegisterMacstsrFieldTssMask  = 0xffffffff
)

// GetTss TSS
func (r *registerMacstsrType) GetTss() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacstsrFieldTssMask) >> RegisterMacstsrFieldTssShift)
}

// SetTss TSS
func (r *registerMacstsrType) SetTss(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacstsrFieldTssMask)|(uint32(value)<<RegisterMacstsrFieldTssShift))
}

// registerMacstnrType System time nanoseconds register
type registerMacstnrType uint32

const (
	RegisterMacstnrFieldTsssShift = 0
	RegisterMacstnrFieldTsssMask  = 0x7fffffff
)

// GetTsss TSSS
func (r *registerMacstnrType) GetTsss() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacstnrFieldTsssMask) >> RegisterMacstnrFieldTsssShift)
}

// SetTsss TSSS
func (r *registerMacstnrType) SetTsss(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacstnrFieldTsssMask)|(uint32(value)<<RegisterMacstnrFieldTsssShift))
}

// registerMacstsurType System time seconds update register
type registerMacstsurType uint32

const (
	RegisterMacstsurFieldTssShift = 0
	RegisterMacstsurFieldTssMask  = 0xffffffff
)

// GetTss TSS
func (r *registerMacstsurType) GetTss() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacstsurFieldTssMask) >> RegisterMacstsurFieldTssShift)
}

// SetTss TSS
func (r *registerMacstsurType) SetTss(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacstsurFieldTssMask)|(uint32(value)<<RegisterMacstsurFieldTssShift))
}

// registerMacstnurType System time nanoseconds update register
type registerMacstnurType uint32

const (
	RegisterMacstnurFieldTsssShift = 0
	RegisterMacstnurFieldTsssMask  = 0x7fffffff
)

// GetTsss TSSS
func (r *registerMacstnurType) GetTsss() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacstnurFieldTsssMask) >> RegisterMacstnurFieldTsssShift)
}

// SetTsss TSSS
func (r *registerMacstnurType) SetTsss(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacstnurFieldTsssMask)|(uint32(value)<<RegisterMacstnurFieldTsssShift))
}

const (
	RegisterMacstnurFieldAddsubShift = 31
	RegisterMacstnurFieldAddsubMask  = 0x80000000
)

// GetAddsub ADDSUB
func (r *registerMacstnurType) GetAddsub() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacstnurFieldAddsubMask) != 0
}

// SetAddsub ADDSUB
func (r *registerMacstnurType) SetAddsub(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacstnurFieldAddsubMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacstnurFieldAddsubMask)
	}
}

// registerMactsarType Timestamp addend register
type registerMactsarType uint32

const (
	RegisterMactsarFieldTsarShift = 0
	RegisterMactsarFieldTsarMask  = 0xffffffff
)

// GetTsar TSAR
func (r *registerMactsarType) GetTsar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactsarFieldTsarMask) >> RegisterMactsarFieldTsarShift)
}

// SetTsar TSAR
func (r *registerMactsarType) SetTsar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactsarFieldTsarMask)|(uint32(value)<<RegisterMactsarFieldTsarShift))
}

// registerMactssrType Timestamp status register
type registerMactssrType uint32

const (
	RegisterMactssrFieldTssovfShift = 0
	RegisterMactssrFieldTssovfMask  = 0x1
)

// GetTssovf TSSOVF
func (r *registerMactssrType) GetTssovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldTssovfMask) != 0
}

// SetTssovf TSSOVF
func (r *registerMactssrType) SetTssovf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactssrFieldTssovfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactssrFieldTssovfMask)
	}
}

const (
	RegisterMactssrFieldTstargt0Shift = 1
	RegisterMactssrFieldTstargt0Mask  = 0x2
)

// GetTstargt0 TSTARGT0
func (r *registerMactssrType) GetTstargt0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldTstargt0Mask) != 0
}

// SetTstargt0 TSTARGT0
func (r *registerMactssrType) SetTstargt0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactssrFieldTstargt0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactssrFieldTstargt0Mask)
	}
}

const (
	RegisterMactssrFieldAuxtstrigShift = 2
	RegisterMactssrFieldAuxtstrigMask  = 0x4
)

// GetAuxtstrig AUXTSTRIG
func (r *registerMactssrType) GetAuxtstrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldAuxtstrigMask) != 0
}

// SetAuxtstrig AUXTSTRIG
func (r *registerMactssrType) SetAuxtstrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactssrFieldAuxtstrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactssrFieldAuxtstrigMask)
	}
}

const (
	RegisterMactssrFieldTstrgterr0Shift = 3
	RegisterMactssrFieldTstrgterr0Mask  = 0x8
)

// GetTstrgterr0 TSTRGTERR0
func (r *registerMactssrType) GetTstrgterr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldTstrgterr0Mask) != 0
}

// SetTstrgterr0 TSTRGTERR0
func (r *registerMactssrType) SetTstrgterr0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactssrFieldTstrgterr0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactssrFieldTstrgterr0Mask)
	}
}

const (
	RegisterMactssrFieldTxtssisShift = 15
	RegisterMactssrFieldTxtssisMask  = 0x8000
)

// GetTxtssis TXTSSIS
func (r *registerMactssrType) GetTxtssis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldTxtssisMask) != 0
}

// SetTxtssis TXTSSIS
func (r *registerMactssrType) SetTxtssis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactssrFieldTxtssisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactssrFieldTxtssisMask)
	}
}

const (
	RegisterMactssrFieldAtsstnShift = 16
	RegisterMactssrFieldAtsstnMask  = 0xf0000
)

// GetAtsstn ATSSTN
func (r *registerMactssrType) GetAtsstn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldAtsstnMask) >> RegisterMactssrFieldAtsstnShift)
}

// SetAtsstn ATSSTN
func (r *registerMactssrType) SetAtsstn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactssrFieldAtsstnMask)|(uint32(value)<<RegisterMactssrFieldAtsstnShift))
}

const (
	RegisterMactssrFieldAtsstmShift = 24
	RegisterMactssrFieldAtsstmMask  = 0x1000000
)

// GetAtsstm ATSSTM
func (r *registerMactssrType) GetAtsstm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldAtsstmMask) != 0
}

// SetAtsstm ATSSTM
func (r *registerMactssrType) SetAtsstm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactssrFieldAtsstmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactssrFieldAtsstmMask)
	}
}

const (
	RegisterMactssrFieldAtsnsShift = 25
	RegisterMactssrFieldAtsnsMask  = 0x3e000000
)

// GetAtsns ATSNS
func (r *registerMactssrType) GetAtsns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldAtsnsMask) >> RegisterMactssrFieldAtsnsShift)
}

// SetAtsns ATSNS
func (r *registerMactssrType) SetAtsns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactssrFieldAtsnsMask)|(uint32(value)<<RegisterMactssrFieldAtsnsShift))
}

// registerMactxtssnrType Tx timestamp status nanoseconds register
type registerMactxtssnrType uint32

const (
	RegisterMactxtssnrFieldTxtssloShift = 0
	RegisterMactxtssnrFieldTxtssloMask  = 0x7fffffff
)

// GetTxtsslo TXTSSLO
func (r *registerMactxtssnrType) GetTxtsslo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactxtssnrFieldTxtssloMask) >> RegisterMactxtssnrFieldTxtssloShift)
}

// SetTxtsslo TXTSSLO
func (r *registerMactxtssnrType) SetTxtsslo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactxtssnrFieldTxtssloMask)|(uint32(value)<<RegisterMactxtssnrFieldTxtssloShift))
}

const (
	RegisterMactxtssnrFieldTxtssmisShift = 31
	RegisterMactxtssnrFieldTxtssmisMask  = 0x80000000
)

// GetTxtssmis TXTSSMIS
func (r *registerMactxtssnrType) GetTxtssmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactxtssnrFieldTxtssmisMask) != 0
}

// SetTxtssmis TXTSSMIS
func (r *registerMactxtssnrType) SetTxtssmis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactxtssnrFieldTxtssmisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactxtssnrFieldTxtssmisMask)
	}
}

// registerMactxtsssrType Tx timestamp status seconds register
type registerMactxtsssrType uint32

const (
	RegisterMactxtsssrFieldTxtsshiShift = 0
	RegisterMactxtsssrFieldTxtsshiMask  = 0xffffffff
)

// GetTxtsshi TXTSSHI
func (r *registerMactxtsssrType) GetTxtsshi() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactxtsssrFieldTxtsshiMask) >> RegisterMactxtsssrFieldTxtsshiShift)
}

// SetTxtsshi TXTSSHI
func (r *registerMactxtsssrType) SetTxtsshi(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactxtsssrFieldTxtsshiMask)|(uint32(value)<<RegisterMactxtsssrFieldTxtsshiShift))
}

// registerMacacrType Auxiliary control register
type registerMacacrType uint32

const (
	RegisterMacacrFieldAtsfcShift = 0
	RegisterMacacrFieldAtsfcMask  = 0x1
)

// GetAtsfc ATSFC
func (r *registerMacacrType) GetAtsfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacacrFieldAtsfcMask) != 0
}

// SetAtsfc ATSFC
func (r *registerMacacrType) SetAtsfc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacacrFieldAtsfcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacacrFieldAtsfcMask)
	}
}

const (
	RegisterMacacrFieldAtsen0Shift = 4
	RegisterMacacrFieldAtsen0Mask  = 0x10
)

// GetAtsen0 ATSEN0
func (r *registerMacacrType) GetAtsen0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacacrFieldAtsen0Mask) != 0
}

// SetAtsen0 ATSEN0
func (r *registerMacacrType) SetAtsen0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacacrFieldAtsen0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacacrFieldAtsen0Mask)
	}
}

const (
	RegisterMacacrFieldAtsen1Shift = 5
	RegisterMacacrFieldAtsen1Mask  = 0x20
)

// GetAtsen1 ATSEN1
func (r *registerMacacrType) GetAtsen1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacacrFieldAtsen1Mask) != 0
}

// SetAtsen1 ATSEN1
func (r *registerMacacrType) SetAtsen1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacacrFieldAtsen1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacacrFieldAtsen1Mask)
	}
}

const (
	RegisterMacacrFieldAtsen2Shift = 6
	RegisterMacacrFieldAtsen2Mask  = 0x40
)

// GetAtsen2 ATSEN2
func (r *registerMacacrType) GetAtsen2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacacrFieldAtsen2Mask) != 0
}

// SetAtsen2 ATSEN2
func (r *registerMacacrType) SetAtsen2(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacacrFieldAtsen2Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacacrFieldAtsen2Mask)
	}
}

const (
	RegisterMacacrFieldAtsen3Shift = 7
	RegisterMacacrFieldAtsen3Mask  = 0x80
)

// GetAtsen3 ATSEN3
func (r *registerMacacrType) GetAtsen3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacacrFieldAtsen3Mask) != 0
}

// SetAtsen3 ATSEN3
func (r *registerMacacrType) SetAtsen3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacacrFieldAtsen3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacacrFieldAtsen3Mask)
	}
}

// registerMacatsnrType Auxiliary timestamp nanoseconds register
type registerMacatsnrType uint32

const (
	RegisterMacatsnrFieldAuxtsloShift = 0
	RegisterMacatsnrFieldAuxtsloMask  = 0x7fffffff
)

// GetAuxtslo AUXTSLO
func (r *registerMacatsnrType) GetAuxtslo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacatsnrFieldAuxtsloMask) >> RegisterMacatsnrFieldAuxtsloShift)
}

// SetAuxtslo AUXTSLO
func (r *registerMacatsnrType) SetAuxtslo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacatsnrFieldAuxtsloMask)|(uint32(value)<<RegisterMacatsnrFieldAuxtsloShift))
}

// registerMacatssrType Auxiliary timestamp seconds register
type registerMacatssrType uint32

const (
	RegisterMacatssrFieldAuxtshiShift = 0
	RegisterMacatssrFieldAuxtshiMask  = 0xffffffff
)

// GetAuxtshi AUXTSHI
func (r *registerMacatssrType) GetAuxtshi() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacatssrFieldAuxtshiMask) >> RegisterMacatssrFieldAuxtshiShift)
}

// SetAuxtshi AUXTSHI
func (r *registerMacatssrType) SetAuxtshi(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacatssrFieldAuxtshiMask)|(uint32(value)<<RegisterMacatssrFieldAuxtshiShift))
}

// registerMactsiacrType Timestamp Ingress asymmetric correction register
type registerMactsiacrType uint32

const (
	RegisterMactsiacrFieldOstiacShift = 0
	RegisterMactsiacrFieldOstiacMask  = 0xffffffff
)

// GetOstiac OSTIAC
func (r *registerMactsiacrType) GetOstiac() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactsiacrFieldOstiacMask) >> RegisterMactsiacrFieldOstiacShift)
}

// SetOstiac OSTIAC
func (r *registerMactsiacrType) SetOstiac(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactsiacrFieldOstiacMask)|(uint32(value)<<RegisterMactsiacrFieldOstiacShift))
}

// registerMactseacrType Timestamp Egress asymmetric correction register
type registerMactseacrType uint32

const (
	RegisterMactseacrFieldOsteacShift = 0
	RegisterMactseacrFieldOsteacMask  = 0xffffffff
)

// GetOsteac OSTEAC
func (r *registerMactseacrType) GetOsteac() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactseacrFieldOsteacMask) >> RegisterMactseacrFieldOsteacShift)
}

// SetOsteac OSTEAC
func (r *registerMactseacrType) SetOsteac(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactseacrFieldOsteacMask)|(uint32(value)<<RegisterMactseacrFieldOsteacShift))
}

// registerMactsicnrType Timestamp Ingress correction nanosecond register
type registerMactsicnrType uint32

const (
	RegisterMactsicnrFieldTsicShift = 0
	RegisterMactsicnrFieldTsicMask  = 0xffffffff
)

// GetTsic TSIC
func (r *registerMactsicnrType) GetTsic() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactsicnrFieldTsicMask) >> RegisterMactsicnrFieldTsicShift)
}

// SetTsic TSIC
func (r *registerMactsicnrType) SetTsic(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactsicnrFieldTsicMask)|(uint32(value)<<RegisterMactsicnrFieldTsicShift))
}

// registerMactsecnrType Timestamp Egress correction nanosecond register
type registerMactsecnrType uint32

const (
	RegisterMactsecnrFieldTsecShift = 0
	RegisterMactsecnrFieldTsecMask  = 0xffffffff
)

// GetTsec TSEC
func (r *registerMactsecnrType) GetTsec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactsecnrFieldTsecMask) >> RegisterMactsecnrFieldTsecShift)
}

// SetTsec TSEC
func (r *registerMactsecnrType) SetTsec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactsecnrFieldTsecMask)|(uint32(value)<<RegisterMactsecnrFieldTsecShift))
}

// registerMacppscrType PPS control register
type registerMacppscrType uint32

const (
	RegisterMacppscrFieldPpsctrlShift = 0
	RegisterMacppscrFieldPpsctrlMask  = 0xf
)

// GetPpsctrl PPSCTRL
func (r *registerMacppscrType) GetPpsctrl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacppscrFieldPpsctrlMask) >> RegisterMacppscrFieldPpsctrlShift)
}

// SetPpsctrl PPSCTRL
func (r *registerMacppscrType) SetPpsctrl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppscrFieldPpsctrlMask)|(uint32(value)<<RegisterMacppscrFieldPpsctrlShift))
}

const (
	RegisterMacppscrFieldPpsen0Shift = 4
	RegisterMacppscrFieldPpsen0Mask  = 0x10
)

// GetPpsen0 PPSEN0
func (r *registerMacppscrType) GetPpsen0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacppscrFieldPpsen0Mask) != 0
}

// SetPpsen0 PPSEN0
func (r *registerMacppscrType) SetPpsen0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacppscrFieldPpsen0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacppscrFieldPpsen0Mask)
	}
}

const (
	RegisterMacppscrFieldTrgtmodsel0Shift = 5
	RegisterMacppscrFieldTrgtmodsel0Mask  = 0x60
)

// GetTrgtmodsel0 TRGTMODSEL0
func (r *registerMacppscrType) GetTrgtmodsel0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacppscrFieldTrgtmodsel0Mask) >> RegisterMacppscrFieldTrgtmodsel0Shift)
}

// SetTrgtmodsel0 TRGTMODSEL0
func (r *registerMacppscrType) SetTrgtmodsel0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppscrFieldTrgtmodsel0Mask)|(uint32(value)<<RegisterMacppscrFieldTrgtmodsel0Shift))
}

// registerMacppsttsrType PPS target time seconds register
type registerMacppsttsrType uint32

const (
	RegisterMacppsttsrFieldTstrh0Shift = 0
	RegisterMacppsttsrFieldTstrh0Mask  = 0x7fffffff
)

// GetTstrh0 TSTRH0
func (r *registerMacppsttsrType) GetTstrh0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacppsttsrFieldTstrh0Mask) >> RegisterMacppsttsrFieldTstrh0Shift)
}

// SetTstrh0 TSTRH0
func (r *registerMacppsttsrType) SetTstrh0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppsttsrFieldTstrh0Mask)|(uint32(value)<<RegisterMacppsttsrFieldTstrh0Shift))
}

// registerMacppsttnrType PPS target time nanoseconds register
type registerMacppsttnrType uint32

const (
	RegisterMacppsttnrFieldTtsl0Shift = 0
	RegisterMacppsttnrFieldTtsl0Mask  = 0x7fffffff
)

// GetTtsl0 TTSL0
func (r *registerMacppsttnrType) GetTtsl0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacppsttnrFieldTtsl0Mask) >> RegisterMacppsttnrFieldTtsl0Shift)
}

// SetTtsl0 TTSL0
func (r *registerMacppsttnrType) SetTtsl0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppsttnrFieldTtsl0Mask)|(uint32(value)<<RegisterMacppsttnrFieldTtsl0Shift))
}

const (
	RegisterMacppsttnrFieldTrgtbusy0Shift = 31
	RegisterMacppsttnrFieldTrgtbusy0Mask  = 0x80000000
)

// GetTrgtbusy0 TRGTBUSY0
func (r *registerMacppsttnrType) GetTrgtbusy0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacppsttnrFieldTrgtbusy0Mask) != 0
}

// SetTrgtbusy0 TRGTBUSY0
func (r *registerMacppsttnrType) SetTrgtbusy0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacppsttnrFieldTrgtbusy0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacppsttnrFieldTrgtbusy0Mask)
	}
}

// registerMacppsirType PPS interval register
type registerMacppsirType uint32

const (
	RegisterMacppsirFieldPpsint0Shift = 0
	RegisterMacppsirFieldPpsint0Mask  = 0xffffffff
)

// GetPpsint0 PPSINT0
func (r *registerMacppsirType) GetPpsint0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacppsirFieldPpsint0Mask) >> RegisterMacppsirFieldPpsint0Shift)
}

// SetPpsint0 PPSINT0
func (r *registerMacppsirType) SetPpsint0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppsirFieldPpsint0Mask)|(uint32(value)<<RegisterMacppsirFieldPpsint0Shift))
}

// registerMacppswrType PPS width register
type registerMacppswrType uint32

const (
	RegisterMacppswrFieldPpswidth0Shift = 0
	RegisterMacppswrFieldPpswidth0Mask  = 0xffffffff
)

// GetPpswidth0 PPSWIDTH0
func (r *registerMacppswrType) GetPpswidth0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacppswrFieldPpswidth0Mask) >> RegisterMacppswrFieldPpswidth0Shift)
}

// SetPpswidth0 PPSWIDTH0
func (r *registerMacppswrType) SetPpswidth0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppswrFieldPpswidth0Mask)|(uint32(value)<<RegisterMacppswrFieldPpswidth0Shift))
}

// registerMacpocrType PTP Offload control register
type registerMacpocrType uint32

const (
	RegisterMacpocrFieldPtoenShift = 0
	RegisterMacpocrFieldPtoenMask  = 0x1
)

// GetPtoen PTOEN
func (r *registerMacpocrType) GetPtoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldPtoenMask) != 0
}

// SetPtoen PTOEN
func (r *registerMacpocrType) SetPtoen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpocrFieldPtoenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpocrFieldPtoenMask)
	}
}

const (
	RegisterMacpocrFieldAsyncenShift = 1
	RegisterMacpocrFieldAsyncenMask  = 0x2
)

// GetAsyncen ASYNCEN
func (r *registerMacpocrType) GetAsyncen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldAsyncenMask) != 0
}

// SetAsyncen ASYNCEN
func (r *registerMacpocrType) SetAsyncen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpocrFieldAsyncenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpocrFieldAsyncenMask)
	}
}

const (
	RegisterMacpocrFieldApdreqenShift = 2
	RegisterMacpocrFieldApdreqenMask  = 0x4
)

// GetApdreqen APDREQEN
func (r *registerMacpocrType) GetApdreqen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldApdreqenMask) != 0
}

// SetApdreqen APDREQEN
func (r *registerMacpocrType) SetApdreqen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpocrFieldApdreqenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpocrFieldApdreqenMask)
	}
}

const (
	RegisterMacpocrFieldAsynctrigShift = 4
	RegisterMacpocrFieldAsynctrigMask  = 0x10
)

// GetAsynctrig ASYNCTRIG
func (r *registerMacpocrType) GetAsynctrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldAsynctrigMask) != 0
}

// SetAsynctrig ASYNCTRIG
func (r *registerMacpocrType) SetAsynctrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpocrFieldAsynctrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpocrFieldAsynctrigMask)
	}
}

const (
	RegisterMacpocrFieldApdreqtrigShift = 5
	RegisterMacpocrFieldApdreqtrigMask  = 0x20
)

// GetApdreqtrig APDREQTRIG
func (r *registerMacpocrType) GetApdreqtrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldApdreqtrigMask) != 0
}

// SetApdreqtrig APDREQTRIG
func (r *registerMacpocrType) SetApdreqtrig(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpocrFieldApdreqtrigMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpocrFieldApdreqtrigMask)
	}
}

const (
	RegisterMacpocrFieldDrrdisShift = 6
	RegisterMacpocrFieldDrrdisMask  = 0x40
)

// GetDrrdis DRRDIS
func (r *registerMacpocrType) GetDrrdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldDrrdisMask) != 0
}

// SetDrrdis DRRDIS
func (r *registerMacpocrType) SetDrrdis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpocrFieldDrrdisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpocrFieldDrrdisMask)
	}
}

const (
	RegisterMacpocrFieldDnShift = 8
	RegisterMacpocrFieldDnMask  = 0xff00
)

// GetDn DN
func (r *registerMacpocrType) GetDn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldDnMask) >> RegisterMacpocrFieldDnShift)
}

// SetDn DN
func (r *registerMacpocrType) SetDn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacpocrFieldDnMask)|(uint32(value)<<RegisterMacpocrFieldDnShift))
}

// registerMacspi0rType PTP Source Port Identity 0 Register
type registerMacspi0rType uint32

const (
	RegisterMacspi0rFieldSpi0Shift = 0
	RegisterMacspi0rFieldSpi0Mask  = 0xffffffff
)

// GetSpi0 SPI0
func (r *registerMacspi0rType) GetSpi0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacspi0rFieldSpi0Mask) >> RegisterMacspi0rFieldSpi0Shift)
}

// SetSpi0 SPI0
func (r *registerMacspi0rType) SetSpi0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacspi0rFieldSpi0Mask)|(uint32(value)<<RegisterMacspi0rFieldSpi0Shift))
}

// registerMacspi1rType PTP Source port identity 1 register
type registerMacspi1rType uint32

const (
	RegisterMacspi1rFieldSpi1Shift = 0
	RegisterMacspi1rFieldSpi1Mask  = 0xffffffff
)

// GetSpi1 SPI1
func (r *registerMacspi1rType) GetSpi1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacspi1rFieldSpi1Mask) >> RegisterMacspi1rFieldSpi1Shift)
}

// SetSpi1 SPI1
func (r *registerMacspi1rType) SetSpi1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacspi1rFieldSpi1Mask)|(uint32(value)<<RegisterMacspi1rFieldSpi1Shift))
}

// registerMacspi2rType PTP Source port identity 2 register
type registerMacspi2rType uint32

const (
	RegisterMacspi2rFieldSpi2Shift = 0
	RegisterMacspi2rFieldSpi2Mask  = 0xffff
)

// GetSpi2 SPI2
func (r *registerMacspi2rType) GetSpi2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacspi2rFieldSpi2Mask) >> RegisterMacspi2rFieldSpi2Shift)
}

// SetSpi2 SPI2
func (r *registerMacspi2rType) SetSpi2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacspi2rFieldSpi2Mask)|(uint32(value)<<RegisterMacspi2rFieldSpi2Shift))
}

// registerMaclmirType Log message interval register
type registerMaclmirType uint32

const (
	RegisterMaclmirFieldLsiShift = 0
	RegisterMaclmirFieldLsiMask  = 0xff
)

// GetLsi LSI
func (r *registerMaclmirType) GetLsi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaclmirFieldLsiMask) >> RegisterMaclmirFieldLsiShift)
}

// SetLsi LSI
func (r *registerMaclmirType) SetLsi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaclmirFieldLsiMask)|(uint32(value)<<RegisterMaclmirFieldLsiShift))
}

const (
	RegisterMaclmirFieldDrsyncrShift = 8
	RegisterMaclmirFieldDrsyncrMask  = 0x700
)

// GetDrsyncr DRSYNCR
func (r *registerMaclmirType) GetDrsyncr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaclmirFieldDrsyncrMask) >> RegisterMaclmirFieldDrsyncrShift)
}

// SetDrsyncr DRSYNCR
func (r *registerMaclmirType) SetDrsyncr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaclmirFieldDrsyncrMask)|(uint32(value)<<RegisterMaclmirFieldDrsyncrShift))
}

const (
	RegisterMaclmirFieldLmpdriShift = 24
	RegisterMaclmirFieldLmpdriMask  = 0xff000000
)

// GetLmpdri LMPDRI
func (r *registerMaclmirType) GetLmpdri() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaclmirFieldLmpdriMask) >> RegisterMaclmirFieldLmpdriShift)
}

// SetLmpdri LMPDRI
func (r *registerMaclmirType) SetLmpdri(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaclmirFieldLmpdriMask)|(uint32(value)<<RegisterMaclmirFieldLmpdriShift))
}

// registerMtlomrType Operating mode Register
type registerMtlomrType uint32

const (
	RegisterMtlomrFieldDtxstsShift = 1
	RegisterMtlomrFieldDtxstsMask  = 0x2
)

// GetDtxsts DTXSTS
func (r *registerMtlomrType) GetDtxsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlomrFieldDtxstsMask) != 0
}

// SetDtxsts DTXSTS
func (r *registerMtlomrType) SetDtxsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlomrFieldDtxstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlomrFieldDtxstsMask)
	}
}

const (
	RegisterMtlomrFieldCntprstShift = 8
	RegisterMtlomrFieldCntprstMask  = 0x100
)

// GetCntprst CNTPRST
func (r *registerMtlomrType) GetCntprst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlomrFieldCntprstMask) != 0
}

// SetCntprst CNTPRST
func (r *registerMtlomrType) SetCntprst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlomrFieldCntprstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlomrFieldCntprstMask)
	}
}

const (
	RegisterMtlomrFieldCntclrShift = 9
	RegisterMtlomrFieldCntclrMask  = 0x200
)

// GetCntclr CNTCLR
func (r *registerMtlomrType) GetCntclr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlomrFieldCntclrMask) != 0
}

// SetCntclr CNTCLR
func (r *registerMtlomrType) SetCntclr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlomrFieldCntclrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlomrFieldCntclrMask)
	}
}

// registerMtlisrType Interrupt status Register
type registerMtlisrType uint32

const (
	RegisterMtlisrFieldQ0isShift = 0
	RegisterMtlisrFieldQ0isMask  = 0x1
)

// GetQ0is Queue interrupt status
func (r *registerMtlisrType) GetQ0is() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlisrFieldQ0isMask) != 0
}

// SetQ0is Queue interrupt status
func (r *registerMtlisrType) SetQ0is(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlisrFieldQ0isMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlisrFieldQ0isMask)
	}
}

// registerMtltxqomrType Tx queue operating mode Register
type registerMtltxqomrType uint32

const (
	RegisterMtltxqomrFieldFtqShift = 0
	RegisterMtltxqomrFieldFtqMask  = 0x1
)

// GetFtq Flush Transmit Queue
func (r *registerMtltxqomrType) GetFtq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqomrFieldFtqMask) != 0
}

// SetFtq Flush Transmit Queue
func (r *registerMtltxqomrType) SetFtq(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtltxqomrFieldFtqMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqomrFieldFtqMask)
	}
}

const (
	RegisterMtltxqomrFieldTsfShift = 1
	RegisterMtltxqomrFieldTsfMask  = 0x2
)

// GetTsf Transmit Store and Forward
func (r *registerMtltxqomrType) GetTsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqomrFieldTsfMask) != 0
}

// SetTsf Transmit Store and Forward
func (r *registerMtltxqomrType) SetTsf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtltxqomrFieldTsfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqomrFieldTsfMask)
	}
}

const (
	RegisterMtltxqomrFieldTxqenShift = 2
	RegisterMtltxqomrFieldTxqenMask  = 0xc
)

// GetTxqen Transmit Queue Enable
func (r *registerMtltxqomrType) GetTxqen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqomrFieldTxqenMask) >> RegisterMtltxqomrFieldTxqenShift)
}

const (
	RegisterMtltxqomrFieldTtcShift = 4
	RegisterMtltxqomrFieldTtcMask  = 0x70
)

// GetTtc Transmit Threshold Control
func (r *registerMtltxqomrType) GetTtc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqomrFieldTtcMask) >> RegisterMtltxqomrFieldTtcShift)
}

// SetTtc Transmit Threshold Control
func (r *registerMtltxqomrType) SetTtc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqomrFieldTtcMask)|(uint32(value)<<RegisterMtltxqomrFieldTtcShift))
}

const (
	RegisterMtltxqomrFieldTqsShift = 16
	RegisterMtltxqomrFieldTqsMask  = 0x1ff0000
)

// GetTqs Transmit Queue Size
func (r *registerMtltxqomrType) GetTqs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqomrFieldTqsMask) >> RegisterMtltxqomrFieldTqsShift)
}

// SetTqs Transmit Queue Size
func (r *registerMtltxqomrType) SetTqs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqomrFieldTqsMask)|(uint32(value)<<RegisterMtltxqomrFieldTqsShift))
}

// registerMtltxqurType Tx queue underflow register
type registerMtltxqurType uint32

const (
	RegisterMtltxqurFieldUffrmcntShift = 0
	RegisterMtltxqurFieldUffrmcntMask  = 0x7ff
)

// GetUffrmcnt Underflow Packet Counter
func (r *registerMtltxqurType) GetUffrmcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqurFieldUffrmcntMask) >> RegisterMtltxqurFieldUffrmcntShift)
}

// SetUffrmcnt Underflow Packet Counter
func (r *registerMtltxqurType) SetUffrmcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqurFieldUffrmcntMask)|(uint32(value)<<RegisterMtltxqurFieldUffrmcntShift))
}

const (
	RegisterMtltxqurFieldUfcntovfShift = 11
	RegisterMtltxqurFieldUfcntovfMask  = 0x800
)

// GetUfcntovf UFCNTOVF
func (r *registerMtltxqurType) GetUfcntovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqurFieldUfcntovfMask) != 0
}

// SetUfcntovf UFCNTOVF
func (r *registerMtltxqurType) SetUfcntovf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtltxqurFieldUfcntovfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqurFieldUfcntovfMask)
	}
}

// registerMtltxqdrType Tx queue debug Register
type registerMtltxqdrType uint32

const (
	RegisterMtltxqdrFieldTxqpausedShift = 0
	RegisterMtltxqdrFieldTxqpausedMask  = 0x1
)

// GetTxqpaused TXQPAUSED
func (r *registerMtltxqdrType) GetTxqpaused() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldTxqpausedMask) != 0
}

// SetTxqpaused TXQPAUSED
func (r *registerMtltxqdrType) SetTxqpaused(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtltxqdrFieldTxqpausedMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqdrFieldTxqpausedMask)
	}
}

const (
	RegisterMtltxqdrFieldTrcstsShift = 1
	RegisterMtltxqdrFieldTrcstsMask  = 0x6
)

// GetTrcsts TRCSTS
func (r *registerMtltxqdrType) GetTrcsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldTrcstsMask) >> RegisterMtltxqdrFieldTrcstsShift)
}

// SetTrcsts TRCSTS
func (r *registerMtltxqdrType) SetTrcsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqdrFieldTrcstsMask)|(uint32(value)<<RegisterMtltxqdrFieldTrcstsShift))
}

const (
	RegisterMtltxqdrFieldTwcstsShift = 3
	RegisterMtltxqdrFieldTwcstsMask  = 0x8
)

// GetTwcsts TWCSTS
func (r *registerMtltxqdrType) GetTwcsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldTwcstsMask) != 0
}

// SetTwcsts TWCSTS
func (r *registerMtltxqdrType) SetTwcsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtltxqdrFieldTwcstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqdrFieldTwcstsMask)
	}
}

const (
	RegisterMtltxqdrFieldTxqstsShift = 4
	RegisterMtltxqdrFieldTxqstsMask  = 0x10
)

// GetTxqsts TXQSTS
func (r *registerMtltxqdrType) GetTxqsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldTxqstsMask) != 0
}

// SetTxqsts TXQSTS
func (r *registerMtltxqdrType) SetTxqsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtltxqdrFieldTxqstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqdrFieldTxqstsMask)
	}
}

const (
	RegisterMtltxqdrFieldTxstsfstsShift = 5
	RegisterMtltxqdrFieldTxstsfstsMask  = 0x20
)

// GetTxstsfsts TXSTSFSTS
func (r *registerMtltxqdrType) GetTxstsfsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldTxstsfstsMask) != 0
}

// SetTxstsfsts TXSTSFSTS
func (r *registerMtltxqdrType) SetTxstsfsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtltxqdrFieldTxstsfstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqdrFieldTxstsfstsMask)
	}
}

const (
	RegisterMtltxqdrFieldPtxqShift = 16
	RegisterMtltxqdrFieldPtxqMask  = 0x70000
)

// GetPtxq PTXQ
func (r *registerMtltxqdrType) GetPtxq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldPtxqMask) >> RegisterMtltxqdrFieldPtxqShift)
}

// SetPtxq PTXQ
func (r *registerMtltxqdrType) SetPtxq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqdrFieldPtxqMask)|(uint32(value)<<RegisterMtltxqdrFieldPtxqShift))
}

const (
	RegisterMtltxqdrFieldStxstsfShift = 20
	RegisterMtltxqdrFieldStxstsfMask  = 0x700000
)

// GetStxstsf STXSTSF
func (r *registerMtltxqdrType) GetStxstsf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldStxstsfMask) >> RegisterMtltxqdrFieldStxstsfShift)
}

// SetStxstsf STXSTSF
func (r *registerMtltxqdrType) SetStxstsf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqdrFieldStxstsfMask)|(uint32(value)<<RegisterMtltxqdrFieldStxstsfShift))
}

// registerMtlqicsrType Queue interrupt control status Register
type registerMtlqicsrType uint32

const (
	RegisterMtlqicsrFieldTxunfisShift = 0
	RegisterMtlqicsrFieldTxunfisMask  = 0x1
)

// GetTxunfis TXUNFIS
func (r *registerMtlqicsrType) GetTxunfis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlqicsrFieldTxunfisMask) != 0
}

// SetTxunfis TXUNFIS
func (r *registerMtlqicsrType) SetTxunfis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlqicsrFieldTxunfisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlqicsrFieldTxunfisMask)
	}
}

const (
	RegisterMtlqicsrFieldTxuieShift = 8
	RegisterMtlqicsrFieldTxuieMask  = 0x100
)

// GetTxuie TXUIE
func (r *registerMtlqicsrType) GetTxuie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlqicsrFieldTxuieMask) != 0
}

// SetTxuie TXUIE
func (r *registerMtlqicsrType) SetTxuie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlqicsrFieldTxuieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlqicsrFieldTxuieMask)
	}
}

const (
	RegisterMtlqicsrFieldRxovfisShift = 16
	RegisterMtlqicsrFieldRxovfisMask  = 0x10000
)

// GetRxovfis RXOVFIS
func (r *registerMtlqicsrType) GetRxovfis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlqicsrFieldRxovfisMask) != 0
}

// SetRxovfis RXOVFIS
func (r *registerMtlqicsrType) SetRxovfis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlqicsrFieldRxovfisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlqicsrFieldRxovfisMask)
	}
}

const (
	RegisterMtlqicsrFieldRxoieShift = 24
	RegisterMtlqicsrFieldRxoieMask  = 0x1000000
)

// GetRxoie RXOIE
func (r *registerMtlqicsrType) GetRxoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlqicsrFieldRxoieMask) != 0
}

// SetRxoie RXOIE
func (r *registerMtlqicsrType) SetRxoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlqicsrFieldRxoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlqicsrFieldRxoieMask)
	}
}

// registerMtlrxqomrType Rx queue operating mode register
type registerMtlrxqomrType uint32

const (
	RegisterMtlrxqomrFieldRtcShift = 0
	RegisterMtlrxqomrFieldRtcMask  = 0x3
)

// GetRtc RTC
func (r *registerMtlrxqomrType) GetRtc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldRtcMask) >> RegisterMtlrxqomrFieldRtcShift)
}

// SetRtc RTC
func (r *registerMtlrxqomrType) SetRtc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldRtcMask)|(uint32(value)<<RegisterMtlrxqomrFieldRtcShift))
}

const (
	RegisterMtlrxqomrFieldFupShift = 3
	RegisterMtlrxqomrFieldFupMask  = 0x8
)

// GetFup FUP
func (r *registerMtlrxqomrType) GetFup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldFupMask) != 0
}

// SetFup FUP
func (r *registerMtlrxqomrType) SetFup(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlrxqomrFieldFupMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldFupMask)
	}
}

const (
	RegisterMtlrxqomrFieldFepShift = 4
	RegisterMtlrxqomrFieldFepMask  = 0x10
)

// GetFep FEP
func (r *registerMtlrxqomrType) GetFep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldFepMask) != 0
}

// SetFep FEP
func (r *registerMtlrxqomrType) SetFep(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlrxqomrFieldFepMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldFepMask)
	}
}

const (
	RegisterMtlrxqomrFieldRsfShift = 5
	RegisterMtlrxqomrFieldRsfMask  = 0x20
)

// GetRsf RSF
func (r *registerMtlrxqomrType) GetRsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldRsfMask) != 0
}

// SetRsf RSF
func (r *registerMtlrxqomrType) SetRsf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlrxqomrFieldRsfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldRsfMask)
	}
}

const (
	RegisterMtlrxqomrFieldDistcpefShift = 6
	RegisterMtlrxqomrFieldDistcpefMask  = 0x40
)

// GetDistcpef DIS_TCP_EF
func (r *registerMtlrxqomrType) GetDistcpef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldDistcpefMask) != 0
}

// SetDistcpef DIS_TCP_EF
func (r *registerMtlrxqomrType) SetDistcpef(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlrxqomrFieldDistcpefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldDistcpefMask)
	}
}

const (
	RegisterMtlrxqomrFieldEhfcShift = 7
	RegisterMtlrxqomrFieldEhfcMask  = 0x80
)

// GetEhfc EHFC
func (r *registerMtlrxqomrType) GetEhfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldEhfcMask) != 0
}

// SetEhfc EHFC
func (r *registerMtlrxqomrType) SetEhfc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlrxqomrFieldEhfcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldEhfcMask)
	}
}

const (
	RegisterMtlrxqomrFieldRfaShift = 8
	RegisterMtlrxqomrFieldRfaMask  = 0x700
)

// GetRfa RFA
func (r *registerMtlrxqomrType) GetRfa() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldRfaMask) >> RegisterMtlrxqomrFieldRfaShift)
}

// SetRfa RFA
func (r *registerMtlrxqomrType) SetRfa(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldRfaMask)|(uint32(value)<<RegisterMtlrxqomrFieldRfaShift))
}

const (
	RegisterMtlrxqomrFieldRfdShift = 14
	RegisterMtlrxqomrFieldRfdMask  = 0x1c000
)

// GetRfd RFD
func (r *registerMtlrxqomrType) GetRfd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldRfdMask) >> RegisterMtlrxqomrFieldRfdShift)
}

// SetRfd RFD
func (r *registerMtlrxqomrType) SetRfd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldRfdMask)|(uint32(value)<<RegisterMtlrxqomrFieldRfdShift))
}

const (
	RegisterMtlrxqomrFieldRqsShift = 20
	RegisterMtlrxqomrFieldRqsMask  = 0x700000
)

// GetRqs RQS
func (r *registerMtlrxqomrType) GetRqs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldRqsMask) >> RegisterMtlrxqomrFieldRqsShift)
}

// registerMtlrxqmpocrType Rx queue missed packet and overflow counter register
type registerMtlrxqmpocrType uint32

const (
	RegisterMtlrxqmpocrFieldOvfpktcntShift = 0
	RegisterMtlrxqmpocrFieldOvfpktcntMask  = 0x7ff
)

// GetOvfpktcnt OVFPKTCNT
func (r *registerMtlrxqmpocrType) GetOvfpktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqmpocrFieldOvfpktcntMask) >> RegisterMtlrxqmpocrFieldOvfpktcntShift)
}

// SetOvfpktcnt OVFPKTCNT
func (r *registerMtlrxqmpocrType) SetOvfpktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqmpocrFieldOvfpktcntMask)|(uint32(value)<<RegisterMtlrxqmpocrFieldOvfpktcntShift))
}

const (
	RegisterMtlrxqmpocrFieldOvfcntovfShift = 11
	RegisterMtlrxqmpocrFieldOvfcntovfMask  = 0x800
)

// GetOvfcntovf OVFCNTOVF
func (r *registerMtlrxqmpocrType) GetOvfcntovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqmpocrFieldOvfcntovfMask) != 0
}

// SetOvfcntovf OVFCNTOVF
func (r *registerMtlrxqmpocrType) SetOvfcntovf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlrxqmpocrFieldOvfcntovfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqmpocrFieldOvfcntovfMask)
	}
}

const (
	RegisterMtlrxqmpocrFieldMispktcntShift = 16
	RegisterMtlrxqmpocrFieldMispktcntMask  = 0x7ff0000
)

// GetMispktcnt MISPKTCNT
func (r *registerMtlrxqmpocrType) GetMispktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqmpocrFieldMispktcntMask) >> RegisterMtlrxqmpocrFieldMispktcntShift)
}

// SetMispktcnt MISPKTCNT
func (r *registerMtlrxqmpocrType) SetMispktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqmpocrFieldMispktcntMask)|(uint32(value)<<RegisterMtlrxqmpocrFieldMispktcntShift))
}

const (
	RegisterMtlrxqmpocrFieldMiscntovfShift = 27
	RegisterMtlrxqmpocrFieldMiscntovfMask  = 0x8000000
)

// GetMiscntovf MISCNTOVF
func (r *registerMtlrxqmpocrType) GetMiscntovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqmpocrFieldMiscntovfMask) != 0
}

// SetMiscntovf MISCNTOVF
func (r *registerMtlrxqmpocrType) SetMiscntovf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlrxqmpocrFieldMiscntovfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqmpocrFieldMiscntovfMask)
	}
}

// registerMtlrxqdrType Rx queue debug register
type registerMtlrxqdrType uint32

const (
	RegisterMtlrxqdrFieldRwcstsShift = 0
	RegisterMtlrxqdrFieldRwcstsMask  = 0x1
)

// GetRwcsts RWCSTS
func (r *registerMtlrxqdrType) GetRwcsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqdrFieldRwcstsMask) != 0
}

// SetRwcsts RWCSTS
func (r *registerMtlrxqdrType) SetRwcsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlrxqdrFieldRwcstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqdrFieldRwcstsMask)
	}
}

const (
	RegisterMtlrxqdrFieldRrcstsShift = 1
	RegisterMtlrxqdrFieldRrcstsMask  = 0x6
)

// GetRrcsts RRCSTS
func (r *registerMtlrxqdrType) GetRrcsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqdrFieldRrcstsMask) >> RegisterMtlrxqdrFieldRrcstsShift)
}

// SetRrcsts RRCSTS
func (r *registerMtlrxqdrType) SetRrcsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqdrFieldRrcstsMask)|(uint32(value)<<RegisterMtlrxqdrFieldRrcstsShift))
}

const (
	RegisterMtlrxqdrFieldRxqstsShift = 4
	RegisterMtlrxqdrFieldRxqstsMask  = 0x30
)

// GetRxqsts RXQSTS
func (r *registerMtlrxqdrType) GetRxqsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqdrFieldRxqstsMask) >> RegisterMtlrxqdrFieldRxqstsShift)
}

// SetRxqsts RXQSTS
func (r *registerMtlrxqdrType) SetRxqsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqdrFieldRxqstsMask)|(uint32(value)<<RegisterMtlrxqdrFieldRxqstsShift))
}

const (
	RegisterMtlrxqdrFieldPrxqShift = 16
	RegisterMtlrxqdrFieldPrxqMask  = 0x3fff0000
)

// GetPrxq PRXQ
func (r *registerMtlrxqdrType) GetPrxq() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqdrFieldPrxqMask) >> RegisterMtlrxqdrFieldPrxqShift)
}

// SetPrxq PRXQ
func (r *registerMtlrxqdrType) SetPrxq(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqdrFieldPrxqMask)|(uint32(value)<<RegisterMtlrxqdrFieldPrxqShift))
}

// registerDmamrType DMA mode register
type registerDmamrType uint32

const (
	RegisterDmamrFieldSwrShift = 0
	RegisterDmamrFieldSwrMask  = 0x1
)

// GetSwr Software Reset
func (r *registerDmamrType) GetSwr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmamrFieldSwrMask) != 0
}

// SetSwr Software Reset
func (r *registerDmamrType) SetSwr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmamrFieldSwrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmamrFieldSwrMask)
	}
}

const (
	RegisterDmamrFieldDaShift = 1
	RegisterDmamrFieldDaMask  = 0x2
)

// GetDa DMA Tx or Rx Arbitration Scheme
func (r *registerDmamrType) GetDa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmamrFieldDaMask) != 0
}

const (
	RegisterDmamrFieldTxprShift = 11
	RegisterDmamrFieldTxprMask  = 0x800
)

// GetTxpr Transmit priority
func (r *registerDmamrType) GetTxpr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmamrFieldTxprMask) != 0
}

const (
	RegisterDmamrFieldPrShift = 12
	RegisterDmamrFieldPrMask  = 0x7000
)

// GetPr Priority ratio
func (r *registerDmamrType) GetPr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmamrFieldPrMask) >> RegisterDmamrFieldPrShift)
}

const (
	RegisterDmamrFieldIntmShift = 16
	RegisterDmamrFieldIntmMask  = 0x10000
)

// GetIntm Interrupt Mode
func (r *registerDmamrType) GetIntm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmamrFieldIntmMask) != 0
}

// SetIntm Interrupt Mode
func (r *registerDmamrType) SetIntm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmamrFieldIntmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmamrFieldIntmMask)
	}
}

// registerDmasbmrType System bus mode register
type registerDmasbmrType uint32

const (
	RegisterDmasbmrFieldFbShift = 0
	RegisterDmasbmrFieldFbMask  = 0x1
)

// GetFb Fixed Burst Length
func (r *registerDmasbmrType) GetFb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmasbmrFieldFbMask) != 0
}

// SetFb Fixed Burst Length
func (r *registerDmasbmrType) SetFb(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmasbmrFieldFbMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmasbmrFieldFbMask)
	}
}

const (
	RegisterDmasbmrFieldAalShift = 12
	RegisterDmasbmrFieldAalMask  = 0x1000
)

// GetAal Address-Aligned Beats
func (r *registerDmasbmrType) GetAal() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmasbmrFieldAalMask) != 0
}

// SetAal Address-Aligned Beats
func (r *registerDmasbmrType) SetAal(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmasbmrFieldAalMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmasbmrFieldAalMask)
	}
}

const (
	RegisterDmasbmrFieldMbShift = 14
	RegisterDmasbmrFieldMbMask  = 0x4000
)

// GetMb Mixed Burst
func (r *registerDmasbmrType) GetMb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmasbmrFieldMbMask) != 0
}

const (
	RegisterDmasbmrFieldRbShift = 15
	RegisterDmasbmrFieldRbMask  = 0x8000
)

// GetRb Rebuild INCRx Burst
func (r *registerDmasbmrType) GetRb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmasbmrFieldRbMask) != 0
}

// registerDmaisrType Interrupt status register
type registerDmaisrType uint32

const (
	RegisterDmaisrFieldDc0isShift = 0
	RegisterDmaisrFieldDc0isMask  = 0x1
)

// GetDc0is DMA Channel Interrupt Status
func (r *registerDmaisrType) GetDc0is() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmaisrFieldDc0isMask) != 0
}

// SetDc0is DMA Channel Interrupt Status
func (r *registerDmaisrType) SetDc0is(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmaisrFieldDc0isMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmaisrFieldDc0isMask)
	}
}

const (
	RegisterDmaisrFieldMtlisShift = 16
	RegisterDmaisrFieldMtlisMask  = 0x10000
)

// GetMtlis MTL Interrupt Status
func (r *registerDmaisrType) GetMtlis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmaisrFieldMtlisMask) != 0
}

// SetMtlis MTL Interrupt Status
func (r *registerDmaisrType) SetMtlis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmaisrFieldMtlisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmaisrFieldMtlisMask)
	}
}

const (
	RegisterDmaisrFieldMacisShift = 17
	RegisterDmaisrFieldMacisMask  = 0x20000
)

// GetMacis MAC Interrupt Status
func (r *registerDmaisrType) GetMacis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmaisrFieldMacisMask) != 0
}

// SetMacis MAC Interrupt Status
func (r *registerDmaisrType) SetMacis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmaisrFieldMacisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmaisrFieldMacisMask)
	}
}

// registerDmadsrType Debug status register
type registerDmadsrType uint32

const (
	RegisterDmadsrFieldAxwhstsShift = 0
	RegisterDmadsrFieldAxwhstsMask  = 0x1
)

// GetAxwhsts AHB Master Write Channel
func (r *registerDmadsrType) GetAxwhsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmadsrFieldAxwhstsMask) != 0
}

// SetAxwhsts AHB Master Write Channel
func (r *registerDmadsrType) SetAxwhsts(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmadsrFieldAxwhstsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmadsrFieldAxwhstsMask)
	}
}

const (
	RegisterDmadsrFieldRps0Shift = 8
	RegisterDmadsrFieldRps0Mask  = 0xf00
)

// GetRps0 DMA Channel Receive Process State
func (r *registerDmadsrType) GetRps0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmadsrFieldRps0Mask) >> RegisterDmadsrFieldRps0Shift)
}

// SetRps0 DMA Channel Receive Process State
func (r *registerDmadsrType) SetRps0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmadsrFieldRps0Mask)|(uint32(value)<<RegisterDmadsrFieldRps0Shift))
}

const (
	RegisterDmadsrFieldTps0Shift = 12
	RegisterDmadsrFieldTps0Mask  = 0xf000
)

// GetTps0 DMA Channel Transmit Process State
func (r *registerDmadsrType) GetTps0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmadsrFieldTps0Mask) >> RegisterDmadsrFieldTps0Shift)
}

// SetTps0 DMA Channel Transmit Process State
func (r *registerDmadsrType) SetTps0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmadsrFieldTps0Mask)|(uint32(value)<<RegisterDmadsrFieldTps0Shift))
}

// registerDmaccrType Channel control register
type registerDmaccrType uint32

const (
	RegisterDmaccrFieldMssShift = 0
	RegisterDmaccrFieldMssMask  = 0x3fff
)

// GetMss Maximum Segment Size
func (r *registerDmaccrType) GetMss() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccrFieldMssMask) >> RegisterDmaccrFieldMssShift)
}

// SetMss Maximum Segment Size
func (r *registerDmaccrType) SetMss(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccrFieldMssMask)|(uint32(value)<<RegisterDmaccrFieldMssShift))
}

const (
	RegisterDmaccrFieldPblx8Shift = 16
	RegisterDmaccrFieldPblx8Mask  = 0x10000
)

// GetPblx8 8xPBL mode
func (r *registerDmaccrType) GetPblx8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmaccrFieldPblx8Mask) != 0
}

// SetPblx8 8xPBL mode
func (r *registerDmaccrType) SetPblx8(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmaccrFieldPblx8Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmaccrFieldPblx8Mask)
	}
}

const (
	RegisterDmaccrFieldDslShift = 18
	RegisterDmaccrFieldDslMask  = 0x1c0000
)

// GetDsl Descriptor Skip Length
func (r *registerDmaccrType) GetDsl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccrFieldDslMask) >> RegisterDmaccrFieldDslShift)
}

// SetDsl Descriptor Skip Length
func (r *registerDmaccrType) SetDsl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccrFieldDslMask)|(uint32(value)<<RegisterDmaccrFieldDslShift))
}

// registerDmactxcrType Channel transmit control register
type registerDmactxcrType uint32

const (
	RegisterDmactxcrFieldStShift = 0
	RegisterDmactxcrFieldStMask  = 0x1
)

// GetSt Start or Stop Transmission Command
func (r *registerDmactxcrType) GetSt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmactxcrFieldStMask) != 0
}

// SetSt Start or Stop Transmission Command
func (r *registerDmactxcrType) SetSt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmactxcrFieldStMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmactxcrFieldStMask)
	}
}

const (
	RegisterDmactxcrFieldOsfShift = 4
	RegisterDmactxcrFieldOsfMask  = 0x10
)

// GetOsf Operate on Second Packet
func (r *registerDmactxcrType) GetOsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmactxcrFieldOsfMask) != 0
}

// SetOsf Operate on Second Packet
func (r *registerDmactxcrType) SetOsf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmactxcrFieldOsfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmactxcrFieldOsfMask)
	}
}

const (
	RegisterDmactxcrFieldTseShift = 12
	RegisterDmactxcrFieldTseMask  = 0x1000
)

// GetTse TCP Segmentation Enabled
func (r *registerDmactxcrType) GetTse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmactxcrFieldTseMask) != 0
}

// SetTse TCP Segmentation Enabled
func (r *registerDmactxcrType) SetTse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmactxcrFieldTseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmactxcrFieldTseMask)
	}
}

const (
	RegisterDmactxcrFieldTxpblShift = 16
	RegisterDmactxcrFieldTxpblMask  = 0x3f0000
)

// GetTxpbl Transmit Programmable Burst Length
func (r *registerDmactxcrType) GetTxpbl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmactxcrFieldTxpblMask) >> RegisterDmactxcrFieldTxpblShift)
}

// SetTxpbl Transmit Programmable Burst Length
func (r *registerDmactxcrType) SetTxpbl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmactxcrFieldTxpblMask)|(uint32(value)<<RegisterDmactxcrFieldTxpblShift))
}

// registerDmacrxcrType Channel receive control register
type registerDmacrxcrType uint32

const (
	RegisterDmacrxcrFieldSrShift = 0
	RegisterDmacrxcrFieldSrMask  = 0x1
)

// GetSr Start or Stop Receive Command
func (r *registerDmacrxcrType) GetSr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxcrFieldSrMask) != 0
}

// SetSr Start or Stop Receive Command
func (r *registerDmacrxcrType) SetSr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacrxcrFieldSrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxcrFieldSrMask)
	}
}

const (
	RegisterDmacrxcrFieldRbszShift = 1
	RegisterDmacrxcrFieldRbszMask  = 0x7ffe
)

// GetRbsz Receive Buffer size
func (r *registerDmacrxcrType) GetRbsz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxcrFieldRbszMask) >> RegisterDmacrxcrFieldRbszShift)
}

// SetRbsz Receive Buffer size
func (r *registerDmacrxcrType) SetRbsz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxcrFieldRbszMask)|(uint32(value)<<RegisterDmacrxcrFieldRbszShift))
}

const (
	RegisterDmacrxcrFieldRxpblShift = 16
	RegisterDmacrxcrFieldRxpblMask  = 0x3f0000
)

// GetRxpbl RXPBL
func (r *registerDmacrxcrType) GetRxpbl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxcrFieldRxpblMask) >> RegisterDmacrxcrFieldRxpblShift)
}

// SetRxpbl RXPBL
func (r *registerDmacrxcrType) SetRxpbl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxcrFieldRxpblMask)|(uint32(value)<<RegisterDmacrxcrFieldRxpblShift))
}

const (
	RegisterDmacrxcrFieldRpfShift = 31
	RegisterDmacrxcrFieldRpfMask  = 0x80000000
)

// GetRpf DMA Rx Channel Packet Flush
func (r *registerDmacrxcrType) GetRpf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxcrFieldRpfMask) != 0
}

// SetRpf DMA Rx Channel Packet Flush
func (r *registerDmacrxcrType) SetRpf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacrxcrFieldRpfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxcrFieldRpfMask)
	}
}

// registerDmactxdlarType Channel Tx descriptor list address register
type registerDmactxdlarType uint32

const (
	RegisterDmactxdlarFieldTdeslaShift = 2
	RegisterDmactxdlarFieldTdeslaMask  = 0xfffffffc
)

// GetTdesla Start of Transmit List
func (r *registerDmactxdlarType) GetTdesla() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmactxdlarFieldTdeslaMask) >> RegisterDmactxdlarFieldTdeslaShift)
}

// SetTdesla Start of Transmit List
func (r *registerDmactxdlarType) SetTdesla(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmactxdlarFieldTdeslaMask)|(uint32(value)<<RegisterDmactxdlarFieldTdeslaShift))
}

// registerDmacrxdlarType Channel Rx descriptor list address register
type registerDmacrxdlarType uint32

const (
	RegisterDmacrxdlarFieldRdeslaShift = 2
	RegisterDmacrxdlarFieldRdeslaMask  = 0xfffffffc
)

// GetRdesla Start of Receive List
func (r *registerDmacrxdlarType) GetRdesla() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxdlarFieldRdeslaMask) >> RegisterDmacrxdlarFieldRdeslaShift)
}

// SetRdesla Start of Receive List
func (r *registerDmacrxdlarType) SetRdesla(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxdlarFieldRdeslaMask)|(uint32(value)<<RegisterDmacrxdlarFieldRdeslaShift))
}

// registerDmactxdtprType Channel Tx descriptor tail pointer register
type registerDmactxdtprType uint32

const (
	RegisterDmactxdtprFieldTdtShift = 2
	RegisterDmactxdtprFieldTdtMask  = 0xfffffffc
)

// GetTdt Transmit Descriptor Tail Pointer
func (r *registerDmactxdtprType) GetTdt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmactxdtprFieldTdtMask) >> RegisterDmactxdtprFieldTdtShift)
}

// SetTdt Transmit Descriptor Tail Pointer
func (r *registerDmactxdtprType) SetTdt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmactxdtprFieldTdtMask)|(uint32(value)<<RegisterDmactxdtprFieldTdtShift))
}

// registerDmacrxdtprType Channel Rx descriptor tail pointer register
type registerDmacrxdtprType uint32

const (
	RegisterDmacrxdtprFieldRdtShift = 2
	RegisterDmacrxdtprFieldRdtMask  = 0xfffffffc
)

// GetRdt Receive Descriptor Tail Pointer
func (r *registerDmacrxdtprType) GetRdt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxdtprFieldRdtMask) >> RegisterDmacrxdtprFieldRdtShift)
}

// SetRdt Receive Descriptor Tail Pointer
func (r *registerDmacrxdtprType) SetRdt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxdtprFieldRdtMask)|(uint32(value)<<RegisterDmacrxdtprFieldRdtShift))
}

// registerDmactxrlrType Channel Tx descriptor ring length register
type registerDmactxrlrType uint32

const (
	RegisterDmactxrlrFieldTdrlShift = 0
	RegisterDmactxrlrFieldTdrlMask  = 0x3ff
)

// GetTdrl Transmit Descriptor Ring Length
func (r *registerDmactxrlrType) GetTdrl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmactxrlrFieldTdrlMask) >> RegisterDmactxrlrFieldTdrlShift)
}

// SetTdrl Transmit Descriptor Ring Length
func (r *registerDmactxrlrType) SetTdrl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmactxrlrFieldTdrlMask)|(uint32(value)<<RegisterDmactxrlrFieldTdrlShift))
}

// registerDmacrxrlrType Channel Rx descriptor ring length register
type registerDmacrxrlrType uint32

const (
	RegisterDmacrxrlrFieldRdrlShift = 0
	RegisterDmacrxrlrFieldRdrlMask  = 0x3ff
)

// GetRdrl Receive Descriptor Ring Length
func (r *registerDmacrxrlrType) GetRdrl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxrlrFieldRdrlMask) >> RegisterDmacrxrlrFieldRdrlShift)
}

// SetRdrl Receive Descriptor Ring Length
func (r *registerDmacrxrlrType) SetRdrl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxrlrFieldRdrlMask)|(uint32(value)<<RegisterDmacrxrlrFieldRdrlShift))
}

// registerDmacierType Channel interrupt enable register
type registerDmacierType uint32

const (
	RegisterDmacierFieldTieShift = 0
	RegisterDmacierFieldTieMask  = 0x1
)

// GetTie Transmit Interrupt Enable
func (r *registerDmacierType) GetTie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldTieMask) != 0
}

// SetTie Transmit Interrupt Enable
func (r *registerDmacierType) SetTie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldTieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldTieMask)
	}
}

const (
	RegisterDmacierFieldTxseShift = 1
	RegisterDmacierFieldTxseMask  = 0x2
)

// GetTxse Transmit Stopped Enable
func (r *registerDmacierType) GetTxse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldTxseMask) != 0
}

// SetTxse Transmit Stopped Enable
func (r *registerDmacierType) SetTxse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldTxseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldTxseMask)
	}
}

const (
	RegisterDmacierFieldTbueShift = 2
	RegisterDmacierFieldTbueMask  = 0x4
)

// GetTbue Transmit Buffer Unavailable Enable
func (r *registerDmacierType) GetTbue() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldTbueMask) != 0
}

// SetTbue Transmit Buffer Unavailable Enable
func (r *registerDmacierType) SetTbue(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldTbueMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldTbueMask)
	}
}

const (
	RegisterDmacierFieldRieShift = 6
	RegisterDmacierFieldRieMask  = 0x40
)

// GetRie Receive Interrupt Enable
func (r *registerDmacierType) GetRie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldRieMask) != 0
}

// SetRie Receive Interrupt Enable
func (r *registerDmacierType) SetRie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldRieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldRieMask)
	}
}

const (
	RegisterDmacierFieldRbueShift = 7
	RegisterDmacierFieldRbueMask  = 0x80
)

// GetRbue Receive Buffer Unavailable Enable
func (r *registerDmacierType) GetRbue() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldRbueMask) != 0
}

// SetRbue Receive Buffer Unavailable Enable
func (r *registerDmacierType) SetRbue(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldRbueMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldRbueMask)
	}
}

const (
	RegisterDmacierFieldRseShift = 8
	RegisterDmacierFieldRseMask  = 0x100
)

// GetRse Receive Stopped Enable
func (r *registerDmacierType) GetRse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldRseMask) != 0
}

// SetRse Receive Stopped Enable
func (r *registerDmacierType) SetRse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldRseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldRseMask)
	}
}

const (
	RegisterDmacierFieldRwteShift = 9
	RegisterDmacierFieldRwteMask  = 0x200
)

// GetRwte Receive Watchdog Timeout Enable
func (r *registerDmacierType) GetRwte() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldRwteMask) != 0
}

// SetRwte Receive Watchdog Timeout Enable
func (r *registerDmacierType) SetRwte(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldRwteMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldRwteMask)
	}
}

const (
	RegisterDmacierFieldEtieShift = 10
	RegisterDmacierFieldEtieMask  = 0x400
)

// GetEtie Early Transmit Interrupt Enable
func (r *registerDmacierType) GetEtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldEtieMask) != 0
}

// SetEtie Early Transmit Interrupt Enable
func (r *registerDmacierType) SetEtie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldEtieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldEtieMask)
	}
}

const (
	RegisterDmacierFieldErieShift = 11
	RegisterDmacierFieldErieMask  = 0x800
)

// GetErie Early Receive Interrupt Enable
func (r *registerDmacierType) GetErie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldErieMask) != 0
}

// SetErie Early Receive Interrupt Enable
func (r *registerDmacierType) SetErie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldErieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldErieMask)
	}
}

const (
	RegisterDmacierFieldFbeeShift = 12
	RegisterDmacierFieldFbeeMask  = 0x1000
)

// GetFbee Fatal Bus Error Enable
func (r *registerDmacierType) GetFbee() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldFbeeMask) != 0
}

// SetFbee Fatal Bus Error Enable
func (r *registerDmacierType) SetFbee(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldFbeeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldFbeeMask)
	}
}

const (
	RegisterDmacierFieldCdeeShift = 13
	RegisterDmacierFieldCdeeMask  = 0x2000
)

// GetCdee Context Descriptor Error Enable
func (r *registerDmacierType) GetCdee() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldCdeeMask) != 0
}

// SetCdee Context Descriptor Error Enable
func (r *registerDmacierType) SetCdee(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldCdeeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldCdeeMask)
	}
}

const (
	RegisterDmacierFieldAieShift = 14
	RegisterDmacierFieldAieMask  = 0x4000
)

// GetAie Abnormal Interrupt Summary Enable
func (r *registerDmacierType) GetAie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldAieMask) != 0
}

// SetAie Abnormal Interrupt Summary Enable
func (r *registerDmacierType) SetAie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldAieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldAieMask)
	}
}

const (
	RegisterDmacierFieldNieShift = 15
	RegisterDmacierFieldNieMask  = 0x8000
)

// GetNie Normal Interrupt Summary Enable
func (r *registerDmacierType) GetNie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldNieMask) != 0
}

// SetNie Normal Interrupt Summary Enable
func (r *registerDmacierType) SetNie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldNieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldNieMask)
	}
}

// registerDmacrxiwtrType Channel Rx interrupt watchdog timer register
type registerDmacrxiwtrType uint32

const (
	RegisterDmacrxiwtrFieldRwtShift = 0
	RegisterDmacrxiwtrFieldRwtMask  = 0xff
)

// GetRwt Receive Interrupt Watchdog Timer Count
func (r *registerDmacrxiwtrType) GetRwt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxiwtrFieldRwtMask) >> RegisterDmacrxiwtrFieldRwtShift)
}

// SetRwt Receive Interrupt Watchdog Timer Count
func (r *registerDmacrxiwtrType) SetRwt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxiwtrFieldRwtMask)|(uint32(value)<<RegisterDmacrxiwtrFieldRwtShift))
}

// registerDmaccatxdrType Channel current application transmit descriptor register
type registerDmaccatxdrType uint32

const (
	RegisterDmaccatxdrFieldCurtdesaptrShift = 0
	RegisterDmaccatxdrFieldCurtdesaptrMask  = 0xffffffff
)

// GetCurtdesaptr Application Transmit Descriptor Address Pointer
func (r *registerDmaccatxdrType) GetCurtdesaptr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccatxdrFieldCurtdesaptrMask) >> RegisterDmaccatxdrFieldCurtdesaptrShift)
}

// SetCurtdesaptr Application Transmit Descriptor Address Pointer
func (r *registerDmaccatxdrType) SetCurtdesaptr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccatxdrFieldCurtdesaptrMask)|(uint32(value)<<RegisterDmaccatxdrFieldCurtdesaptrShift))
}

// registerDmaccarxdrType Channel current application receive descriptor register
type registerDmaccarxdrType uint32

const (
	RegisterDmaccarxdrFieldCurrdesaptrShift = 0
	RegisterDmaccarxdrFieldCurrdesaptrMask  = 0xffffffff
)

// GetCurrdesaptr Application Receive Descriptor Address Pointer
func (r *registerDmaccarxdrType) GetCurrdesaptr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccarxdrFieldCurrdesaptrMask) >> RegisterDmaccarxdrFieldCurrdesaptrShift)
}

// SetCurrdesaptr Application Receive Descriptor Address Pointer
func (r *registerDmaccarxdrType) SetCurrdesaptr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccarxdrFieldCurrdesaptrMask)|(uint32(value)<<RegisterDmaccarxdrFieldCurrdesaptrShift))
}

// registerDmaccatxbrType Channel current application transmit buffer register
type registerDmaccatxbrType uint32

const (
	RegisterDmaccatxbrFieldCurtbufaptrShift = 0
	RegisterDmaccatxbrFieldCurtbufaptrMask  = 0xffffffff
)

// GetCurtbufaptr Application Transmit Buffer Address Pointer
func (r *registerDmaccatxbrType) GetCurtbufaptr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccatxbrFieldCurtbufaptrMask) >> RegisterDmaccatxbrFieldCurtbufaptrShift)
}

// SetCurtbufaptr Application Transmit Buffer Address Pointer
func (r *registerDmaccatxbrType) SetCurtbufaptr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccatxbrFieldCurtbufaptrMask)|(uint32(value)<<RegisterDmaccatxbrFieldCurtbufaptrShift))
}

// registerDmaccarxbrType Channel current application receive buffer register
type registerDmaccarxbrType uint32

const (
	RegisterDmaccarxbrFieldCurrbufaptrShift = 0
	RegisterDmaccarxbrFieldCurrbufaptrMask  = 0xffffffff
)

// GetCurrbufaptr Application Receive Buffer Address Pointer
func (r *registerDmaccarxbrType) GetCurrbufaptr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccarxbrFieldCurrbufaptrMask) >> RegisterDmaccarxbrFieldCurrbufaptrShift)
}

// SetCurrbufaptr Application Receive Buffer Address Pointer
func (r *registerDmaccarxbrType) SetCurrbufaptr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccarxbrFieldCurrbufaptrMask)|(uint32(value)<<RegisterDmaccarxbrFieldCurrbufaptrShift))
}

// registerDmacsrType Channel status register
type registerDmacsrType uint32

const (
	RegisterDmacsrFieldTiShift = 0
	RegisterDmacsrFieldTiMask  = 0x1
)

// GetTi Transmit Interrupt
func (r *registerDmacsrType) GetTi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldTiMask) != 0
}

// SetTi Transmit Interrupt
func (r *registerDmacsrType) SetTi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldTiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldTiMask)
	}
}

const (
	RegisterDmacsrFieldTpsShift = 1
	RegisterDmacsrFieldTpsMask  = 0x2
)

// GetTps Transmit Process Stopped
func (r *registerDmacsrType) GetTps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldTpsMask) != 0
}

// SetTps Transmit Process Stopped
func (r *registerDmacsrType) SetTps(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldTpsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldTpsMask)
	}
}

const (
	RegisterDmacsrFieldTbuShift = 2
	RegisterDmacsrFieldTbuMask  = 0x4
)

// GetTbu Transmit Buffer Unavailable
func (r *registerDmacsrType) GetTbu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldTbuMask) != 0
}

// SetTbu Transmit Buffer Unavailable
func (r *registerDmacsrType) SetTbu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldTbuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldTbuMask)
	}
}

const (
	RegisterDmacsrFieldRiShift = 6
	RegisterDmacsrFieldRiMask  = 0x40
)

// GetRi Receive Interrupt
func (r *registerDmacsrType) GetRi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldRiMask) != 0
}

// SetRi Receive Interrupt
func (r *registerDmacsrType) SetRi(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldRiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldRiMask)
	}
}

const (
	RegisterDmacsrFieldRbuShift = 7
	RegisterDmacsrFieldRbuMask  = 0x80
)

// GetRbu Receive Buffer Unavailable
func (r *registerDmacsrType) GetRbu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldRbuMask) != 0
}

// SetRbu Receive Buffer Unavailable
func (r *registerDmacsrType) SetRbu(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldRbuMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldRbuMask)
	}
}

const (
	RegisterDmacsrFieldRpsShift = 8
	RegisterDmacsrFieldRpsMask  = 0x100
)

// GetRps Receive Process Stopped
func (r *registerDmacsrType) GetRps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldRpsMask) != 0
}

// SetRps Receive Process Stopped
func (r *registerDmacsrType) SetRps(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldRpsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldRpsMask)
	}
}

const (
	RegisterDmacsrFieldRwtShift = 9
	RegisterDmacsrFieldRwtMask  = 0x200
)

// GetRwt Receive Watchdog Timeout
func (r *registerDmacsrType) GetRwt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldRwtMask) != 0
}

// SetRwt Receive Watchdog Timeout
func (r *registerDmacsrType) SetRwt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldRwtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldRwtMask)
	}
}

const (
	RegisterDmacsrFieldEtShift = 10
	RegisterDmacsrFieldEtMask  = 0x400
)

// GetEt Early Transmit Interrupt
func (r *registerDmacsrType) GetEt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldEtMask) != 0
}

// SetEt Early Transmit Interrupt
func (r *registerDmacsrType) SetEt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldEtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldEtMask)
	}
}

const (
	RegisterDmacsrFieldErShift = 11
	RegisterDmacsrFieldErMask  = 0x800
)

// GetEr Early Receive Interrupt
func (r *registerDmacsrType) GetEr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldErMask) != 0
}

// SetEr Early Receive Interrupt
func (r *registerDmacsrType) SetEr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldErMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldErMask)
	}
}

const (
	RegisterDmacsrFieldFbeShift = 12
	RegisterDmacsrFieldFbeMask  = 0x1000
)

// GetFbe Fatal Bus Error
func (r *registerDmacsrType) GetFbe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldFbeMask) != 0
}

// SetFbe Fatal Bus Error
func (r *registerDmacsrType) SetFbe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldFbeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldFbeMask)
	}
}

const (
	RegisterDmacsrFieldCdeShift = 13
	RegisterDmacsrFieldCdeMask  = 0x2000
)

// GetCde Context Descriptor Error
func (r *registerDmacsrType) GetCde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldCdeMask) != 0
}

// SetCde Context Descriptor Error
func (r *registerDmacsrType) SetCde(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldCdeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldCdeMask)
	}
}

const (
	RegisterDmacsrFieldAisShift = 14
	RegisterDmacsrFieldAisMask  = 0x4000
)

// GetAis Abnormal Interrupt Summary
func (r *registerDmacsrType) GetAis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldAisMask) != 0
}

// SetAis Abnormal Interrupt Summary
func (r *registerDmacsrType) SetAis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldAisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldAisMask)
	}
}

const (
	RegisterDmacsrFieldNisShift = 15
	RegisterDmacsrFieldNisMask  = 0x8000
)

// GetNis Normal Interrupt Summary
func (r *registerDmacsrType) GetNis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldNisMask) != 0
}

// SetNis Normal Interrupt Summary
func (r *registerDmacsrType) SetNis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacsrFieldNisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacsrFieldNisMask)
	}
}

const (
	RegisterDmacsrFieldTebShift = 16
	RegisterDmacsrFieldTebMask  = 0x70000
)

// GetTeb Tx DMA Error Bits
func (r *registerDmacsrType) GetTeb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldTebMask) >> RegisterDmacsrFieldTebShift)
}

const (
	RegisterDmacsrFieldRebShift = 19
	RegisterDmacsrFieldRebMask  = 0x380000
)

// GetReb Rx DMA Error Bits
func (r *registerDmacsrType) GetReb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldRebMask) >> RegisterDmacsrFieldRebShift)
}

// registerDmacmfcrType Channel missed frame count register
type registerDmacmfcrType uint32

const (
	RegisterDmacmfcrFieldMfcShift = 0
	RegisterDmacmfcrFieldMfcMask  = 0x7ff
)

// GetMfc Dropped Packet Counters
func (r *registerDmacmfcrType) GetMfc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmacmfcrFieldMfcMask) >> RegisterDmacmfcrFieldMfcShift)
}

// SetMfc Dropped Packet Counters
func (r *registerDmacmfcrType) SetMfc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacmfcrFieldMfcMask)|(uint32(value)<<RegisterDmacmfcrFieldMfcShift))
}

const (
	RegisterDmacmfcrFieldMfcoShift = 15
	RegisterDmacmfcrFieldMfcoMask  = 0x8000
)

// GetMfco Overflow status of the MFC Counter
func (r *registerDmacmfcrType) GetMfco() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacmfcrFieldMfcoMask) != 0
}

// SetMfco Overflow status of the MFC Counter
func (r *registerDmacmfcrType) SetMfco(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacmfcrFieldMfcoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacmfcrFieldMfcoMask)
	}
}
