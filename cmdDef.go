package main

import (
	"bufio"
	"fmt"
)

type cmdFnDef struct {
	id   string
	arg  string
	expr *astExpr
}

func (p *parser) parseDef() *cmdFnDef {
	l := p.lex.peek()
	result := &cmdFnDef{}

	if l.token != tokFn {
		return nil
	}
	p.lex.next()

	if l.token != tokID {
		return nil
	}
	result.id = l.s
	p.lex.next()

	if l.token != '(' {
		return nil
	}
	p.lex.next()

	if l.token != tokID {
		return nil
	}
	result.arg = l.s
	p.lex.next()

	if l.token != ')' {
		return nil
	}
	p.lex.next()

	if l.token != '=' {
		return nil
	}
	p.lex.next()

	result.expr = p.parseExpr(false)

	return result
}

func (c cmdFnDef) receive(g guest) {
	g.visit(c.expr)
}

func (c cmdFnDef) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "float fn_%s(float %s) {\n", c.id, c.arg)
	fmt.Fprintf(wr, "\treturn ")
	c.expr.generateC(wr)
	fmt.Fprintf(wr, ";\n}\n\n")
}
