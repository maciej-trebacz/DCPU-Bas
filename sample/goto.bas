DIM Num, Guess, Count, GetKey

REM Declare out variables
Num = 0
Count = 0

begin:
PRINT "I imagined a number between"
PRINT "1 and 100. Try to guess it!"
PRINT " "

REM Loop start
start:

PRINT "Pick a number [1-100]: "; CONTINUE
Guess = VAL(INPUT)
Count = Count + 1

REM Clear screen
CLS

REM Pick a random number if it's not set yet.
REM We're doing it here, because a random amount of time
REM has passed between start of program and now, because of
REM the user inputting his number
IF Num == 0 THEN
	Num = RND % 100 + 1
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
PRINT " "

REM Again?
PRINT "Do you want to play again? (y/n)"
GetKey = 0
LOOP WHILE GetKey == 0
GetKey = KEY
END LOOP

REM 121 = 'y'
IF GetKey == 121 THEN
	Num = 0
	CLS
	GOTO begin
END IF

Print "Thanks for playing!"

REM End program
END
