package main

import (
	"bufio"
	"fmt"
)

type astMulOpTail struct {
	oper token
	val  *astExpOp
}

type astMulOp struct {
	head *astExpOp
	tail []astMulOpTail
}

func (p *parser) parseMulOp() *astMulOp {
	head := p.parseExpOp()
	if head == nil {
		return nil
	}
	result := &astMulOp{head: head}

	for {
		oper := p.lex.peek().token
		if !isMulOp(oper) {
			break
		}
		p.lex.next()

		val := p.parseExpOp()
		if val == nil {
			return nil
		}
		result.tail = append(result.tail, astMulOpTail{oper: oper, val: val})
	}
	return result
}

func isMulOp(b token) bool {
	return b == '*' || b == '/'
}

func (a astMulOp) receive(g guest) {
	g.visit(a.head)
	for _, t := range a.tail {
		g.visit(t.val)
	}
}

func (a astMulOp) finalType() astType {
	if len(a.tail) == 0 {
		return a.head.finalType()
	}
	return numType
}

func (a astMulOp) generateC(wr *bufio.Writer) {
	a.head.generateC(wr)
	for _, t := range a.tail {
		t.generateC(wr)
	}
}

func (a astMulOpTail) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "%c", a.oper)
	a.val.generateC(wr)
}
