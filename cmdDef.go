package main

type cmdFnDef struct {
	id   string
	arg  string
	expr *astExpr
}

func (p *parser) parseDef() *cmdFnDef {
	l := p.lex.peek()
	result := &cmdFnDef{}

	if l.token != tokFn {
		return nil
	}
	p.lex.next()

	if l.token != tokId {
		return nil
	}
	result.id = l.s
	p.lex.next()

	if l.token != '(' {
		return nil
	}
	p.lex.next()

	if l.token != tokId {
		return nil
	}
	result.arg = l.s
	p.lex.next()

	if l.token != ')' {
		return nil
	}
	p.lex.next()

	if l.token != '=' {
		return nil
	}
	p.lex.next()

	result.expr = p.parseExpr()

	return result
}

func (c cmdFnDef) receive(g guest) {
	g.visit(c.expr)
}
