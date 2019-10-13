package main

import (
	"bufio"
	"fmt"
)

type cmdReturn struct {
}

func (p *parser) parseReturn() *cmdReturn {
	return &cmdReturn{}
}

func (c cmdReturn) receive(g guest) {
}

func (c cmdReturn) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "\tgoto_line = stack[sp--]; break;\n")
}
