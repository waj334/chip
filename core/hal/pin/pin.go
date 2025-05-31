package pin

type base interface {
	High()
	Low()
	Toggle()

	Set(on bool)
	Get() bool

	SetValue(value uint) error
	Value() (uint, error)

	SetMode(dir Mode)
	GetMode() Mode

	SetOutputMode(output OutputMode)
	GetOutputMode() OutputMode

	SetSpeedMode(speed SpeedMode)
	GetSpeedMode() SpeedMode

	SetPullMode(mode PullMode)
	GetPullMode() PullMode
}

type Pin interface {
	base

	SetInterrupt(mode IRQMode, fn InterruptHandler[base])
	ClearInterrupt()
}
