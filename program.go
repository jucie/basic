package main

import (
	"bufio"
	"os"
	"strings"
)

type cmd interface {
	host
}
type cmds []cmd

type progLine struct {
	id int
	cmds
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

func (l *progLine) receive(g guest) {
	for _, cmd := range l.cmds {
		g.visit(cmd)
	}
}

type block struct {
	cmds
	pred []*block
	succ []*block
}
type blocks []*block

type program struct {
	srcPath string
	dstPath string
	lines   progLines
	blocks
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
	solver := newSolver(p)
	scan(p, func(h host) {
		solver.consider(h)
	})
	solver.linkLines(p.lines)
	p.buildBlocks()
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

func (p *program) appendCmds(bl *block, cmds cmds) *block {
	for _, cmd := range cmds {
		bl.cmds = append(bl.cmds, cmd)
		switch c := cmd.(type) {
		case *cmdIf:
			p.blocks = append(p.blocks, bl)
			innerBl := &block{cmds: c.cmds}
			innerBl = p.appendCmds(innerBl, c.cmds)
			p.blocks = append(p.blocks, innerBl)
			bl = &block{}
		case *cmdGo:
			p.blocks = append(p.blocks, bl)
			bl = &block{}
		case *cmdReturn:
			p.blocks = append(p.blocks, bl)
			bl = &block{}
		case *cmdNext:
			p.blocks = append(p.blocks, bl)
			bl = &block{}
		case *cmdEnd:
			p.blocks = append(p.blocks, bl)
			bl = &block{}
		case *cmdStop:
			p.blocks = append(p.blocks, bl)
			bl = &block{}
		case *cmdFor:
			p.blocks = append(p.blocks, bl)
			bl = &block{}
		}
	}
	return bl
}

func (p *program) buildBlocks() {
	bl := &block{}
	for _, l := range p.lines {
		bl = p.appendCmds(bl, l.cmds)
	}
	p.blocks = append(p.blocks, bl)
}
