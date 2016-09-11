package main

type cmdPrint struct {
}

func (p *parser) parsePrint() *cmdPrint {
	return &cmdPrint{}
}
