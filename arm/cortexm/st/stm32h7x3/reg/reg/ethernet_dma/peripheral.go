package ethernet_dma

import (
	"unsafe"
	"volatile"
)

var (
	Ethernet_dma = (*_ethernet_dma)(unsafe.Pointer(uintptr(0x40029000)))
)

type _ethernet_dma struct {
	Dmamr      registerDmamrType
	Dmasbmr    registerDmasbmrType
	Dmaisr     registerDmaisrType
	Dmadsr     registerDmadsrType
	_          [240]byte
	Dmaccr     registerDmaccrType
	Dmactxcr   registerDmactxcrType
	Dmacrxcr   registerDmacrxcrType
	_          [8]byte
	Dmactxdlar registerDmactxdlarType
	_          [4]byte
	Dmacrxdlar registerDmacrxdlarType
	Dmactxdtpr registerDmactxdtprType
	_          [4]byte
	Dmacrxdtpr registerDmacrxdtprType
	Dmactxrlr  registerDmactxrlrType
	Dmacrxrlr  registerDmacrxrlrType
	Dmacier    registerDmacierType
	Dmacrxiwtr registerDmacrxiwtrType
	_          [8]byte
	Dmaccatxdr registerDmaccatxdrType
	_          [4]byte
	Dmaccarxdr registerDmaccarxdrType
	_          [4]byte
	Dmaccatxbr registerDmaccatxbrType
	_          [4]byte
	Dmaccarxbr registerDmaccarxbrType
	Dmacsr     registerDmacsrType
	_          [8]byte
	Dmacmfcr   registerDmacmfcrType
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
