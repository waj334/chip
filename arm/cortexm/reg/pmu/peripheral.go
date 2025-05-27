//go:build arm && cortexm

package pmu

import (
	"unsafe"
	"volatile"
)

var (
	Pmu = (*_pmu)(unsafe.Pointer(uintptr(0xe0003000)))
)

type _pmu struct {
	Evcntr     [8]registerEvcntrType
	_          [92]byte
	Ccntr      registerCcntrType
	_          [896]byte
	Evtyper    [8]registerEvtyperType
	_          [92]byte
	Ccfiltr    registerCcfiltrType
	_          [1920]byte
	Cntenset   registerCntensetType
	_          [28]byte
	Cntenclr   registerCntenclrType
	_          [28]byte
	Intenset   registerIntensetType
	_          [28]byte
	Intenclr   registerIntenclrType
	_          [28]byte
	Ovsclr     registerOvsclrType
	_          [28]byte
	Swinc      registerSwincType
	_          [28]byte
	Ovsset     registerOvssetType
	_          [316]byte
	Type       registerTypeType
	Ctrl       registerCtrlType
	_          [432]byte
	Authstatus registerAuthstatusType
	Devarch    registerDevarchType
	_          [12]byte
	Devtype    registerDevtypeType
	Pidr4      registerPidr4Type
	_          [12]byte
	Pidr0      registerPidr0Type
	Pidr1      registerPidr1Type
	Pidr2      registerPidr2Type
	Pidr3      registerPidr3Type
	Cidr0      registerCidr0Type
	Cidr1      registerCidr1Type
	Cidr2      registerCidr2Type
	Cidr3      registerCidr3Type
}

// registerEvcntrType Holds performance counter, which count events
type registerEvcntrType uint32

const (
	RegisterEvcntrFieldCounterShift = 0
	RegisterEvcntrFieldCounterMask  = 0xffff
)

// GetCounter The counter counts whenever the selected event occurs, and either of the following conditions are met: SecureNoninvasiveDebugAllowed()==TRUE; The source of the event is Non-secure and NoninvasiveDebugAllowed()==TRUE.
func (r *registerEvcntrType) GetCounter() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterEvcntrFieldCounterMask) >> RegisterEvcntrFieldCounterShift)
}

// SetCounter The counter counts whenever the selected event occurs, and either of the following conditions are met: SecureNoninvasiveDebugAllowed()==TRUE; The source of the event is Non-secure and NoninvasiveDebugAllowed()==TRUE.
func (r *registerEvcntrType) SetCounter(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEvcntrFieldCounterMask)|(uint32(value)<<RegisterEvcntrFieldCounterShift))
}

// registerCcntrType Holds the value of the cycle counter, which counts processor clock cycles
type registerCcntrType uint32

// registerEvtyperType Configures the event counter
type registerEvtyperType uint32

const (
	RegisterEvtyperFieldEvtcountShift = 0
	RegisterEvtyperFieldEvtcountMask  = 0xffff
)

// GetEvtcount The event number of the event that is counted by event counter PMU_EVCNTR0-7
func (r *registerEvtyperType) GetEvtcount() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterEvtyperFieldEvtcountMask) >> RegisterEvtyperFieldEvtcountShift)
}

// SetEvtcount The event number of the event that is counted by event counter PMU_EVCNTR0-7
func (r *registerEvtyperType) SetEvtcount(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterEvtyperFieldEvtcountMask)|(uint32(value)<<RegisterEvtyperFieldEvtcountShift))
}

// registerCcfiltrType This register is RES0
type registerCcfiltrType uint32

// registerCntensetType Enables the PMU_CCNTR register and PMU_EVCNTR0-7 registers
type registerCntensetType uint32

const (
	RegisterCntensetFieldPnShift = 0
	RegisterCntensetFieldPnMask  = 0xff
)

// GetPn Event counter PMU_EVCNTR0-7 enable bit
func (r *registerCntensetType) GetPn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCntensetFieldPnMask) >> RegisterCntensetFieldPnShift)
}

// SetPn Event counter PMU_EVCNTR0-7 enable bit
func (r *registerCntensetType) SetPn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntensetFieldPnMask)|(uint32(value)<<RegisterCntensetFieldPnShift))
}

const (
	RegisterCntensetFieldCShift = 8
	RegisterCntensetFieldCMask  = 0xffffff00
)

// GetC CCNTR enable bit
func (r *registerCntensetType) GetC() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCntensetFieldCMask) >> RegisterCntensetFieldCShift)
}

// SetC CCNTR enable bit
func (r *registerCntensetType) SetC(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntensetFieldCMask)|(uint32(value)<<RegisterCntensetFieldCShift))
}

// registerCntenclrType Disables the PMU_CCNTR register and PMU_EVCNTR0-7 registers
type registerCntenclrType uint32

const (
	RegisterCntenclrFieldPnShift = 0
	RegisterCntenclrFieldPnMask  = 0xff
)

// GetPn Event counter PMU_EVCNTR0-7 enable bit
func (r *registerCntenclrType) GetPn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCntenclrFieldPnMask) >> RegisterCntenclrFieldPnShift)
}

// SetPn Event counter PMU_EVCNTR0-7 enable bit
func (r *registerCntenclrType) SetPn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntenclrFieldPnMask)|(uint32(value)<<RegisterCntenclrFieldPnShift))
}

const (
	RegisterCntenclrFieldCShift = 8
	RegisterCntenclrFieldCMask  = 0xffffff00
)

// GetC CCNTR disable bit
func (r *registerCntenclrType) GetC() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterCntenclrFieldCMask) >> RegisterCntenclrFieldCShift)
}

// SetC CCNTR disable bit
func (r *registerCntenclrType) SetC(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCntenclrFieldCMask)|(uint32(value)<<RegisterCntenclrFieldCShift))
}

// registerIntensetType Enables the generation of interrupt requests on overflows from the PMU_CCNTR, and the event counter, PMU_EVCNTR.
type registerIntensetType uint32

const (
	RegisterIntensetFieldPnShift = 0
	RegisterIntensetFieldPnMask  = 0xff
)

// GetPn Event counter overflow interrupt request disable bit for PMU_EVCNTR0-7
func (r *registerIntensetType) GetPn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIntensetFieldPnMask) >> RegisterIntensetFieldPnShift)
}

// SetPn Event counter overflow interrupt request disable bit for PMU_EVCNTR0-7
func (r *registerIntensetType) SetPn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIntensetFieldPnMask)|(uint32(value)<<RegisterIntensetFieldPnShift))
}

const (
	RegisterIntensetFieldCShift = 8
	RegisterIntensetFieldCMask  = 0xffffff00
)

// GetC CCNTR interrupt request enable bit
func (r *registerIntensetType) GetC() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIntensetFieldCMask) >> RegisterIntensetFieldCShift)
}

// SetC CCNTR interrupt request enable bit
func (r *registerIntensetType) SetC(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIntensetFieldCMask)|(uint32(value)<<RegisterIntensetFieldCShift))
}

// registerIntenclrType Disables the generation of interrupt requests on overflows from the PMU_CCNTR, and the event counters, PMU_EVCNTR
type registerIntenclrType uint32

const (
	RegisterIntenclrFieldPnShift = 0
	RegisterIntenclrFieldPnMask  = 0xff
)

// GetPn Event counter overflow interrupt request enable bit for PMU_EVCNTRn. Enable the overflow interrupt for PMU_EVCNTR0-7
func (r *registerIntenclrType) GetPn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterIntenclrFieldPnMask) >> RegisterIntenclrFieldPnShift)
}

// SetPn Event counter overflow interrupt request enable bit for PMU_EVCNTRn. Enable the overflow interrupt for PMU_EVCNTR0-7
func (r *registerIntenclrType) SetPn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIntenclrFieldPnMask)|(uint32(value)<<RegisterIntenclrFieldPnShift))
}

const (
	RegisterIntenclrFieldCShift = 8
	RegisterIntenclrFieldCMask  = 0xffffff00
)

// GetC CCNTR overflow interrupt request disable bit
func (r *registerIntenclrType) GetC() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterIntenclrFieldCMask) >> RegisterIntenclrFieldCShift)
}

// SetC CCNTR overflow interrupt request disable bit
func (r *registerIntenclrType) SetC(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterIntenclrFieldCMask)|(uint32(value)<<RegisterIntenclrFieldCShift))
}

// registerOvsclrType Contains the state of the overflow bit for the PMU_CCNTR, and each of the implemented event counters, PMU_EVCNTRn. Writing to this register clears these bits.
type registerOvsclrType uint32

const (
	RegisterOvsclrFieldPnShift = 0
	RegisterOvsclrFieldPnMask  = 0x7fffffff
)

// GetPn Event counter overflow clear bit for PMU_EVCNTRn. Clears the PMU_EVCNTRn overflow bit
func (r *registerOvsclrType) GetPn() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOvsclrFieldPnMask) >> RegisterOvsclrFieldPnShift)
}

// SetPn Event counter overflow clear bit for PMU_EVCNTRn. Clears the PMU_EVCNTRn overflow bit
func (r *registerOvsclrType) SetPn(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOvsclrFieldPnMask)|(uint32(value)<<RegisterOvsclrFieldPnShift))
}

const (
	RegisterOvsclrFieldCShift = 31
	RegisterOvsclrFieldCMask  = 0x80000000
)

// GetC CCNTR overflow bit. Clears the PMU_CCNTR overflow bit.
func (r *registerOvsclrType) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOvsclrFieldCMask) != 0
}

// SetC CCNTR overflow bit. Clears the PMU_CCNTR overflow bit.
func (r *registerOvsclrType) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOvsclrFieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOvsclrFieldCMask)
	}
}

// registerSwincType Increments a counter that is configured to count the software increment event, event 0x00
type registerSwincType uint32

const (
	RegisterSwincFieldPnShift = 0
	RegisterSwincFieldPnMask  = 0xff
)

// GetPn Event counter software increment bit for PMU_EVCNTRn. An event counter n, configured for SW_INCR events, increments on every write to bit n of this field.
func (r *registerSwincType) GetPn() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterSwincFieldPnMask) >> RegisterSwincFieldPnShift)
}

// SetPn Event counter software increment bit for PMU_EVCNTRn. An event counter n, configured for SW_INCR events, increments on every write to bit n of this field.
func (r *registerSwincType) SetPn(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterSwincFieldPnMask)|(uint32(value)<<RegisterSwincFieldPnShift))
}

// registerOvssetType Performance Monitoring Unit Overflow Flag Status Set Register
type registerOvssetType uint32

const (
	RegisterOvssetFieldPnShift = 0
	RegisterOvssetFieldPnMask  = 0x7fffffff
)

// GetPn Event counter overflow set bit for PMU_EVCNTRn. Sets the overflow status for PMU_EVCNTRn.
func (r *registerOvssetType) GetPn() uint32 {
	return uint32((volatile.LoadUint32((*uint32)(r)) & RegisterOvssetFieldPnMask) >> RegisterOvssetFieldPnShift)
}

// SetPn Event counter overflow set bit for PMU_EVCNTRn. Sets the overflow status for PMU_EVCNTRn.
func (r *registerOvssetType) SetPn(value uint32) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterOvssetFieldPnMask)|(uint32(value)<<RegisterOvssetFieldPnShift))
}

const (
	RegisterOvssetFieldCShift = 31
	RegisterOvssetFieldCMask  = 0x80000000
)

// GetC CCNTR overflow bit. Sets the overflow status for PMU_CCNTR.
func (r *registerOvssetType) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterOvssetFieldCMask) != 0
}

// SetC CCNTR overflow bit. Sets the overflow status for PMU_CCNTR.
func (r *registerOvssetType) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterOvssetFieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterOvssetFieldCMask)
	}
}

// registerTypeType Contains information regarding what the PMU supports
type registerTypeType uint32

const (
	RegisterTypeFieldNShift = 0
	RegisterTypeFieldNMask  = 0xff
)

// GetN Number of counters implemented in addition to the cycle counter, PMU_CCNTR. This field is set to 0b00001000, indicating that 8 16-bit event counters are present in addition to PMU_CCNTR.
func (r *registerTypeType) GetN() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTypeFieldNMask) >> RegisterTypeFieldNShift)
}

// SetN Number of counters implemented in addition to the cycle counter, PMU_CCNTR. This field is set to 0b00001000, indicating that 8 16-bit event counters are present in addition to PMU_CCNTR.
func (r *registerTypeType) SetN(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTypeFieldNMask)|(uint32(value)<<RegisterTypeFieldNShift))
}

const (
	RegisterTypeFieldSizeShift = 8
	RegisterTypeFieldSizeMask  = 0x3f00
)

// GetSize Size of counters. This field determines the spacing of counters in the memory-map. This field is 0b011111.
func (r *registerTypeType) GetSize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterTypeFieldSizeMask) >> RegisterTypeFieldSizeShift)
}

// SetSize Size of counters. This field determines the spacing of counters in the memory-map. This field is 0b011111.
func (r *registerTypeType) SetSize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterTypeFieldSizeMask)|(uint32(value)<<RegisterTypeFieldSizeShift))
}

const (
	RegisterTypeFieldCcShift = 14
	RegisterTypeFieldCcMask  = 0x4000
)

// GetCc A dedicated cycle counter is present. This bit is 0b1.
func (r *registerTypeType) GetCc() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTypeFieldCcMask) != 0
}

// SetCc A dedicated cycle counter is present. This bit is 0b1.
func (r *registerTypeType) SetCc(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTypeFieldCcMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTypeFieldCcMask)
	}
}

const (
	RegisterTypeFieldFzoShift = 21
	RegisterTypeFieldFzoMask  = 0x200000
)

// GetFzo Freeze-on-overflow is supported. This bit is 0b1.
func (r *registerTypeType) GetFzo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTypeFieldFzoMask) != 0
}

// SetFzo Freeze-on-overflow is supported. This bit is 0b1.
func (r *registerTypeType) SetFzo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTypeFieldFzoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTypeFieldFzoMask)
	}
}

const (
	RegisterTypeFieldTroShift = 23
	RegisterTypeFieldTroMask  = 0x800000
)

// GetTro Trace-on-overflow not supported. This bit is 0b0.
func (r *registerTypeType) GetTro() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterTypeFieldTroMask) != 0
}

// SetTro Trace-on-overflow not supported. This bit is 0b0.
func (r *registerTypeType) SetTro(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterTypeFieldTroMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterTypeFieldTroMask)
	}
}

// registerCtrlType Configures and controls the PMU
type registerCtrlType uint32

const (
	RegisterCtrlFieldEShift = 0
	RegisterCtrlFieldEMask  = 0x1
)

// GetE Enable. Enable the event counters.
func (r *registerCtrlType) GetE() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldEMask) != 0
}

// SetE Enable. Enable the event counters.
func (r *registerCtrlType) SetE(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldEMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldEMask)
	}
}

const (
	RegisterCtrlFieldPShift = 1
	RegisterCtrlFieldPMask  = 0x2
)

// GetP Event counter reset
func (r *registerCtrlType) GetP() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldPMask) != 0
}

// SetP Event counter reset
func (r *registerCtrlType) SetP(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldPMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldPMask)
	}
}

const (
	RegisterCtrlFieldCShift = 2
	RegisterCtrlFieldCMask  = 0x4
)

// GetC Cycle counter reset. Reset the PMU_CCNTR counter.
func (r *registerCtrlType) GetC() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldCMask) != 0
}

// SetC Cycle counter reset. Reset the PMU_CCNTR counter.
func (r *registerCtrlType) SetC(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldCMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldCMask)
	}
}

const (
	RegisterCtrlFieldDpShift = 5
	RegisterCtrlFieldDpMask  = 0x20
)

// GetDp Disable cycle counter when event counting is prohibited
func (r *registerCtrlType) GetDp() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldDpMask) != 0
}

// SetDp Disable cycle counter when event counting is prohibited
func (r *registerCtrlType) SetDp(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldDpMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldDpMask)
	}
}

const (
	RegisterCtrlFieldFzoShift = 9
	RegisterCtrlFieldFzoMask  = 0x200
)

// GetFzo Freeze-on-overflow. Stops events being counted once PMU_OVSCLR or PMU_OVSSET is non-zero.
func (r *registerCtrlType) GetFzo() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldFzoMask) != 0
}

// SetFzo Freeze-on-overflow. Stops events being counted once PMU_OVSCLR or PMU_OVSSET is non-zero.
func (r *registerCtrlType) SetFzo(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldFzoMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldFzoMask)
	}
}

const (
	RegisterCtrlFieldTroShift = 11
	RegisterCtrlFieldTroMask  = 0x800
)

// GetTro Trace-on-overflow not supported in Cortex-M55. Therefore, this bit is 0b0.
func (r *registerCtrlType) GetTro() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterCtrlFieldTroMask) != 0
}

// SetTro Trace-on-overflow not supported in Cortex-M55. Therefore, this bit is 0b0.
func (r *registerCtrlType) SetTro(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterCtrlFieldTroMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterCtrlFieldTroMask)
	}
}

// registerAuthstatusType Provides information about the state of the implementation defined authentication interface for the PMU
type registerAuthstatusType uint32

const (
	RegisterAuthstatusFieldNsidShift = 0
	RegisterAuthstatusFieldNsidMask  = 0x3
)

// GetNsid Non-secure Invasive Debug. Indicates whether Non-secure invasive debug is allowed.
func (r *registerAuthstatusType) GetNsid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAuthstatusFieldNsidMask) >> RegisterAuthstatusFieldNsidShift)
}

// SetNsid Non-secure Invasive Debug. Indicates whether Non-secure invasive debug is allowed.
func (r *registerAuthstatusType) SetNsid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAuthstatusFieldNsidMask)|(uint32(value)<<RegisterAuthstatusFieldNsidShift))
}

const (
	RegisterAuthstatusFieldNsnidShift = 2
	RegisterAuthstatusFieldNsnidMask  = 0xc
)

// GetNsnid Non-secure Non-invasive Debug. Indicates whether Non-secure Non-invasive debug is allowed.
func (r *registerAuthstatusType) GetNsnid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAuthstatusFieldNsnidMask) >> RegisterAuthstatusFieldNsnidShift)
}

// SetNsnid Non-secure Non-invasive Debug. Indicates whether Non-secure Non-invasive debug is allowed.
func (r *registerAuthstatusType) SetNsnid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAuthstatusFieldNsnidMask)|(uint32(value)<<RegisterAuthstatusFieldNsnidShift))
}

const (
	RegisterAuthstatusFieldSidShift = 4
	RegisterAuthstatusFieldSidMask  = 0x30
)

// GetSid Secure Invasive Debug. Indicates whether Secure invasive debug is implemented and allowed.
func (r *registerAuthstatusType) GetSid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAuthstatusFieldSidMask) >> RegisterAuthstatusFieldSidShift)
}

// SetSid Secure Invasive Debug. Indicates whether Secure invasive debug is implemented and allowed.
func (r *registerAuthstatusType) SetSid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAuthstatusFieldSidMask)|(uint32(value)<<RegisterAuthstatusFieldSidShift))
}

const (
	RegisterAuthstatusFieldSnidShift = 6
	RegisterAuthstatusFieldSnidMask  = 0xc0
)

// GetSnid Secure Non-invasive Debug. Indicates whether Secure Non-invasive debug is implemented and allowed.
func (r *registerAuthstatusType) GetSnid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAuthstatusFieldSnidMask) >> RegisterAuthstatusFieldSnidShift)
}

// SetSnid Secure Non-invasive Debug. Indicates whether Secure Non-invasive debug is implemented and allowed.
func (r *registerAuthstatusType) SetSnid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAuthstatusFieldSnidMask)|(uint32(value)<<RegisterAuthstatusFieldSnidShift))
}

const (
	RegisterAuthstatusFieldNsuidShift = 16
	RegisterAuthstatusFieldNsuidMask  = 0x30000
)

// GetNsuid Non-secure Unprivileged Invasive Debug Allowed. Indicates that Unprivileged Halting Debug is allowed for the Non-secure state.
func (r *registerAuthstatusType) GetNsuid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAuthstatusFieldNsuidMask) >> RegisterAuthstatusFieldNsuidShift)
}

// SetNsuid Non-secure Unprivileged Invasive Debug Allowed. Indicates that Unprivileged Halting Debug is allowed for the Non-secure state.
func (r *registerAuthstatusType) SetNsuid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAuthstatusFieldNsuidMask)|(uint32(value)<<RegisterAuthstatusFieldNsuidShift))
}

const (
	RegisterAuthstatusFieldNsunidShift = 18
	RegisterAuthstatusFieldNsunidMask  = 0xc0000
)

// GetNsunid Non-secure Unprivileged Non-invasive Debug Allowed. Indicates that Unprivileged Non-invasive Debug is allowed for the Non-secure state.
func (r *registerAuthstatusType) GetNsunid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAuthstatusFieldNsunidMask) >> RegisterAuthstatusFieldNsunidShift)
}

// SetNsunid Non-secure Unprivileged Non-invasive Debug Allowed. Indicates that Unprivileged Non-invasive Debug is allowed for the Non-secure state.
func (r *registerAuthstatusType) SetNsunid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAuthstatusFieldNsunidMask)|(uint32(value)<<RegisterAuthstatusFieldNsunidShift))
}

const (
	RegisterAuthstatusFieldSuidShift = 20
	RegisterAuthstatusFieldSuidMask  = 0x300000
)

// GetSuid Secure Unprivileged Invasive Debug Allowed. Indicates that Unprivileged Invasive Debug is allowed for the Secure state.
func (r *registerAuthstatusType) GetSuid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAuthstatusFieldSuidMask) >> RegisterAuthstatusFieldSuidShift)
}

// SetSuid Secure Unprivileged Invasive Debug Allowed. Indicates that Unprivileged Invasive Debug is allowed for the Secure state.
func (r *registerAuthstatusType) SetSuid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAuthstatusFieldSuidMask)|(uint32(value)<<RegisterAuthstatusFieldSuidShift))
}

const (
	RegisterAuthstatusFieldSunidShift = 22
	RegisterAuthstatusFieldSunidMask  = 0xc00000
)

// GetSunid Secure Unprivileged Non-invasive Debug Allowed. Indicates that Unprivileged Non-invasive Debug is allowed for the Secure state.
func (r *registerAuthstatusType) GetSunid() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterAuthstatusFieldSunidMask) >> RegisterAuthstatusFieldSunidShift)
}

// SetSunid Secure Unprivileged Non-invasive Debug Allowed. Indicates that Unprivileged Non-invasive Debug is allowed for the Secure state.
func (r *registerAuthstatusType) SetSunid(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterAuthstatusFieldSunidMask)|(uint32(value)<<RegisterAuthstatusFieldSunidShift))
}

// registerDevarchType Identifies the programmers model architecture of the PMU
type registerDevarchType uint32

const (
	RegisterDevarchFieldArchidShift = 0
	RegisterDevarchFieldArchidMask  = 0xffff
)

// GetArchid Defines this part to be an Armv8‑M debug component. For the PMU, bits [15:12] are 0x0, bits [11:0] are 0xA06.
func (r *registerDevarchType) GetArchid() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDevarchFieldArchidMask) >> RegisterDevarchFieldArchidShift)
}

// SetArchid Defines this part to be an Armv8‑M debug component. For the PMU, bits [15:12] are 0x0, bits [11:0] are 0xA06.
func (r *registerDevarchType) SetArchid(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDevarchFieldArchidMask)|(uint32(value)<<RegisterDevarchFieldArchidShift))
}

const (
	RegisterDevarchFieldRevisionShift = 16
	RegisterDevarchFieldRevisionMask  = 0xf0000
)

// GetRevision Defines the architecture revision. For the PMU, the revision defined by Armv8.1‑M is 0x2.
func (r *registerDevarchType) GetRevision() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDevarchFieldRevisionMask) >> RegisterDevarchFieldRevisionShift)
}

// SetRevision Defines the architecture revision. For the PMU, the revision defined by Armv8.1‑M is 0x2.
func (r *registerDevarchType) SetRevision(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDevarchFieldRevisionMask)|(uint32(value)<<RegisterDevarchFieldRevisionShift))
}

const (
	RegisterDevarchFieldPresentShift = 20
	RegisterDevarchFieldPresentMask  = 0x100000
)

// GetPresent Determines the presence of DEVARCH. When set to 1, indicates that the DEVARCH is present. This bit reads as 0x1.
func (r *registerDevarchType) GetPresent() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterDevarchFieldPresentMask) != 0
}

// SetPresent Determines the presence of DEVARCH. When set to 1, indicates that the DEVARCH is present. This bit reads as 0x1.
func (r *registerDevarchType) SetPresent(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterDevarchFieldPresentMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterDevarchFieldPresentMask)
	}
}

const (
	RegisterDevarchFieldArchitectShift = 21
	RegisterDevarchFieldArchitectMask  = 0xffe00000
)

// GetArchitect Defines the architecture of the component. For the PMU, it is Arm Limited. Bits[31:28] are the JEP 106 continuation code, 0x4, and bits [27:21] are the JEP 106 ID code, 0x3B.
func (r *registerDevarchType) GetArchitect() uint16 {
	return uint16((volatile.LoadUint32((*uint32)(r)) & RegisterDevarchFieldArchitectMask) >> RegisterDevarchFieldArchitectShift)
}

// SetArchitect Defines the architecture of the component. For the PMU, it is Arm Limited. Bits[31:28] are the JEP 106 continuation code, 0x4, and bits [27:21] are the JEP 106 ID code, 0x3B.
func (r *registerDevarchType) SetArchitect(value uint16) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDevarchFieldArchitectMask)|(uint32(value)<<RegisterDevarchFieldArchitectShift))
}

// registerDevtypeType Indicates to a debugger that the PMU is a part of the processor interface
type registerDevtypeType uint32

const (
	RegisterDevtypeFieldMajorShift = 0
	RegisterDevtypeFieldMajorMask  = 0xf
)

// GetMajor Major type. This field reads as 0x6.
func (r *registerDevtypeType) GetMajor() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDevtypeFieldMajorMask) >> RegisterDevtypeFieldMajorShift)
}

// SetMajor Major type. This field reads as 0x6.
func (r *registerDevtypeType) SetMajor(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDevtypeFieldMajorMask)|(uint32(value)<<RegisterDevtypeFieldMajorShift))
}

const (
	RegisterDevtypeFieldSubShift = 4
	RegisterDevtypeFieldSubMask  = 0xf0
)

// GetSub Subtype. This field reads as 0x1.
func (r *registerDevtypeType) GetSub() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterDevtypeFieldSubMask) >> RegisterDevtypeFieldSubShift)
}

// SetSub Subtype. This field reads as 0x1.
func (r *registerDevtypeType) SetSub(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterDevtypeFieldSubMask)|(uint32(value)<<RegisterDevtypeFieldSubShift))
}

// registerPidr4Type Provides information to identify the PMU.
type registerPidr4Type uint32

const (
	RegisterPidr4FieldDes2Shift = 0
	RegisterPidr4FieldDes2Mask  = 0xf
)

// GetDes2 Designer, JEP 106 continuation code, least significant nibble. For Arm Limited, this field is 0x4.
func (r *registerPidr4Type) GetDes2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr4FieldDes2Mask) >> RegisterPidr4FieldDes2Shift)
}

// SetDes2 Designer, JEP 106 continuation code, least significant nibble. For Arm Limited, this field is 0x4.
func (r *registerPidr4Type) SetDes2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr4FieldDes2Mask)|(uint32(value)<<RegisterPidr4FieldDes2Shift))
}

const (
	RegisterPidr4FieldSizeShift = 4
	RegisterPidr4FieldSizeMask  = 0xf0
)

// GetSize Size of the component. This field is RAZ.
func (r *registerPidr4Type) GetSize() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr4FieldSizeMask) >> RegisterPidr4FieldSizeShift)
}

// SetSize Size of the component. This field is RAZ.
func (r *registerPidr4Type) SetSize(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr4FieldSizeMask)|(uint32(value)<<RegisterPidr4FieldSizeShift))
}

// registerPidr0Type Provides information to identify the PMU
type registerPidr0Type uint32

const (
	RegisterPidr0FieldPart0Shift = 4
	RegisterPidr0FieldPart0Mask  = 0xf0
)

// GetPart0 Part number, least significant byte. This field reads 0x22.
func (r *registerPidr0Type) GetPart0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr0FieldPart0Mask) >> RegisterPidr0FieldPart0Shift)
}

// SetPart0 Part number, least significant byte. This field reads 0x22.
func (r *registerPidr0Type) SetPart0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr0FieldPart0Mask)|(uint32(value)<<RegisterPidr0FieldPart0Shift))
}

// registerPidr1Type Provides information to identify the PMU
type registerPidr1Type uint32

const (
	RegisterPidr1FieldPart1Shift = 0
	RegisterPidr1FieldPart1Mask  = 0xf
)

// GetPart1 Part number, most significant nibble. This field is 0xD.
func (r *registerPidr1Type) GetPart1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr1FieldPart1Mask) >> RegisterPidr1FieldPart1Shift)
}

// SetPart1 Part number, most significant nibble. This field is 0xD.
func (r *registerPidr1Type) SetPart1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr1FieldPart1Mask)|(uint32(value)<<RegisterPidr1FieldPart1Shift))
}

const (
	RegisterPidr1FieldDes0Shift = 4
	RegisterPidr1FieldDes0Mask  = 0xf0
)

// GetDes0 Designer, least significant nibble of JEP 106 ID code. For Arm Limited, this field 0xB.
func (r *registerPidr1Type) GetDes0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr1FieldDes0Mask) >> RegisterPidr1FieldDes0Shift)
}

// SetDes0 Designer, least significant nibble of JEP 106 ID code. For Arm Limited, this field 0xB.
func (r *registerPidr1Type) SetDes0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr1FieldDes0Mask)|(uint32(value)<<RegisterPidr1FieldDes0Shift))
}

// registerPidr2Type Provides information to identify the PMU.
type registerPidr2Type uint32

const (
	RegisterPidr2FieldDes1Shift = 0
	RegisterPidr2FieldDes1Mask  = 0x7
)

// GetDes1 Designer, most significant bits of JEP 106 ID code. For Arm Limited, this field reas as 0b011.
func (r *registerPidr2Type) GetDes1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr2FieldDes1Mask) >> RegisterPidr2FieldDes1Shift)
}

// SetDes1 Designer, most significant bits of JEP 106 ID code. For Arm Limited, this field reas as 0b011.
func (r *registerPidr2Type) SetDes1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr2FieldDes1Mask)|(uint32(value)<<RegisterPidr2FieldDes1Shift))
}

const (
	RegisterPidr2FieldJedecShift = 3
	RegisterPidr2FieldJedecMask  = 0x8
)

// GetJedec Indicates a JEP10 identity code. This is RAO.
func (r *registerPidr2Type) GetJedec() bool {
	return (volatile.LoadUint32((*uint32)(r)) & RegisterPidr2FieldJedecMask) != 0
}

// SetJedec Indicates a JEP10 identity code. This is RAO.
func (r *registerPidr2Type) SetJedec(value bool) {
	if value {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))|RegisterPidr2FieldJedecMask)
	} else {
		volatile.StoreUint32((*uint32)(r), volatile.LoadUint32((*uint32)(r))&^RegisterPidr2FieldJedecMask)
	}
}

const (
	RegisterPidr2FieldRevisionShift = 4
	RegisterPidr2FieldRevisionMask  = 0xf0
)

// GetRevision Part major revision. Parts can also use this field to extend Part number to 16-bits. This field reads as 0x0000.
func (r *registerPidr2Type) GetRevision() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr2FieldRevisionMask) >> RegisterPidr2FieldRevisionShift)
}

// SetRevision Part major revision. Parts can also use this field to extend Part number to 16-bits. This field reads as 0x0000.
func (r *registerPidr2Type) SetRevision(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr2FieldRevisionMask)|(uint32(value)<<RegisterPidr2FieldRevisionShift))
}

// registerPidr3Type Provides information to identify the PMU
type registerPidr3Type uint32

const (
	RegisterPidr3FieldCmodShift = 0
	RegisterPidr3FieldCmodMask  = 0x7
)

// GetCmod Customer modified. This field is 0x0000.
func (r *registerPidr3Type) GetCmod() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr3FieldCmodMask) >> RegisterPidr3FieldCmodShift)
}

// SetCmod Customer modified. This field is 0x0000.
func (r *registerPidr3Type) SetCmod(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr3FieldCmodMask)|(uint32(value)<<RegisterPidr3FieldCmodShift))
}

const (
	RegisterPidr3FieldRevandShift = 4
	RegisterPidr3FieldRevandMask  = 0xf0
)

// GetRevand Part minor revision. This field is 0x0000.
func (r *registerPidr3Type) GetRevand() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterPidr3FieldRevandMask) >> RegisterPidr3FieldRevandShift)
}

// SetRevand Part minor revision. This field is 0x0000.
func (r *registerPidr3Type) SetRevand(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterPidr3FieldRevandMask)|(uint32(value)<<RegisterPidr3FieldRevandShift))
}

// registerCidr0Type Provides information to identify a PMU component
type registerCidr0Type uint32

const (
	RegisterCidr0FieldPrmbl0Shift = 0
	RegisterCidr0FieldPrmbl0Mask  = 0xff
)

// GetPrmbl0 Preamble. This field reads 0x0D.
func (r *registerCidr0Type) GetPrmbl0() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCidr0FieldPrmbl0Mask) >> RegisterCidr0FieldPrmbl0Shift)
}

// SetPrmbl0 Preamble. This field reads 0x0D.
func (r *registerCidr0Type) SetPrmbl0(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCidr0FieldPrmbl0Mask)|(uint32(value)<<RegisterCidr0FieldPrmbl0Shift))
}

// registerCidr1Type Provides information to identify a PMU component
type registerCidr1Type uint32

const (
	RegisterCidr1FieldPrmbl1Shift = 0
	RegisterCidr1FieldPrmbl1Mask  = 0xf
)

// GetPrmbl1 Preamble. This field reads 0x0.
func (r *registerCidr1Type) GetPrmbl1() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCidr1FieldPrmbl1Mask) >> RegisterCidr1FieldPrmbl1Shift)
}

// SetPrmbl1 Preamble. This field reads 0x0.
func (r *registerCidr1Type) SetPrmbl1(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCidr1FieldPrmbl1Mask)|(uint32(value)<<RegisterCidr1FieldPrmbl1Shift))
}

const (
	RegisterCidr1FieldClassShift = 4
	RegisterCidr1FieldClassMask  = 0xf0
)

// GetClass Component class. This field reads 0x0.
func (r *registerCidr1Type) GetClass() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCidr1FieldClassMask) >> RegisterCidr1FieldClassShift)
}

// SetClass Component class. This field reads 0x0.
func (r *registerCidr1Type) SetClass(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCidr1FieldClassMask)|(uint32(value)<<RegisterCidr1FieldClassShift))
}

// registerCidr2Type Provides information to identify a PMU component
type registerCidr2Type uint32

const (
	RegisterCidr2FieldPrmbl2Shift = 0
	RegisterCidr2FieldPrmbl2Mask  = 0xf
)

// GetPrmbl2 Preamble. This field reads 0x05.
func (r *registerCidr2Type) GetPrmbl2() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCidr2FieldPrmbl2Mask) >> RegisterCidr2FieldPrmbl2Shift)
}

// SetPrmbl2 Preamble. This field reads 0x05.
func (r *registerCidr2Type) SetPrmbl2(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCidr2FieldPrmbl2Mask)|(uint32(value)<<RegisterCidr2FieldPrmbl2Shift))
}

// registerCidr3Type Provides information to identify a PMU component
type registerCidr3Type uint32

const (
	RegisterCidr3FieldPrmbl3Shift = 0
	RegisterCidr3FieldPrmbl3Mask  = 0xf
)

// GetPrmbl3 Preamble. This field reads 0xB1.
func (r *registerCidr3Type) GetPrmbl3() uint8 {
	return uint8((volatile.LoadUint32((*uint32)(r)) & RegisterCidr3FieldPrmbl3Mask) >> RegisterCidr3FieldPrmbl3Shift)
}

// SetPrmbl3 Preamble. This field reads 0xB1.
func (r *registerCidr3Type) SetPrmbl3(value uint8) {
	volatile.StoreUint32((*uint32)(r), (volatile.LoadUint32((*uint32)(r))&^RegisterCidr3FieldPrmbl3Mask)|(uint32(value)<<RegisterCidr3FieldPrmbl3Shift))
}
