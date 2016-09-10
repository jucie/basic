package main

import (
	"bufio"
	"unicode"
)

type coord struct {
	row int
	col int
}

type lexer struct {
	rd          *bufio.Reader
	pos         coord
	token       token
	intValue    int
	floatValue  float32
	stringValue string
}

func newLexer(rd *bufio.Reader) *lexer {
	return &lexer{rd: rd}
}

func (lex *lexer) peek() token {
	return lex.token
}

func (lex *lexer) next() {
	for {
		lex.token = lex.walk()
		if lex.token != tokSpace && lex.token != tokComment {
			break
		}
	}
}

func (lex *lexer) walk() token {
	r, _, err := lex.rd.ReadRune()
	if err != nil {
		return tokEof
	}
	if unicode.IsDigit(r) {
		for {
			r, _, err = lex.rd.ReadRune()
			if err != nil {
				break
			}

		}

	}
	return tokEof
}
