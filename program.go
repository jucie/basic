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

func (p *program) newBlock(bl *block, shouldLink bool) *block {
	newBlock := &block{}
	newBlock.pred = append(newBlock.pred, bl)
	if shouldLink {
		bl.succ = append(bl.succ, newBlock)
		p.blocks = append(p.blocks, bl)
	}
	return newBlock
}

func (p *program) appendCmds(bl *block, cmds cmds) *block {
	for _, cmd := range cmds {
		bl.cmds = append(bl.cmds, cmd)
		switch c := cmd.(type) {
		case *cmdIf:
			outterBlock := bl
			innerBl := p.newBlock(bl, true)
			innerBl = p.appendCmds(innerBl, c.cmds)
			bl = p.newBlock(innerBl, true)
			outterBlock.succ = append(outterBlock.succ, bl)
			bl.pred = append(bl.pred, outterBlock)
		case *cmdGo:
			if !c.sub {
				bl = p.newBlock(bl, false)
			}
		case *cmdFor:
			bl = p.newBlock(bl, true)
		case *cmdNext:
			bl = p.newBlock(bl, true)
		case *cmdEnd:
			bl = p.newBlock(bl, false)
		case *cmdStop:
			bl = p.newBlock(bl, false)
		case *cmdReturn:
			bl = p.newBlock(bl, false)
		}
	}
	return bl
}

func (p *program) buildBlocks() {
	bl := &block{}
	for _, l := range p.lines {
		if l.isDst {
			bl = p.newBlock(bl, false)
		}
		bl = p.appendCmds(bl, l.cmds)
	}
}
