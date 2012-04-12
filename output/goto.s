	ADD PC, 3
	:rnd1
	DAT 0x6769
	:rnd2
	DAT 0x1250
	:timer
	DAT 0
	SET PUSH, X
	SET PUSH, Y
	SET PUSH, Z
	SET PUSH, I
	SET PUSH, J
	SET A, SP
	SET PUSH, A
	SET Y, 0x7000
	SET Z, 0x9000
	SUB SP, 3 ; Alloc space on stack
	SET A, 0x0
	SET [0xffff], A
	SET A, 0x0
	SET [0xfffd], A
	ADD PC, 28
	:c0 DAT "I imagined a number between", 0
	SET A, c0
	BOR A, 0x8000
	JSR print
	JSR printnl
	ADD PC, 26
	:c1 DAT "1 and 9. Try to guess it!", 0
	SET A, c1
	BOR A, 0x8000
	JSR print
	JSR printnl
	ADD PC, 2
	:c2 DAT " ", 0
	SET A, c2
	BOR A, 0x8000
	JSR print
	JSR printnl
	:START
	ADD PC, 22
	:c3 DAT "Pick a number [1-9]: ", 0
	SET A, c3
	BOR A, 0x8000
	JSR print
	SET PUSH, 0x0
	SET I, SP
	SUB I, 1
	:input
	SET A, 0
	JSR getkey
	IFE A, 0
	SET PC, input
	IFE A, 0xa
	SET PC, input2
	IFE A, 0x8
	SET PC, inputbsp
	SET PUSH, A
	JSR printchar
	SET PC, input
	:inputbsp
	SET POP, 0
	SUB X, 1
	SET [0x8000+X], 0
	BOR [0x8000+X], Y
	SET PC, input
	:input2
	SET B, SP
	SET J, B
	:input3
	SET A, [B]
	SET [B], [I]
	SET [I], A
	ADD B, 1
	SUB I, 1
	IFG B, I
	SET PC, input4
	SET PC, input3
	:input4
	SET A, J
	BOR A, 0x8000
	IFG 0xF000, A
	AND A, 0x7fff
	SET PUSH, [A]
	SET A, POP
	SET PUSH, A
	SET A, 0x30
	SUB A, POP
	SET PUSH, A
	SET A, 0
	SUB A, POP
	SET [0xfffe], A
	SET A, [0xfffd]
	SET PUSH, A
	SET A, 0x1
	ADD A, POP
	SET [0xfffd], A
	SET I, 0x8220
	:l0
	SUB I, 1
	SET [I], 0
	IFN I, 0x8000
	SET PC, l0
	SET X, 0
	SET A, [0xffff]
	SET PUSH, A
	SET A, 0x0
	SET B, POP
	SET C, 1
	IFE A, B
	SET C, 0
	IFN C, 0
	SET PC, l1
	JSR rand
	SET PUSH, A
	SET A, 0x9
	SET B, POP
	MOD B, A
	SET A, B
	SET PUSH, A
	SET A, 0x1
	ADD A, POP
	SET [0xffff], A
	:l1
	SET A, [0xfffe]
	SET PUSH, A
	SET A, [0xffff]
	SET B, POP
	SET C, 1
	IFG B, A
	SET C, 0
	IFN C, 0
	SET PC, l2
	ADD PC, 27
	:c4 DAT "Wrong! Try a lower number.", 0
	SET A, c4
	BOR A, 0x8000
	JSR print
	JSR printnl
	SET PC, START
	:l2
	SET A, [0xfffe]
	SET PUSH, A
	SET A, [0xffff]
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFN C, 0
	SET PC, l3
	ADD PC, 28
	:c5 DAT "Wrong! Try a higher number.", 0
	SET A, c5
	BOR A, 0x8000
	JSR print
	JSR printnl
	SET PC, START
	:l3
	ADD PC, 23
	:c6 DAT "Success! My number is ", 0
	SET A, c6
	BOR A, 0x8000
	JSR print
	SET A, [0xffff]
	JSR print
	ADD PC, 2
	:c7 DAT "!", 0
	SET A, c7
	BOR A, 0x8000
	JSR print
	JSR printnl
	ADD PC, 13
	:c8 DAT "It took you ", 0
	SET A, c8
	BOR A, 0x8000
	JSR print
	SET A, [0xfffd]
	JSR print
	ADD PC, 10
	:c9 DAT " time(s).", 0
	SET A, c9
	BOR A, 0x8000
	JSR print
	JSR printnl
	SET J, POP
	SET I, POP
	SET Z, POP
	SET Y, POP
	SET X, POP
	SET A, POP
	SET SP, A
	SET PC, end
	
	; compiled functions
	:getkey
	ADD [timer], 1
	IFE [Z], 0
	SET PC, POP
	SET A, [Z]
	SET [Z], 0
	ADD Z, 1
	AND Z, 0x900f
	MUL [rnd1], [timer]
	ADD [rnd2], O
	SET PC, POP
	:strlen
	SET I, A
	:strlen1
	ADD I, 1
	IFN [I], 0x0
	SET PC, strlen1
	SET A, B
	SET PC, POP
	:printchar
	SET [0x8000+X], A
	BOR [0x8000+X], Y
	ADD X, 1
	IFG X, 0x21f
	SET X, 0
	:pnline
	SET PC, POP
	:printint
	SET I, 0
	:printint1
	SET B, A
	MOD A, 0xa
	ADD A, 0x30
	SET PUSH, A
	SET A, B
	DIV A, 0xa
	ADD I, 1
	IFN A, 0
	SET PC, printint1
	:printint2
	SET A, POP
	JSR printchar
	SUB I, 1
	IFN I, 0
	SET PC, printint2
	SET A, POP
	SET PC, POP
	:printstr
	IFG 0xF000, A
	AND A, 0x7fff
	SET I, A
	:printstr1
	IFE [I], 0
	SET PC, POP
	SET A, [I]
	JSR printchar
	ADD I, 1
	SET PC, printstr1
	:printnl
	DIV X, 32
	ADD X, 1
	MUL X, 32
	SET PC, POP
	:print
	SET B, A
	SHR B, 15
	IFE B, 0
	JSR printint
	IFE B, 1
	JSR printstr
	SET PC, POP
	:rand
	SET B, [rnd1]
	SET A, [rnd2]
	MUL [rnd1], 0x660D
	SET C, O
	MUL A, 0x660D
	ADD A, C
	MUL B, 0x0019
	ADD A, B
	ADD [rnd1], 1
	ADD A, O
	SET [rnd2], A
	SET PC, POP
	:end
	SET PC, end
