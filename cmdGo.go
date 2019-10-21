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

var nextLabel = 0

func createLabel() int {
	nextLabel--
	return nextLabel
}

func (c cmdGo) generateC(wr *bufio.Writer) {
	if c.sub {
		returnAddress := createLabel()
		fmt.Fprintf(wr, "\tpush_address(%d);\n", returnAddress)
		fmt.Fprintf(wr, "\ttarget = %d; break;\n", c.dst.adr.switchID)
		fmt.Fprintf(wr, "case %d:\n", returnAddress)
	} else {
		fmt.Fprintf(wr, "\ttarget = %d; break;\n", c.dst.adr.switchID)
	}
}
