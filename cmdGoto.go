package main

import (
	"strconv"
)

type cmdGoto struct {
	line int
}

func (p *parser) parseGoto() *cmdGoto {
	result := &cmdGoto{}
	l := p.lex.peek()

	if l.token != tokNumber {
		return nil
	}
	var err error
	result.line, err = strconv.Atoi(l.s)
	if err != nil {
		return nil
	}
	p.lex.next()

	return result
}

func (c cmdGoto) receive(g guest) {
}
