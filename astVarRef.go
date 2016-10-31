package main

type astVarRef struct {
	coord
	id    string
	type_ astType
	index []*astExpr
}

func (p *parser) parseVarRef() *astVarRef {
	l := p.lex.peek()
	result := &astVarRef{coord: l.pos, type_: numType}

	if l.token != tokId {
		p.unexpected()
		return nil
	}
	result.id = l.s
	p.lex.next()

	if l.token == '$' {
		result.type_ = strType
		result.id += "$"
		p.lex.next()
	}

	if l.token == '(' {
		p.lex.next()

		for {
			expr := p.parseExpr()
			if expr == nil {
				break
			}
			result.index = append(result.index, expr)
			if l.token != ',' {
				break
			}
			p.lex.next()
		}

		if l.token != ')' {
			p.unexpected()
			return nil
		}
		p.lex.next()
	}

	return result
}

func (a astVarRef) receive(g guest) {
	g.visit(a.type_)
	for _, i := range a.index {
		g.visit(i)
	}
}
