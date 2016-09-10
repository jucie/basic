package main

import (
	"io"
)

type parser struct {
	rd io.Reader
}

func newParser(rd io.Reader) *parser {
	return &parser{rd: rd}
}

func (p *parser) parseProgram(prog *program) {

}
