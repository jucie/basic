package main

type astLit struct {
	val   string
	_type astType
}

func (p *parser) parseLit() *astLit {
	l := p.lex.peek()
	val := l.s
	switch l.token {
	case tokNumber:
		p.lex.next()
		return &astLit{val: val, _type: numType}
	case tokString:
		p.lex.next()
		return &astLit{val: val, _type: strType}
	}
	return nil
}

func (a astLit) receive(g guest) {
	g.visit(a._type)
}

func (a astLit) finalType() astType {
	return a._type
}
