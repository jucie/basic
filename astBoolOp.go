package main

import (
	"bufio"
	"fmt"
)

type astBoolOpTail struct {
	oper token
	val  *astRelOp
}
type astBoolOp struct {
	head *astRelOp
	tail []astBoolOpTail
}

func isBoolOp(op token) bool {
	return op == tokOr || op == tokAnd
}

func (p *parser) parseBoolOp() *astBoolOp {
	head := p.parseRelOp()
	if head == nil {
		return nil
	}
	result := &astBoolOp{head: head}

	for {
		oper := p.lex.peek().token
		if !isBoolOp(oper) {
			break
		}
		p.lex.next()

		val := p.parseRelOp()
		if val == nil {
			return nil
		}
		result.tail = append(result.tail, astBoolOpTail{oper: oper, val: val})
	}
	return result
}

func (a astBoolOp) receive(g guest) {
	g.visit(a.head)
	for _, t := range a.tail {
		g.visit(t.val)
	}
}

func (a astBoolOp) generateC(wr *bufio.Writer) {
	a.head.generateC(wr)
	for _, tail := range a.tail {
		tail.generateC(wr)
	}
}

func (a astBoolOp) finalType() astType {
	if len(a.tail) == 0 {
		return a.head.finalType()
	}
	return numType
}

func (a astBoolOpTail) generateC(wr *bufio.Writer) {
	switch a.oper {
	case tokOr:
		fmt.Fprintf(wr, "||")
	case tokAnd:
		fmt.Fprintf(wr, "&&")
	}
	a.val.generateC(wr)
}
