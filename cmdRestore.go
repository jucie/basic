package main

type cmdRestore struct {
}

func (p *parser) parseRestore() *cmdRestore {
	p.consumeCmd()
	return &cmdRestore{}
}
