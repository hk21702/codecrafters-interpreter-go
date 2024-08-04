package parser

import (
	"fmt"
	"strconv"

	"github.com/codecrafters-io/interpreter-starter-go/internal/token"
)

type parser struct {
	tokens  []token.Token
	current int
}

func New(tokens []token.Token) *parser {
	return &parser{tokens: tokens}
}

func (p *parser) Parse() (float64, error) {
	return p.expression()
}

func (p *parser) expression() (float64, error) {
	result, err := p.term()
	if err != nil {
		return 0, err
	}

	for p.match(token.Plus, token.Minus) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return 0, err
		}

		switch operator.Type {
		case token.Plus:
			{
				result += right
			}
		case token.Minus:
			{
				result -= right
			}
		}
	}

	return result, nil
}

func (p *parser) term() (float64, error) {
	result, err := p.factor()
	if err != nil {
		return 0, err
	}

	for p.match(token.Star, token.Slash) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return 0, err
		}

		if operator.Type == token.Star {
			result *= right
		} else if operator.Type == token.Slash {
			result /= right
		}
	}

	return result, nil
}

func (p *parser) factor() (float64, error) {
	if p.match(token.Number) {
		value, err := strconv.ParseFloat(p.previous().Lexeme, 64)
		if err != nil {
			return 0, err
		}
		return value, nil
	}

	if p.match(token.LParen) {
		value, err := p.expression()
		if err != nil {
			return 0, err
		}
		if !p.match(token.RParen) {
			return 0, fmt.Errorf("expected ')' after expression")
		}
		return value, nil
	}

	return 0, fmt.Errorf("expected expression")
}

func (p *parser) match(types ...token.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *parser) check(t token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *parser) isAtEnd() bool {
	return p.peek().Type == token.Eof
}

func (p *parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *parser) previous() token.Token {
	return p.tokens[p.current-1]
}
