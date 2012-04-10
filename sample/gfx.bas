DIM X, I

X = 1
I = 1

LOOP 
	X = I
	X = (X * (X * X * 15731 + 789221) + 1376312589) / 5293 % 3
	I = I + 1
	IF X == 2 THEN X = 8 END IF
	COLOR 1, X + 7
	PUTCHAR " "
END LOOP

END
