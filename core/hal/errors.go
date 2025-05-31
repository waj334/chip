package hal

const (
	ErrInvalidPinout Error = -1
	ErrInvalidConfig Error = -2
	ErrInvalidBuffer Error = -3
	ErrInvalidState  Error = -4
)

type Error int

func (e Error) Error() string {
	switch e {
	case 0:
		return "no error"
	case -1:
		return "invalid pinout"
	case -2:
		return "invalid configuration"
	case -3:
		return "invalid buffer"
	case -4:
		return "invalid state"
	default:
		return "unknown error"
	}
}
