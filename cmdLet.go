package main

type cmdLet struct {
}

func (p *parser) parseLet() *cmdLet {
	p.consumeCmd()
	return &cmdLet{}
}
