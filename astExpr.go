package main

type astExpr struct {
	part *astPart
}

func (p *parser) parseExpr() *astExpr {
	part := p.parsePart()
	if part == nil {
		return nil
	}
	return &astExpr{part: part}
}
