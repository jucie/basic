package main

type cmdRead struct {
}

func (p *parser) parseRead() *cmdRead {
	return &cmdRead{}
}
