package sdio

type Direction uint32

const (
	Read  Direction = 0
	Write Direction = 1
)

type OpCode uint32

const (
	Fixed        OpCode = 0
	Incrementing OpCode = 1
)

type BlockMode uint32

const (
	Bytes  BlockMode = 0
	Blocks BlockMode = 1
)

type RAWMode uint32

const (
	Normal         RAWMode = 0
	ReadAfterWrite RAWMode = 1
)

type CMD5Args struct {
	VoltageBits uint32
	SwitchTo1v8 uint32
}

func (cmd CMD5Args) Value() uint32 {
	var value uint32
	value |= cmd.VoltageBits & 0xFFFFFF
	value |= (cmd.SwitchTo1v8 & 0x1) << 24
	return value
}

type CMD7Args struct {
	RCA uint16
}

func (cmd CMD7Args) Value() uint32 {
	var value uint32
	value |= uint32(cmd.RCA) << 16
	return value
}

type CMD52Args struct {
	Data      uint32
	Address   uint32
	Raw       RAWMode
	Function  uint32
	ReadWrite Direction
}

func (cmd CMD52Args) Value() uint32 {
	var value uint32
	value |= cmd.Data & 0xFF
	value |= (cmd.Address & 0x1FFFF) << 9
	value |= uint32(cmd.Raw&0x1) << 27
	value |= (cmd.Function & 0x7) << 28
	value |= uint32(cmd.ReadWrite&0x1) << 31
	return value
}

type CMD53Args struct {
	Count     uint32
	Address   uint32
	OpCode    OpCode
	BlockMode BlockMode
	Function  uint32
	ReadWrite Direction
}

func (cmd CMD53Args) Value() uint32 {
	var value uint32
	value |= cmd.Count & 0x1FF
	value |= (cmd.Address & 0x1FFFF) << 9
	value |= uint32(cmd.OpCode&0x1) << 26
	value |= uint32(cmd.BlockMode&0x1) << 27
	value |= (cmd.Function & 0x7) << 28
	value |= uint32(cmd.ReadWrite&0x1) << 31
	return value
}
