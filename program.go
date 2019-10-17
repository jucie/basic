package main

import (
	"bufio"
	"fmt"
	"sort"
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
}

func (p program) receive(g guest) {
	for _, l := range p.lines {
		g.visit(l)
	}
}

func (p program) generateC(wr *bufio.Writer) {
	fmt.Fprintf(wr, "#include \"basiclib.h\"\n\n")
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

func (p *program) generateCDataDefinitions(wr *bufio.Writer, type_ astType) {
	size := p.dataCounter[type_]
	fmt.Fprintf(wr, "const size_t data_area_for_%s_cnt=%d;\n", type_, size)
	fmt.Fprintf(wr, "const %s data_area_for_%s[%d]={\n", type_, type_, size)
	scan(p, func(h host) {
		switch v := h.(type) {
		case *cmdData:
			v.generateCDefinition(wr, type_)
		}
	})
	fmt.Fprintf(wr, "};\n\n")
}

func (p *program) generateCFunctionDeclarations(wr *bufio.Writer) {
	b := false
	scan(p, func(h host) {
		switch v := h.(type) {
		case *cmdFnDef:
			v.generateCDeclaration(wr)
			b = true
		}
	})
	if b {
		wr.WriteRune('\n')
	}
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
	p.generateCFunctionDeclarations(wr)
	p.generateCVarDefinitions(wr)
	fmt.Fprintf(wr, "static str temp_str[];\n\n")
}

func (p *program) generateCVarDefinitions(wr *bufio.Writer) {
	m := make(map[string]*astVarRef)
	scan(p, func(h host) {
		switch v := h.(type) {
		case *astVarRef:
			m[v.nameForC()] = v
		}
	})
	var names []string
	for k := range m {
		names = append(names, k)
	}
	sort.Strings(names)
	for _, k := range names {
		v := m[k]
		if v.isArray() {
			fmt.Fprintf(wr, "static arr %s_var;\n", k)
		} else {
			fmt.Fprintf(wr, "static %s %s;\n", v.finalType(), k)
		}
	}
	wr.WriteRune('\n')

	for _, k := range names {
		v := m[k]
		if v.isArray() {
			typeString := v.finalType().String()
			fmt.Fprintf(wr, "static %s *%s(", typeString, k)
			for i := 0; i != len(v.index); i++ {
				if i != 0 {
					wr.WriteRune(',')
				}
				fmt.Fprintf(wr, "num index%d", i)
			}
			fmt.Fprintf(wr, "){ return %s_in_array(&%s_var,%d", typeString, k, len(v.index))
			for i := 0; i != len(v.index); i++ {
				fmt.Fprintf(wr, ",index%d", i)
			}
			fmt.Fprintf(wr, ");}\n")
		}
	}
	wr.WriteRune('\n')
}

func (p *program) incrementDataCounter(type_ astType) {
	p.dataCounter[type_]++
}
