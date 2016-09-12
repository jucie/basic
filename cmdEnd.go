package main

type cmdEnd struct {
}

func (p *parser) parseEnd() *cmdEnd {
	p.consumeCmd()
	return &cmdEnd{}
}
