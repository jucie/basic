package main

import (
	"bufio"
	"fmt"
)

type parser struct {
	lex *lexer
}

func newParser(rd *bufio.Reader) *parser {
	lex := newLexer(rd)
	return &parser{lex: lex}
}

func (p *parser) parseProgram(prog *program) {
	fmt.Printf("Analising program %s\n", prog.srcPath)
	for p.lex.peek().token != tokEof {
		fmt.Printf("%v\n", p.lex.peek())
		p.lex.next()
	}
	fmt.Println("")
	/*
		for p.lex.peek().token == tokInt {
			pl := &progLine{id: strconv.Atoi(p.lex.peek().s)}
		}
	*/
}
