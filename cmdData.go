package main

type cmdData struct {
	line string
}

func (p *parser) parseData() *cmdData {
	line := p.lex.previous.s
	return &cmdData{line: line}
}
