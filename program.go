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

type block struct {
	label string
	cmds
	pred blocks
	succ blocks
}
type blocks []*block

func (bl block) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "block_type block_%s(void){\n", bl.label)
	bl.cmds.generateC(wr)
	fmt.Fprintf(wr, "}\n\n")
}

func (bls blocks) generateC(wr *bufio.Writer) {
	for _, bl := range bls {
		bl.generateC(wr)
	}
}

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
	lines progLines
	ids   map[int]int
	blocks
	orphans int
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
	p.buildBlocks()
	p.removeEmptyBlocks()
	p.coalesceBlocks()
	solver.showStats()
}

func (p program) receive(g guest) {
	for _, l := range p.lines {
		g.visit(l)
	}
}

func linkBlocks(pred, succ *block) {
	if pred == nil || succ == nil {
		return
	}
	for _, bl := range succ.pred {
		if bl == pred {
			return // nothing to do
		}
	}
	succ.pred = append(succ.pred, pred)
}

func (p *program) newBlock(id int, bl *block, shouldLink bool) *block {
	p.addBlock(bl)
	p.ids[id]++
	newBlock := &block{label: fmt.Sprintf("%d_%d", id, p.ids[id])}
	if shouldLink {
		linkBlocks(bl, newBlock)
	}
	return newBlock
}

func (p *program) appendCmds(id int, bl *block, cmds cmds) *block {
	for _, cmd := range cmds {
		bl.cmds = append(bl.cmds, cmd)
		switch c := cmd.(type) {
		case *cmdIf:
			outterBlock := bl
			innerBl := p.newBlock(id, bl, true)
			innerBl = p.appendCmds(id, innerBl, c.cmds)
			bl = p.newBlock(id, innerBl, true)
			linkBlocks(outterBlock, bl)
		case *cmdGo:
			l := p.lines.find(c.dst.nbr)
			if l == nil {
				panic(fmt.Sprintf("coudn't find %s destination line", c.cmdName()))
			}
			l.pred = append(l.pred, bl)
			bl = p.newBlock(id, bl, c.sub)
		case *cmdOn:
			for _, dst := range c.dsts {
				l := p.lines.find(dst.nbr)
				if l == nil {
					panic("coudn't find ON GOTO destination line")
				}
				l.pred = append(l.pred, bl)
			}
			bl = p.newBlock(id, bl, true)
		case *cmdFor:
			bl = p.newBlock(id, bl, true)
		case *cmdNext:
			bl = p.newBlock(id, bl, true)
		case *cmdEnd:
			bl = p.newBlock(id, bl, false)
		case *cmdStop:
			bl = p.newBlock(id, bl, false)
		case *cmdReturn:
			bl = p.newBlock(id, bl, false)
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

func (bls blocks) orphans() int {
	var count int
	for _, bl := range bls {
		if len(bl.pred) == 0 {
			count++
		}
	}
	return count
}

func (p *program) addBlock(bl *block) {
	if bl == nil {
		return
	}
	p.blocks = append(p.blocks, bl)
}

func (p *program) buildBlocks() {
	var bl *block
	for _, l := range p.lines {
		if bl == nil || l.isDst {
			bl = p.newBlock(l.id, bl, false)
			l.firstBlock = bl
		}
		bl = p.appendCmds(l.id, bl, l.cmds)
	}
	p.addBlock(bl)

	for _, l := range p.lines {
		if l.isDst {
			if len(l.pred) == 0 {
				panic(fmt.Sprintf("Destination %d has no predecessors", l.id))
			}
			for _, prBlock := range l.pred {
				linkBlocks(prBlock, l.firstBlock)
			}
		}
	}

	linkBackwards(p.blocks)
	p.orphans = p.blocks.orphans()
}

func (p *program) coalesceBlocks() {
	if len(p.blocks) == 0 {
		return
	}

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

func (bl *block) removeBlock(tbr *block) {
	if tbr == bl {
		return
	}

	var newPred blocks
	for _, pred := range bl.pred {
		if pred == tbr {
			continue
		}
		newPred = append(newPred, pred)
	}
	bl.pred = newPred

	var newSucc blocks
	for _, succ := range bl.succ {
		if succ == tbr {
			continue
		}
		newSucc = append(newSucc, succ)
	}
	bl.succ = newSucc
}

func (p *program) removeBlock(tbr *block) {
	for _, bl := range p.blocks {
		bl.removeBlock(tbr)
	}
}

func (p *program) removeEmptyBlocks() {
	if len(p.blocks) == 0 {
		return
	}

	var v blocks
	for _, bl := range p.blocks {
		if len(bl.cmds) == 0 && len(bl.pred) == 0 {
			p.removeBlock(bl)
		} else {
			v = append(v, bl)
		}
	}
	p.blocks = v
}
