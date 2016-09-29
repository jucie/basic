package main

import (
	"strconv"
)

type cmdGoto struct {
	dst targetLine
}

func (p *parser) parseGoto() *cmdGoto {
	result := &cmdGoto{}
	l := p.lex.peek()

	if l.token != tokNumber {
		return nil
	}
	var err error
	result.dst.nbr, err = strconv.Atoi(l.s)
	if err != nil {
		return nil
	}
	p.lex.next()

	return result
}

func (c cmdGoto) receive(g guest) {
}
