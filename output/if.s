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
	SUB SP, 5 ; Alloc space on stack
	SET A, 0x1
	SET [0xffff], A
	SET A, 0x2
	SET [0xfffe], A
	ADD PC, 6
	:c0 DAT "HELLO", 0
	SET A, c0
	BOR A, 0x8000
	SET [0xfffd], A
	ADD PC, 6
	:c1 DAT "WORLD", 0
	SET A, c1
	BOR A, 0x8000
	SET [0xfffc], A
	ADD PC, 4
	:c2 DAT "A: ", 0
	SET A, c2
	BOR A, 0x8000
	JSR print
	SET A, [0xffff]
	JSR print
	JSR printnl
	ADD PC, 4
	:c3 DAT "B: ", 0
	SET A, c3
	BOR A, 0x8000
	JSR print
	SET A, [0xfffe]
	JSR print
	JSR printnl
	ADD PC, 4
	:c4 DAT "C: ", 0
	SET A, c4
	BOR A, 0x8000
	JSR print
	SET A, [0xfffd]
	JSR print
	JSR printnl
	ADD PC, 4
	:c5 DAT "D: ", 0
	SET A, c5
	BOR A, 0x8000
	JSR print
	SET A, [0xfffc]
	JSR print
	JSR printnl
	ADD PC, 2
	:c6 DAT " ", 0
	SET A, c6
	BOR A, 0x8000
	JSR print
	JSR printnl
	SET A, 0x1
	SET [0xfffb], A
	:l0
	SET A, [0xfffb]
	SET PUSH, A
	SET A, 0x9
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFN C, 0
	SET PC, l1
	SET A, [0xfffb]
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0xa
	SUB A, 1
	ADD X, A
	ADD PC, 2
	:c7 DAT "|", 0
	SET A, c7
	BOR A, 0x8000
	JSR print
	JSR printnl
	SET A, [0xfffb]
	SET PUSH, A
	SET A, 0x1
	ADD A, POP
	SET [0xfffb], A
	SET PC, l0
	:l1
	SET A, [0xffff]
	SET PUSH, A
	SET A, [0xfffe]
	SET B, POP
	SET C, 1
	IFG 0x8000, A
	IFG B, 0x7fff
	JSR comparestr
	IFE A, B
	SET C, 0
	IFN C, 0
	SET PC, l2
	ADD PC, 7
	:c8 DAT "A == B", 0
	SET A, c8
	BOR A, 0x8000
	JSR print
	JSR printnl
	:l2
	SET A, 0x1
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0xc
	SUB A, 1
	ADD X, A
	SET A, [0xffff]
	SET PUSH, A
	SET A, [0xfffe]
	SET B, POP
	SET C, 1
	IFG 0x8000, A
	IFG B, 0x7fff
	JSR comparestr
	IFN A, B
	SET C, 0
	IFN C, 0
	SET PC, l3
	ADD PC, 7
	:c9 DAT "A <> B", 0
	SET A, c9
	BOR A, 0x8000
	JSR print
	JSR printnl
	:l3
	SET A, 0x2
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0xc
	SUB A, 1
	ADD X, A
	SET A, [0xffff]
	SET PUSH, A
	SET A, [0xfffe]
	SET B, POP
	SET C, 1
	IFG B, A
	SET C, 0
	IFN C, 0
	SET PC, l4
	ADD PC, 6
	:c10 DAT "A > B", 0
	SET A, c10
	BOR A, 0x8000
	JSR print
	JSR printnl
	:l4
	SET A, 0x3
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0xc
	SUB A, 1
	ADD X, A
	SET A, [0xffff]
	SET PUSH, A
	SET A, [0xfffe]
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFN C, 0
	SET PC, l5
	ADD PC, 6
	:c11 DAT "A < B", 0
	SET A, c11
	BOR A, 0x8000
	JSR print
	JSR printnl
	:l5
	SET A, 0x4
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0xc
	SUB A, 1
	ADD X, A
	SET A, [0xffff]
	SET PUSH, A
	SET A, [0xfffe]
	SET B, POP
	SET C, 1
	IFG B, A
	SET C, 0
	IFG 0x8000, A
	IFG B, 0x7fff
	JSR comparestr
	IFE A, B
	SET C, 0
	IFN C, 0
	SET PC, l6
	ADD PC, 7
	:c12 DAT "A >= B", 0
	SET A, c12
	BOR A, 0x8000
	JSR print
	JSR printnl
	:l6
	SET A, 0x5
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0xc
	SUB A, 1
	ADD X, A
	SET A, [0xffff]
	SET PUSH, A
	SET A, [0xfffe]
	SET B, POP
	SET C, 1
	IFG A, B
	SET C, 0
	IFG 0x8000, A
	IFG B, 0x7fff
	JSR comparestr
	IFE A, B
	SET C, 0
	IFN C, 0
	SET PC, l7
	ADD PC, 7
	:c13 DAT "A <= B", 0
	SET A, c13
	BOR A, 0x8000
	JSR print
	JSR printnl
	:l7
	SET A, 0x6
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0xc
	SUB A, 1
	ADD X, A
	SET A, [0xfffd]
	SET PUSH, A
	SET A, [0xfffc]
	SET B, POP
	SET C, 1
	IFG 0x8000, A
	IFG B, 0x7fff
	JSR comparestr
	IFE A, B
	SET C, 0
	IFN C, 0
	SET PC, l8
	ADD PC, 7
	:c14 DAT "C == D", 0
	SET A, c14
	BOR A, 0x8000
	JSR print
	JSR printnl
	:l8
	SET A, 0x7
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0xc
	SUB A, 1
	ADD X, A
	SET A, [0xfffd]
	SET PUSH, A
	ADD PC, 6
	:c15 DAT "HELLO", 0
	SET A, c15
	BOR A, 0x8000
	SET B, POP
	SET C, 1
	IFG 0x8000, A
	IFG B, 0x7fff
	JSR comparestr
	IFE A, B
	SET C, 0
	IFN C, 0
	SET PC, l9
	ADD PC, 13
	:c16 DAT "C == 'HELLO'", 0
	SET A, c16
	BOR A, 0x8000
	JSR print
	JSR printnl
	:l9
	SET A, 0x8
	SUB A, 1
	SET PUSH, 0x20
	MUL A, POP
	SET X, A
	SET A, 0xc
	SUB A, 1
	ADD X, A
	SET A, [0xfffc]
	SET PUSH, A
	ADD PC, 6
	:c17 DAT "HELLO", 0
	SET A, c17
	BOR A, 0x8000
	SET B, POP
	SET C, 1
	IFG 0x8000, A
	IFG B, 0x7fff
	JSR comparestr
	IFE A, B
	SET C, 0
	IFN C, 0
	SET PC, l10
	ADD PC, 13
	:c18 DAT "D == 'HELLO'", 0
	SET A, c18
	BOR A, 0x8000
	JSR print
	JSR printnl
	:l10
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
	:input
	SET C, SP
	SET PUSH, 0x0
	SET I, SP
	SUB I, 1
	:input1
	SET A, 0
	JSR getkey
	IFE A, 0
	SET PC, input1
	IFE A, 0xa
	SET PC, input2
	IFE A, 0x8
	SET PC, inputbsp
	SET PUSH, A
	JSR printchar
	SET PC, input1
	:inputbsp
	SET POP, 0
	SUB X, 1
	SET [0x8000+X], 0
	BOR [0x8000+X], Y
	SET PC, input1
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
	SET PC, [C]
	:comparestr
	SET I, POP
	ADD I, 2
	SET PUSH, I
	IFG 0xF000, A
	AND A, 0x7fff
	IFG 0xF000, B
	AND B, 0x7fff
	SET I, 0
	SET C, 0
	:comparestr1
	IFN [A], [B]
	SET PC, comparestr2
	IFN [A], 0
	IFE [B], 0
	SET PC, POP
	ADD A, 1
	ADD B, 1
	SET PC, comparestr1
	:comparestr2
	SET C, 1
	SET PC, POP
	:atoi
	IFE [A], 0
	SET PC, atoi2
	SET C, 0
	:atoi1
	IFG [A], 47
	IFG [A], 57
	SET PC, atoi2
	MUL C, 10
	SET B, [A]
	SUB B, 48
	ADD C, B
	ADD A, 1
	IFE [A], 0
	SET PC, atoi2
	SET PC, atoi1
	:atoi2
	SET A, C
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
