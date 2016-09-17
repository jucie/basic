package main

type astPart struct {
	coord
	signal bool
	val    astValue
}

func (p *parser) parsePart() *astPart {
	l := p.lex.peek()
	result := &astPart{coord: l.pos}

	switch l.token {
	case '-':
		result.signal = true // negative
		p.lex.next()
	case '+':
		result.signal = false // positive
		p.lex.next()
	}

	switch l.token {
	case tokId:
		result.val = p.parseVarRef()
	case '(':
		p.lex.next()
		result.val = p.parseExpr()
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
	default:
		p.unexpected()
	}
	return result
}
