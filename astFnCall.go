package main

import (
	"bufio"
	"fmt"
)

type astFnCall struct {
	id  string
	arg *astExpr
}

func (p *parser) parseFnCall() *astFnCall {
	l := p.lex.peek()

	if l.token != tokID {
		p.unexpected()
		return nil
	}
	id := l.s
	p.lex.next()

	if l.token != '(' {
		p.unexpected()
		return nil
	}
	p.lex.next()

	arg := p.parseExpr(false)
	if arg == nil {
		return nil
	}

	if l.token != ')' {
		p.unexpected()
		return nil
	}
	p.lex.next()

	return &astFnCall{id, arg}
}

func (a astFnCall) receive(g guest) {
	g.visit(a.arg)
}

func (a astFnCall) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "fn_%s(", a.id)
	a.arg.generateC(wr)
	fmt.Fprintf(wr, ")")
}

func (a astFnCall) finalType() astType {
	return numType
}
