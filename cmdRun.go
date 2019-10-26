package main

type cmdRun struct {
	arg string
}

func (p *parser) parseRun() *cmdRun {
	l := p.lex.peek()
	result := &cmdRun{}

	if l.token != tokString {
		return nil
	}
	result.arg = l.s
	p.lex.next()

	return result
}

func (c cmdRun) receive(g guest) {
}
