package main

import (
	"bufio"
	"fmt"
)

type cmdLet struct {
	dst *astVarRef
	src *astExpr
}

func (p *parser) parseLet() *cmdLet {
	dst := p.parseVarRef()

	if p.lex.peek().token != '=' {
		return nil
	}
	p.lex.next()

	src := p.parseExpr(false)
	if src == nil {
		return nil
	}

	return &cmdLet{dst: dst, src: src}
}

func (c cmdLet) receive(g guest) {
	g.visit(c.dst)
	g.visit(c.src)
}

func (c cmdLet) generateC(wr *bufio.Writer) {
	wr.WriteRune('\t')
	parenPending := c.dst.generateCLValue(wr)
	c.src.generateC(wr)
	if parenPending {
		wr.WriteRune(')')
	}
	fmt.Fprintf(wr, ";\n")
}
