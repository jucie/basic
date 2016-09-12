package main

type cmdRead struct {
}

func (p *parser) parseRead() *cmdRead {
	p.consumeCmd()
	return &cmdRead{}
}
