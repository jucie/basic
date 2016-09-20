package main

import (
	"strconv"
)

type cmdOn struct {
	expr *astExpr
	dsts []int
}

func (p *parser) parseOn() *cmdOn {
	result := &cmdOn{}
	l := p.lex.peek()

	result.expr = p.parseExpr()
	if result.expr == nil {
		return nil
	}

	if l.token != tokGoto {
		return nil
	}
	p.lex.next()

	for l.token == tokNumber {
		dst, err := strconv.Atoi(l.s)
		if err != nil {
			return nil
		}
		result.dsts = append(result.dsts, dst)
		p.lex.next()
		if l.token != ',' {
			break
		}
		p.lex.next()
	}
	return result
}
