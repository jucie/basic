package main

type cmdData struct {
	values []*astPart
}

func (p *parser) parseData() *cmdData {
	l := p.lex.peek()
	p.lex.next()
	result := &cmdData{}
	for {
		v := p.parsePart()
		if v == nil {
			break
		}
		p.incrementDataCounter(v.finalType())
		result.values = append(result.values, v)
		if l.token != ',' {
			break
		}
		p.lex.next()
	}
	return result
}

func (c cmdData) receive(g guest) {
}
