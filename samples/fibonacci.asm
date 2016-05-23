; 
; Now I need modify for compile with naken ;)
;  -- test asm---
;
;-------------------------------------------------------------------------------
; MSP430 Assembler Code Template for use with TI Code Composer Studio
;
; This is the Fibonacci Sequence Program. f0=0, f1=1, fn=(fn-1)+(fn-2).
;-------------------------------------------------------------------------------
            .cdecls C,LIST,"msp430.h"       ; Include device header file

;-------------------------------------------------------------------------------
            .text                           ; Assemble into program memory
            .global RESET
            .retain                         ; Override ELF conditional linking
                                            ; and retain current section
            .retainrefs                     ; Additionally retain any sections
                                            ; that have references to current
                                            ; section
;-------------------------------------------------------------------------------
RESET       mov.w   #__STACK_END,SP         ; Initialize stackpointer
StopWDT     mov.w   #WDTPW|WDTHOLD,&WDTCTL  ; Stop watchdog timer

;-------------------------------------------------------------------------------
                                            ; Main loop here
;-------------------------------------------------------------------------------
            mov.w   #0x2400, r9
            mov.w   #10, r10

            mov.w   #0, r11
            mov.w   r11, 0(r9)

            mov.w   #1, r12
            incd    r9
            mov.w   r12, 0(r9)

loop        tst     r10
            jz      forever

            incd    r9
            dec     r10

            mov.w   r12, r13
            add.w   r11, r12
            mov.w   r12, 0(r9)
            mov.w   r13, r11
            jmp     loop

forever     jmp     forever
;-------------------------------------------------------------------------------
;           Stack Pointer definition
;-------------------------------------------------------------------------------
            .global __STACK_END
            .sect 	.stack

;-------------------------------------------------------------------------------
;           Interrupt Vectors
;-------------------------------------------------------------------------------
            .sect   ".reset"                ; MSP430 RESET Vector
            .short  RESET
