package scb_actrl

import (
	"unsafe"
	"volatile"
)

var (
	Scb_actrl = (*_scb_actrl)(unsafe.Pointer(uintptr(0xe000e008)))
)

type _scb_actrl struct {
	Actrl registerActrlType
}

// registerActrlType Auxiliary control register
type registerActrlType uint32

const (
	RegisterActrlFieldDisfoldShift = 2
	RegisterActrlFieldDisfoldMask  = 0x4
)

// GetDisfold DISFOLD
func (r *registerActrlType) GetDisfold() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterActrlFieldDisfoldMask) != 0
}

// SetDisfold DISFOLD
func (r *registerActrlType) SetDisfold(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterActrlFieldDisfoldMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterActrlFieldDisfoldMask)
	}
}

const (
	RegisterActrlFieldFpexcodisShift = 10
	RegisterActrlFieldFpexcodisMask  = 0x400
)

// GetFpexcodis FPEXCODIS
func (r *registerActrlType) GetFpexcodis() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterActrlFieldFpexcodisMask) != 0
}

// SetFpexcodis FPEXCODIS
func (r *registerActrlType) SetFpexcodis(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterActrlFieldFpexcodisMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterActrlFieldFpexcodisMask)
	}
}

const (
	RegisterActrlFieldDisramodeShift = 11
	RegisterActrlFieldDisramodeMask  = 0x800
)

// GetDisramode DISRAMODE
func (r *registerActrlType) GetDisramode() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterActrlFieldDisramodeMask) != 0
}

// SetDisramode DISRAMODE
func (r *registerActrlType) SetDisramode(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterActrlFieldDisramodeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterActrlFieldDisramodeMask)
	}
}

const (
	RegisterActrlFieldDisitmatbflushShift = 12
	RegisterActrlFieldDisitmatbflushMask  = 0x1000
)

// GetDisitmatbflush DISITMATBFLUSH
func (r *registerActrlType) GetDisitmatbflush() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterActrlFieldDisitmatbflushMask) != 0
}

// SetDisitmatbflush DISITMATBFLUSH
func (r *registerActrlType) SetDisitmatbflush(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterActrlFieldDisitmatbflushMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterActrlFieldDisitmatbflushMask)
	}
}
