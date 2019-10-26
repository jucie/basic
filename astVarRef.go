package main

import (
	"fmt"
	"strconv"
)

type astVarRef struct {
	coord
	id    string
	_type astType
	index []*astExpr
}

func (p *parser) parseVarRef() *astVarRef {
	l := p.lex.peek()
	result := &astVarRef{coord: l.pos, _type: numType}

	if l.token != tokID {
		p.unexpected()
		return nil
	}
	result.id = l.s
	p.lex.next()

	if l.token == '$' {
		result._type = strType
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
	g.visit(a._type)
	for _, i := range a.index {
		g.visit(i)
	}
}

func (a astVarRef) isArray() bool {
	return len(a.index) != 0
}

func (a astVarRef) finalType() astType {
	return a._type
}

func (a astVarRef) unambiguousName() string {
	name := fmt.Sprintf("%s_%s", a.id, a._type)
	if a.isArray() {
		name += "_array"
		if len(a.index) > 1 {
			name += strconv.Itoa(len(a.index))
		}
	}
	return name
}

func (a *astVarRef) equals(other *astVarRef) bool {
	return a.id == other.id && a._type == other._type && len(a.index) == len(other.index)
}
