package main

type cmdRem struct {
	text string
}

func (p *parser) parseRem() *cmdRem {
	return &cmdRem{text: p.lex.peek().s}
}

func (c cmdRem) receive(g guest) {
}
