DIM num, pad, test, big

REM Decide if whe should test bigger or smaller variable (0 or 1)
big = 1

REM With how many 0's we should pad the number
pad = 4

IF big == 1 THEN
	num = 32
ELSE
	num = 9
END IF

test = 1

LOOP WHILE pad > 1
	pad = pad - 1
	test = test * 10
END LOOP

pad = num

LOOP WHILE pad < test
	pad = pad * 10
	PRINT "0"; CONTINUE
END LOOP

PRINT num

END
