package main

import (
	"fmt"
	"os"
	"bytes"
	"strings"
	"strconv"
)

var data *os.File
var Look byte
var Keywords = []string { "IF", "ELSE",  "WHILE", "END", "VAR", "CLS", "PRINT", "LOCATE" }
var Tokens = []byte { 'x', 'i', 'l', 'w', 'e', 'v', 'c', 'p', 'o' }
var Token byte
var Value string
var LabelCount = 0
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
	fmt.Printf("\nCurrent token type: %c, value: '%s'\n\n", Token, Value)
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
	Push()
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
	return c == '*' || c == '/'
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
	Token = 'x'
	Value = string(token.Bytes())
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

func GetOp() {
	SkipWhite()
	Token = Look
	Value = string(Look)
	GetChar()
}

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
	EmitLine("SET B, 0")
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
}

func SetEqual() {
	EmitLine("IFE A, B")
	EmitLine("SET A, 0")
}

func SetNotEqual() {
	EmitLine("IFE A, B")
	EmitLine("SET A, 1")
}

func SetGreater() {
	EmitLine("IFG B, A")
	EmitLine("SET A, 0")
}

func SetLess() {
	EmitLine("IFG A, B")
	EmitLine("SET A, 0")
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
	if !InTable(s) {
		Undefined(s)
	}
	symbol := Symbols[Locate(GetSymbol(s))]
	EmitLine(fmt.Sprintf("SET [%#x], A", (0xffff + symbol.l)))
}

func Branch(s string) {
	EmitLine(fmt.Sprintf("SET PC, %s", s))
}

func BranchFalse(s string) {
	EmitLine("IFN A, 0")
	Branch(s)
}

func Prolog() {
	EmitLine("SET PC, begin")
	FuncPrint()
	PostLabel("begin")
}

func Epilog() {
	PostLabel("crash")
	EmitLine("SET PC, crash")
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

func Op_Equal() {
	MatchString("=")
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
		MatchString(")")
	} else {
		if Token == 'x' {
			LoadVar(Value)
		} else if Token == '#' {
			LoadConst(Value)
		} else {
			Expected("Math Factor")
		}
		Next()
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
	Store(name)
}

func NewLabel() string {
	label := fmt.Sprintf("l%c", LabelCount + 97)
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

func While() {
	Next()
	l1 := NewLabel()
	l2 := NewLabel()
	PostLabel(l1)
	BoolExpression()
	BranchFalse(l2)
	Block()
	MatchString("END")
	MatchString("WHILE")
	Branch(l1)
	PostLabel(l2)
}

func Ret() {
	EmitLine("SET PC, POP")
}

func Cls() {
	EmitLine("SET A, 0x200")
	l := NewLabel()
	PostLabel(l)
	EmitLine("SET B, 0x8000")
	EmitLine("ADD B, A")
	EmitLine("SET [B], 0")
	EmitLine("SUB A, 1")
	BranchFalse(l)
	EmitLine("SET [0x8000], 0")
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

func FuncPrint() {
	PostLabel("printnum")
	Push()
	EmitLine("SET I, 0") // Loop counter
	PostLabel("pndiv") // Loop: divide A by 10 until 0 is left
	EmitLine("SET B, A") // Store A (number) for later
	EmitLine("MOD A, 0xa") // Get remainder from division by 10
	EmitLine("ADD A, 0x30") // Add 0x30 to the remainder to get ASCII code
	EmitLine("SET PUSH, A") // Store the remainder (digit) on the stack
	EmitLine("SET A, B") // Get A (number) back
	EmitLine("DIV A, 0xa") // Divide the number by 10
	EmitLine("ADD I, 1") // Increment loop counter
	EmitLine("IFN A, 0") // A > 10: jump to :pnloop1
	EmitLine("SET PC, pndiv")

	PostLabel("pnprint") // Loop: print character by character
	EmitLine("SET A, POP") // Get digit from stack
	EmitLine("SET B, X") // Get current cursor position
	EmitLine("ADD B, 0x8000") // Add video mem address
	EmitLine("SET [B], A") // Set video memory byte to show char
	EmitLine("ADD X, 1") // Increment cursor position
	EmitLine("IFN X, 0x160") // Check if we should do next line (X > 32)
	EmitLine("SET PC, pnline")
	EmitLine("SET X, 0") // First row, first column
	PostLabel("pnline")
	EmitLine("SUB I, 1") // Decrement loop counter
	EmitLine("IFN I, 0")
	EmitLine("SET PC, pnprint") // Jump back to :pnloop2 if there are more chars
	EmitLine("SET A, POP")
	Ret()
}

func Print() {
	Next()
	BoolExpression()
	EmitLine("JSR printnum")
}

func Block() {
	Scan()
	for Token != 'e' && Token != 'l' {
		switch Token {
		case 'i':
			If()
		case 'w':
			While()
		case 'c':
			Cls()
		case 'o':
			Loc()
		case 'p':
			Print()
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
	for Token == 'v' {
		Alloc()
		for Token == ',' {
			Alloc()
		}
	}
}

func Init() {
	GetChar()
	Next()
}

func Program() {
	Init()
	Declarations()
	Prolog()
	Block()
	MatchString("END")
	Epilog()
}

func parse(file *os.File) {
	data = file
	Program()
}
