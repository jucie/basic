package main

type cmdNext struct {
}

func (p *parser) parseNext() *cmdNext {
	p.consumeCmd()
	return &cmdNext{}
}
