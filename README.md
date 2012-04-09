# DCPU-Bas

DCPU-Bas is a simple QuickBASIC-like compiler for virtual DCPU in Notch's [0x10c Game](http://www.0x10c.com), written in [Go language](http://golang.org).
Compiler structure and engine is heavily inspired by [Let's Build a Compiler, by Jack Crenshaw](http://compilers.iecc.com/crenshaw/).

## Features

* Arithmetics: + - * / %
* Boolean operators: & ~ !
* Relational operators: == <> < > <= >=
* Control structures: IF, LOOP
* Variables (both integer and string)
* Statements: CLS, LOCATE, PRINT, COLOR, KEY
* Functions: STR, CHR

### IF

Usage:
	IF _condition_ THEN
		...
	[ELSE
		...]
	END IF

Executes a code block if _condition_ is met. Optional ELSE block executed if _condition_ is NOT met.

### LOOP

Usage:
	LOOP [WHILE _condition_]
		...
	END WHILE

Loops through a code block. Whe _condition_ is supplied, loops while the _condition_ is met.

### CLS

Usage:
	CLS

Clears whole 32x16 screen (video buffer at 0x8000)

### PRINT

Usage:
	PRINT _expression_ [; _expression]
	PRINT

Prints _expression(s)_ at current screen cursor location. Multiple expressions can be joined with semi-colon (;). If no expression is given, it sets cursor to column 1 of next terminal row.

### LOCATE

Usage:
	LOCATE _Y_[, _X_]

Sets current cursor location to _X_, _Y_. Set's only _Y_ if _X_ is not provided.

### COLOR

Usage:
	COLOR _FOREGROUND_, _BACKGROUND_

Sets current output color to _FOREGROUND_ and _BACKGROUND_. Both these values can be 0 to 15.

### KEY

Usage:
	KEY 

Used in an expression, it returns character code of last pressed key.

### STR

Usage:
	STR(_expression_)

Returns an ASCII character from given character code.

### CHR

Usage:
	STR(_expression_)

Returns a character code from first character of an ASCII string (opposite to STR)

### END

Program MUST end with an END statement.

## Demos

You can browse sample/ directory for sample .bas source files, as well as output/ directory for compiled .s (assembly) files of these samples.
There are couple of web-based emulators that you can try these with, most notable: [deNULL's](http://denull.ru/dcpu/dcpu.htm).

## Licence

This code is licenced on the terms of the [MIT Licence](http://www.opensource.org/licenses/mit-license.php).
