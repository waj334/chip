package ethernet_mtl

import (
	"unsafe"
	"volatile"
)

var (
	Ethernet_mtl = (*_ethernet_mtl)(unsafe.Pointer(uintptr(0x40028c00)))
)

type _ethernet_mtl struct {
	Mtlomr      registerMtlomrType
	_           [28]byte
	Mtlisr      registerMtlisrType
	_           [220]byte
	Mtltxqomr   registerMtltxqomrType
	Mtltxqur    registerMtltxqurType
	Mtltxqdr    registerMtltxqdrType
	_           [32]byte
	Mtlqicsr    registerMtlqicsrType
	Mtlrxqomr   registerMtlrxqomrType
	Mtlrxqmpocr registerMtlrxqmpocrType
	Mtlrxqdr    registerMtlrxqdrType
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
	RegisterMtltxqomrFieldTqsMask  = 0x70000
)

// GetTqs Transmit Queue Size
func (r *registerMtltxqomrType) GetTqs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterMtltxqomrFieldTqsMask) >> RegisterMtltxqomrFieldTqsShift)
}

// SetTqs Transmit Queue Size
func (r *registerMtltxqomrType) SetTqs(value uint8) {
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
	RegisterMtlrxqomrFieldDis_tcp_efShift = 6
	RegisterMtlrxqomrFieldDis_tcp_efMask  = 0x40
)

// GetDis_tcp_ef DIS_TCP_EF
func (r *registerMtlrxqomrType) GetDis_tcp_ef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterMtlrxqomrFieldDis_tcp_efMask) != 0
}

// SetDis_tcp_ef DIS_TCP_EF
func (r *registerMtlrxqomrType) SetDis_tcp_ef(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterMtlrxqomrFieldDis_tcp_efMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterMtlrxqomrFieldDis_tcp_efMask)
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
