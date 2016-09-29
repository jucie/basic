package main

import (
	"strconv"
)

type cmdIf struct {
	expr *astExpr
	cmds []cmd
}

func (p *parser) parseIf() *cmdIf {
	result := &cmdIf{}
	l := p.lex.peek()

	result.expr = p.parseExpr()
	if result.expr == nil {
		return nil
	}

	if l.token != tokThen {
		return nil
	}
	p.lex.next()

	if l.token == tokNumber {
		line, err := strconv.Atoi(l.s)
		if err != nil {
			return nil
		}
		p.lex.next()
		goto_ := cmdGoto{}
		goto_.dst.nbr = line
		result.cmds = append(result.cmds, goto_)
	} else {
		result.cmds = p.parseLineTail()
	}

	return result
}

func (c *cmdIf) receive(g guest) {
	g.visit(c.expr)
	for _, cmd := range c.cmds {
		g.visit(cmd)
	}
}
