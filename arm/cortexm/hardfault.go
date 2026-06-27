package cortexm

import "fmt"

//go:export hardfault runtime.hardfault
func hardfault(estack *exceptionFrame) {
	fmt.Printf("lastExcReturn=%#x\n", lastExcReturn)
	fmt.Printf("lastRestorePSP=%#x\n", lastRestorePSP)
	fmt.Println("---- STACK FRAME ----")
	estack.Dump()
	fmt.Println("---------------------")
	abort()
}
