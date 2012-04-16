DIM I, Hello, TermWidth

REM This example requires an emulator that supports custom fonts!
REM Check http://0x10co.de for one that does.

REM Set up variables and colors
Hello = "Hello from 0xBASIC!"
TermWidth = 32
I = 0
COLOR 14, 0 

REM Load font from file
FONT "fonts/box.txt"

REM Top row
PUTCHAR 9
LOOP WHILE I < 30
PUTCHAR 13
I = I + 1
END LOOP
PUTCHAR 7

REM Middle row
PUTCHAR 15
LOCATE 2, 32
PUTCHAR 15

REM bottom row
I = 0
PUTCHAR 8
LOOP WHILE I < 30
PUTCHAR 13
I = I + 1
END LOOP
PUTCHAR 6

REM Print text in the middle of the second line
LOCATE 2, TermWidth / 2 - LEN(Hello) / 2
COLOR 15, 0
PRINT Hello

END

