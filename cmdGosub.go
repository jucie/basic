package main

import (
	"strconv"
)

type targetLine struct {
	nbr int
	adr *progLine
}
type cmdGosub struct {
	dst targetLine
}

func (p *parser) parseGosub() *cmdGosub {
	result := &cmdGosub{}
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

func (c cmdGosub) receive(g guest) {
}
