package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type parser struct {
	lex  *lexer
	prog *program
}

func newParser(rd *bufio.Reader) *parser {
	lex := newLexer(rd)
	return &parser{lex: lex}
}

func (p *parser) parseCmd() cmd {
	l := p.lex.peek()
	switch l.token {
	case tokId:
		return p.parseAssign()
	case tokData:
		return p.parseData()
	case tokDef:
		return p.parseDef()
	case tokDim:
		return p.parseDim()
	case tokEnd:
		return p.parseEnd()
	case tokFor:
		return p.parseFor()
	case tokGosub:
		return p.parseGosub()
	case tokGoto:
		return p.parseGoto()
	case tokIf:
		return p.parseIf()
	case tokInput:
		return p.parseInput()
	case tokLet:
		return p.parseLet()
	case tokNext:
		return p.parseNext()
	case tokPrint:
		return p.parsePrint()
	case tokRead:
		return p.parseRead()
	case tokRem:
		return p.parseRem()
	case tokRestore:
		return p.parseRestore()
	case tokReturn:
		return p.parseReturn()
	case tokStop:
		return p.parseStop()
	default:
		p.unexpected()
		p.lex.consumeLine()
	}
	return nil
}

func (p *parser) unexpected() {
	l := p.lex.peek()

	fmt.Fprintf(os.Stderr, "%s (%d:%d): Unexpected \"%s\".",
		p.prog.srcPath, l.pos.row+1, l.pos.col+1, l.s)
}

func (p *parser) parseLine() *progLine {
	l := p.lex.peek()
	if l.token != tokNumber {
		return nil
	}
	id, err := strconv.Atoi(l.s)
	if err != nil {
		panic(err)
	}
	p.lex.next() // line number

	fmt.Printf("%d\n", id)
	line := &progLine{id: id}
	for {
		cmd := p.parseCmd()
		if cmd == nil {
			break
		}
		line.cmds = append(line.cmds, cmd)
		l = p.lex.peek()
		if l.token == ':' {
			p.lex.next() // skip separator
			continue
		} else if l.token == tokEol {
			p.lex.next() // skip terminator
			break
		} else {
			p.unexpected()
			p.lex.nextLine()
			break
		}
	}
	return line
}

func (p *parser) parseProgram(prog *program) {
	p.prog = prog
	fmt.Printf("Parsing program %s\n", p.prog.srcPath)
	for {
		line := p.parseLine()
		if line == nil {
			break
		}
		p.prog.lines[line.id] = line
	}
}
