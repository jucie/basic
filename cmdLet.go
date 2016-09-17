package main

type cmdLet struct {
	dst *astVarRef
	src *astExpr
}

func (p *parser) parseLet() *cmdLet {
	dst := p.parseVarRef()
	src := p.parseExpr()
	return &cmdLet{dst: dst, src: src}
}
