package main

import (
	"fmt"
	"os"
)

var runeNames map[rune]string
var exitCode int

func init() {
	exitCode = 0

	runeNames = make(map[rune]string)
	runeNames['('] = "LEFT_PAREN"
	runeNames[')'] = "RIGHT_PAREN"
	runeNames['{'] = "LEFT_BRACE"
	runeNames['}'] = "RIGHT_BRACE"
	runeNames['*'] = "STAR"
	runeNames['.'] = "DOT"
	runeNames[','] = "COMMA"
	runeNames['+'] = "PLUS"
	runeNames['-'] = "MINUS"
	runeNames[';'] = "SEMICOLON"
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	strFileContents := string(fileContents)
	for _, current := range strFileContents {
		parseToken(current, 1)
	}

	fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner

	os.Exit(exitCode)
}

func parseToken(ru rune, lineNum int) {
	runeName, exists := runeNames[ru]

	if exists {
		fmt.Printf("%s %c null\n", runeName, ru)
	} else {
		// Unrecognized rune
		fmt.Printf("[line %d] Error: Unexpected character: %c\n", lineNum, ru)
		exitCode = 65
	}
}
