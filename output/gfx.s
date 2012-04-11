	SET PUSH, X
	SET PUSH, Y
	SET PUSH, Z
	SET PUSH, I
	SET PUSH, J
	SET A, SP
	SET PUSH, A
	SET Y, 0x7000
	SET Z, 0x9000
	SUB SP, 2 ; Alloc space on stack
	SET A, 0x1
	SET [0xffff], A
	SET A, 0x1
	SET [0xfffe], A
	:l0
	SET A, [0xfffe]
	SET PUSH, A
	SET A, [0xfffe]
	MUL A, POP
	SET PUSH, A
	SET A, 0x4d2
	ADD A, POP
	SET [0xffff], A
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x1
	ADD A, POP
	SET [0xfffe], A
	SET A, 0x8000
	SET PUSH, A
	SET A, [0xfffe]
	SET PUSH, A
	SET A, 0x200
	SET B, POP
	MOD B, A
	SET A, B
	ADD A, POP
	SET B, A
	SET A, [0xffff]
	SET PUSH, A
	SET A, 0xff00
	AND A, POP
	SET [B], A
	SET PC, l0
	:l1
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
	IFE [Z], 0
	SET PC, POP
	SET A, [Z]
	SET [Z], 0
	ADD Z, 1
	AND Z, 0x900f
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
	SET PC, end
