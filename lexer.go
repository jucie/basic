package main

import (
	"io"
)

type coord struct {
	row int
	col int
}

type lexer struct {
	rd    io.Reader
	pos   coord
	token token
}

func newLexer(rd io.Reader) *lexer {
	return &lexer{rd: rd}
}

func (lex *lexer) peek() token {
	return lex.token
}

func (lex *lexer) next() {
	for {
		lex.walk()
		if lex.token != tokSpace && lex.token != tokComment {
			break
		}
	}
}

func (lex *lexer) walk() {

}
