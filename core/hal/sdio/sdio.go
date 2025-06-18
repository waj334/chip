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

type SDType uint8

const (
	SDMMC SDType = iota
	SDIO
)

type _error uint8

const (
	ErrNone _error = iota
	ErrCommandFail
	ErrCommandCrcFail
	ErrTimeout
	ErrDataError
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
	case ErrDataError:
		return "command data error"
	case ErrNotReady:
		return "the card is not ready"
	default:
		return "unknown error"
	}
}

type Response [4]uint32

type Command struct {
	Data      []byte
	BlockSize uint32
	Argument  uint32
	Class     CommandType
}

type Transfer struct {
	Address    uint32
	BlockSize  uint16
	BlockCount uint16
	Function   uint8
	Increment  bool
}

type Host interface {
	SetBusSpeed(speed BusSpeed) error
	SetBusWidth(width BusWidth) error
	SetClockFrequency(hz uint32) error
	SendCommand(cmd Command) (Response, error)
}
