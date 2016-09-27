package main

type astPredef struct {
	function token
	type_    astType
	args     []*astExpr
}

func (p *parser) parsePredef() *astPredef {
	l := p.lex.peek()
	function := l.token
	p.lex.next()
	result := &astPredef{function: function, type_: numType}

	if l.token == '$' {
		result.type_ = strType
		p.lex.next()
	}

	if l.token != '(' {
		p.unexpected()
		return nil
	}
	p.lex.next()

	for {
		arg := p.parseExpr()
		result.args = append(result.args, arg)
		if l.token == ',' {
			p.lex.next()
			continue
		}
		if l.token == ')' {
			p.lex.next()
			break
		}
	}
	return result
}

func (a astPredef) receive(g guest) {
	g.visit(a.type_)
	for _, arg := range a.args {
		g.visit(*arg)
	}
}
