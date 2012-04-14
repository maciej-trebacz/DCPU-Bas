/****************************************
  DCPU-Bas - QuickBasic DCPU-16 compiler
      by M4v3R (maciej@trebacz.org)

      Functions that print assembly
 ****************************************/

package main

import (
	"fmt"
	"strconv"
)

func Negate() {
	Push()
	EmitLine("SET A, 0")
	EmitLine("SUB A, POP")
}

func Clear() {
	EmitLine("SET A, 0")
}

func Not() {
	EmitLine("XOR A, -1")
}

func LoadConst(s string) {
	val, _ := strconv.Atoi(s)
	EmitLine(fmt.Sprintf("SET A, %#x", val))
}

func LoadConstString(s string) {
	label := NewConst()
	EmitLine(fmt.Sprintf("ADD PC, %d", len(s) + 1))
	EmitLine(fmt.Sprintf(":%s DAT \"%s\", 0", label, s))
	EmitLine(fmt.Sprintf("SET A, %s", label))
	EmitLine(fmt.Sprintf("BOR A, 0x8000"))
}

func LoadVar(s string) {
	if !InTable(s) {
		Undefined(s)
	}
	symbol := Symbols[Locate(GetSymbol(s))]
	EmitLine(fmt.Sprintf("SET A, [%#x]", (0xffff + symbol.l)))
}

func Push() {
	StackDepth++
	EmitLine("SET PUSH, A")
}

func PopAdd() {
	StackDepth--
	EmitLine("ADD A, POP")
}

func PopSub() {
	StackDepth--
	EmitLine("SUB A, POP")
	Negate()
}

func PopMul() {
	StackDepth--
	EmitLine("MUL A, POP")
}

func PopDiv() {
	StackDepth--
	EmitLine("SET B, POP")
	EmitLine("DIV B, A")
	EmitLine("SET A, B")
}

func PopMod() {
	StackDepth--
	EmitLine("SET B, POP")
	EmitLine("MOD B, A")
	EmitLine("SET A, B")
}

func PopPow() {
	StackDepth--
	l := NewLabel()
	EmitLine("SET B, POP")
	EmitLine("SET I, A")
	EmitLine("SET A, 1")
	PostLabel(l)
	EmitLine("MUL A, B")
	EmitLine("SUB I, 1")
	EmitLine("IFN I, 0")
	Branch(l)
}

func PopAnd() {
	StackDepth--
	EmitLine("AND A, POP")
}

func PopOr() {
	StackDepth--
	EmitLine("BOR A, POP")
}

func PopXor() {
	StackDepth--
	EmitLine("XOR A, POP")
}

func PopShl() {
	StackDepth--
	EmitLine("SET B, POP")
	EmitLine("SHL B, A")
	EmitLine("SET A, B")
}

func PopShr() {
	StackDepth--
	EmitLine("SET B, POP")
	EmitLine("SHR B, A")
	EmitLine("SET A, B")
}

func PopCompare() {
	StackDepth--
	EmitLine("SET B, POP")
	EmitLine("SET C, 1")
}

func SetEqual() {
	EmitLine("IFG 0x8000, A")
	EmitLine("IFG B, 0x7fff")
	Call("comparestr")
	EmitLine("IFE A, B")
	EmitLine("SET C, 0")
}

func SetNotEqual() {
	EmitLine("IFG 0x8000, A")
	EmitLine("IFG B, 0x7fff")
	Call("comparestr")
	EmitLine("IFN A, B")
	EmitLine("SET C, 0")
}

func SetGreater() {
	EmitLine("IFG B, A")
	EmitLine("SET C, 0")
}

func SetLess() {
	EmitLine("IFG A, B")
	EmitLine("SET C, 0")
}

func SetGreaterOrEqual() {
	SetGreater()
	SetEqual()
}

func SetLessOrEqual() {
	SetLess()
	SetEqual()
}

func Store(s string) {
	symbol := Symbols[Locate(GetSymbol(s))]
	EmitLine(fmt.Sprintf("SET [%#x], A", (0xffff + symbol.l)))
}

func Call(s string) {
	EmitLine(fmt.Sprintf("JSR %s", s))
}

func Branch(s string) {
	EmitLine(fmt.Sprintf("SET PC, %s", s))
}

func BranchFalse(s string) {
	EmitLine("IFN C, 0")
	Branch(s)
}

func Prolog() {
	EmitLine("ADD PC, 3")
	EmitLine(":rnd1")
	EmitLine("DAT 0x6769")
	EmitLine(":rnd2")
	EmitLine("DAT 0x1250")
	EmitLine(":timer")
	EmitLine("DAT 0")
	EmitLine("SET PUSH, X")
	EmitLine("SET PUSH, Y")
	EmitLine("SET PUSH, Z")
	EmitLine("SET PUSH, I")
	EmitLine("SET PUSH, J")
	EmitLine("SET A, SP")
	EmitLine("SET PUSH, A")
	EmitLine("SET Y, 0x7000") // Set color to white
	EmitLine("SET Z, 0x9000") // Key pointer
}

func Rnd() { // Credits for this function go to Entroper (github.com/Entroper)
	Call("rand")
}

func Ret() {
	EmitLine("SET PC, POP")
}

func Cls() {
	EmitLine("SET I, 0x8220")
	l := NewLabel()
	PostLabel(l)
	EmitLine("SUB I, 1")
	EmitLine("SET [I], 0")
	EmitLine("IFN I, 0x8000")
	Branch(l)
	EmitLine("SET X, 0")
	Next()
}

func Loc() {
	Next()
	BoolExpression()
	EmitLine("SUB A, 1")
	EmitLine("SET PUSH, 0x20")
	EmitLine("MUL A, POP")
	EmitLine("SET X, A")
	if Token == ',' {
		Next()
		BoolExpression()
		EmitLine("SUB A, 1")
		EmitLine("ADD X, A")
	}
}

func Color() {
	Next()
	BoolExpression()
	EmitLine("SET Y, 0")
	EmitLine("SHL A, 12")
	EmitLine("BOR Y, A")
	Next()
	BoolExpression()
	EmitLine("SHL A, 8")
	EmitLine("BOR Y, A")
}

func Poke() {
	Next()
	BoolExpression()
	EmitLine("SET B, A")
	Next()
	BoolExpression()
	EmitLine("SET [B], A")
}

func PutChar() {
	SkipWhite()
	if (Look == '"') {
		GetChar()
		EmitLine(fmt.Sprintf("SET [0x8000+X], %d", Look))
		GetChar()
		GetChar()
		Token = '$'
		Next()
	} else {
		Next()
		Factor()
		EmitLine("SET [0x8000+X], A")
	}
	EmitLine("BOR [0x8000+X], Y")
	EmitLine("ADD X, 1")
	EmitLine("AND X, 0x1ff")
}

func Goto() {
	Next()
	Token = '$'
	Branch(Value)
	Next()
}

func Input() {
	Call("input")
}

func Lib() {
	PostLabel("getkey") // Get key press
	EmitLine("ADD [timer], 1") // Increase "timer"
	EmitLine("IFE [Z], 0")
	Ret()
	EmitLine("SET A, [Z]") // Get key code
	EmitLine("SET [Z], 0")
	EmitLine("ADD Z, 1")
	EmitLine("AND Z, 0x900f") // Make sure the pointer is circular
	EmitLine("MUL [rnd1], [timer]")
	EmitLine("ADD [rnd2], O")
	Ret()

	PostLabel("strlen") // gets string length
	EmitLine("SET I, A")
	PostLabel("strlen1")
	EmitLine("ADD I, 1")
	EmitLine("IFN [I], 0x0")
	Branch("strlen1")
	EmitLine("SET A, B")
	Ret()

	PostLabel("printchar") // Print char
	EmitLine("SET [0x8000+X], A") // Print char at cursor position
	EmitLine("BOR [0x8000+X], Y") // Apply color style
	EmitLine("ADD X, 1") // Increment cursor position
	EmitLine("IFG X, 0x21f") // Check if we should do next line (X > 32)
	EmitLine("SET X, 0") // First row, first column
	PostLabel("pnline")
	Ret()

	PostLabel("printint") // Print integer
	EmitLine("SET I, 0") // Loop counter
	PostLabel("printint1") // Loop: divide A by 10 until 0 is left
	EmitLine("SET B, A") // Store A (number) for later
	EmitLine("MOD A, 0xa") // Get remainder from division by 10
	EmitLine("ADD A, 0x30") // Add 0x30 to the remainder to get ASCII code
	EmitLine("SET PUSH, A") // Store the remainder (digit) on the stack
	EmitLine("SET A, B") // Get A (number) back
	EmitLine("DIV A, 0xa") // Divide the number by 10
	EmitLine("ADD I, 1") // Increment loop counter
	EmitLine("IFN A, 0") // A > 10: jump
	EmitLine("SET PC, printint1")
	PostLabel("printint2") // Loop: print character by character
	EmitLine("SET A, POP") // Get digit from stack
	EmitLine("JSR printchar") // Print character
	EmitLine("SUB I, 1") // Decrement loop counter
	EmitLine("IFN I, 0")
	EmitLine("SET PC, printint2") // Jump back if there are more chars
	EmitLine("SET A, POP")
	Ret()

	PostLabel("printstr") // Print string
	EmitLine("IFG 0xF000, A") // Check if it's not a stack pointer
	EmitLine("AND A, 0x7fff")
	EmitLine("SET I, A") // Get string address
	PostLabel("printstr1")
	EmitLine("IFE [I], 0") // Return if we've reached end of string
	Ret()
	EmitLine("SET A, [I]") // Set A to address of next char
	EmitLine("JSR printchar") // Print char
	EmitLine("ADD I, 1") // Increment char index
	EmitLine("SET PC, printstr1") // Loop

	PostLabel("printnl") // Print new line
	EmitLine("DIV X, 32")
	EmitLine("ADD X, 1")
	EmitLine("MUL X, 32")
	Ret()

	PostLabel("print")
	EmitLine("SET B, A") // Check variable type
	EmitLine("SHR B, 15")
	EmitLine("IFE B, 0") // Integer
	EmitLine("JSR printint")
	EmitLine("IFE B, 1") // String
	EmitLine("JSR printstr")
	Ret()

	PostLabel("input")
	EmitLine("SET C, SP")
	EmitLine("SET PUSH, 0x0")
	EmitLine("SET I, SP")
	EmitLine("SUB I, 1")
	PostLabel("input1")
	EmitLine("SET A, 0")
	Call("getkey")
	EmitLine("IFE A, 0")
	EmitLine("SET PC, input1")
	EmitLine("IFE A, 0xa")
	EmitLine("SET PC, input2")
	EmitLine("IFE A, 0x8")
	EmitLine("SET PC, inputbsp")
	EmitLine("SET PUSH, A")
	Call("printchar")
	EmitLine("SET PC, input1")
	PostLabel("inputbsp")
	EmitLine("SET POP, 0")
	EmitLine("SUB X, 1")
	EmitLine("SET [0x8000+X], 0")
	EmitLine("BOR [0x8000+X], Y")
	EmitLine("SET PC, input1")
	PostLabel("input2")
	EmitLine("SET B, SP")
	EmitLine("SET J, B")
	PostLabel("input3")
	EmitLine("SET A, [B]")
	EmitLine("SET [B], [I]")
	EmitLine("SET [I], A")
	EmitLine("ADD B, 1")
	EmitLine("SUB I, 1")
	EmitLine("IFG B, I")
	EmitLine("SET PC, input4")
	EmitLine("SET PC, input3")
	PostLabel("input4")
	EmitLine("SET A, J")
	EmitLine("BOR A, 0x8000")
	EmitLine("SET PC, [C]")
	PostLabel("comparestr")
	EmitLine("SET I, POP") // adjust return address to bypass int cmp
	EmitLine("ADD I, 2")
	EmitLine("SET PUSH, I")
	EmitLine("IFG 0xF000, A")
	EmitLine("AND A, 0x7fff")
	EmitLine("IFG 0xF000, B")
	EmitLine("AND B, 0x7fff")
	EmitLine("SET I, 0")
	EmitLine("SET C, 0")
	PostLabel("comparestr1")
	EmitLine("IFN [A], [B]")
	Branch("comparestr2")
	EmitLine("IFN [A], 0")
	EmitLine("IFE [B], 0")
	EmitLine("SET PC, POP")
	EmitLine("ADD A, 1")
	EmitLine("ADD B, 1")
	Branch("comparestr1")
    PostLabel("comparestr2")
	EmitLine("SET C, 1")
    Ret()

	PostLabel("atoi")
	EmitLine("IFE [A], 0")
	Branch("atoi2")
	EmitLine("SET C, 0")
	PostLabel("atoi1")
	EmitLine("IFG [A], 47") // Check if character is a digit
	EmitLine("IFG [A], 57")
	Branch("atoi2")
	EmitLine("MUL C, 10")
	EmitLine("SET B, [A]")
	EmitLine("SUB B, 48")
	EmitLine("ADD C, B")
	EmitLine("ADD A, 1")
	EmitLine("IFE [A], 0")
	Branch("atoi2")
	Branch("atoi1")
	PostLabel("atoi2")
	EmitLine("SET A, C")
	Ret()

	PostLabel("rand")
	EmitLine("SET B, [rnd1]")
	EmitLine("SET A, [rnd2]")
	EmitLine("MUL [rnd1], 0x660D")
	EmitLine("SET C, O")
	EmitLine("MUL A, 0x660D")
	EmitLine("ADD A, C")
	EmitLine("MUL B, 0x0019")
	EmitLine("ADD A, B")
	EmitLine("ADD [rnd1], 1")
	EmitLine("ADD A, O")
	EmitLine("SET [rnd2], A")
	Ret()
}

func Epilog() {
	EmitLine("SET J, POP")
	EmitLine("SET I, POP")
	EmitLine("SET Z, POP")
	EmitLine("SET Y, POP")
	EmitLine("SET X, POP")
	EmitLine("SET A, POP")
	EmitLine("SET SP, A")
	EmitLine("SET PC, end")
	EmitLine("")
	EmitLine("; compiled functions")
	Lib()
	PostLabel("end")
//	EmitLine("IFN SP, 0")
//	EmitLine("SET PC, POP")
//	PostLabel("halt")
	EmitLine("SET PC, end")
}
