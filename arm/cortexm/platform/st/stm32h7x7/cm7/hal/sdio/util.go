package sdio

import "unsafe"

func dataPointer(data []byte) unsafe.Pointer {
	if len(data) == 0 {
		return nil
	}
	return unsafe.Pointer(unsafe.SliceData(data))
}

func isWordAligned(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	return uintptr(dataPointer(data))&3 == 0
}

func loadAlignedWord(data []byte, off int) uint32 {
	return *(*uint32)(unsafe.Add(dataPointer(data), off))
}

func storeAlignedWord(data []byte, off int, v uint32) {
	*(*uint32)(unsafe.Add(dataPointer(data), off)) = v
}

func packWordLE(data []byte, off int) uint32 {
	var v uint32
	if off < len(data) {
		v |= uint32(data[off])
	}
	if off+1 < len(data) {
		v |= uint32(data[off+1]) << 8
	}
	if off+2 < len(data) {
		v |= uint32(data[off+2]) << 16
	}
	if off+3 < len(data) {
		v |= uint32(data[off+3]) << 24
	}
	return v
}

func unpackWordLE(data []byte, off int, v uint32) {
	if off < len(data) {
		data[off] = byte(v)
	}
	if off+1 < len(data) {
		data[off+1] = byte(v >> 8)
	}
	if off+2 < len(data) {
		data[off+2] = byte(v >> 16)
	}
	if off+3 < len(data) {
		data[off+3] = byte(v >> 24)
	}
}

func boolToInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func blockSize(size int) uint8 {
	// The SDMMC block size register expects log2(size).
	// Only power-of-two sizes are valid; callers must validate beforehand.
	switch size {
	case 1:
		return 0b0000
	case 2:
		return 0b0001
	case 4:
		return 0b0010
	case 8:
		return 0b0011
	case 16:
		return 0b0100
	case 32:
		return 0b0101
	case 64:
		return 0b0110
	case 128:
		return 0b0111
	case 256:
		return 0b1000
	case 512:
		return 0b1001
	case 1024:
		return 0b1010
	case 2048:
		return 0b1011
	case 4096:
		return 0b1100
	case 8192:
		return 0b1101
	case 16384:
		return 0b1110
	default:
		panic("invalid block size: must be a power of two")
	}
}
