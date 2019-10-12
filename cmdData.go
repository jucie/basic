package main

import (
	"bufio"
	"strings"
)

type cmdData struct {
	values []string
}

func (p *parser) parseData() *cmdData {
	line := p.lex.peek().s
	p.lex.next()
	result := &cmdData{}
	v := strings.Split(line, ",")
	for _, s := range v {
		result.values = append(result.values, strings.Trim(s, " \"\t"))
	}
	return result
}

func (c cmdData) receive(g guest) {
}

func (c cmdData) generateC(wr *bufio.Writer) {
	//does nothing. Data is not generated here.
}
