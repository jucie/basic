package main

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

func (a astBoolOp) finalType() astType {
	if len(a.tail) == 0 {
		return a.head.finalType()
	}
	return numType
}
