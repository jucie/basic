package main

import (
	"bufio"
)

type cmd interface {
	host
}
type cmds []cmd

type progLine struct {
	id int
	cmds
	isDst    bool
	switchID int
}
type progLines []*progLine

func (l *progLine) receive(g guest) {
	for _, cmd := range l.cmds {
		g.visit(cmd)
	}
}

type program struct {
	lines       progLines
	ids         map[int]int
	dataCounter map[astType]int
	loopVars    []string
}

func newProgram() *program {
	return &program{ids: make(map[int]int), dataCounter: make(map[astType]int)}
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
	solver.linkForNext(p)
	p.lines = solver.linkLines(p.lines)
}

func (p program) receive(g guest) {
	for _, l := range p.lines {
		g.visit(l)
	}
}

func (p *program) incrementDataCounter(_type astType) {
	p.dataCounter[_type]++
}
