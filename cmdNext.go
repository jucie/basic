package main

type cmdNext struct {
	vars []*astVarRef
}

func (p *parser) parseNext() *cmdNext {
	result := &cmdNext{}
	l := p.lex.peek()

	for {
		v := p.parseVarRef()
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

func (c cmdNext) receive(g guest) {
	for _, v := range c.vars {
		g.visit(*v)
	}
}
