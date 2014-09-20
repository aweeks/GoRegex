package parser

import (
	"ast"
	"errors"
	"fmt"
	"strings"
)

const debug = false

type Parser struct {
	read *strings.Reader
}

func NewParser(s string) *Parser {
	p := new(Parser)
	p.read = strings.NewReader(s)
	return p
}

func (p *Parser) done() bool {
	return p.read.Len() == 0
}

func (p *Parser) peek() (rune, error) {
	r, _, err := p.read.ReadRune()
	p.read.UnreadRune()
	if debug {
		fmt.Printf("peek %q\n", r)
	}

	return r, err

}

func (p *Parser) eat(expect rune) error {
	if debug {
		fmt.Printf("eat %q\n", expect)
	}

	r, _, err := p.read.ReadRune()
	if err != nil {
		return err
	}

	if r != expect {
		return errors.New(fmt.Sprintf("Expected %q, got %q", expect, r))
	}
	return nil
}

func (p *Parser) next() (rune, error) {
	r, _, err := p.read.ReadRune()
	if debug {
		fmt.Printf("next %q\n", r)
	}
	return r, err
}

func (p *Parser) ParseRegex() (ast.Regex, error) {
	if debug {
		fmt.Println("ParseRegex")
	}

	term, err := p.ParseTerm()
	if err != nil {
		return nil, err
	}
	if !p.done() {
		peek, err := p.peek()
		if err != nil {
			return nil, err
		}

		if peek == '|' {
			err = p.eat('|')
			if err != nil {
				return nil, err
			}

			regex, err := p.ParseRegex()
			if err != nil {
				return nil, err
			}
			return ast.NewRegexOr(term, regex), nil
		}
	}
	return term, nil

}

func (p *Parser) ParseTerm() (ast.Regex, error) {
	if debug {
		fmt.Println("ParseTerm")
	}

	factor := ast.NewRegexNull()

	for {
		if p.done() {
			break
		}

		peek, err := p.peek()
		if err != nil {
			return nil, err
		}
		if peek == '|' || peek == ')' {
			break
		}

		next, err := p.ParseFactor()
		if err != nil {
			return nil, err
		}

		factor = ast.NewRegexConcat(factor, next)
	}
	return factor, nil
}

func (p *Parser) ParseFactor() (ast.Regex, error) {
	if debug {
		fmt.Println("ParseFactor")
	}

	base, err := p.ParseBase()
	if err != nil {
		return nil, err
	}

	if !p.done() {
		peek, err := p.peek()
		if err != nil {
			return nil, err
		}

		if peek == '*' {
			err := p.eat('*')
			if err != nil {
				return nil, err
			}
			return ast.NewRegexStar(base), err
		}
	}

	return base, nil
}

func (p *Parser) ParseBase() (ast.Regex, error) {
	if debug {
		fmt.Println("ParseBase")
	}

	peek, err := p.peek()
	if err != nil {
		return nil, err
	}
	switch peek {
	case '(':
		err = p.eat('(')
		if err != nil {
			return nil, err
		}

		regex, err := p.ParseRegex()

		if err != nil {
			return nil, err
		}

		p.eat(')')
		if err != nil {
			return nil, err
		}

		return regex, err
	case '\\':
		p.eat('\\')
		p.next()
		return ast.NewRegexPrimitive(peek), nil
	default:
		p.next()
		return ast.NewRegexPrimitive(peek), nil
	}
	return nil, nil
}
