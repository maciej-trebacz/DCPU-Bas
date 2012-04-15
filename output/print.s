	ADD PC, 3
	:rnd1
	DAT 0x6769
	:rnd2
	DAT 0x1250
	:timer
	DAT 0
	SET PUSH, X
	SET PUSH, Y
	SET PUSH, Z
	SET PUSH, I
	SET PUSH, J
	SET A, SP
	SET PUSH, A
	SET Y, 0x7000
	SET Z, 0x9000
	SUB SP, 2 ; Alloc space on stack
	ADD PC, 6
	:c0 DAT "HELLO", 0
	SET A, c0
	BOR A, 0x8000
	SET [0xffff], A
	ADD PC, 6
	:c1 DAT "WORLD", 0
	SET A, c1
	BOR A, 0x8000
	SET [0xfffe], A
	SET A, [0xffff]
	JSR print
	ADD PC, 2
	:c2 DAT " ", 0
	SET A, c2
	BOR A, 0x8000
	JSR print
	SET A, [0xfffe]
	JSR print
	JSR printnl
	ADD PC, 8
	:c3 DAT "SUCESS!", 0
	SET A, c3
	BOR A, 0x8000
	JSR print
	JSR printnl
	ADD PC, 16
	:c4 DAT "The answer is: ", 0
	SET A, c4
	BOR A, 0x8000
	JSR print
	SET A, [0xffff]
	IFG 0xF000, A
	AND A, 0x7fff
	SET PUSH, [A]
	SET A, POP
	SET PUSH, A
	SET A, 0x1e
	SUB A, POP
	SET PUSH, A
	SET A, 0
	SUB A, POP
	JSR print
	JSR printnl
	ADD PC, 21
	:c5 DAT "Square root of 144: ", 0
	SET A, c5
	BOR A, 0x8000
	JSR print
	SET A, 0x90
	JSR sqrt
	JSR print
	JSR printnl
	SET J, POP
	SET I, POP
	SET Z, POP
	SET Y, POP
	SET X, POP
	SET A, POP
	SET SP, A
	SET PC, end
	
	; lib.dasm - compiler library
	; get key press
	; also increments timer (for randomization)
	; uses Z register as pointer to keyboard buffer
	; returns the key code in A
	:getkey
	ADD [timer], 1
	IFE [Z], 0
	SET PC, POP
	SET A, [Z]
	SET [Z], 0
	ADD Z, 1
	AND Z, 0x900f ; the buffer is 16 words long
	MUL [rnd1], [timer]
	ADD [rnd2], O
	SET PC, POP

	; get string length
	; A - string address in memory, with highest bit set
	;     to signal string variable type
	:strlen
	IFG 0xF000, A ; get rid of highest bit, unless it's a stack address
	AND A, 0x7fff
	SET B, A
	SET A, 0
	:strlen1
	IFN [B], 0  ; if character is 00 - end of string
	SET PC, POP ; end function
	ADD A, 1 ; increment char count
	ADD B, 1 ; increment char address
	SET PC, strlen1

    ; prints a single character on the screen
	; A - character code
	; X - location on the terminal
	; Y - color/style mask
	:printchar
	SET [0x8000+X], A ; put character to video memory + current location
	BOR [0x8000+X], Y ; apply color/style
	ADD X, 1
	IFG X, 0x1ff ; end of terminal - go back to first row and column
	SET X, 0
	SET PC, POP

	; prints an integer on the screen, converting it to string (itoa)
	; A - number to print
	:printint
	SET I, 0
	:printint1
	SET B, A
	MOD A, 0xa
	ADD A, 0x30
	SET PUSH, A
	SET A, B
	DIV A, 0xa
	ADD I, 1
	IFN A, 0
	SET PC, printint1
	:printint2
	SET A, POP
	JSR printchar
	SUB I, 1
	IFN I, 0
	SET PC, printint2
	SET A, POP
	SET PC, POP

	; prints a string on the screen
	; A - address of 0 terminated string
	:printstr
	IFG 0xF000, A ; see strlen function above
	AND A, 0x7fff
	SET I, A
	:printstr1
	IFE [I], 0
	SET PC, POP
	SET A, [I]
	JSR printchar
	ADD I, 1
	SET PC, printstr1

	; sets location register X to point at new line
	:printnl
	SHR X, 5 ; shifting 5 bits to the right means divide by 32 - terminal width
	ADD X, 1 ; next row
	SHL X, 5 ; multiply by 32 to get address of next row
	SET PC, POP

	; general purpose print function
	; determines which data type is being print and calls specific functions
	; A - address of data to be printed
	:print
	SET B, A
	SHR B, 15 ; get the highest bit
	IFE B, 0  ; if it's 0 - we have a integer
	JSR printint
	IFE B, 1  ; if it's 1 - we have a string
	JSR printstr
	SET PC, POP

	; wait for user input followed by Enter key, stores it on the stack
	; returns pointer to the string in A
	:input
	SET C, SP
	SET PUSH, 0x0
	SET I, SP
	SUB I, 1
	:input1
	SET A, 0
	JSR getkey
	IFE A, 0
	SET PC, input1
	IFE A, 0xa
	SET PC, input2
	IFE A, 0x8
	SET PC, inputbsp
	SET PUSH, A
	JSR printchar
	SET PC, input1
	:inputbsp
	SET POP, 0
	SUB X, 1
	SET [0x8000+X], 0
	BOR [0x8000+X], Y
	SET PC, input1
	:input2
	SET B, SP
	SET J, B
	:input3
	SET A, [B]
	SET [B], [I]
	SET [I], A
	ADD B, 1
	SUB I, 1
	IFG B, I
	SET PC, input4
	SET PC, input3
	:input4
	SET A, J
	BOR A, 0x8000
	SET PC, [C]

	; compare two strings to see if they're equal
	:comparestr
	SET I, POP
	ADD I, 2
	SET PUSH, I
	IFG 0xF000, A ; see strlen function
	AND A, 0x7fff
	IFG 0xF000, B
	AND B, 0x7fff
	SET I, 0
	SET C, 0
	:comparestr1
	IFN [A], [B]
	SET PC, comparestr2
	IFN [A], 0
	IFE [B], 0
	SET PC, POP
	ADD A, 1
	ADD B, 1
	SET PC, comparestr1
	:comparestr2
	SET C, 1
	SET PC, POP

	; convert string representation of a number to an integer
	; A - address of the string
	; returns resulting integer in A
	:atoi
	SET C, 0
	IFG 0xF000, A ; see strlen function
	AND A, 0x7fff
	IFE [A], 0
	SET PC, atoi2
	:atoi1
	IFG [A], 47
	IFG [A], 57
	SET PC, atoi2
	MUL C, 10
	SET B, [A]
	SUB B, 48
	ADD C, B
	ADD A, 1
	IFE [A], 0
	SET PC, atoi2
	SET PC, atoi1
	:atoi2
	SET A, C
	SET PC, POP

	; compute square root of a number
	; A - input number (16-bit integer)
	; returns the result (integer) in A
	; author: Mrrl (on the 0x10cforum.com)
	:sqrt
	SET B, 1
	SET I, A
	IFG 0x100, A
	ADD PC, 2
	SHR A, 8
	SET B, 0x10
	IFG 0x10, A
	ADD PC, 2
	SHR A, 4
	SHL B, 2
	ADD A, 4
	MUL B, A
	SHR B, 2
	SET A, I
	DIV A, B
	ADD B, A
	SHR B, 1
	SET A, I
	DIV A, B
	ADD B, A
	SHR B, 1
	SET A, I
	DIV A, B
	ADD A, B
	SHR A, 1
	SET B, A
	MUL B, A
	IFG B, I
	SUB A, 1
	SET PC, POP

	; get a pseudo-random number
	; returns the number in A
	; author: Entroper (github.com/Entroper)
	:rand
	SET B, [rnd1]
	SET A, [rnd2]
	MUL [rnd1], 0x660D
	SET C, O
	MUL A, 0x660D
	ADD A, C
	MUL B, 0x0019
	ADD A, B
	ADD [rnd1], 1
	ADD A, O
	SET [rnd2], A
	SET PC, POP

	; end of program
	; infinite loop
	:end
	SET PC, end
