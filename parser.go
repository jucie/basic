package main

import (
	"bufio"
)

type parser struct {
	lex *lexer
}

func newParser(rd *bufio.Reader) *parser {
	lex := newLexer(rd)
	return &parser{lex: lex}
}

func (p *parser) parseProgram(prog *program) {

}
