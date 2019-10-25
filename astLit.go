package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type astLit struct {
	val   string
	_type astType
}

func (p *parser) parseLit() *astLit {
	l := p.lex.peek()
	val := l.s
	switch l.token {
	case tokNumber:
		p.lex.next()
		return &astLit{val: val, _type: numType}
	case tokString:
		p.lex.next()
		return &astLit{val: val, _type: strType}
	}
	return nil
}

func (a astLit) receive(g guest) {
	g.visit(a._type)
}

func (a astLit) generateC(wr *bufio.Writer) {
	switch a._type {
	case numType:
		f, err := strconv.ParseFloat(a.val, 32)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(wr, "%.4ff", f)
	case strType:
		wr.WriteRune('"')
		for _, r := range a.val {
			if r == '\'' || r == '\\' {
				wr.WriteRune('\\')
			}
			wr.WriteRune(r)
		}
		wr.WriteRune('"')
	}
}

func (a astLit) finalType() astType {
	return a._type
}
