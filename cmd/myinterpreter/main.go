package main

import (
	"fmt"
	"os"
)

var runeNames map[rune]string

func init() {
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
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

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
		parseToken(current)
	}

	fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner

}

func parseToken(ru rune) {
	fmt.Printf("%s %c null\n", runeNames[ru], ru)
}
