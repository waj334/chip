package hal

const (
	NoError             Error = "no error"
	ErrInvalidPinout    Error = "invalid pinout"
	ErrInvalidConfig    Error = "invalid configuration"
	ErrInvalidBuffer    Error = "invalid buffer"
	ErrInvalidState     Error = "invalid state"
	ErrInvalidParameter Error = "invalid parameter"
	ErrBusy             Error = "busy"
	ErrNotReady         Error = "no ready"
	ErrTimeout          Error = "timeout"
)

type Error string

func (e Error) Error() string {
	return string(e)
}
