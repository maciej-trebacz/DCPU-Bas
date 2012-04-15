# DCPU-Bas - 0xBASIC language for DCPU-16

DCPU-Bas is a simple QuickBASIC-like compiler for virtual DCPU in Notch's [0x10c Game](http://www.0x10c.com), written in [Go language](http://golang.org).  
Compiler structure and engine is heavily inspired by [Let's Build a Compiler, by Jack Crenshaw](http://compilers.iecc.com/crenshaw/).

## Features

* Arithmetics: + - * / %
* Boolean operators: & ~ !
* Bit shifts: << >>
* Relational operators: == <> < > <= >=
* Control structures: IF, LOOP
* Variables (both integer and string)
* Statements: CLS, LOCATE, PRINT, COLOR, KEY, INPUT, POKE, GOTO, RND
* Functions: STR, CHR, VAL, PEEK, SQR

## How the language looks like

Here's a sample program (you can find it in samples/input.bas) that asks for your name and then displays it back to you:

```
DIM MyName

PRINT "Enter your name: "; CONTINUE
MyName = INPUT

PRINT " "
PRINT "Hello, "; MyName; "!"

END
```

## How to get it.

Easiest way is to get the binaries from the [downloads section](https://github.com/M4v3R/DCPU-Bas/downloads).  
You can also build it from the latest sources. In that case, you need to:

* Get and setup [Go](http://golang.org/doc/install)
* Get the [latest sources](https://github.com/M4v3R/DCPU-Bas/zipball/master) and unpack them to a directory
* On the command line, within that directory, type: ```go build```

That should do it.

## Language documentation

Below are language statements and functions explained:

### IF

Usage:

```
IF condition THEN
	...
[ELSE
	...]
END IF
```

Executes a code block if _condition_ is met. Optional ELSE block executed if _condition_ is NOT met.

### LOOP

Usage:

```
LOOP [WHILE condition]
	...
END WHILE
```

Loops through a code block. Whe _condition_ is supplied, loops while the _condition_ is met.

### CLS

Usage:

```
CLS
```

Clears whole 32x16 screen (video buffer at 0x8000)

### PRINT

Usage:

```
PRINT expression [; expression][; CONTINUE]
PRINT
```

Prints _expression(s)_ at current screen cursor location. Multiple expressions can be joined with semi-colon (;).  
After printing all expressions cursor position will be set to next row, column 1, unless CONTINUE keyword is given at the end of expression list.  
Example:

```
A = "World"
PRINT "Hello "; A
PRINT "A sentence within "; CONTINUE
PRINT "the same line."
```

### PUTCHAR

Usage:

```
PUTCHAR "c"
PUTCHAR expression
```

Fast way (5x faster than PRINT "c") to output a single character to video buffer at current cursor position. It can be either a string literal _c_
in double quotes, or a math expression (which could be a single number) returning character ascii code.

### LOCATE

Usage:

```
LOCATE Y[, X]
```

Sets current cursor location to _X_, _Y_. Set's only _Y_ if _X_ is not provided.

### COLOR

Usage:

```
COLOR FOREGROUND, BACKGROUND
```

Sets current output color to _FOREGROUND_ and _BACKGROUND_. Both these values can be 0 to 15.

### KEY

Usage:

```
DIM Code
Code = KEY 
```

Used in an expression, it returns character code of last pressed key.

### STR

Usage:

```
DIM Char
Char = STR(expression)
```

Returns an ASCII character from given character code.

### CHR

Usage:

```
DIM Code
Code = STR(expression)
```

Returns a character code from first character of an ASCII string (opposite to STR)

### VAL

Usage:

```
DIM Number
Number = VAL(expression)
```

Returns an integer representation of a number in string variable (e.g. input from user)

### INPUT

Usage:

```
DIM YourName
YourName = INPUT
```

Waits for user to enter a string followed by ENTER key, and returns this string as expression. User input is displayed on the screen.

### PEEK

Usage:

```
DIM MemoryValue
MemoryValue = PEEK(address)
```

Reads directly the memory and returns a number representing word at given memory _address_.

### POKE

Usage:

```
POKE address, value
```

Writes directly to memory, sets word at _address_ to given _value_.

### GOTO

Usage:

```
:label
GOTO label
```

An unconditional jump into another place in the program code, defined by a _label_.  
Label definitions must start with a colon (:).  
_Use it with caution_, there are no checks performed, so you can break your code easily by abusing this statement.
For example, nothing will prevent you to jump right into a LOOP or IF structure, but it can have unpredictable results.

### RND

Usage:

```
DIM RandomValue
RandomValue = RND
```

Generates a pseudo-random 16-bit integer (from 0 to 65535).  
To get a number from a specific range, use modulo division: ```RandomValue = RND % 100 + 1``` gives you a number from 1 to 100.  
*Note:* Every program always starts with the same seed, meaning that it will have a same set of random numbers generated. 
To update the seed randomly, use user's keyboard input - ask for something using INPUT statement, or wait for a keypress using KEY statement.
These statements update the generator seed by multiplying it with amount of time that has passed before a keystroke, which should be random enough.
Credit: Linear congruentrial generator used by this statement is coded by [Entroper](https://github.com/Entroper/DCPU-16-PRNGs).

### SQR

Usage:

```
DIM Result
Result = SQR(expression)
```

Computes square root of _expression_.

### END

Program MUST end with an ```END``` statement.

## Demos

You can browse sample/ directory for sample .bas source files, as well as output/ directory for compiled .s (assembly) files of these samples.
There are couple of web-based emulators that you can try these with, most notable: [deNULL's](http://denull.ru/dcpu/dcpu.htm).

## Licence

This code is licenced on the terms of the [MIT Licence](http://www.opensource.org/licenses/mit-license.php).
