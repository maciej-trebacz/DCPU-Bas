	SUB SP, 2 ; Alloc space on stack
	:begin
	SET A, 0x1
	SET [0xfffe], A
	:l0
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0xc
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFE A, B
	SET C, 0
	IFN C, 0
	SET PC, l1
	SET A, 0x1
	SET [0xffff], A
	:l2
	SET A, [0xffff]
	SET PUSH, A
	SET A, 0x20
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFE A, B
	SET C, 0
	IFN C, 0
	SET PC, l3
	SET A, [0xfffe]
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, [0xffff]
	SUB A, 1
	ADD X, A
	SET A, [0xffff]
	SET PUSH, A
	SET A, 0x10
	SET B, POP
	SET C, 1
	IFG B, A
	SET C, 0
	IFN C, 0
	SET PC, l4
	SET A, [0xffff]
	SET Y, 0
	SHL A, 12
	BOR Y, A
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0xc
	ADD A, POP
	SHL A, 8
	BOR Y, A
	SET PC, l5
	:l4
	SET A, [0xffff]
	SET Y, 0
	SHL A, 12
	BOR Y, A
	SET A, [0xfffe]
	SHL A, 8
	BOR Y, A
	:l5
	SET A, [0xffff]
	SET PUSH, A
	SET A, 0xa
	SET B, POP
	MOD B, A
	SET A, B
	SET B, 0
	JSR print
	JSR printnl
	SET A, [0xffff]
	SET PUSH, A
	SET A, 0x1
	ADD A, POP
	SET [0xffff], A
	SET PC, l2
	:l3
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x1
	ADD A, POP
	SET [0xfffe], A
	SET PC, l0
	:l1
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
