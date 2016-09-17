package main

type astLit interface {
	Val() string
}

type astLitNumber struct {
	val string
}

func (l *astLitNumber) Val() string {
	return l.val
}

type astLitString struct {
	val string
}

func (l *astLitString) Val() string {
	return l.val
}

func (p *parser) parseLit() astLit {
	l := p.lex.peek()
	switch l.token {
	case tokNumber:
		return &astLitNumber{l.s}
	case tokString:
		return &astLitString{l.s}
	}
	return nil
}
