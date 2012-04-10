DIM X, Y, COL, TICK

X = 1
Y = X
X = X + Y * 2
COL = X - Y + 2
TICK = COL - Y / 1

LOOP
	LOOP
		COL = (X * Y + Y * Y + TICK) + 1
		COLOR 1, COL
		PRINT " "
		X = X + 1
		TICK = TICK + 1
		IF X > 32 THEN
			X = 1
			Y = Y + 1
			IF Y > 12 THEN
				Y = 1
			END IF
		END IF
	END LOOP
END LOOP

END
