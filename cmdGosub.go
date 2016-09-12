package main

type cmdGosub struct {
}

func (p *parser) parseGosub() *cmdGosub {
	p.consumeCmd()
	return &cmdGosub{}
}
