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
	tok := l.token

	switch tok {
	case tokLet:
		p.lex.next()
		fallthrough
	case tokId:
		return p.parseLet()
	case tokData:
		p.lex.next()
		return p.parseData()
	case tokDef:
		p.lex.next()
		return p.parseDef()
	case tokDim:
		p.lex.next()
		return p.parseDim()
	case tokEnd:
		p.lex.next()
		return p.parseEnd()
	case tokFor:
		p.lex.next()
		return p.parseFor()
	case tokGosub:
		p.lex.next()
		return p.parseGosub()
	case tokGoto:
		p.lex.next()
		return p.parseGoto()
	case tokIf:
		p.lex.next()
		return p.parseIf()
	case tokInput:
		p.lex.next()
		return p.parseInput()
	case tokNext:
		p.lex.next()
		return p.parseNext()
	case tokOn:
		p.lex.next()
		return p.parseOn()
	case tokPrint:
		p.lex.next()
		return p.parsePrint()
	case tokRead:
		p.lex.next()
		return p.parseRead()
	case tokRem:
		p.lex.next()
		return p.parseRem()
	case tokRestore:
		p.lex.next()
		return p.parseRestore()
	case tokReturn:
		p.lex.next()
		return p.parseReturn()
	case tokStop:
		p.lex.next()
		return p.parseStop()
	case tokEol:
		fallthrough
	case tokEof:
		return nil
	default:
		p.unexpected()
		p.lex.consumeLine()
	}
	return nil
}

func (p *parser) isEndOfCommand() bool {
	switch p.lex.peek().token {
	case ':':
		fallthrough
	case tokEol:
		fallthrough
	case tokEof:
		return true
	}
	return false
}

func (p *parser) consumeCmd() {
	for !p.isEndOfCommand() {
		p.lex.next()
	}
}

func (p *parser) unexpected() {
	l := p.lex.lexeme

	fmt.Fprintf(os.Stderr, "%s (%d:%d): Unexpected token (%d) \"%s\"\n",
		p.prog.srcPath, l.pos.row+1, l.pos.col+1, l.token, l.s)
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
