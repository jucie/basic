package main

type cmdIf struct {
}

func (p *parser) parseIf() *cmdIf {
	p.consumeCmd()
	return &cmdIf{}
}
