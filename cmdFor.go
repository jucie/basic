package main

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
