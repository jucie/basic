package main

import (
	"fmt"
	"os"
)

type variable struct {
	def  *astVarRef
	dims int
	ref  []*astVarRef
}

type solver struct {
	p        *program
	dsts     []*targetLine
	types    map[string]int
	funcs    map[string]int
	predefs  map[token]int
	vars     map[string]*variable
	notReady map[string]int
}

func newSolver(p *program) *solver {
	return &solver{
		p:        p,
		types:    make(map[string]int),
		funcs:    make(map[string]int),
		predefs:  make(map[token]int),
		vars:     make(map[string]*variable),
		notReady: make(map[string]int),
	}
}

func (s *solver) consider(h host) {
	switch v := h.(type) {
	case *astFnCall:
		s.funcs[v.id]++
	case *astPredef:
		s.predefs[v.function]++
	case astType:
		t := fmt.Stringer(v)
		s.types[t.String()]++
	case *astVarRef:
		vv, ok := s.vars[v.id]
		if !ok {
			vv = &variable{}
			s.vars[v.id] = vv
		}
		vv.ref = append(vv.ref, v)
		vv.dims = len(v.index)
	case cmdDim:
		for _, def := range v.vars {
			vv, ok := s.vars[def.id]
			if !ok {
				vv = &variable{}
				s.vars[def.id] = vv
			}
			if vv.def != nil {
				fmt.Fprintf(os.Stderr, "Multiple definition for variable %s.", def.id)
			}
			vv.def = def
			vv.dims = len(def.index)
		}
	case cmdDef:
		s.funcs[v.id]++
	case cmdGoto:
		s.dsts = append(s.dsts, &v.dst)
	case cmdGosub:
		s.dsts = append(s.dsts, &v.dst)
	case cmdOn:
		for _, dst := range v.dsts {
			s.dsts = append(s.dsts, &dst)
		}
	default:
		s.notReady[fmt.Sprintf("%T", h)]++
	}
}

func (s *solver) showStats() {
	if len(s.vars) > 0 {
		println("\nVars")
		for key, val := range s.vars {
			fmt.Printf("\t%s dims %d refs %v\n", key, val.dims, val.ref)
		}
	}

	if len(s.types) > 0 {
		println("\nTypes")
		for key, val := range s.types {
			println("\t", key, val)
		}
	}

	if len(s.funcs) > 0 {
		println("\nFunctions")
		for key, val := range s.funcs {
			println("\t", key, val)
		}
	}

	if len(s.predefs) > 0 {
		println("\nPredefs")
		for key, val := range s.predefs {
			println("\t", predefs[key].name, val)
		}
	}
}

func (s *solver) showNotReady() {
	for key, _ := range s.notReady {
		println("Solver not ready for type ", key)
	}
}

func (s *solver) linkLines(lines progLines) {
	m := make(map[int]*progLine)
	for _, l := range lines {
		m[l.id] = l
	}
	for i, _ := range s.dsts {
		dst := s.dsts[i]
		l, ok := m[dst.nbr]
		if !ok {
			l = lines.find(dst.nbr)
		}
		if l == nil {
			fmt.Fprintf(os.Stderr, "Target line not found: %d\n", dst.nbr)
		} else {
			dst.adr = l
		}
	}
}
