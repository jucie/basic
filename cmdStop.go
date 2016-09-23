package main

type cmdStop struct {
}

func (p *parser) parseStop() *cmdStop {
	return &cmdStop{}
}
