package main

type cmdFor struct {
}

func (p *parser) parseFor() *cmdFor {
	p.consumeCmd()
	return &cmdFor{}
}
