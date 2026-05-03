package sdio

import "unsafe"

// STM32H747 memory regions accessible by SDMMC IDMA.
//
// SDMMC1 IDMA uses the AXI bus matrix.
// SDMMC2 IDMA uses the AHB bus matrix.
// Both can reach all regions below. The only memories they
// CANNOT reach are ITCM (0x00000000) and DTCM (0x20000000),
// which are wired exclusively to the Cortex-M7 core.

const (
	// Flash (AXI-mapped) — read-only for DMA, used for firmware blobs in .rodata
	flashAXIBase = uintptr(0x08000000)
	flashAXISize = uintptr(2 * 1024 * 1024) // 2MB (bank1 + bank2)

	// AXI SRAM — D1 domain, primary DMA target
	axiSRAMBase = uintptr(0x24000000)
	axiSRAMSize = uintptr(512 * 1024)

	// SRAM1 — D2 domain (AHB)
	sram1Base = uintptr(0x30000000)
	sram1Size = uintptr(128 * 1024)

	// SRAM2 — D2 domain (AHB)
	sram2Base = uintptr(0x30020000)
	sram2Size = uintptr(128 * 1024)

	// SRAM3 — D2 domain (AHB)
	sram3Base = uintptr(0x30040000)
	sram3Size = uintptr(32 * 1024)

	// SRAM4 — D3 domain (AHB)
	sram4Base = uintptr(0x38000000)
	sram4Size = uintptr(64 * 1024)

	// Backup SRAM — D3 domain
	bkpSRAMBase = uintptr(0x38800000)
	bkpSRAMSize = uintptr(4 * 1024)
)

type memRegion struct {
	base uintptr
	size uintptr
}

var idmaRegions = [...]memRegion{
	{flashAXIBase, flashAXISize},
	{axiSRAMBase, axiSRAMSize},
	{sram1Base, sram1Size},
	{sram2Base, sram2Size},
	{sram3Base, sram3Size},
	{sram4Base, sram4Size},
	{bkpSRAMBase, bkpSRAMSize},
}

func isIDMAAccessible(p uintptr, n int) bool {
	if n <= 0 {
		return false
	}
	end := p + uintptr(n)
	if end < p {
		return false // overflow
	}
	for i := range idmaRegions {
		r := &idmaRegions[i]
		if p >= r.base && end <= r.base+r.size {
			return true
		}
	}
	return false
}

func IsSDMMCDMABuffer(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	p := uintptr(unsafe.Pointer(unsafe.SliceData(data)))
	return isIDMAAccessible(p, len(data))
}
