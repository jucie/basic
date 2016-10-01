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
	case tokId:
		{
			c := p.parseLet()
			return c, c != nil
		}
	case tokData:
		{
			p.lex.next()
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
	case tokGosub:
		{
			p.lex.next()
			c := p.parseGosub()
			return c, c != nil
		}
	case tokGoto:
		{
			p.lex.next()
			c := p.parseGoto()
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
			p.lex.next()
			c := p.parseRem()
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
	case tokEol:
		fallthrough
	case tokEof:
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

func (p *parser) parseLineTail() []cmd {
	var result []cmd
	l := p.lex.peek()
	for {
		if l.token == tokEol {
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
		} else if l.token == tokEol {
			p.lex.next() // skip terminator
			break
		} else {
			return nil
		}
	}
	if l.token == tokEol {
		p.lex.next()
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
	return line
}

func (p *parser) parseProgram(prog *program) {
	p.prog = prog
	fmt.Printf("Parsing program %s\n", p.prog.srcPath)
	var previous *progLine
	for {
		line := p.parseLine()
		if line == nil {
			break
		}
		p.prog.lines = append(p.prog.lines, line)
		if previous != nil && line.id <= previous.id {
			fmt.Fprintf(os.Stderr, "%s:%d: out of sequence. Previous is %d\n", prog.srcPath, line.id, previous.id)
		}
		previous = line
	}
}
