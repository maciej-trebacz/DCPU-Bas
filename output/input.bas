	SUB SP, 1 ; Alloc space on stack
	SET Y, 0x7000
	ADD PC, 18
	:c0 DAT "Enter your name: ", 0
	SET A, c0
	BOR A, 0x8000
	JSR print
	SET PUSH, 0x0
	SET I, SP
	SUB I, 1
	:input
	IFE [0x9000], 0
	SET PC, input
	JSR getkey
	IFE A, 0xa
	SET PC, input2
	SET PUSH, A
	JSR printchar
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
	SET [0xffff], A
	JSR printnl
	ADD PC, 8
	:c1 DAT "Hello, ", 0
	SET A, c1
	BOR A, 0x8000
	JSR print
	SET A, [0xffff]
	JSR print
	ADD PC, 2
	:c2 DAT "!", 0
	SET A, c2
	BOR A, 0x8000
	JSR print
	SET PC, crash
	
	; compiled functions
	:getkey
	SET A, [0x9000]
	SET [0x9000], 0
	SET PC, POP
	:printchar
	SET B, X
	ADD B, 0x8000
	BOR A, Y
	SET [B], A
	ADD X, 1
	IFN X, 0x160
	SET PC, pnline
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
