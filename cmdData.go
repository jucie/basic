package main

import (
	"strings"
)

type cmdData struct {
	values []string
}

func (p *parser) parseData() *cmdData {
	result := &cmdData{}
	line := p.lex.previous.s
	v := strings.Split(line, ",")
	for _, s := range v {
		result.values = append(result.values, strings.Trim(s, " \"\t"))
	}
	return result
}

func (c cmdData) receive(g guest) {
}
