package main

import (
	"bufio"
	"fmt"
)

type astVarRef struct {
	coord
	id    string
	type_ astType
	index []*astExpr
}

func (p *parser) parseVarRef() *astVarRef {
	l := p.lex.peek()
	result := &astVarRef{coord: l.pos, type_: numType}

	if l.token != tokID {
		p.unexpected()
		return nil
	}
	result.id = l.s
	p.lex.next()

	if l.token == '$' {
		result.type_ = strType
		p.lex.next()
	}

	if l.token == '(' {
		p.lex.next()

		for {
			expr := p.parseExpr(false)
			if expr == nil {
				break
			}
			result.index = append(result.index, expr)
			if l.token != ',' {
				break
			}
			p.lex.next()
		}

		if l.token != ')' {
			p.unexpected()
			return nil
		}
		p.lex.next()
	}

	return result
}

func (a astVarRef) receive(g guest) {
	g.visit(a.type_)
	for _, i := range a.index {
		g.visit(i)
	}
}

func (a astVarRef) generateC(wr *bufio.Writer) {
	a.generateCVarRef(wr, true)
}

/*
					array        not array     const lit
As an R value
		  str       *s(1)        s             "abc"
		  num       *n(1)        n             1.0000f

As an L value
		  str       let_str(s(1),    let_str(&s,
		  num       let_num(n(1),    let_num(&n,
*/

// generateCVarRef emits C code for access to a variable.
// The intricacies are due to the subtleties of several variable types.
func (a astVarRef) generateCVarRef(wr *bufio.Writer, shouldDeref bool) {
	if !a.isArray() {
		fmt.Fprintf(wr, "%s_%s", a.id, a.type_)
		return
	}
	if shouldDeref {
		wr.WriteRune('*')
	}
	fmt.Fprintf(wr, "%s_%s", a.id, a.type_)
	wr.WriteRune('(')
	for i, v := range a.index {
		if i != 0 {
			wr.WriteRune(',')
		}
		v.generateC(wr)
	}
	wr.WriteRune(')')
}

// gererateCLValue generates code for the left side of an assignment.
// returns a bool indicating wether a closing parenthesis is needed
// The intricacies are due to the subtleties of several variable types.
func (a astVarRef) generateCLValue(wr *bufio.Writer, fname string) {
	fmt.Fprintf(wr, "%s_%s(", fname, a.finalType())
	if !a.isArray() {
		wr.WriteRune('&')
	}
	a.generateCVarRef(wr, false)
}

func (a astVarRef) isArray() bool {
	return len(a.index) != 0
}

func (a astVarRef) finalType() astType {
	return a.type_
}
