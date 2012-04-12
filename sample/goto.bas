DIM Num, Guess, Count

REM Declare out variables
Num = 0
Count = 0

REM We're limited to one digit numbers, because 
REM there's no VAL function right now
PRINT "I imagined a number between"
PRINT "1 and 9. Try to guess it!"
PRINT " "

REM Loop start
start:

PRINT "Pick a number [1-9]: "; CONTINUE
Guess = CHR(INPUT) - 48
Count = Count + 1

REM Clear screen
CLS

REM Pick a random number if it's not set yet.
REM We're doing it here, because a random amount of time
REM has passed between start of program and now, because of
REM the user inputting his number
IF Num == 0 THEN
	Num = RND % 9 + 1
END IF

REM Guessed number is too high
IF Guess > Num THEN
	PRINT "Wrong! Try a lower number."
	GOTO start
END IF

REM Guessed number is too low
IF Guess < Num THEN
	PRINT "Wrong! Try a higher number."
	GOTO start
END IF

REM Good guess!
PRINT "Success! My number is "; Num; "!"
PRINT "It took you "; Count; " time(s)."

END
