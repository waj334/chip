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
	Maccr                          RegisterMaccrType
	Macecr                         RegisterMacecrType
	Macpfr                         RegisterMacpfrType
	Macwtr                         RegisterMacwtrType
	Macht0r                        RegisterMacht0rType
	Macht1r                        RegisterMacht1rType
	_                              [56]byte
	Macvtr                         RegisterMacvtrType
	_                              [4]byte
	Macvhtr                        RegisterMacvhtrType
	_                              [4]byte
	Macvir                         RegisterMacvirType
	Macivir                        RegisterMacivirType
	_                              [8]byte
	Macqtxfcr                      RegisterMacqtxfcrType
	_                              [28]byte
	Macrxfcr                       RegisterMacrxfcrType
	_                              [28]byte
	Macisr                         RegisterMacisrType
	Macier                         RegisterMacierType
	Macrxtxsr                      RegisterMacrxtxsrType
	_                              [4]byte
	Macpcsr                        RegisterMacpcsrType
	Macrwkpfr                      RegisterMacrwkpfrType
	_                              [8]byte
	Maclcsr                        RegisterMaclcsrType
	Macltcr                        RegisterMacltcrType
	Macletr                        RegisterMacletrType
	Mac1ustcr                      RegisterMac1ustcrType
	_                              [48]byte
	Macvr                          RegisterMacvrType
	Macdr                          RegisterMacdrType
	_                              [8]byte
	Machwf1r                       RegisterMachwf1rType
	Machwf2r                       RegisterMachwf2rType
	_                              [216]byte
	Macmdioar                      RegisterMacmdioarType
	Macmdiodr                      RegisterMacmdiodrType
	_                              [248]byte
	Maca0hr                        RegisterMaca0hrType
	Maca0lr                        RegisterMaca0lrType
	Maca1hr                        RegisterMaca1hrType
	Maca1lr                        RegisterMaca1lrType
	Maca2hr                        RegisterMaca2hrType
	Maca2lr                        RegisterMaca2lrType
	Maca3hr                        RegisterMaca3hrType
	Maca3lr                        RegisterMaca3lrType
	_                              [992]byte
	Mmccontrol                     RegisterMmccontrolType
	Mmcrxinterrupt                 RegisterMmcrxinterruptType
	Mmctxinterrupt                 RegisterMmctxinterruptType
	Mmcrxinterruptmask             RegisterMmcrxinterruptmaskType
	Mmctxinterruptmask             RegisterMmctxinterruptmaskType
	_                              [56]byte
	Txsinglecollisiongoodpackets   RegisterTxsinglecollisiongoodpacketsType
	Txmultiplecollisiongoodpackets RegisterTxmultiplecollisiongoodpacketsType
	_                              [20]byte
	Txpacketcountgood              RegisterTxpacketcountgoodType
	_                              [40]byte
	Rxcrcerrorpackets              RegisterRxcrcerrorpacketsType
	Rxalignmenterrorpackets        RegisterRxalignmenterrorpacketsType
	_                              [40]byte
	Rxunicastpacketsgood           RegisterRxunicastpacketsgoodType
	_                              [36]byte
	Txlpiuseccntr                  RegisterTxlpiuseccntrType
	Txlpitrancntr                  RegisterTxlpitrancntrType
	Rxlpiuseccntr                  RegisterRxlpiuseccntrType
	Rxlpitrancntr                  RegisterRxlpitrancntrType
	_                              [260]byte
	Macl3l4c0r                     RegisterMacl3l4c0rType
	Macl4a0r                       RegisterMacl4a0rType
	_                              [8]byte
	Macl3a00r                      RegisterMacl3a00rType
	Macl3a10r                      RegisterMacl3a10rType
	Macl3a20                       RegisterMacl3a20Type
	Macl3a30                       RegisterMacl3a30Type
	_                              [16]byte
	Macl3l4c1r                     RegisterMacl3l4c1rType
	Macl4a1r                       RegisterMacl4a1rType
	_                              [8]byte
	Macl3a01r                      RegisterMacl3a01rType
	Macl3a11r                      RegisterMacl3a11rType
	Macl3a21r                      RegisterMacl3a21rType
	Macl3a31r                      RegisterMacl3a31rType
	_                              [400]byte
	Macarpar                       RegisterMacarparType
	_                              [28]byte
	Mactscr                        RegisterMactscrType
	Macssir                        RegisterMacssirType
	Macstsr                        RegisterMacstsrType
	Macstnr                        RegisterMacstnrType
	Macstsur                       RegisterMacstsurType
	Macstnur                       RegisterMacstnurType
	Mactsar                        RegisterMactsarType
	_                              [4]byte
	Mactssr                        RegisterMactssrType
	_                              [12]byte
	Mactxtssnr                     RegisterMactxtssnrType
	Mactxtsssr                     RegisterMactxtsssrType
	_                              [8]byte
	Macacr                         RegisterMacacrType
	_                              [4]byte
	Macatsnr                       RegisterMacatsnrType
	Macatssr                       RegisterMacatssrType
	Mactsiacr                      RegisterMactsiacrType
	Mactseacr                      RegisterMactseacrType
	Mactsicnr                      RegisterMactsicnrType
	Mactsecnr                      RegisterMactsecnrType
	_                              [16]byte
	Macppscr                       RegisterMacppscrType
	_                              [12]byte
	Macppsttsr                     RegisterMacppsttsrType
	Macppsttnr                     RegisterMacppsttnrType
	Macppsir                       RegisterMacppsirType
	Macppswr                       RegisterMacppswrType
	_                              [48]byte
	Macpocr                        RegisterMacpocrType
	Macspi0r                       RegisterMacspi0rType
	Macspi1r                       RegisterMacspi1rType
	Macspi2r                       RegisterMacspi2rType
	Maclmir                        RegisterMaclmirType
	_                              [44]byte
	Mtlomr                         RegisterMtlomrType
	_                              [28]byte
	Mtlisr                         RegisterMtlisrType
	_                              [220]byte
	Mtltxqomr                      RegisterMtltxqomrType
	Mtltxqur                       RegisterMtltxqurType
	Mtltxqdr                       RegisterMtltxqdrType
	_                              [32]byte
	Mtlqicsr                       RegisterMtlqicsrType
	Mtlrxqomr                      RegisterMtlrxqomrType
	Mtlrxqmpocr                    RegisterMtlrxqmpocrType
	Mtlrxqdr                       RegisterMtlrxqdrType
	_                              [708]byte
	Dmamr                          RegisterDmamrType
	Dmasbmr                        RegisterDmasbmrType
	Dmaisr                         RegisterDmaisrType
	Dmadsr                         RegisterDmadsrType
	_                              [240]byte
	Dmaccr                         RegisterDmaccrType
	Dmactxcr                       RegisterDmactxcrType
	Dmacrxcr                       RegisterDmacrxcrType
	_                              [8]byte
	Dmactxdlar                     RegisterDmactxdlarType
	_                              [4]byte
	Dmacrxdlar                     RegisterDmacrxdlarType
	Dmactxdtpr                     RegisterDmactxdtprType
	_                              [4]byte
	Dmacrxdtpr                     RegisterDmacrxdtprType
	Dmactxrlr                      RegisterDmactxrlrType
	Dmacrxrlr                      RegisterDmacrxrlrType
	Dmacier                        RegisterDmacierType
	Dmacrxiwtr                     RegisterDmacrxiwtrType
	_                              [8]byte
	Dmaccatxdr                     RegisterDmaccatxdrType
	_                              [4]byte
	Dmaccarxdr                     RegisterDmaccarxdrType
	_                              [4]byte
	Dmaccatxbr                     RegisterDmaccatxbrType
	_                              [4]byte
	Dmaccarxbr                     RegisterDmaccarxbrType
	Dmacsr                         RegisterDmacsrType
	_                              [8]byte
	Dmacmfcr                       RegisterDmacmfcrType
}

// RegisterMaccrType Operating mode configuration register
type RegisterMaccrType uint32

func (r *RegisterMaccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaccrFieldReShift = 0
	RegisterMaccrFieldReMask  = 0x1
)

// GetRe Receiver Enable
func (r *RegisterMaccrType) GetRe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldReMask) != 0
}

// SetRe Receiver Enable
func (r *RegisterMaccrType) SetRe(value bool) {
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
func (r *RegisterMaccrType) GetTe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldTeMask) != 0
}

// SetTe TE
func (r *RegisterMaccrType) SetTe(value bool) {
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
func (r *RegisterMaccrType) GetPrelen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldPrelenMask) >> RegisterMaccrFieldPrelenShift)
}

// SetPrelen PRELEN
func (r *RegisterMaccrType) SetPrelen(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldPrelenMask)|(uint32(value)<<RegisterMaccrFieldPrelenShift))
}

const (
	RegisterMaccrFieldDcShift = 4
	RegisterMaccrFieldDcMask  = 0x10
)

// GetDc DC
func (r *RegisterMaccrType) GetDc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldDcMask) != 0
}

// SetDc DC
func (r *RegisterMaccrType) SetDc(value bool) {
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
func (r *RegisterMaccrType) GetBl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldBlMask) >> RegisterMaccrFieldBlShift)
}

// SetBl BL
func (r *RegisterMaccrType) SetBl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldBlMask)|(uint32(value)<<RegisterMaccrFieldBlShift))
}

const (
	RegisterMaccrFieldDrShift = 8
	RegisterMaccrFieldDrMask  = 0x100
)

// GetDr DR
func (r *RegisterMaccrType) GetDr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldDrMask) != 0
}

// SetDr DR
func (r *RegisterMaccrType) SetDr(value bool) {
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
func (r *RegisterMaccrType) GetDcrs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldDcrsMask) != 0
}

// SetDcrs DCRS
func (r *RegisterMaccrType) SetDcrs(value bool) {
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
func (r *RegisterMaccrType) GetDo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldDoMask) != 0
}

// SetDo DO
func (r *RegisterMaccrType) SetDo(value bool) {
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
func (r *RegisterMaccrType) GetEcrsfd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldEcrsfdMask) != 0
}

// SetEcrsfd ECRSFD
func (r *RegisterMaccrType) SetEcrsfd(value bool) {
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
func (r *RegisterMaccrType) GetLm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldLmMask) != 0
}

// SetLm LM
func (r *RegisterMaccrType) SetLm(value bool) {
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
func (r *RegisterMaccrType) GetDm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldDmMask) != 0
}

// SetDm DM
func (r *RegisterMaccrType) SetDm(value bool) {
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
func (r *RegisterMaccrType) GetFes() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldFesMask) != 0
}

// SetFes FES
func (r *RegisterMaccrType) SetFes(value bool) {
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
func (r *RegisterMaccrType) GetJe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldJeMask) != 0
}

// SetJe JE
func (r *RegisterMaccrType) SetJe(value bool) {
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
func (r *RegisterMaccrType) GetJd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldJdMask) != 0
}

// SetJd JD
func (r *RegisterMaccrType) SetJd(value bool) {
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
func (r *RegisterMaccrType) GetWd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldWdMask) != 0
}

// SetWd WD
func (r *RegisterMaccrType) SetWd(value bool) {
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
func (r *RegisterMaccrType) GetAcs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldAcsMask) != 0
}

// SetAcs ACS
func (r *RegisterMaccrType) SetAcs(value bool) {
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
func (r *RegisterMaccrType) GetCst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldCstMask) != 0
}

// SetCst CST
func (r *RegisterMaccrType) SetCst(value bool) {
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
func (r *RegisterMaccrType) GetS2kp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldS2kpMask) != 0
}

// SetS2kp S2KP
func (r *RegisterMaccrType) SetS2kp(value bool) {
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
func (r *RegisterMaccrType) GetGpslce() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldGpslceMask) != 0
}

// SetGpslce GPSLCE
func (r *RegisterMaccrType) SetGpslce(value bool) {
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
func (r *RegisterMaccrType) GetIpg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldIpgMask) >> RegisterMaccrFieldIpgShift)
}

// SetIpg IPG
func (r *RegisterMaccrType) SetIpg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldIpgMask)|(uint32(value)<<RegisterMaccrFieldIpgShift))
}

const (
	RegisterMaccrFieldIpcShift = 27
	RegisterMaccrFieldIpcMask  = 0x8000000
)

// GetIpc IPC
func (r *RegisterMaccrType) GetIpc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldIpcMask) != 0
}

// SetIpc IPC
func (r *RegisterMaccrType) SetIpc(value bool) {
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
func (r *RegisterMaccrType) GetSarc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldSarcMask) >> RegisterMaccrFieldSarcShift)
}

// SetSarc SARC
func (r *RegisterMaccrType) SetSarc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldSarcMask)|(uint32(value)<<RegisterMaccrFieldSarcShift))
}

const (
	RegisterMaccrFieldArpenShift = 31
	RegisterMaccrFieldArpenMask  = 0x80000000
)

// GetArpen ARPEN
func (r *RegisterMaccrType) GetArpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaccrFieldArpenMask) != 0
}

// SetArpen ARPEN
func (r *RegisterMaccrType) SetArpen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaccrFieldArpenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaccrFieldArpenMask)
	}
}

// RegisterMacecrType Extended operating mode configuration register
type RegisterMacecrType uint32

func (r *RegisterMacecrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacecrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacecrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacecrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacecrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacecrFieldGpslShift = 0
	RegisterMacecrFieldGpslMask  = 0x3fff
)

// GetGpsl GPSL
func (r *RegisterMacecrType) GetGpsl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldGpslMask) >> RegisterMacecrFieldGpslShift)
}

// SetGpsl GPSL
func (r *RegisterMacecrType) SetGpsl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacecrFieldGpslMask)|(uint32(value)<<RegisterMacecrFieldGpslShift))
}

const (
	RegisterMacecrFieldDcrccShift = 16
	RegisterMacecrFieldDcrccMask  = 0x10000
)

// GetDcrcc DCRCC
func (r *RegisterMacecrType) GetDcrcc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldDcrccMask) != 0
}

// SetDcrcc DCRCC
func (r *RegisterMacecrType) SetDcrcc(value bool) {
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
func (r *RegisterMacecrType) GetSpen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldSpenMask) != 0
}

// SetSpen SPEN
func (r *RegisterMacecrType) SetSpen(value bool) {
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
func (r *RegisterMacecrType) GetUsp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldUspMask) != 0
}

// SetUsp USP
func (r *RegisterMacecrType) SetUsp(value bool) {
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
func (r *RegisterMacecrType) GetEipgen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldEipgenMask) != 0
}

// SetEipgen EIPGEN
func (r *RegisterMacecrType) SetEipgen(value bool) {
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
func (r *RegisterMacecrType) GetEipg() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacecrFieldEipgMask) >> RegisterMacecrFieldEipgShift)
}

// SetEipg EIPG
func (r *RegisterMacecrType) SetEipg(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacecrFieldEipgMask)|(uint32(value)<<RegisterMacecrFieldEipgShift))
}

// RegisterMacpfrType Packet filtering control register
type RegisterMacpfrType uint32

func (r *RegisterMacpfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacpfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacpfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacpfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacpfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacpfrFieldPrShift = 0
	RegisterMacpfrFieldPrMask  = 0x1
)

// GetPr PR
func (r *RegisterMacpfrType) GetPr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldPrMask) != 0
}

// SetPr PR
func (r *RegisterMacpfrType) SetPr(value bool) {
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
func (r *RegisterMacpfrType) GetHuc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldHucMask) != 0
}

// SetHuc HUC
func (r *RegisterMacpfrType) SetHuc(value bool) {
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
func (r *RegisterMacpfrType) GetHmc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldHmcMask) != 0
}

// SetHmc HMC
func (r *RegisterMacpfrType) SetHmc(value bool) {
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
func (r *RegisterMacpfrType) GetDaif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldDaifMask) != 0
}

// SetDaif DAIF
func (r *RegisterMacpfrType) SetDaif(value bool) {
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
func (r *RegisterMacpfrType) GetPm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldPmMask) != 0
}

// SetPm PM
func (r *RegisterMacpfrType) SetPm(value bool) {
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
func (r *RegisterMacpfrType) GetDbf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldDbfMask) != 0
}

// SetDbf DBF
func (r *RegisterMacpfrType) SetDbf(value bool) {
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
func (r *RegisterMacpfrType) GetPcf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldPcfMask) >> RegisterMacpfrFieldPcfShift)
}

// SetPcf PCF
func (r *RegisterMacpfrType) SetPcf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldPcfMask)|(uint32(value)<<RegisterMacpfrFieldPcfShift))
}

const (
	RegisterMacpfrFieldSaifShift = 8
	RegisterMacpfrFieldSaifMask  = 0x100
)

// GetSaif SAIF
func (r *RegisterMacpfrType) GetSaif() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldSaifMask) != 0
}

// SetSaif SAIF
func (r *RegisterMacpfrType) SetSaif(value bool) {
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
func (r *RegisterMacpfrType) GetSaf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldSafMask) != 0
}

// SetSaf SAF
func (r *RegisterMacpfrType) SetSaf(value bool) {
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
func (r *RegisterMacpfrType) GetHpf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldHpfMask) != 0
}

// SetHpf HPF
func (r *RegisterMacpfrType) SetHpf(value bool) {
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
func (r *RegisterMacpfrType) GetVtfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldVtfeMask) != 0
}

// SetVtfe VTFE
func (r *RegisterMacpfrType) SetVtfe(value bool) {
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
func (r *RegisterMacpfrType) GetIpfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldIpfeMask) != 0
}

// SetIpfe IPFE
func (r *RegisterMacpfrType) SetIpfe(value bool) {
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
func (r *RegisterMacpfrType) GetDntu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldDntuMask) != 0
}

// SetDntu DNTU
func (r *RegisterMacpfrType) SetDntu(value bool) {
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
func (r *RegisterMacpfrType) GetRa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpfrFieldRaMask) != 0
}

// SetRa RA
func (r *RegisterMacpfrType) SetRa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpfrFieldRaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpfrFieldRaMask)
	}
}

// RegisterMacwtrType Watchdog timeout register
type RegisterMacwtrType uint32

func (r *RegisterMacwtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacwtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacwtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacwtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacwtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacwtrFieldWtoShift = 0
	RegisterMacwtrFieldWtoMask  = 0xf
)

// GetWto WTO
func (r *RegisterMacwtrType) GetWto() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacwtrFieldWtoMask) >> RegisterMacwtrFieldWtoShift)
}

// SetWto WTO
func (r *RegisterMacwtrType) SetWto(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacwtrFieldWtoMask)|(uint32(value)<<RegisterMacwtrFieldWtoShift))
}

const (
	RegisterMacwtrFieldPweShift = 8
	RegisterMacwtrFieldPweMask  = 0x100
)

// GetPwe PWE
func (r *RegisterMacwtrType) GetPwe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacwtrFieldPweMask) != 0
}

// SetPwe PWE
func (r *RegisterMacwtrType) SetPwe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacwtrFieldPweMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacwtrFieldPweMask)
	}
}

// RegisterMacht0rType Hash Table 0 register
type RegisterMacht0rType uint32

func (r *RegisterMacht0rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacht0rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacht0rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacht0rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacht0rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacht0rFieldHt31t0Shift = 0
	RegisterMacht0rFieldHt31t0Mask  = 0xffffffff
)

// GetHt31t0 HT31T0
func (r *RegisterMacht0rType) GetHt31t0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacht0rFieldHt31t0Mask) >> RegisterMacht0rFieldHt31t0Shift)
}

// SetHt31t0 HT31T0
func (r *RegisterMacht0rType) SetHt31t0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacht0rFieldHt31t0Mask)|(uint32(value)<<RegisterMacht0rFieldHt31t0Shift))
}

// RegisterMacht1rType Hash Table 1 register
type RegisterMacht1rType uint32

func (r *RegisterMacht1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacht1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacht1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacht1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacht1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacht1rFieldHt63t32Shift = 0
	RegisterMacht1rFieldHt63t32Mask  = 0xffffffff
)

// GetHt63t32 HT63T32
func (r *RegisterMacht1rType) GetHt63t32() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacht1rFieldHt63t32Mask) >> RegisterMacht1rFieldHt63t32Shift)
}

// SetHt63t32 HT63T32
func (r *RegisterMacht1rType) SetHt63t32(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacht1rFieldHt63t32Mask)|(uint32(value)<<RegisterMacht1rFieldHt63t32Shift))
}

// RegisterMacvtrType VLAN tag register
type RegisterMacvtrType uint32

func (r *RegisterMacvtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacvtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacvtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacvtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacvtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacvtrFieldVlShift = 0
	RegisterMacvtrFieldVlMask  = 0xffff
)

// GetVl VL
func (r *RegisterMacvtrType) GetVl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldVlMask) >> RegisterMacvtrFieldVlShift)
}

// SetVl VL
func (r *RegisterMacvtrType) SetVl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldVlMask)|(uint32(value)<<RegisterMacvtrFieldVlShift))
}

const (
	RegisterMacvtrFieldEtvShift = 16
	RegisterMacvtrFieldEtvMask  = 0x10000
)

// GetEtv ETV
func (r *RegisterMacvtrType) GetEtv() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEtvMask) != 0
}

// SetEtv ETV
func (r *RegisterMacvtrType) SetEtv(value bool) {
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
func (r *RegisterMacvtrType) GetVtim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldVtimMask) != 0
}

// SetVtim VTIM
func (r *RegisterMacvtrType) SetVtim(value bool) {
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
func (r *RegisterMacvtrType) GetEsvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEsvlMask) != 0
}

// SetEsvl ESVL
func (r *RegisterMacvtrType) SetEsvl(value bool) {
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
func (r *RegisterMacvtrType) GetErsvlm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldErsvlmMask) != 0
}

// SetErsvlm ERSVLM
func (r *RegisterMacvtrType) SetErsvlm(value bool) {
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
func (r *RegisterMacvtrType) GetDovltc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldDovltcMask) != 0
}

// SetDovltc DOVLTC
func (r *RegisterMacvtrType) SetDovltc(value bool) {
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
func (r *RegisterMacvtrType) GetEvls() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEvlsMask) >> RegisterMacvtrFieldEvlsShift)
}

// SetEvls EVLS
func (r *RegisterMacvtrType) SetEvls(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldEvlsMask)|(uint32(value)<<RegisterMacvtrFieldEvlsShift))
}

const (
	RegisterMacvtrFieldEvlrxsShift = 24
	RegisterMacvtrFieldEvlrxsMask  = 0x1000000
)

// GetEvlrxs EVLRXS
func (r *RegisterMacvtrType) GetEvlrxs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEvlrxsMask) != 0
}

// SetEvlrxs EVLRXS
func (r *RegisterMacvtrType) SetEvlrxs(value bool) {
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
func (r *RegisterMacvtrType) GetVthm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldVthmMask) != 0
}

// SetVthm VTHM
func (r *RegisterMacvtrType) SetVthm(value bool) {
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
func (r *RegisterMacvtrType) GetEdvlp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEdvlpMask) != 0
}

// SetEdvlp EDVLP
func (r *RegisterMacvtrType) SetEdvlp(value bool) {
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
func (r *RegisterMacvtrType) GetErivlt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldErivltMask) != 0
}

// SetErivlt ERIVLT
func (r *RegisterMacvtrType) SetErivlt(value bool) {
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
func (r *RegisterMacvtrType) GetEivls() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEivlsMask) >> RegisterMacvtrFieldEivlsShift)
}

// SetEivls EIVLS
func (r *RegisterMacvtrType) SetEivls(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldEivlsMask)|(uint32(value)<<RegisterMacvtrFieldEivlsShift))
}

const (
	RegisterMacvtrFieldEivlrxsShift = 31
	RegisterMacvtrFieldEivlrxsMask  = 0x80000000
)

// GetEivlrxs EIVLRXS
func (r *RegisterMacvtrType) GetEivlrxs() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvtrFieldEivlrxsMask) != 0
}

// SetEivlrxs EIVLRXS
func (r *RegisterMacvtrType) SetEivlrxs(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvtrFieldEivlrxsMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvtrFieldEivlrxsMask)
	}
}

// RegisterMacvhtrType VLAN Hash table register
type RegisterMacvhtrType uint32

func (r *RegisterMacvhtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacvhtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacvhtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacvhtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacvhtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacvhtrFieldVlhtShift = 0
	RegisterMacvhtrFieldVlhtMask  = 0xffff
)

// GetVlht VLHT
func (r *RegisterMacvhtrType) GetVlht() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacvhtrFieldVlhtMask) >> RegisterMacvhtrFieldVlhtShift)
}

// SetVlht VLHT
func (r *RegisterMacvhtrType) SetVlht(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvhtrFieldVlhtMask)|(uint32(value)<<RegisterMacvhtrFieldVlhtShift))
}

// RegisterMacvirType VLAN inclusion register
type RegisterMacvirType uint32

func (r *RegisterMacvirType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacvirType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacvirType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacvirType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacvirType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacvirFieldVltShift = 0
	RegisterMacvirFieldVltMask  = 0xffff
)

// GetVlt VLT
func (r *RegisterMacvirType) GetVlt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacvirFieldVltMask) >> RegisterMacvirFieldVltShift)
}

// SetVlt VLT
func (r *RegisterMacvirType) SetVlt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvirFieldVltMask)|(uint32(value)<<RegisterMacvirFieldVltShift))
}

const (
	RegisterMacvirFieldVlcShift = 16
	RegisterMacvirFieldVlcMask  = 0x30000
)

// GetVlc VLC
func (r *RegisterMacvirType) GetVlc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacvirFieldVlcMask) >> RegisterMacvirFieldVlcShift)
}

// SetVlc VLC
func (r *RegisterMacvirType) SetVlc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvirFieldVlcMask)|(uint32(value)<<RegisterMacvirFieldVlcShift))
}

const (
	RegisterMacvirFieldVlpShift = 18
	RegisterMacvirFieldVlpMask  = 0x40000
)

// GetVlp VLP
func (r *RegisterMacvirType) GetVlp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvirFieldVlpMask) != 0
}

// SetVlp VLP
func (r *RegisterMacvirType) SetVlp(value bool) {
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
func (r *RegisterMacvirType) GetCsvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvirFieldCsvlMask) != 0
}

// SetCsvl CSVL
func (r *RegisterMacvirType) SetCsvl(value bool) {
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
func (r *RegisterMacvirType) GetVlti() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacvirFieldVltiMask) != 0
}

// SetVlti VLTI
func (r *RegisterMacvirType) SetVlti(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacvirFieldVltiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacvirFieldVltiMask)
	}
}

// RegisterMacivirType Inner VLAN inclusion register
type RegisterMacivirType uint32

func (r *RegisterMacivirType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacivirType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacivirType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacivirType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacivirType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacivirFieldVltShift = 0
	RegisterMacivirFieldVltMask  = 0xffff
)

// GetVlt VLT
func (r *RegisterMacivirType) GetVlt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacivirFieldVltMask) >> RegisterMacivirFieldVltShift)
}

// SetVlt VLT
func (r *RegisterMacivirType) SetVlt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacivirFieldVltMask)|(uint32(value)<<RegisterMacivirFieldVltShift))
}

const (
	RegisterMacivirFieldVlcShift = 16
	RegisterMacivirFieldVlcMask  = 0x30000
)

// GetVlc VLC
func (r *RegisterMacivirType) GetVlc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacivirFieldVlcMask) >> RegisterMacivirFieldVlcShift)
}

// SetVlc VLC
func (r *RegisterMacivirType) SetVlc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacivirFieldVlcMask)|(uint32(value)<<RegisterMacivirFieldVlcShift))
}

const (
	RegisterMacivirFieldVlpShift = 18
	RegisterMacivirFieldVlpMask  = 0x40000
)

// GetVlp VLP
func (r *RegisterMacivirType) GetVlp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacivirFieldVlpMask) != 0
}

// SetVlp VLP
func (r *RegisterMacivirType) SetVlp(value bool) {
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
func (r *RegisterMacivirType) GetCsvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacivirFieldCsvlMask) != 0
}

// SetCsvl CSVL
func (r *RegisterMacivirType) SetCsvl(value bool) {
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
func (r *RegisterMacivirType) GetVlti() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacivirFieldVltiMask) != 0
}

// SetVlti VLTI
func (r *RegisterMacivirType) SetVlti(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacivirFieldVltiMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacivirFieldVltiMask)
	}
}

// RegisterMacqtxfcrType Tx Queue flow control register
type RegisterMacqtxfcrType uint32

func (r *RegisterMacqtxfcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacqtxfcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacqtxfcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacqtxfcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacqtxfcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacqtxfcrFieldFcbbpaShift = 0
	RegisterMacqtxfcrFieldFcbbpaMask  = 0x1
)

// GetFcbbpa FCB_BPA
func (r *RegisterMacqtxfcrType) GetFcbbpa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacqtxfcrFieldFcbbpaMask) != 0
}

// SetFcbbpa FCB_BPA
func (r *RegisterMacqtxfcrType) SetFcbbpa(value bool) {
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
func (r *RegisterMacqtxfcrType) GetTfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacqtxfcrFieldTfeMask) != 0
}

// SetTfe TFE
func (r *RegisterMacqtxfcrType) SetTfe(value bool) {
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
func (r *RegisterMacqtxfcrType) GetPlt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacqtxfcrFieldPltMask) >> RegisterMacqtxfcrFieldPltShift)
}

// SetPlt PLT
func (r *RegisterMacqtxfcrType) SetPlt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacqtxfcrFieldPltMask)|(uint32(value)<<RegisterMacqtxfcrFieldPltShift))
}

const (
	RegisterMacqtxfcrFieldDzpqShift = 7
	RegisterMacqtxfcrFieldDzpqMask  = 0x80
)

// GetDzpq DZPQ
func (r *RegisterMacqtxfcrType) GetDzpq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacqtxfcrFieldDzpqMask) != 0
}

// SetDzpq DZPQ
func (r *RegisterMacqtxfcrType) SetDzpq(value bool) {
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
func (r *RegisterMacqtxfcrType) GetPt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacqtxfcrFieldPtMask) >> RegisterMacqtxfcrFieldPtShift)
}

// SetPt PT
func (r *RegisterMacqtxfcrType) SetPt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacqtxfcrFieldPtMask)|(uint32(value)<<RegisterMacqtxfcrFieldPtShift))
}

// RegisterMacrxfcrType Rx flow control register
type RegisterMacrxfcrType uint32

func (r *RegisterMacrxfcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacrxfcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacrxfcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacrxfcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacrxfcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacrxfcrFieldRfeShift = 0
	RegisterMacrxfcrFieldRfeMask  = 0x1
)

// GetRfe RFE
func (r *RegisterMacrxfcrType) GetRfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxfcrFieldRfeMask) != 0
}

// SetRfe RFE
func (r *RegisterMacrxfcrType) SetRfe(value bool) {
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
func (r *RegisterMacrxfcrType) GetUp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxfcrFieldUpMask) != 0
}

// SetUp UP
func (r *RegisterMacrxfcrType) SetUp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacrxfcrFieldUpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacrxfcrFieldUpMask)
	}
}

// RegisterMacisrType Interrupt status register
type RegisterMacisrType uint32

func (r *RegisterMacisrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacisrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacisrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacisrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacisrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacisrFieldPhyisShift = 3
	RegisterMacisrFieldPhyisMask  = 0x8
)

// GetPhyis PHYIS
func (r *RegisterMacisrType) GetPhyis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldPhyisMask) != 0
}

// SetPhyis PHYIS
func (r *RegisterMacisrType) SetPhyis(value bool) {
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
func (r *RegisterMacisrType) GetPmtis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldPmtisMask) != 0
}

// SetPmtis PMTIS
func (r *RegisterMacisrType) SetPmtis(value bool) {
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
func (r *RegisterMacisrType) GetLpiis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldLpiisMask) != 0
}

// SetLpiis LPIIS
func (r *RegisterMacisrType) SetLpiis(value bool) {
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
func (r *RegisterMacisrType) GetMmcis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldMmcisMask) != 0
}

// SetMmcis MMCIS
func (r *RegisterMacisrType) SetMmcis(value bool) {
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
func (r *RegisterMacisrType) GetMmcrxis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldMmcrxisMask) != 0
}

// SetMmcrxis MMCRXIS
func (r *RegisterMacisrType) SetMmcrxis(value bool) {
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
func (r *RegisterMacisrType) GetMmctxis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldMmctxisMask) != 0
}

// SetMmctxis MMCTXIS
func (r *RegisterMacisrType) SetMmctxis(value bool) {
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
func (r *RegisterMacisrType) GetTsis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldTsisMask) != 0
}

// SetTsis TSIS
func (r *RegisterMacisrType) SetTsis(value bool) {
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
func (r *RegisterMacisrType) GetTxstsis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldTxstsisMask) != 0
}

// SetTxstsis TXSTSIS
func (r *RegisterMacisrType) SetTxstsis(value bool) {
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
func (r *RegisterMacisrType) GetRxstsis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacisrFieldRxstsisMask) != 0
}

// SetRxstsis RXSTSIS
func (r *RegisterMacisrType) SetRxstsis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacisrFieldRxstsisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacisrFieldRxstsisMask)
	}
}

// RegisterMacierType Interrupt enable register
type RegisterMacierType uint32

func (r *RegisterMacierType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacierType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacierType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacierType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacierType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacierFieldPhyieShift = 3
	RegisterMacierFieldPhyieMask  = 0x8
)

// GetPhyie PHYIE
func (r *RegisterMacierType) GetPhyie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldPhyieMask) != 0
}

// SetPhyie PHYIE
func (r *RegisterMacierType) SetPhyie(value bool) {
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
func (r *RegisterMacierType) GetPmtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldPmtieMask) != 0
}

// SetPmtie PMTIE
func (r *RegisterMacierType) SetPmtie(value bool) {
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
func (r *RegisterMacierType) GetLpiie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldLpiieMask) != 0
}

// SetLpiie LPIIE
func (r *RegisterMacierType) SetLpiie(value bool) {
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
func (r *RegisterMacierType) GetTsie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldTsieMask) != 0
}

// SetTsie TSIE
func (r *RegisterMacierType) SetTsie(value bool) {
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
func (r *RegisterMacierType) GetTxstsie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldTxstsieMask) != 0
}

// SetTxstsie TXSTSIE
func (r *RegisterMacierType) SetTxstsie(value bool) {
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
func (r *RegisterMacierType) GetRxstsie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacierFieldRxstsieMask) != 0
}

// SetRxstsie RXSTSIE
func (r *RegisterMacierType) SetRxstsie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacierFieldRxstsieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacierFieldRxstsieMask)
	}
}

// RegisterMacrxtxsrType Rx Tx status register
type RegisterMacrxtxsrType uint32

func (r *RegisterMacrxtxsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacrxtxsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacrxtxsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacrxtxsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacrxtxsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacrxtxsrFieldTjtShift = 0
	RegisterMacrxtxsrFieldTjtMask  = 0x1
)

// GetTjt TJT
func (r *RegisterMacrxtxsrType) GetTjt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldTjtMask) != 0
}

// SetTjt TJT
func (r *RegisterMacrxtxsrType) SetTjt(value bool) {
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
func (r *RegisterMacrxtxsrType) GetNcarr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldNcarrMask) != 0
}

// SetNcarr NCARR
func (r *RegisterMacrxtxsrType) SetNcarr(value bool) {
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
func (r *RegisterMacrxtxsrType) GetLcarr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldLcarrMask) != 0
}

// SetLcarr LCARR
func (r *RegisterMacrxtxsrType) SetLcarr(value bool) {
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
func (r *RegisterMacrxtxsrType) GetExdef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldExdefMask) != 0
}

// SetExdef EXDEF
func (r *RegisterMacrxtxsrType) SetExdef(value bool) {
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
func (r *RegisterMacrxtxsrType) GetLcol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldLcolMask) != 0
}

// SetLcol LCOL
func (r *RegisterMacrxtxsrType) SetLcol(value bool) {
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
func (r *RegisterMacrxtxsrType) GetExcol() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldExcolMask) != 0
}

// SetExcol LCOL
func (r *RegisterMacrxtxsrType) SetExcol(value bool) {
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
func (r *RegisterMacrxtxsrType) GetRwt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacrxtxsrFieldRwtMask) != 0
}

// SetRwt RWT
func (r *RegisterMacrxtxsrType) SetRwt(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacrxtxsrFieldRwtMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacrxtxsrFieldRwtMask)
	}
}

// RegisterMacpcsrType PMT control status register
type RegisterMacpcsrType uint32

func (r *RegisterMacpcsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacpcsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacpcsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacpcsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacpcsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacpcsrFieldPwrdwnShift = 0
	RegisterMacpcsrFieldPwrdwnMask  = 0x1
)

// GetPwrdwn PWRDWN
func (r *RegisterMacpcsrType) GetPwrdwn() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldPwrdwnMask) != 0
}

// SetPwrdwn PWRDWN
func (r *RegisterMacpcsrType) SetPwrdwn(value bool) {
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
func (r *RegisterMacpcsrType) GetMgkpkten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldMgkpktenMask) != 0
}

// SetMgkpkten MGKPKTEN
func (r *RegisterMacpcsrType) SetMgkpkten(value bool) {
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
func (r *RegisterMacpcsrType) GetRwkpkten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldRwkpktenMask) != 0
}

// SetRwkpkten RWKPKTEN
func (r *RegisterMacpcsrType) SetRwkpkten(value bool) {
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
func (r *RegisterMacpcsrType) GetMgkprcvd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldMgkprcvdMask) != 0
}

const (
	RegisterMacpcsrFieldRwkprcvdShift = 6
	RegisterMacpcsrFieldRwkprcvdMask  = 0x40
)

// GetRwkprcvd RWKPRCVD
func (r *RegisterMacpcsrType) GetRwkprcvd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldRwkprcvdMask) != 0
}

const (
	RegisterMacpcsrFieldGlblucastShift = 9
	RegisterMacpcsrFieldGlblucastMask  = 0x200
)

// GetGlblucast GLBLUCAST
func (r *RegisterMacpcsrType) GetGlblucast() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldGlblucastMask) != 0
}

// SetGlblucast GLBLUCAST
func (r *RegisterMacpcsrType) SetGlblucast(value bool) {
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
func (r *RegisterMacpcsrType) GetRwkpfe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldRwkpfeMask) != 0
}

// SetRwkpfe RWKPFE
func (r *RegisterMacpcsrType) SetRwkpfe(value bool) {
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
func (r *RegisterMacpcsrType) GetRwkptr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldRwkptrMask) >> RegisterMacpcsrFieldRwkptrShift)
}

// SetRwkptr RWKPTR
func (r *RegisterMacpcsrType) SetRwkptr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacpcsrFieldRwkptrMask)|(uint32(value)<<RegisterMacpcsrFieldRwkptrShift))
}

const (
	RegisterMacpcsrFieldRwkfiltrstShift = 31
	RegisterMacpcsrFieldRwkfiltrstMask  = 0x80000000
)

// GetRwkfiltrst RWKFILTRST
func (r *RegisterMacpcsrType) GetRwkfiltrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpcsrFieldRwkfiltrstMask) != 0
}

// SetRwkfiltrst RWKFILTRST
func (r *RegisterMacpcsrType) SetRwkfiltrst(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacpcsrFieldRwkfiltrstMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacpcsrFieldRwkfiltrstMask)
	}
}

// RegisterMacrwkpfrType Remove wakeup packet filter register
type RegisterMacrwkpfrType uint32

func (r *RegisterMacrwkpfrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacrwkpfrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacrwkpfrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacrwkpfrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacrwkpfrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacrwkpfrFieldWkupfrmftrShift = 0
	RegisterMacrwkpfrFieldWkupfrmftrMask  = 0xffffffff
)

// GetWkupfrmftr WKUPFRMFTR
func (r *RegisterMacrwkpfrType) GetWkupfrmftr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacrwkpfrFieldWkupfrmftrMask) >> RegisterMacrwkpfrFieldWkupfrmftrShift)
}

// SetWkupfrmftr WKUPFRMFTR
func (r *RegisterMacrwkpfrType) SetWkupfrmftr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacrwkpfrFieldWkupfrmftrMask)|(uint32(value)<<RegisterMacrwkpfrFieldWkupfrmftrShift))
}

// RegisterMaclcsrType LPI control status register
type RegisterMaclcsrType uint32

func (r *RegisterMaclcsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaclcsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaclcsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaclcsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaclcsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaclcsrFieldTlpienShift = 0
	RegisterMaclcsrFieldTlpienMask  = 0x1
)

// GetTlpien TLPIEN
func (r *RegisterMaclcsrType) GetTlpien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldTlpienMask) != 0
}

const (
	RegisterMaclcsrFieldTlpiexShift = 1
	RegisterMaclcsrFieldTlpiexMask  = 0x2
)

// GetTlpiex TLPIEX
func (r *RegisterMaclcsrType) GetTlpiex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldTlpiexMask) != 0
}

const (
	RegisterMaclcsrFieldRlpienShift = 2
	RegisterMaclcsrFieldRlpienMask  = 0x4
)

// GetRlpien RLPIEN
func (r *RegisterMaclcsrType) GetRlpien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldRlpienMask) != 0
}

const (
	RegisterMaclcsrFieldRlpiexShift = 3
	RegisterMaclcsrFieldRlpiexMask  = 0x8
)

// GetRlpiex RLPIEX
func (r *RegisterMaclcsrType) GetRlpiex() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldRlpiexMask) != 0
}

const (
	RegisterMaclcsrFieldTlpistShift = 8
	RegisterMaclcsrFieldTlpistMask  = 0x100
)

// GetTlpist TLPIST
func (r *RegisterMaclcsrType) GetTlpist() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldTlpistMask) != 0
}

const (
	RegisterMaclcsrFieldRlpistShift = 9
	RegisterMaclcsrFieldRlpistMask  = 0x200
)

// GetRlpist RLPIST
func (r *RegisterMaclcsrType) GetRlpist() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldRlpistMask) != 0
}

const (
	RegisterMaclcsrFieldLpienShift = 16
	RegisterMaclcsrFieldLpienMask  = 0x10000
)

// GetLpien LPIEN
func (r *RegisterMaclcsrType) GetLpien() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldLpienMask) != 0
}

// SetLpien LPIEN
func (r *RegisterMaclcsrType) SetLpien(value bool) {
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
func (r *RegisterMaclcsrType) GetPls() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldPlsMask) != 0
}

// SetPls PLS
func (r *RegisterMaclcsrType) SetPls(value bool) {
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
func (r *RegisterMaclcsrType) GetPlsen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldPlsenMask) != 0
}

// SetPlsen PLSEN
func (r *RegisterMaclcsrType) SetPlsen(value bool) {
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
func (r *RegisterMaclcsrType) GetLpitxa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldLpitxaMask) != 0
}

// SetLpitxa LPITXA
func (r *RegisterMaclcsrType) SetLpitxa(value bool) {
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
func (r *RegisterMaclcsrType) GetLpite() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldLpiteMask) != 0
}

// SetLpite LPITE
func (r *RegisterMaclcsrType) SetLpite(value bool) {
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
func (r *RegisterMaclcsrType) GetLpitcse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaclcsrFieldLpitcseMask) != 0
}

// SetLpitcse LPITCSE
func (r *RegisterMaclcsrType) SetLpitcse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaclcsrFieldLpitcseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaclcsrFieldLpitcseMask)
	}
}

// RegisterMacltcrType LPI timers control register
type RegisterMacltcrType uint32

func (r *RegisterMacltcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacltcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacltcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacltcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacltcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacltcrFieldTwtShift = 0
	RegisterMacltcrFieldTwtMask  = 0xffff
)

// GetTwt TWT
func (r *RegisterMacltcrType) GetTwt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacltcrFieldTwtMask) >> RegisterMacltcrFieldTwtShift)
}

// SetTwt TWT
func (r *RegisterMacltcrType) SetTwt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacltcrFieldTwtMask)|(uint32(value)<<RegisterMacltcrFieldTwtShift))
}

const (
	RegisterMacltcrFieldLstShift = 16
	RegisterMacltcrFieldLstMask  = 0x3ff0000
)

// GetLst LST
func (r *RegisterMacltcrType) GetLst() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacltcrFieldLstMask) >> RegisterMacltcrFieldLstShift)
}

// SetLst LST
func (r *RegisterMacltcrType) SetLst(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacltcrFieldLstMask)|(uint32(value)<<RegisterMacltcrFieldLstShift))
}

// RegisterMacletrType LPI entry timer register
type RegisterMacletrType uint32

func (r *RegisterMacletrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacletrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacletrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacletrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacletrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacletrFieldLpietShift = 0
	RegisterMacletrFieldLpietMask  = 0x1ffff
)

// GetLpiet LPIET
func (r *RegisterMacletrType) GetLpiet() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacletrFieldLpietMask) >> RegisterMacletrFieldLpietShift)
}

// SetLpiet LPIET
func (r *RegisterMacletrType) SetLpiet(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacletrFieldLpietMask)|(uint32(value)<<RegisterMacletrFieldLpietShift))
}

// RegisterMac1ustcrType 1-microsecond-tick counter register
type RegisterMac1ustcrType uint32

func (r *RegisterMac1ustcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMac1ustcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMac1ustcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMac1ustcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMac1ustcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMac1ustcrFieldTic1uscntrShift = 0
	RegisterMac1ustcrFieldTic1uscntrMask  = 0xfff
)

// GetTic1uscntr TIC_1US_CNTR
func (r *RegisterMac1ustcrType) GetTic1uscntr() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMac1ustcrFieldTic1uscntrMask) >> RegisterMac1ustcrFieldTic1uscntrShift)
}

// SetTic1uscntr TIC_1US_CNTR
func (r *RegisterMac1ustcrType) SetTic1uscntr(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMac1ustcrFieldTic1uscntrMask)|(uint32(value)<<RegisterMac1ustcrFieldTic1uscntrShift))
}

// RegisterMacvrType Version register
type RegisterMacvrType uint32

func (r *RegisterMacvrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacvrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacvrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacvrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacvrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacvrFieldSnpsverShift = 0
	RegisterMacvrFieldSnpsverMask  = 0xff
)

// GetSnpsver SNPSVER
func (r *RegisterMacvrType) GetSnpsver() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacvrFieldSnpsverMask) >> RegisterMacvrFieldSnpsverShift)
}

// SetSnpsver SNPSVER
func (r *RegisterMacvrType) SetSnpsver(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvrFieldSnpsverMask)|(uint32(value)<<RegisterMacvrFieldSnpsverShift))
}

const (
	RegisterMacvrFieldUserverShift = 8
	RegisterMacvrFieldUserverMask  = 0xff00
)

// GetUserver USERVER
func (r *RegisterMacvrType) GetUserver() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacvrFieldUserverMask) >> RegisterMacvrFieldUserverShift)
}

// SetUserver USERVER
func (r *RegisterMacvrType) SetUserver(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacvrFieldUserverMask)|(uint32(value)<<RegisterMacvrFieldUserverShift))
}

// RegisterMacdrType Debug register
type RegisterMacdrType uint32

func (r *RegisterMacdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacdrFieldRpestsShift = 0
	RegisterMacdrFieldRpestsMask  = 0x1
)

// GetRpests RPESTS
func (r *RegisterMacdrType) GetRpests() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacdrFieldRpestsMask) != 0
}

// SetRpests RPESTS
func (r *RegisterMacdrType) SetRpests(value bool) {
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
func (r *RegisterMacdrType) GetRfcfcsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacdrFieldRfcfcstsMask) >> RegisterMacdrFieldRfcfcstsShift)
}

// SetRfcfcsts RFCFCSTS
func (r *RegisterMacdrType) SetRfcfcsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacdrFieldRfcfcstsMask)|(uint32(value)<<RegisterMacdrFieldRfcfcstsShift))
}

const (
	RegisterMacdrFieldTpestsShift = 16
	RegisterMacdrFieldTpestsMask  = 0x10000
)

// GetTpests TPESTS
func (r *RegisterMacdrType) GetTpests() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacdrFieldTpestsMask) != 0
}

// SetTpests TPESTS
func (r *RegisterMacdrType) SetTpests(value bool) {
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
func (r *RegisterMacdrType) GetTfcsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacdrFieldTfcstsMask) >> RegisterMacdrFieldTfcstsShift)
}

// SetTfcsts TFCSTS
func (r *RegisterMacdrType) SetTfcsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacdrFieldTfcstsMask)|(uint32(value)<<RegisterMacdrFieldTfcstsShift))
}

// RegisterMachwf1rType HW feature 1 register
type RegisterMachwf1rType uint32

func (r *RegisterMachwf1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMachwf1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMachwf1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMachwf1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMachwf1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMachwf1rFieldRxfifosizeShift = 0
	RegisterMachwf1rFieldRxfifosizeMask  = 0x1f
)

// GetRxfifosize RXFIFOSIZE
func (r *RegisterMachwf1rType) GetRxfifosize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldRxfifosizeMask) >> RegisterMachwf1rFieldRxfifosizeShift)
}

// SetRxfifosize RXFIFOSIZE
func (r *RegisterMachwf1rType) SetRxfifosize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldRxfifosizeMask)|(uint32(value)<<RegisterMachwf1rFieldRxfifosizeShift))
}

const (
	RegisterMachwf1rFieldTxfifosizeShift = 6
	RegisterMachwf1rFieldTxfifosizeMask  = 0x7c0
)

// GetTxfifosize TXFIFOSIZE
func (r *RegisterMachwf1rType) GetTxfifosize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldTxfifosizeMask) >> RegisterMachwf1rFieldTxfifosizeShift)
}

// SetTxfifosize TXFIFOSIZE
func (r *RegisterMachwf1rType) SetTxfifosize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldTxfifosizeMask)|(uint32(value)<<RegisterMachwf1rFieldTxfifosizeShift))
}

const (
	RegisterMachwf1rFieldOstenShift = 11
	RegisterMachwf1rFieldOstenMask  = 0x800
)

// GetOsten OSTEN
func (r *RegisterMachwf1rType) GetOsten() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldOstenMask) != 0
}

// SetOsten OSTEN
func (r *RegisterMachwf1rType) SetOsten(value bool) {
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
func (r *RegisterMachwf1rType) GetPtoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldPtoenMask) != 0
}

// SetPtoen PTOEN
func (r *RegisterMachwf1rType) SetPtoen(value bool) {
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
func (r *RegisterMachwf1rType) GetAdvthword() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldAdvthwordMask) != 0
}

// SetAdvthword ADVTHWORD
func (r *RegisterMachwf1rType) SetAdvthword(value bool) {
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
func (r *RegisterMachwf1rType) GetDcben() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldDcbenMask) != 0
}

// SetDcben DCBEN
func (r *RegisterMachwf1rType) SetDcben(value bool) {
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
func (r *RegisterMachwf1rType) GetSphen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldSphenMask) != 0
}

// SetSphen SPHEN
func (r *RegisterMachwf1rType) SetSphen(value bool) {
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
func (r *RegisterMachwf1rType) GetTsoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldTsoenMask) != 0
}

// SetTsoen TSOEN
func (r *RegisterMachwf1rType) SetTsoen(value bool) {
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
func (r *RegisterMachwf1rType) GetDbgmema() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldDbgmemaMask) != 0
}

// SetDbgmema DBGMEMA
func (r *RegisterMachwf1rType) SetDbgmema(value bool) {
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
func (r *RegisterMachwf1rType) GetAvsel() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldAvselMask) != 0
}

// SetAvsel AVSEL
func (r *RegisterMachwf1rType) SetAvsel(value bool) {
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
func (r *RegisterMachwf1rType) GetHashtblsz() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldHashtblszMask) >> RegisterMachwf1rFieldHashtblszShift)
}

// SetHashtblsz HASHTBLSZ
func (r *RegisterMachwf1rType) SetHashtblsz(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldHashtblszMask)|(uint32(value)<<RegisterMachwf1rFieldHashtblszShift))
}

const (
	RegisterMachwf1rFieldL3l4fnumShift = 27
	RegisterMachwf1rFieldL3l4fnumMask  = 0x78000000
)

// GetL3l4fnum L3L4FNUM
func (r *RegisterMachwf1rType) GetL3l4fnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf1rFieldL3l4fnumMask) >> RegisterMachwf1rFieldL3l4fnumShift)
}

// SetL3l4fnum L3L4FNUM
func (r *RegisterMachwf1rType) SetL3l4fnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf1rFieldL3l4fnumMask)|(uint32(value)<<RegisterMachwf1rFieldL3l4fnumShift))
}

// RegisterMachwf2rType HW feature 2 register
type RegisterMachwf2rType uint32

func (r *RegisterMachwf2rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMachwf2rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMachwf2rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMachwf2rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMachwf2rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMachwf2rFieldRxqcntShift = 0
	RegisterMachwf2rFieldRxqcntMask  = 0xf
)

// GetRxqcnt RXQCNT
func (r *RegisterMachwf2rType) GetRxqcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldRxqcntMask) >> RegisterMachwf2rFieldRxqcntShift)
}

// SetRxqcnt RXQCNT
func (r *RegisterMachwf2rType) SetRxqcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldRxqcntMask)|(uint32(value)<<RegisterMachwf2rFieldRxqcntShift))
}

const (
	RegisterMachwf2rFieldTxqcntShift = 6
	RegisterMachwf2rFieldTxqcntMask  = 0x3c0
)

// GetTxqcnt TXQCNT
func (r *RegisterMachwf2rType) GetTxqcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldTxqcntMask) >> RegisterMachwf2rFieldTxqcntShift)
}

// SetTxqcnt TXQCNT
func (r *RegisterMachwf2rType) SetTxqcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldTxqcntMask)|(uint32(value)<<RegisterMachwf2rFieldTxqcntShift))
}

const (
	RegisterMachwf2rFieldRxchcntShift = 12
	RegisterMachwf2rFieldRxchcntMask  = 0xf000
)

// GetRxchcnt RXCHCNT
func (r *RegisterMachwf2rType) GetRxchcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldRxchcntMask) >> RegisterMachwf2rFieldRxchcntShift)
}

// SetRxchcnt RXCHCNT
func (r *RegisterMachwf2rType) SetRxchcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldRxchcntMask)|(uint32(value)<<RegisterMachwf2rFieldRxchcntShift))
}

const (
	RegisterMachwf2rFieldTxchcntShift = 18
	RegisterMachwf2rFieldTxchcntMask  = 0x3c0000
)

// GetTxchcnt TXCHCNT
func (r *RegisterMachwf2rType) GetTxchcnt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldTxchcntMask) >> RegisterMachwf2rFieldTxchcntShift)
}

// SetTxchcnt TXCHCNT
func (r *RegisterMachwf2rType) SetTxchcnt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldTxchcntMask)|(uint32(value)<<RegisterMachwf2rFieldTxchcntShift))
}

const (
	RegisterMachwf2rFieldPpsoutnumShift = 24
	RegisterMachwf2rFieldPpsoutnumMask  = 0x7000000
)

// GetPpsoutnum PPSOUTNUM
func (r *RegisterMachwf2rType) GetPpsoutnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldPpsoutnumMask) >> RegisterMachwf2rFieldPpsoutnumShift)
}

// SetPpsoutnum PPSOUTNUM
func (r *RegisterMachwf2rType) SetPpsoutnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldPpsoutnumMask)|(uint32(value)<<RegisterMachwf2rFieldPpsoutnumShift))
}

const (
	RegisterMachwf2rFieldAuxsnapnumShift = 28
	RegisterMachwf2rFieldAuxsnapnumMask  = 0x70000000
)

// GetAuxsnapnum AUXSNAPNUM
func (r *RegisterMachwf2rType) GetAuxsnapnum() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMachwf2rFieldAuxsnapnumMask) >> RegisterMachwf2rFieldAuxsnapnumShift)
}

// SetAuxsnapnum AUXSNAPNUM
func (r *RegisterMachwf2rType) SetAuxsnapnum(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMachwf2rFieldAuxsnapnumMask)|(uint32(value)<<RegisterMachwf2rFieldAuxsnapnumShift))
}

// RegisterMacmdioarType MDIO address register
type RegisterMacmdioarType uint32

func (r *RegisterMacmdioarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacmdioarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacmdioarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacmdioarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacmdioarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacmdioarFieldMbShift = 0
	RegisterMacmdioarFieldMbMask  = 0x1
)

// GetMb MB
func (r *RegisterMacmdioarType) GetMb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldMbMask) != 0
}

// SetMb MB
func (r *RegisterMacmdioarType) SetMb(value bool) {
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
func (r *RegisterMacmdioarType) GetC45e() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldC45eMask) != 0
}

// SetC45e C45E
func (r *RegisterMacmdioarType) SetC45e(value bool) {
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
func (r *RegisterMacmdioarType) GetGoc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldGocMask) >> RegisterMacmdioarFieldGocShift)
}

// SetGoc GOC
func (r *RegisterMacmdioarType) SetGoc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldGocMask)|(uint32(value)<<RegisterMacmdioarFieldGocShift))
}

const (
	RegisterMacmdioarFieldSkapShift = 4
	RegisterMacmdioarFieldSkapMask  = 0x10
)

// GetSkap SKAP
func (r *RegisterMacmdioarType) GetSkap() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldSkapMask) != 0
}

// SetSkap SKAP
func (r *RegisterMacmdioarType) SetSkap(value bool) {
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
func (r *RegisterMacmdioarType) GetCr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldCrMask) >> RegisterMacmdioarFieldCrShift)
}

// SetCr CR
func (r *RegisterMacmdioarType) SetCr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldCrMask)|(uint32(value)<<RegisterMacmdioarFieldCrShift))
}

const (
	RegisterMacmdioarFieldNtcShift = 12
	RegisterMacmdioarFieldNtcMask  = 0x7000
)

// GetNtc NTC
func (r *RegisterMacmdioarType) GetNtc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldNtcMask) >> RegisterMacmdioarFieldNtcShift)
}

// SetNtc NTC
func (r *RegisterMacmdioarType) SetNtc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldNtcMask)|(uint32(value)<<RegisterMacmdioarFieldNtcShift))
}

const (
	RegisterMacmdioarFieldRdaShift = 16
	RegisterMacmdioarFieldRdaMask  = 0x1f0000
)

// GetRda RDA
func (r *RegisterMacmdioarType) GetRda() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldRdaMask) >> RegisterMacmdioarFieldRdaShift)
}

// SetRda RDA
func (r *RegisterMacmdioarType) SetRda(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldRdaMask)|(uint32(value)<<RegisterMacmdioarFieldRdaShift))
}

const (
	RegisterMacmdioarFieldPaShift = 21
	RegisterMacmdioarFieldPaMask  = 0x3e00000
)

// GetPa PA
func (r *RegisterMacmdioarType) GetPa() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldPaMask) >> RegisterMacmdioarFieldPaShift)
}

// SetPa PA
func (r *RegisterMacmdioarType) SetPa(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldPaMask)|(uint32(value)<<RegisterMacmdioarFieldPaShift))
}

const (
	RegisterMacmdioarFieldBtbShift = 26
	RegisterMacmdioarFieldBtbMask  = 0x4000000
)

// GetBtb BTB
func (r *RegisterMacmdioarType) GetBtb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldBtbMask) != 0
}

// SetBtb BTB
func (r *RegisterMacmdioarType) SetBtb(value bool) {
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
func (r *RegisterMacmdioarType) GetPse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacmdioarFieldPseMask) != 0
}

// SetPse PSE
func (r *RegisterMacmdioarType) SetPse(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacmdioarFieldPseMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacmdioarFieldPseMask)
	}
}

// RegisterMacmdiodrType MDIO data register
type RegisterMacmdiodrType uint32

func (r *RegisterMacmdiodrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacmdiodrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacmdiodrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacmdiodrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacmdiodrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacmdiodrFieldMdShift = 0
	RegisterMacmdiodrFieldMdMask  = 0xffff
)

// GetMd MD
func (r *RegisterMacmdiodrType) GetMd() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdiodrFieldMdMask) >> RegisterMacmdiodrFieldMdShift)
}

// SetMd MD
func (r *RegisterMacmdiodrType) SetMd(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdiodrFieldMdMask)|(uint32(value)<<RegisterMacmdiodrFieldMdShift))
}

const (
	RegisterMacmdiodrFieldRaShift = 16
	RegisterMacmdiodrFieldRaMask  = 0xffff0000
)

// GetRa RA
func (r *RegisterMacmdiodrType) GetRa() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacmdiodrFieldRaMask) >> RegisterMacmdiodrFieldRaShift)
}

// SetRa RA
func (r *RegisterMacmdiodrType) SetRa(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacmdiodrFieldRaMask)|(uint32(value)<<RegisterMacmdiodrFieldRaShift))
}

// RegisterMaca0hrType Address 0 high register
type RegisterMaca0hrType uint32

func (r *RegisterMaca0hrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaca0hrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaca0hrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaca0hrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaca0hrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaca0hrFieldAddrhiShift = 0
	RegisterMaca0hrFieldAddrhiMask  = 0xffff
)

// GetAddrhi ADDRHI
func (r *RegisterMaca0hrType) GetAddrhi() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMaca0hrFieldAddrhiMask) >> RegisterMaca0hrFieldAddrhiShift)
}

// SetAddrhi ADDRHI
func (r *RegisterMaca0hrType) SetAddrhi(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca0hrFieldAddrhiMask)|(uint32(value)<<RegisterMaca0hrFieldAddrhiShift))
}

const (
	RegisterMaca0hrFieldAeShift = 31
	RegisterMaca0hrFieldAeMask  = 0x80000000
)

// GetAe AE
func (r *RegisterMaca0hrType) GetAe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca0hrFieldAeMask) != 0
}

// RegisterMaca0lrType Address 0 low register
type RegisterMaca0lrType uint32

func (r *RegisterMaca0lrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaca0lrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaca0lrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaca0lrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaca0lrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaca0lrFieldAddrloShift = 0
	RegisterMaca0lrFieldAddrloMask  = 0xffffffff
)

// GetAddrlo ADDRLO
func (r *RegisterMaca0lrType) GetAddrlo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMaca0lrFieldAddrloMask) >> RegisterMaca0lrFieldAddrloShift)
}

// SetAddrlo ADDRLO
func (r *RegisterMaca0lrType) SetAddrlo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca0lrFieldAddrloMask)|(uint32(value)<<RegisterMaca0lrFieldAddrloShift))
}

// RegisterMaca1hrType Address 1 high register
type RegisterMaca1hrType uint32

func (r *RegisterMaca1hrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaca1hrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaca1hrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaca1hrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaca1hrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaca1hrFieldAddrhiShift = 0
	RegisterMaca1hrFieldAddrhiMask  = 0xffff
)

// GetAddrhi ADDRHI
func (r *RegisterMaca1hrType) GetAddrhi() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMaca1hrFieldAddrhiMask) >> RegisterMaca1hrFieldAddrhiShift)
}

// SetAddrhi ADDRHI
func (r *RegisterMaca1hrType) SetAddrhi(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca1hrFieldAddrhiMask)|(uint32(value)<<RegisterMaca1hrFieldAddrhiShift))
}

const (
	RegisterMaca1hrFieldMbcShift = 24
	RegisterMaca1hrFieldMbcMask  = 0x3f000000
)

// GetMbc MBC
func (r *RegisterMaca1hrType) GetMbc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaca1hrFieldMbcMask) >> RegisterMaca1hrFieldMbcShift)
}

// SetMbc MBC
func (r *RegisterMaca1hrType) SetMbc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca1hrFieldMbcMask)|(uint32(value)<<RegisterMaca1hrFieldMbcShift))
}

const (
	RegisterMaca1hrFieldSaShift = 30
	RegisterMaca1hrFieldSaMask  = 0x40000000
)

// GetSa SA
func (r *RegisterMaca1hrType) GetSa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca1hrFieldSaMask) != 0
}

// SetSa SA
func (r *RegisterMaca1hrType) SetSa(value bool) {
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
func (r *RegisterMaca1hrType) GetAe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca1hrFieldAeMask) != 0
}

// SetAe AE
func (r *RegisterMaca1hrType) SetAe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaca1hrFieldAeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaca1hrFieldAeMask)
	}
}

// RegisterMaca1lrType Address 1 low register
type RegisterMaca1lrType uint32

func (r *RegisterMaca1lrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaca1lrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaca1lrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaca1lrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaca1lrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaca1lrFieldAddrloShift = 0
	RegisterMaca1lrFieldAddrloMask  = 0xffffffff
)

// GetAddrlo ADDRLO
func (r *RegisterMaca1lrType) GetAddrlo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMaca1lrFieldAddrloMask) >> RegisterMaca1lrFieldAddrloShift)
}

// SetAddrlo ADDRLO
func (r *RegisterMaca1lrType) SetAddrlo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca1lrFieldAddrloMask)|(uint32(value)<<RegisterMaca1lrFieldAddrloShift))
}

// RegisterMaca2hrType Address 2 high register
type RegisterMaca2hrType uint32

func (r *RegisterMaca2hrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaca2hrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaca2hrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaca2hrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaca2hrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaca2hrFieldAddrhiShift = 0
	RegisterMaca2hrFieldAddrhiMask  = 0xffff
)

// GetAddrhi ADDRHI
func (r *RegisterMaca2hrType) GetAddrhi() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMaca2hrFieldAddrhiMask) >> RegisterMaca2hrFieldAddrhiShift)
}

// SetAddrhi ADDRHI
func (r *RegisterMaca2hrType) SetAddrhi(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca2hrFieldAddrhiMask)|(uint32(value)<<RegisterMaca2hrFieldAddrhiShift))
}

const (
	RegisterMaca2hrFieldMbcShift = 24
	RegisterMaca2hrFieldMbcMask  = 0x3f000000
)

// GetMbc MBC
func (r *RegisterMaca2hrType) GetMbc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaca2hrFieldMbcMask) >> RegisterMaca2hrFieldMbcShift)
}

// SetMbc MBC
func (r *RegisterMaca2hrType) SetMbc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca2hrFieldMbcMask)|(uint32(value)<<RegisterMaca2hrFieldMbcShift))
}

const (
	RegisterMaca2hrFieldSaShift = 30
	RegisterMaca2hrFieldSaMask  = 0x40000000
)

// GetSa SA
func (r *RegisterMaca2hrType) GetSa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca2hrFieldSaMask) != 0
}

// SetSa SA
func (r *RegisterMaca2hrType) SetSa(value bool) {
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
func (r *RegisterMaca2hrType) GetAe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca2hrFieldAeMask) != 0
}

// SetAe AE
func (r *RegisterMaca2hrType) SetAe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaca2hrFieldAeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaca2hrFieldAeMask)
	}
}

// RegisterMaca2lrType Address 2 low register
type RegisterMaca2lrType uint32

func (r *RegisterMaca2lrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaca2lrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaca2lrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaca2lrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaca2lrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaca2lrFieldAddrloShift = 0
	RegisterMaca2lrFieldAddrloMask  = 0xffffffff
)

// GetAddrlo ADDRLO
func (r *RegisterMaca2lrType) GetAddrlo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMaca2lrFieldAddrloMask) >> RegisterMaca2lrFieldAddrloShift)
}

// SetAddrlo ADDRLO
func (r *RegisterMaca2lrType) SetAddrlo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca2lrFieldAddrloMask)|(uint32(value)<<RegisterMaca2lrFieldAddrloShift))
}

// RegisterMaca3hrType Address 3 high register
type RegisterMaca3hrType uint32

func (r *RegisterMaca3hrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaca3hrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaca3hrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaca3hrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaca3hrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaca3hrFieldAddrhiShift = 0
	RegisterMaca3hrFieldAddrhiMask  = 0xffff
)

// GetAddrhi ADDRHI
func (r *RegisterMaca3hrType) GetAddrhi() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMaca3hrFieldAddrhiMask) >> RegisterMaca3hrFieldAddrhiShift)
}

// SetAddrhi ADDRHI
func (r *RegisterMaca3hrType) SetAddrhi(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca3hrFieldAddrhiMask)|(uint32(value)<<RegisterMaca3hrFieldAddrhiShift))
}

const (
	RegisterMaca3hrFieldMbcShift = 24
	RegisterMaca3hrFieldMbcMask  = 0x3f000000
)

// GetMbc MBC
func (r *RegisterMaca3hrType) GetMbc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaca3hrFieldMbcMask) >> RegisterMaca3hrFieldMbcShift)
}

// SetMbc MBC
func (r *RegisterMaca3hrType) SetMbc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca3hrFieldMbcMask)|(uint32(value)<<RegisterMaca3hrFieldMbcShift))
}

const (
	RegisterMaca3hrFieldSaShift = 30
	RegisterMaca3hrFieldSaMask  = 0x40000000
)

// GetSa SA
func (r *RegisterMaca3hrType) GetSa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca3hrFieldSaMask) != 0
}

// SetSa SA
func (r *RegisterMaca3hrType) SetSa(value bool) {
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
func (r *RegisterMaca3hrType) GetAe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMaca3hrFieldAeMask) != 0
}

// SetAe AE
func (r *RegisterMaca3hrType) SetAe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMaca3hrFieldAeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMaca3hrFieldAeMask)
	}
}

// RegisterMaca3lrType Address 3 low register
type RegisterMaca3lrType uint32

func (r *RegisterMaca3lrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaca3lrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaca3lrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaca3lrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaca3lrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaca3lrFieldAddrloShift = 0
	RegisterMaca3lrFieldAddrloMask  = 0xffffffff
)

// GetAddrlo ADDRLO
func (r *RegisterMaca3lrType) GetAddrlo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMaca3lrFieldAddrloMask) >> RegisterMaca3lrFieldAddrloShift)
}

// SetAddrlo ADDRLO
func (r *RegisterMaca3lrType) SetAddrlo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaca3lrFieldAddrloMask)|(uint32(value)<<RegisterMaca3lrFieldAddrloShift))
}

// RegisterMmccontrolType MMC control register
type RegisterMmccontrolType uint32

func (r *RegisterMmccontrolType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMmccontrolType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMmccontrolType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMmccontrolType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMmccontrolType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMmccontrolFieldCntrstShift = 0
	RegisterMmccontrolFieldCntrstMask  = 0x1
)

// GetCntrst CNTRST
func (r *RegisterMmccontrolType) GetCntrst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldCntrstMask) != 0
}

// SetCntrst CNTRST
func (r *RegisterMmccontrolType) SetCntrst(value bool) {
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
func (r *RegisterMmccontrolType) GetCntstopro() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldCntstoproMask) != 0
}

// SetCntstopro CNTSTOPRO
func (r *RegisterMmccontrolType) SetCntstopro(value bool) {
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
func (r *RegisterMmccontrolType) GetRstonrd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldRstonrdMask) != 0
}

// SetRstonrd RSTONRD
func (r *RegisterMmccontrolType) SetRstonrd(value bool) {
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
func (r *RegisterMmccontrolType) GetCntfreez() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldCntfreezMask) != 0
}

// SetCntfreez CNTFREEZ
func (r *RegisterMmccontrolType) SetCntfreez(value bool) {
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
func (r *RegisterMmccontrolType) GetCntprst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldCntprstMask) != 0
}

// SetCntprst CNTPRST
func (r *RegisterMmccontrolType) SetCntprst(value bool) {
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
func (r *RegisterMmccontrolType) GetCntprstlvl() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldCntprstlvlMask) != 0
}

// SetCntprstlvl CNTPRSTLVL
func (r *RegisterMmccontrolType) SetCntprstlvl(value bool) {
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
func (r *RegisterMmccontrolType) GetUcdbc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmccontrolFieldUcdbcMask) != 0
}

// SetUcdbc UCDBC
func (r *RegisterMmccontrolType) SetUcdbc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmccontrolFieldUcdbcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmccontrolFieldUcdbcMask)
	}
}

// RegisterMmcrxinterruptType MMC Rx interrupt register
type RegisterMmcrxinterruptType uint32

func (r *RegisterMmcrxinterruptType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMmcrxinterruptType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMmcrxinterruptType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMmcrxinterruptType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMmcrxinterruptType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMmcrxinterruptFieldRxcrcerpisShift = 5
	RegisterMmcrxinterruptFieldRxcrcerpisMask  = 0x20
)

// GetRxcrcerpis RXCRCERPIS
func (r *RegisterMmcrxinterruptType) GetRxcrcerpis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptFieldRxcrcerpisMask) != 0
}

// SetRxcrcerpis RXCRCERPIS
func (r *RegisterMmcrxinterruptType) SetRxcrcerpis(value bool) {
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
func (r *RegisterMmcrxinterruptType) GetRxalgnerpis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptFieldRxalgnerpisMask) != 0
}

// SetRxalgnerpis RXALGNERPIS
func (r *RegisterMmcrxinterruptType) SetRxalgnerpis(value bool) {
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
func (r *RegisterMmcrxinterruptType) GetRxucgpis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptFieldRxucgpisMask) != 0
}

// SetRxucgpis RXUCGPIS
func (r *RegisterMmcrxinterruptType) SetRxucgpis(value bool) {
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
func (r *RegisterMmcrxinterruptType) GetRxlpiuscis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptFieldRxlpiuscisMask) != 0
}

// SetRxlpiuscis RXLPIUSCIS
func (r *RegisterMmcrxinterruptType) SetRxlpiuscis(value bool) {
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
func (r *RegisterMmcrxinterruptType) GetRxlpitrcis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptFieldRxlpitrcisMask) != 0
}

// SetRxlpitrcis RXLPITRCIS
func (r *RegisterMmcrxinterruptType) SetRxlpitrcis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmcrxinterruptFieldRxlpitrcisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmcrxinterruptFieldRxlpitrcisMask)
	}
}

// RegisterMmctxinterruptType MMC Tx interrupt register
type RegisterMmctxinterruptType uint32

func (r *RegisterMmctxinterruptType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMmctxinterruptType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMmctxinterruptType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMmctxinterruptType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMmctxinterruptType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMmctxinterruptFieldTxscolgpisShift = 14
	RegisterMmctxinterruptFieldTxscolgpisMask  = 0x4000
)

// GetTxscolgpis TXSCOLGPIS
func (r *RegisterMmctxinterruptType) GetTxscolgpis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptFieldTxscolgpisMask) != 0
}

// SetTxscolgpis TXSCOLGPIS
func (r *RegisterMmctxinterruptType) SetTxscolgpis(value bool) {
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
func (r *RegisterMmctxinterruptType) GetTxmcolgpis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptFieldTxmcolgpisMask) != 0
}

// SetTxmcolgpis TXMCOLGPIS
func (r *RegisterMmctxinterruptType) SetTxmcolgpis(value bool) {
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
func (r *RegisterMmctxinterruptType) GetTxgpktis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptFieldTxgpktisMask) != 0
}

// SetTxgpktis TXGPKTIS
func (r *RegisterMmctxinterruptType) SetTxgpktis(value bool) {
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
func (r *RegisterMmctxinterruptType) GetTxlpiuscis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptFieldTxlpiuscisMask) != 0
}

// SetTxlpiuscis TXLPIUSCIS
func (r *RegisterMmctxinterruptType) SetTxlpiuscis(value bool) {
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
func (r *RegisterMmctxinterruptType) GetTxlpitrcis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptFieldTxlpitrcisMask) != 0
}

// SetTxlpitrcis TXLPITRCIS
func (r *RegisterMmctxinterruptType) SetTxlpitrcis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMmctxinterruptFieldTxlpitrcisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMmctxinterruptFieldTxlpitrcisMask)
	}
}

// RegisterMmcrxinterruptmaskType MMC Rx interrupt mask register
type RegisterMmcrxinterruptmaskType uint32

func (r *RegisterMmcrxinterruptmaskType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMmcrxinterruptmaskType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMmcrxinterruptmaskType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMmcrxinterruptmaskType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMmcrxinterruptmaskType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMmcrxinterruptmaskFieldRxcrcerpimShift = 5
	RegisterMmcrxinterruptmaskFieldRxcrcerpimMask  = 0x20
)

// GetRxcrcerpim RXCRCERPIM
func (r *RegisterMmcrxinterruptmaskType) GetRxcrcerpim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptmaskFieldRxcrcerpimMask) != 0
}

// SetRxcrcerpim RXCRCERPIM
func (r *RegisterMmcrxinterruptmaskType) SetRxcrcerpim(value bool) {
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
func (r *RegisterMmcrxinterruptmaskType) GetRxalgnerpim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptmaskFieldRxalgnerpimMask) != 0
}

// SetRxalgnerpim RXALGNERPIM
func (r *RegisterMmcrxinterruptmaskType) SetRxalgnerpim(value bool) {
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
func (r *RegisterMmcrxinterruptmaskType) GetRxucgpim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptmaskFieldRxucgpimMask) != 0
}

// SetRxucgpim RXUCGPIM
func (r *RegisterMmcrxinterruptmaskType) SetRxucgpim(value bool) {
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
func (r *RegisterMmcrxinterruptmaskType) GetRxlpiuscim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptmaskFieldRxlpiuscimMask) != 0
}

// SetRxlpiuscim RXLPIUSCIM
func (r *RegisterMmcrxinterruptmaskType) SetRxlpiuscim(value bool) {
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
func (r *RegisterMmcrxinterruptmaskType) GetRxlpitrcim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmcrxinterruptmaskFieldRxlpitrcimMask) != 0
}

// RegisterMmctxinterruptmaskType MMC Tx interrupt mask register
type RegisterMmctxinterruptmaskType uint32

func (r *RegisterMmctxinterruptmaskType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMmctxinterruptmaskType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMmctxinterruptmaskType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMmctxinterruptmaskType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMmctxinterruptmaskType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMmctxinterruptmaskFieldTxscolgpimShift = 14
	RegisterMmctxinterruptmaskFieldTxscolgpimMask  = 0x4000
)

// GetTxscolgpim TXSCOLGPIM
func (r *RegisterMmctxinterruptmaskType) GetTxscolgpim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptmaskFieldTxscolgpimMask) != 0
}

// SetTxscolgpim TXSCOLGPIM
func (r *RegisterMmctxinterruptmaskType) SetTxscolgpim(value bool) {
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
func (r *RegisterMmctxinterruptmaskType) GetTxmcolgpim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptmaskFieldTxmcolgpimMask) != 0
}

// SetTxmcolgpim TXMCOLGPIM
func (r *RegisterMmctxinterruptmaskType) SetTxmcolgpim(value bool) {
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
func (r *RegisterMmctxinterruptmaskType) GetTxgpktim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptmaskFieldTxgpktimMask) != 0
}

// SetTxgpktim TXGPKTIM
func (r *RegisterMmctxinterruptmaskType) SetTxgpktim(value bool) {
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
func (r *RegisterMmctxinterruptmaskType) GetTxlpiuscim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptmaskFieldTxlpiuscimMask) != 0
}

// SetTxlpiuscim TXLPIUSCIM
func (r *RegisterMmctxinterruptmaskType) SetTxlpiuscim(value bool) {
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
func (r *RegisterMmctxinterruptmaskType) GetTxlpitrcim() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMmctxinterruptmaskFieldTxlpitrcimMask) != 0
}

// RegisterTxsinglecollisiongoodpacketsType Tx single collision good packets register
type RegisterTxsinglecollisiongoodpacketsType uint32

func (r *RegisterTxsinglecollisiongoodpacketsType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTxsinglecollisiongoodpacketsType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTxsinglecollisiongoodpacketsType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTxsinglecollisiongoodpacketsType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTxsinglecollisiongoodpacketsType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgShift = 0
	RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgMask  = 0xffffffff
)

// GetTxsnglcolg TXSNGLCOLG
func (r *RegisterTxsinglecollisiongoodpacketsType) GetTxsnglcolg() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgMask) >> RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgShift)
}

// SetTxsnglcolg TXSNGLCOLG
func (r *RegisterTxsinglecollisiongoodpacketsType) SetTxsnglcolg(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgMask)|(uint32(value)<<RegisterTxsinglecollisiongoodpacketsFieldTxsnglcolgShift))
}

// RegisterTxmultiplecollisiongoodpacketsType Tx multiple collision good packets register
type RegisterTxmultiplecollisiongoodpacketsType uint32

func (r *RegisterTxmultiplecollisiongoodpacketsType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTxmultiplecollisiongoodpacketsType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTxmultiplecollisiongoodpacketsType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTxmultiplecollisiongoodpacketsType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTxmultiplecollisiongoodpacketsType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgShift = 0
	RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgMask  = 0xffffffff
)

// GetTxmultcolg TXMULTCOLG
func (r *RegisterTxmultiplecollisiongoodpacketsType) GetTxmultcolg() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgMask) >> RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgShift)
}

// SetTxmultcolg TXMULTCOLG
func (r *RegisterTxmultiplecollisiongoodpacketsType) SetTxmultcolg(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgMask)|(uint32(value)<<RegisterTxmultiplecollisiongoodpacketsFieldTxmultcolgShift))
}

// RegisterTxpacketcountgoodType Tx packet count good register
type RegisterTxpacketcountgoodType uint32

func (r *RegisterTxpacketcountgoodType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTxpacketcountgoodType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTxpacketcountgoodType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTxpacketcountgoodType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTxpacketcountgoodType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTxpacketcountgoodFieldTxpktgShift = 0
	RegisterTxpacketcountgoodFieldTxpktgMask  = 0xffffffff
)

// GetTxpktg TXPKTG
func (r *RegisterTxpacketcountgoodType) GetTxpktg() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxpacketcountgoodFieldTxpktgMask) >> RegisterTxpacketcountgoodFieldTxpktgShift)
}

// SetTxpktg TXPKTG
func (r *RegisterTxpacketcountgoodType) SetTxpktg(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxpacketcountgoodFieldTxpktgMask)|(uint32(value)<<RegisterTxpacketcountgoodFieldTxpktgShift))
}

// RegisterRxcrcerrorpacketsType Rx CRC error packets register
type RegisterRxcrcerrorpacketsType uint32

func (r *RegisterRxcrcerrorpacketsType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRxcrcerrorpacketsType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRxcrcerrorpacketsType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRxcrcerrorpacketsType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRxcrcerrorpacketsType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRxcrcerrorpacketsFieldRxcrcerrShift = 0
	RegisterRxcrcerrorpacketsFieldRxcrcerrMask  = 0xffffffff
)

// GetRxcrcerr RXCRCERR
func (r *RegisterRxcrcerrorpacketsType) GetRxcrcerr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxcrcerrorpacketsFieldRxcrcerrMask) >> RegisterRxcrcerrorpacketsFieldRxcrcerrShift)
}

// SetRxcrcerr RXCRCERR
func (r *RegisterRxcrcerrorpacketsType) SetRxcrcerr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxcrcerrorpacketsFieldRxcrcerrMask)|(uint32(value)<<RegisterRxcrcerrorpacketsFieldRxcrcerrShift))
}

// RegisterRxalignmenterrorpacketsType Rx alignment error packets register
type RegisterRxalignmenterrorpacketsType uint32

func (r *RegisterRxalignmenterrorpacketsType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRxalignmenterrorpacketsType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRxalignmenterrorpacketsType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRxalignmenterrorpacketsType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRxalignmenterrorpacketsType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRxalignmenterrorpacketsFieldRxalgnerrShift = 0
	RegisterRxalignmenterrorpacketsFieldRxalgnerrMask  = 0xffffffff
)

// GetRxalgnerr RXALGNERR
func (r *RegisterRxalignmenterrorpacketsType) GetRxalgnerr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxalignmenterrorpacketsFieldRxalgnerrMask) >> RegisterRxalignmenterrorpacketsFieldRxalgnerrShift)
}

// SetRxalgnerr RXALGNERR
func (r *RegisterRxalignmenterrorpacketsType) SetRxalgnerr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxalignmenterrorpacketsFieldRxalgnerrMask)|(uint32(value)<<RegisterRxalignmenterrorpacketsFieldRxalgnerrShift))
}

// RegisterRxunicastpacketsgoodType Rx unicast packets good register
type RegisterRxunicastpacketsgoodType uint32

func (r *RegisterRxunicastpacketsgoodType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRxunicastpacketsgoodType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRxunicastpacketsgoodType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRxunicastpacketsgoodType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRxunicastpacketsgoodType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRxunicastpacketsgoodFieldRxucastgShift = 0
	RegisterRxunicastpacketsgoodFieldRxucastgMask  = 0xffffffff
)

// GetRxucastg RXUCASTG
func (r *RegisterRxunicastpacketsgoodType) GetRxucastg() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxunicastpacketsgoodFieldRxucastgMask) >> RegisterRxunicastpacketsgoodFieldRxucastgShift)
}

// SetRxucastg RXUCASTG
func (r *RegisterRxunicastpacketsgoodType) SetRxucastg(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxunicastpacketsgoodFieldRxucastgMask)|(uint32(value)<<RegisterRxunicastpacketsgoodFieldRxucastgShift))
}

// RegisterTxlpiuseccntrType Tx LPI microsecond timer register
type RegisterTxlpiuseccntrType uint32

func (r *RegisterTxlpiuseccntrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTxlpiuseccntrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTxlpiuseccntrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTxlpiuseccntrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTxlpiuseccntrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTxlpiuseccntrFieldTxlpiuscShift = 0
	RegisterTxlpiuseccntrFieldTxlpiuscMask  = 0xffffffff
)

// GetTxlpiusc TXLPIUSC
func (r *RegisterTxlpiuseccntrType) GetTxlpiusc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxlpiuseccntrFieldTxlpiuscMask) >> RegisterTxlpiuseccntrFieldTxlpiuscShift)
}

// SetTxlpiusc TXLPIUSC
func (r *RegisterTxlpiuseccntrType) SetTxlpiusc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxlpiuseccntrFieldTxlpiuscMask)|(uint32(value)<<RegisterTxlpiuseccntrFieldTxlpiuscShift))
}

// RegisterTxlpitrancntrType Tx LPI transition counter register
type RegisterTxlpitrancntrType uint32

func (r *RegisterTxlpitrancntrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterTxlpitrancntrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterTxlpitrancntrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterTxlpitrancntrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterTxlpitrancntrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterTxlpitrancntrFieldTxlpitrcShift = 0
	RegisterTxlpitrancntrFieldTxlpitrcMask  = 0xffffffff
)

// GetTxlpitrc TXLPITRC
func (r *RegisterTxlpitrancntrType) GetTxlpitrc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterTxlpitrancntrFieldTxlpitrcMask) >> RegisterTxlpitrancntrFieldTxlpitrcShift)
}

// SetTxlpitrc TXLPITRC
func (r *RegisterTxlpitrancntrType) SetTxlpitrc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTxlpitrancntrFieldTxlpitrcMask)|(uint32(value)<<RegisterTxlpitrancntrFieldTxlpitrcShift))
}

// RegisterRxlpiuseccntrType Rx LPI microsecond counter register
type RegisterRxlpiuseccntrType uint32

func (r *RegisterRxlpiuseccntrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRxlpiuseccntrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRxlpiuseccntrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRxlpiuseccntrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRxlpiuseccntrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRxlpiuseccntrFieldRxlpiuscShift = 0
	RegisterRxlpiuseccntrFieldRxlpiuscMask  = 0xffffffff
)

// GetRxlpiusc RXLPIUSC
func (r *RegisterRxlpiuseccntrType) GetRxlpiusc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxlpiuseccntrFieldRxlpiuscMask) >> RegisterRxlpiuseccntrFieldRxlpiuscShift)
}

// SetRxlpiusc RXLPIUSC
func (r *RegisterRxlpiuseccntrType) SetRxlpiusc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxlpiuseccntrFieldRxlpiuscMask)|(uint32(value)<<RegisterRxlpiuseccntrFieldRxlpiuscShift))
}

// RegisterRxlpitrancntrType Rx LPI transition counter register
type RegisterRxlpitrancntrType uint32

func (r *RegisterRxlpitrancntrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterRxlpitrancntrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterRxlpitrancntrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterRxlpitrancntrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterRxlpitrancntrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterRxlpitrancntrFieldRxlpitrcShift = 0
	RegisterRxlpitrancntrFieldRxlpitrcMask  = 0xffffffff
)

// GetRxlpitrc RXLPITRC
func (r *RegisterRxlpitrancntrType) GetRxlpitrc() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterRxlpitrancntrFieldRxlpitrcMask) >> RegisterRxlpitrancntrFieldRxlpitrcShift)
}

// SetRxlpitrc RXLPITRC
func (r *RegisterRxlpitrancntrType) SetRxlpitrc(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterRxlpitrancntrFieldRxlpitrcMask)|(uint32(value)<<RegisterRxlpitrancntrFieldRxlpitrcShift))
}

// RegisterMacl3l4c0rType L3 and L4 control 0 register
type RegisterMacl3l4c0rType uint32

func (r *RegisterMacl3l4c0rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl3l4c0rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl3l4c0rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl3l4c0rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl3l4c0rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl3l4c0rFieldL3pen0Shift = 0
	RegisterMacl3l4c0rFieldL3pen0Mask  = 0x1
)

// GetL3pen0 L3PEN0
func (r *RegisterMacl3l4c0rType) GetL3pen0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3pen0Mask) != 0
}

// SetL3pen0 L3PEN0
func (r *RegisterMacl3l4c0rType) SetL3pen0(value bool) {
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
func (r *RegisterMacl3l4c0rType) GetL3sam0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3sam0Mask) != 0
}

// SetL3sam0 L3SAM0
func (r *RegisterMacl3l4c0rType) SetL3sam0(value bool) {
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
func (r *RegisterMacl3l4c0rType) GetL3saim0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3saim0Mask) != 0
}

// SetL3saim0 L3SAIM0
func (r *RegisterMacl3l4c0rType) SetL3saim0(value bool) {
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
func (r *RegisterMacl3l4c0rType) GetL3dam0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3dam0Mask) != 0
}

// SetL3dam0 L3DAM0
func (r *RegisterMacl3l4c0rType) SetL3dam0(value bool) {
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
func (r *RegisterMacl3l4c0rType) GetL3daim0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3daim0Mask) != 0
}

// SetL3daim0 L3DAIM0
func (r *RegisterMacl3l4c0rType) SetL3daim0(value bool) {
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
func (r *RegisterMacl3l4c0rType) GetL3hsbm0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3hsbm0Mask) >> RegisterMacl3l4c0rFieldL3hsbm0Shift)
}

// SetL3hsbm0 L3HSBM0
func (r *RegisterMacl3l4c0rType) SetL3hsbm0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL3hsbm0Mask)|(uint32(value)<<RegisterMacl3l4c0rFieldL3hsbm0Shift))
}

const (
	RegisterMacl3l4c0rFieldL3hdbm0Shift = 11
	RegisterMacl3l4c0rFieldL3hdbm0Mask  = 0xf800
)

// GetL3hdbm0 L3HDBM0
func (r *RegisterMacl3l4c0rType) GetL3hdbm0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL3hdbm0Mask) >> RegisterMacl3l4c0rFieldL3hdbm0Shift)
}

// SetL3hdbm0 L3HDBM0
func (r *RegisterMacl3l4c0rType) SetL3hdbm0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL3hdbm0Mask)|(uint32(value)<<RegisterMacl3l4c0rFieldL3hdbm0Shift))
}

const (
	RegisterMacl3l4c0rFieldL4pen0Shift = 16
	RegisterMacl3l4c0rFieldL4pen0Mask  = 0x10000
)

// GetL4pen0 L4PEN0
func (r *RegisterMacl3l4c0rType) GetL4pen0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL4pen0Mask) != 0
}

// SetL4pen0 L4PEN0
func (r *RegisterMacl3l4c0rType) SetL4pen0(value bool) {
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
func (r *RegisterMacl3l4c0rType) GetL4spm0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL4spm0Mask) != 0
}

// SetL4spm0 L4SPM0
func (r *RegisterMacl3l4c0rType) SetL4spm0(value bool) {
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
func (r *RegisterMacl3l4c0rType) GetL4spim0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL4spim0Mask) != 0
}

// SetL4spim0 L4SPIM0
func (r *RegisterMacl3l4c0rType) SetL4spim0(value bool) {
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
func (r *RegisterMacl3l4c0rType) GetL4dpm0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL4dpm0Mask) != 0
}

// SetL4dpm0 L4DPM0
func (r *RegisterMacl3l4c0rType) SetL4dpm0(value bool) {
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
func (r *RegisterMacl3l4c0rType) GetL4dpim0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c0rFieldL4dpim0Mask) != 0
}

// SetL4dpim0 L4DPIM0
func (r *RegisterMacl3l4c0rType) SetL4dpim0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c0rFieldL4dpim0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c0rFieldL4dpim0Mask)
	}
}

// RegisterMacl4a0rType Layer4 address filter 0 register
type RegisterMacl4a0rType uint32

func (r *RegisterMacl4a0rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl4a0rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl4a0rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl4a0rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl4a0rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl4a0rFieldL4sp0Shift = 0
	RegisterMacl4a0rFieldL4sp0Mask  = 0xffff
)

// GetL4sp0 L4SP0
func (r *RegisterMacl4a0rType) GetL4sp0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacl4a0rFieldL4sp0Mask) >> RegisterMacl4a0rFieldL4sp0Shift)
}

// SetL4sp0 L4SP0
func (r *RegisterMacl4a0rType) SetL4sp0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl4a0rFieldL4sp0Mask)|(uint32(value)<<RegisterMacl4a0rFieldL4sp0Shift))
}

const (
	RegisterMacl4a0rFieldL4dp0Shift = 16
	RegisterMacl4a0rFieldL4dp0Mask  = 0xffff0000
)

// GetL4dp0 L4DP0
func (r *RegisterMacl4a0rType) GetL4dp0() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacl4a0rFieldL4dp0Mask) >> RegisterMacl4a0rFieldL4dp0Shift)
}

// SetL4dp0 L4DP0
func (r *RegisterMacl4a0rType) SetL4dp0(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl4a0rFieldL4dp0Mask)|(uint32(value)<<RegisterMacl4a0rFieldL4dp0Shift))
}

// RegisterMacl3a00rType MACL3A00R
type RegisterMacl3a00rType uint32

func (r *RegisterMacl3a00rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl3a00rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl3a00rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl3a00rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl3a00rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl3a00rFieldL3a00Shift = 0
	RegisterMacl3a00rFieldL3a00Mask  = 0xffffffff
)

// GetL3a00 L3A00
func (r *RegisterMacl3a00rType) GetL3a00() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a00rFieldL3a00Mask) >> RegisterMacl3a00rFieldL3a00Shift)
}

// SetL3a00 L3A00
func (r *RegisterMacl3a00rType) SetL3a00(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a00rFieldL3a00Mask)|(uint32(value)<<RegisterMacl3a00rFieldL3a00Shift))
}

// RegisterMacl3a10rType Layer3 address 1 filter 0 register
type RegisterMacl3a10rType uint32

func (r *RegisterMacl3a10rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl3a10rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl3a10rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl3a10rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl3a10rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl3a10rFieldL3a10Shift = 0
	RegisterMacl3a10rFieldL3a10Mask  = 0xffffffff
)

// GetL3a10 L3A10
func (r *RegisterMacl3a10rType) GetL3a10() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a10rFieldL3a10Mask) >> RegisterMacl3a10rFieldL3a10Shift)
}

// SetL3a10 L3A10
func (r *RegisterMacl3a10rType) SetL3a10(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a10rFieldL3a10Mask)|(uint32(value)<<RegisterMacl3a10rFieldL3a10Shift))
}

// RegisterMacl3a20Type Layer3 Address 2 filter 0 register
type RegisterMacl3a20Type uint32

func (r *RegisterMacl3a20Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl3a20Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl3a20Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl3a20Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl3a20Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl3a20FieldL3a20Shift = 0
	RegisterMacl3a20FieldL3a20Mask  = 0xffffffff
)

// GetL3a20 L3A20
func (r *RegisterMacl3a20Type) GetL3a20() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a20FieldL3a20Mask) >> RegisterMacl3a20FieldL3a20Shift)
}

// SetL3a20 L3A20
func (r *RegisterMacl3a20Type) SetL3a20(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a20FieldL3a20Mask)|(uint32(value)<<RegisterMacl3a20FieldL3a20Shift))
}

// RegisterMacl3a30Type Layer3 Address 3 filter 0 register
type RegisterMacl3a30Type uint32

func (r *RegisterMacl3a30Type) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl3a30Type) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl3a30Type) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl3a30Type) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl3a30Type) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl3a30FieldL3a30Shift = 0
	RegisterMacl3a30FieldL3a30Mask  = 0xffffffff
)

// GetL3a30 L3A30
func (r *RegisterMacl3a30Type) GetL3a30() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a30FieldL3a30Mask) >> RegisterMacl3a30FieldL3a30Shift)
}

// SetL3a30 L3A30
func (r *RegisterMacl3a30Type) SetL3a30(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a30FieldL3a30Mask)|(uint32(value)<<RegisterMacl3a30FieldL3a30Shift))
}

// RegisterMacl3l4c1rType L3 and L4 control 1 register
type RegisterMacl3l4c1rType uint32

func (r *RegisterMacl3l4c1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl3l4c1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl3l4c1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl3l4c1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl3l4c1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl3l4c1rFieldL3pen1Shift = 0
	RegisterMacl3l4c1rFieldL3pen1Mask  = 0x1
)

// GetL3pen1 L3PEN1
func (r *RegisterMacl3l4c1rType) GetL3pen1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3pen1Mask) != 0
}

// SetL3pen1 L3PEN1
func (r *RegisterMacl3l4c1rType) SetL3pen1(value bool) {
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
func (r *RegisterMacl3l4c1rType) GetL3sam1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3sam1Mask) != 0
}

// SetL3sam1 L3SAM1
func (r *RegisterMacl3l4c1rType) SetL3sam1(value bool) {
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
func (r *RegisterMacl3l4c1rType) GetL3saim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3saim1Mask) != 0
}

// SetL3saim1 L3SAIM1
func (r *RegisterMacl3l4c1rType) SetL3saim1(value bool) {
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
func (r *RegisterMacl3l4c1rType) GetL3dam1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3dam1Mask) != 0
}

// SetL3dam1 L3DAM1
func (r *RegisterMacl3l4c1rType) SetL3dam1(value bool) {
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
func (r *RegisterMacl3l4c1rType) GetL3daim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3daim1Mask) != 0
}

// SetL3daim1 L3DAIM1
func (r *RegisterMacl3l4c1rType) SetL3daim1(value bool) {
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
func (r *RegisterMacl3l4c1rType) GetL3hsbm1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3hsbm1Mask) >> RegisterMacl3l4c1rFieldL3hsbm1Shift)
}

// SetL3hsbm1 L3HSBM1
func (r *RegisterMacl3l4c1rType) SetL3hsbm1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL3hsbm1Mask)|(uint32(value)<<RegisterMacl3l4c1rFieldL3hsbm1Shift))
}

const (
	RegisterMacl3l4c1rFieldL3hdbm1Shift = 11
	RegisterMacl3l4c1rFieldL3hdbm1Mask  = 0xf800
)

// GetL3hdbm1 L3HDBM1
func (r *RegisterMacl3l4c1rType) GetL3hdbm1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL3hdbm1Mask) >> RegisterMacl3l4c1rFieldL3hdbm1Shift)
}

// SetL3hdbm1 L3HDBM1
func (r *RegisterMacl3l4c1rType) SetL3hdbm1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL3hdbm1Mask)|(uint32(value)<<RegisterMacl3l4c1rFieldL3hdbm1Shift))
}

const (
	RegisterMacl3l4c1rFieldL4pen1Shift = 16
	RegisterMacl3l4c1rFieldL4pen1Mask  = 0x10000
)

// GetL4pen1 L4PEN1
func (r *RegisterMacl3l4c1rType) GetL4pen1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL4pen1Mask) != 0
}

// SetL4pen1 L4PEN1
func (r *RegisterMacl3l4c1rType) SetL4pen1(value bool) {
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
func (r *RegisterMacl3l4c1rType) GetL4spm1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL4spm1Mask) != 0
}

// SetL4spm1 L4SPM1
func (r *RegisterMacl3l4c1rType) SetL4spm1(value bool) {
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
func (r *RegisterMacl3l4c1rType) GetL4spim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL4spim1Mask) != 0
}

// SetL4spim1 L4SPIM1
func (r *RegisterMacl3l4c1rType) SetL4spim1(value bool) {
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
func (r *RegisterMacl3l4c1rType) GetL4dpm1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL4dpm1Mask) != 0
}

// SetL4dpm1 L4DPM1
func (r *RegisterMacl3l4c1rType) SetL4dpm1(value bool) {
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
func (r *RegisterMacl3l4c1rType) GetL4dpim1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacl3l4c1rFieldL4dpim1Mask) != 0
}

// SetL4dpim1 L4DPIM1
func (r *RegisterMacl3l4c1rType) SetL4dpim1(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacl3l4c1rFieldL4dpim1Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacl3l4c1rFieldL4dpim1Mask)
	}
}

// RegisterMacl4a1rType Layer 4 address filter 1 register
type RegisterMacl4a1rType uint32

func (r *RegisterMacl4a1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl4a1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl4a1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl4a1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl4a1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl4a1rFieldL4sp1Shift = 0
	RegisterMacl4a1rFieldL4sp1Mask  = 0xffff
)

// GetL4sp1 L4SP1
func (r *RegisterMacl4a1rType) GetL4sp1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacl4a1rFieldL4sp1Mask) >> RegisterMacl4a1rFieldL4sp1Shift)
}

// SetL4sp1 L4SP1
func (r *RegisterMacl4a1rType) SetL4sp1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl4a1rFieldL4sp1Mask)|(uint32(value)<<RegisterMacl4a1rFieldL4sp1Shift))
}

const (
	RegisterMacl4a1rFieldL4dp1Shift = 16
	RegisterMacl4a1rFieldL4dp1Mask  = 0xffff0000
)

// GetL4dp1 L4DP1
func (r *RegisterMacl4a1rType) GetL4dp1() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacl4a1rFieldL4dp1Mask) >> RegisterMacl4a1rFieldL4dp1Shift)
}

// SetL4dp1 L4DP1
func (r *RegisterMacl4a1rType) SetL4dp1(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl4a1rFieldL4dp1Mask)|(uint32(value)<<RegisterMacl4a1rFieldL4dp1Shift))
}

// RegisterMacl3a01rType Layer3 address 0 filter 1 Register
type RegisterMacl3a01rType uint32

func (r *RegisterMacl3a01rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl3a01rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl3a01rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl3a01rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl3a01rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl3a01rFieldL3a01Shift = 0
	RegisterMacl3a01rFieldL3a01Mask  = 0xffffffff
)

// GetL3a01 L3A01
func (r *RegisterMacl3a01rType) GetL3a01() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a01rFieldL3a01Mask) >> RegisterMacl3a01rFieldL3a01Shift)
}

// SetL3a01 L3A01
func (r *RegisterMacl3a01rType) SetL3a01(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a01rFieldL3a01Mask)|(uint32(value)<<RegisterMacl3a01rFieldL3a01Shift))
}

// RegisterMacl3a11rType Layer3 address 1 filter 1 register
type RegisterMacl3a11rType uint32

func (r *RegisterMacl3a11rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl3a11rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl3a11rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl3a11rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl3a11rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl3a11rFieldL3a11Shift = 0
	RegisterMacl3a11rFieldL3a11Mask  = 0xffffffff
)

// GetL3a11 L3A11
func (r *RegisterMacl3a11rType) GetL3a11() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a11rFieldL3a11Mask) >> RegisterMacl3a11rFieldL3a11Shift)
}

// SetL3a11 L3A11
func (r *RegisterMacl3a11rType) SetL3a11(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a11rFieldL3a11Mask)|(uint32(value)<<RegisterMacl3a11rFieldL3a11Shift))
}

// RegisterMacl3a21rType Layer3 address 2 filter 1 Register
type RegisterMacl3a21rType uint32

func (r *RegisterMacl3a21rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl3a21rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl3a21rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl3a21rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl3a21rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl3a21rFieldL3a21Shift = 0
	RegisterMacl3a21rFieldL3a21Mask  = 0xffffffff
)

// GetL3a21 L3A21
func (r *RegisterMacl3a21rType) GetL3a21() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a21rFieldL3a21Mask) >> RegisterMacl3a21rFieldL3a21Shift)
}

// SetL3a21 L3A21
func (r *RegisterMacl3a21rType) SetL3a21(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a21rFieldL3a21Mask)|(uint32(value)<<RegisterMacl3a21rFieldL3a21Shift))
}

// RegisterMacl3a31rType Layer3 address 3 filter 1 register
type RegisterMacl3a31rType uint32

func (r *RegisterMacl3a31rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacl3a31rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacl3a31rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacl3a31rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacl3a31rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacl3a31rFieldL3a31Shift = 0
	RegisterMacl3a31rFieldL3a31Mask  = 0xffffffff
)

// GetL3a31 L3A31
func (r *RegisterMacl3a31rType) GetL3a31() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacl3a31rFieldL3a31Mask) >> RegisterMacl3a31rFieldL3a31Shift)
}

// SetL3a31 L3A31
func (r *RegisterMacl3a31rType) SetL3a31(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacl3a31rFieldL3a31Mask)|(uint32(value)<<RegisterMacl3a31rFieldL3a31Shift))
}

// RegisterMacarparType ARP address register
type RegisterMacarparType uint32

func (r *RegisterMacarparType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacarparType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacarparType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacarparType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacarparType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacarparFieldArppaShift = 0
	RegisterMacarparFieldArppaMask  = 0xffffffff
)

// GetArppa ARPPA
func (r *RegisterMacarparType) GetArppa() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacarparFieldArppaMask) >> RegisterMacarparFieldArppaShift)
}

// SetArppa ARPPA
func (r *RegisterMacarparType) SetArppa(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacarparFieldArppaMask)|(uint32(value)<<RegisterMacarparFieldArppaShift))
}

// RegisterMactscrType Timestamp control Register
type RegisterMactscrType uint32

func (r *RegisterMactscrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMactscrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMactscrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMactscrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMactscrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMactscrFieldTsenaShift = 0
	RegisterMactscrFieldTsenaMask  = 0x1
)

// GetTsena TSENA
func (r *RegisterMactscrType) GetTsena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsenaMask) != 0
}

// SetTsena TSENA
func (r *RegisterMactscrType) SetTsena(value bool) {
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
func (r *RegisterMactscrType) GetTscfupdt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTscfupdtMask) != 0
}

// SetTscfupdt TSCFUPDT
func (r *RegisterMactscrType) SetTscfupdt(value bool) {
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
func (r *RegisterMactscrType) GetTsinit() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsinitMask) != 0
}

// SetTsinit TSINIT
func (r *RegisterMactscrType) SetTsinit(value bool) {
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
func (r *RegisterMactscrType) GetTsupdt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsupdtMask) != 0
}

// SetTsupdt TSUPDT
func (r *RegisterMactscrType) SetTsupdt(value bool) {
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
func (r *RegisterMactscrType) GetTsaddreg() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsaddregMask) != 0
}

// SetTsaddreg TSADDREG
func (r *RegisterMactscrType) SetTsaddreg(value bool) {
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
func (r *RegisterMactscrType) GetTsenall() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsenallMask) != 0
}

// SetTsenall TSENALL
func (r *RegisterMactscrType) SetTsenall(value bool) {
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
func (r *RegisterMactscrType) GetTsctrlssr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsctrlssrMask) != 0
}

// SetTsctrlssr TSCTRLSSR
func (r *RegisterMactscrType) SetTsctrlssr(value bool) {
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
func (r *RegisterMactscrType) GetTsver2ena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsver2enaMask) != 0
}

// SetTsver2ena TSVER2ENA
func (r *RegisterMactscrType) SetTsver2ena(value bool) {
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
func (r *RegisterMactscrType) GetTsipena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsipenaMask) != 0
}

// SetTsipena TSIPENA
func (r *RegisterMactscrType) SetTsipena(value bool) {
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
func (r *RegisterMactscrType) GetTsipv6ena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsipv6enaMask) != 0
}

// SetTsipv6ena TSIPV6ENA
func (r *RegisterMactscrType) SetTsipv6ena(value bool) {
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
func (r *RegisterMactscrType) GetTsipv4ena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsipv4enaMask) != 0
}

// SetTsipv4ena TSIPV4ENA
func (r *RegisterMactscrType) SetTsipv4ena(value bool) {
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
func (r *RegisterMactscrType) GetTsevntena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsevntenaMask) != 0
}

// SetTsevntena TSEVNTENA
func (r *RegisterMactscrType) SetTsevntena(value bool) {
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
func (r *RegisterMactscrType) GetTsmstrena() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsmstrenaMask) != 0
}

// SetTsmstrena TSMSTRENA
func (r *RegisterMactscrType) SetTsmstrena(value bool) {
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
func (r *RegisterMactscrType) GetSnaptypsel() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldSnaptypselMask) >> RegisterMactscrFieldSnaptypselShift)
}

// SetSnaptypsel SNAPTYPSEL
func (r *RegisterMactscrType) SetSnaptypsel(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldSnaptypselMask)|(uint32(value)<<RegisterMactscrFieldSnaptypselShift))
}

const (
	RegisterMactscrFieldTsenmacaddrShift = 18
	RegisterMactscrFieldTsenmacaddrMask  = 0x40000
)

// GetTsenmacaddr TSENMACADDR
func (r *RegisterMactscrType) GetTsenmacaddr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTsenmacaddrMask) != 0
}

// SetTsenmacaddr TSENMACADDR
func (r *RegisterMactscrType) SetTsenmacaddr(value bool) {
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
func (r *RegisterMactscrType) GetCsc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldCscMask) != 0
}

const (
	RegisterMactscrFieldTxtsstsmShift = 24
	RegisterMactscrFieldTxtsstsmMask  = 0x1000000
)

// GetTxtsstsm TXTSSTSM
func (r *RegisterMactscrType) GetTxtsstsm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactscrFieldTxtsstsmMask) != 0
}

// SetTxtsstsm TXTSSTSM
func (r *RegisterMactscrType) SetTxtsstsm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactscrFieldTxtsstsmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactscrFieldTxtsstsmMask)
	}
}

// RegisterMacssirType Sub-second increment register
type RegisterMacssirType uint32

func (r *RegisterMacssirType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacssirType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacssirType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacssirType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacssirType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacssirFieldSnsincShift = 8
	RegisterMacssirFieldSnsincMask  = 0xff00
)

// GetSnsinc SNSINC
func (r *RegisterMacssirType) GetSnsinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacssirFieldSnsincMask) >> RegisterMacssirFieldSnsincShift)
}

// SetSnsinc SNSINC
func (r *RegisterMacssirType) SetSnsinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacssirFieldSnsincMask)|(uint32(value)<<RegisterMacssirFieldSnsincShift))
}

const (
	RegisterMacssirFieldSsincShift = 16
	RegisterMacssirFieldSsincMask  = 0xff0000
)

// GetSsinc SSINC
func (r *RegisterMacssirType) GetSsinc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacssirFieldSsincMask) >> RegisterMacssirFieldSsincShift)
}

// SetSsinc SSINC
func (r *RegisterMacssirType) SetSsinc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacssirFieldSsincMask)|(uint32(value)<<RegisterMacssirFieldSsincShift))
}

// RegisterMacstsrType System time seconds register
type RegisterMacstsrType uint32

func (r *RegisterMacstsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacstsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacstsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacstsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacstsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacstsrFieldTssShift = 0
	RegisterMacstsrFieldTssMask  = 0xffffffff
)

// GetTss TSS
func (r *RegisterMacstsrType) GetTss() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacstsrFieldTssMask) >> RegisterMacstsrFieldTssShift)
}

// SetTss TSS
func (r *RegisterMacstsrType) SetTss(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacstsrFieldTssMask)|(uint32(value)<<RegisterMacstsrFieldTssShift))
}

// RegisterMacstnrType System time nanoseconds register
type RegisterMacstnrType uint32

func (r *RegisterMacstnrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacstnrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacstnrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacstnrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacstnrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacstnrFieldTsssShift = 0
	RegisterMacstnrFieldTsssMask  = 0x7fffffff
)

// GetTsss TSSS
func (r *RegisterMacstnrType) GetTsss() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacstnrFieldTsssMask) >> RegisterMacstnrFieldTsssShift)
}

// SetTsss TSSS
func (r *RegisterMacstnrType) SetTsss(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacstnrFieldTsssMask)|(uint32(value)<<RegisterMacstnrFieldTsssShift))
}

// RegisterMacstsurType System time seconds update register
type RegisterMacstsurType uint32

func (r *RegisterMacstsurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacstsurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacstsurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacstsurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacstsurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacstsurFieldTssShift = 0
	RegisterMacstsurFieldTssMask  = 0xffffffff
)

// GetTss TSS
func (r *RegisterMacstsurType) GetTss() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacstsurFieldTssMask) >> RegisterMacstsurFieldTssShift)
}

// SetTss TSS
func (r *RegisterMacstsurType) SetTss(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacstsurFieldTssMask)|(uint32(value)<<RegisterMacstsurFieldTssShift))
}

// RegisterMacstnurType System time nanoseconds update register
type RegisterMacstnurType uint32

func (r *RegisterMacstnurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacstnurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacstnurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacstnurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacstnurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacstnurFieldTsssShift = 0
	RegisterMacstnurFieldTsssMask  = 0x7fffffff
)

// GetTsss TSSS
func (r *RegisterMacstnurType) GetTsss() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacstnurFieldTsssMask) >> RegisterMacstnurFieldTsssShift)
}

// SetTsss TSSS
func (r *RegisterMacstnurType) SetTsss(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacstnurFieldTsssMask)|(uint32(value)<<RegisterMacstnurFieldTsssShift))
}

const (
	RegisterMacstnurFieldAddsubShift = 31
	RegisterMacstnurFieldAddsubMask  = 0x80000000
)

// GetAddsub ADDSUB
func (r *RegisterMacstnurType) GetAddsub() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacstnurFieldAddsubMask) != 0
}

// SetAddsub ADDSUB
func (r *RegisterMacstnurType) SetAddsub(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacstnurFieldAddsubMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacstnurFieldAddsubMask)
	}
}

// RegisterMactsarType Timestamp addend register
type RegisterMactsarType uint32

func (r *RegisterMactsarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMactsarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMactsarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMactsarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMactsarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMactsarFieldTsarShift = 0
	RegisterMactsarFieldTsarMask  = 0xffffffff
)

// GetTsar TSAR
func (r *RegisterMactsarType) GetTsar() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactsarFieldTsarMask) >> RegisterMactsarFieldTsarShift)
}

// SetTsar TSAR
func (r *RegisterMactsarType) SetTsar(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactsarFieldTsarMask)|(uint32(value)<<RegisterMactsarFieldTsarShift))
}

// RegisterMactssrType Timestamp status register
type RegisterMactssrType uint32

func (r *RegisterMactssrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMactssrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMactssrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMactssrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMactssrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMactssrFieldTssovfShift = 0
	RegisterMactssrFieldTssovfMask  = 0x1
)

// GetTssovf TSSOVF
func (r *RegisterMactssrType) GetTssovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldTssovfMask) != 0
}

// SetTssovf TSSOVF
func (r *RegisterMactssrType) SetTssovf(value bool) {
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
func (r *RegisterMactssrType) GetTstargt0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldTstargt0Mask) != 0
}

// SetTstargt0 TSTARGT0
func (r *RegisterMactssrType) SetTstargt0(value bool) {
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
func (r *RegisterMactssrType) GetAuxtstrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldAuxtstrigMask) != 0
}

// SetAuxtstrig AUXTSTRIG
func (r *RegisterMactssrType) SetAuxtstrig(value bool) {
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
func (r *RegisterMactssrType) GetTstrgterr0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldTstrgterr0Mask) != 0
}

// SetTstrgterr0 TSTRGTERR0
func (r *RegisterMactssrType) SetTstrgterr0(value bool) {
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
func (r *RegisterMactssrType) GetTxtssis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldTxtssisMask) != 0
}

// SetTxtssis TXTSSIS
func (r *RegisterMactssrType) SetTxtssis(value bool) {
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
func (r *RegisterMactssrType) GetAtsstn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldAtsstnMask) >> RegisterMactssrFieldAtsstnShift)
}

// SetAtsstn ATSSTN
func (r *RegisterMactssrType) SetAtsstn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactssrFieldAtsstnMask)|(uint32(value)<<RegisterMactssrFieldAtsstnShift))
}

const (
	RegisterMactssrFieldAtsstmShift = 24
	RegisterMactssrFieldAtsstmMask  = 0x1000000
)

// GetAtsstm ATSSTM
func (r *RegisterMactssrType) GetAtsstm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldAtsstmMask) != 0
}

// SetAtsstm ATSSTM
func (r *RegisterMactssrType) SetAtsstm(value bool) {
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
func (r *RegisterMactssrType) GetAtsns() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMactssrFieldAtsnsMask) >> RegisterMactssrFieldAtsnsShift)
}

// SetAtsns ATSNS
func (r *RegisterMactssrType) SetAtsns(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactssrFieldAtsnsMask)|(uint32(value)<<RegisterMactssrFieldAtsnsShift))
}

// RegisterMactxtssnrType Tx timestamp status nanoseconds register
type RegisterMactxtssnrType uint32

func (r *RegisterMactxtssnrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMactxtssnrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMactxtssnrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMactxtssnrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMactxtssnrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMactxtssnrFieldTxtssloShift = 0
	RegisterMactxtssnrFieldTxtssloMask  = 0x7fffffff
)

// GetTxtsslo TXTSSLO
func (r *RegisterMactxtssnrType) GetTxtsslo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactxtssnrFieldTxtssloMask) >> RegisterMactxtssnrFieldTxtssloShift)
}

// SetTxtsslo TXTSSLO
func (r *RegisterMactxtssnrType) SetTxtsslo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactxtssnrFieldTxtssloMask)|(uint32(value)<<RegisterMactxtssnrFieldTxtssloShift))
}

const (
	RegisterMactxtssnrFieldTxtssmisShift = 31
	RegisterMactxtssnrFieldTxtssmisMask  = 0x80000000
)

// GetTxtssmis TXTSSMIS
func (r *RegisterMactxtssnrType) GetTxtssmis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMactxtssnrFieldTxtssmisMask) != 0
}

// SetTxtssmis TXTSSMIS
func (r *RegisterMactxtssnrType) SetTxtssmis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMactxtssnrFieldTxtssmisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMactxtssnrFieldTxtssmisMask)
	}
}

// RegisterMactxtsssrType Tx timestamp status seconds register
type RegisterMactxtsssrType uint32

func (r *RegisterMactxtsssrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMactxtsssrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMactxtsssrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMactxtsssrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMactxtsssrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMactxtsssrFieldTxtsshiShift = 0
	RegisterMactxtsssrFieldTxtsshiMask  = 0xffffffff
)

// GetTxtsshi TXTSSHI
func (r *RegisterMactxtsssrType) GetTxtsshi() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactxtsssrFieldTxtsshiMask) >> RegisterMactxtsssrFieldTxtsshiShift)
}

// SetTxtsshi TXTSSHI
func (r *RegisterMactxtsssrType) SetTxtsshi(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactxtsssrFieldTxtsshiMask)|(uint32(value)<<RegisterMactxtsssrFieldTxtsshiShift))
}

// RegisterMacacrType Auxiliary control register
type RegisterMacacrType uint32

func (r *RegisterMacacrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacacrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacacrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacacrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacacrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacacrFieldAtsfcShift = 0
	RegisterMacacrFieldAtsfcMask  = 0x1
)

// GetAtsfc ATSFC
func (r *RegisterMacacrType) GetAtsfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacacrFieldAtsfcMask) != 0
}

// SetAtsfc ATSFC
func (r *RegisterMacacrType) SetAtsfc(value bool) {
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
func (r *RegisterMacacrType) GetAtsen0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacacrFieldAtsen0Mask) != 0
}

// SetAtsen0 ATSEN0
func (r *RegisterMacacrType) SetAtsen0(value bool) {
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
func (r *RegisterMacacrType) GetAtsen1() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacacrFieldAtsen1Mask) != 0
}

// SetAtsen1 ATSEN1
func (r *RegisterMacacrType) SetAtsen1(value bool) {
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
func (r *RegisterMacacrType) GetAtsen2() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacacrFieldAtsen2Mask) != 0
}

// SetAtsen2 ATSEN2
func (r *RegisterMacacrType) SetAtsen2(value bool) {
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
func (r *RegisterMacacrType) GetAtsen3() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacacrFieldAtsen3Mask) != 0
}

// SetAtsen3 ATSEN3
func (r *RegisterMacacrType) SetAtsen3(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacacrFieldAtsen3Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacacrFieldAtsen3Mask)
	}
}

// RegisterMacatsnrType Auxiliary timestamp nanoseconds register
type RegisterMacatsnrType uint32

func (r *RegisterMacatsnrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacatsnrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacatsnrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacatsnrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacatsnrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacatsnrFieldAuxtsloShift = 0
	RegisterMacatsnrFieldAuxtsloMask  = 0x7fffffff
)

// GetAuxtslo AUXTSLO
func (r *RegisterMacatsnrType) GetAuxtslo() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacatsnrFieldAuxtsloMask) >> RegisterMacatsnrFieldAuxtsloShift)
}

// SetAuxtslo AUXTSLO
func (r *RegisterMacatsnrType) SetAuxtslo(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacatsnrFieldAuxtsloMask)|(uint32(value)<<RegisterMacatsnrFieldAuxtsloShift))
}

// RegisterMacatssrType Auxiliary timestamp seconds register
type RegisterMacatssrType uint32

func (r *RegisterMacatssrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacatssrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacatssrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacatssrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacatssrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacatssrFieldAuxtshiShift = 0
	RegisterMacatssrFieldAuxtshiMask  = 0xffffffff
)

// GetAuxtshi AUXTSHI
func (r *RegisterMacatssrType) GetAuxtshi() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacatssrFieldAuxtshiMask) >> RegisterMacatssrFieldAuxtshiShift)
}

// SetAuxtshi AUXTSHI
func (r *RegisterMacatssrType) SetAuxtshi(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacatssrFieldAuxtshiMask)|(uint32(value)<<RegisterMacatssrFieldAuxtshiShift))
}

// RegisterMactsiacrType Timestamp Ingress asymmetric correction register
type RegisterMactsiacrType uint32

func (r *RegisterMactsiacrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMactsiacrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMactsiacrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMactsiacrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMactsiacrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMactsiacrFieldOstiacShift = 0
	RegisterMactsiacrFieldOstiacMask  = 0xffffffff
)

// GetOstiac OSTIAC
func (r *RegisterMactsiacrType) GetOstiac() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactsiacrFieldOstiacMask) >> RegisterMactsiacrFieldOstiacShift)
}

// SetOstiac OSTIAC
func (r *RegisterMactsiacrType) SetOstiac(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactsiacrFieldOstiacMask)|(uint32(value)<<RegisterMactsiacrFieldOstiacShift))
}

// RegisterMactseacrType Timestamp Egress asymmetric correction register
type RegisterMactseacrType uint32

func (r *RegisterMactseacrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMactseacrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMactseacrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMactseacrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMactseacrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMactseacrFieldOsteacShift = 0
	RegisterMactseacrFieldOsteacMask  = 0xffffffff
)

// GetOsteac OSTEAC
func (r *RegisterMactseacrType) GetOsteac() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactseacrFieldOsteacMask) >> RegisterMactseacrFieldOsteacShift)
}

// SetOsteac OSTEAC
func (r *RegisterMactseacrType) SetOsteac(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactseacrFieldOsteacMask)|(uint32(value)<<RegisterMactseacrFieldOsteacShift))
}

// RegisterMactsicnrType Timestamp Ingress correction nanosecond register
type RegisterMactsicnrType uint32

func (r *RegisterMactsicnrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMactsicnrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMactsicnrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMactsicnrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMactsicnrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMactsicnrFieldTsicShift = 0
	RegisterMactsicnrFieldTsicMask  = 0xffffffff
)

// GetTsic TSIC
func (r *RegisterMactsicnrType) GetTsic() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactsicnrFieldTsicMask) >> RegisterMactsicnrFieldTsicShift)
}

// SetTsic TSIC
func (r *RegisterMactsicnrType) SetTsic(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactsicnrFieldTsicMask)|(uint32(value)<<RegisterMactsicnrFieldTsicShift))
}

// RegisterMactsecnrType Timestamp Egress correction nanosecond register
type RegisterMactsecnrType uint32

func (r *RegisterMactsecnrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMactsecnrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMactsecnrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMactsecnrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMactsecnrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMactsecnrFieldTsecShift = 0
	RegisterMactsecnrFieldTsecMask  = 0xffffffff
)

// GetTsec TSEC
func (r *RegisterMactsecnrType) GetTsec() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMactsecnrFieldTsecMask) >> RegisterMactsecnrFieldTsecShift)
}

// SetTsec TSEC
func (r *RegisterMactsecnrType) SetTsec(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMactsecnrFieldTsecMask)|(uint32(value)<<RegisterMactsecnrFieldTsecShift))
}

// RegisterMacppscrType PPS control register
type RegisterMacppscrType uint32

func (r *RegisterMacppscrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacppscrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacppscrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacppscrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacppscrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacppscrFieldPpsctrlShift = 0
	RegisterMacppscrFieldPpsctrlMask  = 0xf
)

// GetPpsctrl PPSCTRL
func (r *RegisterMacppscrType) GetPpsctrl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacppscrFieldPpsctrlMask) >> RegisterMacppscrFieldPpsctrlShift)
}

// SetPpsctrl PPSCTRL
func (r *RegisterMacppscrType) SetPpsctrl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppscrFieldPpsctrlMask)|(uint32(value)<<RegisterMacppscrFieldPpsctrlShift))
}

const (
	RegisterMacppscrFieldPpsen0Shift = 4
	RegisterMacppscrFieldPpsen0Mask  = 0x10
)

// GetPpsen0 PPSEN0
func (r *RegisterMacppscrType) GetPpsen0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacppscrFieldPpsen0Mask) != 0
}

// SetPpsen0 PPSEN0
func (r *RegisterMacppscrType) SetPpsen0(value bool) {
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
func (r *RegisterMacppscrType) GetTrgtmodsel0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacppscrFieldTrgtmodsel0Mask) >> RegisterMacppscrFieldTrgtmodsel0Shift)
}

// SetTrgtmodsel0 TRGTMODSEL0
func (r *RegisterMacppscrType) SetTrgtmodsel0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppscrFieldTrgtmodsel0Mask)|(uint32(value)<<RegisterMacppscrFieldTrgtmodsel0Shift))
}

// RegisterMacppsttsrType PPS target time seconds register
type RegisterMacppsttsrType uint32

func (r *RegisterMacppsttsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacppsttsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacppsttsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacppsttsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacppsttsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacppsttsrFieldTstrh0Shift = 0
	RegisterMacppsttsrFieldTstrh0Mask  = 0x7fffffff
)

// GetTstrh0 TSTRH0
func (r *RegisterMacppsttsrType) GetTstrh0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacppsttsrFieldTstrh0Mask) >> RegisterMacppsttsrFieldTstrh0Shift)
}

// SetTstrh0 TSTRH0
func (r *RegisterMacppsttsrType) SetTstrh0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppsttsrFieldTstrh0Mask)|(uint32(value)<<RegisterMacppsttsrFieldTstrh0Shift))
}

// RegisterMacppsttnrType PPS target time nanoseconds register
type RegisterMacppsttnrType uint32

func (r *RegisterMacppsttnrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacppsttnrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacppsttnrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacppsttnrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacppsttnrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacppsttnrFieldTtsl0Shift = 0
	RegisterMacppsttnrFieldTtsl0Mask  = 0x7fffffff
)

// GetTtsl0 TTSL0
func (r *RegisterMacppsttnrType) GetTtsl0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacppsttnrFieldTtsl0Mask) >> RegisterMacppsttnrFieldTtsl0Shift)
}

// SetTtsl0 TTSL0
func (r *RegisterMacppsttnrType) SetTtsl0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppsttnrFieldTtsl0Mask)|(uint32(value)<<RegisterMacppsttnrFieldTtsl0Shift))
}

const (
	RegisterMacppsttnrFieldTrgtbusy0Shift = 31
	RegisterMacppsttnrFieldTrgtbusy0Mask  = 0x80000000
)

// GetTrgtbusy0 TRGTBUSY0
func (r *RegisterMacppsttnrType) GetTrgtbusy0() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacppsttnrFieldTrgtbusy0Mask) != 0
}

// SetTrgtbusy0 TRGTBUSY0
func (r *RegisterMacppsttnrType) SetTrgtbusy0(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMacppsttnrFieldTrgtbusy0Mask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMacppsttnrFieldTrgtbusy0Mask)
	}
}

// RegisterMacppsirType PPS interval register
type RegisterMacppsirType uint32

func (r *RegisterMacppsirType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacppsirType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacppsirType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacppsirType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacppsirType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacppsirFieldPpsint0Shift = 0
	RegisterMacppsirFieldPpsint0Mask  = 0xffffffff
)

// GetPpsint0 PPSINT0
func (r *RegisterMacppsirType) GetPpsint0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacppsirFieldPpsint0Mask) >> RegisterMacppsirFieldPpsint0Shift)
}

// SetPpsint0 PPSINT0
func (r *RegisterMacppsirType) SetPpsint0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppsirFieldPpsint0Mask)|(uint32(value)<<RegisterMacppsirFieldPpsint0Shift))
}

// RegisterMacppswrType PPS width register
type RegisterMacppswrType uint32

func (r *RegisterMacppswrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacppswrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacppswrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacppswrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacppswrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacppswrFieldPpswidth0Shift = 0
	RegisterMacppswrFieldPpswidth0Mask  = 0xffffffff
)

// GetPpswidth0 PPSWIDTH0
func (r *RegisterMacppswrType) GetPpswidth0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacppswrFieldPpswidth0Mask) >> RegisterMacppswrFieldPpswidth0Shift)
}

// SetPpswidth0 PPSWIDTH0
func (r *RegisterMacppswrType) SetPpswidth0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacppswrFieldPpswidth0Mask)|(uint32(value)<<RegisterMacppswrFieldPpswidth0Shift))
}

// RegisterMacpocrType PTP Offload control register
type RegisterMacpocrType uint32

func (r *RegisterMacpocrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacpocrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacpocrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacpocrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacpocrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacpocrFieldPtoenShift = 0
	RegisterMacpocrFieldPtoenMask  = 0x1
)

// GetPtoen PTOEN
func (r *RegisterMacpocrType) GetPtoen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldPtoenMask) != 0
}

// SetPtoen PTOEN
func (r *RegisterMacpocrType) SetPtoen(value bool) {
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
func (r *RegisterMacpocrType) GetAsyncen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldAsyncenMask) != 0
}

// SetAsyncen ASYNCEN
func (r *RegisterMacpocrType) SetAsyncen(value bool) {
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
func (r *RegisterMacpocrType) GetApdreqen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldApdreqenMask) != 0
}

// SetApdreqen APDREQEN
func (r *RegisterMacpocrType) SetApdreqen(value bool) {
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
func (r *RegisterMacpocrType) GetAsynctrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldAsynctrigMask) != 0
}

// SetAsynctrig ASYNCTRIG
func (r *RegisterMacpocrType) SetAsynctrig(value bool) {
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
func (r *RegisterMacpocrType) GetApdreqtrig() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldApdreqtrigMask) != 0
}

// SetApdreqtrig APDREQTRIG
func (r *RegisterMacpocrType) SetApdreqtrig(value bool) {
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
func (r *RegisterMacpocrType) GetDrrdis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldDrrdisMask) != 0
}

// SetDrrdis DRRDIS
func (r *RegisterMacpocrType) SetDrrdis(value bool) {
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
func (r *RegisterMacpocrType) GetDn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMacpocrFieldDnMask) >> RegisterMacpocrFieldDnShift)
}

// SetDn DN
func (r *RegisterMacpocrType) SetDn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacpocrFieldDnMask)|(uint32(value)<<RegisterMacpocrFieldDnShift))
}

// RegisterMacspi0rType PTP Source Port Identity 0 Register
type RegisterMacspi0rType uint32

func (r *RegisterMacspi0rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacspi0rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacspi0rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacspi0rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacspi0rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacspi0rFieldSpi0Shift = 0
	RegisterMacspi0rFieldSpi0Mask  = 0xffffffff
)

// GetSpi0 SPI0
func (r *RegisterMacspi0rType) GetSpi0() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacspi0rFieldSpi0Mask) >> RegisterMacspi0rFieldSpi0Shift)
}

// SetSpi0 SPI0
func (r *RegisterMacspi0rType) SetSpi0(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacspi0rFieldSpi0Mask)|(uint32(value)<<RegisterMacspi0rFieldSpi0Shift))
}

// RegisterMacspi1rType PTP Source port identity 1 register
type RegisterMacspi1rType uint32

func (r *RegisterMacspi1rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacspi1rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacspi1rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacspi1rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacspi1rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacspi1rFieldSpi1Shift = 0
	RegisterMacspi1rFieldSpi1Mask  = 0xffffffff
)

// GetSpi1 SPI1
func (r *RegisterMacspi1rType) GetSpi1() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterMacspi1rFieldSpi1Mask) >> RegisterMacspi1rFieldSpi1Shift)
}

// SetSpi1 SPI1
func (r *RegisterMacspi1rType) SetSpi1(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacspi1rFieldSpi1Mask)|(uint32(value)<<RegisterMacspi1rFieldSpi1Shift))
}

// RegisterMacspi2rType PTP Source port identity 2 register
type RegisterMacspi2rType uint32

func (r *RegisterMacspi2rType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMacspi2rType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMacspi2rType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMacspi2rType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMacspi2rType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMacspi2rFieldSpi2Shift = 0
	RegisterMacspi2rFieldSpi2Mask  = 0xffff
)

// GetSpi2 SPI2
func (r *RegisterMacspi2rType) GetSpi2() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMacspi2rFieldSpi2Mask) >> RegisterMacspi2rFieldSpi2Shift)
}

// SetSpi2 SPI2
func (r *RegisterMacspi2rType) SetSpi2(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMacspi2rFieldSpi2Mask)|(uint32(value)<<RegisterMacspi2rFieldSpi2Shift))
}

// RegisterMaclmirType Log message interval register
type RegisterMaclmirType uint32

func (r *RegisterMaclmirType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMaclmirType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMaclmirType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMaclmirType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMaclmirType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMaclmirFieldLsiShift = 0
	RegisterMaclmirFieldLsiMask  = 0xff
)

// GetLsi LSI
func (r *RegisterMaclmirType) GetLsi() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaclmirFieldLsiMask) >> RegisterMaclmirFieldLsiShift)
}

// SetLsi LSI
func (r *RegisterMaclmirType) SetLsi(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaclmirFieldLsiMask)|(uint32(value)<<RegisterMaclmirFieldLsiShift))
}

const (
	RegisterMaclmirFieldDrsyncrShift = 8
	RegisterMaclmirFieldDrsyncrMask  = 0x700
)

// GetDrsyncr DRSYNCR
func (r *RegisterMaclmirType) GetDrsyncr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaclmirFieldDrsyncrMask) >> RegisterMaclmirFieldDrsyncrShift)
}

// SetDrsyncr DRSYNCR
func (r *RegisterMaclmirType) SetDrsyncr(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaclmirFieldDrsyncrMask)|(uint32(value)<<RegisterMaclmirFieldDrsyncrShift))
}

const (
	RegisterMaclmirFieldLmpdriShift = 24
	RegisterMaclmirFieldLmpdriMask  = 0xff000000
)

// GetLmpdri LMPDRI
func (r *RegisterMaclmirType) GetLmpdri() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMaclmirFieldLmpdriMask) >> RegisterMaclmirFieldLmpdriShift)
}

// SetLmpdri LMPDRI
func (r *RegisterMaclmirType) SetLmpdri(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMaclmirFieldLmpdriMask)|(uint32(value)<<RegisterMaclmirFieldLmpdriShift))
}

// RegisterMtlomrType Operating mode Register
type RegisterMtlomrType uint32

func (r *RegisterMtlomrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMtlomrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMtlomrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMtlomrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMtlomrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMtlomrFieldDtxstsShift = 1
	RegisterMtlomrFieldDtxstsMask  = 0x2
)

// GetDtxsts DTXSTS
func (r *RegisterMtlomrType) GetDtxsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlomrFieldDtxstsMask) != 0
}

// SetDtxsts DTXSTS
func (r *RegisterMtlomrType) SetDtxsts(value bool) {
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
func (r *RegisterMtlomrType) GetCntprst() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlomrFieldCntprstMask) != 0
}

// SetCntprst CNTPRST
func (r *RegisterMtlomrType) SetCntprst(value bool) {
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
func (r *RegisterMtlomrType) GetCntclr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlomrFieldCntclrMask) != 0
}

// SetCntclr CNTCLR
func (r *RegisterMtlomrType) SetCntclr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlomrFieldCntclrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlomrFieldCntclrMask)
	}
}

// RegisterMtlisrType Interrupt status Register
type RegisterMtlisrType uint32

func (r *RegisterMtlisrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMtlisrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMtlisrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMtlisrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMtlisrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMtlisrFieldQ0isShift = 0
	RegisterMtlisrFieldQ0isMask  = 0x1
)

// GetQ0is Queue interrupt status
func (r *RegisterMtlisrType) GetQ0is() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlisrFieldQ0isMask) != 0
}

// SetQ0is Queue interrupt status
func (r *RegisterMtlisrType) SetQ0is(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlisrFieldQ0isMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlisrFieldQ0isMask)
	}
}

// RegisterMtltxqomrType Tx queue operating mode Register
type RegisterMtltxqomrType uint32

func (r *RegisterMtltxqomrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMtltxqomrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMtltxqomrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMtltxqomrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMtltxqomrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMtltxqomrFieldFtqShift = 0
	RegisterMtltxqomrFieldFtqMask  = 0x1
)

// GetFtq Flush Transmit Queue
func (r *RegisterMtltxqomrType) GetFtq() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqomrFieldFtqMask) != 0
}

// SetFtq Flush Transmit Queue
func (r *RegisterMtltxqomrType) SetFtq(value bool) {
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
func (r *RegisterMtltxqomrType) GetTsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqomrFieldTsfMask) != 0
}

// SetTsf Transmit Store and Forward
func (r *RegisterMtltxqomrType) SetTsf(value bool) {
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
func (r *RegisterMtltxqomrType) GetTxqen() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqomrFieldTxqenMask) >> RegisterMtltxqomrFieldTxqenShift)
}

const (
	RegisterMtltxqomrFieldTtcShift = 4
	RegisterMtltxqomrFieldTtcMask  = 0x70
)

// GetTtc Transmit Threshold Control
func (r *RegisterMtltxqomrType) GetTtc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqomrFieldTtcMask) >> RegisterMtltxqomrFieldTtcShift)
}

// SetTtc Transmit Threshold Control
func (r *RegisterMtltxqomrType) SetTtc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqomrFieldTtcMask)|(uint32(value)<<RegisterMtltxqomrFieldTtcShift))
}

const (
	RegisterMtltxqomrFieldTqsShift = 16
	RegisterMtltxqomrFieldTqsMask  = 0x1ff0000
)

// GetTqs Transmit Queue Size
func (r *RegisterMtltxqomrType) GetTqs() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqomrFieldTqsMask) >> RegisterMtltxqomrFieldTqsShift)
}

// SetTqs Transmit Queue Size
func (r *RegisterMtltxqomrType) SetTqs(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqomrFieldTqsMask)|(uint32(value)<<RegisterMtltxqomrFieldTqsShift))
}

// RegisterMtltxqurType Tx queue underflow register
type RegisterMtltxqurType uint32

func (r *RegisterMtltxqurType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMtltxqurType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMtltxqurType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMtltxqurType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMtltxqurType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMtltxqurFieldUffrmcntShift = 0
	RegisterMtltxqurFieldUffrmcntMask  = 0x7ff
)

// GetUffrmcnt Underflow Packet Counter
func (r *RegisterMtltxqurType) GetUffrmcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqurFieldUffrmcntMask) >> RegisterMtltxqurFieldUffrmcntShift)
}

// SetUffrmcnt Underflow Packet Counter
func (r *RegisterMtltxqurType) SetUffrmcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqurFieldUffrmcntMask)|(uint32(value)<<RegisterMtltxqurFieldUffrmcntShift))
}

const (
	RegisterMtltxqurFieldUfcntovfShift = 11
	RegisterMtltxqurFieldUfcntovfMask  = 0x800
)

// GetUfcntovf UFCNTOVF
func (r *RegisterMtltxqurType) GetUfcntovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqurFieldUfcntovfMask) != 0
}

// SetUfcntovf UFCNTOVF
func (r *RegisterMtltxqurType) SetUfcntovf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtltxqurFieldUfcntovfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqurFieldUfcntovfMask)
	}
}

// RegisterMtltxqdrType Tx queue debug Register
type RegisterMtltxqdrType uint32

func (r *RegisterMtltxqdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMtltxqdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMtltxqdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMtltxqdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMtltxqdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMtltxqdrFieldTxqpausedShift = 0
	RegisterMtltxqdrFieldTxqpausedMask  = 0x1
)

// GetTxqpaused TXQPAUSED
func (r *RegisterMtltxqdrType) GetTxqpaused() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldTxqpausedMask) != 0
}

// SetTxqpaused TXQPAUSED
func (r *RegisterMtltxqdrType) SetTxqpaused(value bool) {
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
func (r *RegisterMtltxqdrType) GetTrcsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldTrcstsMask) >> RegisterMtltxqdrFieldTrcstsShift)
}

// SetTrcsts TRCSTS
func (r *RegisterMtltxqdrType) SetTrcsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqdrFieldTrcstsMask)|(uint32(value)<<RegisterMtltxqdrFieldTrcstsShift))
}

const (
	RegisterMtltxqdrFieldTwcstsShift = 3
	RegisterMtltxqdrFieldTwcstsMask  = 0x8
)

// GetTwcsts TWCSTS
func (r *RegisterMtltxqdrType) GetTwcsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldTwcstsMask) != 0
}

// SetTwcsts TWCSTS
func (r *RegisterMtltxqdrType) SetTwcsts(value bool) {
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
func (r *RegisterMtltxqdrType) GetTxqsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldTxqstsMask) != 0
}

// SetTxqsts TXQSTS
func (r *RegisterMtltxqdrType) SetTxqsts(value bool) {
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
func (r *RegisterMtltxqdrType) GetTxstsfsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldTxstsfstsMask) != 0
}

// SetTxstsfsts TXSTSFSTS
func (r *RegisterMtltxqdrType) SetTxstsfsts(value bool) {
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
func (r *RegisterMtltxqdrType) GetPtxq() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldPtxqMask) >> RegisterMtltxqdrFieldPtxqShift)
}

// SetPtxq PTXQ
func (r *RegisterMtltxqdrType) SetPtxq(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqdrFieldPtxqMask)|(uint32(value)<<RegisterMtltxqdrFieldPtxqShift))
}

const (
	RegisterMtltxqdrFieldStxstsfShift = 20
	RegisterMtltxqdrFieldStxstsfMask  = 0x700000
)

// GetStxstsf STXSTSF
func (r *RegisterMtltxqdrType) GetStxstsf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqdrFieldStxstsfMask) >> RegisterMtltxqdrFieldStxstsfShift)
}

// SetStxstsf STXSTSF
func (r *RegisterMtltxqdrType) SetStxstsf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtltxqdrFieldStxstsfMask)|(uint32(value)<<RegisterMtltxqdrFieldStxstsfShift))
}

// RegisterMtlqicsrType Queue interrupt control status Register
type RegisterMtlqicsrType uint32

func (r *RegisterMtlqicsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMtlqicsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMtlqicsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMtlqicsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMtlqicsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMtlqicsrFieldTxunfisShift = 0
	RegisterMtlqicsrFieldTxunfisMask  = 0x1
)

// GetTxunfis TXUNFIS
func (r *RegisterMtlqicsrType) GetTxunfis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlqicsrFieldTxunfisMask) != 0
}

// SetTxunfis TXUNFIS
func (r *RegisterMtlqicsrType) SetTxunfis(value bool) {
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
func (r *RegisterMtlqicsrType) GetTxuie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlqicsrFieldTxuieMask) != 0
}

// SetTxuie TXUIE
func (r *RegisterMtlqicsrType) SetTxuie(value bool) {
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
func (r *RegisterMtlqicsrType) GetRxovfis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlqicsrFieldRxovfisMask) != 0
}

// SetRxovfis RXOVFIS
func (r *RegisterMtlqicsrType) SetRxovfis(value bool) {
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
func (r *RegisterMtlqicsrType) GetRxoie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlqicsrFieldRxoieMask) != 0
}

// SetRxoie RXOIE
func (r *RegisterMtlqicsrType) SetRxoie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlqicsrFieldRxoieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlqicsrFieldRxoieMask)
	}
}

// RegisterMtlrxqomrType Rx queue operating mode register
type RegisterMtlrxqomrType uint32

func (r *RegisterMtlrxqomrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMtlrxqomrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMtlrxqomrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMtlrxqomrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMtlrxqomrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMtlrxqomrFieldRtcShift = 0
	RegisterMtlrxqomrFieldRtcMask  = 0x3
)

// GetRtc RTC
func (r *RegisterMtlrxqomrType) GetRtc() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldRtcMask) >> RegisterMtlrxqomrFieldRtcShift)
}

// SetRtc RTC
func (r *RegisterMtlrxqomrType) SetRtc(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldRtcMask)|(uint32(value)<<RegisterMtlrxqomrFieldRtcShift))
}

const (
	RegisterMtlrxqomrFieldFupShift = 3
	RegisterMtlrxqomrFieldFupMask  = 0x8
)

// GetFup FUP
func (r *RegisterMtlrxqomrType) GetFup() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldFupMask) != 0
}

// SetFup FUP
func (r *RegisterMtlrxqomrType) SetFup(value bool) {
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
func (r *RegisterMtlrxqomrType) GetFep() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldFepMask) != 0
}

// SetFep FEP
func (r *RegisterMtlrxqomrType) SetFep(value bool) {
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
func (r *RegisterMtlrxqomrType) GetRsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldRsfMask) != 0
}

// SetRsf RSF
func (r *RegisterMtlrxqomrType) SetRsf(value bool) {
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
func (r *RegisterMtlrxqomrType) GetDistcpef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldDistcpefMask) != 0
}

// SetDistcpef DIS_TCP_EF
func (r *RegisterMtlrxqomrType) SetDistcpef(value bool) {
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
func (r *RegisterMtlrxqomrType) GetEhfc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldEhfcMask) != 0
}

// SetEhfc EHFC
func (r *RegisterMtlrxqomrType) SetEhfc(value bool) {
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
func (r *RegisterMtlrxqomrType) GetRfa() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldRfaMask) >> RegisterMtlrxqomrFieldRfaShift)
}

// SetRfa RFA
func (r *RegisterMtlrxqomrType) SetRfa(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldRfaMask)|(uint32(value)<<RegisterMtlrxqomrFieldRfaShift))
}

const (
	RegisterMtlrxqomrFieldRfdShift = 14
	RegisterMtlrxqomrFieldRfdMask  = 0x1c000
)

// GetRfd RFD
func (r *RegisterMtlrxqomrType) GetRfd() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldRfdMask) >> RegisterMtlrxqomrFieldRfdShift)
}

// SetRfd RFD
func (r *RegisterMtlrxqomrType) SetRfd(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldRfdMask)|(uint32(value)<<RegisterMtlrxqomrFieldRfdShift))
}

const (
	RegisterMtlrxqomrFieldRqsShift = 20
	RegisterMtlrxqomrFieldRqsMask  = 0x700000
)

// GetRqs RQS
func (r *RegisterMtlrxqomrType) GetRqs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldRqsMask) >> RegisterMtlrxqomrFieldRqsShift)
}

// RegisterMtlrxqmpocrType Rx queue missed packet and overflow counter register
type RegisterMtlrxqmpocrType uint32

func (r *RegisterMtlrxqmpocrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMtlrxqmpocrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMtlrxqmpocrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMtlrxqmpocrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMtlrxqmpocrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMtlrxqmpocrFieldOvfpktcntShift = 0
	RegisterMtlrxqmpocrFieldOvfpktcntMask  = 0x7ff
)

// GetOvfpktcnt OVFPKTCNT
func (r *RegisterMtlrxqmpocrType) GetOvfpktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqmpocrFieldOvfpktcntMask) >> RegisterMtlrxqmpocrFieldOvfpktcntShift)
}

// SetOvfpktcnt OVFPKTCNT
func (r *RegisterMtlrxqmpocrType) SetOvfpktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqmpocrFieldOvfpktcntMask)|(uint32(value)<<RegisterMtlrxqmpocrFieldOvfpktcntShift))
}

const (
	RegisterMtlrxqmpocrFieldOvfcntovfShift = 11
	RegisterMtlrxqmpocrFieldOvfcntovfMask  = 0x800
)

// GetOvfcntovf OVFCNTOVF
func (r *RegisterMtlrxqmpocrType) GetOvfcntovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqmpocrFieldOvfcntovfMask) != 0
}

// SetOvfcntovf OVFCNTOVF
func (r *RegisterMtlrxqmpocrType) SetOvfcntovf(value bool) {
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
func (r *RegisterMtlrxqmpocrType) GetMispktcnt() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqmpocrFieldMispktcntMask) >> RegisterMtlrxqmpocrFieldMispktcntShift)
}

// SetMispktcnt MISPKTCNT
func (r *RegisterMtlrxqmpocrType) SetMispktcnt(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqmpocrFieldMispktcntMask)|(uint32(value)<<RegisterMtlrxqmpocrFieldMispktcntShift))
}

const (
	RegisterMtlrxqmpocrFieldMiscntovfShift = 27
	RegisterMtlrxqmpocrFieldMiscntovfMask  = 0x8000000
)

// GetMiscntovf MISCNTOVF
func (r *RegisterMtlrxqmpocrType) GetMiscntovf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqmpocrFieldMiscntovfMask) != 0
}

// SetMiscntovf MISCNTOVF
func (r *RegisterMtlrxqmpocrType) SetMiscntovf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlrxqmpocrFieldMiscntovfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqmpocrFieldMiscntovfMask)
	}
}

// RegisterMtlrxqdrType Rx queue debug register
type RegisterMtlrxqdrType uint32

func (r *RegisterMtlrxqdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterMtlrxqdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterMtlrxqdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterMtlrxqdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterMtlrxqdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterMtlrxqdrFieldRwcstsShift = 0
	RegisterMtlrxqdrFieldRwcstsMask  = 0x1
)

// GetRwcsts RWCSTS
func (r *RegisterMtlrxqdrType) GetRwcsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqdrFieldRwcstsMask) != 0
}

// SetRwcsts RWCSTS
func (r *RegisterMtlrxqdrType) SetRwcsts(value bool) {
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
func (r *RegisterMtlrxqdrType) GetRrcsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqdrFieldRrcstsMask) >> RegisterMtlrxqdrFieldRrcstsShift)
}

// SetRrcsts RRCSTS
func (r *RegisterMtlrxqdrType) SetRrcsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqdrFieldRrcstsMask)|(uint32(value)<<RegisterMtlrxqdrFieldRrcstsShift))
}

const (
	RegisterMtlrxqdrFieldRxqstsShift = 4
	RegisterMtlrxqdrFieldRxqstsMask  = 0x30
)

// GetRxqsts RXQSTS
func (r *RegisterMtlrxqdrType) GetRxqsts() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqdrFieldRxqstsMask) >> RegisterMtlrxqdrFieldRxqstsShift)
}

// SetRxqsts RXQSTS
func (r *RegisterMtlrxqdrType) SetRxqsts(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqdrFieldRxqstsMask)|(uint32(value)<<RegisterMtlrxqdrFieldRxqstsShift))
}

const (
	RegisterMtlrxqdrFieldPrxqShift = 16
	RegisterMtlrxqdrFieldPrxqMask  = 0x3fff0000
)

// GetPrxq PRXQ
func (r *RegisterMtlrxqdrType) GetPrxq() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqdrFieldPrxqMask) >> RegisterMtlrxqdrFieldPrxqShift)
}

// SetPrxq PRXQ
func (r *RegisterMtlrxqdrType) SetPrxq(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqdrFieldPrxqMask)|(uint32(value)<<RegisterMtlrxqdrFieldPrxqShift))
}

// RegisterDmamrType DMA mode register
type RegisterDmamrType uint32

func (r *RegisterDmamrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmamrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmamrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmamrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmamrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmamrFieldSwrShift = 0
	RegisterDmamrFieldSwrMask  = 0x1
)

// GetSwr Software Reset
func (r *RegisterDmamrType) GetSwr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmamrFieldSwrMask) != 0
}

// SetSwr Software Reset
func (r *RegisterDmamrType) SetSwr(value bool) {
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
func (r *RegisterDmamrType) GetDa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmamrFieldDaMask) != 0
}

const (
	RegisterDmamrFieldTxprShift = 11
	RegisterDmamrFieldTxprMask  = 0x800
)

// GetTxpr Transmit priority
func (r *RegisterDmamrType) GetTxpr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmamrFieldTxprMask) != 0
}

const (
	RegisterDmamrFieldPrShift = 12
	RegisterDmamrFieldPrMask  = 0x7000
)

// GetPr Priority ratio
func (r *RegisterDmamrType) GetPr() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmamrFieldPrMask) >> RegisterDmamrFieldPrShift)
}

const (
	RegisterDmamrFieldIntmShift = 16
	RegisterDmamrFieldIntmMask  = 0x10000
)

// GetIntm Interrupt Mode
func (r *RegisterDmamrType) GetIntm() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmamrFieldIntmMask) != 0
}

// SetIntm Interrupt Mode
func (r *RegisterDmamrType) SetIntm(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmamrFieldIntmMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmamrFieldIntmMask)
	}
}

// RegisterDmasbmrType System bus mode register
type RegisterDmasbmrType uint32

func (r *RegisterDmasbmrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmasbmrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmasbmrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmasbmrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmasbmrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmasbmrFieldFbShift = 0
	RegisterDmasbmrFieldFbMask  = 0x1
)

// GetFb Fixed Burst Length
func (r *RegisterDmasbmrType) GetFb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmasbmrFieldFbMask) != 0
}

// SetFb Fixed Burst Length
func (r *RegisterDmasbmrType) SetFb(value bool) {
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
func (r *RegisterDmasbmrType) GetAal() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmasbmrFieldAalMask) != 0
}

// SetAal Address-Aligned Beats
func (r *RegisterDmasbmrType) SetAal(value bool) {
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
func (r *RegisterDmasbmrType) GetMb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmasbmrFieldMbMask) != 0
}

const (
	RegisterDmasbmrFieldRbShift = 15
	RegisterDmasbmrFieldRbMask  = 0x8000
)

// GetRb Rebuild INCRx Burst
func (r *RegisterDmasbmrType) GetRb() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmasbmrFieldRbMask) != 0
}

// RegisterDmaisrType Interrupt status register
type RegisterDmaisrType uint32

func (r *RegisterDmaisrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmaisrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmaisrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmaisrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmaisrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmaisrFieldDc0isShift = 0
	RegisterDmaisrFieldDc0isMask  = 0x1
)

// GetDc0is DMA Channel Interrupt Status
func (r *RegisterDmaisrType) GetDc0is() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmaisrFieldDc0isMask) != 0
}

// SetDc0is DMA Channel Interrupt Status
func (r *RegisterDmaisrType) SetDc0is(value bool) {
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
func (r *RegisterDmaisrType) GetMtlis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmaisrFieldMtlisMask) != 0
}

// SetMtlis MTL Interrupt Status
func (r *RegisterDmaisrType) SetMtlis(value bool) {
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
func (r *RegisterDmaisrType) GetMacis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmaisrFieldMacisMask) != 0
}

// SetMacis MAC Interrupt Status
func (r *RegisterDmaisrType) SetMacis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmaisrFieldMacisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmaisrFieldMacisMask)
	}
}

// RegisterDmadsrType Debug status register
type RegisterDmadsrType uint32

func (r *RegisterDmadsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmadsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmadsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmadsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmadsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmadsrFieldAxwhstsShift = 0
	RegisterDmadsrFieldAxwhstsMask  = 0x1
)

// GetAxwhsts AHB Master Write Channel
func (r *RegisterDmadsrType) GetAxwhsts() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmadsrFieldAxwhstsMask) != 0
}

// SetAxwhsts AHB Master Write Channel
func (r *RegisterDmadsrType) SetAxwhsts(value bool) {
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
func (r *RegisterDmadsrType) GetRps0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmadsrFieldRps0Mask) >> RegisterDmadsrFieldRps0Shift)
}

// SetRps0 DMA Channel Receive Process State
func (r *RegisterDmadsrType) SetRps0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmadsrFieldRps0Mask)|(uint32(value)<<RegisterDmadsrFieldRps0Shift))
}

const (
	RegisterDmadsrFieldTps0Shift = 12
	RegisterDmadsrFieldTps0Mask  = 0xf000
)

// GetTps0 DMA Channel Transmit Process State
func (r *RegisterDmadsrType) GetTps0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmadsrFieldTps0Mask) >> RegisterDmadsrFieldTps0Shift)
}

// SetTps0 DMA Channel Transmit Process State
func (r *RegisterDmadsrType) SetTps0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmadsrFieldTps0Mask)|(uint32(value)<<RegisterDmadsrFieldTps0Shift))
}

// RegisterDmaccrType Channel control register
type RegisterDmaccrType uint32

func (r *RegisterDmaccrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmaccrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmaccrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmaccrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmaccrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmaccrFieldMssShift = 0
	RegisterDmaccrFieldMssMask  = 0x3fff
)

// GetMss Maximum Segment Size
func (r *RegisterDmaccrType) GetMss() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccrFieldMssMask) >> RegisterDmaccrFieldMssShift)
}

// SetMss Maximum Segment Size
func (r *RegisterDmaccrType) SetMss(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccrFieldMssMask)|(uint32(value)<<RegisterDmaccrFieldMssShift))
}

const (
	RegisterDmaccrFieldPblx8Shift = 16
	RegisterDmaccrFieldPblx8Mask  = 0x10000
)

// GetPblx8 8xPBL mode
func (r *RegisterDmaccrType) GetPblx8() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmaccrFieldPblx8Mask) != 0
}

// SetPblx8 8xPBL mode
func (r *RegisterDmaccrType) SetPblx8(value bool) {
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
func (r *RegisterDmaccrType) GetDsl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccrFieldDslMask) >> RegisterDmaccrFieldDslShift)
}

// SetDsl Descriptor Skip Length
func (r *RegisterDmaccrType) SetDsl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccrFieldDslMask)|(uint32(value)<<RegisterDmaccrFieldDslShift))
}

// RegisterDmactxcrType Channel transmit control register
type RegisterDmactxcrType uint32

func (r *RegisterDmactxcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmactxcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmactxcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmactxcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmactxcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmactxcrFieldStShift = 0
	RegisterDmactxcrFieldStMask  = 0x1
)

// GetSt Start or Stop Transmission Command
func (r *RegisterDmactxcrType) GetSt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmactxcrFieldStMask) != 0
}

// SetSt Start or Stop Transmission Command
func (r *RegisterDmactxcrType) SetSt(value bool) {
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
func (r *RegisterDmactxcrType) GetOsf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmactxcrFieldOsfMask) != 0
}

// SetOsf Operate on Second Packet
func (r *RegisterDmactxcrType) SetOsf(value bool) {
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
func (r *RegisterDmactxcrType) GetTse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmactxcrFieldTseMask) != 0
}

// SetTse TCP Segmentation Enabled
func (r *RegisterDmactxcrType) SetTse(value bool) {
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
func (r *RegisterDmactxcrType) GetTxpbl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmactxcrFieldTxpblMask) >> RegisterDmactxcrFieldTxpblShift)
}

// SetTxpbl Transmit Programmable Burst Length
func (r *RegisterDmactxcrType) SetTxpbl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmactxcrFieldTxpblMask)|(uint32(value)<<RegisterDmactxcrFieldTxpblShift))
}

// RegisterDmacrxcrType Channel receive control register
type RegisterDmacrxcrType uint32

func (r *RegisterDmacrxcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmacrxcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmacrxcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmacrxcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmacrxcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmacrxcrFieldSrShift = 0
	RegisterDmacrxcrFieldSrMask  = 0x1
)

// GetSr Start or Stop Receive Command
func (r *RegisterDmacrxcrType) GetSr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxcrFieldSrMask) != 0
}

// SetSr Start or Stop Receive Command
func (r *RegisterDmacrxcrType) SetSr(value bool) {
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
func (r *RegisterDmacrxcrType) GetRbsz() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxcrFieldRbszMask) >> RegisterDmacrxcrFieldRbszShift)
}

// SetRbsz Receive Buffer size
func (r *RegisterDmacrxcrType) SetRbsz(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxcrFieldRbszMask)|(uint32(value)<<RegisterDmacrxcrFieldRbszShift))
}

const (
	RegisterDmacrxcrFieldRxpblShift = 16
	RegisterDmacrxcrFieldRxpblMask  = 0x3f0000
)

// GetRxpbl RXPBL
func (r *RegisterDmacrxcrType) GetRxpbl() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxcrFieldRxpblMask) >> RegisterDmacrxcrFieldRxpblShift)
}

// SetRxpbl RXPBL
func (r *RegisterDmacrxcrType) SetRxpbl(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxcrFieldRxpblMask)|(uint32(value)<<RegisterDmacrxcrFieldRxpblShift))
}

const (
	RegisterDmacrxcrFieldRpfShift = 31
	RegisterDmacrxcrFieldRpfMask  = 0x80000000
)

// GetRpf DMA Rx Channel Packet Flush
func (r *RegisterDmacrxcrType) GetRpf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxcrFieldRpfMask) != 0
}

// SetRpf DMA Rx Channel Packet Flush
func (r *RegisterDmacrxcrType) SetRpf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacrxcrFieldRpfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxcrFieldRpfMask)
	}
}

// RegisterDmactxdlarType Channel Tx descriptor list address register
type RegisterDmactxdlarType uint32

func (r *RegisterDmactxdlarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmactxdlarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmactxdlarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmactxdlarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmactxdlarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmactxdlarFieldTdeslaShift = 2
	RegisterDmactxdlarFieldTdeslaMask  = 0xfffffffc
)

// GetTdesla Start of Transmit List
func (r *RegisterDmactxdlarType) GetTdesla() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmactxdlarFieldTdeslaMask) >> RegisterDmactxdlarFieldTdeslaShift)
}

// SetTdesla Start of Transmit List
func (r *RegisterDmactxdlarType) SetTdesla(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmactxdlarFieldTdeslaMask)|(uint32(value)<<RegisterDmactxdlarFieldTdeslaShift))
}

// RegisterDmacrxdlarType Channel Rx descriptor list address register
type RegisterDmacrxdlarType uint32

func (r *RegisterDmacrxdlarType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmacrxdlarType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmacrxdlarType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmacrxdlarType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmacrxdlarType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmacrxdlarFieldRdeslaShift = 2
	RegisterDmacrxdlarFieldRdeslaMask  = 0xfffffffc
)

// GetRdesla Start of Receive List
func (r *RegisterDmacrxdlarType) GetRdesla() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxdlarFieldRdeslaMask) >> RegisterDmacrxdlarFieldRdeslaShift)
}

// SetRdesla Start of Receive List
func (r *RegisterDmacrxdlarType) SetRdesla(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxdlarFieldRdeslaMask)|(uint32(value)<<RegisterDmacrxdlarFieldRdeslaShift))
}

// RegisterDmactxdtprType Channel Tx descriptor tail pointer register
type RegisterDmactxdtprType uint32

func (r *RegisterDmactxdtprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmactxdtprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmactxdtprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmactxdtprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmactxdtprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmactxdtprFieldTdtShift = 2
	RegisterDmactxdtprFieldTdtMask  = 0xfffffffc
)

// GetTdt Transmit Descriptor Tail Pointer
func (r *RegisterDmactxdtprType) GetTdt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmactxdtprFieldTdtMask) >> RegisterDmactxdtprFieldTdtShift)
}

// SetTdt Transmit Descriptor Tail Pointer
func (r *RegisterDmactxdtprType) SetTdt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmactxdtprFieldTdtMask)|(uint32(value)<<RegisterDmactxdtprFieldTdtShift))
}

// RegisterDmacrxdtprType Channel Rx descriptor tail pointer register
type RegisterDmacrxdtprType uint32

func (r *RegisterDmacrxdtprType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmacrxdtprType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmacrxdtprType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmacrxdtprType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmacrxdtprType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmacrxdtprFieldRdtShift = 2
	RegisterDmacrxdtprFieldRdtMask  = 0xfffffffc
)

// GetRdt Receive Descriptor Tail Pointer
func (r *RegisterDmacrxdtprType) GetRdt() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxdtprFieldRdtMask) >> RegisterDmacrxdtprFieldRdtShift)
}

// SetRdt Receive Descriptor Tail Pointer
func (r *RegisterDmacrxdtprType) SetRdt(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxdtprFieldRdtMask)|(uint32(value)<<RegisterDmacrxdtprFieldRdtShift))
}

// RegisterDmactxrlrType Channel Tx descriptor ring length register
type RegisterDmactxrlrType uint32

func (r *RegisterDmactxrlrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmactxrlrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmactxrlrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmactxrlrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmactxrlrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmactxrlrFieldTdrlShift = 0
	RegisterDmactxrlrFieldTdrlMask  = 0x3ff
)

// GetTdrl Transmit Descriptor Ring Length
func (r *RegisterDmactxrlrType) GetTdrl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmactxrlrFieldTdrlMask) >> RegisterDmactxrlrFieldTdrlShift)
}

// SetTdrl Transmit Descriptor Ring Length
func (r *RegisterDmactxrlrType) SetTdrl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmactxrlrFieldTdrlMask)|(uint32(value)<<RegisterDmactxrlrFieldTdrlShift))
}

// RegisterDmacrxrlrType Channel Rx descriptor ring length register
type RegisterDmacrxrlrType uint32

func (r *RegisterDmacrxrlrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmacrxrlrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmacrxrlrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmacrxrlrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmacrxrlrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmacrxrlrFieldRdrlShift = 0
	RegisterDmacrxrlrFieldRdrlMask  = 0x3ff
)

// GetRdrl Receive Descriptor Ring Length
func (r *RegisterDmacrxrlrType) GetRdrl() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxrlrFieldRdrlMask) >> RegisterDmacrxrlrFieldRdrlShift)
}

// SetRdrl Receive Descriptor Ring Length
func (r *RegisterDmacrxrlrType) SetRdrl(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxrlrFieldRdrlMask)|(uint32(value)<<RegisterDmacrxrlrFieldRdrlShift))
}

// RegisterDmacierType Channel interrupt enable register
type RegisterDmacierType uint32

func (r *RegisterDmacierType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmacierType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmacierType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmacierType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmacierType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmacierFieldTieShift = 0
	RegisterDmacierFieldTieMask  = 0x1
)

// GetTie Transmit Interrupt Enable
func (r *RegisterDmacierType) GetTie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldTieMask) != 0
}

// SetTie Transmit Interrupt Enable
func (r *RegisterDmacierType) SetTie(value bool) {
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
func (r *RegisterDmacierType) GetTxse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldTxseMask) != 0
}

// SetTxse Transmit Stopped Enable
func (r *RegisterDmacierType) SetTxse(value bool) {
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
func (r *RegisterDmacierType) GetTbue() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldTbueMask) != 0
}

// SetTbue Transmit Buffer Unavailable Enable
func (r *RegisterDmacierType) SetTbue(value bool) {
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
func (r *RegisterDmacierType) GetRie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldRieMask) != 0
}

// SetRie Receive Interrupt Enable
func (r *RegisterDmacierType) SetRie(value bool) {
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
func (r *RegisterDmacierType) GetRbue() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldRbueMask) != 0
}

// SetRbue Receive Buffer Unavailable Enable
func (r *RegisterDmacierType) SetRbue(value bool) {
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
func (r *RegisterDmacierType) GetRse() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldRseMask) != 0
}

// SetRse Receive Stopped Enable
func (r *RegisterDmacierType) SetRse(value bool) {
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
func (r *RegisterDmacierType) GetRwte() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldRwteMask) != 0
}

// SetRwte Receive Watchdog Timeout Enable
func (r *RegisterDmacierType) SetRwte(value bool) {
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
func (r *RegisterDmacierType) GetEtie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldEtieMask) != 0
}

// SetEtie Early Transmit Interrupt Enable
func (r *RegisterDmacierType) SetEtie(value bool) {
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
func (r *RegisterDmacierType) GetErie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldErieMask) != 0
}

// SetErie Early Receive Interrupt Enable
func (r *RegisterDmacierType) SetErie(value bool) {
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
func (r *RegisterDmacierType) GetFbee() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldFbeeMask) != 0
}

// SetFbee Fatal Bus Error Enable
func (r *RegisterDmacierType) SetFbee(value bool) {
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
func (r *RegisterDmacierType) GetCdee() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldCdeeMask) != 0
}

// SetCdee Context Descriptor Error Enable
func (r *RegisterDmacierType) SetCdee(value bool) {
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
func (r *RegisterDmacierType) GetAie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldAieMask) != 0
}

// SetAie Abnormal Interrupt Summary Enable
func (r *RegisterDmacierType) SetAie(value bool) {
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
func (r *RegisterDmacierType) GetNie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacierFieldNieMask) != 0
}

// SetNie Normal Interrupt Summary Enable
func (r *RegisterDmacierType) SetNie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacierFieldNieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacierFieldNieMask)
	}
}

// RegisterDmacrxiwtrType Channel Rx interrupt watchdog timer register
type RegisterDmacrxiwtrType uint32

func (r *RegisterDmacrxiwtrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmacrxiwtrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmacrxiwtrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmacrxiwtrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmacrxiwtrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmacrxiwtrFieldRwtShift = 0
	RegisterDmacrxiwtrFieldRwtMask  = 0xff
)

// GetRwt Receive Interrupt Watchdog Timer Count
func (r *RegisterDmacrxiwtrType) GetRwt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmacrxiwtrFieldRwtMask) >> RegisterDmacrxiwtrFieldRwtShift)
}

// SetRwt Receive Interrupt Watchdog Timer Count
func (r *RegisterDmacrxiwtrType) SetRwt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacrxiwtrFieldRwtMask)|(uint32(value)<<RegisterDmacrxiwtrFieldRwtShift))
}

// RegisterDmaccatxdrType Channel current application transmit descriptor register
type RegisterDmaccatxdrType uint32

func (r *RegisterDmaccatxdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmaccatxdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmaccatxdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmaccatxdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmaccatxdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmaccatxdrFieldCurtdesaptrShift = 0
	RegisterDmaccatxdrFieldCurtdesaptrMask  = 0xffffffff
)

// GetCurtdesaptr Application Transmit Descriptor Address Pointer
func (r *RegisterDmaccatxdrType) GetCurtdesaptr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccatxdrFieldCurtdesaptrMask) >> RegisterDmaccatxdrFieldCurtdesaptrShift)
}

// SetCurtdesaptr Application Transmit Descriptor Address Pointer
func (r *RegisterDmaccatxdrType) SetCurtdesaptr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccatxdrFieldCurtdesaptrMask)|(uint32(value)<<RegisterDmaccatxdrFieldCurtdesaptrShift))
}

// RegisterDmaccarxdrType Channel current application receive descriptor register
type RegisterDmaccarxdrType uint32

func (r *RegisterDmaccarxdrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmaccarxdrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmaccarxdrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmaccarxdrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmaccarxdrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmaccarxdrFieldCurrdesaptrShift = 0
	RegisterDmaccarxdrFieldCurrdesaptrMask  = 0xffffffff
)

// GetCurrdesaptr Application Receive Descriptor Address Pointer
func (r *RegisterDmaccarxdrType) GetCurrdesaptr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccarxdrFieldCurrdesaptrMask) >> RegisterDmaccarxdrFieldCurrdesaptrShift)
}

// SetCurrdesaptr Application Receive Descriptor Address Pointer
func (r *RegisterDmaccarxdrType) SetCurrdesaptr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccarxdrFieldCurrdesaptrMask)|(uint32(value)<<RegisterDmaccarxdrFieldCurrdesaptrShift))
}

// RegisterDmaccatxbrType Channel current application transmit buffer register
type RegisterDmaccatxbrType uint32

func (r *RegisterDmaccatxbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmaccatxbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmaccatxbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmaccatxbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmaccatxbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmaccatxbrFieldCurtbufaptrShift = 0
	RegisterDmaccatxbrFieldCurtbufaptrMask  = 0xffffffff
)

// GetCurtbufaptr Application Transmit Buffer Address Pointer
func (r *RegisterDmaccatxbrType) GetCurtbufaptr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccatxbrFieldCurtbufaptrMask) >> RegisterDmaccatxbrFieldCurtbufaptrShift)
}

// SetCurtbufaptr Application Transmit Buffer Address Pointer
func (r *RegisterDmaccatxbrType) SetCurtbufaptr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccatxbrFieldCurtbufaptrMask)|(uint32(value)<<RegisterDmaccatxbrFieldCurtbufaptrShift))
}

// RegisterDmaccarxbrType Channel current application receive buffer register
type RegisterDmaccarxbrType uint32

func (r *RegisterDmaccarxbrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmaccarxbrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmaccarxbrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmaccarxbrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmaccarxbrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmaccarxbrFieldCurrbufaptrShift = 0
	RegisterDmaccarxbrFieldCurrbufaptrMask  = 0xffffffff
)

// GetCurrbufaptr Application Receive Buffer Address Pointer
func (r *RegisterDmaccarxbrType) GetCurrbufaptr() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDmaccarxbrFieldCurrbufaptrMask) >> RegisterDmaccarxbrFieldCurrbufaptrShift)
}

// SetCurrbufaptr Application Receive Buffer Address Pointer
func (r *RegisterDmaccarxbrType) SetCurrbufaptr(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmaccarxbrFieldCurrbufaptrMask)|(uint32(value)<<RegisterDmaccarxbrFieldCurrbufaptrShift))
}

// RegisterDmacsrType Channel status register
type RegisterDmacsrType uint32

func (r *RegisterDmacsrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmacsrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmacsrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmacsrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmacsrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmacsrFieldTiShift = 0
	RegisterDmacsrFieldTiMask  = 0x1
)

// GetTi Transmit Interrupt
func (r *RegisterDmacsrType) GetTi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldTiMask) != 0
}

// SetTi Transmit Interrupt
func (r *RegisterDmacsrType) SetTi(value bool) {
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
func (r *RegisterDmacsrType) GetTps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldTpsMask) != 0
}

// SetTps Transmit Process Stopped
func (r *RegisterDmacsrType) SetTps(value bool) {
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
func (r *RegisterDmacsrType) GetTbu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldTbuMask) != 0
}

// SetTbu Transmit Buffer Unavailable
func (r *RegisterDmacsrType) SetTbu(value bool) {
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
func (r *RegisterDmacsrType) GetRi() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldRiMask) != 0
}

// SetRi Receive Interrupt
func (r *RegisterDmacsrType) SetRi(value bool) {
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
func (r *RegisterDmacsrType) GetRbu() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldRbuMask) != 0
}

// SetRbu Receive Buffer Unavailable
func (r *RegisterDmacsrType) SetRbu(value bool) {
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
func (r *RegisterDmacsrType) GetRps() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldRpsMask) != 0
}

// SetRps Receive Process Stopped
func (r *RegisterDmacsrType) SetRps(value bool) {
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
func (r *RegisterDmacsrType) GetRwt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldRwtMask) != 0
}

// SetRwt Receive Watchdog Timeout
func (r *RegisterDmacsrType) SetRwt(value bool) {
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
func (r *RegisterDmacsrType) GetEt() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldEtMask) != 0
}

// SetEt Early Transmit Interrupt
func (r *RegisterDmacsrType) SetEt(value bool) {
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
func (r *RegisterDmacsrType) GetEr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldErMask) != 0
}

// SetEr Early Receive Interrupt
func (r *RegisterDmacsrType) SetEr(value bool) {
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
func (r *RegisterDmacsrType) GetFbe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldFbeMask) != 0
}

// SetFbe Fatal Bus Error
func (r *RegisterDmacsrType) SetFbe(value bool) {
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
func (r *RegisterDmacsrType) GetCde() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldCdeMask) != 0
}

// SetCde Context Descriptor Error
func (r *RegisterDmacsrType) SetCde(value bool) {
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
func (r *RegisterDmacsrType) GetAis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldAisMask) != 0
}

// SetAis Abnormal Interrupt Summary
func (r *RegisterDmacsrType) SetAis(value bool) {
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
func (r *RegisterDmacsrType) GetNis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldNisMask) != 0
}

// SetNis Normal Interrupt Summary
func (r *RegisterDmacsrType) SetNis(value bool) {
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
func (r *RegisterDmacsrType) GetTeb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldTebMask) >> RegisterDmacsrFieldTebShift)
}

const (
	RegisterDmacsrFieldRebShift = 19
	RegisterDmacsrFieldRebMask  = 0x380000
)

// GetReb Rx DMA Error Bits
func (r *RegisterDmacsrType) GetReb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDmacsrFieldRebMask) >> RegisterDmacsrFieldRebShift)
}

// RegisterDmacmfcrType Channel missed frame count register
type RegisterDmacmfcrType uint32

func (r *RegisterDmacmfcrType) Load() uint32 {
	return volatile.LoadUint32((*uint32)(r))
}

func (r *RegisterDmacmfcrType) Store(value uint32) {
	volatile.StoreUint32((*uint32)(r), value)
}

func (r *RegisterDmacmfcrType) StoreBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value|mask)
}

func (r *RegisterDmacmfcrType) ClearBits(mask uint32) {
	value := volatile.LoadUint32((*uint32)(r))
	volatile.StoreUint32((*uint32)(r), value&^mask)
}

func (r *RegisterDmacmfcrType) HasBits(mask uint32) bool {
	value := volatile.LoadUint32((*uint32)(r))
	return value&mask != 0
}

const (
	RegisterDmacmfcrFieldMfcShift = 0
	RegisterDmacmfcrFieldMfcMask  = 0x7ff
)

// GetMfc Dropped Packet Counters
func (r *RegisterDmacmfcrType) GetMfc() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDmacmfcrFieldMfcMask) >> RegisterDmacmfcrFieldMfcShift)
}

// SetMfc Dropped Packet Counters
func (r *RegisterDmacmfcrType) SetMfc(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDmacmfcrFieldMfcMask)|(uint32(value)<<RegisterDmacmfcrFieldMfcShift))
}

const (
	RegisterDmacmfcrFieldMfcoShift = 15
	RegisterDmacmfcrFieldMfcoMask  = 0x8000
)

// GetMfco Overflow status of the MFC Counter
func (r *RegisterDmacmfcrType) GetMfco() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDmacmfcrFieldMfcoMask) != 0
}

// SetMfco Overflow status of the MFC Counter
func (r *RegisterDmacmfcrType) SetMfco(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDmacmfcrFieldMfcoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDmacmfcrFieldMfcoMask)
	}
}
