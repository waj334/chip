package runtime

import (
	"unsafe"
)

//sigo:extern main main.main
func main()

//sigo:extern initgc runtime.initgc
func initgc()

//sigo:extern gcmain runtime.gcmain
func gcmain()

//sigo:extern __libc_init_array __libc_init_array
func __libc_init_array()

//sigo:extern __start_bss __start_bss
var __start_bss unsafe.Pointer

//sigo:extern __end_bss __end_bss
var __end_bss unsafe.Pointer

//sigo:extern __start_data __start_data
var __start_data unsafe.Pointer

//sigo:extern __end_data __end_data
var __end_data unsafe.Pointer

//sigo:extern __data_base_addr __data_base_addr
var __data_base_addr unsafe.Pointer

func initMemory() {
	// Zero init globals
	sbss := unsafe.Pointer(&__start_bss)
	ebss := unsafe.Pointer(&__end_bss)
	for ptr := sbss; ptr != ebss; ptr = unsafe.Add(ptr, 4) {
		*(*uint32)(ptr) = 0
	}

	// Initialize data from flash
	dst := unsafe.Pointer(&__start_data)
	src := unsafe.Pointer(&__data_base_addr)
	edata := unsafe.Pointer(&__end_data)
	for dst != edata {
		*(*uint32)(dst) = *(*uint32)(src)
		dst = unsafe.Add(dst, 4)
		src = unsafe.Add(src, 4)
	}
}

//sigo:interrupt _entry ResetHandler
//sigo:required _entry
func _entry() {
	state := DisableInterrupts()

	// Initialize the global variables.
	initMemory()

	// Initialize the FPU if it was enabled during the build.
	initFPU()

	// Init the GC
	initgc()

	// Initialize all the global variables.
	__libc_init_array()

	go func() {
		// Start the garbage collector.
		go gcmain()

		// Run the main program.
		main()

		// Stop if main returns.
		abort()
	}()

	// Start the SysTick counter.
	initSysTick()

	// Enable interrupts.
	EnableInterrupts(state)

	// Schedule goroutines.
	gosched()
}
