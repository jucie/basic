package main

type subCmd interface {
	host
}

type cmdPrint struct {
	subCmds []subCmd
}

func (p *parser) parsePrint() *cmdPrint {
	result := &cmdPrint{}
	l := p.lex.peek()

Loop:
	for {
		var subCmd subCmd
		switch l.token {
		case ';':
			fallthrough
		case ',':
			subCmd = l.token
			p.lex.next()
		case tokEol:
			fallthrough
		case tokEof:
			fallthrough
		case ':':
			break Loop
		default:
			expr := p.parseExpr()
			if expr == nil {
				break Loop
			}
			subCmd = expr
		}
		result.subCmds = append(result.subCmds, subCmd)
	}
	return result
}

func (c cmdPrint) receive(g guest) {
	for _, subCmd := range c.subCmds {
		g.visit(subCmd)
	}
}
