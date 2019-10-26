package main

import (
	"strconv"
)

type cmdOn struct {
	expr *astExpr
	dsts []targetLine
	sub  bool
}

func (p *parser) parseOn() *cmdOn {
	result := &cmdOn{}
	l := p.lex.peek()

	result.expr = p.parseExpr(false)
	if result.expr == nil {
		return nil
	}

	if l.token != tokGo {
		return nil
	}
	p.lex.next()
	if l.token == tokSub {
		result.sub = true
	} else if l.token != tokTo {
		return nil
	}
	p.lex.next()

	for l.token == tokNumber {
		dst, err := strconv.Atoi(l.s)
		if err != nil {
			return nil
		}
		t := targetLine{}
		t.nbr = dst
		result.dsts = append(result.dsts, t)
		p.lex.next()
		if l.token != ',' {
			break
		}
		p.lex.next()
	}
	return result
}

func (c cmdOn) receive(g guest) {
	g.visit(c.expr)
}
