package main

import (
	"bufio"
	"fmt"
)

type cmdRestore struct {
}

func (p *parser) parseRestore() *cmdRestore {
	return &cmdRestore{}
}

func (c cmdRestore) receive(g guest) {
}

func (c cmdRestore) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "\treset_data_ptr();\n")
}
