package main

type cmdAssign struct {
	rhs *astExpr
	lhs *astExpr
}

func (p *parser) parseAssign() *cmdAssign {
	p.consumeCmd()
	return &cmdAssign{}
}
