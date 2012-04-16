/****************************************
  DCPU-Bas - QuickBasic DCPU-16 compiler
      by M4v3R (maciej@trebacz.org)

      Functions that print assembly
 ****************************************/

package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
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

func Rnd() {
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
	if Look == '"' {
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

func Font() {
	SkipWhite()
	if Look == '"' {
		Next()
		LoadFont(Value)
		Token = '$'
		Next()
	}
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

func LoadFont(filename string) {
	var error error
	var file *os.File
	var charCode = 0
	var charLine uint8 = 0
	var charWord1 uint32 = 0
	var charWord2 uint32 = 0

	EmitLine(fmt.Sprintf("; loading font: %s", filename))
	file, error = os.Open(filename)
	if error != nil {
		Error(fmt.Sprintf("Couldn't open '%s' font file!", filename))
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, _, error := reader.ReadLine()
		if error != nil {
			break
		}
		if len(line) < 4 {
			charCode, _ = strconv.Atoi(string(line))
			charWord1 = 0
			charWord2 = 0
			charLine = 0
		} else {
			if line[0] == 'O' {
				charWord1 = charWord1 | (1 << (charLine + 8))
			}
			if line[1] == 'O' {
				charWord1 = charWord1 | (1 << charLine)
			}
			if line[2] == 'O' {
				charWord2 = charWord2 | (1 << (charLine + 8))
			}
			if line[3] == 'O' {
				charWord2 = charWord2 | (1 << charLine)
			}
			charLine++
			if (charLine == 8) {
				EmitLine(fmt.Sprintf("SET [%#x], %#x", 0x8180 + charCode * 2, charWord1))
				EmitLine(fmt.Sprintf("SET [%#x], %#x", 0x8181 + charCode * 2, charWord2))
			}
		}
	}
}

func Lib() {
	var error error
	var file *os.File

	EmitLine("")
	EmitLine("; lib.dasm - compiler library")
	file, error = os.Open("lib.dasm")
	if error != nil {
		Error("Couldn't open 'lib.dasm' library file!")
	}
	reader := bufio.NewReader(file)
	bytes, _ := ioutil.ReadAll(reader)
	fmt.Printf("%s", bytes)
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
	Lib()
}
