#ifndef _PERIPHERALS_SDMMC1_TD
#define _PERIPHERALS_SDMMC1_TD

include "base.td"

def SDMMC1Peripheral : PeripheralType<"SDMMC1", "SDMMC1"> {
  let accessWidth = 32;
  let registers = [
    Register<"POWER", 0x0, 32, [
      Field<"PWRCTRL", 0, 2, ReadWrite, "SDMMC state control bits. These bits can only be written when the SDMMC is not in the power-on state (PWRCTRL?11). These bits are used to define the functional state of the SDMMC signals: Any further write will be ignored, PWRCTRL value will keep 11.">,
      Field<"VSWITCH", 2, 1, ReadWrite, "Voltage switch sequence start. This bit is used to start the timing critical section of the voltage switch sequence:">,
      Field<"VSWITCHEN", 3, 1, ReadWrite, "Voltage switch procedure enable. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). This bit is used to stop the SDMMC_CK after the voltage switch command response:">,
      Field<"DIRPOL", 4, 1, ReadWrite, "Data and command direction signals polarity selection. This bit can only be written when the SDMMC is in the power-off state (PWRCTRL = 00).">,
    ], "SDMMC power control register">,
    Register<"CLKCR", 0x4, 32, [
      Field<"CLKDIV", 0, 10, ReadWrite, "Clock divide factor This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). This field defines the divide factor between the input clock (SDMMCCLK) and the output clock (SDMMC_CK): SDMMC_CK frequency = SDMMCCLK / [2 * CLKDIV]. 0xx: etc.. xxx: etc..">,
      Field<"PWRSAV", 12, 1, ReadWrite, "Power saving configuration bit This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) For power saving, the SDMMC_CK clock output can be disabled when the bus is idle by setting PWRSAV:">,
      Field<"WIDBUS", 14, 2, ReadWrite, "Wide bus mode enable bit This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)">,
      Field<"NEGEDGE", 16, 1, ReadWrite, "SDMMC_CK dephasing selection bit for data and Command. This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). When clock division = 1 (CLKDIV = 0), this bit has no effect. Data and Command change on SDMMC_CK falling edge. When clock division &gt;1 (CLKDIV &gt; 0) &amp; DDR = 0: - SDMMC_CK edge occurs on SDMMCCLK rising edge. When clock division >1 (CLKDIV > 0) & DDR = 1: - Data changed on the SDMMCCLK falling edge succeeding a SDMMC_CK edge. - SDMMC_CK edge occurs on SDMMCCLK rising edge. - Data changed on the SDMMC_CK falling edge succeeding a SDMMC_CK edge. - SDMMC_CK edge occurs on SDMMCCLK rising edge.">,
      Field<"HWFC_EN", 17, 1, ReadWrite, "Hardware flow control enable This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) When Hardware flow control is enabled, the meaning of the TXFIFOE and RXFIFOF flags change, please see SDMMC status register definition in Section56.8.11.">,
      Field<"DDR", 18, 1, ReadWrite, "Data rate signaling selection This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0) DDR rate shall only be selected with 4-bit or 8-bit wide bus mode. (WIDBUS &gt; 00). DDR = 1 has no effect when WIDBUS = 00 (1-bit wide bus). DDR rate shall only be selected with clock division &gt;1. (CLKDIV &gt; 0)">,
      Field<"BUSSPEED", 19, 1, ReadWrite, "Bus speed mode selection between DS, HS, SDR12, SDR25 and SDR50, DDR50, SDR104. This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)">,
      Field<"SELCLKRX", 20, 2, ReadWrite, "Receive clock selection. These bits can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0)">,
    ], "The SDMMC_CLKCR register controls the SDMMC_CK output clock, the SDMMC_RX_CLK receive clock, and the bus width.">,
    Register<"ARGR", 0x8, 32, [
      Field<"CMDARG", 0, 32, ReadWrite, "Command argument. These bits can only be written by firmware when CPSM is disabled (CPSMEN = 0). Command argument sent to a card as part of a command message. If a command contains an argument, it must be loaded into this register before writing a command to the command register.">,
    ], "The SDMMC_ARGR register contains a 32-bit command argument, which is sent to a card as part of a command message.">,
    Register<"CMDR", 0xc, 32, [
      Field<"CMDINDEX", 0, 6, ReadWrite, "Command index. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). The command index is sent to the card as part of a command message.">,
      Field<"CMDTRANS", 6, 1, ReadWrite, "The CPSM treats the command as a data transfer command, stops the interrupt period, and signals DataEnable to the DPSM This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). If this bit is set, the CPSM issues an end of interrupt period and issues DataEnable signal to the DPSM when the command is sent.">,
      Field<"CMDSTOP", 7, 1, ReadWrite, "The CPSM treats the command as a Stop Transmission command and signals Abort to the DPSM. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). If this bit is set, the CPSM issues the Abort signal to the DPSM when the command is sent.">,
      Field<"WAITRESP", 8, 2, ReadWrite, "Wait for response bits. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). They are used to configure whether the CPSM is to wait for a response, and if yes, which kind of response.">,
      Field<"WAITINT", 10, 1, ReadWrite, "CPSM waits for interrupt request. If this bit is set, the CPSM disables command timeout and waits for an card interrupt request (Response). If this bit is cleared in the CPSM Wait state, will cause the abort of the interrupt mode.">,
      Field<"WAITPEND", 11, 1, ReadWrite, "CPSM Waits for end of data transfer (CmdPend internal signal) from DPSM. This bit when set, the CPSM waits for the end of data transfer trigger before it starts sending a command. WAITPEND is only taken into account when DTMODE = MMC stream data transfer, WIDBUS = 1-bit wide bus mode, DPSMACT = 1 and DTDIR = from host to card.">,
      Field<"CPSMEN", 12, 1, ReadWrite, "Command path state machine (CPSM) Enable bit This bit is written 1 by firmware, and cleared by hardware when the CPSM enters the Idle state. If this bit is set, the CPSM is enabled. When DTEN = 1, no command will be transfered nor boot procedure will be started. CPSMEN is cleared to 0.">,
      Field<"DTHOLD", 13, 1, ReadWrite, "Hold new data block transmission and reception in the DPSM. If this bit is set, the DPSM will not move from the Wait_S state to the Send state or from the Wait_R state to the Receive state.">,
      Field<"BOOTMODE", 14, 1, ReadWrite, "Select the boot mode procedure to be used. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0)">,
      Field<"BOOTEN", 15, 1, ReadWrite, "Enable boot mode procedure.">,
      Field<"CMDSUSPEND", 16, 1, ReadWrite, "The CPSM treats the command as a Suspend or Resume command and signals interrupt period start/end. This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). CMDSUSPEND = 1 and CMDTRANS = 0 Suspend command, start interrupt period when response bit BS=0. CMDSUSPEND = 1 and CMDTRANS = 1 Resume command with data, end interrupt period when response bit DF=1.">,
    ], "The SDMMC_CMDR register contains the command index and command type bits. The command index is sent to a card as part of a command message. The command type bits control the command path state machine (CPSM).">,
    Register<"RESPCMDR", 0x10, 32, [
      Field<"RESPCMD", 0, 6, ReadWrite, "Response command index">,
    ], "SDMMC command response register">,
    Register<"RESP1R", 0x14, 32, [
      Field<"CARDSTATUS1", 0, 32, ReadWrite, "see Table 432">,
    ], "The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.">,
    Register<"RESP2R", 0x18, 32, [
      Field<"CARDSTATUS2", 0, 32, ReadWrite, "see Table404.">,
    ], "The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.">,
    Register<"RESP3R", 0x1c, 32, [
      Field<"CARDSTATUS3", 0, 32, ReadWrite, "see Table404.">,
    ], "The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.">,
    Register<"RESP4R", 0x20, 32, [
      Field<"CARDSTATUS4", 0, 32, ReadWrite, "see Table404.">,
    ], "The SDMMC_RESP1/2/3/4R registers contain the status of a card, which is part of the received response.">,
    Register<"DTIMER", 0x24, 32, [
      Field<"DATATIME", 0, 32, ReadWrite, "Data and R1b busy timeout period This bit can only be written when the CPSM and DPSM are not active (CPSMACT = 0 and DPSMACT = 0). Data and R1b busy timeout period expressed in card bus clock periods.">,
    ], "The SDMMC_DTIMER register contains the data timeout period, in card bus clock periods. A counter loads the value from the SDMMC_DTIMER register, and starts decrementing when the data path state machine (DPSM) enters the Wait_R or Busy state. If the timer reaches 0 while the DPSM is in either of these states, the timeout status flag is set.">,
    Register<"DLENR", 0x28, 32, [
      Field<"DATALENGTH", 0, 25, ReadWrite, "Data length value This register can only be written by firmware when DPSM is inactive (DPSMACT = 0). Number of data bytes to be transferred. When DDR = 1 DATALENGTH is truncated to a multiple of 2. (The last odd byte is not transfered) When DATALENGTH = 0 no data will be transfered, when requested by a CPSMEN and CMDTRANS = 1 also no command will be transfered. DTEN and CPSMEN are cleared to 0.">,
    ], "The SDMMC_DLENR register contains the number of data bytes to be transferred. The value is loaded into the data counter when data transfer starts.">,
    Register<"DCTRL", 0x2c, 32, [
      Field<"DTEN", 0, 1, ReadWrite, "Data transfer enable bit This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). This bit is cleared by Hardware when data transfer completes. This bit shall only be used to transfer data when no associated data transfer command is used, i.e. shall not be used with SD or eMMC cards.">,
      Field<"DTDIR", 1, 1, ReadWrite, "Data transfer direction selection This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).">,
      Field<"DTMODE", 2, 2, ReadWrite, "Data transfer mode selection. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).">,
      Field<"DBLOCKSIZE", 4, 4, ReadWrite, "Data block size This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). Define the data block length when the block data transfer mode is selected: When DATALENGTH is not a multiple of DBLOCKSIZE, the transfered data is truncated at a multiple of DBLOCKSIZE. (Any remain data will not be transfered.) When DDR = 1, DBLOCKSIZE = 0000 shall not be used. (No data will be transfered)">,
      Field<"RWSTART", 8, 1, ReadWrite, "Read wait start. If this bit is set, read wait operation starts.">,
      Field<"RWSTOP", 9, 1, ReadWrite, "Read wait stop This bit is written by firmware and auto cleared by hardware when the DPSM moves from the READ_WAIT state to the WAIT_R or IDLE state.">,
      Field<"RWMOD", 10, 1, ReadWrite, "Read wait mode. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).">,
      Field<"SDIOEN", 11, 1, ReadWrite, "SD I/O interrupt enable functions This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). If this bit is set, the DPSM enables the SD I/O card specific interrupt operation.">,
      Field<"BOOTACKEN", 12, 1, ReadWrite, "Enable the reception of the boot acknowledgment. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).">,
      Field<"FIFORST", 13, 1, ReadWrite, "FIFO reset, will flush any remaining data. This bit can only be written by firmware when IDMAEN= 0 and DPSM is active (DPSMACT = 1). This bit will only take effect when a transfer error or transfer hold occurs.">,
    ], "The SDMMC_DCTRL register control the data path state machine (DPSM).">,
    Register<"DCNTR", 0x30, 32, [
      Field<"DATACOUNT", 0, 25, ReadWrite, "Data count value When read, the number of remaining data bytes to be transferred is returned. Write has no effect.">,
    ], "The SDMMC_DCNTR register loads the value from the data length register (see SDMMC_DLENR) when the DPSM moves from the Idle state to the Wait_R or Wait_S state. As data is transferred, the counter decrements the value until it reaches 0. The DPSM then moves to the Idle state and when there has been no error, the data status end flag (DATAEND) is set.">,
    Register<"STAR", 0x34, 32, [
      Field<"CCRCFAIL", 0, 1, ReadWrite, "Command response received (CRC check failed). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"DCRCFAIL", 1, 1, ReadWrite, "Data block sent/received (CRC check failed). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"CTIMEOUT", 2, 1, ReadWrite, "Command response timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR. The Command Timeout period has a fixed value of 64 SDMMC_CK clock periods.">,
      Field<"DTIMEOUT", 3, 1, ReadWrite, "Data timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"TXUNDERR", 4, 1, ReadWrite, "Transmit FIFO underrun error or IDMA read transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"RXOVERR", 5, 1, ReadWrite, "Received FIFO overrun error or IDMA write transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"CMDREND", 6, 1, ReadWrite, "Command response received (CRC check passed, or no CRC). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"CMDSENT", 7, 1, ReadWrite, "Command sent (no response required). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"DATAEND", 8, 1, ReadWrite, "Data transfer ended correctly. (data counter, DATACOUNT is zero and no errors occur). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"DHOLD", 9, 1, ReadWrite, "Data transfer Hold. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"DBCKEND", 10, 1, ReadWrite, "Data block sent/received. (CRC check passed) and DPSM moves to the READWAIT state. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"DABORT", 11, 1, ReadWrite, "Data transfer aborted by CMD12. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"DPSMACT", 12, 1, ReadWrite, "Data path state machine active, i.e. not in Idle state. This is a hardware status flag only, does not generate an interrupt.">,
      Field<"CPSMACT", 13, 1, ReadWrite, "Command path state machine active, i.e. not in Idle state. This is a hardware status flag only, does not generate an interrupt.">,
      Field<"TXFIFOHE", 14, 1, ReadWrite, "Transmit FIFO half empty At least half the number of words can be written into the FIFO. This bit is cleared when the FIFO becomes half+1 full.">,
      Field<"RXFIFOHF", 15, 1, ReadWrite, "Receive FIFO half full There are at least half the number of words in the FIFO. This bit is cleared when the FIFO becomes half+1 empty.">,
      Field<"TXFIFOF", 16, 1, ReadWrite, "Transmit FIFO full This is a hardware status flag only, does not generate an interrupt. This bit is cleared when one FIFO location becomes empty.">,
      Field<"RXFIFOF", 17, 1, ReadWrite, "Receive FIFO full This bit is cleared when one FIFO location becomes empty.">,
      Field<"TXFIFOE", 18, 1, ReadWrite, "Transmit FIFO empty This bit is cleared when one FIFO location becomes full.">,
      Field<"RXFIFOE", 19, 1, ReadWrite, "Receive FIFO empty This is a hardware status flag only, does not generate an interrupt. This bit is cleared when one FIFO location becomes full.">,
      Field<"BUSYD0", 20, 1, ReadWrite, "Inverted value of SDMMC_D0 line (Busy), sampled at the end of a CMD response and a second time 2 SDMMC_CK cycles after the CMD response. This bit is reset to not busy when the SDMMCD0 line changes from busy to not busy. This bit does not signal busy due to data transfer. This is a hardware status flag only, it does not generate an interrupt.">,
      Field<"BUSYD0END", 21, 1, ReadWrite, "end of SDMMC_D0 Busy following a CMD response detected. This indicates only end of busy following a CMD response. This bit does not signal busy due to data transfer. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"SDIOIT", 22, 1, ReadWrite, "SDIO interrupt received. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"ACKFAIL", 23, 1, ReadWrite, "Boot acknowledgment received (boot acknowledgment check fail). Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"ACKTIMEOUT", 24, 1, ReadWrite, "Boot acknowledgment timeout. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"VSWEND", 25, 1, ReadWrite, "Voltage switch critical timing section completion. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"CKSTOP", 26, 1, ReadWrite, "SDMMC_CK stopped in Voltage switch procedure. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"IDMATE", 27, 1, ReadWrite, "IDMA transfer error. Interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
      Field<"IDMABTC", 28, 1, ReadWrite, "IDMA buffer transfer complete. interrupt flag is cleared by writing corresponding interrupt clear bit in SDMMC_ICR.">,
    ], "The SDMMC_STAR register is a read-only register. It contains two types of flag:Static flags (bits [29,21,11:0]): these bits remain asserted until they are cleared by writing to the SDMMC interrupt Clear register (see SDMMC_ICR)Dynamic flags (bits [20:12]): these bits change state depending on the state of the underlying logic (for example, FIFO full and empty flags are asserted and de-asserted as data while written to the FIFO)">,
    Register<"ICR", 0x38, 32, [
      Field<"CCRCFAILC", 0, 1, ReadWrite, "CCRCFAIL flag clear bit Set by software to clear the CCRCFAIL flag.">,
      Field<"DCRCFAILC", 1, 1, ReadWrite, "DCRCFAIL flag clear bit Set by software to clear the DCRCFAIL flag.">,
      Field<"CTIMEOUTC", 2, 1, ReadWrite, "CTIMEOUT flag clear bit Set by software to clear the CTIMEOUT flag.">,
      Field<"DTIMEOUTC", 3, 1, ReadWrite, "DTIMEOUT flag clear bit Set by software to clear the DTIMEOUT flag.">,
      Field<"TXUNDERRC", 4, 1, ReadWrite, "TXUNDERR flag clear bit Set by software to clear TXUNDERR flag.">,
      Field<"RXOVERRC", 5, 1, ReadWrite, "RXOVERR flag clear bit Set by software to clear the RXOVERR flag.">,
      Field<"CMDRENDC", 6, 1, ReadWrite, "CMDREND flag clear bit Set by software to clear the CMDREND flag.">,
      Field<"CMDSENTC", 7, 1, ReadWrite, "CMDSENT flag clear bit Set by software to clear the CMDSENT flag.">,
      Field<"DATAENDC", 8, 1, ReadWrite, "DATAEND flag clear bit Set by software to clear the DATAEND flag.">,
      Field<"DHOLDC", 9, 1, ReadWrite, "DHOLD flag clear bit Set by software to clear the DHOLD flag.">,
      Field<"DBCKENDC", 10, 1, ReadWrite, "DBCKEND flag clear bit Set by software to clear the DBCKEND flag.">,
      Field<"DABORTC", 11, 1, ReadWrite, "DABORT flag clear bit Set by software to clear the DABORT flag.">,
      Field<"BUSYD0ENDC", 21, 1, ReadWrite, "BUSYD0END flag clear bit Set by software to clear the BUSYD0END flag.">,
      Field<"SDIOITC", 22, 1, ReadWrite, "SDIOIT flag clear bit Set by software to clear the SDIOIT flag.">,
      Field<"ACKFAILC", 23, 1, ReadWrite, "ACKFAIL flag clear bit Set by software to clear the ACKFAIL flag.">,
      Field<"ACKTIMEOUTC", 24, 1, ReadWrite, "ACKTIMEOUT flag clear bit Set by software to clear the ACKTIMEOUT flag.">,
      Field<"VSWENDC", 25, 1, ReadWrite, "VSWEND flag clear bit Set by software to clear the VSWEND flag.">,
      Field<"CKSTOPC", 26, 1, ReadWrite, "CKSTOP flag clear bit Set by software to clear the CKSTOP flag.">,
      Field<"IDMATEC", 27, 1, ReadWrite, "IDMA transfer error clear bit Set by software to clear the IDMATE flag.">,
      Field<"IDMABTCC", 28, 1, ReadWrite, "IDMA buffer transfer complete clear bit Set by software to clear the IDMABTC flag.">,
    ], "The SDMMC_ICR register is a write-only register. Writing a bit with 1 clears the corresponding bit in the SDMMC_STAR status register.">,
    Register<"MASKR", 0x3c, 32, [
      Field<"CCRCFAILIE", 0, 1, ReadWrite, "Command CRC fail interrupt enable Set and cleared by software to enable/disable interrupt caused by command CRC failure.">,
      Field<"DCRCFAILIE", 1, 1, ReadWrite, "Data CRC fail interrupt enable Set and cleared by software to enable/disable interrupt caused by data CRC failure.">,
      Field<"CTIMEOUTIE", 2, 1, ReadWrite, "Command timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by command timeout.">,
      Field<"DTIMEOUTIE", 3, 1, ReadWrite, "Data timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by data timeout.">,
      Field<"TXUNDERRIE", 4, 1, ReadWrite, "Tx FIFO underrun error interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO underrun error.">,
      Field<"RXOVERRIE", 5, 1, ReadWrite, "Rx FIFO overrun error interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO overrun error.">,
      Field<"CMDRENDIE", 6, 1, ReadWrite, "Command response received interrupt enable Set and cleared by software to enable/disable interrupt caused by receiving command response.">,
      Field<"CMDSENTIE", 7, 1, ReadWrite, "Command sent interrupt enable Set and cleared by software to enable/disable interrupt caused by sending command.">,
      Field<"DATAENDIE", 8, 1, ReadWrite, "Data end interrupt enable Set and cleared by software to enable/disable interrupt caused by data end.">,
      Field<"DHOLDIE", 9, 1, ReadWrite, "Data hold interrupt enable Set and cleared by software to enable/disable the interrupt generated when sending new data is hold in the DPSM Wait_S state.">,
      Field<"DBCKENDIE", 10, 1, ReadWrite, "Data block end interrupt enable Set and cleared by software to enable/disable interrupt caused by data block end.">,
      Field<"DABORTIE", 11, 1, ReadWrite, "Data transfer aborted interrupt enable Set and cleared by software to enable/disable interrupt caused by a data transfer being aborted.">,
      Field<"TXFIFOHEIE", 14, 1, ReadWrite, "Tx FIFO half empty interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO half empty.">,
      Field<"RXFIFOHFIE", 15, 1, ReadWrite, "Rx FIFO half full interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO half full.">,
      Field<"RXFIFOFIE", 17, 1, ReadWrite, "Rx FIFO full interrupt enable Set and cleared by software to enable/disable interrupt caused by Rx FIFO full.">,
      Field<"TXFIFOEIE", 18, 1, ReadWrite, "Tx FIFO empty interrupt enable Set and cleared by software to enable/disable interrupt caused by Tx FIFO empty.">,
      Field<"BUSYD0ENDIE", 21, 1, ReadWrite, "BUSYD0END interrupt enable Set and cleared by software to enable/disable the interrupt generated when SDMMC_D0 signal changes from busy to NOT busy following a CMD response.">,
      Field<"SDIOITIE", 22, 1, ReadWrite, "SDIO mode interrupt received interrupt enable Set and cleared by software to enable/disable the interrupt generated when receiving the SDIO mode interrupt.">,
      Field<"ACKFAILIE", 23, 1, ReadWrite, "Acknowledgment Fail interrupt enable Set and cleared by software to enable/disable interrupt caused by acknowledgment Fail.">,
      Field<"ACKTIMEOUTIE", 24, 1, ReadWrite, "Acknowledgment timeout interrupt enable Set and cleared by software to enable/disable interrupt caused by acknowledgment timeout.">,
      Field<"VSWENDIE", 25, 1, ReadWrite, "Voltage switch critical timing section completion interrupt enable Set and cleared by software to enable/disable the interrupt generated when voltage switch critical timing section completion.">,
      Field<"CKSTOPIE", 26, 1, ReadWrite, "Voltage Switch clock stopped interrupt enable Set and cleared by software to enable/disable interrupt caused by Voltage Switch clock stopped.">,
      Field<"IDMABTCIE", 28, 1, ReadWrite, "IDMA buffer transfer complete interrupt enable Set and cleared by software to enable/disable the interrupt generated when the IDMA has transferred all data belonging to a memory buffer.">,
    ], "The interrupt mask register determines which status flags generate an interrupt request by setting the corresponding bit to 1.">,
    Register<"ACKTIMER", 0x40, 32, [
      Field<"ACKTIME", 0, 25, ReadWrite, "Boot acknowledgment timeout period This bit can only be written by firmware when CPSM is disabled (CPSMEN = 0). Boot acknowledgment timeout period expressed in card bus clock periods.">,
    ], "The SDMMC_ACKTIMER register contains the acknowledgment timeout period, in SDMMC_CK bus clock periods. A counter loads the value from the SDMMC_ACKTIMER register, and starts decrementing when the data path state machine (DPSM) enters the Wait_Ack state. If the timer reaches 0 while the DPSM is in this states, the acknowledgment timeout status flag is set.">,
    Register<"IDMACTRLR", 0x50, 32, [
      Field<"IDMAEN", 0, 1, ReadWrite, "IDMA enable This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).">,
      Field<"IDMABMODE", 1, 1, ReadWrite, "Buffer mode selection. This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0).">,
      Field<"IDMABACT", 2, 1, ReadWrite, "Double buffer mode active buffer indication This bit can only be written by firmware when DPSM is inactive (DPSMACT = 0). When IDMA is enabled this bit is toggled by hardware.">,
    ], "The receive and transmit FIFOs can be read or written as 32-bit wide registers. The FIFOs contain 32 entries on 32 sequential addresses. This allows the CPU to use its load and store multiple operands to read from/write to the FIFO.">,
    Register<"IDMABSIZER", 0x54, 32, [
      Field<"IDMABNDT", 5, 8, ReadWrite, "Number of transfers per buffer. This 8-bit value shall be multiplied by 8 to get the size of the buffer in 32-bit words and by 32 to get the size of the buffer in bytes. Example: IDMABNDT = 0x01: buffer size = 8 words = 32 bytes. These bits can only be written by firmware when DPSM is inactive (DPSMACT = 0).">,
    ], "The SDMMC_IDMABSIZER register contains the buffers size when in double buffer configuration.">,
    Register<"IDMABASE0R", 0x58, 32, [
      Field<"IDMABASE0", 0, 32, ReadWrite, "Buffer 0 memory base address bits [31:2], shall be word aligned (bit [1:0] are always 0 and read only). This register can be written by firmware when DPSM is inactive (DPSMACT = 0), and can dynamically be written by firmware when DPSM active (DPSMACT = 1) and memory buffer 0 is inactive (IDMABACT = 1).">,
    ], "The SDMMC_IDMABASE0R register contains the memory buffer base address in single buffer configuration and the buffer 0 base address in double buffer configuration.">,
    Register<"IDMABASE1R", 0x5c, 32, [
      Field<"IDMABASE1", 0, 32, ReadWrite, "Buffer 1 memory base address, shall be word aligned (bit [1:0] are always 0 and read only). This register can be written by firmware when DPSM is inactive (DPSMACT = 0), and can dynamically be written by firmware when DPSM active (DPSMACT = 1) and memory buffer 1 is inactive (IDMABACT = 0).">,
    ], "The SDMMC_IDMABASE1R register contains the double buffer configuration second buffer memory base address.">,
    Register<"FIFOR", 0x80, 32, [
      Field<"FIFODATA", 0, 32, ReadWrite, "Receive and transmit FIFO data This register can only be read or written by firmware when the DPSM is active (DPSMACT=1). The FIFO data occupies 16 entries of 32-bit words.">,
    ], "The receive and transmit FIFOs can be only read or written as word (32-bit) wide registers. The FIFOs contain 16 entries on sequential addresses. This allows the CPU to use its load and store multiple operands to read from/write to the FIFO.When accessing SDMMC_FIFOR with half word or byte access an AHB bus fault is generated.">,
  ];
}

class SDMMC1Instance<string Name, int Base> : PeripheralInstance<Name, Base, SDMMC1Peripheral>;
def SDMMC2 : SDMMC1Instance<"SDMMC2", 0x48022400>;
def SDMMC1 : SDMMC1Instance<"SDMMC1", 0x52007000>;

#endif // _PERIPHERALS_SDMMC1_TD
