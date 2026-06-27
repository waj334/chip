.syntax unified
.thumb

.global runtime.exec
.type runtime.exec, %function

runtime.exec:
    // Check whether runtime.exec itself was entered correctly.
    mov  r12, sp
    tst  r12, #7
    bne  .Lbad_entry_stack_alignment

    push {r7, lr}
    mov  r7, sp

    // Check whether our prologue preserved alignment.
    mov  r12, sp
    tst  r12, #7
    bne  .Lbad_call_stack_alignment

    blx  r1
    pop  {r7, pc}

.Lbad_entry_stack_alignment:
    bkpt #4
    b    .Lbad_entry_stack_alignment

.Lbad_call_stack_alignment:
    bkpt #5
    b    .Lbad_call_stack_alignment

.size runtime.exec, .-runtime.exec