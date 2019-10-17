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
	var type_ astType
	previousIsSemicolon := false
	for _, subCmd := range scs {
		switch cmd := subCmd.(type) {
		case token:
			if cmd == ';' {
				previousIsSemicolon = true
			} else {
				previousIsSemicolon = false
				fmt.Fprintf(wr, "\tprint_char('\t');\n")
			}
		case *astExpr:
			previousIsSemicolon = false
			type_ = cmd.finalType()
			if type_ == voidType {
				fmt.Fprintf(wr, "\t")
				cmd.generateC(wr)
				fmt.Fprintf(wr, ";\n")
			} else {
				fmt.Fprintf(wr, "\tprint_%s(", type_)
				cmd.generateC(wr)
				fmt.Fprintf(wr, ");\n")
			}
		}
	}
	if !previousIsSemicolon {
		fmt.Fprintf(wr, "\tprint_char('\\n');\n")
	}
}
