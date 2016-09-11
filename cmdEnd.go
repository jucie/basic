package main

type cmdEnd struct {
}

func (p *parser) parseEnd() *cmdEnd {
	return &cmdEnd{}
}
