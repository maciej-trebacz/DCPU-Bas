	SUB SP, 5 ; Alloc space on stack
	SET Y, 0x7000
	ADD PC, 6
	:c0 DAT "Hello", 0
	SET A, c0
	BOR A, 0x8000
	JSR print
	ADD PC, 2
	:c1 DAT " ", 0
	SET A, c1
	BOR A, 0x8000
	JSR print
	ADD PC, 7
	:c2 DAT "world!", 0
	SET A, c2
	BOR A, 0x8000
	JSR print
	JSR printnl
	ADD PC, 6
	:c3 DAT "0x10c", 0
	SET A, c3
	BOR A, 0x8000
	SET [0xfffb], A
	ADD PC, 27
	:c4 DAT "Welcome to QBasic compiler", 0
	SET A, c4
	BOR A, 0x8000
	JSR print
	JSR printnl
	ADD PC, 5
	:c5 DAT "for ", 0
	SET A, c5
	BOR A, 0x8000
	JSR print
	SET A, [0xfffb]
	JSR print
	ADD PC, 19
	:c6 DAT " virtual computer!", 0
	SET A, c6
	BOR A, 0x8000
	JSR print
	JSR printnl
	SET A, 0x18
	SET [0xffff], A
	SET A, [0xffff]
	SET PUSH, A
	SET A, 0x2
	MUL A, POP
	SET PUSH, A
	SET A, 0x3
	ADD A, POP
	SET [0xfffe], A
	SET A, 0xa
	SET PUSH, A
	SET A, [0xffff]
	SET PUSH, A
	SET A, 0x2
	SET B, POP
	DIV B, A
	SET A, B
	ADD A, POP
	SET PUSH, A
	SET A, [0xfffe]
	ADD A, POP
	SET [0xfffd], A
	SET A, [0xffff]
	SET PUSH, A
	SET A, 0x2
	SET B, POP
	DIV B, A
	SET A, B
	SET PUSH, A
	SET A, 0x9
	SET PUSH, A
	SET A, 0x3
	SET B, POP
	SET I, A
	SET A, 1
	:l0
	MUL A, B
	SUB I, 1
	IFN I, 0
	SET PC, l0
	MUL A, POP
	SET [0xffff], A
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x2
	SET B, POP
	SET I, A
	SET A, 1
	:l1
	MUL A, B
	SUB I, 1
	IFN I, 0
	SET PC, l1
	SET PUSH, A
	SET A, 0xa
	ADD A, POP
	SET [0xfffe], A
	ADD PC, 4
	:c7 DAT "A: ", 0
	SET A, c7
	BOR A, 0x8000
	JSR print
	SET A, [0xffff]
	JSR print
	ADD PC, 6
	:c8 DAT ", B: ", 0
	SET A, c8
	BOR A, 0x8000
	JSR print
	SET A, [0xfffe]
	JSR print
	ADD PC, 6
	:c9 DAT ", C: ", 0
	SET A, c9
	BOR A, 0x8000
	JSR print
	SET A, [0xfffd]
	JSR print
	JSR printnl
	ADD PC, 32
	:c10 DAT "-------------------------------", 0
	SET A, c10
	BOR A, 0x8000
	JSR print
	JSR printnl
	ADD PC, 17
	:c11 DAT "Expected result:", 0
	SET A, c11
	BOR A, 0x8000
	JSR print
	JSR printnl
	ADD PC, 24
	:c12 DAT "A: 8748, B: 2611, C: 73", 0
	SET A, c12
	BOR A, 0x8000
	JSR print
	JSR printnl
	ADD PC, 32
	:c13 DAT "-------------------------------", 0
	SET A, c13
	BOR A, 0x8000
	JSR print
	JSR printnl
	SET A, 0x0
	SET [0xfffc], A
	:l2
	SET A, [0xfffc]
	SET PUSH, A
	SET A, 0x10
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFN C, 0
	SET PC, l3
	SET A, 0x0
	SET Y, 0
	SHL A, 12
	BOR Y, A
	SET A, [0xfffc]
	SHL A, 8
	BOR Y, A
	SET A, [0xfffc]
	SET PUSH, A
	SET A, 0xa
	SET B, POP
	MOD B, A
	SET A, B
	JSR print
	SET A, [0xfffc]
	SET PUSH, A
	SET A, 0x1
	ADD A, POP
	SET [0xfffc], A
	SET PC, l2
	:l3
	SET A, 0x0
	SET Y, 0
	SHL A, 12
	BOR Y, A
	SET A, 0x0
	SHL A, 8
	BOR Y, A
	JSR printnl
	SET A, 0x0
	SET [0xfffc], A
	:l4
	SET A, [0xfffc]
	SET PUSH, A
	SET A, 0x10
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFN C, 0
	SET PC, l5
	SET A, [0xfffc]
	SET Y, 0
	SHL A, 12
	BOR Y, A
	SET A, 0x0
	SHL A, 8
	BOR Y, A
	SET A, [0xfffc]
	SET PUSH, A
	SET A, 0xa
	SET B, POP
	MOD B, A
	SET A, B
	JSR print
	SET A, [0xfffc]
	SET PUSH, A
	SET A, 0x1
	ADD A, POP
	SET [0xfffc], A
	SET PC, l4
	:l5
	SET PC, crash
	
	; compiled functions
	:getkey
	SET A, [0x9000]
	SET [0x9000], 0
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
	:crash
	SET PC, crash
