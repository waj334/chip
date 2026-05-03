package rng

type RNG interface {
	Uint64() uint64
}
