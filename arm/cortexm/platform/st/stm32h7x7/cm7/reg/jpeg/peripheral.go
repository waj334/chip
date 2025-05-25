//go:build arm && cortexm && armv7m && armv7em && thumb && (stm32h747ag || stm32h747ig || stm32h747bg || stm32h747xg || stm32h747zi || stm32h747ai || stm32h747ii || stm32h747bi || stm32h747xi || stm32h757zi || stm32h757ai || stm32h757ii || stm32h757bi || stm32h757xi)

package jpeg

import (
	"unsafe"
	"volatile"
)

var (
	Jpeg = (*_jpeg)(unsafe.Pointer(uintptr(0x52003000)))
)

type _jpeg struct {
	Confr0  registerConfr0Type
	Confr1  registerConfr1Type
	Confr2  registerConfr2Type
	Confr3  registerConfr3Type
	Confrn1 registerConfrn1Type
	Confrn2 registerConfrn2Type
	Confrn3 registerConfrn3Type
	Confrn4 registerConfrn4Type
	_       [16]byte
	Cr      registerCrType
	Sr      registerSrType
	Cfr     registerCfrType
	_       [4]byte
	Dir     registerDirType
	Dor     registerDorType
}

// registerConfr0Type JPEG codec control register
type registerConfr0Type uint32

const (
	RegisterConfr0FieldStartShift = 0
	RegisterConfr0FieldStartMask  = 0x1
)

// GetStart Start This bit start or stop the encoding or decoding process. Read this register always return 0.
func (r *registerConfr0Type) GetStart() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterConfr0FieldStartMask) != 0
}

// SetStart Start This bit start or stop the encoding or decoding process. Read this register always return 0.
func (r *registerConfr0Type) SetStart(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterConfr0FieldStartMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterConfr0FieldStartMask)
	}
}

// registerConfr1Type JPEG codec configuration register 1
type registerConfr1Type uint32

const (
	RegisterConfr1FieldNfShift = 0
	RegisterConfr1FieldNfMask  = 0x3
)

// GetNf Number of color components This field defines the number of color components minus 1.
func (r *registerConfr1Type) GetNf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfr1FieldNfMask) >> RegisterConfr1FieldNfShift)
}

// SetNf Number of color components This field defines the number of color components minus 1.
func (r *registerConfr1Type) SetNf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfr1FieldNfMask)|(uint32(value)<<RegisterConfr1FieldNfShift))
}

const (
	RegisterConfr1FieldDeShift = 3
	RegisterConfr1FieldDeMask  = 0x8
)

// GetDe Decoding Enable This bit selects the coding or decoding process
func (r *registerConfr1Type) GetDe() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterConfr1FieldDeMask) != 0
}

// SetDe Decoding Enable This bit selects the coding or decoding process
func (r *registerConfr1Type) SetDe(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterConfr1FieldDeMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterConfr1FieldDeMask)
	}
}

const (
	RegisterConfr1FieldColorspaceShift = 4
	RegisterConfr1FieldColorspaceMask  = 0x30
)

// GetColorspace Color Space This filed defines the number of quantization tables minus 1 to insert in the output stream.
func (r *registerConfr1Type) GetColorspace() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfr1FieldColorspaceMask) >> RegisterConfr1FieldColorspaceShift)
}

// SetColorspace Color Space This filed defines the number of quantization tables minus 1 to insert in the output stream.
func (r *registerConfr1Type) SetColorspace(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfr1FieldColorspaceMask)|(uint32(value)<<RegisterConfr1FieldColorspaceShift))
}

const (
	RegisterConfr1FieldNsShift = 6
	RegisterConfr1FieldNsMask  = 0xc0
)

// GetNs Number of components for Scan This field defines the number of components minus 1 for scan header marker segment.
func (r *registerConfr1Type) GetNs() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfr1FieldNsMask) >> RegisterConfr1FieldNsShift)
}

// SetNs Number of components for Scan This field defines the number of components minus 1 for scan header marker segment.
func (r *registerConfr1Type) SetNs(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfr1FieldNsMask)|(uint32(value)<<RegisterConfr1FieldNsShift))
}

const (
	RegisterConfr1FieldHdrShift = 8
	RegisterConfr1FieldHdrMask  = 0x100
)

// GetHdr Header Processing This bit enable the header processing (generation/parsing).
func (r *registerConfr1Type) GetHdr() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterConfr1FieldHdrMask) != 0
}

// SetHdr Header Processing This bit enable the header processing (generation/parsing).
func (r *registerConfr1Type) SetHdr(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterConfr1FieldHdrMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterConfr1FieldHdrMask)
	}
}

const (
	RegisterConfr1FieldYsizeShift = 16
	RegisterConfr1FieldYsizeMask  = 0xffff0000
)

// GetYsize Y Size This field defines the number of lines in source image.
func (r *registerConfr1Type) GetYsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterConfr1FieldYsizeMask) >> RegisterConfr1FieldYsizeShift)
}

// SetYsize Y Size This field defines the number of lines in source image.
func (r *registerConfr1Type) SetYsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfr1FieldYsizeMask)|(uint32(value)<<RegisterConfr1FieldYsizeShift))
}

// registerConfr2Type JPEG codec configuration register 2
type registerConfr2Type uint32

const (
	RegisterConfr2FieldNmcuShift = 0
	RegisterConfr2FieldNmcuMask  = 0x3ffffff
)

// GetNmcu Number of MCU For encoding: this field defines the number of MCU units minus 1 to encode. For decoding: this field indicates the number of complete MCU units minus 1 to be decoded (this field is updated after the JPEG header parsing). If the decoded image size has not a X or Y size multiple of 8 or 16 (depending on the sub-sampling process), the resulting incomplete or empty MCU must be added to this value to get the total number of MCU generated.
func (r *registerConfr2Type) GetNmcu() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterConfr2FieldNmcuMask) >> RegisterConfr2FieldNmcuShift)
}

// SetNmcu Number of MCU For encoding: this field defines the number of MCU units minus 1 to encode. For decoding: this field indicates the number of complete MCU units minus 1 to be decoded (this field is updated after the JPEG header parsing). If the decoded image size has not a X or Y size multiple of 8 or 16 (depending on the sub-sampling process), the resulting incomplete or empty MCU must be added to this value to get the total number of MCU generated.
func (r *registerConfr2Type) SetNmcu(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfr2FieldNmcuMask)|(uint32(value)<<RegisterConfr2FieldNmcuShift))
}

// registerConfr3Type JPEG codec configuration register 3
type registerConfr3Type uint32

const (
	RegisterConfr3FieldXsizeShift = 16
	RegisterConfr3FieldXsizeMask  = 0xffff0000
)

// GetXsize X size This field defines the number of pixels per line.
func (r *registerConfr3Type) GetXsize() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterConfr3FieldXsizeMask) >> RegisterConfr3FieldXsizeShift)
}

// SetXsize X size This field defines the number of pixels per line.
func (r *registerConfr3Type) SetXsize(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfr3FieldXsizeMask)|(uint32(value)<<RegisterConfr3FieldXsizeShift))
}

// registerConfrn1Type JPEG codec configuration register 4-7
type registerConfrn1Type uint32

const (
	RegisterConfrn1FieldHdShift = 0
	RegisterConfrn1FieldHdMask  = 0x1
)

// GetHd Huffman DC Selects the Huffman table for encoding the DC coefficients.
func (r *registerConfrn1Type) GetHd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterConfrn1FieldHdMask) != 0
}

// SetHd Huffman DC Selects the Huffman table for encoding the DC coefficients.
func (r *registerConfrn1Type) SetHd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterConfrn1FieldHdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterConfrn1FieldHdMask)
	}
}

const (
	RegisterConfrn1FieldHaShift = 1
	RegisterConfrn1FieldHaMask  = 0x2
)

// GetHa Huffman AC Selects the Huffman table for encoding the AC coefficients.
func (r *registerConfrn1Type) GetHa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterConfrn1FieldHaMask) != 0
}

// SetHa Huffman AC Selects the Huffman table for encoding the AC coefficients.
func (r *registerConfrn1Type) SetHa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterConfrn1FieldHaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterConfrn1FieldHaMask)
	}
}

const (
	RegisterConfrn1FieldQtShift = 2
	RegisterConfrn1FieldQtMask  = 0xc
)

// GetQt Quantization Table Selects quantization table associated with a color component.
func (r *registerConfrn1Type) GetQt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn1FieldQtMask) >> RegisterConfrn1FieldQtShift)
}

// SetQt Quantization Table Selects quantization table associated with a color component.
func (r *registerConfrn1Type) SetQt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn1FieldQtMask)|(uint32(value)<<RegisterConfrn1FieldQtShift))
}

const (
	RegisterConfrn1FieldNbShift = 4
	RegisterConfrn1FieldNbMask  = 0xf0
)

// GetNb Number of Block Number of data units minus 1 that belong to a particular color in the MCU.
func (r *registerConfrn1Type) GetNb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn1FieldNbMask) >> RegisterConfrn1FieldNbShift)
}

// SetNb Number of Block Number of data units minus 1 that belong to a particular color in the MCU.
func (r *registerConfrn1Type) SetNb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn1FieldNbMask)|(uint32(value)<<RegisterConfrn1FieldNbShift))
}

const (
	RegisterConfrn1FieldVsfShift = 8
	RegisterConfrn1FieldVsfMask  = 0xf00
)

// GetVsf Vertical Sampling Factor Vertical sampling factor for component i.
func (r *registerConfrn1Type) GetVsf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn1FieldVsfMask) >> RegisterConfrn1FieldVsfShift)
}

// SetVsf Vertical Sampling Factor Vertical sampling factor for component i.
func (r *registerConfrn1Type) SetVsf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn1FieldVsfMask)|(uint32(value)<<RegisterConfrn1FieldVsfShift))
}

const (
	RegisterConfrn1FieldHsfShift = 12
	RegisterConfrn1FieldHsfMask  = 0xf000
)

// GetHsf Horizontal Sampling Factor Horizontal sampling factor for component i.
func (r *registerConfrn1Type) GetHsf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn1FieldHsfMask) >> RegisterConfrn1FieldHsfShift)
}

// SetHsf Horizontal Sampling Factor Horizontal sampling factor for component i.
func (r *registerConfrn1Type) SetHsf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn1FieldHsfMask)|(uint32(value)<<RegisterConfrn1FieldHsfShift))
}

// registerConfrn2Type JPEG codec configuration register 4-7
type registerConfrn2Type uint32

const (
	RegisterConfrn2FieldHdShift = 0
	RegisterConfrn2FieldHdMask  = 0x1
)

// GetHd Huffman DC Selects the Huffman table for encoding the DC coefficients.
func (r *registerConfrn2Type) GetHd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterConfrn2FieldHdMask) != 0
}

// SetHd Huffman DC Selects the Huffman table for encoding the DC coefficients.
func (r *registerConfrn2Type) SetHd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterConfrn2FieldHdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterConfrn2FieldHdMask)
	}
}

const (
	RegisterConfrn2FieldHaShift = 1
	RegisterConfrn2FieldHaMask  = 0x2
)

// GetHa Huffman AC Selects the Huffman table for encoding the AC coefficients.
func (r *registerConfrn2Type) GetHa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterConfrn2FieldHaMask) != 0
}

// SetHa Huffman AC Selects the Huffman table for encoding the AC coefficients.
func (r *registerConfrn2Type) SetHa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterConfrn2FieldHaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterConfrn2FieldHaMask)
	}
}

const (
	RegisterConfrn2FieldQtShift = 2
	RegisterConfrn2FieldQtMask  = 0xc
)

// GetQt Quantization Table Selects quantization table associated with a color component.
func (r *registerConfrn2Type) GetQt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn2FieldQtMask) >> RegisterConfrn2FieldQtShift)
}

// SetQt Quantization Table Selects quantization table associated with a color component.
func (r *registerConfrn2Type) SetQt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn2FieldQtMask)|(uint32(value)<<RegisterConfrn2FieldQtShift))
}

const (
	RegisterConfrn2FieldNbShift = 4
	RegisterConfrn2FieldNbMask  = 0xf0
)

// GetNb Number of Block Number of data units minus 1 that belong to a particular color in the MCU.
func (r *registerConfrn2Type) GetNb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn2FieldNbMask) >> RegisterConfrn2FieldNbShift)
}

// SetNb Number of Block Number of data units minus 1 that belong to a particular color in the MCU.
func (r *registerConfrn2Type) SetNb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn2FieldNbMask)|(uint32(value)<<RegisterConfrn2FieldNbShift))
}

const (
	RegisterConfrn2FieldVsfShift = 8
	RegisterConfrn2FieldVsfMask  = 0xf00
)

// GetVsf Vertical Sampling Factor Vertical sampling factor for component i.
func (r *registerConfrn2Type) GetVsf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn2FieldVsfMask) >> RegisterConfrn2FieldVsfShift)
}

// SetVsf Vertical Sampling Factor Vertical sampling factor for component i.
func (r *registerConfrn2Type) SetVsf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn2FieldVsfMask)|(uint32(value)<<RegisterConfrn2FieldVsfShift))
}

const (
	RegisterConfrn2FieldHsfShift = 12
	RegisterConfrn2FieldHsfMask  = 0xf000
)

// GetHsf Horizontal Sampling Factor Horizontal sampling factor for component i.
func (r *registerConfrn2Type) GetHsf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn2FieldHsfMask) >> RegisterConfrn2FieldHsfShift)
}

// SetHsf Horizontal Sampling Factor Horizontal sampling factor for component i.
func (r *registerConfrn2Type) SetHsf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn2FieldHsfMask)|(uint32(value)<<RegisterConfrn2FieldHsfShift))
}

// registerConfrn3Type JPEG codec configuration register 4-7
type registerConfrn3Type uint32

const (
	RegisterConfrn3FieldHdShift = 0
	RegisterConfrn3FieldHdMask  = 0x1
)

// GetHd Huffman DC Selects the Huffman table for encoding the DC coefficients.
func (r *registerConfrn3Type) GetHd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterConfrn3FieldHdMask) != 0
}

// SetHd Huffman DC Selects the Huffman table for encoding the DC coefficients.
func (r *registerConfrn3Type) SetHd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterConfrn3FieldHdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterConfrn3FieldHdMask)
	}
}

const (
	RegisterConfrn3FieldHaShift = 1
	RegisterConfrn3FieldHaMask  = 0x2
)

// GetHa Huffman AC Selects the Huffman table for encoding the AC coefficients.
func (r *registerConfrn3Type) GetHa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterConfrn3FieldHaMask) != 0
}

// SetHa Huffman AC Selects the Huffman table for encoding the AC coefficients.
func (r *registerConfrn3Type) SetHa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterConfrn3FieldHaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterConfrn3FieldHaMask)
	}
}

const (
	RegisterConfrn3FieldQtShift = 2
	RegisterConfrn3FieldQtMask  = 0xc
)

// GetQt Quantization Table Selects quantization table associated with a color component.
func (r *registerConfrn3Type) GetQt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn3FieldQtMask) >> RegisterConfrn3FieldQtShift)
}

// SetQt Quantization Table Selects quantization table associated with a color component.
func (r *registerConfrn3Type) SetQt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn3FieldQtMask)|(uint32(value)<<RegisterConfrn3FieldQtShift))
}

const (
	RegisterConfrn3FieldNbShift = 4
	RegisterConfrn3FieldNbMask  = 0xf0
)

// GetNb Number of Block Number of data units minus 1 that belong to a particular color in the MCU.
func (r *registerConfrn3Type) GetNb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn3FieldNbMask) >> RegisterConfrn3FieldNbShift)
}

// SetNb Number of Block Number of data units minus 1 that belong to a particular color in the MCU.
func (r *registerConfrn3Type) SetNb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn3FieldNbMask)|(uint32(value)<<RegisterConfrn3FieldNbShift))
}

const (
	RegisterConfrn3FieldVsfShift = 8
	RegisterConfrn3FieldVsfMask  = 0xf00
)

// GetVsf Vertical Sampling Factor Vertical sampling factor for component i.
func (r *registerConfrn3Type) GetVsf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn3FieldVsfMask) >> RegisterConfrn3FieldVsfShift)
}

// SetVsf Vertical Sampling Factor Vertical sampling factor for component i.
func (r *registerConfrn3Type) SetVsf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn3FieldVsfMask)|(uint32(value)<<RegisterConfrn3FieldVsfShift))
}

const (
	RegisterConfrn3FieldHsfShift = 12
	RegisterConfrn3FieldHsfMask  = 0xf000
)

// GetHsf Horizontal Sampling Factor Horizontal sampling factor for component i.
func (r *registerConfrn3Type) GetHsf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn3FieldHsfMask) >> RegisterConfrn3FieldHsfShift)
}

// SetHsf Horizontal Sampling Factor Horizontal sampling factor for component i.
func (r *registerConfrn3Type) SetHsf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn3FieldHsfMask)|(uint32(value)<<RegisterConfrn3FieldHsfShift))
}

// registerConfrn4Type JPEG codec configuration register 4-7
type registerConfrn4Type uint32

const (
	RegisterConfrn4FieldHdShift = 0
	RegisterConfrn4FieldHdMask  = 0x1
)

// GetHd Huffman DC Selects the Huffman table for encoding the DC coefficients.
func (r *registerConfrn4Type) GetHd() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterConfrn4FieldHdMask) != 0
}

// SetHd Huffman DC Selects the Huffman table for encoding the DC coefficients.
func (r *registerConfrn4Type) SetHd(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterConfrn4FieldHdMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterConfrn4FieldHdMask)
	}
}

const (
	RegisterConfrn4FieldHaShift = 1
	RegisterConfrn4FieldHaMask  = 0x2
)

// GetHa Huffman AC Selects the Huffman table for encoding the AC coefficients.
func (r *registerConfrn4Type) GetHa() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterConfrn4FieldHaMask) != 0
}

// SetHa Huffman AC Selects the Huffman table for encoding the AC coefficients.
func (r *registerConfrn4Type) SetHa(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterConfrn4FieldHaMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterConfrn4FieldHaMask)
	}
}

const (
	RegisterConfrn4FieldQtShift = 2
	RegisterConfrn4FieldQtMask  = 0xc
)

// GetQt Quantization Table Selects quantization table associated with a color component.
func (r *registerConfrn4Type) GetQt() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn4FieldQtMask) >> RegisterConfrn4FieldQtShift)
}

// SetQt Quantization Table Selects quantization table associated with a color component.
func (r *registerConfrn4Type) SetQt(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn4FieldQtMask)|(uint32(value)<<RegisterConfrn4FieldQtShift))
}

const (
	RegisterConfrn4FieldNbShift = 4
	RegisterConfrn4FieldNbMask  = 0xf0
)

// GetNb Number of Block Number of data units minus 1 that belong to a particular color in the MCU.
func (r *registerConfrn4Type) GetNb() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn4FieldNbMask) >> RegisterConfrn4FieldNbShift)
}

// SetNb Number of Block Number of data units minus 1 that belong to a particular color in the MCU.
func (r *registerConfrn4Type) SetNb(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn4FieldNbMask)|(uint32(value)<<RegisterConfrn4FieldNbShift))
}

const (
	RegisterConfrn4FieldVsfShift = 8
	RegisterConfrn4FieldVsfMask  = 0xf00
)

// GetVsf Vertical Sampling Factor Vertical sampling factor for component i.
func (r *registerConfrn4Type) GetVsf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn4FieldVsfMask) >> RegisterConfrn4FieldVsfShift)
}

// SetVsf Vertical Sampling Factor Vertical sampling factor for component i.
func (r *registerConfrn4Type) SetVsf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn4FieldVsfMask)|(uint32(value)<<RegisterConfrn4FieldVsfShift))
}

const (
	RegisterConfrn4FieldHsfShift = 12
	RegisterConfrn4FieldHsfMask  = 0xf000
)

// GetHsf Horizontal Sampling Factor Horizontal sampling factor for component i.
func (r *registerConfrn4Type) GetHsf() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterConfrn4FieldHsfMask) >> RegisterConfrn4FieldHsfShift)
}

// SetHsf Horizontal Sampling Factor Horizontal sampling factor for component i.
func (r *registerConfrn4Type) SetHsf(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterConfrn4FieldHsfMask)|(uint32(value)<<RegisterConfrn4FieldHsfShift))
}

// registerCrType JPEG control register
type registerCrType uint32

const (
	RegisterCrFieldJcenShift = 0
	RegisterCrFieldJcenMask  = 0x1
)

// GetJcen JPEG Core Enable Enable the JPEG codec Core.
func (r *registerCrType) GetJcen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldJcenMask) != 0
}

// SetJcen JPEG Core Enable Enable the JPEG codec Core.
func (r *registerCrType) SetJcen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldJcenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldJcenMask)
	}
}

const (
	RegisterCrFieldIftieShift = 1
	RegisterCrFieldIftieMask  = 0x2
)

// GetIftie Input FIFO Threshold Interrupt Enable This bit enables the interrupt generation when input FIFO reach the threshold.
func (r *registerCrType) GetIftie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldIftieMask) != 0
}

// SetIftie Input FIFO Threshold Interrupt Enable This bit enables the interrupt generation when input FIFO reach the threshold.
func (r *registerCrType) SetIftie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldIftieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldIftieMask)
	}
}

const (
	RegisterCrFieldIfnfieShift = 2
	RegisterCrFieldIfnfieMask  = 0x4
)

// GetIfnfie Input FIFO Not Full Interrupt Enable This bit enables the interrupt generation when input FIFO is not empty.
func (r *registerCrType) GetIfnfie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldIfnfieMask) != 0
}

// SetIfnfie Input FIFO Not Full Interrupt Enable This bit enables the interrupt generation when input FIFO is not empty.
func (r *registerCrType) SetIfnfie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldIfnfieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldIfnfieMask)
	}
}

const (
	RegisterCrFieldOftieShift = 3
	RegisterCrFieldOftieMask  = 0x8
)

// GetOftie Output FIFO Threshold Interrupt Enable This bit enables the interrupt generation when output FIFO reach the threshold.
func (r *registerCrType) GetOftie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOftieMask) != 0
}

// SetOftie Output FIFO Threshold Interrupt Enable This bit enables the interrupt generation when output FIFO reach the threshold.
func (r *registerCrType) SetOftie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldOftieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldOftieMask)
	}
}

const (
	RegisterCrFieldOfneieShift = 4
	RegisterCrFieldOfneieMask  = 0x10
)

// GetOfneie Output FIFO Not Empty Interrupt Enable This bit enables the interrupt generation when output FIFO is not empty.
func (r *registerCrType) GetOfneie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOfneieMask) != 0
}

// SetOfneie Output FIFO Not Empty Interrupt Enable This bit enables the interrupt generation when output FIFO is not empty.
func (r *registerCrType) SetOfneie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldOfneieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldOfneieMask)
	}
}

const (
	RegisterCrFieldEocieShift = 5
	RegisterCrFieldEocieMask  = 0x20
)

// GetEocie End of Conversion Interrupt Enable This bit enables the interrupt generation on the end of conversion.
func (r *registerCrType) GetEocie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldEocieMask) != 0
}

// SetEocie End of Conversion Interrupt Enable This bit enables the interrupt generation on the end of conversion.
func (r *registerCrType) SetEocie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldEocieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldEocieMask)
	}
}

const (
	RegisterCrFieldHpdieShift = 6
	RegisterCrFieldHpdieMask  = 0x40
)

// GetHpdie Header Parsing Done Interrupt Enable This bit enables the interrupt generation on the Header Parsing Operation.
func (r *registerCrType) GetHpdie() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldHpdieMask) != 0
}

// SetHpdie Header Parsing Done Interrupt Enable This bit enables the interrupt generation on the Header Parsing Operation.
func (r *registerCrType) SetHpdie(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldHpdieMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldHpdieMask)
	}
}

const (
	RegisterCrFieldIdmaenShift = 11
	RegisterCrFieldIdmaenMask  = 0x800
)

// GetIdmaen Input DMA Enable Enable the DMA request generation for the input FIFO.
func (r *registerCrType) GetIdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldIdmaenMask) != 0
}

// SetIdmaen Input DMA Enable Enable the DMA request generation for the input FIFO.
func (r *registerCrType) SetIdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldIdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldIdmaenMask)
	}
}

const (
	RegisterCrFieldOdmaenShift = 12
	RegisterCrFieldOdmaenMask  = 0x1000
)

// GetOdmaen Output DMA Enable Enable the DMA request generation for the output FIFO.
func (r *registerCrType) GetOdmaen() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOdmaenMask) != 0
}

// SetOdmaen Output DMA Enable Enable the DMA request generation for the output FIFO.
func (r *registerCrType) SetOdmaen(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldOdmaenMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldOdmaenMask)
	}
}

const (
	RegisterCrFieldIffShift = 13
	RegisterCrFieldIffMask  = 0x2000
)

// GetIff Input FIFO Flush This bit flush the input FIFO. This bit is always read as 0.
func (r *registerCrType) GetIff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldIffMask) != 0
}

// SetIff Input FIFO Flush This bit flush the input FIFO. This bit is always read as 0.
func (r *registerCrType) SetIff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldIffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldIffMask)
	}
}

const (
	RegisterCrFieldOffShift = 14
	RegisterCrFieldOffMask  = 0x4000
)

// GetOff Output FIFO Flush This bit flush the output FIFO. This bit is always read as 0.
func (r *registerCrType) GetOff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCrFieldOffMask) != 0
}

// SetOff Output FIFO Flush This bit flush the output FIFO. This bit is always read as 0.
func (r *registerCrType) SetOff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCrFieldOffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCrFieldOffMask)
	}
}

// registerSrType JPEG status register
type registerSrType uint32

const (
	RegisterSrFieldIftfShift = 1
	RegisterSrFieldIftfMask  = 0x2
)

// GetIftf Input FIFO Threshold Flag This bit is set when the input FIFO is not full and is bellow its threshold.
func (r *registerSrType) GetIftf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIftfMask) != 0
}

// SetIftf Input FIFO Threshold Flag This bit is set when the input FIFO is not full and is bellow its threshold.
func (r *registerSrType) SetIftf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldIftfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldIftfMask)
	}
}

const (
	RegisterSrFieldIfnffShift = 2
	RegisterSrFieldIfnffMask  = 0x4
)

// GetIfnff Input FIFO Not Full Flag This bit is set when the input FIFO is not full (a data can be written).
func (r *registerSrType) GetIfnff() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldIfnffMask) != 0
}

// SetIfnff Input FIFO Not Full Flag This bit is set when the input FIFO is not full (a data can be written).
func (r *registerSrType) SetIfnff(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldIfnffMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldIfnffMask)
	}
}

const (
	RegisterSrFieldOftfShift = 3
	RegisterSrFieldOftfMask  = 0x8
)

// GetOftf Output FIFO Threshold Flag This bit is set when the output FIFO is not empty and has reach its threshold.
func (r *registerSrType) GetOftf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldOftfMask) != 0
}

// SetOftf Output FIFO Threshold Flag This bit is set when the output FIFO is not empty and has reach its threshold.
func (r *registerSrType) SetOftf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldOftfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldOftfMask)
	}
}

const (
	RegisterSrFieldOfnefShift = 4
	RegisterSrFieldOfnefMask  = 0x10
)

// GetOfnef Output FIFO Not Empty Flag This bit is set when the output FIFO is not empty (a data is available).
func (r *registerSrType) GetOfnef() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldOfnefMask) != 0
}

// SetOfnef Output FIFO Not Empty Flag This bit is set when the output FIFO is not empty (a data is available).
func (r *registerSrType) SetOfnef(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldOfnefMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldOfnefMask)
	}
}

const (
	RegisterSrFieldEocfShift = 5
	RegisterSrFieldEocfMask  = 0x20
)

// GetEocf End of Conversion Flag This bit is set when the JPEG codec core has finished the encoding or the decoding process and than last data has been sent to the output FIFO.
func (r *registerSrType) GetEocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldEocfMask) != 0
}

// SetEocf End of Conversion Flag This bit is set when the JPEG codec core has finished the encoding or the decoding process and than last data has been sent to the output FIFO.
func (r *registerSrType) SetEocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldEocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldEocfMask)
	}
}

const (
	RegisterSrFieldHpdfShift = 6
	RegisterSrFieldHpdfMask  = 0x40
)

// GetHpdf Header Parsing Done Flag This bit is set in decode mode when the JPEG codec has finished the parsing of the headers and the internal registers have been updated.
func (r *registerSrType) GetHpdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldHpdfMask) != 0
}

// SetHpdf Header Parsing Done Flag This bit is set in decode mode when the JPEG codec has finished the parsing of the headers and the internal registers have been updated.
func (r *registerSrType) SetHpdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldHpdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldHpdfMask)
	}
}

const (
	RegisterSrFieldCofShift = 7
	RegisterSrFieldCofMask  = 0x80
)

// GetCof Codec Operation Flag This bit is set when when a JPEG codec operation is on going (encoding or decoding).
func (r *registerSrType) GetCof() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterSrFieldCofMask) != 0
}

// SetCof Codec Operation Flag This bit is set when when a JPEG codec operation is on going (encoding or decoding).
func (r *registerSrType) SetCof(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterSrFieldCofMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterSrFieldCofMask)
	}
}

// registerCfrType JPEG clear flag register
type registerCfrType uint32

const (
	RegisterCfrFieldCeocfShift = 5
	RegisterCfrFieldCeocfMask  = 0x20
)

// GetCeocf Clear End of Conversion Flag Writing 1 clears the End of Conversion Flag of the JPEG Status Register.
func (r *registerCfrType) GetCeocf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfrFieldCeocfMask) != 0
}

// SetCeocf Clear End of Conversion Flag Writing 1 clears the End of Conversion Flag of the JPEG Status Register.
func (r *registerCfrType) SetCeocf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfrFieldCeocfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfrFieldCeocfMask)
	}
}

const (
	RegisterCfrFieldChpdfShift = 6
	RegisterCfrFieldChpdfMask  = 0x40
)

// GetChpdf Clear Header Parsing Done Flag Writing 1 clears the Header Parsing Done Flag of the JPEG Status Register.
func (r *registerCfrType) GetChpdf() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCfrFieldChpdfMask) != 0
}

// SetChpdf Clear Header Parsing Done Flag Writing 1 clears the Header Parsing Done Flag of the JPEG Status Register.
func (r *registerCfrType) SetChpdf(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCfrFieldChpdfMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCfrFieldChpdfMask)
	}
}

// registerDirType JPEG data input register
type registerDirType uint32

const (
	RegisterDirFieldDatainShift = 0
	RegisterDirFieldDatainMask  = 0xffffffff
)

// GetDatain Data Input FIFO Input FIFO data register.
func (r *registerDirType) GetDatain() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDirFieldDatainMask) >> RegisterDirFieldDatainShift)
}

// SetDatain Data Input FIFO Input FIFO data register.
func (r *registerDirType) SetDatain(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDirFieldDatainMask)|(uint32(value)<<RegisterDirFieldDatainShift))
}

// registerDorType JPEG data output register
type registerDorType uint32

const (
	RegisterDorFieldDataoutShift = 0
	RegisterDorFieldDataoutMask  = 0xffffffff
)

// GetDataout Data Output FIFO Output FIFO data register.
func (r *registerDorType) GetDataout() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterDorFieldDataoutMask) >> RegisterDorFieldDataoutShift)
}

// SetDataout Data Output FIFO Output FIFO data register.
func (r *registerDorType) SetDataout(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDorFieldDataoutMask)|(uint32(value)<<RegisterDorFieldDataoutShift))
}
