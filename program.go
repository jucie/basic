package main

import (
	"bufio"
	"os"
	"strings"
)

type cmd interface {
	host
}

type progLine struct {
	id   int
	cmds []cmd
	succ []*progLine
	pred []*progLine
}

func (l *progLine) linkSucc(succ *progLine) {
	l.succ = append(l.succ, succ)
	succ.pred = append(succ.pred, l)
}

func (l *progLine) receive(g guest) {
	for _, cmd := range l.cmds {
		g.visit(cmd)
	}
}

type program struct {
	srcPath  string
	dstPath  string
	lines    []*progLine
	mapLines map[int]int
}

func newProgram() *program {
	return &program{}
}

func loadProgram(path string) *program {
	prog := newProgram()

	prog.srcPath = path
	pos := strings.LastIndexByte(path, '.')
	if pos < 0 {
		pos = len(path)
	}
	prog.dstPath = path[:pos] + ".exe"

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	newParser(rd).parseProgram(prog)
	return prog
}

func (p *program) resolve() {
	p.mapLines = make(map[int]int)
	for i, l := range p.lines {
		p.mapLines[l.id] = i
	}
	solver := newSolver(p)
	solver.visit(p)
	//solver.showStats()
	//solver.showNotReady()
}

func (p *program) generate() {
	// TODO
}

func (p program) receive(g guest) {
	for _, l := range p.lines {
		g.visit(l)
	}
}
