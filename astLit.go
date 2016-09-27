package main

type astLit struct {
	val   string
	type_ astType
}

func (p *parser) parseLit() *astLit {
	l := p.lex.peek()
	val := l.s
	switch l.token {
	case tokNumber:
		p.lex.next()
		return &astLit{val: val, type_: numType}
	case tokString:
		p.lex.next()
		return &astLit{val: val, type_: strType}
	}
	return nil
}

func (a astLit) receive(g guest) {
	g.visit(a.type_)
}
