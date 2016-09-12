package main

type cmdOn struct {
	dst int
}

func (p *parser) parseOn() *cmdOn {
	p.consumeCmd()
	return &cmdOn{}
}
