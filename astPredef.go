package main

import (
	"bufio"
	"fmt"
)

type astPredef struct {
	function token
	args     []*astExpr
}

func (p *parser) parsePredef() *astPredef {
	l := p.lex.peek()
	function := l.token
	p.lex.next()
	result := &astPredef{function: function}

	if l.token == '$' {
		p.lex.next()
	}

	if l.token != '(' {
		p.unexpected()
		return nil
	}
	p.lex.next()

	for {
		arg := p.parseExpr(false)
		if arg == nil {
			return nil
		}
		result.args = append(result.args, arg)
		if l.token == ',' {
			p.lex.next()
			continue
		}
		if l.token == ')' {
			p.lex.next()
			break
		}
	}
	return result
}

func (a astPredef) receive(g guest) {
	for _, arg := range a.args {
		g.visit(arg)
	}
}

func (a astPredef) generateC(wr *bufio.Writer) {
	predef := predefs[a.function]
	fmt.Fprintf(wr, "%s_%s(", predef.name, predef.type_)
	if predef.type_ == strType {
		fmt.Fprintf(wr, "temp_str[%d],", createTemp())
	}
	for i, arg := range a.args {
		if i != 0 {
			wr.WriteRune(',')
		}
		arg.generateC(wr)
	}
	wr.WriteRune(')')
}

func (a astPredef) finalType() astType {
	return predefs[a.function].type_
}

type predef struct {
	token
	name   string
	type_  astType
	params []astType
}

var predefs = make(map[token]predef)

func init() {
	var predefInfo = []predef{
		{tokAbs, "ABS", numType, []astType{numType}},
		{tokAsc, "ASC", numType, []astType{strType}},
		{tokAtn, "ATN", numType, []astType{numType}},
		{tokChr, "CHR", strType, []astType{numType}},
		{tokCos, "COS", numType, []astType{numType}},
		{tokExp, "EXP", numType, []astType{numType}},
		{tokInt, "INT", numType, []astType{numType}},
		{tokLeft, "LEFT", strType, []astType{strType, numType}},
		{tokLen, "LEN", numType, []astType{strType}},
		{tokLog, "LOG", numType, []astType{numType}},
		{tokMid, "MID", strType, []astType{strType, numType, numType}},
		{tokRight, "RIGHT", strType, []astType{strType, numType}},
		{tokRnd, "RND", numType, []astType{numType}},
		{tokSgn, "SGN", numType, []astType{numType}},
		{tokSin, "SIN", numType, []astType{numType}},
		{tokSqr, "SQR", numType, []astType{numType}},
		{tokStr, "STR", strType, []astType{numType}},
		{tokTab, "TAB", voidType, []astType{numType}},
		{tokTan, "TAN", numType, []astType{numType}},
		{tokVal, "VAL", numType, []astType{strType}},
	}
	for _, p := range predefInfo {
		predefs[p.token] = p
	}
}
