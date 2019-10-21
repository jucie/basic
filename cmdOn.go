package main

import (
	"bufio"
	"fmt"
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

func (c cmdOn) generateC(wr *bufio.Writer) {
	labelExit := createLabel()
	fmt.Fprintf(wr, "\ttarget = (int)(")
	c.expr.generateC(wr)
	fmt.Fprintf(wr, ");\n")
	fmt.Fprintf(wr, "\tif (target < 1 || target > %d) {target = %d; break;}\n", len(c.dsts), labelExit)
	if c.sub {
		fmt.Fprintf(wr, "\tpush_address(%d);\n", labelExit)
	}
	fmt.Fprintf(wr, "\ttarget = (const int[]){")
	for i, line := range c.dsts {
		if i != 0 {
			wr.WriteRune(',')
		}
		fmt.Fprintf(wr, "%d", line.adr.switchID)
	}
	fmt.Fprintf(wr, "}[target -1]; break;\n")
	fmt.Fprintf(wr, "case %d:\n", labelExit)
}
