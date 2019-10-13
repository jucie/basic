package main

import (
	"bufio"
	"fmt"
)

type astVarRef struct {
	coord
	id    string
	type_ astType
	index []*astExpr
}

func (p *parser) parseVarRef() *astVarRef {
	l := p.lex.peek()
	result := &astVarRef{coord: l.pos, type_: numType}

	if l.token != tokID {
		p.unexpected()
		return nil
	}
	result.id = l.s
	p.lex.next()

	if l.token == '$' {
		result.type_ = strType
		p.lex.next()
	}

	if l.token == '(' {
		p.lex.next()

		for {
			expr := p.parseExpr(false)
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

func (a astVarRef) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "%s_%s", a.id, a.type_)
	if len(a.index) != 0 {
		wr.WriteRune('(')
		for i, v := range a.index {
			if i != 0 {
				wr.WriteRune(',')
			}
			v.generateC(wr)
		}
		wr.WriteRune(')')
	}
}

func (a astVarRef) finalType() astType {
	return a.type_
}
