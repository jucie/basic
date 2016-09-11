package main

type cmdGoto struct {
}

func (p *parser) parseGoto() *cmdGoto {
	return &cmdGoto{}
}
