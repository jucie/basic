package main

type astMulOp struct {
	lhs  *astPart
	rhs  *astPart
	oper token
}

func (p *parser) parseMulOp() *astMulOp {
	println(">parseMulOp")
	defer println("<parseMulOp")

	lhs := p.parsePart()

	oper := p.lex.peek().token
	if !isMulOp(oper) {
		return &astMulOp{lhs: lhs}
	}
	p.lex.next()

	rhs := p.parsePart()
	if rhs == nil {
		return nil
	}
	return &astMulOp{lhs: lhs, rhs: rhs, oper: oper}
}

func isMulOp(b token) bool {
	return b == '*' || b == '/'
}
