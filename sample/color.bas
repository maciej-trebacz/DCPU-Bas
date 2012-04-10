DIM X, Y

Y = 1
LOOP WHILE Y <= 16
	X = 1
	LOOP WHILE X <= 32
		COLOR X - 1, Y - 1
		LOCATE Y, X
		PRINT (X - 1 + Y - 1) % 10
		X = X + 1
	END LOOP
	Y = Y + 1
END LOOP

END
