package main

type astPredef struct {
	function token
	args     []*astExpr
}

func (p *parser) parsePredef() *astPredef {
	l := p.lex.peek()
	function := l.token
	p.lex.next()
	result := &astPredef{function: function}

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
