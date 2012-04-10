DIM A, B, C, COUNTER, Text, K

REM Check simple print
PRINT "Hello"; " "; "world!"
PRINT

REM Check print with variables
Text = "0x10c"
PRINT "Welcome to QBasic compiler"
PRINT
PRINT "for "; Text; " virtual computer!"
PRINT

REM Math check
A = 24
B = A * 2 + 3
C = 10 + A / 2 + B
A = (A / 2) * 9 ^ 3
B = B ^ 2 + 10

PRINT "A: "; A; ", B: "; B ; ", C: "; C
PRINT
PRINT "-------------------------------"
PRINT
PRINT "Expected result:"
PRINT
PRINT "A: 8748, B: 2611, C: 73"
PRINT
PRINT "-------------------------------"
PRINT

REM COLOR CHECK

COUNTER = 0
LOOP WHILE COUNTER < 16
	COLOR 0, COUNTER
	PRINT COUNTER % 10
	COUNTER = COUNTER + 1
END LOOP

COLOR 0, 0
PRINT

COUNTER = 0
LOOP WHILE COUNTER < 16
	COLOR COUNTER, 0
	PRINT COUNTER % 10
	COUNTER = COUNTER + 1
END LOOP

REM KEY check

K = 0
PRINT
PRINT "-------------------------------"
PRINT
PRINT "Press any key . . ."

LOOP WHILE K == 0
	K = KEY
END LOOP

REM CLS check
CLS

PRINT "You pressed: "; STR(K)
PRINT
PRINT "Test end."

END
