package main

type cmdRun struct {
	arg string
}

func (p *parser) parseRun() *cmdRun {
	p.consumeCmd()
	return &cmdRun{}
}
