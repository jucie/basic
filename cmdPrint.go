package main

type subCmd interface {
	host
}
type printSubCmds []subCmd

type cmdPrint struct {
	printSubCmds
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
		case tokEOL:
			fallthrough
		case tokEOF:
			fallthrough
		case ':':
			break Loop
		default:
			expr := p.parseExpr(false)
			if expr == nil {
				break Loop
			}
			subCmd = expr
		}
		result.printSubCmds = append(result.printSubCmds, subCmd)
	}
	return result
}

func (c cmdPrint) receive(g guest) {
	for _, subCmd := range c.printSubCmds {
		g.visit(subCmd)
	}
}
