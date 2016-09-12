package main

type cmdData struct {
}

func (p *parser) parseData() *cmdData {
	p.consumeCmd()
	return &cmdData{}
}
