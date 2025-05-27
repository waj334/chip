package runtime

import (
	"volatile"

	"pkg.si-go.dev/chip/arm/cortexm/reg/nvic"
)

var IrqPriorityMask uint8 = 0xFF

type Interrupt int16

func (i Interrupt) EnableIRQ() {
	nvic.Nvic.Iser[i>>5].SetSetena(1 << (i & 0x1F))
}

func (i Interrupt) DisableIRQ() {
	nvic.Nvic.Icer[i>>5].SetClrena(1 << (i & 0x1F))
}

func (i Interrupt) SetPriority(priority uint8) {
	const IrqPriorityMask = 0xFF // or mask to implemented bits if needed
	index := int(i) / 4
	offset := int(i) % 4

	// Load the full 32-bit register
	reg := volatile.LoadUint32((*uint32)(&nvic.Nvic.Ipr[index]))

	// Clear the target byte
	shift := offset * 8
	mask := uint32(0xFF) << shift
	reg &^= mask

	// Set the priority byte
	reg |= uint32(priority&IrqPriorityMask) << shift

	// Store it back
	volatile.StoreUint32((*uint32)(&nvic.Nvic.Ipr[index]), reg)
}
