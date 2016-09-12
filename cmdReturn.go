package main

type cmdReturn struct {
}

func (p *parser) parseReturn() *cmdReturn {
	p.consumeCmd()
	return &cmdReturn{}
}
