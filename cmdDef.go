package main

type cmdDef struct {
	id   string
	arg  string
	expr *astExpr
}

func (p *parser) parseDef() *cmdDef {
	l := p.lex.peek()
	result := &cmdDef{}

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
