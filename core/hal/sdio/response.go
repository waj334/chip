package sdio

/* Masks for errors in an R5 response */
const (
	R5ComCrcError    = 0x80
	R5IllegalCommand = 0x40
	R5IoCurrentState = 0x30
	R5Error          = 0x08
	R5FunctionNumber = 0x02
	R5OutOfRange     = 0x01
	R5ErrorBits      = R5ComCrcError | R5IllegalCommand | R5Error | R5FunctionNumber | R5OutOfRange
)

type R4 uint32

func (r R4) Ocr() uint32 {
	return uint32(r) & 0xFFFFFF
}

func (r R4) SwitchTo1v8() bool {
	return (r>>24)&0x1 != 0
}

func (r R4) MemoryPresent() bool {
	return (r>>26)&0x1 != 0
}

func (r R4) NumberOfFunctions() int {
	return int(r>>27) & 0x7
}

func (r R4) Ready() bool {
	return int(r>>31)&0x1 != 0
}

type R5 uint32

func (r R5) Data() uint8 {
	return uint8(r)
}

func (r R5) Flags() uint8 {
	return uint8(r >> 8)
}

type R6 uint32

func (r R6) Status() uint16 {
	return uint16(r)
}

func (r R6) RCA() uint16 {
	return uint16(r >> 16)
}
