package main

type cmdReturn struct {
}

func (p *parser) parseReturn() *cmdReturn {
	return &cmdReturn{}
}

func (c cmdReturn) receive(g guest) {
}
