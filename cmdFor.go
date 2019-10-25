package main

import (
	"bufio"
	"fmt"
)

type cmdFor struct {
	index *astVarRef
	begin *astExpr
	end   *astExpr
	step  *astExpr
	next  *cmdNext
}

var step1 = &astExpr{
	boolOp: &astBoolOp{
		head: &astRelOp{
			lhs: &astAddOp{
				head: &astMulOp{
					head: &astExpOp{
						lhs: &astPart{
							val: &astLit{val: "1", _type: numType},
						},
					},
				},
			},
		},
	},
}

func (p *parser) parseFor() *cmdFor {
	result := &cmdFor{}
	l := p.lex.peek()

	if l.token != tokID {
		return nil
	}
	result.index = p.parseVarRef()
	if result.index == nil {
		return nil
	}

	if l.token != '=' {
		return nil
	}
	p.lex.next()

	result.begin = p.parseExpr(false)

	if l.token != tokTo {
		return nil
	}
	p.lex.next()

	result.end = p.parseExpr(false)

	if l.token == tokStep {
		p.lex.next()
		result.step = p.parseExpr(false)
	} else {
		result.step = step1
	}

	return result
}

func (c cmdFor) receive(g guest) {
	g.visit(c.index)
	g.visit(c.begin)
	g.visit(c.end)
	if c.step != nil {
		g.visit(c.step)
	}
}

func (c cmdFor) generateC(wr *bufio.Writer) {
	labelInc := createLabel()
	labelCond := createLabel()

	fmt.Fprintf(wr, "\t%s_target = %d;\n", c.index.unambiguousName(), labelInc)

	// first index attribution
	wr.WriteRune('\t')
	c.index.generateCLValue(wr, "let")
	wr.WriteRune(',')
	c.begin.generateC(wr)
	fmt.Fprintf(wr, ");\n")
	fmt.Fprintf(wr, "\ttarget = %d; break;\n", labelCond)

	// Increment
	fmt.Fprintf(wr, "case %d:\n", labelInc)
	wr.WriteRune('\t')
	c.index.generateCLValue(wr, "let")
	wr.WriteRune(',')
	c.index.generateC(wr)
	wr.WriteRune('+')
	c.step.generateC(wr)
	fmt.Fprintf(wr, ");\n")

	// index value bounds checking
	fmt.Fprintf(wr, "case %d:\n", labelCond)
	if c.step == step1 {
		fmt.Fprintf(wr, "\tif (")
		c.index.generateC(wr)
		fmt.Fprintf(wr, " > ")
		c.end.generateC(wr)
		fmt.Fprintf(wr, ") { target=%d; break; }\n", c.next.labelExit)
		return
	}
	fmt.Fprintf(wr, "\tif (")
	c.step.generateC(wr)
	fmt.Fprintf(wr, " > 0 && ")
	c.index.generateC(wr)
	fmt.Fprintf(wr, " > ")
	c.end.generateC(wr)
	fmt.Fprintf(wr, ") { target=%d; break; }\n", c.next.labelExit)
	fmt.Fprintf(wr, "\telse if (")
	c.step.generateC(wr)
	fmt.Fprintf(wr, " < 0 && ")
	c.index.generateC(wr)
	fmt.Fprintf(wr, " < ")
	c.end.generateC(wr)
	fmt.Fprintf(wr, ") { target=%d; break; }\n", c.next.labelExit)
}
