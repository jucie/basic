package main

type cmdGoto struct {
	dst int
}

func (p *parser) parseGoto() *cmdGoto {
	p.consumeCmd()
	return &cmdGoto{}
}
