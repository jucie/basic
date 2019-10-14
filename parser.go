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

func (p *parser) parseCmd() (cmd, bool) {
	l := p.lex.peek()
	tok := l.token

	switch tok {
	case tokLet:
		p.lex.next()
		fallthrough
	case tokID:
		{
			c := p.parseLet()
			return c, c != nil
		}
	case tokData:
		{
			c := p.parseData()
			return c, c != nil
		}
	case tokDef:
		{
			p.lex.next()
			c := p.parseDef()
			return c, c != nil
		}
	case tokDim:
		{
			p.lex.next()
			c := p.parseDim()
			return c, c != nil
		}
	case tokEnd:
		{
			p.lex.next()
			c := p.parseEnd()
			return c, c != nil
		}
	case tokFor:
		{
			p.lex.next()
			c := p.parseFor()
			return c, c != nil
		}
	case tokGo:
		{
			p.lex.next()
			c := p.parseGo()
			return c, c != nil
		}
	case tokIf:
		{
			p.lex.next()
			c := p.parseIf()
			return c, c != nil
		}
	case tokInput:
		{
			p.lex.next()
			c := p.parseInput()
			return c, c != nil
		}
	case tokNext:
		{
			p.lex.next()
			c := p.parseNext()
			return c, c != nil
		}
	case tokOn:
		{
			p.lex.next()
			c := p.parseOn()
			return c, c != nil
		}
	case tokPrint:
		{
			p.lex.next()
			c := p.parsePrint()
			return c, c != nil
		}
	case tokRead:
		{
			p.lex.next()
			c := p.parseRead()
			return c, c != nil
		}
	case tokRem:
		{
			c := p.parseRem()
			p.lex.next()
			return c, c != nil
		}
	case tokRestore:
		{
			p.lex.next()
			c := p.parseRestore()
			return c, c != nil
		}
	case tokReturn:
		{
			p.lex.next()
			c := p.parseReturn()
			return c, c != nil
		}
	case tokRun:
		{
			p.lex.next()
			c := p.parseRun()
			return c, c != nil
		}
	case tokStop:
		{
			p.lex.next()
			c := p.parseStop()
			return c, c != nil
		}
	case tokEOL:
		fallthrough
	case tokEOF:
		return nil, true
	default:
		p.unexpected()
		p.lex.consumeLine()
	}
	return nil, false
}

func (p *parser) isEndOfCommand() bool {
	switch p.lex.peek().token {
	case ':':
		fallthrough
	case tokEOL:
		fallthrough
	case tokEOF:
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

	fmt.Fprintf(os.Stderr, "%d:%d: Unexpected token (%d) \"%s\"\n",
		l.pos.row+1, l.pos.col+1, l.token, l.s)
}

func (p *parser) parseLineTail() []cmd {
	var result []cmd
	l := p.lex.peek()
	for {
		if l.token == tokEOL {
			break
		}
		cmd, ok := p.parseCmd()
		if !ok {
			p.unexpected()
			break
		}
		result = append(result, cmd)
		if l.token == ':' {
			p.lex.next() // skip separator
			continue
		} else if l.token == tokEOL {
			break
		} else {
			return nil
		}
	}
	return result
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
	p.lex.next()

	line := &progLine{id: id}
	line.cmds = p.parseLineTail()
	if l.token == tokEOL {
		p.lex.next()
	}
	return line
}

func (p *parser) parseProgram(prog *program) {
	p.prog = prog
	var previous *progLine
	for {
		line := p.parseLine()
		if line == nil {
			break
		}
		p.prog.lines = append(p.prog.lines, line)
		if previous != nil && line.id <= previous.id {
			fmt.Fprintf(os.Stderr, "%d: out of sequence. Previous is %d\n", line.id, previous.id)
		}
		previous = line
	}
}

func (p *parser) incrementDataCounter(type_ astType) {
	p.prog.incrementDataCounter(type_)
}
