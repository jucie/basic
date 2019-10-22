package main

import (
	"bufio"
	"fmt"
)

type cmdStop struct {
}

func (p *parser) parseStop() *cmdStop {
	return &cmdStop{}
}

func (c cmdStop) receive(g guest) {
}

func (c cmdStop) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "\texit(0);\n")
}
