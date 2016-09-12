package main

type cmdDef struct {
}

func (p *parser) parseDef() *cmdDef {
	p.consumeCmd()
	return &cmdDef{}
}
