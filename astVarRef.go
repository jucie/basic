package main

type astVarRef struct {
	coord
	id    string
	type_ astType
	index []*astExpr
}

func (p *parser) parseVarRef() *astVarRef {
	result := &astVarRef{coord: p.lex.pos, type_: numType}

	l := p.lex.peek()
	if l.token != tokId {
		p.unexpected()
		return nil
	}
	result.id = l.s
	p.lex.addId(result.id)
	p.lex.next()

	if l.token == '$' {
		result.type_ = strType
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
