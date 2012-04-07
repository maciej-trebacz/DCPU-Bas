# DCPU-Bas

DCPU-Bas is a simple QuickBASIC-like compiler for virtual DCPU in Notch's [http://www.0x10c.com](0x10c Game), written in [http://golang.org](Go).

## Features

* Arithmetics: + - * /
* Boolean operators: & ~ !
* Control structures: IF, WHILE
* Variables (integer only for now)
* Statements: CLS, LOCATE, PRINT

## Sample Program

	VAR A, B, C
	A = 5
	B = 20+A*10
	C = B - 6
	PRINT C / 2
	END

This program declares three variables, then does some math, and finally prints the result (which is 32) to video memory.
You can test the program with [http://mappum.github.com/DCPU-16/](Mappum's emulator).

### Output of sample program

		SET PUSH, A
		SET PUSH, A
	; VAR A, B, C
		SET PUSH, A
		SET PC, begin
		:print_num
		SET PUSH, A
		SET I, 0
		:pnloop1
		SET B, A
		MOD A, 0xa
		ADD A, 0x30
		SET PUSH, A
		SET A, B
		DIV A, 0xa
		ADD I, 1
		IFN A, 0
		SET PC, pnloop1
		:pnloop2
		SET A, POP
		SET B, X
		ADD B, 0x8000
		SET [B], A
		ADD X, 1
		IFN X, 0x20
		SET PC, pnline
		ADD Y, 1
		SET X, 0
		:pnline
		SUB I, 1
		IFN I, 0
		SET PC, pnloop2
		SET A, POP
		SET PC, POP
		:begin
	; A = 5
		SET A, 0x5
		SET [0xffff], A
		SET A, 0x14
		SET PUSH, A
		SET A, [0xffff]
		SET PUSH, A
	; B = 20+A*10
		SET A, 0xa
		MUL A, POP
		ADD A, POP
		SET [0xfffe], A
		SET A, [0xfffe]
		SET PUSH, A
	; C = B - 6
		SET A, 0x6
		SUB A, POP
		SET PUSH, A
		SET A, 0
		SUB A, POP
		SET [0xfffd], A
		SET A, [0xfffd]
		SET PUSH, A
	; PRINT C / 2
		SET A, 0x2
	; END
		SET X, POP
		DIV X, A
		SET A, X
		SET X, 0
		JSR print_num
		BRK

The output isn't pretty and probably performs badly. But hey, it works!
