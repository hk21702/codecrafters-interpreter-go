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
	Slash TokenType = "SLASH"
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
	Literal string
}
