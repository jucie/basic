package main

type astExpOp struct {
	lhs *astPart
	rhs *astPart
}

func (p *parser) parseExpOp() *astExpOp {
	lhs := p.parsePart()
	if lhs == nil {
		return nil
	}
	result := &astExpOp{lhs: lhs}

	oper := p.lex.peek().token
	if oper == '^' {
		p.lex.next()

		rhs := p.parsePart()
		if rhs == nil {
			return nil
		}
		result.rhs = rhs
	}

	return result
}

func (a astExpOp) receive(g guest) {
	g.visit(a.lhs)
	if a.rhs != nil {
		g.visit(a.rhs)
	}
}

func (a astExpOp) finalType() astType {
	if a.rhs == nil {
		return a.lhs.finalType()
	}
	return numType
}
