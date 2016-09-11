package main

type cmdRem struct {
}

func (p *parser) parseRem() *cmdRem {
	return &cmdRem{}
}
