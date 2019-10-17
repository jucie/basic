package main

import (
	"bufio"
	"fmt"
)

type cmdRem struct {
	text string
}

func (p *parser) parseRem() *cmdRem {
	text := p.lex.peek().s
	if len(text) > 0 {
		if text[0] != ' ' {
			text = "REM" + text
		}
	}
	return &cmdRem{text: text}
}

func (c cmdRem) receive(g guest) {
}

func (c cmdRem) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "\t/*%s*/\n", c.text)
}
