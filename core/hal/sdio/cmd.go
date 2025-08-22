package sdio

type CommandType uint32

const (
	CmdFlagBasic CommandType = 1 << (8 + iota + 1)
	CmdFlagCommAndQueue
	CmdFlagBlockRead
	CmdFlagBlockWrite
	CmdFlagErase
	CmdFlagWriteProtection
	CmdFlagLockCard
	CmdFlagApplicationSpecific
	CmdFlagIOMode
	CmdFlagSwitch
	CmdFlagExtension
)

const (
	CMD0  CommandType = 0 | CmdFlagBasic
	CMD2  CommandType = 2 | CmdFlagBasic
	CMD3  CommandType = 3 | CmdFlagBasic
	CMD4  CommandType = 4 | CmdFlagBasic
	CMD5  CommandType = 5 | CmdFlagIOMode
	CMD6  CommandType = 6 | CmdFlagSwitch
	CMD7  CommandType = 7 | CmdFlagBasic
	CMD8  CommandType = 8 | CmdFlagBasic
	CMD9  CommandType = 9 | CmdFlagBasic
	CMD10 CommandType = 10 | CmdFlagBasic
	CMD11 CommandType = 11 | CmdFlagBasic
	CMD12 CommandType = 12 | CmdFlagBasic
	CMD13 CommandType = 13 | CmdFlagBasic
	CMD15 CommandType = 15 | CmdFlagBasic
	CMD16 CommandType = 16 | CmdFlagLockCard
	CMD17 CommandType = 17 | CmdFlagBlockRead
	CMD18 CommandType = 18 | CmdFlagBlockRead
	CMD19 CommandType = 19 | CmdFlagBlockRead
	CMD20 CommandType = 20
	CMD21 CommandType = 21 | CmdFlagExtension
	CMD22 CommandType = 22
	CMD23 CommandType = 23 | CmdFlagApplicationSpecific
	CMD24 CommandType = 24 | CmdFlagBlockWrite
	CMD25 CommandType = 25 | CmdFlagBlockWrite
	CMD27 CommandType = 27 | CmdFlagBlockWrite
	CMD28 CommandType = 28 | CmdFlagWriteProtection
	CMD29 CommandType = 29 | CmdFlagWriteProtection
	CMD30 CommandType = 30 | CmdFlagWriteProtection
	CMD32 CommandType = 32 | CmdFlagErase
	CMD33 CommandType = 33 | CmdFlagErase
	CMD34 CommandType = 34 | CmdFlagSwitch
	CMD35 CommandType = 35 | CmdFlagSwitch
	CMD36 CommandType = 36 | CmdFlagSwitch
	CMD37 CommandType = 37 | CmdFlagSwitch
	CMD38 CommandType = 38 | CmdFlagErase
	CMD39 CommandType = 39 | CmdFlagExtension
	CMD40 CommandType = 40 | CmdFlagLockCard
	CMD42 CommandType = 42 | CmdFlagLockCard
	CMD43 CommandType = 43 | CmdFlagCommAndQueue
	CMD44 CommandType = 44 | CmdFlagCommAndQueue
	CMD45 CommandType = 45 | CmdFlagCommAndQueue
	CMD46 CommandType = 46 | CmdFlagCommAndQueue
	CMD47 CommandType = 47 | CmdFlagCommAndQueue
	CMD48 CommandType = 48 | CmdFlagExtension
	CMD49 CommandType = 49 | CmdFlagExtension
	CMD50 CommandType = 50 | CmdFlagSwitch
	CMD52 CommandType = 52 | CmdFlagIOMode
	CMD53 CommandType = 53 | CmdFlagIOMode
	CMD55 CommandType = 55 | CmdFlagApplicationSpecific
	CMD56 CommandType = 56 | CmdFlagApplicationSpecific
	CMD57 CommandType = 57 | CmdFlagSwitch
	CMD58 CommandType = 58 | CmdFlagExtension
	CMD59 CommandType = 59 | CmdFlagExtension

	ACMD6  CommandType = 6 | CmdFlagApplicationSpecific
	ACMD13 CommandType = 13 | CmdFlagApplicationSpecific
	ACMD14 CommandType = 14 | CmdFlagApplicationSpecific
	ACMD15 CommandType = 15 | CmdFlagApplicationSpecific
	ACMD16 CommandType = 16 | CmdFlagApplicationSpecific
	ACMD22 CommandType = 22 | CmdFlagApplicationSpecific
	ACMD23 CommandType = 23 | CmdFlagApplicationSpecific
	ACMD28 CommandType = 28 | CmdFlagApplicationSpecific
	ACMD41 CommandType = 41 | CmdFlagApplicationSpecific
	ACMD42 CommandType = 42 | CmdFlagApplicationSpecific
	ACMD51 CommandType = 51 | CmdFlagApplicationSpecific
	ACMD53 CommandType = 53 | CmdFlagApplicationSpecific
	ACMD54 CommandType = 54 | CmdFlagApplicationSpecific
)

func (cmd CommandType) Index() uint32 {
	return uint32(cmd & 0xFF)
}

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
	Data      uint32    // 8 bits (for writes); ignored on reads
	Address   uint32    // 17 bits
	Raw       RAWMode   // 0=Normal, 1=ReadAfterWrite (writes only)
	Function  uint32    // 3 bits (0..7)
	ReadWrite Direction // 0=Read, 1=Write
}

const (
	cmd52DataMask = 0xFF
	cmd52AddrMask = 0x1FFFF
	cmd52FuncMask = 0x7

	cmd52DataShift = 0
	cmd52AddrShift = 9 // leaves bit 8 as required stuff bit (0)
	cmd52RawShift  = 27
	cmd52FuncShift = 28
	cmd52RwShift   = 31
)

func (cmd CMD52Args) Value() uint32 {
	// Enforce spec-ish constraints
	data := cmd.Data & cmd52DataMask
	addr := cmd.Address & cmd52AddrMask
	function := cmd.Function & cmd52FuncMask

	// RAW is only meaningful for writes; clear on reads to be safe.
	raw := uint32(cmd.Raw & 0x1)
	if cmd.ReadWrite&0x1 == 0 {
		raw = 0
	}

	var v uint32
	v |= data << cmd52DataShift
	v |= addr << cmd52AddrShift
	v |= raw << cmd52RawShift
	v |= function << cmd52FuncShift
	v |= uint32(cmd.ReadWrite&0x1) << cmd52RwShift
	return v
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

func DecodeCMD53(args uint32) CMD53Args {
	return CMD53Args{
		Count:     args & 0x1FF,
		Address:   (args >> 9) & 0x1FFFF,
		OpCode:    OpCode(args>>26) & 0x1,
		BlockMode: BlockMode(args>>27) & 0x1,
		Function:  (args >> 28) & 0x7,
		ReadWrite: Direction(args>>31) & 0x1,
	}
}
