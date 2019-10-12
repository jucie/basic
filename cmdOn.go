package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type cmdOn struct {
	expr *astExpr
	dsts []targetLine
}

func (p *parser) parseOn() *cmdOn {
	result := &cmdOn{}
	l := p.lex.peek()

	result.expr = p.parseExpr()
	if result.expr == nil {
		return nil
	}

	if l.token != tokGo {
		return nil
	}
	p.lex.next()
	if l.token != tokTo {
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

func (c cmdOn) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "\tswitch (")
	c.expr.generateC(wr)
	fmt.Fprintf(wr, "){\n")
	for i, line := range c.dsts {
		fmt.Fprintf(wr, "\t\tcase %d: return block_%d_1;\n", i, line.nbr)
	}
	fmt.Fprintf(wr, "\t}\n")
}
