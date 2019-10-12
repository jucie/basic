package main

import "bufio"

type astExpOp struct {
	head *astPart
	tail []*astPart
}

func (p *parser) parseExpOp() *astExpOp {
	head := p.parsePart()
	if head == nil {
		return nil
	}
	result := &astExpOp{head: head}

	for {
		oper := p.lex.peek().token
		if oper != '^' {
			break
		}
		p.lex.next()

		val := p.parsePart()
		if val == nil {
			return nil
		}
		result.tail = append(result.tail, val)
	}
	return result
}

func (a astExpOp) receive(g guest) {
	g.visit(a.head)
	for _, t := range a.tail {
		g.visit(t)
	}
}

func (a astExpOp) generateC(wr *bufio.Writer) {
	a.head.generateC(wr)
	for _, t := range a.tail {
		t.generateC(wr)
	}
}

func (a astExpOp) finalType() astType {
	if len(a.tail) == 0 {
		return a.head.finalType()
	}
	return numType
}
