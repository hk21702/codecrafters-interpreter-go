// Module storing keywork token types
package token

type TokenType string

const (
	UnexpectedChar TokenType = "UNEXPECTED_CHAR"
	LParen         TokenType = "LEFT_PAREN"
	RParen         TokenType = "RIGHT_PAREN"
	LBrace         TokenType = "LEFT_BRACE"
	RBrace         TokenType = "RIGHT_BRACE"
	Star           TokenType = "STAR"
	Dot            TokenType = "DOT"
	Comma          TokenType = "COMMA"
	Plus           TokenType = "PLUS"
	Minus          TokenType = "MINUS"
	Semicolon      TokenType = "SEMICOLON"
	Equal          TokenType = "EQUAL"
	Eof            TokenType = "EOF"
	EqualEqual     TokenType = "EQUAL_EQUAL"
	Bang           TokenType = "BANG"
	BangEqual      TokenType = "BANG_EQUAL"
	Less           TokenType = "LESS"
	LessEqual      TokenType = "LESS_EQUAL"
	Greater        TokenType = "GREATER"
	GreaterEqual   TokenType = "GREATER_EQUAL"
	String         TokenType = "STRING"
	Slash          TokenType = "SLASH"
	Number         TokenType = "NUMBER"
	Identifier     TokenType = "IDENTIFIER"
	And            TokenType = "AND"
	Class          TokenType = "CLASS"
	Else           TokenType = "ELSE"
	False          TokenType = "FALSE"
	For            TokenType = "FOR"
	Fun            TokenType = "FUN"
	If             TokenType = "IF"
	Nil            TokenType = "NIL"
	Or             TokenType = "OR"
	Print          TokenType = "PRINT"
	Return         TokenType = "RETURN"
	Super          TokenType = "SUPER"
	This           TokenType = "THIS"
	True           TokenType = "TRUE"
	Var            TokenType = "VAR"
	While          TokenType = "WHILE"
)

var RuneMap = map[byte]TokenType{
	'(': LParen,
	')': RParen,
	'{': LBrace,
	'}': RBrace,
	'*': Star,
	'.': Dot,
	',': Comma,
	'+': Plus,
	'-': Minus,
	';': Semicolon,
	'=': Equal,
	'!': Bang,
	'<': Less,
	'>': Greater,
	'/': Slash,
	0:   Eof,
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
}

var TokenMap = map[string]TokenType{
	"and":    And,
	"class":  Class,
	"else":   Else,
	"false":  False,
	"for":    For,
	"fun":    Fun,
	"if":     If,
	"nil":    Nil,
	"or":     Or,
	"print":  Print,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Var,
	"while":  While,
}
