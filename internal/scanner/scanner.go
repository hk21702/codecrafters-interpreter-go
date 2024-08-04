package scanner

import (
	"errors"
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/internal/lexer"
	"github.com/codecrafters-io/interpreter-starter-go/internal/parser"
	"github.com/codecrafters-io/interpreter-starter-go/internal/token"
)

// Scan the inputs
func Scan(fileContents []byte, print bool) (exitCode int) {
	lex := lexer.New(string(fileContents))
	var tokens []token.Token

	tk, err := lex.ReadToken()
	for tk.Type != token.Eof {
		if err != nil {
			if errors.As(err, &lexer.UnexpectedChar{}) || errors.As(err, &lexer.UnterminatedStr{}) {
				fmt.Fprint(os.Stderr, err.Error())
				exitCode = 65
			}
		} else {
			if print {
				printToken(tk)
			}
			tokens = append(tokens, tk)
		}

		tk, err = lex.ReadToken()
	}
	if print {
		// parse EOF
		printToken(tk)
	}
	tokens = append(tokens, tk)

	par := parser.New(tokens)
	output, _ := par.Parse()

	fmt.Fprint(os.Stdout, output)

	return exitCode
}

// Parse the input
func printToken(tk token.Token) {
	var literalStr string
	if tk.Literal == nil {
		literalStr = "null"
	} else {
		switch v := tk.Literal.(type) {
		case float64:
			if v == float64(int64(v)) {
				literalStr = fmt.Sprintf("%.1f", v)
			} else {
				literalStr = fmt.Sprintf("%g", v)
			}
		case float32:
			if v == float32(int32(v)) {
				literalStr = fmt.Sprintf("%.1f", v)
			} else {
				literalStr = fmt.Sprintf("%g", v)
			}
		default:
			literalStr = fmt.Sprintf("%v", tk.Literal)
		}
	}

	if tk.Type != token.Eof {
		fmt.Printf("%s %s %v\n", tk.Type, tk.Lexeme, literalStr)
	} else {
		fmt.Printf("%s  %v\n", tk.Type, literalStr)
	}
}
