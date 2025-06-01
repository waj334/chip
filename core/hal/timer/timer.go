package timer

import "time"

type Timer interface {
	Reset(tick uint64)
	Stop()

	Tick() uint64
	Now() time.Time
	Resolution() time.Duration
}
