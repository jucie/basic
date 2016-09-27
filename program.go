package main

import (
	"bufio"
	"os"
	"strings"
)

type cmd interface {
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

type program struct {
	srcPath  string
	dstPath  string
	lines    []*progLine
	mapLines map[int]*progLine
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
	p.mapLines = make(map[int]*progLine)
	for _, l := range p.lines {
		p.mapLines[l.id] = l
	}
	// TODO
}

func (p *program) generate() {
	// TODO
}
