package main

type astPredef struct {
	function token
	arg      *astExpr
}

func (p *parser) parsePredef() *astPredef {
	l := p.lex.peek()
	function := l.token
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

	return &astPredef{function: function, arg: arg}
}
