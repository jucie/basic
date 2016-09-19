package main

type cmdDim struct {
	vars []*astVarRef
}

func (p *parser) parseDim() *cmdDim {
	result := &cmdDim{}
	l := p.lex.peek()

	for {
		v := p.parseVarRef()
		result.vars = append(result.vars, v)
		if l.token != ',' {
			break
		}
		p.lex.next()
	}
	return result
}
