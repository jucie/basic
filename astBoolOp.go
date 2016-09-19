package main

type astBoolOp struct {
	lhs  *astRelOp
	rhs  *astRelOp
	oper token
}

func isBoolOp(op token) bool {
	return op == tokOr || op == tokAnd
}

func (p *parser) parseBoolOp() *astBoolOp {
	println(">parseBoolOp")
	defer println("<parseBoolOp")

	lhs := p.parseRelOp()

	oper := p.lex.peek().token
	if !isBoolOp(oper) {
		return &astBoolOp{lhs: lhs}
	}
	p.lex.next()

	rhs := p.parseRelOp()
	if rhs == nil {
		return nil
	}
	return &astBoolOp{lhs: lhs, rhs: rhs, oper: oper}
}
