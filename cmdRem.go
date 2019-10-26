package main

type cmdRem struct {
	text string
}

func (p *parser) parseRem() *cmdRem {
	text := p.lex.peek().s
	if len(text) > 0 {
		if text[0] != ' ' {
			text = "REM" + text
		}
	}
	return &cmdRem{text: text}
}

func (c cmdRem) receive(g guest) {
}
