package main

import (
	"bufio"
)

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

func (c cmdData) generateC(wr *bufio.Writer) {
	//does nothing. Data is not generated inline.
}

func (c cmdData) generateCDefinition(wr *bufio.Writer, type_ astType) {
	emittedSome := false
	//does nothing. Data is not generated here.
	for _, v := range c.values {
		if v.finalType() == type_ {
			v.generateC(wr)
			wr.WriteRune(',')
			emittedSome = true
		}
	}
	if emittedSome {
		wr.WriteRune('\n')
	}
}
