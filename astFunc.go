package main

type astFunc struct {
	id   string
	arg  string
	expr *astExpr
}

func (p *parser) parseFunc() *astFunc {
	result := &astFunc{}
	l := p.lex.peek()

	if l.token != tokId {
		p.unexpected()
		return nil
	}
	result.id = l.s
	p.lex.next()

	if l.token != '(' {
		p.unexpected()
		return nil
	}
	p.lex.next()

	if l.token != tokId {
		p.unexpected()
		return nil
	}
	result.arg = l.s
	p.lex.next()

	if l.token != ')' {
		p.unexpected()
		return nil
	}
	p.lex.next()

	if l.token != '=' {
		p.unexpected()
		return nil
	}
	p.lex.next()

	result.expr = p.parseExpr()
	if result.expr == nil {
		return nil
	}

	return result
}
