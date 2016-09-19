package main

type astExpr struct {
	boolOp *astBoolOp
}

func (p *parser) parseExpr() *astExpr {
	println(">parseExpr")
	defer println("<parseExpr")

	boolOp := p.parseBoolOp()
	if boolOp == nil {
		return nil
	}
	return &astExpr{boolOp: boolOp}
}
