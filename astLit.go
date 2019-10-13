package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type astLit struct {
	val   string
	type_ astType
}

func (p *parser) parseLit() *astLit {
	l := p.lex.peek()
	val := l.s
	switch l.token {
	case tokNumber:
		p.lex.next()
		return &astLit{val: val, type_: numType}
	case tokString:
		p.lex.next()
		return &astLit{val: val, type_: strType}
	}
	return nil
}

func (a astLit) receive(g guest) {
	g.visit(a.type_)
}

func (a astLit) generateC(wr *bufio.Writer) {
	switch a.type_ {
	case numType:
		f, err := strconv.ParseFloat(a.val, 32)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(wr, "%.4ff", f)
	case strType:
		fmt.Fprintf(wr, "\"%s\"", a.val)
	}
}

func (a astLit) finalType() astType {
	return a.type_
}
