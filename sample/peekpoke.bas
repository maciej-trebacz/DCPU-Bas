DIM Word, Video

REM Get 1st word of program code (3489 in decimal at the moment)
Word = PEEK(0)

REM Print it on the screen
PRINT "10th word: "; Word

REM Access video memory directly
Video = 32768

REM Set bytes to charcode + color flag
POKE Video + 32, 72 + 2048
POKE Video + 33, 105 + 1024
POKE Video + 34, 33 + 4096

END
