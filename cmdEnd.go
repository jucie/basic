package main

import (
	"bufio"
	"fmt"
)

type cmdEnd struct {
}

func (p *parser) parseEnd() *cmdEnd {
	return &cmdEnd{}
}

func (c cmdEnd) receive(g guest) {
}

func (c cmdEnd) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "\texit(0);\n")
}
