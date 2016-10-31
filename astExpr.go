package main

type astExpr struct {
	boolOp *astBoolOp
}

func (p *parser) parseExpr() *astExpr {
	boolOp := p.parseBoolOp()
	if boolOp == nil {
		return nil
	}
	return &astExpr{boolOp: boolOp}
}

func (a astExpr) receive(g guest) {
	g.visit(a.boolOp)
}
