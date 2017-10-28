package main

import (
	"fmt"
	"os"
)

type variable struct {
	def        *astVarRef
	dims       int
	ref        []*astVarRef
	isForIndex bool
}

type function struct {
	id  string
	def *cmdFnDef
	ref []*astFnCall
}

type solver struct {
	p       *program
	dsts    []*targetLine
	types   map[string]int
	funcs   map[string]*function
	predefs map[token]int
	vars    map[string]*variable
}

func newSolver(p *program) *solver {
	return &solver{
		p:       p,
		types:   make(map[string]int),
		funcs:   make(map[string]*function),
		predefs: make(map[token]int),
		vars:    make(map[string]*variable),
	}
}

var mt = make(map[string]int)

func (s *solver) consider(h host) {
	mt[fmt.Sprintf("%T", h)]++
	switch v := h.(type) {
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
	case *cmdDim:
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
	case *cmdFor:
		vv, ok := s.vars[v.index.id]
		if !ok {
			vv = &variable{}
			s.vars[v.index.id] = vv
		}
		vv.isForIndex = true
	case *cmdFnDef:
		vv, ok := s.funcs[v.id]
		if !ok {
			vv = &function{id: v.id}
			s.funcs[v.id] = vv
		}
		if vv.def != nil {
			fmt.Fprintf(os.Stderr, "Multiple definition for function %s.", v.id)
		}
		vv.def = v
	case *astFnCall:
		vv, ok := s.funcs[v.id]
		if !ok {
			vv = &function{id: v.id}
			s.funcs[v.id] = vv
		}
		vv.ref = append(vv.ref, v)
	case *cmdGo:
		s.dsts = append(s.dsts, &v.dst)
	case *cmdOn:
		for _, dst := range v.dsts {
			s.dsts = append(s.dsts, &dst)
		}
	}
}

func (s *solver) showStats() {
	if len(s.vars) > 0 {
		println("\nVars:", len(s.vars))
		for key, val := range s.vars {
			fmt.Printf("\t%s dims %d refs %d\n", key, val.dims, len(val.ref))
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
			fmt.Printf("\t%s refs %d\n", key, len(val.ref))
		}
	}

	if len(s.predefs) > 0 {
		println("\nPredefs")
		for key, val := range s.predefs {
			println("\t", predefs[key].name, val)
		}
	}
}

func (s *solver) linkLines(lines progLines) {
	m := make(map[int]*progLine)
	for _, l := range lines {
		m[l.id] = l
	}
	for i := range s.dsts {
		dst := s.dsts[i]
		l, ok := m[dst.nbr]
		if !ok {
			l = lines.find(dst.nbr)
		}
		if l == nil {
			fmt.Fprintf(os.Stderr, "Target line not found: %d\n", dst.nbr)
		} else {
			dst.adr = l
			l.isDst = true
		}
	}
}
