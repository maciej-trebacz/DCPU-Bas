# DCPU-Bas

DCPU-Bas is a simple QuickBASIC-like compiler for virtual DCPU in Notch's [http://www.0x10c.com](0x10c Game), written in [http://golang.org](Go).

## Features

* Arithmetics: + - * /
* Boolean operators: & ~ !
* Control structures: IF, WHILE
* Variables (integer only for now)
* Statements: CLS, LOCATE, PRINT

## Sample Program

	VAR A, B, C, D
	A = 5
	B = 20+A*10
	C = B - 6
	D = C
	WHILE D > 0
		D = D / 10
		PRINT 0
	END WHILE
	PRINT C / 2
	END

This program declares three variables, then does some math, and finally prints the result (which is 32) to video memory.
You can test the program with [http://mappum.github.com/DCPU-16/](Mappum's emulator).
Output of the program isn't pretty, but hey, it works!

## More complex sample program

This little program gets a number, then pads it with 0's. There's a variable 'big' that let's you choose if you want
a big or small number.

	VAR num, pad, test, big

	big = 1

	pad = 4
	test = 1

	IF big == 1 THEN
		num = 32
	ELSE
		num = 9
	END IF

	WHILE pad <> 0
		pad = pad - 1
		test = test * 10
	END WHILE

	pad = num

	WHILE pad < test
		pad = pad * 10
		PRINT 0
	END WHILE

	PRINT num

	END
