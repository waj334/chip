.global HardfaultHandler
.type HardfaultHandler, %function
HardfaultHandler:
#ifdef __thumb2__
    tst   lr, #4
    ite   eq
    mrseq r0, msp
    mrsne r0, psp
#else
    mov   r1, lr
    lsrs  r1, #2
    bcc   use_msp
    mrs   r0, psp
    b     1f
use_msp:
    mrs   r0, msp
1:
#endif
    b     runtime.hardfault
