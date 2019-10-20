package main

import (
	"bufio"
	"fmt"
)

type cmdInput struct {
	label string // optional
	vars  []*astVarRef
}

func (p *parser) parseInput() *cmdInput {
	result := &cmdInput{}
	l := p.lex.peek()

	if l.token == tokString {
		result.label = l.s
		p.lex.next() // separador
		p.lex.next() // pula o separador
	}
	for {
		v := p.parseVarRef()
		if v == nil {
			break
		}
		result.vars = append(result.vars, v)
		if l.token != ',' {
			break
		}
		p.lex.next()
	}
	if len(result.vars) == 0 {
		return nil
	}
	return result
}

func (c cmdInput) generateC(wr *bufio.Writer) {
	if c.label != "" {
		fmt.Fprintf(wr, "\tprint_str(\"%s\");\n", c.label)
	}
	fmt.Fprintf(wr, "\tinput_to_buffer();\n")

	for _, v := range c.vars {
		fmt.Fprintf(wr, "\t")
		v.generateCLValue(wr, "input")
		fmt.Fprintf(wr, ");\n")
	}
}

func (c cmdInput) receive(g guest) {
	for _, v := range c.vars {
		g.visit(v)
	}
}
