#ifndef _PERIPHERALS_I2C_TD
#define _PERIPHERALS_I2C_TD

include "base.td"

def I2CPeripheral : PeripheralType<"I2C", "I2C"> {
  let accessWidth = 32;
  let registers = [
    Register<"CR1", 0x0, 32, [
      Field<"PE", 0, 1, ReadWrite, "Peripheral enable Note: When PE=0, the I2C SCL and SDA lines are released. Internal state machines and status bits are put back to their reset value. When cleared, PE must be kept low for at least 3 APB clock cycles.">,
      Field<"TXIE", 1, 1, ReadWrite, "TX Interrupt enable">,
      Field<"RXIE", 2, 1, ReadWrite, "RX Interrupt enable">,
      Field<"ADDRIE", 3, 1, ReadWrite, "Address match Interrupt enable (slave only)">,
      Field<"NACKIE", 4, 1, ReadWrite, "Not acknowledge received Interrupt enable">,
      Field<"STOPIE", 5, 1, ReadWrite, "STOP detection Interrupt enable">,
      Field<"TCIE", 6, 1, ReadWrite, "Transfer Complete interrupt enable Note: Any of these events will generate an interrupt: Transfer Complete (TC) Transfer Complete Reload (TCR)">,
      Field<"ERRIE", 7, 1, ReadWrite, "Error interrupts enable Note: Any of these errors generate an interrupt: Arbitration Loss (ARLO) Bus Error detection (BERR) Overrun/Underrun (OVR) Timeout detection (TIMEOUT) PEC error detection (PECERR) Alert pin event detection (ALERT)">,
      Field<"DNF", 8, 4, ReadWrite, "Digital noise filter These bits are used to configure the digital noise filter on SDA and SCL input. The digital filter will filter spikes with a length of up to DNF[3:0] * tI2CCLK ... Note: If the analog filter is also enabled, the digital filter is added to the analog filter. This filter can only be programmed when the I2C is disabled (PE = 0).">,
      Field<"ANFOFF", 12, 1, ReadWrite, "Analog noise filter OFF Note: This bit can only be programmed when the I2C is disabled (PE = 0).">,
      Field<"TXDMAEN", 14, 1, ReadWrite, "DMA transmission requests enable">,
      Field<"RXDMAEN", 15, 1, ReadWrite, "DMA reception requests enable">,
      Field<"SBC", 16, 1, ReadWrite, "Slave byte control This bit is used to enable hardware byte control in slave mode.">,
      Field<"NOSTRETCH", 17, 1, ReadWrite, "Clock stretching disable This bit is used to disable clock stretching in slave mode. It must be kept cleared in master mode. Note: This bit can only be programmed when the I2C is disabled (PE = 0).">,
      Field<"WUPEN", 18, 1, ReadWrite, "Wakeup from Stop mode enable Note: If the Wakeup from Stop mode feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation. Note: WUPEN can be set only when DNF = 0000">,
      Field<"GCEN", 19, 1, ReadWrite, "General call enable">,
      Field<"SMBHEN", 20, 1, ReadWrite, "SMBus Host address enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.">,
      Field<"SMBDEN", 21, 1, ReadWrite, "SMBus Device Default address enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.">,
      Field<"ALERTEN", 22, 1, ReadWrite, "SMBus alert enable Device mode (SMBHEN=0): Host mode (SMBHEN=1): Note: When ALERTEN=0, the SMBA pin can be used as a standard GPIO. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.">,
      Field<"PECEN", 23, 1, ReadWrite, "PEC enable Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.">,
    ], "Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.">,
    Register<"CR2", 0x4, 32, [
      Field<"SADD0", 0, 1, ReadWrite, "Slave address bit 0 (master mode) In 7-bit addressing mode (ADD10 = 0): This bit is dont care In 10-bit addressing mode (ADD10 = 1): This bit should be written with bit 0 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.">,
      Field<"SADD1", 1, 1, ReadWrite, "Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.">,
      Field<"SADD2", 2, 1, ReadWrite, "Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.">,
      Field<"SADD3", 3, 1, ReadWrite, "Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.">,
      Field<"SADD4", 4, 1, ReadWrite, "Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.">,
      Field<"SADD5", 5, 1, ReadWrite, "Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.">,
      Field<"SADD6", 6, 1, ReadWrite, "Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.">,
      Field<"SADD7", 7, 1, ReadWrite, "Slave address bit 7:1 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits should be written with the 7-bit slave address to be sent In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 7:1 of the slave address to be sent. Note: Changing these bits when the START bit is set is not allowed.">,
      Field<"SADD8", 8, 1, ReadWrite, "Slave address bit 9:8 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits are dont care In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 9:8 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.">,
      Field<"SADD9", 9, 1, ReadWrite, "Slave address bit 9:8 (master mode) In 7-bit addressing mode (ADD10 = 0): These bits are dont care In 10-bit addressing mode (ADD10 = 1): These bits should be written with bits 9:8 of the slave address to be sent Note: Changing these bits when the START bit is set is not allowed.">,
      Field<"RD_WRN", 10, 1, ReadWrite, "Transfer direction (master mode) Note: Changing this bit when the START bit is set is not allowed.">,
      Field<"ADD10", 11, 1, ReadWrite, "10-bit addressing mode (master mode) Note: Changing this bit when the START bit is set is not allowed.">,
      Field<"HEAD10R", 12, 1, ReadWrite, "10-bit address header only read direction (master receiver mode) Note: Changing this bit when the START bit is set is not allowed.">,
      Field<"START", 13, 1, ReadWrite, "Start generation This bit is set by software, and cleared by hardware after the Start followed by the address sequence is sent, by an arbitration loss, by a timeout error detection, or when PE = 0. It can also be cleared by software by writing 1 to the ADDRCF bit in the I2C_ICR register. If the I2C is already in master mode with AUTOEND = 0, setting this bit generates a Repeated Start condition when RELOAD=0, after the end of the NBYTES transfer. Otherwise setting this bit will generate a START condition once the bus is free. Note: Writing 0 to this bit has no effect. The START bit can be set even if the bus is BUSY or I2C is in slave mode. This bit has no effect when RELOAD is set.">,
      Field<"STOP", 14, 1, ReadWrite, "Stop generation (master mode) The bit is set by software, cleared by hardware when a Stop condition is detected, or when PE = 0. In Master Mode: Note: Writing 0 to this bit has no effect.">,
      Field<"NACK", 15, 1, ReadWrite, "NACK generation (slave mode) The bit is set by software, cleared by hardware when the NACK is sent, or when a STOP condition or an Address matched is received, or when PE=0. Note: Writing 0 to this bit has no effect. This bit is used in slave mode only: in master receiver mode, NACK is automatically generated after last byte preceding STOP or RESTART condition, whatever the NACK bit value. When an overrun occurs in slave receiver NOSTRETCH mode, a NACK is automatically generated whatever the NACK bit value. When hardware PEC checking is enabled (PECBYTE=1), the PEC acknowledge value does not depend on the NACK value.">,
      Field<"NBYTES", 16, 8, ReadWrite, "Number of bytes The number of bytes to be transmitted/received is programmed there. This field is dont care in slave mode with SBC=0. Note: Changing these bits when the START bit is set is not allowed.">,
      Field<"RELOAD", 24, 1, ReadWrite, "NBYTES reload mode This bit is set and cleared by software.">,
      Field<"AUTOEND", 25, 1, ReadWrite, "Automatic end mode (master mode) This bit is set and cleared by software. Note: This bit has no effect in slave mode or when the RELOAD bit is set.">,
      Field<"PECBYTE", 26, 1, ReadWrite, "Packet error checking byte This bit is set by software, and cleared by hardware when the PEC is transferred, or when a STOP condition or an Address matched is received, also when PE=0. Note: Writing 0 to this bit has no effect. This bit has no effect when RELOAD is set. This bit has no effect is slave mode when SBC=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.">,
    ], "Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.">,
    Register<"OAR1", 0x8, 32, [
      Field<"OA1", 0, 10, ReadWrite, "Interface address 7-bit addressing mode: dont care 10-bit addressing mode: bits 9:8 of address Note: These bits can be written only when OA1EN=0. OA1[7:1]: Interface address Bits 7:1 of address Note: These bits can be written only when OA1EN=0. OA1[0]: Interface address 7-bit addressing mode: dont care 10-bit addressing mode: bit 0 of address Note: This bit can be written only when OA1EN=0.">,
      Field<"OA1MODE", 10, 1, ReadWrite, "Own Address 1 10-bit mode Note: This bit can be written only when OA1EN=0.">,
      Field<"OA1EN", 15, 1, ReadWrite, "Own Address 1 enable">,
    ], "Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.">,
    Register<"OAR2", 0xc, 32, [
      Field<"OA2", 1, 7, ReadWrite, "Interface address bits 7:1 of address Note: These bits can be written only when OA2EN=0.">,
      Field<"OA2MSK", 8, 3, ReadWrite, "Own Address 2 masks Note: These bits can be written only when OA2EN=0. As soon as OA2MSK is not equal to 0, the reserved I2C addresses (0b0000xxx and 0b1111xxx) are not acknowledged even if the comparison matches.">,
      Field<"OA2EN", 15, 1, ReadWrite, "Own Address 2 enable">,
    ], "Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.">,
    Register<"TIMINGR", 0x10, 32, [
      Field<"SCLL", 0, 8, ReadWrite, "SCL low period (master mode) This field is used to generate the SCL low period in master mode. tSCLL = (SCLL+1) x tPRESC Note: SCLL is also used to generate tBUF and tSU:STA timings.">,
      Field<"SCLH", 8, 8, ReadWrite, "SCL high period (master mode) This field is used to generate the SCL high period in master mode. tSCLH = (SCLH+1) x tPRESC Note: SCLH is also used to generate tSU:STO and tHD:STA timing.">,
      Field<"SDADEL", 16, 4, ReadWrite, "Data hold time This field is used to generate the delay tSDADEL between SCL falling edge and SDA edge. In master mode and in slave mode with NOSTRETCH = 0, the SCL line is stretched low during tSDADEL. tSDADEL= SDADEL x tPRESC Note: SDADEL is used to generate tHD:DAT timing.">,
      Field<"SCLDEL", 20, 4, ReadWrite, "Data setup time This field is used to generate a delay tSCLDEL between SDA edge and SCL rising edge. In master mode and in slave mode with NOSTRETCH = 0, the SCL line is stretched low during tSCLDEL. tSCLDEL = (SCLDEL+1) x tPRESC Note: tSCLDEL is used to generate tSU:DAT timing.">,
      Field<"PRESC", 28, 4, ReadWrite, "Timing prescaler This field is used to prescale I2CCLK in order to generate the clock period tPRESC used for data setup and hold counters (refer to I2C timings on page9) and for SCL high and low level counters (refer to I2C master initialization on page24). tPRESC = (PRESC+1) x tI2CCLK">,
    ], "Access: No wait states">,
    Register<"TIMEOUTR", 0x14, 32, [
      Field<"TIMEOUTA", 0, 12, ReadWrite, "Bus Timeout A This field is used to configure: The SCL low timeout condition tTIMEOUT when TIDLE=0 tTIMEOUT= (TIMEOUTA+1) x 2048 x tI2CCLK The bus idle condition (both SCL and SDA high) when TIDLE=1 tIDLE= (TIMEOUTA+1) x 4 x tI2CCLK Note: These bits can be written only when TIMOUTEN=0.">,
      Field<"TIDLE", 12, 1, ReadWrite, "Idle clock timeout detection Note: This bit can be written only when TIMOUTEN=0.">,
      Field<"TIMOUTEN", 15, 1, ReadWrite, "Clock timeout enable">,
      Field<"TIMEOUTB", 16, 12, ReadWrite, "Bus timeout B This field is used to configure the cumulative clock extension timeout: In master mode, the master cumulative clock low extend time (tLOW:MEXT) is detected In slave mode, the slave cumulative clock low extend time (tLOW:SEXT) is detected tLOW:EXT= (TIMEOUTB+1) x 2048 x tI2CCLK Note: These bits can be written only when TEXTEN=0.">,
      Field<"TEXTEN", 31, 1, ReadWrite, "Extended clock timeout enable">,
    ], "Access: No wait states, except if a write access occurs while a write access to this register is ongoing. In this case, wait states are inserted in the second write access until the previous one is completed. The latency of the second write access can be up to 2 x PCLK1 + 6 x I2CCLK.">,
    Register<"ISR", 0x18, 32, [
      Field<"TXE", 0, 1, ReadWrite, "Transmit data register empty (transmitters) This bit is set by hardware when the I2C_TXDR register is empty. It is cleared when the next data to be sent is written in the I2C_TXDR register. This bit can be written to 1 by software in order to flush the transmit data register I2C_TXDR. Note: This bit is set by hardware when PE=0.">,
      Field<"TXIS", 1, 1, ReadWrite, "Transmit interrupt status (transmitters) This bit is set by hardware when the I2C_TXDR register is empty and the data to be transmitted must be written in the I2C_TXDR register. It is cleared when the next data to be sent is written in the I2C_TXDR register. This bit can be written to 1 by software when NOSTRETCH=1 only, in order to generate a TXIS event (interrupt if TXIE=1 or DMA request if TXDMAEN=1). Note: This bit is cleared by hardware when PE=0.">,
      Field<"RXNE", 2, 1, Read, "Receive data register not empty (receivers) This bit is set by hardware when the received data is copied into the I2C_RXDR register, and is ready to be read. It is cleared when I2C_RXDR is read. Note: This bit is cleared by hardware when PE=0.">,
      Field<"ADDR", 3, 1, Read, "Address matched (slave mode) This bit is set by hardware as soon as the received slave address matched with one of the enabled slave addresses. It is cleared by software by setting ADDRCF bit. Note: This bit is cleared by hardware when PE=0.">,
      Field<"NACKF", 4, 1, Read, "Not Acknowledge received flag This flag is set by hardware when a NACK is received after a byte transmission. It is cleared by software by setting the NACKCF bit. Note: This bit is cleared by hardware when PE=0.">,
      Field<"STOPF", 5, 1, Read, "Stop detection flag This flag is set by hardware when a Stop condition is detected on the bus and the peripheral is involved in this transfer: either as a master, provided that the STOP condition is generated by the peripheral. or as a slave, provided that the peripheral has been addressed previously during this transfer. It is cleared by software by setting the STOPCF bit. Note: This bit is cleared by hardware when PE=0.">,
      Field<"TC", 6, 1, Read, "Transfer Complete (master mode) This flag is set by hardware when RELOAD=0, AUTOEND=0 and NBYTES data have been transferred. It is cleared by software when START bit or STOP bit is set. Note: This bit is cleared by hardware when PE=0.">,
      Field<"TCR", 7, 1, Read, "Transfer Complete Reload This flag is set by hardware when RELOAD=1 and NBYTES data have been transferred. It is cleared by software when NBYTES is written to a non-zero value. Note: This bit is cleared by hardware when PE=0. This flag is only for master mode, or for slave mode when the SBC bit is set.">,
      Field<"BERR", 8, 1, Read, "Bus error This flag is set by hardware when a misplaced Start or Stop condition is detected whereas the peripheral is involved in the transfer. The flag is not set during the address phase in slave mode. It is cleared by software by setting BERRCF bit. Note: This bit is cleared by hardware when PE=0.">,
      Field<"ARLO", 9, 1, Read, "Arbitration lost This flag is set by hardware in case of arbitration loss. It is cleared by software by setting the ARLOCF bit. Note: This bit is cleared by hardware when PE=0.">,
      Field<"OVR", 10, 1, Read, "Overrun/Underrun (slave mode) This flag is set by hardware in slave mode with NOSTRETCH=1, when an overrun/underrun error occurs. It is cleared by software by setting the OVRCF bit. Note: This bit is cleared by hardware when PE=0.">,
      Field<"PECERR", 11, 1, Read, "PEC Error in reception This flag is set by hardware when the received PEC does not match with the PEC register content. A NACK is automatically sent after the wrong PEC reception. It is cleared by software by setting the PECCF bit. Note: This bit is cleared by hardware when PE=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.">,
      Field<"TIMEOUT", 12, 1, Read, "Timeout or tLOW detection flag This flag is set by hardware when a timeout or extended clock timeout occurred. It is cleared by software by setting the TIMEOUTCF bit. Note: This bit is cleared by hardware when PE=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.">,
      Field<"ALERT", 13, 1, Read, "SMBus alert This flag is set by hardware when SMBHEN=1 (SMBus host configuration), ALERTEN=1 and a SMBALERT event (falling edge) is detected on SMBA pin. It is cleared by software by setting the ALERTCF bit. Note: This bit is cleared by hardware when PE=0. If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.">,
      Field<"BUSY", 15, 1, Read, "Bus busy This flag indicates that a communication is in progress on the bus. It is set by hardware when a START condition is detected. It is cleared by hardware when a Stop condition is detected, or when PE=0.">,
      Field<"DIR", 16, 1, Read, "Transfer direction (Slave mode) This flag is updated when an address match event occurs (ADDR=1).">,
      Field<"ADDCODE", 17, 7, Read, "Address match code (Slave mode) These bits are updated with the received address when an address match event occurs (ADDR = 1). In the case of a 10-bit address, ADDCODE provides the 10-bit header followed by the 2 MSBs of the address.">,
    ], "Access: No wait states">,
    Register<"ICR", 0x1c, 32, [
      Field<"ADDRCF", 3, 1, ReadWrite, "Address matched flag clear Writing 1 to this bit clears the ADDR flag in the I2C_ISR register. Writing 1 to this bit also clears the START bit in the I2C_CR2 register.">,
      Field<"NACKCF", 4, 1, ReadWrite, "Not Acknowledge flag clear Writing 1 to this bit clears the ACKF flag in I2C_ISR register.">,
      Field<"STOPCF", 5, 1, ReadWrite, "Stop detection flag clear Writing 1 to this bit clears the STOPF flag in the I2C_ISR register.">,
      Field<"BERRCF", 8, 1, ReadWrite, "Bus error flag clear Writing 1 to this bit clears the BERRF flag in the I2C_ISR register.">,
      Field<"ARLOCF", 9, 1, ReadWrite, "Arbitration Lost flag clear Writing 1 to this bit clears the ARLO flag in the I2C_ISR register.">,
      Field<"OVRCF", 10, 1, ReadWrite, "Overrun/Underrun flag clear Writing 1 to this bit clears the OVR flag in the I2C_ISR register.">,
      Field<"PECCF", 11, 1, ReadWrite, "PEC Error flag clear Writing 1 to this bit clears the PECERR flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.">,
      Field<"TIMOUTCF", 12, 1, ReadWrite, "Timeout detection flag clear Writing 1 to this bit clears the TIMEOUT flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.">,
      Field<"ALERTCF", 13, 1, ReadWrite, "Alert flag clear Writing 1 to this bit clears the ALERT flag in the I2C_ISR register. Note: If the SMBus feature is not supported, this bit is reserved and forced by hardware to 0. Please refer to Section25.3: I2C implementation.">,
    ], "Access: No wait states">,
    Register<"PECR", 0x20, 32, [
      Field<"PEC", 0, 8, ReadWrite, "Packet error checking register This field contains the internal PEC when PECEN=1. The PEC is cleared by hardware when PE=0.">,
    ], "Access: No wait states">,
    Register<"RXDR", 0x24, 32, [
      Field<"RXDATA", 0, 8, ReadWrite, "8-bit receive data Data byte received from the I2C bus.">,
    ], "Access: No wait states">,
    Register<"TXDR", 0x28, 32, [
      Field<"TXDATA", 0, 8, ReadWrite, "8-bit transmit data Data byte to be transmitted to the I2C bus. Note: These bits can be written only when TXE=1.">,
    ], "Access: No wait states">,
  ];
}

class I2CInstance<string Name, int Base> : PeripheralInstance<Name, Base, I2CPeripheral>;
def I2C1 : I2CInstance<"I2C1", 0x40005400>;
def I2C2 : I2CInstance<"I2C2", 0x40005800>;
def I2C3 : I2CInstance<"I2C3", 0x40005c00>;
def I2C4 : I2CInstance<"I2C4", 0x58001c00>;

def I2CGroup : PeripheralGroup<[ I2C1, I2C2, I2C3, I2C4 ]>;

#endif // _PERIPHERALS_I2C1_TD
