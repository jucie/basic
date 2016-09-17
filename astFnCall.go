package main

type astFnCall struct {
	fn  string
	arg *astExpr
}

func (p *parser) parseFnCall() *astFnCall {
	l := p.lex.peek()

	if l.token != tokId {
		return nil
	}
	fn := l.s
	arg := p.parseExpr()
	if arg == nil {
		return nil
	}
	return &astFnCall{fn, arg}
}
