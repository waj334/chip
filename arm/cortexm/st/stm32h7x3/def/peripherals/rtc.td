#ifndef _PERIPHERALS_RTC_TD
#define _PERIPHERALS_RTC_TD

include "base.td"

def RTCPeripheral : PeripheralType<"RTC", "RTC"> {
  let accessWidth = 32;
  let registers = [
    Register<"RTC_TR", 0x0, 32, [
      Field<"SU", 0, 4, ReadWrite, "Second units in BCD format">,
      Field<"ST", 4, 3, ReadWrite, "Second tens in BCD format">,
      Field<"MNU", 8, 4, ReadWrite, "Minute units in BCD format">,
      Field<"MNT", 12, 3, ReadWrite, "Minute tens in BCD format">,
      Field<"HU", 16, 4, ReadWrite, "Hour units in BCD format">,
      Field<"HT", 20, 2, ReadWrite, "Hour tens in BCD format">,
      Field<"PM", 22, 1, ReadWrite, "AM/PM notation">,
    ], "The RTC_TR is the calendar time shadow register. This register must be written in initialization mode only. Refer to Calendar initialization and configuration on page9 and Reading the calendar on page10.This register is write protected. The write access procedure is described in RTC register write protection on page9.">,
    Register<"RTC_DR", 0x4, 32, [
      Field<"DU", 0, 4, ReadWrite, "Date units in BCD format">,
      Field<"DT", 4, 2, ReadWrite, "Date tens in BCD format">,
      Field<"MU", 8, 4, ReadWrite, "Month units in BCD format">,
      Field<"MT", 12, 1, ReadWrite, "Month tens in BCD format">,
      Field<"WDU", 13, 3, ReadWrite, "Week day units">,
      Field<"YU", 16, 4, ReadWrite, "Year units in BCD format">,
      Field<"YT", 20, 4, ReadWrite, "Year tens in BCD format">,
    ], "The RTC_DR is the calendar date shadow register. This register must be written in initialization mode only. Refer to Calendar initialization and configuration on page9 and Reading the calendar on page10.This register is write protected. The write access procedure is described in RTC register write protection on page9.">,
    Register<"RTC_CR", 0x8, 32, [
      Field<"WUCKSEL", 0, 3, ReadWrite, "Wakeup clock selection">,
      Field<"TSEDGE", 3, 1, ReadWrite, "Time-stamp event active edge TSE must be reset when TSEDGE is changed to avoid unwanted TSF setting.">,
      Field<"REFCKON", 4, 1, ReadWrite, "RTC_REFIN reference clock detection enable (50 or 60Hz) Note: PREDIV_S must be 0x00FF.">,
      Field<"BYPSHAD", 5, 1, ReadWrite, "Bypass the shadow registers Note: If the frequency of the APB clock is less than seven times the frequency of RTCCLK, BYPSHAD must be set to 1.">,
      Field<"FMT", 6, 1, ReadWrite, "Hour format">,
      Field<"ALRAE", 8, 1, ReadWrite, "Alarm A enable">,
      Field<"ALRBE", 9, 1, ReadWrite, "Alarm B enable">,
      Field<"WUTE", 10, 1, ReadWrite, "Wakeup timer enable">,
      Field<"TSE", 11, 1, ReadWrite, "timestamp enable">,
      Field<"ALRAIE", 12, 1, ReadWrite, "Alarm A interrupt enable">,
      Field<"ALRBIE", 13, 1, ReadWrite, "Alarm B interrupt enable">,
      Field<"WUTIE", 14, 1, ReadWrite, "Wakeup timer interrupt enable">,
      Field<"TSIE", 15, 1, ReadWrite, "Time-stamp interrupt enable">,
      Field<"ADD1H", 16, 1, Write, "Add 1 hour (summer time change) When this bit is set outside initialization mode, 1 hour is added to the calendar time. This bit is always read as 0.">,
      Field<"SUB1H", 17, 1, Write, "Subtract 1 hour (winter time change) When this bit is set outside initialization mode, 1 hour is subtracted to the calendar time if the current hour is not 0. This bit is always read as 0. Setting this bit has no effect when current hour is 0.">,
      Field<"BKP", 18, 1, ReadWrite, "Backup This bit can be written by the user to memorize whether the daylight saving time change has been performed or not.">,
      Field<"COSEL", 19, 1, ReadWrite, "Calibration output selection When COE=1, this bit selects which signal is output on RTC_CALIB. These frequencies are valid for RTCCLK at 32.768 kHz and prescalers at their default values (PREDIV_A=127 and PREDIV_S=255). Refer to Section24.3.15: Calibration clock output">,
      Field<"POL", 20, 1, ReadWrite, "Output polarity This bit is used to configure the polarity of RTC_ALARM output">,
      Field<"OSEL", 21, 2, ReadWrite, "Output selection These bits are used to select the flag to be routed to RTC_ALARM output">,
      Field<"COE", 23, 1, ReadWrite, "Calibration output enable This bit enables the RTC_CALIB output">,
      Field<"ITSE", 24, 1, ReadWrite, "timestamp on internal event enable">,
    ], "RTC control register">,
    Register<"RTC_ISR", 0xc, 32, [
      Field<"ALRAWF", 0, 1, Read, "Alarm A write flag This bit is set by hardware when Alarm A values can be changed, after the ALRAE bit has been set to 0 in RTC_CR. It is cleared by hardware in initialization mode.">,
      Field<"ALRBWF", 1, 1, Read, "Alarm B write flag This bit is set by hardware when Alarm B values can be changed, after the ALRBE bit has been set to 0 in RTC_CR. It is cleared by hardware in initialization mode.">,
      Field<"WUTWF", 2, 1, Read, "Wakeup timer write flag This bit is set by hardware up to 2 RTCCLK cycles after the WUTE bit has been set to 0 in RTC_CR, and is cleared up to 2 RTCCLK cycles after the WUTE bit has been set to 1. The wakeup timer values can be changed when WUTE bit is cleared and WUTWF is set.">,
      Field<"SHPF", 3, 1, Read, "Shift operation pending This flag is set by hardware as soon as a shift operation is initiated by a write to the RTC_SHIFTR register. It is cleared by hardware when the corresponding shift operation has been executed. Writing to the SHPF bit has no effect.">,
      Field<"INITS", 4, 1, Read, "Initialization status flag This bit is set by hardware when the calendar year field is different from 0 (Backup domain reset state).">,
      Field<"RSF", 5, 1, ReadWrite, "Registers synchronization flag This bit is set by hardware each time the calendar registers are copied into the shadow registers (RTC_SSRx, RTC_TRx and RTC_DRx). This bit is cleared by hardware in initialization mode, while a shift operation is pending (SHPF=1), or when in bypass shadow register mode (BYPSHAD=1). This bit can also be cleared by software. It is cleared either by software or by hardware in initialization mode.">,
      Field<"INITF", 6, 1, Read, "Initialization flag When this bit is set to 1, the RTC is in initialization state, and the time, date and prescaler registers can be updated.">,
      Field<"INIT", 7, 1, ReadWrite, "Initialization mode">,
      Field<"ALRAF", 8, 1, ReadWrite, "Alarm A flag This flag is set by hardware when the time/date registers (RTC_TR and RTC_DR) match the Alarm A register (RTC_ALRMAR). This flag is cleared by software by writing 0.">,
      Field<"ALRBF", 9, 1, ReadWrite, "Alarm B flag This flag is set by hardware when the time/date registers (RTC_TR and RTC_DR) match the Alarm B register (RTC_ALRMBR). This flag is cleared by software by writing 0.">,
      Field<"WUTF", 10, 1, ReadWrite, "Wakeup timer flag This flag is set by hardware when the wakeup auto-reload counter reaches 0. This flag is cleared by software by writing 0. This flag must be cleared by software at least 1.5 RTCCLK periods before WUTF is set to 1 again.">,
      Field<"TSF", 11, 1, ReadWrite, "Time-stamp flag This flag is set by hardware when a time-stamp event occurs. This flag is cleared by software by writing 0.">,
      Field<"TSOVF", 12, 1, ReadWrite, "Time-stamp overflow flag This flag is set by hardware when a time-stamp event occurs while TSF is already set. This flag is cleared by software by writing 0. It is recommended to check and then clear TSOVF only after clearing the TSF bit. Otherwise, an overflow might not be noticed if a time-stamp event occurs immediately before the TSF bit is cleared.">,
      Field<"TAMP1F", 13, 1, ReadWrite, "RTC_TAMP1 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP1 input. It is cleared by software writing 0">,
      Field<"TAMP2F", 14, 1, ReadWrite, "RTC_TAMP2 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP2 input. It is cleared by software writing 0">,
      Field<"TAMP3F", 15, 1, ReadWrite, "RTC_TAMP3 detection flag This flag is set by hardware when a tamper detection event is detected on the RTC_TAMP3 input. It is cleared by software writing 0">,
      Field<"RECALPF", 16, 1, Read, "Recalibration pending Flag The RECALPF status flag is automatically set to 1 when software writes to the RTC_CALR register, indicating that the RTC_CALR register is blocked. When the new calibration settings are taken into account, this bit returns to 0. Refer to Re-calibration on-the-fly.">,
      Field<"ITSF", 17, 1, ReadWrite, "Internal tTime-stamp flag">,
    ], "This register is write protected (except for RTC_ISR[13:8] bits). The write access procedure is described in RTC register write protection on page9.">,
    Register<"RTC_PRER", 0x10, 32, [
      Field<"PREDIV_S", 0, 15, ReadWrite, "Synchronous prescaler factor This is the synchronous division factor: ck_spre frequency = ck_apre frequency/(PREDIV_S+1)">,
      Field<"PREDIV_A", 16, 7, ReadWrite, "Asynchronous prescaler factor This is the asynchronous division factor: ck_apre frequency = RTCCLK frequency/(PREDIV_A+1)">,
    ], "This register must be written in initialization mode only. The initialization must be performed in two separate write accesses. Refer to Calendar initialization and configuration on page9.This register is write protected. The write access procedure is described in RTC register write protection on page9.">,
    Register<"RTC_WUTR", 0x14, 32, [
      Field<"WUT", 0, 16, ReadWrite, "Wakeup auto-reload value bits When the wakeup timer is enabled (WUTE set to 1), the WUTF flag is set every (WUT[15:0] + 1) ck_wut cycles. The ck_wut period is selected through WUCKSEL[2:0] bits of the RTC_CR register When WUCKSEL[2] = 1, the wakeup timer becomes 17-bits and WUCKSEL[1] effectively becomes WUT[16] the most-significant bit to be reloaded into the timer. The first assertion of WUTF occurs (WUT+1) ck_wut cycles after WUTE is set. Setting WUT[15:0] to 0x0000 with WUCKSEL[2:0] =011 (RTCCLK/2) is forbidden.">,
    ], "This register can be written only when WUTWF is set to 1 in RTC_ISR.This register is write protected. The write access procedure is described in RTC register write protection on page9.">,
    Register<"RTC_ALRMAR", 0x1c, 32, [
      Field<"SU", 0, 4, ReadWrite, "Second units in BCD format.">,
      Field<"ST", 4, 3, ReadWrite, "Second tens in BCD format.">,
      Field<"MSK1", 7, 1, ReadWrite, "Alarm A seconds mask">,
      Field<"MNU", 8, 4, ReadWrite, "Minute units in BCD format.">,
      Field<"MNT", 12, 3, ReadWrite, "Minute tens in BCD format.">,
      Field<"MSK2", 15, 1, ReadWrite, "Alarm A minutes mask">,
      Field<"HU", 16, 4, ReadWrite, "Hour units in BCD format.">,
      Field<"HT", 20, 2, ReadWrite, "Hour tens in BCD format.">,
      Field<"PM", 22, 1, ReadWrite, "AM/PM notation">,
      Field<"MSK3", 23, 1, ReadWrite, "Alarm A hours mask">,
      Field<"DU", 24, 4, ReadWrite, "Date units or day in BCD format.">,
      Field<"DT", 28, 2, ReadWrite, "Date tens in BCD format.">,
      Field<"WDSEL", 30, 1, ReadWrite, "Week day selection">,
      Field<"MSK4", 31, 1, ReadWrite, "Alarm A date mask">,
    ], "This register can be written only when ALRAWF is set to 1 in RTC_ISR, or in initialization mode.This register is write protected. The write access procedure is described in RTC register write protection on page9.">,
    Register<"RTC_ALRMBR", 0x20, 32, [
      Field<"SU", 0, 4, ReadWrite, "Second units in BCD format">,
      Field<"ST", 4, 3, ReadWrite, "Second tens in BCD format">,
      Field<"MSK1", 7, 1, ReadWrite, "Alarm B seconds mask">,
      Field<"MNU", 8, 4, ReadWrite, "Minute units in BCD format">,
      Field<"MNT", 12, 3, ReadWrite, "Minute tens in BCD format">,
      Field<"MSK2", 15, 1, ReadWrite, "Alarm B minutes mask">,
      Field<"HU", 16, 4, ReadWrite, "Hour units in BCD format">,
      Field<"HT", 20, 2, ReadWrite, "Hour tens in BCD format">,
      Field<"PM", 22, 1, ReadWrite, "AM/PM notation">,
      Field<"MSK3", 23, 1, ReadWrite, "Alarm B hours mask">,
      Field<"DU", 24, 4, ReadWrite, "Date units or day in BCD format">,
      Field<"DT", 28, 2, ReadWrite, "Date tens in BCD format">,
      Field<"WDSEL", 30, 1, ReadWrite, "Week day selection">,
      Field<"MSK4", 31, 1, ReadWrite, "Alarm B date mask">,
    ], "This register can be written only when ALRBWF is set to 1 in RTC_ISR, or in initialization mode.This register is write protected. The write access procedure is described in RTC register write protection on page9.">,
    Register<"RTC_WPR", 0x24, 32, [
      Field<"KEY", 0, 8, ReadWrite, "Write protection key This byte is written by software. Reading this byte always returns 0x00. Refer to RTC register write protection for a description of how to unlock RTC register write protection.">,
    ], "RTC write protection register">,
    Register<"RTC_SSR", 0x28, 32, [
      Field<"SS", 0, 16, ReadWrite, "Sub second value SS[15:0] is the value in the synchronous prescaler counter. The fraction of a second is given by the formula below: Second fraction = (PREDIV_S - SS) / (PREDIV_S + 1) Note: SS can be larger than PREDIV_S only after a shift operation. In that case, the correct time/date is one second less than as indicated by RTC_TR/RTC_DR.">,
    ], "RTC sub second register">,
    Register<"RTC_SHIFTR", 0x2c, 32, [
      Field<"SUBFS", 0, 15, ReadWrite, "Subtract a fraction of a second These bits are write only and is always read as zero. Writing to this bit has no effect when a shift operation is pending (when SHPF=1, in RTC_ISR). The value which is written to SUBFS is added to the synchronous prescaler counter. Since this counter counts down, this operation effectively subtracts from (delays) the clock by: Delay (seconds) = SUBFS / (PREDIV_S + 1) A fraction of a second can effectively be added to the clock (advancing the clock) when the ADD1S function is used in conjunction with SUBFS, effectively advancing the clock by: Advance (seconds) = (1 - (SUBFS / (PREDIV_S + 1))). Note: Writing to SUBFS causes RSF to be cleared. Software can then wait until RSF=1 to be sure that the shadow registers have been updated with the shifted time.">,
      Field<"ADD1S", 31, 1, ReadWrite, "Add one second This bit is write only and is always read as zero. Writing to this bit has no effect when a shift operation is pending (when SHPF=1, in RTC_ISR). This function is intended to be used with SUBFS (see description below) in order to effectively add a fraction of a second to the clock in an atomic operation.">,
    ], "This register is write protected. The write access procedure is described in RTC register write protection on page9.">,
    Register<"RTC_TSTR", 0x30, 32, [
      Field<"SU", 0, 4, ReadWrite, "Second units in BCD format.">,
      Field<"ST", 4, 3, ReadWrite, "Second tens in BCD format.">,
      Field<"MNU", 8, 4, ReadWrite, "Minute units in BCD format.">,
      Field<"MNT", 12, 3, ReadWrite, "Minute tens in BCD format.">,
      Field<"HU", 16, 4, ReadWrite, "Hour units in BCD format.">,
      Field<"HT", 20, 2, ReadWrite, "Hour tens in BCD format.">,
      Field<"PM", 22, 1, ReadWrite, "AM/PM notation">,
    ], "The content of this register is valid only when TSF is set to 1 in RTC_ISR. It is cleared when TSF bit is reset.">,
    Register<"RTC_TSDR", 0x34, 32, [
      Field<"DU", 0, 4, ReadWrite, "Date units in BCD format">,
      Field<"DT", 4, 2, ReadWrite, "Date tens in BCD format">,
      Field<"MU", 8, 4, ReadWrite, "Month units in BCD format">,
      Field<"MT", 12, 1, ReadWrite, "Month tens in BCD format">,
      Field<"WDU", 13, 3, ReadWrite, "Week day units">,
    ], "The content of this register is valid only when TSF is set to 1 in RTC_ISR. It is cleared when TSF bit is reset.">,
    Register<"RTC_TSSSR", 0x38, 32, [
      Field<"SS", 0, 16, ReadWrite, "Sub second value SS[15:0] is the value of the synchronous prescaler counter when the timestamp event occurred.">,
    ], "The content of this register is valid only when RTC_ISR/TSF is set. It is cleared when the RTC_ISR/TSF bit is reset.">,
    Register<"RTC_CALR", 0x3c, 32, [
      Field<"CALM", 0, 9, ReadWrite, "Calibration minus The frequency of the calendar is reduced by masking CALM out of 220 RTCCLK pulses (32 seconds if the input frequency is 32768 Hz). This decreases the frequency of the calendar with a resolution of 0.9537 ppm. To increase the frequency of the calendar, this feature should be used in conjunction with CALP. See Section24.3.12: RTC smooth digital calibration on page13.">,
      Field<"CALW16", 13, 1, ReadWrite, "Use a 16-second calibration cycle period When CALW16 is set to 1, the 16-second calibration cycle period is selected.This bit must not be set to 1 if CALW8=1. Note: CALM[0] is stuck at 0 when CALW16= 1. Refer to Section24.3.12: RTC smooth digital calibration.">,
      Field<"CALW8", 14, 1, ReadWrite, "Use an 8-second calibration cycle period When CALW8 is set to 1, the 8-second calibration cycle period is selected. Note: CALM[1:0] are stuck at 00; when CALW8= 1. Refer to Section24.3.12: RTC smooth digital calibration.">,
      Field<"CALP", 15, 1, ReadWrite, "Increase frequency of RTC by 488.5 ppm This feature is intended to be used in conjunction with CALM, which lowers the frequency of the calendar with a fine resolution. if the input frequency is 32768 Hz, the number of RTCCLK pulses added during a 32-second window is calculated as follows: (512 * CALP) - CALM. Refer to Section24.3.12: RTC smooth digital calibration.">,
    ], "This register is write protected. The write access procedure is described in RTC register write protection on page9.">,
    Register<"RTC_TAMPCR", 0x40, 32, [
      Field<"TAMP1E", 0, 1, ReadWrite, "RTC_TAMP1 input detection enable">,
      Field<"TAMP1TRG", 1, 1, ReadWrite, "Active level for RTC_TAMP1 input If TAMPFLT != 00 if TAMPFLT = 00:">,
      Field<"TAMPIE", 2, 1, ReadWrite, "Tamper interrupt enable">,
      Field<"TAMP2E", 3, 1, ReadWrite, "RTC_TAMP2 input detection enable">,
      Field<"TAMP2TRG", 4, 1, ReadWrite, "Active level for RTC_TAMP2 input if TAMPFLT != 00: if TAMPFLT = 00:">,
      Field<"TAMP3E", 5, 1, ReadWrite, "RTC_TAMP3 detection enable">,
      Field<"TAMP3TRG", 6, 1, ReadWrite, "Active level for RTC_TAMP3 input if TAMPFLT != 00: if TAMPFLT = 00:">,
      Field<"TAMPTS", 7, 1, ReadWrite, "Activate timestamp on tamper detection event TAMPTS is valid even if TSE=0 in the RTC_CR register.">,
      Field<"TAMPFREQ", 8, 3, ReadWrite, "Tamper sampling frequency Determines the frequency at which each of the RTC_TAMPx inputs are sampled.">,
      Field<"TAMPFLT", 11, 2, ReadWrite, "RTC_TAMPx filter count These bits determines the number of consecutive samples at the specified level (TAMP*TRG) needed to activate a Tamper event. TAMPFLT is valid for each of the RTC_TAMPx inputs.">,
      Field<"TAMPPRCH", 13, 2, ReadWrite, "RTC_TAMPx precharge duration These bit determines the duration of time during which the pull-up/is activated before each sample. TAMPPRCH is valid for each of the RTC_TAMPx inputs.">,
      Field<"TAMPPUDIS", 15, 1, ReadWrite, "RTC_TAMPx pull-up disable This bit determines if each of the RTC_TAMPx pins are pre-charged before each sample.">,
      Field<"TAMP1IE", 16, 1, ReadWrite, "Tamper 1 interrupt enable">,
      Field<"TAMP1NOERASE", 17, 1, ReadWrite, "Tamper 1 no erase">,
      Field<"TAMP1MF", 18, 1, ReadWrite, "Tamper 1 mask flag">,
      Field<"TAMP2IE", 19, 1, ReadWrite, "Tamper 2 interrupt enable">,
      Field<"TAMP2NOERASE", 20, 1, ReadWrite, "Tamper 2 no erase">,
      Field<"TAMP2MF", 21, 1, ReadWrite, "Tamper 2 mask flag">,
      Field<"TAMP3IE", 22, 1, ReadWrite, "Tamper 3 interrupt enable">,
      Field<"TAMP3NOERASE", 23, 1, ReadWrite, "Tamper 3 no erase">,
      Field<"TAMP3MF", 24, 1, ReadWrite, "Tamper 3 mask flag">,
    ], "RTC tamper and alternate function configuration register">,
    Register<"RTC_ALRMASSR", 0x44, 32, [
      Field<"SS", 0, 15, ReadWrite, "Sub seconds value This value is compared with the contents of the synchronous prescaler counter to determine if Alarm A is to be activated. Only bits 0 up MASKSS-1 are compared.">,
      Field<"MASKSS", 24, 4, ReadWrite, "Mask the most-significant bits starting at this bit ... The overflow bits of the synchronous counter (bits 15) is never compared. This bit can be different from 0 only after a shift operation.">,
    ], "This register can be written only when ALRAE is reset in RTC_CR register, or in initialization mode.This register is write protected. The write access procedure is described in RTC register write protection on page9">,
    Register<"RTC_ALRMBSSR", 0x48, 32, [
      Field<"SS", 0, 15, ReadWrite, "Sub seconds value This value is compared with the contents of the synchronous prescaler counter to determine if Alarm B is to be activated. Only bits 0 up to MASKSS-1 are compared.">,
      Field<"MASKSS", 24, 4, ReadWrite, "Mask the most-significant bits starting at this bit ... The overflow bits of the synchronous counter (bits 15) is never compared. This bit can be different from 0 only after a shift operation.">,
    ], "This register can be written only when ALRBE is reset in RTC_CR register, or in initialization mode.This register is write protected.The write access procedure is described in Section: RTC register write protection.">,
    Register<"RTC_BKP0R", 0x50, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP1R", 0x54, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP2R", 0x58, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP3R", 0x5c, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP4R", 0x60, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP5R", 0x64, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP6R", 0x68, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP7R", 0x6c, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP8R", 0x70, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP9R", 0x74, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP10R", 0x78, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP11R", 0x7c, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP12R", 0x80, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP13R", 0x84, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP14R", 0x88, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP15R", 0x8c, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_OR", 0x4c, 32, [
      Field<"RTC_ALARM_TYPE", 0, 1, ReadWrite, "RTC_ALARM output type on PC13">,
      Field<"RTC_OUT_RMP", 1, 1, ReadWrite, "RTC_OUT remap">,
    ], "RTC option register">,
    Register<"RTC_BKP16R", 0x90, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP17R", 0x94, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP18R", 0x98, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP19R", 0x9c, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP20R", 0xa0, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP21R", 0xa4, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP22R", 0xa8, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP23R", 0xac, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP24R", 0xb0, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP25R", 0xb4, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP26R", 0xb8, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP27R", 0xbc, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP28R", 0xc0, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP29R", 0xc4, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP30R", 0xc8, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
    Register<"RTC_BKP31R", 0xcc, 32, [
      Field<"BKP", 0, 32, ReadWrite, "The application can write or read data to and from these registers. They are powered-on by VBAT when VDD is switched off, so that they are not reset by System reset, and their contents remain valid when the device operates in low-power mode. This register is reset on a tamper detection event, as long as TAMPxF=1. or when the Flash readout protection is disabled.">,
    ], "RTC backup registers">,
  ];
}

class RTCInstance<string Name, int Base> : PeripheralInstance<Name, Base, RTCPeripheral>;
def RTC : RTCInstance<"RTC", 0x58004000>;

#endif // _PERIPHERALS_RTC_TD
