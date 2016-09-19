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
	println(">parseBoolOp")
	defer println("<parseBoolOp")

	head := p.parseRelOp()
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
