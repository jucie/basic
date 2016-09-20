package main

import (
	"strconv"
)

type cmdGosub struct {
	line int
}

func (p *parser) parseGosub() *cmdGosub {
	result := &cmdGosub{}
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
