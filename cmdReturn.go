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
	fmt.Fprintf(wr, "\tpop_address(&target); break;\n")
}
