package scanner

import (
	"fmt"
	"errors"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/internal/lexer"
	"github.com/codecrafters-io/interpreter-starter-go/internal/token"
)

// Scan the inputs
func Scan(fileContents []byte) (exitCode int){
	lex := lexer.New(string(fileContents))

	for tk, err := lex.ReadToken(); tk.Type != token.Eof; {
		if err != nil {
			if errors.As(err, &lexer.UnexpectedChar{}) {
				fmt.Fprintf(os.Stderr, err.Error())
				exitCode = 65
			}
		}
	}

	return 0
}

// Parse the input
func parseToken(tk token.Token) {
	if tk.Type != token.Eof {
		fmt.Printf("%s %s null\n", tk.Type, tk.Literal)
	} else {
		fmt.Printf("%s null\n", tk.Type)
	}
}
