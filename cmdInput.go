package main

type cmdInput struct {
}

func (p *parser) parseInput() *cmdInput {
	p.consumeCmd()
	return &cmdInput{}
}
