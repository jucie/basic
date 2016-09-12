package main

type cmdDim struct {
}

func (p *parser) parseDim() *cmdDim {
	p.consumeCmd()
	return &cmdDim{}
}
