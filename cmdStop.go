package main

type cmdStop struct {
}

func (p *parser) parseStop() *cmdStop {
	return &cmdStop{}
}

func (c cmdStop) receive(g guest) {
}
