.syntax unified
.thumb

// Symmetric coroutine switch primitive for Cortex-M.
//
// runtime.coroswitch_inner(c *coro)
//   r0 = pointer to the coro struct (coro.sp at offset 0)
//
// Saves the current callee-saved register set + LR onto the current stack.
// Atomically swaps SP with c.sp (so c.sp now records THIS side's saved SP,
// and SP is loaded with the OTHER side's previously-saved SP). Pops the
// other side's callee-saved registers + LR from the new stack and returns.
//
// The result is symmetric: each call to coroswitch_inner pauses the
// caller's execution context on its current stack, resumes whoever was
// previously paused on the other side, and runs until that side calls
// coroswitch_inner in turn.
//
// On thumb2 with FPU, also save/restore callee-saved FP regs (s16..s31).

.global runtime.coroswitch_inner
.type runtime.coroswitch_inner, %function
runtime.coroswitch_inner:
#ifdef __thumb2__
    push    {r4-r11, lr}                // save integer callee-saved + return
#ifdef __ARM_FP
    vstmdb  sp!, {s16-s31}              // save FP callee-saved
#endif
    ldr     r1, [r0]                    // r1 = c.sp (other side's previously-saved SP)
    str     sp, [r0]                    // c.sp = current SP (we are now parked)
    mov     sp, r1                      // SP = other side's SP
#ifdef __ARM_FP
    vldmia  sp!, {s16-s31}              // restore FP callee-saved
#endif
    pop     {r4-r11, lr}                // restore integer callee-saved + return
#else
    // Thumb1 (Cortex-M0+, armv6-m). No FPU on any supported armv6m target.
    // Cannot push high registers (r8-r11) directly; route through r2/r3.
    push    {r4-r7, lr}                 // save r4-r7 + lr (low regs)
    mov     r2, r8
    mov     r3, r9
    push    {r2, r3}                    // save r8, r9
    mov     r2, r10
    mov     r3, r11
    push    {r2, r3}                    // save r10, r11

    ldr     r1, [r0]                    // r1 = c.sp
    str     sp, [r0]                    // c.sp = current SP
    mov     sp, r1                      // SP = other side's SP

    pop     {r2, r3}                    // restore r10, r11
    mov     r10, r2
    mov     r11, r3
    pop     {r2, r3}                    // restore r8, r9
    mov     r8, r2
    mov     r9, r3
    pop     {r4-r7}                     // restore r4-r7
    pop     {r1}                        // restore lr (popped into r1, no high-reg pop on thumb1)
    mov     lr, r1
#endif
    bx      lr
.size runtime.coroswitch_inner, .-runtime.coroswitch_inner
.align 4
