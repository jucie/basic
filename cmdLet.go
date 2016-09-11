package main

type cmdLet struct {
}

func (p *parser) parseLet() *cmdLet {
	return &cmdLet{}
}
