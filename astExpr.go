package main

type astExpr struct {
	boolOp *astBoolOp
	paren  bool
}

func (p *parser) parseExpr(paren bool) *astExpr {
	boolOp := p.parseBoolOp()
	if boolOp == nil {
		return nil
	}
	return &astExpr{boolOp: boolOp, paren: paren}
}

func (a astExpr) receive(g guest) {
	g.visit(a.boolOp)
}

func (a astExpr) finalType() astType {
	return a.boolOp.finalType()
}
