//go:build !norand

package hal

import "pkg.si-go.dev/chip/arm/cortexm/platform/st/stm32h7x7/cm7/hal/rng"

func init() {
	rng.RNG.Init()
}

//go:export rand runtime.rand
func rand() uint64 {
	return rng.RNG.Uint64()
}
