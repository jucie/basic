package main

type cmdFnDef struct {
	id   string
	args []*astVarRef
	expr *astExpr
}

func (p *parser) parseDef() *cmdFnDef {
	l := p.lex.peek()
	result := &cmdFnDef{}

	if l.token != tokFn {
		return nil
	}
	p.lex.next()

	if l.token != tokID {
		return nil
	}
	result.id = l.s
	p.lex.next()

	if l.token != '(' {
		return nil
	}
	p.lex.next()

	if l.token != tokID {
		return nil
	}
	for {
		v := p.parseVarRef()
		if v == nil {
			break
		}
		result.args = append(result.args, v)
		if l.token != ',' {
			break
		}
		p.lex.next()
	}

	if l.token != ')' {
		return nil
	}
	p.lex.next()

	if l.token != '=' {
		return nil
	}
	p.lex.next()

	result.expr = p.parseExpr(false)

	return result
}

func (c cmdFnDef) receive(g guest) {
	g.visit(c.expr)
}
