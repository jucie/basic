package main

import (
	"bufio"
	"fmt"
)

type cmd interface {
	host
	generateC(wr *bufio.Writer)
}
type cmds []cmd

func (cms cmds) generateC(wr *bufio.Writer) {
	for _, cmd := range cms {
		cmd.generateC(wr)
	}
}

type progLine struct {
	id int
	cmds
	isDst bool
}
type progLines []*progLine

func (pl progLines) find(dst int) *progLine {
	for _, l := range pl {
		if l.id >= dst {
			return l
		}
	}
	return nil
}

func (pl progLines) generateC(wr *bufio.Writer) {
	for _, l := range pl {
		l.generateC(wr)
	}
}

func (l *progLine) receive(g guest) {
	for _, cmd := range l.cmds {
		g.visit(cmd)
	}
}

func (l *progLine) generateC(wr *bufio.Writer) {
	if l.isDst {
		fmt.Fprintf(wr, "case %d: ", l.id)
	}
	fmt.Fprintf(wr, "/* line %d */\n", l.id)
	l.cmds.generateC(wr)
}

type program struct {
	lines progLines
	ids   map[int]int
}

func newProgram() *program {
	return &program{ids: make(map[int]int)}
}

func (p *program) loadFrom(src *bufio.Reader) {
	parser := newParser(src)
	parser.parseProgram(p)
	p.resolve()
}

func (p *program) resolve() {
	solver := newSolver(p)
	scan(p, func(h host) {
		solver.consider(h)
	})
	solver.linkLines(p.lines)
	solver.showStats()
}

func (p program) receive(g guest) {
	for _, l := range p.lines {
		g.visit(l)
	}
}

func (p program) generateC(wr *bufio.Writer) {
	p.lines.generateC(wr)
}
