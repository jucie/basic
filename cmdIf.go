package main

import (
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
		_goto := &cmdGo{}
		_goto.dst.nbr = line
		result.cmds = append(result.cmds, _goto)
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
				return v.dst.adr.switchID
			}
		}
	}
	return 0
}
