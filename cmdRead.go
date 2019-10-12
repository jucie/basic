package main

import (
	"bufio"
	"fmt"
)

type cmdRead struct {
	vars []*astVarRef
}

func (p *parser) parseRead() *cmdRead {
	result := &cmdRead{}
	l := p.lex.peek()

	for {
		v := p.parseVarRef()
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

func (c cmdRead) receive(g guest) {
	for _, v := range c.vars {
		g.visit(v)
	}
}

func (c cmdRead) generateC(wr *bufio.Writer) {
	for _, v := range c.vars {
		fmt.Fprintf(wr, "\tread_data_%s(&", v.finalType())
		v.generateC(wr)
		fmt.Fprintf(wr, ");\n")
	}
}
