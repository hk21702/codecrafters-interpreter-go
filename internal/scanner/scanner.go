package scanner

import (
	"errors"
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/internal/lexer"
	"github.com/codecrafters-io/interpreter-starter-go/internal/token"
)

// Scan the inputs
func Scan(fileContents []byte) (exitCode int) {
	lex := lexer.New(string(fileContents))

	tk, err := lex.ReadToken()
	for tk.Type != token.Eof {
		if err != nil {
			if errors.As(err, &lexer.UnexpectedChar{}) || errors.As(err, &lexer.UnterminatedStr{}) {
				fmt.Fprint(os.Stderr, err.Error())
				exitCode = 65
			}
		} else {
			parseToken(tk)
		}

		tk, err = lex.ReadToken()
	}
	// parse EOF
	parseToken(tk)

	return exitCode
}

// Parse the input
func parseToken(tk token.Token) {
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
