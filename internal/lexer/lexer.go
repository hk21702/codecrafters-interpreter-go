package lexer

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/internal/token"
)

type lexer struct {
	input       string
	position    int
	nxtPosition int
	char        byte
	line        int
}

func New(input string) *lexer {
	lex := lexer{input: input, line: 1}
	return &lex
}

// Read and get the current token and prep the lexer for the next
//
//	Returns UnexpectedChar if an unexpected char is encountered
//
// Parameters:
//
//	lex - The lexer to read from
//
// Returns:
//
//	token - The read token
//	err - Error if there is one, nil otherwise.
func (lex *lexer) ReadToken() (tk token.Token, err error) {
	lex.nxtChar()
	lex.skipWhiteSpace()
	tk.Literal = string(lex.char)

	// Check potential multi char tokens
	switch lex.char {
	case '/':
		{
			if lex.peekChar() == '/' {
				lex.nxtLine()
				lex.ReadToken()
			}
		}
	case '=':
		{
			if lex.peekChar() == '=' {
				// Double equals
				return lex.doubleToken(tk, token.EqualEqual)
			}
		}
	case '!':
		{
			if lex.peekChar() == '=' {
				// !=
				return lex.doubleToken(tk, token.BangEqual)
			}
		}
	case '<':
		{
			if lex.peekChar() == '=' {
				return lex.doubleToken(tk, token.LessEqual)
			}
		}
	case '>':
		{
			if lex.peekChar() == '=' {
				return lex.doubleToken(tk, token.GreaterEqual)
			}
		}
	}

	// Check reserved single char tokens
	tokenT, found := token.RuneMap[lex.char]
	if !found {
		tk.Type = token.UnexpectedChar

		return tk, UnexpectedChar{Char: lex.char, Line: lex.line}
	}

	tk.Type = tokenT
	return tk, nil
}

// Helper method to handle when the next token is a double rune/char token.
func (lex *lexer) doubleToken(tk token.Token, tType token.TokenType) (token.Token, error) {
	lex.nxtChar()
	tk.Type = tType
	tk.Literal += string(lex.char)
	return tk, nil
}

// Skip to the end of the line
func (lex *lexer) nxtLine() {
	for lex.char != '\n' && lex.char != 0 {
		lex.nxtChar()
	}
}

// Read the next char/rune in the lexer
// If the previous char was already the last char, sets the char to 0.
// If current position is already EOF, does nothing.
// If the previous position was a new line control char, increments line.
func (lex *lexer) nxtChar() {
	if lex.nxtPosition >= len(lex.input) {
		lex.char = 0
		return
	}
	if lex.char == '\n' {
		lex.line++
	}

	lex.char = lex.input[lex.nxtPosition]

	lex.position = lex.nxtPosition
	lex.nxtPosition++
}

// Peek at the next char/rune in the lexer
// Does not move the position
func (lex *lexer) peekChar() byte {
	if lex.nxtPosition >= len(lex.input) {
		return 0
	}

	return lex.input[lex.nxtPosition]
}

// Moves position until char is no longer whitespace
// Spaces, tabs, new lines
func (lex *lexer) skipWhiteSpace() {
	for lex.char == ' ' || lex.char == '\t' || lex.char == '\n' {
		lex.nxtChar()
	}
}

// Represents an error for when the interpreter encounters an unexpected character in the input.
type UnexpectedChar struct {
	Char byte
	Line int
}

// Implements the Error interface for UnexpectedChar
func (err UnexpectedChar) Error() string {
	return fmt.Sprintf("[line %d] Error: Unexpected character: %c\n", err.Line, err.Char)
}
