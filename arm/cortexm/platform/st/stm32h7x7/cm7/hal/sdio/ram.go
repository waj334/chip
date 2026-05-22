package sdio

import "unsafe"

// Memory regions accessible by both SDMMC1 and SDMMC2 IDMA per RM0399 §2.1.7
// "Bus master peripherals".
//
//   - SDMMC1: 32-bit AXI bus matrix master in D1 domain. Reaches AXI SRAM,
//     Flash, FMC, and QUADSPI directly. Reaches D2 SRAMs via the D1-to-D2
//     AHB bridge and D3 SRAMs via D1-to-D3.
//   - SDMMC2: 32-bit AHB bus matrix master in D2 domain. Reaches D2 SRAMs
//     directly. Reaches AXI SRAM, Flash, FMC via D2-to-D1 bridge and D3
//     SRAMs via D2-to-D3 bridge.
//
// Both peripherals can reach all internal memories except ITCM and DTCM,
// which are private to the Cortex-M7 core and have no bus matrix path.
//
// Note: D3 SRAM (SRAM4, BKPSRAM) is power-gated independently. Accessing
// it from a non-D3 master while D3 is in low-power state will fault.
// Callers using D3 buffers must keep the D3 domain awake during transfers.

const (
	// Flash (AXI-mapped) — read-only for IDMA, used for firmware blobs in .rodata
	flashAXIBase = uintptr(0x08000000)
	flashAXISize = uintptr(2 * 1024 * 1024)

	// AXI SRAM — D1 domain, primary DMA target for SDMMC1
	axiSRAMBase = uintptr(0x24000000)
	axiSRAMSize = uintptr(512 * 1024)

	// SRAM1 — D2 domain
	sram1Base = uintptr(0x30000000)
	sram1Size = uintptr(128 * 1024)

	// SRAM2 — D2 domain
	sram2Base = uintptr(0x30020000)
	sram2Size = uintptr(128 * 1024)

	// SRAM3 — D2 domain
	sram3Base = uintptr(0x30040000)
	sram3Size = uintptr(32 * 1024)

	// SRAM4 — D3 domain (reachable via inter-domain bridge; requires D3 awake)
	sram4Base = uintptr(0x38000000)
	sram4Size = uintptr(64 * 1024)

	// Backup SRAM — D3 domain (reachable via inter-domain bridge; requires D3 awake)
	bkpSRAMBase = uintptr(0x38800000)
	bkpSRAMSize = uintptr(4 * 1024)

	// FMC banks — external memories via FMC controller in D1 domain
	fmcNORBase    = uintptr(0x60000000)
	fmcNORSize    = uintptr(256 * 1024 * 1024)
	fmcNANDBase   = uintptr(0x80000000)
	fmcNANDSize   = uintptr(256 * 1024 * 1024)
	fmcSDRAM1Base = uintptr(0xC0000000)
	fmcSDRAM1Size = uintptr(256 * 1024 * 1024)
	fmcSDRAM2Base = uintptr(0xD0000000)
	fmcSDRAM2Size = uintptr(256 * 1024 * 1024)

	// Quad-SPI memory-mapped region — read-only for IDMA, often used for
	// firmware blob storage. Accessible from both SDMMC peripherals via
	// the QSPI controller in D1.
	qspiBase = uintptr(0x90000000)
	qspiSize = uintptr(256 * 1024 * 1024)
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
	{fmcNORBase, fmcNORSize},
	{fmcNANDBase, fmcNANDSize},
	{fmcSDRAM1Base, fmcSDRAM1Size},
	{fmcSDRAM2Base, fmcSDRAM2Size},
	{qspiBase, qspiSize},
}

// Memory regions reachable by SDMMC1 IDMA per RM0399 §2.1.7.
//
// SDMMC1 is a D1-domain master on the AXI bus matrix. It reaches D1 memories
// natively, D2 memories via the D1-to-D2 AHB bridge, and D3 memories via
// the D1-to-D3 AHB bridge.
//
// D3 entries (SRAM4, BKPSRAM) are included for spec completeness but require
// the D3 domain to be powered. Applications that power-gate D3 should not
// place SDMMC buffers there.
var sdmmc1Regions = [...]memRegion{
	{flashAXIBase, flashAXISize},
	{axiSRAMBase, axiSRAMSize},
	{sram1Base, sram1Size},
	{sram2Base, sram2Size},
	{sram3Base, sram3Size},
	{sram4Base, sram4Size},
	{bkpSRAMBase, bkpSRAMSize},
	{fmcNORBase, fmcNORSize},
	{fmcNANDBase, fmcNANDSize},
	{fmcSDRAM1Base, fmcSDRAM1Size},
	{fmcSDRAM2Base, fmcSDRAM2Size},
	{qspiBase, qspiSize},
}

// Memory regions reachable by SDMMC2 IDMA per RM0399 §2.1.7.
//
// SDMMC2 is a D2-domain master on the AHB bus matrix. It reaches D2 memories
// natively, D1 memories via the D2-to-D1 AHB bridge, and D3 memories via
// the D2-to-D3 AHB bridge.
//
// As with SDMMC1, D3 entries require the D3 domain to be powered.
var sdmmc2Regions = [...]memRegion{
	{flashAXIBase, flashAXISize},
	{axiSRAMBase, axiSRAMSize},
	{sram1Base, sram1Size},
	{sram2Base, sram2Size},
	{sram3Base, sram3Size},
	{sram4Base, sram4Size},
	{bkpSRAMBase, bkpSRAMSize},
	{fmcNORBase, fmcNORSize},
	{fmcNANDBase, fmcNANDSize},
	{fmcSDRAM1Base, fmcSDRAM1Size},
	{fmcSDRAM2Base, fmcSDRAM2Size},
	{qspiBase, qspiSize},
}

func regionsFor(instance SDIO) []memRegion {
	switch instance {
	case SDIO1:
		return sdmmc1Regions[:]
	case SDIO2:
		return sdmmc2Regions[:]
	default:
		return nil
	}
}

func isIDMAAccessible(p uintptr, n int, instance SDIO) bool {
	if n <= 0 {
		return false
	}
	end := p + uintptr(n)
	if end < p {
		return false // overflow
	}
	for _, r := range regionsFor(instance) {
		if p >= r.base && end <= r.base+r.size {
			return true
		}
	}
	return false
}

func IsSDMMCDMABuffer(data []byte, instance SDIO) bool {
	if len(data) == 0 {
		return false
	}
	p := uintptr(unsafe.Pointer(unsafe.SliceData(data)))
	return isIDMAAccessible(p, len(data), instance)
}
