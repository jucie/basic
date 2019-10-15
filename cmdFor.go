package main

import (
	"bufio"
	"fmt"
)

type cmdFor struct {
	index *astVarRef
	begin *astExpr
	end   *astExpr
	step  *astExpr
}

func (p *parser) parseFor() *cmdFor {
	result := &cmdFor{}
	l := p.lex.peek()

	if l.token != tokID {
		return nil
	}
	result.index = p.parseVarRef()
	if result.index == nil {
		return nil
	}

	if l.token != '=' {
		return nil
	}
	p.lex.next()

	result.begin = p.parseExpr(false)

	if l.token != tokTo {
		return nil
	}
	p.lex.next()

	result.end = p.parseExpr(false)

	if l.token == tokStep {
		p.lex.next()
		result.step = p.parseExpr(false)
	}

	return result
}

func (c cmdFor) receive(g guest) {
	g.visit(c.index)
	g.visit(c.begin)
	g.visit(c.end)
	if c.step != nil {
		g.visit(c.step)
	}
}

func (c cmdFor) generateC(wr *bufio.Writer) {
	label := createLabel()
	wr.WriteRune('\t')
	c.index.generateCLValue(wr, "for")
	wr.WriteRune(',')
	c.begin.generateC(wr)
	fmt.Fprintf(wr, ",")
	c.end.generateC(wr)
	fmt.Fprintf(wr, ",")
	if c.step == nil {
		fmt.Fprintf(wr, "1.0f")
	} else {
		c.step.generateC(wr)
	}
	fmt.Fprintf(wr, ",%d);\n", label)
	fmt.Fprintf(wr, "case %d:\n", label)
}
