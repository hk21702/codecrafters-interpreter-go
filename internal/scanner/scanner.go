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
			if errors.As(err, &lexer.UnexpectedChar{}) {
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
	if tk.Type != token.Eof {
		fmt.Printf("%s %s null\n", tk.Type, tk.Literal)
	} else {
		fmt.Printf("%s  null\n", tk.Type)
	}
}
