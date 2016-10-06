package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cmd interface {
	host
}
type cmds []cmd

type block struct {
	cmds
	pred blocks
	succ blocks
}
type blocks []*block

type progLine struct {
	id int
	cmds
	pred       blocks
	isDst      bool
	firstBlock *block
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
	p.generateDotFile()
}

func (p *program) generate() {
	// TODO
}

func (p program) receive(g guest) {
	for _, l := range p.lines {
		g.visit(l)
	}
}

func linkBlocks(pred, succ *block) {
	for _, bl := range succ.pred {
		if bl == pred {
			return // nothing to do
		}
	}
	succ.pred = append(succ.pred, pred)
}

func (p *program) newBlock(bl *block, shouldLink bool) *block {
	p.blocks = append(p.blocks, bl)

	newBlock := &block{}
	if shouldLink {
		linkBlocks(bl, newBlock)
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
			linkBlocks(outterBlock, bl)
		case *cmdGo:
			l := p.lines.find(c.dst.nbr)
			if l == nil {
				panic("coudn't find GOTO destination line")
			}
			l.pred = append(l.pred, bl)
			bl = p.newBlock(bl, c.sub)
		case *cmdOn:
			for _, dst := range c.dsts {
				l := p.lines.find(dst.nbr)
				if l == nil {
					panic("coudn't find ON GOTO destination line")
				}
				l.pred = append(l.pred, bl)
			}
			bl = p.newBlock(bl, true)
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

func linkBackwards(blocks blocks) {
	for _, bl := range blocks {
		bl.succ = nil
	}
	for _, bl := range blocks {
		for _, pred := range bl.pred {
			pred.succ = append(pred.succ, bl)
		}
	}
}

func (p *program) buildBlocks() {
	bl := &block{}
	for _, l := range p.lines {
		if l.isDst {
			bl = p.newBlock(bl, false)
			l.firstBlock = bl
		}
		bl = p.appendCmds(bl, l.cmds)
	}
	p.blocks = append(p.blocks, bl)

	for _, l := range p.lines {
		if l.isDst {
			if len(l.pred) == 0 {
				panic(fmt.Sprintf("%s: Destination %d has no predecessors", p.srcPath, l.id))
			}
			for _, prBlock := range l.pred {
				linkBlocks(prBlock, l.firstBlock)
			}
		}
	}

	if len(p.blocks) == 0 {
		return
	}

	linkBackwards(p.blocks)

	var curr *block
	var v blocks
	for _, bl := range p.blocks {
		if curr == nil {
			curr = bl
			continue
		}
		if len(bl.pred) == 1 && len(curr.succ) == 1 && curr.succ[0] == bl {
			curr.cmds = append(curr.cmds, bl.cmds...)
			curr.succ = bl.succ
		} else {
			v = append(v, curr)
			curr = bl
		}
	}
	v = append(v, curr)
	p.blocks = v

	linkBackwards(p.blocks)
}

func (p *program) generateDotFile() {
	f, err := os.Create(p.srcPath + ".dot")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wr := bufio.NewWriter(f)
	fmt.Fprintf(wr, "digraph G {\n")
	defer wr.Flush()

	for _, bl := range p.blocks {
		for _, pred := range bl.pred {
			fmt.Fprintf(wr, "\t\"%p\" -> \"%p\"\n", pred, bl)
		}
	}
	fmt.Fprintf(wr, "}\n")
}
