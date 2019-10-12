package main

import (
	"bufio"
	"fmt"
)

type cmdRem struct {
	text string
}

func (p *parser) parseRem() *cmdRem {
	return &cmdRem{text: p.lex.peek().s}
}

func (c cmdRem) receive(g guest) {
}

func (c cmdRem) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "\t/*%s*/\n", c.text)
}
