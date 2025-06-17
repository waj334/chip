package sdio

type BusWidth uint8

const (
	Width1Bit BusWidth = iota + 1
	Width4Bit
	Width8Bit
)

type BusSpeed uint8

const (
	Sdr12 BusSpeed = iota + 1
	Sdr25
	Sdr50
	Sdr104
	Ddr50
	DDR200
	Ds
	Hs
)

type CommandIndex uint8

const (
	CMD0  CommandIndex = 0
	CMD2  CommandIndex = 2
	CMD3  CommandIndex = 3
	CMD5  CommandIndex = 5
	CMD7  CommandIndex = 7
	CMD8  CommandIndex = 8
	CMD9  CommandIndex = 9
	CMD12 CommandIndex = 12
	CMD17 CommandIndex = 17
	CMD24 CommandIndex = 24
	CMD41 CommandIndex = 41
	CMD52 CommandIndex = 52
	CMD53 CommandIndex = 53
	CMD55 CommandIndex = 55
)

type _error uint8

const (
	ErrNone _error = iota
	ErrCommandFail
	ErrCommandCrcFail
	ErrTimeout
	ErrNotReady
)

func (e _error) Error() string {
	switch e {
	case ErrNone:
		return "no error"
	case ErrCommandFail:
		return "command fail"
	case ErrCommandCrcFail:
		return "command crc fail"
	case ErrTimeout:
		return "command timeout"
	case ErrNotReady:
		return "the card is not ready"
	default:
		return "unknown error"
	}
}

type Response [4]uint32

type Command struct {
	Argument uint32
	Index    CommandIndex
}

type Host interface {
	SetBusSpeed(speed BusSpeed) error
	SetBusWidth(width BusWidth) error
	SetClockFrequency(hz uint32) error

	SendCommand(cmd Command) (Response, error)

	ReadBlock(buf []byte, blockAddr uint32) error
	WriteBlock(buf []byte, blockAddr uint32) error

	ReadBlocks(buf []byte, blockAddr, count uint32) error
	WriteBlocks(buf []byte, blockAddr, count uint32) error
}
