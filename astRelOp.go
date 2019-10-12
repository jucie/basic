package main

import (
	"bufio"
	"fmt"
)

type astRelOp struct {
	lhs  *astAddOp
	rhs  *astAddOp
	oper token
}

func (p *parser) parseRelOp() *astRelOp {
	lhs := p.parseAddOp()
	if lhs == nil {
		return nil
	}

	oper := p.lex.peek().token
	if !isRelOp(oper) {
		return &astRelOp{lhs: lhs}
	}
	p.lex.next()

	rhs := p.parseAddOp()
	if rhs == nil {
		return nil
	}
	return &astRelOp{lhs: lhs, rhs: rhs, oper: oper}
}

func isRelOp(b token) bool {
	switch b {
	case '=':
		fallthrough
	case '<':
		fallthrough
	case '>':
		fallthrough
	case tokLe:
		fallthrough
	case tokGe:
		fallthrough
	case tokNe:
		return true
	}
	return false
}

func (a astRelOp) receive(g guest) {
	g.visit(a.lhs)
	if a.rhs != nil {
		g.visit(a.rhs)
	}
}

func (a astRelOp) generateC(wr *bufio.Writer) {
	a.lhs.generateC(wr)
	if a.rhs == nil {
		return
	}
	switch a.oper {
	case '=':
		fmt.Fprintf(wr, "==")
	case tokNe:
		fmt.Fprintf(wr, "!=")
	case tokLe:
		fmt.Fprintf(wr, "<=")
	case tokGe:
		fmt.Fprintf(wr, ">=")
	default:
		fmt.Fprintf(wr, "%c", a.oper)
	}
	a.rhs.generateC(wr)
}

func (a astRelOp) finalType() astType {
	if a.rhs != nil {
		return numType
	}
	return a.lhs.finalType()
}
