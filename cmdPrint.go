package main

import (
	"bufio"
	"fmt"
)

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

func (scs printSubCmds) generateC(wr *bufio.Writer) {
	var _type astType
	shouldNL := true
	for _, subCmd := range scs {
		switch cmd := subCmd.(type) {
		case token:
			if cmd == ';' {
				shouldNL = false
			} else if cmd == ',' {
				shouldNL = false
				fmt.Fprintf(wr, "\tprint_char('\\t');\n")
			}
		case *astExpr:
			shouldNL = true
			_type = cmd.finalType()
			if _type == voidType {
				fmt.Fprintf(wr, "\t")
				cmd.generateC(wr)
				fmt.Fprintf(wr, ";\n")
			} else {
				fmt.Fprintf(wr, "\tprint_%s(", _type)
				cmd.generateC(wr)
				fmt.Fprintf(wr, ");\n")
			}
		}
	}
	if shouldNL {
		fmt.Fprintf(wr, "\tprint_char('\\n');\n")
	}
}
