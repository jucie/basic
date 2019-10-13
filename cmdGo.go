package main

import (
	"bufio"
	"fmt"
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

func (c cmdGo) cmdName() string {
	if c.sub {
		return "GOSUB"
	}
	return "GOTO"
}

func (p *parser) parseGo() *cmdGo {
	l := p.lex.peek()
	result := &cmdGo{}
	switch l.token {
	case tokTo:
		result.sub = false
	case tokSub:
		result.sub = true
	default:
		return nil
	}
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

func (c cmdGo) generateC(wr *bufio.Writer) {
	if c.sub {
		fmt.Fprintf(wr, "\treturn_from_subroutine();\n")
	} else {
		fmt.Fprintf(wr, "\tgoto_line = %d; break;\n", c.dst.nbr)
	}
}
