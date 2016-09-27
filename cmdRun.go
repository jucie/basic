package main

type cmdRun struct {
	arg string
}

func (p *parser) parseRun() *cmdRun {
	l := p.lex.peek()

	if l.token != tokString {
		return nil
	}
	return &cmdRun{arg: l.s}
}

func (c cmdRun) receive(g guest) {
}
