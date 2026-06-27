.syntax unified
.thumb
.text

// Keep these offsets in sync with runtime.goroutine on 32-bit Cortex-M:
//
//   type goroutine struct {
//       stackTop unsafe.Pointer // +0
//       __func struct {
//           f         unsafe.Pointer // +4
//           args      unsafe.Pointer // +8
//           stackSize uintptr        // +12
//       }
//       stack unsafe.Pointer // +16
//       ...
//   }
.equ GOROUTINE_STACK_TOP_OFFSET,   0
.equ GOROUTINE_STACK_SIZE_OFFSET, 12
.equ GOROUTINE_STACK_OFFSET,      16

// Software-saved context:
//   r4-r11 + EXC_RETURN = 9 words = 36 bytes
.equ SOFTWARE_CONTEXT_SIZE,       36

// s16-s31 are software-saved only when EXC_RETURN[4] == 0.
.equ FP_HIGH_CONTEXT_SIZE,        64

// Hardware exception-frame sizes, excluding the optional 4-byte alignment pad.
.equ BASIC_EXCEPTION_FRAME_SIZE,  32
.equ FP_EXCEPTION_FRAME_SIZE,    104

// The low end of every goroutine stack is reserved for 16 canary words.
.equ STACK_GUARD_SIZE,            64

.equ XPSR_T_BIT,          0x01000000
.equ XPSR_STACK_ALIGN_BIT,     0x00000200
.equ EXC_RETURN_THREAD_PSP,     0xFFFFFFFD
.equ EXC_RETURN_THREAD_PSP_FP,  0xFFFFFFED

.equ STACKED_XPSR_OFFSET, 28

.global runtime.currentGoroutine
.global runtime.lastGoroutine
.global runtime.lastExcReturn
.global runtime.lastRestorePSP

.global PendSvHandler
.type PendSvHandler, %function
.thumb_func
PendSvHandler:
    // PendSV must run with interrupts masked while the scheduler updates the
    // goroutine ring and while the outgoing/incoming contexts are exchanged.
    cpsid i

    // Preserve the hardware-supplied EXC_RETURN across the Go scheduler call.
    // Push two words so MSP remains 8-byte aligned at the call boundary.
    push {r1, lr}
    bl runtime.schedule
#ifdef __thumb2__
    pop {r1, lr}
#else
    // Thumb-1 POP cannot target LR directly.
    pop {r1, r2}
    mov lr, r2
#endif

    // schedule() returns false when the interrupted goroutine should continue.
    cmp r0, #0
    beq .Lrestore_and_return

    // r2 = incoming/current goroutine
    // r3 = outgoing/last goroutine
    ldr r2, =runtime.currentGoroutine
    ldr r2, [r2]
    ldr r3, =runtime.lastGoroutine
    ldr r3, [r3]

    // A scheduler result that selects the same goroutine needs no switch.
    cmp r3, r2
    beq .Lrestore_and_return

    // The first scheduled goroutine has no outgoing context to save.  Its
    // PendSV may have interrupted Thread mode using MSP (EXC_RETURN F9), which
    // is fine because that MSP context is intentionally discarded.
    cmp r3, #0
    beq .Lrestore_incoming

    // Once a goroutine is running, PendSV must have interrupted Thread mode
    // using PSP.  If it ever preempts another handler, LR will contain F1/E1.
    ldr r0, =EXC_RETURN_THREAD_PSP
    cmp lr, r0
    beq .Lentry_exc_return_valid
#ifdef __ARM_FP
    ldr r0, =EXC_RETURN_THREAD_PSP_FP
    cmp lr, r0
    beq .Lentry_exc_return_valid
#endif
    b .Lbad_entry_exc_return

.Lentry_exc_return_valid:

    //--------------------------------------------------------------------------
    // Save and validate the outgoing context.
    //--------------------------------------------------------------------------
    mrs r1, psp                         // r1 = hardware exception-frame base

    // r0 = allocation base
    // r2 = allocation end
    ldr r0, [r3, #GOROUTINE_STACK_OFFSET]
    ldr r2, [r3, #GOROUTINE_STACK_SIZE_OFFSET]
    adds r2, r0, r2

    // PSP itself must already lie inside the outgoing stack allocation.  This
    // catches an overflow caused by hardware exception stacking before PendSV
    // writes another byte.
    cmp r1, r0
    blo .Lstack_overflow
    cmp r1, r2
    bhs .Lstack_overflow

    // Verify that the complete hardware-created frame lies below allocation
    // end.  The optional alignment word is indicated by stacked xPSR[9].
#ifdef __ARM_FP
    tst lr, #0x10
    beq .Lvalidate_outgoing_fp_frame
#endif

.Lvalidate_outgoing_basic_frame:
    // Require PSP + 32 <= allocation end before reading stacked xPSR.
    // Use only low registers here so this path also assembles for ARMv6-M.
    mov r0, r1
    adds r0, r0, #BASIC_EXCEPTION_FRAME_SIZE
    cmp r0, r2
    bhi .Lstack_overflow

    ldr r0, [r1, #28]                  // stacked xPSR
    ldr r2, =XPSR_STACK_ALIGN_BIT
    tst r0, r2
    beq .Lvalidate_outgoing_save_space

    // Alignment pad present: require PSP + 36 <= allocation end.
    ldr r0, [r3, #GOROUTINE_STACK_OFFSET]
    ldr r2, [r3, #GOROUTINE_STACK_SIZE_OFFSET]
    adds r2, r0, r2
    mov r0, r1
    adds r0, r0, #(BASIC_EXCEPTION_FRAME_SIZE + 4)
    cmp r0, r2
    bhi .Lstack_overflow
    b .Lvalidate_outgoing_save_space

#ifdef __ARM_FP
.Lvalidate_outgoing_fp_frame:
    // Require PSP + 104 <= allocation end.
    mov r12, r1
    add r12, r12, #FP_EXCEPTION_FRAME_SIZE
    cmp r12, r2
    bhi .Lstack_overflow

    // xPSR is at +28 even for an extended frame.
    ldr r12, [r1, #STACKED_XPSR_OFFSET]
    ldr r2, =XPSR_STACK_ALIGN_BIT
    tst r12, r2
    beq .Lvalidate_outgoing_save_space

    // Alignment pad present: require PSP + 108 <= allocation end.
    ldr r2, [r3, #GOROUTINE_STACK_SIZE_OFFSET]
    adds r2, r0, r2
    mov r12, r1
    add r12, r12, #(FP_EXCEPTION_FRAME_SIZE + 4)
    cmp r12, r2
    bhi .Lstack_overflow
#endif

.Lvalidate_outgoing_save_space:
    // Reload allocation base because the frame checks use r0 as a scratch.
    ldr r0, [r3, #GOROUTINE_STACK_OFFSET]

    // Compute the stack pointer that the software save is about to produce.
    // Do this before VSTMDB/STMDB so an overflow cannot corrupt a neighbor.
    mov r2, r1
    subs r2, r2, #SOFTWARE_CONTEXT_SIZE
#ifdef __ARM_FP
    tst lr, #0x10
    it eq
    subeq r2, r2, #FP_HIGH_CONTEXT_SIZE
#endif

    // The prospective saved SP must remain above the 64-byte guard region.
    adds r0, r0, #STACK_GUARD_SIZE
    cmp r2, r0
    blo .Lstack_overflow

#ifdef __thumb2__
#ifdef __ARM_FP
    // Save s16-s31 first.  The integer block is then pushed beneath it, so the
    // restore path reads r4-r11/EXC_RETURN first and the FP block second.
    tst lr, #0x10
    it eq
    vstmdbeq r1!, {s16-s31}
#endif
    stmdb r1!, {r4-r11, lr}
#else
    // ARMv6-M has no supported FPU target here.  Reserve and populate the same
    // nine-word integer context used by initGoroutine.
    subs r1, r1, #SOFTWARE_CONTEXT_SIZE

    str r4, [r1, #0]
    str r5, [r1, #4]
    str r6, [r1, #8]
    str r7, [r1, #12]

    mov r2, r8
    str r2, [r1, #16]
    mov r2, r9
    str r2, [r1, #20]
    mov r2, r10
    str r2, [r1, #24]
    mov r2, r11
    str r2, [r1, #28]
    mov r2, lr
    str r2, [r1, #32]
#endif

    // Publish the outgoing software-context base only after the full save has
    // completed successfully.
    str r1, [r3, #GOROUTINE_STACK_TOP_OFFSET]

.Lrestore_incoming:
    //--------------------------------------------------------------------------
    // Validate the incoming saved context before loading any callee-saved regs.
    //--------------------------------------------------------------------------
    ldr r2, =runtime.currentGoroutine
    ldr r2, [r2]
    cmp r2, #0
    beq .Lbad_saved_context

    ldr r0, [r2, #GOROUTINE_STACK_TOP_OFFSET]  // candidate software-context base
    ldr r1, [r2, #GOROUTINE_STACK_OFFSET]      // allocation base
    ldr r3, [r2, #GOROUTINE_STACK_SIZE_OFFSET]
    adds r3, r1, r3                            // allocation end

    // Saved SP must be above the guard and leave room for the 9-word software
    // context before we read its saved EXC_RETURN word.
    adds r1, r1, #STACK_GUARD_SIZE
    cmp r0, r1
    blo .Lstack_overflow

    mov r1, r0
    adds r1, r1, #SOFTWARE_CONTEXT_SIZE
    cmp r1, r3
    bhi .Lstack_overflow

    // With an 8-byte-aligned hardware frame and a 36-byte software frame, the
    // saved software-context base is naturally 4 modulo 8.
    movs r1, #7
    ands r1, r0
    cmp r1, #4
    bne .Lbad_saved_context

    // Validate the saved EXC_RETURN before it controls FP frame interpretation.
    ldr r1, [r0, #32]
    ldr r2, =EXC_RETURN_THREAD_PSP
    cmp r1, r2
    beq .Lvalidate_incoming_basic_frame
#ifdef __ARM_FP
    ldr r2, =EXC_RETURN_THREAD_PSP_FP
    cmp r1, r2
    beq .Lvalidate_incoming_fp_frame
#endif
    b .Lbad_saved_context

.Lvalidate_incoming_basic_frame:
    // Layout from saved SP:
    //   +0   software r4-r11/EXC_RETURN (36 bytes)
    //   +36  basic hardware frame       (32 bytes)
    //   +68  optional alignment pad
    mov r2, r0
    adds r2, r2, #(SOFTWARE_CONTEXT_SIZE + BASIC_EXCEPTION_FRAME_SIZE)
    cmp r2, r3
    bhi .Lstack_overflow

    ldr r2, [r0, #(SOFTWARE_CONTEXT_SIZE + 28)] // stacked xPSR at +64
    ldr r1, =XPSR_T_BIT
    tst r2, r1
    beq .Lbad_stacked_xpsr

    ldr r1, =XPSR_STACK_ALIGN_BIT
    tst r2, r1
    beq .Lrestore_integer_context

    mov r2, r0
    adds r2, r2, #(SOFTWARE_CONTEXT_SIZE + BASIC_EXCEPTION_FRAME_SIZE + 4)
    cmp r2, r3
    bhi .Lstack_overflow
    b .Lrestore_integer_context

#ifdef __ARM_FP
.Lvalidate_incoming_fp_frame:
    // Layout:
    //   +0    r4-r11/EXC_RETURN       36 bytes
    //   +36   s16-s31                 64 bytes
    //   +100  hardware extended frame 104 bytes
    //   +204  optional alignment word
    mov r2, r0
    add r2, r2, #(SOFTWARE_CONTEXT_SIZE + FP_HIGH_CONTEXT_SIZE + FP_EXCEPTION_FRAME_SIZE)
    cmp r2, r3
    bhi .Lstack_overflow

    // 36 + 64 + 28 = 128.
    ldr r2, [r0, #(SOFTWARE_CONTEXT_SIZE + FP_HIGH_CONTEXT_SIZE + STACKED_XPSR_OFFSET)]
    ldr r1, =XPSR_T_BIT
    tst r2, r1
    beq .Lbad_stacked_xpsr

    ldr r1, =XPSR_STACK_ALIGN_BIT
    tst r2, r1
    beq .Lrestore_integer_context

    mov r2, r0
    add r2, r2, #(SOFTWARE_CONTEXT_SIZE + FP_HIGH_CONTEXT_SIZE + FP_EXCEPTION_FRAME_SIZE + 4)
    cmp r2, r3
    bhi .Lstack_overflow
#endif

.Lrestore_integer_context:
#ifdef __thumb2__
    ldmia r0!, {r4-r12}                // r12 receives saved EXC_RETURN
    mov lr, r12

#ifdef __ARM_FP
    tst lr, #0x10
    it eq
    vldmiaeq r0!, {s16-s31}
#endif
#else
    ldr r4, [r0, #0]
    ldr r5, [r0, #4]
    ldr r6, [r0, #8]
    ldr r7, [r0, #12]

    ldr r3, [r0, #16]
    mov r8, r3
    ldr r3, [r0, #20]
    mov r9, r3
    ldr r3, [r0, #24]
    mov r10, r3
    ldr r3, [r0, #28]
    mov r11, r3
    ldr r3, [r0, #32]
    mov lr, r3
    adds r0, r0, #SOFTWARE_CONTEXT_SIZE
#endif

    // Record the exact state used for exception return.  These writes only use
    // caller-saved registers, which the hardware frame will restore.
    ldr r2, =runtime.lastExcReturn
    mov r3, lr
    str r3, [r2]
    ldr r2, =runtime.lastRestorePSP
    str r0, [r2]

    msr psp, r0
    isb
    cpsie i
    bx lr

.Lrestore_and_return:
    cpsie i
    bx lr

// Distinct breakpoints make the failing invariant obvious when debugging.
.Lbad_entry_exc_return:
    bkpt #1                            // PendSV entered from invalid mode/stack
    b .Lbad_entry_exc_return

.Lbad_saved_context:
    bkpt #0                            // malformed saved software context
    b .Lbad_saved_context

.Lbad_stacked_xpsr:
    bkpt #2                            // stacked xPSR has T-bit clear
    b .Lbad_stacked_xpsr

.Lstack_overflow:
    bkpt #3                            // stack bounds/guard violation
    b .Lstack_overflow

.size PendSvHandler, .-PendSvHandler
.align 4
