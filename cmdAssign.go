package main

type cmdAssign struct {
}

func (p *parser) parseAssign() *cmdAssign {
	return &cmdAssign{}
}
