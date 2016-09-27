package main

import (
	"strconv"
)

type cmdGosub struct {
	dst int
}

func (p *parser) parseGosub() *cmdGosub {
	result := &cmdGosub{}
	l := p.lex.peek()

	if l.token != tokNumber {
		return nil
	}
	var err error
	result.dst, err = strconv.Atoi(l.s)
	if err != nil {
		return nil
	}
	p.lex.next()

	return result
}

func (c cmdGosub) receive(g guest) {
}
