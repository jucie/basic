package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type cmdIf struct {
	expr *astExpr
	cmds
}

func (p *parser) parseIf() *cmdIf {
	result := &cmdIf{}
	l := p.lex.peek()

	result.expr = p.parseExpr(false)
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
		goto_ := &cmdGo{}
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

// condBranchTarget return the target address of the conditional branch
// or zero if this command is not a conditional branch.
func (c *cmdIf) condBranchTarget() int {
	if len(c.cmds) == 1 { // has only 1 conditional command
		switch v := c.cmds[0].(type) {
		case *cmdGo: // it's a GO command
			if !v.sub { // it's not a GOSUB, so it's a GOTO
				return v.dst.nbr
			}
		}
	}
	return 0
}

func (c *cmdIf) generateC(wr *bufio.Writer) {
	label := c.condBranchTarget()
	if label != 0 {
		c.genCondBranch(wr, label)
	} else {
		c.genRegularIf(wr)
	}
}

func (c *cmdIf) genCondBranch(wr *bufio.Writer, label int) {
	fmt.Fprintf(wr, "\tif (")
	c.expr.generateC(wr)
	fmt.Fprintf(wr, "){ target = %d; break; }\n", label)
}

func (c *cmdIf) genRegularIf(wr *bufio.Writer) {
	label := createLabel()
	fmt.Fprintf(wr, "\tif (!(")
	c.expr.generateC(wr)
	fmt.Fprintf(wr, ")){ target = %d; break; }\n", label)
	c.cmds.generateC(wr)
	fmt.Fprintf(wr, "case %d:\n", label)
}
