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

func PopCompare() {
	StackDepth--
	EmitLine("SET B, POP")
	EmitLine("SET C, 1")
}

func SetEqual() {
	EmitLine("IFE A, B")
	EmitLine("SET C, 0")
}

func SetNotEqual() {
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
	EmitLine("SET PUSH, X")
	EmitLine("SET PUSH, Y")
	EmitLine("SET PUSH, Z")
	EmitLine("SET PUSH, I")
	EmitLine("SET PUSH, J")
	EmitLine("SET A, SP")
	EmitLine("SET PUSH, A")
	EmitLine("SET Y, 0x7000") // Set color to white
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

func Key() {
	Next()
}

func Input() {
	EmitLine("SET PUSH, 0x0")
	EmitLine("SET I, SP")
	EmitLine("SUB I, 1")
	PostLabel("input")
	EmitLine("IFE [0x9000], 0")
	EmitLine("SET PC, input")
	Call("getkey")
	EmitLine("IFE A, 0xa")
	EmitLine("SET PC, input2")
	EmitLine("SET PUSH, A")
	Call("printchar")
	EmitLine("SET PC, input")
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
}

func Lib() {
	PostLabel("getkey") // Get key press
	EmitLine("SET A, [0x9000]")
	EmitLine("SET [0x9000], 0")
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
	EmitLine("IFN SP, 0")
	EmitLine("SET PC, POP")
	PostLabel("halt")
	EmitLine("SET PC, halt")
}
