package main

type astPart struct {
	coord
	signal token
	val    astValue
}

func (p *parser) parsePart() *astPart {
	l := p.lex.peek()
	result := &astPart{coord: l.pos}

	switch l.token {
	case '-':
		fallthrough
	case '+':
		result.signal = l.token
		p.lex.next()
	}

	switch l.token {
	case tokID:
		result.val = p.parseVarRef()
	case '(':
		p.lex.next()
		result.val = p.parseExpr(true)
		if l.token != ')' {
			break
		}
		p.lex.next()
	case tokFn:
		p.lex.next()
		result.val = p.parseFnCall()
	case tokNumber:
		fallthrough
	case tokString:
		result.val = p.parseLit()
	case tokSin:
		fallthrough
	case tokCos:
		fallthrough
	case tokAtn:
		fallthrough
	case tokTan:
		fallthrough
	case tokStr:
		fallthrough
	case tokSqr:
		fallthrough
	case tokExp:
		fallthrough
	case tokLog:
		fallthrough
	case tokAbs:
		fallthrough
	case tokInt:
		fallthrough
	case tokRnd:
		fallthrough
	case tokSgn:
		fallthrough
	case tokVal:
		fallthrough
	case tokChr:
		fallthrough
	case tokMid:
		fallthrough
	case tokLeft:
		fallthrough
	case tokRight:
		fallthrough
	case tokLen:
		fallthrough
	case tokAsc:
		fallthrough
	case tokTab:
		result.val = p.parsePredef()

	default:
		p.unexpected()
		return nil
	}
	return result
}

func (a astPart) receive(g guest) {
	g.visit(a.val)
}

func (a astPart) finalType() astType {
	return a.val.finalType()
}
