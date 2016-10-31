package main

type cmdFor struct {
	index *astVarRef
	begin *astExpr
	end   *astExpr
	step  *astExpr
}

func (p *parser) parseFor() *cmdFor {
	result := &cmdFor{}
	l := p.lex.peek()

	if l.token != tokId {
		return nil
	}
	result.index = p.parseVarRef()
	if result.index == nil {
		return nil
	}

	if l.token != '=' {
		return nil
	}
	p.lex.next()

	result.begin = p.parseExpr()

	if l.token != tokTo {
		return nil
	}
	p.lex.next()

	result.end = p.parseExpr()

	if l.token == tokStep {
		p.lex.next()
		result.step = p.parseExpr()
	}

	return result
}

func (c cmdFor) receive(g guest) {
	g.visit(c.index)
	g.visit(c.begin)
	g.visit(c.end)
	if c.step != nil {
		g.visit(c.step)
	}
}
