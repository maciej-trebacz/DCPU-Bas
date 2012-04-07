# DCPU-Bas

DCPU-Bas is a simple QuickBASIC-like compiler for virtual DCPU in Notch's [http://www.0x10c.com](0x10c Game).

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
