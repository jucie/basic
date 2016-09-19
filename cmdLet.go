package main

type cmdLet struct {
	dst *astVarRef
	src *astExpr
}

func (p *parser) parseLet() *cmdLet {
	dst := p.parseVarRef()

	if p.lex.peek().token != '=' {
		p.unexpected()
		return nil
	}
	p.lex.next()

	src := p.parseExpr()

	return &cmdLet{dst: dst, src: src}
}
