package main

import "os"
import "fmt"

// Program header
func header() {
	fmt.Printf("*****************************\n")
	fmt.Printf("* 0x10c QuickBasic compiler *\n")
	fmt.Printf("*         by M4v3R          *\n")
	fmt.Printf("*****************************\n")
	fmt.Printf("\n")
}

// Usage help
func usage() {
	fmt.Printf("Usage: %s <filename.bas>\n\n", os.Args[0])
}

// Compiler errors
func Error(errorString string) {
	fmt.Printf("\nError: %s\n", errorString)
}

// Emit assembly
func Emit(asm string) {
	fmt.Printf("\t%s", asm)
}

func EmitLine(asm string) {
	Emit(asm)
	fmt.Printf("\n")
}

// Compiler main entry point
func main() {
	if len(os.Args) != 2 {
		usage()
		return
	}

	var error error
	var file *os.File

	file, error = os.Open(os.Args[1])

	if error != nil {
		Error(fmt.Sprintf("Couldn't open file: %s", os.Args[1]))
		return
	}

	// Scan and parse source file
	parse(file)
}
