package main

type cmdRem struct {
}

func (p *parser) parseRem() *cmdRem {
	p.consumeCmd()
	return &cmdRem{}
}
