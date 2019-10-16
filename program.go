package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
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
	lines       progLines
	ids         map[int]int
	dataCounter map[astType]int
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
	solver.linkLines(p.lines)
	solver.showStats()
}

func (p program) receive(g guest) {
	for _, l := range p.lines {
		g.visit(l)
	}
}

func (p program) generateC(wr *bufio.Writer) {
	p.generateCPrologue(wr)
	fmt.Fprintf(wr, "int main(){\nint target = 0;\n")
	fmt.Fprintf(wr, "for(;;){switch (target){\ncase 0:\n")
	p.lines.generateC(wr)
	fmt.Fprintf(wr, "case %d: exit(0);\n", createLabel())
	fmt.Fprintf(wr, "default: fprintf(stderr, \"Undefined target line %s\", target); abort(); \n}}\n} /* main */\n\n", "%d")

	p.generateCFunctionDefinitions(wr)
	p.generateCDataDefinitions(wr, strType)
	p.generateCDataDefinitions(wr, numType)
	fmt.Fprintf(wr, "static str temp_str[%d];\n", createTemp())
}

func (p *program) generateCDataDeclarations(wr *bufio.Writer, type_ astType) {
	macro := fmt.Sprintf("SIZE_%s_DATA", strings.ToUpper(type_.String()))
	fmt.Fprintf(wr, "#define %s %d\n", macro, p.dataCounter[type_])
	fmt.Fprintf(wr, "static const %s data_area_for_%s[%s], *data_ptr = data_area_for_%s;\n", type_, type_, macro, type_)
}

func (p *program) generateCDataDefinitions(wr *bufio.Writer, type_ astType) {
	macro := fmt.Sprintf("SIZE_%s_DATA", strings.ToUpper(type_.String()))
	fmt.Fprintf(wr, "static const %s data_area_for_%s[%s]={\n", type_, type_, macro)
	scan(p, func(h host) {
		switch v := h.(type) {
		case *cmdData:
			v.generateCDefinition(wr, type_)
		}
	})
	fmt.Fprintf(wr, "};\n\n")
}

func (p *program) generateCFunctionDeclarations(wr *bufio.Writer) {
	scan(p, func(h host) {
		switch v := h.(type) {
		case *cmdFnDef:
			v.generateCDeclaration(wr)
		}
	})
	wr.WriteRune('\n')

}

func (p *program) generateCFunctionDefinitions(wr *bufio.Writer) {
	scan(p, func(h host) {
		switch v := h.(type) {
		case *cmdFnDef:
			v.generateCDefinition(wr)
		}
	})
}

func (p *program) generateCPrologue(wr *bufio.Writer) {
	p.generateCDataDeclarations(wr, strType)
	p.generateCDataDeclarations(wr, numType)
	p.generateCFunctionDeclarations(wr)
	p.generateCVarDefinitions(wr)
	fmt.Fprintf(wr, "static str temp_str[];\n")
}

func (p *program) generateCVarDefinitions(wr *bufio.Writer) {
	m := make(map[string]struct{})
	scan(p, func(h host) {
		switch v := h.(type) {
		case *astVarRef:
			key := v.finalType().String() + " "
			if v.isArray() {
				key += "*"
			}
			key += v.nameForC()
			m[key] = struct{}{}
		}
	})
	var names []string
	for k := range m {
		names = append(names, k)
	}
	sort.Strings(names)
	for _, k := range names {
		fmt.Fprintf(wr, "static %s;\n", k)
	}
	wr.WriteRune('\n')
}

func (p *program) incrementDataCounter(type_ astType) {
	p.dataCounter[type_]++
}
