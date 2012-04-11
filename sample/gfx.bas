DIM X, I

X = 1
I = 1

LOOP 
	X = (I * I + 1234)
	I = I + 1
	POKE 32768 + (I % 512), X & 65280
END LOOP

END
