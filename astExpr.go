package main

import "bufio"

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

func (a astExpr) generateC(wr *bufio.Writer) {
	a.boolOp.generateC(wr)
}

func (a astExpr) finalType() astType {
	return a.boolOp.finalType()
}
