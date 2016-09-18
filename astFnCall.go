package main

type astFnCall struct {
	fn  string
	arg *astExpr
}

func (p *parser) parseFnCall() *astFnCall {
	l := p.lex.peek()

	if l.token != tokId {
		p.unexpected()
		return nil
	}
	fn := l.s
	p.lex.next()

	if l.token != '(' {
		p.unexpected()
		return nil
	}
	p.lex.next()

	arg := p.parseExpr()
	if arg == nil {
		return nil
	}

	if l.token != ')' {
		p.unexpected()
		return nil
	}
	p.lex.next()

	return &astFnCall{fn, arg}
}
