package main

import (
	"bufio"
	"fmt"
)

type cmdNext struct {
	vars []*astVarRef
}

func (p *parser) parseNext() *cmdNext {
	result := &cmdNext{}
	if p.isEndOfCommand() {
		return result
	}

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
	return result
}

func (c cmdNext) receive(g guest) {
	for _, v := range c.vars {
		g.visit(v)
	}
}

func (c cmdNext) generateC(wr *bufio.Writer) {
	for _, v := range c.vars {
		fmt.Fprintf(wr, "\tif (!end_loop(&target)) break; /* NEXT %s */\n", v.id)
	}
}
