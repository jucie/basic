package main

type cmdPrint struct {
}

func (p *parser) parsePrint() *cmdPrint {
	p.consumeCmd()
	return &cmdPrint{}
}
