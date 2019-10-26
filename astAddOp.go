package main

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
		if !isAddOp(oper) || head.finalType() == strType && oper == '-' {
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

var nextTemp = 0

func createTemp() int {
	val := nextTemp
	nextTemp++
	return val
}
