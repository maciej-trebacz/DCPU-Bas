DIM num, pad, test, big

big = 1

pad = 4
test = 1

IF big == 1 THEN
	num = 32
ELSE
	num = 9
END IF

WHILE pad > 1
	pad = pad - 1
	test = test * 10
END WHILE

pad = num

WHILE pad < test
	pad = pad * 10
	PRINT 0
END WHILE

PRINT num

END
