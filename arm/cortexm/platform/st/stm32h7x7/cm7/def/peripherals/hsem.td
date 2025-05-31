#ifndef _PERIPHERALS_HSEM_TD
#define _PERIPHERALS_HSEM_TD

include "base.td"

def HSEMPeripheral : PeripheralType<"HSEM", "HSEM"> {
  let accessWidth = 32;
  let registers = [
    Register<"HSEM_R0", 0x0, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R1", 0x4, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R2", 0x8, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R3", 0xc, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R4", 0x10, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R5", 0x14, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R6", 0x18, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R7", 0x1c, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R8", 0x20, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R9", 0x24, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R10", 0x28, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R11", 0x2c, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R12", 0x30, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R13", 0x34, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R14", 0x38, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R15", 0x3c, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R16", 0x40, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R17", 0x44, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R18", 0x48, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R19", 0x4c, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R20", 0x50, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R21", 0x54, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R22", 0x58, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R23", 0x5c, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R24", 0x60, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R25", 0x64, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R26", 0x68, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R27", 0x6c, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R28", 0x70, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R29", 0x74, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R30", 0x78, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_R31", 0x7c, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM register HSEM_R0 HSEM_R31">,
    Register<"HSEM_RLR0", 0x80, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR1", 0x84, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR2", 0x88, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR3", 0x8c, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR4", 0x90, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR5", 0x94, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR6", 0x98, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR7", 0x9c, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR8", 0xa0, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR9", 0xa4, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR10", 0xa8, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR11", 0xac, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR12", 0xb0, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR13", 0xb4, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR14", 0xb8, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR15", 0xbc, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR16", 0xc0, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR17", 0xc4, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR18", 0xc8, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR19", 0xcc, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR20", 0xd0, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR21", 0xd4, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR22", 0xd8, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR23", 0xdc, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR24", 0xe0, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR25", 0xe4, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR26", 0xe8, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR27", 0xec, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR28", 0xf0, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR29", 0xf4, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR30", 0xf8, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_RLR31", 0xfc, 32, [
      Field<"PROCID", 0, 8, ReadWrite, "Semaphore ProcessID">,
      Field<"COREID", 8, 4, ReadWrite, "Semaphore COREID">,
      Field<"LOCK", 31, 1, ReadWrite, "Lock indication">,
    ], "HSEM Read lock register">,
    Register<"HSEM_C1IER", 0x100, 32, [
      Field<"ISE", 0, 32, ReadWrite, "Interrupt semaphore x enable bit">,
    ], "HSEM Interrupt enable register">,
    Register<"HSEM_C1ICR", 0x104, 32, [
      Field<"ISC", 0, 32, ReadWrite, "Interrupt semaphore x clear bit">,
    ], "HSEM Interrupt clear register">,
    Register<"HSEM_C1ISR", 0x108, 32, [
      Field<"ISF", 0, 32, ReadWrite, "Interrupt semaphore x status bit before enable (mask)">,
    ], "HSEM Interrupt status register">,
    Register<"HSEM_C1MISR", 0x10c, 32, [
      Field<"MISF", 0, 32, ReadWrite, "masked interrupt semaphore x status bit after enable (mask)">,
    ], "HSEM Masked interrupt status register">,
    Register<"HSEM_C2IER", 0x110, 32, [
      Field<"ISE", 0, 32, ReadWrite, "Interrupt semaphore x enable bit">,
    ], "HSEM Interrupt enable register">,
    Register<"HSEM_C2ICR", 0x114, 32, [
      Field<"ISC", 0, 32, ReadWrite, "Interrupt semaphore x clear bit">,
    ], "HSEM Interrupt clear register">,
    Register<"HSEM_C2ISR", 0x118, 32, [
      Field<"ISF", 0, 32, ReadWrite, "Interrupt semaphore x status bit before enable (mask)">,
    ], "HSEM Interrupt status register">,
    Register<"HSEM_C2MISR", 0x11c, 32, [
      Field<"MISF", 0, 32, ReadWrite, "masked interrupt semaphore x status bit after enable (mask)">,
    ], "HSEM Masked interrupt status register">,
    Register<"HSEM_CR", 0x140, 32, [
      Field<"COREID", 8, 4, ReadWrite, "COREID of semaphores to be cleared">,
      Field<"KEY", 16, 16, ReadWrite, "Semaphore clear Key">,
    ], "HSEM Clear register">,
    Register<"HSEM_KEYR", 0x144, 32, [
      Field<"KEY", 16, 16, ReadWrite, "Semaphore Clear Key">,
    ], "HSEM Interrupt clear register">,
  ];
}

class HSEMInstance<string Name, int Base> : PeripheralInstance<Name, Base, HSEMPeripheral>;
def HSEM : HSEMInstance<"HSEM", 0x58026400>;

#endif // _PERIPHERALS_HSEM_TD
