/****************************************
  DCPU-Bas - QuickBasic DCPU-16 compiler
      by M4v3R (maciej@trebacz.org)

            Main code parser
 ****************************************/

package main

import (
	"fmt"
	"os"
	"bytes"
	"strings"
)

var data *os.File
var Look, Prev byte
var Keywords = []string { "IF", "ELSE",  "LOOP", "END", "DIM", "CLS", "PRINT", "LOCATE", "REM", "COLOR", "POKE", "PUTCHAR", "GOTO" }
var Tokens = []byte { 'x', 'i', 'l', 'w', 'e', 'd', 'c', 'p', 'o', 'r', 'k', 'q', 'u', 'g', 'n' }
var Token byte
var Value string
var LabelCount = 0
var ConstCount = 0
var StackDepth = 0
var Line = bytes.NewBufferString("")

type Symbol struct {
	n string
	t rune
	l int
}

var Symbols = make([]Symbol, 100)

func Abort(errorString string) {
	Error(errorString)
	fmt.Printf("\nCurrent token type: %c, look: %c, value: '%s'\n\n", Token, Look, Value)
	panic("")
}

func Expected(s string) {
	Abort(fmt.Sprintf("%s expected", s))
}

func Undefined(s string) {
	Abort(fmt.Sprintf("Undefined identifier: %s", s))
}

func Duplicate(s string) {
	Abort(fmt.Sprintf("Duplicate identifier: %s", s))
}

func Read() byte {
	bytes := make([]byte, 1)
	count, _ := data.Read(bytes)
	if count == 0 {
		return 0
	}
	return bytes[0]
}

func GetChar() {
	Prev = Look
	Look = Read()
	/* Debug: show basic lines in comments
	fmt.Fprintf(Line, "%c", Look)
	if Look == '\n' {
		fmt.Printf("; %s", Line)
		Line = bytes.NewBufferString("")
	}
	*/
}

func GetSymbol(s string) Symbol {
	var sym Symbol
	sym.n = s
	return sym
}

func UpCase (c byte) byte {
	if c >= 97 && c <= 122 {
		c = c - 32
	}
	return c
}

func MatchString(s string) {
	if Value != s {
		Expected(fmt.Sprintf("'%s'", s))
	}
	Next()
}

func Lookup(s string) int {
	i := 0
	for i = len(Keywords) - 1; i >= 0; i-- {
		if strings.ToUpper(s) == Keywords[i] {
			break
		}
	}
	return i + 1
}

func Locate(n Symbol) int {
	index := -1
	for i := len(Symbols) - 1; i >= 0; i-- {
		if strings.ToUpper(n.n) == strings.ToUpper(Symbols[i].n) {
			index = i
			break
		}
	}
	return index
}

func CheckIdent() {
	if Token != 'x' {
		Expected("Identifier")
	}
}

func InTable(n string) bool {
	symbol := Locate(GetSymbol(n))
	return symbol >= 0
}

func CheckTable(n string) {
	if !InTable(n) {
		Undefined(n)
	}
}

func CheckDup(n string) {
	if InTable(n) {
		Duplicate(n)
	}
}

func AddVar(n string, t rune) {
	CheckDup(n)
	if StackDepth == cap(Symbols) {
		Abort("Symbol Table Full")
	}
	Symbols[StackDepth].n = n
	Symbols[StackDepth].t = t
	Symbols[StackDepth].l = -StackDepth
	StackDepth++
}

func IsAlpha(c byte) bool {
	return UpCase(c) >= 65 && UpCase(c) <= 90
}

func IsDigit(c byte) bool {
	return c >= 48 && c <= 57
}

func IsAlNum(c byte) bool {
	return IsAlpha(c) || IsDigit(c)
}

func IsOp(c byte) bool {
	return IsAddOp(c) || IsMulOp(c) || c == '<' || c == '>' || c == ':' || c == '='
}

func IsAddOp(c byte) bool {
	return c == '+' || c == '-'
}

func IsMulOp(c byte) bool {
	return c == '*' || c == '/' || c == '%'
}

func IsWhite(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func IsBool(c byte) bool {
	return c == 'T' || c == 'F'
}

func IsBoolOr(c byte) bool {
	return c == '|' || c == '~'
}

func IsBoolRel(c byte) bool {
	return c == '=' || c == '<' || c == '>'
}

func SkipWhite() {
	for IsWhite(Look) {
		GetChar()
	}
}

func SkipComma() {
	SkipWhite()
	if Look == ',' {
		GetChar()
		SkipWhite()
	}
}

func GetName() {
	SkipWhite()
	Value = ""
	if !IsAlpha(Look) {
		Expected("Name")
	}

	token := bytes.NewBufferString("")
	for  {
		fmt.Fprintf(token, "%c", UpCase(Look))
		GetChar()
		if !IsAlNum(Look) {
			break
		}
	}

	if (Look == ':') {
		PostLabel(string(token.Bytes()))
		GetChar()
		Next()
	} else {
		Token = 'x'
		Value = string(token.Bytes())
	}
}

func GetNum() {
	SkipWhite()
	Value = ""
	value := bytes.NewBufferString("")
	if !IsDigit(Look) {
		Expected("Integer")
	}
	for {
		fmt.Fprintf(value, "%c", Look)
		GetChar()
		if !IsDigit(Look) {
			break
		}
	}
	Token = '#'
	Value = string(value.Bytes())
}

func GetString() {
	SkipWhite()
	Value = ""
	value := bytes.NewBufferString("")
	if Look != '"' {
		Expected("String Constant")
	}
	GetChar()
	for {
		fmt.Fprintf(value, "%c", Look)
		GetChar()
		if Look == '"' {
			break
		}
	}
	GetChar()
	Token = '$'
	Value = string(value.Bytes())
}

func GetOp() {
	SkipWhite()
	Token = Look
	Value = string(Look)
	GetChar()
	if Look != ' ' && !IsAlNum(Look) {
		Value += string(Look)
	}
}

func Op_Add() {
	Next()
	Term()
	PopAdd()
}

func Op_Subtract() {
	Next()
	Term()
	PopSub()
}

func Op_Multiply() {
	Next()
	Factor()
	PopMul()
}

func Op_Divide() {
	Next()
	Factor()
	PopDiv()
}

func Op_Mod() {
	Next()
	Factor()
	PopMod()
}

func Op_Pow() {
	Next()
	Factor()
	PopPow()
}

func Op_ShiftLeft() {
	Next()
	Next()
	Push()
	Factor()
	PopShl()
}

func Op_ShiftRight() {
	Next()
	Next()
	Push()
	Factor()
	PopShr()
}

func Op_Equal() {
	MatchString("==")
	NextExpression()
	SetEqual()
}

func Op_NotEqual() {
	NextExpression()
	SetNotEqual()
}

func Op_GreaterOrEqual() {
	NextExpression()
	SetGreaterOrEqual()
}

func Op_LessOrEqual() {
	NextExpression()
	SetLessOrEqual()
}

func Op_Less() {
	Next()
	switch Token {
	case '=':
		Op_LessOrEqual()
	case '>':
		Op_NotEqual()
	default:
		CompareExpression()
		SetLess()
	}
}

func Op_Greater() {
	Next()
	switch Token {
	case '=':
		Op_GreaterOrEqual()
	default:
		CompareExpression()
		SetGreater()
	}
}

func BoolOr() {
	Next()
	BoolTerm()
	PopOr()
}

func BoolXor() {
	Next()
	BoolTerm()
	PopXor()
}

func Relation() {
	Expression()
	if IsBoolRel(Token) {
		Push()
		switch Token {
		case '=':
			Op_Equal()
		case '<':
			Op_Less()
		case '>':
			Op_Greater()
		}
	}
}

func NotFactor() {
	if Look == '!' {
		Next()
		Relation()
		Not()
	} else {
		Relation()
	}
}

func BoolTerm() {
	NotFactor()
	SkipWhite()
	for Token == '&' {
		Push()
		Next()
		NotFactor()
		PopAnd()
	}
}

func BoolExpression() {
	BoolTerm()
	for IsBoolOr(Look) {
		Push()
		switch Look {
		case '|':
			BoolOr()
		case '~':
			BoolXor()
		}
	}
}

func Factor() {
	if Token == '(' {
		Next()
		Expression()
		Next()
	} else {
		if Value == "KEY" {
			Call("getkey")
		} else if Value == "INPUT" {
			Input()
		} else if Value == "STR" {
			FuncStr()
		} else if Value == "CHR" {
			FuncChr()
		} else if Value == "VAL" {
			FuncVal()
		} else if Value == "PEEK" {
			FuncPeek()
		} else if Value == "RND" {
			Rnd()
		} else if Token == 'x' {
			LoadVar(Value)
		} else if Token == '#' {
			LoadConst(Value)
		} else if Token == '$' {
			LoadConstString(Value)
		} else {
			Expected("Math Factor")
		}
		Next()
	}
	if Token == '^' {
		Push()
		Op_Pow()
	} else if Value == "<<" {
		Op_ShiftLeft()
	} else if Value == ">>" {
		Op_ShiftRight()
	}
}

func TermCheck() {
	for IsMulOp(Token) {
		Push()
		switch Token {
		case '*':
			Op_Multiply()
		case '/':
			Op_Divide()
		case '%':
			Op_Mod()
		}
	}
}

func Term() {
	Factor()
	TermCheck()
}

func Expression() {
	if IsAddOp(Token) {
		Clear()
	} else {
		Term()
	}
	for IsAddOp(Token) {
		Push()
		switch Token {
		case '+':
			Op_Add()
		case '-':
			Op_Subtract()
		}
	}
}

func CompareExpression() {
	Expression()
	PopCompare()
}

func NextExpression() {
	Next()
	CompareExpression()
}

func Next() {
	SkipWhite()
	if IsAlpha(Look) {
		GetName()
	} else if IsDigit(Look) {
		GetNum()
	} else if Look == '"' {
		GetString()
	} else {
		GetOp()
	}
}

func Scan() {
	k := Lookup(Value)
	if k == 0 {
		Token = 'x'
	} else {
		Token = Tokens[k]
	}
}

func Assignment() {
	// CheckTable(Value)
	name := Value
	Next()
	MatchString("=")
	BoolExpression()
	if !InTable(name) {
		Undefined(name)
	}
	Store(name)
}

func NewConst() string {
	label := fmt.Sprintf("c%d", ConstCount)
	ConstCount += 1
	return label
}

func NewLabel() string {
	label := fmt.Sprintf("l%d", LabelCount)
	LabelCount += 1
	return label
}

func PostLabel(l string) {
	EmitLine(fmt.Sprintf(":%s", l))
}

func If() {
	Next()
	BoolExpression()
	l1 := NewLabel()
	l2 := l1
	BranchFalse(l1)
	MatchString("THEN")
	Block()
	if Token == 'l' {
		Next()
		l2 = NewLabel()
		Branch(l2)
		PostLabel(l1)
		Block()
	}
	PostLabel(l2)
	MatchString("END")
	MatchString("IF")
}

func Loop() {
	l1 := NewLabel()
	l2 := NewLabel()
	PostLabel(l1)
	Next()
	if Value == "WHILE" {
		Next()
		BoolExpression()
		BranchFalse(l2)
	}
	Block()
	MatchString("END")
	MatchString("LOOP")
	Branch(l1)
	PostLabel(l2)
}

func Print() {
	newLine := true
	Next()
	BoolExpression()
	Call("print")
	for Token == ';' {
		Next()
		if Value == "CONTINUE" {
			Next()
			newLine = false
			break
		}
		BoolExpression()
		Call("print")
	}
	if newLine {
		Call("printnl")
	}
}

func Rem() {
	Next()
	for Look != '\n' {
		GetChar()
	}
	Next()
}

func Block() {
	Scan()
	for Token != 'e' && Token != 'l' {
		switch Token {
		case 'i':
			If()
		case 'w':
			Loop()
		case 'c':
			Cls()
		case 'o':
			Loc()
		case 'p':
			Print()
		case 'r':
			Rem()
		case 'k':
			Color()
		case 'q':
			Poke()
		case 'u':
			PutChar()
		case 'g':
			Goto()
		case 'n':
			Rnd()

		default:
			Assignment()
		}
		Scan()
	}
}

func Alloc() {
	Next()
	if Token != 'x' {
		Expected("Variable name")
	}
	CheckDup(Value)
	AddVar(Value, 'i')
	Next()
}

func Declarations() {
	Scan()
	currentStack := StackDepth
	for Token == 'd' {
		Alloc()
		for Token == ',' {
			Alloc()
		}
	}

	// Allocate space on the stack for the vars
	i := StackDepth - currentStack
	if i > 0 {
		EmitLine(fmt.Sprintf("SUB SP, %d ; Alloc space on stack", i))
	}
}

func Init() {
	GetChar()
	Next()
}

func Program() {
	Init()
	Prolog()
	Declarations()
	Block()
	MatchString("END")
	Epilog()
}

func parse(file *os.File) {
	data = file
	Program()
}
