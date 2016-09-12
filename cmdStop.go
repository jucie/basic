package main

type cmdStop struct {
}

func (p *parser) parseStop() *cmdStop {
	p.consumeCmd()
	return &cmdStop{}
}
