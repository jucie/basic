package main

type cmdNext struct {
	ids []string
}

func (p *parser) parseNext() *cmdNext {
	result := &cmdNext{}
	l := p.lex.peek()

	for l.token == tokId {
		result.ids = append(result.ids, l.s)
		p.lex.next()
		if l.token == ',' {
			p.lex.next()
		}
	}
	if len(result.ids) == 0 {
		return nil
	}

	return result
}
