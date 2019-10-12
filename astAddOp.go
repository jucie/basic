package main

import (
	"bufio"
	"fmt"
)

type astAddOpTail struct {
	oper token
	val  *astMulOp
}

type astAddOp struct {
	head *astMulOp
	tail []astAddOpTail
}

func (p *parser) parseAddOp() *astAddOp {
	head := p.parseMulOp()
	if head == nil {
		return nil
	}
	result := &astAddOp{head: head}

	for {
		oper := p.lex.peek().token
		if !isAddOp(oper) {
			break
		}
		p.lex.next()

		val := p.parseMulOp()
		if val == nil {
			return nil
		}
		result.tail = append(result.tail, astAddOpTail{oper: oper, val: val})
	}
	return result
}

func isAddOp(b token) bool {
	return b == '+' || b == '-'
}

func (a astAddOp) receive(g guest) {
	g.visit(a.head)
	for _, t := range a.tail {
		g.visit(t.val)
	}
}

func (a astAddOp) finalType() astType {
	if len(a.tail) == 0 {
		return a.head.finalType()
	}
	return numType
}

func (a astAddOp) generateC(wr *bufio.Writer) {
	a.head.generateC(wr)
	for _, t := range a.tail {
		t.generateC(wr)
	}
}

func (a astAddOpTail) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "%c", a.oper)
	a.val.generateC(wr)
}
