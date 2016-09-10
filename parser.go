package main

import (
	"io"
)

type parser struct {
	lex *lexer
}

func newParser(rd io.Reader) *parser {
	lex := newLexer(rd)
	return &parser{lex: lex}
}

func (p *parser) parseProgram(prog *program) {

}
