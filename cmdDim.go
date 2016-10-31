package main

type cmdDim struct {
	vars []*astVarRef
}

func (p *parser) parseDim() *cmdDim {
	result := &cmdDim{}
	l := p.lex.peek()

	for {
		v := p.parseVarRef()
		if v == nil {
			break
		}
		result.vars = append(result.vars, v)
		if l.token != ',' {
			break
		}
		p.lex.next()
	}
	if len(result.vars) == 0 {
		return nil
	}
	return result
}

func (c cmdDim) receive(g guest) {
	for _, v := range c.vars {
		g.visit(v)
	}
}
