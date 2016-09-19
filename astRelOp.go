package main

type astRelOp struct {
	lhs  *astAddOp
	rhs  *astAddOp
	oper token
}

func (p *parser) parseRelOp() *astRelOp {
	println(">parseRelOp")
	defer println("<parseRelOp")

	lhs := p.parseAddOp()

	oper := p.lex.peek().token
	if !isRelOp(oper) {
		return &astRelOp{lhs: lhs}
	}
	p.lex.next()

	rhs := p.parseAddOp()
	if rhs == nil {
		return nil
	}
	return &astRelOp{lhs: lhs, rhs: rhs, oper: oper}
}

func isRelOp(b token) bool {
	switch b {
	case '=':
		fallthrough
	case '<':
		fallthrough
	case '>':
		fallthrough
	case tokLe:
		fallthrough
	case tokGe:
		fallthrough
	case tokNe:
		return true
	}
	return false
}
