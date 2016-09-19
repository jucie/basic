package main

type astAddOp struct {
	lhs  *astMulOp
	rhs  *astMulOp
	oper token
}

func (p *parser) parseAddOp() *astAddOp {
	println(">parseAddOp")
	defer println("<parseAddOp")

	lhs := p.parseMulOp()

	oper := p.lex.peek().token
	if !isAddOp(oper) {
		return &astAddOp{lhs: lhs}
	}
	p.lex.next()

	rhs := p.parseMulOp()
	if rhs == nil {
		return nil
	}
	return &astAddOp{lhs: lhs, rhs: rhs, oper: oper}
}

func isAddOp(b token) bool {
	return b == '+' || b == '-'
}
