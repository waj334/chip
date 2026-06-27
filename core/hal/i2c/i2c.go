package i2c

type I2C interface {
	SetClockFrequency(clockSpeedHz uint32) bool
	GetClockFrequency() uint32
	WriteAddress(addr uint16, b []byte) (n int, err error)
	ReadAddress(addr uint16, b []byte) (n int, err error)
}
