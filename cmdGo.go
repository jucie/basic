package main

import (
	"strconv"
)

type targetLine struct {
	nbr int
	adr *progLine
}
type cmdGo struct {
	dst targetLine
	sub bool
}

func (p *parser) parseGo() *cmdGo {
	l := p.lex.peek()
	result := &cmdGo{sub: (l.token == tokGosub)}
	p.lex.next()

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

func (c cmdGo) receive(g guest) {
}
