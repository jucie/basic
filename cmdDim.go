package main

import (
	"bufio"
	"fmt"
)

type cmdDim struct {
	vars []*astVarRef
}

func (p *parser) parseDim() *cmdDim {
	result := &cmdDim{}
	l := p.lex.peek()

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

func (c cmdDim) receive(g guest) {
	for _, v := range c.vars {
		g.visit(v)
	}
}

func (c cmdDim) generateC(wr *bufio.Writer) {
	for _, v := range c.vars {
		fmt.Fprintf(wr, "\tdim_%s(&%s_var,%d,", v.type_, v.nameForC(), len(v.index))
		for i, index := range v.index {
			if i != 0 {
				wr.WriteRune(',')
			}
			index.generateC(wr)
		}
		fmt.Fprintf(wr, ");\n")
	}
}
