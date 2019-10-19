package main

import (
	"bufio"
	"fmt"
)

type cmdFnDef struct {
	id   string
	args []*astVarRef
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
	for {
		v := p.parseVarRef()
		if v == nil {
			break
		}
		result.args = append(result.args, v)
		if l.token != ',' {
			break
		}
		p.lex.next()
	}

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
}

func (c cmdFnDef) generateCFunctionHeader(wr *bufio.Writer) {
	fmt.Fprintf(wr, "static num fn_%s(", c.id)
	for i, v := range c.args {
		if i != 0 {
			wr.WriteRune(',')
		}
		fmt.Fprintf(wr, "num %s", v.unambiguousName())
	}
	fmt.Fprintf(wr, ")")
}

func (c cmdFnDef) generateCDeclaration(wr *bufio.Writer) {
	c.generateCFunctionHeader(wr)
	fmt.Fprintf(wr, ";\n")
}

func (c cmdFnDef) generateCDefinition(wr *bufio.Writer) {
	c.generateCFunctionHeader(wr)
	fmt.Fprintf(wr, "{\n")
	for _, v := range c.args {
		fmt.Fprintf(wr, "\t%s=%s;\n", v.unambiguousName(), v.unambiguousName())
	}
	fmt.Fprintf(wr, "\treturn ")
	c.expr.generateC(wr)
	fmt.Fprintf(wr, ";\n}\n\n")
}
