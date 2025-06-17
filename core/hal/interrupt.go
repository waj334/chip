package hal

type Interrupt interface {
	Enable()
	Disable()
	SetPriority(priority uint8)
}
