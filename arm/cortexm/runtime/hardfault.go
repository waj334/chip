package runtime

//go:export _hardfault _hardfault
func hardfault(estack *stackFrame) {
	abort()
}
