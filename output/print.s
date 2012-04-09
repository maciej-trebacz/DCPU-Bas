	SUB SP, 2 ; Alloc space on stack
	:begin
	SET A, 0
	:c0 DAT "HELLO", 0
	SET A, c0
	BOR A, 0x8000
	SET [0xffff], A
	SET A, 0
	:c1 DAT "WORLD", 0
	SET A, c1
	BOR A, 0x8000
	SET [0xfffe], A
	SET A, [0xffff]
	JSR print
	SET A, 0
	:c2 DAT " ", 0
	SET A, c2
	BOR A, 0x8000
	JSR print
	SET A, [0xfffe]
	JSR print
	SET A, 0
	:c3 DAT "! ", 0
	SET A, c3
	BOR A, 0x8000
	JSR print
	JSR printnl
	JSR printnl
	SET A, 0
	:c4 DAT "SUCESS!", 0
	SET A, c4
	BOR A, 0x8000
	JSR print
	JSR printnl
	SET PC, crash
	
	; compiled functions
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
