	SET PUSH, X
	SET PUSH, Y
	SET PUSH, Z
	SET PUSH, I
	SET PUSH, J
	SET A, SP
	SET PUSH, A
	SET Y, 0x7000
	SUB SP, 4 ; Alloc space on stack
	SET A, 0x1
	SET [0xfffc], A
	SET A, 0x4
	SET [0xfffe], A
	SET A, [0xfffc]
	SET PUSH, A
	SET A, 0x1
	SET B, POP
	SET C, 1
	IFE A, B
	SET C, 0
	IFN C, 0
	SET PC, l0
	SET A, 0x20
	SET [0xffff], A
	SET PC, l1
	:l0
	SET A, 0x9
	SET [0xffff], A
	:l1
	SET A, 0x1
	SET [0xfffd], A
	:l2
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x1
	SET B, POP
	SET C, 1
	IFG B, A
	SET C, 0
	IFN C, 0
	SET PC, l3
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x1
	SUB A, POP
	SET PUSH, A
	SET A, 0
	SUB A, POP
	SET [0xfffe], A
	SET A, [0xfffd]
	SET PUSH, A
	SET A, 0xa
	MUL A, POP
	SET [0xfffd], A
	SET PC, l2
	:l3
	SET A, [0xffff]
	SET [0xfffe], A
	:l4
	SET A, [0xfffe]
	SET PUSH, A
	SET A, [0xfffd]
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFN C, 0
	SET PC, l5
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0xa
	MUL A, POP
	SET [0xfffe], A
	ADD PC, 2
	:c0 DAT "0", 0
	SET A, c0
	BOR A, 0x8000
	JSR print
	SET PC, l4
	:l5
	SET A, [0xffff]
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
	SET A, [0x9000]
	SET [0x9000], 0
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
	:end
	IFN SP, 0
	SET PC, POP
	:halt
	SET PC, halt
