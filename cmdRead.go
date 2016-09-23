package main

type cmdRead struct {
	vars []*astVarRef
}

func (p *parser) parseRead() *cmdRead {
	result := &cmdRead{}
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
