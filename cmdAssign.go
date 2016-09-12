package main

type cmdAssign struct {
}

func (p *parser) parseAssign() *cmdAssign {
	p.consumeCmd()
	return &cmdAssign{}
}
